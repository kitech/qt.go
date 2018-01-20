//  header block begin
// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QAccessible struct {
	*qtrt.CObject
}

func (this *QAccessible) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQAccessibleFromPointer(cthis unsafe.Pointer) *QAccessible {
	return &QAccessible{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:411
// index:0
// Public static
// void installActivationObserver(class QAccessible::ActivationObserver *)
func (this *QAccessible) InstallActivationObserver(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible25installActivationObserverEPNS_18ActivationObserverE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QAccessible_InstallActivationObserver(arg0 unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.InstallActivationObserver(arg0)
}

// /usr/include/qt/QtGui/qaccessible.h:412
// index:0
// Public static
// void removeActivationObserver(class QAccessible::ActivationObserver *)
func (this *QAccessible) RemoveActivationObserver(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible24removeActivationObserverEPNS_18ActivationObserverE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QAccessible_RemoveActivationObserver(arg0 unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.RemoveActivationObserver(arg0)
}

// /usr/include/qt/QtGui/qaccessible.h:414
// index:0
// Public static
// QAccessibleInterface * queryAccessibleInterface(class QObject *)
func (this *QAccessible) QueryAccessibleInterface(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible24queryAccessibleInterfaceEP7QObject", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAccessible_QueryAccessibleInterface(arg0 unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.QueryAccessibleInterface(arg0)
}

// /usr/include/qt/QtGui/qaccessible.h:415
// index:0
// Public static
// QAccessible::Id uniqueId(class QAccessibleInterface *)
func (this *QAccessible) UniqueId(iface unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible8uniqueIdEP20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, iface)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAccessible_UniqueId(iface unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.UniqueId(iface)
}

// /usr/include/qt/QtGui/qaccessible.h:416
// index:0
// Public static
// QAccessibleInterface * accessibleInterface(QAccessible::Id)
func (this *QAccessible) AccessibleInterface(uniqueId uint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible19accessibleInterfaceEj", ffiqt.FFI_TYPE_POINTER, uniqueId)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAccessible_AccessibleInterface(uniqueId uint) {
	var nilthis *QAccessible
	nilthis.AccessibleInterface(uniqueId)
}

// /usr/include/qt/QtGui/qaccessible.h:417
// index:0
// Public static
// QAccessible::Id registerAccessibleInterface(class QAccessibleInterface *)
func (this *QAccessible) RegisterAccessibleInterface(iface unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, iface)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAccessible_RegisterAccessibleInterface(iface unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.RegisterAccessibleInterface(iface)
}

// /usr/include/qt/QtGui/qaccessible.h:418
// index:0
// Public static
// void deleteAccessibleInterface(QAccessible::Id)
func (this *QAccessible) DeleteAccessibleInterface(uniqueId uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible25deleteAccessibleInterfaceEj", ffiqt.FFI_TYPE_POINTER, uniqueId)
	gopp.ErrPrint(err, rv)
}
func QAccessible_DeleteAccessibleInterface(uniqueId uint) {
	var nilthis *QAccessible
	nilthis.DeleteAccessibleInterface(uniqueId)
}

// /usr/include/qt/QtGui/qaccessible.h:424
// index:0
// Public static
// void updateAccessibility(class QAccessibleEvent *)
func (this *QAccessible) UpdateAccessibility(event unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent", ffiqt.FFI_TYPE_POINTER, event)
	gopp.ErrPrint(err, rv)
}
func QAccessible_UpdateAccessibility(event unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.UpdateAccessibility(event)
}

// /usr/include/qt/QtGui/qaccessible.h:426
// index:0
// Public static
// bool isActive()
func (this *QAccessible) IsActive() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible8isActiveEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAccessible_IsActive() {
	var nilthis *QAccessible
	nilthis.IsActive()
}

// /usr/include/qt/QtGui/qaccessible.h:427
// index:0
// Public static
// void setActive(_Bool)
func (this *QAccessible) SetActive(active bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible9setActiveEb", ffiqt.FFI_TYPE_POINTER, active)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetActive(active bool) {
	var nilthis *QAccessible
	nilthis.SetActive(active)
}

// /usr/include/qt/QtGui/qaccessible.h:428
// index:0
// Public static
// void setRootObject(class QObject *)
func (this *QAccessible) SetRootObject(object unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible13setRootObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, object)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetRootObject(object unsafe.Pointer) {
	var nilthis *QAccessible
	nilthis.SetRootObject(object)
}

// /usr/include/qt/QtGui/qaccessible.h:430
// index:0
// Public static
// void cleanup()
func (this *QAccessible) Cleanup() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible7cleanupEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QAccessible_Cleanup() {
	var nilthis *QAccessible
	nilthis.Cleanup()
}

// /usr/include/qt/QtGui/qaccessible.h:432
// index:0
// Public static
// QPair<int, int> qAccessibleTextBoundaryHelper(const class QTextCursor &, enum QAccessible::TextBoundaryType)
func (this *QAccessible) QAccessibleTextBoundaryHelper(cursor *QTextCursor, boundaryType int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible29qAccessibleTextBoundaryHelperERK11QTextCursorNS_16TextBoundaryTypeE", ffiqt.FFI_TYPE_POINTER, cursor, boundaryType)
	gopp.ErrPrint(err, rv)
	return rv
}
func QAccessible_QAccessibleTextBoundaryHelper(cursor *QTextCursor, boundaryType int) {
	var nilthis *QAccessible
	nilthis.QAccessibleTextBoundaryHelper(cursor, boundaryType)
}

//  body block end
