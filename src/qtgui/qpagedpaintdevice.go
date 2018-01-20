//  header block begin
// /usr/include/qt/QtGui/qpagedpaintdevice.h
// #include <qpagedpaintdevice.h>
// #include <QtGui>
package qtgui

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
type QPagedPaintDevice struct {
	*QPaintDevice
}

func (this *QPagedPaintDevice) GetCthis() unsafe.Pointer {
	return this.QPaintDevice.GetCthis()
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:58
// index:0
// void QPagedPaintDevice()
func NewQPagedPaintDevice() *QPagedPaintDevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDeviceC1Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPagedPaintDeviceFromPointer(cthis)
	return gothis
}
func NewQPagedPaintDeviceFromPointer(cthis unsafe.Pointer) *QPagedPaintDevice {
	bcthis0 := NewQPaintDeviceFromPointer(cthis)
	return &QPagedPaintDevice{bcthis0}
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:244
// index:1
// void QPagedPaintDevice(class QPagedPaintDevicePrivate *)
func NewQPagedPaintDevice_1(dd unsafe.Pointer) *QPagedPaintDevice {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDeviceC1EP24QPagedPaintDevicePrivate", ffiqt.FFI_TYPE_VOID, cthis, dd)
	gopp.ErrPrint(err, rv)
	gothis := NewQPagedPaintDeviceFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:59
// index:0
// virtual
// void ~QPagedPaintDevice()
func DeleteQPagedPaintDevice(*QPagedPaintDevice) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDeviceD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:61
// index:0
// pure virtual
// bool newPage()
func (this *QPagedPaintDevice) NewPage() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice7newPageEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:219
// index:0
// bool setPageLayout(const class QPageLayout &)
func (this *QPagedPaintDevice) SetPageLayout(pageLayout unsafe.Pointer) {
	// 0: (, pageLayout const QPageLayout &), (pageLayout)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageLayoutERK11QPageLayout", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pageLayout)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:220
// index:0
// bool setPageSize(const class QPageSize &)
func (this *QPagedPaintDevice) SetPageSize(pageSize unsafe.Pointer) {
	// 0: (, pageSize const QPageSize &), (pageSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeERK9QPageSize", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pageSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:226
// index:1
// virtual
// void setPageSize(enum QPagedPaintDevice::PageSize)
func (this *QPagedPaintDevice) SetPageSize_1(size int) {
	// 1: (, size QPagedPaintDevice::PageSize), (&size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice11setPageSizeENS_8PageSizeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:221
// index:0
// bool setPageOrientation(class QPageLayout::Orientation)
func (this *QPagedPaintDevice) SetPageOrientation(orientation int) {
	// 0: (, orientation QPageLayout::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice18setPageOrientationEN11QPageLayout11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:222
// index:0
// bool setPageMargins(const class QMarginsF &)
func (this *QPagedPaintDevice) SetPageMargins(margins unsafe.Pointer) {
	// 0: (, margins const QMarginsF &), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:223
// index:1
// bool setPageMargins(const class QMarginsF &, class QPageLayout::Unit)
func (this *QPagedPaintDevice) SetPageMargins_1(margins unsafe.Pointer, units int) {
	// 1: (, margins const QMarginsF &, units QPageLayout::Unit), (margins, &units)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice14setPageMarginsERK9QMarginsFN11QPageLayout4UnitE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins, &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:224
// index:0
// QPageLayout pageLayout()
func (this *QPagedPaintDevice) PageLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:227
// index:0
// QPagedPaintDevice::PageSize pageSize()
func (this *QPagedPaintDevice) PageSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice8pageSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:229
// index:0
// virtual
// void setPageSizeMM(const class QSizeF &)
func (this *QPagedPaintDevice) SetPageSizeMM(size unsafe.Pointer) {
	// 0: (, size const QSizeF &), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice13setPageSizeMMERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:230
// index:0
// QSizeF pageSizeMM()
func (this *QPagedPaintDevice) PageSizeMM() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice10pageSizeMMEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:240
// index:0
// virtual
// void setMargins(const struct QPagedPaintDevice::Margins &)
func (this *QPagedPaintDevice) SetMargins(margins unsafe.Pointer) {
	// 0: (, margins const QPagedPaintDevice::Margins &), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice10setMarginsERKNS_7MarginsE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:241
// index:0
// QPagedPaintDevice::Margins margins()
func (this *QPagedPaintDevice) Margins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice7marginsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:245
// index:0
// QPagedPaintDevicePrivate * dd()
func (this *QPagedPaintDevice) Dd() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice2ddEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:246
// index:0
// QPageLayout devicePageLayout()
func (this *QPagedPaintDevice) DevicePageLayout() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK17QPagedPaintDevice16devicePageLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagedpaintdevice.h:247
// index:1
// QPageLayout & devicePageLayout()
func (this *QPagedPaintDevice) DevicePageLayout_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN17QPagedPaintDevice16devicePageLayoutEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
