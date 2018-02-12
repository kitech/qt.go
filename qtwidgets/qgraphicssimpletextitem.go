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
// extern C begin: 56
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
func (this *QGraphicsSimpleTextItem) InheritSupportsExtension(f func(extension int) bool) {
	qtrt.SetAllInheritCallback(this, "supportsExtension", f)
}

// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsSimpleTextItem) InheritSetExtension(f func(extension int, variant *qtcore.QVariant) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setExtension", f)
}

// QVariant extension(const class QVariant &)
func (this *QGraphicsSimpleTextItem) InheritExtension(f func(variant *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "extension", f)
}

type QGraphicsSimpleTextItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsSimpleTextItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QAbstractGraphicsShapeItem.GetCthis()
	}
}
func (this *QGraphicsSimpleTextItem) SetCthis(cthis unsafe.Pointer) {
	this.QAbstractGraphicsShapeItem = NewQAbstractGraphicsShapeItemFromPointer(cthis)
}
func NewQGraphicsSimpleTextItemFromPointer(cthis unsafe.Pointer) *QGraphicsSimpleTextItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsSimpleTextItem{bcthis0}
}
func (*QGraphicsSimpleTextItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsSimpleTextItem {
	return NewQGraphicsSimpleTextItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:968
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSimpleTextItem(QGraphicsItem *)
func NewQGraphicsSimpleTextItem(parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsSimpleTextItem {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSimpleTextItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:969
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsSimpleTextItem(const QString &, QGraphicsItem *)
func NewQGraphicsSimpleTextItem_1(text string, parent *QGraphicsItem /*777 QGraphicsItem **/) *QGraphicsSimpleTextItem {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsSimpleTextItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsSimpleTextItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:970
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsSimpleTextItem()
func DeleteQGraphicsSimpleTextItem(this *QGraphicsSimpleTextItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:972
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setText(const QString &)
func (this *QGraphicsSimpleTextItem) SetText(text string) {
	var tmpArg0 = qtcore.NewQString_5(text)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem7setTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:973
// index:0
// Public Visibility=Default Availability=Available
// [8] QString text()
func (this *QGraphicsSimpleTextItem) Text() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4textEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:975
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QGraphicsSimpleTextItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:976
// index:0
// Public Visibility=Default Availability=Available
// [16] QFont font()
func (this *QGraphicsSimpleTextItem) Font() *qtgui.QFont /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:978
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QGraphicsSimpleTextItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:979
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath shape()
func (this *QGraphicsSimpleTextItem) Shape() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem5shapeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:980
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QGraphicsSimpleTextItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:982
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)
func (this *QGraphicsSimpleTextItem) Paint(painter *qtgui.QPainter /*777 QPainter **/, option *QStyleOptionGraphicsItem /*777 const QStyleOptionGraphicsItem **/, widget *QWidget /*777 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:984
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isObscuredBy(const QGraphicsItem *)
func (this *QGraphicsSimpleTextItem) IsObscuredBy(item *QGraphicsItem /*777 const QGraphicsItem **/) bool {
	var convArg0 = item.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:985
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPainterPath opaqueArea()
func (this *QGraphicsSimpleTextItem) OpaqueArea() *qtgui.QPainterPath /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPainterPath)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:988
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type()
func (this *QGraphicsSimpleTextItem) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:991
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsSimpleTextItem) SupportsExtension(extension int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:992
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void setExtension(enum QGraphicsItem::Extension, const QVariant &)
func (this *QGraphicsSimpleTextItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extension, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:993
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant extension(const QVariant &)
func (this *QGraphicsSimpleTextItem) Extension(variant *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	var convArg0 = variant.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

type QGraphicsSimpleTextItem__ = int

const QGraphicsSimpleTextItem__Type QGraphicsSimpleTextItem__ = 9

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
