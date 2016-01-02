package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QImageIOHandler::QImageIOHandler(const QImageIOHandler & );
extern void* dector_ZN15QImageIOHandlerC1ERKS_(void* arg0);
extern void _ZN15QImageIOHandlerC1ERKS_(void* qthis, void* arg0);
  // proto:  int QImageIOHandler::imageCount();
extern void _ZNK15QImageIOHandler10imageCountEv(void* qthis);
  // proto:  QRect QImageIOHandler::currentImageRect();
extern void _ZNK15QImageIOHandler16currentImageRectEv(void* qthis);
  // proto:  bool QImageIOHandler::jumpToImage(int imageNumber);
extern void _ZN15QImageIOHandler11jumpToImageEi(void* qthis, int arg0);
  // proto:  int QImageIOHandler::currentImageNumber();
extern void _ZNK15QImageIOHandler18currentImageNumberEv(void* qthis);
  // proto:  void QImageIOHandler::setFormat(const QByteArray & format);
extern void _ZN15QImageIOHandler9setFormatERK10QByteArray(void* qthis, void* arg0);
  // proto:  bool QImageIOHandler::jumpToNextImage();
extern void _ZN15QImageIOHandler15jumpToNextImageEv(void* qthis);
  // proto:  void QImageIOHandler::~QImageIOHandler();
extern void _ZN15QImageIOHandlerD0Ev(void* qthis);
  // proto:  int QImageIOHandler::loopCount();
extern void _ZNK15QImageIOHandler9loopCountEv(void* qthis);
  // proto:  void QImageIOHandler::QImageIOHandler();
extern void* dector_ZN15QImageIOHandlerC1Ev();
extern void _ZN15QImageIOHandlerC1Ev(void* qthis);
  // proto:  bool QImageIOHandler::read(QImage * image);
extern void _ZN15QImageIOHandler4readEP6QImage(void* qthis, void* arg0);
  // proto:  QByteArray QImageIOHandler::name();
extern void _ZNK15QImageIOHandler4nameEv(void* qthis);
  // proto:  QByteArray QImageIOHandler::format();
extern void _ZNK15QImageIOHandler6formatEv(void* qthis);
  // proto:  int QImageIOHandler::nextImageDelay();
extern void _ZNK15QImageIOHandler14nextImageDelayEv(void* qthis);
  // proto:  void QImageIOHandler::setDevice(QIODevice * device);
extern void _ZN15QImageIOHandler9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  bool QImageIOHandler::canRead();
extern void _ZNK15QImageIOHandler7canReadEv(void* qthis);
  // proto:  QIODevice * QImageIOHandler::device();
extern void _ZNK15QImageIOHandler6deviceEv(void* qthis);
  // proto:  bool QImageIOHandler::write(const QImage & image);
extern void _ZN15QImageIOHandler5writeERK6QImage(void* qthis, void* arg0);
  // proto:  const QMetaObject * QImageIOPlugin::metaObject();
extern void _ZNK14QImageIOPlugin10metaObjectEv(void* qthis);
  // proto:  void QImageIOPlugin::~QImageIOPlugin();
extern void _ZN14QImageIOPluginD0Ev(void* qthis);
  // proto:  QImageIOHandler * QImageIOPlugin::create(QIODevice * device, const QByteArray & format);
extern void _ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  void QImageIOPlugin::QImageIOPlugin(QObject * parent);
extern void* dector_ZN14QImageIOPluginC1EP7QObject(void* arg0);
extern void _ZN14QImageIOPluginC1EP7QObject(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QImageIOPlugin)=1
type QImageIOPlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QImageIOHandler::QImageIOHandler(const QImageIOHandler & );
func NewQImageIOHandler(args ...interface{}) QImageIOHandler {
  return QImageIOHandler{}
}

  // proto:  int QImageIOHandler::imageCount();
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
    C._ZNK15QImageIOHandler10imageCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "imageCount", args)
  }

}

  // proto:  QRect QImageIOHandler::currentImageRect();
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
    C._ZNK15QImageIOHandler16currentImageRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "currentImageRect", args)
  }

}

  // proto:  bool QImageIOHandler::jumpToImage(int imageNumber);
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
    C._ZN15QImageIOHandler11jumpToImageEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "jumpToImage", args)
  }

}

  // proto:  int QImageIOHandler::currentImageNumber();
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
    C._ZNK15QImageIOHandler18currentImageNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "currentImageNumber", args)
  }

}

  // proto:  void QImageIOHandler::setFormat(const QByteArray & format);
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
    // invoke: void setFormat(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QImageIOHandler9setFormatERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "setFormat", args)
  }

}

  // proto:  bool QImageIOHandler::jumpToNextImage();
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
    C._ZN15QImageIOHandler15jumpToNextImageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "jumpToNextImage", args)
  }

}

  // proto:  void QImageIOHandler::~QImageIOHandler();
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

  // proto:  int QImageIOHandler::loopCount();
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
    C._ZNK15QImageIOHandler9loopCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "loopCount", args)
  }

}

  // proto:  bool QImageIOHandler::read(QImage * image);
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
    // invoke: bool read(class QImage *)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN15QImageIOHandler4readEP6QImage(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "read", args)
  }

}

  // proto:  QByteArray QImageIOHandler::name();
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
    C._ZNK15QImageIOHandler4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "name", args)
  }

}

  // proto:  QByteArray QImageIOHandler::format();
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
    C._ZNK15QImageIOHandler6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "format", args)
  }

}

  // proto:  int QImageIOHandler::nextImageDelay();
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
    C._ZNK15QImageIOHandler14nextImageDelayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "nextImageDelay", args)
  }

}

  // proto:  void QImageIOHandler::setDevice(QIODevice * device);
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
    C._ZN15QImageIOHandler9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "setDevice", args)
  }

}

  // proto:  bool QImageIOHandler::canRead();
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
    // invoke: bool canRead()
    C._ZNK15QImageIOHandler7canReadEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "canRead", args)
  }

}

  // proto:  QIODevice * QImageIOHandler::device();
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
    C._ZNK15QImageIOHandler6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "device", args)
  }

}

  // proto:  bool QImageIOHandler::write(const QImage & image);
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
    C._ZN15QImageIOHandler5writeERK6QImage(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageIOHandler", "write", args)
  }

}

  // proto:  const QMetaObject * QImageIOPlugin::metaObject();
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
    C._ZNK14QImageIOPlugin10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "metaObject", args)
  }

}

  // proto:  void QImageIOPlugin::~QImageIOPlugin();
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

  // proto:  QImageIOHandler * QImageIOPlugin::create(QIODevice * device, const QByteArray & format);
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
    // invoke: QImageIOHandler * create(class QIODevice *, const class QByteArray &)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK14QImageIOPlugin6createEP9QIODeviceRK10QByteArray(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImageIOPlugin", "create", args)
  }

}

  // proto:  void QImageIOPlugin::QImageIOPlugin(QObject * parent);
func NewQImageIOPlugin(args ...interface{}) QImageIOPlugin {
  return QImageIOPlugin{}
}

// <= body block end

