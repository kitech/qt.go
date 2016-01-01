package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qitemeditorfactory.h
// dst-file: /src/widgets/qitemeditorfactory.go
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
// class sizeof(QItemEditorCreatorBase)=8
type QItemEditorCreatorBase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QItemEditorFactory)=1
type QItemEditorFactory struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QItemEditorCreatorBase) valuePropertyName(args ...interface{}) () {
  // valuePropertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QItemEditorCreatorBase17valuePropertyNameEv
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "valuePropertyName", args)
  }

}


func (this *QItemEditorCreatorBase) createWidget(args ...interface{}) () {
  // createWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK22QItemEditorCreatorBase12createWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "createWidget", args)
  }

}


func (this *QItemEditorCreatorBase) FreeQItemEditorCreatorBase(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorCreatorBase", "~QItemEditorCreatorBase", args)
  }

}


func NewQItemEditorFactory(args ...interface{}) QItemEditorFactory {
  return QItemEditorFactory{}
}


func (this *QItemEditorFactory) valuePropertyName(args ...interface{}) () {
  // valuePropertyName(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QItemEditorFactory17valuePropertyNameEi
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "valuePropertyName", args)
  }

}


func (this *QItemEditorFactory) defaultFactory_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "defaultFactory", args)
  }

}


func (this *QItemEditorFactory) FreeQItemEditorFactory(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "~QItemEditorFactory", args)
  }

}


func (this *QItemEditorFactory) registerEditor(args ...interface{}) () {
  // registerEditor(int, class QItemEditorCreatorBase *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QItemEditorCreatorBase{}) // "QItemEditorCreatorBase *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QItemEditorFactory14registerEditorEiP22QItemEditorCreatorBase
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "registerEditor", args)
  }

}


func (this *QItemEditorFactory) setDefaultFactory_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "setDefaultFactory", args)
  }

}


func (this *QItemEditorFactory) createEditor(args ...interface{}) () {
  // createEditor(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QItemEditorFactory12createEditorEiP7QWidget
  default:
    qtrt.ErrorResolve("QItemEditorFactory", "createEditor", args)
  }

}

// <= body block end

