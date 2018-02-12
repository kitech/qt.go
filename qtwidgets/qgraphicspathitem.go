package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsPathItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsPathItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsPathItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsPathItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsPathItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsPathItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsPathItemFromPointer(cthis unsafe.Pointer) *QGraphicsPathItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsPathItem{bcthis0}
}
func (*QGraphicsPathItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsPathItem {
	return NewQGraphicsPathItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:628
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsPathItem(QGraphicsItem *)
func NewQGraphicsPathItem(parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsPathItem {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsPathItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsPathItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsPathItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:629
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsPathItem(const QPainterPath &, QGraphicsItem *)
func NewQGraphicsPathItem_1(path *qtgui.QPainterPath, parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsPathItem {
	var convArg0 = path.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsPathItemC2ERK12QPainterPathP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsPathItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsPathItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:630
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsPathItem()
func DeleteQGraphicsPathItem(this *QGraphicsPathItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsPathItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:632
// index:0
// Public Visibility=Default Availability=Available
// [8] QPainterPath path()
func (this *QGraphicsPathItem) Path() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem4pathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:633
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QPainterPath &)
func (this *QGraphicsPathItem) SetPath(path *qtgui.QPainterPath) {
	var convArg0 = path.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsPathItem7setPathERK12QPainterPath", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:635
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsPathItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:636
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape()
func (this *QGraphicsPathItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:637
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QGraphicsPathItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:639
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsPathItem) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsPathItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:641
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsPathItem) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:642
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsPathItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:645
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsPathItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:648
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsPathItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:649
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsPathItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsPathItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:650
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &)
func (this *QGraphicsPathItem) Extension(variant *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsPathItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsPathItem__ = int

const QGraphicsPathItem__Type QGraphicsPathItem__ = 2

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
		qtgui.KeepMe()
	}
}

//  keep block end
