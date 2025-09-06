package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"

	"prideflag.fun/src/database"
	"prideflag.fun/src/pages"
)

//go:embed public/*
var public embed.FS

func main() {
	db, ctx := database.InitDatabase()

	var (
		port = flag.Int("port", 3000, "Specifies the network port the application will use.")
	)

	flag.IntVar(port, "p", 3000, "Specifies the network port the application will use.")
	flag.Parse()

	http.HandleFunc("/", pages.Index)
	http.HandleFunc("/test", pages.Test(db, ctx))
	http.HandleFunc("/flag", pages.Flag(db, ctx, public))

	fileServer := http.FileServer(http.FS(public))
	http.Handle("/public/", fileServer)

	fmt.Printf("The HTTP server now runs on 0.0.0.0:%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
