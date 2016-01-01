package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qtexttable.h
// dst-file: /src/gui/qtexttable.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
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


func NewQTextTableCell(args ...interface{}) QTextTableCell {
  return QTextTableCell{}
}


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
  default:
    qtrt.ErrorResolve("QTextTableCell", "setFormat", args)
  }

}


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
  case 1:
    // invoke: _ZNK10QTextTable6cellAtERK11QTextCursor
  case 2:
    // invoke: _ZNK10QTextTable6cellAtEi
  default:
    qtrt.ErrorResolve("QTextTable", "cellAt", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "removeRows", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "appendRows", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "resize", args)
  }

}


func NewQTextTable(args ...interface{}) QTextTable {
  return QTextTable{}
}


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
  default:
    qtrt.ErrorResolve("QTextTable", "setFormat", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "insertColumns", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "splitCell", args)
  }

}


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
  case 1:
    // invoke: _ZN10QTextTable10mergeCellsERK11QTextCursor
  default:
    qtrt.ErrorResolve("QTextTable", "mergeCells", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "insertRows", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "rowStart", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "rowEnd", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "removeColumns", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QTextTable", "appendColumns", args)
  }

}

// <= body block end

