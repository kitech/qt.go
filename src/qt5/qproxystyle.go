package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qproxystyle.h
// dst-file: /src/widgets/qproxystyle.go
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
  // proto:  void QProxyStyle::unpolish(QWidget * widget);
extern void _ZN11QProxyStyle8unpolishEP7QWidget(void* qthis, void* arg0);
  // proto:  void QProxyStyle::QProxyStyle(const QString & key);
extern void* dector_ZN11QProxyStyleC1ERK7QString(void* arg0);
extern void _ZN11QProxyStyleC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QProxyStyle::unpolish(QApplication * app);
extern void _ZN11QProxyStyle8unpolishEP12QApplication(void* qthis, void* arg0);
  // proto:  QPalette QProxyStyle::standardPalette();
extern void _ZNK11QProxyStyle15standardPaletteEv(void* qthis);
  // proto:  void QProxyStyle::setBaseStyle(QStyle * style);
extern void _ZN11QProxyStyle12setBaseStyleEP6QStyle(void* qthis, void* arg0);
  // proto:  void QProxyStyle::polish(QPalette & pal);
extern void _ZN11QProxyStyle6polishER8QPalette(void* qthis, void* arg0);
  // proto:  void QProxyStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
extern void _ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(void* qthis, void* arg0, void* arg1, int32_t arg2, void* arg3);
  // proto:  void QProxyStyle::~QProxyStyle();
extern void _ZN11QProxyStyleD0Ev(void* qthis);
  // proto:  QStyle * QProxyStyle::baseStyle();
extern void _ZNK11QProxyStyle9baseStyleEv(void* qthis);
  // proto:  void QProxyStyle::QProxyStyle(const QProxyStyle & );
extern void* dector_ZN11QProxyStyleC1ERKS_(void* arg0);
extern void _ZN11QProxyStyleC1ERKS_(void* qthis, void* arg0);
  // proto:  void QProxyStyle::polish(QApplication * app);
extern void _ZN11QProxyStyle6polishEP12QApplication(void* qthis, void* arg0);
  // proto:  void QProxyStyle::polish(QWidget * widget);
extern void _ZN11QProxyStyle6polishEP7QWidget(void* qthis, void* arg0);
  // proto:  QRect QProxyStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
extern void _ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(void* qthis, void* arg0, void* arg1, int32_t arg2, bool arg3, void* arg4);
  // proto:  QRect QProxyStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
extern void _ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  const QMetaObject * QProxyStyle::metaObject();
extern void _ZNK11QProxyStyle10metaObjectEv(void* qthis);
  // proto:  void QProxyStyle::QProxyStyle(QStyle * style);
extern void* dector_ZN11QProxyStyleC1EP6QStyle(void* arg0);
extern void _ZN11QProxyStyleC1EP6QStyle(void* qthis, void* arg0);
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

// class sizeof(QProxyStyle)=1
type QProxyStyle struct {
  /*qbase*/ QCommonStyle;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QProxyStyle::unpolish(QWidget * widget);
func (this *QProxyStyle) unpolish(args ...interface{}) () {
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
    // invoke: _ZN11QProxyStyle8unpolishEP7QWidget
    // invoke: void unpolish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QProxyStyle8unpolishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QProxyStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QProxyStyle8unpolishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "unpolish", args)
  }

}

  // proto:  void QProxyStyle::QProxyStyle(const QString & key);
func NewQProxyStyle(args ...interface{}) QProxyStyle {
  return QProxyStyle{}
}

  // proto:  QPalette QProxyStyle::standardPalette();
func (this *QProxyStyle) standardPalette(args ...interface{}) () {
  // standardPalette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QProxyStyle15standardPaletteEv
    // invoke: QPalette standardPalette()
    C._ZNK11QProxyStyle15standardPaletteEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProxyStyle", "standardPalette", args)
  }

}

  // proto:  void QProxyStyle::setBaseStyle(QStyle * style);
func (this *QProxyStyle) setBaseStyle(args ...interface{}) () {
  // setBaseStyle(class QStyle *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QProxyStyle12setBaseStyleEP6QStyle
    // invoke: void setBaseStyle(class QStyle *)
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QProxyStyle12setBaseStyleEP6QStyle(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "setBaseStyle", args)
  }

}

  // proto:  void QProxyStyle::polish(QPalette & pal);
func (this *QProxyStyle) polish(args ...interface{}) () {
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
    // invoke: _ZN11QProxyStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QProxyStyle6polishER8QPalette(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QProxyStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QProxyStyle6polishEP12QApplication(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN11QProxyStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN11QProxyStyle6polishEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "polish", args)
  }

}

  // proto:  void QProxyStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
func (this *QProxyStyle) drawItemPixmap(args ...interface{}) () {
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
    // invoke: _ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap
    // invoke: void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRect).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QPixmap).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QProxyStyle", "drawItemPixmap", args)
  }

}

  // proto:  void QProxyStyle::~QProxyStyle();
func (this *QProxyStyle) FreeQProxyStyle(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProxyStyle", "~QProxyStyle", args)
  }

}

  // proto:  QStyle * QProxyStyle::baseStyle();
func (this *QProxyStyle) baseStyle(args ...interface{}) () {
  // baseStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QProxyStyle9baseStyleEv
    // invoke: QStyle * baseStyle()
    C._ZNK11QProxyStyle9baseStyleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProxyStyle", "baseStyle", args)
  }

}

  // proto:  QRect QProxyStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
func (this *QProxyStyle) itemTextRect(args ...interface{}) () {
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
    // invoke: _ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString
    // invoke: QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
    var arg0 = args[0].(QFontMetrics).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRect).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.bool(args[3].(bool))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    C._ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QProxyStyle", "itemTextRect", args)
  }

}

  // proto:  QRect QProxyStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
func (this *QProxyStyle) itemPixmapRect(args ...interface{}) () {
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
    // invoke: _ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap
    // invoke: QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QPixmap).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QProxyStyle", "itemPixmapRect", args)
  }

}

  // proto:  const QMetaObject * QProxyStyle::metaObject();
func (this *QProxyStyle) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QProxyStyle10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK11QProxyStyle10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProxyStyle", "metaObject", args)
  }

}

// <= body block end

