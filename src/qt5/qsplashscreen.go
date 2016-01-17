package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qsplashscreen.h
// dst-file: /src/widgets/qsplashscreen.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QPixmap QSplashScreen::pixmap();
extern void _ZNK13QSplashScreen6pixmapEv(void* qthis); // 4
  // proto:  void QSplashScreen::repaint();
extern void _ZN13QSplashScreen7repaintEv(void* qthis); // 4
  // proto:  void QSplashScreen::finish(QWidget * w);
extern void _ZN13QSplashScreen6finishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QSplashScreen::metaObject();
extern void _ZNK13QSplashScreen10metaObjectEv(void* qthis); // 4
  // proto:  void QSplashScreen::setPixmap(const QPixmap & pixmap);
extern void _ZN13QSplashScreen9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  void QSplashScreen::clearMessage();
extern void _ZN13QSplashScreen12clearMessageEv(void* qthis); // 4
  // proto:  QString QSplashScreen::message();
extern void _ZNK13QSplashScreen7messageEv(void* qthis); // 4
  // proto:  void QSplashScreen::showMessage(const QString & message, int alignment, const QColor & color);
extern void _ZN13QSplashScreen11showMessageERK7QStringiRK6QColor(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSplashScreen::~QSplashScreen();
extern void _ZN13QSplashScreenD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QSplashScreen)=1
type QSplashScreen struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _messageChanged QSplashScreen_messageChanged_signal;
}

// pixmap()
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
    // invoke: const QPixmap pixmap()
    C._ZNK13QSplashScreen6pixmapEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "pixmap", args)
  }

}

// repaint()
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
    // invoke: void repaint()
    C._ZN13QSplashScreen7repaintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "repaint", args)
  }

}

// finish(class QWidget *)
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
    // invoke: void finish(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QSplashScreen6finishEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplashScreen", "finish", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QSplashScreen10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "metaObject", args)
  }

}

// setPixmap(const class QPixmap &)
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
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(QPixmap).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QSplashScreen9setPixmapERK7QPixmap(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplashScreen", "setPixmap", args)
  }

}

// clearMessage()
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
    // invoke: void clearMessage()
    C._ZN13QSplashScreen12clearMessageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "clearMessage", args)
  }

}

// message()
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
    // invoke: QString message()
    C._ZNK13QSplashScreen7messageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "message", args)
  }

}

// showMessage(const class QString &, int, const class QColor &)
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
    // invoke: void showMessage(const class QString &, int, const class QColor &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QColor).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN13QSplashScreen11showMessageERK7QStringiRK6QColor(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSplashScreen", "showMessage", args)
  }

}

// ~QSplashScreen()
func (this *QSplashScreen) FreeQSplashScreen(args ...interface{}) () {
  // ~QSplashScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreenD0Ev
    // invoke: void ~QSplashScreen()
    C._ZN13QSplashScreenD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "~QSplashScreen", args)
  }

}

// <= body block end

