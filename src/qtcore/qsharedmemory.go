package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern bool C_ZN13QSharedMemory4lockEv(void* qthis); // 4
  // proto:  bool QSharedMemory::unlock();
extern bool C_ZN13QSharedMemory6unlockEv(void* qthis); // 4
  // proto:  int QSharedMemory::size();
extern int32_t C_ZNK13QSharedMemory4sizeEv(void* qthis); // 4
  // proto:  const void * QSharedMemory::constData();
extern void C_ZNK13QSharedMemory9constDataEv(void* qthis); // 4
  // proto:  void QSharedMemory::setKey(const QString & key);
extern void C_ZN13QSharedMemory6setKeyERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QSharedMemory::isAttached();
extern bool C_ZNK13QSharedMemory10isAttachedEv(void* qthis); // 4
  // proto:  QString QSharedMemory::errorString();
extern void* C_ZNK13QSharedMemory11errorStringEv(void* qthis); // 4
  // proto:  QString QSharedMemory::key();
extern void* C_ZNK13QSharedMemory3keyEv(void* qthis); // 4
  // proto:  void QSharedMemory::QSharedMemory(const QString & key, QObject * parent);
extern void* C_ZN13QSharedMemoryC2ERK7QStringP7QObject(void* arg0, void* arg1); // 3
  // proto:  void QSharedMemory::QSharedMemory(QObject * parent);
extern void* C_ZN13QSharedMemoryC2EP7QObject(void* arg0); // 3
  // proto:  bool QSharedMemory::detach();
extern bool C_ZN13QSharedMemory6detachEv(void* qthis); // 4
  // proto:  void * QSharedMemory::data();
extern void C_ZN13QSharedMemory4dataEv(void* qthis); // 4
  // proto:  QString QSharedMemory::nativeKey();
extern void* C_ZNK13QSharedMemory9nativeKeyEv(void* qthis); // 4
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
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QSharedMemory)=1
type QSharedMemory struct {
  /*qbase*/ QObject;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// lock()
func (this *QSharedMemory) Lock(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QSharedMemory4lockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "lock", args)
  }

  return
}

// unlock()
func (this *QSharedMemory) Unlock(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QSharedMemory6unlockEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "unlock", args)
  }

  return
}

// size()
func (this *QSharedMemory) Size(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSharedMemory4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "size", args)
  }

  return
}

// constData()
func (this *QSharedMemory) ConstData(args ...interface{}) () {
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
    C.C_ZNK13QSharedMemory9constDataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "constData", args)
  }

  return
}

// setKey(const class QString &)
func (this *QSharedMemory) SetKey(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QSharedMemory6setKeyERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "setKey", args)
  }

  return
}

// isAttached()
func (this *QSharedMemory) IsAttached(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSharedMemory10isAttachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "isAttached", args)
  }

  return
}

// errorString()
func (this *QSharedMemory) ErrorString(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSharedMemory11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "errorString", args)
  }

  return
}

// key()
func (this *QSharedMemory) Key(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSharedMemory3keyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "key", args)
  }

  return
}

// QSharedMemory(const class QString &, class QObject *)
func GcfreeQSharedMemory(this *QSharedMemory) {
  qtrt.UniverseFree(this)
}
func NewQSharedMemory(args ...interface{}) *QSharedMemory {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QObject).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QSharedMemoryC2ERK7QStringP7QObject(arg0, arg1)
    this := &QSharedMemory{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQSharedMemory)
    return this
  case 1:
    // invoke: _ZN13QSharedMemoryC1EP7QObject
    // invoke: void QSharedMemory(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QSharedMemoryC2EP7QObject(arg0)
    this := &QSharedMemory{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQSharedMemory)
    return this
  default:
    qtrt.ErrorResolve("QSharedMemory", "QSharedMemory", args)
  }

  return nil // QSharedMemory{Qclsinst:qthis}
}

// detach()
func (this *QSharedMemory) Detach(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN13QSharedMemory6detachEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "detach", args)
  }

  return
}

// data()
func (this *QSharedMemory) Data(args ...interface{}) () {
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
    C.C_ZN13QSharedMemory4dataEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "data", args)
  }

  return
}

// nativeKey()
func (this *QSharedMemory) NativeKey(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSharedMemory9nativeKeyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "nativeKey", args)
  }

  return
}

// metaObject()
func (this *QSharedMemory) MetaObject(args ...interface{}) () {
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
    C.C_ZNK13QSharedMemory10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "metaObject", args)
  }

  return
}

// error()
func (this *QSharedMemory) Error(args ...interface{}) () {
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
    C.C_ZNK13QSharedMemory5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "error", args)
  }

  return
}

// setNativeKey(const class QString &)
func (this *QSharedMemory) SetNativeKey(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QSharedMemory12setNativeKeyERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "setNativeKey", args)
  }

  return
}

// ~QSharedMemory()
func (this *QSharedMemory) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QSharedMemoryD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QSharedMemory", "~QSharedMemory", args)
  }

  return
}

// <= body block end

