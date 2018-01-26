package qtwidgets

// /usr/include/qt/QtWidgets/qlcdnumber.h
// #include <qlcdnumber.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 52
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
		ffiqt.KeepMe()
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QLCDNumber) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:62
// index:0
// Public
// void QLCDNumber(class QWidget *)
func NewQLCDNumber(parent *QWidget /*777 QWidget **/) *QLCDNumber {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:63
// index:1
// Public
// void QLCDNumber(uint, class QWidget *)
func NewQLCDNumber_1(numDigits uint, parent *QWidget /*777 QWidget **/) *QLCDNumber {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg1 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberC2EjP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, numDigits, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:64
// index:0
// Public virtual
// void ~QLCDNumber()
func DeleteQLCDNumber(*QLCDNumber) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:75
// index:0
// Public
// bool smallDecimalPoint()
func (this *QLCDNumber) SmallDecimalPoint() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber17smallDecimalPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:76
// index:0
// Public
// int digitCount()
func (this *QLCDNumber) DigitCount() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber10digitCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:77
// index:0
// Public
// void setDigitCount(int)
func (this *QLCDNumber) SetDigitCount(nDigits int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber13setDigitCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), nDigits)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:79
// index:0
// Public
// bool checkOverflow(double)
func (this *QLCDNumber) CheckOverflow(num float64) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:80
// index:1
// Public
// bool checkOverflow(int)
func (this *QLCDNumber) CheckOverflow_1(num int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:82
// index:0
// Public
// QLCDNumber::Mode mode()
func (this *QLCDNumber) Mode() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber4modeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:83
// index:0
// Public
// void setMode(enum QLCDNumber::Mode)
func (this *QLCDNumber) SetMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7setModeENS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:85
// index:0
// Public
// QLCDNumber::SegmentStyle segmentStyle()
func (this *QLCDNumber) SegmentStyle() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber12segmentStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:86
// index:0
// Public
// void setSegmentStyle(enum QLCDNumber::SegmentStyle)
func (this *QLCDNumber) SetSegmentStyle(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber15setSegmentStyleENS_12SegmentStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:88
// index:0
// Public
// double value()
func (this *QLCDNumber) Value() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:89
// index:0
// Public
// int intValue()
func (this *QLCDNumber) IntValue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber8intValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:91
// index:0
// Public virtual
// QSize sizeHint()
func (this *QLCDNumber) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:94
// index:0
// Public
// void display(const class QString &)
func (this *QLCDNumber) Display(str *qtcore.QString) {
	var convArg0 = str.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:95
// index:1
// Public
// void display(int)
func (this *QLCDNumber) Display_1(num int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:96
// index:2
// Public
// void display(double)
func (this *QLCDNumber) Display_2(num float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:97
// index:0
// Public
// void setHexMode()
func (this *QLCDNumber) SetHexMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setHexModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:98
// index:0
// Public
// void setDecMode()
func (this *QLCDNumber) SetDecMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setDecModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:99
// index:0
// Public
// void setOctMode()
func (this *QLCDNumber) SetOctMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setOctModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:100
// index:0
// Public
// void setBinMode()
func (this *QLCDNumber) SetBinMode() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setBinModeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:101
// index:0
// Public
// void setSmallDecimalPoint(_Bool)
func (this *QLCDNumber) SetSmallDecimalPoint(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber20setSmallDecimalPointEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:104
// index:0
// Public
// void overflow()
func (this *QLCDNumber) Overflow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber8overflowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:107
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QLCDNumber) Event(e *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:108
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QLCDNumber) PaintEvent(arg0 *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
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
