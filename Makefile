
GOPATH := $(PWD):$(GOPATH)

all:
	@echo ${GOPATH}
	# CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

qtrt:
	CC=clang CXX=clang++ go install -v -x qtrt

core:
	CC=clang CXX=clang++ go install -v -x qtcore

gui:
	CC=clang CXX=clang++ go install -v -x qtgui

widgets:
	CC=clang CXX=clang++ go install -v -x qtwidgets

network:
	CC=clang CXX=clang++ go install -v -x qtnetwork

qml:
	CC=clang CXX=clang++ go install -v -x qtqml

quick:
	CC=clang CXX=clang++ go install -v -x qtquick

eg:
	# CC=clang CXX=clang++ go build -v -x src/eg/coreapp.go
	# CC=clang CXX=clang++ go build -v -x src/eg/guiapp.go
	# CC=clang CXX=clang++ go build -v -x src/eg/signal.go
	CC=clang CXX=clang++ go build -v -x src/eg/uicgen.go src/eg/ui_xxx.go src/eg/rcc_rc.go
	# CC=clang CXX=clang++ go build -v -x src/eg/button.go
	# CC=clang CXX=clang++ go build -v -x src/eg/eg00.go
	# CC=clang CXX=clang++ go build -v -x src/eg/eg10.go

tst:
	CC=clang CXX=clang++ go test -v -x tests/qstring_test.go

