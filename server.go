package main

import (
	"context"
	"io"
	"log"
	"net"

	"github.com/BRO3886/grpc-go/addressbook"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	memory addressbook.AddressBook
)

type addressbookServer struct {
	addressbook.UnimplementedAddressBookServiceServer
}

func (s *addressbookServer) CreateAddressBook(c context.Context, req *addressbook.AddressBook) (*addressbook.AddressBook, error) {
	log.Println("revd request", req)
	memory = *req
	return req, nil
}

func (s *addressbookServer) AddPerson(stream addressbook.AddressBookService_AddPersonServer) error {
	for {
		person, err := stream.Recv()
		if err == io.EOF {
			log.Println("[add] client closed stream")
			break
		}

		if err != nil {
			return err
		}

		log.Println("received person", person)
		if person == nil {
			continue
		}

		memory.People = append(memory.People, person)

		log.Println("added")
	}

	return nil
}

func (s *addressbookServer) GetAddressBook(c context.Context, _ *emptypb.Empty) (*addressbook.AddressBook, error) {
	return &memory, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	addressbook.RegisterAddressBookServiceServer(s, &addressbookServer{})
	log.Println("Starting server on port 9000")

	log.Panic(s.Serve(lis))
}
