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
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*

 */
type QByteArrayDataPtr struct {
	*qtrt.CObject
}
type QByteArrayDataPtr_ITF interface {
	QByteArrayDataPtr_PTR() *QByteArrayDataPtr
}

func (ptr *QByteArrayDataPtr) QByteArrayDataPtr_PTR() *QByteArrayDataPtr { return ptr }

func (this *QByteArrayDataPtr) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QByteArrayDataPtr) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQByteArrayDataPtrFromPointer(cthis unsafe.Pointer) *QByteArrayDataPtr {
	return &QByteArrayDataPtr{&qtrt.CObject{cthis}}
}
func (*QByteArrayDataPtr) NewFromPointer(cthis unsafe.Pointer) *QByteArrayDataPtr {
	return NewQByteArrayDataPtrFromPointer(cthis)
}

func DeleteQByteArrayDataPtr(this *QByteArrayDataPtr) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QByteArrayDataPtrD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10177() {
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
