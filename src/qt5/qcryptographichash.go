package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qcryptographichash.h
// dst-file: /src/core/qcryptographichash.go
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
// class sizeof(QCryptographicHash)=8
type QCryptographicHash struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QCryptographicHash) addData(args ...interface{}) () {
  // addData(class QIODevice *)
  // addData(const char *, int)
  // addData(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHash7addDataEP9QIODevice
  case 1:
    // invoke: _ZN18QCryptographicHash7addDataEPKci
  case 2:
    // invoke: _ZN18QCryptographicHash7addDataERK10QByteArray
  default:
    qtrt.ErrorResolve("QCryptographicHash", "addData", args)
  }

}


func (this *QCryptographicHash) FreeQCryptographicHash(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCryptographicHash", "~QCryptographicHash", args)
  }

}


func (this *QCryptographicHash) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCryptographicHash5resetEv
  default:
    qtrt.ErrorResolve("QCryptographicHash", "reset", args)
  }

}


func (this *QCryptographicHash) result(args ...interface{}) () {
  // result()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCryptographicHash6resultEv
  default:
    qtrt.ErrorResolve("QCryptographicHash", "result", args)
  }

}


func NewQCryptographicHash(args ...interface{}) QCryptographicHash {
  return QCryptographicHash{}
}

// <= body block end

