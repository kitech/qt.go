package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
// src-file: /QtWidgets/qwhatsthis.h
// dst-file: /src/widgets/qwhatsthis.go
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
  // proto: static void QWhatsThis::enterWhatsThisMode();
extern void C_ZN10QWhatsThis18enterWhatsThisModeEv(); // 4
  // proto: static void QWhatsThis::leaveWhatsThisMode();
extern void C_ZN10QWhatsThis18leaveWhatsThisModeEv(); // 4
  // proto: static void QWhatsThis::hideText();
extern void C_ZN10QWhatsThis8hideTextEv(); // 4
  // proto: static QAction * QWhatsThis::createAction(QObject * parent);
extern void* C_ZN10QWhatsThis12createActionEP7QObject(void* arg0); // 4
  // proto: static void QWhatsThis::showText(const QPoint & pos, const QString & text, QWidget * w);
extern void C_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2); // 4
  // proto: static bool QWhatsThis::inWhatsThisMode();
extern bool C_ZN10QWhatsThis15inWhatsThisModeEv(); // 4
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

// class sizeof(QWhatsThis)=1
type QWhatsThis struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// enterWhatsThisMode()
func (this *QWhatsThis) Enterwhatsthismode_S(args ...interface{}) () {
  // enterWhatsThisMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QWhatsThis18enterWhatsThisModeEv
    // invoke: void enterWhatsThisMode()
    C.C_ZN10QWhatsThis18enterWhatsThisModeEv()
  default:
    qtrt.ErrorResolve("QWhatsThis", "enterWhatsThisMode", args)
  }

  return
}

// leaveWhatsThisMode()
func (this *QWhatsThis) Leavewhatsthismode_S(args ...interface{}) () {
  // leaveWhatsThisMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QWhatsThis18leaveWhatsThisModeEv
    // invoke: void leaveWhatsThisMode()
    C.C_ZN10QWhatsThis18leaveWhatsThisModeEv()
  default:
    qtrt.ErrorResolve("QWhatsThis", "leaveWhatsThisMode", args)
  }

  return
}

// hideText()
func (this *QWhatsThis) Hidetext_S(args ...interface{}) () {
  // hideText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QWhatsThis8hideTextEv
    // invoke: void hideText()
    C.C_ZN10QWhatsThis8hideTextEv()
  default:
    qtrt.ErrorResolve("QWhatsThis", "hideText", args)
  }

  return
}

// createAction(class QObject *)
func (this *QWhatsThis) Createaction_S(args ...interface{}) (ret interface{}) {
  // createAction(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QWhatsThis12createActionEP7QObject
    // invoke: QAction * createAction(class QObject *)
    var arg0 = args[0].(*QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QWhatsThis12createActionEP7QObject(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWhatsThis", "createAction", args)
  }

  return
}

// showText(const class QPoint &, const class QString &, class QWidget *)
func (this *QWhatsThis) Showtext_S(args ...interface{}) () {
  // showText(const class QPoint &, const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget
    // invoke: void showText(const class QPoint &, const class QString &, class QWidget *)
    var arg0 = args[0].(*QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QWhatsThis", "showText", args)
  }

  return
}

// inWhatsThisMode()
func (this *QWhatsThis) Inwhatsthismode_S(args ...interface{}) (ret interface{}) {
  // inWhatsThisMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QWhatsThis15inWhatsThisModeEv
    // invoke: bool inWhatsThisMode()
    var ret0 = C.C_ZN10QWhatsThis15inWhatsThisModeEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWhatsThis", "inWhatsThisMode", args)
  }

  return
}

// <= body block end

