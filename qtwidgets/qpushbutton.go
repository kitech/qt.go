package qtwidgets

// /usr/include/qt/QtWidgets/qpushbutton.h
// #include <qpushbutton.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 96
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
	if this == nil {
		return nil
	} else {
		return this.QAbstractButton.GetCthis()
	}
}
func (this *QPushButton) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractButton = NewQAbstractButtonFromPointer(cthis)
}
func NewQPushButtonFromPointer(cthis unsafe.Pointer) *QPushButton {
	bcthis0 := NewQAbstractButtonFromPointer(cthis)
	return &QPushButton{bcthis0}
}
func (*QPushButton) NewFromPointer(cthis unsafe.Pointer) *QPushButton {
	return NewQPushButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:57
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPushButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public
// void QPushButton(class QWidget *)
func NewQPushButton(parent *QWidget /*777 QWidget **/) *QPushButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public
// void QPushButton(const class QString &, class QWidget *)
func NewQPushButton_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QPushButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:66
// index:2
// Public
// void QPushButton(const class QIcon &, const class QString &, class QWidget *)
func NewQPushButton_2(icon *qtgui.QIcon, text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QPushButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
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
func (this *QPushButton) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:70
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QPushButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:72
// index:0
// Public
// bool autoDefault()
func (this *QPushButton) AutoDefault() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton11autoDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:73
// index:0
// Public
// void setAutoDefault(_Bool)
func (this *QPushButton) SetAutoDefault(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton14setAutoDefaultEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:74
// index:0
// Public
// bool isDefault()
func (this *QPushButton) IsDefault() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton9isDefaultEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:75
// index:0
// Public
// void setDefault(_Bool)
func (this *QPushButton) SetDefault(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton10setDefaultEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:78
// index:0
// Public
// void setMenu(class QMenu *)
func (this *QPushButton) SetMenu(menu *QMenu /*777 QMenu **/) {
	var convArg0 = menu.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton7setMenuEP5QMenu", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:79
// index:0
// Public
// QMenu * menu()
func (this *QPushButton) Menu() *QMenu /*777 QMenu **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton4menuEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:82
// index:0
// Public
// void setFlat(_Bool)
func (this *QPushButton) SetFlat(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton7setFlatEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:83
// index:0
// Public
// bool isFlat()
func (this *QPushButton) IsFlat() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton6isFlatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
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
func (this *QPushButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:92
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QPushButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:93
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QPushButton) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:94
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QPushButton) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:95
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QPushButton) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPushButton13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:96
// index:0
// Protected
// void initStyleOption(class QStyleOptionButton *)
func (this *QPushButton) InitStyleOption(option *QStyleOptionButton /*777 QStyleOptionButton **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
