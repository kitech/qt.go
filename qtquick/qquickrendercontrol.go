package qtquick

// /usr/include/qt/QtQuick/qquickrendercontrol.h
// #include <qquickrendercontrol.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QQuickRenderControl struct {
	*qtcore.QObject
}
type QQuickRenderControl_ITF interface {
	qtcore.QObject_ITF
	QQuickRenderControl_PTR() *QQuickRenderControl
}

func (ptr *QQuickRenderControl) QQuickRenderControl_PTR() *QQuickRenderControl { return ptr }

func (this *QQuickRenderControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickRenderControl) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQuickRenderControlFromPointer(cthis unsafe.Pointer) *QQuickRenderControl {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQuickRenderControl{bcthis0}
}
func (*QQuickRenderControl) NewFromPointer(cthis unsafe.Pointer) *QQuickRenderControl {
	return NewQQuickRenderControlFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickRenderControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QQuickRenderControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickRenderControl(QObject *)

/*
Constructs a QQuickRenderControl object, with parent object parent.
*/
func (*QQuickRenderControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QQuickRenderControl {
	return NewQQuickRenderControl(parent)
}
func NewQQuickRenderControl(parent qtcore.QObject_ITF /*777 QObject **/) *QQuickRenderControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickRenderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickRenderControl")
	return gothis
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickRenderControl(QObject *)

/*
Constructs a QQuickRenderControl object, with parent object parent.
*/
func (*QQuickRenderControl) NewForInheritp() *QQuickRenderControl {
	return NewQQuickRenderControlp()
}
func NewQQuickRenderControlp() *QQuickRenderControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickRenderControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickRenderControl")
	return gothis
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickRenderControl()

/*

 */
func DeleteQQuickRenderControl(this *QQuickRenderControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void prepareThread(QThread *)

/*
Prepares rendering the Qt Quick scene outside the gui thread.

targetThread specifies the thread on which synchronization and rendering will happen. There is no need to call this function in a single threaded scenario.
*/
func (this *QQuickRenderControl) PrepareThread(targetThread qtcore.QThread_ITF /*777 QThread **/) {
	var convArg0 unsafe.Pointer
	if targetThread != nil && targetThread.QThread_PTR() != nil {
		convArg0 = targetThread.QThread_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl13prepareThreadEP7QThread", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Stop rendering and release resources. Requires a current context.

This is the equivalent of the cleanup operations that happen with a real QQuickWindow when the window becomes hidden.

This function is called from the destructor. Therefore there will typically be no need to call it directly. Pay attention however to the fact that this requires the context, that was passed to initialize(), to be the current one at the time of destroying the QQuickRenderControl instance.

Once invalidate() has been called, it is possible to reuse the QQuickRenderControl instance by calling initialize() again.

Note: This function does not take QQuickWindow::persistentSceneGraph() or QQuickWindow::persistentOpenGLContext() into account. This means that context-specific resources are always released.
*/
func (this *QQuickRenderControl) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void polishItems()

/*
This function should be called as late as possible before sync(). In a threaded scenario, rendering can happen in parallel with this function.
*/
func (this *QQuickRenderControl) PolishItems() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl11polishItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render()

/*
Renders the scenegraph using the current context.
*/
func (this *QQuickRenderControl) Render() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl6renderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:67
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sync()

/*
This function is used to synchronize the QML scene with the rendering scene graph.

If a dedicated render thread is used, the GUI thread should be blocked for the duration of this call.

Returns true if the synchronization changed the scene graph.
*/
func (this *QQuickRenderControl) Sync() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl4syncEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:69
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage grab()

/*
Grabs the contents of the scene and returns it as an image.

Note: Requires the context to be current.
*/
func (this *QQuickRenderControl) Grab() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl4grabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * renderWindowFor(QQuickWindow *, QPoint *)

/*
Returns the real window that win is being rendered to, if any.

If offset in non-null, it is set to the offset of the rendering inside its window.
*/
func (this *QQuickRenderControl) RenderWindowFor(win QQuickWindow_ITF /*777 QQuickWindow **/, offset qtcore.QPoint_ITF /*777 QPoint **/) *qtgui.QWindow /*777 QWindow **/ {
	var convArg0 unsafe.Pointer
	if win != nil && win.QQuickWindow_PTR() != nil {
		convArg0 = win.QQuickWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg1 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl15renderWindowForEP12QQuickWindowP6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QQuickRenderControl_RenderWindowFor(win QQuickWindow_ITF /*777 QQuickWindow **/, offset qtcore.QPoint_ITF /*777 QPoint **/) *qtgui.QWindow /*777 QWindow **/ {
	var nilthis *QQuickRenderControl
	rv := nilthis.RenderWindowFor(win, offset)
	return rv
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:71
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * renderWindowFor(QQuickWindow *, QPoint *)

/*
Returns the real window that win is being rendered to, if any.

If offset in non-null, it is set to the offset of the rendering inside its window.
*/
func (this *QQuickRenderControl) RenderWindowForp(win QQuickWindow_ITF /*777 QQuickWindow **/) *qtgui.QWindow /*777 QWindow **/ {
	var convArg0 unsafe.Pointer
	if win != nil && win.QQuickWindow_PTR() != nil {
		convArg0 = win.QQuickWindow_PTR().GetCthis()
	}
	// arg: 1, QPoint *=Pointer, QPoint=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl15renderWindowForEP12QQuickWindowP6QPoint", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:72
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] QWindow * renderWindow(QPoint *)

/*
Reimplemented in subclasses to return the real window this render control is rendering into.

If offset in non-null, it is set to the offset of the control inside the window.

Note: While not mandatory, reimplementing this function becomes essential for supporting multiple screens with different device pixel ratios and properly positioning popup windows opened from QML. Therefore providing it in subclasses is highly recommended.
*/
func (this *QQuickRenderControl) RenderWindow(offset qtcore.QPoint_ITF /*777 QPoint **/) *qtgui.QWindow /*777 QWindow **/ {
	var convArg0 unsafe.Pointer
	if offset != nil && offset.QPoint_PTR() != nil {
		convArg0 = offset.QPoint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl12renderWindowEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void renderRequested()

/*
This signal is emitted when the scene graph needs to be rendered. It is not necessary to call sync().

Note: Avoid triggering rendering directly when this signal is emitted. Instead, prefer deferring it by using a timer for example. This will lead to better performance.
*/
func (this *QQuickRenderControl) RenderRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl15renderRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickrendercontrol.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneChanged()

/*
This signal is emitted when the scene graph is updated, meaning that polishItems() and sync() needs to be called. If sync() returns true, then render() needs to be called.

Note: Avoid triggering polishing, synchronization and rendering directly when this signal is emitted. Instead, prefer deferring it by using a timer for example. This will lead to better performance.
*/
func (this *QQuickRenderControl) SceneChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QQuickRenderControl12sceneChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

//  body block end

//  keep block begin

func init_unused_11563() {
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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
