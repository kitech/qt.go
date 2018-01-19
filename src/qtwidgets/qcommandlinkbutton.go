//  header block begin
// /usr/include/qt/QtWidgets/qcommandlinkbutton.h
// #include <qcommandlinkbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QCommandLinkButton struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCommandLinkButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:61
// index:0
// void QCommandLinkButton(class QWidget *)
func NewQCommandLinkButton(parent unsafe.Pointer) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QCommandLinkButton{cthis}
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:62
// index:1
// void QCommandLinkButton(const class QString &, class QWidget *)
func NewQCommandLinkButton_1(text unsafe.Pointer, parent unsafe.Pointer) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	return &QCommandLinkButton{cthis}
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:63
// index:2
// void QCommandLinkButton(const class QString &, const class QString &, class QWidget *)
func NewQCommandLinkButton_2(text unsafe.Pointer, description unsafe.Pointer, parent unsafe.Pointer) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringS2_P7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, description, parent)
	gopp.ErrPrint(err, rv)
	return &QCommandLinkButton{cthis}
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:64
// index:0
// virtual
// void ~QCommandLinkButton()
func DeleteQCommandLinkButton(*QCommandLinkButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:66
// index:0
// QString description()
func (this *QCommandLinkButton) Description() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton11descriptionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:67
// index:0
// void setDescription(const class QString &)
func (this *QCommandLinkButton) SetDescription(description unsafe.Pointer) {
	// 0: (, const QString & description), (description)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton14setDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, description)
	gopp.ErrPrint(err, rv)
}

//  body block end
