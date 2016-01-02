package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtWidgets/qcompleter.h
// dst-file: /src/widgets/qcompleter.go
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
  // proto:  void QCompleter::QCompleter(QObject * parent);
extern void* dector_ZN10QCompleterC1EP7QObject(void* arg0);
extern void _ZN10QCompleterC1EP7QObject(void* qthis, void* arg0);
  // proto:  const QMetaObject * QCompleter::metaObject();
extern void _ZNK10QCompleter10metaObjectEv(void* qthis);
  // proto:  QAbstractItemView * QCompleter::popup();
extern void _ZNK10QCompleter5popupEv(void* qthis);
  // proto:  void QCompleter::complete(const QRect & rect);
extern void _ZN10QCompleter8completeERK5QRect(void* qthis, void* arg0);
  // proto:  void QCompleter::setCompletionRole(int role);
extern void _ZN10QCompleter17setCompletionRoleEi(void* qthis, int arg0);
  // proto:  int QCompleter::completionCount();
extern void _ZNK10QCompleter15completionCountEv(void* qthis);
  // proto:  void QCompleter::QCompleter(const QStringList & completions, QObject * parent);
extern void* dector_ZN10QCompleterC1ERK11QStringListP7QObject(void* arg0, void* arg1);
extern void _ZN10QCompleterC1ERK11QStringListP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  QModelIndex QCompleter::currentIndex();
extern void _ZNK10QCompleter12currentIndexEv(void* qthis);
  // proto:  QString QCompleter::pathFromIndex(const QModelIndex & index);
extern void _ZNK10QCompleter13pathFromIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QCompleter::setMaxVisibleItems(int maxItems);
extern void _ZN10QCompleter18setMaxVisibleItemsEi(void* qthis, int arg0);
  // proto:  int QCompleter::completionColumn();
extern void _ZNK10QCompleter16completionColumnEv(void* qthis);
  // proto:  int QCompleter::maxVisibleItems();
extern void _ZNK10QCompleter15maxVisibleItemsEv(void* qthis);
  // proto:  void QCompleter::~QCompleter();
extern void _ZN10QCompleterD0Ev(void* qthis);
  // proto:  void QCompleter::setWrapAround(bool wrap);
extern void _ZN10QCompleter13setWrapAroundEb(void* qthis, bool arg0);
  // proto:  QStringList QCompleter::splitPath(const QString & path);
extern void _ZNK10QCompleter9splitPathERK7QString(void* qthis, void* arg0);
  // proto:  QAbstractItemModel * QCompleter::model();
extern void _ZNK10QCompleter5modelEv(void* qthis);
  // proto:  QString QCompleter::currentCompletion();
extern void _ZNK10QCompleter17currentCompletionEv(void* qthis);
  // proto:  void QCompleter::setCompletionColumn(int column);
extern void _ZN10QCompleter19setCompletionColumnEi(void* qthis, int arg0);
  // proto:  void QCompleter::setCompletionPrefix(const QString & prefix);
extern void _ZN10QCompleter19setCompletionPrefixERK7QString(void* qthis, void* arg0);
  // proto:  QAbstractItemModel * QCompleter::completionModel();
extern void _ZNK10QCompleter15completionModelEv(void* qthis);
  // proto:  bool QCompleter::setCurrentRow(int row);
extern void _ZN10QCompleter13setCurrentRowEi(void* qthis, int arg0);
  // proto:  int QCompleter::currentRow();
extern void _ZNK10QCompleter10currentRowEv(void* qthis);
  // proto:  void QCompleter::setModel(QAbstractItemModel * c);
extern void _ZN10QCompleter8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  bool QCompleter::wrapAround();
extern void _ZNK10QCompleter10wrapAroundEv(void* qthis);
  // proto:  void QCompleter::QCompleter(QAbstractItemModel * model, QObject * parent);
extern void* dector_ZN10QCompleterC1EP18QAbstractItemModelP7QObject(void* arg0, void* arg1);
extern void _ZN10QCompleterC1EP18QAbstractItemModelP7QObject(void* qthis, void* arg0, void* arg1);
  // proto:  void QCompleter::setPopup(QAbstractItemView * popup);
extern void _ZN10QCompleter8setPopupEP17QAbstractItemView(void* qthis, void* arg0);
  // proto:  void QCompleter::QCompleter(const QCompleter & );
extern void* dector_ZN10QCompleterC1ERKS_(void* arg0);
extern void _ZN10QCompleterC1ERKS_(void* qthis, void* arg0);
  // proto:  QWidget * QCompleter::widget();
extern void _ZNK10QCompleter6widgetEv(void* qthis);
  // proto:  int QCompleter::completionRole();
extern void _ZNK10QCompleter14completionRoleEv(void* qthis);
  // proto:  QString QCompleter::completionPrefix();
extern void _ZNK10QCompleter16completionPrefixEv(void* qthis);
  // proto:  void QCompleter::setWidget(QWidget * widget);
extern void _ZN10QCompleter9setWidgetEP7QWidget(void* qthis, void* arg0);
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

// class sizeof(QCompleter)=1
type QCompleter struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
//  _highlighted QCompleter_highlighted_signal;
//  _activated QCompleter_activated_signal;
}

  // proto:  void QCompleter::QCompleter(QObject * parent);
func NewQCompleter(args ...interface{}) QCompleter {
  return QCompleter{}
}

  // proto:  const QMetaObject * QCompleter::metaObject();
func (this *QCompleter) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10metaObjectEv
  default:
    qtrt.ErrorResolve("QCompleter", "metaObject", args)
  }

}

  // proto:  QAbstractItemView * QCompleter::popup();
func (this *QCompleter) popup(args ...interface{}) () {
  // popup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter5popupEv
  default:
    qtrt.ErrorResolve("QCompleter", "popup", args)
  }

}

  // proto:  void QCompleter::complete(const QRect & rect);
func (this *QCompleter) complete(args ...interface{}) () {
  // complete(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter8completeERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "complete", args)
  }

}

  // proto:  void QCompleter::setCompletionRole(int role);
func (this *QCompleter) setCompletionRole(args ...interface{}) () {
  // setCompletionRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter17setCompletionRoleEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionRole", args)
  }

}

  // proto:  int QCompleter::completionCount();
func (this *QCompleter) completionCount(args ...interface{}) () {
  // completionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15completionCountEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionCount", args)
  }

}

  // proto:  QModelIndex QCompleter::currentIndex();
func (this *QCompleter) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter12currentIndexEv
  default:
    qtrt.ErrorResolve("QCompleter", "currentIndex", args)
  }

}

  // proto:  QString QCompleter::pathFromIndex(const QModelIndex & index);
func (this *QCompleter) pathFromIndex(args ...interface{}) () {
  // pathFromIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter13pathFromIndexERK11QModelIndex
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "pathFromIndex", args)
  }

}

  // proto:  void QCompleter::setMaxVisibleItems(int maxItems);
func (this *QCompleter) setMaxVisibleItems(args ...interface{}) () {
  // setMaxVisibleItems(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter18setMaxVisibleItemsEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setMaxVisibleItems", args)
  }

}

  // proto:  int QCompleter::completionColumn();
func (this *QCompleter) completionColumn(args ...interface{}) () {
  // completionColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter16completionColumnEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionColumn", args)
  }

}

  // proto:  int QCompleter::maxVisibleItems();
func (this *QCompleter) maxVisibleItems(args ...interface{}) () {
  // maxVisibleItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15maxVisibleItemsEv
  default:
    qtrt.ErrorResolve("QCompleter", "maxVisibleItems", args)
  }

}

  // proto:  void QCompleter::~QCompleter();
func (this *QCompleter) FreeQCompleter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QCompleter", "~QCompleter", args)
  }

}

  // proto:  void QCompleter::setWrapAround(bool wrap);
func (this *QCompleter) setWrapAround(args ...interface{}) () {
  // setWrapAround(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter13setWrapAroundEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setWrapAround", args)
  }

}

  // proto:  QStringList QCompleter::splitPath(const QString & path);
func (this *QCompleter) splitPath(args ...interface{}) () {
  // splitPath(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter9splitPathERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "splitPath", args)
  }

}

  // proto:  QAbstractItemModel * QCompleter::model();
func (this *QCompleter) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter5modelEv
  default:
    qtrt.ErrorResolve("QCompleter", "model", args)
  }

}

  // proto:  QString QCompleter::currentCompletion();
func (this *QCompleter) currentCompletion(args ...interface{}) () {
  // currentCompletion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter17currentCompletionEv
  default:
    qtrt.ErrorResolve("QCompleter", "currentCompletion", args)
  }

}

  // proto:  void QCompleter::setCompletionColumn(int column);
func (this *QCompleter) setCompletionColumn(args ...interface{}) () {
  // setCompletionColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter19setCompletionColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionColumn", args)
  }

}

  // proto:  void QCompleter::setCompletionPrefix(const QString & prefix);
func (this *QCompleter) setCompletionPrefix(args ...interface{}) () {
  // setCompletionPrefix(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter19setCompletionPrefixERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionPrefix", args)
  }

}

  // proto:  QAbstractItemModel * QCompleter::completionModel();
func (this *QCompleter) completionModel(args ...interface{}) () {
  // completionModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15completionModelEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionModel", args)
  }

}

  // proto:  bool QCompleter::setCurrentRow(int row);
func (this *QCompleter) setCurrentRow(args ...interface{}) () {
  // setCurrentRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter13setCurrentRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setCurrentRow", args)
  }

}

  // proto:  int QCompleter::currentRow();
func (this *QCompleter) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10currentRowEv
  default:
    qtrt.ErrorResolve("QCompleter", "currentRow", args)
  }

}

  // proto:  void QCompleter::setModel(QAbstractItemModel * c);
func (this *QCompleter) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter8setModelEP18QAbstractItemModel
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setModel", args)
  }

}

  // proto:  bool QCompleter::wrapAround();
func (this *QCompleter) wrapAround(args ...interface{}) () {
  // wrapAround()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10wrapAroundEv
  default:
    qtrt.ErrorResolve("QCompleter", "wrapAround", args)
  }

}

  // proto:  void QCompleter::setPopup(QAbstractItemView * popup);
func (this *QCompleter) setPopup(args ...interface{}) () {
  // setPopup(class QAbstractItemView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemView{}) // "QAbstractItemView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter8setPopupEP17QAbstractItemView
    var arg0 = args[0].(QAbstractItemView).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setPopup", args)
  }

}

  // proto:  QWidget * QCompleter::widget();
func (this *QCompleter) widget(args ...interface{}) () {
  // widget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter6widgetEv
  default:
    qtrt.ErrorResolve("QCompleter", "widget", args)
  }

}

  // proto:  int QCompleter::completionRole();
func (this *QCompleter) completionRole(args ...interface{}) () {
  // completionRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter14completionRoleEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionRole", args)
  }

}

  // proto:  QString QCompleter::completionPrefix();
func (this *QCompleter) completionPrefix(args ...interface{}) () {
  // completionPrefix()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter16completionPrefixEv
  default:
    qtrt.ErrorResolve("QCompleter", "completionPrefix", args)
  }

}

  // proto:  void QCompleter::setWidget(QWidget * widget);
func (this *QCompleter) setWidget(args ...interface{}) () {
  // setWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleter9setWidgetEP7QWidget
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QCompleter", "setWidget", args)
  }

}

// <= body block end

