package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Running simple client...")
		simpleMain()
	} else {
		fmt.Printf("Running %s client...\n", os.Args[1])
		if os.Args[1] == "interceptor" {
			interceptorMain()
		} else if os.Args[1] == "http" {
			httpMain()
		} else if os.Args[1] == "tls" {
			tlsMain()
		} else {
			simpleMain()
		}
	}
}
