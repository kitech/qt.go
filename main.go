package main

import "log"

func main() {
	// test()
	// main_a()
	app := NewQApplication5()
	log.Println(app)

	w := NewQWidget5()
	log.Println(w)
	w.Show()

	app.Exec5()
	// NewQWidget()
}
