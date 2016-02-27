package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qmovie.h
// dst-file: /src/gui/qmovie.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "runtime"
import "qtrt"
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QMovie::setScaledSize(const QSize & size);
extern void C_ZN6QMovie13setScaledSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  int QMovie::nextFrameDelay();
extern int32_t C_ZNK6QMovie14nextFrameDelayEv(void* qthis); // 4
  // proto:  int QMovie::speed();
extern int32_t C_ZNK6QMovie5speedEv(void* qthis); // 4
  // proto:  int QMovie::frameCount();
extern int32_t C_ZNK6QMovie10frameCountEv(void* qthis); // 4
  // proto:  int QMovie::currentFrameNumber();
extern int32_t C_ZNK6QMovie18currentFrameNumberEv(void* qthis); // 4
  // proto:  int QMovie::loopCount();
extern int32_t C_ZNK6QMovie9loopCountEv(void* qthis); // 4
  // proto:  void QMovie::setFormat(const QByteArray & format);
extern void C_ZN6QMovie9setFormatERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QByteArray QMovie::format();
extern void* C_ZNK6QMovie6formatEv(void* qthis); // 4
  // proto:  void QMovie::start();
extern void C_ZN6QMovie5startEv(void* qthis); // 4
  // proto:  QMovie::MovieState QMovie::state();
extern void C_ZNK6QMovie5stateEv(void* qthis); // 4
  // proto:  QColor QMovie::backgroundColor();
extern void* C_ZNK6QMovie15backgroundColorEv(void* qthis); // 4
  // proto:  void QMovie::setDevice(QIODevice * device);
extern void C_ZN6QMovie9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  bool QMovie::jumpToFrame(int frameNumber);
extern bool C_ZN6QMovie11jumpToFrameEi(void* qthis, int32_t arg0); // 4
  // proto:  QRect QMovie::frameRect();
extern void* C_ZNK6QMovie9frameRectEv(void* qthis); // 4
  // proto:  bool QMovie::isValid();
extern bool C_ZNK6QMovie7isValidEv(void* qthis); // 4
  // proto:  void QMovie::setBackgroundColor(const QColor & color);
extern void C_ZN6QMovie18setBackgroundColorERK6QColor(void* qthis, void* arg0); // 4
  // proto:  void QMovie::stop();
extern void C_ZN6QMovie4stopEv(void* qthis); // 4
  // proto:  QString QMovie::fileName();
extern void* C_ZNK6QMovie8fileNameEv(void* qthis); // 4
  // proto:  QImage QMovie::currentImage();
extern void* C_ZNK6QMovie12currentImageEv(void* qthis); // 4
  // proto:  QIODevice * QMovie::device();
extern void* C_ZNK6QMovie6deviceEv(void* qthis); // 4
  // proto:  const QMetaObject * QMovie::metaObject();
extern void C_ZNK6QMovie10metaObjectEv(void* qthis); // 4
  // proto:  QPixmap QMovie::currentPixmap();
extern void* C_ZNK6QMovie13currentPixmapEv(void* qthis); // 4
  // proto:  void QMovie::~QMovie();
extern void C_ZN6QMovieD2Ev(void* qthis); // 4
  // proto:  void QMovie::setFileName(const QString & fileName);
extern void C_ZN6QMovie11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QMovie::setSpeed(int percentSpeed);
extern void C_ZN6QMovie8setSpeedEi(void* qthis, int32_t arg0); // 4
  // proto:  QMovie::CacheMode QMovie::cacheMode();
extern void C_ZNK6QMovie9cacheModeEv(void* qthis); // 4
  // proto:  QSize QMovie::scaledSize();
extern void* C_ZN6QMovie10scaledSizeEv(void* qthis); // 4
  // proto:  void QMovie::QMovie(QObject * parent);
extern void* C_ZN6QMovieC2EP7QObject(void* arg0); // 3
  // proto:  void QMovie::QMovie(QIODevice * device, const QByteArray & format, QObject * parent);
extern void* C_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QMovie::QMovie(const QString & fileName, const QByteArray & format, QObject * parent);
extern void* C_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject(void* arg0, void* arg1, void* arg2); // 3
  // proto: static QList<QByteArray> QMovie::supportedFormats();
extern void C_ZN6QMovie16supportedFormatsEv(); // 4
  // proto:  bool QMovie::jumpToNextFrame();
extern bool C_ZN6QMovie15jumpToNextFrameEv(void* qthis); // 4
  // proto:  void QMovie::setPaused(bool paused);
extern void C_ZN6QMovie9setPausedEb(void* qthis, bool arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QMovie)=1
type QMovie struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _updated QMovie_updated_signal;
//  _stateChanged QMovie_stateChanged_signal;
//  _started QMovie_started_signal;
//  _resized QMovie_resized_signal;
//  _finished QMovie_finished_signal;
//  _error QMovie_error_signal;
//  _frameChanged QMovie_frameChanged_signal;
}

// setScaledSize(const class QSize &)
func (this *QMovie) SetScaledSize(args ...interface{}) () {
  // setScaledSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie13setScaledSizeERK5QSize
    // invoke: void setScaledSize(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie13setScaledSizeERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setScaledSize", args)
  }

  return
}

// nextFrameDelay()
func (this *QMovie) NextFrameDelay(args ...interface{}) (ret interface{}) {
  // nextFrameDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie14nextFrameDelayEv
    // invoke: int nextFrameDelay()
    var ret0 = C.C_ZNK6QMovie14nextFrameDelayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "nextFrameDelay", args)
  }

  return
}

// speed()
func (this *QMovie) Speed(args ...interface{}) (ret interface{}) {
  // speed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie5speedEv
    // invoke: int speed()
    var ret0 = C.C_ZNK6QMovie5speedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "speed", args)
  }

  return
}

// frameCount()
func (this *QMovie) FrameCount(args ...interface{}) (ret interface{}) {
  // frameCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie10frameCountEv
    // invoke: int frameCount()
    var ret0 = C.C_ZNK6QMovie10frameCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "frameCount", args)
  }

  return
}

// currentFrameNumber()
func (this *QMovie) CurrentFrameNumber(args ...interface{}) (ret interface{}) {
  // currentFrameNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie18currentFrameNumberEv
    // invoke: int currentFrameNumber()
    var ret0 = C.C_ZNK6QMovie18currentFrameNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "currentFrameNumber", args)
  }

  return
}

// loopCount()
func (this *QMovie) LoopCount(args ...interface{}) (ret interface{}) {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie9loopCountEv
    // invoke: int loopCount()
    var ret0 = C.C_ZNK6QMovie9loopCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "loopCount", args)
  }

  return
}

// setFormat(const class QByteArray &)
func (this *QMovie) SetFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie9setFormatERK10QByteArray
    // invoke: void setFormat(const class QByteArray &)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie9setFormatERK10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setFormat", args)
  }

  return
}

// format()
func (this *QMovie) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie6formatEv
    // invoke: QByteArray format()
    var ret0 = C.C_ZNK6QMovie6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "format", args)
  }

  return
}

// start()
func (this *QMovie) Start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie5startEv
    // invoke: void start()
    C.C_ZN6QMovie5startEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMovie", "start", args)
  }

  return
}

// state()
func (this *QMovie) State(args ...interface{}) () {
  // state()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie5stateEv
    // invoke: QMovie::MovieState state()
    C.C_ZNK6QMovie5stateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMovie", "state", args)
  }

  return
}

// backgroundColor()
func (this *QMovie) BackgroundColor(args ...interface{}) (ret interface{}) {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie15backgroundColorEv
    // invoke: QColor backgroundColor()
    var ret0 = C.C_ZNK6QMovie15backgroundColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "backgroundColor", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QMovie) SetDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie9setDeviceEP9QIODevice
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(*qtcore.QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie9setDeviceEP9QIODevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setDevice", args)
  }

  return
}

// jumpToFrame(int)
func (this *QMovie) JumpToFrame(args ...interface{}) (ret interface{}) {
  // jumpToFrame(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie11jumpToFrameEi
    // invoke: bool jumpToFrame(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QMovie11jumpToFrameEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "jumpToFrame", args)
  }

  return
}

// frameRect()
func (this *QMovie) FrameRect(args ...interface{}) (ret interface{}) {
  // frameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie9frameRectEv
    // invoke: QRect frameRect()
    var ret0 = C.C_ZNK6QMovie9frameRectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "frameRect", args)
  }

  return
}

// isValid()
func (this *QMovie) IsValid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK6QMovie7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "isValid", args)
  }

  return
}

// setBackgroundColor(const class QColor &)
func (this *QMovie) SetBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie18setBackgroundColorERK6QColor
    // invoke: void setBackgroundColor(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie18setBackgroundColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setBackgroundColor", args)
  }

  return
}

// stop()
func (this *QMovie) Stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie4stopEv
    // invoke: void stop()
    C.C_ZN6QMovie4stopEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMovie", "stop", args)
  }

  return
}

// fileName()
func (this *QMovie) FileName(args ...interface{}) (ret interface{}) {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie8fileNameEv
    // invoke: QString fileName()
    var ret0 = C.C_ZNK6QMovie8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "fileName", args)
  }

  return
}

// currentImage()
func (this *QMovie) CurrentImage(args ...interface{}) (ret interface{}) {
  // currentImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie12currentImageEv
    // invoke: QImage currentImage()
    var ret0 = C.C_ZNK6QMovie12currentImageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "currentImage", args)
  }

  return
}

// device()
func (this *QMovie) Device(args ...interface{}) (ret interface{}) {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie6deviceEv
    // invoke: QIODevice * device()
    var ret0 = C.C_ZNK6QMovie6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QIODevice{}) // "QIODevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "device", args)
  }

  return
}

// metaObject()
func (this *QMovie) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QMovie10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMovie", "metaObject", args)
  }

  return
}

// currentPixmap()
func (this *QMovie) CurrentPixmap(args ...interface{}) (ret interface{}) {
  // currentPixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie13currentPixmapEv
    // invoke: QPixmap currentPixmap()
    var ret0 = C.C_ZNK6QMovie13currentPixmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "currentPixmap", args)
  }

  return
}

// ~QMovie()
func (this *QMovie) Free(args ...interface{}) () {
  // ~QMovie()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovieD0Ev
    // invoke: void ~QMovie()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN6QMovieD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QMovie", "~QMovie", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QMovie) SetFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setFileName", args)
  }

  return
}

// setSpeed(int)
func (this *QMovie) SetSpeed(args ...interface{}) () {
  // setSpeed(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie8setSpeedEi
    // invoke: void setSpeed(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie8setSpeedEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setSpeed", args)
  }

  return
}

// cacheMode()
func (this *QMovie) CacheMode(args ...interface{}) () {
  // cacheMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie9cacheModeEv
    // invoke: QMovie::CacheMode cacheMode()
    C.C_ZNK6QMovie9cacheModeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMovie", "cacheMode", args)
  }

  return
}

// scaledSize()
func (this *QMovie) ScaledSize(args ...interface{}) (ret interface{}) {
  // scaledSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie10scaledSizeEv
    // invoke: QSize scaledSize()
    var ret0 = C.C_ZN6QMovie10scaledSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "scaledSize", args)
  }

  return
}

// QMovie(class QObject *)
func GcfreeQMovie(this *QMovie) {
  qtrt.UniverseFree(this)
}
func NewQMovie(args ...interface{}) *QMovie {
  // QMovie(class QObject *)
  // QMovie(class QIODevice *, const class QByteArray &, class QObject *)
  // QMovie(const class QString &, const class QByteArray &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QIODevice{}) // "QIODevice *"
  vtys[1][1] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[1][2] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[2][2] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovieC1EP7QObject
    // invoke: void QMovie(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QMovieC2EP7QObject(arg0)
    this := &QMovie{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQMovie)
    return this
  case 1:
    // invoke: _ZN6QMovieC1EP9QIODeviceRK10QByteArrayP7QObject
    // invoke: void QMovie(class QIODevice *, const class QByteArray &, class QObject *)
    var arg0 = args[0].(*qtcore.QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QMovieC2EP9QIODeviceRK10QByteArrayP7QObject(arg0, arg1, arg2)
    this := &QMovie{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQMovie)
    return this
  case 2:
    // invoke: _ZN6QMovieC1ERK7QStringRK10QByteArrayP7QObject
    // invoke: void QMovie(const class QString &, const class QByteArray &, class QObject *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QMovieC2ERK7QStringRK10QByteArrayP7QObject(arg0, arg1, arg2)
    this := &QMovie{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQMovie)
    return this
  default:
    qtrt.ErrorResolve("QMovie", "QMovie", args)
  }

  return nil // QMovie{Qclsinst:qthis}
}

// supportedFormats()
func (this *QMovie) SupportedFormats_s(args ...interface{}) () {
  // supportedFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie16supportedFormatsEv
    // invoke: QList<QByteArray> supportedFormats()
    C.C_ZN6QMovie16supportedFormatsEv()
  default:
    qtrt.ErrorResolve("QMovie", "supportedFormats", args)
  }

  return
}

// jumpToNextFrame()
func (this *QMovie) JumpToNextFrame(args ...interface{}) (ret interface{}) {
  // jumpToNextFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie15jumpToNextFrameEv
    // invoke: bool jumpToNextFrame()
    var ret0 = C.C_ZN6QMovie15jumpToNextFrameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMovie", "jumpToNextFrame", args)
  }

  return
}

// setPaused(_Bool)
func (this *QMovie) SetPaused(args ...interface{}) () {
  // setPaused(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie9setPausedEb
    // invoke: void setPaused(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN6QMovie9setPausedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMovie", "setPaused", args)
  }

  return
}

// <= body block end

