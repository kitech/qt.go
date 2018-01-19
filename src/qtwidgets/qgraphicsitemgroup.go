//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsitem.h
// #include <qgraphicsitem.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1004
// index:0
// void QGraphicsItemGroup(class QGraphicsItem *)
func NewQGraphicsItemGroup(parent unsafe.Pointer) *QGraphicsItemGroup {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroupC2EP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QGraphicsItemGroup{cthis}
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1005
// index:0
// virtual
// void ~QGraphicsItemGroup()
func DeleteQGraphicsItemGroup(*QGraphicsItemGroup) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroupD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1007
// index:0
// void addToGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) AddToGroup(item unsafe.Pointer) {
	// 0: (, QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup10addToGroupEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1008
// index:0
// void removeFromGroup(class QGraphicsItem *)
func (this *QGraphicsItemGroup) RemoveFromGroup(item unsafe.Pointer) {
	// 0: (, QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup15removeFromGroupEP13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1010
// index:0
// virtual
// QRectF boundingRect()
func (this *QGraphicsItemGroup) BoundingRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12boundingRectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1011
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsItemGroup) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN18QGraphicsItemGroup5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1013
// index:0
// virtual
// bool isObscuredBy(const class QGraphicsItem *)
func (this *QGraphicsItemGroup) IsObscuredBy(item unsafe.Pointer) {
	// 0: (, const QGraphicsItem * item), (item)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup12isObscuredByEPK13QGraphicsItem", ffiqt.FFI_TYPE_VOID, this.cthis, item)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1014
// index:0
// virtual
// QPainterPath opaqueArea()
func (this *QGraphicsItemGroup) OpaqueArea() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup10opaqueAreaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsitem.h:1017
// index:0
// virtual
// int type()
func (this *QGraphicsItemGroup) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK18QGraphicsItemGroup4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
