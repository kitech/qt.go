package qtgui

// /usr/include/qt/QtGui/qicon.h
// #include <qicon.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2ERK7QPixmap", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:68
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QIcon(const QString &)
func NewQIcon_2(fileName *qtcore.QString) *QIcon {
	var convArg0 = fileName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
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
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconC2EP11QIconEngine", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQIconFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQIcon)
	return gothis
}

// /usr/include/qt/QtGui/qicon.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QIcon()
func DeleteQIcon(this *QIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIconD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qicon.h:76
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QIcon &)
func (this *QIcon) Swap(other *QIcon) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:81
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap pixmap(const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap(size *qtcore.QSize, mode int, state int) *QPixmap /*123*/ {
	var convArg0 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapERK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:82
// index:1
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_1(w int, h int, mode int, state int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiiNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), w, h, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:84
// index:2
// Public inline Visibility=Default Availability=Available
// [32] QPixmap pixmap(int, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) Pixmap_2(extent int, mode int, state int) *QPixmap /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEiNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), extent, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6pixmapEP7QWindowRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10actualSizeERK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
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
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10actualSizeEP7QWindowRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQSizeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQSize)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:91
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QIcon) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qicon.h:97
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isNull()
func (this *QIcon) IsNull() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6isNullEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QIcon) IsDetached() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:99
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QIcon) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:104
// index:0
// Public Visibility=Default Availability=Available
// [8] qint64 cacheKey()
func (this *QIcon) CacheKey() int64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon8cacheKeyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int64(rv) // 222
}

// /usr/include/qt/QtGui/qicon.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addPixmap(const QPixmap &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddPixmap(pixmap *QPixmap, mode int, state int) {
	var convArg0 = pixmap.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9addPixmapERK7QPixmapNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:107
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addFile(const QString &, const QSize &, enum QIcon::Mode, enum QIcon::State)
func (this *QIcon) AddFile(fileName *qtcore.QString, size *qtcore.QSize, mode int, state int) {
	var convArg0 = fileName.GetCthis()
	var convArg1 = size.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon7addFileERK7QStringRK5QSizeNS_4ModeENS_5StateE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, mode, state)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIsMask(_Bool)
func (this *QIcon) SetIsMask(isMask bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9setIsMaskEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), isMask)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:112
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isMask()
func (this *QIcon) IsMask() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK5QIcon6isMaskEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qicon.h:114
// index:0
// Public static Visibility=Default Availability=Available
// [8] QIcon fromTheme(const QString &)
func (this *QIcon) FromTheme(name *qtcore.QString) *QIcon /*123*/ {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QIcon_FromTheme(name *qtcore.QString) *QIcon /*123*/ {
	var nilthis *QIcon
	rv := nilthis.FromTheme(name)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:115
// index:1
// Public static Visibility=Default Availability=Available
// [8] QIcon fromTheme(const QString &, const QIcon &)
func (this *QIcon) FromTheme_1(name *qtcore.QString, fallback *QIcon) *QIcon /*123*/ {
	var convArg0 = name.GetCthis()
	var convArg1 = fallback.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9fromThemeERK7QStringRKS_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQIconFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQIcon)
	return rv2
}
func QIcon_FromTheme_1(name *qtcore.QString, fallback *QIcon) *QIcon /*123*/ {
	var nilthis *QIcon
	rv := nilthis.FromTheme_1(name, fallback)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:116
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool hasThemeIcon(const QString &)
func (this *QIcon) HasThemeIcon(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon12hasThemeIconERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QIcon_HasThemeIcon(name *qtcore.QString) bool {
	var nilthis *QIcon
	rv := nilthis.HasThemeIcon(name)
	return rv
}

// /usr/include/qt/QtGui/qicon.h:119
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setThemeSearchPaths(const QStringList &)
func (this *QIcon) SetThemeSearchPaths(searchpath *qtcore.QStringList) {
	var convArg0 = searchpath.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon19setThemeSearchPathsERK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QIcon_SetThemeSearchPaths(searchpath *qtcore.QStringList) {
	var nilthis *QIcon
	nilthis.SetThemeSearchPaths(searchpath)
}

// /usr/include/qt/QtGui/qicon.h:121
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString themeName()
func (this *QIcon) ThemeName() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon9themeNameEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}
func QIcon_ThemeName() *qtcore.QString /*123*/ {
	var nilthis *QIcon
	rv := nilthis.ThemeName()
	return rv
}

// /usr/include/qt/QtGui/qicon.h:122
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setThemeName(const QString &)
func (this *QIcon) SetThemeName(path *qtcore.QString) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN5QIcon12setThemeNameERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QIcon_SetThemeName(path *qtcore.QString) {
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
