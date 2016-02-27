package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qcommandlinkbutton.h
// dst-file: /src/widgets/qcommandlinkbutton.go
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
  // proto:  void QCommandLinkButton::setDescription(const QString & description);
extern void C_ZN18QCommandLinkButton14setDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCommandLinkButton::~QCommandLinkButton();
extern void C_ZN18QCommandLinkButtonD2Ev(void* qthis); // 4
  // proto:  QString QCommandLinkButton::description();
extern void* C_ZNK18QCommandLinkButton11descriptionEv(void* qthis); // 4
  // proto:  void QCommandLinkButton::QCommandLinkButton(QWidget * parent);
extern void* C_ZN18QCommandLinkButtonC2EP7QWidget(void* arg0); // 3
  // proto:  void QCommandLinkButton::QCommandLinkButton(const QString & text, const QString & description, QWidget * parent);
extern void* C_ZN18QCommandLinkButtonC2ERK7QStringS2_P7QWidget(void* arg0, void* arg1, void* arg2); // 3
  // proto:  void QCommandLinkButton::QCommandLinkButton(const QString & text, QWidget * parent);
extern void* C_ZN18QCommandLinkButtonC2ERK7QStringP7QWidget(void* arg0, void* arg1); // 3
  // proto:  const QMetaObject * QCommandLinkButton::metaObject();
extern void C_ZNK18QCommandLinkButton10metaObjectEv(void* qthis); // 4
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

// class sizeof(QCommandLinkButton)=1
type QCommandLinkButton struct {
  /*qbase*/ QPushButton;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setDescription(const class QString &)
func (this *QCommandLinkButton) SetDescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLinkButton14setDescriptionERK7QString
    // invoke: void setDescription(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QCommandLinkButton14setDescriptionERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "setDescription", args)
  }

  return
}

// ~QCommandLinkButton()
func (this *QCommandLinkButton) Free(args ...interface{}) () {
  // ~QCommandLinkButton()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLinkButtonD0Ev
    // invoke: void ~QCommandLinkButton()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN18QCommandLinkButtonD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "~QCommandLinkButton", args)
  }

  return
}

// description()
func (this *QCommandLinkButton) Description(args ...interface{}) (ret interface{}) {
  // description()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLinkButton11descriptionEv
    // invoke: QString description()
    var ret0 = C.C_ZNK18QCommandLinkButton11descriptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "description", args)
  }

  return
}

// QCommandLinkButton(class QWidget *)
func GcfreeQCommandLinkButton(this *QCommandLinkButton) {
  qtrt.UniverseFree(this)
}
func NewQCommandLinkButton(args ...interface{}) *QCommandLinkButton {
  // QCommandLinkButton(class QWidget *)
  // QCommandLinkButton(const class QString &, const class QString &, class QWidget *)
  // QCommandLinkButton(const class QString &, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLinkButtonC1EP7QWidget
    // invoke: void QCommandLinkButton(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLinkButtonC2EP7QWidget(arg0)
    this := &QCommandLinkButton{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCommandLinkButton)
    return this
  case 1:
    // invoke: _ZN18QCommandLinkButtonC1ERK7QStringS2_P7QWidget
    // invoke: void QCommandLinkButton(const class QString &, const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLinkButtonC2ERK7QStringS2_P7QWidget(arg0, arg1, arg2)
    this := &QCommandLinkButton{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCommandLinkButton)
    return this
  case 2:
    // invoke: _ZN18QCommandLinkButtonC1ERK7QStringP7QWidget
    // invoke: void QCommandLinkButton(const class QString &, class QWidget *)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QCommandLinkButtonC2ERK7QStringP7QWidget(arg0, arg1)
    this := &QCommandLinkButton{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQCommandLinkButton)
    return this
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "QCommandLinkButton", args)
  }

  return nil // QCommandLinkButton{Qclsinst:qthis}
}

// metaObject()
func (this *QCommandLinkButton) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QCommandLinkButton10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK18QCommandLinkButton10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "metaObject", args)
  }

  return
}

// <= body block end

