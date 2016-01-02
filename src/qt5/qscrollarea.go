package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  void QScrollArea::QScrollArea(QWidget * parent);
extern void* dector_ZN11QScrollAreaC1EP7QWidget(void* arg0);
extern void _ZN11QScrollAreaC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QScrollArea::setWidgetResizable(bool resizable);
extern void _ZN11QScrollArea18setWidgetResizableEb(void* qthis, bool arg0);
  // proto:  void QScrollArea::QScrollArea(const QScrollArea & );
extern void* dector_ZN11QScrollAreaC1ERKS_(void* arg0);
extern void _ZN11QScrollAreaC1ERKS_(void* qthis, void* arg0);
  // proto:  void QScrollArea::setWidget(QWidget * widget);
extern void _ZN11QScrollArea9setWidgetEP7QWidget(void* qthis, void* arg0);
  // proto:  QWidget * QScrollArea::takeWidget();
extern void _ZN11QScrollArea10takeWidgetEv(void* qthis);
  // proto:  void QScrollArea::ensureVisible(int x, int y, int xmargin, int ymargin);
extern void _ZN11QScrollArea13ensureVisibleEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QScrollArea::ensureWidgetVisible(QWidget * childWidget, int xmargin, int ymargin);
extern void _ZN11QScrollArea19ensureWidgetVisibleEP7QWidgetii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  QWidget * QScrollArea::widget();
extern void _ZNK11QScrollArea6widgetEv(void* qthis);
  // proto:  QSize QScrollArea::sizeHint();
extern void _ZNK11QScrollArea8sizeHintEv(void* qthis);
  // proto:  bool QScrollArea::widgetResizable();
extern void _ZNK11QScrollArea15widgetResizableEv(void* qthis);
  // proto:  void QScrollArea::~QScrollArea();
extern void _ZN11QScrollAreaD0Ev(void* qthis);
  // proto:  bool QScrollArea::focusNextPrevChild(bool next);
extern void _ZN11QScrollArea18focusNextPrevChildEb(void* qthis, bool arg0);
  // proto:  const QMetaObject * QScrollArea::metaObject();
extern void _ZNK11QScrollArea10metaObjectEv(void* qthis);
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

// class sizeof(QScrollArea)=1
type QScrollArea struct {
  /*qbase*/ QAbstractScrollArea;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QScrollArea::QScrollArea(QWidget * parent);
func NewQScrollArea(args ...interface{}) QScrollArea {
  return QScrollArea{}
}

  // proto:  void QScrollArea::setWidgetResizable(bool resizable);
func (this *QScrollArea) setWidgetResizable(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QScrollArea", "setWidgetResizable", args)
  }

}

  // proto:  void QScrollArea::setWidget(QWidget * widget);
func (this *QScrollArea) setWidget(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QScrollArea", "setWidget", args)
  }

}

  // proto:  QWidget * QScrollArea::takeWidget();
func (this *QScrollArea) takeWidget(args ...interface{}) () {
  // takeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QScrollArea10takeWidgetEv
  default:
    qtrt.ErrorResolve("QScrollArea", "takeWidget", args)
  }

}

  // proto:  void QScrollArea::ensureVisible(int x, int y, int xmargin, int ymargin);
func (this *QScrollArea) ensureVisible(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QScrollArea", "ensureVisible", args)
  }

}

  // proto:  void QScrollArea::ensureWidgetVisible(QWidget * childWidget, int xmargin, int ymargin);
func (this *QScrollArea) ensureWidgetVisible(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QScrollArea", "ensureWidgetVisible", args)
  }

}

  // proto:  QWidget * QScrollArea::widget();
func (this *QScrollArea) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea6widgetEv
  default:
    qtrt.ErrorResolve("QScrollArea", "widget", args)
  }

}

  // proto:  QSize QScrollArea::sizeHint();
func (this *QScrollArea) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea8sizeHintEv
  default:
    qtrt.ErrorResolve("QScrollArea", "sizeHint", args)
  }

}

  // proto:  bool QScrollArea::widgetResizable();
func (this *QScrollArea) widgetResizable(args ...interface{}) () {
  // widgetResizable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea15widgetResizableEv
  default:
    qtrt.ErrorResolve("QScrollArea", "widgetResizable", args)
  }

}

  // proto:  void QScrollArea::~QScrollArea();
func (this *QScrollArea) FreeQScrollArea(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QScrollArea", "~QScrollArea", args)
  }

}

  // proto:  bool QScrollArea::focusNextPrevChild(bool next);
func (this *QScrollArea) focusNextPrevChild(args ...interface{}) () {
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
  default:
    qtrt.ErrorResolve("QScrollArea", "focusNextPrevChild", args)
  }

}

  // proto:  const QMetaObject * QScrollArea::metaObject();
func (this *QScrollArea) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QScrollArea10metaObjectEv
  default:
    qtrt.ErrorResolve("QScrollArea", "metaObject", args)
  }

}

// <= body block end

