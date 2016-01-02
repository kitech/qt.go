package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QStatusBar::~QStatusBar();
extern void _ZN10QStatusBarD0Ev(void* qthis);
  // proto:  int QStatusBar::insertPermanentWidget(int index, QWidget * widget, int stretch);
extern void _ZN10QStatusBar21insertPermanentWidgetEiP7QWidgeti(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QStatusBar::removeWidget(QWidget * widget);
extern void _ZN10QStatusBar12removeWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QStatusBar::QStatusBar(const QStatusBar & );
extern void* dector_ZN10QStatusBarC1ERKS_(void* arg0);
extern void _ZN10QStatusBarC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStatusBar::setSizeGripEnabled(bool );
extern void _ZN10QStatusBar18setSizeGripEnabledEb(void* qthis, bool arg0);
  // proto:  void QStatusBar::addPermanentWidget(QWidget * widget, int stretch);
extern void _ZN10QStatusBar18addPermanentWidgetEP7QWidgeti(void* qthis, void* arg0, int arg1);
  // proto:  bool QStatusBar::isSizeGripEnabled();
extern void _ZNK10QStatusBar17isSizeGripEnabledEv(void* qthis);
  // proto:  void QStatusBar::clearMessage();
extern void _ZN10QStatusBar12clearMessageEv(void* qthis);
  // proto:  QString QStatusBar::currentMessage();
extern void _ZNK10QStatusBar14currentMessageEv(void* qthis);
  // proto:  const QMetaObject * QStatusBar::metaObject();
extern void _ZNK10QStatusBar10metaObjectEv(void* qthis);
  // proto:  void QStatusBar::showMessage(const QString & text, int timeout);
extern void _ZN10QStatusBar11showMessageERK7QStringi(void* qthis, void* arg0, int arg1);
  // proto:  int QStatusBar::insertWidget(int index, QWidget * widget, int stretch);
extern void _ZN10QStatusBar12insertWidgetEiP7QWidgeti(void* qthis, int arg0, void* arg1, int arg2);
  // proto:  void QStatusBar::addWidget(QWidget * widget, int stretch);
extern void _ZN10QStatusBar9addWidgetEP7QWidgeti(void* qthis, void* arg0, int arg1);
  // proto:  void QStatusBar::QStatusBar(QWidget * parent);
extern void* dector_ZN10QStatusBarC1EP7QWidget(void* arg0);
extern void _ZN10QStatusBarC1EP7QWidget(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _messageChanged QStatusBar_messageChanged_signal;
}

  // proto:  void QStatusBar::~QStatusBar();
func (this *QStatusBar) FreeQStatusBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStatusBar", "~QStatusBar", args)
  }

}

  // proto:  int QStatusBar::insertPermanentWidget(int index, QWidget * widget, int stretch);
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
  default:
    qtrt.ErrorResolve("QStatusBar", "insertPermanentWidget", args)
  }

}

  // proto:  void QStatusBar::removeWidget(QWidget * widget);
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
  default:
    qtrt.ErrorResolve("QStatusBar", "removeWidget", args)
  }

}

  // proto:  void QStatusBar::QStatusBar(const QStatusBar & );
func NewQStatusBar(args ...interface{}) QStatusBar {
  return QStatusBar{}
}

  // proto:  void QStatusBar::setSizeGripEnabled(bool );
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
  default:
    qtrt.ErrorResolve("QStatusBar", "setSizeGripEnabled", args)
  }

}

  // proto:  void QStatusBar::addPermanentWidget(QWidget * widget, int stretch);
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
  default:
    qtrt.ErrorResolve("QStatusBar", "addPermanentWidget", args)
  }

}

  // proto:  bool QStatusBar::isSizeGripEnabled();
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
  default:
    qtrt.ErrorResolve("QStatusBar", "isSizeGripEnabled", args)
  }

}

  // proto:  void QStatusBar::clearMessage();
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
  default:
    qtrt.ErrorResolve("QStatusBar", "clearMessage", args)
  }

}

  // proto:  QString QStatusBar::currentMessage();
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
  default:
    qtrt.ErrorResolve("QStatusBar", "currentMessage", args)
  }

}

  // proto:  const QMetaObject * QStatusBar::metaObject();
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
  default:
    qtrt.ErrorResolve("QStatusBar", "metaObject", args)
  }

}

  // proto:  void QStatusBar::showMessage(const QString & text, int timeout);
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
  default:
    qtrt.ErrorResolve("QStatusBar", "showMessage", args)
  }

}

  // proto:  int QStatusBar::insertWidget(int index, QWidget * widget, int stretch);
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
  default:
    qtrt.ErrorResolve("QStatusBar", "insertWidget", args)
  }

}

  // proto:  void QStatusBar::addWidget(QWidget * widget, int stretch);
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
  default:
    qtrt.ErrorResolve("QStatusBar", "addWidget", args)
  }

}

// <= body block end

