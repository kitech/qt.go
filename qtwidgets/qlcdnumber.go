package qtwidgets

// /usr/include/qt/QtWidgets/qlcdnumber.h
// #include <qlcdnumber.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 53
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

// bool event(QEvent *)
func (this *QLCDNumber) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(QPaintEvent *)
func (this *QLCDNumber) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

/*

 */
type QLCDNumber struct {
	*QFrame
}
type QLCDNumber_ITF interface {
	QFrame_ITF
	QLCDNumber_PTR() *QLCDNumber
}

func (ptr *QLCDNumber) QLCDNumber_PTR() *QLCDNumber { return ptr }

func (this *QLCDNumber) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QFrame.GetCthis()
	}
}
func (this *QLCDNumber) SetCthis(cthis unsafe.Pointer) {
	this.QFrame = NewQFrameFromPointer(cthis)
}
func NewQLCDNumberFromPointer(cthis unsafe.Pointer) *QLCDNumber {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QLCDNumber{bcthis0}
}
func (*QLCDNumber) NewFromPointer(cthis unsafe.Pointer) *QLCDNumber {
	return NewQLCDNumberFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLCDNumber) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLCDNumber(QWidget *)

/*
Constructs an LCD number, sets the number of digits to 5, the base to decimal, the decimal point mode to 'small' and the frame style to a raised box. The segmentStyle() is set to Outline.

The parent argument is passed to the QFrame constructor.

See also setDigitCount() and setSmallDecimalPoint().
*/
func NewQLCDNumber(parent QWidget_ITF /*777 QWidget **/) *QLCDNumber {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg0 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLCDNumber")
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLCDNumber(QWidget *)

/*
Constructs an LCD number, sets the number of digits to 5, the base to decimal, the decimal point mode to 'small' and the frame style to a raised box. The segmentStyle() is set to Outline.

The parent argument is passed to the QFrame constructor.

See also setDigitCount() and setSmallDecimalPoint().
*/
func NewQLCDNumber__() *QLCDNumber {
	// arg: 0, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLCDNumber")
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLCDNumber(uint, QWidget *)

/*
Constructs an LCD number, sets the number of digits to 5, the base to decimal, the decimal point mode to 'small' and the frame style to a raised box. The segmentStyle() is set to Outline.

The parent argument is passed to the QFrame constructor.

See also setDigitCount() and setSmallDecimalPoint().
*/
func NewQLCDNumber_1(numDigits uint, parent QWidget_ITF /*777 QWidget **/) *QLCDNumber {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QWidget_PTR() != nil {
		convArg1 = parent.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberC2EjP7QWidget", qtrt.FFI_TYPE_POINTER, numDigits, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLCDNumber")
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLCDNumber(uint, QWidget *)

/*
Constructs an LCD number, sets the number of digits to 5, the base to decimal, the decimal point mode to 'small' and the frame style to a raised box. The segmentStyle() is set to Outline.

The parent argument is passed to the QFrame constructor.

See also setDigitCount() and setSmallDecimalPoint().
*/
func NewQLCDNumber_1_(numDigits uint) *QLCDNumber {
	// arg: 1, QWidget *=Pointer, QWidget=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberC2EjP7QWidget", qtrt.FFI_TYPE_POINTER, numDigits, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLCDNumber")
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLCDNumber()

/*

 */
func DeleteQLCDNumber(this *QLCDNumber) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool smallDecimalPoint() const

/*

 */
func (this *QLCDNumber) SmallDecimalPoint() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber17smallDecimalPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int digitCount() const

/*
Returns the current number of digits.

Note: Getter function for property digitCount.

See also setDigitCount().
*/
func (this *QLCDNumber) DigitCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber10digitCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDigitCount(int)

/*
Sets the current number of digits to numDigits. Must be in the range 0..99.

Note: Setter function for property digitCount.

See also digitCount().
*/
func (this *QLCDNumber) SetDigitCount(nDigits int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber13setDigitCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nDigits)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool checkOverflow(double) const

/*
Returns true if num is too big to be displayed in its entirety; otherwise returns false.

See also display(), digitCount(), and smallDecimalPoint().
*/
func (this *QLCDNumber) CheckOverflow(num float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:80
// index:1
// Public Visibility=Default Availability=Available
// [1] bool checkOverflow(int) const

/*
Returns true if num is too big to be displayed in its entirety; otherwise returns false.

See also display(), digitCount(), and smallDecimalPoint().
*/
func (this *QLCDNumber) CheckOverflow_1(num int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QLCDNumber::Mode mode() const

/*

 */
func (this *QLCDNumber) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMode(QLCDNumber::Mode)

/*

 */
func (this *QLCDNumber) SetMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7setModeENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QLCDNumber::SegmentStyle segmentStyle() const

/*

 */
func (this *QLCDNumber) SegmentStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber12segmentStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSegmentStyle(QLCDNumber::SegmentStyle)

/*

 */
func (this *QLCDNumber) SetSegmentStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber15setSegmentStyleENS_12SegmentStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] double value() const

/*

 */
func (this *QLCDNumber) Value() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber5valueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int intValue() const

/*

 */
func (this *QLCDNumber) IntValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber8intValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint() const

/*
Reimplemented from QWidget::sizeHint().
*/
func (this *QLCDNumber) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void display(const QString &)

/*
Displays the number represented by the string s.

This version of the function disregards mode() and smallDecimalPoint().

These digits and other symbols can be shown: 0/O, 1, 2, 3, 4, 5/S, 6, 7, 8, 9/g, minus, decimal point, A, B, C, D, E, F, h, H, L, o, P, r, u, U, Y, colon, degree sign (which is specified as single quote in the string) and space. QLCDNumber substitutes spaces for illegal characters.

Note: Setter function for property intValue. Setter function for property value.
*/
func (this *QLCDNumber) Display(str string) {
	var tmpArg0 = qtcore.NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7displayERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void display(int)

/*
Displays the number represented by the string s.

This version of the function disregards mode() and smallDecimalPoint().

These digits and other symbols can be shown: 0/O, 1, 2, 3, 4, 5/S, 6, 7, 8, 9/g, minus, decimal point, A, B, C, D, E, F, h, H, L, o, P, r, u, U, Y, colon, degree sign (which is specified as single quote in the string) and space. QLCDNumber substitutes spaces for illegal characters.

Note: Setter function for property intValue. Setter function for property value.
*/
func (this *QLCDNumber) Display_1(num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7displayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void display(double)

/*
Displays the number represented by the string s.

This version of the function disregards mode() and smallDecimalPoint().

These digits and other symbols can be shown: 0/O, 1, 2, 3, 4, 5/S, 6, 7, 8, 9/g, minus, decimal point, A, B, C, D, E, F, h, H, L, o, P, r, u, U, Y, colon, degree sign (which is specified as single quote in the string) and space. QLCDNumber substitutes spaces for illegal characters.

Note: Setter function for property intValue. Setter function for property value.
*/
func (this *QLCDNumber) Display_2(num float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7displayEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHexMode()

/*
Calls setMode(Hex). Provided for convenience (e.g. for connecting buttons to it).

See also setMode(), setDecMode(), setOctMode(), setBinMode(), and mode().
*/
func (this *QLCDNumber) SetHexMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setHexModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecMode()

/*
Calls setMode(Dec). Provided for convenience (e.g. for connecting buttons to it).

See also setMode(), setHexMode(), setOctMode(), setBinMode(), and mode().
*/
func (this *QLCDNumber) SetDecMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setDecModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOctMode()

/*
Calls setMode(Oct). Provided for convenience (e.g. for connecting buttons to it).

See also setMode(), setHexMode(), setDecMode(), setBinMode(), and mode().
*/
func (this *QLCDNumber) SetOctMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setOctModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBinMode()

/*
Calls setMode(Bin). Provided for convenience (e.g. for connecting buttons to it).

See also setMode(), setHexMode(), setDecMode(), setOctMode(), and mode().
*/
func (this *QLCDNumber) SetBinMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setBinModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSmallDecimalPoint(bool)

/*

 */
func (this *QLCDNumber) SetSmallDecimalPoint(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber20setSmallDecimalPointEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void overflow()

/*
This signal is emitted whenever the QLCDNumber is asked to display a too-large number or a too-long string.

It is never emitted by setDigitCount().
*/
func (this *QLCDNumber) Overflow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber8overflowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)

/*
Reimplemented from QObject::event().
*/
func (this *QLCDNumber) Event(e qtcore.QEvent_ITF /*777 QEvent **/) bool {
	var convArg0 unsafe.Pointer
	if e != nil && e.QEvent_PTR() != nil {
		convArg0 = e.QEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:108
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)

/*
Reimplemented from QWidget::paintEvent().
*/
func (this *QLCDNumber) PaintEvent(arg0 qtgui.QPaintEvent_ITF /*777 QPaintEvent **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPaintEvent_PTR() != nil {
		convArg0 = arg0.QPaintEvent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
This type determines how numbers are shown.



If the display is set to hexadecimal, octal or binary, the integer equivalent of the value is displayed.

*/
type QLCDNumber__Mode = int

// Hexadecimal
const QLCDNumber__Hex QLCDNumber__Mode = 0

// Decimal
const QLCDNumber__Dec QLCDNumber__Mode = 1

// Octal
const QLCDNumber__Oct QLCDNumber__Mode = 2

// Binary
const QLCDNumber__Bin QLCDNumber__Mode = 3

/*
This type determines the visual appearance of the QLCDNumber widget.


*/
type QLCDNumber__SegmentStyle = int

// gives raised segments filled with the background color.
const QLCDNumber__Outline QLCDNumber__SegmentStyle = 0

// gives raised segments filled with the windowText color.
const QLCDNumber__Filled QLCDNumber__SegmentStyle = 1

// gives flat segments filled with the windowText color.
const QLCDNumber__Flat QLCDNumber__SegmentStyle = 2

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
