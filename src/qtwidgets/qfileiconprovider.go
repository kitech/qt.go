package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
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
extern void* C_ZN17QFileIconProviderC2Ev(); // 3
  // proto:  QString QFileIconProvider::type(const QFileInfo & info);
extern void* C_ZNK17QFileIconProvider4typeERK9QFileInfo(void* qthis, void* arg0); // 4
  // proto:  Options QFileIconProvider::options();
extern void C_ZNK17QFileIconProvider7optionsEv(void* qthis); // 4
  // proto:  QIcon QFileIconProvider::icon(const QFileInfo & info);
extern void* C_ZNK17QFileIconProvider4iconERK9QFileInfo(void* qthis, void* arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QFileIconProvider)=1
type QFileIconProvider struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QFileIconProvider()
func (this *QFileIconProvider) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN17QFileIconProviderD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QFileIconProvider", "~QFileIconProvider", args)
  }

  return
}

// QFileIconProvider()
func GcfreeQFileIconProvider(this *QFileIconProvider) {
  qtrt.UniverseFree(this)
}
func NewQFileIconProvider(args ...interface{}) *QFileIconProvider {
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
    qthis = C.C_ZN17QFileIconProviderC2Ev()
    this := &QFileIconProvider{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQFileIconProvider)
    return this
  default:
    qtrt.ErrorResolve("QFileIconProvider", "QFileIconProvider", args)
  }

  return nil // QFileIconProvider{Qclsinst:qthis}
}

// type(const class QFileInfo &)
func (this *QFileIconProvider) Type_(args ...interface{}) (ret interface{}) {
  // type(const class QFileInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QFileInfo{}) // "const QFileInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider4typeERK9QFileInfo
    // invoke: QString type(const class QFileInfo &)
    var arg0 = args[0].(*qtcore.QFileInfo).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QFileIconProvider4typeERK9QFileInfo(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileIconProvider", "type", args)
  }

  return
}

// options()
func (this *QFileIconProvider) Options(args ...interface{}) () {
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
    C.C_ZNK17QFileIconProvider7optionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QFileIconProvider", "options", args)
  }

  return
}

// icon(const class QFileInfo &)
func (this *QFileIconProvider) Icon(args ...interface{}) (ret interface{}) {
  // icon(const class QFileInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QFileInfo{}) // "const QFileInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK17QFileIconProvider4iconERK9QFileInfo
    // invoke: QIcon icon(const class QFileInfo &)
    var arg0 = args[0].(*qtcore.QFileInfo).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK17QFileIconProvider4iconERK9QFileInfo(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QFileIconProvider", "icon", args)
  }

  return
}

// <= body block end

