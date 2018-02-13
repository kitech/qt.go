package qtcore

// /usr/include/qt/QtCore/qarraydata.h
// #include <qarraydata.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QContainerImplHelper struct {
	*qtrt.CObject
}
type QContainerImplHelper_ITF interface {
	QContainerImplHelper_PTR() *QContainerImplHelper
}

func (ptr *QContainerImplHelper) QContainerImplHelper_PTR() *QContainerImplHelper { return ptr }

func (this *QContainerImplHelper) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QContainerImplHelper) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQContainerImplHelperFromPointer(cthis unsafe.Pointer) *QContainerImplHelper {
	return &QContainerImplHelper{&qtrt.CObject{cthis}}
}
func (*QContainerImplHelper) NewFromPointer(cthis unsafe.Pointer) *QContainerImplHelper {
	return NewQContainerImplHelperFromPointer(cthis)
}

// /usr/include/qt/QtCore/qarraydata.h:373
// index:0
// Public static Visibility=Default Availability=Available
// [4] QtPrivate::QContainerImplHelper::CutResult mid(int, int *, int *)
func (this *QContainerImplHelper) Mid(originalLength int, position unsafe.Pointer /*666*/, length unsafe.Pointer /*666*/) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QtPrivate20QContainerImplHelper3midEiPiS1_", qtrt.FFI_TYPE_POINTER, originalLength, &position, &length)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QContainerImplHelper_Mid(originalLength int, position unsafe.Pointer /*666*/, length unsafe.Pointer /*666*/) int {
	var nilthis *QContainerImplHelper
	rv := nilthis.Mid(originalLength, position, length)
	return rv
}

func DeleteQContainerImplHelper(this *QContainerImplHelper) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QContainerImplHelperD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QContainerImplHelper__CutResult = int

const QContainerImplHelper__Null QContainerImplHelper__CutResult = 0
const QContainerImplHelper__Empty QContainerImplHelper__CutResult = 1
const QContainerImplHelper__Full QContainerImplHelper__CutResult = 2
const QContainerImplHelper__Subset QContainerImplHelper__CutResult = 3

//  body block end

//  keep block begin

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
}

//  keep block end
