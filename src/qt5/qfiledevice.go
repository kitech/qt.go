package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  qint64 QFileDevice::size();
extern void _ZNK11QFileDevice4sizeEv(void* qthis);
  // proto:  bool QFileDevice::seek(qint64 offset);
extern void _ZN11QFileDevice4seekEx(void* qthis, long long arg0);
  // proto:  bool QFileDevice::unmap(uchar * address);
extern void _ZN11QFileDevice5unmapEPh(void* qthis, unsigned char* arg0);
  // proto:  void QFileDevice::close();
extern void _ZN11QFileDevice5closeEv(void* qthis);
  // proto:  qint64 QFileDevice::pos();
extern void _ZNK11QFileDevice3posEv(void* qthis);
  // proto:  int QFileDevice::handle();
extern void _ZNK11QFileDevice6handleEv(void* qthis);
  // proto:  QString QFileDevice::fileName();
extern void _ZNK11QFileDevice8fileNameEv(void* qthis);
  // proto:  void QFileDevice::QFileDevice(QObject * parent);
extern void* dector_ZN11QFileDeviceC1EP7QObject(void* arg0);
extern void _ZN11QFileDeviceC1EP7QObject(void* qthis, void* arg0);
  // proto:  void QFileDevice::~QFileDevice();
extern void _ZN11QFileDeviceD0Ev(void* qthis);
  // proto:  bool QFileDevice::atEnd();
extern void _ZNK11QFileDevice5atEndEv(void* qthis);
  // proto:  bool QFileDevice::isSequential();
extern void _ZNK11QFileDevice12isSequentialEv(void* qthis);
  // proto:  bool QFileDevice::flush();
extern void _ZN11QFileDevice5flushEv(void* qthis);
  // proto:  void QFileDevice::QFileDevice();
extern void* dector_ZN11QFileDeviceC1Ev();
extern void _ZN11QFileDeviceC1Ev(void* qthis);
  // proto:  void QFileDevice::unsetError();
extern void _ZN11QFileDevice10unsetErrorEv(void* qthis);
  // proto:  void QFileDevice::QFileDevice(const QFileDevice & );
extern void* dector_ZN11QFileDeviceC1ERKS_(void* arg0);
extern void _ZN11QFileDeviceC1ERKS_(void* qthis, void* arg0);
  // proto:  const QMetaObject * QFileDevice::metaObject();
extern void _ZNK11QFileDevice10metaObjectEv(void* qthis);
  // proto:  bool QFileDevice::resize(qint64 sz);
extern void _ZN11QFileDevice6resizeEx(void* qthis, long long arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  qint64 QFileDevice::size();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "size", args)
  }

}

  // proto:  bool QFileDevice::seek(qint64 offset);
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
  default:
    qtrt.ErrorResolve("QFileDevice", "seek", args)
  }

}

  // proto:  bool QFileDevice::unmap(uchar * address);
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
  default:
    qtrt.ErrorResolve("QFileDevice", "unmap", args)
  }

}

  // proto:  void QFileDevice::close();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "close", args)
  }

}

  // proto:  qint64 QFileDevice::pos();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "pos", args)
  }

}

  // proto:  int QFileDevice::handle();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "handle", args)
  }

}

  // proto:  QString QFileDevice::fileName();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "fileName", args)
  }

}

  // proto:  void QFileDevice::QFileDevice(QObject * parent);
func NewQFileDevice(args ...interface{}) QFileDevice {
  return QFileDevice{}
}

  // proto:  void QFileDevice::~QFileDevice();
func (this *QFileDevice) FreeQFileDevice(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QFileDevice", "~QFileDevice", args)
  }

}

  // proto:  bool QFileDevice::atEnd();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "atEnd", args)
  }

}

  // proto:  bool QFileDevice::isSequential();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "isSequential", args)
  }

}

  // proto:  bool QFileDevice::flush();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "flush", args)
  }

}

  // proto:  void QFileDevice::unsetError();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "unsetError", args)
  }

}

  // proto:  const QMetaObject * QFileDevice::metaObject();
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
  default:
    qtrt.ErrorResolve("QFileDevice", "metaObject", args)
  }

}

  // proto:  bool QFileDevice::resize(qint64 sz);
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
  default:
    qtrt.ErrorResolve("QFileDevice", "resize", args)
  }

}

// <= body block end

