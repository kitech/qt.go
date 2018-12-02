package qtmultimedia

// /usr/include/qt/QtMultimedia/qradiodata.h
// #include <qradiodata.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
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

// bool setMediaObject(QMediaObject *)
func (this *QRadioData) InheritSetMediaObject(f func(arg0 *QMediaObject /*777 QMediaObject **/) bool) {
	qtrt.SetAllInheritCallback(this, "setMediaObject", f)
}

/*

 */
type QRadioData struct {
	*qtcore.QObject
	*QMediaBindableInterface
}
type QRadioData_ITF interface {
	qtcore.QObject_ITF
	QMediaBindableInterface_ITF
	QRadioData_PTR() *QRadioData
}

func (ptr *QRadioData) QRadioData_PTR() *QRadioData { return ptr }

func (this *QRadioData) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QRadioData) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QMediaBindableInterface = NewQMediaBindableInterfaceFromPointer(cthis)
}
func NewQRadioDataFromPointer(cthis unsafe.Pointer) *QRadioData {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQMediaBindableInterfaceFromPointer(cthis)
	return &QRadioData{bcthis0, bcthis1}
}
func (*QRadioData) NewFromPointer(cthis unsafe.Pointer) *QRadioData {
	return NewQRadioDataFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:55
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRadioData) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qradiodata.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioData(QMediaObject *, QObject *)

/*
Constructs a radio data based on a mediaObject and parent.

The mediaObject should be an instance of QRadioTuner. It is preferable to use the radioData property on a QRadioTuner instance to get an instance of QRadioData.

During construction, this class is bound to the mediaObject using the bind() method.
*/
func (*QRadioData) NewForInherit(mediaObject QMediaObject_ITF /*777 QMediaObject **/, parent qtcore.QObject_ITF /*777 QObject **/) *QRadioData {
	return NewQRadioData(mediaObject, parent)
}
func NewQRadioData(mediaObject QMediaObject_ITF /*777 QMediaObject **/, parent qtcore.QObject_ITF /*777 QObject **/) *QRadioData {
	var convArg0 unsafe.Pointer
	if mediaObject != nil && mediaObject.QMediaObject_PTR() != nil {
		convArg0 = mediaObject.QMediaObject_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioDataC2EP12QMediaObjectP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioData")
	return gothis
}

// /usr/include/qt/QtMultimedia/qradiodata.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QRadioData(QMediaObject *, QObject *)

/*
Constructs a radio data based on a mediaObject and parent.

The mediaObject should be an instance of QRadioTuner. It is preferable to use the radioData property on a QRadioTuner instance to get an instance of QRadioData.

During construction, this class is bound to the mediaObject using the bind() method.
*/
func (*QRadioData) NewForInherit__(mediaObject QMediaObject_ITF /*777 QMediaObject **/) *QRadioData {
	return NewQRadioData__(mediaObject)
}
func NewQRadioData__(mediaObject QMediaObject_ITF /*777 QMediaObject **/) *QRadioData {
	var convArg0 unsafe.Pointer
	if mediaObject != nil && mediaObject.QMediaObject_PTR() != nil {
		convArg0 = mediaObject.QMediaObject_PTR().GetCthis()
	}
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioDataC2EP12QMediaObjectP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioDataFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioData")
	return gothis
}

// /usr/include/qt/QtMultimedia/qradiodata.h:86
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRadioData()

/*

 */
func DeleteQRadioData(this *QRadioData) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioDataD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] QMultimedia::AvailabilityStatus availability() const

/*
Returns the availability of the radio data service.

A long as there is a media service which provides radio functionality, then the availability will be that of the radio tuner.
*/
func (this *QRadioData) Availability() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData12availabilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QMediaObject * mediaObject() const

/*
Reimplemented from QMediaBindableInterface::mediaObject().

See also setMediaObject().
*/
func (this *QRadioData) MediaObject() *QMediaObject /*777 QMediaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData11mediaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMediaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qradiodata.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stationId() const

/*

 */
func (this *QRadioData) StationId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData9stationIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodata.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QRadioData::ProgramType programType() const

/*

 */
func (this *QRadioData) ProgramType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData11programTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QString programTypeName() const

/*

 */
func (this *QRadioData) ProgramTypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData15programTypeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodata.h:95
// index:0
// Public Visibility=Default Availability=Available
// [8] QString stationName() const

/*

 */
func (this *QRadioData) StationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData11stationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodata.h:96
// index:0
// Public Visibility=Default Availability=Available
// [8] QString radioText() const

/*

 */
func (this *QRadioData) RadioText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData9radioTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodata.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAlternativeFrequenciesEnabled() const

/*

 */
func (this *QRadioData) IsAlternativeFrequenciesEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData31isAlternativeFrequenciesEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiodata.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] QRadioData::Error error() const

/*
Returns the error state of a radio data.

See also errorString().
*/
func (this *QRadioData) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:113
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QRadioData::Error)

/*
Returns the error state of a radio data.

See also errorString().
*/
func (this *QRadioData) Error_1(error int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData5errorENS_5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), error)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:100
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a description of a radio data's error state.

See also error().
*/
func (this *QRadioData) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QRadioData11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodata.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlternativeFrequenciesEnabled(bool)

/*

 */
func (this *QRadioData) SetAlternativeFrequenciesEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData32setAlternativeFrequenciesEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stationIdChanged(QString)

/*
Signals that the Program Identification code has changed to stationId

Note: Notifier signal for property stationId.
*/
func (this *QRadioData) StationIdChanged(stationId string) {
	var tmpArg0 = qtcore.NewQString_5(stationId)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData16stationIdChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void programTypeChanged(QRadioData::ProgramType)

/*
Signals that the Program Type code has changed to programType

Note: Notifier signal for property programType.
*/
func (this *QRadioData) ProgramTypeChanged(programType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData18programTypeChangedENS_11ProgramTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), programType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:108
// index:0
// Public Visibility=Default Availability=Available
// [-2] void programTypeNameChanged(QString)

/*
Signals that the Program Type Name has changed to programTypeName

Note: Notifier signal for property programTypeName.
*/
func (this *QRadioData) ProgramTypeNameChanged(programTypeName string) {
	var tmpArg0 = qtcore.NewQString_5(programTypeName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData22programTypeNameChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stationNameChanged(QString)

/*
Signals that the Program Service has changed to stationName

Note: Notifier signal for property stationName.
*/
func (this *QRadioData) StationNameChanged(stationName string) {
	var tmpArg0 = qtcore.NewQString_5(stationName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData18stationNameChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void radioTextChanged(QString)

/*
Signals that the Radio Text property has changed to radioText

Note: Notifier signal for property radioText.
*/
func (this *QRadioData) RadioTextChanged(radioText string) {
	var tmpArg0 = qtcore.NewQString_5(radioText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData16radioTextChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void alternativeFrequenciesEnabledChanged(bool)

/*
Signals that automatically tuning to alternative frequencies has been enabled or disabled according to enabled.

Note: Notifier signal for property alternativeFrequenciesEnabled.
*/
func (this *QRadioData) AlternativeFrequenciesEnabledChanged(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData36alternativeFrequenciesEnabledChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodata.h:116
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool setMediaObject(QMediaObject *)

/*
Reimplemented from QMediaBindableInterface::setMediaObject().

See also mediaObject().
*/
func (this *QRadioData) SetMediaObject(arg0 QMediaObject_ITF /*777 QMediaObject **/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QMediaObject_PTR() != nil {
		convArg0 = arg0.QMediaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QRadioData14setMediaObjectEP12QMediaObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

/*
Enumerates radio data error conditions.


*/
type QRadioData__Error = int

// No errors have occurred.
const QRadioData__NoError QRadioData__Error = 0

// There is no radio service available.
const QRadioData__ResourceError QRadioData__Error = 1

// Unable to open radio device.
const QRadioData__OpenError QRadioData__Error = 2

// An attempt to set a frequency or band that is not supported by radio device.
const QRadioData__OutOfRangeError QRadioData__Error = 3

func (this *QRadioData) ErrorItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioData_ErrorItemName(val int) string {
	var nilthis *QRadioData
	return nilthis.ErrorItemName(val)
}

/*
This property holds the type of the currently playing program as transmitted by the radio station. The value can be any one of the values defined in the table below.

ConstantValue
QRadioData::Undefined0
QRadioData::News1
QRadioData::CurrentAffairs2
QRadioData::Information3
QRadioData::Sport4
QRadioData::Education5
QRadioData::Drama6
QRadioData::Culture7
QRadioData::Science8
QRadioData::Varied9

*/
type QRadioData__ProgramType = int

//
const QRadioData__Undefined QRadioData__ProgramType = 0

//
const QRadioData__News QRadioData__ProgramType = 1

//
const QRadioData__CurrentAffairs QRadioData__ProgramType = 2

//
const QRadioData__Information QRadioData__ProgramType = 3

//
const QRadioData__Sport QRadioData__ProgramType = 4

//
const QRadioData__Education QRadioData__ProgramType = 5

//
const QRadioData__Drama QRadioData__ProgramType = 6

//
const QRadioData__Culture QRadioData__ProgramType = 7

//
const QRadioData__Science QRadioData__ProgramType = 8

//
const QRadioData__Varied QRadioData__ProgramType = 9

// 0
const QRadioData__PopMusic QRadioData__ProgramType = 10

// 1
const QRadioData__RockMusic QRadioData__ProgramType = 11

// 2
const QRadioData__EasyListening QRadioData__ProgramType = 12

// 3
const QRadioData__LightClassical QRadioData__ProgramType = 13

// 4
const QRadioData__SeriousClassical QRadioData__ProgramType = 14

// 5
const QRadioData__OtherMusic QRadioData__ProgramType = 15

// 6
const QRadioData__Weather QRadioData__ProgramType = 16

// 7
const QRadioData__Finance QRadioData__ProgramType = 17

// 8
const QRadioData__ChildrensProgrammes QRadioData__ProgramType = 18

// 9
const QRadioData__SocialAffairs QRadioData__ProgramType = 19

// 0
const QRadioData__Religion QRadioData__ProgramType = 20

// 1
const QRadioData__PhoneIn QRadioData__ProgramType = 21

// 2
const QRadioData__Travel QRadioData__ProgramType = 22

// 3
const QRadioData__Leisure QRadioData__ProgramType = 23

// 4
const QRadioData__JazzMusic QRadioData__ProgramType = 24

// 5
const QRadioData__CountryMusic QRadioData__ProgramType = 25

// 6
const QRadioData__NationalMusic QRadioData__ProgramType = 26

// 7
const QRadioData__OldiesMusic QRadioData__ProgramType = 27

// 8
const QRadioData__FolkMusic QRadioData__ProgramType = 28

// 9
const QRadioData__Documentary QRadioData__ProgramType = 29

// 0
const QRadioData__AlarmTest QRadioData__ProgramType = 30

// 1
const QRadioData__Alarm QRadioData__ProgramType = 31

// 2
const QRadioData__Talk QRadioData__ProgramType = 32

// 3
const QRadioData__ClassicRock QRadioData__ProgramType = 33

// 4
const QRadioData__AdultHits QRadioData__ProgramType = 34

// 5
const QRadioData__SoftRock QRadioData__ProgramType = 35

// 6
const QRadioData__Top40 QRadioData__ProgramType = 36

// 7
const QRadioData__Soft QRadioData__ProgramType = 37

// 8
const QRadioData__Nostalgia QRadioData__ProgramType = 38

// 9
const QRadioData__Classical QRadioData__ProgramType = 39

// 0
const QRadioData__RhythmAndBlues QRadioData__ProgramType = 40

// 1
const QRadioData__SoftRhythmAndBlues QRadioData__ProgramType = 41

// 2
const QRadioData__Language QRadioData__ProgramType = 42

// 3
const QRadioData__ReligiousMusic QRadioData__ProgramType = 43

// 4
const QRadioData__ReligiousTalk QRadioData__ProgramType = 44

// 5
const QRadioData__Personality QRadioData__ProgramType = 45

// 6
const QRadioData__Public QRadioData__ProgramType = 46

// 7
const QRadioData__College QRadioData__ProgramType = 47

func (this *QRadioData) ProgramTypeItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QRadioData_ProgramTypeItemName(val int) string {
	var nilthis *QRadioData
	return nilthis.ProgramTypeItemName(val)
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
