package qt5
// auto generated, do not modify.
// created: Sat Jan  2 01:07:50 2016
// src-file: /QtWidgets/qsplashscreen.h
// dst-file: /src/widgets/qsplashscreen.go
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
// class sizeof(QSplashScreen)=1
type QSplashScreen struct {
  /*qbase*/ QWidget;
  qclsinst uint64 /* *mut c_void*/;
//  _messageChanged QSplashScreen_messageChanged_signal;
}


func (this *QSplashScreen) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSplashScreen10metaObjectEv
  default:
    qtrt.ErrorResolve("QSplashScreen", "metaObject", args)
  }

}


func (this *QSplashScreen) FreeQSplashScreen(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QSplashScreen", "~QSplashScreen", args)
  }

}


func (this *QSplashScreen) clearMessage(args ...interface{}) () {
  // clearMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen12clearMessageEv
  default:
    qtrt.ErrorResolve("QSplashScreen", "clearMessage", args)
  }

}


func NewQSplashScreen(args ...interface{}) QSplashScreen {
  return QSplashScreen{}
}


func (this *QSplashScreen) pixmap(args ...interface{}) () {
  // pixmap()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSplashScreen6pixmapEv
  default:
    qtrt.ErrorResolve("QSplashScreen", "pixmap", args)
  }

}


func (this *QSplashScreen) showMessage(args ...interface{}) () {
  // showMessage(const class QString &, int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen11showMessageERK7QStringiRK6QColor
  default:
    qtrt.ErrorResolve("QSplashScreen", "showMessage", args)
  }

}


func (this *QSplashScreen) setPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen9setPixmapERK7QPixmap
  default:
    qtrt.ErrorResolve("QSplashScreen", "setPixmap", args)
  }

}


func (this *QSplashScreen) message(args ...interface{}) () {
  // message()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QSplashScreen7messageEv
  default:
    qtrt.ErrorResolve("QSplashScreen", "message", args)
  }

}


func (this *QSplashScreen) repaint(args ...interface{}) () {
  // repaint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen7repaintEv
  default:
    qtrt.ErrorResolve("QSplashScreen", "repaint", args)
  }

}


func (this *QSplashScreen) finish(args ...interface{}) () {
  // finish(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen6finishEP7QWidget
  default:
    qtrt.ErrorResolve("QSplashScreen", "finish", args)
  }

}

// <= body block end

