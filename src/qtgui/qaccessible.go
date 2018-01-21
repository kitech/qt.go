package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQAccessibleFromPointer(cthis unsafe.Pointer) *QAccessible {
	return &QAccessible{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qaccessible.h:414
// index:0
// Public static
// QAccessibleInterface * queryAccessibleInterface(class QObject *)
func (this *QAccessible) QueryAccessibleInterface(arg0 *qtcore.QObject /*444 QObject **/) *QAccessibleInterface /*444 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible24queryAccessibleInterfaceEP7QObject", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QAccessible_QueryAccessibleInterface(arg0 *qtcore.QObject /*444 QObject **/) *QAccessibleInterface /*444 QAccessibleInterface **/ {
	var nilthis *QAccessible
	rv := nilthis.QueryAccessibleInterface(arg0)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:415
// index:0
// Public static
// QAccessible::Id uniqueId(class QAccessibleInterface *)
func (this *QAccessible) UniqueId(iface *QAccessibleInterface /*444 QAccessibleInterface **/) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible8uniqueIdEP20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, iface)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QAccessible_UniqueId(iface *QAccessibleInterface /*444 QAccessibleInterface **/) uint {
	var nilthis *QAccessible
	rv := nilthis.UniqueId(iface)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:416
// index:0
// Public static
// QAccessibleInterface * accessibleInterface(QAccessible::Id)
func (this *QAccessible) AccessibleInterface(uniqueId uint) *QAccessibleInterface /*444 QAccessibleInterface **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible19accessibleInterfaceEj", ffiqt.FFI_TYPE_POINTER, uniqueId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QAccessible_AccessibleInterface(uniqueId uint) *QAccessibleInterface /*444 QAccessibleInterface **/ {
	var nilthis *QAccessible
	rv := nilthis.AccessibleInterface(uniqueId)
	return rv
}

// /usr/include/qt/QtGui/qaccessible.h:417
// index:0
// Public static
// QAccessible::Id registerAccessibleInterface(class QAccessibleInterface *)
func (this *QAccessible) RegisterAccessibleInterface(iface *QAccessibleInterface /*444 QAccessibleInterface **/) uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible27registerAccessibleInterfaceEP20QAccessibleInterface", ffiqt.FFI_TYPE_POINTER, iface)
	gopp.ErrPrint(err, rv)
	// return rv
	return uint(rv) // 222
}
func QAccessible_RegisterAccessibleInterface(iface *QAccessibleInterface /*444 QAccessibleInterface **/) uint {
	var nilthis *QAccessible
	rv := nilthis.RegisterAccessibleInterface(iface)
	return rv
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
func (this *QAccessible) UpdateAccessibility(event *QAccessibleEvent /*444 QAccessibleEvent **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible19updateAccessibilityEP16QAccessibleEvent", ffiqt.FFI_TYPE_POINTER, event)
	gopp.ErrPrint(err, rv)
}
func QAccessible_UpdateAccessibility(event *QAccessibleEvent /*444 QAccessibleEvent **/) {
	var nilthis *QAccessible
	nilthis.UpdateAccessibility(event)
}

// /usr/include/qt/QtGui/qaccessible.h:426
// index:0
// Public static
// bool isActive()
func (this *QAccessible) IsActive() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible8isActiveEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QAccessible_IsActive() bool {
	var nilthis *QAccessible
	rv := nilthis.IsActive()
	return rv
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
func (this *QAccessible) SetRootObject(object *qtcore.QObject /*444 QObject **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QAccessible13setRootObjectEP7QObject", ffiqt.FFI_TYPE_POINTER, object)
	gopp.ErrPrint(err, rv)
}
func QAccessible_SetRootObject(object *qtcore.QObject /*444 QObject **/) {
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

//  body block end
