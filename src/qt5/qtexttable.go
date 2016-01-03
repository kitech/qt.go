package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
// src-file: /QtGui/qtexttable.h
// dst-file: /src/gui/qtexttable.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTextTableCell::QTextTableCell(const QTextTableCell & o);
extern void* dector_ZN14QTextTableCellC1ERKS_(void* arg0);
extern void _ZN14QTextTableCellC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextTableCell::setFormat(const QTextCharFormat & format);
extern void _ZN14QTextTableCell9setFormatERK15QTextCharFormat(void* qthis, void* arg0);
  // proto:  int QTextTableCell::lastPosition();
extern void _ZNK14QTextTableCell12lastPositionEv(void* qthis);
  // proto:  void QTextTableCell::~QTextTableCell();
extern void _ZN14QTextTableCellD0Ev(void* qthis);
  // proto:  int QTextTableCell::rowSpan();
extern void _ZNK14QTextTableCell7rowSpanEv(void* qthis);
  // proto:  int QTextTableCell::firstPosition();
extern void _ZNK14QTextTableCell13firstPositionEv(void* qthis);
  // proto:  void QTextTableCell::QTextTableCell(const QTextTable * t, int f);
extern void* dector_ZN14QTextTableCellC1EPK10QTextTablei(void* arg0, int32_t arg1);
extern void _ZN14QTextTableCellC1EPK10QTextTablei(void* qthis, void* arg0, int32_t arg1);
  // proto:  int QTextTableCell::tableCellFormatIndex();
extern void _ZNK14QTextTableCell20tableCellFormatIndexEv(void* qthis);
  // proto:  int QTextTableCell::columnSpan();
extern void _ZNK14QTextTableCell10columnSpanEv(void* qthis);
  // proto:  QTextCharFormat QTextTableCell::format();
extern void _ZNK14QTextTableCell6formatEv(void* qthis);
  // proto:  int QTextTableCell::row();
extern void _ZNK14QTextTableCell3rowEv(void* qthis);
  // proto:  bool QTextTableCell::isValid();
extern void demth_ZNK14QTextTableCell7isValidEv(void* qthis);
  // proto:  QTextCursor QTextTableCell::lastCursorPosition();
extern void _ZNK14QTextTableCell18lastCursorPositionEv(void* qthis);
  // proto:  int QTextTableCell::column();
extern void _ZNK14QTextTableCell6columnEv(void* qthis);
  // proto:  QTextCursor QTextTableCell::firstCursorPosition();
extern void _ZNK14QTextTableCell19firstCursorPositionEv(void* qthis);
  // proto:  void QTextTableCell::QTextTableCell();
extern void* dector_ZN14QTextTableCellC1Ev();
extern void _ZN14QTextTableCellC1Ev(void* qthis);
  // proto:  QTextTableCell QTextTable::cellAt(int row, int col);
extern void _ZNK10QTextTable6cellAtEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  int QTextTable::rows();
extern void _ZNK10QTextTable4rowsEv(void* qthis);
  // proto:  void QTextTable::removeRows(int pos, int num);
extern void _ZN10QTextTable10removeRowsEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  int QTextTable::columns();
extern void _ZNK10QTextTable7columnsEv(void* qthis);
  // proto:  void QTextTable::appendRows(int count);
extern void _ZN10QTextTable10appendRowsEi(void* qthis, int32_t arg0);
  // proto:  void QTextTable::resize(int rows, int cols);
extern void _ZN10QTextTable6resizeEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QTextTableCell QTextTable::cellAt(const QTextCursor & c);
extern void _ZNK10QTextTable6cellAtERK11QTextCursor(void* qthis, void* arg0);
  // proto:  void QTextTable::QTextTable(const QTextTable & );
extern void* dector_ZN10QTextTableC1ERKS_(void* arg0);
extern void _ZN10QTextTableC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextTable::setFormat(const QTextTableFormat & format);
extern void _ZN10QTextTable9setFormatERK16QTextTableFormat(void* qthis, void* arg0);
  // proto:  void QTextTable::insertColumns(int pos, int num);
extern void _ZN10QTextTable13insertColumnsEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QTextTable::splitCell(int row, int col, int numRows, int numCols);
extern void _ZN10QTextTable9splitCellEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QTextTable::mergeCells(int row, int col, int numRows, int numCols);
extern void _ZN10QTextTable10mergeCellsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto:  void QTextTable::insertRows(int pos, int num);
extern void _ZN10QTextTable10insertRowsEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QTextTable::~QTextTable();
extern void _ZN10QTextTableD0Ev(void* qthis);
  // proto:  void QTextTable::QTextTable(QTextDocument * doc);
extern void* dector_ZN10QTextTableC1EP13QTextDocument(void* arg0);
extern void _ZN10QTextTableC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  QTextTableCell QTextTable::cellAt(int position);
extern void _ZNK10QTextTable6cellAtEi(void* qthis, int32_t arg0);
  // proto:  QTextCursor QTextTable::rowStart(const QTextCursor & c);
extern void _ZNK10QTextTable8rowStartERK11QTextCursor(void* qthis, void* arg0);
  // proto:  QTextTableFormat QTextTable::format();
extern void demth_ZNK10QTextTable6formatEv(void* qthis);
  // proto:  QTextCursor QTextTable::rowEnd(const QTextCursor & c);
extern void _ZNK10QTextTable6rowEndERK11QTextCursor(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTextTable::metaObject();
extern void _ZNK10QTextTable10metaObjectEv(void* qthis);
  // proto:  void QTextTable::removeColumns(int pos, int num);
extern void _ZN10QTextTable13removeColumnsEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QTextTable::appendColumns(int count);
extern void _ZN10QTextTable13appendColumnsEi(void* qthis, int32_t arg0);
  // proto:  void QTextTable::mergeCells(const QTextCursor & cursor);
extern void _ZN10QTextTable10mergeCellsERK11QTextCursor(void* qthis, void* arg0);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextTableCell)=16
type QTextTableCell struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextTable)=1
type QTextTable struct {
  /*qbase*/ QTextFrame;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QTextTableCell::QTextTableCell(const QTextTableCell & o);
func NewQTextTableCell(args ...interface{}) QTextTableCell {
  return QTextTableCell{}
}

  // proto:  void QTextTableCell::setFormat(const QTextCharFormat & format);
func (this *QTextTableCell) setFormat(args ...interface{}) () {
  // setFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTextTableCell9setFormatERK15QTextCharFormat
    // invoke: void setFormat(const class QTextCharFormat &)
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN14QTextTableCell9setFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCell", "setFormat", args)
  }

}

  // proto:  int QTextTableCell::lastPosition();
func (this *QTextTableCell) lastPosition(args ...interface{}) () {
  // lastPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell12lastPositionEv
    // invoke: int lastPosition()
    C._ZNK14QTextTableCell12lastPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "lastPosition", args)
  }

}

  // proto:  void QTextTableCell::~QTextTableCell();
func (this *QTextTableCell) FreeQTextTableCell(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextTableCell", "~QTextTableCell", args)
  }

}

  // proto:  int QTextTableCell::rowSpan();
func (this *QTextTableCell) rowSpan(args ...interface{}) () {
  // rowSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell7rowSpanEv
    // invoke: int rowSpan()
    C._ZNK14QTextTableCell7rowSpanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "rowSpan", args)
  }

}

  // proto:  int QTextTableCell::firstPosition();
func (this *QTextTableCell) firstPosition(args ...interface{}) () {
  // firstPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell13firstPositionEv
    // invoke: int firstPosition()
    C._ZNK14QTextTableCell13firstPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "firstPosition", args)
  }

}

  // proto:  int QTextTableCell::tableCellFormatIndex();
func (this *QTextTableCell) tableCellFormatIndex(args ...interface{}) () {
  // tableCellFormatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell20tableCellFormatIndexEv
    // invoke: int tableCellFormatIndex()
    C._ZNK14QTextTableCell20tableCellFormatIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "tableCellFormatIndex", args)
  }

}

  // proto:  int QTextTableCell::columnSpan();
func (this *QTextTableCell) columnSpan(args ...interface{}) () {
  // columnSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell10columnSpanEv
    // invoke: int columnSpan()
    C._ZNK14QTextTableCell10columnSpanEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "columnSpan", args)
  }

}

  // proto:  QTextCharFormat QTextTableCell::format();
func (this *QTextTableCell) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell6formatEv
    // invoke: QTextCharFormat format()
    C._ZNK14QTextTableCell6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "format", args)
  }

}

  // proto:  int QTextTableCell::row();
func (this *QTextTableCell) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell3rowEv
    // invoke: int row()
    C._ZNK14QTextTableCell3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "row", args)
  }

}

  // proto:  bool QTextTableCell::isValid();
func (this *QTextTableCell) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell7isValidEv
    // invoke: bool isValid()
    C.demth_ZNK14QTextTableCell7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "isValid", args)
  }

}

  // proto:  QTextCursor QTextTableCell::lastCursorPosition();
func (this *QTextTableCell) lastCursorPosition(args ...interface{}) () {
  // lastCursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell18lastCursorPositionEv
    // invoke: QTextCursor lastCursorPosition()
    C._ZNK14QTextTableCell18lastCursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "lastCursorPosition", args)
  }

}

  // proto:  int QTextTableCell::column();
func (this *QTextTableCell) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell6columnEv
    // invoke: int column()
    C._ZNK14QTextTableCell6columnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "column", args)
  }

}

  // proto:  QTextCursor QTextTableCell::firstCursorPosition();
func (this *QTextTableCell) firstCursorPosition(args ...interface{}) () {
  // firstCursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell19firstCursorPositionEv
    // invoke: QTextCursor firstCursorPosition()
    C._ZNK14QTextTableCell19firstCursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "firstCursorPosition", args)
  }

}

  // proto:  QTextTableCell QTextTable::cellAt(int row, int col);
func (this *QTextTable) cellAt(args ...interface{}) () {
  // cellAt(int, int)
  // cellAt(const class QTextCursor &)
  // cellAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6cellAtEii
    // invoke: QTextTableCell cellAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QTextTable6cellAtEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK10QTextTable6cellAtERK11QTextCursor
    // invoke: QTextTableCell cellAt(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextTable6cellAtERK11QTextCursor(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QTextTable6cellAtEi
    // invoke: QTextTableCell cellAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QTextTable6cellAtEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "cellAt", args)
  }

}

  // proto:  int QTextTable::rows();
func (this *QTextTable) rows(args ...interface{}) () {
  // rows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable4rowsEv
    // invoke: int rows()
    C._ZNK10QTextTable4rowsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "rows", args)
  }

}

  // proto:  void QTextTable::removeRows(int pos, int num);
func (this *QTextTable) removeRows(args ...interface{}) () {
  // removeRows(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10removeRowsEii
    // invoke: void removeRows(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTextTable10removeRowsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "removeRows", args)
  }

}

  // proto:  int QTextTable::columns();
func (this *QTextTable) columns(args ...interface{}) () {
  // columns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable7columnsEv
    // invoke: int columns()
    C._ZNK10QTextTable7columnsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "columns", args)
  }

}

  // proto:  void QTextTable::appendRows(int count);
func (this *QTextTable) appendRows(args ...interface{}) () {
  // appendRows(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10appendRowsEi
    // invoke: void appendRows(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTextTable10appendRowsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "appendRows", args)
  }

}

  // proto:  void QTextTable::resize(int rows, int cols);
func (this *QTextTable) resize(args ...interface{}) () {
  // resize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTextTable6resizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "resize", args)
  }

}

  // proto:  void QTextTable::QTextTable(const QTextTable & );
func NewQTextTable(args ...interface{}) QTextTable {
  return QTextTable{}
}

  // proto:  void QTextTable::setFormat(const QTextTableFormat & format);
func (this *QTextTable) setFormat(args ...interface{}) () {
  // setFormat(const class QTextTableFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextTableFormat{}) // "const QTextTableFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable9setFormatERK16QTextTableFormat
    // invoke: void setFormat(const class QTextTableFormat &)
    var arg0 = args[0].(QTextTableFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextTable9setFormatERK16QTextTableFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "setFormat", args)
  }

}

  // proto:  void QTextTable::insertColumns(int pos, int num);
func (this *QTextTable) insertColumns(args ...interface{}) () {
  // insertColumns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable13insertColumnsEii
    // invoke: void insertColumns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTextTable13insertColumnsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "insertColumns", args)
  }

}

  // proto:  void QTextTable::splitCell(int row, int col, int numRows, int numCols);
func (this *QTextTable) splitCell(args ...interface{}) () {
  // splitCell(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable9splitCellEiiii
    // invoke: void splitCell(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QTextTable9splitCellEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextTable", "splitCell", args)
  }

}

  // proto:  void QTextTable::mergeCells(int row, int col, int numRows, int numCols);
func (this *QTextTable) mergeCells(args ...interface{}) () {
  // mergeCells(int, int, int, int)
  // mergeCells(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10mergeCellsEiiii
    // invoke: void mergeCells(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QTextTable10mergeCellsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN10QTextTable10mergeCellsERK11QTextCursor
    // invoke: void mergeCells(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextTable10mergeCellsERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "mergeCells", args)
  }

}

  // proto:  void QTextTable::insertRows(int pos, int num);
func (this *QTextTable) insertRows(args ...interface{}) () {
  // insertRows(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10insertRowsEii
    // invoke: void insertRows(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTextTable10insertRowsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "insertRows", args)
  }

}

  // proto:  void QTextTable::~QTextTable();
func (this *QTextTable) FreeQTextTable(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextTable", "~QTextTable", args)
  }

}

  // proto:  QTextCursor QTextTable::rowStart(const QTextCursor & c);
func (this *QTextTable) rowStart(args ...interface{}) () {
  // rowStart(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable8rowStartERK11QTextCursor
    // invoke: QTextCursor rowStart(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextTable8rowStartERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "rowStart", args)
  }

}

  // proto:  QTextTableFormat QTextTable::format();
func (this *QTextTable) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6formatEv
    // invoke: QTextTableFormat format()
    C.demth_ZNK10QTextTable6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "format", args)
  }

}

  // proto:  QTextCursor QTextTable::rowEnd(const QTextCursor & c);
func (this *QTextTable) rowEnd(args ...interface{}) () {
  // rowEnd(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6rowEndERK11QTextCursor
    // invoke: QTextCursor rowEnd(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextTable6rowEndERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "rowEnd", args)
  }

}

  // proto:  const QMetaObject * QTextTable::metaObject();
func (this *QTextTable) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QTextTable10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "metaObject", args)
  }

}

  // proto:  void QTextTable::removeColumns(int pos, int num);
func (this *QTextTable) removeColumns(args ...interface{}) () {
  // removeColumns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable13removeColumnsEii
    // invoke: void removeColumns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QTextTable13removeColumnsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "removeColumns", args)
  }

}

  // proto:  void QTextTable::appendColumns(int count);
func (this *QTextTable) appendColumns(args ...interface{}) () {
  // appendColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable13appendColumnsEi
    // invoke: void appendColumns(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTextTable13appendColumnsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "appendColumns", args)
  }

}

// <= body block end

