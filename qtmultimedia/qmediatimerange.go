package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediatimerange.h
// #include <qmediatimerange.h>
// #include <QtMultimedia>

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
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QMediaTimeRange struct {
	*qtrt.CObject
}
type QMediaTimeRange_ITF interface {
	QMediaTimeRange_PTR() *QMediaTimeRange
}

func (ptr *QMediaTimeRange) QMediaTimeRange_PTR() *QMediaTimeRange { return ptr }

func (this *QMediaTimeRange) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaTimeRange) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaTimeRangeFromPointer(cthis unsafe.Pointer) *QMediaTimeRange {
	return &QMediaTimeRange{&qtrt.CObject{cthis}}
}
func (*QMediaTimeRange) NewFromPointer(cthis unsafe.Pointer) *QMediaTimeRange {
	return NewQMediaTimeRangeFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaTimeRange()

/*
Constructs an empty time range.
*/
func (*QMediaTimeRange) NewForInherit() *QMediaTimeRange {
	return NewQMediaTimeRange()
}
func NewQMediaTimeRange() *QMediaTimeRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaTimeRange)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:84
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMediaTimeRange(qint64, qint64)

/*
Constructs an empty time range.
*/
func (*QMediaTimeRange) NewForInherit_1(start int64, end int64) *QMediaTimeRange {
	return NewQMediaTimeRange_1(start, end)
}
func NewQMediaTimeRange_1(start int64, end int64) *QMediaTimeRange {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangeC2Exx", qtrt.FFI_TYPE_POINTER, start, end)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaTimeRange)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:85
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMediaTimeRange(const QMediaTimeInterval &)

/*
Constructs an empty time range.
*/
func (*QMediaTimeRange) NewForInherit_2(arg0 QMediaTimeInterval_ITF) *QMediaTimeRange {
	return NewQMediaTimeRange_2(arg0)
}
func NewQMediaTimeRange_2(arg0 QMediaTimeInterval_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeInterval_PTR() != nil {
		convArg0 = arg0.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangeC2ERK18QMediaTimeInterval", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaTimeRange)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMediaTimeRange()

/*

 */
func DeleteQMediaTimeRange(this *QMediaTimeRange) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:89
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaTimeRange & operator=(const QMediaTimeRange &)

/*

 */
func (this *QMediaTimeRange) Operator_equal(arg0 QMediaTimeRange_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:90
// index:1
// Public Visibility=Default Availability=Available
// [8] QMediaTimeRange & operator=(const QMediaTimeInterval &)

/*

 */
func (this *QMediaTimeRange) Operator_equal_1(arg0 QMediaTimeInterval_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeInterval_PTR() != nil {
		convArg0 = arg0.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangeaSERK18QMediaTimeInterval", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 earliestTime() const

/*
Returns the earliest time within the time range.

For empty time ranges, this value is equal to zero.

See also latestTime().
*/
func (this *QMediaTimeRange) EarliestTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMediaTimeRange12earliestTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 latestTime() const

/*
Returns the latest time within the time range.

For empty time ranges, this value is equal to zero.

See also earliestTime().
*/
func (this *QMediaTimeRange) LatestTime() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMediaTimeRange10latestTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if there are no intervals within the time range.

See also intervals().
*/
func (this *QMediaTimeRange) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMediaTimeRange7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isContinuous() const

/*
Returns true if the time range consists of a continuous interval. That is, there is one or fewer disjoint intervals within the time range.
*/
func (this *QMediaTimeRange) IsContinuous() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMediaTimeRange12isContinuousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool contains(qint64) const

/*
Returns true if the specified time lies within the time range.
*/
func (this *QMediaTimeRange) Contains(time int64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QMediaTimeRange8containsEx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), time)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addInterval(qint64, qint64)

/*
Adds the specified interval to the time range.

Adding intervals which are not normal is invalid, and will be ignored.

If the specified interval is adjacent to, or overlaps existing intervals within the time range, these intervals will be merged.

This operation takes linear time.

See also removeInterval().
*/
func (this *QMediaTimeRange) AddInterval(start int64, end int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange11addIntervalExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:102
// index:1
// Public Visibility=Default Availability=Available
// [-2] void addInterval(const QMediaTimeInterval &)

/*
Adds the specified interval to the time range.

Adding intervals which are not normal is invalid, and will be ignored.

If the specified interval is adjacent to, or overlaps existing intervals within the time range, these intervals will be merged.

This operation takes linear time.

See also removeInterval().
*/
func (this *QMediaTimeRange) AddInterval_1(interval QMediaTimeInterval_ITF) {
	var convArg0 unsafe.Pointer
	if interval != nil && interval.QMediaTimeInterval_PTR() != nil {
		convArg0 = interval.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange11addIntervalERK18QMediaTimeInterval", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addTimeRange(const QMediaTimeRange &)

/*
Adds each of the intervals in range to this time range.

Equivalent to calling addInterval() for each interval in range.
*/
func (this *QMediaTimeRange) AddTimeRange(arg0 QMediaTimeRange_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange12addTimeRangeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeInterval(qint64, qint64)

/*
Removes the specified interval from the time range.

Removing intervals which are not normal is invalid, and will be ignored.

Intervals within the time range will be trimmed, split or deleted such that no intervals within the time range include any part of the target interval.

This operation takes linear time.

See also addInterval().
*/
func (this *QMediaTimeRange) RemoveInterval(start int64, end int64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange14removeIntervalExx", qtrt.FFI_TYPE_POINTER, this.GetCthis(), start, end)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:106
// index:1
// Public Visibility=Default Availability=Available
// [-2] void removeInterval(const QMediaTimeInterval &)

/*
Removes the specified interval from the time range.

Removing intervals which are not normal is invalid, and will be ignored.

Intervals within the time range will be trimmed, split or deleted such that no intervals within the time range include any part of the target interval.

This operation takes linear time.

See also addInterval().
*/
func (this *QMediaTimeRange) RemoveInterval_1(interval QMediaTimeInterval_ITF) {
	var convArg0 unsafe.Pointer
	if interval != nil && interval.QMediaTimeInterval_PTR() != nil {
		convArg0 = interval.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange14removeIntervalERK18QMediaTimeInterval", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeTimeRange(const QMediaTimeRange &)

/*
Removes each of the intervals in range from this time range.

Equivalent to calling removeInterval() for each interval in range.
*/
func (this *QMediaTimeRange) RemoveTimeRange(arg0 QMediaTimeRange_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange15removeTimeRangeERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaTimeRange & operator+=(const QMediaTimeRange &)

/*

 */
func (this *QMediaTimeRange) Operator_add_equal(arg0 QMediaTimeRange_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangepLERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:110
// index:1
// Public Visibility=Default Availability=Available
// [8] QMediaTimeRange & operator+=(const QMediaTimeInterval &)

/*

 */
func (this *QMediaTimeRange) Operator_add_equal_1(arg0 QMediaTimeInterval_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeInterval_PTR() != nil {
		convArg0 = arg0.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangepLERK18QMediaTimeInterval", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:111
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaTimeRange & operator-=(const QMediaTimeRange &)

/*

 */
func (this *QMediaTimeRange) Operator_minus_equal(arg0 QMediaTimeRange_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeRange_PTR() != nil {
		convArg0 = arg0.QMediaTimeRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangemIERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:112
// index:1
// Public Visibility=Default Availability=Available
// [8] QMediaTimeRange & operator-=(const QMediaTimeInterval &)

/*

 */
func (this *QMediaTimeRange) Operator_minus_equal_1(arg0 QMediaTimeInterval_ITF) *QMediaTimeRange {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaTimeInterval_PTR() != nil {
		convArg0 = arg0.QMediaTimeInterval_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRangemIERK18QMediaTimeInterval", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaTimeRangeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaTimeRange)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediatimerange.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all intervals from the time range.

See also removeInterval().
*/
func (this *QMediaTimeRange) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QMediaTimeRange5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
