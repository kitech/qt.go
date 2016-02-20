package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qfiledialog.h
// dst-file: /src/widgets/qfiledialog.go
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
  // proto:  void QFileDialog::selectFile(const QString & filename);
extern void C_ZN11QFileDialog10selectFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  QUrl QFileDialog::directoryUrl();
extern void* C_ZNK11QFileDialog12directoryUrlEv(void* qthis); // 4
  // proto:  QList<QUrl> QFileDialog::sidebarUrls();
extern void C_ZNK11QFileDialog11sidebarUrlsEv(void* qthis); // 4
  // proto:  void QFileDialog::selectUrl(const QUrl & url);
extern void C_ZN11QFileDialog9selectUrlERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::~QFileDialog();
extern void C_ZN11QFileDialogD2Ev(void* qthis); // 4
  // proto:  void QFileDialog::QFileDialog(QWidget * parent, const QString & caption, const QString & directory, const QString & filter);
extern void* C_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_(void* arg0, void* arg1, void* arg2, void* arg3); // 3
  // proto:  QByteArray QFileDialog::saveState();
extern void* C_ZNK11QFileDialog9saveStateEv(void* qthis); // 4
  // proto:  QString QFileDialog::selectedNameFilter();
extern void* C_ZNK11QFileDialog18selectedNameFilterEv(void* qthis); // 4
  // proto:  QFileDialog::FileMode QFileDialog::fileMode();
extern void C_ZNK11QFileDialog8fileModeEv(void* qthis); // 4
  // proto:  bool QFileDialog::resolveSymlinks();
extern bool C_ZNK11QFileDialog15resolveSymlinksEv(void* qthis); // 4
  // proto:  void QFileDialog::open(QObject * receiver, const char * member);
extern void C_ZN11QFileDialog4openEP7QObjectPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QFileDialog::isNameFilterDetailsVisible();
extern bool C_ZNK11QFileDialog26isNameFilterDetailsVisibleEv(void* qthis); // 4
  // proto:  QString QFileDialog::defaultSuffix();
extern void* C_ZNK11QFileDialog13defaultSuffixEv(void* qthis); // 4
  // proto:  void QFileDialog::setProxyModel(QAbstractProxyModel * model);
extern void C_ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setDirectoryUrl(const QUrl & directory);
extern void C_ZN11QFileDialog15setDirectoryUrlERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setResolveSymlinks(bool enabled);
extern void C_ZN11QFileDialog18setResolveSymlinksEb(void* qthis, bool arg0); // 4
  // proto:  bool QFileDialog::isReadOnly();
extern bool C_ZNK11QFileDialog10isReadOnlyEv(void* qthis); // 4
  // proto:  void QFileDialog::setItemDelegate(QAbstractItemDelegate * delegate);
extern void C_ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setNameFilter(const QString & filter);
extern void C_ZN11QFileDialog13setNameFilterERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAbstractProxyModel * QFileDialog::proxyModel();
extern void C_ZNK11QFileDialog10proxyModelEv(void* qthis); // 4
  // proto:  void QFileDialog::setMimeTypeFilters(const QStringList & filters);
extern void C_ZN11QFileDialog18setMimeTypeFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::selectNameFilter(const QString & filter);
extern void C_ZN11QFileDialog16selectNameFilterERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setIconProvider(QFileIconProvider * provider);
extern void C_ZN11QFileDialog15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0); // 4
  // proto:  bool QFileDialog::restoreState(const QByteArray & state);
extern bool C_ZN11QFileDialog12restoreStateERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  bool QFileDialog::confirmOverwrite();
extern bool C_ZNK11QFileDialog16confirmOverwriteEv(void* qthis); // 4
  // proto:  QList<QUrl> QFileDialog::selectedUrls();
extern void C_ZNK11QFileDialog12selectedUrlsEv(void* qthis); // 4
  // proto:  QAbstractItemDelegate * QFileDialog::itemDelegate();
extern void C_ZNK11QFileDialog12itemDelegateEv(void* qthis); // 4
  // proto:  void QFileDialog::selectMimeTypeFilter(const QString & filter);
extern void C_ZN11QFileDialog20selectMimeTypeFilterERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QFileDialog::selectedFiles();
extern void C_ZNK11QFileDialog13selectedFilesEv(void* qthis); // 4
  // proto:  void QFileDialog::setNameFilterDetailsVisible(bool enabled);
extern void C_ZN11QFileDialog27setNameFilterDetailsVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QFileDialog::setConfirmOverwrite(bool enabled);
extern void C_ZN11QFileDialog19setConfirmOverwriteEb(void* qthis, bool arg0); // 4
  // proto:  void QFileDialog::setNameFilters(const QStringList & filters);
extern void C_ZN11QFileDialog14setNameFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setDefaultSuffix(const QString & suffix);
extern void C_ZN11QFileDialog16setDefaultSuffixERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setDirectory(const QDir & directory);
extern void C_ZN11QFileDialog12setDirectoryERK4QDir(void* qthis, void* arg0); // 2
  // proto:  void QFileDialog::setDirectory(const QString & directory);
extern void C_ZN11QFileDialog12setDirectoryERK7QString(void* qthis, void* arg0); // 4
  // proto:  QFileIconProvider * QFileDialog::iconProvider();
extern void* C_ZNK11QFileDialog12iconProviderEv(void* qthis); // 4
  // proto:  const QMetaObject * QFileDialog::metaObject();
extern void C_ZNK11QFileDialog10metaObjectEv(void* qthis); // 4
  // proto:  QStringList QFileDialog::nameFilters();
extern void C_ZNK11QFileDialog11nameFiltersEv(void* qthis); // 4
  // proto:  QStringList QFileDialog::mimeTypeFilters();
extern void C_ZNK11QFileDialog15mimeTypeFiltersEv(void* qthis); // 4
  // proto:  void QFileDialog::setHistory(const QStringList & paths);
extern void C_ZN11QFileDialog10setHistoryERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QDir::Filters QFileDialog::filter();
extern void C_ZNK11QFileDialog6filterEv(void* qthis); // 4
  // proto:  void QFileDialog::setReadOnly(bool enabled);
extern void C_ZN11QFileDialog11setReadOnlyEb(void* qthis, bool arg0); // 4
  // proto:  QFileDialog::AcceptMode QFileDialog::acceptMode();
extern void C_ZNK11QFileDialog10acceptModeEv(void* qthis); // 4
  // proto:  QDir QFileDialog::directory();
extern void* C_ZNK11QFileDialog9directoryEv(void* qthis); // 4
  // proto:  QFileDialog::ViewMode QFileDialog::viewMode();
extern void C_ZNK11QFileDialog8viewModeEv(void* qthis); // 4
  // proto:  void QFileDialog::setVisible(bool visible);
extern void C_ZN11QFileDialog10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  Options QFileDialog::options();
extern void C_ZNK11QFileDialog7optionsEv(void* qthis); // 4
  // proto:  QStringList QFileDialog::history();
extern void C_ZNK11QFileDialog7historyEv(void* qthis); // 4
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

// class sizeof(QFileDialog)=1
type QFileDialog struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
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

// selectFile(const class QString &)
func (this *QFileDialog) Selectfile(args ...interface{}) () {
  // selectFile(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog10selectFileERK7QString
    // invoke: void selectFile(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog10selectFileERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectFile", args)
  }

  return
}

// directoryUrl()
func (this *QFileDialog) Directoryurl(args ...interface{}) (ret interface{}) {
  // directoryUrl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12directoryUrlEv
    // invoke: QUrl directoryUrl()
    var ret0 = C.C_ZNK11QFileDialog12directoryUrlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QUrl{}) // "QUrl"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "directoryUrl", args)
  }

  return
}

// sidebarUrls()
func (this *QFileDialog) Sidebarurls(args ...interface{}) () {
  // sidebarUrls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog11sidebarUrlsEv
    // invoke: QList<QUrl> sidebarUrls()
    C.C_ZNK11QFileDialog11sidebarUrlsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "sidebarUrls", args)
  }

  return
}

// selectUrl(const class QUrl &)
func (this *QFileDialog) Selecturl(args ...interface{}) () {
  // selectUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog9selectUrlERK4QUrl
    // invoke: void selectUrl(const class QUrl &)
    var arg0 = args[0].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog9selectUrlERK4QUrl(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectUrl", args)
  }

  return
}

// ~QFileDialog()
func (this *QFileDialog) Freeqfiledialog(args ...interface{}) () {
  // ~QFileDialog()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialogD0Ev
    // invoke: void ~QFileDialog()
    C.C_ZN11QFileDialogD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "~QFileDialog", args)
  }

  return
}

// QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
func NewQFileDialog(args ...interface{}) *QFileDialog {
  // QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialogC1EP7QWidgetRK7QStringS4_S4_
    // invoke: void QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_(arg0, arg1, arg2, arg3)
    return &QFileDialog{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFileDialog", "QFileDialog", args)
  }

  return nil // QFileDialog{Qclsinst:qthis}
}

// saveState()
func (this *QFileDialog) Savestate(args ...interface{}) (ret interface{}) {
  // saveState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog9saveStateEv
    // invoke: QByteArray saveState()
    var ret0 = C.C_ZNK11QFileDialog9saveStateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "saveState", args)
  }

  return
}

// selectedNameFilter()
func (this *QFileDialog) Selectednamefilter(args ...interface{}) (ret interface{}) {
  // selectedNameFilter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog18selectedNameFilterEv
    // invoke: QString selectedNameFilter()
    var ret0 = C.C_ZNK11QFileDialog18selectedNameFilterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedNameFilter", args)
  }

  return
}

// fileMode()
func (this *QFileDialog) Filemode(args ...interface{}) () {
  // fileMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog8fileModeEv
    // invoke: QFileDialog::FileMode fileMode()
    C.C_ZNK11QFileDialog8fileModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "fileMode", args)
  }

  return
}

// resolveSymlinks()
func (this *QFileDialog) Resolvesymlinks(args ...interface{}) (ret interface{}) {
  // resolveSymlinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog15resolveSymlinksEv
    // invoke: bool resolveSymlinks()
    var ret0 = C.C_ZNK11QFileDialog15resolveSymlinksEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "resolveSymlinks", args)
  }

  return
}

// open(class QObject *, const char *)
func (this *QFileDialog) Open(args ...interface{}) () {
  // open(class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog4openEP7QObjectPKc
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    C.C_ZN11QFileDialog4openEP7QObjectPKc(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFileDialog", "open", args)
  }

  return
}

// isNameFilterDetailsVisible()
func (this *QFileDialog) Isnamefilterdetailsvisible(args ...interface{}) (ret interface{}) {
  // isNameFilterDetailsVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog26isNameFilterDetailsVisibleEv
    // invoke: bool isNameFilterDetailsVisible()
    var ret0 = C.C_ZNK11QFileDialog26isNameFilterDetailsVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "isNameFilterDetailsVisible", args)
  }

  return
}

// defaultSuffix()
func (this *QFileDialog) Defaultsuffix(args ...interface{}) (ret interface{}) {
  // defaultSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog13defaultSuffixEv
    // invoke: QString defaultSuffix()
    var ret0 = C.C_ZNK11QFileDialog13defaultSuffixEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "defaultSuffix", args)
  }

  return
}

// setProxyModel(class QAbstractProxyModel *)
func (this *QFileDialog) Setproxymodel(args ...interface{}) () {
  // setProxyModel(class QAbstractProxyModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QAbstractProxyModel{}) // "QAbstractProxyModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel
    // invoke: void setProxyModel(class QAbstractProxyModel *)
    var arg0 = args[0].(*qtcore.QAbstractProxyModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setProxyModel", args)
  }

  return
}

// setDirectoryUrl(const class QUrl &)
func (this *QFileDialog) Setdirectoryurl(args ...interface{}) () {
  // setDirectoryUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog15setDirectoryUrlERK4QUrl
    // invoke: void setDirectoryUrl(const class QUrl &)
    var arg0 = args[0].(*qtcore.QUrl).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog15setDirectoryUrlERK4QUrl(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectoryUrl", args)
  }

  return
}

// setResolveSymlinks(_Bool)
func (this *QFileDialog) Setresolvesymlinks(args ...interface{}) () {
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
    // invoke: void setResolveSymlinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog18setResolveSymlinksEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setResolveSymlinks", args)
  }

  return
}

// isReadOnly()
func (this *QFileDialog) Isreadonly(args ...interface{}) (ret interface{}) {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10isReadOnlyEv
    // invoke: bool isReadOnly()
    var ret0 = C.C_ZNK11QFileDialog10isReadOnlyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "isReadOnly", args)
  }

  return
}

// setItemDelegate(class QAbstractItemDelegate *)
func (this *QFileDialog) Setitemdelegate(args ...interface{}) () {
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
    // invoke: void setItemDelegate(class QAbstractItemDelegate *)
    var arg0 = args[0].(*QAbstractItemDelegate).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setItemDelegate", args)
  }

  return
}

// setNameFilter(const class QString &)
func (this *QFileDialog) Setnamefilter(args ...interface{}) () {
  // setNameFilter(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog13setNameFilterERK7QString
    // invoke: void setNameFilter(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog13setNameFilterERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilter", args)
  }

  return
}

// proxyModel()
func (this *QFileDialog) Proxymodel(args ...interface{}) () {
  // proxyModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10proxyModelEv
    // invoke: QAbstractProxyModel * proxyModel()
    C.C_ZNK11QFileDialog10proxyModelEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "proxyModel", args)
  }

  return
}

// setMimeTypeFilters(const class QStringList &)
func (this *QFileDialog) Setmimetypefilters(args ...interface{}) () {
  // setMimeTypeFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog18setMimeTypeFiltersERK11QStringList
    // invoke: void setMimeTypeFilters(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog18setMimeTypeFiltersERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setMimeTypeFilters", args)
  }

  return
}

// selectNameFilter(const class QString &)
func (this *QFileDialog) Selectnamefilter(args ...interface{}) () {
  // selectNameFilter(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog16selectNameFilterERK7QString
    // invoke: void selectNameFilter(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog16selectNameFilterERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectNameFilter", args)
  }

  return
}

// setIconProvider(class QFileIconProvider *)
func (this *QFileDialog) Seticonprovider(args ...interface{}) () {
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
    // invoke: void setIconProvider(class QFileIconProvider *)
    var arg0 = args[0].(*QFileIconProvider).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog15setIconProviderEP17QFileIconProvider(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setIconProvider", args)
  }

  return
}

// restoreState(const class QByteArray &)
func (this *QFileDialog) Restorestate(args ...interface{}) (ret interface{}) {
  // restoreState(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog12restoreStateERK10QByteArray
    // invoke: bool restoreState(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QFileDialog12restoreStateERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "restoreState", args)
  }

  return
}

// confirmOverwrite()
func (this *QFileDialog) Confirmoverwrite(args ...interface{}) (ret interface{}) {
  // confirmOverwrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog16confirmOverwriteEv
    // invoke: bool confirmOverwrite()
    var ret0 = C.C_ZNK11QFileDialog16confirmOverwriteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "confirmOverwrite", args)
  }

  return
}

// selectedUrls()
func (this *QFileDialog) Selectedurls(args ...interface{}) () {
  // selectedUrls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12selectedUrlsEv
    // invoke: QList<QUrl> selectedUrls()
    C.C_ZNK11QFileDialog12selectedUrlsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedUrls", args)
  }

  return
}

// itemDelegate()
func (this *QFileDialog) Itemdelegate(args ...interface{}) () {
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12itemDelegateEv
    // invoke: QAbstractItemDelegate * itemDelegate()
    C.C_ZNK11QFileDialog12itemDelegateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "itemDelegate", args)
  }

  return
}

// selectMimeTypeFilter(const class QString &)
func (this *QFileDialog) Selectmimetypefilter(args ...interface{}) () {
  // selectMimeTypeFilter(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog20selectMimeTypeFilterERK7QString
    // invoke: void selectMimeTypeFilter(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog20selectMimeTypeFilterERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectMimeTypeFilter", args)
  }

  return
}

// selectedFiles()
func (this *QFileDialog) Selectedfiles(args ...interface{}) () {
  // selectedFiles()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog13selectedFilesEv
    // invoke: QStringList selectedFiles()
    C.C_ZNK11QFileDialog13selectedFilesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedFiles", args)
  }

  return
}

// setNameFilterDetailsVisible(_Bool)
func (this *QFileDialog) Setnamefilterdetailsvisible(args ...interface{}) () {
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
    // invoke: void setNameFilterDetailsVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog27setNameFilterDetailsVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilterDetailsVisible", args)
  }

  return
}

// setConfirmOverwrite(_Bool)
func (this *QFileDialog) Setconfirmoverwrite(args ...interface{}) () {
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
    // invoke: void setConfirmOverwrite(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog19setConfirmOverwriteEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setConfirmOverwrite", args)
  }

  return
}

// setNameFilters(const class QStringList &)
func (this *QFileDialog) Setnamefilters(args ...interface{}) () {
  // setNameFilters(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog14setNameFiltersERK11QStringList
    // invoke: void setNameFilters(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog14setNameFiltersERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilters", args)
  }

  return
}

// setDefaultSuffix(const class QString &)
func (this *QFileDialog) Setdefaultsuffix(args ...interface{}) () {
  // setDefaultSuffix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog16setDefaultSuffixERK7QString
    // invoke: void setDefaultSuffix(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog16setDefaultSuffixERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDefaultSuffix", args)
  }

  return
}

// setDirectory(const class QDir &)
func (this *QFileDialog) Setdirectory(args ...interface{}) () {
  // setDirectory(const class QDir &)
  // setDirectory(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDir{}) // "const QDir &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog12setDirectoryERK4QDir
    // invoke: void setDirectory(const class QDir &)
    var arg0 = args[0].(*qtcore.QDir).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog12setDirectoryERK4QDir(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QFileDialog12setDirectoryERK7QString
    // invoke: void setDirectory(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog12setDirectoryERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectory", args)
  }

  return
}

// iconProvider()
func (this *QFileDialog) Iconprovider(args ...interface{}) (ret interface{}) {
  // iconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog12iconProviderEv
    // invoke: QFileIconProvider * iconProvider()
    var ret0 = C.C_ZNK11QFileDialog12iconProviderEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFileIconProvider{}) // "QFileIconProvider *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "iconProvider", args)
  }

  return
}

// metaObject()
func (this *QFileDialog) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QFileDialog10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "metaObject", args)
  }

  return
}

// nameFilters()
func (this *QFileDialog) Namefilters(args ...interface{}) () {
  // nameFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog11nameFiltersEv
    // invoke: QStringList nameFilters()
    C.C_ZNK11QFileDialog11nameFiltersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "nameFilters", args)
  }

  return
}

// mimeTypeFilters()
func (this *QFileDialog) Mimetypefilters(args ...interface{}) () {
  // mimeTypeFilters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog15mimeTypeFiltersEv
    // invoke: QStringList mimeTypeFilters()
    C.C_ZNK11QFileDialog15mimeTypeFiltersEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "mimeTypeFilters", args)
  }

  return
}

// setHistory(const class QStringList &)
func (this *QFileDialog) Sethistory(args ...interface{}) () {
  // setHistory(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog10setHistoryERK11QStringList
    // invoke: void setHistory(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog10setHistoryERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setHistory", args)
  }

  return
}

// filter()
func (this *QFileDialog) Filter(args ...interface{}) () {
  // filter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog6filterEv
    // invoke: QDir::Filters filter()
    C.C_ZNK11QFileDialog6filterEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "filter", args)
  }

  return
}

// setReadOnly(_Bool)
func (this *QFileDialog) Setreadonly(args ...interface{}) () {
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
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog11setReadOnlyEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setReadOnly", args)
  }

  return
}

// acceptMode()
func (this *QFileDialog) Acceptmode(args ...interface{}) () {
  // acceptMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog10acceptModeEv
    // invoke: QFileDialog::AcceptMode acceptMode()
    C.C_ZNK11QFileDialog10acceptModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "acceptMode", args)
  }

  return
}

// directory()
func (this *QFileDialog) Directory(args ...interface{}) (ret interface{}) {
  // directory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog9directoryEv
    // invoke: QDir directory()
    var ret0 = C.C_ZNK11QFileDialog9directoryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QDir{}) // "QDir"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDialog", "directory", args)
  }

  return
}

// viewMode()
func (this *QFileDialog) Viewmode(args ...interface{}) () {
  // viewMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog8viewModeEv
    // invoke: QFileDialog::ViewMode viewMode()
    C.C_ZNK11QFileDialog8viewModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "viewMode", args)
  }

  return
}

// setVisible(_Bool)
func (this *QFileDialog) Setvisible(args ...interface{}) () {
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QFileDialog10setVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setVisible", args)
  }

  return
}

// options()
func (this *QFileDialog) Options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog7optionsEv
    // invoke: Options options()
    C.C_ZNK11QFileDialog7optionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "options", args)
  }

  return
}

// history()
func (this *QFileDialog) History(args ...interface{}) () {
  // history()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDialog7historyEv
    // invoke: QStringList history()
    C.C_ZNK11QFileDialog7historyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "history", args)
  }

  return
}

// <= body block end

