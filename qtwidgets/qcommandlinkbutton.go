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
import "qt.go/cffiqt"
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QCommandLinkButton) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:61
// index:0
// Public
// void QCommandLinkButton(QWidget *)
func NewQCommandLinkButton(parent *QWidget /*777 QWidget **/) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:62
// index:1
// Public
// void QCommandLinkButton(const QString &, QWidget *)
func NewQCommandLinkButton_1(text *qtcore.QString, parent *QWidget /*777 QWidget **/) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:63
// index:2
// Public
// void QCommandLinkButton(const QString &, const QString &, QWidget *)
func NewQCommandLinkButton_2(text *qtcore.QString, description *qtcore.QString, parent *QWidget /*777 QWidget **/) *QCommandLinkButton {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = text.GetCthis()
	var convArg1 = description.GetCthis()
	var convArg2 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonC2ERK7QStringS2_P7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQCommandLinkButtonFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:64
// index:0
// Public virtual
// void ~QCommandLinkButton()
func DeleteQCommandLinkButton(*QCommandLinkButton) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButtonD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:66
// index:0
// Public
// QString description()
func (this *QCommandLinkButton) Description() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton11descriptionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:67
// index:0
// Public
// void setDescription(const QString &)
func (this *QCommandLinkButton) SetDescription(description *qtcore.QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton14setDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:70
// index:0
// Protected virtual
// QSize sizeHint()
func (this *QCommandLinkButton) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:71
// index:0
// Protected virtual
// int heightForWidth(int)
func (this *QCommandLinkButton) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:72
// index:0
// Protected virtual
// QSize minimumSizeHint()
func (this *QCommandLinkButton) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QCommandLinkButton15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:73
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QCommandLinkButton) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qcommandlinkbutton.h:74
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QCommandLinkButton) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QCommandLinkButton10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
