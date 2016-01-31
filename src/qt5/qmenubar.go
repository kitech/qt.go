package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZN8QMenuBar9addActionERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenuBar::addAction(const QString & text, const QObject * receiver, const char * member);
extern void* C_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, void* arg2); // 4
  // proto:  bool QMenuBar::isDefaultUp();
extern bool C_ZNK8QMenuBar11isDefaultUpEv(void* qthis); // 4
  // proto:  QMenu * QMenuBar::addMenu(const QIcon & icon, const QString & title);
extern void* C_ZN8QMenuBar7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QMenu * QMenuBar::addMenu(const QString & title);
extern void* C_ZN8QMenuBar7addMenuERK7QString(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenuBar::addMenu(QMenu * menu);
extern void* C_ZN8QMenuBar7addMenuEP5QMenu(void* qthis, void* arg0); // 4
  // proto:  QPlatformMenuBar * QMenuBar::platformMenuBar();
extern void C_ZN8QMenuBar15platformMenuBarEv(void* qthis); // 4
  // proto:  QAction * QMenuBar::actionAt(const QPoint & );
extern void* C_ZNK8QMenuBar8actionAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QRect QMenuBar::actionGeometry(QAction * );
extern void* C_ZNK8QMenuBar14actionGeometryEP7QAction(void* qthis, void* arg0); // 4
  // proto:  QAction * QMenuBar::activeAction();
extern void* C_ZNK8QMenuBar12activeActionEv(void* qthis); // 4
  // proto:  QAction * QMenuBar::addSeparator();
extern void* C_ZN8QMenuBar12addSeparatorEv(void* qthis); // 4
  // proto:  void QMenuBar::setActiveAction(QAction * action);
extern void C_ZN8QMenuBar15setActiveActionEP7QAction(void* qthis, void* arg0); // 4
  // proto:  bool QMenuBar::isNativeMenuBar();
extern bool C_ZNK8QMenuBar15isNativeMenuBarEv(void* qthis); // 4
  // proto:  void QMenuBar::setNativeMenuBar(bool nativeMenuBar);
extern void C_ZN8QMenuBar16setNativeMenuBarEb(void* qthis, bool arg0); // 4
  // proto:  const QMetaObject * QMenuBar::metaObject();
extern void C_ZNK8QMenuBar10metaObjectEv(void* qthis); // 4
  // proto:  void QMenuBar::QMenuBar(QWidget * parent);
extern void* C_ZN8QMenuBarC2EP7QWidget(void* arg0); // 3
  // proto:  QSize QMenuBar::sizeHint();
extern void* C_ZNK8QMenuBar8sizeHintEv(void* qthis); // 4
  // proto:  QSize QMenuBar::minimumSizeHint();
extern void* C_ZNK8QMenuBar15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QMenuBar::clear();
extern void C_ZN8QMenuBar5clearEv(void* qthis); // 4
  // proto:  QAction * QMenuBar::insertSeparator(QAction * before);
extern void* C_ZN8QMenuBar15insertSeparatorEP7QAction(void* qthis, void* arg0); // 4
  // proto:  int QMenuBar::heightForWidth(int );
extern int32_t C_ZNK8QMenuBar14heightForWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QMenuBar::setDefaultUp(bool );
extern void C_ZN8QMenuBar12setDefaultUpEb(void* qthis, bool arg0); // 4
  // proto:  void QMenuBar::~QMenuBar();
extern void C_ZN8QMenuBarD2Ev(void* qthis); // 4
  // proto:  void QMenuBar::setVisible(bool visible);
extern void C_ZN8QMenuBar10setVisibleEb(void* qthis, bool arg0); // 4
  // proto:  QAction * QMenuBar::insertMenu(QAction * before, QMenu * menu);
extern void* C_ZN8QMenuBar10insertMenuEP7QActionP5QMenu(void* qthis, void* arg0, void* arg1); // 4
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
func (this *QMenuBar) Addaction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QMenuBar9addActionERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc
    // invoke: QAction * addAction(const class QString &, const class QObject *, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[1][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var ret0 = C.C_ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc(this.qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "addAction", args)
  }

  return
}

// isDefaultUp()
func (this *QMenuBar) Isdefaultup(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar11isDefaultUpEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "isDefaultUp", args)
  }

  return
}

// addMenu(const class QIcon &, const class QString &)
func (this *QMenuBar) Addmenu(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QMenuBar7addMenuERK5QIconRK7QString(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN8QMenuBar7addMenuERK7QString
    // invoke: QMenu * addMenu(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QMenuBar7addMenuERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QMenu{}) // "QMenu *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 2:
    // invoke: _ZN8QMenuBar7addMenuEP5QMenu
    // invoke: QAction * addMenu(class QMenu *)
    var arg0 = args[0].(QMenu).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QMenuBar7addMenuEP5QMenu(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "addMenu", args)
  }

  return
}

// platformMenuBar()
func (this *QMenuBar) Platformmenubar(args ...interface{}) () {
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

  return
}

// actionAt(const class QPoint &)
func (this *QMenuBar) Actionat(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar8actionAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "actionAt", args)
  }

  return
}

// actionGeometry(class QAction *)
func (this *QMenuBar) Actiongeometry(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar14actionGeometryEP7QAction(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "actionGeometry", args)
  }

  return
}

// activeAction()
func (this *QMenuBar) Activeaction(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar12activeActionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "activeAction", args)
  }

  return
}

// addSeparator()
func (this *QMenuBar) Addseparator(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QMenuBar12addSeparatorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "addSeparator", args)
  }

  return
}

// setActiveAction(class QAction *)
func (this *QMenuBar) Setactiveaction(args ...interface{}) () {
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

  return
}

// isNativeMenuBar()
func (this *QMenuBar) Isnativemenubar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar15isNativeMenuBarEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "isNativeMenuBar", args)
  }

  return
}

// setNativeMenuBar(_Bool)
func (this *QMenuBar) Setnativemenubar(args ...interface{}) () {
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

  return
}

// metaObject()
func (this *QMenuBar) Metaobject(args ...interface{}) () {
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

  return
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
func (this *QMenuBar) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "sizeHint", args)
  }

  return
}

// minimumSizeHint()
func (this *QMenuBar) Minimumsizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "minimumSizeHint", args)
  }

  return
}

// clear()
func (this *QMenuBar) Clear(args ...interface{}) () {
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

  return
}

// insertSeparator(class QAction *)
func (this *QMenuBar) Insertseparator(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QMenuBar15insertSeparatorEP7QAction(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "insertSeparator", args)
  }

  return
}

// heightForWidth(int)
func (this *QMenuBar) Heightforwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK8QMenuBar14heightForWidthEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "heightForWidth", args)
  }

  return
}

// setDefaultUp(_Bool)
func (this *QMenuBar) Setdefaultup(args ...interface{}) () {
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

  return
}

// ~QMenuBar()
func (this *QMenuBar) Freeqmenubar(args ...interface{}) () {
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

  return
}

// setVisible(_Bool)
func (this *QMenuBar) Setvisible(args ...interface{}) () {
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

  return
}

// insertMenu(class QAction *, class QMenu *)
func (this *QMenuBar) Insertmenu(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN8QMenuBar10insertMenuEP7QActionP5QMenu(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAction{}) // "QAction *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QMenuBar", "insertMenu", args)
  }

  return
}

// <= body block end

