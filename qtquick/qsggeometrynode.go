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
// Public
// void QSGGeometryNode()
func NewQSGGeometryNode() *QSGGeometryNode {
	cthis := qtrt.Calloc(1, 256) // 144
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSGGeometryNodeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGGeometryNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:232
// index:0
// Public virtual
// void ~QSGGeometryNode()
func DeleteQSGGeometryNode(*QSGGeometryNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSGGeometryNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:234
// index:0
// Public
// void setMaterial(class QSGMaterial *)
func (this *QSGGeometryNode) SetMaterial(material *QSGMaterial /*777 QSGMaterial **/) {
	var convArg0 = material.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSGGeometryNode11setMaterialEP11QSGMaterial", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:235
// index:0
// Public inline
// QSGMaterial * material()
func (this *QSGGeometryNode) Material() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSGGeometryNode8materialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:237
// index:0
// Public
// void setOpaqueMaterial(class QSGMaterial *)
func (this *QSGGeometryNode) SetOpaqueMaterial(material *QSGMaterial /*777 QSGMaterial **/) {
	var convArg0 = material.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSGGeometryNode17setOpaqueMaterialEP11QSGMaterial", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:238
// index:0
// Public inline
// QSGMaterial * opaqueMaterial()
func (this *QSGGeometryNode) OpaqueMaterial() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSGGeometryNode14opaqueMaterialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:240
// index:0
// Public
// QSGMaterial * activeMaterial()
func (this *QSGGeometryNode) ActiveMaterial() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSGGeometryNode14activeMaterialEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:242
// index:0
// Public
// void setRenderOrder(int)
func (this *QSGGeometryNode) SetRenderOrder(order int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSGGeometryNode14setRenderOrderEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), order)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:243
// index:0
// Public inline
// int renderOrder()
func (this *QSGGeometryNode) RenderOrder() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSGGeometryNode11renderOrderEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtQuick/qsgnode.h:245
// index:0
// Public
// void setInheritedOpacity(qreal)
func (this *QSGGeometryNode) SetInheritedOpacity(opacity float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QSGGeometryNode19setInheritedOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:246
// index:0
// Public inline
// qreal inheritedOpacity()
func (this *QSGGeometryNode) InheritedOpacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QSGGeometryNode16inheritedOpacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

//  body block end
