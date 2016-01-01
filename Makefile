
GOPATH := $(PWD):$(GOPATH)

all:
	@echo ${GOPATH}
	go install -v -x qt5
	# go build -v -x core
	# go build -v -x gui
	# go build -v -x widgets

