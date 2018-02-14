package qtquick

// /usr/include/qt/QtQuick/qsgengine.h
// #include <qsgengine.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGEngine struct {
	*qtcore.QObject
}
type QSGEngine_ITF interface {
	qtcore.QObject_ITF
	QSGEngine_PTR() *QSGEngine
}

func (ptr *QSGEngine) QSGEngine_PTR() *QSGEngine { return ptr }

func (this *QSGEngine) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QSGEngine) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQSGEngineFromPointer(cthis unsafe.Pointer) *QSGEngine {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QSGEngine{bcthis0}
}
func (*QSGEngine) NewFromPointer(cthis unsafe.Pointer) *QSGEngine {
	return NewQSGEngineFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgengine.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QSGEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGEngine(QObject *)
func NewQSGEngine(parent qtcore.QObject_ITF /*777 QObject **/) *QSGEngine {
	var convArg0 = parent.QObject_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSGEngineC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGEngineFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qsgengine.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGEngine()
func DeleteQSGEngine(this *QSGEngine) {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSGEngineD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgengine.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void invalidate()
func (this *QSGEngine) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN9QSGEngine10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgengine.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGAbstractRenderer * createRenderer()
func (this *QSGEngine) CreateRenderer() *QSGAbstractRenderer /*777 QSGAbstractRenderer **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine14createRendererEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGAbstractRendererFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:77
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromImage(const QImage &, QSGEngine::CreateTextureOptions)
func (this *QSGEngine) CreateTextureFromImage(image qtgui.QImage_ITF, options int) *QSGTexture /*777 QSGTexture **/ {
	var convArg0 = image.QImage_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine22createTextureFromImageERK6QImage6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:78
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGTexture * createTextureFromId(uint, const QSize &, QSGEngine::CreateTextureOptions)
func (this *QSGEngine) CreateTextureFromId(id uint, size qtcore.QSize_ITF, options int) *QSGTexture /*777 QSGTexture **/ {
	var convArg1 = size.QSize_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine19createTextureFromIdEjRK5QSize6QFlagsINS_19CreateTextureOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, convArg1, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRendererInterface * rendererInterface()
func (this *QSGEngine) RendererInterface() *QSGRendererInterface /*777 QSGRendererInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine17rendererInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRendererInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:80
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGRectangleNode * createRectangleNode()
func (this *QSGEngine) CreateRectangleNode() *QSGRectangleNode /*777 QSGRectangleNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine19createRectangleNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGRectangleNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:81
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGImageNode * createImageNode()
func (this *QSGEngine) CreateImageNode() *QSGImageNode /*777 QSGImageNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine15createImageNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGImageNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgengine.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGNinePatchNode * createNinePatchNode()
func (this *QSGEngine) CreateNinePatchNode() *QSGNinePatchNode /*777 QSGNinePatchNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK9QSGEngine19createNinePatchNodeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGNinePatchNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

type QSGEngine__CreateTextureOption = int

const QSGEngine__TextureHasAlphaChannel QSGEngine__CreateTextureOption = 1
const QSGEngine__TextureOwnsGLTexture QSGEngine__CreateTextureOption = 4
const QSGEngine__TextureCanUseAtlas QSGEngine__CreateTextureOption = 8
const QSGEngine__TextureIsOpaque QSGEngine__CreateTextureOption = 16

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
