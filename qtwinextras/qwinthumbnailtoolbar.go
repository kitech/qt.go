package qtwinextras

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h
// #include <qwinthumbnailtoolbar.h>
// #include <Qtwinextras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 27
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

//  ext block end

//  body block begin

/*

 */
type QWinThumbnailToolBar struct {
	*qtcore.QObject
}
type QWinThumbnailToolBar_ITF interface {
	qtcore.QObject_ITF
	QWinThumbnailToolBar_PTR() *QWinThumbnailToolBar
}

func (ptr *QWinThumbnailToolBar) QWinThumbnailToolBar_PTR() *QWinThumbnailToolBar { return ptr }

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
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QWinThumbnailToolBar) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinThumbnailToolBar(QObject *)

/*
Constructs a QWinThumbnailToolBar with the specified parent.

If parent is an instance of QWindow, it is automatically assigned as the thumbnail toolbar's window.
*/
func NewQWinThumbnailToolBar(parent qtcore.QObject_ITF /*777 QObject **/) *QWinThumbnailToolBar {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinThumbnailToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWinThumbnailToolBar")
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:64
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWinThumbnailToolBar(QObject *)

/*
Constructs a QWinThumbnailToolBar with the specified parent.

If parent is an instance of QWindow, it is automatically assigned as the thumbnail toolbar's window.
*/
func NewQWinThumbnailToolBar__() *QWinThumbnailToolBar {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBarC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWinThumbnailToolBarFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QWinThumbnailToolBar")
	return gothis
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:65
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QWinThumbnailToolBar()

/*

 */
func DeleteQWinThumbnailToolBar(this *QWinThumbnailToolBar) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBarD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:67
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWindow(QWindow *)

/*

 */
func (this *QWinThumbnailToolBar) SetWindow(window qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar9setWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:68
// index:0
// Public Visibility=Default Availability=Available
// [8] QWindow * window() const

/*

 */
func (this *QWinThumbnailToolBar) Window() *qtgui.QWindow /*777 QWindow **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar6windowEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtgui.NewQWindowFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:70
// index:0
// Public Visibility=Default Availability=Available
// [-2] void addButton(QWinThumbnailToolButton *)

/*
Adds a button to the thumbnail toolbar.

Note: The number of buttons is limited to 7.
*/
func (this *QWinThumbnailToolBar) AddButton(button QWinThumbnailToolButton_ITF /*777 QWinThumbnailToolButton **/) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QWinThumbnailToolButton_PTR() != nil {
		convArg0 = button.QWinThumbnailToolButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar9addButtonEP23QWinThumbnailToolButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void removeButton(QWinThumbnailToolButton *)

/*
Removes the button from the thumbnail toolbar.
*/
func (this *QWinThumbnailToolBar) RemoveButton(button QWinThumbnailToolButton_ITF /*777 QWinThumbnailToolButton **/) {
	var convArg0 unsafe.Pointer
	if button != nil && button.QWinThumbnailToolButton_PTR() != nil {
		convArg0 = button.QWinThumbnailToolButton_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar12removeButtonEP23QWinThumbnailToolButton", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:74
// index:0
// Public Visibility=Default Availability=Available
// [4] int count() const

/*

 */
func (this *QWinThumbnailToolBar) Count() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar5countEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:76
// index:0
// Public Visibility=Default Availability=Available
// [1] bool iconicPixmapNotificationsEnabled() const

/*

 */
func (this *QWinThumbnailToolBar) IconicPixmapNotificationsEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar32iconicPixmapNotificationsEnabledEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:77
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconicPixmapNotificationsEnabled(bool)

/*

 */
func (this *QWinThumbnailToolBar) SetIconicPixmapNotificationsEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar35setIconicPixmapNotificationsEnabledEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:79
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap iconicThumbnailPixmap() const

/*

 */
func (this *QWinThumbnailToolBar) IconicThumbnailPixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar21iconicThumbnailPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:80
// index:0
// Public Visibility=Default Availability=Available
// [32] QPixmap iconicLivePreviewPixmap() const

/*

 */
func (this *QWinThumbnailToolBar) IconicLivePreviewPixmap() *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK20QWinThumbnailToolBar23iconicLivePreviewPixmapEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:83
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Removes all buttons from the thumbnail toolbar.
*/
func (this *QWinThumbnailToolBar) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconicThumbnailPixmap(const QPixmap &)

/*

 */
func (this *QWinThumbnailToolBar) SetIconicThumbnailPixmap(arg0 qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar24setIconicThumbnailPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setIconicLivePreviewPixmap(const QPixmap &)

/*

 */
func (this *QWinThumbnailToolBar) SetIconicLivePreviewPixmap(arg0 qtgui.QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QPixmap_PTR() != nil {
		convArg0 = arg0.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar26setIconicLivePreviewPixmapERK7QPixmap", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconicThumbnailPixmapRequested()

/*
This signal is emitted when the operating system requests a new iconic thumbnail pixmap, typically when the thumbnail is shown.

This function was introduced in  Qt 5.4.

See also iconicThumbnailPixmap.
*/
func (this *QWinThumbnailToolBar) IconicThumbnailPixmapRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar30iconicThumbnailPixmapRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinthumbnailtoolbar.h:89
// index:0
// Public Visibility=Default Availability=Available
// [-2] void iconicLivePreviewPixmapRequested()

/*
This signal is emitted when the operating system requests a new iconic live preview pixmap, typically when the user ALT-tabs to the application.

This function was introduced in  Qt 5.4.

See also iconicLivePreviewPixmap.
*/
func (this *QWinThumbnailToolBar) IconicLivePreviewPixmapRequested() {
	rv, err := qtrt.InvokeQtFunc6("_ZN20QWinThumbnailToolBar32iconicLivePreviewPixmapRequestedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  keep block end
