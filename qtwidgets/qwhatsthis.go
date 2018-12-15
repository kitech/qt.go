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
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
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

// /usr/include/qt/QtWidgets/qwhatsthis.h:59
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void enterWhatsThisMode()

/*
This function switches the user interface into "What's This?" mode. The user interface can be switched back into normal mode by the user (e.g. by them clicking or pressing Esc), or programmatically by calling leaveWhatsThisMode().

When entering "What's This?" mode, a QEvent of type Qt::EnterWhatsThisMode is sent to all toplevel widgets.

See also inWhatsThisMode() and leaveWhatsThisMode().
*/
func (this *QWhatsThis) EnterWhatsThisMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis18enterWhatsThisModeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_EnterWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.EnterWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:60
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool inWhatsThisMode()

/*
Returns true if the user interface is in "What's This?" mode; otherwise returns false.

See also enterWhatsThisMode().
*/
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

// /usr/include/qt/QtWidgets/qwhatsthis.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void leaveWhatsThisMode()

/*
If the user interface is in "What's This?" mode, this function switches back to normal mode; otherwise it does nothing.

When leaving "What's This?" mode, a QEvent of type Qt::LeaveWhatsThisMode is sent to all toplevel widgets.

See also enterWhatsThisMode() and inWhatsThisMode().
*/
func (this *QWhatsThis) LeaveWhatsThisMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis18leaveWhatsThisModeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_LeaveWhatsThisMode() {
	var nilthis *QWhatsThis
	nilthis.LeaveWhatsThisMode()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void showText(const QPoint &, const QString &, QWidget *)

/*
Shows text as a "What's This?" window, at global position pos. The optional widget argument, w, is used to determine the appropriate screen on multi-head systems.

See also hideText().
*/
func (this *QWhatsThis) ShowText(pos qtcore.QPoint_ITF, text string, w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg2 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_ShowText(pos qtcore.QPoint_ITF, text string, w QWidget_ITF /*777 QWidget **/) {
	var nilthis *QWhatsThis
	nilthis.ShowText(pos, text, w)
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void showText(const QPoint &, const QString &, QWidget *)

/*
Shows text as a "What's This?" window, at global position pos. The optional widget argument, w, is used to determine the appropriate screen on multi-head systems.

See also hideText().
*/
func (this *QWhatsThis) ShowTextp(pos qtcore.QPoint_ITF, text string) {
	var convArg0 unsafe.Pointer
	if pos != nil && pos.QPoint_PTR() != nil {
		convArg0 = pos.QPoint_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString5(text)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:64
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void hideText()

/*
If a "What's This?" window is showing, this destroys it.

See also showText().
*/
func (this *QWhatsThis) HideText() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis8hideTextEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QWhatsThis_HideText() {
	var nilthis *QWhatsThis
	nilthis.HideText()
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:66
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAction * createAction(QObject *)

/*
Returns a ready-made QAction, used to invoke "What's This?" context help, with the given parent.

The returned QAction provides a convenient way to let users enter "What's This?" mode.
*/
func (this *QWhatsThis) CreateAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWhatsThis_CreateAction(parent qtcore.QObject_ITF /*777 QObject **/) *QAction /*777 QAction **/ {
	var nilthis *QWhatsThis
	rv := nilthis.CreateAction(parent)
	return rv
}

// /usr/include/qt/QtWidgets/qwhatsthis.h:66
// index:0
// Public static Visibility=Default Availability=Available
// [8] QAction * createAction(QObject *)

/*
Returns a ready-made QAction, used to invoke "What's This?" context help, with the given parent.

The returned QAction provides a convenient way to let users enter "What's This?" mode.
*/
func (this *QWhatsThis) CreateActionp() *QAction /*777 QAction **/ {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QWhatsThis12createActionEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQActionFromPointer(unsafe.Pointer(uintptr(rv))) // 444
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
