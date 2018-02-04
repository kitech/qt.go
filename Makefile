
# GOPATH := $(PWD):$(GOPATH)

all:
	@echo ${GOPATH}
	# CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

qtrt-:
	CC=clang CXX=clang++ go install -v -x ./qtrt
ffiqt:
	CC=clang CXX=clang++ go install -v -x ./cffiqt
mock-:
	CC=clang CXX=clang++ go install -v -x ./qtmock

core-:
	CC=clang CXX=clang++ go install -v -x ./qtcore

gui-:
	CC=clang CXX=clang++ go install -v -x ./qtgui

widgets-:
	CC=clang CXX=clang++ go install -v -x ./qtwidgets

network-:
	CC=clang CXX=clang++ go install -v -x ./qtnetwork

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
	CC=clang CXX=clang++ go build -v -x -o bin/go-uic go-uic/uic-tph.go go-uic/codepager.go
	CC=clang CXX=clang++ go build -v -x -o bin/go-rcc go-uic/rcc-tph.go go-uic/codepager.go

tst:
	CC=clang CXX=clang++ go test -v -x tests/qstring_test.go

