package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
  // proto:  void QWidgetAction::releaseWidget(QWidget * widget);
extern void C_ZN13QWidgetAction13releaseWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QWidgetAction::metaObject();
extern void C_ZNK13QWidgetAction10metaObjectEv(void* qthis); // 4
  // proto:  void QWidgetAction::QWidgetAction(QObject * parent);
extern void* C_ZN13QWidgetActionC2EP7QObject(void* arg0); // 3
  // proto:  QWidget * QWidgetAction::defaultWidget();
extern void* C_ZNK13QWidgetAction13defaultWidgetEv(void* qthis); // 4
  // proto:  QWidget * QWidgetAction::requestWidget(QWidget * parent);
extern void* C_ZN13QWidgetAction13requestWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QWidgetAction::~QWidgetAction();
extern void C_ZN13QWidgetActionD2Ev(void* qthis); // 4
  // proto:  void QWidgetAction::setDefaultWidget(QWidget * w);
extern void C_ZN13QWidgetAction16setDefaultWidgetEP7QWidget(void* qthis, void* arg0); // 4
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

// class sizeof(QWidgetAction)=1
type QWidgetAction struct {
  /*qbase*/ QAction;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// releaseWidget(class QWidget *)
func (this *QWidgetAction) ReleaseWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QWidgetAction13releaseWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "releaseWidget", args)
  }

  return
}

// metaObject()
func (this *QWidgetAction) MetaObject(args ...interface{}) () {
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
    C.C_ZNK13QWidgetAction10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QWidgetAction", "metaObject", args)
  }

  return
}

// QWidgetAction(class QObject *)
func GcfreeQWidgetAction(this *QWidgetAction) {
  qtrt.UniverseFree(this)
}
func NewQWidgetAction(args ...interface{}) *QWidgetAction {
  // QWidgetAction(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetActionC1EP7QObject
    // invoke: void QWidgetAction(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QWidgetActionC2EP7QObject(arg0)
    this := &QWidgetAction{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQWidgetAction)
    return this
  default:
    qtrt.ErrorResolve("QWidgetAction", "QWidgetAction", args)
  }

  return nil // QWidgetAction{Qclsinst:qthis}
}

// defaultWidget()
func (this *QWidgetAction) DefaultWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QWidgetAction13defaultWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetAction", "defaultWidget", args)
  }

  return
}

// requestWidget(class QWidget *)
func (this *QWidgetAction) RequestWidget(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN13QWidgetAction13requestWidgetEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QWidgetAction", "requestWidget", args)
  }

  return
}

// ~QWidgetAction()
func (this *QWidgetAction) Free(args ...interface{}) () {
  // ~QWidgetAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QWidgetActionD0Ev
    // invoke: void ~QWidgetAction()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN13QWidgetActionD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QWidgetAction", "~QWidgetAction", args)
  }

  return
}

// setDefaultWidget(class QWidget *)
func (this *QWidgetAction) SetDefaultWidget(args ...interface{}) () {
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
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QWidgetAction16setDefaultWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QWidgetAction", "setDefaultWidget", args)
  }

  return
}

// <= body block end

