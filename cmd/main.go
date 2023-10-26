package main

import (
	app "github.com/Srgkharkov/test-game/internal/app"
	"log"
)

func main() {
	app := app.NewApp()
	if err := app.APIServer.Run(); err != nil {
		log.Fatal(err)
	}
	//router := sw.NewRouter()
	//
	//log.Fatal(http.ListenAndServe(":8080", router))

}
