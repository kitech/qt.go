package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  bool QImage::load(QIODevice * device, const char * format);
extern bool C_ZN6QImage4loadEP9QIODevicePKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QImage::load(const QString & fileName, const char * format);
extern bool C_ZN6QImage4loadERK7QStringPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto: static QImage::Format QImage::toImageFormat(QPixelFormat format);
extern void C_ZN6QImage13toImageFormatE12QPixelFormat(void* arg0); // 4
  // proto:  QRgb QImage::color(int i);
extern int32_t C_ZNK6QImage5colorEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QImage::text(const QString & key);
extern void* C_ZNK6QImage4textERK7QString(void* qthis, void* arg0); // 4
  // proto:  QPoint QImage::offset();
extern void* C_ZNK6QImage6offsetEv(void* qthis); // 4
  // proto:  const uchar * QImage::constScanLine(int );
extern void* C_ZNK6QImage13constScanLineEi(void* qthis, int32_t arg0); // 4
  // proto:  int QImage::height();
extern int32_t C_ZNK6QImage6heightEv(void* qthis); // 4
  // proto: static QImage QImage::fromData(const uchar * data, int size, const char * format);
extern void* C_ZN6QImage8fromDataEPKhiPKc(void* arg0, int32_t arg1, void* arg2); // 4
  // proto: static QImage QImage::fromData(const QByteArray & data, const char * format);
extern void* C_ZN6QImage8fromDataERK10QByteArrayPKc(void* arg0, void* arg1); // 2
  // proto:  uchar * QImage::scanLine(int );
extern void* C_ZN6QImage8scanLineEi(void* qthis, int32_t arg0); // 4
  // proto:  int QImage::byteCount();
extern int32_t C_ZNK6QImage9byteCountEv(void* qthis); // 4
  // proto:  qint64 QImage::cacheKey();
extern int64_t C_ZNK6QImage8cacheKeyEv(void* qthis); // 4
  // proto:  void QImage::setColorCount(int );
extern void C_ZN6QImage13setColorCountEi(void* qthis, int32_t arg0); // 4
  // proto:  const uchar * QImage::constBits();
extern void* C_ZNK6QImage9constBitsEv(void* qthis); // 4
  // proto: static QMatrix QImage::trueMatrix(const QMatrix & , int w, int h);
extern void* C_ZN6QImage10trueMatrixERK7QMatrixii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto: static QTransform QImage::trueMatrix(const QTransform & , int w, int h);
extern void* C_ZN6QImage10trueMatrixERK10QTransformii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QImage::detach();
extern void C_ZN6QImage6detachEv(void* qthis); // 4
  // proto:  QPaintEngine * QImage::paintEngine();
extern void* C_ZNK6QImage11paintEngineEv(void* qthis); // 4
  // proto:  QSize QImage::size();
extern void* C_ZNK6QImage4sizeEv(void* qthis); // 4
  // proto:  int QImage::devType();
extern int32_t C_ZNK6QImage7devTypeEv(void* qthis); // 4
  // proto:  QImage QImage::mirrored(bool horizontally, bool vertically);
extern void* C_ZNKR6QImage8mirroredEbb(void* qthis, bool arg0, bool arg1); // 2
  // proto:  int QImage::width();
extern int32_t C_ZNK6QImage5widthEv(void* qthis); // 4
  // proto:  bool QImage::valid(int x, int y);
extern bool C_ZNK6QImage5validEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  bool QImage::valid(const QPoint & pt);
extern bool C_ZNK6QImage5validERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QImage::setDevicePixelRatio(qreal scaleFactor);
extern void C_ZN6QImage19setDevicePixelRatioEd(void* qthis, double arg0); // 4
  // proto:  QStringList QImage::textKeys();
extern void C_ZNK6QImage8textKeysEv(void* qthis); // 4
  // proto:  void QImage::setOffset(const QPoint & );
extern void C_ZN6QImage9setOffsetERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  bool QImage::save(const QString & fileName, const char * format, int quality);
extern bool C_ZNK6QImage4saveERK7QStringPKci(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QImage::save(QIODevice * device, const char * format, int quality);
extern bool C_ZNK6QImage4saveEP9QIODevicePKci(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  uchar * QImage::bits();
extern void* C_ZN6QImage4bitsEv(void* qthis); // 4
  // proto:  QRgb QImage::pixel(const QPoint & pt);
extern int32_t C_ZNK6QImage5pixelERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  QRgb QImage::pixel(int x, int y);
extern int32_t C_ZNK6QImage5pixelEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QImage::swap(QImage & other);
extern void C_ZN6QImage4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QImage::isDetached();
extern bool C_ZNK6QImage10isDetachedEv(void* qthis); // 4
  // proto:  QImage QImage::alphaChannel();
extern void* C_ZNK6QImage12alphaChannelEv(void* qthis); // 4
  // proto:  void QImage::setPixel(int x, int y, uint index_or_rgb);
extern void C_ZN6QImage8setPixelEiij(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QImage::setPixel(const QPoint & pt, uint index_or_rgb);
extern void C_ZN6QImage8setPixelERK6QPointj(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QImage::setDotsPerMeterY(int );
extern void C_ZN6QImage16setDotsPerMeterYEi(void* qthis, int32_t arg0); // 4
  // proto:  QImage::Format QImage::format();
extern void C_ZNK6QImage6formatEv(void* qthis); // 4
  // proto:  bool QImage::hasAlphaChannel();
extern bool C_ZNK6QImage15hasAlphaChannelEv(void* qthis); // 4
  // proto:  int QImage::bitPlaneCount();
extern int32_t C_ZNK6QImage13bitPlaneCountEv(void* qthis); // 4
  // proto:  QPixelFormat QImage::pixelFormat();
extern void* C_ZNK6QImage11pixelFormatEv(void* qthis); // 4
  // proto:  QImage QImage::rgbSwapped();
extern void* C_ZNKR6QImage10rgbSwappedEv(void* qthis); // 2
  // proto:  bool QImage::isGrayscale();
extern bool C_ZNK6QImage11isGrayscaleEv(void* qthis); // 4
  // proto:  QImage QImage::createHeuristicMask(bool clipTight);
extern void* C_ZNK6QImage19createHeuristicMaskEb(void* qthis, bool arg0); // 4
  // proto:  int QImage::bytesPerLine();
extern int32_t C_ZNK6QImage12bytesPerLineEv(void* qthis); // 4
  // proto:  QImage QImage::copy(const QRect & rect);
extern void* C_ZNK6QImage4copyERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QImage QImage::copy(int x, int y, int w, int h);
extern void* C_ZNK6QImage4copyEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  int QImage::dotsPerMeterY();
extern int32_t C_ZNK6QImage13dotsPerMeterYEv(void* qthis); // 4
  // proto:  int QImage::dotsPerMeterX();
extern int32_t C_ZNK6QImage13dotsPerMeterXEv(void* qthis); // 4
  // proto:  QRect QImage::rect();
extern void* C_ZNK6QImage4rectEv(void* qthis); // 4
  // proto:  void QImage::setDotsPerMeterX(int );
extern void C_ZN6QImage16setDotsPerMeterXEi(void* qthis, int32_t arg0); // 4
  // proto:  void QImage::setAlphaChannel(const QImage & alphaChannel);
extern void C_ZN6QImage15setAlphaChannelERKS_(void* qthis, void* arg0); // 4
  // proto:  void QImage::setColor(int i, QRgb c);
extern void C_ZN6QImage8setColorEij(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  qreal QImage::devicePixelRatio();
extern double C_ZNK6QImage16devicePixelRatioEv(void* qthis); // 4
  // proto:  bool QImage::loadFromData(const uchar * buf, int len, const char * format);
extern bool C_ZN6QImage12loadFromDataEPKhiPKc(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QImage::loadFromData(const QByteArray & data, const char * aformat);
extern bool C_ZN6QImage12loadFromDataERK10QByteArrayPKc(void* qthis, void* arg0, void* arg1); // 2
  // proto:  void QImage::setText(const QString & key, const QString & value);
extern void C_ZN6QImage7setTextERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QImage::allGray();
extern bool C_ZNK6QImage7allGrayEv(void* qthis); // 4
  // proto:  int QImage::colorCount();
extern int32_t C_ZNK6QImage10colorCountEv(void* qthis); // 4
  // proto:  bool QImage::isNull();
extern bool C_ZNK6QImage6isNullEv(void* qthis); // 4
  // proto:  int QImage::depth();
extern int32_t C_ZNK6QImage5depthEv(void* qthis); // 4
  // proto:  void QImage::~QImage();
extern void C_ZN6QImageD2Ev(void* qthis); // 4
  // proto:  int QImage::pixelIndex(int x, int y);
extern int32_t C_ZNK6QImage10pixelIndexEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QImage::pixelIndex(const QPoint & pt);
extern int32_t C_ZNK6QImage10pixelIndexERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  void QImage::fill(uint pixel);
extern void C_ZN6QImage4fillEj(void* qthis, int32_t arg0); // 4
  // proto:  void QImage::fill(const QColor & color);
extern void C_ZN6QImage4fillERK6QColor(void* qthis, void* arg0); // 4
  // proto:  QVector<QRgb> QImage::colorTable();
extern void C_ZNK6QImage10colorTableEv(void* qthis); // 4
  // proto:  void QImage::QImage();
extern void* C_ZN6QImageC2Ev(); // 3
  // proto:  void QImage::QImage(const QString & fileName, const char * format);
extern void* C_ZN6QImageC2ERK7QStringPKc(void* arg0, void* arg1); // 3
  // proto:  void QImage::QImage(const QImage & );
extern void* C_ZN6QImageC2ERKS_(void* arg0); // 3
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

// class sizeof(QImage)=32
type QImage struct {
  /*qbase*/ QPaintDevice;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// load(class QIODevice *, const char *)
func (this *QImage) Load(args ...interface{}) (ret interface{}) {
  // load(class QIODevice *, const char *)
  // load(const class QString &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QIODevice{}) // "QIODevice *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage4loadEP9QIODevicePKc
    // invoke: bool load(class QIODevice *, const char *)
    var arg0 = args[0].(*qtcore.QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN6QImage4loadEP9QIODevicePKc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN6QImage4loadERK7QStringPKc
    // invoke: bool load(const class QString &, const char *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN6QImage4loadERK7QStringPKc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "load", args)
  }

  return
}

// toImageFormat(class QPixelFormat)
func (this *QImage) ToImageFormat_s(args ...interface{}) () {
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
    var arg0 = args[0].(*QPixelFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage13toImageFormatE12QPixelFormat(arg0)
  default:
    qtrt.ErrorResolve("QImage", "toImageFormat", args)
  }

  return
}

// color(int)
func (this *QImage) Color(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage5colorEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "QRgb"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "color", args)
  }

  return
}

// text(const class QString &)
func (this *QImage) Text(args ...interface{}) (ret interface{}) {
  // text(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4textERK7QString
    // invoke: QString text(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage4textERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "text", args)
  }

  return
}

// offset()
func (this *QImage) Offset(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage6offsetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QPoint{}) // "QPoint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "offset", args)
  }

  return
}

// constScanLine(int)
func (this *QImage) ConstScanLine(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage13constScanLineEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const uchar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "constScanLine", args)
  }

  return
}

// height()
func (this *QImage) Height(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "height", args)
  }

  return
}

// fromData(const uchar *, int, const char *)
func (this *QImage) FromData_s(args ...interface{}) (ret interface{}) {
  // fromData(const uchar *, int, const char *)
  // fromData(const class QByteArray &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8fromDataEPKhiPKc
    // invoke: QImage fromData(const uchar *, int, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var ret0 = C.C_ZN6QImage8fromDataEPKhiPKc(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN6QImage8fromDataERK10QByteArrayPKc
    // invoke: QImage fromData(const class QByteArray &, const char *)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN6QImage8fromDataERK10QByteArrayPKc(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "fromData", args)
  }

  return
}

// scanLine(int)
func (this *QImage) ScanLine(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QImage8scanLineEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "uchar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "scanLine", args)
  }

  return
}

// byteCount()
func (this *QImage) ByteCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage9byteCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "byteCount", args)
  }

  return
}

// cacheKey()
func (this *QImage) CacheKey(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage8cacheKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "cacheKey", args)
  }

  return
}

// setColorCount(int)
func (this *QImage) SetColorCount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage13setColorCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setColorCount", args)
  }

  return
}

// constBits()
func (this *QImage) ConstBits(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage9constBitsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const uchar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "constBits", args)
  }

  return
}

// trueMatrix(const class QMatrix &, int, int)
func (this *QImage) TrueMatrix_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN6QImage10trueMatrixERK7QMatrixii(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN6QImage10trueMatrixERK10QTransformii
    // invoke: QTransform trueMatrix(const class QTransform &, int, int)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN6QImage10trueMatrixERK10QTransformii(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "trueMatrix", args)
  }

  return
}

// detach()
func (this *QImage) Detach(args ...interface{}) () {
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
    C.C_ZN6QImage6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "detach", args)
  }

  return
}

// paintEngine()
func (this *QImage) PaintEngine(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage11paintEngineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "paintEngine", args)
  }

  return
}

// size()
func (this *QImage) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "size", args)
  }

  return
}

// devType()
func (this *QImage) DevType(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage7devTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "devType", args)
  }

  return
}

// mirrored(_Bool, _Bool)
func (this *QImage) Mirrored(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNKR6QImage8mirroredEbb(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "mirrored", args)
  }

  return
}

// width()
func (this *QImage) Width(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "width", args)
  }

  return
}

// valid(int, int)
func (this *QImage) Valid(args ...interface{}) (ret interface{}) {
  // valid(int, int)
  // valid(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5validEii
    // invoke: bool valid(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK6QImage5validEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK6QImage5validERK6QPoint
    // invoke: bool valid(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage5validERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "valid", args)
  }

  return
}

// setDevicePixelRatio(qreal)
func (this *QImage) SetDevicePixelRatio(args ...interface{}) () {
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
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage19setDevicePixelRatioEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setDevicePixelRatio", args)
  }

  return
}

// textKeys()
func (this *QImage) TextKeys(args ...interface{}) () {
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
    C.C_ZNK6QImage8textKeysEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "textKeys", args)
  }

  return
}

// setOffset(const class QPoint &)
func (this *QImage) SetOffset(args ...interface{}) () {
  // setOffset(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage9setOffsetERK6QPoint
    // invoke: void setOffset(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage9setOffsetERK6QPoint(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setOffset", args)
  }

  return
}

// save(const class QString &, const char *, int)
func (this *QImage) Save(args ...interface{}) (ret interface{}) {
  // save(const class QString &, const char *, int)
  // save(class QIODevice *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage4saveERK7QStringPKci
    // invoke: bool save(const class QString &, const char *, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK6QImage4saveERK7QStringPKci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK6QImage4saveEP9QIODevicePKci
    // invoke: bool save(class QIODevice *, const char *, int)
    var arg0 = args[0].(*qtcore.QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK6QImage4saveEP9QIODevicePKci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "save", args)
  }

  return
}

// bits()
func (this *QImage) Bits(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN6QImage4bitsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "uchar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "bits", args)
  }

  return
}

// pixel(const class QPoint &)
func (this *QImage) Pixel(args ...interface{}) (ret interface{}) {
  // pixel(const class QPoint &)
  // pixel(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage5pixelERK6QPoint
    // invoke: QRgb pixel(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage5pixelERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "QRgb"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK6QImage5pixelEii
    // invoke: QRgb pixel(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK6QImage5pixelEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "QRgb"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "pixel", args)
  }

  return
}

// swap(class QImage &)
func (this *QImage) Swap(args ...interface{}) () {
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
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "swap", args)
  }

  return
}

// isDetached()
func (this *QImage) IsDetached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "isDetached", args)
  }

  return
}

// alphaChannel()
func (this *QImage) AlphaChannel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage12alphaChannelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "alphaChannel", args)
  }

  return
}

// setPixel(int, int, uint)
func (this *QImage) SetPixel(args ...interface{}) () {
  // setPixel(int, int, uint)
  // setPixel(const class QPoint &, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "uint"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage8setPixelEiij
    // invoke: void setPixel(int, int, uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN6QImage8setPixelEiij(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN6QImage8setPixelERK6QPointj
    // invoke: void setPixel(const class QPoint &, uint)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN6QImage8setPixelERK6QPointj(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "setPixel", args)
  }

  return
}

// setDotsPerMeterY(int)
func (this *QImage) SetDotsPerMeterY(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage16setDotsPerMeterYEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setDotsPerMeterY", args)
  }

  return
}

// format()
func (this *QImage) Format(args ...interface{}) () {
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
    C.C_ZNK6QImage6formatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "format", args)
  }

  return
}

// hasAlphaChannel()
func (this *QImage) HasAlphaChannel(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage15hasAlphaChannelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "hasAlphaChannel", args)
  }

  return
}

// bitPlaneCount()
func (this *QImage) BitPlaneCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage13bitPlaneCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "bitPlaneCount", args)
  }

  return
}

// pixelFormat()
func (this *QImage) PixelFormat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage11pixelFormatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixelFormat{}) // "QPixelFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "pixelFormat", args)
  }

  return
}

// rgbSwapped()
func (this *QImage) RgbSwapped(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNKR6QImage10rgbSwappedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "rgbSwapped", args)
  }

  return
}

// isGrayscale()
func (this *QImage) IsGrayscale(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage11isGrayscaleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "isGrayscale", args)
  }

  return
}

// createHeuristicMask(_Bool)
func (this *QImage) CreateHeuristicMask(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage19createHeuristicMaskEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "createHeuristicMask", args)
  }

  return
}

// bytesPerLine()
func (this *QImage) BytesPerLine(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage12bytesPerLineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "bytesPerLine", args)
  }

  return
}

// copy(const class QRect &)
func (this *QImage) Copy(args ...interface{}) (ret interface{}) {
  // copy(const class QRect &)
  // copy(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
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
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage4copyERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK6QImage4copyEiiii
    // invoke: QImage copy(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK6QImage4copyEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "copy", args)
  }

  return
}

// dotsPerMeterY()
func (this *QImage) DotsPerMeterY(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage13dotsPerMeterYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "dotsPerMeterY", args)
  }

  return
}

// dotsPerMeterX()
func (this *QImage) DotsPerMeterX(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage13dotsPerMeterXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "dotsPerMeterX", args)
  }

  return
}

// rect()
func (this *QImage) Rect(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "rect", args)
  }

  return
}

// setDotsPerMeterX(int)
func (this *QImage) SetDotsPerMeterX(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage16setDotsPerMeterXEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setDotsPerMeterX", args)
  }

  return
}

// setAlphaChannel(const class QImage &)
func (this *QImage) SetAlphaChannel(args ...interface{}) () {
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
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage15setAlphaChannelERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "setAlphaChannel", args)
  }

  return
}

// setColor(int, QRgb)
func (this *QImage) SetColor(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN6QImage8setColorEij(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "setColor", args)
  }

  return
}

// devicePixelRatio()
func (this *QImage) DevicePixelRatio(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage16devicePixelRatioEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "devicePixelRatio", args)
  }

  return
}

// loadFromData(const uchar *, int, const char *)
func (this *QImage) LoadFromData(args ...interface{}) (ret interface{}) {
  // loadFromData(const uchar *, int, const char *)
  // loadFromData(const class QByteArray &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const uchar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage12loadFromDataEPKhiPKc
    // invoke: bool loadFromData(const uchar *, int, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[0][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var ret0 = C.C_ZN6QImage12loadFromDataEPKhiPKc(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN6QImage12loadFromDataERK10QByteArrayPKc
    // invoke: bool loadFromData(const class QByteArray &, const char *)
    var arg0 = args[0].(*qtcore.QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN6QImage12loadFromDataERK10QByteArrayPKc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "loadFromData", args)
  }

  return
}

// setText(const class QString &, const class QString &)
func (this *QImage) SetText(args ...interface{}) () {
  // setText(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QImage7setTextERK7QStringS2_
    // invoke: void setText(const class QString &, const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN6QImage7setTextERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImage", "setText", args)
  }

  return
}

// allGray()
func (this *QImage) AllGray(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage7allGrayEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "allGray", args)
  }

  return
}

// colorCount()
func (this *QImage) ColorCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage10colorCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "colorCount", args)
  }

  return
}

// isNull()
func (this *QImage) IsNull(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "isNull", args)
  }

  return
}

// depth()
func (this *QImage) Depth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK6QImage5depthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "depth", args)
  }

  return
}

// ~QImage()
func (this *QImage) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN6QImageD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QImage", "~QImage", args)
  }

  return
}

// pixelIndex(int, int)
func (this *QImage) PixelIndex(args ...interface{}) (ret interface{}) {
  // pixelIndex(int, int)
  // pixelIndex(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QImage10pixelIndexEii
    // invoke: int pixelIndex(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK6QImage10pixelIndexEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK6QImage10pixelIndexERK6QPoint
    // invoke: int pixelIndex(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QImage10pixelIndexERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImage", "pixelIndex", args)
  }

  return
}

// fill(uint)
func (this *QImage) Fill(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage4fillEj(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QImage4fillERK6QColor
    // invoke: void fill(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QImage4fillERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImage", "fill", args)
  }

  return
}

// colorTable()
func (this *QImage) ColorTable(args ...interface{}) () {
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
    C.C_ZNK6QImage10colorTableEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImage", "colorTable", args)
  }

  return
}

// QImage()
func GcfreeQImage(this *QImage) {
  qtrt.UniverseFree(this)
}
func NewQImage(args ...interface{}) *QImage {
  // QImage()
  // QImage(const class QString &, const char *)
  // QImage(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
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
    qthis = C.C_ZN6QImageC2Ev()
    this := &QImage{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQImage)
    return this
  case 1:
    // invoke: _ZN6QImageC1ERK7QStringPKc
    // invoke: void QImage(const class QString &, const char *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QImageC2ERK7QStringPKc(arg0, arg1)
    this := &QImage{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQImage)
    return this
  case 2:
    // invoke: _ZN6QImageC1ERKS_
    // invoke: void QImage(const class QImage &)
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QImageC2ERKS_(arg0)
    this := &QImage{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQImage)
    return this
  default:
    qtrt.ErrorResolve("QImage", "QImage", args)
  }

  return nil // QImage{Qclsinst:qthis}
}

// <= body block end

