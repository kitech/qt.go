package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtGui/qgenericpluginfactory.h
// dst-file: /src/gui/qgenericpluginfactory.go
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
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QStringList QGenericPluginFactory::keys();
extern void C_ZN21QGenericPluginFactory4keysEv(); // 4
  // proto: static QObject * QGenericPluginFactory::create(const QString & , const QString & );
extern void* C_ZN21QGenericPluginFactory6createERK7QStringS2_(void* arg0, void* arg1); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QGenericPluginFactory)=1
type QGenericPluginFactory struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// keys()
func (this *QGenericPluginFactory) Keys_s(args ...interface{}) () {
  // keys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGenericPluginFactory4keysEv
    // invoke: QStringList keys()
    C.C_ZN21QGenericPluginFactory4keysEv()
  default:
    qtrt.ErrorResolve("QGenericPluginFactory", "keys", args)
  }

  return
}

// create(const class QString &, const class QString &)
func (this *QGenericPluginFactory) Create_s(args ...interface{}) (ret interface{}) {
  // create(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGenericPluginFactory6createERK7QStringS2_
    // invoke: QObject * create(const class QString &, const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN21QGenericPluginFactory6createERK7QStringS2_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QGenericPluginFactory", "create", args)
  }

  return
}

// <= body block end

