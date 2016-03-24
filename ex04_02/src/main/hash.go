package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
)

type ArgError string

func (e ArgError) Error() string {
	return string(e)
}

func main() {

	if len(os.Args) <= 1 {
		fmt.Printf("Usage : %s <crypt_method>", os.Args[0])
		return
	}

	fmt.Print(">> ");

	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadSlice('\n')

	method := strings.ToUpper(os.Args[1])
	sum, err := crypt(method, string(word))

	if(err != nil) {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
	} else {
		fmt.Printf("%s >> %x\n", method, sum)
	}
}

func crypt(method, value string) ([]byte, *ArgError) {
	switch method {
		case "MD5":
			sum := md5.Sum([]byte(value))
			return sum[:], nil

		case "SHA224":
			sum := sha256.Sum224([]byte(value))
			return sum[:], nil

			return sum[:], nil
		case "SHA256":
			sum := sha256.Sum256([]byte(value))
			return sum[:], nil

		case "SHA384":
			sum := sha512.Sum512([]byte(value))
			return sum[:], nil

		case "SHA512":
			sum := sha512.Sum512([]byte(value))
			return sum[:], nil

		default:
			err := ArgError(fmt.Sprintf("Invalid argument : %s", method))
			return nil, &err
	}
}