package qtquick

// /usr/include/qt/QtQuick/qsgnode.h
// #include <qsgnode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 25
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
type QSGBasicGeometryNode struct {
	*QSGNode
}

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
// Public virtual
// void ~QSGBasicGeometryNode()
func DeleteQSGBasicGeometryNode(*QSGBasicGeometryNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGBasicGeometryNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:200
// index:0
// Public
// void setGeometry(QSGGeometry *)
func (this *QSGBasicGeometryNode) SetGeometry(geometry *QSGGeometry /*777 QSGGeometry **/) {
	var convArg0 = geometry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode11setGeometryEP11QSGGeometry", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:201
// index:0
// Public inline
// const QSGGeometry * geometry()
func (this *QSGBasicGeometryNode) Geometry() *QSGGeometry /*777 const QSGGeometry **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGBasicGeometryNode8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGGeometryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:202
// index:1
// Public inline
// QSGGeometry * geometry()
func (this *QSGBasicGeometryNode) Geometry_1() *QSGGeometry /*777 QSGGeometry **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGGeometryFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:204
// index:0
// Public inline
// const QMatrix4x4 * matrix()
func (this *QSGBasicGeometryNode) Matrix() *qtgui.QMatrix4x4 /*777 const QMatrix4x4 **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGBasicGeometryNode6matrixEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:205
// index:0
// Public inline
// const QSGClipNode * clipList()
func (this *QSGBasicGeometryNode) ClipList() *QSGClipNode /*777 const QSGClipNode **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QSGBasicGeometryNode8clipListEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGClipNodeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:207
// index:0
// Public inline
// void setRendererMatrix(const QMatrix4x4 *)
func (this *QSGBasicGeometryNode) SetRendererMatrix(m *qtgui.QMatrix4x4 /*777 const QMatrix4x4 **/) {
	var convArg0 = m.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode17setRendererMatrixEPK10QMatrix4x4", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:208
// index:0
// Public inline
// void setRendererClipList(const QSGClipNode *)
func (this *QSGBasicGeometryNode) SetRendererClipList(c *QSGClipNode /*777 const QSGClipNode **/) {
	var convArg0 = c.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGBasicGeometryNode19setRendererClipListEPK11QSGClipNode", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:211
// index:0
// Protected
// void QSGBasicGeometryNode(enum QSGNode::NodeType)
func NewQSGBasicGeometryNode(type_ int) *QSGBasicGeometryNode {
	cthis := qtrt.Calloc(1, 256) // 112
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QSGBasicGeometryNodeC2EN7QSGNode8NodeTypeE", ffiqt.FFI_TYPE_VOID, cthis, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGBasicGeometryNodeFromPointer(cthis)
	return gothis
}

//  body block end
