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
// extern C begin: 26
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

/*

 */
type QSGBasicGeometryNode struct {
	*QSGNode
}
type QSGBasicGeometryNode_ITF interface {
	QSGNode_ITF
	QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode
}

func (ptr *QSGBasicGeometryNode) QSGBasicGeometryNode_PTR() *QSGBasicGeometryNode { return ptr }

func (this *QSGBasicGeometryNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGBasicGeometryNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGBasicGeometryNodeFromPointer(cthis unsafe.Pointer) *QSGBasicGeometryNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGBasicGeometryNode{bcthis0}
}
func (*QSGBasicGeometryNode) NewFromPointer(cthis unsafe.Pointer) *QSGBasicGeometryNode {
	return NewQSGBasicGeometryNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:198
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGBasicGeometryNode()

/*

 */
func DeleteQSGBasicGeometryNode(this *QSGBasicGeometryNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGBasicGeometryNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 112)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:200
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(QSGGeometry *)

/*

 */
func (this *QSGBasicGeometryNode) SetGeometry(geometry QSGGeometry_ITF /*777 QSGGeometry **/) {
	var convArg0 unsafe.Pointer
	if geometry != nil && geometry.QSGGeometry_PTR() != nil {
		convArg0 = geometry.QSGGeometry_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode11setGeometryEP11QSGGeometry", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:201
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QSGGeometry * geometry() const

/*

 */
func (this *QSGBasicGeometryNode) Geometry() *QSGGeometry /*777 const QSGGeometry **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGBasicGeometryNode8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGGeometryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:202
// index:1
// Public inline Visibility=Default Availability=Available
// [8] QSGGeometry * geometry()

/*

 */
func (this *QSGBasicGeometryNode) Geometry_1() *QSGGeometry /*777 QSGGeometry **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGGeometryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:204
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QMatrix4x4 * matrix() const

/*

 */
func (this *QSGBasicGeometryNode) Matrix() *qtgui.QMatrix4x4 /*777 const QMatrix4x4 **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGBasicGeometryNode6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:205
// index:0
// Public inline Visibility=Default Availability=Available
// [8] const QSGClipNode * clipList() const

/*

 */
func (this *QSGBasicGeometryNode) ClipList() *QSGClipNode /*777 const QSGClipNode **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QSGBasicGeometryNode8clipListEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGClipNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgnode.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRendererMatrix(const QMatrix4x4 *)

/*

 */
func (this *QSGBasicGeometryNode) SetRendererMatrix(m qtgui.QMatrix4x4_ITF /*777 const QMatrix4x4 **/) {
	var convArg0 unsafe.Pointer
	if m != nil && m.QMatrix4x4_PTR() != nil {
		convArg0 = m.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode17setRendererMatrixEPK10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRendererClipList(const QSGClipNode *)

/*

 */
func (this *QSGBasicGeometryNode) SetRendererClipList(c QSGClipNode_ITF /*777 const QSGClipNode **/) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QSGClipNode_PTR() != nil {
		convArg0 = c.QSGClipNode_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode19setRendererClipListEPK11QSGClipNode", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:211
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QSGBasicGeometryNode(QSGNode::NodeType)

/*

 */
func NewQSGBasicGeometryNode(type_ int) *QSGBasicGeometryNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QSGBasicGeometryNodeC2EN7QSGNode8NodeTypeE", qtrt.FFI_TYPE_POINTER, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGBasicGeometryNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGBasicGeometryNode)
	return gothis
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
