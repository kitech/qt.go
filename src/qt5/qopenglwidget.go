package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qopenglwidget.h
// dst-file: /src/widgets/qopenglwidget.go
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
  // proto:  QOpenGLWidget::UpdateBehavior QOpenGLWidget::updateBehavior();
extern void C_ZNK13QOpenGLWidget14updateBehaviorEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::doneCurrent();
extern void C_ZN13QOpenGLWidget11doneCurrentEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::setFormat(const QSurfaceFormat & format);
extern void C_ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat(void* qthis, void* arg0); // 4
  // proto:  QSurfaceFormat QOpenGLWidget::format();
extern void C_ZNK13QOpenGLWidget6formatEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::~QOpenGLWidget();
extern void C_ZN13QOpenGLWidgetD2Ev(void* qthis); // 4
  // proto:  bool QOpenGLWidget::isValid();
extern void C_ZNK13QOpenGLWidget7isValidEv(void* qthis); // 4
  // proto:  GLuint QOpenGLWidget::defaultFramebufferObject();
extern void C_ZNK13QOpenGLWidget24defaultFramebufferObjectEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLWidget::metaObject();
extern void C_ZNK13QOpenGLWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::makeCurrent();
extern void C_ZN13QOpenGLWidget11makeCurrentEv(void* qthis); // 4
  // proto:  QImage QOpenGLWidget::grabFramebuffer();
extern void C_ZN13QOpenGLWidget15grabFramebufferEv(void* qthis); // 4
  // proto:  QOpenGLContext * QOpenGLWidget::context();
extern void C_ZNK13QOpenGLWidget7contextEv(void* qthis); // 4
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

// class sizeof(QOpenGLWidget)=1
type QOpenGLWidget struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToResize QOpenGLWidget_aboutToResize_signal;
//  _resized QOpenGLWidget_resized_signal;
//  _frameSwapped QOpenGLWidget_frameSwapped_signal;
//  _aboutToCompose QOpenGLWidget_aboutToCompose_signal;
}

// updateBehavior()
func (this *QOpenGLWidget) updateBehavior(args ...interface{}) () {
  // updateBehavior()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget14updateBehaviorEv
    // invoke: QOpenGLWidget::UpdateBehavior updateBehavior()
    C.C_ZNK13QOpenGLWidget14updateBehaviorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "updateBehavior", args)
  }

}

// doneCurrent()
func (this *QOpenGLWidget) doneCurrent(args ...interface{}) () {
  // doneCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget11doneCurrentEv
    // invoke: void doneCurrent()
    C.C_ZN13QOpenGLWidget11doneCurrentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "doneCurrent", args)
  }

}

// setFormat(const class QSurfaceFormat &)
func (this *QOpenGLWidget) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat
    // invoke: void setFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "setFormat", args)
  }

}

// format()
func (this *QOpenGLWidget) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget6formatEv
    // invoke: QSurfaceFormat format()
    var ret = C.C_ZNK13QOpenGLWidget6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "format", args)
  }

}

// ~QOpenGLWidget()
func (this *QOpenGLWidget) FreeQOpenGLWidget(args ...interface{}) () {
  // ~QOpenGLWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidgetD0Ev
    // invoke: void ~QOpenGLWidget()
    C.C_ZN13QOpenGLWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "~QOpenGLWidget", args)
  }

}

// isValid()
func (this *QOpenGLWidget) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget7isValidEv
    // invoke: bool isValid()
    var ret = C.C_ZNK13QOpenGLWidget7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "isValid", args)
  }

}

// defaultFramebufferObject()
func (this *QOpenGLWidget) defaultFramebufferObject(args ...interface{}) () {
  // defaultFramebufferObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget24defaultFramebufferObjectEv
    // invoke: GLuint defaultFramebufferObject()
    var ret = C.C_ZNK13QOpenGLWidget24defaultFramebufferObjectEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "defaultFramebufferObject", args)
  }

}

// metaObject()
func (this *QOpenGLWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QOpenGLWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "metaObject", args)
  }

}

// makeCurrent()
func (this *QOpenGLWidget) makeCurrent(args ...interface{}) () {
  // makeCurrent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget11makeCurrentEv
    // invoke: void makeCurrent()
    C.C_ZN13QOpenGLWidget11makeCurrentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "makeCurrent", args)
  }

}

// grabFramebuffer()
func (this *QOpenGLWidget) grabFramebuffer(args ...interface{}) () {
  // grabFramebuffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget15grabFramebufferEv
    // invoke: QImage grabFramebuffer()
    var ret = C.C_ZN13QOpenGLWidget15grabFramebufferEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "grabFramebuffer", args)
  }

}

// context()
func (this *QOpenGLWidget) context(args ...interface{}) () {
  // context()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QOpenGLWidget7contextEv
    // invoke: QOpenGLContext * context()
    C.C_ZNK13QOpenGLWidget7contextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "context", args)
  }

}

// <= body block end

