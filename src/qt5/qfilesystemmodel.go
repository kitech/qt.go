package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qfilesystemmodel.h
// dst-file: /src/widgets/qfilesystemmodel.go
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
  // proto:  QString QFileSystemModel::fileName(const QModelIndex & index);
extern void demth_ZNK16QFileSystemModel8fileNameERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::hasChildren(const QModelIndex & parent);
extern void _ZNK16QFileSystemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QStringList QFileSystemModel::mimeTypes();
extern void _ZNK16QFileSystemModel9mimeTypesEv(void* qthis);
  // proto:  void QFileSystemModel::~QFileSystemModel();
extern void _ZN16QFileSystemModelD0Ev(void* qthis);
  // proto:  QModelIndex QFileSystemModel::index(const QString & path, int column);
extern void _ZNK16QFileSystemModel5indexERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  void QFileSystemModel::setNameFilterDisables(bool enable);
extern void _ZN16QFileSystemModel21setNameFilterDisablesEb(void* qthis, bool arg0);
  // proto:  QIcon QFileSystemModel::fileIcon(const QModelIndex & index);
extern void demth_ZNK16QFileSystemModel8fileIconERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::resolveSymlinks();
extern void _ZNK16QFileSystemModel15resolveSymlinksEv(void* qthis);
  // proto:  QString QFileSystemModel::filePath(const QModelIndex & index);
extern void _ZNK16QFileSystemModel8filePathERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QFileSystemModel::parent(const QModelIndex & child);
extern void _ZNK16QFileSystemModel6parentERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::nameFilterDisables();
extern void _ZNK16QFileSystemModel18nameFilterDisablesEv(void* qthis);
  // proto:  const QMetaObject * QFileSystemModel::metaObject();
extern void _ZNK16QFileSystemModel10metaObjectEv(void* qthis);
  // proto:  void QFileSystemModel::fetchMore(const QModelIndex & parent);
extern void _ZN16QFileSystemModel9fetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QFileSystemModel::QFileSystemModel(const QFileSystemModel & );
extern void* dector_ZN16QFileSystemModelC1ERKS_(void* arg0);
extern void _ZN16QFileSystemModelC1ERKS_(void* qthis, void* arg0);
  // proto:  qint64 QFileSystemModel::size(const QModelIndex & index);
extern void _ZNK16QFileSystemModel4sizeERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QFileIconProvider * QFileSystemModel::iconProvider();
extern void _ZNK16QFileSystemModel12iconProviderEv(void* qthis);
  // proto:  void QFileSystemModel::setNameFilters(const QStringList & filters);
extern void _ZN16QFileSystemModel14setNameFiltersERK11QStringList(void* qthis, void* arg0);
  // proto:  QVariant QFileSystemModel::data(const QModelIndex & index, int role);
extern void _ZNK16QFileSystemModel4dataERK11QModelIndexi(void* qthis, void* arg0, int arg1);
  // proto:  QDir QFileSystemModel::rootDirectory();
extern void _ZNK16QFileSystemModel13rootDirectoryEv(void* qthis);
  // proto:  QModelIndex QFileSystemModel::mkdir(const QModelIndex & parent, const QString & name);
extern void _ZN16QFileSystemModel5mkdirERK11QModelIndexRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  int QFileSystemModel::rowCount(const QModelIndex & parent);
extern void _ZNK16QFileSystemModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN16QFileSystemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  int QFileSystemModel::columnCount(const QModelIndex & parent);
extern void _ZNK16QFileSystemModel11columnCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QFileSystemModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK16QFileSystemModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QFileSystemModel::canFetchMore(const QModelIndex & parent);
extern void _ZNK16QFileSystemModel12canFetchMoreERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::remove(const QModelIndex & index);
extern void _ZN16QFileSystemModel6removeERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QFileSystemModel::setRootPath(const QString & path);
extern void _ZN16QFileSystemModel11setRootPathERK7QString(void* qthis, void* arg0);
  // proto:  void QFileSystemModel::setResolveSymlinks(bool enable);
extern void _ZN16QFileSystemModel18setResolveSymlinksEb(void* qthis, bool arg0);
  // proto:  void QFileSystemModel::setReadOnly(bool enable);
extern void _ZN16QFileSystemModel11setReadOnlyEb(void* qthis, bool arg0);
  // proto:  QStringList QFileSystemModel::nameFilters();
extern void _ZNK16QFileSystemModel11nameFiltersEv(void* qthis);
  // proto:  void QFileSystemModel::QFileSystemModel(QObject * parent);
extern void* dector_ZN16QFileSystemModelC1EP7QObject(void* arg0);
extern void _ZN16QFileSystemModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  QFileInfo QFileSystemModel::fileInfo(const QModelIndex & index);
extern void demth_ZNK16QFileSystemModel8fileInfoERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QVariant QFileSystemModel::myComputer(int role);
extern void _ZNK16QFileSystemModel10myComputerEi(void* qthis, int arg0);
  // proto:  bool QFileSystemModel::isReadOnly();
extern void _ZNK16QFileSystemModel10isReadOnlyEv(void* qthis);
  // proto:  QString QFileSystemModel::type(const QModelIndex & index);
extern void _ZNK16QFileSystemModel4typeERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QString QFileSystemModel::rootPath();
extern void _ZNK16QFileSystemModel8rootPathEv(void* qthis);
  // proto:  QDateTime QFileSystemModel::lastModified(const QModelIndex & index);
extern void _ZNK16QFileSystemModel12lastModifiedERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::isDir(const QModelIndex & index);
extern void _ZNK16QFileSystemModel5isDirERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QFileSystemModel::rmdir(const QModelIndex & index);
extern void _ZN16QFileSystemModel5rmdirERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QFileSystemModel::setIconProvider(QFileIconProvider * provider);
extern void _ZN16QFileSystemModel15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0);
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

// class sizeof(QFileSystemModel)=1
type QFileSystemModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
//  _rootPathChanged QFileSystemModel_rootPathChanged_signal;
//  _directoryLoaded QFileSystemModel_directoryLoaded_signal;
//  _fileRenamed QFileSystemModel_fileRenamed_signal;
}

  // proto:  QString QFileSystemModel::fileName(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileName", args)
  }

}

  // proto:  bool QFileSystemModel::hasChildren(const QModelIndex & parent);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "hasChildren", args)
  }

}

  // proto:  QStringList QFileSystemModel::mimeTypes();
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

  // proto:  void QFileSystemModel::~QFileSystemModel();
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

  // proto:  QModelIndex QFileSystemModel::index(const QString & path, int column);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZNK16QFileSystemModel5indexEiiRK11QModelIndex
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "index", args)
  }

}

  // proto:  void QFileSystemModel::setNameFilterDisables(bool enable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setNameFilterDisables", args)
  }

}

  // proto:  QIcon QFileSystemModel::fileIcon(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileIcon", args)
  }

}

  // proto:  bool QFileSystemModel::resolveSymlinks();
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

  // proto:  QString QFileSystemModel::filePath(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "filePath", args)
  }

}

  // proto:  QModelIndex QFileSystemModel::parent(const QModelIndex & child);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "parent", args)
  }

}

  // proto:  bool QFileSystemModel::nameFilterDisables();
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

  // proto:  const QMetaObject * QFileSystemModel::metaObject();
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

  // proto:  void QFileSystemModel::fetchMore(const QModelIndex & parent);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fetchMore", args)
  }

}

  // proto:  void QFileSystemModel::QFileSystemModel(const QFileSystemModel & );
func NewQFileSystemModel(args ...interface{}) QFileSystemModel {
  return QFileSystemModel{}
}

  // proto:  qint64 QFileSystemModel::size(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "size", args)
  }

}

  // proto:  QFileIconProvider * QFileSystemModel::iconProvider();
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

  // proto:  void QFileSystemModel::setNameFilters(const QStringList & filters);
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
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setNameFilters", args)
  }

}

  // proto:  QVariant QFileSystemModel::data(const QModelIndex & index, int role);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "data", args)
  }

}

  // proto:  QDir QFileSystemModel::rootDirectory();
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

  // proto:  QModelIndex QFileSystemModel::mkdir(const QModelIndex & parent, const QString & name);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "mkdir", args)
  }

}

  // proto:  int QFileSystemModel::rowCount(const QModelIndex & parent);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rowCount", args)
  }

}

  // proto:  bool QFileSystemModel::setData(const QModelIndex & index, const QVariant & value, int role);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setData", args)
  }

}

  // proto:  int QFileSystemModel::columnCount(const QModelIndex & parent);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "columnCount", args)
  }

}

  // proto:  bool QFileSystemModel::canFetchMore(const QModelIndex & parent);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "canFetchMore", args)
  }

}

  // proto:  bool QFileSystemModel::remove(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "remove", args)
  }

}

  // proto:  QModelIndex QFileSystemModel::setRootPath(const QString & path);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setRootPath", args)
  }

}

  // proto:  void QFileSystemModel::setResolveSymlinks(bool enable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setResolveSymlinks", args)
  }

}

  // proto:  void QFileSystemModel::setReadOnly(bool enable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setReadOnly", args)
  }

}

  // proto:  QStringList QFileSystemModel::nameFilters();
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

  // proto:  QFileInfo QFileSystemModel::fileInfo(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "fileInfo", args)
  }

}

  // proto:  QVariant QFileSystemModel::myComputer(int role);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "myComputer", args)
  }

}

  // proto:  bool QFileSystemModel::isReadOnly();
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

  // proto:  QString QFileSystemModel::type(const QModelIndex & index);
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

  // proto:  QString QFileSystemModel::rootPath();
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

  // proto:  QDateTime QFileSystemModel::lastModified(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "lastModified", args)
  }

}

  // proto:  bool QFileSystemModel::isDir(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "isDir", args)
  }

}

  // proto:  bool QFileSystemModel::rmdir(const QModelIndex & index);
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
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "rmdir", args)
  }

}

  // proto:  void QFileSystemModel::setIconProvider(QFileIconProvider * provider);
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
    var arg0 = args[0].(QFileIconProvider).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QFileSystemModel", "setIconProvider", args)
  }

}

// <= body block end

