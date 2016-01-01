package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qabstractitemmodel.h
// dst-file: /src/core/qabstractitemmodel.go
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
// class sizeof(QModelIndex)=24
type QModelIndex struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPersistentModelIndex)=8
type QPersistentModelIndex struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractTableModel)=1
type QAbstractTableModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAbstractItemModel)=1
type QAbstractItemModel struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _modelReset QAbstractItemModel_modelReset_signal;
//  _columnsAboutToBeMoved QAbstractItemModel_columnsAboutToBeMoved_signal;
//  _columnsInserted QAbstractItemModel_columnsInserted_signal;
//  _rowsAboutToBeRemoved QAbstractItemModel_rowsAboutToBeRemoved_signal;
//  _columnsAboutToBeRemoved QAbstractItemModel_columnsAboutToBeRemoved_signal;
//  _rowsAboutToBeMoved QAbstractItemModel_rowsAboutToBeMoved_signal;
//  _layoutChanged QAbstractItemModel_layoutChanged_signal;
//  _columnsRemoved QAbstractItemModel_columnsRemoved_signal;
//  _rowsInserted QAbstractItemModel_rowsInserted_signal;
//  _columnsAboutToBeInserted QAbstractItemModel_columnsAboutToBeInserted_signal;
//  _layoutAboutToBeChanged QAbstractItemModel_layoutAboutToBeChanged_signal;
//  _rowsRemoved QAbstractItemModel_rowsRemoved_signal;
//  _rowsMoved QAbstractItemModel_rowsMoved_signal;
//  _headerDataChanged QAbstractItemModel_headerDataChanged_signal;
//  _columnsMoved QAbstractItemModel_columnsMoved_signal;
//  _rowsAboutToBeInserted QAbstractItemModel_rowsAboutToBeInserted_signal;
//  _modelAboutToBeReset QAbstractItemModel_modelAboutToBeReset_signal;
//  _dataChanged QAbstractItemModel_dataChanged_signal;
}

// class sizeof(QAbstractListModel)=1
type QAbstractListModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQModelIndex(args ...interface{}) QModelIndex {
  return QModelIndex{}
}


func (this *QModelIndex) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex6columnEv
  default:
    qtrt.ErrorResolve("QModelIndex", "column", args)
  }

}


func (this *QModelIndex) internalId(args ...interface{}) () {
  // internalId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex10internalIdEv
  default:
    qtrt.ErrorResolve("QModelIndex", "internalId", args)
  }

}


func (this *QModelIndex) child(args ...interface{}) () {
  // child(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex5childEii
  default:
    qtrt.ErrorResolve("QModelIndex", "child", args)
  }

}


func (this *QModelIndex) internalPointer(args ...interface{}) () {
  // internalPointer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex15internalPointerEv
  default:
    qtrt.ErrorResolve("QModelIndex", "internalPointer", args)
  }

}


func (this *QModelIndex) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex7isValidEv
  default:
    qtrt.ErrorResolve("QModelIndex", "isValid", args)
  }

}


func (this *QModelIndex) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex6parentEv
  default:
    qtrt.ErrorResolve("QModelIndex", "parent", args)
  }

}


func (this *QModelIndex) sibling(args ...interface{}) () {
  // sibling(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex7siblingEii
  default:
    qtrt.ErrorResolve("QModelIndex", "sibling", args)
  }

}


func (this *QModelIndex) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex5modelEv
  default:
    qtrt.ErrorResolve("QModelIndex", "model", args)
  }

}


func (this *QModelIndex) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex4dataEi
  default:
    qtrt.ErrorResolve("QModelIndex", "data", args)
  }

}


func (this *QModelIndex) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QModelIndex3rowEv
  default:
    qtrt.ErrorResolve("QModelIndex", "row", args)
  }

}


func (this *QPersistentModelIndex) sibling(args ...interface{}) () {
  // sibling(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex7siblingEii
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "sibling", args)
  }

}


func (this *QPersistentModelIndex) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex4dataEi
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "data", args)
  }

}


func (this *QPersistentModelIndex) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex6parentEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "parent", args)
  }

}


func NewQPersistentModelIndex(args ...interface{}) QPersistentModelIndex {
  return QPersistentModelIndex{}
}


func (this *QPersistentModelIndex) internalPointer(args ...interface{}) () {
  // internalPointer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex15internalPointerEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalPointer", args)
  }

}


func (this *QPersistentModelIndex) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex3rowEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "row", args)
  }

}


func (this *QPersistentModelIndex) internalId(args ...interface{}) () {
  // internalId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex10internalIdEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "internalId", args)
  }

}


func (this *QPersistentModelIndex) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex5modelEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "model", args)
  }

}


func (this *QPersistentModelIndex) FreeQPersistentModelIndex(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "~QPersistentModelIndex", args)
  }

}


func (this *QPersistentModelIndex) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex6columnEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "column", args)
  }

}


func (this *QPersistentModelIndex) swap(args ...interface{}) () {
  // swap(class QPersistentModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPersistentModelIndex{}) // "QPersistentModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QPersistentModelIndex4swapERS_
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "swap", args)
  }

}


func (this *QPersistentModelIndex) child(args ...interface{}) () {
  // child(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex5childEii
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "child", args)
  }

}


func (this *QPersistentModelIndex) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK21QPersistentModelIndex7isValidEv
  default:
    qtrt.ErrorResolve("QPersistentModelIndex", "isValid", args)
  }

}


func NewQAbstractTableModel(args ...interface{}) QAbstractTableModel {
  return QAbstractTableModel{}
}


func (this *QAbstractTableModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTableModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "index", args)
  }

}


func (this *QAbstractTableModel) sibling(args ...interface{}) () {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTableModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "sibling", args)
  }

}


func (this *QAbstractTableModel) FreeQAbstractTableModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "~QAbstractTableModel", args)
  }

}


func (this *QAbstractTableModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractTableModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractTableModel", "metaObject", args)
  }

}


func (this *QAbstractItemModel) removeColumns(args ...interface{}) () {
  // removeColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel13removeColumnsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumns", args)
  }

}


func (this *QAbstractItemModel) FreeQAbstractItemModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "~QAbstractItemModel", args)
  }

}


func (this *QAbstractItemModel) canFetchMore(args ...interface{}) () {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel12canFetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "canFetchMore", args)
  }

}


func (this *QAbstractItemModel) submit(args ...interface{}) () {
  // submit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel6submitEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "submit", args)
  }

}


func (this *QAbstractItemModel) hasIndex(args ...interface{}) () {
  // hasIndex(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel8hasIndexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasIndex", args)
  }

}


func (this *QAbstractItemModel) roleNames(args ...interface{}) () {
  // roleNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel9roleNamesEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "roleNames", args)
  }

}


func (this *QAbstractItemModel) moveColumn(args ...interface{}) () {
  // moveColumn(const class QModelIndex &, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel10moveColumnERK11QModelIndexiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumn", args)
  }

}


func (this *QAbstractItemModel) fetchMore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel9fetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "fetchMore", args)
  }

}


func (this *QAbstractItemModel) insertRows(args ...interface{}) () {
  // insertRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel10insertRowsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRows", args)
  }

}


func (this *QAbstractItemModel) span(args ...interface{}) () {
  // span(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel4spanERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "span", args)
  }

}


func NewQAbstractItemModel(args ...interface{}) QAbstractItemModel {
  return QAbstractItemModel{}
}


func (this *QAbstractItemModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "index", args)
  }

}


func (this *QAbstractItemModel) insertRow(args ...interface{}) () {
  // insertRow(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel9insertRowEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertRow", args)
  }

}


func (this *QAbstractItemModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "metaObject", args)
  }

}


func (this *QAbstractItemModel) removeRow(args ...interface{}) () {
  // removeRow(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel9removeRowEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRow", args)
  }

}


func (this *QAbstractItemModel) setData(args ...interface{}) () {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel7setDataERK11QModelIndexRK8QVarianti
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "setData", args)
  }

}


func (this *QAbstractItemModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel8rowCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "rowCount", args)
  }

}


func (this *QAbstractItemModel) removeRows(args ...interface{}) () {
  // removeRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel10removeRowsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeRows", args)
  }

}


func (this *QAbstractItemModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel11hasChildrenERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "hasChildren", args)
  }

}


func (this *QAbstractItemModel) moveRow(args ...interface{}) () {
  // moveRow(const class QModelIndex &, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel7moveRowERK11QModelIndexiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRow", args)
  }

}


func (this *QAbstractItemModel) revert(args ...interface{}) () {
  // revert()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel6revertEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "revert", args)
  }

}


func (this *QAbstractItemModel) removeColumn(args ...interface{}) () {
  // removeColumn(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel12removeColumnEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "removeColumn", args)
  }

}


func (this *QAbstractItemModel) insertColumns(args ...interface{}) () {
  // insertColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel13insertColumnsEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumns", args)
  }

}


func (this *QAbstractItemModel) insertColumn(args ...interface{}) () {
  // insertColumn(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel12insertColumnEiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "insertColumn", args)
  }

}


func (this *QAbstractItemModel) moveColumns(args ...interface{}) () {
  // moveColumns(const class QModelIndex &, int, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel11moveColumnsERK11QModelIndexiiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveColumns", args)
  }

}


func (this *QAbstractItemModel) itemData(args ...interface{}) () {
  // itemData(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel8itemDataERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "itemData", args)
  }

}


func (this *QAbstractItemModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel9mimeTypesEv
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "mimeTypes", args)
  }

}


func (this *QAbstractItemModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel6parentERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "parent", args)
  }

}


func (this *QAbstractItemModel) buddy(args ...interface{}) () {
  // buddy(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel5buddyERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "buddy", args)
  }

}


func (this *QAbstractItemModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel11columnCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "columnCount", args)
  }

}


func (this *QAbstractItemModel) data(args ...interface{}) () {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel4dataERK11QModelIndexi
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "data", args)
  }

}


func (this *QAbstractItemModel) sibling(args ...interface{}) () {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractItemModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "sibling", args)
  }

}


func (this *QAbstractItemModel) moveRows(args ...interface{}) () {
  // moveRows(const class QModelIndex &, int, int, const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QAbstractItemModel8moveRowsERK11QModelIndexiiS2_i
  default:
    qtrt.ErrorResolve("QAbstractItemModel", "moveRows", args)
  }

}


func (this *QAbstractListModel) sibling(args ...interface{}) () {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractListModel7siblingEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractListModel", "sibling", args)
  }

}


func NewQAbstractListModel(args ...interface{}) QAbstractListModel {
  return QAbstractListModel{}
}


func (this *QAbstractListModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractListModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractListModel", "metaObject", args)
  }

}


func (this *QAbstractListModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QAbstractListModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QAbstractListModel", "index", args)
  }

}


func (this *QAbstractListModel) FreeQAbstractListModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractListModel", "~QAbstractListModel", args)
  }

}

// <= body block end

