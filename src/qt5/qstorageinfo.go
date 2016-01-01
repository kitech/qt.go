package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qstorageinfo.h
// dst-file: /src/core/qstorageinfo.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QStorageInfo)=1
type QStorageInfo struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QStorageInfo) bytesFree(args ...interface{}) () {
  // bytesFree()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo9bytesFreeEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesFree", args)
 }

}


func NewQStorageInfo(args ...interface{})() {
}


func (this *QStorageInfo) isRoot(args ...interface{}) () {
  // isRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo6isRootEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "isRoot", args)
 }

}


func (this *QStorageInfo) isReadOnly(args ...interface{}) () {
  // isReadOnly()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo10isReadOnlyEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReadOnly", args)
 }

}


func (this *QStorageInfo) fileSystemType(args ...interface{}) () {
  // fileSystemType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo14fileSystemTypeEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "fileSystemType", args)
 }

}


func (this *QStorageInfo) setPath(args ...interface{}) () {
  // setPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo7setPathERK7QString
  default:
    qtrt.ErrorResolve("QStorageInfo", "setPath", args)
 }

}


func (this *QStorageInfo) mountedVolumes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStorageInfo", "mountedVolumes", args)
 }

}


func (this *QStorageInfo) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo4nameEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "name", args)
 }

}


func (this *QStorageInfo) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo7refreshEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "refresh", args)
 }

}


func (this *QStorageInfo) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo7isValidEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "isValid", args)
 }

}


func (this *QStorageInfo) isReady(args ...interface{}) () {
  // isReady()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo7isReadyEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "isReady", args)
 }

}


func (this *QStorageInfo) bytesTotal(args ...interface{}) () {
  // bytesTotal()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo10bytesTotalEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesTotal", args)
 }

}


func (this *QStorageInfo) rootPath(args ...interface{}) () {
  // rootPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo8rootPathEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "rootPath", args)
 }

}


func (this *QStorageInfo) FreeQStorageInfo(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStorageInfo", "~QStorageInfo", args)
 }

}


func (this *QStorageInfo) bytesAvailable(args ...interface{}) () {
  // bytesAvailable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo14bytesAvailableEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "bytesAvailable", args)
 }

}


func (this *QStorageInfo) root_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QStorageInfo", "root", args)
 }

}


func (this *QStorageInfo) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo6deviceEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "device", args)
 }

}


func (this *QStorageInfo) displayName(args ...interface{}) () {
  // displayName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK12QStorageInfo11displayNameEv
  default:
    qtrt.ErrorResolve("QStorageInfo", "displayName", args)
 }

}


func (this *QStorageInfo) swap(args ...interface{}) () {
  // swap(class QStorageInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStorageInfo{}) // "QStorageInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QStorageInfo4swapERS_
  default:
    qtrt.ErrorResolve("QStorageInfo", "swap", args)
 }

}

// <= body block end

