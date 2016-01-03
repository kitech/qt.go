package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtWidgets/qcombobox.h
// dst-file: /src/widgets/qcombobox.go
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
  // proto:  void QComboBox::setModel(QAbstractItemModel * model);
extern void _ZN9QComboBox8setModelEP18QAbstractItemModel(void* qthis, void* arg0);
  // proto:  void QComboBox::clearEditText();
extern void _ZN9QComboBox13clearEditTextEv(void* qthis);
  // proto:  void QComboBox::setAutoCompletion(bool enable);
extern void _ZN9QComboBox17setAutoCompletionEb(void* qthis, bool arg0);
  // proto:  void QComboBox::setFrame(bool );
extern void _ZN9QComboBox8setFrameEb(void* qthis, bool arg0);
  // proto:  void QComboBox::setIconSize(const QSize & size);
extern void _ZN9QComboBox11setIconSizeERK5QSize(void* qthis, void* arg0);
  // proto:  void QComboBox::setView(QAbstractItemView * itemView);
extern void _ZN9QComboBox7setViewEP17QAbstractItemView(void* qthis, void* arg0);
  // proto:  QAbstractItemView * QComboBox::view();
extern void _ZNK9QComboBox4viewEv(void* qthis);
  // proto:  QSize QComboBox::minimumSizeHint();
extern void _ZNK9QComboBox15minimumSizeHintEv(void* qthis);
  // proto:  void QComboBox::clear();
extern void _ZN9QComboBox5clearEv(void* qthis);
  // proto:  int QComboBox::maxCount();
extern void _ZNK9QComboBox8maxCountEv(void* qthis);
  // proto:  void QComboBox::addItem(const QIcon & icon, const QString & text, const QVariant & userData);
extern void demth_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  void QComboBox::insertItems(int index, const QStringList & texts);
extern void _ZN9QComboBox11insertItemsEiRK11QStringList(void* qthis, int32_t arg0, void* arg1);
  // proto:  QSize QComboBox::iconSize();
extern void _ZNK9QComboBox8iconSizeEv(void* qthis);
  // proto:  QModelIndex QComboBox::rootModelIndex();
extern void _ZNK9QComboBox14rootModelIndexEv(void* qthis);
  // proto:  void QComboBox::setEditable(bool editable);
extern void _ZN9QComboBox11setEditableEb(void* qthis, bool arg0);
  // proto:  void QComboBox::setItemIcon(int index, const QIcon & icon);
extern void _ZN9QComboBox11setItemIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1);
  // proto:  bool QComboBox::autoCompletion();
extern void _ZNK9QComboBox14autoCompletionEv(void* qthis);
  // proto:  QVariant QComboBox::currentData(int role);
extern void _ZNK9QComboBox11currentDataEi(void* qthis, int32_t arg0);
  // proto:  bool QComboBox::hasFrame();
extern void _ZNK9QComboBox8hasFrameEv(void* qthis);
  // proto:  const QValidator * QComboBox::validator();
extern void _ZNK9QComboBox9validatorEv(void* qthis);
  // proto:  QString QComboBox::itemText(int index);
extern void _ZNK9QComboBox8itemTextEi(void* qthis, int32_t arg0);
  // proto:  void QComboBox::setItemData(int index, const QVariant & value, int role);
extern void _ZN9QComboBox11setItemDataEiRK8QVarianti(void* qthis, int32_t arg0, void* arg1, int32_t arg2);
  // proto:  void QComboBox::hidePopup();
extern void _ZN9QComboBox9hidePopupEv(void* qthis);
  // proto:  void QComboBox::insertItem(int index, const QIcon & icon, const QString & text, const QVariant & userData);
extern void _ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3);
  // proto:  void QComboBox::setCurrentText(const QString & text);
extern void _ZN9QComboBox14setCurrentTextERK7QString(void* qthis, void* arg0);
  // proto:  int QComboBox::modelColumn();
extern void _ZNK9QComboBox11modelColumnEv(void* qthis);
  // proto:  QSize QComboBox::sizeHint();
extern void _ZNK9QComboBox8sizeHintEv(void* qthis);
  // proto:  QVariant QComboBox::itemData(int index, int role);
extern void _ZNK9QComboBox8itemDataEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QComboBox::setCompleter(QCompleter * c);
extern void _ZN9QComboBox12setCompleterEP10QCompleter(void* qthis, void* arg0);
  // proto:  int QComboBox::maxVisibleItems();
extern void _ZNK9QComboBox15maxVisibleItemsEv(void* qthis);
  // proto:  void QComboBox::QComboBox(QWidget * parent);
extern void* dector_ZN9QComboBoxC1EP7QWidget(void* arg0);
extern void _ZN9QComboBoxC1EP7QWidget(void* qthis, void* arg0);
  // proto:  void QComboBox::setCurrentIndex(int index);
extern void _ZN9QComboBox15setCurrentIndexEi(void* qthis, int32_t arg0);
  // proto:  void QComboBox::QComboBox(const QComboBox & );
extern void* dector_ZN9QComboBoxC1ERKS_(void* arg0);
extern void _ZN9QComboBoxC1ERKS_(void* qthis, void* arg0);
  // proto:  void QComboBox::setRootModelIndex(const QModelIndex & index);
extern void _ZN9QComboBox17setRootModelIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QComboBox::setEditText(const QString & text);
extern void _ZN9QComboBox11setEditTextERK7QString(void* qthis, void* arg0);
  // proto:  void QComboBox::addItem(const QString & text, const QVariant & userData);
extern void demth_ZN9QComboBox7addItemERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  QCompleter * QComboBox::completer();
extern void _ZNK9QComboBox9completerEv(void* qthis);
  // proto:  void QComboBox::removeItem(int index);
extern void _ZN9QComboBox10removeItemEi(void* qthis, int32_t arg0);
  // proto:  int QComboBox::count();
extern void _ZNK9QComboBox5countEv(void* qthis);
  // proto:  void QComboBox::setItemDelegate(QAbstractItemDelegate * delegate);
extern void _ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0);
  // proto:  void QComboBox::addItems(const QStringList & texts);
extern void demth_ZN9QComboBox8addItemsERK11QStringList(void* qthis, void* arg0);
  // proto:  void QComboBox::setMinimumContentsLength(int characters);
extern void _ZN9QComboBox24setMinimumContentsLengthEi(void* qthis, int32_t arg0);
  // proto:  bool QComboBox::duplicatesEnabled();
extern void _ZNK9QComboBox17duplicatesEnabledEv(void* qthis);
  // proto:  void QComboBox::~QComboBox();
extern void _ZN9QComboBoxD0Ev(void* qthis);
  // proto:  QAbstractItemModel * QComboBox::model();
extern void _ZNK9QComboBox5modelEv(void* qthis);
  // proto:  int QComboBox::minimumContentsLength();
extern void _ZNK9QComboBox21minimumContentsLengthEv(void* qthis);
  // proto:  bool QComboBox::isEditable();
extern void _ZNK9QComboBox10isEditableEv(void* qthis);
  // proto:  void QComboBox::setMaxCount(int max);
extern void _ZN9QComboBox11setMaxCountEi(void* qthis, int32_t arg0);
  // proto:  int QComboBox::currentIndex();
extern void _ZNK9QComboBox12currentIndexEv(void* qthis);
  // proto:  void QComboBox::setDuplicatesEnabled(bool enable);
extern void _ZN9QComboBox20setDuplicatesEnabledEb(void* qthis, bool arg0);
  // proto:  QString QComboBox::currentText();
extern void _ZNK9QComboBox11currentTextEv(void* qthis);
  // proto:  void QComboBox::showPopup();
extern void _ZN9QComboBox9showPopupEv(void* qthis);
  // proto:  QLineEdit * QComboBox::lineEdit();
extern void _ZNK9QComboBox8lineEditEv(void* qthis);
  // proto:  QAbstractItemDelegate * QComboBox::itemDelegate();
extern void _ZNK9QComboBox12itemDelegateEv(void* qthis);
  // proto:  void QComboBox::setMaxVisibleItems(int maxItems);
extern void _ZN9QComboBox18setMaxVisibleItemsEi(void* qthis, int32_t arg0);
  // proto:  bool QComboBox::event(QEvent * event);
extern void _ZN9QComboBox5eventEP6QEvent(void* qthis, void* arg0);
  // proto:  void QComboBox::setModelColumn(int visibleColumn);
extern void _ZN9QComboBox14setModelColumnEi(void* qthis, int32_t arg0);
  // proto:  void QComboBox::setItemText(int index, const QString & text);
extern void _ZN9QComboBox11setItemTextEiRK7QString(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QComboBox::setLineEdit(QLineEdit * edit);
extern void _ZN9QComboBox11setLineEditEP9QLineEdit(void* qthis, void* arg0);
  // proto:  QIcon QComboBox::itemIcon(int index);
extern void _ZNK9QComboBox8itemIconEi(void* qthis, int32_t arg0);
  // proto:  void QComboBox::insertItem(int index, const QString & text, const QVariant & userData);
extern void demth_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant(void* qthis, int32_t arg0, void* arg1, void* arg2);
  // proto:  void QComboBox::setValidator(const QValidator * v);
extern void _ZN9QComboBox12setValidatorEPK10QValidator(void* qthis, void* arg0);
  // proto:  void QComboBox::insertSeparator(int index);
extern void _ZN9QComboBox15insertSeparatorEi(void* qthis, int32_t arg0);
  // proto:  const QMetaObject * QComboBox::metaObject();
extern void _ZNK9QComboBox10metaObjectEv(void* qthis);
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

// class sizeof(QComboBox)=1
type QComboBox struct {
  /*qbase*/ QWidget;
  qclsinst unsafe.Pointer /* *C.void */;
//  _currentIndexChanged QComboBox_currentIndexChanged_signal;
//  _currentTextChanged QComboBox_currentTextChanged_signal;
//  _highlighted QComboBox_highlighted_signal;
//  _activated QComboBox_activated_signal;
//  _editTextChanged QComboBox_editTextChanged_signal;
}

  // proto:  void QComboBox::setModel(QAbstractItemModel * model);
func (this *QComboBox) setModel(args ...interface{}) () {
  // setModel(class QAbstractItemModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemModel{}) // "QAbstractItemModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox8setModelEP18QAbstractItemModel
    // invoke: void setModel(class QAbstractItemModel *)
    var arg0 = args[0].(QAbstractItemModel).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setModel", args)
  }

}

  // proto:  void QComboBox::clearEditText();
func (this *QComboBox) clearEditText(args ...interface{}) () {
  // clearEditText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox13clearEditTextEv
    // invoke: void clearEditText()
    C._ZN9QComboBox13clearEditTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "clearEditText", args)
  }

}

  // proto:  void QComboBox::setAutoCompletion(bool enable);
func (this *QComboBox) setAutoCompletion(args ...interface{}) () {
  // setAutoCompletion(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox17setAutoCompletionEb
    // invoke: void setAutoCompletion(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox17setAutoCompletionEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setAutoCompletion", args)
  }

}

  // proto:  void QComboBox::setFrame(bool );
func (this *QComboBox) setFrame(args ...interface{}) () {
  // setFrame(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox8setFrameEb
    // invoke: void setFrame(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox8setFrameEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setFrame", args)
  }

}

  // proto:  void QComboBox::setIconSize(const QSize & size);
func (this *QComboBox) setIconSize(args ...interface{}) () {
  // setIconSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setIconSizeERK5QSize
    // invoke: void setIconSize(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setIconSize", args)
  }

}

  // proto:  void QComboBox::setView(QAbstractItemView * itemView);
func (this *QComboBox) setView(args ...interface{}) () {
  // setView(class QAbstractItemView *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemView{}) // "QAbstractItemView *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox7setViewEP17QAbstractItemView
    // invoke: void setView(class QAbstractItemView *)
    var arg0 = args[0].(QAbstractItemView).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox7setViewEP17QAbstractItemView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setView", args)
  }

}

  // proto:  QAbstractItemView * QComboBox::view();
func (this *QComboBox) view(args ...interface{}) () {
  // view()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox4viewEv
    // invoke: QAbstractItemView * view()
    C._ZNK9QComboBox4viewEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "view", args)
  }

}

  // proto:  QSize QComboBox::minimumSizeHint();
func (this *QComboBox) minimumSizeHint(args ...interface{}) () {
  // minimumSizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox15minimumSizeHintEv
    // invoke: QSize minimumSizeHint()
    C._ZNK9QComboBox15minimumSizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "minimumSizeHint", args)
  }

}

  // proto:  void QComboBox::clear();
func (this *QComboBox) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox5clearEv
    // invoke: void clear()
    C._ZN9QComboBox5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "clear", args)
  }

}

  // proto:  int QComboBox::maxCount();
func (this *QComboBox) maxCount(args ...interface{}) () {
  // maxCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8maxCountEv
    // invoke: int maxCount()
    C._ZNK9QComboBox8maxCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "maxCount", args)
  }

}

  // proto:  void QComboBox::addItem(const QIcon & icon, const QString & text, const QVariant & userData);
func (this *QComboBox) addItem(args ...interface{}) () {
  // addItem(const class QIcon &, const class QString &, const class QVariant &)
  // addItem(const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant
    // invoke: void addItem(const class QIcon &, const class QString &, const class QVariant &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN9QComboBox7addItemERK7QStringRK8QVariant
    // invoke: void addItem(const class QString &, const class QVariant &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN9QComboBox7addItemERK7QStringRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "addItem", args)
  }

}

  // proto:  void QComboBox::insertItems(int index, const QStringList & texts);
func (this *QComboBox) insertItems(args ...interface{}) () {
  // insertItems(int, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11insertItemsEiRK11QStringList
    // invoke: void insertItems(int, const class QStringList &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QComboBox11insertItemsEiRK11QStringList(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "insertItems", args)
  }

}

  // proto:  QSize QComboBox::iconSize();
func (this *QComboBox) iconSize(args ...interface{}) () {
  // iconSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8iconSizeEv
    // invoke: QSize iconSize()
    C._ZNK9QComboBox8iconSizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "iconSize", args)
  }

}

  // proto:  QModelIndex QComboBox::rootModelIndex();
func (this *QComboBox) rootModelIndex(args ...interface{}) () {
  // rootModelIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox14rootModelIndexEv
    // invoke: QModelIndex rootModelIndex()
    C._ZNK9QComboBox14rootModelIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "rootModelIndex", args)
  }

}

  // proto:  void QComboBox::setEditable(bool editable);
func (this *QComboBox) setEditable(args ...interface{}) () {
  // setEditable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setEditableEb
    // invoke: void setEditable(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox11setEditableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setEditable", args)
  }

}

  // proto:  void QComboBox::setItemIcon(int index, const QIcon & icon);
func (this *QComboBox) setItemIcon(args ...interface{}) () {
  // setItemIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setItemIconEiRK5QIcon
    // invoke: void setItemIcon(int, const class QIcon &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QComboBox11setItemIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemIcon", args)
  }

}

  // proto:  bool QComboBox::autoCompletion();
func (this *QComboBox) autoCompletion(args ...interface{}) () {
  // autoCompletion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox14autoCompletionEv
    // invoke: bool autoCompletion()
    C._ZNK9QComboBox14autoCompletionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "autoCompletion", args)
  }

}

  // proto:  QVariant QComboBox::currentData(int role);
func (this *QComboBox) currentData(args ...interface{}) () {
  // currentData(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox11currentDataEi
    // invoke: QVariant currentData(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QComboBox11currentDataEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "currentData", args)
  }

}

  // proto:  bool QComboBox::hasFrame();
func (this *QComboBox) hasFrame(args ...interface{}) () {
  // hasFrame()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8hasFrameEv
    // invoke: bool hasFrame()
    C._ZNK9QComboBox8hasFrameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "hasFrame", args)
  }

}

  // proto:  const QValidator * QComboBox::validator();
func (this *QComboBox) validator(args ...interface{}) () {
  // validator()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox9validatorEv
    // invoke: const QValidator * validator()
    C._ZNK9QComboBox9validatorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "validator", args)
  }

}

  // proto:  QString QComboBox::itemText(int index);
func (this *QComboBox) itemText(args ...interface{}) () {
  // itemText(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8itemTextEi
    // invoke: QString itemText(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QComboBox8itemTextEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "itemText", args)
  }

}

  // proto:  void QComboBox::setItemData(int index, const QVariant & value, int role);
func (this *QComboBox) setItemData(args ...interface{}) () {
  // setItemData(int, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setItemDataEiRK8QVarianti
    // invoke: void setItemData(int, const class QVariant &, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN9QComboBox11setItemDataEiRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemData", args)
  }

}

  // proto:  void QComboBox::hidePopup();
func (this *QComboBox) hidePopup(args ...interface{}) () {
  // hidePopup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox9hidePopupEv
    // invoke: void hidePopup()
    C._ZN9QComboBox9hidePopupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "hidePopup", args)
  }

}

  // proto:  void QComboBox::insertItem(int index, const QIcon & icon, const QString & text, const QVariant & userData);
func (this *QComboBox) insertItem(args ...interface{}) () {
  // insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
  // insertItem(int, const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant
    // invoke: void insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QVariant).qclsinst
    if false {fmt.Println(arg3)}
    C._ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant(this.qclsinst, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN9QComboBox10insertItemEiRK7QStringRK8QVariant
    // invoke: void insertItem(int, const class QString &, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QComboBox", "insertItem", args)
  }

}

  // proto:  void QComboBox::setCurrentText(const QString & text);
func (this *QComboBox) setCurrentText(args ...interface{}) () {
  // setCurrentText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox14setCurrentTextERK7QString
    // invoke: void setCurrentText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox14setCurrentTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setCurrentText", args)
  }

}

  // proto:  int QComboBox::modelColumn();
func (this *QComboBox) modelColumn(args ...interface{}) () {
  // modelColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox11modelColumnEv
    // invoke: int modelColumn()
    C._ZNK9QComboBox11modelColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "modelColumn", args)
  }

}

  // proto:  QSize QComboBox::sizeHint();
func (this *QComboBox) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8sizeHintEv
    // invoke: QSize sizeHint()
    C._ZNK9QComboBox8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "sizeHint", args)
  }

}

  // proto:  QVariant QComboBox::itemData(int index, int role);
func (this *QComboBox) itemData(args ...interface{}) () {
  // itemData(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8itemDataEii
    // invoke: QVariant itemData(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK9QComboBox8itemDataEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "itemData", args)
  }

}

  // proto:  void QComboBox::setCompleter(QCompleter * c);
func (this *QComboBox) setCompleter(args ...interface{}) () {
  // setCompleter(class QCompleter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QCompleter{}) // "QCompleter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox12setCompleterEP10QCompleter
    // invoke: void setCompleter(class QCompleter *)
    var arg0 = args[0].(QCompleter).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox12setCompleterEP10QCompleter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setCompleter", args)
  }

}

  // proto:  int QComboBox::maxVisibleItems();
func (this *QComboBox) maxVisibleItems(args ...interface{}) () {
  // maxVisibleItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox15maxVisibleItemsEv
    // invoke: int maxVisibleItems()
    C._ZNK9QComboBox15maxVisibleItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "maxVisibleItems", args)
  }

}

  // proto:  void QComboBox::QComboBox(QWidget * parent);
func NewQComboBox(args ...interface{}) QComboBox {
  return QComboBox{}
}

  // proto:  void QComboBox::setCurrentIndex(int index);
func (this *QComboBox) setCurrentIndex(args ...interface{}) () {
  // setCurrentIndex(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox15setCurrentIndexEi
    // invoke: void setCurrentIndex(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setCurrentIndex", args)
  }

}

  // proto:  void QComboBox::setRootModelIndex(const QModelIndex & index);
func (this *QComboBox) setRootModelIndex(args ...interface{}) () {
  // setRootModelIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox17setRootModelIndexERK11QModelIndex
    // invoke: void setRootModelIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox17setRootModelIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setRootModelIndex", args)
  }

}

  // proto:  void QComboBox::setEditText(const QString & text);
func (this *QComboBox) setEditText(args ...interface{}) () {
  // setEditText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setEditTextERK7QString
    // invoke: void setEditText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox11setEditTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setEditText", args)
  }

}

  // proto:  QCompleter * QComboBox::completer();
func (this *QComboBox) completer(args ...interface{}) () {
  // completer()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox9completerEv
    // invoke: QCompleter * completer()
    C._ZNK9QComboBox9completerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "completer", args)
  }

}

  // proto:  void QComboBox::removeItem(int index);
func (this *QComboBox) removeItem(args ...interface{}) () {
  // removeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox10removeItemEi
    // invoke: void removeItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox10removeItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "removeItem", args)
  }

}

  // proto:  int QComboBox::count();
func (this *QComboBox) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox5countEv
    // invoke: int count()
    C._ZNK9QComboBox5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "count", args)
  }

}

  // proto:  void QComboBox::setItemDelegate(QAbstractItemDelegate * delegate);
func (this *QComboBox) setItemDelegate(args ...interface{}) () {
  // setItemDelegate(class QAbstractItemDelegate *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAbstractItemDelegate{}) // "QAbstractItemDelegate *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate
    // invoke: void setItemDelegate(class QAbstractItemDelegate *)
    var arg0 = args[0].(QAbstractItemDelegate).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemDelegate", args)
  }

}

  // proto:  void QComboBox::addItems(const QStringList & texts);
func (this *QComboBox) addItems(args ...interface{}) () {
  // addItems(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox8addItemsERK11QStringList
    // invoke: void addItems(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN9QComboBox8addItemsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "addItems", args)
  }

}

  // proto:  void QComboBox::setMinimumContentsLength(int characters);
func (this *QComboBox) setMinimumContentsLength(args ...interface{}) () {
  // setMinimumContentsLength(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox24setMinimumContentsLengthEi
    // invoke: void setMinimumContentsLength(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox24setMinimumContentsLengthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setMinimumContentsLength", args)
  }

}

  // proto:  bool QComboBox::duplicatesEnabled();
func (this *QComboBox) duplicatesEnabled(args ...interface{}) () {
  // duplicatesEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox17duplicatesEnabledEv
    // invoke: bool duplicatesEnabled()
    C._ZNK9QComboBox17duplicatesEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "duplicatesEnabled", args)
  }

}

  // proto:  void QComboBox::~QComboBox();
func (this *QComboBox) FreeQComboBox(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QComboBox", "~QComboBox", args)
  }

}

  // proto:  QAbstractItemModel * QComboBox::model();
func (this *QComboBox) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox5modelEv
    // invoke: QAbstractItemModel * model()
    C._ZNK9QComboBox5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "model", args)
  }

}

  // proto:  int QComboBox::minimumContentsLength();
func (this *QComboBox) minimumContentsLength(args ...interface{}) () {
  // minimumContentsLength()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox21minimumContentsLengthEv
    // invoke: int minimumContentsLength()
    C._ZNK9QComboBox21minimumContentsLengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "minimumContentsLength", args)
  }

}

  // proto:  bool QComboBox::isEditable();
func (this *QComboBox) isEditable(args ...interface{}) () {
  // isEditable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox10isEditableEv
    // invoke: bool isEditable()
    C._ZNK9QComboBox10isEditableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "isEditable", args)
  }

}

  // proto:  void QComboBox::setMaxCount(int max);
func (this *QComboBox) setMaxCount(args ...interface{}) () {
  // setMaxCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setMaxCountEi
    // invoke: void setMaxCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox11setMaxCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setMaxCount", args)
  }

}

  // proto:  int QComboBox::currentIndex();
func (this *QComboBox) currentIndex(args ...interface{}) () {
  // currentIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox12currentIndexEv
    // invoke: int currentIndex()
    C._ZNK9QComboBox12currentIndexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "currentIndex", args)
  }

}

  // proto:  void QComboBox::setDuplicatesEnabled(bool enable);
func (this *QComboBox) setDuplicatesEnabled(args ...interface{}) () {
  // setDuplicatesEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox20setDuplicatesEnabledEb
    // invoke: void setDuplicatesEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox20setDuplicatesEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setDuplicatesEnabled", args)
  }

}

  // proto:  QString QComboBox::currentText();
func (this *QComboBox) currentText(args ...interface{}) () {
  // currentText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox11currentTextEv
    // invoke: QString currentText()
    C._ZNK9QComboBox11currentTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "currentText", args)
  }

}

  // proto:  void QComboBox::showPopup();
func (this *QComboBox) showPopup(args ...interface{}) () {
  // showPopup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox9showPopupEv
    // invoke: void showPopup()
    C._ZN9QComboBox9showPopupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "showPopup", args)
  }

}

  // proto:  QLineEdit * QComboBox::lineEdit();
func (this *QComboBox) lineEdit(args ...interface{}) () {
  // lineEdit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8lineEditEv
    // invoke: QLineEdit * lineEdit()
    C._ZNK9QComboBox8lineEditEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "lineEdit", args)
  }

}

  // proto:  QAbstractItemDelegate * QComboBox::itemDelegate();
func (this *QComboBox) itemDelegate(args ...interface{}) () {
  // itemDelegate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox12itemDelegateEv
    // invoke: QAbstractItemDelegate * itemDelegate()
    C._ZNK9QComboBox12itemDelegateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "itemDelegate", args)
  }

}

  // proto:  void QComboBox::setMaxVisibleItems(int maxItems);
func (this *QComboBox) setMaxVisibleItems(args ...interface{}) () {
  // setMaxVisibleItems(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox18setMaxVisibleItemsEi
    // invoke: void setMaxVisibleItems(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox18setMaxVisibleItemsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setMaxVisibleItems", args)
  }

}

  // proto:  bool QComboBox::event(QEvent * event);
func (this *QComboBox) event(args ...interface{}) () {
  // event(class QEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QEvent{}) // "QEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox5eventEP6QEvent
    // invoke: bool event(class QEvent *)
    var arg0 = args[0].(QEvent).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox5eventEP6QEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "event", args)
  }

}

  // proto:  void QComboBox::setModelColumn(int visibleColumn);
func (this *QComboBox) setModelColumn(args ...interface{}) () {
  // setModelColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox14setModelColumnEi
    // invoke: void setModelColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox14setModelColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setModelColumn", args)
  }

}

  // proto:  void QComboBox::setItemText(int index, const QString & text);
func (this *QComboBox) setItemText(args ...interface{}) () {
  // setItemText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setItemTextEiRK7QString
    // invoke: void setItemText(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN9QComboBox11setItemTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemText", args)
  }

}

  // proto:  void QComboBox::setLineEdit(QLineEdit * edit);
func (this *QComboBox) setLineEdit(args ...interface{}) () {
  // setLineEdit(class QLineEdit *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLineEdit{}) // "QLineEdit *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox11setLineEditEP9QLineEdit
    // invoke: void setLineEdit(class QLineEdit *)
    var arg0 = args[0].(QLineEdit).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox11setLineEditEP9QLineEdit(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setLineEdit", args)
  }

}

  // proto:  QIcon QComboBox::itemIcon(int index);
func (this *QComboBox) itemIcon(args ...interface{}) () {
  // itemIcon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox8itemIconEi
    // invoke: QIcon itemIcon(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK9QComboBox8itemIconEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "itemIcon", args)
  }

}

  // proto:  void QComboBox::setValidator(const QValidator * v);
func (this *QComboBox) setValidator(args ...interface{}) () {
  // setValidator(const class QValidator *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QValidator{}) // "const QValidator *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox12setValidatorEPK10QValidator
    // invoke: void setValidator(const class QValidator *)
    var arg0 = args[0].(QValidator).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox12setValidatorEPK10QValidator(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setValidator", args)
  }

}

  // proto:  void QComboBox::insertSeparator(int index);
func (this *QComboBox) insertSeparator(args ...interface{}) () {
  // insertSeparator(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox15insertSeparatorEi
    // invoke: void insertSeparator(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN9QComboBox15insertSeparatorEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "insertSeparator", args)
  }

}

  // proto:  const QMetaObject * QComboBox::metaObject();
func (this *QComboBox) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK9QComboBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "metaObject", args)
  }

}

// <= body block end

