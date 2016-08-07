package qtwidgets
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
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
  // proto:  void QListWidgetItem::setTextAlignment(int alignment);
extern void C_ZN15QListWidgetItem16setTextAlignmentEi(void* qthis, int32_t arg0); // 2
  // proto:  void QListWidgetItem::setSizeHint(const QSize & size);
extern void C_ZN15QListWidgetItem11setSizeHintERK5QSize(void* qthis, void* arg0); // 2
  // proto:  void QListWidgetItem::setSelected(bool select);
extern void C_ZN15QListWidgetItem11setSelectedEb(void* qthis, bool arg0); // 2
  // proto:  QString QListWidgetItem::text();
extern void* C_ZNK15QListWidgetItem4textEv(void* qthis); // 2
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
extern void* C_ZNK15QListWidgetItem4fontEv(void* qthis); // 2
  // proto:  void QListWidgetItem::write(QDataStream & out);
extern void C_ZNK15QListWidgetItem5writeER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  void QListWidgetItem::setForeground(const QBrush & brush);
extern void C_ZN15QListWidgetItem13setForegroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  void QListWidgetItem::setBackground(const QBrush & brush);
extern void C_ZN15QListWidgetItem13setBackgroundERK6QBrush(void* qthis, void* arg0); // 2
  // proto:  QString QListWidgetItem::whatsThis();
extern void* C_ZNK15QListWidgetItem9whatsThisEv(void* qthis); // 2
  // proto:  bool QListWidgetItem::isSelected();
extern bool C_ZNK15QListWidgetItem10isSelectedEv(void* qthis); // 2
  // proto:  QColor QListWidgetItem::backgroundColor();
extern void* C_ZNK15QListWidgetItem15backgroundColorEv(void* qthis); // 2
  // proto:  void QListWidgetItem::~QListWidgetItem();
extern void C_ZN15QListWidgetItemD2Ev(void* qthis); // 4
  // proto:  QColor QListWidgetItem::textColor();
extern void* C_ZNK15QListWidgetItem9textColorEv(void* qthis); // 2
  // proto:  int QListWidgetItem::type();
extern int32_t C_ZNK15QListWidgetItem4typeEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setData(int role, const QVariant & value);
extern void C_ZN15QListWidgetItem7setDataEiRK8QVariant(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString QListWidgetItem::statusTip();
extern void* C_ZNK15QListWidgetItem9statusTipEv(void* qthis); // 2
  // proto:  QBrush QListWidgetItem::foreground();
extern void* C_ZNK15QListWidgetItem10foregroundEv(void* qthis); // 2
  // proto:  QListWidget * QListWidgetItem::listWidget();
extern void* C_ZNK15QListWidgetItem10listWidgetEv(void* qthis); // 2
  // proto:  void QListWidgetItem::read(QDataStream & in);
extern void C_ZN15QListWidgetItem4readER11QDataStream(void* qthis, void* arg0); // 4
  // proto:  QListWidgetItem * QListWidgetItem::clone();
extern void* C_ZNK15QListWidgetItem5cloneEv(void* qthis); // 4
  // proto:  void QListWidgetItem::setBackgroundColor(const QColor & color);
extern void C_ZN15QListWidgetItem18setBackgroundColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QString QListWidgetItem::toolTip();
extern void* C_ZNK15QListWidgetItem7toolTipEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setWhatsThis(const QString & whatsThis);
extern void C_ZN15QListWidgetItem12setWhatsThisERK7QString(void* qthis, void* arg0); // 2
  // proto:  QBrush QListWidgetItem::background();
extern void* C_ZNK15QListWidgetItem10backgroundEv(void* qthis); // 2
  // proto:  QVariant QListWidgetItem::data(int role);
extern void* C_ZNK15QListWidgetItem4dataEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListWidgetItem::setTextColor(const QColor & color);
extern void C_ZN15QListWidgetItem12setTextColorERK6QColor(void* qthis, void* arg0); // 2
  // proto:  QIcon QListWidgetItem::icon();
extern void* C_ZNK15QListWidgetItem4iconEv(void* qthis); // 2
  // proto:  void QListWidgetItem::setToolTip(const QString & toolTip);
extern void C_ZN15QListWidgetItem10setToolTipERK7QString(void* qthis, void* arg0); // 2
  // proto:  QSize QListWidgetItem::sizeHint();
extern void* C_ZNK15QListWidgetItem8sizeHintEv(void* qthis); // 2
  // proto:  bool QListWidgetItem::isHidden();
extern bool C_ZNK15QListWidgetItem8isHiddenEv(void* qthis); // 2
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
extern int32_t C_ZNK15QListWidgetItem13textAlignmentEv(void* qthis); // 2
  // proto:  void QListWidget::~QListWidget();
extern void C_ZN11QListWidgetD2Ev(void* qthis); // 4
  // proto:  void QListWidget::removeItemWidget(QListWidgetItem * item);
extern void C_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem(void* qthis, void* arg0); // 2
  // proto:  bool QListWidget::isItemHidden(const QListWidgetItem * item);
extern bool C_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::dropEvent(QDropEvent * event);
extern void C_ZN11QListWidget9dropEventEP10QDropEvent(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::setItemSelected(const QListWidgetItem * item, bool select);
extern void C_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
  // proto:  QWidget * QListWidget::itemWidget(QListWidgetItem * item);
extern void* C_ZNK11QListWidget10itemWidgetEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::openPersistentEditor(QListWidgetItem * item);
extern void C_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QListWidget::row(const QListWidgetItem * item);
extern int32_t C_ZNK11QListWidget3rowEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  bool QListWidget::isSortingEnabled();
extern bool C_ZNK11QListWidget16isSortingEnabledEv(void* qthis); // 4
  // proto:  void QListWidget::setItemWidget(QListWidgetItem * item, QWidget * widget);
extern void C_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QListWidget::editItem(QListWidgetItem * item);
extern void C_ZN11QListWidget8editItemEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QListWidgetItem * QListWidget::currentItem();
extern void* C_ZNK11QListWidget11currentItemEv(void* qthis); // 4
  // proto:  void QListWidget::addItems(const QStringList & labels);
extern void C_ZN11QListWidget8addItemsERK11QStringList(void* qthis, void* arg0); // 2
  // proto:  bool QListWidget::isItemSelected(const QListWidgetItem * item);
extern bool C_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::insertItems(int row, const QStringList & labels);
extern void C_ZN11QListWidget11insertItemsEiRK11QStringList(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  void QListWidget::QListWidget(QWidget * parent);
extern void* C_ZN11QListWidgetC2EP7QWidget(void* arg0); // 3
  // proto:  QListWidgetItem * QListWidget::takeItem(int row);
extern void* C_ZN11QListWidget8takeItemEi(void* qthis, int32_t arg0); // 4
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
extern int32_t C_ZNK11QListWidget10currentRowEv(void* qthis); // 4
  // proto:  void QListWidget::setSelectionModel(QItemSelectionModel * selectionModel);
extern void C_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel(void* qthis, void* arg0); // 4
  // proto:  QRect QListWidget::visualItemRect(const QListWidgetItem * item);
extern void* C_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::setCurrentItem(QListWidgetItem * item);
extern void C_ZN11QListWidget14setCurrentItemEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  int QListWidget::count();
extern int32_t C_ZNK11QListWidget5countEv(void* qthis); // 4
  // proto:  const QMetaObject * QListWidget::metaObject();
extern void C_ZNK11QListWidget10metaObjectEv(void* qthis); // 4
  // proto:  void QListWidget::closePersistentEditor(QListWidgetItem * item);
extern void C_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem(void* qthis, void* arg0); // 4
  // proto:  QListWidgetItem * QListWidget::itemAt(int x, int y);
extern void* C_ZNK11QListWidget6itemAtEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QListWidgetItem * QListWidget::itemAt(const QPoint & p);
extern void* C_ZNK11QListWidget6itemAtERK6QPoint(void* qthis, void* arg0); // 4
  // proto:  void QListWidget::clear();
extern void C_ZN11QListWidget5clearEv(void* qthis); // 4
  // proto:  void QListWidget::setSortingEnabled(bool enable);
extern void C_ZN11QListWidget17setSortingEnabledEb(void* qthis, bool arg0); // 4
  // proto:  QListWidgetItem * QListWidget::item(int row);
extern void* C_ZNK11QListWidget4itemEi(void* qthis, int32_t arg0); // 4
  // proto:  void QListWidget::setItemHidden(const QListWidgetItem * item, bool hide);
extern void C_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb(void* qthis, void* arg0, bool arg1); // 4
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
}

// class sizeof(QListWidgetItem)=1
type QListWidgetItem struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QListWidget)=1
type QListWidget struct {
  /*qbase*/ QListView;
  Qclsinst unsafe.Pointer /* *C.void */;
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
func (this *QListWidgetItem) Settextalignment(args ...interface{}) () {
  // setTextAlignment(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem16setTextAlignmentEi
    // invoke: void setTextAlignment(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem16setTextAlignmentEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextAlignment", args)
  }

  return
}

// setSizeHint(const class QSize &)
func (this *QListWidgetItem) Setsizehint(args ...interface{}) () {
  // setSizeHint(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QSize{}) // "const QSize &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem11setSizeHintERK5QSize
    // invoke: void setSizeHint(const class QSize &)
    var arg0 = args[0].(*qtcore.QSize).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem11setSizeHintERK5QSize(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSizeHint", args)
  }

  return
}

// setSelected(_Bool)
func (this *QListWidgetItem) Setselected(args ...interface{}) () {
  // setSelected(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem11setSelectedEb
    // invoke: void setSelected(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem11setSelectedEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSelected", args)
  }

  return
}

// text()
func (this *QListWidgetItem) Text(args ...interface{}) (ret interface{}) {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4textEv
    // invoke: QString text()
    var ret0 = C.C_ZNK15QListWidgetItem4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "text", args)
  }

  return
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
  vtys[1][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QListWidget{}) // "QListWidget *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QListWidget{}) // "QListWidget *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QListWidget{}) // "QListWidget *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItemC1ERKS_
    // invoke: void QListWidgetItem(const class QListWidgetItem &)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2ERKS_(arg0)
    return &QListWidgetItem{Qclsinst:qthis}
  case 1:
    // invoke: _ZN15QListWidgetItemC1ERK5QIconRK7QStringP11QListWidgeti
    // invoke: void QListWidgetItem(const class QIcon &, const class QString &, class QListWidget *, int)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QListWidget).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2ERK5QIconRK7QStringP11QListWidgeti(arg0, arg1, arg2, arg3)
    return &QListWidgetItem{Qclsinst:qthis}
  case 2:
    // invoke: _ZN15QListWidgetItemC1ERK7QStringP11QListWidgeti
    // invoke: void QListWidgetItem(const class QString &, class QListWidget *, int)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QListWidget).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2ERK7QStringP11QListWidgeti(arg0, arg1, arg2)
    return &QListWidgetItem{Qclsinst:qthis}
  case 3:
    // invoke: _ZN15QListWidgetItemC1EP11QListWidgeti
    // invoke: void QListWidgetItem(class QListWidget *, int)
    var arg0 = args[0].(*QListWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN15QListWidgetItemC2EP11QListWidgeti(arg0, arg1)
    return &QListWidgetItem{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "QListWidgetItem", args)
  }

  return nil // QListWidgetItem{Qclsinst:qthis}
}

// setHidden(_Bool)
func (this *QListWidgetItem) Sethidden(args ...interface{}) () {
  // setHidden(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem9setHiddenEb
    // invoke: void setHidden(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem9setHiddenEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setHidden", args)
  }

  return
}

// font()
func (this *QListWidgetItem) Font(args ...interface{}) (ret interface{}) {
  // font()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4fontEv
    // invoke: QFont font()
    var ret0 = C.C_ZNK15QListWidgetItem4fontEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QFont{}) // "QFont"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "font", args)
  }

  return
}

// write(class QDataStream &)
func (this *QListWidgetItem) Write(args ...interface{}) () {
  // write(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDataStream{}) // "QDataStream &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5writeER11QDataStream
    // invoke: void write(class QDataStream &)
    var arg0 = args[0].(*qtcore.QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK15QListWidgetItem5writeER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "write", args)
  }

  return
}

// setForeground(const class QBrush &)
func (this *QListWidgetItem) Setforeground(args ...interface{}) () {
  // setForeground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem13setForegroundERK6QBrush
    // invoke: void setForeground(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem13setForegroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setForeground", args)
  }

  return
}

// setBackground(const class QBrush &)
func (this *QListWidgetItem) Setbackground(args ...interface{}) () {
  // setBackground(const class QBrush &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QBrush{}) // "const QBrush &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem13setBackgroundERK6QBrush
    // invoke: void setBackground(const class QBrush &)
    var arg0 = args[0].(*qtgui.QBrush).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem13setBackgroundERK6QBrush(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackground", args)
  }

  return
}

// whatsThis()
func (this *QListWidgetItem) Whatsthis(args ...interface{}) (ret interface{}) {
  // whatsThis()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9whatsThisEv
    // invoke: QString whatsThis()
    var ret0 = C.C_ZNK15QListWidgetItem9whatsThisEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "whatsThis", args)
  }

  return
}

// isSelected()
func (this *QListWidgetItem) Isselected(args ...interface{}) (ret interface{}) {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10isSelectedEv
    // invoke: bool isSelected()
    var ret0 = C.C_ZNK15QListWidgetItem10isSelectedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isSelected", args)
  }

  return
}

// backgroundColor()
func (this *QListWidgetItem) Backgroundcolor(args ...interface{}) (ret interface{}) {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem15backgroundColorEv
    // invoke: QColor backgroundColor()
    var ret0 = C.C_ZNK15QListWidgetItem15backgroundColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "backgroundColor", args)
  }

  return
}

// ~QListWidgetItem()
func (this *QListWidgetItem) Freeqlistwidgetitem(args ...interface{}) () {
  // ~QListWidgetItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItemD0Ev
    // invoke: void ~QListWidgetItem()
    C.C_ZN15QListWidgetItemD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "~QListWidgetItem", args)
  }

  return
}

// textColor()
func (this *QListWidgetItem) Textcolor(args ...interface{}) (ret interface{}) {
  // textColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9textColorEv
    // invoke: QColor textColor()
    var ret0 = C.C_ZNK15QListWidgetItem9textColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textColor", args)
  }

  return
}

// type()
func (this *QListWidgetItem) Type_(args ...interface{}) (ret interface{}) {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4typeEv
    // invoke: int type()
    var ret0 = C.C_ZNK15QListWidgetItem4typeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "type", args)
  }

  return
}

// setData(int, const class QVariant &)
func (this *QListWidgetItem) Setdata(args ...interface{}) () {
  // setData(int, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setDataEiRK8QVariant
    // invoke: void setData(int, const class QVariant &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN15QListWidgetItem7setDataEiRK8QVariant(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setData", args)
  }

  return
}

// statusTip()
func (this *QListWidgetItem) Statustip(args ...interface{}) (ret interface{}) {
  // statusTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem9statusTipEv
    // invoke: QString statusTip()
    var ret0 = C.C_ZNK15QListWidgetItem9statusTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "statusTip", args)
  }

  return
}

// foreground()
func (this *QListWidgetItem) Foreground(args ...interface{}) (ret interface{}) {
  // foreground()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10foregroundEv
    // invoke: QBrush foreground()
    var ret0 = C.C_ZNK15QListWidgetItem10foregroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "foreground", args)
  }

  return
}

// listWidget()
func (this *QListWidgetItem) Listwidget(args ...interface{}) (ret interface{}) {
  // listWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10listWidgetEv
    // invoke: QListWidget * listWidget()
    var ret0 = C.C_ZNK15QListWidgetItem10listWidgetEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidget{}) // "QListWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "listWidget", args)
  }

  return
}

// read(class QDataStream &)
func (this *QListWidgetItem) Read(args ...interface{}) () {
  // read(class QDataStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QDataStream{}) // "QDataStream &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem4readER11QDataStream
    // invoke: void read(class QDataStream &)
    var arg0 = args[0].(*qtcore.QDataStream).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem4readER11QDataStream(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "read", args)
  }

  return
}

// clone()
func (this *QListWidgetItem) Clone(args ...interface{}) (ret interface{}) {
  // clone()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5cloneEv
    // invoke: QListWidgetItem * clone()
    var ret0 = C.C_ZNK15QListWidgetItem5cloneEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "clone", args)
  }

  return
}

// setBackgroundColor(const class QColor &)
func (this *QListWidgetItem) Setbackgroundcolor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem18setBackgroundColorERK6QColor
    // invoke: void setBackgroundColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem18setBackgroundColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackgroundColor", args)
  }

  return
}

// toolTip()
func (this *QListWidgetItem) Tooltip(args ...interface{}) (ret interface{}) {
  // toolTip()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem7toolTipEv
    // invoke: QString toolTip()
    var ret0 = C.C_ZNK15QListWidgetItem7toolTipEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "toolTip", args)
  }

  return
}

// setWhatsThis(const class QString &)
func (this *QListWidgetItem) Setwhatsthis(args ...interface{}) () {
  // setWhatsThis(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setWhatsThisERK7QString
    // invoke: void setWhatsThis(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem12setWhatsThisERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setWhatsThis", args)
  }

  return
}

// background()
func (this *QListWidgetItem) Background(args ...interface{}) (ret interface{}) {
  // background()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10backgroundEv
    // invoke: QBrush background()
    var ret0 = C.C_ZNK15QListWidgetItem10backgroundEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QBrush{}) // "QBrush"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "background", args)
  }

  return
}

// data(int)
func (this *QListWidgetItem) Data(args ...interface{}) (ret interface{}) {
  // data(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4dataEi
    // invoke: QVariant data(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK15QListWidgetItem4dataEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "data", args)
  }

  return
}

// setTextColor(const class QColor &)
func (this *QListWidgetItem) Settextcolor(args ...interface{}) () {
  // setTextColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QColor{}) // "const QColor &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setTextColorERK6QColor
    // invoke: void setTextColor(const class QColor &)
    var arg0 = args[0].(*qtgui.QColor).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem12setTextColorERK6QColor(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextColor", args)
  }

  return
}

// icon()
func (this *QListWidgetItem) Icon(args ...interface{}) (ret interface{}) {
  // icon()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem4iconEv
    // invoke: QIcon icon()
    var ret0 = C.C_ZNK15QListWidgetItem4iconEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtgui.QIcon{}) // "QIcon"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "icon", args)
  }

  return
}

// setToolTip(const class QString &)
func (this *QListWidgetItem) Settooltip(args ...interface{}) () {
  // setToolTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem10setToolTipERK7QString
    // invoke: void setToolTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem10setToolTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setToolTip", args)
  }

  return
}

// sizeHint()
func (this *QListWidgetItem) Sizehint(args ...interface{}) (ret interface{}) {
  // sizeHint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem8sizeHintEv
    // invoke: QSize sizeHint()
    var ret0 = C.C_ZNK15QListWidgetItem8sizeHintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QSize{}) // "QSize"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "sizeHint", args)
  }

  return
}

// isHidden()
func (this *QListWidgetItem) Ishidden(args ...interface{}) (ret interface{}) {
  // isHidden()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem8isHiddenEv
    // invoke: bool isHidden()
    var ret0 = C.C_ZNK15QListWidgetItem8isHiddenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isHidden", args)
  }

  return
}

// setText(const class QString &)
func (this *QListWidgetItem) Settext(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setText", args)
  }

  return
}

// setStatusTip(const class QString &)
func (this *QListWidgetItem) Setstatustip(args ...interface{}) () {
  // setStatusTip(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem12setStatusTipERK7QString
    // invoke: void setStatusTip(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem12setStatusTipERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setStatusTip", args)
  }

  return
}

// setIcon(const class QIcon &)
func (this *QListWidgetItem) Seticon(args ...interface{}) () {
  // setIcon(const class QIcon &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QIcon{}) // "const QIcon &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setIconERK5QIcon
    // invoke: void setIcon(const class QIcon &)
    var arg0 = args[0].(*qtgui.QIcon).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem7setIconERK5QIcon(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setIcon", args)
  }

  return
}

// flags()
func (this *QListWidgetItem) Flags(args ...interface{}) () {
  // flags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem5flagsEv
    // invoke: Qt::ItemFlags flags()
    C.C_ZNK15QListWidgetItem5flagsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "flags", args)
  }

  return
}

// checkState()
func (this *QListWidgetItem) Checkstate(args ...interface{}) () {
  // checkState()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem10checkStateEv
    // invoke: Qt::CheckState checkState()
    C.C_ZNK15QListWidgetItem10checkStateEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "checkState", args)
  }

  return
}

// setFont(const class QFont &)
func (this *QListWidgetItem) Setfont(args ...interface{}) () {
  // setFont(const class QFont &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QFont{}) // "const QFont &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN15QListWidgetItem7setFontERK5QFont
    // invoke: void setFont(const class QFont &)
    var arg0 = args[0].(*qtgui.QFont).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN15QListWidgetItem7setFontERK5QFont(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setFont", args)
  }

  return
}

// textAlignment()
func (this *QListWidgetItem) Textalignment(args ...interface{}) (ret interface{}) {
  // textAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK15QListWidgetItem13textAlignmentEv
    // invoke: int textAlignment()
    var ret0 = C.C_ZNK15QListWidgetItem13textAlignmentEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textAlignment", args)
  }

  return
}

// ~QListWidget()
func (this *QListWidget) Freeqlistwidget(args ...interface{}) () {
  // ~QListWidget()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidgetD0Ev
    // invoke: void ~QListWidget()
    C.C_ZN11QListWidgetD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "~QListWidget", args)
  }

  return
}

// removeItemWidget(class QListWidgetItem *)
func (this *QListWidget) Removeitemwidget(args ...interface{}) () {
  // removeItemWidget(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget16removeItemWidgetEP15QListWidgetItem
    // invoke: void removeItemWidget(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "removeItemWidget", args)
  }

  return
}

// isItemHidden(const class QListWidgetItem *)
func (this *QListWidget) Isitemhidden(args ...interface{}) (ret interface{}) {
  // isItemHidden(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem
    // invoke: bool isItemHidden(const class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "isItemHidden", args)
  }

  return
}

// dropEvent(class QDropEvent *)
func (this *QListWidget) Dropevent(args ...interface{}) () {
  // dropEvent(class QDropEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtgui.QDropEvent{}) // "QDropEvent *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget9dropEventEP10QDropEvent
    // invoke: void dropEvent(class QDropEvent *)
    var arg0 = args[0].(*qtgui.QDropEvent).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget9dropEventEP10QDropEvent(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "dropEvent", args)
  }

  return
}

// setItemSelected(const class QListWidgetItem *, _Bool)
func (this *QListWidget) Setitemselected(args ...interface{}) () {
  // setItemSelected(const class QListWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb
    // invoke: void setItemSelected(const class QListWidgetItem *, _Bool)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "setItemSelected", args)
  }

  return
}

// itemWidget(class QListWidgetItem *)
func (this *QListWidget) Itemwidget(args ...interface{}) (ret interface{}) {
  // itemWidget(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10itemWidgetEP15QListWidgetItem
    // invoke: QWidget * itemWidget(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget10itemWidgetEP15QListWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWidget{}) // "QWidget *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "itemWidget", args)
  }

  return
}

// openPersistentEditor(class QListWidgetItem *)
func (this *QListWidget) Openpersistenteditor(args ...interface{}) () {
  // openPersistentEditor(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget20openPersistentEditorEP15QListWidgetItem
    // invoke: void openPersistentEditor(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget20openPersistentEditorEP15QListWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "openPersistentEditor", args)
  }

  return
}

// row(const class QListWidgetItem *)
func (this *QListWidget) Row(args ...interface{}) (ret interface{}) {
  // row(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget3rowEPK15QListWidgetItem
    // invoke: int row(const class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget3rowEPK15QListWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "row", args)
  }

  return
}

// isSortingEnabled()
func (this *QListWidget) Issortingenabled(args ...interface{}) (ret interface{}) {
  // isSortingEnabled()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget16isSortingEnabledEv
    // invoke: bool isSortingEnabled()
    var ret0 = C.C_ZNK11QListWidget16isSortingEnabledEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "isSortingEnabled", args)
  }

  return
}

// setItemWidget(class QListWidgetItem *, class QWidget *)
func (this *QListWidget) Setitemwidget(args ...interface{}) () {
  // setItemWidget(class QListWidgetItem *, class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[0][1] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget
    // invoke: void setItemWidget(class QListWidgetItem *, class QWidget *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QWidget).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "setItemWidget", args)
  }

  return
}

// editItem(class QListWidgetItem *)
func (this *QListWidget) Edititem(args ...interface{}) () {
  // editItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8editItemEP15QListWidgetItem
    // invoke: void editItem(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget8editItemEP15QListWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "editItem", args)
  }

  return
}

// currentItem()
func (this *QListWidget) Currentitem(args ...interface{}) (ret interface{}) {
  // currentItem()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget11currentItemEv
    // invoke: QListWidgetItem * currentItem()
    var ret0 = C.C_ZNK11QListWidget11currentItemEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "currentItem", args)
  }

  return
}

// addItems(const class QStringList &)
func (this *QListWidget) Additems(args ...interface{}) () {
  // addItems(const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8addItemsERK11QStringList
    // invoke: void addItems(const class QStringList &)
    var arg0 = args[0].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget8addItemsERK11QStringList(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "addItems", args)
  }

  return
}

// isItemSelected(const class QListWidgetItem *)
func (this *QListWidget) Isitemselected(args ...interface{}) (ret interface{}) {
  // isItemSelected(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem
    // invoke: bool isItemSelected(const class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "isItemSelected", args)
  }

  return
}

// insertItems(int, const class QStringList &)
func (this *QListWidget) Insertitems(args ...interface{}) () {
  // insertItems(int, const class QStringList &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(qtcore.QStringList{}) // "const QStringList &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget11insertItemsEiRK11QStringList
    // invoke: void insertItems(int, const class QStringList &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QStringList).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget11insertItemsEiRK11QStringList(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "insertItems", args)
  }

  return
}

// QListWidget(class QWidget *)
func NewQListWidget(args ...interface{}) *QListWidget {
  // QListWidget(class QWidget *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QWidget{}) // "QWidget *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidgetC1EP7QWidget
    // invoke: void QListWidget(class QWidget *)
    var arg0 = args[0].(*QWidget).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QListWidgetC2EP7QWidget(arg0)
    return &QListWidget{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QListWidget", "QListWidget", args)
  }

  return nil // QListWidget{Qclsinst:qthis}
}

// takeItem(int)
func (this *QListWidget) Takeitem(args ...interface{}) (ret interface{}) {
  // takeItem(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget8takeItemEi
    // invoke: QListWidgetItem * takeItem(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QListWidget8takeItemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "takeItem", args)
  }

  return
}

// insertItem(int, class QListWidgetItem *)
func (this *QListWidget) Insertitem(args ...interface{}) () {
  // insertItem(int, class QListWidgetItem *)
  // insertItem(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget10insertItemEiP15QListWidgetItem
    // invoke: void insertItem(int, class QListWidgetItem *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget10insertItemEiP15QListWidgetItem(this.Qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN11QListWidget10insertItemEiRK7QString
    // invoke: void insertItem(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget10insertItemEiRK7QString(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "insertItem", args)
  }

  return
}

// setCurrentRow(int)
func (this *QListWidget) Setcurrentrow(args ...interface{}) () {
  // setCurrentRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setCurrentRowEi
    // invoke: void setCurrentRow(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget13setCurrentRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentRow", args)
  }

  return
}

// addItem(class QListWidgetItem *)
func (this *QListWidget) Additem(args ...interface{}) () {
  // addItem(class QListWidgetItem *)
  // addItem(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget7addItemEP15QListWidgetItem
    // invoke: void addItem(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget7addItemEP15QListWidgetItem(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN11QListWidget7addItemERK7QString
    // invoke: void addItem(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget7addItemERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "addItem", args)
  }

  return
}

// selectedItems()
func (this *QListWidget) Selecteditems(args ...interface{}) () {
  // selectedItems()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget13selectedItemsEv
    // invoke: QList<QListWidgetItem *> selectedItems()
    C.C_ZNK11QListWidget13selectedItemsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "selectedItems", args)
  }

  return
}

// currentRow()
func (this *QListWidget) Currentrow(args ...interface{}) (ret interface{}) {
  // currentRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10currentRowEv
    // invoke: int currentRow()
    var ret0 = C.C_ZNK11QListWidget10currentRowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "currentRow", args)
  }

  return
}

// setSelectionModel(class QItemSelectionModel *)
func (this *QListWidget) Setselectionmodel(args ...interface{}) () {
  // setSelectionModel(class QItemSelectionModel *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QItemSelectionModel{}) // "QItemSelectionModel *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget17setSelectionModelEP19QItemSelectionModel
    // invoke: void setSelectionModel(class QItemSelectionModel *)
    var arg0 = args[0].(*qtcore.QItemSelectionModel).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget17setSelectionModelEP19QItemSelectionModel(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setSelectionModel", args)
  }

  return
}

// visualItemRect(const class QListWidgetItem *)
func (this *QListWidget) Visualitemrect(args ...interface{}) (ret interface{}) {
  // visualItemRect(const class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget14visualItemRectEPK15QListWidgetItem
    // invoke: QRect visualItemRect(const class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget14visualItemRectEPK15QListWidgetItem(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QRect{}) // "QRect"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "visualItemRect", args)
  }

  return
}

// setCurrentItem(class QListWidgetItem *)
func (this *QListWidget) Setcurrentitem(args ...interface{}) () {
  // setCurrentItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget14setCurrentItemEP15QListWidgetItem
    // invoke: void setCurrentItem(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget14setCurrentItemEP15QListWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentItem", args)
  }

  return
}

// count()
func (this *QListWidget) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK11QListWidget5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "count", args)
  }

  return
}

// metaObject()
func (this *QListWidget) Metaobject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK11QListWidget10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "metaObject", args)
  }

  return
}

// closePersistentEditor(class QListWidgetItem *)
func (this *QListWidget) Closepersistenteditor(args ...interface{}) () {
  // closePersistentEditor(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget21closePersistentEditorEP15QListWidgetItem
    // invoke: void closePersistentEditor(class QListWidgetItem *)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget21closePersistentEditorEP15QListWidgetItem(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "closePersistentEditor", args)
  }

  return
}

// itemAt(int, int)
func (this *QListWidget) Itemat(args ...interface{}) (ret interface{}) {
  // itemAt(int, int)
  // itemAt(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QPoint{}) // "const QPoint &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget6itemAtEii
    // invoke: QListWidgetItem * itemAt(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK11QListWidget6itemAtEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK11QListWidget6itemAtERK6QPoint
    // invoke: QListWidgetItem * itemAt(const class QPoint &)
    var arg0 = args[0].(*qtcore.QPoint).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget6itemAtERK6QPoint(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "itemAt", args)
  }

  return
}

// clear()
func (this *QListWidget) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget5clearEv
    // invoke: void clear()
    C.C_ZN11QListWidget5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QListWidget", "clear", args)
  }

  return
}

// setSortingEnabled(_Bool)
func (this *QListWidget) Setsortingenabled(args ...interface{}) () {
  // setSortingEnabled(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget17setSortingEnabledEb
    // invoke: void setSortingEnabled(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QListWidget17setSortingEnabledEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QListWidget", "setSortingEnabled", args)
  }

  return
}

// item(int)
func (this *QListWidget) Item(args ...interface{}) (ret interface{}) {
  // item(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QListWidget4itemEi
    // invoke: QListWidgetItem * item(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK11QListWidget4itemEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QListWidget", "item", args)
  }

  return
}

// setItemHidden(const class QListWidgetItem *, _Bool)
func (this *QListWidget) Setitemhidden(args ...interface{}) () {
  // setItemHidden(const class QListWidgetItem *, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "const QListWidgetItem *"
  vtys[0][1] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb
    // invoke: void setItemHidden(const class QListWidgetItem *, _Bool)
    var arg0 = args[0].(*QListWidgetItem).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.bool(args[1].(bool))
    if false {fmt.Println(arg1)}
    C.C_ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QListWidget", "setItemHidden", args)
  }

  return
}

// <= body block end

