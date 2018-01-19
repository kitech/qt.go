//  header block begin
// /usr/include/qt/QtGui/qwindow.h
// #include <qwindow.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QWindow struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qwindow.h:97
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QWindow) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:144
// index:0
// void QWindow(class QScreen *)
func NewQWindow(screen unsafe.Pointer) *QWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindowC2EP7QScreen", ffiqt.FFI_TYPE_VOID, cthis, screen)
	gopp.ErrPrint(err, rv)
	return &QWindow{cthis}
}

// /usr/include/qt/QtGui/qwindow.h:145
// index:1
// void QWindow(class QWindow *)
func NewQWindow_1(parent unsafe.Pointer) *QWindow {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindowC2EPS_", ffiqt.FFI_TYPE_VOID, cthis, parent)
	gopp.ErrPrint(err, rv)
	return &QWindow{cthis}
}

// /usr/include/qt/QtGui/qwindow.h:146
// index:0
// virtual
// void ~QWindow()
func DeleteQWindow(*QWindow) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindowD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:148
// index:0
// void setSurfaceType(enum QSurface::SurfaceType)
func (this *QWindow) SetSurfaceType(surfaceType int) {
	// 0: (, QSurface::SurfaceType surfaceType), (&surfaceType)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setSurfaceTypeEN8QSurface11SurfaceTypeE", ffiqt.FFI_TYPE_VOID, this.cthis, &surfaceType)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:149
// index:0
// virtual
// QSurface::SurfaceType surfaceType()
func (this *QWindow) SurfaceType() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11surfaceTypeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:151
// index:0
// bool isVisible()
func (this *QWindow) IsVisible() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow9isVisibleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:153
// index:0
// QWindow::Visibility visibility()
func (this *QWindow) Visibility() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow10visibilityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:154
// index:0
// void setVisibility(enum QWindow::Visibility)
func (this *QWindow) SetVisibility(v int) {
	// 0: (, QWindow::Visibility v), (&v)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13setVisibilityENS_10VisibilityE", ffiqt.FFI_TYPE_VOID, this.cthis, &v)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:156
// index:0
// void create()
func (this *QWindow) Create() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow6createEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:158
// index:0
// WId winId()
func (this *QWindow) WinId() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5winIdEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:160
// index:0
// QWindow * parent(enum QWindow::AncestorMode)
func (this *QWindow) Parent(mode int) {
	// 0: (, QWindow::AncestorMode mode), (&mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6parentENS_12AncestorModeE", ffiqt.FFI_TYPE_VOID, this.cthis, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:161
// index:1
// QWindow * parent()
func (this *QWindow) Parent_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6parentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:162
// index:0
// void setParent(class QWindow *)
func (this *QWindow) SetParent(parent unsafe.Pointer) {
	// 0: (, QWindow * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setParentEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:164
// index:0
// bool isTopLevel()
func (this *QWindow) IsTopLevel() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow10isTopLevelEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:166
// index:0
// bool isModal()
func (this *QWindow) IsModal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow7isModalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:167
// index:0
// Qt::WindowModality modality()
func (this *QWindow) Modality() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8modalityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:168
// index:0
// void setModality(Qt::WindowModality)
func (this *QWindow) SetModality(modality int) {
	// 0: (, Qt::WindowModality modality), (&modality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_VOID, this.cthis, &modality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:170
// index:0
// void setFormat(const class QSurfaceFormat &)
func (this *QWindow) SetFormat(format unsafe.Pointer) {
	// 0: (, const QSurfaceFormat & format), (format)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setFormatERK14QSurfaceFormat", ffiqt.FFI_TYPE_VOID, this.cthis, format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:171
// index:0
// virtual
// QSurfaceFormat format()
func (this *QWindow) Format() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6formatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:172
// index:0
// QSurfaceFormat requestedFormat()
func (this *QWindow) RequestedFormat() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow15requestedFormatEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:175
// index:0
// Qt::WindowFlags flags()
func (this *QWindow) Flags() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5flagsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:176
// index:0
// void setFlag(Qt::WindowType, _Bool)
func (this *QWindow) SetFlag(arg0 int, on bool) {
	// 0: (, Qt::WindowType arg0, bool on), (&arg0, &on)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7setFlagEN2Qt10WindowTypeEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg0, &on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:177
// index:0
// Qt::WindowType type()
func (this *QWindow) Type() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4typeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:179
// index:0
// QString title()
func (this *QWindow) Title() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5titleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:181
// index:0
// void setOpacity(qreal)
func (this *QWindow) SetOpacity(level float64) {
	// 0: (, qreal level), (&level)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10setOpacityEd", ffiqt.FFI_TYPE_VOID, this.cthis, &level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:182
// index:0
// qreal opacity()
func (this *QWindow) Opacity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow7opacityEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:184
// index:0
// void setMask(const class QRegion &)
func (this *QWindow) SetMask(region unsafe.Pointer) {
	// 0: (, const QRegion & region), (region)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7setMaskERK7QRegion", ffiqt.FFI_TYPE_VOID, this.cthis, region)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:185
// index:0
// QRegion mask()
func (this *QWindow) Mask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4maskEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:187
// index:0
// bool isActive()
func (this *QWindow) IsActive() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8isActiveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:189
// index:0
// void reportContentOrientationChange(Qt::ScreenOrientation)
func (this *QWindow) ReportContentOrientationChange(orientation int) {
	// 0: (, Qt::ScreenOrientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow30reportContentOrientationChangeEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:190
// index:0
// Qt::ScreenOrientation contentOrientation()
func (this *QWindow) ContentOrientation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow18contentOrientationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:192
// index:0
// qreal devicePixelRatio()
func (this *QWindow) DevicePixelRatio() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow16devicePixelRatioEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:194
// index:0
// Qt::WindowState windowState()
func (this *QWindow) WindowState() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11windowStateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:195
// index:0
// Qt::WindowStates windowStates()
func (this *QWindow) WindowStates() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12windowStatesEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:196
// index:0
// void setWindowState(Qt::WindowState)
func (this *QWindow) SetWindowState(state int) {
	// 0: (, Qt::WindowState state), (&state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setWindowStateEN2Qt11WindowStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:199
// index:0
// void setTransientParent(class QWindow *)
func (this *QWindow) SetTransientParent(parent unsafe.Pointer) {
	// 0: (, QWindow * parent), (parent)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18setTransientParentEPS_", ffiqt.FFI_TYPE_VOID, this.cthis, parent)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:200
// index:0
// QWindow * transientParent()
func (this *QWindow) TransientParent() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow15transientParentEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:202
// index:0
// bool isAncestorOf(const class QWindow *, enum QWindow::AncestorMode)
func (this *QWindow) IsAncestorOf(child unsafe.Pointer, mode int) {
	// 0: (, const QWindow * child, QWindow::AncestorMode mode), (child, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12isAncestorOfEPKS_NS_12AncestorModeE", ffiqt.FFI_TYPE_VOID, this.cthis, child, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:204
// index:0
// bool isExposed()
func (this *QWindow) IsExposed() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow9isExposedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:206
// index:0
// inline
// int minimumWidth()
func (this *QWindow) MinimumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12minimumWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:207
// index:0
// inline
// int minimumHeight()
func (this *QWindow) MinimumHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13minimumHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:208
// index:0
// inline
// int maximumWidth()
func (this *QWindow) MaximumWidth() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12maximumWidthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:209
// index:0
// inline
// int maximumHeight()
func (this *QWindow) MaximumHeight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13maximumHeightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:211
// index:0
// QSize minimumSize()
func (this *QWindow) MinimumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11minimumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:212
// index:0
// QSize maximumSize()
func (this *QWindow) MaximumSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11maximumSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:213
// index:0
// QSize baseSize()
func (this *QWindow) BaseSize() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8baseSizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:214
// index:0
// QSize sizeIncrement()
func (this *QWindow) SizeIncrement() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13sizeIncrementEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:216
// index:0
// void setMinimumSize(const class QSize &)
func (this *QWindow) SetMinimumSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setMinimumSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:217
// index:0
// void setMaximumSize(const class QSize &)
func (this *QWindow) SetMaximumSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14setMaximumSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:218
// index:0
// void setBaseSize(const class QSize &)
func (this *QWindow) SetBaseSize(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setBaseSizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:219
// index:0
// void setSizeIncrement(const class QSize &)
func (this *QWindow) SetSizeIncrement(size unsafe.Pointer) {
	// 0: (, const QSize & size), (size)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setSizeIncrementERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, size)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:221
// index:0
// QRect geometry()
func (this *QWindow) Geometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8geometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:223
// index:0
// QMargins frameMargins()
func (this *QWindow) FrameMargins() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow12frameMarginsEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:224
// index:0
// QRect frameGeometry()
func (this *QWindow) FrameGeometry() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13frameGeometryEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:226
// index:0
// QPoint framePosition()
func (this *QWindow) FramePosition() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13framePositionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:227
// index:0
// void setFramePosition(const class QPoint &)
func (this *QWindow) SetFramePosition(point unsafe.Pointer) {
	// 0: (, const QPoint & point), (point)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setFramePositionERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, point)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:229
// index:0
// inline
// int width()
func (this *QWindow) Width() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow5widthEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:230
// index:0
// inline
// int height()
func (this *QWindow) Height() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6heightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:231
// index:0
// inline
// int x()
func (this *QWindow) X() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow1xEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:232
// index:0
// inline
// int y()
func (this *QWindow) Y() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow1yEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:234
// index:0
// inline virtual
// QSize size()
func (this *QWindow) Size() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4sizeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:235
// index:0
// inline
// QPoint position()
func (this *QWindow) Position() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8positionEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:237
// index:0
// void setPosition(const class QPoint &)
func (this *QWindow) SetPosition(pt unsafe.Pointer) {
	// 0: (, const QPoint & pt), (pt)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setPositionERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pt)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:238
// index:1
// void setPosition(int, int)
func (this *QWindow) SetPosition_1(posx int, posy int) {
	// 1: (, int posx, int posy), (&posx, &posy)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setPositionEii", ffiqt.FFI_TYPE_VOID, this.cthis, &posx, &posy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:240
// index:0
// void resize(const class QSize &)
func (this *QWindow) Resize(newSize unsafe.Pointer) {
	// 0: (, const QSize & newSize), (newSize)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow6resizeERK5QSize", ffiqt.FFI_TYPE_VOID, this.cthis, newSize)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:241
// index:1
// void resize(int, int)
func (this *QWindow) Resize_1(w int, h int) {
	// 1: (, int w, int h), (&w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow6resizeEii", ffiqt.FFI_TYPE_VOID, this.cthis, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:243
// index:0
// void setFilePath(const class QString &)
func (this *QWindow) SetFilePath(filePath unsafe.Pointer) {
	// 0: (, const QString & filePath), (filePath)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setFilePathERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, filePath)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:244
// index:0
// QString filePath()
func (this *QWindow) FilePath() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow8filePathEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:246
// index:0
// void setIcon(const class QIcon &)
func (this *QWindow) SetIcon(icon unsafe.Pointer) {
	// 0: (, const QIcon & icon), (icon)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7setIconERK5QIcon", ffiqt.FFI_TYPE_VOID, this.cthis, icon)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:247
// index:0
// QIcon icon()
func (this *QWindow) Icon() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow4iconEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:249
// index:0
// void destroy()
func (this *QWindow) Destroy() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow7destroyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:251
// index:0
// QPlatformWindow * handle()
func (this *QWindow) Handle() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6handleEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:253
// index:0
// bool setKeyboardGrabEnabled(_Bool)
func (this *QWindow) SetKeyboardGrabEnabled(grab bool) {
	// 0: (, bool grab), (&grab)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow22setKeyboardGrabEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &grab)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:254
// index:0
// bool setMouseGrabEnabled(_Bool)
func (this *QWindow) SetMouseGrabEnabled(grab bool) {
	// 0: (, bool grab), (&grab)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow19setMouseGrabEnabledEb", ffiqt.FFI_TYPE_VOID, this.cthis, &grab)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:256
// index:0
// QScreen * screen()
func (this *QWindow) Screen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6screenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:257
// index:0
// void setScreen(class QScreen *)
func (this *QWindow) SetScreen(screen unsafe.Pointer) {
	// 0: (, QScreen * screen), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setScreenEP7QScreen", ffiqt.FFI_TYPE_VOID, this.cthis, screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:259
// index:0
// virtual
// QAccessibleInterface * accessibleRoot()
func (this *QWindow) AccessibleRoot() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow14accessibleRootEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:260
// index:0
// virtual
// QObject * focusObject()
func (this *QWindow) FocusObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11focusObjectEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:262
// index:0
// QPoint mapToGlobal(const class QPoint &)
func (this *QWindow) MapToGlobal(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow11mapToGlobalERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:263
// index:0
// QPoint mapFromGlobal(const class QPoint &)
func (this *QWindow) MapFromGlobal(pos unsafe.Pointer) {
	// 0: (, const QPoint & pos), (pos)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow13mapFromGlobalERK6QPoint", ffiqt.FFI_TYPE_VOID, this.cthis, pos)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:266
// index:0
// QCursor cursor()
func (this *QWindow) Cursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow6cursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:267
// index:0
// void setCursor(const class QCursor &)
func (this *QWindow) SetCursor(arg0 unsafe.Pointer) {
	// 0: (, const QCursor & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setCursorERK7QCursor", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:268
// index:0
// void unsetCursor()
func (this *QWindow) UnsetCursor() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11unsetCursorEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:271
// index:0
// static
// QWindow * fromWinId(WId)
func (this *QWindow) FromWinId(id uint64) {
	// 0: (WId id), (id)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9fromWinIdEy", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QWindow_FromWinId(id uint64) {
	// 0: (WId id), (id)
	var nilthis *QWindow
	nilthis.FromWinId(id)
}

// /usr/include/qt/QtGui/qwindow.h:274
// index:0
// void setVulkanInstance(class QVulkanInstance *)
func (this *QWindow) SetVulkanInstance(instance unsafe.Pointer) {
	// 0: (, QVulkanInstance * instance), (instance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow17setVulkanInstanceEP15QVulkanInstance", ffiqt.FFI_TYPE_VOID, this.cthis, instance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:275
// index:0
// QVulkanInstance * vulkanInstance()
func (this *QWindow) VulkanInstance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWindow14vulkanInstanceEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:279
// index:0
// void requestActivate()
func (this *QWindow) RequestActivate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15requestActivateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:281
// index:0
// void setVisible(_Bool)
func (this *QWindow) SetVisible(visible bool) {
	// 0: (, bool visible), (&visible)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10setVisibleEb", ffiqt.FFI_TYPE_VOID, this.cthis, &visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:283
// index:0
// void show()
func (this *QWindow) Show() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4showEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:284
// index:0
// void hide()
func (this *QWindow) Hide() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4hideEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:286
// index:0
// void showMinimized()
func (this *QWindow) ShowMinimized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13showMinimizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:287
// index:0
// void showMaximized()
func (this *QWindow) ShowMaximized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13showMaximizedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:288
// index:0
// void showFullScreen()
func (this *QWindow) ShowFullScreen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14showFullScreenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:289
// index:0
// void showNormal()
func (this *QWindow) ShowNormal() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow10showNormalEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:291
// index:0
// bool close()
func (this *QWindow) Close() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5closeEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:292
// index:0
// void raise()
func (this *QWindow) Raise() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5raiseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:293
// index:0
// void lower()
func (this *QWindow) Lower() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5lowerEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:295
// index:0
// void setTitle(const class QString &)
func (this *QWindow) SetTitle(arg0 unsafe.Pointer) {
	// 0: (, const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8setTitleERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:297
// index:0
// void setX(int)
func (this *QWindow) SetX(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4setXEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:298
// index:0
// void setY(int)
func (this *QWindow) SetY(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow4setYEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:299
// index:0
// void setWidth(int)
func (this *QWindow) SetWidth(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8setWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:300
// index:0
// void setHeight(int)
func (this *QWindow) SetHeight(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow9setHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:301
// index:0
// void setGeometry(int, int, int, int)
func (this *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	// 0: (, int posx, int posy, int w, int h), (&posx, &posy, &w, &h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setGeometryEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &posx, &posy, &w, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:302
// index:1
// void setGeometry(const class QRect &)
func (this *QWindow) SetGeometry_1(rect unsafe.Pointer) {
	// 1: (, const QRect & rect), (rect)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow11setGeometryERK5QRect", ffiqt.FFI_TYPE_VOID, this.cthis, rect)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:304
// index:0
// void setMinimumWidth(int)
func (this *QWindow) SetMinimumWidth(w int) {
	// 0: (, int w), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15setMinimumWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:305
// index:0
// void setMinimumHeight(int)
func (this *QWindow) SetMinimumHeight(h int) {
	// 0: (, int h), (&h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setMinimumHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:306
// index:0
// void setMaximumWidth(int)
func (this *QWindow) SetMaximumWidth(w int) {
	// 0: (, int w), (&w)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15setMaximumWidthEi", ffiqt.FFI_TYPE_VOID, this.cthis, &w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:307
// index:0
// void setMaximumHeight(int)
func (this *QWindow) SetMaximumHeight(h int) {
	// 0: (, int h), (&h)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow16setMaximumHeightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:309
// index:0
// void alert(int)
func (this *QWindow) Alert(msec int) {
	// 0: (, int msec), (&msec)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow5alertEi", ffiqt.FFI_TYPE_VOID, this.cthis, &msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:311
// index:0
// void requestUpdate()
func (this *QWindow) RequestUpdate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13requestUpdateEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:314
// index:0
// void screenChanged(class QScreen *)
func (this *QWindow) ScreenChanged(screen unsafe.Pointer) {
	// 0: (, QScreen * screen), (screen)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13screenChangedEP7QScreen", ffiqt.FFI_TYPE_VOID, this.cthis, screen)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:315
// index:0
// void modalityChanged(Qt::WindowModality)
func (this *QWindow) ModalityChanged(modality int) {
	// 0: (, Qt::WindowModality modality), (&modality)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow15modalityChangedEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_VOID, this.cthis, &modality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:316
// index:0
// void windowStateChanged(Qt::WindowState)
func (this *QWindow) WindowStateChanged(windowState int) {
	// 0: (, Qt::WindowState windowState), (&windowState)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18windowStateChangedEN2Qt11WindowStateE", ffiqt.FFI_TYPE_VOID, this.cthis, &windowState)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:317
// index:0
// void windowTitleChanged(const class QString &)
func (this *QWindow) WindowTitleChanged(title unsafe.Pointer) {
	// 0: (, const QString & title), (title)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18windowTitleChangedERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, title)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:319
// index:0
// void xChanged(int)
func (this *QWindow) XChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8xChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:320
// index:0
// void yChanged(int)
func (this *QWindow) YChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow8yChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:322
// index:0
// void widthChanged(int)
func (this *QWindow) WidthChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow12widthChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:323
// index:0
// void heightChanged(int)
func (this *QWindow) HeightChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13heightChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:325
// index:0
// void minimumWidthChanged(int)
func (this *QWindow) MinimumWidthChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow19minimumWidthChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:326
// index:0
// void minimumHeightChanged(int)
func (this *QWindow) MinimumHeightChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow20minimumHeightChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:327
// index:0
// void maximumWidthChanged(int)
func (this *QWindow) MaximumWidthChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow19maximumWidthChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:328
// index:0
// void maximumHeightChanged(int)
func (this *QWindow) MaximumHeightChanged(arg int) {
	// 0: (, int arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow20maximumHeightChangedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:330
// index:0
// void visibleChanged(_Bool)
func (this *QWindow) VisibleChanged(arg bool) {
	// 0: (, bool arg), (&arg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14visibleChangedEb", ffiqt.FFI_TYPE_VOID, this.cthis, &arg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:331
// index:0
// void visibilityChanged(class QWindow::Visibility)
func (this *QWindow) VisibilityChanged(visibility int) {
	// 0: (, QWindow::Visibility visibility), (&visibility)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow17visibilityChangedENS_10VisibilityE", ffiqt.FFI_TYPE_VOID, this.cthis, &visibility)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:332
// index:0
// void activeChanged()
func (this *QWindow) ActiveChanged() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow13activeChangedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:333
// index:0
// void contentOrientationChanged(Qt::ScreenOrientation)
func (this *QWindow) ContentOrientationChanged(orientation int) {
	// 0: (, Qt::ScreenOrientation orientation), (&orientation)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow25contentOrientationChangedEN2Qt17ScreenOrientationE", ffiqt.FFI_TYPE_VOID, this.cthis, &orientation)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:335
// index:0
// void focusObjectChanged(class QObject *)
func (this *QWindow) FocusObjectChanged(object unsafe.Pointer) {
	// 0: (, QObject * object), (object)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow18focusObjectChangedEP7QObject", ffiqt.FFI_TYPE_VOID, this.cthis, object)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qwindow.h:337
// index:0
// void opacityChanged(qreal)
func (this *QWindow) OpacityChanged(opacity float64) {
	// 0: (, qreal opacity), (&opacity)
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWindow14opacityChangedEd", ffiqt.FFI_TYPE_VOID, this.cthis, &opacity)
	gopp.ErrPrint(err, rv)
}

//  body block end
