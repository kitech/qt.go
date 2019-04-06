package qtmultimedia

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h
// #include <qcustomaudiorolecontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QCustomAudioRoleControl struct {
	*QMediaControl
}
type QCustomAudioRoleControl_ITF interface {
	QMediaControl_ITF
	QCustomAudioRoleControl_PTR() *QCustomAudioRoleControl
}

func (ptr *QCustomAudioRoleControl) QCustomAudioRoleControl_PTR() *QCustomAudioRoleControl { return ptr }

func (this *QCustomAudioRoleControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCustomAudioRoleControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCustomAudioRoleControlFromPointer(cthis unsafe.Pointer) *QCustomAudioRoleControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCustomAudioRoleControl{bcthis0}
}
func (*QCustomAudioRoleControl) NewFromPointer(cthis unsafe.Pointer) *QCustomAudioRoleControl {
	return NewQCustomAudioRoleControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCustomAudioRoleControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QCustomAudioRoleControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:57
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCustomAudioRoleControl()

/*

 */
func DeleteQCustomAudioRoleControl(this *QCustomAudioRoleControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QCustomAudioRoleControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString customAudioRole() const

/*
Returns the audio role of the media played by the media service.

See also setCustomAudioRole().
*/
func (this *QCustomAudioRoleControl) CustomAudioRole() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QCustomAudioRoleControl15customAudioRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCustomAudioRole(const QString &)

/*
Sets the audio role of the media played by the media service.

See also customAudioRole().
*/
func (this *QCustomAudioRoleControl) SetCustomAudioRole(role string) {
	var tmpArg0 = qtcore.NewQString5(role)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QCustomAudioRoleControl18setCustomAudioRoleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedCustomAudioRoles() const

/*
Returns a list of custom audio roles that the media service supports. An empty list may indicate that the supported custom audio roles aren't known. The list may not be complete.
*/
func (this *QCustomAudioRoleControl) SupportedCustomAudioRoles() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QCustomAudioRoleControl25supportedCustomAudioRolesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void customAudioRoleChanged(const QString &)

/*
Signal emitted when the audio role has changed.
*/
func (this *QCustomAudioRoleControl) CustomAudioRoleChanged(role string) {
	var tmpArg0 = qtcore.NewQString5(role)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN23QCustomAudioRoleControl22customAudioRoleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:68
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCustomAudioRoleControl(QObject *)

/*
Construct a QCustomAudioRoleControl with the given parent.
*/
func (*QCustomAudioRoleControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCustomAudioRoleControl {
	return NewQCustomAudioRoleControl(parent)
}
func NewQCustomAudioRoleControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCustomAudioRoleControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QCustomAudioRoleControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCustomAudioRoleControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCustomAudioRoleControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcustomaudiorolecontrol.h:68
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCustomAudioRoleControl(QObject *)

/*
Construct a QCustomAudioRoleControl with the given parent.
*/
func (*QCustomAudioRoleControl) NewForInheritp() *QCustomAudioRoleControl {
	return NewQCustomAudioRoleControlp()
}
func NewQCustomAudioRoleControlp() *QCustomAudioRoleControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QCustomAudioRoleControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCustomAudioRoleControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCustomAudioRoleControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11843() {
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
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
