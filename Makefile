.PHONY: clean test

PWD=$(shell pwd)
GOPATH=$(PWD)
PY3LIB=$(shell pkg-config --cflags --libs python3)

test: pygo.so
	python3 -c 'import pygo, sys; print(pygo.sum(40, 2)); pygo.gil(); pygo.tick(); sys.stdin.read(1);'

pygo.so: py.go
	@go build -v -buildmode=c-shared -o pygo.so

clean:
	@rm -f pygo.so; true
	@rm -f pygo.h; true
