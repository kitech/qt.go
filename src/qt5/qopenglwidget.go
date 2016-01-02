package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QOpenGLWidget::~QOpenGLWidget();
extern void _ZN13QOpenGLWidgetD0Ev(void* qthis);
  // proto:  GLuint QOpenGLWidget::defaultFramebufferObject();
extern void _ZNK13QOpenGLWidget24defaultFramebufferObjectEv(void* qthis);
  // proto:  void QOpenGLWidget::QOpenGLWidget(const QOpenGLWidget & );
extern void* dector_ZN13QOpenGLWidgetC1ERKS_(void* arg0);
extern void _ZN13QOpenGLWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QOpenGLWidget::isValid();
extern void _ZNK13QOpenGLWidget7isValidEv(void* qthis);
  // proto:  QOpenGLContext * QOpenGLWidget::context();
extern void _ZNK13QOpenGLWidget7contextEv(void* qthis);
  // proto:  void QOpenGLWidget::doneCurrent();
extern void _ZN13QOpenGLWidget11doneCurrentEv(void* qthis);
  // proto:  void QOpenGLWidget::makeCurrent();
extern void _ZN13QOpenGLWidget11makeCurrentEv(void* qthis);
  // proto:  QImage QOpenGLWidget::grabFramebuffer();
extern void _ZN13QOpenGLWidget15grabFramebufferEv(void* qthis);
  // proto:  const QMetaObject * QOpenGLWidget::metaObject();
extern void _ZNK13QOpenGLWidget10metaObjectEv(void* qthis);
  // proto:  void QOpenGLWidget::setFormat(const QSurfaceFormat & format);
extern void _ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat(void* qthis, void* arg0);
  // proto:  QSurfaceFormat QOpenGLWidget::format();
extern void _ZNK13QOpenGLWidget6formatEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
//  _aboutToResize QOpenGLWidget_aboutToResize_signal;
//  _resized QOpenGLWidget_resized_signal;
//  _frameSwapped QOpenGLWidget_frameSwapped_signal;
//  _aboutToCompose QOpenGLWidget_aboutToCompose_signal;
}

  // proto:  void QOpenGLWidget::~QOpenGLWidget();
func (this *QOpenGLWidget) FreeQOpenGLWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "~QOpenGLWidget", args)
  }

}

  // proto:  GLuint QOpenGLWidget::defaultFramebufferObject();
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
    C._ZNK13QOpenGLWidget24defaultFramebufferObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "defaultFramebufferObject", args)
  }

}

  // proto:  void QOpenGLWidget::QOpenGLWidget(const QOpenGLWidget & );
func NewQOpenGLWidget(args ...interface{}) QOpenGLWidget {
  return QOpenGLWidget{}
}

  // proto:  bool QOpenGLWidget::isValid();
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
    C._ZNK13QOpenGLWidget7isValidEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "isValid", args)
  }

}

  // proto:  QOpenGLContext * QOpenGLWidget::context();
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
    C._ZNK13QOpenGLWidget7contextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "context", args)
  }

}

  // proto:  void QOpenGLWidget::doneCurrent();
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
    C._ZN13QOpenGLWidget11doneCurrentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "doneCurrent", args)
  }

}

  // proto:  void QOpenGLWidget::makeCurrent();
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
    C._ZN13QOpenGLWidget11makeCurrentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "makeCurrent", args)
  }

}

  // proto:  QImage QOpenGLWidget::grabFramebuffer();
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
    C._ZN13QOpenGLWidget15grabFramebufferEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "grabFramebuffer", args)
  }

}

  // proto:  const QMetaObject * QOpenGLWidget::metaObject();
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
    C._ZNK13QOpenGLWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "metaObject", args)
  }

}

  // proto:  void QOpenGLWidget::setFormat(const QSurfaceFormat & format);
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
    C._ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "setFormat", args)
  }

}

  // proto:  QSurfaceFormat QOpenGLWidget::format();
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
    C._ZNK13QOpenGLWidget6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "format", args)
  }

}

// <= body block end

