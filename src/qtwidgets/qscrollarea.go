package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtWidgets/qscrollarea.h
// dst-file: /src/widgets/qscrollarea.go
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
  // proto:  void QScrollArea::ensureWidgetVisible(QWidget * childWidget, int xmargin, int ymargin);
extern void C_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 4
  // proto:  void QScrollArea::QScrollArea(QWidget * parent);
extern void* C_ZN11QScrollAreaC2EP7QWidget(void* arg0); // 3
  // proto:  void QScrollArea::~QScrollArea();
extern void C_ZN11QScrollAreaD2Ev(void* qthis); // 4
  // proto:  void QScrollArea::setWidget(QWidget * widget);
extern void C_ZN11QScrollArea9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  bool QScrollArea::focusNextPrevChild(bool next);
extern bool C_ZN11QScrollArea18focusNextPrevChildEb(void* qthis, bool arg0); // 4
  // proto:  Qt::Alignment QScrollArea::alignment();
extern void C_ZNK11QScrollArea9alignmentEv(void* qthis); // 4
  // proto:  bool QScrollArea::widgetResizable();
extern bool C_ZNK11QScrollArea15widgetResizableEv(void* qthis); // 4
  // proto:  void QScrollArea::setWidgetResizable(bool resizable);
extern void C_ZN11QScrollArea18setWidgetResizableEb(void* qthis, bool arg0); // 4
  // proto:  QWidget * QScrollArea::widget();
extern void* C_ZNK11QScrollArea6widgetEv(void* qthis); // 4
  // proto:  void QScrollArea::ensureVisible(int x, int y, int xmargin, int ymargin);
extern void C_ZN11QScrollArea13ensureVisibleEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  const QMetaObject * QScrollArea::metaObject();
extern void C_ZNK11QScrollArea10metaObjectEv(void* qthis); // 4
  // proto:  QWidget * QScrollArea::takeWidget();
extern void* C_ZN11QScrollArea10takeWidgetEv(void* qthis); // 4
  // proto:  QSize QScrollArea::sizeHint();
extern void* C_ZNK11QScrollArea8sizeHintEv(void* qthis); // 4
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

// class sizeof(QScrollArea)=1
type QScrollArea struct {
  /*qbase*/ QAbstractScrollArea;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ensureWidgetVisible(class QWidget *, int, int)
func (this *QScrollArea) Ensurewidgetvisible(args ...interface{}) () {
  // ensureWidgetVisible(class QWidget *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii
    // invoke: void ensureWidgetVisible(class QWidget *, int, int)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QScrollArea", "ensureWidgetVisible", args)
  }

  return
}

// QScrollArea(class QWidget *)
func NewQScrollArea(args ...interface{}) *QScrollArea {
  // QScrollArea(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollAreaC1EP7QWidget
    // invoke: void QScrollArea(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QScrollAreaC2EP7QWidget(arg0)
    return &QScrollArea{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QScrollArea", "QScrollArea", args)
  }

  return nil // QScrollArea{Qclsinst:qthis}
}

// ~QScrollArea()
func (this *QScrollArea) Freeqscrollarea(args ...interface{}) () {
  // ~QScrollArea()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollAreaD0Ev
    // invoke: void ~QScrollArea()
    C.C_ZN11QScrollAreaD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollArea", "~QScrollArea", args)
  }

  return
}

// setWidget(class QWidget *)
func (this *QScrollArea) Setwidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea9setWidgetEP7QWidget
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QScrollArea9setWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollArea", "setWidget", args)
  }

  return
}

// focusNextPrevChild(_Bool)
func (this *QScrollArea) Focusnextprevchild(args ...interface{}) (ret interface{}) {
  // focusNextPrevChild(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea18focusNextPrevChildEb
    // invoke: bool focusNextPrevChild(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QScrollArea18focusNextPrevChildEb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollArea", "focusNextPrevChild", args)
  }

  return
}

// alignment()
func (this *QScrollArea) Alignment(args ...interface{}) () {
  // alignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea9alignmentEv
    // invoke: Qt::Alignment alignment()
    C.C_ZNK11QScrollArea9alignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollArea", "alignment", args)
  }

  return
}

// widgetResizable()
func (this *QScrollArea) Widgetresizable(args ...interface{}) (ret interface{}) {
  // widgetResizable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea15widgetResizableEv
    // invoke: bool widgetResizable()
    var ret0 = C.C_ZNK11QScrollArea15widgetResizableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollArea", "widgetResizable", args)
  }

  return
}

// setWidgetResizable(_Bool)
func (this *QScrollArea) Setwidgetresizable(args ...interface{}) () {
  // setWidgetResizable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea18setWidgetResizableEb
    // invoke: void setWidgetResizable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QScrollArea18setWidgetResizableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QScrollArea", "setWidgetResizable", args)
  }

  return
}

// widget()
func (this *QScrollArea) Widget(args ...interface{}) (ret interface{}) {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea6widgetEv
    // invoke: QWidget * widget()
    var ret0 = C.C_ZNK11QScrollArea6widgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollArea", "widget", args)
  }

  return
}

// ensureVisible(int, int, int, int)
func (this *QScrollArea) Ensurevisible(args ...interface{}) () {
  // ensureVisible(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea13ensureVisibleEiiii
    // invoke: void ensureVisible(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN11QScrollArea13ensureVisibleEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QScrollArea", "ensureVisible", args)
  }

  return
}

// metaObject()
func (this *QScrollArea) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QScrollArea10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QScrollArea", "metaObject", args)
  }

  return
}

// takeWidget()
func (this *QScrollArea) Takewidget(args ...interface{}) (ret interface{}) {
  // takeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea10takeWidgetEv
    // invoke: QWidget * takeWidget()
    var ret0 = C.C_ZN11QScrollArea10takeWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollArea", "takeWidget", args)
  }

  return
}

// sizeHint()
func (this *QScrollArea) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK11QScrollArea8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QScrollArea", "sizeHint", args)
  }

  return
}

// <= body block end

