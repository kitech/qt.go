package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  Qt::Orientation QProgressBar::orientation();
extern void _ZNK12QProgressBar11orientationEv(void* qthis); // 4
  // proto:  QString QProgressBar::text();
extern void _ZNK12QProgressBar4textEv(void* qthis); // 4
  // proto:  bool QProgressBar::isTextVisible();
extern void _ZNK12QProgressBar13isTextVisibleEv(void* qthis); // 4
  // proto:  void QProgressBar::QProgressBar(QWidget * parent);
extern void _ZN12QProgressBarC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  int QProgressBar::minimum();
extern void _ZNK12QProgressBar7minimumEv(void* qthis); // 4
  // proto:  Qt::Alignment QProgressBar::alignment();
extern void _ZNK12QProgressBar9alignmentEv(void* qthis); // 4
  // proto:  QProgressBar::Direction QProgressBar::textDirection();
extern void _ZNK12QProgressBar13textDirectionEv(void* qthis); // 4
  // proto:  void QProgressBar::resetFormat();
extern void _ZN12QProgressBar11resetFormatEv(void* qthis); // 4
  // proto:  void QProgressBar::~QProgressBar();
extern void _ZN12QProgressBarD2Ev(void* qthis); // 4
  // proto:  void QProgressBar::setFormat(const QString & format);
extern void _ZN12QProgressBar9setFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QProgressBar::setMaximum(int maximum);
extern void _ZN12QProgressBar10setMaximumEi(void* qthis, int32_t arg0); // 4
  // proto:  void QProgressBar::setValue(int value);
extern void _ZN12QProgressBar8setValueEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QProgressBar::format();
extern void _ZNK12QProgressBar6formatEv(void* qthis); // 4
  // proto:  QSize QProgressBar::sizeHint();
extern void _ZNK12QProgressBar8sizeHintEv(void* qthis); // 4
  // proto:  void QProgressBar::setRange(int minimum, int maximum);
extern void _ZN12QProgressBar8setRangeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QProgressBar::reset();
extern void _ZN12QProgressBar5resetEv(void* qthis); // 4
  // proto:  void QProgressBar::setInvertedAppearance(bool invert);
extern void _ZN12QProgressBar21setInvertedAppearanceEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QProgressBar::metaObject();
extern void _ZNK12QProgressBar10metaObjectEv(void* qthis); // 4
  // proto:  void QProgressBar::setTextVisible(bool visible);
extern void _ZN12QProgressBar14setTextVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QSize QProgressBar::minimumSizeHint();
extern void _ZNK12QProgressBar15minimumSizeHintEv(void* qthis); // 4
  // proto:  int QProgressBar::maximum();
extern void _ZNK12QProgressBar7maximumEv(void* qthis); // 4
  // proto:  int QProgressBar::value();
extern void _ZNK12QProgressBar5valueEv(void* qthis); // 4
  // proto:  bool QProgressBar::invertedAppearance();
extern void _ZNK12QProgressBar18invertedAppearanceEv(void* qthis); // 4
  // proto:  void QProgressBar::setMinimum(int minimum);
extern void _ZN12QProgressBar10setMinimumEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QProgressBar)=1
type QProgressBar struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _valueChanged QProgressBar_valueChanged_signal;
}

// orientation()
func (this *QProgressBar) orientation(args ...interface{}) () {
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
    C._ZNK12QProgressBar11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "orientation", args)
  }

}

// text()
func (this *QProgressBar) text(args ...interface{}) () {
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
    C._ZNK12QProgressBar4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "text", args)
  }

}

// isTextVisible()
func (this *QProgressBar) isTextVisible(args ...interface{}) () {
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
    C._ZNK12QProgressBar13isTextVisibleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "isTextVisible", args)
  }

}

// QProgressBar(class QWidget *)
func NewQProgressBar(args ...interface{}) QProgressBar {
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
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QProgressBarC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "QProgressBar", args)
  }

  return QProgressBar{}
}

// minimum()
func (this *QProgressBar) minimum(args ...interface{}) () {
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
    C._ZNK12QProgressBar7minimumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "minimum", args)
  }

}

// alignment()
func (this *QProgressBar) alignment(args ...interface{}) () {
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
    C._ZNK12QProgressBar9alignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "alignment", args)
  }

}

// textDirection()
func (this *QProgressBar) textDirection(args ...interface{}) () {
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
    C._ZNK12QProgressBar13textDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "textDirection", args)
  }

}

// resetFormat()
func (this *QProgressBar) resetFormat(args ...interface{}) () {
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
    C._ZN12QProgressBar11resetFormatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "resetFormat", args)
  }

}

// ~QProgressBar()
func (this *QProgressBar) FreeQProgressBar(args ...interface{}) () {
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
    C._ZN12QProgressBarD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "~QProgressBar", args)
  }

}

// setFormat(const class QString &)
func (this *QProgressBar) setFormat(args ...interface{}) () {
  // setFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QProgressBar9setFormatERK7QString
    // invoke: void setFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QProgressBar9setFormatERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setFormat", args)
  }

}

// setMaximum(int)
func (this *QProgressBar) setMaximum(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN12QProgressBar10setMaximumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setMaximum", args)
  }

}

// setValue(int)
func (this *QProgressBar) setValue(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN12QProgressBar8setValueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setValue", args)
  }

}

// format()
func (this *QProgressBar) format(args ...interface{}) () {
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
    C._ZNK12QProgressBar6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "format", args)
  }

}

// sizeHint()
func (this *QProgressBar) sizeHint(args ...interface{}) () {
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
    C._ZNK12QProgressBar8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "sizeHint", args)
  }

}

// setRange(int, int)
func (this *QProgressBar) setRange(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN12QProgressBar8setRangeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QProgressBar", "setRange", args)
  }

}

// reset()
func (this *QProgressBar) reset(args ...interface{}) () {
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
    C._ZN12QProgressBar5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "reset", args)
  }

}

// setInvertedAppearance(_Bool)
func (this *QProgressBar) setInvertedAppearance(args ...interface{}) () {
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
    C._ZN12QProgressBar21setInvertedAppearanceEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setInvertedAppearance", args)
  }

}

// metaObject()
func (this *QProgressBar) metaObject(args ...interface{}) () {
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
    C._ZNK12QProgressBar10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "metaObject", args)
  }

}

// setTextVisible(_Bool)
func (this *QProgressBar) setTextVisible(args ...interface{}) () {
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
    C._ZN12QProgressBar14setTextVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setTextVisible", args)
  }

}

// minimumSizeHint()
func (this *QProgressBar) minimumSizeHint(args ...interface{}) () {
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
    C._ZNK12QProgressBar15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "minimumSizeHint", args)
  }

}

// maximum()
func (this *QProgressBar) maximum(args ...interface{}) () {
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
    C._ZNK12QProgressBar7maximumEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "maximum", args)
  }

}

// value()
func (this *QProgressBar) value(args ...interface{}) () {
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
    C._ZNK12QProgressBar5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "value", args)
  }

}

// invertedAppearance()
func (this *QProgressBar) invertedAppearance(args ...interface{}) () {
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
    C._ZNK12QProgressBar18invertedAppearanceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QProgressBar", "invertedAppearance", args)
  }

}

// setMinimum(int)
func (this *QProgressBar) setMinimum(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN12QProgressBar10setMinimumEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QProgressBar", "setMinimum", args)
  }

}

// <= body block end

