package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qscreen.h
// dst-file: /src/gui/qscreen.go
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
  // proto:  qreal QScreen::logicalDotsPerInchY();
extern void _ZNK7QScreen19logicalDotsPerInchYEv(void* qthis);
  // proto:  QRect QScreen::geometry();
extern void _ZNK7QScreen8geometryEv(void* qthis);
  // proto:  QPixmap QScreen::grabWindow(WId window, int x, int y, int w, int h);
extern void _ZN7QScreen10grabWindowEiiiii(void* qthis, uint32_t* arg0, int arg1, int arg2, int arg3, int arg4);
  // proto:  QSize QScreen::size();
extern void _ZNK7QScreen4sizeEv(void* qthis);
  // proto:  QSizeF QScreen::physicalSize();
extern void _ZNK7QScreen12physicalSizeEv(void* qthis);
  // proto:  QPlatformScreen * QScreen::handle();
extern void _ZNK7QScreen6handleEv(void* qthis);
  // proto:  QRect QScreen::availableVirtualGeometry();
extern void _ZNK7QScreen24availableVirtualGeometryEv(void* qthis);
  // proto:  void QScreen::~QScreen();
extern void _ZN7QScreenD0Ev(void* qthis);
  // proto:  QSize QScreen::virtualSize();
extern void _ZNK7QScreen11virtualSizeEv(void* qthis);
  // proto:  qreal QScreen::devicePixelRatio();
extern void _ZNK7QScreen16devicePixelRatioEv(void* qthis);
  // proto:  QList<QScreen *> QScreen::virtualSiblings();
extern void _ZNK7QScreen15virtualSiblingsEv(void* qthis);
  // proto:  void QScreen::QScreen(const QScreen & );
extern void* dector_ZN7QScreenC1ERKS_(void* arg0);
extern void _ZN7QScreenC1ERKS_(void* qthis, void* arg0);
  // proto:  QRect QScreen::virtualGeometry();
extern void _ZNK7QScreen15virtualGeometryEv(void* qthis);
  // proto:  qreal QScreen::logicalDotsPerInch();
extern void _ZNK7QScreen18logicalDotsPerInchEv(void* qthis);
  // proto:  qreal QScreen::physicalDotsPerInch();
extern void _ZNK7QScreen19physicalDotsPerInchEv(void* qthis);
  // proto:  qreal QScreen::refreshRate();
extern void _ZNK7QScreen11refreshRateEv(void* qthis);
  // proto:  const QMetaObject * QScreen::metaObject();
extern void _ZNK7QScreen10metaObjectEv(void* qthis);
  // proto:  QSize QScreen::availableSize();
extern void _ZNK7QScreen13availableSizeEv(void* qthis);
  // proto:  QString QScreen::name();
extern void _ZNK7QScreen4nameEv(void* qthis);
  // proto:  QSize QScreen::availableVirtualSize();
extern void _ZNK7QScreen20availableVirtualSizeEv(void* qthis);
  // proto:  qreal QScreen::logicalDotsPerInchX();
extern void _ZNK7QScreen19logicalDotsPerInchXEv(void* qthis);
  // proto:  QRect QScreen::availableGeometry();
extern void _ZNK7QScreen17availableGeometryEv(void* qthis);
  // proto:  qreal QScreen::physicalDotsPerInchX();
extern void _ZNK7QScreen20physicalDotsPerInchXEv(void* qthis);
  // proto:  qreal QScreen::physicalDotsPerInchY();
extern void _ZNK7QScreen20physicalDotsPerInchYEv(void* qthis);
  // proto:  int QScreen::depth();
extern void _ZNK7QScreen5depthEv(void* qthis);
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

// class sizeof(QScreen)=1
type QScreen struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _geometryChanged QScreen_geometryChanged_signal;
//  _virtualGeometryChanged QScreen_virtualGeometryChanged_signal;
//  _refreshRateChanged QScreen_refreshRateChanged_signal;
//  _availableGeometryChanged QScreen_availableGeometryChanged_signal;
//  _physicalSizeChanged QScreen_physicalSizeChanged_signal;
//  _physicalDotsPerInchChanged QScreen_physicalDotsPerInchChanged_signal;
//  _logicalDotsPerInchChanged QScreen_logicalDotsPerInchChanged_signal;
//  _primaryOrientationChanged QScreen_primaryOrientationChanged_signal;
//  _orientationChanged QScreen_orientationChanged_signal;
}

  // proto:  qreal QScreen::logicalDotsPerInchY();
func (this *QScreen) logicalDotsPerInchY(args ...interface{}) () {
  // logicalDotsPerInchY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen19logicalDotsPerInchYEv
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchY", args)
  }

}

  // proto:  QRect QScreen::geometry();
func (this *QScreen) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen8geometryEv
  default:
    qtrt.ErrorResolve("QScreen", "geometry", args)
  }

}

  // proto:  QPixmap QScreen::grabWindow(WId window, int x, int y, int w, int h);
func (this *QScreen) grabWindow(args ...interface{}) () {
  // grabWindow(WId, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "WId"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QScreen10grabWindowEiiiii
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
  default:
    qtrt.ErrorResolve("QScreen", "grabWindow", args)
  }

}

  // proto:  QSize QScreen::size();
func (this *QScreen) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen4sizeEv
  default:
    qtrt.ErrorResolve("QScreen", "size", args)
  }

}

  // proto:  QSizeF QScreen::physicalSize();
func (this *QScreen) physicalSize(args ...interface{}) () {
  // physicalSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen12physicalSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalSize", args)
  }

}

  // proto:  QPlatformScreen * QScreen::handle();
func (this *QScreen) handle(args ...interface{}) () {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen6handleEv
  default:
    qtrt.ErrorResolve("QScreen", "handle", args)
  }

}

  // proto:  QRect QScreen::availableVirtualGeometry();
func (this *QScreen) availableVirtualGeometry(args ...interface{}) () {
  // availableVirtualGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen24availableVirtualGeometryEv
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualGeometry", args)
  }

}

  // proto:  void QScreen::~QScreen();
func (this *QScreen) FreeQScreen(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScreen", "~QScreen", args)
  }

}

  // proto:  QSize QScreen::virtualSize();
func (this *QScreen) virtualSize(args ...interface{}) () {
  // virtualSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen11virtualSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "virtualSize", args)
  }

}

  // proto:  qreal QScreen::devicePixelRatio();
func (this *QScreen) devicePixelRatio(args ...interface{}) () {
  // devicePixelRatio()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen16devicePixelRatioEv
  default:
    qtrt.ErrorResolve("QScreen", "devicePixelRatio", args)
  }

}

  // proto:  QList<QScreen *> QScreen::virtualSiblings();
func (this *QScreen) virtualSiblings(args ...interface{}) () {
  // virtualSiblings()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen15virtualSiblingsEv
  default:
    qtrt.ErrorResolve("QScreen", "virtualSiblings", args)
  }

}

  // proto:  void QScreen::QScreen(const QScreen & );
func NewQScreen(args ...interface{}) QScreen {
  return QScreen{}
}

  // proto:  QRect QScreen::virtualGeometry();
func (this *QScreen) virtualGeometry(args ...interface{}) () {
  // virtualGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen15virtualGeometryEv
  default:
    qtrt.ErrorResolve("QScreen", "virtualGeometry", args)
  }

}

  // proto:  qreal QScreen::logicalDotsPerInch();
func (this *QScreen) logicalDotsPerInch(args ...interface{}) () {
  // logicalDotsPerInch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen18logicalDotsPerInchEv
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInch", args)
  }

}

  // proto:  qreal QScreen::physicalDotsPerInch();
func (this *QScreen) physicalDotsPerInch(args ...interface{}) () {
  // physicalDotsPerInch()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen19physicalDotsPerInchEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInch", args)
  }

}

  // proto:  qreal QScreen::refreshRate();
func (this *QScreen) refreshRate(args ...interface{}) () {
  // refreshRate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen11refreshRateEv
  default:
    qtrt.ErrorResolve("QScreen", "refreshRate", args)
  }

}

  // proto:  const QMetaObject * QScreen::metaObject();
func (this *QScreen) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen10metaObjectEv
  default:
    qtrt.ErrorResolve("QScreen", "metaObject", args)
  }

}

  // proto:  QSize QScreen::availableSize();
func (this *QScreen) availableSize(args ...interface{}) () {
  // availableSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen13availableSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "availableSize", args)
  }

}

  // proto:  QString QScreen::name();
func (this *QScreen) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen4nameEv
  default:
    qtrt.ErrorResolve("QScreen", "name", args)
  }

}

  // proto:  QSize QScreen::availableVirtualSize();
func (this *QScreen) availableVirtualSize(args ...interface{}) () {
  // availableVirtualSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen20availableVirtualSizeEv
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualSize", args)
  }

}

  // proto:  qreal QScreen::logicalDotsPerInchX();
func (this *QScreen) logicalDotsPerInchX(args ...interface{}) () {
  // logicalDotsPerInchX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen19logicalDotsPerInchXEv
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchX", args)
  }

}

  // proto:  QRect QScreen::availableGeometry();
func (this *QScreen) availableGeometry(args ...interface{}) () {
  // availableGeometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen17availableGeometryEv
  default:
    qtrt.ErrorResolve("QScreen", "availableGeometry", args)
  }

}

  // proto:  qreal QScreen::physicalDotsPerInchX();
func (this *QScreen) physicalDotsPerInchX(args ...interface{}) () {
  // physicalDotsPerInchX()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen20physicalDotsPerInchXEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchX", args)
  }

}

  // proto:  qreal QScreen::physicalDotsPerInchY();
func (this *QScreen) physicalDotsPerInchY(args ...interface{}) () {
  // physicalDotsPerInchY()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen20physicalDotsPerInchYEv
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchY", args)
  }

}

  // proto:  int QScreen::depth();
func (this *QScreen) depth(args ...interface{}) () {
  // depth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen5depthEv
  default:
    qtrt.ErrorResolve("QScreen", "depth", args)
  }

}

// <= body block end

