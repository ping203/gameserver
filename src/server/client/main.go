package main

import (
	"server/client/internal"
)

func StringHash(s string) (hash uint16) {
	for _, c := range s {
		ch := uint16(c)
		hash = hash + ((hash) << 5) + ch + (ch << 7)
	}
	return
}

func main() {
	client := internal.Client{}
	client.Init()

	sig := make(chan bool)
	client.Run(sig)
}
