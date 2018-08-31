package qtmultimedia

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h
// #include <qradiotunercontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 41
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
type QRadioTunerControl struct {
	*QMediaControl
}
type QRadioTunerControl_ITF interface {
	QMediaControl_ITF
	QRadioTunerControl_PTR() *QRadioTunerControl
}

func (ptr *QRadioTunerControl) QRadioTunerControl_PTR() *QRadioTunerControl { return ptr }

func (this *QRadioTunerControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QRadioTunerControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQRadioTunerControlFromPointer(cthis unsafe.Pointer) *QRadioTunerControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QRadioTunerControl{bcthis0}
}
func (*QRadioTunerControl) NewFromPointer(cthis unsafe.Pointer) *QRadioTunerControl {
	return NewQRadioTunerControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRadioTunerControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRadioTunerControl()

/*

 */
func DeleteQRadioTunerControl(this *QRadioTunerControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QRadioTuner::State state() const

/*
Returns the current radio tuner state.
*/
func (this *QRadioTunerControl) State() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl5stateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QRadioTuner::Band band() const

/*
Returns the frequency band a radio tuner is tuned to.

See also setBand().
*/
func (this *QRadioTunerControl) Band() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl4bandEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setBand(QRadioTuner::Band)

/*
Sets the frequecy band a radio tuner is tuned to.

Changing the frequency band will reset the frequency to the minimum frequency of the new band.

See also band().
*/
func (this *QRadioTunerControl) SetBand(b int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl7setBandEN11QRadioTuner4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isBandSupported(QRadioTuner::Band) const

/*
Identifies if a frequency band is supported.

Returns true if the band is supported, and false if it is not.
*/
func (this *QRadioTunerControl) IsBandSupported(b int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl15isBandSupportedEN11QRadioTuner4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int frequency() const

/*
Returns the frequency a radio tuner is tuned to.

See also setFrequency().
*/
func (this *QRadioTunerControl) Frequency() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl9frequencyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:65
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int frequencyStep(QRadioTuner::Band) const

/*
Returns the number of Hertz to increment the frequency by when stepping through frequencies within a given band.
*/
func (this *QRadioTunerControl) FrequencyStep(b int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl13frequencyStepEN11QRadioTuner4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), b)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setFrequency(int)

/*
Sets the frequency a radio tuner is tuned to.

See also frequency().
*/
func (this *QRadioTunerControl) SetFrequency(frequency int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl12setFrequencyEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:69
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isStereo() const

/*
Identifies if a radio tuner is receiving a stereo signal.

Returns true if the tuner is receiving a stereo signal, and false if it is not.
*/
func (this *QRadioTunerControl) IsStereo() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl8isStereoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:70
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QRadioTuner::StereoMode stereoMode() const

/*
Returns a radio tuner's stereo mode.

See also setStereoMode() and QRadioTuner::StereoMode.
*/
func (this *QRadioTunerControl) StereoMode() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl10stereoModeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:71
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setStereoMode(QRadioTuner::StereoMode)

/*
Sets a radio tuner's stereo mode.

See also stereoMode() and QRadioTuner::StereoMode.
*/
func (this *QRadioTunerControl) SetStereoMode(mode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl13setStereoModeEN11QRadioTuner10StereoModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:73
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int signalStrength() const

/*
Return a radio tuner's current signal strength as a percentage.
*/
func (this *QRadioTunerControl) SignalStrength() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl14signalStrengthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int volume() const

/*
Returns the volume of a radio tuner's audio output as a percentage.

See also setVolume().
*/
func (this *QRadioTunerControl) Volume() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl6volumeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:76
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setVolume(int)

/*
Sets the percentage volume of a radio tuner's audio output.

See also volume().
*/
func (this *QRadioTunerControl) SetVolume(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl9setVolumeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:78
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isMuted() const

/*
Identifies if a radio tuner's audio output is muted.

Returns true if the audio is muted, and false if it is not.
*/
func (this *QRadioTunerControl) IsMuted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl7isMutedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:79
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setMuted(bool)

/*
Sets the muted state of a radio tuner's audio output.

See also isMuted().
*/
func (this *QRadioTunerControl) SetMuted(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl8setMutedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:81
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isSearching() const

/*
Identifies if a radio tuner is currently scanning for signal.

Returns true if the tuner is scanning, and false if it is not.
*/
func (this *QRadioTunerControl) IsSearching() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl11isSearchingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:83
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [1] bool isAntennaConnected() const

/*
Identifies if there is an antenna connected to the device.

Returns true if there is a connected antenna, and false otherwise.
*/
func (this *QRadioTunerControl) IsAntennaConnected() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl18isAntennaConnectedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:85
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void searchForward()

/*
Starts a forward scan for a signal, starting from the current frequency().
*/
func (this *QRadioTunerControl) SearchForward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl13searchForwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:86
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void searchBackward()

/*
Starts a backwards scan for a signal, starting from the current frequency().
*/
func (this *QRadioTunerControl) SearchBackward() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl14searchBackwardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void searchAllStations(QRadioTuner::SearchMode)

/*
Starts a scan through the whole frequency band searching all stations with a specific searchMode.
*/
func (this *QRadioTunerControl) SearchAllStations(searchMode int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl17searchAllStationsEN11QRadioTuner10SearchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), searchMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:87
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void searchAllStations(QRadioTuner::SearchMode)

/*
Starts a scan through the whole frequency band searching all stations with a specific searchMode.
*/
func (this *QRadioTunerControl) SearchAllStations__() {
	// arg: 0, QRadioTuner::SearchMode=Elaborated, QRadioTuner::SearchMode=Enum, , Invalid
	searchMode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl17searchAllStationsEN11QRadioTuner10SearchModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), searchMode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:88
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void cancelSearch()

/*
Stops scanning for a signal.
*/
func (this *QRadioTunerControl) CancelSearch() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl12cancelSearchEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:90
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void start()

/*
Activate the radio device.
*/
func (this *QRadioTunerControl) Start() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl5startEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:91
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void stop()

/*
Deactivate the radio device.
*/
func (this *QRadioTunerControl) Stop() {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl4stopEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:93
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QRadioTuner::Error error() const

/*
Returns the error state of a radio tuner.
*/
func (this *QRadioTunerControl) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:105
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QRadioTuner::Error)

/*
Returns the error state of a radio tuner.
*/
func (this *QRadioTunerControl) Error_1(err_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl5errorEN11QRadioTuner5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), err_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:94
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a string describing a radio tuner's error state.
*/
func (this *QRadioTunerControl) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QRadioTunerControl11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stateChanged(QRadioTuner::State)

/*
Signals that the state of a radio tuner has changed.
*/
func (this *QRadioTunerControl) StateChanged(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl12stateChangedEN11QRadioTuner5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:98
// index:0
// Public Visibility=Default Availability=Available
// [-2] void bandChanged(QRadioTuner::Band)

/*
Signals that the frequency band a radio tuner is tuned to has changed.
*/
func (this *QRadioTunerControl) BandChanged(band int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl11bandChangedEN11QRadioTuner4BandE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), band)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void frequencyChanged(int)

/*
Signals that the frequency a radio tuner is tuned to has changed.
*/
func (this *QRadioTunerControl) FrequencyChanged(frequency int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl16frequencyChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stereoStatusChanged(bool)

/*
Signals that the stereo state of a radio tuner has changed.
*/
func (this *QRadioTunerControl) StereoStatusChanged(stereo bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl19stereoStatusChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), stereo)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void searchingChanged(bool)

/*
Signals that the searching state of a radio tuner has changed.
*/
func (this *QRadioTunerControl) SearchingChanged(searching bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl16searchingChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), searching)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void signalStrengthChanged(int)

/*
Signals that the percentage strength of the signal received by a radio tuner has changed.
*/
func (this *QRadioTunerControl) SignalStrengthChanged(signalStrength int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl21signalStrengthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), signalStrength)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void volumeChanged(int)

/*
Signals that the percentage volume of radio tuner's audio output has changed.
*/
func (this *QRadioTunerControl) VolumeChanged(volume int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl13volumeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), volume)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mutedChanged(bool)

/*
Signals that the muted state of a radio tuner's audio output has changed.
*/
func (this *QRadioTunerControl) MutedChanged(muted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl12mutedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), muted)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stationFound(int, QString)

/*
Signals that new station with frequency and stationId was found when scanning
*/
func (this *QRadioTunerControl) StationFound(frequency int, stationId string) {
	var tmpArg1 = qtcore.NewQString_5(stationId)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl12stationFoundEi7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), frequency, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void antennaConnectedChanged(bool)

/*
Signals that the antenna has either been connected or disconnected as reflected with the connectionStatus.
*/
func (this *QRadioTunerControl) AntennaConnectedChanged(connectionStatus bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControl23antennaConnectedChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), connectionStatus)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:110
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QRadioTunerControl(QObject *)

/*
Constructs a radio tuner control with the given parent.
*/
func (*QRadioTunerControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QRadioTunerControl {
	return NewQRadioTunerControl(parent)
}
func NewQRadioTunerControl(parent qtcore.QObject_ITF /*777 QObject **/) *QRadioTunerControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioTunerControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioTunerControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qradiotunercontrol.h:110
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QRadioTunerControl(QObject *)

/*
Constructs a radio tuner control with the given parent.
*/
func (*QRadioTunerControl) NewForInherit__() *QRadioTunerControl {
	return NewQRadioTunerControl__()
}
func NewQRadioTunerControl__() *QRadioTunerControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN18QRadioTunerControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioTunerControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioTunerControl")
	return gothis
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
