package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtCore/qsysinfo.h
// dst-file: /src/core/qsysinfo.go
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
  // proto: static QString QSysInfo::kernelType();
extern void _ZN8QSysInfo10kernelTypeEv();
  // proto: static QString QSysInfo::productType();
extern void _ZN8QSysInfo11productTypeEv();
  // proto: static QString QSysInfo::prettyProductName();
extern void _ZN8QSysInfo17prettyProductNameEv();
  // proto: static QString QSysInfo::currentCpuArchitecture();
extern void _ZN8QSysInfo22currentCpuArchitectureEv();
  // proto: static QString QSysInfo::buildCpuArchitecture();
extern void _ZN8QSysInfo20buildCpuArchitectureEv();
  // proto: static QString QSysInfo::kernelVersion();
extern void _ZN8QSysInfo13kernelVersionEv();
  // proto: static QString QSysInfo::productVersion();
extern void _ZN8QSysInfo14productVersionEv();
  // proto: static QString QSysInfo::buildAbi();
extern void _ZN8QSysInfo8buildAbiEv();
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

// class sizeof(QSysInfo)=1
type QSysInfo struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static QString QSysInfo::kernelType();
func (this *QSysInfo) kernelType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "kernelType", args)
  }

}

  // proto: static QString QSysInfo::productType();
func (this *QSysInfo) productType_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "productType", args)
  }

}

  // proto: static QString QSysInfo::prettyProductName();
func (this *QSysInfo) prettyProductName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "prettyProductName", args)
  }

}

  // proto: static QString QSysInfo::currentCpuArchitecture();
func (this *QSysInfo) currentCpuArchitecture_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "currentCpuArchitecture", args)
  }

}

  // proto: static QString QSysInfo::buildCpuArchitecture();
func (this *QSysInfo) buildCpuArchitecture_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "buildCpuArchitecture", args)
  }

}

  // proto: static QString QSysInfo::kernelVersion();
func (this *QSysInfo) kernelVersion_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "kernelVersion", args)
  }

}

  // proto: static QString QSysInfo::productVersion();
func (this *QSysInfo) productVersion_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "productVersion", args)
  }

}

  // proto: static QString QSysInfo::buildAbi();
func (this *QSysInfo) buildAbi_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSysInfo", "buildAbi", args)
  }

}

// <= body block end

