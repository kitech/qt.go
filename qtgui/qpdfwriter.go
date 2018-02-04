package qtgui

// /usr/include/qt/QtGui/qpdfwriter.h
// #include <qpdfwriter.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 19
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
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
// QPaintEngine * paintEngine()
func (this *QPdfWriter) InheritPaintEngine(f func() unsafe.Pointer /*666*/) {
	ffiqt.SetAllInheritCallback(this, "paintEngine", f)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPdfWriter) InheritMetric(f func(id int) int) {
	ffiqt.SetAllInheritCallback(this, "metric", f)
}

type QPdfWriter struct {
	*qtcore.QObject
	*QPagedPaintDevice
}

func (this *QPdfWriter) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QPdfWriter) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QPagedPaintDevice = NewQPagedPaintDeviceFromPointer(cthis)
}
func NewQPdfWriterFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQPagedPaintDeviceFromPointer(cthis)
	return &QPdfWriter{bcthis0, bcthis1}
}
func (*QPdfWriter) NewFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	return NewQPdfWriterFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpdfwriter.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QPdfWriter) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPdfWriter(const QString &)
func NewQPdfWriter(filename *qtcore.QString) *QPdfWriter {
	var convArg0 = filename.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QPdfWriter(QIODevice *)
func NewQPdfWriter_1(device *qtcore.QIODevice /*777 QIODevice **/) *QPdfWriter {
	var convArg0 = device.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2EP9QIODevice", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:62
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPdfWriter()
func DeleteQPdfWriter(this *QPdfWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpdfwriter.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPdfVersion(enum QPagedPaintDevice::PdfVersion)
func (this *QPdfWriter) SetPdfVersion(version int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPdfVersionEN17QPagedPaintDevice10PdfVersionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:65
// index:0
// Public Visibility=Default Availability=Available
// [4] QPagedPaintDevice::PdfVersion pdfVersion()
func (this *QPdfWriter) PdfVersion() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10pdfVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:67
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title()
func (this *QPdfWriter) Title() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter5titleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QPdfWriter) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:70
// index:0
// Public Visibility=Default Availability=Available
// [8] QString creator()
func (this *QPdfWriter) Creator() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter7creatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCreator(const QString &)
func (this *QPdfWriter) SetCreator(creator *qtcore.QString) {
	var convArg0 = creator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter10setCreatorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:73
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool newPage()
func (this *QPdfWriter) NewPage() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter7newPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpdfwriter.h:75
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setResolution(int)
func (this *QPdfWriter) SetResolution(resolution int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setResolutionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:76
// index:0
// Public Visibility=Default Availability=Available
// [4] int resolution()
func (this *QPdfWriter) Resolution() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10resolutionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpdfwriter.h:89
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPdfWriter) SetPageSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter11setPageSizeEN17QPagedPaintDevice8PageSizeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:90
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSizeMM(const QSizeF &)
func (this *QPdfWriter) SetPageSizeMM(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPageSizeMMERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:95
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine()
func (this *QPdfWriter) PaintEngine() *QPaintEngine /*777 QPaintEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qpdfwriter.h:96
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPdfWriter) Metric(id int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

//  body block end
