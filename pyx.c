#define Py_LIMITED_API
#include <Python.h>

PyObject * sum(PyObject *, PyObject *);

// Workaround missing variadic function support
// https://github.com/golang/go/issues/975
int PyArg_ParseTuple_LL(PyObject * args, long long * a, long long * b) {
  return PyArg_ParseTuple(args, "LL", a, b);
}

static PyMethodDef PyGoMethods[] = {
	{"sum", sum, METH_VARARGS, "Add two numbers."},
	{NULL, NULL, 0, NULL}
};

static struct PyModuleDef pygomodule = {
	PyModuleDef_HEAD_INIT, "pygo", NULL, -1, PyGoMethods
};

PyMODINIT_FUNC PyInit_pygo(void) {
	return PyModule_Create(&pygomodule);
}
