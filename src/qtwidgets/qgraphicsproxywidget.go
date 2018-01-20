//  header block begin
// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h
// #include <qgraphicsproxywidget.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 81
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
	*QGraphicsWidget
}

func (this *QGraphicsProxyWidget) GetCthis() unsafe.Pointer {
	return this.QGraphicsWidget.GetCthis()
}
func NewQGraphicsProxyWidgetFromPointer(cthis unsafe.Pointer) *QGraphicsProxyWidget {
	bcthis0 := NewQGraphicsWidgetFromPointer(cthis)
	return &QGraphicsProxyWidget{bcthis0}
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:54
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsProxyWidget) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:57
// index:0
// Public virtual
// void ~QGraphicsProxyWidget()
func DeleteQGraphicsProxyWidget(*QGraphicsProxyWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:59
// index:0
// Public
// void setWidget(class QWidget *)
func (this *QGraphicsProxyWidget) SetWidget(widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9setWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:60
// index:0
// Public
// QWidget * widget()
func (this *QGraphicsProxyWidget) Widget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget6widgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:62
// index:0
// Public
// QRectF subWidgetRect(const class QWidget *)
func (this *QGraphicsProxyWidget) SubWidgetRect(widget unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget13subWidgetRectEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), widget)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:64
// index:0
// Public virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsProxyWidget) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11setGeometryERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:66
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsProxyWidget) Paint(painter unsafe.Pointer, option unsafe.Pointer, widget unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), painter, option, widget)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:71
// index:0
// Public virtual
// int type()
func (this *QGraphicsProxyWidget) Type() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:73
// index:0
// Public
// QGraphicsProxyWidget * createProxyForChildWidget(class QWidget *)
func (this *QGraphicsProxyWidget) CreateProxyForChildWidget(child unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget25createProxyForChildWidgetEP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), child)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:76
// index:0
// Protected virtual
// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
func (this *QGraphicsProxyWidget) ItemChange(change int, value *qtcore.QVariant) interface{} {
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &change, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:78
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QGraphicsProxyWidget) Event(event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:79
// index:0
// Protected virtual
// bool eventFilter(class QObject *, class QEvent *)
func (this *QGraphicsProxyWidget) EventFilter(object unsafe.Pointer, event unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), object, event)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:81
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QGraphicsProxyWidget) ShowEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:82
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QGraphicsProxyWidget) HideEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:85
// index:0
// Protected virtual
// void contextMenuEvent(class QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsProxyWidget) ContextMenuEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:89
// index:0
// Protected virtual
// void dragEnterEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) DragEnterEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:90
// index:0
// Protected virtual
// void dragLeaveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) DragLeaveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:91
// index:0
// Protected virtual
// void dragMoveEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) DragMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:92
// index:0
// Protected virtual
// void dropEvent(class QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) DropEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:95
// index:0
// Protected virtual
// void hoverEnterEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsProxyWidget) HoverEnterEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:96
// index:0
// Protected virtual
// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsProxyWidget) HoverLeaveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:97
// index:0
// Protected virtual
// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsProxyWidget) HoverMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:98
// index:0
// Protected virtual
// void grabMouseEvent(class QEvent *)
func (this *QGraphicsProxyWidget) GrabMouseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:99
// index:0
// Protected virtual
// void ungrabMouseEvent(class QEvent *)
func (this *QGraphicsProxyWidget) UngrabMouseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:101
// index:0
// Protected virtual
// void mouseMoveEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) MouseMoveEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:102
// index:0
// Protected virtual
// void mousePressEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) MousePressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:103
// index:0
// Protected virtual
// void mouseReleaseEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) MouseReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:104
// index:0
// Protected virtual
// void mouseDoubleClickEvent(class QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) MouseDoubleClickEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:106
// index:0
// Protected virtual
// void wheelEvent(class QGraphicsSceneWheelEvent *)
func (this *QGraphicsProxyWidget) WheelEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:109
// index:0
// Protected virtual
// void keyPressEvent(class QKeyEvent *)
func (this *QGraphicsProxyWidget) KeyPressEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:110
// index:0
// Protected virtual
// void keyReleaseEvent(class QKeyEvent *)
func (this *QGraphicsProxyWidget) KeyReleaseEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:112
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsProxyWidget) FocusInEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:113
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsProxyWidget) FocusOutEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:114
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QGraphicsProxyWidget) FocusNextPrevChild(next bool) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &next)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:116
// index:0
// Protected virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsProxyWidget) InputMethodQuery(query int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &query)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:117
// index:0
// Protected virtual
// void inputMethodEvent(class QInputMethodEvent *)
func (this *QGraphicsProxyWidget) InputMethodEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:119
// index:0
// Protected virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsProxyWidget) SizeHint(which int, constraint *qtcore.QSizeF) interface{} {
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &which, convArg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:120
// index:0
// Protected virtual
// void resizeEvent(class QGraphicsSceneResizeEvent *)
func (this *QGraphicsProxyWidget) ResizeEvent(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), event)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:123
// index:0
// Protected
// QGraphicsProxyWidget * newProxyWidget(const class QWidget *)
func (this *QGraphicsProxyWidget) NewProxyWidget(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
