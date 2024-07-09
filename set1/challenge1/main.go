package main

import (
	"fmt"
	"os"

	"edwinbuck.com/set1/challenge1/pkg/cryptobuffer"
)

func main() {
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	buffer, err := cryptobuffer.FromHexString(hexString)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	for _, b := range []byte(buffer) {
		fmt.Printf("%0.2X", b)
	}
	fmt.Printf("\n")
	base64String, err := buffer.ToBase64String()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", base64String)
	os.Exit(0)
}
