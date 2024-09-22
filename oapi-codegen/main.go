package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	hc := http.Client{}

	c, err := NewClientWithResponses("http://localhost:4200", WithHTTPClient(&hc))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.GetUserWithResponse(context.Background(), 3)
	if err != nil {
		log.Fatal(err)
	}

	var user MainUser
	err = json.Unmarshal(resp.Body, &user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("forename: %s surname: %s", *user.Forename, *user.Surname)
}
