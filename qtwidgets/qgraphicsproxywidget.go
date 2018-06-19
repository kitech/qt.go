package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h
// #include <qgraphicsproxywidget.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 82
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

// QVariant itemChange(QGraphicsItem::GraphicsItemChange, const QVariant &)
func (this *QGraphicsProxyWidget) InheritItemChange(f func(change int, value *qtcore.QVariant) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "itemChange", f)
}

// bool event(QEvent *)
func (this *QGraphicsProxyWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool eventFilter(QObject *, QEvent *)
func (this *QGraphicsProxyWidget) InheritEventFilter(f func(object *qtcore.QObject /*777 QObject **/, event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "eventFilter", f)
}

// void showEvent(QShowEvent *)
func (this *QGraphicsProxyWidget) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(QHideEvent *)
func (this *QGraphicsProxyWidget) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// void contextMenuEvent(QGraphicsSceneContextMenuEvent *)
func (this *QGraphicsProxyWidget) InheritContextMenuEvent(f func(event *QGraphicsSceneContextMenuEvent /*777 QGraphicsSceneContextMenuEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void dragEnterEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) InheritDragEnterEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragLeaveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) InheritDragLeaveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dragMoveEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) InheritDragMoveEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dropEvent(QGraphicsSceneDragDropEvent *)
func (this *QGraphicsProxyWidget) InheritDropEvent(f func(event *QGraphicsSceneDragDropEvent /*777 QGraphicsSceneDragDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void hoverEnterEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsProxyWidget) InheritHoverEnterEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverLeaveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsProxyWidget) InheritHoverLeaveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void hoverMoveEvent(QGraphicsSceneHoverEvent *)
func (this *QGraphicsProxyWidget) InheritHoverMoveEvent(f func(event *QGraphicsSceneHoverEvent /*777 QGraphicsSceneHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void grabMouseEvent(QEvent *)
func (this *QGraphicsProxyWidget) InheritGrabMouseEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "grabMouseEvent", f)
}

// void ungrabMouseEvent(QEvent *)
func (this *QGraphicsProxyWidget) InheritUngrabMouseEvent(f func(event *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "ungrabMouseEvent", f)
}

// void mouseMoveEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) InheritMouseMoveEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mousePressEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) InheritMousePressEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) InheritMouseReleaseEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)
func (this *QGraphicsProxyWidget) InheritMouseDoubleClickEvent(f func(event *QGraphicsSceneMouseEvent /*777 QGraphicsSceneMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void wheelEvent(QGraphicsSceneWheelEvent *)
func (this *QGraphicsProxyWidget) InheritWheelEvent(f func(event *QGraphicsSceneWheelEvent /*777 QGraphicsSceneWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QGraphicsProxyWidget) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QGraphicsProxyWidget) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QGraphicsProxyWidget) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QGraphicsProxyWidget) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// bool focusNextPrevChild(bool)
func (this *QGraphicsProxyWidget) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QGraphicsProxyWidget) InheritInputMethodQuery(f func(query int) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "inputMethodQuery", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QGraphicsProxyWidget) InheritInputMethodEvent(f func(event *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsProxyWidget) InheritSizeHint(f func(which int, constraint *qtcore.QSizeF) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

// void resizeEvent(QGraphicsSceneResizeEvent *)
func (this *QGraphicsProxyWidget) InheritResizeEvent(f func(event *QGraphicsSceneResizeEvent /*777 QGraphicsSceneResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// QGraphicsProxyWidget * newProxyWidget(const QWidget *)
func (this *QGraphicsProxyWidget) InheritNewProxyWidget(f func(arg0 *QWidget /*777 const QWidget **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "newProxyWidget", f)
}

/*

 */
type QGraphicsProxyWidget struct {
	*QGraphicsWidget
}
type QGraphicsProxyWidget_ITF interface {
	QGraphicsWidget_ITF
	QGraphicsProxyWidget_PTR() *QGraphicsProxyWidget
}

func (ptr *QGraphicsProxyWidget) QGraphicsProxyWidget_PTR() *QGraphicsProxyWidget { return ptr }

func (this *QGraphicsProxyWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsWidget.GetCthis()
	}
}
func (this *QGraphicsProxyWidget) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsWidget = NewQGraphicsWidgetFromPointer(cthis)
}
func NewQGraphicsProxyWidgetFromPointer(cthis unsafe.Pointer) *QGraphicsProxyWidget {
	bcthis0 := NewQGraphicsWidgetFromPointer(cthis)
	return &QGraphicsProxyWidget{bcthis0}
}
func (*QGraphicsProxyWidget) NewFromPointer(cthis unsafe.Pointer) *QGraphicsProxyWidget {
	return NewQGraphicsProxyWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGraphicsProxyWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsProxyWidget(QGraphicsItem *, Qt::WindowFlags)

/*
Constructs a new QGraphicsProxy widget. parent and wFlags are passed to QGraphicsItem's constructor.
*/
func NewQGraphicsProxyWidget(parent QGraphicsItem_ITF /*777 QGraphicsItem **/, wFlags int) *QGraphicsProxyWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsProxyWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsProxyWidget(QGraphicsItem *, Qt::WindowFlags)

/*
Constructs a new QGraphicsProxy widget. parent and wFlags are passed to QGraphicsItem's constructor.
*/
func NewQGraphicsProxyWidget__() *QGraphicsProxyWidget {
	// arg: 0, QGraphicsItem *=Pointer, QGraphicsItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	wFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsProxyWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsProxyWidget(QGraphicsItem *, Qt::WindowFlags)

/*
Constructs a new QGraphicsProxy widget. parent and wFlags are passed to QGraphicsItem's constructor.
*/
func NewQGraphicsProxyWidget__1(parent QGraphicsItem_ITF /*777 QGraphicsItem **/) *QGraphicsProxyWidget {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsItem_PTR() != nil {
		convArg0 = parent.QGraphicsItem_PTR().GetCthis()
	}
	// arg: 1, Qt::WindowFlags=Elaborated, Qt::WindowFlags=Typedef, QFlags<Qt::WindowType>, Unexposed
	wFlags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, wFlags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGraphicsProxyWidget")
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsProxyWidget()

/*

 */
func DeleteQGraphicsProxyWidget(this *QGraphicsProxyWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:59
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidget(QWidget *)

/*
Embeds widget into this proxy widget. The embedded widget must reside exclusively either inside or outside of Graphics View. You cannot embed a widget as long as it is is visible elsewhere in the UI, at the same time.

widget must be a top-level widget whose parent is 0.

When the widget is embedded, its state (e.g., visible, enabled, geometry, size hints) is copied into the proxy widget. If the embedded widget is explicitly hidden or disabled, the proxy widget will become explicitly hidden or disabled after embedding is complete. The class documentation has a full overview over the shared state.

QGraphicsProxyWidget's window flags determine whether the widget, after embedding, will be given window decorations or not.

After this function returns, QGraphicsProxyWidget will keep its state synchronized with that of widget whenever possible.

If a widget is already embedded by this proxy when this function is called, that widget will first be automatically unembedded. Passing 0 for the widget argument will only unembed the widget, and the ownership of the currently embedded widget will be passed on to the caller. Every child widget that are embedded will also be embedded and their proxy widget destroyed.

Note that widgets with the Qt::WA_PaintOnScreen widget attribute set and widgets that wrap an external application or controller cannot be embedded. Examples are QGLWidget and QAxWidget.

See also widget().
*/
func (this *QGraphicsProxyWidget) SetWidget(widget QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9setWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:60
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * widget() const

/*
Returns a pointer to the embedded widget.

See also setWidget().
*/
func (this *QGraphicsProxyWidget) Widget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget6widgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:62
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF subWidgetRect(const QWidget *) const

/*
Returns the rectangle for widget, which must be a descendant of widget(), or widget() itself, in this proxy item's local coordinates.

If no widget is embedded, widget is 0, or widget is not a descendant of the embedded widget, this function returns an empty QRectF.

See also widget().
*/
func (this *QGraphicsProxyWidget) SubWidgetRect(widget QWidget_ITF /*777 const QWidget **/) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg0 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget13subWidgetRectEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)

/*
Reimplemented from QGraphicsLayoutItem::setGeometry().
*/
func (this *QGraphicsProxyWidget) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QStyleOptionGraphicsItem *, QWidget *)

/*
Reimplemented from QGraphicsItem::paint().
*/
func (this *QGraphicsProxyWidget) Paint(painter qtgui.QPainter_ITF /*777 QPainter **/, option QStyleOptionGraphicsItem_ITF /*777 const QStyleOptionGraphicsItem **/, widget QWidget_ITF /*777 QWidget **/) {
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
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int type() const

/*
Reimplemented from QGraphicsItem::type().
*/
func (this *QGraphicsProxyWidget) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * createProxyForChildWidget(QWidget *)

/*
Creates a proxy widget for the given child of the widget contained in this proxy.

This function makes it possible to acquire proxies for non top-level widgets. For instance, you can embed a dialog, and then transform only one of its widgets.

If the widget is already embedded, return the existing proxy widget.

This function was introduced in  Qt 4.5.

See also newProxyWidget() and QGraphicsScene::addWidget().
*/
func (this *QGraphicsProxyWidget) CreateProxyForChildWidget(child QWidget_ITF /*777 QWidget **/) *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	var convArg0 unsafe.Pointer
	if child != nil && child.QWidget_PTR() != nil {
		convArg0 = child.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget25createProxyForChildWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:76
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant itemChange(QGraphicsItem::GraphicsItemChange, const QVariant &)

/*
Reimplemented from QGraphicsItem::itemChange().
*/
func (this *QGraphicsProxyWidget) ItemChange(change int, value qtcore.QVariant_ITF) *qtcore.QVariant /*123*/ {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), change, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:78
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QGraphicsProxyWidget) Event(event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:79
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool eventFilter(QObject *, QEvent *)

/*
Reimplemented from QObject::eventFilter().
*/
func (this *QGraphicsProxyWidget) EventFilter(object qtcore.QObject_ITF /*777 QObject **/, event qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if object != nil && object.QObject_PTR() != nil {
		convArg0 = object.QObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg1 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11eventFilterEP7QObjectP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:81
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)

/*
Reimplemented from QGraphicsWidget::showEvent().
*/
func (this *QGraphicsProxyWidget) ShowEvent(event qtgui.QShowEvent_ITF /*777 QShowEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QShowEvent_PTR() != nil {
		convArg0 = event.QShowEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:82
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)

/*
Reimplemented from QGraphicsWidget::hideEvent().
*/
func (this *QGraphicsProxyWidget) HideEvent(event qtgui.QHideEvent_ITF /*777 QHideEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHideEvent_PTR() != nil {
		convArg0 = event.QHideEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:85
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QGraphicsSceneContextMenuEvent *)

/*
Reimplemented from QGraphicsItem::contextMenuEvent().
*/
func (this *QGraphicsProxyWidget) ContextMenuEvent(event QGraphicsSceneContextMenuEvent_ITF /*777 QGraphicsSceneContextMenuEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneContextMenuEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneContextMenuEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget16contextMenuEventEP30QGraphicsSceneContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:89
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dragEnterEvent().
*/
func (this *QGraphicsProxyWidget) DragEnterEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14dragEnterEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:90
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dragLeaveEvent().
*/
func (this *QGraphicsProxyWidget) DragLeaveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14dragLeaveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:91
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dragMoveEvent().
*/
func (this *QGraphicsProxyWidget) DragMoveEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13dragMoveEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:92
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QGraphicsSceneDragDropEvent *)

/*
Reimplemented from QGraphicsItem::dropEvent().
*/
func (this *QGraphicsProxyWidget) DropEvent(event QGraphicsSceneDragDropEvent_ITF /*777 QGraphicsSceneDragDropEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneDragDropEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneDragDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget9dropEventEP27QGraphicsSceneDragDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverEnterEvent(QGraphicsSceneHoverEvent *)

/*
Reimplemented from QGraphicsItem::hoverEnterEvent().
*/
func (this *QGraphicsProxyWidget) HoverEnterEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15hoverEnterEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QGraphicsSceneHoverEvent *)

/*
Reimplemented from QGraphicsItem::hoverLeaveEvent().
*/
func (this *QGraphicsProxyWidget) HoverLeaveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QGraphicsSceneHoverEvent *)

/*
Reimplemented from QGraphicsItem::hoverMoveEvent().
*/
func (this *QGraphicsProxyWidget) HoverMoveEvent(event QGraphicsSceneHoverEvent_ITF /*777 QGraphicsSceneHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneHoverEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:98
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void grabMouseEvent(QEvent *)

/*
Reimplemented from QGraphicsWidget::grabMouseEvent().
*/
func (this *QGraphicsProxyWidget) GrabMouseEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14grabMouseEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:99
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ungrabMouseEvent(QEvent *)

/*
Reimplemented from QGraphicsWidget::ungrabMouseEvent().
*/
func (this *QGraphicsProxyWidget) UngrabMouseEvent(event qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QEvent_PTR() != nil {
		convArg0 = event.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget16ungrabMouseEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:101
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QGraphicsSceneMouseEvent *)

/*
Reimplemented from QGraphicsItem::mouseMoveEvent().
*/
func (this *QGraphicsProxyWidget) MouseMoveEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14mouseMoveEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:102
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QGraphicsSceneMouseEvent *)

/*
Reimplemented from QGraphicsItem::mousePressEvent().
*/
func (this *QGraphicsProxyWidget) MousePressEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15mousePressEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:103
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QGraphicsSceneMouseEvent *)

/*
Reimplemented from QGraphicsItem::mouseReleaseEvent().
*/
func (this *QGraphicsProxyWidget) MouseReleaseEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget17mouseReleaseEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:104
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QGraphicsSceneMouseEvent *)

/*
Reimplemented from QGraphicsItem::mouseDoubleClickEvent().
*/
func (this *QGraphicsProxyWidget) MouseDoubleClickEvent(event QGraphicsSceneMouseEvent_ITF /*777 QGraphicsSceneMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneMouseEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget21mouseDoubleClickEventEP24QGraphicsSceneMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:106
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QGraphicsSceneWheelEvent *)

/*
Reimplemented from QGraphicsItem::wheelEvent().
*/
func (this *QGraphicsProxyWidget) WheelEvent(event QGraphicsSceneWheelEvent_ITF /*777 QGraphicsSceneWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneWheelEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget10wheelEventEP24QGraphicsSceneWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:109
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
Reimplemented from QGraphicsItem::keyPressEvent().
*/
func (this *QGraphicsProxyWidget) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:110
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
Reimplemented from QGraphicsItem::keyReleaseEvent().
*/
func (this *QGraphicsProxyWidget) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:112
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
Reimplemented from QGraphicsItem::focusInEvent().
*/
func (this *QGraphicsProxyWidget) FocusInEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:113
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
Reimplemented from QGraphicsItem::focusOutEvent().
*/
func (this *QGraphicsProxyWidget) FocusOutEvent(event qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QFocusEvent_PTR() != nil {
		convArg0 = event.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:114
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(bool)

/*
Reimplemented from QGraphicsWidget::focusNextPrevChild().
*/
func (this *QGraphicsProxyWidget) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
Reimplemented from QGraphicsItem::inputMethodQuery().
*/
func (this *QGraphicsProxyWidget) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:117
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
Reimplemented from QGraphicsItem::inputMethodEvent().
*/
func (this *QGraphicsProxyWidget) InputMethodEvent(event qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QInputMethodEvent_PTR() != nil {
		convArg0 = event.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsProxyWidget) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:119
// index:0
// Protected virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsProxyWidget) SizeHint__(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QGraphicsProxyWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:120
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QGraphicsSceneResizeEvent *)

/*
Reimplemented from QGraphicsWidget::resizeEvent().
*/
func (this *QGraphicsProxyWidget) ResizeEvent(event QGraphicsSceneResizeEvent_ITF /*777 QGraphicsSceneResizeEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QGraphicsSceneResizeEvent_PTR() != nil {
		convArg0 = event.QGraphicsSceneResizeEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget11resizeEventEP25QGraphicsSceneResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicsproxywidget.h:123
// index:0
// Protected Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * newProxyWidget(const QWidget *)

/*
Creates a proxy widget for the given child of the widget contained in this proxy.

You should not call this function directly; use QGraphicsProxyWidget::createProxyForChildWidget() instead.

This function is a fake virtual slot that you can reimplement in your subclass in order to control how new proxy widgets are created. The default implementation returns a proxy created with the QGraphicsProxyWidget() constructor with this proxy widget as the parent.

This function was introduced in  Qt 4.5.

See also createProxyForChildWidget().
*/
func (this *QGraphicsProxyWidget) NewProxyWidget(arg0 QWidget_ITF /*777 const QWidget **/) *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QGraphicsProxyWidget14newProxyWidgetEPK7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

/*


 */
type QGraphicsProxyWidget__ = int

//
const QGraphicsProxyWidget__Type QGraphicsProxyWidget__ = 12

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
