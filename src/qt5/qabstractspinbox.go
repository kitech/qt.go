package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qabstractspinbox.h
// dst-file: /src/widgets/qabstractspinbox.go
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
  // proto:  void QAbstractSpinBox::stepBy(int steps);
extern void _ZN16QAbstractSpinBox6stepByEi(void* qthis, int arg0);
  // proto:  void QAbstractSpinBox::setReadOnly(bool r);
extern void _ZN16QAbstractSpinBox11setReadOnlyEb(void* qthis, bool arg0);
  // proto:  void QAbstractSpinBox::setFrame(bool );
extern void _ZN16QAbstractSpinBox8setFrameEb(void* qthis, bool arg0);
  // proto:  void QAbstractSpinBox::setSpecialValueText(const QString & txt);
extern void _ZN16QAbstractSpinBox19setSpecialValueTextERK7QString(void* qthis, void* arg0);
  // proto:  void QAbstractSpinBox::setAccelerated(bool on);
extern void _ZN16QAbstractSpinBox14setAcceleratedEb(void* qthis, bool arg0);
  // proto:  void QAbstractSpinBox::interpretText();
extern void _ZN16QAbstractSpinBox13interpretTextEv(void* qthis);
  // proto:  bool QAbstractSpinBox::event(QEvent * event);
extern void _ZN16QAbstractSpinBox5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  bool QAbstractSpinBox::keyboardTracking();
extern void _ZNK16QAbstractSpinBox16keyboardTrackingEv(void* qthis);
  // proto:  void QAbstractSpinBox::QAbstractSpinBox(const QAbstractSpinBox & );
extern void* dector_ZN16QAbstractSpinBoxC1ERKS_(void* arg0);
extern void _ZN16QAbstractSpinBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QAbstractSpinBox::metaObject();
extern void _ZNK16QAbstractSpinBox10metaObjectEv(void* qthis);
  // proto:  QSize QAbstractSpinBox::sizeHint();
extern void _ZNK16QAbstractSpinBox8sizeHintEv(void* qthis);
  // proto:  void QAbstractSpinBox::~QAbstractSpinBox();
extern void _ZN16QAbstractSpinBoxD0Ev(void* qthis);
  // proto:  void QAbstractSpinBox::fixup(QString & input);
extern void _ZNK16QAbstractSpinBox5fixupER7QString(void* qthis, void* arg0);
  // proto:  void QAbstractSpinBox::selectAll();
extern void _ZN16QAbstractSpinBox9selectAllEv(void* qthis);
  // proto:  void QAbstractSpinBox::stepDown();
extern void _ZN16QAbstractSpinBox8stepDownEv(void* qthis);
  // proto:  void QAbstractSpinBox::clear();
extern void _ZN16QAbstractSpinBox5clearEv(void* qthis);
  // proto:  QString QAbstractSpinBox::text();
extern void _ZNK16QAbstractSpinBox4textEv(void* qthis);
  // proto:  QString QAbstractSpinBox::specialValueText();
extern void _ZNK16QAbstractSpinBox16specialValueTextEv(void* qthis);
  // proto:  QSize QAbstractSpinBox::minimumSizeHint();
extern void _ZNK16QAbstractSpinBox15minimumSizeHintEv(void* qthis);
  // proto:  bool QAbstractSpinBox::wrapping();
extern void _ZNK16QAbstractSpinBox8wrappingEv(void* qthis);
  // proto:  void QAbstractSpinBox::QAbstractSpinBox(QWidget * parent);
extern void* dector_ZN16QAbstractSpinBoxC1EP7QWidget(void* arg0);
extern void _ZN16QAbstractSpinBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QAbstractSpinBox::stepUp();
extern void _ZN16QAbstractSpinBox6stepUpEv(void* qthis);
  // proto:  void QAbstractSpinBox::setWrapping(bool w);
extern void _ZN16QAbstractSpinBox11setWrappingEb(void* qthis, bool arg0);
  // proto:  void QAbstractSpinBox::setKeyboardTracking(bool kt);
extern void _ZN16QAbstractSpinBox19setKeyboardTrackingEb(void* qthis, bool arg0);
  // proto:  bool QAbstractSpinBox::isAccelerated();
extern void _ZNK16QAbstractSpinBox13isAcceleratedEv(void* qthis);
  // proto:  void QAbstractSpinBox::setGroupSeparatorShown(bool shown);
extern void _ZN16QAbstractSpinBox22setGroupSeparatorShownEb(void* qthis, bool arg0);
  // proto:  bool QAbstractSpinBox::isReadOnly();
extern void _ZNK16QAbstractSpinBox10isReadOnlyEv(void* qthis);
  // proto:  bool QAbstractSpinBox::hasAcceptableInput();
extern void _ZNK16QAbstractSpinBox18hasAcceptableInputEv(void* qthis);
  // proto:  bool QAbstractSpinBox::isGroupSeparatorShown();
extern void _ZNK16QAbstractSpinBox21isGroupSeparatorShownEv(void* qthis);
  // proto:  bool QAbstractSpinBox::hasFrame();
extern void _ZNK16QAbstractSpinBox8hasFrameEv(void* qthis);
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

// class sizeof(QAbstractSpinBox)=1
type QAbstractSpinBox struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _editingFinished QAbstractSpinBox_editingFinished_signal;
}

  // proto:  void QAbstractSpinBox::stepBy(int steps);
func (this *QAbstractSpinBox) stepBy(args ...interface{}) () {
  // stepBy(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox6stepByEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepBy", args)
  }

}

  // proto:  void QAbstractSpinBox::setReadOnly(bool r);
func (this *QAbstractSpinBox) setReadOnly(args ...interface{}) () {
  // setReadOnly(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox11setReadOnlyEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setReadOnly", args)
  }

}

  // proto:  void QAbstractSpinBox::setFrame(bool );
func (this *QAbstractSpinBox) setFrame(args ...interface{}) () {
  // setFrame(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox8setFrameEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setFrame", args)
  }

}

  // proto:  void QAbstractSpinBox::setSpecialValueText(const QString & txt);
func (this *QAbstractSpinBox) setSpecialValueText(args ...interface{}) () {
  // setSpecialValueText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox19setSpecialValueTextERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setSpecialValueText", args)
  }

}

  // proto:  void QAbstractSpinBox::setAccelerated(bool on);
func (this *QAbstractSpinBox) setAccelerated(args ...interface{}) () {
  // setAccelerated(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox14setAcceleratedEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setAccelerated", args)
  }

}

  // proto:  void QAbstractSpinBox::interpretText();
func (this *QAbstractSpinBox) interpretText(args ...interface{}) () {
  // interpretText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox13interpretTextEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "interpretText", args)
  }

}

  // proto:  bool QAbstractSpinBox::event(QEvent * event);
func (this *QAbstractSpinBox) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox5eventEP6QEvent
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "event", args)
  }

}

  // proto:  bool QAbstractSpinBox::keyboardTracking();
func (this *QAbstractSpinBox) keyboardTracking(args ...interface{}) () {
  // keyboardTracking()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox16keyboardTrackingEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "keyboardTracking", args)
  }

}

  // proto:  void QAbstractSpinBox::QAbstractSpinBox(const QAbstractSpinBox & );
func NewQAbstractSpinBox(args ...interface{}) QAbstractSpinBox {
  return QAbstractSpinBox{}
}

  // proto:  const QMetaObject * QAbstractSpinBox::metaObject();
func (this *QAbstractSpinBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox10metaObjectEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "metaObject", args)
  }

}

  // proto:  QSize QAbstractSpinBox::sizeHint();
func (this *QAbstractSpinBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8sizeHintEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "sizeHint", args)
  }

}

  // proto:  void QAbstractSpinBox::~QAbstractSpinBox();
func (this *QAbstractSpinBox) FreeQAbstractSpinBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "~QAbstractSpinBox", args)
  }

}

  // proto:  void QAbstractSpinBox::fixup(QString & input);
func (this *QAbstractSpinBox) fixup(args ...interface{}) () {
  // fixup(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox5fixupER7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "fixup", args)
  }

}

  // proto:  void QAbstractSpinBox::selectAll();
func (this *QAbstractSpinBox) selectAll(args ...interface{}) () {
  // selectAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox9selectAllEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "selectAll", args)
  }

}

  // proto:  void QAbstractSpinBox::stepDown();
func (this *QAbstractSpinBox) stepDown(args ...interface{}) () {
  // stepDown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox8stepDownEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepDown", args)
  }

}

  // proto:  void QAbstractSpinBox::clear();
func (this *QAbstractSpinBox) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox5clearEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "clear", args)
  }

}

  // proto:  QString QAbstractSpinBox::text();
func (this *QAbstractSpinBox) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox4textEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "text", args)
  }

}

  // proto:  QString QAbstractSpinBox::specialValueText();
func (this *QAbstractSpinBox) specialValueText(args ...interface{}) () {
  // specialValueText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox16specialValueTextEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "specialValueText", args)
  }

}

  // proto:  QSize QAbstractSpinBox::minimumSizeHint();
func (this *QAbstractSpinBox) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox15minimumSizeHintEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "minimumSizeHint", args)
  }

}

  // proto:  bool QAbstractSpinBox::wrapping();
func (this *QAbstractSpinBox) wrapping(args ...interface{}) () {
  // wrapping()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8wrappingEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "wrapping", args)
  }

}

  // proto:  void QAbstractSpinBox::stepUp();
func (this *QAbstractSpinBox) stepUp(args ...interface{}) () {
  // stepUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox6stepUpEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "stepUp", args)
  }

}

  // proto:  void QAbstractSpinBox::setWrapping(bool w);
func (this *QAbstractSpinBox) setWrapping(args ...interface{}) () {
  // setWrapping(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox11setWrappingEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setWrapping", args)
  }

}

  // proto:  void QAbstractSpinBox::setKeyboardTracking(bool kt);
func (this *QAbstractSpinBox) setKeyboardTracking(args ...interface{}) () {
  // setKeyboardTracking(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox19setKeyboardTrackingEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setKeyboardTracking", args)
  }

}

  // proto:  bool QAbstractSpinBox::isAccelerated();
func (this *QAbstractSpinBox) isAccelerated(args ...interface{}) () {
  // isAccelerated()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox13isAcceleratedEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isAccelerated", args)
  }

}

  // proto:  void QAbstractSpinBox::setGroupSeparatorShown(bool shown);
func (this *QAbstractSpinBox) setGroupSeparatorShown(args ...interface{}) () {
  // setGroupSeparatorShown(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAbstractSpinBox22setGroupSeparatorShownEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "setGroupSeparatorShown", args)
  }

}

  // proto:  bool QAbstractSpinBox::isReadOnly();
func (this *QAbstractSpinBox) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isReadOnly", args)
  }

}

  // proto:  bool QAbstractSpinBox::hasAcceptableInput();
func (this *QAbstractSpinBox) hasAcceptableInput(args ...interface{}) () {
  // hasAcceptableInput()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox18hasAcceptableInputEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "hasAcceptableInput", args)
  }

}

  // proto:  bool QAbstractSpinBox::isGroupSeparatorShown();
func (this *QAbstractSpinBox) isGroupSeparatorShown(args ...interface{}) () {
  // isGroupSeparatorShown()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox21isGroupSeparatorShownEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "isGroupSeparatorShown", args)
  }

}

  // proto:  bool QAbstractSpinBox::hasFrame();
func (this *QAbstractSpinBox) hasFrame(args ...interface{}) () {
  // hasFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAbstractSpinBox8hasFrameEv
  default:
    qtrt.ErrorResolve("QAbstractSpinBox", "hasFrame", args)
  }

}

// <= body block end

