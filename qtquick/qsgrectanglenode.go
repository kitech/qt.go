package qtquick

// /usr/include/qt/QtQuick/qsgrectanglenode.h
// #include <qsgrectanglenode.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
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
type QSGRectangleNode struct {
	*QSGGeometryNode
}

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
func DeleteQSGRectangleNode(*QSGRectangleNode) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGRectangleNodeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:52
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)
func (this *QSGRectangleNode) SetRect(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGRectangleNode7setRectERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:53
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QSGRectangleNode) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGRectangleNode7setRectEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:54
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QSGRectangleNode) Rect() *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSGRectangleNode4rectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:56
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setColor(const QColor &)
func (this *QSGRectangleNode) SetColor(color *qtgui.QColor) {
	var convArg0 = color.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN16QSGRectangleNode8setColorERK6QColor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgrectanglenode.h:57
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QColor color()
func (this *QSGRectangleNode) Color() *qtgui.QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK16QSGRectangleNode5colorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
