package qtgui
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QStandardItemModel::columnCount(const QModelIndex & parent);
extern int32_t C_ZNK18QStandardItemModel11columnCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QStandardItemModel::QStandardItemModel(int rows, int columns, QObject * parent);
extern void* C_ZN18QStandardItemModelC2EiiP7QObject(int32_t arg0, int32_t arg1, void* arg2); // 3
  // proto:  void QStandardItemModel::QStandardItemModel(QObject * parent);
extern void* C_ZN18QStandardItemModelC2EP7QObject(void* arg0); // 3
  // proto:  void QStandardItemModel::setSortRole(int role);
extern void C_ZN18QStandardItemModel11setSortRoleEi(void* qthis, int32_t arg0); // 4
  // proto:  QStandardItem * QStandardItemModel::verticalHeaderItem(int row);
extern void* C_ZNK18QStandardItemModel18verticalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QStandardItemModel::insertColumn(int column, const QModelIndex & parent);
extern bool C_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QStandardItemModel::insertColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QModelIndex QStandardItemModel::index(int row, int column, const QModelIndex & parent);
extern void* C_ZNK18QStandardItemModel5indexEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QStandardItemModel::setColumnCount(int columns);
extern void C_ZN18QStandardItemModel14setColumnCountEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QStandardItemModel::removeRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QStandardItemModel::sortRole();
extern int32_t C_ZNK18QStandardItemModel8sortRoleEv(void* qthis); // 4
  // proto:  void QStandardItemModel::setHorizontalHeaderItem(int column, QStandardItem * item);
extern void C_ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QList<QStandardItem *> QStandardItemModel::takeRow(int row);
extern void C_ZN18QStandardItemModel7takeRowEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QStandardItemModel::setData(const QModelIndex & index, const QVariant & value, int role);
extern bool C_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
  // proto:  QStandardItem * QStandardItemModel::takeItem(int row, int column);
extern void* C_ZN18QStandardItemModel8takeItemEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QMap<int, QVariant> QStandardItemModel::itemData(const QModelIndex & index);
extern void C_ZNK18QStandardItemModel8itemDataERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QStandardItemModel::parent(const QModelIndex & child);
extern void* C_ZNK18QStandardItemModel6parentERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  QModelIndex QStandardItemModel::sibling(int row, int column, const QModelIndex & idx);
extern void* C_ZNK18QStandardItemModel7siblingEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QStandardItem * QStandardItemModel::invisibleRootItem();
extern void* C_ZNK18QStandardItemModel17invisibleRootItemEv(void* qthis); // 4
  // proto:  int QStandardItemModel::rowCount(const QModelIndex & parent);
extern int32_t C_ZNK18QStandardItemModel8rowCountERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QStandardItemModel::setItem(int row, int column, QStandardItem * item);
extern void C_ZN18QStandardItemModel7setItemEiiP13QStandardItem(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QStandardItemModel::setItem(int row, QStandardItem * item);
extern void C_ZN18QStandardItemModel7setItemEiP13QStandardItem(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QStandardItem * QStandardItemModel::takeHorizontalHeaderItem(int column);
extern void* C_ZN18QStandardItemModel24takeHorizontalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  QModelIndex QStandardItemModel::indexFromItem(const QStandardItem * item);
extern void* C_ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem(void* qthis, void* arg0); // 4
  // proto:  QStandardItem * QStandardItemModel::horizontalHeaderItem(int column);
extern void* C_ZNK18QStandardItemModel20horizontalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStandardItemModel::setHorizontalHeaderLabels(const QStringList & labels);
extern void C_ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QVariant QStandardItemModel::data(const QModelIndex & index, int role);
extern void* C_ZNK18QStandardItemModel4dataERK11QModelIndexi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QStandardItemModel::appendRow(QStandardItem * item);
extern void C_ZN18QStandardItemModel9appendRowEP13QStandardItem(void* qthis, void* arg0); // 2
  // proto:  QStringList QStandardItemModel::mimeTypes();
extern void C_ZNK18QStandardItemModel9mimeTypesEv(void* qthis); // 4
  // proto:  QStandardItem * QStandardItemModel::takeVerticalHeaderItem(int row);
extern void* C_ZN18QStandardItemModel22takeVerticalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  const QMetaObject * QStandardItemModel::metaObject();
extern void C_ZNK18QStandardItemModel10metaObjectEv(void* qthis); // 4
  // proto:  bool QStandardItemModel::hasChildren(const QModelIndex & parent);
extern bool C_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QStandardItemModel::setItemPrototype(const QStandardItem * item);
extern void C_ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem(void* qthis, void* arg0); // 4
  // proto:  void QStandardItemModel::setVerticalHeaderItem(int row, QStandardItem * item);
extern void C_ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QStandardItemModel::clear();
extern void C_ZN18QStandardItemModel5clearEv(void* qthis); // 4
  // proto:  const QStandardItem * QStandardItemModel::itemPrototype();
extern void* C_ZNK18QStandardItemModel13itemPrototypeEv(void* qthis); // 4
  // proto:  bool QStandardItemModel::insertRows(int row, int count, const QModelIndex & parent);
extern bool C_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  bool QStandardItemModel::insertRow(int row, const QModelIndex & parent);
extern bool C_ZN18QStandardItemModel9insertRowEiRK11QModelIndex(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QStandardItemModel::insertRow(int row, QStandardItem * item);
extern void C_ZN18QStandardItemModel9insertRowEiP13QStandardItem(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QStandardItem * QStandardItemModel::item(int row, int column);
extern void* C_ZNK18QStandardItemModel4itemEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QStandardItemModel::setRowCount(int rows);
extern void C_ZN18QStandardItemModel11setRowCountEi(void* qthis, int32_t arg0); // 4
  // proto:  Qt::ItemFlags QStandardItemModel::flags(const QModelIndex & index);
extern void C_ZNK18QStandardItemModel5flagsERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  Qt::DropActions QStandardItemModel::supportedDropActions();
extern void C_ZNK18QStandardItemModel20supportedDropActionsEv(void* qthis); // 4
  // proto:  QList<QStandardItem *> QStandardItemModel::takeColumn(int column);
extern void C_ZN18QStandardItemModel10takeColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStandardItemModel::~QStandardItemModel();
extern void C_ZN18QStandardItemModelD2Ev(void* qthis); // 4
  // proto:  bool QStandardItemModel::removeColumns(int column, int count, const QModelIndex & parent);
extern bool C_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QStandardItem * QStandardItemModel::itemFromIndex(const QModelIndex & index);
extern void* C_ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex(void* qthis, void* arg0); // 4
  // proto:  void QStandardItemModel::setVerticalHeaderLabels(const QStringList & labels);
extern void C_ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QStandardItem * QStandardItem::takeChild(int row, int column);
extern void* C_ZN13QStandardItem9takeChildEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QStandardItem::columnCount();
extern int32_t C_ZNK13QStandardItem11columnCountEv(void* qthis); // 4
  // proto:  void QStandardItem::setText(const QString & text);
extern void C_ZN13QStandardItem7setTextERK7QString(void* qthis, void* arg0); // 2
  // proto:  bool QStandardItem::isEditable();
extern bool C_ZNK13QStandardItem10isEditableEv(void* qthis); // 2
  // proto:  void QStandardItem::setEnabled(bool enabled);
extern void C_ZN13QStandardItem10setEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QString QStandardItem::text();
extern void* C_ZNK13QStandardItem4textEv(void* qthis); // 2
  // proto:  void QStandardItem::setDragEnabled(bool dragEnabled);
extern void C_ZN13QStandardItem14setDragEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QStandardItem::isTristate();
extern bool C_ZNK13QStandardItem10isTristateEv(void* qthis); // 2
  // proto:  void QStandardItem::insertRow(int row, QStandardItem * item);
extern void C_ZN13QStandardItem9insertRowEiPS_(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QStandardItem::isCheckable();
extern bool C_ZNK13QStandardItem11isCheckableEv(void* qthis); // 2
  // proto:  bool QStandardItem::isSelectable();
extern bool C_ZNK13QStandardItem12isSelectableEv(void* qthis); // 2
  // proto:  void QStandardItem::setAccessibleDescription(const QString & accessibleDescription);
extern void C_ZN13QStandardItem24setAccessibleDescriptionERK7QString(void* qthis, void* arg0); // 2
  // proto:  QFont QStandardItem::font();
extern void* C_ZNK13QStandardItem4fontEv(void* qthis); // 2
  // proto:  QString QStandardItem::accessibleText();
extern void* C_ZNK13QStandardItem14accessibleTextEv(void* qthis); // 2
  // proto:  void QStandardItem::QStandardItem(int rows, int columns);
extern void* C_ZN13QStandardItemC2Eii(int32_t arg0, int32_t arg1); // 3
  // proto:  void QStandardItem::QStandardItem(const QIcon & icon, const QString & text);
extern void* C_ZN13QStandardItemC2ERK5QIconRK7QString(void* arg0, void* arg1); // 3
  // proto:  void QStandardItem::QStandardItem(const QString & text);
extern void* C_ZN13QStandardItemC2ERK7QString(void* arg0); // 3
  // proto:  void QStandardItem::QStandardItem();
extern void* C_ZN13QStandardItemC2Ev(); // 3
  // proto:  int QStandardItem::row();
extern int32_t C_ZNK13QStandardItem3rowEv(void* qthis); // 4
  // proto:  QModelIndex QStandardItem::index();
extern void* C_ZNK13QStandardItem5indexEv(void* qthis); // 4
  // proto:  void QStandardItem::setColumnCount(int columns);
extern void C_ZN13QStandardItem14setColumnCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStandardItem::setForeground(const QBrush & brush);
extern void C_ZN13QStandardItem13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  void QStandardItem::~QStandardItem();
extern void C_ZN13QStandardItemD2Ev(void* qthis); // 4
  // proto:  bool QStandardItem::isEnabled();
extern bool C_ZNK13QStandardItem9isEnabledEv(void* qthis); // 2
  // proto:  void QStandardItem::setEditable(bool editable);
extern void C_ZN13QStandardItem11setEditableEb(void* qthis, bool arg0); // 4
  // proto:  QString QStandardItem::whatsThis();
extern void* C_ZNK13QStandardItem9whatsThisEv(void* qthis); // 2
  // proto:  void QStandardItem::read(QDataStream & in);
extern void C_ZN13QStandardItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QStandardItem::removeRows(int row, int count);
extern void C_ZN13QStandardItem10removeRowsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QStandardItem::setCheckable(bool checkable);
extern void C_ZN13QStandardItem12setCheckableEb(void* qthis, bool arg0); // 4
  // proto:  QList<QStandardItem *> QStandardItem::takeRow(int row);
extern void C_ZN13QStandardItem7takeRowEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStandardItem::setChild(int row, QStandardItem * item);
extern void C_ZN13QStandardItem8setChildEiPS_(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QStandardItem::setChild(int row, int column, QStandardItem * item);
extern void C_ZN13QStandardItem8setChildEiiPS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QStandardItem::type();
extern int32_t C_ZNK13QStandardItem4typeEv(void* qthis); // 4
  // proto:  void QStandardItem::setSelectable(bool selectable);
extern void C_ZN13QStandardItem13setSelectableEb(void* qthis, bool arg0); // 4
  // proto:  void QStandardItem::setBackground(const QBrush & brush);
extern void C_ZN13QStandardItem13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  void QStandardItem::setIcon(const QIcon & icon);
extern void C_ZN13QStandardItem7setIconERK5QIcon(void* qthis, void* arg0); // 2
  // proto:  QBrush QStandardItem::foreground();
extern void* C_ZNK13QStandardItem10foregroundEv(void* qthis); // 2
  // proto:  QStandardItem * QStandardItem::parent();
extern void* C_ZNK13QStandardItem6parentEv(void* qthis); // 4
  // proto:  bool QStandardItem::isDropEnabled();
extern bool C_ZNK13QStandardItem13isDropEnabledEv(void* qthis); // 2
  // proto:  QString QStandardItem::statusTip();
extern void* C_ZNK13QStandardItem9statusTipEv(void* qthis); // 2
  // proto:  QStandardItem * QStandardItem::clone();
extern void* C_ZNK13QStandardItem5cloneEv(void* qthis); // 4
  // proto:  void QStandardItem::removeRow(int row);
extern void C_ZN13QStandardItem9removeRowEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QStandardItem::toolTip();
extern void* C_ZNK13QStandardItem7toolTipEv(void* qthis); // 2
  // proto:  int QStandardItem::rowCount();
extern int32_t C_ZNK13QStandardItem8rowCountEv(void* qthis); // 4
  // proto:  void QStandardItem::insertRows(int row, int count);
extern void C_ZN13QStandardItem10insertRowsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QStandardItem::setWhatsThis(const QString & whatsThis);
extern void C_ZN13QStandardItem12setWhatsThisERK7QString(void* qthis, void* arg0); // 2
  // proto:  QBrush QStandardItem::background();
extern void* C_ZNK13QStandardItem10backgroundEv(void* qthis); // 2
  // proto:  QStandardItem * QStandardItem::child(int row, int column);
extern void* C_ZNK13QStandardItem5childEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QStandardItem::setTristate(bool tristate);
extern void C_ZN13QStandardItem11setTristateEb(void* qthis, bool arg0); // 4
  // proto:  void QStandardItem::removeColumn(int column);
extern void C_ZN13QStandardItem12removeColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QStandardItem::accessibleDescription();
extern void* C_ZNK13QStandardItem21accessibleDescriptionEv(void* qthis); // 2
  // proto:  QVariant QStandardItem::data(int role);
extern void* C_ZNK13QStandardItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStandardItem::appendRow(QStandardItem * item);
extern void C_ZN13QStandardItem9appendRowEPS_(void* qthis, void* arg0); // 2
  // proto:  void QStandardItem::setAccessibleText(const QString & accessibleText);
extern void C_ZN13QStandardItem17setAccessibleTextERK7QString(void* qthis, void* arg0); // 2
  // proto:  QIcon QStandardItem::icon();
extern void* C_ZNK13QStandardItem4iconEv(void* qthis); // 2
  // proto:  void QStandardItem::setSizeHint(const QSize & sizeHint);
extern void C_ZN13QStandardItem11setSizeHintERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QStandardItem::setToolTip(const QString & toolTip);
extern void C_ZN13QStandardItem10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  QSize QStandardItem::sizeHint();
extern void* C_ZNK13QStandardItem8sizeHintEv(void* qthis); // 2
  // proto:  bool QStandardItem::hasChildren();
extern bool C_ZNK13QStandardItem11hasChildrenEv(void* qthis); // 4
  // proto:  int QStandardItem::column();
extern int32_t C_ZNK13QStandardItem6columnEv(void* qthis); // 4
  // proto:  void QStandardItem::insertColumns(int column, int count);
extern void C_ZN13QStandardItem13insertColumnsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QStandardItem::setStatusTip(const QString & statusTip);
extern void C_ZN13QStandardItem12setStatusTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QStandardItem::setDropEnabled(bool dropEnabled);
extern void C_ZN13QStandardItem14setDropEnabledEb(void* qthis, bool arg0); // 4
  // proto:  bool QStandardItem::isDragEnabled();
extern bool C_ZNK13QStandardItem13isDragEnabledEv(void* qthis); // 2
  // proto:  void QStandardItem::setRowCount(int rows);
extern void C_ZN13QStandardItem11setRowCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QStandardItem::write(QDataStream & out);
extern void C_ZNK13QStandardItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  Qt::ItemFlags QStandardItem::flags();
extern void C_ZNK13QStandardItem5flagsEv(void* qthis); // 4
  // proto:  Qt::CheckState QStandardItem::checkState();
extern void C_ZNK13QStandardItem10checkStateEv(void* qthis); // 2
  // proto:  void QStandardItem::setFont(const QFont & font);
extern void C_ZN13QStandardItem7setFontERK5QFont(void* qthis, void* arg0); // 2
  // proto:  QList<QStandardItem *> QStandardItem::takeColumn(int column);
extern void C_ZN13QStandardItem10takeColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  QStandardItemModel * QStandardItem::model();
extern void* C_ZNK13QStandardItem5modelEv(void* qthis); // 4
  // proto:  void QStandardItem::setData(const QVariant & value, int role);
extern void C_ZN13QStandardItem7setDataERK8QVarianti(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QStandardItem::removeColumns(int column, int count);
extern void C_ZN13QStandardItem13removeColumnsEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  Qt::Alignment QStandardItem::textAlignment();
extern void C_ZNK13QStandardItem13textAlignmentEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QStandardItemModel)=1
type QStandardItemModel struct {
  /*qbase*/ qtcore.QAbstractItemModel;
  Qclsinst unsafe.Pointer /* *C.void */;
//  _itemChanged QStandardItemModel_itemChanged_signal;
}

// class sizeof(QStandardItem)=1
type QStandardItem struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount(const class QModelIndex &)
func (this *QStandardItemModel) Columncount(args ...interface{}) (ret interface{}) {
  // columnCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel11columnCountERK11QModelIndex
    // invoke: int columnCount(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel11columnCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "columnCount", args)
  }

  return
}

// QStandardItemModel(int, int, class QObject *)
func NewQStandardItemModel(args ...interface{}) *QStandardItemModel {
  // QStandardItemModel(int, int, class QObject *)
  // QStandardItemModel(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModelC1EiiP7QObject
    // invoke: void QStandardItemModel(int, int, class QObject *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStandardItemModelC2EiiP7QObject(arg0, arg1, arg2)
    return &QStandardItemModel{Qclsinst:qthis}
  case 1:
    // invoke: _ZN18QStandardItemModelC1EP7QObject
    // invoke: void QStandardItemModel(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN18QStandardItemModelC2EP7QObject(arg0)
    return &QStandardItemModel{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStandardItemModel", "QStandardItemModel", args)
  }

  return nil // QStandardItemModel{Qclsinst:qthis}
}

// setSortRole(int)
func (this *QStandardItemModel) Setsortrole(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel11setSortRoleEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setSortRole", args)
  }

  return
}

// verticalHeaderItem(int)
func (this *QStandardItemModel) Verticalheaderitem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel18verticalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "verticalHeaderItem", args)
  }

  return
}

// insertColumn(int, const class QModelIndex &)
func (this *QStandardItemModel) Insertcolumn(args ...interface{}) (ret interface{}) {
  // insertColumn(int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel12insertColumnEiRK11QModelIndex
    // invoke: bool insertColumn(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN18QStandardItemModel12insertColumnEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertColumn", args)
  }

  return
}

// insertColumns(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Insertcolumns(args ...interface{}) (ret interface{}) {
  // insertColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex
    // invoke: bool insertColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN18QStandardItemModel13insertColumnsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertColumns", args)
  }

  return
}

// index(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Index(args ...interface{}) (ret interface{}) {
  // index(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel5indexEiiRK11QModelIndex
    // invoke: QModelIndex index(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK18QStandardItemModel5indexEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "index", args)
  }

  return
}

// setColumnCount(int)
func (this *QStandardItemModel) Setcolumncount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel14setColumnCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setColumnCount", args)
  }

  return
}

// removeRows(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Removerows(args ...interface{}) (ret interface{}) {
  // removeRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex
    // invoke: bool removeRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN18QStandardItemModel10removeRowsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "removeRows", args)
  }

  return
}

// sortRole()
func (this *QStandardItemModel) Sortrole(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QStandardItemModel8sortRoleEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "sortRole", args)
  }

  return
}

// setHorizontalHeaderItem(int, class QStandardItem *)
func (this *QStandardItemModel) Sethorizontalheaderitem(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QStandardItemModel23setHorizontalHeaderItemEiP13QStandardItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setHorizontalHeaderItem", args)
  }

  return
}

// takeRow(int)
func (this *QStandardItemModel) Takerow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel7takeRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeRow", args)
  }

  return
}

// setData(const class QModelIndex &, const class QVariant &, int)
func (this *QStandardItemModel) Setdata(args ...interface{}) (ret interface{}) {
  // setData(const class QModelIndex &, const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti
    // invoke: bool setData(const class QModelIndex &, const class QVariant &, int)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN18QStandardItemModel7setDataERK11QModelIndexRK8QVarianti(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setData", args)
  }

  return
}

// takeItem(int, int)
func (this *QStandardItemModel) Takeitem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN18QStandardItemModel8takeItemEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeItem", args)
  }

  return
}

// itemData(const class QModelIndex &)
func (this *QStandardItemModel) Itemdata(args ...interface{}) () {
  // itemData(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel8itemDataERK11QModelIndex
    // invoke: QMap<int, QVariant> itemData(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QStandardItemModel8itemDataERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "itemData", args)
  }

  return
}

// parent(const class QModelIndex &)
func (this *QStandardItemModel) Parent(args ...interface{}) (ret interface{}) {
  // parent(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel6parentERK11QModelIndex
    // invoke: QModelIndex parent(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel6parentERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "parent", args)
  }

  return
}

// sibling(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Sibling(args ...interface{}) (ret interface{}) {
  // sibling(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel7siblingEiiRK11QModelIndex
    // invoke: QModelIndex sibling(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK18QStandardItemModel7siblingEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "sibling", args)
  }

  return
}

// invisibleRootItem()
func (this *QStandardItemModel) Invisiblerootitem(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QStandardItemModel17invisibleRootItemEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "invisibleRootItem", args)
  }

  return
}

// rowCount(const class QModelIndex &)
func (this *QStandardItemModel) Rowcount(args ...interface{}) (ret interface{}) {
  // rowCount(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel8rowCountERK11QModelIndex
    // invoke: int rowCount(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel8rowCountERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "rowCount", args)
  }

  return
}

// setItem(int, int, class QStandardItem *)
func (this *QStandardItemModel) Setitem(args ...interface{}) () {
  // setItem(int, int, class QStandardItem *)
  // setItem(int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel7setItemEiiP13QStandardItem
    // invoke: void setItem(int, int, class QStandardItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN18QStandardItemModel7setItemEiiP13QStandardItem(this.Qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN18QStandardItemModel7setItemEiP13QStandardItem
    // invoke: void setItem(int, class QStandardItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QStandardItemModel7setItemEiP13QStandardItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setItem", args)
  }

  return
}

// takeHorizontalHeaderItem(int)
func (this *QStandardItemModel) Takehorizontalheaderitem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QStandardItemModel24takeHorizontalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeHorizontalHeaderItem", args)
  }

  return
}

// indexFromItem(const class QStandardItem *)
func (this *QStandardItemModel) Indexfromitem(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel13indexFromItemEPK13QStandardItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "indexFromItem", args)
  }

  return
}

// horizontalHeaderItem(int)
func (this *QStandardItemModel) Horizontalheaderitem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel20horizontalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "horizontalHeaderItem", args)
  }

  return
}

// setHorizontalHeaderLabels(const class QStringList &)
func (this *QStandardItemModel) Sethorizontalheaderlabels(args ...interface{}) () {
  // setHorizontalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList
    // invoke: void setHorizontalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel25setHorizontalHeaderLabelsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setHorizontalHeaderLabels", args)
  }

  return
}

// data(const class QModelIndex &, int)
func (this *QStandardItemModel) Data(args ...interface{}) (ret interface{}) {
  // data(const class QModelIndex &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel4dataERK11QModelIndexi
    // invoke: QVariant data(const class QModelIndex &, int)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK18QStandardItemModel4dataERK11QModelIndexi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "data", args)
  }

  return
}

// appendRow(class QStandardItem *)
func (this *QStandardItemModel) Appendrow(args ...interface{}) () {
  // appendRow(class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel9appendRowEP13QStandardItem
    // invoke: void appendRow(class QStandardItem *)
    var arg0 = args[0].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel9appendRowEP13QStandardItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "appendRow", args)
  }

  return
}

// mimeTypes()
func (this *QStandardItemModel) Mimetypes(args ...interface{}) () {
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
    C.C_ZNK18QStandardItemModel9mimeTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "mimeTypes", args)
  }

  return
}

// takeVerticalHeaderItem(int)
func (this *QStandardItemModel) Takeverticalheaderitem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN18QStandardItemModel22takeVerticalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeVerticalHeaderItem", args)
  }

  return
}

// metaObject()
func (this *QStandardItemModel) Metaobject(args ...interface{}) () {
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
    C.C_ZNK18QStandardItemModel10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "metaObject", args)
  }

  return
}

// hasChildren(const class QModelIndex &)
func (this *QStandardItemModel) Haschildren(args ...interface{}) (ret interface{}) {
  // hasChildren(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel11hasChildrenERK11QModelIndex
    // invoke: bool hasChildren(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel11hasChildrenERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "hasChildren", args)
  }

  return
}

// setItemPrototype(const class QStandardItem *)
func (this *QStandardItemModel) Setitemprototype(args ...interface{}) () {
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
    var arg0 = args[0].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel16setItemPrototypeEPK13QStandardItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setItemPrototype", args)
  }

  return
}

// setVerticalHeaderItem(int, class QStandardItem *)
func (this *QStandardItemModel) Setverticalheaderitem(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QStandardItemModel21setVerticalHeaderItemEiP13QStandardItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setVerticalHeaderItem", args)
  }

  return
}

// clear()
func (this *QStandardItemModel) Clear(args ...interface{}) () {
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
    C.C_ZN18QStandardItemModel5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "clear", args)
  }

  return
}

// itemPrototype()
func (this *QStandardItemModel) Itemprototype(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK18QStandardItemModel13itemPrototypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "const QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "itemPrototype", args)
  }

  return
}

// insertRows(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Insertrows(args ...interface{}) (ret interface{}) {
  // insertRows(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex
    // invoke: bool insertRows(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN18QStandardItemModel10insertRowsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertRows", args)
  }

  return
}

// insertRow(int, const class QModelIndex &)
func (this *QStandardItemModel) Insertrow(args ...interface{}) (ret interface{}) {
  // insertRow(int, const class QModelIndex &)
  // insertRow(int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel9insertRowEiRK11QModelIndex
    // invoke: bool insertRow(int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN18QStandardItemModel9insertRowEiRK11QModelIndex(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN18QStandardItemModel9insertRowEiP13QStandardItem
    // invoke: void insertRow(int, class QStandardItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN18QStandardItemModel9insertRowEiP13QStandardItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "insertRow", args)
  }

  return
}

// item(int, int)
func (this *QStandardItemModel) Item(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK18QStandardItemModel4itemEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "item", args)
  }

  return
}

// setRowCount(int)
func (this *QStandardItemModel) Setrowcount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel11setRowCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setRowCount", args)
  }

  return
}

// flags(const class QModelIndex &)
func (this *QStandardItemModel) Flags(args ...interface{}) () {
  // flags(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel5flagsERK11QModelIndex
    // invoke: Qt::ItemFlags flags(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK18QStandardItemModel5flagsERK11QModelIndex(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "flags", args)
  }

  return
}

// supportedDropActions()
func (this *QStandardItemModel) Supporteddropactions(args ...interface{}) () {
  // supportedDropActions()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel20supportedDropActionsEv
    // invoke: Qt::DropActions supportedDropActions()
    C.C_ZNK18QStandardItemModel20supportedDropActionsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "supportedDropActions", args)
  }

  return
}

// takeColumn(int)
func (this *QStandardItemModel) Takecolumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel10takeColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "takeColumn", args)
  }

  return
}

// ~QStandardItemModel()
func (this *QStandardItemModel) Freeqstandarditemmodel(args ...interface{}) () {
  // ~QStandardItemModel()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModelD0Ev
    // invoke: void ~QStandardItemModel()
    C.C_ZN18QStandardItemModelD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "~QStandardItemModel", args)
  }

  return
}

// removeColumns(int, int, const class QModelIndex &)
func (this *QStandardItemModel) Removecolumns(args ...interface{}) (ret interface{}) {
  // removeColumns(int, int, const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex
    // invoke: bool removeColumns(int, int, const class QModelIndex &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN18QStandardItemModel13removeColumnsEiiRK11QModelIndex(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "removeColumns", args)
  }

  return
}

// itemFromIndex(const class QModelIndex &)
func (this *QStandardItemModel) Itemfromindex(args ...interface{}) (ret interface{}) {
  // itemFromIndex(const class QModelIndex &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QModelIndex{}) // "const QModelIndex &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex
    // invoke: QStandardItem * itemFromIndex(const class QModelIndex &)
    var arg0 = args[0].(*qtcore.QModelIndex).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK18QStandardItemModel13itemFromIndexERK11QModelIndex(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItemModel", "itemFromIndex", args)
  }

  return
}

// setVerticalHeaderLabels(const class QStringList &)
func (this *QStandardItemModel) Setverticalheaderlabels(args ...interface{}) () {
  // setVerticalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList
    // invoke: void setVerticalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN18QStandardItemModel23setVerticalHeaderLabelsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItemModel", "setVerticalHeaderLabels", args)
  }

  return
}

// takeChild(int, int)
func (this *QStandardItem) Takechild(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN13QStandardItem9takeChildEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "takeChild", args)
  }

  return
}

// columnCount()
func (this *QStandardItem) Columncount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem11columnCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "columnCount", args)
  }

  return
}

// setText(const class QString &)
func (this *QStandardItem) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setText", args)
  }

  return
}

// isEditable()
func (this *QStandardItem) Iseditable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem10isEditableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isEditable", args)
  }

  return
}

// setEnabled(_Bool)
func (this *QStandardItem) Setenabled(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem10setEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setEnabled", args)
  }

  return
}

// text()
func (this *QStandardItem) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "text", args)
  }

  return
}

// setDragEnabled(_Bool)
func (this *QStandardItem) Setdragenabled(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem14setDragEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setDragEnabled", args)
  }

  return
}

// isTristate()
func (this *QStandardItem) Istristate(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem10isTristateEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isTristate", args)
  }

  return
}

// insertRow(int, class QStandardItem *)
func (this *QStandardItem) Insertrow(args ...interface{}) () {
  // insertRow(int, class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem9insertRowEiPS_
    // invoke: void insertRow(int, class QStandardItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem9insertRowEiPS_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "insertRow", args)
  }

  return
}

// isCheckable()
func (this *QStandardItem) Ischeckable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem11isCheckableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isCheckable", args)
  }

  return
}

// isSelectable()
func (this *QStandardItem) Isselectable(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem12isSelectableEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isSelectable", args)
  }

  return
}

// setAccessibleDescription(const class QString &)
func (this *QStandardItem) Setaccessibledescription(args ...interface{}) () {
  // setAccessibleDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem24setAccessibleDescriptionERK7QString
    // invoke: void setAccessibleDescription(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem24setAccessibleDescriptionERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setAccessibleDescription", args)
  }

  return
}

// font()
func (this *QStandardItem) Font(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "font", args)
  }

  return
}

// accessibleText()
func (this *QStandardItem) Accessibletext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem14accessibleTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "accessibleText", args)
  }

  return
}

// QStandardItem(int, int)
func NewQStandardItem(args ...interface{}) *QStandardItem {
  // QStandardItem(int, int)
  // QStandardItem(const class QIcon &, const class QString &)
  // QStandardItem(const class QString &)
  // QStandardItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItemC1Eii
    // invoke: void QStandardItem(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStandardItemC2Eii(arg0, arg1)
    return &QStandardItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QStandardItemC1ERK5QIconRK7QString
    // invoke: void QStandardItem(const class QIcon &, const class QString &)
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStandardItemC2ERK5QIconRK7QString(arg0, arg1)
    return &QStandardItem{Qclsinst:qthis}
  case 2:
    // invoke: _ZN13QStandardItemC1ERK7QString
    // invoke: void QStandardItem(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStandardItemC2ERK7QString(arg0)
    return &QStandardItem{Qclsinst:qthis}
  case 3:
    // invoke: _ZN13QStandardItemC1Ev
    // invoke: void QStandardItem()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QStandardItemC2Ev()
    return &QStandardItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStandardItem", "QStandardItem", args)
  }

  return nil // QStandardItem{Qclsinst:qthis}
}

// row()
func (this *QStandardItem) Row(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem3rowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "row", args)
  }

  return
}

// index()
func (this *QStandardItem) Index(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem5indexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QModelIndex{}) // "QModelIndex"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "index", args)
  }

  return
}

// setColumnCount(int)
func (this *QStandardItem) Setcolumncount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem14setColumnCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setColumnCount", args)
  }

  return
}

// setForeground(const class QBrush &)
func (this *QStandardItem) Setforeground(args ...interface{}) () {
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
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem13setForegroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setForeground", args)
  }

  return
}

// ~QStandardItem()
func (this *QStandardItem) Freeqstandarditem(args ...interface{}) () {
  // ~QStandardItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItemD0Ev
    // invoke: void ~QStandardItem()
    C.C_ZN13QStandardItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "~QStandardItem", args)
  }

  return
}

// isEnabled()
func (this *QStandardItem) Isenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem9isEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isEnabled", args)
  }

  return
}

// setEditable(_Bool)
func (this *QStandardItem) Seteditable(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem11setEditableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setEditable", args)
  }

  return
}

// whatsThis()
func (this *QStandardItem) Whatsthis(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem9whatsThisEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "whatsThis", args)
  }

  return
}

// read(class QDataStream &)
func (this *QStandardItem) Read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(*qtcore.QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem4readER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "read", args)
  }

  return
}

// removeRows(int, int)
func (this *QStandardItem) Removerows(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem10removeRowsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeRows", args)
  }

  return
}

// setCheckable(_Bool)
func (this *QStandardItem) Setcheckable(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem12setCheckableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setCheckable", args)
  }

  return
}

// takeRow(int)
func (this *QStandardItem) Takerow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem7takeRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "takeRow", args)
  }

  return
}

// setChild(int, class QStandardItem *)
func (this *QStandardItem) Setchild(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem8setChildEiPS_(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN13QStandardItem8setChildEiiPS_
    // invoke: void setChild(int, int, class QStandardItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN13QStandardItem8setChildEiiPS_(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStandardItem", "setChild", args)
  }

  return
}

// type()
func (this *QStandardItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK13QStandardItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "type", args)
  }

  return
}

// setSelectable(_Bool)
func (this *QStandardItem) Setselectable(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem13setSelectableEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setSelectable", args)
  }

  return
}

// setBackground(const class QBrush &)
func (this *QStandardItem) Setbackground(args ...interface{}) () {
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
    var arg0 = args[0].(*QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem13setBackgroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setBackground", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QStandardItem) Seticon(args ...interface{}) () {
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
    var arg0 = args[0].(*QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem7setIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setIcon", args)
  }

  return
}

// foreground()
func (this *QStandardItem) Foreground(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem10foregroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "foreground", args)
  }

  return
}

// parent()
func (this *QStandardItem) Parent(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem6parentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "parent", args)
  }

  return
}

// isDropEnabled()
func (this *QStandardItem) Isdropenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem13isDropEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isDropEnabled", args)
  }

  return
}

// statusTip()
func (this *QStandardItem) Statustip(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem9statusTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "statusTip", args)
  }

  return
}

// clone()
func (this *QStandardItem) Clone(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem5cloneEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "clone", args)
  }

  return
}

// removeRow(int)
func (this *QStandardItem) Removerow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem9removeRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeRow", args)
  }

  return
}

// toolTip()
func (this *QStandardItem) Tooltip(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem7toolTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "toolTip", args)
  }

  return
}

// rowCount()
func (this *QStandardItem) Rowcount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem8rowCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "rowCount", args)
  }

  return
}

// insertRows(int, int)
func (this *QStandardItem) Insertrows(args ...interface{}) () {
  // insertRows(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10insertRowsEii
    // invoke: void insertRows(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem10insertRowsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "insertRows", args)
  }

  return
}

// setWhatsThis(const class QString &)
func (this *QStandardItem) Setwhatsthis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem12setWhatsThisERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setWhatsThis", args)
  }

  return
}

// background()
func (this *QStandardItem) Background(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem10backgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "background", args)
  }

  return
}

// child(int, int)
func (this *QStandardItem) Child(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK13QStandardItem5childEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "child", args)
  }

  return
}

// setTristate(_Bool)
func (this *QStandardItem) Settristate(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem11setTristateEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setTristate", args)
  }

  return
}

// removeColumn(int)
func (this *QStandardItem) Removecolumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem12removeColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeColumn", args)
  }

  return
}

// accessibleDescription()
func (this *QStandardItem) Accessibledescription(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem21accessibleDescriptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "accessibleDescription", args)
  }

  return
}

// data(int)
func (this *QStandardItem) Data(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK13QStandardItem4dataEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "data", args)
  }

  return
}

// appendRow(class QStandardItem *)
func (this *QStandardItem) Appendrow(args ...interface{}) () {
  // appendRow(class QStandardItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStandardItem{}) // "QStandardItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem9appendRowEPS_
    // invoke: void appendRow(class QStandardItem *)
    var arg0 = args[0].(*QStandardItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem9appendRowEPS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "appendRow", args)
  }

  return
}

// setAccessibleText(const class QString &)
func (this *QStandardItem) Setaccessibletext(args ...interface{}) () {
  // setAccessibleText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem17setAccessibleTextERK7QString
    // invoke: void setAccessibleText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem17setAccessibleTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setAccessibleText", args)
  }

  return
}

// icon()
func (this *QStandardItem) Icon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem4iconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "icon", args)
  }

  return
}

// setSizeHint(const class QSize &)
func (this *QStandardItem) Setsizehint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem11setSizeHintERK5QSize
    // invoke: void setSizeHint(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem11setSizeHintERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setSizeHint", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QStandardItem) Settooltip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem10setToolTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setToolTip", args)
  }

  return
}

// sizeHint()
func (this *QStandardItem) Sizehint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "sizeHint", args)
  }

  return
}

// hasChildren()
func (this *QStandardItem) Haschildren(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem11hasChildrenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "hasChildren", args)
  }

  return
}

// column()
func (this *QStandardItem) Column(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem6columnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "column", args)
  }

  return
}

// insertColumns(int, int)
func (this *QStandardItem) Insertcolumns(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem13insertColumnsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "insertColumns", args)
  }

  return
}

// setStatusTip(const class QString &)
func (this *QStandardItem) Setstatustip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem12setStatusTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setStatusTip", args)
  }

  return
}

// setDropEnabled(_Bool)
func (this *QStandardItem) Setdropenabled(args ...interface{}) () {
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
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem14setDropEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setDropEnabled", args)
  }

  return
}

// isDragEnabled()
func (this *QStandardItem) Isdragenabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem13isDragEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "isDragEnabled", args)
  }

  return
}

// setRowCount(int)
func (this *QStandardItem) Setrowcount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem11setRowCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setRowCount", args)
  }

  return
}

// write(class QDataStream &)
func (this *QStandardItem) Write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(*qtcore.QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK13QStandardItem5writeER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "write", args)
  }

  return
}

// flags()
func (this *QStandardItem) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK13QStandardItem5flagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "flags", args)
  }

  return
}

// checkState()
func (this *QStandardItem) Checkstate(args ...interface{}) () {
  // checkState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem10checkStateEv
    // invoke: Qt::CheckState checkState()
    C.C_ZNK13QStandardItem10checkStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "checkState", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QStandardItem) Setfont(args ...interface{}) () {
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
    var arg0 = args[0].(*QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "setFont", args)
  }

  return
}

// takeColumn(int)
func (this *QStandardItem) Takecolumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN13QStandardItem10takeColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStandardItem", "takeColumn", args)
  }

  return
}

// model()
func (this *QStandardItem) Model(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK13QStandardItem5modelEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QStandardItemModel{}) // "QStandardItemModel *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStandardItem", "model", args)
  }

  return
}

// setData(const class QVariant &, int)
func (this *QStandardItem) Setdata(args ...interface{}) () {
  // setData(const class QVariant &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QStandardItem7setDataERK8QVarianti
    // invoke: void setData(const class QVariant &, int)
    var arg0 = args[0].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem7setDataERK8QVarianti(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "setData", args)
  }

  return
}

// removeColumns(int, int)
func (this *QStandardItem) Removecolumns(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN13QStandardItem13removeColumnsEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStandardItem", "removeColumns", args)
  }

  return
}

// textAlignment()
func (this *QStandardItem) Textalignment(args ...interface{}) () {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QStandardItem13textAlignmentEv
    // invoke: Qt::Alignment textAlignment()
    C.C_ZNK13QStandardItem13textAlignmentEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStandardItem", "textAlignment", args)
  }

  return
}

// <= body block end

