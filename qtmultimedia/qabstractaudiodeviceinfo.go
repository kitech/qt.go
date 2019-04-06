package qtmultimedia

// /usr/include/qt/QtMultimedia/qaudiosystem.h
// #include <qaudiosystem.h>
// #include <QtMultimedia>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 6
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
type QAbstractAudioDeviceInfo struct {
	*qtcore.QObject
}
type QAbstractAudioDeviceInfo_ITF interface {
	qtcore.QObject_ITF
	QAbstractAudioDeviceInfo_PTR() *QAbstractAudioDeviceInfo
}

func (ptr *QAbstractAudioDeviceInfo) QAbstractAudioDeviceInfo_PTR() *QAbstractAudioDeviceInfo {
	return ptr
}

func (this *QAbstractAudioDeviceInfo) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QAbstractAudioDeviceInfo) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQAbstractAudioDeviceInfoFromPointer(cthis unsafe.Pointer) *QAbstractAudioDeviceInfo {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QAbstractAudioDeviceInfo{bcthis0}
}
func (*QAbstractAudioDeviceInfo) NewFromPointer(cthis unsafe.Pointer) *QAbstractAudioDeviceInfo {
	return NewQAbstractAudioDeviceInfoFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QAbstractAudioDeviceInfo) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractAudioDeviceInfo10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:59
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QAudioFormat preferredFormat() const

/*

 */
func (this *QAbstractAudioDeviceInfo) PreferredFormat() *QAudioFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractAudioDeviceInfo15preferredFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQAudioFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQAudioFormat)
	return rv2
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:60
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool isFormatSupported(const QAudioFormat &) const

/*

 */
func (this *QAbstractAudioDeviceInfo) IsFormatSupported(format QAudioFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if format != nil && format.QAudioFormat_PTR() != nil {
		convArg0 = format.QAudioFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractAudioDeviceInfo17isFormatSupportedERK12QAudioFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QString deviceName() const

/*

 */
func (this *QAbstractAudioDeviceInfo) DeviceName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK24QAbstractAudioDeviceInfo10deviceNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qaudiosystem.h:62
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QStringList supportedCodecs()

/*

 */
func (this *QAbstractAudioDeviceInfo) SupportedCodecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractAudioDeviceInfo15supportedCodecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

func DeleteQAbstractAudioDeviceInfo(this *QAbstractAudioDeviceInfo) {
	rv, err := qtrt.InvokeQtFunc6("_ZN24QAbstractAudioDeviceInfoD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

//  body block end

//  keep block begin

func init_unused_11787() {
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
