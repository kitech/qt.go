package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qopenglwidget.h
// dst-file: /src/widgets/qopenglwidget.go
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
// class sizeof(QOpenGLWidget)=1
type QOpenGLWidget struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToResize QOpenGLWidget_aboutToResize_signal;
//  _resized QOpenGLWidget_resized_signal;
//  _frameSwapped QOpenGLWidget_frameSwapped_signal;
//  _aboutToCompose QOpenGLWidget_aboutToCompose_signal;
}


func (this *QOpenGLWidget) FreeQOpenGLWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "~QOpenGLWidget", args)
  }

}


func (this *QOpenGLWidget) defaultFramebufferObject(args ...interface{}) () {
  // defaultFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget24defaultFramebufferObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "defaultFramebufferObject", args)
  }

}


func NewQOpenGLWidget(args ...interface{}) QOpenGLWidget {
  return QOpenGLWidget{}
}


func (this *QOpenGLWidget) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget7isValidEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "isValid", args)
  }

}


func (this *QOpenGLWidget) context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget7contextEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "context", args)
  }

}


func (this *QOpenGLWidget) doneCurrent(args ...interface{}) () {
  // doneCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget11doneCurrentEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "doneCurrent", args)
  }

}


func (this *QOpenGLWidget) makeCurrent(args ...interface{}) () {
  // makeCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget11makeCurrentEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "makeCurrent", args)
  }

}


func (this *QOpenGLWidget) grabFramebuffer(args ...interface{}) () {
  // grabFramebuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget15grabFramebufferEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "grabFramebuffer", args)
  }

}


func (this *QOpenGLWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "metaObject", args)
  }

}


func (this *QOpenGLWidget) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "setFormat", args)
  }

}


func (this *QOpenGLWidget) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget6formatEv
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "format", args)
  }

}

// <= body block end

