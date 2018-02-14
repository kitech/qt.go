package qtgui

// /usr/include/qt/QtGui/qstylehints.h
// #include <qstylehints.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 20
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QStyleHints struct {
	*qtcore.QObject
}
type QStyleHints_ITF interface {
	qtcore.QObject_ITF
	QStyleHints_PTR() *QStyleHints
}

func (ptr *QStyleHints) QStyleHints_PTR() *QStyleHints { return ptr }

func (this *QStyleHints) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QStyleHints) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQStyleHintsFromPointer(cthis unsafe.Pointer) *QStyleHints {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QStyleHints{bcthis0}
}
func (*QStyleHints) NewFromPointer(cthis unsafe.Pointer) *QStyleHints {
	return NewQStyleHintsFromPointer(cthis)
}

// /usr/include/qt/QtGui/qstylehints.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QStyleHints) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qstylehints.h:78
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMouseDoubleClickInterval(int)
func (this *QStyleHints) SetMouseDoubleClickInterval(mouseDoubleClickInterval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints27setMouseDoubleClickIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mouseDoubleClickInterval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:79
// index:0
// Public Visibility=Default Availability=Available
// [4] int mouseDoubleClickInterval()
func (this *QStyleHints) MouseDoubleClickInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints24mouseDoubleClickIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:80
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setMousePressAndHoldInterval(int)
func (this *QStyleHints) SetMousePressAndHoldInterval(mousePressAndHoldInterval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints28setMousePressAndHoldIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mousePressAndHoldInterval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:81
// index:0
// Public Visibility=Default Availability=Available
// [4] int mousePressAndHoldInterval()
func (this *QStyleHints) MousePressAndHoldInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints25mousePressAndHoldIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:82
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartDragDistance(int)
func (this *QStyleHints) SetStartDragDistance(startDragDistance int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints20setStartDragDistanceEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startDragDistance)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:83
// index:0
// Public Visibility=Default Availability=Available
// [4] int startDragDistance()
func (this *QStyleHints) StartDragDistance() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints17startDragDistanceEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:84
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setStartDragTime(int)
func (this *QStyleHints) SetStartDragTime(startDragTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints16setStartDragTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startDragTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:85
// index:0
// Public Visibility=Default Availability=Available
// [4] int startDragTime()
func (this *QStyleHints) StartDragTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints13startDragTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:86
// index:0
// Public Visibility=Default Availability=Available
// [4] int startDragVelocity()
func (this *QStyleHints) StartDragVelocity() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints17startDragVelocityEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:87
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setKeyboardInputInterval(int)
func (this *QStyleHints) SetKeyboardInputInterval(keyboardInputInterval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints24setKeyboardInputIntervalEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), keyboardInputInterval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:88
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyboardInputInterval()
func (this *QStyleHints) KeyboardInputInterval() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints21keyboardInputIntervalEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:89
// index:0
// Public Visibility=Default Availability=Available
// [4] int keyboardAutoRepeatRate()
func (this *QStyleHints) KeyboardAutoRepeatRate() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints22keyboardAutoRepeatRateEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:90
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCursorFlashTime(int)
func (this *QStyleHints) SetCursorFlashTime(cursorFlashTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints18setCursorFlashTimeEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorFlashTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:91
// index:0
// Public Visibility=Default Availability=Available
// [4] int cursorFlashTime()
func (this *QStyleHints) CursorFlashTime() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints15cursorFlashTimeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:92
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showIsFullScreen()
func (this *QStyleHints) ShowIsFullScreen() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints16showIsFullScreenEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:93
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showIsMaximized()
func (this *QStyleHints) ShowIsMaximized() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints15showIsMaximizedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:94
// index:0
// Public Visibility=Default Availability=Available
// [1] bool showShortcutsInContextMenus()
func (this *QStyleHints) ShowShortcutsInContextMenus() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints27showShortcutsInContextMenusEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:95
// index:0
// Public Visibility=Default Availability=Available
// [4] int passwordMaskDelay()
func (this *QStyleHints) PasswordMaskDelay() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints17passwordMaskDelayEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:96
// index:0
// Public Visibility=Default Availability=Available
// [2] QChar passwordMaskCharacter()
func (this *QStyleHints) PasswordMaskCharacter() *qtcore.QChar /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints21passwordMaskCharacterEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQCharFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQChar)
	return rv2
}

// /usr/include/qt/QtGui/qstylehints.h:97
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal fontSmoothingGamma()
func (this *QStyleHints) FontSmoothingGamma() float64 {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints18fontSmoothingGammaEv", qtrt.FFI_TYPE_DOUBLE, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:98
// index:0
// Public Visibility=Default Availability=Available
// [1] bool useRtlExtensions()
func (this *QStyleHints) UseRtlExtensions() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints16useRtlExtensionsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:99
// index:0
// Public Visibility=Default Availability=Available
// [1] bool setFocusOnTouchRelease()
func (this *QStyleHints) SetFocusOnTouchRelease() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints22setFocusOnTouchReleaseEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:100
// index:0
// Public Visibility=Default Availability=Available
// [4] Qt::TabFocusBehavior tabFocusBehavior()
func (this *QStyleHints) TabFocusBehavior() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints16tabFocusBehaviorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtGui/qstylehints.h:101
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setTabFocusBehavior(Qt::TabFocusBehavior)
func (this *QStyleHints) SetTabFocusBehavior(tabFocusBehavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints19setTabFocusBehaviorEN2Qt16TabFocusBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabFocusBehavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:102
// index:0
// Public Visibility=Default Availability=Available
// [1] bool singleClickActivation()
func (this *QStyleHints) SingleClickActivation() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints21singleClickActivationEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:103
// index:0
// Public Visibility=Default Availability=Available
// [1] bool useHoverEffects()
func (this *QStyleHints) UseHoverEffects() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints15useHoverEffectsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qstylehints.h:104
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUseHoverEffects(_Bool)
func (this *QStyleHints) SetUseHoverEffects(useHoverEffects bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints18setUseHoverEffectsEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), useHoverEffects)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:105
// index:0
// Public Visibility=Default Availability=Available
// [4] int wheelScrollLines()
func (this *QStyleHints) WheelScrollLines() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QStyleHints16wheelScrollLinesEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qstylehints.h:106
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setWheelScrollLines(int)
func (this *QStyleHints) SetWheelScrollLines(scrollLines int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints19setWheelScrollLinesEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scrollLines)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void cursorFlashTimeChanged(int)
func (this *QStyleHints) CursorFlashTimeChanged(cursorFlashTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints22cursorFlashTimeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), cursorFlashTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:110
// index:0
// Public Visibility=Default Availability=Available
// [-2] void keyboardInputIntervalChanged(int)
func (this *QStyleHints) KeyboardInputIntervalChanged(keyboardInputInterval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints28keyboardInputIntervalChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), keyboardInputInterval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:111
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mouseDoubleClickIntervalChanged(int)
func (this *QStyleHints) MouseDoubleClickIntervalChanged(mouseDoubleClickInterval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints31mouseDoubleClickIntervalChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mouseDoubleClickInterval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:112
// index:0
// Public Visibility=Default Availability=Available
// [-2] void mousePressAndHoldIntervalChanged(int)
func (this *QStyleHints) MousePressAndHoldIntervalChanged(mousePressAndHoldInterval int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints32mousePressAndHoldIntervalChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), mousePressAndHoldInterval)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startDragDistanceChanged(int)
func (this *QStyleHints) StartDragDistanceChanged(startDragDistance int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints24startDragDistanceChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startDragDistance)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:114
// index:0
// Public Visibility=Default Availability=Available
// [-2] void startDragTimeChanged(int)
func (this *QStyleHints) StartDragTimeChanged(startDragTime int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints20startDragTimeChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), startDragTime)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:115
// index:0
// Public Visibility=Default Availability=Available
// [-2] void tabFocusBehaviorChanged(Qt::TabFocusBehavior)
func (this *QStyleHints) TabFocusBehaviorChanged(tabFocusBehavior int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints23tabFocusBehaviorChangedEN2Qt16TabFocusBehaviorE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), tabFocusBehavior)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void useHoverEffectsChanged(_Bool)
func (this *QStyleHints) UseHoverEffectsChanged(useHoverEffects bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints22useHoverEffectsChangedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), useHoverEffects)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstylehints.h:117
// index:0
// Public Visibility=Default Availability=Available
// [-2] void wheelScrollLinesChanged(int)
func (this *QStyleHints) WheelScrollLinesChanged(scrollLines int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHints23wheelScrollLinesChangedEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), scrollLines)
	qtrt.ErrPrint(err, rv)
}

func DeleteQStyleHints(this *QStyleHints) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QStyleHintsD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
}

//  keep block end
