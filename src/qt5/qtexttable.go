package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
extern void* dector_ZN14QTextTableCellC1EPK10QTextTablei(void* arg0, int arg1);
extern void _ZN14QTextTableCellC1EPK10QTextTablei(void* qthis, void* arg0, int arg1);
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
extern void _ZNK10QTextTable6cellAtEii(void* qthis, int arg0, int arg1);
  // proto:  int QTextTable::rows();
extern void _ZNK10QTextTable4rowsEv(void* qthis);
  // proto:  void QTextTable::removeRows(int pos, int num);
extern void _ZN10QTextTable10removeRowsEii(void* qthis, int arg0, int arg1);
  // proto:  int QTextTable::columns();
extern void _ZNK10QTextTable7columnsEv(void* qthis);
  // proto:  void QTextTable::appendRows(int count);
extern void _ZN10QTextTable10appendRowsEi(void* qthis, int arg0);
  // proto:  void QTextTable::resize(int rows, int cols);
extern void _ZN10QTextTable6resizeEii(void* qthis, int arg0, int arg1);
  // proto:  QTextTableCell QTextTable::cellAt(const QTextCursor & c);
extern void _ZNK10QTextTable6cellAtERK11QTextCursor(void* qthis, void* arg0);
  // proto:  void QTextTable::QTextTable(const QTextTable & );
extern void* dector_ZN10QTextTableC1ERKS_(void* arg0);
extern void _ZN10QTextTableC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextTable::setFormat(const QTextTableFormat & format);
extern void _ZN10QTextTable9setFormatERK16QTextTableFormat(void* qthis, void* arg0);
  // proto:  void QTextTable::insertColumns(int pos, int num);
extern void _ZN10QTextTable13insertColumnsEii(void* qthis, int arg0, int arg1);
  // proto:  void QTextTable::splitCell(int row, int col, int numRows, int numCols);
extern void _ZN10QTextTable9splitCellEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QTextTable::mergeCells(int row, int col, int numRows, int numCols);
extern void _ZN10QTextTable10mergeCellsEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QTextTable::insertRows(int pos, int num);
extern void _ZN10QTextTable10insertRowsEii(void* qthis, int arg0, int arg1);
  // proto:  void QTextTable::~QTextTable();
extern void _ZN10QTextTableD0Ev(void* qthis);
  // proto:  void QTextTable::QTextTable(QTextDocument * doc);
extern void* dector_ZN10QTextTableC1EP13QTextDocument(void* arg0);
extern void _ZN10QTextTableC1EP13QTextDocument(void* qthis, void* arg0);
  // proto:  QTextTableCell QTextTable::cellAt(int position);
extern void _ZNK10QTextTable6cellAtEi(void* qthis, int arg0);
  // proto:  QTextCursor QTextTable::rowStart(const QTextCursor & c);
extern void _ZNK10QTextTable8rowStartERK11QTextCursor(void* qthis, void* arg0);
  // proto:  QTextTableFormat QTextTable::format();
extern void _ZNK10QTextTable6formatEv(void* qthis);
  // proto:  QTextCursor QTextTable::rowEnd(const QTextCursor & c);
extern void _ZNK10QTextTable6rowEndERK11QTextCursor(void* qthis, void* arg0);
  // proto:  const QMetaObject * QTextTable::metaObject();
extern void _ZNK10QTextTable10metaObjectEv(void* qthis);
  // proto:  void QTextTable::removeColumns(int pos, int num);
extern void _ZN10QTextTable13removeColumnsEii(void* qthis, int arg0, int arg1);
  // proto:  void QTextTable::appendColumns(int count);
extern void _ZN10QTextTable13appendColumnsEi(void* qthis, int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextTable)=1
type QTextTable struct {
  /*qbase*/ QTextFrame;
  qclsinst uint64 /* *mut c_void*/;
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
    var arg0 = args[0].(QTextCharFormat).qclsinst
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZNK10QTextTable6cellAtERK11QTextCursor
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
  case 2:
    // invoke: _ZNK10QTextTable6cellAtEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
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
    var arg0 = args[0].(QTextTableFormat).qclsinst
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
  case 1:
    // invoke: _ZN10QTextTable10mergeCellsERK11QTextCursor
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
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
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
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
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextTable", "appendColumns", args)
  }

}

// <= body block end

