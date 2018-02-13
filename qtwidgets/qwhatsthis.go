package qtwidgets

// /usr/include/qt/QtWidgets/qwhatsthis.h
// #include <qwhatsthis.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

type QWhatsThis struct {
	*qtrt.CObject
}
type QWhatsThis_ITF interface {
	QWhatsThis_PTR() *QWhatsThis
}

func (ptr *QWhatsThis) QWhatsThis_PTR() *QWhatsThis { return ptr }

func (this *QWhatsThis) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWhatsThis) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWhatsThisFromPointer(cthis unsafe.Pointer) *QWhatsThis {
	return &QWhatsThis{&qtrt.CObject{cthis}}
}
func (*QWhatsThis) NewFromPointer(cthis unsafe.Pointer) *QWhatsThis {
	return NewQWhatsThisFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:58
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void enterWhatsThisMode()
func (this *QWhatsThis) EnterWhatsThisMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis18enterWhatsThisModeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_EnterWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.EnterWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool inWhatsThisMode()
func (this *QWhatsThis) InWhatsThisMode() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis15inWhatsThisModeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QWhatsThis_InWhatsThisMode() bool {
	var nilthis *QWhatsThis
	rv := nilthis.InWhatsThisMode()
	return rv
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:60
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void leaveWhatsThisMode()
func (this *QWhatsThis) LeaveWhatsThisMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis18leaveWhatsThisModeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_LeaveWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.LeaveWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:62
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void showText(const QPoint &, const QString &, QWidget *)
func (this *QWhatsThis) ShowText(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/) {
	var convArg0 = pos.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = w.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_ShowText(pos *qtcore.QPoint, text string, w *QWidget /*777 QWidget **/) {
	var nilthis *QWhatsThis
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void hideText()
func (this *QWhatsThis) HideText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis8hideTextEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_HideText() {
	var nilthis *QWhatsThis
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:65
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAction * createAction(QObject *)
func (this *QWhatsThis) CreateAction(parent *qtcore.QObject /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWhatsThis_CreateAction(parent *qtcore.QObject /*777 QObject **/) *QAction /*777 QAction **/ {
	var nilthis *QWhatsThis
	rv := nilthis.CreateAction(parent)
	return rv
}

func DeleteQWhatsThis(this *QWhatsThis) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThisD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
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
