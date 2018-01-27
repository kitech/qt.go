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
import "qt.go/cffiqt"
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
		ffiqt.KeepMe()
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
// Public virtual
// const QMetaObject * metaObject()
func (this *QWidget) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:214
// index:0
// Public
// void QWidget(QWidget *, Qt::WindowFlags)
func NewQWidget(parent *QWidget /*777 QWidget **/, f int) *QWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidgetC2EPS_6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, f)
	gopp.ErrPrint(err, rv)
	gothis := NewQWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qwidget.h:215
// index:0
// Public virtual
// void ~QWidget()
func DeleteQWidget(*QWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:217
// index:0
// Public virtual
// int devType()
func (this *QWidget) DevType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7devTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:219
// index:0
// Public
// WId winId()
func (this *QWidget) WinId() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5winIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:220
// index:0
// Public
// void createWinId()
func (this *QWidget) CreateWinId() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11createWinIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:221
// index:0
// Public inline
// WId internalWinId()
func (this *QWidget) InternalWinId() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13internalWinIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:222
// index:0
// Public
// WId effectiveWinId()
func (this *QWidget) EffectiveWinId() uint64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14effectiveWinIdEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:225
// index:0
// Public
// QStyle * style()
func (this *QWidget) Style() *QStyle /*777 QStyle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:226
// index:0
// Public
// void setStyle(QStyle *)
func (this *QWidget) SetStyle(arg0 *QStyle /*777 QStyle **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget8setStyleEP6QStyle", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:229
// index:0
// Public
// bool isTopLevel()
func (this *QWidget) IsTopLevel() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10isTopLevelEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:230
// index:0
// Public
// bool isWindow()
func (this *QWidget) IsWindow() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8isWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:232
// index:0
// Public
// bool isModal()
func (this *QWidget) IsModal() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7isModalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:233
// index:0
// Public
// Qt::WindowModality windowModality()
func (this *QWidget) WindowModality() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14windowModalityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:234
// index:0
// Public
// void setWindowModality(Qt::WindowModality)
func (this *QWidget) SetWindowModality(windowModality int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowModalityEN2Qt14WindowModalityE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), windowModality)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:236
// index:0
// Public
// bool isEnabled()
func (this *QWidget) IsEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9isEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:237
// index:0
// Public
// bool isEnabledTo(const QWidget *)
func (this *QWidget) IsEnabledTo(arg0 *QWidget /*777 const QWidget **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isEnabledToEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:238
// index:0
// Public
// bool isEnabledToTLW()
func (this *QWidget) IsEnabledToTLW() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14isEnabledToTLWEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:241
// index:0
// Public
// void setEnabled(bool)
func (this *QWidget) SetEnabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:242
// index:0
// Public
// void setDisabled(bool)
func (this *QWidget) SetDisabled(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setDisabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:243
// index:0
// Public
// void setWindowModified(bool)
func (this *QWidget) SetWindowModified(arg0 bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowModifiedEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:248
// index:0
// Public
// QRect frameGeometry()
func (this *QWidget) FrameGeometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13frameGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:249
// index:0
// Public
// const QRect & geometry()
func (this *QWidget) Geometry() *qtcore.QRect {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8geometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:250
// index:0
// Public
// QRect normalGeometry()
func (this *QWidget) NormalGeometry() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14normalGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:252
// index:0
// Public
// int x()
func (this *QWidget) X() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget1xEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:253
// index:0
// Public
// int y()
func (this *QWidget) Y() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget1yEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:254
// index:0
// Public
// QPoint pos()
func (this *QWidget) Pos() *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget3posEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:255
// index:0
// Public
// QSize frameSize()
func (this *QWidget) FrameSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9frameSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:256
// index:0
// Public
// QSize size()
func (this *QWidget) Size() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:257
// index:0
// Public inline
// int width()
func (this *QWidget) Width() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5widthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:258
// index:0
// Public inline
// int height()
func (this *QWidget) Height() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6heightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:259
// index:0
// Public inline
// QRect rect()
func (this *QWidget) Rect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:260
// index:0
// Public
// QRect childrenRect()
func (this *QWidget) ChildrenRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12childrenRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:261
// index:0
// Public
// QRegion childrenRegion()
func (this *QWidget) ChildrenRegion() *qtgui.QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14childrenRegionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:263
// index:0
// Public
// QSize minimumSize()
func (this *QWidget) MinimumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11minimumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:264
// index:0
// Public
// QSize maximumSize()
func (this *QWidget) MaximumSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11maximumSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:265
// index:0
// Public
// int minimumWidth()
func (this *QWidget) MinimumWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12minimumWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:266
// index:0
// Public
// int minimumHeight()
func (this *QWidget) MinimumHeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13minimumHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:267
// index:0
// Public
// int maximumWidth()
func (this *QWidget) MaximumWidth() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12maximumWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:268
// index:0
// Public
// int maximumHeight()
func (this *QWidget) MaximumHeight() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13maximumHeightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:269
// index:0
// Public
// void setMinimumSize(const QSize &)
func (this *QWidget) SetMinimumSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:270
// index:1
// Public
// void setMinimumSize(int, int)
func (this *QWidget) SetMinimumSize_1(minw int, minh int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMinimumSizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minw, minh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:271
// index:0
// Public
// void setMaximumSize(const QSize &)
func (this *QWidget) SetMaximumSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:272
// index:1
// Public
// void setMaximumSize(int, int)
func (this *QWidget) SetMaximumSize_1(maxw int, maxh int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setMaximumSizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maxw, maxh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:273
// index:0
// Public
// void setMinimumWidth(int)
func (this *QWidget) SetMinimumWidth(minw int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15setMinimumWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minw)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:274
// index:0
// Public
// void setMinimumHeight(int)
func (this *QWidget) SetMinimumHeight(minh int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setMinimumHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), minh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:275
// index:0
// Public
// void setMaximumWidth(int)
func (this *QWidget) SetMaximumWidth(maxw int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15setMaximumWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maxw)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:276
// index:0
// Public
// void setMaximumHeight(int)
func (this *QWidget) SetMaximumHeight(maxh int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setMaximumHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), maxh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:282
// index:0
// Public
// QSize sizeIncrement()
func (this *QWidget) SizeIncrement() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13sizeIncrementEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:283
// index:0
// Public
// void setSizeIncrement(const QSize &)
func (this *QWidget) SetSizeIncrement(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:284
// index:1
// Public
// void setSizeIncrement(int, int)
func (this *QWidget) SetSizeIncrement_1(w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setSizeIncrementEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:285
// index:0
// Public
// QSize baseSize()
func (this *QWidget) BaseSize() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8baseSizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:286
// index:0
// Public
// void setBaseSize(const QSize &)
func (this *QWidget) SetBaseSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:287
// index:1
// Public
// void setBaseSize(int, int)
func (this *QWidget) SetBaseSize_1(basew int, baseh int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setBaseSizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), basew, baseh)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:289
// index:0
// Public
// void setFixedSize(const QSize &)
func (this *QWidget) SetFixedSize(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:290
// index:1
// Public
// void setFixedSize(int, int)
func (this *QWidget) SetFixedSize_1(w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setFixedSizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:291
// index:0
// Public
// void setFixedWidth(int)
func (this *QWidget) SetFixedWidth(w int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setFixedWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:292
// index:0
// Public
// void setFixedHeight(int)
func (this *QWidget) SetFixedHeight(h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setFixedHeightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:296
// index:0
// Public
// QPoint mapToGlobal(const QPoint &)
func (this *QWidget) MapToGlobal(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11mapToGlobalERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:297
// index:0
// Public
// QPoint mapFromGlobal(const QPoint &)
func (this *QWidget) MapFromGlobal(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13mapFromGlobalERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:298
// index:0
// Public
// QPoint mapToParent(const QPoint &)
func (this *QWidget) MapToParent(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11mapToParentERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:299
// index:0
// Public
// QPoint mapFromParent(const QPoint &)
func (this *QWidget) MapFromParent(arg0 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13mapFromParentERK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:300
// index:0
// Public
// QPoint mapTo(const QWidget *, const QPoint &)
func (this *QWidget) MapTo(arg0 *QWidget /*777 const QWidget **/, arg1 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget5mapToEPKS_RK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:301
// index:0
// Public
// QPoint mapFrom(const QWidget *, const QPoint &)
func (this *QWidget) MapFrom(arg0 *QWidget /*777 const QWidget **/, arg1 *qtcore.QPoint) *qtcore.QPoint /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = arg0.GetCthis()
	var convArg1 = arg1.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7mapFromEPKS_RK6QPoint", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQPointFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:303
// index:0
// Public
// QWidget * window()
func (this *QWidget) Window() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6windowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:304
// index:0
// Public
// QWidget * nativeParentWidget()
func (this *QWidget) NativeParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget18nativeParentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:305
// index:0
// Public inline
// QWidget * topLevelWidget()
func (this *QWidget) TopLevelWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14topLevelWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:308
// index:0
// Public
// const QPalette & palette()
func (this *QWidget) Palette() *qtgui.QPalette {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7paletteEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:309
// index:0
// Public
// void setPalette(const QPalette &)
func (this *QWidget) SetPalette(arg0 *qtgui.QPalette) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setPaletteERK8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:311
// index:0
// Public
// void setBackgroundRole(QPalette::ColorRole)
func (this *QWidget) SetBackgroundRole(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setBackgroundRoleEN8QPalette9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:312
// index:0
// Public
// QPalette::ColorRole backgroundRole()
func (this *QWidget) BackgroundRole() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14backgroundRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:314
// index:0
// Public
// void setForegroundRole(QPalette::ColorRole)
func (this *QWidget) SetForegroundRole(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setForegroundRoleEN8QPalette9ColorRoleE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:315
// index:0
// Public
// QPalette::ColorRole foregroundRole()
func (this *QWidget) ForegroundRole() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14foregroundRoleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:317
// index:0
// Public
// const QFont & font()
func (this *QWidget) Font() *qtgui.QFont {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4fontEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:318
// index:0
// Public
// void setFont(const QFont &)
func (this *QWidget) SetFont(arg0 *qtgui.QFont) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:319
// index:0
// Public
// QFontMetrics fontMetrics()
func (this *QWidget) FontMetrics() *qtgui.QFontMetrics /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11fontMetricsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQFontMetricsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:320
// index:0
// Public
// QFontInfo fontInfo()
func (this *QWidget) FontInfo() *qtgui.QFontInfo /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8fontInfoEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQFontInfoFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:323
// index:0
// Public
// QCursor cursor()
func (this *QWidget) Cursor() *qtgui.QCursor /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6cursorEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:324
// index:0
// Public
// void setCursor(const QCursor &)
func (this *QWidget) SetCursor(arg0 *qtgui.QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setCursorERK7QCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:325
// index:0
// Public
// void unsetCursor()
func (this *QWidget) UnsetCursor() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11unsetCursorEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:328
// index:0
// Public
// void setMouseTracking(bool)
func (this *QWidget) SetMouseTracking(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setMouseTrackingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:329
// index:0
// Public
// bool hasMouseTracking()
func (this *QWidget) HasMouseTracking() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16hasMouseTrackingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:330
// index:0
// Public
// bool underMouse()
func (this *QWidget) UnderMouse() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10underMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:332
// index:0
// Public
// void setTabletTracking(bool)
func (this *QWidget) SetTabletTracking(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setTabletTrackingEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:333
// index:0
// Public
// bool hasTabletTracking()
func (this *QWidget) HasTabletTracking() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget17hasTabletTrackingEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:335
// index:0
// Public
// void setMask(const QBitmap &)
func (this *QWidget) SetMask(arg0 *qtgui.QBitmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QBitmap", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:336
// index:1
// Public
// void setMask(const QRegion &)
func (this *QWidget) SetMask_1(arg0 *qtgui.QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7setMaskERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:337
// index:0
// Public
// QRegion mask()
func (this *QWidget) Mask() *qtgui.QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget4maskEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:338
// index:0
// Public
// void clearMask()
func (this *QWidget) ClearMask() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9clearMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:340
// index:0
// Public
// void render(QPaintDevice *, const QPoint &, const QRegion &, QWidget::RenderFlags)
func (this *QWidget) Render(target *qtgui.QPaintDevice /*777 QPaintDevice **/, targetOffset *qtcore.QPoint, sourceRegion *qtgui.QRegion, renderFlags int) {
	var convArg0 = target.GetCthis()
	var convArg1 = targetOffset.GetCthis()
	var convArg2 = sourceRegion.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6renderEP12QPaintDeviceRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:344
// index:1
// Public
// void render(QPainter *, const QPoint &, const QRegion &, QWidget::RenderFlags)
func (this *QWidget) Render_1(painter *qtgui.QPainter /*777 QPainter **/, targetOffset *qtcore.QPoint, sourceRegion *qtgui.QRegion, renderFlags int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = targetOffset.GetCthis()
	var convArg2 = sourceRegion.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6renderEP8QPainterRK6QPointRK7QRegion6QFlagsINS_10RenderFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, renderFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:348
// index:0
// Public
// QPixmap grab(const QRect &)
func (this *QWidget) Grab(rectangle *qtcore.QRect) *qtgui.QPixmap /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = rectangle.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4grabERK5QRect", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:351
// index:0
// Public
// QGraphicsEffect * graphicsEffect()
func (this *QWidget) GraphicsEffect() *QGraphicsEffect /*777 QGraphicsEffect **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14graphicsEffectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsEffectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:352
// index:0
// Public
// void setGraphicsEffect(QGraphicsEffect *)
func (this *QWidget) SetGraphicsEffect(effect *QGraphicsEffect /*777 QGraphicsEffect **/) {
	var convArg0 = effect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setGraphicsEffectEP15QGraphicsEffect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:356
// index:0
// Public
// void grabGesture(Qt::GestureType, Qt::GestureFlags)
func (this *QWidget) GrabGesture(type_ int, flags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11grabGestureEN2Qt11GestureTypeE6QFlagsINS0_11GestureFlagEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_, flags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:357
// index:0
// Public
// void ungrabGesture(Qt::GestureType)
func (this *QWidget) UngrabGesture(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13ungrabGestureEN2Qt11GestureTypeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:361
// index:0
// Public
// void setWindowTitle(const QString &)
func (this *QWidget) SetWindowTitle(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setWindowTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:363
// index:0
// Public
// void setStyleSheet(const QString &)
func (this *QWidget) SetStyleSheet(styleSheet *qtcore.QString) {
	var convArg0 = styleSheet.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setStyleSheetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:367
// index:0
// Public
// QString styleSheet()
func (this *QWidget) StyleSheet() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10styleSheetEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:369
// index:0
// Public
// QString windowTitle()
func (this *QWidget) WindowTitle() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11windowTitleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:370
// index:0
// Public
// void setWindowIcon(const QIcon &)
func (this *QWidget) SetWindowIcon(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:371
// index:0
// Public
// QIcon windowIcon()
func (this *QWidget) WindowIcon() *qtgui.QIcon /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10windowIconEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:372
// index:0
// Public
// void setWindowIconText(const QString &)
func (this *QWidget) SetWindowIconText(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowIconTextERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:373
// index:0
// Public
// QString windowIconText()
func (this *QWidget) WindowIconText() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14windowIconTextEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:374
// index:0
// Public
// void setWindowRole(const QString &)
func (this *QWidget) SetWindowRole(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setWindowRoleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:375
// index:0
// Public
// QString windowRole()
func (this *QWidget) WindowRole() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10windowRoleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:376
// index:0
// Public
// void setWindowFilePath(const QString &)
func (this *QWidget) SetWindowFilePath(filePath *qtcore.QString) {
	var convArg0 = filePath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setWindowFilePathERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:377
// index:0
// Public
// QString windowFilePath()
func (this *QWidget) WindowFilePath() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14windowFilePathEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:379
// index:0
// Public
// void setWindowOpacity(qreal)
func (this *QWidget) SetWindowOpacity(level float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16setWindowOpacityEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), level)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:380
// index:0
// Public
// qreal windowOpacity()
func (this *QWidget) WindowOpacity() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13windowOpacityEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return float64(rv) // 222
}

// /usr/include/qt/QtWidgets/qwidget.h:382
// index:0
// Public
// bool isWindowModified()
func (this *QWidget) IsWindowModified() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16isWindowModifiedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:384
// index:0
// Public
// void setToolTip(const QString &)
func (this *QWidget) SetToolTip(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setToolTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:385
// index:0
// Public
// QString toolTip()
func (this *QWidget) ToolTip() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7toolTipEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:386
// index:0
// Public
// void setToolTipDuration(int)
func (this *QWidget) SetToolTipDuration(msec int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setToolTipDurationEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), msec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:387
// index:0
// Public
// int toolTipDuration()
func (this *QWidget) ToolTipDuration() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15toolTipDurationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:390
// index:0
// Public
// void setStatusTip(const QString &)
func (this *QWidget) SetStatusTip(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setStatusTipERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:391
// index:0
// Public
// QString statusTip()
func (this *QWidget) StatusTip() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9statusTipEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:394
// index:0
// Public
// void setWhatsThis(const QString &)
func (this *QWidget) SetWhatsThis(arg0 *qtcore.QString) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setWhatsThisERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:395
// index:0
// Public
// QString whatsThis()
func (this *QWidget) WhatsThis() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9whatsThisEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:398
// index:0
// Public
// QString accessibleName()
func (this *QWidget) AccessibleName() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14accessibleNameEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:399
// index:0
// Public
// void setAccessibleName(const QString &)
func (this *QWidget) SetAccessibleName(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setAccessibleNameERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:400
// index:0
// Public
// QString accessibleDescription()
func (this *QWidget) AccessibleDescription() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget21accessibleDescriptionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:401
// index:0
// Public
// void setAccessibleDescription(const QString &)
func (this *QWidget) SetAccessibleDescription(description *qtcore.QString) {
	var convArg0 = description.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget24setAccessibleDescriptionERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:404
// index:0
// Public
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QWidget) SetLayoutDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:405
// index:0
// Public
// Qt::LayoutDirection layoutDirection()
func (this *QWidget) LayoutDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:406
// index:0
// Public
// void unsetLayoutDirection()
func (this *QWidget) UnsetLayoutDirection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget20unsetLayoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:408
// index:0
// Public
// void setLocale(const QLocale &)
func (this *QWidget) SetLocale(locale *qtcore.QLocale) {
	var convArg0 = locale.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setLocaleERK7QLocale", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:409
// index:0
// Public
// QLocale locale()
func (this *QWidget) Locale() *qtcore.QLocale /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6localeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQLocaleFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:410
// index:0
// Public
// void unsetLocale()
func (this *QWidget) UnsetLocale() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11unsetLocaleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:412
// index:0
// Public inline
// bool isRightToLeft()
func (this *QWidget) IsRightToLeft() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13isRightToLeftEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:413
// index:0
// Public inline
// bool isLeftToRight()
func (this *QWidget) IsLeftToRight() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13isLeftToRightEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:416
// index:0
// Public inline
// void setFocus()
func (this *QWidget) SetFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget8setFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:423
// index:1
// Public
// void setFocus(Qt::FocusReason)
func (this *QWidget) SetFocus_1(reason int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget8setFocusEN2Qt11FocusReasonE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), reason)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:419
// index:0
// Public
// bool isActiveWindow()
func (this *QWidget) IsActiveWindow() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14isActiveWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:420
// index:0
// Public
// void activateWindow()
func (this *QWidget) ActivateWindow() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14activateWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:421
// index:0
// Public
// void clearFocus()
func (this *QWidget) ClearFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10clearFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:424
// index:0
// Public
// Qt::FocusPolicy focusPolicy()
func (this *QWidget) FocusPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11focusPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:425
// index:0
// Public
// void setFocusPolicy(Qt::FocusPolicy)
func (this *QWidget) SetFocusPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setFocusPolicyEN2Qt11FocusPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:426
// index:0
// Public
// bool hasFocus()
func (this *QWidget) HasFocus() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8hasFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:427
// index:0
// Public static
// void setTabOrder(QWidget *, QWidget *)
func (this *QWidget) SetTabOrder(arg0 *QWidget /*777 QWidget **/, arg1 *QWidget /*777 QWidget **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setTabOrderEPS_S0_", ffiqt.FFI_TYPE_POINTER, arg0, arg1)
	gopp.ErrPrint(err, rv)
}
func QWidget_SetTabOrder(arg0 *QWidget /*777 QWidget **/, arg1 *QWidget /*777 QWidget **/) {
	var nilthis *QWidget
	nilthis.SetTabOrder(arg0, arg1)
}

// /usr/include/qt/QtWidgets/qwidget.h:428
// index:0
// Public
// void setFocusProxy(QWidget *)
func (this *QWidget) SetFocusProxy(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setFocusProxyEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:429
// index:0
// Public
// QWidget * focusProxy()
func (this *QWidget) FocusProxy() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10focusProxyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:430
// index:0
// Public
// Qt::ContextMenuPolicy contextMenuPolicy()
func (this *QWidget) ContextMenuPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget17contextMenuPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:431
// index:0
// Public
// void setContextMenuPolicy(Qt::ContextMenuPolicy)
func (this *QWidget) SetContextMenuPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget20setContextMenuPolicyEN2Qt17ContextMenuPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:434
// index:0
// Public
// void grabMouse()
func (this *QWidget) GrabMouse() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9grabMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:436
// index:1
// Public
// void grabMouse(const QCursor &)
func (this *QWidget) GrabMouse_1(arg0 *qtgui.QCursor) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9grabMouseERK7QCursor", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:438
// index:0
// Public
// void releaseMouse()
func (this *QWidget) ReleaseMouse() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12releaseMouseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:439
// index:0
// Public
// void grabKeyboard()
func (this *QWidget) GrabKeyboard() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12grabKeyboardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:440
// index:0
// Public
// void releaseKeyboard()
func (this *QWidget) ReleaseKeyboard() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15releaseKeyboardEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:442
// index:0
// Public
// int grabShortcut(const QKeySequence &, Qt::ShortcutContext)
func (this *QWidget) GrabShortcut(key *qtgui.QKeySequence, context int) int {
	var convArg0 = key.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:443
// index:0
// Public
// void releaseShortcut(int)
func (this *QWidget) ReleaseShortcut(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15releaseShortcutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:444
// index:0
// Public
// void setShortcutEnabled(int, bool)
func (this *QWidget) SetShortcutEnabled(id int, enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setShortcutEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:445
// index:0
// Public
// void setShortcutAutoRepeat(int, bool)
func (this *QWidget) SetShortcutAutoRepeat(id int, enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21setShortcutAutoRepeatEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id, enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:447
// index:0
// Public static
// QWidget * mouseGrabber()
func (this *QWidget) MouseGrabber() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12mouseGrabberEv", ffiqt.FFI_TYPE_POINTER)
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
// Public static
// QWidget * keyboardGrabber()
func (this *QWidget) KeyboardGrabber() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15keyboardGrabberEv", ffiqt.FFI_TYPE_POINTER)
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
// Public inline
// bool updatesEnabled()
func (this *QWidget) UpdatesEnabled() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14updatesEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:452
// index:0
// Public
// void setUpdatesEnabled(bool)
func (this *QWidget) SetUpdatesEnabled(enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17setUpdatesEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enable)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:455
// index:0
// Public
// QGraphicsProxyWidget * graphicsProxyWidget()
func (this *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget /*777 QGraphicsProxyWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget19graphicsProxyWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsProxyWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:459
// index:0
// Public
// void update()
func (this *QWidget) Update() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:463
// index:1
// Public inline
// void update(int, int, int, int)
func (this *QWidget) Update_1(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:464
// index:2
// Public
// void update(const QRect &)
func (this *QWidget) Update_2(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:465
// index:3
// Public
// void update(const QRegion &)
func (this *QWidget) Update_3(arg0 *qtgui.QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6updateERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:460
// index:0
// Public
// void repaint()
func (this *QWidget) Repaint() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:467
// index:1
// Public
// void repaint(int, int, int, int)
func (this *QWidget) Repaint_1(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:468
// index:2
// Public
// void repaint(const QRect &)
func (this *QWidget) Repaint_2(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:469
// index:3
// Public
// void repaint(const QRegion &)
func (this *QWidget) Repaint_3(arg0 *qtgui.QRegion) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7repaintERK7QRegion", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:474
// index:0
// Public virtual
// void setVisible(bool)
func (this *QWidget) SetVisible(visible bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10setVisibleEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), visible)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:475
// index:0
// Public
// void setHidden(bool)
func (this *QWidget) SetHidden(hidden bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setHiddenEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), hidden)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:476
// index:0
// Public
// void show()
func (this *QWidget) Show() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4showEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:477
// index:0
// Public
// void hide()
func (this *QWidget) Hide() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4hideEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:479
// index:0
// Public
// void showMinimized()
func (this *QWidget) ShowMinimized() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13showMinimizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:480
// index:0
// Public
// void showMaximized()
func (this *QWidget) ShowMaximized() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13showMaximizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:481
// index:0
// Public
// void showFullScreen()
func (this *QWidget) ShowFullScreen() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14showFullScreenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:482
// index:0
// Public
// void showNormal()
func (this *QWidget) ShowNormal() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10showNormalEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:484
// index:0
// Public
// bool close()
func (this *QWidget) Close() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:485
// index:0
// Public
// void raise()
func (this *QWidget) Raise() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5raiseEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:486
// index:0
// Public
// void lower()
func (this *QWidget) Lower() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5lowerEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:489
// index:0
// Public
// void stackUnder(QWidget *)
func (this *QWidget) StackUnder(arg0 *QWidget /*777 QWidget **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10stackUnderEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:490
// index:0
// Public
// void move(int, int)
func (this *QWidget) Move(x int, y int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4moveEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:491
// index:1
// Public
// void move(const QPoint &)
func (this *QWidget) Move_1(arg0 *qtcore.QPoint) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4moveERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:492
// index:0
// Public
// void resize(int, int)
func (this *QWidget) Resize(w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6resizeEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:493
// index:1
// Public
// void resize(const QSize &)
func (this *QWidget) Resize_1(arg0 *qtcore.QSize) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6resizeERK5QSize", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:494
// index:0
// Public inline
// void setGeometry(int, int, int, int)
func (this *QWidget) SetGeometry(x int, y int, w int, h int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setGeometryEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:495
// index:1
// Public
// void setGeometry(const QRect &)
func (this *QWidget) SetGeometry_1(arg0 *qtcore.QRect) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11setGeometryERK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:496
// index:0
// Public
// QByteArray saveGeometry()
func (this *QWidget) SaveGeometry() *qtcore.QByteArray /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12saveGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:497
// index:0
// Public
// bool restoreGeometry(const QByteArray &)
func (this *QWidget) RestoreGeometry(geometry *qtcore.QByteArray) bool {
	var convArg0 = geometry.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15restoreGeometryERK10QByteArray", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:498
// index:0
// Public
// void adjustSize()
func (this *QWidget) AdjustSize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10adjustSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:499
// index:0
// Public
// bool isVisible()
func (this *QWidget) IsVisible() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget9isVisibleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:500
// index:0
// Public
// bool isVisibleTo(const QWidget *)
func (this *QWidget) IsVisibleTo(arg0 *QWidget /*777 const QWidget **/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isVisibleToEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:501
// index:0
// Public inline
// bool isHidden()
func (this *QWidget) IsHidden() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8isHiddenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:503
// index:0
// Public
// bool isMinimized()
func (this *QWidget) IsMinimized() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isMinimizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:504
// index:0
// Public
// bool isMaximized()
func (this *QWidget) IsMaximized() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11isMaximizedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:505
// index:0
// Public
// bool isFullScreen()
func (this *QWidget) IsFullScreen() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12isFullScreenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:507
// index:0
// Public
// Qt::WindowStates windowState()
func (this *QWidget) WindowState() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11windowStateEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:511
// index:0
// Public virtual
// QSize sizeHint()
func (this *QWidget) SizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget8sizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:512
// index:0
// Public virtual
// QSize minimumSizeHint()
func (this *QWidget) MinimumSizeHint() *qtcore.QSize /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15minimumSizeHintEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:514
// index:0
// Public
// QSizePolicy sizePolicy()
func (this *QWidget) SizePolicy() *QSizePolicy /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10sizePolicyEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := /*==*/ NewQSizePolicyFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:515
// index:0
// Public
// void setSizePolicy(QSizePolicy)
func (this *QWidget) SetSizePolicy(arg0 *QSizePolicy /*123*/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyE11QSizePolicy", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:516
// index:1
// Public inline
// void setSizePolicy(QSizePolicy::Policy, QSizePolicy::Policy)
func (this *QWidget) SetSizePolicy_1(horizontal int, vertical int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setSizePolicyEN11QSizePolicy6PolicyES1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), horizontal, vertical)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:517
// index:0
// Public virtual
// int heightForWidth(int)
func (this *QWidget) HeightForWidth(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14heightForWidthEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:518
// index:0
// Public virtual
// bool hasHeightForWidth()
func (this *QWidget) HasHeightForWidth() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget17hasHeightForWidthEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:520
// index:0
// Public
// QRegion visibleRegion()
func (this *QWidget) VisibleRegion() *qtgui.QRegion /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13visibleRegionEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:522
// index:0
// Public
// void setContentsMargins(int, int, int, int)
func (this *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:523
// index:1
// Public
// void setContentsMargins(const QMargins &)
func (this *QWidget) SetContentsMargins_1(margins *qtcore.QMargins) {
	var convArg0 = margins.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18setContentsMarginsERK8QMargins", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:524
// index:0
// Public
// void getContentsMargins(int *, int *, int *, int *)
func (this *QWidget) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget18getContentsMarginsEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:525
// index:0
// Public
// QMargins contentsMargins()
func (this *QWidget) ContentsMargins() *qtcore.QMargins /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget15contentsMarginsEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQMarginsFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:527
// index:0
// Public
// QRect contentsRect()
func (this *QWidget) ContentsRect() *qtcore.QRect /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12contentsRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:530
// index:0
// Public
// QLayout * layout()
func (this *QWidget) Layout() *QLayout /*777 QLayout **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6layoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:531
// index:0
// Public
// void setLayout(QLayout *)
func (this *QWidget) SetLayout(arg0 *QLayout /*777 QLayout **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setLayoutEP7QLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:532
// index:0
// Public
// void updateGeometry()
func (this *QWidget) UpdateGeometry() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14updateGeometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:534
// index:0
// Public
// void setParent(QWidget *)
func (this *QWidget) SetParent(parent *QWidget /*777 QWidget **/) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:535
// index:1
// Public
// void setParent(QWidget *, Qt::WindowFlags)
func (this *QWidget) SetParent_1(parent *QWidget /*777 QWidget **/, f int) {
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9setParentEPS_6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:537
// index:0
// Public
// void scroll(int, int)
func (this *QWidget) Scroll(dx int, dy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6scrollEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:538
// index:1
// Public
// void scroll(int, int, const QRect &)
func (this *QWidget) Scroll_1(dx int, dy int, arg2 *qtcore.QRect) {
	var convArg2 = arg2.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6scrollEiiRK5QRect", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), dx, dy, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:542
// index:0
// Public
// QWidget * focusWidget()
func (this *QWidget) FocusWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11focusWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:543
// index:0
// Public
// QWidget * nextInFocusChain()
func (this *QWidget) NextInFocusChain() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16nextInFocusChainEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:544
// index:0
// Public
// QWidget * previousInFocusChain()
func (this *QWidget) PreviousInFocusChain() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget20previousInFocusChainEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:547
// index:0
// Public
// bool acceptDrops()
func (this *QWidget) AcceptDrops() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11acceptDropsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:548
// index:0
// Public
// void setAcceptDrops(bool)
func (this *QWidget) SetAcceptDrops(on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setAcceptDropsEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:552
// index:0
// Public
// void addAction(QAction *)
func (this *QWidget) AddAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9addActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:560
// index:0
// Public
// void insertAction(QAction *, QAction *)
func (this *QWidget) InsertAction(before *QAction /*777 QAction **/, action *QAction /*777 QAction **/) {
	var convArg0 = before.GetCthis()
	var convArg1 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12insertActionEP7QActionS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:561
// index:0
// Public
// void removeAction(QAction *)
func (this *QWidget) RemoveAction(action *QAction /*777 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12removeActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:565
// index:0
// Public
// QWidget * parentWidget()
func (this *QWidget) ParentWidget() *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12parentWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:567
// index:0
// Public
// void setWindowFlags(Qt::WindowFlags)
func (this *QWidget) SetWindowFlags(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14setWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:568
// index:0
// Public inline
// Qt::WindowFlags windowFlags()
func (this *QWidget) WindowFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11windowFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:569
// index:0
// Public
// void setWindowFlag(Qt::WindowType, bool)
func (this *QWidget) SetWindowFlag(arg0 int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13setWindowFlagEN2Qt10WindowTypeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:570
// index:0
// Public
// void overrideWindowFlags(Qt::WindowFlags)
func (this *QWidget) OverrideWindowFlags(type_ int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget19overrideWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:572
// index:0
// Public inline
// Qt::WindowType windowType()
func (this *QWidget) WindowType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10windowTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:574
// index:0
// Public static
// QWidget * find(WId)
func (this *QWidget) Find(arg0 uint64) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget4findEy", ffiqt.FFI_TYPE_POINTER, arg0)
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
// Public inline
// QWidget * childAt(int, int)
func (this *QWidget) ChildAt(x int, y int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7childAtEii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:576
// index:1
// Public
// QWidget * childAt(const QPoint &)
func (this *QWidget) ChildAt_1(p *qtcore.QPoint) *QWidget /*777 QWidget **/ {
	var convArg0 = p.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget7childAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:578
// index:0
// Public
// void setAttribute(Qt::WidgetAttribute, bool)
func (this *QWidget) SetAttribute(arg0 int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12setAttributeEN2Qt15WidgetAttributeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:579
// index:0
// Public inline
// bool testAttribute(Qt::WidgetAttribute)
func (this *QWidget) TestAttribute(arg0 int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13testAttributeEN2Qt15WidgetAttributeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:581
// index:0
// Public virtual
// QPaintEngine * paintEngine()
func (this *QWidget) PaintEngine() *qtgui.QPaintEngine /*777 QPaintEngine **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11paintEngineEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaintEngineFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:583
// index:0
// Public
// void ensurePolished()
func (this *QWidget) EnsurePolished() {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget14ensurePolishedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:585
// index:0
// Public
// bool isAncestorOf(const QWidget *)
func (this *QWidget) IsAncestorOf(child *QWidget /*777 const QWidget **/) bool {
	var convArg0 = child.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12isAncestorOfEPKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:592
// index:0
// Public
// bool autoFillBackground()
func (this *QWidget) AutoFillBackground() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget18autoFillBackgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:593
// index:0
// Public
// void setAutoFillBackground(bool)
func (this *QWidget) SetAutoFillBackground(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21setAutoFillBackgroundEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:595
// index:0
// Public
// QBackingStore * backingStore()
func (this *QWidget) BackingStore() *qtgui.QBackingStore /*777 QBackingStore **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12backingStoreEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQBackingStoreFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:597
// index:0
// Public
// QWindow * windowHandle()
func (this *QWidget) WindowHandle() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget12windowHandleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:599
// index:0
// Public static
// QWidget * createWindowContainer(QWindow *, QWidget *, Qt::WindowFlags)
func (this *QWidget) CreateWindowContainer(window *qtgui.QWindow /*777 QWindow **/, parent *QWidget /*777 QWidget **/, flags int) *QWidget /*777 QWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21createWindowContainerEP7QWindowPS_6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, window, parent, flags)
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
// Public
// void windowTitleChanged(const QString &)
func (this *QWidget) WindowTitleChanged(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18windowTitleChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:605
// index:0
// Public
// void windowIconChanged(const QIcon &)
func (this *QWidget) WindowIconChanged(icon *qtgui.QIcon) {
	var convArg0 = icon.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17windowIconChangedERK5QIcon", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:606
// index:0
// Public
// void windowIconTextChanged(const QString &)
func (this *QWidget) WindowIconTextChanged(iconText *qtcore.QString) {
	var convArg0 = iconText.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21windowIconTextChangedERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:607
// index:0
// Public
// void customContextMenuRequested(const QPoint &)
func (this *QWidget) CustomContextMenuRequested(pos *qtcore.QPoint) {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget26customContextMenuRequestedERK6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:611
// index:0
// Protected virtual
// bool event(QEvent *)
func (this *QWidget) Event(event *qtcore.QEvent /*777 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:612
// index:0
// Protected virtual
// void mousePressEvent(QMouseEvent *)
func (this *QWidget) MousePressEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15mousePressEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:613
// index:0
// Protected virtual
// void mouseReleaseEvent(QMouseEvent *)
func (this *QWidget) MouseReleaseEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget17mouseReleaseEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:614
// index:0
// Protected virtual
// void mouseDoubleClickEvent(QMouseEvent *)
func (this *QWidget) MouseDoubleClickEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget21mouseDoubleClickEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:615
// index:0
// Protected virtual
// void mouseMoveEvent(QMouseEvent *)
func (this *QWidget) MouseMoveEvent(event *qtgui.QMouseEvent /*777 QMouseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14mouseMoveEventEP11QMouseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:617
// index:0
// Protected virtual
// void wheelEvent(QWheelEvent *)
func (this *QWidget) WheelEvent(event *qtgui.QWheelEvent /*777 QWheelEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10wheelEventEP11QWheelEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:619
// index:0
// Protected virtual
// void keyPressEvent(QKeyEvent *)
func (this *QWidget) KeyPressEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13keyPressEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:620
// index:0
// Protected virtual
// void keyReleaseEvent(QKeyEvent *)
func (this *QWidget) KeyReleaseEvent(event *qtgui.QKeyEvent /*777 QKeyEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget15keyReleaseEventEP9QKeyEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:621
// index:0
// Protected virtual
// void focusInEvent(QFocusEvent *)
func (this *QWidget) FocusInEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:622
// index:0
// Protected virtual
// void focusOutEvent(QFocusEvent *)
func (this *QWidget) FocusOutEvent(event *qtgui.QFocusEvent /*777 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:623
// index:0
// Protected virtual
// void enterEvent(QEvent *)
func (this *QWidget) EnterEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10enterEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:624
// index:0
// Protected virtual
// void leaveEvent(QEvent *)
func (this *QWidget) LeaveEvent(event *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10leaveEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:625
// index:0
// Protected virtual
// void paintEvent(QPaintEvent *)
func (this *QWidget) PaintEvent(event *qtgui.QPaintEvent /*777 QPaintEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10paintEventEP11QPaintEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:626
// index:0
// Protected virtual
// void moveEvent(QMoveEvent *)
func (this *QWidget) MoveEvent(event *qtgui.QMoveEvent /*777 QMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9moveEventEP10QMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:627
// index:0
// Protected virtual
// void resizeEvent(QResizeEvent *)
func (this *QWidget) ResizeEvent(event *qtgui.QResizeEvent /*777 QResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11resizeEventEP12QResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:628
// index:0
// Protected virtual
// void closeEvent(QCloseEvent *)
func (this *QWidget) CloseEvent(event *qtgui.QCloseEvent /*777 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:630
// index:0
// Protected virtual
// void contextMenuEvent(QContextMenuEvent *)
func (this *QWidget) ContextMenuEvent(event *qtgui.QContextMenuEvent /*777 QContextMenuEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16contextMenuEventEP17QContextMenuEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:633
// index:0
// Protected virtual
// void tabletEvent(QTabletEvent *)
func (this *QWidget) TabletEvent(event *qtgui.QTabletEvent /*777 QTabletEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11tabletEventEP12QTabletEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:636
// index:0
// Protected virtual
// void actionEvent(QActionEvent *)
func (this *QWidget) ActionEvent(event *qtgui.QActionEvent /*777 QActionEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11actionEventEP12QActionEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:640
// index:0
// Protected virtual
// void dragEnterEvent(QDragEnterEvent *)
func (this *QWidget) DragEnterEvent(event *qtgui.QDragEnterEvent /*777 QDragEnterEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14dragEnterEventEP15QDragEnterEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:641
// index:0
// Protected virtual
// void dragMoveEvent(QDragMoveEvent *)
func (this *QWidget) DragMoveEvent(event *qtgui.QDragMoveEvent /*777 QDragMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget13dragMoveEventEP14QDragMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:642
// index:0
// Protected virtual
// void dragLeaveEvent(QDragLeaveEvent *)
func (this *QWidget) DragLeaveEvent(event *qtgui.QDragLeaveEvent /*777 QDragLeaveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14dragLeaveEventEP15QDragLeaveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:643
// index:0
// Protected virtual
// void dropEvent(QDropEvent *)
func (this *QWidget) DropEvent(event *qtgui.QDropEvent /*777 QDropEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9dropEventEP10QDropEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:646
// index:0
// Protected virtual
// void showEvent(QShowEvent *)
func (this *QWidget) ShowEvent(event *qtgui.QShowEvent /*777 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:647
// index:0
// Protected virtual
// void hideEvent(QHideEvent *)
func (this *QWidget) HideEvent(event *qtgui.QHideEvent /*777 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:648
// index:0
// Protected virtual
// bool nativeEvent(const QByteArray &, void *, long *)
func (this *QWidget) NativeEvent(eventType *qtcore.QByteArray, message unsafe.Pointer /*666*/, result unsafe.Pointer /*666*/) bool {
	var convArg0 = eventType.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11nativeEventERK10QByteArrayPvPl", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, message, &result)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:651
// index:0
// Protected virtual
// void changeEvent(QEvent *)
func (this *QWidget) ChangeEvent(arg0 *qtcore.QEvent /*777 QEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:653
// index:0
// Protected virtual
// int metric(QPaintDevice::PaintDeviceMetric)
func (this *QWidget) Metric(arg0 int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget6metricEN12QPaintDevice17PaintDeviceMetricE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qwidget.h:654
// index:0
// Protected virtual
// void initPainter(QPainter *)
func (this *QWidget) InitPainter(painter *qtgui.QPainter /*777 QPainter **/) {
	var convArg0 = painter.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget11initPainterEP8QPainter", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:655
// index:0
// Protected virtual
// QPaintDevice * redirected(QPoint *)
func (this *QWidget) Redirected(offset *qtcore.QPoint /*777 QPoint **/) *qtgui.QPaintDevice /*777 QPaintDevice **/ {
	var convArg0 = offset.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget10redirectedEP6QPoint", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPaintDeviceFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:656
// index:0
// Protected virtual
// QPainter * sharedPainter()
func (this *QWidget) SharedPainter() *qtgui.QPainter /*777 QPainter **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget13sharedPainterEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPainterFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:658
// index:0
// Protected virtual
// void inputMethodEvent(QInputMethodEvent *)
func (this *QWidget) InputMethodEvent(arg0 *qtgui.QInputMethodEvent /*777 QInputMethodEvent **/) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16inputMethodEventEP17QInputMethodEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:660
// index:0
// Public virtual
// QVariant inputMethodQuery(Qt::InputMethodQuery)
func (this *QWidget) InputMethodQuery(arg0 int) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16inputMethodQueryEN2Qt16InputMethodQueryE", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qwidget.h:662
// index:0
// Public
// Qt::InputMethodHints inputMethodHints()
func (this *QWidget) InputMethodHints() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK7QWidget16inputMethodHintsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:666
// index:0
// Protected
// void updateMicroFocus()
func (this *QWidget) UpdateMicroFocus() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget16updateMicroFocusEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:669
// index:0
// Protected
// void create(WId, bool, bool)
func (this *QWidget) Create(arg0 uint64, initializeWindow bool, destroyOldWindow bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget6createEybb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, initializeWindow, destroyOldWindow)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:671
// index:0
// Protected
// void destroy(bool, bool)
func (this *QWidget) Destroy(destroyWindow bool, destroySubWindows bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget7destroyEbb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), destroyWindow, destroySubWindows)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qwidget.h:675
// index:0
// Protected virtual
// bool focusNextPrevChild(bool)
func (this *QWidget) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:676
// index:0
// Protected inline
// bool focusNextChild()
func (this *QWidget) FocusNextChild() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget14focusNextChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qwidget.h:677
// index:0
// Protected inline
// bool focusPreviousChild()
func (this *QWidget) FocusPreviousChild() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN7QWidget18focusPreviousChildEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

type QWidget__RenderFlag = int

const QWidget__DrawWindowBackground QWidget__RenderFlag = 1
const QWidget__DrawChildren QWidget__RenderFlag = 2
const QWidget__IgnoreMask QWidget__RenderFlag = 4

//  body block end
