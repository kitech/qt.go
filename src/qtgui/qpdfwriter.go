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

// /usr/include/qt/QtGui/qpdfwriter.h:58
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QPdfWriter) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:60
// index:0
// void QPdfWriter(const class QString &)
func NewQPdfWriter(filename unsafe.Pointer) *QPdfWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, filename)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(cthis)
	return gothis
}
func NewQPdfWriterFromPointer(cthis unsafe.Pointer) *QPdfWriter {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQPagedPaintDeviceFromPointer(cthis)
	return &QPdfWriter{bcthis0, bcthis1}
}

// /usr/include/qt/QtGui/qpdfwriter.h:61
// index:1
// void QPdfWriter(class QIODevice *)
func NewQPdfWriter_1(device unsafe.Pointer) *QPdfWriter {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterC2EP9QIODevice", ffiqt.FFI_TYPE_VOID, cthis, device)
	gopp.ErrPrint(err, rv)
	gothis := NewQPdfWriterFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpdfwriter.h:62
// index:0
// virtual
// void ~QPdfWriter()
func DeleteQPdfWriter(*QPdfWriter) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriterD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:64
// index:0
// void setPdfVersion(enum QPagedPaintDevice::PdfVersion)
func (this *QPdfWriter) SetPdfVersion(version int) {
	// 0: (, version QPagedPaintDevice::PdfVersion), (&version)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPdfVersionEN17QPagedPaintDevice10PdfVersionE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &version)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:65
// index:0
// QPagedPaintDevice::PdfVersion pdfVersion()
func (this *QPdfWriter) PdfVersion() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10pdfVersionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:67
// index:0
// QString title()
func (this *QPdfWriter) Title() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter5titleEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:68
// index:0
// void setTitle(const class QString &)
func (this *QPdfWriter) SetTitle(title unsafe.Pointer) {
	// 0: (, title const QString &), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter8setTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:70
// index:0
// QString creator()
func (this *QPdfWriter) Creator() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter7creatorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:71
// index:0
// void setCreator(const class QString &)
func (this *QPdfWriter) SetCreator(creator unsafe.Pointer) {
	// 0: (, creator const QString &), (creator)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter10setCreatorERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), creator)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:73
// index:0
// virtual
// bool newPage()
func (this *QPdfWriter) NewPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter7newPageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:75
// index:0
// void setResolution(int)
func (this *QPdfWriter) SetResolution(resolution int) {
	// 0: (, resolution int), (&resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setResolutionEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:76
// index:0
// int resolution()
func (this *QPdfWriter) Resolution() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter10resolutionEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:89
// index:0
// virtual
// void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPdfWriter) SetPageSize(size int) {
	// 0: (, size QPagedPaintDevice::PageSize), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter11setPageSizeEN17QPagedPaintDevice8PageSizeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:90
// index:0
// virtual
// void setPageSizeMM(const class QSizeF &)
func (this *QPdfWriter) SetPageSizeMM(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter13setPageSizeMMERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:92
// index:0
// virtual
// void setMargins(const struct QPagedPaintDevice::Margins &)
func (this *QPdfWriter) SetMargins(m unsafe.Pointer) {
	// 0: (, m const QPagedPaintDevice::Margins &), (m)
	rv, err := ffiqt.InvokeQtFunc6("_ZN10QPdfWriter10setMarginsERKN17QPagedPaintDevice7MarginsE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), m)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:95
// index:0
// virtual
// QPaintEngine * paintEngine()
func (this *QPdfWriter) PaintEngine() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter11paintEngineEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpdfwriter.h:96
// index:0
// virtual
// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QPdfWriter) Metric(id int) {
	// 0: (, id QPaintDevice::PaintDeviceMetric), (&id)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK10QPdfWriter6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &id)
	gopp.ErrPrint(err, rv)
}

//  body block end
