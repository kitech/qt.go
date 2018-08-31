package qtmultimedia

// /usr/include/qt/QtMultimedia/qradiotuner.h
// #include <qradiotuner.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
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
type QRadioTuner struct {
	*QMediaObject
}
type QRadioTuner_ITF interface {
	QMediaObject_ITF
	QRadioTuner_PTR() *QRadioTuner
}

func (ptr *QRadioTuner) QRadioTuner_PTR() *QRadioTuner { return ptr }

func (this *QRadioTuner) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaObject.GetCthis()
	}
}
func (this *QRadioTuner) SetCthis(cthis unsafe.Pointer) {
	this.QMediaObject = NewQMediaObjectFromPointer(cthis)
}
func NewQRadioTunerFromPointer(cthis unsafe.Pointer) *QRadioTuner {
	bcthis0 := NewQMediaObjectFromPointer(cthis)
	return &QRadioTuner{bcthis0}
}
func (*QRadioTuner) NewFromPointer(cthis unsafe.Pointer) *QRadioTuner {
	return NewQRadioTunerFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRadioTuner) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioTuner(QObject *)

/*
Constructs a radio tuner based on a media service allocated by the default media service provider.

The parent is passed to QMediaObject.
*/
func (*QRadioTuner) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QRadioTuner {
	return NewQRadioTuner(parent)
}
func NewQRadioTuner(parent qtcore.QObject_ITF /*777 QObject **/) *QRadioTuner {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTunerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioTunerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioTuner")
	return gothis
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:81
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioTuner(QObject *)

/*
Constructs a radio tuner based on a media service allocated by the default media service provider.

The parent is passed to QMediaObject.
*/
func (*QRadioTuner) NewForInherit__() *QRadioTuner {
	return NewQRadioTuner__()
}
func NewQRadioTuner__() *QRadioTuner {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTunerC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioTunerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioTuner")
	return gothis
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:82
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRadioTuner()

/*

 */
func DeleteQRadioTuner(this *QRadioTuner) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTunerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:84
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Reimplemented from QMediaObject::availability().

Returns the availability of the radio tuner.
*/
func (this *QRadioTuner) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] QRadioTuner::State state() const

/*

 */
func (this *QRadioTuner) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] QRadioTuner::Band band() const

/*

 */
func (this *QRadioTuner) Band() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner4bandEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isBandSupported(QRadioTuner::Band) const

/*
Identifies if a frequency band is supported by a radio tuner.

Returns true if the band is supported, and false if it is not.
*/
func (this *QRadioTuner) IsBandSupported(b int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner15isBandSupportedENS_4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:92
// index:0
// Public Visibility=Default Availability=Available
// [4] int frequency() const

/*

 */
func (this *QRadioTuner) Frequency() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner9frequencyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] int frequencyStep(QRadioTuner::Band) const

/*
Returns the number of Hertz to increment the frequency by when stepping through frequencies within a given band.
*/
func (this *QRadioTuner) FrequencyStep(band int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner13frequencyStepENS_4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), band)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:96
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isStereo() const

/*

 */
func (this *QRadioTuner) IsStereo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner8isStereoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStereoMode(QRadioTuner::StereoMode)

/*

 */
func (this *QRadioTuner) SetStereoMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner13setStereoModeENS_10StereoModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:98
// index:0
// Public Visibility=Default Availability=Available
// [4] QRadioTuner::StereoMode stereoMode() const

/*

 */
func (this *QRadioTuner) StereoMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner10stereoModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] int signalStrength() const

/*

 */
func (this *QRadioTuner) SignalStrength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner14signalStrengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] int volume() const

/*

 */
func (this *QRadioTuner) Volume() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner6volumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMuted() const

/*

 */
func (this *QRadioTuner) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:105
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isSearching() const

/*

 */
func (this *QRadioTuner) IsSearching() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner11isSearchingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:107
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAntennaConnected() const

/*

 */
func (this *QRadioTuner) IsAntennaConnected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner18isAntennaConnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:109
// index:0
// Public Visibility=Default Availability=Available
// [4] QRadioTuner::Error error() const

/*
Returns the error state of a radio tuner.

See also errorString().
*/
func (this *QRadioTuner) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:141
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QRadioTuner::Error)

/*
Returns the error state of a radio tuner.

See also errorString().
*/
func (this *QRadioTuner) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a description of a radio tuner's error state.

See also error().
*/
func (this *QRadioTuner) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:112
// index:0
// Public Visibility=Default Availability=Available
// [8] QRadioData * radioData() const

/*

 */
func (this *QRadioTuner) RadioData() *QRadioData /*777 QRadioData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QRadioTuner9radioDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQRadioDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchForward()

/*
Starts a forward scan for a signal, starting from the current frequency.

See also searchBackward(), cancelSearch(), and searching.
*/
func (this *QRadioTuner) SearchForward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner13searchForwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchBackward()

/*
Starts a backwards scan for a signal, starting from the current frequency.

See also searchForward(), cancelSearch(), and searching.
*/
func (this *QRadioTuner) SearchBackward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner14searchBackwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchAllStations(QRadioTuner::SearchMode)

/*
Search all stations in current band

Emits QRadioTuner::stationFound(int, QString) for every found station. After searching is completed, QRadioTuner::searchingChanged(bool) is emitted (false). If searchMode is set to SearchGetStationId, searching waits for station id (PI) on each frequency.

See also searchForward(), searchBackward(), and searching.
*/
func (this *QRadioTuner) SearchAllStations(searchMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner17searchAllStationsENS_10SearchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), searchMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchAllStations(QRadioTuner::SearchMode)

/*
Search all stations in current band

Emits QRadioTuner::stationFound(int, QString) for every found station. After searching is completed, QRadioTuner::searchingChanged(bool) is emitted (false). If searchMode is set to SearchGetStationId, searching waits for station id (PI) on each frequency.

See also searchForward(), searchBackward(), and searching.
*/
func (this *QRadioTuner) SearchAllStations__() {
	// arg: 0, QRadioTuner::SearchMode=Elaborated, QRadioTuner::SearchMode=Enum, , Invalid
	searchMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner17searchAllStationsENS_10SearchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), searchMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:118
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cancelSearch()

/*
Stops scanning for a signal.

See also searchForward(), searchBackward(), and searching.
*/
func (this *QRadioTuner) CancelSearch() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner12cancelSearchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:120
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBand(QRadioTuner::Band)

/*
Sets a radio tuner's frequency band.

Changing the band will reset the frequency to the new band's minimum frequency.

Note: Setter function for property band.

See also band().
*/
func (this *QRadioTuner) SetBand(band int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner7setBandENS_4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), band)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFrequency(int)

/*
Sets a radio tuner's frequency.

If the tuner is set to a frequency outside the current band, the band will be changed to one occupied by the new frequency.

Note: Setter function for property frequency.

See also frequency().
*/
func (this *QRadioTuner) SetFrequency(frequency int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner12setFrequencyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVolume(int)

/*

 */
func (this *QRadioTuner) SetVolume(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner9setVolumeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:124
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*

 */
func (this *QRadioTuner) SetMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:126
// index:0
// Public Visibility=Default Availability=Available
// [-2] void start()

/*
Activate the radio device.
*/
func (this *QRadioTuner) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:127
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stop()

/*
Deactivate the radio device.
*/
func (this *QRadioTuner) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QRadioTuner::State)

/*
This signal is emitted when the state changes to state.

Note: Notifier signal for property state.
*/
func (this *QRadioTuner) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner12stateChangedENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:131
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bandChanged(QRadioTuner::Band)

/*
Signals a radio tuner's band has changed.

Note: Notifier signal for property band.
*/
func (this *QRadioTuner) BandChanged(band int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner11bandChangedENS_4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), band)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frequencyChanged(int)

/*
Signals that the frequency a radio tuner is tuned to has changed.

Note: Notifier signal for property frequency.
*/
func (this *QRadioTuner) FrequencyChanged(frequency int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner16frequencyChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stereoStatusChanged(bool)

/*
Signals that the stereo state of a radio tuner has changed.

Note: Notifier signal for property stereo.
*/
func (this *QRadioTuner) StereoStatusChanged(stereo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner19stereoStatusChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stereo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchingChanged(bool)

/*
Signals that the searching state of a radio tuner has changed.

Note: Notifier signal for property searching.
*/
func (this *QRadioTuner) SearchingChanged(searching bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner16searchingChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), searching)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void signalStrengthChanged(int)

/*
Signals that the strength of the signal received by a radio tuner has changed.

Note: Notifier signal for property signalStrength.
*/
func (this *QRadioTuner) SignalStrengthChanged(signalStrength int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner21signalStrengthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), signalStrength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged(int)

/*
Signals that the volume of a radio tuner's audio output has changed.

Note: Notifier signal for property volume.
*/
func (this *QRadioTuner) VolumeChanged(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner13volumeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged(bool)

/*
Signals that the muted state of a radio tuner's audio output has changed.

Note: Notifier signal for property muted.
*/
func (this *QRadioTuner) MutedChanged(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner12mutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:138
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stationFound(int, QString)

/*
Signals that a station was found in frequency with stationId Program Identification code.
*/
func (this *QRadioTuner) StationFound(frequency int, stationId string) {
	var tmpArg1 = qtcore.NewQString_5(stationId)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner12stationFoundEi7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotuner.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void antennaConnectedChanged(bool)

/*

 */
func (this *QRadioTuner) AntennaConnectedChanged(connectionStatus bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QRadioTuner23antennaConnectedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), connectionStatus)
	qtrt.ErrPrint(err, rv)
}

/*
Enumerates radio tuner states.


*/
type QRadioTuner__State = int

// The tuner is started and active.
const QRadioTuner__ActiveState QRadioTuner__State = 0

// The tuner device is stopped.
const QRadioTuner__StoppedState QRadioTuner__State = 1

func (this *QRadioTuner) StateItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioTuner_StateItemName(val int) string {
	var nilthis *QRadioTuner
	return nilthis.StateItemName(val)
}

/*
Enumerates radio frequency bands.


*/
type QRadioTuner__Band = int

//
const QRadioTuner__AM QRadioTuner__Band = 0

//
const QRadioTuner__FM QRadioTuner__Band = 1

//
const QRadioTuner__SW QRadioTuner__Band = 2

//
const QRadioTuner__LW QRadioTuner__Band = 3

// range not defined, used when area supports more than one FM range.
const QRadioTuner__FM2 QRadioTuner__Band = 4

func (this *QRadioTuner) BandItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioTuner_BandItemName(val int) string {
	var nilthis *QRadioTuner
	return nilthis.BandItemName(val)
}

/*
Enumerates radio tuner error conditions.


*/
type QRadioTuner__Error = int

// No errors have occurred.
const QRadioTuner__NoError QRadioTuner__Error = 0

// There is no radio service available.
const QRadioTuner__ResourceError QRadioTuner__Error = 1

// Unable to open radio device.
const QRadioTuner__OpenError QRadioTuner__Error = 2

// An attempt to set a frequency or band that is not supported by radio device.
const QRadioTuner__OutOfRangeError QRadioTuner__Error = 3

func (this *QRadioTuner) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioTuner_ErrorItemName(val int) string {
	var nilthis *QRadioTuner
	return nilthis.ErrorItemName(val)
}

/*
Enumerates radio tuner policy for receiving stereo signals.


*/
type QRadioTuner__StereoMode = int

// Provide stereo mode, converting if required.
const QRadioTuner__ForceStereo QRadioTuner__StereoMode = 0

// Provide mono mode, converting if required.
const QRadioTuner__ForceMono QRadioTuner__StereoMode = 1

// Uses the stereo mode matching the station.
const QRadioTuner__Auto QRadioTuner__StereoMode = 2

func (this *QRadioTuner) StereoModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioTuner_StereoModeItemName(val int) string {
	var nilthis *QRadioTuner
	return nilthis.StereoModeItemName(val)
}

/*
Enumerates how the radio tuner should search for stations.


*/
type QRadioTuner__SearchMode = int

// Use only signal strength when searching.
const QRadioTuner__SearchFast QRadioTuner__SearchMode = 0

// After finding a strong signal, wait for the RDS station id (PI) before continuing.
const QRadioTuner__SearchGetStationId QRadioTuner__SearchMode = 1

func (this *QRadioTuner) SearchModeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioTuner_SearchModeItemName(val int) string {
	var nilthis *QRadioTuner
	return nilthis.SearchModeItemName(val)
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
