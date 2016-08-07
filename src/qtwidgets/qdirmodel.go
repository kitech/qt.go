package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QDirModel::columnCount(const QModelIndex & parent);
extern int32_t C_ZNK9QDirModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDirModel::setNameFilters(const QStringList & filters);
extern void C_ZN9QDirModel14setNameFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QDirModel::setIconProvider(QFileIconProvider * provider);
extern void C_ZN9QDirModel15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0); // 4
  // proto:  void QDirModel::setLazyChildCount(bool enable);
extern void C_ZN9QDirModel17setLazyChildCountEb(void* qthis, bool arg0); // 4
  // proto:  bool QDirModel::isDir(const QModelIndex & index);
extern bool C_ZNK9QDirModel5isDirERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QDirModel::lazyChildCount();
extern bool C_ZNK9QDirModel14lazyChildCountEv(void* qthis); // 4
  // proto:  void QDirModel::refresh(const QModelIndex & parent);
extern void C_ZN9QDirModel7refreshERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QDirModel::resolveSymlinks();
extern bool C_ZNK9QDirModel15resolveSymlinksEv(void* qthis); // 4
  // proto:  QModelIndex QDirModel::index(int row, int column, const QModelIndex & parent);
extern void* C_ZNK9QDirModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QDirModel::index(const QString & path, int column);
extern void* C_ZNK9QDirModel5indexERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QDirModel::hasChildren(const QModelIndex & index);
extern bool C_ZNK9QDirModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QDirModel::isReadOnly();
extern bool C_ZNK9QDirModel10isReadOnlyEv(void* qthis); // 4
  // proto:  QString QDirModel::filePath(const QModelIndex & index);
extern void* C_ZNK9QDirModel8filePathERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QDirModel::mkdir(const QModelIndex & parent, const QString & name);
extern void* C_ZN9QDirModel5mkdirERK11QModelIndexRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QDirModel::rmdir(const QModelIndex & index);
extern bool C_ZN9QDirModel5rmdirERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDirModel::setResolveSymlinks(bool enable);
extern void C_ZN9QDirModel18setResolveSymlinksEb(void* qthis, bool arg0); // 4
  // proto:  QFileInfo QDirModel::fileInfo(const QModelIndex & index);
extern void* C_ZNK9QDirModel8fileInfoERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  bool QDirModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern bool C_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  QModelIndex QDirModel::parent(const QModelIndex & child);
extern void* C_ZNK9QDirModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QString QDirModel::fileName(const QModelIndex & index);
extern void* C_ZNK9QDirModel8fileNameERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  int QDirModel::rowCount(const QModelIndex & parent);
extern int32_t C_ZNK9QDirModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QDirModel::~QDirModel();
extern void C_ZN9QDirModelD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QDirModel::metaObject();
extern void C_ZNK9QDirModel10metaObjectEv(void* qthis); // 4
  // proto:  QVariant QDirModel::data(const QModelIndex & index, int role);
extern void* C_ZNK9QDirModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringList QDirModel::mimeTypes();
extern void C_ZNK9QDirModel9mimeTypesEv(void* qthis); // 4
  // proto:  QDir::SortFlags QDirModel::sorting();
extern void C_ZNK9QDirModel7sortingEv(void* qthis); // 4
  // proto:  QFileIconProvider * QDirModel::iconProvider();
extern void* C_ZNK9QDirModel12iconProviderEv(void* qthis); // 4
  // proto:  QStringList QDirModel::nameFilters();
extern void C_ZNK9QDirModel11nameFiltersEv(void* qthis); // 4
  // proto:  void QDirModel::QDirModel(QObject * parent);
extern void* C_ZN9QDirModelC2EP7QObject(void* arg0); // 3
  // proto:  bool QDirModel::remove(const QModelIndex & index);
extern bool C_ZN9QDirModel6removeERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QDir::Filters QDirModel::filter();
extern void C_ZNK9QDirModel6filterEv(void* qthis); // 4
  // proto:  void QDirModel::setReadOnly(bool enable);
extern void C_ZN9QDirModel11setReadOnlyEb(void* qthis, bool arg0); // 4
  // proto:  Qt::ItemFlags QDirModel::flags(const QModelIndex & index);
extern void C_ZNK9QDirModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QIcon QDirModel::fileIcon(const QModelIndex & index);
extern void* C_ZNK9QDirModel8fileIconERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QDirModel::supportedDropActions();
extern void C_ZNK9QDirModel20supportedDropActionsEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QDirModel)=1
type QDirModel struct {
  /*qbase*/ qtcore.QAbstractItemModel;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount(const class QModelIndex &)
func (this *QDirModel) Columncount(args ...interface{}) (ret interface{}) {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel11columnCountERK11QModelIndex
    // invoke: int columnCount(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel11columnCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "columnCount", args)
  }

  return
}

// setNameFilters(const class QStringList &)
func (this *QDirModel) Setnamefilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel14setNameFiltersERK11QStringList
    // invoke: void setNameFilters(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDirModel14setNameFiltersERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "setNameFilters", args)
  }

  return
}

// setIconProvider(class QFileIconProvider *)
func (this *QDirModel) Seticonprovider(args ...interface{}) () {
  // setIconProvider(class QFileIconProvider *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel15setIconProviderEP17QFileIconProvider
    // invoke: void setIconProvider(class QFileIconProvider *)
    var arg0 = args[0].(*QFileIconProvider).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDirModel15setIconProviderEP17QFileIconProvider(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "setIconProvider", args)
  }

  return
}

// setLazyChildCount(_Bool)
func (this *QDirModel) Setlazychildcount(args ...interface{}) () {
  // setLazyChildCount(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel17setLazyChildCountEb
    // invoke: void setLazyChildCount(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDirModel17setLazyChildCountEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "setLazyChildCount", args)
  }

  return
}

// isDir(const class QModelIndex &)
func (this *QDirModel) Isdir(args ...interface{}) (ret interface{}) {
  // isDir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel5isDirERK11QModelIndex
    // invoke: bool isDir(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel5isDirERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "isDir", args)
  }

  return
}

// lazyChildCount()
func (this *QDirModel) Lazychildcount(args ...interface{}) (ret interface{}) {
  // lazyChildCount()
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
    // invoke: _ZNK9QDirModel14lazyChildCountEv
    // invoke: bool lazyChildCount()
    var ret0 = C.C_ZNK9QDirModel14lazyChildCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "lazyChildCount", args)
  }

  return
}

// refresh(const class QModelIndex &)
func (this *QDirModel) Refresh(args ...interface{}) () {
  // refresh(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel7refreshERK11QModelIndex
    // invoke: void refresh(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QDirModel7refreshERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "refresh", args)
  }

  return
}

// resolveSymlinks()
func (this *QDirModel) Resolvesymlinks(args ...interface{}) (ret interface{}) {
  // resolveSymlinks()
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
    // invoke: _ZNK9QDirModel15resolveSymlinksEv
    // invoke: bool resolveSymlinks()
    var ret0 = C.C_ZNK9QDirModel15resolveSymlinksEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "resolveSymlinks", args)
  }

  return
}

// index(int, int, const class QModelIndex &)
func (this *QDirModel) Index(args ...interface{}) (ret interface{}) {
  // index(int, int, const class QModelIndex &)
  // index(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel5indexEiiRK11QModelIndex
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK9QDirModel5indexEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK9QDirModel5indexERK7QStringi
    // invoke: QModelIndex index(const class QString &, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QDirModel5indexERK7QStringi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "index", args)
  }

  return
}

// hasChildren(const class QModelIndex &)
func (this *QDirModel) Haschildren(args ...interface{}) (ret interface{}) {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel11hasChildrenERK11QModelIndex
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel11hasChildrenERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "hasChildren", args)
  }

  return
}

// isReadOnly()
func (this *QDirModel) Isreadonly(args ...interface{}) (ret interface{}) {
  // isReadOnly()
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
    // invoke: _ZNK9QDirModel10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret0 = C.C_ZNK9QDirModel10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "isReadOnly", args)
  }

  return
}

// filePath(const class QModelIndex &)
func (this *QDirModel) Filepath(args ...interface{}) (ret interface{}) {
  // filePath(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8filePathERK11QModelIndex
    // invoke: QString filePath(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel8filePathERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "filePath", args)
  }

  return
}

// mkdir(const class QModelIndex &, const class QString &)
func (this *QDirModel) Mkdir(args ...interface{}) (ret interface{}) {
  // mkdir(const class QModelIndex &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel5mkdirERK11QModelIndexRK7QString
    // invoke: QModelIndex mkdir(const class QModelIndex &, const class QString &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN9QDirModel5mkdirERK11QModelIndexRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "mkdir", args)
  }

  return
}

// rmdir(const class QModelIndex &)
func (this *QDirModel) Rmdir(args ...interface{}) (ret interface{}) {
  // rmdir(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel5rmdirERK11QModelIndex
    // invoke: bool rmdir(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QDirModel5rmdirERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "rmdir", args)
  }

  return
}

// setResolveSymlinks(_Bool)
func (this *QDirModel) Setresolvesymlinks(args ...interface{}) () {
  // setResolveSymlinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel18setResolveSymlinksEb
    // invoke: void setResolveSymlinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDirModel18setResolveSymlinksEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "setResolveSymlinks", args)
  }

  return
}

// fileInfo(const class QModelIndex &)
func (this *QDirModel) Fileinfo(args ...interface{}) (ret interface{}) {
  // fileInfo(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8fileInfoERK11QModelIndex
    // invoke: QFileInfo fileInfo(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel8fileInfoERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QFileInfo{}) // "QFileInfo"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "fileInfo", args)
  }

  return
}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QDirModel) Setdata(args ...interface{}) (ret interface{}) {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN9QDirModel7setDataERK11QModelIndexRK8QVarianti(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "setData", args)
  }

  return
}

// parent(const class QModelIndex &)
func (this *QDirModel) Parent(args ...interface{}) (ret interface{}) {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel6parentERK11QModelIndex
    // invoke: QModelIndex parent(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel6parentERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "parent", args)
  }

  return
}

// fileName(const class QModelIndex &)
func (this *QDirModel) Filename(args ...interface{}) (ret interface{}) {
  // fileName(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8fileNameERK11QModelIndex
    // invoke: QString fileName(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel8fileNameERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "fileName", args)
  }

  return
}

// rowCount(const class QModelIndex &)
func (this *QDirModel) Rowcount(args ...interface{}) (ret interface{}) {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel8rowCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "rowCount", args)
  }

  return
}

// ~QDirModel()
func (this *QDirModel) Freeqdirmodel(args ...interface{}) () {
  // ~QDirModel()
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
    // invoke: _ZN9QDirModelD0Ev
    // invoke: void ~QDirModel()
    C.C_ZN9QDirModelD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "~QDirModel", args)
  }

  return
}

// metaObject()
func (this *QDirModel) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK9QDirModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QDirModel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "metaObject", args)
  }

  return
}

// data(const class QModelIndex &, int)
func (this *QDirModel) Data(args ...interface{}) (ret interface{}) {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK9QDirModel4dataERK11QModelIndexi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "data", args)
  }

  return
}

// mimeTypes()
func (this *QDirModel) Mimetypes(args ...interface{}) () {
  // mimeTypes()
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
    // invoke: _ZNK9QDirModel9mimeTypesEv
    // invoke: QStringList mimeTypes()
    C.C_ZNK9QDirModel9mimeTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "mimeTypes", args)
  }

  return
}

// sorting()
func (this *QDirModel) Sorting(args ...interface{}) () {
  // sorting()
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
    // invoke: _ZNK9QDirModel7sortingEv
    // invoke: QDir::SortFlags sorting()
    C.C_ZNK9QDirModel7sortingEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "sorting", args)
  }

  return
}

// iconProvider()
func (this *QDirModel) Iconprovider(args ...interface{}) (ret interface{}) {
  // iconProvider()
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
    // invoke: _ZNK9QDirModel12iconProviderEv
    // invoke: QFileIconProvider * iconProvider()
    var ret0 = C.C_ZNK9QDirModel12iconProviderEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "iconProvider", args)
  }

  return
}

// nameFilters()
func (this *QDirModel) Namefilters(args ...interface{}) () {
  // nameFilters()
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
    // invoke: _ZNK9QDirModel11nameFiltersEv
    // invoke: QStringList nameFilters()
    C.C_ZNK9QDirModel11nameFiltersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "nameFilters", args)
  }

  return
}

// QDirModel(class QObject *)
func NewQDirModel(args ...interface{}) *QDirModel {
  // QDirModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModelC1EP7QObject
    // invoke: void QDirModel(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QDirModelC2EP7QObject(arg0)
    return &QDirModel{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDirModel", "QDirModel", args)
  }

  return nil // QDirModel{Qclsinst:qthis}
}

// remove(const class QModelIndex &)
func (this *QDirModel) Remove(args ...interface{}) (ret interface{}) {
  // remove(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel6removeERK11QModelIndex
    // invoke: bool remove(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN9QDirModel6removeERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "remove", args)
  }

  return
}

// filter()
func (this *QDirModel) Filter(args ...interface{}) () {
  // filter()
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
    // invoke: _ZNK9QDirModel6filterEv
    // invoke: QDir::Filters filter()
    C.C_ZNK9QDirModel6filterEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "filter", args)
  }

  return
}

// setReadOnly(_Bool)
func (this *QDirModel) Setreadonly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QDirModel11setReadOnlyEb
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN9QDirModel11setReadOnlyEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "setReadOnly", args)
  }

  return
}

// flags(const class QModelIndex &)
func (this *QDirModel) Flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK9QDirModel5flagsERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDirModel", "flags", args)
  }

  return
}

// fileIcon(const class QModelIndex &)
func (this *QDirModel) Fileicon(args ...interface{}) (ret interface{}) {
  // fileIcon(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QDirModel8fileIconERK11QModelIndex
    // invoke: QIcon fileIcon(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QDirModel8fileIconERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QDirModel", "fileIcon", args)
  }

  return
}

// supportedDropActions()
func (this *QDirModel) Supporteddropactions(args ...interface{}) () {
  // supportedDropActions()
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
    // invoke: _ZNK9QDirModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C.C_ZNK9QDirModel20supportedDropActionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QDirModel", "supportedDropActions", args)
  }

  return
}

// <= body block end

