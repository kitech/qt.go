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
// extern C begin: 6
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

type QSGTransformNode struct {
	*QSGNode
}
type QSGTransformNode_ITF interface {
	QSGNode_ITF
	QSGTransformNode_PTR() *QSGTransformNode
}

func (ptr *QSGTransformNode) QSGTransformNode_PTR() *QSGTransformNode { return ptr }

func (this *QSGTransformNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGNode.GetCthis()
	}
}
func (this *QSGTransformNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGNode = NewQSGNodeFromPointer(cthis)
}
func NewQSGTransformNodeFromPointer(cthis unsafe.Pointer) *QSGTransformNode {
	bcthis0 := NewQSGNodeFromPointer(cthis)
	return &QSGTransformNode{bcthis0}
}
func (*QSGTransformNode) NewFromPointer(cthis unsafe.Pointer) *QSGTransformNode {
	return NewQSGTransformNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGTransformNode()
func NewQSGTransformNode() *QSGTransformNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGTransformNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGTransformNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:285
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGTransformNode()
func DeleteQSGTransformNode(this *QSGTransformNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 216)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMatrix(const QMatrix4x4 &)
func (this *QSGTransformNode) SetMatrix(matrix qtgui.QMatrix4x4_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix4x4_PTR() != nil {
		convArg0 = matrix.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNode9setMatrixERK10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:288
// index:0
// Public inline Visibility=Default Availability=Available
// [68] const QMatrix4x4 & matrix() const
func (this *QSGTransformNode) Matrix() *qtgui.QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSGTransformNode6matrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix4x4)
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:290
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCombinedMatrix(const QMatrix4x4 &)
func (this *QSGTransformNode) SetCombinedMatrix(matrix qtgui.QMatrix4x4_ITF) {
	var convArg0 unsafe.Pointer
	if matrix != nil && matrix.QMatrix4x4_PTR() != nil {
		convArg0 = matrix.QMatrix4x4_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGTransformNode17setCombinedMatrixERK10QMatrix4x4", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:291
// index:0
// Public inline Visibility=Default Availability=Available
// [68] const QMatrix4x4 & combinedMatrix() const
func (this *QSGTransformNode) CombinedMatrix() *qtgui.QMatrix4x4 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSGTransformNode14combinedMatrixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQMatrix4x4FromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQMatrix4x4)
	return rv2
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
