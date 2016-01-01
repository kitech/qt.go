package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qbuffer.h
// dst-file: /src/core/qbuffer.go
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
// class sizeof(QBuffer)=1
type QBuffer struct {
  /*qbase*/ QIODevice;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QBuffer) seek(args ...interface{}) () {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer4seekEx
  default:
    qtrt.ErrorResolve("QBuffer", "seek", args)
  }

}


func (this *QBuffer) canReadLine(args ...interface{}) () {
  // canReadLine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer11canReadLineEv
  default:
    qtrt.ErrorResolve("QBuffer", "canReadLine", args)
  }

}


func (this *QBuffer) FreeQBuffer(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QBuffer", "~QBuffer", args)
  }

}


func (this *QBuffer) setData(args ...interface{}) () {
  // setData(const class QByteArray &)
  // setData(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer7setDataERK10QByteArray
  case 1:
    // invoke: _ZN7QBuffer7setDataEPKci
  default:
    qtrt.ErrorResolve("QBuffer", "setData", args)
  }

}


func (this *QBuffer) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer4dataEv
  default:
    qtrt.ErrorResolve("QBuffer", "data", args)
  }

}


func NewQBuffer(args ...interface{}) QBuffer {
  return QBuffer{}
}


func (this *QBuffer) setBuffer(args ...interface{}) () {
  // setBuffer(class QByteArray *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer9setBufferEP10QByteArray
  default:
    qtrt.ErrorResolve("QBuffer", "setBuffer", args)
  }

}


func (this *QBuffer) buffer(args ...interface{}) () {
  // buffer()
  // buffer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer6bufferEv
  case 1:
    // invoke: _ZN7QBuffer6bufferEv
  default:
    qtrt.ErrorResolve("QBuffer", "buffer", args)
  }

}


func (this *QBuffer) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer3posEv
  default:
    qtrt.ErrorResolve("QBuffer", "pos", args)
  }

}


func (this *QBuffer) close(args ...interface{}) () {
  // close()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QBuffer5closeEv
  default:
    qtrt.ErrorResolve("QBuffer", "close", args)
  }

}


func (this *QBuffer) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer10metaObjectEv
  default:
    qtrt.ErrorResolve("QBuffer", "metaObject", args)
  }

}


func (this *QBuffer) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer4sizeEv
  default:
    qtrt.ErrorResolve("QBuffer", "size", args)
  }

}


func (this *QBuffer) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QBuffer5atEndEv
  default:
    qtrt.ErrorResolve("QBuffer", "atEnd", args)
  }

}

// <= body block end

