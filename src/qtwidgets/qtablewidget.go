package qtwidgets
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
import "qtgui"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  int QTableWidgetSelectionRange::columnCount();
extern int32_t C_ZNK26QTableWidgetSelectionRange11columnCountEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::leftColumn();
extern int32_t C_ZNK26QTableWidgetSelectionRange10leftColumnEv(void* qthis); // 2
  // proto:  void QTableWidgetSelectionRange::~QTableWidgetSelectionRange();
extern void C_ZN26QTableWidgetSelectionRangeD2Ev(void* qthis); // 4
  // proto:  int QTableWidgetSelectionRange::bottomRow();
extern int32_t C_ZNK26QTableWidgetSelectionRange9bottomRowEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::rightColumn();
extern int32_t C_ZNK26QTableWidgetSelectionRange11rightColumnEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::topRow();
extern int32_t C_ZNK26QTableWidgetSelectionRange6topRowEv(void* qthis); // 2
  // proto:  int QTableWidgetSelectionRange::rowCount();
extern int32_t C_ZNK26QTableWidgetSelectionRange8rowCountEv(void* qthis); // 2
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(int top, int left, int bottom, int right);
extern void* C_ZN26QTableWidgetSelectionRangeC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 3
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange(const QTableWidgetSelectionRange & other);
extern void* C_ZN26QTableWidgetSelectionRangeC2ERKS_(void* arg0); // 3
  // proto:  void QTableWidgetSelectionRange::QTableWidgetSelectionRange();
extern void* C_ZN26QTableWidgetSelectionRangeC2Ev(); // 3
  // proto:  int QTableWidget::columnCount();
extern int32_t C_ZNK12QTableWidget11columnCountEv(void* qthis); // 4
  // proto:  void QTableWidget::removeCellWidget(int row, int column);
extern void C_ZN12QTableWidget16removeCellWidgetEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  void QTableWidget::setCellWidget(int row, int column, QWidget * widget);
extern void C_ZN12QTableWidget13setCellWidgetEiiP7QWidget(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QTableWidgetItem * QTableWidget::verticalHeaderItem(int row);
extern void* C_ZNK12QTableWidget18verticalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setItemPrototype(const QTableWidgetItem * item);
extern void C_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::~QTableWidget();
extern void C_ZN12QTableWidgetD2Ev(void* qthis); // 4
  // proto:  void QTableWidget::setItemSelected(const QTableWidgetItem * item, bool select);
extern void C_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTableWidget::openPersistentEditor(QTableWidgetItem * item);
extern void C_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTableWidget::row(const QTableWidgetItem * item);
extern int32_t C_ZNK12QTableWidget3rowEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
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
extern void* C_ZNK12QTableWidget11currentItemEv(void* qthis); // 4
  // proto:  void QTableWidget::insertColumn(int column);
extern void C_ZN12QTableWidget12insertColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTableWidget::isItemSelected(const QTableWidgetItem * item);
extern bool C_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::clearContents();
extern void C_ZN12QTableWidget13clearContentsEv(void* qthis); // 4
  // proto:  int QTableWidget::currentColumn();
extern int32_t C_ZNK12QTableWidget13currentColumnEv(void* qthis); // 4
  // proto:  QTableWidgetItem * QTableWidget::itemAt(const QPoint & p);
extern void* C_ZNK12QTableWidget6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  QTableWidgetItem * QTableWidget::itemAt(int x, int y);
extern void* C_ZNK12QTableWidget6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QTableWidgetItem * QTableWidget::takeItem(int row, int column);
extern void* C_ZN12QTableWidget8takeItemEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QTableWidgetItem * QTableWidget::takeHorizontalHeaderItem(int column);
extern void* C_ZN12QTableWidget24takeHorizontalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QTableWidget::isSortingEnabled();
extern bool C_ZNK12QTableWidget16isSortingEnabledEv(void* qthis); // 4
  // proto:  QList<QTableWidgetSelectionRange> QTableWidget::selectedRanges();
extern void C_ZNK12QTableWidget14selectedRangesEv(void* qthis); // 4
  // proto:  QList<QTableWidgetItem *> QTableWidget::selectedItems();
extern void C_ZNK12QTableWidget13selectedItemsEv(void* qthis); // 4
  // proto:  void QTableWidget::removeRow(int row);
extern void C_ZN12QTableWidget9removeRowEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTableWidget::rowCount();
extern int32_t C_ZNK12QTableWidget8rowCountEv(void* qthis); // 4
  // proto:  void QTableWidget::setItem(int row, int column, QTableWidgetItem * item);
extern void C_ZN12QTableWidget7setItemEiiP16QTableWidgetItem(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QTableWidgetItem * QTableWidget::takeVerticalHeaderItem(int row);
extern void* C_ZN12QTableWidget22takeVerticalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  QTableWidgetItem * QTableWidget::horizontalHeaderItem(int column);
extern void* C_ZNK12QTableWidget20horizontalHeaderItemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setHorizontalHeaderLabels(const QStringList & labels);
extern void C_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList(void* qthis, void* arg0); // 4
  // proto:  QRect QTableWidget::visualItemRect(const QTableWidgetItem * item);
extern void* C_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QTableWidget::removeColumn(int column);
extern void C_ZN12QTableWidget12removeColumnEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidget::setCurrentItem(QTableWidgetItem * item);
extern void C_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QTableWidget::currentRow();
extern int32_t C_ZNK12QTableWidget10currentRowEv(void* qthis); // 4
  // proto:  const QMetaObject * QTableWidget::metaObject();
extern void C_ZNK12QTableWidget10metaObjectEv(void* qthis); // 4
  // proto:  const QTableWidgetItem * QTableWidget::itemPrototype();
extern void* C_ZNK12QTableWidget13itemPrototypeEv(void* qthis); // 4
  // proto:  int QTableWidget::visualRow(int logicalRow);
extern int32_t C_ZNK12QTableWidget9visualRowEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTableWidget::column(const QTableWidgetItem * item);
extern int32_t C_ZNK12QTableWidget6columnEPK16QTableWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QWidget * QTableWidget::cellWidget(int row, int column);
extern void* C_ZNK12QTableWidget10cellWidgetEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  void QTableWidget::setRangeSelected(const QTableWidgetSelectionRange & range, bool select);
extern void C_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  void QTableWidget::setSortingEnabled(bool enable);
extern void C_ZN12QTableWidget17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  void QTableWidget::insertRow(int row);
extern void C_ZN12QTableWidget9insertRowEi(void* qthis, int32_t arg0); // 4
  // proto:  QTableWidgetItem * QTableWidget::item(int row, int column);
extern void* C_ZNK12QTableWidget4itemEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  int QTableWidget::visualColumn(int logicalColumn);
extern int32_t C_ZNK12QTableWidget12visualColumnEi(void* qthis, int32_t arg0); // 4
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
extern void* C_ZNK16QTableWidgetItem4textEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QTableWidgetItem & other);
extern void* C_ZN16QTableWidgetItemC2ERKS_(void* arg0); // 3
  // proto:  void QTableWidgetItem::QTableWidgetItem(int type);
extern void* C_ZN16QTableWidgetItemC2Ei(int32_t arg0); // 3
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QIcon & icon, const QString & text, int type);
extern void* C_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QTableWidgetItem::QTableWidgetItem(const QString & text, int type);
extern void* C_ZN16QTableWidgetItemC2ERK7QStringi(void* arg0, int32_t arg1); // 3
  // proto:  QFont QTableWidgetItem::font();
extern void* C_ZNK16QTableWidgetItem4fontEv(void* qthis); // 2
  // proto:  int QTableWidgetItem::row();
extern int32_t C_ZNK16QTableWidgetItem3rowEv(void* qthis); // 2
  // proto:  QString QTableWidgetItem::whatsThis();
extern void* C_ZNK16QTableWidgetItem9whatsThisEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setForeground(const QBrush & brush);
extern void C_ZN16QTableWidgetItem13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  bool QTableWidgetItem::isSelected();
extern bool C_ZNK16QTableWidgetItem10isSelectedEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::write(QDataStream & out);
extern void C_ZNK16QTableWidgetItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QTableWidgetItem::read(QDataStream & in);
extern void C_ZN16QTableWidgetItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  QColor QTableWidgetItem::backgroundColor();
extern void* C_ZNK16QTableWidgetItem15backgroundColorEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setBackground(const QBrush & brush);
extern void C_ZN16QTableWidgetItem13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  QColor QTableWidgetItem::textColor();
extern void* C_ZNK16QTableWidgetItem9textColorEv(void* qthis); // 2
  // proto:  int QTableWidgetItem::type();
extern int32_t C_ZNK16QTableWidgetItem4typeEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setData(int role, const QVariant & value);
extern void C_ZN16QTableWidgetItem7setDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QBrush QTableWidgetItem::foreground();
extern void* C_ZNK16QTableWidgetItem10foregroundEv(void* qthis); // 2
  // proto:  QString QTableWidgetItem::statusTip();
extern void* C_ZNK16QTableWidgetItem9statusTipEv(void* qthis); // 2
  // proto:  QTableWidgetItem * QTableWidgetItem::clone();
extern void* C_ZNK16QTableWidgetItem5cloneEv(void* qthis); // 4
  // proto:  void QTableWidgetItem::setBackgroundColor(const QColor & color);
extern void C_ZN16QTableWidgetItem18setBackgroundColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QString QTableWidgetItem::toolTip();
extern void* C_ZNK16QTableWidgetItem7toolTipEv(void* qthis); // 2
  // proto:  QTableWidget * QTableWidgetItem::tableWidget();
extern void* C_ZNK16QTableWidgetItem11tableWidgetEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setWhatsThis(const QString & whatsThis);
extern void C_ZN16QTableWidgetItem12setWhatsThisERK7QString(void* qthis, void* arg0); // 2
  // proto:  QBrush QTableWidgetItem::background();
extern void* C_ZNK16QTableWidgetItem10backgroundEv(void* qthis); // 2
  // proto:  QVariant QTableWidgetItem::data(int role);
extern void* C_ZNK16QTableWidgetItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTableWidgetItem::setTextColor(const QColor & color);
extern void C_ZN16QTableWidgetItem12setTextColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QIcon QTableWidgetItem::icon();
extern void* C_ZNK16QTableWidgetItem4iconEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::setToolTip(const QString & toolTip);
extern void C_ZN16QTableWidgetItem10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  QSize QTableWidgetItem::sizeHint();
extern void* C_ZNK16QTableWidgetItem8sizeHintEv(void* qthis); // 2
  // proto:  void QTableWidgetItem::~QTableWidgetItem();
extern void C_ZN16QTableWidgetItemD2Ev(void* qthis); // 4
  // proto:  int QTableWidgetItem::column();
extern int32_t C_ZNK16QTableWidgetItem6columnEv(void* qthis); // 2
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
extern int32_t C_ZNK16QTableWidgetItem13textAlignmentEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {qtgui.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTableWidgetSelectionRange)=16
type QTableWidgetSelectionRange struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTableWidget)=1
type QTableWidget struct {
  /*qbase*/ QTableView;
  Qclsinst unsafe.Pointer /* *C.void */;
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// columnCount()
func (this *QTableWidgetSelectionRange) ColumnCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QTableWidgetSelectionRange11columnCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "columnCount", args)
  }

  return
}

// leftColumn()
func (this *QTableWidgetSelectionRange) LeftColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QTableWidgetSelectionRange10leftColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "leftColumn", args)
  }

  return
}

// ~QTableWidgetSelectionRange()
func (this *QTableWidgetSelectionRange) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN26QTableWidgetSelectionRangeD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "~QTableWidgetSelectionRange", args)
  }

  return
}

// bottomRow()
func (this *QTableWidgetSelectionRange) BottomRow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QTableWidgetSelectionRange9bottomRowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "bottomRow", args)
  }

  return
}

// rightColumn()
func (this *QTableWidgetSelectionRange) RightColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QTableWidgetSelectionRange11rightColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rightColumn", args)
  }

  return
}

// topRow()
func (this *QTableWidgetSelectionRange) TopRow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QTableWidgetSelectionRange6topRowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "topRow", args)
  }

  return
}

// rowCount()
func (this *QTableWidgetSelectionRange) RowCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QTableWidgetSelectionRange8rowCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "rowCount", args)
  }

  return
}

// QTableWidgetSelectionRange(int, int, int, int)
func GcfreeQTableWidgetSelectionRange(this *QTableWidgetSelectionRange) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QTableWidgetSelectionRangeC2Eiiii(arg0, arg1, arg2, arg3)
    this := &QTableWidgetSelectionRange{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetSelectionRange)
    return this
  case 1:
    // invoke: _ZN26QTableWidgetSelectionRangeC1ERKS_
    // invoke: void QTableWidgetSelectionRange(const class QTableWidgetSelectionRange &)
    var arg0 = args[0].(*QTableWidgetSelectionRange).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QTableWidgetSelectionRangeC2ERKS_(arg0)
    this := &QTableWidgetSelectionRange{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetSelectionRange)
    return this
  case 2:
    // invoke: _ZN26QTableWidgetSelectionRangeC1Ev
    // invoke: void QTableWidgetSelectionRange()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QTableWidgetSelectionRangeC2Ev()
    this := &QTableWidgetSelectionRange{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetSelectionRange)
    return this
  default:
    qtrt.ErrorResolve("QTableWidgetSelectionRange", "QTableWidgetSelectionRange", args)
  }

  return nil // QTableWidgetSelectionRange{Qclsinst:qthis}
}

// columnCount()
func (this *QTableWidget) ColumnCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget11columnCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "columnCount", args)
  }

  return
}

// removeCellWidget(int, int)
func (this *QTableWidget) RemoveCellWidget(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget16removeCellWidgetEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "removeCellWidget", args)
  }

  return
}

// setCellWidget(int, int, class QWidget *)
func (this *QTableWidget) SetCellWidget(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QTableWidget13setCellWidgetEiiP7QWidget(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTableWidget", "setCellWidget", args)
  }

  return
}

// verticalHeaderItem(int)
func (this *QTableWidget) VerticalHeaderItem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget18verticalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "verticalHeaderItem", args)
  }

  return
}

// setItemPrototype(const class QTableWidgetItem *)
func (this *QTableWidget) SetItemPrototype(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget16setItemPrototypeEPK16QTableWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemPrototype", args)
  }

  return
}

// ~QTableWidget()
func (this *QTableWidget) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN12QTableWidgetD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "~QTableWidget", args)
  }

  return
}

// setItemSelected(const class QTableWidgetItem *, _Bool)
func (this *QTableWidget) SetItemSelected(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget15setItemSelectedEPK16QTableWidgetItemb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setItemSelected", args)
  }

  return
}

// openPersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) OpenPersistentEditor(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget20openPersistentEditorEP16QTableWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "openPersistentEditor", args)
  }

  return
}

// row(const class QTableWidgetItem *)
func (this *QTableWidget) Row(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget3rowEPK16QTableWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "row", args)
  }

  return
}

// setColumnCount(int)
func (this *QTableWidget) SetColumnCount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget14setColumnCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setColumnCount", args)
  }

  return
}

// setCurrentCell(int, int)
func (this *QTableWidget) SetCurrentCell(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget14setCurrentCellEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentCell", args)
  }

  return
}

// QTableWidget(int, int, class QWidget *)
func GcfreeQTableWidget(this *QTableWidget) {
  qtrt.UniverseFree(this)
}
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QWidget).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTableWidgetC2EiiP7QWidget(arg0, arg1, arg2)
    this := &QTableWidget{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidget)
    return this
  case 1:
    // invoke: _ZN12QTableWidgetC1EP7QWidget
    // invoke: void QTableWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTableWidgetC2EP7QWidget(arg0)
    this := &QTableWidget{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidget)
    return this
  default:
    qtrt.ErrorResolve("QTableWidget", "QTableWidget", args)
  }

  return nil // QTableWidget{Qclsinst:qthis}
}

// setRowCount(int)
func (this *QTableWidget) SetRowCount(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget11setRowCountEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setRowCount", args)
  }

  return
}

// editItem(class QTableWidgetItem *)
func (this *QTableWidget) EditItem(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget8editItemEP16QTableWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "editItem", args)
  }

  return
}

// setHorizontalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) SetHorizontalHeaderItem(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget23setHorizontalHeaderItemEiP16QTableWidgetItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderItem", args)
  }

  return
}

// currentItem()
func (this *QTableWidget) CurrentItem(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget11currentItemEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "currentItem", args)
  }

  return
}

// insertColumn(int)
func (this *QTableWidget) InsertColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget12insertColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "insertColumn", args)
  }

  return
}

// isItemSelected(const class QTableWidgetItem *)
func (this *QTableWidget) IsItemSelected(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget14isItemSelectedEPK16QTableWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "isItemSelected", args)
  }

  return
}

// clearContents()
func (this *QTableWidget) ClearContents(args ...interface{}) () {
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
    C.C_ZN12QTableWidget13clearContentsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "clearContents", args)
  }

  return
}

// currentColumn()
func (this *QTableWidget) CurrentColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget13currentColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "currentColumn", args)
  }

  return
}

// itemAt(const class QPoint &)
func (this *QTableWidget) ItemAt(args ...interface{}) (ret interface{}) {
  // itemAt(const class QPoint &)
  // itemAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTableWidget6itemAtERK6QPoint
    // invoke: QTableWidgetItem * itemAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget6itemAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK12QTableWidget6itemAtEii
    // invoke: QTableWidgetItem * itemAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK12QTableWidget6itemAtEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "itemAt", args)
  }

  return
}

// takeItem(int, int)
func (this *QTableWidget) TakeItem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QTableWidget8takeItemEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "takeItem", args)
  }

  return
}

// takeHorizontalHeaderItem(int)
func (this *QTableWidget) TakeHorizontalHeaderItem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QTableWidget24takeHorizontalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "takeHorizontalHeaderItem", args)
  }

  return
}

// isSortingEnabled()
func (this *QTableWidget) IsSortingEnabled(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget16isSortingEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "isSortingEnabled", args)
  }

  return
}

// selectedRanges()
func (this *QTableWidget) SelectedRanges(args ...interface{}) () {
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
    C.C_ZNK12QTableWidget14selectedRangesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedRanges", args)
  }

  return
}

// selectedItems()
func (this *QTableWidget) SelectedItems(args ...interface{}) () {
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
    C.C_ZNK12QTableWidget13selectedItemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "selectedItems", args)
  }

  return
}

// removeRow(int)
func (this *QTableWidget) RemoveRow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget9removeRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "removeRow", args)
  }

  return
}

// rowCount()
func (this *QTableWidget) RowCount(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget8rowCountEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "rowCount", args)
  }

  return
}

// setItem(int, int, class QTableWidgetItem *)
func (this *QTableWidget) SetItem(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg2)}
    C.C_ZN12QTableWidget7setItemEiiP16QTableWidgetItem(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTableWidget", "setItem", args)
  }

  return
}

// takeVerticalHeaderItem(int)
func (this *QTableWidget) TakeVerticalHeaderItem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QTableWidget22takeVerticalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "takeVerticalHeaderItem", args)
  }

  return
}

// horizontalHeaderItem(int)
func (this *QTableWidget) HorizontalHeaderItem(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget20horizontalHeaderItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "horizontalHeaderItem", args)
  }

  return
}

// setHorizontalHeaderLabels(const class QStringList &)
func (this *QTableWidget) SetHorizontalHeaderLabels(args ...interface{}) () {
  // setHorizontalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList
    // invoke: void setHorizontalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget25setHorizontalHeaderLabelsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setHorizontalHeaderLabels", args)
  }

  return
}

// visualItemRect(const class QTableWidgetItem *)
func (this *QTableWidget) VisualItemRect(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget14visualItemRectEPK16QTableWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "visualItemRect", args)
  }

  return
}

// removeColumn(int)
func (this *QTableWidget) RemoveColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget12removeColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "removeColumn", args)
  }

  return
}

// setCurrentItem(class QTableWidgetItem *)
func (this *QTableWidget) SetCurrentItem(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget14setCurrentItemEP16QTableWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setCurrentItem", args)
  }

  return
}

// currentRow()
func (this *QTableWidget) CurrentRow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget10currentRowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "currentRow", args)
  }

  return
}

// metaObject()
func (this *QTableWidget) MetaObject(args ...interface{}) () {
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
    C.C_ZNK12QTableWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "metaObject", args)
  }

  return
}

// itemPrototype()
func (this *QTableWidget) ItemPrototype(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTableWidget13itemPrototypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "const QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "itemPrototype", args)
  }

  return
}

// visualRow(int)
func (this *QTableWidget) VisualRow(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget9visualRowEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "visualRow", args)
  }

  return
}

// column(const class QTableWidgetItem *)
func (this *QTableWidget) Column(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget6columnEPK16QTableWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "column", args)
  }

  return
}

// cellWidget(int, int)
func (this *QTableWidget) CellWidget(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK12QTableWidget10cellWidgetEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "cellWidget", args)
  }

  return
}

// setRangeSelected(const class QTableWidgetSelectionRange &, _Bool)
func (this *QTableWidget) SetRangeSelected(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetSelectionRange).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget16setRangeSelectedERK26QTableWidgetSelectionRangeb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setRangeSelected", args)
  }

  return
}

// setSortingEnabled(_Bool)
func (this *QTableWidget) SetSortingEnabled(args ...interface{}) () {
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
    C.C_ZN12QTableWidget17setSortingEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setSortingEnabled", args)
  }

  return
}

// insertRow(int)
func (this *QTableWidget) InsertRow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget9insertRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "insertRow", args)
  }

  return
}

// item(int, int)
func (this *QTableWidget) Item(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK12QTableWidget4itemEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "item", args)
  }

  return
}

// visualColumn(int)
func (this *QTableWidget) VisualColumn(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK12QTableWidget12visualColumnEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidget", "visualColumn", args)
  }

  return
}

// setVerticalHeaderItem(int, class QTableWidgetItem *)
func (this *QTableWidget) SetVerticalHeaderItem(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QTableWidget21setVerticalHeaderItemEiP16QTableWidgetItem(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderItem", args)
  }

  return
}

// closePersistentEditor(class QTableWidgetItem *)
func (this *QTableWidget) ClosePersistentEditor(args ...interface{}) () {
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
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget21closePersistentEditorEP16QTableWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "closePersistentEditor", args)
  }

  return
}

// clear()
func (this *QTableWidget) Clear(args ...interface{}) () {
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
    C.C_ZN12QTableWidget5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidget", "clear", args)
  }

  return
}

// setVerticalHeaderLabels(const class QStringList &)
func (this *QTableWidget) SetVerticalHeaderLabels(args ...interface{}) () {
  // setVerticalHeaderLabels(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList
    // invoke: void setVerticalHeaderLabels(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QTableWidget23setVerticalHeaderLabelsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidget", "setVerticalHeaderLabels", args)
  }

  return
}

// setTextAlignment(int)
func (this *QTableWidgetItem) SetTextAlignment(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem16setTextAlignmentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextAlignment", args)
  }

  return
}

// setSizeHint(const class QSize &)
func (this *QTableWidgetItem) SetSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem11setSizeHintERK5QSize
    // invoke: void setSizeHint(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem11setSizeHintERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSizeHint", args)
  }

  return
}

// setSelected(_Bool)
func (this *QTableWidgetItem) SetSelected(args ...interface{}) () {
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
    C.C_ZN16QTableWidgetItem11setSelectedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setSelected", args)
  }

  return
}

// text()
func (this *QTableWidgetItem) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "text", args)
  }

  return
}

// QTableWidgetItem(const class QTableWidgetItem &)
func GcfreeQTableWidgetItem(this *QTableWidgetItem) {
  qtrt.UniverseFree(this)
}
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
  vtys[2][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[2][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItemC1ERKS_
    // invoke: void QTableWidgetItem(const class QTableWidgetItem &)
    var arg0 = args[0].(*QTableWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2ERKS_(arg0)
    this := &QTableWidgetItem{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetItem)
    return this
  case 1:
    // invoke: _ZN16QTableWidgetItemC1Ei
    // invoke: void QTableWidgetItem(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2Ei(arg0)
    this := &QTableWidgetItem{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetItem)
    return this
  case 2:
    // invoke: _ZN16QTableWidgetItemC1ERK5QIconRK7QStringi
    // invoke: void QTableWidgetItem(const class QIcon &, const class QString &, int)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2ERK5QIconRK7QStringi(arg0, arg1, arg2)
    this := &QTableWidgetItem{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetItem)
    return this
  case 3:
    // invoke: _ZN16QTableWidgetItemC1ERK7QStringi
    // invoke: void QTableWidgetItem(const class QString &, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN16QTableWidgetItemC2ERK7QStringi(arg0, arg1)
    this := &QTableWidgetItem{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTableWidgetItem)
    return this
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "QTableWidgetItem", args)
  }

  return nil // QTableWidgetItem{Qclsinst:qthis}
}

// font()
func (this *QTableWidgetItem) Font(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "font", args)
  }

  return
}

// row()
func (this *QTableWidgetItem) Row(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem3rowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "row", args)
  }

  return
}

// whatsThis()
func (this *QTableWidgetItem) WhatsThis(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem9whatsThisEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "whatsThis", args)
  }

  return
}

// setForeground(const class QBrush &)
func (this *QTableWidgetItem) SetForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setForegroundERK6QBrush
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem13setForegroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setForeground", args)
  }

  return
}

// isSelected()
func (this *QTableWidgetItem) IsSelected(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem10isSelectedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "isSelected", args)
  }

  return
}

// write(class QDataStream &)
func (this *QTableWidgetItem) Write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QTableWidgetItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(*qtcore.QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK16QTableWidgetItem5writeER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "write", args)
  }

  return
}

// read(class QDataStream &)
func (this *QTableWidgetItem) Read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(*qtcore.QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem4readER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "read", args)
  }

  return
}

// backgroundColor()
func (this *QTableWidgetItem) BackgroundColor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem15backgroundColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "backgroundColor", args)
  }

  return
}

// setBackground(const class QBrush &)
func (this *QTableWidgetItem) SetBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem13setBackgroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackground", args)
  }

  return
}

// textColor()
func (this *QTableWidgetItem) TextColor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem9textColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textColor", args)
  }

  return
}

// type()
func (this *QTableWidgetItem) Type_(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "type", args)
  }

  return
}

// setData(int, const class QVariant &)
func (this *QTableWidgetItem) SetData(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setDataEiRK8QVariant
    // invoke: void setData(int, const class QVariant &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN16QTableWidgetItem7setDataEiRK8QVariant(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setData", args)
  }

  return
}

// foreground()
func (this *QTableWidgetItem) Foreground(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem10foregroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "foreground", args)
  }

  return
}

// statusTip()
func (this *QTableWidgetItem) StatusTip(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem9statusTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "statusTip", args)
  }

  return
}

// clone()
func (this *QTableWidgetItem) Clone(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem5cloneEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidgetItem{}) // "QTableWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "clone", args)
  }

  return
}

// setBackgroundColor(const class QColor &)
func (this *QTableWidgetItem) SetBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem18setBackgroundColorERK6QColor
    // invoke: void setBackgroundColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem18setBackgroundColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setBackgroundColor", args)
  }

  return
}

// toolTip()
func (this *QTableWidgetItem) ToolTip(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem7toolTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "toolTip", args)
  }

  return
}

// tableWidget()
func (this *QTableWidgetItem) TableWidget(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem11tableWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTableWidget{}) // "QTableWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "tableWidget", args)
  }

  return
}

// setWhatsThis(const class QString &)
func (this *QTableWidgetItem) SetWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem12setWhatsThisERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setWhatsThis", args)
  }

  return
}

// background()
func (this *QTableWidgetItem) Background(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem10backgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "background", args)
  }

  return
}

// data(int)
func (this *QTableWidgetItem) Data(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK16QTableWidgetItem4dataEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "data", args)
  }

  return
}

// setTextColor(const class QColor &)
func (this *QTableWidgetItem) SetTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setTextColorERK6QColor
    // invoke: void setTextColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem12setTextColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setTextColor", args)
  }

  return
}

// icon()
func (this *QTableWidgetItem) Icon(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem4iconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "icon", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QTableWidgetItem) SetToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem10setToolTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setToolTip", args)
  }

  return
}

// sizeHint()
func (this *QTableWidgetItem) SizeHint(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "sizeHint", args)
  }

  return
}

// ~QTableWidgetItem()
func (this *QTableWidgetItem) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN16QTableWidgetItemD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "~QTableWidgetItem", args)
  }

  return
}

// column()
func (this *QTableWidgetItem) Column(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem6columnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "column", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QTableWidgetItem) SetIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem7setIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setIcon", args)
  }

  return
}

// setText(const class QString &)
func (this *QTableWidgetItem) SetText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setText", args)
  }

  return
}

// setStatusTip(const class QString &)
func (this *QTableWidgetItem) SetStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem12setStatusTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setStatusTip", args)
  }

  return
}

// flags()
func (this *QTableWidgetItem) Flags(args ...interface{}) () {
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
    C.C_ZNK16QTableWidgetItem5flagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "flags", args)
  }

  return
}

// checkState()
func (this *QTableWidgetItem) CheckState(args ...interface{}) () {
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
    C.C_ZNK16QTableWidgetItem10checkStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "checkState", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QTableWidgetItem) SetFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QTableWidgetItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN16QTableWidgetItem7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "setFont", args)
  }

  return
}

// textAlignment()
func (this *QTableWidgetItem) TextAlignment(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QTableWidgetItem13textAlignmentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTableWidgetItem", "textAlignment", args)
  }

  return
}

// <= body block end

