package qtgui

// /usr/include/qt/QtGui/qpagedpaintdevice.h
// #include <qpagedpaintdevice.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 36
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

// QPageLayout devicePageLayout()
func (this *QPagedPaintDevice) InheritDevicePageLayout(f func() unsafe.Pointer) {
	qtrt.SetAllInheritCallback(this, "devicePageLayout", f)
}

/*

 */
type QPagedPaintDevice struct {
	*QPaintDevice
}
type QPagedPaintDevice_ITF interface {
	QPaintDevice_ITF
	QPagedPaintDevice_PTR() *QPagedPaintDevice
}

func (ptr *QPagedPaintDevice) QPagedPaintDevice_PTR() *QPagedPaintDevice { return ptr }

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

/*
Constructs a new paged paint device.
*/
func NewQPagedPaintDevice() *QPagedPaintDevice {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDeviceC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQPagedPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQPagedPaintDevice)
	return gothis
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QPagedPaintDevice()

/*

 */
func DeleteQPagedPaintDevice(this *QPagedPaintDevice) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDeviceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:61
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [1] bool newPage()

/*
Starts a new page. Returns true on success.
*/
func (this *QPagedPaintDevice) NewPage() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice7newPageEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:219
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageLayout(const QPageLayout &)

/*
Sets the page layout to newPageLayout.

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new page layout to a new page. You should not call any painting methods between a call to setPageLayout() and newPage() as the wrong paint metrics may be used.

Returns true if the page layout was successfully set to newPageLayout.

This function was introduced in  Qt 5.3.

See also pageLayout().
*/
func (this *QPagedPaintDevice) SetPageLayout(pageLayout QPageLayout_ITF) bool {
	var convArg0 unsafe.Pointer
	if pageLayout != nil && pageLayout.QPageLayout_PTR() != nil {
		convArg0 = pageLayout.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:220
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageSize(const QPageSize &)

/*
Sets the page size to pageSize.

To get the current QPageSize use pageLayout().pageSize().

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new page size to a new page. You should not call any painting methods between a call to setPageSize() and newPage() as the wrong paint metrics may be used.

Returns true if the page size was successfully set to pageSize.

This function was introduced in  Qt 5.3.

See also pageSize() and pageLayout().
*/
func (this *QPagedPaintDevice) SetPageSize(pageSize QPageSize_ITF) bool {
	var convArg0 unsafe.Pointer
	if pageSize != nil && pageSize.QPageSize_PTR() != nil {
		convArg0 = pageSize.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeERK9QPageSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:226
// index:1
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSize(QPagedPaintDevice::PageSize)

/*
Sets the page size to pageSize.

To get the current QPageSize use pageLayout().pageSize().

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new page size to a new page. You should not call any painting methods between a call to setPageSize() and newPage() as the wrong paint metrics may be used.

Returns true if the page size was successfully set to pageSize.

This function was introduced in  Qt 5.3.

See also pageSize() and pageLayout().
*/
func (this *QPagedPaintDevice) SetPageSize_1(size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeENS_8PageSizeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:221
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageOrientation(QPageLayout::Orientation)

/*
Sets the page orientation.

The page orientation is used to define the orientation of the page size when obtaining the page rect.

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new orientation to a new page. You should not call any painting methods between a call to setPageOrientation() and newPage() as the wrong paint metrics may be used.

To get the current QPageLayout::Orientation use pageLayout().pageOrientation().

Returns true if the page orientation was successfully set to orientation.

This function was introduced in  Qt 5.3.

See also pageLayout().
*/
func (this *QPagedPaintDevice) SetPageOrientation(orientation int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice18setPageOrientationEN11QPageLayout11OrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:222
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setPageMargins(const QMarginsF &)

/*
Set the page margins in the current page layout units.

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new margins to a new page. You should not call any painting methods between a call to setPageMargins() and newPage() as the wrong paint metrics may be used.

To get the current page margins use pageLayout().pageMargins().

Returns true if the page margins were successfully set to margins.

This function was introduced in  Qt 5.3.

See also pageLayout().
*/
func (this *QPagedPaintDevice) SetPageMargins(margins qtcore.QMarginsF_ITF) bool {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:223
// index:1
// Public Visibility=Default Availability=Available
// [1] bool setPageMargins(const QMarginsF &, QPageLayout::Unit)

/*
Set the page margins in the current page layout units.

You should call this before calling QPainter::begin(), or immediately before calling newPage() to apply the new margins to a new page. You should not call any painting methods between a call to setPageMargins() and newPage() as the wrong paint metrics may be used.

To get the current page margins use pageLayout().pageMargins().

Returns true if the page margins were successfully set to margins.

This function was introduced in  Qt 5.3.

See also pageLayout().
*/
func (this *QPagedPaintDevice) SetPageMargins_1(margins qtcore.QMarginsF_ITF, units int) bool {
	var convArg0 unsafe.Pointer
	if margins != nil && margins.QMarginsF_PTR() != nil {
		convArg0 = margins.QMarginsF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsFN11QPageLayout4UnitE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, units)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:224
// index:0
// Public Visibility=Default Availability=Available
// [8] QPageLayout pageLayout() const

/*
Returns the current page layout. Use this method to access the current QPageSize, QPageLayout::Orientation, QMarginsF, fullRect() and paintRect().

Note that you cannot use the setters on the returned object, you must either call the individual QPagedPaintDevice setters or use setPageLayout().

This function was introduced in  Qt 5.3.

See also setPageLayout(), setPageSize(), setPageOrientation(), and setPageMargins().
*/
func (this *QPagedPaintDevice) PageLayout() *QPageLayout /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:227
// index:0
// Public Visibility=Default Availability=Available
// [4] QPagedPaintDevice::PageSize pageSize() const

/*
Returns the currently used page size.

See also setPageSize().
*/
func (this *QPagedPaintDevice) PageSize() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice8pageSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:229
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setPageSizeMM(const QSizeF &)

/*
Sets the page size to size. size is specified in millimeters.

If the size matches a standard QPagedPaintDevice::PageSize then that page size will be used, otherwise QPagedPaintDevice::Custom will be set.

See also pageSizeMM().
*/
func (this *QPagedPaintDevice) SetPageSizeMM(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:230
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF pageSizeMM() const

/*
Returns the page size in millimeters.

See also setPageSizeMM().
*/
func (this *QPagedPaintDevice) PageSizeMM() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageSizeMMEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:241
// index:0
// Public Visibility=Default Availability=Available
// [32] QPagedPaintDevice::Margins margins() const

/*
Returns the current margins of the paint device. The default is 0.

Margins are specified in millimeters.

See also setMargins().
*/
func (this *QPagedPaintDevice) Margins() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice7marginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:246
// index:0
// Protected Visibility=Default Availability=Available
// [8] QPageLayout devicePageLayout() const

/*

 */
func (this *QPagedPaintDevice) DevicePageLayout() *QPageLayout /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK17QPagedPaintDevice16devicePageLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:247
// index:1
// Protected Visibility=Default Availability=Available
// [8] QPageLayout & devicePageLayout()

/*

 */
func (this *QPagedPaintDevice) DevicePageLayout_1() *QPageLayout {
	rv, err := qtrt.InvokeQtFunc6("_ZN17QPagedPaintDevice16devicePageLayoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPageLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPageLayout)
	return rv2
}

/*
This enum type lists the available page sizes as defined in the Postscript PPD standard. These values are duplicated in QPageSize and QPrinter.

The defined sizes are:

QPagedPaintDevice::AnsiALetter= Letter
QPagedPaintDevice::AnsiBLedger= Ledger
QPagedPaintDevice::EnvelopeDLDLE= DLE


Due to historic reasons QPageSize::Executive is not the same as the standard Postscript and Windows Executive size, use QPageSize::ExecutiveStandard instead.

The Postscript standard size QPageSize::Folio is different to the Windows DMPAPER_FOLIO size, use the Postscript standard size QPageSize::FanFoldGermanLegal if needed.

*/
type QPagedPaintDevice__PageSize = int

//
const QPagedPaintDevice__A4 QPagedPaintDevice__PageSize = 0

//
const QPagedPaintDevice__B5 QPagedPaintDevice__PageSize = 1

//
const QPagedPaintDevice__Letter QPagedPaintDevice__PageSize = 2

//
const QPagedPaintDevice__Legal QPagedPaintDevice__PageSize = 3

//
const QPagedPaintDevice__Executive QPagedPaintDevice__PageSize = 4

//
const QPagedPaintDevice__A0 QPagedPaintDevice__PageSize = 5

//
const QPagedPaintDevice__A1 QPagedPaintDevice__PageSize = 6

//
const QPagedPaintDevice__A2 QPagedPaintDevice__PageSize = 7

//
const QPagedPaintDevice__A3 QPagedPaintDevice__PageSize = 8

//
const QPagedPaintDevice__A5 QPagedPaintDevice__PageSize = 9

//
const QPagedPaintDevice__A6 QPagedPaintDevice__PageSize = 10

//
const QPagedPaintDevice__A7 QPagedPaintDevice__PageSize = 11

//
const QPagedPaintDevice__A8 QPagedPaintDevice__PageSize = 12

//
const QPagedPaintDevice__A9 QPagedPaintDevice__PageSize = 13

//
const QPagedPaintDevice__B0 QPagedPaintDevice__PageSize = 14

//
const QPagedPaintDevice__B1 QPagedPaintDevice__PageSize = 15

//
const QPagedPaintDevice__B10 QPagedPaintDevice__PageSize = 16

//
const QPagedPaintDevice__B2 QPagedPaintDevice__PageSize = 17

//
const QPagedPaintDevice__B3 QPagedPaintDevice__PageSize = 18

//
const QPagedPaintDevice__B4 QPagedPaintDevice__PageSize = 19

//
const QPagedPaintDevice__B6 QPagedPaintDevice__PageSize = 20

//
const QPagedPaintDevice__B7 QPagedPaintDevice__PageSize = 21

//
const QPagedPaintDevice__B8 QPagedPaintDevice__PageSize = 22

//
const QPagedPaintDevice__B9 QPagedPaintDevice__PageSize = 23

//
const QPagedPaintDevice__C5E QPagedPaintDevice__PageSize = 24

//
const QPagedPaintDevice__Comm10E QPagedPaintDevice__PageSize = 25

//
const QPagedPaintDevice__DLE QPagedPaintDevice__PageSize = 26

//
const QPagedPaintDevice__Folio QPagedPaintDevice__PageSize = 27

//
const QPagedPaintDevice__Ledger QPagedPaintDevice__PageSize = 28

//
const QPagedPaintDevice__Tabloid QPagedPaintDevice__PageSize = 29

//
const QPagedPaintDevice__Custom QPagedPaintDevice__PageSize = 30

//
const QPagedPaintDevice__A10 QPagedPaintDevice__PageSize = 31

//
const QPagedPaintDevice__A3Extra QPagedPaintDevice__PageSize = 32

//
const QPagedPaintDevice__A4Extra QPagedPaintDevice__PageSize = 33

//
const QPagedPaintDevice__A4Plus QPagedPaintDevice__PageSize = 34

//
const QPagedPaintDevice__A4Small QPagedPaintDevice__PageSize = 35

//
const QPagedPaintDevice__A5Extra QPagedPaintDevice__PageSize = 36

//
const QPagedPaintDevice__B5Extra QPagedPaintDevice__PageSize = 37

//
const QPagedPaintDevice__JisB0 QPagedPaintDevice__PageSize = 38

//
const QPagedPaintDevice__JisB1 QPagedPaintDevice__PageSize = 39

//
const QPagedPaintDevice__JisB2 QPagedPaintDevice__PageSize = 40

//
const QPagedPaintDevice__JisB3 QPagedPaintDevice__PageSize = 41

//
const QPagedPaintDevice__JisB4 QPagedPaintDevice__PageSize = 42

//
const QPagedPaintDevice__JisB5 QPagedPaintDevice__PageSize = 43

//
const QPagedPaintDevice__JisB6 QPagedPaintDevice__PageSize = 44

//
const QPagedPaintDevice__JisB7 QPagedPaintDevice__PageSize = 45

//
const QPagedPaintDevice__JisB8 QPagedPaintDevice__PageSize = 46

//
const QPagedPaintDevice__JisB9 QPagedPaintDevice__PageSize = 47

//
const QPagedPaintDevice__JisB10 QPagedPaintDevice__PageSize = 48

//
const QPagedPaintDevice__AnsiC QPagedPaintDevice__PageSize = 49

//
const QPagedPaintDevice__AnsiD QPagedPaintDevice__PageSize = 50

//
const QPagedPaintDevice__AnsiE QPagedPaintDevice__PageSize = 51

//
const QPagedPaintDevice__LegalExtra QPagedPaintDevice__PageSize = 52

//
const QPagedPaintDevice__LetterExtra QPagedPaintDevice__PageSize = 53

//
const QPagedPaintDevice__LetterPlus QPagedPaintDevice__PageSize = 54

//
const QPagedPaintDevice__LetterSmall QPagedPaintDevice__PageSize = 55

//
const QPagedPaintDevice__TabloidExtra QPagedPaintDevice__PageSize = 56

//
const QPagedPaintDevice__ArchA QPagedPaintDevice__PageSize = 57

//
const QPagedPaintDevice__ArchB QPagedPaintDevice__PageSize = 58

//
const QPagedPaintDevice__ArchC QPagedPaintDevice__PageSize = 59

//
const QPagedPaintDevice__ArchD QPagedPaintDevice__PageSize = 60

//
const QPagedPaintDevice__ArchE QPagedPaintDevice__PageSize = 61

//
const QPagedPaintDevice__Imperial7x9 QPagedPaintDevice__PageSize = 62

//
const QPagedPaintDevice__Imperial8x10 QPagedPaintDevice__PageSize = 63

//
const QPagedPaintDevice__Imperial9x11 QPagedPaintDevice__PageSize = 64

//
const QPagedPaintDevice__Imperial9x12 QPagedPaintDevice__PageSize = 65

//
const QPagedPaintDevice__Imperial10x11 QPagedPaintDevice__PageSize = 66

//
const QPagedPaintDevice__Imperial10x13 QPagedPaintDevice__PageSize = 67

//
const QPagedPaintDevice__Imperial10x14 QPagedPaintDevice__PageSize = 68

//
const QPagedPaintDevice__Imperial12x11 QPagedPaintDevice__PageSize = 69

//
const QPagedPaintDevice__Imperial15x11 QPagedPaintDevice__PageSize = 70

//
const QPagedPaintDevice__ExecutiveStandard QPagedPaintDevice__PageSize = 71

//
const QPagedPaintDevice__Note QPagedPaintDevice__PageSize = 72

//
const QPagedPaintDevice__Quarto QPagedPaintDevice__PageSize = 73

//
const QPagedPaintDevice__Statement QPagedPaintDevice__PageSize = 74

//
const QPagedPaintDevice__SuperA QPagedPaintDevice__PageSize = 75

//
const QPagedPaintDevice__SuperB QPagedPaintDevice__PageSize = 76

//
const QPagedPaintDevice__Postcard QPagedPaintDevice__PageSize = 77

//
const QPagedPaintDevice__DoublePostcard QPagedPaintDevice__PageSize = 78

//
const QPagedPaintDevice__Prc16K QPagedPaintDevice__PageSize = 79

//
const QPagedPaintDevice__Prc32K QPagedPaintDevice__PageSize = 80

//
const QPagedPaintDevice__Prc32KBig QPagedPaintDevice__PageSize = 81

//
const QPagedPaintDevice__FanFoldUS QPagedPaintDevice__PageSize = 82

//
const QPagedPaintDevice__FanFoldGerman QPagedPaintDevice__PageSize = 83

//
const QPagedPaintDevice__FanFoldGermanLegal QPagedPaintDevice__PageSize = 84

//
const QPagedPaintDevice__EnvelopeB4 QPagedPaintDevice__PageSize = 85

//
const QPagedPaintDevice__EnvelopeB5 QPagedPaintDevice__PageSize = 86

//
const QPagedPaintDevice__EnvelopeB6 QPagedPaintDevice__PageSize = 87

//
const QPagedPaintDevice__EnvelopeC0 QPagedPaintDevice__PageSize = 88

//
const QPagedPaintDevice__EnvelopeC1 QPagedPaintDevice__PageSize = 89

//
const QPagedPaintDevice__EnvelopeC2 QPagedPaintDevice__PageSize = 90

//
const QPagedPaintDevice__EnvelopeC3 QPagedPaintDevice__PageSize = 91

//
const QPagedPaintDevice__EnvelopeC4 QPagedPaintDevice__PageSize = 92

//
const QPagedPaintDevice__EnvelopeC6 QPagedPaintDevice__PageSize = 93

//
const QPagedPaintDevice__EnvelopeC65 QPagedPaintDevice__PageSize = 94

//
const QPagedPaintDevice__EnvelopeC7 QPagedPaintDevice__PageSize = 95

//
const QPagedPaintDevice__Envelope9 QPagedPaintDevice__PageSize = 96

//
const QPagedPaintDevice__Envelope11 QPagedPaintDevice__PageSize = 97

//
const QPagedPaintDevice__Envelope12 QPagedPaintDevice__PageSize = 98

//
const QPagedPaintDevice__Envelope14 QPagedPaintDevice__PageSize = 99

//
const QPagedPaintDevice__EnvelopeMonarch QPagedPaintDevice__PageSize = 100

//
const QPagedPaintDevice__EnvelopePersonal QPagedPaintDevice__PageSize = 101

//
const QPagedPaintDevice__EnvelopeChou3 QPagedPaintDevice__PageSize = 102

//
const QPagedPaintDevice__EnvelopeChou4 QPagedPaintDevice__PageSize = 103

//
const QPagedPaintDevice__EnvelopeInvite QPagedPaintDevice__PageSize = 104

//
const QPagedPaintDevice__EnvelopeItalian QPagedPaintDevice__PageSize = 105

//
const QPagedPaintDevice__EnvelopeKaku2 QPagedPaintDevice__PageSize = 106

//
const QPagedPaintDevice__EnvelopeKaku3 QPagedPaintDevice__PageSize = 107

//
const QPagedPaintDevice__EnvelopePrc1 QPagedPaintDevice__PageSize = 108

//
const QPagedPaintDevice__EnvelopePrc2 QPagedPaintDevice__PageSize = 109

//
const QPagedPaintDevice__EnvelopePrc3 QPagedPaintDevice__PageSize = 110

//
const QPagedPaintDevice__EnvelopePrc4 QPagedPaintDevice__PageSize = 111

//
const QPagedPaintDevice__EnvelopePrc5 QPagedPaintDevice__PageSize = 112

//
const QPagedPaintDevice__EnvelopePrc6 QPagedPaintDevice__PageSize = 113

//
const QPagedPaintDevice__EnvelopePrc7 QPagedPaintDevice__PageSize = 114

//
const QPagedPaintDevice__EnvelopePrc8 QPagedPaintDevice__PageSize = 115

//
const QPagedPaintDevice__EnvelopePrc9 QPagedPaintDevice__PageSize = 116

//
const QPagedPaintDevice__EnvelopePrc10 QPagedPaintDevice__PageSize = 117

//
const QPagedPaintDevice__EnvelopeYou4 QPagedPaintDevice__PageSize = 118

//
const QPagedPaintDevice__LastPageSize QPagedPaintDevice__PageSize = 118

//
const QPagedPaintDevice__NPageSize QPagedPaintDevice__PageSize = 118

//
const QPagedPaintDevice__NPaperSize QPagedPaintDevice__PageSize = 118

//
const QPagedPaintDevice__AnsiA QPagedPaintDevice__PageSize = 2

//
const QPagedPaintDevice__AnsiB QPagedPaintDevice__PageSize = 28

//
const QPagedPaintDevice__EnvelopeC5 QPagedPaintDevice__PageSize = 24

//
const QPagedPaintDevice__EnvelopeDL QPagedPaintDevice__PageSize = 26

//
const QPagedPaintDevice__Envelope10 QPagedPaintDevice__PageSize = 25

func (this *QPagedPaintDevice) PageSizeItemName(val int) string {
	switch val {
	case QPagedPaintDevice__A4: // 0
		return "A4"
	case QPagedPaintDevice__B5: // 1
		return "B5"
	case QPagedPaintDevice__Letter: // 2
		return "Letter,AnsiA"
	case QPagedPaintDevice__Legal: // 3
		return "Legal"
	case QPagedPaintDevice__Executive: // 4
		return "Executive"
	case QPagedPaintDevice__A0: // 5
		return "A0"
	case QPagedPaintDevice__A1: // 6
		return "A1"
	case QPagedPaintDevice__A2: // 7
		return "A2"
	case QPagedPaintDevice__A3: // 8
		return "A3"
	case QPagedPaintDevice__A5: // 9
		return "A5"
	case QPagedPaintDevice__A6: // 10
		return "A6"
	case QPagedPaintDevice__A7: // 11
		return "A7"
	case QPagedPaintDevice__A8: // 12
		return "A8"
	case QPagedPaintDevice__A9: // 13
		return "A9"
	case QPagedPaintDevice__B0: // 14
		return "B0"
	case QPagedPaintDevice__B1: // 15
		return "B1"
	case QPagedPaintDevice__B10: // 16
		return "B10"
	case QPagedPaintDevice__B2: // 17
		return "B2"
	case QPagedPaintDevice__B3: // 18
		return "B3"
	case QPagedPaintDevice__B4: // 19
		return "B4"
	case QPagedPaintDevice__B6: // 20
		return "B6"
	case QPagedPaintDevice__B7: // 21
		return "B7"
	case QPagedPaintDevice__B8: // 22
		return "B8"
	case QPagedPaintDevice__B9: // 23
		return "B9"
	case QPagedPaintDevice__C5E: // 24
		return "C5E,EnvelopeC5"
	case QPagedPaintDevice__Comm10E: // 25
		return "Comm10E,Envelope10"
	case QPagedPaintDevice__DLE: // 26
		return "DLE,EnvelopeDL"
	case QPagedPaintDevice__Folio: // 27
		return "Folio"
	case QPagedPaintDevice__Ledger: // 28
		return "Ledger,AnsiB"
	case QPagedPaintDevice__Tabloid: // 29
		return "Tabloid"
	case QPagedPaintDevice__Custom: // 30
		return "Custom"
	case QPagedPaintDevice__A10: // 31
		return "A10"
	case QPagedPaintDevice__A3Extra: // 32
		return "A3Extra"
	case QPagedPaintDevice__A4Extra: // 33
		return "A4Extra"
	case QPagedPaintDevice__A4Plus: // 34
		return "A4Plus"
	case QPagedPaintDevice__A4Small: // 35
		return "A4Small"
	case QPagedPaintDevice__A5Extra: // 36
		return "A5Extra"
	case QPagedPaintDevice__B5Extra: // 37
		return "B5Extra"
	case QPagedPaintDevice__JisB0: // 38
		return "JisB0"
	case QPagedPaintDevice__JisB1: // 39
		return "JisB1"
	case QPagedPaintDevice__JisB2: // 40
		return "JisB2"
	case QPagedPaintDevice__JisB3: // 41
		return "JisB3"
	case QPagedPaintDevice__JisB4: // 42
		return "JisB4"
	case QPagedPaintDevice__JisB5: // 43
		return "JisB5"
	case QPagedPaintDevice__JisB6: // 44
		return "JisB6"
	case QPagedPaintDevice__JisB7: // 45
		return "JisB7"
	case QPagedPaintDevice__JisB8: // 46
		return "JisB8"
	case QPagedPaintDevice__JisB9: // 47
		return "JisB9"
	case QPagedPaintDevice__JisB10: // 48
		return "JisB10"
	case QPagedPaintDevice__AnsiC: // 49
		return "AnsiC"
	case QPagedPaintDevice__AnsiD: // 50
		return "AnsiD"
	case QPagedPaintDevice__AnsiE: // 51
		return "AnsiE"
	case QPagedPaintDevice__LegalExtra: // 52
		return "LegalExtra"
	case QPagedPaintDevice__LetterExtra: // 53
		return "LetterExtra"
	case QPagedPaintDevice__LetterPlus: // 54
		return "LetterPlus"
	case QPagedPaintDevice__LetterSmall: // 55
		return "LetterSmall"
	case QPagedPaintDevice__TabloidExtra: // 56
		return "TabloidExtra"
	case QPagedPaintDevice__ArchA: // 57
		return "ArchA"
	case QPagedPaintDevice__ArchB: // 58
		return "ArchB"
	case QPagedPaintDevice__ArchC: // 59
		return "ArchC"
	case QPagedPaintDevice__ArchD: // 60
		return "ArchD"
	case QPagedPaintDevice__ArchE: // 61
		return "ArchE"
	case QPagedPaintDevice__Imperial7x9: // 62
		return "Imperial7x9"
	case QPagedPaintDevice__Imperial8x10: // 63
		return "Imperial8x10"
	case QPagedPaintDevice__Imperial9x11: // 64
		return "Imperial9x11"
	case QPagedPaintDevice__Imperial9x12: // 65
		return "Imperial9x12"
	case QPagedPaintDevice__Imperial10x11: // 66
		return "Imperial10x11"
	case QPagedPaintDevice__Imperial10x13: // 67
		return "Imperial10x13"
	case QPagedPaintDevice__Imperial10x14: // 68
		return "Imperial10x14"
	case QPagedPaintDevice__Imperial12x11: // 69
		return "Imperial12x11"
	case QPagedPaintDevice__Imperial15x11: // 70
		return "Imperial15x11"
	case QPagedPaintDevice__ExecutiveStandard: // 71
		return "ExecutiveStandard"
	case QPagedPaintDevice__Note: // 72
		return "Note"
	case QPagedPaintDevice__Quarto: // 73
		return "Quarto"
	case QPagedPaintDevice__Statement: // 74
		return "Statement"
	case QPagedPaintDevice__SuperA: // 75
		return "SuperA"
	case QPagedPaintDevice__SuperB: // 76
		return "SuperB"
	case QPagedPaintDevice__Postcard: // 77
		return "Postcard"
	case QPagedPaintDevice__DoublePostcard: // 78
		return "DoublePostcard"
	case QPagedPaintDevice__Prc16K: // 79
		return "Prc16K"
	case QPagedPaintDevice__Prc32K: // 80
		return "Prc32K"
	case QPagedPaintDevice__Prc32KBig: // 81
		return "Prc32KBig"
	case QPagedPaintDevice__FanFoldUS: // 82
		return "FanFoldUS"
	case QPagedPaintDevice__FanFoldGerman: // 83
		return "FanFoldGerman"
	case QPagedPaintDevice__FanFoldGermanLegal: // 84
		return "FanFoldGermanLegal"
	case QPagedPaintDevice__EnvelopeB4: // 85
		return "EnvelopeB4"
	case QPagedPaintDevice__EnvelopeB5: // 86
		return "EnvelopeB5"
	case QPagedPaintDevice__EnvelopeB6: // 87
		return "EnvelopeB6"
	case QPagedPaintDevice__EnvelopeC0: // 88
		return "EnvelopeC0"
	case QPagedPaintDevice__EnvelopeC1: // 89
		return "EnvelopeC1"
	case QPagedPaintDevice__EnvelopeC2: // 90
		return "EnvelopeC2"
	case QPagedPaintDevice__EnvelopeC3: // 91
		return "EnvelopeC3"
	case QPagedPaintDevice__EnvelopeC4: // 92
		return "EnvelopeC4"
	case QPagedPaintDevice__EnvelopeC6: // 93
		return "EnvelopeC6"
	case QPagedPaintDevice__EnvelopeC65: // 94
		return "EnvelopeC65"
	case QPagedPaintDevice__EnvelopeC7: // 95
		return "EnvelopeC7"
	case QPagedPaintDevice__Envelope9: // 96
		return "Envelope9"
	case QPagedPaintDevice__Envelope11: // 97
		return "Envelope11"
	case QPagedPaintDevice__Envelope12: // 98
		return "Envelope12"
	case QPagedPaintDevice__Envelope14: // 99
		return "Envelope14"
	case QPagedPaintDevice__EnvelopeMonarch: // 100
		return "EnvelopeMonarch"
	case QPagedPaintDevice__EnvelopePersonal: // 101
		return "EnvelopePersonal"
	case QPagedPaintDevice__EnvelopeChou3: // 102
		return "EnvelopeChou3"
	case QPagedPaintDevice__EnvelopeChou4: // 103
		return "EnvelopeChou4"
	case QPagedPaintDevice__EnvelopeInvite: // 104
		return "EnvelopeInvite"
	case QPagedPaintDevice__EnvelopeItalian: // 105
		return "EnvelopeItalian"
	case QPagedPaintDevice__EnvelopeKaku2: // 106
		return "EnvelopeKaku2"
	case QPagedPaintDevice__EnvelopeKaku3: // 107
		return "EnvelopeKaku3"
	case QPagedPaintDevice__EnvelopePrc1: // 108
		return "EnvelopePrc1"
	case QPagedPaintDevice__EnvelopePrc2: // 109
		return "EnvelopePrc2"
	case QPagedPaintDevice__EnvelopePrc3: // 110
		return "EnvelopePrc3"
	case QPagedPaintDevice__EnvelopePrc4: // 111
		return "EnvelopePrc4"
	case QPagedPaintDevice__EnvelopePrc5: // 112
		return "EnvelopePrc5"
	case QPagedPaintDevice__EnvelopePrc6: // 113
		return "EnvelopePrc6"
	case QPagedPaintDevice__EnvelopePrc7: // 114
		return "EnvelopePrc7"
	case QPagedPaintDevice__EnvelopePrc8: // 115
		return "EnvelopePrc8"
	case QPagedPaintDevice__EnvelopePrc9: // 116
		return "EnvelopePrc9"
	case QPagedPaintDevice__EnvelopePrc10: // 117
		return "EnvelopePrc10"
	case QPagedPaintDevice__EnvelopeYou4: // 118
		return "EnvelopeYou4,LastPageSize,NPageSize,NPaperSize"
		// case QPagedPaintDevice__LastPageSize: // 118
		// return ""
		// case QPagedPaintDevice__NPageSize: // 118
		// return ""
		// case QPagedPaintDevice__NPaperSize: // 118
		// return ""
		// case QPagedPaintDevice__AnsiA: // 2
		// return ""
		// case QPagedPaintDevice__AnsiB: // 28
		// return ""
		// case QPagedPaintDevice__EnvelopeC5: // 24
		// return ""
		// case QPagedPaintDevice__EnvelopeDL: // 26
		// return ""
		// case QPagedPaintDevice__Envelope10: // 25
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPagedPaintDevice_PageSizeItemName(val int) string {
	var nilthis *QPagedPaintDevice
	return nilthis.PageSizeItemName(val)
}

/*
The PdfVersion enum describes the version of the PDF file that is produced by QPrinter or QPdfWriter.



This enum was introduced or modified in  Qt 5.10.

*/
type QPagedPaintDevice__PdfVersion = int

//
const QPagedPaintDevice__PdfVersion_1_4 QPagedPaintDevice__PdfVersion = 0

//
const QPagedPaintDevice__PdfVersion_A1b QPagedPaintDevice__PdfVersion = 1

func (this *QPagedPaintDevice) PdfVersionItemName(val int) string {
	switch val {
	case QPagedPaintDevice__PdfVersion_1_4: // 0
		return "PdfVersion_1_4"
	case QPagedPaintDevice__PdfVersion_A1b: // 1
		return "PdfVersion_A1b"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QPagedPaintDevice_PdfVersionItemName(val int) string {
	var nilthis *QPagedPaintDevice
	return nilthis.PdfVersionItemName(val)
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
