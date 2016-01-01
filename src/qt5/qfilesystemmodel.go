package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qfilesystemmodel.h
// dst-file: /src/widgets/qfilesystemmodel.go
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
// class sizeof(QFileSystemModel)=1
type QFileSystemModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
//  _rootPathChanged QFileSystemModel_rootPathChanged_signal;
//  _directoryLoaded QFileSystemModel_directoryLoaded_signal;
//  _fileRenamed QFileSystemModel_fileRenamed_signal;
}


func (this *QFileSystemModel) fileName(args ...interface{}) () {
  // fileName(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8fileNameERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileName", args)
  }

}


func (this *QFileSystemModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11hasChildrenERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "hasChildren", args)
  }

}


func (this *QFileSystemModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel9mimeTypesEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "mimeTypes", args)
  }

}


func (this *QFileSystemModel) FreeQFileSystemModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileSystemModel", "~QFileSystemModel", args)
  }

}


func (this *QFileSystemModel) index(args ...interface{}) () {
  // index(const class QString &, int)
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel5indexERK7QStringi
  case 1:
    // invoke: _ZNK16QFileSystemModel5indexEiiRK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "index", args)
  }

}


func (this *QFileSystemModel) setNameFilterDisables(args ...interface{}) () {
  // setNameFilterDisables(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel21setNameFilterDisablesEb
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setNameFilterDisables", args)
  }

}


func (this *QFileSystemModel) fileIcon(args ...interface{}) () {
  // fileIcon(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8fileIconERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileIcon", args)
  }

}


func (this *QFileSystemModel) resolveSymlinks(args ...interface{}) () {
  // resolveSymlinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel15resolveSymlinksEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "resolveSymlinks", args)
  }

}


func (this *QFileSystemModel) filePath(args ...interface{}) () {
  // filePath(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8filePathERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "filePath", args)
  }

}


func (this *QFileSystemModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel6parentERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "parent", args)
  }

}


func (this *QFileSystemModel) nameFilterDisables(args ...interface{}) () {
  // nameFilterDisables()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel18nameFilterDisablesEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "nameFilterDisables", args)
  }

}


func (this *QFileSystemModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "metaObject", args)
  }

}


func (this *QFileSystemModel) fetchMore(args ...interface{}) () {
  // fetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel9fetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fetchMore", args)
  }

}


func NewQFileSystemModel(args ...interface{}) QFileSystemModel {
  return QFileSystemModel{}
}


func (this *QFileSystemModel) size(args ...interface{}) () {
  // size(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel4sizeERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "size", args)
  }

}


func (this *QFileSystemModel) iconProvider(args ...interface{}) () {
  // iconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel12iconProviderEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "iconProvider", args)
  }

}


func (this *QFileSystemModel) setNameFilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel14setNameFiltersERK11QStringList
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setNameFilters", args)
  }

}


func (this *QFileSystemModel) data(args ...interface{}) () {
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
    // invoke: _ZNK16QFileSystemModel4dataERK11QModelIndexi
  default:
    qtrt.ErrorResolve("QFileSystemModel", "data", args)
  }

}


func (this *QFileSystemModel) rootDirectory(args ...interface{}) () {
  // rootDirectory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel13rootDirectoryEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rootDirectory", args)
  }

}


func (this *QFileSystemModel) mkdir(args ...interface{}) () {
  // mkdir(const class QModelIndex &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString
  default:
    qtrt.ErrorResolve("QFileSystemModel", "mkdir", args)
  }

}


func (this *QFileSystemModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8rowCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rowCount", args)
  }

}


func (this *QFileSystemModel) setData(args ...interface{}) () {
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
    // invoke: _ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setData", args)
  }

}


func (this *QFileSystemModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11columnCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "columnCount", args)
  }

}


func (this *QFileSystemModel) canFetchMore(args ...interface{}) () {
  // canFetchMore(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "canFetchMore", args)
  }

}


func (this *QFileSystemModel) remove(args ...interface{}) () {
  // remove(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel6removeERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "remove", args)
  }

}


func (this *QFileSystemModel) setRootPath(args ...interface{}) () {
  // setRootPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel11setRootPathERK7QString
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setRootPath", args)
  }

}


func (this *QFileSystemModel) setResolveSymlinks(args ...interface{}) () {
  // setResolveSymlinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel18setResolveSymlinksEb
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setResolveSymlinks", args)
  }

}


func (this *QFileSystemModel) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel11setReadOnlyEb
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setReadOnly", args)
  }

}


func (this *QFileSystemModel) nameFilters(args ...interface{}) () {
  // nameFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel11nameFiltersEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "nameFilters", args)
  }

}


func (this *QFileSystemModel) fileInfo(args ...interface{}) () {
  // fileInfo(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8fileInfoERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileInfo", args)
  }

}


func (this *QFileSystemModel) myComputer(args ...interface{}) () {
  // myComputer(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel10myComputerEi
  default:
    qtrt.ErrorResolve("QFileSystemModel", "myComputer", args)
  }

}


func (this *QFileSystemModel) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "isReadOnly", args)
  }

}


func (this *QFileSystemModel) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileSystemModel", "type", args)
  }

}


func (this *QFileSystemModel) rootPath(args ...interface{}) () {
  // rootPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel8rootPathEv
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rootPath", args)
  }

}


func (this *QFileSystemModel) lastModified(args ...interface{}) () {
  // lastModified(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel12lastModifiedERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "lastModified", args)
  }

}


func (this *QFileSystemModel) isDir(args ...interface{}) () {
  // isDir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QFileSystemModel5isDirERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "isDir", args)
  }

}


func (this *QFileSystemModel) rmdir(args ...interface{}) () {
  // rmdir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel5rmdirERK11QModelIndex
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rmdir", args)
  }

}


func (this *QFileSystemModel) setIconProvider(args ...interface{}) () {
  // setIconProvider(class QFileIconProvider *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setIconProvider", args)
  }

}

// <= body block end

