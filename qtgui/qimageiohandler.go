package qtgui

// /usr/include/qt/QtGui/qimageiohandler.h
// #include <qimageiohandler.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 4
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QImageIOHandler struct {
	*qtrt.CObject
}

func (this *QImageIOHandler) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QImageIOHandler) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQImageIOHandlerFromPointer(cthis unsafe.Pointer) *QImageIOHandler {
	return &QImageIOHandler{&qtrt.CObject{cthis}}
}
func (*QImageIOHandler) NewFromPointer(cthis unsafe.Pointer) *QImageIOHandler {
	return NewQImageIOHandlerFromPointer(cthis)
}

// /usr/include/qt/QtGui/qimageiohandler.h:62
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QImageIOHandler()
func NewQImageIOHandler() *QImageIOHandler {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandlerC2Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQImageIOHandlerFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQImageIOHandler)
	return gothis
}

// /usr/include/qt/QtGui/qimageiohandler.h:63
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QImageIOHandler()
func DeleteQImageIOHandler(this *QImageIOHandler) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandlerD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qimageiohandler.h:65
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDevice(QIODevice *)
func (this *QImageIOHandler) SetDevice(device *qtcore.QIODevice /*777 QIODevice **/) {
	var convArg0 = device.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler9setDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:66
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * device()
func (this *QImageIOHandler) Device() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler6deviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)
func (this *QImageIOHandler) SetFormat(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QByteArray &)
func (this *QImageIOHandler) SetFormat_1(format *qtcore.QByteArray) {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler9setFormatERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray format()
func (this *QImageIOHandler) Format() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QByteArray name()
func (this *QImageIOHandler) Name() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:74
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool canRead()
func (this *QImageIOHandler) CanRead() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler7canReadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:75
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool read(QImage *)
func (this *QImageIOHandler) Read(image *QImage /*777 QImage **/) bool {
	var convArg0 = image.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler4readEP6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:76
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool write(const QImage &)
func (this *QImageIOHandler) Write(image *QImage) bool {
	var convArg0 = image.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler5writeERK6QImage", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:115
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant option(enum QImageIOHandler::ImageOption)
func (this *QImageIOHandler) Option(option int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler6optionENS_11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:116
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setOption(enum QImageIOHandler::ImageOption, const QVariant &)
func (this *QImageIOHandler) SetOption(option int, value *qtcore.QVariant) {
	var convArg1 = value.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler9setOptionENS_11ImageOptionERK8QVariant", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimageiohandler.h:117
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool supportsOption(enum QImageIOHandler::ImageOption)
func (this *QImageIOHandler) SupportsOption(option int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler14supportsOptionENS_11ImageOptionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), option)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool jumpToNextImage()
func (this *QImageIOHandler) JumpToNextImage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler15jumpToNextImageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:121
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool jumpToImage(int)
func (this *QImageIOHandler) JumpToImage(imageNumber int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QImageIOHandler11jumpToImageEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), imageNumber)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qimageiohandler.h:122
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int loopCount()
func (this *QImageIOHandler) LoopCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler9loopCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:123
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int imageCount()
func (this *QImageIOHandler) ImageCount() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler10imageCountEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int nextImageDelay()
func (this *QImageIOHandler) NextImageDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler14nextImageDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:125
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int currentImageNumber()
func (this *QImageIOHandler) CurrentImageNumber() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler18currentImageNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qimageiohandler.h:126
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QRect currentImageRect()
func (this *QImageIOHandler) CurrentImageRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK15QImageIOHandler16currentImageRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

type QImageIOHandler__ImageOption = int

const QImageIOHandler__Size QImageIOHandler__ImageOption = 0
const QImageIOHandler__ClipRect QImageIOHandler__ImageOption = 1
const QImageIOHandler__Description QImageIOHandler__ImageOption = 2
const QImageIOHandler__ScaledClipRect QImageIOHandler__ImageOption = 3
const QImageIOHandler__ScaledSize QImageIOHandler__ImageOption = 4
const QImageIOHandler__CompressionRatio QImageIOHandler__ImageOption = 5
const QImageIOHandler__Gamma QImageIOHandler__ImageOption = 6
const QImageIOHandler__Quality QImageIOHandler__ImageOption = 7
const QImageIOHandler__Name QImageIOHandler__ImageOption = 8
const QImageIOHandler__SubType QImageIOHandler__ImageOption = 9
const QImageIOHandler__IncrementalReading QImageIOHandler__ImageOption = 10
const QImageIOHandler__Endianness QImageIOHandler__ImageOption = 11
const QImageIOHandler__Animation QImageIOHandler__ImageOption = 12
const QImageIOHandler__BackgroundColor QImageIOHandler__ImageOption = 13
const QImageIOHandler__ImageFormat QImageIOHandler__ImageOption = 14
const QImageIOHandler__SupportedSubTypes QImageIOHandler__ImageOption = 15
const QImageIOHandler__OptimizedWrite QImageIOHandler__ImageOption = 16
const QImageIOHandler__ProgressiveScanWrite QImageIOHandler__ImageOption = 17
const QImageIOHandler__ImageTransformation QImageIOHandler__ImageOption = 18
const QImageIOHandler__TransformedByDefault QImageIOHandler__ImageOption = 19

type QImageIOHandler__Transformation = int

const QImageIOHandler__TransformationNone QImageIOHandler__Transformation = 0
const QImageIOHandler__TransformationMirror QImageIOHandler__Transformation = 1
const QImageIOHandler__TransformationFlip QImageIOHandler__Transformation = 2
const QImageIOHandler__TransformationRotate180 QImageIOHandler__Transformation = 3
const QImageIOHandler__TransformationRotate90 QImageIOHandler__Transformation = 4
const QImageIOHandler__TransformationMirrorAndRotate90 QImageIOHandler__Transformation = 5
const QImageIOHandler__TransformationFlipAndRotate90 QImageIOHandler__Transformation = 6
const QImageIOHandler__TransformationRotate270 QImageIOHandler__Transformation = 7

//  body block end
