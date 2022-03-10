package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/BRO3886/grpc-go-example/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	option = flag.String("option", "create", "[create, list, add]")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := addressbook.NewAddressBookServiceClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	switch *option {
	case "create":
		createAddrBook(client, ctx)

	case "list":
		list(client, ctx)

	case "add":
		addPerson(client, ctx)
	}
}

func createAddrBook(client addressbook.AddressBookServiceClient, ctx context.Context) {
	resp, err := client.CreateAddressBook(ctx, &addressbook.AddressBook{
		People: []*addressbook.Person{
			&addressbook.Person{
				Name:  "John Doe",
				Id:    1,
				Email: "hello@email.com",
				Phones: []*addressbook.Person_PhoneNumber{
					&addressbook.Person_PhoneNumber{
						Number: "123456789",
						Type:   addressbook.Person_MOBILE,
					},
				},
			},
		},
	})
	if err != nil {
		log.Println("[create] error from server: ", err)
	}
	log.Println("[create] response from server: ", resp)
}

func addPerson(client addressbook.AddressBookServiceClient, ctx context.Context) {
	stream, err := client.AddPerson(ctx)
	if err != nil {
		log.Println("[add] opening stream: ", err)
	}

	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		fmt.Println("enter name: ")
		var name string
		fmt.Scanln(&name)

		fmt.Println("enter email: ")
		var email string
		fmt.Scanln(&email)

		id := r.Int31()

		var phone string
		fmt.Println("enter phone number: ")
		fmt.Scanln(&phone)

		p := addressbook.Person{
			Name:  name,
			Email: email,
			Id:    id,
			Phones: []*addressbook.Person_PhoneNumber{
				&addressbook.Person_PhoneNumber{
					Number: phone,
					Type:   addressbook.Person_MOBILE,
				},
			},
		}

		if err := stream.Send(&p); err != nil {
			log.Println("[add] error from server: ", err)
		}

		log.Println("[add] sent: ", p.Name)
	}
}

func list(client addressbook.AddressBookServiceClient, ctx context.Context) {
	resp, err := client.GetAddressBook(ctx, &emptypb.Empty{})
	if err != nil {
		log.Println("[list] error from server: ", err)
	}
	for _, person := range resp.People {
		log.Println("[list] person: ", person.Id, person.Name, person.Email)
	}
}
