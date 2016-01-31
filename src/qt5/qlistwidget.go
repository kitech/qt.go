package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtWidgets/qlistwidget.h
// dst-file: /src/widgets/qlistwidget.go
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
  // proto:  void QListWidgetItem::setTextAlignment(int alignment);
extern void C_ZN15QListWidgetItem16setTextAlignmentEi(void* qthis, int32_t arg0); // 2
  // proto:  void QListWidgetItem::setSizeHint(const QSize & size);
extern void C_ZN15QListWidgetItem11setSizeHintERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QListWidgetItem::setSelected(bool select);
extern void C_ZN15QListWidgetItem11setSelectedEb(void* qthis, bool arg0); // 2
  // proto:  QString QListWidgetItem::text();
extern void C_ZNK15QListWidgetItem4textEv(void* qthis); // 2
  // proto:  void QListWidgetItem::QListWidgetItem(const QListWidgetItem & other);
extern void* C_ZN15QListWidgetItemC2ERKS_(void* arg0); // 3
  // proto:  void QListWidgetItem::QListWidgetItem(const QIcon & icon, const QString & text, QListWidget * view, int type);
extern void* C_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti(void* arg0, void* arg1, void* arg2, int32_t arg3); // 3
  // proto:  void QListWidgetItem::QListWidgetItem(const QString & text, QListWidget * view, int type);
extern void* C_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti(void* arg0, void* arg1, int32_t arg2); // 3
  // proto:  void QListWidgetItem::QListWidgetItem(QListWidget * view, int type);
extern void* C_ZN15QListWidgetItemC2EP11QListWidgeti(void* arg0, int32_t arg1); // 3
  // proto:  void QListWidgetItem::setHidden(bool hide);
extern void C_ZN15QListWidgetItem9setHiddenEb(void* qthis, bool arg0); // 2
  // proto:  QFont QListWidgetItem::font();
extern void C_ZNK15QListWidgetItem4fontEv(void* qthis); // 2
  // proto:  void QListWidgetItem::write(QDataStream & out);
extern void C_ZNK15QListWidgetItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QListWidgetItem::setForeground(const QBrush & brush);
extern void C_ZN15QListWidgetItem13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  void QListWidgetItem::setBackground(const QBrush & brush);
extern void C_ZN15QListWidgetItem13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  QString QListWidgetItem::whatsThis();
extern void C_ZNK15QListWidgetItem9whatsThisEv(void* qthis); // 2
  // proto:  bool QListWidgetItem::isSelected();
extern void C_ZNK15QListWidgetItem10isSelectedEv(void* qthis); // 2
  // proto:  QColor QListWidgetItem::backgroundColor();
extern void C_ZNK15QListWidgetItem15backgroundColorEv(void* qthis); // 2
  // proto:  void QListWidgetItem::~QListWidgetItem();
extern void C_ZN15QListWidgetItemD2Ev(void* qthis); // 4
  // proto:  QColor QListWidgetItem::textColor();
extern void C_ZNK15QListWidgetItem9textColorEv(void* qthis); // 2
  // proto:  int QListWidgetItem::type();
extern void C_ZNK15QListWidgetItem4typeEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setData(int role, const QVariant & value);
extern void C_ZN15QListWidgetItem7setDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString QListWidgetItem::statusTip();
extern void C_ZNK15QListWidgetItem9statusTipEv(void* qthis); // 2
  // proto:  QBrush QListWidgetItem::foreground();
extern void C_ZNK15QListWidgetItem10foregroundEv(void* qthis); // 2
  // proto:  QListWidget * QListWidgetItem::listWidget();
extern void C_ZNK15QListWidgetItem10listWidgetEv(void* qthis); // 2
  // proto:  void QListWidgetItem::read(QDataStream & in);
extern void C_ZN15QListWidgetItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  QListWidgetItem * QListWidgetItem::clone();
extern void C_ZNK15QListWidgetItem5cloneEv(void* qthis); // 4
  // proto:  void QListWidgetItem::setBackgroundColor(const QColor & color);
extern void C_ZN15QListWidgetItem18setBackgroundColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QString QListWidgetItem::toolTip();
extern void C_ZNK15QListWidgetItem7toolTipEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setWhatsThis(const QString & whatsThis);
extern void C_ZN15QListWidgetItem12setWhatsThisERK7QString(void* qthis, void* arg0); // 2
  // proto:  QBrush QListWidgetItem::background();
extern void C_ZNK15QListWidgetItem10backgroundEv(void* qthis); // 2
  // proto:  QVariant QListWidgetItem::data(int role);
extern void C_ZNK15QListWidgetItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListWidgetItem::setTextColor(const QColor & color);
extern void C_ZN15QListWidgetItem12setTextColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QIcon QListWidgetItem::icon();
extern void C_ZNK15QListWidgetItem4iconEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setToolTip(const QString & toolTip);
extern void C_ZN15QListWidgetItem10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  QSize QListWidgetItem::sizeHint();
extern void C_ZNK15QListWidgetItem8sizeHintEv(void* qthis); // 2
  // proto:  bool QListWidgetItem::isHidden();
extern void C_ZNK15QListWidgetItem8isHiddenEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setText(const QString & text);
extern void C_ZN15QListWidgetItem7setTextERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QListWidgetItem::setStatusTip(const QString & statusTip);
extern void C_ZN15QListWidgetItem12setStatusTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  void QListWidgetItem::setIcon(const QIcon & icon);
extern void C_ZN15QListWidgetItem7setIconERK5QIcon(void* qthis, void* arg0); // 2
  // proto:  Qt::ItemFlags QListWidgetItem::flags();
extern void C_ZNK15QListWidgetItem5flagsEv(void* qthis); // 2
  // proto:  Qt::CheckState QListWidgetItem::checkState();
extern void C_ZNK15QListWidgetItem10checkStateEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setFont(const QFont & font);
extern void C_ZN15QListWidgetItem7setFontERK5QFont(void* qthis, void* arg0); // 2
  // proto:  int QListWidgetItem::textAlignment();
extern void C_ZNK15QListWidgetItem13textAlignmentEv(void* qthis); // 2
  // proto:  void QListWidget::~QListWidget();
extern void C_ZN11QListWidgetD2Ev(void* qthis); // 4
  // proto:  void QListWidget::removeItemWidget(QListWidgetItem * item);
extern void C_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem(void* qthis, void* arg0); // 2
  // proto:  bool QListWidget::isItemHidden(const QListWidgetItem * item);
extern void C_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::dropEvent(QDropEvent * event);
extern void C_ZN11QListWidget9dropEventEP10QDropEvent(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::setItemSelected(const QListWidgetItem * item, bool select);
extern void C_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QWidget * QListWidget::itemWidget(QListWidgetItem * item);
extern void C_ZNK11QListWidget10itemWidgetEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::openPersistentEditor(QListWidgetItem * item);
extern void C_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QListWidget::row(const QListWidgetItem * item);
extern void C_ZNK11QListWidget3rowEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  bool QListWidget::isSortingEnabled();
extern void C_ZNK11QListWidget16isSortingEnabledEv(void* qthis); // 4
  // proto:  void QListWidget::setItemWidget(QListWidgetItem * item, QWidget * widget);
extern void C_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QListWidget::editItem(QListWidgetItem * item);
extern void C_ZN11QListWidget8editItemEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QListWidgetItem * QListWidget::currentItem();
extern void C_ZNK11QListWidget11currentItemEv(void* qthis); // 4
  // proto:  void QListWidget::addItems(const QStringList & labels);
extern void C_ZN11QListWidget8addItemsERK11QStringList(void* qthis, void* arg0); // 2
  // proto:  bool QListWidget::isItemSelected(const QListWidgetItem * item);
extern void C_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::insertItems(int row, const QStringList & labels);
extern void C_ZN11QListWidget11insertItemsEiRK11QStringList(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QListWidget::QListWidget(QWidget * parent);
extern void* C_ZN11QListWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  QListWidgetItem * QListWidget::takeItem(int row);
extern void C_ZN11QListWidget8takeItemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListWidget::insertItem(int row, QListWidgetItem * item);
extern void C_ZN11QListWidget10insertItemEiP15QListWidgetItem(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QListWidget::insertItem(int row, const QString & label);
extern void C_ZN11QListWidget10insertItemEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QListWidget::setCurrentRow(int row);
extern void C_ZN11QListWidget13setCurrentRowEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListWidget::addItem(QListWidgetItem * item);
extern void C_ZN11QListWidget7addItemEP15QListWidgetItem(void* qthis, void* arg0); // 2
  // proto:  void QListWidget::addItem(const QString & label);
extern void C_ZN11QListWidget7addItemERK7QString(void* qthis, void* arg0); // 2
  // proto:  QList<QListWidgetItem *> QListWidget::selectedItems();
extern void C_ZNK11QListWidget13selectedItemsEv(void* qthis); // 4
  // proto:  int QListWidget::currentRow();
extern void C_ZNK11QListWidget10currentRowEv(void* qthis); // 4
  // proto:  QRect QListWidget::visualItemRect(const QListWidgetItem * item);
extern void C_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::setCurrentItem(QListWidgetItem * item);
extern void C_ZN11QListWidget14setCurrentItemEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QListWidget::count();
extern void C_ZNK11QListWidget5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QListWidget::metaObject();
extern void C_ZNK11QListWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QListWidget::closePersistentEditor(QListWidgetItem * item);
extern void C_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QListWidgetItem * QListWidget::itemAt(int x, int y);
extern void C_ZNK11QListWidget6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QListWidgetItem * QListWidget::itemAt(const QPoint & p);
extern void C_ZNK11QListWidget6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::clear();
extern void C_ZN11QListWidget5clearEv(void* qthis); // 4
  // proto:  void QListWidget::setSortingEnabled(bool enable);
extern void C_ZN11QListWidget17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QListWidgetItem * QListWidget::item(int row);
extern void C_ZNK11QListWidget4itemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListWidget::setItemHidden(const QListWidgetItem * item, bool hide);
extern void C_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
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

// class sizeof(QListWidgetItem)=1
type QListWidgetItem struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QListWidget)=1
type QListWidget struct {
  /*qbase*/ QListView;
  qclsinst unsafe.Pointer /* *C.void */;
//  _itemDoubleClicked QListWidget_itemDoubleClicked_signal;
//  _itemClicked QListWidget_itemClicked_signal;
//  _currentItemChanged QListWidget_currentItemChanged_signal;
//  _itemEntered QListWidget_itemEntered_signal;
//  _itemPressed QListWidget_itemPressed_signal;
//  _itemSelectionChanged QListWidget_itemSelectionChanged_signal;
//  _itemActivated QListWidget_itemActivated_signal;
//  _itemChanged QListWidget_itemChanged_signal;
//  _currentRowChanged QListWidget_currentRowChanged_signal;
//  _currentTextChanged QListWidget_currentTextChanged_signal;
}

// setTextAlignment(int)
func (this *QListWidgetItem) setTextAlignment(args ...interface{}) () {
  // setTextAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem16setTextAlignmentEi
    // invoke: void setTextAlignment(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem16setTextAlignmentEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextAlignment", args)
  }

}

// setSizeHint(const class QSize &)
func (this *QListWidgetItem) setSizeHint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem11setSizeHintERK5QSize
    // invoke: void setSizeHint(const class QSize &)
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem11setSizeHintERK5QSize(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSizeHint", args)
  }

}

// setSelected(_Bool)
func (this *QListWidgetItem) setSelected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem11setSelectedEb
    // invoke: void setSelected(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem11setSelectedEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSelected", args)
  }

}

// text()
func (this *QListWidgetItem) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4textEv
    // invoke: QString text()
    var ret = C.C_ZNK15QListWidgetItem4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "text", args)
  }

}

// QListWidgetItem(const class QListWidgetItem &)
func NewQListWidgetItem(args ...interface{}) *QListWidgetItem {
  // QListWidgetItem(const class QListWidgetItem &)
  // QListWidgetItem(const class QIcon &, const class QString &, class QListWidget *, int)
  // QListWidgetItem(const class QString &, class QListWidget *, int)
  // QListWidgetItem(class QListWidget *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QListWidget{}) // "QListWidget *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QListWidget{}) // "QListWidget *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QListWidget{}) // "QListWidget *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItemC1ERKS_
    // invoke: void QListWidgetItem(const class QListWidgetItem &)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2ERKS_(arg0)
    return &QListWidgetItem{qclsinst:qthis}
  case 1:
    // invoke: _ZN15QListWidgetItemC1ERK5QIconRK7QStringP11QListWidgeti
    // invoke: void QListWidgetItem(const class QIcon &, const class QString &, class QListWidget *, int)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QListWidget).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti(arg0, arg1, arg2, arg3)
    return &QListWidgetItem{qclsinst:qthis}
  case 2:
    // invoke: _ZN15QListWidgetItemC1ERK7QStringP11QListWidgeti
    // invoke: void QListWidgetItem(const class QString &, class QListWidget *, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QListWidget).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti(arg0, arg1, arg2)
    return &QListWidgetItem{qclsinst:qthis}
  case 3:
    // invoke: _ZN15QListWidgetItemC1EP11QListWidgeti
    // invoke: void QListWidgetItem(class QListWidget *, int)
    var arg0 = args[0].(QListWidget).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2EP11QListWidgeti(arg0, arg1)
    return &QListWidgetItem{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "QListWidgetItem", args)
  }

  return nil // QListWidgetItem{qclsinst:qthis}
}

// setHidden(_Bool)
func (this *QListWidgetItem) setHidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem9setHiddenEb
    // invoke: void setHidden(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem9setHiddenEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setHidden", args)
  }

}

// font()
func (this *QListWidgetItem) font(args ...interface{}) () {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4fontEv
    // invoke: QFont font()
    var ret = C.C_ZNK15QListWidgetItem4fontEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "font", args)
  }

}

// write(class QDataStream &)
func (this *QListWidgetItem) write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK15QListWidgetItem5writeER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "write", args)
  }

}

// setForeground(const class QBrush &)
func (this *QListWidgetItem) setForeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem13setForegroundERK6QBrush
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem13setForegroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setForeground", args)
  }

}

// setBackground(const class QBrush &)
func (this *QListWidgetItem) setBackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QBrush{}) // "const QBrush &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem13setBackgroundERK6QBrush(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackground", args)
  }

}

// whatsThis()
func (this *QListWidgetItem) whatsThis(args ...interface{}) () {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9whatsThisEv
    // invoke: QString whatsThis()
    var ret = C.C_ZNK15QListWidgetItem9whatsThisEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "whatsThis", args)
  }

}

// isSelected()
func (this *QListWidgetItem) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10isSelectedEv
    // invoke: bool isSelected()
    var ret = C.C_ZNK15QListWidgetItem10isSelectedEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isSelected", args)
  }

}

// backgroundColor()
func (this *QListWidgetItem) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem15backgroundColorEv
    // invoke: QColor backgroundColor()
    var ret = C.C_ZNK15QListWidgetItem15backgroundColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "backgroundColor", args)
  }

}

// ~QListWidgetItem()
func (this *QListWidgetItem) FreeQListWidgetItem(args ...interface{}) () {
  // ~QListWidgetItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItemD0Ev
    // invoke: void ~QListWidgetItem()
    C.C_ZN15QListWidgetItemD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "~QListWidgetItem", args)
  }

}

// textColor()
func (this *QListWidgetItem) textColor(args ...interface{}) () {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9textColorEv
    // invoke: QColor textColor()
    var ret = C.C_ZNK15QListWidgetItem9textColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textColor", args)
  }

}

// type()
func (this *QListWidgetItem) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4typeEv
    // invoke: int type()
    var ret = C.C_ZNK15QListWidgetItem4typeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "type", args)
  }

}

// setData(int, const class QVariant &)
func (this *QListWidgetItem) setData(args ...interface{}) () {
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
    // invoke: _ZN15QListWidgetItem7setDataEiRK8QVariant
    // invoke: void setData(int, const class QVariant &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QListWidgetItem7setDataEiRK8QVariant(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setData", args)
  }

}

// statusTip()
func (this *QListWidgetItem) statusTip(args ...interface{}) () {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9statusTipEv
    // invoke: QString statusTip()
    var ret = C.C_ZNK15QListWidgetItem9statusTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "statusTip", args)
  }

}

// foreground()
func (this *QListWidgetItem) foreground(args ...interface{}) () {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10foregroundEv
    // invoke: QBrush foreground()
    var ret = C.C_ZNK15QListWidgetItem10foregroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "foreground", args)
  }

}

// listWidget()
func (this *QListWidgetItem) listWidget(args ...interface{}) () {
  // listWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10listWidgetEv
    // invoke: QListWidget * listWidget()
    var ret = C.C_ZNK15QListWidgetItem10listWidgetEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "listWidget", args)
  }

}

// read(class QDataStream &)
func (this *QListWidgetItem) read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDataStream{}) // "QDataStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem4readER11QDataStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "read", args)
  }

}

// clone()
func (this *QListWidgetItem) clone(args ...interface{}) () {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5cloneEv
    // invoke: QListWidgetItem * clone()
    var ret = C.C_ZNK15QListWidgetItem5cloneEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "clone", args)
  }

}

// setBackgroundColor(const class QColor &)
func (this *QListWidgetItem) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem18setBackgroundColorERK6QColor
    // invoke: void setBackgroundColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem18setBackgroundColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackgroundColor", args)
  }

}

// toolTip()
func (this *QListWidgetItem) toolTip(args ...interface{}) () {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem7toolTipEv
    // invoke: QString toolTip()
    var ret = C.C_ZNK15QListWidgetItem7toolTipEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "toolTip", args)
  }

}

// setWhatsThis(const class QString &)
func (this *QListWidgetItem) setWhatsThis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem12setWhatsThisERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setWhatsThis", args)
  }

}

// background()
func (this *QListWidgetItem) background(args ...interface{}) () {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10backgroundEv
    // invoke: QBrush background()
    var ret = C.C_ZNK15QListWidgetItem10backgroundEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "background", args)
  }

}

// data(int)
func (this *QListWidgetItem) data(args ...interface{}) () {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4dataEi
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK15QListWidgetItem4dataEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "data", args)
  }

}

// setTextColor(const class QColor &)
func (this *QListWidgetItem) setTextColor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setTextColorERK6QColor
    // invoke: void setTextColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem12setTextColorERK6QColor(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextColor", args)
  }

}

// icon()
func (this *QListWidgetItem) icon(args ...interface{}) () {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4iconEv
    // invoke: QIcon icon()
    var ret = C.C_ZNK15QListWidgetItem4iconEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "icon", args)
  }

}

// setToolTip(const class QString &)
func (this *QListWidgetItem) setToolTip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem10setToolTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setToolTip", args)
  }

}

// sizeHint()
func (this *QListWidgetItem) sizeHint(args ...interface{}) () {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem8sizeHintEv
    // invoke: QSize sizeHint()
    var ret = C.C_ZNK15QListWidgetItem8sizeHintEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "sizeHint", args)
  }

}

// isHidden()
func (this *QListWidgetItem) isHidden(args ...interface{}) () {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem8isHiddenEv
    // invoke: bool isHidden()
    var ret = C.C_ZNK15QListWidgetItem8isHiddenEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isHidden", args)
  }

}

// setText(const class QString &)
func (this *QListWidgetItem) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setText", args)
  }

}

// setStatusTip(const class QString &)
func (this *QListWidgetItem) setStatusTip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem12setStatusTipERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setStatusTip", args)
  }

}

// setIcon(const class QIcon &)
func (this *QListWidgetItem) setIcon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIcon{}) // "const QIcon &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem7setIconERK5QIcon(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setIcon", args)
  }

}

// flags()
func (this *QListWidgetItem) flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK15QListWidgetItem5flagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "flags", args)
  }

}

// checkState()
func (this *QListWidgetItem) checkState(args ...interface{}) () {
  // checkState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10checkStateEv
    // invoke: Qt::CheckState checkState()
    C.C_ZNK15QListWidgetItem10checkStateEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "checkState", args)
  }

}

// setFont(const class QFont &)
func (this *QListWidgetItem) setFont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QFont{}) // "const QFont &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem7setFontERK5QFont(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setFont", args)
  }

}

// textAlignment()
func (this *QListWidgetItem) textAlignment(args ...interface{}) () {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem13textAlignmentEv
    // invoke: int textAlignment()
    var ret = C.C_ZNK15QListWidgetItem13textAlignmentEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textAlignment", args)
  }

}

// ~QListWidget()
func (this *QListWidget) FreeQListWidget(args ...interface{}) () {
  // ~QListWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidgetD0Ev
    // invoke: void ~QListWidget()
    C.C_ZN11QListWidgetD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "~QListWidget", args)
  }

}

// removeItemWidget(class QListWidgetItem *)
func (this *QListWidget) removeItemWidget(args ...interface{}) () {
  // removeItemWidget(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget16removeItemWidgetEP15QListWidgetItem
    // invoke: void removeItemWidget(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "removeItemWidget", args)
  }

}

// isItemHidden(const class QListWidgetItem *)
func (this *QListWidget) isItemHidden(args ...interface{}) () {
  // isItemHidden(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem
    // invoke: bool isItemHidden(const class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "isItemHidden", args)
  }

}

// dropEvent(class QDropEvent *)
func (this *QListWidget) dropEvent(args ...interface{}) () {
  // dropEvent(class QDropEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QDropEvent{}) // "QDropEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget9dropEventEP10QDropEvent
    // invoke: void dropEvent(class QDropEvent *)
    var arg0 = args[0].(QDropEvent).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget9dropEventEP10QDropEvent(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "dropEvent", args)
  }

}

// setItemSelected(const class QListWidgetItem *, _Bool)
func (this *QListWidget) setItemSelected(args ...interface{}) () {
  // setItemSelected(const class QListWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb
    // invoke: void setItemSelected(const class QListWidgetItem *, _Bool)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "setItemSelected", args)
  }

}

// itemWidget(class QListWidgetItem *)
func (this *QListWidget) itemWidget(args ...interface{}) () {
  // itemWidget(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10itemWidgetEP15QListWidgetItem
    // invoke: QWidget * itemWidget(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget10itemWidgetEP15QListWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "itemWidget", args)
  }

}

// openPersistentEditor(class QListWidgetItem *)
func (this *QListWidget) openPersistentEditor(args ...interface{}) () {
  // openPersistentEditor(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget20openPersistentEditorEP15QListWidgetItem
    // invoke: void openPersistentEditor(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "openPersistentEditor", args)
  }

}

// row(const class QListWidgetItem *)
func (this *QListWidget) row(args ...interface{}) () {
  // row(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget3rowEPK15QListWidgetItem
    // invoke: int row(const class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget3rowEPK15QListWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "row", args)
  }

}

// isSortingEnabled()
func (this *QListWidget) isSortingEnabled(args ...interface{}) () {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget16isSortingEnabledEv
    // invoke: bool isSortingEnabled()
    var ret = C.C_ZNK11QListWidget16isSortingEnabledEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "isSortingEnabled", args)
  }

}

// setItemWidget(class QListWidgetItem *, class QWidget *)
func (this *QListWidget) setItemWidget(args ...interface{}) () {
  // setItemWidget(class QListWidgetItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget
    // invoke: void setItemWidget(class QListWidgetItem *, class QWidget *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "setItemWidget", args)
  }

}

// editItem(class QListWidgetItem *)
func (this *QListWidget) editItem(args ...interface{}) () {
  // editItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8editItemEP15QListWidgetItem
    // invoke: void editItem(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget8editItemEP15QListWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "editItem", args)
  }

}

// currentItem()
func (this *QListWidget) currentItem(args ...interface{}) () {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget11currentItemEv
    // invoke: QListWidgetItem * currentItem()
    var ret = C.C_ZNK11QListWidget11currentItemEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "currentItem", args)
  }

}

// addItems(const class QStringList &)
func (this *QListWidget) addItems(args ...interface{}) () {
  // addItems(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringList{}) // "const QStringList &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8addItemsERK11QStringList
    // invoke: void addItems(const class QStringList &)
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget8addItemsERK11QStringList(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "addItems", args)
  }

}

// isItemSelected(const class QListWidgetItem *)
func (this *QListWidget) isItemSelected(args ...interface{}) () {
  // isItemSelected(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem
    // invoke: bool isItemSelected(const class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "isItemSelected", args)
  }

}

// insertItems(int, const class QStringList &)
func (this *QListWidget) insertItems(args ...interface{}) () {
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
    // invoke: _ZN11QListWidget11insertItemsEiRK11QStringList
    // invoke: void insertItems(int, const class QStringList &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget11insertItemsEiRK11QStringList(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "insertItems", args)
  }

}

// QListWidget(class QWidget *)
func NewQListWidget(args ...interface{}) *QListWidget {
  // QListWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidgetC1EP7QWidget
    // invoke: void QListWidget(class QWidget *)
    var arg0 = args[0].(QWidget).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QListWidgetC2EP7QWidget(arg0)
    return &QListWidget{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QListWidget", "QListWidget", args)
  }

  return nil // QListWidget{qclsinst:qthis}
}

// takeItem(int)
func (this *QListWidget) takeItem(args ...interface{}) () {
  // takeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8takeItemEi
    // invoke: QListWidgetItem * takeItem(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN11QListWidget8takeItemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "takeItem", args)
  }

}

// insertItem(int, class QListWidgetItem *)
func (this *QListWidget) insertItem(args ...interface{}) () {
  // insertItem(int, class QListWidgetItem *)
  // insertItem(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget10insertItemEiP15QListWidgetItem
    // invoke: void insertItem(int, class QListWidgetItem *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget10insertItemEiP15QListWidgetItem(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QListWidget10insertItemEiRK7QString
    // invoke: void insertItem(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget10insertItemEiRK7QString(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "insertItem", args)
  }

}

// setCurrentRow(int)
func (this *QListWidget) setCurrentRow(args ...interface{}) () {
  // setCurrentRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setCurrentRowEi
    // invoke: void setCurrentRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget13setCurrentRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentRow", args)
  }

}

// addItem(class QListWidgetItem *)
func (this *QListWidget) addItem(args ...interface{}) () {
  // addItem(class QListWidgetItem *)
  // addItem(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget7addItemEP15QListWidgetItem
    // invoke: void addItem(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget7addItemEP15QListWidgetItem(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QListWidget7addItemERK7QString
    // invoke: void addItem(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget7addItemERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "addItem", args)
  }

}

// selectedItems()
func (this *QListWidget) selectedItems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget13selectedItemsEv
    // invoke: QList<QListWidgetItem *> selectedItems()
    C.C_ZNK11QListWidget13selectedItemsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "selectedItems", args)
  }

}

// currentRow()
func (this *QListWidget) currentRow(args ...interface{}) () {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10currentRowEv
    // invoke: int currentRow()
    var ret = C.C_ZNK11QListWidget10currentRowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "currentRow", args)
  }

}

// visualItemRect(const class QListWidgetItem *)
func (this *QListWidget) visualItemRect(args ...interface{}) () {
  // visualItemRect(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget14visualItemRectEPK15QListWidgetItem
    // invoke: QRect visualItemRect(const class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "visualItemRect", args)
  }

}

// setCurrentItem(class QListWidgetItem *)
func (this *QListWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget14setCurrentItemEP15QListWidgetItem
    // invoke: void setCurrentItem(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget14setCurrentItemEP15QListWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentItem", args)
  }

}

// count()
func (this *QListWidget) count(args ...interface{}) () {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget5countEv
    // invoke: int count()
    var ret = C.C_ZNK11QListWidget5countEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "count", args)
  }

}

// metaObject()
func (this *QListWidget) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QListWidget10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "metaObject", args)
  }

}

// closePersistentEditor(class QListWidgetItem *)
func (this *QListWidget) closePersistentEditor(args ...interface{}) () {
  // closePersistentEditor(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget21closePersistentEditorEP15QListWidgetItem
    // invoke: void closePersistentEditor(class QListWidgetItem *)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "closePersistentEditor", args)
  }

}

// itemAt(int, int)
func (this *QListWidget) itemAt(args ...interface{}) () {
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
    // invoke: _ZNK11QListWidget6itemAtEii
    // invoke: QListWidgetItem * itemAt(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var ret = C.C_ZNK11QListWidget6itemAtEii(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZNK11QListWidget6itemAtERK6QPoint
    // invoke: QListWidgetItem * itemAt(const class QPoint &)
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget6itemAtERK6QPoint(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "itemAt", args)
  }

}

// clear()
func (this *QListWidget) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget5clearEv
    // invoke: void clear()
    C.C_ZN11QListWidget5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "clear", args)
  }

}

// setSortingEnabled(_Bool)
func (this *QListWidget) setSortingEnabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget17setSortingEnabledEb
    // invoke: void setSortingEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget17setSortingEnabledEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setSortingEnabled", args)
  }

}

// item(int)
func (this *QListWidget) item(args ...interface{}) () {
  // item(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget4itemEi
    // invoke: QListWidgetItem * item(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK11QListWidget4itemEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QListWidget", "item", args)
  }

}

// setItemHidden(const class QListWidgetItem *, _Bool)
func (this *QListWidget) setItemHidden(args ...interface{}) () {
  // setItemHidden(const class QListWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb
    // invoke: void setItemHidden(const class QListWidgetItem *, _Bool)
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "setItemHidden", args)
  }

}

// <= body block end

