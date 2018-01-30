package main

import (
	"log"
	"time"

	"qt.go/qtcore"
)

func main() {
	argv := []string{"./coreapp", "-v", "-x"}
	app := qtcore.NewQCoreApplication(len(argv), argv, 0)
	log.Println(app)
	go func() {
		time.Sleep(5 * time.Second)
		app.Exit(0)
	}()
	app.Exec()
}
