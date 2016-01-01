package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qaccessiblebridge.h
// dst-file: /src/gui/qaccessiblebridge.go
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
// class sizeof(QAccessibleBridgePlugin)=1
type QAccessibleBridgePlugin struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleBridge)=8
type QAccessibleBridge struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func NewQAccessibleBridgePlugin(args ...interface{})() {
}


func (this *QAccessibleBridgePlugin) create(args ...interface{}) () {
  // create(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN23QAccessibleBridgePlugin6createERK7QString
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "create", args)
 }

}


func (this *QAccessibleBridgePlugin) FreeQAccessibleBridgePlugin(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "~QAccessibleBridgePlugin", args)
 }

}


func (this *QAccessibleBridgePlugin) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK23QAccessibleBridgePlugin10metaObjectEv
  default:
    qtrt.ErrorResolve("QAccessibleBridgePlugin", "metaObject", args)
 }

}


func (this *QAccessibleBridge) FreeQAccessibleBridge(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleBridge", "~QAccessibleBridge", args)
 }

}


func (this *QAccessibleBridge) notifyAccessibilityUpdate(args ...interface{}) () {
  // notifyAccessibilityUpdate(class QAccessibleEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleEvent{}) // "QAccessibleEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QAccessibleBridge25notifyAccessibilityUpdateEP16QAccessibleEvent
  default:
    qtrt.ErrorResolve("QAccessibleBridge", "notifyAccessibilityUpdate", args)
 }

}


func (this *QAccessibleBridge) setRootObject(args ...interface{}) () {
  // setRootObject(class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN17QAccessibleBridge13setRootObjectEP20QAccessibleInterface
  default:
    qtrt.ErrorResolve("QAccessibleBridge", "setRootObject", args)
 }

}

// <= body block end

