package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK11QFileDevice3posEv(void* qthis); // 4
  // proto:  bool QFileDevice::flush();
extern void _ZN11QFileDevice5flushEv(void* qthis); // 4
  // proto:  void QFileDevice::close();
extern void _ZN11QFileDevice5closeEv(void* qthis); // 4
  // proto:  bool QFileDevice::seek(qint64 offset);
extern void _ZN11QFileDevice4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  bool QFileDevice::isSequential();
extern void _ZNK11QFileDevice12isSequentialEv(void* qthis); // 4
  // proto:  qint64 QFileDevice::size();
extern void _ZNK11QFileDevice4sizeEv(void* qthis); // 4
  // proto:  void QFileDevice::unsetError();
extern void _ZN11QFileDevice10unsetErrorEv(void* qthis); // 4
  // proto:  int QFileDevice::handle();
extern void _ZNK11QFileDevice6handleEv(void* qthis); // 4
  // proto:  QString QFileDevice::fileName();
extern void _ZNK11QFileDevice8fileNameEv(void* qthis); // 4
  // proto:  const QMetaObject * QFileDevice::metaObject();
extern void _ZNK11QFileDevice10metaObjectEv(void* qthis); // 4
  // proto:  bool QFileDevice::resize(qint64 sz);
extern void _ZN11QFileDevice6resizeEx(void* qthis, int64_t arg0); // 4
  // proto:  Permissions QFileDevice::permissions();
extern void _ZNK11QFileDevice11permissionsEv(void* qthis); // 4
  // proto:  bool QFileDevice::unmap(uchar * address);
extern void _ZN11QFileDevice5unmapEPh(void* qthis, unsigned char* arg0); // 4
  // proto:  bool QFileDevice::atEnd();
extern void _ZNK11QFileDevice5atEndEv(void* qthis); // 4
  // proto:  void QFileDevice::~QFileDevice();
extern void _ZN11QFileDeviceD2Ev(void* qthis); // 4
  // proto:  QFileDevice::FileError QFileDevice::error();
extern void _ZNK11QFileDevice5errorEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// pos()
func (this *QFileDevice) pos(args ...interface{}) () {
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
    C._ZNK11QFileDevice3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "pos", args)
  }

}

// flush()
func (this *QFileDevice) flush(args ...interface{}) () {
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
    C._ZN11QFileDevice5flushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "flush", args)
  }

}

// close()
func (this *QFileDevice) close(args ...interface{}) () {
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
    C._ZN11QFileDevice5closeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "close", args)
  }

}

// seek(qint64)
func (this *QFileDevice) seek(args ...interface{}) () {
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDevice4seekEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDevice", "seek", args)
  }

}

// isSequential()
func (this *QFileDevice) isSequential(args ...interface{}) () {
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
    C._ZNK11QFileDevice12isSequentialEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "isSequential", args)
  }

}

// size()
func (this *QFileDevice) size(args ...interface{}) () {
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
    C._ZNK11QFileDevice4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "size", args)
  }

}

// unsetError()
func (this *QFileDevice) unsetError(args ...interface{}) () {
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
    C._ZN11QFileDevice10unsetErrorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "unsetError", args)
  }

}

// handle()
func (this *QFileDevice) handle(args ...interface{}) () {
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
    C._ZNK11QFileDevice6handleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "handle", args)
  }

}

// fileName()
func (this *QFileDevice) fileName(args ...interface{}) () {
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
    C._ZNK11QFileDevice8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "fileName", args)
  }

}

// metaObject()
func (this *QFileDevice) metaObject(args ...interface{}) () {
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
    C._ZNK11QFileDevice10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "metaObject", args)
  }

}

// resize(qint64)
func (this *QFileDevice) resize(args ...interface{}) () {
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
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDevice6resizeEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDevice", "resize", args)
  }

}

// permissions()
func (this *QFileDevice) permissions(args ...interface{}) () {
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
    C._ZNK11QFileDevice11permissionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "permissions", args)
  }

}

// unmap(uchar *)
func (this *QFileDevice) unmap(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN11QFileDevice5unmapEPh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileDevice", "unmap", args)
  }

}

// atEnd()
func (this *QFileDevice) atEnd(args ...interface{}) () {
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
    C._ZNK11QFileDevice5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "atEnd", args)
  }

}

// ~QFileDevice()
func (this *QFileDevice) FreeQFileDevice(args ...interface{}) () {
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
    C._ZN11QFileDeviceD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "~QFileDevice", args)
  }

}

// error()
func (this *QFileDevice) error(args ...interface{}) () {
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
    C._ZNK11QFileDevice5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileDevice", "error", args)
  }

}

// <= body block end

