package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h
// #include <qwinthumbnailtoolbar.h>
// #include <Qtwinextras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
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

type QWinThumbnailToolBar struct {
	*qtcore.QObject
}

func (this *QWinThumbnailToolBar) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QWinThumbnailToolBar) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQWinThumbnailToolBarFromPointer(cthis unsafe.Pointer) *QWinThumbnailToolBar {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QWinThumbnailToolBar{bcthis0}
}
func (*QWinThumbnailToolBar) NewFromPointer(cthis unsafe.Pointer) *QWinThumbnailToolBar {
	return NewQWinThumbnailToolBarFromPointer(cthis)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QWinThumbnailToolBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinThumbnailToolBar(QObject *)
func NewQWinThumbnailToolBar(parent *qtcore.QObject /*777 QObject **/) *QWinThumbnailToolBar {
	var convArg0 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQWinThumbnailToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinThumbnailToolBar()
func DeleteQWinThumbnailToolBar(this *QWinThumbnailToolBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindow(QWindow *)
func (this *QWinThumbnailToolBar) SetWindow(window *qtgui.QWindow /*777 QWindow **/) {
	var convArg0 = window.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar9setWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * window()
func (this *QWinThumbnailToolBar) Window() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addButton(QWinThumbnailToolButton *)
func (this *QWinThumbnailToolBar) AddButton(button *QWinThumbnailToolButton /*777 QWinThumbnailToolButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar9addButtonEP23QWinThumbnailToolButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeButton(QWinThumbnailToolButton *)
func (this *QWinThumbnailToolBar) RemoveButton(button *QWinThumbnailToolButton /*777 QWinThumbnailToolButton **/) {
	var convArg0 = button.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar12removeButtonEP23QWinThumbnailToolButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int count()
func (this *QWinThumbnailToolBar) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool iconicPixmapNotificationsEnabled()
func (this *QWinThumbnailToolBar) IconicPixmapNotificationsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar32iconicPixmapNotificationsEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconicPixmapNotificationsEnabled(_Bool)
func (this *QWinThumbnailToolBar) SetIconicPixmapNotificationsEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar35setIconicPixmapNotificationsEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap iconicThumbnailPixmap()
func (this *QWinThumbnailToolBar) IconicThumbnailPixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar21iconicThumbnailPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:80
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap iconicLivePreviewPixmap()
func (this *QWinThumbnailToolBar) IconicLivePreviewPixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar23iconicLivePreviewPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QWinThumbnailToolBar) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconicThumbnailPixmap(const QPixmap &)
func (this *QWinThumbnailToolBar) SetIconicThumbnailPixmap(arg0 *qtgui.QPixmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar24setIconicThumbnailPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconicLivePreviewPixmap(const QPixmap &)
func (this *QWinThumbnailToolBar) SetIconicLivePreviewPixmap(arg0 *qtgui.QPixmap) {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar26setIconicLivePreviewPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconicThumbnailPixmapRequested()
func (this *QWinThumbnailToolBar) IconicThumbnailPixmapRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar30iconicThumbnailPixmapRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconicLivePreviewPixmapRequested()
func (this *QWinThumbnailToolBar) IconicLivePreviewPixmapRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar32iconicLivePreviewPixmapRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

//  body block end
