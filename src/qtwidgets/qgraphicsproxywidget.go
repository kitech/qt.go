//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h
// #include <qgraphicsproxywidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 56
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
type QGraphicsProxyWidget struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QGraphicsProxyWidget) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:57
// index:0
// virtual
// void ~QGraphicsProxyWidget()
func DeleteQGraphicsProxyWidget(*QGraphicsProxyWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:59
// index:0
// void setWidget(class QWidget *)
func (this *QGraphicsProxyWidget) SetWidget(widget unsafe.Pointer) {
	// 0: (, QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9setWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:60
// index:0
// QWidget * widget()
func (this *QGraphicsProxyWidget) Widget() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget6widgetEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:62
// index:0
// QRectF subWidgetRect(const class QWidget *)
func (this *QGraphicsProxyWidget) SubWidgetRect(widget unsafe.Pointer) {
	// 0: (, const QWidget * widget), (widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget13subWidgetRectEPK7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:64
// index:0
// virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsProxyWidget) SetGeometry(rect unsafe.Pointer) {
	// 0: (, const QRectF & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11setGeometryERK6QRectF", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:66
// index:0
// virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsProxyWidget) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	// 0: (, QPainter * painter, const QStyleOptionGraphicsItem * option, QWidget * widget), (painter, option, widget)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:71
// index:0
// virtual
// int type()
func (this *QGraphicsProxyWidget) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:73
// index:0
// QGraphicsProxyWidget * createProxyForChildWidget(class QWidget *)
func (this *QGraphicsProxyWidget) CreateProxyForChildWidget(child unsafe.Pointer) {
	// 0: (, QWidget * child), (child)
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget25createProxyForChildWidgetEP7QWidget", ffiqt.FFI_TYPE_VOID, this.cthis, child)
	gopp.ErrPrint(err, rv)
}

//  body block end
