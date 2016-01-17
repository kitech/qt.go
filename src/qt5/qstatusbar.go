package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
// src-file: /QtWidgets/qstatusbar.h
// dst-file: /src/widgets/qstatusbar.go
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
  // proto:  void QStatusBar::~QStatusBar();
extern void _ZN10QStatusBarD2Ev(void* qthis); // 4
  // proto:  void QStatusBar::removeWidget(QWidget * widget);
extern void _ZN10QStatusBar12removeWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStatusBar::QStatusBar(QWidget * parent);
extern void _ZN10QStatusBarC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  void QStatusBar::clearMessage();
extern void _ZN10QStatusBar12clearMessageEv(void* qthis); // 4
  // proto:  int QStatusBar::insertWidget(int index, QWidget * widget, int stretch);
extern void _ZN10QStatusBar12insertWidgetEiP7QWidgeti(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QStatusBar::addPermanentWidget(QWidget * widget, int stretch);
extern void _ZN10QStatusBar18addPermanentWidgetEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QString QStatusBar::currentMessage();
extern void _ZNK10QStatusBar14currentMessageEv(void* qthis); // 4
  // proto:  void QStatusBar::addWidget(QWidget * widget, int stretch);
extern void _ZN10QStatusBar9addWidgetEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QStatusBar::setSizeGripEnabled(bool );
extern void _ZN10QStatusBar18setSizeGripEnabledEb(void* qthis, bool arg0); // 4
  // proto:  int QStatusBar::insertPermanentWidget(int index, QWidget * widget, int stretch);
extern void _ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  const QMetaObject * QStatusBar::metaObject();
extern void _ZNK10QStatusBar10metaObjectEv(void* qthis); // 4
  // proto:  bool QStatusBar::isSizeGripEnabled();
extern void _ZNK10QStatusBar17isSizeGripEnabledEv(void* qthis); // 4
  // proto:  void QStatusBar::showMessage(const QString & text, int timeout);
extern void _ZN10QStatusBar11showMessageERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
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

// class sizeof(QStatusBar)=1
type QStatusBar struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _messageChanged QStatusBar_messageChanged_signal;
}

// ~QStatusBar()
func (this *QStatusBar) FreeQStatusBar(args ...interface{}) () {
  // ~QStatusBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBarD0Ev
    // invoke: void ~QStatusBar()
    C._ZN10QStatusBarD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "~QStatusBar", args)
  }

}

// removeWidget(class QWidget *)
func (this *QStatusBar) removeWidget(args ...interface{}) () {
  // removeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar12removeWidgetEP7QWidget
    // invoke: void removeWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QStatusBar12removeWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStatusBar", "removeWidget", args)
  }

}

// QStatusBar(class QWidget *)
func NewQStatusBar(args ...interface{}) QStatusBar {
  // QStatusBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBarC1EP7QWidget
    // invoke: void QStatusBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QStatusBarC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QStatusBar", "QStatusBar", args)
  }

  return QStatusBar{}
}

// clearMessage()
func (this *QStatusBar) clearMessage(args ...interface{}) () {
  // clearMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar12clearMessageEv
    // invoke: void clearMessage()
    C._ZN10QStatusBar12clearMessageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "clearMessage", args)
  }

}

// insertWidget(int, class QWidget *, int)
func (this *QStatusBar) insertWidget(args ...interface{}) () {
  // insertWidget(int, class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar12insertWidgetEiP7QWidgeti
    // invoke: int insertWidget(int, class QWidget *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QStatusBar12insertWidgetEiP7QWidgeti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStatusBar", "insertWidget", args)
  }

}

// addPermanentWidget(class QWidget *, int)
func (this *QStatusBar) addPermanentWidget(args ...interface{}) () {
  // addPermanentWidget(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar18addPermanentWidgetEP7QWidgeti
    // invoke: void addPermanentWidget(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QStatusBar18addPermanentWidgetEP7QWidgeti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStatusBar", "addPermanentWidget", args)
  }

}

// currentMessage()
func (this *QStatusBar) currentMessage(args ...interface{}) () {
  // currentMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStatusBar14currentMessageEv
    // invoke: QString currentMessage()
    C._ZNK10QStatusBar14currentMessageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "currentMessage", args)
  }

}

// addWidget(class QWidget *, int)
func (this *QStatusBar) addWidget(args ...interface{}) () {
  // addWidget(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar9addWidgetEP7QWidgeti
    // invoke: void addWidget(class QWidget *, int)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QStatusBar9addWidgetEP7QWidgeti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStatusBar", "addWidget", args)
  }

}

// setSizeGripEnabled(_Bool)
func (this *QStatusBar) setSizeGripEnabled(args ...interface{}) () {
  // setSizeGripEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar18setSizeGripEnabledEb
    // invoke: void setSizeGripEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QStatusBar18setSizeGripEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStatusBar", "setSizeGripEnabled", args)
  }

}

// insertPermanentWidget(int, class QWidget *, int)
func (this *QStatusBar) insertPermanentWidget(args ...interface{}) () {
  // insertPermanentWidget(int, class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti
    // invoke: int insertPermanentWidget(int, class QWidget *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStatusBar", "insertPermanentWidget", args)
  }

}

// metaObject()
func (this *QStatusBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStatusBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QStatusBar10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "metaObject", args)
  }

}

// isSizeGripEnabled()
func (this *QStatusBar) isSizeGripEnabled(args ...interface{}) () {
  // isSizeGripEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStatusBar17isSizeGripEnabledEv
    // invoke: bool isSizeGripEnabled()
    C._ZNK10QStatusBar17isSizeGripEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "isSizeGripEnabled", args)
  }

}

// showMessage(const class QString &, int)
func (this *QStatusBar) showMessage(args ...interface{}) () {
  // showMessage(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar11showMessageERK7QStringi
    // invoke: void showMessage(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QStatusBar11showMessageERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStatusBar", "showMessage", args)
  }

}

// <= body block end

