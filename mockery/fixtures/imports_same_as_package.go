package test

import (
	"github.com/TheMacies/mockery/mockery/fixtures/test"
)

type C int

type ImportsSameAsPackage interface {
	A() test.B
	B() KeyManager
	C(C)
}
