package qtcore

// /usr/include/qt/QtCore/qshareddata.h
// #include <qshareddata.h>
// #include <QtCore>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
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
type QSharedData struct {
	*qtrt.CObject
}

func (this *QSharedData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSharedData) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQSharedDataFromPointer(cthis unsafe.Pointer) *QSharedData {
	return &QSharedData{&qtrt.CObject{cthis}}
}
func (*QSharedData) NewFromPointer(cthis unsafe.Pointer) *QSharedData {
	return NewQSharedDataFromPointer(cthis)
}

// /usr/include/qt/QtCore/qshareddata.h:60
// index:0
// Public inline
// void QSharedData()
func NewQSharedData() *QSharedData {
	cthis := qtrt.Calloc(1, 256) // 4
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QSharedDataC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSharedDataFromPointer(cthis)
	return gothis
}

//  body block end
