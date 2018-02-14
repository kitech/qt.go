package qtquick

// /usr/include/qt/QtQuick/qsgrectanglenode.h
// #include <qsgrectanglenode.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

type QSGRectangleNode struct {
	*QSGGeometryNode
}
type QSGRectangleNode_ITF interface {
	QSGGeometryNode_ITF
	QSGRectangleNode_PTR() *QSGRectangleNode
}

func (ptr *QSGRectangleNode) QSGRectangleNode_PTR() *QSGRectangleNode { return ptr }

func (this *QSGRectangleNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGGeometryNode.GetCthis()
	}
}
func (this *QSGRectangleNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGGeometryNode = NewQSGGeometryNodeFromPointer(cthis)
}
func NewQSGRectangleNodeFromPointer(cthis unsafe.Pointer) *QSGRectangleNode {
	bcthis0 := NewQSGGeometryNodeFromPointer(cthis)
	return &QSGRectangleNode{bcthis0}
}
func (*QSGRectangleNode) NewFromPointer(cthis unsafe.Pointer) *QSGRectangleNode {
	return NewQSGRectangleNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:50
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [-2] void ~QSGRectangleNode()
func DeleteQSGRectangleNode(this *QSGRectangleNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGRectangleNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 144)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:52
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)
func (this *QSGRectangleNode) SetRect(rect qtcore.QRectF_ITF) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGRectangleNode7setRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:53
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QSGRectangleNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGRectangleNode7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:54
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QSGRectangleNode) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSGRectangleNode4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QSGRectangleNode) SetColor(color qtgui.QColor_ITF) {
	var convArg0 = color.QColor_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QSGRectangleNode8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QColor color()
func (this *QSGRectangleNode) Color() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QSGRectangleNode5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
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
