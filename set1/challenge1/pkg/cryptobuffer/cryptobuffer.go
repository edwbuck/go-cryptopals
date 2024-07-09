package cryptobuffer

import (
	"errors"
)

const (
	base64chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
	paddingIndex = 64
)
type cryptobuffer []byte

func getHexValue(b byte) (int, error) {
	if 'a' <= b && b <= 'f' {
		return int(10 + b - 'a'), nil
	}
	if 'A' <= b && b <= 'F' {
		return int(10 + b - 'A'), nil
	}
	if '0' <= b && b <= '9' {
		return int(b - '0'), nil
	}
	return 0, errors.New("character out of hex digit range")
}

func FromHexString(input string) (cryptobuffer, error) {
	charCount := len(input)
	byteCount := (charCount + 1) / 2

	var newbuffer = make([]byte, byteCount)
	highNybble := true
	for idx, byt := range []byte(input) {
		value, err := getHexValue(byt)
		if err != nil {
			return nil, err
		}
		if highNybble {
			newbuffer[idx/2] |= byte(0x000000F0 & (value << 4))
			highNybble = false
		} else {
			newbuffer[idx/2] |= byte(0x0000000F & (value))
			highNybble = true
		}
	}
	return newbuffer, nil
}

func (c *cryptobuffer) ToBase64String() (string, error) {
	byteCount := len(*c)
	charCount := 4*((byteCount+2)/3)

	var base64buffer = make([]byte, charCount)
	byteIndex := 0
	for idx, _ := range base64buffer {
		if byteIndex >= byteCount {
			base64buffer[idx] = base64buffer[paddingIndex]
			continue
		}
		var base64idx byte
		switch idx % 4 {
			case 0:
				base64idx |= (0x000000FC & ((*c)[byteIndex])) >> 2
				base64buffer[idx] = base64chars[base64idx]
			case 1:
				base64idx |= (0x00000003 & ((*c)[byteIndex])) << 4
				base64idx |= (0x000000F0 & ((*c)[byteIndex+1])) >> 4
				base64buffer[idx] = base64chars[base64idx]
			case 2:
				base64idx |= (0x0000000F & ((*c)[byteIndex+1])) << 2
				base64idx |= (0x000000C0 & ((*c)[byteIndex+2])) >> 6
				base64buffer[idx] = base64chars[base64idx]
			case 3:
				base64idx |= (0x0000002F & ((*c)[byteIndex+2]))
				base64buffer[idx] = base64chars[base64idx]
				byteIndex = byteIndex + 3
		}
	}
	return string(base64buffer), nil
}
