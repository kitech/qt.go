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

// /usr/include/qt/QtGui/qaccessible.h:411
// index:0
// static
// void installActivationObserver(class QAccessible::ActivationObserver *)
func (this *QAccessible) InstallActivationObserver(arg0 unsafe.Pointer) {
	// 0: (QAccessible::ActivationObserver * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible25installActivationObserverEPNS_18ActivationObserverE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_InstallActivationObserver(arg0 unsafe.Pointer) {
	// 0: (QAccessible::ActivationObserver * arg0), (arg0)
	var nilthis *QAccessible
	nilthis.InstallActivationObserver(arg0)
}

// /usr/include/qt/QtGui/qaccessible.h:412
// index:0
// static
// void removeActivationObserver(class QAccessible::ActivationObserver *)
func (this *QAccessible) RemoveActivationObserver(arg0 unsafe.Pointer) {
	// 0: (QAccessible::ActivationObserver * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible24removeActivationObserverEPNS_18ActivationObserverE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_RemoveActivationObserver(arg0 unsafe.Pointer) {
	// 0: (QAccessible::ActivationObserver * arg0), (arg0)
	var nilthis *QAccessible
	nilthis.RemoveActivationObserver(arg0)
}

// /usr/include/qt/QtGui/qaccessible.h:414
// index:0
// static
// QAccessibleInterface * queryAccessibleInterface(class QObject *)
func (this *QAccessible) QueryAccessibleInterface(arg0 unsafe.Pointer) {
	// 0: (QObject * arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible24queryAccessibleInterfaceEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_QueryAccessibleInterface(arg0 unsafe.Pointer) {
	// 0: (QObject * arg0), (arg0)
	var nilthis *QAccessible
	nilthis.QueryAccessibleInterface(arg0)
}

// /usr/include/qt/QtGui/qaccessible.h:415
// index:0
// static
// QAccessible::Id uniqueId(class QAccessibleInterface *)
func (this *QAccessible) UniqueId(iface unsafe.Pointer) {
	// 0: (iface QAccessibleInterface *), (iface)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible8uniqueIdEP20QAccessibleInterface", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_UniqueId(iface unsafe.Pointer) {
	// 0: (iface QAccessibleInterface *), (iface)
	var nilthis *QAccessible
	nilthis.UniqueId(iface)
}

// /usr/include/qt/QtGui/qaccessible.h:416
// index:0
// static
// QAccessibleInterface * accessibleInterface(QAccessible::Id)
func (this *QAccessible) AccessibleInterface(uniqueId uint) {
	// 0: (uniqueId QAccessible::Id), (uniqueId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible19accessibleInterfaceEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_AccessibleInterface(uniqueId uint) {
	// 0: (uniqueId QAccessible::Id), (uniqueId)
	var nilthis *QAccessible
	nilthis.AccessibleInterface(uniqueId)
}

// /usr/include/qt/QtGui/qaccessible.h:417
// index:0
// static
// QAccessible::Id registerAccessibleInterface(class QAccessibleInterface *)
func (this *QAccessible) RegisterAccessibleInterface(iface unsafe.Pointer) {
	// 0: (iface QAccessibleInterface *), (iface)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_RegisterAccessibleInterface(iface unsafe.Pointer) {
	// 0: (iface QAccessibleInterface *), (iface)
	var nilthis *QAccessible
	nilthis.RegisterAccessibleInterface(iface)
}

// /usr/include/qt/QtGui/qaccessible.h:418
// index:0
// static
// void deleteAccessibleInterface(QAccessible::Id)
func (this *QAccessible) DeleteAccessibleInterface(uniqueId uint) {
	// 0: (uniqueId QAccessible::Id), (uniqueId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible25deleteAccessibleInterfaceEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_DeleteAccessibleInterface(uniqueId uint) {
	// 0: (uniqueId QAccessible::Id), (uniqueId)
	var nilthis *QAccessible
	nilthis.DeleteAccessibleInterface(uniqueId)
}

// /usr/include/qt/QtGui/qaccessible.h:424
// index:0
// static
// void updateAccessibility(class QAccessibleEvent *)
func (this *QAccessible) UpdateAccessibility(event unsafe.Pointer) {
	// 0: (event QAccessibleEvent *), (event)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_UpdateAccessibility(event unsafe.Pointer) {
	// 0: (event QAccessibleEvent *), (event)
	var nilthis *QAccessible
	nilthis.UpdateAccessibility(event)
}

// /usr/include/qt/QtGui/qaccessible.h:426
// index:0
// static
// bool isActive()
func (this *QAccessible) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible8isActiveEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_IsActive() {
	// 0: (), ()
	var nilthis *QAccessible
	nilthis.IsActive()
}

// /usr/include/qt/QtGui/qaccessible.h:427
// index:0
// static
// void setActive(_Bool)
func (this *QAccessible) SetActive(active bool) {
	// 0: (active bool), (active)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible9setActiveEb", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetActive(active bool) {
	// 0: (active bool), (active)
	var nilthis *QAccessible
	nilthis.SetActive(active)
}

// /usr/include/qt/QtGui/qaccessible.h:428
// index:0
// static
// void setRootObject(class QObject *)
func (this *QAccessible) SetRootObject(object unsafe.Pointer) {
	// 0: (object QObject *), (object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible13setRootObjectEP7QObject", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetRootObject(object unsafe.Pointer) {
	// 0: (object QObject *), (object)
	var nilthis *QAccessible
	nilthis.SetRootObject(object)
}

// /usr/include/qt/QtGui/qaccessible.h:430
// index:0
// static
// void cleanup()
func (this *QAccessible) Cleanup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible7cleanupEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_Cleanup() {
	// 0: (), ()
	var nilthis *QAccessible
	nilthis.Cleanup()
}

// /usr/include/qt/QtGui/qaccessible.h:432
// index:0
// static
// QPair<int, int> qAccessibleTextBoundaryHelper(const class QTextCursor &, enum QAccessible::TextBoundaryType)
func (this *QAccessible) QAccessibleTextBoundaryHelper(cursor unsafe.Pointer, boundaryType int) {
	// 0: (cursor const QTextCursor &, boundaryType QAccessible::TextBoundaryType), (cursor, boundaryType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible29qAccessibleTextBoundaryHelperERK11QTextCursorNS_16TextBoundaryTypeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QAccessible_QAccessibleTextBoundaryHelper(cursor unsafe.Pointer, boundaryType int) {
	// 0: (cursor const QTextCursor &, boundaryType QAccessible::TextBoundaryType), (cursor, boundaryType)
	var nilthis *QAccessible
	nilthis.QAccessibleTextBoundaryHelper(cursor, boundaryType)
}

//  body block end
