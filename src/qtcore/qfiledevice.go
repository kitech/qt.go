package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qfiledevice.h
// dst-file: /src/core/qfiledevice.go
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
  // proto:  qint64 QFileDevice::pos();
extern int64_t C_ZNK11QFileDevice3posEv(void* qthis); // 4
  // proto:  bool QFileDevice::flush();
extern bool C_ZN11QFileDevice5flushEv(void* qthis); // 4
  // proto:  void QFileDevice::close();
extern void C_ZN11QFileDevice5closeEv(void* qthis); // 4
  // proto:  bool QFileDevice::seek(qint64 offset);
extern bool C_ZN11QFileDevice4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QFileDevice::isSequential();
extern bool C_ZNK11QFileDevice12isSequentialEv(void* qthis); // 4
  // proto:  qint64 QFileDevice::size();
extern int64_t C_ZNK11QFileDevice4sizeEv(void* qthis); // 4
  // proto:  void QFileDevice::unsetError();
extern void C_ZN11QFileDevice10unsetErrorEv(void* qthis); // 4
  // proto:  int QFileDevice::handle();
extern int32_t C_ZNK11QFileDevice6handleEv(void* qthis); // 4
  // proto:  QString QFileDevice::fileName();
extern void* C_ZNK11QFileDevice8fileNameEv(void* qthis); // 4
  // proto:  const QMetaObject * QFileDevice::metaObject();
extern void C_ZNK11QFileDevice10metaObjectEv(void* qthis); // 4
  // proto:  bool QFileDevice::resize(qint64 sz);
extern bool C_ZN11QFileDevice6resizeEx(void* qthis, int64_t arg0); // 4
  // proto:  Permissions QFileDevice::permissions();
extern void C_ZNK11QFileDevice11permissionsEv(void* qthis); // 4
  // proto:  bool QFileDevice::unmap(uchar * address);
extern bool C_ZN11QFileDevice5unmapEPh(void* qthis, void* arg0); // 4
  // proto:  bool QFileDevice::atEnd();
extern bool C_ZNK11QFileDevice5atEndEv(void* qthis); // 4
  // proto:  void QFileDevice::~QFileDevice();
extern void C_ZN11QFileDeviceD2Ev(void* qthis); // 4
  // proto:  QFileDevice::FileError QFileDevice::error();
extern void C_ZNK11QFileDevice5errorEv(void* qthis); // 4
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

// class sizeof(QFileDevice)=1
type QFileDevice struct {
  /*qbase*/ QIODevice;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// pos()
func (this *QFileDevice) Pos(args ...interface{}) (ret interface{}) {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice3posEv
    // invoke: qint64 pos()
    var ret0 = C.C_ZNK11QFileDevice3posEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "pos", args)
  }

  return
}

// flush()
func (this *QFileDevice) Flush(args ...interface{}) (ret interface{}) {
  // flush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDevice5flushEv
    // invoke: bool flush()
    var ret0 = C.C_ZN11QFileDevice5flushEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "flush", args)
  }

  return
}

// close()
func (this *QFileDevice) Close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDevice5closeEv
    // invoke: void close()
    C.C_ZN11QFileDevice5closeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "close", args)
  }

  return
}

// seek(qint64)
func (this *QFileDevice) Seek(args ...interface{}) (ret interface{}) {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDevice4seekEx
    // invoke: bool seek(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QFileDevice4seekEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "seek", args)
  }

  return
}

// isSequential()
func (this *QFileDevice) Issequential(args ...interface{}) (ret interface{}) {
  // isSequential()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice12isSequentialEv
    // invoke: bool isSequential()
    var ret0 = C.C_ZNK11QFileDevice12isSequentialEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "isSequential", args)
  }

  return
}

// size()
func (this *QFileDevice) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice4sizeEv
    // invoke: qint64 size()
    var ret0 = C.C_ZNK11QFileDevice4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "size", args)
  }

  return
}

// unsetError()
func (this *QFileDevice) Unseterror(args ...interface{}) () {
  // unsetError()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDevice10unsetErrorEv
    // invoke: void unsetError()
    C.C_ZN11QFileDevice10unsetErrorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "unsetError", args)
  }

  return
}

// handle()
func (this *QFileDevice) Handle(args ...interface{}) (ret interface{}) {
  // handle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice6handleEv
    // invoke: int handle()
    var ret0 = C.C_ZNK11QFileDevice6handleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "handle", args)
  }

  return
}

// fileName()
func (this *QFileDevice) Filename(args ...interface{}) (ret interface{}) {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice8fileNameEv
    // invoke: QString fileName()
    var ret0 = C.C_ZNK11QFileDevice8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "fileName", args)
  }

  return
}

// metaObject()
func (this *QFileDevice) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QFileDevice10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "metaObject", args)
  }

  return
}

// resize(qint64)
func (this *QFileDevice) Resize(args ...interface{}) (ret interface{}) {
  // resize(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDevice6resizeEx
    // invoke: bool resize(qint64)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QFileDevice6resizeEx(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "resize", args)
  }

  return
}

// permissions()
func (this *QFileDevice) Permissions(args ...interface{}) () {
  // permissions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice11permissionsEv
    // invoke: Permissions permissions()
    C.C_ZNK11QFileDevice11permissionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "permissions", args)
  }

  return
}

// unmap(uchar *)
func (this *QFileDevice) Unmap(args ...interface{}) (ret interface{}) {
  // unmap(uchar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "uchar *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDevice5unmapEPh
    // invoke: bool unmap(uchar *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN11QFileDevice5unmapEPh(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "unmap", args)
  }

  return
}

// atEnd()
func (this *QFileDevice) Atend(args ...interface{}) (ret interface{}) {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice5atEndEv
    // invoke: bool atEnd()
    var ret0 = C.C_ZNK11QFileDevice5atEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileDevice", "atEnd", args)
  }

  return
}

// ~QFileDevice()
func (this *QFileDevice) Freeqfiledevice(args ...interface{}) () {
  // ~QFileDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QFileDeviceD0Ev
    // invoke: void ~QFileDevice()
    C.C_ZN11QFileDeviceD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "~QFileDevice", args)
  }

  return
}

// error()
func (this *QFileDevice) Error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QFileDevice5errorEv
    // invoke: QFileDevice::FileError error()
    C.C_ZNK11QFileDevice5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "error", args)
  }

  return
}

// <= body block end

