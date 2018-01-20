//  header block begin
// /usr/include/qt/QtGui/qstylehints.h
// #include <qstylehints.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QStyleHints struct {
	*qtcore.QObject
}

func (this *QStyleHints) GetCthis() unsafe.Pointer {
	return this.QObject.GetCthis()
}

// /usr/include/qt/QtGui/qstylehints.h:54
// index:0
// virtual
// const QMetaObject * metaObject()
func (this *QStyleHints) MetaObject() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints10metaObjectEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:78
// index:0
// void setMouseDoubleClickInterval(int)
func (this *QStyleHints) SetMouseDoubleClickInterval(mouseDoubleClickInterval int) {
	// 0: (, mouseDoubleClickInterval int), (&mouseDoubleClickInterval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints27setMouseDoubleClickIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mouseDoubleClickInterval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:79
// index:0
// int mouseDoubleClickInterval()
func (this *QStyleHints) MouseDoubleClickInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints24mouseDoubleClickIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:80
// index:0
// void setMousePressAndHoldInterval(int)
func (this *QStyleHints) SetMousePressAndHoldInterval(mousePressAndHoldInterval int) {
	// 0: (, mousePressAndHoldInterval int), (&mousePressAndHoldInterval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints28setMousePressAndHoldIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mousePressAndHoldInterval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:81
// index:0
// int mousePressAndHoldInterval()
func (this *QStyleHints) MousePressAndHoldInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints25mousePressAndHoldIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:82
// index:0
// void setStartDragDistance(int)
func (this *QStyleHints) SetStartDragDistance(startDragDistance int) {
	// 0: (, startDragDistance int), (&startDragDistance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints20setStartDragDistanceEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startDragDistance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:83
// index:0
// int startDragDistance()
func (this *QStyleHints) StartDragDistance() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints17startDragDistanceEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:84
// index:0
// void setStartDragTime(int)
func (this *QStyleHints) SetStartDragTime(startDragTime int) {
	// 0: (, startDragTime int), (&startDragTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints16setStartDragTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startDragTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:85
// index:0
// int startDragTime()
func (this *QStyleHints) StartDragTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints13startDragTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:86
// index:0
// int startDragVelocity()
func (this *QStyleHints) StartDragVelocity() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints17startDragVelocityEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:87
// index:0
// void setKeyboardInputInterval(int)
func (this *QStyleHints) SetKeyboardInputInterval(keyboardInputInterval int) {
	// 0: (, keyboardInputInterval int), (&keyboardInputInterval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints24setKeyboardInputIntervalEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &keyboardInputInterval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:88
// index:0
// int keyboardInputInterval()
func (this *QStyleHints) KeyboardInputInterval() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints21keyboardInputIntervalEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:89
// index:0
// int keyboardAutoRepeatRate()
func (this *QStyleHints) KeyboardAutoRepeatRate() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints22keyboardAutoRepeatRateEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:90
// index:0
// void setCursorFlashTime(int)
func (this *QStyleHints) SetCursorFlashTime(cursorFlashTime int) {
	// 0: (, cursorFlashTime int), (&cursorFlashTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints18setCursorFlashTimeEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorFlashTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:91
// index:0
// int cursorFlashTime()
func (this *QStyleHints) CursorFlashTime() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints15cursorFlashTimeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:92
// index:0
// bool showIsFullScreen()
func (this *QStyleHints) ShowIsFullScreen() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints16showIsFullScreenEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:93
// index:0
// bool showIsMaximized()
func (this *QStyleHints) ShowIsMaximized() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints15showIsMaximizedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:94
// index:0
// bool showShortcutsInContextMenus()
func (this *QStyleHints) ShowShortcutsInContextMenus() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints27showShortcutsInContextMenusEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:95
// index:0
// int passwordMaskDelay()
func (this *QStyleHints) PasswordMaskDelay() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints17passwordMaskDelayEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:96
// index:0
// QChar passwordMaskCharacter()
func (this *QStyleHints) PasswordMaskCharacter() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints21passwordMaskCharacterEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:97
// index:0
// qreal fontSmoothingGamma()
func (this *QStyleHints) FontSmoothingGamma() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints18fontSmoothingGammaEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:98
// index:0
// bool useRtlExtensions()
func (this *QStyleHints) UseRtlExtensions() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints16useRtlExtensionsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:99
// index:0
// bool setFocusOnTouchRelease()
func (this *QStyleHints) SetFocusOnTouchRelease() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints22setFocusOnTouchReleaseEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:100
// index:0
// Qt::TabFocusBehavior tabFocusBehavior()
func (this *QStyleHints) TabFocusBehavior() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints16tabFocusBehaviorEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:101
// index:0
// void setTabFocusBehavior(Qt::TabFocusBehavior)
func (this *QStyleHints) SetTabFocusBehavior(tabFocusBehavior int) {
	// 0: (, tabFocusBehavior Qt::TabFocusBehavior), (&tabFocusBehavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints19setTabFocusBehaviorEN2Qt16TabFocusBehaviorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tabFocusBehavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:102
// index:0
// bool singleClickActivation()
func (this *QStyleHints) SingleClickActivation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints21singleClickActivationEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:103
// index:0
// bool useHoverEffects()
func (this *QStyleHints) UseHoverEffects() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints15useHoverEffectsEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:104
// index:0
// void setUseHoverEffects(_Bool)
func (this *QStyleHints) SetUseHoverEffects(useHoverEffects bool) {
	// 0: (, useHoverEffects bool), (&useHoverEffects)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints18setUseHoverEffectsEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &useHoverEffects)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:105
// index:0
// int wheelScrollLines()
func (this *QStyleHints) WheelScrollLines() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK11QStyleHints16wheelScrollLinesEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:106
// index:0
// void setWheelScrollLines(int)
func (this *QStyleHints) SetWheelScrollLines(scrollLines int) {
	// 0: (, scrollLines int), (&scrollLines)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints19setWheelScrollLinesEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &scrollLines)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:109
// index:0
// void cursorFlashTimeChanged(int)
func (this *QStyleHints) CursorFlashTimeChanged(cursorFlashTime int) {
	// 0: (, cursorFlashTime int), (&cursorFlashTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints22cursorFlashTimeChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &cursorFlashTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:110
// index:0
// void keyboardInputIntervalChanged(int)
func (this *QStyleHints) KeyboardInputIntervalChanged(keyboardInputInterval int) {
	// 0: (, keyboardInputInterval int), (&keyboardInputInterval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints28keyboardInputIntervalChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &keyboardInputInterval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:111
// index:0
// void mouseDoubleClickIntervalChanged(int)
func (this *QStyleHints) MouseDoubleClickIntervalChanged(mouseDoubleClickInterval int) {
	// 0: (, mouseDoubleClickInterval int), (&mouseDoubleClickInterval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints31mouseDoubleClickIntervalChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mouseDoubleClickInterval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:112
// index:0
// void mousePressAndHoldIntervalChanged(int)
func (this *QStyleHints) MousePressAndHoldIntervalChanged(mousePressAndHoldInterval int) {
	// 0: (, mousePressAndHoldInterval int), (&mousePressAndHoldInterval)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints32mousePressAndHoldIntervalChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &mousePressAndHoldInterval)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:113
// index:0
// void startDragDistanceChanged(int)
func (this *QStyleHints) StartDragDistanceChanged(startDragDistance int) {
	// 0: (, startDragDistance int), (&startDragDistance)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints24startDragDistanceChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startDragDistance)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:114
// index:0
// void startDragTimeChanged(int)
func (this *QStyleHints) StartDragTimeChanged(startDragTime int) {
	// 0: (, startDragTime int), (&startDragTime)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints20startDragTimeChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &startDragTime)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:115
// index:0
// void tabFocusBehaviorChanged(Qt::TabFocusBehavior)
func (this *QStyleHints) TabFocusBehaviorChanged(tabFocusBehavior int) {
	// 0: (, tabFocusBehavior Qt::TabFocusBehavior), (&tabFocusBehavior)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints23tabFocusBehaviorChangedEN2Qt16TabFocusBehaviorE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &tabFocusBehavior)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:116
// index:0
// void useHoverEffectsChanged(_Bool)
func (this *QStyleHints) UseHoverEffectsChanged(useHoverEffects bool) {
	// 0: (, useHoverEffects bool), (&useHoverEffects)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints22useHoverEffectsChangedEb", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &useHoverEffects)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:117
// index:0
// void wheelScrollLinesChanged(int)
func (this *QStyleHints) WheelScrollLinesChanged(scrollLines int) {
	// 0: (, scrollLines int), (&scrollLines)
	rv, err := ffiqt.InvokeQtFunc6("_ZN11QStyleHints23wheelScrollLinesChangedEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &scrollLines)
	gopp.ErrPrint(err, rv)
}

//  body block end
