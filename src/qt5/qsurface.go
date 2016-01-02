package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qsurface.h
// dst-file: /src/gui/qsurface.go
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
  // proto:  bool QSurface::supportsOpenGL();
extern void _ZNK8QSurface14supportsOpenGLEv(void* qthis);
  // proto:  QSurfaceFormat QSurface::format();
extern void _ZNK8QSurface6formatEv(void* qthis);
  // proto:  QPlatformSurface * QSurface::surfaceHandle();
extern void _ZNK8QSurface13surfaceHandleEv(void* qthis);
  // proto:  QSize QSurface::size();
extern void _ZNK8QSurface4sizeEv(void* qthis);
  // proto:  void QSurface::~QSurface();
extern void _ZN8QSurfaceD0Ev(void* qthis);
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

// class sizeof(QSurface)=24
type QSurface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  bool QSurface::supportsOpenGL();
func (this *QSurface) supportsOpenGL(args ...interface{}) () {
  // supportsOpenGL()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSurface14supportsOpenGLEv
    // invoke: bool supportsOpenGL()
    C._ZNK8QSurface14supportsOpenGLEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurface", "supportsOpenGL", args)
  }

}

  // proto:  QSurfaceFormat QSurface::format();
func (this *QSurface) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSurface6formatEv
    // invoke: QSurfaceFormat format()
    C._ZNK8QSurface6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurface", "format", args)
  }

}

  // proto:  QPlatformSurface * QSurface::surfaceHandle();
func (this *QSurface) surfaceHandle(args ...interface{}) () {
  // surfaceHandle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSurface13surfaceHandleEv
    // invoke: QPlatformSurface * surfaceHandle()
    C._ZNK8QSurface13surfaceHandleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurface", "surfaceHandle", args)
  }

}

  // proto:  QSize QSurface::size();
func (this *QSurface) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QSurface4sizeEv
    // invoke: QSize size()
    C._ZNK8QSurface4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSurface", "size", args)
  }

}

  // proto:  void QSurface::~QSurface();
func (this *QSurface) FreeQSurface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSurface", "~QSurface", args)
  }

}

// <= body block end

