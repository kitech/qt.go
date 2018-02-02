package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
}

//  ext block end

//  body block begin

type QTextTableCellFormat struct {
	*QTextCharFormat
}

func (this *QTextTableCellFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextCharFormat.GetCthis()
	}
}
func (this *QTextTableCellFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextCharFormat = NewQTextCharFormatFromPointer(cthis)
}
func NewQTextTableCellFormatFromPointer(cthis unsafe.Pointer) *QTextTableCellFormat {
	bcthis0 := NewQTextCharFormatFromPointer(cthis)
	return &QTextTableCellFormat{bcthis0}
}
func (*QTextTableCellFormat) NewFromPointer(cthis unsafe.Pointer) *QTextTableCellFormat {
	return NewQTextTableCellFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:945
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextTableCellFormat()
func NewQTextTableCellFormat() *QTextTableCellFormat {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormatC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableCellFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:964
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextTableCellFormat(const QTextFormat &)
func NewQTextTableCellFormat_1(fmt *QTextFormat) *QTextTableCellFormat {
	var convArg0 = fmt.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormatC2ERK11QTextFormat", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableCellFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:947
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QTextTableCellFormat) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:949
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopPadding(qreal)
func (this *QTextTableCellFormat) SetTopPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat13setTopPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:950
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal topPadding()
func (this *QTextTableCellFormat) TopPadding() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat10topPaddingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:952
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomPadding(qreal)
func (this *QTextTableCellFormat) SetBottomPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat16setBottomPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:953
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottomPadding()
func (this *QTextTableCellFormat) BottomPadding() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat13bottomPaddingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:955
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeftPadding(qreal)
func (this *QTextTableCellFormat) SetLeftPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat14setLeftPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:956
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal leftPadding()
func (this *QTextTableCellFormat) LeftPadding() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat11leftPaddingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:958
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRightPadding(qreal)
func (this *QTextTableCellFormat) SetRightPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat15setRightPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:959
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rightPadding()
func (this *QTextTableCellFormat) RightPadding() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK20QTextTableCellFormat12rightPaddingEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:961
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPadding(qreal)
func (this *QTextTableCellFormat) SetPadding(padding float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormat10setPaddingEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	gopp.ErrPrint(err, rv)
}

func DeleteQTextTableCellFormat(this *QTextTableCellFormat) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextTableCellFormatD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end
