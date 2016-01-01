package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qfiledialog.h
// dst-file: /src/widgets/qfiledialog.go
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
// class sizeof(QFileDialog)=1
type QFileDialog struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
//  _filesSelected QFileDialog_filesSelected_signal;
//  _fileSelected QFileDialog_fileSelected_signal;
//  _currentChanged QFileDialog_currentChanged_signal;
//  _directoryEntered QFileDialog_directoryEntered_signal;
//  _currentUrlChanged QFileDialog_currentUrlChanged_signal;
//  _directoryUrlEntered QFileDialog_directoryUrlEntered_signal;
//  _filterSelected QFileDialog_filterSelected_signal;
//  _urlsSelected QFileDialog_urlsSelected_signal;
//  _urlSelected QFileDialog_urlSelected_signal;
}


func (this *QFileDialog) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10metaObjectEv
  default:
    qtrt.ErrorResolve("QFileDialog", "metaObject", args)
  }

}


func (this *QFileDialog) setDirectoryUrl(args ...interface{}) () {
  // setDirectoryUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog15setDirectoryUrlERK4QUrl
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectoryUrl", args)
  }

}


func (this *QFileDialog) nameFilters(args ...interface{}) () {
  // nameFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog11nameFiltersEv
  default:
    qtrt.ErrorResolve("QFileDialog", "nameFilters", args)
  }

}


func (this *QFileDialog) setConfirmOverwrite(args ...interface{}) () {
  // setConfirmOverwrite(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog19setConfirmOverwriteEb
  default:
    qtrt.ErrorResolve("QFileDialog", "setConfirmOverwrite", args)
  }

}


func (this *QFileDialog) setDefaultSuffix(args ...interface{}) () {
  // setDefaultSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog16setDefaultSuffixERK7QString
  default:
    qtrt.ErrorResolve("QFileDialog", "setDefaultSuffix", args)
  }

}


func (this *QFileDialog) setItemDelegate(args ...interface{}) () {
  // setItemDelegate(class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate
  default:
    qtrt.ErrorResolve("QFileDialog", "setItemDelegate", args)
  }

}


func (this *QFileDialog) sidebarUrls(args ...interface{}) () {
  // sidebarUrls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog11sidebarUrlsEv
  default:
    qtrt.ErrorResolve("QFileDialog", "sidebarUrls", args)
  }

}


func (this *QFileDialog) defaultSuffix(args ...interface{}) () {
  // defaultSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog13defaultSuffixEv
  default:
    qtrt.ErrorResolve("QFileDialog", "defaultSuffix", args)
  }

}


func (this *QFileDialog) setProxyModel(args ...interface{}) () {
  // setProxyModel(class QAbstractProxyModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractProxyModel{}) // "QAbstractProxyModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel
  default:
    qtrt.ErrorResolve("QFileDialog", "setProxyModel", args)
  }

}


func (this *QFileDialog) selectFile(args ...interface{}) () {
  // selectFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog10selectFileERK7QString
  default:
    qtrt.ErrorResolve("QFileDialog", "selectFile", args)
  }

}


func (this *QFileDialog) resolveSymlinks(args ...interface{}) () {
  // resolveSymlinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog15resolveSymlinksEv
  default:
    qtrt.ErrorResolve("QFileDialog", "resolveSymlinks", args)
  }

}


func (this *QFileDialog) setDirectory(args ...interface{}) () {
  // setDirectory(const class QString &)
  // setDirectory(const class QDir &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QDir{}) // "const QDir &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog12setDirectoryERK7QString
  case 1:
    // invoke: _ZN11QFileDialog12setDirectoryERK4QDir
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectory", args)
  }

}


func (this *QFileDialog) selectUrl(args ...interface{}) () {
  // selectUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog9selectUrlERK4QUrl
  default:
    qtrt.ErrorResolve("QFileDialog", "selectUrl", args)
  }

}


func (this *QFileDialog) selectedNameFilter(args ...interface{}) () {
  // selectedNameFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog18selectedNameFilterEv
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedNameFilter", args)
  }

}


func (this *QFileDialog) directoryUrl(args ...interface{}) () {
  // directoryUrl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12directoryUrlEv
  default:
    qtrt.ErrorResolve("QFileDialog", "directoryUrl", args)
  }

}


func (this *QFileDialog) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QFileDialog", "isReadOnly", args)
  }

}


func (this *QFileDialog) saveState(args ...interface{}) () {
  // saveState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog9saveStateEv
  default:
    qtrt.ErrorResolve("QFileDialog", "saveState", args)
  }

}


func (this *QFileDialog) open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog4openEP7QObjectPKc
  default:
    qtrt.ErrorResolve("QFileDialog", "open", args)
  }

}


func (this *QFileDialog) directory(args ...interface{}) () {
  // directory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog9directoryEv
  default:
    qtrt.ErrorResolve("QFileDialog", "directory", args)
  }

}


func (this *QFileDialog) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog10setVisibleEb
  default:
    qtrt.ErrorResolve("QFileDialog", "setVisible", args)
  }

}


func (this *QFileDialog) setIconProvider(args ...interface{}) () {
  // setIconProvider(class QFileIconProvider *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog15setIconProviderEP17QFileIconProvider
  default:
    qtrt.ErrorResolve("QFileDialog", "setIconProvider", args)
  }

}


func (this *QFileDialog) selectMimeTypeFilter(args ...interface{}) () {
  // selectMimeTypeFilter(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog20selectMimeTypeFilterERK7QString
  default:
    qtrt.ErrorResolve("QFileDialog", "selectMimeTypeFilter", args)
  }

}


func (this *QFileDialog) mimeTypeFilters(args ...interface{}) () {
  // mimeTypeFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog15mimeTypeFiltersEv
  default:
    qtrt.ErrorResolve("QFileDialog", "mimeTypeFilters", args)
  }

}


func (this *QFileDialog) setMimeTypeFilters(args ...interface{}) () {
  // setMimeTypeFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog18setMimeTypeFiltersERK11QStringList
  default:
    qtrt.ErrorResolve("QFileDialog", "setMimeTypeFilters", args)
  }

}


func (this *QFileDialog) setResolveSymlinks(args ...interface{}) () {
  // setResolveSymlinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog18setResolveSymlinksEb
  default:
    qtrt.ErrorResolve("QFileDialog", "setResolveSymlinks", args)
  }

}


func (this *QFileDialog) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog11setReadOnlyEb
  default:
    qtrt.ErrorResolve("QFileDialog", "setReadOnly", args)
  }

}


func (this *QFileDialog) setNameFilterDetailsVisible(args ...interface{}) () {
  // setNameFilterDetailsVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog27setNameFilterDetailsVisibleEb
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilterDetailsVisible", args)
  }

}


func (this *QFileDialog) selectNameFilter(args ...interface{}) () {
  // selectNameFilter(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog16selectNameFilterERK7QString
  default:
    qtrt.ErrorResolve("QFileDialog", "selectNameFilter", args)
  }

}


func (this *QFileDialog) restoreState(args ...interface{}) () {
  // restoreState(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog12restoreStateERK10QByteArray
  default:
    qtrt.ErrorResolve("QFileDialog", "restoreState", args)
  }

}


func (this *QFileDialog) iconProvider(args ...interface{}) () {
  // iconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12iconProviderEv
  default:
    qtrt.ErrorResolve("QFileDialog", "iconProvider", args)
  }

}


func (this *QFileDialog) selectedFiles(args ...interface{}) () {
  // selectedFiles()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog13selectedFilesEv
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedFiles", args)
  }

}


func (this *QFileDialog) FreeQFileDialog(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileDialog", "~QFileDialog", args)
  }

}


func (this *QFileDialog) itemDelegate(args ...interface{}) () {
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12itemDelegateEv
  default:
    qtrt.ErrorResolve("QFileDialog", "itemDelegate", args)
  }

}


func (this *QFileDialog) confirmOverwrite(args ...interface{}) () {
  // confirmOverwrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog16confirmOverwriteEv
  default:
    qtrt.ErrorResolve("QFileDialog", "confirmOverwrite", args)
  }

}


func (this *QFileDialog) setHistory(args ...interface{}) () {
  // setHistory(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog10setHistoryERK11QStringList
  default:
    qtrt.ErrorResolve("QFileDialog", "setHistory", args)
  }

}


func (this *QFileDialog) setNameFilter(args ...interface{}) () {
  // setNameFilter(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog13setNameFilterERK7QString
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilter", args)
  }

}


func (this *QFileDialog) proxyModel(args ...interface{}) () {
  // proxyModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10proxyModelEv
  default:
    qtrt.ErrorResolve("QFileDialog", "proxyModel", args)
  }

}


func (this *QFileDialog) setNameFilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog14setNameFiltersERK11QStringList
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilters", args)
  }

}


func (this *QFileDialog) selectedUrls(args ...interface{}) () {
  // selectedUrls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12selectedUrlsEv
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedUrls", args)
  }

}


func (this *QFileDialog) history(args ...interface{}) () {
  // history()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog7historyEv
  default:
    qtrt.ErrorResolve("QFileDialog", "history", args)
  }

}


func (this *QFileDialog) isNameFilterDetailsVisible(args ...interface{}) () {
  // isNameFilterDetailsVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog26isNameFilterDetailsVisibleEv
  default:
    qtrt.ErrorResolve("QFileDialog", "isNameFilterDetailsVisible", args)
  }

}


func NewQFileDialog(args ...interface{}) QFileDialog {
  return QFileDialog{}
}

// <= body block end

