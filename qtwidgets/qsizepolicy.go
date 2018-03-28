package qtwidgets

// /usr/include/qt/QtWidgets/qsizepolicy.h
// #include <qsizepolicy.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QSizePolicy struct {
	*qtrt.CObject
}
type QSizePolicy_ITF interface {
	QSizePolicy_PTR() *QSizePolicy
}

func (ptr *QSizePolicy) QSizePolicy_PTR() *QSizePolicy { return ptr }

func (this *QSizePolicy) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSizePolicy) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSizePolicyFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return &QSizePolicy{&qtrt.CObject{cthis}}
}
func (*QSizePolicy) NewFromPointer(cthis unsafe.Pointer) *QSizePolicy {
	return NewQSizePolicyFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:113
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QSizePolicy()

/*
Constructs a QSizePolicy object with Fixed as its horizontal and vertical policies.

The policies can be altered using the setHorizontalPolicy() and setVerticalPolicy() functions. Use the setHeightForWidth() function if the preferred height of the widget is dependent on the width of the widget (for example, a QLabel with line wrapping).

See also setHorizontalStretch() and setVerticalStretch().
*/
func NewQSizePolicy() *QSizePolicy {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizePolicy)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy, QSizePolicy::ControlType)

/*
Constructs a QSizePolicy object with Fixed as its horizontal and vertical policies.

The policies can be altered using the setHorizontalPolicy() and setVerticalPolicy() functions. Use the setHeightForWidth() function if the preferred height of the widget is dependent on the width of the widget (for example, a QLabel with line wrapping).

See also setHorizontalStretch() and setVerticalStretch().
*/
func NewQSizePolicy_1(horizontal int, vertical int, type_ int) *QSizePolicy {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyC2ENS_6PolicyES0_NS_11ControlTypeE", qtrt.FFI_TYPE_POINTER, horizontal, vertical, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizePolicy)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:116
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy, QSizePolicy::ControlType)

/*
Constructs a QSizePolicy object with Fixed as its horizontal and vertical policies.

The policies can be altered using the setHorizontalPolicy() and setVerticalPolicy() functions. Use the setHeightForWidth() function if the preferred height of the widget is dependent on the width of the widget (for example, a QLabel with line wrapping).

See also setHorizontalStretch() and setVerticalStretch().
*/
func NewQSizePolicy_1_(horizontal int, vertical int) *QSizePolicy {
	// arg: 2, QSizePolicy::ControlType=Enum, QSizePolicy::ControlType=Enum,
	type_ := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyC2ENS_6PolicyES0_NS_11ControlTypeE", qtrt.FFI_TYPE_POINTER, horizontal, vertical, type_)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSizePolicy)
	return gothis
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSizePolicy::Policy horizontalPolicy() const

/*
Returns the horizontal component of the size policy.

See also setHorizontalPolicy(), verticalPolicy(), and horizontalStretch().
*/
func (this *QSizePolicy) HorizontalPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy16horizontalPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSizePolicy::Policy verticalPolicy() const

/*
Returns the vertical component of the size policy.

See also setVerticalPolicy(), horizontalPolicy(), and verticalStretch().
*/
func (this *QSizePolicy) VerticalPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy14verticalPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy::ControlType controlType() const

/*
Returns the control type associated with the widget for which this size policy applies.

This function was introduced in  Qt 4.3.

See also setControlType().
*/
func (this *QSizePolicy) ControlType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy11controlTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:132
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHorizontalPolicy(QSizePolicy::Policy)

/*
Sets the horizontal component to the given policy.

See also horizontalPolicy(), setVerticalPolicy(), and setHorizontalStretch().
*/
func (this *QSizePolicy) SetHorizontalPolicy(d int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy19setHorizontalPolicyENS_6PolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), d)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:133
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setVerticalPolicy(QSizePolicy::Policy)

/*
Sets the vertical component to the given policy.

See also verticalPolicy(), setHorizontalPolicy(), and setVerticalStretch().
*/
func (this *QSizePolicy) SetVerticalPolicy(d int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy17setVerticalPolicyENS_6PolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), d)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setControlType(QSizePolicy::ControlType)

/*
Sets the control type associated with the widget for which this size policy applies to type.

The control type specifies the type of the widget for which this size policy applies. It is used by some styles, notably QMacStyle, to insert proper spacing between widgets. For example, the macOS Aqua guidelines specify that push buttons should be separated by 12 pixels, whereas vertically stacked radio buttons only require 6 pixels.

This function was introduced in  Qt 4.3.

See also controlType() and QStyle::layoutSpacing().
*/
func (this *QSizePolicy) SetControlType(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy14setControlTypeENS_11ControlTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:136
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Orientations expandingDirections() const

/*
Returns whether a widget can make use of more space than the QWidget::sizeHint() function indicates.

A value of Qt::Horizontal or Qt::Vertical means that the widget can grow horizontally or vertically (i.e., the horizontal or vertical policy is Expanding or MinimumExpanding), whereas Qt::Horizontal | Qt::Vertical means that it can grow in both dimensions.

See also horizontalPolicy() and verticalPolicy().
*/
func (this *QSizePolicy) ExpandingDirections() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy19expandingDirectionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeightForWidth(bool)

/*
Sets the flag determining whether the widget's preferred height depends on its width, to dependent.

See also hasHeightForWidth() and setWidthForHeight().
*/
func (this *QSizePolicy) SetHeightForWidth(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy17setHeightForWidthEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasHeightForWidth() const

/*
Returns true if the widget's preferred height depends on its width; otherwise returns false.

See also setHeightForWidth().
*/
func (this *QSizePolicy) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWidthForHeight(bool)

/*
Sets the flag determining whether the widget's width depends on its height, to dependent.

This is only supported for QGraphicsLayout's subclasses. It is not possible to have a layout with both height-for-width and width-for-height constraints at the same time.

See also hasWidthForHeight() and setHeightForWidth().
*/
func (this *QSizePolicy) SetWidthForHeight(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy17setWidthForHeightEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:144
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool hasWidthForHeight() const

/*
Returns true if the widget's width depends on its height; otherwise returns false.

See also setWidthForHeight().
*/
func (this *QSizePolicy) HasWidthForHeight() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy17hasWidthForHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:146
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QSizePolicy &) const

/*

 */
func (this *QSizePolicy) Operator_equal_equal(s QSizePolicy_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizePolicy_PTR() != nil {
		convArg0 = s.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicyeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:147
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QSizePolicy &) const

/*

 */
func (this *QSizePolicy) Operator_not_equal(s QSizePolicy_ITF) bool {
	var convArg0 unsafe.Pointer
	if s != nil && s.QSizePolicy_PTR() != nil {
		convArg0 = s.QSizePolicy_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicyneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int horizontalStretch() const

/*
Returns the horizontal stretch factor of the size policy.

See also setHorizontalStretch(), verticalStretch(), and horizontalPolicy().
*/
func (this *QSizePolicy) HorizontalStretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy17horizontalStretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:154
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int verticalStretch() const

/*
Returns the vertical stretch factor of the size policy.

See also setVerticalStretch(), horizontalStretch(), and verticalPolicy().
*/
func (this *QSizePolicy) VerticalStretch() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy15verticalStretchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:155
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHorizontalStretch(int)

/*
Sets the horizontal stretch factor of the size policy to the given stretchFactor. stretchFactor must be in the range [0,255].

When two widgets are adjacent to each other in a horizontal layout, setting the horizontal stretch factor of the widget on the left to 2 and the factor of widget on the right to 1 will ensure that the widget on the left will always be twice the size of the one on the right.

See also horizontalStretch(), setVerticalStretch(), and setHorizontalPolicy().
*/
func (this *QSizePolicy) SetHorizontalStretch(stretchFactor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy20setHorizontalStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretchFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:156
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setVerticalStretch(int)

/*
Sets the vertical stretch factor of the size policy to the given stretchFactor. stretchFactor must be in the range [0,255].

When two widgets are adjacent to each other in a vertical layout, setting the vertical stretch factor of the widget on the top to 2 and the factor of widget on the bottom to 1 will ensure that the widget on the top will always be twice the size of the one on the bottom.

See also verticalStretch(), setHorizontalStretch(), and setVerticalPolicy().
*/
func (this *QSizePolicy) SetVerticalStretch(stretchFactor int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy18setVerticalStretchEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stretchFactor)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:158
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool retainSizeWhenHidden() const

/*
Returns whether the layout should retain the widget's size when it is hidden. This is false by default.

This function was introduced in  Qt 5.2.

See also setRetainSizeWhenHidden().
*/
func (this *QSizePolicy) RetainSizeWhenHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy20retainSizeWhenHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:159
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRetainSizeWhenHidden(bool)

/*
Sets whether a layout should retain the widget's size when it is hidden. If retainSize is true, the layout will not be changed by hiding the widget.

This function was introduced in  Qt 5.2.

See also retainSizeWhenHidden().
*/
func (this *QSizePolicy) SetRetainSizeWhenHidden(retainSize bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy23setRetainSizeWhenHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), retainSize)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:161
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void transpose()

/*
Swaps the horizontal and vertical policies and stretches.

See also transposed().
*/
func (this *QSizePolicy) Transpose() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicy9transposeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qsizepolicy.h:166
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSizePolicy transposed() const

/*
Returns a size policy object with the horizontal and vertical policies and stretches swapped.

This function was introduced in  Qt 5.9.

See also transpose().
*/
func (this *QSizePolicy) Transposed() *QSizePolicy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSizePolicy10transposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizePolicy)
	return rv2
}

func DeleteQSizePolicy(this *QSizePolicy) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSizePolicyD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*
These flags are combined together to form the various Policy values:



See also Policy.

*/
type QSizePolicy__PolicyFlag = int

// The widget can grow beyond its size hint if necessary.
const QSizePolicy__GrowFlag QSizePolicy__PolicyFlag = 1

// The widget should get as much space as possible.
const QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = 2

// The widget can shrink below its size hint if necessary.
const QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = 4

// The widget's size hint is ignored. The widget will get as much space as possible.
const QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = 8

/*
This enum describes the various per-dimension sizing types used when constructing a QSizePolicy.

QSizePolicy::MinimumGrowFlagThe sizeHint() is minimal, and sufficient. The widget can be expanded, but there is no advantage to it being larger (e.g. the horizontal direction of a push button). It cannot be smaller than the size provided by sizeHint().
QSizePolicy::MaximumShrinkFlagThe sizeHint() is a maximum. The widget can be shrunk any amount without detriment if other widgets need the space (e.g. a separator line). It cannot be larger than the size provided by sizeHint().
QSizePolicy::PreferredGrowFlag | ShrinkFlagThe sizeHint() is best, but the widget can be shrunk and still be useful. The widget can be expanded, but there is no advantage to it being larger than sizeHint() (the default QWidget policy).
QSizePolicy::ExpandingGrowFlag | ShrinkFlag | ExpandFlagThe sizeHint() is a sensible size, but the widget can be shrunk and still be useful. The widget can make use of extra space, so it should get as much space as possible (e.g. the horizontal direction of a horizontal slider).
QSizePolicy::MinimumExpandingGrowFlag | ExpandFlagThe sizeHint() is minimal, and sufficient. The widget can make use of extra space, so it should get as much space as possible (e.g. the horizontal direction of a horizontal slider).
QSizePolicy::IgnoredShrinkFlag | GrowFlag | IgnoreFlagThe sizeHint() is ignored. The widget will get as much space as possible.


See also PolicyFlag, setHorizontalPolicy(), and setVerticalPolicy().

*/
type QSizePolicy__Policy = int

// The QWidget::sizeHint() is the only acceptable alternative, so the widget can never grow or shrink (e.g. the vertical direction of a push button).
const QSizePolicy__Fixed QSizePolicy__Policy = 0

//
const QSizePolicy__Minimum QSizePolicy__Policy = 1

//
const QSizePolicy__Maximum QSizePolicy__Policy = 4

//
const QSizePolicy__Preferred QSizePolicy__Policy = 5

//
const QSizePolicy__MinimumExpanding QSizePolicy__Policy = 3

//
const QSizePolicy__Expanding QSizePolicy__Policy = 7

//
const QSizePolicy__Ignored QSizePolicy__Policy = 13

/*


 */
type QSizePolicy__ControlType = int

//
const QSizePolicy__DefaultType QSizePolicy__ControlType = 1

//
const QSizePolicy__ButtonBox QSizePolicy__ControlType = 2

//
const QSizePolicy__CheckBox QSizePolicy__ControlType = 4

//
const QSizePolicy__ComboBox QSizePolicy__ControlType = 8

//
const QSizePolicy__Frame QSizePolicy__ControlType = 16

//
const QSizePolicy__GroupBox QSizePolicy__ControlType = 32

//
const QSizePolicy__Label QSizePolicy__ControlType = 64

//
const QSizePolicy__Line QSizePolicy__ControlType = 128

//
const QSizePolicy__LineEdit QSizePolicy__ControlType = 256

//
const QSizePolicy__PushButton QSizePolicy__ControlType = 512

//
const QSizePolicy__RadioButton QSizePolicy__ControlType = 1024

//
const QSizePolicy__Slider QSizePolicy__ControlType = 2048

//
const QSizePolicy__SpinBox QSizePolicy__ControlType = 4096

//
const QSizePolicy__TabWidget QSizePolicy__ControlType = 8192

//
const QSizePolicy__ToolButton QSizePolicy__ControlType = 16384

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
