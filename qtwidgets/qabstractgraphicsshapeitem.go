package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 205
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
}

//  ext block end

//  body block begin
type QAbstractGraphicsShapeItem struct {
	*QGraphicsItem
}

func (this *QAbstractGraphicsShapeItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsItem.GetCthis()
	}
}
func (this *QAbstractGraphicsShapeItem) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsItem = NewQGraphicsItemFromPointer(cthis)
}
func NewQAbstractGraphicsShapeItemFromPointer(cthis unsafe.Pointer) *QAbstractGraphicsShapeItem {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QAbstractGraphicsShapeItem{bcthis0}
}
func (*QAbstractGraphicsShapeItem) NewFromPointer(cthis unsafe.Pointer) *QAbstractGraphicsShapeItem {
	return NewQAbstractGraphicsShapeItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:603
// index:0
// Public
// void QAbstractGraphicsShapeItem(class QGraphicsItem *)
func NewQAbstractGraphicsShapeItem(parent *QGraphicsItem /*777 QGraphicsItem **/) *QAbstractGraphicsShapeItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemC1EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:604
// index:0
// Public virtual
// void ~QAbstractGraphicsShapeItem()
func DeleteQAbstractGraphicsShapeItem(*QAbstractGraphicsShapeItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:606
// index:0
// Public
// QPen pen()
func (this *QAbstractGraphicsShapeItem) Pen() *qtgui.QPen /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem3penEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPenFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:607
// index:0
// Public
// void setPen(const class QPen &)
func (this *QAbstractGraphicsShapeItem) SetPen(pen *qtgui.QPen) {
	var convArg0 = pen.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:609
// index:0
// Public
// QBrush brush()
func (this *QAbstractGraphicsShapeItem) Brush() *qtgui.QBrush /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem5brushEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQBrushFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:610
// index:0
// Public
// void setBrush(const class QBrush &)
func (this *QAbstractGraphicsShapeItem) SetBrush(brush *qtgui.QBrush) {
	var convArg0 = brush.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:612
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QAbstractGraphicsShapeItem) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:613
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QAbstractGraphicsShapeItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
