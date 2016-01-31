package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtWidgets/qtreewidget.h
// dst-file: /src/widgets/qtreewidget.go
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
  // proto:  int QTreeWidget::sortColumn();
extern int32_t C_ZNK11QTreeWidget10sortColumnEv(void* qthis); // 4
  // proto:  void QTreeWidget::setHeaderLabel(const QString & label);
extern void C_ZN11QTreeWidget14setHeaderLabelERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTreeWidget::setFirstItemColumnSpanned(const QTreeWidgetItem * item, bool span);
extern void C_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  int QTreeWidget::indexOfTopLevelItem(QTreeWidgetItem * item);
extern int32_t C_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setHeaderLabels(const QStringList & labels);
extern void C_ZN11QTreeWidget15setHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setItemExpanded(const QTreeWidgetItem * item, bool expand);
extern void C_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeWidget::setItemSelected(const QTreeWidgetItem * item, bool select);
extern void C_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeWidget::QTreeWidget(QWidget * parent);
extern void* C_ZN11QTreeWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  QWidget * QTreeWidget::itemWidget(QTreeWidgetItem * item, int column);
extern void* C_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QList<QTreeWidgetItem *> QTreeWidget::selectedItems();
extern void C_ZNK11QTreeWidget13selectedItemsEv(void* qthis); // 4
  // proto:  void QTreeWidget::openPersistentEditor(QTreeWidgetItem * item, int column);
extern void C_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QRect QTreeWidget::visualItemRect(const QTreeWidgetItem * item);
extern void* C_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::removeItemWidget(QTreeWidgetItem * item, int column);
extern void C_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QTreeWidget::setColumnCount(int columns);
extern void C_ZN11QTreeWidget14setColumnCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeWidget::expandItem(const QTreeWidgetItem * item);
extern void C_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::collapseItem(const QTreeWidgetItem * item);
extern void C_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setItemWidget(QTreeWidgetItem * item, int column, QWidget * widget);
extern void C_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QTreeWidget::editItem(QTreeWidgetItem * item, int column);
extern void C_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QTreeWidget::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::currentItem();
extern void* C_ZNK11QTreeWidget11currentItemEv(void* qthis); // 4
  // proto:  bool QTreeWidget::isItemSelected(const QTreeWidgetItem * item);
extern bool C_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::itemBelow(const QTreeWidgetItem * item);
extern void* C_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTreeWidget::currentColumn();
extern int32_t C_ZNK11QTreeWidget13currentColumnEv(void* qthis); // 4
  // proto:  bool QTreeWidget::isItemHidden(const QTreeWidgetItem * item);
extern bool C_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTreeWidget::columnCount();
extern int32_t C_ZNK11QTreeWidget11columnCountEv(void* qthis); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::itemAbove(const QTreeWidgetItem * item);
extern void* C_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::invisibleRootItem();
extern void* C_ZNK11QTreeWidget17invisibleRootItemEv(void* qthis); // 4
  // proto:  bool QTreeWidget::isFirstItemColumnSpanned(const QTreeWidgetItem * item);
extern bool C_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  bool QTreeWidget::isItemExpanded(const QTreeWidgetItem * item);
extern bool C_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::takeTopLevelItem(int index);
extern void* C_ZN11QTreeWidget16takeTopLevelItemEi(void* qthis, int32_t arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::topLevelItem(int index);
extern void* C_ZNK11QTreeWidget12topLevelItemEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeWidget::topLevelItemCount();
extern int32_t C_ZNK11QTreeWidget17topLevelItemCountEv(void* qthis); // 4
  // proto:  void QTreeWidget::insertTopLevelItem(int index, QTreeWidgetItem * item);
extern void C_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTreeWidget::addTopLevelItem(QTreeWidgetItem * item);
extern void C_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::headerItem();
extern void* C_ZNK11QTreeWidget10headerItemEv(void* qthis); // 4
  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item);
extern void C_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item, int column);
extern void C_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QTreeWidget::metaObject();
extern void C_ZNK11QTreeWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QTreeWidget::~QTreeWidget();
extern void C_ZN11QTreeWidgetD2Ev(void* qthis); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(int x, int y);
extern void* C_ZNK11QTreeWidget6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(const QPoint & p);
extern void* C_ZNK11QTreeWidget6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::clear();
extern void C_ZN11QTreeWidget5clearEv(void* qthis); // 4
  // proto:  void QTreeWidget::setItemHidden(const QTreeWidgetItem * item, bool hide);
extern void C_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeWidget::closePersistentEditor(QTreeWidgetItem * item, int column);
extern void C_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QTreeWidget::setHeaderItem(QTreeWidgetItem * item);
extern void C_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setTextAlignment(int column, int alignment);
extern void C_ZN15QTreeWidgetItem16setTextAlignmentEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  int QTreeWidgetItem::columnCount();
extern int32_t C_ZNK15QTreeWidgetItem11columnCountEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::addChild(QTreeWidgetItem * child);
extern void C_ZN15QTreeWidgetItem8addChildEPS_(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setSelected(bool select);
extern void C_ZN15QTreeWidgetItem11setSelectedEb(void* qthis, bool arg0); // 2
  // proto:  QString QTreeWidgetItem::text(int column);
extern void* C_ZNK15QTreeWidgetItem4textEi(void* qthis, int32_t arg0); // 2
  // proto:  QString QTreeWidgetItem::whatsThis(int column);
extern void* C_ZNK15QTreeWidgetItem9whatsThisEi(void* qthis, int32_t arg0); // 2
  // proto:  int QTreeWidgetItem::type();
extern int32_t C_ZNK15QTreeWidgetItem4typeEv(void* qthis); // 2
  // proto:  QFont QTreeWidgetItem::font(int column);
extern void* C_ZNK15QTreeWidgetItem4fontEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::~QTreeWidgetItem();
extern void C_ZN15QTreeWidgetItemD2Ev(void* qthis); // 4
  // proto:  void QTreeWidgetItem::setForeground(int column, const QBrush & brush);
extern void C_ZN15QTreeWidgetItem13setForegroundEiRK6QBrush(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QTreeWidgetItem::isSelected();
extern bool C_ZNK15QTreeWidgetItem10isSelectedEv(void* qthis); // 2
  // proto:  bool QTreeWidgetItem::isDisabled();
extern bool C_ZNK15QTreeWidgetItem10isDisabledEv(void* qthis); // 2
  // proto:  int QTreeWidgetItem::indexOfChild(QTreeWidgetItem * child);
extern int32_t C_ZNK15QTreeWidgetItem12indexOfChildEPS_(void* qthis, void* arg0); // 2
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, int type);
extern void* C_ZN15QTreeWidgetItemC2EPS_i(void* arg0, int32_t arg1); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(int type);
extern void* C_ZN15QTreeWidgetItemC2Ei(int32_t arg0); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, QTreeWidgetItem * after, int type);
extern void* C_ZN15QTreeWidgetItemC2EPS_S0_i(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, int type);
extern void* C_ZN15QTreeWidgetItemC2EP11QTreeWidgeti(void* arg0, int32_t arg1); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, const QStringList & strings, int type);
extern void* C_ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, QTreeWidgetItem * after, int type);
extern void* C_ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(const QStringList & strings, int type);
extern void* C_ZN15QTreeWidgetItemC2ERK11QStringListi(void* arg0, int32_t arg1); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(const QTreeWidgetItem & other);
extern void* C_ZN15QTreeWidgetItemC2ERKS_(void* arg0); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, const QStringList & strings, int type);
extern void* C_ZN15QTreeWidgetItemC2EPS_RK11QStringListi(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::write(QDataStream & out);
extern void C_ZNK15QTreeWidgetItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setExpanded(bool expand);
extern void C_ZN15QTreeWidgetItem11setExpandedEb(void* qthis, bool arg0); // 2
  // proto:  void QTreeWidgetItem::insertChild(int index, QTreeWidgetItem * child);
extern void C_ZN15QTreeWidgetItem11insertChildEiPS_(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QColor QTreeWidgetItem::backgroundColor(int column);
extern void* C_ZNK15QTreeWidgetItem15backgroundColorEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setBackground(int column, const QBrush & brush);
extern void C_ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QColor QTreeWidgetItem::textColor(int column);
extern void* C_ZNK15QTreeWidgetItem9textColorEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setIcon(int column, const QIcon & icon);
extern void C_ZN15QTreeWidgetItem7setIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  int QTreeWidgetItem::childCount();
extern int32_t C_ZNK15QTreeWidgetItem10childCountEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setData(int column, int role, const QVariant & value);
extern void C_ZN15QTreeWidgetItem7setDataEiiRK8QVariant(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QTreeWidgetItem::statusTip(int column);
extern void* C_ZNK15QTreeWidgetItem9statusTipEi(void* qthis, int32_t arg0); // 2
  // proto:  QBrush QTreeWidgetItem::foreground(int column);
extern void* C_ZNK15QTreeWidgetItem10foregroundEi(void* qthis, int32_t arg0); // 2
  // proto:  QTreeWidget * QTreeWidgetItem::treeWidget();
extern void* C_ZNK15QTreeWidgetItem10treeWidgetEv(void* qthis); // 2
  // proto:  QTreeWidgetItem * QTreeWidgetItem::parent();
extern void* C_ZNK15QTreeWidgetItem6parentEv(void* qthis); // 2
  // proto:  bool QTreeWidgetItem::isFirstColumnSpanned();
extern bool C_ZNK15QTreeWidgetItem20isFirstColumnSpannedEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setFirstColumnSpanned(bool span);
extern void C_ZN15QTreeWidgetItem21setFirstColumnSpannedEb(void* qthis, bool arg0); // 2
  // proto:  QTreeWidgetItem * QTreeWidgetItem::clone();
extern void* C_ZNK15QTreeWidgetItem5cloneEv(void* qthis); // 4
  // proto:  void QTreeWidgetItem::setBackgroundColor(int column, const QColor & color);
extern void C_ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QString QTreeWidgetItem::toolTip(int column);
extern void* C_ZNK15QTreeWidgetItem7toolTipEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setFont(int column, const QFont & font);
extern void C_ZN15QTreeWidgetItem7setFontEiRK5QFont(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QTreeWidgetItem::ChildIndicatorPolicy QTreeWidgetItem::childIndicatorPolicy();
extern void C_ZNK15QTreeWidgetItem20childIndicatorPolicyEv(void* qthis); // 4
  // proto:  void QTreeWidgetItem::setWhatsThis(int column, const QString & whatsThis);
extern void C_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QBrush QTreeWidgetItem::background(int column);
extern void* C_ZNK15QTreeWidgetItem10backgroundEi(void* qthis, int32_t arg0); // 2
  // proto:  QTreeWidgetItem * QTreeWidgetItem::child(int index);
extern void* C_ZNK15QTreeWidgetItem5childEi(void* qthis, int32_t arg0); // 2
  // proto:  QVariant QTreeWidgetItem::data(int column, int role);
extern void* C_ZNK15QTreeWidgetItem4dataEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTreeWidgetItem::setTextColor(int column, const QColor & color);
extern void C_ZN15QTreeWidgetItem12setTextColorEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QIcon QTreeWidgetItem::icon(int column);
extern void* C_ZNK15QTreeWidgetItem4iconEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QTreeWidgetItem::isExpanded();
extern bool C_ZNK15QTreeWidgetItem10isExpandedEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setSizeHint(int column, const QSize & size);
extern void C_ZN15QTreeWidgetItem11setSizeHintEiRK5QSize(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QTreeWidgetItem::setToolTip(int column, const QString & toolTip);
extern void C_ZN15QTreeWidgetItem10setToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QSize QTreeWidgetItem::sizeHint(int column);
extern void* C_ZNK15QTreeWidgetItem8sizeHintEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::read(QDataStream & in);
extern void C_ZN15QTreeWidgetItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  bool QTreeWidgetItem::isHidden();
extern bool C_ZNK15QTreeWidgetItem8isHiddenEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setText(int column, const QString & text);
extern void C_ZN15QTreeWidgetItem7setTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QTreeWidgetItem::setStatusTip(int column, const QString & statusTip);
extern void C_ZN15QTreeWidgetItem12setStatusTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QTreeWidgetItem::removeChild(QTreeWidgetItem * child);
extern void C_ZN15QTreeWidgetItem11removeChildEPS_(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setDisabled(bool disabled);
extern void C_ZN15QTreeWidgetItem11setDisabledEb(void* qthis, bool arg0); // 2
  // proto:  Qt::ItemFlags QTreeWidgetItem::flags();
extern void C_ZNK15QTreeWidgetItem5flagsEv(void* qthis); // 4
  // proto:  Qt::CheckState QTreeWidgetItem::checkState(int column);
extern void C_ZNK15QTreeWidgetItem10checkStateEi(void* qthis, int32_t arg0); // 2
  // proto:  QList<QTreeWidgetItem *> QTreeWidgetItem::takeChildren();
extern void C_ZN15QTreeWidgetItem12takeChildrenEv(void* qthis); // 4
  // proto:  QTreeWidgetItem * QTreeWidgetItem::takeChild(int index);
extern void* C_ZN15QTreeWidgetItem9takeChildEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeWidgetItem::textAlignment(int column);
extern int32_t C_ZNK15QTreeWidgetItem13textAlignmentEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setHidden(bool hide);
extern void C_ZN15QTreeWidgetItem9setHiddenEb(void* qthis, bool arg0); // 2
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

// class sizeof(QTreeWidget)=1
type QTreeWidget struct {
  /*qbase*/ QTreeView;
  qclsinst unsafe.Pointer /* *C.void */;
//  _itemDoubleClicked QTreeWidget_itemDoubleClicked_signal;
//  _itemClicked QTreeWidget_itemClicked_signal;
//  _currentItemChanged QTreeWidget_currentItemChanged_signal;
//  _itemEntered QTreeWidget_itemEntered_signal;
//  _itemPressed QTreeWidget_itemPressed_signal;
//  _itemSelectionChanged QTreeWidget_itemSelectionChanged_signal;
//  _itemActivated QTreeWidget_itemActivated_signal;
//  _itemExpanded QTreeWidget_itemExpanded_signal;
//  _itemChanged QTreeWidget_itemChanged_signal;
//  _itemCollapsed QTreeWidget_itemCollapsed_signal;
}

// class sizeof(QTreeWidgetItem)=1
type QTreeWidgetItem struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// sortColumn()
func (this *QTreeWidget) Sortcolumn(args ...interface{}) (ret interface{}) {
  // sortColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10sortColumnEv
    // invoke: int sortColumn()
    var ret0 = C.C_ZNK11QTreeWidget10sortColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "sortColumn", args)
  }

  return
}

// setHeaderLabel(const class QString &)
func (this *QTreeWidget) Setheaderlabel(args ...interface{}) () {
  // setHeaderLabel(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setHeaderLabelERK7QString
    // invoke: void setHeaderLabel(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget14setHeaderLabelERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabel", args)
  }

  return
}

// setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) Setfirstitemcolumnspanned(args ...interface{}) () {
  // setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb
    // invoke: void setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setFirstItemColumnSpanned", args)
  }

  return
}

// indexOfTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) Indexoftoplevelitem(args ...interface{}) (ret interface{}) {
  // indexOfTopLevelItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem
    // invoke: int indexOfTopLevelItem(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "indexOfTopLevelItem", args)
  }

  return
}

// setHeaderLabels(const class QStringList &)
func (this *QTreeWidget) Setheaderlabels(args ...interface{}) () {
  // setHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setHeaderLabelsERK11QStringList
    // invoke: void setHeaderLabels(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget15setHeaderLabelsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabels", args)
  }

  return
}

// setItemExpanded(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) Setitemexpanded(args ...interface{}) () {
  // setItemExpanded(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb
    // invoke: void setItemExpanded(const class QTreeWidgetItem *, _Bool)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemExpanded", args)
  }

  return
}

// setItemSelected(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) Setitemselected(args ...interface{}) () {
  // setItemSelected(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb
    // invoke: void setItemSelected(const class QTreeWidgetItem *, _Bool)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemSelected", args)
  }

  return
}

// QTreeWidget(class QWidget *)
func NewQTreeWidget(args ...interface{}) *QTreeWidget {
  // QTreeWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidgetC1EP7QWidget
    // invoke: void QTreeWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTreeWidgetC2EP7QWidget(arg0)
    return &QTreeWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTreeWidget", "QTreeWidget", args)
  }

  return nil // QTreeWidget{qclsinst:qthis}
}

// itemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) Itemwidget(args ...interface{}) (ret interface{}) {
  // itemWidget(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi
    // invoke: QWidget * itemWidget(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemWidget", args)
  }

  return
}

// selectedItems()
func (this *QTreeWidget) Selecteditems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget13selectedItemsEv
    // invoke: QList<QTreeWidgetItem *> selectedItems()
    C.C_ZNK11QTreeWidget13selectedItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "selectedItems", args)
  }

  return
}

// openPersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) Openpersistenteditor(args ...interface{}) () {
  // openPersistentEditor(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi
    // invoke: void openPersistentEditor(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "openPersistentEditor", args)
  }

  return
}

// visualItemRect(const class QTreeWidgetItem *)
func (this *QTreeWidget) Visualitemrect(args ...interface{}) (ret interface{}) {
  // visualItemRect(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem
    // invoke: QRect visualItemRect(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "visualItemRect", args)
  }

  return
}

// removeItemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) Removeitemwidget(args ...interface{}) () {
  // removeItemWidget(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi
    // invoke: void removeItemWidget(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "removeItemWidget", args)
  }

  return
}

// setColumnCount(int)
func (this *QTreeWidget) Setcolumncount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setColumnCountEi
    // invoke: void setColumnCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget14setColumnCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setColumnCount", args)
  }

  return
}

// expandItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) Expanditem(args ...interface{}) () {
  // expandItem(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem
    // invoke: void expandItem(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "expandItem", args)
  }

  return
}

// collapseItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) Collapseitem(args ...interface{}) () {
  // collapseItem(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem
    // invoke: void collapseItem(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "collapseItem", args)
  }

  return
}

// setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
func (this *QTreeWidget) Setitemwidget(args ...interface{}) () {
  // setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget
    // invoke: void setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemWidget", args)
  }

  return
}

// editItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) Edititem(args ...interface{}) () {
  // editItem(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget8editItemEP15QTreeWidgetItemi
    // invoke: void editItem(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget8editItemEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "editItem", args)
  }

  return
}

// setSelectionModel(class QItemSelectionModel *)
func (this *QTreeWidget) Setselectionmodel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QItemSelectionModel{}) // "QItemSelectionModel *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(QItemSelectionModel).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setSelectionModel", args)
  }

  return
}

// currentItem()
func (this *QTreeWidget) Currentitem(args ...interface{}) (ret interface{}) {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget11currentItemEv
    // invoke: QTreeWidgetItem * currentItem()
    var ret0 = C.C_ZNK11QTreeWidget11currentItemEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentItem", args)
  }

  return
}

// isItemSelected(const class QTreeWidgetItem *)
func (this *QTreeWidget) Isitemselected(args ...interface{}) (ret interface{}) {
  // isItemSelected(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem
    // invoke: bool isItemSelected(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemSelected", args)
  }

  return
}

// itemBelow(const class QTreeWidgetItem *)
func (this *QTreeWidget) Itembelow(args ...interface{}) (ret interface{}) {
  // itemBelow(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem
    // invoke: QTreeWidgetItem * itemBelow(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemBelow", args)
  }

  return
}

// currentColumn()
func (this *QTreeWidget) Currentcolumn(args ...interface{}) (ret interface{}) {
  // currentColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget13currentColumnEv
    // invoke: int currentColumn()
    var ret0 = C.C_ZNK11QTreeWidget13currentColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentColumn", args)
  }

  return
}

// isItemHidden(const class QTreeWidgetItem *)
func (this *QTreeWidget) Isitemhidden(args ...interface{}) (ret interface{}) {
  // isItemHidden(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem
    // invoke: bool isItemHidden(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemHidden", args)
  }

  return
}

// columnCount()
func (this *QTreeWidget) Columncount(args ...interface{}) (ret interface{}) {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget11columnCountEv
    // invoke: int columnCount()
    var ret0 = C.C_ZNK11QTreeWidget11columnCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "columnCount", args)
  }

  return
}

// itemAbove(const class QTreeWidgetItem *)
func (this *QTreeWidget) Itemabove(args ...interface{}) (ret interface{}) {
  // itemAbove(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem
    // invoke: QTreeWidgetItem * itemAbove(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAbove", args)
  }

  return
}

// invisibleRootItem()
func (this *QTreeWidget) Invisiblerootitem(args ...interface{}) (ret interface{}) {
  // invisibleRootItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget17invisibleRootItemEv
    // invoke: QTreeWidgetItem * invisibleRootItem()
    var ret0 = C.C_ZNK11QTreeWidget17invisibleRootItemEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "invisibleRootItem", args)
  }

  return
}

// isFirstItemColumnSpanned(const class QTreeWidgetItem *)
func (this *QTreeWidget) Isfirstitemcolumnspanned(args ...interface{}) (ret interface{}) {
  // isFirstItemColumnSpanned(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem
    // invoke: bool isFirstItemColumnSpanned(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "isFirstItemColumnSpanned", args)
  }

  return
}

// isItemExpanded(const class QTreeWidgetItem *)
func (this *QTreeWidget) Isitemexpanded(args ...interface{}) (ret interface{}) {
  // isItemExpanded(const class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem
    // invoke: bool isItemExpanded(const class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemExpanded", args)
  }

  return
}

// takeTopLevelItem(int)
func (this *QTreeWidget) Taketoplevelitem(args ...interface{}) (ret interface{}) {
  // takeTopLevelItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget16takeTopLevelItemEi
    // invoke: QTreeWidgetItem * takeTopLevelItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QTreeWidget16takeTopLevelItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "takeTopLevelItem", args)
  }

  return
}

// topLevelItem(int)
func (this *QTreeWidget) Toplevelitem(args ...interface{}) (ret interface{}) {
  // topLevelItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget12topLevelItemEi
    // invoke: QTreeWidgetItem * topLevelItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget12topLevelItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItem", args)
  }

  return
}

// topLevelItemCount()
func (this *QTreeWidget) Toplevelitemcount(args ...interface{}) (ret interface{}) {
  // topLevelItemCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget17topLevelItemCountEv
    // invoke: int topLevelItemCount()
    var ret0 = C.C_ZNK11QTreeWidget17topLevelItemCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItemCount", args)
  }

  return
}

// insertTopLevelItem(int, class QTreeWidgetItem *)
func (this *QTreeWidget) Inserttoplevelitem(args ...interface{}) () {
  // insertTopLevelItem(int, class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem
    // invoke: void insertTopLevelItem(int, class QTreeWidgetItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "insertTopLevelItem", args)
  }

  return
}

// addTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) Addtoplevelitem(args ...interface{}) () {
  // addTopLevelItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem
    // invoke: void addTopLevelItem(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "addTopLevelItem", args)
  }

  return
}

// headerItem()
func (this *QTreeWidget) Headeritem(args ...interface{}) (ret interface{}) {
  // headerItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10headerItemEv
    // invoke: QTreeWidgetItem * headerItem()
    var ret0 = C.C_ZNK11QTreeWidget10headerItemEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "headerItem", args)
  }

  return
}

// setCurrentItem(class QTreeWidgetItem *)
func (this *QTreeWidget) Setcurrentitem(args ...interface{}) () {
  // setCurrentItem(class QTreeWidgetItem *)
  // setCurrentItem(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem
    // invoke: void setCurrentItem(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi
    // invoke: void setCurrentItem(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setCurrentItem", args)
  }

  return
}

// metaObject()
func (this *QTreeWidget) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QTreeWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "metaObject", args)
  }

  return
}

// ~QTreeWidget()
func (this *QTreeWidget) Freeqtreewidget(args ...interface{}) () {
  // ~QTreeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidgetD0Ev
    // invoke: void ~QTreeWidget()
    C.C_ZN11QTreeWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "~QTreeWidget", args)
  }

  return
}

// itemAt(int, int)
func (this *QTreeWidget) Itemat(args ...interface{}) (ret interface{}) {
  // itemAt(int, int)
  // itemAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTreeWidget6itemAtEii
    // invoke: QTreeWidgetItem * itemAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QTreeWidget6itemAtEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZNK11QTreeWidget6itemAtERK6QPoint
    // invoke: QTreeWidgetItem * itemAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QTreeWidget6itemAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAt", args)
  }

  return
}

// clear()
func (this *QTreeWidget) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget5clearEv
    // invoke: void clear()
    C.C_ZN11QTreeWidget5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "clear", args)
  }

  return
}

// setItemHidden(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) Setitemhidden(args ...interface{}) () {
  // setItemHidden(const class QTreeWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb
    // invoke: void setItemHidden(const class QTreeWidgetItem *, _Bool)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemHidden", args)
  }

  return
}

// closePersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) Closepersistenteditor(args ...interface{}) () {
  // closePersistentEditor(class QTreeWidgetItem *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi
    // invoke: void closePersistentEditor(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "closePersistentEditor", args)
  }

  return
}

// setHeaderItem(class QTreeWidgetItem *)
func (this *QTreeWidget) Setheaderitem(args ...interface{}) () {
  // setHeaderItem(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem
    // invoke: void setHeaderItem(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderItem", args)
  }

  return
}

// setTextAlignment(int, int)
func (this *QTreeWidgetItem) Settextalignment(args ...interface{}) () {
  // setTextAlignment(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem16setTextAlignmentEii
    // invoke: void setTextAlignment(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem16setTextAlignmentEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextAlignment", args)
  }

  return
}

// columnCount()
func (this *QTreeWidgetItem) Columncount(args ...interface{}) (ret interface{}) {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem11columnCountEv
    // invoke: int columnCount()
    var ret0 = C.C_ZNK15QTreeWidgetItem11columnCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "columnCount", args)
  }

  return
}

// addChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) Addchild(args ...interface{}) () {
  // addChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem8addChildEPS_
    // invoke: void addChild(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem8addChildEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "addChild", args)
  }

  return
}

// setSelected(_Bool)
func (this *QTreeWidgetItem) Setselected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setSelectedEb
    // invoke: void setSelected(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem11setSelectedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSelected", args)
  }

  return
}

// text(int)
func (this *QTreeWidgetItem) Text(args ...interface{}) (ret interface{}) {
  // text(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4textEi
    // invoke: QString text(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem4textEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "text", args)
  }

  return
}

// whatsThis(int)
func (this *QTreeWidgetItem) Whatsthis(args ...interface{}) (ret interface{}) {
  // whatsThis(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9whatsThisEi
    // invoke: QString whatsThis(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem9whatsThisEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "whatsThis", args)
  }

  return
}

// type()
func (this *QTreeWidgetItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK15QTreeWidgetItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "type", args)
  }

  return
}

// font(int)
func (this *QTreeWidgetItem) Font(args ...interface{}) (ret interface{}) {
  // font(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4fontEi
    // invoke: QFont font(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem4fontEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QFont{}) // "QFont"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "font", args)
  }

  return
}

// ~QTreeWidgetItem()
func (this *QTreeWidgetItem) Freeqtreewidgetitem(args ...interface{}) () {
  // ~QTreeWidgetItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItemD0Ev
    // invoke: void ~QTreeWidgetItem()
    C.C_ZN15QTreeWidgetItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "~QTreeWidgetItem", args)
  }

  return
}

// setForeground(int, const class QBrush &)
func (this *QTreeWidgetItem) Setforeground(args ...interface{}) () {
  // setForeground(int, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem13setForegroundEiRK6QBrush
    // invoke: void setForeground(int, const class QBrush &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QBrush).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem13setForegroundEiRK6QBrush(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setForeground", args)
  }

  return
}

// isSelected()
func (this *QTreeWidgetItem) Isselected(args ...interface{}) (ret interface{}) {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isSelectedEv
    // invoke: bool isSelected()
    var ret0 = C.C_ZNK15QTreeWidgetItem10isSelectedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isSelected", args)
  }

  return
}

// isDisabled()
func (this *QTreeWidgetItem) Isdisabled(args ...interface{}) (ret interface{}) {
  // isDisabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isDisabledEv
    // invoke: bool isDisabled()
    var ret0 = C.C_ZNK15QTreeWidgetItem10isDisabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isDisabled", args)
  }

  return
}

// indexOfChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) Indexofchild(args ...interface{}) (ret interface{}) {
  // indexOfChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem12indexOfChildEPS_
    // invoke: int indexOfChild(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem12indexOfChildEPS_(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "indexOfChild", args)
  }

  return
}

// QTreeWidgetItem(class QTreeWidgetItem *, int)
func NewQTreeWidgetItem(args ...interface{}) *QTreeWidgetItem {
  // QTreeWidgetItem(class QTreeWidgetItem *, int)
  // QTreeWidgetItem(int)
  // QTreeWidgetItem(class QTreeWidgetItem *, class QTreeWidgetItem *, int)
  // QTreeWidgetItem(class QTreeWidget *, int)
  // QTreeWidgetItem(class QTreeWidget *, const class QStringList &, int)
  // QTreeWidgetItem(class QTreeWidget *, class QTreeWidgetItem *, int)
  // QTreeWidgetItem(const class QStringList &, int)
  // QTreeWidgetItem(const class QTreeWidgetItem &)
  // QTreeWidgetItem(class QTreeWidgetItem *, const class QStringList &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[2][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QTreeWidget{}) // "QTreeWidget *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QTreeWidget{}) // "QTreeWidget *"
  vtys[4][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QTreeWidget{}) // "QTreeWidget *"
  vtys[5][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QTreeWidgetItem{}) // "const QTreeWidgetItem &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
  vtys[8][1] = reflect.TypeOf(QStringList{}) // "const QStringList &"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItemC1EPS_i
    // invoke: void QTreeWidgetItem(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2EPS_i(arg0, arg1)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN15QTreeWidgetItemC1Ei
    // invoke: void QTreeWidgetItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2Ei(arg0)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 2:
    // invoke: _ZN15QTreeWidgetItemC1EPS_S0_i
    // invoke: void QTreeWidgetItem(class QTreeWidgetItem *, class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2EPS_S0_i(arg0, arg1, arg2)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 3:
    // invoke: _ZN15QTreeWidgetItemC1EP11QTreeWidgeti
    // invoke: void QTreeWidgetItem(class QTreeWidget *, int)
    var arg0 = args[0].(QTreeWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2EP11QTreeWidgeti(arg0, arg1)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 4:
    // invoke: _ZN15QTreeWidgetItemC1EP11QTreeWidgetRK11QStringListi
    // invoke: void QTreeWidgetItem(class QTreeWidget *, const class QStringList &, int)
    var arg0 = args[0].(QTreeWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi(arg0, arg1, arg2)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 5:
    // invoke: _ZN15QTreeWidgetItemC1EP11QTreeWidgetPS_i
    // invoke: void QTreeWidgetItem(class QTreeWidget *, class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i(arg0, arg1, arg2)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 6:
    // invoke: _ZN15QTreeWidgetItemC1ERK11QStringListi
    // invoke: void QTreeWidgetItem(const class QStringList &, int)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2ERK11QStringListi(arg0, arg1)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 7:
    // invoke: _ZN15QTreeWidgetItemC1ERKS_
    // invoke: void QTreeWidgetItem(const class QTreeWidgetItem &)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2ERKS_(arg0)
    return &QTreeWidgetItem{qclsinst:qthis}
  case 8:
    // invoke: _ZN15QTreeWidgetItemC1EPS_RK11QStringListi
    // invoke: void QTreeWidgetItem(class QTreeWidgetItem *, const class QStringList &, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QTreeWidgetItemC2EPS_RK11QStringListi(arg0, arg1, arg2)
    return &QTreeWidgetItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "QTreeWidgetItem", args)
  }

  return nil // QTreeWidgetItem{qclsinst:qthis}
}

// write(class QDataStream &)
func (this *QTreeWidgetItem) Write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK15QTreeWidgetItem5writeER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "write", args)
  }

  return
}

// setExpanded(_Bool)
func (this *QTreeWidgetItem) Setexpanded(args ...interface{}) () {
  // setExpanded(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setExpandedEb
    // invoke: void setExpanded(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem11setExpandedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setExpanded", args)
  }

  return
}

// insertChild(int, class QTreeWidgetItem *)
func (this *QTreeWidgetItem) Insertchild(args ...interface{}) () {
  // insertChild(int, class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11insertChildEiPS_
    // invoke: void insertChild(int, class QTreeWidgetItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem11insertChildEiPS_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "insertChild", args)
  }

  return
}

// backgroundColor(int)
func (this *QTreeWidgetItem) Backgroundcolor(args ...interface{}) (ret interface{}) {
  // backgroundColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem15backgroundColorEi
    // invoke: QColor backgroundColor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem15backgroundColorEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "backgroundColor", args)
  }

  return
}

// setBackground(int, const class QBrush &)
func (this *QTreeWidgetItem) Setbackground(args ...interface{}) () {
  // setBackground(int, const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush
    // invoke: void setBackground(int, const class QBrush &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QBrush).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackground", args)
  }

  return
}

// textColor(int)
func (this *QTreeWidgetItem) Textcolor(args ...interface{}) (ret interface{}) {
  // textColor(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9textColorEi
    // invoke: QColor textColor(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem9textColorEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textColor", args)
  }

  return
}

// setIcon(int, const class QIcon &)
func (this *QTreeWidgetItem) Seticon(args ...interface{}) () {
  // setIcon(int, const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setIconEiRK5QIcon
    // invoke: void setIcon(int, const class QIcon &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QIcon).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem7setIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setIcon", args)
  }

  return
}

// childCount()
func (this *QTreeWidgetItem) Childcount(args ...interface{}) (ret interface{}) {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10childCountEv
    // invoke: int childCount()
    var ret0 = C.C_ZNK15QTreeWidgetItem10childCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "childCount", args)
  }

  return
}

// setData(int, int, const class QVariant &)
func (this *QTreeWidgetItem) Setdata(args ...interface{}) () {
  // setData(int, int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setDataEiiRK8QVariant
    // invoke: void setData(int, int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QVariant).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN15QTreeWidgetItem7setDataEiiRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setData", args)
  }

  return
}

// statusTip(int)
func (this *QTreeWidgetItem) Statustip(args ...interface{}) (ret interface{}) {
  // statusTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem9statusTipEi
    // invoke: QString statusTip(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem9statusTipEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "statusTip", args)
  }

  return
}

// foreground(int)
func (this *QTreeWidgetItem) Foreground(args ...interface{}) (ret interface{}) {
  // foreground(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10foregroundEi
    // invoke: QBrush foreground(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem10foregroundEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "foreground", args)
  }

  return
}

// treeWidget()
func (this *QTreeWidgetItem) Treewidget(args ...interface{}) (ret interface{}) {
  // treeWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10treeWidgetEv
    // invoke: QTreeWidget * treeWidget()
    var ret0 = C.C_ZNK15QTreeWidgetItem10treeWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidget{}) // "QTreeWidget *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "treeWidget", args)
  }

  return
}

// parent()
func (this *QTreeWidgetItem) Parent(args ...interface{}) (ret interface{}) {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem6parentEv
    // invoke: QTreeWidgetItem * parent()
    var ret0 = C.C_ZNK15QTreeWidgetItem6parentEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "parent", args)
  }

  return
}

// isFirstColumnSpanned()
func (this *QTreeWidgetItem) Isfirstcolumnspanned(args ...interface{}) (ret interface{}) {
  // isFirstColumnSpanned()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem20isFirstColumnSpannedEv
    // invoke: bool isFirstColumnSpanned()
    var ret0 = C.C_ZNK15QTreeWidgetItem20isFirstColumnSpannedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isFirstColumnSpanned", args)
  }

  return
}

// setFirstColumnSpanned(_Bool)
func (this *QTreeWidgetItem) Setfirstcolumnspanned(args ...interface{}) () {
  // setFirstColumnSpanned(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem21setFirstColumnSpannedEb
    // invoke: void setFirstColumnSpanned(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem21setFirstColumnSpannedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFirstColumnSpanned", args)
  }

  return
}

// clone()
func (this *QTreeWidgetItem) Clone(args ...interface{}) (ret interface{}) {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5cloneEv
    // invoke: QTreeWidgetItem * clone()
    var ret0 = C.C_ZNK15QTreeWidgetItem5cloneEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "clone", args)
  }

  return
}

// setBackgroundColor(int, const class QColor &)
func (this *QTreeWidgetItem) Setbackgroundcolor(args ...interface{}) () {
  // setBackgroundColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor
    // invoke: void setBackgroundColor(int, const class QColor &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackgroundColor", args)
  }

  return
}

// toolTip(int)
func (this *QTreeWidgetItem) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem7toolTipEi
    // invoke: QString toolTip(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem7toolTipEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "toolTip", args)
  }

  return
}

// setFont(int, const class QFont &)
func (this *QTreeWidgetItem) Setfont(args ...interface{}) () {
  // setFont(int, const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setFontEiRK5QFont
    // invoke: void setFont(int, const class QFont &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QFont).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem7setFontEiRK5QFont(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFont", args)
  }

  return
}

// childIndicatorPolicy()
func (this *QTreeWidgetItem) Childindicatorpolicy(args ...interface{}) () {
  // childIndicatorPolicy()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem20childIndicatorPolicyEv
    // invoke: QTreeWidgetItem::ChildIndicatorPolicy childIndicatorPolicy()
    C.C_ZNK15QTreeWidgetItem20childIndicatorPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "childIndicatorPolicy", args)
  }

  return
}

// setWhatsThis(int, const class QString &)
func (this *QTreeWidgetItem) Setwhatsthis(args ...interface{}) () {
  // setWhatsThis(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setWhatsThisEiRK7QString
    // invoke: void setWhatsThis(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem12setWhatsThisEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setWhatsThis", args)
  }

  return
}

// background(int)
func (this *QTreeWidgetItem) Background(args ...interface{}) (ret interface{}) {
  // background(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10backgroundEi
    // invoke: QBrush background(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem10backgroundEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QBrush{}) // "QBrush"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "background", args)
  }

  return
}

// child(int)
func (this *QTreeWidgetItem) Child(args ...interface{}) (ret interface{}) {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5childEi
    // invoke: QTreeWidgetItem * child(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem5childEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "child", args)
  }

  return
}

// data(int, int)
func (this *QTreeWidgetItem) Data(args ...interface{}) (ret interface{}) {
  // data(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4dataEii
    // invoke: QVariant data(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK15QTreeWidgetItem4dataEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "data", args)
  }

  return
}

// setTextColor(int, const class QColor &)
func (this *QTreeWidgetItem) Settextcolor(args ...interface{}) () {
  // setTextColor(int, const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setTextColorEiRK6QColor
    // invoke: void setTextColor(int, const class QColor &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QColor).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem12setTextColorEiRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextColor", args)
  }

  return
}

// icon(int)
func (this *QTreeWidgetItem) Icon(args ...interface{}) (ret interface{}) {
  // icon(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem4iconEi
    // invoke: QIcon icon(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem4iconEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIcon{}) // "QIcon"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "icon", args)
  }

  return
}

// isExpanded()
func (this *QTreeWidgetItem) Isexpanded(args ...interface{}) (ret interface{}) {
  // isExpanded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10isExpandedEv
    // invoke: bool isExpanded()
    var ret0 = C.C_ZNK15QTreeWidgetItem10isExpandedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isExpanded", args)
  }

  return
}

// setSizeHint(int, const class QSize &)
func (this *QTreeWidgetItem) Setsizehint(args ...interface{}) () {
  // setSizeHint(int, const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setSizeHintEiRK5QSize
    // invoke: void setSizeHint(int, const class QSize &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QSize).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem11setSizeHintEiRK5QSize(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSizeHint", args)
  }

  return
}

// setToolTip(int, const class QString &)
func (this *QTreeWidgetItem) Settooltip(args ...interface{}) () {
  // setToolTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem10setToolTipEiRK7QString
    // invoke: void setToolTip(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem10setToolTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setToolTip", args)
  }

  return
}

// sizeHint(int)
func (this *QTreeWidgetItem) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem8sizeHintEi
    // invoke: QSize sizeHint(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem8sizeHintEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QSize{}) // "QSize"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "sizeHint", args)
  }

  return
}

// read(class QDataStream &)
func (this *QTreeWidgetItem) Read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem4readER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "read", args)
  }

  return
}

// isHidden()
func (this *QTreeWidgetItem) Ishidden(args ...interface{}) (ret interface{}) {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem8isHiddenEv
    // invoke: bool isHidden()
    var ret0 = C.C_ZNK15QTreeWidgetItem8isHiddenEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isHidden", args)
  }

  return
}

// setText(int, const class QString &)
func (this *QTreeWidgetItem) Settext(args ...interface{}) () {
  // setText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem7setTextEiRK7QString
    // invoke: void setText(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem7setTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setText", args)
  }

  return
}

// setStatusTip(int, const class QString &)
func (this *QTreeWidgetItem) Setstatustip(args ...interface{}) () {
  // setStatusTip(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12setStatusTipEiRK7QString
    // invoke: void setStatusTip(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QTreeWidgetItem12setStatusTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setStatusTip", args)
  }

  return
}

// removeChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) Removechild(args ...interface{}) () {
  // removeChild(class QTreeWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11removeChildEPS_
    // invoke: void removeChild(class QTreeWidgetItem *)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem11removeChildEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "removeChild", args)
  }

  return
}

// setDisabled(_Bool)
func (this *QTreeWidgetItem) Setdisabled(args ...interface{}) () {
  // setDisabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem11setDisabledEb
    // invoke: void setDisabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem11setDisabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setDisabled", args)
  }

  return
}

// flags()
func (this *QTreeWidgetItem) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK15QTreeWidgetItem5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "flags", args)
  }

  return
}

// checkState(int)
func (this *QTreeWidgetItem) Checkstate(args ...interface{}) () {
  // checkState(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem10checkStateEi
    // invoke: Qt::CheckState checkState(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK15QTreeWidgetItem10checkStateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "checkState", args)
  }

  return
}

// takeChildren()
func (this *QTreeWidgetItem) Takechildren(args ...interface{}) () {
  // takeChildren()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem12takeChildrenEv
    // invoke: QList<QTreeWidgetItem *> takeChildren()
    C.C_ZN15QTreeWidgetItem12takeChildrenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChildren", args)
  }

  return
}

// takeChild(int)
func (this *QTreeWidgetItem) Takechild(args ...interface{}) (ret interface{}) {
  // takeChild(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem9takeChildEi
    // invoke: QTreeWidgetItem * takeChild(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN15QTreeWidgetItem9takeChildEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTreeWidgetItem{}) // "QTreeWidgetItem *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChild", args)
  }

  return
}

// textAlignment(int)
func (this *QTreeWidgetItem) Textalignment(args ...interface{}) (ret interface{}) {
  // textAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QTreeWidgetItem13textAlignmentEi
    // invoke: int textAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QTreeWidgetItem13textAlignmentEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textAlignment", args)
  }

  return
}

// setHidden(_Bool)
func (this *QTreeWidgetItem) Sethidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QTreeWidgetItem9setHiddenEb
    // invoke: void setHidden(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QTreeWidgetItem9setHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setHidden", args)
  }

  return
}

// <= body block end

