package qtpositioning

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h
// #include <qgeosatelliteinfosource.h>
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
type QGeoSatelliteInfoSource struct {
	*qtcore.QObject
}
type QGeoSatelliteInfoSource_ITF interface {
	qtcore.QObject_ITF
	QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource
}

func (ptr *QGeoSatelliteInfoSource) QGeoSatelliteInfoSource_PTR() *QGeoSatelliteInfoSource { return ptr }

func (this *QGeoSatelliteInfoSource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGeoSatelliteInfoSource) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGeoSatelliteInfoSourceFromPointer(cthis unsafe.Pointer) *QGeoSatelliteInfoSource {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGeoSatelliteInfoSource{bcthis0}
}
func (*QGeoSatelliteInfoSource) NewFromPointer(cthis unsafe.Pointer) *QGeoSatelliteInfoSource {
	return NewQGeoSatelliteInfoSourceFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:52
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGeoSatelliteInfoSource) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGeoSatelliteInfoSource10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoSatelliteInfoSource(QObject *)

/*
Creates a satellite source with the specified parent.
*/
func (*QGeoSatelliteInfoSource) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource {
	return NewQGeoSatelliteInfoSource(parent)
}
func NewQGeoSatelliteInfoSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSourceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGeoSatelliteInfoSource")
	return gothis
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGeoSatelliteInfoSource()

/*

 */
func DeleteQGeoSatelliteInfoSource(this *QGeoSatelliteInfoSource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QGeoSatelliteInfoSource * createDefaultSource(QObject *)

/*
Creates and returns a source with the specified parent that reads from the system's default source of satellite update information, or the highest priority available plugin.

Returns 0 if the system has no default satellite source, no valid plugins could be found or the user does not have the permission to access the satellite data.
*/
func (this *QGeoSatelliteInfoSource) CreateDefaultSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource /*777 QGeoSatelliteInfoSource **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource19createDefaultSourceEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGeoSatelliteInfoSource_CreateDefaultSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource /*777 QGeoSatelliteInfoSource **/ {
	var nilthis *QGeoSatelliteInfoSource
	rv := nilthis.CreateDefaultSource(parent)
	return rv
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [8] QGeoSatelliteInfoSource * createSource(const QString &, QObject *)

/*
Creates and returns a source with the given parent, by loading the plugin named sourceName.

Returns 0 if the plugin cannot be found.
*/
func (this *QGeoSatelliteInfoSource) CreateSource(sourceName string, parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource /*777 QGeoSatelliteInfoSource **/ {
	var tmpArg0 = qtcore.NewQString_5(sourceName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource12createSourceERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoSatelliteInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGeoSatelliteInfoSource_CreateSource(sourceName string, parent qtcore.QObject_ITF /*777 QObject **/) *QGeoSatelliteInfoSource /*777 QGeoSatelliteInfoSource **/ {
	var nilthis *QGeoSatelliteInfoSource
	rv := nilthis.CreateSource(sourceName, parent)
	return rv
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:70
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList availableSources()

/*
Returns a list of available source plugins, including the default system backend if one is available.
*/
func (this *QGeoSatelliteInfoSource) AvailableSources() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource16availableSourcesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QGeoSatelliteInfoSource_AvailableSources() *qtcore.QStringList /*123*/ {
	var nilthis *QGeoSatelliteInfoSource
	rv := nilthis.AvailableSources()
	return rv
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:72
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceName() const

/*
Returns the unique name of the satellite source implementation in use.

This is the same name that can be passed to createSource() in order to create a new instance of a particular satellite source implementation.
*/
func (this *QGeoSatelliteInfoSource) SourceName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGeoSatelliteInfoSource10sourceNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:74
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setUpdateInterval(int)

/*

 */
func (this *QGeoSatelliteInfoSource) SetUpdateInterval(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource17setUpdateIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:75
// index:0
// Public Visibility=Default Availability=Available
// [4] int updateInterval() const

/*

 */
func (this *QGeoSatelliteInfoSource) UpdateInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGeoSatelliteInfoSource14updateIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int minimumUpdateInterval() const

/*

 */
func (this *QGeoSatelliteInfoSource) MinimumUpdateInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGeoSatelliteInfoSource21minimumUpdateIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:77
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QGeoSatelliteInfoSource::Error error() const

/*
Returns the last error that occurred.

This signal is not emitted when a requestTimeout() has occurred.
*/
func (this *QGeoSatelliteInfoSource) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK23QGeoSatelliteInfoSource5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QGeoSatelliteInfoSource::Error)

/*
Returns the last error that occurred.

This signal is not emitted when a requestTimeout() has occurred.
*/
func (this *QGeoSatelliteInfoSource) Error_1(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:80
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void startUpdates()

/*
Starts emitting updates at regular intervals. The updates will be provided whenever new satellite information becomes available.

If satellite information cannot be retrieved or some other form of timeout has occurred the satellitesInViewUpdated() and satellitesInUseUpdated() signals may be emitted with empty parameter lists.

See also satellitesInViewUpdated() and satellitesInUseUpdated().
*/
func (this *QGeoSatelliteInfoSource) StartUpdates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource12startUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stopUpdates()

/*
Stops emitting updates at regular intervals.
*/
func (this *QGeoSatelliteInfoSource) StopUpdates() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource11stopUpdatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void requestUpdate(int)

/*
Attempts to get the current satellite information and emit satellitesInViewUpdated() and satellitesInUseUpdated() with this information. If the current satellite information cannot be found within the given timeout (in milliseconds) or if timeout is less than the value returned by minimumUpdateInterval(), requestTimeout() is emitted.

If the timeout is zero, the timeout defaults to a reasonable timeout period as appropriate for the source.

This does nothing if another update request is in progress. However it can be called even if startUpdates() has already been called and regular updates are in progress.
*/
func (this *QGeoSatelliteInfoSource) RequestUpdate(timeout int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource13requestUpdateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:83
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void requestUpdate(int)

/*
Attempts to get the current satellite information and emit satellitesInViewUpdated() and satellitesInUseUpdated() with this information. If the current satellite information cannot be found within the given timeout (in milliseconds) or if timeout is less than the value returned by minimumUpdateInterval(), requestTimeout() is emitted.

If the timeout is zero, the timeout defaults to a reasonable timeout period as appropriate for the source.

This does nothing if another update request is in progress. However it can be called even if startUpdates() has already been called and regular updates are in progress.
*/
func (this *QGeoSatelliteInfoSource) RequestUpdate__() {
	// arg: 0, int=Int, =Invalid, , Invalid
	timeout := int(0)
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource13requestUpdateEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), timeout)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeosatelliteinfosource.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestTimeout()

/*
Emitted if requestUpdate() was called and the current satellite information could not be retrieved within the specified timeout.

While the triggering of this signal may be considered an error condition, it does not imply the emission of the error() signal. Only the emission of requestTimeout() is required to indicate a timeout.
*/
func (this *QGeoSatelliteInfoSource) RequestTimeout() {
	rv, err := qtrt.InvokeQtFunc6("_ZN23QGeoSatelliteInfoSource14requestTimeoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

/*
The Error enumeration represents the errors which can occur.


*/
type QGeoSatelliteInfoSource__Error = int

// The connection setup to the satellite backend failed because the application lacked the required privileges.
const QGeoSatelliteInfoSource__AccessError QGeoSatelliteInfoSource__Error = 0

// The satellite backend closed the connection, which happens for example in case the user is switching location services to off. This object becomes invalid and should be deleted. A new satellite source can be created by calling createDefaultSource() later on.
const QGeoSatelliteInfoSource__ClosedError QGeoSatelliteInfoSource__Error = 1

// No error has occurred.
const QGeoSatelliteInfoSource__NoError QGeoSatelliteInfoSource__Error = 2

//
const QGeoSatelliteInfoSource__UnknownSourceError QGeoSatelliteInfoSource__Error = -1

func (this *QGeoSatelliteInfoSource) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QGeoSatelliteInfoSource_ErrorItemName(val int) string {
	var nilthis *QGeoSatelliteInfoSource
	return nilthis.ErrorItemName(val)
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
