package qtwinextras

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"

func init() {
	if false {
		_ = unsafe.Pointer(uintptr(0))
	}
	if false {
		qtrt.KeepMe()
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

//  header block end

//  body block begin
// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:81
// index:0
// Invalid Visibility=Default Availability=Available
// [8] HBITMAP toHBITMAP(const QPixmap &, QtWin::HBitmapFormat)

/*

 */
func ToHBITMAP(p qtgui.QPixmap_ITF, format int) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPixmap_PTR() != nil {
		convArg0 = p.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin9toHBITMAPERK7QPixmapNS_13HBitmapFormatE", qtrt.FFI_TYPE_POINTER, convArg0, format)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:85
// index:0
// Invalid Visibility=Default Availability=Available
// [32] QPixmap fromHICON(HICON)

/*

 */
func FromHICON(icon unsafe.Pointer /*666*/) *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin9fromHICONEP7HICON__", qtrt.FFI_TYPE_POINTER, icon)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:87
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QRegion fromHRGN(HRGN)

/*

 */
func FromHRGN(hrgn unsafe.Pointer /*666*/) *qtgui.QRegion /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin8fromHRGNEP6HRGN__", qtrt.FFI_TYPE_POINTER, hrgn)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQRegionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQRegion)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:83
// index:0
// Invalid Visibility=Default Availability=Available
// [8] HICON toHICON(const QPixmap &)

/*

 */
func ToHICON(p qtgui.QPixmap_ITF) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if p != nil && p.QPixmap_PTR() != nil {
		convArg0 = p.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin7toHICONERK7QPixmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:86
// index:0
// Invalid Visibility=Default Availability=Available
// [8] HRGN toHRGN(const QRegion &)

/*

 */
func ToHRGN(region qtgui.QRegion_ITF) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg0 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin6toHRGNERK7QRegion", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:114
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void setCurrentProcessExplicitAppUserModelID(const QString &)

/*

 */
func SetCurrentProcessExplicitAppUserModelID(id string) {
	var tmpArg0 = qtcore.NewQString_5(id)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin39setCurrentProcessExplicitAppUserModelIDERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:95
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void setWindowExcludedFromPeek(QWindow *, bool)

/*

 */
func SetWindowExcludedFromPeek(window qtgui.QWindow_ITF /*777 QWindow **/, exclude bool) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin25setWindowExcludedFromPeekEP7QWindowb", qtrt.FFI_TYPE_POINTER, convArg0, exclude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:124
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void setWindowExcludedFromPeek(QWidget *, bool)

/*

 */
func SetWindowExcludedFromPeek_1(window qtwidgets.QWidget_ITF /*777 QWidget **/, exclude bool) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin25setWindowExcludedFromPeekEP7QWidgetb", qtrt.FFI_TYPE_POINTER, convArg0, exclude)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:102
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void extendFrameIntoClientArea(QWindow *, int, int, int, int)

/*

 */
func ExtendFrameIntoClientArea(window qtgui.QWindow_ITF /*777 QWindow **/, left int, top int, right int, bottom int) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin25extendFrameIntoClientAreaEP7QWindowiiii", qtrt.FFI_TYPE_POINTER, convArg0, left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:103
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void extendFrameIntoClientArea(QWindow *, const QMargins &)

/*

 */
func ExtendFrameIntoClientArea_1(window qtgui.QWindow_ITF /*777 QWindow **/, margins qtcore.QMargins_ITF) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg1 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin25extendFrameIntoClientAreaEP7QWindowRK8QMargins", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:172
// index:2
// Invalid inline Visibility=Default Availability=Available
// [-2] void extendFrameIntoClientArea(QWidget *, int, int, int, int)

/*

 */
func ExtendFrameIntoClientArea_2(window qtwidgets.QWidget_ITF /*777 QWidget **/, left int, top int, right int, bottom int) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin25extendFrameIntoClientAreaEP7QWidgetiiii", qtrt.FFI_TYPE_POINTER, convArg0, left, top, right, bottom)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:166
// index:3
// Invalid inline Visibility=Default Availability=Available
// [-2] void extendFrameIntoClientArea(QWidget *, const QMargins &)

/*

 */
func ExtendFrameIntoClientArea_3(window qtwidgets.QWidget_ITF /*777 QWidget **/, margins qtcore.QMargins_ITF) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if margins != nil && margins.QMargins_PTR() != nil {
		convArg1 = margins.QMargins_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin25extendFrameIntoClientAreaEP7QWidgetRK8QMargins", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:96
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool isWindowExcludedFromPeek(QWindow *)

/*

 */
func IsWindowExcludedFromPeek(window qtgui.QWindow_ITF /*777 QWindow **/) bool {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin24isWindowExcludedFromPeekEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:130
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool isWindowExcludedFromPeek(QWidget *)

/*

 */
func IsWindowExcludedFromPeek_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin24isWindowExcludedFromPeekEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:108
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void disableBlurBehindWindow(QWindow *)

/*

 */
func DisableBlurBehindWindow(window qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin23disableBlurBehindWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:196
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void disableBlurBehindWindow(QWidget *)

/*

 */
func DisableBlurBehindWindow_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin23disableBlurBehindWindowEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:98
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool isWindowPeekDisallowed(QWindow *)

/*

 */
func IsWindowPeekDisallowed(window qtgui.QWindow_ITF /*777 QWindow **/) bool {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22isWindowPeekDisallowedEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:144
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool isWindowPeekDisallowed(QWidget *)

/*

 */
func IsWindowPeekDisallowed_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) bool {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22isWindowPeekDisallowedEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:90
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString errorStringFromHresult(HRESULT)

/*

 */
func ErrorStringFromHresult(hresult int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22errorStringFromHresultEi", qtrt.FFI_TYPE_POINTER, hresult)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:106
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void enableBlurBehindWindow(QWindow *, const QRegion &)

/*

 */
func EnableBlurBehindWindow(window qtgui.QWindow_ITF /*777 QWindow **/, region qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg1 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22enableBlurBehindWindowEP7QWindowRK7QRegion", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:107
// index:1
// Invalid Visibility=Default Availability=Available
// [-2] void enableBlurBehindWindow(QWindow *)

/*

 */
func EnableBlurBehindWindow_1(window qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22enableBlurBehindWindowEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:184
// index:2
// Invalid inline Visibility=Default Availability=Available
// [-2] void enableBlurBehindWindow(QWidget *, const QRegion &)

/*

 */
func EnableBlurBehindWindow_2(window qtwidgets.QWidget_ITF /*777 QWidget **/, region qtgui.QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if region != nil && region.QRegion_PTR() != nil {
		convArg1 = region.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22enableBlurBehindWindowEP7QWidgetRK7QRegion", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:190
// index:3
// Invalid inline Visibility=Default Availability=Available
// [-2] void enableBlurBehindWindow(QWidget *)

/*

 */
func EnableBlurBehindWindow_3(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin22enableBlurBehindWindowEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:119
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void taskbarActivateTabAlt(QWindow *)

/*

 */
func TaskbarActivateTabAlt(arg0 qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWindow_PTR() != nil {
		convArg0 = arg0.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21taskbarActivateTabAltEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:214
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void taskbarActivateTabAlt(QWidget *)

/*

 */
func TaskbarActivateTabAlt_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21taskbarActivateTabAltEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:99
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void setWindowFlip3DPolicy(QWindow *, QtWin::WindowFlip3DPolicy)

/*

 */
func SetWindowFlip3DPolicy(window qtgui.QWindow_ITF /*777 QWindow **/, policy int) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21setWindowFlip3DPolicyEP7QWindowNS_18WindowFlip3DPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:152
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void setWindowFlip3DPolicy(QWidget *, QtWin::WindowFlip3DPolicy)

/*

 */
func SetWindowFlip3DPolicy_1(window qtwidgets.QWidget_ITF /*777 QWidget **/, policy int) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21setWindowFlip3DPolicyEP7QWidgetNS_18WindowFlip3DPolicyE", qtrt.FFI_TYPE_POINTER, convArg0, policy)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:97
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void setWindowDisallowPeek(QWindow *, bool)

/*

 */
func SetWindowDisallowPeek(window qtgui.QWindow_ITF /*777 QWindow **/, disallow bool) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21setWindowDisallowPeekEP7QWindowb", qtrt.FFI_TYPE_POINTER, convArg0, disallow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:138
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void setWindowDisallowPeek(QWidget *, bool)

/*

 */
func SetWindowDisallowPeek_1(window qtwidgets.QWidget_ITF /*777 QWidget **/, disallow bool) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21setWindowDisallowPeekEP7QWidgetb", qtrt.FFI_TYPE_POINTER, convArg0, disallow)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:111
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void setCompositionEnabled(bool)

/*

 */
func SetCompositionEnabled(enabled bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21setCompositionEnabledEb", qtrt.FFI_TYPE_POINTER, enabled)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:93
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QColor realColorizationColor()

/*

 */
func RealColorizationColor() *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin21realColorizationColorEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:116
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void markFullscreenWindow(QWindow *, bool)

/*

 */
func MarkFullscreenWindow(arg0 qtgui.QWindow_ITF /*777 QWindow **/, fullscreen bool) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWindow_PTR() != nil {
		convArg0 = arg0.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin20markFullscreenWindowEP7QWindowb", qtrt.FFI_TYPE_POINTER, convArg0, fullscreen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:202
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void markFullscreenWindow(QWidget *, bool)

/*

 */
func MarkFullscreenWindow_1(window qtwidgets.QWidget_ITF /*777 QWidget **/, fullscreen bool) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin20markFullscreenWindowEP7QWidgetb", qtrt.FFI_TYPE_POINTER, convArg0, fullscreen)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:110
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool isCompositionEnabled()

/*

 */
func IsCompositionEnabled() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin20isCompositionEnabledEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:112
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool isCompositionOpaque()

/*

 */
func IsCompositionOpaque() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin19isCompositionOpaqueEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:100
// index:0
// Invalid Visibility=Default Availability=Available
// [4] QtWin::WindowFlip3DPolicy windowFlip3DPolicy(QWindow *)

/*

 */
func WindowFlip3DPolicy(arg0 qtgui.QWindow_ITF /*777 QWindow **/) int {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWindow_PTR() != nil {
		convArg0 = arg0.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin18windowFlip3DPolicyEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:158
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] QtWin::WindowFlip3DPolicy windowFlip3DPolicy(QWidget *)

/*

 */
func WindowFlip3DPolicy_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) int {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin18windowFlip3DPolicyEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:118
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void taskbarActivateTab(QWindow *)

/*

 */
func TaskbarActivateTab(arg0 qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWindow_PTR() != nil {
		convArg0 = arg0.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin18taskbarActivateTabEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:208
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void taskbarActivateTab(QWidget *)

/*

 */
func TaskbarActivateTab_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin18taskbarActivateTabEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:104
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void resetExtendedFrame(QWindow *)

/*

 */
func ResetExtendedFrame(window qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWindow_PTR() != nil {
		convArg0 = window.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin18resetExtendedFrameEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:178
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void resetExtendedFrame(QWidget *)

/*

 */
func ResetExtendedFrame_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin18resetExtendedFrameEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:89
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString stringFromHresult(HRESULT)

/*

 */
func StringFromHresult(hresult int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin17stringFromHresultEi", qtrt.FFI_TYPE_POINTER, hresult)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:92
// index:0
// Invalid Visibility=Default Availability=Available
// [16] QColor colorizationColor(bool *)

/*

 */
func ColorizationColor(opaqueBlend *bool) *qtgui.QColor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin17colorizationColorEPb", qtrt.FFI_TYPE_POINTER, opaqueBlend)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQColor)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:121
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void taskbarDeleteTab(QWindow *)

/*

 */
func TaskbarDeleteTab(arg0 qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWindow_PTR() != nil {
		convArg0 = arg0.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin16taskbarDeleteTabEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:226
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void taskbarDeleteTab(QWidget *)

/*

 */
func TaskbarDeleteTab_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin16taskbarDeleteTabEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:84
// index:0
// Invalid Visibility=Default Availability=Available
// [32] QImage imageFromHBITMAP(HDC, HBITMAP, int, int)

/*

 */
func ImageFromHBITMAP(hdc unsafe.Pointer /*666*/, bitmap unsafe.Pointer /*666*/, width int, height int) *qtgui.QImage /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin16imageFromHBITMAPEP5HDC__P9HBITMAP__ii", qtrt.FFI_TYPE_POINTER, hdc, bitmap, width, height)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQImageFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQImage)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:120
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void taskbarAddTab(QWindow *)

/*

 */
func TaskbarAddTab(arg0 qtgui.QWindow_ITF /*777 QWindow **/) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QWindow_PTR() != nil {
		convArg0 = arg0.QWindow_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin13taskbarAddTabEP7QWindow", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:220
// index:1
// Invalid inline Visibility=Default Availability=Available
// [-2] void taskbarAddTab(QWidget *)

/*

 */
func TaskbarAddTab_1(window qtwidgets.QWidget_ITF /*777 QWidget **/) {
	var convArg0 unsafe.Pointer
	if window != nil && window.QWidget_PTR() != nil {
		convArg0 = window.QWidget_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin13taskbarAddTabEP7QWidget", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:82
// index:0
// Invalid Visibility=Default Availability=Available
// [32] QPixmap fromHBITMAP(HBITMAP, QtWin::HBitmapFormat)

/*

 */
func FromHBITMAP(bitmap unsafe.Pointer /*666*/, format int) *qtgui.QPixmap /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin11fromHBITMAPEP9HBITMAP__NS_13HBitmapFormatE", qtrt.FFI_TYPE_POINTER, bitmap, format)
	qtrt.ErrPrint(err, rv)
	rv2 := qtgui.NewQPixmapFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtgui.DeleteQPixmap)
	return rv2
}

// /usr/include/qt/QtWinExtras/../../src/winextras/qwinfunctions.h:80
// index:0
// Invalid Visibility=Default Availability=Available
// [8] HBITMAP createMask(const QBitmap &)

/*

 */
func CreateMask(bitmap qtgui.QBitmap_ITF) unsafe.Pointer /*666*/ {
	var convArg0 unsafe.Pointer
	if bitmap != nil && bitmap.QBitmap_PTR() != nil {
		convArg0 = bitmap.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN5QtWin10createMaskERK7QBitmap", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

//  body block end
