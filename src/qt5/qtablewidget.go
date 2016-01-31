package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qtablewidget.h
// dst-file: /src/widgets/qtablewidget.go
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
  // proto:  int QTableWidgetSelectionRange::columnCount();
extern void C_ZNK26QTableWidgetSelectionRange11columnCountEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::leftColumn();
extern void C_ZNK26QTableWidgetSelectionRange10leftColumnEv(void* qthis); // 2
  // proto:  void QTableWidgetSelectionRange::~QTableWidgetSelectionRange();
extern void C_ZN26QTableWidgetSelectionRangeD2Ev(void* qthis); // 4
  // proto:  int QTableWidgetSelectionRange::bottomRow();
extern void C_ZNK26QTableWidgetSelectionRange9bottomRowEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::rightColumn();
extern void C_ZNK26QTableWidgetSelectionRange11rightColumnEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::topRow();
extern void C_ZNK26QTableWidgetSelectionRange6topRowEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::rowCount();
extern void C_ZNK26QTableWidgetSelectionRange8rowCountEv(void* qthis); // 2
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(int top, int left, int bottom, int right);
extern void* C_ZN26QTableWidgetSelectionRangeC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 3
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(const QTableWidgetSelectionRange & other);
extern void* C_ZN26QTableWidgetSelectionRangeC2ERKS_(void* arg0); // 3
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange();
extern void* C_ZN26QTableWidgetSelectionRangeC2Ev(); // 3
  // proto:  int QTableWidget::columnCount();
extern void C_ZNK12QTableWidget11columnCountEv(void* qthis); // 4
  // proto:  void QTableWidget::removeCellWidget(int row, int column);
extern void C_ZN12QTableWidget16removeCellWidgetEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QTableWidget::setCellWidget(int row, int column, QWidget * widget);
extern void C_ZN12QTableWidget13setCellWidgetEiiP7QWidget(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QTableWidgetItem * QTableWidget::verticalHeaderItem(int row);
extern void C_ZNK12QTableWidget18verticalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setItemPrototype(const QTableWidgetItem * item);
extern void C_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::~QTableWidget();
extern void C_ZN12QTableWidgetD2Ev(void* qthis); // 4
  // proto:  void QTableWidget::setItemSelected(const QTableWidgetItem * item, bool select);
extern void C_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTableWidget::openPersistentEditor(QTableWidgetItem * item);
extern void C_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTableWidget::row(const QTableWidgetItem * item);
extern void C_ZNK12QTableWidget3rowEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::setColumnCount(int columns);
extern void C_ZN12QTableWidget14setColumnCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setCurrentCell(int row, int column);
extern void C_ZN12QTableWidget14setCurrentCellEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTableWidget::QTableWidget(int rows, int columns, QWidget * parent);
extern void* C_ZN12QTableWidgetC2EiiP7QWidget(int32_t arg0, int32_t arg1, void* arg2); // 3
  // proto:  void QTableWidget::QTableWidget(QWidget * parent);
extern void* C_ZN12QTableWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  void QTableWidget::setRowCount(int rows);
extern void C_ZN12QTableWidget11setRowCountEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::editItem(QTableWidgetItem * item);
extern void C_ZN12QTableWidget8editItemEP16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::setHorizontalHeaderItem(int column, QTableWidgetItem * item);
extern void C_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QTableWidgetItem * QTableWidget::currentItem();
extern void C_ZNK12QTableWidget11currentItemEv(void* qthis); // 4
  // proto:  void QTableWidget::insertColumn(int column);
extern void C_ZN12QTableWidget12insertColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTableWidget::isItemSelected(const QTableWidgetItem * item);
extern void C_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::clearContents();
extern void C_ZN12QTableWidget13clearContentsEv(void* qthis); // 4
  // proto:  int QTableWidget::currentColumn();
extern void C_ZNK12QTableWidget13currentColumnEv(void* qthis); // 4
  // proto:  QTableWidgetItem * QTableWidget::itemAt(const QPoint & p);
extern void C_ZNK12QTableWidget6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QTableWidgetItem * QTableWidget::itemAt(int x, int y);
extern void C_ZNK12QTableWidget6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QTableWidgetItem * QTableWidget::takeItem(int row, int column);
extern void C_ZN12QTableWidget8takeItemEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTableWidgetItem * QTableWidget::takeHorizontalHeaderItem(int column);
extern void C_ZN12QTableWidget24takeHorizontalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTableWidget::isSortingEnabled();
extern void C_ZNK12QTableWidget16isSortingEnabledEv(void* qthis); // 4
  // proto:  QList<QTableWidgetSelectionRange> QTableWidget::selectedRanges();
extern void C_ZNK12QTableWidget14selectedRangesEv(void* qthis); // 4
  // proto:  QList<QTableWidgetItem *> QTableWidget::selectedItems();
extern void C_ZNK12QTableWidget13selectedItemsEv(void* qthis); // 4
  // proto:  void QTableWidget::removeRow(int row);
extern void C_ZN12QTableWidget9removeRowEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTableWidget::rowCount();
extern void C_ZNK12QTableWidget8rowCountEv(void* qthis); // 4
  // proto:  void QTableWidget::setItem(int row, int column, QTableWidgetItem * item);
extern void C_ZN12QTableWidget7setItemEiiP16QTableWidgetItem(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QTableWidgetItem * QTableWidget::takeVerticalHeaderItem(int row);
extern void C_ZN12QTableWidget22takeVerticalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  QTableWidgetItem * QTableWidget::horizontalHeaderItem(int column);
extern void C_ZNK12QTableWidget20horizontalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setHorizontalHeaderLabels(const QStringList & labels);
extern void C_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QRect QTableWidget::visualItemRect(const QTableWidgetItem * item);
extern void C_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::removeColumn(int column);
extern void C_ZN12QTableWidget12removeColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setCurrentItem(QTableWidgetItem * item);
extern void C_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTableWidget::currentRow();
extern void C_ZNK12QTableWidget10currentRowEv(void* qthis); // 4
  // proto:  const QMetaObject * QTableWidget::metaObject();
extern void C_ZNK12QTableWidget10metaObjectEv(void* qthis); // 4
  // proto:  const QTableWidgetItem * QTableWidget::itemPrototype();
extern void C_ZNK12QTableWidget13itemPrototypeEv(void* qthis); // 4
  // proto:  int QTableWidget::visualRow(int logicalRow);
extern void C_ZNK12QTableWidget9visualRowEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTableWidget::column(const QTableWidgetItem * item);
extern void C_ZNK12QTableWidget6columnEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QWidget * QTableWidget::cellWidget(int row, int column);
extern void C_ZNK12QTableWidget10cellWidgetEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTableWidget::setRangeSelected(const QTableWidgetSelectionRange & range, bool select);
extern void C_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTableWidget::setSortingEnabled(bool enable);
extern void C_ZN12QTableWidget17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QTableWidget::insertRow(int row);
extern void C_ZN12QTableWidget9insertRowEi(void* qthis, int32_t arg0); // 4
  // proto:  QTableWidgetItem * QTableWidget::item(int row, int column);
extern void C_ZNK12QTableWidget4itemEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTableWidget::visualColumn(int logicalColumn);
extern void C_ZNK12QTableWidget12visualColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setVerticalHeaderItem(int row, QTableWidgetItem * item);
extern void C_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QTableWidget::closePersistentEditor(QTableWidgetItem * item);
extern void C_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::clear();
extern void C_ZN12QTableWidget5clearEv(void* qthis); // 4
  // proto:  void QTableWidget::setVerticalHeaderLabels(const QStringList & labels);
extern void C_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  void QTableWidgetItem::setTextAlignment(int alignment);
extern void C_ZN16QTableWidgetItem16setTextAlignmentEi(void* qthis, int32_t arg0); // 2
  // proto:  void QTableWidgetItem::setSizeHint(const QSize & size);
extern void C_ZN16QTableWidgetItem11setSizeHintERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QTableWidgetItem::setSelected(bool select);
extern void C_ZN16QTableWidgetItem11setSelectedEb(void* qthis, bool arg0); // 2
  // proto:  QString QTableWidgetItem::text();
extern void C_ZNK16QTableWidgetItem4textEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QTableWidgetItem & other);
extern void* C_ZN16QTableWidgetItemC2ERKS_(void* arg0); // 3
  // proto:  void QTableWidgetItem::QTableWidgetItem(int type);
extern void* C_ZN16QTableWidgetItemC2Ei(int32_t arg0); // 3
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QIcon & icon, const QString & text, int type);
extern void* C_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QString & text, int type);
extern void* C_ZN16QTableWidgetItemC2ERK7QStringi(void* arg0, int32_t arg1); // 3
  // proto:  QFont QTableWidgetItem::font();
extern void C_ZNK16QTableWidgetItem4fontEv(void* qthis); // 2
  // proto:  int QTableWidgetItem::row();
extern void C_ZNK16QTableWidgetItem3rowEv(void* qthis); // 2
  // proto:  QString QTableWidgetItem::whatsThis();
extern void C_ZNK16QTableWidgetItem9whatsThisEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setForeground(const QBrush & brush);
extern void C_ZN16QTableWidgetItem13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  bool QTableWidgetItem::isSelected();
extern void C_ZNK16QTableWidgetItem10isSelectedEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::write(QDataStream & out);
extern void C_ZNK16QTableWidgetItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QTableWidgetItem::read(QDataStream & in);
extern void C_ZN16QTableWidgetItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  QColor QTableWidgetItem::backgroundColor();
extern void C_ZNK16QTableWidgetItem15backgroundColorEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setBackground(const QBrush & brush);
extern void C_ZN16QTableWidgetItem13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  QColor QTableWidgetItem::textColor();
extern void C_ZNK16QTableWidgetItem9textColorEv(void* qthis); // 2
  // proto:  int QTableWidgetItem::type();
extern void C_ZNK16QTableWidgetItem4typeEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setData(int role, const QVariant & value);
extern void C_ZN16QTableWidgetItem7setDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QBrush QTableWidgetItem::foreground();
extern void C_ZNK16QTableWidgetItem10foregroundEv(void* qthis); // 2
  // proto:  QString QTableWidgetItem::statusTip();
extern void C_ZNK16QTableWidgetItem9statusTipEv(void* qthis); // 2
  // proto:  QTableWidgetItem * QTableWidgetItem::clone();
extern void C_ZNK16QTableWidgetItem5cloneEv(void* qthis); // 4
  // proto:  void QTableWidgetItem::setBackgroundColor(const QColor & color);
extern void C_ZN16QTableWidgetItem18setBackgroundColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QString QTableWidgetItem::toolTip();
extern void C_ZNK16QTableWidgetItem7toolTipEv(void* qthis); // 2
  // proto:  QTableWidget * QTableWidgetItem::tableWidget();
extern void C_ZNK16QTableWidgetItem11tableWidgetEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setWhatsThis(const QString & whatsThis);
extern void C_ZN16QTableWidgetItem12setWhatsThisERK7QString(void* qthis, void* arg0); // 2
  // proto:  QBrush QTableWidgetItem::background();
extern void C_ZNK16QTableWidgetItem10backgroundEv(void* qthis); // 2
  // proto:  QVariant QTableWidgetItem::data(int role);
extern void C_ZNK16QTableWidgetItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidgetItem::setTextColor(const QColor & color);
extern void C_ZN16QTableWidgetItem12setTextColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QIcon QTableWidgetItem::icon();
extern void C_ZNK16QTableWidgetItem4iconEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setToolTip(const QString & toolTip);
extern void C_ZN16QTableWidgetItem10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  QSize QTableWidgetItem::sizeHint();
extern void C_ZNK16QTableWidgetItem8sizeHintEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::~QTableWidgetItem();
extern void C_ZN16QTableWidgetItemD2Ev(void* qthis); // 4
  // proto:  int QTableWidgetItem::column();
extern void C_ZNK16QTableWidgetItem6columnEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setIcon(const QIcon & icon);
extern void C_ZN16QTableWidgetItem7setIconERK5QIcon(void* qthis, void* arg0); // 2
  // proto:  void QTableWidgetItem::setText(const QString & text);
extern void C_ZN16QTableWidgetItem7setTextERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QTableWidgetItem::setStatusTip(const QString & statusTip);
extern void C_ZN16QTableWidgetItem12setStatusTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  Qt::ItemFlags QTableWidgetItem::flags();
extern void C_ZNK16QTableWidgetItem5flagsEv(void* qthis); // 2
  // proto:  Qt::CheckState QTableWidgetItem::checkState();
extern void C_ZNK16QTableWidgetItem10checkStateEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setFont(const QFont & font);
extern void C_ZN16QTableWidgetItem7setFontERK5QFont(void* qthis, void* arg0); // 2
  // proto:  int QTableWidgetItem::textAlignment();
extern void C_ZNK16QTableWidgetItem13textAlignmentEv(void* qthis); // 2
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

// class sizeof(QTableWidgetSelectionRange)=16
type QTableWidgetSelectionRange struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTableWidget)=1
type QTableWidget struct {
  /*qbase*/ QTableView;
  qclsinst unsafe.Pointer /* *C.void */;
//  _itemDoubleClicked QTableWidget_itemDoubleClicked_signal;
//  _cellEntered QTableWidget_cellEntered_signal;
//  _itemClicked QTableWidget_itemClicked_signal;
//  _currentItemChanged QTableWidget_currentItemChanged_signal;
//  _itemEntered QTableWidget_itemEntered_signal;
//  _itemPressed QTableWidget_itemPressed_signal;
//  _cellClicked QTableWidget_cellClicked_signal;
//  _itemSelectionChanged QTableWidget_itemSelectionChanged_signal;
//  _cellChanged QTableWidget_cellChanged_signal;
//  _itemActivated QTableWidget_itemActivated_signal;
//  _cellActivated QTableWidget_cellActivated_signal;
//  _itemChanged QTableWidget_itemChanged_signal;
//  _currentCellChanged QTableWidget_currentCellChanged_signal;
//  _cellDoubleClicked QTableWidget_cellDoubleClicked_signal;
//  _cellPressed QTableWidget_cellPressed_signal;
}

// class sizeof(QTableWidgetItem)=1
type QTableWidgetItem struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount()
func (this *QTableWidgetSelectionRange) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange11columnCountEv
    // invoke: int columnCount()
    var ret = C.C_ZNK26QTableWidgetSelectionRange11columnCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "columnCount", args)
  }

}

// leftColumn()
func (this *QTableWidgetSelectionRange) leftColumn(args ...interface{}) () {
  // leftColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange10leftColumnEv
    // invoke: int leftColumn()
    var ret = C.C_ZNK26QTableWidgetSelectionRange10leftColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "leftColumn", args)
  }

}

// ~QTableWidgetSelectionRange()
func (this *QTableWidgetSelectionRange) FreeQTableWidgetSelectionRange(args ...interface{}) () {
  // ~QTableWidgetSelectionRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QTableWidgetSelectionRangeD0Ev
    // invoke: void ~QTableWidgetSelectionRange()
    C.C_ZN26QTableWidgetSelectionRangeD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "~QTableWidgetSelectionRange", args)
  }

}

// bottomRow()
func (this *QTableWidgetSelectionRange) bottomRow(args ...interface{}) () {
  // bottomRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange9bottomRowEv
    // invoke: int bottomRow()
    var ret = C.C_ZNK26QTableWidgetSelectionRange9bottomRowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "bottomRow", args)
  }

}

// rightColumn()
func (this *QTableWidgetSelectionRange) rightColumn(args ...interface{}) () {
  // rightColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange11rightColumnEv
    // invoke: int rightColumn()
    var ret = C.C_ZNK26QTableWidgetSelectionRange11rightColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rightColumn", args)
  }

}

// topRow()
func (this *QTableWidgetSelectionRange) topRow(args ...interface{}) () {
  // topRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange6topRowEv
    // invoke: int topRow()
    var ret = C.C_ZNK26QTableWidgetSelectionRange6topRowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "topRow", args)
  }

}

// rowCount()
func (this *QTableWidgetSelectionRange) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QTableWidgetSelectionRange8rowCountEv
    // invoke: int rowCount()
    var ret = C.C_ZNK26QTableWidgetSelectionRange8rowCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rowCount", args)
  }

}

// QTableWidgetSelectionRange(int, int, int, int)
func NewQTableWidgetSelectionRange(args ...interface{}) *QTableWidgetSelectionRange {
  // QTableWidgetSelectionRange(int, int, int, int)
  // QTableWidgetSelectionRange(const class QTableWidgetSelectionRange &)
  // QTableWidgetSelectionRange()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTableWidgetSelectionRange{}) // "const QTableWidgetSelectionRange &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QTableWidgetSelectionRangeC1Eiiii
    // invoke: void QTableWidgetSelectionRange(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QTableWidgetSelectionRangeC2Eiiii(arg0, arg1, arg2, arg3)
    return &QTableWidgetSelectionRange{qclsinst:qthis}
  case 1:
    // invoke: _ZN26QTableWidgetSelectionRangeC1ERKS_
    // invoke: void QTableWidgetSelectionRange(const class QTableWidgetSelectionRange &)
    var arg0 = args[0].(QTableWidgetSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QTableWidgetSelectionRangeC2ERKS_(arg0)
    return &QTableWidgetSelectionRange{qclsinst:qthis}
  case 2:
    // invoke: _ZN26QTableWidgetSelectionRangeC1Ev
    // invoke: void QTableWidgetSelectionRange()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QTableWidgetSelectionRangeC2Ev()
    return &QTableWidgetSelectionRange{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "QTableWidgetSelectionRange", args)
  }

  return nil // QTableWidgetSelectionRange{qclsinst:qthis}
}

// columnCount()
func (this *QTableWidget) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget11columnCountEv
    // invoke: int columnCount()
    var ret = C.C_ZNK12QTableWidget11columnCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "columnCount", args)
  }

}

// removeCellWidget(int, int)
func (this *QTableWidget) removeCellWidget(args ...interface{}) () {
  // removeCellWidget(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16removeCellWidgetEii
    // invoke: void removeCellWidget(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget16removeCellWidgetEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "removeCellWidget", args)
  }

}

// setCellWidget(int, int, class QWidget *)
func (this *QTableWidget) setCellWidget(args ...interface{}) () {
  // setCellWidget(int, int, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget13setCellWidgetEiiP7QWidget
    // invoke: void setCellWidget(int, int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QTableWidget13setCellWidgetEiiP7QWidget(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTableWidget", "setCellWidget", args)
  }

}

// verticalHeaderItem(int)
func (this *QTableWidget) verticalHeaderItem(args ...interface{}) () {
  // verticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget18verticalHeaderItemEi
    // invoke: QTableWidgetItem * verticalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget18verticalHeaderItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "verticalHeaderItem", args)
  }

}

// setItemPrototype(const class QTableWidgetItem *)
func (this *QTableWidget) setItemPrototype(args ...interface{}) () {
  // setItemPrototype(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem
    // invoke: void setItemPrototype(const class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemPrototype", args)
  }

}

// ~QTableWidget()
func (this *QTableWidget) FreeQTableWidget(args ...interface{}) () {
  // ~QTableWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidgetD0Ev
    // invoke: void ~QTableWidget()
    C.C_ZN12QTableWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "~QTableWidget", args)
  }

}

// setItemSelected(const class QTableWidgetItem *, _Bool)
func (this *QTableWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QTableWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb
    // invoke: void setItemSelected(const class QTableWidgetItem *, _Bool)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemSelected", args)
  }

}

// openPersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem
    // invoke: void openPersistentEditor(class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "openPersistentEditor", args)
  }

}

// row(const class QTableWidgetItem *)
func (this *QTableWidget) row(args ...interface{}) () {
  // row(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget3rowEPK16QTableWidgetItem
    // invoke: int row(const class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget3rowEPK16QTableWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "row", args)
  }

}

// setColumnCount(int)
func (this *QTableWidget) setColumnCount(args ...interface{}) () {
  // setColumnCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setColumnCountEi
    // invoke: void setColumnCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget14setColumnCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setColumnCount", args)
  }

}

// setCurrentCell(int, int)
func (this *QTableWidget) setCurrentCell(args ...interface{}) () {
  // setCurrentCell(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setCurrentCellEii
    // invoke: void setCurrentCell(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget14setCurrentCellEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentCell", args)
  }

}

// QTableWidget(int, int, class QWidget *)
func NewQTableWidget(args ...interface{}) *QTableWidget {
  // QTableWidget(int, int, class QWidget *)
  // QTableWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QWidget{}) // "QWidget *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidgetC1EiiP7QWidget
    // invoke: void QTableWidget(int, int, class QWidget *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QWidget).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTableWidgetC2EiiP7QWidget(arg0, arg1, arg2)
    return &QTableWidget{qclsinst:qthis}
  case 1:
    // invoke: _ZN12QTableWidgetC1EP7QWidget
    // invoke: void QTableWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTableWidgetC2EP7QWidget(arg0)
    return &QTableWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTableWidget", "QTableWidget", args)
  }

  return nil // QTableWidget{qclsinst:qthis}
}

// setRowCount(int)
func (this *QTableWidget) setRowCount(args ...interface{}) () {
  // setRowCount(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget11setRowCountEi
    // invoke: void setRowCount(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget11setRowCountEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setRowCount", args)
  }

}

// editItem(class QTableWidgetItem *)
func (this *QTableWidget) editItem(args ...interface{}) () {
  // editItem(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget8editItemEP16QTableWidgetItem
    // invoke: void editItem(class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget8editItemEP16QTableWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "editItem", args)
  }

}

// setHorizontalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) setHorizontalHeaderItem(args ...interface{}) () {
  // setHorizontalHeaderItem(int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem
    // invoke: void setHorizontalHeaderItem(int, class QTableWidgetItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderItem", args)
  }

}

// currentItem()
func (this *QTableWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget11currentItemEv
    // invoke: QTableWidgetItem * currentItem()
    var ret = C.C_ZNK12QTableWidget11currentItemEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "currentItem", args)
  }

}

// insertColumn(int)
func (this *QTableWidget) insertColumn(args ...interface{}) () {
  // insertColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget12insertColumnEi
    // invoke: void insertColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget12insertColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "insertColumn", args)
  }

}

// isItemSelected(const class QTableWidgetItem *)
func (this *QTableWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem
    // invoke: bool isItemSelected(const class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "isItemSelected", args)
  }

}

// clearContents()
func (this *QTableWidget) clearContents(args ...interface{}) () {
  // clearContents()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget13clearContentsEv
    // invoke: void clearContents()
    C.C_ZN12QTableWidget13clearContentsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "clearContents", args)
  }

}

// currentColumn()
func (this *QTableWidget) currentColumn(args ...interface{}) () {
  // currentColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13currentColumnEv
    // invoke: int currentColumn()
    var ret = C.C_ZNK12QTableWidget13currentColumnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "currentColumn", args)
  }

}

// itemAt(const class QPoint &)
func (this *QTableWidget) itemAt(args ...interface{}) () {
  // itemAt(const class QPoint &)
  // itemAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget6itemAtERK6QPoint
    // invoke: QTableWidgetItem * itemAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget6itemAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK12QTableWidget6itemAtEii
    // invoke: QTableWidgetItem * itemAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK12QTableWidget6itemAtEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "itemAt", args)
  }

}

// takeItem(int, int)
func (this *QTableWidget) takeItem(args ...interface{}) () {
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
    // invoke: _ZN12QTableWidget8takeItemEii
    // invoke: QTableWidgetItem * takeItem(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZN12QTableWidget8takeItemEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "takeItem", args)
  }

}

// takeHorizontalHeaderItem(int)
func (this *QTableWidget) takeHorizontalHeaderItem(args ...interface{}) () {
  // takeHorizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget24takeHorizontalHeaderItemEi
    // invoke: QTableWidgetItem * takeHorizontalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN12QTableWidget24takeHorizontalHeaderItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "takeHorizontalHeaderItem", args)
  }

}

// isSortingEnabled()
func (this *QTableWidget) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget16isSortingEnabledEv
    // invoke: bool isSortingEnabled()
    var ret = C.C_ZNK12QTableWidget16isSortingEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "isSortingEnabled", args)
  }

}

// selectedRanges()
func (this *QTableWidget) selectedRanges(args ...interface{}) () {
  // selectedRanges()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14selectedRangesEv
    // invoke: QList<QTableWidgetSelectionRange> selectedRanges()
    C.C_ZNK12QTableWidget14selectedRangesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedRanges", args)
  }

}

// selectedItems()
func (this *QTableWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13selectedItemsEv
    // invoke: QList<QTableWidgetItem *> selectedItems()
    C.C_ZNK12QTableWidget13selectedItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedItems", args)
  }

}

// removeRow(int)
func (this *QTableWidget) removeRow(args ...interface{}) () {
  // removeRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget9removeRowEi
    // invoke: void removeRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget9removeRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "removeRow", args)
  }

}

// rowCount()
func (this *QTableWidget) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget8rowCountEv
    // invoke: int rowCount()
    var ret = C.C_ZNK12QTableWidget8rowCountEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "rowCount", args)
  }

}

// setItem(int, int, class QTableWidgetItem *)
func (this *QTableWidget) setItem(args ...interface{}) () {
  // setItem(int, int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget7setItemEiiP16QTableWidgetItem
    // invoke: void setItem(int, int, class QTableWidgetItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QTableWidget7setItemEiiP16QTableWidgetItem(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTableWidget", "setItem", args)
  }

}

// takeVerticalHeaderItem(int)
func (this *QTableWidget) takeVerticalHeaderItem(args ...interface{}) () {
  // takeVerticalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget22takeVerticalHeaderItemEi
    // invoke: QTableWidgetItem * takeVerticalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN12QTableWidget22takeVerticalHeaderItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "takeVerticalHeaderItem", args)
  }

}

// horizontalHeaderItem(int)
func (this *QTableWidget) horizontalHeaderItem(args ...interface{}) () {
  // horizontalHeaderItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget20horizontalHeaderItemEi
    // invoke: QTableWidgetItem * horizontalHeaderItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget20horizontalHeaderItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "horizontalHeaderItem", args)
  }

}

// setHorizontalHeaderLabels(const class QStringList &)
func (this *QTableWidget) setHorizontalHeaderLabels(args ...interface{}) () {
  // setHorizontalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList
    // invoke: void setHorizontalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderLabels", args)
  }

}

// visualItemRect(const class QTableWidgetItem *)
func (this *QTableWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem
    // invoke: QRect visualItemRect(const class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "visualItemRect", args)
  }

}

// removeColumn(int)
func (this *QTableWidget) removeColumn(args ...interface{}) () {
  // removeColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget12removeColumnEi
    // invoke: void removeColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget12removeColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "removeColumn", args)
  }

}

// setCurrentItem(class QTableWidgetItem *)
func (this *QTableWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem
    // invoke: void setCurrentItem(class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentItem", args)
  }

}

// currentRow()
func (this *QTableWidget) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10currentRowEv
    // invoke: int currentRow()
    var ret = C.C_ZNK12QTableWidget10currentRowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "currentRow", args)
  }

}

// metaObject()
func (this *QTableWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK12QTableWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "metaObject", args)
  }

}

// itemPrototype()
func (this *QTableWidget) itemPrototype(args ...interface{}) () {
  // itemPrototype()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget13itemPrototypeEv
    // invoke: const QTableWidgetItem * itemPrototype()
    var ret = C.C_ZNK12QTableWidget13itemPrototypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "itemPrototype", args)
  }

}

// visualRow(int)
func (this *QTableWidget) visualRow(args ...interface{}) () {
  // visualRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget9visualRowEi
    // invoke: int visualRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget9visualRowEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "visualRow", args)
  }

}

// column(const class QTableWidgetItem *)
func (this *QTableWidget) column(args ...interface{}) () {
  // column(const class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget6columnEPK16QTableWidgetItem
    // invoke: int column(const class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget6columnEPK16QTableWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "column", args)
  }

}

// cellWidget(int, int)
func (this *QTableWidget) cellWidget(args ...interface{}) () {
  // cellWidget(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget10cellWidgetEii
    // invoke: QWidget * cellWidget(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK12QTableWidget10cellWidgetEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "cellWidget", args)
  }

}

// setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
func (this *QTableWidget) setRangeSelected(args ...interface{}) () {
  // setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetSelectionRange{}) // "const QTableWidgetSelectionRange &"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb
    // invoke: void setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
    var arg0 = args[0].(QTableWidgetSelectionRange).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setRangeSelected", args)
  }

}

// setSortingEnabled(_Bool)
func (this *QTableWidget) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget17setSortingEnabledEb
    // invoke: void setSortingEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget17setSortingEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setSortingEnabled", args)
  }

}

// insertRow(int)
func (this *QTableWidget) insertRow(args ...interface{}) () {
  // insertRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget9insertRowEi
    // invoke: void insertRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget9insertRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "insertRow", args)
  }

}

// item(int, int)
func (this *QTableWidget) item(args ...interface{}) () {
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
    // invoke: _ZNK12QTableWidget4itemEii
    // invoke: QTableWidgetItem * item(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK12QTableWidget4itemEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "item", args)
  }

}

// visualColumn(int)
func (this *QTableWidget) visualColumn(args ...interface{}) () {
  // visualColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget12visualColumnEi
    // invoke: int visualColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK12QTableWidget12visualColumnEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidget", "visualColumn", args)
  }

}

// setVerticalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) setVerticalHeaderItem(args ...interface{}) () {
  // setVerticalHeaderItem(int, class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem
    // invoke: void setVerticalHeaderItem(int, class QTableWidgetItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderItem", args)
  }

}

// closePersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QTableWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem
    // invoke: void closePersistentEditor(class QTableWidgetItem *)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "closePersistentEditor", args)
  }

}

// clear()
func (this *QTableWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget5clearEv
    // invoke: void clear()
    C.C_ZN12QTableWidget5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "clear", args)
  }

}

// setVerticalHeaderLabels(const class QStringList &)
func (this *QTableWidget) setVerticalHeaderLabels(args ...interface{}) () {
  // setVerticalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList
    // invoke: void setVerticalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderLabels", args)
  }

}

// setTextAlignment(int)
func (this *QTableWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem16setTextAlignmentEi
    // invoke: void setTextAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem16setTextAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextAlignment", args)
  }

}

// setSizeHint(const class QSize &)
func (this *QTableWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSizeHintERK5QSize
    // invoke: void setSizeHint(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem11setSizeHintERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSizeHint", args)
  }

}

// setSelected(_Bool)
func (this *QTableWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSelectedEb
    // invoke: void setSelected(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem11setSelectedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSelected", args)
  }

}

// text()
func (this *QTableWidgetItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4textEv
    // invoke: QString text()
    var ret = C.C_ZNK16QTableWidgetItem4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "text", args)
  }

}

// QTableWidgetItem(const class QTableWidgetItem &)
func NewQTableWidgetItem(args ...interface{}) *QTableWidgetItem {
  // QTableWidgetItem(const class QTableWidgetItem &)
  // QTableWidgetItem(int)
  // QTableWidgetItem(const class QIcon &, const class QString &, int)
  // QTableWidgetItem(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItemC1ERKS_
    // invoke: void QTableWidgetItem(const class QTableWidgetItem &)
    var arg0 = args[0].(QTableWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2ERKS_(arg0)
    return &QTableWidgetItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN16QTableWidgetItemC1Ei
    // invoke: void QTableWidgetItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2Ei(arg0)
    return &QTableWidgetItem{qclsinst:qthis}
  case 2:
    // invoke: _ZN16QTableWidgetItemC1ERK5QIconRK7QStringi
    // invoke: void QTableWidgetItem(const class QIcon &, const class QString &, int)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi(arg0, arg1, arg2)
    return &QTableWidgetItem{qclsinst:qthis}
  case 3:
    // invoke: _ZN16QTableWidgetItemC1ERK7QStringi
    // invoke: void QTableWidgetItem(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2ERK7QStringi(arg0, arg1)
    return &QTableWidgetItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "QTableWidgetItem", args)
  }

  return nil // QTableWidgetItem{qclsinst:qthis}
}

// font()
func (this *QTableWidgetItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4fontEv
    // invoke: QFont font()
    var ret = C.C_ZNK16QTableWidgetItem4fontEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "font", args)
  }

}

// row()
func (this *QTableWidgetItem) row(args ...interface{}) () {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem3rowEv
    // invoke: int row()
    var ret = C.C_ZNK16QTableWidgetItem3rowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "row", args)
  }

}

// whatsThis()
func (this *QTableWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9whatsThisEv
    // invoke: QString whatsThis()
    var ret = C.C_ZNK16QTableWidgetItem9whatsThisEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "whatsThis", args)
  }

}

// setForeground(const class QBrush &)
func (this *QTableWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setForegroundERK6QBrush
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem13setForegroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setForeground", args)
  }

}

// isSelected()
func (this *QTableWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10isSelectedEv
    // invoke: bool isSelected()
    var ret = C.C_ZNK16QTableWidgetItem10isSelectedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "isSelected", args)
  }

}

// write(class QDataStream &)
func (this *QTableWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QTableWidgetItem5writeER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "write", args)
  }

}

// read(class QDataStream &)
func (this *QTableWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem4readER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "read", args)
  }

}

// backgroundColor()
func (this *QTableWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem15backgroundColorEv
    // invoke: QColor backgroundColor()
    var ret = C.C_ZNK16QTableWidgetItem15backgroundColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "backgroundColor", args)
  }

}

// setBackground(const class QBrush &)
func (this *QTableWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackground", args)
  }

}

// textColor()
func (this *QTableWidgetItem) textColor(args ...interface{}) () {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9textColorEv
    // invoke: QColor textColor()
    var ret = C.C_ZNK16QTableWidgetItem9textColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textColor", args)
  }

}

// type()
func (this *QTableWidgetItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK16QTableWidgetItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "type", args)
  }

}

// setData(int, const class QVariant &)
func (this *QTableWidgetItem) setData(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setDataEiRK8QVariant
    // invoke: void setData(int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QTableWidgetItem7setDataEiRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setData", args)
  }

}

// foreground()
func (this *QTableWidgetItem) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10foregroundEv
    // invoke: QBrush foreground()
    var ret = C.C_ZNK16QTableWidgetItem10foregroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "foreground", args)
  }

}

// statusTip()
func (this *QTableWidgetItem) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem9statusTipEv
    // invoke: QString statusTip()
    var ret = C.C_ZNK16QTableWidgetItem9statusTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "statusTip", args)
  }

}

// clone()
func (this *QTableWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5cloneEv
    // invoke: QTableWidgetItem * clone()
    var ret = C.C_ZNK16QTableWidgetItem5cloneEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "clone", args)
  }

}

// setBackgroundColor(const class QColor &)
func (this *QTableWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem18setBackgroundColorERK6QColor
    // invoke: void setBackgroundColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem18setBackgroundColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackgroundColor", args)
  }

}

// toolTip()
func (this *QTableWidgetItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem7toolTipEv
    // invoke: QString toolTip()
    var ret = C.C_ZNK16QTableWidgetItem7toolTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "toolTip", args)
  }

}

// tableWidget()
func (this *QTableWidgetItem) tableWidget(args ...interface{}) () {
  // tableWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem11tableWidgetEv
    // invoke: QTableWidget * tableWidget()
    var ret = C.C_ZNK16QTableWidgetItem11tableWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "tableWidget", args)
  }

}

// setWhatsThis(const class QString &)
func (this *QTableWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem12setWhatsThisERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setWhatsThis", args)
  }

}

// background()
func (this *QTableWidgetItem) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10backgroundEv
    // invoke: QBrush background()
    var ret = C.C_ZNK16QTableWidgetItem10backgroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "background", args)
  }

}

// data(int)
func (this *QTableWidgetItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4dataEi
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK16QTableWidgetItem4dataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "data", args)
  }

}

// setTextColor(const class QColor &)
func (this *QTableWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setTextColorERK6QColor
    // invoke: void setTextColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem12setTextColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextColor", args)
  }

}

// icon()
func (this *QTableWidgetItem) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem4iconEv
    // invoke: QIcon icon()
    var ret = C.C_ZNK16QTableWidgetItem4iconEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "icon", args)
  }

}

// setToolTip(const class QString &)
func (this *QTableWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setToolTip", args)
  }

}

// sizeHint()
func (this *QTableWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem8sizeHintEv
    // invoke: QSize sizeHint()
    var ret = C.C_ZNK16QTableWidgetItem8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "sizeHint", args)
  }

}

// ~QTableWidgetItem()
func (this *QTableWidgetItem) FreeQTableWidgetItem(args ...interface{}) () {
  // ~QTableWidgetItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItemD0Ev
    // invoke: void ~QTableWidgetItem()
    C.C_ZN16QTableWidgetItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "~QTableWidgetItem", args)
  }

}

// column()
func (this *QTableWidgetItem) column(args ...interface{}) () {
  // column()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem6columnEv
    // invoke: int column()
    var ret = C.C_ZNK16QTableWidgetItem6columnEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "column", args)
  }

}

// setIcon(const class QIcon &)
func (this *QTableWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setIcon", args)
  }

}

// setText(const class QString &)
func (this *QTableWidgetItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setText", args)
  }

}

// setStatusTip(const class QString &)
func (this *QTableWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem12setStatusTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setStatusTip", args)
  }

}

// flags()
func (this *QTableWidgetItem) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK16QTableWidgetItem5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "flags", args)
  }

}

// checkState()
func (this *QTableWidgetItem) checkState(args ...interface{}) () {
  // checkState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem10checkStateEv
    // invoke: Qt::CheckState checkState()
    C.C_ZNK16QTableWidgetItem10checkStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "checkState", args)
  }

}

// setFont(const class QFont &)
func (this *QTableWidgetItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setFont", args)
  }

}

// textAlignment()
func (this *QTableWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem13textAlignmentEv
    // invoke: int textAlignment()
    var ret = C.C_ZNK16QTableWidgetItem13textAlignmentEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textAlignment", args)
  }

}

// <= body block end

