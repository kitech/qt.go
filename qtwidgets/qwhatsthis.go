package qtwidgets

// /usr/include/qt/QtWidgets/qwhatsthis.h
// #include <qwhatsthis.h>
// #include <QtWidgets>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWhatsThis) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQWhatsThisFromPointer(cthis unsafe.Pointer) *QWhatsThis {
	return &QWhatsThis{&qtrt.CObject{cthis}}
}
func (*QWhatsThis) NewFromPointer(cthis unsafe.Pointer) *QWhatsThis {
	return NewQWhatsThisFromPointer(cthis)
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
func (this *QWhatsThis) InWhatsThisMode() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis15inWhatsThisModeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QWhatsThis_InWhatsThisMode() bool {
	var nilthis *QWhatsThis
	rv := nilthis.InWhatsThisMode()
	return rv
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
func (this *QWhatsThis) ShowText(pos *qtcore.QPoint, text *qtcore.QString, w *QWidget /*777 QWidget **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget", ffiqt.FFI_TYPE_POINTER, pos, text, w)
	gopp.ErrPrint(err, rv)
}
func QWhatsThis_ShowText(pos *qtcore.QPoint, text *qtcore.QString, w *QWidget /*777 QWidget **/) {
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
func (this *QWhatsThis) CreateAction(parent *qtcore.QObject /*777 QObject **/) *QAction /*777 QAction **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", ffiqt.FFI_TYPE_POINTER, parent)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QWhatsThis_CreateAction(parent *qtcore.QObject /*777 QObject **/) *QAction /*777 QAction **/ {
	var nilthis *QWhatsThis
	rv := nilthis.CreateAction(parent)
	return rv
}

//  body block end
