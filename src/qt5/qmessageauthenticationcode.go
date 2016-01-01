package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qmessageauthenticationcode.h
// dst-file: /src/core/qmessageauthenticationcode.go
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
// class sizeof(QMessageAuthenticationCode)=8
type QMessageAuthenticationCode struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QMessageAuthenticationCode) result(args ...interface{}) () {
  // result()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QMessageAuthenticationCode6resultEv
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "result", args)
  }

}


func (this *QMessageAuthenticationCode) addData(args ...interface{}) () {
  // addData(const class QByteArray &)
  // addData(const char *, int)
  // addData(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QMessageAuthenticationCode7addDataERK10QByteArray
  case 1:
    // invoke: _ZN26QMessageAuthenticationCode7addDataEPKci
  case 2:
    // invoke: _ZN26QMessageAuthenticationCode7addDataEP9QIODevice
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "addData", args)
  }

}


func NewQMessageAuthenticationCode(args ...interface{}) QMessageAuthenticationCode {
  return QMessageAuthenticationCode{}
}


func (this *QMessageAuthenticationCode) FreeQMessageAuthenticationCode(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "~QMessageAuthenticationCode", args)
  }

}


func (this *QMessageAuthenticationCode) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QMessageAuthenticationCode5resetEv
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "reset", args)
  }

}


func (this *QMessageAuthenticationCode) setKey(args ...interface{}) () {
  // setKey(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QMessageAuthenticationCode6setKeyERK10QByteArray
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "setKey", args)
  }

}

// <= body block end

