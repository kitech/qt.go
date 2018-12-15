package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

// bool event(QEvent *)
func (this *QQuickItem) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// bool isComponentComplete()
func (this *QQuickItem) InheritIsComponentComplete(f func() bool) {
	qtrt.SetAllInheritCallback(this, "isComponentComplete", f)
}

// void updateInputMethod(Qt::InputMethodQueries)
func (this *QQuickItem) InheritUpdateInputMethod(f func(queries int) /*void*/) {
	qtrt.SetAllInheritCallback(this, "updateInputMethod", f)
}

// bool widthValid()
func (this *QQuickItem) InheritWidthValid(f func() bool) {
	qtrt.SetAllInheritCallback(this, "widthValid", f)
}

// bool heightValid()
func (this *QQuickItem) InheritHeightValid(f func() bool) {
	qtrt.SetAllInheritCallback(this, "heightValid", f)
}

// void setImplicitSize(qreal, qreal)
func (this *QQuickItem) InheritSetImplicitSize(f func(arg0 float64, arg1 float64) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setImplicitSize", f)
}

// void classBegin()
func (this *QQuickItem) InheritClassBegin(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "classBegin", f)
}

// void componentComplete()
func (this *QQuickItem) InheritComponentComplete(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "componentComplete", f)
}

// void keyPressEvent(QKeyEvent *)
func (this *QQuickItem) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(QKeyEvent *)
func (this *QQuickItem) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void inputMethodEvent(QInputMethodEvent *)
func (this *QQuickItem) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void focusInEvent(QFocusEvent *)
func (this *QQuickItem) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(QFocusEvent *)
func (this *QQuickItem) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void mousePressEvent(QMouseEvent *)
func (this *QQuickItem) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(QMouseEvent *)
func (this *QQuickItem) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickItem) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickItem) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseUngrabEvent()
func (this *QQuickItem) InheritMouseUngrabEvent(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseUngrabEvent", f)
}

// void touchUngrabEvent()
func (this *QQuickItem) InheritTouchUngrabEvent(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "touchUngrabEvent", f)
}

// void wheelEvent(QWheelEvent *)
func (this *QQuickItem) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void touchEvent(QTouchEvent *)
func (this *QQuickItem) InheritTouchEvent(f func(event *qtgui.QTouchEvent /*777 QTouchEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "touchEvent", f)
}

// void hoverEnterEvent(QHoverEvent *)
func (this *QQuickItem) InheritHoverEnterEvent(f func(event *qtgui.QHoverEvent /*777 QHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverMoveEvent(QHoverEvent *)
func (this *QQuickItem) InheritHoverMoveEvent(f func(event *qtgui.QHoverEvent /*777 QHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(QHoverEvent *)
func (this *QQuickItem) InheritHoverLeaveEvent(f func(event *qtgui.QHoverEvent /*777 QHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void dragEnterEvent(QDragEnterEvent *)
func (this *QQuickItem) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(QDragMoveEvent *)
func (this *QQuickItem) InheritDragMoveEvent(f func(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QQuickItem) InheritDragLeaveEvent(f func(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(QDropEvent *)
func (this *QQuickItem) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// bool childMouseEventFilter(QQuickItem *, QEvent *)
func (this *QQuickItem) InheritChildMouseEventFilter(f func(arg0 *QQuickItem /*777 QQuickItem **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "childMouseEventFilter", f)
}

// void windowDeactivateEvent()
func (this *QQuickItem) InheritWindowDeactivateEvent(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "windowDeactivateEvent", f)
}

// void geometryChanged(const QRectF &, const QRectF &)
func (this *QQuickItem) InheritGeometryChanged(f func(newGeometry *qtcore.QRectF, oldGeometry *qtcore.QRectF) /*void*/) {
	qtrt.SetAllInheritCallback(this, "geometryChanged", f)
}

// void releaseResources()
func (this *QQuickItem) InheritReleaseResources(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "releaseResources", f)
}

// void updatePolish()
func (this *QQuickItem) InheritUpdatePolish(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "updatePolish", f)
}

/*

 */
type QQuickItem struct {
	*qtcore.QObject
	*qtqml.QQmlParserStatus
}
type QQuickItem_ITF interface {
	qtcore.QObject_ITF
	qtqml.QQmlParserStatus_ITF
	QQuickItem_PTR() *QQuickItem
}

func (ptr *QQuickItem) QQuickItem_PTR() *QQuickItem { return ptr }

func (this *QQuickItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQuickItem) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QQmlParserStatus = qtqml.NewQQmlParserStatusFromPointer(cthis)
}
func NewQQuickItemFromPointer(cthis unsafe.Pointer) *QQuickItem {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := qtqml.NewQQmlParserStatusFromPointer(cthis)
	return &QQuickItem{bcthis0, bcthis1}
}
func (*QQuickItem) NewFromPointer(cthis unsafe.Pointer) *QQuickItem {
	return NewQQuickItemFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qquickitem.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QQuickItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickItem(QQuickItem *)

/*
Constructs a QQuickItem with the given parent.
*/
func (*QQuickItem) NewForInherit(parent QQuickItem_ITF /*777 QQuickItem **/) *QQuickItem {
	return NewQQuickItem(parent)
}
func NewQQuickItem(parent QQuickItem_ITF /*777 QQuickItem **/) *QQuickItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQuickItem_PTR() != nil {
		convArg0 = parent.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItemC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickItem")
	return gothis
}

// /usr/include/qt/QtQuick/qquickitem.h:198
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickItem(QQuickItem *)

/*
Constructs a QQuickItem with the given parent.
*/
func (*QQuickItem) NewForInheritp() *QQuickItem {
	return NewQQuickItemp()
}
func NewQQuickItemp() *QQuickItem {
	// arg: 0, QQuickItem *=Pointer, QQuickItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItemC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QQuickItem")
	return gothis
}

// /usr/include/qt/QtQuick/qquickitem.h:199
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickItem()

/*

 */
func DeleteQQuickItem(this *QQuickItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickitem.h:201
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickWindow * window() const

/*
Returns the window in which this item is rendered.

The item does not have a window until it has been assigned into a scene. The windowChanged() signal provides a notification both when the item is entered into a scene and when it is removed from a scene.
*/
func (this *QQuickItem) Window() *QQuickWindow /*777 QQuickWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * parentItem() const

/*

 */
func (this *QQuickItem) ParentItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10parentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:203
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentItem(QQuickItem *)

/*

 */
func (this *QQuickItem) SetParentItem(parent QQuickItem_ITF /*777 QQuickItem **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QQuickItem_PTR() != nil {
		convArg0 = parent.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13setParentItemEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackBefore(const QQuickItem *)

/*
Moves the specified sibling item to the index before this item within the list of children. The order of children affects both the visual stacking order and tab focus navigation order.

Assuming the z values of both items are the same, this will cause sibling to be rendered above this item.

If both items have activeFocusOnTab set to true, this will also cause the tab focus order to change, with sibling receiving focus after this item.

The given sibling must be a sibling of this item; that is, they must have the same immediate parent.

See also Concepts - Visual Parent in Qt Quick.
*/
func (this *QQuickItem) StackBefore(arg0 QQuickItem_ITF /*777 const QQuickItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11stackBeforeEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackAfter(const QQuickItem *)

/*
Moves the specified sibling item to the index after this item within the list of children. The order of children affects both the visual stacking order and tab focus navigation order.

Assuming the z values of both items are the same, this will cause sibling to be rendered below this item.

If both items have activeFocusOnTab set to true, this will also cause the tab focus order to change, with sibling receiving focus before this item.

The given sibling must be a sibling of this item; that is, they must have the same immediate parent.

See also Concepts - Visual Parent in Qt Quick.
*/
func (this *QQuickItem) StackAfter(arg0 QQuickItem_ITF /*777 const QQuickItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10stackAfterEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:207
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF childrenRect()

/*

 */
func (this *QQuickItem) ChildrenRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12childrenRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:208
// index:0
// Public Visibility=Default Availability=Available
// [-2] QList<QQuickItem *> childItems() const

/*
Returns the children of this item.
*/
func (this *QQuickItem) ChildItems() *QQuickItemList /*lll*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10childItemsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuickItemListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:210
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clip() const

/*

 */
func (this *QQuickItem) Clip() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem4clipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:211
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClip(bool)

/*

 */
func (this *QQuickItem) SetClip(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setClipEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QString state() const

/*

 */
func (this *QQuickItem) State() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtQuick/qquickitem.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setState(const QString &)

/*

 */
func (this *QQuickItem) SetState(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setStateERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:216
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal baselineOffset() const

/*

 */
func (this *QQuickItem) BaselineOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14baselineOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaselineOffset(qreal)

/*

 */
func (this *QQuickItem) SetBaselineOffset(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17setBaselineOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:221
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal x() const

/*

 */
func (this *QQuickItem) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal y() const

/*

 */
func (this *QQuickItem) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:223
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position() const

/*

 */
func (this *QQuickItem) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(qreal)

/*

 */
func (this *QQuickItem) SetX(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(qreal)

/*

 */
func (this *QQuickItem) SetY(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)

/*

 */
func (this *QQuickItem) SetPosition(arg0 qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPointF_PTR() != nil {
		convArg0 = arg0.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:228
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width() const

/*

 */
func (this *QQuickItem) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:229
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(qreal)

/*

 */
func (this *QQuickItem) SetWidth(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetWidth()

/*

 */
func (this *QQuickItem) ResetWidth() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10resetWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImplicitWidth(qreal)

/*

 */
func (this *QQuickItem) SetImplicitWidth(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16setImplicitWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:232
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal implicitWidth() const

/*
Returns the width of the item that is implied by other properties that determine the content.

Note: Getter function for property implicitWidth.

See also setImplicitWidth().
*/
func (this *QQuickItem) ImplicitWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13implicitWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:234
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height() const

/*

 */
func (this *QQuickItem) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:235
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeight(qreal)

/*

 */
func (this *QQuickItem) SetHeight(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetHeight()

/*

 */
func (this *QQuickItem) ResetHeight() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11resetHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImplicitHeight(qreal)

/*

 */
func (this *QQuickItem) SetImplicitHeight(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17setImplicitHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:238
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal implicitHeight() const

/*

 */
func (this *QQuickItem) ImplicitHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14implicitHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSize(const QSizeF &)

/*

 */
func (this *QQuickItem) SetSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:242
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickItem::TransformOrigin transformOrigin() const

/*

 */
func (this *QQuickItem) TransformOrigin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15transformOriginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformOrigin(QQuickItem::TransformOrigin)

/*

 */
func (this *QQuickItem) SetTransformOrigin(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem18setTransformOriginENS_15TransformOriginE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:244
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF transformOriginPoint() const

/*

 */
func (this *QQuickItem) TransformOriginPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem20transformOriginPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformOriginPoint(const QPointF &)

/*

 */
func (this *QQuickItem) SetTransformOriginPoint(arg0 qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPointF_PTR() != nil {
		convArg0 = arg0.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem23setTransformOriginPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:247
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal z() const

/*

 */
func (this *QQuickItem) Z() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:248
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(qreal)

/*

 */
func (this *QQuickItem) SetZ(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem4setZEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:250
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rotation() const

/*

 */
func (this *QQuickItem) Rotation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8rotationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:251
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRotation(qreal)

/*

 */
func (this *QQuickItem) SetRotation(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11setRotationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:252
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal scale() const

/*

 */
func (this *QQuickItem) Scale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5scaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScale(qreal)

/*

 */
func (this *QQuickItem) SetScale(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:255
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity() const

/*

 */
func (this *QQuickItem) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)

/*

 */
func (this *QQuickItem) SetOpacity(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:258
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible() const

/*

 */
func (this *QQuickItem) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:259
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(bool)

/*

 */
func (this *QQuickItem) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*

 */
func (this *QQuickItem) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:262
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*

 */
func (this *QQuickItem) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:264
// index:0
// Public Visibility=Default Availability=Available
// [1] bool smooth() const

/*

 */
func (this *QQuickItem) Smooth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6smoothEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:265
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSmooth(bool)

/*

 */
func (this *QQuickItem) SetSmooth(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setSmoothEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:267
// index:0
// Public Visibility=Default Availability=Available
// [1] bool activeFocusOnTab() const

/*

 */
func (this *QQuickItem) ActiveFocusOnTab() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16activeFocusOnTabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveFocusOnTab(bool)

/*

 */
func (this *QQuickItem) SetActiveFocusOnTab(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem19setActiveFocusOnTabEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:270
// index:0
// Public Visibility=Default Availability=Available
// [1] bool antialiasing() const

/*

 */
func (this *QQuickItem) Antialiasing() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12antialiasingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAntialiasing(bool)

/*

 */
func (this *QQuickItem) SetAntialiasing(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15setAntialiasingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:272
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetAntialiasing()

/*

 */
func (this *QQuickItem) ResetAntialiasing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17resetAntialiasingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:274
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickItem::Flags flags() const

/*
Returns the item flags for this item.

See also setFlags() and setFlag().
*/
func (this *QQuickItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QQuickItem::Flag, bool)

/*
Enables the specified flag for this item if enabled is true; if enabled is false, the flag is disabled.

These provide various hints for the item; for example, the ItemClipsChildrenToShape flag indicates that all children of this item should be clipped to fit within the item area.
*/
func (this *QQuickItem) SetFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setFlagENS_4FlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QQuickItem::Flag, bool)

/*
Enables the specified flag for this item if enabled is true; if enabled is false, the flag is disabled.

These provide various hints for the item; for example, the ItemClipsChildrenToShape flag indicates that all children of this item should be clipped to fit within the item area.
*/
func (this *QQuickItem) SetFlagp(flag int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	enabled := true
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setFlagENS_4FlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QQuickItem::Flags)

/*
Enables the specified flags for this item.

See also flags() and setFlag().
*/
func (this *QQuickItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setFlagsE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:278
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect() const

/*

 */
func (this *QQuickItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:279
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF clipRect() const

/*

 */
func (this *QQuickItem) ClipRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8clipRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:281
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasActiveFocus() const

/*

 */
func (this *QQuickItem) HasActiveFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14hasActiveFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:282
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus() const

/*

 */
func (this *QQuickItem) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(bool)

/*

 */
func (this *QQuickItem) SetFocus(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:284
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFocus(bool, Qt::FocusReason)

/*

 */
func (this *QQuickItem) SetFocus1(focus bool, reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setFocusEbN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focus, reason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:285
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFocusScope() const

/*
Returns true if this item is a focus scope, and false otherwise.
*/
func (this *QQuickItem) IsFocusScope() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12isFocusScopeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:286
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * scopedFocusItem() const

/*
If this item is a focus scope, this returns the item in its focus chain that currently has focus.

Returns 0 if this item is not a focus scope.
*/
func (this *QQuickItem) ScopedFocusItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15scopedFocusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:288
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QQuickItem *) const

/*
Returns true if this item is an ancestor of child (i.e., if this item is child's parent, or one of child's parent's ancestors).

This function was introduced in  Qt 5.7.

See also parentItem().
*/
func (this *QQuickItem) IsAncestorOf(child QQuickItem_ITF /*777 const QQuickItem **/) bool {
	var convArg0 unsafe.Pointer
	if child != nil && child.QQuickItem_PTR() != nil {
		convArg0 = child.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12isAncestorOfEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:290
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons acceptedMouseButtons() const

/*
Returns the mouse buttons accepted by this item.

The default value is Qt::NoButton; that is, no mouse buttons are accepted.

If an item does not accept the mouse button for a particular mouse event, the mouse event will not be delivered to the item and will be delivered to the next item in the item hierarchy instead.

See also setAcceptedMouseButtons().
*/
func (this *QQuickItem) AcceptedMouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem20acceptedMouseButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptedMouseButtons(Qt::MouseButtons)

/*
Sets the mouse buttons accepted by this item to buttons.

See also acceptedMouseButtons().
*/
func (this *QQuickItem) SetAcceptedMouseButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem23setAcceptedMouseButtonsE6QFlagsIN2Qt11MouseButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:292
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptHoverEvents() const

/*
Returns whether hover events are accepted by this item.

The default value is false.

If this is false, then the item will not receive any hover events through the hoverEnterEvent(), hoverMoveEvent() and hoverLeaveEvent() functions.

See also setAcceptHoverEvents().
*/
func (this *QQuickItem) AcceptHoverEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem17acceptHoverEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptHoverEvents(bool)

/*
If enabled is true, this sets the item to accept hover events; otherwise, hover events are not accepted by this item.

See also acceptHoverEvents().
*/
func (this *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20setAcceptHoverEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:296
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor() const

/*
Returns the cursor shape for this item.

The mouse cursor will assume this shape when it is over this item, unless an override cursor is set. See the list of predefined cursor objects for a range of useful shapes.

If no cursor shape has been set this returns a cursor with the Qt::ArrowCursor shape, however another cursor shape may be displayed if an overlapping item has a valid cursor.

See also setCursor() and unsetCursor().
*/
func (this *QQuickItem) Cursor() *qtgui.QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)

/*
Sets the cursor shape for this item.

See also cursor() and unsetCursor().
*/
func (this *QQuickItem) SetCursor(cursor qtgui.QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QCursor_PTR() != nil {
		convArg0 = cursor.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()

/*
Clears the cursor shape for this item.

See also cursor() and setCursor().
*/
func (this *QQuickItem) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:301
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUnderMouse() const

/*

 */
func (this *QQuickItem) IsUnderMouse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12isUnderMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabMouse()

/*
Grabs the mouse input.

This item will receive all mouse events until ungrabMouse() is called. Usually this function should not be called, since accepting for example a mouse press event makes sure that the following events are delivered to the item. If an item wants to take over mouse events from the current receiver, it needs to call this function.

Warning: This function should be used with caution.
*/
func (this *QQuickItem) GrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9grabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:303
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabMouse()

/*
Releases the mouse grab following a call to grabMouse().

Note that this function should only be called when the item wants to stop handling further events. There is no need to call this function after a release or cancel event since no future events will be received in any case. No move or release events will be delivered after this function was called.
*/
func (this *QQuickItem) UngrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11ungrabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:304
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keepMouseGrab() const

/*
Returns whether mouse input should exclusively remain with this item.

See also setKeepMouseGrab().
*/
func (this *QQuickItem) KeepMouseGrab() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13keepMouseGrabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeepMouseGrab(bool)

/*
Sets whether the mouse input should remain exclusively with this item.

This is useful for items that wish to grab and keep mouse interaction following a predefined gesture. For example, an item that is interested in horizontal mouse movement may set keepMouseGrab to true once a threshold has been exceeded. Once keepMouseGrab has been set to true, filtering items will not react to mouse events.

If keep is false, a filtering item may steal the grab. For example, Flickable may attempt to steal a mouse grab if it detects that the user has begun to move the viewport.

See also keepMouseGrab().
*/
func (this *QQuickItem) SetKeepMouseGrab(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16setKeepMouseGrabEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:306
// index:0
// Public Visibility=Default Availability=Available
// [1] bool filtersChildMouseEvents() const

/*
Returns whether mouse and touch events of this item's children should be filtered through this item.

See also setFiltersChildMouseEvents() and childMouseEventFilter().
*/
func (this *QQuickItem) FiltersChildMouseEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem23filtersChildMouseEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltersChildMouseEvents(bool)

/*
Sets whether mouse and touch events of this item's children should be filtered through this item.

If filter is true, childMouseEventFilter() will be called when a mouse event is triggered for a child item.

See also filtersChildMouseEvents().
*/
func (this *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem26setFiltersChildMouseEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:310
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabTouchPoints()

/*
Ungrabs the touch points owned by this item.

Note: there is hardly any reason to call this function. It should only be called when an item does not want to receive any further events, so no move or release events will be delivered after calling this function.

See also grabTouchPoints().
*/
func (this *QQuickItem) UngrabTouchPoints() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17ungrabTouchPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:311
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keepTouchGrab() const

/*
Returns whether the touch points grabbed by this item should exclusively remain with this item.

See also setKeepTouchGrab() and keepMouseGrab().
*/
func (this *QQuickItem) KeepTouchGrab() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13keepTouchGrabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:312
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeepTouchGrab(bool)

/*
Sets whether the touch points grabbed by this item should remain exclusively with this item.

This is useful for items that wish to grab and keep specific touch points following a predefined gesture. For example, an item that is interested in horizontal touch point movement may set setKeepTouchGrab to true once a threshold has been exceeded. Once setKeepTouchGrab has been set to true, filtering items will not react to the relevant touch points.

If keep is false, a filtering item may steal the grab. For example, Flickable may attempt to steal a touch point grab if it detects that the user has begun to move the viewport.

See also keepTouchGrab() and setKeepMouseGrab().
*/
func (this *QQuickItem) SetKeepTouchGrab(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16setKeepTouchGrabEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:315
// index:0
// Public Visibility=Default Availability=Available
// [1] bool grabToImage(const QJSValue &, const QSize &)

/*
Grabs the item into an in-memory image.

The grab happens asynchronously and the signal QQuickItemGrabResult::ready() is emitted when the grab has been completed.

Use targetSize to specify the size of the target image. By default, the result will have the same size as item.

If the grab could not be initiated, the function returns null.

Note: This function will render the item to an offscreen surface and copy that surface from the GPU's memory into the CPU's memory, which can be quite costly. For "live" preview, use layers or ShaderEffectSource.

See also QQuickWindow::grabWindow().
*/
func (this *QQuickItem) GrabToImage(callback qtqml.QJSValue_ITF, targetSize qtcore.QSize_ITF) bool {
	var convArg0 unsafe.Pointer
	if callback != nil && callback.QJSValue_PTR() != nil {
		convArg0 = callback.QJSValue_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if targetSize != nil && targetSize.QSize_PTR() != nil {
		convArg1 = targetSize.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11grabToImageERK8QJSValueRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:315
// index:0
// Public Visibility=Default Availability=Available
// [1] bool grabToImage(const QJSValue &, const QSize &)

/*
Grabs the item into an in-memory image.

The grab happens asynchronously and the signal QQuickItemGrabResult::ready() is emitted when the grab has been completed.

Use targetSize to specify the size of the target image. By default, the result will have the same size as item.

If the grab could not be initiated, the function returns null.

Note: This function will render the item to an offscreen surface and copy that surface from the GPU's memory into the CPU's memory, which can be quite costly. For "live" preview, use layers or ShaderEffectSource.

See also QQuickWindow::grabWindow().
*/
func (this *QQuickItem) GrabToImagep(callback qtqml.QJSValue_ITF) bool {
	var convArg0 unsafe.Pointer
	if callback != nil && callback.QJSValue_PTR() != nil {
		convArg0 = callback.QJSValue_PTR().GetCthis()
	}
	// arg: 1, const QSize &=LValueReference, QSize=Record, , Invalid
	var convArg1 = qtcore.NewQSize()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11grabToImageERK8QJSValueRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:318
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &) const

/*
Returns true if this item contains point, which is in local coordinates; returns false otherwise.

This function can be overwritten in order to handle point collisions in items with custom shapes. The default implementation checks if the point is inside the item's bounding rect.

Note that this method is generally used to check whether the item is under the mouse cursor, and for that reason, the implementation of this function should be as light-weight as possible.
*/
func (this *QQuickItem) Contains(point qtcore.QPointF_ITF) bool {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:320
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform itemTransform(QQuickItem *, bool *) const

/*

 */
func (this *QQuickItem) ItemTransform(arg0 QQuickItem_ITF /*777 QQuickItem **/, arg1 *bool) *qtgui.QTransform /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13itemTransformEPS_Pb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, arg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:321
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToItem(const QQuickItem *, const QPointF &) const

/*
Maps the given point in this item's coordinate system to the equivalent point within item's coordinate system, and returns the mapped coordinate.

If item is 0, this maps point to the coordinate system of the scene.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapToItem(item QQuickItem_ITF /*777 const QQuickItem **/, point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QQuickItem_PTR() != nil {
		convArg0 = item.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg1 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem9mapToItemEPKS_RK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:322
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToScene(const QPointF &) const

/*
Maps the given point in this item's coordinate system to the equivalent point within the scene's coordinate system, and returns the mapped coordinate.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapToScene(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10mapToSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:323
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToGlobal(const QPointF &) const

/*
Maps the given point in this item's coordinate system to the equivalent point within global screen coordinate system, and returns the mapped coordinate.

For example, this may be helpful to add a popup to a Qt Quick component.

Note: Window positioning is done by the window manager and this value is treated only as a hint. So, the resulting window position may differ from what is expected.

This function was introduced in  Qt 5.7.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapToGlobal(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem11mapToGlobalERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:324
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToItem(const QQuickItem *, const QRectF &) const

/*
Maps the given rect in this item's coordinate system to the equivalent rectangular area within item's coordinate system, and returns the mapped rectangle value.

If item is 0, this maps rect to the coordinate system of the scene.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapRectToItem(item QQuickItem_ITF /*777 const QQuickItem **/, rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QQuickItem_PTR() != nil {
		convArg0 = item.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13mapRectToItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:325
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToScene(const QRectF &) const

/*
Maps the given rect in this item's coordinate system to the equivalent rectangular area within the scene's coordinate system, and returns the mapped rectangle value.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapRectToScene(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14mapRectToSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:326
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromItem(const QQuickItem *, const QPointF &) const

/*
Maps the given point in item's coordinate system to the equivalent point within this item's coordinate system, and returns the mapped coordinate.

If item is 0, this maps point from the coordinate system of the scene.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapFromItem(item QQuickItem_ITF /*777 const QQuickItem **/, point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QQuickItem_PTR() != nil {
		convArg0 = item.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg1 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem11mapFromItemEPKS_RK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:327
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromScene(const QPointF &) const

/*
Maps the given point in the scene's coordinate system to the equivalent point within this item's coordinate system, and returns the mapped coordinate.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapFromScene(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12mapFromSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:328
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromGlobal(const QPointF &) const

/*
Maps the given point in the global screen coordinate system to the equivalent point within this item's coordinate system, and returns the mapped coordinate.

For example, this may be helpful to add a popup to a Qt Quick component.

Note: Window positioning is done by the window manager and this value is treated only as a hint. So, the resulting window position may differ from what is expected.

This function was introduced in  Qt 5.7.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapFromGlobal(point qtcore.QPointF_ITF) *qtcore.QPointF /*123*/ {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13mapFromGlobalERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:329
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromItem(const QQuickItem *, const QRectF &) const

/*
Maps the given rect in item's coordinate system to the equivalent rectangular area within this item's coordinate system, and returns the mapped rectangle value.

If item is 0, this maps rect from the coordinate system of the scene.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapRectFromItem(item QQuickItem_ITF /*777 const QQuickItem **/, rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if item != nil && item.QQuickItem_PTR() != nil {
		convArg0 = item.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg1 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15mapRectFromItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:330
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromScene(const QRectF &) const

/*
Maps the given rect in the scene's coordinate system to the equivalent rectangular area within this item's coordinate system, and returns the mapped rectangle value.

See also Concepts - Visual Coordinates in Qt Quick.
*/
func (this *QQuickItem) MapRectFromScene(rect qtcore.QRectF_ITF) *qtcore.QRectF /*123*/ {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16mapRectFromSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void polish()

/*
Schedules a polish event for this item.

When the scene graph processes the request, it will call updatePolish() on this item.
*/
func (this *QQuickItem) Polish() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem6polishEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forceActiveFocus()

/*
Forces active focus on the item.

This method sets focus on the item and ensures that all ancestor FocusScope objects in the object hierarchy are also given focus.

The reason for the focus change will be Qt::OtherFocusReason. Use the overloaded method to specify the focus reason to enable better handling of the focus change.

See also activeFocus.
*/
func (this *QQuickItem) ForceActiveFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16forceActiveFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:339
// index:1
// Public Visibility=Default Availability=Available
// [-2] void forceActiveFocus(Qt::FocusReason)

/*
Forces active focus on the item.

This method sets focus on the item and ensures that all ancestor FocusScope objects in the object hierarchy are also given focus.

The reason for the focus change will be Qt::OtherFocusReason. Use the overloaded method to specify the focus reason to enable better handling of the focus change.

See also activeFocus.
*/
func (this *QQuickItem) ForceActiveFocus1(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16forceActiveFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:340
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * nextItemInFocusChain(bool)

/*
Returns the item in the focus chain which is next to this item. If forward is true, or not supplied, it is the next item in the forwards direction. If forward is false, it is the next item in the backwards direction.
*/
func (this *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20nextItemInFocusChainEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), forward)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:340
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * nextItemInFocusChain(bool)

/*
Returns the item in the focus chain which is next to this item. If forward is true, or not supplied, it is the next item in the forwards direction. If forward is false, it is the next item in the backwards direction.
*/
func (this *QQuickItem) NextItemInFocusChainp() *QQuickItem /*777 QQuickItem **/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	forward := true
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20nextItemInFocusChainEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), forward)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:341
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * childAt(qreal, qreal) const

/*
Returns the first visible child item found at point (x, y) within the coordinate system of this item.

Returns 0 if there is no such item.
*/
func (this *QQuickItem) ChildAt(x float64, y float64) *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem7childAtEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:344
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery) const

/*
This method is only relevant for input items.

If this item is an input item, this method should be reimplemented to return the relevant input method flags for the given query.

See also QWidget::inputMethodQuery().
*/
func (this *QQuickItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:354
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isTextureProvider() const

/*
Returns true if this item is a texture provider. The default implementation returns false.

This function can be called from any thread.
*/
func (this *QQuickItem) IsTextureProvider() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem17isTextureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:355
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGTextureProvider * textureProvider() const

/*
Returns the texture provider for an item. The default implementation returns 0.

This function may only be called on the rendering thread.
*/
func (this *QQuickItem) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15textureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qquickitem.h:358
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()

/*
Schedules a call to updatePaintNode() for this item.

The call to QQuickItem::updatePaintNode() will always happen if the item is showing in a QQuickWindow.

Only items which specify QQuickItem::ItemHasContents are allowed to call QQuickItem::update().
*/
func (this *QQuickItem) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:361
// index:0
// Public Visibility=Default Availability=Available
// [-2] void childrenRectChanged(const QRectF &)

/*

 */
func (this *QQuickItem) ChildrenRectChanged(arg0 qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRectF_PTR() != nil {
		convArg0 = arg0.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem19childrenRectChangedERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:362
// index:0
// Public Visibility=Default Availability=Available
// [-2] void baselineOffsetChanged(qreal)

/*

 */
func (this *QQuickItem) BaselineOffsetChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21baselineOffsetChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:363
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(const QString &)

/*

 */
func (this *QQuickItem) StateChanged(arg0 string) {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12stateChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:364
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusChanged(bool)

/*

 */
func (this *QQuickItem) FocusChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12focusChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:365
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusChanged(bool)

/*

 */
func (this *QQuickItem) ActiveFocusChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem18activeFocusChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:366
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusOnTabChanged(bool)

/*

 */
func (this *QQuickItem) ActiveFocusOnTabChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem23activeFocusOnTabChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:367
// index:0
// Public Visibility=Default Availability=Available
// [-2] void parentChanged(QQuickItem *)

/*

 */
func (this *QQuickItem) ParentChanged(arg0 QQuickItem_ITF /*777 QQuickItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13parentChangedEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:368
// index:0
// Public Visibility=Default Availability=Available
// [-2] void transformOriginChanged(QQuickItem::TransformOrigin)

/*

 */
func (this *QQuickItem) TransformOriginChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem22transformOriginChangedENS_15TransformOriginE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:369
// index:0
// Public Visibility=Default Availability=Available
// [-2] void smoothChanged(bool)

/*

 */
func (this *QQuickItem) SmoothChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13smoothChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:370
// index:0
// Public Visibility=Default Availability=Available
// [-2] void antialiasingChanged(bool)

/*

 */
func (this *QQuickItem) AntialiasingChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem19antialiasingChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:371
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clipChanged(bool)

/*

 */
func (this *QQuickItem) ClipChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11clipChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:372
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowChanged(QQuickWindow *)

/*
This signal is emitted when the item's window changes.
*/
func (this *QQuickItem) WindowChanged(window QQuickWindow_ITF /*777 QQuickWindow **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QQuickWindow_PTR() != nil {
		convArg0 = window.QQuickWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13windowChangedEP12QQuickWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:374
// index:0
// Public Visibility=Default Availability=Available
// [-2] void childrenChanged()

/*

 */
func (this *QQuickItem) ChildrenChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15childrenChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:375
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityChanged()

/*

 */
func (this *QQuickItem) OpacityChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14opacityChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:376
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enabledChanged()

/*

 */
func (this *QQuickItem) EnabledChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14enabledChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:377
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChanged()

/*

 */
func (this *QQuickItem) VisibleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14visibleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:378
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChildrenChanged()

/*

 */
func (this *QQuickItem) VisibleChildrenChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem22visibleChildrenChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotationChanged()

/*

 */
func (this *QQuickItem) RotationChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15rotationChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:380
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scaleChanged()

/*

 */
func (this *QQuickItem) ScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12scaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:382
// index:0
// Public Visibility=Default Availability=Available
// [-2] void xChanged()

/*

 */
func (this *QQuickItem) XChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8xChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:383
// index:0
// Public Visibility=Default Availability=Available
// [-2] void yChanged()

/*

 */
func (this *QQuickItem) YChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8yChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widthChanged()

/*

 */
func (this *QQuickItem) WidthChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12widthChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:385
// index:0
// Public Visibility=Default Availability=Available
// [-2] void heightChanged()

/*

 */
func (this *QQuickItem) HeightChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13heightChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:386
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zChanged()

/*

 */
func (this *QQuickItem) ZChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8zChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void implicitWidthChanged()

/*

 */
func (this *QQuickItem) ImplicitWidthChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20implicitWidthChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:388
// index:0
// Public Visibility=Default Availability=Available
// [-2] void implicitHeightChanged()

/*

 */
func (this *QQuickItem) ImplicitHeightChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21implicitHeightChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:391
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QQuickItem) Event(arg0 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:393
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool isComponentComplete() const

/*
Returns true if construction of the QML component is complete; otherwise returns false.

It is often desirable to delay some processing until the component is completed.

See also componentComplete().
*/
func (this *QQuickItem) IsComponentComplete() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem19isComponentCompleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:397
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateInputMethod(Qt::InputMethodQueries)

/*
Notify input method on updated query values if needed. queries indicates the changed attributes.
*/
func (this *QQuickItem) UpdateInputMethod(queries int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17updateInputMethodE6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), queries)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:397
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateInputMethod(Qt::InputMethodQueries)

/*
Notify input method on updated query values if needed. queries indicates the changed attributes.
*/
func (this *QQuickItem) UpdateInputMethodp() {
	// arg: 0, Qt::InputMethodQueries=Elaborated, Qt::InputMethodQueries=Typedef, QFlags<Qt::InputMethodQuery>, Unexposed
	queries := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17updateInputMethodE6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), queries)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:400
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool widthValid() const

/*
Returns whether the width property has been set explicitly.
*/
func (this *QQuickItem) WidthValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10widthValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:401
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool heightValid() const

/*
Returns whether the height property has been set explicitly.
*/
func (this *QQuickItem) HeightValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem11heightValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:402
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setImplicitSize(qreal, qreal)

/*

 */
func (this *QQuickItem) SetImplicitSize(arg0 float64, arg1 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15setImplicitSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:404
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void classBegin()

/*
Reimplemented from QQmlParserStatus::classBegin().

Derived classes should call the base class method before adding their own action to perform at classBegin.
*/
func (this *QQuickItem) ClassBegin() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10classBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:405
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void componentComplete()

/*
Reimplemented from QQmlParserStatus::componentComplete().

Derived classes should call the base class method before adding their own actions to perform at componentComplete.
*/
func (this *QQuickItem) ComponentComplete() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17componentCompleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:407
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)

/*
This event handler can be reimplemented in a subclass to receive key press events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) KeyPressEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:408
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)

/*
This event handler can be reimplemented in a subclass to receive key release events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) KeyReleaseEvent(event qtgui.QKeyEvent_ITF /*777 QKeyEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QKeyEvent_PTR() != nil {
		convArg0 = event.QKeyEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:410
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)

/*
This event handler can be reimplemented in a subclass to receive input method events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) InputMethodEvent(arg0 qtgui.QInputMethodEvent_ITF /*777 QInputMethodEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QInputMethodEvent_PTR() != nil {
		convArg0 = arg0.QInputMethodEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:412
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)

/*
This event handler can be reimplemented in a subclass to receive focus-in events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) FocusInEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:413
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)

/*
This event handler can be reimplemented in a subclass to receive focus-out events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) FocusOutEvent(arg0 qtgui.QFocusEvent_ITF /*777 QFocusEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QFocusEvent_PTR() != nil {
		convArg0 = arg0.QFocusEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:414
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)

/*
This event handler can be reimplemented in a subclass to receive mouse press events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) MousePressEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:415
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)

/*
This event handler can be reimplemented in a subclass to receive mouse move events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) MouseMoveEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:416
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)

/*
This event handler can be reimplemented in a subclass to receive mouse release events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) MouseReleaseEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:417
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)

/*
This event handler can be reimplemented in a subclass to receive mouse double-click events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) MouseDoubleClickEvent(event qtgui.QMouseEvent_ITF /*777 QMouseEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QMouseEvent_PTR() != nil {
		convArg0 = event.QMouseEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:418
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseUngrabEvent()

/*
This event handler can be reimplemented in a subclass to be notified when a mouse ungrab event has occurred on this item.

See also ungrabMouse().
*/
func (this *QQuickItem) MouseUngrabEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16mouseUngrabEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:419
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void touchUngrabEvent()

/*
This event handler can be reimplemented in a subclass to be notified when a touch ungrab event has occurred on this item.
*/
func (this *QQuickItem) TouchUngrabEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16touchUngrabEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:421
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)

/*
This event handler can be reimplemented in a subclass to receive wheel events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) WheelEvent(event qtgui.QWheelEvent_ITF /*777 QWheelEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QWheelEvent_PTR() != nil {
		convArg0 = event.QWheelEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:423
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void touchEvent(QTouchEvent *)

/*
This event handler can be reimplemented in a subclass to receive touch events for an item. The event information is provided by the event parameter.
*/
func (this *QQuickItem) TouchEvent(event qtgui.QTouchEvent_ITF /*777 QTouchEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QTouchEvent_PTR() != nil {
		convArg0 = event.QTouchEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10touchEventEP11QTouchEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:424
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverEnterEvent(QHoverEvent *)

/*
This event handler can be reimplemented in a subclass to receive hover-enter events for an item. The event information is provided by the event parameter.

Hover events are only provided if acceptHoverEvents() is true.
*/
func (this *QQuickItem) HoverEnterEvent(event qtgui.QHoverEvent_ITF /*777 QHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHoverEvent_PTR() != nil {
		convArg0 = event.QHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15hoverEnterEventEP11QHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:425
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QHoverEvent *)

/*
This event handler can be reimplemented in a subclass to receive hover-move events for an item. The event information is provided by the event parameter.

Hover events are only provided if acceptHoverEvents() is true.
*/
func (this *QQuickItem) HoverMoveEvent(event qtgui.QHoverEvent_ITF /*777 QHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHoverEvent_PTR() != nil {
		convArg0 = event.QHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14hoverMoveEventEP11QHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:426
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QHoverEvent *)

/*
This event handler can be reimplemented in a subclass to receive hover-leave events for an item. The event information is provided by the event parameter.

Hover events are only provided if acceptHoverEvents() is true.
*/
func (this *QQuickItem) HoverLeaveEvent(event qtgui.QHoverEvent_ITF /*777 QHoverEvent **/) {
	var convArg0 unsafe.Pointer
	if event != nil && event.QHoverEvent_PTR() != nil {
		convArg0 = event.QHoverEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15hoverLeaveEventEP11QHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:428
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)

/*
This event handler can be reimplemented in a subclass to receive drag-enter events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DragEnterEvent(arg0 qtgui.QDragEnterEvent_ITF /*777 QDragEnterEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragEnterEvent_PTR() != nil {
		convArg0 = arg0.QDragEnterEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:429
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)

/*
This event handler can be reimplemented in a subclass to receive drag-move events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DragMoveEvent(arg0 qtgui.QDragMoveEvent_ITF /*777 QDragMoveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragMoveEvent_PTR() != nil {
		convArg0 = arg0.QDragMoveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:430
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)

/*
This event handler can be reimplemented in a subclass to receive drag-leave events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DragLeaveEvent(arg0 qtgui.QDragLeaveEvent_ITF /*777 QDragLeaveEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDragLeaveEvent_PTR() != nil {
		convArg0 = arg0.QDragLeaveEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:431
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)

/*
This event handler can be reimplemented in a subclass to receive drop events for an item. The event information is provided by the event parameter.

Drag and drop events are only provided if the ItemAcceptsDrops flag has been set for this item.

See also Drag and Drag and Drop.
*/
func (this *QQuickItem) DropEvent(arg0 qtgui.QDropEvent_ITF /*777 QDropEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QDropEvent_PTR() != nil {
		convArg0 = arg0.QDropEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:433
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool childMouseEventFilter(QQuickItem *, QEvent *)

/*
Reimplement this method to filter the mouse events that are received by this item's children.

This method will only be called if filtersChildMouseEvents() is true.

Return true if the specified event should not be passed onto the specified child item, and false otherwise.

See also setFiltersChildMouseEvents().
*/
func (this *QQuickItem) ChildMouseEventFilter(arg0 QQuickItem_ITF /*777 QQuickItem **/, arg1 qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QQuickItem_PTR() != nil {
		convArg0 = arg0.QQuickItem_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QEvent_PTR() != nil {
		convArg1 = arg1.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21childMouseEventFilterEPS_P6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:434
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void windowDeactivateEvent()

/*

 */
func (this *QQuickItem) WindowDeactivateEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21windowDeactivateEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:436
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void geometryChanged(const QRectF &, const QRectF &)

/*
This function is called to handle this item's changes in geometry from oldGeometry to newGeometry. If the two geometries are the same, it doesn't do anything.

Derived classes must call the base class method within their implementation.
*/
func (this *QQuickItem) GeometryChanged(newGeometry qtcore.QRectF_ITF, oldGeometry qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if newGeometry != nil && newGeometry.QRectF_PTR() != nil {
		convArg0 = newGeometry.QRectF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if oldGeometry != nil && oldGeometry.QRectF_PTR() != nil {
		convArg1 = oldGeometry.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15geometryChangedERK6QRectFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:440
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void releaseResources()

/*
This function is called when an item should release graphics resources which are not already managed by the nodes returned from QQuickItem::updatePaintNode().

This happens when the item is about to be removed from the window it was previously rendering to. The item is guaranteed to have a window when the function is called.

The function is called on the GUI thread and the state of the rendering thread, when it is used, is unknown. Objects should not be deleted directly, but instead scheduled for cleanup using QQuickWindow::scheduleRenderJob().

See also Graphics Resource Handling.
*/
func (this *QQuickItem) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:441
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updatePolish()

/*
This function should perform any layout as required for this item.

When polish() is called, the scene graph schedules a polish event for this item. When the scene graph is ready to render this item, it calls updatePolish() to do any item layout as required before it renders the next frame.
*/
func (this *QQuickItem) UpdatePolish() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12updatePolishEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QQuickItem__Flag = int

//
const QQuickItem__ItemClipsChildrenToShape QQuickItem__Flag = 1

//
const QQuickItem__ItemAcceptsInputMethod QQuickItem__Flag = 2

//
const QQuickItem__ItemIsFocusScope QQuickItem__Flag = 4

//
const QQuickItem__ItemHasContents QQuickItem__Flag = 8

//
const QQuickItem__ItemAcceptsDrops QQuickItem__Flag = 16

func (this *QQuickItem) FlagItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickItem_FlagItemName(val int) string {
	var nilthis *QQuickItem
	return nilthis.FlagItemName(val)
}

/*
Used in conjunction with QQuickItem::itemChange() to notify the item about certain types of changes.


*/
type QQuickItem__ItemChange = int

// A child was added. ItemChangeData::item contains the added child.
const QQuickItem__ItemChildAddedChange QQuickItem__ItemChange = 0

// A child was removed. ItemChangeData::item contains the removed child.
const QQuickItem__ItemChildRemovedChange QQuickItem__ItemChange = 1

// The item was added to or removed from a scene. The QQuickWindow rendering the scene is specified in using ItemChangeData::window. The window parameter is null when the item is removed from a scene.
const QQuickItem__ItemSceneChange QQuickItem__ItemChange = 2

// The item's visibility has changed. ItemChangeData::boolValue contains the new visibility.
const QQuickItem__ItemVisibleHasChanged QQuickItem__ItemChange = 3

// The item's parent has changed. ItemChangeData::item contains the new parent.
const QQuickItem__ItemParentHasChanged QQuickItem__ItemChange = 4

// The item's opacity has changed. ItemChangeData::realValue contains the new opacity.
const QQuickItem__ItemOpacityHasChanged QQuickItem__ItemChange = 5

// The item's focus has changed. ItemChangeData::boolValue contains whether the item has focus or not.
const QQuickItem__ItemActiveFocusHasChanged QQuickItem__ItemChange = 6

// The item's rotation has changed. ItemChangeData::realValue contains the new rotation.
const QQuickItem__ItemRotationHasChanged QQuickItem__ItemChange = 7

// The antialiasing has changed. The current (boolean) value can be found in QQuickItem::antialiasing.
const QQuickItem__ItemAntialiasingHasChanged QQuickItem__ItemChange = 8

// The device pixel ratio of the screen the item is on has changed. ItemChangedData::realValue contains the new device pixel ratio.
const QQuickItem__ItemDevicePixelRatioHasChanged QQuickItem__ItemChange = 9

func (this *QQuickItem) ItemChangeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickItem_ItemChangeItemName(val int) string {
	var nilthis *QQuickItem
	return nilthis.ItemChangeItemName(val)
}

/*
Controls the point about which simple transforms like scale apply.



See also transformOrigin() and setTransformOrigin().

*/
type QQuickItem__TransformOrigin = int

// The top-left corner of the item.
const QQuickItem__TopLeft QQuickItem__TransformOrigin = 0

// The center point of the top of the item.
const QQuickItem__Top QQuickItem__TransformOrigin = 1

// The top-right corner of the item.
const QQuickItem__TopRight QQuickItem__TransformOrigin = 2

// The left most point of the vertical middle.
const QQuickItem__Left QQuickItem__TransformOrigin = 3

// The center of the item.
const QQuickItem__Center QQuickItem__TransformOrigin = 4

// The right most point of the vertical middle.
const QQuickItem__Right QQuickItem__TransformOrigin = 5

// The bottom-left corner of the item.
const QQuickItem__BottomLeft QQuickItem__TransformOrigin = 6

// The center point of the bottom of the item.
const QQuickItem__Bottom QQuickItem__TransformOrigin = 7

// The bottom-right corner of the item.
const QQuickItem__BottomRight QQuickItem__TransformOrigin = 8

func (this *QQuickItem) TransformOriginItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QQuickItem_TransformOriginItemName(val int) string {
	var nilthis *QQuickItem
	return nilthis.TransformOriginItemName(val)
}

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
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  keep block end
