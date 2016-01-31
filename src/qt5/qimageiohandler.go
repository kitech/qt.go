package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtGui/qimageiohandler.h
// dst-file: /src/gui/qimageiohandler.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QIODevice * QImageIOHandler::device();
extern void C_ZNK15QImageIOHandler6deviceEv(void* qthis); // 4
  // proto:  int QImageIOHandler::imageCount();
extern void C_ZNK15QImageIOHandler10imageCountEv(void* qthis); // 4
  // proto:  int QImageIOHandler::loopCount();
extern void C_ZNK15QImageIOHandler9loopCountEv(void* qthis); // 4
  // proto:  void QImageIOHandler::setFormat(const QByteArray & format);
extern void C_ZN15QImageIOHandler9setFormatERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QByteArray QImageIOHandler::format();
extern void C_ZNK15QImageIOHandler6formatEv(void* qthis); // 4
  // proto:  bool QImageIOHandler::write(const QImage & image);
extern void C_ZN15QImageIOHandler5writeERK6QImage(void* qthis, void* arg0); // 4
  // proto:  bool QImageIOHandler::jumpToImage(int imageNumber);
extern void C_ZN15QImageIOHandler11jumpToImageEi(void* qthis, int32_t arg0); // 4
  // proto:  void QImageIOHandler::setDevice(QIODevice * device);
extern void C_ZN15QImageIOHandler9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  bool QImageIOHandler::jumpToNextImage();
extern void C_ZN15QImageIOHandler15jumpToNextImageEv(void* qthis); // 4
  // proto:  int QImageIOHandler::currentImageNumber();
extern void C_ZNK15QImageIOHandler18currentImageNumberEv(void* qthis); // 4
  // proto:  void QImageIOHandler::~QImageIOHandler();
extern void C_ZN15QImageIOHandlerD2Ev(void* qthis); // 4
  // proto:  void QImageIOHandler::QImageIOHandler();
extern void* C_ZN15QImageIOHandlerC2Ev(); // 3
  // proto:  QRect QImageIOHandler::currentImageRect();
extern void C_ZNK15QImageIOHandler16currentImageRectEv(void* qthis); // 4
  // proto:  QByteArray QImageIOHandler::name();
extern void C_ZNK15QImageIOHandler4nameEv(void* qthis); // 4
  // proto:  int QImageIOHandler::nextImageDelay();
extern void C_ZNK15QImageIOHandler14nextImageDelayEv(void* qthis); // 4
  // proto:  void QImageIOPlugin::~QImageIOPlugin();
extern void C_ZN14QImageIOPluginD2Ev(void* qthis); // 4
  // proto:  void QImageIOPlugin::QImageIOPlugin(QObject * parent);
extern void* C_ZN14QImageIOPluginC2EP7QObject(void* arg0); // 3
  // proto:  const QMetaObject * QImageIOPlugin::metaObject();
extern void C_ZNK14QImageIOPlugin10metaObjectEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QImageIOHandler)=1
type QImageIOHandler struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QImageIOPlugin)=1
type QImageIOPlugin struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// device()
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
    // invoke: QIODevice * device()
    var ret = C.C_ZNK15QImageIOHandler6deviceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "device", args)
  }

}

// imageCount()
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
    // invoke: int imageCount()
    var ret = C.C_ZNK15QImageIOHandler10imageCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "imageCount", args)
  }

}

// loopCount()
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
    // invoke: int loopCount()
    var ret = C.C_ZNK15QImageIOHandler9loopCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "loopCount", args)
  }

}

// setFormat(const class QByteArray &)
func (this *QImageIOHandler) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandler9setFormatERK10QByteArray
    // invoke: void setFormat(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QImageIOHandler9setFormatERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "setFormat", args)
  }

}

// format()
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
    // invoke: QByteArray format()
    var ret = C.C_ZNK15QImageIOHandler6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "format", args)
  }

}

// write(const class QImage &)
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
    // invoke: bool write(const class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN15QImageIOHandler5writeERK6QImage(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "write", args)
  }

}

// jumpToImage(int)
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
    // invoke: bool jumpToImage(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN15QImageIOHandler11jumpToImageEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "jumpToImage", args)
  }

}

// setDevice(class QIODevice *)
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
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QImageIOHandler9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "setDevice", args)
  }

}

// jumpToNextImage()
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
    // invoke: bool jumpToNextImage()
    var ret = C.C_ZN15QImageIOHandler15jumpToNextImageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "jumpToNextImage", args)
  }

}

// currentImageNumber()
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
    // invoke: int currentImageNumber()
    var ret = C.C_ZNK15QImageIOHandler18currentImageNumberEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "currentImageNumber", args)
  }

}

// ~QImageIOHandler()
func (this *QImageIOHandler) FreeQImageIOHandler(args ...interface{}) () {
  // ~QImageIOHandler()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandlerD0Ev
    // invoke: void ~QImageIOHandler()
    C.C_ZN15QImageIOHandlerD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "~QImageIOHandler", args)
  }

}

// QImageIOHandler()
func NewQImageIOHandler(args ...interface{}) *QImageIOHandler {
  // QImageIOHandler()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QImageIOHandlerC1Ev
    // invoke: void QImageIOHandler()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QImageIOHandlerC2Ev()
    return &QImageIOHandler{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "QImageIOHandler", args)
  }

  return nil // QImageIOHandler{qclsinst:qthis}
}

// currentImageRect()
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
    // invoke: QRect currentImageRect()
    var ret = C.C_ZNK15QImageIOHandler16currentImageRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "currentImageRect", args)
  }

}

// name()
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
    // invoke: QByteArray name()
    var ret = C.C_ZNK15QImageIOHandler4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "name", args)
  }

}

// nextImageDelay()
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
    // invoke: int nextImageDelay()
    var ret = C.C_ZNK15QImageIOHandler14nextImageDelayEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageIOHandler", "nextImageDelay", args)
  }

}

// ~QImageIOPlugin()
func (this *QImageIOPlugin) FreeQImageIOPlugin(args ...interface{}) () {
  // ~QImageIOPlugin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QImageIOPluginD0Ev
    // invoke: void ~QImageIOPlugin()
    C.C_ZN14QImageIOPluginD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "~QImageIOPlugin", args)
  }

}

// QImageIOPlugin(class QObject *)
func NewQImageIOPlugin(args ...interface{}) *QImageIOPlugin {
  // QImageIOPlugin(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QImageIOPluginC1EP7QObject
    // invoke: void QImageIOPlugin(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QImageIOPluginC2EP7QObject(arg0)
    return &QImageIOPlugin{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "QImageIOPlugin", args)
  }

  return nil // QImageIOPlugin{qclsinst:qthis}
}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QImageIOPlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "metaObject", args)
  }

}

// <= body block end

