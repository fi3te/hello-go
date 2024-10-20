package main

import (
	"log"

	"github.com/google/uuid"
)

func main() {
	log.Println("random UUID:", uuid.NewString())

	givenUUID := "a24cc8d4-6d69-4e22-a08d-b"
	if err := uuid.Validate(givenUUID); err != nil {
		log.Printf("given UUID %v was not valid (%v)\n", givenUUID, err)
	}

	parsedUUID := uuid.MustParse("77d09d4e-2b50-4259-af41-138a0b0c1d48")
	log.Println("parsed UUID:", parsedUUID)
}
