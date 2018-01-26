package qtcore

// /usr/include/qt/QtCore/qbytearray.h
// #include <qbytearray.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QByteArrayDataPtr struct {
	*qtrt.CObject
}

func (this *QByteArrayDataPtr) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteArrayDataPtr) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQByteArrayDataPtrFromPointer(cthis unsafe.Pointer) *QByteArrayDataPtr {
	return &QByteArrayDataPtr{&qtrt.CObject{cthis}}
}
func (*QByteArrayDataPtr) NewFromPointer(cthis unsafe.Pointer) *QByteArrayDataPtr {
	return NewQByteArrayDataPtrFromPointer(cthis)
}

//  body block end
