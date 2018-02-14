package qtquick

// /usr/include/qt/QtQuick/qquickwindow.h
// #include <qquickwindow.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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

// void exposeEvent(class QExposeEvent *)
func (this *QQuickWindow) InheritExposeEvent(f func(arg0 *qtgui.QExposeEvent /*777 QExposeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "exposeEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QQuickWindow) InheritResizeEvent(f func(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QQuickWindow) InheritShowEvent(f func(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QQuickWindow) InheritHideEvent(f func(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QQuickWindow) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QQuickWindow) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool event(class QEvent *)
func (this *QQuickWindow) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QQuickWindow) InheritKeyPressEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QQuickWindow) InheritKeyReleaseEvent(f func(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QQuickWindow) InheritMousePressEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QQuickWindow) InheritMouseReleaseEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QQuickWindow) InheritMouseDoubleClickEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QQuickWindow) InheritMouseMoveEvent(f func(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QQuickWindow) InheritWheelEvent(f func(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

type QQuickWindow struct {
	*qtgui.QWindow
}
type QQuickWindow_ITF interface {
	qtgui.QWindow_ITF
	QQuickWindow_PTR() *QQuickWindow
}

func (ptr *QQuickWindow) QQuickWindow_PTR() *QQuickWindow { return ptr }

func (this *QQuickWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QWindow.GetCthis()
	}
}
func (this *QQuickWindow) SetCthis(cthis unsafe.Pointer) {
	this.QWindow = qtgui.NewQWindowFromPointer(cthis)
}
func NewQQuickWindowFromPointer(cthis unsafe.Pointer) *QQuickWindow {
	bcthis0 := qtgui.NewQWindowFromPointer(cthis)
	return &QQuickWindow{bcthis0}
}
func (*QQuickWindow) NewFromPointer(cthis unsafe.Pointer) *QQuickWindow {
	return NewQQuickWindowFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickwindow.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQuickWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QWindow *)
func NewQQuickWindow(parent qtgui.QWindow_ITF /*777 QWindow **/) *QQuickWindow {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWindow_PTR() != nil {
		convArg0 = parent.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowC2EP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWindow")
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:110
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QQuickRenderControl *)
func NewQQuickWindow_1(renderControl QQuickRenderControl_ITF /*777 QQuickRenderControl **/) *QQuickWindow {
	var convArg0 unsafe.Pointer
	if renderControl != nil && renderControl.QQuickRenderControl_PTR() != nil {
		convArg0 = renderControl.QQuickRenderControl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowC2EP19QQuickRenderControl", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickWindow")
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickWindow()
func DeleteQQuickWindow(this *QQuickWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickwindow.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * contentItem()
func (this *QQuickWindow) ContentItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow11contentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * activeFocusItem()
func (this *QQuickWindow) ActiveFocusItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow15activeFocusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * focusObject()
func (this *QQuickWindow) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow11focusObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * mouseGrabberItem()
func (this *QQuickWindow) MouseGrabberItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow16mouseGrabberItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sendEvent(QQuickItem *, QEvent *)
func (this *QQuickWindow) SendEvent(arg0 QQuickItem_ITF /*777 QQuickItem **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow9sendEventEP10QQuickItemP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage grabWindow()
func (this *QQuickWindow) GrabWindow() *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow10grabWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderTarget(uint, const QSize &)
func (this *QQuickWindow) SetRenderTarget(fboId uint, size qtcore.QSize_ITF) {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15setRenderTargetEjRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fboId, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] uint renderTargetId()
func (this *QQuickWindow) RenderTargetId() uint {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow14renderTargetIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtQuick/qquickwindow.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize renderTargetSize()
func (this *QQuickWindow) RenderTargetSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow16renderTargetSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetOpenGLState()
func (this *QQuickWindow) ResetOpenGLState() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow16resetOpenGLStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlIncubationController * incubationController()
func (this *QQuickWindow) IncubationController() *qtqml.QQmlIncubationController /*777 QQmlIncubationController **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow20incubationControllerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtqml.NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleRoot()
func (this *QQuickWindow) AccessibleRoot() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow14accessibleRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromImage(const QImage &)
func (this *QQuickWindow) CreateTextureFromImage(image qtgui.QImage_ITF) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow22createTextureFromImageERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:144
// index:1
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromImage(const QImage &, QQuickWindow::CreateTextureOptions)
func (this *QQuickWindow) CreateTextureFromImage_1(image qtgui.QImage_ITF, options int) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 unsafe.Pointer
	if image != nil && image.QImage_PTR() != nil {
		convArg0 = image.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow22createTextureFromImageERK6QImage6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromId(uint, const QSize &, QQuickWindow::CreateTextureOptions)
func (this *QQuickWindow) CreateTextureFromId(id uint, size qtcore.QSize_ITF, options int) *QSGTexture /*777 QSGTexture **/ {
	var convArg1 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg1 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createTextureFromIdEjRK5QSize6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearBeforeRendering(_Bool)
func (this *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow23setClearBeforeRenderingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clearBeforeRendering()
func (this *QQuickWindow) ClearBeforeRendering() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow20clearBeforeRenderingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QQuickWindow) SetColor(color qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if color != nil && color.QColor_PTR() != nil {
		convArg0 = color.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:151
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color()
func (this *QQuickWindow) Color() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasDefaultAlphaBuffer()
func (this *QQuickWindow) HasDefaultAlphaBuffer() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21hasDefaultAlphaBufferEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QQuickWindow_HasDefaultAlphaBuffer() bool {
	var nilthis *QQuickWindow
	rv := nilthis.HasDefaultAlphaBuffer()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:154
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultAlphaBuffer(_Bool)
func (this *QQuickWindow) SetDefaultAlphaBuffer(useAlpha bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21setDefaultAlphaBufferEb", qtrt.FFI_TYPE_POINTER, useAlpha)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetDefaultAlphaBuffer(useAlpha bool) {
	var nilthis *QQuickWindow
	nilthis.SetDefaultAlphaBuffer(useAlpha)
}

// /usr/include/qt/QtQuick/qquickwindow.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentOpenGLContext(_Bool)
func (this *QQuickWindow) SetPersistentOpenGLContext(persistent bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow26setPersistentOpenGLContextEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), persistent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:157
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentOpenGLContext()
func (this *QQuickWindow) IsPersistentOpenGLContext() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow25isPersistentOpenGLContextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentSceneGraph(_Bool)
func (this *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow23setPersistentSceneGraphEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), persistent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentSceneGraph()
func (this *QQuickWindow) IsPersistentSceneGraph() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow22isPersistentSceneGraphEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSceneGraphInitialized()
func (this *QQuickWindow) IsSceneGraphInitialized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow23isSceneGraphInitializedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scheduleRenderJob(QRunnable *, enum QQuickWindow::RenderStage)
func (this *QQuickWindow) ScheduleRenderJob(job qtcore.QRunnable_ITF /*777 QRunnable **/, schedule int) {
	var convArg0 unsafe.Pointer
	if job != nil && job.QRunnable_PTR() != nil {
		convArg0 = job.QRunnable_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17scheduleRenderJobEP9QRunnableNS_11RenderStageE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, schedule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal effectiveDevicePixelRatio()
func (this *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow25effectiveDevicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickwindow.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRendererInterface * rendererInterface()
func (this *QQuickWindow) RendererInterface() *QSGRendererInterface /*777 QSGRendererInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow17rendererInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRendererInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSceneGraphBackend(QSGRendererInterface::GraphicsApi)
func (this *QQuickWindow) SetSceneGraphBackend(api int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow20setSceneGraphBackendEN20QSGRendererInterface11GraphicsApiE", qtrt.FFI_TYPE_POINTER, api)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetSceneGraphBackend(api int) {
	var nilthis *QQuickWindow
	nilthis.SetSceneGraphBackend(api)
}

// /usr/include/qt/QtQuick/qquickwindow.h:172
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void setSceneGraphBackend(const QString &)
func (this *QQuickWindow) SetSceneGraphBackend_1(backend string) {
	var tmpArg0 = qtcore.NewQString_5(backend)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow20setSceneGraphBackendERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetSceneGraphBackend_1(backend string) {
	var nilthis *QQuickWindow
	nilthis.SetSceneGraphBackend_1(backend)
}

// /usr/include/qt/QtQuick/qquickwindow.h:173
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sceneGraphBackend()
func (this *QQuickWindow) SceneGraphBackend() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17sceneGraphBackendEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QQuickWindow_SceneGraphBackend() string {
	var nilthis *QQuickWindow
	rv := nilthis.SceneGraphBackend()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRectangleNode * createRectangleNode()
func (this *QQuickWindow) CreateRectangleNode() *QSGRectangleNode /*777 QSGRectangleNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createRectangleNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRectangleNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGImageNode * createImageNode()
func (this *QQuickWindow) CreateImageNode() *QSGImageNode /*777 QSGImageNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow15createImageNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGImageNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGNinePatchNode * createNinePatchNode()
func (this *QQuickWindow) CreateNinePatchNode() *QSGNinePatchNode /*777 QSGNinePatchNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK12QQuickWindow19createNinePatchNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNinePatchNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickwindow.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [4] QQuickWindow::TextRenderType textRenderType()
func (this *QQuickWindow) TextRenderType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14textRenderTypeEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}
func QQuickWindow_TextRenderType() int {
	var nilthis *QQuickWindow
	rv := nilthis.TextRenderType()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:180
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTextRenderType(enum QQuickWindow::TextRenderType)
func (this *QQuickWindow) SetTextRenderType(renderType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17setTextRenderTypeENS_14TextRenderTypeE", qtrt.FFI_TYPE_POINTER, renderType)
	qtrt.ErrPrint(err, rv)
}
func QQuickWindow_SetTextRenderType(renderType int) {
	var nilthis *QQuickWindow
	nilthis.SetTextRenderType(renderType)
}

// /usr/include/qt/QtQuick/qquickwindow.h:183
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frameSwapped()
func (this *QQuickWindow) FrameSwapped() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow12frameSwappedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphInitialized()
func (this *QQuickWindow) SceneGraphInitialized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphInitializedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphInvalidated()
func (this *QQuickWindow) SceneGraphInvalidated() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphInvalidatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beforeSynchronizing()
func (this *QQuickWindow) BeforeSynchronizing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow19beforeSynchronizingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterSynchronizing()
func (this *QQuickWindow) AfterSynchronizing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow18afterSynchronizingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beforeRendering()
func (this *QQuickWindow) BeforeRendering() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15beforeRenderingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterRendering()
func (this *QQuickWindow) AfterRendering() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14afterRenderingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterAnimating()
func (this *QQuickWindow) AfterAnimating() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14afterAnimatingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphAboutToStop()
func (this *QQuickWindow) SceneGraphAboutToStop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphAboutToStopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void colorChanged(const QColor &)
func (this *QQuickWindow) ColorChanged(arg0 qtgui.QColor_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QColor_PTR() != nil {
		convArg0 = arg0.QColor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow12colorChangedERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusItemChanged()
func (this *QQuickWindow) ActiveFocusItemChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow22activeFocusItemChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphError(QQuickWindow::SceneGraphError, const QString &)
func (this *QQuickWindow) SceneGraphError(error int, message string) {
	var tmpArg1 = qtcore.NewQString_5(message)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15sceneGraphErrorENS_15SceneGraphErrorERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QQuickWindow) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseResources()
func (this *QQuickWindow) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:207
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void exposeEvent(QExposeEvent *)
func (this *QQuickWindow) ExposeEvent(arg0 qtgui.QExposeEvent_ITF /*777 QExposeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QExposeEvent_PTR() != nil {
		convArg0 = arg0.QExposeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow11exposeEventEP12QExposeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QQuickWindow) ResizeEvent(arg0 qtgui.QResizeEvent_ITF /*777 QResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QResizeEvent_PTR() != nil {
		convArg0 = arg0.QResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QQuickWindow) ShowEvent(arg0 qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QShowEvent_PTR() != nil {
		convArg0 = arg0.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:211
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QQuickWindow) HideEvent(arg0 qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QHideEvent_PTR() != nil {
		convArg0 = arg0.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QQuickWindow) FocusInEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QQuickWindow) FocusOutEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QQuickWindow) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QQuickWindow) KeyPressEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QQuickWindow) KeyReleaseEvent(arg0 qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QKeyEvent_PTR() != nil {
		convArg0 = arg0.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QQuickWindow) MousePressEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickWindow) MouseReleaseEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickWindow) MouseDoubleClickEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QQuickWindow) MouseMoveEvent(arg0 qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMouseEvent_PTR() != nil {
		convArg0 = arg0.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:225
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QQuickWindow) WheelEvent(arg0 qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWheelEvent_PTR() != nil {
		convArg0 = arg0.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN12QQuickWindow10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

type QQuickWindow__CreateTextureOption = int

const QQuickWindow__TextureHasAlphaChannel QQuickWindow__CreateTextureOption = 1
const QQuickWindow__TextureHasMipmaps QQuickWindow__CreateTextureOption = 2
const QQuickWindow__TextureOwnsGLTexture QQuickWindow__CreateTextureOption = 4
const QQuickWindow__TextureCanUseAtlas QQuickWindow__CreateTextureOption = 8
const QQuickWindow__TextureIsOpaque QQuickWindow__CreateTextureOption = 16

type QQuickWindow__RenderStage = int

const QQuickWindow__BeforeSynchronizingStage QQuickWindow__RenderStage = 0
const QQuickWindow__AfterSynchronizingStage QQuickWindow__RenderStage = 1
const QQuickWindow__BeforeRenderingStage QQuickWindow__RenderStage = 2
const QQuickWindow__AfterRenderingStage QQuickWindow__RenderStage = 3
const QQuickWindow__AfterSwapStage QQuickWindow__RenderStage = 4
const QQuickWindow__NoStage QQuickWindow__RenderStage = 5

type QQuickWindow__SceneGraphError = int

const QQuickWindow__ContextNotAvailable QQuickWindow__SceneGraphError = 1

type QQuickWindow__TextRenderType = int

const QQuickWindow__QtTextRendering QQuickWindow__TextRenderType = 0
const QQuickWindow__NativeTextRendering QQuickWindow__TextRenderType = 1

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
