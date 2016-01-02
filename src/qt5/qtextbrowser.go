package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  bool QTextBrowser::isBackwardAvailable();
extern void _ZNK12QTextBrowser19isBackwardAvailableEv(void* qthis);
  // proto:  void QTextBrowser::reload();
extern void _ZN12QTextBrowser6reloadEv(void* qthis);
  // proto:  bool QTextBrowser::openLinks();
extern void _ZNK12QTextBrowser9openLinksEv(void* qthis);
  // proto:  void QTextBrowser::clearHistory();
extern void _ZN12QTextBrowser12clearHistoryEv(void* qthis);
  // proto:  const QMetaObject * QTextBrowser::metaObject();
extern void _ZNK12QTextBrowser10metaObjectEv(void* qthis);
  // proto:  QUrl QTextBrowser::historyUrl(int );
extern void _ZNK12QTextBrowser10historyUrlEi(void* qthis, int arg0);
  // proto:  bool QTextBrowser::isForwardAvailable();
extern void _ZNK12QTextBrowser18isForwardAvailableEv(void* qthis);
  // proto:  bool QTextBrowser::openExternalLinks();
extern void _ZNK12QTextBrowser17openExternalLinksEv(void* qthis);
  // proto:  void QTextBrowser::QTextBrowser(QWidget * parent);
extern void* dector_ZN12QTextBrowserC1EP7QWidget(void* arg0);
extern void _ZN12QTextBrowserC1EP7QWidget(void* qthis, void* arg0);
  // proto:  int QTextBrowser::backwardHistoryCount();
extern void _ZNK12QTextBrowser20backwardHistoryCountEv(void* qthis);
  // proto:  void QTextBrowser::home();
extern void _ZN12QTextBrowser4homeEv(void* qthis);
  // proto:  void QTextBrowser::~QTextBrowser();
extern void _ZN12QTextBrowserD0Ev(void* qthis);
  // proto:  void QTextBrowser::QTextBrowser(const QTextBrowser & );
extern void* dector_ZN12QTextBrowserC1ERKS_(void* arg0);
extern void _ZN12QTextBrowserC1ERKS_(void* qthis, void* arg0);
  // proto:  void QTextBrowser::setOpenLinks(bool open);
extern void _ZN12QTextBrowser12setOpenLinksEb(void* qthis, bool arg0);
  // proto:  void QTextBrowser::forward();
extern void _ZN12QTextBrowser7forwardEv(void* qthis);
  // proto:  QString QTextBrowser::historyTitle(int );
extern void _ZNK12QTextBrowser12historyTitleEi(void* qthis, int arg0);
  // proto:  void QTextBrowser::setSearchPaths(const QStringList & paths);
extern void _ZN12QTextBrowser14setSearchPathsERK11QStringList(void* qthis, void* arg0);
  // proto:  QVariant QTextBrowser::loadResource(int type, const QUrl & name);
extern void _ZN12QTextBrowser12loadResourceEiRK4QUrl(void* qthis, int arg0, void* arg1);
  // proto:  QUrl QTextBrowser::source();
extern void _ZNK12QTextBrowser6sourceEv(void* qthis);
  // proto:  void QTextBrowser::setOpenExternalLinks(bool open);
extern void _ZN12QTextBrowser20setOpenExternalLinksEb(void* qthis, bool arg0);
  // proto:  void QTextBrowser::setSource(const QUrl & name);
extern void _ZN12QTextBrowser9setSourceERK4QUrl(void* qthis, void* arg0);
  // proto:  QStringList QTextBrowser::searchPaths();
extern void _ZNK12QTextBrowser11searchPathsEv(void* qthis);
  // proto:  void QTextBrowser::backward();
extern void _ZN12QTextBrowser8backwardEv(void* qthis);
  // proto:  int QTextBrowser::forwardHistoryCount();
extern void _ZNK12QTextBrowser19forwardHistoryCountEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
//  _forwardAvailable QTextBrowser_forwardAvailable_signal;
//  _sourceChanged QTextBrowser_sourceChanged_signal;
//  _highlighted QTextBrowser_highlighted_signal;
//  _anchorClicked QTextBrowser_anchorClicked_signal;
//  _historyChanged QTextBrowser_historyChanged_signal;
//  _backwardAvailable QTextBrowser_backwardAvailable_signal;
}

  // proto:  bool QTextBrowser::isBackwardAvailable();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "isBackwardAvailable", args)
  }

}

  // proto:  void QTextBrowser::reload();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "reload", args)
  }

}

  // proto:  bool QTextBrowser::openLinks();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "openLinks", args)
  }

}

  // proto:  void QTextBrowser::clearHistory();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "clearHistory", args)
  }

}

  // proto:  const QMetaObject * QTextBrowser::metaObject();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "metaObject", args)
  }

}

  // proto:  QUrl QTextBrowser::historyUrl(int );
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyUrl", args)
  }

}

  // proto:  bool QTextBrowser::isForwardAvailable();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "isForwardAvailable", args)
  }

}

  // proto:  bool QTextBrowser::openExternalLinks();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "openExternalLinks", args)
  }

}

  // proto:  void QTextBrowser::QTextBrowser(QWidget * parent);
func NewQTextBrowser(args ...interface{}) QTextBrowser {
  return QTextBrowser{}
}

  // proto:  int QTextBrowser::backwardHistoryCount();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "backwardHistoryCount", args)
  }

}

  // proto:  void QTextBrowser::home();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "home", args)
  }

}

  // proto:  void QTextBrowser::~QTextBrowser();
func (this *QTextBrowser) FreeQTextBrowser(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTextBrowser", "~QTextBrowser", args)
  }

}

  // proto:  void QTextBrowser::setOpenLinks(bool open);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "setOpenLinks", args)
  }

}

  // proto:  void QTextBrowser::forward();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "forward", args)
  }

}

  // proto:  QString QTextBrowser::historyTitle(int );
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyTitle", args)
  }

}

  // proto:  void QTextBrowser::setSearchPaths(const QStringList & paths);
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
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "setSearchPaths", args)
  }

}

  // proto:  QVariant QTextBrowser::loadResource(int type, const QUrl & name);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QUrl).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "loadResource", args)
  }

}

  // proto:  QUrl QTextBrowser::source();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "source", args)
  }

}

  // proto:  void QTextBrowser::setOpenExternalLinks(bool open);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "setOpenExternalLinks", args)
  }

}

  // proto:  void QTextBrowser::setSource(const QUrl & name);
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
    var arg0 = args[0].(QUrl).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QTextBrowser", "setSource", args)
  }

}

  // proto:  QStringList QTextBrowser::searchPaths();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "searchPaths", args)
  }

}

  // proto:  void QTextBrowser::backward();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "backward", args)
  }

}

  // proto:  int QTextBrowser::forwardHistoryCount();
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "forwardHistoryCount", args)
  }

}

// <= body block end

