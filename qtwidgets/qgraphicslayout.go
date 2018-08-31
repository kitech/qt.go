package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslayout.h
// #include <qgraphicslayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 40
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

// void addChildLayoutItem(QGraphicsLayoutItem *)
func (this *QGraphicsLayout) InheritAddChildLayoutItem(f func(layoutItem *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addChildLayoutItem", f)
}

/*

 */
type QGraphicsLayout struct {
	*QGraphicsLayoutItem
}
type QGraphicsLayout_ITF interface {
	QGraphicsLayoutItem_ITF
	QGraphicsLayout_PTR() *QGraphicsLayout
}

func (ptr *QGraphicsLayout) QGraphicsLayout_PTR() *QGraphicsLayout { return ptr }

func (this *QGraphicsLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayoutItem.GetCthis()
	}
}
func (this *QGraphicsLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayoutItem = NewQGraphicsLayoutItemFromPointer(cthis)
}
func NewQGraphicsLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLayout {
	bcthis0 := NewQGraphicsLayoutItemFromPointer(cthis)
	return &QGraphicsLayout{bcthis0}
}
func (*QGraphicsLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLayout {
	return NewQGraphicsLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayout(QGraphicsLayoutItem *)

/*
Contructs a QGraphicsLayout object.

parent is passed to QGraphicsLayoutItem's constructor and the QGraphicsLayoutItem's isLayout argument is set to true.

If parent is a QGraphicsWidget the layout will be installed on that widget. (Note that installing a layout will delete the old one installed.)
*/
func (*QGraphicsLayout) NewForInherit(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLayout {
	return NewQGraphicsLayout(parent)
}
func NewQGraphicsLayout(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayout(QGraphicsLayoutItem *)

/*
Contructs a QGraphicsLayout object.

parent is passed to QGraphicsLayoutItem's constructor and the QGraphicsLayoutItem's isLayout argument is set to true.

If parent is a QGraphicsWidget the layout will be installed on that widget. (Note that installing a layout will delete the old one installed.)
*/
func (*QGraphicsLayout) NewForInherit__() *QGraphicsLayout {
	return NewQGraphicsLayout__()
}
func NewQGraphicsLayout__() *QGraphicsLayout {
	// arg: 0, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLayout()

/*

 */
func DeleteQGraphicsLayout(this *QGraphicsLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(qreal, qreal, qreal, qreal)

/*
Sets the contents margins to left, top, right and bottom. The default contents margins for toplevel layouts are style dependent (by querying the pixelMetric for QStyle::PM_LayoutLeftMargin, QStyle::PM_LayoutTopMargin, QStyle::PM_LayoutRightMargin and QStyle::PM_LayoutBottomMargin).

For sublayouts the default margins are 0.

Changing the contents margins automatically invalidates the layout.

See also invalidate().
*/
func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout18setContentsMarginsEdddd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *) const

/*
Reimplemented from QGraphicsLayoutItem::getContentsMargins().
*/
func (this *QGraphicsLayout) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout18getContentsMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activate()

/*
Activates the layout, causing all items in the layout to be immediately rearranged. This function is based on calling count() and itemAt(), and then calling setGeometry() on all items sequentially. When activated, the layout will adjust its geometry to its parent's contentsRect(). The parent will then invalidate any layout of its own.

If called in sequence or recursively, e.g., by one of the arranged items in response to being resized, this function will do nothing.

Note that the layout is free to use geometry caching to optimize this process. To forcefully invalidate any such cache, you can call invalidate() before calling activate().

See also invalidate().
*/
func (this *QGraphicsLayout) Activate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout8activateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActivated() const

/*
Returns true if the layout is currently being activated; otherwise, returns false. If the layout is being activated, this means that it is currently in the process of rearranging its items (i.e., the activate() function has been called, and has not yet returned).

See also activate() and invalidate().
*/
func (this *QGraphicsLayout) IsActivated() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout11isActivatedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Clears any cached geometry and size hint information in the layout, and posts a LayoutRequest event to the managed parent QGraphicsLayoutItem.

See also activate() and setGeometry().
*/
func (this *QGraphicsLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()

/*
Reimplemented from QGraphicsLayoutItem::updateGeometry().
*/
func (this *QGraphicsLayout) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void widgetEvent(QEvent *)

/*
This virtual event handler receives all events for the managed widget. QGraphicsLayout uses this event handler to listen for layout related events such as geometry changes, layout changes or layout direction changes.

e is a pointer to the event.

You can reimplement this event handler to track similar events for your own custom layout.

See also QGraphicsWidget::event() and QGraphicsItem::sceneEvent().
*/
func (this *QGraphicsLayout) WidgetEvent(e qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout11widgetEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int count() const

/*
This pure virtual function must be reimplemented in a subclass of QGraphicsLayout to return the number of items in the layout.

The subclass is free to decide how to store the items.

See also itemAt() and removeAt().
*/
func (this *QGraphicsLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int) const

/*
This pure virtual function must be reimplemented in a subclass of QGraphicsLayout to return a pointer to the item at index i. The reimplementation can assume that i is valid (i.e., it respects the value of count()). Together with count(), it is provided as a means of iterating over all items in a layout.

The subclass is free to decide how to store the items, and the visual arrangement does not have to be reflected through this function.

See also count() and removeAt().
*/
func (this *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QGraphicsLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:72
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)

/*
This pure virtual function must be reimplemented in a subclass of QGraphicsLayout to remove the item at index. The reimplementation can assume that index is valid (i.e., it respects the value of count()).

The implementation must ensure that the parentLayoutItem() of the removed item does not point to this layout, since the item is considered to be removed from the layout hierarchy.

If the layout is to be reused between applications, we recommend that the layout deletes the item, but the graphics view framework does not depend on this.

The subclass is free to decide how to store the items.

See also itemAt() and count().
*/
func (this *QGraphicsLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:74
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setInstantInvalidatePropagation(bool)

/*

 */
func (this *QGraphicsLayout) SetInstantInvalidatePropagation(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout31setInstantInvalidatePropagationEb", qtrt.FFI_TYPE_POINTER, enable)
	qtrt.ErrPrint(err, rv)
}
func QGraphicsLayout_SetInstantInvalidatePropagation(enable bool) {
	var nilthis *QGraphicsLayout
	nilthis.SetInstantInvalidatePropagation(enable)
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:75
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool instantInvalidatePropagation()

/*

 */
func (this *QGraphicsLayout) InstantInvalidatePropagation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout28instantInvalidatePropagationEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QGraphicsLayout_InstantInvalidatePropagation() bool {
	var nilthis *QGraphicsLayout
	rv := nilthis.InstantInvalidatePropagation()
	return rv
}

// /usr/include/qt/QtWidgets/qgraphicslayout.h:78
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addChildLayoutItem(QGraphicsLayoutItem *)

/*
This function is a convenience function provided for custom layouts, and will go through all items in the layout and reparent their graphics items to the closest QGraphicsWidget ancestor of the layout.

If layoutItem is already in a different layout, it will be removed from that layout.

If custom layouts want special behaviour they can ignore to use this function, and implement their own behaviour.

This function was introduced in  Qt 4.6.

See also graphicsItem().
*/
func (this *QGraphicsLayout) AddChildLayoutItem(layoutItem QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if layoutItem != nil && layoutItem.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = layoutItem.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QGraphicsLayout18addChildLayoutItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
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
		qtgui.KeepMe()
	}
}

//  keep block end
