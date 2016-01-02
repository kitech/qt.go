package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
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
  // proto:  void QListWidgetItem::~QListWidgetItem();
extern void _ZN15QListWidgetItemD0Ev(void* qthis);
  // proto:  bool QListWidgetItem::isHidden();
extern void demth_ZNK15QListWidgetItem8isHiddenEv(void* qthis);
  // proto:  void QListWidgetItem::setData(int role, const QVariant & value);
extern void _ZN15QListWidgetItem7setDataEiRK8QVariant(void* qthis, int arg0, void* arg1);
  // proto:  void QListWidgetItem::setBackground(const QBrush & brush);
extern void demth_ZN15QListWidgetItem13setBackgroundERK6QBrush(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::setSelected(bool select);
extern void demth_ZN15QListWidgetItem11setSelectedEb(void* qthis, bool arg0);
  // proto:  QFont QListWidgetItem::font();
extern void demth_ZNK15QListWidgetItem4fontEv(void* qthis);
  // proto:  void QListWidgetItem::setTextAlignment(int alignment);
extern void demth_ZN15QListWidgetItem16setTextAlignmentEi(void* qthis, int arg0);
  // proto:  void QListWidgetItem::QListWidgetItem(QListWidget * view, int type);
extern void* dector_ZN15QListWidgetItemC1EP11QListWidgeti(void* arg0, int arg1);
extern void _ZN15QListWidgetItemC1EP11QListWidgeti(void* qthis, void* arg0, int arg1);
  // proto:  void QListWidgetItem::write(QDataStream & out);
extern void _ZNK15QListWidgetItem5writeER11QDataStream(void* qthis, void* arg0);
  // proto:  QString QListWidgetItem::whatsThis();
extern void demth_ZNK15QListWidgetItem9whatsThisEv(void* qthis);
  // proto:  int QListWidgetItem::type();
extern void demth_ZNK15QListWidgetItem4typeEv(void* qthis);
  // proto:  void QListWidgetItem::QListWidgetItem(const QIcon & icon, const QString & text, QListWidget * view, int type);
extern void* dector_ZN15QListWidgetItemC1ERK5QIconRK7QStringP11QListWidgeti(void* arg0, void* arg1, void* arg2, int arg3);
extern void _ZN15QListWidgetItemC1ERK5QIconRK7QStringP11QListWidgeti(void* qthis, void* arg0, void* arg1, void* arg2, int arg3);
  // proto:  QIcon QListWidgetItem::icon();
extern void demth_ZNK15QListWidgetItem4iconEv(void* qthis);
  // proto:  QColor QListWidgetItem::textColor();
extern void demth_ZNK15QListWidgetItem9textColorEv(void* qthis);
  // proto:  QBrush QListWidgetItem::foreground();
extern void demth_ZNK15QListWidgetItem10foregroundEv(void* qthis);
  // proto:  QBrush QListWidgetItem::background();
extern void demth_ZNK15QListWidgetItem10backgroundEv(void* qthis);
  // proto:  void QListWidgetItem::setStatusTip(const QString & statusTip);
extern void demth_ZN15QListWidgetItem12setStatusTipERK7QString(void* qthis, void* arg0);
  // proto:  QString QListWidgetItem::text();
extern void demth_ZNK15QListWidgetItem4textEv(void* qthis);
  // proto:  QColor QListWidgetItem::backgroundColor();
extern void demth_ZNK15QListWidgetItem15backgroundColorEv(void* qthis);
  // proto:  bool QListWidgetItem::isSelected();
extern void demth_ZNK15QListWidgetItem10isSelectedEv(void* qthis);
  // proto:  void QListWidgetItem::setFont(const QFont & font);
extern void demth_ZN15QListWidgetItem7setFontERK5QFont(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::setText(const QString & text);
extern void demth_ZN15QListWidgetItem7setTextERK7QString(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::QListWidgetItem(const QString & text, QListWidget * view, int type);
extern void* dector_ZN15QListWidgetItemC1ERK7QStringP11QListWidgeti(void* arg0, void* arg1, int arg2);
extern void _ZN15QListWidgetItemC1ERK7QStringP11QListWidgeti(void* qthis, void* arg0, void* arg1, int arg2);
  // proto:  QVariant QListWidgetItem::data(int role);
extern void _ZNK15QListWidgetItem4dataEi(void* qthis, int arg0);
  // proto:  QSize QListWidgetItem::sizeHint();
extern void demth_ZNK15QListWidgetItem8sizeHintEv(void* qthis);
  // proto:  void QListWidgetItem::setWhatsThis(const QString & whatsThis);
extern void demth_ZN15QListWidgetItem12setWhatsThisERK7QString(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::read(QDataStream & in);
extern void _ZN15QListWidgetItem4readER11QDataStream(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::setTextColor(const QColor & color);
extern void demth_ZN15QListWidgetItem12setTextColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::setSizeHint(const QSize & size);
extern void demth_ZN15QListWidgetItem11setSizeHintERK5QSize(void* qthis, void* arg0);
  // proto:  QListWidget * QListWidgetItem::listWidget();
extern void demth_ZNK15QListWidgetItem10listWidgetEv(void* qthis);
  // proto:  void QListWidgetItem::setIcon(const QIcon & icon);
extern void demth_ZN15QListWidgetItem7setIconERK5QIcon(void* qthis, void* arg0);
  // proto:  QListWidgetItem * QListWidgetItem::clone();
extern void _ZNK15QListWidgetItem5cloneEv(void* qthis);
  // proto:  void QListWidgetItem::setBackgroundColor(const QColor & color);
extern void demth_ZN15QListWidgetItem18setBackgroundColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::setForeground(const QBrush & brush);
extern void demth_ZN15QListWidgetItem13setForegroundERK6QBrush(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::QListWidgetItem(const QListWidgetItem & other);
extern void* dector_ZN15QListWidgetItemC1ERKS_(void* arg0);
extern void _ZN15QListWidgetItemC1ERKS_(void* qthis, void* arg0);
  // proto:  void QListWidgetItem::setHidden(bool hide);
extern void demth_ZN15QListWidgetItem9setHiddenEb(void* qthis, bool arg0);
  // proto:  QString QListWidgetItem::toolTip();
extern void demth_ZNK15QListWidgetItem7toolTipEv(void* qthis);
  // proto:  int QListWidgetItem::textAlignment();
extern void demth_ZNK15QListWidgetItem13textAlignmentEv(void* qthis);
  // proto:  QString QListWidgetItem::statusTip();
extern void demth_ZNK15QListWidgetItem9statusTipEv(void* qthis);
  // proto:  void QListWidgetItem::setToolTip(const QString & toolTip);
extern void demth_ZN15QListWidgetItem10setToolTipERK7QString(void* qthis, void* arg0);
  // proto:  void QListWidget::dropEvent(QDropEvent * event);
extern void _ZN11QListWidget9dropEventEP10QDropEvent(void* qthis, void* arg0);
  // proto:  QWidget * QListWidget::itemWidget(QListWidgetItem * item);
extern void _ZNK11QListWidget10itemWidgetEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::QListWidget(QWidget * parent);
extern void* dector_ZN11QListWidgetC1EP7QWidget(void* arg0);
extern void _ZN11QListWidgetC1EP7QWidget(void* qthis, void* arg0);
  // proto:  int QListWidget::currentRow();
extern void _ZNK11QListWidget10currentRowEv(void* qthis);
  // proto:  QListWidgetItem * QListWidget::item(int row);
extern void _ZNK11QListWidget4itemEi(void* qthis, int arg0);
  // proto:  QListWidgetItem * QListWidget::itemAt(const QPoint & p);
extern void _ZNK11QListWidget6itemAtERK6QPoint(void* qthis, void* arg0);
  // proto:  void QListWidget::insertItem(int row, const QString & label);
extern void _ZN11QListWidget10insertItemEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  int QListWidget::row(const QListWidgetItem * item);
extern void _ZNK11QListWidget3rowEPK15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::openPersistentEditor(QListWidgetItem * item);
extern void _ZN11QListWidget20openPersistentEditorEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::clear();
extern void _ZN11QListWidget5clearEv(void* qthis);
  // proto:  void QListWidget::editItem(QListWidgetItem * item);
extern void _ZN11QListWidget8editItemEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  int QListWidget::count();
extern void _ZNK11QListWidget5countEv(void* qthis);
  // proto:  void QListWidget::setItemHidden(const QListWidgetItem * item, bool hide);
extern void _ZN11QListWidget13setItemHiddenEPK15QListWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  void QListWidget::~QListWidget();
extern void _ZN11QListWidgetD0Ev(void* qthis);
  // proto:  void QListWidget::addItem(QListWidgetItem * item);
extern void demth_ZN11QListWidget7addItemEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  QListWidgetItem * QListWidget::takeItem(int row);
extern void _ZN11QListWidget8takeItemEi(void* qthis, int arg0);
  // proto:  bool QListWidget::isSortingEnabled();
extern void _ZNK11QListWidget16isSortingEnabledEv(void* qthis);
  // proto:  void QListWidget::addItems(const QStringList & labels);
extern void demth_ZN11QListWidget8addItemsERK11QStringList(void* qthis, void* arg0);
  // proto:  QList<QListWidgetItem *> QListWidget::selectedItems();
extern void _ZNK11QListWidget13selectedItemsEv(void* qthis);
  // proto:  const QMetaObject * QListWidget::metaObject();
extern void _ZNK11QListWidget10metaObjectEv(void* qthis);
  // proto:  void QListWidget::setItemSelected(const QListWidgetItem * item, bool select);
extern void _ZN11QListWidget15setItemSelectedEPK15QListWidgetItemb(void* qthis, void* arg0, bool arg1);
  // proto:  void QListWidget::insertItem(int row, QListWidgetItem * item);
extern void _ZN11QListWidget10insertItemEiP15QListWidgetItem(void* qthis, int arg0, void* arg1);
  // proto:  void QListWidget::setCurrentRow(int row);
extern void _ZN11QListWidget13setCurrentRowEi(void* qthis, int arg0);
  // proto:  void QListWidget::setSortingEnabled(bool enable);
extern void _ZN11QListWidget17setSortingEnabledEb(void* qthis, bool arg0);
  // proto:  QRect QListWidget::visualItemRect(const QListWidgetItem * item);
extern void _ZNK11QListWidget14visualItemRectEPK15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::removeItemWidget(QListWidgetItem * item);
extern void demth_ZN11QListWidget16removeItemWidgetEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::closePersistentEditor(QListWidgetItem * item);
extern void _ZN11QListWidget21closePersistentEditorEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  bool QListWidget::isItemHidden(const QListWidgetItem * item);
extern void _ZNK11QListWidget12isItemHiddenEPK15QListWidgetItem(void* qthis, void* arg0);
  // proto:  QListWidgetItem * QListWidget::itemAt(int x, int y);
extern void demth_ZNK11QListWidget6itemAtEii(void* qthis, int arg0, int arg1);
  // proto:  void QListWidget::addItem(const QString & label);
extern void demth_ZN11QListWidget7addItemERK7QString(void* qthis, void* arg0);
  // proto:  void QListWidget::insertItems(int row, const QStringList & labels);
extern void _ZN11QListWidget11insertItemsEiRK11QStringList(void* qthis, int arg0, void* arg1);
  // proto:  QListWidgetItem * QListWidget::currentItem();
extern void _ZNK11QListWidget11currentItemEv(void* qthis);
  // proto:  void QListWidget::setCurrentItem(QListWidgetItem * item);
extern void _ZN11QListWidget14setCurrentItemEP15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::setItemWidget(QListWidgetItem * item, QWidget * widget);
extern void _ZN11QListWidget13setItemWidgetEP15QListWidgetItemP7QWidget(void* qthis, void* arg0, void* arg1);
  // proto:  bool QListWidget::isItemSelected(const QListWidgetItem * item);
extern void _ZNK11QListWidget14isItemSelectedEPK15QListWidgetItem(void* qthis, void* arg0);
  // proto:  void QListWidget::QListWidget(const QListWidget & );
extern void* dector_ZN11QListWidgetC1ERKS_(void* arg0);
extern void _ZN11QListWidgetC1ERKS_(void* qthis, void* arg0);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QListWidget)=1
type QListWidget struct {
  /*qbase*/ QListView;
  qclsinst uint64 /* *mut c_void*/;
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

  // proto:  void QListWidgetItem::~QListWidgetItem();
func (this *QListWidgetItem) FreeQListWidgetItem(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListWidgetItem", "~QListWidgetItem", args)
  }

}

  // proto:  bool QListWidgetItem::isHidden();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isHidden", args)
  }

}

  // proto:  void QListWidgetItem::setData(int role, const QVariant & value);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setData", args)
  }

}

  // proto:  void QListWidgetItem::setBackground(const QBrush & brush);
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
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackground", args)
  }

}

  // proto:  void QListWidgetItem::setSelected(bool select);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSelected", args)
  }

}

  // proto:  QFont QListWidgetItem::font();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "font", args)
  }

}

  // proto:  void QListWidgetItem::setTextAlignment(int alignment);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextAlignment", args)
  }

}

  // proto:  void QListWidgetItem::QListWidgetItem(QListWidget * view, int type);
func NewQListWidgetItem(args ...interface{}) QListWidgetItem {
  return QListWidgetItem{}
}

  // proto:  void QListWidgetItem::write(QDataStream & out);
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
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "write", args)
  }

}

  // proto:  QString QListWidgetItem::whatsThis();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "whatsThis", args)
  }

}

  // proto:  int QListWidgetItem::type();
func (this *QListWidgetItem) type_(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListWidgetItem", "type", args)
  }

}

  // proto:  QIcon QListWidgetItem::icon();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "icon", args)
  }

}

  // proto:  QColor QListWidgetItem::textColor();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textColor", args)
  }

}

  // proto:  QBrush QListWidgetItem::foreground();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "foreground", args)
  }

}

  // proto:  QBrush QListWidgetItem::background();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "background", args)
  }

}

  // proto:  void QListWidgetItem::setStatusTip(const QString & statusTip);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setStatusTip", args)
  }

}

  // proto:  QString QListWidgetItem::text();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "text", args)
  }

}

  // proto:  QColor QListWidgetItem::backgroundColor();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "backgroundColor", args)
  }

}

  // proto:  bool QListWidgetItem::isSelected();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "isSelected", args)
  }

}

  // proto:  void QListWidgetItem::setFont(const QFont & font);
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
    var arg0 = args[0].(QFont).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setFont", args)
  }

}

  // proto:  void QListWidgetItem::setText(const QString & text);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setText", args)
  }

}

  // proto:  QVariant QListWidgetItem::data(int role);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "data", args)
  }

}

  // proto:  QSize QListWidgetItem::sizeHint();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "sizeHint", args)
  }

}

  // proto:  void QListWidgetItem::setWhatsThis(const QString & whatsThis);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setWhatsThis", args)
  }

}

  // proto:  void QListWidgetItem::read(QDataStream & in);
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
    var arg0 = args[0].(QDataStream).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "read", args)
  }

}

  // proto:  void QListWidgetItem::setTextColor(const QColor & color);
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
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setTextColor", args)
  }

}

  // proto:  void QListWidgetItem::setSizeHint(const QSize & size);
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
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setSizeHint", args)
  }

}

  // proto:  QListWidget * QListWidgetItem::listWidget();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "listWidget", args)
  }

}

  // proto:  void QListWidgetItem::setIcon(const QIcon & icon);
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
    var arg0 = args[0].(QIcon).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setIcon", args)
  }

}

  // proto:  QListWidgetItem * QListWidgetItem::clone();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "clone", args)
  }

}

  // proto:  void QListWidgetItem::setBackgroundColor(const QColor & color);
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
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setBackgroundColor", args)
  }

}

  // proto:  void QListWidgetItem::setForeground(const QBrush & brush);
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
    var arg0 = args[0].(QBrush).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setForeground", args)
  }

}

  // proto:  void QListWidgetItem::setHidden(bool hide);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setHidden", args)
  }

}

  // proto:  QString QListWidgetItem::toolTip();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "toolTip", args)
  }

}

  // proto:  int QListWidgetItem::textAlignment();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "textAlignment", args)
  }

}

  // proto:  QString QListWidgetItem::statusTip();
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
  default:
    qtrt.ErrorResolve("QListWidgetItem", "statusTip", args)
  }

}

  // proto:  void QListWidgetItem::setToolTip(const QString & toolTip);
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidgetItem", "setToolTip", args)
  }

}

  // proto:  void QListWidget::dropEvent(QDropEvent * event);
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
    var arg0 = args[0].(QDropEvent).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "dropEvent", args)
  }

}

  // proto:  QWidget * QListWidget::itemWidget(QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "itemWidget", args)
  }

}

  // proto:  void QListWidget::QListWidget(QWidget * parent);
func NewQListWidget(args ...interface{}) QListWidget {
  return QListWidget{}
}

  // proto:  int QListWidget::currentRow();
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
  default:
    qtrt.ErrorResolve("QListWidget", "currentRow", args)
  }

}

  // proto:  QListWidgetItem * QListWidget::item(int row);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "item", args)
  }

}

  // proto:  QListWidgetItem * QListWidget::itemAt(const QPoint & p);
func (this *QListWidget) itemAt(args ...interface{}) () {
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
    // invoke: _ZNK11QListWidget6itemAtERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZNK11QListWidget6itemAtEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidget", "itemAt", args)
  }

}

  // proto:  void QListWidget::insertItem(int row, const QString & label);
func (this *QListWidget) insertItem(args ...interface{}) () {
  // insertItem(int, const class QString &)
  // insertItem(int, class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget10insertItemEiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN11QListWidget10insertItemEiP15QListWidgetItem
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidget", "insertItem", args)
  }

}

  // proto:  int QListWidget::row(const QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "row", args)
  }

}

  // proto:  void QListWidget::openPersistentEditor(QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "openPersistentEditor", args)
  }

}

  // proto:  void QListWidget::clear();
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
  default:
    qtrt.ErrorResolve("QListWidget", "clear", args)
  }

}

  // proto:  void QListWidget::editItem(QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "editItem", args)
  }

}

  // proto:  int QListWidget::count();
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
  default:
    qtrt.ErrorResolve("QListWidget", "count", args)
  }

}

  // proto:  void QListWidget::setItemHidden(const QListWidgetItem * item, bool hide);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidget", "setItemHidden", args)
  }

}

  // proto:  void QListWidget::~QListWidget();
func (this *QListWidget) FreeQListWidget(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QListWidget", "~QListWidget", args)
  }

}

  // proto:  void QListWidget::addItem(QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN11QListWidget7addItemERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "addItem", args)
  }

}

  // proto:  QListWidgetItem * QListWidget::takeItem(int row);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "takeItem", args)
  }

}

  // proto:  bool QListWidget::isSortingEnabled();
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
  default:
    qtrt.ErrorResolve("QListWidget", "isSortingEnabled", args)
  }

}

  // proto:  void QListWidget::addItems(const QStringList & labels);
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
    var arg0 = args[0].(QStringList).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "addItems", args)
  }

}

  // proto:  QList<QListWidgetItem *> QListWidget::selectedItems();
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
  default:
    qtrt.ErrorResolve("QListWidget", "selectedItems", args)
  }

}

  // proto:  const QMetaObject * QListWidget::metaObject();
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
  default:
    qtrt.ErrorResolve("QListWidget", "metaObject", args)
  }

}

  // proto:  void QListWidget::setItemSelected(const QListWidgetItem * item, bool select);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int8_t(args[1].(int8))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidget", "setItemSelected", args)
  }

}

  // proto:  void QListWidget::setCurrentRow(int row);
func (this *QListWidget) setCurrentRow(args ...interface{}) () {
  // setCurrentRow(int, class QItemSelectionModel::SelectionFlags)
  // setCurrentRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget13setCurrentRowEi6QFlagsIN19QItemSelectionModel13SelectionFlagEE
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN11QListWidget13setCurrentRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentRow", args)
  }

}

  // proto:  void QListWidget::setSortingEnabled(bool enable);
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
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "setSortingEnabled", args)
  }

}

  // proto:  QRect QListWidget::visualItemRect(const QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "visualItemRect", args)
  }

}

  // proto:  void QListWidget::removeItemWidget(QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "removeItemWidget", args)
  }

}

  // proto:  void QListWidget::closePersistentEditor(QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "closePersistentEditor", args)
  }

}

  // proto:  bool QListWidget::isItemHidden(const QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "isItemHidden", args)
  }

}

  // proto:  void QListWidget::insertItems(int row, const QStringList & labels);
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QStringList).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidget", "insertItems", args)
  }

}

  // proto:  QListWidgetItem * QListWidget::currentItem();
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
  default:
    qtrt.ErrorResolve("QListWidget", "currentItem", args)
  }

}

  // proto:  void QListWidget::setCurrentItem(QListWidgetItem * item);
func (this *QListWidget) setCurrentItem(args ...interface{}) () {
  // setCurrentItem(class QListWidgetItem *, class QItemSelectionModel::SelectionFlags)
  // setCurrentItem(class QListWidgetItem *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"
  vtys[0][1] = qtrt.Int64Ty(false) // "QItemSelectionModel::SelectionFlags"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QListWidgetItem{}) // "QListWidgetItem *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QListWidget14setCurrentItemEP15QListWidgetItem6QFlagsIN19QItemSelectionModel13SelectionFlagEE
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  case 1:
    // invoke: _ZN11QListWidget14setCurrentItemEP15QListWidgetItem
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "setCurrentItem", args)
  }

}

  // proto:  void QListWidget::setItemWidget(QListWidgetItem * item, QWidget * widget);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QWidget).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QListWidget", "setItemWidget", args)
  }

}

  // proto:  bool QListWidget::isItemSelected(const QListWidgetItem * item);
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
    var arg0 = args[0].(QListWidgetItem).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QListWidget", "isItemSelected", args)
  }

}

// <= body block end

