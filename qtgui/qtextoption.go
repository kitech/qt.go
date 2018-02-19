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
func NewQTextOption_1(alignment int) *QTextOption {
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
func (this *QTextOption) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:93
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const
func (this *QTextOption) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:95
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTextDirection(Qt::LayoutDirection)
func (this *QTextOption) SetTextDirection(aDirection int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption16setTextDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aDirection)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:96
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::LayoutDirection textDirection() const
func (this *QTextOption) TextDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption13textDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:105
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setWrapMode(enum QTextOption::WrapMode)
func (this *QTextOption) SetWrapMode(wrap int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption11setWrapModeENS_8WrapModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), wrap)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:106
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextOption::WrapMode wrapMode() const
func (this *QTextOption) WrapMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption8wrapModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:117
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFlags(QTextOption::Flags)
func (this *QTextOption) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption8setFlagsE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:118
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QTextOption::Flags flags() const
func (this *QTextOption) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextoption.h:121
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTabStop(qreal)
func (this *QTextOption) SetTabStop(tabStop float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption10setTabStopEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabStop)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:122
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal tabStop() const
func (this *QTextOption) TabStop() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption7tabStopEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextoption.h:125
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTabStopDistance(qreal)
func (this *QTextOption) SetTabStopDistance(tabStopDistance float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption18setTabStopDistanceEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabStopDistance)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal tabStopDistance() const
func (this *QTextOption) TabStopDistance() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption15tabStopDistanceEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextoption.h:134
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setUseDesignMetrics(_Bool)
func (this *QTextOption) SetUseDesignMetrics(b bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTextOption19setUseDesignMetricsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextoption.h:135
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool useDesignMetrics() const
func (this *QTextOption) UseDesignMetrics() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTextOption16useDesignMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QTextOption__TabType = int

const QTextOption__LeftTab QTextOption__TabType = 0
const QTextOption__RightTab QTextOption__TabType = 1
const QTextOption__CenterTab QTextOption__TabType = 2
const QTextOption__DelimiterTab QTextOption__TabType = 3

type QTextOption__WrapMode = int

const QTextOption__NoWrap QTextOption__WrapMode = 0
const QTextOption__WordWrap QTextOption__WrapMode = 1
const QTextOption__ManualWrap QTextOption__WrapMode = 2
const QTextOption__WrapAnywhere QTextOption__WrapMode = 3
const QTextOption__WrapAtWordBoundaryOrAnywhere QTextOption__WrapMode = 4

type QTextOption__Flag = int

const QTextOption__ShowTabsAndSpaces QTextOption__Flag = 1
const QTextOption__ShowLineAndParagraphSeparators QTextOption__Flag = 2
const QTextOption__AddSpaceForLineAndParagraphSeparators QTextOption__Flag = 4
const QTextOption__SuppressColors QTextOption__Flag = 8
const QTextOption__ShowDocumentTerminator QTextOption__Flag = 16
const QTextOption__IncludeTrailingSpaces QTextOption__Flag = -2147483648

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
