package qtpositioning

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h
// #include <qnmeapositioninfosource.h>
// #include <QtPositioning>

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

// bool parsePosInfoFromNmeaData(const char *, int, QGeoPositionInfo *, bool *)
func (this *QNmeaPositionInfoSource) InheritParsePosInfoFromNmeaData(f func(data string, size int, posInfo *QGeoPositionInfo /*777 QGeoPositionInfo **/, hasFix *bool) bool) {
	qtrt.SetAllInheritCallback(this, "parsePosInfoFromNmeaData", f)
}

/*

 */
type QNmeaPositionInfoSource struct {
	*QGeoPositionInfoSource
}
type QNmeaPositionInfoSource_ITF interface {
	QGeoPositionInfoSource_ITF
	QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource
}

func (ptr *QNmeaPositionInfoSource) QNmeaPositionInfoSource_PTR() *QNmeaPositionInfoSource { return ptr }

func (this *QNmeaPositionInfoSource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGeoPositionInfoSource.GetCthis()
	}
}
func (this *QNmeaPositionInfoSource) SetCthis(cthis unsafe.Pointer) {
	this.QGeoPositionInfoSource = NewQGeoPositionInfoSourceFromPointer(cthis)
}
func NewQNmeaPositionInfoSourceFromPointer(cthis unsafe.Pointer) *QNmeaPositionInfoSource {
	bcthis0 := NewQGeoPositionInfoSourceFromPointer(cthis)
	return &QNmeaPositionInfoSource{bcthis0}
}
func (*QNmeaPositionInfoSource) NewFromPointer(cthis unsafe.Pointer) *QNmeaPositionInfoSource {
	return NewQNmeaPositionInfoSourceFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:51
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QNmeaPositionInfoSource) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNmeaPositionInfoSource(QNmeaPositionInfoSource::UpdateMode, QObject *)

/*
Constructs a QNmeaPositionInfoSource instance with the given parent and updateMode.
*/
func (*QNmeaPositionInfoSource) NewForInherit(updateMode int, parent qtcore.QObject_ITF /*777 QObject **/) *QNmeaPositionInfoSource {
	return NewQNmeaPositionInfoSource(updateMode, parent)
}
func NewQNmeaPositionInfoSource(updateMode int, parent qtcore.QObject_ITF /*777 QObject **/) *QNmeaPositionInfoSource {
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSourceC2ENS_10UpdateModeEP7QObject", qtrt.FFI_TYPE_POINTER, updateMode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNmeaPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNmeaPositionInfoSource")
	return gothis
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QNmeaPositionInfoSource(QNmeaPositionInfoSource::UpdateMode, QObject *)

/*
Constructs a QNmeaPositionInfoSource instance with the given parent and updateMode.
*/
func (*QNmeaPositionInfoSource) NewForInheritp(updateMode int) *QNmeaPositionInfoSource {
	return NewQNmeaPositionInfoSourcep(updateMode)
}
func NewQNmeaPositionInfoSourcep(updateMode int) *QNmeaPositionInfoSource {
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSourceC2ENS_10UpdateModeEP7QObject", qtrt.FFI_TYPE_POINTER, updateMode, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQNmeaPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QNmeaPositionInfoSource")
	return gothis
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QNmeaPositionInfoSource()

/*

 */
func DeleteQNmeaPositionInfoSource(this *QNmeaPositionInfoSource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserEquivalentRangeError(double)

/*
Sets the User Equivalent Range Error (UERE) to uere. The UERE is used in calculating an estimate of the accuracy of the position information reported by the position info source. The UERE should be set to a value appropriate for the GPS device which generated the NMEA stream.

The true UERE value is calculated from multiple error sources including errors introduced by the satellites and signal propogation delays through the atmosphere as well as errors introduced by the receiving GPS equipment. For details on GPS accuracy see http://edu-observatory.org/gps/gps_accuracy.html.

A typical value for UERE is approximately 5.1.

This function was introduced in  Qt 5.3.

See also userEquivalentRangeError().
*/
func (this *QNmeaPositionInfoSource) SetUserEquivalentRangeError(uere float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource27setUserEquivalentRangeErrorEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), uere)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:62
// index:0
// Public Visibility=Default Availability=Available
// [8] double userEquivalentRangeError() const

/*
Returns the current User Equivalent Range Error (UERE). The UERE is used in calculating an estimate of the accuracy of the position information reported by the position info source. The default value is NaN which means no accuracy information will be provided.

This function was introduced in  Qt 5.3.

See also setUserEquivalentRangeError().
*/
func (this *QNmeaPositionInfoSource) UserEquivalentRangeError() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource24userEquivalentRangeErrorEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:64
// index:0
// Public Visibility=Default Availability=Available
// [4] QNmeaPositionInfoSource::UpdateMode updateMode() const

/*
Returns the update mode.
*/
func (this *QNmeaPositionInfoSource) UpdateMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource10updateModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:66
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)

/*
Sets the NMEA data source to device. If the device is not open, it will be opened in QIODevice::ReadOnly mode.

The source device can only be set once and must be set before calling startUpdates() or requestUpdate().

Note: The device must emit QIODevice::readyRead() for the source to be notified when data is available for reading. QNmeaPositionInfoSource does not assume the ownership of the device, and hence does not deallocate it upon destruction.

See also device().
*/
func (this *QNmeaPositionInfoSource) SetDevice(source qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if source != nil && source.QIODevice_PTR() != nil {
		convArg0 = source.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device() const

/*
Returns the NMEA data source.

See also setDevice().
*/
func (this *QNmeaPositionInfoSource) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:69
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setUpdateInterval(int)

/*
Reimplemented from QGeoPositionInfoSource::setUpdateInterval().
*/
func (this *QNmeaPositionInfoSource) SetUpdateInterval(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource17setUpdateIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGeoPositionInfo lastKnownPosition(bool) const

/*
Reimplemented from QGeoPositionInfoSource::lastKnownPosition().
*/
func (this *QNmeaPositionInfoSource) LastKnownPosition(fromSatellitePositioningMethodsOnly bool) *QGeoPositionInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource17lastKnownPositionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fromSatellitePositioningMethodsOnly)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPositionInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGeoPositionInfo lastKnownPosition(bool) const

/*
Reimplemented from QGeoPositionInfoSource::lastKnownPosition().
*/
func (this *QNmeaPositionInfoSource) LastKnownPositionp() *QGeoPositionInfo /*123*/ {
	// arg: 0, bool=Bool, =Invalid, , Invalid
	fromSatellitePositioningMethodsOnly := false
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource17lastKnownPositionEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), fromSatellitePositioningMethodsOnly)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQGeoPositionInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQGeoPositionInfo)
	return rv2
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QGeoPositionInfoSource::PositioningMethods supportedPositioningMethods() const

/*
Reimplemented from QGeoPositionInfoSource::supportedPositioningMethods().
*/
func (this *QNmeaPositionInfoSource) SupportedPositioningMethods() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource27supportedPositioningMethodsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int minimumUpdateInterval() const

/*
Reimplemented from QGeoPositionInfoSource::minimumUpdateInterval().
*/
func (this *QNmeaPositionInfoSource) MinimumUpdateInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource21minimumUpdateIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QGeoPositionInfoSource::Error error() const

/*
Reimplemented from QGeoPositionInfoSource::error().
*/
func (this *QNmeaPositionInfoSource) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QNmeaPositionInfoSource5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:78
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void startUpdates()

/*
Reimplemented from QGeoPositionInfoSource::startUpdates().
*/
func (this *QNmeaPositionInfoSource) StartUpdates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource12startUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void stopUpdates()

/*
Reimplemented from QGeoPositionInfoSource::stopUpdates().
*/
func (this *QNmeaPositionInfoSource) StopUpdates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource11stopUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void requestUpdate(int)

/*
Reimplemented from QGeoPositionInfoSource::requestUpdate().
*/
func (this *QNmeaPositionInfoSource) RequestUpdate(timeout int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource13requestUpdateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void requestUpdate(int)

/*
Reimplemented from QGeoPositionInfoSource::requestUpdate().
*/
func (this *QNmeaPositionInfoSource) RequestUpdatep() {
	// arg: 0, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource13requestUpdateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qnmeapositioninfosource.h:83
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool parsePosInfoFromNmeaData(const char *, int, QGeoPositionInfo *, bool *)

/*
Parses an NMEA sentence string into a QGeoPositionInfo.

The default implementation will parse standard NMEA sentences. This method should be reimplemented in a subclass whenever the need to deal with non-standard NMEA sentences arises.

The parser reads size bytes from data and uses that information to setup posInfo and hasFix. If hasFix is set to false then posInfo may contain only the time or the date and the time.

Returns true if the sentence was succsesfully parsed, otherwise returns false and should not modifiy posInfo or hasFix.
*/
func (this *QNmeaPositionInfoSource) ParsePosInfoFromNmeaData(data string, size int, posInfo QGeoPositionInfo_ITF /*777 QGeoPositionInfo **/, hasFix *bool) bool {
	var convArg0 = qtrt.CString(data)
	defer qtrt.FreeMem(convArg0)
	var convArg2 unsafe.Pointer
	if posInfo != nil && posInfo.QGeoPositionInfo_PTR() != nil {
		convArg2 = posInfo.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QNmeaPositionInfoSource24parsePosInfoFromNmeaDataEPKciP16QGeoPositionInfoPb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, size, convArg2, hasFix)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
Defines the available update modes.


*/
type QNmeaPositionInfoSource__UpdateMode = int

// Positional data is read and distributed from the data source as it becomes available. Use this mode if you are using a live source of positional data (for example, a GPS hardware device).
const QNmeaPositionInfoSource__RealTimeMode QNmeaPositionInfoSource__UpdateMode = 1

// The data and time information in the NMEA source data is used to provide positional updates at the rate at which the data was originally recorded. Use this mode if the data source contains previously recorded NMEA data and you want to replay the data for simulation purposes.
const QNmeaPositionInfoSource__SimulationMode QNmeaPositionInfoSource__UpdateMode = 2

func (this *QNmeaPositionInfoSource) UpdateModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QNmeaPositionInfoSource_UpdateModeItemName(val int) string {
	var nilthis *QNmeaPositionInfoSource
	return nilthis.UpdateModeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11661() {
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
