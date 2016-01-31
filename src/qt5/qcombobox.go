package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  bool QComboBox::hasFrame();
extern bool C_ZNK9QComboBox8hasFrameEv(void* qthis); // 4
  // proto:  bool QComboBox::duplicatesEnabled();
extern bool C_ZNK9QComboBox17duplicatesEnabledEv(void* qthis); // 4
  // proto:  void QComboBox::setView(QAbstractItemView * itemView);
extern void C_ZN9QComboBox7setViewEP17QAbstractItemView(void* qthis, void* arg0); // 4
  // proto:  void QComboBox::QComboBox(QWidget * parent);
extern void* C_ZN9QComboBoxC2EP7QWidget(void* arg0); // 3
  // proto:  Qt::CaseSensitivity QComboBox::autoCompletionCaseSensitivity();
extern void C_ZNK9QComboBox29autoCompletionCaseSensitivityEv(void* qthis); // 4
  // proto:  int QComboBox::count();
extern int32_t C_ZNK9QComboBox5countEv(void* qthis); // 4
  // proto:  QIcon QComboBox::itemIcon(int index);
extern void* C_ZNK9QComboBox8itemIconEi(void* qthis, int32_t arg0); // 4
  // proto:  void QComboBox::insertItem(int index, const QString & text, const QVariant & userData);
extern void C_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant(void* qthis, int32_t arg0, void* arg1, void* arg2); // 2
  // proto:  void QComboBox::insertItem(int index, const QIcon & icon, const QString & text, const QVariant & userData);
extern void C_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant(void* qthis, int32_t arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  void QComboBox::removeItem(int index);
extern void C_ZN9QComboBox10removeItemEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QComboBox::metaObject();
extern void C_ZNK9QComboBox10metaObjectEv(void* qthis); // 4
  // proto:  void QComboBox::setItemData(int index, const QVariant & value, int role);
extern void C_ZN9QComboBox11setItemDataEiRK8QVarianti(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  void QComboBox::setIconSize(const QSize & size);
extern void C_ZN9QComboBox11setIconSizeERK5QSize(void* qthis, void* arg0); // 4
  // proto:  int QComboBox::currentIndex();
extern int32_t C_ZNK9QComboBox12currentIndexEv(void* qthis); // 4
  // proto:  QAbstractItemView * QComboBox::view();
extern void C_ZNK9QComboBox4viewEv(void* qthis); // 4
  // proto:  bool QComboBox::isEditable();
extern bool C_ZNK9QComboBox10isEditableEv(void* qthis); // 4
  // proto:  void QComboBox::setEditText(const QString & text);
extern void C_ZN9QComboBox11setEditTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QComboBox::maxVisibleItems();
extern int32_t C_ZNK9QComboBox15maxVisibleItemsEv(void* qthis); // 4
  // proto:  void QComboBox::setCurrentText(const QString & text);
extern void C_ZN9QComboBox14setCurrentTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QComboBox::event(QEvent * event);
extern bool C_ZN9QComboBox5eventEP6QEvent(void* qthis, void* arg0); // 4
  // proto:  void QComboBox::setDuplicatesEnabled(bool enable);
extern void C_ZN9QComboBox20setDuplicatesEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QComboBox::setEditable(bool editable);
extern void C_ZN9QComboBox11setEditableEb(void* qthis, bool arg0); // 4
  // proto:  QVariant QComboBox::itemData(int index, int role);
extern void* C_ZNK9QComboBox8itemDataEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QSize QComboBox::iconSize();
extern void* C_ZNK9QComboBox8iconSizeEv(void* qthis); // 4
  // proto:  void QComboBox::addItem(const QIcon & icon, const QString & text, const QVariant & userData);
extern void C_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  void QComboBox::addItem(const QString & text, const QVariant & userData);
extern void C_ZN9QComboBox7addItemERK7QStringRK8QVariant(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QSize QComboBox::sizeHint();
extern void* C_ZNK9QComboBox8sizeHintEv(void* qthis); // 4
  // proto:  void QComboBox::setMaxVisibleItems(int maxItems);
extern void C_ZN9QComboBox18setMaxVisibleItemsEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QComboBox::currentText();
extern void* C_ZNK9QComboBox11currentTextEv(void* qthis); // 4
  // proto:  void QComboBox::setMaxCount(int max);
extern void C_ZN9QComboBox11setMaxCountEi(void* qthis, int32_t arg0); // 4
  // proto:  QModelIndex QComboBox::rootModelIndex();
extern void* C_ZNK9QComboBox14rootModelIndexEv(void* qthis); // 4
  // proto:  int QComboBox::modelColumn();
extern int32_t C_ZNK9QComboBox11modelColumnEv(void* qthis); // 4
  // proto:  void QComboBox::setItemText(int index, const QString & text);
extern void C_ZN9QComboBox11setItemTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QComboBox::hidePopup();
extern void C_ZN9QComboBox9hidePopupEv(void* qthis); // 4
  // proto:  void QComboBox::showPopup();
extern void C_ZN9QComboBox9showPopupEv(void* qthis); // 4
  // proto:  int QComboBox::minimumContentsLength();
extern int32_t C_ZNK9QComboBox21minimumContentsLengthEv(void* qthis); // 4
  // proto:  void QComboBox::addItems(const QStringList & texts);
extern void C_ZN9QComboBox8addItemsERK11QStringList(void* qthis, void* arg0); // 2
  // proto:  void QComboBox::setCompleter(QCompleter * c);
extern void C_ZN9QComboBox12setCompleterEP10QCompleter(void* qthis, void* arg0); // 4
  // proto:  void QComboBox::insertItems(int index, const QStringList & texts);
extern void C_ZN9QComboBox11insertItemsEiRK11QStringList(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QComboBox::setRootModelIndex(const QModelIndex & index);
extern void C_ZN9QComboBox17setRootModelIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QComboBox::setValidator(const QValidator * v);
extern void C_ZN9QComboBox12setValidatorEPK10QValidator(void* qthis, void* arg0); // 4
  // proto:  int QComboBox::maxCount();
extern int32_t C_ZNK9QComboBox8maxCountEv(void* qthis); // 4
  // proto:  QCompleter * QComboBox::completer();
extern void* C_ZNK9QComboBox9completerEv(void* qthis); // 4
  // proto:  QString QComboBox::itemText(int index);
extern void* C_ZNK9QComboBox8itemTextEi(void* qthis, int32_t arg0); // 4
  // proto:  QSize QComboBox::minimumSizeHint();
extern void* C_ZNK9QComboBox15minimumSizeHintEv(void* qthis); // 4
  // proto:  void QComboBox::insertSeparator(int index);
extern void C_ZN9QComboBox15insertSeparatorEi(void* qthis, int32_t arg0); // 4
  // proto:  QComboBox::SizeAdjustPolicy QComboBox::sizeAdjustPolicy();
extern void C_ZNK9QComboBox16sizeAdjustPolicyEv(void* qthis); // 4
  // proto:  QLineEdit * QComboBox::lineEdit();
extern void* C_ZNK9QComboBox8lineEditEv(void* qthis); // 4
  // proto:  QVariant QComboBox::currentData(int role);
extern void* C_ZNK9QComboBox11currentDataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QComboBox::setItemDelegate(QAbstractItemDelegate * delegate);
extern void C_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate(void* qthis, void* arg0); // 4
  // proto:  void QComboBox::~QComboBox();
extern void C_ZN9QComboBoxD2Ev(void* qthis); // 4
  // proto:  void QComboBox::setCurrentIndex(int index);
extern void C_ZN9QComboBox15setCurrentIndexEi(void* qthis, int32_t arg0); // 4
  // proto:  void QComboBox::setLineEdit(QLineEdit * edit);
extern void C_ZN9QComboBox11setLineEditEP9QLineEdit(void* qthis, void* arg0); // 4
  // proto:  void QComboBox::setModelColumn(int visibleColumn);
extern void C_ZN9QComboBox14setModelColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QComboBox::autoCompletion();
extern bool C_ZNK9QComboBox14autoCompletionEv(void* qthis); // 4
  // proto:  void QComboBox::setAutoCompletion(bool enable);
extern void C_ZN9QComboBox17setAutoCompletionEb(void* qthis, bool arg0); // 4
  // proto:  void QComboBox::setMinimumContentsLength(int characters);
extern void C_ZN9QComboBox24setMinimumContentsLengthEi(void* qthis, int32_t arg0); // 4
  // proto:  void QComboBox::clearEditText();
extern void C_ZN9QComboBox13clearEditTextEv(void* qthis); // 4
  // proto:  void QComboBox::setModel(QAbstractItemModel * model);
extern void C_ZN9QComboBox8setModelEP18QAbstractItemModel(void* qthis, void* arg0); // 4
  // proto:  QComboBox::InsertPolicy QComboBox::insertPolicy();
extern void C_ZNK9QComboBox12insertPolicyEv(void* qthis); // 4
  // proto:  void QComboBox::setItemIcon(int index, const QIcon & icon);
extern void C_ZN9QComboBox11setItemIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QAbstractItemDelegate * QComboBox::itemDelegate();
extern void C_ZNK9QComboBox12itemDelegateEv(void* qthis); // 4
  // proto:  void QComboBox::setFrame(bool );
extern void C_ZN9QComboBox8setFrameEb(void* qthis, bool arg0); // 4
  // proto:  void QComboBox::clear();
extern void C_ZN9QComboBox5clearEv(void* qthis); // 4
  // proto:  const QValidator * QComboBox::validator();
extern void* C_ZNK9QComboBox9validatorEv(void* qthis); // 4
  // proto:  QAbstractItemModel * QComboBox::model();
extern void C_ZNK9QComboBox5modelEv(void* qthis); // 4
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

// hasFrame()
func (this *QComboBox) Hasframe(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8hasFrameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "hasFrame", args)
  }

  return
}

// duplicatesEnabled()
func (this *QComboBox) Duplicatesenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox17duplicatesEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "duplicatesEnabled", args)
  }

  return
}

// setView(class QAbstractItemView *)
func (this *QComboBox) Setview(args ...interface{}) () {
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
    C.C_ZN9QComboBox7setViewEP17QAbstractItemView(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setView", args)
  }

  return
}

// QComboBox(class QWidget *)
func NewQComboBox(args ...interface{}) *QComboBox {
  // QComboBox(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBoxC1EP7QWidget
    // invoke: void QComboBox(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QComboBoxC2EP7QWidget(arg0)
    return &QComboBox{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QComboBox", "QComboBox", args)
  }

  return nil // QComboBox{qclsinst:qthis}
}

// autoCompletionCaseSensitivity()
func (this *QComboBox) Autocompletioncasesensitivity(args ...interface{}) () {
  // autoCompletionCaseSensitivity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox29autoCompletionCaseSensitivityEv
    // invoke: Qt::CaseSensitivity autoCompletionCaseSensitivity()
    C.C_ZNK9QComboBox29autoCompletionCaseSensitivityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "autoCompletionCaseSensitivity", args)
  }

  return
}

// count()
func (this *QComboBox) Count(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "count", args)
  }

  return
}

// itemIcon(int)
func (this *QComboBox) Itemicon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8itemIconEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "itemIcon", args)
  }

  return
}

// insertItem(int, const class QString &, const class QVariant &)
func (this *QComboBox) Insertitem(args ...interface{}) () {
  // insertItem(int, const class QString &, const class QVariant &)
  // insertItem(int, const class QIcon &, const class QString &, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBox10insertItemEiRK7QStringRK8QVariant
    // invoke: void insertItem(int, const class QString &, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN9QComboBox10insertItemEiRK7QStringRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  case 1:
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
    C.C_ZN9QComboBox10insertItemEiRK5QIconRK7QStringRK8QVariant(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QComboBox", "insertItem", args)
  }

  return
}

// removeItem(int)
func (this *QComboBox) Removeitem(args ...interface{}) () {
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
    C.C_ZN9QComboBox10removeItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "removeItem", args)
  }

  return
}

// metaObject()
func (this *QComboBox) Metaobject(args ...interface{}) () {
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
    C.C_ZNK9QComboBox10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "metaObject", args)
  }

  return
}

// setItemData(int, const class QVariant &, int)
func (this *QComboBox) Setitemdata(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setItemDataEiRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemData", args)
  }

  return
}

// setIconSize(const class QSize &)
func (this *QComboBox) Seticonsize(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setIconSizeERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setIconSize", args)
  }

  return
}

// currentIndex()
func (this *QComboBox) Currentindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox12currentIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "currentIndex", args)
  }

  return
}

// view()
func (this *QComboBox) View(args ...interface{}) () {
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
    C.C_ZNK9QComboBox4viewEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "view", args)
  }

  return
}

// isEditable()
func (this *QComboBox) Iseditable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox10isEditableEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "isEditable", args)
  }

  return
}

// setEditText(const class QString &)
func (this *QComboBox) Setedittext(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setEditTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setEditText", args)
  }

  return
}

// maxVisibleItems()
func (this *QComboBox) Maxvisibleitems(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox15maxVisibleItemsEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "maxVisibleItems", args)
  }

  return
}

// setCurrentText(const class QString &)
func (this *QComboBox) Setcurrenttext(args ...interface{}) () {
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
    C.C_ZN9QComboBox14setCurrentTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setCurrentText", args)
  }

  return
}

// event(class QEvent *)
func (this *QComboBox) Event(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN9QComboBox5eventEP6QEvent(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "event", args)
  }

  return
}

// setDuplicatesEnabled(_Bool)
func (this *QComboBox) Setduplicatesenabled(args ...interface{}) () {
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
    C.C_ZN9QComboBox20setDuplicatesEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setDuplicatesEnabled", args)
  }

  return
}

// setEditable(_Bool)
func (this *QComboBox) Seteditable(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setEditableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setEditable", args)
  }

  return
}

// itemData(int, int)
func (this *QComboBox) Itemdata(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8itemDataEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "itemData", args)
  }

  return
}

// iconSize()
func (this *QComboBox) Iconsize(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8iconSizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "iconSize", args)
  }

  return
}

// addItem(const class QIcon &, const class QString &, const class QVariant &)
func (this *QComboBox) Additem(args ...interface{}) () {
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
    C.C_ZN9QComboBox7addItemERK5QIconRK7QStringRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN9QComboBox7addItemERK7QStringRK8QVariant
    // invoke: void addItem(const class QString &, const class QVariant &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QComboBox7addItemERK7QStringRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "addItem", args)
  }

  return
}

// sizeHint()
func (this *QComboBox) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "sizeHint", args)
  }

  return
}

// setMaxVisibleItems(int)
func (this *QComboBox) Setmaxvisibleitems(args ...interface{}) () {
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
    C.C_ZN9QComboBox18setMaxVisibleItemsEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setMaxVisibleItems", args)
  }

  return
}

// currentText()
func (this *QComboBox) Currenttext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox11currentTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "currentText", args)
  }

  return
}

// setMaxCount(int)
func (this *QComboBox) Setmaxcount(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setMaxCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setMaxCount", args)
  }

  return
}

// rootModelIndex()
func (this *QComboBox) Rootmodelindex(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox14rootModelIndexEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QModelIndex{}) // "QModelIndex"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "rootModelIndex", args)
  }

  return
}

// modelColumn()
func (this *QComboBox) Modelcolumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox11modelColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "modelColumn", args)
  }

  return
}

// setItemText(int, const class QString &)
func (this *QComboBox) Setitemtext(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setItemTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemText", args)
  }

  return
}

// hidePopup()
func (this *QComboBox) Hidepopup(args ...interface{}) () {
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
    C.C_ZN9QComboBox9hidePopupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "hidePopup", args)
  }

  return
}

// showPopup()
func (this *QComboBox) Showpopup(args ...interface{}) () {
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
    C.C_ZN9QComboBox9showPopupEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "showPopup", args)
  }

  return
}

// minimumContentsLength()
func (this *QComboBox) Minimumcontentslength(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox21minimumContentsLengthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "minimumContentsLength", args)
  }

  return
}

// addItems(const class QStringList &)
func (this *QComboBox) Additems(args ...interface{}) () {
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
    C.C_ZN9QComboBox8addItemsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "addItems", args)
  }

  return
}

// setCompleter(class QCompleter *)
func (this *QComboBox) Setcompleter(args ...interface{}) () {
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
    C.C_ZN9QComboBox12setCompleterEP10QCompleter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setCompleter", args)
  }

  return
}

// insertItems(int, const class QStringList &)
func (this *QComboBox) Insertitems(args ...interface{}) () {
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
    C.C_ZN9QComboBox11insertItemsEiRK11QStringList(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "insertItems", args)
  }

  return
}

// setRootModelIndex(const class QModelIndex &)
func (this *QComboBox) Setrootmodelindex(args ...interface{}) () {
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
    C.C_ZN9QComboBox17setRootModelIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setRootModelIndex", args)
  }

  return
}

// setValidator(const class QValidator *)
func (this *QComboBox) Setvalidator(args ...interface{}) () {
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
    C.C_ZN9QComboBox12setValidatorEPK10QValidator(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setValidator", args)
  }

  return
}

// maxCount()
func (this *QComboBox) Maxcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8maxCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "maxCount", args)
  }

  return
}

// completer()
func (this *QComboBox) Completer(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox9completerEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QCompleter{}) // "QCompleter *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "completer", args)
  }

  return
}

// itemText(int)
func (this *QComboBox) Itemtext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8itemTextEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "itemText", args)
  }

  return
}

// minimumSizeHint()
func (this *QComboBox) Minimumsizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox15minimumSizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "minimumSizeHint", args)
  }

  return
}

// insertSeparator(int)
func (this *QComboBox) Insertseparator(args ...interface{}) () {
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
    C.C_ZN9QComboBox15insertSeparatorEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "insertSeparator", args)
  }

  return
}

// sizeAdjustPolicy()
func (this *QComboBox) Sizeadjustpolicy(args ...interface{}) () {
  // sizeAdjustPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox16sizeAdjustPolicyEv
    // invoke: QComboBox::SizeAdjustPolicy sizeAdjustPolicy()
    C.C_ZNK9QComboBox16sizeAdjustPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "sizeAdjustPolicy", args)
  }

  return
}

// lineEdit()
func (this *QComboBox) Lineedit(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox8lineEditEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLineEdit{}) // "QLineEdit *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "lineEdit", args)
  }

  return
}

// currentData(int)
func (this *QComboBox) Currentdata(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox11currentDataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "currentData", args)
  }

  return
}

// setItemDelegate(class QAbstractItemDelegate *)
func (this *QComboBox) Setitemdelegate(args ...interface{}) () {
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
    C.C_ZN9QComboBox15setItemDelegateEP21QAbstractItemDelegate(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemDelegate", args)
  }

  return
}

// ~QComboBox()
func (this *QComboBox) Freeqcombobox(args ...interface{}) () {
  // ~QComboBox()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QComboBoxD0Ev
    // invoke: void ~QComboBox()
    C.C_ZN9QComboBoxD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "~QComboBox", args)
  }

  return
}

// setCurrentIndex(int)
func (this *QComboBox) Setcurrentindex(args ...interface{}) () {
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
    C.C_ZN9QComboBox15setCurrentIndexEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setCurrentIndex", args)
  }

  return
}

// setLineEdit(class QLineEdit *)
func (this *QComboBox) Setlineedit(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setLineEditEP9QLineEdit(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setLineEdit", args)
  }

  return
}

// setModelColumn(int)
func (this *QComboBox) Setmodelcolumn(args ...interface{}) () {
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
    C.C_ZN9QComboBox14setModelColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setModelColumn", args)
  }

  return
}

// autoCompletion()
func (this *QComboBox) Autocompletion(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox14autoCompletionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "autoCompletion", args)
  }

  return
}

// setAutoCompletion(_Bool)
func (this *QComboBox) Setautocompletion(args ...interface{}) () {
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
    C.C_ZN9QComboBox17setAutoCompletionEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setAutoCompletion", args)
  }

  return
}

// setMinimumContentsLength(int)
func (this *QComboBox) Setminimumcontentslength(args ...interface{}) () {
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
    C.C_ZN9QComboBox24setMinimumContentsLengthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setMinimumContentsLength", args)
  }

  return
}

// clearEditText()
func (this *QComboBox) Clearedittext(args ...interface{}) () {
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
    C.C_ZN9QComboBox13clearEditTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "clearEditText", args)
  }

  return
}

// setModel(class QAbstractItemModel *)
func (this *QComboBox) Setmodel(args ...interface{}) () {
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
    C.C_ZN9QComboBox8setModelEP18QAbstractItemModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setModel", args)
  }

  return
}

// insertPolicy()
func (this *QComboBox) Insertpolicy(args ...interface{}) () {
  // insertPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QComboBox12insertPolicyEv
    // invoke: QComboBox::InsertPolicy insertPolicy()
    C.C_ZNK9QComboBox12insertPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "insertPolicy", args)
  }

  return
}

// setItemIcon(int, const class QIcon &)
func (this *QComboBox) Setitemicon(args ...interface{}) () {
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
    C.C_ZN9QComboBox11setItemIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QComboBox", "setItemIcon", args)
  }

  return
}

// itemDelegate()
func (this *QComboBox) Itemdelegate(args ...interface{}) () {
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
    C.C_ZNK9QComboBox12itemDelegateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "itemDelegate", args)
  }

  return
}

// setFrame(_Bool)
func (this *QComboBox) Setframe(args ...interface{}) () {
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
    C.C_ZN9QComboBox8setFrameEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QComboBox", "setFrame", args)
  }

  return
}

// clear()
func (this *QComboBox) Clear(args ...interface{}) () {
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
    C.C_ZN9QComboBox5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "clear", args)
  }

  return
}

// validator()
func (this *QComboBox) Validator(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QComboBox9validatorEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QValidator{}) // "const QValidator *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QComboBox", "validator", args)
  }

  return
}

// model()
func (this *QComboBox) Model(args ...interface{}) () {
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
    C.C_ZNK9QComboBox5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QComboBox", "model", args)
  }

  return
}

// <= body block end

