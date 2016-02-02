package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern void* C_ZNK12QTouchDevice4nameEv(void* qthis); // 4
  // proto:  void QTouchDevice::setName(const QString & name);
extern void C_ZN12QTouchDevice7setNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTouchDevice::~QTouchDevice();
extern void C_ZN12QTouchDeviceD2Ev(void* qthis); // 4
  // proto:  void QTouchDevice::setMaximumTouchPoints(int max);
extern void C_ZN12QTouchDevice21setMaximumTouchPointsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTouchDevice::maximumTouchPoints();
extern int32_t C_ZNK12QTouchDevice18maximumTouchPointsEv(void* qthis); // 4
  // proto: static QList<const QTouchDevice *> QTouchDevice::devices();
extern void C_ZN12QTouchDevice7devicesEv(); // 4
  // proto:  Capabilities QTouchDevice::capabilities();
extern void C_ZNK12QTouchDevice12capabilitiesEv(void* qthis); // 4
  // proto:  void QTouchDevice::QTouchDevice();
extern void* C_ZN12QTouchDeviceC2Ev(); // 3
  // proto:  QTouchDevice::DeviceType QTouchDevice::type();
extern void C_ZNK12QTouchDevice4typeEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// name()
func (this *QTouchDevice) Name(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTouchDevice4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTouchDevice", "name", args)
  }

  return
}

// setName(const class QString &)
func (this *QTouchDevice) Setname(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTouchDevice7setNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchDevice", "setName", args)
  }

  return
}

// ~QTouchDevice()
func (this *QTouchDevice) Freeqtouchdevice(args ...interface{}) () {
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
    C.C_ZN12QTouchDeviceD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "~QTouchDevice", args)
  }

  return
}

// setMaximumTouchPoints(int)
func (this *QTouchDevice) Setmaximumtouchpoints(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTouchDevice21setMaximumTouchPointsEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTouchDevice", "setMaximumTouchPoints", args)
  }

  return
}

// maximumTouchPoints()
func (this *QTouchDevice) Maximumtouchpoints(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTouchDevice18maximumTouchPointsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTouchDevice", "maximumTouchPoints", args)
  }

  return
}

// devices()
func (this *QTouchDevice) Devices_S(args ...interface{}) () {
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
    C.C_ZN12QTouchDevice7devicesEv()
  default:
    qtrt.ErrorResolve("QTouchDevice", "devices", args)
  }

  return
}

// capabilities()
func (this *QTouchDevice) Capabilities(args ...interface{}) () {
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
    C.C_ZNK12QTouchDevice12capabilitiesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "capabilities", args)
  }

  return
}

// QTouchDevice()
func NewQTouchDevice(args ...interface{}) *QTouchDevice {
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
    qthis = C.C_ZN12QTouchDeviceC2Ev()
    return &QTouchDevice{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTouchDevice", "QTouchDevice", args)
  }

  return nil // QTouchDevice{Qclsinst:qthis}
}

// type()
func (this *QTouchDevice) Type_(args ...interface{}) () {
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
    C.C_ZNK12QTouchDevice4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTouchDevice", "type", args)
  }

  return
}

// <= body block end

