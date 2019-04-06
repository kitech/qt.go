package qtmultimedia

// /usr/include/qt/QtMultimedia/qcameraexposure.h
// #include <qcameraexposure.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QCameraExposure struct {
	*qtcore.QObject
}
type QCameraExposure_ITF interface {
	qtcore.QObject_ITF
	QCameraExposure_PTR() *QCameraExposure
}

func (ptr *QCameraExposure) QCameraExposure_PTR() *QCameraExposure { return ptr }

func (this *QCameraExposure) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QCameraExposure) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQCameraExposureFromPointer(cthis unsafe.Pointer) *QCameraExposure {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QCameraExposure{bcthis0}
}
func (*QCameraExposure) NewFromPointer(cthis unsafe.Pointer) *QCameraExposure {
	return NewQCameraExposureFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QCameraExposure) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:113
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAvailable() const

/*
Returns true if exposure settings are supported by this camera.
*/
func (this *QCameraExposure) IsAvailable() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure11isAvailableEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:115
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraExposure::FlashModes flashMode() const

/*

 */
func (this *QCameraExposure) FlashMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure9flashModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:116
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlashModeSupported(QCameraExposure::FlashModes) const

/*
Returns true if the flash mode is supported.
*/
func (this *QCameraExposure) IsFlashModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure20isFlashModeSupportedE6QFlagsINS_9FlashModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:117
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFlashReady() const

/*
Returns true if flash is charged.

Note: Getter function for property flashReady.
*/
func (this *QCameraExposure) IsFlashReady() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure12isFlashReadyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraExposure::ExposureMode exposureMode() const

/*

 */
func (this *QCameraExposure) ExposureMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure12exposureModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:120
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExposureModeSupported(QCameraExposure::ExposureMode) const

/*
Returns true if the exposure mode is supported.
*/
func (this *QCameraExposure) IsExposureModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure23isExposureModeSupportedENS_12ExposureModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal exposureCompensation() const

/*

 */
func (this *QCameraExposure) ExposureCompensation() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure20exposureCompensationEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:124
// index:0
// Public Visibility=Default Availability=Available
// [4] QCameraExposure::MeteringMode meteringMode() const

/*

 */
func (this *QCameraExposure) MeteringMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure12meteringModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:125
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMeteringModeSupported(QCameraExposure::MeteringMode) const

/*
Returns true if the metering mode is supported.
*/
func (this *QCameraExposure) IsMeteringModeSupported(mode int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure23isMeteringModeSupportedENS_12MeteringModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:127
// index:0
// Public Visibility=Default Availability=Available
// [16] QPointF spotMeteringPoint() const

/*
When supported, the spot metering point is the (normalized) position of the point of the image where exposure metering will be performed. This is typically used to indicate an "interesting" area of the image that should be exposed properly.

The coordinates are relative frame coordinates: QPointF(0,0) points to the left top frame point, QPointF(0.5,0.5) points to the frame center, which is typically the default spot metering point.

The spot metering point is only used with spot metering mode.

See also setSpotMeteringPoint().
*/
func (this *QCameraExposure) SpotMeteringPoint() *qtcore.QPointF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure17spotMeteringPointEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPointF)
	return rv2
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSpotMeteringPoint(const QPointF &)

/*
Allows setting the spot metering point to point.

See also spotMeteringPoint().
*/
func (this *QCameraExposure) SetSpotMeteringPoint(point qtcore.QPointF_ITF) {
	var convArg0 unsafe.Pointer
	if point != nil && point.QPointF_PTR() != nil {
		convArg0 = point.QPointF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure20setSpotMeteringPointERK7QPointF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:130
// index:0
// Public Visibility=Default Availability=Available
// [4] int isoSensitivity() const

/*

 */
func (this *QCameraExposure) IsoSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure14isoSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:131
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal aperture() const

/*

 */
func (this *QCameraExposure) Aperture() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure8apertureEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal shutterSpeed() const

/*
Returns the current shutter speed in seconds.

Note: Getter function for property shutterSpeed.
*/
func (this *QCameraExposure) ShutterSpeed() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure12shutterSpeedEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:134
// index:0
// Public Visibility=Default Availability=Available
// [4] int requestedIsoSensitivity() const

/*
Returns the requested ISO sensitivity or -1 if automatic ISO is turned on.
*/
func (this *QCameraExposure) RequestedIsoSensitivity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure23requestedIsoSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:135
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal requestedAperture() const

/*
Returns the requested manual aperture or -1.0 if automatic aperture is turned on.
*/
func (this *QCameraExposure) RequestedAperture() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure17requestedApertureEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:136
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal requestedShutterSpeed() const

/*
Returns the requested manual shutter speed in seconds or -1.0 if automatic shutter speed is turned on.
*/
func (this *QCameraExposure) RequestedShutterSpeed() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QCameraExposure21requestedShutterSpeedEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:143
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlashMode(QCameraExposure::FlashModes)

/*

 */
func (this *QCameraExposure) SetFlashMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure12setFlashModeE6QFlagsINS_9FlashModeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExposureMode(QCameraExposure::ExposureMode)

/*

 */
func (this *QCameraExposure) SetExposureMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure15setExposureModeENS_12ExposureModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:145
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMeteringMode(QCameraExposure::MeteringMode)

/*

 */
func (this *QCameraExposure) SetMeteringMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure15setMeteringModeENS_12MeteringModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:147
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExposureCompensation(qreal)

/*

 */
func (this *QCameraExposure) SetExposureCompensation(ev float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure23setExposureCompensationEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), ev)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setManualIsoSensitivity(int)

/*
Sets the manual sensitivity to iso
*/
func (this *QCameraExposure) SetManualIsoSensitivity(iso int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure23setManualIsoSensitivityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), iso)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoIsoSensitivity()

/*
Turn on auto sensitivity
*/
func (this *QCameraExposure) SetAutoIsoSensitivity() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure21setAutoIsoSensitivityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setManualAperture(qreal)

/*
Sets the manual camera aperture value.
*/
func (this *QCameraExposure) SetManualAperture(aperture float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure17setManualApertureEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), aperture)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoAperture()

/*
Turn on auto aperture
*/
func (this *QCameraExposure) SetAutoAperture() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure15setAutoApertureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setManualShutterSpeed(qreal)

/*
Set the manual shutter speed to seconds
*/
func (this *QCameraExposure) SetManualShutterSpeed(seconds float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure21setManualShutterSpeedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), seconds)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoShutterSpeed()

/*
Turn on auto shutter speed
*/
func (this *QCameraExposure) SetAutoShutterSpeed() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure19setAutoShutterSpeedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void flashReady(bool)

/*
Signal the flash ready status has changed.

Note: Notifier signal for property flashReady.
*/
func (this *QCameraExposure) FlashReady(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure10flashReadyEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void apertureChanged(qreal)

/*
Signal emitted when aperature changes to value.

Note: Notifier signal for property aperture.
*/
func (this *QCameraExposure) ApertureChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure15apertureChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void apertureRangeChanged()

/*
Signal emitted when aperature range has changed.
*/
func (this *QCameraExposure) ApertureRangeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure20apertureRangeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shutterSpeedChanged(qreal)

/*
Signals that a camera's shutter speed has changed.

Note: Notifier signal for property shutterSpeed.
*/
func (this *QCameraExposure) ShutterSpeedChanged(speed float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure19shutterSpeedChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), speed)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:164
// index:0
// Public Visibility=Default Availability=Available
// [-2] void shutterSpeedRangeChanged()

/*
Signal emitted when the shutter speed range has changed.
*/
func (this *QCameraExposure) ShutterSpeedRangeChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure24shutterSpeedRangeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:165
// index:0
// Public Visibility=Default Availability=Available
// [-2] void isoSensitivityChanged(int)

/*
Signal emitted when sensitivity changes to value.

Note: Notifier signal for property isoSensitivity.
*/
func (this *QCameraExposure) IsoSensitivityChanged(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure21isoSensitivityChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qcameraexposure.h:166
// index:0
// Public Visibility=Default Availability=Available
// [-2] void exposureCompensationChanged(qreal)

/*
Signal emitted when the exposure compensation changes to value.

Note: Notifier signal for property exposureCompensation.
*/
func (this *QCameraExposure) ExposureCompensationChanged(arg0 float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposure27exposureCompensationChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQCameraExposure(this *QCameraExposure) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QCameraExposureD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QCameraExposure__FlashMode = int

//
const QCameraExposure__FlashAuto QCameraExposure__FlashMode = 1

//
const QCameraExposure__FlashOff QCameraExposure__FlashMode = 2

//
const QCameraExposure__FlashOn QCameraExposure__FlashMode = 4

//
const QCameraExposure__FlashRedEyeReduction QCameraExposure__FlashMode = 8

//
const QCameraExposure__FlashFill QCameraExposure__FlashMode = 16

//
const QCameraExposure__FlashTorch QCameraExposure__FlashMode = 32

//
const QCameraExposure__FlashVideoLight QCameraExposure__FlashMode = 64

//
const QCameraExposure__FlashSlowSyncFrontCurtain QCameraExposure__FlashMode = 128

//
const QCameraExposure__FlashSlowSyncRearCurtain QCameraExposure__FlashMode = 256

//
const QCameraExposure__FlashManual QCameraExposure__FlashMode = 512

func (this *QCameraExposure) FlashModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraExposure_FlashModeItemName(val int) string {
	var nilthis *QCameraExposure
	return nilthis.FlashModeItemName(val)
}

/*

 */
type QCameraExposure__ExposureMode = int

// Automatic mode.
const QCameraExposure__ExposureAuto QCameraExposure__ExposureMode = 0

// Manual mode.
const QCameraExposure__ExposureManual QCameraExposure__ExposureMode = 1

// Portrait exposure mode.
const QCameraExposure__ExposurePortrait QCameraExposure__ExposureMode = 2

// Night mode.
const QCameraExposure__ExposureNight QCameraExposure__ExposureMode = 3

// Backlight exposure mode.
const QCameraExposure__ExposureBacklight QCameraExposure__ExposureMode = 4

// Spotlight exposure mode.
const QCameraExposure__ExposureSpotlight QCameraExposure__ExposureMode = 5

// Spots exposure mode.
const QCameraExposure__ExposureSports QCameraExposure__ExposureMode = 6

// Snow exposure mode.
const QCameraExposure__ExposureSnow QCameraExposure__ExposureMode = 7

// Beach exposure mode.
const QCameraExposure__ExposureBeach QCameraExposure__ExposureMode = 8

// Use larger aperture with small depth of field.
const QCameraExposure__ExposureLargeAperture QCameraExposure__ExposureMode = 9

//
const QCameraExposure__ExposureSmallAperture QCameraExposure__ExposureMode = 10

//
const QCameraExposure__ExposureAction QCameraExposure__ExposureMode = 11

//
const QCameraExposure__ExposureLandscape QCameraExposure__ExposureMode = 12

//
const QCameraExposure__ExposureNightPortrait QCameraExposure__ExposureMode = 13

//
const QCameraExposure__ExposureTheatre QCameraExposure__ExposureMode = 14

//
const QCameraExposure__ExposureSunset QCameraExposure__ExposureMode = 15

//
const QCameraExposure__ExposureSteadyPhoto QCameraExposure__ExposureMode = 16

//
const QCameraExposure__ExposureFireworks QCameraExposure__ExposureMode = 17

//
const QCameraExposure__ExposureParty QCameraExposure__ExposureMode = 18

//
const QCameraExposure__ExposureCandlelight QCameraExposure__ExposureMode = 19

//
const QCameraExposure__ExposureBarcode QCameraExposure__ExposureMode = 20

//
const QCameraExposure__ExposureModeVendor QCameraExposure__ExposureMode = 1000

func (this *QCameraExposure) ExposureModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraExposure_ExposureModeItemName(val int) string {
	var nilthis *QCameraExposure
	return nilthis.ExposureModeItemName(val)
}

/*

 */
type QCameraExposure__MeteringMode = int

// Matrix metering mode.
const QCameraExposure__MeteringMatrix QCameraExposure__MeteringMode = 1

// Center weighted average metering mode.
const QCameraExposure__MeteringAverage QCameraExposure__MeteringMode = 2

// Spot metering mode.
const QCameraExposure__MeteringSpot QCameraExposure__MeteringMode = 3

func (this *QCameraExposure) MeteringModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCameraExposure_MeteringModeItemName(val int) string {
	var nilthis *QCameraExposure
	return nilthis.MeteringModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11799() {
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
