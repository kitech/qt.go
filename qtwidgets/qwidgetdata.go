package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QWidgetData struct {
	*qtrt.CObject
}
type QWidgetData_ITF interface {
	QWidgetData_PTR() *QWidgetData
}

func (ptr *QWidgetData) QWidgetData_PTR() *QWidgetData { return ptr }

func (this *QWidgetData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWidgetData) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWidgetDataFromPointer(cthis unsafe.Pointer) *QWidgetData {
	return &QWidgetData{&qtrt.CObject{cthis}}
}
func (*QWidgetData) NewFromPointer(cthis unsafe.Pointer) *QWidgetData {
	return NewQWidgetDataFromPointer(cthis)
}

func DeleteQWidgetData(this *QWidgetData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QWidgetDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10981() {
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
