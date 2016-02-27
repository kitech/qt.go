package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  qreal QScreen::refreshRate();
extern double C_ZNK7QScreen11refreshRateEv(void* qthis); // 4
  // proto:  QSize QScreen::availableVirtualSize();
extern void* C_ZNK7QScreen20availableVirtualSizeEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QScreen::orientation();
extern void C_ZNK7QScreen11orientationEv(void* qthis); // 4
  // proto:  QPixmap QScreen::grabWindow(WId window, int x, int y, int w, int h);
extern void* C_ZN7QScreen10grabWindowEiiiii(void* qthis, void* arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  Qt::ScreenOrientation QScreen::nativeOrientation();
extern void C_ZNK7QScreen17nativeOrientationEv(void* qthis); // 4
  // proto:  QSize QScreen::size();
extern void* C_ZNK7QScreen4sizeEv(void* qthis); // 4
  // proto:  qreal QScreen::physicalDotsPerInchY();
extern double C_ZNK7QScreen20physicalDotsPerInchYEv(void* qthis); // 4
  // proto:  qreal QScreen::logicalDotsPerInch();
extern double C_ZNK7QScreen18logicalDotsPerInchEv(void* qthis); // 4
  // proto:  QList<QScreen *> QScreen::virtualSiblings();
extern void C_ZNK7QScreen15virtualSiblingsEv(void* qthis); // 4
  // proto:  QRect QScreen::availableGeometry();
extern void* C_ZNK7QScreen17availableGeometryEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientations QScreen::orientationUpdateMask();
extern void C_ZNK7QScreen21orientationUpdateMaskEv(void* qthis); // 4
  // proto:  Qt::ScreenOrientation QScreen::primaryOrientation();
extern void C_ZNK7QScreen18primaryOrientationEv(void* qthis); // 4
  // proto:  QPlatformScreen * QScreen::handle();
extern void C_ZNK7QScreen6handleEv(void* qthis); // 4
  // proto:  QRect QScreen::availableVirtualGeometry();
extern void* C_ZNK7QScreen24availableVirtualGeometryEv(void* qthis); // 4
  // proto:  void QScreen::~QScreen();
extern void C_ZN7QScreenD2Ev(void* qthis); // 4
  // proto:  QSize QScreen::availableSize();
extern void* C_ZNK7QScreen13availableSizeEv(void* qthis); // 4
  // proto:  QRect QScreen::virtualGeometry();
extern void* C_ZNK7QScreen15virtualGeometryEv(void* qthis); // 4
  // proto:  qreal QScreen::physicalDotsPerInch();
extern double C_ZNK7QScreen19physicalDotsPerInchEv(void* qthis); // 4
  // proto:  QSizeF QScreen::physicalSize();
extern void* C_ZNK7QScreen12physicalSizeEv(void* qthis); // 4
  // proto:  const QMetaObject * QScreen::metaObject();
extern void C_ZNK7QScreen10metaObjectEv(void* qthis); // 4
  // proto:  QString QScreen::name();
extern void* C_ZNK7QScreen4nameEv(void* qthis); // 4
  // proto:  qreal QScreen::logicalDotsPerInchY();
extern double C_ZNK7QScreen19logicalDotsPerInchYEv(void* qthis); // 4
  // proto:  qreal QScreen::logicalDotsPerInchX();
extern double C_ZNK7QScreen19logicalDotsPerInchXEv(void* qthis); // 4
  // proto:  qreal QScreen::devicePixelRatio();
extern double C_ZNK7QScreen16devicePixelRatioEv(void* qthis); // 4
  // proto:  QRect QScreen::geometry();
extern void* C_ZNK7QScreen8geometryEv(void* qthis); // 4
  // proto:  int QScreen::depth();
extern int32_t C_ZNK7QScreen5depthEv(void* qthis); // 4
  // proto:  QSize QScreen::virtualSize();
extern void* C_ZNK7QScreen11virtualSizeEv(void* qthis); // 4
  // proto:  qreal QScreen::physicalDotsPerInchX();
extern double C_ZNK7QScreen20physicalDotsPerInchXEv(void* qthis); // 4
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

// class sizeof(QScreen)=1
type QScreen struct {
  /*qbase*/ qtcore.QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
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
func (this *QScreen) RefreshRate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen11refreshRateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "refreshRate", args)
  }

  return
}

// availableVirtualSize()
func (this *QScreen) AvailableVirtualSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen20availableVirtualSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualSize", args)
  }

  return
}

// orientation()
func (this *QScreen) Orientation(args ...interface{}) () {
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
    C.C_ZNK7QScreen11orientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "orientation", args)
  }

  return
}

// grabWindow(WId, int, int, int, int)
func (this *QScreen) GrabWindow(args ...interface{}) (ret interface{}) {
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
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN7QScreen10grabWindowEiiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPixmap{}) // "QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "grabWindow", args)
  }

  return
}

// nativeOrientation()
func (this *QScreen) NativeOrientation(args ...interface{}) () {
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
    C.C_ZNK7QScreen17nativeOrientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "nativeOrientation", args)
  }

  return
}

// size()
func (this *QScreen) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "size", args)
  }

  return
}

// physicalDotsPerInchY()
func (this *QScreen) PhysicalDotsPerInchY(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen20physicalDotsPerInchYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchY", args)
  }

  return
}

// logicalDotsPerInch()
func (this *QScreen) LogicalDotsPerInch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen18logicalDotsPerInchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInch", args)
  }

  return
}

// virtualSiblings()
func (this *QScreen) VirtualSiblings(args ...interface{}) () {
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
    C.C_ZNK7QScreen15virtualSiblingsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "virtualSiblings", args)
  }

  return
}

// availableGeometry()
func (this *QScreen) AvailableGeometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen17availableGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "availableGeometry", args)
  }

  return
}

// orientationUpdateMask()
func (this *QScreen) OrientationUpdateMask(args ...interface{}) () {
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
    C.C_ZNK7QScreen21orientationUpdateMaskEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "orientationUpdateMask", args)
  }

  return
}

// primaryOrientation()
func (this *QScreen) PrimaryOrientation(args ...interface{}) () {
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
    C.C_ZNK7QScreen18primaryOrientationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "primaryOrientation", args)
  }

  return
}

// handle()
func (this *QScreen) Handle(args ...interface{}) () {
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
    C.C_ZNK7QScreen6handleEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "handle", args)
  }

  return
}

// availableVirtualGeometry()
func (this *QScreen) AvailableVirtualGeometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen24availableVirtualGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "availableVirtualGeometry", args)
  }

  return
}

// ~QScreen()
func (this *QScreen) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN7QScreenD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QScreen", "~QScreen", args)
  }

  return
}

// availableSize()
func (this *QScreen) AvailableSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen13availableSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "availableSize", args)
  }

  return
}

// virtualGeometry()
func (this *QScreen) VirtualGeometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen15virtualGeometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "virtualGeometry", args)
  }

  return
}

// physicalDotsPerInch()
func (this *QScreen) PhysicalDotsPerInch(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen19physicalDotsPerInchEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInch", args)
  }

  return
}

// physicalSize()
func (this *QScreen) PhysicalSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen12physicalSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSizeF{}) // "QSizeF"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "physicalSize", args)
  }

  return
}

// metaObject()
func (this *QScreen) MetaObject(args ...interface{}) () {
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
    C.C_ZNK7QScreen10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScreen", "metaObject", args)
  }

  return
}

// name()
func (this *QScreen) Name(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "name", args)
  }

  return
}

// logicalDotsPerInchY()
func (this *QScreen) LogicalDotsPerInchY(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen19logicalDotsPerInchYEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchY", args)
  }

  return
}

// logicalDotsPerInchX()
func (this *QScreen) LogicalDotsPerInchX(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen19logicalDotsPerInchXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "logicalDotsPerInchX", args)
  }

  return
}

// devicePixelRatio()
func (this *QScreen) DevicePixelRatio(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen16devicePixelRatioEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "devicePixelRatio", args)
  }

  return
}

// geometry()
func (this *QScreen) Geometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen8geometryEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "geometry", args)
  }

  return
}

// depth()
func (this *QScreen) Depth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen5depthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "depth", args)
  }

  return
}

// virtualSize()
func (this *QScreen) VirtualSize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen11virtualSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "virtualSize", args)
  }

  return
}

// physicalDotsPerInchX()
func (this *QScreen) PhysicalDotsPerInchX(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK7QScreen20physicalDotsPerInchXEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScreen", "physicalDotsPerInchX", args)
  }

  return
}

// <= body block end

