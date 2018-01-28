package qtcore

// /usr/include/qt/QtCore/qarraydata.h
// #include <qarraydata.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QContainerImplHelper struct {
	*qtrt.CObject
}

func (this *QContainerImplHelper) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QContainerImplHelper) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQContainerImplHelperFromPointer(cthis unsafe.Pointer) *QContainerImplHelper {
	return &QContainerImplHelper{&qtrt.CObject{cthis}}
}
func (*QContainerImplHelper) NewFromPointer(cthis unsafe.Pointer) *QContainerImplHelper {
	return NewQContainerImplHelperFromPointer(cthis)
}

// /usr/include/qt/QtCore/qarraydata.h:373
// index:0
// Public static
// QtPrivate::QContainerImplHelper::CutResult mid(int, int *, int *)
func (this *QContainerImplHelper) Mid(originalLength int, position unsafe.Pointer /*666*/, length unsafe.Pointer /*666*/) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QtPrivate20QContainerImplHelper3midEiPiS1_", ffiqt.FFI_TYPE_POINTER, originalLength, &position, &length)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QContainerImplHelper_Mid(originalLength int, position unsafe.Pointer /*666*/, length unsafe.Pointer /*666*/) int {
	var nilthis *QContainerImplHelper
	rv := nilthis.Mid(originalLength, position, length)
	return rv
}

type QContainerImplHelper__CutResult = int

const QContainerImplHelper__Null QContainerImplHelper__CutResult = 0
const QContainerImplHelper__Empty QContainerImplHelper__CutResult = 1
const QContainerImplHelper__Full QContainerImplHelper__CutResult = 2
const QContainerImplHelper__Subset QContainerImplHelper__CutResult = 3

//  body block end
