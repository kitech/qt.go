package qtquick

// /usr/include/qt/QtQuick/qquickwindow.h
// #include <qquickwindow.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
import "qt.go/qtnetwork"
import "qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  ext block end

//  body block begin
type QQuickWindow struct {
	*qtgui.QWindow
}

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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QWindow *)
func NewQQuickWindow(parent *qtgui.QWindow /*777 QWindow **/) *QQuickWindow {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindowC2EP7QWindow", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:110
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QQuickWindow(QQuickRenderControl *)
func NewQQuickWindow_1(renderControl *QQuickRenderControl /*777 QQuickRenderControl **/) *QQuickWindow {
	var convArg0 = renderControl.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindowC2EP19QQuickRenderControl", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquickwindow.h:112
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickWindow()
func DeleteQQuickWindow(*QQuickWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * contentItem()
func (this *QQuickWindow) ContentItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow11contentItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * activeFocusItem()
func (this *QQuickWindow) ActiveFocusItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow15activeFocusItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * focusObject()
func (this *QQuickWindow) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow11focusObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * mouseGrabberItem()
func (this *QQuickWindow) MouseGrabberItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow16mouseGrabberItemEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:122
// index:0
// Public Visibility=Default Availability=Available
// [1] bool sendEvent(QQuickItem *, QEvent *)
func (this *QQuickWindow) SendEvent(arg0 *QQuickItem /*777 QQuickItem **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow9sendEventEP10QQuickItemP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:125
// index:0
// Public Visibility=Default Availability=Available
// [32] QImage grabWindow()
func (this *QQuickWindow) GrabWindow() *qtgui.QImage /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow10grabWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderTarget(uint, const QSize &)
func (this *QQuickWindow) SetRenderTarget(fboId uint, size *qtcore.QSize) {
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow15setRenderTargetEjRK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), fboId, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] uint renderTargetId()
func (this *QQuickWindow) RenderTargetId() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow14renderTargetIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtQuick/qquickwindow.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize renderTargetSize()
func (this *QQuickWindow) RenderTargetSize() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow16renderTargetSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetOpenGLState()
func (this *QQuickWindow) ResetOpenGLState() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow16resetOpenGLStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] QQmlIncubationController * incubationController()
func (this *QQuickWindow) IncubationController() *qtqml.QQmlIncubationController /*777 QQmlIncubationController **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow20incubationControllerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtqml.NewQQmlIncubationControllerFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:139
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleRoot()
func (this *QQuickWindow) AccessibleRoot() *qtgui.QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow14accessibleRootEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromImage(const QImage &)
func (this *QQuickWindow) CreateTextureFromImage(image *qtgui.QImage) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 = image.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow22createTextureFromImageERK6QImage", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearBeforeRendering(_Bool)
func (this *QQuickWindow) SetClearBeforeRendering(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow23setClearBeforeRenderingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:148
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clearBeforeRendering()
func (this *QQuickWindow) ClearBeforeRendering() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow20clearBeforeRenderingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QQuickWindow) SetColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow8setColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:151
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color()
func (this *QQuickWindow) Color() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow5colorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:153
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasDefaultAlphaBuffer()
func (this *QQuickWindow) HasDefaultAlphaBuffer() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow21hasDefaultAlphaBufferEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow21setDefaultAlphaBufferEb", ffiqt.FFI_TYPE_POINTER, useAlpha)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow26setPersistentOpenGLContextEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), persistent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:157
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentOpenGLContext()
func (this *QQuickWindow) IsPersistentOpenGLContext() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow25isPersistentOpenGLContextEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPersistentSceneGraph(_Bool)
func (this *QQuickWindow) SetPersistentSceneGraph(persistent bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow23setPersistentSceneGraphEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), persistent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPersistentSceneGraph()
func (this *QQuickWindow) IsPersistentSceneGraph() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow22isPersistentSceneGraphEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:163
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSceneGraphInitialized()
func (this *QQuickWindow) IsSceneGraphInitialized() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow23isSceneGraphInitializedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scheduleRenderJob(QRunnable *, enum QQuickWindow::RenderStage)
func (this *QQuickWindow) ScheduleRenderJob(job *qtcore.QRunnable /*777 QRunnable **/, schedule int) {
	var convArg0 = job.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow17scheduleRenderJobEP9QRunnableNS_11RenderStageE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, schedule)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:167
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal effectiveDevicePixelRatio()
func (this *QQuickWindow) EffectiveDevicePixelRatio() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow25effectiveDevicePixelRatioEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickwindow.h:169
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRendererInterface * rendererInterface()
func (this *QQuickWindow) RendererInterface() *QSGRendererInterface /*777 QSGRendererInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow17rendererInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGRendererInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:171
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setSceneGraphBackend(QSGRendererInterface::GraphicsApi)
func (this *QQuickWindow) SetSceneGraphBackend(api int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow20setSceneGraphBackendEN20QSGRendererInterface11GraphicsApiE", ffiqt.FFI_TYPE_POINTER, api)
	gopp.ErrPrint(err, rv)
}
func QQuickWindow_SetSceneGraphBackend(api int) {
	var nilthis *QQuickWindow
	nilthis.SetSceneGraphBackend(api)
}

// /usr/include/qt/QtQuick/qquickwindow.h:172
// index:1
// Public static Visibility=Default Availability=Available
// [-2] void setSceneGraphBackend(const QString &)
func (this *QQuickWindow) SetSceneGraphBackend_1(backend *qtcore.QString) {
	var convArg0 = backend.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow20setSceneGraphBackendERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QQuickWindow_SetSceneGraphBackend_1(backend *qtcore.QString) {
	var nilthis *QQuickWindow
	nilthis.SetSceneGraphBackend_1(backend)
}

// /usr/include/qt/QtQuick/qquickwindow.h:173
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString sceneGraphBackend()
func (this *QQuickWindow) SceneGraphBackend() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow17sceneGraphBackendEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QQuickWindow_SceneGraphBackend() *qtcore.QString /*123*/ {
	var nilthis *QQuickWindow
	rv := nilthis.SceneGraphBackend()
	return rv
}

// /usr/include/qt/QtQuick/qquickwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRectangleNode * createRectangleNode()
func (this *QQuickWindow) CreateRectangleNode() *QSGRectangleNode /*777 QSGRectangleNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow19createRectangleNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGRectangleNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:176
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGImageNode * createImageNode()
func (this *QQuickWindow) CreateImageNode() *QSGImageNode /*777 QSGImageNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow15createImageNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGImageNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:177
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGNinePatchNode * createNinePatchNode()
func (this *QQuickWindow) CreateNinePatchNode() *QSGNinePatchNode /*777 QSGNinePatchNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QQuickWindow19createNinePatchNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNinePatchNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickwindow.h:179
// index:0
// Public static Visibility=Default Availability=Available
// [4] QQuickWindow::TextRenderType textRenderType()
func (this *QQuickWindow) TextRenderType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow14textRenderTypeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow17setTextRenderTypeENS_14TextRenderTypeE", ffiqt.FFI_TYPE_POINTER, renderType)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow12frameSwappedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:185
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphInitialized()
func (this *QQuickWindow) SceneGraphInitialized() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphInitializedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:186
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphInvalidated()
func (this *QQuickWindow) SceneGraphInvalidated() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphInvalidatedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:187
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beforeSynchronizing()
func (this *QQuickWindow) BeforeSynchronizing() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow19beforeSynchronizingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:188
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterSynchronizing()
func (this *QQuickWindow) AfterSynchronizing() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow18afterSynchronizingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void beforeRendering()
func (this *QQuickWindow) BeforeRendering() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow15beforeRenderingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterRendering()
func (this *QQuickWindow) AfterRendering() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow14afterRenderingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void afterAnimating()
func (this *QQuickWindow) AfterAnimating() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow14afterAnimatingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphAboutToStop()
func (this *QQuickWindow) SceneGraphAboutToStop() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow21sceneGraphAboutToStopEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void colorChanged(const QColor &)
func (this *QQuickWindow) ColorChanged(arg0 *qtgui.QColor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow12colorChangedERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusItemChanged()
func (this *QQuickWindow) ActiveFocusItemChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow22activeFocusItemChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphError(QQuickWindow::SceneGraphError, const QString &)
func (this *QQuickWindow) SceneGraphError(error int, message *qtcore.QString) {
	var convArg1 = message.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow15sceneGraphErrorENS_15SceneGraphErrorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), error, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:201
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QQuickWindow) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:202
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseResources()
func (this *QQuickWindow) ReleaseResources() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow16releaseResourcesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:207
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void exposeEvent(QExposeEvent *)
func (this *QQuickWindow) ExposeEvent(arg0 *qtgui.QExposeEvent /*777 QExposeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow11exposeEventEP12QExposeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:208
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QQuickWindow) ResizeEvent(arg0 *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:210
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QQuickWindow) ShowEvent(arg0 *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:211
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QQuickWindow) HideEvent(arg0 *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:214
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QQuickWindow) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:215
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QQuickWindow) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:217
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QQuickWindow) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickwindow.h:218
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QQuickWindow) KeyPressEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:219
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QQuickWindow) KeyReleaseEvent(arg0 *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:220
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QQuickWindow) MousePressEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:221
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickWindow) MouseReleaseEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:222
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickWindow) MouseDoubleClickEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:223
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QQuickWindow) MouseMoveEvent(arg0 *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickwindow.h:225
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QQuickWindow) WheelEvent(arg0 *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QQuickWindow10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
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
