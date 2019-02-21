package main

import(
	"crypto/rand"
	"fmt"
	)

func keyGenerate() {
    key := make([]byte, 32)

    _, err := rand.Read(key)
    if err != nil {
        // handle error here
    }
		fmt.Println(key)
}
