package main

import "testing"
import "reflect"
import "unsafe"
import "github.com/kitech/qt.go/qt5"

func init() {
	if false {
		print(reflect.TypeOf(0))
	}
	if false {
		print(unsafe.Pointer(uintptr(0)))
	}
}

// using go test tests/qxxx_test.go

func Test1(t *testing.T) {
	s := qt5.NewQString("vioug")
	if s == nil {
		t.Fail()
	}
}

func Test2(t *testing.T) {
	s := qt5.NewQString("vioug")
	if s == nil {
		t.Fail()
	}

	len := s.Length()
	if len.(int32) != 5 {
		t.Fail()
	}
}

func Test3(t *testing.T) {
	s := qt5.NewQString("vioug")
	if s == nil {
		t.Fail()
	}

	s2 := qt5.NewQString(s)
	if s2 == nil {
		t.Fail()
	}

	len := s2.Length()
	if len.(int32) != 5 {
		t.Fail()
	}
}

func Test4(t *testing.T) {
	s := qt5.NewQString("vioug")
	if s == nil {
		t.Fail()
	}

	s2 := s.Append("123")
	if s2 == nil {
		t.Fail()
	}

	switch s2.(type) {
	case *qt5.QString:
	default:
		t.Fail()
	}

	len := s.Length().(int32)
	if len != 8 {
		t.Fail()
	}

}

func Test5(t *testing.T) {
	s := qt5.NewQString("vio%1ug")
	if s == nil {
		t.Fail()
	}

	qc := qt5.NewQChar()
	s2 := s.Arg(32132, 0, 10, qc)
	if s2 == nil {
		t.Fail()
	}

	len := s2.(*qt5.QString).Length()
	if len.(int32) != 10 {
		t.Fail()
	}
}
