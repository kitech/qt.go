package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h
// #include <qaudiosystemplugin.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
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
type QAudioSystemFactoryInterface struct {
	*qtrt.CObject
}
type QAudioSystemFactoryInterface_ITF interface {
	QAudioSystemFactoryInterface_PTR() *QAudioSystemFactoryInterface
}

func (ptr *QAudioSystemFactoryInterface) QAudioSystemFactoryInterface_PTR() *QAudioSystemFactoryInterface {
	return ptr
}

func (this *QAudioSystemFactoryInterface) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QAudioSystemFactoryInterface) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQAudioSystemFactoryInterfaceFromPointer(cthis unsafe.Pointer) *QAudioSystemFactoryInterface {
	return &QAudioSystemFactoryInterface{&qtrt.CObject{cthis}}
}
func (*QAudioSystemFactoryInterface) NewFromPointer(cthis unsafe.Pointer) *QAudioSystemFactoryInterface {
	return NewQAudioSystemFactoryInterfaceFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractAudioInput * createInput(const QByteArray &)

/*
Returns a pointer to a QAbstractAudioInput created using device identifier
*/
func (this *QAudioSystemFactoryInterface) CreateInput(device qtcore.QByteArray_ITF) *QAbstractAudioInput /*777 QAbstractAudioInput **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioSystemFactoryInterface11createInputERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAudioInputFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractAudioOutput * createOutput(const QByteArray &)

/*
Returns a pointer to a QAbstractAudioOutput created using device identifier
*/
func (this *QAudioSystemFactoryInterface) CreateOutput(device qtcore.QByteArray_ITF) *QAbstractAudioOutput /*777 QAbstractAudioOutput **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioSystemFactoryInterface12createOutputERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAudioOutputFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAbstractAudioDeviceInfo * createDeviceInfo(const QByteArray &, QAudio::Mode)

/*
Returns a pointer to a QAbstractAudioDeviceInfo created using device and mode
*/
func (this *QAudioSystemFactoryInterface) CreateDeviceInfo(device qtcore.QByteArray_ITF, mode int) *QAbstractAudioDeviceInfo /*777 QAbstractAudioDeviceInfo **/ {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioSystemFactoryInterface16createDeviceInfoERK10QByteArrayN6QAudio4ModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAbstractAudioDeviceInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystemplugin.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAudioSystemFactoryInterface()

/*

 */
func DeleteQAudioSystemFactoryInterface(this *QAudioSystemFactoryInterface) {
	rv, err := qtrt.InvokeQtFunc6("_ZN28QAudioSystemFactoryInterfaceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
