package qtcore

// /usr/include/qt/QtCore/qstringliteral.h
// #include <qstringliteral.h>
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

type QStringDataPtr struct {
	*qtrt.CObject
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QStringDataPtrD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
