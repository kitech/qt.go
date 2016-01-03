package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  const QRect QDesktopWidget::screenGeometry(const QPoint & point);
extern void demth_ZNK14QDesktopWidget14screenGeometryERK6QPoint(void* qthis, void* arg0);
  // proto:  QWidget * QDesktopWidget::screen(int screen);
extern void _ZN14QDesktopWidget6screenEi(void* qthis, int32_t arg0);
  // proto:  const QRect QDesktopWidget::screenGeometry(const QWidget * widget);
extern void _ZNK14QDesktopWidget14screenGeometryEPK7QWidget(void* qthis, void* arg0);
  // proto:  int QDesktopWidget::numScreens();
extern void _ZNK14QDesktopWidget10numScreensEv(void* qthis);
  // proto:  void QDesktopWidget::~QDesktopWidget();
extern void _ZN14QDesktopWidgetD0Ev(void* qthis);
  // proto:  const QRect QDesktopWidget::screenGeometry(int screen);
extern void _ZNK14QDesktopWidget14screenGeometryEi(void* qthis, int32_t arg0);
  // proto:  const QRect QDesktopWidget::availableGeometry(const QWidget * widget);
extern void _ZNK14QDesktopWidget17availableGeometryEPK7QWidget(void* qthis, void* arg0);
  // proto:  void QDesktopWidget::QDesktopWidget(const QDesktopWidget & );
extern void* dector_ZN14QDesktopWidgetC1ERKS_(void* arg0);
extern void _ZN14QDesktopWidgetC1ERKS_(void* qthis, void* arg0);
  // proto:  int QDesktopWidget::screenNumber(const QPoint & );
extern void _ZNK14QDesktopWidget12screenNumberERK6QPoint(void* qthis, void* arg0);
  // proto:  int QDesktopWidget::screenCount();
extern void demth_ZNK14QDesktopWidget11screenCountEv(void* qthis);
  // proto:  bool QDesktopWidget::isVirtualDesktop();
extern void _ZNK14QDesktopWidget16isVirtualDesktopEv(void* qthis);
  // proto:  int QDesktopWidget::screenNumber(const QWidget * widget);
extern void _ZNK14QDesktopWidget12screenNumberEPK7QWidget(void* qthis, void* arg0);
  // proto:  int QDesktopWidget::primaryScreen();
extern void _ZNK14QDesktopWidget13primaryScreenEv(void* qthis);
  // proto:  void QDesktopWidget::QDesktopWidget();
extern void* dector_ZN14QDesktopWidgetC1Ev();
extern void _ZN14QDesktopWidgetC1Ev(void* qthis);
  // proto:  const QRect QDesktopWidget::availableGeometry(const QPoint & point);
extern void demth_ZNK14QDesktopWidget17availableGeometryERK6QPoint(void* qthis, void* arg0);
  // proto:  const QRect QDesktopWidget::availableGeometry(int screen);
extern void _ZNK14QDesktopWidget17availableGeometryEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QDesktopWidget::metaObject();
extern void _ZNK14QDesktopWidget10metaObjectEv(void* qthis);
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

  // proto:  const QRect QDesktopWidget::screenGeometry(const QPoint & point);
func (this *QDesktopWidget) screenGeometry(args ...interface{}) () {
  // screenGeometry(const class QPoint &)
  // screenGeometry(const class QWidget *)
  // screenGeometry(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget14screenGeometryERK6QPoint
    // invoke: const QRect screenGeometry(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK14QDesktopWidget14screenGeometryERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QDesktopWidget14screenGeometryEPK7QWidget
    // invoke: const QRect screenGeometry(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QDesktopWidget14screenGeometryEPK7QWidget(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QDesktopWidget14screenGeometryEi
    // invoke: const QRect screenGeometry(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK14QDesktopWidget14screenGeometryEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenGeometry", args)
  }

}

  // proto:  QWidget * QDesktopWidget::screen(int screen);
func (this *QDesktopWidget) screen(args ...interface{}) () {
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
    C._ZN14QDesktopWidget6screenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screen", args)
  }

}

  // proto:  int QDesktopWidget::numScreens();
func (this *QDesktopWidget) numScreens(args ...interface{}) () {
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
    C._ZNK14QDesktopWidget10numScreensEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "numScreens", args)
  }

}

  // proto:  void QDesktopWidget::~QDesktopWidget();
func (this *QDesktopWidget) FreeQDesktopWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QDesktopWidget", "~QDesktopWidget", args)
  }

}

  // proto:  const QRect QDesktopWidget::availableGeometry(const QWidget * widget);
func (this *QDesktopWidget) availableGeometry(args ...interface{}) () {
  // availableGeometry(const class QWidget *)
  // availableGeometry(const class QPoint &)
  // availableGeometry(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget17availableGeometryEPK7QWidget
    // invoke: const QRect availableGeometry(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QDesktopWidget17availableGeometryEPK7QWidget(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QDesktopWidget17availableGeometryERK6QPoint
    // invoke: const QRect availableGeometry(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK14QDesktopWidget17availableGeometryERK6QPoint(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK14QDesktopWidget17availableGeometryEi
    // invoke: const QRect availableGeometry(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK14QDesktopWidget17availableGeometryEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "availableGeometry", args)
  }

}

  // proto:  void QDesktopWidget::QDesktopWidget(const QDesktopWidget & );
func NewQDesktopWidget(args ...interface{}) QDesktopWidget {
  return QDesktopWidget{}
}

  // proto:  int QDesktopWidget::screenNumber(const QPoint & );
func (this *QDesktopWidget) screenNumber(args ...interface{}) () {
  // screenNumber(const class QPoint &)
  // screenNumber(const class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "const QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK14QDesktopWidget12screenNumberERK6QPoint
    // invoke: int screenNumber(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QDesktopWidget12screenNumberERK6QPoint(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK14QDesktopWidget12screenNumberEPK7QWidget
    // invoke: int screenNumber(const class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK14QDesktopWidget12screenNumberEPK7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenNumber", args)
  }

}

  // proto:  int QDesktopWidget::screenCount();
func (this *QDesktopWidget) screenCount(args ...interface{}) () {
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
    C.demth_ZNK14QDesktopWidget11screenCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "screenCount", args)
  }

}

  // proto:  bool QDesktopWidget::isVirtualDesktop();
func (this *QDesktopWidget) isVirtualDesktop(args ...interface{}) () {
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
    C._ZNK14QDesktopWidget16isVirtualDesktopEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "isVirtualDesktop", args)
  }

}

  // proto:  int QDesktopWidget::primaryScreen();
func (this *QDesktopWidget) primaryScreen(args ...interface{}) () {
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
    C._ZNK14QDesktopWidget13primaryScreenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "primaryScreen", args)
  }

}

  // proto:  const QMetaObject * QDesktopWidget::metaObject();
func (this *QDesktopWidget) metaObject(args ...interface{}) () {
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
    C._ZNK14QDesktopWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QDesktopWidget", "metaObject", args)
  }

}

// <= body block end

