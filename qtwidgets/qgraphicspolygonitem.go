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
// extern C begin: 21
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsPolygonItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsPolygonItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsPolygonItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsPolygonItem struct {
	*QAbstractGraphicsShapeItem
}
type QGraphicsPolygonItem_ITF interface {
	QAbstractGraphicsShapeItem_ITF
	QGraphicsPolygonItem_PTR() *QGraphicsPolygonItem
}

func (ptr *QGraphicsPolygonItem) QGraphicsPolygonItem_PTR() *QGraphicsPolygonItem { return ptr }

func (this *QGraphicsPolygonItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsPolygonItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsPolygonItemFromPointer(cthis unsafe.Pointer) *QGraphicsPolygonItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsPolygonItem{bcthis0}
}
func (*QGraphicsPolygonItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsPolygonItem {
	return NewQGraphicsPolygonItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:743
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsPolygonItem(QGraphicsItem *)
func NewQGraphicsPolygonItem(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsPolygonItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsPolygonItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:743
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsPolygonItem(QGraphicsItem *)
func NewQGraphicsPolygonItem__() *QGraphicsPolygonItem {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record,
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsPolygonItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:744
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsPolygonItem(const QPolygonF &, QGraphicsItem *)
func NewQGraphicsPolygonItem_1(polygon qtgui.QPolygonF_ITF, parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsPolygonItem {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg1 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsPolygonItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:744
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsPolygonItem(const QPolygonF &, QGraphicsItem *)
func NewQGraphicsPolygonItem_1_(polygon qtgui.QPolygonF_ITF) *QGraphicsPolygonItem {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	// arg: 1, QGraphicsItem *=Pointer, QGraphicsItem=Record,
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemC2ERK9QPolygonFP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsPolygonItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsPolygonItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:746
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsPolygonItem()
func DeleteQGraphicsPolygonItem(this *QGraphicsPolygonItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:748
// index:0
// Public Visibility=Default Availability=Available
// [8] QPolygonF polygon() const
func (this *QGraphicsPolygonItem) Polygon() *qtgui.QPolygonF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem7polygonEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPolygonFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPolygonF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:749
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPolygon(const QPolygonF &)
func (this *QGraphicsPolygonItem) SetPolygon(polygon qtgui.QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if polygon != nil && polygon.QPolygonF_PTR() != nil {
		convArg0 = polygon.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem10setPolygonERK9QPolygonF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:751
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FillRule fillRule() const
func (this *QGraphicsPolygonItem) FillRule() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem8fillRuleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:752
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFillRule(Qt::FillRule)
func (this *QGraphicsPolygonItem) SetFillRule(rule int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem11setFillRuleEN2Qt8FillRuleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rule)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:754
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const
func (this *QGraphicsPolygonItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:755
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape() const
func (this *QGraphicsPolygonItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:756
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const
func (this *QGraphicsPolygonItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:758
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsPolygonItem) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg2 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:758
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsPolygonItem) Paint__(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if painter != nil && painter.QPainter_PTR() != nil {
		convArg0 = painter.QPainter_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if option != nil && option.QStyleOptionGraphicsItem_PTR() != nil {
		convArg1 = option.QStyleOptionGraphicsItem_PTR().GetCthis()
	}
	// arg: 2, QWidget *=Pointer, QWidget=Record,
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:760
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *) const
func (this *QGraphicsPolygonItem) IsObscuredBy(item QGraphicsItem_ITF /*777 const QGraphicsItem **/) bool {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:761
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea() const
func (this *QGraphicsPolygonItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:764
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const
func (this *QGraphicsPolygonItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:767
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension) const
func (this *QGraphicsPolygonItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:768
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsPolygonItem) SetExtension(extension int, variant qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg1 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsPolygonItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:769
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &) const
func (this *QGraphicsPolygonItem) Extension(variant qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg0 unsafe.Pointer
	if variant != nil && variant.QVariant_PTR() != nil {
		convArg0 = variant.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsPolygonItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsPolygonItem__ = int

const QGraphicsPolygonItem__Type QGraphicsPolygonItem__ = 5

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
		qtgui.KeepMe()
	}
}

//  keep block end
