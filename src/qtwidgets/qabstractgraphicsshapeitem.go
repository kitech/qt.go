//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 214
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
	return this.QGraphicsItem.GetCthis()
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:603
// index:0
// void QAbstractGraphicsShapeItem(class QGraphicsItem *)
func NewQAbstractGraphicsShapeItem(parent unsafe.Pointer) *QAbstractGraphicsShapeItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemC1EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return gothis
}
func NewQAbstractGraphicsShapeItemFromPointer(cthis unsafe.Pointer) *QAbstractGraphicsShapeItem {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QAbstractGraphicsShapeItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:616
// index:1
// void QAbstractGraphicsShapeItem(class QAbstractGraphicsShapeItemPrivate &, class QGraphicsItem *)
func NewQAbstractGraphicsShapeItem_1(dd unsafe.Pointer, parent unsafe.Pointer) *QAbstractGraphicsShapeItem {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemC1ER33QAbstractGraphicsShapeItemPrivateP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, dd, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:604
// index:0
// virtual
// void ~QAbstractGraphicsShapeItem()
func DeleteQAbstractGraphicsShapeItem(*QAbstractGraphicsShapeItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:606
// index:0
// QPen pen()
func (this *QAbstractGraphicsShapeItem) Pen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem3penEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:607
// index:0
// void setPen(const class QPen &)
func (this *QAbstractGraphicsShapeItem) SetPen(pen unsafe.Pointer) {
	// 0: (, pen const QPen &), (pen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItem6setPenERK4QPen", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:609
// index:0
// QBrush brush()
func (this *QAbstractGraphicsShapeItem) Brush() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem5brushEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:610
// index:0
// void setBrush(const class QBrush &)
func (this *QAbstractGraphicsShapeItem) SetBrush(brush unsafe.Pointer) {
	// 0: (, brush const QBrush &), (brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN26QAbstractGraphicsShapeItem8setBrushERK6QBrush", ffiqt.FFI_TYPE_VOID, this.GetCthis(), brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:612
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QAbstractGraphicsShapeItem) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, item const QGraphicsItem *), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:613
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QAbstractGraphicsShapeItem) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK26QAbstractGraphicsShapeItem10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
