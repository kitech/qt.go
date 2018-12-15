package qtpositioning

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h
// #include <qgeopositioninfosource.h>
// #include <QtPositioning>

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

//  ext block end

//  body block begin

/*

 */
type QGeoPositionInfoSource struct {
	*qtcore.QObject
}
type QGeoPositionInfoSource_ITF interface {
	qtcore.QObject_ITF
	QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource
}

func (ptr *QGeoPositionInfoSource) QGeoPositionInfoSource_PTR() *QGeoPositionInfoSource { return ptr }

func (this *QGeoPositionInfoSource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGeoPositionInfoSource) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGeoPositionInfoSourceFromPointer(cthis unsafe.Pointer) *QGeoPositionInfoSource {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGeoPositionInfoSource{bcthis0}
}
func (*QGeoPositionInfoSource) NewFromPointer(cthis unsafe.Pointer) *QGeoPositionInfoSource {
	return NewQGeoPositionInfoSourceFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:51
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGeoPositionInfoSource) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoPositionInfoSource(QObject *)

/*
Creates a position source with the specified parent.
*/
func (*QGeoPositionInfoSource) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource {
	return NewQGeoPositionInfoSource(parent)
}
func NewQGeoPositionInfoSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSourceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGeoPositionInfoSource")
	return gothis
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGeoPositionInfoSource()

/*

 */
func DeleteQGeoPositionInfoSource(this *QGeoPositionInfoSource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setUpdateInterval(int)

/*

 */
func (this *QGeoPositionInfoSource) SetUpdateInterval(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource17setUpdateIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:77
// index:0
// Public Visibility=Default Availability=Available
// [4] int updateInterval() const

/*

 */
func (this *QGeoPositionInfoSource) UpdateInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource14updateIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods)

/*
Sets the preferred positioning methods for this source to methods.

If methods includes a method that is not supported by the source, the unsupported method will be ignored.

If methods does not include any methods supported by the source, the preferred methods will be set to the set of methods which the source supports.

Note: When reimplementing this method, subclasses must call the base method implementation to ensure preferredPositioningMethods() returns the correct value.

See also preferredPositioningMethods() and supportedPositioningMethods().
*/
func (this *QGeoPositionInfoSource) SetPreferredPositioningMethods(methods int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource30setPreferredPositioningMethodsE6QFlagsINS_17PositioningMethodEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), methods)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:80
// index:0
// Public Visibility=Default Availability=Available
// [4] QGeoPositionInfoSource::PositioningMethods preferredPositioningMethods() const

/*
Returns the positioning methods set by setPreferredPositioningMethods().

See also setPreferredPositioningMethods().
*/
func (this *QGeoPositionInfoSource) PreferredPositioningMethods() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource27preferredPositioningMethodsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGeoPositionInfo lastKnownPosition(bool) const

/*
Returns an update containing the last known position, or a null update if none is available.

If fromSatellitePositioningMethodsOnly is true, this returns the last known position received from a satellite positioning method; if none is available, a null update is returned.
*/
func (this *QGeoPositionInfoSource) LastKnownPosition(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource17lastKnownPositionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fromSatellitePositioningMethodsOnly)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPositionInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:82
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QGeoPositionInfo lastKnownPosition(bool) const

/*
Returns an update containing the last known position, or a null update if none is available.

If fromSatellitePositioningMethodsOnly is true, this returns the last known position received from a satellite positioning method; if none is available, a null update is returned.
*/
func (this *QGeoPositionInfoSource) LastKnownPositionp() *QGeoPositionInfo /*123*/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	fromSatellitePositioningMethodsOnly := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource17lastKnownPositionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fromSatellitePositioningMethodsOnly)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPositionInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:84
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QGeoPositionInfoSource::PositioningMethods supportedPositioningMethods() const

/*
Returns the positioning methods available to this source.

See also setPreferredPositioningMethods().
*/
func (this *QGeoPositionInfoSource) SupportedPositioningMethods() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource27supportedPositioningMethodsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int minimumUpdateInterval() const

/*

 */
func (this *QGeoPositionInfoSource) MinimumUpdateInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource21minimumUpdateIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceName() const

/*

 */
func (this *QGeoPositionInfoSource) SourceName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource10sourceNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:89
// index:0
// Public static Visibility=Default Availability=Available
// [8] QGeoPositionInfoSource * createDefaultSource(QObject *)

/*
Creates and returns a position source with the given parent that reads from the system's default sources of location data, or the plugin with the highest available priority.

Returns 0 if the system has no default position source, no valid plugins could be found or the user does not have the permission to access the current position.
*/
func (this *QGeoPositionInfoSource) CreateDefaultSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource /*777 QGeoPositionInfoSource **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource19createDefaultSourceEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGeoPositionInfoSource_CreateDefaultSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource /*777 QGeoPositionInfoSource **/ {
	var nilthis *QGeoPositionInfoSource
	rv := nilthis.CreateDefaultSource(parent)
	return rv
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:90
// index:0
// Public static Visibility=Default Availability=Available
// [8] QGeoPositionInfoSource * createSource(const QString &, QObject *)

/*
Creates and returns a position source with the given parent, by loading the plugin named sourceName.

Returns 0 if the plugin cannot be found.
*/
func (this *QGeoPositionInfoSource) CreateSource(sourceName string, parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource /*777 QGeoPositionInfoSource **/ {
	var tmpArg0 = qtcore.NewQString5(sourceName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource12createSourceERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGeoPositionInfoSource_CreateSource(sourceName string, parent qtcore.QObject_ITF /*777 QObject **/) *QGeoPositionInfoSource /*777 QGeoPositionInfoSource **/ {
	var nilthis *QGeoPositionInfoSource
	rv := nilthis.CreateSource(sourceName, parent)
	return rv
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:91
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList availableSources()

/*
Returns a list of available source plugins. This includes any default backend plugin for the current platform.
*/
func (this *QGeoPositionInfoSource) AvailableSources() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource16availableSourcesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QGeoPositionInfoSource_AvailableSources() *qtcore.QStringList /*123*/ {
	var nilthis *QGeoPositionInfoSource
	rv := nilthis.AvailableSources()
	return rv
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:92
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QGeoPositionInfoSource::Error error() const

/*
Returns the type of error that last occurred.
*/
func (this *QGeoPositionInfoSource) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK22QGeoPositionInfoSource5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:103
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QGeoPositionInfoSource::Error)

/*
Returns the type of error that last occurred.
*/
func (this *QGeoPositionInfoSource) Error1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:95
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void startUpdates()

/*
Starts emitting updates at regular intervals as specified by setUpdateInterval().

If setUpdateInterval() has not been called, the source will emit updates as soon as they become available.

An updateTimeout() signal will be emitted if this QGeoPositionInfoSource subclass determines that it will not be able to provide regular updates. This could happen if a satellite fix is lost or if a hardware error is detected. Position updates will recommence if the data becomes available later on. The updateTimeout() signal will not be emitted again until after the periodic updates resume.

On iOS, starting from version 8, Core Location framework requires additional entries in the application's Info.plist with keys NSLocationAlwaysUsageDescription or NSLocationWhenInUseUsageDescription and a string to be displayed in the authorization prompt. The key NSLocationWhenInUseUsageDescription is used when requesting permission to use location services while the app is in the foreground. The key NSLocationAlwaysUsageDescription is used when requesting permission to use location services whenever the app is running (both the foreground and the background). If both entries are defined, NSLocationWhenInUseUsageDescription has a priority in the foreground mode.
*/
func (this *QGeoPositionInfoSource) StartUpdates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource12startUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:96
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stopUpdates()

/*
Stops emitting updates at regular intervals.
*/
func (this *QGeoPositionInfoSource) StopUpdates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource11stopUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:98
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void requestUpdate(int)

/*
Attempts to get the current position and emit positionUpdated() with this information. If the current position cannot be found within the given timeout (in milliseconds) or if timeout is less than the value returned by minimumUpdateInterval(), updateTimeout() is emitted.

If the timeout is zero, the timeout defaults to a reasonable timeout period as appropriate for the source.

This does nothing if another update request is in progress. However it can be called even if startUpdates() has already been called and regular updates are in progress.

If the source uses multiple positioning methods, it tries to get the current position from the most accurate positioning method within the given timeout.
*/
func (this *QGeoPositionInfoSource) RequestUpdate(timeout int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource13requestUpdateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:98
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void requestUpdate(int)

/*
Attempts to get the current position and emit positionUpdated() with this information. If the current position cannot be found within the given timeout (in milliseconds) or if timeout is less than the value returned by minimumUpdateInterval(), updateTimeout() is emitted.

If the timeout is zero, the timeout defaults to a reasonable timeout period as appropriate for the source.

This does nothing if another update request is in progress. However it can be called even if startUpdates() has already been called and regular updates are in progress.

If the source uses multiple positioning methods, it tries to get the current position from the most accurate positioning method within the given timeout.
*/
func (this *QGeoPositionInfoSource) RequestUpdatep() {
	// arg: 0, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource13requestUpdateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void positionUpdated(const QGeoPositionInfo &)

/*
If startUpdates() or requestUpdate() is called, this signal is emitted when an update becomes available.

The update value holds the value of the new update.
*/
func (this *QGeoPositionInfoSource) PositionUpdated(update QGeoPositionInfo_ITF) {
	var convArg0 unsafe.Pointer
	if update != nil && update.QGeoPositionInfo_PTR() != nil {
		convArg0 = update.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource15positionUpdatedERK16QGeoPositionInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateTimeout()

/*
If requestUpdate() was called, this signal will be emitted if the current position could not be retrieved within the specified timeout.

If startUpdates() has been called, this signal will be emitted if this QGeoPositionInfoSource subclass determines that it will not be able to provide further regular updates. This signal will not be emitted again until after the regular updates resume.

While the triggering of this signal may be considered an error condition, it does not imply the emission of the error() signal. Only the emission of updateTimeout() is required to indicate a timeout.
*/
func (this *QGeoPositionInfoSource) UpdateTimeout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource13updateTimeoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeopositioninfosource.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void supportedPositioningMethodsChanged()

/*
This signal is emitted after the supportedPositioningMethods change.

This function was introduced in  Qt 5.12.
*/
func (this *QGeoPositionInfoSource) SupportedPositioningMethodsChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN22QGeoPositionInfoSource34supportedPositioningMethodsChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
The Error enumeration represents the errors which can occur.


*/
type QGeoPositionInfoSource__Error = int

// The connection setup to the remote positioning backend failed because the application lacked the required privileges.
const QGeoPositionInfoSource__AccessError QGeoPositionInfoSource__Error = 0

// The remote positioning backend closed the connection, which happens for example in case the user is switching location services to off. As soon as the location service is re-enabled regular updates will resume.
const QGeoPositionInfoSource__ClosedError QGeoPositionInfoSource__Error = 1

// An unidentified error occurred.
const QGeoPositionInfoSource__UnknownSourceError QGeoPositionInfoSource__Error = 2

// No error has occurred.
const QGeoPositionInfoSource__NoError QGeoPositionInfoSource__Error = 3

func (this *QGeoPositionInfoSource) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGeoPositionInfoSource_ErrorItemName(val int) string {
	var nilthis *QGeoPositionInfoSource
	return nilthis.ErrorItemName(val)
}

/*


 */
type QGeoPositionInfoSource__PositioningMethod = int

//
const QGeoPositionInfoSource__NoPositioningMethods QGeoPositionInfoSource__PositioningMethod = 0

//
const QGeoPositionInfoSource__SatellitePositioningMethods QGeoPositionInfoSource__PositioningMethod = 255

//
const QGeoPositionInfoSource__NonSatellitePositioningMethods QGeoPositionInfoSource__PositioningMethod = -256

//
const QGeoPositionInfoSource__AllPositioningMethods QGeoPositionInfoSource__PositioningMethod = -1

func (this *QGeoPositionInfoSource) PositioningMethodItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGeoPositionInfoSource_PositioningMethodItemName(val int) string {
	var nilthis *QGeoPositionInfoSource
	return nilthis.PositioningMethodItemName(val)
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
