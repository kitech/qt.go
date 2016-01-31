package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qdesktopwidget.h
// dst-file: /src/widgets/qdesktopwidget.go
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
  // proto:  int QDesktopWidget::screenNumber(const QWidget * widget);
extern int32_t C_ZNK14QDesktopWidget12screenNumberEPK7QWidget(void* qthis, void* arg0); // 4
  // proto:  int QDesktopWidget::screenNumber(const QPoint & );
extern int32_t C_ZNK14QDesktopWidget12screenNumberERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  int QDesktopWidget::screenCount();
extern int32_t C_ZNK14QDesktopWidget11screenCountEv(void* qthis); // 2
  // proto:  const QRect QDesktopWidget::availableGeometry(int screen);
extern void* C_ZNK14QDesktopWidget17availableGeometryEi(void* qthis, int32_t arg0); // 4
  // proto:  const QRect QDesktopWidget::availableGeometry(const QWidget * widget);
extern void* C_ZNK14QDesktopWidget17availableGeometryEPK7QWidget(void* qthis, void* arg0); // 4
  // proto:  const QRect QDesktopWidget::availableGeometry(const QPoint & point);
extern void* C_ZNK14QDesktopWidget17availableGeometryERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  const QRect QDesktopWidget::screenGeometry(const QPoint & point);
extern void* C_ZNK14QDesktopWidget14screenGeometryERK6QPoint(void* qthis, void* arg0); // 2
  // proto:  const QRect QDesktopWidget::screenGeometry(int screen);
extern void* C_ZNK14QDesktopWidget14screenGeometryEi(void* qthis, int32_t arg0); // 4
  // proto:  const QRect QDesktopWidget::screenGeometry(const QWidget * widget);
extern void* C_ZNK14QDesktopWidget14screenGeometryEPK7QWidget(void* qthis, void* arg0); // 4
  // proto:  bool QDesktopWidget::isVirtualDesktop();
extern bool C_ZNK14QDesktopWidget16isVirtualDesktopEv(void* qthis); // 4
  // proto:  QWidget * QDesktopWidget::screen(int screen);
extern void* C_ZN14QDesktopWidget6screenEi(void* qthis, int32_t arg0); // 4
  // proto:  int QDesktopWidget::primaryScreen();
extern int32_t C_ZNK14QDesktopWidget13primaryScreenEv(void* qthis); // 4
  // proto:  int QDesktopWidget::numScreens();
extern int32_t C_ZNK14QDesktopWidget10numScreensEv(void* qthis); // 4
  // proto:  void QDesktopWidget::QDesktopWidget();
extern void* C_ZN14QDesktopWidgetC2Ev(); // 3
  // proto:  const QMetaObject * QDesktopWidget::metaObject();
extern void C_ZNK14QDesktopWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QDesktopWidget::~QDesktopWidget();
extern void C_ZN14QDesktopWidgetD2Ev(void* qthis); // 4
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

// class sizeof(QDesktopWidget)=1
type QDesktopWidget struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _screenCountChanged QDesktopWidget_screenCountChanged_signal;
//  _resized QDesktopWidget_resized_signal;
//  _workAreaResized QDesktopWidget_workAreaResized_signal;
}

// screenNumber(const class QWidget *)
func (this *QDesktopWidget) Screennumber(args ...interface{}) (ret interface{}) {
  // screenNumber(const class QWidget *)
  // screenNumber(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget12screenNumberEPK7QWidget
    // invoke: int screenNumber(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget12screenNumberEPK7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK14QDesktopWidget12screenNumberERK6QPoint
    // invoke: int screenNumber(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget12screenNumberERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenNumber", args)
  }

  return
}

// screenCount()
func (this *QDesktopWidget) Screencount(args ...interface{}) (ret interface{}) {
  // screenCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget11screenCountEv
    // invoke: int screenCount()
    var ret0 = C.C_ZNK14QDesktopWidget11screenCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenCount", args)
  }

  return
}

// availableGeometry(int)
func (this *QDesktopWidget) Availablegeometry(args ...interface{}) (ret interface{}) {
  // availableGeometry(int)
  // availableGeometry(const class QWidget *)
  // availableGeometry(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget17availableGeometryEi
    // invoke: const QRect availableGeometry(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget17availableGeometryEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK14QDesktopWidget17availableGeometryEPK7QWidget
    // invoke: const QRect availableGeometry(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget17availableGeometryEPK7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK14QDesktopWidget17availableGeometryERK6QPoint
    // invoke: const QRect availableGeometry(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget17availableGeometryERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "availableGeometry", args)
  }

  return
}

// screenGeometry(const class QPoint &)
func (this *QDesktopWidget) Screengeometry(args ...interface{}) (ret interface{}) {
  // screenGeometry(const class QPoint &)
  // screenGeometry(int)
  // screenGeometry(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget14screenGeometryERK6QPoint
    // invoke: const QRect screenGeometry(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget14screenGeometryERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK14QDesktopWidget14screenGeometryEi
    // invoke: const QRect screenGeometry(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget14screenGeometryEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZNK14QDesktopWidget14screenGeometryEPK7QWidget
    // invoke: const QRect screenGeometry(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK14QDesktopWidget14screenGeometryEPK7QWidget(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "const QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenGeometry", args)
  }

  return
}

// isVirtualDesktop()
func (this *QDesktopWidget) Isvirtualdesktop(args ...interface{}) (ret interface{}) {
  // isVirtualDesktop()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget16isVirtualDesktopEv
    // invoke: bool isVirtualDesktop()
    var ret0 = C.C_ZNK14QDesktopWidget16isVirtualDesktopEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "isVirtualDesktop", args)
  }

  return
}

// screen(int)
func (this *QDesktopWidget) Screen(args ...interface{}) (ret interface{}) {
  // screen(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDesktopWidget6screenEi
    // invoke: QWidget * screen(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN14QDesktopWidget6screenEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screen", args)
  }

  return
}

// primaryScreen()
func (this *QDesktopWidget) Primaryscreen(args ...interface{}) (ret interface{}) {
  // primaryScreen()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget13primaryScreenEv
    // invoke: int primaryScreen()
    var ret0 = C.C_ZNK14QDesktopWidget13primaryScreenEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "primaryScreen", args)
  }

  return
}

// numScreens()
func (this *QDesktopWidget) Numscreens(args ...interface{}) (ret interface{}) {
  // numScreens()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget10numScreensEv
    // invoke: int numScreens()
    var ret0 = C.C_ZNK14QDesktopWidget10numScreensEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QDesktopWidget", "numScreens", args)
  }

  return
}

// QDesktopWidget()
func NewQDesktopWidget(args ...interface{}) *QDesktopWidget {
  // QDesktopWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDesktopWidgetC1Ev
    // invoke: void QDesktopWidget()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN14QDesktopWidgetC2Ev()
    return &QDesktopWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QDesktopWidget", "QDesktopWidget", args)
  }

  return nil // QDesktopWidget{qclsinst:qthis}
}

// metaObject()
func (this *QDesktopWidget) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK14QDesktopWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "metaObject", args)
  }

  return
}

// ~QDesktopWidget()
func (this *QDesktopWidget) Freeqdesktopwidget(args ...interface{}) () {
  // ~QDesktopWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN14QDesktopWidgetD0Ev
    // invoke: void ~QDesktopWidget()
    C.C_ZN14QDesktopWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "~QDesktopWidget", args)
  }

  return
}

// <= body block end

