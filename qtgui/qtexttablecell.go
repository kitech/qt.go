package qtgui

// /usr/include/qt/QtGui/qtexttable.h
// #include <qtexttable.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 13
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

type QTextTableCell struct {
	*qtrt.CObject
}
type QTextTableCell_ITF interface {
	QTextTableCell_PTR() *QTextTableCell
}

func (ptr *QTextTableCell) QTextTableCell_PTR() *QTextTableCell { return ptr }

func (this *QTextTableCell) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextTableCell) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextTableCellFromPointer(cthis unsafe.Pointer) *QTextTableCell {
	return &QTextTableCell{&qtrt.CObject{cthis}}
}
func (*QTextTableCell) NewFromPointer(cthis unsafe.Pointer) *QTextTableCell {
	return NewQTextTableCellFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtexttable.h:57
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QTextTableCell()
func NewQTextTableCell() *QTextTableCell {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTextTableCellC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextTableCellFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextTableCell)
	return gothis
}

// /usr/include/qt/QtGui/qtexttable.h:58
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ~QTextTableCell()
func DeleteQTextTableCell(this *QTextTableCell) {
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTextTableCellD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtexttable.h:60
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextTableCell & operator=(const QTextTableCell &)
func (this *QTextTableCell) Operator_equal(o QTextTableCell_ITF) *QTextTableCell {
	var convArg0 unsafe.Pointer
	if o != nil && o.QTextTableCell_PTR() != nil {
		convArg0 = o.QTextTableCell_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTextTableCellaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableCellFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCell)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:63
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QTextCharFormat &)
func (this *QTextTableCell) SetFormat(format QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextCharFormat_PTR() != nil {
		convArg0 = format.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN14QTextTableCell9setFormatERK15QTextCharFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:64
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextCharFormat format() const
func (this *QTextTableCell) Format() *QTextCharFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCharFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCharFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:66
// index:0
// Public Visibility=Default Availability=Available
// [4] int row() const
func (this *QTextTableCell) Row() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell3rowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:67
// index:0
// Public Visibility=Default Availability=Available
// [4] int column() const
func (this *QTextTableCell) Column() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell6columnEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:69
// index:0
// Public Visibility=Default Availability=Available
// [4] int rowSpan() const
func (this *QTextTableCell) RowSpan() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell7rowSpanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:70
// index:0
// Public Visibility=Default Availability=Available
// [4] int columnSpan() const
func (this *QTextTableCell) ColumnSpan() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell10columnSpanEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:72
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QTextTableCell) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtexttable.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor firstCursorPosition() const
func (this *QTextTableCell) FirstCursorPosition() *QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell19firstCursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:75
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor lastCursorPosition() const
func (this *QTextTableCell) LastCursorPosition() *QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell18lastCursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int firstPosition() const
func (this *QTextTableCell) FirstPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell13firstPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastPosition() const
func (this *QTextTableCell) LastPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell12lastPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:79
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator==(const QTextTableCell &) const
func (this *QTextTableCell) Operator_equal_equal(other QTextTableCell_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextTableCell_PTR() != nil {
		convArg0 = other.QTextTableCell_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCelleqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtexttable.h:81
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QTextTableCell &) const
func (this *QTextTableCell) Operator_not_equal(other QTextTableCell_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QTextTableCell_PTR() != nil {
		convArg0 = other.QTextTableCell_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCellneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtexttable.h:84
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextFrame::iterator begin() const
func (this *QTextTableCell) Begin() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtexttable.h:85
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextFrame::iterator end() const
func (this *QTextTableCell) End() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qtexttable.h:87
// index:0
// Public Visibility=Default Availability=Available
// [4] int tableCellFormatIndex() const
func (this *QTextTableCell) TableCellFormatIndex() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK14QTextTableCell20tableCellFormatIndexEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
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
