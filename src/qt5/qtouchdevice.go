package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qtouchdevice.h
// dst-file: /src/gui/qtouchdevice.go
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
  // proto:  void QTouchDevice::setName(const QString & name);
extern void _ZN12QTouchDevice7setNameERK7QString(void* qthis, void* arg0);
  // proto:  QString QTouchDevice::name();
extern void _ZNK12QTouchDevice4nameEv(void* qthis);
  // proto:  void QTouchDevice::setMaximumTouchPoints(int max);
extern void _ZN12QTouchDevice21setMaximumTouchPointsEi(void* qthis, int arg0);
  // proto: static QList<const QTouchDevice *> QTouchDevice::devices();
extern void _ZN12QTouchDevice7devicesEv();
  // proto:  void QTouchDevice::QTouchDevice();
extern void* dector_ZN12QTouchDeviceC1Ev();
extern void _ZN12QTouchDeviceC1Ev(void* qthis);
  // proto:  void QTouchDevice::~QTouchDevice();
extern void _ZN12QTouchDeviceD0Ev(void* qthis);
  // proto:  int QTouchDevice::maximumTouchPoints();
extern void _ZNK12QTouchDevice18maximumTouchPointsEv(void* qthis);
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

// class sizeof(QTouchDevice)=8
type QTouchDevice struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QTouchDevice::setName(const QString & name);
func (this *QTouchDevice) setName(args ...interface{}) () {
  // setName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDevice7setNameERK7QString
    // invoke: void setName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QTouchDevice7setNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchDevice", "setName", args)
  }

}

  // proto:  QString QTouchDevice::name();
func (this *QTouchDevice) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTouchDevice4nameEv
    // invoke: QString name()
    C._ZNK12QTouchDevice4nameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "name", args)
  }

}

  // proto:  void QTouchDevice::setMaximumTouchPoints(int max);
func (this *QTouchDevice) setMaximumTouchPoints(args ...interface{}) () {
  // setMaximumTouchPoints(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDevice21setMaximumTouchPointsEi
    // invoke: void setMaximumTouchPoints(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN12QTouchDevice21setMaximumTouchPointsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchDevice", "setMaximumTouchPoints", args)
  }

}

  // proto: static QList<const QTouchDevice *> QTouchDevice::devices();
func (this *QTouchDevice) devices_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTouchDevice", "devices", args)
  }

}

  // proto:  void QTouchDevice::QTouchDevice();
func NewQTouchDevice(args ...interface{}) QTouchDevice {
  return QTouchDevice{}
}

  // proto:  void QTouchDevice::~QTouchDevice();
func (this *QTouchDevice) FreeQTouchDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTouchDevice", "~QTouchDevice", args)
  }

}

  // proto:  int QTouchDevice::maximumTouchPoints();
func (this *QTouchDevice) maximumTouchPoints(args ...interface{}) () {
  // maximumTouchPoints()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTouchDevice18maximumTouchPointsEv
    // invoke: int maximumTouchPoints()
    C._ZNK12QTouchDevice18maximumTouchPointsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "maximumTouchPoints", args)
  }

}

// <= body block end

