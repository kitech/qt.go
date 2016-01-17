package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QCompleter::setCurrentRow(int row);
extern void _ZN10QCompleter13setCurrentRowEi(void* qthis, int32_t arg0); // 4
  // proto:  QAbstractItemModel * QCompleter::completionModel();
extern void _ZNK10QCompleter15completionModelEv(void* qthis); // 4
  // proto:  int QCompleter::completionRole();
extern void _ZNK10QCompleter14completionRoleEv(void* qthis); // 4
  // proto:  int QCompleter::maxVisibleItems();
extern void _ZNK10QCompleter15maxVisibleItemsEv(void* qthis); // 4
  // proto:  bool QCompleter::wrapAround();
extern void _ZNK10QCompleter10wrapAroundEv(void* qthis); // 4
  // proto:  QString QCompleter::completionPrefix();
extern void _ZNK10QCompleter16completionPrefixEv(void* qthis); // 4
  // proto:  void QCompleter::setCompletionColumn(int column);
extern void _ZN10QCompleter19setCompletionColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  QStringList QCompleter::splitPath(const QString & path);
extern void _ZNK10QCompleter9splitPathERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCompleter::setWidget(QWidget * widget);
extern void _ZN10QCompleter9setWidgetEP7QWidget(void* qthis, void* arg0); // 4
  // proto:  void QCompleter::setPopup(QAbstractItemView * popup);
extern void _ZN10QCompleter8setPopupEP17QAbstractItemView(void* qthis, void* arg0); // 4
  // proto:  void QCompleter::setWrapAround(bool wrap);
extern void _ZN10QCompleter13setWrapAroundEb(void* qthis, bool arg0); // 4
  // proto:  void QCompleter::setCompletionPrefix(const QString & prefix);
extern void _ZN10QCompleter19setCompletionPrefixERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QCompleter::~QCompleter();
extern void _ZN10QCompleterD2Ev(void* qthis); // 4
  // proto:  QCompleter::CompletionMode QCompleter::completionMode();
extern void _ZNK10QCompleter14completionModeEv(void* qthis); // 4
  // proto:  Qt::CaseSensitivity QCompleter::caseSensitivity();
extern void _ZNK10QCompleter15caseSensitivityEv(void* qthis); // 4
  // proto:  QWidget * QCompleter::widget();
extern void _ZNK10QCompleter6widgetEv(void* qthis); // 4
  // proto:  QAbstractItemView * QCompleter::popup();
extern void _ZNK10QCompleter5popupEv(void* qthis); // 4
  // proto:  void QCompleter::complete(const QRect & rect);
extern void _ZN10QCompleter8completeERK5QRect(void* qthis, void* arg0); // 4
  // proto:  int QCompleter::completionColumn();
extern void _ZNK10QCompleter16completionColumnEv(void* qthis); // 4
  // proto:  void QCompleter::setModel(QAbstractItemModel * c);
extern void _ZN10QCompleter8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  void QCompleter::QCompleter(QObject * parent);
extern void _ZN10QCompleterC2EP7QObject(void* qthis, void* arg0); // 3
  // proto:  void QCompleter::QCompleter(QAbstractItemModel * model, QObject * parent);
extern void _ZN10QCompleterC2EP18QAbstractItemModelP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  void QCompleter::QCompleter(const QStringList & completions, QObject * parent);
extern void _ZN10QCompleterC2ERK11QStringListP7QObject(void* qthis, void* arg0, void* arg1); // 3
  // proto:  QCompleter::ModelSorting QCompleter::modelSorting();
extern void _ZNK10QCompleter12modelSortingEv(void* qthis); // 4
  // proto:  int QCompleter::currentRow();
extern void _ZNK10QCompleter10currentRowEv(void* qthis); // 4
  // proto:  const QMetaObject * QCompleter::metaObject();
extern void _ZNK10QCompleter10metaObjectEv(void* qthis); // 4
  // proto:  QString QCompleter::currentCompletion();
extern void _ZNK10QCompleter17currentCompletionEv(void* qthis); // 4
  // proto:  void QCompleter::setMaxVisibleItems(int maxItems);
extern void _ZN10QCompleter18setMaxVisibleItemsEi(void* qthis, int32_t arg0); // 4
  // proto:  int QCompleter::completionCount();
extern void _ZNK10QCompleter15completionCountEv(void* qthis); // 4
  // proto:  void QCompleter::setCompletionRole(int role);
extern void _ZN10QCompleter17setCompletionRoleEi(void* qthis, int32_t arg0); // 4
  // proto:  QModelIndex QCompleter::currentIndex();
extern void _ZNK10QCompleter12currentIndexEv(void* qthis); // 4
  // proto:  QString QCompleter::pathFromIndex(const QModelIndex & index);
extern void _ZNK10QCompleter13pathFromIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::MatchFlags QCompleter::filterMode();
extern void _ZNK10QCompleter10filterModeEv(void* qthis); // 4
  // proto:  QAbstractItemModel * QCompleter::model();
extern void _ZNK10QCompleter5modelEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
//  _highlighted QCompleter_highlighted_signal;
//  _activated QCompleter_activated_signal;
}

// setCurrentRow(int)
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
    // invoke: bool setCurrentRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter13setCurrentRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setCurrentRow", args)
  }

}

// completionModel()
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
    // invoke: QAbstractItemModel * completionModel()
    C._ZNK10QCompleter15completionModelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "completionModel", args)
  }

}

// completionRole()
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
    // invoke: int completionRole()
    C._ZNK10QCompleter14completionRoleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "completionRole", args)
  }

}

// maxVisibleItems()
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
    // invoke: int maxVisibleItems()
    C._ZNK10QCompleter15maxVisibleItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "maxVisibleItems", args)
  }

}

// wrapAround()
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
    // invoke: bool wrapAround()
    C._ZNK10QCompleter10wrapAroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "wrapAround", args)
  }

}

// completionPrefix()
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
    // invoke: QString completionPrefix()
    C._ZNK10QCompleter16completionPrefixEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "completionPrefix", args)
  }

}

// setCompletionColumn(int)
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
    // invoke: void setCompletionColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter19setCompletionColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionColumn", args)
  }

}

// splitPath(const class QString &)
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
    // invoke: QStringList splitPath(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QCompleter9splitPathERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "splitPath", args)
  }

}

// setWidget(class QWidget *)
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
    // invoke: void setWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter9setWidgetEP7QWidget(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setWidget", args)
  }

}

// setPopup(class QAbstractItemView *)
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
    // invoke: void setPopup(class QAbstractItemView *)
    var arg0 = args[0].(QAbstractItemView).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter8setPopupEP17QAbstractItemView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setPopup", args)
  }

}

// setWrapAround(_Bool)
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
    // invoke: void setWrapAround(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter13setWrapAroundEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setWrapAround", args)
  }

}

// setCompletionPrefix(const class QString &)
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
    // invoke: void setCompletionPrefix(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter19setCompletionPrefixERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionPrefix", args)
  }

}

// ~QCompleter()
func (this *QCompleter) FreeQCompleter(args ...interface{}) () {
  // ~QCompleter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleterD0Ev
    // invoke: void ~QCompleter()
    C._ZN10QCompleterD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "~QCompleter", args)
  }

}

// completionMode()
func (this *QCompleter) completionMode(args ...interface{}) () {
  // completionMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter14completionModeEv
    // invoke: QCompleter::CompletionMode completionMode()
    C._ZNK10QCompleter14completionModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "completionMode", args)
  }

}

// caseSensitivity()
func (this *QCompleter) caseSensitivity(args ...interface{}) () {
  // caseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter15caseSensitivityEv
    // invoke: Qt::CaseSensitivity caseSensitivity()
    C._ZNK10QCompleter15caseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "caseSensitivity", args)
  }

}

// widget()
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
    // invoke: QWidget * widget()
    C._ZNK10QCompleter6widgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "widget", args)
  }

}

// popup()
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
    // invoke: QAbstractItemView * popup()
    C._ZNK10QCompleter5popupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "popup", args)
  }

}

// complete(const class QRect &)
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
    // invoke: void complete(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter8completeERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "complete", args)
  }

}

// completionColumn()
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
    // invoke: int completionColumn()
    C._ZNK10QCompleter16completionColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "completionColumn", args)
  }

}

// setModel(class QAbstractItemModel *)
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
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setModel", args)
  }

}

// QCompleter(class QObject *)
func NewQCompleter(args ...interface{}) QCompleter {
  // QCompleter(class QObject *)
  // QCompleter(class QAbstractItemModel *, class QObject *)
  // QCompleter(const class QStringList &, class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"
  vtys[1][1] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[2][1] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QCompleterC1EP7QObject
    // invoke: void QCompleter(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QCompleterC2EP7QObject(qthis, arg0)
  case 1:
    // invoke: _ZN10QCompleterC1EP18QAbstractItemModelP7QObject
    // invoke: void QCompleter(class QAbstractItemModel *, class QObject *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QCompleterC2EP18QAbstractItemModelP7QObject(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN10QCompleterC1ERK11QStringListP7QObject
    // invoke: void QCompleter(const class QStringList &, class QObject *)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QObject).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QCompleterC2ERK11QStringListP7QObject(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QCompleter", "QCompleter", args)
  }

  return QCompleter{}
}

// modelSorting()
func (this *QCompleter) modelSorting(args ...interface{}) () {
  // modelSorting()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter12modelSortingEv
    // invoke: QCompleter::ModelSorting modelSorting()
    C._ZNK10QCompleter12modelSortingEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "modelSorting", args)
  }

}

// currentRow()
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
    // invoke: int currentRow()
    C._ZNK10QCompleter10currentRowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "currentRow", args)
  }

}

// metaObject()
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
    // invoke: const QMetaObject * metaObject()
    C._ZNK10QCompleter10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "metaObject", args)
  }

}

// currentCompletion()
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
    // invoke: QString currentCompletion()
    C._ZNK10QCompleter17currentCompletionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "currentCompletion", args)
  }

}

// setMaxVisibleItems(int)
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
    // invoke: void setMaxVisibleItems(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter18setMaxVisibleItemsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setMaxVisibleItems", args)
  }

}

// completionCount()
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
    // invoke: int completionCount()
    C._ZNK10QCompleter15completionCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "completionCount", args)
  }

}

// setCompletionRole(int)
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
    // invoke: void setCompletionRole(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QCompleter17setCompletionRoleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "setCompletionRole", args)
  }

}

// currentIndex()
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
    // invoke: QModelIndex currentIndex()
    C._ZNK10QCompleter12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "currentIndex", args)
  }

}

// pathFromIndex(const class QModelIndex &)
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
    // invoke: QString pathFromIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QCompleter13pathFromIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCompleter", "pathFromIndex", args)
  }

}

// filterMode()
func (this *QCompleter) filterMode(args ...interface{}) () {
  // filterMode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QCompleter10filterModeEv
    // invoke: Qt::MatchFlags filterMode()
    C._ZNK10QCompleter10filterModeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "filterMode", args)
  }

}

// model()
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
    // invoke: QAbstractItemModel * model()
    C._ZNK10QCompleter5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCompleter", "model", args)
  }

}

// <= body block end

