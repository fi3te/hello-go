package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fi3te/hello-go/google-protobuf/user"
	"google.golang.org/protobuf/proto"
)

const filename = "data"

func main() {
	u := &user.User{
		Id:            1,
		Forename:      "Max",
		Surname:       "Mustermann",
		Age:           20,
		MailAddresses: []string{"max@muster.mann", "max@mustermann.de"},
	}

	bytes, err := proto.Marshal(u)
	if err != nil {
		log.Fatalln("failed to encode user", err)
	}

	if err := os.WriteFile(filename, bytes, os.ModeTemporary); err != nil {
		log.Fatalln("failed to write data to file", err)
	}
	bytes, err = os.ReadFile(filename)
	if err != nil {
		log.Fatalln("failed to read data from file", err)
	}

	u = &user.User{}
	if err := proto.Unmarshal(bytes, u); err != nil {
		log.Fatalln("failed to decode user", err)
	}
	fmt.Println(u.String())
}
