package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
// bool event(class QEvent *)
func (this *QWidget) InheritEvent(f func(event *qtcore.QEvent /*777 QEvent **/) bool) {
	qtrt.SetAllInheritCallback(this, "event", f)
}

// void mousePressEvent(class QMouseEvent *)
func (this *QWidget) InheritMousePressEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mousePressEvent", f)
}

// void mouseReleaseEvent(class QMouseEvent *)
func (this *QWidget) InheritMouseReleaseEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseReleaseEvent", f)
}

// void mouseDoubleClickEvent(class QMouseEvent *)
func (this *QWidget) InheritMouseDoubleClickEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseDoubleClickEvent", f)
}

// void mouseMoveEvent(class QMouseEvent *)
func (this *QWidget) InheritMouseMoveEvent(f func(event *qtgui.QMouseEvent /*777 QMouseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "mouseMoveEvent", f)
}

// void wheelEvent(class QWheelEvent *)
func (this *QWidget) InheritWheelEvent(f func(event *qtgui.QWheelEvent /*777 QWheelEvent **/)) {
	qtrt.SetAllInheritCallback(this, "wheelEvent", f)
}

// void keyPressEvent(class QKeyEvent *)
func (this *QWidget) InheritKeyPressEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyPressEvent", f)
}

// void keyReleaseEvent(class QKeyEvent *)
func (this *QWidget) InheritKeyReleaseEvent(f func(event *qtgui.QKeyEvent /*777 QKeyEvent **/)) {
	qtrt.SetAllInheritCallback(this, "keyReleaseEvent", f)
}

// void focusInEvent(class QFocusEvent *)
func (this *QWidget) InheritFocusInEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusInEvent", f)
}

// void focusOutEvent(class QFocusEvent *)
func (this *QWidget) InheritFocusOutEvent(f func(event *qtgui.QFocusEvent /*777 QFocusEvent **/)) {
	qtrt.SetAllInheritCallback(this, "focusOutEvent", f)
}

// void enterEvent(class QEvent *)
func (this *QWidget) InheritEnterEvent(f func(event *qtcore.QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "enterEvent", f)
}

// void leaveEvent(class QEvent *)
func (this *QWidget) InheritLeaveEvent(f func(event *qtcore.QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "leaveEvent", f)
}

// void paintEvent(class QPaintEvent *)
func (this *QWidget) InheritPaintEvent(f func(event *qtgui.QPaintEvent /*777 QPaintEvent **/)) {
	qtrt.SetAllInheritCallback(this, "paintEvent", f)
}

// void moveEvent(class QMoveEvent *)
func (this *QWidget) InheritMoveEvent(f func(event *qtgui.QMoveEvent /*777 QMoveEvent **/)) {
	qtrt.SetAllInheritCallback(this, "moveEvent", f)
}

// void resizeEvent(class QResizeEvent *)
func (this *QWidget) InheritResizeEvent(f func(event *qtgui.QResizeEvent /*777 QResizeEvent **/)) {
	qtrt.SetAllInheritCallback(this, "resizeEvent", f)
}

// void closeEvent(class QCloseEvent *)
func (this *QWidget) InheritCloseEvent(f func(event *qtgui.QCloseEvent /*777 QCloseEvent **/)) {
	qtrt.SetAllInheritCallback(this, "closeEvent", f)
}

// void contextMenuEvent(class QContextMenuEvent *)
func (this *QWidget) InheritContextMenuEvent(f func(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/)) {
	qtrt.SetAllInheritCallback(this, "contextMenuEvent", f)
}

// void tabletEvent(class QTabletEvent *)
func (this *QWidget) InheritTabletEvent(f func(event *qtgui.QTabletEvent /*777 QTabletEvent **/)) {
	qtrt.SetAllInheritCallback(this, "tabletEvent", f)
}

// void actionEvent(class QActionEvent *)
func (this *QWidget) InheritActionEvent(f func(event *qtgui.QActionEvent /*777 QActionEvent **/)) {
	qtrt.SetAllInheritCallback(this, "actionEvent", f)
}

// void dragEnterEvent(class QDragEnterEvent *)
func (this *QWidget) InheritDragEnterEvent(f func(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dragEnterEvent", f)
}

// void dragMoveEvent(class QDragMoveEvent *)
func (this *QWidget) InheritDragMoveEvent(f func(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dragMoveEvent", f)
}

// void dragLeaveEvent(class QDragLeaveEvent *)
func (this *QWidget) InheritDragLeaveEvent(f func(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dragLeaveEvent", f)
}

// void dropEvent(class QDropEvent *)
func (this *QWidget) InheritDropEvent(f func(event *qtgui.QDropEvent /*777 QDropEvent **/)) {
	qtrt.SetAllInheritCallback(this, "dropEvent", f)
}

// void showEvent(class QShowEvent *)
func (this *QWidget) InheritShowEvent(f func(event *qtgui.QShowEvent /*777 QShowEvent **/)) {
	qtrt.SetAllInheritCallback(this, "showEvent", f)
}

// void hideEvent(class QHideEvent *)
func (this *QWidget) InheritHideEvent(f func(event *qtgui.QHideEvent /*777 QHideEvent **/)) {
	qtrt.SetAllInheritCallback(this, "hideEvent", f)
}

// bool nativeEvent(const class QByteArray &, void *, long *)
func (this *QWidget) InheritNativeEvent(f func(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool) {
	qtrt.SetAllInheritCallback(this, "nativeEvent", f)
}

// void changeEvent(class QEvent *)
func (this *QWidget) InheritChangeEvent(f func(arg0 *qtcore.QEvent /*777 QEvent **/)) {
	qtrt.SetAllInheritCallback(this, "changeEvent", f)
}

// int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QWidget) InheritMetric(f func(arg0 int) int) {
	qtrt.SetAllInheritCallback(this, "metric", f)
}

// void initPainter(class QPainter *)
func (this *QWidget) InheritInitPainter(f func(painter *qtgui.QPainter /*777 QPainter **/)) {
	qtrt.SetAllInheritCallback(this, "initPainter", f)
}

// QPaintDevice * redirected(class QPoint *)
func (this *QWidget) InheritRedirected(f func(offset *qtcore.QPoint /*777 QPoint **/) unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "redirected", f)
}

// QPainter * sharedPainter()
func (this *QWidget) InheritSharedPainter(f func() unsafe.Pointer /*666*/) {
	qtrt.SetAllInheritCallback(this, "sharedPainter", f)
}

// void inputMethodEvent(class QInputMethodEvent *)
func (this *QWidget) InheritInputMethodEvent(f func(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/)) {
	qtrt.SetAllInheritCallback(this, "inputMethodEvent", f)
}

// void updateMicroFocus()
func (this *QWidget) InheritUpdateMicroFocus(f func()) {
	qtrt.SetAllInheritCallback(this, "updateMicroFocus", f)
}

// void create(WId, _Bool, _Bool)
func (this *QWidget) InheritCreate(f func(arg0 uint64, initializeWindow bool, destroyOldWindow bool)) {
	qtrt.SetAllInheritCallback(this, "create", f)
}

// void destroy(_Bool, _Bool)
func (this *QWidget) InheritDestroy(f func(destroyWindow bool, destroySubWindows bool)) {
	qtrt.SetAllInheritCallback(this, "destroy", f)
}

// bool focusNextPrevChild(_Bool)
func (this *QWidget) InheritFocusNextPrevChild(f func(next bool) bool) {
	qtrt.SetAllInheritCallback(this, "focusNextPrevChild", f)
}

// bool focusNextChild()
func (this *QWidget) InheritFocusNextChild(f func() bool) {
	qtrt.SetAllInheritCallback(this, "focusNextChild", f)
}

// bool focusPreviousChild()
func (this *QWidget) InheritFocusPreviousChild(f func() bool) {
	qtrt.SetAllInheritCallback(this, "focusPreviousChild", f)
}

type QWidget struct {
	*qtcore.QObject
	*qtgui.QPaintDevice
}

func (this *QWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWidget) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
	this.QPaintDevice = qtgui.NewQPaintDeviceFromPointer(cthis)
}
func NewQWidgetFromPointer(cthis unsafe.Pointer) *QWidget {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	bcthis1 := qtgui.NewQPaintDeviceFromPointer(cthis)
	return &QWidget{bcthis0, bcthis1}
}
func (*QWidget) NewFromPointer(cthis unsafe.Pointer) *QWidget {
	return NewQWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qwidget.h:130
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:214
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWidget(QWidget *, Qt::WindowFlags)
func NewQWidget(parent *QWidget /*777 QWidget **/, f int) *QWidget {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWidget()
func DeleteQWidget(this *QWidget) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidgetD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 48)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWidgets/qwidget.h:217
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int devType()
func (this *QWidget) DevType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7devTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:219
// index:0
// Public Visibility=Default Availability=Available
// [8] WId winId()
func (this *QWidget) WinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5winIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:220
// index:0
// Public Visibility=Default Availability=Available
// [-2] void createWinId()
func (this *QWidget) CreateWinId() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11createWinIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:221
// index:0
// Public inline Visibility=Default Availability=Available
// [8] WId internalWinId()
func (this *QWidget) InternalWinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13internalWinIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] WId effectiveWinId()
func (this *QWidget) EffectiveWinId() uint64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14effectiveWinIdEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:225
// index:0
// Public Visibility=Default Availability=Available
// [8] QStyle * style()
func (this *QWidget) Style() *QStyle /*777 QStyle **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5styleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:226
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyle(QStyle *)
func (this *QWidget) SetStyle(arg0 *QStyle /*777 QStyle **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget8setStyleEP6QStyle", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:229
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isTopLevel()
func (this *QWidget) IsTopLevel() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10isTopLevelEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:230
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWindow()
func (this *QWidget) IsWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8isWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:232
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isModal()
func (this *QWidget) IsModal() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7isModalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:233
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowModality windowModality()
func (this *QWidget) WindowModality() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14windowModalityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:234
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowModality(Qt::WindowModality)
func (this *QWidget) SetWindowModality(windowModality int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowModalityEN2Qt14WindowModalityE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:236
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabled()
func (this *QWidget) IsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9isEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:237
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabledTo(const QWidget *)
func (this *QWidget) IsEnabledTo(arg0 *QWidget /*777 const QWidget **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isEnabledToEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:238
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEnabledToTLW()
func (this *QWidget) IsEnabledToTLW() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14isEnabledToTLWEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:241
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setEnabled(_Bool)
func (this *QWidget) SetEnabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDisabled(_Bool)
func (this *QWidget) SetDisabled(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setDisabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:243
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowModified(_Bool)
func (this *QWidget) SetWindowModified(arg0 bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowModifiedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:248
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect frameGeometry()
func (this *QWidget) FrameGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13frameGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:249
// index:0
// Public Visibility=Default Availability=Available
// [16] const QRect & geometry()
func (this *QWidget) Geometry() *qtcore.QRect {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8geometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:250
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect normalGeometry()
func (this *QWidget) NormalGeometry() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14normalGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:252
// index:0
// Public Visibility=Default Availability=Available
// [4] int x()
func (this *QWidget) X() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget1xEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:253
// index:0
// Public Visibility=Default Availability=Available
// [4] int y()
func (this *QWidget) Y() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget1yEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:254
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint pos()
func (this *QWidget) Pos() *qtcore.QPoint /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget3posEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:255
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize frameSize()
func (this *QWidget) FrameSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9frameSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:256
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize size()
func (this *QWidget) Size() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4sizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:257
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int width()
func (this *QWidget) Width() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5widthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:258
// index:0
// Public inline Visibility=Default Availability=Available
// [4] int height()
func (this *QWidget) Height() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6heightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:259
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QRect rect()
func (this *QWidget) Rect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4rectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:260
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect childrenRect()
func (this *QWidget) ChildrenRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12childrenRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:261
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion childrenRegion()
func (this *QWidget) ChildrenRegion() *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14childrenRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:263
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize minimumSize()
func (this *QWidget) MinimumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11minimumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:264
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize maximumSize()
func (this *QWidget) MaximumSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11maximumSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:265
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumWidth()
func (this *QWidget) MinimumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12minimumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:266
// index:0
// Public Visibility=Default Availability=Available
// [4] int minimumHeight()
func (this *QWidget) MinimumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13minimumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:267
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumWidth()
func (this *QWidget) MaximumWidth() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12maximumWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:268
// index:0
// Public Visibility=Default Availability=Available
// [4] int maximumHeight()
func (this *QWidget) MaximumHeight() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13maximumHeightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:269
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(const QSize &)
func (this *QWidget) SetMinimumSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:270
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMinimumSize(int, int)
func (this *QWidget) SetMinimumSize_1(minw int, minh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minw, minh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:271
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(const QSize &)
func (this *QWidget) SetMaximumSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:272
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMaximumSize(int, int)
func (this *QWidget) SetMaximumSize_1(maxw int, maxh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxw, maxh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:273
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumWidth(int)
func (this *QWidget) SetMinimumWidth(minw int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15setMinimumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minw)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:274
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMinimumHeight(int)
func (this *QWidget) SetMinimumHeight(minh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setMinimumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), minh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:275
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumWidth(int)
func (this *QWidget) SetMaximumWidth(maxw int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15setMaximumWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxw)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:276
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMaximumHeight(int)
func (this *QWidget) SetMaximumHeight(maxh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setMaximumHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), maxh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:282
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize sizeIncrement()
func (this *QWidget) SizeIncrement() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13sizeIncrementEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:283
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizeIncrement(const QSize &)
func (this *QWidget) SetSizeIncrement(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:284
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setSizeIncrement(int, int)
func (this *QWidget) SetSizeIncrement_1(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:285
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize baseSize()
func (this *QWidget) BaseSize() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8baseSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:286
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBaseSize(const QSize &)
func (this *QWidget) SetBaseSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:287
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setBaseSize(int, int)
func (this *QWidget) SetBaseSize_1(basew int, baseh int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), basew, baseh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:289
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedSize(const QSize &)
func (this *QWidget) SetFixedSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:290
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFixedSize(int, int)
func (this *QWidget) SetFixedSize_1(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:291
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedWidth(int)
func (this *QWidget) SetFixedWidth(w int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setFixedWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:292
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFixedHeight(int)
func (this *QWidget) SetFixedHeight(h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setFixedHeightEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:296
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapToGlobal(const QPoint &)
func (this *QWidget) MapToGlobal(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11mapToGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:297
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromGlobal(const QPoint &)
func (this *QWidget) MapFromGlobal(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13mapFromGlobalERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:298
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapToParent(const QPoint &)
func (this *QWidget) MapToParent(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11mapToParentERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:299
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFromParent(const QPoint &)
func (this *QWidget) MapFromParent(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13mapFromParentERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:300
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapTo(const QWidget *, const QPoint &)
func (this *QWidget) MapTo(arg0 *QWidget /*777 const QWidget **/, arg1 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget5mapToEPKS_RK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:301
// index:0
// Public Visibility=Default Availability=Available
// [8] QPoint mapFrom(const QWidget *, const QPoint &)
func (this *QWidget) MapFrom(arg0 *QWidget /*777 const QWidget **/, arg1 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7mapFromEPKS_RK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQPoint)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:303
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * window()
func (this *QWidget) Window() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:304
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * nativeParentWidget()
func (this *QWidget) NativeParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget18nativeParentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:305
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWidget * topLevelWidget()
func (this *QWidget) TopLevelWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14topLevelWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:308
// index:0
// Public Visibility=Default Availability=Available
// [16] const QPalette & palette()
func (this *QWidget) Palette() *qtgui.QPalette {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7paletteEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPalette)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:309
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPalette(const QPalette &)
func (this *QWidget) SetPalette(arg0 *qtgui.QPalette) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setPaletteERK8QPalette", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:311
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBackgroundRole(QPalette::ColorRole)
func (this *QWidget) SetBackgroundRole(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setBackgroundRoleEN8QPalette9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:312
// index:0
// Public Visibility=Default Availability=Available
// [4] QPalette::ColorRole backgroundRole()
func (this *QWidget) BackgroundRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14backgroundRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:314
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setForegroundRole(QPalette::ColorRole)
func (this *QWidget) SetForegroundRole(arg0 int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setForegroundRoleEN8QPalette9ColorRoleE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:315
// index:0
// Public Visibility=Default Availability=Available
// [4] QPalette::ColorRole foregroundRole()
func (this *QWidget) ForegroundRole() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14foregroundRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:317
// index:0
// Public Visibility=Default Availability=Available
// [16] const QFont & font()
func (this *QWidget) Font() *qtgui.QFont {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4fontEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFont)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:318
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFont(const QFont &)
func (this *QWidget) SetFont(arg0 *qtgui.QFont) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7setFontERK5QFont", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:319
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontMetrics fontMetrics()
func (this *QWidget) FontMetrics() *qtgui.QFontMetrics /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11fontMetricsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFontMetrics)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:320
// index:0
// Public Visibility=Default Availability=Available
// [8] QFontInfo fontInfo()
func (this *QWidget) FontInfo() *qtgui.QFontInfo /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8fontInfoEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQFontInfo)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:323
// index:0
// Public Visibility=Default Availability=Available
// [8] QCursor cursor()
func (this *QWidget) Cursor() *qtgui.QCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6cursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQCursor)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:324
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursor(const QCursor &)
func (this *QWidget) SetCursor(arg0 *qtgui.QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setCursorERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:325
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetCursor()
func (this *QWidget) UnsetCursor() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11unsetCursorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:328
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMouseTracking(_Bool)
func (this *QWidget) SetMouseTracking(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setMouseTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:329
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasMouseTracking()
func (this *QWidget) HasMouseTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16hasMouseTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:330
// index:0
// Public Visibility=Default Availability=Available
// [1] bool underMouse()
func (this *QWidget) UnderMouse() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10underMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:332
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabletTracking(_Bool)
func (this *QWidget) SetTabletTracking(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setTabletTrackingEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:333
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasTabletTracking()
func (this *QWidget) HasTabletTracking() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget17hasTabletTrackingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:335
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QBitmap &)
func (this *QWidget) SetMask(arg0 *qtgui.QBitmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QBitmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:336
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setMask(const QRegion &)
func (this *QWidget) SetMask_1(arg0 *qtgui.QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:337
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion mask()
func (this *QWidget) Mask() *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget4maskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:338
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearMask()
func (this *QWidget) ClearMask() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9clearMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:340
// index:0
// Public Visibility=Default Availability=Available
// [-2] void render(QPaintDevice *, const QPoint &, const QRegion &, QWidget::RenderFlags)
func (this *QWidget) Render(target *qtgui.QPaintDevice /*777 QPaintDevice **/, targetOffset *qtcore.QPoint, sourceRegion *qtgui.QRegion, renderFlags int) {
	var convArg0 = target.GetCthis()
	var convArg1 = targetOffset.GetCthis()
	var convArg2 = sourceRegion.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP12QPaintDeviceRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:344
// index:1
// Public Visibility=Default Availability=Available
// [-2] void render(QPainter *, const QPoint &, const QRegion &, QWidget::RenderFlags)
func (this *QWidget) Render_1(painter *qtgui.QPainter /*777 QPainter **/, targetOffset *qtcore.QPoint, sourceRegion *qtgui.QRegion, renderFlags int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = targetOffset.GetCthis()
	var convArg2 = sourceRegion.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6renderEP8QPainterRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:348
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap grab(const QRect &)
func (this *QWidget) Grab(rectangle *qtcore.QRect) *qtgui.QPixmap /*123*/ {
	var convArg0 = rectangle.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4grabERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:351
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsEffect * graphicsEffect()
func (this *QWidget) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14graphicsEffectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:352
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGraphicsEffect(QGraphicsEffect *)
func (this *QWidget) SetGraphicsEffect(effect *QGraphicsEffect /*777 QGraphicsEffect **/) {
	var convArg0 = effect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:356
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabGesture(Qt::GestureType, Qt::GestureFlags)
func (this *QWidget) GrabGesture(type_ int, flags int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11grabGestureEN2Qt11GestureTypeE6QFlagsINS0_11GestureFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:357
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ungrabGesture(Qt::GestureType)
func (this *QWidget) UngrabGesture(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13ungrabGestureEN2Qt11GestureTypeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:361
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowTitle(const QString &)
func (this *QWidget) SetWindowTitle(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setWindowTitleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:363
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStyleSheet(const QString &)
func (this *QWidget) SetStyleSheet(styleSheet *qtcore.QString) {
	var convArg0 = styleSheet.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setStyleSheetERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:367
// index:0
// Public Visibility=Default Availability=Available
// [8] QString styleSheet()
func (this *QWidget) StyleSheet() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10styleSheetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:369
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowTitle()
func (this *QWidget) WindowTitle() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11windowTitleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:370
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowIcon(const QIcon &)
func (this *QWidget) SetWindowIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowIconERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:371
// index:0
// Public Visibility=Default Availability=Available
// [8] QIcon windowIcon()
func (this *QWidget) WindowIcon() *qtgui.QIcon /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10windowIconEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQIcon)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:372
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowIconText(const QString &)
func (this *QWidget) SetWindowIconText(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowIconTextERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:373
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowIconText()
func (this *QWidget) WindowIconText() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14windowIconTextEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:374
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowRole(const QString &)
func (this *QWidget) SetWindowRole(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowRoleERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:375
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowRole()
func (this *QWidget) WindowRole() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10windowRoleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:376
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFilePath(const QString &)
func (this *QWidget) SetWindowFilePath(filePath *qtcore.QString) {
	var convArg0 = filePath.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setWindowFilePathERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:377
// index:0
// Public Visibility=Default Availability=Available
// [8] QString windowFilePath()
func (this *QWidget) WindowFilePath() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14windowFilePathEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:379
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowOpacity(qreal)
func (this *QWidget) SetWindowOpacity(level float64) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16setWindowOpacityEd", qtrt.FFI_TYPE_POINTER, this.GetCthis(), level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:380
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal windowOpacity()
func (this *QWidget) WindowOpacity() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13windowOpacityEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:382
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isWindowModified()
func (this *QWidget) IsWindowModified() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16isWindowModifiedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:384
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTip(const QString &)
func (this *QWidget) SetToolTip(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setToolTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:385
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toolTip()
func (this *QWidget) ToolTip() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7toolTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:386
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setToolTipDuration(int)
func (this *QWidget) SetToolTipDuration(msec int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setToolTipDurationEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:387
// index:0
// Public Visibility=Default Availability=Available
// [4] int toolTipDuration()
func (this *QWidget) ToolTipDuration() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15toolTipDurationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:390
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStatusTip(const QString &)
func (this *QWidget) SetStatusTip(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setStatusTipERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:391
// index:0
// Public Visibility=Default Availability=Available
// [8] QString statusTip()
func (this *QWidget) StatusTip() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9statusTipEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:394
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWhatsThis(const QString &)
func (this *QWidget) SetWhatsThis(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setWhatsThisERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:395
// index:0
// Public Visibility=Default Availability=Available
// [8] QString whatsThis()
func (this *QWidget) WhatsThis() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9whatsThisEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:398
// index:0
// Public Visibility=Default Availability=Available
// [8] QString accessibleName()
func (this *QWidget) AccessibleName() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14accessibleNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:399
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccessibleName(const QString &)
func (this *QWidget) SetAccessibleName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setAccessibleNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:400
// index:0
// Public Visibility=Default Availability=Available
// [8] QString accessibleDescription()
func (this *QWidget) AccessibleDescription() *qtcore.QString /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget21accessibleDescriptionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:401
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAccessibleDescription(const QString &)
func (this *QWidget) SetAccessibleDescription(description *qtcore.QString) {
	var convArg0 = description.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget24setAccessibleDescriptionERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:404
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutDirection(Qt::LayoutDirection)
func (this *QWidget) SetLayoutDirection(direction int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:405
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::LayoutDirection layoutDirection()
func (this *QWidget) LayoutDirection() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15layoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:406
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetLayoutDirection()
func (this *QWidget) UnsetLayoutDirection() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget20unsetLayoutDirectionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:408
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLocale(const QLocale &)
func (this *QWidget) SetLocale(locale *qtcore.QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setLocaleERK7QLocale", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:409
// index:0
// Public Visibility=Default Availability=Available
// [8] QLocale locale()
func (this *QWidget) Locale() *qtcore.QLocale /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6localeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQLocale)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:410
// index:0
// Public Visibility=Default Availability=Available
// [-2] void unsetLocale()
func (this *QWidget) UnsetLocale() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11unsetLocaleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:412
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isRightToLeft()
func (this *QWidget) IsRightToLeft() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13isRightToLeftEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:413
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isLeftToRight()
func (this *QWidget) IsLeftToRight() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13isLeftToRightEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:416
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFocus()
func (this *QWidget) SetFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget8setFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:423
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFocus(Qt::FocusReason)
func (this *QWidget) SetFocus_1(reason int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget8setFocusEN2Qt11FocusReasonE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:419
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isActiveWindow()
func (this *QWidget) IsActiveWindow() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14isActiveWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:420
// index:0
// Public Visibility=Default Availability=Available
// [-2] void activateWindow()
func (this *QWidget) ActivateWindow() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14activateWindowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:421
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clearFocus()
func (this *QWidget) ClearFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10clearFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:424
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::FocusPolicy focusPolicy()
func (this *QWidget) FocusPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11focusPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:425
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusPolicy(Qt::FocusPolicy)
func (this *QWidget) SetFocusPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setFocusPolicyEN2Qt11FocusPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:426
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFocus()
func (this *QWidget) HasFocus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8hasFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:427
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setTabOrder(QWidget *, QWidget *)
func (this *QWidget) SetTabOrder(arg0 *QWidget /*777 QWidget **/, arg1 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setTabOrderEPS_S0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}
func QWidget_SetTabOrder(arg0 *QWidget /*777 QWidget **/, arg1 *QWidget /*777 QWidget **/) {
	var nilthis *QWidget
	nilthis.SetTabOrder(arg0, arg1)
}

// /usr/include/qt/QtWidgets/qwidget.h:428
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFocusProxy(QWidget *)
func (this *QWidget) SetFocusProxy(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setFocusProxyEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:429
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * focusProxy()
func (this *QWidget) FocusProxy() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10focusProxyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:430
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::ContextMenuPolicy contextMenuPolicy()
func (this *QWidget) ContextMenuPolicy() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget17contextMenuPolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:431
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContextMenuPolicy(Qt::ContextMenuPolicy)
func (this *QWidget) SetContextMenuPolicy(policy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget20setContextMenuPolicyEN2Qt17ContextMenuPolicyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:434
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabMouse()
func (this *QWidget) GrabMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9grabMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:436
// index:1
// Public Visibility=Default Availability=Available
// [-2] void grabMouse(const QCursor &)
func (this *QWidget) GrabMouse_1(arg0 *qtgui.QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9grabMouseERK7QCursor", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:438
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseMouse()
func (this *QWidget) ReleaseMouse() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12releaseMouseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:439
// index:0
// Public Visibility=Default Availability=Available
// [-2] void grabKeyboard()
func (this *QWidget) GrabKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12grabKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:440
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseKeyboard()
func (this *QWidget) ReleaseKeyboard() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15releaseKeyboardEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:442
// index:0
// Public Visibility=Default Availability=Available
// [4] int grabShortcut(const QKeySequence &, Qt::ShortcutContext)
func (this *QWidget) GrabShortcut(key *qtgui.QKeySequence, context int) int {
	var convArg0 = key.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:443
// index:0
// Public Visibility=Default Availability=Available
// [-2] void releaseShortcut(int)
func (this *QWidget) ReleaseShortcut(id int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15releaseShortcutEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:444
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutEnabled(int, _Bool)
func (this *QWidget) SetShortcutEnabled(id int, enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setShortcutEnabledEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:445
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setShortcutAutoRepeat(int, _Bool)
func (this *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21setShortcutAutoRepeatEib", qtrt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:447
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * mouseGrabber()
func (this *QWidget) MouseGrabber() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12mouseGrabberEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QWidget_MouseGrabber() *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.MouseGrabber()
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:448
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * keyboardGrabber()
func (this *QWidget) KeyboardGrabber() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15keyboardGrabberEv", qtrt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QWidget_KeyboardGrabber() *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.KeyboardGrabber()
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:451
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool updatesEnabled()
func (this *QWidget) UpdatesEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14updatesEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:452
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUpdatesEnabled(_Bool)
func (this *QWidget) SetUpdatesEnabled(enable bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17setUpdatesEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:455
// index:0
// Public Visibility=Default Availability=Available
// [8] QGraphicsProxyWidget * graphicsProxyWidget()
func (this *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget19graphicsProxyWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:459
// index:0
// Public Visibility=Default Availability=Available
// [-2] void update()
func (this *QWidget) Update() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:463
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void update(int, int, int, int)
func (this *QWidget) Update_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:464
// index:2
// Public Visibility=Default Availability=Available
// [-2] void update(const QRect &)
func (this *QWidget) Update_2(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:465
// index:3
// Public Visibility=Default Availability=Available
// [-2] void update(const QRegion &)
func (this *QWidget) Update_3(arg0 *qtgui.QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6updateERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:460
// index:0
// Public Visibility=Default Availability=Available
// [-2] void repaint()
func (this *QWidget) Repaint() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:467
// index:1
// Public Visibility=Default Availability=Available
// [-2] void repaint(int, int, int, int)
func (this *QWidget) Repaint_1(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:468
// index:2
// Public Visibility=Default Availability=Available
// [-2] void repaint(const QRect &)
func (this *QWidget) Repaint_2(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:469
// index:3
// Public Visibility=Default Availability=Available
// [-2] void repaint(const QRegion &)
func (this *QWidget) Repaint_3(arg0 *qtgui.QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7repaintERK7QRegion", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:474
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void setVisible(_Bool)
func (this *QWidget) SetVisible(visible bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10setVisibleEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:475
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHidden(_Bool)
func (this *QWidget) SetHidden(hidden bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setHiddenEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hidden)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:476
// index:0
// Public Visibility=Default Availability=Available
// [-2] void show()
func (this *QWidget) Show() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4showEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:477
// index:0
// Public Visibility=Default Availability=Available
// [-2] void hide()
func (this *QWidget) Hide() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4hideEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:479
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMinimized()
func (this *QWidget) ShowMinimized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13showMinimizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:480
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showMaximized()
func (this *QWidget) ShowMaximized() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13showMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:481
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showFullScreen()
func (this *QWidget) ShowFullScreen() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14showFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:482
// index:0
// Public Visibility=Default Availability=Available
// [-2] void showNormal()
func (this *QWidget) ShowNormal() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10showNormalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:484
// index:0
// Public Visibility=Default Availability=Available
// [1] bool close()
func (this *QWidget) Close() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5closeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:485
// index:0
// Public Visibility=Default Availability=Available
// [-2] void raise()
func (this *QWidget) Raise() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5raiseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:486
// index:0
// Public Visibility=Default Availability=Available
// [-2] void lower()
func (this *QWidget) Lower() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5lowerEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:489
// index:0
// Public Visibility=Default Availability=Available
// [-2] void stackUnder(QWidget *)
func (this *QWidget) StackUnder(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10stackUnderEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:490
// index:0
// Public Visibility=Default Availability=Available
// [-2] void move(int, int)
func (this *QWidget) Move(x int, y int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4moveEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:491
// index:1
// Public Visibility=Default Availability=Available
// [-2] void move(const QPoint &)
func (this *QWidget) Move_1(arg0 *qtcore.QPoint) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4moveERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:492
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resize(int, int)
func (this *QWidget) Resize(w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6resizeEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:493
// index:1
// Public Visibility=Default Availability=Available
// [-2] void resize(const QSize &)
func (this *QWidget) Resize_1(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6resizeERK5QSize", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:494
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setGeometry(int, int, int, int)
func (this *QWidget) SetGeometry(x int, y int, w int, h int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setGeometryEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:495
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setGeometry(const QRect &)
func (this *QWidget) SetGeometry_1(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11setGeometryERK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:496
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray saveGeometry()
func (this *QWidget) SaveGeometry() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12saveGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:497
// index:0
// Public Visibility=Default Availability=Available
// [1] bool restoreGeometry(const QByteArray &)
func (this *QWidget) RestoreGeometry(geometry *qtcore.QByteArray) bool {
	var convArg0 = geometry.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15restoreGeometryERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:498
// index:0
// Public Visibility=Default Availability=Available
// [-2] void adjustSize()
func (this *QWidget) AdjustSize() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10adjustSizeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:499
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisible()
func (this *QWidget) IsVisible() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget9isVisibleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:500
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isVisibleTo(const QWidget *)
func (this *QWidget) IsVisibleTo(arg0 *QWidget /*777 const QWidget **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isVisibleToEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:501
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isHidden()
func (this *QWidget) IsHidden() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8isHiddenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:503
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMinimized()
func (this *QWidget) IsMinimized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isMinimizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:504
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMaximized()
func (this *QWidget) IsMaximized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11isMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:505
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isFullScreen()
func (this *QWidget) IsFullScreen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12isFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:507
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::WindowStates windowState()
func (this *QWidget) WindowState() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11windowStateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:508
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowState(Qt::WindowStates)
func (this *QWidget) SetWindowState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setWindowStateE6QFlagsIN2Qt11WindowStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:509
// index:0
// Public Visibility=Default Availability=Available
// [-2] void overrideWindowState(Qt::WindowStates)
func (this *QWidget) OverrideWindowState(state int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget19overrideWindowStateE6QFlagsIN2Qt11WindowStateEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:511
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize sizeHint()
func (this *QWidget) SizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget8sizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:512
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QSize minimumSizeHint()
func (this *QWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15minimumSizeHintEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:514
// index:0
// Public Visibility=Default Availability=Available
// [4] QSizePolicy sizePolicy()
func (this *QWidget) SizePolicy() *QSizePolicy /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10sizePolicyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQSizePolicy)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:515
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy)
func (this *QWidget) SetSizePolicy(arg0 *QSizePolicy /*123*/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyE11QSizePolicy", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:516
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy)
func (this *QWidget) SetSizePolicy_1(horizontal int, vertical int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyEN11QSizePolicy6PolicyES1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:517
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int heightForWidth(int)
func (this *QWidget) HeightForWidth(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14heightForWidthEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:518
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool hasHeightForWidth()
func (this *QWidget) HasHeightForWidth() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget17hasHeightForWidthEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:520
// index:0
// Public Visibility=Default Availability=Available
// [8] QRegion visibleRegion()
func (this *QWidget) VisibleRegion() *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13visibleRegionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:522
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(int, int, int, int)
func (this *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsEiiii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:523
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setContentsMargins(const QMargins &)
func (this *QWidget) SetContentsMargins_1(margins *qtcore.QMargins) {
	var convArg0 = margins.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsERK8QMargins", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:524
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getContentsMargins(int *, int *, int *, int *)
func (this *QWidget) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:525
// index:0
// Public Visibility=Default Availability=Available
// [16] QMargins contentsMargins()
func (this *QWidget) ContentsMargins() *qtcore.QMargins /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget15contentsMarginsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQMargins)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:527
// index:0
// Public Visibility=Default Availability=Available
// [16] QRect contentsRect()
func (this *QWidget) ContentsRect() *qtcore.QRect /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12contentsRectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQRect)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:530
// index:0
// Public Visibility=Default Availability=Available
// [8] QLayout * layout()
func (this *QWidget) Layout() *QLayout /*777 QLayout **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6layoutEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:531
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayout(QLayout *)
func (this *QWidget) SetLayout(arg0 *QLayout /*777 QLayout **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setLayoutEP7QLayout", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:532
// index:0
// Public Visibility=Default Availability=Available
// [-2] void updateGeometry()
func (this *QWidget) UpdateGeometry() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14updateGeometryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:534
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setParent(QWidget *)
func (this *QWidget) SetParent(parent *QWidget /*777 QWidget **/) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:535
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setParent(QWidget *, Qt::WindowFlags)
func (this *QWidget) SetParent_1(parent *QWidget /*777 QWidget **/, f int) {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:537
// index:0
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int)
func (this *QWidget) Scroll(dx int, dy int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6scrollEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:538
// index:1
// Public Visibility=Default Availability=Available
// [-2] void scroll(int, int, const QRect &)
func (this *QWidget) Scroll_1(dx int, dy int, arg2 *qtcore.QRect) {
	var convArg2 = arg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6scrollEiiRK5QRect", qtrt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:542
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * focusWidget()
func (this *QWidget) FocusWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11focusWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:543
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * nextInFocusChain()
func (this *QWidget) NextInFocusChain() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16nextInFocusChainEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:544
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * previousInFocusChain()
func (this *QWidget) PreviousInFocusChain() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget20previousInFocusChainEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:547
// index:0
// Public Visibility=Default Availability=Available
// [1] bool acceptDrops()
func (this *QWidget) AcceptDrops() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11acceptDropsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:548
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAcceptDrops(_Bool)
func (this *QWidget) SetAcceptDrops(on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setAcceptDropsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:552
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addAction(QAction *)
func (this *QWidget) AddAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9addActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:560
// index:0
// Public Visibility=Default Availability=Available
// [-2] void insertAction(QAction *, QAction *)
func (this *QWidget) InsertAction(before *QAction /*777 QAction **/, action *QAction /*777 QAction **/) {
	var convArg0 = before.GetCthis()
	var convArg1 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12insertActionEP7QActionS1_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:561
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeAction(QAction *)
func (this *QWidget) RemoveAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12removeActionEP7QAction", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:565
// index:0
// Public Visibility=Default Availability=Available
// [8] QWidget * parentWidget()
func (this *QWidget) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12parentWidgetEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:567
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlags(Qt::WindowFlags)
func (this *QWidget) SetWindowFlags(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14setWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:568
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::WindowFlags windowFlags()
func (this *QWidget) WindowFlags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11windowFlagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:569
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindowFlag(Qt::WindowType, _Bool)
func (this *QWidget) SetWindowFlag(arg0 int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13setWindowFlagEN2Qt10WindowTypeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:570
// index:0
// Public Visibility=Default Availability=Available
// [-2] void overrideWindowFlags(Qt::WindowFlags)
func (this *QWidget) OverrideWindowFlags(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget19overrideWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:572
// index:0
// Public inline Visibility=Default Availability=Available
// [4] Qt::WindowType windowType()
func (this *QWidget) WindowType() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10windowTypeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:574
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * find(WId)
func (this *QWidget) Find(arg0 uint64) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget4findEy", qtrt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QWidget_Find(arg0 uint64) *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.Find(arg0)
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:575
// index:0
// Public inline Visibility=Default Availability=Available
// [8] QWidget * childAt(int, int)
func (this *QWidget) ChildAt(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7childAtEii", qtrt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:576
// index:1
// Public Visibility=Default Availability=Available
// [8] QWidget * childAt(const QPoint &)
func (this *QWidget) ChildAt_1(p *qtcore.QPoint) *QWidget /*777 QWidget **/ {
	var convArg0 = p.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget7childAtERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:578
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(Qt::WidgetAttribute, _Bool)
func (this *QWidget) SetAttribute(arg0 int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12setAttributeEN2Qt15WidgetAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:579
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool testAttribute(Qt::WidgetAttribute)
func (this *QWidget) TestAttribute(arg0 int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13testAttributeEN2Qt15WidgetAttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:581
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QPaintEngine * paintEngine()
func (this *QWidget) PaintEngine() *qtgui.QPaintEngine /*777 QPaintEngine **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11paintEngineEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:583
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ensurePolished()
func (this *QWidget) EnsurePolished() {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget14ensurePolishedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:585
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isAncestorOf(const QWidget *)
func (this *QWidget) IsAncestorOf(child *QWidget /*777 const QWidget **/) bool {
	var convArg0 = child.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12isAncestorOfEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:592
// index:0
// Public Visibility=Default Availability=Available
// [1] bool autoFillBackground()
func (this *QWidget) AutoFillBackground() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget18autoFillBackgroundEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:593
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAutoFillBackground(_Bool)
func (this *QWidget) SetAutoFillBackground(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21setAutoFillBackgroundEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:595
// index:0
// Public Visibility=Default Availability=Available
// [8] QBackingStore * backingStore()
func (this *QWidget) BackingStore() *qtgui.QBackingStore /*777 QBackingStore **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12backingStoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBackingStoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:597
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * windowHandle()
func (this *QWidget) WindowHandle() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget12windowHandleEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:599
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWidget * createWindowContainer(QWindow *, QWidget *, Qt::WindowFlags)
func (this *QWidget) CreateWindowContainer(window *qtgui.QWindow /*777 QWindow **/, parent *QWidget /*777 QWidget **/, flags int) *QWidget /*777 QWidget **/ {
	var convArg0 = window.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21createWindowContainerEP7QWindowPS_6QFlagsIN2Qt10WindowTypeEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, flags)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}
func QWidget_CreateWindowContainer(window *qtgui.QWindow /*777 QWindow **/, parent *QWidget /*777 QWidget **/, flags int) *QWidget /*777 QWidget **/ {
	var nilthis *QWidget
	rv := nilthis.CreateWindowContainer(window, parent, flags)
	return rv
}

// /usr/include/qt/QtWidgets/qwidget.h:604
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowTitleChanged(const QString &)
func (this *QWidget) WindowTitleChanged(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18windowTitleChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:605
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowIconChanged(const QIcon &)
func (this *QWidget) WindowIconChanged(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17windowIconChangedERK5QIcon", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:606
// index:0
// Public Visibility=Default Availability=Available
// [-2] void windowIconTextChanged(const QString &)
func (this *QWidget) WindowIconTextChanged(iconText *qtcore.QString) {
	var convArg0 = iconText.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21windowIconTextChangedERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:607
// index:0
// Public Visibility=Default Availability=Available
// [-2] void customContextMenuRequested(const QPoint &)
func (this *QWidget) CustomContextMenuRequested(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget26customContextMenuRequestedERK6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:611
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool event(QEvent *)
func (this *QWidget) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget5eventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:612
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mousePressEvent(QMouseEvent *)
func (this *QWidget) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15mousePressEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:613
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseReleaseEvent(QMouseEvent *)
func (this *QWidget) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget17mouseReleaseEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:614
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseDoubleClickEvent(QMouseEvent *)
func (this *QWidget) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:615
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void mouseMoveEvent(QMouseEvent *)
func (this *QWidget) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14mouseMoveEventEP11QMouseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:617
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void wheelEvent(QWheelEvent *)
func (this *QWidget) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10wheelEventEP11QWheelEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:619
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyPressEvent(QKeyEvent *)
func (this *QWidget) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13keyPressEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:620
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void keyReleaseEvent(QKeyEvent *)
func (this *QWidget) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget15keyReleaseEventEP9QKeyEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:621
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusInEvent(QFocusEvent *)
func (this *QWidget) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget12focusInEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:622
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void focusOutEvent(QFocusEvent *)
func (this *QWidget) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13focusOutEventEP11QFocusEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:623
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void enterEvent(QEvent *)
func (this *QWidget) EnterEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10enterEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:624
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void leaveEvent(QEvent *)
func (this *QWidget) LeaveEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10leaveEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:625
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void paintEvent(QPaintEvent *)
func (this *QWidget) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10paintEventEP11QPaintEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:626
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void moveEvent(QMoveEvent *)
func (this *QWidget) MoveEvent(event *qtgui.QMoveEvent /*777 QMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9moveEventEP10QMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:627
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void resizeEvent(QResizeEvent *)
func (this *QWidget) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11resizeEventEP12QResizeEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:628
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void closeEvent(QCloseEvent *)
func (this *QWidget) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget10closeEventEP11QCloseEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:630
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void contextMenuEvent(QContextMenuEvent *)
func (this *QWidget) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16contextMenuEventEP17QContextMenuEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:633
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void tabletEvent(QTabletEvent *)
func (this *QWidget) TabletEvent(event *qtgui.QTabletEvent /*777 QTabletEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11tabletEventEP12QTabletEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:636
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void actionEvent(QActionEvent *)
func (this *QWidget) ActionEvent(event *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11actionEventEP12QActionEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:640
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragEnterEvent(QDragEnterEvent *)
func (this *QWidget) DragEnterEvent(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14dragEnterEventEP15QDragEnterEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:641
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragMoveEvent(QDragMoveEvent *)
func (this *QWidget) DragMoveEvent(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget13dragMoveEventEP14QDragMoveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:642
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dragLeaveEvent(QDragLeaveEvent *)
func (this *QWidget) DragLeaveEvent(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:643
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void dropEvent(QDropEvent *)
func (this *QWidget) DropEvent(event *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9dropEventEP10QDropEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:646
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void showEvent(QShowEvent *)
func (this *QWidget) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9showEventEP10QShowEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:647
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void hideEvent(QHideEvent *)
func (this *QWidget) HideEvent(event *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget9hideEventEP10QHideEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:648
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool nativeEvent(const QByteArray &, void *, long *)
func (this *QWidget) NativeEvent(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11nativeEventERK10QByteArrayPvPl", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:651
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void changeEvent(QEvent *)
func (this *QWidget) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget11changeEventEP6QEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:653
// index:0
// Protected virtual Visibility=Default Availability=Available
// [4] int metric(enum QPaintDevice::PaintDeviceMetric)
func (this *QWidget) Metric(arg0 int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWidgets/qwidget.h:654
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void initPainter(QPainter *)
func (this *QWidget) InitPainter(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget11initPainterEP8QPainter", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:655
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPaintDevice * redirected(QPoint *)
func (this *QWidget) Redirected(offset *qtcore.QPoint /*777 QPoint **/) *qtgui.QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 = offset.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget10redirectedEP6QPoint", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:656
// index:0
// Protected virtual Visibility=Default Availability=Available
// [8] QPainter * sharedPainter()
func (this *QWidget) SharedPainter() *qtgui.QPainter /*777 QPainter **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget13sharedPainterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:658
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void inputMethodEvent(QInputMethodEvent *)
func (this *QWidget) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16inputMethodEventEP17QInputMethodEvent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:660
// index:0
// Public virtual Visibility=Default Availability=Available
// [16] QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QWidget) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16inputMethodQueryEN2Qt16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQVariant)
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:662
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::InputMethodHints inputMethodHints()
func (this *QWidget) InputMethodHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK7QWidget16inputMethodHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:663
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setInputMethodHints(Qt::InputMethodHints)
func (this *QWidget) SetInputMethodHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget19setInputMethodHintsE6QFlagsIN2Qt15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:666
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void updateMicroFocus()
func (this *QWidget) UpdateMicroFocus() {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget16updateMicroFocusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:669
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void create(WId, _Bool, _Bool)
func (this *QWidget) Create(arg0 uint64, initializeWindow bool, destroyOldWindow bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget6createEybb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), arg0, initializeWindow, destroyOldWindow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:671
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void destroy(_Bool, _Bool)
func (this *QWidget) Destroy(destroyWindow bool, destroySubWindows bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget7destroyEbb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), destroyWindow, destroySubWindows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:675
// index:0
// Protected virtual Visibility=Default Availability=Available
// [1] bool focusNextPrevChild(_Bool)
func (this *QWidget) FocusNextPrevChild(next bool) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18focusNextPrevChildEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:676
// index:0
// Protected inline Visibility=Default Availability=Available
// [1] bool focusNextChild()
func (this *QWidget) FocusNextChild() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget14focusNextChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:677
// index:0
// Protected inline Visibility=Default Availability=Available
// [1] bool focusPreviousChild()
func (this *QWidget) FocusPreviousChild() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN7QWidget18focusPreviousChildEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QWidget__RenderFlag = int

const QWidget__DrawWindowBackground QWidget__RenderFlag = 1
const QWidget__DrawChildren QWidget__RenderFlag = 2
const QWidget__IgnoreMask QWidget__RenderFlag = 4

//  body block end
