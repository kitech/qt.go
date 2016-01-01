package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qsurfaceformat.h
// dst-file: /src/gui/qsurfaceformat.go
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
// class sizeof(QSurfaceFormat)=8
type QSurfaceFormat struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QSurfaceFormat) defaultFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "defaultFormat", args)
 }

}


func (this *QSurfaceFormat) setAlphaBufferSize(args ...interface{}) () {
  // setAlphaBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat18setAlphaBufferSizeEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setAlphaBufferSize", args)
 }

}


func (this *QSurfaceFormat) setMinorVersion(args ...interface{}) () {
  // setMinorVersion(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat15setMinorVersionEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setMinorVersion", args)
 }

}


func (this *QSurfaceFormat) stencilBufferSize(args ...interface{}) () {
  // stencilBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat17stencilBufferSizeEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stencilBufferSize", args)
 }

}


func (this *QSurfaceFormat) setRedBufferSize(args ...interface{}) () {
  // setRedBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat16setRedBufferSizeEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setRedBufferSize", args)
 }

}


func (this *QSurfaceFormat) setDepthBufferSize(args ...interface{}) () {
  // setDepthBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat18setDepthBufferSizeEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setDepthBufferSize", args)
 }

}


func (this *QSurfaceFormat) majorVersion(args ...interface{}) () {
  // majorVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12majorVersionEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "majorVersion", args)
 }

}


func (this *QSurfaceFormat) setSamples(args ...interface{}) () {
  // setSamples(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat10setSamplesEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setSamples", args)
 }

}


func (this *QSurfaceFormat) setMajorVersion(args ...interface{}) () {
  // setMajorVersion(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat15setMajorVersionEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setMajorVersion", args)
 }

}


func (this *QSurfaceFormat) setDefaultFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setDefaultFormat", args)
 }

}


func (this *QSurfaceFormat) greenBufferSize(args ...interface{}) () {
  // greenBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat15greenBufferSizeEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "greenBufferSize", args)
 }

}


func (this *QSurfaceFormat) minorVersion(args ...interface{}) () {
  // minorVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12minorVersionEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "minorVersion", args)
 }

}


func (this *QSurfaceFormat) setStencilBufferSize(args ...interface{}) () {
  // setStencilBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat20setStencilBufferSizeEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setStencilBufferSize", args)
 }

}


func (this *QSurfaceFormat) swapInterval(args ...interface{}) () {
  // swapInterval()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat12swapIntervalEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "swapInterval", args)
 }

}


func (this *QSurfaceFormat) setVersion(args ...interface{}) () {
  // setVersion(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat10setVersionEii
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setVersion", args)
 }

}


func (this *QSurfaceFormat) hasAlpha(args ...interface{}) () {
  // hasAlpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat8hasAlphaEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "hasAlpha", args)
 }

}


func NewQSurfaceFormat(args ...interface{})() {
}


func (this *QSurfaceFormat) version(args ...interface{}) () {
  // version()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat7versionEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "version", args)
 }

}


func (this *QSurfaceFormat) blueBufferSize(args ...interface{}) () {
  // blueBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat14blueBufferSizeEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "blueBufferSize", args)
 }

}


func (this *QSurfaceFormat) redBufferSize(args ...interface{}) () {
  // redBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat13redBufferSizeEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "redBufferSize", args)
 }

}


func (this *QSurfaceFormat) FreeQSurfaceFormat(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "~QSurfaceFormat", args)
 }

}


func (this *QSurfaceFormat) setGreenBufferSize(args ...interface{}) () {
  // setGreenBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat18setGreenBufferSizeEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setGreenBufferSize", args)
 }

}


func (this *QSurfaceFormat) samples(args ...interface{}) () {
  // samples()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat7samplesEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "samples", args)
 }

}


func (this *QSurfaceFormat) depthBufferSize(args ...interface{}) () {
  // depthBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat15depthBufferSizeEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "depthBufferSize", args)
 }

}


func (this *QSurfaceFormat) setBlueBufferSize(args ...interface{}) () {
  // setBlueBufferSize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat17setBlueBufferSizeEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setBlueBufferSize", args)
 }

}


func (this *QSurfaceFormat) alphaBufferSize(args ...interface{}) () {
  // alphaBufferSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat15alphaBufferSizeEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "alphaBufferSize", args)
 }

}


func (this *QSurfaceFormat) stereo(args ...interface{}) () {
  // stereo()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QSurfaceFormat6stereoEv
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "stereo", args)
 }

}


func (this *QSurfaceFormat) setSwapInterval(args ...interface{}) () {
  // setSwapInterval(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat15setSwapIntervalEi
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setSwapInterval", args)
 }

}


func (this *QSurfaceFormat) setStereo(args ...interface{}) () {
  // setStereo(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QSurfaceFormat9setStereoEb
  default:
    qtrt.ErrorResolve("QSurfaceFormat", "setStereo", args)
 }

}

// <= body block end

