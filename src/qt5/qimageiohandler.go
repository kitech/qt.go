package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qimageiohandler.h
// dst-file: /src/gui/qimageiohandler.go
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
// class sizeof(QImageIOHandler)=1
type QImageIOHandler struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QImageIOPlugin)=1
type QImageIOPlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQImageIOHandler(args ...interface{})() {
}


func (this *QImageIOHandler) imageCount(args ...interface{}) () {
  // imageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler10imageCountEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "imageCount", args)
 }

}


func (this *QImageIOHandler) currentImageRect(args ...interface{}) () {
  // currentImageRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler16currentImageRectEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "currentImageRect", args)
 }

}


func (this *QImageIOHandler) jumpToImage(args ...interface{}) () {
  // jumpToImage(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler11jumpToImageEi
  default:
    qtrt.ErrorResolve("QImageIOHandler", "jumpToImage", args)
 }

}


func (this *QImageIOHandler) currentImageNumber(args ...interface{}) () {
  // currentImageNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler18currentImageNumberEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "currentImageNumber", args)
 }

}


func (this *QImageIOHandler) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler9setFormatERK10QByteArray
  case 1:
    // invoke: _ZNK15QImageIOHandler9setFormatERK10QByteArray
  default:
    qtrt.ErrorResolve("QImageIOHandler", "setFormat", args)
 }

}


func (this *QImageIOHandler) jumpToNextImage(args ...interface{}) () {
  // jumpToNextImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler15jumpToNextImageEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "jumpToNextImage", args)
 }

}


func (this *QImageIOHandler) FreeQImageIOHandler(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QImageIOHandler", "~QImageIOHandler", args)
 }

}


func (this *QImageIOHandler) loopCount(args ...interface{}) () {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler9loopCountEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "loopCount", args)
 }

}


func (this *QImageIOHandler) read(args ...interface{}) () {
  // read(class QImage *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "QImage *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler4readEP6QImage
  default:
    qtrt.ErrorResolve("QImageIOHandler", "read", args)
 }

}


func (this *QImageIOHandler) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler4nameEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "name", args)
 }

}


func (this *QImageIOHandler) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler6formatEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "format", args)
 }

}


func (this *QImageIOHandler) nextImageDelay(args ...interface{}) () {
  // nextImageDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler14nextImageDelayEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "nextImageDelay", args)
 }

}


func (this *QImageIOHandler) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QImageIOHandler", "setDevice", args)
 }

}


func (this *QImageIOHandler) canRead(args ...interface{}) () {
  // canRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler7canReadEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "canRead", args)
 }

}


func (this *QImageIOHandler) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QImageIOHandler6deviceEv
  default:
    qtrt.ErrorResolve("QImageIOHandler", "device", args)
 }

}


func (this *QImageIOHandler) write(args ...interface{}) () {
  // write(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler5writeERK6QImage
  default:
    qtrt.ErrorResolve("QImageIOHandler", "write", args)
 }

}


func (this *QImageIOPlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QImageIOPlugin10metaObjectEv
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "metaObject", args)
 }

}


func (this *QImageIOPlugin) FreeQImageIOPlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "~QImageIOPlugin", args)
 }

}


func (this *QImageIOPlugin) create(args ...interface{}) () {
  // create(class QIODevice *, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "create", args)
 }

}


func NewQImageIOPlugin(args ...interface{})() {
}

// <= body block end

