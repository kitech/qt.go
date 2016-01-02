package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qoffscreensurface.h
// dst-file: /src/gui/qoffscreensurface.go
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
  // proto:  void QOffscreenSurface::~QOffscreenSurface();
extern void _ZN17QOffscreenSurfaceD0Ev(void* qthis);
  // proto:  void QOffscreenSurface::QOffscreenSurface(const QOffscreenSurface & );
extern void* dector_ZN17QOffscreenSurfaceC1ERKS_(void* arg0);
extern void _ZN17QOffscreenSurfaceC1ERKS_(void* qthis, void* arg0);
  // proto:  void QOffscreenSurface::QOffscreenSurface(QScreen * screen);
extern void* dector_ZN17QOffscreenSurfaceC1EP7QScreen(void* arg0);
extern void _ZN17QOffscreenSurfaceC1EP7QScreen(void* qthis, void* arg0);
  // proto:  QScreen * QOffscreenSurface::screen();
extern void _ZNK17QOffscreenSurface6screenEv(void* qthis);
  // proto:  void QOffscreenSurface::setFormat(const QSurfaceFormat & format);
extern void _ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat(void* qthis, void* arg0);
  // proto:  void QOffscreenSurface::setScreen(QScreen * screen);
extern void _ZN17QOffscreenSurface9setScreenEP7QScreen(void* qthis, void* arg0);
  // proto:  QSurfaceFormat QOffscreenSurface::requestedFormat();
extern void _ZNK17QOffscreenSurface15requestedFormatEv(void* qthis);
  // proto:  QSurfaceFormat QOffscreenSurface::format();
extern void _ZNK17QOffscreenSurface6formatEv(void* qthis);
  // proto:  QPlatformOffscreenSurface * QOffscreenSurface::handle();
extern void _ZNK17QOffscreenSurface6handleEv(void* qthis);
  // proto:  const QMetaObject * QOffscreenSurface::metaObject();
extern void _ZNK17QOffscreenSurface10metaObjectEv(void* qthis);
  // proto:  void QOffscreenSurface::destroy();
extern void _ZN17QOffscreenSurface7destroyEv(void* qthis);
  // proto:  bool QOffscreenSurface::isValid();
extern void _ZNK17QOffscreenSurface7isValidEv(void* qthis);
  // proto:  QSize QOffscreenSurface::size();
extern void _ZNK17QOffscreenSurface4sizeEv(void* qthis);
  // proto:  void QOffscreenSurface::create();
extern void _ZN17QOffscreenSurface6createEv(void* qthis);
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

// class sizeof(QOffscreenSurface)=1
type QOffscreenSurface struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _screenChanged QOffscreenSurface_screenChanged_signal;
}

  // proto:  void QOffscreenSurface::~QOffscreenSurface();
func (this *QOffscreenSurface) FreeQOffscreenSurface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "~QOffscreenSurface", args)
  }

}

  // proto:  void QOffscreenSurface::QOffscreenSurface(const QOffscreenSurface & );
func NewQOffscreenSurface(args ...interface{}) QOffscreenSurface {
  return QOffscreenSurface{}
}

  // proto:  QScreen * QOffscreenSurface::screen();
func (this *QOffscreenSurface) screen(args ...interface{}) () {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6screenEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "screen", args)
  }

}

  // proto:  void QOffscreenSurface::setFormat(const QSurfaceFormat & format);
func (this *QOffscreenSurface) setFormat(args ...interface{}) () {
  // setFormat(const class QSurfaceFormat &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSurfaceFormat{}) // "const QSurfaceFormat &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "setFormat", args)
  }

}

  // proto:  void QOffscreenSurface::setScreen(QScreen * screen);
func (this *QOffscreenSurface) setScreen(args ...interface{}) () {
  // setScreen(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface9setScreenEP7QScreen
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "setScreen", args)
  }

}

  // proto:  QSurfaceFormat QOffscreenSurface::requestedFormat();
func (this *QOffscreenSurface) requestedFormat(args ...interface{}) () {
  // requestedFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface15requestedFormatEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "requestedFormat", args)
  }

}

  // proto:  QSurfaceFormat QOffscreenSurface::format();
func (this *QOffscreenSurface) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6formatEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "format", args)
  }

}

  // proto:  QPlatformOffscreenSurface * QOffscreenSurface::handle();
func (this *QOffscreenSurface) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6handleEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "handle", args)
  }

}

  // proto:  const QMetaObject * QOffscreenSurface::metaObject();
func (this *QOffscreenSurface) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface10metaObjectEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "metaObject", args)
  }

}

  // proto:  void QOffscreenSurface::destroy();
func (this *QOffscreenSurface) destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface7destroyEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "destroy", args)
  }

}

  // proto:  bool QOffscreenSurface::isValid();
func (this *QOffscreenSurface) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface7isValidEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "isValid", args)
  }

}

  // proto:  QSize QOffscreenSurface::size();
func (this *QOffscreenSurface) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface4sizeEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "size", args)
  }

}

  // proto:  void QOffscreenSurface::create();
func (this *QOffscreenSurface) create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface6createEv
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "create", args)
  }

}

// <= body block end

