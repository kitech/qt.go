package qtmultimedia

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h
// #include <qradiodatacontrol.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 22
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
type QRadioDataControl struct {
	*QMediaControl
}
type QRadioDataControl_ITF interface {
	QMediaControl_ITF
	QRadioDataControl_PTR() *QRadioDataControl
}

func (ptr *QRadioDataControl) QRadioDataControl_PTR() *QRadioDataControl { return ptr }

func (this *QRadioDataControl) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QMediaControl.GetCthis()
	}
}
func (this *QRadioDataControl) SetCthis(cthis unsafe.Pointer) {
	this.QMediaControl = NewQMediaControlFromPointer(cthis)
}
func NewQRadioDataControlFromPointer(cthis unsafe.Pointer) *QRadioDataControl {
	bcthis0 := NewQMediaControlFromPointer(cthis)
	return &QRadioDataControl{bcthis0}
}
func (*QRadioDataControl) NewFromPointer(cthis unsafe.Pointer) *QRadioDataControl {
	return NewQRadioDataControlFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QRadioDataControl) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QRadioDataControl()

/*

 */
func DeleteQRadioDataControl(this *QRadioDataControl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:58
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString stationId() const

/*
Returns the current Program Identification
*/
func (this *QRadioDataControl) StationId() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl9stationIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QRadioData::ProgramType programType() const

/*
Returns the current Program Type
*/
func (this *QRadioDataControl) ProgramType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl11programTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString programTypeName() const

/*
Returns the current Program Type Name
*/
func (this *QRadioDataControl) ProgramTypeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl15programTypeNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString stationName() const

/*
Returns the current Program Service
*/
func (this *QRadioDataControl) StationName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl11stationNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString radioText() const

/*
Returns the current Radio Text
*/
func (this *QRadioDataControl) RadioText() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl9radioTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:63
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [-2] void setAlternativeFrequenciesEnabled(bool)

/*
Sets the Alternative Frequency to enabled

See also isAlternativeFrequenciesEnabled().
*/
func (this *QRadioDataControl) SetAlternativeFrequenciesEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl32setAlternativeFrequenciesEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:64
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isAlternativeFrequenciesEnabled() const

/*
Returns true if Alternative Frequency is currently enabled
*/
func (this *QRadioDataControl) IsAlternativeFrequenciesEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl31isAlternativeFrequenciesEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:66
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] QRadioData::Error error() const

/*
Returns the error state of a radio data.
*/
func (this *QRadioDataControl) Error() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl5errorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:76
// index:1
// Public Visibility=Default Availability=Available
// [-2] void error(QRadioData::Error)

/*
Returns the error state of a radio data.
*/
func (this *QRadioDataControl) Error1(err_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl5errorEN10QRadioData5ErrorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), err_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:67
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a string describing a radio data's error state.
*/
func (this *QRadioDataControl) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QRadioDataControl11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stationIdChanged(QString)

/*
Signals that the Program Identification stationId has changed
*/
func (this *QRadioDataControl) StationIdChanged(stationId string) {
	var tmpArg0 = qtcore.NewQString5(stationId)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl16stationIdChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void programTypeChanged(QRadioData::ProgramType)

/*
Signals that the Program Type programType has changed
*/
func (this *QRadioDataControl) ProgramTypeChanged(programType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl18programTypeChangedEN10QRadioData11ProgramTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), programType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:72
// index:0
// Public Visibility=Default Availability=Available
// [-2] void programTypeNameChanged(QString)

/*
Signals that the Program Type Name programTypeName has changed
*/
func (this *QRadioDataControl) ProgramTypeNameChanged(programTypeName string) {
	var tmpArg0 = qtcore.NewQString5(programTypeName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl22programTypeNameChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:73
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stationNameChanged(QString)

/*
Signals that the Program Service stationName has changed
*/
func (this *QRadioDataControl) StationNameChanged(stationName string) {
	var tmpArg0 = qtcore.NewQString5(stationName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl18stationNameChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void radioTextChanged(QString)

/*
Signals that the Radio Text radioText has changed
*/
func (this *QRadioDataControl) RadioTextChanged(radioText string) {
	var tmpArg0 = qtcore.NewQString5(radioText)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl16radioTextChangedE7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void alternativeFrequenciesEnabledChanged(bool)

/*
Signals that the alternative frequencies setting has changed to the value of enabled.
*/
func (this *QRadioDataControl) AlternativeFrequenciesEnabledChanged(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControl36alternativeFrequenciesEnabledChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:79
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QRadioDataControl(QObject *)

/*
Constructs a radio data control with the given parent.
*/
func (*QRadioDataControl) NewForInherit(parent qtcore.QObject_ITF /*777 QObject **/) *QRadioDataControl {
	return NewQRadioDataControl(parent)
}
func NewQRadioDataControl(parent qtcore.QObject_ITF /*777 QObject **/) *QRadioDataControl {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioDataControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioDataControl")
	return gothis
}

// /usr/include/qt/QtMultimedia/qradiodatacontrol.h:79
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QRadioDataControl(QObject *)

/*
Constructs a radio data control with the given parent.
*/
func (*QRadioDataControl) NewForInheritp() *QRadioDataControl {
	return NewQRadioDataControlp()
}
func NewQRadioDataControlp() *QRadioDataControl {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN17QRadioDataControlC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQRadioDataControlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QRadioDataControl")
	return gothis
}

//  body block end

//  keep block begin

func init_unused_11899() {
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
