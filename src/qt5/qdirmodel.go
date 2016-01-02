package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qdirmodel.h
// dst-file: /src/widgets/qdirmodel.go
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
  // proto:  QFileIconProvider * QDirModel::iconProvider();
extern void _ZNK9QDirModel12iconProviderEv(void* qthis);
  // proto:  QModelIndex QDirModel::parent(const QModelIndex & child);
extern void _ZNK9QDirModel6parentERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QVariant QDirModel::data(const QModelIndex & index, int role);
extern void _ZNK9QDirModel4dataERK11QModelIndexi(void* qthis, void* arg0, int arg1);
  // proto:  QStringList QDirModel::nameFilters();
extern void _ZNK9QDirModel11nameFiltersEv(void* qthis);
  // proto:  bool QDirModel::isReadOnly();
extern void _ZNK9QDirModel10isReadOnlyEv(void* qthis);
  // proto:  int QDirModel::columnCount(const QModelIndex & parent);
extern void _ZNK9QDirModel11columnCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QStringList QDirModel::mimeTypes();
extern void _ZNK9QDirModel9mimeTypesEv(void* qthis);
  // proto:  void QDirModel::~QDirModel();
extern void _ZN9QDirModelD0Ev(void* qthis);
  // proto:  bool QDirModel::remove(const QModelIndex & index);
extern void _ZN9QDirModel6removeERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QString QDirModel::fileName(const QModelIndex & index);
extern void _ZNK9QDirModel8fileNameERK11QModelIndex(void* qthis, void* arg0);
  // proto:  const QMetaObject * QDirModel::metaObject();
extern void _ZNK9QDirModel10metaObjectEv(void* qthis);
  // proto:  bool QDirModel::resolveSymlinks();
extern void _ZNK9QDirModel15resolveSymlinksEv(void* qthis);
  // proto:  void QDirModel::refresh(const QModelIndex & parent);
extern void _ZN9QDirModel7refreshERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QDirModel::setNameFilters(const QStringList & filters);
extern void _ZN9QDirModel14setNameFiltersERK11QStringList(void* qthis, void* arg0);
  // proto:  void QDirModel::setIconProvider(QFileIconProvider * provider);
extern void _ZN9QDirModel15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0);
  // proto:  QModelIndex QDirModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK9QDirModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QDirModel::QDirModel(QObject * parent);
extern void* dector_ZN9QDirModelC1EP7QObject(void* arg0);
extern void _ZN9QDirModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  QModelIndex QDirModel::index(const QString & path, int column);
extern void _ZNK9QDirModel5indexERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  bool QDirModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  void QDirModel::setLazyChildCount(bool enable);
extern void _ZN9QDirModel17setLazyChildCountEb(void* qthis, bool arg0);
  // proto:  QIcon QDirModel::fileIcon(const QModelIndex & index);
extern void _ZNK9QDirModel8fileIconERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QDirModel::hasChildren(const QModelIndex & index);
extern void _ZNK9QDirModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QDirModel::isDir(const QModelIndex & index);
extern void _ZNK9QDirModel5isDirERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QModelIndex QDirModel::mkdir(const QModelIndex & parent, const QString & name);
extern void _ZN9QDirModel5mkdirERK11QModelIndexRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  bool QDirModel::rmdir(const QModelIndex & index);
extern void _ZN9QDirModel5rmdirERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QString QDirModel::filePath(const QModelIndex & index);
extern void _ZNK9QDirModel8filePathERK11QModelIndex(void* qthis, void* arg0);
  // proto:  int QDirModel::rowCount(const QModelIndex & parent);
extern void _ZNK9QDirModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QDirModel::setReadOnly(bool enable);
extern void _ZN9QDirModel11setReadOnlyEb(void* qthis, bool arg0);
  // proto:  void QDirModel::QDirModel(const QDirModel & );
extern void* dector_ZN9QDirModelC1ERKS_(void* arg0);
extern void _ZN9QDirModelC1ERKS_(void* qthis, void* arg0);
  // proto:  void QDirModel::setResolveSymlinks(bool enable);
extern void _ZN9QDirModel18setResolveSymlinksEb(void* qthis, bool arg0);
  // proto:  bool QDirModel::lazyChildCount();
extern void _ZNK9QDirModel14lazyChildCountEv(void* qthis);
  // proto:  QFileInfo QDirModel::fileInfo(const QModelIndex & index);
extern void _ZNK9QDirModel8fileInfoERK11QModelIndex(void* qthis, void* arg0);
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

// class sizeof(QDirModel)=1
type QDirModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QFileIconProvider * QDirModel::iconProvider();
func (this *QDirModel) iconProvider(args ...interface{}) () {
  // iconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel12iconProviderEv
  default:
    qtrt.ErrorResolve("QDirModel", "iconProvider", args)
  }

}

  // proto:  QModelIndex QDirModel::parent(const QModelIndex & child);
func (this *QDirModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel6parentERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "parent", args)
  }

}

  // proto:  QVariant QDirModel::data(const QModelIndex & index, int role);
func (this *QDirModel) data(args ...interface{}) () {
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
    // invoke: _ZNK9QDirModel4dataERK11QModelIndexi
  default:
    qtrt.ErrorResolve("QDirModel", "data", args)
  }

}

  // proto:  QStringList QDirModel::nameFilters();
func (this *QDirModel) nameFilters(args ...interface{}) () {
  // nameFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel11nameFiltersEv
  default:
    qtrt.ErrorResolve("QDirModel", "nameFilters", args)
  }

}

  // proto:  bool QDirModel::isReadOnly();
func (this *QDirModel) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QDirModel", "isReadOnly", args)
  }

}

  // proto:  int QDirModel::columnCount(const QModelIndex & parent);
func (this *QDirModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel11columnCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "columnCount", args)
  }

}

  // proto:  QStringList QDirModel::mimeTypes();
func (this *QDirModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel9mimeTypesEv
  default:
    qtrt.ErrorResolve("QDirModel", "mimeTypes", args)
  }

}

  // proto:  void QDirModel::~QDirModel();
func (this *QDirModel) FreeQDirModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDirModel", "~QDirModel", args)
  }

}

  // proto:  bool QDirModel::remove(const QModelIndex & index);
func (this *QDirModel) remove(args ...interface{}) () {
  // remove(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel6removeERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "remove", args)
  }

}

  // proto:  QString QDirModel::fileName(const QModelIndex & index);
func (this *QDirModel) fileName(args ...interface{}) () {
  // fileName(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8fileNameERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "fileName", args)
  }

}

  // proto:  const QMetaObject * QDirModel::metaObject();
func (this *QDirModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel10metaObjectEv
  default:
    qtrt.ErrorResolve("QDirModel", "metaObject", args)
  }

}

  // proto:  bool QDirModel::resolveSymlinks();
func (this *QDirModel) resolveSymlinks(args ...interface{}) () {
  // resolveSymlinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel15resolveSymlinksEv
  default:
    qtrt.ErrorResolve("QDirModel", "resolveSymlinks", args)
  }

}

  // proto:  void QDirModel::refresh(const QModelIndex & parent);
func (this *QDirModel) refresh(args ...interface{}) () {
  // refresh(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel7refreshERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "refresh", args)
  }

}

  // proto:  void QDirModel::setNameFilters(const QStringList & filters);
func (this *QDirModel) setNameFilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel14setNameFiltersERK11QStringList
  default:
    qtrt.ErrorResolve("QDirModel", "setNameFilters", args)
  }

}

  // proto:  void QDirModel::setIconProvider(QFileIconProvider * provider);
func (this *QDirModel) setIconProvider(args ...interface{}) () {
  // setIconProvider(class QFileIconProvider *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel15setIconProviderEP17QFileIconProvider
  default:
    qtrt.ErrorResolve("QDirModel", "setIconProvider", args)
  }

}

  // proto:  QModelIndex QDirModel::index(int row, int column, const QModelIndex & parent);
func (this *QDirModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
  // index(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel5indexEiiRK11QModelIndex
  case 1:
    // invoke: _ZNK9QDirModel5indexERK7QStringi
  default:
    qtrt.ErrorResolve("QDirModel", "index", args)
  }

}

  // proto:  void QDirModel::QDirModel(QObject * parent);
func NewQDirModel(args ...interface{}) QDirModel {
  return QDirModel{}
}

  // proto:  bool QDirModel::setData(const QModelIndex & index, const QVariant & value, int role);
func (this *QDirModel) setData(args ...interface{}) () {
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
    // invoke: _ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti
  default:
    qtrt.ErrorResolve("QDirModel", "setData", args)
  }

}

  // proto:  void QDirModel::setLazyChildCount(bool enable);
func (this *QDirModel) setLazyChildCount(args ...interface{}) () {
  // setLazyChildCount(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel17setLazyChildCountEb
  default:
    qtrt.ErrorResolve("QDirModel", "setLazyChildCount", args)
  }

}

  // proto:  QIcon QDirModel::fileIcon(const QModelIndex & index);
func (this *QDirModel) fileIcon(args ...interface{}) () {
  // fileIcon(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8fileIconERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "fileIcon", args)
  }

}

  // proto:  bool QDirModel::hasChildren(const QModelIndex & index);
func (this *QDirModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel11hasChildrenERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "hasChildren", args)
  }

}

  // proto:  bool QDirModel::isDir(const QModelIndex & index);
func (this *QDirModel) isDir(args ...interface{}) () {
  // isDir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel5isDirERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "isDir", args)
  }

}

  // proto:  QModelIndex QDirModel::mkdir(const QModelIndex & parent, const QString & name);
func (this *QDirModel) mkdir(args ...interface{}) () {
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
    // invoke: _ZN9QDirModel5mkdirERK11QModelIndexRK7QString
  default:
    qtrt.ErrorResolve("QDirModel", "mkdir", args)
  }

}

  // proto:  bool QDirModel::rmdir(const QModelIndex & index);
func (this *QDirModel) rmdir(args ...interface{}) () {
  // rmdir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel5rmdirERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "rmdir", args)
  }

}

  // proto:  QString QDirModel::filePath(const QModelIndex & index);
func (this *QDirModel) filePath(args ...interface{}) () {
  // filePath(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8filePathERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "filePath", args)
  }

}

  // proto:  int QDirModel::rowCount(const QModelIndex & parent);
func (this *QDirModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8rowCountERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "rowCount", args)
  }

}

  // proto:  void QDirModel::setReadOnly(bool enable);
func (this *QDirModel) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel11setReadOnlyEb
  default:
    qtrt.ErrorResolve("QDirModel", "setReadOnly", args)
  }

}

  // proto:  void QDirModel::setResolveSymlinks(bool enable);
func (this *QDirModel) setResolveSymlinks(args ...interface{}) () {
  // setResolveSymlinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel18setResolveSymlinksEb
  default:
    qtrt.ErrorResolve("QDirModel", "setResolveSymlinks", args)
  }

}

  // proto:  bool QDirModel::lazyChildCount();
func (this *QDirModel) lazyChildCount(args ...interface{}) () {
  // lazyChildCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel14lazyChildCountEv
  default:
    qtrt.ErrorResolve("QDirModel", "lazyChildCount", args)
  }

}

  // proto:  QFileInfo QDirModel::fileInfo(const QModelIndex & index);
func (this *QDirModel) fileInfo(args ...interface{}) () {
  // fileInfo(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8fileInfoERK11QModelIndex
  default:
    qtrt.ErrorResolve("QDirModel", "fileInfo", args)
  }

}

// <= body block end

