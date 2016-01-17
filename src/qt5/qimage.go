package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QImage::load(QIODevice * device, const char * format);
extern void _ZN6QImage4loadEP9QIODevicePKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  bool QImage::load(const QString & fileName, const char * format);
extern void _ZN6QImage4loadERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto: static QImage::Format QImage::toImageFormat(QPixelFormat format);
extern void _ZN6QImage13toImageFormatE12QPixelFormat(void* arg0); // 4
  // proto:  QRgb QImage::color(int i);
extern void _ZNK6QImage5colorEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QImage::text(const QString & key);
extern void _ZNK6QImage4textERK7QString(void* qthis, void* arg0); // 4
  // proto:  QPoint QImage::offset();
extern void _ZNK6QImage6offsetEv(void* qthis); // 4
  // proto:  const uchar * QImage::constScanLine(int );
extern void _ZNK6QImage13constScanLineEi(void* qthis, int32_t arg0); // 4
  // proto:  int QImage::height();
extern void _ZNK6QImage6heightEv(void* qthis); // 4
  // proto: static QImage QImage::fromData(const uchar * data, int size, const char * format);
extern void _ZN6QImage8fromDataEPKhiPKc(unsigned char* arg0, int32_t arg1, unsigned char* arg2); // 4
  // proto: static QImage QImage::fromData(const QByteArray & data, const char * format);
extern void _ZN6QImage8fromDataERK10QByteArrayPKc(void* arg0, unsigned char* arg1); // 2
  // proto:  uchar * QImage::scanLine(int );
extern void _ZN6QImage8scanLineEi(void* qthis, int32_t arg0); // 4
  // proto:  int QImage::byteCount();
extern void _ZNK6QImage9byteCountEv(void* qthis); // 4
  // proto:  qint64 QImage::cacheKey();
extern void _ZNK6QImage8cacheKeyEv(void* qthis); // 4
  // proto:  void QImage::setColorCount(int );
extern void _ZN6QImage13setColorCountEi(void* qthis, int32_t arg0); // 4
  // proto:  const uchar * QImage::constBits();
extern void _ZNK6QImage9constBitsEv(void* qthis); // 4
  // proto: static QMatrix QImage::trueMatrix(const QMatrix & , int w, int h);
extern void _ZN6QImage10trueMatrixERK7QMatrixii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto: static QTransform QImage::trueMatrix(const QTransform & , int w, int h);
extern void _ZN6QImage10trueMatrixERK10QTransformii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QImage::detach();
extern void _ZN6QImage6detachEv(void* qthis); // 4
  // proto:  QPaintEngine * QImage::paintEngine();
extern void _ZNK6QImage11paintEngineEv(void* qthis); // 4
  // proto:  QSize QImage::size();
extern void _ZNK6QImage4sizeEv(void* qthis); // 4
  // proto:  int QImage::devType();
extern void _ZNK6QImage7devTypeEv(void* qthis); // 4
  // proto:  QImage QImage::mirrored(bool horizontally, bool vertically);
extern void _ZNKR6QImage8mirroredEbb(void* qthis, bool arg0, bool arg1); // 2
  // proto:  int QImage::width();
extern void _ZNK6QImage5widthEv(void* qthis); // 4
  // proto:  bool QImage::valid(int x, int y);
extern void _ZNK6QImage5validEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QImage::valid(const QPoint & pt);
extern void _ZNK6QImage5validERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QImage::setDevicePixelRatio(qreal scaleFactor);
extern void _ZN6QImage19setDevicePixelRatioEd(void* qthis, double arg0); // 4
  // proto:  QStringList QImage::textKeys();
extern void _ZNK6QImage8textKeysEv(void* qthis); // 4
  // proto:  void QImage::setOffset(const QPoint & );
extern void _ZN6QImage9setOffsetERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QImage::save(const QString & fileName, const char * format, int quality);
extern void _ZNK6QImage4saveERK7QStringPKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2); // 4
  // proto:  bool QImage::save(QIODevice * device, const char * format, int quality);
extern void _ZNK6QImage4saveEP9QIODevicePKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2); // 4
  // proto:  uchar * QImage::bits();
extern void _ZN6QImage4bitsEv(void* qthis); // 4
  // proto:  QRgb QImage::pixel(const QPoint & pt);
extern void _ZNK6QImage5pixelERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QRgb QImage::pixel(int x, int y);
extern void _ZNK6QImage5pixelEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QImage::swap(QImage & other);
extern void _ZN6QImage4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QImage::isDetached();
extern void _ZNK6QImage10isDetachedEv(void* qthis); // 4
  // proto:  QImage QImage::alphaChannel();
extern void _ZNK6QImage12alphaChannelEv(void* qthis); // 4
  // proto:  void QImage::setPixel(int x, int y, uint index_or_rgb);
extern void _ZN6QImage8setPixelEiij(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QImage::setPixel(const QPoint & pt, uint index_or_rgb);
extern void _ZN6QImage8setPixelERK6QPointj(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QImage::setDotsPerMeterY(int );
extern void _ZN6QImage16setDotsPerMeterYEi(void* qthis, int32_t arg0); // 4
  // proto:  QImage::Format QImage::format();
extern void _ZNK6QImage6formatEv(void* qthis); // 4
  // proto:  bool QImage::hasAlphaChannel();
extern void _ZNK6QImage15hasAlphaChannelEv(void* qthis); // 4
  // proto:  int QImage::bitPlaneCount();
extern void _ZNK6QImage13bitPlaneCountEv(void* qthis); // 4
  // proto:  QPixelFormat QImage::pixelFormat();
extern void _ZNK6QImage11pixelFormatEv(void* qthis); // 4
  // proto:  QImage QImage::rgbSwapped();
extern void _ZNKR6QImage10rgbSwappedEv(void* qthis); // 2
  // proto:  bool QImage::isGrayscale();
extern void _ZNK6QImage11isGrayscaleEv(void* qthis); // 4
  // proto:  QImage QImage::createHeuristicMask(bool clipTight);
extern void _ZNK6QImage19createHeuristicMaskEb(void* qthis, bool arg0); // 4
  // proto:  int QImage::bytesPerLine();
extern void _ZNK6QImage12bytesPerLineEv(void* qthis); // 4
  // proto:  QImage QImage::copy(const QRect & rect);
extern void _ZNK6QImage4copyERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QImage QImage::copy(int x, int y, int w, int h);
extern void _ZNK6QImage4copyEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QImage::dotsPerMeterY();
extern void _ZNK6QImage13dotsPerMeterYEv(void* qthis); // 4
  // proto:  int QImage::dotsPerMeterX();
extern void _ZNK6QImage13dotsPerMeterXEv(void* qthis); // 4
  // proto:  QRect QImage::rect();
extern void _ZNK6QImage4rectEv(void* qthis); // 4
  // proto:  void QImage::setDotsPerMeterX(int );
extern void _ZN6QImage16setDotsPerMeterXEi(void* qthis, int32_t arg0); // 4
  // proto:  void QImage::setAlphaChannel(const QImage & alphaChannel);
extern void _ZN6QImage15setAlphaChannelERKS_(void* qthis, void* arg0); // 4
  // proto:  void QImage::setColor(int i, QRgb c);
extern void _ZN6QImage8setColorEij(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  qreal QImage::devicePixelRatio();
extern void _ZNK6QImage16devicePixelRatioEv(void* qthis); // 4
  // proto:  bool QImage::loadFromData(const uchar * buf, int len, const char * format);
extern void _ZN6QImage12loadFromDataEPKhiPKc(void* qthis, unsigned char* arg0, int32_t arg1, unsigned char* arg2); // 4
  // proto:  bool QImage::loadFromData(const QByteArray & data, const char * aformat);
extern void _ZN6QImage12loadFromDataERK10QByteArrayPKc(void* qthis, void* arg0, unsigned char* arg1); // 2
  // proto:  void QImage::setText(const QString & key, const QString & value);
extern void _ZN6QImage7setTextERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QImage::allGray();
extern void _ZNK6QImage7allGrayEv(void* qthis); // 4
  // proto:  int QImage::colorCount();
extern void _ZNK6QImage10colorCountEv(void* qthis); // 4
  // proto:  bool QImage::isNull();
extern void _ZNK6QImage6isNullEv(void* qthis); // 4
  // proto:  int QImage::depth();
extern void _ZNK6QImage5depthEv(void* qthis); // 4
  // proto:  void QImage::~QImage();
extern void _ZN6QImageD2Ev(void* qthis); // 4
  // proto:  int QImage::pixelIndex(int x, int y);
extern void _ZNK6QImage10pixelIndexEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QImage::pixelIndex(const QPoint & pt);
extern void _ZNK6QImage10pixelIndexERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QImage::fill(uint pixel);
extern void _ZN6QImage4fillEj(void* qthis, int32_t arg0); // 4
  // proto:  void QImage::fill(const QColor & color);
extern void _ZN6QImage4fillERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QVector<QRgb> QImage::colorTable();
extern void _ZNK6QImage10colorTableEv(void* qthis); // 4
  // proto:  void QImage::QImage();
extern void _ZN6QImageC2Ev(void* qthis); // 3
  // proto:  void QImage::QImage(const QString & fileName, const char * format);
extern void _ZN6QImageC2ERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1); // 3
  // proto:  void QImage::QImage(const QImage & );
extern void _ZN6QImageC2ERKS_(void* qthis, void* arg0); // 3
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// load(class QIODevice *, const char *)
func (this *QImage) load(args ...interface{}) () {
  // load(class QIODevice *, const char *)
  // load(const class QString &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4loadEP9QIODevicePKc
    // invoke: bool load(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN6QImage4loadEP9QIODevicePKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN6QImage4loadERK7QStringPKc
    // invoke: bool load(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN6QImage4loadERK7QStringPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "load", args)
  }

}

// toImageFormat(class QPixelFormat)
func (this *QImage) toImageFormat_s(args ...interface{}) () {
  // toImageFormat(class QPixelFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixelFormat{}) // "QPixelFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage13toImageFormatE12QPixelFormat
    // invoke: QImage::Format toImageFormat(class QPixelFormat)
    var arg0 = args[0].(QPixelFormat).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN6QImage13toImageFormatE12QPixelFormat(arg0)
  default:
    qtrt.ErrorResolve("QImage", "toImageFormat", args)
  }

}

// color(int)
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

// text(const class QString &)
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

// offset()
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

// constScanLine(int)
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

// height()
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

// fromData(const uchar *, int, const char *)
func (this *QImage) fromData_s(args ...interface{}) () {
  // fromData(const uchar *, int, const char *)
  // fromData(const class QByteArray &, const char *)
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
    // invoke: _ZN6QImage8fromDataEPKhiPKc
    // invoke: QImage fromData(const uchar *, int, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C._ZN6QImage8fromDataEPKhiPKc(arg0, arg1, arg2)
  case 1:
    // invoke: _ZN6QImage8fromDataERK10QByteArrayPKc
    // invoke: QImage fromData(const class QByteArray &, const char *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN6QImage8fromDataERK10QByteArrayPKc(arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "fromData", args)
  }

}

// scanLine(int)
func (this *QImage) scanLine(args ...interface{}) () {
  // scanLine(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

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

// byteCount()
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

// cacheKey()
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

// setColorCount(int)
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

// constBits()
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

// trueMatrix(const class QMatrix &, int, int)
func (this *QImage) trueMatrix_s(args ...interface{}) () {
  // trueMatrix(const class QMatrix &, int, int)
  // trueMatrix(const class QTransform &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMatrix{}) // "const QMatrix &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage10trueMatrixERK7QMatrixii
    // invoke: QMatrix trueMatrix(const class QMatrix &, int, int)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN6QImage10trueMatrixERK7QMatrixii(arg0, arg1, arg2)
  case 1:
    // invoke: _ZN6QImage10trueMatrixERK10QTransformii
    // invoke: QTransform trueMatrix(const class QTransform &, int, int)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN6QImage10trueMatrixERK10QTransformii(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QImage", "trueMatrix", args)
  }

}

// detach()
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

// paintEngine()
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

// size()
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

// devType()
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

// mirrored(_Bool, _Bool)
func (this *QImage) mirrored(args ...interface{}) () {
  // mirrored(_Bool, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR6QImage8mirroredEbb
    // invoke: QImage mirrored(_Bool, _Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C._ZNKR6QImage8mirroredEbb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "mirrored", args)
  }

}

// width()
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

// valid(int, int)
func (this *QImage) valid(args ...interface{}) () {
  // valid(int, int)
  // valid(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5validEii
    // invoke: bool valid(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK6QImage5validEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK6QImage5validERK6QPoint
    // invoke: bool valid(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QImage5validERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "valid", args)
  }

}

// setDevicePixelRatio(qreal)
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

// textKeys()
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

// setOffset(const class QPoint &)
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

// save(const class QString &, const char *, int)
func (this *QImage) save(args ...interface{}) () {
  // save(const class QString &, const char *, int)
  // save(class QIODevice *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4saveERK7QStringPKci
    // invoke: bool save(const class QString &, const char *, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK6QImage4saveERK7QStringPKci(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK6QImage4saveEP9QIODevicePKci
    // invoke: bool save(class QIODevice *, const char *, int)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK6QImage4saveEP9QIODevicePKci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QImage", "save", args)
  }

}

// bits()
func (this *QImage) bits(args ...interface{}) () {
  // bits()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

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

// pixel(const class QPoint &)
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
    C._ZNK6QImage5pixelERK6QPoint(this.qclsinst, arg0)
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

// swap(class QImage &)
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
    C._ZN6QImage4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "swap", args)
  }

}

// isDetached()
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

// alphaChannel()
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

// setPixel(int, int, uint)
func (this *QImage) setPixel(args ...interface{}) () {
  // setPixel(int, int, uint)
  // setPixel(const class QPoint &, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8setPixelEiij
    // invoke: void setPixel(int, int, uint)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN6QImage8setPixelEiij(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN6QImage8setPixelERK6QPointj
    // invoke: void setPixel(const class QPoint &, uint)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN6QImage8setPixelERK6QPointj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "setPixel", args)
  }

}

// setDotsPerMeterY(int)
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

// format()
func (this *QImage) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage6formatEv
    // invoke: QImage::Format format()
    C._ZNK6QImage6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "format", args)
  }

}

// hasAlphaChannel()
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

// bitPlaneCount()
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

// pixelFormat()
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

// rgbSwapped()
func (this *QImage) rgbSwapped(args ...interface{}) () {
  // rgbSwapped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR6QImage10rgbSwappedEv
    // invoke: QImage rgbSwapped()
    C._ZNKR6QImage10rgbSwappedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "rgbSwapped", args)
  }

}

// isGrayscale()
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

// createHeuristicMask(_Bool)
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZNK6QImage19createHeuristicMaskEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "createHeuristicMask", args)
  }

}

// bytesPerLine()
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

// copy(const class QRect &)
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
    C._ZNK6QImage4copyEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QImage", "copy", args)
  }

}

// dotsPerMeterY()
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

// dotsPerMeterX()
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

// rect()
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

// setDotsPerMeterX(int)
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

// setAlphaChannel(const class QImage &)
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

// setColor(int, QRgb)
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

// devicePixelRatio()
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

// loadFromData(const uchar *, int, const char *)
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C._ZN6QImage12loadFromDataEPKhiPKc(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN6QImage12loadFromDataERK10QByteArrayPKc
    // invoke: bool loadFromData(const class QByteArray &, const char *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN6QImage12loadFromDataERK10QByteArrayPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "loadFromData", args)
  }

}

// setText(const class QString &, const class QString &)
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

// allGray()
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

// colorCount()
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

// isNull()
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

// depth()
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

// ~QImage()
func (this *QImage) FreeQImage(args ...interface{}) () {
  // ~QImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImageD0Ev
    // invoke: void ~QImage()
    C._ZN6QImageD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "~QImage", args)
  }

}

// pixelIndex(int, int)
func (this *QImage) pixelIndex(args ...interface{}) () {
  // pixelIndex(int, int)
  // pixelIndex(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10pixelIndexEii
    // invoke: int pixelIndex(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK6QImage10pixelIndexEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK6QImage10pixelIndexERK6QPoint
    // invoke: int pixelIndex(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK6QImage10pixelIndexERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "pixelIndex", args)
  }

}

// fill(uint)
func (this *QImage) fill(args ...interface{}) () {
  // fill(uint)
  // fill(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"

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

// colorTable()
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

// QImage()
func NewQImage(args ...interface{}) QImage {
  // QImage()
  // QImage(const class QString &, const char *)
  // QImage(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImageC1Ev
    // invoke: void QImage()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QImageC2Ev(qthis)
  case 1:
    // invoke: _ZN6QImageC1ERK7QStringPKc
    // invoke: void QImage(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QImageC2ERK7QStringPKc(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN6QImageC1ERKS_
    // invoke: void QImage(const class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN6QImageC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QImage", "QImage", args)
  }

  return QImage{}
}

// <= body block end

