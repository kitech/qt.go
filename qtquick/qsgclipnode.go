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
// extern C begin: 11
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
type QSGClipNode struct {
	*QSGBasicGeometryNode
}
type QSGClipNode_ITF interface {
	QSGBasicGeometryNode_ITF
	QSGClipNode_PTR() *QSGClipNode
}

func (ptr *QSGClipNode) QSGClipNode_PTR() *QSGClipNode { return ptr }

func (this *QSGClipNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGBasicGeometryNode.GetCthis()
	}
}
func (this *QSGClipNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGBasicGeometryNode = NewQSGBasicGeometryNodeFromPointer(cthis)
}
func NewQSGClipNodeFromPointer(cthis unsafe.Pointer) *QSGClipNode {
	bcthis0 := NewQSGBasicGeometryNodeFromPointer(cthis)
	return &QSGClipNode{bcthis0}
}
func (*QSGClipNode) NewFromPointer(cthis unsafe.Pointer) *QSGClipNode {
	return NewQSGClipNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgnode.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGClipNode()

/*

 */
func (*QSGClipNode) NewForInherit() *QSGClipNode {
	return NewQSGClipNode()
}
func NewQSGClipNode() *QSGClipNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGClipNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGClipNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGClipNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgnode.h:265
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGClipNode()

/*

 */
func DeleteQSGClipNode(this *QSGClipNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGClipNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 152)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgnode.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIsRectangular(bool)

/*

 */
func (this *QSGClipNode) SetIsRectangular(rectHint bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGClipNode16setIsRectangularEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rectHint)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:268
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRectangular() const

/*

 */
func (this *QSGClipNode) IsRectangular() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGClipNode13isRectangularEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qsgnode.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClipRect(const QRectF &)

/*

 */
func (this *QSGClipNode) SetClipRect(arg0 qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGClipNode11setClipRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgnode.h:271
// index:0
// Public inline Visibility=Default Availability=Available
// [32] QRectF clipRect() const

/*

 */
func (this *QSGClipNode) ClipRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGClipNode8clipRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
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
