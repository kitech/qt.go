package qtquick

// /usr/include/qt/QtQuick/qquickframebufferobject.h
// #include <qquickframebufferobject.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 172
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

// void geometryChanged(const QRectF &, const QRectF &)
func (this *QQuickFramebufferObject) InheritGeometryChanged(f func(newGeometry *qtcore.QRectF, oldGeometry *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "geometryChanged", f)
}

/*

 */
type QQuickFramebufferObject struct {
	*QQuickItem
}
type QQuickFramebufferObject_ITF interface {
	QQuickItem_ITF
	QQuickFramebufferObject_PTR() *QQuickFramebufferObject
}

func (ptr *QQuickFramebufferObject) QQuickFramebufferObject_PTR() *QQuickFramebufferObject { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickFramebufferObject) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickFramebufferObject(QQuickItem *)

/*
Constructs a new QQuickFramebufferObject with parent parent.
*/
func (*QQuickFramebufferObject) NewForInherit(parent QQuickItem_ITF /*777 QQuickItem **/) *QQuickFramebufferObject {
	return NewQQuickFramebufferObject(parent)
}
func NewQQuickFramebufferObject(parent QQuickItem_ITF /*777 QQuickItem **/) *QQuickFramebufferObject {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQuickItem_PTR() != nil {
		convArg0 = parent.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObjectC2EP10QQuickItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickFramebufferObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickFramebufferObject")
	return gothis
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickFramebufferObject(QQuickItem *)

/*
Constructs a new QQuickFramebufferObject with parent parent.
*/
func (*QQuickFramebufferObject) NewForInheritp() *QQuickFramebufferObject {
	return NewQQuickFramebufferObjectp()
}
func NewQQuickFramebufferObjectp() *QQuickFramebufferObject {
	// arg: 0, QQuickItem *=Pointer, QQuickItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObjectC2EP10QQuickItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickFramebufferObjectFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickFramebufferObject")
	return gothis
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool textureFollowsItemSize() const

/*

 */
func (this *QQuickFramebufferObject) TextureFollowsItemSize() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject22textureFollowsItemSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTextureFollowsItemSize(bool)

/*

 */
func (this *QQuickFramebufferObject) SetTextureFollowsItemSize(follows bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObject25setTextureFollowsItemSizeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), follows)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:83
// index:0
// Public Visibility=Default Availability=Available
// [1] bool mirrorVertically() const

/*

 */
func (this *QQuickFramebufferObject) MirrorVertically() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject16mirrorVerticallyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMirrorVertically(bool)

/*

 */
func (this *QQuickFramebufferObject) SetMirrorVertically(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObject19setMirrorVerticallyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QQuickFramebufferObject::Renderer * createRenderer() const

/*
Reimplement this function to create a renderer used to render into the FBO.

This function will be called on the rendering thread while the GUI thread is blocked.
*/
func (this *QQuickFramebufferObject) CreateRenderer() unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject14createRendererEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isTextureProvider() const

/*
Reimplemented from QQuickItem::isTextureProvider().
*/
func (this *QQuickFramebufferObject) IsTextureProvider() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject17isTextureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGTextureProvider * textureProvider() const

/*
Reimplemented from QQuickItem::textureProvider().
*/
func (this *QQuickFramebufferObject) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QQuickFramebufferObject15textureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void releaseResources()

/*
Reimplemented from QQuickItem::releaseResources().
*/
func (this *QQuickFramebufferObject) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObject16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:93
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void geometryChanged(const QRectF &, const QRectF &)

/*

 */
func (this *QQuickFramebufferObject) GeometryChanged(newGeometry qtcore.QRectF_ITF, oldGeometry qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if newGeometry != nil && newGeometry.QRectF_PTR() != nil {
		convArg0 = newGeometry.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if oldGeometry != nil && oldGeometry.QRectF_PTR() != nil {
		convArg1 = oldGeometry.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObject15geometryChangedERK6QRectFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void textureFollowsItemSizeChanged(bool)

/*

 */
func (this *QQuickFramebufferObject) TextureFollowsItemSizeChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObject29textureFollowsItemSizeChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickframebufferobject.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mirrorVerticallyChanged(bool)

/*

 */
func (this *QQuickFramebufferObject) MirrorVerticallyChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObject23mirrorVerticallyChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQQuickFramebufferObject(this *QQuickFramebufferObject) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QQuickFramebufferObjectD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11549() {
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
