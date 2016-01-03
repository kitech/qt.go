package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
extern void _ZN11QAccessible25deleteAccessibleInterfaceEj(int32_t arg0);
  // proto: static QAccessibleInterface * QAccessible::accessibleInterface(Id uniqueId);
extern void _ZN11QAccessible19accessibleInterfaceEj(int32_t arg0);
  // proto:  void QAccessible::QAccessible();
extern void* dector_ZN11QAccessibleC1Ev();
extern void _ZN11QAccessibleC1Ev(void* qthis);
  // proto:  void QAccessibleTableModelChangeEvent::setFirstColumn(int col);
extern void demth_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(void* qthis, int32_t arg0);
  // proto:  void QAccessibleTableModelChangeEvent::setFirstRow(int row);
extern void demth_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(void* qthis, int32_t arg0);
  // proto:  int QAccessibleTableModelChangeEvent::firstRow();
extern void demth_ZNK32QAccessibleTableModelChangeEvent8firstRowEv(void* qthis);
  // proto:  void QAccessibleTableModelChangeEvent::setLastColumn(int col);
extern void demth_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(void* qthis, int32_t arg0);
  // proto:  int QAccessibleTableModelChangeEvent::firstColumn();
extern void demth_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(void* qthis);
  // proto:  int QAccessibleTableModelChangeEvent::lastColumn();
extern void demth_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(void* qthis);
  // proto:  void QAccessibleTableModelChangeEvent::setLastRow(int row);
extern void demth_ZN32QAccessibleTableModelChangeEvent10setLastRowEi(void* qthis, int32_t arg0);
  // proto:  int QAccessibleTableModelChangeEvent::lastRow();
extern void demth_ZNK32QAccessibleTableModelChangeEvent7lastRowEv(void* qthis);
  // proto:  void QAccessibleTextInterface::~QAccessibleTextInterface();
extern void _ZN24QAccessibleTextInterfaceD0Ev(void* qthis);
  // proto:  QObject * QAccessibleEvent::object();
extern void demth_ZNK16QAccessibleEvent6objectEv(void* qthis);
  // proto:  void QAccessibleEvent::setChild(int chld);
extern void demth_ZN16QAccessibleEvent8setChildEi(void* qthis, int32_t arg0);
  // proto:  QAccessibleInterface * QAccessibleEvent::accessibleInterface();
extern void _ZNK16QAccessibleEvent19accessibleInterfaceEv(void* qthis);
  // proto:  void QAccessibleEvent::QAccessibleEvent(const QAccessibleEvent & );
extern void* dector_ZN16QAccessibleEventC1ERKS_(void* arg0);
extern void _ZN16QAccessibleEventC1ERKS_(void* qthis, void* arg0);
  // proto:  int QAccessibleEvent::child();
extern void demth_ZNK16QAccessibleEvent5childEv(void* qthis);
  // proto:  void QAccessibleEvent::~QAccessibleEvent();
extern void _ZN16QAccessibleEventD0Ev(void* qthis);
  // proto: static QString QAccessibleActionInterface::scrollUpAction();
extern void _ZN26QAccessibleActionInterface14scrollUpActionEv();
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
  // proto: static const QString & QAccessibleActionInterface::increaseAction();
extern void _ZN26QAccessibleActionInterface14increaseActionEv();
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
  // proto:  QAccessibleActionInterface * QAccessibleInterface::actionInterface();
extern void demth_ZN20QAccessibleInterface15actionInterfaceEv(void* qthis);
  // proto:  QAccessibleTableCellInterface * QAccessibleInterface::tableCellInterface();
extern void demth_ZN20QAccessibleInterface18tableCellInterfaceEv(void* qthis);
  // proto:  QColor QAccessibleInterface::foregroundColor();
extern void _ZNK20QAccessibleInterface15foregroundColorEv(void* qthis);
  // proto:  QWindow * QAccessibleInterface::window();
extern void _ZNK20QAccessibleInterface6windowEv(void* qthis);
  // proto:  void QAccessibleInterface::virtual_hook(int id, void * data);
extern void _ZN20QAccessibleInterface12virtual_hookEiPv(void* qthis, int32_t arg0, void* arg1);
  // proto:  QAccessibleInterface * QAccessibleInterface::focusChild();
extern void _ZNK20QAccessibleInterface10focusChildEv(void* qthis);
  // proto:  QAccessibleTextInterface * QAccessibleInterface::textInterface();
extern void demth_ZN20QAccessibleInterface13textInterfaceEv(void* qthis);
  // proto:  QColor QAccessibleInterface::backgroundColor();
extern void _ZNK20QAccessibleInterface15backgroundColorEv(void* qthis);
  // proto:  void QAccessibleInterface::~QAccessibleInterface();
extern void _ZN20QAccessibleInterfaceD0Ev(void* qthis);
  // proto:  void QAccessibleEditableTextInterface::~QAccessibleEditableTextInterface();
extern void _ZN32QAccessibleEditableTextInterfaceD0Ev(void* qthis);
  // proto:  void QAccessibleTableCellInterface::~QAccessibleTableCellInterface();
extern void _ZN29QAccessibleTableCellInterfaceD0Ev(void* qthis);
  // proto:  void QAccessibleTableInterface::~QAccessibleTableInterface();
extern void _ZN25QAccessibleTableInterfaceD0Ev(void* qthis);
  // proto:  QString QAccessibleTextUpdateEvent::textInserted();
extern void demth_ZNK26QAccessibleTextUpdateEvent12textInsertedEv(void* qthis);
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QAccessibleInterface * iface, int position, const QString & oldText, const QString & text);
extern void* dector_ZN26QAccessibleTextUpdateEventC1EP20QAccessibleInterfaceiRK7QStringS4_(void* arg0, int32_t arg1, void* arg2, void* arg3);
extern void demth_ZN26QAccessibleTextUpdateEventC1EP20QAccessibleInterfaceiRK7QStringS4_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3);
  // proto:  QString QAccessibleTextUpdateEvent::textRemoved();
extern void demth_ZNK26QAccessibleTextUpdateEvent11textRemovedEv(void* qthis);
  // proto:  int QAccessibleTextUpdateEvent::changePosition();
extern void demth_ZNK26QAccessibleTextUpdateEvent14changePositionEv(void* qthis);
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QObject * obj, int position, const QString & oldText, const QString & text);
extern void* dector_ZN26QAccessibleTextUpdateEventC1EP7QObjectiRK7QStringS4_(void* arg0, int32_t arg1, void* arg2, void* arg3);
extern void demth_ZN26QAccessibleTextUpdateEventC1EP7QObjectiRK7QStringS4_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3);
  // proto:  void QAccessibleImageInterface::~QAccessibleImageInterface();
extern void _ZN25QAccessibleImageInterfaceD0Ev(void* qthis);
  // proto:  QString QAccessibleTextInsertEvent::textInserted();
extern void demth_ZNK26QAccessibleTextInsertEvent12textInsertedEv(void* qthis);
  // proto:  int QAccessibleTextInsertEvent::changePosition();
extern void demth_ZNK26QAccessibleTextInsertEvent14changePositionEv(void* qthis);
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextInsertEventC1EP20QAccessibleInterfaceiRK7QString(void* arg0, int32_t arg1, void* arg2);
extern void demth_ZN26QAccessibleTextInsertEventC1EP20QAccessibleInterfaceiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QObject * obj, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextInsertEventC1EP7QObjectiRK7QString(void* arg0, int32_t arg1, void* arg2);
extern void demth_ZN26QAccessibleTextInsertEventC1EP7QObjectiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  void QAccessibleValueInterface::~QAccessibleValueInterface();
extern void _ZN25QAccessibleValueInterfaceD0Ev(void* qthis);
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QObject * obj, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextRemoveEventC1EP7QObjectiRK7QString(void* arg0, int32_t arg1, void* arg2);
extern void demth_ZN26QAccessibleTextRemoveEventC1EP7QObjectiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  QString QAccessibleTextRemoveEvent::textRemoved();
extern void demth_ZNK26QAccessibleTextRemoveEvent11textRemovedEv(void* qthis);
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void* dector_ZN26QAccessibleTextRemoveEventC1EP20QAccessibleInterfaceiRK7QString(void* arg0, int32_t arg1, void* arg2);
extern void demth_ZN26QAccessibleTextRemoveEventC1EP20QAccessibleInterfaceiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  int QAccessibleTextRemoveEvent::changePosition();
extern void demth_ZNK26QAccessibleTextRemoveEvent14changePositionEv(void* qthis);
  // proto:  int QAccessibleTextSelectionEvent::selectionEnd();
extern void demth_ZNK29QAccessibleTextSelectionEvent12selectionEndEv(void* qthis);
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QAccessibleInterface * iface, int start, int end);
extern void* dector_ZN29QAccessibleTextSelectionEventC1EP20QAccessibleInterfaceii(void* arg0, int32_t arg1, int32_t arg2);
extern void demth_ZN29QAccessibleTextSelectionEventC1EP20QAccessibleInterfaceii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  int QAccessibleTextSelectionEvent::selectionStart();
extern void demth_ZNK29QAccessibleTextSelectionEvent14selectionStartEv(void* qthis);
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QObject * obj, int start, int end);
extern void* dector_ZN29QAccessibleTextSelectionEventC1EP7QObjectii(void* arg0, int32_t arg1, int32_t arg2);
extern void demth_ZN29QAccessibleTextSelectionEventC1EP7QObjectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QAccessibleTextSelectionEvent::setSelection(int start, int end);
extern void demth_ZN29QAccessibleTextSelectionEvent12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QAccessibleInterface * iface, int cursorPos);
extern void* dector_ZN26QAccessibleTextCursorEventC1EP20QAccessibleInterfacei(void* arg0, int32_t arg1);
extern void demth_ZN26QAccessibleTextCursorEventC1EP20QAccessibleInterfacei(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QAccessibleTextCursorEvent::setCursorPosition(int position);
extern void demth_ZN26QAccessibleTextCursorEvent17setCursorPositionEi(void* qthis, int32_t arg0);
  // proto:  int QAccessibleTextCursorEvent::cursorPosition();
extern void demth_ZNK26QAccessibleTextCursorEvent14cursorPositionEv(void* qthis);
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QObject * obj, int cursorPos);
extern void* dector_ZN26QAccessibleTextCursorEventC1EP7QObjecti(void* arg0, int32_t arg1);
extern void demth_ZN26QAccessibleTextCursorEventC1EP7QObjecti(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QObject * obj, const QVariant & val);
extern void* dector_ZN27QAccessibleValueChangeEventC1EP7QObjectRK8QVariant(void* arg0, void* arg1);
extern void demth_ZN27QAccessibleValueChangeEventC1EP7QObjectRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QAccessibleInterface * iface, const QVariant & val);
extern void* dector_ZN27QAccessibleValueChangeEventC1EP20QAccessibleInterfaceRK8QVariant(void* arg0, void* arg1);
extern void demth_ZN27QAccessibleValueChangeEventC1EP20QAccessibleInterfaceRK8QVariant(void* qthis, void* arg0, void* arg1);
  // proto:  void QAccessibleValueChangeEvent::setValue(const QVariant & val);
extern void demth_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(void* qthis, void* arg0);
  // proto:  QVariant QAccessibleValueChangeEvent::value();
extern void demth_ZNK27QAccessibleValueChangeEvent5valueEv(void* qthis);
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTableModelChangeEvent)=48
type QAccessibleTableModelChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextInterface)=8
type QAccessibleTextInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleEvent)=32
type QAccessibleEvent struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleActionInterface)=8
type QAccessibleActionInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleInterface)=8
type QAccessibleInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleEditableTextInterface)=8
type QAccessibleEditableTextInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTableCellInterface)=8
type QAccessibleTableCellInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTableInterface)=8
type QAccessibleTableInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextUpdateEvent)=56
type QAccessibleTextUpdateEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleStateChangeEvent)=40
type QAccessibleStateChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleImageInterface)=8
type QAccessibleImageInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextInsertEvent)=48
type QAccessibleTextInsertEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleValueInterface)=8
type QAccessibleValueInterface struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextRemoveEvent)=48
type QAccessibleTextRemoveEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextSelectionEvent)=40
type QAccessibleTextSelectionEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextCursorEvent)=32
type QAccessibleTextCursorEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleValueChangeEvent)=48
type QAccessibleValueChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  qclsinst unsafe.Pointer /* *C.void */;
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
    // invoke: void setFirstColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(this.qclsinst, arg0)
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
    // invoke: void setFirstRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(this.qclsinst, arg0)
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
    // invoke: int firstRow()
    C.demth_ZNK32QAccessibleTableModelChangeEvent8firstRowEv(this.qclsinst)
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
    // invoke: void setLastColumn(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(this.qclsinst, arg0)
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
    // invoke: int firstColumn()
    C.demth_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(this.qclsinst)
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
    // invoke: int lastColumn()
    C.demth_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(this.qclsinst)
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
    // invoke: void setLastRow(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN32QAccessibleTableModelChangeEvent10setLastRowEi(this.qclsinst, arg0)
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
    // invoke: int lastRow()
    C.demth_ZNK32QAccessibleTableModelChangeEvent7lastRowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastRow", args)
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
    // invoke: QObject * object()
    C.demth_ZNK16QAccessibleEvent6objectEv(this.qclsinst)
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
    // invoke: void setChild(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN16QAccessibleEvent8setChildEi(this.qclsinst, arg0)
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
    // invoke: QAccessibleInterface * accessibleInterface()
    C._ZNK16QAccessibleEvent19accessibleInterfaceEv(this.qclsinst)
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
    // invoke: int child()
    C.demth_ZNK16QAccessibleEvent5childEv(this.qclsinst)
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
    // invoke: QString localizedActionName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK26QAccessibleActionInterface19localizedActionNameERK7QString(this.qclsinst, arg0)
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
    // invoke: QString localizedActionDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString(this.qclsinst, arg0)
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
    // invoke: QAccessibleImageInterface * imageInterface()
    C.demth_ZN20QAccessibleInterface14imageInterfaceEv(this.qclsinst)
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
    // invoke: QAccessibleTableInterface * tableInterface()
    C.demth_ZN20QAccessibleInterface14tableInterfaceEv(this.qclsinst)
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
    // invoke: QAccessibleEditableTextInterface * editableTextInterface()
    C.demth_ZN20QAccessibleInterface21editableTextInterfaceEv(this.qclsinst)
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
    // invoke: QAccessibleValueInterface * valueInterface()
    C.demth_ZN20QAccessibleInterface14valueInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "valueInterface", args)
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
    // invoke: QAccessibleActionInterface * actionInterface()
    C.demth_ZN20QAccessibleInterface15actionInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "actionInterface", args)
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
    // invoke: QAccessibleTableCellInterface * tableCellInterface()
    C.demth_ZN20QAccessibleInterface18tableCellInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableCellInterface", args)
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
    // invoke: QColor foregroundColor()
    C._ZNK20QAccessibleInterface15foregroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "foregroundColor", args)
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
    // invoke: QWindow * window()
    C._ZNK20QAccessibleInterface6windowEv(this.qclsinst)
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
    // invoke: void virtual_hook(int, void *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C._ZN20QAccessibleInterface12virtual_hookEiPv(this.qclsinst, arg0, arg1)
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
    // invoke: QAccessibleInterface * focusChild()
    C._ZNK20QAccessibleInterface10focusChildEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "focusChild", args)
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
    // invoke: QAccessibleTextInterface * textInterface()
    C.demth_ZN20QAccessibleInterface13textInterfaceEv(this.qclsinst)
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
    // invoke: QColor backgroundColor()
    C._ZNK20QAccessibleInterface15backgroundColorEv(this.qclsinst)
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
    // invoke: QString textInserted()
    C.demth_ZNK26QAccessibleTextUpdateEvent12textInsertedEv(this.qclsinst)
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
    // invoke: QString textRemoved()
    C.demth_ZNK26QAccessibleTextUpdateEvent11textRemovedEv(this.qclsinst)
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
    // invoke: int changePosition()
    C.demth_ZNK26QAccessibleTextUpdateEvent14changePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "changePosition", args)
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
    // invoke: QString textInserted()
    C.demth_ZNK26QAccessibleTextInsertEvent12textInsertedEv(this.qclsinst)
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
    // invoke: int changePosition()
    C.demth_ZNK26QAccessibleTextInsertEvent14changePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "changePosition", args)
  }

}

  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QAccessibleInterface * iface, int position, const QString & text);
func NewQAccessibleTextInsertEvent(args ...interface{}) QAccessibleTextInsertEvent {
  return QAccessibleTextInsertEvent{}
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
    // invoke: QString textRemoved()
    C.demth_ZNK26QAccessibleTextRemoveEvent11textRemovedEv(this.qclsinst)
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
    // invoke: int changePosition()
    C.demth_ZNK26QAccessibleTextRemoveEvent14changePositionEv(this.qclsinst)
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
    // invoke: int selectionEnd()
    C.demth_ZNK29QAccessibleTextSelectionEvent12selectionEndEv(this.qclsinst)
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
    // invoke: int selectionStart()
    C.demth_ZNK29QAccessibleTextSelectionEvent14selectionStartEv(this.qclsinst)
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
    // invoke: void setSelection(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN29QAccessibleTextSelectionEvent12setSelectionEii(this.qclsinst, arg0, arg1)
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
    // invoke: void setCursorPosition(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.demth_ZN26QAccessibleTextCursorEvent17setCursorPositionEi(this.qclsinst, arg0)
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
    // invoke: int cursorPosition()
    C.demth_ZNK26QAccessibleTextCursorEvent14cursorPositionEv(this.qclsinst)
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
    // invoke: void setValue(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(this.qclsinst, arg0)
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
    // invoke: QVariant value()
    C.demth_ZNK27QAccessibleValueChangeEvent5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "value", args)
  }

}

// <= body block end

