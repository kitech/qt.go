package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QCommandLinkButton::QCommandLinkButton(const QString & text, const QString & description, QWidget * parent);
extern void* dector_ZN18QCommandLinkButtonC1ERK7QStringS2_P7QWidget(void* arg0, void* arg1, void* arg2);
extern void _ZN18QCommandLinkButtonC1ERK7QStringS2_P7QWidget(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  const QMetaObject * QCommandLinkButton::metaObject();
extern void _ZNK18QCommandLinkButton10metaObjectEv(void* qthis);
  // proto:  void QCommandLinkButton::~QCommandLinkButton();
extern void _ZN18QCommandLinkButtonD0Ev(void* qthis);
  // proto:  void QCommandLinkButton::QCommandLinkButton(const QCommandLinkButton & );
extern void* dector_ZN18QCommandLinkButtonC1ERKS_(void* arg0);
extern void _ZN18QCommandLinkButtonC1ERKS_(void* qthis, void* arg0);
  // proto:  void QCommandLinkButton::QCommandLinkButton(QWidget * parent);
extern void* dector_ZN18QCommandLinkButtonC1EP7QWidget(void* arg0);
extern void _ZN18QCommandLinkButtonC1EP7QWidget(void* qthis, void* arg0);
  // proto:  QString QCommandLinkButton::description();
extern void _ZNK18QCommandLinkButton11descriptionEv(void* qthis);
  // proto:  void QCommandLinkButton::QCommandLinkButton(const QString & text, QWidget * parent);
extern void* dector_ZN18QCommandLinkButtonC1ERK7QStringP7QWidget(void* arg0, void* arg1);
extern void _ZN18QCommandLinkButtonC1ERK7QStringP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  void QCommandLinkButton::setDescription(const QString & description);
extern void _ZN18QCommandLinkButton14setDescriptionERK7QString(void* qthis, void* arg0);
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

// class sizeof(QCommandLinkButton)=1
type QCommandLinkButton struct {
  /*qbase*/ QPushButton;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QCommandLinkButton::QCommandLinkButton(const QString & text, const QString & description, QWidget * parent);
func NewQCommandLinkButton(args ...interface{}) QCommandLinkButton {
  return QCommandLinkButton{}
}

  // proto:  const QMetaObject * QCommandLinkButton::metaObject();
func (this *QCommandLinkButton) metaObject(args ...interface{}) () {
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
    C._ZNK18QCommandLinkButton10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "metaObject", args)
  }

}

  // proto:  void QCommandLinkButton::~QCommandLinkButton();
func (this *QCommandLinkButton) FreeQCommandLinkButton(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "~QCommandLinkButton", args)
  }

}

  // proto:  QString QCommandLinkButton::description();
func (this *QCommandLinkButton) description(args ...interface{}) () {
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
    C._ZNK18QCommandLinkButton11descriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "description", args)
  }

}

  // proto:  void QCommandLinkButton::setDescription(const QString & description);
func (this *QCommandLinkButton) setDescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QCommandLinkButton14setDescriptionERK7QString
    // invoke: void setDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QCommandLinkButton14setDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCommandLinkButton", "setDescription", args)
  }

}

// <= body block end

