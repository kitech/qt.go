//  header block begin
// /usr/include/qt/QtGui/qpdfwriter.h
// #include <qpdfwriter.h>
// #include <QtGui>
package qtgui

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
type QPdfWriter struct {
	*qtcore.QObject
	*QPagedPaintDevice
}

func (this *QPdfWriter) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}
func NewQPdfWriterFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQPagedPaintDeviceFromPointer(cthis)
	return &QPdfWriter{bcthis0, bcthis1}
}

// /usr/include/qt/QtGui/qpdfwriter.h:58
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QPdfWriter) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:60
// index:0
// Public
// void QPdfWriter(const class QString &)
func NewQPdfWriter(filename *qtcore.QString) *QPdfWriter {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = filename.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:61
// index:1
// Public
// void QPdfWriter(class QIODevice *)
func NewQPdfWriter_1(device unsafe.Pointer) *QPdfWriter {
	cthis := qtrt.Calloc(1, 256) // 48
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, device)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:62
// index:0
// Public virtual
// void ~QPdfWriter()
func DeleteQPdfWriter(*QPdfWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:64
// index:0
// Public
// void setPdfVersion(enum QPagedPaintDevice::PdfVersion)
func (this *QPdfWriter) SetPdfVersion(version int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPdfVersionEN17QPagedPaintDevice10PdfVersionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:65
// index:0
// Public
// QPagedPaintDevice::PdfVersion pdfVersion()
func (this *QPdfWriter) PdfVersion() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10pdfVersionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:67
// index:0
// Public
// QString title()
func (this *QPdfWriter) Title() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter5titleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:68
// index:0
// Public
// void setTitle(const class QString &)
func (this *QPdfWriter) SetTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter8setTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:70
// index:0
// Public
// QString creator()
func (this *QPdfWriter) Creator() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter7creatorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:71
// index:0
// Public
// void setCreator(const class QString &)
func (this *QPdfWriter) SetCreator(creator *qtcore.QString) {
	var convArg0 = creator.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter10setCreatorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:73
// index:0
// Public virtual
// bool newPage()
func (this *QPdfWriter) NewPage() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter7newPageEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:75
// index:0
// Public
// void setResolution(int)
func (this *QPdfWriter) SetResolution(resolution int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setResolutionEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:76
// index:0
// Public
// int resolution()
func (this *QPdfWriter) Resolution() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10resolutionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:89
// index:0
// Public virtual
// void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPdfWriter) SetPageSize(size int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter11setPageSizeEN17QPagedPaintDevice8PageSizeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:90
// index:0
// Public virtual
// void setPageSizeMM(const class QSizeF &)
func (this *QPdfWriter) SetPageSizeMM(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPageSizeMMERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:95
// index:0
// Protected virtual
// QPaintEngine * paintEngine()
func (this *QPdfWriter) PaintEngine() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtGui/qpdfwriter.h:96
// index:0
// Protected virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPdfWriter) Metric(id int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
