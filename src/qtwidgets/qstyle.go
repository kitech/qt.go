package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QStyle::drawItemPixmap(QPainter * painter, const QRect & rect, int alignment, const QPixmap & pixmap);
extern void C_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(void* qthis, void* arg0, void* arg1, int32_t arg2, void* arg3); // 4
  // proto:  void QStyle::~QStyle();
extern void C_ZN6QStyleD2Ev(void* qthis); // 4
  // proto:  void QStyle::polish(QWidget * );
extern void C_ZN6QStyle6polishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStyle::polish(QPalette & );
extern void C_ZN6QStyle6polishER8QPalette(void* qthis, void* arg0); // 4
  // proto:  void QStyle::polish(QApplication * );
extern void C_ZN6QStyle6polishEP12QApplication(void* qthis, void* arg0); // 4
  // proto:  QPalette QStyle::standardPalette();
extern void* C_ZNK6QStyle15standardPaletteEv(void* qthis); // 4
  // proto:  void QStyle::QStyle();
extern void* C_ZN6QStyleC2Ev(); // 3
  // proto: static int QStyle::sliderValueFromPosition(int min, int max, int pos, int space, bool upsideDown);
extern int32_t C_ZN6QStyle23sliderValueFromPositionEiiiib(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, bool arg4); // 4
  // proto:  const QStyle * QStyle::proxy();
extern void* C_ZNK6QStyle5proxyEv(void* qthis); // 4
  // proto:  const QMetaObject * QStyle::metaObject();
extern void C_ZNK6QStyle10metaObjectEv(void* qthis); // 4
  // proto:  QRect QStyle::itemTextRect(const QFontMetrics & fm, const QRect & r, int flags, bool enabled, const QString & text);
extern void* C_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(void* qthis, void* arg0, void* arg1, int32_t arg2, bool arg3, void* arg4); // 4
  // proto:  void QStyle::unpolish(QWidget * );
extern void C_ZN6QStyle8unpolishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStyle::unpolish(QApplication * );
extern void C_ZN6QStyle8unpolishEP12QApplication(void* qthis, void* arg0); // 4
  // proto: static int QStyle::sliderPositionFromValue(int min, int max, int val, int space, bool upsideDown);
extern int32_t C_ZN6QStyle23sliderPositionFromValueEiiiib(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, bool arg4); // 4
  // proto:  QRect QStyle::itemPixmapRect(const QRect & r, int flags, const QPixmap & pixmap);
extern void* C_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
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

// class sizeof(QStyle)=1
type QStyle struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
func (this *QStyle) DrawItemPixmap(args ...interface{}) () {
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
    // invoke: _ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap
    // invoke: void drawItemPixmap(class QPainter *, const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(*qtgui.QPainter).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg3)}
    C.C_ZNK6QStyle14drawItemPixmapEP8QPainterRK5QRectiRK7QPixmap(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QStyle", "drawItemPixmap", args)
  }

  return
}

// ~QStyle()
func (this *QStyle) Free(args ...interface{}) () {
  // ~QStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyleD0Ev
    // invoke: void ~QStyle()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN6QStyleD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QStyle", "~QStyle", args)
  }

  return
}

// polish(class QWidget *)
func (this *QStyle) Polish(args ...interface{}) () {
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
    // invoke: _ZN6QStyle6polishEP7QWidget
    // invoke: void polish(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle6polishEP7QWidget(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QStyle6polishER8QPalette
    // invoke: void polish(class QPalette &)
    var arg0 = args[0].(*qtgui.QPalette).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle6polishER8QPalette(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN6QStyle6polishEP12QApplication
    // invoke: void polish(class QApplication *)
    var arg0 = args[0].(*QApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle6polishEP12QApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyle", "polish", args)
  }

  return
}

// standardPalette()
func (this *QStyle) StandardPalette(args ...interface{}) (ret interface{}) {
  // standardPalette()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle15standardPaletteEv
    // invoke: QPalette standardPalette()
    var ret0 = C.C_ZNK6QStyle15standardPaletteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPalette{}) // "QPalette"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyle", "standardPalette", args)
  }

  return
}

// QStyle()
func GcfreeQStyle(this *QStyle) {
  qtrt.UniverseFree(this)
}
func NewQStyle(args ...interface{}) *QStyle {
  // QStyle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyleC1Ev
    // invoke: void QStyle()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QStyleC2Ev()
    this := &QStyle{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStyle)
    return this
  default:
    qtrt.ErrorResolve("QStyle", "QStyle", args)
  }

  return nil // QStyle{Qclsinst:qthis}
}

// sliderValueFromPosition(int, int, int, int, _Bool)
func (this *QStyle) SliderValueFromPosition_s(args ...interface{}) (ret interface{}) {
  // sliderValueFromPosition(int, int, int, int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle23sliderValueFromPositionEiiiib
    // invoke: int sliderValueFromPosition(int, int, int, int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.bool(args[4].(bool))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN6QStyle23sliderValueFromPositionEiiiib(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyle", "sliderValueFromPosition", args)
  }

  return
}

// proxy()
func (this *QStyle) Proxy(args ...interface{}) (ret interface{}) {
  // proxy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle5proxyEv
    // invoke: const QStyle * proxy()
    var ret0 = C.C_ZNK6QStyle5proxyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "const QStyle *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyle", "proxy", args)
  }

  return
}

// metaObject()
func (this *QStyle) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QStyle10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK6QStyle10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStyle", "metaObject", args)
  }

  return
}

// itemTextRect(const class QFontMetrics &, const class QRect &, int, _Bool, const class QString &)
func (this *QStyle) ItemTextRect(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString
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
    var ret0 = C.C_ZNK6QStyle12itemTextRectERK12QFontMetricsRK5QRectibRK7QString(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyle", "itemTextRect", args)
  }

  return
}

// unpolish(class QWidget *)
func (this *QStyle) Unpolish(args ...interface{}) () {
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
    // invoke: void unpolish(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle8unpolishEP7QWidget(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QStyle8unpolishEP12QApplication
    // invoke: void unpolish(class QApplication *)
    var arg0 = args[0].(*QApplication).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QStyle8unpolishEP12QApplication(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStyle", "unpolish", args)
  }

  return
}

// sliderPositionFromValue(int, int, int, int, _Bool)
func (this *QStyle) SliderPositionFromValue_s(args ...interface{}) (ret interface{}) {
  // sliderPositionFromValue(int, int, int, int, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QStyle23sliderPositionFromValueEiiiib
    // invoke: int sliderPositionFromValue(int, int, int, int, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.bool(args[4].(bool))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN6QStyle23sliderPositionFromValueEiiiib(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyle", "sliderPositionFromValue", args)
  }

  return
}

// itemPixmapRect(const class QRect &, int, const class QPixmap &)
func (this *QStyle) ItemPixmapRect(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap
    // invoke: QRect itemPixmapRect(const class QRect &, int, const class QPixmap &)
    var arg0 = args[0].(*qtcore.QRect).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK6QStyle14itemPixmapRectERK5QRectiRK7QPixmap(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStyle", "itemPixmapRect", args)
  }

  return
}

// <= body block end

