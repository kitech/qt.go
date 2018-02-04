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
// bool event(class QEvent *)
func (this *QPushButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QPushButton) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QPushButton) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QPushButton) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QPushButton) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void initStyleOption(class QStyleOptionButton *)
func (this *QPushButton) InheritInitStyleOption(f func(option *QStyleOptionButton /*777 QStyleOptionButton **/)) {
	qtrt.SetAllInheritCallback(this, "initStyleOption", f)
}

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QPushButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(QWidget *)
func NewQPushButton(parent *QWidget /*777 QWidget **/) *QPushButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:65
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QString &, QWidget *)
func NewQPushButton_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QPushButton {
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:66
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPushButton(const QIcon &, const QString &, QWidget *)
func NewQPushButton_2(icon *qtgui.QIcon, text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QPushButton {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonC2ERK5QIconRK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQPushButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qpushbutton.h:67
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPushButton()
func DeleteQPushButton(this *QPushButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QPushButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QPushButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:72
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoDefault()
func (this *QPushButton) AutoDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton11autoDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoDefault(_Bool)
func (this *QPushButton) SetAutoDefault(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton14setAutoDefaultEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:74
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDefault()
func (this *QPushButton) IsDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton9isDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefault(_Bool)
func (this *QPushButton) SetDefault(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton10setDefaultEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenu(QMenu *)
func (this *QPushButton) SetMenu(menu *QMenu /*777 QMenu **/) {
	var convArg0 = menu.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton7setMenuEP5QMenu", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QMenu * menu()
func (this *QPushButton) Menu() *QMenu /*777 QMenu **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton4menuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMenuFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qpushbutton.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlat(_Bool)
func (this *QPushButton) SetFlat(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton7setFlatEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlat()
func (this *QPushButton) IsFlat() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton6isFlatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMenu()
func (this *QPushButton) ShowMenu() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton8showMenuEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QPushButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qpushbutton.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QPushButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QPushButton) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:94
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QPushButton) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QPushButton) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QPushButton13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qpushbutton.h:96
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void initStyleOption(QStyleOptionButton *)
func (this *QPushButton) InitStyleOption(option *QStyleOptionButton /*777 QStyleOptionButton **/) {
	var convArg0 = option.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QPushButton15initStyleOptionEP18QStyleOptionButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
