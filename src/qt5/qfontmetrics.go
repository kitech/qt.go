package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  int QFontMetrics::maxWidth();
extern void _ZNK12QFontMetrics8maxWidthEv(void* qthis);
  // proto:  void QFontMetrics::~QFontMetrics();
extern void _ZN12QFontMetricsD0Ev(void* qthis);
  // proto:  int QFontMetrics::lineWidth();
extern void _ZNK12QFontMetrics9lineWidthEv(void* qthis);
  // proto:  QRect QFontMetrics::boundingRect(const QRect & r, int flags, const QString & text, int tabstops, int * tabarray);
extern void _ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t* arg4);
  // proto:  int QFontMetrics::minLeftBearing();
extern void _ZNK12QFontMetrics14minLeftBearingEv(void* qthis);
  // proto:  int QFontMetrics::rightBearing(QChar );
extern void _ZNK12QFontMetrics12rightBearingE5QChar(void* qthis, void* arg0);
  // proto:  int QFontMetrics::ascent();
extern void _ZNK12QFontMetrics6ascentEv(void* qthis);
  // proto:  QSize QFontMetrics::size(int flags, const QString & str, int tabstops, int * tabarray);
extern void _ZNK12QFontMetrics4sizeEiRK7QStringiPi(void* qthis, int32_t arg0, void* arg1, int32_t arg2, int32_t* arg3);
  // proto:  int QFontMetrics::overlinePos();
extern void _ZNK12QFontMetrics11overlinePosEv(void* qthis);
  // proto:  int QFontMetrics::leading();
extern void _ZNK12QFontMetrics7leadingEv(void* qthis);
  // proto:  QRect QFontMetrics::tightBoundingRect(const QString & text);
extern void _ZNK12QFontMetrics17tightBoundingRectERK7QString(void* qthis, void* arg0);
  // proto:  int QFontMetrics::averageCharWidth();
extern void _ZNK12QFontMetrics16averageCharWidthEv(void* qthis);
  // proto:  int QFontMetrics::underlinePos();
extern void _ZNK12QFontMetrics12underlinePosEv(void* qthis);
  // proto:  bool QFontMetrics::inFont(QChar );
extern void _ZNK12QFontMetrics6inFontE5QChar(void* qthis, void* arg0);
  // proto:  int QFontMetrics::height();
extern void _ZNK12QFontMetrics6heightEv(void* qthis);
  // proto:  int QFontMetrics::width(QChar );
extern void _ZNK12QFontMetrics5widthE5QChar(void* qthis, void* arg0);
  // proto:  QRect QFontMetrics::boundingRect(const QString & text);
extern void _ZNK12QFontMetrics12boundingRectERK7QString(void* qthis, void* arg0);
  // proto:  int QFontMetrics::xHeight();
extern void _ZNK12QFontMetrics7xHeightEv(void* qthis);
  // proto:  int QFontMetrics::width(const QString & , int len, int flags);
extern void _ZNK12QFontMetrics5widthERK7QStringii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  int QFontMetrics::strikeOutPos();
extern void _ZNK12QFontMetrics12strikeOutPosEv(void* qthis);
  // proto:  int QFontMetrics::lineSpacing();
extern void _ZNK12QFontMetrics11lineSpacingEv(void* qthis);
  // proto:  void QFontMetrics::QFontMetrics(const QFontMetrics & );
extern void* dector_ZN12QFontMetricsC1ERKS_(void* arg0);
extern void _ZN12QFontMetricsC1ERKS_(void* qthis, void* arg0);
  // proto:  void QFontMetrics::QFontMetrics(const QFont & , QPaintDevice * pd);
extern void* dector_ZN12QFontMetricsC1ERK5QFontP12QPaintDevice(void* arg0, void* arg1);
extern void _ZN12QFontMetricsC1ERK5QFontP12QPaintDevice(void* qthis, void* arg0, void* arg1);
  // proto:  int QFontMetrics::minRightBearing();
extern void _ZNK12QFontMetrics15minRightBearingEv(void* qthis);
  // proto:  void QFontMetrics::swap(QFontMetrics & other);
extern void demth_ZN12QFontMetrics4swapERS_(void* qthis, void* arg0);
  // proto:  QRect QFontMetrics::boundingRect(QChar );
extern void _ZNK12QFontMetrics12boundingRectE5QChar(void* qthis, void* arg0);
  // proto:  void QFontMetrics::QFontMetrics(const QFont & );
extern void* dector_ZN12QFontMetricsC1ERK5QFont(void* arg0);
extern void _ZN12QFontMetricsC1ERK5QFont(void* qthis, void* arg0);
  // proto:  int QFontMetrics::width(const QString & , int len);
extern void _ZNK12QFontMetrics5widthERK7QStringi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QRect QFontMetrics::boundingRect(int x, int y, int w, int h, int flags, const QString & text, int tabstops, int * tabarray);
extern void demth_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4, void* arg5, int32_t arg6, int32_t* arg7);
  // proto:  int QFontMetrics::charWidth(const QString & str, int pos);
extern void _ZNK12QFontMetrics9charWidthERK7QStringi(void* qthis, void* arg0, int32_t arg1);
  // proto:  int QFontMetrics::leftBearing(QChar );
extern void _ZNK12QFontMetrics11leftBearingE5QChar(void* qthis, void* arg0);
  // proto:  bool QFontMetrics::inFontUcs4(uint ucs4);
extern void _ZNK12QFontMetrics10inFontUcs4Ej(void* qthis, int32_t arg0);
  // proto:  int QFontMetrics::descent();
extern void _ZNK12QFontMetrics7descentEv(void* qthis);
  // proto:  bool QFontMetricsF::inFont(QChar );
extern void _ZNK13QFontMetricsF6inFontE5QChar(void* qthis, void* arg0);
  // proto:  QSizeF QFontMetricsF::size(int flags, const QString & str, int tabstops, int * tabarray);
extern void _ZNK13QFontMetricsF4sizeEiRK7QStringiPi(void* qthis, int32_t arg0, void* arg1, int32_t arg2, int32_t* arg3);
  // proto:  qreal QFontMetricsF::minRightBearing();
extern void _ZNK13QFontMetricsF15minRightBearingEv(void* qthis);
  // proto:  void QFontMetricsF::QFontMetricsF(const QFontMetricsF & );
extern void* dector_ZN13QFontMetricsFC1ERKS_(void* arg0);
extern void _ZN13QFontMetricsFC1ERKS_(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::xHeight();
extern void _ZNK13QFontMetricsF7xHeightEv(void* qthis);
  // proto:  qreal QFontMetricsF::width(QChar );
extern void _ZNK13QFontMetricsF5widthE5QChar(void* qthis, void* arg0);
  // proto:  void QFontMetricsF::~QFontMetricsF();
extern void _ZN13QFontMetricsFD0Ev(void* qthis);
  // proto:  QRectF QFontMetricsF::boundingRect(const QRectF & r, int flags, const QString & string, int tabstops, int * tabarray);
extern void _ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3, int32_t* arg4);
  // proto:  void QFontMetricsF::swap(QFontMetricsF & other);
extern void demth_ZN13QFontMetricsF4swapERS_(void* qthis, void* arg0);
  // proto:  QRectF QFontMetricsF::tightBoundingRect(const QString & text);
extern void _ZNK13QFontMetricsF17tightBoundingRectERK7QString(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::leftBearing(QChar );
extern void _ZNK13QFontMetricsF11leftBearingE5QChar(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::rightBearing(QChar );
extern void _ZNK13QFontMetricsF12rightBearingE5QChar(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::overlinePos();
extern void _ZNK13QFontMetricsF11overlinePosEv(void* qthis);
  // proto:  qreal QFontMetricsF::height();
extern void _ZNK13QFontMetricsF6heightEv(void* qthis);
  // proto:  qreal QFontMetricsF::descent();
extern void _ZNK13QFontMetricsF7descentEv(void* qthis);
  // proto:  QRectF QFontMetricsF::boundingRect(const QString & string);
extern void _ZNK13QFontMetricsF12boundingRectERK7QString(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::lineWidth();
extern void _ZNK13QFontMetricsF9lineWidthEv(void* qthis);
  // proto:  void QFontMetricsF::QFontMetricsF(const QFontMetrics & );
extern void* dector_ZN13QFontMetricsFC1ERK12QFontMetrics(void* arg0);
extern void _ZN13QFontMetricsFC1ERK12QFontMetrics(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::width(const QString & string);
extern void _ZNK13QFontMetricsF5widthERK7QString(void* qthis, void* arg0);
  // proto:  qreal QFontMetricsF::strikeOutPos();
extern void _ZNK13QFontMetricsF12strikeOutPosEv(void* qthis);
  // proto:  qreal QFontMetricsF::lineSpacing();
extern void _ZNK13QFontMetricsF11lineSpacingEv(void* qthis);
  // proto:  qreal QFontMetricsF::averageCharWidth();
extern void _ZNK13QFontMetricsF16averageCharWidthEv(void* qthis);
  // proto:  void QFontMetricsF::QFontMetricsF(const QFont & , QPaintDevice * pd);
extern void* dector_ZN13QFontMetricsFC1ERK5QFontP12QPaintDevice(void* arg0, void* arg1);
extern void _ZN13QFontMetricsFC1ERK5QFontP12QPaintDevice(void* qthis, void* arg0, void* arg1);
  // proto:  qreal QFontMetricsF::leading();
extern void _ZNK13QFontMetricsF7leadingEv(void* qthis);
  // proto:  void QFontMetricsF::QFontMetricsF(const QFont & );
extern void* dector_ZN13QFontMetricsFC1ERK5QFont(void* arg0);
extern void _ZN13QFontMetricsFC1ERK5QFont(void* qthis, void* arg0);
  // proto:  QRectF QFontMetricsF::boundingRect(QChar );
extern void _ZNK13QFontMetricsF12boundingRectE5QChar(void* qthis, void* arg0);
  // proto:  bool QFontMetricsF::inFontUcs4(uint ucs4);
extern void _ZNK13QFontMetricsF10inFontUcs4Ej(void* qthis, int32_t arg0);
  // proto:  qreal QFontMetricsF::minLeftBearing();
extern void _ZNK13QFontMetricsF14minLeftBearingEv(void* qthis);
  // proto:  qreal QFontMetricsF::ascent();
extern void _ZNK13QFontMetricsF6ascentEv(void* qthis);
  // proto:  qreal QFontMetricsF::maxWidth();
extern void _ZNK13QFontMetricsF8maxWidthEv(void* qthis);
  // proto:  qreal QFontMetricsF::underlinePos();
extern void _ZNK13QFontMetricsF12underlinePosEv(void* qthis);
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QFontMetricsF)=1
type QFontMetricsF struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  int QFontMetrics::maxWidth();
func (this *QFontMetrics) maxWidth(args ...interface{}) () {
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
    C._ZNK12QFontMetrics8maxWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "maxWidth", args)
  }

}

  // proto:  void QFontMetrics::~QFontMetrics();
func (this *QFontMetrics) FreeQFontMetrics(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontMetrics", "~QFontMetrics", args)
  }

}

  // proto:  int QFontMetrics::lineWidth();
func (this *QFontMetrics) lineWidth(args ...interface{}) () {
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
    C._ZNK12QFontMetrics9lineWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "lineWidth", args)
  }

}

  // proto:  QRect QFontMetrics::boundingRect(const QRect & r, int flags, const QString & text, int tabstops, int * tabarray);
func (this *QFontMetrics) boundingRect(args ...interface{}) () {
  // boundingRect(const class QRect &, int, const class QString &, int, int *)
  // boundingRect(const class QString &)
  // boundingRect(class QChar)
  // boundingRect(int, int, int, int, int, const class QString &, int, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[3][4] = qtrt.Int32Ty(false) // "int"
  vtys[3][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][6] = qtrt.Int32Ty(false) // "int"
  vtys[3][7] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi
    // invoke: QRect boundingRect(const class QRect &, int, const class QString &, int, int *)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.int32_t)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    C._ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZNK12QFontMetrics12boundingRectERK7QString
    // invoke: QRect boundingRect(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics12boundingRectERK7QString(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK12QFontMetrics12boundingRectE5QChar
    // invoke: QRect boundingRect(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics12boundingRectE5QChar(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi
    // invoke: QRect boundingRect(int, int, int, int, int, const class QString &, int, int *)
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
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = C.int32_t(args[6].(int32))
    if false {fmt.Println(arg6)}
    var arg7 = (*C.int32_t)(args[7].(*int32))
    if false {fmt.Println(arg7)}
    C.demth_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
  default:
    qtrt.ErrorResolve("QFontMetrics", "boundingRect", args)
  }

}

  // proto:  int QFontMetrics::minLeftBearing();
func (this *QFontMetrics) minLeftBearing(args ...interface{}) () {
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
    C._ZNK12QFontMetrics14minLeftBearingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "minLeftBearing", args)
  }

}

  // proto:  int QFontMetrics::rightBearing(QChar );
func (this *QFontMetrics) rightBearing(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics12rightBearingE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "rightBearing", args)
  }

}

  // proto:  int QFontMetrics::ascent();
func (this *QFontMetrics) ascent(args ...interface{}) () {
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
    C._ZNK12QFontMetrics6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "ascent", args)
  }

}

  // proto:  QSize QFontMetrics::size(int flags, const QString & str, int tabstops, int * tabarray);
func (this *QFontMetrics) size(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK12QFontMetrics4sizeEiRK7QStringiPi(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QFontMetrics", "size", args)
  }

}

  // proto:  int QFontMetrics::overlinePos();
func (this *QFontMetrics) overlinePos(args ...interface{}) () {
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
    C._ZNK12QFontMetrics11overlinePosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "overlinePos", args)
  }

}

  // proto:  int QFontMetrics::leading();
func (this *QFontMetrics) leading(args ...interface{}) () {
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
    C._ZNK12QFontMetrics7leadingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "leading", args)
  }

}

  // proto:  QRect QFontMetrics::tightBoundingRect(const QString & text);
func (this *QFontMetrics) tightBoundingRect(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics17tightBoundingRectERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "tightBoundingRect", args)
  }

}

  // proto:  int QFontMetrics::averageCharWidth();
func (this *QFontMetrics) averageCharWidth(args ...interface{}) () {
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
    C._ZNK12QFontMetrics16averageCharWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "averageCharWidth", args)
  }

}

  // proto:  int QFontMetrics::underlinePos();
func (this *QFontMetrics) underlinePos(args ...interface{}) () {
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
    C._ZNK12QFontMetrics12underlinePosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "underlinePos", args)
  }

}

  // proto:  bool QFontMetrics::inFont(QChar );
func (this *QFontMetrics) inFont(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics6inFontE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "inFont", args)
  }

}

  // proto:  int QFontMetrics::height();
func (this *QFontMetrics) height(args ...interface{}) () {
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
    C._ZNK12QFontMetrics6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "height", args)
  }

}

  // proto:  int QFontMetrics::width(QChar );
func (this *QFontMetrics) width(args ...interface{}) () {
  // width(class QChar)
  // width(const class QString &, int, int)
  // width(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QFontMetrics5widthE5QChar
    // invoke: int width(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics5widthE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringii
    // invoke: int width(const class QString &, int, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZNK12QFontMetrics5widthERK7QStringii(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringi
    // invoke: int width(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK12QFontMetrics5widthERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontMetrics", "width", args)
  }

}

  // proto:  int QFontMetrics::xHeight();
func (this *QFontMetrics) xHeight(args ...interface{}) () {
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
    C._ZNK12QFontMetrics7xHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "xHeight", args)
  }

}

  // proto:  int QFontMetrics::strikeOutPos();
func (this *QFontMetrics) strikeOutPos(args ...interface{}) () {
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
    C._ZNK12QFontMetrics12strikeOutPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "strikeOutPos", args)
  }

}

  // proto:  int QFontMetrics::lineSpacing();
func (this *QFontMetrics) lineSpacing(args ...interface{}) () {
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
    C._ZNK12QFontMetrics11lineSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "lineSpacing", args)
  }

}

  // proto:  void QFontMetrics::QFontMetrics(const QFontMetrics & );
func NewQFontMetrics(args ...interface{}) QFontMetrics {
  return QFontMetrics{}
}

  // proto:  int QFontMetrics::minRightBearing();
func (this *QFontMetrics) minRightBearing(args ...interface{}) () {
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
    C._ZNK12QFontMetrics15minRightBearingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "minRightBearing", args)
  }

}

  // proto:  void QFontMetrics::swap(QFontMetrics & other);
func (this *QFontMetrics) swap(args ...interface{}) () {
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
    var arg0 = args[0].(QFontMetrics).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN12QFontMetrics4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "swap", args)
  }

}

  // proto:  int QFontMetrics::charWidth(const QString & str, int pos);
func (this *QFontMetrics) charWidth(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK12QFontMetrics9charWidthERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QFontMetrics", "charWidth", args)
  }

}

  // proto:  int QFontMetrics::leftBearing(QChar );
func (this *QFontMetrics) leftBearing(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics11leftBearingE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "leftBearing", args)
  }

}

  // proto:  bool QFontMetrics::inFontUcs4(uint ucs4);
func (this *QFontMetrics) inFontUcs4(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK12QFontMetrics10inFontUcs4Ej(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetrics", "inFontUcs4", args)
  }

}

  // proto:  int QFontMetrics::descent();
func (this *QFontMetrics) descent(args ...interface{}) () {
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
    C._ZNK12QFontMetrics7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetrics", "descent", args)
  }

}

  // proto:  bool QFontMetricsF::inFont(QChar );
func (this *QFontMetricsF) inFont(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF6inFontE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "inFont", args)
  }

}

  // proto:  QSizeF QFontMetricsF::size(int flags, const QString & str, int tabstops, int * tabarray);
func (this *QFontMetricsF) size(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C._ZNK13QFontMetricsF4sizeEiRK7QStringiPi(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "size", args)
  }

}

  // proto:  qreal QFontMetricsF::minRightBearing();
func (this *QFontMetricsF) minRightBearing(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF15minRightBearingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "minRightBearing", args)
  }

}

  // proto:  void QFontMetricsF::QFontMetricsF(const QFontMetricsF & );
func NewQFontMetricsF(args ...interface{}) QFontMetricsF {
  return QFontMetricsF{}
}

  // proto:  qreal QFontMetricsF::xHeight();
func (this *QFontMetricsF) xHeight(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF7xHeightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "xHeight", args)
  }

}

  // proto:  qreal QFontMetricsF::width(QChar );
func (this *QFontMetricsF) width(args ...interface{}) () {
  // width(class QChar)
  // width(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF5widthE5QChar
    // invoke: qreal width(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF5widthE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK13QFontMetricsF5widthERK7QString
    // invoke: qreal width(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF5widthERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "width", args)
  }

}

  // proto:  void QFontMetricsF::~QFontMetricsF();
func (this *QFontMetricsF) FreeQFontMetricsF(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFontMetricsF", "~QFontMetricsF", args)
  }

}

  // proto:  QRectF QFontMetricsF::boundingRect(const QRectF & r, int flags, const QString & string, int tabstops, int * tabarray);
func (this *QFontMetricsF) boundingRect(args ...interface{}) () {
  // boundingRect(const class QRectF &, int, const class QString &, int, int *)
  // boundingRect(const class QString &)
  // boundingRect(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi
    // invoke: QRectF boundingRect(const class QRectF &, int, const class QString &, int, int *)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.int32_t)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    C._ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 1:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK7QString
    // invoke: QRectF boundingRect(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF12boundingRectERK7QString(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK13QFontMetricsF12boundingRectE5QChar
    // invoke: QRectF boundingRect(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF12boundingRectE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "boundingRect", args)
  }

}

  // proto:  void QFontMetricsF::swap(QFontMetricsF & other);
func (this *QFontMetricsF) swap(args ...interface{}) () {
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
    var arg0 = args[0].(QFontMetricsF).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QFontMetricsF4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "swap", args)
  }

}

  // proto:  QRectF QFontMetricsF::tightBoundingRect(const QString & text);
func (this *QFontMetricsF) tightBoundingRect(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF17tightBoundingRectERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "tightBoundingRect", args)
  }

}

  // proto:  qreal QFontMetricsF::leftBearing(QChar );
func (this *QFontMetricsF) leftBearing(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF11leftBearingE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "leftBearing", args)
  }

}

  // proto:  qreal QFontMetricsF::rightBearing(QChar );
func (this *QFontMetricsF) rightBearing(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF12rightBearingE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "rightBearing", args)
  }

}

  // proto:  qreal QFontMetricsF::overlinePos();
func (this *QFontMetricsF) overlinePos(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF11overlinePosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "overlinePos", args)
  }

}

  // proto:  qreal QFontMetricsF::height();
func (this *QFontMetricsF) height(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF6heightEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "height", args)
  }

}

  // proto:  qreal QFontMetricsF::descent();
func (this *QFontMetricsF) descent(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF7descentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "descent", args)
  }

}

  // proto:  qreal QFontMetricsF::lineWidth();
func (this *QFontMetricsF) lineWidth(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF9lineWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "lineWidth", args)
  }

}

  // proto:  qreal QFontMetricsF::strikeOutPos();
func (this *QFontMetricsF) strikeOutPos(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF12strikeOutPosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "strikeOutPos", args)
  }

}

  // proto:  qreal QFontMetricsF::lineSpacing();
func (this *QFontMetricsF) lineSpacing(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF11lineSpacingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "lineSpacing", args)
  }

}

  // proto:  qreal QFontMetricsF::averageCharWidth();
func (this *QFontMetricsF) averageCharWidth(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF16averageCharWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "averageCharWidth", args)
  }

}

  // proto:  qreal QFontMetricsF::leading();
func (this *QFontMetricsF) leading(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF7leadingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "leading", args)
  }

}

  // proto:  bool QFontMetricsF::inFontUcs4(uint ucs4);
func (this *QFontMetricsF) inFontUcs4(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QFontMetricsF10inFontUcs4Ej(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "inFontUcs4", args)
  }

}

  // proto:  qreal QFontMetricsF::minLeftBearing();
func (this *QFontMetricsF) minLeftBearing(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF14minLeftBearingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "minLeftBearing", args)
  }

}

  // proto:  qreal QFontMetricsF::ascent();
func (this *QFontMetricsF) ascent(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF6ascentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "ascent", args)
  }

}

  // proto:  qreal QFontMetricsF::maxWidth();
func (this *QFontMetricsF) maxWidth(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF8maxWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "maxWidth", args)
  }

}

  // proto:  qreal QFontMetricsF::underlinePos();
func (this *QFontMetricsF) underlinePos(args ...interface{}) () {
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
    C._ZNK13QFontMetricsF12underlinePosEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFontMetricsF", "underlinePos", args)
  }

}

// <= body block end

