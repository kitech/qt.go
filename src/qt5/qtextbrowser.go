package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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
extern void C_ZNK12QTextBrowser9openLinksEv(void* qthis); // 4
  // proto:  bool QTextBrowser::isBackwardAvailable();
extern void C_ZNK12QTextBrowser19isBackwardAvailableEv(void* qthis); // 4
  // proto:  void QTextBrowser::home();
extern void C_ZN12QTextBrowser4homeEv(void* qthis); // 4
  // proto:  void QTextBrowser::setSource(const QUrl & name);
extern void C_ZN12QTextBrowser9setSourceERK4QUrl(void* qthis, void* arg0); // 4
  // proto:  bool QTextBrowser::openExternalLinks();
extern void C_ZNK12QTextBrowser17openExternalLinksEv(void* qthis); // 4
  // proto:  QStringList QTextBrowser::searchPaths();
extern void C_ZNK12QTextBrowser11searchPathsEv(void* qthis); // 4
  // proto:  void QTextBrowser::setOpenLinks(bool open);
extern void C_ZN12QTextBrowser12setOpenLinksEb(void* qthis, bool arg0); // 4
  // proto:  QUrl QTextBrowser::historyUrl(int );
extern void C_ZNK12QTextBrowser10historyUrlEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTextBrowser::isForwardAvailable();
extern void C_ZNK12QTextBrowser18isForwardAvailableEv(void* qthis); // 4
  // proto:  QUrl QTextBrowser::source();
extern void C_ZNK12QTextBrowser6sourceEv(void* qthis); // 4
  // proto:  void QTextBrowser::forward();
extern void C_ZN12QTextBrowser7forwardEv(void* qthis); // 4
  // proto:  void QTextBrowser::setSearchPaths(const QStringList & paths);
extern void C_ZN12QTextBrowser14setSearchPathsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QTextBrowser::QTextBrowser(QWidget * parent);
extern void* C_ZN12QTextBrowserC2EP7QWidget(void* arg0); // 3
  // proto:  int QTextBrowser::forwardHistoryCount();
extern void C_ZNK12QTextBrowser19forwardHistoryCountEv(void* qthis); // 4
  // proto:  QVariant QTextBrowser::loadResource(int type, const QUrl & name);
extern void C_ZN12QTextBrowser12loadResourceEiRK4QUrl(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTextBrowser::~QTextBrowser();
extern void C_ZN12QTextBrowserD2Ev(void* qthis); // 4
  // proto:  int QTextBrowser::backwardHistoryCount();
extern void C_ZNK12QTextBrowser20backwardHistoryCountEv(void* qthis); // 4
  // proto:  void QTextBrowser::setOpenExternalLinks(bool open);
extern void C_ZN12QTextBrowser20setOpenExternalLinksEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QTextBrowser::metaObject();
extern void C_ZNK12QTextBrowser10metaObjectEv(void* qthis); // 4
  // proto:  QString QTextBrowser::historyTitle(int );
extern void C_ZNK12QTextBrowser12historyTitleEi(void* qthis, int32_t arg0); // 4
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
func (this *QTextBrowser) openLinks(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser9openLinksEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "openLinks", args)
  }

}

// isBackwardAvailable()
func (this *QTextBrowser) isBackwardAvailable(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser19isBackwardAvailableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "isBackwardAvailable", args)
  }

}

// home()
func (this *QTextBrowser) home(args ...interface{}) () {
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

}

// setSource(const class QUrl &)
func (this *QTextBrowser) setSource(args ...interface{}) () {
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

}

// openExternalLinks()
func (this *QTextBrowser) openExternalLinks(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser17openExternalLinksEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "openExternalLinks", args)
  }

}

// searchPaths()
func (this *QTextBrowser) searchPaths(args ...interface{}) () {
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

}

// setOpenLinks(_Bool)
func (this *QTextBrowser) setOpenLinks(args ...interface{}) () {
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

}

// historyUrl(int)
func (this *QTextBrowser) historyUrl(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser10historyUrlEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyUrl", args)
  }

}

// isForwardAvailable()
func (this *QTextBrowser) isForwardAvailable(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser18isForwardAvailableEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "isForwardAvailable", args)
  }

}

// source()
func (this *QTextBrowser) source(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser6sourceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "source", args)
  }

}

// forward()
func (this *QTextBrowser) forward(args ...interface{}) () {
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

}

// setSearchPaths(const class QStringList &)
func (this *QTextBrowser) setSearchPaths(args ...interface{}) () {
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
func (this *QTextBrowser) forwardHistoryCount(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser19forwardHistoryCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "forwardHistoryCount", args)
  }

}

// loadResource(int, const class QUrl &)
func (this *QTextBrowser) loadResource(args ...interface{}) () {
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
    var ret = C.C_ZN12QTextBrowser12loadResourceEiRK4QUrl(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "loadResource", args)
  }

}

// ~QTextBrowser()
func (this *QTextBrowser) FreeQTextBrowser(args ...interface{}) () {
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

}

// backwardHistoryCount()
func (this *QTextBrowser) backwardHistoryCount(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser20backwardHistoryCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "backwardHistoryCount", args)
  }

}

// setOpenExternalLinks(_Bool)
func (this *QTextBrowser) setOpenExternalLinks(args ...interface{}) () {
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

}

// metaObject()
func (this *QTextBrowser) metaObject(args ...interface{}) () {
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

}

// historyTitle(int)
func (this *QTextBrowser) historyTitle(args ...interface{}) () {
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
    var ret = C.C_ZNK12QTextBrowser12historyTitleEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyTitle", args)
  }

}

// reload()
func (this *QTextBrowser) reload(args ...interface{}) () {
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

}

// clearHistory()
func (this *QTextBrowser) clearHistory(args ...interface{}) () {
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

}

// backward()
func (this *QTextBrowser) backward(args ...interface{}) () {
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

}

// <= body block end

