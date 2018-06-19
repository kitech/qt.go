package qtwidgets

// /usr/include/qt/QtWidgets/qlayout.h
// #include <qlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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

// void widgetEvent(QEvent *)
func (this *QLayout) InheritWidgetEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "widgetEvent", f)
}

// void childEvent(QChildEvent *)
func (this *QLayout) InheritChildEvent(f func(e *qtcore.QChildEvent /*777 QChildEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "childEvent", f)
}

// void addChildLayout(QLayout *)
func (this *QLayout) InheritAddChildLayout(f func(l *QLayout /*777 QLayout **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addChildLayout", f)
}

// void addChildWidget(QWidget *)
func (this *QLayout) InheritAddChildWidget(f func(w *QWidget /*777 QWidget **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "addChildWidget", f)
}

// bool adoptLayout(QLayout *)
func (this *QLayout) InheritAdoptLayout(f func(layout *QLayout /*777 QLayout **/) bool) {
	qtrt.SetAllInheritCallback(this, "adoptLayout", f)
}

// QRect alignmentRect(const QRect &)
func (this *QLayout) InheritAlignmentRect(f func(arg0 *qtcore.QRect) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "alignmentRect", f)
}

/*

 */
type QLayout struct {
	*qtcore.QObject
	*QLayoutItem
}
type QLayout_ITF interface {
	qtcore.QObject_ITF
	QLayoutItem_ITF
	QLayout_PTR() *QLayout
}

func (ptr *QLayout) QLayout_PTR() *QLayout { return ptr }

func (this *QLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QLayout) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QLayoutItem = NewQLayoutItemFromPointer(cthis)
}
func NewQLayoutFromPointer(cthis unsafe.Pointer) *QLayout {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQLayoutItemFromPointer(cthis)
	return &QLayout{bcthis0, bcthis1}
}
func (*QLayout) NewFromPointer(cthis unsafe.Pointer) *QLayout {
	return NewQLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlayout.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLayout(QWidget *)

/*
Constructs a new top-level QLayout, with parent parent. parent may not be 0.

There can be only one top-level layout for a widget. It is returned by QWidget::layout().
*/
func NewQLayout(parent QWidget_ITF /*777 QWidget **/) *QLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayoutC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qlayout.h:81
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLayout()

/*
Constructs a new top-level QLayout, with parent parent. parent may not be 0.

There can be only one top-level layout for a widget. It is returned by QWidget::layout().
*/
func NewQLayout_1() *QLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayoutC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qlayout.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLayout()

/*

 */
func DeleteQLayout(this *QLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [4] int margin() const

/*

 */
func (this *QLayout) Margin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout6marginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing() const

/*

 */
func (this *QLayout) Spacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout7spacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMargin(int)

/*

 */
func (this *QLayout) SetMargin(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout9setMarginEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)

/*

 */
func (this *QLayout) SetSpacing(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout10setSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(int, int, int, int)

/*
Sets the left, top, right, and bottom margins to use around the layout.

By default, QLayout uses the values provided by the style. On most platforms, the margin is 11 pixels in all directions.

This function was introduced in  Qt 4.3.

See also contentsMargins(), getContentsMargins(), QStyle::pixelMetric(), PM_LayoutLeftMargin, PM_LayoutTopMargin, PM_LayoutRightMargin, and PM_LayoutBottomMargin.
*/
func (this *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout18setContentsMarginsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:91
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(const QMargins &)

/*
Sets the left, top, right, and bottom margins to use around the layout.

By default, QLayout uses the values provided by the style. On most platforms, the margin is 11 pixels in all directions.

This function was introduced in  Qt 4.3.

See also contentsMargins(), getContentsMargins(), QStyle::pixelMetric(), PM_LayoutLeftMargin, PM_LayoutTopMargin, PM_LayoutRightMargin, and PM_LayoutBottomMargin.
*/
func (this *QLayout) SetContentsMargins_1(margins qtcore.QMargins_ITF) {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg0 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout18setContentsMarginsERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getContentsMargins(int *, int *, int *, int *) const

/*
Extracts the left, top, right, and bottom margins used around the layout, and assigns them to *left, *top, *right, and *bottom (unless they are null pointers).

By default, QLayout uses the values provided by the style. On most platforms, the margin is 11 pixels in all directions.

This function was introduced in  Qt 4.3.

See also setContentsMargins(), QStyle::pixelMetric(), PM_LayoutLeftMargin, PM_LayoutTopMargin, PM_LayoutRightMargin, and PM_LayoutBottomMargin.
*/
func (this *QLayout) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout18getContentsMarginsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:93
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins contentsMargins() const

/*
Returns the margins used around the layout.

By default, QLayout uses the values provided by the style. On most platforms, the margin is 11 pixels in all directions.

This function was introduced in  Qt 4.6.

See also setContentsMargins().
*/
func (this *QLayout) ContentsMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout15contentsMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:94
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect contentsRect() const

/*
Returns the layout's geometry() rectangle, but taking into account the contents margins.

This function was introduced in  Qt 4.3.

See also setContentsMargins() and getContentsMargins().
*/
func (this *QLayout) ContentsRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout12contentsRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setAlignment(QWidget *, Qt::Alignment)

/*
Sets the alignment for widget w to alignment and returns true if w is found in this layout (not including child layouts); otherwise returns false.
*/
func (this *QLayout) SetAlignment(w QWidget_ITF /*777 QWidget **/, alignment int) bool {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout12setAlignmentEP7QWidget6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:97
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setAlignment(QLayout *, Qt::Alignment)

/*
Sets the alignment for widget w to alignment and returns true if w is found in this layout (not including child layouts); otherwise returns false.
*/
func (this *QLayout) SetAlignment_1(l QLayout_ITF /*777 QLayout **/, alignment int) bool {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLayout_PTR() != nil {
		convArg0 = l.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout12setAlignmentEPS_6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeConstraint(QLayout::SizeConstraint)

/*

 */
func (this *QLayout) SetSizeConstraint(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout17setSizeConstraintENS_14SizeConstraintE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:101
// index:0
// Public Visibility=Default Availability=Available
// [4] QLayout::SizeConstraint sizeConstraint() const

/*

 */
func (this *QLayout) SizeConstraint() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout14sizeConstraintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMenuBar(QWidget *)

/*
Tells the geometry manager to place the menu bar widget at the top of parentWidget(), outside QWidget::contentsMargins(). All child widgets are placed below the bottom edge of the menu bar.

See also menuBar().
*/
func (this *QLayout) SetMenuBar(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout10setMenuBarEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * menuBar() const

/*
Returns the menu bar set for this layout, or 0 if no menu bar is set.

See also setMenuBar().
*/
func (this *QLayout) MenuBar() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout7menuBarEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * parentWidget() const

/*
Returns the parent widget of this layout, or 0 if this layout is not installed on any widget.

If the layout is a sub-layout, this function returns the parent widget of the parent layout.

See also parent().
*/
func (this *QLayout) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Reimplemented from QLayoutItem::invalidate().
*/
func (this *QLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:108
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect geometry() const

/*
Reimplemented from QLayoutItem::geometry().

See also setGeometry().
*/
func (this *QLayout) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:109
// index:0
// Public Visibility=Default Availability=Available
// [1] bool activate()

/*
Redoes the layout for parentWidget() if necessary.

You should generally not need to call this because it is automatically called at the most appropriate times. It returns true if the layout was redone.

See also update() and QWidget::updateGeometry().
*/
func (this *QLayout) Activate() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout8activateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()

/*
Updates the layout for parentWidget().

You should generally not need to call this because it is automatically called at the most appropriate times.

See also activate() and invalidate().
*/
func (this *QLayout) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *)

/*
Adds widget w to this layout in a manner specific to the layout. This function uses addItem().
*/
func (this *QLayout) AddWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout9addWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:113
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)

/*
Implemented in subclasses to add an item. How it is added is specific to each subclass.

This function is not usually called in application code. To add a widget to a layout, use the addWidget() function; to add a child layout, use the addLayout() function provided by the relevant QLayout subclass.

Note: The ownership of item is transferred to the layout, and it's the layout's responsibility to delete it.

See also addWidget(), QBoxLayout::addLayout(), and QGridLayout::addLayout().
*/
func (this *QLayout) AddItem(arg0 QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QLayoutItem_PTR() != nil {
		convArg0 = arg0.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout7addItemEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeWidget(QWidget *)

/*
Removes the widget widget from the layout. After this call, it is the caller's responsibility to give the widget a reasonable geometry or to put the widget back into a layout or to explicitly hide it if necessary.

Note: The ownership of widget remains the same as when it was added.

See also removeItem(), QWidget::setGeometry(), and addWidget().
*/
func (this *QLayout) RemoveWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout12removeWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QLayoutItem *)

/*
Removes the layout item item from the layout. It is the caller's responsibility to delete the item.

Notice that item can be a layout (since QLayout inherits QLayoutItem).

See also removeWidget() and addItem().
*/
func (this *QLayout) RemoveItem(arg0 QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QLayoutItem_PTR() != nil {
		convArg0 = arg0.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout10removeItemEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:118
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
Reimplemented from QLayoutItem::expandingDirections().

Returns whether this layout can make use of more space than sizeHint(). A value of Qt::Vertical or Qt::Horizontal means that it wants to grow in only one dimension, whereas Qt::Vertical | Qt::Horizontal means that it wants to grow in both dimensions.

The default implementation returns Qt::Horizontal | Qt::Vertical. Subclasses reimplement it to return a meaningful value based on their child widgets's size policies.

See also sizeHint().
*/
func (this *QLayout) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:119
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Reimplemented from QLayoutItem::minimumSize().

Returns the minimum size of this layout. This is the smallest size that the layout can have while still respecting the specifications.

The returned value doesn't include the space required by QWidget::setContentsMargins() or menuBar().

The default implementation allows unlimited resizing.
*/
func (this *QLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
Reimplemented from QLayoutItem::maximumSize().

Returns the maximum size of this layout. This is the largest size that the layout can have while still respecting the specifications.

The returned value doesn't include the space required by QWidget::setContentsMargins() or menuBar().

The default implementation allows unlimited resizing.
*/
func (this *QLayout) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Reimplemented from QLayoutItem::setGeometry().

See also geometry().
*/
func (this *QLayout) SetGeometry(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:122
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int) const

/*
Must be implemented in subclasses to return the layout item at index. If there is no such item, the function must return 0. Items are numbered consecutively from 0. If an item is deleted, other items will be renumbered.

This function can be used to iterate over a layout. The following code will draw a rectangle for each layout item in the layout structure of the widget.


  static void paintLayout(QPainter *painter, QLayoutItem *item)
  {
      QLayout *layout = item->layout();
      if (layout) {
          for (int i = 0; i < layout->count(); ++i)
              paintLayout(painter, layout->itemAt(i));
      }
      painter->drawRect(item->geometry());
  }

  void MyWidget::paintEvent(QPaintEvent *)
  {
      QPainter painter(this);
      if (layout())
          paintLayout(&painter, layout());
  }



See also count() and takeAt().
*/
func (this *QLayout) ItemAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:123
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)

/*
Must be implemented in subclasses to remove the layout item at index from the layout, and return the item. If there is no such item, the function must do nothing and return 0. Items are numbered consecutively from 0. If an item is removed, other items will be renumbered.

The following code fragment shows a safe way to remove all items from a layout:


  QLayoutItem *child;
  while ((child = layout->takeAt(0)) != 0) {
      ...
      delete child;
  }



See also itemAt() and count().
*/
func (this *QLayout) TakeAt(index int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int indexOf(QWidget *) const

/*
Searches for widget widget in this layout (not including child layouts).

Returns the index of widget, or -1 if widget is not found.

The default implementation iterates over all items using itemAt()
*/
func (this *QLayout) IndexOf(arg0 QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout7indexOfEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:125
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int count() const

/*
Must be implemented in subclasses to return the number of items in the layout.

See also itemAt().
*/
func (this *QLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Reimplemented from QLayoutItem::isEmpty().
*/
func (this *QLayout) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:127
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSizePolicy::ControlTypes controlTypes() const

/*
Reimplemented from QLayoutItem::controlTypes().
*/
func (this *QLayout) ControlTypes() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout12controlTypesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayoutItem * replaceWidget(QWidget *, QWidget *, Qt::FindChildOptions)

/*
Searches for widget from and replaces it with widget to if found. Returns the layout item that contains the widget from on success. Otherwise 0 is returned. If options contains Qt::FindChildrenRecursively (the default), sub-layouts are searched for doing the replacement. Any other flag in options is ignored.

Notice that the returned item therefore might not belong to this layout, but to a sub-layout.

The returned layout item is no longer owned by the layout and should be either deleted or inserted to another layout. The widget from is no longer managed by the layout and may need to be deleted or hidden. The parent of widget from is left unchanged.

This function works for the built-in Qt layouts, but might not work for custom layouts.

This function was introduced in  Qt 5.2.

See also indexOf().
*/
func (this *QLayout) ReplaceWidget(from QWidget_ITF /*777 QWidget **/, to QWidget_ITF /*777 QWidget **/, options int) *QLayoutItem /*777 QLayoutItem **/ {
	var convArg0 unsafe.Pointer
	if from != nil && from.QWidget_PTR() != nil {
		convArg0 = from.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if to != nil && to.QWidget_PTR() != nil {
		convArg1 = to.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout13replaceWidgetEP7QWidgetS1_6QFlagsIN2Qt15FindChildOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayoutItem * replaceWidget(QWidget *, QWidget *, Qt::FindChildOptions)

/*
Searches for widget from and replaces it with widget to if found. Returns the layout item that contains the widget from on success. Otherwise 0 is returned. If options contains Qt::FindChildrenRecursively (the default), sub-layouts are searched for doing the replacement. Any other flag in options is ignored.

Notice that the returned item therefore might not belong to this layout, but to a sub-layout.

The returned layout item is no longer owned by the layout and should be either deleted or inserted to another layout. The widget from is no longer managed by the layout and may need to be deleted or hidden. The parent of widget from is left unchanged.

This function works for the built-in Qt layouts, but might not work for custom layouts.

This function was introduced in  Qt 5.2.

See also indexOf().
*/
func (this *QLayout) ReplaceWidget__(from QWidget_ITF /*777 QWidget **/, to QWidget_ITF /*777 QWidget **/) *QLayoutItem /*777 QLayoutItem **/ {
	var convArg0 unsafe.Pointer
	if from != nil && from.QWidget_PTR() != nil {
		convArg0 = from.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if to != nil && to.QWidget_PTR() != nil {
		convArg1 = to.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::FindChildOptions=Elaborated, Qt::FindChildOptions=Typedef, QFlags<Qt::FindChildOption>, Unexposed
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout13replaceWidgetEP7QWidgetS1_6QFlagsIN2Qt15FindChildOptionEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int totalHeightForWidth(int) const

/*

 */
func (this *QLayout) TotalHeightForWidth(w int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout19totalHeightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlayout.h:133
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize totalMinimumSize() const

/*

 */
func (this *QLayout) TotalMinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout16totalMinimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize totalMaximumSize() const

/*

 */
func (this *QLayout) TotalMaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout16totalMaximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize totalSizeHint() const

/*

 */
func (this *QLayout) TotalSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout13totalSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlayout.h:136
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayout * layout()

/*
Reimplemented from QLayoutItem::layout().
*/
func (this *QLayout) Layout() *QLayout /*777 QLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlayout.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(bool)

/*
Enables this layout if enable is true, otherwise disables it.

An enabled layout adjusts dynamically to changes; a disabled layout acts as if it did not exist.

By default all layouts are enabled.

See also isEnabled().
*/
func (this *QLayout) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:139
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled() const

/*
Returns true if the layout is enabled; otherwise returns false.

See also setEnabled().
*/
func (this *QLayout) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:142
// index:0
// Public static Visibility=Default Availability=Available
// [8] QSize closestAcceptableSize(const QWidget *, const QSize &)

/*
Returns a size that satisfies all size constraints on widget, including heightForWidth() and that is as close as possible to size.
*/
func (this *QLayout) ClosestAcceptableSize(w QWidget_ITF /*777 const QWidget **/, s qtcore.QSize_ITF) *qtcore.QSize /*123*/ {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if s != nil && s.QSize_PTR() != nil {
		convArg1 = s.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout21closestAcceptableSizeEPK7QWidgetRK5QSize", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}
func QLayout_ClosestAcceptableSize(w QWidget_ITF /*777 const QWidget **/, s qtcore.QSize_ITF) *qtcore.QSize /*123*/ {
	var nilthis *QLayout
	rv := nilthis.ClosestAcceptableSize(w, s)
	return rv
}

// /usr/include/qt/QtWidgets/qlayout.h:145
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void widgetEvent(QEvent *)

/*

 */
func (this *QLayout) WidgetEvent(arg0 qtcore.QEvent_ITF /*777 QEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QEvent_PTR() != nil {
		convArg0 = arg0.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout11widgetEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:146
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void childEvent(QChildEvent *)

/*
Reimplemented from QObject::childEvent().
*/
func (this *QLayout) ChildEvent(e qtcore.QChildEvent_ITF /*777 QChildEvent **/) {
	var convArg0 unsafe.Pointer
	if e != nil && e.QChildEvent_PTR() != nil {
		convArg0 = e.QChildEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout10childEventEP11QChildEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:147
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addChildLayout(QLayout *)

/*
This function is called from addLayout() or insertLayout() functions in subclasses to add layout l as a sub-layout.

The only scenario in which you need to call it directly is if you implement a custom layout that supports nested layouts.

See also QBoxLayout::addLayout(), QBoxLayout::insertLayout(), and QGridLayout::addLayout().
*/
func (this *QLayout) AddChildLayout(l QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLayout_PTR() != nil {
		convArg0 = l.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout14addChildLayoutEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:148
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void addChildWidget(QWidget *)

/*
This function is called from addWidget() functions in subclasses to add w as a managed widget of a layout.

If w is already managed by a layout, this function will give a warning and remove w from that layout. This function must therefore be called before adding w to the layout's data structure.
*/
func (this *QLayout) AddChildWidget(w QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout14addChildWidgetEP7QWidget", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlayout.h:149
// index:0
// Protected Visibility=Default Availability=Available
// [1] bool adoptLayout(QLayout *)

/*

 */
func (this *QLayout) AdoptLayout(layout QLayout_ITF /*777 QLayout **/) bool {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QLayout11adoptLayoutEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlayout.h:151
// index:0
// Protected Visibility=Default Availability=Available
// [16] QRect alignmentRect(const QRect &) const

/*
Returns the rectangle that should be covered when the geometry of this layout is set to r, provided that this layout supports setAlignment().

The result is derived from sizeHint() and expanding(). It is never larger than r.
*/
func (this *QLayout) AlignmentRect(arg0 qtcore.QRect_ITF) *qtcore.QRect /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QLayout13alignmentRectERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

/*
The possible values are:



See also setSizeConstraint().

*/
type QLayout__SizeConstraint = int

// The main widget's minimum size is set to minimumSize(), unless the widget already has a minimum size.
const QLayout__SetDefaultConstraint QLayout__SizeConstraint = 0

// The widget is not constrained.
const QLayout__SetNoConstraint QLayout__SizeConstraint = 1

// The main widget's minimum size is set to minimumSize(); it cannot be smaller.
const QLayout__SetMinimumSize QLayout__SizeConstraint = 2

// The main widget's size is set to sizeHint(); it cannot be resized at all.
const QLayout__SetFixedSize QLayout__SizeConstraint = 3

// The main widget's maximum size is set to maximumSize(); it cannot be larger.
const QLayout__SetMaximumSize QLayout__SizeConstraint = 4

// The main widget's minimum size is set to minimumSize() and its maximum size is set to maximumSize().
const QLayout__SetMinAndMaxSize QLayout__SizeConstraint = 5

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
