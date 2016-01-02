package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qinputmethod.h
// dst-file: /src/gui/qinputmethod.go
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
  // proto:  QRectF QInputMethod::inputItemRectangle();
extern void _ZNK12QInputMethod18inputItemRectangleEv(void* qthis);
  // proto:  const QMetaObject * QInputMethod::metaObject();
extern void _ZNK12QInputMethod10metaObjectEv(void* qthis);
  // proto:  QTransform QInputMethod::inputItemTransform();
extern void _ZNK12QInputMethod18inputItemTransformEv(void* qthis);
  // proto:  void QInputMethod::hide();
extern void _ZN12QInputMethod4hideEv(void* qthis);
  // proto:  QRectF QInputMethod::keyboardRectangle();
extern void _ZNK12QInputMethod17keyboardRectangleEv(void* qthis);
  // proto:  void QInputMethod::show();
extern void _ZN12QInputMethod4showEv(void* qthis);
  // proto:  void QInputMethod::QInputMethod();
extern void* dector_ZN12QInputMethodC1Ev();
extern void _ZN12QInputMethodC1Ev(void* qthis);
  // proto:  bool QInputMethod::isAnimating();
extern void _ZNK12QInputMethod11isAnimatingEv(void* qthis);
  // proto:  void QInputMethod::setVisible(bool visible);
extern void _ZN12QInputMethod10setVisibleEb(void* qthis, bool arg0);
  // proto:  void QInputMethod::setInputItemRectangle(const QRectF & rect);
extern void _ZN12QInputMethod21setInputItemRectangleERK6QRectF(void* qthis, void* arg0);
  // proto:  void QInputMethod::commit();
extern void _ZN12QInputMethod6commitEv(void* qthis);
  // proto:  void QInputMethod::setInputItemTransform(const QTransform & transform);
extern void _ZN12QInputMethod21setInputItemTransformERK10QTransform(void* qthis, void* arg0);
  // proto:  QRectF QInputMethod::cursorRectangle();
extern void _ZNK12QInputMethod15cursorRectangleEv(void* qthis);
  // proto:  bool QInputMethod::isVisible();
extern void _ZNK12QInputMethod9isVisibleEv(void* qthis);
  // proto:  void QInputMethod::~QInputMethod();
extern void _ZN12QInputMethodD0Ev(void* qthis);
  // proto:  QLocale QInputMethod::locale();
extern void _ZNK12QInputMethod6localeEv(void* qthis);
  // proto:  void QInputMethod::reset();
extern void _ZN12QInputMethod5resetEv(void* qthis);
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

// class sizeof(QInputMethod)=1
type QInputMethod struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _cursorRectangleChanged QInputMethod_cursorRectangleChanged_signal;
//  _localeChanged QInputMethod_localeChanged_signal;
//  _inputDirectionChanged QInputMethod_inputDirectionChanged_signal;
//  _animatingChanged QInputMethod_animatingChanged_signal;
//  _keyboardRectangleChanged QInputMethod_keyboardRectangleChanged_signal;
//  _visibleChanged QInputMethod_visibleChanged_signal;
}

  // proto:  QRectF QInputMethod::inputItemRectangle();
func (this *QInputMethod) inputItemRectangle(args ...interface{}) () {
  // inputItemRectangle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod18inputItemRectangleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "inputItemRectangle", args)
  }

}

  // proto:  const QMetaObject * QInputMethod::metaObject();
func (this *QInputMethod) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod10metaObjectEv
  default:
    qtrt.ErrorResolve("QInputMethod", "metaObject", args)
  }

}

  // proto:  QTransform QInputMethod::inputItemTransform();
func (this *QInputMethod) inputItemTransform(args ...interface{}) () {
  // inputItemTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod18inputItemTransformEv
  default:
    qtrt.ErrorResolve("QInputMethod", "inputItemTransform", args)
  }

}

  // proto:  void QInputMethod::hide();
func (this *QInputMethod) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod4hideEv
  default:
    qtrt.ErrorResolve("QInputMethod", "hide", args)
  }

}

  // proto:  QRectF QInputMethod::keyboardRectangle();
func (this *QInputMethod) keyboardRectangle(args ...interface{}) () {
  // keyboardRectangle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod17keyboardRectangleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "keyboardRectangle", args)
  }

}

  // proto:  void QInputMethod::show();
func (this *QInputMethod) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod4showEv
  default:
    qtrt.ErrorResolve("QInputMethod", "show", args)
  }

}

  // proto:  void QInputMethod::QInputMethod();
func NewQInputMethod(args ...interface{}) QInputMethod {
  return QInputMethod{}
}

  // proto:  bool QInputMethod::isAnimating();
func (this *QInputMethod) isAnimating(args ...interface{}) () {
  // isAnimating()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod11isAnimatingEv
  default:
    qtrt.ErrorResolve("QInputMethod", "isAnimating", args)
  }

}

  // proto:  void QInputMethod::setVisible(bool visible);
func (this *QInputMethod) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod10setVisibleEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QInputMethod", "setVisible", args)
  }

}

  // proto:  void QInputMethod::setInputItemRectangle(const QRectF & rect);
func (this *QInputMethod) setInputItemRectangle(args ...interface{}) () {
  // setInputItemRectangle(const class QRectF &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRectF{}) // "const QRectF &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod21setInputItemRectangleERK6QRectF
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QInputMethod", "setInputItemRectangle", args)
  }

}

  // proto:  void QInputMethod::commit();
func (this *QInputMethod) commit(args ...interface{}) () {
  // commit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod6commitEv
  default:
    qtrt.ErrorResolve("QInputMethod", "commit", args)
  }

}

  // proto:  void QInputMethod::setInputItemTransform(const QTransform & transform);
func (this *QInputMethod) setInputItemTransform(args ...interface{}) () {
  // setInputItemTransform(const class QTransform &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTransform{}) // "const QTransform &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod21setInputItemTransformERK10QTransform
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QInputMethod", "setInputItemTransform", args)
  }

}

  // proto:  QRectF QInputMethod::cursorRectangle();
func (this *QInputMethod) cursorRectangle(args ...interface{}) () {
  // cursorRectangle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod15cursorRectangleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "cursorRectangle", args)
  }

}

  // proto:  bool QInputMethod::isVisible();
func (this *QInputMethod) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod9isVisibleEv
  default:
    qtrt.ErrorResolve("QInputMethod", "isVisible", args)
  }

}

  // proto:  void QInputMethod::~QInputMethod();
func (this *QInputMethod) FreeQInputMethod(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QInputMethod", "~QInputMethod", args)
  }

}

  // proto:  QLocale QInputMethod::locale();
func (this *QInputMethod) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod6localeEv
  default:
    qtrt.ErrorResolve("QInputMethod", "locale", args)
  }

}

  // proto:  void QInputMethod::reset();
func (this *QInputMethod) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QInputMethod5resetEv
  default:
    qtrt.ErrorResolve("QInputMethod", "reset", args)
  }

}

// <= body block end

