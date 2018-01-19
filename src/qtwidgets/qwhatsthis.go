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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:58
// index:0
// static
// void enterWhatsThisMode()
func (this *QWhatsThis) EnterWhatsThisMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis18enterWhatsThisModeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_EnterWhatsThisMode() {
	// 0: (), ()
	var nilthis *QWhatsThis
	nilthis.EnterWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:59
// index:0
// static
// bool inWhatsThisMode()
func (this *QWhatsThis) InWhatsThisMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis15inWhatsThisModeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_InWhatsThisMode() {
	// 0: (), ()
	var nilthis *QWhatsThis
	nilthis.InWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:60
// index:0
// static
// void leaveWhatsThisMode()
func (this *QWhatsThis) LeaveWhatsThisMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis18leaveWhatsThisModeEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_LeaveWhatsThisMode() {
	// 0: (), ()
	var nilthis *QWhatsThis
	nilthis.LeaveWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:62
// index:0
// static
// void showText(const class QPoint &, const class QString &, class QWidget *)
func (this *QWhatsThis) ShowText(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer) {
	// 0: (const QPoint & pos, const QString & text, QWidget * w), (pos, text, w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_ShowText(pos unsafe.Pointer, text unsafe.Pointer, w unsafe.Pointer) {
	// 0: (const QPoint & pos, const QString & text, QWidget * w), (pos, text, w)
	var nilthis *QWhatsThis
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:63
// index:0
// static
// void hideText()
func (this *QWhatsThis) HideText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis8hideTextEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_HideText() {
	// 0: (), ()
	var nilthis *QWhatsThis
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:65
// index:0
// static
// QAction * createAction(class QObject *)
func (this *QWhatsThis) CreateAction(parent unsafe.Pointer) {
	// 0: (QObject * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_CreateAction(parent unsafe.Pointer) {
	// 0: (QObject * parent), (parent)
	var nilthis *QWhatsThis
	nilthis.CreateAction(parent)
}

//  body block end
