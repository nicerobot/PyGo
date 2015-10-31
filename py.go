// https://blog.filippo.io/building-python-modules-with-go-1-5/
package main

/*

#cgo pkg-config: python3
#define Py_LIMITED_API
#include <Python.h>

int PyArg_ParseTuple_LL(PyObject *, long long *, long long *);

PyObject * sum(PyObject *, PyObject *);

*/
import "C"
import (
	"log"
	"time"
	"sync"
)

func init() {
	log.Println("ok")
	go func() {
	  for range time.NewTicker(time.Second).C {
      log.Println("tick")
	  }
	}()
}

//export sum
func sum(self, args *C.PyObject) *C.PyObject {
	var a, b C.longlong
	if C.PyArg_ParseTuple_LL(args, &a, &b) == 0 {
	  return nil
	}
	return C.PyLong_FromLongLong(a + b)
}

//export tick
func tick(self, args *C.PyObject) *C.PyObject {
	go func() {
	  for range time.NewTicker(time.Second).C {
      log.Println("pygo ticker")
	  }
	}()
	return C.PyLong_FromLong(0)
}

//export gil
func gil(self, args *C.PyObject) *C.PyObject {
	var res *C.PyObject

	tState := C.PyEval_SaveThread()

	var mu sync.Mutex
	mu.Lock()

	go func() {
	  C.PyEval_RestoreThread(tState)
	  res = C.PyLong_FromLong(1)
	  mu.Unlock()
	}()

	mu.Lock()

	return res
}

func main() {}
