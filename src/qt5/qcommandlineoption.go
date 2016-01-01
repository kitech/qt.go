package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qcommandlineoption.h
// dst-file: /src/core/qcommandlineoption.go
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
// class sizeof(QCommandLineOption)=1
type QCommandLineOption struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QCommandLineOption) setValueName(args ...interface{}) () {
  // setValueName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption12setValueNameERK7QString
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setValueName", args)
 }

}


func (this *QCommandLineOption) names(args ...interface{}) () {
  // names()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption5namesEv
  default:
    qtrt.ErrorResolve("QCommandLineOption", "names", args)
 }

}


func NewQCommandLineOption(args ...interface{})() {
}


func (this *QCommandLineOption) setDescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption14setDescriptionERK7QString
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDescription", args)
 }

}


func (this *QCommandLineOption) valueName(args ...interface{}) () {
  // valueName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption9valueNameEv
  default:
    qtrt.ErrorResolve("QCommandLineOption", "valueName", args)
 }

}


func (this *QCommandLineOption) swap(args ...interface{}) () {
  // swap(class QCommandLineOption &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCommandLineOption{}) // "QCommandLineOption &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption4swapERS_
  default:
    qtrt.ErrorResolve("QCommandLineOption", "swap", args)
 }

}


func (this *QCommandLineOption) description(args ...interface{}) () {
  // description()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption11descriptionEv
  default:
    qtrt.ErrorResolve("QCommandLineOption", "description", args)
 }

}


func (this *QCommandLineOption) FreeQCommandLineOption(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QCommandLineOption", "~QCommandLineOption", args)
 }

}


func (this *QCommandLineOption) setDefaultValue(args ...interface{}) () {
  // setDefaultValue(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption15setDefaultValueERK7QString
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDefaultValue", args)
 }

}


func (this *QCommandLineOption) setDefaultValues(args ...interface{}) () {
  // setDefaultValues(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLineOption16setDefaultValuesERK11QStringList
  default:
    qtrt.ErrorResolve("QCommandLineOption", "setDefaultValues", args)
 }

}


func (this *QCommandLineOption) defaultValues(args ...interface{}) () {
  // defaultValues()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLineOption13defaultValuesEv
  default:
    qtrt.ErrorResolve("QCommandLineOption", "defaultValues", args)
 }

}

// <= body block end

