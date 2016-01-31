package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QString QSysInfo::kernelVersion();
extern void C_ZN8QSysInfo13kernelVersionEv(); // 4
  // proto: static QString QSysInfo::buildCpuArchitecture();
extern void C_ZN8QSysInfo20buildCpuArchitectureEv(); // 4
  // proto: static QString QSysInfo::prettyProductName();
extern void C_ZN8QSysInfo17prettyProductNameEv(); // 4
  // proto: static QString QSysInfo::productVersion();
extern void C_ZN8QSysInfo14productVersionEv(); // 4
  // proto: static QString QSysInfo::buildAbi();
extern void C_ZN8QSysInfo8buildAbiEv(); // 4
  // proto: static QString QSysInfo::kernelType();
extern void C_ZN8QSysInfo10kernelTypeEv(); // 4
  // proto: static QString QSysInfo::productType();
extern void C_ZN8QSysInfo11productTypeEv(); // 4
  // proto: static QSysInfo::WinVersion QSysInfo::windowsVersion();
extern void C_ZN8QSysInfo14windowsVersionEv(); // 2
  // proto: static QString QSysInfo::currentCpuArchitecture();
extern void C_ZN8QSysInfo22currentCpuArchitectureEv(); // 4
  // proto: static QSysInfo::MacVersion QSysInfo::macVersion();
extern void C_ZN8QSysInfo10macVersionEv(); // 2
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// kernelVersion()
func (this *QSysInfo) kernelVersion_s(args ...interface{}) () {
  // kernelVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo13kernelVersionEv
    // invoke: QString kernelVersion()
    var ret = C.C_ZN8QSysInfo13kernelVersionEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "kernelVersion", args)
  }

}

// buildCpuArchitecture()
func (this *QSysInfo) buildCpuArchitecture_s(args ...interface{}) () {
  // buildCpuArchitecture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo20buildCpuArchitectureEv
    // invoke: QString buildCpuArchitecture()
    var ret = C.C_ZN8QSysInfo20buildCpuArchitectureEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "buildCpuArchitecture", args)
  }

}

// prettyProductName()
func (this *QSysInfo) prettyProductName_s(args ...interface{}) () {
  // prettyProductName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo17prettyProductNameEv
    // invoke: QString prettyProductName()
    var ret = C.C_ZN8QSysInfo17prettyProductNameEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "prettyProductName", args)
  }

}

// productVersion()
func (this *QSysInfo) productVersion_s(args ...interface{}) () {
  // productVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo14productVersionEv
    // invoke: QString productVersion()
    var ret = C.C_ZN8QSysInfo14productVersionEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "productVersion", args)
  }

}

// buildAbi()
func (this *QSysInfo) buildAbi_s(args ...interface{}) () {
  // buildAbi()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo8buildAbiEv
    // invoke: QString buildAbi()
    var ret = C.C_ZN8QSysInfo8buildAbiEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "buildAbi", args)
  }

}

// kernelType()
func (this *QSysInfo) kernelType_s(args ...interface{}) () {
  // kernelType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo10kernelTypeEv
    // invoke: QString kernelType()
    var ret = C.C_ZN8QSysInfo10kernelTypeEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "kernelType", args)
  }

}

// productType()
func (this *QSysInfo) productType_s(args ...interface{}) () {
  // productType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo11productTypeEv
    // invoke: QString productType()
    var ret = C.C_ZN8QSysInfo11productTypeEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "productType", args)
  }

}

// windowsVersion()
func (this *QSysInfo) windowsVersion_s(args ...interface{}) () {
  // windowsVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo14windowsVersionEv
    // invoke: QSysInfo::WinVersion windowsVersion()
    C.C_ZN8QSysInfo14windowsVersionEv()
  default:
    qtrt.ErrorResolve("QSysInfo", "windowsVersion", args)
  }

}

// currentCpuArchitecture()
func (this *QSysInfo) currentCpuArchitecture_s(args ...interface{}) () {
  // currentCpuArchitecture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo22currentCpuArchitectureEv
    // invoke: QString currentCpuArchitecture()
    var ret = C.C_ZN8QSysInfo22currentCpuArchitectureEv()
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QSysInfo", "currentCpuArchitecture", args)
  }

}

// macVersion()
func (this *QSysInfo) macVersion_s(args ...interface{}) () {
  // macVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo10macVersionEv
    // invoke: QSysInfo::MacVersion macVersion()
    C.C_ZN8QSysInfo10macVersionEv()
  default:
    qtrt.ErrorResolve("QSysInfo", "macVersion", args)
  }

}

// <= body block end

