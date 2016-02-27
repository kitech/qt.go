package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qopenglwindow.h
// dst-file: /src/gui/qopenglwindow.go
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QOpenGLWindow::UpdateBehavior QOpenGLWindow::updateBehavior();
extern void C_ZNK13QOpenGLWindow14updateBehaviorEv(void* qthis); // 4
  // proto:  QOpenGLContext * QOpenGLWindow::shareContext();
extern void C_ZNK13QOpenGLWindow12shareContextEv(void* qthis); // 4
  // proto:  void QOpenGLWindow::doneCurrent();
extern void C_ZN13QOpenGLWindow11doneCurrentEv(void* qthis); // 4
  // proto:  bool QOpenGLWindow::isValid();
extern bool C_ZNK13QOpenGLWindow7isValidEv(void* qthis); // 4
  // proto:  GLuint QOpenGLWindow::defaultFramebufferObject();
extern int32_t C_ZNK13QOpenGLWindow24defaultFramebufferObjectEv(void* qthis); // 4
  // proto:  void QOpenGLWindow::~QOpenGLWindow();
extern void C_ZN13QOpenGLWindowD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLWindow::metaObject();
extern void C_ZNK13QOpenGLWindow10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLWindow::makeCurrent();
extern void C_ZN13QOpenGLWindow11makeCurrentEv(void* qthis); // 4
  // proto:  QImage QOpenGLWindow::grabFramebuffer();
extern void* C_ZN13QOpenGLWindow15grabFramebufferEv(void* qthis); // 4
  // proto:  QOpenGLContext * QOpenGLWindow::context();
extern void C_ZNK13QOpenGLWindow7contextEv(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QOpenGLWindow)=1
type QOpenGLWindow struct {
  /*qbase*/ QPaintDeviceWindow;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _frameSwapped QOpenGLWindow_frameSwapped_signal;
}

// updateBehavior()
func (this *QOpenGLWindow) UpdateBehavior(args ...interface{}) () {
  // updateBehavior()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow14updateBehaviorEv
    // invoke: QOpenGLWindow::UpdateBehavior updateBehavior()
    C.C_ZNK13QOpenGLWindow14updateBehaviorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "updateBehavior", args)
  }

  return
}

// shareContext()
func (this *QOpenGLWindow) ShareContext(args ...interface{}) () {
  // shareContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow12shareContextEv
    // invoke: QOpenGLContext * shareContext()
    C.C_ZNK13QOpenGLWindow12shareContextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "shareContext", args)
  }

  return
}

// doneCurrent()
func (this *QOpenGLWindow) DoneCurrent(args ...interface{}) () {
  // doneCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow11doneCurrentEv
    // invoke: void doneCurrent()
    C.C_ZN13QOpenGLWindow11doneCurrentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "doneCurrent", args)
  }

  return
}

// isValid()
func (this *QOpenGLWindow) IsValid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK13QOpenGLWindow7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "isValid", args)
  }

  return
}

// defaultFramebufferObject()
func (this *QOpenGLWindow) DefaultFramebufferObject(args ...interface{}) (ret interface{}) {
  // defaultFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow24defaultFramebufferObjectEv
    // invoke: GLuint defaultFramebufferObject()
    var ret0 = C.C_ZNK13QOpenGLWindow24defaultFramebufferObjectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "defaultFramebufferObject", args)
  }

  return
}

// ~QOpenGLWindow()
func (this *QOpenGLWindow) Free(args ...interface{}) () {
  // ~QOpenGLWindow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindowD0Ev
    // invoke: void ~QOpenGLWindow()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QOpenGLWindowD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "~QOpenGLWindow", args)
  }

  return
}

// metaObject()
func (this *QOpenGLWindow) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QOpenGLWindow10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "metaObject", args)
  }

  return
}

// makeCurrent()
func (this *QOpenGLWindow) MakeCurrent(args ...interface{}) () {
  // makeCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow11makeCurrentEv
    // invoke: void makeCurrent()
    C.C_ZN13QOpenGLWindow11makeCurrentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "makeCurrent", args)
  }

  return
}

// grabFramebuffer()
func (this *QOpenGLWindow) GrabFramebuffer(args ...interface{}) (ret interface{}) {
  // grabFramebuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow15grabFramebufferEv
    // invoke: QImage grabFramebuffer()
    var ret0 = C.C_ZN13QOpenGLWindow15grabFramebufferEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "grabFramebuffer", args)
  }

  return
}

// context()
func (this *QOpenGLWindow) Context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow7contextEv
    // invoke: QOpenGLContext * context()
    C.C_ZNK13QOpenGLWindow7contextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "context", args)
  }

  return
}

// <= body block end

