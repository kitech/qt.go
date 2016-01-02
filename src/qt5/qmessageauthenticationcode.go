package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtCore/qmessageauthenticationcode.h
// dst-file: /src/core/qmessageauthenticationcode.go
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
  // proto:  QByteArray QMessageAuthenticationCode::result();
extern void _ZNK26QMessageAuthenticationCode6resultEv(void* qthis);
  // proto:  void QMessageAuthenticationCode::addData(const QByteArray & data);
extern void _ZN26QMessageAuthenticationCode7addDataERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QMessageAuthenticationCode::QMessageAuthenticationCode(const QMessageAuthenticationCode & );
extern void* dector_ZN26QMessageAuthenticationCodeC1ERKS_(void* arg0);
extern void _ZN26QMessageAuthenticationCodeC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMessageAuthenticationCode::addData(const char * data, int length);
extern void _ZN26QMessageAuthenticationCode7addDataEPKci(void* qthis, char* arg0, int arg1);
  // proto:  void QMessageAuthenticationCode::~QMessageAuthenticationCode();
extern void _ZN26QMessageAuthenticationCodeD0Ev(void* qthis);
  // proto:  void QMessageAuthenticationCode::reset();
extern void _ZN26QMessageAuthenticationCode5resetEv(void* qthis);
  // proto:  bool QMessageAuthenticationCode::addData(QIODevice * device);
extern void _ZN26QMessageAuthenticationCode7addDataEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QMessageAuthenticationCode::setKey(const QByteArray & key);
extern void _ZN26QMessageAuthenticationCode6setKeyERK10QByteArray(void* qthis, void* arg0);
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

// class sizeof(QMessageAuthenticationCode)=8
type QMessageAuthenticationCode struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QByteArray QMessageAuthenticationCode::result();
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

  // proto:  void QMessageAuthenticationCode::addData(const QByteArray & data);
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN26QMessageAuthenticationCode7addDataEPKci
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  case 2:
    // invoke: _ZN26QMessageAuthenticationCode7addDataEP9QIODevice
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "addData", args)
  }

}

  // proto:  void QMessageAuthenticationCode::QMessageAuthenticationCode(const QMessageAuthenticationCode & );
func NewQMessageAuthenticationCode(args ...interface{}) QMessageAuthenticationCode {
  return QMessageAuthenticationCode{}
}

  // proto:  void QMessageAuthenticationCode::~QMessageAuthenticationCode();
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

  // proto:  void QMessageAuthenticationCode::reset();
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

  // proto:  void QMessageAuthenticationCode::setKey(const QByteArray & key);
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QMessageAuthenticationCode", "setKey", args)
  }

}

// <= body block end

