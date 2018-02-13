package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 130
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QByteRef struct {
	*qtrt.CObject
}
type QByteRef_ITF interface {
	QByteRef_PTR() *QByteRef
}

func (ptr *QByteRef) QByteRef_PTR() *QByteRef { return ptr }

func (this *QByteRef) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteRef) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQByteRefFromPointer(cthis unsafe.Pointer) *QByteRef {
	return &QByteRef{&qtrt.CObject{cthis}}
}
func (*QByteRef) NewFromPointer(cthis unsafe.Pointer) *QByteRef {
	return NewQByteRefFromPointer(cthis)
}

func DeleteQByteRef(this *QByteRef) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QByteRefD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

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
