package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtWidgets/qwidgetaction.h
// dst-file: /src/widgets/qwidgetaction.go
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
  // proto:  const QMetaObject * QWidgetAction::metaObject();
extern void _ZNK13QWidgetAction10metaObjectEv(void* qthis);
  // proto:  void QWidgetAction::~QWidgetAction();
extern void _ZN13QWidgetActionD0Ev(void* qthis);
  // proto:  void QWidgetAction::setDefaultWidget(QWidget * w);
extern void _ZN13QWidgetAction16setDefaultWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QWidgetAction::releaseWidget(QWidget * widget);
extern void _ZN13QWidgetAction13releaseWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  void QWidgetAction::QWidgetAction(const QWidgetAction & );
extern void* dector_ZN13QWidgetActionC1ERKS_(void* arg0);
extern void _ZN13QWidgetActionC1ERKS_(void* qthis, void* arg0);
  // proto:  void QWidgetAction::QWidgetAction(QObject * parent);
extern void* dector_ZN13QWidgetActionC1EP7QObject(void* arg0);
extern void _ZN13QWidgetActionC1EP7QObject(void* qthis, void* arg0);
  // proto:  QWidget * QWidgetAction::requestWidget(QWidget * parent);
extern void _ZN13QWidgetAction13requestWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QWidget * QWidgetAction::defaultWidget();
extern void _ZNK13QWidgetAction13defaultWidgetEv(void* qthis);
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

// class sizeof(QWidgetAction)=1
type QWidgetAction struct {
  /*qbase*/ QAction;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  const QMetaObject * QWidgetAction::metaObject();
func (this *QWidgetAction) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetAction10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK13QWidgetAction10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetAction", "metaObject", args)
  }

}

  // proto:  void QWidgetAction::~QWidgetAction();
func (this *QWidgetAction) FreeQWidgetAction(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QWidgetAction", "~QWidgetAction", args)
  }

}

  // proto:  void QWidgetAction::setDefaultWidget(QWidget * w);
func (this *QWidgetAction) setDefaultWidget(args ...interface{}) () {
  // setDefaultWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction16setDefaultWidgetEP7QWidget
    // invoke: void setDefaultWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QWidgetAction16setDefaultWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "setDefaultWidget", args)
  }

}

  // proto:  void QWidgetAction::releaseWidget(QWidget * widget);
func (this *QWidgetAction) releaseWidget(args ...interface{}) () {
  // releaseWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction13releaseWidgetEP7QWidget
    // invoke: void releaseWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QWidgetAction13releaseWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "releaseWidget", args)
  }

}

  // proto:  void QWidgetAction::QWidgetAction(const QWidgetAction & );
func NewQWidgetAction(args ...interface{}) QWidgetAction {
  return QWidgetAction{}
}

  // proto:  QWidget * QWidgetAction::requestWidget(QWidget * parent);
func (this *QWidgetAction) requestWidget(args ...interface{}) () {
  // requestWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetAction13requestWidgetEP7QWidget
    // invoke: QWidget * requestWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QWidgetAction13requestWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "requestWidget", args)
  }

}

  // proto:  QWidget * QWidgetAction::defaultWidget();
func (this *QWidgetAction) defaultWidget(args ...interface{}) () {
  // defaultWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QWidgetAction13defaultWidgetEv
    // invoke: QWidget * defaultWidget()
    C._ZNK13QWidgetAction13defaultWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetAction", "defaultWidget", args)
  }

}

// <= body block end

