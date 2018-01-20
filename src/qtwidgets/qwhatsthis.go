//  header block begin
// /usr/include/qt/QtWidgets/qwhatsthis.h
// #include <qwhatsthis.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QWhatsThis struct {
	*qtrt.CObject
}

func (this *QWhatsThis) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQWhatsThisFromPointer(cthis unsafe.Pointer) *QWhatsThis {
	return &QWhatsThis{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:58
// index:0
// Public static
// void enterWhatsThisMode()
func (this *QWhatsThis) EnterWhatsThisMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis18enterWhatsThisModeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_EnterWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.EnterWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:59
// index:0
// Public static
// bool inWhatsThisMode()
func (this *QWhatsThis) InWhatsThisMode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis15inWhatsThisModeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QWhatsThis_InWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.InWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:60
// index:0
// Public static
// void leaveWhatsThisMode()
func (this *QWhatsThis) LeaveWhatsThisMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis18leaveWhatsThisModeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_LeaveWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.LeaveWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:62
// index:0
// Public static
// void showText(const class QPoint &, const class QString &, class QWidget *)
func (this *QWhatsThis) ShowText(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, pos, text, w)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_ShowText(pos *qtcore.QPoint, text *qtcore.QString, w unsafe.Pointer) {
	var nilthis *QWhatsThis
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:63
// index:0
// Public static
// void hideText()
func (this *QWhatsThis) HideText() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis8hideTextEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_HideText() {
	var nilthis *QWhatsThis
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:65
// index:0
// Public static
// QAction * createAction(class QObject *)
func (this *QWhatsThis) CreateAction(parent unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", ffiqt.FFI_TYPE_POINTER, parent)
	gopp.ErrPrint(err, rv)
	return rv
}
func QWhatsThis_CreateAction(parent unsafe.Pointer) {
	var nilthis *QWhatsThis
	nilthis.CreateAction(parent)
}

//  body block end
