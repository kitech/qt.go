package qtgui
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QPixmap::height();
extern int32_t C_ZNK7QPixmap6heightEv(void* qthis); // 4
  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, int x, int y, int w, int h);
extern void* C_ZN7QPixmap10grabWidgetEP7QObjectiiii(void* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 2
  // proto: static QPixmap QPixmap::grabWidget(QObject * widget, const QRect & rect);
extern void* C_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect(void* arg0, void* arg1); // 4
  // proto:  QPaintEngine * QPixmap::paintEngine();
extern void* C_ZNK7QPixmap11paintEngineEv(void* qthis); // 4
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
extern void* C_ZNK7QPixmap4rectEv(void* qthis); // 4
  // proto: static QMatrix QPixmap::trueMatrix(const QMatrix & m, int w, int h);
extern void* C_ZN7QPixmap10trueMatrixERK7QMatrixii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto: static QTransform QPixmap::trueMatrix(const QTransform & m, int w, int h);
extern void* C_ZN7QPixmap10trueMatrixERK10QTransformii(void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QPixmap::detach();
extern void C_ZN7QPixmap6detachEv(void* qthis); // 4
  // proto:  QSize QPixmap::size();
extern void* C_ZNK7QPixmap4sizeEv(void* qthis); // 4
  // proto:  int QPixmap::depth();
extern int32_t C_ZNK7QPixmap5depthEv(void* qthis); // 4
  // proto:  int QPixmap::width();
extern int32_t C_ZNK7QPixmap5widthEv(void* qthis); // 4
  // proto:  void QPixmap::swap(QPixmap & other);
extern void C_ZN7QPixmap4swapERS_(void* qthis, void* arg0); // 2
  // proto:  bool QPixmap::isDetached();
extern bool C_ZNK7QPixmap10isDetachedEv(void* qthis); // 4
  // proto:  qint64 QPixmap::cacheKey();
extern int64_t C_ZNK7QPixmap8cacheKeyEv(void* qthis); // 4
  // proto:  void QPixmap::setDevicePixelRatio(qreal scaleFactor);
extern void C_ZN7QPixmap19setDevicePixelRatioEd(void* qthis, double arg0); // 4
  // proto:  bool QPixmap::hasAlphaChannel();
extern bool C_ZNK7QPixmap15hasAlphaChannelEv(void* qthis); // 4
  // proto:  bool QPixmap::hasAlpha();
extern bool C_ZNK7QPixmap8hasAlphaEv(void* qthis); // 4
  // proto:  void QPixmap::setMask(const QBitmap & );
extern void C_ZN7QPixmap7setMaskERK7QBitmap(void* qthis, void* arg0); // 4
  // proto:  QBitmap QPixmap::createHeuristicMask(bool clipTight);
extern void* C_ZNK7QPixmap19createHeuristicMaskEb(void* qthis, bool arg0); // 4
  // proto: static int QPixmap::defaultDepth();
extern int32_t C_ZN7QPixmap12defaultDepthEv(); // 4
  // proto:  QPixmap QPixmap::copy(const QRect & rect);
extern void* C_ZNK7QPixmap4copyERK5QRect(void* qthis, void* arg0); // 4
  // proto:  QPixmap QPixmap::copy(int x, int y, int width, int height);
extern void* C_ZNK7QPixmap4copyEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 2
  // proto:  bool QPixmap::save(const QString & fileName, const char * format, int quality);
extern bool C_ZNK7QPixmap4saveERK7QStringPKci(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  bool QPixmap::save(QIODevice * device, const char * format, int quality);
extern bool C_ZNK7QPixmap4saveEP9QIODevicePKci(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  qreal QPixmap::devicePixelRatio();
extern double C_ZNK7QPixmap16devicePixelRatioEv(void* qthis); // 4
  // proto:  bool QPixmap::isQBitmap();
extern bool C_ZNK7QPixmap9isQBitmapEv(void* qthis); // 4
  // proto:  QBitmap QPixmap::mask();
extern void* C_ZNK7QPixmap4maskEv(void* qthis); // 4
  // proto:  QImage QPixmap::toImage();
extern void* C_ZNK7QPixmap7toImageEv(void* qthis); // 4
  // proto:  void QPixmap::QPixmap();
extern void* C_ZN7QPixmapC2Ev(); // 3
  // proto:  void QPixmap::QPixmap(int w, int h);
extern void* C_ZN7QPixmapC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  void QPixmap::QPixmap(const QPixmap & );
extern void* C_ZN7QPixmapC2ERKS_(void* arg0); // 3
  // proto:  void QPixmap::QPixmap(const QSize & );
extern void* C_ZN7QPixmapC2ERK5QSize(void* arg0); // 3
  // proto:  bool QPixmap::isNull();
extern bool C_ZNK7QPixmap6isNullEv(void* qthis); // 4
  // proto:  int QPixmap::devType();
extern int32_t C_ZNK7QPixmap7devTypeEv(void* qthis); // 4
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
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QPixmap)=1
type QPixmap struct {
  /*qbase*/ QPaintDevice;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// height()
func (this *QPixmap) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK7QPixmap6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "height", args)
  }

  return
}

// grabWidget(class QObject *, int, int, int, int)
func (this *QPixmap) Grabwidget_S(args ...interface{}) (ret interface{}) {
  // grabWidget(class QObject *, int, int, int, int)
  // grabWidget(class QObject *, const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1][1] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap10grabWidgetEP7QObjectiiii
    // invoke: QPixmap grabWidget(class QObject *, int, int, int, int)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN7QPixmap10grabWidgetEP7QObjectiiii(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QPixmap10grabWidgetEP7QObjectRK5QRect
    // invoke: QPixmap grabWidget(class QObject *, const class QRect &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QPixmap10grabWidgetEP7QObjectRK5QRect(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "grabWidget", args)
  }

  return
}

// paintEngine()
func (this *QPixmap) Paintengine(args ...interface{}) (ret interface{}) {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    var ret0 = C.C_ZNK7QPixmap11paintEngineEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "paintEngine", args)
  }

  return
}

// handle()
func (this *QPixmap) Handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6handleEv
    // invoke: QPlatformPixmap * handle()
    C.C_ZNK7QPixmap6handleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "handle", args)
  }

  return
}

// ~QPixmap()
func (this *QPixmap) Freeqpixmap(args ...interface{}) () {
  // ~QPixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmapD0Ev
    // invoke: void ~QPixmap()
    C.C_ZN7QPixmapD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "~QPixmap", args)
  }

  return
}

// fill(const class QColor &)
func (this *QPixmap) Fill(args ...interface{}) () {
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
  vtys[2][1] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4fillERK6QColor
    // invoke: void fill(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QPixmap4fillERK6QColor(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceii
    // invoke: void fill(const class QPaintDevice *, int, int)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN7QPixmap4fillEPK12QPaintDeviceii(this.Qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint
    // invoke: void fill(const class QPaintDevice *, const class QPoint &)
    var arg0 = args[0].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN7QPixmap4fillEPK12QPaintDeviceRK6QPoint(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPixmap", "fill", args)
  }

  return
}

// rect()
func (this *QPixmap) Rect(args ...interface{}) (ret interface{}) {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4rectEv
    // invoke: QRect rect()
    var ret0 = C.C_ZNK7QPixmap4rectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "rect", args)
  }

  return
}

// trueMatrix(const class QMatrix &, int, int)
func (this *QPixmap) Truematrix_S(args ...interface{}) (ret interface{}) {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap10trueMatrixERK7QMatrixii
    // invoke: QMatrix trueMatrix(const class QMatrix &, int, int)
    var arg0 = args[0].(*QMatrix).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QPixmap10trueMatrixERK7QMatrixii(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMatrix{}) // "QMatrix"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QPixmap10trueMatrixERK10QTransformii
    // invoke: QTransform trueMatrix(const class QTransform &, int, int)
    var arg0 = args[0].(*QTransform).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QPixmap10trueMatrixERK10QTransformii(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTransform{}) // "QTransform"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "trueMatrix", args)
  }

  return
}

// detach()
func (this *QPixmap) Detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap6detachEv
    // invoke: void detach()
    C.C_ZN7QPixmap6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QPixmap", "detach", args)
  }

  return
}

// size()
func (this *QPixmap) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4sizeEv
    // invoke: QSize size()
    var ret0 = C.C_ZNK7QPixmap4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "size", args)
  }

  return
}

// depth()
func (this *QPixmap) Depth(args ...interface{}) (ret interface{}) {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap5depthEv
    // invoke: int depth()
    var ret0 = C.C_ZNK7QPixmap5depthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "depth", args)
  }

  return
}

// width()
func (this *QPixmap) Width(args ...interface{}) (ret interface{}) {
  // width()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap5widthEv
    // invoke: int width()
    var ret0 = C.C_ZNK7QPixmap5widthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "width", args)
  }

  return
}

// swap(class QPixmap &)
func (this *QPixmap) Swap(args ...interface{}) () {
  // swap(class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "QPixmap &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap4swapERS_
    // invoke: void swap(class QPixmap &)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QPixmap4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "swap", args)
  }

  return
}

// isDetached()
func (this *QPixmap) Isdetached(args ...interface{}) (ret interface{}) {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap10isDetachedEv
    // invoke: bool isDetached()
    var ret0 = C.C_ZNK7QPixmap10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "isDetached", args)
  }

  return
}

// cacheKey()
func (this *QPixmap) Cachekey(args ...interface{}) (ret interface{}) {
  // cacheKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap8cacheKeyEv
    // invoke: qint64 cacheKey()
    var ret0 = C.C_ZNK7QPixmap8cacheKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "cacheKey", args)
  }

  return
}

// setDevicePixelRatio(qreal)
func (this *QPixmap) Setdevicepixelratio(args ...interface{}) () {
  // setDevicePixelRatio(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap19setDevicePixelRatioEd
    // invoke: void setDevicePixelRatio(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN7QPixmap19setDevicePixelRatioEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "setDevicePixelRatio", args)
  }

  return
}

// hasAlphaChannel()
func (this *QPixmap) Hasalphachannel(args ...interface{}) (ret interface{}) {
  // hasAlphaChannel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap15hasAlphaChannelEv
    // invoke: bool hasAlphaChannel()
    var ret0 = C.C_ZNK7QPixmap15hasAlphaChannelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlphaChannel", args)
  }

  return
}

// hasAlpha()
func (this *QPixmap) Hasalpha(args ...interface{}) (ret interface{}) {
  // hasAlpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap8hasAlphaEv
    // invoke: bool hasAlpha()
    var ret0 = C.C_ZNK7QPixmap8hasAlphaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "hasAlpha", args)
  }

  return
}

// setMask(const class QBitmap &)
func (this *QPixmap) Setmask(args ...interface{}) () {
  // setMask(const class QBitmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBitmap{}) // "const QBitmap &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap7setMaskERK7QBitmap
    // invoke: void setMask(const class QBitmap &)
    var arg0 = args[0].(*QBitmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QPixmap7setMaskERK7QBitmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPixmap", "setMask", args)
  }

  return
}

// createHeuristicMask(_Bool)
func (this *QPixmap) Createheuristicmask(args ...interface{}) (ret interface{}) {
  // createHeuristicMask(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap19createHeuristicMaskEb
    // invoke: QBitmap createHeuristicMask(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QPixmap19createHeuristicMaskEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitmap{}) // "QBitmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "createHeuristicMask", args)
  }

  return
}

// defaultDepth()
func (this *QPixmap) Defaultdepth_S(args ...interface{}) (ret interface{}) {
  // defaultDepth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap12defaultDepthEv
    // invoke: int defaultDepth()
    var ret0 = C.C_ZN7QPixmap12defaultDepthEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "defaultDepth", args)
  }

  return
}

// copy(const class QRect &)
func (this *QPixmap) Copy(args ...interface{}) (ret interface{}) {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4copyERK5QRect
    // invoke: QPixmap copy(const class QRect &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QPixmap4copyERK5QRect(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QPixmap4copyEiiii
    // invoke: QPixmap copy(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QPixmap4copyEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "copy", args)
  }

  return
}

// save(const class QString &, const char *, int)
func (this *QPixmap) Save(args ...interface{}) (ret interface{}) {
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
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4saveERK7QStringPKci
    // invoke: bool save(const class QString &, const char *, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QPixmap4saveERK7QStringPKci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QPixmap4saveEP9QIODevicePKci
    // invoke: bool save(class QIODevice *, const char *, int)
    var arg0 = args[0].(*qtcore.QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QPixmap4saveEP9QIODevicePKci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "save", args)
  }

  return
}

// devicePixelRatio()
func (this *QPixmap) Devicepixelratio(args ...interface{}) (ret interface{}) {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap16devicePixelRatioEv
    // invoke: qreal devicePixelRatio()
    var ret0 = C.C_ZNK7QPixmap16devicePixelRatioEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "devicePixelRatio", args)
  }

  return
}

// isQBitmap()
func (this *QPixmap) Isqbitmap(args ...interface{}) (ret interface{}) {
  // isQBitmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap9isQBitmapEv
    // invoke: bool isQBitmap()
    var ret0 = C.C_ZNK7QPixmap9isQBitmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "isQBitmap", args)
  }

  return
}

// mask()
func (this *QPixmap) Mask(args ...interface{}) (ret interface{}) {
  // mask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap4maskEv
    // invoke: QBitmap mask()
    var ret0 = C.C_ZNK7QPixmap4maskEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBitmap{}) // "QBitmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "mask", args)
  }

  return
}

// toImage()
func (this *QPixmap) Toimage(args ...interface{}) (ret interface{}) {
  // toImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap7toImageEv
    // invoke: QImage toImage()
    var ret0 = C.C_ZNK7QPixmap7toImageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "toImage", args)
  }

  return
}

// QPixmap()
func NewQPixmap(args ...interface{}) *QPixmap {
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
  vtys[3][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmapC1Ev
    // invoke: void QPixmap()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPixmapC2Ev()
    return &QPixmap{Qclsinst:qthis}
  case 1:
    // invoke: _ZN7QPixmapC1Eii
    // invoke: void QPixmap(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPixmapC2Eii(arg0, arg1)
    return &QPixmap{Qclsinst:qthis}
  case 2:
    // invoke: _ZN7QPixmapC1ERKS_
    // invoke: void QPixmap(const class QPixmap &)
    var arg0 = args[0].(*QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPixmapC2ERKS_(arg0)
    return &QPixmap{Qclsinst:qthis}
  case 3:
    // invoke: _ZN7QPixmapC1ERK5QSize
    // invoke: void QPixmap(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QPixmapC2ERK5QSize(arg0)
    return &QPixmap{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPixmap", "QPixmap", args)
  }

  return nil // QPixmap{Qclsinst:qthis}
}

// isNull()
func (this *QPixmap) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK7QPixmap6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "isNull", args)
  }

  return
}

// devType()
func (this *QPixmap) Devtype(args ...interface{}) (ret interface{}) {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QPixmap7devTypeEv
    // invoke: int devType()
    var ret0 = C.C_ZNK7QPixmap7devTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QPixmap", "devType", args)
  }

  return
}

// scroll(int, int, int, int, int, int, class QRegion *)
func (this *QPixmap) Scroll(args ...interface{}) () {
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
  vtys[1][2] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[1][3] = reflect.TypeOf(QRegion{}) // "QRegion *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QPixmap6scrollEiiiiiiP7QRegion
    // invoke: void scroll(int, int, int, int, int, int, class QRegion *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var arg5 = C.int32_t(qtrt.PrimConv(args[5], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(*QRegion).Qclsinst
    if false {fmt.Println(arg6)}
    C.C_ZN7QPixmap6scrollEiiiiiiP7QRegion(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 1:
    // invoke: _ZN7QPixmap6scrollEiiRK5QRectP7QRegion
    // invoke: void scroll(int, int, const class QRect &, class QRegion *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QRegion).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZN7QPixmap6scrollEiiRK5QRectP7QRegion(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QPixmap", "scroll", args)
  }

  return
}

// <= body block end

