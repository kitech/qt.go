package qtgui

// /usr/include/qt/QtGui/qicon.h
// #include <qicon.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QIcon struct {
	*qtrt.CObject
}
type QIcon_ITF interface {
	QIcon_PTR() *QIcon
}

func (ptr *QIcon) QIcon_PTR() *QIcon { return ptr }

func (this *QIcon) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QIcon) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQIconFromPointer(cthis unsafe.Pointer) *QIcon {
	return &QIcon{&qtrt.CObject{cthis}}
}
func (*QIcon) NewFromPointer(cthis unsafe.Pointer) *QIcon {
	return NewQIconFromPointer(cthis)
}

// /usr/include/qt/QtGui/qicon.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QIcon()
func NewQIcon() *QIcon {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:61
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QIcon(const QPixmap &)
func NewQIcon_1(pixmap *QPixmap) *QIcon {
	var convArg0 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2ERK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QIcon(const QString &)
func NewQIcon_2(fileName string) *QIcon {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2ERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:69
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QIcon(QIconEngine *)
func NewQIcon_3(engine *QIconEngine /*777 QIconEngine **/) *QIcon {
	var convArg0 = engine.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconC2EP11QIconEngine", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QIcon()
func DeleteQIcon(this *QIcon) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIconD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qicon.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QIcon &)
func (this *QIcon) Swap(other *QIcon) {
	var convArg0 = other.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap(size *qtcore.QSize, mode int, state int) *QPixmap /*123*/ {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_1(w int, h int, mode int, state int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_2(extent int, mode int, state int) *QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), extent, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:86
// index:3
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_3(window *QWindow /*777 QWindow **/, size *qtcore.QSize, mode int, state int) *QPixmap /*123*/ {
	var convArg0 = window.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:88
// index:0
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) ActualSize(size *qtcore.QSize, mode int, state int) *qtcore.QSize /*123*/ {
	var convArg0 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:89
// index:1
// Public Visibility=Default Availability=Available
// [8] QSize actualSize(QWindow *, const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) ActualSize_1(window *QWindow /*777 QWindow **/, size *qtcore.QSize, mode int, state int) *qtcore.QSize /*123*/ {
	var convArg0 = window.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QIcon) Name() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qicon.h:93
// index:0
// Public Visibility=Default Availability=Available
// [-2] void paint(QPainter *, const QRect &, Qt::Alignment, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Paint(painter *QPainter /*777 QPainter **/, rect *qtcore.QRect, alignment int, mode int, state int) {
	var convArg0 = painter.GetCthis()
	var convArg1 = rect.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainterRK5QRect6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:94
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void paint(QPainter *, int, int, int, int, Qt::Alignment, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Paint_1(painter *QPainter /*777 QPainter **/, x int, y int, w int, h int, alignment int, mode int, state int) {
	var convArg0 = painter.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon5paintEP8QPainteriiii6QFlagsIN2Qt13AlignmentFlagEENS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, x, y, w, h, alignment, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QIcon) IsNull() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6isNullEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QIcon) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QIcon) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey()
func (this *QIcon) CacheKey() int64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon8cacheKeyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPixmap(const QPixmap &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddPixmap(pixmap *QPixmap, mode int, state int) {
	var convArg0 = pixmap.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddFile(fileName string, size *qtcore.QSize, mode int, state int) {
	var tmpArg0 = qtcore.NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIsMask(_Bool)
func (this *QIcon) SetIsMask(isMask bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9setIsMaskEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), isMask)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMask()
func (this *QIcon) IsMask() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK5QIcon6isMaskEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon fromTheme(const QString &)
func (this *QIcon) FromTheme(name string) *QIcon /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QIcon_FromTheme(name string) *QIcon /*123*/ {
	var nilthis *QIcon
	rv := nilthis.FromTheme(name)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:115
// index:1
// Public static Visibility=Default Availability=Available
// [8] QIcon fromTheme(const QString &, const QIcon &)
func (this *QIcon) FromTheme_1(name string, fallback *QIcon) *QIcon /*123*/ {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = fallback.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QStringRKS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QIcon_FromTheme_1(name string, fallback *QIcon) *QIcon /*123*/ {
	var nilthis *QIcon
	rv := nilthis.FromTheme_1(name, fallback)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:116
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasThemeIcon(const QString &)
func (this *QIcon) HasThemeIcon(name string) bool {
	var tmpArg0 = qtcore.NewQString_5(name)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon12hasThemeIconERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QIcon_HasThemeIcon(name string) bool {
	var nilthis *QIcon
	rv := nilthis.HasThemeIcon(name)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:118
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList themeSearchPaths()
func (this *QIcon) ThemeSearchPaths() *qtcore.QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon16themeSearchPathsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQStringList)
	return rv2
}
func QIcon_ThemeSearchPaths() *qtcore.QStringList /*123*/ {
	var nilthis *QIcon
	rv := nilthis.ThemeSearchPaths()
	return rv
}

// /usr/include/qt/QtGui/qicon.h:119
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setThemeSearchPaths(const QStringList &)
func (this *QIcon) SetThemeSearchPaths(searchpath *qtcore.QStringList) {
	var convArg0 = searchpath.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon19setThemeSearchPathsERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QIcon_SetThemeSearchPaths(searchpath *qtcore.QStringList) {
	var nilthis *QIcon
	nilthis.SetThemeSearchPaths(searchpath)
}

// /usr/include/qt/QtGui/qicon.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString themeName()
func (this *QIcon) ThemeName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon9themeNameEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}
func QIcon_ThemeName() string {
	var nilthis *QIcon
	rv := nilthis.ThemeName()
	return rv
}

// /usr/include/qt/QtGui/qicon.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setThemeName(const QString &)
func (this *QIcon) SetThemeName(path string) {
	var tmpArg0 = qtcore.NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QIcon12setThemeNameERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QIcon_SetThemeName(path string) {
	var nilthis *QIcon
	nilthis.SetThemeName(path)
}

type QIcon__Mode = int

const QIcon__Normal QIcon__Mode = 0
const QIcon__Disabled QIcon__Mode = 1
const QIcon__Active QIcon__Mode = 2
const QIcon__Selected QIcon__Mode = 3

type QIcon__State = int

const QIcon__On QIcon__State = 0
const QIcon__Off QIcon__State = 1

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
