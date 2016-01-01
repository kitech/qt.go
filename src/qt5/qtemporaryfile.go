package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qtemporaryfile.h
// dst-file: /src/core/qtemporaryfile.go
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
// class sizeof(QTemporaryFile)=1
type QTemporaryFile struct {
  /*qbase*/ QFile;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTemporaryFile) autoRemove(args ...interface{}) () {
  // autoRemove()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile10autoRemoveEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "autoRemove", args)
  }

}


func (this *QTemporaryFile) createLocalFile_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryFile", "createLocalFile", args)
  }

}


func NewQTemporaryFile(args ...interface{}) QTemporaryFile {
  return QTemporaryFile{}
}


func (this *QTemporaryFile) FreeQTemporaryFile(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryFile", "~QTemporaryFile", args)
  }

}


func (this *QTemporaryFile) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile10metaObjectEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "metaObject", args)
  }

}


func (this *QTemporaryFile) setAutoRemove(args ...interface{}) () {
  // setAutoRemove(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile13setAutoRemoveEb
  default:
    qtrt.ErrorResolve("QTemporaryFile", "setAutoRemove", args)
  }

}


func (this *QTemporaryFile) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile8fileNameEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "fileName", args)
  }

}


func (this *QTemporaryFile) fileTemplate(args ...interface{}) () {
  // fileTemplate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QTemporaryFile12fileTemplateEv
  default:
    qtrt.ErrorResolve("QTemporaryFile", "fileTemplate", args)
  }

}


func (this *QTemporaryFile) createNativeFile_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QTemporaryFile", "createNativeFile", args)
  }

}


func (this *QTemporaryFile) open(args ...interface{}) () {
  // open()
  // open(OpenMode)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "OpenMode"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile4openEv
  case 1:
    // invoke: _ZN14QTemporaryFile4openE6QFlagsIN9QIODevice12OpenModeFlagEE
  default:
    qtrt.ErrorResolve("QTemporaryFile", "open", args)
  }

}


func (this *QTemporaryFile) setFileTemplate(args ...interface{}) () {
  // setFileTemplate(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QTemporaryFile15setFileTemplateERK7QString
  default:
    qtrt.ErrorResolve("QTemporaryFile", "setFileTemplate", args)
  }

}

// <= body block end

