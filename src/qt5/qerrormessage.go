package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qerrormessage.h
// dst-file: /src/widgets/qerrormessage.go
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
// class sizeof(QErrorMessage)=1
type QErrorMessage struct {
  /*qbase*/ QDialog;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QErrorMessage) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QErrorMessage10metaObjectEv
  default:
    qtrt.ErrorResolve("QErrorMessage", "metaObject", args)
  }

}


func NewQErrorMessage(args ...interface{}) QErrorMessage {
  return QErrorMessage{}
}


func (this *QErrorMessage) qtHandler_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QErrorMessage", "qtHandler", args)
  }

}


func (this *QErrorMessage) showMessage(args ...interface{}) () {
  // showMessage(const class QString &, const class QString &)
  // showMessage(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QErrorMessage11showMessageERK7QStringS2_
  case 1:
    // invoke: _ZN13QErrorMessage11showMessageERK7QString
  default:
    qtrt.ErrorResolve("QErrorMessage", "showMessage", args)
  }

}


func (this *QErrorMessage) FreeQErrorMessage(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QErrorMessage", "~QErrorMessage", args)
  }

}

// <= body block end

