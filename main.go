package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func main() {
	arg := os.Args[1]

	id, err := uuid.Parse(arg)
	if err != nil {
		parseBase64(arg)
		return
	}

	makeBase64(id)
}

func parseBase64(data string) {
	bs, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}

	parsed, err := uuid.FromBytes(bs)
	if err != nil {
		panic(err)
	}
	fmt.Println(parsed.String())
}

func makeBase64(id uuid.UUID) {
	bs, err := id.MarshalBinary()
	if err != nil {
		panic(err)
	}

	str := base64.StdEncoding.EncodeToString(bs)
	fmt.Println(str)
}
