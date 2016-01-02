package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QLibraryInfo::QLibraryInfo();
extern void* dector_ZN12QLibraryInfoC1Ev();
extern void _ZN12QLibraryInfoC1Ev(void* qthis);
  // proto: static QStringList QLibraryInfo::platformPluginArguments(const QString & platformName);
extern void _ZN12QLibraryInfo23platformPluginArgumentsERK7QString(void* arg0);
  // proto: static QString QLibraryInfo::licensee();
extern void _ZN12QLibraryInfo8licenseeEv();
  // proto: static QString QLibraryInfo::licensedProducts();
extern void _ZN12QLibraryInfo16licensedProductsEv();
  // proto: static bool QLibraryInfo::isDebugBuild();
extern void _ZN12QLibraryInfo12isDebugBuildEv();
  // proto: static const char * QLibraryInfo::build();
extern void _ZN12QLibraryInfo5buildEv();
  // proto: static QDate QLibraryInfo::buildDate();
extern void _ZN12QLibraryInfo9buildDateEv();
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
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QLibraryInfo::QLibraryInfo();
func NewQLibraryInfo(args ...interface{}) QLibraryInfo {
  return QLibraryInfo{}
}

  // proto: static QStringList QLibraryInfo::platformPluginArguments(const QString & platformName);
func (this *QLibraryInfo) platformPluginArguments_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLibraryInfo", "platformPluginArguments", args)
  }

}

  // proto: static QString QLibraryInfo::licensee();
func (this *QLibraryInfo) licensee_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLibraryInfo", "licensee", args)
  }

}

  // proto: static QString QLibraryInfo::licensedProducts();
func (this *QLibraryInfo) licensedProducts_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLibraryInfo", "licensedProducts", args)
  }

}

  // proto: static bool QLibraryInfo::isDebugBuild();
func (this *QLibraryInfo) isDebugBuild_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLibraryInfo", "isDebugBuild", args)
  }

}

  // proto: static const char * QLibraryInfo::build();
func (this *QLibraryInfo) build_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLibraryInfo", "build", args)
  }

}

  // proto: static QDate QLibraryInfo::buildDate();
func (this *QLibraryInfo) buildDate_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QLibraryInfo", "buildDate", args)
  }

}

// <= body block end

