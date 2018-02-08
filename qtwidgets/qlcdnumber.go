package qtwidgets

// /usr/include/qt/QtWidgets/qlcdnumber.h
// #include <qlcdnumber.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 53
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		qtrt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QLCDNumber) InheritEvent(f func(e *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QLCDNumber) InheritPaintEvent(f func(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

type QLCDNumber struct {
	*QFrame
}

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
// [8] const QMetaObject * metaObject()
func (this *QLCDNumber) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLCDNumber(QWidget *)
func NewQLCDNumber(parent *QWidget /*777 QWidget **/) *QLCDNumber {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberC2EP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:63
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLCDNumber(uint, QWidget *)
func NewQLCDNumber_1(numDigits uint, parent *QWidget /*777 QWidget **/) *QLCDNumber {
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberC2EjP7QWidget", qtrt.FFI_TYPE_POINTER, numDigits, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLCDNumber()
func DeleteQLCDNumber(this *QLCDNumber) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumberD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:75
// index:0
// Public Visibility=Default Availability=Available
// [1] bool smallDecimalPoint()
func (this *QLCDNumber) SmallDecimalPoint() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber17smallDecimalPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int digitCount()
func (this *QLCDNumber) DigitCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber10digitCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDigitCount(int)
func (this *QLCDNumber) SetDigitCount(nDigits int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber13setDigitCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), nDigits)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool checkOverflow(double)
func (this *QLCDNumber) CheckOverflow(num float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:80
// index:1
// Public Visibility=Default Availability=Available
// [1] bool checkOverflow(int)
func (this *QLCDNumber) CheckOverflow_1(num int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:82
// index:0
// Public Visibility=Default Availability=Available
// [4] QLCDNumber::Mode mode()
func (this *QLCDNumber) Mode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber4modeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMode(enum QLCDNumber::Mode)
func (this *QLCDNumber) SetMode(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7setModeENS_4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] QLCDNumber::SegmentStyle segmentStyle()
func (this *QLCDNumber) SegmentStyle() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber12segmentStyleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:86
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSegmentStyle(enum QLCDNumber::SegmentStyle)
func (this *QLCDNumber) SetSegmentStyle(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber15setSegmentStyleENS_12SegmentStyleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] double value()
func (this *QLCDNumber) Value() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber5valueEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int intValue()
func (this *QLCDNumber) IntValue() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber8intValueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:91
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QLCDNumber) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QLCDNumber8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void display(const QString &)
func (this *QLCDNumber) Display(str string) {
	var tmpArg0 = qtcore.NewQString_5(str)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7displayERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:95
// index:1
// Public Visibility=Default Availability=Available
// [-2] void display(int)
func (this *QLCDNumber) Display_1(num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7displayEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:96
// index:2
// Public Visibility=Default Availability=Available
// [-2] void display(double)
func (this *QLCDNumber) Display_2(num float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber7displayEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHexMode()
func (this *QLCDNumber) SetHexMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setHexModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDecMode()
func (this *QLCDNumber) SetDecMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setDecModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOctMode()
func (this *QLCDNumber) SetOctMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setOctModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBinMode()
func (this *QLCDNumber) SetBinMode() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10setBinModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSmallDecimalPoint(_Bool)
func (this *QLCDNumber) SetSmallDecimalPoint(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber20setSmallDecimalPointEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void overflow()
func (this *QLCDNumber) Overflow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber8overflowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:107
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QLCDNumber) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:108
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QLCDNumber) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QLCDNumber10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QLCDNumber__Mode = int

const QLCDNumber__Hex QLCDNumber__Mode = 0
const QLCDNumber__Dec QLCDNumber__Mode = 1
const QLCDNumber__Oct QLCDNumber__Mode = 2
const QLCDNumber__Bin QLCDNumber__Mode = 3

type QLCDNumber__SegmentStyle = int

const QLCDNumber__Outline QLCDNumber__SegmentStyle = 0
const QLCDNumber__Filled QLCDNumber__SegmentStyle = 1
const QLCDNumber__Flat QLCDNumber__SegmentStyle = 2

//  body block end
