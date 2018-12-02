package qtgui

// /usr/include/qt/QtGui/qscreen.h
// #include <qscreen.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 5
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

/*

 */
type QScreen struct {
	*qtcore.QObject
}
type QScreen_ITF interface {
	qtcore.QObject_ITF
	QScreen_PTR() *QScreen
}

func (ptr *QScreen) QScreen_PTR() *QScreen { return ptr }

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
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QScreen) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qscreen.h:98
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QScreen()

/*

 */
func DeleteQScreen(this *QScreen) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreenD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qscreen.h:101
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name() const

/*

 */
func (this *QScreen) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qscreen.h:103
// index:0
// Public Visibility=Default Availability=Available
// [8] QString manufacturer() const

/*

 */
func (this *QScreen) Manufacturer() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen12manufacturerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qscreen.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] QString model() const

/*

 */
func (this *QScreen) Model() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen5modelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qscreen.h:105
// index:0
// Public Visibility=Default Availability=Available
// [8] QString serialNumber() const

/*

 */
func (this *QScreen) SerialNumber() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen12serialNumberEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qscreen.h:107
// index:0
// Public Visibility=Default Availability=Available
// [4] int depth() const

/*

 */
func (this *QScreen) Depth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen5depthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:109
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size() const

/*

 */
func (this *QScreen) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:110
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect geometry() const

/*

 */
func (this *QScreen) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:112
// index:0
// Public Visibility=Default Availability=Available
// [16] QSizeF physicalSize() const

/*

 */
func (this *QScreen) PhysicalSize() *qtcore.QSizeF /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen12physicalSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSizeF)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:114
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal physicalDotsPerInchX() const

/*

 */
func (this *QScreen) PhysicalDotsPerInchX() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen20physicalDotsPerInchXEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal physicalDotsPerInchY() const

/*

 */
func (this *QScreen) PhysicalDotsPerInchY() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen20physicalDotsPerInchYEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:116
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal physicalDotsPerInch() const

/*

 */
func (this *QScreen) PhysicalDotsPerInch() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen19physicalDotsPerInchEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:118
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal logicalDotsPerInchX() const

/*

 */
func (this *QScreen) LogicalDotsPerInchX() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen19logicalDotsPerInchXEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal logicalDotsPerInchY() const

/*

 */
func (this *QScreen) LogicalDotsPerInchY() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen19logicalDotsPerInchYEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal logicalDotsPerInch() const

/*

 */
func (this *QScreen) LogicalDotsPerInch() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen18logicalDotsPerInchEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio() const

/*

 */
func (this *QScreen) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:124
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize availableSize() const

/*

 */
func (this *QScreen) AvailableSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen13availableSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:125
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect availableGeometry() const

/*

 */
func (this *QScreen) AvailableGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen17availableGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize virtualSize() const

/*

 */
func (this *QScreen) VirtualSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen11virtualSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:130
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect virtualGeometry() const

/*

 */
func (this *QScreen) VirtualGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen15virtualGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:132
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize availableVirtualSize() const

/*

 */
func (this *QScreen) AvailableVirtualSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen20availableVirtualSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:133
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect availableVirtualGeometry() const

/*

 */
func (this *QScreen) AvailableVirtualGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen24availableVirtualGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:135
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation primaryOrientation() const

/*

 */
func (this *QScreen) PrimaryOrientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen18primaryOrientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:136
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation orientation() const

/*

 */
func (this *QScreen) Orientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen11orientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:137
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation nativeOrientation() const

/*

 */
func (this *QScreen) NativeOrientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen17nativeOrientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:139
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientations orientationUpdateMask() const

/*
Returns the currently set orientation update mask.

See also setOrientationUpdateMask().
*/
func (this *QScreen) OrientationUpdateMask() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen21orientationUpdateMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qscreen.h:140
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOrientationUpdateMask(Qt::ScreenOrientations)

/*
Sets the orientations that the application is interested in receiving updates for in conjunction with this screen.

For example, to receive orientation() updates and thus have orientationChanged() signals being emitted for LandscapeOrientation and InvertedLandscapeOrientation, call setOrientationUpdateMask() with mask set to Qt::LandscapeOrientation | Qt::InvertedLandscapeOrientation.

The default, 0, means no orientationChanged() signals are fired.

See also orientationUpdateMask().
*/
func (this *QScreen) SetOrientationUpdateMask(mask int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen24setOrientationUpdateMaskE6QFlagsIN2Qt17ScreenOrientationEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mask)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:142
// index:0
// Public Visibility=Default Availability=Available
// [4] int angleBetween(Qt::ScreenOrientation, Qt::ScreenOrientation) const

/*
Convenience function to compute the angle of rotation to get from rotation a to rotation b.

The result will be 0, 90, 180, or 270.

Qt::PrimaryOrientation is interpreted as the screen's primaryOrientation().
*/
func (this *QScreen) AngleBetween(a int, b int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen12angleBetweenEN2Qt17ScreenOrientationES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, b)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:143
// index:0
// Public Visibility=Default Availability=Available
// [88] QTransform transformBetween(Qt::ScreenOrientation, Qt::ScreenOrientation, const QRect &) const

/*
Convenience function to compute a transform that maps from the coordinate system defined by orientation a into the coordinate system defined by orientation b and target dimensions target.

Example, a is Qt::Landscape, b is Qt::Portrait, and target is QRect(0, 0, w, h) the resulting transform will be such that the point QPoint(0, 0) is mapped to QPoint(0, w), and QPoint(h, w) is mapped to QPoint(0, h). Thus, the landscape coordinate system QRect(0, 0, h, w) is mapped (with a 90 degree rotation) into the portrait coordinate system QRect(0, 0, w, h).

Qt::PrimaryOrientation is interpreted as the screen's primaryOrientation().
*/
func (this *QScreen) TransformBetween(a int, b int, target qtcore.QRect_ITF) *QTransform /*123*/ {
	var convArg2 unsafe.Pointer
	if target != nil && target.QRect_PTR() != nil {
		convArg2 = target.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen16transformBetweenEN2Qt17ScreenOrientationES1_RK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, b, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTransformFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTransform)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:144
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect mapBetween(Qt::ScreenOrientation, Qt::ScreenOrientation, const QRect &) const

/*
Maps the rect between two screen orientations.

This will flip the x and y dimensions of the rectangle rect if the orientation a is Qt::PortraitOrientation or Qt::InvertedPortraitOrientation and orientation b is Qt::LandscapeOrientation or Qt::InvertedLandscapeOrientation, or vice versa.

Qt::PrimaryOrientation is interpreted as the screen's primaryOrientation().
*/
func (this *QScreen) MapBetween(a int, b int, rect qtcore.QRect_ITF) *qtcore.QRect /*123*/ {
	var convArg2 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg2 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen10mapBetweenEN2Qt17ScreenOrientationES1_RK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), a, b, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:146
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isPortrait(Qt::ScreenOrientation) const

/*
Convenience function that returns true if o is either portrait or inverted portrait; otherwise returns false.

Qt::PrimaryOrientation is interpreted as the screen's primaryOrientation().
*/
func (this *QScreen) IsPortrait(orientation int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen10isPortraitEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qscreen.h:147
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLandscape(Qt::ScreenOrientation) const

/*
Convenience function that returns true if o is either landscape or inverted landscape; otherwise returns false.

Qt::PrimaryOrientation is interpreted as the screen's primaryOrientation().
*/
func (this *QScreen) IsLandscape(orientation int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen11isLandscapeEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*
Creates and returns a pixmap constructed by grabbing the contents of the given window restricted by QRect(x, y, width, height).

The arguments (x, y) specify the offset in the window, whereas (width, height) specify the area to be copied. If width is negative, the function copies everything to the right border of the window. If height is negative, the function copies everything to the bottom of the window.

The window system identifier (WId) can be retrieved using the QWidget::winId() function. The rationale for using a window identifier and not a QWidget, is to enable grabbing of windows that are not part of the application, window system frames, and so on.

Warning: Grabbing windows that are not part of the application is not supported on systems such as iOS, where sandboxing/security prevents reading pixels of windows not owned by the application.

The grabWindow() function grabs pixels from the screen, not from the window, i.e. if there is another window partially or entirely over the one you grab, you get pixels from the overlying window, too. The mouse cursor is generally not grabbed.

Note on X11 that if the given window doesn't have the same depth as the root window, and another window partially or entirely obscures the one you grab, you will not get pixels from the overlying window. The contents of the obscured areas in the pixmap will be undefined and uninitialized.

On Windows Vista and above grabbing a layered window, which is created by setting the Qt::WA_TranslucentBackground attribute, will not work. Instead grabbing the desktop widget should work.

Warning: In general, grabbing an area outside the screen is not safe. This depends on the underlying window system.
*/
func (this *QScreen) GrabWindow(window uint64, x int, y int, w int, h int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), window, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*
Creates and returns a pixmap constructed by grabbing the contents of the given window restricted by QRect(x, y, width, height).

The arguments (x, y) specify the offset in the window, whereas (width, height) specify the area to be copied. If width is negative, the function copies everything to the right border of the window. If height is negative, the function copies everything to the bottom of the window.

The window system identifier (WId) can be retrieved using the QWidget::winId() function. The rationale for using a window identifier and not a QWidget, is to enable grabbing of windows that are not part of the application, window system frames, and so on.

Warning: Grabbing windows that are not part of the application is not supported on systems such as iOS, where sandboxing/security prevents reading pixels of windows not owned by the application.

The grabWindow() function grabs pixels from the screen, not from the window, i.e. if there is another window partially or entirely over the one you grab, you get pixels from the overlying window, too. The mouse cursor is generally not grabbed.

Note on X11 that if the given window doesn't have the same depth as the root window, and another window partially or entirely obscures the one you grab, you will not get pixels from the overlying window. The contents of the obscured areas in the pixmap will be undefined and uninitialized.

On Windows Vista and above grabbing a layered window, which is created by setting the Qt::WA_TranslucentBackground attribute, will not work. Instead grabbing the desktop widget should work.

Warning: In general, grabbing an area outside the screen is not safe. This depends on the underlying window system.
*/
func (this *QScreen) GrabWindow__(window uint64) *QPixmap /*123*/ {
	// arg: 1, int=Int, =Invalid, , Invalid
	x := int(0)
	// arg: 2, int=Int, =Invalid, , Invalid
	y := int(0)
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), window, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*
Creates and returns a pixmap constructed by grabbing the contents of the given window restricted by QRect(x, y, width, height).

The arguments (x, y) specify the offset in the window, whereas (width, height) specify the area to be copied. If width is negative, the function copies everything to the right border of the window. If height is negative, the function copies everything to the bottom of the window.

The window system identifier (WId) can be retrieved using the QWidget::winId() function. The rationale for using a window identifier and not a QWidget, is to enable grabbing of windows that are not part of the application, window system frames, and so on.

Warning: Grabbing windows that are not part of the application is not supported on systems such as iOS, where sandboxing/security prevents reading pixels of windows not owned by the application.

The grabWindow() function grabs pixels from the screen, not from the window, i.e. if there is another window partially or entirely over the one you grab, you get pixels from the overlying window, too. The mouse cursor is generally not grabbed.

Note on X11 that if the given window doesn't have the same depth as the root window, and another window partially or entirely obscures the one you grab, you will not get pixels from the overlying window. The contents of the obscured areas in the pixmap will be undefined and uninitialized.

On Windows Vista and above grabbing a layered window, which is created by setting the Qt::WA_TranslucentBackground attribute, will not work. Instead grabbing the desktop widget should work.

Warning: In general, grabbing an area outside the screen is not safe. This depends on the underlying window system.
*/
func (this *QScreen) GrabWindow__1(window uint64, x int) *QPixmap /*123*/ {
	// arg: 2, int=Int, =Invalid, , Invalid
	y := int(0)
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), window, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*
Creates and returns a pixmap constructed by grabbing the contents of the given window restricted by QRect(x, y, width, height).

The arguments (x, y) specify the offset in the window, whereas (width, height) specify the area to be copied. If width is negative, the function copies everything to the right border of the window. If height is negative, the function copies everything to the bottom of the window.

The window system identifier (WId) can be retrieved using the QWidget::winId() function. The rationale for using a window identifier and not a QWidget, is to enable grabbing of windows that are not part of the application, window system frames, and so on.

Warning: Grabbing windows that are not part of the application is not supported on systems such as iOS, where sandboxing/security prevents reading pixels of windows not owned by the application.

The grabWindow() function grabs pixels from the screen, not from the window, i.e. if there is another window partially or entirely over the one you grab, you get pixels from the overlying window, too. The mouse cursor is generally not grabbed.

Note on X11 that if the given window doesn't have the same depth as the root window, and another window partially or entirely obscures the one you grab, you will not get pixels from the overlying window. The contents of the obscured areas in the pixmap will be undefined and uninitialized.

On Windows Vista and above grabbing a layered window, which is created by setting the Qt::WA_TranslucentBackground attribute, will not work. Instead grabbing the desktop widget should work.

Warning: In general, grabbing an area outside the screen is not safe. This depends on the underlying window system.
*/
func (this *QScreen) GrabWindow__2(window uint64, x int, y int) *QPixmap /*123*/ {
	// arg: 3, int=Int, =Invalid, , Invalid
	w := int(-1)
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), window, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:149
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grabWindow(WId, int, int, int, int)

/*
Creates and returns a pixmap constructed by grabbing the contents of the given window restricted by QRect(x, y, width, height).

The arguments (x, y) specify the offset in the window, whereas (width, height) specify the area to be copied. If width is negative, the function copies everything to the right border of the window. If height is negative, the function copies everything to the bottom of the window.

The window system identifier (WId) can be retrieved using the QWidget::winId() function. The rationale for using a window identifier and not a QWidget, is to enable grabbing of windows that are not part of the application, window system frames, and so on.

Warning: Grabbing windows that are not part of the application is not supported on systems such as iOS, where sandboxing/security prevents reading pixels of windows not owned by the application.

The grabWindow() function grabs pixels from the screen, not from the window, i.e. if there is another window partially or entirely over the one you grab, you get pixels from the overlying window, too. The mouse cursor is generally not grabbed.

Note on X11 that if the given window doesn't have the same depth as the root window, and another window partially or entirely obscures the one you grab, you will not get pixels from the overlying window. The contents of the obscured areas in the pixmap will be undefined and uninitialized.

On Windows Vista and above grabbing a layered window, which is created by setting the Qt::WA_TranslucentBackground attribute, will not work. Instead grabbing the desktop widget should work.

Warning: In general, grabbing an area outside the screen is not safe. This depends on the underlying window system.
*/
func (this *QScreen) GrabWindow__3(window uint64, x int, y int, w int) *QPixmap /*123*/ {
	// arg: 4, int=Int, =Invalid, , Invalid
	h := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen10grabWindowEyiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), window, x, y, w, h)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qscreen.h:151
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal refreshRate() const

/*

 */
func (this *QScreen) RefreshRate() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QScreen11refreshRateEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qscreen.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void geometryChanged(const QRect &)

/*

 */
func (this *QScreen) GeometryChanged(geometry qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if geometry != nil && geometry.QRect_PTR() != nil {
		convArg0 = geometry.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen15geometryChangedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:155
// index:0
// Public Visibility=Default Availability=Available
// [-2] void availableGeometryChanged(const QRect &)

/*

 */
func (this *QScreen) AvailableGeometryChanged(geometry qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if geometry != nil && geometry.QRect_PTR() != nil {
		convArg0 = geometry.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen24availableGeometryChangedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void physicalSizeChanged(const QSizeF &)

/*

 */
func (this *QScreen) PhysicalSizeChanged(size qtcore.QSizeF_ITF) {
	var convArg0 unsafe.Pointer
	if size != nil && size.QSizeF_PTR() != nil {
		convArg0 = size.QSizeF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen19physicalSizeChangedERK6QSizeF", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:157
// index:0
// Public Visibility=Default Availability=Available
// [-2] void physicalDotsPerInchChanged(qreal)

/*

 */
func (this *QScreen) PhysicalDotsPerInchChanged(dpi float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen26physicalDotsPerInchChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dpi)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:158
// index:0
// Public Visibility=Default Availability=Available
// [-2] void logicalDotsPerInchChanged(qreal)

/*

 */
func (this *QScreen) LogicalDotsPerInchChanged(dpi float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen25logicalDotsPerInchChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dpi)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void virtualGeometryChanged(const QRect &)

/*

 */
func (this *QScreen) VirtualGeometryChanged(rect qtcore.QRect_ITF) {
	var convArg0 unsafe.Pointer
	if rect != nil && rect.QRect_PTR() != nil {
		convArg0 = rect.QRect_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen22virtualGeometryChangedERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void primaryOrientationChanged(Qt::ScreenOrientation)

/*
This signal is emitted when the primary orientation of the screen changes with orientation as an argument.

Note: Notifier signal for property primaryOrientation.

See also primaryOrientation().
*/
func (this *QScreen) PrimaryOrientationChanged(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen25primaryOrientationChangedEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:161
// index:0
// Public Visibility=Default Availability=Available
// [-2] void orientationChanged(Qt::ScreenOrientation)

/*
This signal is emitted when the orientation of the screen changes with orientation as an argument.

Note: Notifier signal for property orientation.

See also orientation().
*/
func (this *QScreen) OrientationChanged(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen18orientationChangedEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qscreen.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void refreshRateChanged(qreal)

/*

 */
func (this *QScreen) RefreshRateChanged(refreshRate float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QScreen18refreshRateChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), refreshRate)
	qtrt.ErrPrint(err, rv)
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
