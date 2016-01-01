package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtCore/qpropertyanimation.h
// dst-file: /src/core/qpropertyanimation.go
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
// class sizeof(QPropertyAnimation)=1
type QPropertyAnimation struct {
  /*qbase*/ QVariantAnimation;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QPropertyAnimation) propertyName(args ...interface{}) () {
  // propertyName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPropertyAnimation12propertyNameEv
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "propertyName", args)
  }

}


func (this *QPropertyAnimation) setTargetObject(args ...interface{}) () {
  // setTargetObject(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPropertyAnimation15setTargetObjectEP7QObject
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "setTargetObject", args)
  }

}


func NewQPropertyAnimation(args ...interface{}) QPropertyAnimation {
  return QPropertyAnimation{}
}


func (this *QPropertyAnimation) FreeQPropertyAnimation(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "~QPropertyAnimation", args)
  }

}


func (this *QPropertyAnimation) targetObject(args ...interface{}) () {
  // targetObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPropertyAnimation12targetObjectEv
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "targetObject", args)
  }

}


func (this *QPropertyAnimation) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QPropertyAnimation10metaObjectEv
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "metaObject", args)
  }

}


func (this *QPropertyAnimation) setPropertyName(args ...interface{}) () {
  // setPropertyName(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QPropertyAnimation15setPropertyNameERK10QByteArray
  default:
    qtrt.ErrorResolve("QPropertyAnimation", "setPropertyName", args)
  }

}

// <= body block end

