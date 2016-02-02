
GOPATH := $(PWD):$(GOPATH)

all:
	@echo ${GOPATH}
	CC=clang CXX=clang++ go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

qtrt:
	CC=clang CXX=clang++ go install -v -x qtrt

eg:
	CC=clang CXX=clang++ go build -v -x src/eg/eg00.go
	# CC=clang CXX=clang++ go build -v -x src/eg/eg10.go

tst:
	CC=clang CXX=clang++ go test -v -x tests/qstring_test.go

