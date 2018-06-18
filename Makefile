
# GOPATH := $(PWD):$(GOPATH)

all:
	@echo ${GOPATH}
	# CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

realall: rts bases qmls extras webengines tools

rts: qtrt- mock-
qtrt-:
	go install -v -i -x ./qtqt
	go install -v -i -x ./qtrt
mock-:
	go install -v -i -x ./qtmock

bases: qtrt- core- gui- widgets-
core-:
	go install -v -i -x ./qtcore
gui-:
	go install -v -i -x ./qtgui
widgets-:
	go install -v -i -x ./qtwidgets

network-:
	go install -v -i -x ./qtnetwork

qmls: qml- quick- quickctrl- quickwgt-
qml-:
	go install -v -i -x ./qtqml
quick-:
	go install -v -i -x ./qtquick
quickctrl-:
	go install -v -i -x ./qtquicktemplates2
	go install -v -i -x ./qtquickcontrols2
quickwgt-:
	go install -v -i -x ./qtquickwidgets

extras:
	go install -v -i -x ./qtandroidextras
	go install -v -i -x ./qtmacextras
	go install -v -i -x ./qtwinextras

webengines:
	go install -v -i -x ./qtpositioning
	go install -v -i -x ./qtprintsupport
	go install -v -i -x ./qtwebchannel
	go install -v -i -x ./qtwebenginecore
	go install -v -i -x ./qtwebengine
	go install -v -i -x ./qtwebenginewidgets

multimedias:
	go install -v -i -x ./qtsvg
	go install -v -i -x ./qtmultimedia

eg-:
	go build -v -x eg/coreapp.go
	# go build -v -x eg/guiapp.go
	# go build -v -x eg/signal.go
	# go build -v -x eg/pmthor.go
	# go build -v -x -o bin/uicgen eg/uicgen.go eg/ui_xxx.go eg/rcc_rc.go
	# go build -v -x eg/button.go
	# go build -v -x eg/eg00.go
	# go build -v -x eg/eg10.go
	go build -v -x -o bui bigui/*.go

tools:
	go build -v -o bin/go-uic ./cmd/go-uic
	go build -v -o bin/go-rcc ./cmd/go-rcc
	go build -v -o bin/go-qmlviewer ./cmd/go-qmlviewer
	go build -v -o bin/go-dir2qrc ./cmd/dir2qrc

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
	curl -POST -d "path=github.com/kitech/qt.go/qtprintsupport" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtpositioning" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtwebchannel" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtwebenginecore" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtwebengine" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtwebenginewidgets" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtsvg" "https://godoc.org/-/refresh"
	curl -POST -d "path=github.com/kitech/qt.go/qtmultimedia" "https://godoc.org/-/refresh"

wcs: wcbases wcqmls wcextras wcwebengines
wcbases:
	wc -l qt{core,gui,widgets}/* | tail -n 1
wcqmls:
	wc -l qt{qml,quick,quickwidgets,network}/* | tail -n 1
wcextras:
	wc -l qt{androidextras,winextras,macextras}/* | tail -n 1
wcwebengines:
	wc -l qtweb{channel,enginecore,engine,enginewidgets}/* | tail -n 1

