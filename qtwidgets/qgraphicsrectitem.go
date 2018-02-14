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
// extern C begin: 15
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
func (this *QGraphicsRectItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsRectItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsRectItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsRectItem struct {
	*QAbstractGraphicsShapeItem
}
type QGraphicsRectItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsRectItem_PTR() *QGraphicsRectItem
}

func (ptr *QGraphicsRectItem) QGraphicsRectItem_PTR() *QGraphicsRectItem { return ptr }

func (this *QGraphicsRectItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsRectItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsRectItemFromPointer(cthis unsafe.Pointer) *QGraphicsRectItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsRectItem{bcthis0}
}
func (*QGraphicsRectItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsRectItem {
	return NewQGraphicsRectItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:661
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(QGraphicsItem *)
func NewQGraphicsRectItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	var convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:662
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(const QRectF &, QGraphicsItem *)
func NewQGraphicsRectItem_1(rect qtcore.QRectF_ITF, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	var convArg1 = parent.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2ERK6QRectFP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:663
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsRectItem(qreal, qreal, qreal, qreal, QGraphicsItem *)
func NewQGraphicsRectItem_2(x float64, y float64, w float64, h float64, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsRectItem {
	var convArg4 = parent.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemC2EddddP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, x, y, w, h, convArg4)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsRectItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsRectItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:664
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsRectItem()
func DeleteQGraphicsRectItem(this *QGraphicsRectItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:666
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF rect()
func (this *QGraphicsRectItem) Rect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:667
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRect(const QRectF &)
func (this *QGraphicsRectItem) SetRect(rect qtcore.QRectF_ITF) {
	var convArg0 = rect.QRectF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:668
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setRect(qreal, qreal, qreal, qreal)
func (this *QGraphicsRectItem) SetRect_1(x float64, y float64, w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem7setRectEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:670
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsRectItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:671
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape()
func (this *QGraphicsRectItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:672
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QGraphicsRectItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 = point.QPointF_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:674
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsRectItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 = painter.QPainter_PTR().GetCthis()
	var convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	var convArg2 = widget.QWidget_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:676
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsRectItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.QGraphicsItem_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:677
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsRectItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:680
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsRectItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:683
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsRectItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:684
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsRectItem) SetExtension(extension int, variant qtcore.QVariant_ITF) {
	var convArg1 = variant.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QGraphicsRectItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:685
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &)
func (this *QGraphicsRectItem) Extension(variant qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.QVariant_PTR().GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QGraphicsRectItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsRectItem__ = int

const QGraphicsRectItem__Type QGraphicsRectItem__ = 3

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
