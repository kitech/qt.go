package qtsvg

// /usr/include/qt/QtSvg/qsvggenerator.h
// #include <qsvggenerator.h>
// #include <QtSvg>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 14
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"

//  ext block end

//  body block begin

// QPaintEngine * paintEngine()
func (this *QSvgGenerator) InheritPaintEngine(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "paintEngine", f)
}

// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QSvgGenerator) InheritMetric(f func(metric int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

/*

 */
type QSvgGenerator struct {
	*qtgui.QPaintDevice
}
type QSvgGenerator_ITF interface {
	qtgui.QPaintDevice_ITF
	QSvgGenerator_PTR() *QSvgGenerator
}

func (ptr *QSvgGenerator) QSvgGenerator_PTR() *QSvgGenerator { return ptr }

func (this *QSvgGenerator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QSvgGenerator) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = qtgui.NewQPaintDeviceFromPointer(cthis)
}
func NewQSvgGeneratorFromPointer(cthis unsafe.Pointer) *QSvgGenerator {
	bcthis0 := qtgui.NewQPaintDeviceFromPointer(cthis)
	return &QSvgGenerator{bcthis0}
}
func (*QSvgGenerator) NewFromPointer(cthis unsafe.Pointer) *QSvgGenerator {
	return NewQSvgGeneratorFromPointer(cthis)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSvgGenerator()

/*
Constructs a new generator.
*/
func (*QSvgGenerator) NewForInherit() *QSvgGenerator {
	return NewQSvgGenerator()
}
func NewQSvgGenerator() *QSvgGenerator {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGeneratorC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSvgGeneratorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSvgGenerator)
	return gothis
}

// /usr/include/qt/QtSvg/qsvggenerator.h:71
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSvgGenerator()

/*

 */
func DeleteQSvgGenerator(this *QSvgGenerator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGeneratorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:73
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title() const

/*

 */
func (this *QSvgGenerator) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtSvg/qsvggenerator.h:74
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)

/*

 */
func (this *QSvgGenerator) SetTitle(title string) {
	var tmpArg0 = qtcore.NewQString5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:76
// index:0
// Public Visibility=Default Availability=Available
// [8] QString description() const

/*

 */
func (this *QSvgGenerator) Description() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator11descriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtSvg/qsvggenerator.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDescription(const QString &)

/*

 */
func (this *QSvgGenerator) SetDescription(description string) {
	var tmpArg0 = qtcore.NewQString5(description)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator14setDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:79
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*

 */
func (this *QSvgGenerator) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtSvg/qsvggenerator.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSize(const QSize &)

/*

 */
func (this *QSvgGenerator) SetSize(size qtcore.QSize_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSize_PTR() != nil {
		convArg0 = size.QSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator7setSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:82
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect viewBox() const

/*
Returns viewBoxF().toRect().

This function was introduced in  Qt 4.5.

See also setViewBox() and viewBoxF().
*/
func (this *QSvgGenerator) ViewBox() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator7viewBoxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtSvg/qsvggenerator.h:83
// index:0
// Public Visibility=Default Availability=Available
// [32] QRectF viewBoxF() const

/*

 */
func (this *QSvgGenerator) ViewBoxF() *qtcore.QRectF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator8viewBoxFEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRectF)
	return rv2
}

// /usr/include/qt/QtSvg/qsvggenerator.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setViewBox(const QRect &)

/*

 */
func (this *QSvgGenerator) SetViewBox(viewBox qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if viewBox != nil && viewBox.QRect_PTR() != nil {
		convArg0 = viewBox.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator10setViewBoxERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:85
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setViewBox(const QRectF &)

/*

 */
func (this *QSvgGenerator) SetViewBox1(viewBox qtcore.QRectF_ITF) {
	var convArg0 unsafe.Pointer
	if viewBox != nil && viewBox.QRectF_PTR() != nil {
		convArg0 = viewBox.QRectF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator10setViewBoxERK6QRectF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*

 */
func (this *QSvgGenerator) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtSvg/qsvggenerator.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*

 */
func (this *QSvgGenerator) SetFileName(fileName string) {
	var tmpArg0 = qtcore.NewQString5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QIODevice * outputDevice() const

/*

 */
func (this *QSvgGenerator) OutputDevice() *qtcore.QIODevice /*777 QIODevice **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator12outputDeviceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQIODeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qsvggenerator.h:91
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOutputDevice(QIODevice *)

/*

 */
func (this *QSvgGenerator) SetOutputDevice(outputDevice qtcore.QIODevice_ITF /*777 QIODevice **/) {
	var convArg0 unsafe.Pointer
	if outputDevice != nil && outputDevice.QIODevice_PTR() != nil {
		convArg0 = outputDevice.QIODevice_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator15setOutputDeviceEP9QIODevice", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(int)

/*

 */
func (this *QSvgGenerator) SetResolution(dpi int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN13QSvgGenerator13setResolutionEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dpi)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtSvg/qsvggenerator.h:94
// index:0
// Public Visibility=Default Availability=Available
// [4] int resolution() const

/*

 */
func (this *QSvgGenerator) Resolution() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator10resolutionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtSvg/qsvggenerator.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine() const

/*
Reimplemented from QPaintDevice::paintEngine().

Returns the paint engine used to render graphics to be converted to SVG format information.
*/
func (this *QSvgGenerator) PaintEngine() *qtgui.QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtSvg/qsvggenerator.h:97
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(QPaintDevice::PaintDeviceMetric) const

/*
Reimplemented from QPaintDevice::metric().
*/
func (this *QSvgGenerator) Metric(metric int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK13QSvgGenerator6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), metric)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end

//  keep block begin

func init_unused_11921() {
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
		qtwidgets.KeepMe()
	}
}

//  keep block end
