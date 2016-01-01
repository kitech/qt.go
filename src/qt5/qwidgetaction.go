package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qwidgetaction.h
// dst-file: /src/widgets/qwidgetaction.go
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
// class sizeof(QWidgetAction)=1
type QWidgetAction struct {
  /*qbase*/ QAction;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QWidgetAction) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetAction10metaObjectEv
  default:
    qtrt.ErrorResolve("QWidgetAction", "metaObject", args)
  }

}


func (this *QWidgetAction) FreeQWidgetAction(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidgetAction", "~QWidgetAction", args)
  }

}


func (this *QWidgetAction) setDefaultWidget(args ...interface{}) () {
  // setDefaultWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction16setDefaultWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QWidgetAction", "setDefaultWidget", args)
  }

}


func (this *QWidgetAction) releaseWidget(args ...interface{}) () {
  // releaseWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction13releaseWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QWidgetAction", "releaseWidget", args)
  }

}


func NewQWidgetAction(args ...interface{}) QWidgetAction {
  return QWidgetAction{}
}


func (this *QWidgetAction) requestWidget(args ...interface{}) () {
  // requestWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction13requestWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QWidgetAction", "requestWidget", args)
  }

}


func (this *QWidgetAction) defaultWidget(args ...interface{}) () {
  // defaultWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetAction13defaultWidgetEv
  default:
    qtrt.ErrorResolve("QWidgetAction", "defaultWidget", args)
  }

}

// <= body block end

