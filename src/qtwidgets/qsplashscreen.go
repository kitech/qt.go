package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  const QPixmap QSplashScreen::pixmap();
extern void* C_ZNK13QSplashScreen6pixmapEv(void* qthis); // 4
  // proto:  void QSplashScreen::repaint();
extern void C_ZN13QSplashScreen7repaintEv(void* qthis); // 4
  // proto:  void QSplashScreen::finish(QWidget * w);
extern void C_ZN13QSplashScreen6finishEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QSplashScreen::metaObject();
extern void C_ZNK13QSplashScreen10metaObjectEv(void* qthis); // 4
  // proto:  void QSplashScreen::setPixmap(const QPixmap & pixmap);
extern void C_ZN13QSplashScreen9setPixmapERK7QPixmap(void* qthis, void* arg0); // 4
  // proto:  void QSplashScreen::clearMessage();
extern void C_ZN13QSplashScreen12clearMessageEv(void* qthis); // 4
  // proto:  QString QSplashScreen::message();
extern void* C_ZNK13QSplashScreen7messageEv(void* qthis); // 4
  // proto:  void QSplashScreen::showMessage(const QString & message, int alignment, const QColor & color);
extern void C_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QSplashScreen::~QSplashScreen();
extern void C_ZN13QSplashScreenD2Ev(void* qthis); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QSplashScreen)=1
type QSplashScreen struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _messageChanged QSplashScreen_messageChanged_signal;
}

// pixmap()
func (this *QSplashScreen) Pixmap(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSplashScreen6pixmapEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplashScreen", "pixmap", args)
  }

  return
}

// repaint()
func (this *QSplashScreen) Repaint(args ...interface{}) () {
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
    C.C_ZN13QSplashScreen7repaintEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "repaint", args)
  }

  return
}

// finish(class QWidget *)
func (this *QSplashScreen) Finish(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QSplashScreen6finishEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplashScreen", "finish", args)
  }

  return
}

// metaObject()
func (this *QSplashScreen) MetaObject(args ...interface{}) () {
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
    C.C_ZNK13QSplashScreen10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "metaObject", args)
  }

  return
}

// setPixmap(const class QPixmap &)
func (this *QSplashScreen) SetPixmap(args ...interface{}) () {
  // setPixmap(const class QPixmap &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QPixmap{}) // "const QPixmap &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen9setPixmapERK7QPixmap
    // invoke: void setPixmap(const class QPixmap &)
    var arg0 = args[0].(*qtgui.QPixmap).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QSplashScreen9setPixmapERK7QPixmap(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QSplashScreen", "setPixmap", args)
  }

  return
}

// clearMessage()
func (this *QSplashScreen) ClearMessage(args ...interface{}) () {
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
    C.C_ZN13QSplashScreen12clearMessageEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QSplashScreen", "clearMessage", args)
  }

  return
}

// message()
func (this *QSplashScreen) Message(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QSplashScreen7messageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QSplashScreen", "message", args)
  }

  return
}

// showMessage(const class QString &, int, const class QColor &)
func (this *QSplashScreen) ShowMessage(args ...interface{}) () {
  // showMessage(const class QString &, int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QSplashScreen11showMessageERK7QStringiRK6QColor
    // invoke: void showMessage(const class QString &, int, const class QColor &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QSplashScreen11showMessageERK7QStringiRK6QColor(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QSplashScreen", "showMessage", args)
  }

  return
}

// ~QSplashScreen()
func (this *QSplashScreen) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QSplashScreenD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QSplashScreen", "~QSplashScreen", args)
  }

  return
}

// <= body block end

