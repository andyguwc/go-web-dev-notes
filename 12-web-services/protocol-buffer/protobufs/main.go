/* 

We import protocol buffer from the protofiles package. These are structus that are mapped to given protobuf in the proto files

We use the Person struct and initialized it. Then we serialized the struct using the proto.Marshal function

*/ 


package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
	pb "github.com/narenaryan/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Roger F",
		Email: "rf@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p, "\n")
	fmt.Println("Marshalled proto data: ", body, "\n")
	fmt.Println("Unmarshalled struct: ", p1)
}