package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qscrollarea.h
// dst-file: /src/widgets/qscrollarea.go
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
// class sizeof(QScrollArea)=1
type QScrollArea struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQScrollArea(args ...interface{}) QScrollArea {
  return QScrollArea{}
}


func (this *QScrollArea) setWidgetResizable(args ...interface{}) () {
  // setWidgetResizable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea18setWidgetResizableEb
  default:
    qtrt.ErrorResolve("QScrollArea", "setWidgetResizable", args)
  }

}


func (this *QScrollArea) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea9setWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QScrollArea", "setWidget", args)
  }

}


func (this *QScrollArea) takeWidget(args ...interface{}) () {
  // takeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea10takeWidgetEv
  default:
    qtrt.ErrorResolve("QScrollArea", "takeWidget", args)
  }

}


func (this *QScrollArea) ensureVisible(args ...interface{}) () {
  // ensureVisible(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea13ensureVisibleEiiii
  default:
    qtrt.ErrorResolve("QScrollArea", "ensureVisible", args)
  }

}


func (this *QScrollArea) ensureWidgetVisible(args ...interface{}) () {
  // ensureWidgetVisible(class QWidget *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii
  default:
    qtrt.ErrorResolve("QScrollArea", "ensureWidgetVisible", args)
  }

}


func (this *QScrollArea) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea6widgetEv
  default:
    qtrt.ErrorResolve("QScrollArea", "widget", args)
  }

}


func (this *QScrollArea) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea8sizeHintEv
  default:
    qtrt.ErrorResolve("QScrollArea", "sizeHint", args)
  }

}


func (this *QScrollArea) widgetResizable(args ...interface{}) () {
  // widgetResizable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea15widgetResizableEv
  default:
    qtrt.ErrorResolve("QScrollArea", "widgetResizable", args)
  }

}


func (this *QScrollArea) FreeQScrollArea(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollArea", "~QScrollArea", args)
  }

}


func (this *QScrollArea) focusNextPrevChild(args ...interface{}) () {
  // focusNextPrevChild(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea18focusNextPrevChildEb
  default:
    qtrt.ErrorResolve("QScrollArea", "focusNextPrevChild", args)
  }

}


func (this *QScrollArea) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea10metaObjectEv
  default:
    qtrt.ErrorResolve("QScrollArea", "metaObject", args)
  }

}

// <= body block end

