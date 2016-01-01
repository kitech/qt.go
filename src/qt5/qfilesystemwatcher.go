package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qfilesystemwatcher.h
// dst-file: /src/core/qfilesystemwatcher.go
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
// class sizeof(QFileSystemWatcher)=1
type QFileSystemWatcher struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _directoryChanged QFileSystemWatcher_directoryChanged_signal;
//  _fileChanged QFileSystemWatcher_fileChanged_signal;
}


func (this *QFileSystemWatcher) FreeQFileSystemWatcher(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "~QFileSystemWatcher", args)
 }

}


func (this *QFileSystemWatcher) removePath(args ...interface{}) () {
  // removePath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher10removePathERK7QString
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "removePath", args)
 }

}


func (this *QFileSystemWatcher) directories(args ...interface{}) () {
  // directories()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFileSystemWatcher11directoriesEv
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "directories", args)
 }

}


func (this *QFileSystemWatcher) files(args ...interface{}) () {
  // files()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFileSystemWatcher5filesEv
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "files", args)
 }

}


func (this *QFileSystemWatcher) addPaths(args ...interface{}) () {
  // addPaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher8addPathsERK11QStringList
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "addPaths", args)
 }

}


func (this *QFileSystemWatcher) removePaths(args ...interface{}) () {
  // removePaths(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher11removePathsERK11QStringList
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "removePaths", args)
 }

}


func NewQFileSystemWatcher(args ...interface{})() {
}


func (this *QFileSystemWatcher) addPath(args ...interface{}) () {
  // addPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QFileSystemWatcher7addPathERK7QString
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "addPath", args)
 }

}


func (this *QFileSystemWatcher) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QFileSystemWatcher10metaObjectEv
  default:
    qtrt.ErrorResolve("QFileSystemWatcher", "metaObject", args)
 }

}

// <= body block end

