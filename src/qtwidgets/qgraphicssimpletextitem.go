//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 55
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
type QGraphicsSimpleTextItem struct {
	*QAbstractGraphicsShapeItem
}

func (this *QGraphicsSimpleTextItem) GetCthis() unsafe.Pointer {
	return this.QAbstractGraphicsShapeItem.GetCthis()
}
func NewQGraphicsSimpleTextItemFromPointer(cthis unsafe.Pointer) *QGraphicsSimpleTextItem {
	bcthis0 := NewQAbstractGraphicsShapeItemFromPointer(cthis)
	return &QGraphicsSimpleTextItem{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:968
// index:0
// Public
// void QGraphicsSimpleTextItem(class QGraphicsItem *)
func NewQGraphicsSimpleTextItem(parent unsafe.Pointer) *QGraphicsSimpleTextItem {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSimpleTextItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:969
// index:1
// Public
// void QGraphicsSimpleTextItem(const class QString &, class QGraphicsItem *)
func NewQGraphicsSimpleTextItem_1(text *qtcore.QString, parent unsafe.Pointer) *QGraphicsSimpleTextItem {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemC2ERK7QStringP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, convArg0, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsSimpleTextItemFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:970
// index:0
// Public virtual
// void ~QGraphicsSimpleTextItem()
func DeleteQGraphicsSimpleTextItem(*QGraphicsSimpleTextItem) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItemD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:972
// index:0
// Public
// void setText(const class QString &)
func (this *QGraphicsSimpleTextItem) SetText(text *qtcore.QString) {
	var convArg0 = text.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem7setTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:973
// index:0
// Public
// QString text()
func (this *QGraphicsSimpleTextItem) Text() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4textEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:975
// index:0
// Public
// void setFont(const class QFont &)
func (this *QGraphicsSimpleTextItem) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:976
// index:0
// Public
// QFont font()
func (this *QGraphicsSimpleTextItem) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:978
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsSimpleTextItem) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:979
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsSimpleTextItem) Shape() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem5shapeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:980
// index:0
// Public virtual
// bool contains(const class QPointF &)
func (this *QGraphicsSimpleTextItem) Contains(point *qtcore.QPointF) interface{} {
	var convArg0 = point.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem8containsERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:982
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsSimpleTextItem) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:984
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsSimpleTextItem) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:985
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsSimpleTextItem) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:988
// index:0
// Public virtual
// int type()
func (this *QGraphicsSimpleTextItem) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:991
// index:0
// Protected virtual
// bool supportsExtension(enum QGraphicsItem::Extension)
func (this *QGraphicsSimpleTextItem) SupportsExtension(extension int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem17supportsExtensionEN13QGraphicsItem9ExtensionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:992
// index:0
// Protected virtual
// void setExtension(enum QGraphicsItem::Extension, const class QVariant &)
func (this *QGraphicsSimpleTextItem) SetExtension(extension int, variant *qtcore.QVariant) {
	var convArg1 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN23QGraphicsSimpleTextItem12setExtensionEN13QGraphicsItem9ExtensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &extension, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:993
// index:0
// Protected virtual
// QVariant extension(const class QVariant &)
func (this *QGraphicsSimpleTextItem) Extension(variant *qtcore.QVariant) interface{} {
	var convArg0 = variant.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK23QGraphicsSimpleTextItem9extensionERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
