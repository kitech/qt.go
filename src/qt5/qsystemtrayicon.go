package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qsystemtrayicon.h
// dst-file: /src/widgets/qsystemtrayicon.go
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
// class sizeof(QSystemTrayIcon)=1
type QSystemTrayIcon struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _activated QSystemTrayIcon_activated_signal;
//  _messageClicked QSystemTrayIcon_messageClicked_signal;
}


func (this *QSystemTrayIcon) FreeQSystemTrayIcon(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "~QSystemTrayIcon", args)
 }

}


func (this *QSystemTrayIcon) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon10setVisibleEb
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setVisible", args)
 }

}


func (this *QSystemTrayIcon) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon7toolTipEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "toolTip", args)
 }

}


func NewQSystemTrayIcon(args ...interface{})() {
}


func (this *QSystemTrayIcon) hide(args ...interface{}) () {
  // hide()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon4hideEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "hide", args)
 }

}


func (this *QSystemTrayIcon) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon10metaObjectEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "metaObject", args)
 }

}


func (this *QSystemTrayIcon) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon7setIconERK5QIcon
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setIcon", args)
 }

}


func (this *QSystemTrayIcon) isVisible(args ...interface{}) () {
  // isVisible()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon9isVisibleEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "isVisible", args)
 }

}


func (this *QSystemTrayIcon) show(args ...interface{}) () {
  // show()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon4showEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "show", args)
 }

}


func (this *QSystemTrayIcon) supportsMessages_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "supportsMessages", args)
 }

}


func (this *QSystemTrayIcon) setContextMenu(args ...interface{}) () {
  // setContextMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon14setContextMenuEP5QMenu
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setContextMenu", args)
 }

}


func (this *QSystemTrayIcon) geometry(args ...interface{}) () {
  // geometry()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon8geometryEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "geometry", args)
 }

}


func (this *QSystemTrayIcon) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN15QSystemTrayIcon10setToolTipERK7QString
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "setToolTip", args)
 }

}


func (this *QSystemTrayIcon) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon4iconEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "icon", args)
 }

}


func (this *QSystemTrayIcon) contextMenu(args ...interface{}) () {
  // contextMenu()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK15QSystemTrayIcon11contextMenuEv
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "contextMenu", args)
 }

}


func (this *QSystemTrayIcon) isSystemTrayAvailable_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QSystemTrayIcon", "isSystemTrayAvailable", args)
 }

}

// <= body block end

