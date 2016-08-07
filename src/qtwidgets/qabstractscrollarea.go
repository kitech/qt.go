package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtWidgets/qabstractscrollarea.h
// dst-file: /src/widgets/qabstractscrollarea.go
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
  // proto:  Qt::ScrollBarPolicy QAbstractScrollArea::horizontalScrollBarPolicy();
extern void C_ZNK19QAbstractScrollArea25horizontalScrollBarPolicyEv(void* qthis); // 4
  // proto:  QSize QAbstractScrollArea::maximumViewportSize();
extern void* C_ZNK19QAbstractScrollArea19maximumViewportSizeEv(void* qthis); // 4
  // proto:  QWidget * QAbstractScrollArea::cornerWidget();
extern void* C_ZNK19QAbstractScrollArea12cornerWidgetEv(void* qthis); // 4
  // proto:  void QAbstractScrollArea::setVerticalScrollBar(QScrollBar * scrollbar);
extern void C_ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar(void* qthis, void* arg0); // 4
  // proto:  void QAbstractScrollArea::setHorizontalScrollBar(QScrollBar * scrollbar);
extern void C_ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar(void* qthis, void* arg0); // 4
  // proto:  void QAbstractScrollArea::setViewport(QWidget * widget);
extern void C_ZN19QAbstractScrollArea11setViewportEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QAbstractScrollArea::setCornerWidget(QWidget * widget);
extern void C_ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QAbstractScrollArea::setupViewport(QWidget * viewport);
extern void C_ZN19QAbstractScrollArea13setupViewportEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QAbstractScrollArea::QAbstractScrollArea(QWidget * parent);
extern void* C_ZN19QAbstractScrollAreaC2EP7QWidget(void* arg0); // 3
  // proto:  void QAbstractScrollArea::~QAbstractScrollArea();
extern void C_ZN19QAbstractScrollAreaD2Ev(void* qthis); // 4
  // proto:  QScrollBar * QAbstractScrollArea::verticalScrollBar();
extern void* C_ZNK19QAbstractScrollArea17verticalScrollBarEv(void* qthis); // 4
  // proto:  QScrollBar * QAbstractScrollArea::horizontalScrollBar();
extern void* C_ZNK19QAbstractScrollArea19horizontalScrollBarEv(void* qthis); // 4
  // proto:  QWidget * QAbstractScrollArea::viewport();
extern void* C_ZNK19QAbstractScrollArea8viewportEv(void* qthis); // 4
  // proto:  const QMetaObject * QAbstractScrollArea::metaObject();
extern void C_ZNK19QAbstractScrollArea10metaObjectEv(void* qthis); // 4
  // proto:  Qt::ScrollBarPolicy QAbstractScrollArea::verticalScrollBarPolicy();
extern void C_ZNK19QAbstractScrollArea23verticalScrollBarPolicyEv(void* qthis); // 4
  // proto:  QSize QAbstractScrollArea::sizeHint();
extern void* C_ZNK19QAbstractScrollArea8sizeHintEv(void* qthis); // 4
  // proto:  QSize QAbstractScrollArea::minimumSizeHint();
extern void* C_ZNK19QAbstractScrollArea15minimumSizeHintEv(void* qthis); // 4
  // proto:  QAbstractScrollArea::SizeAdjustPolicy QAbstractScrollArea::sizeAdjustPolicy();
extern void C_ZNK19QAbstractScrollArea16sizeAdjustPolicyEv(void* qthis); // 4
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

// class sizeof(QAbstractScrollArea)=1
type QAbstractScrollArea struct {
  /*qbase*/ QFrame;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// horizontalScrollBarPolicy()
func (this *QAbstractScrollArea) Horizontalscrollbarpolicy(args ...interface{}) () {
  // horizontalScrollBarPolicy()
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
    // invoke: _ZNK19QAbstractScrollArea25horizontalScrollBarPolicyEv
    // invoke: Qt::ScrollBarPolicy horizontalScrollBarPolicy()
    C.C_ZNK19QAbstractScrollArea25horizontalScrollBarPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "horizontalScrollBarPolicy", args)
  }

  return
}

// maximumViewportSize()
func (this *QAbstractScrollArea) Maximumviewportsize(args ...interface{}) (ret interface{}) {
  // maximumViewportSize()
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
    // invoke: _ZNK19QAbstractScrollArea19maximumViewportSizeEv
    // invoke: QSize maximumViewportSize()
    var ret0 = C.C_ZNK19QAbstractScrollArea19maximumViewportSizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "maximumViewportSize", args)
  }

  return
}

// cornerWidget()
func (this *QAbstractScrollArea) Cornerwidget(args ...interface{}) (ret interface{}) {
  // cornerWidget()
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
    // invoke: _ZNK19QAbstractScrollArea12cornerWidgetEv
    // invoke: QWidget * cornerWidget()
    var ret0 = C.C_ZNK19QAbstractScrollArea12cornerWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "cornerWidget", args)
  }

  return
}

// setVerticalScrollBar(class QScrollBar *)
func (this *QAbstractScrollArea) Setverticalscrollbar(args ...interface{}) () {
  // setVerticalScrollBar(class QScrollBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar
    // invoke: void setVerticalScrollBar(class QScrollBar *)
    var arg0 = args[0].(*QScrollBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractScrollArea20setVerticalScrollBarEP10QScrollBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setVerticalScrollBar", args)
  }

  return
}

// setHorizontalScrollBar(class QScrollBar *)
func (this *QAbstractScrollArea) Sethorizontalscrollbar(args ...interface{}) () {
  // setHorizontalScrollBar(class QScrollBar *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar
    // invoke: void setHorizontalScrollBar(class QScrollBar *)
    var arg0 = args[0].(*QScrollBar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractScrollArea22setHorizontalScrollBarEP10QScrollBar(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setHorizontalScrollBar", args)
  }

  return
}

// setViewport(class QWidget *)
func (this *QAbstractScrollArea) Setviewport(args ...interface{}) () {
  // setViewport(class QWidget *)
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
    // invoke: _ZN19QAbstractScrollArea11setViewportEP7QWidget
    // invoke: void setViewport(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractScrollArea11setViewportEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setViewport", args)
  }

  return
}

// setCornerWidget(class QWidget *)
func (this *QAbstractScrollArea) Setcornerwidget(args ...interface{}) () {
  // setCornerWidget(class QWidget *)
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
    // invoke: _ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget
    // invoke: void setCornerWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractScrollArea15setCornerWidgetEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setCornerWidget", args)
  }

  return
}

// setupViewport(class QWidget *)
func (this *QAbstractScrollArea) Setupviewport(args ...interface{}) () {
  // setupViewport(class QWidget *)
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
    // invoke: _ZN19QAbstractScrollArea13setupViewportEP7QWidget
    // invoke: void setupViewport(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QAbstractScrollArea13setupViewportEP7QWidget(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "setupViewport", args)
  }

  return
}

// QAbstractScrollArea(class QWidget *)
func NewQAbstractScrollArea(args ...interface{}) *QAbstractScrollArea {
  // QAbstractScrollArea(class QWidget *)
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
    // invoke: _ZN19QAbstractScrollAreaC1EP7QWidget
    // invoke: void QAbstractScrollArea(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QAbstractScrollAreaC2EP7QWidget(arg0)
    return &QAbstractScrollArea{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "QAbstractScrollArea", args)
  }

  return nil // QAbstractScrollArea{Qclsinst:qthis}
}

// ~QAbstractScrollArea()
func (this *QAbstractScrollArea) Freeqabstractscrollarea(args ...interface{}) () {
  // ~QAbstractScrollArea()
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
    // invoke: _ZN19QAbstractScrollAreaD0Ev
    // invoke: void ~QAbstractScrollArea()
    C.C_ZN19QAbstractScrollAreaD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "~QAbstractScrollArea", args)
  }

  return
}

// verticalScrollBar()
func (this *QAbstractScrollArea) Verticalscrollbar(args ...interface{}) (ret interface{}) {
  // verticalScrollBar()
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
    // invoke: _ZNK19QAbstractScrollArea17verticalScrollBarEv
    // invoke: QScrollBar * verticalScrollBar()
    var ret0 = C.C_ZNK19QAbstractScrollArea17verticalScrollBarEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "verticalScrollBar", args)
  }

  return
}

// horizontalScrollBar()
func (this *QAbstractScrollArea) Horizontalscrollbar(args ...interface{}) (ret interface{}) {
  // horizontalScrollBar()
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
    // invoke: _ZNK19QAbstractScrollArea19horizontalScrollBarEv
    // invoke: QScrollBar * horizontalScrollBar()
    var ret0 = C.C_ZNK19QAbstractScrollArea19horizontalScrollBarEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QScrollBar{}) // "QScrollBar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "horizontalScrollBar", args)
  }

  return
}

// viewport()
func (this *QAbstractScrollArea) Viewport(args ...interface{}) (ret interface{}) {
  // viewport()
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
    // invoke: _ZNK19QAbstractScrollArea8viewportEv
    // invoke: QWidget * viewport()
    var ret0 = C.C_ZNK19QAbstractScrollArea8viewportEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "viewport", args)
  }

  return
}

// metaObject()
func (this *QAbstractScrollArea) Metaobject(args ...interface{}) () {
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
    // invoke: _ZNK19QAbstractScrollArea10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK19QAbstractScrollArea10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "metaObject", args)
  }

  return
}

// verticalScrollBarPolicy()
func (this *QAbstractScrollArea) Verticalscrollbarpolicy(args ...interface{}) () {
  // verticalScrollBarPolicy()
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
    // invoke: _ZNK19QAbstractScrollArea23verticalScrollBarPolicyEv
    // invoke: Qt::ScrollBarPolicy verticalScrollBarPolicy()
    C.C_ZNK19QAbstractScrollArea23verticalScrollBarPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "verticalScrollBarPolicy", args)
  }

  return
}

// sizeHint()
func (this *QAbstractScrollArea) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
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
    // invoke: _ZNK19QAbstractScrollArea8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK19QAbstractScrollArea8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "sizeHint", args)
  }

  return
}

// minimumSizeHint()
func (this *QAbstractScrollArea) Minimumsizehint(args ...interface{}) (ret interface{}) {
  // minimumSizeHint()
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
    // invoke: _ZNK19QAbstractScrollArea15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret0 = C.C_ZNK19QAbstractScrollArea15minimumSizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "minimumSizeHint", args)
  }

  return
}

// sizeAdjustPolicy()
func (this *QAbstractScrollArea) Sizeadjustpolicy(args ...interface{}) () {
  // sizeAdjustPolicy()
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
    // invoke: _ZNK19QAbstractScrollArea16sizeAdjustPolicyEv
    // invoke: QAbstractScrollArea::SizeAdjustPolicy sizeAdjustPolicy()
    C.C_ZNK19QAbstractScrollArea16sizeAdjustPolicyEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAbstractScrollArea", "sizeAdjustPolicy", args)
  }

  return
}

// <= body block end

