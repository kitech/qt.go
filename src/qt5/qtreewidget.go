package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK11QTreeWidget10sortColumnEv(void* qthis); // 4
  // proto:  void QTreeWidget::setHeaderLabel(const QString & label);
extern void _ZN11QTreeWidget14setHeaderLabelERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTreeWidget::setFirstItemColumnSpanned(const QTreeWidgetItem * item, bool span);
extern void _ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  int QTreeWidget::indexOfTopLevelItem(QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setHeaderLabels(const QStringList & labels);
extern void _ZN11QTreeWidget15setHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setItemExpanded(const QTreeWidgetItem * item, bool expand);
extern void _ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeWidget::setItemSelected(const QTreeWidgetItem * item, bool select);
extern void _ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeWidget::QTreeWidget(QWidget * parent);
extern void _ZN11QTreeWidgetC2EP7QWidget(void* qthis, void* arg0); // 3
  // proto:  QWidget * QTreeWidget::itemWidget(QTreeWidgetItem * item, int column);
extern void _ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QList<QTreeWidgetItem *> QTreeWidget::selectedItems();
extern void _ZNK11QTreeWidget13selectedItemsEv(void* qthis); // 4
  // proto:  void QTreeWidget::openPersistentEditor(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QRect QTreeWidget::visualItemRect(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::removeItemWidget(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  void QTreeWidget::setColumnCount(int columns);
extern void _ZN11QTreeWidget14setColumnCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTreeWidget::expandItem(const QTreeWidgetItem * item);
extern void _ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::collapseItem(const QTreeWidgetItem * item);
extern void _ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setItemWidget(QTreeWidgetItem * item, int column, QWidget * widget);
extern void _ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QTreeWidget::editItem(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget8editItemEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QTreeWidget::setSelectionModel(QItemSelectionModel * selectionModel);
extern void _ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::currentItem();
extern void _ZNK11QTreeWidget11currentItemEv(void* qthis); // 4
  // proto:  bool QTreeWidget::isItemSelected(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::itemBelow(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTreeWidget::currentColumn();
extern void _ZNK11QTreeWidget13currentColumnEv(void* qthis); // 4
  // proto:  bool QTreeWidget::isItemHidden(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTreeWidget::columnCount();
extern void _ZNK11QTreeWidget11columnCountEv(void* qthis); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::itemAbove(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::invisibleRootItem();
extern void _ZNK11QTreeWidget17invisibleRootItemEv(void* qthis); // 4
  // proto:  bool QTreeWidget::isFirstItemColumnSpanned(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  bool QTreeWidget::isItemExpanded(const QTreeWidgetItem * item);
extern void _ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::takeTopLevelItem(int index);
extern void _ZN11QTreeWidget16takeTopLevelItemEi(void* qthis, int32_t arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::topLevelItem(int index);
extern void _ZNK11QTreeWidget12topLevelItemEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeWidget::topLevelItemCount();
extern void _ZNK11QTreeWidget17topLevelItemCountEv(void* qthis); // 4
  // proto:  void QTreeWidget::insertTopLevelItem(int index, QTreeWidgetItem * item);
extern void _ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTreeWidget::addTopLevelItem(QTreeWidgetItem * item);
extern void _ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::headerItem();
extern void _ZNK11QTreeWidget10headerItemEv(void* qthis); // 4
  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item);
extern void _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::setCurrentItem(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  const QMetaObject * QTreeWidget::metaObject();
extern void _ZNK11QTreeWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QTreeWidget::~QTreeWidget();
extern void _ZN11QTreeWidgetD2Ev(void* qthis); // 4
  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(int x, int y);
extern void _ZNK11QTreeWidget6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QTreeWidgetItem * QTreeWidget::itemAt(const QPoint & p);
extern void _ZNK11QTreeWidget6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidget::clear();
extern void _ZN11QTreeWidget5clearEv(void* qthis); // 4
  // proto:  void QTreeWidget::setItemHidden(const QTreeWidgetItem * item, bool hide);
extern void _ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTreeWidget::closePersistentEditor(QTreeWidgetItem * item, int column);
extern void _ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QTreeWidget::setHeaderItem(QTreeWidgetItem * item);
extern void _ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setTextAlignment(int column, int alignment);
extern void _ZN15QTreeWidgetItem16setTextAlignmentEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  int QTreeWidgetItem::columnCount();
extern void _ZNK15QTreeWidgetItem11columnCountEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::addChild(QTreeWidgetItem * child);
extern void _ZN15QTreeWidgetItem8addChildEPS_(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setSelected(bool select);
extern void _ZN15QTreeWidgetItem11setSelectedEb(void* qthis, bool arg0); // 2
  // proto:  QString QTreeWidgetItem::text(int column);
extern void _ZNK15QTreeWidgetItem4textEi(void* qthis, int32_t arg0); // 2
  // proto:  QString QTreeWidgetItem::whatsThis(int column);
extern void _ZNK15QTreeWidgetItem9whatsThisEi(void* qthis, int32_t arg0); // 2
  // proto:  int QTreeWidgetItem::type();
extern void _ZNK15QTreeWidgetItem4typeEv(void* qthis); // 2
  // proto:  QFont QTreeWidgetItem::font(int column);
extern void _ZNK15QTreeWidgetItem4fontEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::~QTreeWidgetItem();
extern void _ZN15QTreeWidgetItemD2Ev(void* qthis); // 4
  // proto:  void QTreeWidgetItem::setForeground(int column, const QBrush & brush);
extern void _ZN15QTreeWidgetItem13setForegroundEiRK6QBrush(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  bool QTreeWidgetItem::isSelected();
extern void _ZNK15QTreeWidgetItem10isSelectedEv(void* qthis); // 2
  // proto:  bool QTreeWidgetItem::isDisabled();
extern void _ZNK15QTreeWidgetItem10isDisabledEv(void* qthis); // 2
  // proto:  int QTreeWidgetItem::indexOfChild(QTreeWidgetItem * child);
extern void _ZNK15QTreeWidgetItem12indexOfChildEPS_(void* qthis, void* arg0); // 2
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, int type);
extern void _ZN15QTreeWidgetItemC2EPS_i(void* qthis, void* arg0, int32_t arg1); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(int type);
extern void _ZN15QTreeWidgetItemC2Ei(void* qthis, int32_t arg0); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, QTreeWidgetItem * after, int type);
extern void _ZN15QTreeWidgetItemC2EPS_S0_i(void* qthis, void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, int type);
extern void _ZN15QTreeWidgetItemC2EP11QTreeWidgeti(void* qthis, void* arg0, int32_t arg1); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, const QStringList & strings, int type);
extern void _ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi(void* qthis, void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidget * view, QTreeWidgetItem * after, int type);
extern void _ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i(void* qthis, void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(const QStringList & strings, int type);
extern void _ZN15QTreeWidgetItemC2ERK11QStringListi(void* qthis, void* arg0, int32_t arg1); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(const QTreeWidgetItem & other);
extern void _ZN15QTreeWidgetItemC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  void QTreeWidgetItem::QTreeWidgetItem(QTreeWidgetItem * parent, const QStringList & strings, int type);
extern void _ZN15QTreeWidgetItemC2EPS_RK11QStringListi(void* qthis, void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTreeWidgetItem::write(QDataStream & out);
extern void _ZNK15QTreeWidgetItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setExpanded(bool expand);
extern void _ZN15QTreeWidgetItem11setExpandedEb(void* qthis, bool arg0); // 2
  // proto:  void QTreeWidgetItem::insertChild(int index, QTreeWidgetItem * child);
extern void _ZN15QTreeWidgetItem11insertChildEiPS_(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QColor QTreeWidgetItem::backgroundColor(int column);
extern void _ZNK15QTreeWidgetItem15backgroundColorEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setBackground(int column, const QBrush & brush);
extern void _ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QColor QTreeWidgetItem::textColor(int column);
extern void _ZNK15QTreeWidgetItem9textColorEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setIcon(int column, const QIcon & icon);
extern void _ZN15QTreeWidgetItem7setIconEiRK5QIcon(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  int QTreeWidgetItem::childCount();
extern void _ZNK15QTreeWidgetItem10childCountEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setData(int column, int role, const QVariant & value);
extern void _ZN15QTreeWidgetItem7setDataEiiRK8QVariant(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QTreeWidgetItem::statusTip(int column);
extern void _ZNK15QTreeWidgetItem9statusTipEi(void* qthis, int32_t arg0); // 2
  // proto:  QBrush QTreeWidgetItem::foreground(int column);
extern void _ZNK15QTreeWidgetItem10foregroundEi(void* qthis, int32_t arg0); // 2
  // proto:  QTreeWidget * QTreeWidgetItem::treeWidget();
extern void _ZNK15QTreeWidgetItem10treeWidgetEv(void* qthis); // 2
  // proto:  QTreeWidgetItem * QTreeWidgetItem::parent();
extern void _ZNK15QTreeWidgetItem6parentEv(void* qthis); // 2
  // proto:  bool QTreeWidgetItem::isFirstColumnSpanned();
extern void _ZNK15QTreeWidgetItem20isFirstColumnSpannedEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setFirstColumnSpanned(bool span);
extern void _ZN15QTreeWidgetItem21setFirstColumnSpannedEb(void* qthis, bool arg0); // 2
  // proto:  QTreeWidgetItem * QTreeWidgetItem::clone();
extern void _ZNK15QTreeWidgetItem5cloneEv(void* qthis); // 4
  // proto:  void QTreeWidgetItem::setBackgroundColor(int column, const QColor & color);
extern void _ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QString QTreeWidgetItem::toolTip(int column);
extern void _ZNK15QTreeWidgetItem7toolTipEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setFont(int column, const QFont & font);
extern void _ZN15QTreeWidgetItem7setFontEiRK5QFont(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QTreeWidgetItem::ChildIndicatorPolicy QTreeWidgetItem::childIndicatorPolicy();
extern void _ZNK15QTreeWidgetItem20childIndicatorPolicyEv(void* qthis); // 4
  // proto:  void QTreeWidgetItem::setWhatsThis(int column, const QString & whatsThis);
extern void _ZN15QTreeWidgetItem12setWhatsThisEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QBrush QTreeWidgetItem::background(int column);
extern void _ZNK15QTreeWidgetItem10backgroundEi(void* qthis, int32_t arg0); // 2
  // proto:  QTreeWidgetItem * QTreeWidgetItem::child(int index);
extern void _ZNK15QTreeWidgetItem5childEi(void* qthis, int32_t arg0); // 2
  // proto:  QVariant QTreeWidgetItem::data(int column, int role);
extern void _ZNK15QTreeWidgetItem4dataEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTreeWidgetItem::setTextColor(int column, const QColor & color);
extern void _ZN15QTreeWidgetItem12setTextColorEiRK6QColor(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QIcon QTreeWidgetItem::icon(int column);
extern void _ZNK15QTreeWidgetItem4iconEi(void* qthis, int32_t arg0); // 2
  // proto:  bool QTreeWidgetItem::isExpanded();
extern void _ZNK15QTreeWidgetItem10isExpandedEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setSizeHint(int column, const QSize & size);
extern void _ZN15QTreeWidgetItem11setSizeHintEiRK5QSize(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QTreeWidgetItem::setToolTip(int column, const QString & toolTip);
extern void _ZN15QTreeWidgetItem10setToolTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QSize QTreeWidgetItem::sizeHint(int column);
extern void _ZNK15QTreeWidgetItem8sizeHintEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::read(QDataStream & in);
extern void _ZN15QTreeWidgetItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  bool QTreeWidgetItem::isHidden();
extern void _ZNK15QTreeWidgetItem8isHiddenEv(void* qthis); // 2
  // proto:  void QTreeWidgetItem::setText(int column, const QString & text);
extern void _ZN15QTreeWidgetItem7setTextEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QTreeWidgetItem::setStatusTip(int column, const QString & statusTip);
extern void _ZN15QTreeWidgetItem12setStatusTipEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  void QTreeWidgetItem::removeChild(QTreeWidgetItem * child);
extern void _ZN15QTreeWidgetItem11removeChildEPS_(void* qthis, void* arg0); // 4
  // proto:  void QTreeWidgetItem::setDisabled(bool disabled);
extern void _ZN15QTreeWidgetItem11setDisabledEb(void* qthis, bool arg0); // 2
  // proto:  Qt::ItemFlags QTreeWidgetItem::flags();
extern void _ZNK15QTreeWidgetItem5flagsEv(void* qthis); // 4
  // proto:  Qt::CheckState QTreeWidgetItem::checkState(int column);
extern void _ZNK15QTreeWidgetItem10checkStateEi(void* qthis, int32_t arg0); // 2
  // proto:  QList<QTreeWidgetItem *> QTreeWidgetItem::takeChildren();
extern void _ZN15QTreeWidgetItem12takeChildrenEv(void* qthis); // 4
  // proto:  QTreeWidgetItem * QTreeWidgetItem::takeChild(int index);
extern void _ZN15QTreeWidgetItem9takeChildEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTreeWidgetItem::textAlignment(int column);
extern void _ZNK15QTreeWidgetItem13textAlignmentEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTreeWidgetItem::setHidden(bool hide);
extern void _ZN15QTreeWidgetItem9setHiddenEb(void* qthis, bool arg0); // 2
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
func (this *QTreeWidget) sortColumn(args ...interface{}) () {
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
    C._ZNK11QTreeWidget10sortColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "sortColumn", args)
  }

}

// setHeaderLabel(const class QString &)
func (this *QTreeWidget) setHeaderLabel(args ...interface{}) () {
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
    C._ZN11QTreeWidget14setHeaderLabelERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabel", args)
  }

}

// setFirstItemColumnSpanned(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) setFirstItemColumnSpanned(args ...interface{}) () {
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
    C._ZN11QTreeWidget25setFirstItemColumnSpannedEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setFirstItemColumnSpanned", args)
  }

}

// indexOfTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) indexOfTopLevelItem(args ...interface{}) () {
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
    C._ZNK11QTreeWidget19indexOfTopLevelItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "indexOfTopLevelItem", args)
  }

}

// setHeaderLabels(const class QStringList &)
func (this *QTreeWidget) setHeaderLabels(args ...interface{}) () {
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
    C._ZN11QTreeWidget15setHeaderLabelsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderLabels", args)
  }

}

// setItemExpanded(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) setItemExpanded(args ...interface{}) () {
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
    C._ZN11QTreeWidget15setItemExpandedEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemExpanded", args)
  }

}

// setItemSelected(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) setItemSelected(args ...interface{}) () {
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
    C._ZN11QTreeWidget15setItemSelectedEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemSelected", args)
  }

}

// QTreeWidget(class QWidget *)
func NewQTreeWidget(args ...interface{}) QTreeWidget {
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
    C._ZN11QTreeWidgetC2EP7QWidget(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "QTreeWidget", args)
  }

  return QTreeWidget{}
}

// itemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) itemWidget(args ...interface{}) () {
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
    C._ZNK11QTreeWidget10itemWidgetEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemWidget", args)
  }

}

// selectedItems()
func (this *QTreeWidget) selectedItems(args ...interface{}) () {
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
    C._ZNK11QTreeWidget13selectedItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "selectedItems", args)
  }

}

// openPersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) openPersistentEditor(args ...interface{}) () {
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
    C._ZN11QTreeWidget20openPersistentEditorEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "openPersistentEditor", args)
  }

}

// visualItemRect(const class QTreeWidgetItem *)
func (this *QTreeWidget) visualItemRect(args ...interface{}) () {
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
    C._ZNK11QTreeWidget14visualItemRectEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "visualItemRect", args)
  }

}

// removeItemWidget(class QTreeWidgetItem *, int)
func (this *QTreeWidget) removeItemWidget(args ...interface{}) () {
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
    C._ZN11QTreeWidget16removeItemWidgetEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "removeItemWidget", args)
  }

}

// setColumnCount(int)
func (this *QTreeWidget) setColumnCount(args ...interface{}) () {
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
    C._ZN11QTreeWidget14setColumnCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setColumnCount", args)
  }

}

// expandItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) expandItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget10expandItemEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "expandItem", args)
  }

}

// collapseItem(const class QTreeWidgetItem *)
func (this *QTreeWidget) collapseItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget12collapseItemEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "collapseItem", args)
  }

}

// setItemWidget(class QTreeWidgetItem *, int, class QWidget *)
func (this *QTreeWidget) setItemWidget(args ...interface{}) () {
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
    C._ZN11QTreeWidget13setItemWidgetEP15QTreeWidgetItemiP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemWidget", args)
  }

}

// editItem(class QTreeWidgetItem *, int)
func (this *QTreeWidget) editItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget8editItemEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "editItem", args)
  }

}

// setSelectionModel(class QItemSelectionModel *)
func (this *QTreeWidget) setSelectionModel(args ...interface{}) () {
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
    C._ZN11QTreeWidget17setSelectionModelEP19QItemSelectionModel(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setSelectionModel", args)
  }

}

// currentItem()
func (this *QTreeWidget) currentItem(args ...interface{}) () {
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
    C._ZNK11QTreeWidget11currentItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentItem", args)
  }

}

// isItemSelected(const class QTreeWidgetItem *)
func (this *QTreeWidget) isItemSelected(args ...interface{}) () {
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
    C._ZNK11QTreeWidget14isItemSelectedEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemSelected", args)
  }

}

// itemBelow(const class QTreeWidgetItem *)
func (this *QTreeWidget) itemBelow(args ...interface{}) () {
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
    C._ZNK11QTreeWidget9itemBelowEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemBelow", args)
  }

}

// currentColumn()
func (this *QTreeWidget) currentColumn(args ...interface{}) () {
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
    C._ZNK11QTreeWidget13currentColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "currentColumn", args)
  }

}

// isItemHidden(const class QTreeWidgetItem *)
func (this *QTreeWidget) isItemHidden(args ...interface{}) () {
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
    C._ZNK11QTreeWidget12isItemHiddenEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemHidden", args)
  }

}

// columnCount()
func (this *QTreeWidget) columnCount(args ...interface{}) () {
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
    C._ZNK11QTreeWidget11columnCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "columnCount", args)
  }

}

// itemAbove(const class QTreeWidgetItem *)
func (this *QTreeWidget) itemAbove(args ...interface{}) () {
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
    C._ZNK11QTreeWidget9itemAboveEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAbove", args)
  }

}

// invisibleRootItem()
func (this *QTreeWidget) invisibleRootItem(args ...interface{}) () {
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
    C._ZNK11QTreeWidget17invisibleRootItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "invisibleRootItem", args)
  }

}

// isFirstItemColumnSpanned(const class QTreeWidgetItem *)
func (this *QTreeWidget) isFirstItemColumnSpanned(args ...interface{}) () {
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
    C._ZNK11QTreeWidget24isFirstItemColumnSpannedEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "isFirstItemColumnSpanned", args)
  }

}

// isItemExpanded(const class QTreeWidgetItem *)
func (this *QTreeWidget) isItemExpanded(args ...interface{}) () {
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
    C._ZNK11QTreeWidget14isItemExpandedEPK15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "isItemExpanded", args)
  }

}

// takeTopLevelItem(int)
func (this *QTreeWidget) takeTopLevelItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget16takeTopLevelItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "takeTopLevelItem", args)
  }

}

// topLevelItem(int)
func (this *QTreeWidget) topLevelItem(args ...interface{}) () {
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
    C._ZNK11QTreeWidget12topLevelItemEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItem", args)
  }

}

// topLevelItemCount()
func (this *QTreeWidget) topLevelItemCount(args ...interface{}) () {
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
    C._ZNK11QTreeWidget17topLevelItemCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "topLevelItemCount", args)
  }

}

// insertTopLevelItem(int, class QTreeWidgetItem *)
func (this *QTreeWidget) insertTopLevelItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget18insertTopLevelItemEiP15QTreeWidgetItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "insertTopLevelItem", args)
  }

}

// addTopLevelItem(class QTreeWidgetItem *)
func (this *QTreeWidget) addTopLevelItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget15addTopLevelItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "addTopLevelItem", args)
  }

}

// headerItem()
func (this *QTreeWidget) headerItem(args ...interface{}) () {
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
    C._ZNK11QTreeWidget10headerItemEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "headerItem", args)
  }

}

// setCurrentItem(class QTreeWidgetItem *)
func (this *QTreeWidget) setCurrentItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi
    // invoke: void setCurrentItem(class QTreeWidgetItem *, int)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN11QTreeWidget14setCurrentItemEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setCurrentItem", args)
  }

}

// metaObject()
func (this *QTreeWidget) metaObject(args ...interface{}) () {
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
    C._ZNK11QTreeWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "metaObject", args)
  }

}

// ~QTreeWidget()
func (this *QTreeWidget) FreeQTreeWidget(args ...interface{}) () {
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
    C._ZN11QTreeWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "~QTreeWidget", args)
  }

}

// itemAt(int, int)
func (this *QTreeWidget) itemAt(args ...interface{}) () {
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
    C._ZNK11QTreeWidget6itemAtEii(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK11QTreeWidget6itemAtERK6QPoint
    // invoke: QTreeWidgetItem * itemAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK11QTreeWidget6itemAtERK6QPoint(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "itemAt", args)
  }

}

// clear()
func (this *QTreeWidget) clear(args ...interface{}) () {
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
    C._ZN11QTreeWidget5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidget", "clear", args)
  }

}

// setItemHidden(const class QTreeWidgetItem *, _Bool)
func (this *QTreeWidget) setItemHidden(args ...interface{}) () {
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
    C._ZN11QTreeWidget13setItemHiddenEPK15QTreeWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setItemHidden", args)
  }

}

// closePersistentEditor(class QTreeWidgetItem *, int)
func (this *QTreeWidget) closePersistentEditor(args ...interface{}) () {
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
    C._ZN11QTreeWidget21closePersistentEditorEP15QTreeWidgetItemi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidget", "closePersistentEditor", args)
  }

}

// setHeaderItem(class QTreeWidgetItem *)
func (this *QTreeWidget) setHeaderItem(args ...interface{}) () {
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
    C._ZN11QTreeWidget13setHeaderItemEP15QTreeWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidget", "setHeaderItem", args)
  }

}

// setTextAlignment(int, int)
func (this *QTreeWidgetItem) setTextAlignment(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem16setTextAlignmentEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextAlignment", args)
  }

}

// columnCount()
func (this *QTreeWidgetItem) columnCount(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem11columnCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "columnCount", args)
  }

}

// addChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) addChild(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem8addChildEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "addChild", args)
  }

}

// setSelected(_Bool)
func (this *QTreeWidgetItem) setSelected(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem11setSelectedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSelected", args)
  }

}

// text(int)
func (this *QTreeWidgetItem) text(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem4textEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "text", args)
  }

}

// whatsThis(int)
func (this *QTreeWidgetItem) whatsThis(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem9whatsThisEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "whatsThis", args)
  }

}

// type()
func (this *QTreeWidgetItem) type_(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "type", args)
  }

}

// font(int)
func (this *QTreeWidgetItem) font(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem4fontEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "font", args)
  }

}

// ~QTreeWidgetItem()
func (this *QTreeWidgetItem) FreeQTreeWidgetItem(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "~QTreeWidgetItem", args)
  }

}

// setForeground(int, const class QBrush &)
func (this *QTreeWidgetItem) setForeground(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem13setForegroundEiRK6QBrush(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setForeground", args)
  }

}

// isSelected()
func (this *QTreeWidgetItem) isSelected(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10isSelectedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isSelected", args)
  }

}

// isDisabled()
func (this *QTreeWidgetItem) isDisabled(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10isDisabledEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isDisabled", args)
  }

}

// indexOfChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) indexOfChild(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem12indexOfChildEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "indexOfChild", args)
  }

}

// QTreeWidgetItem(class QTreeWidgetItem *, int)
func NewQTreeWidgetItem(args ...interface{}) QTreeWidgetItem {
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
    C._ZN15QTreeWidgetItemC2EPS_i(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN15QTreeWidgetItemC1Ei
    // invoke: void QTreeWidgetItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QTreeWidgetItemC2Ei(qthis, arg0)
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
    C._ZN15QTreeWidgetItemC2EPS_S0_i(qthis, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN15QTreeWidgetItemC1EP11QTreeWidgeti
    // invoke: void QTreeWidgetItem(class QTreeWidget *, int)
    var arg0 = args[0].(QTreeWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QTreeWidgetItemC2EP11QTreeWidgeti(qthis, arg0, arg1)
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
    C._ZN15QTreeWidgetItemC2EP11QTreeWidgetRK11QStringListi(qthis, arg0, arg1, arg2)
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
    C._ZN15QTreeWidgetItemC2EP11QTreeWidgetPS_i(qthis, arg0, arg1, arg2)
  case 6:
    // invoke: _ZN15QTreeWidgetItemC1ERK11QStringListi
    // invoke: void QTreeWidgetItem(const class QStringList &, int)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QTreeWidgetItemC2ERK11QStringListi(qthis, arg0, arg1)
  case 7:
    // invoke: _ZN15QTreeWidgetItemC1ERKS_
    // invoke: void QTreeWidgetItem(const class QTreeWidgetItem &)
    var arg0 = args[0].(QTreeWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN15QTreeWidgetItemC2ERKS_(qthis, arg0)
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
    C._ZN15QTreeWidgetItemC2EPS_RK11QStringListi(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "QTreeWidgetItem", args)
  }

  return QTreeWidgetItem{}
}

// write(class QDataStream &)
func (this *QTreeWidgetItem) write(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem5writeER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "write", args)
  }

}

// setExpanded(_Bool)
func (this *QTreeWidgetItem) setExpanded(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem11setExpandedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setExpanded", args)
  }

}

// insertChild(int, class QTreeWidgetItem *)
func (this *QTreeWidgetItem) insertChild(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem11insertChildEiPS_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "insertChild", args)
  }

}

// backgroundColor(int)
func (this *QTreeWidgetItem) backgroundColor(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem15backgroundColorEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "backgroundColor", args)
  }

}

// setBackground(int, const class QBrush &)
func (this *QTreeWidgetItem) setBackground(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem13setBackgroundEiRK6QBrush(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackground", args)
  }

}

// textColor(int)
func (this *QTreeWidgetItem) textColor(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem9textColorEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textColor", args)
  }

}

// setIcon(int, const class QIcon &)
func (this *QTreeWidgetItem) setIcon(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem7setIconEiRK5QIcon(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setIcon", args)
  }

}

// childCount()
func (this *QTreeWidgetItem) childCount(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10childCountEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "childCount", args)
  }

}

// setData(int, int, const class QVariant &)
func (this *QTreeWidgetItem) setData(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem7setDataEiiRK8QVariant(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setData", args)
  }

}

// statusTip(int)
func (this *QTreeWidgetItem) statusTip(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem9statusTipEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "statusTip", args)
  }

}

// foreground(int)
func (this *QTreeWidgetItem) foreground(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10foregroundEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "foreground", args)
  }

}

// treeWidget()
func (this *QTreeWidgetItem) treeWidget(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10treeWidgetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "treeWidget", args)
  }

}

// parent()
func (this *QTreeWidgetItem) parent(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem6parentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "parent", args)
  }

}

// isFirstColumnSpanned()
func (this *QTreeWidgetItem) isFirstColumnSpanned(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem20isFirstColumnSpannedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isFirstColumnSpanned", args)
  }

}

// setFirstColumnSpanned(_Bool)
func (this *QTreeWidgetItem) setFirstColumnSpanned(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem21setFirstColumnSpannedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFirstColumnSpanned", args)
  }

}

// clone()
func (this *QTreeWidgetItem) clone(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem5cloneEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "clone", args)
  }

}

// setBackgroundColor(int, const class QColor &)
func (this *QTreeWidgetItem) setBackgroundColor(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem18setBackgroundColorEiRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setBackgroundColor", args)
  }

}

// toolTip(int)
func (this *QTreeWidgetItem) toolTip(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem7toolTipEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "toolTip", args)
  }

}

// setFont(int, const class QFont &)
func (this *QTreeWidgetItem) setFont(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem7setFontEiRK5QFont(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setFont", args)
  }

}

// childIndicatorPolicy()
func (this *QTreeWidgetItem) childIndicatorPolicy(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem20childIndicatorPolicyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "childIndicatorPolicy", args)
  }

}

// setWhatsThis(int, const class QString &)
func (this *QTreeWidgetItem) setWhatsThis(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem12setWhatsThisEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setWhatsThis", args)
  }

}

// background(int)
func (this *QTreeWidgetItem) background(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10backgroundEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "background", args)
  }

}

// child(int)
func (this *QTreeWidgetItem) child(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem5childEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "child", args)
  }

}

// data(int, int)
func (this *QTreeWidgetItem) data(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem4dataEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "data", args)
  }

}

// setTextColor(int, const class QColor &)
func (this *QTreeWidgetItem) setTextColor(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem12setTextColorEiRK6QColor(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setTextColor", args)
  }

}

// icon(int)
func (this *QTreeWidgetItem) icon(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem4iconEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "icon", args)
  }

}

// isExpanded()
func (this *QTreeWidgetItem) isExpanded(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10isExpandedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isExpanded", args)
  }

}

// setSizeHint(int, const class QSize &)
func (this *QTreeWidgetItem) setSizeHint(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem11setSizeHintEiRK5QSize(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setSizeHint", args)
  }

}

// setToolTip(int, const class QString &)
func (this *QTreeWidgetItem) setToolTip(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem10setToolTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setToolTip", args)
  }

}

// sizeHint(int)
func (this *QTreeWidgetItem) sizeHint(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem8sizeHintEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "sizeHint", args)
  }

}

// read(class QDataStream &)
func (this *QTreeWidgetItem) read(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem4readER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "read", args)
  }

}

// isHidden()
func (this *QTreeWidgetItem) isHidden(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem8isHiddenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "isHidden", args)
  }

}

// setText(int, const class QString &)
func (this *QTreeWidgetItem) setText(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem7setTextEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setText", args)
  }

}

// setStatusTip(int, const class QString &)
func (this *QTreeWidgetItem) setStatusTip(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem12setStatusTipEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setStatusTip", args)
  }

}

// removeChild(class QTreeWidgetItem *)
func (this *QTreeWidgetItem) removeChild(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem11removeChildEPS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "removeChild", args)
  }

}

// setDisabled(_Bool)
func (this *QTreeWidgetItem) setDisabled(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem11setDisabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setDisabled", args)
  }

}

// flags()
func (this *QTreeWidgetItem) flags(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "flags", args)
  }

}

// checkState(int)
func (this *QTreeWidgetItem) checkState(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem10checkStateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "checkState", args)
  }

}

// takeChildren()
func (this *QTreeWidgetItem) takeChildren(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem12takeChildrenEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChildren", args)
  }

}

// takeChild(int)
func (this *QTreeWidgetItem) takeChild(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem9takeChildEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "takeChild", args)
  }

}

// textAlignment(int)
func (this *QTreeWidgetItem) textAlignment(args ...interface{}) () {
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
    C._ZNK15QTreeWidgetItem13textAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "textAlignment", args)
  }

}

// setHidden(_Bool)
func (this *QTreeWidgetItem) setHidden(args ...interface{}) () {
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
    C._ZN15QTreeWidgetItem9setHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTreeWidgetItem", "setHidden", args)
  }

}

// <= body block end

