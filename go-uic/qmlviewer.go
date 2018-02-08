package main

import (
	"gopp"
	"io/ioutil"
	"log"
	"os"
	"strings"

	funk "github.com/thoas/go-funk"

	"qt.go/qtcore"
	"qt.go/qtqml"
	"qt.go/qtquickwidgets"
	"qt.go/qtwidgets"
)

func main() {
	file := "./main.qml"
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	bcc, err := ioutil.ReadFile(file)
	gopp.ErrPrint(err, file)
	if err != nil {
		os.Exit(1)
	}
	iswin := funk.Contains(strings.Split(string(bcc), "\n"), "Window {")

	app := qtwidgets.NewQApplication(len(os.Args), os.Args, 0)

	if iswin {
		log.Println("app mod qml:")
		qmle := qtqml.NewQQmlApplicationEngine(nil)
		qmle.AddImportPath("./")
		qmle.Load_1(file)
	} else {
		log.Println("rect mod qml:")
		qw := qtquickwidgets.NewQQuickWidget(nil)
		qw.Engine().AddImportPath("./")
		qw.SetSource(qtcore.NewQUrl_1(file, 0))
		qw.Show()
	}

	app.Exec()
}
