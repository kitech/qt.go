package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qprogressbar.h
// dst-file: /src/widgets/qprogressbar.go
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
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  Qt::Orientation QProgressBar::orientation();
extern void C_ZNK12QProgressBar11orientationEv(void* qthis); // 4
  // proto:  QString QProgressBar::text();
extern void* C_ZNK12QProgressBar4textEv(void* qthis); // 4
  // proto:  bool QProgressBar::isTextVisible();
extern bool C_ZNK12QProgressBar13isTextVisibleEv(void* qthis); // 4
  // proto:  void QProgressBar::QProgressBar(QWidget * parent);
extern void* C_ZN12QProgressBarC2EP7QWidget(void* arg0); // 3
  // proto:  int QProgressBar::minimum();
extern int32_t C_ZNK12QProgressBar7minimumEv(void* qthis); // 4
  // proto:  Qt::Alignment QProgressBar::alignment();
extern void C_ZNK12QProgressBar9alignmentEv(void* qthis); // 4
  // proto:  QProgressBar::Direction QProgressBar::textDirection();
extern void C_ZNK12QProgressBar13textDirectionEv(void* qthis); // 4
  // proto:  void QProgressBar::resetFormat();
extern void C_ZN12QProgressBar11resetFormatEv(void* qthis); // 4
  // proto:  void QProgressBar::~QProgressBar();
extern void C_ZN12QProgressBarD2Ev(void* qthis); // 4
  // proto:  void QProgressBar::setFormat(const QString & format);
extern void C_ZN12QProgressBar9setFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProgressBar::setMaximum(int maximum);
extern void C_ZN12QProgressBar10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressBar::setValue(int value);
extern void C_ZN12QProgressBar8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QProgressBar::format();
extern void* C_ZNK12QProgressBar6formatEv(void* qthis); // 4
  // proto:  QSize QProgressBar::sizeHint();
extern void* C_ZNK12QProgressBar8sizeHintEv(void* qthis); // 4
  // proto:  void QProgressBar::setRange(int minimum, int maximum);
extern void C_ZN12QProgressBar8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QProgressBar::reset();
extern void C_ZN12QProgressBar5resetEv(void* qthis); // 4
  // proto:  void QProgressBar::setInvertedAppearance(bool invert);
extern void C_ZN12QProgressBar21setInvertedAppearanceEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QProgressBar::metaObject();
extern void C_ZNK12QProgressBar10metaObjectEv(void* qthis); // 4
  // proto:  void QProgressBar::setTextVisible(bool visible);
extern void C_ZN12QProgressBar14setTextVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QSize QProgressBar::minimumSizeHint();
extern void* C_ZNK12QProgressBar15minimumSizeHintEv(void* qthis); // 4
  // proto:  int QProgressBar::maximum();
extern int32_t C_ZNK12QProgressBar7maximumEv(void* qthis); // 4
  // proto:  int QProgressBar::value();
extern int32_t C_ZNK12QProgressBar5valueEv(void* qthis); // 4
  // proto:  bool QProgressBar::invertedAppearance();
extern bool C_ZNK12QProgressBar18invertedAppearanceEv(void* qthis); // 4
  // proto:  void QProgressBar::setMinimum(int minimum);
extern void C_ZN12QProgressBar10setMinimumEi(void* qthis, int32_t arg0); // 4
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
}

// class sizeof(QProgressBar)=1
type QProgressBar struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QProgressBar_valueChanged_signal;
}

// orientation()
func (this *QProgressBar) Orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar11orientationEv
    // invoke: Qt::Orientation orientation()
    C.C_ZNK12QProgressBar11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "orientation", args)
  }

  return
}

// text()
func (this *QProgressBar) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK12QProgressBar4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "text", args)
  }

  return
}

// isTextVisible()
func (this *QProgressBar) Istextvisible(args ...interface{}) (ret interface{}) {
  // isTextVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar13isTextVisibleEv
    // invoke: bool isTextVisible()
    var ret0 = C.C_ZNK12QProgressBar13isTextVisibleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "isTextVisible", args)
  }

  return
}

// QProgressBar(class QWidget *)
func NewQProgressBar(args ...interface{}) *QProgressBar {
  // QProgressBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBarC1EP7QWidget
    // invoke: void QProgressBar(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QProgressBarC2EP7QWidget(arg0)
    return &QProgressBar{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QProgressBar", "QProgressBar", args)
  }

  return nil // QProgressBar{Qclsinst:qthis}
}

// minimum()
func (this *QProgressBar) Minimum(args ...interface{}) (ret interface{}) {
  // minimum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar7minimumEv
    // invoke: int minimum()
    var ret0 = C.C_ZNK12QProgressBar7minimumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "minimum", args)
  }

  return
}

// alignment()
func (this *QProgressBar) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK12QProgressBar9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "alignment", args)
  }

  return
}

// textDirection()
func (this *QProgressBar) Textdirection(args ...interface{}) () {
  // textDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar13textDirectionEv
    // invoke: QProgressBar::Direction textDirection()
    C.C_ZNK12QProgressBar13textDirectionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "textDirection", args)
  }

  return
}

// resetFormat()
func (this *QProgressBar) Resetformat(args ...interface{}) () {
  // resetFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar11resetFormatEv
    // invoke: void resetFormat()
    C.C_ZN12QProgressBar11resetFormatEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "resetFormat", args)
  }

  return
}

// ~QProgressBar()
func (this *QProgressBar) Freeqprogressbar(args ...interface{}) () {
  // ~QProgressBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBarD0Ev
    // invoke: void ~QProgressBar()
    C.C_ZN12QProgressBarD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "~QProgressBar", args)
  }

  return
}

// setFormat(const class QString &)
func (this *QProgressBar) Setformat(args ...interface{}) () {
  // setFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar9setFormatERK7QString
    // invoke: void setFormat(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QProgressBar9setFormatERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setFormat", args)
  }

  return
}

// setMaximum(int)
func (this *QProgressBar) Setmaximum(args ...interface{}) () {
  // setMaximum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar10setMaximumEi
    // invoke: void setMaximum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QProgressBar10setMaximumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setMaximum", args)
  }

  return
}

// setValue(int)
func (this *QProgressBar) Setvalue(args ...interface{}) () {
  // setValue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar8setValueEi
    // invoke: void setValue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QProgressBar8setValueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setValue", args)
  }

  return
}

// format()
func (this *QProgressBar) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar6formatEv
    // invoke: QString format()
    var ret0 = C.C_ZNK12QProgressBar6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "format", args)
  }

  return
}

// sizeHint()
func (this *QProgressBar) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK12QProgressBar8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "sizeHint", args)
  }

  return
}

// setRange(int, int)
func (this *QProgressBar) Setrange(args ...interface{}) () {
  // setRange(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar8setRangeEii
    // invoke: void setRange(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QProgressBar8setRangeEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProgressBar", "setRange", args)
  }

  return
}

// reset()
func (this *QProgressBar) Reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar5resetEv
    // invoke: void reset()
    C.C_ZN12QProgressBar5resetEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "reset", args)
  }

  return
}

// setInvertedAppearance(_Bool)
func (this *QProgressBar) Setinvertedappearance(args ...interface{}) () {
  // setInvertedAppearance(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar21setInvertedAppearanceEb
    // invoke: void setInvertedAppearance(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QProgressBar21setInvertedAppearanceEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setInvertedAppearance", args)
  }

  return
}

// metaObject()
func (this *QProgressBar) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QProgressBar10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "metaObject", args)
  }

  return
}

// setTextVisible(_Bool)
func (this *QProgressBar) Settextvisible(args ...interface{}) () {
  // setTextVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar14setTextVisibleEb
    // invoke: void setTextVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QProgressBar14setTextVisibleEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setTextVisible", args)
  }

  return
}

// minimumSizeHint()
func (this *QProgressBar) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK12QProgressBar15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "minimumSizeHint", args)
  }

  return
}

// maximum()
func (this *QProgressBar) Maximum(args ...interface{}) (ret interface{}) {
  // maximum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar7maximumEv
    // invoke: int maximum()
    var ret0 = C.C_ZNK12QProgressBar7maximumEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "maximum", args)
  }

  return
}

// value()
func (this *QProgressBar) Value(args ...interface{}) (ret interface{}) {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar5valueEv
    // invoke: int value()
    var ret0 = C.C_ZNK12QProgressBar5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "value", args)
  }

  return
}

// invertedAppearance()
func (this *QProgressBar) Invertedappearance(args ...interface{}) (ret interface{}) {
  // invertedAppearance()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QProgressBar18invertedAppearanceEv
    // invoke: bool invertedAppearance()
    var ret0 = C.C_ZNK12QProgressBar18invertedAppearanceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QProgressBar", "invertedAppearance", args)
  }

  return
}

// setMinimum(int)
func (this *QProgressBar) Setminimum(args ...interface{}) () {
  // setMinimum(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar10setMinimumEi
    // invoke: void setMinimum(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QProgressBar10setMinimumEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setMinimum", args)
  }

  return
}

// <= body block end

