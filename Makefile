
# GOPATH := $(PWD):$(GOPATH)

all:
	@echo ${GOPATH}
	# CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

realall: rts bases qmls extras tools

rts: qtrt- mock-
qtrt-:
	go install -v -x ./qtqt
	go install -v -x ./qtrt
mock-:
	go install -v -x ./qtmock

bases: qtrt- core- gui- widgets-
core-:
	go install -v -x ./qtcore
gui-:
	go install -v -x ./qtgui
widgets-:
	go install -v -x ./qtwidgets

network-:
	go install -v -x ./qtnetwork

qmls: qml- quick- quickctrl- quickwgt-
qml-:
	go install -v -x ./qtqml
quick-:
	go install -v -x ./qtquick
quickctrl-:
	go install -v -x ./qtquicktemplates2
	go install -v -x ./qtquickcontrols2
quickwgt-:
	go install -v -x ./qtquickwidgets

extras:
	go install -v -x ./qtandroidextras
	go install -v -x ./qtmacextras
	go install -v -x ./qtwinextras


eg-:
	# go build -v -x eg/coreapp.go
	# go build -v -x eg/guiapp.go
	# go build -v -x eg/signal.go
	# go build -v -x eg/pmthor.go
	# go build -v -x -o bin/uicgen eg/uicgen.go eg/ui_xxx.go eg/rcc_rc.go
	# go build -v -x eg/button.go
	# go build -v -x eg/eg00.go
	# go build -v -x eg/eg10.go
	go build -v -x -o bui bigui/*.go

tools:
	go build -v -x -o bin/go-uic go-uic/uic-tph.go go-uic/codepager.go go-uic/util.go
	go build -v -x -o bin/go-rcc go-uic/rcc-tph.go go-uic/codepager.go go-uic/util.go
	go build -v -o bin/go-qmlviewer go-uic/qmlviewer.go

tst:
	go test -v -x tests/qstring_test.go

updoc:
	curl -POST -d "path=github.com/kitech/qt.go/qtcore" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtgui" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtwidgets" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtnetwork" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtqml" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtquick" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtquickcontrols2" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtquickwidgets" "https://godoc.org/-/refresh"

wcs: wcbases wcqmls wcextras
wcbases:
	wc -l qt{core,gui,widgets}/* | tail -n 1
wcqmls:
	wc -l qt{qml,quick,quickwidgets,network}/* | tail -n 1
wcextras:
	wc -l qt{androidextras,winextras,macextras}/* | tail -n 1

