package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  int QSharedMemory::size();
extern void _ZNK13QSharedMemory4sizeEv(void* qthis);
  // proto:  void QSharedMemory::setNativeKey(const QString & key);
extern void _ZN13QSharedMemory12setNativeKeyERK7QString(void* qthis, void* arg0);
  // proto:  void QSharedMemory::QSharedMemory(const QString & key, QObject * parent);
extern void* dector_ZN13QSharedMemoryC1ERK7QStringP7QObject(void* arg0, void* arg1);
extern void _ZN13QSharedMemoryC1ERK7QStringP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  QString QSharedMemory::errorString();
extern void _ZNK13QSharedMemory11errorStringEv(void* qthis);
  // proto:  void QSharedMemory::setKey(const QString & key);
extern void _ZN13QSharedMemory6setKeyERK7QString(void* qthis, void* arg0);
  // proto:  QString QSharedMemory::key();
extern void _ZNK13QSharedMemory3keyEv(void* qthis);
  // proto:  const void * QSharedMemory::constData();
extern void _ZNK13QSharedMemory9constDataEv(void* qthis);
  // proto:  void * QSharedMemory::data();
extern void _ZN13QSharedMemory4dataEv(void* qthis);
  // proto:  void QSharedMemory::QSharedMemory(const QSharedMemory & );
extern void* dector_ZN13QSharedMemoryC1ERKS_(void* arg0);
extern void _ZN13QSharedMemoryC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QSharedMemory::isAttached();
extern void _ZNK13QSharedMemory10isAttachedEv(void* qthis);
  // proto:  bool QSharedMemory::lock();
extern void _ZN13QSharedMemory4lockEv(void* qthis);
  // proto:  void QSharedMemory::~QSharedMemory();
extern void _ZN13QSharedMemoryD0Ev(void* qthis);
  // proto:  bool QSharedMemory::unlock();
extern void _ZN13QSharedMemory6unlockEv(void* qthis);
  // proto:  bool QSharedMemory::detach();
extern void _ZN13QSharedMemory6detachEv(void* qthis);
  // proto:  QString QSharedMemory::nativeKey();
extern void _ZNK13QSharedMemory9nativeKeyEv(void* qthis);
  // proto:  const QMetaObject * QSharedMemory::metaObject();
extern void _ZNK13QSharedMemory10metaObjectEv(void* qthis);
  // proto:  void QSharedMemory::QSharedMemory(QObject * parent);
extern void* dector_ZN13QSharedMemoryC1EP7QObject(void* arg0);
extern void _ZN13QSharedMemoryC1EP7QObject(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  int QSharedMemory::size();
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
    C._ZNK13QSharedMemory4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "size", args)
  }

}

  // proto:  void QSharedMemory::setNativeKey(const QString & key);
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
    C._ZN13QSharedMemory12setNativeKeyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "setNativeKey", args)
  }

}

  // proto:  void QSharedMemory::QSharedMemory(const QString & key, QObject * parent);
func NewQSharedMemory(args ...interface{}) QSharedMemory {
  return QSharedMemory{}
}

  // proto:  QString QSharedMemory::errorString();
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
    C._ZNK13QSharedMemory11errorStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "errorString", args)
  }

}

  // proto:  void QSharedMemory::setKey(const QString & key);
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
    C._ZN13QSharedMemory6setKeyERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSharedMemory", "setKey", args)
  }

}

  // proto:  QString QSharedMemory::key();
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
    C._ZNK13QSharedMemory3keyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "key", args)
  }

}

  // proto:  const void * QSharedMemory::constData();
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
    C._ZNK13QSharedMemory9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "constData", args)
  }

}

  // proto:  void * QSharedMemory::data();
func (this *QSharedMemory) data(args ...interface{}) () {
  // data()
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSharedMemory4dataEv
    // invoke: void * data()
    C._ZN13QSharedMemory4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "data", args)
  }

}

  // proto:  bool QSharedMemory::isAttached();
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
    C._ZNK13QSharedMemory10isAttachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "isAttached", args)
  }

}

  // proto:  bool QSharedMemory::lock();
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
    C._ZN13QSharedMemory4lockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "lock", args)
  }

}

  // proto:  void QSharedMemory::~QSharedMemory();
func (this *QSharedMemory) FreeQSharedMemory(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSharedMemory", "~QSharedMemory", args)
  }

}

  // proto:  bool QSharedMemory::unlock();
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
    C._ZN13QSharedMemory6unlockEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "unlock", args)
  }

}

  // proto:  bool QSharedMemory::detach();
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
    C._ZN13QSharedMemory6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "detach", args)
  }

}

  // proto:  QString QSharedMemory::nativeKey();
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
    C._ZNK13QSharedMemory9nativeKeyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "nativeKey", args)
  }

}

  // proto:  const QMetaObject * QSharedMemory::metaObject();
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
    C._ZNK13QSharedMemory10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSharedMemory", "metaObject", args)
  }

}

// <= body block end

