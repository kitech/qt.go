package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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

type QSGGeometryNode struct {
	*QSGBasicGeometryNode
}
type QSGGeometryNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGGeometryNode_PTR() *QSGGeometryNode
}

func (ptr *QSGGeometryNode) QSGGeometryNode_PTR() *QSGGeometryNode { return ptr }

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
	qtrt.ErrPrint(err, rv)
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
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaterial(QSGMaterial *)
func (this *QSGGeometryNode) SetMaterial(material QSGMaterial_ITF /*777 QSGMaterial **/) {
	var convArg0 unsafe.Pointer
	if material != nil && material.QSGMaterial_PTR() != nil {
		convArg0 = material.QSGMaterial_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode11setMaterialEP11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGMaterial * material() const
func (this *QSGGeometryNode) Material() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode8materialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpaqueMaterial(QSGMaterial *)
func (this *QSGGeometryNode) SetOpaqueMaterial(material QSGMaterial_ITF /*777 QSGMaterial **/) {
	var convArg0 unsafe.Pointer
	if material != nil && material.QSGMaterial_PTR() != nil {
		convArg0 = material.QSGMaterial_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode17setOpaqueMaterialEP11QSGMaterial", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:238
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QSGMaterial * opaqueMaterial() const
func (this *QSGGeometryNode) OpaqueMaterial() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode14opaqueMaterialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:240
// index:0
// Public Visibility=Default Availability=Available
// [8] QSGMaterial * activeMaterial() const
func (this *QSGGeometryNode) ActiveMaterial() *QSGMaterial /*777 QSGMaterial **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode14activeMaterialEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRenderOrder(int)
func (this *QSGGeometryNode) SetRenderOrder(order int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode14setRenderOrderEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), order)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:243
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int renderOrder() const
func (this *QSGGeometryNode) RenderOrder() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode11renderOrderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgnode.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInheritedOpacity(qreal)
func (this *QSGGeometryNode) SetInheritedOpacity(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QSGGeometryNode19setInheritedOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:246
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal inheritedOpacity() const
func (this *QSGGeometryNode) InheritedOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QSGGeometryNode16inheritedOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

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
