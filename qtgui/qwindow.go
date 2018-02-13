package qtgui

// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

// void exposeEvent(class QExposeEvent *)
func (this *QWindow) InheritExposeEvent(f func(arg0 *QExposeEvent /*777 QExposeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "exposeEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QWindow) InheritResizeEvent(f func(arg0 *QResizeEvent /*777 QResizeEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void moveEvent(class QMoveEvent *)
func (this *QWindow) InheritMoveEvent(f func(arg0 *QMoveEvent /*777 QMoveEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QWindow) InheritFocusInEvent(f func(arg0 *QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QWindow) InheritFocusOutEvent(f func(arg0 *QFocusEvent /*777 QFocusEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QWindow) InheritShowEvent(f func(arg0 *QShowEvent /*777 QShowEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QWindow) InheritHideEvent(f func(arg0 *QHideEvent /*777 QHideEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// bool event(class QEvent *)
func (this *QWindow) InheritEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QWindow) InheritKeyPressEvent(f func(arg0 *QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QWindow) InheritKeyReleaseEvent(f func(arg0 *QKeyEvent /*777 QKeyEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QWindow) InheritMousePressEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QWindow) InheritMouseReleaseEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QWindow) InheritMouseDoubleClickEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QWindow) InheritMouseMoveEvent(f func(arg0 *QMouseEvent /*777 QMouseEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QWindow) InheritWheelEvent(f func(arg0 *QWheelEvent /*777 QWheelEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void touchEvent(class QTouchEvent *)
func (this *QWindow) InheritTouchEvent(f func(arg0 *QTouchEvent /*777 QTouchEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "touchEvent", f)
}

// void tabletEvent(class QTabletEvent *)
func (this *QWindow) InheritTabletEvent(f func(arg0 *QTabletEvent /*777 QTabletEvent **/) /*void*/) {
	qtrt.SetAllInheritCallback(this, "tabletEvent", f)
}

// bool nativeEvent(const class QByteArray &, void *, long *)
func (this *QWindow) InheritNativeEvent(f func(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool) {
	qtrt.SetAllInheritCallback(this, "nativeEvent", f)
}

type QWindow struct {
	*qtcore.QObject
	*QSurface
}
type QWindow_ITF interface {
	qtcore.QObject_ITF
	QSurface_ITF
	QWindow_PTR() *QWindow
}

func (ptr *QWindow) QWindow_PTR() *QWindow { return ptr }

func (this *QWindow) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWindow) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QSurface = NewQSurfaceFromPointer(cthis)
}
func NewQWindowFromPointer(cthis unsafe.Pointer) *QWindow {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := NewQSurfaceFromPointer(cthis)
	return &QWindow{bcthis0, bcthis1}
}
func (*QWindow) NewFromPointer(cthis unsafe.Pointer) *QWindow {
	return NewQWindowFromPointer(cthis)
}

// /usr/include/qt/QtGui/qwindow.h:97
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWindow) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:144
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWindow(QScreen *)
func NewQWindow(screen *QScreen /*777 QScreen **/) *QWindow {
	var convArg0 = screen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowC2EP7QScreen", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:145
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWindow(QWindow *)
func NewQWindow_1(parent *QWindow /*777 QWindow **/) *QWindow {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowC2EPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWindowFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qwindow.h:146
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWindow()
func DeleteQWindow(this *QWindow) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindowD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 40)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qwindow.h:148
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSurfaceType(enum QSurface::SurfaceType)
func (this *QWindow) SetSurfaceType(surfaceType int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setSurfaceTypeEN8QSurface11SurfaceTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), surfaceType)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:149
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] QSurface::SurfaceType surfaceType()
func (this *QWindow) SurfaceType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11surfaceTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:151
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QWindow) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:153
// index:0
// Public Visibility=Default Availability=Available
// [4] QWindow::Visibility visibility()
func (this *QWindow) Visibility() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow10visibilityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisibility(enum QWindow::Visibility)
func (this *QWindow) SetVisibility(v int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13setVisibilityENS_10VisibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), v)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:156
// index:0
// Public Visibility=Default Availability=Available
// [-2] void create()
func (this *QWindow) Create() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow6createEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:158
// index:0
// Public Visibility=Default Availability=Available
// [8] WId winId()
func (this *QWindow) WinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5winIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtGui/qwindow.h:160
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * parent(enum QWindow::AncestorMode)
func (this *QWindow) Parent(mode int) *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6parentENS_12AncestorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mode)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:161
// index:1
// Public Visibility=Default Availability=Available
// [8] QWindow * parent()
func (this *QWindow) Parent_1() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6parentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParent(QWindow *)
func (this *QWindow) SetParent(parent *QWindow /*777 QWindow **/) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:164
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTopLevel()
func (this *QWindow) IsTopLevel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow10isTopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:166
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModal()
func (this *QWindow) IsModal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow7isModalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowModality modality()
func (this *QWindow) Modality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8modalityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:168
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setModality(Qt::WindowModality)
func (this *QWindow) SetModality(modality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setModalityEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:170
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFormat(const QSurfaceFormat &)
func (this *QWindow) SetFormat(format *QSurfaceFormat) {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setFormatERK14QSurfaceFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:171
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSurfaceFormat format()
func (this *QWindow) Format() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6formatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] QSurfaceFormat requestedFormat()
func (this *QWindow) RequestedFormat() *QSurfaceFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow15requestedFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQSurfaceFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSurfaceFormat)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:174
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(Qt::WindowFlags)
func (this *QWindow) SetFlags(flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8setFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:175
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowFlags flags()
func (this *QWindow) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(Qt::WindowType, _Bool)
func (this *QWindow) SetFlag(arg0 int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setFlagEN2Qt10WindowTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:177
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowType type()
func (this *QWindow) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:179
// index:0
// Public Visibility=Default Availability=Available
// [8] QString title()
func (this *QWindow) Title() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5titleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qwindow.h:181
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setOpacity(qreal)
func (this *QWindow) SetOpacity(level float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10setOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:182
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal opacity()
func (this *QWindow) Opacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow7opacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:184
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QRegion &)
func (this *QWindow) SetMask(region *QRegion) {
	var convArg0 = region.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setMaskERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion mask()
func (this *QWindow) Mask() *QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:187
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActive()
func (this *QWindow) IsActive() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8isActiveEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void reportContentOrientationChange(Qt::ScreenOrientation)
func (this *QWindow) ReportContentOrientationChange(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow30reportContentOrientationChangeEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:190
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ScreenOrientation contentOrientation()
func (this *QWindow) ContentOrientation() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow18contentOrientationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:192
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal devicePixelRatio()
func (this *QWindow) DevicePixelRatio() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow16devicePixelRatioEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:194
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowState windowState()
func (this *QWindow) WindowState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11windowStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:195
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowStates windowStates()
func (this *QWindow) WindowStates() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12windowStatesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qwindow.h:196
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowState(Qt::WindowState)
func (this *QWindow) SetWindowState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setWindowStateEN2Qt11WindowStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:197
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowStates(Qt::WindowStates)
func (this *QWindow) SetWindowStates(states int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15setWindowStatesE6QFlagsIN2Qt11WindowStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), states)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:199
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTransientParent(QWindow *)
func (this *QWindow) SetTransientParent(parent *QWindow /*777 QWindow **/) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18setTransientParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:200
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * transientParent()
func (this *QWindow) TransientParent() *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow15transientParentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:202
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QWindow *, enum QWindow::AncestorMode)
func (this *QWindow) IsAncestorOf(child *QWindow /*777 const QWindow **/, mode int) bool {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12isAncestorOfEPKS_NS_12AncestorModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:204
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isExposed()
func (this *QWindow) IsExposed() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow9isExposedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:206
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minimumWidth()
func (this *QWindow) MinimumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12minimumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:207
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int minimumHeight()
func (this *QWindow) MinimumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13minimumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int maximumWidth()
func (this *QWindow) MaximumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12maximumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:209
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int maximumHeight()
func (this *QWindow) MaximumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13maximumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:211
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QWindow) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:212
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QWindow) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:213
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize baseSize()
func (this *QWindow) BaseSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8baseSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:214
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeIncrement()
func (this *QWindow) SizeIncrement() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13sizeIncrementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:216
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSize &)
func (this *QWindow) SetMinimumSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setMinimumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:217
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSize &)
func (this *QWindow) SetMaximumSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14setMaximumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:218
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseSize(const QSize &)
func (this *QWindow) SetBaseSize(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setBaseSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeIncrement(const QSize &)
func (this *QWindow) SetSizeIncrement(size *qtcore.QSize) {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setSizeIncrementERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:221
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect geometry()
func (this *QWindow) Geometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:223
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins frameMargins()
func (this *QWindow) FrameMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow12frameMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:224
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameGeometry()
func (this *QWindow) FrameGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13frameGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:226
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint framePosition()
func (this *QWindow) FramePosition() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13framePositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFramePosition(const QPoint &)
func (this *QWindow) SetFramePosition(point *qtcore.QPoint) {
	var convArg0 = point.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setFramePositionERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:229
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width()
func (this *QWindow) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:230
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height()
func (this *QWindow) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:231
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int x()
func (this *QWindow) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:232
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int y()
func (this *QWindow) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qwindow.h:234
// index:0
// Public inline virtual Visibility=Default Availability=Available
// [8] QSize size()
func (this *QWindow) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:235
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QPoint position()
func (this *QWindow) Position() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8positionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:237
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPosition(const QPoint &)
func (this *QWindow) SetPosition(pt *qtcore.QPoint) {
	var convArg0 = pt.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setPositionERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:238
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setPosition(int, int)
func (this *QWindow) SetPosition_1(posx int, posy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setPositionEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), posx, posy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:240
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSize &)
func (this *QWindow) Resize(newSize *qtcore.QSize) {
	var convArg0 = newSize.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow6resizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:241
// index:1
// Public Visibility=Default Availability=Available
// [-2] void resize(int, int)
func (this *QWindow) Resize_1(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow6resizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFilePath(const QString &)
func (this *QWindow) SetFilePath(filePath string) {
	var tmpArg0 = qtcore.NewQString_5(filePath)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:244
// index:0
// Public Visibility=Default Availability=Available
// [8] QString filePath()
func (this *QWindow) FilePath() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow8filePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qwindow.h:246
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIcon(const QIcon &)
func (this *QWindow) SetIcon(icon *QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7setIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:247
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon icon()
func (this *QWindow) Icon() *QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow4iconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:249
// index:0
// Public Visibility=Default Availability=Available
// [-2] void destroy()
func (this *QWindow) Destroy() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow7destroyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:253
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setKeyboardGrabEnabled(_Bool)
func (this *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow22setKeyboardGrabEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), grab)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setMouseGrabEnabled(_Bool)
func (this *QWindow) SetMouseGrabEnabled(grab bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow19setMouseGrabEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), grab)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:256
// index:0
// Public Visibility=Default Availability=Available
// [8] QScreen * screen()
func (this *QWindow) Screen() *QScreen /*777 QScreen **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6screenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQScreenFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:257
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScreen(QScreen *)
func (this *QWindow) SetScreen(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setScreenEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:259
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAccessibleInterface * accessibleRoot()
func (this *QWindow) AccessibleRoot() *QAccessibleInterface /*777 QAccessibleInterface **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow14accessibleRootEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAccessibleInterfaceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:260
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QObject * focusObject()
func (this *QWindow) FocusObject() *qtcore.QObject /*777 QObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11focusObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qwindow.h:262
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapToGlobal(const QPoint &)
func (this *QWindow) MapToGlobal(pos *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow11mapToGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromGlobal(const QPoint &)
func (this *QWindow) MapFromGlobal(pos *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow13mapFromGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:266
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor()
func (this *QWindow) Cursor() *QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWindow6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtGui/qwindow.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)
func (this *QWindow) SetCursor(arg0 *QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:268
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()
func (this *QWindow) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:271
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWindow * fromWinId(WId)
func (this *QWindow) FromWinId(id uint64) *QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9fromWinIdEy", qtrt.FFI_TYPE_POINTER, id)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWindow_FromWinId(id uint64) *QWindow /*777 QWindow **/ {
	var nilthis *QWindow
	rv := nilthis.FromWinId(id)
	return rv
}

// /usr/include/qt/QtGui/qwindow.h:279
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestActivate()
func (this *QWindow) RequestActivate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15requestActivateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:281
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QWindow) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()
func (this *QWindow) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:284
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()
func (this *QWindow) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:286
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMinimized()
func (this *QWindow) ShowMinimized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13showMinimizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:287
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMaximized()
func (this *QWindow) ShowMaximized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13showMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:288
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showFullScreen()
func (this *QWindow) ShowFullScreen() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14showFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNormal()
func (this *QWindow) ShowNormal() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10showNormalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:291
// index:0
// Public Visibility=Default Availability=Available
// [1] bool close()
func (this *QWindow) Close() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void raise()
func (this *QWindow) Raise() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5raiseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:293
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lower()
func (this *QWindow) Lower() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5lowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:295
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTitle(const QString &)
func (this *QWindow) SetTitle(arg0 string) {
	var tmpArg0 = qtcore.NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8setTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setX(int)
func (this *QWindow) SetX(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4setXEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:298
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setY(int)
func (this *QWindow) SetY(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow4setYEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:299
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWidth(int)
func (this *QWindow) SetWidth(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8setWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:300
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHeight(int)
func (this *QWindow) SetHeight(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9setHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:301
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(int, int, int, int)
func (this *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setGeometryEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), posx, posy, w, h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:302
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QWindow) SetGeometry_1(rect *qtcore.QRect) {
	var convArg0 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:304
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(int)
func (this *QWindow) SetMinimumWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15setMinimumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:305
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(int)
func (this *QWindow) SetMinimumHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setMinimumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:306
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(int)
func (this *QWindow) SetMaximumWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15setMaximumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:307
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(int)
func (this *QWindow) SetMaximumHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow16setMaximumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:309
// index:0
// Public Visibility=Default Availability=Available
// [-2] void alert(int)
func (this *QWindow) Alert(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5alertEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void requestUpdate()
func (this *QWindow) RequestUpdate() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13requestUpdateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void screenChanged(QScreen *)
func (this *QWindow) ScreenChanged(screen *QScreen /*777 QScreen **/) {
	var convArg0 = screen.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13screenChangedEP7QScreen", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:315
// index:0
// Public Visibility=Default Availability=Available
// [-2] void modalityChanged(Qt::WindowModality)
func (this *QWindow) ModalityChanged(modality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15modalityChangedEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), modality)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:316
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowStateChanged(Qt::WindowState)
func (this *QWindow) WindowStateChanged(windowState int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18windowStateChangedEN2Qt11WindowStateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), windowState)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:317
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowTitleChanged(const QString &)
func (this *QWindow) WindowTitleChanged(title string) {
	var tmpArg0 = qtcore.NewQString_5(title)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18windowTitleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:319
// index:0
// Public Visibility=Default Availability=Available
// [-2] void xChanged(int)
func (this *QWindow) XChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8xChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:320
// index:0
// Public Visibility=Default Availability=Available
// [-2] void yChanged(int)
func (this *QWindow) YChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow8yChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:322
// index:0
// Public Visibility=Default Availability=Available
// [-2] void widthChanged(int)
func (this *QWindow) WidthChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow12widthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:323
// index:0
// Public Visibility=Default Availability=Available
// [-2] void heightChanged(int)
func (this *QWindow) HeightChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13heightChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:325
// index:0
// Public Visibility=Default Availability=Available
// [-2] void minimumWidthChanged(int)
func (this *QWindow) MinimumWidthChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow19minimumWidthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:326
// index:0
// Public Visibility=Default Availability=Available
// [-2] void minimumHeightChanged(int)
func (this *QWindow) MinimumHeightChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow20minimumHeightChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:327
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumWidthChanged(int)
func (this *QWindow) MaximumWidthChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow19maximumWidthChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void maximumHeightChanged(int)
func (this *QWindow) MaximumHeightChanged(arg int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow20maximumHeightChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:330
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibleChanged(_Bool)
func (this *QWindow) VisibleChanged(arg bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14visibleChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:331
// index:0
// Public Visibility=Default Availability=Available
// [-2] void visibilityChanged(QWindow::Visibility)
func (this *QWindow) VisibilityChanged(visibility int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow17visibilityChangedENS_10VisibilityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visibility)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activeChanged()
func (this *QWindow) ActiveChanged() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13activeChangedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:333
// index:0
// Public Visibility=Default Availability=Available
// [-2] void contentOrientationChanged(Qt::ScreenOrientation)
func (this *QWindow) ContentOrientationChanged(orientation int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow25contentOrientationChangedEN2Qt17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), orientation)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:335
// index:0
// Public Visibility=Default Availability=Available
// [-2] void focusObjectChanged(QObject *)
func (this *QWindow) FocusObjectChanged(object *qtcore.QObject /*777 QObject **/) {
	var convArg0 = object.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow18focusObjectChangedEP7QObject", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:337
// index:0
// Public Visibility=Default Availability=Available
// [-2] void opacityChanged(qreal)
func (this *QWindow) OpacityChanged(opacity float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14opacityChangedEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), opacity)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:340
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void exposeEvent(QExposeEvent *)
func (this *QWindow) ExposeEvent(arg0 *QExposeEvent /*777 QExposeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11exposeEventEP12QExposeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:341
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QWindow) ResizeEvent(arg0 *QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:342
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)
func (this *QWindow) MoveEvent(arg0 *QMoveEvent /*777 QMoveEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:343
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QWindow) FocusInEvent(arg0 *QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:344
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QWindow) FocusOutEvent(arg0 *QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:346
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QWindow) ShowEvent(arg0 *QShowEvent /*777 QShowEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:347
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QWindow) HideEvent(arg0 *QHideEvent /*777 QHideEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:350
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QWindow) Event(arg0 *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qwindow.h:351
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QWindow) KeyPressEvent(arg0 *QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:352
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QWindow) KeyReleaseEvent(arg0 *QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:353
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QWindow) MousePressEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:354
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QWindow) MouseReleaseEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:355
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QWindow) MouseDoubleClickEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:356
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QWindow) MouseMoveEvent(arg0 *QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:358
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QWindow) WheelEvent(arg0 *QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:360
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void touchEvent(QTouchEvent *)
func (this *QWindow) TouchEvent(arg0 *QTouchEvent /*777 QTouchEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow10touchEventEP11QTouchEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:362
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabletEvent(QTabletEvent *)
func (this *QWindow) TabletEvent(arg0 *QTabletEvent /*777 QTabletEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11tabletEventEP12QTabletEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:364
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool nativeEvent(const QByteArray &, void *, long *)
func (this *QWindow) NativeEvent(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWindow11nativeEventERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

type QWindow__Visibility = int

const QWindow__Hidden QWindow__Visibility = 0
const QWindow__AutomaticVisibility QWindow__Visibility = 1
const QWindow__Windowed QWindow__Visibility = 2
const QWindow__Minimized QWindow__Visibility = 3
const QWindow__Maximized QWindow__Visibility = 4
const QWindow__FullScreen QWindow__Visibility = 5

type QWindow__AncestorMode = int

const QWindow__ExcludeTransients QWindow__AncestorMode = 0
const QWindow__IncludeTransients QWindow__AncestorMode = 1

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
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
