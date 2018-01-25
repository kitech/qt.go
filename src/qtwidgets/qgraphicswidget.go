package qtwidgets

// /usr/include/qt/QtWidgets/qgraphicswidget.h
// #include <qgraphicswidget.h>
// #include <QtWidgets>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 24
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"
import "qtgui"

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
type QGraphicsWidget struct {
	*QGraphicsLayoutItem
}

func (this *QGraphicsWidget) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QGraphicsLayoutItem.GetCthis()
	}
}
func (this *QGraphicsWidget) SetCthis(cthis unsafe.Pointer) {
	this.QGraphicsLayoutItem = NewQGraphicsLayoutItemFromPointer(cthis)
}
func NewQGraphicsWidgetFromPointer(cthis unsafe.Pointer) *QGraphicsWidget {
	bcthis0 := NewQGraphicsLayoutItemFromPointer(cthis)
	return &QGraphicsWidget{bcthis0}
}
func (*QGraphicsWidget) NewFromPointer(cthis unsafe.Pointer) *QGraphicsWidget {
	return NewQGraphicsWidgetFromPointer(cthis)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:66
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QGraphicsWidget) MetaObject() *qtcore.QMetaObject /*444 const QMetaObject **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:83
// index:0
// Public
// void QGraphicsWidget(class QGraphicsItem *, Qt::WindowFlags)
func NewQGraphicsWidget(parent *QGraphicsItem /*444 QGraphicsItem **/, wFlags int) *QGraphicsWidget {
	cthis := qtrt.Calloc(1, 256) // 48
	var convArg0 = parent.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidgetC2EP13QGraphicsItem6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_VOID, cthis, convArg0, wFlags)
	gopp.ErrPrint(err, rv)
	gothis := NewQGraphicsWidgetFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:84
// index:0
// Public virtual
// void ~QGraphicsWidget()
func DeleteQGraphicsWidget(*QGraphicsWidget) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidgetD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:85
// index:0
// Public
// QGraphicsLayout * layout()
func (this *QGraphicsWidget) Layout() *QGraphicsLayout /*444 QGraphicsLayout **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget6layoutEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsLayoutFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:86
// index:0
// Public
// void setLayout(class QGraphicsLayout *)
func (this *QGraphicsWidget) SetLayout(layout *QGraphicsLayout /*444 QGraphicsLayout **/) {
	var convArg0 = layout.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9setLayoutEP15QGraphicsLayout", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:87
// index:0
// Public
// void adjustSize()
func (this *QGraphicsWidget) AdjustSize() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10adjustSizeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:89
// index:0
// Public
// Qt::LayoutDirection layoutDirection()
func (this *QGraphicsWidget) LayoutDirection() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget15layoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:90
// index:0
// Public
// void setLayoutDirection(Qt::LayoutDirection)
func (this *QGraphicsWidget) SetLayoutDirection(direction int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18setLayoutDirectionEN2Qt15LayoutDirectionE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), direction)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:91
// index:0
// Public
// void unsetLayoutDirection()
func (this *QGraphicsWidget) UnsetLayoutDirection() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget20unsetLayoutDirectionEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:93
// index:0
// Public
// QStyle * style()
func (this *QGraphicsWidget) Style() *QStyle /*444 QStyle **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget5styleEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStyleFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:94
// index:0
// Public
// void setStyle(class QStyle *)
func (this *QGraphicsWidget) SetStyle(style *QStyle /*444 QStyle **/) {
	var convArg0 = style.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget8setStyleEP6QStyle", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:96
// index:0
// Public
// QFont font()
func (this *QGraphicsWidget) Font() *qtgui.QFont /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4fontEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQFontFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:97
// index:0
// Public
// void setFont(const class QFont &)
func (this *QGraphicsWidget) SetFont(font *qtgui.QFont) {
	var convArg0 = font.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget7setFontERK5QFont", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:99
// index:0
// Public
// QPalette palette()
func (this *QGraphicsWidget) Palette() *qtgui.QPalette /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget7paletteEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPaletteFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:100
// index:0
// Public
// void setPalette(const class QPalette &)
func (this *QGraphicsWidget) SetPalette(palette *qtgui.QPalette) {
	var convArg0 = palette.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10setPaletteERK8QPalette", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:102
// index:0
// Public
// bool autoFillBackground()
func (this *QGraphicsWidget) AutoFillBackground() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget18autoFillBackgroundEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:103
// index:0
// Public
// void setAutoFillBackground(_Bool)
func (this *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget21setAutoFillBackgroundEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:105
// index:0
// Public
// void resize(const class QSizeF &)
func (this *QGraphicsWidget) Resize(size *qtcore.QSizeF) {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeERK6QSizeF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:106
// index:1
// Public inline
// void resize(qreal, qreal)
func (this *QGraphicsWidget) Resize_1(w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget6resizeEdd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:107
// index:0
// Public
// QSizeF size()
func (this *QGraphicsWidget) Size() *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4sizeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:109
// index:0
// Public virtual
// void setGeometry(const class QRectF &)
func (this *QGraphicsWidget) SetGeometry(rect *qtcore.QRectF) {
	var convArg0 = rect.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryERK6QRectF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:110
// index:1
// Public inline
// void setGeometry(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetGeometry_1(x float64, y float64, w float64, h float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11setGeometryEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), x, y, w, h)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:111
// index:0
// Public inline
// QRectF rect()
func (this *QGraphicsWidget) Rect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4rectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:113
// index:0
// Public
// void setContentsMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18setContentsMarginsEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:114
// index:0
// Public virtual
// void getContentsMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) GetContentsMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget18getContentsMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:116
// index:0
// Public
// void setWindowFrameMargins(qreal, qreal, qreal, qreal)
func (this *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget21setWindowFrameMarginsEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), left, top, right, bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:117
// index:0
// Public
// void getWindowFrameMargins(qreal *, qreal *, qreal *, qreal *)
func (this *QGraphicsWidget) GetWindowFrameMargins(left unsafe.Pointer /*666*/, top unsafe.Pointer /*666*/, right unsafe.Pointer /*666*/, bottom unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget21getWindowFrameMarginsEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &left, &top, &right, &bottom)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:118
// index:0
// Public
// void unsetWindowFrameMargins()
func (this *QGraphicsWidget) UnsetWindowFrameMargins() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget23unsetWindowFrameMarginsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:119
// index:0
// Public
// QRectF windowFrameGeometry()
func (this *QGraphicsWidget) WindowFrameGeometry() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget19windowFrameGeometryEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:120
// index:0
// Public
// QRectF windowFrameRect()
func (this *QGraphicsWidget) WindowFrameRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget15windowFrameRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:123
// index:0
// Public
// Qt::WindowFlags windowFlags()
func (this *QGraphicsWidget) WindowFlags() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowFlagsEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:124
// index:0
// Public
// Qt::WindowType windowType()
func (this *QGraphicsWidget) WindowType() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget10windowTypeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:125
// index:0
// Public
// void setWindowFlags(Qt::WindowFlags)
func (this *QGraphicsWidget) SetWindowFlags(wFlags int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowFlagsE6QFlagsIN2Qt10WindowTypeEE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), wFlags)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:126
// index:0
// Public
// bool isActiveWindow()
func (this *QGraphicsWidget) IsActiveWindow() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget14isActiveWindowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:127
// index:0
// Public
// void setWindowTitle(const class QString &)
func (this *QGraphicsWidget) SetWindowTitle(title *qtcore.QString) {
	var convArg0 = title.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14setWindowTitleERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:128
// index:0
// Public
// QString windowTitle()
func (this *QGraphicsWidget) WindowTitle() *qtcore.QString /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11windowTitleEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:131
// index:0
// Public
// Qt::FocusPolicy focusPolicy()
func (this *QGraphicsWidget) FocusPolicy() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusPolicyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:132
// index:0
// Public
// void setFocusPolicy(Qt::FocusPolicy)
func (this *QGraphicsWidget) SetFocusPolicy(policy int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14setFocusPolicyEN2Qt11FocusPolicyE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), policy)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:133
// index:0
// Public static
// void setTabOrder(class QGraphicsWidget *, class QGraphicsWidget *)
func (this *QGraphicsWidget) SetTabOrder(first *QGraphicsWidget /*444 QGraphicsWidget **/, second *QGraphicsWidget /*444 QGraphicsWidget **/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11setTabOrderEPS_S0_", ffiqt.FFI_TYPE_POINTER, first, second)
	gopp.ErrPrint(err, rv)
}
func QGraphicsWidget_SetTabOrder(first *QGraphicsWidget /*444 QGraphicsWidget **/, second *QGraphicsWidget /*444 QGraphicsWidget **/) {
	var nilthis *QGraphicsWidget
	nilthis.SetTabOrder(first, second)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:134
// index:0
// Public
// QGraphicsWidget * focusWidget()
func (this *QGraphicsWidget) FocusWidget() *QGraphicsWidget /*444 QGraphicsWidget **/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget11focusWidgetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQGraphicsWidgetFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:137
// index:0
// Public
// int grabShortcut(const class QKeySequence &, Qt::ShortcutContext)
func (this *QGraphicsWidget) GrabShortcut(sequence *qtgui.QKeySequence, context int) int {
	var convArg0 = sequence.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12grabShortcutERK12QKeySequenceN2Qt15ShortcutContextE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, context)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:138
// index:0
// Public
// void releaseShortcut(int)
func (this *QGraphicsWidget) ReleaseShortcut(id int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget15releaseShortcutEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:139
// index:0
// Public
// void setShortcutEnabled(int, _Bool)
func (this *QGraphicsWidget) SetShortcutEnabled(id int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18setShortcutEnabledEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:140
// index:0
// Public
// void setShortcutAutoRepeat(int, _Bool)
func (this *QGraphicsWidget) SetShortcutAutoRepeat(id int, enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget21setShortcutAutoRepeatEib", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), id, enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:145
// index:0
// Public
// void addAction(class QAction *)
func (this *QGraphicsWidget) AddAction(action *QAction /*444 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9addActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:153
// index:0
// Public
// void insertAction(class QAction *, class QAction *)
func (this *QGraphicsWidget) InsertAction(before *QAction /*444 QAction **/, action *QAction /*444 QAction **/) {
	var convArg0 = before.GetCthis()
	var convArg1 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12insertActionEP7QActionS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:154
// index:0
// Public
// void removeAction(class QAction *)
func (this *QGraphicsWidget) RemoveAction(action *QAction /*444 QAction **/) {
	var convArg0 = action.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12removeActionEP7QAction", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:158
// index:0
// Public
// void setAttribute(Qt::WidgetAttribute, _Bool)
func (this *QGraphicsWidget) SetAttribute(attribute int, on bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12setAttributeEN2Qt15WidgetAttributeEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), attribute, on)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:159
// index:0
// Public
// bool testAttribute(Qt::WidgetAttribute)
func (this *QGraphicsWidget) TestAttribute(attribute int) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget13testAttributeEN2Qt15WidgetAttributeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), attribute)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:164
// index:0
// Public virtual
// int type()
func (this *QGraphicsWidget) Type() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget4typeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv) // 111
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:166
// index:0
// Public virtual
// void paint(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) Paint(painter *qtgui.QPainter /*444 QPainter **/, option *QStyleOptionGraphicsItem /*444 const QStyleOptionGraphicsItem **/, widget *QWidget /*444 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget5paintEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:167
// index:0
// Public virtual
// void paintWindowFrame(class QPainter *, const class QStyleOptionGraphicsItem *, class QWidget *)
func (this *QGraphicsWidget) PaintWindowFrame(painter *qtgui.QPainter /*444 QPainter **/, option *QStyleOptionGraphicsItem /*444 const QStyleOptionGraphicsItem **/, widget *QWidget /*444 QWidget **/) {
	var convArg0 = painter.GetCthis()
	var convArg1 = option.GetCthis()
	var convArg2 = widget.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget16paintWindowFrameEP8QPainterPK24QStyleOptionGraphicsItemP7QWidget", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:168
// index:0
// Public virtual
// QRectF boundingRect()
func (this *QGraphicsWidget) BoundingRect() *qtcore.QRectF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget12boundingRectEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQRectFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:169
// index:0
// Public virtual
// QPainterPath shape()
func (this *QGraphicsWidget) Shape() *qtgui.QPainterPath /*123*/ {
	mv := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget5shapeEv", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtgui.NewQPainterPathFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:178
// index:0
// Public
// void geometryChanged()
func (this *QGraphicsWidget) GeometryChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget15geometryChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:179
// index:0
// Public
// void layoutChanged()
func (this *QGraphicsWidget) LayoutChanged() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget13layoutChangedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:182
// index:0
// Public
// bool close()
func (this *QGraphicsWidget) Close() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget5closeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:185
// index:0
// Protected virtual
// void initStyleOption(class QStyleOption *)
func (this *QGraphicsWidget) InitStyleOption(option *QStyleOption /*444 QStyleOption **/) {
	var convArg0 = option.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget15initStyleOptionEP12QStyleOption", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:187
// index:0
// Protected virtual
// QSizeF sizeHint(Qt::SizeHint, const class QSizeF &)
func (this *QGraphicsWidget) SizeHint(which int, constraint *qtcore.QSizeF) *qtcore.QSizeF /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = constraint.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget8sizeHintEN2Qt8SizeHintERK6QSizeF", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), which, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQSizeFFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:188
// index:0
// Protected virtual
// void updateGeometry()
func (this *QGraphicsWidget) UpdateGeometry() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14updateGeometryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:191
// index:0
// Protected virtual
// QVariant itemChange(enum QGraphicsItem::GraphicsItemChange, const class QVariant &)
func (this *QGraphicsWidget) ItemChange(change int, value *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10itemChangeEN13QGraphicsItem18GraphicsItemChangeERK8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), change, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:192
// index:0
// Protected virtual
// QVariant propertyChange(const class QString &, const class QVariant &)
func (this *QGraphicsWidget) PropertyChange(propertyName *qtcore.QString, value *qtcore.QVariant) *qtcore.QVariant /*123*/ {
	mv := qtrt.Calloc(1, 256)
	var convArg0 = propertyName.GetCthis()
	var convArg1 = value.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14propertyChangeERK7QStringRK8QVariant", ffiqt.FFI_TYPE_POINTER, mv, this.GetCthis(), convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv = uint64(uintptr(mv))
	rv2 := qtcore.NewQVariantFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	return rv2
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:195
// index:0
// Protected virtual
// bool sceneEvent(class QEvent *)
func (this *QGraphicsWidget) SceneEvent(event *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10sceneEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:196
// index:0
// Protected virtual
// bool windowFrameEvent(class QEvent *)
func (this *QGraphicsWidget) WindowFrameEvent(e *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = e.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget16windowFrameEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:197
// index:0
// Protected virtual
// Qt::WindowFrameSection windowFrameSectionAt(const class QPointF &)
func (this *QGraphicsWidget) WindowFrameSectionAt(pos *qtcore.QPointF) int {
	var convArg0 = pos.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK15QGraphicsWidget20windowFrameSectionAtERK7QPointF", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:200
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QGraphicsWidget) Event(event *qtcore.QEvent /*444 QEvent **/) bool {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:202
// index:0
// Protected virtual
// void changeEvent(class QEvent *)
func (this *QGraphicsWidget) ChangeEvent(event *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11changeEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:203
// index:0
// Protected virtual
// void closeEvent(class QCloseEvent *)
func (this *QGraphicsWidget) CloseEvent(event *qtgui.QCloseEvent /*444 QCloseEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget10closeEventEP11QCloseEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:206
// index:0
// Protected virtual
// void focusInEvent(class QFocusEvent *)
func (this *QGraphicsWidget) FocusInEvent(event *qtgui.QFocusEvent /*444 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget12focusInEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:207
// index:0
// Protected virtual
// bool focusNextPrevChild(_Bool)
func (this *QGraphicsWidget) FocusNextPrevChild(next bool) bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget18focusNextPrevChildEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), next)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:208
// index:0
// Protected virtual
// void focusOutEvent(class QFocusEvent *)
func (this *QGraphicsWidget) FocusOutEvent(event *qtgui.QFocusEvent /*444 QFocusEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget13focusOutEventEP11QFocusEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:209
// index:0
// Protected virtual
// void hideEvent(class QHideEvent *)
func (this *QGraphicsWidget) HideEvent(event *qtgui.QHideEvent /*444 QHideEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9hideEventEP10QHideEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:211
// index:0
// Protected virtual
// void moveEvent(class QGraphicsSceneMoveEvent *)
func (this *QGraphicsWidget) MoveEvent(event *QGraphicsSceneMoveEvent /*444 QGraphicsSceneMoveEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9moveEventEP23QGraphicsSceneMoveEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:212
// index:0
// Protected virtual
// void polishEvent()
func (this *QGraphicsWidget) PolishEvent() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11polishEventEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:214
// index:0
// Protected virtual
// void resizeEvent(class QGraphicsSceneResizeEvent *)
func (this *QGraphicsWidget) ResizeEvent(event *QGraphicsSceneResizeEvent /*444 QGraphicsSceneResizeEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget11resizeEventEP25QGraphicsSceneResizeEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:215
// index:0
// Protected virtual
// void showEvent(class QShowEvent *)
func (this *QGraphicsWidget) ShowEvent(event *qtgui.QShowEvent /*444 QShowEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget9showEventEP10QShowEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:217
// index:0
// Protected virtual
// void hoverMoveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) HoverMoveEvent(event *QGraphicsSceneHoverEvent /*444 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14hoverMoveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:218
// index:0
// Protected virtual
// void hoverLeaveEvent(class QGraphicsSceneHoverEvent *)
func (this *QGraphicsWidget) HoverLeaveEvent(event *QGraphicsSceneHoverEvent /*444 QGraphicsSceneHoverEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget15hoverLeaveEventEP24QGraphicsSceneHoverEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:219
// index:0
// Protected virtual
// void grabMouseEvent(class QEvent *)
func (this *QGraphicsWidget) GrabMouseEvent(event *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget14grabMouseEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:220
// index:0
// Protected virtual
// void ungrabMouseEvent(class QEvent *)
func (this *QGraphicsWidget) UngrabMouseEvent(event *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget16ungrabMouseEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:221
// index:0
// Protected virtual
// void grabKeyboardEvent(class QEvent *)
func (this *QGraphicsWidget) GrabKeyboardEvent(event *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget17grabKeyboardEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qgraphicswidget.h:222
// index:0
// Protected virtual
// void ungrabKeyboardEvent(class QEvent *)
func (this *QGraphicsWidget) UngrabKeyboardEvent(event *qtcore.QEvent /*444 QEvent **/) {
	var convArg0 = event.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN15QGraphicsWidget19ungrabKeyboardEventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

type QGraphicsWidget__ = int

const QGraphicsWidget__Type QGraphicsWidget__ = 11

//  body block end
