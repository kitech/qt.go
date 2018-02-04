package qtgui

// /usr/include/qt/QtGui/qpagedpaintdevice.h
// #include <qpagedpaintdevice.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 34
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
// QPageLayout devicePageLayout()
func (this *QPagedPaintDevice) InheritDevicePageLayout(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "devicePageLayout", f)
}

type QPagedPaintDevice struct {
	*QPaintDevice
}

func (this *QPagedPaintDevice) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QPaintDevice.GetCthis()
	}
}
func (this *QPagedPaintDevice) SetCthis(cthis unsafe.Pointer) {
	this.QPaintDevice = NewQPaintDeviceFromPointer(cthis)
}
func NewQPagedPaintDeviceFromPointer(cthis unsafe.Pointer) *QPagedPaintDevice {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPagedPaintDevice{bcthis0}
}
func (*QPagedPaintDevice) NewFromPointer(cthis unsafe.Pointer) *QPagedPaintDevice {
	return NewQPagedPaintDeviceFromPointer(cthis)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QPagedPaintDevice()
func NewQPagedPaintDevice() *QPagedPaintDevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDeviceC1Ev", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQPagedPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPagedPaintDevice)
	return gothis
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPagedPaintDevice()
func DeleteQPagedPaintDevice(this *QPagedPaintDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool newPage()
func (this *QPagedPaintDevice) NewPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice7newPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:219
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageLayout(const QPageLayout &)
func (this *QPagedPaintDevice) SetPageLayout(pageLayout *QPageLayout) bool {
	var convArg0 = pageLayout.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:220
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageSize(const QPageSize &)
func (this *QPagedPaintDevice) SetPageSize(pageSize *QPageSize) bool {
	var convArg0 = pageSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeERK9QPageSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:226
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPagedPaintDevice) SetPageSize_1(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeENS_8PageSizeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:221
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageOrientation(QPageLayout::Orientation)
func (this *QPagedPaintDevice) SetPageOrientation(orientation int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice18setPageOrientationEN11QPageLayout11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:222
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageMargins(const QMarginsF &)
func (this *QPagedPaintDevice) SetPageMargins(margins *qtcore.QMarginsF) bool {
	var convArg0 = margins.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:223
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setPageMargins(const QMarginsF &, QPageLayout::Unit)
func (this *QPagedPaintDevice) SetPageMargins_1(margins *qtcore.QMarginsF, units int) bool {
	var convArg0 = margins.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsFN11QPageLayout4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, units)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QPageLayout pageLayout()
func (this *QPagedPaintDevice) PageLayout() *QPageLayout /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:227
// index:0
// Public Visibility=Default Availability=Available
// [4] QPagedPaintDevice::PageSize pageSize()
func (this *QPagedPaintDevice) PageSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice8pageSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:229
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSizeMM(const QSizeF &)
func (this *QPagedPaintDevice) SetPageSizeMM(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:230
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF pageSizeMM()
func (this *QPagedPaintDevice) PageSizeMM() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageSizeMMEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:241
// index:0
// Public Visibility=Default Availability=Available
// [32] QPagedPaintDevice::Margins margins()
func (this *QPagedPaintDevice) Margins() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice7marginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPageLayout devicePageLayout()
func (this *QPagedPaintDevice) DevicePageLayout() *QPageLayout /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice16devicePageLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:247
// index:1
// Protected Visibility=Default Availability=Available
// [8] QPageLayout & devicePageLayout()
func (this *QPagedPaintDevice) DevicePageLayout_1() *QPageLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice16devicePageLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

type QPagedPaintDevice__PageSize = int

const QPagedPaintDevice__A4 QPagedPaintDevice__PageSize = 0
const QPagedPaintDevice__B5 QPagedPaintDevice__PageSize = 1
const QPagedPaintDevice__Letter QPagedPaintDevice__PageSize = 2
const QPagedPaintDevice__Legal QPagedPaintDevice__PageSize = 3
const QPagedPaintDevice__Executive QPagedPaintDevice__PageSize = 4
const QPagedPaintDevice__A0 QPagedPaintDevice__PageSize = 5
const QPagedPaintDevice__A1 QPagedPaintDevice__PageSize = 6
const QPagedPaintDevice__A2 QPagedPaintDevice__PageSize = 7
const QPagedPaintDevice__A3 QPagedPaintDevice__PageSize = 8
const QPagedPaintDevice__A5 QPagedPaintDevice__PageSize = 9
const QPagedPaintDevice__A6 QPagedPaintDevice__PageSize = 10
const QPagedPaintDevice__A7 QPagedPaintDevice__PageSize = 11
const QPagedPaintDevice__A8 QPagedPaintDevice__PageSize = 12
const QPagedPaintDevice__A9 QPagedPaintDevice__PageSize = 13
const QPagedPaintDevice__B0 QPagedPaintDevice__PageSize = 14
const QPagedPaintDevice__B1 QPagedPaintDevice__PageSize = 15
const QPagedPaintDevice__B10 QPagedPaintDevice__PageSize = 16
const QPagedPaintDevice__B2 QPagedPaintDevice__PageSize = 17
const QPagedPaintDevice__B3 QPagedPaintDevice__PageSize = 18
const QPagedPaintDevice__B4 QPagedPaintDevice__PageSize = 19
const QPagedPaintDevice__B6 QPagedPaintDevice__PageSize = 20
const QPagedPaintDevice__B7 QPagedPaintDevice__PageSize = 21
const QPagedPaintDevice__B8 QPagedPaintDevice__PageSize = 22
const QPagedPaintDevice__B9 QPagedPaintDevice__PageSize = 23
const QPagedPaintDevice__C5E QPagedPaintDevice__PageSize = 24
const QPagedPaintDevice__Comm10E QPagedPaintDevice__PageSize = 25
const QPagedPaintDevice__DLE QPagedPaintDevice__PageSize = 26
const QPagedPaintDevice__Folio QPagedPaintDevice__PageSize = 27
const QPagedPaintDevice__Ledger QPagedPaintDevice__PageSize = 28
const QPagedPaintDevice__Tabloid QPagedPaintDevice__PageSize = 29
const QPagedPaintDevice__Custom QPagedPaintDevice__PageSize = 30
const QPagedPaintDevice__A10 QPagedPaintDevice__PageSize = 31
const QPagedPaintDevice__A3Extra QPagedPaintDevice__PageSize = 32
const QPagedPaintDevice__A4Extra QPagedPaintDevice__PageSize = 33
const QPagedPaintDevice__A4Plus QPagedPaintDevice__PageSize = 34
const QPagedPaintDevice__A4Small QPagedPaintDevice__PageSize = 35
const QPagedPaintDevice__A5Extra QPagedPaintDevice__PageSize = 36
const QPagedPaintDevice__B5Extra QPagedPaintDevice__PageSize = 37
const QPagedPaintDevice__JisB0 QPagedPaintDevice__PageSize = 38
const QPagedPaintDevice__JisB1 QPagedPaintDevice__PageSize = 39
const QPagedPaintDevice__JisB2 QPagedPaintDevice__PageSize = 40
const QPagedPaintDevice__JisB3 QPagedPaintDevice__PageSize = 41
const QPagedPaintDevice__JisB4 QPagedPaintDevice__PageSize = 42
const QPagedPaintDevice__JisB5 QPagedPaintDevice__PageSize = 43
const QPagedPaintDevice__JisB6 QPagedPaintDevice__PageSize = 44
const QPagedPaintDevice__JisB7 QPagedPaintDevice__PageSize = 45
const QPagedPaintDevice__JisB8 QPagedPaintDevice__PageSize = 46
const QPagedPaintDevice__JisB9 QPagedPaintDevice__PageSize = 47
const QPagedPaintDevice__JisB10 QPagedPaintDevice__PageSize = 48
const QPagedPaintDevice__AnsiC QPagedPaintDevice__PageSize = 49
const QPagedPaintDevice__AnsiD QPagedPaintDevice__PageSize = 50
const QPagedPaintDevice__AnsiE QPagedPaintDevice__PageSize = 51
const QPagedPaintDevice__LegalExtra QPagedPaintDevice__PageSize = 52
const QPagedPaintDevice__LetterExtra QPagedPaintDevice__PageSize = 53
const QPagedPaintDevice__LetterPlus QPagedPaintDevice__PageSize = 54
const QPagedPaintDevice__LetterSmall QPagedPaintDevice__PageSize = 55
const QPagedPaintDevice__TabloidExtra QPagedPaintDevice__PageSize = 56
const QPagedPaintDevice__ArchA QPagedPaintDevice__PageSize = 57
const QPagedPaintDevice__ArchB QPagedPaintDevice__PageSize = 58
const QPagedPaintDevice__ArchC QPagedPaintDevice__PageSize = 59
const QPagedPaintDevice__ArchD QPagedPaintDevice__PageSize = 60
const QPagedPaintDevice__ArchE QPagedPaintDevice__PageSize = 61
const QPagedPaintDevice__Imperial7x9 QPagedPaintDevice__PageSize = 62
const QPagedPaintDevice__Imperial8x10 QPagedPaintDevice__PageSize = 63
const QPagedPaintDevice__Imperial9x11 QPagedPaintDevice__PageSize = 64
const QPagedPaintDevice__Imperial9x12 QPagedPaintDevice__PageSize = 65
const QPagedPaintDevice__Imperial10x11 QPagedPaintDevice__PageSize = 66
const QPagedPaintDevice__Imperial10x13 QPagedPaintDevice__PageSize = 67
const QPagedPaintDevice__Imperial10x14 QPagedPaintDevice__PageSize = 68
const QPagedPaintDevice__Imperial12x11 QPagedPaintDevice__PageSize = 69
const QPagedPaintDevice__Imperial15x11 QPagedPaintDevice__PageSize = 70
const QPagedPaintDevice__ExecutiveStandard QPagedPaintDevice__PageSize = 71
const QPagedPaintDevice__Note QPagedPaintDevice__PageSize = 72
const QPagedPaintDevice__Quarto QPagedPaintDevice__PageSize = 73
const QPagedPaintDevice__Statement QPagedPaintDevice__PageSize = 74
const QPagedPaintDevice__SuperA QPagedPaintDevice__PageSize = 75
const QPagedPaintDevice__SuperB QPagedPaintDevice__PageSize = 76
const QPagedPaintDevice__Postcard QPagedPaintDevice__PageSize = 77
const QPagedPaintDevice__DoublePostcard QPagedPaintDevice__PageSize = 78
const QPagedPaintDevice__Prc16K QPagedPaintDevice__PageSize = 79
const QPagedPaintDevice__Prc32K QPagedPaintDevice__PageSize = 80
const QPagedPaintDevice__Prc32KBig QPagedPaintDevice__PageSize = 81
const QPagedPaintDevice__FanFoldUS QPagedPaintDevice__PageSize = 82
const QPagedPaintDevice__FanFoldGerman QPagedPaintDevice__PageSize = 83
const QPagedPaintDevice__FanFoldGermanLegal QPagedPaintDevice__PageSize = 84
const QPagedPaintDevice__EnvelopeB4 QPagedPaintDevice__PageSize = 85
const QPagedPaintDevice__EnvelopeB5 QPagedPaintDevice__PageSize = 86
const QPagedPaintDevice__EnvelopeB6 QPagedPaintDevice__PageSize = 87
const QPagedPaintDevice__EnvelopeC0 QPagedPaintDevice__PageSize = 88
const QPagedPaintDevice__EnvelopeC1 QPagedPaintDevice__PageSize = 89
const QPagedPaintDevice__EnvelopeC2 QPagedPaintDevice__PageSize = 90
const QPagedPaintDevice__EnvelopeC3 QPagedPaintDevice__PageSize = 91
const QPagedPaintDevice__EnvelopeC4 QPagedPaintDevice__PageSize = 92
const QPagedPaintDevice__EnvelopeC6 QPagedPaintDevice__PageSize = 93
const QPagedPaintDevice__EnvelopeC65 QPagedPaintDevice__PageSize = 94
const QPagedPaintDevice__EnvelopeC7 QPagedPaintDevice__PageSize = 95
const QPagedPaintDevice__Envelope9 QPagedPaintDevice__PageSize = 96
const QPagedPaintDevice__Envelope11 QPagedPaintDevice__PageSize = 97
const QPagedPaintDevice__Envelope12 QPagedPaintDevice__PageSize = 98
const QPagedPaintDevice__Envelope14 QPagedPaintDevice__PageSize = 99
const QPagedPaintDevice__EnvelopeMonarch QPagedPaintDevice__PageSize = 100
const QPagedPaintDevice__EnvelopePersonal QPagedPaintDevice__PageSize = 101
const QPagedPaintDevice__EnvelopeChou3 QPagedPaintDevice__PageSize = 102
const QPagedPaintDevice__EnvelopeChou4 QPagedPaintDevice__PageSize = 103
const QPagedPaintDevice__EnvelopeInvite QPagedPaintDevice__PageSize = 104
const QPagedPaintDevice__EnvelopeItalian QPagedPaintDevice__PageSize = 105
const QPagedPaintDevice__EnvelopeKaku2 QPagedPaintDevice__PageSize = 106
const QPagedPaintDevice__EnvelopeKaku3 QPagedPaintDevice__PageSize = 107
const QPagedPaintDevice__EnvelopePrc1 QPagedPaintDevice__PageSize = 108
const QPagedPaintDevice__EnvelopePrc2 QPagedPaintDevice__PageSize = 109
const QPagedPaintDevice__EnvelopePrc3 QPagedPaintDevice__PageSize = 110
const QPagedPaintDevice__EnvelopePrc4 QPagedPaintDevice__PageSize = 111
const QPagedPaintDevice__EnvelopePrc5 QPagedPaintDevice__PageSize = 112
const QPagedPaintDevice__EnvelopePrc6 QPagedPaintDevice__PageSize = 113
const QPagedPaintDevice__EnvelopePrc7 QPagedPaintDevice__PageSize = 114
const QPagedPaintDevice__EnvelopePrc8 QPagedPaintDevice__PageSize = 115
const QPagedPaintDevice__EnvelopePrc9 QPagedPaintDevice__PageSize = 116
const QPagedPaintDevice__EnvelopePrc10 QPagedPaintDevice__PageSize = 117
const QPagedPaintDevice__EnvelopeYou4 QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__LastPageSize QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__NPageSize QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__NPaperSize QPagedPaintDevice__PageSize = 118
const QPagedPaintDevice__AnsiA QPagedPaintDevice__PageSize = 2
const QPagedPaintDevice__AnsiB QPagedPaintDevice__PageSize = 28
const QPagedPaintDevice__EnvelopeC5 QPagedPaintDevice__PageSize = 24
const QPagedPaintDevice__EnvelopeDL QPagedPaintDevice__PageSize = 26
const QPagedPaintDevice__Envelope10 QPagedPaintDevice__PageSize = 25

type QPagedPaintDevice__PdfVersion = int

const QPagedPaintDevice__PdfVersion_1_4 QPagedPaintDevice__PdfVersion = 0
const QPagedPaintDevice__PdfVersion_A1b QPagedPaintDevice__PdfVersion = 1

//  body block end
