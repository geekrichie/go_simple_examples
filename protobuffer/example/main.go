package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "go_simple_examples/protobuffer"
	"io/ioutil"
	"log"
)


func main() {
	 p := pb.Person{
	 	Id: 1234,
	 	Name : "John Doe",
	 	Email : "jdoe@example.com",
	 	Phones:[]*pb.Person_PhoneNumber{
	 		{
	 			Number: "555-4321",
	 			Type : pb.Person_HOME,
			},
		},
	 }
	 out, err := proto.Marshal(&p)
	 if err != nil {
	 	log.Fatalln("Failed to encode address book", err)
	 }
	 fmt.Println(out)
	 fname := "person.proto"
	 if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		 log.Fatalln("Failed to write address book:", err)
	 }
}