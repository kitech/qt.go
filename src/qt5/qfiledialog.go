package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  const QMetaObject * QFileDialog::metaObject();
extern void _ZNK11QFileDialog10metaObjectEv(void* qthis);
  // proto:  void QFileDialog::setDirectoryUrl(const QUrl & directory);
extern void _ZN11QFileDialog15setDirectoryUrlERK4QUrl(void* qthis, void* arg0);
  // proto:  QStringList QFileDialog::nameFilters();
extern void _ZNK11QFileDialog11nameFiltersEv(void* qthis);
  // proto:  void QFileDialog::setConfirmOverwrite(bool enabled);
extern void _ZN11QFileDialog19setConfirmOverwriteEb(void* qthis, bool arg0);
  // proto:  void QFileDialog::setDefaultSuffix(const QString & suffix);
extern void _ZN11QFileDialog16setDefaultSuffixERK7QString(void* qthis, void* arg0);
  // proto:  void QFileDialog::setItemDelegate(QAbstractItemDelegate * delegate);
extern void _ZN11QFileDialog15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0);
  // proto:  QList<QUrl> QFileDialog::sidebarUrls();
extern void _ZNK11QFileDialog11sidebarUrlsEv(void* qthis);
  // proto:  QString QFileDialog::defaultSuffix();
extern void _ZNK11QFileDialog13defaultSuffixEv(void* qthis);
  // proto:  void QFileDialog::setProxyModel(QAbstractProxyModel * model);
extern void _ZN11QFileDialog13setProxyModelEP19QAbstractProxyModel(void* qthis, void* arg0);
  // proto:  void QFileDialog::selectFile(const QString & filename);
extern void _ZN11QFileDialog10selectFileERK7QString(void* qthis, void* arg0);
  // proto:  bool QFileDialog::resolveSymlinks();
extern void _ZNK11QFileDialog15resolveSymlinksEv(void* qthis);
  // proto:  void QFileDialog::setDirectory(const QString & directory);
extern void _ZN11QFileDialog12setDirectoryERK7QString(void* qthis, void* arg0);
  // proto:  void QFileDialog::selectUrl(const QUrl & url);
extern void _ZN11QFileDialog9selectUrlERK4QUrl(void* qthis, void* arg0);
  // proto:  QString QFileDialog::selectedNameFilter();
extern void _ZNK11QFileDialog18selectedNameFilterEv(void* qthis);
  // proto:  QUrl QFileDialog::directoryUrl();
extern void _ZNK11QFileDialog12directoryUrlEv(void* qthis);
  // proto:  bool QFileDialog::isReadOnly();
extern void _ZNK11QFileDialog10isReadOnlyEv(void* qthis);
  // proto:  QByteArray QFileDialog::saveState();
extern void _ZNK11QFileDialog9saveStateEv(void* qthis);
  // proto:  void QFileDialog::open(QObject * receiver, const char * member);
extern void _ZN11QFileDialog4openEP7QObjectPKc(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  QDir QFileDialog::directory();
extern void _ZNK11QFileDialog9directoryEv(void* qthis);
  // proto:  void QFileDialog::setDirectory(const QDir & directory);
extern void demth_ZN11QFileDialog12setDirectoryERK4QDir(void* qthis, void* arg0);
  // proto:  void QFileDialog::setVisible(bool visible);
extern void _ZN11QFileDialog10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QFileDialog::setIconProvider(QFileIconProvider * provider);
extern void _ZN11QFileDialog15setIconProviderEP17QFileIconProvider(void* qthis, void* arg0);
  // proto:  void QFileDialog::selectMimeTypeFilter(const QString & filter);
extern void _ZN11QFileDialog20selectMimeTypeFilterERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QFileDialog::mimeTypeFilters();
extern void _ZNK11QFileDialog15mimeTypeFiltersEv(void* qthis);
  // proto:  void QFileDialog::setMimeTypeFilters(const QStringList & filters);
extern void _ZN11QFileDialog18setMimeTypeFiltersERK11QStringList(void* qthis, void* arg0);
  // proto:  void QFileDialog::setResolveSymlinks(bool enabled);
extern void _ZN11QFileDialog18setResolveSymlinksEb(void* qthis, bool arg0);
  // proto:  void QFileDialog::setReadOnly(bool enabled);
extern void _ZN11QFileDialog11setReadOnlyEb(void* qthis, bool arg0);
  // proto:  void QFileDialog::setNameFilterDetailsVisible(bool enabled);
extern void _ZN11QFileDialog27setNameFilterDetailsVisibleEb(void* qthis, bool arg0);
  // proto:  void QFileDialog::selectNameFilter(const QString & filter);
extern void _ZN11QFileDialog16selectNameFilterERK7QString(void* qthis, void* arg0);
  // proto:  bool QFileDialog::restoreState(const QByteArray & state);
extern void _ZN11QFileDialog12restoreStateERK10QByteArray(void* qthis, void* arg0);
  // proto:  QFileIconProvider * QFileDialog::iconProvider();
extern void _ZNK11QFileDialog12iconProviderEv(void* qthis);
  // proto:  QStringList QFileDialog::selectedFiles();
extern void _ZNK11QFileDialog13selectedFilesEv(void* qthis);
  // proto:  void QFileDialog::~QFileDialog();
extern void _ZN11QFileDialogD0Ev(void* qthis);
  // proto:  QAbstractItemDelegate * QFileDialog::itemDelegate();
extern void _ZNK11QFileDialog12itemDelegateEv(void* qthis);
  // proto:  bool QFileDialog::confirmOverwrite();
extern void _ZNK11QFileDialog16confirmOverwriteEv(void* qthis);
  // proto:  void QFileDialog::setHistory(const QStringList & paths);
extern void _ZN11QFileDialog10setHistoryERK11QStringList(void* qthis, void* arg0);
  // proto:  void QFileDialog::setNameFilter(const QString & filter);
extern void _ZN11QFileDialog13setNameFilterERK7QString(void* qthis, void* arg0);
  // proto:  QAbstractProxyModel * QFileDialog::proxyModel();
extern void _ZNK11QFileDialog10proxyModelEv(void* qthis);
  // proto:  void QFileDialog::setNameFilters(const QStringList & filters);
extern void _ZN11QFileDialog14setNameFiltersERK11QStringList(void* qthis, void* arg0);
  // proto:  QList<QUrl> QFileDialog::selectedUrls();
extern void _ZNK11QFileDialog12selectedUrlsEv(void* qthis);
  // proto:  QStringList QFileDialog::history();
extern void _ZNK11QFileDialog7historyEv(void* qthis);
  // proto:  bool QFileDialog::isNameFilterDetailsVisible();
extern void _ZNK11QFileDialog26isNameFilterDetailsVisibleEv(void* qthis);
  // proto:  void QFileDialog::QFileDialog(const QFileDialog & );
extern void* dector_ZN11QFileDialogC1ERKS_(void* arg0);
extern void _ZN11QFileDialogC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFileDialog::QFileDialog(QWidget * parent, const QString & caption, const QString & directory, const QString & filter);
extern void* dector_ZN11QFileDialogC1EP7QWidgetRK7QStringS4_S4_(void* arg0, void* arg1, void* arg2, void* arg3);
extern void _ZN11QFileDialogC1EP7QWidgetRK7QStringS4_S4_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
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

  // proto:  const QMetaObject * QFileDialog::metaObject();
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

  // proto:  void QFileDialog::setDirectoryUrl(const QUrl & directory);
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

  // proto:  QStringList QFileDialog::nameFilters();
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

  // proto:  void QFileDialog::setConfirmOverwrite(bool enabled);
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

  // proto:  void QFileDialog::setDefaultSuffix(const QString & suffix);
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

  // proto:  void QFileDialog::setItemDelegate(QAbstractItemDelegate * delegate);
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

  // proto:  QList<QUrl> QFileDialog::sidebarUrls();
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

  // proto:  QString QFileDialog::defaultSuffix();
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

  // proto:  void QFileDialog::setProxyModel(QAbstractProxyModel * model);
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

  // proto:  void QFileDialog::selectFile(const QString & filename);
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

  // proto:  bool QFileDialog::resolveSymlinks();
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

  // proto:  void QFileDialog::setDirectory(const QString & directory);
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
    // invoke: void setDirectory(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QFileDialog12setDirectoryERK7QString(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QFileDialog12setDirectoryERK4QDir
    // invoke: void setDirectory(const class QDir &)
    var arg0 = args[0].(QDir).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN11QFileDialog12setDirectoryERK4QDir(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDialog", "setDirectory", args)
  }

}

  // proto:  void QFileDialog::selectUrl(const QUrl & url);
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

  // proto:  QString QFileDialog::selectedNameFilter();
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

  // proto:  QUrl QFileDialog::directoryUrl();
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

  // proto:  bool QFileDialog::isReadOnly();
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

  // proto:  QByteArray QFileDialog::saveState();
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

  // proto:  void QFileDialog::open(QObject * receiver, const char * member);
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
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZN11QFileDialog4openEP7QObjectPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFileDialog", "open", args)
  }

}

  // proto:  QDir QFileDialog::directory();
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

  // proto:  void QFileDialog::setVisible(bool visible);
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

  // proto:  void QFileDialog::setIconProvider(QFileIconProvider * provider);
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

  // proto:  void QFileDialog::selectMimeTypeFilter(const QString & filter);
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

  // proto:  QStringList QFileDialog::mimeTypeFilters();
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

  // proto:  void QFileDialog::setMimeTypeFilters(const QStringList & filters);
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

  // proto:  void QFileDialog::setResolveSymlinks(bool enabled);
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

  // proto:  void QFileDialog::setReadOnly(bool enabled);
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

  // proto:  void QFileDialog::setNameFilterDetailsVisible(bool enabled);
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

  // proto:  void QFileDialog::selectNameFilter(const QString & filter);
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

  // proto:  bool QFileDialog::restoreState(const QByteArray & state);
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

  // proto:  QFileIconProvider * QFileDialog::iconProvider();
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

  // proto:  QStringList QFileDialog::selectedFiles();
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

  // proto:  void QFileDialog::~QFileDialog();
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

  // proto:  QAbstractItemDelegate * QFileDialog::itemDelegate();
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

  // proto:  bool QFileDialog::confirmOverwrite();
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

  // proto:  void QFileDialog::setHistory(const QStringList & paths);
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

  // proto:  void QFileDialog::setNameFilter(const QString & filter);
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

  // proto:  QAbstractProxyModel * QFileDialog::proxyModel();
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

  // proto:  void QFileDialog::setNameFilters(const QStringList & filters);
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

  // proto:  QList<QUrl> QFileDialog::selectedUrls();
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

  // proto:  QStringList QFileDialog::history();
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

  // proto:  bool QFileDialog::isNameFilterDetailsVisible();
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

  // proto:  void QFileDialog::QFileDialog(const QFileDialog & );
func NewQFileDialog(args ...interface{}) QFileDialog {
  return QFileDialog{}
}

// <= body block end

