package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qpixmap.h
// dst-file: /src/gui/qpixmap.go
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
  // proto:  bool QPixmap::save(const QString & fileName, const char * format, int quality);
extern void _ZNK7QPixmap4saveERK7QStringPKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2);
  // proto:  void QPixmap::swap(QPixmap & other);
extern void demth_ZN7QPixmap4swapERS_(void* qthis, void* arg0);
  // proto:  bool QPixmap::isQBitmap();
extern void _ZNK7QPixmap9isQBitmapEv(void* qthis);
  // proto:  qreal QPixmap::devicePixelRatio();
extern void _ZNK7QPixmap16devicePixelRatioEv(void* qthis);
  // proto:  void QPixmap::QPixmap(const QSize & );
extern void* dector_ZN7QPixmapC1ERK5QSize(void* arg0);
extern void _ZN7QPixmapC1ERK5QSize(void* qthis, void* arg0);
  // proto:  void QPixmap::fill(const QPaintDevice * device, int xofs, int yofs);
extern void demth_ZN7QPixmap4fillEPK12QPaintDeviceii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPixmap::QPixmap(const QSize & s, int type);
extern void* dector_ZN7QPixmapC1ERK5QSizei(void* arg0, int32_t arg1);
extern void _ZN7QPixmapC1ERK5QSizei(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QPixmap::fill(const QColor & fillColor);
extern void _ZN7QPixmap4fillERK6QColor(void* qthis, void* arg0);
  // proto:  int QPixmap::devType();
extern void _ZNK7QPixmap7devTypeEv(void* qthis);
  // proto:  void QPixmap::scroll(int dx, int dy, int x, int y, int width, int height, QRegion * exposed);
extern void demth_ZN7QPixmap6scrollEiiiiiiP7QRegion(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5, void* arg6);
  // proto:  QPixmap QPixmap::copy(const QRect & rect);
extern void _ZNK7QPixmap4copyERK5QRect(void* qthis, void* arg0);
  // proto: static QTransform QPixmap::trueMatrix(const QTransform & m, int w, int h);
extern void _ZN7QPixmap10trueMatrixERK10QTransformii(void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QPixmap::QPixmap(int w, int h);
extern void* dector_ZN7QPixmapC1Eii(int32_t arg0, int32_t arg1);
extern void _ZN7QPixmapC1Eii(void* qthis, int32_t arg0, int32_t arg1);
  // proto: static QPixmap QPixmap::grabWindow(WId , int x, int y, int w, int h);
extern void _ZN7QPixmap10grabWindowEiiiii(int32_t* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4);
  // proto:  void QPixmap::fill(const QPaintDevice * device, const QPoint & ofs);
extern void _ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint(void* qthis, void* arg0, void* arg1);
  // proto:  bool QPixmap::isDetached();
extern void _ZNK7QPixmap10isDetachedEv(void* qthis);
  // proto:  bool QPixmap::isNull();
extern void _ZNK7QPixmap6isNullEv(void* qthis);
  // proto:  QPixmap QPixmap::copy(int x, int y, int width, int height);
extern void demth_ZNK7QPixmap4copyEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3);
  // proto: static int QPixmap::defaultDepth();
extern void _ZN7QPixmap12defaultDepthEv();
  // proto:  void QPixmap::detach();
extern void _ZN7QPixmap6detachEv(void* qthis);
  // proto:  void QPixmap::scroll(int dx, int dy, const QRect & rect, QRegion * exposed);
extern void _ZN7QPixmap6scrollEiiRK5QRectP7QRegion(void* qthis, int32_t arg0, int32_t arg1, void* arg2, void* arg3);
  // proto:  void QPixmap::setMask(const QBitmap & );
extern void _ZN7QPixmap7setMaskERK7QBitmap(void* qthis, void* arg0);
  // proto:  void QPixmap::QPixmap();
extern void* dector_ZN7QPixmapC1Ev();
extern void _ZN7QPixmapC1Ev(void* qthis);
  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, const QRect & rect);
extern void _ZN7QPixmap10grabWidgetEP7QObjectRK5QRect(void* arg0, void* arg1);
  // proto:  void QPixmap::QPixmap(const QPixmap & );
extern void* dector_ZN7QPixmapC1ERKS_(void* arg0);
extern void _ZN7QPixmapC1ERKS_(void* qthis, void* arg0);
  // proto:  void QPixmap::setDevicePixelRatio(qreal scaleFactor);
extern void _ZN7QPixmap19setDevicePixelRatioEd(void* qthis, double arg0);
  // proto:  void QPixmap::QPixmap(const char *const [] xpm);
extern void* dector_ZN7QPixmapC1EPKPKc(char** arg0);
extern void _ZN7QPixmapC1EPKPKc(void* qthis, char** arg0);
  // proto:  qint64 QPixmap::cacheKey();
extern void _ZNK7QPixmap8cacheKeyEv(void* qthis);
  // proto:  QBitmap QPixmap::createHeuristicMask(bool clipTight);
extern void _ZNK7QPixmap19createHeuristicMaskEb(void* qthis, bool arg0);
  // proto:  int QPixmap::depth();
extern void _ZNK7QPixmap5depthEv(void* qthis);
  // proto:  QImage QPixmap::toImage();
extern void _ZNK7QPixmap7toImageEv(void* qthis);
  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, int x, int y, int w, int h);
extern void demth_ZN7QPixmap10grabWidgetEP7QObjectiiii(void* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4);
  // proto:  QPlatformPixmap * QPixmap::handle();
extern void _ZNK7QPixmap6handleEv(void* qthis);
  // proto:  bool QPixmap::hasAlphaChannel();
extern void _ZNK7QPixmap15hasAlphaChannelEv(void* qthis);
  // proto:  QRect QPixmap::rect();
extern void _ZNK7QPixmap4rectEv(void* qthis);
  // proto: static QMatrix QPixmap::trueMatrix(const QMatrix & m, int w, int h);
extern void _ZN7QPixmap10trueMatrixERK7QMatrixii(void* arg0, int32_t arg1, int32_t arg2);
  // proto:  QBitmap QPixmap::mask();
extern void _ZNK7QPixmap4maskEv(void* qthis);
  // proto:  int QPixmap::width();
extern void _ZNK7QPixmap5widthEv(void* qthis);
  // proto:  QPaintEngine * QPixmap::paintEngine();
extern void _ZNK7QPixmap11paintEngineEv(void* qthis);
  // proto:  void QPixmap::~QPixmap();
extern void _ZN7QPixmapD0Ev(void* qthis);
  // proto:  int QPixmap::height();
extern void _ZNK7QPixmap6heightEv(void* qthis);
  // proto:  bool QPixmap::save(QIODevice * device, const char * format, int quality);
extern void _ZNK7QPixmap4saveEP9QIODevicePKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2);
  // proto:  QSize QPixmap::size();
extern void _ZNK7QPixmap4sizeEv(void* qthis);
  // proto:  bool QPixmap::hasAlpha();
extern void _ZNK7QPixmap8hasAlphaEv(void* qthis);
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

// class sizeof(QPixmap)=1
type QPixmap struct {
  /*qbase*/ QPaintDevice;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  bool QPixmap::save(const QString & fileName, const char * format, int quality);
func (this *QPixmap) save(args ...interface{}) () {
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
    // invoke: _ZNK7QPixmap4saveERK7QStringPKci
    // invoke: bool save(const class QString &, const char *, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK7QPixmap4saveERK7QStringPKci(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK7QPixmap4saveEP9QIODevicePKci
    // invoke: bool save(class QIODevice *, const char *, int)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK7QPixmap4saveEP9QIODevicePKci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPixmap", "save", args)
  }

}

  // proto:  void QPixmap::swap(QPixmap & other);
func (this *QPixmap) swap(args ...interface{}) () {
  // swap(class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4swapERS_
    // invoke: void swap(class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QPixmap4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "swap", args)
  }

}

  // proto:  bool QPixmap::isQBitmap();
func (this *QPixmap) isQBitmap(args ...interface{}) () {
  // isQBitmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap9isQBitmapEv
    // invoke: bool isQBitmap()
    C._ZNK7QPixmap9isQBitmapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "isQBitmap", args)
  }

}

  // proto:  qreal QPixmap::devicePixelRatio();
func (this *QPixmap) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    C._ZNK7QPixmap16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "devicePixelRatio", args)
  }

}

  // proto:  void QPixmap::QPixmap(const QSize & );
func NewQPixmap(args ...interface{}) QPixmap {
  return QPixmap{}
}

  // proto:  void QPixmap::fill(const QPaintDevice * device, int xofs, int yofs);
func (this *QPixmap) fill(args ...interface{}) () {
  // fill(const class QPaintDevice *, int, int)
  // fill(const class QColor &)
  // fill(const class QPaintDevice *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[2][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceii
    // invoke: void fill(const class QPaintDevice *, int, int)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN7QPixmap4fillEPK12QPaintDeviceii(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN7QPixmap4fillERK6QColor
    // invoke: void fill(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QPixmap4fillERK6QColor(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint
    // invoke: void fill(const class QPaintDevice *, const class QPoint &)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPixmap", "fill", args)
  }

}

  // proto:  int QPixmap::devType();
func (this *QPixmap) devType(args ...interface{}) () {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap7devTypeEv
    // invoke: int devType()
    C._ZNK7QPixmap7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "devType", args)
  }

}

  // proto:  void QPixmap::scroll(int dx, int dy, int x, int y, int width, int height, QRegion * exposed);
func (this *QPixmap) scroll(args ...interface{}) () {
  // scroll(int, int, int, int, int, int, class QRegion *)
  // scroll(int, int, const class QRect &, class QRegion *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = qtrt.Int32Ty(false) // "int"
  vtys[0][6] = reflect.TypeOf(QRegion{}) // "QRegion *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][3] = reflect.TypeOf(QRegion{}) // "QRegion *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap6scrollEiiiiiiP7QRegion
    // invoke: void scroll(int, int, int, int, int, int, class QRegion *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(args[5].(int32))
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QRegion).qclsinst
    if false {fmt.Println(arg6)}
    C.demth_ZN7QPixmap6scrollEiiiiiiP7QRegion(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 1:
    // invoke: _ZN7QPixmap6scrollEiiRK5QRectP7QRegion
    // invoke: void scroll(int, int, const class QRect &, class QRegion *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRect).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QRegion).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN7QPixmap6scrollEiiRK5QRectP7QRegion(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPixmap", "scroll", args)
  }

}

  // proto:  QPixmap QPixmap::copy(const QRect & rect);
func (this *QPixmap) copy(args ...interface{}) () {
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
    // invoke: _ZNK7QPixmap4copyERK5QRect
    // invoke: QPixmap copy(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QPixmap4copyERK5QRect(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QPixmap4copyEiiii
    // invoke: QPixmap copy(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.demth_ZNK7QPixmap4copyEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPixmap", "copy", args)
  }

}

  // proto: static QTransform QPixmap::trueMatrix(const QTransform & m, int w, int h);
func (this *QPixmap) trueMatrix_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "trueMatrix", args)
  }

}

  // proto: static QPixmap QPixmap::grabWindow(WId , int x, int y, int w, int h);
func (this *QPixmap) grabWindow_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "grabWindow", args)
  }

}

  // proto:  bool QPixmap::isDetached();
func (this *QPixmap) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap10isDetachedEv
    // invoke: bool isDetached()
    C._ZNK7QPixmap10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "isDetached", args)
  }

}

  // proto:  bool QPixmap::isNull();
func (this *QPixmap) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6isNullEv
    // invoke: bool isNull()
    C._ZNK7QPixmap6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "isNull", args)
  }

}

  // proto: static int QPixmap::defaultDepth();
func (this *QPixmap) defaultDepth_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "defaultDepth", args)
  }

}

  // proto:  void QPixmap::detach();
func (this *QPixmap) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap6detachEv
    // invoke: void detach()
    C._ZN7QPixmap6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "detach", args)
  }

}

  // proto:  void QPixmap::setMask(const QBitmap & );
func (this *QPixmap) setMask(args ...interface{}) () {
  // setMask(const class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap7setMaskERK7QBitmap
    // invoke: void setMask(const class QBitmap &)
    var arg0 = args[0].(QBitmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QPixmap7setMaskERK7QBitmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "setMask", args)
  }

}

  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, const QRect & rect);
func (this *QPixmap) grabWidget_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "grabWidget", args)
  }

}

  // proto:  void QPixmap::setDevicePixelRatio(qreal scaleFactor);
func (this *QPixmap) setDevicePixelRatio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap19setDevicePixelRatioEd
    // invoke: void setDevicePixelRatio(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C._ZN7QPixmap19setDevicePixelRatioEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "setDevicePixelRatio", args)
  }

}

  // proto:  qint64 QPixmap::cacheKey();
func (this *QPixmap) cacheKey(args ...interface{}) () {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap8cacheKeyEv
    // invoke: qint64 cacheKey()
    C._ZNK7QPixmap8cacheKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "cacheKey", args)
  }

}

  // proto:  QBitmap QPixmap::createHeuristicMask(bool clipTight);
func (this *QPixmap) createHeuristicMask(args ...interface{}) () {
  // createHeuristicMask(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap19createHeuristicMaskEb
    // invoke: QBitmap createHeuristicMask(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZNK7QPixmap19createHeuristicMaskEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "createHeuristicMask", args)
  }

}

  // proto:  int QPixmap::depth();
func (this *QPixmap) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap5depthEv
    // invoke: int depth()
    C._ZNK7QPixmap5depthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "depth", args)
  }

}

  // proto:  QImage QPixmap::toImage();
func (this *QPixmap) toImage(args ...interface{}) () {
  // toImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap7toImageEv
    // invoke: QImage toImage()
    C._ZNK7QPixmap7toImageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "toImage", args)
  }

}

  // proto:  QPlatformPixmap * QPixmap::handle();
func (this *QPixmap) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6handleEv
    // invoke: QPlatformPixmap * handle()
    C._ZNK7QPixmap6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "handle", args)
  }

}

  // proto:  bool QPixmap::hasAlphaChannel();
func (this *QPixmap) hasAlphaChannel(args ...interface{}) () {
  // hasAlphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap15hasAlphaChannelEv
    // invoke: bool hasAlphaChannel()
    C._ZNK7QPixmap15hasAlphaChannelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlphaChannel", args)
  }

}

  // proto:  QRect QPixmap::rect();
func (this *QPixmap) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4rectEv
    // invoke: QRect rect()
    C._ZNK7QPixmap4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "rect", args)
  }

}

  // proto:  QBitmap QPixmap::mask();
func (this *QPixmap) mask(args ...interface{}) () {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4maskEv
    // invoke: QBitmap mask()
    C._ZNK7QPixmap4maskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "mask", args)
  }

}

  // proto:  int QPixmap::width();
func (this *QPixmap) width(args ...interface{}) () {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap5widthEv
    // invoke: int width()
    C._ZNK7QPixmap5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "width", args)
  }

}

  // proto:  QPaintEngine * QPixmap::paintEngine();
func (this *QPixmap) paintEngine(args ...interface{}) () {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    C._ZNK7QPixmap11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "paintEngine", args)
  }

}

  // proto:  void QPixmap::~QPixmap();
func (this *QPixmap) FreeQPixmap(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPixmap", "~QPixmap", args)
  }

}

  // proto:  int QPixmap::height();
func (this *QPixmap) height(args ...interface{}) () {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6heightEv
    // invoke: int height()
    C._ZNK7QPixmap6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "height", args)
  }

}

  // proto:  QSize QPixmap::size();
func (this *QPixmap) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4sizeEv
    // invoke: QSize size()
    C._ZNK7QPixmap4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "size", args)
  }

}

  // proto:  bool QPixmap::hasAlpha();
func (this *QPixmap) hasAlpha(args ...interface{}) () {
  // hasAlpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap8hasAlphaEv
    // invoke: bool hasAlpha()
    C._ZNK7QPixmap8hasAlphaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlpha", args)
  }

}

// <= body block end

