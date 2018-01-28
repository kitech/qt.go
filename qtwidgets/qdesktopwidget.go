package qtwidgets

// /usr/include/qt/QtWidgets/qdesktopwidget.h
// #include <qdesktopwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QDesktopWidget struct {
	*QWidget
}

func (this *QDesktopWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWidget.GetCthis()
	}
}
func (this *QDesktopWidget) SetCthis(cthis unsafe.Pointer) {
	this.QWidget = NewQWidgetFromPointer(cthis)
}
func NewQDesktopWidgetFromPointer(cthis unsafe.Pointer) *QDesktopWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDesktopWidget{bcthis0}
}
func (*QDesktopWidget) NewFromPointer(cthis unsafe.Pointer) *QDesktopWidget {
	return NewQDesktopWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QDesktopWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:59
// index:0
// Public
// void QDesktopWidget()
func NewQDesktopWidget() *QDesktopWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidgetC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDesktopWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:60
// index:0
// Public virtual
// void ~QDesktopWidget()
func DeleteQDesktopWidget(*QDesktopWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:62
// index:0
// Public
// bool isVirtualDesktop()
func (this *QDesktopWidget) IsVirtualDesktop() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget16isVirtualDesktopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:64
// index:0
// Public
// int numScreens()
func (this *QDesktopWidget) NumScreens() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget10numScreensEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:65
// index:0
// Public
// int screenCount()
func (this *QDesktopWidget) ScreenCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget11screenCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:66
// index:0
// Public
// int primaryScreen()
func (this *QDesktopWidget) PrimaryScreen() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget13primaryScreenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:68
// index:0
// Public
// int screenNumber(const QWidget *)
func (this *QDesktopWidget) ScreenNumber(widget *QWidget /*777 const QWidget **/) int {
	var convArg0 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:69
// index:1
// Public
// int screenNumber(const QPoint &)
func (this *QDesktopWidget) ScreenNumber_1(arg0 *qtcore.QPoint) int {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:71
// index:0
// Public
// QWidget * screen(int)
func (this *QDesktopWidget) Screen(screen int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget6screenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:73
// index:0
// Public
// const QRect screenGeometry(int)
func (this *QDesktopWidget) ScreenGeometry(screen int) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:74
// index:1
// Public
// const QRect screenGeometry(const QWidget *)
func (this *QDesktopWidget) ScreenGeometry_1(widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg0 = widget.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEPK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:75
// index:2
// Public inline
// const QRect screenGeometry(const QPoint &)
func (this *QDesktopWidget) ScreenGeometry_2(point *qtcore.QPoint) *qtcore.QRect /*123*/ {
	var convArg0 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:78
// index:0
// Public
// const QRect availableGeometry(int)
func (this *QDesktopWidget) AvailableGeometry(screen int) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEi", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), screen)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:79
// index:1
// Public
// const QRect availableGeometry(const QWidget *)
func (this *QDesktopWidget) AvailableGeometry_1(widget *QWidget /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg0 = widget.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEPK7QWidget", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:80
// index:2
// Public inline
// const QRect availableGeometry(const QPoint &)
func (this *QDesktopWidget) AvailableGeometry_2(point *qtcore.QPoint) *qtcore.QRect /*123*/ {
	var convArg0 = point.GetCthis()
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:84
// index:0
// Public
// void resized(int)
func (this *QDesktopWidget) Resized(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget7resizedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:85
// index:0
// Public
// void workAreaResized(int)
func (this *QDesktopWidget) WorkAreaResized(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget15workAreaResizedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:86
// index:0
// Public
// void screenCountChanged(int)
func (this *QDesktopWidget) ScreenCountChanged(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget18screenCountChangedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:87
// index:0
// Public
// void primaryScreenChanged()
func (this *QDesktopWidget) PrimaryScreenChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget20primaryScreenChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:90
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QDesktopWidget) ResizeEvent(e *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
