package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtGui/qfontmetrics.h
// dst-file: /src/gui/qfontmetrics.go
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
  // proto:  void QFontMetrics::~QFontMetrics();
extern void C_ZN12QFontMetricsD2Ev(void* qthis); // 4
  // proto:  int QFontMetrics::height();
extern int32_t C_ZNK12QFontMetrics6heightEv(void* qthis); // 4
  // proto:  int QFontMetrics::ascent();
extern int32_t C_ZNK12QFontMetrics6ascentEv(void* qthis); // 4
  // proto:  QSize QFontMetrics::size(int flags, const QString & str, int tabstops, int * tabarray);
extern void* C_ZNK12QFontMetrics4sizeEiRK7QStringiPi(void* qthis, int32_t arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto:  int QFontMetrics::leftBearing(QChar );
extern int32_t C_ZNK12QFontMetrics11leftBearingE5QChar(void* qthis, void* arg0); // 4
  // proto:  int QFontMetrics::minRightBearing();
extern int32_t C_ZNK12QFontMetrics15minRightBearingEv(void* qthis); // 4
  // proto:  QRect QFontMetrics::boundingRect(int x, int y, int w, int h, int flags, const QString & text, int tabstops, int * tabarray);
extern void* C_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, void* arg5, int32_t arg6, void* arg7); // 2
  // proto:  QRect QFontMetrics::boundingRect(const QRect & r, int flags, const QString & text, int tabstops, int * tabarray);
extern void* C_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3, void* arg4); // 4
  // proto:  QRect QFontMetrics::boundingRect(QChar );
extern void* C_ZNK12QFontMetrics12boundingRectE5QChar(void* qthis, void* arg0); // 4
  // proto:  QRect QFontMetrics::boundingRect(const QString & text);
extern void* C_ZNK12QFontMetrics12boundingRectERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QFontMetrics::QFontMetrics(const QFont & );
extern void* C_ZN12QFontMetricsC2ERK5QFont(void* arg0); // 3
  // proto:  void QFontMetrics::QFontMetrics(const QFont & , QPaintDevice * pd);
extern void* C_ZN12QFontMetricsC2ERK5QFontP12QPaintDevice(void* arg0, void* arg1); // 3
  // proto:  void QFontMetrics::QFontMetrics(const QFontMetrics & );
extern void* C_ZN12QFontMetricsC2ERKS_(void* arg0); // 3
  // proto:  int QFontMetrics::leading();
extern int32_t C_ZNK12QFontMetrics7leadingEv(void* qthis); // 4
  // proto:  QRect QFontMetrics::tightBoundingRect(const QString & text);
extern void* C_ZNK12QFontMetrics17tightBoundingRectERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QFontMetrics::width(QChar );
extern int32_t C_ZNK12QFontMetrics5widthE5QChar(void* qthis, void* arg0); // 4
  // proto:  int QFontMetrics::width(const QString & , int len);
extern int32_t C_ZNK12QFontMetrics5widthERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QFontMetrics::width(const QString & , int len, int flags);
extern int32_t C_ZNK12QFontMetrics5widthERK7QStringii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QFontMetrics::swap(QFontMetrics & other);
extern void C_ZN12QFontMetrics4swapERS_(void* qthis, void* arg0); // 2
  // proto:  int QFontMetrics::lineSpacing();
extern int32_t C_ZNK12QFontMetrics11lineSpacingEv(void* qthis); // 4
  // proto:  int QFontMetrics::underlinePos();
extern int32_t C_ZNK12QFontMetrics12underlinePosEv(void* qthis); // 4
  // proto:  int QFontMetrics::charWidth(const QString & str, int pos);
extern int32_t C_ZNK12QFontMetrics9charWidthERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QFontMetrics::maxWidth();
extern int32_t C_ZNK12QFontMetrics8maxWidthEv(void* qthis); // 4
  // proto:  int QFontMetrics::minLeftBearing();
extern int32_t C_ZNK12QFontMetrics14minLeftBearingEv(void* qthis); // 4
  // proto:  bool QFontMetrics::inFont(QChar );
extern bool C_ZNK12QFontMetrics6inFontE5QChar(void* qthis, void* arg0); // 4
  // proto:  int QFontMetrics::descent();
extern int32_t C_ZNK12QFontMetrics7descentEv(void* qthis); // 4
  // proto:  int QFontMetrics::averageCharWidth();
extern int32_t C_ZNK12QFontMetrics16averageCharWidthEv(void* qthis); // 4
  // proto:  int QFontMetrics::overlinePos();
extern int32_t C_ZNK12QFontMetrics11overlinePosEv(void* qthis); // 4
  // proto:  int QFontMetrics::strikeOutPos();
extern int32_t C_ZNK12QFontMetrics12strikeOutPosEv(void* qthis); // 4
  // proto:  int QFontMetrics::xHeight();
extern int32_t C_ZNK12QFontMetrics7xHeightEv(void* qthis); // 4
  // proto:  bool QFontMetrics::inFontUcs4(uint ucs4);
extern bool C_ZNK12QFontMetrics10inFontUcs4Ej(void* qthis, int32_t arg0); // 4
  // proto:  int QFontMetrics::lineWidth();
extern int32_t C_ZNK12QFontMetrics9lineWidthEv(void* qthis); // 4
  // proto:  int QFontMetrics::rightBearing(QChar );
extern int32_t C_ZNK12QFontMetrics12rightBearingE5QChar(void* qthis, void* arg0); // 4
  // proto:  QRectF QFontMetricsF::boundingRect(const QString & string);
extern void* C_ZNK13QFontMetricsF12boundingRectERK7QString(void* qthis, void* arg0); // 4
  // proto:  QRectF QFontMetricsF::boundingRect(const QRectF & r, int flags, const QString & string, int tabstops, int * tabarray);
extern void* C_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3, void* arg4); // 4
  // proto:  QRectF QFontMetricsF::boundingRect(QChar );
extern void* C_ZNK13QFontMetricsF12boundingRectE5QChar(void* qthis, void* arg0); // 4
  // proto:  qreal QFontMetricsF::height();
extern double C_ZNK13QFontMetricsF6heightEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::ascent();
extern double C_ZNK13QFontMetricsF6ascentEv(void* qthis); // 4
  // proto:  QSizeF QFontMetricsF::size(int flags, const QString & str, int tabstops, int * tabarray);
extern void* C_ZNK13QFontMetricsF4sizeEiRK7QStringiPi(void* qthis, int32_t arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto:  qreal QFontMetricsF::leftBearing(QChar );
extern double C_ZNK13QFontMetricsF11leftBearingE5QChar(void* qthis, void* arg0); // 4
  // proto:  qreal QFontMetricsF::minRightBearing();
extern double C_ZNK13QFontMetricsF15minRightBearingEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::descent();
extern double C_ZNK13QFontMetricsF7descentEv(void* qthis); // 4
  // proto:  void QFontMetricsF::~QFontMetricsF();
extern void C_ZN13QFontMetricsFD2Ev(void* qthis); // 4
  // proto:  QRectF QFontMetricsF::tightBoundingRect(const QString & text);
extern void* C_ZNK13QFontMetricsF17tightBoundingRectERK7QString(void* qthis, void* arg0); // 4
  // proto:  qreal QFontMetricsF::width(const QString & string);
extern double C_ZNK13QFontMetricsF5widthERK7QString(void* qthis, void* arg0); // 4
  // proto:  qreal QFontMetricsF::width(QChar );
extern double C_ZNK13QFontMetricsF5widthE5QChar(void* qthis, void* arg0); // 4
  // proto:  void QFontMetricsF::swap(QFontMetricsF & other);
extern void C_ZN13QFontMetricsF4swapERS_(void* qthis, void* arg0); // 2
  // proto:  qreal QFontMetricsF::lineSpacing();
extern double C_ZNK13QFontMetricsF11lineSpacingEv(void* qthis); // 4
  // proto:  void QFontMetricsF::QFontMetricsF(const QFont & );
extern void* C_ZN13QFontMetricsFC2ERK5QFont(void* arg0); // 3
  // proto:  void QFontMetricsF::QFontMetricsF(const QFontMetrics & );
extern void* C_ZN13QFontMetricsFC2ERK12QFontMetrics(void* arg0); // 3
  // proto:  void QFontMetricsF::QFontMetricsF(const QFont & , QPaintDevice * pd);
extern void* C_ZN13QFontMetricsFC2ERK5QFontP12QPaintDevice(void* arg0, void* arg1); // 3
  // proto:  void QFontMetricsF::QFontMetricsF(const QFontMetricsF & );
extern void* C_ZN13QFontMetricsFC2ERKS_(void* arg0); // 3
  // proto:  qreal QFontMetricsF::underlinePos();
extern double C_ZNK13QFontMetricsF12underlinePosEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::maxWidth();
extern double C_ZNK13QFontMetricsF8maxWidthEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::minLeftBearing();
extern double C_ZNK13QFontMetricsF14minLeftBearingEv(void* qthis); // 4
  // proto:  bool QFontMetricsF::inFont(QChar );
extern bool C_ZNK13QFontMetricsF6inFontE5QChar(void* qthis, void* arg0); // 4
  // proto:  qreal QFontMetricsF::averageCharWidth();
extern double C_ZNK13QFontMetricsF16averageCharWidthEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::overlinePos();
extern double C_ZNK13QFontMetricsF11overlinePosEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::strikeOutPos();
extern double C_ZNK13QFontMetricsF12strikeOutPosEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::xHeight();
extern double C_ZNK13QFontMetricsF7xHeightEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::leading();
extern double C_ZNK13QFontMetricsF7leadingEv(void* qthis); // 4
  // proto:  bool QFontMetricsF::inFontUcs4(uint ucs4);
extern bool C_ZNK13QFontMetricsF10inFontUcs4Ej(void* qthis, int32_t arg0); // 4
  // proto:  qreal QFontMetricsF::lineWidth();
extern double C_ZNK13QFontMetricsF9lineWidthEv(void* qthis); // 4
  // proto:  qreal QFontMetricsF::rightBearing(QChar );
extern double C_ZNK13QFontMetricsF12rightBearingE5QChar(void* qthis, void* arg0); // 4
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

// class sizeof(QFontMetrics)=1
type QFontMetrics struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFontMetricsF)=1
type QFontMetricsF struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QFontMetrics()
func (this *QFontMetrics) Freeqfontmetrics(args ...interface{}) () {
  // ~QFontMetrics()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QFontMetricsD0Ev
    // invoke: void ~QFontMetrics()
    C.C_ZN12QFontMetricsD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "~QFontMetrics", args)
  }

  return
}

// height()
func (this *QFontMetrics) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics6heightEv
    // invoke: int height()
    var ret0 = C.C_ZNK12QFontMetrics6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "height", args)
  }

  return
}

// ascent()
func (this *QFontMetrics) Ascent(args ...interface{}) (ret interface{}) {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics6ascentEv
    // invoke: int ascent()
    var ret0 = C.C_ZNK12QFontMetrics6ascentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "ascent", args)
  }

  return
}

// size(int, const class QString &, int, int *)
func (this *QFontMetrics) Size(args ...interface{}) (ret interface{}) {
  // size(int, const class QString &, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics4sizeEiRK7QStringiPi
    // invoke: QSize size(int, const class QString &, int, int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK12QFontMetrics4sizeEiRK7QStringiPi(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "size", args)
  }

  return
}

// leftBearing(class QChar)
func (this *QFontMetrics) Leftbearing(args ...interface{}) (ret interface{}) {
  // leftBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics11leftBearingE5QChar
    // invoke: int leftBearing(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics11leftBearingE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "leftBearing", args)
  }

  return
}

// minRightBearing()
func (this *QFontMetrics) Minrightbearing(args ...interface{}) (ret interface{}) {
  // minRightBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics15minRightBearingEv
    // invoke: int minRightBearing()
    var ret0 = C.C_ZNK12QFontMetrics15minRightBearingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "minRightBearing", args)
  }

  return
}

// boundingRect(int, int, int, int, int, const class QString &, int, int *)
func (this *QFontMetrics) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect(int, int, int, int, int, const class QString &, int, int *)
  // boundingRect(const class QRect &, int, const class QString &, int, int *)
  // boundingRect(class QChar)
  // boundingRect(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"
  vtys[0][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][6] = qtrt.Int32Ty(false) // "int"
  vtys[0][7] = qtrt.Int32Ty(true) // "int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(true) // "int *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi
    // invoke: QRect boundingRect(int, int, int, int, int, const class QString &, int, int *)
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
    var arg5 = args[5].(*QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(qtrt.PrimConv(args[6], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg6)}
    var arg7 = (unsafe.Pointer)(args[7].(*int32))
    if false {fmt.Println(arg7)}
    var ret0 = C.C_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi
    // invoke: QRect boundingRect(const class QRect &, int, const class QString &, int, int *)
    var arg0 = args[0].(*QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = (unsafe.Pointer)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK12QFontMetrics12boundingRectE5QChar
    // invoke: QRect boundingRect(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics12boundingRectE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK12QFontMetrics12boundingRectERK7QString
    // invoke: QRect boundingRect(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics12boundingRectERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "boundingRect", args)
  }

  return
}

// QFontMetrics(const class QFont &)
func NewQFontMetrics(args ...interface{}) *QFontMetrics {
  // QFontMetrics(const class QFont &)
  // QFontMetrics(const class QFont &, class QPaintDevice *)
  // QFontMetrics(const class QFontMetrics &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1][1] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QFontMetrics{}) // "const QFontMetrics &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QFontMetricsC1ERK5QFont
    // invoke: void QFontMetrics(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QFontMetricsC2ERK5QFont(arg0)
    return &QFontMetrics{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QFontMetricsC1ERK5QFontP12QPaintDevice
    // invoke: void QFontMetrics(const class QFont &, class QPaintDevice *)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QFontMetricsC2ERK5QFontP12QPaintDevice(arg0, arg1)
    return &QFontMetrics{Qclsinst:qthis}
  case 2:
    // invoke: _ZN12QFontMetricsC1ERKS_
    // invoke: void QFontMetrics(const class QFontMetrics &)
    var arg0 = args[0].(*QFontMetrics).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QFontMetricsC2ERKS_(arg0)
    return &QFontMetrics{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFontMetrics", "QFontMetrics", args)
  }

  return nil // QFontMetrics{Qclsinst:qthis}
}

// leading()
func (this *QFontMetrics) Leading(args ...interface{}) (ret interface{}) {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics7leadingEv
    // invoke: int leading()
    var ret0 = C.C_ZNK12QFontMetrics7leadingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "leading", args)
  }

  return
}

// tightBoundingRect(const class QString &)
func (this *QFontMetrics) Tightboundingrect(args ...interface{}) (ret interface{}) {
  // tightBoundingRect(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics17tightBoundingRectERK7QString
    // invoke: QRect tightBoundingRect(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics17tightBoundingRectERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "tightBoundingRect", args)
  }

  return
}

// width(class QChar)
func (this *QFontMetrics) Width(args ...interface{}) (ret interface{}) {
  // width(class QChar)
  // width(const class QString &, int)
  // width(const class QString &, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics5widthE5QChar
    // invoke: int width(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics5widthE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringi
    // invoke: int width(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK12QFontMetrics5widthERK7QStringi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringii
    // invoke: int width(const class QString &, int, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK12QFontMetrics5widthERK7QStringii(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "width", args)
  }

  return
}

// swap(class QFontMetrics &)
func (this *QFontMetrics) Swap(args ...interface{}) () {
  // swap(class QFontMetrics &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontMetrics{}) // "QFontMetrics &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QFontMetrics4swapERS_
    // invoke: void swap(class QFontMetrics &)
    var arg0 = args[0].(*QFontMetrics).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QFontMetrics4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "swap", args)
  }

  return
}

// lineSpacing()
func (this *QFontMetrics) Linespacing(args ...interface{}) (ret interface{}) {
  // lineSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics11lineSpacingEv
    // invoke: int lineSpacing()
    var ret0 = C.C_ZNK12QFontMetrics11lineSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "lineSpacing", args)
  }

  return
}

// underlinePos()
func (this *QFontMetrics) Underlinepos(args ...interface{}) (ret interface{}) {
  // underlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12underlinePosEv
    // invoke: int underlinePos()
    var ret0 = C.C_ZNK12QFontMetrics12underlinePosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "underlinePos", args)
  }

  return
}

// charWidth(const class QString &, int)
func (this *QFontMetrics) Charwidth(args ...interface{}) (ret interface{}) {
  // charWidth(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics9charWidthERK7QStringi
    // invoke: int charWidth(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK12QFontMetrics9charWidthERK7QStringi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "charWidth", args)
  }

  return
}

// maxWidth()
func (this *QFontMetrics) Maxwidth(args ...interface{}) (ret interface{}) {
  // maxWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics8maxWidthEv
    // invoke: int maxWidth()
    var ret0 = C.C_ZNK12QFontMetrics8maxWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "maxWidth", args)
  }

  return
}

// minLeftBearing()
func (this *QFontMetrics) Minleftbearing(args ...interface{}) (ret interface{}) {
  // minLeftBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics14minLeftBearingEv
    // invoke: int minLeftBearing()
    var ret0 = C.C_ZNK12QFontMetrics14minLeftBearingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "minLeftBearing", args)
  }

  return
}

// inFont(class QChar)
func (this *QFontMetrics) Infont(args ...interface{}) (ret interface{}) {
  // inFont(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics6inFontE5QChar
    // invoke: bool inFont(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics6inFontE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "inFont", args)
  }

  return
}

// descent()
func (this *QFontMetrics) Descent(args ...interface{}) (ret interface{}) {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics7descentEv
    // invoke: int descent()
    var ret0 = C.C_ZNK12QFontMetrics7descentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "descent", args)
  }

  return
}

// averageCharWidth()
func (this *QFontMetrics) Averagecharwidth(args ...interface{}) (ret interface{}) {
  // averageCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics16averageCharWidthEv
    // invoke: int averageCharWidth()
    var ret0 = C.C_ZNK12QFontMetrics16averageCharWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "averageCharWidth", args)
  }

  return
}

// overlinePos()
func (this *QFontMetrics) Overlinepos(args ...interface{}) (ret interface{}) {
  // overlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics11overlinePosEv
    // invoke: int overlinePos()
    var ret0 = C.C_ZNK12QFontMetrics11overlinePosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "overlinePos", args)
  }

  return
}

// strikeOutPos()
func (this *QFontMetrics) Strikeoutpos(args ...interface{}) (ret interface{}) {
  // strikeOutPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12strikeOutPosEv
    // invoke: int strikeOutPos()
    var ret0 = C.C_ZNK12QFontMetrics12strikeOutPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "strikeOutPos", args)
  }

  return
}

// xHeight()
func (this *QFontMetrics) Xheight(args ...interface{}) (ret interface{}) {
  // xHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics7xHeightEv
    // invoke: int xHeight()
    var ret0 = C.C_ZNK12QFontMetrics7xHeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "xHeight", args)
  }

  return
}

// inFontUcs4(uint)
func (this *QFontMetrics) Infontucs4(args ...interface{}) (ret interface{}) {
  // inFontUcs4(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics10inFontUcs4Ej
    // invoke: bool inFontUcs4(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics10inFontUcs4Ej(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "inFontUcs4", args)
  }

  return
}

// lineWidth()
func (this *QFontMetrics) Linewidth(args ...interface{}) (ret interface{}) {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics9lineWidthEv
    // invoke: int lineWidth()
    var ret0 = C.C_ZNK12QFontMetrics9lineWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "lineWidth", args)
  }

  return
}

// rightBearing(class QChar)
func (this *QFontMetrics) Rightbearing(args ...interface{}) (ret interface{}) {
  // rightBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12rightBearingE5QChar
    // invoke: int rightBearing(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QFontMetrics12rightBearingE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetrics", "rightBearing", args)
  }

  return
}

// boundingRect(const class QString &)
func (this *QFontMetricsF) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect(const class QString &)
  // boundingRect(const class QRectF &, int, const class QString &, int, int *)
  // boundingRect(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[1][4] = qtrt.Int32Ty(true) // "int *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK7QString
    // invoke: QRectF boundingRect(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF12boundingRectERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi
    // invoke: QRectF boundingRect(const class QRectF &, int, const class QString &, int, int *)
    var arg0 = args[0].(*QRectF).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = (unsafe.Pointer)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK13QFontMetricsF12boundingRectE5QChar
    // invoke: QRectF boundingRect(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF12boundingRectE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "boundingRect", args)
  }

  return
}

// height()
func (this *QFontMetricsF) Height(args ...interface{}) (ret interface{}) {
  // height()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF6heightEv
    // invoke: qreal height()
    var ret0 = C.C_ZNK13QFontMetricsF6heightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "height", args)
  }

  return
}

// ascent()
func (this *QFontMetricsF) Ascent(args ...interface{}) (ret interface{}) {
  // ascent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF6ascentEv
    // invoke: qreal ascent()
    var ret0 = C.C_ZNK13QFontMetricsF6ascentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "ascent", args)
  }

  return
}

// size(int, const class QString &, int, int *)
func (this *QFontMetricsF) Size(args ...interface{}) (ret interface{}) {
  // size(int, const class QString &, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF4sizeEiRK7QStringiPi
    // invoke: QSizeF size(int, const class QString &, int, int *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK13QFontMetricsF4sizeEiRK7QStringiPi(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "size", args)
  }

  return
}

// leftBearing(class QChar)
func (this *QFontMetricsF) Leftbearing(args ...interface{}) (ret interface{}) {
  // leftBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF11leftBearingE5QChar
    // invoke: qreal leftBearing(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF11leftBearingE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "leftBearing", args)
  }

  return
}

// minRightBearing()
func (this *QFontMetricsF) Minrightbearing(args ...interface{}) (ret interface{}) {
  // minRightBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF15minRightBearingEv
    // invoke: qreal minRightBearing()
    var ret0 = C.C_ZNK13QFontMetricsF15minRightBearingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "minRightBearing", args)
  }

  return
}

// descent()
func (this *QFontMetricsF) Descent(args ...interface{}) (ret interface{}) {
  // descent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF7descentEv
    // invoke: qreal descent()
    var ret0 = C.C_ZNK13QFontMetricsF7descentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "descent", args)
  }

  return
}

// ~QFontMetricsF()
func (this *QFontMetricsF) Freeqfontmetricsf(args ...interface{}) () {
  // ~QFontMetricsF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontMetricsFD0Ev
    // invoke: void ~QFontMetricsF()
    C.C_ZN13QFontMetricsFD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "~QFontMetricsF", args)
  }

  return
}

// tightBoundingRect(const class QString &)
func (this *QFontMetricsF) Tightboundingrect(args ...interface{}) (ret interface{}) {
  // tightBoundingRect(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF17tightBoundingRectERK7QString
    // invoke: QRectF tightBoundingRect(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF17tightBoundingRectERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRectF{}) // "QRectF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "tightBoundingRect", args)
  }

  return
}

// width(const class QString &)
func (this *QFontMetricsF) Width(args ...interface{}) (ret interface{}) {
  // width(const class QString &)
  // width(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF5widthERK7QString
    // invoke: qreal width(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF5widthERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK13QFontMetricsF5widthE5QChar
    // invoke: qreal width(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF5widthE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "width", args)
  }

  return
}

// swap(class QFontMetricsF &)
func (this *QFontMetricsF) Swap(args ...interface{}) () {
  // swap(class QFontMetricsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontMetricsF{}) // "QFontMetricsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontMetricsF4swapERS_
    // invoke: void swap(class QFontMetricsF &)
    var arg0 = args[0].(*QFontMetricsF).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QFontMetricsF4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "swap", args)
  }

  return
}

// lineSpacing()
func (this *QFontMetricsF) Linespacing(args ...interface{}) (ret interface{}) {
  // lineSpacing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF11lineSpacingEv
    // invoke: qreal lineSpacing()
    var ret0 = C.C_ZNK13QFontMetricsF11lineSpacingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "lineSpacing", args)
  }

  return
}

// QFontMetricsF(const class QFont &)
func NewQFontMetricsF(args ...interface{}) *QFontMetricsF {
  // QFontMetricsF(const class QFont &)
  // QFontMetricsF(const class QFontMetrics &)
  // QFontMetricsF(const class QFont &, class QPaintDevice *)
  // QFontMetricsF(const class QFontMetricsF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFontMetrics{}) // "const QFontMetrics &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QFont{}) // "const QFont &"
  vtys[2][1] = reflect.TypeOf(QPaintDevice{}) // "QPaintDevice *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QFontMetricsF{}) // "const QFontMetricsF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QFontMetricsFC1ERK5QFont
    // invoke: void QFontMetricsF(const class QFont &)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QFontMetricsFC2ERK5QFont(arg0)
    return &QFontMetricsF{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QFontMetricsFC1ERK12QFontMetrics
    // invoke: void QFontMetricsF(const class QFontMetrics &)
    var arg0 = args[0].(*QFontMetrics).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QFontMetricsFC2ERK12QFontMetrics(arg0)
    return &QFontMetricsF{Qclsinst:qthis}
  case 2:
    // invoke: _ZN13QFontMetricsFC1ERK5QFontP12QPaintDevice
    // invoke: void QFontMetricsF(const class QFont &, class QPaintDevice *)
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QPaintDevice).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QFontMetricsFC2ERK5QFontP12QPaintDevice(arg0, arg1)
    return &QFontMetricsF{Qclsinst:qthis}
  case 3:
    // invoke: _ZN13QFontMetricsFC1ERKS_
    // invoke: void QFontMetricsF(const class QFontMetricsF &)
    var arg0 = args[0].(*QFontMetricsF).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QFontMetricsFC2ERKS_(arg0)
    return &QFontMetricsF{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QFontMetricsF", "QFontMetricsF", args)
  }

  return nil // QFontMetricsF{Qclsinst:qthis}
}

// underlinePos()
func (this *QFontMetricsF) Underlinepos(args ...interface{}) (ret interface{}) {
  // underlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12underlinePosEv
    // invoke: qreal underlinePos()
    var ret0 = C.C_ZNK13QFontMetricsF12underlinePosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "underlinePos", args)
  }

  return
}

// maxWidth()
func (this *QFontMetricsF) Maxwidth(args ...interface{}) (ret interface{}) {
  // maxWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF8maxWidthEv
    // invoke: qreal maxWidth()
    var ret0 = C.C_ZNK13QFontMetricsF8maxWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "maxWidth", args)
  }

  return
}

// minLeftBearing()
func (this *QFontMetricsF) Minleftbearing(args ...interface{}) (ret interface{}) {
  // minLeftBearing()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF14minLeftBearingEv
    // invoke: qreal minLeftBearing()
    var ret0 = C.C_ZNK13QFontMetricsF14minLeftBearingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "minLeftBearing", args)
  }

  return
}

// inFont(class QChar)
func (this *QFontMetricsF) Infont(args ...interface{}) (ret interface{}) {
  // inFont(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF6inFontE5QChar
    // invoke: bool inFont(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF6inFontE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "inFont", args)
  }

  return
}

// averageCharWidth()
func (this *QFontMetricsF) Averagecharwidth(args ...interface{}) (ret interface{}) {
  // averageCharWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF16averageCharWidthEv
    // invoke: qreal averageCharWidth()
    var ret0 = C.C_ZNK13QFontMetricsF16averageCharWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "averageCharWidth", args)
  }

  return
}

// overlinePos()
func (this *QFontMetricsF) Overlinepos(args ...interface{}) (ret interface{}) {
  // overlinePos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF11overlinePosEv
    // invoke: qreal overlinePos()
    var ret0 = C.C_ZNK13QFontMetricsF11overlinePosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "overlinePos", args)
  }

  return
}

// strikeOutPos()
func (this *QFontMetricsF) Strikeoutpos(args ...interface{}) (ret interface{}) {
  // strikeOutPos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12strikeOutPosEv
    // invoke: qreal strikeOutPos()
    var ret0 = C.C_ZNK13QFontMetricsF12strikeOutPosEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "strikeOutPos", args)
  }

  return
}

// xHeight()
func (this *QFontMetricsF) Xheight(args ...interface{}) (ret interface{}) {
  // xHeight()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF7xHeightEv
    // invoke: qreal xHeight()
    var ret0 = C.C_ZNK13QFontMetricsF7xHeightEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "xHeight", args)
  }

  return
}

// leading()
func (this *QFontMetricsF) Leading(args ...interface{}) (ret interface{}) {
  // leading()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF7leadingEv
    // invoke: qreal leading()
    var ret0 = C.C_ZNK13QFontMetricsF7leadingEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "leading", args)
  }

  return
}

// inFontUcs4(uint)
func (this *QFontMetricsF) Infontucs4(args ...interface{}) (ret interface{}) {
  // inFontUcs4(uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF10inFontUcs4Ej
    // invoke: bool inFontUcs4(uint)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF10inFontUcs4Ej(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "inFontUcs4", args)
  }

  return
}

// lineWidth()
func (this *QFontMetricsF) Linewidth(args ...interface{}) (ret interface{}) {
  // lineWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF9lineWidthEv
    // invoke: qreal lineWidth()
    var ret0 = C.C_ZNK13QFontMetricsF9lineWidthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "lineWidth", args)
  }

  return
}

// rightBearing(class QChar)
func (this *QFontMetricsF) Rightbearing(args ...interface{}) (ret interface{}) {
  // rightBearing(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12rightBearingE5QChar
    // invoke: qreal rightBearing(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QFontMetricsF12rightBearingE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFontMetricsF", "rightBearing", args)
  }

  return
}

// <= body block end

