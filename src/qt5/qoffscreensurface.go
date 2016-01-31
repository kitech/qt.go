package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QOffscreenSurface::QOffscreenSurface(QScreen * screen);
extern void* C_ZN17QOffscreenSurfaceC2EP7QScreen(void* arg0); // 3
  // proto:  QSurface::SurfaceType QOffscreenSurface::surfaceType();
extern void C_ZNK17QOffscreenSurface11surfaceTypeEv(void* qthis); // 4
  // proto:  QSize QOffscreenSurface::size();
extern void* C_ZNK17QOffscreenSurface4sizeEv(void* qthis); // 4
  // proto:  void QOffscreenSurface::setScreen(QScreen * screen);
extern void C_ZN17QOffscreenSurface9setScreenEP7QScreen(void* qthis, void* arg0); // 4
  // proto:  void QOffscreenSurface::setFormat(const QSurfaceFormat & format);
extern void C_ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat(void* qthis, void* arg0); // 4
  // proto:  void QOffscreenSurface::create();
extern void C_ZN17QOffscreenSurface6createEv(void* qthis); // 4
  // proto:  QSurfaceFormat QOffscreenSurface::format();
extern void* C_ZNK17QOffscreenSurface6formatEv(void* qthis); // 4
  // proto:  void QOffscreenSurface::destroy();
extern void C_ZN17QOffscreenSurface7destroyEv(void* qthis); // 4
  // proto:  QPlatformOffscreenSurface * QOffscreenSurface::handle();
extern void C_ZNK17QOffscreenSurface6handleEv(void* qthis); // 4
  // proto:  bool QOffscreenSurface::isValid();
extern bool C_ZNK17QOffscreenSurface7isValidEv(void* qthis); // 4
  // proto:  QScreen * QOffscreenSurface::screen();
extern void* C_ZNK17QOffscreenSurface6screenEv(void* qthis); // 4
  // proto:  void QOffscreenSurface::~QOffscreenSurface();
extern void C_ZN17QOffscreenSurfaceD2Ev(void* qthis); // 4
  // proto:  const QMetaObject * QOffscreenSurface::metaObject();
extern void C_ZNK17QOffscreenSurface10metaObjectEv(void* qthis); // 4
  // proto:  QSurfaceFormat QOffscreenSurface::requestedFormat();
extern void* C_ZNK17QOffscreenSurface15requestedFormatEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _screenChanged QOffscreenSurface_screenChanged_signal;
}

// QOffscreenSurface(class QScreen *)
func NewQOffscreenSurface(args ...interface{}) *QOffscreenSurface {
  // QOffscreenSurface(class QScreen *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScreen{}) // "QScreen *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurfaceC1EP7QScreen
    // invoke: void QOffscreenSurface(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN17QOffscreenSurfaceC2EP7QScreen(arg0)
    return &QOffscreenSurface{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "QOffscreenSurface", args)
  }

  return nil // QOffscreenSurface{qclsinst:qthis}
}

// surfaceType()
func (this *QOffscreenSurface) Surfacetype(args ...interface{}) () {
  // surfaceType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface11surfaceTypeEv
    // invoke: QSurface::SurfaceType surfaceType()
    C.C_ZNK17QOffscreenSurface11surfaceTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "surfaceType", args)
  }

  return
}

// size()
func (this *QOffscreenSurface) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface4sizeEv
    // invoke: QSize size()
    var ret0 = C.C_ZNK17QOffscreenSurface4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "size", args)
  }

  return
}

// setScreen(class QScreen *)
func (this *QOffscreenSurface) Setscreen(args ...interface{}) () {
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
    // invoke: void setScreen(class QScreen *)
    var arg0 = args[0].(QScreen).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QOffscreenSurface9setScreenEP7QScreen(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "setScreen", args)
  }

  return
}

// setFormat(const class QSurfaceFormat &)
func (this *QOffscreenSurface) Setformat(args ...interface{}) () {
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
    // invoke: void setFormat(const class QSurfaceFormat &)
    var arg0 = args[0].(QSurfaceFormat).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN17QOffscreenSurface9setFormatERK14QSurfaceFormat(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "setFormat", args)
  }

  return
}

// create()
func (this *QOffscreenSurface) Create(args ...interface{}) () {
  // create()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface6createEv
    // invoke: void create()
    C.C_ZN17QOffscreenSurface6createEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "create", args)
  }

  return
}

// format()
func (this *QOffscreenSurface) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6formatEv
    // invoke: QSurfaceFormat format()
    var ret0 = C.C_ZNK17QOffscreenSurface6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurfaceFormat{}) // "QSurfaceFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "format", args)
  }

  return
}

// destroy()
func (this *QOffscreenSurface) Destroy(args ...interface{}) () {
  // destroy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurface7destroyEv
    // invoke: void destroy()
    C.C_ZN17QOffscreenSurface7destroyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "destroy", args)
  }

  return
}

// handle()
func (this *QOffscreenSurface) Handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6handleEv
    // invoke: QPlatformOffscreenSurface * handle()
    C.C_ZNK17QOffscreenSurface6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "handle", args)
  }

  return
}

// isValid()
func (this *QOffscreenSurface) Isvalid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK17QOffscreenSurface7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "isValid", args)
  }

  return
}

// screen()
func (this *QOffscreenSurface) Screen(args ...interface{}) (ret interface{}) {
  // screen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface6screenEv
    // invoke: QScreen * screen()
    var ret0 = C.C_ZNK17QOffscreenSurface6screenEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScreen{}) // "QScreen *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "screen", args)
  }

  return
}

// ~QOffscreenSurface()
func (this *QOffscreenSurface) Freeqoffscreensurface(args ...interface{}) () {
  // ~QOffscreenSurface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QOffscreenSurfaceD0Ev
    // invoke: void ~QOffscreenSurface()
    C.C_ZN17QOffscreenSurfaceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "~QOffscreenSurface", args)
  }

  return
}

// metaObject()
func (this *QOffscreenSurface) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK17QOffscreenSurface10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "metaObject", args)
  }

  return
}

// requestedFormat()
func (this *QOffscreenSurface) Requestedformat(args ...interface{}) (ret interface{}) {
  // requestedFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QOffscreenSurface15requestedFormatEv
    // invoke: QSurfaceFormat requestedFormat()
    var ret0 = C.C_ZNK17QOffscreenSurface15requestedFormatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSurfaceFormat{}) // "QSurfaceFormat"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QOffscreenSurface", "requestedFormat", args)
  }

  return
}

// <= body block end

