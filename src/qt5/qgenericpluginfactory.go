package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
import "qtrt"
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
extern void C_ZN21QGenericPluginFactory6createERK7QStringS2_(void* arg0, void* arg1); // 4
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

// class sizeof(QGenericPluginFactory)=1
type QGenericPluginFactory struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// keys()
func (this *QGenericPluginFactory) keys_s(args ...interface{}) () {
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

}

// create(const class QString &, const class QString &)
func (this *QGenericPluginFactory) create_s(args ...interface{}) () {
  // create(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN21QGenericPluginFactory6createERK7QStringS2_
    // invoke: QObject * create(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN21QGenericPluginFactory6createERK7QStringS2_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QGenericPluginFactory", "create", args)
  }

}

// <= body block end

