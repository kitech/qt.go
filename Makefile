
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
	CC=clang CXX=clang++ go install -v -x ./qtrt
mock-:
	CC=clang CXX=clang++ go install -v -x ./qtmock

bases: core- gui- widgets-
core-:
	CC=clang CXX=clang++ go install -v -x ./qtcore
gui-:
	CC=clang CXX=clang++ go install -v -x ./qtgui
widgets-:
	CC=clang CXX=clang++ go install -v -x ./qtwidgets

network-:
	CC=clang CXX=clang++ go install -v -x ./qtnetwork

qmls: qml- quick- quicktmpl- quickctrl- quickwgt-
qml-:
	CC=clang CXX=clang++ go install -v -x ./qtqml
quick-:
	CC=clang CXX=clang++ go install -v -x ./qtquick
quicktmpl-:
	CC=clang CXX=clang++ go install -v -x ./qtquicktemplates2
quickctrl-:
	CC=clang CXX=clang++ go install -v -x ./qtquickcontrols2
quickwgt-:
	CC=clang CXX=clang++ go install -v -x ./qtquickwidgets

extras:
	CC=clang CXX=clang++ go install -v -x ./qtandroidextras
	CC=clang CXX=clang++ go install -v -x ./qtmacextras
	CC=clang CXX=clang++ go install -v -x ./qtwinextras


eg-:
	# CC=clang CXX=clang++ go build -v -x eg/coreapp.go
	# CC=clang CXX=clang++ go build -v -x eg/guiapp.go
	# CC=clang CXX=clang++ go build -v -x eg/signal.go
	# CC=clang CXX=clang++ go build -v -x eg/pmthor.go
	# CC=clang CXX=clang++ go build -v -x -o bin/uicgen eg/uicgen.go eg/ui_xxx.go eg/rcc_rc.go
	CC=clang CXX=clang++ go build -v -x eg/button.go
	# CC=clang CXX=clang++ go build -v -x eg/eg00.go
	# CC=clang CXX=clang++ go build -v -x eg/eg10.go

tools:
	CC=clang CXX=clang++ go build -v -x -o bin/go-uic go-uic/uic-tph.go go-uic/codepager.go go-uic/util.go
	CC=clang CXX=clang++ go build -v -x -o bin/go-rcc go-uic/rcc-tph.go go-uic/codepager.go go-uic/util.go
	go build -v -o bin/go-qmlviewer go-uic/qmlviewer.go

tst:
	CC=clang CXX=clang++ go test -v -x tests/qstring_test.go

wcs: wcbases wcqmls wcextras
wcbases:
	wc -l qt{core,gui,widgets}/* | tail -n 1
wcqmls:
	wc -l qt{qml,quick,quickwidgets,network}/* | tail -n 1
wcextras:
	wc -l qt{androidextras,winextras,macextras}/* | tail -n 1

