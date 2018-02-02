package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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

type QTextBlockUserData struct {
	*qtrt.CObject
}

func (this *QTextBlockUserData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextBlockUserData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextBlockUserDataFromPointer(cthis unsafe.Pointer) *QTextBlockUserData {
	return &QTextBlockUserData{&qtrt.CObject{cthis}}
}
func (*QTextBlockUserData) NewFromPointer(cthis unsafe.Pointer) *QTextBlockUserData {
	return NewQTextBlockUserDataFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:198
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextBlockUserData()
func DeleteQTextBlockUserData(this *QTextBlockUserData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QTextBlockUserDataD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
