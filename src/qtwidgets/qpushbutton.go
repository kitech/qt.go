//  header block begin
// /usr/include/qt/QtWidgets/qpushbutton.h
// #include <qpushbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 97
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
type QPushButton struct {
	*QAbstractButton
}

func (this *QPushButton) GetCthis() unsafe.Pointer {
	return this.QAbstractButton.GetCthis()
}

// /usr/include/qt/QtWidgets/qpushbutton.h:57
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPushButton) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// void QPushButton(class QWidget *)
func NewQPushButton(parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}
func NewQPushButtonFromPointer(cthis unsafe.Pointer) *QPushButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QPushButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// void QPushButton(const class QString &, class QWidget *)
func NewQPushButton_1(text unsafe.Pointer, parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:66
// index:2
// void QPushButton(const class QIcon &, const class QString &, class QWidget *)
func NewQPushButton_2(icon unsafe.Pointer, text unsafe.Pointer, parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, icon, text, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:97
// index:3
// void QPushButton(class QPushButtonPrivate &, class QWidget *)
func NewQPushButton_3(dd unsafe.Pointer, parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ER18QPushButtonPrivateP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:67
// index:0
// virtual
// void ~QPushButton()
func DeleteQPushButton(*QPushButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:69
// index:0
// virtual
// QSize sizeHint()
func (this *QPushButton) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:70
// index:0
// virtual
// QSize minimumSizeHint()
func (this *QPushButton) MinimumSizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton15minimumSizeHintEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:72
// index:0
// bool autoDefault()
func (this *QPushButton) AutoDefault() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton11autoDefaultEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:73
// index:0
// void setAutoDefault(_Bool)
func (this *QPushButton) SetAutoDefault(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton14setAutoDefaultEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:74
// index:0
// bool isDefault()
func (this *QPushButton) IsDefault() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton9isDefaultEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:75
// index:0
// void setDefault(_Bool)
func (this *QPushButton) SetDefault(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton10setDefaultEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:78
// index:0
// void setMenu(class QMenu *)
func (this *QPushButton) SetMenu(menu unsafe.Pointer) {
	// 0: (, menu QMenu *), (menu)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton7setMenuEP5QMenu", ffiqt.FFI_TYPE_VOID, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:79
// index:0
// QMenu * menu()
func (this *QPushButton) Menu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton4menuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:82
// index:0
// void setFlat(_Bool)
func (this *QPushButton) SetFlat(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton7setFlatEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:83
// index:0
// bool isFlat()
func (this *QPushButton) IsFlat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton6isFlatEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:87
// index:0
// void showMenu()
func (this *QPushButton) ShowMenu() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton8showMenuEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:91
// index:0
// virtual
// bool event(class QEvent *)
func (this *QPushButton) Event(e unsafe.Pointer) {
	// 0: (, e QEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton5eventEP6QEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:92
// index:0
// virtual
// void paintEvent(class QPaintEvent *)
func (this *QPushButton) PaintEvent(arg0 unsafe.Pointer) {
	// 0: (, QPaintEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:93
// index:0
// virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QPushButton) KeyPressEvent(arg0 unsafe.Pointer) {
	// 0: (, QKeyEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:94
// index:0
// virtual
// void focusInEvent(class QFocusEvent *)
func (this *QPushButton) FocusInEvent(arg0 unsafe.Pointer) {
	// 0: (, QFocusEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:95
// index:0
// virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QPushButton) FocusOutEvent(arg0 unsafe.Pointer) {
	// 0: (, QFocusEvent * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:96
// index:0
// void initStyleOption(class QStyleOptionButton *)
func (this *QPushButton) InitStyleOption(option unsafe.Pointer) {
	// 0: (, option QStyleOptionButton *), (option)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_VOID, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
