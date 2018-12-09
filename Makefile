
# GOPATH := $(PWD):$(GOPATH)
MINORVER := $(shell go version|awk '{print $$3}'|awk -F. '{print $$2}')
USEI=$(shell expr "${MINORVER}" ">=" "10")
ifeq ($(USEI), 1)
	ARGI="-i"
endif

all:
	@echo ${GOPATH}
	@echo ${MINORVER}
	@echo "dd ${ARGI} dd"
	# CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

realall: rts bases qmls extras webengines multimedias tools

rts: qtrt- mock-
qtrt-:
	go install -v -x ${ARGI} ./qtqt
	go install -v -x ${ARGI} ./qtrt
mock-:
	go install -v -x ${ARGI} ./qtmock

bases: qtrt- core- gui- widgets-
core-:
	go install -v -x ${ARGI} ./qtcore
gui-:
	go install -v -x ${ARGI} ./qtgui
widgets-:
	go install -v -x ${ARGI} ./qtwidgets

network-:
	go install -v -x ${ARGI} ./qtnetwork

qmls: qml- quick- quickctrl- quickwgt-
qml-:
	go install -v -x ${ARGI} ./qtqml
quick-:
	go install -v -x ${ARGI} ./qtquick
quickctrl-:
	go install -v -x ${ARGI} ./qtquicktemplates2
	go install -v -x ${ARGI} ./qtquickcontrols2
quickwgt-:
	go install -v -x ${ARGI} ./qtquickwidgets

extras:
	go install -v -x ${ARGI} ./qtandroidextras
	go install -v -x ${ARGI} ./qtmacextras
	go install -v -x ${ARGI} ./qtwinextras

webengines:
	go install -v -x ${ARGI} ./qtpositioning
	go install -v -x ${ARGI} ./qtprintsupport
	go install -v -x ${ARGI} ./qtwebchannel
	go install -v -x ${ARGI} ./qtwebenginecore
	go install -v -x ${ARGI} ./qtwebengine
	go install -v -x ${ARGI} ./qtwebenginewidgets

multimedias:
	go install -v -x ${ARGI} ./qtsvg
	go install -v -x ${ARGI} ./qtmultimedia

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
	go build -p 1 -v -i -o bin/go-uic ./cmd/go-uic
	go build -p 1 -v -i -o bin/go-rcc ./cmd/go-rcc
	go build -p 1 -v -i -o bin/cgo-rcc ./cmd/cgo-rcc
	go build -p 1 -v -i -o bin/go-qmlviewer ./cmd/go-qmlviewer
	go build -p 1 -v -i -o bin/go-dir2qrc ./cmd/dir2qrc

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
	curl -POST -d "path=github.com/kitech/qt.go/qtmeta" "https://godoc.org/-/refresh"

wcs: wcbases wcqmls wcextras wcwebengines
wcbases:
	wc -l qt{core,gui,widgets}/* | tail -n 1
wcqmls:
	wc -l qt{qml,quick,quickwidgets,network}/* | tail -n 1
wcextras:
	wc -l qt{androidextras,winextras,macextras}/* | tail -n 1
wcwebengines:
	wc -l qtweb{channel,enginecore,engine,enginewidgets}/* | tail -n 1

