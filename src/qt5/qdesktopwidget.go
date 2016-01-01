package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtWidgets/qdesktopwidget.h
// dst-file: /src/widgets/qdesktopwidget.go
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
// class sizeof(QDesktopWidget)=1
type QDesktopWidget struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _screenCountChanged QDesktopWidget_screenCountChanged_signal;
//  _resized QDesktopWidget_resized_signal;
//  _workAreaResized QDesktopWidget_workAreaResized_signal;
}


func (this *QDesktopWidget) screenGeometry(args ...interface{}) () {
  // screenGeometry(const class QPoint &)
  // screenGeometry(const class QWidget *)
  // screenGeometry(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget14screenGeometryERK6QPoint
  case 1:
    // invoke: _ZNK14QDesktopWidget14screenGeometryEPK7QWidget
  case 2:
    // invoke: _ZNK14QDesktopWidget14screenGeometryEi
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenGeometry", args)
 }

}


func (this *QDesktopWidget) screen(args ...interface{}) () {
  // screen(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN14QDesktopWidget6screenEi
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screen", args)
 }

}


func (this *QDesktopWidget) numScreens(args ...interface{}) () {
  // numScreens()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget10numScreensEv
  default:
    qtrt.ErrorResolve("QDesktopWidget", "numScreens", args)
 }

}


func (this *QDesktopWidget) FreeQDesktopWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QDesktopWidget", "~QDesktopWidget", args)
 }

}


func (this *QDesktopWidget) availableGeometry(args ...interface{}) () {
  // availableGeometry(const class QWidget *)
  // availableGeometry(const class QPoint &)
  // availableGeometry(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget17availableGeometryEPK7QWidget
  case 1:
    // invoke: _ZNK14QDesktopWidget17availableGeometryERK6QPoint
  case 2:
    // invoke: _ZNK14QDesktopWidget17availableGeometryEi
  default:
    qtrt.ErrorResolve("QDesktopWidget", "availableGeometry", args)
 }

}


func NewQDesktopWidget(args ...interface{})() {
}


func (this *QDesktopWidget) screenNumber(args ...interface{}) () {
  // screenNumber(const class QPoint &)
  // screenNumber(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget12screenNumberERK6QPoint
  case 1:
    // invoke: _ZNK14QDesktopWidget12screenNumberEPK7QWidget
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenNumber", args)
 }

}


func (this *QDesktopWidget) screenCount(args ...interface{}) () {
  // screenCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget11screenCountEv
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenCount", args)
 }

}


func (this *QDesktopWidget) isVirtualDesktop(args ...interface{}) () {
  // isVirtualDesktop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget16isVirtualDesktopEv
  default:
    qtrt.ErrorResolve("QDesktopWidget", "isVirtualDesktop", args)
 }

}


func (this *QDesktopWidget) primaryScreen(args ...interface{}) () {
  // primaryScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget13primaryScreenEv
  default:
    qtrt.ErrorResolve("QDesktopWidget", "primaryScreen", args)
 }

}


func (this *QDesktopWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget10metaObjectEv
  default:
    qtrt.ErrorResolve("QDesktopWidget", "metaObject", args)
 }

}

// <= body block end

