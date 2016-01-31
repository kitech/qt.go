package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QTextTableCell::QTextTableCell(const QTextTableCell & o);
extern void* C_ZN14QTextTableCellC2ERKS_(void* arg0); // 1
  // proto:  void QTextTableCell::QTextTableCell();
extern void* C_ZN14QTextTableCellC2Ev(); // 1
  // proto:  int QTextTableCell::rowSpan();
extern void C_ZNK14QTextTableCell7rowSpanEv(void* qthis); // 4
  // proto:  QTextCursor QTextTableCell::firstCursorPosition();
extern void C_ZNK14QTextTableCell19firstCursorPositionEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextTableCell::format();
extern void C_ZNK14QTextTableCell6formatEv(void* qthis); // 4
  // proto:  void QTextTableCell::setFormat(const QTextCharFormat & format);
extern void C_ZN14QTextTableCell9setFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  bool QTextTableCell::isValid();
extern void C_ZNK14QTextTableCell7isValidEv(void* qthis); // 2
  // proto:  QTextFrame::iterator QTextTableCell::begin();
extern void C_ZNK14QTextTableCell5beginEv(void* qthis); // 4
  // proto:  QTextCursor QTextTableCell::lastCursorPosition();
extern void C_ZNK14QTextTableCell18lastCursorPositionEv(void* qthis); // 4
  // proto:  int QTextTableCell::tableCellFormatIndex();
extern void C_ZNK14QTextTableCell20tableCellFormatIndexEv(void* qthis); // 4
  // proto:  int QTextTableCell::column();
extern void C_ZNK14QTextTableCell6columnEv(void* qthis); // 4
  // proto:  int QTextTableCell::columnSpan();
extern void C_ZNK14QTextTableCell10columnSpanEv(void* qthis); // 4
  // proto:  QTextFrame::iterator QTextTableCell::end();
extern void C_ZNK14QTextTableCell3endEv(void* qthis); // 4
  // proto:  int QTextTableCell::lastPosition();
extern void C_ZNK14QTextTableCell12lastPositionEv(void* qthis); // 4
  // proto:  int QTextTableCell::firstPosition();
extern void C_ZNK14QTextTableCell13firstPositionEv(void* qthis); // 4
  // proto:  int QTextTableCell::row();
extern void C_ZNK14QTextTableCell3rowEv(void* qthis); // 4
  // proto:  void QTextTable::appendRows(int count);
extern void C_ZN10QTextTable10appendRowsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextTable::~QTextTable();
extern void C_ZN10QTextTableD2Ev(void* qthis); // 4
  // proto:  void QTextTable::splitCell(int row, int col, int numRows, int numCols);
extern void C_ZN10QTextTable9splitCellEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QTextTable::rows();
extern void C_ZNK10QTextTable4rowsEv(void* qthis); // 4
  // proto:  QTextTableCell QTextTable::cellAt(const QTextCursor & c);
extern void C_ZNK10QTextTable6cellAtERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QTextTableCell QTextTable::cellAt(int row, int col);
extern void C_ZNK10QTextTable6cellAtEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTextTableCell QTextTable::cellAt(int position);
extern void C_ZNK10QTextTable6cellAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextTable::setFormat(const QTextTableFormat & format);
extern void C_ZN10QTextTable9setFormatERK16QTextTableFormat(void* qthis, void* arg0); // 4
  // proto:  QTextTableFormat QTextTable::format();
extern void C_ZNK10QTextTable6formatEv(void* qthis); // 2
  // proto:  void QTextTable::removeRows(int pos, int num);
extern void C_ZN10QTextTable10removeRowsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTextTable::columns();
extern void C_ZNK10QTextTable7columnsEv(void* qthis); // 4
  // proto:  void QTextTable::appendColumns(int count);
extern void C_ZN10QTextTable13appendColumnsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextTable::insertRows(int pos, int num);
extern void C_ZN10QTextTable10insertRowsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTextTable::mergeCells(int row, int col, int numRows, int numCols);
extern void C_ZN10QTextTable10mergeCellsEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QTextTable::mergeCells(const QTextCursor & cursor);
extern void C_ZN10QTextTable10mergeCellsERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  void QTextTable::resize(int rows, int cols);
extern void C_ZN10QTextTable6resizeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTextCursor QTextTable::rowStart(const QTextCursor & c);
extern void C_ZNK10QTextTable8rowStartERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QTextTable::metaObject();
extern void C_ZNK10QTextTable10metaObjectEv(void* qthis); // 4
  // proto:  void QTextTable::QTextTable(QTextDocument * doc);
extern void* C_ZN10QTextTableC2EP13QTextDocument(void* arg0); // 3
  // proto:  QTextCursor QTextTable::rowEnd(const QTextCursor & c);
extern void C_ZNK10QTextTable6rowEndERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  void QTextTable::insertColumns(int pos, int num);
extern void C_ZN10QTextTable13insertColumnsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTextTable::removeColumns(int pos, int num);
extern void C_ZN10QTextTable13removeColumnsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
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

// QTextTableCell(const class QTextTableCell &)
func NewQTextTableCell(args ...interface{}) *QTextTableCell {
  // QTextTableCell(const class QTextTableCell &)
  // QTextTableCell()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextTableCell{}) // "const QTextTableCell &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTextTableCellC1ERKS_
    // invoke: void QTextTableCell(const class QTextTableCell &)
    var arg0 = args[0].(QTextTableCell).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTextTableCellC2ERKS_(arg0)
    return &QTextTableCell{qclsinst:qthis}
  case 1:
    // invoke: _ZN14QTextTableCellC1Ev
    // invoke: void QTextTableCell()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTextTableCellC2Ev()
    return &QTextTableCell{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextTableCell", "QTextTableCell", args)
  }

  return nil // QTextTableCell{qclsinst:qthis}
}

// rowSpan()
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
    var ret = C.C_ZNK14QTextTableCell7rowSpanEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "rowSpan", args)
  }

}

// firstCursorPosition()
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
    var ret = C.C_ZNK14QTextTableCell19firstCursorPositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "firstCursorPosition", args)
  }

}

// format()
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
    var ret = C.C_ZNK14QTextTableCell6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "format", args)
  }

}

// setFormat(const class QTextCharFormat &)
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
    C.C_ZN14QTextTableCell9setFormatERK15QTextCharFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCell", "setFormat", args)
  }

}

// isValid()
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
    var ret = C.C_ZNK14QTextTableCell7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "isValid", args)
  }

}

// begin()
func (this *QTextTableCell) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell5beginEv
    // invoke: QTextFrame::iterator begin()
    C.C_ZNK14QTextTableCell5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "begin", args)
  }

}

// lastCursorPosition()
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
    var ret = C.C_ZNK14QTextTableCell18lastCursorPositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "lastCursorPosition", args)
  }

}

// tableCellFormatIndex()
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
    var ret = C.C_ZNK14QTextTableCell20tableCellFormatIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "tableCellFormatIndex", args)
  }

}

// column()
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
    var ret = C.C_ZNK14QTextTableCell6columnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "column", args)
  }

}

// columnSpan()
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
    var ret = C.C_ZNK14QTextTableCell10columnSpanEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "columnSpan", args)
  }

}

// end()
func (this *QTextTableCell) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell3endEv
    // invoke: QTextFrame::iterator end()
    C.C_ZNK14QTextTableCell3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "end", args)
  }

}

// lastPosition()
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
    var ret = C.C_ZNK14QTextTableCell12lastPositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "lastPosition", args)
  }

}

// firstPosition()
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
    var ret = C.C_ZNK14QTextTableCell13firstPositionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "firstPosition", args)
  }

}

// row()
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
    var ret = C.C_ZNK14QTextTableCell3rowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTableCell", "row", args)
  }

}

// appendRows(int)
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
    C.C_ZN10QTextTable10appendRowsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "appendRows", args)
  }

}

// ~QTextTable()
func (this *QTextTable) FreeQTextTable(args ...interface{}) () {
  // ~QTextTable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTableD0Ev
    // invoke: void ~QTextTable()
    C.C_ZN10QTextTableD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "~QTextTable", args)
  }

}

// splitCell(int, int, int, int)
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
    C.C_ZN10QTextTable9splitCellEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextTable", "splitCell", args)
  }

}

// rows()
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
    var ret = C.C_ZNK10QTextTable4rowsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTable", "rows", args)
  }

}

// cellAt(const class QTextCursor &)
func (this *QTextTable) cellAt(args ...interface{}) () {
  // cellAt(const class QTextCursor &)
  // cellAt(int, int)
  // cellAt(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6cellAtERK11QTextCursor
    // invoke: QTextTableCell cellAt(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTextTable6cellAtERK11QTextCursor(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK10QTextTable6cellAtEii
    // invoke: QTextTableCell cellAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK10QTextTable6cellAtEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZNK10QTextTable6cellAtEi
    // invoke: QTextTableCell cellAt(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK10QTextTable6cellAtEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTable", "cellAt", args)
  }

}

// setFormat(const class QTextTableFormat &)
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
    C.C_ZN10QTextTable9setFormatERK16QTextTableFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "setFormat", args)
  }

}

// format()
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
    var ret = C.C_ZNK10QTextTable6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTable", "format", args)
  }

}

// removeRows(int, int)
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
    C.C_ZN10QTextTable10removeRowsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "removeRows", args)
  }

}

// columns()
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
    var ret = C.C_ZNK10QTextTable7columnsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTable", "columns", args)
  }

}

// appendColumns(int)
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
    C.C_ZN10QTextTable13appendColumnsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "appendColumns", args)
  }

}

// insertRows(int, int)
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
    C.C_ZN10QTextTable10insertRowsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "insertRows", args)
  }

}

// mergeCells(int, int, int, int)
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
    C.C_ZN10QTextTable10mergeCellsEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN10QTextTable10mergeCellsERK11QTextCursor
    // invoke: void mergeCells(const class QTextCursor &)
    var arg0 = args[0].(QTextCursor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextTable10mergeCellsERK11QTextCursor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "mergeCells", args)
  }

}

// resize(int, int)
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
    C.C_ZN10QTextTable6resizeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "resize", args)
  }

}

// rowStart(const class QTextCursor &)
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
    var ret = C.C_ZNK10QTextTable8rowStartERK11QTextCursor(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTable", "rowStart", args)
  }

}

// metaObject()
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
    C.C_ZNK10QTextTable10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "metaObject", args)
  }

}

// QTextTable(class QTextDocument *)
func NewQTextTable(args ...interface{}) *QTextTable {
  // QTextTable(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTableC1EP13QTextDocument
    // invoke: void QTextTable(class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTextTableC2EP13QTextDocument(arg0)
    return &QTextTable{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextTable", "QTextTable", args)
  }

  return nil // QTextTable{qclsinst:qthis}
}

// rowEnd(const class QTextCursor &)
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
    var ret = C.C_ZNK10QTextTable6rowEndERK11QTextCursor(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextTable", "rowEnd", args)
  }

}

// insertColumns(int, int)
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
    C.C_ZN10QTextTable13insertColumnsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "insertColumns", args)
  }

}

// removeColumns(int, int)
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
    C.C_ZN10QTextTable13removeColumnsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "removeColumns", args)
  }

}

// <= body block end

