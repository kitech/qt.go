package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtCore/qlibraryinfo.h
// dst-file: /src/core/qlibraryinfo.go
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
  // proto: static QDate QLibraryInfo::buildDate();
extern void* C_ZN12QLibraryInfo9buildDateEv(); // 4
  // proto: static bool QLibraryInfo::isDebugBuild();
extern bool C_ZN12QLibraryInfo12isDebugBuildEv(); // 4
  // proto: static QString QLibraryInfo::licensee();
extern void* C_ZN12QLibraryInfo8licenseeEv(); // 4
  // proto: static QString QLibraryInfo::licensedProducts();
extern void* C_ZN12QLibraryInfo16licensedProductsEv(); // 4
  // proto: static const char * QLibraryInfo::build();
extern void* C_ZN12QLibraryInfo5buildEv(); // 4
  // proto: static QStringList QLibraryInfo::platformPluginArguments(const QString & platformName);
extern void C_ZN12QLibraryInfo23platformPluginArgumentsERK7QString(void* arg0); // 4
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

// class sizeof(QLibraryInfo)=1
type QLibraryInfo struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// buildDate()
func (this *QLibraryInfo) Builddate_S(args ...interface{}) (ret interface{}) {
  // buildDate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QLibraryInfo9buildDateEv
    // invoke: QDate buildDate()
    var ret0 = C.C_ZN12QLibraryInfo9buildDateEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QDate{}) // "QDate"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibraryInfo", "buildDate", args)
  }

  return
}

// isDebugBuild()
func (this *QLibraryInfo) Isdebugbuild_S(args ...interface{}) (ret interface{}) {
  // isDebugBuild()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QLibraryInfo12isDebugBuildEv
    // invoke: bool isDebugBuild()
    var ret0 = C.C_ZN12QLibraryInfo12isDebugBuildEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibraryInfo", "isDebugBuild", args)
  }

  return
}

// licensee()
func (this *QLibraryInfo) Licensee_S(args ...interface{}) (ret interface{}) {
  // licensee()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QLibraryInfo8licenseeEv
    // invoke: QString licensee()
    var ret0 = C.C_ZN12QLibraryInfo8licenseeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibraryInfo", "licensee", args)
  }

  return
}

// licensedProducts()
func (this *QLibraryInfo) Licensedproducts_S(args ...interface{}) (ret interface{}) {
  // licensedProducts()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QLibraryInfo16licensedProductsEv
    // invoke: QString licensedProducts()
    var ret0 = C.C_ZN12QLibraryInfo16licensedProductsEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibraryInfo", "licensedProducts", args)
  }

  return
}

// build()
func (this *QLibraryInfo) Build_S(args ...interface{}) (ret interface{}) {
  // build()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QLibraryInfo5buildEv
    // invoke: const char * build()
    var ret0 = C.C_ZN12QLibraryInfo5buildEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLibraryInfo", "build", args)
  }

  return
}

// platformPluginArguments(const class QString &)
func (this *QLibraryInfo) Platformpluginarguments_S(args ...interface{}) () {
  // platformPluginArguments(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QLibraryInfo23platformPluginArgumentsERK7QString
    // invoke: QStringList platformPluginArguments(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QLibraryInfo23platformPluginArgumentsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QLibraryInfo", "platformPluginArguments", args)
  }

  return
}

// <= body block end

