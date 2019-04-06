package qtgui

// /usr/include/qt/QtGui/qtextformat.h
// #include <qtextformat.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
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
type QTextTableCellFormat struct {
	*QTextCharFormat
}
type QTextTableCellFormat_ITF interface {
	QTextCharFormat_ITF
	QTextTableCellFormat_PTR() *QTextTableCellFormat
}

func (ptr *QTextTableCellFormat) QTextTableCellFormat_PTR() *QTextTableCellFormat { return ptr }

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

// /usr/include/qt/QtGui/qtextformat.h:959
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextTableCellFormat()

/*

 */
func (*QTextTableCellFormat) NewForInherit() *QTextTableCellFormat {
	return NewQTextTableCellFormat()
}
func NewQTextTableCellFormat() *QTextTableCellFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableCellFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:978
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextTableCellFormat(const QTextFormat &)

/*

 */
func (*QTextTableCellFormat) NewForInherit1(fmt_ QTextFormat_ITF) *QTextTableCellFormat {
	return NewQTextTableCellFormat1(fmt_)
}
func NewQTextTableCellFormat1(fmt_ QTextFormat_ITF) *QTextTableCellFormat {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QTextFormat_PTR() != nil {
		convArg0 = fmt_.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextTableCellFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableCellFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:961
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextTableCellFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QTextTableCellFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:963
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setTopPadding(qreal)

/*

 */
func (this *QTextTableCellFormat) SetTopPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormat13setTopPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:964
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal topPadding() const

/*

 */
func (this *QTextTableCellFormat) TopPadding() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QTextTableCellFormat10topPaddingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:966
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setBottomPadding(qreal)

/*

 */
func (this *QTextTableCellFormat) SetBottomPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormat16setBottomPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:967
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal bottomPadding() const

/*

 */
func (this *QTextTableCellFormat) BottomPadding() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QTextTableCellFormat13bottomPaddingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:969
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setLeftPadding(qreal)

/*

 */
func (this *QTextTableCellFormat) SetLeftPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormat14setLeftPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:970
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal leftPadding() const

/*

 */
func (this *QTextTableCellFormat) LeftPadding() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QTextTableCellFormat11leftPaddingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:972
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setRightPadding(qreal)

/*

 */
func (this *QTextTableCellFormat) SetRightPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormat15setRightPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:973
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal rightPadding() const

/*

 */
func (this *QTextTableCellFormat) RightPadding() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QTextTableCellFormat12rightPaddingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:975
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setPadding(qreal)

/*

 */
func (this *QTextTableCellFormat) SetPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormat10setPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

func DeleteQTextTableCellFormat(this *QTextTableCellFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QTextTableCellFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_10775() {
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
