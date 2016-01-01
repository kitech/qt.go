package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qsharedmemory.h
// dst-file: /src/core/qsharedmemory.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QSharedMemory)=1
type QSharedMemory struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "size", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "setNativeKey", args)
  }

}


func NewQSharedMemory(args ...interface{}) QSharedMemory {
  return QSharedMemory{}
}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "errorString", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "setKey", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "key", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "constData", args)
  }

}


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
    // invoke: _ZNK13QSharedMemory4dataEv
  case 1:
    // invoke: _ZN13QSharedMemory4dataEv
  default:
    qtrt.ErrorResolve("QSharedMemory", "data", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "isAttached", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "lock", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "unlock", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "detach", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "nativeKey", args)
  }

}


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
  default:
    qtrt.ErrorResolve("QSharedMemory", "metaObject", args)
  }

}

// <= body block end

