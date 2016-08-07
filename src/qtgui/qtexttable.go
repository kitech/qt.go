package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
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
extern int32_t C_ZNK14QTextTableCell7rowSpanEv(void* qthis); // 4
  // proto:  QTextCursor QTextTableCell::firstCursorPosition();
extern void* C_ZNK14QTextTableCell19firstCursorPositionEv(void* qthis); // 4
  // proto:  QTextCharFormat QTextTableCell::format();
extern void* C_ZNK14QTextTableCell6formatEv(void* qthis); // 4
  // proto:  void QTextTableCell::setFormat(const QTextCharFormat & format);
extern void C_ZN14QTextTableCell9setFormatERK15QTextCharFormat(void* qthis, void* arg0); // 4
  // proto:  bool QTextTableCell::isValid();
extern bool C_ZNK14QTextTableCell7isValidEv(void* qthis); // 2
  // proto:  QTextFrame::iterator QTextTableCell::begin();
extern void C_ZNK14QTextTableCell5beginEv(void* qthis); // 4
  // proto:  QTextCursor QTextTableCell::lastCursorPosition();
extern void* C_ZNK14QTextTableCell18lastCursorPositionEv(void* qthis); // 4
  // proto:  int QTextTableCell::tableCellFormatIndex();
extern int32_t C_ZNK14QTextTableCell20tableCellFormatIndexEv(void* qthis); // 4
  // proto:  int QTextTableCell::column();
extern int32_t C_ZNK14QTextTableCell6columnEv(void* qthis); // 4
  // proto:  int QTextTableCell::columnSpan();
extern int32_t C_ZNK14QTextTableCell10columnSpanEv(void* qthis); // 4
  // proto:  QTextFrame::iterator QTextTableCell::end();
extern void C_ZNK14QTextTableCell3endEv(void* qthis); // 4
  // proto:  int QTextTableCell::lastPosition();
extern int32_t C_ZNK14QTextTableCell12lastPositionEv(void* qthis); // 4
  // proto:  int QTextTableCell::firstPosition();
extern int32_t C_ZNK14QTextTableCell13firstPositionEv(void* qthis); // 4
  // proto:  int QTextTableCell::row();
extern int32_t C_ZNK14QTextTableCell3rowEv(void* qthis); // 4
  // proto:  void QTextTable::appendRows(int count);
extern void C_ZN10QTextTable10appendRowsEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextTable::~QTextTable();
extern void C_ZN10QTextTableD2Ev(void* qthis); // 4
  // proto:  void QTextTable::splitCell(int row, int col, int numRows, int numCols);
extern void C_ZN10QTextTable9splitCellEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QTextTable::rows();
extern int32_t C_ZNK10QTextTable4rowsEv(void* qthis); // 4
  // proto:  QTextTableCell QTextTable::cellAt(const QTextCursor & c);
extern void* C_ZNK10QTextTable6cellAtERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  QTextTableCell QTextTable::cellAt(int row, int col);
extern void* C_ZNK10QTextTable6cellAtEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTextTableCell QTextTable::cellAt(int position);
extern void* C_ZNK10QTextTable6cellAtEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextTable::setFormat(const QTextTableFormat & format);
extern void C_ZN10QTextTable9setFormatERK16QTextTableFormat(void* qthis, void* arg0); // 4
  // proto:  QTextTableFormat QTextTable::format();
extern void* C_ZNK10QTextTable6formatEv(void* qthis); // 2
  // proto:  void QTextTable::removeRows(int pos, int num);
extern void C_ZN10QTextTable10removeRowsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTextTable::columns();
extern int32_t C_ZNK10QTextTable7columnsEv(void* qthis); // 4
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
extern void* C_ZNK10QTextTable8rowStartERK11QTextCursor(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QTextTable::metaObject();
extern void C_ZNK10QTextTable10metaObjectEv(void* qthis); // 4
  // proto:  void QTextTable::QTextTable(QTextDocument * doc);
extern void* C_ZN10QTextTableC2EP13QTextDocument(void* arg0); // 3
  // proto:  QTextCursor QTextTable::rowEnd(const QTextCursor & c);
extern void* C_ZNK10QTextTable6rowEndERK11QTextCursor(void* qthis, void* arg0); // 4
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
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QTextTableCell)=16
type QTextTableCell struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextTable)=1
type QTextTable struct {
  /*qbase*/ QTextFrame;
  Qclsinst unsafe.Pointer /* *C.void */;
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTextTableCellC1ERKS_
    // invoke: void QTextTableCell(const class QTextTableCell &)
    var arg0 = args[0].(*QTextTableCell).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTextTableCellC2ERKS_(arg0)
    return &QTextTableCell{Qclsinst:qthis}
  case 1:
    // invoke: _ZN14QTextTableCellC1Ev
    // invoke: void QTextTableCell()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QTextTableCellC2Ev()
    return &QTextTableCell{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextTableCell", "QTextTableCell", args)
  }

  return nil // QTextTableCell{Qclsinst:qthis}
}

// rowSpan()
func (this *QTextTableCell) Rowspan(args ...interface{}) (ret interface{}) {
  // rowSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell7rowSpanEv
    // invoke: int rowSpan()
    var ret0 = C.C_ZNK14QTextTableCell7rowSpanEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "rowSpan", args)
  }

  return
}

// firstCursorPosition()
func (this *QTextTableCell) Firstcursorposition(args ...interface{}) (ret interface{}) {
  // firstCursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell19firstCursorPositionEv
    // invoke: QTextCursor firstCursorPosition()
    var ret0 = C.C_ZNK14QTextTableCell19firstCursorPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "firstCursorPosition", args)
  }

  return
}

// format()
func (this *QTextTableCell) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell6formatEv
    // invoke: QTextCharFormat format()
    var ret0 = C.C_ZNK14QTextTableCell6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCharFormat{}) // "QTextCharFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "format", args)
  }

  return
}

// setFormat(const class QTextCharFormat &)
func (this *QTextTableCell) Setformat(args ...interface{}) () {
  // setFormat(const class QTextCharFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCharFormat{}) // "const QTextCharFormat &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTextTableCell9setFormatERK15QTextCharFormat
    // invoke: void setFormat(const class QTextCharFormat &)
    var arg0 = args[0].(*QTextCharFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QTextTableCell9setFormatERK15QTextCharFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTableCell", "setFormat", args)
  }

  return
}

// isValid()
func (this *QTextTableCell) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK14QTextTableCell7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "isValid", args)
  }

  return
}

// begin()
func (this *QTextTableCell) Begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell5beginEv
    // invoke: QTextFrame::iterator begin()
    C.C_ZNK14QTextTableCell5beginEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "begin", args)
  }

  return
}

// lastCursorPosition()
func (this *QTextTableCell) Lastcursorposition(args ...interface{}) (ret interface{}) {
  // lastCursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell18lastCursorPositionEv
    // invoke: QTextCursor lastCursorPosition()
    var ret0 = C.C_ZNK14QTextTableCell18lastCursorPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "lastCursorPosition", args)
  }

  return
}

// tableCellFormatIndex()
func (this *QTextTableCell) Tablecellformatindex(args ...interface{}) (ret interface{}) {
  // tableCellFormatIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell20tableCellFormatIndexEv
    // invoke: int tableCellFormatIndex()
    var ret0 = C.C_ZNK14QTextTableCell20tableCellFormatIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "tableCellFormatIndex", args)
  }

  return
}

// column()
func (this *QTextTableCell) Column(args ...interface{}) (ret interface{}) {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell6columnEv
    // invoke: int column()
    var ret0 = C.C_ZNK14QTextTableCell6columnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "column", args)
  }

  return
}

// columnSpan()
func (this *QTextTableCell) Columnspan(args ...interface{}) (ret interface{}) {
  // columnSpan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell10columnSpanEv
    // invoke: int columnSpan()
    var ret0 = C.C_ZNK14QTextTableCell10columnSpanEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "columnSpan", args)
  }

  return
}

// end()
func (this *QTextTableCell) End(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell3endEv
    // invoke: QTextFrame::iterator end()
    C.C_ZNK14QTextTableCell3endEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextTableCell", "end", args)
  }

  return
}

// lastPosition()
func (this *QTextTableCell) Lastposition(args ...interface{}) (ret interface{}) {
  // lastPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell12lastPositionEv
    // invoke: int lastPosition()
    var ret0 = C.C_ZNK14QTextTableCell12lastPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "lastPosition", args)
  }

  return
}

// firstPosition()
func (this *QTextTableCell) Firstposition(args ...interface{}) (ret interface{}) {
  // firstPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell13firstPositionEv
    // invoke: int firstPosition()
    var ret0 = C.C_ZNK14QTextTableCell13firstPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "firstPosition", args)
  }

  return
}

// row()
func (this *QTextTableCell) Row(args ...interface{}) (ret interface{}) {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTextTableCell3rowEv
    // invoke: int row()
    var ret0 = C.C_ZNK14QTextTableCell3rowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTableCell", "row", args)
  }

  return
}

// appendRows(int)
func (this *QTextTable) Appendrows(args ...interface{}) () {
  // appendRows(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10appendRowsEi
    // invoke: void appendRows(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextTable10appendRowsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "appendRows", args)
  }

  return
}

// ~QTextTable()
func (this *QTextTable) Freeqtexttable(args ...interface{}) () {
  // ~QTextTable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTableD0Ev
    // invoke: void ~QTextTable()
    C.C_ZN10QTextTableD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "~QTextTable", args)
  }

  return
}

// splitCell(int, int, int, int)
func (this *QTextTable) Splitcell(args ...interface{}) () {
  // splitCell(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable9splitCellEiiii
    // invoke: void splitCell(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN10QTextTable9splitCellEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QTextTable", "splitCell", args)
  }

  return
}

// rows()
func (this *QTextTable) Rows(args ...interface{}) (ret interface{}) {
  // rows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable4rowsEv
    // invoke: int rows()
    var ret0 = C.C_ZNK10QTextTable4rowsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTable", "rows", args)
  }

  return
}

// cellAt(const class QTextCursor &)
func (this *QTextTable) Cellat(args ...interface{}) (ret interface{}) {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6cellAtERK11QTextCursor
    // invoke: QTextTableCell cellAt(const class QTextCursor &)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextTable6cellAtERK11QTextCursor(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTableCell{}) // "QTextTableCell"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QTextTable6cellAtEii
    // invoke: QTextTableCell cellAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QTextTable6cellAtEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTableCell{}) // "QTextTableCell"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QTextTable6cellAtEi
    // invoke: QTextTableCell cellAt(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextTable6cellAtEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTableCell{}) // "QTextTableCell"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTable", "cellAt", args)
  }

  return
}

// setFormat(const class QTextTableFormat &)
func (this *QTextTable) Setformat(args ...interface{}) () {
  // setFormat(const class QTextTableFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextTableFormat{}) // "const QTextTableFormat &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable9setFormatERK16QTextTableFormat
    // invoke: void setFormat(const class QTextTableFormat &)
    var arg0 = args[0].(*QTextTableFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextTable9setFormatERK16QTextTableFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "setFormat", args)
  }

  return
}

// format()
func (this *QTextTable) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6formatEv
    // invoke: QTextTableFormat format()
    var ret0 = C.C_ZNK10QTextTable6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextTableFormat{}) // "QTextTableFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTable", "format", args)
  }

  return
}

// removeRows(int, int)
func (this *QTextTable) Removerows(args ...interface{}) () {
  // removeRows(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10removeRowsEii
    // invoke: void removeRows(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTextTable10removeRowsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "removeRows", args)
  }

  return
}

// columns()
func (this *QTextTable) Columns(args ...interface{}) (ret interface{}) {
  // columns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable7columnsEv
    // invoke: int columns()
    var ret0 = C.C_ZNK10QTextTable7columnsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTable", "columns", args)
  }

  return
}

// appendColumns(int)
func (this *QTextTable) Appendcolumns(args ...interface{}) () {
  // appendColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable13appendColumnsEi
    // invoke: void appendColumns(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextTable13appendColumnsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "appendColumns", args)
  }

  return
}

// insertRows(int, int)
func (this *QTextTable) Insertrows(args ...interface{}) () {
  // insertRows(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10insertRowsEii
    // invoke: void insertRows(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTextTable10insertRowsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "insertRows", args)
  }

  return
}

// mergeCells(int, int, int, int)
func (this *QTextTable) Mergecells(args ...interface{}) () {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable10mergeCellsEiiii
    // invoke: void mergeCells(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN10QTextTable10mergeCellsEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN10QTextTable10mergeCellsERK11QTextCursor
    // invoke: void mergeCells(const class QTextCursor &)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextTable10mergeCellsERK11QTextCursor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextTable", "mergeCells", args)
  }

  return
}

// resize(int, int)
func (this *QTextTable) Resize(args ...interface{}) () {
  // resize(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable6resizeEii
    // invoke: void resize(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTextTable6resizeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "resize", args)
  }

  return
}

// rowStart(const class QTextCursor &)
func (this *QTextTable) Rowstart(args ...interface{}) (ret interface{}) {
  // rowStart(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable8rowStartERK11QTextCursor
    // invoke: QTextCursor rowStart(const class QTextCursor &)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextTable8rowStartERK11QTextCursor(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTable", "rowStart", args)
  }

  return
}

// metaObject()
func (this *QTextTable) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QTextTable10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextTable", "metaObject", args)
  }

  return
}

// QTextTable(class QTextDocument *)
func NewQTextTable(args ...interface{}) *QTextTable {
  // QTextTable(class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "QTextDocument *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTableC1EP13QTextDocument
    // invoke: void QTextTable(class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QTextTableC2EP13QTextDocument(arg0)
    return &QTextTable{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextTable", "QTextTable", args)
  }

  return nil // QTextTable{Qclsinst:qthis}
}

// rowEnd(const class QTextCursor &)
func (this *QTextTable) Rowend(args ...interface{}) (ret interface{}) {
  // rowEnd(const class QTextCursor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCursor{}) // "const QTextCursor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextTable6rowEndERK11QTextCursor
    // invoke: QTextCursor rowEnd(const class QTextCursor &)
    var arg0 = args[0].(*QTextCursor).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextTable6rowEndERK11QTextCursor(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCursor{}) // "QTextCursor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextTable", "rowEnd", args)
  }

  return
}

// insertColumns(int, int)
func (this *QTextTable) Insertcolumns(args ...interface{}) () {
  // insertColumns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable13insertColumnsEii
    // invoke: void insertColumns(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTextTable13insertColumnsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "insertColumns", args)
  }

  return
}

// removeColumns(int, int)
func (this *QTextTable) Removecolumns(args ...interface{}) () {
  // removeColumns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextTable13removeColumnsEii
    // invoke: void removeColumns(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QTextTable13removeColumnsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextTable", "removeColumns", args)
  }

  return
}

// <= body block end

