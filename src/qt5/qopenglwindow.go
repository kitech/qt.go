package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  bool QOpenGLWindow::isValid();
extern void _ZNK13QOpenGLWindow7isValidEv(void* qthis);
  // proto:  QImage QOpenGLWindow::grabFramebuffer();
extern void _ZN13QOpenGLWindow15grabFramebufferEv(void* qthis);
  // proto:  void QOpenGLWindow::QOpenGLWindow(const QOpenGLWindow & );
extern void* dector_ZN13QOpenGLWindowC1ERKS_(void* arg0);
extern void _ZN13QOpenGLWindowC1ERKS_(void* qthis, void* arg0);
  // proto:  QOpenGLContext * QOpenGLWindow::shareContext();
extern void _ZNK13QOpenGLWindow12shareContextEv(void* qthis);
  // proto:  void QOpenGLWindow::makeCurrent();
extern void _ZN13QOpenGLWindow11makeCurrentEv(void* qthis);
  // proto:  QOpenGLContext * QOpenGLWindow::context();
extern void _ZNK13QOpenGLWindow7contextEv(void* qthis);
  // proto:  void QOpenGLWindow::doneCurrent();
extern void _ZN13QOpenGLWindow11doneCurrentEv(void* qthis);
  // proto:  GLuint QOpenGLWindow::defaultFramebufferObject();
extern void _ZNK13QOpenGLWindow24defaultFramebufferObjectEv(void* qthis);
  // proto:  void QOpenGLWindow::~QOpenGLWindow();
extern void _ZN13QOpenGLWindowD0Ev(void* qthis);
  // proto:  const QMetaObject * QOpenGLWindow::metaObject();
extern void _ZNK13QOpenGLWindow10metaObjectEv(void* qthis);
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

// class sizeof(QOpenGLWindow)=1
type QOpenGLWindow struct {
  /*qbase*/ QPaintDeviceWindow;
  qclsinst uint64 /* *mut c_void*/;
//  _frameSwapped QOpenGLWindow_frameSwapped_signal;
}

  // proto:  bool QOpenGLWindow::isValid();
func (this *QOpenGLWindow) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow7isValidEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "isValid", args)
  }

}

  // proto:  QImage QOpenGLWindow::grabFramebuffer();
func (this *QOpenGLWindow) grabFramebuffer(args ...interface{}) () {
  // grabFramebuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow15grabFramebufferEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "grabFramebuffer", args)
  }

}

  // proto:  void QOpenGLWindow::QOpenGLWindow(const QOpenGLWindow & );
func NewQOpenGLWindow(args ...interface{}) QOpenGLWindow {
  return QOpenGLWindow{}
}

  // proto:  QOpenGLContext * QOpenGLWindow::shareContext();
func (this *QOpenGLWindow) shareContext(args ...interface{}) () {
  // shareContext()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow12shareContextEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "shareContext", args)
  }

}

  // proto:  void QOpenGLWindow::makeCurrent();
func (this *QOpenGLWindow) makeCurrent(args ...interface{}) () {
  // makeCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow11makeCurrentEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "makeCurrent", args)
  }

}

  // proto:  QOpenGLContext * QOpenGLWindow::context();
func (this *QOpenGLWindow) context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow7contextEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "context", args)
  }

}

  // proto:  void QOpenGLWindow::doneCurrent();
func (this *QOpenGLWindow) doneCurrent(args ...interface{}) () {
  // doneCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWindow11doneCurrentEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "doneCurrent", args)
  }

}

  // proto:  GLuint QOpenGLWindow::defaultFramebufferObject();
func (this *QOpenGLWindow) defaultFramebufferObject(args ...interface{}) () {
  // defaultFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow24defaultFramebufferObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "defaultFramebufferObject", args)
  }

}

  // proto:  void QOpenGLWindow::~QOpenGLWindow();
func (this *QOpenGLWindow) FreeQOpenGLWindow(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "~QOpenGLWindow", args)
  }

}

  // proto:  const QMetaObject * QOpenGLWindow::metaObject();
func (this *QOpenGLWindow) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWindow10metaObjectEv
  default:
    qtrt.ErrorResolve("QOpenGLWindow", "metaObject", args)
  }

}

// <= body block end

