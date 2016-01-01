package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qmdiarea.h
// dst-file: /src/widgets/qmdiarea.go
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
// class sizeof(QMdiArea)=1
type QMdiArea struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
//  _subWindowActivated QMdiArea_subWindowActivated_signal;
}


func (this *QMdiArea) activateNextSubWindow(args ...interface{}) () {
  // activateNextSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea21activateNextSubWindowEv
  default:
    qtrt.ErrorResolve("QMdiArea", "activateNextSubWindow", args)
 }

}


func (this *QMdiArea) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea13setBackgroundERK6QBrush
  default:
    qtrt.ErrorResolve("QMdiArea", "setBackground", args)
 }

}


func (this *QMdiArea) FreeQMdiArea(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QMdiArea", "~QMdiArea", args)
 }

}


func (this *QMdiArea) removeSubWindow(args ...interface{}) () {
  // removeSubWindow(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea15removeSubWindowEP7QWidget
  default:
    qtrt.ErrorResolve("QMdiArea", "removeSubWindow", args)
 }

}


func (this *QMdiArea) setTabsClosable(args ...interface{}) () {
  // setTabsClosable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea15setTabsClosableEb
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsClosable", args)
 }

}


func (this *QMdiArea) currentSubWindow(args ...interface{}) () {
  // currentSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea16currentSubWindowEv
  default:
    qtrt.ErrorResolve("QMdiArea", "currentSubWindow", args)
 }

}


func (this *QMdiArea) tabsMovable(args ...interface{}) () {
  // tabsMovable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea11tabsMovableEv
  default:
    qtrt.ErrorResolve("QMdiArea", "tabsMovable", args)
 }

}


func (this *QMdiArea) activatePreviousSubWindow(args ...interface{}) () {
  // activatePreviousSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea25activatePreviousSubWindowEv
  default:
    qtrt.ErrorResolve("QMdiArea", "activatePreviousSubWindow", args)
 }

}


func (this *QMdiArea) setDocumentMode(args ...interface{}) () {
  // setDocumentMode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea15setDocumentModeEb
  default:
    qtrt.ErrorResolve("QMdiArea", "setDocumentMode", args)
 }

}


func (this *QMdiArea) documentMode(args ...interface{}) () {
  // documentMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea12documentModeEv
  default:
    qtrt.ErrorResolve("QMdiArea", "documentMode", args)
 }

}


func (this *QMdiArea) setActiveSubWindow(args ...interface{}) () {
  // setActiveSubWindow(class QMdiSubWindow *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMdiSubWindow{}) // "QMdiSubWindow *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea18setActiveSubWindowEP13QMdiSubWindow
  default:
    qtrt.ErrorResolve("QMdiArea", "setActiveSubWindow", args)
 }

}


func (this *QMdiArea) activeSubWindow(args ...interface{}) () {
  // activeSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea15activeSubWindowEv
  default:
    qtrt.ErrorResolve("QMdiArea", "activeSubWindow", args)
 }

}


func (this *QMdiArea) setTabsMovable(args ...interface{}) () {
  // setTabsMovable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea14setTabsMovableEb
  default:
    qtrt.ErrorResolve("QMdiArea", "setTabsMovable", args)
 }

}


func (this *QMdiArea) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea10metaObjectEv
  default:
    qtrt.ErrorResolve("QMdiArea", "metaObject", args)
 }

}


func NewQMdiArea(args ...interface{})() {
}


func (this *QMdiArea) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea8sizeHintEv
  default:
    qtrt.ErrorResolve("QMdiArea", "sizeHint", args)
 }

}


func (this *QMdiArea) closeAllSubWindows(args ...interface{}) () {
  // closeAllSubWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea18closeAllSubWindowsEv
  default:
    qtrt.ErrorResolve("QMdiArea", "closeAllSubWindows", args)
 }

}


func (this *QMdiArea) cascadeSubWindows(args ...interface{}) () {
  // cascadeSubWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea17cascadeSubWindowsEv
  default:
    qtrt.ErrorResolve("QMdiArea", "cascadeSubWindows", args)
 }

}


func (this *QMdiArea) closeActiveSubWindow(args ...interface{}) () {
  // closeActiveSubWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea20closeActiveSubWindowEv
  default:
    qtrt.ErrorResolve("QMdiArea", "closeActiveSubWindow", args)
 }

}


func (this *QMdiArea) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea10backgroundEv
  default:
    qtrt.ErrorResolve("QMdiArea", "background", args)
 }

}


func (this *QMdiArea) tileSubWindows(args ...interface{}) () {
  // tileSubWindows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN8QMdiArea14tileSubWindowsEv
  default:
    qtrt.ErrorResolve("QMdiArea", "tileSubWindows", args)
 }

}


func (this *QMdiArea) tabsClosable(args ...interface{}) () {
  // tabsClosable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea12tabsClosableEv
  default:
    qtrt.ErrorResolve("QMdiArea", "tabsClosable", args)
 }

}


func (this *QMdiArea) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK8QMdiArea15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QMdiArea", "minimumSizeHint", args)
 }

}

// <= body block end

