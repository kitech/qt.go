//  header block begin
// /usr/include/qt/QtWidgets/qcommandlinkbutton.h
// #include <qcommandlinkbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
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
	*QPushButton
}

func (this *QCommandLinkButton) GetCthis() unsafe.Pointer {
	return this.QPushButton.GetCthis()
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:55
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QCommandLinkButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:61
// index:0
// void QCommandLinkButton(class QWidget *)
func NewQCommandLinkButton(parent unsafe.Pointer) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(cthis)
	return gothis
}
func NewQCommandLinkButtonFromPointer(cthis unsafe.Pointer) *QCommandLinkButton {
	bcthis0 := NewQPushButtonFromPointer(cthis)
	return &QCommandLinkButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:62
// index:1
// void QCommandLinkButton(const class QString &, class QWidget *)
func NewQCommandLinkButton_1(text unsafe.Pointer, parent unsafe.Pointer) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:63
// index:2
// void QCommandLinkButton(const class QString &, const class QString &, class QWidget *)
func NewQCommandLinkButton_2(text unsafe.Pointer, description unsafe.Pointer, parent unsafe.Pointer) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringS2_P7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, description, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(cthis)
	return gothis
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton11descriptionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:67
// index:0
// void setDescription(const class QString &)
func (this *QCommandLinkButton) SetDescription(description unsafe.Pointer) {
	// 0: (, description const QString &), (description)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton14setDescriptionERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), description)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:70
// index:0
// virtual
// QSize sizeHint()
func (this *QCommandLinkButton) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:71
// index:0
// virtual
// int heightForWidth(int)
func (this *QCommandLinkButton) HeightForWidth(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton14heightForWidthEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:72
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QCommandLinkButton) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:73
// index:0
// virtual
// bool event(class QEvent *)
func (this *QCommandLinkButton) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:74
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QCommandLinkButton) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
