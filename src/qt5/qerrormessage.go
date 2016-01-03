package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qerrormessage.h
// dst-file: /src/widgets/qerrormessage.go
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
  // proto:  const QMetaObject * QErrorMessage::metaObject();
extern void _ZNK13QErrorMessage10metaObjectEv(void* qthis);
  // proto:  void QErrorMessage::QErrorMessage(QWidget * parent);
extern void* dector_ZN13QErrorMessageC1EP7QWidget(void* arg0);
extern void _ZN13QErrorMessageC1EP7QWidget(void* qthis, void* arg0);
  // proto: static QErrorMessage * QErrorMessage::qtHandler();
extern void _ZN13QErrorMessage9qtHandlerEv();
  // proto:  void QErrorMessage::showMessage(const QString & message, const QString & type);
extern void _ZN13QErrorMessage11showMessageERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QErrorMessage::showMessage(const QString & message);
extern void _ZN13QErrorMessage11showMessageERK7QString(void* qthis, void* arg0);
  // proto:  void QErrorMessage::~QErrorMessage();
extern void _ZN13QErrorMessageD0Ev(void* qthis);
  // proto:  void QErrorMessage::QErrorMessage(const QErrorMessage & );
extern void* dector_ZN13QErrorMessageC1ERKS_(void* arg0);
extern void _ZN13QErrorMessageC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QErrorMessage)=1
type QErrorMessage struct {
  /*qbase*/ QDialog;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  const QMetaObject * QErrorMessage::metaObject();
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QErrorMessage10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QErrorMessage", "metaObject", args)
  }

}

  // proto:  void QErrorMessage::QErrorMessage(QWidget * parent);
func NewQErrorMessage(args ...interface{}) QErrorMessage {
  return QErrorMessage{}
}

  // proto: static QErrorMessage * QErrorMessage::qtHandler();
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

  // proto:  void QErrorMessage::showMessage(const QString & message, const QString & type);
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
    // invoke: void showMessage(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN13QErrorMessage11showMessageERK7QStringS2_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QErrorMessage11showMessageERK7QString
    // invoke: void showMessage(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QErrorMessage11showMessageERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QErrorMessage", "showMessage", args)
  }

}

  // proto:  void QErrorMessage::~QErrorMessage();
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

