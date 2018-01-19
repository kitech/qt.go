//  header block begin
// /usr/include/qt/QtGui/qpagesize.h
// #include <qpagesize.h>
// #include <QtGui>
package qtgui

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
type QPageSize struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpagesize.h:230
// index:0
// void QPageSize()
func NewQPageSize() *QPageSize {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPageSize{cthis}
}

// /usr/include/qt/QtGui/qpagesize.h:231
// index:1
// void QPageSize(enum QPageSize::PageSizeId)
func NewQPageSize_1(pageSizeId int) *QPageSize {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2ENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID, cthis, &pageSizeId)
	gopp.ErrPrint(err, rv)
	return &QPageSize{cthis}
}

// /usr/include/qt/QtGui/qpagesize.h:232
// index:2
// void QPageSize(const class QSize &, const class QString &, enum QPageSize::SizeMatchPolicy)
func NewQPageSize_2(pointSize unsafe.Pointer, name unsafe.Pointer, matchPolicy int) *QPageSize {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2ERK5QSizeRK7QStringNS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_VOID, cthis, pointSize, name, &matchPolicy)
	gopp.ErrPrint(err, rv)
	return &QPageSize{cthis}
}

// /usr/include/qt/QtGui/qpagesize.h:235
// index:3
// void QPageSize(const class QSizeF &, enum QPageSize::Unit, const class QString &, enum QPageSize::SizeMatchPolicy)
func NewQPageSize_3(size unsafe.Pointer, units int, name unsafe.Pointer, matchPolicy int) *QPageSize {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2ERK6QSizeFNS_4UnitERK7QStringNS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_VOID, cthis, size, &units, name, &matchPolicy)
	gopp.ErrPrint(err, rv)
	return &QPageSize{cthis}
}

// /usr/include/qt/QtGui/qpagesize.h:243
// index:0
// void ~QPageSize()
func DeleteQPageSize(*QPageSize) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:246
// index:0
// inline
// void swap(class QPageSize &)
func (this *QPageSize) Swap(other unsafe.Pointer) {
	// 0: (, QPageSize & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:249
// index:0
// bool isEquivalentTo(const class QPageSize &)
func (this *QPageSize) IsEquivalentTo(other unsafe.Pointer) {
	// 0: (, const QPageSize & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize14isEquivalentToERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:251
// index:0
// bool isValid()
func (this *QPageSize) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:253
// index:0
// QString key()
func (this *QPageSize) Key() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize3keyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:271
// index:1
// static
// QString key(enum QPageSize::PageSizeId)
func (this *QPageSize) Key_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize3keyENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_Key_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	var nilthis *QPageSize
	nilthis.Key_1(pageSizeId)
}

// /usr/include/qt/QtGui/qpagesize.h:254
// index:0
// QString name()
func (this *QPageSize) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:272
// index:1
// static
// QString name(enum QPageSize::PageSizeId)
func (this *QPageSize) Name_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize4nameENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_Name_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	var nilthis *QPageSize
	nilthis.Name_1(pageSizeId)
}

// /usr/include/qt/QtGui/qpagesize.h:256
// index:0
// QPageSize::PageSizeId id()
func (this *QPageSize) Id() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize2idEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:274
// index:1
// static
// QPageSize::PageSizeId id(const class QSize &, enum QPageSize::SizeMatchPolicy)
func (this *QPageSize) Id_1(pointSize unsafe.Pointer, matchPolicy int) {
	// 1: (const QSize & pointSize, QPageSize::SizeMatchPolicy matchPolicy), (pointSize, matchPolicy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize2idERK5QSizeNS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_Id_1(pointSize unsafe.Pointer, matchPolicy int) {
	// 1: (const QSize & pointSize, QPageSize::SizeMatchPolicy matchPolicy), (pointSize, matchPolicy)
	var nilthis *QPageSize
	nilthis.Id_1(pointSize, matchPolicy)
}

// /usr/include/qt/QtGui/qpagesize.h:276
// index:2
// static
// QPageSize::PageSizeId id(const class QSizeF &, enum QPageSize::Unit, enum QPageSize::SizeMatchPolicy)
func (this *QPageSize) Id_2(size unsafe.Pointer, units int, matchPolicy int) {
	// 2: (const QSizeF & size, QPageSize::Unit units, QPageSize::SizeMatchPolicy matchPolicy), (size, units, matchPolicy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize2idERK6QSizeFNS_4UnitENS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_Id_2(size unsafe.Pointer, units int, matchPolicy int) {
	// 2: (const QSizeF & size, QPageSize::Unit units, QPageSize::SizeMatchPolicy matchPolicy), (size, units, matchPolicy)
	var nilthis *QPageSize
	nilthis.Id_2(size, units, matchPolicy)
}

// /usr/include/qt/QtGui/qpagesize.h:279
// index:3
// static
// QPageSize::PageSizeId id(int)
func (this *QPageSize) Id_3(windowsId int) {
	// 3: (int windowsId), (windowsId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize2idEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_Id_3(windowsId int) {
	// 3: (int windowsId), (windowsId)
	var nilthis *QPageSize
	nilthis.Id_3(windowsId)
}

// /usr/include/qt/QtGui/qpagesize.h:258
// index:0
// int windowsId()
func (this *QPageSize) WindowsId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize9windowsIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:280
// index:1
// static
// int windowsId(enum QPageSize::PageSizeId)
func (this *QPageSize) WindowsId_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize9windowsIdENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_WindowsId_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	var nilthis *QPageSize
	nilthis.WindowsId_1(pageSizeId)
}

// /usr/include/qt/QtGui/qpagesize.h:260
// index:0
// QSizeF definitionSize()
func (this *QPageSize) DefinitionSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize14definitionSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:282
// index:1
// static
// QSizeF definitionSize(enum QPageSize::PageSizeId)
func (this *QPageSize) DefinitionSize_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize14definitionSizeENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_DefinitionSize_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	var nilthis *QPageSize
	nilthis.DefinitionSize_1(pageSizeId)
}

// /usr/include/qt/QtGui/qpagesize.h:261
// index:0
// QPageSize::Unit definitionUnits()
func (this *QPageSize) DefinitionUnits() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize15definitionUnitsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:283
// index:1
// static
// QPageSize::Unit definitionUnits(enum QPageSize::PageSizeId)
func (this *QPageSize) DefinitionUnits_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize15definitionUnitsENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_DefinitionUnits_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	var nilthis *QPageSize
	nilthis.DefinitionUnits_1(pageSizeId)
}

// /usr/include/qt/QtGui/qpagesize.h:263
// index:0
// QSizeF size(enum QPageSize::Unit)
func (this *QPageSize) Size(units int) {
	// 0: (, QPageSize::Unit units), (&units)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize4sizeENS_4UnitE", ffiqt.FFI_TYPE_VOID, this.cthis, &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:285
// index:1
// static
// QSizeF size(enum QPageSize::PageSizeId, enum QPageSize::Unit)
func (this *QPageSize) Size_1(pageSizeId int, units int) {
	// 1: (QPageSize::PageSizeId pageSizeId, QPageSize::Unit units), (pageSizeId, units)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize4sizeENS_10PageSizeIdENS_4UnitE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_Size_1(pageSizeId int, units int) {
	// 1: (QPageSize::PageSizeId pageSizeId, QPageSize::Unit units), (pageSizeId, units)
	var nilthis *QPageSize
	nilthis.Size_1(pageSizeId, units)
}

// /usr/include/qt/QtGui/qpagesize.h:264
// index:0
// QSize sizePoints()
func (this *QPageSize) SizePoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10sizePointsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:286
// index:1
// static
// QSize sizePoints(enum QPageSize::PageSizeId)
func (this *QPageSize) SizePoints_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize10sizePointsENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_SizePoints_1(pageSizeId int) {
	// 1: (QPageSize::PageSizeId pageSizeId), (pageSizeId)
	var nilthis *QPageSize
	nilthis.SizePoints_1(pageSizeId)
}

// /usr/include/qt/QtGui/qpagesize.h:265
// index:0
// QSize sizePixels(int)
func (this *QPageSize) SizePixels(resolution int) {
	// 0: (, int resolution), (&resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10sizePixelsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &resolution)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:287
// index:1
// static
// QSize sizePixels(enum QPageSize::PageSizeId, int)
func (this *QPageSize) SizePixels_1(pageSizeId int, resolution int) {
	// 1: (QPageSize::PageSizeId pageSizeId, int resolution), (pageSizeId, resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize10sizePixelsENS_10PageSizeIdEi", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QPageSize_SizePixels_1(pageSizeId int, resolution int) {
	// 1: (QPageSize::PageSizeId pageSizeId, int resolution), (pageSizeId, resolution)
	var nilthis *QPageSize
	nilthis.SizePixels_1(pageSizeId, resolution)
}

// /usr/include/qt/QtGui/qpagesize.h:267
// index:0
// QRectF rect(enum QPageSize::Unit)
func (this *QPageSize) Rect(units int) {
	// 0: (, QPageSize::Unit units), (&units)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize4rectENS_4UnitE", ffiqt.FFI_TYPE_VOID, this.cthis, &units)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:268
// index:0
// QRect rectPoints()
func (this *QPageSize) RectPoints() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10rectPointsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:269
// index:0
// QRect rectPixels(int)
func (this *QPageSize) RectPixels(resolution int) {
	// 0: (, int resolution), (&resolution)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10rectPixelsEi", ffiqt.FFI_TYPE_VOID, this.cthis, &resolution)
	gopp.ErrPrint(err, rv)
}

//  body block end
