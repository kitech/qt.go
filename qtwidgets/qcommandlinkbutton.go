package qtwidgets

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h
// #include <qcommandlinkbutton.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
// QSize sizeHint()
func (this *QCommandLinkButton) InheritSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

// int heightForWidth(int)
func (this *QCommandLinkButton) InheritHeightForWidth(f func(arg0 int) int) {
	qtrt.SetAllInheritCallback(this, "heightForWidth", f)
}

// QSize minimumSizeHint()
func (this *QCommandLinkButton) InheritMinimumSizeHint(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "minimumSizeHint", f)
}

// bool event(class QEvent *)
func (this *QCommandLinkButton) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QCommandLinkButton) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

type QCommandLinkButton struct {
	*QPushButton
}

func (this *QCommandLinkButton) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPushButton.GetCthis()
	}
}
func (this *QCommandLinkButton) SetCthis(cthis unsafe.Pointer) {
	this.QPushButton = NewQPushButtonFromPointer(cthis)
}
func NewQCommandLinkButtonFromPointer(cthis unsafe.Pointer) *QCommandLinkButton {
	bcthis0 := NewQPushButtonFromPointer(cthis)
	return &QCommandLinkButton{bcthis0}
}
func (*QCommandLinkButton) NewFromPointer(cthis unsafe.Pointer) *QCommandLinkButton {
	return NewQCommandLinkButtonFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QCommandLinkButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLinkButton10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QCommandLinkButton(QWidget *)
func NewQCommandLinkButton(parent *QWidget /*777 QWidget **/) *QCommandLinkButton {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QCommandLinkButton(const QString &, QWidget *)
func NewQCommandLinkButton_1(text string, parent *QWidget /*777 QWidget **/) *QCommandLinkButton {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:63
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QCommandLinkButton(const QString &, const QString &, QWidget *)
func NewQCommandLinkButton_2(text string, description string, parent *QWidget /*777 QWidget **/) *QCommandLinkButton {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = qtcore.NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringS2_P7QWidget", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCommandLinkButton()
func DeleteQCommandLinkButton(this *QCommandLinkButton) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButtonD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description()
func (this *QCommandLinkButton) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLinkButton11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QCommandLinkButton) SetDescription(description string) {
	var tmpArg0 = qtcore.NewQString_5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButton14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:70
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QCommandLinkButton) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLinkButton8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:71
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QCommandLinkButton) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLinkButton14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:72
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QCommandLinkButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QCommandLinkButton15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:73
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QCommandLinkButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButton5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:74
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QCommandLinkButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QCommandLinkButton10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
