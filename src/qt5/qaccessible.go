package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qaccessible.h
// dst-file: /src/gui/qaccessible.go
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
  // proto: static bool QAccessible::isActive();
extern void _ZN11QAccessible8isActiveEv();
  // proto: static Id QAccessible::uniqueId(QAccessibleInterface * iface);
extern void _ZN11QAccessible8uniqueIdEP20QAccessibleInterface(void* arg0);
  // proto: static Id QAccessible::registerAccessibleInterface(QAccessibleInterface * iface);
extern void _ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface(void* arg0);
  // proto: static void QAccessible::setActive(bool active);
extern void _ZN11QAccessible9setActiveEb(bool arg0);
  // proto: static QAccessibleInterface * QAccessible::queryAccessibleInterface(QObject * );
extern void _ZN11QAccessible24queryAccessibleInterfaceEP7QObject(void* arg0);
  // proto: static void QAccessible::updateAccessibility(QAccessibleEvent * event);
extern void _ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent(void* arg0);
  // proto: static void QAccessible::cleanup();
extern void _ZN11QAccessible7cleanupEv();
  // proto: static void QAccessible::setRootObject(QObject * object);
extern void _ZN11QAccessible13setRootObjectEP7QObject(void* arg0);
  // proto: static void QAccessible::deleteAccessibleInterface(Id uniqueId);
extern void _ZN11QAccessible25deleteAccessibleInterfaceEj(unsigned int arg0);
  // proto: static QAccessibleInterface * QAccessible::accessibleInterface(Id uniqueId);
extern void _ZN11QAccessible19accessibleInterfaceEj(unsigned int arg0);
  // proto:  void QAccessible::QAccessible();
extern void* dector_ZN11QAccessibleC1Ev();
extern void _ZN11QAccessibleC1Ev(void* qthis);
  // proto:  void QAccessibleTableModelChangeEvent::setFirstColumn(int col);
extern void _ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(void* qthis, int arg0);
  // proto:  void QAccessibleTableModelChangeEvent::setFirstRow(int row);
extern void _ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableModelChangeEvent::firstRow();
extern void _ZNK32QAccessibleTableModelChangeEvent8firstRowEv(void* qthis);
  // proto:  void QAccessibleTableModelChangeEvent::setLastColumn(int col);
extern void _ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableModelChangeEvent::firstColumn();
extern void _ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(void* qthis);
  // proto:  int QAccessibleTableModelChangeEvent::lastColumn();
extern void _ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(void* qthis);
  // proto:  void QAccessibleTableModelChangeEvent::setLastRow(int row);
extern void _ZN32QAccessibleTableModelChangeEvent10setLastRowEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableModelChangeEvent::lastRow();
extern void _ZNK32QAccessibleTableModelChangeEvent7lastRowEv(void* qthis);
  // proto:  void QAccessibleTextInterface::selection(int selectionIndex, int * startOffset, int * endOffset);
extern void _ZNK24QAccessibleTextInterface9selectionEiPiS0_(void* qthis, int arg0, int* arg1, int* arg2);
  // proto:  void QAccessibleTextInterface::setCursorPosition(int position);
extern void _ZN24QAccessibleTextInterface17setCursorPositionEi(void* qthis, int arg0);
  // proto:  int QAccessibleTextInterface::offsetAtPoint(const QPoint & point);
extern void _ZNK24QAccessibleTextInterface13offsetAtPointERK6QPoint(void* qthis, void* arg0);
  // proto:  QString QAccessibleTextInterface::attributes(int offset, int * startOffset, int * endOffset);
extern void _ZNK24QAccessibleTextInterface10attributesEiPiS0_(void* qthis, int arg0, int* arg1, int* arg2);
  // proto:  int QAccessibleTextInterface::selectionCount();
extern void _ZNK24QAccessibleTextInterface14selectionCountEv(void* qthis);
  // proto:  int QAccessibleTextInterface::characterCount();
extern void _ZNK24QAccessibleTextInterface14characterCountEv(void* qthis);
  // proto:  void QAccessibleTextInterface::~QAccessibleTextInterface();
extern void _ZN24QAccessibleTextInterfaceD0Ev(void* qthis);
  // proto:  QString QAccessibleTextInterface::text(int startOffset, int endOffset);
extern void _ZNK24QAccessibleTextInterface4textEii(void* qthis, int arg0, int arg1);
  // proto:  QRect QAccessibleTextInterface::characterRect(int offset);
extern void _ZNK24QAccessibleTextInterface13characterRectEi(void* qthis, int arg0);
  // proto:  void QAccessibleTextInterface::removeSelection(int selectionIndex);
extern void _ZN24QAccessibleTextInterface15removeSelectionEi(void* qthis, int arg0);
  // proto:  void QAccessibleTextInterface::addSelection(int startOffset, int endOffset);
extern void _ZN24QAccessibleTextInterface12addSelectionEii(void* qthis, int arg0, int arg1);
  // proto:  void QAccessibleTextInterface::scrollToSubstring(int startIndex, int endIndex);
extern void _ZN24QAccessibleTextInterface17scrollToSubstringEii(void* qthis, int arg0, int arg1);
  // proto:  int QAccessibleTextInterface::cursorPosition();
extern void _ZNK24QAccessibleTextInterface14cursorPositionEv(void* qthis);
  // proto:  void QAccessibleTextInterface::setSelection(int selectionIndex, int startOffset, int endOffset);
extern void _ZN24QAccessibleTextInterface12setSelectionEiii(void* qthis, int arg0, int arg1, int arg2);
  // proto:  QObject * QAccessibleEvent::object();
extern void _ZNK16QAccessibleEvent6objectEv(void* qthis);
  // proto:  void QAccessibleEvent::setChild(int chld);
extern void _ZN16QAccessibleEvent8setChildEi(void* qthis, int arg0);
  // proto:  QAccessibleInterface * QAccessibleEvent::accessibleInterface();
extern void _ZNK16QAccessibleEvent19accessibleInterfaceEv(void* qthis);
  // proto:  void QAccessibleEvent::QAccessibleEvent(const QAccessibleEvent & );
extern void* dector_ZN16QAccessibleEventC1ERKS_(void* arg0);
extern void _ZN16QAccessibleEventC1ERKS_(void* qthis, void* arg0);
  // proto:  int QAccessibleEvent::child();
extern void _ZNK16QAccessibleEvent5childEv(void* qthis);
  // proto:  void QAccessibleEvent::~QAccessibleEvent();
extern void _ZN16QAccessibleEventD0Ev(void* qthis);
  // proto: static QString QAccessibleActionInterface::scrollUpAction();
extern void _ZN26QAccessibleActionInterface14scrollUpActionEv();
  // proto:  QStringList QAccessibleActionInterface::actionNames();
extern void _ZNK26QAccessibleActionInterface11actionNamesEv(void* qthis);
  // proto: static const QString & QAccessibleActionInterface::decreaseAction();
extern void _ZN26QAccessibleActionInterface14decreaseActionEv();
  // proto: static const QString & QAccessibleActionInterface::toggleAction();
extern void _ZN26QAccessibleActionInterface12toggleActionEv();
  // proto:  QString QAccessibleActionInterface::localizedActionName(const QString & name);
extern void _ZNK26QAccessibleActionInterface19localizedActionNameERK7QString(void* qthis, void* arg0);
  // proto:  QString QAccessibleActionInterface::localizedActionDescription(const QString & name);
extern void _ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString(void* qthis, void* arg0);
  // proto: static QString QAccessibleActionInterface::scrollLeftAction();
extern void _ZN26QAccessibleActionInterface16scrollLeftActionEv();
  // proto: static QString QAccessibleActionInterface::previousPageAction();
extern void _ZN26QAccessibleActionInterface18previousPageActionEv();
  // proto: static const QString & QAccessibleActionInterface::showMenuAction();
extern void _ZN26QAccessibleActionInterface14showMenuActionEv();
  // proto: static QString QAccessibleActionInterface::scrollRightAction();
extern void _ZN26QAccessibleActionInterface17scrollRightActionEv();
  // proto: static const QString & QAccessibleActionInterface::setFocusAction();
extern void _ZN26QAccessibleActionInterface14setFocusActionEv();
  // proto: static QString QAccessibleActionInterface::nextPageAction();
extern void _ZN26QAccessibleActionInterface14nextPageActionEv();
  // proto:  void QAccessibleActionInterface::~QAccessibleActionInterface();
extern void _ZN26QAccessibleActionInterfaceD0Ev(void* qthis);
  // proto: static const QString & QAccessibleActionInterface::pressAction();
extern void _ZN26QAccessibleActionInterface11pressActionEv();
  // proto:  void QAccessibleActionInterface::doAction(const QString & actionName);
extern void _ZN26QAccessibleActionInterface8doActionERK7QString(void* qthis, void* arg0);
  // proto: static const QString & QAccessibleActionInterface::increaseAction();
extern void _ZN26QAccessibleActionInterface14increaseActionEv();
  // proto:  QStringList QAccessibleActionInterface::keyBindingsForAction(const QString & actionName);
extern void _ZNK26QAccessibleActionInterface20keyBindingsForActionERK7QString(void* qthis, void* arg0);
  // proto: static QString QAccessibleActionInterface::scrollDownAction();
extern void _ZN26QAccessibleActionInterface16scrollDownActionEv();
  // proto:  QAccessibleImageInterface * QAccessibleInterface::imageInterface();
extern void demth_ZN20QAccessibleInterface14imageInterfaceEv(void* qthis);
  // proto:  QAccessibleTableInterface * QAccessibleInterface::tableInterface();
extern void demth_ZN20QAccessibleInterface14tableInterfaceEv(void* qthis);
  // proto:  QAccessibleEditableTextInterface * QAccessibleInterface::editableTextInterface();
extern void demth_ZN20QAccessibleInterface21editableTextInterfaceEv(void* qthis);
  // proto:  QAccessibleValueInterface * QAccessibleInterface::valueInterface();
extern void demth_ZN20QAccessibleInterface14valueInterfaceEv(void* qthis);
  // proto:  QRect QAccessibleInterface::rect();
extern void _ZNK20QAccessibleInterface4rectEv(void* qthis);
  // proto:  QObject * QAccessibleInterface::object();
extern void _ZNK20QAccessibleInterface6objectEv(void* qthis);
  // proto:  QAccessibleActionInterface * QAccessibleInterface::actionInterface();
extern void demth_ZN20QAccessibleInterface15actionInterfaceEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleInterface::parent();
extern void _ZNK20QAccessibleInterface6parentEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleInterface::childAt(int x, int y);
extern void _ZNK20QAccessibleInterface7childAtEii(void* qthis, int arg0, int arg1);
  // proto:  int QAccessibleInterface::childCount();
extern void _ZNK20QAccessibleInterface10childCountEv(void* qthis);
  // proto:  QAccessibleTableCellInterface * QAccessibleInterface::tableCellInterface();
extern void demth_ZN20QAccessibleInterface18tableCellInterfaceEv(void* qthis);
  // proto:  int QAccessibleInterface::indexOfChild(const QAccessibleInterface * );
extern void _ZNK20QAccessibleInterface12indexOfChildEPKS_(void* qthis, void* arg0);
  // proto:  QColor QAccessibleInterface::foregroundColor();
extern void _ZNK20QAccessibleInterface15foregroundColorEv(void* qthis);
  // proto:  bool QAccessibleInterface::isValid();
extern void _ZNK20QAccessibleInterface7isValidEv(void* qthis);
  // proto:  QWindow * QAccessibleInterface::window();
extern void _ZNK20QAccessibleInterface6windowEv(void* qthis);
  // proto:  void QAccessibleInterface::virtual_hook(int id, void * data);
extern void _ZN20QAccessibleInterface12virtual_hookEiPv(void* qthis, int arg0, void* arg1);
  // proto:  QAccessibleInterface * QAccessibleInterface::focusChild();
extern void _ZNK20QAccessibleInterface10focusChildEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleInterface::child(int index);
extern void _ZNK20QAccessibleInterface5childEi(void* qthis, int arg0);
  // proto:  QAccessibleTextInterface * QAccessibleInterface::textInterface();
extern void demth_ZN20QAccessibleInterface13textInterfaceEv(void* qthis);
  // proto:  QColor QAccessibleInterface::backgroundColor();
extern void _ZNK20QAccessibleInterface15backgroundColorEv(void* qthis);
  // proto:  void QAccessibleInterface::~QAccessibleInterface();
extern void _ZN20QAccessibleInterfaceD0Ev(void* qthis);
  // proto:  void QAccessibleEditableTextInterface::insertText(int offset, const QString & text);
extern void _ZN32QAccessibleEditableTextInterface10insertTextEiRK7QString(void* qthis, int arg0, void* arg1);
  // proto:  void QAccessibleEditableTextInterface::replaceText(int startOffset, int endOffset, const QString & text);
extern void _ZN32QAccessibleEditableTextInterface11replaceTextEiiRK7QString(void* qthis, int arg0, int arg1, void* arg2);
  // proto:  void QAccessibleEditableTextInterface::deleteText(int startOffset, int endOffset);
extern void _ZN32QAccessibleEditableTextInterface10deleteTextEii(void* qthis, int arg0, int arg1);
  // proto:  void QAccessibleEditableTextInterface::~QAccessibleEditableTextInterface();
extern void _ZN32QAccessibleEditableTextInterfaceD0Ev(void* qthis);
  // proto:  int QAccessibleTableCellInterface::columnIndex();
extern void _ZNK29QAccessibleTableCellInterface11columnIndexEv(void* qthis);
  // proto:  void QAccessibleTableCellInterface::~QAccessibleTableCellInterface();
extern void _ZN29QAccessibleTableCellInterfaceD0Ev(void* qthis);
  // proto:  int QAccessibleTableCellInterface::columnExtent();
extern void _ZNK29QAccessibleTableCellInterface12columnExtentEv(void* qthis);
  // proto:  int QAccessibleTableCellInterface::rowIndex();
extern void _ZNK29QAccessibleTableCellInterface8rowIndexEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleTableCellInterface::table();
extern void _ZNK29QAccessibleTableCellInterface5tableEv(void* qthis);
  // proto:  int QAccessibleTableCellInterface::rowExtent();
extern void _ZNK29QAccessibleTableCellInterface9rowExtentEv(void* qthis);
  // proto:  QList<QAccessibleInterface *> QAccessibleTableCellInterface::rowHeaderCells();
extern void _ZNK29QAccessibleTableCellInterface14rowHeaderCellsEv(void* qthis);
  // proto:  QList<QAccessibleInterface *> QAccessibleTableCellInterface::columnHeaderCells();
extern void _ZNK29QAccessibleTableCellInterface17columnHeaderCellsEv(void* qthis);
  // proto:  bool QAccessibleTableCellInterface::isSelected();
extern void _ZNK29QAccessibleTableCellInterface10isSelectedEv(void* qthis);
  // proto:  bool QAccessibleTableInterface::unselectColumn(int column);
extern void _ZN25QAccessibleTableInterface14unselectColumnEi(void* qthis, int arg0);
  // proto:  QString QAccessibleTableInterface::columnDescription(int column);
extern void _ZNK25QAccessibleTableInterface17columnDescriptionEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableInterface::selectedCellCount();
extern void _ZNK25QAccessibleTableInterface17selectedCellCountEv(void* qthis);
  // proto:  QList<QAccessibleInterface *> QAccessibleTableInterface::selectedCells();
extern void _ZNK25QAccessibleTableInterface13selectedCellsEv(void* qthis);
  // proto:  bool QAccessibleTableInterface::selectRow(int row);
extern void _ZN25QAccessibleTableInterface9selectRowEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableInterface::selectedRowCount();
extern void _ZNK25QAccessibleTableInterface16selectedRowCountEv(void* qthis);
  // proto:  void QAccessibleTableInterface::~QAccessibleTableInterface();
extern void _ZN25QAccessibleTableInterfaceD0Ev(void* qthis);
  // proto:  QList<int> QAccessibleTableInterface::selectedColumns();
extern void _ZNK25QAccessibleTableInterface15selectedColumnsEv(void* qthis);
  // proto:  QAccessibleInterface * QAccessibleTableInterface::cellAt(int row, int column);
extern void _ZNK25QAccessibleTableInterface6cellAtEii(void* qthis, int arg0, int arg1);
  // proto:  QList<int> QAccessibleTableInterface::selectedRows();
extern void _ZNK25QAccessibleTableInterface12selectedRowsEv(void* qthis);
  // proto:  void QAccessibleTableInterface::modelChange(QAccessibleTableModelChangeEvent * event);
extern void _ZN25QAccessibleTableInterface11modelChangeEP32QAccessibleTableModelChangeEvent(void* qthis, void* arg0);
  // proto:  int QAccessibleTableInterface::columnCount();
extern void _ZNK25QAccessibleTableInterface11columnCountEv(void* qthis);
  // proto:  bool QAccessibleTableInterface::selectColumn(int column);
extern void _ZN25QAccessibleTableInterface12selectColumnEi(void* qthis, int arg0);
  // proto:  bool QAccessibleTableInterface::unselectRow(int row);
extern void _ZN25QAccessibleTableInterface11unselectRowEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableInterface::rowCount();
extern void _ZNK25QAccessibleTableInterface8rowCountEv(void* qthis);
  // proto:  QString QAccessibleTableInterface::rowDescription(int row);
extern void _ZNK25QAccessibleTableInterface14rowDescriptionEi(void* qthis, int arg0);
  // proto:  QAccessibleInterface * QAccessibleTableInterface::summary();
extern void _ZNK25QAccessibleTableInterface7summaryEv(void* qthis);
  // proto:  bool QAccessibleTableInterface::isColumnSelected(int column);
extern void _ZNK25QAccessibleTableInterface16isColumnSelectedEi(void* qthis, int arg0);
  // proto:  QAccessibleInterface * QAccessibleTableInterface::caption();
extern void _ZNK25QAccessibleTableInterface7captionEv(void* qthis);
  // proto:  bool QAccessibleTableInterface::isRowSelected(int row);
extern void _ZNK25QAccessibleTableInterface13isRowSelectedEi(void* qthis, int arg0);
  // proto:  int QAccessibleTableInterface::selectedColumnCount();
extern void _ZNK25QAccessibleTableInterface19selectedColumnCountEv(void* qthis);
  // proto:  QString QAccessibleTextUpdateEvent::textInserted();
extern void _ZNK26QAccessibleTextUpdateEvent12textInsertedEv(void* qthis);
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QAccessibleInterface * iface, int position, const QString & oldText, const QString & text);
extern void* dector_ZN26QAccessibleTextUpdateEventC1EP20QAccessibleInterfaceiRK7QStringS4_(void* arg0, int arg1, void* arg2, void* arg3);
extern void demth_ZN26QAccessibleTextUpdateEventC1EP20QAccessibleInterfaceiRK7QStringS4_(void* qthis, void* arg0, int arg1, void* arg2, void* arg3);
  // proto:  QString QAccessibleTextUpdateEvent::textRemoved();
extern void _ZNK26QAccessibleTextUpdateEvent11textRemovedEv(void* qthis);
  // proto:  int QAccessibleTextUpdateEvent::changePosition();
extern void _ZNK26QAccessibleTextUpdateEvent14changePositionEv(void* qthis);
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QObject * obj, int position, const QString & oldText, const QString & text);
extern void* dector_ZN26QAccessibleTextUpdateEventC1EP7QObjectiRK7QStringS4_(void* arg0, int arg1, void* arg2, void* arg3);
extern void demth_ZN26QAccessibleTextUpdateEventC1EP7QObjectiRK7QStringS4_(void* qthis, void* arg0, int arg1, void* arg2, void* arg3);
  // proto:  QString QAccessibleImageInterface::imageDescription();
extern void _ZNK25QAccessibleImageInterface16imageDescriptionEv(void* qthis);
  // proto:  QPoint QAccessibleImageInterface::imagePosition();
extern void _ZNK25QAccessibleImageInterface13imagePositionEv(void* qthis);
  // proto:  void QAccessibleImageInterface::~QAccessibleImageInterface();
extern void _ZN25QAccessibleImageInterfaceD0Ev(void* qthis);
  // proto:  QSize QAccessibleImageInterface::imageSize();
extern void _ZNK25QAccessibleImageInterface9imageSizeEv(void* qthis);
  // proto:  QString QAccessibleTextInsertEvent::textInserted();
extern void _ZNK26QAccessibleTextInsertEvent12textInsertedEv(void* qthis);
  // proto:  int QAccessibleTextInsertEvent::changePosition();
extern void _ZNK26QAccessibleTextInsertEvent14changePositionEv(void* qthis);
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextInsertEventC1EP20QAccessibleInterfaceiRK7QString(void* arg0, int arg1, void* arg2);
extern void demth_ZN26QAccessibleTextInsertEventC1EP20QAccessibleInterfaceiRK7QString(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QObject * obj, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextInsertEventC1EP7QObjectiRK7QString(void* arg0, int arg1, void* arg2);
extern void demth_ZN26QAccessibleTextInsertEventC1EP7QObjectiRK7QString(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  QVariant QAccessibleValueInterface::maximumValue();
extern void _ZNK25QAccessibleValueInterface12maximumValueEv(void* qthis);
  // proto:  QVariant QAccessibleValueInterface::minimumStepSize();
extern void _ZNK25QAccessibleValueInterface15minimumStepSizeEv(void* qthis);
  // proto:  QVariant QAccessibleValueInterface::currentValue();
extern void _ZNK25QAccessibleValueInterface12currentValueEv(void* qthis);
  // proto:  QVariant QAccessibleValueInterface::minimumValue();
extern void _ZNK25QAccessibleValueInterface12minimumValueEv(void* qthis);
  // proto:  void QAccessibleValueInterface::~QAccessibleValueInterface();
extern void _ZN25QAccessibleValueInterfaceD0Ev(void* qthis);
  // proto:  void QAccessibleValueInterface::setCurrentValue(const QVariant & value);
extern void _ZN25QAccessibleValueInterface15setCurrentValueERK8QVariant(void* qthis, void* arg0);
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QObject * obj, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextRemoveEventC1EP7QObjectiRK7QString(void* arg0, int arg1, void* arg2);
extern void demth_ZN26QAccessibleTextRemoveEventC1EP7QObjectiRK7QString(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  QString QAccessibleTextRemoveEvent::textRemoved();
extern void _ZNK26QAccessibleTextRemoveEvent11textRemovedEv(void* qthis);
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextRemoveEventC1EP20QAccessibleInterfaceiRK7QString(void* arg0, int arg1, void* arg2);
extern void demth_ZN26QAccessibleTextRemoveEventC1EP20QAccessibleInterfaceiRK7QString(void* qthis, void* arg0, int arg1, void* arg2);
  // proto:  int QAccessibleTextRemoveEvent::changePosition();
extern void _ZNK26QAccessibleTextRemoveEvent14changePositionEv(void* qthis);
  // proto:  int QAccessibleTextSelectionEvent::selectionEnd();
extern void _ZNK29QAccessibleTextSelectionEvent12selectionEndEv(void* qthis);
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QAccessibleInterface * iface, int start, int end);
extern void* dector_ZN29QAccessibleTextSelectionEventC1EP20QAccessibleInterfaceii(void* arg0, int arg1, int arg2);
extern void demth_ZN29QAccessibleTextSelectionEventC1EP20QAccessibleInterfaceii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  int QAccessibleTextSelectionEvent::selectionStart();
extern void _ZNK29QAccessibleTextSelectionEvent14selectionStartEv(void* qthis);
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QObject * obj, int start, int end);
extern void* dector_ZN29QAccessibleTextSelectionEventC1EP7QObjectii(void* arg0, int arg1, int arg2);
extern void demth_ZN29QAccessibleTextSelectionEventC1EP7QObjectii(void* qthis, void* arg0, int arg1, int arg2);
  // proto:  void QAccessibleTextSelectionEvent::setSelection(int start, int end);
extern void _ZN29QAccessibleTextSelectionEvent12setSelectionEii(void* qthis, int arg0, int arg1);
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QAccessibleInterface * iface, int cursorPos);
extern void* dector_ZN26QAccessibleTextCursorEventC1EP20QAccessibleInterfacei(void* arg0, int arg1);
extern void demth_ZN26QAccessibleTextCursorEventC1EP20QAccessibleInterfacei(void* qthis, void* arg0, int arg1);
  // proto:  void QAccessibleTextCursorEvent::setCursorPosition(int position);
extern void _ZN26QAccessibleTextCursorEvent17setCursorPositionEi(void* qthis, int arg0);
  // proto:  int QAccessibleTextCursorEvent::cursorPosition();
extern void _ZNK26QAccessibleTextCursorEvent14cursorPositionEv(void* qthis);
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QObject * obj, int cursorPos);
extern void* dector_ZN26QAccessibleTextCursorEventC1EP7QObjecti(void* arg0, int arg1);
extern void demth_ZN26QAccessibleTextCursorEventC1EP7QObjecti(void* qthis, void* arg0, int arg1);
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QObject * obj, const QVariant & val);
extern void* dector_ZN27QAccessibleValueChangeEventC1EP7QObjectRK8QVariant(void* arg0, void* arg1);
extern void demth_ZN27QAccessibleValueChangeEventC1EP7QObjectRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QAccessibleInterface * iface, const QVariant & val);
extern void* dector_ZN27QAccessibleValueChangeEventC1EP20QAccessibleInterfaceRK8QVariant(void* arg0, void* arg1);
extern void demth_ZN27QAccessibleValueChangeEventC1EP20QAccessibleInterfaceRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  void QAccessibleValueChangeEvent::setValue(const QVariant & val);
extern void _ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(void* qthis, void* arg0);
  // proto:  QVariant QAccessibleValueChangeEvent::value();
extern void _ZNK27QAccessibleValueChangeEvent5valueEv(void* qthis);
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

// class sizeof(QAccessible)=1
type QAccessible struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTableModelChangeEvent)=48
type QAccessibleTableModelChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextInterface)=8
type QAccessibleTextInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleEvent)=32
type QAccessibleEvent struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleActionInterface)=8
type QAccessibleActionInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleInterface)=8
type QAccessibleInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleEditableTextInterface)=8
type QAccessibleEditableTextInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTableCellInterface)=8
type QAccessibleTableCellInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTableInterface)=8
type QAccessibleTableInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextUpdateEvent)=56
type QAccessibleTextUpdateEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleStateChangeEvent)=40
type QAccessibleStateChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleImageInterface)=8
type QAccessibleImageInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextInsertEvent)=48
type QAccessibleTextInsertEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleValueInterface)=8
type QAccessibleValueInterface struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextRemoveEvent)=48
type QAccessibleTextRemoveEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextSelectionEvent)=40
type QAccessibleTextSelectionEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleTextCursorEvent)=32
type QAccessibleTextCursorEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QAccessibleValueChangeEvent)=48
type QAccessibleValueChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto: static bool QAccessible::isActive();
func (this *QAccessible) isActive_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "isActive", args)
  }

}

  // proto: static Id QAccessible::uniqueId(QAccessibleInterface * iface);
func (this *QAccessible) uniqueId_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "uniqueId", args)
  }

}

  // proto: static Id QAccessible::registerAccessibleInterface(QAccessibleInterface * iface);
func (this *QAccessible) registerAccessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "registerAccessibleInterface", args)
  }

}

  // proto: static void QAccessible::setActive(bool active);
func (this *QAccessible) setActive_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "setActive", args)
  }

}

  // proto: static QAccessibleInterface * QAccessible::queryAccessibleInterface(QObject * );
func (this *QAccessible) queryAccessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "queryAccessibleInterface", args)
  }

}

  // proto: static void QAccessible::updateAccessibility(QAccessibleEvent * event);
func (this *QAccessible) updateAccessibility_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "updateAccessibility", args)
  }

}

  // proto: static void QAccessible::cleanup();
func (this *QAccessible) cleanup_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "cleanup", args)
  }

}

  // proto: static void QAccessible::setRootObject(QObject * object);
func (this *QAccessible) setRootObject_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "setRootObject", args)
  }

}

  // proto: static void QAccessible::deleteAccessibleInterface(Id uniqueId);
func (this *QAccessible) deleteAccessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "deleteAccessibleInterface", args)
  }

}

  // proto: static QAccessibleInterface * QAccessible::accessibleInterface(Id uniqueId);
func (this *QAccessible) accessibleInterface_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessible", "accessibleInterface", args)
  }

}

  // proto:  void QAccessible::QAccessible();
func NewQAccessible(args ...interface{}) QAccessible {
  return QAccessible{}
}

  // proto:  void QAccessibleTableModelChangeEvent::setFirstColumn(int col);
func (this *QAccessibleTableModelChangeEvent) setFirstColumn(args ...interface{}) () {
  // setFirstColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstColumn", args)
  }

}

  // proto:  void QAccessibleTableModelChangeEvent::setFirstRow(int row);
func (this *QAccessibleTableModelChangeEvent) setFirstRow(args ...interface{}) () {
  // setFirstRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent11setFirstRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstRow", args)
  }

}

  // proto:  int QAccessibleTableModelChangeEvent::firstRow();
func (this *QAccessibleTableModelChangeEvent) firstRow(args ...interface{}) () {
  // firstRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent8firstRowEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstRow", args)
  }

}

  // proto:  void QAccessibleTableModelChangeEvent::setLastColumn(int col);
func (this *QAccessibleTableModelChangeEvent) setLastColumn(args ...interface{}) () {
  // setLastColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent13setLastColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastColumn", args)
  }

}

  // proto:  int QAccessibleTableModelChangeEvent::firstColumn();
func (this *QAccessibleTableModelChangeEvent) firstColumn(args ...interface{}) () {
  // firstColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent11firstColumnEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstColumn", args)
  }

}

  // proto:  int QAccessibleTableModelChangeEvent::lastColumn();
func (this *QAccessibleTableModelChangeEvent) lastColumn(args ...interface{}) () {
  // lastColumn()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent10lastColumnEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastColumn", args)
  }

}

  // proto:  void QAccessibleTableModelChangeEvent::setLastRow(int row);
func (this *QAccessibleTableModelChangeEvent) setLastRow(args ...interface{}) () {
  // setLastRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleTableModelChangeEvent10setLastRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastRow", args)
  }

}

  // proto:  int QAccessibleTableModelChangeEvent::lastRow();
func (this *QAccessibleTableModelChangeEvent) lastRow(args ...interface{}) () {
  // lastRow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent7lastRowEv
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastRow", args)
  }

}

  // proto:  void QAccessibleTextInterface::selection(int selectionIndex, int * startOffset, int * endOffset);
func (this *QAccessibleTextInterface) selection(args ...interface{}) () {
  // selection(int, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface9selectionEiPiS0_
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "selection", args)
  }

}

  // proto:  void QAccessibleTextInterface::setCursorPosition(int position);
func (this *QAccessibleTextInterface) setCursorPosition(args ...interface{}) () {
  // setCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface17setCursorPositionEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "setCursorPosition", args)
  }

}

  // proto:  int QAccessibleTextInterface::offsetAtPoint(const QPoint & point);
func (this *QAccessibleTextInterface) offsetAtPoint(args ...interface{}) () {
  // offsetAtPoint(const class QPoint &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPoint{}) // "const QPoint &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface13offsetAtPointERK6QPoint
    var arg0 = args[0].(QPoint).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "offsetAtPoint", args)
  }

}

  // proto:  QString QAccessibleTextInterface::attributes(int offset, int * startOffset, int * endOffset);
func (this *QAccessibleTextInterface) attributes(args ...interface{}) () {
  // attributes(int, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface10attributesEiPiS0_
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "attributes", args)
  }

}

  // proto:  int QAccessibleTextInterface::selectionCount();
func (this *QAccessibleTextInterface) selectionCount(args ...interface{}) () {
  // selectionCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface14selectionCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "selectionCount", args)
  }

}

  // proto:  int QAccessibleTextInterface::characterCount();
func (this *QAccessibleTextInterface) characterCount(args ...interface{}) () {
  // characterCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface14characterCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "characterCount", args)
  }

}

  // proto:  void QAccessibleTextInterface::~QAccessibleTextInterface();
func (this *QAccessibleTextInterface) FreeQAccessibleTextInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "~QAccessibleTextInterface", args)
  }

}

  // proto:  QString QAccessibleTextInterface::text(int startOffset, int endOffset);
func (this *QAccessibleTextInterface) text(args ...interface{}) () {
  // text(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface4textEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "text", args)
  }

}

  // proto:  QRect QAccessibleTextInterface::characterRect(int offset);
func (this *QAccessibleTextInterface) characterRect(args ...interface{}) () {
  // characterRect(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface13characterRectEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "characterRect", args)
  }

}

  // proto:  void QAccessibleTextInterface::removeSelection(int selectionIndex);
func (this *QAccessibleTextInterface) removeSelection(args ...interface{}) () {
  // removeSelection(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface15removeSelectionEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "removeSelection", args)
  }

}

  // proto:  void QAccessibleTextInterface::addSelection(int startOffset, int endOffset);
func (this *QAccessibleTextInterface) addSelection(args ...interface{}) () {
  // addSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface12addSelectionEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "addSelection", args)
  }

}

  // proto:  void QAccessibleTextInterface::scrollToSubstring(int startIndex, int endIndex);
func (this *QAccessibleTextInterface) scrollToSubstring(args ...interface{}) () {
  // scrollToSubstring(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface17scrollToSubstringEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "scrollToSubstring", args)
  }

}

  // proto:  int QAccessibleTextInterface::cursorPosition();
func (this *QAccessibleTextInterface) cursorPosition(args ...interface{}) () {
  // cursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK24QAccessibleTextInterface14cursorPositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "cursorPosition", args)
  }

}

  // proto:  void QAccessibleTextInterface::setSelection(int selectionIndex, int startOffset, int endOffset);
func (this *QAccessibleTextInterface) setSelection(args ...interface{}) () {
  // setSelection(int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN24QAccessibleTextInterface12setSelectionEiii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QAccessibleTextInterface", "setSelection", args)
  }

}

  // proto:  QObject * QAccessibleEvent::object();
func (this *QAccessibleEvent) object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent6objectEv
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "object", args)
  }

}

  // proto:  void QAccessibleEvent::setChild(int chld);
func (this *QAccessibleEvent) setChild(args ...interface{}) () {
  // setChild(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAccessibleEvent8setChildEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "setChild", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleEvent::accessibleInterface();
func (this *QAccessibleEvent) accessibleInterface(args ...interface{}) () {
  // accessibleInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent19accessibleInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "accessibleInterface", args)
  }

}

  // proto:  void QAccessibleEvent::QAccessibleEvent(const QAccessibleEvent & );
func NewQAccessibleEvent(args ...interface{}) QAccessibleEvent {
  return QAccessibleEvent{}
}

  // proto:  int QAccessibleEvent::child();
func (this *QAccessibleEvent) child(args ...interface{}) () {
  // child()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent5childEv
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "child", args)
  }

}

  // proto:  void QAccessibleEvent::~QAccessibleEvent();
func (this *QAccessibleEvent) FreeQAccessibleEvent(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "~QAccessibleEvent", args)
  }

}

  // proto: static QString QAccessibleActionInterface::scrollUpAction();
func (this *QAccessibleActionInterface) scrollUpAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollUpAction", args)
  }

}

  // proto:  QStringList QAccessibleActionInterface::actionNames();
func (this *QAccessibleActionInterface) actionNames(args ...interface{}) () {
  // actionNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface11actionNamesEv
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "actionNames", args)
  }

}

  // proto: static const QString & QAccessibleActionInterface::decreaseAction();
func (this *QAccessibleActionInterface) decreaseAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "decreaseAction", args)
  }

}

  // proto: static const QString & QAccessibleActionInterface::toggleAction();
func (this *QAccessibleActionInterface) toggleAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "toggleAction", args)
  }

}

  // proto:  QString QAccessibleActionInterface::localizedActionName(const QString & name);
func (this *QAccessibleActionInterface) localizedActionName(args ...interface{}) () {
  // localizedActionName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface19localizedActionNameERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionName", args)
  }

}

  // proto:  QString QAccessibleActionInterface::localizedActionDescription(const QString & name);
func (this *QAccessibleActionInterface) localizedActionDescription(args ...interface{}) () {
  // localizedActionDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionDescription", args)
  }

}

  // proto: static QString QAccessibleActionInterface::scrollLeftAction();
func (this *QAccessibleActionInterface) scrollLeftAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollLeftAction", args)
  }

}

  // proto: static QString QAccessibleActionInterface::previousPageAction();
func (this *QAccessibleActionInterface) previousPageAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "previousPageAction", args)
  }

}

  // proto: static const QString & QAccessibleActionInterface::showMenuAction();
func (this *QAccessibleActionInterface) showMenuAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "showMenuAction", args)
  }

}

  // proto: static QString QAccessibleActionInterface::scrollRightAction();
func (this *QAccessibleActionInterface) scrollRightAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollRightAction", args)
  }

}

  // proto: static const QString & QAccessibleActionInterface::setFocusAction();
func (this *QAccessibleActionInterface) setFocusAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "setFocusAction", args)
  }

}

  // proto: static QString QAccessibleActionInterface::nextPageAction();
func (this *QAccessibleActionInterface) nextPageAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "nextPageAction", args)
  }

}

  // proto:  void QAccessibleActionInterface::~QAccessibleActionInterface();
func (this *QAccessibleActionInterface) FreeQAccessibleActionInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "~QAccessibleActionInterface", args)
  }

}

  // proto: static const QString & QAccessibleActionInterface::pressAction();
func (this *QAccessibleActionInterface) pressAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "pressAction", args)
  }

}

  // proto:  void QAccessibleActionInterface::doAction(const QString & actionName);
func (this *QAccessibleActionInterface) doAction(args ...interface{}) () {
  // doAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface8doActionERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "doAction", args)
  }

}

  // proto: static const QString & QAccessibleActionInterface::increaseAction();
func (this *QAccessibleActionInterface) increaseAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "increaseAction", args)
  }

}

  // proto:  QStringList QAccessibleActionInterface::keyBindingsForAction(const QString & actionName);
func (this *QAccessibleActionInterface) keyBindingsForAction(args ...interface{}) () {
  // keyBindingsForAction(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface20keyBindingsForActionERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "keyBindingsForAction", args)
  }

}

  // proto: static QString QAccessibleActionInterface::scrollDownAction();
func (this *QAccessibleActionInterface) scrollDownAction_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollDownAction", args)
  }

}

  // proto:  QAccessibleImageInterface * QAccessibleInterface::imageInterface();
func (this *QAccessibleInterface) imageInterface(args ...interface{}) () {
  // imageInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface14imageInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "imageInterface", args)
  }

}

  // proto:  QAccessibleTableInterface * QAccessibleInterface::tableInterface();
func (this *QAccessibleInterface) tableInterface(args ...interface{}) () {
  // tableInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface14tableInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableInterface", args)
  }

}

  // proto:  QAccessibleEditableTextInterface * QAccessibleInterface::editableTextInterface();
func (this *QAccessibleInterface) editableTextInterface(args ...interface{}) () {
  // editableTextInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface21editableTextInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "editableTextInterface", args)
  }

}

  // proto:  QAccessibleValueInterface * QAccessibleInterface::valueInterface();
func (this *QAccessibleInterface) valueInterface(args ...interface{}) () {
  // valueInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface14valueInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "valueInterface", args)
  }

}

  // proto:  QRect QAccessibleInterface::rect();
func (this *QAccessibleInterface) rect(args ...interface{}) () {
  // rect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface4rectEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "rect", args)
  }

}

  // proto:  QObject * QAccessibleInterface::object();
func (this *QAccessibleInterface) object(args ...interface{}) () {
  // object()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface6objectEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "object", args)
  }

}

  // proto:  QAccessibleActionInterface * QAccessibleInterface::actionInterface();
func (this *QAccessibleInterface) actionInterface(args ...interface{}) () {
  // actionInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface15actionInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "actionInterface", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleInterface::parent();
func (this *QAccessibleInterface) parent(args ...interface{}) () {
  // parent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface6parentEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "parent", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleInterface::childAt(int x, int y);
func (this *QAccessibleInterface) childAt(args ...interface{}) () {
  // childAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface7childAtEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "childAt", args)
  }

}

  // proto:  int QAccessibleInterface::childCount();
func (this *QAccessibleInterface) childCount(args ...interface{}) () {
  // childCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface10childCountEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "childCount", args)
  }

}

  // proto:  QAccessibleTableCellInterface * QAccessibleInterface::tableCellInterface();
func (this *QAccessibleInterface) tableCellInterface(args ...interface{}) () {
  // tableCellInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface18tableCellInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableCellInterface", args)
  }

}

  // proto:  int QAccessibleInterface::indexOfChild(const QAccessibleInterface * );
func (this *QAccessibleInterface) indexOfChild(args ...interface{}) () {
  // indexOfChild(const class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "const QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface12indexOfChildEPKS_
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "indexOfChild", args)
  }

}

  // proto:  QColor QAccessibleInterface::foregroundColor();
func (this *QAccessibleInterface) foregroundColor(args ...interface{}) () {
  // foregroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface15foregroundColorEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "foregroundColor", args)
  }

}

  // proto:  bool QAccessibleInterface::isValid();
func (this *QAccessibleInterface) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface7isValidEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "isValid", args)
  }

}

  // proto:  QWindow * QAccessibleInterface::window();
func (this *QAccessibleInterface) window(args ...interface{}) () {
  // window()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface6windowEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "window", args)
  }

}

  // proto:  void QAccessibleInterface::virtual_hook(int id, void * data);
func (this *QAccessibleInterface) virtual_hook(args ...interface{}) () {
  // virtual_hook(int, void *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.VoidpTy() // "void *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface12virtual_hookEiPv
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "virtual_hook", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleInterface::focusChild();
func (this *QAccessibleInterface) focusChild(args ...interface{}) () {
  // focusChild()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface10focusChildEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "focusChild", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleInterface::child(int index);
func (this *QAccessibleInterface) child(args ...interface{}) () {
  // child(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface5childEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "child", args)
  }

}

  // proto:  QAccessibleTextInterface * QAccessibleInterface::textInterface();
func (this *QAccessibleInterface) textInterface(args ...interface{}) () {
  // textInterface()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN20QAccessibleInterface13textInterfaceEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "textInterface", args)
  }

}

  // proto:  QColor QAccessibleInterface::backgroundColor();
func (this *QAccessibleInterface) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK20QAccessibleInterface15backgroundColorEv
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "backgroundColor", args)
  }

}

  // proto:  void QAccessibleInterface::~QAccessibleInterface();
func (this *QAccessibleInterface) FreeQAccessibleInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "~QAccessibleInterface", args)
  }

}

  // proto:  void QAccessibleEditableTextInterface::insertText(int offset, const QString & text);
func (this *QAccessibleEditableTextInterface) insertText(args ...interface{}) () {
  // insertText(int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleEditableTextInterface10insertTextEiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "insertText", args)
  }

}

  // proto:  void QAccessibleEditableTextInterface::replaceText(int startOffset, int endOffset, const QString & text);
func (this *QAccessibleEditableTextInterface) replaceText(args ...interface{}) () {
  // replaceText(int, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleEditableTextInterface11replaceTextEiiRK7QString
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "replaceText", args)
  }

}

  // proto:  void QAccessibleEditableTextInterface::deleteText(int startOffset, int endOffset);
func (this *QAccessibleEditableTextInterface) deleteText(args ...interface{}) () {
  // deleteText(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN32QAccessibleEditableTextInterface10deleteTextEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "deleteText", args)
  }

}

  // proto:  void QAccessibleEditableTextInterface::~QAccessibleEditableTextInterface();
func (this *QAccessibleEditableTextInterface) FreeQAccessibleEditableTextInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleEditableTextInterface", "~QAccessibleEditableTextInterface", args)
  }

}

  // proto:  int QAccessibleTableCellInterface::columnIndex();
func (this *QAccessibleTableCellInterface) columnIndex(args ...interface{}) () {
  // columnIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface11columnIndexEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "columnIndex", args)
  }

}

  // proto:  void QAccessibleTableCellInterface::~QAccessibleTableCellInterface();
func (this *QAccessibleTableCellInterface) FreeQAccessibleTableCellInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "~QAccessibleTableCellInterface", args)
  }

}

  // proto:  int QAccessibleTableCellInterface::columnExtent();
func (this *QAccessibleTableCellInterface) columnExtent(args ...interface{}) () {
  // columnExtent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface12columnExtentEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "columnExtent", args)
  }

}

  // proto:  int QAccessibleTableCellInterface::rowIndex();
func (this *QAccessibleTableCellInterface) rowIndex(args ...interface{}) () {
  // rowIndex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface8rowIndexEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "rowIndex", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleTableCellInterface::table();
func (this *QAccessibleTableCellInterface) table(args ...interface{}) () {
  // table()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface5tableEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "table", args)
  }

}

  // proto:  int QAccessibleTableCellInterface::rowExtent();
func (this *QAccessibleTableCellInterface) rowExtent(args ...interface{}) () {
  // rowExtent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface9rowExtentEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "rowExtent", args)
  }

}

  // proto:  QList<QAccessibleInterface *> QAccessibleTableCellInterface::rowHeaderCells();
func (this *QAccessibleTableCellInterface) rowHeaderCells(args ...interface{}) () {
  // rowHeaderCells()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface14rowHeaderCellsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "rowHeaderCells", args)
  }

}

  // proto:  QList<QAccessibleInterface *> QAccessibleTableCellInterface::columnHeaderCells();
func (this *QAccessibleTableCellInterface) columnHeaderCells(args ...interface{}) () {
  // columnHeaderCells()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface17columnHeaderCellsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "columnHeaderCells", args)
  }

}

  // proto:  bool QAccessibleTableCellInterface::isSelected();
func (this *QAccessibleTableCellInterface) isSelected(args ...interface{}) () {
  // isSelected()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTableCellInterface10isSelectedEv
  default:
    qtrt.ErrorResolve("QAccessibleTableCellInterface", "isSelected", args)
  }

}

  // proto:  bool QAccessibleTableInterface::unselectColumn(int column);
func (this *QAccessibleTableInterface) unselectColumn(args ...interface{}) () {
  // unselectColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface14unselectColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "unselectColumn", args)
  }

}

  // proto:  QString QAccessibleTableInterface::columnDescription(int column);
func (this *QAccessibleTableInterface) columnDescription(args ...interface{}) () {
  // columnDescription(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface17columnDescriptionEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "columnDescription", args)
  }

}

  // proto:  int QAccessibleTableInterface::selectedCellCount();
func (this *QAccessibleTableInterface) selectedCellCount(args ...interface{}) () {
  // selectedCellCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface17selectedCellCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedCellCount", args)
  }

}

  // proto:  QList<QAccessibleInterface *> QAccessibleTableInterface::selectedCells();
func (this *QAccessibleTableInterface) selectedCells(args ...interface{}) () {
  // selectedCells()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface13selectedCellsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedCells", args)
  }

}

  // proto:  bool QAccessibleTableInterface::selectRow(int row);
func (this *QAccessibleTableInterface) selectRow(args ...interface{}) () {
  // selectRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface9selectRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectRow", args)
  }

}

  // proto:  int QAccessibleTableInterface::selectedRowCount();
func (this *QAccessibleTableInterface) selectedRowCount(args ...interface{}) () {
  // selectedRowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface16selectedRowCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedRowCount", args)
  }

}

  // proto:  void QAccessibleTableInterface::~QAccessibleTableInterface();
func (this *QAccessibleTableInterface) FreeQAccessibleTableInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "~QAccessibleTableInterface", args)
  }

}

  // proto:  QList<int> QAccessibleTableInterface::selectedColumns();
func (this *QAccessibleTableInterface) selectedColumns(args ...interface{}) () {
  // selectedColumns()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface15selectedColumnsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedColumns", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleTableInterface::cellAt(int row, int column);
func (this *QAccessibleTableInterface) cellAt(args ...interface{}) () {
  // cellAt(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface6cellAtEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "cellAt", args)
  }

}

  // proto:  QList<int> QAccessibleTableInterface::selectedRows();
func (this *QAccessibleTableInterface) selectedRows(args ...interface{}) () {
  // selectedRows()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface12selectedRowsEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedRows", args)
  }

}

  // proto:  void QAccessibleTableInterface::modelChange(QAccessibleTableModelChangeEvent * event);
func (this *QAccessibleTableInterface) modelChange(args ...interface{}) () {
  // modelChange(class QAccessibleTableModelChangeEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleTableModelChangeEvent{}) // "QAccessibleTableModelChangeEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface11modelChangeEP32QAccessibleTableModelChangeEvent
    var arg0 = args[0].(QAccessibleTableModelChangeEvent).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "modelChange", args)
  }

}

  // proto:  int QAccessibleTableInterface::columnCount();
func (this *QAccessibleTableInterface) columnCount(args ...interface{}) () {
  // columnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface11columnCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "columnCount", args)
  }

}

  // proto:  bool QAccessibleTableInterface::selectColumn(int column);
func (this *QAccessibleTableInterface) selectColumn(args ...interface{}) () {
  // selectColumn(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface12selectColumnEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectColumn", args)
  }

}

  // proto:  bool QAccessibleTableInterface::unselectRow(int row);
func (this *QAccessibleTableInterface) unselectRow(args ...interface{}) () {
  // unselectRow(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleTableInterface11unselectRowEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "unselectRow", args)
  }

}

  // proto:  int QAccessibleTableInterface::rowCount();
func (this *QAccessibleTableInterface) rowCount(args ...interface{}) () {
  // rowCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface8rowCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "rowCount", args)
  }

}

  // proto:  QString QAccessibleTableInterface::rowDescription(int row);
func (this *QAccessibleTableInterface) rowDescription(args ...interface{}) () {
  // rowDescription(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface14rowDescriptionEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "rowDescription", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleTableInterface::summary();
func (this *QAccessibleTableInterface) summary(args ...interface{}) () {
  // summary()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface7summaryEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "summary", args)
  }

}

  // proto:  bool QAccessibleTableInterface::isColumnSelected(int column);
func (this *QAccessibleTableInterface) isColumnSelected(args ...interface{}) () {
  // isColumnSelected(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface16isColumnSelectedEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "isColumnSelected", args)
  }

}

  // proto:  QAccessibleInterface * QAccessibleTableInterface::caption();
func (this *QAccessibleTableInterface) caption(args ...interface{}) () {
  // caption()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface7captionEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "caption", args)
  }

}

  // proto:  bool QAccessibleTableInterface::isRowSelected(int row);
func (this *QAccessibleTableInterface) isRowSelected(args ...interface{}) () {
  // isRowSelected(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface13isRowSelectedEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "isRowSelected", args)
  }

}

  // proto:  int QAccessibleTableInterface::selectedColumnCount();
func (this *QAccessibleTableInterface) selectedColumnCount(args ...interface{}) () {
  // selectedColumnCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleTableInterface19selectedColumnCountEv
  default:
    qtrt.ErrorResolve("QAccessibleTableInterface", "selectedColumnCount", args)
  }

}

  // proto:  QString QAccessibleTextUpdateEvent::textInserted();
func (this *QAccessibleTextUpdateEvent) textInserted(args ...interface{}) () {
  // textInserted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextUpdateEvent12textInsertedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textInserted", args)
  }

}

  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QAccessibleInterface * iface, int position, const QString & oldText, const QString & text);
func NewQAccessibleTextUpdateEvent(args ...interface{}) QAccessibleTextUpdateEvent {
  return QAccessibleTextUpdateEvent{}
}

  // proto:  QString QAccessibleTextUpdateEvent::textRemoved();
func (this *QAccessibleTextUpdateEvent) textRemoved(args ...interface{}) () {
  // textRemoved()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextUpdateEvent11textRemovedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textRemoved", args)
  }

}

  // proto:  int QAccessibleTextUpdateEvent::changePosition();
func (this *QAccessibleTextUpdateEvent) changePosition(args ...interface{}) () {
  // changePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextUpdateEvent14changePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "changePosition", args)
  }

}

  // proto:  QString QAccessibleImageInterface::imageDescription();
func (this *QAccessibleImageInterface) imageDescription(args ...interface{}) () {
  // imageDescription()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleImageInterface16imageDescriptionEv
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "imageDescription", args)
  }

}

  // proto:  QPoint QAccessibleImageInterface::imagePosition();
func (this *QAccessibleImageInterface) imagePosition(args ...interface{}) () {
  // imagePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleImageInterface13imagePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "imagePosition", args)
  }

}

  // proto:  void QAccessibleImageInterface::~QAccessibleImageInterface();
func (this *QAccessibleImageInterface) FreeQAccessibleImageInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "~QAccessibleImageInterface", args)
  }

}

  // proto:  QSize QAccessibleImageInterface::imageSize();
func (this *QAccessibleImageInterface) imageSize(args ...interface{}) () {
  // imageSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleImageInterface9imageSizeEv
  default:
    qtrt.ErrorResolve("QAccessibleImageInterface", "imageSize", args)
  }

}

  // proto:  QString QAccessibleTextInsertEvent::textInserted();
func (this *QAccessibleTextInsertEvent) textInserted(args ...interface{}) () {
  // textInserted()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextInsertEvent12textInsertedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "textInserted", args)
  }

}

  // proto:  int QAccessibleTextInsertEvent::changePosition();
func (this *QAccessibleTextInsertEvent) changePosition(args ...interface{}) () {
  // changePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextInsertEvent14changePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "changePosition", args)
  }

}

  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QAccessibleInterface * iface, int position, const QString & text);
func NewQAccessibleTextInsertEvent(args ...interface{}) QAccessibleTextInsertEvent {
  return QAccessibleTextInsertEvent{}
}

  // proto:  QVariant QAccessibleValueInterface::maximumValue();
func (this *QAccessibleValueInterface) maximumValue(args ...interface{}) () {
  // maximumValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface12maximumValueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "maximumValue", args)
  }

}

  // proto:  QVariant QAccessibleValueInterface::minimumStepSize();
func (this *QAccessibleValueInterface) minimumStepSize(args ...interface{}) () {
  // minimumStepSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface15minimumStepSizeEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "minimumStepSize", args)
  }

}

  // proto:  QVariant QAccessibleValueInterface::currentValue();
func (this *QAccessibleValueInterface) currentValue(args ...interface{}) () {
  // currentValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface12currentValueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "currentValue", args)
  }

}

  // proto:  QVariant QAccessibleValueInterface::minimumValue();
func (this *QAccessibleValueInterface) minimumValue(args ...interface{}) () {
  // minimumValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK25QAccessibleValueInterface12minimumValueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "minimumValue", args)
  }

}

  // proto:  void QAccessibleValueInterface::~QAccessibleValueInterface();
func (this *QAccessibleValueInterface) FreeQAccessibleValueInterface(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "~QAccessibleValueInterface", args)
  }

}

  // proto:  void QAccessibleValueInterface::setCurrentValue(const QVariant & value);
func (this *QAccessibleValueInterface) setCurrentValue(args ...interface{}) () {
  // setCurrentValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN25QAccessibleValueInterface15setCurrentValueERK8QVariant
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleValueInterface", "setCurrentValue", args)
  }

}

  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QObject * obj, int position, const QString & text);
func NewQAccessibleTextRemoveEvent(args ...interface{}) QAccessibleTextRemoveEvent {
  return QAccessibleTextRemoveEvent{}
}

  // proto:  QString QAccessibleTextRemoveEvent::textRemoved();
func (this *QAccessibleTextRemoveEvent) textRemoved(args ...interface{}) () {
  // textRemoved()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextRemoveEvent11textRemovedEv
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "textRemoved", args)
  }

}

  // proto:  int QAccessibleTextRemoveEvent::changePosition();
func (this *QAccessibleTextRemoveEvent) changePosition(args ...interface{}) () {
  // changePosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextRemoveEvent14changePositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "changePosition", args)
  }

}

  // proto:  int QAccessibleTextSelectionEvent::selectionEnd();
func (this *QAccessibleTextSelectionEvent) selectionEnd(args ...interface{}) () {
  // selectionEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTextSelectionEvent12selectionEndEv
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionEnd", args)
  }

}

  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QAccessibleInterface * iface, int start, int end);
func NewQAccessibleTextSelectionEvent(args ...interface{}) QAccessibleTextSelectionEvent {
  return QAccessibleTextSelectionEvent{}
}

  // proto:  int QAccessibleTextSelectionEvent::selectionStart();
func (this *QAccessibleTextSelectionEvent) selectionStart(args ...interface{}) () {
  // selectionStart()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK29QAccessibleTextSelectionEvent14selectionStartEv
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionStart", args)
  }

}

  // proto:  void QAccessibleTextSelectionEvent::setSelection(int start, int end);
func (this *QAccessibleTextSelectionEvent) setSelection(args ...interface{}) () {
  // setSelection(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QAccessibleTextSelectionEvent12setSelectionEii
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "setSelection", args)
  }

}

  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QAccessibleInterface * iface, int cursorPos);
func NewQAccessibleTextCursorEvent(args ...interface{}) QAccessibleTextCursorEvent {
  return QAccessibleTextCursorEvent{}
}

  // proto:  void QAccessibleTextCursorEvent::setCursorPosition(int position);
func (this *QAccessibleTextCursorEvent) setCursorPosition(args ...interface{}) () {
  // setCursorPosition(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextCursorEvent17setCursorPositionEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "setCursorPosition", args)
  }

}

  // proto:  int QAccessibleTextCursorEvent::cursorPosition();
func (this *QAccessibleTextCursorEvent) cursorPosition(args ...interface{}) () {
  // cursorPosition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleTextCursorEvent14cursorPositionEv
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "cursorPosition", args)
  }

}

  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QObject * obj, const QVariant & val);
func NewQAccessibleValueChangeEvent(args ...interface{}) QAccessibleValueChangeEvent {
  return QAccessibleValueChangeEvent{}
}

  // proto:  void QAccessibleValueChangeEvent::setValue(const QVariant & val);
func (this *QAccessibleValueChangeEvent) setValue(args ...interface{}) () {
  // setValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAccessibleValueChangeEvent8setValueERK8QVariant
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "setValue", args)
  }

}

  // proto:  QVariant QAccessibleValueChangeEvent::value();
func (this *QAccessibleValueChangeEvent) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAccessibleValueChangeEvent5valueEv
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "value", args)
  }

}

// <= body block end

