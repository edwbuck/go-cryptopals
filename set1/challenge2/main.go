package main

import (
	"fmt"
	"os"

	"edwinbuck.com/set1/challenge2/pkg/cryptobuffer"
)

func main() {
	hexString := "1c0111001f010100061a024b53535009181c"
	xorString := "686974207468652062756c6c277320657965"

	buffer, err := cryptobuffer.FromHexString(hexString)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		os.Exit(1)
	}

	xor, err := cryptobuffer.FromHexString(xorString)
	if err != nil {
		fmt.Printf("error: %s\n", xor)
		os.Exit(1)
	}

	result, err := buffer.Xor(xor)
	if err != nil {
		fmt.Printf("error: %s\n", xor)
		os.Exit(1)
	}

	output, err := result.ToHexString()
	if err != nil {
		fmt.Printf("error: %s\n", xor)
		os.Exit(1)
	}

	fmt.Printf("%s\n", output)
	os.Exit(0)
}
