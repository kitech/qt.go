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
// Public
// void QSGSimpleRectNode(const class QRectF &, const class QColor &)
func NewQSGSimpleRectNode(rect *qtcore.QRectF, color *qtgui.QColor) *QSGSimpleRectNode {
	cthis := qtrt.Calloc(1, 256) // 320
	var convArg0 = rect.GetCthis()
	var convArg1 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGSimpleRectNodeC2ERK6QRectFRK6QColor", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGSimpleRectNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:52
// index:1
// Public
// void QSGSimpleRectNode()
func NewQSGSimpleRectNode_1() *QSGSimpleRectNode {
	cthis := qtrt.Calloc(1, 256) // 320
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGSimpleRectNodeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQSGSimpleRectNodeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:54
// index:0
// Public
// void setRect(const class QRectF &)
func (this *QSGSimpleRectNode) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGSimpleRectNode7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:55
// index:1
// Public inline
// void setRect(qreal, qreal, qreal, qreal)
func (this *QSGSimpleRectNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGSimpleRectNode7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:56
// index:0
// Public
// QRectF rect()
func (this *QSGSimpleRectNode) Rect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSGSimpleRectNode4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:58
// index:0
// Public
// void setColor(const class QColor &)
func (this *QSGSimpleRectNode) SetColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QSGSimpleRectNode8setColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgsimplerectnode.h:59
// index:0
// Public
// QColor color()
func (this *QSGSimpleRectNode) Color() *qtgui.QColor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QSGSimpleRectNode5colorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
