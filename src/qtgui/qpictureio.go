package qtgui

// /usr/include/qt/QtGui/qpicture.h
// #include <qpicture.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QPictureIO struct {
	*qtrt.CObject
}

func (this *QPictureIO) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QPictureIO) SetCthis(cthis unsafe.Pointer) {
	this.CObject = &qtrt.CObject{cthis}
}
func NewQPictureIOFromPointer(cthis unsafe.Pointer) *QPictureIO {
	return &QPictureIO{&qtrt.CObject{cthis}}
}
func (*QPictureIO) NewFromPointer(cthis unsafe.Pointer) *QPictureIO {
	return NewQPictureIOFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpicture.h:134
// index:0
// Public
// void QPictureIO()
func NewQPictureIO() *QPictureIO {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPictureIOFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:135
// index:1
// Public
// void QPictureIO(class QIODevice *, const char *)
func NewQPictureIO_1(ioDevice *qtcore.QIODevice /*444 QIODevice **/, format string) *QPictureIO {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = ioDevice.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOC2EP9QIODevicePKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPictureIOFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:136
// index:2
// Public
// void QPictureIO(const class QString &, const char *)
func NewQPictureIO_2(fileName *qtcore.QString, format string) *QPictureIO {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = fileName.GetCthis()
	var convArg1 = qtrt.CString(format)
	defer qtrt.FreeMem(convArg1)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOC2ERK7QStringPKc", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	gothis := NewQPictureIOFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpicture.h:137
// index:0
// Public
// void ~QPictureIO()
func DeleteQPictureIO(*QPictureIO) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIOD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:139
// index:0
// Public
// const QPicture & picture()
func (this *QPictureIO) Picture() *QPicture {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO7pictureEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPictureFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:140
// index:0
// Public
// int status()
func (this *QPictureIO) Status() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO6statusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpicture.h:141
// index:0
// Public
// const char * format()
func (this *QPictureIO) Format() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO6formatEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qpicture.h:142
// index:0
// Public
// QIODevice * ioDevice()
func (this *QPictureIO) IoDevice() *qtcore.QIODevice /*444 QIODevice **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO8ioDeviceEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:143
// index:0
// Public
// QString fileName()
func (this *QPictureIO) FileName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO8fileNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:144
// index:0
// Public
// int quality()
func (this *QPictureIO) Quality() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO7qualityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpicture.h:145
// index:0
// Public
// QString description()
func (this *QPictureIO) Description() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO11descriptionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpicture.h:146
// index:0
// Public
// const char * parameters()
func (this *QPictureIO) Parameters() string {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO10parametersEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qpicture.h:147
// index:0
// Public
// float gamma()
func (this *QPictureIO) Gamma() float32 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPictureIO5gammaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float32(rv) // 111
}

// /usr/include/qt/QtGui/qpicture.h:149
// index:0
// Public
// void setPicture(const class QPicture &)
func (this *QPictureIO) SetPicture(arg0 *QPicture) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO10setPictureERK8QPicture", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:150
// index:0
// Public
// void setStatus(int)
func (this *QPictureIO) SetStatus(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO9setStatusEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:151
// index:0
// Public
// void setFormat(const char *)
func (this *QPictureIO) SetFormat(arg0 string) {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO9setFormatEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:152
// index:0
// Public
// void setIODevice(class QIODevice *)
func (this *QPictureIO) SetIODevice(arg0 *qtcore.QIODevice /*444 QIODevice **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO11setIODeviceEP9QIODevice", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:153
// index:0
// Public
// void setFileName(const class QString &)
func (this *QPictureIO) SetFileName(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO11setFileNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:154
// index:0
// Public
// void setQuality(int)
func (this *QPictureIO) SetQuality(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO10setQualityEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:155
// index:0
// Public
// void setDescription(const class QString &)
func (this *QPictureIO) SetDescription(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO14setDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:156
// index:0
// Public
// void setParameters(const char *)
func (this *QPictureIO) SetParameters(arg0 string) {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13setParametersEPKc", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:157
// index:0
// Public
// void setGamma(float)
func (this *QPictureIO) SetGamma(arg0 float32) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO8setGammaEf", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:159
// index:0
// Public
// bool read()
func (this *QPictureIO) Read() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO4readEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:160
// index:0
// Public
// bool write()
func (this *QPictureIO) Write() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO5writeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpicture.h:162
// index:0
// Public static
// QByteArray pictureFormat(const class QString &)
func (this *QPictureIO) PictureFormat(fileName *qtcore.QString) *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13pictureFormatERK7QString", ffiqt.FFI_TYPE_POINTER, fileName)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPictureIO_PictureFormat(fileName *qtcore.QString) *qtcore.QByteArray /*123*/ {
	var nilthis *QPictureIO
	rv := nilthis.PictureFormat(fileName)
	return rv
}

// /usr/include/qt/QtGui/qpicture.h:163
// index:1
// Public static
// QByteArray pictureFormat(class QIODevice *)
func (this *QPictureIO) PictureFormat_1(arg0 *qtcore.QIODevice /*444 QIODevice **/) *qtcore.QByteArray /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPictureIO13pictureFormatEP9QIODevice", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPictureIO_PictureFormat_1(arg0 *qtcore.QIODevice /*444 QIODevice **/) *qtcore.QByteArray /*123*/ {
	var nilthis *QPictureIO
	rv := nilthis.PictureFormat_1(arg0)
	return rv
}

//  body block end
