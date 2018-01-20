//  header block begin
// /usr/include/qt/QtGui/qpagelayout.h
// #include <qpagelayout.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 31
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
type QPageLayout struct {
	*qtrt.CObject
}

func (this *QPageLayout) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qpagelayout.h:80
// index:0
// void QPageLayout()
func NewQPageLayout() *QPageLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayoutC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(cthis)
	return gothis
}
func NewQPageLayoutFromPointer(cthis unsafe.Pointer) *QPageLayout {
	return &QPageLayout{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpagelayout.h:81
// index:1
// void QPageLayout(const class QPageSize &, enum QPageLayout::Orientation, const class QMarginsF &, enum QPageLayout::Unit, const class QMarginsF &)
func NewQPageLayout_1(pageSize unsafe.Pointer, orientation int, margins unsafe.Pointer, units int, minMargins unsafe.Pointer) *QPageLayout {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayoutC2ERK9QPageSizeNS_11OrientationERK9QMarginsFNS_4UnitES6_", ffiqt.FFI_TYPE_VOID, cthis, pageSize, &orientation, margins, &units, minMargins)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageLayoutFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagelayout.h:89
// index:0
// void ~QPageLayout()
func DeleteQPageLayout(*QPageLayout) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayoutD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:91
// index:0
// inline
// void swap(class QPageLayout &)
func (this *QPageLayout) Swap(other unsafe.Pointer) {
	// 0: (, other QPageLayout &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:94
// index:0
// bool isEquivalentTo(const class QPageLayout &)
func (this *QPageLayout) IsEquivalentTo(other unsafe.Pointer) {
	// 0: (, other const QPageLayout &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14isEquivalentToERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:96
// index:0
// bool isValid()
func (this *QPageLayout) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:98
// index:0
// void setMode(enum QPageLayout::Mode)
func (this *QPageLayout) SetMode(mode int) {
	// 0: (, mode QPageLayout::Mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout7setModeENS_4ModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:99
// index:0
// QPageLayout::Mode mode()
func (this *QPageLayout) Mode() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout4modeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:101
// index:0
// void setPageSize(const class QPageSize &, const class QMarginsF &)
func (this *QPageLayout) SetPageSize(pageSize unsafe.Pointer, minMargins unsafe.Pointer) {
	// 0: (, pageSize const QPageSize &, minMargins const QMarginsF &), (pageSize, minMargins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout11setPageSizeERK9QPageSizeRK9QMarginsF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pageSize, minMargins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:103
// index:0
// QPageSize pageSize()
func (this *QPageLayout) PageSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout8pageSizeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:105
// index:0
// void setOrientation(enum QPageLayout::Orientation)
func (this *QPageLayout) SetOrientation(orientation int) {
	// 0: (, orientation QPageLayout::Orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout14setOrientationENS_11OrientationE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:106
// index:0
// QPageLayout::Orientation orientation()
func (this *QPageLayout) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout11orientationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:108
// index:0
// void setUnits(enum QPageLayout::Unit)
func (this *QPageLayout) SetUnits(units int) {
	// 0: (, units QPageLayout::Unit), (&units)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout8setUnitsENS_4UnitE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:109
// index:0
// QPageLayout::Unit units()
func (this *QPageLayout) Units() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout5unitsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:111
// index:0
// bool setMargins(const class QMarginsF &)
func (this *QPageLayout) SetMargins(margins unsafe.Pointer) {
	// 0: (, margins const QMarginsF &), (margins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout10setMarginsERK9QMarginsF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), margins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:112
// index:0
// bool setLeftMargin(qreal)
func (this *QPageLayout) SetLeftMargin(leftMargin float64) {
	// 0: (, leftMargin qreal), (&leftMargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout13setLeftMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &leftMargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:113
// index:0
// bool setRightMargin(qreal)
func (this *QPageLayout) SetRightMargin(rightMargin float64) {
	// 0: (, rightMargin qreal), (&rightMargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout14setRightMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &rightMargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:114
// index:0
// bool setTopMargin(qreal)
func (this *QPageLayout) SetTopMargin(topMargin float64) {
	// 0: (, topMargin qreal), (&topMargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout12setTopMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &topMargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:115
// index:0
// bool setBottomMargin(qreal)
func (this *QPageLayout) SetBottomMargin(bottomMargin float64) {
	// 0: (, bottomMargin qreal), (&bottomMargin)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout15setBottomMarginEd", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &bottomMargin)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:117
// index:0
// QMarginsF margins()
func (this *QPageLayout) Margins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout7marginsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:118
// index:1
// QMarginsF margins(enum QPageLayout::Unit)
func (this *QPageLayout) Margins_1(units int) {
	// 1: (, units QPageLayout::Unit), (&units)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout7marginsENS_4UnitE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:119
// index:0
// QMargins marginsPoints()
func (this *QPageLayout) MarginsPoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout13marginsPointsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:120
// index:0
// QMargins marginsPixels(int)
func (this *QPageLayout) MarginsPixels(resolution int) {
	// 0: (, resolution int), (&resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout13marginsPixelsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:122
// index:0
// void setMinimumMargins(const class QMarginsF &)
func (this *QPageLayout) SetMinimumMargins(minMargins unsafe.Pointer) {
	// 0: (, minMargins const QMarginsF &), (minMargins)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QPageLayout17setMinimumMarginsERK9QMarginsF", ffiqt.FFI_TYPE_VOID, this.GetCthis(), minMargins)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:123
// index:0
// QMarginsF minimumMargins()
func (this *QPageLayout) MinimumMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14minimumMarginsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:124
// index:0
// QMarginsF maximumMargins()
func (this *QPageLayout) MaximumMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14maximumMarginsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:126
// index:0
// QRectF fullRect()
func (this *QPageLayout) FullRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout8fullRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:127
// index:1
// QRectF fullRect(enum QPageLayout::Unit)
func (this *QPageLayout) FullRect_1(units int) {
	// 1: (, units QPageLayout::Unit), (&units)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout8fullRectENS_4UnitE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:128
// index:0
// QRect fullRectPoints()
func (this *QPageLayout) FullRectPoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14fullRectPointsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:129
// index:0
// QRect fullRectPixels(int)
func (this *QPageLayout) FullRectPixels(resolution int) {
	// 0: (, resolution int), (&resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout14fullRectPixelsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:131
// index:0
// QRectF paintRect()
func (this *QPageLayout) PaintRect() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout9paintRectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:132
// index:1
// QRectF paintRect(enum QPageLayout::Unit)
func (this *QPageLayout) PaintRect_1(units int) {
	// 1: (, units QPageLayout::Unit), (&units)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout9paintRectENS_4UnitE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:133
// index:0
// QRect paintRectPoints()
func (this *QPageLayout) PaintRectPoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPointsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:134
// index:0
// QRect paintRectPixels(int)
func (this *QPageLayout) PaintRectPixels(resolution int) {
	// 0: (, resolution int), (&resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QPageLayout15paintRectPixelsEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
}

//  body block end
