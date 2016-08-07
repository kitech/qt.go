package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  void QStatusBar::~QStatusBar();
extern void C_ZN10QStatusBarD2Ev(void* qthis); // 4
  // proto:  void QStatusBar::removeWidget(QWidget * widget);
extern void C_ZN10QStatusBar12removeWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStatusBar::QStatusBar(QWidget * parent);
extern void* C_ZN10QStatusBarC2EP7QWidget(void* arg0); // 3
  // proto:  void QStatusBar::clearMessage();
extern void C_ZN10QStatusBar12clearMessageEv(void* qthis); // 4
  // proto:  int QStatusBar::insertWidget(int index, QWidget * widget, int stretch);
extern int32_t C_ZN10QStatusBar12insertWidgetEiP7QWidgeti(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QStatusBar::addPermanentWidget(QWidget * widget, int stretch);
extern void C_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QString QStatusBar::currentMessage();
extern void* C_ZNK10QStatusBar14currentMessageEv(void* qthis); // 4
  // proto:  void QStatusBar::addWidget(QWidget * widget, int stretch);
extern void C_ZN10QStatusBar9addWidgetEP7QWidgeti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QStatusBar::setSizeGripEnabled(bool );
extern void C_ZN10QStatusBar18setSizeGripEnabledEb(void* qthis, bool arg0); // 4
  // proto:  int QStatusBar::insertPermanentWidget(int index, QWidget * widget, int stretch);
extern int32_t C_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  const QMetaObject * QStatusBar::metaObject();
extern void C_ZNK10QStatusBar10metaObjectEv(void* qthis); // 4
  // proto:  bool QStatusBar::isSizeGripEnabled();
extern bool C_ZNK10QStatusBar17isSizeGripEnabledEv(void* qthis); // 4
  // proto:  void QStatusBar::showMessage(const QString & text, int timeout);
extern void C_ZN10QStatusBar11showMessageERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 4
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

// class sizeof(QStatusBar)=1
type QStatusBar struct {
  /*qbase*/ QWidget;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _messageChanged QStatusBar_messageChanged_signal;
}

// ~QStatusBar()
func (this *QStatusBar) Freeqstatusbar(args ...interface{}) () {
  // ~QStatusBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBarD0Ev
    // invoke: void ~QStatusBar()
    C.C_ZN10QStatusBarD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "~QStatusBar", args)
  }

  return
}

// removeWidget(class QWidget *)
func (this *QStatusBar) Removewidget(args ...interface{}) () {
  // removeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar12removeWidgetEP7QWidget
    // invoke: void removeWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QStatusBar12removeWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStatusBar", "removeWidget", args)
  }

  return
}

// QStatusBar(class QWidget *)
func NewQStatusBar(args ...interface{}) *QStatusBar {
  // QStatusBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBarC1EP7QWidget
    // invoke: void QStatusBar(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QStatusBarC2EP7QWidget(arg0)
    return &QStatusBar{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStatusBar", "QStatusBar", args)
  }

  return nil // QStatusBar{Qclsinst:qthis}
}

// clearMessage()
func (this *QStatusBar) Clearmessage(args ...interface{}) () {
  // clearMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar12clearMessageEv
    // invoke: void clearMessage()
    C.C_ZN10QStatusBar12clearMessageEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "clearMessage", args)
  }

  return
}

// insertWidget(int, class QWidget *, int)
func (this *QStatusBar) Insertwidget(args ...interface{}) (ret interface{}) {
  // insertWidget(int, class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar12insertWidgetEiP7QWidgeti
    // invoke: int insertWidget(int, class QWidget *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QStatusBar12insertWidgetEiP7QWidgeti(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStatusBar", "insertWidget", args)
  }

  return
}

// addPermanentWidget(class QWidget *, int)
func (this *QStatusBar) Addpermanentwidget(args ...interface{}) () {
  // addPermanentWidget(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar18addPermanentWidgetEP7QWidgeti
    // invoke: void addPermanentWidget(class QWidget *, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QStatusBar18addPermanentWidgetEP7QWidgeti(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStatusBar", "addPermanentWidget", args)
  }

  return
}

// currentMessage()
func (this *QStatusBar) Currentmessage(args ...interface{}) (ret interface{}) {
  // currentMessage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStatusBar14currentMessageEv
    // invoke: QString currentMessage()
    var ret0 = C.C_ZNK10QStatusBar14currentMessageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStatusBar", "currentMessage", args)
  }

  return
}

// addWidget(class QWidget *, int)
func (this *QStatusBar) Addwidget(args ...interface{}) () {
  // addWidget(class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar9addWidgetEP7QWidgeti
    // invoke: void addWidget(class QWidget *, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QStatusBar9addWidgetEP7QWidgeti(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStatusBar", "addWidget", args)
  }

  return
}

// setSizeGripEnabled(_Bool)
func (this *QStatusBar) Setsizegripenabled(args ...interface{}) () {
  // setSizeGripEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar18setSizeGripEnabledEb
    // invoke: void setSizeGripEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN10QStatusBar18setSizeGripEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStatusBar", "setSizeGripEnabled", args)
  }

  return
}

// insertPermanentWidget(int, class QWidget *, int)
func (this *QStatusBar) Insertpermanentwidget(args ...interface{}) (ret interface{}) {
  // insertPermanentWidget(int, class QWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti
    // invoke: int insertPermanentWidget(int, class QWidget *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStatusBar", "insertPermanentWidget", args)
  }

  return
}

// metaObject()
func (this *QStatusBar) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStatusBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK10QStatusBar10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStatusBar", "metaObject", args)
  }

  return
}

// isSizeGripEnabled()
func (this *QStatusBar) Issizegripenabled(args ...interface{}) (ret interface{}) {
  // isSizeGripEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStatusBar17isSizeGripEnabledEv
    // invoke: bool isSizeGripEnabled()
    var ret0 = C.C_ZNK10QStatusBar17isSizeGripEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStatusBar", "isSizeGripEnabled", args)
  }

  return
}

// showMessage(const class QString &, int)
func (this *QStatusBar) Showmessage(args ...interface{}) () {
  // showMessage(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStatusBar11showMessageERK7QStringi
    // invoke: void showMessage(const class QString &, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN10QStatusBar11showMessageERK7QStringi(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStatusBar", "showMessage", args)
  }

  return
}

// <= body block end

