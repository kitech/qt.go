package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qmenubar.h
// dst-file: /src/widgets/qmenubar.go
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
  // proto:  QAction * QMenuBar::addAction(const QString & text);
extern void C_ZN8QMenuBar9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenuBar::addAction(const QString & text, const QObject * receiver, const char * member);
extern void C_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, unsigned char* arg2); // 4
  // proto:  bool QMenuBar::isDefaultUp();
extern void C_ZNK8QMenuBar11isDefaultUpEv(void* qthis); // 4
  // proto:  QMenu * QMenuBar::addMenu(const QIcon & icon, const QString & title);
extern void C_ZN8QMenuBar7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QMenu * QMenuBar::addMenu(const QString & title);
extern void C_ZN8QMenuBar7addMenuERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenuBar::addMenu(QMenu * menu);
extern void C_ZN8QMenuBar7addMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  QPlatformMenuBar * QMenuBar::platformMenuBar();
extern void C_ZN8QMenuBar15platformMenuBarEv(void* qthis); // 4
  // proto:  QAction * QMenuBar::actionAt(const QPoint & );
extern void C_ZNK8QMenuBar8actionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QRect QMenuBar::actionGeometry(QAction * );
extern void C_ZNK8QMenuBar14actionGeometryEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenuBar::activeAction();
extern void C_ZNK8QMenuBar12activeActionEv(void* qthis); // 4
  // proto:  QAction * QMenuBar::addSeparator();
extern void C_ZN8QMenuBar12addSeparatorEv(void* qthis); // 4
  // proto:  void QMenuBar::setActiveAction(QAction * action);
extern void C_ZN8QMenuBar15setActiveActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  bool QMenuBar::isNativeMenuBar();
extern void C_ZNK8QMenuBar15isNativeMenuBarEv(void* qthis); // 4
  // proto:  void QMenuBar::setNativeMenuBar(bool nativeMenuBar);
extern void C_ZN8QMenuBar16setNativeMenuBarEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QMenuBar::metaObject();
extern void C_ZNK8QMenuBar10metaObjectEv(void* qthis); // 4
  // proto:  void QMenuBar::QMenuBar(QWidget * parent);
extern void* C_ZN8QMenuBarC2EP7QWidget(void* arg0); // 3
  // proto:  QSize QMenuBar::sizeHint();
extern void C_ZNK8QMenuBar8sizeHintEv(void* qthis); // 4
  // proto:  QSize QMenuBar::minimumSizeHint();
extern void C_ZNK8QMenuBar15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QMenuBar::clear();
extern void C_ZN8QMenuBar5clearEv(void* qthis); // 4
  // proto:  QAction * QMenuBar::insertSeparator(QAction * before);
extern void C_ZN8QMenuBar15insertSeparatorEP7QAction(void* qthis, void* arg0); // 4
  // proto:  int QMenuBar::heightForWidth(int );
extern void C_ZNK8QMenuBar14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMenuBar::setDefaultUp(bool );
extern void C_ZN8QMenuBar12setDefaultUpEb(void* qthis, bool arg0); // 4
  // proto:  void QMenuBar::~QMenuBar();
extern void C_ZN8QMenuBarD2Ev(void* qthis); // 4
  // proto:  void QMenuBar::setVisible(bool visible);
extern void C_ZN8QMenuBar10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QAction * QMenuBar::insertMenu(QAction * before, QMenu * menu);
extern void C_ZN8QMenuBar10insertMenuEP7QActionP5QMenu(void* qthis, void* arg0, void* arg1); // 4
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

// class sizeof(QMenuBar)=1
type QMenuBar struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _hovered QMenuBar_hovered_signal;
//  _triggered QMenuBar_triggered_signal;
}

// addAction(const class QString &)
func (this *QMenuBar) addAction(args ...interface{}) () {
  // addAction(const class QString &)
  // addAction(const class QString &, const class QObject *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "const QObject *"
  vtys[1][2] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar9addActionERK7QString
    // invoke: QAction * addAction(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN8QMenuBar9addActionERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var ret = C.C_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "addAction", args)
  }

}

// isDefaultUp()
func (this *QMenuBar) isDefaultUp(args ...interface{}) () {
  // isDefaultUp()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar11isDefaultUpEv
    // invoke: bool isDefaultUp()
    var ret = C.C_ZNK8QMenuBar11isDefaultUpEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "isDefaultUp", args)
  }

}

// addMenu(const class QIcon &, const class QString &)
func (this *QMenuBar) addMenu(args ...interface{}) () {
  // addMenu(const class QIcon &, const class QString &)
  // addMenu(const class QString &)
  // addMenu(class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar7addMenuERK5QIconRK7QString
    // invoke: QMenu * addMenu(const class QIcon &, const class QString &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN8QMenuBar7addMenuERK5QIconRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN8QMenuBar7addMenuERK7QString
    // invoke: QMenu * addMenu(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN8QMenuBar7addMenuERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 2:
    // invoke: _ZN8QMenuBar7addMenuEP5QMenu
    // invoke: QAction * addMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN8QMenuBar7addMenuEP5QMenu(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "addMenu", args)
  }

}

// platformMenuBar()
func (this *QMenuBar) platformMenuBar(args ...interface{}) () {
  // platformMenuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar15platformMenuBarEv
    // invoke: QPlatformMenuBar * platformMenuBar()
    C.C_ZN8QMenuBar15platformMenuBarEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenuBar", "platformMenuBar", args)
  }

}

// actionAt(const class QPoint &)
func (this *QMenuBar) actionAt(args ...interface{}) () {
  // actionAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar8actionAtERK6QPoint
    // invoke: QAction * actionAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK8QMenuBar8actionAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "actionAt", args)
  }

}

// actionGeometry(class QAction *)
func (this *QMenuBar) actionGeometry(args ...interface{}) () {
  // actionGeometry(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar14actionGeometryEP7QAction
    // invoke: QRect actionGeometry(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK8QMenuBar14actionGeometryEP7QAction(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "actionGeometry", args)
  }

}

// activeAction()
func (this *QMenuBar) activeAction(args ...interface{}) () {
  // activeAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar12activeActionEv
    // invoke: QAction * activeAction()
    var ret = C.C_ZNK8QMenuBar12activeActionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "activeAction", args)
  }

}

// addSeparator()
func (this *QMenuBar) addSeparator(args ...interface{}) () {
  // addSeparator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar12addSeparatorEv
    // invoke: QAction * addSeparator()
    var ret = C.C_ZN8QMenuBar12addSeparatorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "addSeparator", args)
  }

}

// setActiveAction(class QAction *)
func (this *QMenuBar) setActiveAction(args ...interface{}) () {
  // setActiveAction(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar15setActiveActionEP7QAction
    // invoke: void setActiveAction(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QMenuBar15setActiveActionEP7QAction(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenuBar", "setActiveAction", args)
  }

}

// isNativeMenuBar()
func (this *QMenuBar) isNativeMenuBar(args ...interface{}) () {
  // isNativeMenuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar15isNativeMenuBarEv
    // invoke: bool isNativeMenuBar()
    var ret = C.C_ZNK8QMenuBar15isNativeMenuBarEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "isNativeMenuBar", args)
  }

}

// setNativeMenuBar(_Bool)
func (this *QMenuBar) setNativeMenuBar(args ...interface{}) () {
  // setNativeMenuBar(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar16setNativeMenuBarEb
    // invoke: void setNativeMenuBar(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMenuBar16setNativeMenuBarEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenuBar", "setNativeMenuBar", args)
  }

}

// metaObject()
func (this *QMenuBar) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK8QMenuBar10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenuBar", "metaObject", args)
  }

}

// QMenuBar(class QWidget *)
func NewQMenuBar(args ...interface{}) *QMenuBar {
  // QMenuBar(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBarC1EP7QWidget
    // invoke: void QMenuBar(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QMenuBarC2EP7QWidget(arg0)
    return &QMenuBar{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMenuBar", "QMenuBar", args)
  }

  return nil // QMenuBar{qclsinst:qthis}
}

// sizeHint()
func (this *QMenuBar) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar8sizeHintEv
    // invoke: QSize sizeHint()
    var ret = C.C_ZNK8QMenuBar8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "sizeHint", args)
  }

}

// minimumSizeHint()
func (this *QMenuBar) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    var ret = C.C_ZNK8QMenuBar15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "minimumSizeHint", args)
  }

}

// clear()
func (this *QMenuBar) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar5clearEv
    // invoke: void clear()
    C.C_ZN8QMenuBar5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenuBar", "clear", args)
  }

}

// insertSeparator(class QAction *)
func (this *QMenuBar) insertSeparator(args ...interface{}) () {
  // insertSeparator(class QAction *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar15insertSeparatorEP7QAction
    // invoke: QAction * insertSeparator(class QAction *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN8QMenuBar15insertSeparatorEP7QAction(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "insertSeparator", args)
  }

}

// heightForWidth(int)
func (this *QMenuBar) heightForWidth(args ...interface{}) () {
  // heightForWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QMenuBar14heightForWidthEi
    // invoke: int heightForWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK8QMenuBar14heightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "heightForWidth", args)
  }

}

// setDefaultUp(_Bool)
func (this *QMenuBar) setDefaultUp(args ...interface{}) () {
  // setDefaultUp(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar12setDefaultUpEb
    // invoke: void setDefaultUp(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMenuBar12setDefaultUpEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenuBar", "setDefaultUp", args)
  }

}

// ~QMenuBar()
func (this *QMenuBar) FreeQMenuBar(args ...interface{}) () {
  // ~QMenuBar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBarD0Ev
    // invoke: void ~QMenuBar()
    C.C_ZN8QMenuBarD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMenuBar", "~QMenuBar", args)
  }

}

// setVisible(_Bool)
func (this *QMenuBar) setVisible(args ...interface{}) () {
  // setVisible(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar10setVisibleEb
    // invoke: void setVisible(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN8QMenuBar10setVisibleEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMenuBar", "setVisible", args)
  }

}

// insertMenu(class QAction *, class QMenu *)
func (this *QMenuBar) insertMenu(args ...interface{}) () {
  // insertMenu(class QAction *, class QMenu *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAction{}) // "QAction *"
  vtys[0][1] = reflect.TypeOf(QMenu{}) // "QMenu *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar10insertMenuEP7QActionP5QMenu
    // invoke: QAction * insertMenu(class QAction *, class QMenu *)
    var arg0 = args[0].(QAction).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QMenu).qclsinst
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN8QMenuBar10insertMenuEP7QActionP5QMenu(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMenuBar", "insertMenu", args)
  }

}

// <= body block end

