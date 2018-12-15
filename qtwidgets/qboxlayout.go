package qtwidgets

// /usr/include/qt/QtWidgets/qboxlayout.h
// #include <qboxlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 45
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

/*

 */
type QBoxLayout struct {
	*QLayout
}
type QBoxLayout_ITF interface {
	QLayout_ITF
	QBoxLayout_PTR() *QBoxLayout
}

func (ptr *QBoxLayout) QBoxLayout_PTR() *QBoxLayout { return ptr }

func (this *QBoxLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QLayout.GetCthis()
	}
}
func (this *QBoxLayout) SetCthis(cthis unsafe.Pointer) {
	this.QLayout = NewQLayoutFromPointer(cthis)
}
func NewQBoxLayoutFromPointer(cthis unsafe.Pointer) *QBoxLayout {
	bcthis0 := NewQLayoutFromPointer(cthis)
	return &QBoxLayout{bcthis0}
}
func (*QBoxLayout) NewFromPointer(cthis unsafe.Pointer) *QBoxLayout {
	return NewQBoxLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QBoxLayout) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBoxLayout(QBoxLayout::Direction, QWidget *)

/*
Constructs a new QBoxLayout with direction dir and parent widget parent.

See also direction().
*/
func (*QBoxLayout) NewForInherit(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QBoxLayout {
	return NewQBoxLayout(arg0, parent)
}
func NewQBoxLayout(arg0 int, parent QWidget_ITF /*777 QWidget **/) *QBoxLayout {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QBoxLayout(QBoxLayout::Direction, QWidget *)

/*
Constructs a new QBoxLayout with direction dir and parent widget parent.

See also direction().
*/
func (*QBoxLayout) NewForInheritp(arg0 int) *QBoxLayout {
	return NewQBoxLayoutp(arg0)
}
func NewQBoxLayoutp(arg0 int) *QBoxLayout {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayoutC2ENS_9DirectionEP7QWidget", qtrt.FFI_TYPE_POINTER, arg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQBoxLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QBoxLayout")
	return gothis
}

// /usr/include/qt/QtWidgets/qboxlayout.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QBoxLayout()

/*

 */
func DeleteQBoxLayout(this *QBoxLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:68
// index:0
// Public Visibility=Default Availability=Available
// [4] QBoxLayout::Direction direction() const

/*
Returns the direction of the box. addWidget() and addSpacing() work in this direction; the stretch stretches in this direction.

See also setDirection(), QBoxLayout::Direction, addWidget(), and addSpacing().
*/
func (this *QBoxLayout) Direction() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout9directionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDirection(QBoxLayout::Direction)

/*
Sets the direction of this layout to direction.

See also direction().
*/
func (this *QBoxLayout) SetDirection(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout12setDirectionENS_9DirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addSpacing(int)

/*
Adds a non-stretchable space (a QSpacerItem) with size size to the end of this box layout. QBoxLayout provides default margin and spacing. This function adds additional space.

See also insertSpacing(), addItem(), and QSpacerItem.
*/
func (this *QBoxLayout) AddSpacing(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10addSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStretch(int)

/*
Adds a stretchable space (a QSpacerItem) with zero minimum size and stretch factor stretch to the end of this box layout.

See also insertStretch(), addItem(), and QSpacerItem.
*/
func (this *QBoxLayout) AddStretch(stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10addStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStretch(int)

/*
Adds a stretchable space (a QSpacerItem) with zero minimum size and stretch factor stretch to the end of this box layout.

See also insertStretch(), addItem(), and QSpacerItem.
*/
func (this *QBoxLayout) AddStretchp() {
	// arg: 0, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10addStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addSpacerItem(QSpacerItem *)

/*
Adds spacerItem to the end of this box layout.

This function was introduced in  Qt 4.4.

See also addSpacing() and addStretch().
*/
func (this *QBoxLayout) AddSpacerItem(spacerItem QSpacerItem_ITF /*777 QSpacerItem **/) {
	var convArg0 unsafe.Pointer
	if spacerItem != nil && spacerItem.QSpacerItem_PTR() != nil {
		convArg0 = spacerItem.QSpacerItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout13addSpacerItemEP11QSpacerItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, Qt::Alignment)

/*
Adds widget to the end of this box layout, with a stretch factor of stretch and alignment alignment.

The stretch factor applies only in the direction of the QBoxLayout, and is relative to the other boxes and widgets in this QBoxLayout. Widgets and boxes with higher stretch factors grow more.

If the stretch factor is 0 and nothing else in the QBoxLayout has a stretch factor greater than zero, the space is distributed according to the QWidget:sizePolicy() of each widget that's involved.

The alignment is specified by alignment. The default alignment is 0, which means that the widget fills the entire cell.

See also insertWidget(), addItem(), addLayout(), addStretch(), addSpacing(), and addStrut().
*/
func (this *QBoxLayout) AddWidget(arg0 QWidget_ITF /*777 QWidget **/, stretch int, alignment int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout9addWidgetEP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, Qt::Alignment)

/*
Adds widget to the end of this box layout, with a stretch factor of stretch and alignment alignment.

The stretch factor applies only in the direction of the QBoxLayout, and is relative to the other boxes and widgets in this QBoxLayout. Widgets and boxes with higher stretch factors grow more.

If the stretch factor is 0 and nothing else in the QBoxLayout has a stretch factor greater than zero, the space is distributed according to the QWidget:sizePolicy() of each widget that's involved.

The alignment is specified by alignment. The default alignment is 0, which means that the widget fills the entire cell.

See also insertWidget(), addItem(), addLayout(), addStretch(), addSpacing(), and addStrut().
*/
func (this *QBoxLayout) AddWidgetp(arg0 QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	// arg: 2, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout9addWidgetEP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addWidget(QWidget *, int, Qt::Alignment)

/*
Adds widget to the end of this box layout, with a stretch factor of stretch and alignment alignment.

The stretch factor applies only in the direction of the QBoxLayout, and is relative to the other boxes and widgets in this QBoxLayout. Widgets and boxes with higher stretch factors grow more.

If the stretch factor is 0 and nothing else in the QBoxLayout has a stretch factor greater than zero, the space is distributed according to the QWidget:sizePolicy() of each widget that's involved.

The alignment is specified by alignment. The default alignment is 0, which means that the widget fills the entire cell.

See also insertWidget(), addItem(), addLayout(), addStretch(), addSpacing(), and addStrut().
*/
func (this *QBoxLayout) AddWidgetp1(arg0 QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWidget_PTR() != nil {
		convArg0 = arg0.QWidget_PTR().GetCthis()
	}
	// arg: 2, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout9addWidgetEP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int)

/*
Adds layout to the end of the box, with serial stretch factor stretch.

See also insertLayout(), addItem(), and addWidget().
*/
func (this *QBoxLayout) AddLayout(layout QLayout_ITF /*777 QLayout **/, stretch int) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout9addLayoutEP7QLayouti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addLayout(QLayout *, int)

/*
Adds layout to the end of the box, with serial stretch factor stretch.

See also insertLayout(), addItem(), and addWidget().
*/
func (this *QBoxLayout) AddLayoutp(layout QLayout_ITF /*777 QLayout **/) {
	var convArg0 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg0 = layout.QLayout_PTR().GetCthis()
	}
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout9addLayoutEP7QLayouti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addStrut(int)

/*
Limits the perpendicular dimension of the box (e.g. height if the box is LeftToRight) to a minimum of size. Other constraints may increase the limit.

See also addItem().
*/
func (this *QBoxLayout) AddStrut(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout8addStrutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:77
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void addItem(QLayoutItem *)

/*
Reimplemented from QLayout::addItem().
*/
func (this *QBoxLayout) AddItem(arg0 QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QLayoutItem_PTR() != nil {
		convArg0 = arg0.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout7addItemEP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertSpacing(int, int)

/*
Inserts a non-stretchable space (a QSpacerItem) at position index, with size size. If index is negative the space is added at the end.

The box layout has default margin and spacing. This function adds additional space.

See also addSpacing(), insertItem(), and QSpacerItem.
*/
func (this *QBoxLayout) InsertSpacing(index int, size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout13insertSpacingEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertStretch(int, int)

/*
Inserts a stretchable space (a QSpacerItem) at position index, with zero minimum size and stretch factor stretch. If index is negative the space is added at the end.

See also addStretch(), insertItem(), and QSpacerItem.
*/
func (this *QBoxLayout) InsertStretch(index int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout13insertStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertStretch(int, int)

/*
Inserts a stretchable space (a QSpacerItem) at position index, with zero minimum size and stretch factor stretch. If index is negative the space is added at the end.

See also addStretch(), insertItem(), and QSpacerItem.
*/
func (this *QBoxLayout) InsertStretchp(index int) {
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout13insertStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertSpacerItem(int, QSpacerItem *)

/*
Inserts spacerItem at position index, with zero minimum size and stretch factor. If index is negative the space is added at the end.

This function was introduced in  Qt 4.4.

See also addSpacerItem(), insertStretch(), and insertSpacing().
*/
func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem QSpacerItem_ITF /*777 QSpacerItem **/) {
	var convArg1 unsafe.Pointer
	if spacerItem != nil && spacerItem.QSpacerItem_PTR() != nil {
		convArg1 = spacerItem.QSpacerItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout16insertSpacerItemEiP11QSpacerItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertWidget(int, QWidget *, int, Qt::Alignment)

/*
Inserts widget at position index, with stretch factor stretch and alignment alignment. If index is negative, the widget is added at the end.

The stretch factor applies only in the direction of the QBoxLayout, and is relative to the other boxes and widgets in this QBoxLayout. Widgets and boxes with higher stretch factors grow more.

If the stretch factor is 0 and nothing else in the QBoxLayout has a stretch factor greater than zero, the space is distributed according to the QWidget:sizePolicy() of each widget that's involved.

The alignment is specified by alignment. The default alignment is 0, which means that the widget fills the entire cell.

See also addWidget() and insertItem().
*/
func (this *QBoxLayout) InsertWidget(index int, widget QWidget_ITF /*777 QWidget **/, stretch int, alignment int) {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout12insertWidgetEiP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertWidget(int, QWidget *, int, Qt::Alignment)

/*
Inserts widget at position index, with stretch factor stretch and alignment alignment. If index is negative, the widget is added at the end.

The stretch factor applies only in the direction of the QBoxLayout, and is relative to the other boxes and widgets in this QBoxLayout. Widgets and boxes with higher stretch factors grow more.

If the stretch factor is 0 and nothing else in the QBoxLayout has a stretch factor greater than zero, the space is distributed according to the QWidget:sizePolicy() of each widget that's involved.

The alignment is specified by alignment. The default alignment is 0, which means that the widget fills the entire cell.

See also addWidget() and insertItem().
*/
func (this *QBoxLayout) InsertWidgetp(index int, widget QWidget_ITF /*777 QWidget **/) {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	stretch := int(0)
	// arg: 3, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout12insertWidgetEiP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertWidget(int, QWidget *, int, Qt::Alignment)

/*
Inserts widget at position index, with stretch factor stretch and alignment alignment. If index is negative, the widget is added at the end.

The stretch factor applies only in the direction of the QBoxLayout, and is relative to the other boxes and widgets in this QBoxLayout. Widgets and boxes with higher stretch factors grow more.

If the stretch factor is 0 and nothing else in the QBoxLayout has a stretch factor greater than zero, the space is distributed according to the QWidget:sizePolicy() of each widget that's involved.

The alignment is specified by alignment. The default alignment is 0, which means that the widget fills the entire cell.

See also addWidget() and insertItem().
*/
func (this *QBoxLayout) InsertWidgetp1(index int, widget QWidget_ITF /*777 QWidget **/, stretch int) {
	var convArg1 unsafe.Pointer
	if widget != nil && widget.QWidget_PTR() != nil {
		convArg1 = widget.QWidget_PTR().GetCthis()
	}
	// arg: 3, Qt::Alignment=Elaborated, Qt::Alignment=Typedef, QFlags<Qt::AlignmentFlag>, Unexposed
	alignment := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout12insertWidgetEiP7QWidgeti6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertLayout(int, QLayout *, int)

/*
Inserts layout at position index, with stretch factor stretch. If index is negative, the layout is added at the end.

layout becomes a child of the box layout.

See also addLayout() and insertItem().
*/
func (this *QBoxLayout) InsertLayout(index int, layout QLayout_ITF /*777 QLayout **/, stretch int) {
	var convArg1 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg1 = layout.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout12insertLayoutEiP7QLayouti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertLayout(int, QLayout *, int)

/*
Inserts layout at position index, with stretch factor stretch. If index is negative, the layout is added at the end.

layout becomes a child of the box layout.

See also addLayout() and insertItem().
*/
func (this *QBoxLayout) InsertLayoutp(index int, layout QLayout_ITF /*777 QLayout **/) {
	var convArg1 unsafe.Pointer
	if layout != nil && layout.QLayout_PTR() != nil {
		convArg1 = layout.QLayout_PTR().GetCthis()
	}
	// arg: 2, int=Int, =Invalid, , Invalid
	stretch := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout12insertLayoutEiP7QLayouti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, QLayoutItem *)

/*
Inserts item into this box layout at position index. If index is negative, the item is added at the end.

See also addItem(), insertWidget(), insertLayout(), insertStretch(), and insertSpacing().
*/
func (this *QBoxLayout) InsertItem(index int, arg1 QLayoutItem_ITF /*777 QLayoutItem **/) {
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QLayoutItem_PTR() != nil {
		convArg1 = arg1.QLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10insertItemEiP11QLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int spacing() const

/*
Reimplements QLayout::spacing(). If the spacing property is valid, that value is returned. Otherwise, a value for the spacing property is computed and returned. Since layout spacing in a widget is style dependent, if the parent is a widget, it queries the style for the (horizontal or vertical) spacing of the layout. Otherwise, the parent is a layout, and it queries the parent layout for the spacing().

See also QLayout::spacing() and setSpacing().
*/
func (this *QBoxLayout) Spacing() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout7spacingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(int)

/*
Reimplements QLayout::setSpacing(). Sets the spacing property to spacing.

See also QLayout::setSpacing() and spacing().
*/
func (this *QBoxLayout) SetSpacing(spacing int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10setSpacingEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:89
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setStretchFactor(QWidget *, int)

/*
Sets the stretch factor for widget to stretch and returns true if widget is found in this layout (not including child layouts); otherwise returns false.

See also setAlignment().
*/
func (this *QBoxLayout) SetStretchFactor(w QWidget_ITF /*777 QWidget **/, stretch int) bool {
	var convArg0 unsafe.Pointer
	if w != nil && w.QWidget_PTR() != nil {
		convArg0 = w.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QWidgeti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:90
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setStretchFactor(QLayout *, int)

/*
Sets the stretch factor for widget to stretch and returns true if widget is found in this layout (not including child layouts); otherwise returns false.

See also setAlignment().
*/
func (this *QBoxLayout) SetStretchFactor1(l QLayout_ITF /*777 QLayout **/, stretch int) bool {
	var convArg0 unsafe.Pointer
	if l != nil && l.QLayout_PTR() != nil {
		convArg0 = l.QLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout16setStretchFactorEP7QLayouti", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretch(int, int)

/*
Sets the stretch factor at position index. to stretch.

This function was introduced in  Qt 4.5.

See also stretch().
*/
func (this *QBoxLayout) SetStretch(index int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10setStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretch(int) const

/*
Returns the stretch factor at position index.

This function was introduced in  Qt 4.5.

See also setStretch().
*/
func (this *QBoxLayout) Stretch(index int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout7stretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:94
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QLayoutItem::sizeHint().
*/
func (this *QBoxLayout) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:95
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSize() const

/*
Reimplemented from QLayout::minimumSize().
*/
func (this *QBoxLayout) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize maximumSize() const

/*
Reimplemented from QLayout::maximumSize().
*/
func (this *QBoxLayout) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qboxlayout.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Reimplemented from QLayoutItem::hasHeightForWidth().
*/
func (this *QBoxLayout) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qboxlayout.h:99
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int) const

/*
Reimplemented from QLayoutItem::heightForWidth().
*/
func (this *QBoxLayout) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int minimumHeightForWidth(int) const

/*
Reimplemented from QLayoutItem::minimumHeightForWidth().
*/
func (this *QBoxLayout) MinimumHeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout21minimumHeightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:102
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
Reimplemented from QLayout::expandingDirections().
*/
func (this *QBoxLayout) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Reimplemented from QLayout::invalidate().

Resets cached information.
*/
func (this *QBoxLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qboxlayout.h:104
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * itemAt(int) const

/*
Reimplemented from QLayout::itemAt().
*/
func (this *QBoxLayout) ItemAt(arg0 int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qboxlayout.h:105
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QLayoutItem * takeAt(int)

/*
Reimplemented from QLayout::takeAt().
*/
func (this *QBoxLayout) TakeAt(arg0 int) *QLayoutItem /*777 QLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout6takeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qboxlayout.h:106
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const

/*
Reimplemented from QLayout::count().
*/
func (this *QBoxLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QBoxLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qboxlayout.h:107
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)

/*
Reimplemented from QLayout::setGeometry().
*/
func (this *QBoxLayout) SetGeometry(arg0 qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QRect_PTR() != nil {
		convArg0 = arg0.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QBoxLayout11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This type is used to determine the direction of a box layout.


*/
type QBoxLayout__Direction = int

// Horizontal from left to right.
const QBoxLayout__LeftToRight QBoxLayout__Direction = 0

// Horizontal from right to left.
const QBoxLayout__RightToLeft QBoxLayout__Direction = 1

// Vertical from top to bottom.
const QBoxLayout__TopToBottom QBoxLayout__Direction = 2

// Vertical from bottom to top.
const QBoxLayout__BottomToTop QBoxLayout__Direction = 3

//
const QBoxLayout__Down QBoxLayout__Direction = 2

//
const QBoxLayout__Up QBoxLayout__Direction = 3

func (this *QBoxLayout) DirectionItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QBoxLayout_DirectionItemName(val int) string {
	var nilthis *QBoxLayout
	return nilthis.DirectionItemName(val)
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
