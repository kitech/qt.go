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
// extern C begin: 6
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

/*

 */
type QTextFrameLayoutData struct {
	*qtrt.CObject
}
type QTextFrameLayoutData_ITF interface {
	QTextFrameLayoutData_PTR() *QTextFrameLayoutData
}

func (ptr *QTextFrameLayoutData) QTextFrameLayoutData_PTR() *QTextFrameLayoutData { return ptr }

func (this *QTextFrameLayoutData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextFrameLayoutData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextFrameLayoutDataFromPointer(cthis unsafe.Pointer) *QTextFrameLayoutData {
	return &QTextFrameLayoutData{&qtrt.CObject{cthis}}
}
func (*QTextFrameLayoutData) NewFromPointer(cthis unsafe.Pointer) *QTextFrameLayoutData {
	return NewQTextFrameLayoutDataFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:114
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextFrameLayoutData()

/*

 */
func DeleteQTextFrameLayoutData(this *QTextFrameLayoutData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextFrameLayoutDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10947() {
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
