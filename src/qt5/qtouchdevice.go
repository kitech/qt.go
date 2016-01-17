package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QString QTouchDevice::name();
extern void _ZNK12QTouchDevice4nameEv(void* qthis); // 4
  // proto:  void QTouchDevice::setName(const QString & name);
extern void _ZN12QTouchDevice7setNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTouchDevice::~QTouchDevice();
extern void _ZN12QTouchDeviceD2Ev(void* qthis); // 4
  // proto:  void QTouchDevice::setMaximumTouchPoints(int max);
extern void _ZN12QTouchDevice21setMaximumTouchPointsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTouchDevice::maximumTouchPoints();
extern void _ZNK12QTouchDevice18maximumTouchPointsEv(void* qthis); // 4
  // proto: static QList<const QTouchDevice *> QTouchDevice::devices();
extern void _ZN12QTouchDevice7devicesEv(); // 4
  // proto:  Capabilities QTouchDevice::capabilities();
extern void _ZNK12QTouchDevice12capabilitiesEv(void* qthis); // 4
  // proto:  void QTouchDevice::QTouchDevice();
extern void _ZN12QTouchDeviceC2Ev(void* qthis); // 3
  // proto:  QTouchDevice::DeviceType QTouchDevice::type();
extern void _ZNK12QTouchDevice4typeEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// name()
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

// setName(const class QString &)
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

// ~QTouchDevice()
func (this *QTouchDevice) FreeQTouchDevice(args ...interface{}) () {
  // ~QTouchDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDeviceD0Ev
    // invoke: void ~QTouchDevice()
    C._ZN12QTouchDeviceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "~QTouchDevice", args)
  }

}

// setMaximumTouchPoints(int)
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

// maximumTouchPoints()
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

// devices()
func (this *QTouchDevice) devices_s(args ...interface{}) () {
  // devices()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDevice7devicesEv
    // invoke: QList<const QTouchDevice *> devices()
    C._ZN12QTouchDevice7devicesEv()
  default:
    qtrt.ErrorResolve("QTouchDevice", "devices", args)
  }

}

// capabilities()
func (this *QTouchDevice) capabilities(args ...interface{}) () {
  // capabilities()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTouchDevice12capabilitiesEv
    // invoke: Capabilities capabilities()
    C._ZNK12QTouchDevice12capabilitiesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "capabilities", args)
  }

}

// QTouchDevice()
func NewQTouchDevice(args ...interface{}) QTouchDevice {
  // QTouchDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTouchDeviceC1Ev
    // invoke: void QTouchDevice()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QTouchDeviceC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTouchDevice", "QTouchDevice", args)
  }

  return QTouchDevice{}
}

// type()
func (this *QTouchDevice) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTouchDevice4typeEv
    // invoke: QTouchDevice::DeviceType type()
    C._ZNK12QTouchDevice4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "type", args)
  }

}

// <= body block end

