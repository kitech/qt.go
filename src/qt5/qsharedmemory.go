package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qsharedmemory.h
// dst-file: /src/core/qsharedmemory.go
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
  // proto:  bool QSharedMemory::lock();
extern void C_ZN13QSharedMemory4lockEv(void* qthis); // 4
  // proto:  bool QSharedMemory::unlock();
extern void C_ZN13QSharedMemory6unlockEv(void* qthis); // 4
  // proto:  int QSharedMemory::size();
extern void C_ZNK13QSharedMemory4sizeEv(void* qthis); // 4
  // proto:  const void * QSharedMemory::constData();
extern void C_ZNK13QSharedMemory9constDataEv(void* qthis); // 4
  // proto:  void QSharedMemory::setKey(const QString & key);
extern void C_ZN13QSharedMemory6setKeyERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QSharedMemory::isAttached();
extern void C_ZNK13QSharedMemory10isAttachedEv(void* qthis); // 4
  // proto:  QString QSharedMemory::errorString();
extern void C_ZNK13QSharedMemory11errorStringEv(void* qthis); // 4
  // proto:  QString QSharedMemory::key();
extern void C_ZNK13QSharedMemory3keyEv(void* qthis); // 4
  // proto:  void QSharedMemory::QSharedMemory(const QString & key, QObject * parent);
extern void C_ZN13QSharedMemoryC2ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QSharedMemory::QSharedMemory(QObject * parent);
extern void C_ZN13QSharedMemoryC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  bool QSharedMemory::detach();
extern void C_ZN13QSharedMemory6detachEv(void* qthis); // 4
  // proto:  void * QSharedMemory::data();
extern void C_ZN13QSharedMemory4dataEv(void* qthis); // 4
  // proto:  QString QSharedMemory::nativeKey();
extern void C_ZNK13QSharedMemory9nativeKeyEv(void* qthis); // 4
  // proto:  const QMetaObject * QSharedMemory::metaObject();
extern void C_ZNK13QSharedMemory10metaObjectEv(void* qthis); // 4
  // proto:  QSharedMemory::SharedMemoryError QSharedMemory::error();
extern void C_ZNK13QSharedMemory5errorEv(void* qthis); // 4
  // proto:  void QSharedMemory::setNativeKey(const QString & key);
extern void C_ZN13QSharedMemory12setNativeKeyERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QSharedMemory::~QSharedMemory();
extern void C_ZN13QSharedMemoryD2Ev(void* qthis); // 4
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

// class sizeof(QSharedMemory)=1
type QSharedMemory struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// lock()
func (this *QSharedMemory) lock(args ...interface{}) () {
  // lock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory4lockEv
    // invoke: bool lock()
    C.C_ZN13QSharedMemory4lockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "lock", args)
  }

}

// unlock()
func (this *QSharedMemory) unlock(args ...interface{}) () {
  // unlock()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory6unlockEv
    // invoke: bool unlock()
    C.C_ZN13QSharedMemory6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "unlock", args)
  }

}

// size()
func (this *QSharedMemory) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory4sizeEv
    // invoke: int size()
    C.C_ZNK13QSharedMemory4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "size", args)
  }

}

// constData()
func (this *QSharedMemory) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory9constDataEv
    // invoke: const void * constData()
    C.C_ZNK13QSharedMemory9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "constData", args)
  }

}

// setKey(const class QString &)
func (this *QSharedMemory) setKey(args ...interface{}) () {
  // setKey(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory6setKeyERK7QString
    // invoke: void setKey(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QSharedMemory6setKeyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "setKey", args)
  }

}

// isAttached()
func (this *QSharedMemory) isAttached(args ...interface{}) () {
  // isAttached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory10isAttachedEv
    // invoke: bool isAttached()
    C.C_ZNK13QSharedMemory10isAttachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "isAttached", args)
  }

}

// errorString()
func (this *QSharedMemory) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory11errorStringEv
    // invoke: QString errorString()
    C.C_ZNK13QSharedMemory11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "errorString", args)
  }

}

// key()
func (this *QSharedMemory) key(args ...interface{}) () {
  // key()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory3keyEv
    // invoke: QString key()
    C.C_ZNK13QSharedMemory3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "key", args)
  }

}

// QSharedMemory(const class QString &, class QObject *)
func NewQSharedMemory(args ...interface{}) QSharedMemory {
  // QSharedMemory(const class QString &, class QObject *)
  // QSharedMemory(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemoryC1ERK7QStringP7QObject
    // invoke: void QSharedMemory(const class QString &, class QObject *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QSharedMemoryC2ERK7QStringP7QObject(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN13QSharedMemoryC1EP7QObject
    // invoke: void QSharedMemory(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN13QSharedMemoryC2EP7QObject(qthis, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "QSharedMemory", args)
  }

  return QSharedMemory{}
}

// detach()
func (this *QSharedMemory) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory6detachEv
    // invoke: bool detach()
    C.C_ZN13QSharedMemory6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "detach", args)
  }

}

// data()
func (this *QSharedMemory) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory4dataEv
    // invoke: void * data()
    C.C_ZN13QSharedMemory4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "data", args)
  }

}

// nativeKey()
func (this *QSharedMemory) nativeKey(args ...interface{}) () {
  // nativeKey()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory9nativeKeyEv
    // invoke: QString nativeKey()
    C.C_ZNK13QSharedMemory9nativeKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "nativeKey", args)
  }

}

// metaObject()
func (this *QSharedMemory) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK13QSharedMemory10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "metaObject", args)
  }

}

// error()
func (this *QSharedMemory) error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSharedMemory5errorEv
    // invoke: QSharedMemory::SharedMemoryError error()
    C.C_ZNK13QSharedMemory5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "error", args)
  }

}

// setNativeKey(const class QString &)
func (this *QSharedMemory) setNativeKey(args ...interface{}) () {
  // setNativeKey(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory12setNativeKeyERK7QString
    // invoke: void setNativeKey(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QSharedMemory12setNativeKeyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "setNativeKey", args)
  }

}

// ~QSharedMemory()
func (this *QSharedMemory) FreeQSharedMemory(args ...interface{}) () {
  // ~QSharedMemory()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemoryD0Ev
    // invoke: void ~QSharedMemory()
    C.C_ZN13QSharedMemoryD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "~QSharedMemory", args)
  }

}

// <= body block end

