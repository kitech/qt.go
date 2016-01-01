package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qdockwidget.h
// dst-file: /src/widgets/qdockwidget.go
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
// class sizeof(QDockWidget)=1
type QDockWidget struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _featuresChanged QDockWidget_featuresChanged_signal;
//  _visibilityChanged QDockWidget_visibilityChanged_signal;
//  _topLevelChanged QDockWidget_topLevelChanged_signal;
//  _allowedAreasChanged QDockWidget_allowedAreasChanged_signal;
//  _dockLocationChanged QDockWidget_dockLocationChanged_signal;
}


func (this *QDockWidget) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget6widgetEv
  default:
    qtrt.ErrorResolve("QDockWidget", "widget", args)
  }

}


func (this *QDockWidget) setFloating(args ...interface{}) () {
  // setFloating(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidget11setFloatingEb
  default:
    qtrt.ErrorResolve("QDockWidget", "setFloating", args)
  }

}


func (this *QDockWidget) titleBarWidget(args ...interface{}) () {
  // titleBarWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget14titleBarWidgetEv
  default:
    qtrt.ErrorResolve("QDockWidget", "titleBarWidget", args)
  }

}


func (this *QDockWidget) FreeQDockWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDockWidget", "~QDockWidget", args)
  }

}


func (this *QDockWidget) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidget9setWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QDockWidget", "setWidget", args)
  }

}


func (this *QDockWidget) isFloating(args ...interface{}) () {
  // isFloating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget10isFloatingEv
  default:
    qtrt.ErrorResolve("QDockWidget", "isFloating", args)
  }

}


func (this *QDockWidget) toggleViewAction(args ...interface{}) () {
  // toggleViewAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget16toggleViewActionEv
  default:
    qtrt.ErrorResolve("QDockWidget", "toggleViewAction", args)
  }

}


func NewQDockWidget(args ...interface{}) QDockWidget {
  return QDockWidget{}
}


func (this *QDockWidget) setTitleBarWidget(args ...interface{}) () {
  // setTitleBarWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QDockWidget17setTitleBarWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QDockWidget", "setTitleBarWidget", args)
  }

}


func (this *QDockWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QDockWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QDockWidget", "metaObject", args)
  }

}

// <= body block end

