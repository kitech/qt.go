package qtwidgets

// /usr/include/qt/QtWidgets/qdesktopwidget.h
// #include <qdesktopwidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// void resizeEvent(class QResizeEvent *)
func (this *QDesktopWidget) InheritResizeEvent(f func(e *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

/*

 */
type QDesktopWidget struct {
	*QWidget
}
type QDesktopWidget_ITF interface {
	QWidget_ITF
	QDesktopWidget_PTR() *QDesktopWidget
}

func (ptr *QDesktopWidget) QDesktopWidget_PTR() *QDesktopWidget { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QDesktopWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QDesktopWidget()

/*

 */
func NewQDesktopWidget() *QDesktopWidget {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidgetC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQDesktopWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QDesktopWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:60
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QDesktopWidget()

/*

 */
func DeleteQDesktopWidget(this *QDesktopWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:62
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVirtualDesktop() const

/*

 */
func (this *QDesktopWidget) IsVirtualDesktop() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget16isVirtualDesktopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] int numScreens() const

/*

 */
func (this *QDesktopWidget) NumScreens() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget10numScreensEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] int screenCount() const

/*

 */
func (this *QDesktopWidget) ScreenCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget11screenCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int primaryScreen() const

/*

 */
func (this *QDesktopWidget) PrimaryScreen() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget13primaryScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int screenNumber(const QWidget *) const

/*
Returns the index of the screen that contains the largest part of widget, or -1 if the widget not on a screen.

See also primaryScreen.
*/
func (this *QDesktopWidget) ScreenNumber(widget QWidget_ITF /*777 const QWidget **/) int {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] int screenNumber(const QWidget *) const

/*
Returns the index of the screen that contains the largest part of widget, or -1 if the widget not on a screen.

See also primaryScreen.
*/
func (this *QDesktopWidget) ScreenNumber__() int {
	// arg: 0, const QWidget *=Pointer, QWidget=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:69
// index:1
// Public Visibility=Default Availability=Available
// [4] int screenNumber(const QPoint &) const

/*
Returns the index of the screen that contains the largest part of widget, or -1 if the widget not on a screen.

See also primaryScreen.
*/
func (this *QDesktopWidget) ScreenNumber_1(arg0 qtcore.QPoint_ITF) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPoint_PTR() != nil {
		convArg0 = arg0.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget12screenNumberERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * screen(int)

/*
Returns a widget that represents the screen with index screen (a value of -1 means the default screen).

If the system uses a virtual desktop, the returned widget will have the geometry of the entire virtual desktop; i.e., bounding every screen.

See also primaryScreen, screenCount, and virtualDesktop.
*/
func (this *QDesktopWidget) Screen(screen int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget6screenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:71
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * screen(int)

/*
Returns a widget that represents the screen with index screen (a value of -1 means the default screen).

If the system uses a virtual desktop, the returned widget will have the geometry of the entire virtual desktop; i.e., bounding every screen.

See also primaryScreen, screenCount, and virtualDesktop.
*/
func (this *QDesktopWidget) Screen__() *QWidget /*777 QWidget **/ {
	// arg: 0, int=Int, =Invalid,
	screen := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget6screenEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:73
// index:0
// Public Visibility=Default Availability=Available
// [16] const QRect screenGeometry(int) const

/*
Returns the geometry of the screen with index screen. The default screen is used if screen is -1.

See also screenNumber().
*/
func (this *QDesktopWidget) ScreenGeometry(screen int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:73
// index:0
// Public Visibility=Default Availability=Available
// [16] const QRect screenGeometry(int) const

/*
Returns the geometry of the screen with index screen. The default screen is used if screen is -1.

See also screenNumber().
*/
func (this *QDesktopWidget) ScreenGeometry__() *qtcore.QRect /*123*/ {
	// arg: 0, int=Int, =Invalid,
	screen := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:74
// index:1
// Public Visibility=Default Availability=Available
// [16] const QRect screenGeometry(const QWidget *) const

/*
Returns the geometry of the screen with index screen. The default screen is used if screen is -1.

See also screenNumber().
*/
func (this *QDesktopWidget) ScreenGeometry_1(widget QWidget_ITF /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:75
// index:2
// Public inline Visibility=Default Availability=Available
// [16] const QRect screenGeometry(const QPoint &) const

/*
Returns the geometry of the screen with index screen. The default screen is used if screen is -1.

See also screenNumber().
*/
func (this *QDesktopWidget) ScreenGeometry_2(point qtcore.QPoint_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget14screenGeometryERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:78
// index:0
// Public Visibility=Default Availability=Available
// [16] const QRect availableGeometry(int) const

/*
Returns the available geometry of the screen with index screen. What is available will be subrect of screenGeometry() based on what the platform decides is available (for example excludes the dock and menu bar on macOS, or the task bar on Windows). The default screen is used if screen is -1.

See also screenNumber(), screenGeometry(), and QScreen::availableGeometry().
*/
func (this *QDesktopWidget) AvailableGeometry(screen int) *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:78
// index:0
// Public Visibility=Default Availability=Available
// [16] const QRect availableGeometry(int) const

/*
Returns the available geometry of the screen with index screen. What is available will be subrect of screenGeometry() based on what the platform decides is available (for example excludes the dock and menu bar on macOS, or the task bar on Windows). The default screen is used if screen is -1.

See also screenNumber(), screenGeometry(), and QScreen::availableGeometry().
*/
func (this *QDesktopWidget) AvailableGeometry__() *qtcore.QRect /*123*/ {
	// arg: 0, int=Int, =Invalid,
	screen := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), screen)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:79
// index:1
// Public Visibility=Default Availability=Available
// [16] const QRect availableGeometry(const QWidget *) const

/*
Returns the available geometry of the screen with index screen. What is available will be subrect of screenGeometry() based on what the platform decides is available (for example excludes the dock and menu bar on macOS, or the task bar on Windows). The default screen is used if screen is -1.

See also screenNumber(), screenGeometry(), and QScreen::availableGeometry().
*/
func (this *QDesktopWidget) AvailableGeometry_1(widget QWidget_ITF /*777 const QWidget **/) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:80
// index:2
// Public inline Visibility=Default Availability=Available
// [16] const QRect availableGeometry(const QPoint &) const

/*
Returns the available geometry of the screen with index screen. What is available will be subrect of screenGeometry() based on what the platform decides is available (for example excludes the dock and menu bar on macOS, or the task bar on Windows). The default screen is used if screen is -1.

See also screenNumber(), screenGeometry(), and QScreen::availableGeometry().
*/
func (this *QDesktopWidget) AvailableGeometry_2(point qtcore.QPoint_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPoint_PTR() != nil {
		convArg0 = point.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QDesktopWidget17availableGeometryERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resized(int)

/*
This signal is emitted when the size of screen changes.
*/
func (this *QDesktopWidget) Resized(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget7resizedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void workAreaResized(int)

/*
This signal is emitted when the work area available on screen changes.
*/
func (this *QDesktopWidget) WorkAreaResized(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget15workAreaResizedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenCountChanged(int)

/*
This signal is emitted when the number of screens changes to newCount.

This function was introduced in  Qt 4.6.

Note: Notifier signal for property screenCount.

See also screenCount.
*/
func (this *QDesktopWidget) ScreenCountChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget18screenCountChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void primaryScreenChanged()

/*
This signal is emitted whenever the primary screen changes.

Note: This doesn't mean the QDesktopWidget::primaryScreen index will necessarily be different, but now it will refer to the new primary screen.

This function was introduced in  Qt 5.6.

Note: Notifier signal for property primaryScreen.

See also primaryScreen and screenGeometry().
*/
func (this *QDesktopWidget) PrimaryScreenChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget20primaryScreenChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qdesktopwidget.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)

/*
Reimplemented from QWidget::resizeEvent().
*/
func (this *QDesktopWidget) ResizeEvent(e qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QResizeEvent_PTR() != nil {
		convArg0 = e.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QDesktopWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
