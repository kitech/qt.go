//  header block begin
// /usr/include/qt/QtWidgets/qdesktopwidget.h
// #include <qdesktopwidget.h>
// #include <QtWidgets>
package qtwidgets

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
type QDesktopWidget struct {
	*QWidget
}

func (this *QDesktopWidget) GetCthis() unsafe.Pointer {
	return this.QWidget.GetCthis()
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QDesktopWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:59
// index:0
// void QDesktopWidget()
func NewQDesktopWidget() *QDesktopWidget {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidgetC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQDesktopWidgetFromPointer(cthis)
	return gothis
}
func NewQDesktopWidgetFromPointer(cthis unsafe.Pointer) *QDesktopWidget {
	bcthis0 := NewQWidgetFromPointer(cthis)
	return &QDesktopWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:60
// index:0
// virtual
// void ~QDesktopWidget()
func DeleteQDesktopWidget(*QDesktopWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:62
// index:0
// bool isVirtualDesktop()
func (this *QDesktopWidget) IsVirtualDesktop() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget16isVirtualDesktopEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:64
// index:0
// int numScreens()
func (this *QDesktopWidget) NumScreens() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget10numScreensEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:65
// index:0
// int screenCount()
func (this *QDesktopWidget) ScreenCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget11screenCountEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:66
// index:0
// int primaryScreen()
func (this *QDesktopWidget) PrimaryScreen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget13primaryScreenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:68
// index:0
// int screenNumber(const class QWidget *)
func (this *QDesktopWidget) ScreenNumber(widget unsafe.Pointer) {
	// 0: (, widget const QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:69
// index:1
// int screenNumber(const class QPoint &)
func (this *QDesktopWidget) ScreenNumber_1(arg0 unsafe.Pointer) {
	// 1: (, const QPoint & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:71
// index:0
// QWidget * screen(int)
func (this *QDesktopWidget) Screen(screen int) {
	// 0: (, screen int), (&screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget6screenEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:73
// index:0
// const QRect screenGeometry(int)
func (this *QDesktopWidget) ScreenGeometry(screen int) {
	// 0: (, screen int), (&screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:74
// index:1
// const QRect screenGeometry(const class QWidget *)
func (this *QDesktopWidget) ScreenGeometry_1(widget unsafe.Pointer) {
	// 1: (, widget const QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:75
// index:2
// inline
// const QRect screenGeometry(const class QPoint &)
func (this *QDesktopWidget) ScreenGeometry_2(point unsafe.Pointer) {
	// 2: (, point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:78
// index:0
// const QRect availableGeometry(int)
func (this *QDesktopWidget) AvailableGeometry(screen int) {
	// 0: (, screen int), (&screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:79
// index:1
// const QRect availableGeometry(const class QWidget *)
func (this *QDesktopWidget) AvailableGeometry_1(widget unsafe.Pointer) {
	// 1: (, widget const QWidget *), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:80
// index:2
// inline
// const QRect availableGeometry(const class QPoint &)
func (this *QDesktopWidget) AvailableGeometry_2(point unsafe.Pointer) {
	// 2: (, point const QPoint &), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryERK6QPoint", ffiqt.FFI_TYPE_VOID, this.GetCthis(), point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:84
// index:0
// void resized(int)
func (this *QDesktopWidget) Resized(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget7resizedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:85
// index:0
// void workAreaResized(int)
func (this *QDesktopWidget) WorkAreaResized(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget15workAreaResizedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:86
// index:0
// void screenCountChanged(int)
func (this *QDesktopWidget) ScreenCountChanged(arg0 int) {
	// 0: (, int arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget18screenCountChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:87
// index:0
// void primaryScreenChanged()
func (this *QDesktopWidget) PrimaryScreenChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget20primaryScreenChangedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:90
// index:0
// virtual
// void resizeEvent(class QResizeEvent *)
func (this *QDesktopWidget) ResizeEvent(e unsafe.Pointer) {
	// 0: (, e QResizeEvent *), (e)
	rv, err := ffiqt.InvokeQtFunc6("_ZN14QDesktopWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_VOID, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
}

//  body block end
