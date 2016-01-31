package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZN12QLibraryInfo9buildDateEv(); // 4
  // proto: static bool QLibraryInfo::isDebugBuild();
extern void C_ZN12QLibraryInfo12isDebugBuildEv(); // 4
  // proto: static QString QLibraryInfo::licensee();
extern void C_ZN12QLibraryInfo8licenseeEv(); // 4
  // proto: static QString QLibraryInfo::licensedProducts();
extern void C_ZN12QLibraryInfo16licensedProductsEv(); // 4
  // proto: static const char * QLibraryInfo::build();
extern void C_ZN12QLibraryInfo5buildEv(); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// buildDate()
func (this *QLibraryInfo) buildDate_s(args ...interface{}) () {
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
    C.C_ZN12QLibraryInfo9buildDateEv()
  default:
    qtrt.ErrorResolve("QLibraryInfo", "buildDate", args)
  }

}

// isDebugBuild()
func (this *QLibraryInfo) isDebugBuild_s(args ...interface{}) () {
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
    C.C_ZN12QLibraryInfo12isDebugBuildEv()
  default:
    qtrt.ErrorResolve("QLibraryInfo", "isDebugBuild", args)
  }

}

// licensee()
func (this *QLibraryInfo) licensee_s(args ...interface{}) () {
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
    C.C_ZN12QLibraryInfo8licenseeEv()
  default:
    qtrt.ErrorResolve("QLibraryInfo", "licensee", args)
  }

}

// licensedProducts()
func (this *QLibraryInfo) licensedProducts_s(args ...interface{}) () {
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
    C.C_ZN12QLibraryInfo16licensedProductsEv()
  default:
    qtrt.ErrorResolve("QLibraryInfo", "licensedProducts", args)
  }

}

// build()
func (this *QLibraryInfo) build_s(args ...interface{}) () {
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
    C.C_ZN12QLibraryInfo5buildEv()
  default:
    qtrt.ErrorResolve("QLibraryInfo", "build", args)
  }

}

// platformPluginArguments(const class QString &)
func (this *QLibraryInfo) platformPluginArguments_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QLibraryInfo23platformPluginArgumentsERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QLibraryInfo", "platformPluginArguments", args)
  }

}

// <= body block end

