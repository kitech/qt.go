//  header block begin
// /usr/include/qt/QtWidgets/qlcdnumber.h
// #include <qlcdnumber.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 38
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:53
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QLCDNumber) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:62
// index:0
// void QLCDNumber(class QWidget *)
func NewQLCDNumber(parent unsafe.Pointer) *QLCDNumber {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberC2EP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QLCDNumber{cthis}
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:63
// index:1
// void QLCDNumber(uint, class QWidget *)
func NewQLCDNumber_1(numDigits uint, parent unsafe.Pointer) *QLCDNumber {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberC2EjP7QWidget", ffiqt.FFI_TYPE_VOID, cthis, &numDigits, parent)
	gopp.ErrPrint(err, rv)
	return &QLCDNumber{cthis}
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:64
// index:0
// virtual
// void ~QLCDNumber()
func DeleteQLCDNumber(*QLCDNumber) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumberD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:75
// index:0
// bool smallDecimalPoint()
func (this *QLCDNumber) SmallDecimalPoint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber17smallDecimalPointEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:76
// index:0
// int digitCount()
func (this *QLCDNumber) DigitCount() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber10digitCountEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:77
// index:0
// void setDigitCount(int)
func (this *QLCDNumber) SetDigitCount(nDigits int) {
	// 0: (, int nDigits), (&nDigits)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber13setDigitCountEi", ffiqt.FFI_TYPE_VOID, this.cthis, &nDigits)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:79
// index:0
// bool checkOverflow(double)
func (this *QLCDNumber) CheckOverflow(num float64) {
	// 0: (, double num), (&num)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEd", ffiqt.FFI_TYPE_VOID, this.cthis, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:80
// index:1
// bool checkOverflow(int)
func (this *QLCDNumber) CheckOverflow_1(num int) {
	// 1: (, int num), (&num)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber13checkOverflowEi", ffiqt.FFI_TYPE_VOID, this.cthis, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:82
// index:0
// QLCDNumber::Mode mode()
func (this *QLCDNumber) Mode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber4modeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:83
// index:0
// void setMode(enum QLCDNumber::Mode)
func (this *QLCDNumber) SetMode(arg0 int) {
	// 0: (, QLCDNumber::Mode arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7setModeENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:85
// index:0
// QLCDNumber::SegmentStyle segmentStyle()
func (this *QLCDNumber) SegmentStyle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber12segmentStyleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:86
// index:0
// void setSegmentStyle(enum QLCDNumber::SegmentStyle)
func (this *QLCDNumber) SetSegmentStyle(arg0 int) {
	// 0: (, QLCDNumber::SegmentStyle arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber15setSegmentStyleENS_12SegmentStyleE", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:88
// index:0
// double value()
func (this *QLCDNumber) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:89
// index:0
// int intValue()
func (this *QLCDNumber) IntValue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber8intValueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:91
// index:0
// virtual
// QSize sizeHint()
func (this *QLCDNumber) SizeHint() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QLCDNumber8sizeHintEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:94
// index:0
// void display(const class QString &)
func (this *QLCDNumber) Display(str unsafe.Pointer) {
	// 0: (, const QString & str), (str)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, str)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:95
// index:1
// void display(int)
func (this *QLCDNumber) Display_1(num int) {
	// 1: (, int num), (&num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayEi", ffiqt.FFI_TYPE_VOID, this.cthis, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:96
// index:2
// void display(double)
func (this *QLCDNumber) Display_2(num float64) {
	// 2: (, double num), (&num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber7displayEd", ffiqt.FFI_TYPE_VOID, this.cthis, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:97
// index:0
// void setHexMode()
func (this *QLCDNumber) SetHexMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setHexModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:98
// index:0
// void setDecMode()
func (this *QLCDNumber) SetDecMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setDecModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:99
// index:0
// void setOctMode()
func (this *QLCDNumber) SetOctMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setOctModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:100
// index:0
// void setBinMode()
func (this *QLCDNumber) SetBinMode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber10setBinModeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:101
// index:0
// void setSmallDecimalPoint(_Bool)
func (this *QLCDNumber) SetSmallDecimalPoint(arg0 bool) {
	// 0: (, bool arg0), (&arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber20setSmallDecimalPointEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qlcdnumber.h:104
// index:0
// void overflow()
func (this *QLCDNumber) Overflow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QLCDNumber8overflowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
