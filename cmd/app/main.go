package main

import (
	sw "github.com/Srgkharkov/test-game/internal/sw_codegen"
	"log"
	"net/http"
)

func main() {
	//app := NewApp()
	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))

}
