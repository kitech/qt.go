package qtquick

// /usr/include/qt/QtQuick/qsgengine.h
// #include <qsgengine.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QSGEngine struct {
	*qtcore.QObject
}

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
// Public virtual
// const QMetaObject * metaObject()
func (this *QSGEngine) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSGEngine10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgengine.h:70
// index:0
// Public
// void QSGEngine(QObject *)
func NewQSGEngine(parent *qtcore.QObject /*777 QObject **/) *QSGEngine {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSGEngineC2EP7QObject", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGEngineFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgengine.h:71
// index:0
// Public virtual
// void ~QSGEngine()
func DeleteQSGEngine(*QSGEngine) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSGEngineD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgengine.h:74
// index:0
// Public
// void invalidate()
func (this *QSGEngine) Invalidate() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QSGEngine10invalidateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgengine.h:76
// index:0
// Public
// QSGAbstractRenderer * createRenderer()
func (this *QSGEngine) CreateRenderer() *QSGAbstractRenderer /*777 QSGAbstractRenderer **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSGEngine14createRendererEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGAbstractRendererFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgengine.h:79
// index:0
// Public
// QSGRendererInterface * rendererInterface()
func (this *QSGEngine) RendererInterface() *QSGRendererInterface /*777 QSGRendererInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSGEngine17rendererInterfaceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGRendererInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgengine.h:80
// index:0
// Public
// QSGRectangleNode * createRectangleNode()
func (this *QSGEngine) CreateRectangleNode() *QSGRectangleNode /*777 QSGRectangleNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSGEngine19createRectangleNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGRectangleNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgengine.h:81
// index:0
// Public
// QSGImageNode * createImageNode()
func (this *QSGEngine) CreateImageNode() *QSGImageNode /*777 QSGImageNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSGEngine15createImageNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGImageNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgengine.h:82
// index:0
// Public
// QSGNinePatchNode * createNinePatchNode()
func (this *QSGEngine) CreateNinePatchNode() *QSGNinePatchNode /*777 QSGNinePatchNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QSGEngine19createNinePatchNodeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGNinePatchNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

type QSGEngine__CreateTextureOption = int

const QSGEngine__TextureHasAlphaChannel QSGEngine__CreateTextureOption = 1
const QSGEngine__TextureOwnsGLTexture QSGEngine__CreateTextureOption = 4
const QSGEngine__TextureCanUseAtlas QSGEngine__CreateTextureOption = 8
const QSGEngine__TextureIsOpaque QSGEngine__CreateTextureOption = 16

//  body block end
