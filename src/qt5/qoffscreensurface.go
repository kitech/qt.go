package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qoffscreensurface.h
// dst-file: /src/gui/qoffscreensurface.go
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
// class sizeof(QOffscreenSurface)=1
type QOffscreenSurface struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _screenChanged QOffscreenSurface_screenChanged_signal;
}


func (this *QOffscreenSurface) FreeQOffscreenSurface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "~QOffscreenSurface", args)
  }

}


func NewQOffscreenSurface(args ...interface{}) QOffscreenSurface {
  return QOffscreenSurface{}
}


func (this *QOffscreenSurface) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6screenEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "screen", args)
  }

}


func (this *QOffscreenSurface) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "setFormat", args)
  }

}


func (this *QOffscreenSurface) setScreen(args ...interface{}) () {
  // setScreen(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface9setScreenEP7QScreen
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "setScreen", args)
  }

}


func (this *QOffscreenSurface) requestedFormat(args ...interface{}) () {
  // requestedFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface15requestedFormatEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "requestedFormat", args)
  }

}


func (this *QOffscreenSurface) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6formatEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "format", args)
  }

}


func (this *QOffscreenSurface) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6handleEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "handle", args)
  }

}


func (this *QOffscreenSurface) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface10metaObjectEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "metaObject", args)
  }

}


func (this *QOffscreenSurface) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface7destroyEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "destroy", args)
  }

}


func (this *QOffscreenSurface) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface7isValidEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "isValid", args)
  }

}


func (this *QOffscreenSurface) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface4sizeEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "size", args)
  }

}


func (this *QOffscreenSurface) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface6createEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "create", args)
  }

}

// <= body block end

