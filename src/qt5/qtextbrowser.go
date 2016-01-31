package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qtextbrowser.h
// dst-file: /src/widgets/qtextbrowser.go
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
  // proto:  bool QTextBrowser::openLinks();
extern bool C_ZNK12QTextBrowser9openLinksEv(void* qthis); // 4
  // proto:  bool QTextBrowser::isBackwardAvailable();
extern bool C_ZNK12QTextBrowser19isBackwardAvailableEv(void* qthis); // 4
  // proto:  void QTextBrowser::home();
extern void C_ZN12QTextBrowser4homeEv(void* qthis); // 4
  // proto:  void QTextBrowser::setSource(const QUrl & name);
extern void C_ZN12QTextBrowser9setSourceERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  bool QTextBrowser::openExternalLinks();
extern bool C_ZNK12QTextBrowser17openExternalLinksEv(void* qthis); // 4
  // proto:  QStringList QTextBrowser::searchPaths();
extern void C_ZNK12QTextBrowser11searchPathsEv(void* qthis); // 4
  // proto:  void QTextBrowser::setOpenLinks(bool open);
extern void C_ZN12QTextBrowser12setOpenLinksEb(void* qthis, bool arg0); // 4
  // proto:  QUrl QTextBrowser::historyUrl(int );
extern void* C_ZNK12QTextBrowser10historyUrlEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextBrowser::isForwardAvailable();
extern bool C_ZNK12QTextBrowser18isForwardAvailableEv(void* qthis); // 4
  // proto:  QUrl QTextBrowser::source();
extern void* C_ZNK12QTextBrowser6sourceEv(void* qthis); // 4
  // proto:  void QTextBrowser::forward();
extern void C_ZN12QTextBrowser7forwardEv(void* qthis); // 4
  // proto:  void QTextBrowser::setSearchPaths(const QStringList & paths);
extern void C_ZN12QTextBrowser14setSearchPathsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QTextBrowser::QTextBrowser(QWidget * parent);
extern void* C_ZN12QTextBrowserC2EP7QWidget(void* arg0); // 3
  // proto:  int QTextBrowser::forwardHistoryCount();
extern int32_t C_ZNK12QTextBrowser19forwardHistoryCountEv(void* qthis); // 4
  // proto:  QVariant QTextBrowser::loadResource(int type, const QUrl & name);
extern void* C_ZN12QTextBrowser12loadResourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTextBrowser::~QTextBrowser();
extern void C_ZN12QTextBrowserD2Ev(void* qthis); // 4
  // proto:  int QTextBrowser::backwardHistoryCount();
extern int32_t C_ZNK12QTextBrowser20backwardHistoryCountEv(void* qthis); // 4
  // proto:  void QTextBrowser::setOpenExternalLinks(bool open);
extern void C_ZN12QTextBrowser20setOpenExternalLinksEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTextBrowser::metaObject();
extern void C_ZNK12QTextBrowser10metaObjectEv(void* qthis); // 4
  // proto:  QString QTextBrowser::historyTitle(int );
extern void* C_ZNK12QTextBrowser12historyTitleEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextBrowser::reload();
extern void C_ZN12QTextBrowser6reloadEv(void* qthis); // 4
  // proto:  void QTextBrowser::clearHistory();
extern void C_ZN12QTextBrowser12clearHistoryEv(void* qthis); // 4
  // proto:  void QTextBrowser::backward();
extern void C_ZN12QTextBrowser8backwardEv(void* qthis); // 4
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

// class sizeof(QTextBrowser)=1
type QTextBrowser struct {
  /*qbase*/ QTextEdit;
  qclsinst unsafe.Pointer /* *C.void */;
//  _forwardAvailable QTextBrowser_forwardAvailable_signal;
//  _sourceChanged QTextBrowser_sourceChanged_signal;
//  _highlighted QTextBrowser_highlighted_signal;
//  _anchorClicked QTextBrowser_anchorClicked_signal;
//  _historyChanged QTextBrowser_historyChanged_signal;
//  _backwardAvailable QTextBrowser_backwardAvailable_signal;
}

// openLinks()
func (this *QTextBrowser) Openlinks(args ...interface{}) (ret interface{}) {
  // openLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser9openLinksEv
    // invoke: bool openLinks()
    var ret0 = C.C_ZNK12QTextBrowser9openLinksEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "openLinks", args)
  }

  return
}

// isBackwardAvailable()
func (this *QTextBrowser) Isbackwardavailable(args ...interface{}) (ret interface{}) {
  // isBackwardAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser19isBackwardAvailableEv
    // invoke: bool isBackwardAvailable()
    var ret0 = C.C_ZNK12QTextBrowser19isBackwardAvailableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "isBackwardAvailable", args)
  }

  return
}

// home()
func (this *QTextBrowser) Home(args ...interface{}) () {
  // home()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser4homeEv
    // invoke: void home()
    C.C_ZN12QTextBrowser4homeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "home", args)
  }

  return
}

// setSource(const class QUrl &)
func (this *QTextBrowser) Setsource(args ...interface{}) () {
  // setSource(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser9setSourceERK4QUrl
    // invoke: void setSource(const class QUrl &)
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTextBrowser9setSourceERK4QUrl(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBrowser", "setSource", args)
  }

  return
}

// openExternalLinks()
func (this *QTextBrowser) Openexternallinks(args ...interface{}) (ret interface{}) {
  // openExternalLinks()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser17openExternalLinksEv
    // invoke: bool openExternalLinks()
    var ret0 = C.C_ZNK12QTextBrowser17openExternalLinksEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "openExternalLinks", args)
  }

  return
}

// searchPaths()
func (this *QTextBrowser) Searchpaths(args ...interface{}) () {
  // searchPaths()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser11searchPathsEv
    // invoke: QStringList searchPaths()
    C.C_ZNK12QTextBrowser11searchPathsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "searchPaths", args)
  }

  return
}

// setOpenLinks(_Bool)
func (this *QTextBrowser) Setopenlinks(args ...interface{}) () {
  // setOpenLinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser12setOpenLinksEb
    // invoke: void setOpenLinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTextBrowser12setOpenLinksEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBrowser", "setOpenLinks", args)
  }

  return
}

// historyUrl(int)
func (this *QTextBrowser) Historyurl(args ...interface{}) (ret interface{}) {
  // historyUrl(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser10historyUrlEi
    // invoke: QUrl historyUrl(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTextBrowser10historyUrlEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUrl{}) // "QUrl"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyUrl", args)
  }

  return
}

// isForwardAvailable()
func (this *QTextBrowser) Isforwardavailable(args ...interface{}) (ret interface{}) {
  // isForwardAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser18isForwardAvailableEv
    // invoke: bool isForwardAvailable()
    var ret0 = C.C_ZNK12QTextBrowser18isForwardAvailableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "isForwardAvailable", args)
  }

  return
}

// source()
func (this *QTextBrowser) Source(args ...interface{}) (ret interface{}) {
  // source()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser6sourceEv
    // invoke: QUrl source()
    var ret0 = C.C_ZNK12QTextBrowser6sourceEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QUrl{}) // "QUrl"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "source", args)
  }

  return
}

// forward()
func (this *QTextBrowser) Forward(args ...interface{}) () {
  // forward()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser7forwardEv
    // invoke: void forward()
    C.C_ZN12QTextBrowser7forwardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "forward", args)
  }

  return
}

// setSearchPaths(const class QStringList &)
func (this *QTextBrowser) Setsearchpaths(args ...interface{}) () {
  // setSearchPaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser14setSearchPathsERK11QStringList
    // invoke: void setSearchPaths(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTextBrowser14setSearchPathsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBrowser", "setSearchPaths", args)
  }

  return
}

// QTextBrowser(class QWidget *)
func NewQTextBrowser(args ...interface{}) *QTextBrowser {
  // QTextBrowser(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowserC1EP7QWidget
    // invoke: void QTextBrowser(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTextBrowserC2EP7QWidget(arg0)
    return &QTextBrowser{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextBrowser", "QTextBrowser", args)
  }

  return nil // QTextBrowser{qclsinst:qthis}
}

// forwardHistoryCount()
func (this *QTextBrowser) Forwardhistorycount(args ...interface{}) (ret interface{}) {
  // forwardHistoryCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser19forwardHistoryCountEv
    // invoke: int forwardHistoryCount()
    var ret0 = C.C_ZNK12QTextBrowser19forwardHistoryCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "forwardHistoryCount", args)
  }

  return
}

// loadResource(int, const class QUrl &)
func (this *QTextBrowser) Loadresource(args ...interface{}) (ret interface{}) {
  // loadResource(int, const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser12loadResourceEiRK4QUrl
    // invoke: QVariant loadResource(int, const class QUrl &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QTextBrowser12loadResourceEiRK4QUrl(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "loadResource", args)
  }

  return
}

// ~QTextBrowser()
func (this *QTextBrowser) Freeqtextbrowser(args ...interface{}) () {
  // ~QTextBrowser()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowserD0Ev
    // invoke: void ~QTextBrowser()
    C.C_ZN12QTextBrowserD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "~QTextBrowser", args)
  }

  return
}

// backwardHistoryCount()
func (this *QTextBrowser) Backwardhistorycount(args ...interface{}) (ret interface{}) {
  // backwardHistoryCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser20backwardHistoryCountEv
    // invoke: int backwardHistoryCount()
    var ret0 = C.C_ZNK12QTextBrowser20backwardHistoryCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "backwardHistoryCount", args)
  }

  return
}

// setOpenExternalLinks(_Bool)
func (this *QTextBrowser) Setopenexternallinks(args ...interface{}) () {
  // setOpenExternalLinks(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser20setOpenExternalLinksEb
    // invoke: void setOpenExternalLinks(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTextBrowser20setOpenExternalLinksEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextBrowser", "setOpenExternalLinks", args)
  }

  return
}

// metaObject()
func (this *QTextBrowser) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QTextBrowser10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "metaObject", args)
  }

  return
}

// historyTitle(int)
func (this *QTextBrowser) Historytitle(args ...interface{}) (ret interface{}) {
  // historyTitle(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextBrowser12historyTitleEi
    // invoke: QString historyTitle(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTextBrowser12historyTitleEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyTitle", args)
  }

  return
}

// reload()
func (this *QTextBrowser) Reload(args ...interface{}) () {
  // reload()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser6reloadEv
    // invoke: void reload()
    C.C_ZN12QTextBrowser6reloadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "reload", args)
  }

  return
}

// clearHistory()
func (this *QTextBrowser) Clearhistory(args ...interface{}) () {
  // clearHistory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser12clearHistoryEv
    // invoke: void clearHistory()
    C.C_ZN12QTextBrowser12clearHistoryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "clearHistory", args)
  }

  return
}

// backward()
func (this *QTextBrowser) Backward(args ...interface{}) () {
  // backward()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextBrowser8backwardEv
    // invoke: void backward()
    C.C_ZN12QTextBrowser8backwardEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextBrowser", "backward", args)
  }

  return
}

// <= body block end

