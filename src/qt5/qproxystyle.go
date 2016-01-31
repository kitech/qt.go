package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QProxyStyle::QProxyStyle(QStyle * style);
extern void* C_ZN11QProxyStyleC2EP6QStyle(void* arg0); // 3
  // proto:  void QProxyStyle::QProxyStyle(const QString & key);
extern void* C_ZN11QProxyStyleC2ERK7QString(void* arg0); // 3
  // proto:  void QProxyStyle::polish(QWidget * widget);
extern void C_ZN11QProxyStyle6polishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QProxyStyle::polish(QPalette & pal);
extern void C_ZN11QProxyStyle6polishER8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QProxyStyle::polish(QApplication * app);
extern void C_ZN11QProxyStyle6polishEP12QApplication(void* qthis, void* arg0); // 4
  // proto:  QStyle * QProxyStyle::baseStyle();
extern void C_ZNK11QProxyStyle9baseStyleEv(void* qthis); // 4
  // proto:  QPalette QProxyStyle::standardPalette();
extern void C_ZNK11QProxyStyle15standardPaletteEv(void* qthis); // 4
  // proto:  void QProxyStyle::setBaseStyle(QStyle * style);
extern void C_ZN11QProxyStyle12setBaseStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto:  void QProxyStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
extern void C_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(void* qthis, void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto:  const QMetaObject * QProxyStyle::metaObject();
extern void C_ZNK11QProxyStyle10metaObjectEv(void* qthis); // 4
  // proto:  void QProxyStyle::~QProxyStyle();
extern void C_ZN11QProxyStyleD2Ev(void* qthis); // 4
  // proto:  QRect QProxyStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
extern void C_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(void* qthis, void* arg0, void* arg1, int32_t arg2, bool arg3, void* arg4); // 4
  // proto:  QRect QProxyStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
extern void C_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QProxyStyle::unpolish(QWidget * widget);
extern void C_ZN11QProxyStyle8unpolishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QProxyStyle::unpolish(QApplication * app);
extern void C_ZN11QProxyStyle8unpolishEP12QApplication(void* qthis, void* arg0); // 4
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

// QProxyStyle(class QStyle *)
func NewQProxyStyle(args ...interface{}) *QProxyStyle {
  // QProxyStyle(class QStyle *)
  // QProxyStyle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QProxyStyleC1EP6QStyle
    // invoke: void QProxyStyle(class QStyle *)
    var arg0 = args[0].(QStyle).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QProxyStyleC2EP6QStyle(arg0)
    return &QProxyStyle{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QProxyStyleC1ERK7QString
    // invoke: void QProxyStyle(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QProxyStyleC2ERK7QString(arg0)
    return &QProxyStyle{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QProxyStyle", "QProxyStyle", args)
  }

  return nil // QProxyStyle{qclsinst:qthis}
}

// polish(class QWidget *)
func (this *QProxyStyle) polish(args ...interface{}) () {
  // polish(class QWidget *)
  // polish(class QPalette &)
  // polish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPalette{}) // "QPalette &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QProxyStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle6polishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QProxyStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(QPalette).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle6polishER8QPalette(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN11QProxyStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle6polishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "polish", args)
  }

}

// baseStyle()
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
    var ret = C.C_ZNK11QProxyStyle9baseStyleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QProxyStyle", "baseStyle", args)
  }

}

// standardPalette()
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
    var ret = C.C_ZNK11QProxyStyle15standardPaletteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QProxyStyle", "standardPalette", args)
  }

}

// setBaseStyle(class QStyle *)
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
    C.C_ZN11QProxyStyle12setBaseStyleEP6QStyle(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "setBaseStyle", args)
  }

}

// drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
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
    C.C_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QProxyStyle", "drawItemPixmap", args)
  }

}

// metaObject()
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
    C.C_ZNK11QProxyStyle10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProxyStyle", "metaObject", args)
  }

}

// ~QProxyStyle()
func (this *QProxyStyle) FreeQProxyStyle(args ...interface{}) () {
  // ~QProxyStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QProxyStyleD0Ev
    // invoke: void ~QProxyStyle()
    C.C_ZN11QProxyStyleD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProxyStyle", "~QProxyStyle", args)
  }

}

// itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
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
    var ret = C.C_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QProxyStyle", "itemTextRect", args)
  }

}

// itemPixmapRect(const class QRect &, int, const class QPixmap &)
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
    var ret = C.C_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QProxyStyle", "itemPixmapRect", args)
  }

}

// unpolish(class QWidget *)
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
    C.C_ZN11QProxyStyle8unpolishEP7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QProxyStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(QApplication).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle8unpolishEP12QApplication(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "unpolish", args)
  }

}

// <= body block end

