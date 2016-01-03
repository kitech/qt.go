package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  QString QProgressBar::format();
extern void _ZNK12QProgressBar6formatEv(void* qthis);
  // proto:  void QProgressBar::reset();
extern void _ZN12QProgressBar5resetEv(void* qthis);
  // proto:  const QMetaObject * QProgressBar::metaObject();
extern void _ZNK12QProgressBar10metaObjectEv(void* qthis);
  // proto:  int QProgressBar::maximum();
extern void _ZNK12QProgressBar7maximumEv(void* qthis);
  // proto:  void QProgressBar::setFormat(const QString & format);
extern void _ZN12QProgressBar9setFormatERK7QString(void* qthis, void* arg0);
  // proto:  bool QProgressBar::invertedAppearance();
extern void _ZNK12QProgressBar18invertedAppearanceEv(void* qthis);
  // proto:  QString QProgressBar::text();
extern void _ZNK12QProgressBar4textEv(void* qthis);
  // proto:  void QProgressBar::setMinimum(int minimum);
extern void _ZN12QProgressBar10setMinimumEi(void* qthis, int32_t arg0);
  // proto:  void QProgressBar::QProgressBar(const QProgressBar & );
extern void* dector_ZN12QProgressBarC1ERKS_(void* arg0);
extern void _ZN12QProgressBarC1ERKS_(void* qthis, void* arg0);
  // proto:  void QProgressBar::QProgressBar(QWidget * parent);
extern void* dector_ZN12QProgressBarC1EP7QWidget(void* arg0);
extern void _ZN12QProgressBarC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QProgressBar::setTextVisible(bool visible);
extern void _ZN12QProgressBar14setTextVisibleEb(void* qthis, bool arg0);
  // proto:  int QProgressBar::value();
extern void _ZNK12QProgressBar5valueEv(void* qthis);
  // proto:  void QProgressBar::setValue(int value);
extern void _ZN12QProgressBar8setValueEi(void* qthis, int32_t arg0);
  // proto:  QSize QProgressBar::minimumSizeHint();
extern void _ZNK12QProgressBar15minimumSizeHintEv(void* qthis);
  // proto:  int QProgressBar::minimum();
extern void _ZNK12QProgressBar7minimumEv(void* qthis);
  // proto:  void QProgressBar::setRange(int minimum, int maximum);
extern void _ZN12QProgressBar8setRangeEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QSize QProgressBar::sizeHint();
extern void _ZNK12QProgressBar8sizeHintEv(void* qthis);
  // proto:  void QProgressBar::resetFormat();
extern void _ZN12QProgressBar11resetFormatEv(void* qthis);
  // proto:  bool QProgressBar::isTextVisible();
extern void _ZNK12QProgressBar13isTextVisibleEv(void* qthis);
  // proto:  void QProgressBar::setInvertedAppearance(bool invert);
extern void _ZN12QProgressBar21setInvertedAppearanceEb(void* qthis, bool arg0);
  // proto:  void QProgressBar::~QProgressBar();
extern void _ZN12QProgressBarD0Ev(void* qthis);
  // proto:  void QProgressBar::setMaximum(int maximum);
extern void _ZN12QProgressBar10setMaximumEi(void* qthis, int32_t arg0);
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

  // proto:  QString QProgressBar::format();
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

  // proto:  void QProgressBar::reset();
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

  // proto:  const QMetaObject * QProgressBar::metaObject();
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

  // proto:  int QProgressBar::maximum();
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

  // proto:  void QProgressBar::setFormat(const QString & format);
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

  // proto:  bool QProgressBar::invertedAppearance();
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

  // proto:  QString QProgressBar::text();
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

  // proto:  void QProgressBar::setMinimum(int minimum);
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

  // proto:  void QProgressBar::QProgressBar(const QProgressBar & );
func NewQProgressBar(args ...interface{}) QProgressBar {
  return QProgressBar{}
}

  // proto:  void QProgressBar::setTextVisible(bool visible);
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

  // proto:  int QProgressBar::value();
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

  // proto:  void QProgressBar::setValue(int value);
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

  // proto:  QSize QProgressBar::minimumSizeHint();
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

  // proto:  int QProgressBar::minimum();
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

  // proto:  void QProgressBar::setRange(int minimum, int maximum);
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

  // proto:  QSize QProgressBar::sizeHint();
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

  // proto:  void QProgressBar::resetFormat();
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

  // proto:  bool QProgressBar::isTextVisible();
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

  // proto:  void QProgressBar::setInvertedAppearance(bool invert);
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

  // proto:  void QProgressBar::~QProgressBar();
func (this *QProgressBar) FreeQProgressBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QProgressBar", "~QProgressBar", args)
  }

}

  // proto:  void QProgressBar::setMaximum(int maximum);
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

// <= body block end

