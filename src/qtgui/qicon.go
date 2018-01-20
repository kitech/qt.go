//  header block begin
// /usr/include/qt/QtGui/qicon.h
// #include <qicon.h>
// #include <QtGui>
package qtgui

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
type QIcon struct {
	*qtrt.CObject
}

func (this *QIcon) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtGui/qicon.h:60
// index:0
// void QIcon()
func NewQIcon() *QIcon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}
func NewQIconFromPointer(cthis unsafe.Pointer) *QIcon {
	return &QIcon{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qicon.h:61
// index:1
// void QIcon(const class QPixmap &)
func NewQIcon_1(pixmap unsafe.Pointer) *QIcon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2ERK7QPixmap", ffiqt.FFI_TYPE_VOID, cthis, pixmap)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:68
// index:2
// void QIcon(const class QString &)
func NewQIcon_2(fileName unsafe.Pointer) *QIcon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, fileName)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:69
// index:3
// void QIcon(class QIconEngine *)
func NewQIcon_3(engine unsafe.Pointer) *QIcon {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2EP11QIconEngine", ffiqt.FFI_TYPE_VOID, cthis, engine)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:70
// index:0
// void ~QIcon()
func DeleteQIcon(*QIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:76
// index:0
// inline
// void swap(class QIcon &)
func (this *QIcon) Swap(other unsafe.Pointer) {
	// 0: (, other QIcon &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// QPixmap pixmap(const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap(size unsafe.Pointer, mode int, state int) {
	// 0: (, size const QSize &, mode QIcon::Mode, state QIcon::State), (size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// inline
// QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_1(w int, h int, mode int, state int) {
	// 1: (, w int, h int, mode QIcon::Mode, state QIcon::State), (&w, &h, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &w, &h, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// inline
// QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_2(extent int, mode int, state int) {
	// 2: (, extent int, mode QIcon::Mode, state QIcon::State), (&extent, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &extent, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:86
// index:3
// QPixmap pixmap(class QWindow *, const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_3(window unsafe.Pointer, size unsafe.Pointer, mode int, state int) {
	// 3: (, window QWindow *, size const QSize &, mode QIcon::Mode, state QIcon::State), (window, size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), window, size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:88
// index:0
// QSize actualSize(const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) ActualSize(size unsafe.Pointer, mode int, state int) {
	// 0: (, size const QSize &, mode QIcon::Mode, state QIcon::State), (size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:89
// index:1
// QSize actualSize(class QWindow *, const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) ActualSize_1(window unsafe.Pointer, size unsafe.Pointer, mode int, state int) {
	// 1: (, window QWindow *, size const QSize &, mode QIcon::Mode, state QIcon::State), (window, size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), window, size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:91
// index:0
// QString name()
func (this *QIcon) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon4nameEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:93
// index:0
// void paint(class QPainter *, const class QRect &, Qt::Alignment, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Paint(painter unsafe.Pointer, rect unsafe.Pointer, alignment int, mode int, state int) {
	// 0: (, painter QPainter *, rect const QRect &, QFlags<Qt::AlignmentFlag> alignment, mode QIcon::Mode, state QIcon::State), (painter, rect, &alignment, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainterRK5QRect6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, rect, &alignment, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:94
// index:1
// inline
// void paint(class QPainter *, int, int, int, int, Qt::Alignment, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Paint_1(painter unsafe.Pointer, x int, y int, w int, h int, alignment int, mode int, state int) {
	// 1: (, painter QPainter *, x int, y int, w int, h int, QFlags<Qt::AlignmentFlag> alignment, mode QIcon::Mode, state QIcon::State), (painter, &x, &y, &w, &h, &alignment, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainteriiii6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), painter, &x, &y, &w, &h, &alignment, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:97
// index:0
// bool isNull()
func (this *QIcon) IsNull() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6isNullEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:98
// index:0
// bool isDetached()
func (this *QIcon) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:99
// index:0
// void detach()
func (this *QIcon) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon6detachEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:104
// index:0
// qint64 cacheKey()
func (this *QIcon) CacheKey() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon8cacheKeyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// void addPixmap(const class QPixmap &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddPixmap(pixmap unsafe.Pointer, mode int, state int) {
	// 0: (, pixmap const QPixmap &, mode QIcon::Mode, state QIcon::State), (pixmap, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), pixmap, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// void addFile(const class QString &, const class QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddFile(fileName unsafe.Pointer, size unsafe.Pointer, mode int, state int) {
	// 0: (, fileName const QString &, size const QSize &, mode QIcon::Mode, state QIcon::State), (fileName, size, &mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fileName, size, &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:109
// index:0
// QList<QSize> availableSizes(enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AvailableSizes(mode int, state int) {
	// 0: (, mode QIcon::Mode, state QIcon::State), (&mode, &state)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon14availableSizesENS_4ModeENS_5StateE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mode, &state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:111
// index:0
// void setIsMask(_Bool)
func (this *QIcon) SetIsMask(isMask bool) {
	// 0: (, isMask bool), (&isMask)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9setIsMaskEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &isMask)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:112
// index:0
// bool isMask()
func (this *QIcon) IsMask() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6isMaskEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:114
// index:0
// static
// QIcon fromTheme(const class QString &)
func (this *QIcon) FromTheme(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_FromTheme(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	var nilthis *QIcon
	nilthis.FromTheme(name)
}

// /usr/include/qt/QtGui/qicon.h:115
// index:1
// static
// QIcon fromTheme(const class QString &, const class QIcon &)
func (this *QIcon) FromTheme_1(name unsafe.Pointer, fallback unsafe.Pointer) {
	// 1: (name const QString &, fallback const QIcon &), (name, fallback)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QStringRKS_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_FromTheme_1(name unsafe.Pointer, fallback unsafe.Pointer) {
	// 1: (name const QString &, fallback const QIcon &), (name, fallback)
	var nilthis *QIcon
	nilthis.FromTheme_1(name, fallback)
}

// /usr/include/qt/QtGui/qicon.h:116
// index:0
// static
// bool hasThemeIcon(const class QString &)
func (this *QIcon) HasThemeIcon(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon12hasThemeIconERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_HasThemeIcon(name unsafe.Pointer) {
	// 0: (name const QString &), (name)
	var nilthis *QIcon
	nilthis.HasThemeIcon(name)
}

// /usr/include/qt/QtGui/qicon.h:118
// index:0
// static
// QStringList themeSearchPaths()
func (this *QIcon) ThemeSearchPaths() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon16themeSearchPathsEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_ThemeSearchPaths() {
	// 0: (), ()
	var nilthis *QIcon
	nilthis.ThemeSearchPaths()
}

// /usr/include/qt/QtGui/qicon.h:119
// index:0
// static
// void setThemeSearchPaths(const class QStringList &)
func (this *QIcon) SetThemeSearchPaths(searchpath unsafe.Pointer) {
	// 0: (searchpath const QStringList &), (searchpath)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon19setThemeSearchPathsERK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_SetThemeSearchPaths(searchpath unsafe.Pointer) {
	// 0: (searchpath const QStringList &), (searchpath)
	var nilthis *QIcon
	nilthis.SetThemeSearchPaths(searchpath)
}

// /usr/include/qt/QtGui/qicon.h:121
// index:0
// static
// QString themeName()
func (this *QIcon) ThemeName() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9themeNameEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_ThemeName() {
	// 0: (), ()
	var nilthis *QIcon
	nilthis.ThemeName()
}

// /usr/include/qt/QtGui/qicon.h:122
// index:0
// static
// void setThemeName(const class QString &)
func (this *QIcon) SetThemeName(path unsafe.Pointer) {
	// 0: (path const QString &), (path)
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon12setThemeNameERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QIcon_SetThemeName(path unsafe.Pointer) {
	// 0: (path const QString &), (path)
	var nilthis *QIcon
	nilthis.SetThemeName(path)
}

//  body block end
