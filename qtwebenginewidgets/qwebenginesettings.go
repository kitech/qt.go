package qtwebenginewidgets

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h
// #include <qwebenginesettings.h>
// #include <QtWebEngineWidgets>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtquickwidgets"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtwidgets"
import "github.com/kitech/qt.go/qtprintsupport"
import "github.com/kitech/qt.go/qtquick"
import "github.com/kitech/qt.go/qtwebenginecore"

//  ext block end

//  body block begin

/*

 */
type QWebEngineSettings struct {
	*qtrt.CObject
}
type QWebEngineSettings_ITF interface {
	QWebEngineSettings_PTR() *QWebEngineSettings
}

func (ptr *QWebEngineSettings) QWebEngineSettings_PTR() *QWebEngineSettings { return ptr }

func (this *QWebEngineSettings) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineSettings) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineSettingsFromPointer(cthis unsafe.Pointer) *QWebEngineSettings {
	return &QWebEngineSettings{&qtrt.CObject{cthis}}
}
func (*QWebEngineSettings) NewFromPointer(cthis unsafe.Pointer) *QWebEngineSettings {
	return NewQWebEngineSettingsFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:105
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWebEngineSettings * globalSettings()

/*

 */
func (this *QWebEngineSettings) GlobalSettings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings14globalSettingsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWebEngineSettings_GlobalSettings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	var nilthis *QWebEngineSettings
	rv := nilthis.GlobalSettings()
	return rv
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:107
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWebEngineSettings * defaultSettings()

/*

 */
func (this *QWebEngineSettings) DefaultSettings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings15defaultSettingsEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQWebEngineSettingsFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QWebEngineSettings_DefaultSettings() *QWebEngineSettings /*777 QWebEngineSettings **/ {
	var nilthis *QWebEngineSettings
	rv := nilthis.DefaultSettings()
	return rv
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontFamily(QWebEngineSettings::FontFamily, const QString &)

/*

 */
func (this *QWebEngineSettings) SetFontFamily(which int, family string) {
	var tmpArg1 = qtcore.NewQString_5(family)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings13setFontFamilyENS_10FontFamilyERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:110
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fontFamily(QWebEngineSettings::FontFamily) const

/*

 */
func (this *QWebEngineSettings) FontFamily(which int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QWebEngineSettings10fontFamilyENS_10FontFamilyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetFontFamily(QWebEngineSettings::FontFamily)

/*

 */
func (this *QWebEngineSettings) ResetFontFamily(which int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings15resetFontFamilyENS_10FontFamilyE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), which)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFontSize(QWebEngineSettings::FontSize, int)

/*

 */
func (this *QWebEngineSettings) SetFontSize(type_ int, size int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings11setFontSizeENS_8FontSizeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_, size)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:114
// index:0
// Public Visibility=Default Availability=Available
// [4] int fontSize(QWebEngineSettings::FontSize) const

/*

 */
func (this *QWebEngineSettings) FontSize(type_ int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QWebEngineSettings8fontSizeENS_8FontSizeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetFontSize(QWebEngineSettings::FontSize)

/*

 */
func (this *QWebEngineSettings) ResetFontSize(type_ int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings13resetFontSizeENS_8FontSizeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), type_)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAttribute(QWebEngineSettings::WebAttribute, bool)

/*

 */
func (this *QWebEngineSettings) SetAttribute(attr int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings12setAttributeENS_12WebAttributeEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attr, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:118
// index:0
// Public Visibility=Default Availability=Available
// [1] bool testAttribute(QWebEngineSettings::WebAttribute) const

/*

 */
func (this *QWebEngineSettings) TestAttribute(attr int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QWebEngineSettings13testAttributeENS_12WebAttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attr)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:119
// index:0
// Public Visibility=Default Availability=Available
// [-2] void resetAttribute(QWebEngineSettings::WebAttribute)

/*

 */
func (this *QWebEngineSettings) ResetAttribute(attr int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings14resetAttributeENS_12WebAttributeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), attr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultTextEncoding(const QString &)

/*

 */
func (this *QWebEngineSettings) SetDefaultTextEncoding(encoding string) {
	var tmpArg0 = qtcore.NewQString_5(encoding)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettings22setDefaultTextEncodingERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineWidgets/qwebenginesettings.h:122
// index:0
// Public Visibility=Default Availability=Available
// [8] QString defaultTextEncoding() const

/*

 */
func (this *QWebEngineSettings) DefaultTextEncoding() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK18QWebEngineSettings19defaultTextEncodingEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

func DeleteQWebEngineSettings(this *QWebEngineSettings) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QWebEngineSettingsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QWebEngineSettings__FontFamily = int

//
const QWebEngineSettings__StandardFont QWebEngineSettings__FontFamily = 0

//
const QWebEngineSettings__FixedFont QWebEngineSettings__FontFamily = 1

//
const QWebEngineSettings__SerifFont QWebEngineSettings__FontFamily = 2

//
const QWebEngineSettings__SansSerifFont QWebEngineSettings__FontFamily = 3

//
const QWebEngineSettings__CursiveFont QWebEngineSettings__FontFamily = 4

//
const QWebEngineSettings__FantasyFont QWebEngineSettings__FontFamily = 5

//
const QWebEngineSettings__PictographFont QWebEngineSettings__FontFamily = 6

func (this *QWebEngineSettings) FontFamilyItemName(val int) string {
	switch val {
	case QWebEngineSettings__StandardFont: // 0
		return "StandardFont"
	case QWebEngineSettings__FixedFont: // 1
		return "FixedFont"
	case QWebEngineSettings__SerifFont: // 2
		return "SerifFont"
	case QWebEngineSettings__SansSerifFont: // 3
		return "SansSerifFont"
	case QWebEngineSettings__CursiveFont: // 4
		return "CursiveFont"
	case QWebEngineSettings__FantasyFont: // 5
		return "FantasyFont"
	case QWebEngineSettings__PictographFont: // 6
		return "PictographFont"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineSettings_FontFamilyItemName(val int) string {
	var nilthis *QWebEngineSettings
	return nilthis.FontFamilyItemName(val)
}

/*


 */
type QWebEngineSettings__WebAttribute = int

//
const QWebEngineSettings__AutoLoadImages QWebEngineSettings__WebAttribute = 0

//
const QWebEngineSettings__JavascriptEnabled QWebEngineSettings__WebAttribute = 1

//
const QWebEngineSettings__JavascriptCanOpenWindows QWebEngineSettings__WebAttribute = 2

//
const QWebEngineSettings__JavascriptCanAccessClipboard QWebEngineSettings__WebAttribute = 3

//
const QWebEngineSettings__LinksIncludedInFocusChain QWebEngineSettings__WebAttribute = 4

//
const QWebEngineSettings__LocalStorageEnabled QWebEngineSettings__WebAttribute = 5

//
const QWebEngineSettings__LocalContentCanAccessRemoteUrls QWebEngineSettings__WebAttribute = 6

//
const QWebEngineSettings__XSSAuditingEnabled QWebEngineSettings__WebAttribute = 7

//
const QWebEngineSettings__SpatialNavigationEnabled QWebEngineSettings__WebAttribute = 8

//
const QWebEngineSettings__LocalContentCanAccessFileUrls QWebEngineSettings__WebAttribute = 9

//
const QWebEngineSettings__HyperlinkAuditingEnabled QWebEngineSettings__WebAttribute = 10

//
const QWebEngineSettings__ScrollAnimatorEnabled QWebEngineSettings__WebAttribute = 11

//
const QWebEngineSettings__ErrorPageEnabled QWebEngineSettings__WebAttribute = 12

//
const QWebEngineSettings__PluginsEnabled QWebEngineSettings__WebAttribute = 13

//
const QWebEngineSettings__FullScreenSupportEnabled QWebEngineSettings__WebAttribute = 14

//
const QWebEngineSettings__ScreenCaptureEnabled QWebEngineSettings__WebAttribute = 15

//
const QWebEngineSettings__WebGLEnabled QWebEngineSettings__WebAttribute = 16

//
const QWebEngineSettings__Accelerated2dCanvasEnabled QWebEngineSettings__WebAttribute = 17

//
const QWebEngineSettings__AutoLoadIconsForPage QWebEngineSettings__WebAttribute = 18

//
const QWebEngineSettings__TouchIconsEnabled QWebEngineSettings__WebAttribute = 19

//
const QWebEngineSettings__FocusOnNavigationEnabled QWebEngineSettings__WebAttribute = 20

//
const QWebEngineSettings__PrintElementBackgrounds QWebEngineSettings__WebAttribute = 21

//
const QWebEngineSettings__AllowRunningInsecureContent QWebEngineSettings__WebAttribute = 22

//
const QWebEngineSettings__AllowGeolocationOnInsecureOrigins QWebEngineSettings__WebAttribute = 23

//
const QWebEngineSettings__AllowWindowActivationFromJavaScript QWebEngineSettings__WebAttribute = 24

//
const QWebEngineSettings__ShowScrollBars QWebEngineSettings__WebAttribute = 25

func (this *QWebEngineSettings) WebAttributeItemName(val int) string {
	switch val {
	case QWebEngineSettings__AutoLoadImages: // 0
		return "AutoLoadImages"
	case QWebEngineSettings__JavascriptEnabled: // 1
		return "JavascriptEnabled"
	case QWebEngineSettings__JavascriptCanOpenWindows: // 2
		return "JavascriptCanOpenWindows"
	case QWebEngineSettings__JavascriptCanAccessClipboard: // 3
		return "JavascriptCanAccessClipboard"
	case QWebEngineSettings__LinksIncludedInFocusChain: // 4
		return "LinksIncludedInFocusChain"
	case QWebEngineSettings__LocalStorageEnabled: // 5
		return "LocalStorageEnabled"
	case QWebEngineSettings__LocalContentCanAccessRemoteUrls: // 6
		return "LocalContentCanAccessRemoteUrls"
	case QWebEngineSettings__XSSAuditingEnabled: // 7
		return "XSSAuditingEnabled"
	case QWebEngineSettings__SpatialNavigationEnabled: // 8
		return "SpatialNavigationEnabled"
	case QWebEngineSettings__LocalContentCanAccessFileUrls: // 9
		return "LocalContentCanAccessFileUrls"
	case QWebEngineSettings__HyperlinkAuditingEnabled: // 10
		return "HyperlinkAuditingEnabled"
	case QWebEngineSettings__ScrollAnimatorEnabled: // 11
		return "ScrollAnimatorEnabled"
	case QWebEngineSettings__ErrorPageEnabled: // 12
		return "ErrorPageEnabled"
	case QWebEngineSettings__PluginsEnabled: // 13
		return "PluginsEnabled"
	case QWebEngineSettings__FullScreenSupportEnabled: // 14
		return "FullScreenSupportEnabled"
	case QWebEngineSettings__ScreenCaptureEnabled: // 15
		return "ScreenCaptureEnabled"
	case QWebEngineSettings__WebGLEnabled: // 16
		return "WebGLEnabled"
	case QWebEngineSettings__Accelerated2dCanvasEnabled: // 17
		return "Accelerated2dCanvasEnabled"
	case QWebEngineSettings__AutoLoadIconsForPage: // 18
		return "AutoLoadIconsForPage"
	case QWebEngineSettings__TouchIconsEnabled: // 19
		return "TouchIconsEnabled"
	case QWebEngineSettings__FocusOnNavigationEnabled: // 20
		return "FocusOnNavigationEnabled"
	case QWebEngineSettings__PrintElementBackgrounds: // 21
		return "PrintElementBackgrounds"
	case QWebEngineSettings__AllowRunningInsecureContent: // 22
		return "AllowRunningInsecureContent"
	case QWebEngineSettings__AllowGeolocationOnInsecureOrigins: // 23
		return "AllowGeolocationOnInsecureOrigins"
	case QWebEngineSettings__AllowWindowActivationFromJavaScript: // 24
		return "AllowWindowActivationFromJavaScript"
	case QWebEngineSettings__ShowScrollBars: // 25
		return "ShowScrollBars"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineSettings_WebAttributeItemName(val int) string {
	var nilthis *QWebEngineSettings
	return nilthis.WebAttributeItemName(val)
}

/*


 */
type QWebEngineSettings__FontSize = int

//
const QWebEngineSettings__MinimumFontSize QWebEngineSettings__FontSize = 0

//
const QWebEngineSettings__MinimumLogicalFontSize QWebEngineSettings__FontSize = 1

//
const QWebEngineSettings__DefaultFontSize QWebEngineSettings__FontSize = 2

//
const QWebEngineSettings__DefaultFixedFontSize QWebEngineSettings__FontSize = 3

func (this *QWebEngineSettings) FontSizeItemName(val int) string {
	switch val {
	case QWebEngineSettings__MinimumFontSize: // 0
		return "MinimumFontSize"
	case QWebEngineSettings__MinimumLogicalFontSize: // 1
		return "MinimumLogicalFontSize"
	case QWebEngineSettings__DefaultFontSize: // 2
		return "DefaultFontSize"
	case QWebEngineSettings__DefaultFixedFontSize: // 3
		return "DefaultFixedFontSize"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineSettings_FontSizeItemName(val int) string {
	var nilthis *QWebEngineSettings
	return nilthis.FontSizeItemName(val)
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
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtquickwidgets.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtwidgets.KeepMe()
	}
	if false {
		qtprintsupport.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
	if false {
		qtwebenginecore.KeepMe()
	}
}

//  keep block end
