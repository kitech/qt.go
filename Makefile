all:

test:
	go test -v -run Gen0
	# g++ -o tsym.so -shared tsym/tsym.cpp -Itsym
