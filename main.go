package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct {
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 9999)
	// resp.Body.Read(bs)

	// fmt.Println(string(bs))

	lw := logWriter{}

	// Copy does the same thing as the above by first creating a byte slice,
	// then passing the byte slice to the Reader,
	// before sending the byte slice to the Writer in os.Stdout
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))

	return len(bs), nil
}
