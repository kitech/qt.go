package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextBlockUserData struct {
	*qtrt.CObject
}
type QTextBlockUserData_ITF interface {
	QTextBlockUserData_PTR() *QTextBlockUserData
}

func (ptr *QTextBlockUserData) QTextBlockUserData_PTR() *QTextBlockUserData { return ptr }

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
	rv, err := qtrt.InvokeQtFunc6("_ZN18QTextBlockUserDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
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
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
