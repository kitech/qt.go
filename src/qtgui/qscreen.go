//  header block begin
// /usr/include/qt/QtGui/qscreen.h
// #include <qscreen.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 3
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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qscreen.h:68
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QScreen) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:98
// index:0
// virtual
// void ~QScreen()
func DeleteQScreen(*QScreen) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreenD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:99
// index:0
// QPlatformScreen * handle()
func (this *QScreen) Handle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen6handleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:101
// index:0
// QString name()
func (this *QScreen) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:103
// index:0
// QString manufacturer()
func (this *QScreen) Manufacturer() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12manufacturerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:104
// index:0
// QString model()
func (this *QScreen) Model() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen5modelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:105
// index:0
// QString serialNumber()
func (this *QScreen) SerialNumber() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12serialNumberEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:107
// index:0
// int depth()
func (this *QScreen) Depth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen5depthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:109
// index:0
// QSize size()
func (this *QScreen) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:110
// index:0
// QRect geometry()
func (this *QScreen) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen8geometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:112
// index:0
// QSizeF physicalSize()
func (this *QScreen) PhysicalSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12physicalSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:114
// index:0
// qreal physicalDotsPerInchX()
func (this *QScreen) PhysicalDotsPerInchX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen20physicalDotsPerInchXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:115
// index:0
// qreal physicalDotsPerInchY()
func (this *QScreen) PhysicalDotsPerInchY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen20physicalDotsPerInchYEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:116
// index:0
// qreal physicalDotsPerInch()
func (this *QScreen) PhysicalDotsPerInch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen19physicalDotsPerInchEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:118
// index:0
// qreal logicalDotsPerInchX()
func (this *QScreen) LogicalDotsPerInchX() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen19logicalDotsPerInchXEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:119
// index:0
// qreal logicalDotsPerInchY()
func (this *QScreen) LogicalDotsPerInchY() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen19logicalDotsPerInchYEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:120
// index:0
// qreal logicalDotsPerInch()
func (this *QScreen) LogicalDotsPerInch() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen18logicalDotsPerInchEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:122
// index:0
// qreal devicePixelRatio()
func (this *QScreen) DevicePixelRatio() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen16devicePixelRatioEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:124
// index:0
// QSize availableSize()
func (this *QScreen) AvailableSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen13availableSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:125
// index:0
// QRect availableGeometry()
func (this *QScreen) AvailableGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen17availableGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:127
// index:0
// QList<QScreen *> virtualSiblings()
func (this *QScreen) VirtualSiblings() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen15virtualSiblingsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:129
// index:0
// QSize virtualSize()
func (this *QScreen) VirtualSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11virtualSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:130
// index:0
// QRect virtualGeometry()
func (this *QScreen) VirtualGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen15virtualGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:132
// index:0
// QSize availableVirtualSize()
func (this *QScreen) AvailableVirtualSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen20availableVirtualSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:133
// index:0
// QRect availableVirtualGeometry()
func (this *QScreen) AvailableVirtualGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen24availableVirtualGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:135
// index:0
// Qt::ScreenOrientation primaryOrientation()
func (this *QScreen) PrimaryOrientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen18primaryOrientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:136
// index:0
// Qt::ScreenOrientation orientation()
func (this *QScreen) Orientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11orientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:137
// index:0
// Qt::ScreenOrientation nativeOrientation()
func (this *QScreen) NativeOrientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen17nativeOrientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:139
// index:0
// Qt::ScreenOrientations orientationUpdateMask()
func (this *QScreen) OrientationUpdateMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen21orientationUpdateMaskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:142
// index:0
// int angleBetween(Qt::ScreenOrientation, Qt::ScreenOrientation)
func (this *QScreen) AngleBetween(a int, b int) {
	// 0: (, Qt::ScreenOrientation a, Qt::ScreenOrientation b), (&a, &b)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen12angleBetweenEN2Qt17ScreenOrientationES1_", ffiqt.FFI_TYPE_VOID, this.cthis, &a, &b)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:143
// index:0
// QTransform transformBetween(Qt::ScreenOrientation, Qt::ScreenOrientation, const class QRect &)
func (this *QScreen) TransformBetween(a int, b int, target unsafe.Pointer) {
	// 0: (, Qt::ScreenOrientation a, Qt::ScreenOrientation b, const QRect & target), (&a, &b, target)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen16transformBetweenEN2Qt17ScreenOrientationES1_RK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, &a, &b, target)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:144
// index:0
// QRect mapBetween(Qt::ScreenOrientation, Qt::ScreenOrientation, const class QRect &)
func (this *QScreen) MapBetween(a int, b int, rect unsafe.Pointer) {
	// 0: (, Qt::ScreenOrientation a, Qt::ScreenOrientation b, const QRect & rect), (&a, &b, rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen10mapBetweenEN2Qt17ScreenOrientationES1_RK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, &a, &b, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:146
// index:0
// bool isPortrait(Qt::ScreenOrientation)
func (this *QScreen) IsPortrait(orientation int) {
	// 0: (, Qt::ScreenOrientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen10isPortraitEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:147
// index:0
// bool isLandscape(Qt::ScreenOrientation)
func (this *QScreen) IsLandscape(orientation int) {
	// 0: (, Qt::ScreenOrientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11isLandscapeEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// QPixmap grabWindow(WId, int, int, int, int)
func (this *QScreen) GrabWindow(window uint64, x int, y int, w int, h int) {
	// 0: (, WId window, int x, int y, int w, int h), (&window, &x, &y, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &window, &x, &y, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:151
// index:0
// qreal refreshRate()
func (this *QScreen) RefreshRate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QScreen11refreshRateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:154
// index:0
// void geometryChanged(const class QRect &)
func (this *QScreen) GeometryChanged(geometry unsafe.Pointer) {
	// 0: (, const QRect & geometry), (geometry)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen15geometryChangedERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, geometry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:155
// index:0
// void availableGeometryChanged(const class QRect &)
func (this *QScreen) AvailableGeometryChanged(geometry unsafe.Pointer) {
	// 0: (, const QRect & geometry), (geometry)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen24availableGeometryChangedERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, geometry)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:156
// index:0
// void physicalSizeChanged(const class QSizeF &)
func (this *QScreen) PhysicalSizeChanged(size unsafe.Pointer) {
	// 0: (, const QSizeF & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen19physicalSizeChangedERK6QSizeF", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:157
// index:0
// void physicalDotsPerInchChanged(qreal)
func (this *QScreen) PhysicalDotsPerInchChanged(dpi float64) {
	// 0: (, qreal dpi), (&dpi)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen26physicalDotsPerInchChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &dpi)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:158
// index:0
// void logicalDotsPerInchChanged(qreal)
func (this *QScreen) LogicalDotsPerInchChanged(dpi float64) {
	// 0: (, qreal dpi), (&dpi)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen25logicalDotsPerInchChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &dpi)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:159
// index:0
// void virtualGeometryChanged(const class QRect &)
func (this *QScreen) VirtualGeometryChanged(rect unsafe.Pointer) {
	// 0: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen22virtualGeometryChangedERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:160
// index:0
// void primaryOrientationChanged(Qt::ScreenOrientation)
func (this *QScreen) PrimaryOrientationChanged(orientation int) {
	// 0: (, Qt::ScreenOrientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen25primaryOrientationChangedEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:161
// index:0
// void orientationChanged(Qt::ScreenOrientation)
func (this *QScreen) OrientationChanged(orientation int) {
	// 0: (, Qt::ScreenOrientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen18orientationChangedEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:162
// index:0
// void refreshRateChanged(qreal)
func (this *QScreen) RefreshRateChanged(refreshRate float64) {
	// 0: (, qreal refreshRate), (&refreshRate)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QScreen18refreshRateChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &refreshRate)
	gopp.ErrPrint(err, rv)
}

//  body block end
