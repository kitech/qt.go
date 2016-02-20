package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
  // proto:  const QMetaObject * QErrorMessage::metaObject();
extern void C_ZNK13QErrorMessage10metaObjectEv(void* qthis); // 4
  // proto:  void QErrorMessage::QErrorMessage(QWidget * parent);
extern void* C_ZN13QErrorMessageC2EP7QWidget(void* arg0); // 3
  // proto:  void QErrorMessage::~QErrorMessage();
extern void C_ZN13QErrorMessageD2Ev(void* qthis); // 4
  // proto: static QErrorMessage * QErrorMessage::qtHandler();
extern void* C_ZN13QErrorMessage9qtHandlerEv(); // 4
  // proto:  void QErrorMessage::showMessage(const QString & message, const QString & type);
extern void C_ZN13QErrorMessage11showMessageERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QErrorMessage::showMessage(const QString & message);
extern void C_ZN13QErrorMessage11showMessageERK7QString(void* qthis, void* arg0); // 4
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
}

// class sizeof(QErrorMessage)=1
type QErrorMessage struct {
  /*qbase*/ QDialog;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// metaObject()
func (this *QErrorMessage) Metaobject(args ...interface{}) () {
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
    C.C_ZNK13QErrorMessage10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QErrorMessage", "metaObject", args)
  }

  return
}

// QErrorMessage(class QWidget *)
func NewQErrorMessage(args ...interface{}) *QErrorMessage {
  // QErrorMessage(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QErrorMessageC1EP7QWidget
    // invoke: void QErrorMessage(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QErrorMessageC2EP7QWidget(arg0)
    return &QErrorMessage{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QErrorMessage", "QErrorMessage", args)
  }

  return nil // QErrorMessage{Qclsinst:qthis}
}

// ~QErrorMessage()
func (this *QErrorMessage) Freeqerrormessage(args ...interface{}) () {
  // ~QErrorMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QErrorMessageD0Ev
    // invoke: void ~QErrorMessage()
    C.C_ZN13QErrorMessageD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QErrorMessage", "~QErrorMessage", args)
  }

  return
}

// qtHandler()
func (this *QErrorMessage) Qthandler_S(args ...interface{}) (ret interface{}) {
  // qtHandler()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QErrorMessage9qtHandlerEv
    // invoke: QErrorMessage * qtHandler()
    var ret0 = C.C_ZN13QErrorMessage9qtHandlerEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QErrorMessage{}) // "QErrorMessage *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QErrorMessage", "qtHandler", args)
  }

  return
}

// showMessage(const class QString &, const class QString &)
func (this *QErrorMessage) Showmessage(args ...interface{}) () {
  // showMessage(const class QString &, const class QString &)
  // showMessage(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QErrorMessage11showMessageERK7QStringS2_
    // invoke: void showMessage(const class QString &, const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QErrorMessage11showMessageERK7QStringS2_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QErrorMessage11showMessageERK7QString
    // invoke: void showMessage(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QErrorMessage11showMessageERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QErrorMessage", "showMessage", args)
  }

  return
}

// <= body block end

