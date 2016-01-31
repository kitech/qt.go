package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qstylefactory.h
// dst-file: /src/widgets/qstylefactory.go
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
  // proto: static QStringList QStyleFactory::keys();
extern void C_ZN13QStyleFactory4keysEv(); // 4
  // proto: static QStyle * QStyleFactory::create(const QString & );
extern void* C_ZN13QStyleFactory6createERK7QString(void* arg0); // 4
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

// class sizeof(QStyleFactory)=1
type QStyleFactory struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// keys()
func (this *QStyleFactory) Keys_S(args ...interface{}) () {
  // keys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStyleFactory4keysEv
    // invoke: QStringList keys()
    C.C_ZN13QStyleFactory4keysEv()
  default:
    qtrt.ErrorResolve("QStyleFactory", "keys", args)
  }

  return
}

// create(const class QString &)
func (this *QStyleFactory) Create_S(args ...interface{}) (ret interface{}) {
  // create(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStyleFactory6createERK7QString
    // invoke: QStyle * create(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QStyleFactory6createERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStyle{}) // "QStyle *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QStyleFactory", "create", args)
  }

  return
}

// <= body block end

