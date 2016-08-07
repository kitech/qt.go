package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
extern void* C_ZN8QSysInfo13kernelVersionEv(); // 4
  // proto: static QString QSysInfo::machineHostName();
extern void* C_ZN8QSysInfo15machineHostNameEv(); // 4
  // proto: static QString QSysInfo::prettyProductName();
extern void* C_ZN8QSysInfo17prettyProductNameEv(); // 4
  // proto: static QString QSysInfo::productVersion();
extern void* C_ZN8QSysInfo14productVersionEv(); // 4
  // proto: static QString QSysInfo::buildAbi();
extern void* C_ZN8QSysInfo8buildAbiEv(); // 4
  // proto: static QString QSysInfo::buildCpuArchitecture();
extern void* C_ZN8QSysInfo20buildCpuArchitectureEv(); // 4
  // proto: static QString QSysInfo::kernelType();
extern void* C_ZN8QSysInfo10kernelTypeEv(); // 4
  // proto: static QString QSysInfo::productType();
extern void* C_ZN8QSysInfo11productTypeEv(); // 4
  // proto: static QSysInfo::WinVersion QSysInfo::windowsVersion();
extern void C_ZN8QSysInfo14windowsVersionEv(); // 2
  // proto: static QString QSysInfo::currentCpuArchitecture();
extern void* C_ZN8QSysInfo22currentCpuArchitectureEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// kernelVersion()
func (this *QSysInfo) Kernelversion_S(args ...interface{}) (ret interface{}) {
  // kernelVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo13kernelVersionEv
    // invoke: QString kernelVersion()
    var ret0 = C.C_ZN8QSysInfo13kernelVersionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "kernelVersion", args)
  }

  return
}

// machineHostName()
func (this *QSysInfo) Machinehostname_S(args ...interface{}) (ret interface{}) {
  // machineHostName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo15machineHostNameEv
    // invoke: QString machineHostName()
    var ret0 = C.C_ZN8QSysInfo15machineHostNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "machineHostName", args)
  }

  return
}

// prettyProductName()
func (this *QSysInfo) Prettyproductname_S(args ...interface{}) (ret interface{}) {
  // prettyProductName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo17prettyProductNameEv
    // invoke: QString prettyProductName()
    var ret0 = C.C_ZN8QSysInfo17prettyProductNameEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "prettyProductName", args)
  }

  return
}

// productVersion()
func (this *QSysInfo) Productversion_S(args ...interface{}) (ret interface{}) {
  // productVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo14productVersionEv
    // invoke: QString productVersion()
    var ret0 = C.C_ZN8QSysInfo14productVersionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "productVersion", args)
  }

  return
}

// buildAbi()
func (this *QSysInfo) Buildabi_S(args ...interface{}) (ret interface{}) {
  // buildAbi()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo8buildAbiEv
    // invoke: QString buildAbi()
    var ret0 = C.C_ZN8QSysInfo8buildAbiEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "buildAbi", args)
  }

  return
}

// buildCpuArchitecture()
func (this *QSysInfo) Buildcpuarchitecture_S(args ...interface{}) (ret interface{}) {
  // buildCpuArchitecture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo20buildCpuArchitectureEv
    // invoke: QString buildCpuArchitecture()
    var ret0 = C.C_ZN8QSysInfo20buildCpuArchitectureEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "buildCpuArchitecture", args)
  }

  return
}

// kernelType()
func (this *QSysInfo) Kerneltype_S(args ...interface{}) (ret interface{}) {
  // kernelType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo10kernelTypeEv
    // invoke: QString kernelType()
    var ret0 = C.C_ZN8QSysInfo10kernelTypeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "kernelType", args)
  }

  return
}

// productType()
func (this *QSysInfo) Producttype_S(args ...interface{}) (ret interface{}) {
  // productType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo11productTypeEv
    // invoke: QString productType()
    var ret0 = C.C_ZN8QSysInfo11productTypeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "productType", args)
  }

  return
}

// windowsVersion()
func (this *QSysInfo) Windowsversion_S(args ...interface{}) () {
  // windowsVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

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

  return
}

// currentCpuArchitecture()
func (this *QSysInfo) Currentcpuarchitecture_S(args ...interface{}) (ret interface{}) {
  // currentCpuArchitecture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QSysInfo22currentCpuArchitectureEv
    // invoke: QString currentCpuArchitecture()
    var ret0 = C.C_ZN8QSysInfo22currentCpuArchitectureEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSysInfo", "currentCpuArchitecture", args)
  }

  return
}

// macVersion()
func (this *QSysInfo) Macversion_S(args ...interface{}) () {
  // macVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

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

  return
}

// <= body block end

