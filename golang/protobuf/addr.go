package main

import "fmt"
import "log"
import "io/ioutil"
import "github.com/golang/protobuf/proto"
import pb "github.com/rchirakk/animated/golang/protobuf/addrBook"

var fileName = "serialized"

func decode() {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
	}
	person := &pb.Person{}
	if err := proto.Unmarshal(in, person); err != nil {
		log.Fatalf("failed to parse address book: %s", err)
	}
	fmt.Printf("decoded: %+v \n", person)
}

func main() {
	p := &pb.Person{
		Name:  "abc",
		Id:    123,
		Email: "def@woot.com",
		Number: []*pb.Person_PhoneNumber{
			{
				Number: "12345678",
				Ptype:  pb.Person_MOBILE,
			},
			{
				Number: "22345678",
				Ptype:  pb.Person_MOBILE,
			},
		},
	}

	out, err := proto.Marshal(p)
	if err != nil {
		log.Fatalf("failed to marshal %s\n", err)
	}
	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalf("failed to save to file %s \n", err)
	}
	fmt.Printf("serialized %+v \n", p)
	decode()
}
