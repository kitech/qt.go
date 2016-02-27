package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
// src-file: /QtWidgets/qstackedwidget.h
// dst-file: /src/widgets/qstackedwidget.go
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
  // proto:  void QStackedWidget::setCurrentIndex(int index);
extern void C_ZN14QStackedWidget15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  QWidget * QStackedWidget::currentWidget();
extern void* C_ZNK14QStackedWidget13currentWidgetEv(void* qthis); // 4
  // proto:  void QStackedWidget::QStackedWidget(QWidget * parent);
extern void* C_ZN14QStackedWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  void QStackedWidget::setCurrentWidget(QWidget * w);
extern void C_ZN14QStackedWidget16setCurrentWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStackedWidget::removeWidget(QWidget * w);
extern void C_ZN14QStackedWidget12removeWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  QWidget * QStackedWidget::widget(int );
extern void* C_ZNK14QStackedWidget6widgetEi(void* qthis, int32_t arg0); // 4
  // proto:  int QStackedWidget::indexOf(QWidget * );
extern int32_t C_ZNK14QStackedWidget7indexOfEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  int QStackedWidget::insertWidget(int index, QWidget * w);
extern int32_t C_ZN14QStackedWidget12insertWidgetEiP7QWidget(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  int QStackedWidget::addWidget(QWidget * w);
extern int32_t C_ZN14QStackedWidget9addWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QStackedWidget::~QStackedWidget();
extern void C_ZN14QStackedWidgetD2Ev(void* qthis); // 4
  // proto:  int QStackedWidget::count();
extern int32_t C_ZNK14QStackedWidget5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QStackedWidget::metaObject();
extern void C_ZNK14QStackedWidget10metaObjectEv(void* qthis); // 4
  // proto:  int QStackedWidget::currentIndex();
extern int32_t C_ZNK14QStackedWidget12currentIndexEv(void* qthis); // 4
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

// class sizeof(QStackedWidget)=1
type QStackedWidget struct {
  /*qbase*/ QFrame;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _widgetRemoved QStackedWidget_widgetRemoved_signal;
//  _currentChanged QStackedWidget_currentChanged_signal;
}

// setCurrentIndex(int)
func (this *QStackedWidget) SetCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedWidget15setCurrentIndexEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "setCurrentIndex", args)
  }

  return
}

// currentWidget()
func (this *QStackedWidget) CurrentWidget(args ...interface{}) (ret interface{}) {
  // currentWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget13currentWidgetEv
    // invoke: QWidget * currentWidget()
    var ret0 = C.C_ZNK14QStackedWidget13currentWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "currentWidget", args)
  }

  return
}

// QStackedWidget(class QWidget *)
func GcfreeQStackedWidget(this *QStackedWidget) {
  qtrt.UniverseFree(this)
}
func NewQStackedWidget(args ...interface{}) *QStackedWidget {
  // QStackedWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidgetC1EP7QWidget
    // invoke: void QStackedWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QStackedWidgetC2EP7QWidget(arg0)
    this := &QStackedWidget{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQStackedWidget)
    return this
  default:
    qtrt.ErrorResolve("QStackedWidget", "QStackedWidget", args)
  }

  return nil // QStackedWidget{Qclsinst:qthis}
}

// setCurrentWidget(class QWidget *)
func (this *QStackedWidget) SetCurrentWidget(args ...interface{}) () {
  // setCurrentWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget16setCurrentWidgetEP7QWidget
    // invoke: void setCurrentWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedWidget16setCurrentWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "setCurrentWidget", args)
  }

  return
}

// removeWidget(class QWidget *)
func (this *QStackedWidget) RemoveWidget(args ...interface{}) () {
  // removeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget12removeWidgetEP7QWidget
    // invoke: void removeWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN14QStackedWidget12removeWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStackedWidget", "removeWidget", args)
  }

  return
}

// widget(int)
func (this *QStackedWidget) Widget(args ...interface{}) (ret interface{}) {
  // widget(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget6widgetEi
    // invoke: QWidget * widget(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QStackedWidget6widgetEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "widget", args)
  }

  return
}

// indexOf(class QWidget *)
func (this *QStackedWidget) IndexOf(args ...interface{}) (ret interface{}) {
  // indexOf(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget7indexOfEP7QWidget
    // invoke: int indexOf(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QStackedWidget7indexOfEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "indexOf", args)
  }

  return
}

// insertWidget(int, class QWidget *)
func (this *QStackedWidget) InsertWidget(args ...interface{}) (ret interface{}) {
  // insertWidget(int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget12insertWidgetEiP7QWidget
    // invoke: int insertWidget(int, class QWidget *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN14QStackedWidget12insertWidgetEiP7QWidget(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "insertWidget", args)
  }

  return
}

// addWidget(class QWidget *)
func (this *QStackedWidget) AddWidget(args ...interface{}) (ret interface{}) {
  // addWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidget9addWidgetEP7QWidget
    // invoke: int addWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN14QStackedWidget9addWidgetEP7QWidget(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "addWidget", args)
  }

  return
}

// ~QStackedWidget()
func (this *QStackedWidget) Free(args ...interface{}) () {
  // ~QStackedWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QStackedWidgetD0Ev
    // invoke: void ~QStackedWidget()
    if this != nil && this.Qclsinst != nil {
      C.C_ZN14QStackedWidgetD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "~QStackedWidget", args)
  }

  return
}

// count()
func (this *QStackedWidget) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK14QStackedWidget5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "count", args)
  }

  return
}

// metaObject()
func (this *QStackedWidget) MetaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QStackedWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStackedWidget", "metaObject", args)
  }

  return
}

// currentIndex()
func (this *QStackedWidget) CurrentIndex(args ...interface{}) (ret interface{}) {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QStackedWidget12currentIndexEv
    // invoke: int currentIndex()
    var ret0 = C.C_ZNK14QStackedWidget12currentIndexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStackedWidget", "currentIndex", args)
  }

  return
}

// <= body block end

