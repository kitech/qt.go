package qtquick

// /usr/include/qt/QtQuick/qsgsimplerectnode.h
// #include <qsgsimplerectnode.h>
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

type QSGSimpleRectNode struct {
	*QSGGeometryNode
}

func (this *QSGSimpleRectNode) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QSGGeometryNode.GetCthis()
	}
}
func (this *QSGSimpleRectNode) SetCthis(cthis unsafe.Pointer) {
	this.QSGGeometryNode = NewQSGGeometryNodeFromPointer(cthis)
}
func NewQSGSimpleRectNodeFromPointer(cthis unsafe.Pointer) *QSGSimpleRectNode {
	bcthis0 := NewQSGGeometryNodeFromPointer(cthis)
	return &QSGSimpleRectNode{bcthis0}
}
func (*QSGSimpleRectNode) NewFromPointer(cthis unsafe.Pointer) *QSGSimpleRectNode {
	return NewQSGSimpleRectNodeFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:51
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGSimpleRectNode(const QRectF &, const QColor &)
func NewQSGSimpleRectNode(rect *qtcore.QRectF, color *qtgui.QColor) *QSGSimpleRectNode {
	var convArg0 = rect.GetCthis()
	var convArg1 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGSimpleRectNodeC2ERK6QRectFRK6QColor", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGSimpleRectNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGSimpleRectNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:52
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QSGSimpleRectNode()
func NewQSGSimpleRectNode_1() *QSGSimpleRectNode {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGSimpleRectNodeC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGSimpleRectNodeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGSimpleRectNode)
	return gothis
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)
func (this *QSGSimpleRectNode) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGSimpleRectNode7setRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:55
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QSGSimpleRectNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGSimpleRectNode7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:56
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QSGSimpleRectNode) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGSimpleRectNode4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QSGSimpleRectNode) SetColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGSimpleRectNode8setColorERK6QColor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:59
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor color()
func (this *QSGSimpleRectNode) Color() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QSGSimpleRectNode5colorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

func DeleteQSGSimpleRectNode(this *QSGSimpleRectNode) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QSGSimpleRectNodeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
