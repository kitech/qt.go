package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QAccessibleInterface * QAccessible::accessibleInterface(Id uniqueId);
extern void C_ZN11QAccessible19accessibleInterfaceEj(int32_t arg0); // 4
  // proto: static void QAccessible::deleteAccessibleInterface(Id uniqueId);
extern void C_ZN11QAccessible25deleteAccessibleInterfaceEj(int32_t arg0); // 4
  // proto: static void QAccessible::setRootObject(QObject * object);
extern void C_ZN11QAccessible13setRootObjectEP7QObject(void* arg0); // 4
  // proto: static QAccessibleInterface * QAccessible::queryAccessibleInterface(QObject * );
extern void C_ZN11QAccessible24queryAccessibleInterfaceEP7QObject(void* arg0); // 4
  // proto: static Id QAccessible::uniqueId(QAccessibleInterface * iface);
extern void C_ZN11QAccessible8uniqueIdEP20QAccessibleInterface(void* arg0); // 4
  // proto: static void QAccessible::setActive(bool active);
extern void C_ZN11QAccessible9setActiveEb(bool arg0); // 4
  // proto: static void QAccessible::updateAccessibility(QAccessibleEvent * event);
extern void C_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent(void* arg0); // 4
  // proto: static void QAccessible::cleanup();
extern void C_ZN11QAccessible7cleanupEv(); // 4
  // proto: static bool QAccessible::isActive();
extern void C_ZN11QAccessible8isActiveEv(); // 4
  // proto: static Id QAccessible::registerAccessibleInterface(QAccessibleInterface * iface);
extern void C_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface(void* arg0); // 4
  // proto:  QAccessibleTableModelChangeEvent::ModelChangeType QAccessibleTableModelChangeEvent::modelChangeType();
extern void C_ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv(void* qthis); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setFirstColumn(int col);
extern void C_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(void* qthis, int32_t arg0); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setLastColumn(int col);
extern void C_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(void* qthis, int32_t arg0); // 2
  // proto:  int QAccessibleTableModelChangeEvent::lastColumn();
extern void C_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(void* qthis); // 2
  // proto:  int QAccessibleTableModelChangeEvent::firstColumn();
extern void C_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(void* qthis); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setFirstRow(int row);
extern void C_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(void* qthis, int32_t arg0); // 2
  // proto:  int QAccessibleTableModelChangeEvent::firstRow();
extern void C_ZNK32QAccessibleTableModelChangeEvent8firstRowEv(void* qthis); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setLastRow(int row);
extern void C_ZN32QAccessibleTableModelChangeEvent10setLastRowEi(void* qthis, int32_t arg0); // 2
  // proto:  int QAccessibleTableModelChangeEvent::lastRow();
extern void C_ZNK32QAccessibleTableModelChangeEvent7lastRowEv(void* qthis); // 2
  // proto:  QAccessibleInterface * QAccessibleEvent::accessibleInterface();
extern void C_ZNK16QAccessibleEvent19accessibleInterfaceEv(void* qthis); // 4
  // proto:  QObject * QAccessibleEvent::object();
extern void C_ZNK16QAccessibleEvent6objectEv(void* qthis); // 2
  // proto:  QAccessible::Id QAccessibleEvent::uniqueId();
extern void C_ZNK16QAccessibleEvent8uniqueIdEv(void* qthis); // 4
  // proto:  int QAccessibleEvent::child();
extern void C_ZNK16QAccessibleEvent5childEv(void* qthis); // 2
  // proto:  void QAccessibleEvent::setChild(int chld);
extern void C_ZN16QAccessibleEvent8setChildEi(void* qthis, int32_t arg0); // 2
  // proto:  QAccessible::Event QAccessibleEvent::type();
extern void C_ZNK16QAccessibleEvent4typeEv(void* qthis); // 2
  // proto:  void QAccessibleEvent::~QAccessibleEvent();
extern void C_ZN16QAccessibleEventD2Ev(void* qthis); // 4
  // proto: static const QString & QAccessibleActionInterface::increaseAction();
extern void C_ZN26QAccessibleActionInterface14increaseActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::pressAction();
extern void C_ZN26QAccessibleActionInterface11pressActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::previousPageAction();
extern void C_ZN26QAccessibleActionInterface18previousPageActionEv(); // 4
  // proto:  QString QAccessibleActionInterface::localizedActionName(const QString & name);
extern void C_ZNK26QAccessibleActionInterface19localizedActionNameERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QAccessibleActionInterface::scrollRightAction();
extern void C_ZN26QAccessibleActionInterface17scrollRightActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::decreaseAction();
extern void C_ZN26QAccessibleActionInterface14decreaseActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::showMenuAction();
extern void C_ZN26QAccessibleActionInterface14showMenuActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::toggleAction();
extern void C_ZN26QAccessibleActionInterface12toggleActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::setFocusAction();
extern void C_ZN26QAccessibleActionInterface14setFocusActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::scrollLeftAction();
extern void C_ZN26QAccessibleActionInterface16scrollLeftActionEv(); // 4
  // proto:  QString QAccessibleActionInterface::localizedActionDescription(const QString & name);
extern void C_ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QAccessibleActionInterface::scrollUpAction();
extern void C_ZN26QAccessibleActionInterface14scrollUpActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::scrollDownAction();
extern void C_ZN26QAccessibleActionInterface16scrollDownActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::nextPageAction();
extern void C_ZN26QAccessibleActionInterface14nextPageActionEv(); // 4
  // proto:  QAccessibleTableInterface * QAccessibleInterface::tableInterface();
extern void C_ZN20QAccessibleInterface14tableInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleActionInterface * QAccessibleInterface::actionInterface();
extern void C_ZN20QAccessibleInterface15actionInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleTextInterface * QAccessibleInterface::textInterface();
extern void C_ZN20QAccessibleInterface13textInterfaceEv(void* qthis); // 2
  // proto:  QColor QAccessibleInterface::foregroundColor();
extern void C_ZNK20QAccessibleInterface15foregroundColorEv(void* qthis); // 4
  // proto:  QAccessibleEditableTextInterface * QAccessibleInterface::editableTextInterface();
extern void C_ZN20QAccessibleInterface21editableTextInterfaceEv(void* qthis); // 2
  // proto:  QWindow * QAccessibleInterface::window();
extern void C_ZNK20QAccessibleInterface6windowEv(void* qthis); // 4
  // proto:  void QAccessibleInterface::virtual_hook(int id, void * data);
extern void C_ZN20QAccessibleInterface12virtual_hookEiPv(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QColor QAccessibleInterface::backgroundColor();
extern void C_ZNK20QAccessibleInterface15backgroundColorEv(void* qthis); // 4
  // proto:  QAccessibleValueInterface * QAccessibleInterface::valueInterface();
extern void C_ZN20QAccessibleInterface14valueInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleInterface * QAccessibleInterface::focusChild();
extern void C_ZNK20QAccessibleInterface10focusChildEv(void* qthis); // 4
  // proto:  QAccessibleImageInterface * QAccessibleInterface::imageInterface();
extern void C_ZN20QAccessibleInterface14imageInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleTableCellInterface * QAccessibleInterface::tableCellInterface();
extern void C_ZN20QAccessibleInterface18tableCellInterfaceEv(void* qthis); // 2
  // proto:  QString QAccessibleTextUpdateEvent::textInserted();
extern void C_ZNK26QAccessibleTextUpdateEvent12textInsertedEv(void* qthis); // 2
  // proto:  int QAccessibleTextUpdateEvent::changePosition();
extern void C_ZNK26QAccessibleTextUpdateEvent14changePositionEv(void* qthis); // 2
  // proto:  QString QAccessibleTextUpdateEvent::textRemoved();
extern void C_ZNK26QAccessibleTextUpdateEvent11textRemovedEv(void* qthis); // 2
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QAccessibleInterface * iface, int position, const QString & oldText, const QString & text);
extern void C_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3); // 1
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QObject * obj, int position, const QString & oldText, const QString & text);
extern void C_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_(void* qthis, void* arg0, int32_t arg1, void* arg2, void* arg3); // 1
  // proto:  QAccessible::State QAccessibleStateChangeEvent::changedStates();
extern void C_ZNK27QAccessibleStateChangeEvent13changedStatesEv(void* qthis); // 2
  // proto:  QString QAccessibleTextInsertEvent::textInserted();
extern void C_ZNK26QAccessibleTextInsertEvent12textInsertedEv(void* qthis); // 2
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void C_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QObject * obj, int position, const QString & text);
extern void C_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  int QAccessibleTextInsertEvent::changePosition();
extern void C_ZNK26QAccessibleTextInsertEvent14changePositionEv(void* qthis); // 2
  // proto:  QString QAccessibleTextRemoveEvent::textRemoved();
extern void C_ZNK26QAccessibleTextRemoveEvent11textRemovedEv(void* qthis); // 2
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QObject * obj, int position, const QString & text);
extern void C_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void C_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString(void* qthis, void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  int QAccessibleTextRemoveEvent::changePosition();
extern void C_ZNK26QAccessibleTextRemoveEvent14changePositionEv(void* qthis); // 2
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QAccessibleInterface * iface, int start, int end);
extern void C_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QObject * obj, int start, int end);
extern void C_ZN29QAccessibleTextSelectionEventC2EP7QObjectii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  int QAccessibleTextSelectionEvent::selectionStart();
extern void C_ZNK29QAccessibleTextSelectionEvent14selectionStartEv(void* qthis); // 2
  // proto:  int QAccessibleTextSelectionEvent::selectionEnd();
extern void C_ZNK29QAccessibleTextSelectionEvent12selectionEndEv(void* qthis); // 2
  // proto:  void QAccessibleTextSelectionEvent::setSelection(int start, int end);
extern void C_ZN29QAccessibleTextSelectionEvent12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  int QAccessibleTextCursorEvent::cursorPosition();
extern void C_ZNK26QAccessibleTextCursorEvent14cursorPositionEv(void* qthis); // 2
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QAccessibleInterface * iface, int cursorPos);
extern void C_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei(void* qthis, void* arg0, int32_t arg1); // 1
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QObject * obj, int cursorPos);
extern void C_ZN26QAccessibleTextCursorEventC2EP7QObjecti(void* qthis, void* arg0, int32_t arg1); // 1
  // proto:  void QAccessibleTextCursorEvent::setCursorPosition(int position);
extern void C_ZN26QAccessibleTextCursorEvent17setCursorPositionEi(void* qthis, int32_t arg0); // 2
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QObject * obj, const QVariant & val);
extern void C_ZN27QAccessibleValueChangeEventC2EP7QObjectRK8QVariant(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QAccessibleInterface * iface, const QVariant & val);
extern void C_ZN27QAccessibleValueChangeEventC2EP20QAccessibleInterfaceRK8QVariant(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QAccessibleValueChangeEvent::setValue(const QVariant & val);
extern void C_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(void* qthis, void* arg0); // 2
  // proto:  QVariant QAccessibleValueChangeEvent::value();
extern void C_ZNK27QAccessibleValueChangeEvent5valueEv(void* qthis); // 2
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

// accessibleInterface(Id)
func (this *QAccessible) accessibleInterface_s(args ...interface{}) () {
  // accessibleInterface(Id)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Id"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible19accessibleInterfaceEj
    // invoke: QAccessibleInterface * accessibleInterface(Id)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible19accessibleInterfaceEj(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "accessibleInterface", args)
  }

}

// deleteAccessibleInterface(Id)
func (this *QAccessible) deleteAccessibleInterface_s(args ...interface{}) () {
  // deleteAccessibleInterface(Id)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "Id"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible25deleteAccessibleInterfaceEj
    // invoke: void deleteAccessibleInterface(Id)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible25deleteAccessibleInterfaceEj(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "deleteAccessibleInterface", args)
  }

}

// setRootObject(class QObject *)
func (this *QAccessible) setRootObject_s(args ...interface{}) () {
  // setRootObject(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible13setRootObjectEP7QObject
    // invoke: void setRootObject(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible13setRootObjectEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "setRootObject", args)
  }

}

// queryAccessibleInterface(class QObject *)
func (this *QAccessible) queryAccessibleInterface_s(args ...interface{}) () {
  // queryAccessibleInterface(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible24queryAccessibleInterfaceEP7QObject
    // invoke: QAccessibleInterface * queryAccessibleInterface(class QObject *)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible24queryAccessibleInterfaceEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "queryAccessibleInterface", args)
  }

}

// uniqueId(class QAccessibleInterface *)
func (this *QAccessible) uniqueId_s(args ...interface{}) () {
  // uniqueId(class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible8uniqueIdEP20QAccessibleInterface
    // invoke: Id uniqueId(class QAccessibleInterface *)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible8uniqueIdEP20QAccessibleInterface(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "uniqueId", args)
  }

}

// setActive(_Bool)
func (this *QAccessible) setActive_s(args ...interface{}) () {
  // setActive(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible9setActiveEb
    // invoke: void setActive(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible9setActiveEb(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "setActive", args)
  }

}

// updateAccessibility(class QAccessibleEvent *)
func (this *QAccessible) updateAccessibility_s(args ...interface{}) () {
  // updateAccessibility(class QAccessibleEvent *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleEvent{}) // "QAccessibleEvent *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent
    // invoke: void updateAccessibility(class QAccessibleEvent *)
    var arg0 = args[0].(QAccessibleEvent).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "updateAccessibility", args)
  }

}

// cleanup()
func (this *QAccessible) cleanup_s(args ...interface{}) () {
  // cleanup()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible7cleanupEv
    // invoke: void cleanup()
    C.C_ZN11QAccessible7cleanupEv()
  default:
    qtrt.ErrorResolve("QAccessible", "cleanup", args)
  }

}

// isActive()
func (this *QAccessible) isActive_s(args ...interface{}) () {
  // isActive()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible8isActiveEv
    // invoke: bool isActive()
    C.C_ZN11QAccessible8isActiveEv()
  default:
    qtrt.ErrorResolve("QAccessible", "isActive", args)
  }

}

// registerAccessibleInterface(class QAccessibleInterface *)
func (this *QAccessible) registerAccessibleInterface_s(args ...interface{}) () {
  // registerAccessibleInterface(class QAccessibleInterface *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface
    // invoke: Id registerAccessibleInterface(class QAccessibleInterface *)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "registerAccessibleInterface", args)
  }

}

// modelChangeType()
func (this *QAccessibleTableModelChangeEvent) modelChangeType(args ...interface{}) () {
  // modelChangeType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv
    // invoke: QAccessibleTableModelChangeEvent::ModelChangeType modelChangeType()
    C.C_ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "modelChangeType", args)
  }

}

// setFirstColumn(int)
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
    C.C_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstColumn", args)
  }

}

// setLastColumn(int)
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
    C.C_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastColumn", args)
  }

}

// lastColumn()
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
    C.C_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastColumn", args)
  }

}

// firstColumn()
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
    C.C_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstColumn", args)
  }

}

// setFirstRow(int)
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
    C.C_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstRow", args)
  }

}

// firstRow()
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
    C.C_ZNK32QAccessibleTableModelChangeEvent8firstRowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstRow", args)
  }

}

// setLastRow(int)
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
    C.C_ZN32QAccessibleTableModelChangeEvent10setLastRowEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastRow", args)
  }

}

// lastRow()
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
    C.C_ZNK32QAccessibleTableModelChangeEvent7lastRowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastRow", args)
  }

}

// accessibleInterface()
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
    C.C_ZNK16QAccessibleEvent19accessibleInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "accessibleInterface", args)
  }

}

// object()
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
    C.C_ZNK16QAccessibleEvent6objectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "object", args)
  }

}

// uniqueId()
func (this *QAccessibleEvent) uniqueId(args ...interface{}) () {
  // uniqueId()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent8uniqueIdEv
    // invoke: QAccessible::Id uniqueId()
    C.C_ZNK16QAccessibleEvent8uniqueIdEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "uniqueId", args)
  }

}

// child()
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
    C.C_ZNK16QAccessibleEvent5childEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "child", args)
  }

}

// setChild(int)
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
    C.C_ZN16QAccessibleEvent8setChildEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "setChild", args)
  }

}

// type()
func (this *QAccessibleEvent) type_(args ...interface{}) () {
  // type()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK16QAccessibleEvent4typeEv
    // invoke: QAccessible::Event type()
    C.C_ZNK16QAccessibleEvent4typeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "type", args)
  }

}

// ~QAccessibleEvent()
func (this *QAccessibleEvent) FreeQAccessibleEvent(args ...interface{}) () {
  // ~QAccessibleEvent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN16QAccessibleEventD0Ev
    // invoke: void ~QAccessibleEvent()
    C.C_ZN16QAccessibleEventD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "~QAccessibleEvent", args)
  }

}

// increaseAction()
func (this *QAccessibleActionInterface) increaseAction_s(args ...interface{}) () {
  // increaseAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface14increaseActionEv
    // invoke: const QString & increaseAction()
    C.C_ZN26QAccessibleActionInterface14increaseActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "increaseAction", args)
  }

}

// pressAction()
func (this *QAccessibleActionInterface) pressAction_s(args ...interface{}) () {
  // pressAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface11pressActionEv
    // invoke: const QString & pressAction()
    C.C_ZN26QAccessibleActionInterface11pressActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "pressAction", args)
  }

}

// previousPageAction()
func (this *QAccessibleActionInterface) previousPageAction_s(args ...interface{}) () {
  // previousPageAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface18previousPageActionEv
    // invoke: QString previousPageAction()
    C.C_ZN26QAccessibleActionInterface18previousPageActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "previousPageAction", args)
  }

}

// localizedActionName(const class QString &)
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
    C.C_ZNK26QAccessibleActionInterface19localizedActionNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionName", args)
  }

}

// scrollRightAction()
func (this *QAccessibleActionInterface) scrollRightAction_s(args ...interface{}) () {
  // scrollRightAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface17scrollRightActionEv
    // invoke: QString scrollRightAction()
    C.C_ZN26QAccessibleActionInterface17scrollRightActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollRightAction", args)
  }

}

// decreaseAction()
func (this *QAccessibleActionInterface) decreaseAction_s(args ...interface{}) () {
  // decreaseAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface14decreaseActionEv
    // invoke: const QString & decreaseAction()
    C.C_ZN26QAccessibleActionInterface14decreaseActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "decreaseAction", args)
  }

}

// showMenuAction()
func (this *QAccessibleActionInterface) showMenuAction_s(args ...interface{}) () {
  // showMenuAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface14showMenuActionEv
    // invoke: const QString & showMenuAction()
    C.C_ZN26QAccessibleActionInterface14showMenuActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "showMenuAction", args)
  }

}

// toggleAction()
func (this *QAccessibleActionInterface) toggleAction_s(args ...interface{}) () {
  // toggleAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface12toggleActionEv
    // invoke: const QString & toggleAction()
    C.C_ZN26QAccessibleActionInterface12toggleActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "toggleAction", args)
  }

}

// setFocusAction()
func (this *QAccessibleActionInterface) setFocusAction_s(args ...interface{}) () {
  // setFocusAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface14setFocusActionEv
    // invoke: const QString & setFocusAction()
    C.C_ZN26QAccessibleActionInterface14setFocusActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "setFocusAction", args)
  }

}

// scrollLeftAction()
func (this *QAccessibleActionInterface) scrollLeftAction_s(args ...interface{}) () {
  // scrollLeftAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface16scrollLeftActionEv
    // invoke: QString scrollLeftAction()
    C.C_ZN26QAccessibleActionInterface16scrollLeftActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollLeftAction", args)
  }

}

// localizedActionDescription(const class QString &)
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
    C.C_ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionDescription", args)
  }

}

// scrollUpAction()
func (this *QAccessibleActionInterface) scrollUpAction_s(args ...interface{}) () {
  // scrollUpAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface14scrollUpActionEv
    // invoke: QString scrollUpAction()
    C.C_ZN26QAccessibleActionInterface14scrollUpActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollUpAction", args)
  }

}

// scrollDownAction()
func (this *QAccessibleActionInterface) scrollDownAction_s(args ...interface{}) () {
  // scrollDownAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface16scrollDownActionEv
    // invoke: QString scrollDownAction()
    C.C_ZN26QAccessibleActionInterface16scrollDownActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollDownAction", args)
  }

}

// nextPageAction()
func (this *QAccessibleActionInterface) nextPageAction_s(args ...interface{}) () {
  // nextPageAction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleActionInterface14nextPageActionEv
    // invoke: QString nextPageAction()
    C.C_ZN26QAccessibleActionInterface14nextPageActionEv()
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "nextPageAction", args)
  }

}

// tableInterface()
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
    C.C_ZN20QAccessibleInterface14tableInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableInterface", args)
  }

}

// actionInterface()
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
    C.C_ZN20QAccessibleInterface15actionInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "actionInterface", args)
  }

}

// textInterface()
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
    C.C_ZN20QAccessibleInterface13textInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "textInterface", args)
  }

}

// foregroundColor()
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
    C.C_ZNK20QAccessibleInterface15foregroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "foregroundColor", args)
  }

}

// editableTextInterface()
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
    C.C_ZN20QAccessibleInterface21editableTextInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "editableTextInterface", args)
  }

}

// window()
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
    C.C_ZNK20QAccessibleInterface6windowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "window", args)
  }

}

// virtual_hook(int, void *)
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
    C.C_ZN20QAccessibleInterface12virtual_hookEiPv(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "virtual_hook", args)
  }

}

// backgroundColor()
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
    C.C_ZNK20QAccessibleInterface15backgroundColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "backgroundColor", args)
  }

}

// valueInterface()
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
    C.C_ZN20QAccessibleInterface14valueInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "valueInterface", args)
  }

}

// focusChild()
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
    C.C_ZNK20QAccessibleInterface10focusChildEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "focusChild", args)
  }

}

// imageInterface()
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
    C.C_ZN20QAccessibleInterface14imageInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "imageInterface", args)
  }

}

// tableCellInterface()
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
    C.C_ZN20QAccessibleInterface18tableCellInterfaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableCellInterface", args)
  }

}

// textInserted()
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
    C.C_ZNK26QAccessibleTextUpdateEvent12textInsertedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textInserted", args)
  }

}

// changePosition()
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
    C.C_ZNK26QAccessibleTextUpdateEvent14changePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "changePosition", args)
  }

}

// textRemoved()
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
    C.C_ZNK26QAccessibleTextUpdateEvent11textRemovedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textRemoved", args)
  }

}

// QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
func NewQAccessibleTextUpdateEvent(args ...interface{}) QAccessibleTextUpdateEvent {
  // QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
  // QAccessibleTextUpdateEvent(class QObject *, int, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextUpdateEventC1EP20QAccessibleInterfaceiRK7QStringS4_
    // invoke: void QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_(qthis, arg0, arg1, arg2, arg3)
  case 1:
    // invoke: _ZN26QAccessibleTextUpdateEventC1EP7QObjectiRK7QStringS4_
    // invoke: void QAccessibleTextUpdateEvent(class QObject *, int, const class QString &, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_(qthis, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "QAccessibleTextUpdateEvent", args)
  }

  return QAccessibleTextUpdateEvent{}
}

// changedStates()
func (this *QAccessibleStateChangeEvent) changedStates(args ...interface{}) () {
  // changedStates()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK27QAccessibleStateChangeEvent13changedStatesEv
    // invoke: QAccessible::State changedStates()
    C.C_ZNK27QAccessibleStateChangeEvent13changedStatesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleStateChangeEvent", "changedStates", args)
  }

}

// textInserted()
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
    C.C_ZNK26QAccessibleTextInsertEvent12textInsertedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "textInserted", args)
  }

}

// QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
func NewQAccessibleTextInsertEvent(args ...interface{}) QAccessibleTextInsertEvent {
  // QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
  // QAccessibleTextInsertEvent(class QObject *, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextInsertEventC1EP20QAccessibleInterfaceiRK7QString
    // invoke: void QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString(qthis, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN26QAccessibleTextInsertEventC1EP7QObjectiRK7QString
    // invoke: void QAccessibleTextInsertEvent(class QObject *, int, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "QAccessibleTextInsertEvent", args)
  }

  return QAccessibleTextInsertEvent{}
}

// changePosition()
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
    C.C_ZNK26QAccessibleTextInsertEvent14changePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "changePosition", args)
  }

}

// textRemoved()
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
    C.C_ZNK26QAccessibleTextRemoveEvent11textRemovedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "textRemoved", args)
  }

}

// QAccessibleTextRemoveEvent(class QObject *, int, const class QString &)
func NewQAccessibleTextRemoveEvent(args ...interface{}) QAccessibleTextRemoveEvent {
  // QAccessibleTextRemoveEvent(class QObject *, int, const class QString &)
  // QAccessibleTextRemoveEvent(class QAccessibleInterface *, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextRemoveEventC1EP7QObjectiRK7QString
    // invoke: void QAccessibleTextRemoveEvent(class QObject *, int, const class QString &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString(qthis, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN26QAccessibleTextRemoveEventC1EP20QAccessibleInterfaceiRK7QString
    // invoke: void QAccessibleTextRemoveEvent(class QAccessibleInterface *, int, const class QString &)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "QAccessibleTextRemoveEvent", args)
  }

  return QAccessibleTextRemoveEvent{}
}

// changePosition()
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
    C.C_ZNK26QAccessibleTextRemoveEvent14changePositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "changePosition", args)
  }

}

// QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
func NewQAccessibleTextSelectionEvent(args ...interface{}) QAccessibleTextSelectionEvent {
  // QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
  // QAccessibleTextSelectionEvent(class QObject *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QAccessibleTextSelectionEventC1EP20QAccessibleInterfaceii
    // invoke: void QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii(qthis, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN29QAccessibleTextSelectionEventC1EP7QObjectii
    // invoke: void QAccessibleTextSelectionEvent(class QObject *, int, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN29QAccessibleTextSelectionEventC2EP7QObjectii(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "QAccessibleTextSelectionEvent", args)
  }

  return QAccessibleTextSelectionEvent{}
}

// selectionStart()
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
    C.C_ZNK29QAccessibleTextSelectionEvent14selectionStartEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionStart", args)
  }

}

// selectionEnd()
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
    C.C_ZNK29QAccessibleTextSelectionEvent12selectionEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionEnd", args)
  }

}

// setSelection(int, int)
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
    C.C_ZN29QAccessibleTextSelectionEvent12setSelectionEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "setSelection", args)
  }

}

// cursorPosition()
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
    C.C_ZNK26QAccessibleTextCursorEvent14cursorPositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "cursorPosition", args)
  }

}

// QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
func NewQAccessibleTextCursorEvent(args ...interface{}) QAccessibleTextCursorEvent {
  // QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
  // QAccessibleTextCursorEvent(class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextCursorEventC1EP20QAccessibleInterfacei
    // invoke: void QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN26QAccessibleTextCursorEventC1EP7QObjecti
    // invoke: void QAccessibleTextCursorEvent(class QObject *, int)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN26QAccessibleTextCursorEventC2EP7QObjecti(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "QAccessibleTextCursorEvent", args)
  }

  return QAccessibleTextCursorEvent{}
}

// setCursorPosition(int)
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
    C.C_ZN26QAccessibleTextCursorEvent17setCursorPositionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "setCursorPosition", args)
  }

}

// QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
func NewQAccessibleValueChangeEvent(args ...interface{}) QAccessibleValueChangeEvent {
  // QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
  // QAccessibleValueChangeEvent(class QAccessibleInterface *, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[1][1] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAccessibleValueChangeEventC1EP7QObjectRK8QVariant
    // invoke: void QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
    var arg0 = args[0].(QObject).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN27QAccessibleValueChangeEventC2EP7QObjectRK8QVariant(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN27QAccessibleValueChangeEventC1EP20QAccessibleInterfaceRK8QVariant
    // invoke: void QAccessibleValueChangeEvent(class QAccessibleInterface *, const class QVariant &)
    var arg0 = args[0].(QAccessibleInterface).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QVariant).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN27QAccessibleValueChangeEventC2EP20QAccessibleInterfaceRK8QVariant(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "QAccessibleValueChangeEvent", args)
  }

  return QAccessibleValueChangeEvent{}
}

// setValue(const class QVariant &)
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
    C.C_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "setValue", args)
  }

}

// value()
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
    C.C_ZNK27QAccessibleValueChangeEvent5valueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "value", args)
  }

}

// <= body block end

