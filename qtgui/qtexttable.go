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
// extern C begin: 19
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

type QTextTable struct {
	*QTextFrame
}
type QTextTable_ITF interface {
	QTextFrame_ITF
	QTextTable_PTR() *QTextTable
}

func (ptr *QTextTable) QTextTable_PTR() *QTextTable { return ptr }

func (this *QTextTable) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextFrame.GetCthis()
	}
}
func (this *QTextTable) SetCthis(cthis unsafe.Pointer) {
	this.QTextFrame = NewQTextFrameFromPointer(cthis)
}
func NewQTextTableFromPointer(cthis unsafe.Pointer) *QTextTable {
	bcthis0 := NewQTextFrameFromPointer(cthis)
	return &QTextTable{bcthis0}
}
func (*QTextTable) NewFromPointer(cthis unsafe.Pointer) *QTextTable {
	return NewQTextTableFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtexttable.h:100
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTextTable) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtexttable.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextTable(QTextDocument *)
func NewQTextTable(doc QTextDocument_ITF /*777 QTextDocument **/) *QTextTable {
	var convArg0 unsafe.Pointer
	if doc != nil && doc.QTextDocument_PTR() != nil {
		convArg0 = doc.QTextDocument_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTableC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextTableFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTextTable")
	return gothis
}

// /usr/include/qt/QtGui/qtexttable.h:103
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextTable()
func DeleteQTextTable(this *QTextTable) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTableD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtexttable.h:105
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(int, int)
func (this *QTextTable) Resize(rows int, cols int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable6resizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), rows, cols)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertRows(int, int)
func (this *QTextTable) InsertRows(pos int, num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable10insertRowsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertColumns(int, int)
func (this *QTextTable) InsertColumns(pos int, num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable13insertColumnsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendRows(int)
func (this *QTextTable) AppendRows(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable10appendRowsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void appendColumns(int)
func (this *QTextTable) AppendColumns(count int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable13appendColumnsEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), count)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeRows(int, int)
func (this *QTextTable) RemoveRows(pos int, num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable10removeRowsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeColumns(int, int)
func (this *QTextTable) RemoveColumns(pos int, num int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable13removeColumnsEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), pos, num)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mergeCells(int, int, int, int)
func (this *QTextTable) MergeCells(row int, col int, numRows int, numCols int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable10mergeCellsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, col, numRows, numCols)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:114
// index:1
// Public Visibility=Default Availability=Available
// [-2] void mergeCells(const QTextCursor &)
func (this *QTextTable) MergeCells_1(cursor QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if cursor != nil && cursor.QTextCursor_PTR() != nil {
		convArg0 = cursor.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable10mergeCellsERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void splitCell(int, int, int, int)
func (this *QTextTable) SplitCell(row int, col int, numRows int, numCols int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable9splitCellEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, col, numRows, numCols)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:117
// index:0
// Public Visibility=Default Availability=Available
// [4] int rows()
func (this *QTextTable) Rows() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable4rowsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] int columns()
func (this *QTextTable) Columns() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable7columnsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtexttable.h:120
// index:0
// Public Visibility=Default Availability=Available
// [16] QTextTableCell cellAt(int, int)
func (this *QTextTable) CellAt(row int, col int) *QTextTableCell /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable6cellAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), row, col)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableCellFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCell)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:121
// index:1
// Public Visibility=Default Availability=Available
// [16] QTextTableCell cellAt(int)
func (this *QTextTable) CellAt_1(position int) *QTextTableCell /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable6cellAtEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), position)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableCellFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCell)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:122
// index:2
// Public Visibility=Default Availability=Available
// [16] QTextTableCell cellAt(const QTextCursor &)
func (this *QTextTable) CellAt_2(c QTextCursor_ITF) *QTextTableCell /*123*/ {
	var convArg0 unsafe.Pointer
	if c != nil && c.QTextCursor_PTR() != nil {
		convArg0 = c.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable6cellAtERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableCellFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableCell)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor rowStart(const QTextCursor &)
func (this *QTextTable) RowStart(c QTextCursor_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if c != nil && c.QTextCursor_PTR() != nil {
		convArg0 = c.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable8rowStartERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor rowEnd(const QTextCursor &)
func (this *QTextTable) RowEnd(c QTextCursor_ITF) *QTextCursor /*123*/ {
	var convArg0 unsafe.Pointer
	if c != nil && c.QTextCursor_PTR() != nil {
		convArg0 = c.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable6rowEndERK11QTextCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtexttable.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QTextTableFormat &)
func (this *QTextTable) SetFormat(format QTextTableFormat_ITF) {
	var convArg0 unsafe.Pointer
	if format != nil && format.QTextTableFormat_PTR() != nil {
		convArg0 = format.QTextTableFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextTable9setFormatERK16QTextTableFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:128
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextTableFormat format()
func (this *QTextTable) Format() *QTextTableFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextTable6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextTableFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextTableFormat)
	return rv2
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
