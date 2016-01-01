package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtGui/qpaintdevice.h
// dst-file: /src/gui/qpaintdevice.go
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
// class sizeof(QPaintDevice)=24
type QPaintDevice struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPaintDevice) physicalDpiY(args ...interface{}) () {
  // physicalDpiY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice12physicalDpiYEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiY", args)
  }

}


func (this *QPaintDevice) heightMM(args ...interface{}) () {
  // heightMM()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice8heightMMEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "heightMM", args)
  }

}


func (this *QPaintDevice) colorCount(args ...interface{}) () {
  // colorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice10colorCountEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "colorCount", args)
  }

}


func (this *QPaintDevice) physicalDpiX(args ...interface{}) () {
  // physicalDpiX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice12physicalDpiXEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "physicalDpiX", args)
  }

}


func (this *QPaintDevice) widthMM(args ...interface{}) () {
  // widthMM()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice7widthMMEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "widthMM", args)
  }

}


func (this *QPaintDevice) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice7devTypeEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "devType", args)
  }

}


func (this *QPaintDevice) paintingActive(args ...interface{}) () {
  // paintingActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice14paintingActiveEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "paintingActive", args)
  }

}


func (this *QPaintDevice) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice5widthEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "width", args)
  }

}


func NewQPaintDevice(args ...interface{}) QPaintDevice {
  return QPaintDevice{}
}


func (this *QPaintDevice) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice16devicePixelRatioEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "devicePixelRatio", args)
  }

}


func (this *QPaintDevice) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice6heightEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "height", args)
  }

}


func (this *QPaintDevice) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice5depthEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "depth", args)
  }

}


func (this *QPaintDevice) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice11paintEngineEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "paintEngine", args)
  }

}


func (this *QPaintDevice) logicalDpiY(args ...interface{}) () {
  // logicalDpiY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice11logicalDpiYEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiY", args)
  }

}


func (this *QPaintDevice) FreeQPaintDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPaintDevice", "~QPaintDevice", args)
  }

}


func (this *QPaintDevice) logicalDpiX(args ...interface{}) () {
  // logicalDpiX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QPaintDevice11logicalDpiXEv
  default:
    qtrt.ErrorResolve("QPaintDevice", "logicalDpiX", args)
  }

}

// <= body block end

