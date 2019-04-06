package qtmultimedia

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h
// #include <qmediaserviceproviderplugin.h>
// #include <QtMultimedia>

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
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

/*

 */
type QMediaServiceProviderHint struct {
	*qtrt.CObject
}
type QMediaServiceProviderHint_ITF interface {
	QMediaServiceProviderHint_PTR() *QMediaServiceProviderHint
}

func (ptr *QMediaServiceProviderHint) QMediaServiceProviderHint_PTR() *QMediaServiceProviderHint {
	return ptr
}

func (this *QMediaServiceProviderHint) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaServiceProviderHint) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaServiceProviderHintFromPointer(cthis unsafe.Pointer) *QMediaServiceProviderHint {
	return &QMediaServiceProviderHint{&qtrt.CObject{cthis}}
}
func (*QMediaServiceProviderHint) NewFromPointer(cthis unsafe.Pointer) *QMediaServiceProviderHint {
	return NewQMediaServiceProviderHintFromPointer(cthis)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QMediaServiceProviderHint()

/*

 */
func (*QMediaServiceProviderHint) NewForInherit() *QMediaServiceProviderHint {
	return NewQMediaServiceProviderHint()
}
func NewQMediaServiceProviderHint() *QMediaServiceProviderHint {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaServiceProviderHintFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaServiceProviderHint)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:75
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QMediaServiceProviderHint(const QString &, const QStringList &)

/*

 */
func (*QMediaServiceProviderHint) NewForInherit1(mimeType string, codecs qtcore.QStringList_ITF) *QMediaServiceProviderHint {
	return NewQMediaServiceProviderHint1(mimeType, codecs)
}
func NewQMediaServiceProviderHint1(mimeType string, codecs qtcore.QStringList_ITF) *QMediaServiceProviderHint {
	var tmpArg0 = qtcore.NewQString5(mimeType)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if codecs != nil && codecs.QStringList_PTR() != nil {
		convArg1 = codecs.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintC2ERK7QStringRK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaServiceProviderHintFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaServiceProviderHint)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:76
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QMediaServiceProviderHint(const QByteArray &)

/*

 */
func (*QMediaServiceProviderHint) NewForInherit2(device qtcore.QByteArray_ITF) *QMediaServiceProviderHint {
	return NewQMediaServiceProviderHint2(device)
}
func NewQMediaServiceProviderHint2(device qtcore.QByteArray_ITF) *QMediaServiceProviderHint {
	var convArg0 unsafe.Pointer
	if device != nil && device.QByteArray_PTR() != nil {
		convArg0 = device.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaServiceProviderHintFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaServiceProviderHint)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:77
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QMediaServiceProviderHint(QCamera::Position)

/*

 */
func (*QMediaServiceProviderHint) NewForInherit3(position int) *QMediaServiceProviderHint {
	return NewQMediaServiceProviderHint3(position)
}
func NewQMediaServiceProviderHint3(position int) *QMediaServiceProviderHint {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintC2EN7QCamera8PositionE", qtrt.FFI_TYPE_POINTER, position)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaServiceProviderHintFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaServiceProviderHint)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:78
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QMediaServiceProviderHint(QMediaServiceProviderHint::Features)

/*

 */
func (*QMediaServiceProviderHint) NewForInherit4(features int) *QMediaServiceProviderHint {
	return NewQMediaServiceProviderHint4(features)
}
func NewQMediaServiceProviderHint4(features int) *QMediaServiceProviderHint {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintC2E6QFlagsINS_7FeatureEE", qtrt.FFI_TYPE_POINTER, features)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMediaServiceProviderHintFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMediaServiceProviderHint)
	return gothis
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QMediaServiceProviderHint()

/*

 */
func DeleteQMediaServiceProviderHint(this *QMediaServiceProviderHint) {
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:82
// index:0
// Public Visibility=Default Availability=Available
// [8] QMediaServiceProviderHint & operator=(const QMediaServiceProviderHint &)

/*

 */
func (this *QMediaServiceProviderHint) Operator_equal(other QMediaServiceProviderHint_ITF) *QMediaServiceProviderHint {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaServiceProviderHint_PTR() != nil {
		convArg0 = other.QMediaServiceProviderHint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN25QMediaServiceProviderHintaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQMediaServiceProviderHintFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQMediaServiceProviderHint)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:84
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QMediaServiceProviderHint &) const

/*

 */
func (this *QMediaServiceProviderHint) Operator_equal_equal(other QMediaServiceProviderHint_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaServiceProviderHint_PTR() != nil {
		convArg0 = other.QMediaServiceProviderHint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHinteqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:85
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QMediaServiceProviderHint &) const

/*

 */
func (this *QMediaServiceProviderHint) Operator_not_equal(other QMediaServiceProviderHint_ITF) bool {
	var convArg0 unsafe.Pointer
	if other != nil && other.QMediaServiceProviderHint_PTR() != nil {
		convArg0 = other.QMediaServiceProviderHint_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHintneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:87
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull() const

/*

 */
func (this *QMediaServiceProviderHint) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaServiceProviderHint::Type type() const

/*

 */
func (this *QMediaServiceProviderHint) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString mimeType() const

/*

 */
func (this *QMediaServiceProviderHint) MimeType() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint8mimeTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:92
// index:0
// Public Visibility=Default Availability=Available
// [8] QStringList codecs() const

/*

 */
func (this *QMediaServiceProviderHint) Codecs() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint6codecsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:94
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray device() const

/*

 */
func (this *QMediaServiceProviderHint) Device() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] QCamera::Position cameraPosition() const

/*

 */
func (this *QMediaServiceProviderHint) CameraPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint14cameraPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtMultimedia/qmediaserviceproviderplugin.h:97
// index:0
// Public Visibility=Default Availability=Available
// [4] QMediaServiceProviderHint::Features features() const

/*

 */
func (this *QMediaServiceProviderHint) Features() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK25QMediaServiceProviderHint8featuresEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*


 */
type QMediaServiceProviderHint__Type = int

//
const QMediaServiceProviderHint__Null QMediaServiceProviderHint__Type = 0

//
const QMediaServiceProviderHint__ContentType QMediaServiceProviderHint__Type = 1

//
const QMediaServiceProviderHint__Device QMediaServiceProviderHint__Type = 2

//
const QMediaServiceProviderHint__SupportedFeatures QMediaServiceProviderHint__Type = 3

//
const QMediaServiceProviderHint__CameraPosition QMediaServiceProviderHint__Type = 4

func (this *QMediaServiceProviderHint) TypeItemName(val int) string {
	switch val {
	case QMediaServiceProviderHint__Null: // 0
		return "Null"
	case QMediaServiceProviderHint__ContentType: // 1
		return "ContentType"
	case QMediaServiceProviderHint__Device: // 2
		return "Device"
	case QMediaServiceProviderHint__SupportedFeatures: // 3
		return "SupportedFeatures"
	case QMediaServiceProviderHint__CameraPosition: // 4
		return "CameraPosition"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMediaServiceProviderHint_TypeItemName(val int) string {
	var nilthis *QMediaServiceProviderHint
	return nilthis.TypeItemName(val)
}

/*


 */
type QMediaServiceProviderHint__Feature = int

//
const QMediaServiceProviderHint__LowLatencyPlayback QMediaServiceProviderHint__Feature = 1

//
const QMediaServiceProviderHint__RecordingSupport QMediaServiceProviderHint__Feature = 2

//
const QMediaServiceProviderHint__StreamPlayback QMediaServiceProviderHint__Feature = 4

//
const QMediaServiceProviderHint__VideoSurface QMediaServiceProviderHint__Feature = 8

func (this *QMediaServiceProviderHint) FeatureItemName(val int) string {
	switch val {
	case QMediaServiceProviderHint__LowLatencyPlayback: // 1
		return "LowLatencyPlayback"
	case QMediaServiceProviderHint__RecordingSupport: // 2
		return "RecordingSupport"
	case QMediaServiceProviderHint__StreamPlayback: // 4
		return "StreamPlayback"
	case QMediaServiceProviderHint__VideoSurface: // 8
		return "VideoSurface"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QMediaServiceProviderHint_FeatureItemName(val int) string {
	var nilthis *QMediaServiceProviderHint
	return nilthis.FeatureItemName(val)
}

//  body block end

//  keep block begin

func init_unused_11873() {
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
