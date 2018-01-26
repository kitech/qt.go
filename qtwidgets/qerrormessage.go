package qtwidgets

// /usr/include/qt/QtWidgets/qerrormessage.h
// #include <qerrormessage.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 2
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
type QErrorMessage struct {
	*QDialog
}

func (this *QErrorMessage) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QDialog.GetCthis()
	}
}
func (this *QErrorMessage) SetCthis(cthis unsafe.Pointer) {
	this.QDialog = NewQDialogFromPointer(cthis)
}
func NewQErrorMessageFromPointer(cthis unsafe.Pointer) *QErrorMessage {
	bcthis0 := NewQDialogFromPointer(cthis)
	return &QErrorMessage{bcthis0}
}
func (*QErrorMessage) NewFromPointer(cthis unsafe.Pointer) *QErrorMessage {
	return NewQErrorMessageFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:55
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QErrorMessage) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QErrorMessage10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qerrormessage.h:58
// index:0
// Public
// void QErrorMessage(class QWidget *)
func NewQErrorMessage(parent *QWidget /*777 QWidget **/) *QErrorMessage {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessageC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQErrorMessageFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qerrormessage.h:59
// index:0
// Public virtual
// void ~QErrorMessage()
func DeleteQErrorMessage(*QErrorMessage) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessageD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:61
// index:0
// Public static
// QErrorMessage * qtHandler()
func (this *QErrorMessage) QtHandler() *QErrorMessage /*777 QErrorMessage **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage9qtHandlerEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQErrorMessageFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QErrorMessage_QtHandler() *QErrorMessage /*777 QErrorMessage **/ {
	var nilthis *QErrorMessage
	rv := nilthis.QtHandler()
	return rv
}

// /usr/include/qt/QtWidgets/qerrormessage.h:64
// index:0
// Public
// void showMessage(const class QString &)
func (this *QErrorMessage) ShowMessage(message *qtcore.QString) {
	var convArg0 = message.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage11showMessageERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:65
// index:1
// Public
// void showMessage(const class QString &, const class QString &)
func (this *QErrorMessage) ShowMessage_1(message *qtcore.QString, type_ *qtcore.QString) {
	var convArg0 = message.GetCthis()
	var convArg1 = type_.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage11showMessageERK7QStringS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:68
// index:0
// Protected virtual
// void done(int)
func (this *QErrorMessage) Done(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage4doneEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:69
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QErrorMessage) ChangeEvent(e *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
