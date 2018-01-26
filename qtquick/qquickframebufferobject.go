package qtquick

// /usr/include/qt/QtQuick/qquickframebufferobject.h
// #include <qquickframebufferobject.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 166
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "mkuse/cffiqt"
import "gopp"
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
type QQuickFramebufferObject struct {
	*QQuickItem
}

func (this *QQuickFramebufferObject) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QQuickItem.GetCthis()
	}
}
func (this *QQuickFramebufferObject) SetCthis(cthis unsafe.Pointer) {
	this.QQuickItem = NewQQuickItemFromPointer(cthis)
}
func NewQQuickFramebufferObjectFromPointer(cthis unsafe.Pointer) *QQuickFramebufferObject {
	bcthis0 := NewQQuickItemFromPointer(cthis)
	return &QQuickFramebufferObject{bcthis0}
}
func (*QQuickFramebufferObject) NewFromPointer(cthis unsafe.Pointer) *QQuickFramebufferObject {
	return NewQQuickFramebufferObjectFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QQuickFramebufferObject) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:78
// index:0
// Public
// void QQuickFramebufferObject(class QQuickItem *)
func NewQQuickFramebufferObject(parent *QQuickItem /*777 QQuickItem **/) *QQuickFramebufferObject {
	cthis := qtrt.Calloc(1, 256) // 32
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObjectC1EP10QQuickItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickFramebufferObjectFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:80
// index:0
// Public
// bool textureFollowsItemSize()
func (this *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject22textureFollowsItemSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:81
// index:0
// Public
// void setTextureFollowsItemSize(_Bool)
func (this *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObject25setTextureFollowsItemSizeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), follows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:83
// index:0
// Public
// bool mirrorVertically()
func (this *QQuickFramebufferObject) MirrorVertically() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject16mirrorVerticallyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:84
// index:0
// Public
// void setMirrorVertically(_Bool)
func (this *QQuickFramebufferObject) SetMirrorVertically(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObject19setMirrorVerticallyEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:86
// index:0
// Public pure virtual
// QQuickFramebufferObject::Renderer * createRenderer()
func (this *QQuickFramebufferObject) CreateRenderer() unsafe.Pointer /*666*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject14createRendererEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:88
// index:0
// Public virtual
// bool isTextureProvider()
func (this *QQuickFramebufferObject) IsTextureProvider() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject17isTextureProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:89
// index:0
// Public virtual
// QSGTextureProvider * textureProvider()
func (this *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject15textureProviderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:90
// index:0
// Public virtual
// void releaseResources()
func (this *QQuickFramebufferObject) ReleaseResources() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObject16releaseResourcesEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:93
// index:0
// Protected virtual
// void geometryChanged(const class QRectF &, const class QRectF &)
func (this *QQuickFramebufferObject) GeometryChanged(newGeometry *qtcore.QRectF, oldGeometry *qtcore.QRectF) {
	var convArg0 = newGeometry.GetCthis()
	var convArg1 = oldGeometry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObject15geometryChangedERK6QRectFS2_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:99
// index:0
// Public
// void textureFollowsItemSizeChanged(_Bool)
func (this *QQuickFramebufferObject) TextureFollowsItemSizeChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObject29textureFollowsItemSizeChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:100
// index:0
// Public
// void mirrorVerticallyChanged(_Bool)
func (this *QQuickFramebufferObject) MirrorVerticallyChanged(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QQuickFramebufferObject23mirrorVerticallyChangedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
