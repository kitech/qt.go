//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QGraphicsPathItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsPathItem) GetCthis() unsafe.Pointer {
	return this.QAbstractGraphicsShapeItem.GetCthis()
}
func NewQGraphicsPathItemFromPointer(cthis unsafe.Pointer) *QGraphicsPathItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsPathItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:628
// index:0
// Public
// void QGraphicsPathItem(class QGraphicsItem *)
func NewQGraphicsPathItem(parent unsafe.Pointer) *QGraphicsPathItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsPathItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsPathItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:629
// index:1
// Public
// void QGraphicsPathItem(const class QPainterPath &, class QGraphicsItem *)
func NewQGraphicsPathItem_1(path *qtgui.QPainterPath, parent unsafe.Pointer) *QGraphicsPathItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsPathItemC2ERK12QPainterPathP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsPathItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:630
// index:0
// Public virtual
// void ~QGraphicsPathItem()
func DeleteQGraphicsPathItem(*QGraphicsPathItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsPathItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:632
// index:0
// Public
// QPainterPath path()
func (this *QGraphicsPathItem) Path() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem4pathEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:633
// index:0
// Public
// void setPath(const class QPainterPath &)
func (this *QGraphicsPathItem) SetPath(path *qtgui.QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsPathItem7setPathERK12QPainterPath", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:635
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsPathItem) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:636
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsPathItem) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:637
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsPathItem) Contains(point *qtcore.QPointF) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:639
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsPathItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:641
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsPathItem) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:642
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsPathItem) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:645
// index:0
// Public virtual
// int type()
func (this *QGraphicsPathItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:648
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsPathItem) SupportsExtension(extension int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:649
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsPathItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:650
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsPathItem) Extension(variant *qtcore.QVariant) interface{} {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QGraphicsPathItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
