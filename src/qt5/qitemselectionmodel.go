package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qitemselectionmodel.h
// dst-file: /src/core/qitemselectionmodel.go
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
// class sizeof(QItemSelection)=1
type QItemSelection struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QItemSelectionRange)=16
type QItemSelectionRange struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QItemSelectionModel)=1
type QItemSelectionModel struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _currentRowChanged QItemSelectionModel_currentRowChanged_signal;
//  _currentColumnChanged QItemSelectionModel_currentColumnChanged_signal;
//  _modelChanged QItemSelectionModel_modelChanged_signal;
//  _selectionChanged QItemSelectionModel_selectionChanged_signal;
//  _currentChanged QItemSelectionModel_currentChanged_signal;
}


func (this *QItemSelection) split_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemSelection", "split", args)
  }

}


func (this *QItemSelection) indexes(args ...interface{}) () {
  // indexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QItemSelection7indexesEv
  default:
    qtrt.ErrorResolve("QItemSelection", "indexes", args)
  }

}


func NewQItemSelection(args ...interface{}) QItemSelection {
  return QItemSelection{}
}


func (this *QItemSelection) contains(args ...interface{}) () {
  // contains(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QItemSelection8containsERK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelection", "contains", args)
  }

}


func (this *QItemSelection) select_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemSelection", "select", args)
  }

}


func (this *QItemSelectionRange) left(args ...interface{}) () {
  // left()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange4leftEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "left", args)
  }

}


func (this *QItemSelectionRange) contains(args ...interface{}) () {
  // contains(const class QModelIndex &)
  // contains(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange8containsERK11QModelIndex
  case 1:
    // invoke: _ZNK19QItemSelectionRange8containsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "contains", args)
  }

}


func (this *QItemSelectionRange) intersected(args ...interface{}) () {
  // intersected(const class QItemSelectionRange &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange11intersectedERKS_
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "intersected", args)
  }

}


func (this *QItemSelectionRange) bottom(args ...interface{}) () {
  // bottom()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange6bottomEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "bottom", args)
  }

}


func (this *QItemSelectionRange) indexes(args ...interface{}) () {
  // indexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7indexesEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "indexes", args)
  }

}


func (this *QItemSelectionRange) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7isValidEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "isValid", args)
  }

}


func (this *QItemSelectionRange) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange5modelEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "model", args)
  }

}


func (this *QItemSelectionRange) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange6heightEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "height", args)
  }

}


func (this *QItemSelectionRange) right(args ...interface{}) () {
  // right()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange5rightEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "right", args)
  }

}


func (this *QItemSelectionRange) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange6parentEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "parent", args)
  }

}


func NewQItemSelectionRange(args ...interface{}) QItemSelectionRange {
  return QItemSelectionRange{}
}


func (this *QItemSelectionRange) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange5widthEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "width", args)
  }

}


func (this *QItemSelectionRange) topLeft(args ...interface{}) () {
  // topLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7topLeftEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "topLeft", args)
  }

}


func (this *QItemSelectionRange) intersects(args ...interface{}) () {
  // intersects(const class QItemSelectionRange &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionRange{}) // "const QItemSelectionRange &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange10intersectsERKS_
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "intersects", args)
  }

}


func (this *QItemSelectionRange) bottomRight(args ...interface{}) () {
  // bottomRight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange11bottomRightEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "bottomRight", args)
  }

}


func (this *QItemSelectionRange) top(args ...interface{}) () {
  // top()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange3topEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "top", args)
  }

}


func (this *QItemSelectionRange) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionRange7isEmptyEv
  default:
    qtrt.ErrorResolve("QItemSelectionRange", "isEmpty", args)
  }

}


func (this *QItemSelectionModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "metaObject", args)
  }

}


func (this *QItemSelectionModel) selection(args ...interface{}) () {
  // selection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel9selectionEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selection", args)
  }

}


func (this *QItemSelectionModel) isSelected(args ...interface{}) () {
  // isSelected(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel10isSelectedERK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isSelected", args)
  }

}


func (this *QItemSelectionModel) clearSelection(args ...interface{}) () {
  // clearSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel14clearSelectionEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clearSelection", args)
  }

}


func (this *QItemSelectionModel) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel12currentIndexEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "currentIndex", args)
  }

}


func (this *QItemSelectionModel) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel8setModelEP18QAbstractItemModel
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "setModel", args)
  }

}


func NewQItemSelectionModel(args ...interface{}) QItemSelectionModel {
  return QItemSelectionModel{}
}


func (this *QItemSelectionModel) rowIntersectsSelection(args ...interface{}) () {
  // rowIntersectsSelection(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel22rowIntersectsSelectionEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "rowIntersectsSelection", args)
  }

}


func (this *QItemSelectionModel) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel5resetEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "reset", args)
  }

}


func (this *QItemSelectionModel) clearCurrentIndex(args ...interface{}) () {
  // clearCurrentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel17clearCurrentIndexEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clearCurrentIndex", args)
  }

}


func (this *QItemSelectionModel) selectedIndexes(args ...interface{}) () {
  // selectedIndexes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel15selectedIndexesEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedIndexes", args)
  }

}


func (this *QItemSelectionModel) selectedColumns(args ...interface{}) () {
  // selectedColumns(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel15selectedColumnsEi
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedColumns", args)
  }

}


func (this *QItemSelectionModel) isColumnSelected(args ...interface{}) () {
  // isColumnSelected(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel16isColumnSelectedEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isColumnSelected", args)
  }

}


func (this *QItemSelectionModel) columnIntersectsSelection(args ...interface{}) () {
  // columnIntersectsSelection(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel25columnIntersectsSelectionEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "columnIntersectsSelection", args)
  }

}


func (this *QItemSelectionModel) FreeQItemSelectionModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "~QItemSelectionModel", args)
  }

}


func (this *QItemSelectionModel) isRowSelected(args ...interface{}) () {
  // isRowSelected(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel13isRowSelectedEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "isRowSelected", args)
  }

}


func (this *QItemSelectionModel) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QItemSelectionModel5clearEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "clear", args)
  }

}


func (this *QItemSelectionModel) hasSelection(args ...interface{}) () {
  // hasSelection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel12hasSelectionEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "hasSelection", args)
  }

}


func (this *QItemSelectionModel) model(args ...interface{}) () {
  // model()
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel5modelEv
  case 1:
    // invoke: _ZN19QItemSelectionModel5modelEv
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "model", args)
  }

}


func (this *QItemSelectionModel) selectedRows(args ...interface{}) () {
  // selectedRows(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QItemSelectionModel12selectedRowsEi
  default:
    qtrt.ErrorResolve("QItemSelectionModel", "selectedRows", args)
  }

}

// <= body block end

