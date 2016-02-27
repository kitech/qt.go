package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
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
extern void* C_ZNK11QProxyStyle9baseStyleEv(void* qthis); // 4
  // proto:  QPalette QProxyStyle::standardPalette();
extern void* C_ZNK11QProxyStyle15standardPaletteEv(void* qthis); // 4
  // proto:  void QProxyStyle::setBaseStyle(QStyle * style);
extern void C_ZN11QProxyStyle12setBaseStyleEP6QStyle(void* qthis, void* arg0); // 4
  // proto:  void QProxyStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
extern void C_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(void* qthis, void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto:  const QMetaObject * QProxyStyle::metaObject();
extern void C_ZNK11QProxyStyle10metaObjectEv(void* qthis); // 4
  // proto:  void QProxyStyle::~QProxyStyle();
extern void C_ZN11QProxyStyleD2Ev(void* qthis); // 4
  // proto:  QRect QProxyStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
extern void* C_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(void* qthis, void* arg0, void* arg1, int32_t arg2, bool arg3, void* arg4); // 4
  // proto:  QRect QProxyStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
extern void* C_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
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
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QProxyStyle)=1
type QProxyStyle struct {
  /*qbase*/ QCommonStyle;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QProxyStyle(class QStyle *)
func GcfreeQProxyStyle(this *QProxyStyle) {
  qtrt.UniverseFree(this)
}
func NewQProxyStyle(args ...interface{}) *QProxyStyle {
  // QProxyStyle(class QStyle *)
  // QProxyStyle(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStyle{}) // "QStyle *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QProxyStyleC1EP6QStyle
    // invoke: void QProxyStyle(class QStyle *)
    var arg0 = args[0].(*QStyle).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QProxyStyleC2EP6QStyle(arg0)
    this := &QProxyStyle{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQProxyStyle)
    return this
  case 1:
    // invoke: _ZN11QProxyStyleC1ERK7QString
    // invoke: void QProxyStyle(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QProxyStyleC2ERK7QString(arg0)
    this := &QProxyStyle{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQProxyStyle)
    return this
  default:
    qtrt.ErrorResolve("QProxyStyle", "QProxyStyle", args)
  }

  return nil // QProxyStyle{Qclsinst:qthis}
}

// polish(class QWidget *)
func (this *QProxyStyle) Polish(args ...interface{}) () {
  // polish(class QWidget *)
  // polish(class QPalette &)
  // polish(class QApplication *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtgui.QPalette{}) // "QPalette &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QApplication{}) // "QApplication *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QProxyStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle6polishEP7QWidget(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QProxyStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(*qtgui.QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle6polishER8QPalette(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN11QProxyStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(*QApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle6polishEP12QApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "polish", args)
  }

  return
}

// baseStyle()
func (this *QProxyStyle) BaseStyle(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QProxyStyle9baseStyleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProxyStyle", "baseStyle", args)
  }

  return
}

// standardPalette()
func (this *QProxyStyle) StandardPalette(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QProxyStyle15standardPaletteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProxyStyle", "standardPalette", args)
  }

  return
}

// setBaseStyle(class QStyle *)
func (this *QProxyStyle) SetBaseStyle(args ...interface{}) () {
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
    var arg0 = args[0].(*QStyle).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle12setBaseStyleEP6QStyle(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "setBaseStyle", args)
  }

  return
}

// drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QProxyStyle) DrawItemPixmap(args ...interface{}) () {
  // drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPainter{}) // "QPainter *"
  vtys[0][1] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap
    // invoke: void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(*qtgui.QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZNK11QProxyStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QProxyStyle", "drawItemPixmap", args)
  }

  return
}

// metaObject()
func (this *QProxyStyle) MetaObject(args ...interface{}) () {
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
    C.C_ZNK11QProxyStyle10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProxyStyle", "metaObject", args)
  }

  return
}

// ~QProxyStyle()
func (this *QProxyStyle) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN11QProxyStyleD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QProxyStyle", "~QProxyStyle", args)
  }

  return
}

// itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QProxyStyle) ItemTextRect(args ...interface{}) (ret interface{}) {
  // itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFontMetrics{}) // "const QFontMetrics &"
  vtys[0][1] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.BoolTy(false) // "bool"
  vtys[0][4] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString
    // invoke: QRect itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
    var arg0 = args[0].(*qtgui.QFontMetrics).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.bool(args[3].(bool))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK11QProxyStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProxyStyle", "itemTextRect", args)
  }

  return
}

// itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QProxyStyle) ItemPixmapRect(args ...interface{}) (ret interface{}) {
  // itemPixmapRect(const class QRect &, int, const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QRect{}) // "const QRect &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap
    // invoke: QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK11QProxyStyle14itemPixmapRectERK5QRectiRK7QPixmap(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProxyStyle", "itemPixmapRect", args)
  }

  return
}

// unpolish(class QWidget *)
func (this *QProxyStyle) Unpolish(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle8unpolishEP7QWidget(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QProxyStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(*QApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QProxyStyle8unpolishEP12QApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProxyStyle", "unpolish", args)
  }

  return
}

// <= body block end

