package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qscreen.h
// dst-file: /src/gui/qscreen.go
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
// class sizeof(QScreen)=1
type QScreen struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _geometryChanged QScreen_geometryChanged_signal;
//  _virtualGeometryChanged QScreen_virtualGeometryChanged_signal;
//  _refreshRateChanged QScreen_refreshRateChanged_signal;
//  _availableGeometryChanged QScreen_availableGeometryChanged_signal;
//  _physicalSizeChanged QScreen_physicalSizeChanged_signal;
//  _physicalDotsPerInchChanged QScreen_physicalDotsPerInchChanged_signal;
//  _logicalDotsPerInchChanged QScreen_logicalDotsPerInchChanged_signal;
//  _primaryOrientationChanged QScreen_primaryOrientationChanged_signal;
//  _orientationChanged QScreen_orientationChanged_signal;
}


func (this *QScreen) logicalDotsPerInchY(args ...interface{}) () {
  // logicalDotsPerInchY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen19logicalDotsPerInchYEv
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchY", args)
 }

}


func (this *QScreen) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen8geometryEv
  default:
    qtrt.ErrorResolve("QScreen", "geometry", args)
 }

}


func (this *QScreen) grabWindow(args ...interface{}) () {
  // grabWindow(WId, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "WId"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN7QScreen10grabWindowEiiiii
  default:
    qtrt.ErrorResolve("QScreen", "grabWindow", args)
 }

}


func (this *QScreen) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen4sizeEv
  default:
    qtrt.ErrorResolve("QScreen", "size", args)
 }

}


func (this *QScreen) physicalSize(args ...interface{}) () {
  // physicalSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen12physicalSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalSize", args)
 }

}


func (this *QScreen) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen6handleEv
  default:
    qtrt.ErrorResolve("QScreen", "handle", args)
 }

}


func (this *QScreen) availableVirtualGeometry(args ...interface{}) () {
  // availableVirtualGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen24availableVirtualGeometryEv
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualGeometry", args)
 }

}


func (this *QScreen) FreeQScreen(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QScreen", "~QScreen", args)
 }

}


func (this *QScreen) virtualSize(args ...interface{}) () {
  // virtualSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen11virtualSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "virtualSize", args)
 }

}


func (this *QScreen) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen16devicePixelRatioEv
  default:
    qtrt.ErrorResolve("QScreen", "devicePixelRatio", args)
 }

}


func (this *QScreen) virtualSiblings(args ...interface{}) () {
  // virtualSiblings()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen15virtualSiblingsEv
  default:
    qtrt.ErrorResolve("QScreen", "virtualSiblings", args)
 }

}


func NewQScreen(args ...interface{})() {
}


func (this *QScreen) virtualGeometry(args ...interface{}) () {
  // virtualGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen15virtualGeometryEv
  default:
    qtrt.ErrorResolve("QScreen", "virtualGeometry", args)
 }

}


func (this *QScreen) logicalDotsPerInch(args ...interface{}) () {
  // logicalDotsPerInch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen18logicalDotsPerInchEv
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInch", args)
 }

}


func (this *QScreen) physicalDotsPerInch(args ...interface{}) () {
  // physicalDotsPerInch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen19physicalDotsPerInchEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInch", args)
 }

}


func (this *QScreen) refreshRate(args ...interface{}) () {
  // refreshRate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen11refreshRateEv
  default:
    qtrt.ErrorResolve("QScreen", "refreshRate", args)
 }

}


func (this *QScreen) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen10metaObjectEv
  default:
    qtrt.ErrorResolve("QScreen", "metaObject", args)
 }

}


func (this *QScreen) availableSize(args ...interface{}) () {
  // availableSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen13availableSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "availableSize", args)
 }

}


func (this *QScreen) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen4nameEv
  default:
    qtrt.ErrorResolve("QScreen", "name", args)
 }

}


func (this *QScreen) availableVirtualSize(args ...interface{}) () {
  // availableVirtualSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen20availableVirtualSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualSize", args)
 }

}


func (this *QScreen) logicalDotsPerInchX(args ...interface{}) () {
  // logicalDotsPerInchX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen19logicalDotsPerInchXEv
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchX", args)
 }

}


func (this *QScreen) availableGeometry(args ...interface{}) () {
  // availableGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen17availableGeometryEv
  default:
    qtrt.ErrorResolve("QScreen", "availableGeometry", args)
 }

}


func (this *QScreen) physicalDotsPerInchX(args ...interface{}) () {
  // physicalDotsPerInchX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen20physicalDotsPerInchXEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchX", args)
 }

}


func (this *QScreen) physicalDotsPerInchY(args ...interface{}) () {
  // physicalDotsPerInchY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen20physicalDotsPerInchYEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchY", args)
 }

}


func (this *QScreen) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen5depthEv
  default:
    qtrt.ErrorResolve("QScreen", "depth", args)
 }

}

// <= body block end

