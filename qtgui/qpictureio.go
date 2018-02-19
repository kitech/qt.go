package qtgui

// /usr/include/qt/QtGui/qpicture.h
// #include <qpicture.h>
// #include <QtGui>

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

//  ext block end

//  body block begin

type QPictureIO struct {
	*qtrt.CObject
}
type QPictureIO_ITF interface {
	QPictureIO_PTR() *QPictureIO
}

func (ptr *QPictureIO) QPictureIO_PTR() *QPictureIO { return ptr }

func (this *QPictureIO) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPictureIO) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQPictureIOFromPointer(cthis unsafe.Pointer) *QPictureIO {
	return &QPictureIO{&qtrt.CObject{cthis}}
}
func (*QPictureIO) NewFromPointer(cthis unsafe.Pointer) *QPictureIO {
	return NewQPictureIOFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpicture.h:134
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPictureIO()
func NewQPictureIO() *QPictureIO {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIOC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureIOFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPictureIO)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:135
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPictureIO(QIODevice *, const char *)
func NewQPictureIO_1(ioDevice qtcore.QIODevice_ITF /*777 QIODevice **/, format string) *QPictureIO {
	var convArg0 unsafe.Pointer
	if ioDevice != nil && ioDevice.QIODevice_PTR() != nil {
		convArg0 = ioDevice.QIODevice_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIOC2EP9QIODevicePKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureIOFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPictureIO)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:136
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QPictureIO(const QString &, const char *)
func NewQPictureIO_2(fileName string, format string) *QPictureIO {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIOC2ERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPictureIOFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPictureIO)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QPictureIO()
func DeleteQPictureIO(this *QPictureIO) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIOD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpicture.h:139
// index:0
// Public Visibility=Default Availability=Available
// [32] const QPicture & picture() const
func (this *QPictureIO) Picture() *QPicture {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO7pictureEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPicture)
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:140
// index:0
// Public Visibility=Default Availability=Available
// [4] int status() const
func (this *QPictureIO) Status() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO6statusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpicture.h:141
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * format() const
func (this *QPictureIO) Format() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qpicture.h:142
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * ioDevice() const
func (this *QPictureIO) IoDevice() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO8ioDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qpicture.h:143
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const
func (this *QPictureIO) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpicture.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] int quality() const
func (this *QPictureIO) Quality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO7qualityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpicture.h:145
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description() const
func (this *QPictureIO) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qpicture.h:146
// index:0
// Public Visibility=Default Availability=Available
// [8] const char * parameters() const
func (this *QPictureIO) Parameters() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO10parametersEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qpicture.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] float gamma() const
func (this *QPictureIO) Gamma() float32 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QPictureIO5gammaEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtGui/qpicture.h:149
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPicture(const QPicture &)
func (this *QPictureIO) SetPicture(arg0 QPicture_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPicture_PTR() != nil {
		convArg0 = arg0.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO10setPictureERK8QPicture", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:150
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatus(int)
func (this *QPictureIO) SetStatus(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO9setStatusEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:151
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const char *)
func (this *QPictureIO) SetFormat(arg0 string) {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO9setFormatEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:152
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIODevice(QIODevice *)
func (this *QPictureIO) SetIODevice(arg0 qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QIODevice_PTR() != nil {
		convArg0 = arg0.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO11setIODeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:153
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)
func (this *QPictureIO) SetFileName(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuality(int)
func (this *QPictureIO) SetQuality(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO10setQualityEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)
func (this *QPictureIO) SetDescription(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParameters(const char *)
func (this *QPictureIO) SetParameters(arg0 string) {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO13setParametersEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGamma(float)
func (this *QPictureIO) SetGamma(arg0 float32) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO8setGammaEf", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:159
// index:0
// Public Visibility=Default Availability=Available
// [1] bool read()
func (this *QPictureIO) Read() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO4readEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:160
// index:0
// Public Visibility=Default Availability=Available
// [1] bool write()
func (this *QPictureIO) Write() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO5writeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:162
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray pictureFormat(const QString &)
func (this *QPictureIO) PictureFormat(fileName string) *qtcore.QByteArray /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO13pictureFormatERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}
func QPictureIO_PictureFormat(fileName string) *qtcore.QByteArray /*123*/ {
	var nilthis *QPictureIO
	rv := nilthis.PictureFormat(fileName)
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:163
// index:1
// Public static Visibility=Default Availability=Available
// [8] QByteArray pictureFormat(QIODevice *)
func (this *QPictureIO) PictureFormat_1(arg0 qtcore.QIODevice_ITF /*777 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QIODevice_PTR() != nil {
		convArg0 = arg0.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QPictureIO13pictureFormatEP9QIODevice", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}
func QPictureIO_PictureFormat_1(arg0 qtcore.QIODevice_ITF /*777 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var nilthis *QPictureIO
	rv := nilthis.PictureFormat_1(arg0)
	return rv
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
