package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qmimedatabase.h
// dst-file: /src/core/qmimedatabase.go
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
// class sizeof(QMimeDatabase)=8
type QMimeDatabase struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QMimeDatabase) mimeTypeForUrl(args ...interface{}) () {
  // mimeTypeForUrl(const class QUrl &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QUrl{}) // "const QUrl &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase14mimeTypeForUrlERK4QUrl
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForUrl", args)
  }

}


func NewQMimeDatabase(args ...interface{}) QMimeDatabase {
  return QMimeDatabase{}
}


func (this *QMimeDatabase) mimeTypeForData(args ...interface{}) () {
  // mimeTypeForData(const class QByteArray &)
  // mimeTypeForData(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase15mimeTypeForDataERK10QByteArray
  case 1:
    // invoke: _ZNK13QMimeDatabase15mimeTypeForDataEP9QIODevice
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForData", args)
  }

}


func (this *QMimeDatabase) mimeTypeForName(args ...interface{}) () {
  // mimeTypeForName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase15mimeTypeForNameERK7QString
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForName", args)
  }

}


func (this *QMimeDatabase) suffixForFileName(args ...interface{}) () {
  // suffixForFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase17suffixForFileNameERK7QString
  default:
    qtrt.ErrorResolve("QMimeDatabase", "suffixForFileName", args)
  }

}


func (this *QMimeDatabase) mimeTypesForFileName(args ...interface{}) () {
  // mimeTypesForFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase20mimeTypesForFileNameERK7QString
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypesForFileName", args)
  }

}


func (this *QMimeDatabase) mimeTypeForFileNameAndData(args ...interface{}) () {
  // mimeTypeForFileNameAndData(const class QString &, class QIODevice *)
  // mimeTypeForFileNameAndData(const class QString &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringP9QIODevice
  case 1:
    // invoke: _ZNK13QMimeDatabase26mimeTypeForFileNameAndDataERK7QStringRK10QByteArray
  default:
    qtrt.ErrorResolve("QMimeDatabase", "mimeTypeForFileNameAndData", args)
  }

}


func (this *QMimeDatabase) FreeQMimeDatabase(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMimeDatabase", "~QMimeDatabase", args)
  }

}


func (this *QMimeDatabase) allMimeTypes(args ...interface{}) () {
  // allMimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QMimeDatabase12allMimeTypesEv
  default:
    qtrt.ErrorResolve("QMimeDatabase", "allMimeTypes", args)
  }

}

// <= body block end

