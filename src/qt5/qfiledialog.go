package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QFileDialog::selectFile(const QString & filename);
extern void _ZN11QFileDialog10selectFileERK7QString(void* qthis, void* arg0); // 4
  // proto:  QUrl QFileDialog::directoryUrl();
extern void _ZNK11QFileDialog12directoryUrlEv(void* qthis); // 4
  // proto:  QList<QUrl> QFileDialog::sidebarUrls();
extern void _ZNK11QFileDialog11sidebarUrlsEv(void* qthis); // 4
  // proto:  void QFileDialog::selectUrl(const QUrl & url);
extern void _ZN11QFileDialog9selectUrlERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::~QFileDialog();
extern void _ZN11QFileDialogD2Ev(void* qthis); // 4
  // proto:  void QFileDialog::QFileDialog(QWidget * parent, const QString & caption, const QString & directory, const QString & filter);
extern void _ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 3
  // proto:  QByteArray QFileDialog::saveState();
extern void _ZNK11QFileDialog9saveStateEv(void* qthis); // 4
  // proto:  QString QFileDialog::selectedNameFilter();
extern void _ZNK11QFileDialog18selectedNameFilterEv(void* qthis); // 4
  // proto:  QFileDialog::FileMode QFileDialog::fileMode();
extern void _ZNK11QFileDialog8fileModeEv(void* qthis); // 4
  // proto:  bool QFileDialog::resolveSymlinks();
extern void _ZNK11QFileDialog15resolveSymlinksEv(void* qthis); // 4
  // proto:  void QFileDialog::open(QObject * receiver, const char * member);
extern void _ZN11QFileDialog4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  bool QFileDialog::isNameFilterDetailsVisible();
extern void _ZNK11QFileDialog26isNameFilterDetailsVisibleEv(void* qthis); // 4
  // proto:  QString QFileDialog::defaultSuffix();
extern void _ZNK11QFileDialog13defaultSuffixEv(void* qthis); // 4
  // proto:  void QFileDialog::setProxyModel(QAbstractProxyModel * model);
extern void _ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setDirectoryUrl(const QUrl & directory);
extern void _ZN11QFileDialog15setDirectoryUrlERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setResolveSymlinks(bool enabled);
extern void _ZN11QFileDialog18setResolveSymlinksEb(void* qthis, bool arg0); // 4
  // proto:  bool QFileDialog::isReadOnly();
extern void _ZNK11QFileDialog10isReadOnlyEv(void* qthis); // 4
  // proto:  void QFileDialog::setItemDelegate(QAbstractItemDelegate * delegate);
extern void _ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setNameFilter(const QString & filter);
extern void _ZN11QFileDialog13setNameFilterERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAbstractProxyModel * QFileDialog::proxyModel();
extern void _ZNK11QFileDialog10proxyModelEv(void* qthis); // 4
  // proto:  void QFileDialog::setMimeTypeFilters(const QStringList & filters);
extern void _ZN11QFileDialog18setMimeTypeFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::selectNameFilter(const QString & filter);
extern void _ZN11QFileDialog16selectNameFilterERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setIconProvider(QFileIconProvider * provider);
extern void _ZN11QFileDialog15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0); // 4
  // proto:  bool QFileDialog::restoreState(const QByteArray & state);
extern void _ZN11QFileDialog12restoreStateERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  bool QFileDialog::confirmOverwrite();
extern void _ZNK11QFileDialog16confirmOverwriteEv(void* qthis); // 4
  // proto:  QList<QUrl> QFileDialog::selectedUrls();
extern void _ZNK11QFileDialog12selectedUrlsEv(void* qthis); // 4
  // proto:  QAbstractItemDelegate * QFileDialog::itemDelegate();
extern void _ZNK11QFileDialog12itemDelegateEv(void* qthis); // 4
  // proto:  void QFileDialog::selectMimeTypeFilter(const QString & filter);
extern void _ZN11QFileDialog20selectMimeTypeFilterERK7QString(void* qthis, void* arg0); // 4
  // proto:  QStringList QFileDialog::selectedFiles();
extern void _ZNK11QFileDialog13selectedFilesEv(void* qthis); // 4
  // proto:  void QFileDialog::setNameFilterDetailsVisible(bool enabled);
extern void _ZN11QFileDialog27setNameFilterDetailsVisibleEb(void* qthis, bool arg0); // 4
  // proto:  void QFileDialog::setConfirmOverwrite(bool enabled);
extern void _ZN11QFileDialog19setConfirmOverwriteEb(void* qthis, bool arg0); // 4
  // proto:  void QFileDialog::setNameFilters(const QStringList & filters);
extern void _ZN11QFileDialog14setNameFiltersERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setDefaultSuffix(const QString & suffix);
extern void _ZN11QFileDialog16setDefaultSuffixERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFileDialog::setDirectory(const QDir & directory);
extern void _ZN11QFileDialog12setDirectoryERK4QDir(void* qthis, void* arg0); // 2
  // proto:  void QFileDialog::setDirectory(const QString & directory);
extern void _ZN11QFileDialog12setDirectoryERK7QString(void* qthis, void* arg0); // 4
  // proto:  QFileIconProvider * QFileDialog::iconProvider();
extern void _ZNK11QFileDialog12iconProviderEv(void* qthis); // 4
  // proto:  const QMetaObject * QFileDialog::metaObject();
extern void _ZNK11QFileDialog10metaObjectEv(void* qthis); // 4
  // proto:  QStringList QFileDialog::nameFilters();
extern void _ZNK11QFileDialog11nameFiltersEv(void* qthis); // 4
  // proto:  QStringList QFileDialog::mimeTypeFilters();
extern void _ZNK11QFileDialog15mimeTypeFiltersEv(void* qthis); // 4
  // proto:  void QFileDialog::setHistory(const QStringList & paths);
extern void _ZN11QFileDialog10setHistoryERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QDir::Filters QFileDialog::filter();
extern void _ZNK11QFileDialog6filterEv(void* qthis); // 4
  // proto:  void QFileDialog::setReadOnly(bool enabled);
extern void _ZN11QFileDialog11setReadOnlyEb(void* qthis, bool arg0); // 4
  // proto:  QFileDialog::AcceptMode QFileDialog::acceptMode();
extern void _ZNK11QFileDialog10acceptModeEv(void* qthis); // 4
  // proto:  QDir QFileDialog::directory();
extern void _ZNK11QFileDialog9directoryEv(void* qthis); // 4
  // proto:  QFileDialog::ViewMode QFileDialog::viewMode();
extern void _ZNK11QFileDialog8viewModeEv(void* qthis); // 4
  // proto:  void QFileDialog::setVisible(bool visible);
extern void _ZN11QFileDialog10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  Options QFileDialog::options();
extern void _ZNK11QFileDialog7optionsEv(void* qthis); // 4
  // proto:  QStringList QFileDialog::history();
extern void _ZNK11QFileDialog7historyEv(void* qthis); // 4
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

// class sizeof(QFileDialog)=1
type QFileDialog struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
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
    // invoke: void selectFile(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog10selectFileERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectFile", args)
  }

}

// directoryUrl()
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
    // invoke: QUrl directoryUrl()
    C._ZNK11QFileDialog12directoryUrlEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "directoryUrl", args)
  }

}

// sidebarUrls()
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
    // invoke: QList<QUrl> sidebarUrls()
    C._ZNK11QFileDialog11sidebarUrlsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "sidebarUrls", args)
  }

}

// selectUrl(const class QUrl &)
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
    // invoke: void selectUrl(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog9selectUrlERK4QUrl(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectUrl", args)
  }

}

// ~QFileDialog()
func (this *QFileDialog) FreeQFileDialog(args ...interface{}) () {
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
    C._ZN11QFileDialogD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "~QFileDialog", args)
  }

}

// QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
func NewQFileDialog(args ...interface{}) QFileDialog {
  // QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialogC1EP7QWidgetRK7QStringS4_S4_
    // invoke: void QFileDialog(class QWidget *, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN11QFileDialogC2EP7QWidgetRK7QStringS4_S4_(qthis, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QFileDialog", "QFileDialog", args)
  }

  return QFileDialog{}
}

// saveState()
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
    // invoke: QByteArray saveState()
    C._ZNK11QFileDialog9saveStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "saveState", args)
  }

}

// selectedNameFilter()
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
    // invoke: QString selectedNameFilter()
    C._ZNK11QFileDialog18selectedNameFilterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedNameFilter", args)
  }

}

// fileMode()
func (this *QFileDialog) fileMode(args ...interface{}) () {
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
    C._ZNK11QFileDialog8fileModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "fileMode", args)
  }

}

// resolveSymlinks()
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
    // invoke: bool resolveSymlinks()
    C._ZNK11QFileDialog15resolveSymlinksEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "resolveSymlinks", args)
  }

}

// open(class QObject *, const char *)
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
    // invoke: void open(class QObject *, const char *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN11QFileDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFileDialog", "open", args)
  }

}

// isNameFilterDetailsVisible()
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
    // invoke: bool isNameFilterDetailsVisible()
    C._ZNK11QFileDialog26isNameFilterDetailsVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "isNameFilterDetailsVisible", args)
  }

}

// defaultSuffix()
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
    // invoke: QString defaultSuffix()
    C._ZNK11QFileDialog13defaultSuffixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "defaultSuffix", args)
  }

}

// setProxyModel(class QAbstractProxyModel *)
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
    // invoke: void setProxyModel(class QAbstractProxyModel *)
    var arg0 = args[0].(QAbstractProxyModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setProxyModel", args)
  }

}

// setDirectoryUrl(const class QUrl &)
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
    // invoke: void setDirectoryUrl(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog15setDirectoryUrlERK4QUrl(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectoryUrl", args)
  }

}

// setResolveSymlinks(_Bool)
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
    // invoke: void setResolveSymlinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog18setResolveSymlinksEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setResolveSymlinks", args)
  }

}

// isReadOnly()
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
    // invoke: bool isReadOnly()
    C._ZNK11QFileDialog10isReadOnlyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "isReadOnly", args)
  }

}

// setItemDelegate(class QAbstractItemDelegate *)
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
    // invoke: void setItemDelegate(class QAbstractItemDelegate *)
    var arg0 = args[0].(QAbstractItemDelegate).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setItemDelegate", args)
  }

}

// setNameFilter(const class QString &)
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
    // invoke: void setNameFilter(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog13setNameFilterERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilter", args)
  }

}

// proxyModel()
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
    // invoke: QAbstractProxyModel * proxyModel()
    C._ZNK11QFileDialog10proxyModelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "proxyModel", args)
  }

}

// setMimeTypeFilters(const class QStringList &)
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
    // invoke: void setMimeTypeFilters(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog18setMimeTypeFiltersERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setMimeTypeFilters", args)
  }

}

// selectNameFilter(const class QString &)
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
    // invoke: void selectNameFilter(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog16selectNameFilterERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectNameFilter", args)
  }

}

// setIconProvider(class QFileIconProvider *)
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
    // invoke: void setIconProvider(class QFileIconProvider *)
    var arg0 = args[0].(QFileIconProvider).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog15setIconProviderEP17QFileIconProvider(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setIconProvider", args)
  }

}

// restoreState(const class QByteArray &)
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
    // invoke: bool restoreState(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog12restoreStateERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "restoreState", args)
  }

}

// confirmOverwrite()
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
    // invoke: bool confirmOverwrite()
    C._ZNK11QFileDialog16confirmOverwriteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "confirmOverwrite", args)
  }

}

// selectedUrls()
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
    // invoke: QList<QUrl> selectedUrls()
    C._ZNK11QFileDialog12selectedUrlsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedUrls", args)
  }

}

// itemDelegate()
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
    // invoke: QAbstractItemDelegate * itemDelegate()
    C._ZNK11QFileDialog12itemDelegateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "itemDelegate", args)
  }

}

// selectMimeTypeFilter(const class QString &)
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
    // invoke: void selectMimeTypeFilter(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog20selectMimeTypeFilterERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectMimeTypeFilter", args)
  }

}

// selectedFiles()
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
    // invoke: QStringList selectedFiles()
    C._ZNK11QFileDialog13selectedFilesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "selectedFiles", args)
  }

}

// setNameFilterDetailsVisible(_Bool)
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
    // invoke: void setNameFilterDetailsVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog27setNameFilterDetailsVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilterDetailsVisible", args)
  }

}

// setConfirmOverwrite(_Bool)
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
    // invoke: void setConfirmOverwrite(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog19setConfirmOverwriteEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setConfirmOverwrite", args)
  }

}

// setNameFilters(const class QStringList &)
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
    // invoke: void setNameFilters(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog14setNameFiltersERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setNameFilters", args)
  }

}

// setDefaultSuffix(const class QString &)
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
    // invoke: void setDefaultSuffix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog16setDefaultSuffixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDefaultSuffix", args)
  }

}

// setDirectory(const class QDir &)
func (this *QFileDialog) setDirectory(args ...interface{}) () {
  // setDirectory(const class QDir &)
  // setDirectory(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDialog12setDirectoryERK4QDir
    // invoke: void setDirectory(const class QDir &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog12setDirectoryERK4QDir(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QFileDialog12setDirectoryERK7QString
    // invoke: void setDirectory(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog12setDirectoryERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectory", args)
  }

}

// iconProvider()
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
    // invoke: QFileIconProvider * iconProvider()
    C._ZNK11QFileDialog12iconProviderEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "iconProvider", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QFileDialog10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "metaObject", args)
  }

}

// nameFilters()
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
    // invoke: QStringList nameFilters()
    C._ZNK11QFileDialog11nameFiltersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "nameFilters", args)
  }

}

// mimeTypeFilters()
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
    // invoke: QStringList mimeTypeFilters()
    C._ZNK11QFileDialog15mimeTypeFiltersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "mimeTypeFilters", args)
  }

}

// setHistory(const class QStringList &)
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
    // invoke: void setHistory(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog10setHistoryERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setHistory", args)
  }

}

// filter()
func (this *QFileDialog) filter(args ...interface{}) () {
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
    C._ZNK11QFileDialog6filterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "filter", args)
  }

}

// setReadOnly(_Bool)
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
    // invoke: void setReadOnly(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog11setReadOnlyEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setReadOnly", args)
  }

}

// acceptMode()
func (this *QFileDialog) acceptMode(args ...interface{}) () {
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
    C._ZNK11QFileDialog10acceptModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "acceptMode", args)
  }

}

// directory()
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
    // invoke: QDir directory()
    C._ZNK11QFileDialog9directoryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "directory", args)
  }

}

// viewMode()
func (this *QFileDialog) viewMode(args ...interface{}) () {
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
    C._ZNK11QFileDialog8viewModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "viewMode", args)
  }

}

// setVisible(_Bool)
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setVisible", args)
  }

}

// options()
func (this *QFileDialog) options(args ...interface{}) () {
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
    C._ZNK11QFileDialog7optionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "options", args)
  }

}

// history()
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
    // invoke: QStringList history()
    C._ZNK11QFileDialog7historyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDialog", "history", args)
  }

}

// <= body block end

