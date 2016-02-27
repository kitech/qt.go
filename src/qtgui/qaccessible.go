package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto: static QAccessibleInterface * QAccessible::accessibleInterface(Id uniqueId);
extern void* C_ZN11QAccessible19accessibleInterfaceEj(int32_t arg0); // 4
  // proto: static void QAccessible::deleteAccessibleInterface(Id uniqueId);
extern void C_ZN11QAccessible25deleteAccessibleInterfaceEj(int32_t arg0); // 4
  // proto: static void QAccessible::setRootObject(QObject * object);
extern void C_ZN11QAccessible13setRootObjectEP7QObject(void* arg0); // 4
  // proto: static QAccessibleInterface * QAccessible::queryAccessibleInterface(QObject * );
extern void* C_ZN11QAccessible24queryAccessibleInterfaceEP7QObject(void* arg0); // 4
  // proto: static Id QAccessible::uniqueId(QAccessibleInterface * iface);
extern int32_t C_ZN11QAccessible8uniqueIdEP20QAccessibleInterface(void* arg0); // 4
  // proto: static void QAccessible::setActive(bool active);
extern void C_ZN11QAccessible9setActiveEb(bool arg0); // 4
  // proto: static void QAccessible::updateAccessibility(QAccessibleEvent * event);
extern void C_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent(void* arg0); // 4
  // proto: static void QAccessible::cleanup();
extern void C_ZN11QAccessible7cleanupEv(); // 4
  // proto: static bool QAccessible::isActive();
extern bool C_ZN11QAccessible8isActiveEv(); // 4
  // proto: static Id QAccessible::registerAccessibleInterface(QAccessibleInterface * iface);
extern int32_t C_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface(void* arg0); // 4
  // proto:  QAccessibleTableModelChangeEvent::ModelChangeType QAccessibleTableModelChangeEvent::modelChangeType();
extern void C_ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv(void* qthis); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setFirstColumn(int col);
extern void C_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(void* qthis, int32_t arg0); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setLastColumn(int col);
extern void C_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(void* qthis, int32_t arg0); // 2
  // proto:  int QAccessibleTableModelChangeEvent::lastColumn();
extern int32_t C_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(void* qthis); // 2
  // proto:  int QAccessibleTableModelChangeEvent::firstColumn();
extern int32_t C_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(void* qthis); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setFirstRow(int row);
extern void C_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(void* qthis, int32_t arg0); // 2
  // proto:  int QAccessibleTableModelChangeEvent::firstRow();
extern int32_t C_ZNK32QAccessibleTableModelChangeEvent8firstRowEv(void* qthis); // 2
  // proto:  void QAccessibleTableModelChangeEvent::setLastRow(int row);
extern void C_ZN32QAccessibleTableModelChangeEvent10setLastRowEi(void* qthis, int32_t arg0); // 2
  // proto:  int QAccessibleTableModelChangeEvent::lastRow();
extern int32_t C_ZNK32QAccessibleTableModelChangeEvent7lastRowEv(void* qthis); // 2
  // proto:  QAccessibleInterface * QAccessibleEvent::accessibleInterface();
extern void* C_ZNK16QAccessibleEvent19accessibleInterfaceEv(void* qthis); // 4
  // proto:  QObject * QAccessibleEvent::object();
extern void* C_ZNK16QAccessibleEvent6objectEv(void* qthis); // 2
  // proto:  QAccessible::Id QAccessibleEvent::uniqueId();
extern void C_ZNK16QAccessibleEvent8uniqueIdEv(void* qthis); // 4
  // proto:  int QAccessibleEvent::child();
extern int32_t C_ZNK16QAccessibleEvent5childEv(void* qthis); // 2
  // proto:  void QAccessibleEvent::setChild(int chld);
extern void C_ZN16QAccessibleEvent8setChildEi(void* qthis, int32_t arg0); // 2
  // proto:  QAccessible::Event QAccessibleEvent::type();
extern void C_ZNK16QAccessibleEvent4typeEv(void* qthis); // 2
  // proto:  void QAccessibleEvent::~QAccessibleEvent();
extern void C_ZN16QAccessibleEventD2Ev(void* qthis); // 4
  // proto: static const QString & QAccessibleActionInterface::increaseAction();
extern void* C_ZN26QAccessibleActionInterface14increaseActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::pressAction();
extern void* C_ZN26QAccessibleActionInterface11pressActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::previousPageAction();
extern void* C_ZN26QAccessibleActionInterface18previousPageActionEv(); // 4
  // proto:  QString QAccessibleActionInterface::localizedActionName(const QString & name);
extern void* C_ZNK26QAccessibleActionInterface19localizedActionNameERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QAccessibleActionInterface::scrollRightAction();
extern void* C_ZN26QAccessibleActionInterface17scrollRightActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::decreaseAction();
extern void* C_ZN26QAccessibleActionInterface14decreaseActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::showMenuAction();
extern void* C_ZN26QAccessibleActionInterface14showMenuActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::toggleAction();
extern void* C_ZN26QAccessibleActionInterface12toggleActionEv(); // 4
  // proto: static const QString & QAccessibleActionInterface::setFocusAction();
extern void* C_ZN26QAccessibleActionInterface14setFocusActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::scrollLeftAction();
extern void* C_ZN26QAccessibleActionInterface16scrollLeftActionEv(); // 4
  // proto:  QString QAccessibleActionInterface::localizedActionDescription(const QString & name);
extern void* C_ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto: static QString QAccessibleActionInterface::scrollUpAction();
extern void* C_ZN26QAccessibleActionInterface14scrollUpActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::scrollDownAction();
extern void* C_ZN26QAccessibleActionInterface16scrollDownActionEv(); // 4
  // proto: static QString QAccessibleActionInterface::nextPageAction();
extern void* C_ZN26QAccessibleActionInterface14nextPageActionEv(); // 4
  // proto:  QAccessibleTableInterface * QAccessibleInterface::tableInterface();
extern void* C_ZN20QAccessibleInterface14tableInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleActionInterface * QAccessibleInterface::actionInterface();
extern void* C_ZN20QAccessibleInterface15actionInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleTextInterface * QAccessibleInterface::textInterface();
extern void* C_ZN20QAccessibleInterface13textInterfaceEv(void* qthis); // 2
  // proto:  QColor QAccessibleInterface::foregroundColor();
extern void* C_ZNK20QAccessibleInterface15foregroundColorEv(void* qthis); // 4
  // proto:  QAccessibleEditableTextInterface * QAccessibleInterface::editableTextInterface();
extern void* C_ZN20QAccessibleInterface21editableTextInterfaceEv(void* qthis); // 2
  // proto:  QWindow * QAccessibleInterface::window();
extern void* C_ZNK20QAccessibleInterface6windowEv(void* qthis); // 4
  // proto:  void QAccessibleInterface::virtual_hook(int id, void * data);
extern void C_ZN20QAccessibleInterface12virtual_hookEiPv(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QColor QAccessibleInterface::backgroundColor();
extern void* C_ZNK20QAccessibleInterface15backgroundColorEv(void* qthis); // 4
  // proto:  QAccessibleValueInterface * QAccessibleInterface::valueInterface();
extern void* C_ZN20QAccessibleInterface14valueInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleInterface * QAccessibleInterface::focusChild();
extern void* C_ZNK20QAccessibleInterface10focusChildEv(void* qthis); // 4
  // proto:  QAccessibleImageInterface * QAccessibleInterface::imageInterface();
extern void* C_ZN20QAccessibleInterface14imageInterfaceEv(void* qthis); // 2
  // proto:  QAccessibleTableCellInterface * QAccessibleInterface::tableCellInterface();
extern void* C_ZN20QAccessibleInterface18tableCellInterfaceEv(void* qthis); // 2
  // proto:  QString QAccessibleTextUpdateEvent::textInserted();
extern void* C_ZNK26QAccessibleTextUpdateEvent12textInsertedEv(void* qthis); // 2
  // proto:  int QAccessibleTextUpdateEvent::changePosition();
extern int32_t C_ZNK26QAccessibleTextUpdateEvent14changePositionEv(void* qthis); // 2
  // proto:  QString QAccessibleTextUpdateEvent::textRemoved();
extern void* C_ZNK26QAccessibleTextUpdateEvent11textRemovedEv(void* qthis); // 2
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QAccessibleInterface * iface, int position, const QString & oldText, const QString & text);
extern void* C_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_(void* arg0, int32_t arg1, void* arg2, void* arg3); // 1
  // proto:  void QAccessibleTextUpdateEvent::QAccessibleTextUpdateEvent(QObject * obj, int position, const QString & oldText, const QString & text);
extern void* C_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_(void* arg0, int32_t arg1, void* arg2, void* arg3); // 1
  // proto:  QAccessible::State QAccessibleStateChangeEvent::changedStates();
extern void C_ZNK27QAccessibleStateChangeEvent13changedStatesEv(void* qthis); // 2
  // proto:  QString QAccessibleTextInsertEvent::textInserted();
extern void* C_ZNK26QAccessibleTextInsertEvent12textInsertedEv(void* qthis); // 2
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void* C_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString(void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  void QAccessibleTextInsertEvent::QAccessibleTextInsertEvent(QObject * obj, int position, const QString & text);
extern void* C_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString(void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  int QAccessibleTextInsertEvent::changePosition();
extern int32_t C_ZNK26QAccessibleTextInsertEvent14changePositionEv(void* qthis); // 2
  // proto:  QString QAccessibleTextRemoveEvent::textRemoved();
extern void* C_ZNK26QAccessibleTextRemoveEvent11textRemovedEv(void* qthis); // 2
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QObject * obj, int position, const QString & text);
extern void* C_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString(void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  void QAccessibleTextRemoveEvent::QAccessibleTextRemoveEvent(QAccessibleInterface * iface, int position, const QString & text);
extern void* C_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString(void* arg0, int32_t arg1, void* arg2); // 1
  // proto:  int QAccessibleTextRemoveEvent::changePosition();
extern int32_t C_ZNK26QAccessibleTextRemoveEvent14changePositionEv(void* qthis); // 2
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QAccessibleInterface * iface, int start, int end);
extern void* C_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii(void* arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  void QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent(QObject * obj, int start, int end);
extern void* C_ZN29QAccessibleTextSelectionEventC2EP7QObjectii(void* arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  int QAccessibleTextSelectionEvent::selectionStart();
extern int32_t C_ZNK29QAccessibleTextSelectionEvent14selectionStartEv(void* qthis); // 2
  // proto:  int QAccessibleTextSelectionEvent::selectionEnd();
extern int32_t C_ZNK29QAccessibleTextSelectionEvent12selectionEndEv(void* qthis); // 2
  // proto:  void QAccessibleTextSelectionEvent::setSelection(int start, int end);
extern void C_ZN29QAccessibleTextSelectionEvent12setSelectionEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  int QAccessibleTextCursorEvent::cursorPosition();
extern int32_t C_ZNK26QAccessibleTextCursorEvent14cursorPositionEv(void* qthis); // 2
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QAccessibleInterface * iface, int cursorPos);
extern void* C_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei(void* arg0, int32_t arg1); // 1
  // proto:  void QAccessibleTextCursorEvent::QAccessibleTextCursorEvent(QObject * obj, int cursorPos);
extern void* C_ZN26QAccessibleTextCursorEventC2EP7QObjecti(void* arg0, int32_t arg1); // 1
  // proto:  void QAccessibleTextCursorEvent::setCursorPosition(int position);
extern void C_ZN26QAccessibleTextCursorEvent17setCursorPositionEi(void* qthis, int32_t arg0); // 2
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QObject * obj, const QVariant & val);
extern void* C_ZN27QAccessibleValueChangeEventC2EP7QObjectRK8QVariant(void* arg0, void* arg1); // 1
  // proto:  void QAccessibleValueChangeEvent::QAccessibleValueChangeEvent(QAccessibleInterface * iface, const QVariant & val);
extern void* C_ZN27QAccessibleValueChangeEventC2EP20QAccessibleInterfaceRK8QVariant(void* arg0, void* arg1); // 1
  // proto:  void QAccessibleValueChangeEvent::setValue(const QVariant & val);
extern void C_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(void* qthis, void* arg0); // 2
  // proto:  QVariant QAccessibleValueChangeEvent::value();
extern void* C_ZNK27QAccessibleValueChangeEvent5valueEv(void* qthis); // 2
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QAccessible)=1
type QAccessible struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTableModelChangeEvent)=48
type QAccessibleTableModelChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextInterface)=8
type QAccessibleTextInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleEvent)=32
type QAccessibleEvent struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleActionInterface)=8
type QAccessibleActionInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleInterface)=8
type QAccessibleInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleEditableTextInterface)=8
type QAccessibleEditableTextInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTableCellInterface)=8
type QAccessibleTableCellInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTableInterface)=8
type QAccessibleTableInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextUpdateEvent)=56
type QAccessibleTextUpdateEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleStateChangeEvent)=40
type QAccessibleStateChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleImageInterface)=8
type QAccessibleImageInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextInsertEvent)=48
type QAccessibleTextInsertEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleValueInterface)=8
type QAccessibleValueInterface struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextRemoveEvent)=48
type QAccessibleTextRemoveEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextSelectionEvent)=40
type QAccessibleTextSelectionEvent struct {
  /*qbase*/ QAccessibleTextCursorEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleTextCursorEvent)=32
type QAccessibleTextCursorEvent struct {
  /*qbase*/ QAccessibleEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QAccessibleValueChangeEvent)=48
type QAccessibleValueChangeEvent struct {
  /*qbase*/ QAccessibleEvent;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// accessibleInterface(Id)
func (this *QAccessible) AccessibleInterface_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QAccessible19accessibleInterfaceEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessible", "accessibleInterface", args)
  }

  return
}

// deleteAccessibleInterface(Id)
func (this *QAccessible) DeleteAccessibleInterface_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible25deleteAccessibleInterfaceEj(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "deleteAccessibleInterface", args)
  }

  return
}

// setRootObject(class QObject *)
func (this *QAccessible) SetRootObject_s(args ...interface{}) () {
  // setRootObject(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible13setRootObjectEP7QObject
    // invoke: void setRootObject(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible13setRootObjectEP7QObject(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "setRootObject", args)
  }

  return
}

// queryAccessibleInterface(class QObject *)
func (this *QAccessible) QueryAccessibleInterface_s(args ...interface{}) (ret interface{}) {
  // queryAccessibleInterface(class QObject *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QAccessible24queryAccessibleInterfaceEP7QObject
    // invoke: QAccessibleInterface * queryAccessibleInterface(class QObject *)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QAccessible24queryAccessibleInterfaceEP7QObject(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessible", "queryAccessibleInterface", args)
  }

  return
}

// uniqueId(class QAccessibleInterface *)
func (this *QAccessible) UniqueId_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QAccessible8uniqueIdEP20QAccessibleInterface(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "Id"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessible", "uniqueId", args)
  }

  return
}

// setActive(_Bool)
func (this *QAccessible) SetActive_s(args ...interface{}) () {
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

  return
}

// updateAccessibility(class QAccessibleEvent *)
func (this *QAccessible) UpdateAccessibility_s(args ...interface{}) () {
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
    var arg0 = args[0].(*QAccessibleEvent).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent(arg0)
  default:
    qtrt.ErrorResolve("QAccessible", "updateAccessibility", args)
  }

  return
}

// cleanup()
func (this *QAccessible) Cleanup_s(args ...interface{}) () {
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

  return
}

// isActive()
func (this *QAccessible) IsActive_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QAccessible8isActiveEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessible", "isActive", args)
  }

  return
}

// registerAccessibleInterface(class QAccessibleInterface *)
func (this *QAccessible) RegisterAccessibleInterface_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "Id"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessible", "registerAccessibleInterface", args)
  }

  return
}

// modelChangeType()
func (this *QAccessibleTableModelChangeEvent) ModelChangeType(args ...interface{}) () {
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
    C.C_ZNK32QAccessibleTableModelChangeEvent15modelChangeTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "modelChangeType", args)
  }

  return
}

// setFirstColumn(int)
func (this *QAccessibleTableModelChangeEvent) SetFirstColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN32QAccessibleTableModelChangeEvent14setFirstColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstColumn", args)
  }

  return
}

// setLastColumn(int)
func (this *QAccessibleTableModelChangeEvent) SetLastColumn(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN32QAccessibleTableModelChangeEvent13setLastColumnEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastColumn", args)
  }

  return
}

// lastColumn()
func (this *QAccessibleTableModelChangeEvent) LastColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK32QAccessibleTableModelChangeEvent10lastColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastColumn", args)
  }

  return
}

// firstColumn()
func (this *QAccessibleTableModelChangeEvent) FirstColumn(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK32QAccessibleTableModelChangeEvent11firstColumnEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstColumn", args)
  }

  return
}

// setFirstRow(int)
func (this *QAccessibleTableModelChangeEvent) SetFirstRow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN32QAccessibleTableModelChangeEvent11setFirstRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setFirstRow", args)
  }

  return
}

// firstRow()
func (this *QAccessibleTableModelChangeEvent) FirstRow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK32QAccessibleTableModelChangeEvent8firstRowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "firstRow", args)
  }

  return
}

// setLastRow(int)
func (this *QAccessibleTableModelChangeEvent) SetLastRow(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN32QAccessibleTableModelChangeEvent10setLastRowEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "setLastRow", args)
  }

  return
}

// lastRow()
func (this *QAccessibleTableModelChangeEvent) LastRow(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK32QAccessibleTableModelChangeEvent7lastRowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTableModelChangeEvent", "lastRow", args)
  }

  return
}

// accessibleInterface()
func (this *QAccessibleEvent) AccessibleInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QAccessibleEvent19accessibleInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "accessibleInterface", args)
  }

  return
}

// object()
func (this *QAccessibleEvent) Object(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QAccessibleEvent6objectEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "object", args)
  }

  return
}

// uniqueId()
func (this *QAccessibleEvent) UniqueId(args ...interface{}) () {
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
    C.C_ZNK16QAccessibleEvent8uniqueIdEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "uniqueId", args)
  }

  return
}

// child()
func (this *QAccessibleEvent) Child(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK16QAccessibleEvent5childEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "child", args)
  }

  return
}

// setChild(int)
func (this *QAccessibleEvent) SetChild(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN16QAccessibleEvent8setChildEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "setChild", args)
  }

  return
}

// type()
func (this *QAccessibleEvent) Type_(args ...interface{}) () {
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
    C.C_ZNK16QAccessibleEvent4typeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "type", args)
  }

  return
}

// ~QAccessibleEvent()
func (this *QAccessibleEvent) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN16QAccessibleEventD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QAccessibleEvent", "~QAccessibleEvent", args)
  }

  return
}

// increaseAction()
func (this *QAccessibleActionInterface) IncreaseAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface14increaseActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "increaseAction", args)
  }

  return
}

// pressAction()
func (this *QAccessibleActionInterface) PressAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface11pressActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "pressAction", args)
  }

  return
}

// previousPageAction()
func (this *QAccessibleActionInterface) PreviousPageAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface18previousPageActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "previousPageAction", args)
  }

  return
}

// localizedActionName(const class QString &)
func (this *QAccessibleActionInterface) LocalizedActionName(args ...interface{}) (ret interface{}) {
  // localizedActionName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface19localizedActionNameERK7QString
    // invoke: QString localizedActionName(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK26QAccessibleActionInterface19localizedActionNameERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionName", args)
  }

  return
}

// scrollRightAction()
func (this *QAccessibleActionInterface) ScrollRightAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface17scrollRightActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollRightAction", args)
  }

  return
}

// decreaseAction()
func (this *QAccessibleActionInterface) DecreaseAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface14decreaseActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "decreaseAction", args)
  }

  return
}

// showMenuAction()
func (this *QAccessibleActionInterface) ShowMenuAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface14showMenuActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "showMenuAction", args)
  }

  return
}

// toggleAction()
func (this *QAccessibleActionInterface) ToggleAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface12toggleActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "toggleAction", args)
  }

  return
}

// setFocusAction()
func (this *QAccessibleActionInterface) SetFocusAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface14setFocusActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "const QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "setFocusAction", args)
  }

  return
}

// scrollLeftAction()
func (this *QAccessibleActionInterface) ScrollLeftAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface16scrollLeftActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollLeftAction", args)
  }

  return
}

// localizedActionDescription(const class QString &)
func (this *QAccessibleActionInterface) LocalizedActionDescription(args ...interface{}) (ret interface{}) {
  // localizedActionDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString
    // invoke: QString localizedActionDescription(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK26QAccessibleActionInterface26localizedActionDescriptionERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "localizedActionDescription", args)
  }

  return
}

// scrollUpAction()
func (this *QAccessibleActionInterface) ScrollUpAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface14scrollUpActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollUpAction", args)
  }

  return
}

// scrollDownAction()
func (this *QAccessibleActionInterface) ScrollDownAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface16scrollDownActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "scrollDownAction", args)
  }

  return
}

// nextPageAction()
func (this *QAccessibleActionInterface) NextPageAction_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN26QAccessibleActionInterface14nextPageActionEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleActionInterface", "nextPageAction", args)
  }

  return
}

// tableInterface()
func (this *QAccessibleInterface) TableInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface14tableInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleTableInterface{}) // "QAccessibleTableInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableInterface", args)
  }

  return
}

// actionInterface()
func (this *QAccessibleInterface) ActionInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface15actionInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleActionInterface{}) // "QAccessibleActionInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "actionInterface", args)
  }

  return
}

// textInterface()
func (this *QAccessibleInterface) TextInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface13textInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleTextInterface{}) // "QAccessibleTextInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "textInterface", args)
  }

  return
}

// foregroundColor()
func (this *QAccessibleInterface) ForegroundColor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK20QAccessibleInterface15foregroundColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "foregroundColor", args)
  }

  return
}

// editableTextInterface()
func (this *QAccessibleInterface) EditableTextInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface21editableTextInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleEditableTextInterface{}) // "QAccessibleEditableTextInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "editableTextInterface", args)
  }

  return
}

// window()
func (this *QAccessibleInterface) Window(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK20QAccessibleInterface6windowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QWindow{}) // "QWindow *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "window", args)
  }

  return
}

// virtual_hook(int, void *)
func (this *QAccessibleInterface) Virtual_hook(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(unsafe.Pointer)
    if false {fmt.Println(arg1)}
    C.C_ZN20QAccessibleInterface12virtual_hookEiPv(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "virtual_hook", args)
  }

  return
}

// backgroundColor()
func (this *QAccessibleInterface) BackgroundColor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK20QAccessibleInterface15backgroundColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "backgroundColor", args)
  }

  return
}

// valueInterface()
func (this *QAccessibleInterface) ValueInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface14valueInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleValueInterface{}) // "QAccessibleValueInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "valueInterface", args)
  }

  return
}

// focusChild()
func (this *QAccessibleInterface) FocusChild(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK20QAccessibleInterface10focusChildEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "focusChild", args)
  }

  return
}

// imageInterface()
func (this *QAccessibleInterface) ImageInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface14imageInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleImageInterface{}) // "QAccessibleImageInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "imageInterface", args)
  }

  return
}

// tableCellInterface()
func (this *QAccessibleInterface) TableCellInterface(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN20QAccessibleInterface18tableCellInterfaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QAccessibleTableCellInterface{}) // "QAccessibleTableCellInterface *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleInterface", "tableCellInterface", args)
  }

  return
}

// textInserted()
func (this *QAccessibleTextUpdateEvent) TextInserted(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextUpdateEvent12textInsertedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textInserted", args)
  }

  return
}

// changePosition()
func (this *QAccessibleTextUpdateEvent) ChangePosition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextUpdateEvent14changePositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "changePosition", args)
  }

  return
}

// textRemoved()
func (this *QAccessibleTextUpdateEvent) TextRemoved(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextUpdateEvent11textRemovedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "textRemoved", args)
  }

  return
}

// QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
func GcfreeQAccessibleTextUpdateEvent(this *QAccessibleTextUpdateEvent) {
  qtrt.UniverseFree(this)
}
func NewQAccessibleTextUpdateEvent(args ...interface{}) *QAccessibleTextUpdateEvent {
  // QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
  // QAccessibleTextUpdateEvent(class QObject *, int, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextUpdateEventC1EP20QAccessibleInterfaceiRK7QStringS4_
    // invoke: void QAccessibleTextUpdateEvent(class QAccessibleInterface *, int, const class QString &, const class QString &)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextUpdateEventC2EP20QAccessibleInterfaceiRK7QStringS4_(arg0, arg1, arg2, arg3)
    this := &QAccessibleTextUpdateEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextUpdateEvent)
    return this
  case 1:
    // invoke: _ZN26QAccessibleTextUpdateEventC1EP7QObjectiRK7QStringS4_
    // invoke: void QAccessibleTextUpdateEvent(class QObject *, int, const class QString &, const class QString &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextUpdateEventC2EP7QObjectiRK7QStringS4_(arg0, arg1, arg2, arg3)
    this := &QAccessibleTextUpdateEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextUpdateEvent)
    return this
  default:
    qtrt.ErrorResolve("QAccessibleTextUpdateEvent", "QAccessibleTextUpdateEvent", args)
  }

  return nil // QAccessibleTextUpdateEvent{Qclsinst:qthis}
}

// changedStates()
func (this *QAccessibleStateChangeEvent) ChangedStates(args ...interface{}) () {
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
    C.C_ZNK27QAccessibleStateChangeEvent13changedStatesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QAccessibleStateChangeEvent", "changedStates", args)
  }

  return
}

// textInserted()
func (this *QAccessibleTextInsertEvent) TextInserted(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextInsertEvent12textInsertedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "textInserted", args)
  }

  return
}

// QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
func GcfreeQAccessibleTextInsertEvent(this *QAccessibleTextInsertEvent) {
  qtrt.UniverseFree(this)
}
func NewQAccessibleTextInsertEvent(args ...interface{}) *QAccessibleTextInsertEvent {
  // QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
  // QAccessibleTextInsertEvent(class QObject *, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextInsertEventC1EP20QAccessibleInterfaceiRK7QString
    // invoke: void QAccessibleTextInsertEvent(class QAccessibleInterface *, int, const class QString &)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextInsertEventC2EP20QAccessibleInterfaceiRK7QString(arg0, arg1, arg2)
    this := &QAccessibleTextInsertEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextInsertEvent)
    return this
  case 1:
    // invoke: _ZN26QAccessibleTextInsertEventC1EP7QObjectiRK7QString
    // invoke: void QAccessibleTextInsertEvent(class QObject *, int, const class QString &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextInsertEventC2EP7QObjectiRK7QString(arg0, arg1, arg2)
    this := &QAccessibleTextInsertEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextInsertEvent)
    return this
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "QAccessibleTextInsertEvent", args)
  }

  return nil // QAccessibleTextInsertEvent{Qclsinst:qthis}
}

// changePosition()
func (this *QAccessibleTextInsertEvent) ChangePosition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextInsertEvent14changePositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextInsertEvent", "changePosition", args)
  }

  return
}

// textRemoved()
func (this *QAccessibleTextRemoveEvent) TextRemoved(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextRemoveEvent11textRemovedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "textRemoved", args)
  }

  return
}

// QAccessibleTextRemoveEvent(class QObject *, int, const class QString &)
func GcfreeQAccessibleTextRemoveEvent(this *QAccessibleTextRemoveEvent) {
  qtrt.UniverseFree(this)
}
func NewQAccessibleTextRemoveEvent(args ...interface{}) *QAccessibleTextRemoveEvent {
  // QAccessibleTextRemoveEvent(class QObject *, int, const class QString &)
  // QAccessibleTextRemoveEvent(class QAccessibleInterface *, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextRemoveEventC1EP7QObjectiRK7QString
    // invoke: void QAccessibleTextRemoveEvent(class QObject *, int, const class QString &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextRemoveEventC2EP7QObjectiRK7QString(arg0, arg1, arg2)
    this := &QAccessibleTextRemoveEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextRemoveEvent)
    return this
  case 1:
    // invoke: _ZN26QAccessibleTextRemoveEventC1EP20QAccessibleInterfaceiRK7QString
    // invoke: void QAccessibleTextRemoveEvent(class QAccessibleInterface *, int, const class QString &)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextRemoveEventC2EP20QAccessibleInterfaceiRK7QString(arg0, arg1, arg2)
    this := &QAccessibleTextRemoveEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextRemoveEvent)
    return this
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "QAccessibleTextRemoveEvent", args)
  }

  return nil // QAccessibleTextRemoveEvent{Qclsinst:qthis}
}

// changePosition()
func (this *QAccessibleTextRemoveEvent) ChangePosition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextRemoveEvent14changePositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextRemoveEvent", "changePosition", args)
  }

  return
}

// QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
func GcfreeQAccessibleTextSelectionEvent(this *QAccessibleTextSelectionEvent) {
  qtrt.UniverseFree(this)
}
func NewQAccessibleTextSelectionEvent(args ...interface{}) *QAccessibleTextSelectionEvent {
  // QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
  // QAccessibleTextSelectionEvent(class QObject *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN29QAccessibleTextSelectionEventC1EP20QAccessibleInterfaceii
    // invoke: void QAccessibleTextSelectionEvent(class QAccessibleInterface *, int, int)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN29QAccessibleTextSelectionEventC2EP20QAccessibleInterfaceii(arg0, arg1, arg2)
    this := &QAccessibleTextSelectionEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextSelectionEvent)
    return this
  case 1:
    // invoke: _ZN29QAccessibleTextSelectionEventC1EP7QObjectii
    // invoke: void QAccessibleTextSelectionEvent(class QObject *, int, int)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN29QAccessibleTextSelectionEventC2EP7QObjectii(arg0, arg1, arg2)
    this := &QAccessibleTextSelectionEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextSelectionEvent)
    return this
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "QAccessibleTextSelectionEvent", args)
  }

  return nil // QAccessibleTextSelectionEvent{Qclsinst:qthis}
}

// selectionStart()
func (this *QAccessibleTextSelectionEvent) SelectionStart(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK29QAccessibleTextSelectionEvent14selectionStartEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionStart", args)
  }

  return
}

// selectionEnd()
func (this *QAccessibleTextSelectionEvent) SelectionEnd(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK29QAccessibleTextSelectionEvent12selectionEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "selectionEnd", args)
  }

  return
}

// setSelection(int, int)
func (this *QAccessibleTextSelectionEvent) SetSelection(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN29QAccessibleTextSelectionEvent12setSelectionEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QAccessibleTextSelectionEvent", "setSelection", args)
  }

  return
}

// cursorPosition()
func (this *QAccessibleTextCursorEvent) CursorPosition(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK26QAccessibleTextCursorEvent14cursorPositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "cursorPosition", args)
  }

  return
}

// QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
func GcfreeQAccessibleTextCursorEvent(this *QAccessibleTextCursorEvent) {
  qtrt.UniverseFree(this)
}
func NewQAccessibleTextCursorEvent(args ...interface{}) *QAccessibleTextCursorEvent {
  // QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
  // QAccessibleTextCursorEvent(class QObject *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN26QAccessibleTextCursorEventC1EP20QAccessibleInterfacei
    // invoke: void QAccessibleTextCursorEvent(class QAccessibleInterface *, int)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextCursorEventC2EP20QAccessibleInterfacei(arg0, arg1)
    this := &QAccessibleTextCursorEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextCursorEvent)
    return this
  case 1:
    // invoke: _ZN26QAccessibleTextCursorEventC1EP7QObjecti
    // invoke: void QAccessibleTextCursorEvent(class QObject *, int)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN26QAccessibleTextCursorEventC2EP7QObjecti(arg0, arg1)
    this := &QAccessibleTextCursorEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleTextCursorEvent)
    return this
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "QAccessibleTextCursorEvent", args)
  }

  return nil // QAccessibleTextCursorEvent{Qclsinst:qthis}
}

// setCursorPosition(int)
func (this *QAccessibleTextCursorEvent) SetCursorPosition(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN26QAccessibleTextCursorEvent17setCursorPositionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleTextCursorEvent", "setCursorPosition", args)
  }

  return
}

// QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
func GcfreeQAccessibleValueChangeEvent(this *QAccessibleValueChangeEvent) {
  qtrt.UniverseFree(this)
}
func NewQAccessibleValueChangeEvent(args ...interface{}) *QAccessibleValueChangeEvent {
  // QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
  // QAccessibleValueChangeEvent(class QAccessibleInterface *, const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QObject{}) // "QObject *"
  vtys[0][1] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QAccessibleInterface{}) // "QAccessibleInterface *"
  vtys[1][1] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAccessibleValueChangeEventC1EP7QObjectRK8QVariant
    // invoke: void QAccessibleValueChangeEvent(class QObject *, const class QVariant &)
    var arg0 = args[0].(*qtcore.QObject).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QAccessibleValueChangeEventC2EP7QObjectRK8QVariant(arg0, arg1)
    this := &QAccessibleValueChangeEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleValueChangeEvent)
    return this
  case 1:
    // invoke: _ZN27QAccessibleValueChangeEventC1EP20QAccessibleInterfaceRK8QVariant
    // invoke: void QAccessibleValueChangeEvent(class QAccessibleInterface *, const class QVariant &)
    var arg0 = args[0].(*QAccessibleInterface).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN27QAccessibleValueChangeEventC2EP20QAccessibleInterfaceRK8QVariant(arg0, arg1)
    this := &QAccessibleValueChangeEvent{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQAccessibleValueChangeEvent)
    return this
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "QAccessibleValueChangeEvent", args)
  }

  return nil // QAccessibleValueChangeEvent{Qclsinst:qthis}
}

// setValue(const class QVariant &)
func (this *QAccessibleValueChangeEvent) SetValue(args ...interface{}) () {
  // setValue(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN27QAccessibleValueChangeEvent8setValueERK8QVariant
    // invoke: void setValue(const class QVariant &)
    var arg0 = args[0].(*qtcore.QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN27QAccessibleValueChangeEvent8setValueERK8QVariant(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "setValue", args)
  }

  return
}

// value()
func (this *QAccessibleValueChangeEvent) Value(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK27QAccessibleValueChangeEvent5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QAccessibleValueChangeEvent", "value", args)
  }

  return
}

// <= body block end

