package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static void QWhatsThis::hideText();
extern void _ZN10QWhatsThis8hideTextEv();
  // proto: static void QWhatsThis::enterWhatsThisMode();
extern void _ZN10QWhatsThis18enterWhatsThisModeEv();
  // proto: static bool QWhatsThis::inWhatsThisMode();
extern void _ZN10QWhatsThis15inWhatsThisModeEv();
  // proto: static void QWhatsThis::leaveWhatsThisMode();
extern void _ZN10QWhatsThis18leaveWhatsThisModeEv();
  // proto:  void QWhatsThis::QWhatsThis();
extern void* dector_ZN10QWhatsThisC1Ev();
extern void _ZN10QWhatsThisC1Ev(void* qthis);
  // proto: static void QWhatsThis::showText(const QPoint & pos, const QString & text, QWidget * w);
extern void _ZN10QWhatsThis8showTextERK6QPointRK7QStringP7QWidget(void* arg0, void* arg1, void* arg2);
  // proto: static QAction * QWhatsThis::createAction(QObject * parent);
extern void _ZN10QWhatsThis12createActionEP7QObject(void* arg0);
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
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto: static void QWhatsThis::hideText();
func (this *QWhatsThis) hideText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThis", "hideText", args)
  }

}

  // proto: static void QWhatsThis::enterWhatsThisMode();
func (this *QWhatsThis) enterWhatsThisMode_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThis", "enterWhatsThisMode", args)
  }

}

  // proto: static bool QWhatsThis::inWhatsThisMode();
func (this *QWhatsThis) inWhatsThisMode_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThis", "inWhatsThisMode", args)
  }

}

  // proto: static void QWhatsThis::leaveWhatsThisMode();
func (this *QWhatsThis) leaveWhatsThisMode_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThis", "leaveWhatsThisMode", args)
  }

}

  // proto:  void QWhatsThis::QWhatsThis();
func NewQWhatsThis(args ...interface{}) QWhatsThis {
  return QWhatsThis{}
}

  // proto: static void QWhatsThis::showText(const QPoint & pos, const QString & text, QWidget * w);
func (this *QWhatsThis) showText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThis", "showText", args)
  }

}

  // proto: static QAction * QWhatsThis::createAction(QObject * parent);
func (this *QWhatsThis) createAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWhatsThis", "createAction", args)
  }

}

// <= body block end

