package main

import (
	fmt "fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	william := &Person{
		Name: "William",
		Age:  34,
		SocialFollowers: &SocialFollowers{
			Youtube: 2500,
			Twitter: 1400,
		},
	}

	data, err := proto.Marshal(william)
	if err != nil {
		log.Fatal("Marshalling error: ", err)
	}

	fmt.Println(data)

	newWilliam := &Person{}
	err = proto.Unmarshal(data, newWilliam)
	if err != nil {
		log.Fatal("Unmarshalling error:", err)
	}

	fmt.Println(newWilliam.GetAge())
	fmt.Println(newWilliam.GetName())
	fmt.Println(newWilliam.SocialFollowers.GetTwitter())
}
