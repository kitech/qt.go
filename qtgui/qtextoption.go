package qtgui

// /usr/include/qt/QtGui/qtextoption.h
// #include <qtextoption.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 32
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QTextOption struct {
	*qtrt.CObject
}
type QTextOption_ITF interface {
	QTextOption_PTR() *QTextOption
}

func (ptr *QTextOption) QTextOption_PTR() *QTextOption { return ptr }

func (this *QTextOption) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextOption) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextOptionFromPointer(cthis unsafe.Pointer) *QTextOption {
	return &QTextOption{&qtrt.CObject{cthis}}
}
func (*QTextOption) NewFromPointer(cthis unsafe.Pointer) *QTextOption {
	return NewQTextOptionFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextoption.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextOption()

/*
Constructs a text option with default properties for text. The text alignment property is set to Qt::AlignLeft. The word wrap property is set to QTextOption::WordWrap. The using of design metrics flag is set to false.
*/
func (*QTextOption) NewForInherit() *QTextOption {
	return NewQTextOption()
}
func NewQTextOption() *QTextOption {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOptionC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextOption)
	return gothis
}

// /usr/include/qt/QtGui/qtextoption.h:86
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QTextOption(Qt::Alignment)

/*
Constructs a text option with default properties for text. The text alignment property is set to Qt::AlignLeft. The word wrap property is set to QTextOption::WordWrap. The using of design metrics flag is set to false.
*/
func (*QTextOption) NewForInherit1(alignment int) *QTextOption {
	return NewQTextOption1(alignment)
}
func NewQTextOption1(alignment int) *QTextOption {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOptionC2E6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, alignment)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextOption)
	return gothis
}

// /usr/include/qt/QtGui/qtextoption.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QTextOption()

/*

 */
func DeleteQTextOption(this *QTextOption) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOptionD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextoption.h:90
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextOption & operator=(const QTextOption &)

/*

 */
func (this *QTextOption) Operator_equal(o QTextOption_ITF) *QTextOption {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextOption_PTR() != nil {
		convArg0 = o.QTextOption_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOptionaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextOptionFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextOption)
	return rv2
}

// /usr/include/qt/QtGui/qtextoption.h:92
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*
Sets the option's text alignment to the specified alignment.

See also alignment().
*/
func (this *QTextOption) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*
Returns the text alignment defined by the option.

See also setAlignment().
*/
func (this *QTextOption) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextDirection(Qt::LayoutDirection)

/*
Sets the direction of the text layout defined by the option to the given direction.

See also textDirection().
*/
func (this *QTextOption) SetTextDirection(aDirection int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption16setTextDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aDirection)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::LayoutDirection textDirection() const

/*
Returns the direction of the text layout defined by the option.

See also setTextDirection().
*/
func (this *QTextOption) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWrapMode(QTextOption::WrapMode)

/*
Sets the option's text wrap mode to the given mode.

See also wrapMode().
*/
func (this *QTextOption) SetWrapMode(wrap int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption11setWrapModeENS_8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), wrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextOption::WrapMode wrapMode() const

/*
Returns the text wrap mode defined by the option.

See also setWrapMode().
*/
func (this *QTextOption) WrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption8wrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFlags(QTextOption::Flags)

/*
Sets the flags associated with the option to the given flags.

See also flags().
*/
func (this *QTextOption) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption8setFlagsE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextOption::Flags flags() const

/*
Returns the flags associated with the option.

See also setFlags().
*/
func (this *QTextOption) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:120
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTabStop(qreal)

/*
Sets the default distance in device units between tab stops to the value specified by tabStop.

See also tabStop(), setTabArray(), setTabs(), and tabs().
*/
func (this *QTextOption) SetTabStop(tabStop float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption10setTabStopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabStop)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal tabStop() const

/*
Returns the distance in device units between tab stops. Convenient function for the above method

See also setTabStop(), tabArray(), setTabs(), and tabs().
*/
func (this *QTextOption) TabStop() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption7tabStopEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextoption.h:129
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setUseDesignMetrics(bool)

/*
If enable is true then the layout will use design metrics; otherwise it will use the metrics of the paint device (which is the default behavior).

See also useDesignMetrics().
*/
func (this *QTextOption) SetUseDesignMetrics(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption19setUseDesignMetricsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:130
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool useDesignMetrics() const

/*
Returns true if the layout uses design rather than device metrics; otherwise returns false.

See also setUseDesignMetrics().
*/
func (this *QTextOption) UseDesignMetrics() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption16useDesignMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
This enum holds the different types of tabulator



This enum was introduced or modified in  Qt 4.4.

*/
type QTextOption__TabType = int

// A left-tab
const QTextOption__LeftTab QTextOption__TabType = 0

// A right-tab
const QTextOption__RightTab QTextOption__TabType = 1

// A centered-tab
const QTextOption__CenterTab QTextOption__TabType = 2

// A tab stopping at a certain delimiter-character
const QTextOption__DelimiterTab QTextOption__TabType = 3

func (this *QTextOption) TabTypeItemName(val int) string {
	switch val {
	case QTextOption__LeftTab: // 0
		return "LeftTab"
	case QTextOption__RightTab: // 1
		return "RightTab"
	case QTextOption__CenterTab: // 2
		return "CenterTab"
	case QTextOption__DelimiterTab: // 3
		return "DelimiterTab"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextOption_TabTypeItemName(val int) string {
	var nilthis *QTextOption
	return nilthis.TabTypeItemName(val)
}

/*
This enum describes how text is wrapped in a document.


*/
type QTextOption__WrapMode = int

// Text is not wrapped at all.
const QTextOption__NoWrap QTextOption__WrapMode = 0

// Text is wrapped at word boundaries.
const QTextOption__WordWrap QTextOption__WrapMode = 1

// Same as QTextOption::NoWrap
const QTextOption__ManualWrap QTextOption__WrapMode = 2

// Text can be wrapped at any point on a line, even if it occurs in the middle of a word.
const QTextOption__WrapAnywhere QTextOption__WrapMode = 3

// If possible, wrapping occurs at a word boundary; otherwise it will occur at the appropriate point on the line, even in the middle of a word.
const QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = 4

func (this *QTextOption) WrapModeItemName(val int) string {
	switch val {
	case QTextOption__NoWrap: // 0
		return "NoWrap"
	case QTextOption__WordWrap: // 1
		return "WordWrap"
	case QTextOption__ManualWrap: // 2
		return "ManualWrap"
	case QTextOption__WrapAnywhere: // 3
		return "WrapAnywhere"
	case QTextOption__WrapAtWordBoundaryOrAnywhere: // 4
		return "WrapAtWordBoundaryOrAnywhere"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextOption_WrapModeItemName(val int) string {
	var nilthis *QTextOption
	return nilthis.WrapModeItemName(val)
}

/*


 */
type QTextOption__Flag = int

//
const QTextOption__ShowTabsAndSpaces QTextOption__Flag = 1

//
const QTextOption__ShowLineAndParagraphSeparators QTextOption__Flag = 2

//
const QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = 4

//
const QTextOption__SuppressColors QTextOption__Flag = 8

//
const QTextOption__ShowDocumentTerminator QTextOption__Flag = 16

//
const QTextOption__IncludeTrailingSpaces QTextOption__Flag = -2147483648

func (this *QTextOption) FlagItemName(val int) string {
	switch val {
	case QTextOption__ShowTabsAndSpaces: // 1
		return "ShowTabsAndSpaces"
	case QTextOption__ShowLineAndParagraphSeparators: // 2
		return "ShowLineAndParagraphSeparators"
	case QTextOption__AddSpaceForLineAndParagraphSeparators: // 4
		return "AddSpaceForLineAndParagraphSeparators"
	case QTextOption__SuppressColors: // 8
		return "SuppressColors"
	case QTextOption__ShowDocumentTerminator: // 16
		return "ShowDocumentTerminator"
	case QTextOption__IncludeTrailingSpaces: // -2147483648
		return "IncludeTrailingSpaces"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QTextOption_FlagItemName(val int) string {
	var nilthis *QTextOption
	return nilthis.FlagItemName(val)
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
}

//  keep block end
