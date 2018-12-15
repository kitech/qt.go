package qtcore

// /usr/include/qt/QtCore/qstring.h
// #include <qstring.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QStringDataPtr struct {
	*qtrt.CObject
}
type QStringDataPtr_ITF interface {
	QStringDataPtr_PTR() *QStringDataPtr
}

func (ptr *QStringDataPtr) QStringDataPtr_PTR() *QStringDataPtr { return ptr }

func (this *QStringDataPtr) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QStringDataPtr) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQStringDataPtrFromPointer(cthis unsafe.Pointer) *QStringDataPtr {
	return &QStringDataPtr{&qtrt.CObject{cthis}}
}
func (*QStringDataPtr) NewFromPointer(cthis unsafe.Pointer) *QStringDataPtr {
	return NewQStringDataPtrFromPointer(cthis)
}

func DeleteQStringDataPtr(this *QStringDataPtr) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QStringDataPtrD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
