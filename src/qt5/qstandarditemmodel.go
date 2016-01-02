package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
// src-file: /QtGui/qstandarditemmodel.h
// dst-file: /src/gui/qstandarditemmodel.go
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
  // proto:  void QStandardItemModel::QStandardItemModel(int rows, int columns, QObject * parent);
extern void* dector_ZN18QStandardItemModelC1EiiP7QObject(int arg0, int arg1, void* arg2);
extern void _ZN18QStandardItemModelC1EiiP7QObject(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QStandardItemModel::clear();
extern void _ZN18QStandardItemModel5clearEv(void* qthis);
  // proto:  QStandardItem * QStandardItemModel::item(int row, int column);
extern void _ZNK18QStandardItemModel4itemEii(void* qthis, int arg0, int arg1);
  // proto:  bool QStandardItemModel::insertRow(int row, const QModelIndex & parent);
extern void demth_ZN18QStandardItemModel9insertRowEiRK11QModelIndex(void* qthis, int arg0, void* arg1);
  // proto:  void QStandardItemModel::setItem(int row, QStandardItem * item);
extern void demth_ZN18QStandardItemModel7setItemEiP13QStandardItem(void* qthis, int arg0, void* arg1);
  // proto:  QModelIndex QStandardItemModel::index(int row, int column, const QModelIndex & parent);
extern void _ZNK18QStandardItemModel5indexEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  bool QStandardItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern void _ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  int QStandardItemModel::columnCount(const QModelIndex & parent);
extern void _ZNK18QStandardItemModel11columnCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  QStandardItem * QStandardItemModel::takeItem(int row, int column);
extern void _ZN18QStandardItemModel8takeItemEii(void* qthis, int arg0, int arg1);
  // proto:  void QStandardItemModel::setRowCount(int rows);
extern void _ZN18QStandardItemModel11setRowCountEi(void* qthis, int arg0);
  // proto:  QStandardItem * QStandardItemModel::itemFromIndex(const QModelIndex & index);
extern void _ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QStandardItemModel::insertColumn(int column, const QModelIndex & parent);
extern void demth_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex(void* qthis, int arg0, void* arg1);
  // proto:  void QStandardItemModel::setVerticalHeaderItem(int row, QStandardItem * item);
extern void _ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem(void* qthis, int arg0, void* arg1);
  // proto:  void QStandardItemModel::QStandardItemModel(const QStandardItemModel & );
extern void* dector_ZN18QStandardItemModelC1ERKS_(void* arg0);
extern void _ZN18QStandardItemModelC1ERKS_(void* qthis, void* arg0);
  // proto:  void QStandardItemModel::QStandardItemModel(QObject * parent);
extern void* dector_ZN18QStandardItemModelC1EP7QObject(void* arg0);
extern void _ZN18QStandardItemModelC1EP7QObject(void* qthis, void* arg0);
  // proto:  QList<QStandardItem *> QStandardItemModel::takeColumn(int column);
extern void _ZN18QStandardItemModel10takeColumnEi(void* qthis, int arg0);
  // proto:  QStandardItem * QStandardItemModel::takeVerticalHeaderItem(int row);
extern void _ZN18QStandardItemModel22takeVerticalHeaderItemEi(void* qthis, int arg0);
  // proto:  bool QStandardItemModel::insertColumns(int column, int count, const QModelIndex & parent);
extern void _ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  const QMetaObject * QStandardItemModel::metaObject();
extern void _ZNK18QStandardItemModel10metaObjectEv(void* qthis);
  // proto:  bool QStandardItemModel::insertRows(int row, int count, const QModelIndex & parent);
extern void _ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QStandardItemModel::insertRow(int row, QStandardItem * item);
extern void demth_ZN18QStandardItemModel9insertRowEiP13QStandardItem(void* qthis, int arg0, void* arg1);
  // proto:  QStandardItem * QStandardItemModel::invisibleRootItem();
extern void _ZNK18QStandardItemModel17invisibleRootItemEv(void* qthis);
  // proto:  void QStandardItemModel::setItemPrototype(const QStandardItem * item);
extern void _ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem(void* qthis, void* arg0);
  // proto:  void QStandardItemModel::setHorizontalHeaderLabels(const QStringList & labels);
extern void _ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList(void* qthis, void* arg0);
  // proto:  QModelIndex QStandardItemModel::parent(const QModelIndex & child);
extern void _ZNK18QStandardItemModel6parentERK11QModelIndex(void* qthis, void* arg0);
  // proto:  bool QStandardItemModel::removeColumns(int column, int count, const QModelIndex & parent);
extern void _ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QModelIndex QStandardItemModel::sibling(int row, int column, const QModelIndex & idx);
extern void _ZNK18QStandardItemModel7siblingEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  int QStandardItemModel::sortRole();
extern void _ZNK18QStandardItemModel8sortRoleEv(void* qthis);
  // proto:  QStandardItem * QStandardItemModel::takeHorizontalHeaderItem(int column);
extern void _ZN18QStandardItemModel24takeHorizontalHeaderItemEi(void* qthis, int arg0);
  // proto:  QModelIndex QStandardItemModel::indexFromItem(const QStandardItem * item);
extern void _ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem(void* qthis, void* arg0);
  // proto:  const QStandardItem * QStandardItemModel::itemPrototype();
extern void _ZNK18QStandardItemModel13itemPrototypeEv(void* qthis);
  // proto:  void QStandardItemModel::setHorizontalHeaderItem(int column, QStandardItem * item);
extern void _ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem(void* qthis, int arg0, void* arg1);
  // proto:  QStandardItem * QStandardItemModel::horizontalHeaderItem(int column);
extern void _ZNK18QStandardItemModel20horizontalHeaderItemEi(void* qthis, int arg0);
  // proto:  void QStandardItemModel::appendRow(QStandardItem * item);
extern void demth_ZN18QStandardItemModel9appendRowEP13QStandardItem(void* qthis, void* arg0);
  // proto:  QMap<int, QVariant> QStandardItemModel::itemData(const QModelIndex & index);
extern void _ZNK18QStandardItemModel8itemDataERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QStandardItemModel::setSortRole(int role);
extern void _ZN18QStandardItemModel11setSortRoleEi(void* qthis, int arg0);
  // proto:  bool QStandardItemModel::hasChildren(const QModelIndex & parent);
extern void _ZNK18QStandardItemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QStandardItemModel::~QStandardItemModel();
extern void _ZN18QStandardItemModelD0Ev(void* qthis);
  // proto:  QVariant QStandardItemModel::data(const QModelIndex & index, int role);
extern void _ZNK18QStandardItemModel4dataERK11QModelIndexi(void* qthis, void* arg0, int arg1);
  // proto:  QList<QStandardItem *> QStandardItemModel::takeRow(int row);
extern void _ZN18QStandardItemModel7takeRowEi(void* qthis, int arg0);
  // proto:  void QStandardItemModel::setColumnCount(int columns);
extern void _ZN18QStandardItemModel14setColumnCountEi(void* qthis, int arg0);
  // proto:  QStandardItem * QStandardItemModel::verticalHeaderItem(int row);
extern void _ZNK18QStandardItemModel18verticalHeaderItemEi(void* qthis, int arg0);
  // proto:  bool QStandardItemModel::removeRows(int row, int count, const QModelIndex & parent);
extern void _ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QStandardItemModel::setVerticalHeaderLabels(const QStringList & labels);
extern void _ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList(void* qthis, void* arg0);
  // proto:  void QStandardItemModel::setItem(int row, int column, QStandardItem * item);
extern void _ZN18QStandardItemModel7setItemEiiP13QStandardItem(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QStringList QStandardItemModel::mimeTypes();
extern void _ZNK18QStandardItemModel9mimeTypesEv(void* qthis);
  // proto:  int QStandardItemModel::rowCount(const QModelIndex & parent);
extern void _ZNK18QStandardItemModel8rowCountERK11QModelIndex(void* qthis, void* arg0);
  // proto:  void QStandardItem::setChild(int row, QStandardItem * item);
extern void demth_ZN13QStandardItem8setChildEiPS_(void* qthis, int arg0, void* arg1);
  // proto:  QStandardItemModel * QStandardItem::model();
extern void _ZNK13QStandardItem5modelEv(void* qthis);
  // proto:  void QStandardItem::insertColumns(int column, int count);
extern void _ZN13QStandardItem13insertColumnsEii(void* qthis, int arg0, int arg1);
  // proto:  void QStandardItem::setSelectable(bool selectable);
extern void _ZN13QStandardItem13setSelectableEb(void* qthis, bool arg0);
  // proto:  int QStandardItem::column();
extern void _ZNK13QStandardItem6columnEv(void* qthis);
  // proto:  QString QStandardItem::whatsThis();
extern void demth_ZNK13QStandardItem9whatsThisEv(void* qthis);
  // proto:  QList<QStandardItem *> QStandardItem::takeColumn(int column);
extern void _ZN13QStandardItem10takeColumnEi(void* qthis, int arg0);
  // proto:  void QStandardItem::setForeground(const QBrush & brush);
extern void demth_ZN13QStandardItem13setForegroundERK6QBrush(void* qthis, void* arg0);
  // proto:  bool QStandardItem::isEditable();
extern void demth_ZNK13QStandardItem10isEditableEv(void* qthis);
  // proto:  QIcon QStandardItem::icon();
extern void demth_ZNK13QStandardItem4iconEv(void* qthis);
  // proto:  void QStandardItem::setWhatsThis(const QString & whatsThis);
extern void demth_ZN13QStandardItem12setWhatsThisERK7QString(void* qthis, void* arg0);
  // proto:  QStandardItem * QStandardItem::takeChild(int row, int column);
extern void _ZN13QStandardItem9takeChildEii(void* qthis, int arg0, int arg1);
  // proto:  int QStandardItem::type();
extern void _ZNK13QStandardItem4typeEv(void* qthis);
  // proto:  QList<QStandardItem *> QStandardItem::takeRow(int row);
extern void _ZN13QStandardItem7takeRowEi(void* qthis, int arg0);
  // proto:  int QStandardItem::row();
extern void _ZNK13QStandardItem3rowEv(void* qthis);
  // proto:  bool QStandardItem::isCheckable();
extern void demth_ZNK13QStandardItem11isCheckableEv(void* qthis);
  // proto:  QString QStandardItem::text();
extern void demth_ZNK13QStandardItem4textEv(void* qthis);
  // proto:  void QStandardItem::insertRows(int row, int count);
extern void _ZN13QStandardItem10insertRowsEii(void* qthis, int arg0, int arg1);
  // proto:  bool QStandardItem::isDropEnabled();
extern void demth_ZNK13QStandardItem13isDropEnabledEv(void* qthis);
  // proto:  bool QStandardItem::hasChildren();
extern void _ZNK13QStandardItem11hasChildrenEv(void* qthis);
  // proto:  QString QStandardItem::statusTip();
extern void demth_ZNK13QStandardItem9statusTipEv(void* qthis);
  // proto:  void QStandardItem::setStatusTip(const QString & statusTip);
extern void demth_ZN13QStandardItem12setStatusTipERK7QString(void* qthis, void* arg0);
  // proto:  void QStandardItem::appendRow(QStandardItem * item);
extern void demth_ZN13QStandardItem9appendRowEPS_(void* qthis, void* arg0);
  // proto:  void QStandardItem::setChild(int row, int column, QStandardItem * item);
extern void _ZN13QStandardItem8setChildEiiPS_(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  QModelIndex QStandardItem::index();
extern void _ZNK13QStandardItem5indexEv(void* qthis);
  // proto:  void QStandardItem::setIcon(const QIcon & icon);
extern void demth_ZN13QStandardItem7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  void QStandardItem::setToolTip(const QString & toolTip);
extern void demth_ZN13QStandardItem10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  void QStandardItem::setData(const QVariant & value, int role);
extern void _ZN13QStandardItem7setDataERK8QVarianti(void* qthis, void* arg0, int arg1);
  // proto:  QBrush QStandardItem::background();
extern void demth_ZNK13QStandardItem10backgroundEv(void* qthis);
  // proto:  QVariant QStandardItem::data(int role);
extern void _ZNK13QStandardItem4dataEi(void* qthis, int arg0);
  // proto:  void QStandardItem::QStandardItem(const QStandardItem & other);
extern void* dector_ZN13QStandardItemC1ERKS_(void* arg0);
extern void _ZN13QStandardItemC1ERKS_(void* qthis, void* arg0);
  // proto:  QStandardItem * QStandardItem::child(int row, int column);
extern void _ZNK13QStandardItem5childEii(void* qthis, int arg0, int arg1);
  // proto:  bool QStandardItem::isSelectable();
extern void demth_ZNK13QStandardItem12isSelectableEv(void* qthis);
  // proto:  QString QStandardItem::toolTip();
extern void demth_ZNK13QStandardItem7toolTipEv(void* qthis);
  // proto:  void QStandardItem::setRowCount(int rows);
extern void _ZN13QStandardItem11setRowCountEi(void* qthis, int arg0);
  // proto:  void QStandardItem::QStandardItem(const QString & text);
extern void* dector_ZN13QStandardItemC1ERK7QString(void* arg0);
extern void _ZN13QStandardItemC1ERK7QString(void* qthis, void* arg0);
  // proto:  void QStandardItem::write(QDataStream & out);
extern void _ZNK13QStandardItem5writeER11QDataStream(void* qthis, void* arg0);
  // proto:  bool QStandardItem::isDragEnabled();
extern void demth_ZNK13QStandardItem13isDragEnabledEv(void* qthis);
  // proto:  void QStandardItem::setAccessibleText(const QString & accessibleText);
extern void demth_ZN13QStandardItem17setAccessibleTextERK7QString(void* qthis, void* arg0);
  // proto:  int QStandardItem::rowCount();
extern void _ZNK13QStandardItem8rowCountEv(void* qthis);
  // proto:  void QStandardItem::removeColumn(int column);
extern void _ZN13QStandardItem12removeColumnEi(void* qthis, int arg0);
  // proto:  void QStandardItem::removeRow(int row);
extern void _ZN13QStandardItem9removeRowEi(void* qthis, int arg0);
  // proto:  int QStandardItem::columnCount();
extern void _ZNK13QStandardItem11columnCountEv(void* qthis);
  // proto:  bool QStandardItem::isTristate();
extern void demth_ZNK13QStandardItem10isTristateEv(void* qthis);
  // proto:  QStandardItem * QStandardItem::parent();
extern void _ZNK13QStandardItem6parentEv(void* qthis);
  // proto:  void QStandardItem::insertRow(int row, QStandardItem * item);
extern void demth_ZN13QStandardItem9insertRowEiPS_(void* qthis, int arg0, void* arg1);
  // proto:  void QStandardItem::QStandardItem(const QIcon & icon, const QString & text);
extern void* dector_ZN13QStandardItemC1ERK5QIconRK7QString(void* arg0, void* arg1);
extern void _ZN13QStandardItemC1ERK5QIconRK7QString(void* qthis, void* arg0, void* arg1);
  // proto:  void QStandardItem::setFont(const QFont & font);
extern void demth_ZN13QStandardItem7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QStandardItem::removeColumns(int column, int count);
extern void _ZN13QStandardItem13removeColumnsEii(void* qthis, int arg0, int arg1);
  // proto:  void QStandardItem::~QStandardItem();
extern void _ZN13QStandardItemD0Ev(void* qthis);
  // proto:  void QStandardItem::QStandardItem();
extern void* dector_ZN13QStandardItemC1Ev();
extern void _ZN13QStandardItemC1Ev(void* qthis);
  // proto:  QFont QStandardItem::font();
extern void demth_ZNK13QStandardItem4fontEv(void* qthis);
  // proto:  void QStandardItem::setEditable(bool editable);
extern void _ZN13QStandardItem11setEditableEb(void* qthis, bool arg0);
  // proto:  void QStandardItem::setText(const QString & text);
extern void demth_ZN13QStandardItem7setTextERK7QString(void* qthis, void* arg0);
  // proto:  void QStandardItem::QStandardItem(int rows, int columns);
extern void* dector_ZN13QStandardItemC1Eii(int arg0, int arg1);
extern void _ZN13QStandardItemC1Eii(void* qthis, int arg0, int arg1);
  // proto:  bool QStandardItem::isEnabled();
extern void demth_ZNK13QStandardItem9isEnabledEv(void* qthis);
  // proto:  void QStandardItem::setDropEnabled(bool dropEnabled);
extern void _ZN13QStandardItem14setDropEnabledEb(void* qthis, bool arg0);
  // proto:  void QStandardItem::setColumnCount(int columns);
extern void _ZN13QStandardItem14setColumnCountEi(void* qthis, int arg0);
  // proto:  QString QStandardItem::accessibleText();
extern void demth_ZNK13QStandardItem14accessibleTextEv(void* qthis);
  // proto:  void QStandardItem::read(QDataStream & in);
extern void _ZN13QStandardItem4readER11QDataStream(void* qthis, void* arg0);
  // proto:  void QStandardItem::setCheckable(bool checkable);
extern void _ZN13QStandardItem12setCheckableEb(void* qthis, bool arg0);
  // proto:  void QStandardItem::setDragEnabled(bool dragEnabled);
extern void _ZN13QStandardItem14setDragEnabledEb(void* qthis, bool arg0);
  // proto:  QBrush QStandardItem::foreground();
extern void demth_ZNK13QStandardItem10foregroundEv(void* qthis);
  // proto:  QStandardItem * QStandardItem::clone();
extern void _ZNK13QStandardItem5cloneEv(void* qthis);
  // proto:  void QStandardItem::removeRows(int row, int count);
extern void _ZN13QStandardItem10removeRowsEii(void* qthis, int arg0, int arg1);
  // proto:  QSize QStandardItem::sizeHint();
extern void demth_ZNK13QStandardItem8sizeHintEv(void* qthis);
  // proto:  void QStandardItem::setEnabled(bool enabled);
extern void _ZN13QStandardItem10setEnabledEb(void* qthis, bool arg0);
  // proto:  void QStandardItem::setBackground(const QBrush & brush);
extern void demth_ZN13QStandardItem13setBackgroundERK6QBrush(void* qthis, void* arg0);
  // proto:  void QStandardItem::setAccessibleDescription(const QString & accessibleDescription);
extern void demth_ZN13QStandardItem24setAccessibleDescriptionERK7QString(void* qthis, void* arg0);
  // proto:  void QStandardItem::setSizeHint(const QSize & sizeHint);
extern void demth_ZN13QStandardItem11setSizeHintERK5QSize(void* qthis, void* arg0);
  // proto:  QString QStandardItem::accessibleDescription();
extern void demth_ZNK13QStandardItem21accessibleDescriptionEv(void* qthis);
  // proto:  void QStandardItem::setTristate(bool tristate);
extern void _ZN13QStandardItem11setTristateEb(void* qthis, bool arg0);
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

// class sizeof(QStandardItemModel)=1
type QStandardItemModel struct {
  /*qbase*/ QAbstractItemModel;
  qclsinst uint64 /* *mut c_void*/;
//  _itemChanged QStandardItemModel_itemChanged_signal;
}

// class sizeof(QStandardItem)=1
type QStandardItem struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QStandardItemModel::QStandardItemModel(int rows, int columns, QObject * parent);
func NewQStandardItemModel(args ...interface{}) QStandardItemModel {
  return QStandardItemModel{}
}

  // proto:  void QStandardItemModel::clear();
func (this *QStandardItemModel) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel5clearEv
    // invoke: void clear()
    C._ZN18QStandardItemModel5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "clear", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::item(int row, int column);
func (this *QStandardItemModel) item(args ...interface{}) () {
  // item(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel4itemEii
    // invoke: QStandardItem * item(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK18QStandardItemModel4itemEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "item", args)
  }

}

  // proto:  bool QStandardItemModel::insertRow(int row, const QModelIndex & parent);
func (this *QStandardItemModel) insertRow(args ...interface{}) () {
  // insertRow(int, const class QModelIndex &)
  // insertRow(int, class QStandardItem *)
  // insertRow(int, const QList<class QStandardItem *> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  // vtys[2][1] = reflect.TypeOf(QList<QStandardItem *>{}) // "const QList<QStandardItem *> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel9insertRowEiRK11QModelIndex
    // invoke: bool insertRow(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN18QStandardItemModel9insertRowEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN18QStandardItemModel9insertRowEiP13QStandardItem
    // invoke: void insertRow(int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStandardItem).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN18QStandardItemModel9insertRowEiP13QStandardItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertRow", args)
  }

}

  // proto:  void QStandardItemModel::setItem(int row, QStandardItem * item);
func (this *QStandardItemModel) setItem(args ...interface{}) () {
  // setItem(int, class QStandardItem *)
  // setItem(int, int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel7setItemEiP13QStandardItem
    // invoke: void setItem(int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStandardItem).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN18QStandardItemModel7setItemEiP13QStandardItem(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN18QStandardItemModel7setItemEiiP13QStandardItem
    // invoke: void setItem(int, int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QStandardItem).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN18QStandardItemModel7setItemEiiP13QStandardItem(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setItem", args)
  }

}

  // proto:  QModelIndex QStandardItemModel::index(int row, int column, const QModelIndex & parent);
func (this *QStandardItemModel) index(args ...interface{}) () {
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel5indexEiiRK11QModelIndex
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK18QStandardItemModel5indexEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "index", args)
  }

}

  // proto:  bool QStandardItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
func (this *QStandardItemModel) setData(args ...interface{}) () {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setData", args)
  }

}

  // proto:  int QStandardItemModel::columnCount(const QModelIndex & parent);
func (this *QStandardItemModel) columnCount(args ...interface{}) () {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel11columnCountERK11QModelIndex
    // invoke: int columnCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel11columnCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "columnCount", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::takeItem(int row, int column);
func (this *QStandardItemModel) takeItem(args ...interface{}) () {
  // takeItem(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel8takeItemEii
    // invoke: QStandardItem * takeItem(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN18QStandardItemModel8takeItemEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeItem", args)
  }

}

  // proto:  void QStandardItemModel::setRowCount(int rows);
func (this *QStandardItemModel) setRowCount(args ...interface{}) () {
  // setRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel11setRowCountEi
    // invoke: void setRowCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel11setRowCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setRowCount", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::itemFromIndex(const QModelIndex & index);
func (this *QStandardItemModel) itemFromIndex(args ...interface{}) () {
  // itemFromIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex
    // invoke: QStandardItem * itemFromIndex(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "itemFromIndex", args)
  }

}

  // proto:  bool QStandardItemModel::insertColumn(int column, const QModelIndex & parent);
func (this *QStandardItemModel) insertColumn(args ...interface{}) () {
  // insertColumn(int, const QList<class QStandardItem *> &)
  // insertColumn(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  // vtys[0][1] = reflect.TypeOf(QList<QStandardItem *>{}) // "const QList<QStandardItem *> &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel12insertColumnEiRK11QModelIndex
    // invoke: bool insertColumn(int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QModelIndex).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertColumn", args)
  }

}

  // proto:  void QStandardItemModel::setVerticalHeaderItem(int row, QStandardItem * item);
func (this *QStandardItemModel) setVerticalHeaderItem(args ...interface{}) () {
  // setVerticalHeaderItem(int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem
    // invoke: void setVerticalHeaderItem(int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStandardItem).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setVerticalHeaderItem", args)
  }

}

  // proto:  QList<QStandardItem *> QStandardItemModel::takeColumn(int column);
func (this *QStandardItemModel) takeColumn(args ...interface{}) () {
  // takeColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel10takeColumnEi
    // invoke: QList<QStandardItem *> takeColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel10takeColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeColumn", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::takeVerticalHeaderItem(int row);
func (this *QStandardItemModel) takeVerticalHeaderItem(args ...interface{}) () {
  // takeVerticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel22takeVerticalHeaderItemEi
    // invoke: QStandardItem * takeVerticalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel22takeVerticalHeaderItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeVerticalHeaderItem", args)
  }

}

  // proto:  bool QStandardItemModel::insertColumns(int column, int count, const QModelIndex & parent);
func (this *QStandardItemModel) insertColumns(args ...interface{}) () {
  // insertColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex
    // invoke: bool insertColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertColumns", args)
  }

}

  // proto:  const QMetaObject * QStandardItemModel::metaObject();
func (this *QStandardItemModel) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C._ZNK18QStandardItemModel10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "metaObject", args)
  }

}

  // proto:  bool QStandardItemModel::insertRows(int row, int count, const QModelIndex & parent);
func (this *QStandardItemModel) insertRows(args ...interface{}) () {
  // insertRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertRows", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::invisibleRootItem();
func (this *QStandardItemModel) invisibleRootItem(args ...interface{}) () {
  // invisibleRootItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel17invisibleRootItemEv
    // invoke: QStandardItem * invisibleRootItem()
    C._ZNK18QStandardItemModel17invisibleRootItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "invisibleRootItem", args)
  }

}

  // proto:  void QStandardItemModel::setItemPrototype(const QStandardItem * item);
func (this *QStandardItemModel) setItemPrototype(args ...interface{}) () {
  // setItemPrototype(const class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStandardItem{}) // "const QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem
    // invoke: void setItemPrototype(const class QStandardItem *)
    var arg0 = args[0].(QStandardItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setItemPrototype", args)
  }

}

  // proto:  void QStandardItemModel::setHorizontalHeaderLabels(const QStringList & labels);
func (this *QStandardItemModel) setHorizontalHeaderLabels(args ...interface{}) () {
  // setHorizontalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList
    // invoke: void setHorizontalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setHorizontalHeaderLabels", args)
  }

}

  // proto:  QModelIndex QStandardItemModel::parent(const QModelIndex & child);
func (this *QStandardItemModel) parent(args ...interface{}) () {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel6parentERK11QModelIndex
    // invoke: QModelIndex parent(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel6parentERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "parent", args)
  }

}

  // proto:  bool QStandardItemModel::removeColumns(int column, int count, const QModelIndex & parent);
func (this *QStandardItemModel) removeColumns(args ...interface{}) () {
  // removeColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex
    // invoke: bool removeColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "removeColumns", args)
  }

}

  // proto:  QModelIndex QStandardItemModel::sibling(int row, int column, const QModelIndex & idx);
func (this *QStandardItemModel) sibling(args ...interface{}) () {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK18QStandardItemModel7siblingEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "sibling", args)
  }

}

  // proto:  int QStandardItemModel::sortRole();
func (this *QStandardItemModel) sortRole(args ...interface{}) () {
  // sortRole()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel8sortRoleEv
    // invoke: int sortRole()
    C._ZNK18QStandardItemModel8sortRoleEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "sortRole", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::takeHorizontalHeaderItem(int column);
func (this *QStandardItemModel) takeHorizontalHeaderItem(args ...interface{}) () {
  // takeHorizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel24takeHorizontalHeaderItemEi
    // invoke: QStandardItem * takeHorizontalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel24takeHorizontalHeaderItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeHorizontalHeaderItem", args)
  }

}

  // proto:  QModelIndex QStandardItemModel::indexFromItem(const QStandardItem * item);
func (this *QStandardItemModel) indexFromItem(args ...interface{}) () {
  // indexFromItem(const class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStandardItem{}) // "const QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem
    // invoke: QModelIndex indexFromItem(const class QStandardItem *)
    var arg0 = args[0].(QStandardItem).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "indexFromItem", args)
  }

}

  // proto:  const QStandardItem * QStandardItemModel::itemPrototype();
func (this *QStandardItemModel) itemPrototype(args ...interface{}) () {
  // itemPrototype()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel13itemPrototypeEv
    // invoke: const QStandardItem * itemPrototype()
    C._ZNK18QStandardItemModel13itemPrototypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "itemPrototype", args)
  }

}

  // proto:  void QStandardItemModel::setHorizontalHeaderItem(int column, QStandardItem * item);
func (this *QStandardItemModel) setHorizontalHeaderItem(args ...interface{}) () {
  // setHorizontalHeaderItem(int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem
    // invoke: void setHorizontalHeaderItem(int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStandardItem).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setHorizontalHeaderItem", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::horizontalHeaderItem(int column);
func (this *QStandardItemModel) horizontalHeaderItem(args ...interface{}) () {
  // horizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel20horizontalHeaderItemEi
    // invoke: QStandardItem * horizontalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel20horizontalHeaderItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "horizontalHeaderItem", args)
  }

}

  // proto:  void QStandardItemModel::appendRow(QStandardItem * item);
func (this *QStandardItemModel) appendRow(args ...interface{}) () {
  // appendRow(class QStandardItem *)
  // appendRow(const QList<class QStandardItem *> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QList<QStandardItem *>{}) // "const QList<QStandardItem *> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel9appendRowEP13QStandardItem
    // invoke: void appendRow(class QStandardItem *)
    var arg0 = args[0].(QStandardItem).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN18QStandardItemModel9appendRowEP13QStandardItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "appendRow", args)
  }

}

  // proto:  QMap<int, QVariant> QStandardItemModel::itemData(const QModelIndex & index);
func (this *QStandardItemModel) itemData(args ...interface{}) () {
  // itemData(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel8itemDataERK11QModelIndex
    // invoke: QMap<int, QVariant> itemData(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel8itemDataERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "itemData", args)
  }

}

  // proto:  void QStandardItemModel::setSortRole(int role);
func (this *QStandardItemModel) setSortRole(args ...interface{}) () {
  // setSortRole(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel11setSortRoleEi
    // invoke: void setSortRole(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel11setSortRoleEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setSortRole", args)
  }

}

  // proto:  bool QStandardItemModel::hasChildren(const QModelIndex & parent);
func (this *QStandardItemModel) hasChildren(args ...interface{}) () {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel11hasChildrenERK11QModelIndex
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel11hasChildrenERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "hasChildren", args)
  }

}

  // proto:  void QStandardItemModel::~QStandardItemModel();
func (this *QStandardItemModel) FreeQStandardItemModel(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardItemModel", "~QStandardItemModel", args)
  }

}

  // proto:  QVariant QStandardItemModel::data(const QModelIndex & index, int role);
func (this *QStandardItemModel) data(args ...interface{}) () {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK18QStandardItemModel4dataERK11QModelIndexi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "data", args)
  }

}

  // proto:  QList<QStandardItem *> QStandardItemModel::takeRow(int row);
func (this *QStandardItemModel) takeRow(args ...interface{}) () {
  // takeRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel7takeRowEi
    // invoke: QList<QStandardItem *> takeRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel7takeRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeRow", args)
  }

}

  // proto:  void QStandardItemModel::setColumnCount(int columns);
func (this *QStandardItemModel) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel14setColumnCountEi
    // invoke: void setColumnCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel14setColumnCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setColumnCount", args)
  }

}

  // proto:  QStandardItem * QStandardItemModel::verticalHeaderItem(int row);
func (this *QStandardItemModel) verticalHeaderItem(args ...interface{}) () {
  // verticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel18verticalHeaderItemEi
    // invoke: QStandardItem * verticalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel18verticalHeaderItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "verticalHeaderItem", args)
  }

}

  // proto:  bool QStandardItemModel::removeRows(int row, int count, const QModelIndex & parent);
func (this *QStandardItemModel) removeRows(args ...interface{}) () {
  // removeRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QModelIndex).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "removeRows", args)
  }

}

  // proto:  void QStandardItemModel::setVerticalHeaderLabels(const QStringList & labels);
func (this *QStandardItemModel) setVerticalHeaderLabels(args ...interface{}) () {
  // setVerticalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList
    // invoke: void setVerticalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setVerticalHeaderLabels", args)
  }

}

  // proto:  QStringList QStandardItemModel::mimeTypes();
func (this *QStandardItemModel) mimeTypes(args ...interface{}) () {
  // mimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel9mimeTypesEv
    // invoke: QStringList mimeTypes()
    C._ZNK18QStandardItemModel9mimeTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "mimeTypes", args)
  }

}

  // proto:  int QStandardItemModel::rowCount(const QModelIndex & parent);
func (this *QStandardItemModel) rowCount(args ...interface{}) () {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(QModelIndex).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK18QStandardItemModel8rowCountERK11QModelIndex(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "rowCount", args)
  }

}

  // proto:  void QStandardItem::setChild(int row, QStandardItem * item);
func (this *QStandardItem) setChild(args ...interface{}) () {
  // setChild(int, class QStandardItem *)
  // setChild(int, int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem8setChildEiPS_
    // invoke: void setChild(int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStandardItem).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN13QStandardItem8setChildEiPS_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QStandardItem8setChildEiiPS_
    // invoke: void setChild(int, int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QStandardItem).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN13QStandardItem8setChildEiiPS_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItem", "setChild", args)
  }

}

  // proto:  QStandardItemModel * QStandardItem::model();
func (this *QStandardItem) model(args ...interface{}) () {
  // model()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5modelEv
    // invoke: QStandardItemModel * model()
    C._ZNK13QStandardItem5modelEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "model", args)
  }

}

  // proto:  void QStandardItem::insertColumns(int column, int count);
func (this *QStandardItem) insertColumns(args ...interface{}) () {
  // insertColumns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem13insertColumnsEii
    // invoke: void insertColumns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStandardItem13insertColumnsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "insertColumns", args)
  }

}

  // proto:  void QStandardItem::setSelectable(bool selectable);
func (this *QStandardItem) setSelectable(args ...interface{}) () {
  // setSelectable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem13setSelectableEb
    // invoke: void setSelectable(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem13setSelectableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setSelectable", args)
  }

}

  // proto:  int QStandardItem::column();
func (this *QStandardItem) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem6columnEv
    // invoke: int column()
    C._ZNK13QStandardItem6columnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "column", args)
  }

}

  // proto:  QString QStandardItem::whatsThis();
func (this *QStandardItem) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem9whatsThisEv
    // invoke: QString whatsThis()
    C.demth_ZNK13QStandardItem9whatsThisEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "whatsThis", args)
  }

}

  // proto:  QList<QStandardItem *> QStandardItem::takeColumn(int column);
func (this *QStandardItem) takeColumn(args ...interface{}) () {
  // takeColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10takeColumnEi
    // invoke: QList<QStandardItem *> takeColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem10takeColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "takeColumn", args)
  }

}

  // proto:  void QStandardItem::setForeground(const QBrush & brush);
func (this *QStandardItem) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem13setForegroundERK6QBrush
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem13setForegroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setForeground", args)
  }

}

  // proto:  bool QStandardItem::isEditable();
func (this *QStandardItem) isEditable(args ...interface{}) () {
  // isEditable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem10isEditableEv
    // invoke: bool isEditable()
    C.demth_ZNK13QStandardItem10isEditableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isEditable", args)
  }

}

  // proto:  QIcon QStandardItem::icon();
func (this *QStandardItem) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem4iconEv
    // invoke: QIcon icon()
    C.demth_ZNK13QStandardItem4iconEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "icon", args)
  }

}

  // proto:  void QStandardItem::setWhatsThis(const QString & whatsThis);
func (this *QStandardItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem12setWhatsThisERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setWhatsThis", args)
  }

}

  // proto:  QStandardItem * QStandardItem::takeChild(int row, int column);
func (this *QStandardItem) takeChild(args ...interface{}) () {
  // takeChild(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem9takeChildEii
    // invoke: QStandardItem * takeChild(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStandardItem9takeChildEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "takeChild", args)
  }

}

  // proto:  int QStandardItem::type();
func (this *QStandardItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardItem", "type", args)
  }

}

  // proto:  QList<QStandardItem *> QStandardItem::takeRow(int row);
func (this *QStandardItem) takeRow(args ...interface{}) () {
  // takeRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7takeRowEi
    // invoke: QList<QStandardItem *> takeRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem7takeRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "takeRow", args)
  }

}

  // proto:  int QStandardItem::row();
func (this *QStandardItem) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem3rowEv
    // invoke: int row()
    C._ZNK13QStandardItem3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "row", args)
  }

}

  // proto:  bool QStandardItem::isCheckable();
func (this *QStandardItem) isCheckable(args ...interface{}) () {
  // isCheckable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem11isCheckableEv
    // invoke: bool isCheckable()
    C.demth_ZNK13QStandardItem11isCheckableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isCheckable", args)
  }

}

  // proto:  QString QStandardItem::text();
func (this *QStandardItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem4textEv
    // invoke: QString text()
    C.demth_ZNK13QStandardItem4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "text", args)
  }

}

  // proto:  void QStandardItem::insertRows(int row, int count);
func (this *QStandardItem) insertRows(args ...interface{}) () {
  // insertRows(int, int)
  // insertRows(int, const QList<class QStandardItem *> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  // vtys[1][1] = reflect.TypeOf(QList<QStandardItem *>{}) // "const QList<QStandardItem *> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10insertRowsEii
    // invoke: void insertRows(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStandardItem10insertRowsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "insertRows", args)
  }

}

  // proto:  bool QStandardItem::isDropEnabled();
func (this *QStandardItem) isDropEnabled(args ...interface{}) () {
  // isDropEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem13isDropEnabledEv
    // invoke: bool isDropEnabled()
    C.demth_ZNK13QStandardItem13isDropEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isDropEnabled", args)
  }

}

  // proto:  bool QStandardItem::hasChildren();
func (this *QStandardItem) hasChildren(args ...interface{}) () {
  // hasChildren()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem11hasChildrenEv
    // invoke: bool hasChildren()
    C._ZNK13QStandardItem11hasChildrenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "hasChildren", args)
  }

}

  // proto:  QString QStandardItem::statusTip();
func (this *QStandardItem) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem9statusTipEv
    // invoke: QString statusTip()
    C.demth_ZNK13QStandardItem9statusTipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "statusTip", args)
  }

}

  // proto:  void QStandardItem::setStatusTip(const QString & statusTip);
func (this *QStandardItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem12setStatusTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setStatusTip", args)
  }

}

  // proto:  void QStandardItem::appendRow(QStandardItem * item);
func (this *QStandardItem) appendRow(args ...interface{}) () {
  // appendRow(class QStandardItem *)
  // appendRow(const QList<class QStandardItem *> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[1] = make(map[int32]reflect.Type)
  // vtys[1][0] = reflect.TypeOf(QList<QStandardItem *>{}) // "const QList<QStandardItem *> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem9appendRowEPS_
    // invoke: void appendRow(class QStandardItem *)
    var arg0 = args[0].(QStandardItem).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem9appendRowEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "appendRow", args)
  }

}

  // proto:  QModelIndex QStandardItem::index();
func (this *QStandardItem) index(args ...interface{}) () {
  // index()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5indexEv
    // invoke: QModelIndex index()
    C._ZNK13QStandardItem5indexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "index", args)
  }

}

  // proto:  void QStandardItem::setIcon(const QIcon & icon);
func (this *QStandardItem) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setIcon", args)
  }

}

  // proto:  void QStandardItem::setToolTip(const QString & toolTip);
func (this *QStandardItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setToolTip", args)
  }

}

  // proto:  void QStandardItem::setData(const QVariant & value, int role);
func (this *QStandardItem) setData(args ...interface{}) () {
  // setData(const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7setDataERK8QVarianti
    // invoke: void setData(const class QVariant &, int)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStandardItem7setDataERK8QVarianti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "setData", args)
  }

}

  // proto:  QBrush QStandardItem::background();
func (this *QStandardItem) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem10backgroundEv
    // invoke: QBrush background()
    C.demth_ZNK13QStandardItem10backgroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "background", args)
  }

}

  // proto:  QVariant QStandardItem::data(int role);
func (this *QStandardItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem4dataEi
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK13QStandardItem4dataEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "data", args)
  }

}

  // proto:  void QStandardItem::QStandardItem(const QStandardItem & other);
func NewQStandardItem(args ...interface{}) QStandardItem {
  return QStandardItem{}
}

  // proto:  QStandardItem * QStandardItem::child(int row, int column);
func (this *QStandardItem) child(args ...interface{}) () {
  // child(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5childEii
    // invoke: QStandardItem * child(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK13QStandardItem5childEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "child", args)
  }

}

  // proto:  bool QStandardItem::isSelectable();
func (this *QStandardItem) isSelectable(args ...interface{}) () {
  // isSelectable()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem12isSelectableEv
    // invoke: bool isSelectable()
    C.demth_ZNK13QStandardItem12isSelectableEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isSelectable", args)
  }

}

  // proto:  QString QStandardItem::toolTip();
func (this *QStandardItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem7toolTipEv
    // invoke: QString toolTip()
    C.demth_ZNK13QStandardItem7toolTipEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "toolTip", args)
  }

}

  // proto:  void QStandardItem::setRowCount(int rows);
func (this *QStandardItem) setRowCount(args ...interface{}) () {
  // setRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem11setRowCountEi
    // invoke: void setRowCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem11setRowCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setRowCount", args)
  }

}

  // proto:  void QStandardItem::write(QDataStream & out);
func (this *QStandardItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK13QStandardItem5writeER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "write", args)
  }

}

  // proto:  bool QStandardItem::isDragEnabled();
func (this *QStandardItem) isDragEnabled(args ...interface{}) () {
  // isDragEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem13isDragEnabledEv
    // invoke: bool isDragEnabled()
    C.demth_ZNK13QStandardItem13isDragEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isDragEnabled", args)
  }

}

  // proto:  void QStandardItem::setAccessibleText(const QString & accessibleText);
func (this *QStandardItem) setAccessibleText(args ...interface{}) () {
  // setAccessibleText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem17setAccessibleTextERK7QString
    // invoke: void setAccessibleText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem17setAccessibleTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setAccessibleText", args)
  }

}

  // proto:  int QStandardItem::rowCount();
func (this *QStandardItem) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem8rowCountEv
    // invoke: int rowCount()
    C._ZNK13QStandardItem8rowCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "rowCount", args)
  }

}

  // proto:  void QStandardItem::removeColumn(int column);
func (this *QStandardItem) removeColumn(args ...interface{}) () {
  // removeColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem12removeColumnEi
    // invoke: void removeColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem12removeColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeColumn", args)
  }

}

  // proto:  void QStandardItem::removeRow(int row);
func (this *QStandardItem) removeRow(args ...interface{}) () {
  // removeRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem9removeRowEi
    // invoke: void removeRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem9removeRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeRow", args)
  }

}

  // proto:  int QStandardItem::columnCount();
func (this *QStandardItem) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem11columnCountEv
    // invoke: int columnCount()
    C._ZNK13QStandardItem11columnCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "columnCount", args)
  }

}

  // proto:  bool QStandardItem::isTristate();
func (this *QStandardItem) isTristate(args ...interface{}) () {
  // isTristate()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem10isTristateEv
    // invoke: bool isTristate()
    C.demth_ZNK13QStandardItem10isTristateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isTristate", args)
  }

}

  // proto:  QStandardItem * QStandardItem::parent();
func (this *QStandardItem) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem6parentEv
    // invoke: QStandardItem * parent()
    C._ZNK13QStandardItem6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "parent", args)
  }

}

  // proto:  void QStandardItem::insertRow(int row, QStandardItem * item);
func (this *QStandardItem) insertRow(args ...interface{}) () {
  // insertRow(int, class QStandardItem *)
  // insertRow(int, const QList<class QStandardItem *> &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  // vtys[1][1] = reflect.TypeOf(QList<QStandardItem *>{}) // "const QList<QStandardItem *> &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem9insertRowEiPS_
    // invoke: void insertRow(int, class QStandardItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStandardItem).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN13QStandardItem9insertRowEiPS_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "insertRow", args)
  }

}

  // proto:  void QStandardItem::setFont(const QFont & font);
func (this *QStandardItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setFont", args)
  }

}

  // proto:  void QStandardItem::removeColumns(int column, int count);
func (this *QStandardItem) removeColumns(args ...interface{}) () {
  // removeColumns(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem13removeColumnsEii
    // invoke: void removeColumns(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStandardItem13removeColumnsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeColumns", args)
  }

}

  // proto:  void QStandardItem::~QStandardItem();
func (this *QStandardItem) FreeQStandardItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStandardItem", "~QStandardItem", args)
  }

}

  // proto:  QFont QStandardItem::font();
func (this *QStandardItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem4fontEv
    // invoke: QFont font()
    C.demth_ZNK13QStandardItem4fontEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "font", args)
  }

}

  // proto:  void QStandardItem::setEditable(bool editable);
func (this *QStandardItem) setEditable(args ...interface{}) () {
  // setEditable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem11setEditableEb
    // invoke: void setEditable(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem11setEditableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setEditable", args)
  }

}

  // proto:  void QStandardItem::setText(const QString & text);
func (this *QStandardItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setText", args)
  }

}

  // proto:  bool QStandardItem::isEnabled();
func (this *QStandardItem) isEnabled(args ...interface{}) () {
  // isEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem9isEnabledEv
    // invoke: bool isEnabled()
    C.demth_ZNK13QStandardItem9isEnabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "isEnabled", args)
  }

}

  // proto:  void QStandardItem::setDropEnabled(bool dropEnabled);
func (this *QStandardItem) setDropEnabled(args ...interface{}) () {
  // setDropEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem14setDropEnabledEb
    // invoke: void setDropEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem14setDropEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setDropEnabled", args)
  }

}

  // proto:  void QStandardItem::setColumnCount(int columns);
func (this *QStandardItem) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem14setColumnCountEi
    // invoke: void setColumnCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem14setColumnCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setColumnCount", args)
  }

}

  // proto:  QString QStandardItem::accessibleText();
func (this *QStandardItem) accessibleText(args ...interface{}) () {
  // accessibleText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem14accessibleTextEv
    // invoke: QString accessibleText()
    C.demth_ZNK13QStandardItem14accessibleTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "accessibleText", args)
  }

}

  // proto:  void QStandardItem::read(QDataStream & in);
func (this *QStandardItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem4readER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "read", args)
  }

}

  // proto:  void QStandardItem::setCheckable(bool checkable);
func (this *QStandardItem) setCheckable(args ...interface{}) () {
  // setCheckable(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem12setCheckableEb
    // invoke: void setCheckable(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem12setCheckableEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setCheckable", args)
  }

}

  // proto:  void QStandardItem::setDragEnabled(bool dragEnabled);
func (this *QStandardItem) setDragEnabled(args ...interface{}) () {
  // setDragEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem14setDragEnabledEb
    // invoke: void setDragEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem14setDragEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setDragEnabled", args)
  }

}

  // proto:  QBrush QStandardItem::foreground();
func (this *QStandardItem) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem10foregroundEv
    // invoke: QBrush foreground()
    C.demth_ZNK13QStandardItem10foregroundEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "foreground", args)
  }

}

  // proto:  QStandardItem * QStandardItem::clone();
func (this *QStandardItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5cloneEv
    // invoke: QStandardItem * clone()
    C._ZNK13QStandardItem5cloneEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "clone", args)
  }

}

  // proto:  void QStandardItem::removeRows(int row, int count);
func (this *QStandardItem) removeRows(args ...interface{}) () {
  // removeRows(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10removeRowsEii
    // invoke: void removeRows(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN13QStandardItem10removeRowsEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeRows", args)
  }

}

  // proto:  QSize QStandardItem::sizeHint();
func (this *QStandardItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem8sizeHintEv
    // invoke: QSize sizeHint()
    C.demth_ZNK13QStandardItem8sizeHintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "sizeHint", args)
  }

}

  // proto:  void QStandardItem::setEnabled(bool enabled);
func (this *QStandardItem) setEnabled(args ...interface{}) () {
  // setEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10setEnabledEb
    // invoke: void setEnabled(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem10setEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setEnabled", args)
  }

}

  // proto:  void QStandardItem::setBackground(const QBrush & brush);
func (this *QStandardItem) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setBackground", args)
  }

}

  // proto:  void QStandardItem::setAccessibleDescription(const QString & accessibleDescription);
func (this *QStandardItem) setAccessibleDescription(args ...interface{}) () {
  // setAccessibleDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem24setAccessibleDescriptionERK7QString
    // invoke: void setAccessibleDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem24setAccessibleDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setAccessibleDescription", args)
  }

}

  // proto:  void QStandardItem::setSizeHint(const QSize & sizeHint);
func (this *QStandardItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem11setSizeHintERK5QSize
    // invoke: void setSizeHint(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN13QStandardItem11setSizeHintERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setSizeHint", args)
  }

}

  // proto:  QString QStandardItem::accessibleDescription();
func (this *QStandardItem) accessibleDescription(args ...interface{}) () {
  // accessibleDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem21accessibleDescriptionEv
    // invoke: QString accessibleDescription()
    C.demth_ZNK13QStandardItem21accessibleDescriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "accessibleDescription", args)
  }

}

  // proto:  void QStandardItem::setTristate(bool tristate);
func (this *QStandardItem) setTristate(args ...interface{}) () {
  // setTristate(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem11setTristateEb
    // invoke: void setTristate(_Bool)
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
    C._ZN13QStandardItem11setTristateEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setTristate", args)
  }

}

// <= body block end

