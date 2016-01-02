package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
extern void _ZNK12QFontMetrics12boundingRectERK5QRectiRK7QStringiPi(void* qthis, void* arg0, int arg1, void* arg2, int arg3, int* arg4);
  // proto:  int QFontMetrics::minLeftBearing();
extern void _ZNK12QFontMetrics14minLeftBearingEv(void* qthis);
  // proto:  int QFontMetrics::rightBearing(QChar );
extern void _ZNK12QFontMetrics12rightBearingE5QChar(void* qthis, void* arg0);
  // proto:  int QFontMetrics::ascent();
extern void _ZNK12QFontMetrics6ascentEv(void* qthis);
  // proto:  QSize QFontMetrics::size(int flags, const QString & str, int tabstops, int * tabarray);
extern void _ZNK12QFontMetrics4sizeEiRK7QStringiPi(void* qthis, int arg0, void* arg1, int arg2, int* arg3);
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
extern void _ZNK12QFontMetrics5widthERK7QStringii(void* qthis, void* arg0, int arg1, int arg2);
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
extern void _ZN12QFontMetrics4swapERS_(void* qthis, void* arg0);
  // proto:  QRect QFontMetrics::boundingRect(QChar );
extern void _ZNK12QFontMetrics12boundingRectE5QChar(void* qthis, void* arg0);
  // proto:  void QFontMetrics::QFontMetrics(const QFont & );
extern void* dector_ZN12QFontMetricsC1ERK5QFont(void* arg0);
extern void _ZN12QFontMetricsC1ERK5QFont(void* qthis, void* arg0);
  // proto:  int QFontMetrics::width(const QString & , int len);
extern void _ZNK12QFontMetrics5widthERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  QRect QFontMetrics::boundingRect(int x, int y, int w, int h, int flags, const QString & text, int tabstops, int * tabarray);
extern void demth_ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi(void* qthis, int arg0, int arg1, int arg2, int arg3, int arg4, void* arg5, int arg6, int* arg7);
  // proto:  int QFontMetrics::charWidth(const QString & str, int pos);
extern void _ZNK12QFontMetrics9charWidthERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  int QFontMetrics::leftBearing(QChar );
extern void _ZNK12QFontMetrics11leftBearingE5QChar(void* qthis, void* arg0);
  // proto:  bool QFontMetrics::inFontUcs4(uint ucs4);
extern void _ZNK12QFontMetrics10inFontUcs4Ej(void* qthis, unsigned int arg0);
  // proto:  int QFontMetrics::descent();
extern void _ZNK12QFontMetrics7descentEv(void* qthis);
  // proto:  bool QFontMetricsF::inFont(QChar );
extern void _ZNK13QFontMetricsF6inFontE5QChar(void* qthis, void* arg0);
  // proto:  QSizeF QFontMetricsF::size(int flags, const QString & str, int tabstops, int * tabarray);
extern void _ZNK13QFontMetricsF4sizeEiRK7QStringiPi(void* qthis, int arg0, void* arg1, int arg2, int* arg3);
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
extern void _ZNK13QFontMetricsF12boundingRectERK6QRectFiRK7QStringiPi(void* qthis, void* arg0, int arg1, void* arg2, int arg3, int* arg4);
  // proto:  void QFontMetricsF::swap(QFontMetricsF & other);
extern void _ZN13QFontMetricsF4swapERS_(void* qthis, void* arg0);
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
extern void _ZNK13QFontMetricsF10inFontUcs4Ej(void* qthis, unsigned int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QFontMetricsF)=1
type QFontMetricsF struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
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
  case 1:
    // invoke: _ZNK12QFontMetrics12boundingRectERK7QString
  case 2:
    // invoke: _ZNK12QFontMetrics12boundingRectE5QChar
  case 3:
    // invoke: _ZNK12QFontMetrics12boundingRectEiiiiiRK7QStringiPi
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
  case 1:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringii
  case 2:
    // invoke: _ZNK12QFontMetrics5widthERK7QStringi
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
  case 1:
    // invoke: _ZNK13QFontMetricsF5widthERK7QString
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
  case 1:
    // invoke: _ZNK13QFontMetricsF12boundingRectERK7QString
  case 2:
    // invoke: _ZNK13QFontMetricsF12boundingRectE5QChar
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
  default:
    qtrt.ErrorResolve("QFontMetricsF", "underlinePos", args)
  }

}

// <= body block end

