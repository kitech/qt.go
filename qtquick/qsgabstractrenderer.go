package qtquick

// /usr/include/qt/QtQuick/qsgabstractrenderer.h
// #include <qsgabstractrenderer.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QSGAbstractRenderer struct {
	*qtcore.QObject
}

func (this *QSGAbstractRenderer) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSGAbstractRenderer) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSGAbstractRendererFromPointer(cthis unsafe.Pointer) *QSGAbstractRenderer {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSGAbstractRenderer{bcthis0}
}
func (*QSGAbstractRenderer) NewFromPointer(cthis unsafe.Pointer) *QSGAbstractRenderer {
	return NewQSGAbstractRendererFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSGAbstractRenderer) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGAbstractRenderer()
func DeleteQSGAbstractRenderer(*QSGAbstractRenderer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRendererD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRootNode(QSGRootNode *)
func (this *QSGAbstractRenderer) SetRootNode(node *QSGRootNode /*777 QSGRootNode **/) {
	var convArg0 = node.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer11setRootNodeEP11QSGRootNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRootNode * rootNode()
func (this *QSGAbstractRenderer) RootNode() *QSGRootNode /*777 QSGRootNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer8rootNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGRootNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDeviceRect(const QRect &)
func (this *QSGAbstractRenderer) SetDeviceRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer13setDeviceRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:70
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setDeviceRect(const QSize &)
func (this *QSGAbstractRenderer) SetDeviceRect_1(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer13setDeviceRectERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:71
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect deviceRect()
func (this *QSGAbstractRenderer) DeviceRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer10deviceRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewportRect(const QRect &)
func (this *QSGAbstractRenderer) SetViewportRect(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer15setViewportRectERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setViewportRect(const QSize &)
func (this *QSGAbstractRenderer) SetViewportRect_1(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer15setViewportRectERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:75
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewportRect()
func (this *QSGAbstractRenderer) ViewportRect() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer12viewportRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProjectionMatrixToRect(const QRectF &)
func (this *QSGAbstractRenderer) SetProjectionMatrixToRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer25setProjectionMatrixToRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setProjectionMatrix(const QMatrix4x4 &)
func (this *QSGAbstractRenderer) SetProjectionMatrix(matrix *qtgui.QMatrix4x4) {
	var convArg0 = matrix.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer19setProjectionMatrixERK10QMatrix4x4", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:79
// index:0
// Public Visibility=Default Availability=Available
// [68] QMatrix4x4 projectionMatrix()
func (this *QSGAbstractRenderer) ProjectionMatrix() *qtgui.QMatrix4x4 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer16projectionMatrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClearColor(const QColor &)
func (this *QSGAbstractRenderer) SetClearColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer13setClearColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor clearColor()
func (this *QSGAbstractRenderer) ClearColor() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer10clearColorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QSGAbstractRenderer::ClearMode clearMode()
func (this *QSGAbstractRenderer) ClearMode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK19QSGAbstractRenderer9clearModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void renderScene(uint)
func (this *QSGAbstractRenderer) RenderScene(fboId uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer11renderSceneEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), fboId)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void sceneGraphChanged()
func (this *QSGAbstractRenderer) SceneGraphChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRenderer17sceneGraphChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:93
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSGAbstractRenderer(QObject *)
func NewQSGAbstractRenderer(parent *qtcore.QObject /*777 QObject **/) *QSGAbstractRenderer {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN19QSGAbstractRendererC1EP7QObject", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGAbstractRendererFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

type QSGAbstractRenderer__ClearModeBit = int

const QSGAbstractRenderer__ClearColorBuffer QSGAbstractRenderer__ClearModeBit = 1
const QSGAbstractRenderer__ClearDepthBuffer QSGAbstractRenderer__ClearModeBit = 2
const QSGAbstractRenderer__ClearStencilBuffer QSGAbstractRenderer__ClearModeBit = 4

//  body block end
