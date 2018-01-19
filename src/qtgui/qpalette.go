//  header block begin
// /usr/include/qt/QtGui/qpalette.h
// #include <qpalette.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 89
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
type QPalette struct {
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qpalette.h:58
// index:0
// void QPalette()
func NewQPalette() *QPalette {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QPalette{cthis}
}

// /usr/include/qt/QtGui/qpalette.h:59
// index:1
// void QPalette(const class QColor &)
func NewQPalette_1(button unsafe.Pointer) *QPalette {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColor", ffiqt.FFI_TYPE_VOID, cthis, button)
	gopp.ErrPrint(err, rv)
	return &QPalette{cthis}
}

// /usr/include/qt/QtGui/qpalette.h:60
// index:2
// void QPalette(Qt::GlobalColor)
func NewQPalette_2(button int) *QPalette {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2EN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, cthis, &button)
	gopp.ErrPrint(err, rv)
	return &QPalette{cthis}
}

// /usr/include/qt/QtGui/qpalette.h:61
// index:3
// void QPalette(const class QColor &, const class QColor &)
func NewQPalette_3(button unsafe.Pointer, window unsafe.Pointer) *QPalette {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColorS2_", ffiqt.FFI_TYPE_VOID, cthis, button, window)
	gopp.ErrPrint(err, rv)
	return &QPalette{cthis}
}

// /usr/include/qt/QtGui/qpalette.h:62
// index:4
// void QPalette(const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &)
func NewQPalette_4(windowText unsafe.Pointer, button unsafe.Pointer, light unsafe.Pointer, dark unsafe.Pointer, mid unsafe.Pointer, text unsafe.Pointer, bright_text unsafe.Pointer, base unsafe.Pointer, window unsafe.Pointer) *QPalette {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QBrushS2_S2_S2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_VOID, cthis, windowText, button, light, dark, mid, text, bright_text, base, window)
	gopp.ErrPrint(err, rv)
	return &QPalette{cthis}
}

// /usr/include/qt/QtGui/qpalette.h:65
// index:5
// void QPalette(const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &, const class QColor &)
func NewQPalette_5(windowText unsafe.Pointer, window unsafe.Pointer, light unsafe.Pointer, dark unsafe.Pointer, mid unsafe.Pointer, text unsafe.Pointer, base unsafe.Pointer) *QPalette {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteC2ERK6QColorS2_S2_S2_S2_S2_S2_", ffiqt.FFI_TYPE_VOID, cthis, windowText, window, light, dark, mid, text, base)
	gopp.ErrPrint(err, rv)
	return &QPalette{cthis}
}

// /usr/include/qt/QtGui/qpalette.h:68
// index:0
// void ~QPalette()
func DeleteQPalette(*QPalette) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPaletteD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:81
// index:0
// inline
// void swap(class QPalette &)
func (this *QPalette) Swap(other unsafe.Pointer) {
	// 0: (, QPalette & other), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette4swapERS_", ffiqt.FFI_TYPE_VOID, this.cthis, other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:104
// index:0
// inline
// QPalette::ColorGroup currentColorGroup()
func (this *QPalette) CurrentColorGroup() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette17currentColorGroupEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:105
// index:0
// inline
// void setCurrentColorGroup(enum QPalette::ColorGroup)
func (this *QPalette) SetCurrentColorGroup(cg int) {
	// 0: (, QPalette::ColorGroup cg), (&cg)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette20setCurrentColorGroupENS_10ColorGroupE", ffiqt.FFI_TYPE_VOID, this.cthis, &cg)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:107
// index:0
// inline
// const QColor & color(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) Color(cg int, cr int) {
	// 0: (, QPalette::ColorGroup cg, QPalette::ColorRole cr), (&cg, &cr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5colorENS_10ColorGroupENS_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &cg, &cr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:121
// index:1
// inline
// const QColor & color(enum QPalette::ColorRole)
func (this *QPalette) Color_1(cr int) {
	// 1: (, QPalette::ColorRole cr), (&cr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5colorENS_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &cr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:109
// index:0
// const QBrush & brush(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) Brush(cg int, cr int) {
	// 0: (, QPalette::ColorGroup cg, QPalette::ColorRole cr), (&cg, &cr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5brushENS_10ColorGroupENS_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &cg, &cr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:122
// index:1
// inline
// const QBrush & brush(enum QPalette::ColorRole)
func (this *QPalette) Brush_1(cr int) {
	// 1: (, QPalette::ColorRole cr), (&cr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5brushENS_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &cr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:110
// index:0
// inline
// void setColor(enum QPalette::ColorGroup, enum QPalette::ColorRole, const class QColor &)
func (this *QPalette) SetColor(cg int, cr int, color unsafe.Pointer) {
	// 0: (, QPalette::ColorGroup cg, QPalette::ColorRole cr, const QColor & color), (&cg, &cr, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setColorENS_10ColorGroupENS_9ColorRoleERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, &cg, &cr, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:111
// index:1
// inline
// void setColor(enum QPalette::ColorRole, const class QColor &)
func (this *QPalette) SetColor_1(cr int, color unsafe.Pointer) {
	// 1: (, QPalette::ColorRole cr, const QColor & color), (&cr, color)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setColorENS_9ColorRoleERK6QColor", ffiqt.FFI_TYPE_VOID, this.cthis, &cr, color)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:112
// index:0
// inline
// void setBrush(enum QPalette::ColorRole, const class QBrush &)
func (this *QPalette) SetBrush(cr int, brush unsafe.Pointer) {
	// 0: (, QPalette::ColorRole cr, const QBrush & brush), (&cr, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setBrushENS_9ColorRoleERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, &cr, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:114
// index:1
// void setBrush(enum QPalette::ColorGroup, enum QPalette::ColorRole, const class QBrush &)
func (this *QPalette) SetBrush_1(cg int, cr int, brush unsafe.Pointer) {
	// 1: (, QPalette::ColorGroup cg, QPalette::ColorRole cr, const QBrush & brush), (&cg, &cr, brush)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette8setBrushENS_10ColorGroupENS_9ColorRoleERK6QBrush", ffiqt.FFI_TYPE_VOID, this.cthis, &cg, &cr, brush)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:113
// index:0
// bool isBrushSet(enum QPalette::ColorGroup, enum QPalette::ColorRole)
func (this *QPalette) IsBrushSet(cg int, cr int) {
	// 0: (, QPalette::ColorGroup cg, QPalette::ColorRole cr), (&cg, &cr)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10isBrushSetENS_10ColorGroupENS_9ColorRoleE", ffiqt.FFI_TYPE_VOID, this.cthis, &cg, &cr)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:115
// index:0
// void setColorGroup(enum QPalette::ColorGroup, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &, const class QBrush &)
func (this *QPalette) SetColorGroup(cr int, windowText unsafe.Pointer, button unsafe.Pointer, light unsafe.Pointer, dark unsafe.Pointer, mid unsafe.Pointer, text unsafe.Pointer, bright_text unsafe.Pointer, base unsafe.Pointer, window unsafe.Pointer) {
	// 0: (, QPalette::ColorGroup cr, const QBrush & windowText, const QBrush & button, const QBrush & light, const QBrush & dark, const QBrush & mid, const QBrush & text, const QBrush & bright_text, const QBrush & base, const QBrush & window), (&cr, windowText, button, light, dark, mid, text, bright_text, base, window)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette13setColorGroupENS_10ColorGroupERK6QBrushS3_S3_S3_S3_S3_S3_S3_S3_", ffiqt.FFI_TYPE_VOID, this.cthis, &cr, windowText, button, light, dark, mid, text, bright_text, base, window)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:119
// index:0
// bool isEqual(enum QPalette::ColorGroup, enum QPalette::ColorGroup)
func (this *QPalette) IsEqual(cr1 int, cr2 int) {
	// 0: (, QPalette::ColorGroup cr1, QPalette::ColorGroup cr2), (&cr1, &cr2)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette7isEqualENS_10ColorGroupES0_", ffiqt.FFI_TYPE_VOID, this.cthis, &cr1, &cr2)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:123
// index:0
// inline
// const QBrush & foreground()
func (this *QPalette) Foreground() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10foregroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:124
// index:0
// inline
// const QBrush & windowText()
func (this *QPalette) WindowText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10windowTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:125
// index:0
// inline
// const QBrush & button()
func (this *QPalette) Button() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette6buttonEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:126
// index:0
// inline
// const QBrush & light()
func (this *QPalette) Light() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette5lightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:127
// index:0
// inline
// const QBrush & dark()
func (this *QPalette) Dark() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4darkEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:128
// index:0
// inline
// const QBrush & mid()
func (this *QPalette) Mid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette3midEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:129
// index:0
// inline
// const QBrush & text()
func (this *QPalette) Text() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4textEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:130
// index:0
// inline
// const QBrush & base()
func (this *QPalette) Base() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4baseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:131
// index:0
// inline
// const QBrush & alternateBase()
func (this *QPalette) AlternateBase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette13alternateBaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:132
// index:0
// inline
// const QBrush & toolTipBase()
func (this *QPalette) ToolTipBase() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette11toolTipBaseEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:133
// index:0
// inline
// const QBrush & toolTipText()
func (this *QPalette) ToolTipText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette11toolTipTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:134
// index:0
// inline
// const QBrush & background()
func (this *QPalette) Background() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10backgroundEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:135
// index:0
// inline
// const QBrush & window()
func (this *QPalette) Window() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette6windowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:136
// index:0
// inline
// const QBrush & midlight()
func (this *QPalette) Midlight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette8midlightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:137
// index:0
// inline
// const QBrush & brightText()
func (this *QPalette) BrightText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10brightTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:138
// index:0
// inline
// const QBrush & buttonText()
func (this *QPalette) ButtonText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette10buttonTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:139
// index:0
// inline
// const QBrush & shadow()
func (this *QPalette) Shadow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette6shadowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:140
// index:0
// inline
// const QBrush & highlight()
func (this *QPalette) Highlight() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette9highlightEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:141
// index:0
// inline
// const QBrush & highlightedText()
func (this *QPalette) HighlightedText() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette15highlightedTextEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:142
// index:0
// inline
// const QBrush & link()
func (this *QPalette) Link() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette4linkEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:143
// index:0
// inline
// const QBrush & linkVisited()
func (this *QPalette) LinkVisited() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette11linkVisitedEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:147
// index:0
// bool isCopyOf(const class QPalette &)
func (this *QPalette) IsCopyOf(p unsafe.Pointer) {
	// 0: (, const QPalette & p), (p)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette8isCopyOfERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, p)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:152
// index:0
// qint64 cacheKey()
func (this *QPalette) CacheKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette8cacheKeyEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:154
// index:0
// QPalette resolve(const class QPalette &)
func (this *QPalette) Resolve(arg0 unsafe.Pointer) {
	// 0: (, const QPalette & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette7resolveERKS_", ffiqt.FFI_TYPE_VOID, this.cthis, arg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:155
// index:1
// inline
// uint resolve()
func (this *QPalette) Resolve_1() {
	// 1: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK8QPalette7resolveEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:156
// index:2
// inline
// void resolve(uint)
func (this *QPalette) Resolve_2(mask uint) {
	// 2: (, uint mask), (&mask)
	rv, err := ffiqt.InvokeQtFunc6("_ZN8QPalette7resolveEj", ffiqt.FFI_TYPE_VOID, this.cthis, &mask)
	gopp.ErrPrint(err, rv)
}

//  body block end
