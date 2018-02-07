package qtmacextras

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h
// #include <qmactoolbar.h>
// #include <Qtmacextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

type QMacToolBar struct {
	*qtcore.QObject
}

func (this *QMacToolBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QMacToolBar) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQMacToolBarFromPointer(cthis unsafe.Pointer) *QMacToolBar {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QMacToolBar{bcthis0}
}
func (*QMacToolBar) NewFromPointer(cthis unsafe.Pointer) *QMacToolBar {
	return NewQMacToolBarFromPointer(cthis)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QMacToolBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMacToolBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMacToolBar(QObject *)
func NewQMacToolBar(parent *qtcore.QObject /*777 QObject **/) *QMacToolBar {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQMacToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMacToolBar(const QString &, QObject *)
func NewQMacToolBar_1(identifier *qtcore.QString, parent *qtcore.QObject /*777 QObject **/) *QMacToolBar {
	var convArg0 = identifier.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBarC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQMacToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QMacToolBar()
func DeleteQMacToolBar(this *QMacToolBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QMacToolBarItem * addItem(const QIcon &, const QString &)
func (this *QMacToolBar) AddItem(icon *qtgui.QIcon, text *qtcore.QString) *QMacToolBarItem /*777 QMacToolBarItem **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar7addItemERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QMacToolBarItem * addAllowedItem(const QIcon &, const QString &)
func (this *QMacToolBar) AddAllowedItem(icon *qtgui.QIcon, text *qtcore.QString) *QMacToolBarItem /*777 QMacToolBarItem **/ {
	var convArg0 = icon.GetCthis()
	var convArg1 = text.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar14addAllowedItemERK5QIconRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QMacToolBarItem * addStandardItem(QMacToolBarItem::StandardItem)
func (this *QMacToolBar) AddStandardItem(standardItem int) *QMacToolBarItem /*777 QMacToolBarItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar15addStandardItemEN15QMacToolBarItem12StandardItemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardItem)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:69
// index:0
// Public Visibility=Default Availability=Available
// [8] QMacToolBarItem * addAllowedStandardItem(QMacToolBarItem::StandardItem)
func (this *QMacToolBar) AddAllowedStandardItem(standardItem int) *QMacToolBarItem /*777 QMacToolBarItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar22addAllowedStandardItemEN15QMacToolBarItem12StandardItemE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), standardItem)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQMacToolBarItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addSeparator()
func (this *QMacToolBar) AddSeparator() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar12addSeparatorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void attachToWindow(QWindow *)
func (this *QMacToolBar) AttachToWindow(window *qtgui.QWindow /*777 QWindow **/) {
	var convArg0 = window.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar14attachToWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detachFromWindow()
func (this *QMacToolBar) DetachFromWindow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QMacToolBar16detachFromWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtMacExtras/../../src/macextras/qmactoolbar.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] NSToolbar * nativeToolbar()
func (this *QMacToolBar) NativeToolbar() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QMacToolBar13nativeToolbarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

//  body block end
