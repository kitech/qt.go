package qtquick

// /usr/include/qt/QtQuick/qquickitem.h
// #include <qquickitem.h>
// #include <QtQuick>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtnetwork"
import "qt.go/qtgui"
import "qt.go/qtqml"

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
		gopp.KeepMe()
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

//  ext block end

//  body block begin
// bool event(class QEvent *)
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

// void keyPressEvent(class QKeyEvent *)
func (this *QQuickItem) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QQuickItem) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QQuickItem) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QQuickItem) InheritFocusInEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QQuickItem) InheritFocusOutEvent(f func(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QQuickItem) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QQuickItem) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QQuickItem) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
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

// void wheelEvent(class QWheelEvent *)
func (this *QQuickItem) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void touchEvent(class QTouchEvent *)
func (this *QQuickItem) InheritTouchEvent(f func(event *qtgui.QTouchEvent /*777 QTouchEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "touchEvent", f)
}

// void hoverEnterEvent(class QHoverEvent *)
func (this *QQuickItem) InheritHoverEnterEvent(f func(event *qtgui.QHoverEvent /*777 QHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverEnterEvent", f)
}

// void hoverMoveEvent(class QHoverEvent *)
func (this *QQuickItem) InheritHoverMoveEvent(f func(event *qtgui.QHoverEvent /*777 QHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverMoveEvent", f)
}

// void hoverLeaveEvent(class QHoverEvent *)
func (this *QQuickItem) InheritHoverLeaveEvent(f func(event *qtgui.QHoverEvent /*777 QHoverEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hoverLeaveEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QQuickItem) InheritDragEnterEvent(f func(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QQuickItem) InheritDragMoveEvent(f func(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QQuickItem) InheritDragLeaveEvent(f func(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QQuickItem) InheritDropEvent(f func(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// bool childMouseEventFilter(class QQuickItem *, class QEvent *)
func (this *QQuickItem) InheritChildMouseEventFilter(f func(arg0 *QQuickItem /*777 QQuickItem **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "childMouseEventFilter", f)
}

// void windowDeactivateEvent()
func (this *QQuickItem) InheritWindowDeactivateEvent(f func() /*void*/) {
	qtrt.SetAllInheritCallback(this, "windowDeactivateEvent", f)
}

// void geometryChanged(const class QRectF &, const class QRectF &)
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

type QQuickItem struct {
	*qtcore.QObject
	*qtqml.QQmlParserStatus
}

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
// [8] const QMetaObject * metaObject()
func (this *QQuickItem) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQuickItem(QQuickItem *)
func NewQQuickItem(parent *QQuickItem /*777 QQuickItem **/) *QQuickItem {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItemC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQuick/qquickitem.h:200
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQuickItem()
func DeleteQQuickItem(this *QQuickItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qquickitem.h:202
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickWindow * window()
func (this *QQuickItem) Window() *QQuickWindow /*777 QQuickWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:203
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * parentItem()
func (this *QQuickItem) ParentItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10parentItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:204
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentItem(QQuickItem *)
func (this *QQuickItem) SetParentItem(parent *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13setParentItemEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:205
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackBefore(const QQuickItem *)
func (this *QQuickItem) StackBefore(arg0 *QQuickItem /*777 const QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11stackBeforeEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:206
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackAfter(const QQuickItem *)
func (this *QQuickItem) StackAfter(arg0 *QQuickItem /*777 const QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10stackAfterEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:208
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF childrenRect()
func (this *QQuickItem) ChildrenRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12childrenRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:211
// index:0
// Public Visibility=Default Availability=Available
// [1] bool clip()
func (this *QQuickItem) Clip() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem4clipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:212
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setClip(_Bool)
func (this *QQuickItem) SetClip(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setClipEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QString state()
func (this *QQuickItem) State() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:215
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setState(const QString &)
func (this *QQuickItem) SetState(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setStateERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:217
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal baselineOffset()
func (this *QQuickItem) BaselineOffset() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14baselineOffsetEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaselineOffset(qreal)
func (this *QQuickItem) SetBaselineOffset(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17setBaselineOffsetEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal x()
func (this *QQuickItem) X() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem1xEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:223
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal y()
func (this *QQuickItem) Y() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem1yEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:224
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF position()
func (this *QQuickItem) Position() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:225
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(qreal)
func (this *QQuickItem) SetX(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem4setXEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(qreal)
func (this *QQuickItem) SetY(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem4setYEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPointF &)
func (this *QQuickItem) SetPosition(arg0 *qtcore.QPointF) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11setPositionERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:229
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal width()
func (this *QQuickItem) Width() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5widthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(qreal)
func (this *QQuickItem) SetWidth(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:231
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetWidth()
func (this *QQuickItem) ResetWidth() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10resetWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:232
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImplicitWidth(qreal)
func (this *QQuickItem) SetImplicitWidth(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16setImplicitWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:233
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal implicitWidth()
func (this *QQuickItem) ImplicitWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13implicitWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:235
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal height()
func (this *QQuickItem) Height() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6heightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeight(qreal)
func (this *QQuickItem) SetHeight(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetHeight()
func (this *QQuickItem) ResetHeight() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11resetHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:238
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setImplicitHeight(qreal)
func (this *QQuickItem) SetImplicitHeight(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17setImplicitHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:239
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal implicitHeight()
func (this *QQuickItem) ImplicitHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14implicitHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:241
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF size()
func (this *QQuickItem) Size() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSize(const QSizeF &)
func (this *QQuickItem) SetSize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:244
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickItem::TransformOrigin transformOrigin()
func (this *QQuickItem) TransformOrigin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15transformOriginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformOrigin(enum QQuickItem::TransformOrigin)
func (this *QQuickItem) SetTransformOrigin(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem18setTransformOriginENS_15TransformOriginE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:246
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF transformOriginPoint()
func (this *QQuickItem) TransformOriginPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem20transformOriginPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:247
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransformOriginPoint(const QPointF &)
func (this *QQuickItem) SetTransformOriginPoint(arg0 *qtcore.QPointF) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem23setTransformOriginPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:249
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal z()
func (this *QQuickItem) Z() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem1zEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setZ(qreal)
func (this *QQuickItem) SetZ(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem4setZEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:252
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal rotation()
func (this *QQuickItem) Rotation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8rotationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:253
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRotation(qreal)
func (this *QQuickItem) SetRotation(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11setRotationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:254
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal scale()
func (this *QQuickItem) Scale() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5scaleEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:255
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScale(qreal)
func (this *QQuickItem) SetScale(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setScaleEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:257
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QQuickItem) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtQuick/qquickitem.h:258
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)
func (this *QQuickItem) SetOpacity(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QQuickItem) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:261
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QQuickItem) SetVisible(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:263
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QQuickItem) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:264
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QQuickItem) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:266
// index:0
// Public Visibility=Default Availability=Available
// [1] bool smooth()
func (this *QQuickItem) Smooth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6smoothEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSmooth(_Bool)
func (this *QQuickItem) SetSmooth(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setSmoothEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:269
// index:0
// Public Visibility=Default Availability=Available
// [1] bool activeFocusOnTab()
func (this *QQuickItem) ActiveFocusOnTab() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16activeFocusOnTabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:270
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setActiveFocusOnTab(_Bool)
func (this *QQuickItem) SetActiveFocusOnTab(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem19setActiveFocusOnTabEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:272
// index:0
// Public Visibility=Default Availability=Available
// [1] bool antialiasing()
func (this *QQuickItem) Antialiasing() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12antialiasingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAntialiasing(_Bool)
func (this *QQuickItem) SetAntialiasing(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15setAntialiasingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetAntialiasing()
func (this *QQuickItem) ResetAntialiasing() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17resetAntialiasingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:276
// index:0
// Public Visibility=Default Availability=Available
// [4] QQuickItem::Flags flags()
func (this *QQuickItem) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:277
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(enum QQuickItem::Flag, _Bool)
func (this *QQuickItem) SetFlag(flag int, enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem7setFlagENS_4FlagEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flag, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:278
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QQuickItem::Flags)
func (this *QQuickItem) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setFlagsE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:280
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF boundingRect()
func (this *QQuickItem) BoundingRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12boundingRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:281
// index:0
// Public virtual Visibility=Default Availability=Available
// [32] QRectF clipRect()
func (this *QQuickItem) ClipRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8clipRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:283
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasActiveFocus()
func (this *QQuickItem) HasActiveFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14hasActiveFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:284
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus()
func (this *QQuickItem) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:285
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocus(_Bool)
func (this *QQuickItem) SetFocus(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setFocusEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:286
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFocus(_Bool, Qt::FocusReason)
func (this *QQuickItem) SetFocus_1(focus bool, reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8setFocusEbN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), focus, reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:287
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFocusScope()
func (this *QQuickItem) IsFocusScope() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12isFocusScopeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:288
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * scopedFocusItem()
func (this *QQuickItem) ScopedFocusItem() *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15scopedFocusItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:290
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QQuickItem *)
func (this *QQuickItem) IsAncestorOf(child *QQuickItem /*777 const QQuickItem **/) bool {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12isAncestorOfEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:292
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::MouseButtons acceptedMouseButtons()
func (this *QQuickItem) AcceptedMouseButtons() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem20acceptedMouseButtonsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptedMouseButtons(Qt::MouseButtons)
func (this *QQuickItem) SetAcceptedMouseButtons(buttons int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem23setAcceptedMouseButtonsE6QFlagsIN2Qt11MouseButtonEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), buttons)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:294
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptHoverEvents()
func (this *QQuickItem) AcceptHoverEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem17acceptHoverEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptHoverEvents(_Bool)
func (this *QQuickItem) SetAcceptHoverEvents(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20setAcceptHoverEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:296
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptTouchEvents()
func (this *QQuickItem) AcceptTouchEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem17acceptTouchEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptTouchEvents(_Bool)
func (this *QQuickItem) SetAcceptTouchEvents(accept bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20setAcceptTouchEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accept)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:300
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor()
func (this *QQuickItem) Cursor() *qtgui.QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)
func (this *QQuickItem) SetCursor(cursor *qtgui.QCursor) {
	var convArg0 = cursor.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:302
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()
func (this *QQuickItem) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:305
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isUnderMouse()
func (this *QQuickItem) IsUnderMouse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12isUnderMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:306
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabMouse()
func (this *QQuickItem) GrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9grabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabMouse()
func (this *QQuickItem) UngrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11ungrabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:308
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keepMouseGrab()
func (this *QQuickItem) KeepMouseGrab() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13keepMouseGrabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:309
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeepMouseGrab(_Bool)
func (this *QQuickItem) SetKeepMouseGrab(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16setKeepMouseGrabEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:310
// index:0
// Public Visibility=Default Availability=Available
// [1] bool filtersChildMouseEvents()
func (this *QQuickItem) FiltersChildMouseEvents() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem23filtersChildMouseEventsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFiltersChildMouseEvents(_Bool)
func (this *QQuickItem) SetFiltersChildMouseEvents(filter bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem26setFiltersChildMouseEventsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), filter)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabTouchPoints()
func (this *QQuickItem) UngrabTouchPoints() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17ungrabTouchPointsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:315
// index:0
// Public Visibility=Default Availability=Available
// [1] bool keepTouchGrab()
func (this *QQuickItem) KeepTouchGrab() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13keepTouchGrabEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeepTouchGrab(_Bool)
func (this *QQuickItem) SetKeepTouchGrab(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16setKeepTouchGrabEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:319
// index:0
// Public Visibility=Default Availability=Available
// [1] bool grabToImage(const QJSValue &, const QSize &)
func (this *QQuickItem) GrabToImage(callback *qtqml.QJSValue, targetSize *qtcore.QSize) bool {
	var convArg0 = callback.GetCthis()
	var convArg1 = targetSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11grabToImageERK8QJSValueRK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:322
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool contains(const QPointF &)
func (this *QQuickItem) Contains(point *qtcore.QPointF) bool {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem8containsERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:324
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform itemTransform(QQuickItem *, _Bool *)
func (this *QQuickItem) ItemTransform(arg0 *QQuickItem /*777 QQuickItem **/, arg1 unsafe.Pointer /*666*/) *qtgui.QTransform /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13itemTransformEPS_Pb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, &arg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:325
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToItem(const QQuickItem *, const QPointF &)
func (this *QQuickItem) MapToItem(item *QQuickItem /*777 const QQuickItem **/, point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem9mapToItemEPKS_RK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:326
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToScene(const QPointF &)
func (this *QQuickItem) MapToScene(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10mapToSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:327
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapToGlobal(const QPointF &)
func (this *QQuickItem) MapToGlobal(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem11mapToGlobalERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:328
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToItem(const QQuickItem *, const QRectF &)
func (this *QQuickItem) MapRectToItem(item *QQuickItem /*777 const QQuickItem **/, rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13mapRectToItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:329
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectToScene(const QRectF &)
func (this *QQuickItem) MapRectToScene(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem14mapRectToSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:330
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromItem(const QQuickItem *, const QPointF &)
func (this *QQuickItem) MapFromItem(item *QQuickItem /*777 const QQuickItem **/, point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem11mapFromItemEPKS_RK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:331
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromScene(const QPointF &)
func (this *QQuickItem) MapFromScene(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem12mapFromSceneERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:332
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF mapFromGlobal(const QPointF &)
func (this *QQuickItem) MapFromGlobal(point *qtcore.QPointF) *qtcore.QPointF /*123*/ {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem13mapFromGlobalERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:333
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromItem(const QQuickItem *, const QRectF &)
func (this *QQuickItem) MapRectFromItem(item *QQuickItem /*777 const QQuickItem **/, rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = item.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15mapRectFromItemEPKS_RK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:334
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF mapRectFromScene(const QRectF &)
func (this *QQuickItem) MapRectFromScene(rect *qtcore.QRectF) *qtcore.QRectF /*123*/ {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16mapRectFromSceneERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:336
// index:0
// Public Visibility=Default Availability=Available
// [-2] void polish()
func (this *QQuickItem) Polish() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem6polishEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:342
// index:0
// Public Visibility=Default Availability=Available
// [-2] void forceActiveFocus()
func (this *QQuickItem) ForceActiveFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16forceActiveFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:343
// index:1
// Public Visibility=Default Availability=Available
// [-2] void forceActiveFocus(Qt::FocusReason)
func (this *QQuickItem) ForceActiveFocus_1(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16forceActiveFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:344
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * nextItemInFocusChain(_Bool)
func (this *QQuickItem) NextItemInFocusChain(forward bool) *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20nextItemInFocusChainEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), forward)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:345
// index:0
// Public Visibility=Default Availability=Available
// [8] QQuickItem * childAt(qreal, qreal)
func (this *QQuickItem) ChildAt(x float64, y float64) *QQuickItem /*777 QQuickItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem7childAtEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQQuickItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:348
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QQuickItem) InputMethodQuery(query int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:358
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isTextureProvider()
func (this *QQuickItem) IsTextureProvider() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem17isTextureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:359
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSGTextureProvider * textureProvider()
func (this *QQuickItem) TextureProvider() *QSGTextureProvider /*777 QSGTextureProvider **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem15textureProviderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSGTextureProviderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:362
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QQuickItem) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:365
// index:0
// Public Visibility=Default Availability=Available
// [-2] void childrenRectChanged(const QRectF &)
func (this *QQuickItem) ChildrenRectChanged(arg0 *qtcore.QRectF) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem19childrenRectChangedERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:366
// index:0
// Public Visibility=Default Availability=Available
// [-2] void baselineOffsetChanged(qreal)
func (this *QQuickItem) BaselineOffsetChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21baselineOffsetChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:367
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(const QString &)
func (this *QQuickItem) StateChanged(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12stateChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:368
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusChanged(_Bool)
func (this *QQuickItem) FocusChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12focusChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:369
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusChanged(_Bool)
func (this *QQuickItem) ActiveFocusChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem18activeFocusChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:370
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeFocusOnTabChanged(_Bool)
func (this *QQuickItem) ActiveFocusOnTabChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem23activeFocusOnTabChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:371
// index:0
// Public Visibility=Default Availability=Available
// [-2] void parentChanged(QQuickItem *)
func (this *QQuickItem) ParentChanged(arg0 *QQuickItem /*777 QQuickItem **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13parentChangedEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:372
// index:0
// Public Visibility=Default Availability=Available
// [-2] void transformOriginChanged(enum QQuickItem::TransformOrigin)
func (this *QQuickItem) TransformOriginChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem22transformOriginChangedENS_15TransformOriginE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:373
// index:0
// Public Visibility=Default Availability=Available
// [-2] void smoothChanged(_Bool)
func (this *QQuickItem) SmoothChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13smoothChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:374
// index:0
// Public Visibility=Default Availability=Available
// [-2] void antialiasingChanged(_Bool)
func (this *QQuickItem) AntialiasingChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem19antialiasingChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:375
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clipChanged(_Bool)
func (this *QQuickItem) ClipChanged(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem11clipChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:376
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowChanged(QQuickWindow *)
func (this *QQuickItem) WindowChanged(window *QQuickWindow /*777 QQuickWindow **/) {
	var convArg0 = window.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13windowChangedEP12QQuickWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:378
// index:0
// Public Visibility=Default Availability=Available
// [-2] void childrenChanged()
func (this *QQuickItem) ChildrenChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15childrenChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityChanged()
func (this *QQuickItem) OpacityChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14opacityChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:380
// index:0
// Public Visibility=Default Availability=Available
// [-2] void enabledChanged()
func (this *QQuickItem) EnabledChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14enabledChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:381
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChanged()
func (this *QQuickItem) VisibleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14visibleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:382
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChildrenChanged()
func (this *QQuickItem) VisibleChildrenChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem22visibleChildrenChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:383
// index:0
// Public Visibility=Default Availability=Available
// [-2] void rotationChanged()
func (this *QQuickItem) RotationChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15rotationChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scaleChanged()
func (this *QQuickItem) ScaleChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12scaleChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:386
// index:0
// Public Visibility=Default Availability=Available
// [-2] void xChanged()
func (this *QQuickItem) XChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8xChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:387
// index:0
// Public Visibility=Default Availability=Available
// [-2] void yChanged()
func (this *QQuickItem) YChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8yChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:388
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widthChanged()
func (this *QQuickItem) WidthChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12widthChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:389
// index:0
// Public Visibility=Default Availability=Available
// [-2] void heightChanged()
func (this *QQuickItem) HeightChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13heightChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:390
// index:0
// Public Visibility=Default Availability=Available
// [-2] void zChanged()
func (this *QQuickItem) ZChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem8zChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:391
// index:0
// Public Visibility=Default Availability=Available
// [-2] void implicitWidthChanged()
func (this *QQuickItem) ImplicitWidthChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem20implicitWidthChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:392
// index:0
// Public Visibility=Default Availability=Available
// [-2] void implicitHeightChanged()
func (this *QQuickItem) ImplicitHeightChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21implicitHeightChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:395
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QQuickItem) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:397
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool isComponentComplete()
func (this *QQuickItem) IsComponentComplete() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem19isComponentCompleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:401
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateInputMethod(Qt::InputMethodQueries)
func (this *QQuickItem) UpdateInputMethod(queries int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17updateInputMethodE6QFlagsIN2Qt16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), queries)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:404
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool widthValid()
func (this *QQuickItem) WidthValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem10widthValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:405
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool heightValid()
func (this *QQuickItem) HeightValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QQuickItem11heightValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:406
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setImplicitSize(qreal, qreal)
func (this *QQuickItem) SetImplicitSize(arg0 float64, arg1 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15setImplicitSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:408
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void classBegin()
func (this *QQuickItem) ClassBegin() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10classBeginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:409
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void componentComplete()
func (this *QQuickItem) ComponentComplete() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17componentCompleteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:411
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QQuickItem) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:412
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QQuickItem) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:414
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QQuickItem) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:416
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QQuickItem) FocusInEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:417
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QQuickItem) FocusOutEvent(arg0 *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:418
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QQuickItem) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:419
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QQuickItem) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:420
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QQuickItem) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:421
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QQuickItem) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:422
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseUngrabEvent()
func (this *QQuickItem) MouseUngrabEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16mouseUngrabEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:423
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void touchUngrabEvent()
func (this *QQuickItem) TouchUngrabEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16touchUngrabEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:425
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QQuickItem) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:427
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void touchEvent(QTouchEvent *)
func (this *QQuickItem) TouchEvent(event *qtgui.QTouchEvent /*777 QTouchEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem10touchEventEP11QTouchEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:428
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverEnterEvent(QHoverEvent *)
func (this *QQuickItem) HoverEnterEvent(event *qtgui.QHoverEvent /*777 QHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15hoverEnterEventEP11QHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:429
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverMoveEvent(QHoverEvent *)
func (this *QQuickItem) HoverMoveEvent(event *qtgui.QHoverEvent /*777 QHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14hoverMoveEventEP11QHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:430
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hoverLeaveEvent(QHoverEvent *)
func (this *QQuickItem) HoverLeaveEvent(event *qtgui.QHoverEvent /*777 QHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15hoverLeaveEventEP11QHoverEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:432
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)
func (this *QQuickItem) DragEnterEvent(arg0 *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:433
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QQuickItem) DragMoveEvent(arg0 *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:434
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QQuickItem) DragLeaveEvent(arg0 *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:435
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QQuickItem) DropEvent(arg0 *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:437
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool childMouseEventFilter(QQuickItem *, QEvent *)
func (this *QQuickItem) ChildMouseEventFilter(arg0 *QQuickItem /*777 QQuickItem **/, arg1 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21childMouseEventFilterEPS_P6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtQuick/qquickitem.h:438
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void windowDeactivateEvent()
func (this *QQuickItem) WindowDeactivateEvent() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem21windowDeactivateEventEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:440
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void geometryChanged(const QRectF &, const QRectF &)
func (this *QQuickItem) GeometryChanged(newGeometry *qtcore.QRectF, oldGeometry *qtcore.QRectF) {
	var convArg0 = newGeometry.GetCthis()
	var convArg1 = oldGeometry.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem15geometryChangedERK6QRectFS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:444
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void releaseResources()
func (this *QQuickItem) ReleaseResources() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem16releaseResourcesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qquickitem.h:445
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void updatePolish()
func (this *QQuickItem) UpdatePolish() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QQuickItem12updatePolishEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

type QQuickItem__Flag = int

const QQuickItem__ItemClipsChildrenToShape QQuickItem__Flag = 1
const QQuickItem__ItemAcceptsInputMethod QQuickItem__Flag = 2
const QQuickItem__ItemIsFocusScope QQuickItem__Flag = 4
const QQuickItem__ItemHasContents QQuickItem__Flag = 8
const QQuickItem__ItemAcceptsDrops QQuickItem__Flag = 16

type QQuickItem__ItemChange = int

const QQuickItem__ItemChildAddedChange QQuickItem__ItemChange = 0
const QQuickItem__ItemChildRemovedChange QQuickItem__ItemChange = 1
const QQuickItem__ItemSceneChange QQuickItem__ItemChange = 2
const QQuickItem__ItemVisibleHasChanged QQuickItem__ItemChange = 3
const QQuickItem__ItemParentHasChanged QQuickItem__ItemChange = 4
const QQuickItem__ItemOpacityHasChanged QQuickItem__ItemChange = 5
const QQuickItem__ItemActiveFocusHasChanged QQuickItem__ItemChange = 6
const QQuickItem__ItemRotationHasChanged QQuickItem__ItemChange = 7
const QQuickItem__ItemAntialiasingHasChanged QQuickItem__ItemChange = 8
const QQuickItem__ItemDevicePixelRatioHasChanged QQuickItem__ItemChange = 9
const QQuickItem__ItemEnabledHasChanged QQuickItem__ItemChange = 10

type QQuickItem__TransformOrigin = int

const QQuickItem__TopLeft QQuickItem__TransformOrigin = 0
const QQuickItem__Top QQuickItem__TransformOrigin = 1
const QQuickItem__TopRight QQuickItem__TransformOrigin = 2
const QQuickItem__Left QQuickItem__TransformOrigin = 3
const QQuickItem__Center QQuickItem__TransformOrigin = 4
const QQuickItem__Right QQuickItem__TransformOrigin = 5
const QQuickItem__BottomLeft QQuickItem__TransformOrigin = 6
const QQuickItem__Bottom QQuickItem__TransformOrigin = 7
const QQuickItem__BottomRight QQuickItem__TransformOrigin = 8

//  body block end
