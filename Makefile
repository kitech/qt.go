# make [target] [V=0|1|2 P=n MIN=x]

# GOPATH := $(PWD):$(GOPATH)
MINORVER := $(shell go version|awk '{print $$3}'|awk -F. '{print $$2}')
USEI=$(shell expr "${MINORVER}" ">=" "10")
ifeq ($(USEI), 1)
		ARGI=-i
		# ARGI=-i -tags minimal
endif
ifeq ($(V), 1)
	ARGV=-v
endif
ifeq ($(V), 2)
	ARGV=-v -x
endif
ifneq ($(P), )
	ARGP=-p $(P)
endif
ifneq ($(MIN), )
	ARGMIN=-tags minimal
endif
ifneq ($(GOSO), )
	ARGGOSO=-pkgdir ~/oss/pkg/linux_amd64 -ldflags "-w -s"
	ARGGOSOLIB=${ARGGOSO} -buildmode=shared
	ARGGOSOEXE=${ARGGOSO} -linkshared
endif
ARGA=${ARGI} ${ARGV} ${ARGP} ${ARGMIN} ${ARGGOSOLIB}
ARGEXE=${ARGI} ${ARGV} ${ARGP} ${ARGMIN} # ${ARGGOSOEXE} # For RAM seems not good

all:
	@echo ${GOPATH}
	@echo ${MINORVER}
	@echo "dd ${ARGA} dd"
	@echo "${V} "
	@echo "all target do nothing, valid targets: realall"
	# CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

realall: rts bases qmls extras webengines multimedias tools

rts: qtrt- mock-
qtrt-:
	go install ${ARGA} ./qtqt
	go install ${ARGA} ./qtrt
mock-:
	go install ${ARGA} ./qtmock

bases: qtrt- core- gui- widgets-
core-:
	go install ${ARGA} ./qtcore
gui-:
	go install ${ARGA} ./qtgui
widgets-:
	go install ${ARGA} ./qtwidgets

network-:
	go install ${ARGA} ./qtnetwork

qmls: qml- quick- quickctrl- quickwgt-
qml-:
	go install ${ARGA} ./qtqml
quick-:
	go install ${ARGA} ./qtquick
quickctrl-:
	go install ${ARGA} ./qtquicktemplates2
	go install ${ARGA} ./qtquickcontrols2
quickwgt-:
	go install ${ARGA} ./qtquickwidgets

extras:
	go install ${ARGA} ./qtandroidextras
	go install ${ARGA} ./qtmacextras
	go install ${ARGA} ./qtwinextras

webengines:
	go install ${ARGA} ./qtpositioning
	go install ${ARGA} ./qtprintsupport
	go install ${ARGA} ./qtwebchannel
	go install ${ARGA} ./qtwebenginecore
	go install ${ARGA} ./qtwebengine
	go install ${ARGA} ./qtwebenginewidgets

multimedias:
	go install ${ARGA} ./qtsvg
	go install ${ARGA} ./qtmultimedia

cleansos:
	rm -vf ~/oss/pkg/linux_amd64/libgithub.com-kitech-qt.go-*.so

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
	go build ${ARGEXE} -o bin/go-uic ./cmd/go-uic
	go build ${ARGEXE} -o bin/go-rcc ./cmd/go-rcc
	go build ${ARGEXE} -o bin/cgo-rcc ./cmd/cgo-rcc
	go build ${ARGEXE} -o bin/go-qmlviewer ./cmd/go-qmlviewer
	go build ${ARGEXE} -o bin/go-dir2qrc ./cmd/dir2qrc

tst:
	go test ${ARGEXE} tests/qstring_test.go

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

