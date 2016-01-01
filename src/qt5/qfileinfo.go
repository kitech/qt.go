package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qfileinfo.h
// dst-file: /src/core/qfileinfo.go
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
// class sizeof(QFileInfo)=1
type QFileInfo struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QFileInfo) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8isHiddenEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isHidden", args)
 }

}


func NewQFileInfo(args ...interface{})() {
}


func (this *QFileInfo) completeSuffix(args ...interface{}) () {
  // completeSuffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo14completeSuffixEv
  default:
    qtrt.ErrorResolve("QFileInfo", "completeSuffix", args)
 }

}


func (this *QFileInfo) canonicalPath(args ...interface{}) () {
  // canonicalPath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo13canonicalPathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalPath", args)
 }

}


func (this *QFileInfo) exists_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFileInfo", "exists", args)
 }

}


func (this *QFileInfo) makeAbsolute(args ...interface{}) () {
  // makeAbsolute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo12makeAbsoluteEv
  default:
    qtrt.ErrorResolve("QFileInfo", "makeAbsolute", args)
 }

}


func (this *QFileInfo) isRoot(args ...interface{}) () {
  // isRoot()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo6isRootEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isRoot", args)
 }

}


func (this *QFileInfo) canonicalFilePath(args ...interface{}) () {
  // canonicalFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo17canonicalFilePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "canonicalFilePath", args)
 }

}


func (this *QFileInfo) isDir(args ...interface{}) () {
  // isDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo5isDirEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isDir", args)
 }

}


func (this *QFileInfo) symLinkTarget(args ...interface{}) () {
  // symLinkTarget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo13symLinkTargetEv
  default:
    qtrt.ErrorResolve("QFileInfo", "symLinkTarget", args)
 }

}


func (this *QFileInfo) setFile(args ...interface{}) () {
  // setFile(const class QString &)
  // setFile(const class QFile &)
  // setFile(const class QDir &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QFile{}) // "const QFile &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QDir{}) // "const QDir &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo7setFileERK7QString
  case 1:
    // invoke: _ZN9QFileInfo7setFileERK5QFile
  case 2:
    // invoke: _ZN9QFileInfo7setFileERK4QDirRK7QString
  default:
    qtrt.ErrorResolve("QFileInfo", "setFile", args)
 }

}


func (this *QFileInfo) isFile(args ...interface{}) () {
  // isFile()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo6isFileEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isFile", args)
 }

}


func (this *QFileInfo) ownerId(args ...interface{}) () {
  // ownerId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7ownerIdEv
  default:
    qtrt.ErrorResolve("QFileInfo", "ownerId", args)
 }

}


func (this *QFileInfo) readLink(args ...interface{}) () {
  // readLink()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8readLinkEv
  default:
    qtrt.ErrorResolve("QFileInfo", "readLink", args)
 }

}


func (this *QFileInfo) filePath(args ...interface{}) () {
  // filePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8filePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "filePath", args)
 }

}


func (this *QFileInfo) isSymLink(args ...interface{}) () {
  // isSymLink()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo9isSymLinkEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isSymLink", args)
 }

}


func (this *QFileInfo) lastRead(args ...interface{}) () {
  // lastRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8lastReadEv
  default:
    qtrt.ErrorResolve("QFileInfo", "lastRead", args)
 }

}


func (this *QFileInfo) refresh(args ...interface{}) () {
  // refresh()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo7refreshEv
  default:
    qtrt.ErrorResolve("QFileInfo", "refresh", args)
 }

}


func (this *QFileInfo) path(args ...interface{}) () {
  // path()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo4pathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "path", args)
 }

}


func (this *QFileInfo) absoluteDir(args ...interface{}) () {
  // absoluteDir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo11absoluteDirEv
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteDir", args)
 }

}


func (this *QFileInfo) isBundle(args ...interface{}) () {
  // isBundle()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8isBundleEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isBundle", args)
 }

}


func (this *QFileInfo) isRelative(args ...interface{}) () {
  // isRelative()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isRelativeEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isRelative", args)
 }

}


func (this *QFileInfo) setCaching(args ...interface{}) () {
  // setCaching(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo10setCachingEb
  default:
    qtrt.ErrorResolve("QFileInfo", "setCaching", args)
 }

}


func (this *QFileInfo) created(args ...interface{}) () {
  // created()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7createdEv
  default:
    qtrt.ErrorResolve("QFileInfo", "created", args)
 }

}


func (this *QFileInfo) caching(args ...interface{}) () {
  // caching()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7cachingEv
  default:
    qtrt.ErrorResolve("QFileInfo", "caching", args)
 }

}


func (this *QFileInfo) FreeQFileInfo(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFileInfo", "~QFileInfo", args)
 }

}


func (this *QFileInfo) completeBaseName(args ...interface{}) () {
  // completeBaseName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo16completeBaseNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "completeBaseName", args)
 }

}


func (this *QFileInfo) baseName(args ...interface{}) () {
  // baseName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8baseNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "baseName", args)
 }

}


func (this *QFileInfo) isExecutable(args ...interface{}) () {
  // isExecutable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12isExecutableEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isExecutable", args)
 }

}


func (this *QFileInfo) bundleName(args ...interface{}) () {
  // bundleName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10bundleNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "bundleName", args)
 }

}


func (this *QFileInfo) groupId(args ...interface{}) () {
  // groupId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo7groupIdEv
  default:
    qtrt.ErrorResolve("QFileInfo", "groupId", args)
 }

}


func (this *QFileInfo) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo8fileNameEv
  default:
    qtrt.ErrorResolve("QFileInfo", "fileName", args)
 }

}


func (this *QFileInfo) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo4sizeEv
  default:
    qtrt.ErrorResolve("QFileInfo", "size", args)
 }

}


func (this *QFileInfo) absoluteFilePath(args ...interface{}) () {
  // absoluteFilePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo16absoluteFilePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "absoluteFilePath", args)
 }

}


func (this *QFileInfo) suffix(args ...interface{}) () {
  // suffix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo6suffixEv
  default:
    qtrt.ErrorResolve("QFileInfo", "suffix", args)
 }

}


func (this *QFileInfo) group(args ...interface{}) () {
  // group()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo5groupEv
  default:
    qtrt.ErrorResolve("QFileInfo", "group", args)
 }

}


func (this *QFileInfo) isAbsolute(args ...interface{}) () {
  // isAbsolute()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isAbsoluteEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isAbsolute", args)
 }

}


func (this *QFileInfo) isNativePath(args ...interface{}) () {
  // isNativePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12isNativePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isNativePath", args)
 }

}


func (this *QFileInfo) isWritable(args ...interface{}) () {
  // isWritable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isWritableEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isWritable", args)
 }

}


func (this *QFileInfo) owner(args ...interface{}) () {
  // owner()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo5ownerEv
  default:
    qtrt.ErrorResolve("QFileInfo", "owner", args)
 }

}


func (this *QFileInfo) isReadable(args ...interface{}) () {
  // isReadable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo10isReadableEv
  default:
    qtrt.ErrorResolve("QFileInfo", "isReadable", args)
 }

}


func (this *QFileInfo) dir(args ...interface{}) () {
  // dir()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo3dirEv
  default:
    qtrt.ErrorResolve("QFileInfo", "dir", args)
 }

}


func (this *QFileInfo) swap(args ...interface{}) () {
  // swap(class QFileInfo &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFileInfo{}) // "QFileInfo &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo4swapERS_
  default:
    qtrt.ErrorResolve("QFileInfo", "swap", args)
 }

}


func (this *QFileInfo) exists(args ...interface{}) () {
  // exists(const class QString &)
  // exists()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN9QFileInfo6existsERK7QString
  case 1:
    // invoke: _ZNK9QFileInfo6existsEv
  default:
    qtrt.ErrorResolve("QFileInfo", "exists", args)
 }

}


func (this *QFileInfo) lastModified(args ...interface{}) () {
  // lastModified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12lastModifiedEv
  default:
    qtrt.ErrorResolve("QFileInfo", "lastModified", args)
 }

}


func (this *QFileInfo) absolutePath(args ...interface{}) () {
  // absolutePath()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK9QFileInfo12absolutePathEv
  default:
    qtrt.ErrorResolve("QFileInfo", "absolutePath", args)
 }

}

// <= body block end

