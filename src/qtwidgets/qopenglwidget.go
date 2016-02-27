package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
  // proto:  QOpenGLWidget::UpdateBehavior QOpenGLWidget::updateBehavior();
extern void C_ZNK13QOpenGLWidget14updateBehaviorEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::doneCurrent();
extern void C_ZN13QOpenGLWidget11doneCurrentEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::setFormat(const QSurfaceFormat & format);
extern void C_ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat(void* qthis, void* arg0); // 4
  // proto:  QSurfaceFormat QOpenGLWidget::format();
extern void* C_ZNK13QOpenGLWidget6formatEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::~QOpenGLWidget();
extern void C_ZN13QOpenGLWidgetD2Ev(void* qthis); // 4
  // proto:  bool QOpenGLWidget::isValid();
extern bool C_ZNK13QOpenGLWidget7isValidEv(void* qthis); // 4
  // proto:  GLuint QOpenGLWidget::defaultFramebufferObject();
extern int32_t C_ZNK13QOpenGLWidget24defaultFramebufferObjectEv(void* qthis); // 4
  // proto:  const QMetaObject * QOpenGLWidget::metaObject();
extern void C_ZNK13QOpenGLWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QOpenGLWidget::makeCurrent();
extern void C_ZN13QOpenGLWidget11makeCurrentEv(void* qthis); // 4
  // proto:  QImage QOpenGLWidget::grabFramebuffer();
extern void* C_ZN13QOpenGLWidget15grabFramebufferEv(void* qthis); // 4
  // proto:  QOpenGLContext * QOpenGLWidget::context();
extern void C_ZNK13QOpenGLWidget7contextEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QOpenGLWidget)=1
type QOpenGLWidget struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _aboutToResize QOpenGLWidget_aboutToResize_signal;
//  _resized QOpenGLWidget_resized_signal;
//  _frameSwapped QOpenGLWidget_frameSwapped_signal;
//  _aboutToCompose QOpenGLWidget_aboutToCompose_signal;
}

// updateBehavior()
func (this *QOpenGLWidget) UpdateBehavior(args ...interface{}) () {
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
    C.C_ZNK13QOpenGLWidget14updateBehaviorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "updateBehavior", args)
  }

  return
}

// doneCurrent()
func (this *QOpenGLWidget) DoneCurrent(args ...interface{}) () {
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
    C.C_ZN13QOpenGLWidget11doneCurrentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "doneCurrent", args)
  }

  return
}

// setFormat(const class QSurfaceFormat &)
func (this *QOpenGLWidget) SetFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat
    // invoke: void setFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(*qtgui.QSurfaceFormat).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QOpenGLWidget9setFormatERK14QSurfaceFormat(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "setFormat", args)
  }

  return
}

// format()
func (this *QOpenGLWidget) Format(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QOpenGLWidget6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QSurfaceFormat{}) // "QSurfaceFormat"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "format", args)
  }

  return
}

// ~QOpenGLWidget()
func (this *QOpenGLWidget) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QOpenGLWidgetD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "~QOpenGLWidget", args)
  }

  return
}

// isValid()
func (this *QOpenGLWidget) IsValid(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QOpenGLWidget7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "isValid", args)
  }

  return
}

// defaultFramebufferObject()
func (this *QOpenGLWidget) DefaultFramebufferObject(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QOpenGLWidget24defaultFramebufferObjectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "GLuint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "defaultFramebufferObject", args)
  }

  return
}

// metaObject()
func (this *QOpenGLWidget) MetaObject(args ...interface{}) () {
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
    C.C_ZNK13QOpenGLWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "metaObject", args)
  }

  return
}

// makeCurrent()
func (this *QOpenGLWidget) MakeCurrent(args ...interface{}) () {
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
    C.C_ZN13QOpenGLWidget11makeCurrentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "makeCurrent", args)
  }

  return
}

// grabFramebuffer()
func (this *QOpenGLWidget) GrabFramebuffer(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QOpenGLWidget15grabFramebufferEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QImage{}) // "QImage"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "grabFramebuffer", args)
  }

  return
}

// context()
func (this *QOpenGLWidget) Context(args ...interface{}) () {
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
    C.C_ZNK13QOpenGLWidget7contextEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QOpenGLWidget", "context", args)
  }

  return
}

// <= body block end

