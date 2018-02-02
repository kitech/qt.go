package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 130
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

type QByteRef struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QByteRefD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
