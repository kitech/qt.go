//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 17
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
type QGraphicsItemGroup struct {
	*QGraphicsItem
}

func (this *QGraphicsItemGroup) GetCthis() unsafe.Pointer {
	return this.QGraphicsItem.GetCthis()
}
func NewQGraphicsItemGroupFromPointer(cthis unsafe.Pointer) *QGraphicsItemGroup {
	bcthis0 := NewQGraphicsItemFromPointer(cthis)
	return &QGraphicsItemGroup{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1004
// index:0
// Public
// void QGraphicsItemGroup(class QGraphicsItem *)
func NewQGraphicsItemGroup(parent unsafe.Pointer) *QGraphicsItemGroup {
	cthis := qtrt.Calloc(1, 256) // 16
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroupC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsItemGroupFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1005
// index:0
// Public virtual
// void ~QGraphicsItemGroup()
func DeleteQGraphicsItemGroup(*QGraphicsItemGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1007
// index:0
// Public
// void addToGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) AddToGroup(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1008
// index:0
// Public
// void removeFromGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) RemoveFromGroup(item unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1010
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsItemGroup) BoundingRect() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12boundingRectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1011
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsItemGroup) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1013
// index:0
// Public virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItemGroup) IsObscuredBy(item unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), item)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1014
// index:0
// Public virtual
// QPainterPath opaqueArea()
func (this *QGraphicsItemGroup) OpaqueArea() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup10opaqueAreaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1017
// index:0
// Public virtual
// int type()
func (this *QGraphicsItemGroup) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
