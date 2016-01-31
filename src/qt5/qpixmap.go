package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QPixmap::height();
extern void C_ZNK7QPixmap6heightEv(void* qthis); // 4
  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, int x, int y, int w, int h);
extern void C_ZN7QPixmap10grabWidgetEP7QObjectiiii(void* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 2
  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, const QRect & rect);
extern void C_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect(void* arg0, void* arg1); // 4
  // proto:  QPaintEngine * QPixmap::paintEngine();
extern void C_ZNK7QPixmap11paintEngineEv(void* qthis); // 4
  // proto:  QPlatformPixmap * QPixmap::handle();
extern void C_ZNK7QPixmap6handleEv(void* qthis); // 4
  // proto:  void QPixmap::~QPixmap();
extern void C_ZN7QPixmapD2Ev(void* qthis); // 4
  // proto:  void QPixmap::fill(const QColor & fillColor);
extern void C_ZN7QPixmap4fillERK6QColor(void* qthis, void* arg0); // 4
  // proto:  void QPixmap::fill(const QPaintDevice * device, int xofs, int yofs);
extern void C_ZN7QPixmap4fillEPK12QPaintDeviceii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 2
  // proto:  void QPixmap::fill(const QPaintDevice * device, const QPoint & ofs);
extern void C_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRect QPixmap::rect();
extern void C_ZNK7QPixmap4rectEv(void* qthis); // 4
  // proto: static QMatrix QPixmap::trueMatrix(const QMatrix & m, int w, int h);
extern void C_ZN7QPixmap10trueMatrixERK7QMatrixii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto: static QTransform QPixmap::trueMatrix(const QTransform & m, int w, int h);
extern void C_ZN7QPixmap10trueMatrixERK10QTransformii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPixmap::detach();
extern void C_ZN7QPixmap6detachEv(void* qthis); // 4
  // proto:  QSize QPixmap::size();
extern void C_ZNK7QPixmap4sizeEv(void* qthis); // 4
  // proto:  int QPixmap::depth();
extern void C_ZNK7QPixmap5depthEv(void* qthis); // 4
  // proto:  int QPixmap::width();
extern void C_ZNK7QPixmap5widthEv(void* qthis); // 4
  // proto:  void QPixmap::swap(QPixmap & other);
extern void C_ZN7QPixmap4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QPixmap::isDetached();
extern void C_ZNK7QPixmap10isDetachedEv(void* qthis); // 4
  // proto:  qint64 QPixmap::cacheKey();
extern void C_ZNK7QPixmap8cacheKeyEv(void* qthis); // 4
  // proto:  void QPixmap::setDevicePixelRatio(qreal scaleFactor);
extern void C_ZN7QPixmap19setDevicePixelRatioEd(void* qthis, double arg0); // 4
  // proto:  bool QPixmap::hasAlphaChannel();
extern void C_ZNK7QPixmap15hasAlphaChannelEv(void* qthis); // 4
  // proto:  bool QPixmap::hasAlpha();
extern void C_ZNK7QPixmap8hasAlphaEv(void* qthis); // 4
  // proto:  void QPixmap::setMask(const QBitmap & );
extern void C_ZN7QPixmap7setMaskERK7QBitmap(void* qthis, void* arg0); // 4
  // proto:  QBitmap QPixmap::createHeuristicMask(bool clipTight);
extern void C_ZNK7QPixmap19createHeuristicMaskEb(void* qthis, bool arg0); // 4
  // proto: static int QPixmap::defaultDepth();
extern void C_ZN7QPixmap12defaultDepthEv(); // 4
  // proto:  QPixmap QPixmap::copy(const QRect & rect);
extern void C_ZNK7QPixmap4copyERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPixmap QPixmap::copy(int x, int y, int width, int height);
extern void C_ZNK7QPixmap4copyEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  bool QPixmap::save(const QString & fileName, const char * format, int quality);
extern void C_ZNK7QPixmap4saveERK7QStringPKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2); // 4
  // proto:  bool QPixmap::save(QIODevice * device, const char * format, int quality);
extern void C_ZNK7QPixmap4saveEP9QIODevicePKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2); // 4
  // proto:  qreal QPixmap::devicePixelRatio();
extern void C_ZNK7QPixmap16devicePixelRatioEv(void* qthis); // 4
  // proto:  bool QPixmap::isQBitmap();
extern void C_ZNK7QPixmap9isQBitmapEv(void* qthis); // 4
  // proto:  QBitmap QPixmap::mask();
extern void C_ZNK7QPixmap4maskEv(void* qthis); // 4
  // proto:  QImage QPixmap::toImage();
extern void C_ZNK7QPixmap7toImageEv(void* qthis); // 4
  // proto:  void QPixmap::QPixmap();
extern void C_ZN7QPixmapC2Ev(void* qthis); // 3
  // proto:  void QPixmap::QPixmap(int w, int h);
extern void C_ZN7QPixmapC2Eii(void* qthis, int32_t arg0, int32_t arg1); // 3
  // proto:  void QPixmap::QPixmap(const QPixmap & );
extern void C_ZN7QPixmapC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QPixmap::QPixmap(const QSize & );
extern void C_ZN7QPixmapC2ERK5QSize(void* qthis, void* arg0); // 3
  // proto:  bool QPixmap::isNull();
extern void C_ZNK7QPixmap6isNullEv(void* qthis); // 4
  // proto:  int QPixmap::devType();
extern void C_ZNK7QPixmap7devTypeEv(void* qthis); // 4
  // proto:  void QPixmap::scroll(int dx, int dy, int x, int y, int width, int height, QRegion * exposed);
extern void C_ZN7QPixmap6scrollEiiiiiiP7QRegion(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, int32_t arg5, void* arg6); // 2
  // proto:  void QPixmap::scroll(int dx, int dy, const QRect & rect, QRegion * exposed);
extern void C_ZN7QPixmap6scrollEiiRK5QRectP7QRegion(void* qthis, int32_t arg0, int32_t arg1, void* arg2, void* arg3); // 4
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

// height()
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
    C.C_ZNK7QPixmap6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "height", args)
  }

}

// grabWidget(class QObject *, int, int, int, int)
func (this *QPixmap) grabWidget_s(args ...interface{}) () {
  // grabWidget(class QObject *, int, int, int, int)
  // grabWidget(class QObject *, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap10grabWidgetEP7QObjectiiii
    // invoke: QPixmap grabWidget(class QObject *, int, int, int, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN7QPixmap10grabWidgetEP7QObjectiiii(arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZN7QPixmap10grabWidgetEP7QObjectRK5QRect
    // invoke: QPixmap grabWidget(class QObject *, const class QRect &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRect).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect(arg0, arg1)
  default:
    qtrt.ErrorResolve("QPixmap", "grabWidget", args)
  }

}

// paintEngine()
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
    C.C_ZNK7QPixmap11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "paintEngine", args)
  }

}

// handle()
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
    C.C_ZNK7QPixmap6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "handle", args)
  }

}

// ~QPixmap()
func (this *QPixmap) FreeQPixmap(args ...interface{}) () {
  // ~QPixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmapD0Ev
    // invoke: void ~QPixmap()
    C.C_ZN7QPixmapD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "~QPixmap", args)
  }

}

// fill(const class QColor &)
func (this *QPixmap) fill(args ...interface{}) () {
  // fill(const class QColor &)
  // fill(const class QPaintDevice *, int, int)
  // fill(const class QPaintDevice *, const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPaintDevice{}) // "const QPaintDevice *"
  vtys[2][1] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4fillERK6QColor
    // invoke: void fill(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QPixmap4fillERK6QColor(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceii
    // invoke: void fill(const class QPaintDevice *, int, int)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN7QPixmap4fillEPK12QPaintDeviceii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint
    // invoke: void fill(const class QPaintDevice *, const class QPoint &)
    var arg0 = args[0].(QPaintDevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QPoint).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPixmap", "fill", args)
  }

}

// rect()
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
    C.C_ZNK7QPixmap4rectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "rect", args)
  }

}

// trueMatrix(const class QMatrix &, int, int)
func (this *QPixmap) trueMatrix_s(args ...interface{}) () {
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
    // invoke: _ZN7QPixmap10trueMatrixERK7QMatrixii
    // invoke: QMatrix trueMatrix(const class QMatrix &, int, int)
    var arg0 = args[0].(QMatrix).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN7QPixmap10trueMatrixERK7QMatrixii(arg0, arg1, arg2)
  case 1:
    // invoke: _ZN7QPixmap10trueMatrixERK10QTransformii
    // invoke: QTransform trueMatrix(const class QTransform &, int, int)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN7QPixmap10trueMatrixERK10QTransformii(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPixmap", "trueMatrix", args)
  }

}

// detach()
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
    C.C_ZN7QPixmap6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "detach", args)
  }

}

// size()
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
    C.C_ZNK7QPixmap4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "size", args)
  }

}

// depth()
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
    C.C_ZNK7QPixmap5depthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "depth", args)
  }

}

// width()
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
    C.C_ZNK7QPixmap5widthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "width", args)
  }

}

// swap(class QPixmap &)
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
    C.C_ZN7QPixmap4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "swap", args)
  }

}

// isDetached()
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
    C.C_ZNK7QPixmap10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "isDetached", args)
  }

}

// cacheKey()
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
    C.C_ZNK7QPixmap8cacheKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "cacheKey", args)
  }

}

// setDevicePixelRatio(qreal)
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
    C.C_ZN7QPixmap19setDevicePixelRatioEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "setDevicePixelRatio", args)
  }

}

// hasAlphaChannel()
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
    C.C_ZNK7QPixmap15hasAlphaChannelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlphaChannel", args)
  }

}

// hasAlpha()
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
    C.C_ZNK7QPixmap8hasAlphaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlpha", args)
  }

}

// setMask(const class QBitmap &)
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
    C.C_ZN7QPixmap7setMaskERK7QBitmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "setMask", args)
  }

}

// createHeuristicMask(_Bool)
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
    C.C_ZNK7QPixmap19createHeuristicMaskEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "createHeuristicMask", args)
  }

}

// defaultDepth()
func (this *QPixmap) defaultDepth_s(args ...interface{}) () {
  // defaultDepth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap12defaultDepthEv
    // invoke: int defaultDepth()
    C.C_ZN7QPixmap12defaultDepthEv()
  default:
    qtrt.ErrorResolve("QPixmap", "defaultDepth", args)
  }

}

// copy(const class QRect &)
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
    C.C_ZNK7QPixmap4copyERK5QRect(this.qclsinst, arg0)
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
    C.C_ZNK7QPixmap4copyEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPixmap", "copy", args)
  }

}

// save(const class QString &, const char *, int)
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
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK7QPixmap4saveERK7QStringPKci(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK7QPixmap4saveEP9QIODevicePKci
    // invoke: bool save(class QIODevice *, const char *, int)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.C_ZNK7QPixmap4saveEP9QIODevicePKci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QPixmap", "save", args)
  }

}

// devicePixelRatio()
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
    C.C_ZNK7QPixmap16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "devicePixelRatio", args)
  }

}

// isQBitmap()
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
    C.C_ZNK7QPixmap9isQBitmapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "isQBitmap", args)
  }

}

// mask()
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
    C.C_ZNK7QPixmap4maskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "mask", args)
  }

}

// toImage()
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
    C.C_ZNK7QPixmap7toImageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "toImage", args)
  }

}

// QPixmap()
func NewQPixmap(args ...interface{}) QPixmap {
  // QPixmap()
  // QPixmap(int, int)
  // QPixmap(const class QPixmap &)
  // QPixmap(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmapC1Ev
    // invoke: void QPixmap()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPixmapC2Ev(qthis)
  case 1:
    // invoke: _ZN7QPixmapC1Eii
    // invoke: void QPixmap(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPixmapC2Eii(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN7QPixmapC1ERKS_
    // invoke: void QPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPixmapC2ERKS_(qthis, arg0)
  case 3:
    // invoke: _ZN7QPixmapC1ERK5QSize
    // invoke: void QPixmap(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN7QPixmapC2ERK5QSize(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "QPixmap", args)
  }

  return QPixmap{}
}

// isNull()
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
    C.C_ZNK7QPixmap6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "isNull", args)
  }

}

// devType()
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
    C.C_ZNK7QPixmap7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "devType", args)
  }

}

// scroll(int, int, int, int, int, int, class QRegion *)
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
    C.C_ZN7QPixmap6scrollEiiiiiiP7QRegion(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
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
    C.C_ZN7QPixmap6scrollEiiRK5QRectP7QRegion(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPixmap", "scroll", args)
  }

}

// <= body block end

