package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qimage.h
// dst-file: /src/gui/qimage.go
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
  // proto:  QImage QImage::copy(const QRect & rect);
extern void _ZNK6QImage4copyERK5QRect(void* qthis, void* arg0);
  // proto: static QTransform QImage::trueMatrix(const QTransform & , int w, int h);
extern void _ZN6QImage10trueMatrixERK10QTransformii(void* arg0, int arg1, int arg2);
  // proto:  uchar * QImage::bits();
extern void _ZN6QImage4bitsEv(void* qthis);
  // proto:  void QImage::setAlphaChannel(const QImage & alphaChannel);
extern void _ZN6QImage15setAlphaChannelERKS_(void* qthis, void* arg0);
  // proto:  QString QImage::text(const QString & key);
extern void _ZNK6QImage4textERK7QString(void* qthis, void* arg0);
  // proto:  QRect QImage::rect();
extern void _ZNK6QImage4rectEv(void* qthis);
  // proto:  void QImage::QImage(const char *const [] xpm);
extern void* dector_ZN6QImageC1EPKPKc(char** arg0);
extern void _ZN6QImageC1EPKPKc(void* qthis, char** arg0);
  // proto:  QImage QImage::createHeuristicMask(bool clipTight);
extern void _ZNK6QImage19createHeuristicMaskEb(void* qthis, bool arg0);
  // proto:  const uchar * QImage::constBits();
extern void _ZNK6QImage9constBitsEv(void* qthis);
  // proto: static QImage QImage::fromData(const QByteArray & data, const char * format);
extern void demth_ZN6QImage8fromDataERK10QByteArrayPKc(void* arg0, char* arg1);
  // proto: static QImage QImage::fromData(const uchar * data, int size, const char * format);
extern void _ZN6QImage8fromDataEPKhiPKc(unsigned char* arg0, int arg1, char* arg2);
  // proto:  bool QImage::isDetached();
extern void _ZNK6QImage10isDetachedEv(void* qthis);
  // proto:  void QImage::setOffset(const QPoint & );
extern void _ZN6QImage9setOffsetERK6QPoint(void* qthis, void* arg0);
  // proto: static QMatrix QImage::trueMatrix(const QMatrix & , int w, int h);
extern void _ZN6QImage10trueMatrixERK7QMatrixii(void* arg0, int arg1, int arg2);
  // proto:  bool QImage::isGrayscale();
extern void _ZNK6QImage11isGrayscaleEv(void* qthis);
  // proto:  bool QImage::save(QIODevice * device, const char * format, int quality);
extern void _ZNK6QImage4saveEP9QIODevicePKci(void* qthis, void* arg0, char* arg1, int arg2);
  // proto:  int QImage::depth();
extern void _ZNK6QImage5depthEv(void* qthis);
  // proto:  QImage QImage::alphaChannel();
extern void _ZNK6QImage12alphaChannelEv(void* qthis);
  // proto:  bool QImage::hasAlphaChannel();
extern void _ZNK6QImage15hasAlphaChannelEv(void* qthis);
  // proto:  bool QImage::loadFromData(const uchar * buf, int len, const char * format);
extern void _ZN6QImage12loadFromDataEPKhiPKc(void* qthis, unsigned char* arg0, int arg1, char* arg2);
  // proto:  int QImage::colorCount();
extern void _ZNK6QImage10colorCountEv(void* qthis);
  // proto:  bool QImage::allGray();
extern void _ZNK6QImage7allGrayEv(void* qthis);
  // proto:  void QImage::setColorCount(int );
extern void _ZN6QImage13setColorCountEi(void* qthis, int arg0);
  // proto:  QRgb QImage::pixel(const QPoint & pt);
extern void demth_ZNK6QImage5pixelERK6QPoint(void* qthis, void* arg0);
  // proto:  void QImage::setDevicePixelRatio(qreal scaleFactor);
extern void _ZN6QImage19setDevicePixelRatioEd(void* qthis, double arg0);
  // proto:  QImage QImage::copy(int x, int y, int w, int h);
extern void demth_ZNK6QImage4copyEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QImage::setText(const QString & key, const QString & value);
extern void _ZN6QImage7setTextERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  QRgb QImage::color(int i);
extern void _ZNK6QImage5colorEi(void* qthis, int arg0);
  // proto:  void QImage::setPixel(const QPoint & pt, uint index_or_rgb);
extern void demth_ZN6QImage8setPixelERK6QPointj(void* qthis, void* arg0, unsigned int arg1);
  // proto:  QPoint QImage::offset();
extern void _ZNK6QImage6offsetEv(void* qthis);
  // proto:  const uchar * QImage::constScanLine(int );
extern void _ZNK6QImage13constScanLineEi(void* qthis, int arg0);
  // proto:  QStringList QImage::textKeys();
extern void _ZNK6QImage8textKeysEv(void* qthis);
  // proto:  int QImage::dotsPerMeterY();
extern void _ZNK6QImage13dotsPerMeterYEv(void* qthis);
  // proto:  void QImage::fill(uint pixel);
extern void _ZN6QImage4fillEj(void* qthis, unsigned int arg0);
  // proto:  QPixelFormat QImage::pixelFormat();
extern void _ZNK6QImage11pixelFormatEv(void* qthis);
  // proto:  int QImage::dotsPerMeterX();
extern void _ZNK6QImage13dotsPerMeterXEv(void* qthis);
  // proto:  void QImage::setDotsPerMeterY(int );
extern void _ZN6QImage16setDotsPerMeterYEi(void* qthis, int arg0);
  // proto:  int QImage::bitPlaneCount();
extern void _ZNK6QImage13bitPlaneCountEv(void* qthis);
  // proto:  void QImage::fill(const QColor & color);
extern void _ZN6QImage4fillERK6QColor(void* qthis, void* arg0);
  // proto:  void QImage::detach();
extern void _ZN6QImage6detachEv(void* qthis);
  // proto:  bool QImage::loadFromData(const QByteArray & data, const char * aformat);
extern void demth_ZN6QImage12loadFromDataERK10QByteArrayPKc(void* qthis, void* arg0, char* arg1);
  // proto:  void QImage::QImage(const QString & fileName, const char * format);
extern void* dector_ZN6QImageC1ERK7QStringPKc(void* arg0, char* arg1);
extern void _ZN6QImageC1ERK7QStringPKc(void* qthis, void* arg0, char* arg1);
  // proto:  QPaintEngine * QImage::paintEngine();
extern void _ZNK6QImage11paintEngineEv(void* qthis);
  // proto:  void QImage::QImage(const QImage & );
extern void* dector_ZN6QImageC1ERKS_(void* arg0);
extern void _ZN6QImageC1ERKS_(void* qthis, void* arg0);
  // proto:  void QImage::swap(QImage & other);
extern void demth_ZN6QImage4swapERS_(void* qthis, void* arg0);
  // proto:  qreal QImage::devicePixelRatio();
extern void _ZNK6QImage16devicePixelRatioEv(void* qthis);
  // proto:  int QImage::devType();
extern void _ZNK6QImage7devTypeEv(void* qthis);
  // proto:  bool QImage::valid(const QPoint & pt);
extern void demth_ZNK6QImage5validERK6QPoint(void* qthis, void* arg0);
  // proto:  int QImage::pixelIndex(const QPoint & pt);
extern void demth_ZNK6QImage10pixelIndexERK6QPoint(void* qthis, void* arg0);
  // proto:  void QImage::setDotsPerMeterX(int );
extern void _ZN6QImage16setDotsPerMeterXEi(void* qthis, int arg0);
  // proto:  void QImage::setPixel(int x, int y, uint index_or_rgb);
extern void _ZN6QImage8setPixelEiij(void* qthis, int arg0, int arg1, unsigned int arg2);
  // proto:  bool QImage::load(const QString & fileName, const char * format);
extern void _ZN6QImage4loadERK7QStringPKc(void* qthis, void* arg0, char* arg1);
  // proto:  QVector<QRgb> QImage::colorTable();
extern void _ZNK6QImage10colorTableEv(void* qthis);
  // proto:  QSize QImage::size();
extern void _ZNK6QImage4sizeEv(void* qthis);
  // proto:  int QImage::height();
extern void _ZNK6QImage6heightEv(void* qthis);
  // proto:  int QImage::pixelIndex(int x, int y);
extern void _ZNK6QImage10pixelIndexEii(void* qthis, int arg0, int arg1);
  // proto:  int QImage::width();
extern void _ZNK6QImage5widthEv(void* qthis);
  // proto:  bool QImage::load(QIODevice * device, const char * format);
extern void _ZN6QImage4loadEP9QIODevicePKc(void* qthis, void* arg0, char* arg1);
  // proto:  void QImage::QImage();
extern void* dector_ZN6QImageC1Ev();
extern void _ZN6QImageC1Ev(void* qthis);
  // proto:  uchar * QImage::scanLine(int );
extern void _ZN6QImage8scanLineEi(void* qthis, int arg0);
  // proto:  int QImage::bytesPerLine();
extern void _ZNK6QImage12bytesPerLineEv(void* qthis);
  // proto:  qint64 QImage::cacheKey();
extern void _ZNK6QImage8cacheKeyEv(void* qthis);
  // proto:  QRgb QImage::pixel(int x, int y);
extern void _ZNK6QImage5pixelEii(void* qthis, int arg0, int arg1);
  // proto:  void QImage::~QImage();
extern void _ZN6QImageD0Ev(void* qthis);
  // proto:  bool QImage::save(const QString & fileName, const char * format, int quality);
extern void _ZNK6QImage4saveERK7QStringPKci(void* qthis, void* arg0, char* arg1, int arg2);
  // proto:  void QImage::setColor(int i, QRgb c);
extern void _ZN6QImage8setColorEij(void* qthis, int arg0, unsigned int arg1);
  // proto:  bool QImage::isNull();
extern void _ZNK6QImage6isNullEv(void* qthis);
  // proto:  int QImage::byteCount();
extern void _ZNK6QImage9byteCountEv(void* qthis);
  // proto:  bool QImage::valid(int x, int y);
extern void _ZNK6QImage5validEii(void* qthis, int arg0, int arg1);
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

// class sizeof(QImage)=32
type QImage struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QImage QImage::copy(const QRect & rect);
func (this *QImage) copy(args ...interface{}) () {
  // copy(const class QRect &)
  // copy(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4copyERK5QRect
    // invoke: QImage copy(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QImage4copyERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK6QImage4copyEiiii
    // invoke: QImage copy(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZNK6QImage4copyEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QImage", "copy", args)
  }

}

  // proto: static QTransform QImage::trueMatrix(const QTransform & , int w, int h);
func (this *QImage) trueMatrix_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImage", "trueMatrix", args)
  }

}

  // proto:  uchar * QImage::bits();
func (this *QImage) bits(args ...interface{}) () {
  // bits()
  // bits()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4bitsEv
    // invoke: uchar * bits()
    C._ZN6QImage4bitsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "bits", args)
  }

}

  // proto:  void QImage::setAlphaChannel(const QImage & alphaChannel);
func (this *QImage) setAlphaChannel(args ...interface{}) () {
  // setAlphaChannel(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage15setAlphaChannelERKS_
    // invoke: void setAlphaChannel(const class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QImage15setAlphaChannelERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setAlphaChannel", args)
  }

}

  // proto:  QString QImage::text(const QString & key);
func (this *QImage) text(args ...interface{}) () {
  // text(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4textERK7QString
    // invoke: QString text(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QImage4textERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "text", args)
  }

}

  // proto:  QRect QImage::rect();
func (this *QImage) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4rectEv
    // invoke: QRect rect()
    C._ZNK6QImage4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "rect", args)
  }

}

  // proto:  void QImage::QImage(const char *const [] xpm);
func NewQImage(args ...interface{}) QImage {
  return QImage{}
}

  // proto:  QImage QImage::createHeuristicMask(bool clipTight);
func (this *QImage) createHeuristicMask(args ...interface{}) () {
  // createHeuristicMask(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage19createHeuristicMaskEb
    // invoke: QImage createHeuristicMask(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZNK6QImage19createHeuristicMaskEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "createHeuristicMask", args)
  }

}

  // proto:  const uchar * QImage::constBits();
func (this *QImage) constBits(args ...interface{}) () {
  // constBits()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage9constBitsEv
    // invoke: const uchar * constBits()
    C._ZNK6QImage9constBitsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "constBits", args)
  }

}

  // proto: static QImage QImage::fromData(const QByteArray & data, const char * format);
func (this *QImage) fromData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImage", "fromData", args)
  }

}

  // proto:  bool QImage::isDetached();
func (this *QImage) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10isDetachedEv
    // invoke: bool isDetached()
    C._ZNK6QImage10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "isDetached", args)
  }

}

  // proto:  void QImage::setOffset(const QPoint & );
func (this *QImage) setOffset(args ...interface{}) () {
  // setOffset(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage9setOffsetERK6QPoint
    // invoke: void setOffset(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QImage9setOffsetERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setOffset", args)
  }

}

  // proto:  bool QImage::isGrayscale();
func (this *QImage) isGrayscale(args ...interface{}) () {
  // isGrayscale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage11isGrayscaleEv
    // invoke: bool isGrayscale()
    C._ZNK6QImage11isGrayscaleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "isGrayscale", args)
  }

}

  // proto:  bool QImage::save(QIODevice * device, const char * format, int quality);
func (this *QImage) save(args ...interface{}) () {
  // save(class QIODevice *, const char *, int)
  // save(const class QString &, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4saveEP9QIODevicePKci
    // invoke: bool save(class QIODevice *, const char *, int)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK6QImage4saveEP9QIODevicePKci(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK6QImage4saveERK7QStringPKci
    // invoke: bool save(const class QString &, const char *, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK6QImage4saveERK7QStringPKci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QImage", "save", args)
  }

}

  // proto:  int QImage::depth();
func (this *QImage) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5depthEv
    // invoke: int depth()
    C._ZNK6QImage5depthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "depth", args)
  }

}

  // proto:  QImage QImage::alphaChannel();
func (this *QImage) alphaChannel(args ...interface{}) () {
  // alphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage12alphaChannelEv
    // invoke: QImage alphaChannel()
    C._ZNK6QImage12alphaChannelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "alphaChannel", args)
  }

}

  // proto:  bool QImage::hasAlphaChannel();
func (this *QImage) hasAlphaChannel(args ...interface{}) () {
  // hasAlphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage15hasAlphaChannelEv
    // invoke: bool hasAlphaChannel()
    C._ZNK6QImage15hasAlphaChannelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "hasAlphaChannel", args)
  }

}

  // proto:  bool QImage::loadFromData(const uchar * buf, int len, const char * format);
func (this *QImage) loadFromData(args ...interface{}) () {
  // loadFromData(const uchar *, int, const char *)
  // loadFromData(const class QByteArray &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage12loadFromDataEPKhiPKc
    // invoke: bool loadFromData(const uchar *, int, const char *)
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.CString(args[2].(string))
    if false {fmt.Println(arg2)}
    C._ZN6QImage12loadFromDataEPKhiPKc(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN6QImage12loadFromDataERK10QByteArrayPKc
    // invoke: bool loadFromData(const class QByteArray &, const char *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C.demth_ZN6QImage12loadFromDataERK10QByteArrayPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "loadFromData", args)
  }

}

  // proto:  int QImage::colorCount();
func (this *QImage) colorCount(args ...interface{}) () {
  // colorCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10colorCountEv
    // invoke: int colorCount()
    C._ZNK6QImage10colorCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "colorCount", args)
  }

}

  // proto:  bool QImage::allGray();
func (this *QImage) allGray(args ...interface{}) () {
  // allGray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage7allGrayEv
    // invoke: bool allGray()
    C._ZNK6QImage7allGrayEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "allGray", args)
  }

}

  // proto:  void QImage::setColorCount(int );
func (this *QImage) setColorCount(args ...interface{}) () {
  // setColorCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage13setColorCountEi
    // invoke: void setColorCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QImage13setColorCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setColorCount", args)
  }

}

  // proto:  QRgb QImage::pixel(const QPoint & pt);
func (this *QImage) pixel(args ...interface{}) () {
  // pixel(const class QPoint &)
  // pixel(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5pixelERK6QPoint
    // invoke: QRgb pixel(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK6QImage5pixelERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK6QImage5pixelEii
    // invoke: QRgb pixel(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK6QImage5pixelEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "pixel", args)
  }

}

  // proto:  void QImage::setDevicePixelRatio(qreal scaleFactor);
func (this *QImage) setDevicePixelRatio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage19setDevicePixelRatioEd
    // invoke: void setDevicePixelRatio(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN6QImage19setDevicePixelRatioEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setDevicePixelRatio", args)
  }

}

  // proto:  void QImage::setText(const QString & key, const QString & value);
func (this *QImage) setText(args ...interface{}) () {
  // setText(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage7setTextERK7QStringS2_
    // invoke: void setText(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN6QImage7setTextERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "setText", args)
  }

}

  // proto:  QRgb QImage::color(int i);
func (this *QImage) color(args ...interface{}) () {
  // color(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5colorEi
    // invoke: QRgb color(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK6QImage5colorEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "color", args)
  }

}

  // proto:  void QImage::setPixel(const QPoint & pt, uint index_or_rgb);
func (this *QImage) setPixel(args ...interface{}) () {
  // setPixel(const class QPoint &, uint)
  // setPixel(int, int, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8setPixelERK6QPointj
    // invoke: void setPixel(const class QPoint &, uint)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN6QImage8setPixelERK6QPointj(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN6QImage8setPixelEiij
    // invoke: void setPixel(int, int, uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN6QImage8setPixelEiij(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QImage", "setPixel", args)
  }

}

  // proto:  QPoint QImage::offset();
func (this *QImage) offset(args ...interface{}) () {
  // offset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6offsetEv
    // invoke: QPoint offset()
    C._ZNK6QImage6offsetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "offset", args)
  }

}

  // proto:  const uchar * QImage::constScanLine(int );
func (this *QImage) constScanLine(args ...interface{}) () {
  // constScanLine(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13constScanLineEi
    // invoke: const uchar * constScanLine(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK6QImage13constScanLineEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "constScanLine", args)
  }

}

  // proto:  QStringList QImage::textKeys();
func (this *QImage) textKeys(args ...interface{}) () {
  // textKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage8textKeysEv
    // invoke: QStringList textKeys()
    C._ZNK6QImage8textKeysEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "textKeys", args)
  }

}

  // proto:  int QImage::dotsPerMeterY();
func (this *QImage) dotsPerMeterY(args ...interface{}) () {
  // dotsPerMeterY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13dotsPerMeterYEv
    // invoke: int dotsPerMeterY()
    C._ZNK6QImage13dotsPerMeterYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "dotsPerMeterY", args)
  }

}

  // proto:  void QImage::fill(uint pixel);
func (this *QImage) fill(args ...interface{}) () {
  // fill(uint)
  // fill(const class QColor &)
  // fill(Qt::GlobalColor)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "Qt::GlobalColor"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4fillEj
    // invoke: void fill(uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QImage4fillEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QImage4fillERK6QColor
    // invoke: void fill(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QImage4fillERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "fill", args)
  }

}

  // proto:  QPixelFormat QImage::pixelFormat();
func (this *QImage) pixelFormat(args ...interface{}) () {
  // pixelFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage11pixelFormatEv
    // invoke: QPixelFormat pixelFormat()
    C._ZNK6QImage11pixelFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "pixelFormat", args)
  }

}

  // proto:  int QImage::dotsPerMeterX();
func (this *QImage) dotsPerMeterX(args ...interface{}) () {
  // dotsPerMeterX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13dotsPerMeterXEv
    // invoke: int dotsPerMeterX()
    C._ZNK6QImage13dotsPerMeterXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "dotsPerMeterX", args)
  }

}

  // proto:  void QImage::setDotsPerMeterY(int );
func (this *QImage) setDotsPerMeterY(args ...interface{}) () {
  // setDotsPerMeterY(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage16setDotsPerMeterYEi
    // invoke: void setDotsPerMeterY(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QImage16setDotsPerMeterYEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setDotsPerMeterY", args)
  }

}

  // proto:  int QImage::bitPlaneCount();
func (this *QImage) bitPlaneCount(args ...interface{}) () {
  // bitPlaneCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage13bitPlaneCountEv
    // invoke: int bitPlaneCount()
    C._ZNK6QImage13bitPlaneCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "bitPlaneCount", args)
  }

}

  // proto:  void QImage::detach();
func (this *QImage) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage6detachEv
    // invoke: void detach()
    C._ZN6QImage6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "detach", args)
  }

}

  // proto:  QPaintEngine * QImage::paintEngine();
func (this *QImage) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    C._ZNK6QImage11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "paintEngine", args)
  }

}

  // proto:  void QImage::swap(QImage & other);
func (this *QImage) swap(args ...interface{}) () {
  // swap(class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4swapERS_
    // invoke: void swap(class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN6QImage4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "swap", args)
  }

}

  // proto:  qreal QImage::devicePixelRatio();
func (this *QImage) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    C._ZNK6QImage16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "devicePixelRatio", args)
  }

}

  // proto:  int QImage::devType();
func (this *QImage) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage7devTypeEv
    // invoke: int devType()
    C._ZNK6QImage7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "devType", args)
  }

}

  // proto:  bool QImage::valid(const QPoint & pt);
func (this *QImage) valid(args ...interface{}) () {
  // valid(const class QPoint &)
  // valid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5validERK6QPoint
    // invoke: bool valid(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK6QImage5validERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK6QImage5validEii
    // invoke: bool valid(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK6QImage5validEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "valid", args)
  }

}

  // proto:  int QImage::pixelIndex(const QPoint & pt);
func (this *QImage) pixelIndex(args ...interface{}) () {
  // pixelIndex(const class QPoint &)
  // pixelIndex(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10pixelIndexERK6QPoint
    // invoke: int pixelIndex(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK6QImage10pixelIndexERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK6QImage10pixelIndexEii
    // invoke: int pixelIndex(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK6QImage10pixelIndexEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "pixelIndex", args)
  }

}

  // proto:  void QImage::setDotsPerMeterX(int );
func (this *QImage) setDotsPerMeterX(args ...interface{}) () {
  // setDotsPerMeterX(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage16setDotsPerMeterXEi
    // invoke: void setDotsPerMeterX(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QImage16setDotsPerMeterXEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setDotsPerMeterX", args)
  }

}

  // proto:  bool QImage::load(const QString & fileName, const char * format);
func (this *QImage) load(args ...interface{}) () {
  // load(const class QString &, const char *)
  // load(class QIODevice *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4loadERK7QStringPKc
    // invoke: bool load(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN6QImage4loadERK7QStringPKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN6QImage4loadEP9QIODevicePKc
    // invoke: bool load(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN6QImage4loadEP9QIODevicePKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "load", args)
  }

}

  // proto:  QVector<QRgb> QImage::colorTable();
func (this *QImage) colorTable(args ...interface{}) () {
  // colorTable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10colorTableEv
    // invoke: QVector<QRgb> colorTable()
    C._ZNK6QImage10colorTableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "colorTable", args)
  }

}

  // proto:  QSize QImage::size();
func (this *QImage) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4sizeEv
    // invoke: QSize size()
    C._ZNK6QImage4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "size", args)
  }

}

  // proto:  int QImage::height();
func (this *QImage) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6heightEv
    // invoke: int height()
    C._ZNK6QImage6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "height", args)
  }

}

  // proto:  int QImage::width();
func (this *QImage) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5widthEv
    // invoke: int width()
    C._ZNK6QImage5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "width", args)
  }

}

  // proto:  uchar * QImage::scanLine(int );
func (this *QImage) scanLine(args ...interface{}) () {
  // scanLine(int)
  // scanLine(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8scanLineEi
    // invoke: uchar * scanLine(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN6QImage8scanLineEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "scanLine", args)
  }

}

  // proto:  int QImage::bytesPerLine();
func (this *QImage) bytesPerLine(args ...interface{}) () {
  // bytesPerLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage12bytesPerLineEv
    // invoke: int bytesPerLine()
    C._ZNK6QImage12bytesPerLineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "bytesPerLine", args)
  }

}

  // proto:  qint64 QImage::cacheKey();
func (this *QImage) cacheKey(args ...interface{}) () {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage8cacheKeyEv
    // invoke: qint64 cacheKey()
    C._ZNK6QImage8cacheKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "cacheKey", args)
  }

}

  // proto:  void QImage::~QImage();
func (this *QImage) FreeQImage(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImage", "~QImage", args)
  }

}

  // proto:  void QImage::setColor(int i, QRgb c);
func (this *QImage) setColor(args ...interface{}) () {
  // setColor(int, QRgb)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "QRgb"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8setColorEij
    // invoke: void setColor(int, QRgb)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN6QImage8setColorEij(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "setColor", args)
  }

}

  // proto:  bool QImage::isNull();
func (this *QImage) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6isNullEv
    // invoke: bool isNull()
    C._ZNK6QImage6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "isNull", args)
  }

}

  // proto:  int QImage::byteCount();
func (this *QImage) byteCount(args ...interface{}) () {
  // byteCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage9byteCountEv
    // invoke: int byteCount()
    C._ZNK6QImage9byteCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "byteCount", args)
  }

}

// <= body block end

