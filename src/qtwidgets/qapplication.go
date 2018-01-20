//  header block begin
// /usr/include/qt/QtWidgets/qapplication.h
// #include <qapplication.h>
// #include <QtWidgets>
package qtwidgets

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 18
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
type QApplication struct {
	*qtgui.QGuiApplication
}

func (this *QApplication) GetCthis() unsafe.Pointer {
	return this.QGuiApplication.GetCthis()
}
func NewQApplicationFromPointer(cthis unsafe.Pointer) *QApplication {
	bcthis0 := qtgui.NewQGuiApplicationFromPointer(cthis)
	return &QApplication{bcthis0}
}

// /usr/include/qt/QtWidgets/qapplication.h:74
// index:0
// Public virtual
// const QMetaObject * metaObject()
func (this *QApplication) MetaObject() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QApplication10metaObjectEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:94
// index:0
// Public
// void QApplication(int &, char **, int)
func NewQApplication(argc int, argv []string, arg2 int) *QApplication {
	cthis := qtrt.Calloc(1, 256) // 16
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplicationC2ERiPPci", ffiqt.FFI_TYPE_VOID, cthis, &argc, convArg1, &arg2)
	gopp.ErrPrint(err, rv)
	gothis := NewQApplicationFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtWidgets/qapplication.h:96
// index:0
// Public virtual
// void ~QApplication()
func DeleteQApplication(*QApplication) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplicationD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:98
// index:0
// Public static
// QStyle * style()
func (this *QApplication) Style() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication5styleEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Style() {
	var nilthis *QApplication
	nilthis.Style()
}

// /usr/include/qt/QtWidgets/qapplication.h:99
// index:0
// Public static
// void setStyle(class QStyle *)
func (this *QApplication) SetStyle(arg0 unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8setStyleEP6QStyle", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetStyle(arg0 unsafe.Pointer) {
	var nilthis *QApplication
	nilthis.SetStyle(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:100
// index:1
// Public static
// QStyle * setStyle(const class QString &)
func (this *QApplication) SetStyle_1(arg0 *qtcore.QString) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8setStyleERK7QString", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_SetStyle_1(arg0 *qtcore.QString) {
	var nilthis *QApplication
	nilthis.SetStyle_1(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:103
// index:0
// Public static
// int colorSpec()
func (this *QApplication) ColorSpec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication9colorSpecEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_ColorSpec() {
	var nilthis *QApplication
	nilthis.ColorSpec()
}

// /usr/include/qt/QtWidgets/qapplication.h:104
// index:0
// Public static
// void setColorSpec(int)
func (this *QApplication) SetColorSpec(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication12setColorSpecEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetColorSpec(arg0 int) {
	var nilthis *QApplication
	nilthis.SetColorSpec(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:111
// index:0
// Public static
// QPalette palette(const class QWidget *)
func (this *QApplication) Palette(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7paletteEPK7QWidget", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Palette(arg0 unsafe.Pointer) {
	var nilthis *QApplication
	nilthis.Palette(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:112
// index:1
// Public static
// QPalette palette(const char *)
func (this *QApplication) Palette_1(className string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7paletteEPKc", ffiqt.FFI_TYPE_POINTER, className)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Palette_1(className string) {
	var nilthis *QApplication
	nilthis.Palette_1(className)
}

// /usr/include/qt/QtWidgets/qapplication.h:113
// index:0
// Public static
// void setPalette(const class QPalette &, const char *)
func (this *QApplication) SetPalette(arg0 *qtgui.QPalette, className string) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10setPaletteERK8QPalettePKc", ffiqt.FFI_TYPE_POINTER, arg0, className)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetPalette(arg0 *qtgui.QPalette, className string) {
	var nilthis *QApplication
	nilthis.SetPalette(arg0, className)
}

// /usr/include/qt/QtWidgets/qapplication.h:114
// index:0
// Public static
// QFont font()
func (this *QApplication) Font() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4fontEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Font() {
	var nilthis *QApplication
	nilthis.Font()
}

// /usr/include/qt/QtWidgets/qapplication.h:115
// index:1
// Public static
// QFont font(const class QWidget *)
func (this *QApplication) Font_1(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4fontEPK7QWidget", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Font_1(arg0 unsafe.Pointer) {
	var nilthis *QApplication
	nilthis.Font_1(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:116
// index:2
// Public static
// QFont font(const char *)
func (this *QApplication) Font_2(className string) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4fontEPKc", ffiqt.FFI_TYPE_POINTER, className)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Font_2(className string) {
	var nilthis *QApplication
	nilthis.Font_2(className)
}

// /usr/include/qt/QtWidgets/qapplication.h:117
// index:0
// Public static
// void setFont(const class QFont &, const char *)
func (this *QApplication) SetFont(arg0 *qtgui.QFont, className string) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7setFontERK5QFontPKc", ffiqt.FFI_TYPE_POINTER, arg0, className)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetFont(arg0 *qtgui.QFont, className string) {
	var nilthis *QApplication
	nilthis.SetFont(arg0, className)
}

// /usr/include/qt/QtWidgets/qapplication.h:118
// index:0
// Public static
// QFontMetrics fontMetrics()
func (this *QApplication) FontMetrics() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication11fontMetricsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_FontMetrics() {
	var nilthis *QApplication
	nilthis.FontMetrics()
}

// /usr/include/qt/QtWidgets/qapplication.h:121
// index:0
// Public static
// void setWindowIcon(const class QIcon &)
func (this *QApplication) SetWindowIcon(icon *qtgui.QIcon) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13setWindowIconERK5QIcon", ffiqt.FFI_TYPE_POINTER, icon)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetWindowIcon(icon *qtgui.QIcon) {
	var nilthis *QApplication
	nilthis.SetWindowIcon(icon)
}

// /usr/include/qt/QtWidgets/qapplication.h:122
// index:0
// Public static
// QIcon windowIcon()
func (this *QApplication) WindowIcon() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10windowIconEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_WindowIcon() {
	var nilthis *QApplication
	nilthis.WindowIcon()
}

// /usr/include/qt/QtWidgets/qapplication.h:125
// index:0
// Public static
// QWidgetList allWidgets()
func (this *QApplication) AllWidgets() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10allWidgetsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_AllWidgets() {
	var nilthis *QApplication
	nilthis.AllWidgets()
}

// /usr/include/qt/QtWidgets/qapplication.h:126
// index:0
// Public static
// QWidgetList topLevelWidgets()
func (this *QApplication) TopLevelWidgets() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15topLevelWidgetsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_TopLevelWidgets() {
	var nilthis *QApplication
	nilthis.TopLevelWidgets()
}

// /usr/include/qt/QtWidgets/qapplication.h:128
// index:0
// Public static
// QDesktopWidget * desktop()
func (this *QApplication) Desktop() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7desktopEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Desktop() {
	var nilthis *QApplication
	nilthis.Desktop()
}

// /usr/include/qt/QtWidgets/qapplication.h:130
// index:0
// Public static
// QWidget * activePopupWidget()
func (this *QApplication) ActivePopupWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17activePopupWidgetEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_ActivePopupWidget() {
	var nilthis *QApplication
	nilthis.ActivePopupWidget()
}

// /usr/include/qt/QtWidgets/qapplication.h:131
// index:0
// Public static
// QWidget * activeModalWidget()
func (this *QApplication) ActiveModalWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17activeModalWidgetEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_ActiveModalWidget() {
	var nilthis *QApplication
	nilthis.ActiveModalWidget()
}

// /usr/include/qt/QtWidgets/qapplication.h:132
// index:0
// Public static
// QWidget * focusWidget()
func (this *QApplication) FocusWidget() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication11focusWidgetEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_FocusWidget() {
	var nilthis *QApplication
	nilthis.FocusWidget()
}

// /usr/include/qt/QtWidgets/qapplication.h:134
// index:0
// Public static
// QWidget * activeWindow()
func (this *QApplication) ActiveWindow() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication12activeWindowEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_ActiveWindow() {
	var nilthis *QApplication
	nilthis.ActiveWindow()
}

// /usr/include/qt/QtWidgets/qapplication.h:135
// index:0
// Public static
// void setActiveWindow(class QWidget *)
func (this *QApplication) SetActiveWindow(act unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15setActiveWindowEP7QWidget", ffiqt.FFI_TYPE_POINTER, act)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetActiveWindow(act unsafe.Pointer) {
	var nilthis *QApplication
	nilthis.SetActiveWindow(act)
}

// /usr/include/qt/QtWidgets/qapplication.h:137
// index:0
// Public static
// QWidget * widgetAt(const class QPoint &)
func (this *QApplication) WidgetAt(p *qtcore.QPoint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8widgetAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, p)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_WidgetAt(p *qtcore.QPoint) {
	var nilthis *QApplication
	nilthis.WidgetAt(p)
}

// /usr/include/qt/QtWidgets/qapplication.h:138
// index:1
// Public static inline
// QWidget * widgetAt(int, int)
func (this *QApplication) WidgetAt_1(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication8widgetAtEii", ffiqt.FFI_TYPE_POINTER, x, y)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_WidgetAt_1(x int, y int) {
	var nilthis *QApplication
	nilthis.WidgetAt_1(x, y)
}

// /usr/include/qt/QtWidgets/qapplication.h:139
// index:0
// Public static
// QWidget * topLevelAt(const class QPoint &)
func (this *QApplication) TopLevelAt(p *qtcore.QPoint) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10topLevelAtERK6QPoint", ffiqt.FFI_TYPE_POINTER, p)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_TopLevelAt(p *qtcore.QPoint) {
	var nilthis *QApplication
	nilthis.TopLevelAt(p)
}

// /usr/include/qt/QtWidgets/qapplication.h:140
// index:1
// Public static inline
// QWidget * topLevelAt(int, int)
func (this *QApplication) TopLevelAt_1(x int, y int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication10topLevelAtEii", ffiqt.FFI_TYPE_POINTER, x, y)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_TopLevelAt_1(x int, y int) {
	var nilthis *QApplication
	nilthis.TopLevelAt_1(x, y)
}

// /usr/include/qt/QtWidgets/qapplication.h:145
// index:0
// Public static
// void beep()
func (this *QApplication) Beep() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4beepEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QApplication_Beep() {
	var nilthis *QApplication
	nilthis.Beep()
}

// /usr/include/qt/QtWidgets/qapplication.h:146
// index:0
// Public static
// void alert(class QWidget *, int)
func (this *QApplication) Alert(widget unsafe.Pointer, duration int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication5alertEP7QWidgeti", ffiqt.FFI_TYPE_POINTER, widget, duration)
	gopp.ErrPrint(err, rv)
}
func QApplication_Alert(widget unsafe.Pointer, duration int) {
	var nilthis *QApplication
	nilthis.Alert(widget, duration)
}

// /usr/include/qt/QtWidgets/qapplication.h:148
// index:0
// Public static
// void setCursorFlashTime(int)
func (this *QApplication) SetCursorFlashTime(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication18setCursorFlashTimeEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetCursorFlashTime(arg0 int) {
	var nilthis *QApplication
	nilthis.SetCursorFlashTime(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:149
// index:0
// Public static
// int cursorFlashTime()
func (this *QApplication) CursorFlashTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15cursorFlashTimeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_CursorFlashTime() {
	var nilthis *QApplication
	nilthis.CursorFlashTime()
}

// /usr/include/qt/QtWidgets/qapplication.h:151
// index:0
// Public static
// void setDoubleClickInterval(int)
func (this *QApplication) SetDoubleClickInterval(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication22setDoubleClickIntervalEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetDoubleClickInterval(arg0 int) {
	var nilthis *QApplication
	nilthis.SetDoubleClickInterval(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:152
// index:0
// Public static
// int doubleClickInterval()
func (this *QApplication) DoubleClickInterval() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication19doubleClickIntervalEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_DoubleClickInterval() {
	var nilthis *QApplication
	nilthis.DoubleClickInterval()
}

// /usr/include/qt/QtWidgets/qapplication.h:154
// index:0
// Public static
// void setKeyboardInputInterval(int)
func (this *QApplication) SetKeyboardInputInterval(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication24setKeyboardInputIntervalEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetKeyboardInputInterval(arg0 int) {
	var nilthis *QApplication
	nilthis.SetKeyboardInputInterval(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:155
// index:0
// Public static
// int keyboardInputInterval()
func (this *QApplication) KeyboardInputInterval() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication21keyboardInputIntervalEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_KeyboardInputInterval() {
	var nilthis *QApplication
	nilthis.KeyboardInputInterval()
}

// /usr/include/qt/QtWidgets/qapplication.h:158
// index:0
// Public static
// void setWheelScrollLines(int)
func (this *QApplication) SetWheelScrollLines(arg0 int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication19setWheelScrollLinesEi", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetWheelScrollLines(arg0 int) {
	var nilthis *QApplication
	nilthis.SetWheelScrollLines(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:159
// index:0
// Public static
// int wheelScrollLines()
func (this *QApplication) WheelScrollLines() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication16wheelScrollLinesEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_WheelScrollLines() {
	var nilthis *QApplication
	nilthis.WheelScrollLines()
}

// /usr/include/qt/QtWidgets/qapplication.h:161
// index:0
// Public static
// void setGlobalStrut(const class QSize &)
func (this *QApplication) SetGlobalStrut(arg0 *qtcore.QSize) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication14setGlobalStrutERK5QSize", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetGlobalStrut(arg0 *qtcore.QSize) {
	var nilthis *QApplication
	nilthis.SetGlobalStrut(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:162
// index:0
// Public static
// QSize globalStrut()
func (this *QApplication) GlobalStrut() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication11globalStrutEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_GlobalStrut() {
	var nilthis *QApplication
	nilthis.GlobalStrut()
}

// /usr/include/qt/QtWidgets/qapplication.h:164
// index:0
// Public static
// void setStartDragTime(int)
func (this *QApplication) SetStartDragTime(ms int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication16setStartDragTimeEi", ffiqt.FFI_TYPE_POINTER, ms)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetStartDragTime(ms int) {
	var nilthis *QApplication
	nilthis.SetStartDragTime(ms)
}

// /usr/include/qt/QtWidgets/qapplication.h:165
// index:0
// Public static
// int startDragTime()
func (this *QApplication) StartDragTime() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13startDragTimeEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_StartDragTime() {
	var nilthis *QApplication
	nilthis.StartDragTime()
}

// /usr/include/qt/QtWidgets/qapplication.h:166
// index:0
// Public static
// void setStartDragDistance(int)
func (this *QApplication) SetStartDragDistance(l int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication20setStartDragDistanceEi", ffiqt.FFI_TYPE_POINTER, l)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetStartDragDistance(l int) {
	var nilthis *QApplication
	nilthis.SetStartDragDistance(l)
}

// /usr/include/qt/QtWidgets/qapplication.h:167
// index:0
// Public static
// int startDragDistance()
func (this *QApplication) StartDragDistance() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17startDragDistanceEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_StartDragDistance() {
	var nilthis *QApplication
	nilthis.StartDragDistance()
}

// /usr/include/qt/QtWidgets/qapplication.h:169
// index:0
// Public static
// bool isEffectEnabled(Qt::UIEffect)
func (this *QApplication) IsEffectEnabled(arg0 int) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15isEffectEnabledEN2Qt8UIEffectE", ffiqt.FFI_TYPE_POINTER, arg0)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_IsEffectEnabled(arg0 int) {
	var nilthis *QApplication
	nilthis.IsEffectEnabled(arg0)
}

// /usr/include/qt/QtWidgets/qapplication.h:170
// index:0
// Public static
// void setEffectEnabled(Qt::UIEffect, _Bool)
func (this *QApplication) SetEffectEnabled(arg0 int, enable bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication16setEffectEnabledEN2Qt8UIEffectEb", ffiqt.FFI_TYPE_POINTER, arg0, enable)
	gopp.ErrPrint(err, rv)
}
func QApplication_SetEffectEnabled(arg0 int, enable bool) {
	var nilthis *QApplication
	nilthis.SetEffectEnabled(arg0, enable)
}

// /usr/include/qt/QtWidgets/qapplication.h:179
// index:0
// Public static
// int exec()
func (this *QApplication) Exec() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication4execEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	return rv
}
func QApplication_Exec() {
	var nilthis *QApplication
	nilthis.Exec()
}

// /usr/include/qt/QtWidgets/qapplication.h:180
// index:0
// Public virtual
// bool notify(class QObject *, class QEvent *)
func (this *QApplication) Notify(arg0 unsafe.Pointer, arg1 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication6notifyEP7QObjectP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, arg1)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:190
// index:0
// Public
// void focusChanged(class QWidget *, class QWidget *)
func (this *QApplication) FocusChanged(old unsafe.Pointer, now unsafe.Pointer) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication12focusChangedEP7QWidgetS1_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), old, now)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:193
// index:0
// Public
// QString styleSheet()
func (this *QApplication) StyleSheet() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QApplication10styleSheetEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:196
// index:0
// Public
// void setStyleSheet(const class QString &)
func (this *QApplication) SetStyleSheet(sheet *qtcore.QString) {
	var convArg0 = sheet.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13setStyleSheetERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:198
// index:0
// Public
// void setAutoSipEnabled(const _Bool)
func (this *QApplication) SetAutoSipEnabled(enabled bool) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication17setAutoSipEnabledEb", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &enabled)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtWidgets/qapplication.h:199
// index:0
// Public
// bool autoSipEnabled()
func (this *QApplication) AutoSipEnabled() interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK12QApplication14autoSipEnabledEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:200
// index:0
// Public static
// void closeAllWindows()
func (this *QApplication) CloseAllWindows() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication15closeAllWindowsEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QApplication_CloseAllWindows() {
	var nilthis *QApplication
	nilthis.CloseAllWindows()
}

// /usr/include/qt/QtWidgets/qapplication.h:201
// index:0
// Public static
// void aboutQt()
func (this *QApplication) AboutQt() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication7aboutQtEv", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
}
func QApplication_AboutQt() {
	var nilthis *QApplication
	nilthis.AboutQt()
}

// /usr/include/qt/QtWidgets/qapplication.h:204
// index:0
// Protected virtual
// bool event(class QEvent *)
func (this *QApplication) Event(arg0 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication5eventEP6QEvent", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0)
	gopp.ErrPrint(err, rv)
	return rv
}

// /usr/include/qt/QtWidgets/qapplication.h:205
// index:0
// Protected virtual
// bool compressEvent(class QEvent *, class QObject *, class QPostEventList *)
func (this *QApplication) CompressEvent(arg0 unsafe.Pointer, receiver unsafe.Pointer, arg2 unsafe.Pointer) interface{} {
	rv, err := ffiqt.InvokeQtFunc6("_ZN12QApplication13compressEventEP6QEventP7QObjectP14QPostEventList", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), arg0, receiver, arg2)
	gopp.ErrPrint(err, rv)
	return rv
}

//  body block end
