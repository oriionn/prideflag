package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	var (
		port = flag.Int("port", 3000, "Specifies the network port the application will use.")
	)

	flag.IntVar(port, "p", 3000, "Specifies the network port the application will use.")
	flag.Parse()

	fmt.Printf("The HTTP server now runs on 0.0.0.0:%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
