//  header block begin
// /usr/include/qt/QtWidgets/qlcdnumber.h
// #include <qlcdnumber.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 49
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
	return this.QFrame.GetCthis()
}
func NewQLCDNumberFromPointer(cthis unsafe.Pointer) *QLCDNumber {
	bcthis0 := NewQFrameFromPointer(cthis)
	return &QLCDNumber{bcthis0}
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:53
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QLCDNumber) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:62
// index:0
// Public
// void QLCDNumber(class QWidget *)
func NewQLCDNumber(parent unsafe.Pointer) *QLCDNumber {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	gothis := NewQLCDNumberFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:63
// index:1
// Public
// void QLCDNumber(uint, class QWidget *)
func NewQLCDNumber_1(numDigits uint, parent unsafe.Pointer) *QLCDNumber {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberC2EjP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &numDigits, parent)
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
func (this *QLCDNumber) SmallDecimalPoint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber17smallDecimalPointEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:76
// index:0
// Public
// int digitCount()
func (this *QLCDNumber) DigitCount() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber10digitCountEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:77
// index:0
// Public
// void setDigitCount(int)
func (this *QLCDNumber) SetDigitCount(nDigits int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber13setDigitCountEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &nDigits)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:79
// index:0
// Public
// bool checkOverflow(double)
func (this *QLCDNumber) CheckOverflow(num float64) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &num)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:80
// index:1
// Public
// bool checkOverflow(int)
func (this *QLCDNumber) CheckOverflow_1(num int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &num)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:82
// index:0
// Public
// QLCDNumber::Mode mode()
func (this *QLCDNumber) Mode() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber4modeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:83
// index:0
// Public
// void setMode(enum QLCDNumber::Mode)
func (this *QLCDNumber) SetMode(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7setModeENS_4ModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:85
// index:0
// Public
// QLCDNumber::SegmentStyle segmentStyle()
func (this *QLCDNumber) SegmentStyle() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber12segmentStyleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:86
// index:0
// Public
// void setSegmentStyle(enum QLCDNumber::SegmentStyle)
func (this *QLCDNumber) SetSegmentStyle(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber15setSegmentStyleENS_12SegmentStyleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:88
// index:0
// Public
// double value()
func (this *QLCDNumber) Value() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:89
// index:0
// Public
// int intValue()
func (this *QLCDNumber) IntValue() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber8intValueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:91
// index:0
// Public virtual
// QSize sizeHint()
func (this *QLCDNumber) SizeHint() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber8sizeHintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:96
// index:2
// Public
// void display(double)
func (this *QLCDNumber) Display_2(num float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &num)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber20setSmallDecimalPointEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &arg0)
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
func (this *QLCDNumber) Event(e unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), e)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:108
// index:0
// Protected virtual
// void paintEvent(class QPaintEvent *)
func (this *QLCDNumber) PaintEvent(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

//  body block end
