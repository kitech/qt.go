package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qtextbrowser.h
// dst-file: /src/widgets/qtextbrowser.go
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
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyUrl", args)
 }

}


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


func NewQTextBrowser(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextBrowser", "setOpenLinks", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBrowser", "historyTitle", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBrowser", "setSearchPaths", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBrowser", "loadResource", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBrowser", "setOpenExternalLinks", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextBrowser", "setSource", args)
 }

}


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

