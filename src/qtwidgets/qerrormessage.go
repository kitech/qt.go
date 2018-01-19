//  header block begin
// /usr/include/qt/QtWidgets/qerrormessage.h
// #include <qerrormessage.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
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
type QErrorMessage struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qerrormessage.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QErrorMessage) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK13QErrorMessage10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:58
// index:0
// void QErrorMessage(class QWidget *)
func NewQErrorMessage(parent unsafe.Pointer) *QErrorMessage {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessageC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QErrorMessage{cthis}
}

// /usr/include/qt/QtWidgets/qerrormessage.h:59
// index:0
// virtual
// void ~QErrorMessage()
func DeleteQErrorMessage(*QErrorMessage) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessageD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:61
// index:0
// static
// QErrorMessage * qtHandler()
func (this *QErrorMessage) QtHandler() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage9qtHandlerEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QErrorMessage_QtHandler() {
	// 0: (), ()
	var nilthis *QErrorMessage
	nilthis.QtHandler()
}

// /usr/include/qt/QtWidgets/qerrormessage.h:64
// index:0
// void showMessage(const class QString &)
func (this *QErrorMessage) ShowMessage(message unsafe.Pointer) {
	// 0: (, const QString & message), (message)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage11showMessageERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, message)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qerrormessage.h:65
// index:1
// void showMessage(const class QString &, const class QString &)
func (this *QErrorMessage) ShowMessage_1(message unsafe.Pointer, type_ unsafe.Pointer) {
	// 1: (, const QString & message, const QString & type), (message, type_)
	rv, err := ffiqt.InvokeQtFunc6("_ZN13QErrorMessage11showMessageERK7QStringS2_", ffiqt.FFI_TYPE_VOID, this.cthis, message, type_)
	gopp.ErrPrint(err, rv)
}

//  body block end
