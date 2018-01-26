package qtgui

// /usr/include/qt/QtGui/qscreen.h
// #include <qscreen.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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
type QScreen struct {
	*qtcore.QObject
}

func (this *QScreen) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QScreen) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQScreenFromPointer(cthis unsafe.Pointer) *QScreen {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QScreen{bcthis0}
}
func (*QScreen) NewFromPointer(cthis unsafe.Pointer) *QScreen {
	return NewQScreenFromPointer(cthis)
}

// /usr/include/qt/QtGui/qscreen.h:68
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QScreen) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:98
// index:0
// Public virtual
// void ~QScreen()
func DeleteQScreen(*QScreen) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreenD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:101
// index:0
// Public
// QString name()
func (this *QScreen) Name() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen4nameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:103
// index:0
// Public
// QString manufacturer()
func (this *QScreen) Manufacturer() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12manufacturerEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:104
// index:0
// Public
// QString model()
func (this *QScreen) Model() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen5modelEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:105
// index:0
// Public
// QString serialNumber()
func (this *QScreen) SerialNumber() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12serialNumberEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:107
// index:0
// Public
// int depth()
func (this *QScreen) Depth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen5depthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qscreen.h:109
// index:0
// Public
// QSize size()
func (this *QScreen) Size() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:110
// index:0
// Public
// QRect geometry()
func (this *QScreen) Geometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen8geometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:112
// index:0
// Public
// QSizeF physicalSize()
func (this *QScreen) PhysicalSize() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12physicalSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:114
// index:0
// Public
// qreal physicalDotsPerInchX()
func (this *QScreen) PhysicalDotsPerInchX() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen20physicalDotsPerInchXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:115
// index:0
// Public
// qreal physicalDotsPerInchY()
func (this *QScreen) PhysicalDotsPerInchY() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen20physicalDotsPerInchYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:116
// index:0
// Public
// qreal physicalDotsPerInch()
func (this *QScreen) PhysicalDotsPerInch() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen19physicalDotsPerInchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:118
// index:0
// Public
// qreal logicalDotsPerInchX()
func (this *QScreen) LogicalDotsPerInchX() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen19logicalDotsPerInchXEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:119
// index:0
// Public
// qreal logicalDotsPerInchY()
func (this *QScreen) LogicalDotsPerInchY() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen19logicalDotsPerInchYEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:120
// index:0
// Public
// qreal logicalDotsPerInch()
func (this *QScreen) LogicalDotsPerInch() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen18logicalDotsPerInchEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:122
// index:0
// Public
// qreal devicePixelRatio()
func (this *QScreen) DevicePixelRatio() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen16devicePixelRatioEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:124
// index:0
// Public
// QSize availableSize()
func (this *QScreen) AvailableSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen13availableSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:125
// index:0
// Public
// QRect availableGeometry()
func (this *QScreen) AvailableGeometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen17availableGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:129
// index:0
// Public
// QSize virtualSize()
func (this *QScreen) VirtualSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11virtualSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:130
// index:0
// Public
// QRect virtualGeometry()
func (this *QScreen) VirtualGeometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen15virtualGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:132
// index:0
// Public
// QSize availableVirtualSize()
func (this *QScreen) AvailableVirtualSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen20availableVirtualSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:133
// index:0
// Public
// QRect availableVirtualGeometry()
func (this *QScreen) AvailableVirtualGeometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen24availableVirtualGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:135
// index:0
// Public
// Qt::ScreenOrientation primaryOrientation()
func (this *QScreen) PrimaryOrientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen18primaryOrientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:136
// index:0
// Public
// Qt::ScreenOrientation orientation()
func (this *QScreen) Orientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11orientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:137
// index:0
// Public
// Qt::ScreenOrientation nativeOrientation()
func (this *QScreen) NativeOrientation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen17nativeOrientationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:139
// index:0
// Public
// Qt::ScreenOrientations orientationUpdateMask()
func (this *QScreen) OrientationUpdateMask() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen21orientationUpdateMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:142
// index:0
// Public
// int angleBetween(Qt::ScreenOrientation, Qt::ScreenOrientation)
func (this *QScreen) AngleBetween(a int, b int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12angleBetweenEN2Qt17ScreenOrientationES1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), a, b)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtGui/qscreen.h:143
// index:0
// Public
// QTransform transformBetween(Qt::ScreenOrientation, Qt::ScreenOrientation, const class QRect &)
func (this *QScreen) TransformBetween(a int, b int, target *qtcore.QRect) *QTransform /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = target.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen16transformBetweenEN2Qt17ScreenOrientationES1_RK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), a, b, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:144
// index:0
// Public
// QRect mapBetween(Qt::ScreenOrientation, Qt::ScreenOrientation, const class QRect &)
func (this *QScreen) MapBetween(a int, b int, rect *qtcore.QRect) *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg2 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen10mapBetweenEN2Qt17ScreenOrientationES1_RK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), a, b, convArg2)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:146
// index:0
// Public
// bool isPortrait(Qt::ScreenOrientation)
func (this *QScreen) IsPortrait(orientation int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen10isPortraitEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qscreen.h:147
// index:0
// Public
// bool isLandscape(Qt::ScreenOrientation)
func (this *QScreen) IsLandscape(orientation int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11isLandscapeEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// Public
// QPixmap grabWindow(WId, int, int, int, int)
func (this *QScreen) GrabWindow(window uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), window, x, y, w, h)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:151
// index:0
// Public
// qreal refreshRate()
func (this *QScreen) RefreshRate() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11refreshRateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtGui/qscreen.h:154
// index:0
// Public
// void geometryChanged(const class QRect &)
func (this *QScreen) GeometryChanged(geometry *qtcore.QRect) {
	var convArg0 = geometry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen15geometryChangedERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:155
// index:0
// Public
// void availableGeometryChanged(const class QRect &)
func (this *QScreen) AvailableGeometryChanged(geometry *qtcore.QRect) {
	var convArg0 = geometry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen24availableGeometryChangedERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:156
// index:0
// Public
// void physicalSizeChanged(const class QSizeF &)
func (this *QScreen) PhysicalSizeChanged(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen19physicalSizeChangedERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:157
// index:0
// Public
// void physicalDotsPerInchChanged(qreal)
func (this *QScreen) PhysicalDotsPerInchChanged(dpi float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen26physicalDotsPerInchChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dpi)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:158
// index:0
// Public
// void logicalDotsPerInchChanged(qreal)
func (this *QScreen) LogicalDotsPerInchChanged(dpi float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen25logicalDotsPerInchChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dpi)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:159
// index:0
// Public
// void virtualGeometryChanged(const class QRect &)
func (this *QScreen) VirtualGeometryChanged(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen22virtualGeometryChangedERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:160
// index:0
// Public
// void primaryOrientationChanged(Qt::ScreenOrientation)
func (this *QScreen) PrimaryOrientationChanged(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen25primaryOrientationChangedEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:161
// index:0
// Public
// void orientationChanged(Qt::ScreenOrientation)
func (this *QScreen) OrientationChanged(orientation int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen18orientationChangedEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:162
// index:0
// Public
// void refreshRateChanged(qreal)
func (this *QScreen) RefreshRateChanged(refreshRate float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen18refreshRateChangedEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), refreshRate)
	gopp.ErrPrint(err, rv)
}

//  body block end
