package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qmdisubwindow.h
// dst-file: /src/widgets/qmdisubwindow.go
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
// class sizeof(QMdiSubWindow)=1
type QMdiSubWindow struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToActivate QMdiSubWindow_aboutToActivate_signal;
//  _windowStateChanged QMdiSubWindow_windowStateChanged_signal;
}


func (this *QMdiSubWindow) keyboardSingleStep(args ...interface{}) () {
  // keyboardSingleStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow18keyboardSingleStepEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "keyboardSingleStep", args)
  }

}


func (this *QMdiSubWindow) keyboardPageStep(args ...interface{}) () {
  // keyboardPageStep()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow16keyboardPageStepEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "keyboardPageStep", args)
  }

}


func (this *QMdiSubWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow10metaObjectEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "metaObject", args)
  }

}


func (this *QMdiSubWindow) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow8sizeHintEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "sizeHint", args)
  }

}


func (this *QMdiSubWindow) maximizedSystemMenuIconWidget(args ...interface{}) () {
  // maximizedSystemMenuIconWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow29maximizedSystemMenuIconWidgetEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "maximizedSystemMenuIconWidget", args)
  }

}


func (this *QMdiSubWindow) setSystemMenu(args ...interface{}) () {
  // setSystemMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow13setSystemMenuEP5QMenu
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setSystemMenu", args)
  }

}


func (this *QMdiSubWindow) FreeQMdiSubWindow(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "~QMdiSubWindow", args)
  }

}


func (this *QMdiSubWindow) setKeyboardSingleStep(args ...interface{}) () {
  // setKeyboardSingleStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow21setKeyboardSingleStepEi
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setKeyboardSingleStep", args)
  }

}


func (this *QMdiSubWindow) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow6widgetEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "widget", args)
  }

}


func (this *QMdiSubWindow) showShaded(args ...interface{}) () {
  // showShaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow10showShadedEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "showShaded", args)
  }

}


func (this *QMdiSubWindow) maximizedButtonsWidget(args ...interface{}) () {
  // maximizedButtonsWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow22maximizedButtonsWidgetEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "maximizedButtonsWidget", args)
  }

}


func (this *QMdiSubWindow) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "minimumSizeHint", args)
  }

}


func (this *QMdiSubWindow) showSystemMenu(args ...interface{}) () {
  // showSystemMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow14showSystemMenuEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "showSystemMenu", args)
  }

}


func (this *QMdiSubWindow) systemMenu(args ...interface{}) () {
  // systemMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow10systemMenuEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "systemMenu", args)
  }

}


func NewQMdiSubWindow(args ...interface{}) QMdiSubWindow {
  return QMdiSubWindow{}
}


func (this *QMdiSubWindow) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow9setWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setWidget", args)
  }

}


func (this *QMdiSubWindow) isShaded(args ...interface{}) () {
  // isShaded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow8isShadedEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "isShaded", args)
  }

}


func (this *QMdiSubWindow) mdiArea(args ...interface{}) () {
  // mdiArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMdiSubWindow7mdiAreaEv
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "mdiArea", args)
  }

}


func (this *QMdiSubWindow) setKeyboardPageStep(args ...interface{}) () {
  // setKeyboardPageStep(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QMdiSubWindow19setKeyboardPageStepEi
  default:
    qtrt.ErrorResolve("QMdiSubWindow", "setKeyboardPageStep", args)
  }

}

// <= body block end

