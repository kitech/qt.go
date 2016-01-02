package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtWidgets/qstyle.h
// dst-file: /src/widgets/qstyle.go
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
  // proto:  void QStyle::QStyle(const QStyle & );
extern void* dector_ZN6QStyleC1ERKS_(void* arg0);
extern void _ZN6QStyleC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStyle::unpolish(QWidget * );
extern void _ZN6QStyle8unpolishEP7QWidget(void* qthis, void* arg0);
  // proto:  void QStyle::~QStyle();
extern void _ZN6QStyleD0Ev(void* qthis);
  // proto:  void QStyle::polish(QPalette & );
extern void _ZN6QStyle6polishER8QPalette(void* qthis, void* arg0);
  // proto:  void QStyle::QStyle();
extern void* dector_ZN6QStyleC1Ev();
extern void _ZN6QStyleC1Ev(void* qthis);
  // proto:  QRect QStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
extern void _ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  QRect QStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
extern void _ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(void* qthis, void* arg0, void* arg1, int arg2, bool arg3, void* arg4);
  // proto:  const QStyle * QStyle::proxy();
extern void _ZNK6QStyle5proxyEv(void* qthis);
  // proto:  QPalette QStyle::standardPalette();
extern void _ZNK6QStyle15standardPaletteEv(void* qthis);
  // proto:  const QMetaObject * QStyle::metaObject();
extern void _ZNK6QStyle10metaObjectEv(void* qthis);
  // proto:  void QStyle::polish(QApplication * );
extern void _ZN6QStyle6polishEP12QApplication(void* qthis, void* arg0);
  // proto:  void QStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
extern void _ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(void* qthis, void* arg0, void* arg1, int arg2, void* arg3);
  // proto:  void QStyle::polish(QWidget * );
extern void _ZN6QStyle6polishEP7QWidget(void* qthis, void* arg0);
  // proto: static int QStyle::sliderPositionFromValue(int min, int max, int val, int space, bool upsideDown);
extern void _ZN6QStyle23sliderPositionFromValueEiiiib(int arg0, int arg1, int arg2, int arg3, bool arg4);
  // proto: static int QStyle::sliderValueFromPosition(int min, int max, int pos, int space, bool upsideDown);
extern void _ZN6QStyle23sliderValueFromPositionEiiiib(int arg0, int arg1, int arg2, int arg3, bool arg4);
  // proto:  void QStyle::unpolish(QApplication * );
extern void _ZN6QStyle8unpolishEP12QApplication(void* qthis, void* arg0);
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

// class sizeof(QStyle)=1
type QStyle struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QStyle::QStyle(const QStyle & );
func NewQStyle(args ...interface{}) QStyle {
  return QStyle{}
}

  // proto:  void QStyle::unpolish(QWidget * );
func (this *QStyle) unpolish(args ...interface{}) () {
  // unpolish(class QWidget *)
  // unpolish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle8unpolishEP7QWidget
  case 1:
    // invoke: _ZN6QStyle8unpolishEP12QApplication
  default:
    qtrt.ErrorResolve("QStyle", "unpolish", args)
  }

}

  // proto:  void QStyle::~QStyle();
func (this *QStyle) FreeQStyle(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyle", "~QStyle", args)
  }

}

  // proto:  void QStyle::polish(QPalette & );
func (this *QStyle) polish(args ...interface{}) () {
  // polish(class QPalette &)
  // polish(class QApplication *)
  // polish(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPalette{}) // "QPalette &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QApplication{}) // "QApplication *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle6polishER8QPalette
  case 1:
    // invoke: _ZN6QStyle6polishEP12QApplication
  case 2:
    // invoke: _ZN6QStyle6polishEP7QWidget
  default:
    qtrt.ErrorResolve("QStyle", "polish", args)
  }

}

  // proto:  QRect QStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
func (this *QStyle) itemPixmapRect(args ...interface{}) () {
  // itemPixmapRect(const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap
  default:
    qtrt.ErrorResolve("QStyle", "itemPixmapRect", args)
  }

}

  // proto:  QRect QStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
func (this *QStyle) itemTextRect(args ...interface{}) () {
  // itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFontMetrics{}) // "const QFontMetrics &"
  vtys[0][1] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.BoolTy(false) // "bool"
  vtys[0][4] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString
  default:
    qtrt.ErrorResolve("QStyle", "itemTextRect", args)
  }

}

  // proto:  const QStyle * QStyle::proxy();
func (this *QStyle) proxy(args ...interface{}) () {
  // proxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle5proxyEv
  default:
    qtrt.ErrorResolve("QStyle", "proxy", args)
  }

}

  // proto:  QPalette QStyle::standardPalette();
func (this *QStyle) standardPalette(args ...interface{}) () {
  // standardPalette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle15standardPaletteEv
  default:
    qtrt.ErrorResolve("QStyle", "standardPalette", args)
  }

}

  // proto:  const QMetaObject * QStyle::metaObject();
func (this *QStyle) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle10metaObjectEv
  default:
    qtrt.ErrorResolve("QStyle", "metaObject", args)
  }

}

  // proto:  void QStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
func (this *QStyle) drawItemPixmap(args ...interface{}) () {
  // drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(QRect{}) // "const QRect &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap
  default:
    qtrt.ErrorResolve("QStyle", "drawItemPixmap", args)
  }

}

  // proto: static int QStyle::sliderPositionFromValue(int min, int max, int val, int space, bool upsideDown);
func (this *QStyle) sliderPositionFromValue_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyle", "sliderPositionFromValue", args)
  }

}

  // proto: static int QStyle::sliderValueFromPosition(int min, int max, int pos, int space, bool upsideDown);
func (this *QStyle) sliderValueFromPosition_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStyle", "sliderValueFromPosition", args)
  }

}

// <= body block end

