//  header block begin
// /usr/include/qt/QtGui/qtexttable.h
// #include <qtexttable.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
type QTextTable struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qtexttable.h:100
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QTextTable) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:102
// index:0
// void QTextTable(class QTextDocument *)
func NewQTextTable(doc unsafe.Pointer) *QTextTable {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTableC2EP13QTextDocument", ffiqt.FFI_TYPE_VOID, cthis, doc)
	gopp.ErrPrint(err, rv)
	return &QTextTable{cthis}
}

// /usr/include/qt/QtGui/qtexttable.h:103
// index:0
// virtual
// void ~QTextTable()
func DeleteQTextTable(*QTextTable) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTableD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:105
// index:0
// void resize(int, int)
func (this *QTextTable) Resize(rows int, cols int) {
	// 0: (, int rows, int cols), (&rows, &cols)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable6resizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &rows, &cols)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:106
// index:0
// void insertRows(int, int)
func (this *QTextTable) InsertRows(pos int, num int) {
	// 0: (, int pos, int num), (&pos, &num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable10insertRowsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:107
// index:0
// void insertColumns(int, int)
func (this *QTextTable) InsertColumns(pos int, num int) {
	// 0: (, int pos, int num), (&pos, &num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable13insertColumnsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:108
// index:0
// void appendRows(int)
func (this *QTextTable) AppendRows(count int) {
	// 0: (, int count), (&count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable10appendRowsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:109
// index:0
// void appendColumns(int)
func (this *QTextTable) AppendColumns(count int) {
	// 0: (, int count), (&count)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable13appendColumnsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &count)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:110
// index:0
// void removeRows(int, int)
func (this *QTextTable) RemoveRows(pos int, num int) {
	// 0: (, int pos, int num), (&pos, &num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable10removeRowsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:111
// index:0
// void removeColumns(int, int)
func (this *QTextTable) RemoveColumns(pos int, num int) {
	// 0: (, int pos, int num), (&pos, &num)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable13removeColumnsEii", ffiqt.FFI_TYPE_VOID, this.cthis, &pos, &num)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:113
// index:0
// void mergeCells(int, int, int, int)
func (this *QTextTable) MergeCells(row int, col int, numRows int, numCols int) {
	// 0: (, int row, int col, int numRows, int numCols), (&row, &col, &numRows, &numCols)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable10mergeCellsEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &col, &numRows, &numCols)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:114
// index:1
// void mergeCells(const class QTextCursor &)
func (this *QTextTable) MergeCells_1(cursor unsafe.Pointer) {
	// 1: (, const QTextCursor & cursor), (cursor)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable10mergeCellsERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, cursor)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:115
// index:0
// void splitCell(int, int, int, int)
func (this *QTextTable) SplitCell(row int, col int, numRows int, numCols int) {
	// 0: (, int row, int col, int numRows, int numCols), (&row, &col, &numRows, &numCols)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable9splitCellEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &col, &numRows, &numCols)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:117
// index:0
// int rows()
func (this *QTextTable) Rows() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable4rowsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:118
// index:0
// int columns()
func (this *QTextTable) Columns() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable7columnsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:120
// index:0
// QTextTableCell cellAt(int, int)
func (this *QTextTable) CellAt(row int, col int) {
	// 0: (, int row, int col), (&row, &col)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable6cellAtEii", ffiqt.FFI_TYPE_VOID, this.cthis, &row, &col)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:121
// index:1
// QTextTableCell cellAt(int)
func (this *QTextTable) CellAt_1(position int) {
	// 1: (, int position), (&position)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable6cellAtEi", ffiqt.FFI_TYPE_VOID, this.cthis, &position)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:122
// index:2
// QTextTableCell cellAt(const class QTextCursor &)
func (this *QTextTable) CellAt_2(c unsafe.Pointer) {
	// 2: (, const QTextCursor & c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable6cellAtERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:124
// index:0
// QTextCursor rowStart(const class QTextCursor &)
func (this *QTextTable) RowStart(c unsafe.Pointer) {
	// 0: (, const QTextCursor & c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable8rowStartERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:125
// index:0
// QTextCursor rowEnd(const class QTextCursor &)
func (this *QTextTable) RowEnd(c unsafe.Pointer) {
	// 0: (, const QTextCursor & c), (c)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable6rowEndERK11QTextCursor", ffiqt.FFI_TYPE_VOID, this.cthis, c)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:127
// index:0
// void setFormat(const class QTextTableFormat &)
func (this *QTextTable) SetFormat(format unsafe.Pointer) {
	// 0: (, const QTextTableFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QTextTable9setFormatERK16QTextTableFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtexttable.h:128
// index:0
// inline
// QTextTableFormat format()
func (this *QTextTable) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QTextTable6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

//  body block end
