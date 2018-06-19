package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h
// #include <qgraphicslinearlayout.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 26
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
type QGraphicsLinearLayout struct {
	*QGraphicsLayout
}
type QGraphicsLinearLayout_ITF interface {
	QGraphicsLayout_ITF
	QGraphicsLinearLayout_PTR() *QGraphicsLinearLayout
}

func (ptr *QGraphicsLinearLayout) QGraphicsLinearLayout_PTR() *QGraphicsLinearLayout { return ptr }

func (this *QGraphicsLinearLayout) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayout.GetCthis()
	}
}
func (this *QGraphicsLinearLayout) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayout = NewQGraphicsLayoutFromPointer(cthis)
}
func NewQGraphicsLinearLayoutFromPointer(cthis unsafe.Pointer) *QGraphicsLinearLayout {
	bcthis0 := NewQGraphicsLayoutFromPointer(cthis)
	return &QGraphicsLinearLayout{bcthis0}
}
func (*QGraphicsLinearLayout) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLinearLayout {
	return NewQGraphicsLinearLayoutFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLinearLayout(QGraphicsLayoutItem *)

/*
Constructs a QGraphicsLinearLayout instance using Qt::Horizontal orientation. parent is passed to QGraphicsLayout's constructor.
*/
func NewQGraphicsLinearLayout(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLinearLayout {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLinearLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:56
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLinearLayout(QGraphicsLayoutItem *)

/*
Constructs a QGraphicsLinearLayout instance using Qt::Horizontal orientation. parent is passed to QGraphicsLayout's constructor.
*/
func NewQGraphicsLinearLayout__() *QGraphicsLinearLayout {
	// arg: 0, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLinearLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLinearLayout(Qt::Orientation, QGraphicsLayoutItem *)

/*
Constructs a QGraphicsLinearLayout instance using Qt::Horizontal orientation. parent is passed to QGraphicsLayout's constructor.
*/
func NewQGraphicsLinearLayout_1(orientation int, parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLinearLayout {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg1 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EN2Qt11OrientationEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLinearLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:57
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLinearLayout(Qt::Orientation, QGraphicsLayoutItem *)

/*
Constructs a QGraphicsLinearLayout instance using Qt::Horizontal orientation. parent is passed to QGraphicsLayout's constructor.
*/
func NewQGraphicsLinearLayout_1_(orientation int) *QGraphicsLinearLayout {
	// arg: 1, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutC2EN2Qt11OrientationEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, orientation, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLinearLayoutFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLinearLayout)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLinearLayout()

/*

 */
func DeleteQGraphicsLinearLayout(this *QGraphicsLinearLayout) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayoutD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientation(Qt::Orientation)

/*
Change the layout orientation to orientation. Changing the layout orientation will automatically invalidate the layout.

See also orientation().
*/
func (this *QGraphicsLinearLayout) SetOrientation(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setOrientationEN2Qt11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:61
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Orientation orientation() const

/*
Returns the layout orientation.

See also setOrientation().
*/
func (this *QGraphicsLinearLayout) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:63
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addItem(QGraphicsLayoutItem *)

/*
This convenience function is equivalent to calling insertItem(-1, item).
*/
func (this *QGraphicsLinearLayout) AddItem(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout7addItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addStretch(int)

/*
This convenience function is equivalent to calling insertStretch(-1, stretch).
*/
func (this *QGraphicsLinearLayout) AddStretch(stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10addStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:64
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void addStretch(int)

/*
This convenience function is equivalent to calling insertStretch(-1, stretch).
*/
func (this *QGraphicsLinearLayout) AddStretch__() {
	// arg: 0, int=Int, =Invalid, , Invalid
	stretch := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10addStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertItem(int, QGraphicsLayoutItem *)

/*
Inserts item into the layout at index, or before any item that is currently at index.

See also addItem(), itemAt(), insertStretch(), and setItemSpacing().
*/
func (this *QGraphicsLinearLayout) InsertItem(index int, item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg1 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg1 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10insertItemEiP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertStretch(int, int)

/*
Inserts a stretch of stretch at index, or before any item that is currently at index.

See also addStretch(), setStretchFactor(), setItemSpacing(), and insertItem().
*/
func (this *QGraphicsLinearLayout) InsertStretch(index int, stretch int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout13insertStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertStretch(int, int)

/*
Inserts a stretch of stretch at index, or before any item that is currently at index.

See also addStretch(), setStretchFactor(), setItemSpacing(), and insertItem().
*/
func (this *QGraphicsLinearLayout) InsertStretch__(index int) {
	// arg: 1, int=Int, =Invalid, , Invalid
	stretch := int(1)
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout13insertStretchEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:69
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeItem(QGraphicsLayoutItem *)

/*
Removes item from the layout without destroying it. Ownership of item is transferred to the caller.

See also removeAt() and insertItem().
*/
func (this *QGraphicsLinearLayout) RemoveItem(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10removeItemEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:70
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void removeAt(int)

/*
Reimplemented from QGraphicsLayout::removeAt().

Removes the item at index without destroying it. Ownership of the item is transferred to the caller.

See also removeItem() and insertItem().
*/
func (this *QGraphicsLinearLayout) RemoveAt(index int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout8removeAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpacing(qreal)

/*
Sets the layout's spacing to spacing. Spacing refers to the vertical and horizontal distances between items.

See also spacing(), setItemSpacing(), setStretchFactor(), and QGraphicsGridLayout::setSpacing().
*/
func (this *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10setSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal spacing() const

/*
Returns the layout's spacing. Spacing refers to the vertical and horizontal distances between items.

See also setSpacing().
*/
func (this *QGraphicsLinearLayout) Spacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout7spacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setItemSpacing(int, qreal)

/*
Sets the spacing after item at index to spacing.

See also itemSpacing().
*/
func (this *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout14setItemSpacingEid", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index, spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal itemSpacing(int) const

/*
Returns the spacing after item at index.

See also setItemSpacing().
*/
func (this *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout11itemSpacingEi", qtrt.FFI_TYPE_DOUBLE, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStretchFactor(QGraphicsLayoutItem *, int)

/*
Sets the stretch factor for item to stretch. If an item's stretch factor changes, this function will invalidate the layout.

Setting stretch to 0 removes the stretch factor from the item, and is effectively equivalent to setting stretch to 1.

See also stretchFactor().
*/
func (this *QGraphicsLinearLayout) SetStretchFactor(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, stretch int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout16setStretchFactorEP19QGraphicsLayoutItemi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, stretch)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:78
// index:0
// Public Visibility=Default Availability=Available
// [4] int stretchFactor(QGraphicsLayoutItem *) const

/*
Returns the stretch factor for item. The default stretch factor is 0, meaning that the item has no assigned stretch factor.

See also setStretchFactor().
*/
func (this *QGraphicsLinearLayout) StretchFactor(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout13stretchFactorEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlignment(QGraphicsLayoutItem *, Qt::Alignment)

/*
Sets the alignment of item to alignment. If item's alignment changes, the layout is automatically invalidated.

See also alignment() and invalidate().
*/
func (this *QGraphicsLinearLayout) SetAlignment(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, alignment int) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout12setAlignmentEP19QGraphicsLayoutItem6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::Alignment alignment(QGraphicsLayoutItem *) const

/*
Returns the alignment for item. The default alignment is Qt::AlignTop | Qt::AlignLeft.

The alignment decides how the item is positioned within its assigned space in the case where there's more space available in the layout than the widgets can occupy.

See also setAlignment().
*/
func (this *QGraphicsLinearLayout) Alignment(item QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) int {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = item.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout9alignmentEP19QGraphicsLayoutItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:83
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)

/*
Reimplemented from QGraphicsLayoutItem::setGeometry().
*/
func (this *QGraphicsLinearLayout) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:85
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int count() const

/*
Reimplemented from QGraphicsLayout::count().
*/
func (this *QGraphicsLinearLayout) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * itemAt(int) const

/*
Reimplemented from QGraphicsLayout::itemAt().

When iterating from 0 and up, it will return the items in the visual arranged order.
*/
func (this *QGraphicsLinearLayout) ItemAt(index int) *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout6itemAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), index)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:88
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void invalidate()

/*
Reimplemented from QGraphicsLayout::invalidate().
*/
func (this *QGraphicsLinearLayout) Invalidate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGraphicsLinearLayout10invalidateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsLinearLayout) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
Reimplemented from QGraphicsLayoutItem::sizeHint().
*/
func (this *QGraphicsLinearLayout) SizeHint__(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dump(int) const

/*

 */
func (this *QGraphicsLinearLayout) Dump(indent int) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout4dumpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), indent)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslinearlayout.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void dump(int) const

/*

 */
func (this *QGraphicsLinearLayout) Dump__() {
	// arg: 0, int=Int, =Invalid, , Invalid
	indent := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGraphicsLinearLayout4dumpEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), indent)
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
