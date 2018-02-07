package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
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
		gopp.KeepMe()
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

//  ext block end

//  body block begin

type QSGGeometryNode struct {
	*QSGBasicGeometryNode
}

func (this *QSGGeometryNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGBasicGeometryNode.GetCthis()
	}
}
func (this *QSGGeometryNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGBasicGeometryNode = NewQSGBasicGeometryNodeFromPointer(cthis)
}
func NewQSGGeometryNodeFromPointer(cthis unsafe.Pointer) *QSGGeometryNode {
	bcthis0 := NewQSGBasicGeometryNodeFromPointer(cthis)
	return &QSGGeometryNode{bcthis0}
}
func (*QSGGeometryNode) NewFromPointer(cthis unsafe.Pointer) *QSGGeometryNode {
	return NewQSGGeometryNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGGeometryNode()
func NewQSGGeometryNode() *QSGGeometryNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGGeometryNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGGeometryNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:232
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGGeometryNode()
func DeleteQSGGeometryNode(this *QSGGeometryNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 144)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaterial(QSGMaterial *)
func (this *QSGGeometryNode) SetMaterial(material *QSGMaterial /*777 QSGMaterial **/) {
	var convArg0 = material.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode11setMaterialEP11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGMaterial * material()
func (this *QSGGeometryNode) Material() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode8materialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueMaterial(QSGMaterial *)
func (this *QSGGeometryNode) SetOpaqueMaterial(material *QSGMaterial /*777 QSGMaterial **/) {
	var convArg0 = material.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode17setOpaqueMaterialEP11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGMaterial * opaqueMaterial()
func (this *QSGGeometryNode) OpaqueMaterial() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode14opaqueMaterialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:240
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGMaterial * activeMaterial()
func (this *QSGGeometryNode) ActiveMaterial() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode14activeMaterialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderOrder(int)
func (this *QSGGeometryNode) SetRenderOrder(order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode14setRenderOrderEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:243
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int renderOrder()
func (this *QSGGeometryNode) RenderOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode11renderOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgnode.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInheritedOpacity(qreal)
func (this *QSGGeometryNode) SetInheritedOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode19setInheritedOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal inheritedOpacity()
func (this *QSGGeometryNode) InheritedOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode16inheritedOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

//  body block end
