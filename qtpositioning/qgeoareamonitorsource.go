package qtpositioning

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h
// #include <qgeoareamonitorsource.h>
// #include <QtPositioning>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 21
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
type QGeoAreaMonitorSource struct {
	*qtcore.QObject
}
type QGeoAreaMonitorSource_ITF interface {
	qtcore.QObject_ITF
	QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource
}

func (ptr *QGeoAreaMonitorSource) QGeoAreaMonitorSource_PTR() *QGeoAreaMonitorSource { return ptr }

func (this *QGeoAreaMonitorSource) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QGeoAreaMonitorSource) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQGeoAreaMonitorSourceFromPointer(cthis unsafe.Pointer) *QGeoAreaMonitorSource {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QGeoAreaMonitorSource{bcthis0}
}
func (*QGeoAreaMonitorSource) NewFromPointer(cthis unsafe.Pointer) *QGeoAreaMonitorSource {
	return NewQGeoAreaMonitorSourceFromPointer(cthis)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QGeoAreaMonitorSource) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGeoAreaMonitorSource10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QGeoAreaMonitorSource(QObject *)

/*
Creates a monitor with the given parent.
*/
func NewQGeoAreaMonitorSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoAreaMonitorSource {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSourceC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQGeoAreaMonitorSourceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QGeoAreaMonitorSource")
	return gothis
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QGeoAreaMonitorSource()

/*

 */
func DeleteQGeoAreaMonitorSource(this *QGeoAreaMonitorSource) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSourceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:75
// index:0
// Public static Visibility=Default Availability=Available
// [8] QGeoAreaMonitorSource * createDefaultSource(QObject *)

/*
Creates and returns a monitor with the given parent that monitors areas using resources on the underlying system.

Returns 0 if the system has no support for position monitoring.
*/
func (this *QGeoAreaMonitorSource) CreateDefaultSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoAreaMonitorSource /*777 QGeoAreaMonitorSource **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource19createDefaultSourceEP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoAreaMonitorSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGeoAreaMonitorSource_CreateDefaultSource(parent qtcore.QObject_ITF /*777 QObject **/) *QGeoAreaMonitorSource /*777 QGeoAreaMonitorSource **/ {
	var nilthis *QGeoAreaMonitorSource
	rv := nilthis.CreateDefaultSource(parent)
	return rv
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:76
// index:0
// Public static Visibility=Default Availability=Available
// [8] QGeoAreaMonitorSource * createSource(const QString &, QObject *)

/*
Creates and returns a monitor with the given parent, by loading the plugin named sourceName.

Returns 0 if the plugin cannot be found.
*/
func (this *QGeoAreaMonitorSource) CreateSource(sourceName string, parent qtcore.QObject_ITF /*777 QObject **/) *QGeoAreaMonitorSource /*777 QGeoAreaMonitorSource **/ {
	var tmpArg0 = qtcore.NewQString_5(sourceName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource12createSourceERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoAreaMonitorSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QGeoAreaMonitorSource_CreateSource(sourceName string, parent qtcore.QObject_ITF /*777 QObject **/) *QGeoAreaMonitorSource /*777 QGeoAreaMonitorSource **/ {
	var nilthis *QGeoAreaMonitorSource
	rv := nilthis.CreateSource(sourceName, parent)
	return rv
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:77
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList availableSources()

/*
Returns a list of available monitor plugins, including the default system backend if one is available.
*/
func (this *QGeoAreaMonitorSource) AvailableSources() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource16availableSourcesEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QGeoAreaMonitorSource_AvailableSources() *qtcore.QStringList /*123*/ {
	var nilthis *QGeoAreaMonitorSource
	rv := nilthis.AvailableSources()
	return rv
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:79
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPositionInfoSource(QGeoPositionInfoSource *)

/*
Sets the new QGeoPositionInfoSource to be used by this QGeoAreaMonitorSource object. The area monitoring backend becomes the new QObject parent for newSource. The previous QGeoPositionInfoSource object will be deleted. All QGeoAreaMonitorSource instances based on the same sourceName() share the same QGeoPositionInfoSource instance.

This may be useful when it is desirable to manipulate the positioning system used by the area monitoring engine.

Note that ownership must be taken care of by subclasses of QGeoAreaMonitorSource. Due to the singleton pattern behind this class newSource may be moved to a new thread.

See also positionInfoSource().
*/
func (this *QGeoAreaMonitorSource) SetPositionInfoSource(source QGeoPositionInfoSource_ITF /*777 QGeoPositionInfoSource **/) {
	var convArg0 unsafe.Pointer
	if source != nil && source.QGeoPositionInfoSource_PTR() != nil {
		convArg0 = source.QGeoPositionInfoSource_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource21setPositionInfoSourceEP22QGeoPositionInfoSource", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:80
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QGeoPositionInfoSource * positionInfoSource() const

/*
Returns the current QGeoPositionInfoSource used by this QGeoAreaMonitorSource object. The function will return QGeoPositionInfoSource::createDefaultSource() if no other object has been set.

The function returns 0 if not even a default QGeoPositionInfoSource exists.

Any usage of the returned QGeoPositionInfoSource instance should account for the fact that it may reside in a different thread.

See also QGeoPositionInfoSource and setPositionInfoSource().
*/
func (this *QGeoAreaMonitorSource) PositionInfoSource() *QGeoPositionInfoSource /*777 QGeoPositionInfoSource **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGeoAreaMonitorSource18positionInfoSourceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQGeoPositionInfoSourceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QString sourceName() const

/*
Returns the unique name of the area monitor source implementation in use.

This is the same name that can be passed to createSource() in order to create a new instance of a particular area monitor source implementation.
*/
func (this *QGeoAreaMonitorSource) SourceName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGeoAreaMonitorSource10sourceNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:84
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QGeoAreaMonitorSource::Error error() const

/*
Returns the type of error that last occurred.
*/
func (this *QGeoAreaMonitorSource) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGeoAreaMonitorSource5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:98
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QGeoAreaMonitorSource::Error)

/*
Returns the type of error that last occurred.
*/
func (this *QGeoAreaMonitorSource) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] QGeoAreaMonitorSource::AreaMonitorFeatures supportedAreaMonitorFeatures() const

/*
Returns the area monitoring features available to this source.
*/
func (this *QGeoAreaMonitorSource) SupportedAreaMonitorFeatures() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK21QGeoAreaMonitorSource28supportedAreaMonitorFeaturesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool startMonitoring(const QGeoAreaMonitorInfo &)

/*
Returns true if the monitoring of monitor could be successfully started; otherwise returns false. A reason for not being able to start monitoring could be the unavailability of an appropriate default position info source while no alternative QGeoPositionInfoSource has been set via setPositionInfoSource().

If monitor is already active the existing monitor object will be replaced by the new monitor reference. The identification of QGeoAreaMonitorInfo instances happens via QGeoAreaMonitorInfo::identifier(). Therefore this function can also be used to update active monitors.

If monitor has an expiry date that has been passed this function returns false. Calling this function for an already via requestUpdate() registered single shot monitor switches the monitor to a permanent monitoring mode.

Requesting persistent monitoring on a QGeoAreaMonitorSource instance fails if the area monitoring backend doesn't support QGeoAreaMonitorSource::PersistentAreaMonitorFeature.

See also stopMonitoring().
*/
func (this *QGeoAreaMonitorSource) StartMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if monitor != nil && monitor.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = monitor.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource15startMonitoringERK19QGeoAreaMonitorInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool stopMonitoring(const QGeoAreaMonitorInfo &)

/*
Returns true if monitor was successfully removed from the list of activeMonitors(); otherwise returns false. This behavior is independent on whether monitor was registered via startMonitoring() or requestUpdate().
*/
func (this *QGeoAreaMonitorSource) StopMonitoring(monitor QGeoAreaMonitorInfo_ITF) bool {
	var convArg0 unsafe.Pointer
	if monitor != nil && monitor.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = monitor.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource14stopMonitoringERK19QGeoAreaMonitorInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:89
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool requestUpdate(const QGeoAreaMonitorInfo &, const char *)

/*
Enables single shot area monitoring. Area monitoring for monitor will be performed until this QGeoAreaMonitorSource instance emits signal for the first time. Once the signal was emitted, monitor is automatically removed from the list of activeMonitors(). If monitor is invalid or has an expiry date that has been passed this function returns false.


  QGeoAreaMonitor singleShotMonitor;
  QGeoAreaMonitorSource * source = QGeoAreaMonitorSource::createDefaultSource(this);
  //...
  bool ret = source->requestUpdate(singleShotMonitor,
                        SIGNAL(areaExited(QGeoAreaMonitor,QGeoPositionInfo)));



The above singleShotMonitor object will cease to send updates once the areaExited() signal was emitted for the first time. Until this point in time any other signal may be emitted zero or more times depending on the area context.

It is not possible to simultanously request updates for more than one signal of the same monitor object. The last call to this function determines the signal upon which the updates cease to continue. At this stage only the areaEntered() and areaExited() signals can be used to terminate the monitoring process.

Requesting persistent monitoring on a QGeoAreaMonitorSource instance fails if the area monitoring backend doesn't support QGeoAreaMonitorSource::PersistentAreaMonitorFeature.

If monitor was already registered via startMonitoring() it is converted to a single shot behavior.

See also startMonitoring() and stopMonitoring().
*/
func (this *QGeoAreaMonitorSource) RequestUpdate(monitor QGeoAreaMonitorInfo_ITF, signal string) bool {
	var convArg0 unsafe.Pointer
	if monitor != nil && monitor.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = monitor.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(signal)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource13requestUpdateERK19QGeoAreaMonitorInfoPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:95
// index:0
// Public Visibility=Default Availability=Available
// [-2] void areaEntered(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)

/*
Emitted when the current position has moved from a position outside of the active monitor to a position within the monitored area.

The update holds the new position.
*/
func (this *QGeoAreaMonitorSource) AreaEntered(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {
	var convArg0 unsafe.Pointer
	if monitor != nil && monitor.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = monitor.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if update != nil && update.QGeoPositionInfo_PTR() != nil {
		convArg1 = update.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource11areaEnteredERK19QGeoAreaMonitorInfoRK16QGeoPositionInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:96
// index:0
// Public Visibility=Default Availability=Available
// [-2] void areaExited(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)

/*
Emitted when the current position has moved from a position within the active monitor to a position outside the monitored area.

The update holds the new position.
*/
func (this *QGeoAreaMonitorSource) AreaExited(monitor QGeoAreaMonitorInfo_ITF, update QGeoPositionInfo_ITF) {
	var convArg0 unsafe.Pointer
	if monitor != nil && monitor.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = monitor.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if update != nil && update.QGeoPositionInfo_PTR() != nil {
		convArg1 = update.QGeoPositionInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource10areaExitedERK19QGeoAreaMonitorInfoRK16QGeoPositionInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtPositioning/qgeoareamonitorsource.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void monitorExpired(const QGeoAreaMonitorInfo &)

/*
Emitted when monitor has expired. An expired area monitor is automatically removed from the list of activeMonitors().

See also activeMonitors().
*/
func (this *QGeoAreaMonitorSource) MonitorExpired(monitor QGeoAreaMonitorInfo_ITF) {
	var convArg0 unsafe.Pointer
	if monitor != nil && monitor.QGeoAreaMonitorInfo_PTR() != nil {
		convArg0 = monitor.QGeoAreaMonitorInfo_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN21QGeoAreaMonitorSource14monitorExpiredERK19QGeoAreaMonitorInfo", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

/*
Defines the types of positioning methods.

The Error enumeration represents the errors which can occur.


*/
type QGeoAreaMonitorSource__Error = int

// The connection setup to the remote area monitoring backend failed because the application lacked the required privileges.
const QGeoAreaMonitorSource__AccessError QGeoAreaMonitorSource__Error = 0

// The area monitoring source could not retrieve a location fix or the accuracy of the fix is not high enough to provide an effective area monitoring.
const QGeoAreaMonitorSource__InsufficientPositionInfo QGeoAreaMonitorSource__Error = 1

// An unidentified error occurred.
const QGeoAreaMonitorSource__UnknownSourceError QGeoAreaMonitorSource__Error = 2

// No error has occurred.
const QGeoAreaMonitorSource__NoError QGeoAreaMonitorSource__Error = 3

/*


 */
type QGeoAreaMonitorSource__AreaMonitorFeature = int

//
const QGeoAreaMonitorSource__PersistentAreaMonitorFeature QGeoAreaMonitorSource__AreaMonitorFeature = 1

//
const QGeoAreaMonitorSource__AnyAreaMonitorFeature QGeoAreaMonitorSource__AreaMonitorFeature = -1

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
