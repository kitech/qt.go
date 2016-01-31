package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtWidgets/qfileiconprovider.h
// dst-file: /src/widgets/qfileiconprovider.go
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
  // proto:  void QFileIconProvider::~QFileIconProvider();
extern void C_ZN17QFileIconProviderD2Ev(void* qthis); // 4
  // proto:  void QFileIconProvider::QFileIconProvider();
extern void C_ZN17QFileIconProviderC2Ev(void* qthis); // 3
  // proto:  QString QFileIconProvider::type(const QFileInfo & info);
extern void C_ZNK17QFileIconProvider4typeERK9QFileInfo(void* qthis, void* arg0); // 4
  // proto:  Options QFileIconProvider::options();
extern void C_ZNK17QFileIconProvider7optionsEv(void* qthis); // 4
  // proto:  QIcon QFileIconProvider::icon(const QFileInfo & info);
extern void C_ZNK17QFileIconProvider4iconERK9QFileInfo(void* qthis, void* arg0); // 4
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

// class sizeof(QFileIconProvider)=1
type QFileIconProvider struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QFileIconProvider()
func (this *QFileIconProvider) FreeQFileIconProvider(args ...interface{}) () {
  // ~QFileIconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QFileIconProviderD0Ev
    // invoke: void ~QFileIconProvider()
    C.C_ZN17QFileIconProviderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "~QFileIconProvider", args)
  }

}

// QFileIconProvider()
func NewQFileIconProvider(args ...interface{}) QFileIconProvider {
  // QFileIconProvider()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN17QFileIconProviderC1Ev
    // invoke: void QFileIconProvider()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN17QFileIconProviderC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "QFileIconProvider", args)
  }

  return QFileIconProvider{}
}

// type(const class QFileInfo &)
func (this *QFileIconProvider) type_(args ...interface{}) () {
  // type(const class QFileInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "const QFileInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider4typeERK9QFileInfo
    // invoke: QString type(const class QFileInfo &)
    var arg0 = args[0].(QFileInfo).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK17QFileIconProvider4typeERK9QFileInfo(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "type", args)
  }

}

// options()
func (this *QFileIconProvider) options(args ...interface{}) () {
  // options()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider7optionsEv
    // invoke: Options options()
    C.C_ZNK17QFileIconProvider7optionsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "options", args)
  }

}

// icon(const class QFileInfo &)
func (this *QFileIconProvider) icon(args ...interface{}) () {
  // icon(const class QFileInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "const QFileInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider4iconERK9QFileInfo
    // invoke: QIcon icon(const class QFileInfo &)
    var arg0 = args[0].(QFileInfo).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK17QFileIconProvider4iconERK9QFileInfo(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "icon", args)
  }

}

// <= body block end

