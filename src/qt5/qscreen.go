package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QScreen::refreshRate();
extern void C_ZNK7QScreen11refreshRateEv(void* qthis); // 4
  // proto:  QSize QScreen::availableVirtualSize();
extern void C_ZNK7QScreen20availableVirtualSizeEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QScreen::orientation();
extern void C_ZNK7QScreen11orientationEv(void* qthis); // 4
  // proto:  QPixmap QScreen::grabWindow(WId window, int x, int y, int w, int h);
extern void C_ZN7QScreen10grabWindowEiiiii(void* qthis, int32_t* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  Qt::ScreenOrientation QScreen::nativeOrientation();
extern void C_ZNK7QScreen17nativeOrientationEv(void* qthis); // 4
  // proto:  QSize QScreen::size();
extern void C_ZNK7QScreen4sizeEv(void* qthis); // 4
  // proto:  qreal QScreen::physicalDotsPerInchY();
extern void C_ZNK7QScreen20physicalDotsPerInchYEv(void* qthis); // 4
  // proto:  qreal QScreen::logicalDotsPerInch();
extern void C_ZNK7QScreen18logicalDotsPerInchEv(void* qthis); // 4
  // proto:  QList<QScreen *> QScreen::virtualSiblings();
extern void C_ZNK7QScreen15virtualSiblingsEv(void* qthis); // 4
  // proto:  QRect QScreen::availableGeometry();
extern void C_ZNK7QScreen17availableGeometryEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientations QScreen::orientationUpdateMask();
extern void C_ZNK7QScreen21orientationUpdateMaskEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QScreen::primaryOrientation();
extern void C_ZNK7QScreen18primaryOrientationEv(void* qthis); // 4
  // proto:  QPlatformScreen * QScreen::handle();
extern void C_ZNK7QScreen6handleEv(void* qthis); // 4
  // proto:  QRect QScreen::availableVirtualGeometry();
extern void C_ZNK7QScreen24availableVirtualGeometryEv(void* qthis); // 4
  // proto:  void QScreen::~QScreen();
extern void C_ZN7QScreenD2Ev(void* qthis); // 4
  // proto:  QSize QScreen::availableSize();
extern void C_ZNK7QScreen13availableSizeEv(void* qthis); // 4
  // proto:  QRect QScreen::virtualGeometry();
extern void C_ZNK7QScreen15virtualGeometryEv(void* qthis); // 4
  // proto:  qreal QScreen::physicalDotsPerInch();
extern void C_ZNK7QScreen19physicalDotsPerInchEv(void* qthis); // 4
  // proto:  QSizeF QScreen::physicalSize();
extern void C_ZNK7QScreen12physicalSizeEv(void* qthis); // 4
  // proto:  const QMetaObject * QScreen::metaObject();
extern void C_ZNK7QScreen10metaObjectEv(void* qthis); // 4
  // proto:  QString QScreen::name();
extern void C_ZNK7QScreen4nameEv(void* qthis); // 4
  // proto:  qreal QScreen::logicalDotsPerInchY();
extern void C_ZNK7QScreen19logicalDotsPerInchYEv(void* qthis); // 4
  // proto:  qreal QScreen::logicalDotsPerInchX();
extern void C_ZNK7QScreen19logicalDotsPerInchXEv(void* qthis); // 4
  // proto:  qreal QScreen::devicePixelRatio();
extern void C_ZNK7QScreen16devicePixelRatioEv(void* qthis); // 4
  // proto:  QRect QScreen::geometry();
extern void C_ZNK7QScreen8geometryEv(void* qthis); // 4
  // proto:  int QScreen::depth();
extern void C_ZNK7QScreen5depthEv(void* qthis); // 4
  // proto:  QSize QScreen::virtualSize();
extern void C_ZNK7QScreen11virtualSizeEv(void* qthis); // 4
  // proto:  qreal QScreen::physicalDotsPerInchX();
extern void C_ZNK7QScreen20physicalDotsPerInchXEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
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

// refreshRate()
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
    // invoke: qreal refreshRate()
    C.C_ZNK7QScreen11refreshRateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "refreshRate", args)
  }

}

// availableVirtualSize()
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
    // invoke: QSize availableVirtualSize()
    C.C_ZNK7QScreen20availableVirtualSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualSize", args)
  }

}

// orientation()
func (this *QScreen) orientation(args ...interface{}) () {
  // orientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen11orientationEv
    // invoke: Qt::ScreenOrientation orientation()
    C.C_ZNK7QScreen11orientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "orientation", args)
  }

}

// grabWindow(WId, int, int, int, int)
func (this *QScreen) grabWindow(args ...interface{}) () {
  // grabWindow(WId, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "WId"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QScreen10grabWindowEiiiii
    // invoke: QPixmap grabWindow(WId, int, int, int, int)
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
    C.C_ZN7QScreen10grabWindowEiiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QScreen", "grabWindow", args)
  }

}

// nativeOrientation()
func (this *QScreen) nativeOrientation(args ...interface{}) () {
  // nativeOrientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen17nativeOrientationEv
    // invoke: Qt::ScreenOrientation nativeOrientation()
    C.C_ZNK7QScreen17nativeOrientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "nativeOrientation", args)
  }

}

// size()
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
    // invoke: QSize size()
    C.C_ZNK7QScreen4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "size", args)
  }

}

// physicalDotsPerInchY()
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
    // invoke: qreal physicalDotsPerInchY()
    C.C_ZNK7QScreen20physicalDotsPerInchYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchY", args)
  }

}

// logicalDotsPerInch()
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
    // invoke: qreal logicalDotsPerInch()
    C.C_ZNK7QScreen18logicalDotsPerInchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInch", args)
  }

}

// virtualSiblings()
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
    // invoke: QList<QScreen *> virtualSiblings()
    C.C_ZNK7QScreen15virtualSiblingsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "virtualSiblings", args)
  }

}

// availableGeometry()
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
    // invoke: QRect availableGeometry()
    C.C_ZNK7QScreen17availableGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "availableGeometry", args)
  }

}

// orientationUpdateMask()
func (this *QScreen) orientationUpdateMask(args ...interface{}) () {
  // orientationUpdateMask()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen21orientationUpdateMaskEv
    // invoke: Qt::ScreenOrientations orientationUpdateMask()
    C.C_ZNK7QScreen21orientationUpdateMaskEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "orientationUpdateMask", args)
  }

}

// primaryOrientation()
func (this *QScreen) primaryOrientation(args ...interface{}) () {
  // primaryOrientation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QScreen18primaryOrientationEv
    // invoke: Qt::ScreenOrientation primaryOrientation()
    C.C_ZNK7QScreen18primaryOrientationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "primaryOrientation", args)
  }

}

// handle()
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
    // invoke: QPlatformScreen * handle()
    C.C_ZNK7QScreen6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "handle", args)
  }

}

// availableVirtualGeometry()
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
    // invoke: QRect availableVirtualGeometry()
    C.C_ZNK7QScreen24availableVirtualGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualGeometry", args)
  }

}

// ~QScreen()
func (this *QScreen) FreeQScreen(args ...interface{}) () {
  // ~QScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QScreenD0Ev
    // invoke: void ~QScreen()
    C.C_ZN7QScreenD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "~QScreen", args)
  }

}

// availableSize()
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
    // invoke: QSize availableSize()
    C.C_ZNK7QScreen13availableSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "availableSize", args)
  }

}

// virtualGeometry()
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
    // invoke: QRect virtualGeometry()
    C.C_ZNK7QScreen15virtualGeometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "virtualGeometry", args)
  }

}

// physicalDotsPerInch()
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
    // invoke: qreal physicalDotsPerInch()
    C.C_ZNK7QScreen19physicalDotsPerInchEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInch", args)
  }

}

// physicalSize()
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
    // invoke: QSizeF physicalSize()
    C.C_ZNK7QScreen12physicalSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "physicalSize", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK7QScreen10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "metaObject", args)
  }

}

// name()
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
    // invoke: QString name()
    C.C_ZNK7QScreen4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "name", args)
  }

}

// logicalDotsPerInchY()
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
    // invoke: qreal logicalDotsPerInchY()
    C.C_ZNK7QScreen19logicalDotsPerInchYEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchY", args)
  }

}

// logicalDotsPerInchX()
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
    // invoke: qreal logicalDotsPerInchX()
    C.C_ZNK7QScreen19logicalDotsPerInchXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchX", args)
  }

}

// devicePixelRatio()
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
    // invoke: qreal devicePixelRatio()
    C.C_ZNK7QScreen16devicePixelRatioEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "devicePixelRatio", args)
  }

}

// geometry()
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
    // invoke: QRect geometry()
    C.C_ZNK7QScreen8geometryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "geometry", args)
  }

}

// depth()
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
    // invoke: int depth()
    C.C_ZNK7QScreen5depthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "depth", args)
  }

}

// virtualSize()
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
    // invoke: QSize virtualSize()
    C.C_ZNK7QScreen11virtualSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "virtualSize", args)
  }

}

// physicalDotsPerInchX()
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
    // invoke: qreal physicalDotsPerInchX()
    C.C_ZNK7QScreen20physicalDotsPerInchXEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchX", args)
  }

}

// <= body block end

