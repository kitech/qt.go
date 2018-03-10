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
// extern C begin: 31
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
type QTextTableFormat struct {
	*QTextFrameFormat
}
type QTextTableFormat_ITF interface {
	QTextFrameFormat_ITF
	QTextTableFormat_PTR() *QTextTableFormat
}

func (ptr *QTextTableFormat) QTextTableFormat_PTR() *QTextTableFormat { return ptr }

func (this *QTextTableFormat) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFrameFormat.GetCthis()
	}
}
func (this *QTextTableFormat) SetCthis(cthis unsafe.Pointer) {
	this.QTextFrameFormat = NewQTextFrameFormatFromPointer(cthis)
}
func NewQTextTableFormatFromPointer(cthis unsafe.Pointer) *QTextTableFormat {
	bcthis0 := NewQTextFrameFormatFromPointer(cthis)
	return &QTextTableFormat{bcthis0}
}
func (*QTextTableFormat) NewFromPointer(cthis unsafe.Pointer) *QTextTableFormat {
	return NewQTextTableFormatFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextformat.h:887
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextTableFormat()

/*

 */
func NewQTextTableFormat() *QTextTableFormat {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormatC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:923
// index:1
// Protected Visibility=Default Availability=Available
// [-2] void QTextTableFormat(const QTextFormat &)

/*

 */
func NewQTextTableFormat_1(fmt_ QTextFormat_ITF) *QTextTableFormat {
	var convArg0 unsafe.Pointer
	if fmt_ != nil && fmt_.QTextFormat_PTR() != nil {
		convArg0 = fmt_.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormatC2ERK11QTextFormat", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableFormat)
	return gothis
}

// /usr/include/qt/QtGui/qtextformat.h:889
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const

/*
Returns true if the format is valid (i.e. is not InvalidFormat); otherwise returns false.
*/
func (this *QTextTableFormat) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextTableFormat7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextformat.h:891
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int columns() const

/*

 */
func (this *QTextTableFormat) Columns() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextTableFormat7columnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:893
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setColumns(int)

/*

 */
func (this *QTextTableFormat) SetColumns(columns int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormat10setColumnsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), columns)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:901
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void clearColumnWidthConstraints()

/*

 */
func (this *QTextTableFormat) ClearColumnWidthConstraints() {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormat27clearColumnWidthConstraintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:904
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal cellSpacing() const

/*

 */
func (this *QTextTableFormat) CellSpacing() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextTableFormat11cellSpacingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:906
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCellSpacing(qreal)

/*

 */
func (this *QTextTableFormat) SetCellSpacing(spacing float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormat14setCellSpacingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), spacing)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:909
// index:0
// Public inline Visibility=Default Availability=Available
// [8] qreal cellPadding() const

/*

 */
func (this *QTextTableFormat) CellPadding() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextTableFormat11cellPaddingEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qtextformat.h:911
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setCellPadding(qreal)

/*

 */
func (this *QTextTableFormat) SetCellPadding(padding float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormat14setCellPaddingEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), padding)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:913
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAlignment(Qt::Alignment)

/*

 */
func (this *QTextTableFormat) SetAlignment(alignment int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormat12setAlignmentE6QFlagsIN2Qt13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), alignment)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:914
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::Alignment alignment() const

/*

 */
func (this *QTextTableFormat) Alignment() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextTableFormat9alignmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtextformat.h:917
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setHeaderRowCount(int)

/*

 */
func (this *QTextTableFormat) SetHeaderRowCount(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormat17setHeaderRowCountEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:919
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int headerRowCount() const

/*

 */
func (this *QTextTableFormat) HeaderRowCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QTextTableFormat14headerRowCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

func DeleteQTextTableFormat(this *QTextTableFormat) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QTextTableFormatD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
