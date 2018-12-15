package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h
// #include <qgraphicslayoutitem.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
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

// void setGraphicsItem(QGraphicsItem *)
func (this *QGraphicsLayoutItem) InheritSetGraphicsItem(f func(item *QGraphicsItem /*777 QGraphicsItem **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setGraphicsItem", f)
}

// void setOwnedByLayout(bool)
func (this *QGraphicsLayoutItem) InheritSetOwnedByLayout(f func(ownedByLayout bool) /*void*/) {
	qtrt.SetAllInheritCallback(this, "setOwnedByLayout", f)
}

// QSizeF sizeHint(Qt::SizeHint, const QSizeF &)
func (this *QGraphicsLayoutItem) InheritSizeHint(f func(which int, constraint *qtcore.QSizeF) unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "sizeHint", f)
}

/*

 */
type QGraphicsLayoutItem struct {
	*qtrt.CObject
}
type QGraphicsLayoutItem_ITF interface {
	QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem
}

func (ptr *QGraphicsLayoutItem) QGraphicsLayoutItem_PTR() *QGraphicsLayoutItem { return ptr }

func (this *QGraphicsLayoutItem) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QGraphicsLayoutItem) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQGraphicsLayoutItemFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return &QGraphicsLayoutItem{&qtrt.CObject{cthis}}
}
func (*QGraphicsLayoutItem) NewFromPointer(cthis unsafe.Pointer) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayoutItem(QGraphicsLayoutItem *, bool)

/*
Constructs the QGraphicsLayoutItem object. parent becomes the object's parent. If isLayout is true the item is a layout, otherwise isLayout is false.
*/
func (*QGraphicsLayoutItem) NewForInherit(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, isLayout bool) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItem(parent, isLayout)
}
func NewQGraphicsLayoutItem(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/, isLayout bool) *QGraphicsLayoutItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC2EPS_b", qtrt.FFI_TYPE_POINTER, convArg0, isLayout)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayoutItem(QGraphicsLayoutItem *, bool)

/*
Constructs the QGraphicsLayoutItem object. parent becomes the object's parent. If isLayout is true the item is a layout, otherwise isLayout is false.
*/
func (*QGraphicsLayoutItem) NewForInheritp() *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemp()
}
func NewQGraphicsLayoutItemp() *QGraphicsLayoutItem {
	// arg: 0, QGraphicsLayoutItem *=Pointer, QGraphicsLayoutItem=Record, , Invalid
	var convArg0 unsafe.Pointer
	// arg: 1, bool=Bool, =Invalid, , Invalid
	isLayout := false
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC2EPS_b", qtrt.FFI_TYPE_POINTER, convArg0, isLayout)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGraphicsLayoutItem(QGraphicsLayoutItem *, bool)

/*
Constructs the QGraphicsLayoutItem object. parent becomes the object's parent. If isLayout is true the item is a layout, otherwise isLayout is false.
*/
func (*QGraphicsLayoutItem) NewForInheritp1(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLayoutItem {
	return NewQGraphicsLayoutItemp1(parent)
}
func NewQGraphicsLayoutItemp1(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) *QGraphicsLayoutItem {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	// arg: 1, bool=Bool, =Invalid, , Invalid
	isLayout := false
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemC2EPS_b", qtrt.FFI_TYPE_POINTER, convArg0, isLayout)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQGraphicsLayoutItem)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGraphicsLayoutItem()

/*

 */
func DeleteQGraphicsLayoutItem(this *QGraphicsLayoutItem) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItemD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(const QSizePolicy &)

/*
Sets the size policy to policy. The size policy describes how the item should grow horizontally and vertically when arranged in a layout.

QGraphicsLayoutItem's default size policy is (QSizePolicy::Fixed, QSizePolicy::Fixed, QSizePolicy::DefaultType), but it is common for subclasses to change the default. For example, QGraphicsWidget defaults to (QSizePolicy::Preferred, QSizePolicy::Preferred, QSizePolicy::DefaultType).

See also sizePolicy() and QWidget::sizePolicy().
*/
func (this *QGraphicsLayoutItem) SetSizePolicy(policy QSizePolicy_ITF) {
	var convArg0 unsafe.Pointer
	if policy != nil && policy.QSizePolicy_PTR() != nil {
		convArg0 = policy.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyERK11QSizePolicy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy, QSizePolicy::ControlType)

/*
Sets the size policy to policy. The size policy describes how the item should grow horizontally and vertically when arranged in a layout.

QGraphicsLayoutItem's default size policy is (QSizePolicy::Fixed, QSizePolicy::Fixed, QSizePolicy::DefaultType), but it is common for subclasses to change the default. For example, QGraphicsWidget defaults to (QSizePolicy::Preferred, QSizePolicy::Preferred, QSizePolicy::DefaultType).

See also sizePolicy() and QWidget::sizePolicy().
*/
func (this *QGraphicsLayoutItem) SetSizePolicy1(hPolicy int, vPolicy int, controlType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hPolicy, vPolicy, controlType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy, QSizePolicy::ControlType)

/*
Sets the size policy to policy. The size policy describes how the item should grow horizontally and vertically when arranged in a layout.

QGraphicsLayoutItem's default size policy is (QSizePolicy::Fixed, QSizePolicy::Fixed, QSizePolicy::DefaultType), but it is common for subclasses to change the default. For example, QGraphicsWidget defaults to (QSizePolicy::Preferred, QSizePolicy::Preferred, QSizePolicy::DefaultType).

See also sizePolicy() and QWidget::sizePolicy().
*/
func (this *QGraphicsLayoutItem) SetSizePolicy1p(hPolicy int, vPolicy int) {
	// arg: 2, QSizePolicy::ControlType=Elaborated, QSizePolicy::ControlType=Enum, , Invalid
	controlType := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem13setSizePolicyEN11QSizePolicy6PolicyES1_NS0_11ControlTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hPolicy, vPolicy, controlType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:63
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy sizePolicy() const

/*
Returns the current size policy.

See also setSizePolicy() and QWidget::sizePolicy().
*/
func (this *QGraphicsLayoutItem) SizePolicy() *QSizePolicy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem10sizePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizePolicy)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSizeF &)

/*
Sets the minimum size to size. This property overrides sizeHint() for Qt::MinimumSize and ensures that effectiveSizeHint() will never return a size smaller than size. In order to unset the minimum size, use an invalid size.

See also minimumSize(), maximumSize(), preferredSize(), Qt::MinimumSize, sizeHint(), setMinimumWidth(), and setMinimumHeight().
*/
func (this *QGraphicsLayoutItem) SetMinimumSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:66
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setMinimumSize(qreal, qreal)

/*
Sets the minimum size to size. This property overrides sizeHint() for Qt::MinimumSize and ensures that effectiveSizeHint() will never return a size smaller than size. In order to unset the minimum size, use an invalid size.

See also minimumSize(), maximumSize(), preferredSize(), Qt::MinimumSize, sizeHint(), setMinimumWidth(), and setMinimumHeight().
*/
func (this *QGraphicsLayoutItem) SetMinimumSize1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMinimumSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:67
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF minimumSize() const

/*
Returns the minimum size.

See also setMinimumSize(), preferredSize(), maximumSize(), Qt::MinimumSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) MinimumSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(qreal)

/*
Sets the minimum width to width.

See also minimumWidth(), setMinimumSize(), and minimumSize().
*/
func (this *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMinimumWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:69
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal minimumWidth() const

/*
Returns the minimum width.

See also setMinimumWidth(), setMinimumSize(), and minimumSize().
*/
func (this *QGraphicsLayoutItem) MinimumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12minimumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(qreal)

/*
Sets the minimum height to height.

See also minimumHeight(), setMinimumSize(), and minimumSize().
*/
func (this *QGraphicsLayoutItem) SetMinimumHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMinimumHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:71
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal minimumHeight() const

/*
Returns the minimum height.

See also setMinimumHeight(), setMinimumSize(), and minimumSize().
*/
func (this *QGraphicsLayoutItem) MinimumHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13minimumHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredSize(const QSizeF &)

/*
Sets the preferred size to size. This property overrides sizeHint() for Qt::PreferredSize and provides the default value for effectiveSizeHint(). In order to unset the preferred size, use an invalid size.

See also preferredSize(), minimumSize(), maximumSize(), Qt::PreferredSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) SetPreferredSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:74
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setPreferredSize(qreal, qreal)

/*
Sets the preferred size to size. This property overrides sizeHint() for Qt::PreferredSize and provides the default value for effectiveSizeHint(). In order to unset the preferred size, use an invalid size.

See also preferredSize(), minimumSize(), maximumSize(), Qt::PreferredSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) SetPreferredSize1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setPreferredSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:75
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF preferredSize() const

/*
Returns the preferred size.

See also setPreferredSize(), minimumSize(), maximumSize(), Qt::PreferredSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) PreferredSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13preferredSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:76
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredWidth(qreal)

/*
Sets the preferred width to width.

See also preferredWidth(), preferredHeight(), setPreferredSize(), and preferredSize().
*/
func (this *QGraphicsLayoutItem) SetPreferredWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem17setPreferredWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:77
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal preferredWidth() const

/*
Returns the preferred width.

See also setPreferredWidth(), setPreferredSize(), and preferredSize().
*/
func (this *QGraphicsLayoutItem) PreferredWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem14preferredWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPreferredHeight(qreal)

/*
Sets the preferred height to height.

See also preferredHeight(), preferredWidth(), setPreferredSize(), and preferredSize().
*/
func (this *QGraphicsLayoutItem) SetPreferredHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem18setPreferredHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal preferredHeight() const

/*
Returns the preferred height.

See also setPreferredHeight(), setPreferredSize(), and preferredSize().
*/
func (this *QGraphicsLayoutItem) PreferredHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem15preferredHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSizeF &)

/*
Sets the maximum size to size. This property overrides sizeHint() for Qt::MaximumSize and ensures that effectiveSizeHint() will never return a size larger than size. In order to unset the maximum size, use an invalid size.

See also maximumSize(), minimumSize(), preferredSize(), Qt::MaximumSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) SetMaximumSize(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setMaximumSize(qreal, qreal)

/*
Sets the maximum size to size. This property overrides sizeHint() for Qt::MaximumSize and ensures that effectiveSizeHint() will never return a size larger than size. In order to unset the maximum size, use an invalid size.

See also maximumSize(), minimumSize(), preferredSize(), Qt::MaximumSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) SetMaximumSize1(w float64, h float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14setMaximumSizeEdd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:83
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF maximumSize() const

/*
Returns the maximum size.

See also setMaximumSize(), minimumSize(), preferredSize(), Qt::MaximumSize, and sizeHint().
*/
func (this *QGraphicsLayoutItem) MaximumSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(qreal)

/*
Sets the maximum width to width.

See also maximumWidth(), setMaximumSize(), and maximumSize().
*/
func (this *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setMaximumWidthEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), width)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:85
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal maximumWidth() const

/*
Returns the maximum width.

See also setMaximumWidth(), setMaximumSize(), and maximumSize().
*/
func (this *QGraphicsLayoutItem) MaximumWidth() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12maximumWidthEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(qreal)

/*
Sets the maximum height to height.

See also maximumHeight(), setMaximumSize(), and maximumSize().
*/
func (this *QGraphicsLayoutItem) SetMaximumHeight(height float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setMaximumHeightEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), height)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:87
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal maximumHeight() const

/*
Returns the maximum height.

See also setMaximumHeight(), setMaximumSize(), and maximumSize().
*/
func (this *QGraphicsLayoutItem) MaximumHeight() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13maximumHeightEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setGeometry(const QRectF &)

/*
This virtual function sets the geometry of the QGraphicsLayoutItem to rect, which is in parent coordinates (e.g., the top-left corner of rect is equivalent to the item's position in parent coordinates).

You must reimplement this function in a subclass of QGraphicsLayoutItem to receive geometry updates. The layout will call this function when it does a rearrangement.

If rect is outside of the bounds of minimumSize and maximumSize, it will be adjusted to its closest size so that it is within the legal bounds.

See also geometry().
*/
func (this *QGraphicsLayoutItem) SetGeometry(rect qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRectF_PTR() != nil {
		convArg0 = rect.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem11setGeometryERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:90
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF geometry() const

/*
Returns the item's geometry (e.g., position and size) as a QRectF. This function is equivalent to QRectF(pos(), size()).

See also setGeometry().
*/
func (this *QGraphicsLayoutItem) Geometry() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void getContentsMargins(qreal *, qreal *, qreal *, qreal *) const

/*
This virtual function provides the left, top, right and bottom contents margins for this QGraphicsLayoutItem. The default implementation assumes all contents margins are 0. The parameters point to values stored in qreals. If any of the pointers is 0, that value will not be updated.

See also QGraphicsWidget::setContentsMargins().
*/
func (this *QGraphicsLayoutItem) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem18getContentsMarginsEPdS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:92
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF contentsRect() const

/*
Returns the contents rect in local coordinates.

The contents rect defines the subrectangle used by an associated layout when arranging subitems. This function is a convenience function that adjusts the item's geometry() by its contents margins. Note that getContentsMargins() is a virtual function that you can reimplement to return the item's contents margins.

See also getContentsMargins() and geometry().
*/
func (this *QGraphicsLayoutItem) ContentsRect() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12contentsRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:94
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF effectiveSizeHint(Qt::SizeHint, const QSizeF &) const

/*
Returns the effective size hint for this QGraphicsLayoutItem.

which is the size hint in question. constraint is an optional argument that defines a special constrain when calculating the effective size hint. By default, constraint is QSizeF(-1, -1), which means there is no constraint to the size hint.

If you want to specify the widget's size hint for a given width or height, you can provide the fixed dimension in constraint. This is useful for widgets that can grow only either vertically or horizontally, and need to set either their width or their height to a special value.

For example, a text paragraph item fit into a column width of 200 may grow vertically. You can pass QSizeF(200, -1) as a constraint to get a suitable minimum, preferred and maximum height).

You can adjust the effective size hint by reimplementing sizeHint() in a QGraphicsLayoutItem subclass, or by calling one of the following functions: setMinimumSize(), setPreferredSize, or setMaximumSize() (or a combination of both).

This function caches each of the size hints and guarantees that sizeHint() will be called only once for each value of which - unless constraint is not specified and updateGeometry() has been called.

See also sizeHint().
*/
func (this *QGraphicsLayoutItem) EffectiveSizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem17effectiveSizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:94
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF effectiveSizeHint(Qt::SizeHint, const QSizeF &) const

/*
Returns the effective size hint for this QGraphicsLayoutItem.

which is the size hint in question. constraint is an optional argument that defines a special constrain when calculating the effective size hint. By default, constraint is QSizeF(-1, -1), which means there is no constraint to the size hint.

If you want to specify the widget's size hint for a given width or height, you can provide the fixed dimension in constraint. This is useful for widgets that can grow only either vertically or horizontally, and need to set either their width or their height to a special value.

For example, a text paragraph item fit into a column width of 200 may grow vertically. You can pass QSizeF(200, -1) as a constraint to get a suitable minimum, preferred and maximum height).

You can adjust the effective size hint by reimplementing sizeHint() in a QGraphicsLayoutItem subclass, or by calling one of the following functions: setMinimumSize(), setPreferredSize, or setMaximumSize() (or a combination of both).

This function caches each of the size hints and guarantees that sizeHint() will be called only once for each value of which - unless constraint is not specified and updateGeometry() has been called.

See also sizeHint().
*/
func (this *QGraphicsLayoutItem) EffectiveSizeHintp(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem17effectiveSizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:96
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void updateGeometry()

/*
This virtual function discards any cached size hint information. You should always call this function if you change the return value of the sizeHint() function. Subclasses must always call the base implementation when reimplementing this function.

See also effectiveSizeHint().
*/
func (this *QGraphicsLayoutItem) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsLayoutItem * parentLayoutItem() const

/*
Returns the parent of this QGraphicsLayoutItem, or 0 if there is no parent, or if the parent does not inherit from QGraphicsLayoutItem (QGraphicsLayoutItem is often used through multiple inheritance with QObject-derived classes).

See also setParentLayoutItem().
*/
func (this *QGraphicsLayoutItem) ParentLayoutItem() *QGraphicsLayoutItem /*777 QGraphicsLayoutItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem16parentLayoutItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsLayoutItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParentLayoutItem(QGraphicsLayoutItem *)

/*
Sets the parent of this QGraphicsLayoutItem to parent.

See also parentLayoutItem().
*/
func (this *QGraphicsLayoutItem) SetParentLayoutItem(parent QGraphicsLayoutItem_ITF /*777 QGraphicsLayoutItem **/) {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QGraphicsLayoutItem_PTR() != nil {
		convArg0 = parent.QGraphicsLayoutItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem19setParentLayoutItemEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:101
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLayout() const

/*
Returns true if this QGraphicsLayoutItem is a layout (e.g., is inherited by an object that arranges other QGraphicsLayoutItem objects); otherwise returns false.

See also QGraphicsLayout.
*/
func (this *QGraphicsLayoutItem) IsLayout() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8isLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:102
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsItem * graphicsItem() const

/*
Returns the QGraphicsItem that this layout item represents. For QGraphicsWidget it will return itself. For custom items it can return an aggregated value.

See also setGraphicsItem().
*/
func (this *QGraphicsLayoutItem) GraphicsItem() *QGraphicsItem /*777 QGraphicsItem **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem12graphicsItemEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGraphicsItemFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool ownedByLayout() const

/*
Returns whether a layout should delete this item in its destructor. If its true, then the layout will delete it. If its false, then it is assumed that another object has the ownership of it, and the layout won't delete this item.

If the item inherits both QGraphicsItem and QGraphicsLayoutItem (such as QGraphicsWidget does) the item is really part of two ownership hierarchies. This property informs what the layout should do with its child items when it is destructed. In the case of QGraphicsWidget, it is preferred that when the layout is deleted it won't delete its children (since they are also part of the graphics item hierarchy).

By default this value is initialized to false in QGraphicsLayoutItem, but it is overridden by QGraphicsLayout to return true. This is because QGraphicsLayout is not normally part of the QGraphicsItem hierarchy, so the parent layout should delete it. Subclasses might override this default behaviour by calling setOwnedByLayout(true).

This function was introduced in  Qt 4.6.

See also setOwnedByLayout().
*/
func (this *QGraphicsLayoutItem) OwnedByLayout() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem13ownedByLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:106
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setGraphicsItem(QGraphicsItem *)

/*
If the QGraphicsLayoutItem represents a QGraphicsItem, and it wants to take advantage of the automatic reparenting capabilities of QGraphicsLayout it should set this value. Note that if you delete item and not delete the layout item, you are responsible of calling setGraphicsItem(0) in order to avoid having a dangling pointer.

See also graphicsItem().
*/
func (this *QGraphicsLayoutItem) SetGraphicsItem(item QGraphicsItem_ITF /*777 QGraphicsItem **/) {
	var convArg0 unsafe.Pointer
	if item != nil && item.QGraphicsItem_PTR() != nil {
		convArg0 = item.QGraphicsItem_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem15setGraphicsItemEP13QGraphicsItem", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:107
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void setOwnedByLayout(bool)

/*
Sets whether a layout should delete this item in its destructor or not. ownership must be true to in order for the layout to delete it.

This function was introduced in  Qt 4.6.

See also ownedByLayout().
*/
func (this *QGraphicsLayoutItem) SetOwnedByLayout(ownedByLayout bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QGraphicsLayoutItem16setOwnedByLayoutEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ownedByLayout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:110
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
This pure virtual function returns the size hint for which of the QGraphicsLayoutItem, using the width or height of constraint to constrain the output.

Reimplement this function in a subclass of QGraphicsLayoutItem to provide the necessary size hints for your items.

See also effectiveSizeHint().
*/
func (this *QGraphicsLayoutItem) SizeHint(which int, constraint qtcore.QSizeF_ITF) *qtcore.QSizeF /*123*/ {
	var convArg1 unsafe.Pointer
	if constraint != nil && constraint.QSizeF_PTR() != nil {
		convArg1 = constraint.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicslayoutitem.h:110
// index:0
// Protected purevirtual virtual Visibility=Default Availability=Available
// [16] QSizeF sizeHint(Qt::SizeHint, const QSizeF &) const

/*
This pure virtual function returns the size hint for which of the QGraphicsLayoutItem, using the width or height of constraint to constrain the output.

Reimplement this function in a subclass of QGraphicsLayoutItem to provide the necessary size hints for your items.

See also effectiveSizeHint().
*/
func (this *QGraphicsLayoutItem) SizeHintp(which int) *qtcore.QSizeF /*123*/ {
	// arg: 1, const QSizeF &=LValueReference, QSizeF=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QGraphicsLayoutItem8sizeHintEN2Qt8SizeHintERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
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
