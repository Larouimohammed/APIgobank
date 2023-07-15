package main

import (
	"fmt"
	"log"
	"github.com/anthdm/Gobank/storage"
"github.com/anthdm/Gobank/api"
)

func main() {
	store, err := storage.NewPostgresStore()
	if err != nil {log.Fatal("gdfgd",err)}
    if err := store.Init(); err != nil {log.Fatal(err)}
	fmt.Println("fgfd",store)
	server := api.NewAPIServer(":8080")
	err1 := server.Run()
	if err1 != nil {
	fmt.Println(err1)
	}
	fmt.Println("Hello, World!")
}
