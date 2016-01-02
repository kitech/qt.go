package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QMovie::QMovie(QObject * parent);
extern void* dector_ZN6QMovieC1EP7QObject(void* arg0);
extern void _ZN6QMovieC1EP7QObject(void* qthis, void* arg0);
  // proto:  int QMovie::speed();
extern void _ZNK6QMovie5speedEv(void* qthis);
  // proto:  bool QMovie::jumpToNextFrame();
extern void _ZN6QMovie15jumpToNextFrameEv(void* qthis);
  // proto:  int QMovie::frameCount();
extern void _ZNK6QMovie10frameCountEv(void* qthis);
  // proto:  void QMovie::setScaledSize(const QSize & size);
extern void _ZN6QMovie13setScaledSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QMovie::setDevice(QIODevice * device);
extern void _ZN6QMovie9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  QImage QMovie::currentImage();
extern void _ZNK6QMovie12currentImageEv(void* qthis);
  // proto:  bool QMovie::jumpToFrame(int frameNumber);
extern void _ZN6QMovie11jumpToFrameEi(void* qthis, int arg0);
  // proto:  void QMovie::QMovie(const QString & fileName, const QByteArray & format, QObject * parent);
extern void* dector_ZN6QMovieC1ERK7QStringRK10QByteArrayP7QObject(void* arg0, void* arg1, void* arg2);
extern void _ZN6QMovieC1ERK7QStringRK10QByteArrayP7QObject(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  const QMetaObject * QMovie::metaObject();
extern void _ZNK6QMovie10metaObjectEv(void* qthis);
  // proto:  void QMovie::~QMovie();
extern void _ZN6QMovieD0Ev(void* qthis);
  // proto:  void QMovie::start();
extern void _ZN6QMovie5startEv(void* qthis);
  // proto:  int QMovie::loopCount();
extern void _ZNK6QMovie9loopCountEv(void* qthis);
  // proto:  void QMovie::QMovie(QIODevice * device, const QByteArray & format, QObject * parent);
extern void* dector_ZN6QMovieC1EP9QIODeviceRK10QByteArrayP7QObject(void* arg0, void* arg1, void* arg2);
extern void _ZN6QMovieC1EP9QIODeviceRK10QByteArrayP7QObject(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QMovie::setFormat(const QByteArray & format);
extern void _ZN6QMovie9setFormatERK10QByteArray(void* qthis, void* arg0);
  // proto: static QList<QByteArray> QMovie::supportedFormats();
extern void _ZN6QMovie16supportedFormatsEv();
  // proto:  QRect QMovie::frameRect();
extern void _ZNK6QMovie9frameRectEv(void* qthis);
  // proto:  void QMovie::setPaused(bool paused);
extern void _ZN6QMovie9setPausedEb(void* qthis, bool arg0);
  // proto:  QSize QMovie::scaledSize();
extern void _ZN6QMovie10scaledSizeEv(void* qthis);
  // proto:  QIODevice * QMovie::device();
extern void _ZNK6QMovie6deviceEv(void* qthis);
  // proto:  void QMovie::setBackgroundColor(const QColor & color);
extern void _ZN6QMovie18setBackgroundColorERK6QColor(void* qthis, void* arg0);
  // proto:  bool QMovie::isValid();
extern void _ZNK6QMovie7isValidEv(void* qthis);
  // proto:  void QMovie::setSpeed(int percentSpeed);
extern void _ZN6QMovie8setSpeedEi(void* qthis, int arg0);
  // proto:  void QMovie::QMovie(const QMovie & );
extern void* dector_ZN6QMovieC1ERKS_(void* arg0);
extern void _ZN6QMovieC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMovie::stop();
extern void _ZN6QMovie4stopEv(void* qthis);
  // proto:  int QMovie::currentFrameNumber();
extern void _ZNK6QMovie18currentFrameNumberEv(void* qthis);
  // proto:  int QMovie::nextFrameDelay();
extern void _ZNK6QMovie14nextFrameDelayEv(void* qthis);
  // proto:  QPixmap QMovie::currentPixmap();
extern void _ZNK6QMovie13currentPixmapEv(void* qthis);
  // proto:  QByteArray QMovie::format();
extern void _ZNK6QMovie6formatEv(void* qthis);
  // proto:  QString QMovie::fileName();
extern void _ZNK6QMovie8fileNameEv(void* qthis);
  // proto:  QColor QMovie::backgroundColor();
extern void _ZNK6QMovie15backgroundColorEv(void* qthis);
  // proto:  void QMovie::setFileName(const QString & fileName);
extern void _ZN6QMovie11setFileNameERK7QString(void* qthis, void* arg0);
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

// class sizeof(QMovie)=1
type QMovie struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _updated QMovie_updated_signal;
//  _stateChanged QMovie_stateChanged_signal;
//  _started QMovie_started_signal;
//  _resized QMovie_resized_signal;
//  _finished QMovie_finished_signal;
//  _error QMovie_error_signal;
//  _frameChanged QMovie_frameChanged_signal;
}

  // proto:  void QMovie::QMovie(QObject * parent);
func NewQMovie(args ...interface{}) QMovie {
  return QMovie{}
}

  // proto:  int QMovie::speed();
func (this *QMovie) speed(args ...interface{}) () {
  // speed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie5speedEv
  default:
    qtrt.ErrorResolve("QMovie", "speed", args)
  }

}

  // proto:  bool QMovie::jumpToNextFrame();
func (this *QMovie) jumpToNextFrame(args ...interface{}) () {
  // jumpToNextFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie15jumpToNextFrameEv
  default:
    qtrt.ErrorResolve("QMovie", "jumpToNextFrame", args)
  }

}

  // proto:  int QMovie::frameCount();
func (this *QMovie) frameCount(args ...interface{}) () {
  // frameCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie10frameCountEv
  default:
    qtrt.ErrorResolve("QMovie", "frameCount", args)
  }

}

  // proto:  void QMovie::setScaledSize(const QSize & size);
func (this *QMovie) setScaledSize(args ...interface{}) () {
  // setScaledSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie13setScaledSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setScaledSize", args)
  }

}

  // proto:  void QMovie::setDevice(QIODevice * device);
func (this *QMovie) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie9setDeviceEP9QIODevice
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setDevice", args)
  }

}

  // proto:  QImage QMovie::currentImage();
func (this *QMovie) currentImage(args ...interface{}) () {
  // currentImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie12currentImageEv
  default:
    qtrt.ErrorResolve("QMovie", "currentImage", args)
  }

}

  // proto:  bool QMovie::jumpToFrame(int frameNumber);
func (this *QMovie) jumpToFrame(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "jumpToFrame", args)
  }

}

  // proto:  const QMetaObject * QMovie::metaObject();
func (this *QMovie) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie10metaObjectEv
  default:
    qtrt.ErrorResolve("QMovie", "metaObject", args)
  }

}

  // proto:  void QMovie::~QMovie();
func (this *QMovie) FreeQMovie(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMovie", "~QMovie", args)
  }

}

  // proto:  void QMovie::start();
func (this *QMovie) start(args ...interface{}) () {
  // start()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie5startEv
  default:
    qtrt.ErrorResolve("QMovie", "start", args)
  }

}

  // proto:  int QMovie::loopCount();
func (this *QMovie) loopCount(args ...interface{}) () {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie9loopCountEv
  default:
    qtrt.ErrorResolve("QMovie", "loopCount", args)
  }

}

  // proto:  void QMovie::setFormat(const QByteArray & format);
func (this *QMovie) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie9setFormatERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setFormat", args)
  }

}

  // proto: static QList<QByteArray> QMovie::supportedFormats();
func (this *QMovie) supportedFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMovie", "supportedFormats", args)
  }

}

  // proto:  QRect QMovie::frameRect();
func (this *QMovie) frameRect(args ...interface{}) () {
  // frameRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie9frameRectEv
  default:
    qtrt.ErrorResolve("QMovie", "frameRect", args)
  }

}

  // proto:  void QMovie::setPaused(bool paused);
func (this *QMovie) setPaused(args ...interface{}) () {
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setPaused", args)
  }

}

  // proto:  QSize QMovie::scaledSize();
func (this *QMovie) scaledSize(args ...interface{}) () {
  // scaledSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie10scaledSizeEv
  default:
    qtrt.ErrorResolve("QMovie", "scaledSize", args)
  }

}

  // proto:  QIODevice * QMovie::device();
func (this *QMovie) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie6deviceEv
  default:
    qtrt.ErrorResolve("QMovie", "device", args)
  }

}

  // proto:  void QMovie::setBackgroundColor(const QColor & color);
func (this *QMovie) setBackgroundColor(args ...interface{}) () {
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
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setBackgroundColor", args)
  }

}

  // proto:  bool QMovie::isValid();
func (this *QMovie) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie7isValidEv
  default:
    qtrt.ErrorResolve("QMovie", "isValid", args)
  }

}

  // proto:  void QMovie::setSpeed(int percentSpeed);
func (this *QMovie) setSpeed(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setSpeed", args)
  }

}

  // proto:  void QMovie::stop();
func (this *QMovie) stop(args ...interface{}) () {
  // stop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie4stopEv
  default:
    qtrt.ErrorResolve("QMovie", "stop", args)
  }

}

  // proto:  int QMovie::currentFrameNumber();
func (this *QMovie) currentFrameNumber(args ...interface{}) () {
  // currentFrameNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie18currentFrameNumberEv
  default:
    qtrt.ErrorResolve("QMovie", "currentFrameNumber", args)
  }

}

  // proto:  int QMovie::nextFrameDelay();
func (this *QMovie) nextFrameDelay(args ...interface{}) () {
  // nextFrameDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie14nextFrameDelayEv
  default:
    qtrt.ErrorResolve("QMovie", "nextFrameDelay", args)
  }

}

  // proto:  QPixmap QMovie::currentPixmap();
func (this *QMovie) currentPixmap(args ...interface{}) () {
  // currentPixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie13currentPixmapEv
  default:
    qtrt.ErrorResolve("QMovie", "currentPixmap", args)
  }

}

  // proto:  QByteArray QMovie::format();
func (this *QMovie) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie6formatEv
  default:
    qtrt.ErrorResolve("QMovie", "format", args)
  }

}

  // proto:  QString QMovie::fileName();
func (this *QMovie) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie8fileNameEv
  default:
    qtrt.ErrorResolve("QMovie", "fileName", args)
  }

}

  // proto:  QColor QMovie::backgroundColor();
func (this *QMovie) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QMovie15backgroundColorEv
  default:
    qtrt.ErrorResolve("QMovie", "backgroundColor", args)
  }

}

  // proto:  void QMovie::setFileName(const QString & fileName);
func (this *QMovie) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QMovie11setFileNameERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMovie", "setFileName", args)
  }

}

// <= body block end

