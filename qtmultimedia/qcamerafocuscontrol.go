package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h
// #include <qcamerafocuscontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
type QCameraFocusControl struct {
	*QMediaControl
}
type QCameraFocusControl_ITF interface {
	QMediaControl_ITF
	QCameraFocusControl_PTR() *QCameraFocusControl
}

func (ptr *QCameraFocusControl) QCameraFocusControl_PTR() *QCameraFocusControl { return ptr }

func (this *QCameraFocusControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QCameraFocusControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQCameraFocusControlFromPointer(cthis unsafe.Pointer) *QCameraFocusControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QCameraFocusControl{bcthis0}
}
func (*QCameraFocusControl) NewFromPointer(cthis unsafe.Pointer) *QCameraFocusControl {
	return NewQCameraFocusControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraFocusControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QCameraFocusControl()

/*

 */
func DeleteQCameraFocusControl(this *QCameraFocusControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCameraFocus::FocusModes focusMode() const

/*
Returns the focus mode being used.

See also setFocusMode().
*/
func (this *QCameraFocusControl) FocusMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl9focusModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFocusMode(QCameraFocus::FocusModes)

/*
Set the focus mode to mode.

See also focusMode().
*/
func (this *QCameraFocusControl) SetFocusMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl12setFocusModeE6QFlagsIN12QCameraFocus9FocusModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isFocusModeSupported(QCameraFocus::FocusModes) const

/*
Returns true if focus mode is supported.
*/
func (this *QCameraFocusControl) IsFocusModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl20isFocusModeSupportedE6QFlagsIN12QCameraFocus9FocusModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QCameraFocus::FocusPointMode focusPointMode() const

/*
Returns the camera focus point selection mode.

See also setFocusPointMode().
*/
func (this *QCameraFocusControl) FocusPointMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl14focusPointModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFocusPointMode(QCameraFocus::FocusPointMode)

/*
Sets the camera focus point selection mode.

See also focusPointMode().
*/
func (this *QCameraFocusControl) SetFocusPointMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl17setFocusPointModeEN12QCameraFocus14FocusPointModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isFocusPointModeSupported(QCameraFocus::FocusPointMode) const

/*
Returns true if the camera focus point mode is supported.
*/
func (this *QCameraFocusControl) IsFocusPointModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl25isFocusPointModeSupportedEN12QCameraFocus14FocusPointModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [16] QPointF customFocusPoint() const

/*
Return the position of custom focus point, in relative frame coordinates: QPointF(0,0) points to the left top frame point, QPointF(0.5,0.5) points to the frame center.

Custom focus point is used only in FocusPointCustom focus mode.

See also setCustomFocusPoint().
*/
func (this *QCameraFocusControl) CustomFocusPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl16customFocusPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:68
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setCustomFocusPoint(const QPointF &)

/*
Sets the custom focus point.

If camera supports fixed set of focus points, it should use the nearest supported focus point, and return the actual focus point with QCameraFocusControl::focusZones().

See also QCameraFocusControl::customFocusPoint() and QCameraFocusControl::focusZones().
*/
func (this *QCameraFocusControl) SetCustomFocusPoint(point qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl19setCustomFocusPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] QCameraFocusZoneList focusZones() const

/*
Returns the list of zones, the camera is using for focusing or focused on.
*/
func (this *QCameraFocusControl) FocusZones() *QCameraFocusZoneList /*667*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QCameraFocusControl10focusZonesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCameraFocusZoneListFromPointer(unsafe.Pointer(uintptr(rv))) //5551
	return rv2
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusModeChanged(QCameraFocus::FocusModes)

/*
Signal is emitted when the focus mode is changed, usually in result of QCameraFocusControl::setFocusMode call or capture mode changes.

See also QCameraFocusControl::focusMode() and QCameraFocusControl::setFocusMode().
*/
func (this *QCameraFocusControl) FocusModeChanged(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl16focusModeChangedE6QFlagsIN12QCameraFocus9FocusModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusPointModeChanged(QCameraFocus::FocusPointMode)

/*
Signal is emitted when the focus point mode is changed, usually in result of QCameraFocusControl::setFocusPointMode call or capture mode changes.

See also QCameraFocusControl::focusPointMode() and QCameraFocusControl::setFocusPointMode().
*/
func (this *QCameraFocusControl) FocusPointModeChanged(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl21focusPointModeChangedEN12QCameraFocus14FocusPointModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void customFocusPointChanged(const QPointF &)

/*
Signal is emitted when the custom focus point is changed.

See also QCameraFocusControl::customFocusPoint() and QCameraFocusControl::setCustomFocusPoint().
*/
func (this *QCameraFocusControl) CustomFocusPointChanged(point qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl23customFocusPointChangedERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusZonesChanged()

/*
Signal is emitted when the set of zones, camera focused on is changed.

Usually the zones list is changed when the camera is focused.

See also QCameraFocusControl::focusZones().
*/
func (this *QCameraFocusControl) FocusZonesChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControl17focusZonesChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:80
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraFocusControl(QObject *)

/*
Constructs a camera control object with parent.
*/
func (*QCameraFocusControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraFocusControl {
	return NewQCameraFocusControl(parent)
}
func NewQCameraFocusControl(parent qtcore.QObject_ITF /*777 QObject **/) *QCameraFocusControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFocusControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraFocusControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qcamerafocuscontrol.h:80
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QCameraFocusControl(QObject *)

/*
Constructs a camera control object with parent.
*/
func (*QCameraFocusControl) NewForInheritp() *QCameraFocusControl {
	return NewQCameraFocusControlp()
}
func NewQCameraFocusControlp() *QCameraFocusControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN19QCameraFocusControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQCameraFocusControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QCameraFocusControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11825() {
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
