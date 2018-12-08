package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAccessibleEvent struct {
	*qtrt.CObject
}
type QAccessibleEvent_ITF interface {
	QAccessibleEvent_PTR() *QAccessibleEvent
}

func (ptr *QAccessibleEvent) QAccessibleEvent_PTR() *QAccessibleEvent { return ptr }

func (this *QAccessibleEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAccessibleEvent) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAccessibleEventFromPointer(cthis unsafe.Pointer) *QAccessibleEvent {
	return &QAccessibleEvent{&qtrt.CObject{cthis}}
}
func (*QAccessibleEvent) NewFromPointer(cthis unsafe.Pointer) *QAccessibleEvent {
	return NewQAccessibleEventFromPointer(cthis)
}

// /usr/include/qt/QtGui/qaccessible.h:668
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleEvent(QObject *, QAccessible::Event)

/*

 */
func (*QAccessibleEvent) NewForInherit(obj qtcore.QObject_ITF /*777 QObject **/, typ int) *QAccessibleEvent {
	return NewQAccessibleEvent(obj, typ)
}
func NewQAccessibleEvent(obj qtcore.QObject_ITF /*777 QObject **/, typ int) *QAccessibleEvent {
	var convArg0 unsafe.Pointer
	if obj != nil && obj.QObject_PTR() != nil {
		convArg0 = obj.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP7QObjectN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, convArg0, typ)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:684
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QAccessibleEvent(QAccessibleInterface *, QAccessible::Event)

/*

 */
func (*QAccessibleEvent) NewForInherit1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, typ int) *QAccessibleEvent {
	return NewQAccessibleEvent1(iface, typ)
}
func NewQAccessibleEvent1(iface QAccessibleInterface_ITF /*777 QAccessibleInterface **/, typ int) *QAccessibleEvent {
	var convArg0 unsafe.Pointer
	if iface != nil && iface.QAccessibleInterface_PTR() != nil {
		convArg0 = iface.QAccessibleInterface_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEventC2EP20QAccessibleInterfaceN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, convArg0, typ)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAccessibleEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQAccessibleEvent)
	return gothis
}

// /usr/include/qt/QtGui/qaccessible.h:699
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAccessibleEvent()

/*

 */
func DeleteQAccessibleEvent(this *QAccessibleEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qaccessible.h:701
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QAccessible::Event type() const

/*

 */
func (this *QAccessibleEvent) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:702
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QObject * object() const

/*

 */
func (this *QAccessibleEvent) Object() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent6objectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qaccessible.h:703
// index:0
// Public Visibility=Default Availability=Available
// [4] QAccessible::Id uniqueId() const

/*
Returns the unique ID for the QAccessibleInterface iface.
*/
func (this *QAccessibleEvent) UniqueId() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent8uniqueIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:705
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setChild(int)

/*

 */
func (this *QAccessibleEvent) SetChild(chld int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QAccessibleEvent8setChildEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), chld)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qaccessible.h:706
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int child() const

/*

 */
func (this *QAccessibleEvent) Child() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent5childEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:708
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleInterface() const

/*
Returns the QAccessibleInterface belonging to the id.

Returns 0 if the id is invalid.
*/
func (this *QAccessibleEvent) AccessibleInterface() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QAccessibleEvent19accessibleInterfaceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
