//  header block begin
// /usr/include/qt/QtWidgets/qpushbutton.h
// #include <qpushbutton.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 94
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
func NewQPushButtonFromPointer(cthis unsafe.Pointer) *QPushButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QPushButton{bcthis0}
}

// /usr/include/qt/QtWidgets/qpushbutton.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPushButton) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public
// void QPushButton(class QWidget *)
func NewQPushButton(parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public
// void QPushButton(const class QString &, class QWidget *)
func NewQPushButton_1(text *qtcore.QString, parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:66
// index:2
// Public
// void QPushButton(const class QIcon &, const class QString &, class QWidget *)
func NewQPushButton_2(icon *qtgui.QIcon, text *qtcore.QString, parent unsafe.Pointer) *QPushButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:67
// index:0
// Public virtual
// void ~QPushButton()
func DeleteQPushButton(*QPushButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:69
// index:0
// Public virtual
// QSize sizeHint()
func (this *QPushButton) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:70
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QPushButton) MinimumSizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:72
// index:0
// Public
// bool autoDefault()
func (this *QPushButton) AutoDefault() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton11autoDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:73
// index:0
// Public
// void setAutoDefault(_Bool)
func (this *QPushButton) SetAutoDefault(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton14setAutoDefaultEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:74
// index:0
// Public
// bool isDefault()
func (this *QPushButton) IsDefault() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton9isDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:75
// index:0
// Public
// void setDefault(_Bool)
func (this *QPushButton) SetDefault(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton10setDefaultEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:78
// index:0
// Public
// void setMenu(class QMenu *)
func (this *QPushButton) SetMenu(menu unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton7setMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), menu)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:79
// index:0
// Public
// QMenu * menu()
func (this *QPushButton) Menu() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton4menuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:82
// index:0
// Public
// void setFlat(_Bool)
func (this *QPushButton) SetFlat(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton7setFlatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:83
// index:0
// Public
// bool isFlat()
func (this *QPushButton) IsFlat() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton6isFlatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:87
// index:0
// Public
// void showMenu()
func (this *QPushButton) ShowMenu() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton8showMenuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:91
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QPushButton) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qpushbutton.h:92
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QPushButton) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:93
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QPushButton) KeyPressEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:94
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QPushButton) FocusInEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:95
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QPushButton) FocusOutEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:96
// index:0
// Protected
// void initStyleOption(class QStyleOptionButton *)
func (this *QPushButton) InitStyleOption(option unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
}

//  body block end
