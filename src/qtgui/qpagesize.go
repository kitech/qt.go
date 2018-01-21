package qtgui

// /usr/include/qt/QtGui/qpagesize.h
// #include <qpagesize.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
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
	*qtrt.CObject
}

func (this *QPageSize) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func NewQPageSizeFromPointer(cthis unsafe.Pointer) *QPageSize {
	return &QPageSize{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qpagesize.h:230
// index:0
// Public
// void QPageSize()
func NewQPageSize() *QPageSize {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:231
// index:1
// Public
// void QPageSize(enum QPageSize::PageSizeId)
func NewQPageSize_1(pageSizeId int) *QPageSize {
	cthis := qtrt.Calloc(1, 256) // 8
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2ENS_10PageSizeIdE", ffiqt.FFI_TYPE_VOID, cthis, &pageSizeId)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:232
// index:2
// Public
// void QPageSize(const class QSize &, const class QString &, enum QPageSize::SizeMatchPolicy)
func NewQPageSize_2(pointSize *qtcore.QSize, name *qtcore.QString, matchPolicy int) *QPageSize {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = pointSize.GetCthis()
	var convArg1 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2ERK5QSizeRK7QStringNS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, convArg1, &matchPolicy)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:235
// index:3
// Public
// void QPageSize(const class QSizeF &, enum QPageSize::Unit, const class QString &, enum QPageSize::SizeMatchPolicy)
func NewQPageSize_3(size *qtcore.QSizeF, units int, name *qtcore.QString, matchPolicy int) *QPageSize {
	cthis := qtrt.Calloc(1, 256) // 8
	var convArg0 = size.GetCthis()
	var convArg2 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeC2ERK6QSizeFNS_4UnitERK7QStringNS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, &units, convArg2, &matchPolicy)
	gopp.ErrPrint(err, rv)
	gothis := NewQPageSizeFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qpagesize.h:243
// index:0
// Public
// void ~QPageSize()
func DeleteQPageSize(*QPageSize) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSizeD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:246
// index:0
// Public inline
// void swap(class QPageSize &)
func (this *QPageSize) Swap(other *QPageSize) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:249
// index:0
// Public
// bool isEquivalentTo(const class QPageSize &)
func (this *QPageSize) IsEquivalentTo(other *QPageSize) bool {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize14isEquivalentToERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:251
// index:0
// Public
// bool isValid()
func (this *QPageSize) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:253
// index:0
// Public
// QString key()
func (this *QPageSize) Key() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize3keyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:271
// index:1
// Public static
// QString key(enum QPageSize::PageSizeId)
func (this *QPageSize) Key_1(pageSizeId int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize3keyENS_10PageSizeIdE", ffiqt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPageSize_Key_1(pageSizeId int) *qtcore.QString /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.Key_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:254
// index:0
// Public
// QString name()
func (this *QPageSize) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:272
// index:1
// Public static
// QString name(enum QPageSize::PageSizeId)
func (this *QPageSize) Name_1(pageSizeId int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize4nameENS_10PageSizeIdE", ffiqt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPageSize_Name_1(pageSizeId int) *qtcore.QString /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.Name_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:256
// index:0
// Public
// QPageSize::PageSizeId id()
func (this *QPageSize) Id() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize2idEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:274
// index:1
// Public static
// QPageSize::PageSizeId id(const class QSize &, enum QPageSize::SizeMatchPolicy)
func (this *QPageSize) Id_1(pointSize *qtcore.QSize, matchPolicy int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize2idERK5QSizeNS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_POINTER, pointSize, matchPolicy)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QPageSize_Id_1(pointSize *qtcore.QSize, matchPolicy int) int {
	var nilthis *QPageSize
	rv := nilthis.Id_1(pointSize, matchPolicy)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:276
// index:2
// Public static
// QPageSize::PageSizeId id(const class QSizeF &, enum QPageSize::Unit, enum QPageSize::SizeMatchPolicy)
func (this *QPageSize) Id_2(size *qtcore.QSizeF, units int, matchPolicy int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize2idERK6QSizeFNS_4UnitENS_15SizeMatchPolicyE", ffiqt.FFI_TYPE_POINTER, size, units, matchPolicy)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QPageSize_Id_2(size *qtcore.QSizeF, units int, matchPolicy int) int {
	var nilthis *QPageSize
	rv := nilthis.Id_2(size, units, matchPolicy)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:279
// index:3
// Public static
// QPageSize::PageSizeId id(int)
func (this *QPageSize) Id_3(windowsId int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize2idEi", ffiqt.FFI_TYPE_POINTER, windowsId)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QPageSize_Id_3(windowsId int) int {
	var nilthis *QPageSize
	rv := nilthis.Id_3(windowsId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:258
// index:0
// Public
// int windowsId()
func (this *QPageSize) WindowsId() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize9windowsIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qpagesize.h:280
// index:1
// Public static
// int windowsId(enum QPageSize::PageSizeId)
func (this *QPageSize) WindowsId_1(pageSizeId int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize9windowsIdENS_10PageSizeIdE", ffiqt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv) // 111
}
func QPageSize_WindowsId_1(pageSizeId int) int {
	var nilthis *QPageSize
	rv := nilthis.WindowsId_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:260
// index:0
// Public
// QSizeF definitionSize()
func (this *QPageSize) DefinitionSize() *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize14definitionSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:282
// index:1
// Public static
// QSizeF definitionSize(enum QPageSize::PageSizeId)
func (this *QPageSize) DefinitionSize_1(pageSizeId int) *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize14definitionSizeENS_10PageSizeIdE", ffiqt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPageSize_DefinitionSize_1(pageSizeId int) *qtcore.QSizeF /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.DefinitionSize_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:261
// index:0
// Public
// QPageSize::Unit definitionUnits()
func (this *QPageSize) DefinitionUnits() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize15definitionUnitsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qpagesize.h:283
// index:1
// Public static
// QPageSize::Unit definitionUnits(enum QPageSize::PageSizeId)
func (this *QPageSize) DefinitionUnits_1(pageSizeId int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize15definitionUnitsENS_10PageSizeIdE", ffiqt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	// return rv
	return int(rv)
}
func QPageSize_DefinitionUnits_1(pageSizeId int) int {
	var nilthis *QPageSize
	rv := nilthis.DefinitionUnits_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:263
// index:0
// Public
// QSizeF size(enum QPageSize::Unit)
func (this *QPageSize) Size(units int) *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize4sizeENS_4UnitE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &units)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:285
// index:1
// Public static
// QSizeF size(enum QPageSize::PageSizeId, enum QPageSize::Unit)
func (this *QPageSize) Size_1(pageSizeId int, units int) *qtcore.QSizeF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize4sizeENS_10PageSizeIdENS_4UnitE", ffiqt.FFI_TYPE_POINTER, pageSizeId, units)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPageSize_Size_1(pageSizeId int, units int) *qtcore.QSizeF /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.Size_1(pageSizeId, units)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:264
// index:0
// Public
// QSize sizePoints()
func (this *QPageSize) SizePoints() *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10sizePointsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:286
// index:1
// Public static
// QSize sizePoints(enum QPageSize::PageSizeId)
func (this *QPageSize) SizePoints_1(pageSizeId int) *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize10sizePointsENS_10PageSizeIdE", ffiqt.FFI_TYPE_POINTER, pageSizeId)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPageSize_SizePoints_1(pageSizeId int) *qtcore.QSize /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.SizePoints_1(pageSizeId)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:265
// index:0
// Public
// QSize sizePixels(int)
func (this *QPageSize) SizePixels(resolution int) *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10sizePixelsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:287
// index:1
// Public static
// QSize sizePixels(enum QPageSize::PageSizeId, int)
func (this *QPageSize) SizePixels_1(pageSizeId int, resolution int) *qtcore.QSize /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN9QPageSize10sizePixelsENS_10PageSizeIdEi", ffiqt.FFI_TYPE_POINTER, pageSizeId, resolution)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}
func QPageSize_SizePixels_1(pageSizeId int, resolution int) *qtcore.QSize /*123*/ {
	var nilthis *QPageSize
	rv := nilthis.SizePixels_1(pageSizeId, resolution)
	return rv
}

// /usr/include/qt/QtGui/qpagesize.h:267
// index:0
// Public
// QRectF rect(enum QPageSize::Unit)
func (this *QPageSize) Rect(units int) *qtcore.QRectF /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize4rectENS_4UnitE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &units)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:268
// index:0
// Public
// QRect rectPoints()
func (this *QPageSize) RectPoints() *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10rectPointsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:269
// index:0
// Public
// QRect rectPixels(int)
func (this *QPageSize) RectPixels(resolution int) *qtcore.QRect /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK9QPageSize10rectPixelsEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &resolution)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

//  body block end
