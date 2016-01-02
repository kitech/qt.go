package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
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
  // proto:  QAction * QMenuBar::addAction(const QString & text);
extern void _ZN8QMenuBar9addActionERK7QString(void* qthis, void* arg0);
  // proto:  QPlatformMenuBar * QMenuBar::platformMenuBar();
extern void _ZN8QMenuBar15platformMenuBarEv(void* qthis);
  // proto:  void QMenuBar::setNativeMenuBar(bool nativeMenuBar);
extern void _ZN8QMenuBar16setNativeMenuBarEb(void* qthis, bool arg0);
  // proto:  void QMenuBar::~QMenuBar();
extern void _ZN8QMenuBarD0Ev(void* qthis);
  // proto:  QAction * QMenuBar::addMenu(QMenu * menu);
extern void _ZN8QMenuBar7addMenuEP5QMenu(void* qthis, void* arg0);
  // proto:  QSize QMenuBar::sizeHint();
extern void _ZNK8QMenuBar8sizeHintEv(void* qthis);
  // proto:  QAction * QMenuBar::actionAt(const QPoint & );
extern void _ZNK8QMenuBar8actionAtERK6QPoint(void* qthis, void* arg0);
  // proto:  const QMetaObject * QMenuBar::metaObject();
extern void _ZNK8QMenuBar10metaObjectEv(void* qthis);
  // proto:  bool QMenuBar::isNativeMenuBar();
extern void _ZNK8QMenuBar15isNativeMenuBarEv(void* qthis);
  // proto:  QAction * QMenuBar::insertSeparator(QAction * before);
extern void _ZN8QMenuBar15insertSeparatorEP7QAction(void* qthis, void* arg0);
  // proto:  QAction * QMenuBar::addSeparator();
extern void _ZN8QMenuBar12addSeparatorEv(void* qthis);
  // proto:  QSize QMenuBar::minimumSizeHint();
extern void _ZNK8QMenuBar15minimumSizeHintEv(void* qthis);
  // proto:  bool QMenuBar::isDefaultUp();
extern void _ZNK8QMenuBar11isDefaultUpEv(void* qthis);
  // proto:  void QMenuBar::QMenuBar(const QMenuBar & );
extern void* dector_ZN8QMenuBarC1ERKS_(void* arg0);
extern void _ZN8QMenuBarC1ERKS_(void* qthis, void* arg0);
  // proto:  void QMenuBar::QMenuBar(QWidget * parent);
extern void* dector_ZN8QMenuBarC1EP7QWidget(void* arg0);
extern void _ZN8QMenuBarC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QMenuBar::setActiveAction(QAction * action);
extern void _ZN8QMenuBar15setActiveActionEP7QAction(void* qthis, void* arg0);
  // proto:  void QMenuBar::clear();
extern void _ZN8QMenuBar5clearEv(void* qthis);
  // proto:  QAction * QMenuBar::activeAction();
extern void _ZNK8QMenuBar12activeActionEv(void* qthis);
  // proto:  QMenu * QMenuBar::addMenu(const QIcon & icon, const QString & title);
extern void _ZN8QMenuBar7addMenuERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  QMenu * QMenuBar::addMenu(const QString & title);
extern void _ZN8QMenuBar7addMenuERK7QString(void* qthis, void* arg0);
  // proto:  QRect QMenuBar::actionGeometry(QAction * );
extern void _ZNK8QMenuBar14actionGeometryEP7QAction(void* qthis, void* arg0);
  // proto:  QAction * QMenuBar::insertMenu(QAction * before, QMenu * menu);
extern void _ZN8QMenuBar10insertMenuEP7QActionP5QMenu(void* qthis, void* arg0, void* arg1);
  // proto:  void QMenuBar::setDefaultUp(bool );
extern void _ZN8QMenuBar12setDefaultUpEb(void* qthis, bool arg0);
  // proto:  void QMenuBar::setVisible(bool visible);
extern void _ZN8QMenuBar10setVisibleEb(void* qthis, bool arg0);
  // proto:  QAction * QMenuBar::addAction(const QString & text, const QObject * receiver, const char * member);
extern void _ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc(void* qthis, void* arg0, void* arg1, char* arg2);
  // proto:  int QMenuBar::heightForWidth(int );
extern void _ZNK8QMenuBar14heightForWidthEi(void* qthis, int arg0);
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
  qclsinst uint64 /* *mut c_void*/;
//  _hovered QMenuBar_hovered_signal;
//  _triggered QMenuBar_triggered_signal;
}

  // proto:  QAction * QMenuBar::addAction(const QString & text);
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
  case 1:
    // invoke: _ZN8QMenuBar9addActionERK7QStringPK7QObjectPKc
  default:
    qtrt.ErrorResolve("QMenuBar", "addAction", args)
  }

}

  // proto:  QPlatformMenuBar * QMenuBar::platformMenuBar();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "platformMenuBar", args)
  }

}

  // proto:  void QMenuBar::setNativeMenuBar(bool nativeMenuBar);
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
  default:
    qtrt.ErrorResolve("QMenuBar", "setNativeMenuBar", args)
  }

}

  // proto:  void QMenuBar::~QMenuBar();
func (this *QMenuBar) FreeQMenuBar(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMenuBar", "~QMenuBar", args)
  }

}

  // proto:  QAction * QMenuBar::addMenu(QMenu * menu);
func (this *QMenuBar) addMenu(args ...interface{}) () {
  // addMenu(class QMenu *)
  // addMenu(const class QIcon &, const class QString &)
  // addMenu(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QMenu{}) // "QMenu *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QMenuBar7addMenuEP5QMenu
  case 1:
    // invoke: _ZN8QMenuBar7addMenuERK5QIconRK7QString
  case 2:
    // invoke: _ZN8QMenuBar7addMenuERK7QString
  default:
    qtrt.ErrorResolve("QMenuBar", "addMenu", args)
  }

}

  // proto:  QSize QMenuBar::sizeHint();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "sizeHint", args)
  }

}

  // proto:  QAction * QMenuBar::actionAt(const QPoint & );
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
  default:
    qtrt.ErrorResolve("QMenuBar", "actionAt", args)
  }

}

  // proto:  const QMetaObject * QMenuBar::metaObject();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "metaObject", args)
  }

}

  // proto:  bool QMenuBar::isNativeMenuBar();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "isNativeMenuBar", args)
  }

}

  // proto:  QAction * QMenuBar::insertSeparator(QAction * before);
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
  default:
    qtrt.ErrorResolve("QMenuBar", "insertSeparator", args)
  }

}

  // proto:  QAction * QMenuBar::addSeparator();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "addSeparator", args)
  }

}

  // proto:  QSize QMenuBar::minimumSizeHint();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "minimumSizeHint", args)
  }

}

  // proto:  bool QMenuBar::isDefaultUp();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "isDefaultUp", args)
  }

}

  // proto:  void QMenuBar::QMenuBar(const QMenuBar & );
func NewQMenuBar(args ...interface{}) QMenuBar {
  return QMenuBar{}
}

  // proto:  void QMenuBar::setActiveAction(QAction * action);
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
  default:
    qtrt.ErrorResolve("QMenuBar", "setActiveAction", args)
  }

}

  // proto:  void QMenuBar::clear();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "clear", args)
  }

}

  // proto:  QAction * QMenuBar::activeAction();
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
  default:
    qtrt.ErrorResolve("QMenuBar", "activeAction", args)
  }

}

  // proto:  QRect QMenuBar::actionGeometry(QAction * );
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
  default:
    qtrt.ErrorResolve("QMenuBar", "actionGeometry", args)
  }

}

  // proto:  QAction * QMenuBar::insertMenu(QAction * before, QMenu * menu);
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
  default:
    qtrt.ErrorResolve("QMenuBar", "insertMenu", args)
  }

}

  // proto:  void QMenuBar::setDefaultUp(bool );
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
  default:
    qtrt.ErrorResolve("QMenuBar", "setDefaultUp", args)
  }

}

  // proto:  void QMenuBar::setVisible(bool visible);
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
  default:
    qtrt.ErrorResolve("QMenuBar", "setVisible", args)
  }

}

  // proto:  int QMenuBar::heightForWidth(int );
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
  default:
    qtrt.ErrorResolve("QMenuBar", "heightForWidth", args)
  }

}

// <= body block end

