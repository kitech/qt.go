package qtgui

// /usr/include/qt/QtGui/qbrush.h
// #include <qbrush.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QBrushData struct {
	*qtrt.CObject
}

func (this *QBrushData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QBrushData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQBrushDataFromPointer(cthis unsafe.Pointer) *QBrushData {
	return &QBrushData{&qtrt.CObject{cthis}}
}
func (*QBrushData) NewFromPointer(cthis unsafe.Pointer) *QBrushData {
	return NewQBrushDataFromPointer(cthis)
}

func DeleteQBrushData(this *QBrushData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBrushDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
