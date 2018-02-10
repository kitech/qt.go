package qtcore

// /usr/include/qt/QtCore/qcoreevent.h
// #include <qcoreevent.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin

type QEvent struct {
	*qtrt.CObject
}

func (this *QEvent) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QEvent) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQEventFromPointer(cthis unsafe.Pointer) *QEvent {
	return &QEvent{&qtrt.CObject{cthis}}
}
func (*QEvent) NewFromPointer(cthis unsafe.Pointer) *QEvent {
	return NewQEventFromPointer(cthis)
}

// /usr/include/qt/QtCore/qcoreevent.h:297
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QEvent(enum QEvent::Type)
func NewQEvent(type_ int) *QEvent {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QEventC2ENS_4TypeE", qtrt.FFI_TYPE_POINTER, type_)
	gopp.ErrPrint(err, rv)
	gothis := NewQEventFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQEvent)
	return gothis
}

// /usr/include/qt/QtCore/qcoreevent.h:299
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QEvent()
func DeleteQEvent(this *QEvent) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QEventD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qcoreevent.h:301
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QEvent::Type type()
func (this *QEvent) Type() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QEvent4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:302
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool spontaneous()
func (this *QEvent) Spontaneous() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QEvent11spontaneousEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreevent.h:304
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setAccepted(_Bool)
func (this *QEvent) SetAccepted(accepted bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QEvent11setAcceptedEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), accepted)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:305
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isAccepted()
func (this *QEvent) IsAccepted() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK6QEvent10isAcceptedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qcoreevent.h:307
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void accept()
func (this *QEvent) Accept() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QEvent6acceptEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:308
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void ignore()
func (this *QEvent) Ignore() {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QEvent6ignoreEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreevent.h:310
// index:0
// Public static Visibility=Default Availability=Available
// [4] int registerEventType(int)
func (this *QEvent) RegisterEventType(hint int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZN6QEvent17registerEventTypeEi", qtrt.FFI_TYPE_POINTER, hint)
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}
func QEvent_RegisterEventType(hint int) int {
	var nilthis *QEvent
	rv := nilthis.RegisterEventType(hint)
	return rv
}

type QEvent__Type = int

const QEvent__None QEvent__Type = 0
const QEvent__Timer QEvent__Type = 1
const QEvent__MouseButtonPress QEvent__Type = 2
const QEvent__MouseButtonRelease QEvent__Type = 3
const QEvent__MouseButtonDblClick QEvent__Type = 4
const QEvent__MouseMove QEvent__Type = 5
const QEvent__KeyPress QEvent__Type = 6
const QEvent__KeyRelease QEvent__Type = 7
const QEvent__FocusIn QEvent__Type = 8
const QEvent__FocusOut QEvent__Type = 9
const QEvent__FocusAboutToChange QEvent__Type = 23
const QEvent__Enter QEvent__Type = 10
const QEvent__Leave QEvent__Type = 11
const QEvent__Paint QEvent__Type = 12
const QEvent__Move QEvent__Type = 13
const QEvent__Resize QEvent__Type = 14
const QEvent__Create QEvent__Type = 15
const QEvent__Destroy QEvent__Type = 16
const QEvent__Show QEvent__Type = 17
const QEvent__Hide QEvent__Type = 18
const QEvent__Close QEvent__Type = 19
const QEvent__Quit QEvent__Type = 20
const QEvent__ParentChange QEvent__Type = 21
const QEvent__ParentAboutToChange QEvent__Type = 131
const QEvent__ThreadChange QEvent__Type = 22
const QEvent__WindowActivate QEvent__Type = 24
const QEvent__WindowDeactivate QEvent__Type = 25
const QEvent__ShowToParent QEvent__Type = 26
const QEvent__HideToParent QEvent__Type = 27
const QEvent__Wheel QEvent__Type = 31
const QEvent__WindowTitleChange QEvent__Type = 33
const QEvent__WindowIconChange QEvent__Type = 34
const QEvent__ApplicationWindowIconChange QEvent__Type = 35
const QEvent__ApplicationFontChange QEvent__Type = 36
const QEvent__ApplicationLayoutDirectionChange QEvent__Type = 37
const QEvent__ApplicationPaletteChange QEvent__Type = 38
const QEvent__PaletteChange QEvent__Type = 39
const QEvent__Clipboard QEvent__Type = 40
const QEvent__Speech QEvent__Type = 42
const QEvent__MetaCall QEvent__Type = 43
const QEvent__SockAct QEvent__Type = 50
const QEvent__WinEventAct QEvent__Type = 132
const QEvent__DeferredDelete QEvent__Type = 52
const QEvent__DragEnter QEvent__Type = 60
const QEvent__DragMove QEvent__Type = 61
const QEvent__DragLeave QEvent__Type = 62
const QEvent__Drop QEvent__Type = 63
const QEvent__DragResponse QEvent__Type = 64
const QEvent__ChildAdded QEvent__Type = 68
const QEvent__ChildPolished QEvent__Type = 69
const QEvent__ChildRemoved QEvent__Type = 71
const QEvent__ShowWindowRequest QEvent__Type = 73
const QEvent__PolishRequest QEvent__Type = 74
const QEvent__Polish QEvent__Type = 75
const QEvent__LayoutRequest QEvent__Type = 76
const QEvent__UpdateRequest QEvent__Type = 77
const QEvent__UpdateLater QEvent__Type = 78
const QEvent__EmbeddingControl QEvent__Type = 79
const QEvent__ActivateControl QEvent__Type = 80
const QEvent__DeactivateControl QEvent__Type = 81
const QEvent__ContextMenu QEvent__Type = 82
const QEvent__InputMethod QEvent__Type = 83
const QEvent__TabletMove QEvent__Type = 87
const QEvent__LocaleChange QEvent__Type = 88
const QEvent__LanguageChange QEvent__Type = 89
const QEvent__LayoutDirectionChange QEvent__Type = 90
const QEvent__Style QEvent__Type = 91
const QEvent__TabletPress QEvent__Type = 92
const QEvent__TabletRelease QEvent__Type = 93
const QEvent__OkRequest QEvent__Type = 94
const QEvent__HelpRequest QEvent__Type = 95
const QEvent__IconDrag QEvent__Type = 96
const QEvent__FontChange QEvent__Type = 97
const QEvent__EnabledChange QEvent__Type = 98
const QEvent__ActivationChange QEvent__Type = 99
const QEvent__StyleChange QEvent__Type = 100
const QEvent__IconTextChange QEvent__Type = 101
const QEvent__ModifiedChange QEvent__Type = 102
const QEvent__MouseTrackingChange QEvent__Type = 109
const QEvent__WindowBlocked QEvent__Type = 103
const QEvent__WindowUnblocked QEvent__Type = 104
const QEvent__WindowStateChange QEvent__Type = 105
const QEvent__ReadOnlyChange QEvent__Type = 106
const QEvent__ToolTip QEvent__Type = 110
const QEvent__WhatsThis QEvent__Type = 111
const QEvent__StatusTip QEvent__Type = 112
const QEvent__ActionChanged QEvent__Type = 113
const QEvent__ActionAdded QEvent__Type = 114
const QEvent__ActionRemoved QEvent__Type = 115
const QEvent__FileOpen QEvent__Type = 116
const QEvent__Shortcut QEvent__Type = 117
const QEvent__ShortcutOverride QEvent__Type = 51
const QEvent__WhatsThisClicked QEvent__Type = 118
const QEvent__ToolBarChange QEvent__Type = 120
const QEvent__ApplicationActivate QEvent__Type = 121
const QEvent__ApplicationActivated QEvent__Type = 121
const QEvent__ApplicationDeactivate QEvent__Type = 122
const QEvent__ApplicationDeactivated QEvent__Type = 122
const QEvent__QueryWhatsThis QEvent__Type = 123
const QEvent__EnterWhatsThisMode QEvent__Type = 124
const QEvent__LeaveWhatsThisMode QEvent__Type = 125
const QEvent__ZOrderChange QEvent__Type = 126
const QEvent__HoverEnter QEvent__Type = 127
const QEvent__HoverLeave QEvent__Type = 128
const QEvent__HoverMove QEvent__Type = 129
const QEvent__AcceptDropsChange QEvent__Type = 152
const QEvent__ZeroTimerEvent QEvent__Type = 154
const QEvent__GraphicsSceneMouseMove QEvent__Type = 155
const QEvent__GraphicsSceneMousePress QEvent__Type = 156
const QEvent__GraphicsSceneMouseRelease QEvent__Type = 157
const QEvent__GraphicsSceneMouseDoubleClick QEvent__Type = 158
const QEvent__GraphicsSceneContextMenu QEvent__Type = 159
const QEvent__GraphicsSceneHoverEnter QEvent__Type = 160
const QEvent__GraphicsSceneHoverMove QEvent__Type = 161
const QEvent__GraphicsSceneHoverLeave QEvent__Type = 162
const QEvent__GraphicsSceneHelp QEvent__Type = 163
const QEvent__GraphicsSceneDragEnter QEvent__Type = 164
const QEvent__GraphicsSceneDragMove QEvent__Type = 165
const QEvent__GraphicsSceneDragLeave QEvent__Type = 166
const QEvent__GraphicsSceneDrop QEvent__Type = 167
const QEvent__GraphicsSceneWheel QEvent__Type = 168
const QEvent__KeyboardLayoutChange QEvent__Type = 169
const QEvent__DynamicPropertyChange QEvent__Type = 170
const QEvent__TabletEnterProximity QEvent__Type = 171
const QEvent__TabletLeaveProximity QEvent__Type = 172
const QEvent__NonClientAreaMouseMove QEvent__Type = 173
const QEvent__NonClientAreaMouseButtonPress QEvent__Type = 174
const QEvent__NonClientAreaMouseButtonRelease QEvent__Type = 175
const QEvent__NonClientAreaMouseButtonDblClick QEvent__Type = 176
const QEvent__MacSizeChange QEvent__Type = 177
const QEvent__ContentsRectChange QEvent__Type = 178
const QEvent__MacGLWindowChange QEvent__Type = 179
const QEvent__FutureCallOut QEvent__Type = 180
const QEvent__GraphicsSceneResize QEvent__Type = 181
const QEvent__GraphicsSceneMove QEvent__Type = 182
const QEvent__CursorChange QEvent__Type = 183
const QEvent__ToolTipChange QEvent__Type = 184
const QEvent__NetworkReplyUpdated QEvent__Type = 185
const QEvent__GrabMouse QEvent__Type = 186
const QEvent__UngrabMouse QEvent__Type = 187
const QEvent__GrabKeyboard QEvent__Type = 188
const QEvent__UngrabKeyboard QEvent__Type = 189
const QEvent__MacGLClearDrawable QEvent__Type = 191
const QEvent__StateMachineSignal QEvent__Type = 192
const QEvent__StateMachineWrapped QEvent__Type = 193
const QEvent__TouchBegin QEvent__Type = 194
const QEvent__TouchUpdate QEvent__Type = 195
const QEvent__TouchEnd QEvent__Type = 196
const QEvent__NativeGesture QEvent__Type = 197
const QEvent__RequestSoftwareInputPanel QEvent__Type = 199
const QEvent__CloseSoftwareInputPanel QEvent__Type = 200
const QEvent__WinIdChange QEvent__Type = 203
const QEvent__Gesture QEvent__Type = 198
const QEvent__GestureOverride QEvent__Type = 202
const QEvent__ScrollPrepare QEvent__Type = 204
const QEvent__Scroll QEvent__Type = 205
const QEvent__Expose QEvent__Type = 206
const QEvent__InputMethodQuery QEvent__Type = 207
const QEvent__OrientationChange QEvent__Type = 208
const QEvent__TouchCancel QEvent__Type = 209
const QEvent__ThemeChange QEvent__Type = 210
const QEvent__SockClose QEvent__Type = 211
const QEvent__PlatformPanel QEvent__Type = 212
const QEvent__StyleAnimationUpdate QEvent__Type = 213
const QEvent__ApplicationStateChange QEvent__Type = 214
const QEvent__WindowChangeInternal QEvent__Type = 215
const QEvent__ScreenChangeInternal QEvent__Type = 216
const QEvent__PlatformSurface QEvent__Type = 217
const QEvent__Pointer QEvent__Type = 218
const QEvent__TabletTrackingChange QEvent__Type = 219
const QEvent__User QEvent__Type = 1000
const QEvent__MaxUser QEvent__Type = 65535

//  body block end
