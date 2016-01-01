package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qabstractscrollarea.h
// dst-file: /src/widgets/qabstractscrollarea.go
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
// class sizeof(QAbstractScrollArea)=1
type QAbstractScrollArea struct {
  /*qbase*/ QFrame;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QAbstractScrollArea) horizontalScrollBar(args ...interface{}) () {
  // horizontalScrollBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea19horizontalScrollBarEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "horizontalScrollBar", args)
 }

}


func (this *QAbstractScrollArea) maximumViewportSize(args ...interface{}) () {
  // maximumViewportSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea19maximumViewportSizeEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "maximumViewportSize", args)
 }

}


func NewQAbstractScrollArea(args ...interface{})() {
}


func (this *QAbstractScrollArea) setViewport(args ...interface{}) () {
  // setViewport(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea11setViewportEP7QWidget
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setViewport", args)
 }

}


func (this *QAbstractScrollArea) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "minimumSizeHint", args)
 }

}


func (this *QAbstractScrollArea) setCornerWidget(args ...interface{}) () {
  // setCornerWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setCornerWidget", args)
 }

}


func (this *QAbstractScrollArea) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "metaObject", args)
 }

}


func (this *QAbstractScrollArea) setupViewport(args ...interface{}) () {
  // setupViewport(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea13setupViewportEP7QWidget
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setupViewport", args)
 }

}


func (this *QAbstractScrollArea) cornerWidget(args ...interface{}) () {
  // cornerWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea12cornerWidgetEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "cornerWidget", args)
 }

}


func (this *QAbstractScrollArea) verticalScrollBar(args ...interface{}) () {
  // verticalScrollBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea17verticalScrollBarEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "verticalScrollBar", args)
 }

}


func (this *QAbstractScrollArea) viewport(args ...interface{}) () {
  // viewport()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea8viewportEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "viewport", args)
 }

}


func (this *QAbstractScrollArea) FreeQAbstractScrollArea(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "~QAbstractScrollArea", args)
 }

}


func (this *QAbstractScrollArea) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QAbstractScrollArea8sizeHintEv
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "sizeHint", args)
 }

}


func (this *QAbstractScrollArea) setHorizontalScrollBar(args ...interface{}) () {
  // setHorizontalScrollBar(class QScrollBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setHorizontalScrollBar", args)
 }

}


func (this *QAbstractScrollArea) setVerticalScrollBar(args ...interface{}) () {
  // setVerticalScrollBar(class QScrollBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setVerticalScrollBar", args)
 }

}

// <= body block end

