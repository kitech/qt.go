// +build !minimal

package qtwidgets

// /usr/include/qt/QtWidgets/qscrollerproperties.h
// #include <qscrollerproperties.h>
// #include <QtWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 15
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QScrollerProperties struct {
	*qtrt.CObject
}
type QScrollerProperties_ITF interface {
	QScrollerProperties_PTR() *QScrollerProperties
}

func (ptr *QScrollerProperties) QScrollerProperties_PTR() *QScrollerProperties { return ptr }

func (this *QScrollerProperties) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QScrollerProperties) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQScrollerPropertiesFromPointer(cthis unsafe.Pointer) *QScrollerProperties {
	return &QScrollerProperties{&qtrt.CObject{cthis}}
}
func (*QScrollerProperties) NewFromPointer(cthis unsafe.Pointer) *QScrollerProperties {
	return NewQScrollerPropertiesFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QScrollerProperties()

/*
Constructs new scroller properties.
*/
func (*QScrollerProperties) NewForInherit() *QScrollerProperties {
	return NewQScrollerProperties()
}
func NewQScrollerProperties() *QScrollerProperties {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollerPropertiesC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQScrollerPropertiesFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQScrollerProperties)
	return gothis
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:62
// index:0
// Public Visibility=Default Availability=Available
// [16] QScrollerProperties & operator=(const QScrollerProperties &)

/*

 */
func (this *QScrollerProperties) Operator_equal(sp QScrollerProperties_ITF) *QScrollerProperties {
	var convArg0 unsafe.Pointer
	if sp != nil && sp.QScrollerProperties_PTR() != nil {
		convArg0 = sp.QScrollerProperties_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollerPropertiesaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQScrollerPropertiesFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQScrollerProperties)
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScrollerProperties()

/*

 */
func DeleteQScrollerProperties(this *QScrollerProperties) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollerPropertiesD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:65
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QScrollerProperties &) const

/*

 */
func (this *QScrollerProperties) Operator_equal_equal(sp QScrollerProperties_ITF) bool {
	var convArg0 unsafe.Pointer
	if sp != nil && sp.QScrollerProperties_PTR() != nil {
		convArg0 = sp.QScrollerProperties_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollerPropertieseqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QScrollerProperties &) const

/*

 */
func (this *QScrollerProperties) Operator_not_equal(sp QScrollerProperties_ITF) bool {
	var convArg0 unsafe.Pointer
	if sp != nil && sp.QScrollerProperties_PTR() != nil {
		convArg0 = sp.QScrollerProperties_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollerPropertiesneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setDefaultScrollerProperties(const QScrollerProperties &)

/*
Sets the scroller properties for all new QScrollerProperties objects to sp.

Use this function to override the platform default properties returned by the default constructor. If you only want to change the scroller properties of a single scroller, use QScroller::setScrollerProperties()

Note: Calling this function will not change the content of already existing QScrollerProperties objects.

See also unsetDefaultScrollerProperties().
*/
func (this *QScrollerProperties) SetDefaultScrollerProperties(sp QScrollerProperties_ITF) {
	var convArg0 unsafe.Pointer
	if sp != nil && sp.QScrollerProperties_PTR() != nil {
		convArg0 = sp.QScrollerProperties_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollerProperties28setDefaultScrollerPropertiesERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QScrollerProperties_SetDefaultScrollerProperties(sp QScrollerProperties_ITF) {
	var nilthis *QScrollerProperties
	nilthis.SetDefaultScrollerProperties(sp)
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void unsetDefaultScrollerProperties()

/*
Sets the scroller properties returned by the default constructor back to the platform default properties.

See also setDefaultScrollerProperties().
*/
func (this *QScrollerProperties) UnsetDefaultScrollerProperties() {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollerProperties30unsetDefaultScrollerPropertiesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}
func QScrollerProperties_UnsetDefaultScrollerProperties() {
	var nilthis *QScrollerProperties
	nilthis.UnsetDefaultScrollerProperties()
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:117
// index:0
// Public Visibility=Default Availability=Available
// [16] QVariant scrollMetric(QScrollerProperties::ScrollMetric) const

/*
Query the metric value of the scroller properties.

See also setScrollMetric() and ScrollMetric.
*/
func (this *QScrollerProperties) ScrollMetric(metric int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QScrollerProperties12scrollMetricENS_12ScrollMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qscrollerproperties.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScrollMetric(QScrollerProperties::ScrollMetric, const QVariant &)

/*
Set a specific value of the metric ScrollerMetric to value.

See also scrollMetric() and ScrollMetric.
*/
func (this *QScrollerProperties) SetScrollMetric(metric int, value qtcore.QVariant_ITF) {
	var convArg1 unsafe.Pointer
	if value != nil && value.QVariant_PTR() != nil {
		convArg1 = value.QVariant_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QScrollerProperties15setScrollMetricENS_12ScrollMetricERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric, convArg1)
	qtrt.ErrPrint(err, rv)
}

/*
This enum describes the various modes of overshooting.


*/
type QScrollerProperties__OvershootPolicy = int

// Overshooting is possible when the content is scrollable. This is the default.
const QScrollerProperties__OvershootWhenScrollable QScrollerProperties__OvershootPolicy = 0

// Overshooting is never enabled, even when the content is scrollable.
const QScrollerProperties__OvershootAlwaysOff QScrollerProperties__OvershootPolicy = 1

// Overshooting is always enabled, even when the content is not scrollable.
const QScrollerProperties__OvershootAlwaysOn QScrollerProperties__OvershootPolicy = 2

func (this *QScrollerProperties) OvershootPolicyItemName(val int) string {
	switch val {
	case QScrollerProperties__OvershootWhenScrollable: // 0
		return "OvershootWhenScrollable"
	case QScrollerProperties__OvershootAlwaysOff: // 1
		return "OvershootAlwaysOff"
	case QScrollerProperties__OvershootAlwaysOn: // 2
		return "OvershootAlwaysOn"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QScrollerProperties_OvershootPolicyItemName(val int) string {
	var nilthis *QScrollerProperties
	return nilthis.OvershootPolicyItemName(val)
}

/*
This enum describes the available frame rates used while dragging or scrolling.


*/
type QScrollerProperties__FrameRates = int

//
const QScrollerProperties__Standard QScrollerProperties__FrameRates = 0

//
const QScrollerProperties__Fps60 QScrollerProperties__FrameRates = 1

//
const QScrollerProperties__Fps30 QScrollerProperties__FrameRates = 2

//
const QScrollerProperties__Fps20 QScrollerProperties__FrameRates = 3

func (this *QScrollerProperties) FrameRatesItemName(val int) string {
	switch val {
	case QScrollerProperties__Standard: // 0
		return "Standard"
	case QScrollerProperties__Fps60: // 1
		return "Fps60"
	case QScrollerProperties__Fps30: // 2
		return "Fps30"
	case QScrollerProperties__Fps20: // 3
		return "Fps20"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QScrollerProperties_FrameRatesItemName(val int) string {
	var nilthis *QScrollerProperties
	return nilthis.FrameRatesItemName(val)
}

/*
This enum contains the different scroll metric types. When not indicated otherwise the setScrollMetric function expects a QVariant of type qreal.

See the QScroller documentation for further details of the concepts behind the different values.


*/
type QScrollerProperties__ScrollMetric = int

// This is the time a mouse press event is delayed when starting a flick gesture in [s]. If the gesture is triggered within that time, no mouse press or release is sent to the scrolled object. If it triggers after that delay the delayed mouse press plus a faked release event at global position QPoint(-QWIDGETSIZE_MAX, -QWIDGETSIZE_MAX) is sent. If the gesture is canceled, then both the delayed mouse press plus the real release event are delivered.
const QScrollerProperties__MousePressEventDelay QScrollerProperties__ScrollMetric = 0

// This is the minimum distance the touch or mouse point needs to be moved before the flick gesture is triggered in m.
const QScrollerProperties__DragStartDistance QScrollerProperties__ScrollMetric = 1

//
const QScrollerProperties__DragVelocitySmoothingFactor QScrollerProperties__ScrollMetric = 2

//
const QScrollerProperties__AxisLockThreshold QScrollerProperties__ScrollMetric = 3

//
const QScrollerProperties__ScrollingCurve QScrollerProperties__ScrollMetric = 4

//
const QScrollerProperties__DecelerationFactor QScrollerProperties__ScrollMetric = 5

// The minimum velocity that is needed after ending the touch or releasing the mouse to start scrolling in m/s.
const QScrollerProperties__MinimumVelocity QScrollerProperties__ScrollMetric = 6

// This is the maximum velocity that can be reached in m/s.
const QScrollerProperties__MaximumVelocity QScrollerProperties__ScrollMetric = 7

// This is the maximum allowed scroll speed for a click-through in m/s. This means that a click on a currently (slowly) scrolling object will not only stop the scrolling but the click event will also be delivered to the UI control. This is useful when using exponential-type scrolling curves.
const QScrollerProperties__MaximumClickThroughVelocity QScrollerProperties__ScrollMetric = 8

// This is the maximum time in seconds that a flick gesture can take to be recognized as an accelerating flick. If set to zero no such gesture is detected. An "accelerating flick" is a flick gesture executed on an already scrolling object. In such cases the scrolling speed is multiplied by AcceleratingFlickSpeedupFactor in order to accelerate it.
const QScrollerProperties__AcceleratingFlickMaximumTime QScrollerProperties__ScrollMetric = 9

//
const QScrollerProperties__AcceleratingFlickSpeedupFactor QScrollerProperties__ScrollMetric = 10

//
const QScrollerProperties__SnapPositionRatio QScrollerProperties__ScrollMetric = 11

//
const QScrollerProperties__SnapTime QScrollerProperties__ScrollMetric = 12

//
const QScrollerProperties__OvershootDragResistanceFactor QScrollerProperties__ScrollMetric = 13

//
const QScrollerProperties__OvershootDragDistanceFactor QScrollerProperties__ScrollMetric = 14

//
const QScrollerProperties__OvershootScrollDistanceFactor QScrollerProperties__ScrollMetric = 15

//
const QScrollerProperties__OvershootScrollTime QScrollerProperties__ScrollMetric = 16

//
const QScrollerProperties__HorizontalOvershootPolicy QScrollerProperties__ScrollMetric = 17

//
const QScrollerProperties__VerticalOvershootPolicy QScrollerProperties__ScrollMetric = 18

//
const QScrollerProperties__FrameRate QScrollerProperties__ScrollMetric = 19

//
const QScrollerProperties__ScrollMetricCount QScrollerProperties__ScrollMetric = 20

func (this *QScrollerProperties) ScrollMetricItemName(val int) string {
	switch val {
	case QScrollerProperties__MousePressEventDelay: // 0
		return "MousePressEventDelay"
	case QScrollerProperties__DragStartDistance: // 1
		return "DragStartDistance"
	case QScrollerProperties__DragVelocitySmoothingFactor: // 2
		return "DragVelocitySmoothingFactor"
	case QScrollerProperties__AxisLockThreshold: // 3
		return "AxisLockThreshold"
	case QScrollerProperties__ScrollingCurve: // 4
		return "ScrollingCurve"
	case QScrollerProperties__DecelerationFactor: // 5
		return "DecelerationFactor"
	case QScrollerProperties__MinimumVelocity: // 6
		return "MinimumVelocity"
	case QScrollerProperties__MaximumVelocity: // 7
		return "MaximumVelocity"
	case QScrollerProperties__MaximumClickThroughVelocity: // 8
		return "MaximumClickThroughVelocity"
	case QScrollerProperties__AcceleratingFlickMaximumTime: // 9
		return "AcceleratingFlickMaximumTime"
	case QScrollerProperties__AcceleratingFlickSpeedupFactor: // 10
		return "AcceleratingFlickSpeedupFactor"
	case QScrollerProperties__SnapPositionRatio: // 11
		return "SnapPositionRatio"
	case QScrollerProperties__SnapTime: // 12
		return "SnapTime"
	case QScrollerProperties__OvershootDragResistanceFactor: // 13
		return "OvershootDragResistanceFactor"
	case QScrollerProperties__OvershootDragDistanceFactor: // 14
		return "OvershootDragDistanceFactor"
	case QScrollerProperties__OvershootScrollDistanceFactor: // 15
		return "OvershootScrollDistanceFactor"
	case QScrollerProperties__OvershootScrollTime: // 16
		return "OvershootScrollTime"
	case QScrollerProperties__HorizontalOvershootPolicy: // 17
		return "HorizontalOvershootPolicy"
	case QScrollerProperties__VerticalOvershootPolicy: // 18
		return "VerticalOvershootPolicy"
	case QScrollerProperties__FrameRate: // 19
		return "FrameRate"
	case QScrollerProperties__ScrollMetricCount: // 20
		return "ScrollMetricCount"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QScrollerProperties_ScrollMetricItemName(val int) string {
	var nilthis *QScrollerProperties
	return nilthis.ScrollMetricItemName(val)
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
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
