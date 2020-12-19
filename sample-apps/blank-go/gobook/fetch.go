// fetch - a Go clone of curl
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	const buffSize = 100
	const prefix = "http://"

	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, prefix) {
			url = fmt.Sprintf("%s%s", prefix, url)
		}

		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		buffer := make([]byte, buffSize)

		for {
			n, err := resp.Body.Read(buffer)

			if err != nil {
				if err != io.EOF {
					fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
				}
				break
			}
			_, errw := os.Stdout.Write(buffer[:n])

			if errw != nil {
				fmt.Fprintf(os.Stderr, "fetch: writing failed %v\n", errw)
				break
			}

		}

		resp.Body.Close()

	}
}
