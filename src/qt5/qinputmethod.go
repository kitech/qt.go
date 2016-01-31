package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QInputMethod::show();
extern void C_ZN12QInputMethod4showEv(void* qthis); // 4
  // proto:  QLocale QInputMethod::locale();
extern void C_ZNK12QInputMethod6localeEv(void* qthis); // 4
  // proto:  void QInputMethod::setInputItemRectangle(const QRectF & rect);
extern void C_ZN12QInputMethod21setInputItemRectangleERK6QRectF(void* qthis, void* arg0); // 4
  // proto:  void QInputMethod::hide();
extern void C_ZN12QInputMethod4hideEv(void* qthis); // 4
  // proto:  QRectF QInputMethod::cursorRectangle();
extern void C_ZNK12QInputMethod15cursorRectangleEv(void* qthis); // 4
  // proto:  QRectF QInputMethod::inputItemRectangle();
extern void C_ZNK12QInputMethod18inputItemRectangleEv(void* qthis); // 4
  // proto:  bool QInputMethod::isAnimating();
extern void C_ZNK12QInputMethod11isAnimatingEv(void* qthis); // 4
  // proto:  Qt::LayoutDirection QInputMethod::inputDirection();
extern void C_ZNK12QInputMethod14inputDirectionEv(void* qthis); // 4
  // proto:  void QInputMethod::setInputItemTransform(const QTransform & transform);
extern void C_ZN12QInputMethod21setInputItemTransformERK10QTransform(void* qthis, void* arg0); // 4
  // proto:  QRectF QInputMethod::keyboardRectangle();
extern void C_ZNK12QInputMethod17keyboardRectangleEv(void* qthis); // 4
  // proto:  void QInputMethod::reset();
extern void C_ZN12QInputMethod5resetEv(void* qthis); // 4
  // proto:  const QMetaObject * QInputMethod::metaObject();
extern void C_ZNK12QInputMethod10metaObjectEv(void* qthis); // 4
  // proto:  QTransform QInputMethod::inputItemTransform();
extern void C_ZNK12QInputMethod18inputItemTransformEv(void* qthis); // 4
  // proto:  bool QInputMethod::isVisible();
extern void C_ZNK12QInputMethod9isVisibleEv(void* qthis); // 4
  // proto:  void QInputMethod::commit();
extern void C_ZN12QInputMethod6commitEv(void* qthis); // 4
  // proto:  void QInputMethod::setVisible(bool visible);
extern void C_ZN12QInputMethod10setVisibleEb(void* qthis, bool arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _cursorRectangleChanged QInputMethod_cursorRectangleChanged_signal;
//  _localeChanged QInputMethod_localeChanged_signal;
//  _inputDirectionChanged QInputMethod_inputDirectionChanged_signal;
//  _animatingChanged QInputMethod_animatingChanged_signal;
//  _keyboardRectangleChanged QInputMethod_keyboardRectangleChanged_signal;
//  _visibleChanged QInputMethod_visibleChanged_signal;
}

// show()
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
    // invoke: void show()
    C.C_ZN12QInputMethod4showEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethod", "show", args)
  }

}

// locale()
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
    // invoke: QLocale locale()
    var ret = C.C_ZNK12QInputMethod6localeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "locale", args)
  }

}

// setInputItemRectangle(const class QRectF &)
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
    // invoke: void setInputItemRectangle(const class QRectF &)
    var arg0 = args[0].(QRectF).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputMethod21setInputItemRectangleERK6QRectF(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputMethod", "setInputItemRectangle", args)
  }

}

// hide()
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
    // invoke: void hide()
    C.C_ZN12QInputMethod4hideEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethod", "hide", args)
  }

}

// cursorRectangle()
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
    // invoke: QRectF cursorRectangle()
    var ret = C.C_ZNK12QInputMethod15cursorRectangleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "cursorRectangle", args)
  }

}

// inputItemRectangle()
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
    // invoke: QRectF inputItemRectangle()
    var ret = C.C_ZNK12QInputMethod18inputItemRectangleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "inputItemRectangle", args)
  }

}

// isAnimating()
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
    // invoke: bool isAnimating()
    var ret = C.C_ZNK12QInputMethod11isAnimatingEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "isAnimating", args)
  }

}

// inputDirection()
func (this *QInputMethod) inputDirection(args ...interface{}) () {
  // inputDirection()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QInputMethod14inputDirectionEv
    // invoke: Qt::LayoutDirection inputDirection()
    C.C_ZNK12QInputMethod14inputDirectionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethod", "inputDirection", args)
  }

}

// setInputItemTransform(const class QTransform &)
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
    // invoke: void setInputItemTransform(const class QTransform &)
    var arg0 = args[0].(QTransform).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputMethod21setInputItemTransformERK10QTransform(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputMethod", "setInputItemTransform", args)
  }

}

// keyboardRectangle()
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
    // invoke: QRectF keyboardRectangle()
    var ret = C.C_ZNK12QInputMethod17keyboardRectangleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "keyboardRectangle", args)
  }

}

// reset()
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
    // invoke: void reset()
    C.C_ZN12QInputMethod5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethod", "reset", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QInputMethod10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethod", "metaObject", args)
  }

}

// inputItemTransform()
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
    // invoke: QTransform inputItemTransform()
    var ret = C.C_ZNK12QInputMethod18inputItemTransformEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "inputItemTransform", args)
  }

}

// isVisible()
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
    // invoke: bool isVisible()
    var ret = C.C_ZNK12QInputMethod9isVisibleEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QInputMethod", "isVisible", args)
  }

}

// commit()
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
    // invoke: void commit()
    C.C_ZN12QInputMethod6commitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QInputMethod", "commit", args)
  }

}

// setVisible(_Bool)
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
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QInputMethod10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QInputMethod", "setVisible", args)
  }

}

// <= body block end

