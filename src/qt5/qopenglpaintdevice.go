package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qopenglpaintdevice.h
// dst-file: /src/gui/qopenglpaintdevice.go
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
// class sizeof(QOpenGLPaintDevice)=1
type QOpenGLPaintDevice struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QOpenGLPaintDevice) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice11paintEngineEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintEngine", args)
  }

}


func (this *QOpenGLPaintDevice) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice4sizeEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "size", args)
  }

}


func (this *QOpenGLPaintDevice) setPaintFlipped(args ...interface{}) () {
  // setPaintFlipped(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice15setPaintFlippedEb
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setPaintFlipped", args)
  }

}


func (this *QOpenGLPaintDevice) FreeQOpenGLPaintDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "~QOpenGLPaintDevice", args)
  }

}


func NewQOpenGLPaintDevice(args ...interface{}) QOpenGLPaintDevice {
  return QOpenGLPaintDevice{}
}


func (this *QOpenGLPaintDevice) context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice7contextEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "context", args)
  }

}


func (this *QOpenGLPaintDevice) setDevicePixelRatio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice19setDevicePixelRatioEd
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDevicePixelRatio", args)
  }

}


func (this *QOpenGLPaintDevice) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice7devTypeEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "devType", args)
  }

}


func (this *QOpenGLPaintDevice) dotsPerMeterX(args ...interface{}) () {
  // dotsPerMeterX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice13dotsPerMeterXEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterX", args)
  }

}


func (this *QOpenGLPaintDevice) setDotsPerMeterX(args ...interface{}) () {
  // setDotsPerMeterX(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice16setDotsPerMeterXEd
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterX", args)
  }

}


func (this *QOpenGLPaintDevice) dotsPerMeterY(args ...interface{}) () {
  // dotsPerMeterY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice13dotsPerMeterYEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "dotsPerMeterY", args)
  }

}


func (this *QOpenGLPaintDevice) setDotsPerMeterY(args ...interface{}) () {
  // setDotsPerMeterY(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice16setDotsPerMeterYEd
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setDotsPerMeterY", args)
  }

}


func (this *QOpenGLPaintDevice) paintFlipped(args ...interface{}) () {
  // paintFlipped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QOpenGLPaintDevice12paintFlippedEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "paintFlipped", args)
  }

}


func (this *QOpenGLPaintDevice) setSize(args ...interface{}) () {
  // setSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice7setSizeERK5QSize
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "setSize", args)
  }

}


func (this *QOpenGLPaintDevice) ensureActiveTarget(args ...interface{}) () {
  // ensureActiveTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QOpenGLPaintDevice18ensureActiveTargetEv
  default:
    qtrt.ErrorResolve("QOpenGLPaintDevice", "ensureActiveTarget", args)
  }

}

// <= body block end

