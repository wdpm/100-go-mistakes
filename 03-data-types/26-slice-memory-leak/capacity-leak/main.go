package main

import (
	"fmt"
	"runtime"
)

func consumeMessages() {
	for {
		msg := receiveMessage()
		// Do something with msg
		storeMessageType(getMessageType(msg))
	}
}

func getMessageType(msg []byte) []byte {
	// The slicing operation on msg using msg[:5] creates a five-length slice. However, its
	// capacity remains the same as the initial slice.
	return msg[:5]
}

func getMessageTypeWithCopy(msg []byte) []byte {
	msgType := make([]byte, 5)
	copy(msgType, msg)
	// here msgType is better for memory-wise
	return msgType
}

func receiveMessage() []byte {
	return make([]byte, 1_000_000)
}

func storeMessageType([]byte) {}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
