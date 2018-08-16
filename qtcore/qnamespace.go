package qtcore

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

//  ext block end

//  body block begin

/*


 */
type Qt__float_round_style = int // stdglobal
//
const Qt__round_indeterminate Qt__float_round_style = -1

//
const Qt__round_toward_zero Qt__float_round_style = 0

//
const Qt__round_to_nearest Qt__float_round_style = 1

//
const Qt__round_toward_infinity Qt__float_round_style = 2

//
const Qt__round_toward_neg_infinity Qt__float_round_style = 3

/*


 */
type Qt__float_denorm_style = int // stdglobal
//
const Qt__denorm_indeterminate Qt__float_denorm_style = -1

//
const Qt__denorm_absent Qt__float_denorm_style = 0

//
const Qt__denorm_present Qt__float_denorm_style = 1

/*


 */
type Qt__QtMsgType = int // core
//
const Qt__QtDebugMsg Qt__QtMsgType = 0

//
const Qt__QtWarningMsg Qt__QtMsgType = 1

//
const Qt__QtCriticalMsg Qt__QtMsgType = 2

//
const Qt__QtFatalMsg Qt__QtMsgType = 3

//
const Qt__QtInfoMsg Qt__QtMsgType = 4

//
const Qt__QtSystemMsg Qt__QtMsgType = 2

/*


 */
type Qt__memory_order = int // stdglobal
//
const Qt__memory_order_relaxed Qt__memory_order = 0

//
const Qt__memory_order_consume Qt__memory_order = 1

//
const Qt__memory_order_acquire Qt__memory_order = 2

//
const Qt__memory_order_release Qt__memory_order = 3

//
const Qt__memory_order_acq_rel Qt__memory_order = 4

//
const Qt__memory_order_seq_cst Qt__memory_order = 5

/*


 */
type Qt____memory_order_modifier = int // stdglobal
//
const Qt____memory_order_mask Qt____memory_order_modifier = 65535

//
const Qt____memory_order_modifier_mask Qt____memory_order_modifier = -65536

//
const Qt____memory_order_hle_acquire Qt____memory_order_modifier = 65536

//
const Qt____memory_order_hle_release Qt____memory_order_modifier = 131072

/*
Qt's predefined QColor objects:



See also QColor.

*/
type Qt__GlobalColor = int // core
//
const Qt__color0 Qt__GlobalColor = 0

//
const Qt__color1 Qt__GlobalColor = 1

//
const Qt__black Qt__GlobalColor = 2

// White (#ffffff)
const Qt__white Qt__GlobalColor = 3

//
const Qt__darkGray Qt__GlobalColor = 4

//
const Qt__gray Qt__GlobalColor = 5

//
const Qt__lightGray Qt__GlobalColor = 6

//
const Qt__red Qt__GlobalColor = 7

//
const Qt__green Qt__GlobalColor = 8

//
const Qt__blue Qt__GlobalColor = 9

//
const Qt__cyan Qt__GlobalColor = 10

//
const Qt__magenta Qt__GlobalColor = 11

//
const Qt__yellow Qt__GlobalColor = 12

//
const Qt__darkRed Qt__GlobalColor = 13

//
const Qt__darkGreen Qt__GlobalColor = 14

//
const Qt__darkBlue Qt__GlobalColor = 15

//
const Qt__darkCyan Qt__GlobalColor = 16

//
const Qt__darkMagenta Qt__GlobalColor = 17

//
const Qt__darkYellow Qt__GlobalColor = 18

//
const Qt__transparent Qt__GlobalColor = 19

/*


 */
type Qt__KeyboardModifier = int // core
//
const Qt__NoModifier Qt__KeyboardModifier = 0

//
const Qt__ShiftModifier Qt__KeyboardModifier = 33554432

//
const Qt__ControlModifier Qt__KeyboardModifier = 67108864

//
const Qt__AltModifier Qt__KeyboardModifier = 134217728

//
const Qt__MetaModifier Qt__KeyboardModifier = 268435456

//
const Qt__KeypadModifier Qt__KeyboardModifier = 536870912

//
const Qt__GroupSwitchModifier Qt__KeyboardModifier = 1073741824

//
const Qt__KeyboardModifierMask Qt__KeyboardModifier = -33554432

/*
This enum provides shorter names for the keyboard modifier keys supported by Qt.

Note: On macOS, the CTRL value corresponds to the Command keys on the keyboard, and the META value corresponds to the Control keys.

Qt::SHIFTQt::ShiftModifierThe Shift keys provided on all standard keyboards.
Qt::METAQt::MetaModifierThe Meta keys.
Qt::CTRLQt::ControlModifierThe Ctrl keys.
Qt::ALTQt::AltModifierThe normal Alt keys, but not keys like AltGr.


See also KeyboardModifier and MouseButton.

*/
type Qt__Modifier = int // core
//
const Qt__META Qt__Modifier = 268435456

//
const Qt__SHIFT Qt__Modifier = 33554432

//
const Qt__CTRL Qt__Modifier = 67108864

//
const Qt__ALT Qt__Modifier = 134217728

//
const Qt__MODIFIER_MASK Qt__Modifier = -33554432

//
const Qt__UNICODE_ACCEL Qt__Modifier = 0

/*


 */
type Qt__MouseButton = int // core
//
const Qt__NoButton Qt__MouseButton = 0

//
const Qt__LeftButton Qt__MouseButton = 1

//
const Qt__RightButton Qt__MouseButton = 2

//
const Qt__MidButton Qt__MouseButton = 4

//
const Qt__MiddleButton Qt__MouseButton = 4

//
const Qt__BackButton Qt__MouseButton = 8

//
const Qt__XButton1 Qt__MouseButton = 8

//
const Qt__ExtraButton1 Qt__MouseButton = 8

//
const Qt__ForwardButton Qt__MouseButton = 16

//
const Qt__XButton2 Qt__MouseButton = 16

//
const Qt__ExtraButton2 Qt__MouseButton = 16

//
const Qt__TaskButton Qt__MouseButton = 32

//
const Qt__ExtraButton3 Qt__MouseButton = 32

//
const Qt__ExtraButton4 Qt__MouseButton = 64

//
const Qt__ExtraButton5 Qt__MouseButton = 128

//
const Qt__ExtraButton6 Qt__MouseButton = 256

//
const Qt__ExtraButton7 Qt__MouseButton = 512

//
const Qt__ExtraButton8 Qt__MouseButton = 1024

//
const Qt__ExtraButton9 Qt__MouseButton = 2048

//
const Qt__ExtraButton10 Qt__MouseButton = 4096

//
const Qt__ExtraButton11 Qt__MouseButton = 8192

//
const Qt__ExtraButton12 Qt__MouseButton = 16384

//
const Qt__ExtraButton13 Qt__MouseButton = 32768

//
const Qt__ExtraButton14 Qt__MouseButton = 65536

//
const Qt__ExtraButton15 Qt__MouseButton = 131072

//
const Qt__ExtraButton16 Qt__MouseButton = 262144

//
const Qt__ExtraButton17 Qt__MouseButton = 524288

//
const Qt__ExtraButton18 Qt__MouseButton = 1048576

//
const Qt__ExtraButton19 Qt__MouseButton = 2097152

//
const Qt__ExtraButton20 Qt__MouseButton = 4194304

//
const Qt__ExtraButton21 Qt__MouseButton = 8388608

//
const Qt__ExtraButton22 Qt__MouseButton = 16777216

//
const Qt__ExtraButton23 Qt__MouseButton = 33554432

//
const Qt__ExtraButton24 Qt__MouseButton = 67108864

//
const Qt__AllButtons Qt__MouseButton = 134217727

//
const Qt__MaxMouseButton Qt__MouseButton = 67108864

//
const Qt__MouseButtonMask Qt__MouseButton = -1

/*


 */
type Qt__Orientation = int // core
//
const Qt__Horizontal Qt__Orientation = 1

//
const Qt__Vertical Qt__Orientation = 2

/*
This enum type defines the various policies a widget can have with respect to acquiring keyboard focus.


*/
type Qt__FocusPolicy = int // core
// the widget does not accept focus.
const Qt__NoFocus Qt__FocusPolicy = 0

//
const Qt__TabFocus Qt__FocusPolicy = 1

//
const Qt__ClickFocus Qt__FocusPolicy = 2

//
const Qt__StrongFocus Qt__FocusPolicy = 11

//
const Qt__WheelFocus Qt__FocusPolicy = 15

/*
This enum type provides different focus behaviors for tab navigation.



This enum was introduced or modified in  Qt 5.5.

*/
type Qt__TabFocusBehavior = int // core
//
const Qt__NoTabFocus Qt__TabFocusBehavior = 0

//
const Qt__TabFocusTextControls Qt__TabFocusBehavior = 1

//
const Qt__TabFocusListControls Qt__TabFocusBehavior = 2

// xffiterate all controls and widgets.
const Qt__TabFocusAllControls Qt__TabFocusBehavior = 255

/*
This enum describes how the items in a widget are sorted.


*/
type Qt__SortOrder = int // core
//
const Qt__AscendingOrder Qt__SortOrder = 0

//
const Qt__DescendingOrder Qt__SortOrder = 1

/*
This enum describes how to repeat or stretch the parts of an image when drawing.



This enum was introduced or modified in  Qt 4.6.

*/
type Qt__TileRule = int // core
// Scale the image to fit to the available area.
const Qt__StretchTile Qt__TileRule = 0

// Repeat the image until there is no more space. May crop the last image.
const Qt__RepeatTile Qt__TileRule = 1

// Similar to Repeat, but scales the image down to ensure that the last tile is not cropped.
const Qt__RoundTile Qt__TileRule = 2

/*


 */
type Qt__AlignmentFlag = int // core
//
const Qt__AlignLeft Qt__AlignmentFlag = 1

//
const Qt__AlignLeading Qt__AlignmentFlag = 1

//
const Qt__AlignRight Qt__AlignmentFlag = 2

//
const Qt__AlignTrailing Qt__AlignmentFlag = 2

//
const Qt__AlignHCenter Qt__AlignmentFlag = 4

//
const Qt__AlignJustify Qt__AlignmentFlag = 8

//
const Qt__AlignAbsolute Qt__AlignmentFlag = 16

//
const Qt__AlignHorizontal_Mask Qt__AlignmentFlag = 31

//
const Qt__AlignTop Qt__AlignmentFlag = 32

//
const Qt__AlignBottom Qt__AlignmentFlag = 64

//
const Qt__AlignVCenter Qt__AlignmentFlag = 128

//
const Qt__AlignBaseline Qt__AlignmentFlag = 256

//
const Qt__AlignVertical_Mask Qt__AlignmentFlag = 480

//
const Qt__AlignCenter Qt__AlignmentFlag = 132

/*
This enum type is used to define some modifier flags. Some of these flags only make sense in the context of printing:



You can use as many modifier flags as you want, except that Qt::TextSingleLine and Qt::TextWordWrap cannot be combined.

Flags that are inappropriate for a given use are generally ignored.

*/
type Qt__TextFlag = int // core
//
const Qt__TextSingleLine Qt__TextFlag = 256

//
const Qt__TextDontClip Qt__TextFlag = 512

//
const Qt__TextExpandTabs Qt__TextFlag = 1024

//
const Qt__TextShowMnemonic Qt__TextFlag = 2048

//
const Qt__TextWordWrap Qt__TextFlag = 4096

//
const Qt__TextWrapAnywhere Qt__TextFlag = 8192

//
const Qt__TextDontPrint Qt__TextFlag = 16384

//
const Qt__TextIncludeTrailingSpaces Qt__TextFlag = 134217728

//
const Qt__TextHideMnemonic Qt__TextFlag = 32768

//
const Qt__TextJustificationForced Qt__TextFlag = 65536

//
const Qt__TextForceLeftToRight Qt__TextFlag = 131072

//
const Qt__TextForceRightToLeft Qt__TextFlag = 262144

//
const Qt__TextLongestVariant Qt__TextFlag = 524288

//
const Qt__TextBypassShaping Qt__TextFlag = 1048576

/*
This enum specifies where the ellipsis should appear when displaying texts that don't fit:




See also QAbstractItemView::textElideMode, QFontMetrics::elidedText(), AlignmentFlag, and QTabBar::elideMode.

*/
type Qt__TextElideMode = int // core
// The ellipsis should appear at the beginning of the text.
const Qt__ElideLeft Qt__TextElideMode = 0

// The ellipsis should appear at the end of the text.
const Qt__ElideRight Qt__TextElideMode = 1

// The ellipsis should appear in the middle of the text.
const Qt__ElideMiddle Qt__TextElideMode = 2

// Ellipsis should NOT appear in the text.
const Qt__ElideNone Qt__TextElideMode = 3

/*
This enum describes the types of whitespace mode that are used by the QTextDocument class to meet the requirements of different kinds of textual information.

Qt::WhiteSpaceNoWrap2

*/
type Qt__WhiteSpaceMode = int // core
// The whitespace mode used to display normal word wrapped text in paragraphs.
const Qt__WhiteSpaceNormal Qt__WhiteSpaceMode = 0

// A preformatted text mode in which whitespace is reproduced exactly.
const Qt__WhiteSpacePre Qt__WhiteSpaceMode = 1

//
const Qt__WhiteSpaceNoWrap Qt__WhiteSpaceMode = 2

//
const Qt__WhiteSpaceModeUndefined Qt__WhiteSpaceMode = -1

/*
This enum contains the types of accuracy that can be used by the QTextDocument class when testing for mouse clicks on text documents.


*/
type Qt__HitTestAccuracy = int // core
// The point at which input occurred must coincide exactly with input-sensitive parts of the document.
const Qt__ExactHit Qt__HitTestAccuracy = 0

// The point at which input occurred can lie close to input-sensitive parts of the document.
const Qt__FuzzyHit Qt__HitTestAccuracy = 1

/*


 */
type Qt__WindowType = int // core
//
const Qt__Widget Qt__WindowType = 0

//
const Qt__Window Qt__WindowType = 1

//
const Qt__Dialog Qt__WindowType = 3

//
const Qt__Sheet Qt__WindowType = 5

//
const Qt__Drawer Qt__WindowType = 7

//
const Qt__Popup Qt__WindowType = 9

//
const Qt__Tool Qt__WindowType = 11

//
const Qt__ToolTip Qt__WindowType = 13

//
const Qt__SplashScreen Qt__WindowType = 15

//
const Qt__Desktop Qt__WindowType = 17

//
const Qt__SubWindow Qt__WindowType = 18

//
const Qt__ForeignWindow Qt__WindowType = 33

//
const Qt__CoverWindow Qt__WindowType = 65

//
const Qt__WindowType_Mask Qt__WindowType = 255

//
const Qt__MSWindowsFixedSizeDialogHint Qt__WindowType = 256

//
const Qt__MSWindowsOwnDC Qt__WindowType = 512

//
const Qt__BypassWindowManagerHint Qt__WindowType = 1024

//
const Qt__X11BypassWindowManagerHint Qt__WindowType = 1024

//
const Qt__FramelessWindowHint Qt__WindowType = 2048

//
const Qt__WindowTitleHint Qt__WindowType = 4096

//
const Qt__WindowSystemMenuHint Qt__WindowType = 8192

//
const Qt__WindowMinimizeButtonHint Qt__WindowType = 16384

//
const Qt__WindowMaximizeButtonHint Qt__WindowType = 32768

//
const Qt__WindowMinMaxButtonsHint Qt__WindowType = 49152

//
const Qt__WindowContextHelpButtonHint Qt__WindowType = 65536

//
const Qt__WindowShadeButtonHint Qt__WindowType = 131072

//
const Qt__WindowStaysOnTopHint Qt__WindowType = 262144

//
const Qt__WindowTransparentForInput Qt__WindowType = 524288

//
const Qt__WindowOverridesSystemGestures Qt__WindowType = 1048576

//
const Qt__WindowDoesNotAcceptFocus Qt__WindowType = 2097152

//
const Qt__MaximizeUsingFullscreenGeometryHint Qt__WindowType = 4194304

//
const Qt__CustomizeWindowHint Qt__WindowType = 33554432

//
const Qt__WindowStaysOnBottomHint Qt__WindowType = 67108864

//
const Qt__WindowCloseButtonHint Qt__WindowType = 134217728

//
const Qt__MacWindowToolBarButtonHint Qt__WindowType = 268435456

//
const Qt__BypassGraphicsProxyWidget Qt__WindowType = 536870912

//
const Qt__NoDropShadowWindowHint Qt__WindowType = 1073741824

//
const Qt__WindowFullscreenButtonHint Qt__WindowType = -2147483648

/*


 */
type Qt__WindowState = int // core
//
const Qt__WindowNoState Qt__WindowState = 0

//
const Qt__WindowMinimized Qt__WindowState = 1

//
const Qt__WindowMaximized Qt__WindowState = 2

//
const Qt__WindowFullScreen Qt__WindowState = 4

//
const Qt__WindowActive Qt__WindowState = 8

/*


 */
type Qt__ApplicationState = int // core
//
const Qt__ApplicationSuspended Qt__ApplicationState = 0

//
const Qt__ApplicationHidden Qt__ApplicationState = 1

//
const Qt__ApplicationInactive Qt__ApplicationState = 2

//
const Qt__ApplicationActive Qt__ApplicationState = 4

/*


 */
type Qt__ScreenOrientation = int // core
//
const Qt__PrimaryOrientation Qt__ScreenOrientation = 0

//
const Qt__PortraitOrientation Qt__ScreenOrientation = 1

//
const Qt__LandscapeOrientation Qt__ScreenOrientation = 2

//
const Qt__InvertedPortraitOrientation Qt__ScreenOrientation = 4

//
const Qt__InvertedLandscapeOrientation Qt__ScreenOrientation = 8

/*
This enum type is used to specify various widget attributes. Attributes are set and cleared with QWidget::setAttribute(), and queried with QWidget::testAttribute(), although some have special convenience functions which are mentioned below.

Qt::WA_NoBackgroundWA_OpaquePaintEventThis value is obsolete. Use WA_OpaquePaintEvent instead.

*/
type Qt__WidgetAttribute = int // core
// Indicates that the widget is disabled, i.e. it does not receive any mouse or keyboard events. There is also a getter functions QWidget::isEnabled(). This is set/cleared by the Qt kernel.
const Qt__WA_Disabled Qt__WidgetAttribute = 0

// Indicates that the widget is under the mouse cursor. The value is not updated correctly during drag and drop operations. There is also a getter function, QWidget::underMouse(). This flag is set or cleared by the Qt kernel.
const Qt__WA_UnderMouse Qt__WidgetAttribute = 1

// Indicates that the widget has mouse tracking enabled. See QWidget::mouseTracking.
const Qt__WA_MouseTracking Qt__WidgetAttribute = 2

//
const Qt__WA_ContentsPropagated Qt__WidgetAttribute = 3

// Indicates that the widget paints all its pixels when it receives a paint event. Thus, it is not required for operations like updating, resizing, scrolling and focus changes to erase the widget before generating paint events. The use of WA_OpaquePaintEvent provides a small optimization by helping to reduce flicker on systems that do not support double buffering and avoiding computational cycles necessary to erase the background prior to painting. Note: Unlike WA_NoSystemBackground, WA_OpaquePaintEvent makes an effort to avoid transparent window backgrounds. This flag is set or cleared by the widget's author.
const Qt__WA_OpaquePaintEvent Qt__WidgetAttribute = 4

//
const Qt__WA_NoBackground Qt__WidgetAttribute = 4

// Indicates that the widget contents are north-west aligned and static. On resize, such a widget will receive paint events only for parts of itself that are newly visible. This flag is set or cleared by the widget's author.
const Qt__WA_StaticContents Qt__WidgetAttribute = 5

//
const Qt__WA_LaidOut Qt__WidgetAttribute = 7

//
const Qt__WA_PaintOnScreen Qt__WidgetAttribute = 8

// Indicates that the widget has no background, i.e. when the widget receives paint events, the background is not automatically repainted. Note: Unlike WA_OpaquePaintEvent, newly exposed areas are never filled with the background (e.g., after showing a window for the first time the user can see "through" it until the application processes the paint events). This flag is set or cleared by the widget's author.
const Qt__WA_NoSystemBackground Qt__WidgetAttribute = 9

//
const Qt__WA_UpdatesDisabled Qt__WidgetAttribute = 10

//
const Qt__WA_Mapped Qt__WidgetAttribute = 11

//
const Qt__WA_MacNoClickThrough Qt__WidgetAttribute = 12

//
const Qt__WA_InputMethodEnabled Qt__WidgetAttribute = 14

//
const Qt__WA_WState_Visible Qt__WidgetAttribute = 15

//
const Qt__WA_WState_Hidden Qt__WidgetAttribute = 16

//
const Qt__WA_ForceDisabled Qt__WidgetAttribute = 32

//
const Qt__WA_KeyCompression Qt__WidgetAttribute = 33

//
const Qt__WA_PendingMoveEvent Qt__WidgetAttribute = 34

//
const Qt__WA_PendingResizeEvent Qt__WidgetAttribute = 35

//
const Qt__WA_SetPalette Qt__WidgetAttribute = 36

//
const Qt__WA_SetFont Qt__WidgetAttribute = 37

//
const Qt__WA_SetCursor Qt__WidgetAttribute = 38

//
const Qt__WA_NoChildEventsFromChildren Qt__WidgetAttribute = 39

//
const Qt__WA_WindowModified Qt__WidgetAttribute = 41

//
const Qt__WA_Resized Qt__WidgetAttribute = 42

//
const Qt__WA_Moved Qt__WidgetAttribute = 43

//
const Qt__WA_PendingUpdate Qt__WidgetAttribute = 44

//
const Qt__WA_InvalidSize Qt__WidgetAttribute = 45

//
const Qt__WA_MacBrushedMetal Qt__WidgetAttribute = 46

//
const Qt__WA_MacMetalStyle Qt__WidgetAttribute = 46

//
const Qt__WA_CustomWhatsThis Qt__WidgetAttribute = 47

//
const Qt__WA_LayoutOnEntireRect Qt__WidgetAttribute = 48

//
const Qt__WA_OutsideWSRange Qt__WidgetAttribute = 49

//
const Qt__WA_GrabbedShortcut Qt__WidgetAttribute = 50

//
const Qt__WA_TransparentForMouseEvents Qt__WidgetAttribute = 51

//
const Qt__WA_PaintUnclipped Qt__WidgetAttribute = 52

//
const Qt__WA_SetWindowIcon Qt__WidgetAttribute = 53

//
const Qt__WA_NoMouseReplay Qt__WidgetAttribute = 54

//
const Qt__WA_DeleteOnClose Qt__WidgetAttribute = 55

//
const Qt__WA_RightToLeft Qt__WidgetAttribute = 56

//
const Qt__WA_SetLayoutDirection Qt__WidgetAttribute = 57

//
const Qt__WA_NoChildEventsForParent Qt__WidgetAttribute = 58

//
const Qt__WA_ForceUpdatesDisabled Qt__WidgetAttribute = 59

//
const Qt__WA_WState_Created Qt__WidgetAttribute = 60

//
const Qt__WA_WState_CompressKeys Qt__WidgetAttribute = 61

//
const Qt__WA_WState_InPaintEvent Qt__WidgetAttribute = 62

//
const Qt__WA_WState_Reparented Qt__WidgetAttribute = 63

//
const Qt__WA_WState_ConfigPending Qt__WidgetAttribute = 64

//
const Qt__WA_WState_Polished Qt__WidgetAttribute = 66

//
const Qt__WA_WState_DND Qt__WidgetAttribute = 67

//
const Qt__WA_WState_OwnSizePolicy Qt__WidgetAttribute = 68

//
const Qt__WA_WState_ExplicitShowHide Qt__WidgetAttribute = 69

//
const Qt__WA_ShowModal Qt__WidgetAttribute = 70

//
const Qt__WA_MouseNoMask Qt__WidgetAttribute = 71

//
const Qt__WA_GroupLeader Qt__WidgetAttribute = 72

//
const Qt__WA_NoMousePropagation Qt__WidgetAttribute = 73

//
const Qt__WA_Hover Qt__WidgetAttribute = 74

//
const Qt__WA_InputMethodTransparent Qt__WidgetAttribute = 75

//
const Qt__WA_QuitOnClose Qt__WidgetAttribute = 76

//
const Qt__WA_KeyboardFocusChange Qt__WidgetAttribute = 77

//
const Qt__WA_AcceptDrops Qt__WidgetAttribute = 78

//
const Qt__WA_DropSiteRegistered Qt__WidgetAttribute = 79

//
const Qt__WA_ForceAcceptDrops Qt__WidgetAttribute = 79

//
const Qt__WA_WindowPropagation Qt__WidgetAttribute = 80

//
const Qt__WA_NoX11EventCompression Qt__WidgetAttribute = 81

//
const Qt__WA_TintedBackground Qt__WidgetAttribute = 82

//
const Qt__WA_X11OpenGLOverlay Qt__WidgetAttribute = 83

//
const Qt__WA_AlwaysShowToolTips Qt__WidgetAttribute = 84

//
const Qt__WA_MacOpaqueSizeGrip Qt__WidgetAttribute = 85

//
const Qt__WA_SetStyle Qt__WidgetAttribute = 86

//
const Qt__WA_SetLocale Qt__WidgetAttribute = 87

//
const Qt__WA_MacShowFocusRect Qt__WidgetAttribute = 88

//
const Qt__WA_MacNormalSize Qt__WidgetAttribute = 89

//
const Qt__WA_MacSmallSize Qt__WidgetAttribute = 90

//
const Qt__WA_MacMiniSize Qt__WidgetAttribute = 91

//
const Qt__WA_LayoutUsesWidgetRect Qt__WidgetAttribute = 92

//
const Qt__WA_StyledBackground Qt__WidgetAttribute = 93

//
const Qt__WA_MSWindowsUseDirect3D Qt__WidgetAttribute = 94

//
const Qt__WA_CanHostQMdiSubWindowTitleBar Qt__WidgetAttribute = 95

//
const Qt__WA_MacAlwaysShowToolWindow Qt__WidgetAttribute = 96

//
const Qt__WA_StyleSheet Qt__WidgetAttribute = 97

//
const Qt__WA_ShowWithoutActivating Qt__WidgetAttribute = 98

//
const Qt__WA_X11BypassTransientForHint Qt__WidgetAttribute = 99

//
const Qt__WA_NativeWindow Qt__WidgetAttribute = 100

//
const Qt__WA_DontCreateNativeAncestors Qt__WidgetAttribute = 101

//
const Qt__WA_MacVariableSize Qt__WidgetAttribute = 102

//
const Qt__WA_DontShowOnScreen Qt__WidgetAttribute = 103

//
const Qt__WA_X11NetWmWindowTypeDesktop Qt__WidgetAttribute = 104

//
const Qt__WA_X11NetWmWindowTypeDock Qt__WidgetAttribute = 105

//
const Qt__WA_X11NetWmWindowTypeToolBar Qt__WidgetAttribute = 106

//
const Qt__WA_X11NetWmWindowTypeMenu Qt__WidgetAttribute = 107

//
const Qt__WA_X11NetWmWindowTypeUtility Qt__WidgetAttribute = 108

//
const Qt__WA_X11NetWmWindowTypeSplash Qt__WidgetAttribute = 109

//
const Qt__WA_X11NetWmWindowTypeDialog Qt__WidgetAttribute = 110

//
const Qt__WA_X11NetWmWindowTypeDropDownMenu Qt__WidgetAttribute = 111

//
const Qt__WA_X11NetWmWindowTypePopupMenu Qt__WidgetAttribute = 112

//
const Qt__WA_X11NetWmWindowTypeToolTip Qt__WidgetAttribute = 113

//
const Qt__WA_X11NetWmWindowTypeNotification Qt__WidgetAttribute = 114

//
const Qt__WA_X11NetWmWindowTypeCombo Qt__WidgetAttribute = 115

//
const Qt__WA_X11NetWmWindowTypeDND Qt__WidgetAttribute = 116

//
const Qt__WA_MacFrameworkScaled Qt__WidgetAttribute = 117

//
const Qt__WA_SetWindowModality Qt__WidgetAttribute = 118

//
const Qt__WA_WState_WindowOpacitySet Qt__WidgetAttribute = 119

//
const Qt__WA_TranslucentBackground Qt__WidgetAttribute = 120

//
const Qt__WA_AcceptTouchEvents Qt__WidgetAttribute = 121

//
const Qt__WA_WState_AcceptedTouchBeginEvent Qt__WidgetAttribute = 122

//
const Qt__WA_TouchPadAcceptSingleTouchEvents Qt__WidgetAttribute = 123

//
const Qt__WA_X11DoNotAcceptFocus Qt__WidgetAttribute = 126

//
const Qt__WA_MacNoShadow Qt__WidgetAttribute = 127

//
const Qt__WA_AlwaysStackOnTop Qt__WidgetAttribute = 128

//
const Qt__WA_TabletTracking Qt__WidgetAttribute = 129

//
const Qt__WA_ContentsMarginsRespectsSafeArea Qt__WidgetAttribute = 130

//
const Qt__WA_AttributeCount Qt__WidgetAttribute = 131

/*
This enum describes attributes that change the behavior of application-wide features. These are enabled and disabled using QCoreApplication::setAttribute(), and can be tested for with QCoreApplication::testAttribute().



The following values are deprecated or obsolete:

Qt::AA_MacPluginApplicationAA_PluginApplicationThis attribute has been deprecated. Use AA_PluginApplication instead.

*/
type Qt__ApplicationAttribute = int // core
//
const Qt__AA_ImmediateWidgetCreation Qt__ApplicationAttribute = 0

// This value is obsolete and has no effect.
const Qt__AA_MSWindowsUseDirect3DByDefault Qt__ApplicationAttribute = 1

// Actions with the Icon property won't be shown in any menus unless specifically set by the QAction::iconVisibleInMenu property. Menus that are currently open or menus already created in the native macOS menubar may not pick up a change in this attribute. Changes in the QAction::iconVisibleInMenu property will always be picked up.
const Qt__AA_DontShowIconsInMenus Qt__ApplicationAttribute = 2

// Ensures that widgets have native windows.
const Qt__AA_NativeWindows Qt__ApplicationAttribute = 3

// Ensures that siblings of native widgets stay non-native unless specifically set by the Qt::WA_NativeWindow attribute.
const Qt__AA_DontCreateNativeWidgetSiblings Qt__ApplicationAttribute = 4

//
const Qt__AA_PluginApplication Qt__ApplicationAttribute = 5

//
const Qt__AA_MacPluginApplication Qt__ApplicationAttribute = 5

// All menubars created while this attribute is set to true won't be used as a native menubar (e.g, the menubar at the top of the main screen on macOS).
const Qt__AA_DontUseNativeMenuBar Qt__ApplicationAttribute = 6

// On macOS by default, Qt swaps the Control and Meta (Command) keys (i.e., whenever Control is pressed, Qt sends Meta, and whenever Meta is pressed Control is sent). When this attribute is true, Qt will not do the flip. QKeySequence::StandardKey will also flip accordingly (i.e., QKeySequence::Copy will be Command+C on the keyboard regardless of the value set, though what is output for QKeySequence::toString() will be different).
const Qt__AA_MacDontSwapCtrlAndMeta Qt__ApplicationAttribute = 7

//
const Qt__AA_Use96Dpi Qt__ApplicationAttribute = 8

//
const Qt__AA_X11InitThreads Qt__ApplicationAttribute = 10

//
const Qt__AA_SynthesizeTouchForUnhandledMouseEvents Qt__ApplicationAttribute = 11

//
const Qt__AA_SynthesizeMouseForUnhandledTouchEvents Qt__ApplicationAttribute = 12

//
const Qt__AA_UseHighDpiPixmaps Qt__ApplicationAttribute = 13

//
const Qt__AA_ForceRasterWidgets Qt__ApplicationAttribute = 14

//
const Qt__AA_UseDesktopOpenGL Qt__ApplicationAttribute = 15

//
const Qt__AA_UseOpenGLES Qt__ApplicationAttribute = 16

//
const Qt__AA_UseSoftwareOpenGL Qt__ApplicationAttribute = 17

//
const Qt__AA_ShareOpenGLContexts Qt__ApplicationAttribute = 18

//
const Qt__AA_SetPalette Qt__ApplicationAttribute = 19

//
const Qt__AA_EnableHighDpiScaling Qt__ApplicationAttribute = 20

//
const Qt__AA_DisableHighDpiScaling Qt__ApplicationAttribute = 21

//
const Qt__AA_UseStyleSheetPropagationInWidgetStyles Qt__ApplicationAttribute = 22

//
const Qt__AA_DontUseNativeDialogs Qt__ApplicationAttribute = 23

//
const Qt__AA_SynthesizeMouseForUnhandledTabletEvents Qt__ApplicationAttribute = 24

//
const Qt__AA_CompressHighFrequencyEvents Qt__ApplicationAttribute = 25

//
const Qt__AA_DontCheckOpenGLContextThreadAffinity Qt__ApplicationAttribute = 26

//
const Qt__AA_DisableShaderDiskCache Qt__ApplicationAttribute = 27

//
const Qt__AA_DontShowShortcutsInContextMenus Qt__ApplicationAttribute = 28

//
const Qt__AA_CompressTabletEvents Qt__ApplicationAttribute = 29

//
const Qt__AA_DisableWindowContextHelpButton Qt__ApplicationAttribute = 30

//
const Qt__AA_AttributeCount Qt__ApplicationAttribute = 31

/*


 */
type Qt__ImageConversionFlag = int // core
//
const Qt__ColorMode_Mask Qt__ImageConversionFlag = 3

//
const Qt__AutoColor Qt__ImageConversionFlag = 0

//
const Qt__ColorOnly Qt__ImageConversionFlag = 3

//
const Qt__MonoOnly Qt__ImageConversionFlag = 2

//
const Qt__AlphaDither_Mask Qt__ImageConversionFlag = 12

//
const Qt__ThresholdAlphaDither Qt__ImageConversionFlag = 0

//
const Qt__OrderedAlphaDither Qt__ImageConversionFlag = 4

//
const Qt__DiffuseAlphaDither Qt__ImageConversionFlag = 8

//
const Qt__NoAlpha Qt__ImageConversionFlag = 12

//
const Qt__Dither_Mask Qt__ImageConversionFlag = 48

//
const Qt__DiffuseDither Qt__ImageConversionFlag = 0

//
const Qt__OrderedDither Qt__ImageConversionFlag = 16

//
const Qt__ThresholdDither Qt__ImageConversionFlag = 32

//
const Qt__DitherMode_Mask Qt__ImageConversionFlag = 192

//
const Qt__AutoDither Qt__ImageConversionFlag = 0

//
const Qt__PreferDither Qt__ImageConversionFlag = 64

//
const Qt__AvoidDither Qt__ImageConversionFlag = 128

//
const Qt__NoOpaqueDetection Qt__ImageConversionFlag = 256

//
const Qt__NoFormatConversion Qt__ImageConversionFlag = 512

/*
Background mode:

ConstantValue
Qt::TransparentMode0
Qt::OpaqueMode1

*/
type Qt__BGMode = int // core
//
const Qt__TransparentMode Qt__BGMode = 0

//
const Qt__OpaqueMode Qt__BGMode = 1

/*
The key names used by Qt.

Qt::Key_AnyKey_Space


See also QKeyEvent::key().

*/
type Qt__Key = int // core
//
const Qt__Key_Escape Qt__Key = 16777216

//
const Qt__Key_Tab Qt__Key = 16777217

//
const Qt__Key_Backtab Qt__Key = 16777218

//
const Qt__Key_Backspace Qt__Key = 16777219

//
const Qt__Key_Return Qt__Key = 16777220

//
const Qt__Key_Enter Qt__Key = 16777221

//
const Qt__Key_Insert Qt__Key = 16777222

//
const Qt__Key_Delete Qt__Key = 16777223

//
const Qt__Key_Pause Qt__Key = 16777224

//
const Qt__Key_Print Qt__Key = 16777225

//
const Qt__Key_SysReq Qt__Key = 16777226

//
const Qt__Key_Clear Qt__Key = 16777227

//
const Qt__Key_Home Qt__Key = 16777232

//
const Qt__Key_End Qt__Key = 16777233

//
const Qt__Key_Left Qt__Key = 16777234

//
const Qt__Key_Up Qt__Key = 16777235

//
const Qt__Key_Right Qt__Key = 16777236

//
const Qt__Key_Down Qt__Key = 16777237

//
const Qt__Key_PageUp Qt__Key = 16777238

//
const Qt__Key_PageDown Qt__Key = 16777239

//
const Qt__Key_Shift Qt__Key = 16777248

//
const Qt__Key_Control Qt__Key = 16777249

//
const Qt__Key_Meta Qt__Key = 16777250

//
const Qt__Key_Alt Qt__Key = 16777251

//
const Qt__Key_CapsLock Qt__Key = 16777252

//
const Qt__Key_NumLock Qt__Key = 16777253

//
const Qt__Key_ScrollLock Qt__Key = 16777254

//
const Qt__Key_F1 Qt__Key = 16777264

//
const Qt__Key_F2 Qt__Key = 16777265

//
const Qt__Key_F3 Qt__Key = 16777266

//
const Qt__Key_F4 Qt__Key = 16777267

//
const Qt__Key_F5 Qt__Key = 16777268

//
const Qt__Key_F6 Qt__Key = 16777269

//
const Qt__Key_F7 Qt__Key = 16777270

//
const Qt__Key_F8 Qt__Key = 16777271

//
const Qt__Key_F9 Qt__Key = 16777272

//
const Qt__Key_F10 Qt__Key = 16777273

//
const Qt__Key_F11 Qt__Key = 16777274

//
const Qt__Key_F12 Qt__Key = 16777275

//
const Qt__Key_F13 Qt__Key = 16777276

//
const Qt__Key_F14 Qt__Key = 16777277

//
const Qt__Key_F15 Qt__Key = 16777278

//
const Qt__Key_F16 Qt__Key = 16777279

//
const Qt__Key_F17 Qt__Key = 16777280

//
const Qt__Key_F18 Qt__Key = 16777281

//
const Qt__Key_F19 Qt__Key = 16777282

//
const Qt__Key_F20 Qt__Key = 16777283

//
const Qt__Key_F21 Qt__Key = 16777284

//
const Qt__Key_F22 Qt__Key = 16777285

//
const Qt__Key_F23 Qt__Key = 16777286

//
const Qt__Key_F24 Qt__Key = 16777287

//
const Qt__Key_F25 Qt__Key = 16777288

//
const Qt__Key_F26 Qt__Key = 16777289

//
const Qt__Key_F27 Qt__Key = 16777290

//
const Qt__Key_F28 Qt__Key = 16777291

//
const Qt__Key_F29 Qt__Key = 16777292

//
const Qt__Key_F30 Qt__Key = 16777293

//
const Qt__Key_F31 Qt__Key = 16777294

//
const Qt__Key_F32 Qt__Key = 16777295

//
const Qt__Key_F33 Qt__Key = 16777296

//
const Qt__Key_F34 Qt__Key = 16777297

//
const Qt__Key_F35 Qt__Key = 16777298

//
const Qt__Key_Super_L Qt__Key = 16777299

//
const Qt__Key_Super_R Qt__Key = 16777300

//
const Qt__Key_Menu Qt__Key = 16777301

//
const Qt__Key_Hyper_L Qt__Key = 16777302

//
const Qt__Key_Hyper_R Qt__Key = 16777303

//
const Qt__Key_Help Qt__Key = 16777304

//
const Qt__Key_Direction_L Qt__Key = 16777305

//
const Qt__Key_Direction_R Qt__Key = 16777312

//
const Qt__Key_Space Qt__Key = 32

//
const Qt__Key_Any Qt__Key = 32

//
const Qt__Key_Exclam Qt__Key = 33

//
const Qt__Key_QuoteDbl Qt__Key = 34

//
const Qt__Key_NumberSign Qt__Key = 35

//
const Qt__Key_Dollar Qt__Key = 36

//
const Qt__Key_Percent Qt__Key = 37

//
const Qt__Key_Ampersand Qt__Key = 38

//
const Qt__Key_Apostrophe Qt__Key = 39

//
const Qt__Key_ParenLeft Qt__Key = 40

//
const Qt__Key_ParenRight Qt__Key = 41

//
const Qt__Key_Asterisk Qt__Key = 42

//
const Qt__Key_Plus Qt__Key = 43

//
const Qt__Key_Comma Qt__Key = 44

//
const Qt__Key_Minus Qt__Key = 45

//
const Qt__Key_Period Qt__Key = 46

//
const Qt__Key_Slash Qt__Key = 47

//
const Qt__Key_0 Qt__Key = 48

//
const Qt__Key_1 Qt__Key = 49

//
const Qt__Key_2 Qt__Key = 50

//
const Qt__Key_3 Qt__Key = 51

//
const Qt__Key_4 Qt__Key = 52

//
const Qt__Key_5 Qt__Key = 53

//
const Qt__Key_6 Qt__Key = 54

//
const Qt__Key_7 Qt__Key = 55

//
const Qt__Key_8 Qt__Key = 56

//
const Qt__Key_9 Qt__Key = 57

//
const Qt__Key_Colon Qt__Key = 58

//
const Qt__Key_Semicolon Qt__Key = 59

//
const Qt__Key_Less Qt__Key = 60

//
const Qt__Key_Equal Qt__Key = 61

//
const Qt__Key_Greater Qt__Key = 62

//
const Qt__Key_Question Qt__Key = 63

//
const Qt__Key_At Qt__Key = 64

//
const Qt__Key_A Qt__Key = 65

//
const Qt__Key_B Qt__Key = 66

//
const Qt__Key_C Qt__Key = 67

//
const Qt__Key_D Qt__Key = 68

//
const Qt__Key_E Qt__Key = 69

//
const Qt__Key_F Qt__Key = 70

//
const Qt__Key_G Qt__Key = 71

//
const Qt__Key_H Qt__Key = 72

//
const Qt__Key_I Qt__Key = 73

//
const Qt__Key_J Qt__Key = 74

//
const Qt__Key_K Qt__Key = 75

//
const Qt__Key_L Qt__Key = 76

//
const Qt__Key_M Qt__Key = 77

//
const Qt__Key_N Qt__Key = 78

//
const Qt__Key_O Qt__Key = 79

//
const Qt__Key_P Qt__Key = 80

//
const Qt__Key_Q Qt__Key = 81

//
const Qt__Key_R Qt__Key = 82

//
const Qt__Key_S Qt__Key = 83

//
const Qt__Key_T Qt__Key = 84

//
const Qt__Key_U Qt__Key = 85

//
const Qt__Key_V Qt__Key = 86

//
const Qt__Key_W Qt__Key = 87

//
const Qt__Key_X Qt__Key = 88

//
const Qt__Key_Y Qt__Key = 89

//
const Qt__Key_Z Qt__Key = 90

//
const Qt__Key_BracketLeft Qt__Key = 91

//
const Qt__Key_Backslash Qt__Key = 92

//
const Qt__Key_BracketRight Qt__Key = 93

//
const Qt__Key_AsciiCircum Qt__Key = 94

//
const Qt__Key_Underscore Qt__Key = 95

//
const Qt__Key_QuoteLeft Qt__Key = 96

//
const Qt__Key_BraceLeft Qt__Key = 123

//
const Qt__Key_Bar Qt__Key = 124

//
const Qt__Key_BraceRight Qt__Key = 125

//
const Qt__Key_AsciiTilde Qt__Key = 126

//
const Qt__Key_nobreakspace Qt__Key = 160

//
const Qt__Key_exclamdown Qt__Key = 161

//
const Qt__Key_cent Qt__Key = 162

//
const Qt__Key_sterling Qt__Key = 163

//
const Qt__Key_currency Qt__Key = 164

//
const Qt__Key_yen Qt__Key = 165

//
const Qt__Key_brokenbar Qt__Key = 166

//
const Qt__Key_section Qt__Key = 167

//
const Qt__Key_diaeresis Qt__Key = 168

//
const Qt__Key_copyright Qt__Key = 169

//
const Qt__Key_ordfeminine Qt__Key = 170

//
const Qt__Key_guillemotleft Qt__Key = 171

//
const Qt__Key_notsign Qt__Key = 172

//
const Qt__Key_hyphen Qt__Key = 173

//
const Qt__Key_registered Qt__Key = 174

//
const Qt__Key_macron Qt__Key = 175

//
const Qt__Key_degree Qt__Key = 176

//
const Qt__Key_plusminus Qt__Key = 177

//
const Qt__Key_twosuperior Qt__Key = 178

//
const Qt__Key_threesuperior Qt__Key = 179

//
const Qt__Key_acute Qt__Key = 180

//
const Qt__Key_mu Qt__Key = 181

//
const Qt__Key_paragraph Qt__Key = 182

//
const Qt__Key_periodcentered Qt__Key = 183

//
const Qt__Key_cedilla Qt__Key = 184

//
const Qt__Key_onesuperior Qt__Key = 185

//
const Qt__Key_masculine Qt__Key = 186

//
const Qt__Key_guillemotright Qt__Key = 187

//
const Qt__Key_onequarter Qt__Key = 188

//
const Qt__Key_onehalf Qt__Key = 189

//
const Qt__Key_threequarters Qt__Key = 190

//
const Qt__Key_questiondown Qt__Key = 191

//
const Qt__Key_Agrave Qt__Key = 192

//
const Qt__Key_Aacute Qt__Key = 193

//
const Qt__Key_Acircumflex Qt__Key = 194

//
const Qt__Key_Atilde Qt__Key = 195

//
const Qt__Key_Adiaeresis Qt__Key = 196

//
const Qt__Key_Aring Qt__Key = 197

//
const Qt__Key_AE Qt__Key = 198

//
const Qt__Key_Ccedilla Qt__Key = 199

//
const Qt__Key_Egrave Qt__Key = 200

//
const Qt__Key_Eacute Qt__Key = 201

//
const Qt__Key_Ecircumflex Qt__Key = 202

//
const Qt__Key_Ediaeresis Qt__Key = 203

//
const Qt__Key_Igrave Qt__Key = 204

//
const Qt__Key_Iacute Qt__Key = 205

//
const Qt__Key_Icircumflex Qt__Key = 206

//
const Qt__Key_Idiaeresis Qt__Key = 207

//
const Qt__Key_ETH Qt__Key = 208

//
const Qt__Key_Ntilde Qt__Key = 209

//
const Qt__Key_Ograve Qt__Key = 210

//
const Qt__Key_Oacute Qt__Key = 211

//
const Qt__Key_Ocircumflex Qt__Key = 212

//
const Qt__Key_Otilde Qt__Key = 213

//
const Qt__Key_Odiaeresis Qt__Key = 214

//
const Qt__Key_multiply Qt__Key = 215

//
const Qt__Key_Ooblique Qt__Key = 216

//
const Qt__Key_Ugrave Qt__Key = 217

//
const Qt__Key_Uacute Qt__Key = 218

//
const Qt__Key_Ucircumflex Qt__Key = 219

//
const Qt__Key_Udiaeresis Qt__Key = 220

//
const Qt__Key_Yacute Qt__Key = 221

//
const Qt__Key_THORN Qt__Key = 222

//
const Qt__Key_ssharp Qt__Key = 223

//
const Qt__Key_division Qt__Key = 247

//
const Qt__Key_ydiaeresis Qt__Key = 255

//
const Qt__Key_AltGr Qt__Key = 16781571

//
const Qt__Key_Multi_key Qt__Key = 16781600

//
const Qt__Key_Codeinput Qt__Key = 16781623

//
const Qt__Key_SingleCandidate Qt__Key = 16781628

//
const Qt__Key_MultipleCandidate Qt__Key = 16781629

//
const Qt__Key_PreviousCandidate Qt__Key = 16781630

//
const Qt__Key_Mode_switch Qt__Key = 16781694

//
const Qt__Key_Kanji Qt__Key = 16781601

//
const Qt__Key_Muhenkan Qt__Key = 16781602

//
const Qt__Key_Henkan Qt__Key = 16781603

//
const Qt__Key_Romaji Qt__Key = 16781604

//
const Qt__Key_Hiragana Qt__Key = 16781605

//
const Qt__Key_Katakana Qt__Key = 16781606

//
const Qt__Key_Hiragana_Katakana Qt__Key = 16781607

//
const Qt__Key_Zenkaku Qt__Key = 16781608

//
const Qt__Key_Hankaku Qt__Key = 16781609

//
const Qt__Key_Zenkaku_Hankaku Qt__Key = 16781610

//
const Qt__Key_Touroku Qt__Key = 16781611

//
const Qt__Key_Massyo Qt__Key = 16781612

//
const Qt__Key_Kana_Lock Qt__Key = 16781613

//
const Qt__Key_Kana_Shift Qt__Key = 16781614

//
const Qt__Key_Eisu_Shift Qt__Key = 16781615

//
const Qt__Key_Eisu_toggle Qt__Key = 16781616

//
const Qt__Key_Hangul Qt__Key = 16781617

//
const Qt__Key_Hangul_Start Qt__Key = 16781618

//
const Qt__Key_Hangul_End Qt__Key = 16781619

//
const Qt__Key_Hangul_Hanja Qt__Key = 16781620

//
const Qt__Key_Hangul_Jamo Qt__Key = 16781621

//
const Qt__Key_Hangul_Romaja Qt__Key = 16781622

//
const Qt__Key_Hangul_Jeonja Qt__Key = 16781624

//
const Qt__Key_Hangul_Banja Qt__Key = 16781625

//
const Qt__Key_Hangul_PreHanja Qt__Key = 16781626

//
const Qt__Key_Hangul_PostHanja Qt__Key = 16781627

//
const Qt__Key_Hangul_Special Qt__Key = 16781631

//
const Qt__Key_Dead_Grave Qt__Key = 16781904

//
const Qt__Key_Dead_Acute Qt__Key = 16781905

//
const Qt__Key_Dead_Circumflex Qt__Key = 16781906

//
const Qt__Key_Dead_Tilde Qt__Key = 16781907

//
const Qt__Key_Dead_Macron Qt__Key = 16781908

//
const Qt__Key_Dead_Breve Qt__Key = 16781909

//
const Qt__Key_Dead_Abovedot Qt__Key = 16781910

//
const Qt__Key_Dead_Diaeresis Qt__Key = 16781911

//
const Qt__Key_Dead_Abovering Qt__Key = 16781912

//
const Qt__Key_Dead_Doubleacute Qt__Key = 16781913

//
const Qt__Key_Dead_Caron Qt__Key = 16781914

//
const Qt__Key_Dead_Cedilla Qt__Key = 16781915

//
const Qt__Key_Dead_Ogonek Qt__Key = 16781916

//
const Qt__Key_Dead_Iota Qt__Key = 16781917

//
const Qt__Key_Dead_Voiced_Sound Qt__Key = 16781918

//
const Qt__Key_Dead_Semivoiced_Sound Qt__Key = 16781919

//
const Qt__Key_Dead_Belowdot Qt__Key = 16781920

//
const Qt__Key_Dead_Hook Qt__Key = 16781921

//
const Qt__Key_Dead_Horn Qt__Key = 16781922

//
const Qt__Key_Back Qt__Key = 16777313

//
const Qt__Key_Forward Qt__Key = 16777314

//
const Qt__Key_Stop Qt__Key = 16777315

//
const Qt__Key_Refresh Qt__Key = 16777316

//
const Qt__Key_VolumeDown Qt__Key = 16777328

//
const Qt__Key_VolumeMute Qt__Key = 16777329

//
const Qt__Key_VolumeUp Qt__Key = 16777330

//
const Qt__Key_BassBoost Qt__Key = 16777331

//
const Qt__Key_BassUp Qt__Key = 16777332

//
const Qt__Key_BassDown Qt__Key = 16777333

//
const Qt__Key_TrebleUp Qt__Key = 16777334

//
const Qt__Key_TrebleDown Qt__Key = 16777335

//
const Qt__Key_MediaPlay Qt__Key = 16777344

//
const Qt__Key_MediaStop Qt__Key = 16777345

//
const Qt__Key_MediaPrevious Qt__Key = 16777346

//
const Qt__Key_MediaNext Qt__Key = 16777347

//
const Qt__Key_MediaRecord Qt__Key = 16777348

//
const Qt__Key_MediaPause Qt__Key = 16777349

//
const Qt__Key_MediaTogglePlayPause Qt__Key = 16777350

//
const Qt__Key_HomePage Qt__Key = 16777360

//
const Qt__Key_Favorites Qt__Key = 16777361

//
const Qt__Key_Search Qt__Key = 16777362

//
const Qt__Key_Standby Qt__Key = 16777363

//
const Qt__Key_OpenUrl Qt__Key = 16777364

//
const Qt__Key_LaunchMail Qt__Key = 16777376

//
const Qt__Key_LaunchMedia Qt__Key = 16777377

//
const Qt__Key_Launch0 Qt__Key = 16777378

//
const Qt__Key_Launch1 Qt__Key = 16777379

//
const Qt__Key_Launch2 Qt__Key = 16777380

//
const Qt__Key_Launch3 Qt__Key = 16777381

//
const Qt__Key_Launch4 Qt__Key = 16777382

//
const Qt__Key_Launch5 Qt__Key = 16777383

//
const Qt__Key_Launch6 Qt__Key = 16777384

//
const Qt__Key_Launch7 Qt__Key = 16777385

//
const Qt__Key_Launch8 Qt__Key = 16777386

//
const Qt__Key_Launch9 Qt__Key = 16777387

//
const Qt__Key_LaunchA Qt__Key = 16777388

//
const Qt__Key_LaunchB Qt__Key = 16777389

//
const Qt__Key_LaunchC Qt__Key = 16777390

//
const Qt__Key_LaunchD Qt__Key = 16777391

//
const Qt__Key_LaunchE Qt__Key = 16777392

//
const Qt__Key_LaunchF Qt__Key = 16777393

//
const Qt__Key_MonBrightnessUp Qt__Key = 16777394

//
const Qt__Key_MonBrightnessDown Qt__Key = 16777395

//
const Qt__Key_KeyboardLightOnOff Qt__Key = 16777396

//
const Qt__Key_KeyboardBrightnessUp Qt__Key = 16777397

//
const Qt__Key_KeyboardBrightnessDown Qt__Key = 16777398

//
const Qt__Key_PowerOff Qt__Key = 16777399

//
const Qt__Key_WakeUp Qt__Key = 16777400

//
const Qt__Key_Eject Qt__Key = 16777401

//
const Qt__Key_ScreenSaver Qt__Key = 16777402

//
const Qt__Key_WWW Qt__Key = 16777403

//
const Qt__Key_Memo Qt__Key = 16777404

//
const Qt__Key_LightBulb Qt__Key = 16777405

//
const Qt__Key_Shop Qt__Key = 16777406

//
const Qt__Key_History Qt__Key = 16777407

//
const Qt__Key_AddFavorite Qt__Key = 16777408

//
const Qt__Key_HotLinks Qt__Key = 16777409

//
const Qt__Key_BrightnessAdjust Qt__Key = 16777410

//
const Qt__Key_Finance Qt__Key = 16777411

//
const Qt__Key_Community Qt__Key = 16777412

//
const Qt__Key_AudioRewind Qt__Key = 16777413

//
const Qt__Key_BackForward Qt__Key = 16777414

//
const Qt__Key_ApplicationLeft Qt__Key = 16777415

//
const Qt__Key_ApplicationRight Qt__Key = 16777416

//
const Qt__Key_Book Qt__Key = 16777417

//
const Qt__Key_CD Qt__Key = 16777418

//
const Qt__Key_Calculator Qt__Key = 16777419

//
const Qt__Key_ToDoList Qt__Key = 16777420

//
const Qt__Key_ClearGrab Qt__Key = 16777421

//
const Qt__Key_Close Qt__Key = 16777422

//
const Qt__Key_Copy Qt__Key = 16777423

//
const Qt__Key_Cut Qt__Key = 16777424

//
const Qt__Key_Display Qt__Key = 16777425

//
const Qt__Key_DOS Qt__Key = 16777426

//
const Qt__Key_Documents Qt__Key = 16777427

//
const Qt__Key_Excel Qt__Key = 16777428

//
const Qt__Key_Explorer Qt__Key = 16777429

//
const Qt__Key_Game Qt__Key = 16777430

//
const Qt__Key_Go Qt__Key = 16777431

//
const Qt__Key_iTouch Qt__Key = 16777432

//
const Qt__Key_LogOff Qt__Key = 16777433

//
const Qt__Key_Market Qt__Key = 16777434

//
const Qt__Key_Meeting Qt__Key = 16777435

//
const Qt__Key_MenuKB Qt__Key = 16777436

//
const Qt__Key_MenuPB Qt__Key = 16777437

//
const Qt__Key_MySites Qt__Key = 16777438

//
const Qt__Key_News Qt__Key = 16777439

//
const Qt__Key_OfficeHome Qt__Key = 16777440

//
const Qt__Key_Option Qt__Key = 16777441

//
const Qt__Key_Paste Qt__Key = 16777442

//
const Qt__Key_Phone Qt__Key = 16777443

//
const Qt__Key_Calendar Qt__Key = 16777444

//
const Qt__Key_Reply Qt__Key = 16777445

//
const Qt__Key_Reload Qt__Key = 16777446

//
const Qt__Key_RotateWindows Qt__Key = 16777447

//
const Qt__Key_RotationPB Qt__Key = 16777448

//
const Qt__Key_RotationKB Qt__Key = 16777449

//
const Qt__Key_Save Qt__Key = 16777450

//
const Qt__Key_Send Qt__Key = 16777451

//
const Qt__Key_Spell Qt__Key = 16777452

//
const Qt__Key_SplitScreen Qt__Key = 16777453

//
const Qt__Key_Support Qt__Key = 16777454

//
const Qt__Key_TaskPane Qt__Key = 16777455

//
const Qt__Key_Terminal Qt__Key = 16777456

//
const Qt__Key_Tools Qt__Key = 16777457

//
const Qt__Key_Travel Qt__Key = 16777458

//
const Qt__Key_Video Qt__Key = 16777459

//
const Qt__Key_Word Qt__Key = 16777460

//
const Qt__Key_Xfer Qt__Key = 16777461

//
const Qt__Key_ZoomIn Qt__Key = 16777462

//
const Qt__Key_ZoomOut Qt__Key = 16777463

//
const Qt__Key_Away Qt__Key = 16777464

//
const Qt__Key_Messenger Qt__Key = 16777465

//
const Qt__Key_WebCam Qt__Key = 16777466

//
const Qt__Key_MailForward Qt__Key = 16777467

//
const Qt__Key_Pictures Qt__Key = 16777468

//
const Qt__Key_Music Qt__Key = 16777469

//
const Qt__Key_Battery Qt__Key = 16777470

//
const Qt__Key_Bluetooth Qt__Key = 16777471

//
const Qt__Key_WLAN Qt__Key = 16777472

//
const Qt__Key_UWB Qt__Key = 16777473

//
const Qt__Key_AudioForward Qt__Key = 16777474

//
const Qt__Key_AudioRepeat Qt__Key = 16777475

//
const Qt__Key_AudioRandomPlay Qt__Key = 16777476

//
const Qt__Key_Subtitle Qt__Key = 16777477

//
const Qt__Key_AudioCycleTrack Qt__Key = 16777478

//
const Qt__Key_Time Qt__Key = 16777479

//
const Qt__Key_Hibernate Qt__Key = 16777480

//
const Qt__Key_View Qt__Key = 16777481

//
const Qt__Key_TopMenu Qt__Key = 16777482

//
const Qt__Key_PowerDown Qt__Key = 16777483

//
const Qt__Key_Suspend Qt__Key = 16777484

//
const Qt__Key_ContrastAdjust Qt__Key = 16777485

//
const Qt__Key_LaunchG Qt__Key = 16777486

//
const Qt__Key_LaunchH Qt__Key = 16777487

//
const Qt__Key_TouchpadToggle Qt__Key = 16777488

//
const Qt__Key_TouchpadOn Qt__Key = 16777489

//
const Qt__Key_TouchpadOff Qt__Key = 16777490

//
const Qt__Key_MicMute Qt__Key = 16777491

//
const Qt__Key_Red Qt__Key = 16777492

//
const Qt__Key_Green Qt__Key = 16777493

//
const Qt__Key_Yellow Qt__Key = 16777494

//
const Qt__Key_Blue Qt__Key = 16777495

//
const Qt__Key_ChannelUp Qt__Key = 16777496

//
const Qt__Key_ChannelDown Qt__Key = 16777497

//
const Qt__Key_Guide Qt__Key = 16777498

//
const Qt__Key_Info Qt__Key = 16777499

//
const Qt__Key_Settings Qt__Key = 16777500

//
const Qt__Key_MicVolumeUp Qt__Key = 16777501

//
const Qt__Key_MicVolumeDown Qt__Key = 16777502

//
const Qt__Key_New Qt__Key = 16777504

//
const Qt__Key_Open Qt__Key = 16777505

//
const Qt__Key_Find Qt__Key = 16777506

//
const Qt__Key_Undo Qt__Key = 16777507

//
const Qt__Key_Redo Qt__Key = 16777508

//
const Qt__Key_MediaLast Qt__Key = 16842751

//
const Qt__Key_Select Qt__Key = 16842752

//
const Qt__Key_Yes Qt__Key = 16842753

//
const Qt__Key_No Qt__Key = 16842754

//
const Qt__Key_Cancel Qt__Key = 16908289

//
const Qt__Key_Printer Qt__Key = 16908290

//
const Qt__Key_Execute Qt__Key = 16908291

//
const Qt__Key_Sleep Qt__Key = 16908292

//
const Qt__Key_Play Qt__Key = 16908293

//
const Qt__Key_Zoom Qt__Key = 16908294

//
const Qt__Key_Exit Qt__Key = 16908298

//
const Qt__Key_Context1 Qt__Key = 17825792

//
const Qt__Key_Context2 Qt__Key = 17825793

//
const Qt__Key_Context3 Qt__Key = 17825794

//
const Qt__Key_Context4 Qt__Key = 17825795

//
const Qt__Key_Call Qt__Key = 17825796

//
const Qt__Key_Hangup Qt__Key = 17825797

//
const Qt__Key_Flip Qt__Key = 17825798

//
const Qt__Key_ToggleCallHangup Qt__Key = 17825799

//
const Qt__Key_VoiceDial Qt__Key = 17825800

//
const Qt__Key_LastNumberRedial Qt__Key = 17825801

//
const Qt__Key_Camera Qt__Key = 17825824

//
const Qt__Key_CameraFocus Qt__Key = 17825825

//
const Qt__Key_unknown Qt__Key = 33554431

/*
ConstantValue
Qt::NoArrow0
Qt::UpArrow1
Qt::DownArrow2
Qt::LeftArrow3
Qt::RightArrow4

*/
type Qt__ArrowType = int // core
//
const Qt__NoArrow Qt__ArrowType = 0

//
const Qt__UpArrow Qt__ArrowType = 1

//
const Qt__DownArrow Qt__ArrowType = 2

//
const Qt__LeftArrow Qt__ArrowType = 3

//
const Qt__RightArrow Qt__ArrowType = 4

/*
This enum type defines the pen styles that can be drawn using QPainter. The styles are:



Qt::SolidLineQt::DashLineQt::DotLine

Qt::DashDotLineQt::DashDotDotLineQt::CustomDashLine




See also QPen.

*/
type Qt__PenStyle = int // core
// no line at all. For example, QPainter::drawRect() fills but does not draw any boundary line.
const Qt__NoPen Qt__PenStyle = 0

// A plain line.
const Qt__SolidLine Qt__PenStyle = 1

// Dashes separated by a few pixels.
const Qt__DashLine Qt__PenStyle = 2

// Dots separated by a few pixels.
const Qt__DotLine Qt__PenStyle = 3

// Alternate dots and dashes.
const Qt__DashDotLine Qt__PenStyle = 4

// One dash, two dots, one dash, two dots.
const Qt__DashDotDotLine Qt__PenStyle = 5

// A custom pattern defined using QPainterPathStroker::setDashPattern().
const Qt__CustomDashLine Qt__PenStyle = 6

//
const Qt__MPenStyle Qt__PenStyle = 15

/*
This enum type defines the pen cap styles supported by Qt, i.e. the line end caps that can be drawn using QPainter.



Qt::SquareCapQt::FlatCapQt::RoundCap




See also QPen.

*/
type Qt__PenCapStyle = int // core
//
const Qt__FlatCap Qt__PenCapStyle = 0

//
const Qt__SquareCap Qt__PenCapStyle = 16

//
const Qt__RoundCap Qt__PenCapStyle = 32

//
const Qt__MPenCapStyle Qt__PenCapStyle = 48

/*
This enum type defines the pen join styles supported by Qt, i.e. which joins between two connected lines can be drawn using QPainter.



Qt::BevelJoinQt::MiterJoinQt::RoundJoin




See also QPen.

*/
type Qt__PenJoinStyle = int // core
//
const Qt__MiterJoin Qt__PenJoinStyle = 0

//
const Qt__BevelJoin Qt__PenJoinStyle = 64

//
const Qt__RoundJoin Qt__PenJoinStyle = 128

//
const Qt__SvgMiterJoin Qt__PenJoinStyle = 256

//
const Qt__MPenJoinStyle Qt__PenJoinStyle = 448

/*
This enum type defines the brush styles supported by Qt, i.e. the fill pattern of shapes drawn using QPainter.





See also QBrush.

*/
type Qt__BrushStyle = int // core
// No brush pattern.
const Qt__NoBrush Qt__BrushStyle = 0

// Uniform color.
const Qt__SolidPattern Qt__BrushStyle = 1

// Extremely dense brush pattern.
const Qt__Dense1Pattern Qt__BrushStyle = 2

// Very dense brush pattern.
const Qt__Dense2Pattern Qt__BrushStyle = 3

// Somewhat dense brush pattern.
const Qt__Dense3Pattern Qt__BrushStyle = 4

// Half dense brush pattern.
const Qt__Dense4Pattern Qt__BrushStyle = 5

// Somewhat sparse brush pattern.
const Qt__Dense5Pattern Qt__BrushStyle = 6

// Very sparse brush pattern.
const Qt__Dense6Pattern Qt__BrushStyle = 7

// Extremely sparse brush pattern.
const Qt__Dense7Pattern Qt__BrushStyle = 8

// Horizontal lines.
const Qt__HorPattern Qt__BrushStyle = 9

//
const Qt__VerPattern Qt__BrushStyle = 10

//
const Qt__CrossPattern Qt__BrushStyle = 11

//
const Qt__BDiagPattern Qt__BrushStyle = 12

//
const Qt__FDiagPattern Qt__BrushStyle = 13

//
const Qt__DiagCrossPattern Qt__BrushStyle = 14

//
const Qt__LinearGradientPattern Qt__BrushStyle = 15

//
const Qt__RadialGradientPattern Qt__BrushStyle = 16

//
const Qt__ConicalGradientPattern Qt__BrushStyle = 17

//
const Qt__TexturePattern Qt__BrushStyle = 24

/*
This enum is used by QPainter::drawRoundedRect() and QPainterPath::addRoundedRect() functions to specify the radii of rectangle corners with respect to the dimensions of the bounding rectangles specified.



This enum was introduced or modified in  Qt 4.4.

*/
type Qt__SizeMode = int // core
// Specifies the size using absolute measurements.
const Qt__AbsoluteSize Qt__SizeMode = 0

// Specifies the size relative to the bounding rectangle, typically using percentage measurements.
const Qt__RelativeSize Qt__SizeMode = 1

/*
This enum describes the available UI effects.

By default, Qt will try to use the platform specific desktop settings for each effect. Use the QApplication::setDesktopSettingsAware() function (passing false as argument) to prevent this, and the QApplication::setEffectEnabled() to enable or disable a particular effect.

Note that all effects are disabled on screens running at less than 16-bit color depth.



See also QApplication::setEffectEnabled() and QGuiApplication::setDesktopSettingsAware().

*/
type Qt__UIEffect = int // core
//
const Qt__UI_General Qt__UIEffect = 0

// Show animated menus.
const Qt__UI_AnimateMenu Qt__UIEffect = 1

// Show faded menus.
const Qt__UI_FadeMenu Qt__UIEffect = 2

// Show animated comboboxes.
const Qt__UI_AnimateCombo Qt__UIEffect = 3

// Show tooltip animations.
const Qt__UI_AnimateTooltip Qt__UIEffect = 4

// Show tooltip fading effects.
const Qt__UI_FadeTooltip Qt__UIEffect = 5

// Reserved
const Qt__UI_AnimateToolBox Qt__UIEffect = 6

/*
This enum type defines the various cursors that can be used.

The standard arrow cursor is the default for widgets in a normal state.


*/
type Qt__CursorShape = int // core
//  The standard arrow cursor.
const Qt__ArrowCursor Qt__CursorShape = 0

//  An arrow pointing upwards toward the top of the screen.
const Qt__UpArrowCursor Qt__CursorShape = 1

//  A crosshair cursor, typically used to help the user accurately select a point on the screen.
const Qt__CrossCursor Qt__CursorShape = 2

//  An hourglass or watch cursor, usually shown during operations that prevent the user from interacting with the application.
const Qt__WaitCursor Qt__CursorShape = 3

//  A caret or ibeam cursor, indicating that a widget can accept and display text input.
const Qt__IBeamCursor Qt__CursorShape = 4

//  A cursor used for elements that are used to vertically resize top-level windows.
const Qt__SizeVerCursor Qt__CursorShape = 5

//  A cursor used for elements that are used to horizontally resize top-level windows.
const Qt__SizeHorCursor Qt__CursorShape = 6

//  A cursor used for elements that are used to diagonally resize top-level windows at their top-right and bottom-left corners.
const Qt__SizeBDiagCursor Qt__CursorShape = 7

//  A cursor used for elements that are used to diagonally resize top-level windows at their top-left and bottom-right corners.
const Qt__SizeFDiagCursor Qt__CursorShape = 8

//  A cursor used for elements that are used to resize top-level windows in any direction.
const Qt__SizeAllCursor Qt__CursorShape = 9

//
const Qt__BlankCursor Qt__CursorShape = 10

//
const Qt__SplitVCursor Qt__CursorShape = 11

//
const Qt__SplitHCursor Qt__CursorShape = 12

//
const Qt__PointingHandCursor Qt__CursorShape = 13

//
const Qt__ForbiddenCursor Qt__CursorShape = 14

//
const Qt__WhatsThisCursor Qt__CursorShape = 15

//
const Qt__BusyCursor Qt__CursorShape = 16

//
const Qt__OpenHandCursor Qt__CursorShape = 17

//
const Qt__ClosedHandCursor Qt__CursorShape = 18

//
const Qt__DragCopyCursor Qt__CursorShape = 19

//
const Qt__DragMoveCursor Qt__CursorShape = 20

//
const Qt__DragLinkCursor Qt__CursorShape = 21

//
const Qt__LastCursor Qt__CursorShape = 21

// 4
const Qt__BitmapCursor Qt__CursorShape = 24

//
const Qt__CustomCursor Qt__CursorShape = 25

/*
This enum is used in widgets that can display both plain text and rich text, for example QLabel. It is used for deciding whether a text string should be interpreted as one or the other. This is normally done by passing one of the enum values to a QTextEdit::setTextFormat() function.


*/
type Qt__TextFormat = int // core
// The text string is interpreted as a plain text string.
const Qt__PlainText Qt__TextFormat = 0

// The text string is interpreted as a rich text string. See Supported HTML Subset for the definition of rich text.
const Qt__RichText Qt__TextFormat = 1

// The text string is interpreted as for Qt::RichText if Qt::mightBeRichText() returns true, otherwise as Qt::PlainText.
const Qt__AutoText Qt__TextFormat = 2

/*
This enum type defines what happens to the aspect ratio when scaling an rectangle.





See also QSize::scale() and QImage::scaled().

*/
type Qt__AspectRatioMode = int // core
// The size is scaled freely. The aspect ratio is not preserved.
const Qt__IgnoreAspectRatio Qt__AspectRatioMode = 0

// The size is scaled to a rectangle as large as possible inside a given rectangle, preserving the aspect ratio.
const Qt__KeepAspectRatio Qt__AspectRatioMode = 1

// The size is scaled to a rectangle as small as possible outside a given rectangle, preserving the aspect ratio.
const Qt__KeepAspectRatioByExpanding Qt__AspectRatioMode = 2

/*


 */
type Qt__DockWidgetArea = int // core
//
const Qt__LeftDockWidgetArea Qt__DockWidgetArea = 1

//
const Qt__RightDockWidgetArea Qt__DockWidgetArea = 2

//
const Qt__TopDockWidgetArea Qt__DockWidgetArea = 4

//
const Qt__BottomDockWidgetArea Qt__DockWidgetArea = 8

//
const Qt__DockWidgetArea_Mask Qt__DockWidgetArea = 15

//
const Qt__AllDockWidgetAreas Qt__DockWidgetArea = 15

//
const Qt__NoDockWidgetArea Qt__DockWidgetArea = 0

/*


 */
type Qt__DockWidgetAreaSizes = int // core
//
const Qt__NDockWidgetAreas Qt__DockWidgetAreaSizes = 4

/*


 */
type Qt__ToolBarArea = int // core
//
const Qt__LeftToolBarArea Qt__ToolBarArea = 1

//
const Qt__RightToolBarArea Qt__ToolBarArea = 2

//
const Qt__TopToolBarArea Qt__ToolBarArea = 4

//
const Qt__BottomToolBarArea Qt__ToolBarArea = 8

//
const Qt__ToolBarArea_Mask Qt__ToolBarArea = 15

//
const Qt__AllToolBarAreas Qt__ToolBarArea = 15

//
const Qt__NoToolBarArea Qt__ToolBarArea = 0

/*


 */
type Qt__ToolBarAreaSizes = int // core
//
const Qt__NToolBarAreas Qt__ToolBarAreaSizes = 4

/*
Qt::SystemLocaleShortDate?The short format used by the operating system.
Qt::SystemLocaleLongDate?The long format used by the operating system.
Qt::DefaultLocaleShortDate?The short format specified by the application's locale.
Qt::DefaultLocaleLongDate?The long format used by the application's locale.
Qt::LocaleDate?This enum value is deprecated. Use Qt::DefaultLocaleShortDate instead (or Qt::DefaultLocaleLongDate if you want long dates).
Qt::LocalDateSystemLocaleDateThis enum value is deprecated. Use Qt::SystemLocaleShortDate instead (or Qt::SystemLocaleLongDate if you want long dates).


Note: For ISODate formats, each Y, M and D represents a single digit of the year, month and day used to specify the date. Each H, M and S represents a single digit of the hour, minute and second used to specify the time. The presence of a literal T character is used to separate the date and time when both are specified.

*/
type Qt__DateFormat = int // core
// The default Qt format, which includes the day and month name, the day number in the month, and the year in full. The day and month names will be short, localized names. This is basically equivalent to using the date format string, "ddd MMM d yyyy". See QDate::toString() for more information.
const Qt__TextDate Qt__DateFormat = 0

//
const Qt__ISODate Qt__DateFormat = 1

// This enum value is deprecated. Use Qt::SystemLocaleShortDate instead (or Qt::SystemLocaleLongDate if you want long dates).
const Qt__SystemLocaleDate Qt__DateFormat = 2

//
const Qt__LocalDate Qt__DateFormat = 2

//
const Qt__LocaleDate Qt__DateFormat = 3

//
const Qt__SystemLocaleShortDate Qt__DateFormat = 4

//
const Qt__SystemLocaleLongDate Qt__DateFormat = 5

//
const Qt__DefaultLocaleShortDate Qt__DateFormat = 6

//
const Qt__DefaultLocaleLongDate Qt__DateFormat = 7

//
const Qt__RFC2822Date Qt__DateFormat = 8

//
const Qt__ISODateWithMs Qt__DateFormat = 9

/*

 */
type Qt__TimeSpec = int // core
// Locale dependent time (Timezones and Daylight Savings Time).
const Qt__LocalTime Qt__TimeSpec = 0

// Coordinated Universal Time, replaces Greenwich Mean Time.
const Qt__UTC Qt__TimeSpec = 1

// An offset in seconds from Coordinated Universal Time.
const Qt__OffsetFromUTC Qt__TimeSpec = 2

// A named time zone using a specific set of Daylight Savings rules.
const Qt__TimeZone Qt__TimeSpec = 3

/*
ConstantValue
Qt::Monday1
Qt::Tuesday2
Qt::Wednesday3
Qt::Thursday4
Qt::Friday5
Qt::Saturday6
Qt::Sunday7

*/
type Qt__DayOfWeek = int // core
//
const Qt__Monday Qt__DayOfWeek = 1

//
const Qt__Tuesday Qt__DayOfWeek = 2

//
const Qt__Wednesday Qt__DayOfWeek = 3

//
const Qt__Thursday Qt__DayOfWeek = 4

//
const Qt__Friday Qt__DayOfWeek = 5

//
const Qt__Saturday Qt__DayOfWeek = 6

//
const Qt__Sunday Qt__DayOfWeek = 7

/*
This enum type describes the various modes of QAbstractScrollArea's scroll bars.



(The modes for the horizontal and vertical scroll bars are independent.)

*/
type Qt__ScrollBarPolicy = int // core
// QAbstractScrollArea shows a scroll bar when the content is too large to fit and not otherwise. This is the default.
const Qt__ScrollBarAsNeeded Qt__ScrollBarPolicy = 0

// QAbstractScrollArea never shows a scroll bar.
const Qt__ScrollBarAlwaysOff Qt__ScrollBarPolicy = 1

//
const Qt__ScrollBarAlwaysOn Qt__ScrollBarPolicy = 2

/*
ConstantValue
Qt::CaseInsensitive0
Qt::CaseSensitive1

*/
type Qt__CaseSensitivity = int // core
//
const Qt__CaseInsensitive Qt__CaseSensitivity = 0

//
const Qt__CaseSensitive Qt__CaseSensitivity = 1

/*
This enum type specifies a corner in a rectangle:


*/
type Qt__Corner = int // core
//
const Qt__TopLeftCorner Qt__Corner = 0

//
const Qt__TopRightCorner Qt__Corner = 1

//
const Qt__BottomLeftCorner Qt__Corner = 2

//
const Qt__BottomRightCorner Qt__Corner = 3

/*


 */
type Qt__Edge = int // core
//
const Qt__TopEdge Qt__Edge = 1

//
const Qt__LeftEdge Qt__Edge = 2

//
const Qt__RightEdge Qt__Edge = 4

//
const Qt__BottomEdge Qt__Edge = 8

/*
This enum describes the types of connection that can be used between signals and slots. In particular, it determines whether a particular signal is delivered to a slot immediately or queued for delivery at a later time.



With queued connections, the parameters must be of types that are known to Qt's meta-object system, because Qt needs to copy the arguments to store them in an event behind the scenes. If you try to use a queued connection and get the error message:


  QObject::connect: Cannot queue arguments of type 'MyType'



Call qRegisterMetaType() to register the data type before you establish the connection.

When using signals and slots with multiple threads, see Signals and Slots Across Threads.

See also Thread Support in Qt, QObject::connect(), qRegisterMetaType(), and Q_DECLARE_METATYPE().

*/
type Qt__ConnectionType = int // core
// (Default) If the receiver lives in the thread that emits the signal, Qt::DirectConnection is used. Otherwise, Qt::QueuedConnection is used. The connection type is determined when the signal is emitted.
const Qt__AutoConnection Qt__ConnectionType = 0

// The slot is invoked immediately when the signal is emitted. The slot is executed in the signalling thread.
const Qt__DirectConnection Qt__ConnectionType = 1

// The slot is invoked when control returns to the event loop of the receiver's thread. The slot is executed in the receiver's thread.
const Qt__QueuedConnection Qt__ConnectionType = 2

// Same as Qt::QueuedConnection, except that the signalling thread blocks until the slot returns. This connection must not be used if the receiver lives in the signalling thread, or else the application will deadlock.
const Qt__BlockingQueuedConnection Qt__ConnectionType = 3

//
const Qt__UniqueConnection Qt__ConnectionType = 128

/*
For a QEvent::Shortcut event to occur, the shortcut's key sequence must be entered by the user in a context where the shortcut is active. The possible contexts are these:


*/
type Qt__ShortcutContext = int // core
// The shortcut is active when its parent widget has focus.
const Qt__WidgetShortcut Qt__ShortcutContext = 0

// The shortcut is active when its parent widget is a logical subwidget of the active top-level window.
const Qt__WindowShortcut Qt__ShortcutContext = 1

// The shortcut is active when one of the applications windows are active.
const Qt__ApplicationShortcut Qt__ShortcutContext = 2

// The shortcut is active when its parent widget, or any of its children has focus. Children which are top-level widgets, except pop-ups, are not affected by this shortcut context.
const Qt__WidgetWithChildrenShortcut Qt__ShortcutContext = 3

/*
Specifies which method should be used to fill the paths and polygons.


*/
type Qt__FillRule = int // core
// Specifies that the region is filled using the odd even fill rule. With this rule, we determine whether a point is inside the shape by using the following method. Draw a horizontal line from the point to a location outside the shape, and count the number of intersections. If the number of intersections is an odd number, the point is inside the shape. This mode is the default.
const Qt__OddEvenFill Qt__FillRule = 0

// Specifies that the region is filled using the non zero winding rule. With this rule, we determine whether a point is inside the shape by using the following method. Draw a horizontal line from the point to a location outside the shape. Determine whether the direction of the line at each intersection point is up or down. The winding number is determined by summing the direction of each intersection. If the number is non zero, the point is inside the shape. This fill mode can also in most cases be considered as the intersection of closed shapes.
const Qt__WindingFill Qt__FillRule = 1

/*
This enum specifies the behavior of the QPixmap::createMaskFromColor() and QImage::createMaskFromColor() functions.


*/
type Qt__MaskMode = int // core
// Creates a mask where all pixels matching the given color are opaque.
const Qt__MaskInColor Qt__MaskMode = 0

// Creates a mask where all pixels matching the given color are transparent.
const Qt__MaskOutColor Qt__MaskMode = 1

/*

 */
type Qt__ClipOperation = int // core
// This operation turns clipping off.
const Qt__NoClip Qt__ClipOperation = 0

// Replaces the current clip path/rect/region with the one supplied in the function call.
const Qt__ReplaceClip Qt__ClipOperation = 1

// Intersects the current clip path/rect/region with the one supplied in the function call.
const Qt__IntersectClip Qt__ClipOperation = 2

/*
This enum is used in QGraphicsItem, QGraphicsScene and QGraphicsView to specify how items are selected, or how to determine if shapes and items collide.



See also QGraphicsScene::items(), QGraphicsScene::collidingItems(), QGraphicsView::items(), QGraphicsItem::collidesWithItem(), and QGraphicsItem::collidesWithPath().

*/
type Qt__ItemSelectionMode = int // core
//
const Qt__ContainsItemShape Qt__ItemSelectionMode = 0

//
const Qt__IntersectsItemShape Qt__ItemSelectionMode = 1

//
const Qt__ContainsItemBoundingRect Qt__ItemSelectionMode = 2

//
const Qt__IntersectsItemBoundingRect Qt__ItemSelectionMode = 3

/*
This enum is used in QGraphicsScene to specify what to do with currently selected items when setting a selection area.



See also QGraphicsScene::setSelectionArea().

*/
type Qt__ItemSelectionOperation = int // core
// The currently selected items are replaced by items in the selection area.
const Qt__ReplaceSelection Qt__ItemSelectionOperation = 0

// The items in the selection area are added to the currently selected items.
const Qt__AddToSelection Qt__ItemSelectionOperation = 1

/*
This enum type defines whether image transformations (e.g., scaling) should be smooth or not.



See also QImage::scaled().

*/
type Qt__TransformationMode = int // core
// The transformation is performed quickly, with no smoothing.
const Qt__FastTransformation Qt__TransformationMode = 0

// The resulting image is transformed using bilinear filtering.
const Qt__SmoothTransformation Qt__TransformationMode = 1

/*
This enum type defines three values to represent the three axes in the cartesian coordinate system.



See also QTransform::rotate() and QTransform::rotateRadians().

*/
type Qt__Axis = int // core
// The X axis.
const Qt__XAxis Qt__Axis = 0

// The Y axis.
const Qt__YAxis Qt__Axis = 1

// The Z axis.
const Qt__ZAxis Qt__Axis = 2

/*
This enum specifies why the focus changed. It will be passed through QWidget::setFocus and can be retrieved in the QFocusEvent sent to the widget upon focus change.



See also Keyboard Focus in Widgets.

*/
type Qt__FocusReason = int // core
// A mouse action occurred.
const Qt__MouseFocusReason Qt__FocusReason = 0

// The Tab key was pressed.
const Qt__TabFocusReason Qt__FocusReason = 1

// A Backtab occurred. The input for this may include the Shift or Control keys; e.g. Shift+Tab.
const Qt__BacktabFocusReason Qt__FocusReason = 2

// The window system made this window either active or inactive.
const Qt__ActiveWindowFocusReason Qt__FocusReason = 3

// The application opened/closed a pop-up that grabbed/released the keyboard focus.
const Qt__PopupFocusReason Qt__FocusReason = 4

// The user typed a label's buddy shortcut
const Qt__ShortcutFocusReason Qt__FocusReason = 5

// The menu bar took focus.
const Qt__MenuBarFocusReason Qt__FocusReason = 6

// Another reason, usually application-specific.
const Qt__OtherFocusReason Qt__FocusReason = 7

//
const Qt__NoFocusReason Qt__FocusReason = 8

/*
This enum type defines the various policies a widget can have with respect to showing a context menu.


*/
type Qt__ContextMenuPolicy = int // core
// the widget does not feature a context menu, context menu handling is deferred to the widget's parent.
const Qt__NoContextMenu Qt__ContextMenuPolicy = 0

// the widget's QWidget::contextMenuEvent() handler is called.
const Qt__DefaultContextMenu Qt__ContextMenuPolicy = 1

// the widget displays its QWidget::actions() as context menu.
const Qt__ActionsContextMenu Qt__ContextMenuPolicy = 2

// the widget emits the QWidget::customContextMenuRequested() signal.
const Qt__CustomContextMenu Qt__ContextMenuPolicy = 3

// the widget does not feature a context menu, and in contrast to NoContextMenu, the handling is not deferred to the widget's parent. This means that all right mouse button events are guaranteed to be delivered to the widget itself through QWidget::mousePressEvent(), and QWidget::mouseReleaseEvent().
const Qt__PreventContextMenu Qt__ContextMenuPolicy = 4

/*


 */
type Qt__InputMethodQuery = int // core
//
const Qt__ImEnabled Qt__InputMethodQuery = 1

//
const Qt__ImCursorRectangle Qt__InputMethodQuery = 2

//
const Qt__ImMicroFocus Qt__InputMethodQuery = 2

//
const Qt__ImFont Qt__InputMethodQuery = 4

//
const Qt__ImCursorPosition Qt__InputMethodQuery = 8

//
const Qt__ImSurroundingText Qt__InputMethodQuery = 16

//
const Qt__ImCurrentSelection Qt__InputMethodQuery = 32

//
const Qt__ImMaximumTextLength Qt__InputMethodQuery = 64

//
const Qt__ImAnchorPosition Qt__InputMethodQuery = 128

//
const Qt__ImHints Qt__InputMethodQuery = 256

//
const Qt__ImPreferredLanguage Qt__InputMethodQuery = 512

//
const Qt__ImAbsolutePosition Qt__InputMethodQuery = 1024

//
const Qt__ImTextBeforeCursor Qt__InputMethodQuery = 2048

//
const Qt__ImTextAfterCursor Qt__InputMethodQuery = 4096

//
const Qt__ImEnterKeyType Qt__InputMethodQuery = 8192

//
const Qt__ImAnchorRectangle Qt__InputMethodQuery = 16384

//
const Qt__ImInputItemClipRectangle Qt__InputMethodQuery = 32768

//
const Qt__ImPlatformData Qt__InputMethodQuery = -2147483648

//
const Qt__ImQueryInput Qt__InputMethodQuery = 16570

//
const Qt__ImQueryAll Qt__InputMethodQuery = -1

/*


 */
type Qt__InputMethodHint = int // core
//
const Qt__ImhNone Qt__InputMethodHint = 0

//
const Qt__ImhHiddenText Qt__InputMethodHint = 1

//
const Qt__ImhSensitiveData Qt__InputMethodHint = 2

//
const Qt__ImhNoAutoUppercase Qt__InputMethodHint = 4

//
const Qt__ImhPreferNumbers Qt__InputMethodHint = 8

//
const Qt__ImhPreferUppercase Qt__InputMethodHint = 16

//
const Qt__ImhPreferLowercase Qt__InputMethodHint = 32

//
const Qt__ImhNoPredictiveText Qt__InputMethodHint = 64

//
const Qt__ImhDate Qt__InputMethodHint = 128

//
const Qt__ImhTime Qt__InputMethodHint = 256

//
const Qt__ImhPreferLatin Qt__InputMethodHint = 512

//
const Qt__ImhMultiLine Qt__InputMethodHint = 1024

//
const Qt__ImhDigitsOnly Qt__InputMethodHint = 65536

//
const Qt__ImhFormattedNumbersOnly Qt__InputMethodHint = 131072

//
const Qt__ImhUppercaseOnly Qt__InputMethodHint = 262144

//
const Qt__ImhLowercaseOnly Qt__InputMethodHint = 524288

//
const Qt__ImhDialableCharactersOnly Qt__InputMethodHint = 1048576

//
const Qt__ImhEmailCharactersOnly Qt__InputMethodHint = 2097152

//
const Qt__ImhUrlCharactersOnly Qt__InputMethodHint = 4194304

//
const Qt__ImhLatinOnly Qt__InputMethodHint = 8388608

//
const Qt__ImhExclusiveInputMask Qt__InputMethodHint = -65536

/*
This can be used to alter the appearance of the Return key on an on-screen keyboard.

Note: Not all of these values are supported on all platforms. For unsupported values the default key will be used instead.



This enum was introduced or modified in  Qt 5.6.

*/
type Qt__EnterKeyType = int // core
// The default Enter key. This can either be a button closing the keyboard, or a Return button causing a new line in case of a multi-line input field.
const Qt__EnterKeyDefault Qt__EnterKeyType = 0

// Show a Return button that inserts a new line. The keyboard will not close when this button is pressed.
const Qt__EnterKeyReturn Qt__EnterKeyType = 1

// Show a "Done" button. The keyboard will close when this button is pressed.
const Qt__EnterKeyDone Qt__EnterKeyType = 2

// Show a "Go" button. Typically used in an address bar when entering a URL; the keyboard will close when this button is pressed.
const Qt__EnterKeyGo Qt__EnterKeyType = 3

// Show a "Send" button. The keyboard will close when this button is pressed.
const Qt__EnterKeySend Qt__EnterKeyType = 4

// Show a "Search" button. The keyboard will close when this button is pressed.
const Qt__EnterKeySearch Qt__EnterKeyType = 5

// Show a "Next" button. Typically used in a form to allow navigating to the next input field; the keyboard will not close when this button is pressed.
const Qt__EnterKeyNext Qt__EnterKeyType = 6

// Show a "Previous" button. The keyboard will not close when this button is pressed.
const Qt__EnterKeyPrevious Qt__EnterKeyType = 7

/*
The style of the tool button, describing how the button's text and icon should be displayed.


*/
type Qt__ToolButtonStyle = int // core
// Only display the icon.
const Qt__ToolButtonIconOnly Qt__ToolButtonStyle = 0

// Only display the text.
const Qt__ToolButtonTextOnly Qt__ToolButtonStyle = 1

// The text appears beside the icon.
const Qt__ToolButtonTextBesideIcon Qt__ToolButtonStyle = 2

// The text appears under the icon.
const Qt__ToolButtonTextUnderIcon Qt__ToolButtonStyle = 3

// Follow the style.
const Qt__ToolButtonFollowStyle Qt__ToolButtonStyle = 4

/*
Specifies the direction of Qt's layouts and text handling.



Right-to-left layouts are necessary for certain languages, notably Arabic and Hebrew.

LayoutDirectionAuto serves two purposes. When used in conjunction with widgets and layouts, it will imply to use the layout direction set on the parent widget or QApplication. This has the same effect as QWidget::unsetLayoutDirection().

When LayoutDirectionAuto is used in conjunction with text layouting, it will imply that the text directionality is determined from the content of the string to be layouted.

See also QGuiApplication::setLayoutDirection(), QWidget::setLayoutDirection(), QTextOption::setTextDirection(), and QString::isRightToLeft().

*/
type Qt__LayoutDirection = int // core
// Left-to-right layout.
const Qt__LeftToRight Qt__LayoutDirection = 0

// Right-to-left layout.
const Qt__RightToLeft Qt__LayoutDirection = 1

// Automatic layout.
const Qt__LayoutDirectionAuto Qt__LayoutDirection = 2

/*
Specifies a side of a layout item that can be anchored. This is used by QGraphicsAnchorLayout.



See also QGraphicsAnchorLayout.

*/
type Qt__AnchorPoint = int // core
// The left side of a layout item.
const Qt__AnchorLeft Qt__AnchorPoint = 0

// A "virtual" side that is centered between the left and the right side of a layout item.
const Qt__AnchorHorizontalCenter Qt__AnchorPoint = 1

// The right side of a layout item.
const Qt__AnchorRight Qt__AnchorPoint = 2

// The top side of a layout item.
const Qt__AnchorTop Qt__AnchorPoint = 3

// A "virtual" side that is centered between the top and the bottom side of a layout item.
const Qt__AnchorVerticalCenter Qt__AnchorPoint = 4

// The bottom side of a layout item.
const Qt__AnchorBottom Qt__AnchorPoint = 5

/*


 */
type Qt__FindChildOption = int // core
//
const Qt__FindDirectChildrenOnly Qt__FindChildOption = 0

//
const Qt__FindChildrenRecursively Qt__FindChildOption = 1

/*


 */
type Qt__DropAction = int // core
//
const Qt__CopyAction Qt__DropAction = 1

//
const Qt__MoveAction Qt__DropAction = 2

//
const Qt__LinkAction Qt__DropAction = 4

//
const Qt__ActionMask Qt__DropAction = 255

//
const Qt__TargetMoveAction Qt__DropAction = 32770

//
const Qt__IgnoreAction Qt__DropAction = 0

/*
This enum describes the state of checkable items, controls, and widgets.



See also QCheckBox, Qt::ItemFlags, and Qt::ItemDataRole.

*/
type Qt__CheckState = int // core
// The item is unchecked.
const Qt__Unchecked Qt__CheckState = 0

// The item is partially checked. Items in hierarchical models may be partially checked if some, but not all, of their children are checked.
const Qt__PartiallyChecked Qt__CheckState = 1

// The item is checked.
const Qt__Checked Qt__CheckState = 2

/*
Each item in the model has a set of data elements associated with it, each with its own role. The roles are used by the view to indicate to the model which type of data it needs. Custom models should return data in these types.

The general purpose roles (and the associated types) are:



Roles describing appearance and meta data (with associated types):



Accessibility roles (with associated types):



User roles:



For user roles, it is up to the developer to decide which types to use and ensure that components use the correct types when accessing and setting data.

*/
type Qt__ItemDataRole = int // core
// The key data to be rendered in the form of text. (QString)
const Qt__DisplayRole Qt__ItemDataRole = 0

// The data to be rendered as a decoration in the form of an icon. (QColor, QIcon or QPixmap)
const Qt__DecorationRole Qt__ItemDataRole = 1

// The data in a form suitable for editing in an editor. (QString)
const Qt__EditRole Qt__ItemDataRole = 2

// The data displayed in the item's tooltip. (QString)
const Qt__ToolTipRole Qt__ItemDataRole = 3

// The data displayed in the status bar. (QString)
const Qt__StatusTipRole Qt__ItemDataRole = 4

// The data displayed for the item in "What's This?" mode. (QString)
const Qt__WhatsThisRole Qt__ItemDataRole = 5

// The font used for items rendered with the default delegate. (QFont)
const Qt__FontRole Qt__ItemDataRole = 6

// The alignment of the text for items rendered with the default delegate. (Qt::Alignment)
const Qt__TextAlignmentRole Qt__ItemDataRole = 7

// This role is obsolete. Use BackgroundRole instead.
const Qt__BackgroundColorRole Qt__ItemDataRole = 8

// The background brush used for items rendered with the default delegate. (QBrush)
const Qt__BackgroundRole Qt__ItemDataRole = 8

// This role is obsolete. Use ForegroundRole instead.
const Qt__TextColorRole Qt__ItemDataRole = 9

// The foreground brush (text color, typically) used for items rendered with the default delegate. (QBrush)
const Qt__ForegroundRole Qt__ItemDataRole = 9

//
const Qt__CheckStateRole Qt__ItemDataRole = 10

//
const Qt__AccessibleTextRole Qt__ItemDataRole = 11

//
const Qt__AccessibleDescriptionRole Qt__ItemDataRole = 12

//
const Qt__SizeHintRole Qt__ItemDataRole = 13

//
const Qt__InitialSortOrderRole Qt__ItemDataRole = 14

//
const Qt__DisplayPropertyRole Qt__ItemDataRole = 27

//
const Qt__DecorationPropertyRole Qt__ItemDataRole = 28

//
const Qt__ToolTipPropertyRole Qt__ItemDataRole = 29

//
const Qt__StatusTipPropertyRole Qt__ItemDataRole = 30

//
const Qt__WhatsThisPropertyRole Qt__ItemDataRole = 31

//
const Qt__UserRole Qt__ItemDataRole = 256

/*


 */
type Qt__ItemFlag = int // core
//
const Qt__NoItemFlags Qt__ItemFlag = 0

//
const Qt__ItemIsSelectable Qt__ItemFlag = 1

//
const Qt__ItemIsEditable Qt__ItemFlag = 2

//
const Qt__ItemIsDragEnabled Qt__ItemFlag = 4

//
const Qt__ItemIsDropEnabled Qt__ItemFlag = 8

//
const Qt__ItemIsUserCheckable Qt__ItemFlag = 16

//
const Qt__ItemIsEnabled Qt__ItemFlag = 32

//
const Qt__ItemIsAutoTristate Qt__ItemFlag = 64

//
const Qt__ItemIsTristate Qt__ItemFlag = 64

//
const Qt__ItemNeverHasChildren Qt__ItemFlag = 128

//
const Qt__ItemIsUserTristate Qt__ItemFlag = 256

/*


 */
type Qt__MatchFlag = int // core
//
const Qt__MatchExactly Qt__MatchFlag = 0

//
const Qt__MatchContains Qt__MatchFlag = 1

//
const Qt__MatchStartsWith Qt__MatchFlag = 2

//
const Qt__MatchEndsWith Qt__MatchFlag = 3

//
const Qt__MatchRegExp Qt__MatchFlag = 4

//
const Qt__MatchWildcard Qt__MatchFlag = 5

//
const Qt__MatchFixedString Qt__MatchFlag = 8

//
const Qt__MatchCaseSensitive Qt__MatchFlag = 16

//
const Qt__MatchWrap Qt__MatchFlag = 32

//
const Qt__MatchRecursive Qt__MatchFlag = 64

/*
This enum specifies the behavior of a modal window. A modal window is one that blocks input to other windows. Note that windows that are children of a modal window are not blocked.

The values are:



See also QWidget::windowModality and QDialog.

*/
type Qt__WindowModality = int // core
// The window is not modal and does not block input to other windows.
const Qt__NonModal Qt__WindowModality = 0

// The window is modal to a single window hierarchy and blocks input to its parent window, all grandparent windows, and all siblings of its parent and grandparent windows.
const Qt__WindowModal Qt__WindowModality = 1

// The window is modal to the application and blocks input to all windows.
const Qt__ApplicationModal Qt__WindowModality = 2

/*


 */
type Qt__TextInteractionFlag = int // core
//
const Qt__NoTextInteraction Qt__TextInteractionFlag = 0

//
const Qt__TextSelectableByMouse Qt__TextInteractionFlag = 1

//
const Qt__TextSelectableByKeyboard Qt__TextInteractionFlag = 2

//
const Qt__LinksAccessibleByMouse Qt__TextInteractionFlag = 4

//
const Qt__LinksAccessibleByKeyboard Qt__TextInteractionFlag = 8

//
const Qt__TextEditable Qt__TextInteractionFlag = 16

//
const Qt__TextEditorInteraction Qt__TextInteractionFlag = 19

//
const Qt__TextBrowserInteraction Qt__TextInteractionFlag = 13

/*
This enum can be used to specify event priorities.



Note that these values are provided purely for convenience, since event priorities can be any value between INT_MAX and INT_MIN, inclusive. For example, you can define custom priorities as being relative to each other:


  enum CustomEventPriority
  {
      // An important event
      ImportantEventPriority = Qt::HighEventPriority,

      // A more important event
      MoreImportantEventPriority = ImportantEventPriority + 1,

      // A critical event
      CriticalEventPriority = 100 * MoreImportantEventPriority,

      // Not that important
      StatusEventPriority = Qt::LowEventPriority,

      // These are less important than Status events
      IdleProcessingDoneEventPriority = StatusEventPriority - 1
  };



See also QCoreApplication::postEvent().

*/
type Qt__EventPriority = int // core
// Events with this priority are sent before events with NormalEventPriority or LowEventPriority.
const Qt__HighEventPriority Qt__EventPriority = 1

// Events with this priority are sent after events with HighEventPriority, but before events with LowEventPriority.
const Qt__NormalEventPriority Qt__EventPriority = 0

//
const Qt__LowEventPriority Qt__EventPriority = -1

/*
This enum is used by QGraphicsLayoutItem::sizeHint()



This enum was introduced or modified in  Qt 4.4.

See also QGraphicsLayoutItem::sizeHint().

*/
type Qt__SizeHint = int // core
// is used to specify the minimum size of a graphics layout item.
const Qt__MinimumSize Qt__SizeHint = 0

// is used to specify the preferred size of a graphics layout item.
const Qt__PreferredSize Qt__SizeHint = 1

// is used to specify the maximum size of a graphics layout item.
const Qt__MaximumSize Qt__SizeHint = 2

// is used to specify the minimum descent of a text string in a graphics layout item.
const Qt__MinimumDescent Qt__SizeHint = 3

//
const Qt__NSizeHints Qt__SizeHint = 4

/*
This enum is used to describe parts of a window frame. It is returned by QGraphicsWidget::windowFrameSectionAt() to describe what section of the window frame is under the mouse.

ConstantValue
Qt::NoSection0
Qt::LeftSection1
Qt::TopLeftSection2
Qt::TopSection3
Qt::TopRightSection4
Qt::RightSection5
Qt::BottomRightSection6
Qt::BottomSection7
Qt::BottomLeftSection8
Qt::TitleBarArea9


This enum was introduced or modified in  Qt 4.4.

See also QGraphicsWidget::windowFrameEvent(), QGraphicsWidget::paintWindowFrame(), and QGraphicsWidget::windowFrameSectionAt().

*/
type Qt__WindowFrameSection = int // core
//
const Qt__NoSection Qt__WindowFrameSection = 0

//
const Qt__LeftSection Qt__WindowFrameSection = 1

//
const Qt__TopLeftSection Qt__WindowFrameSection = 2

//
const Qt__TopSection Qt__WindowFrameSection = 3

//
const Qt__TopRightSection Qt__WindowFrameSection = 4

//
const Qt__RightSection Qt__WindowFrameSection = 5

//
const Qt__BottomRightSection Qt__WindowFrameSection = 6

//
const Qt__BottomSection Qt__WindowFrameSection = 7

//
const Qt__BottomLeftSection Qt__WindowFrameSection = 8

//
const Qt__TitleBarArea Qt__WindowFrameSection = 9

/*


 */
type Qt__Initialization = int // core
//
const Qt__Uninitialized Qt__Initialization = 0

/*
This enum specifies the coordinate system.



This enum was introduced or modified in  Qt 4.6.

*/
type Qt__CoordinateSystem = int // core
// Coordinates are relative to the top-left corner of the object's paint device.
const Qt__DeviceCoordinates Qt__CoordinateSystem = 0

// Coordinates are relative to the top-left corner of the object.
const Qt__LogicalCoordinates Qt__CoordinateSystem = 1

/*


 */
type Qt__TouchPointState = int // core
//
const Qt__TouchPointPressed Qt__TouchPointState = 1

//
const Qt__TouchPointMoved Qt__TouchPointState = 2

//
const Qt__TouchPointStationary Qt__TouchPointState = 4

//
const Qt__TouchPointReleased Qt__TouchPointState = 8

/*
This enum type describes the state of a gesture.



This enum was introduced or modified in  Qt 4.6.

See also QGesture.

*/
type Qt__GestureState = int // core
// No gesture has been detected.
const Qt__NoGesture Qt__GestureState = 0

// A continuous gesture has started.
const Qt__GestureStarted Qt__GestureState = 1

// A gesture continues.
const Qt__GestureUpdated Qt__GestureState = 2

// A gesture has finished.
const Qt__GestureFinished Qt__GestureState = 3

// A gesture was canceled.
const Qt__GestureCanceled Qt__GestureState = 4

/*
This enum type describes the standard gestures.



User-defined gestures are registered with the QGestureRecognizer::registerRecognizer() function which generates a custom gesture ID with the Qt::CustomGesture flag set.

This enum was introduced or modified in  Qt 4.6.

See also QGesture, QWidget::grabGesture(), and QGraphicsObject::grabGesture().

*/
type Qt__GestureType = int // core
// A Tap gesture.
const Qt__TapGesture Qt__GestureType = 1

// A Tap-And-Hold (Long-Tap) gesture.
const Qt__TapAndHoldGesture Qt__GestureType = 2

// A Pan gesture.
const Qt__PanGesture Qt__GestureType = 3

// A Pinch gesture.
const Qt__PinchGesture Qt__GestureType = 4

// A Swipe gesture.
const Qt__SwipeGesture Qt__GestureType = 5

//
const Qt__CustomGesture Qt__GestureType = 256

//
const Qt__LastGestureType Qt__GestureType = -1

/*


 */
type Qt__GestureFlag = int // core
//
const Qt__DontStartGestureOnChildren Qt__GestureFlag = 1

//
const Qt__ReceivePartialGestures Qt__GestureFlag = 2

//
const Qt__IgnoredGesturesPropagateToParent Qt__GestureFlag = 4

/*
This enum returns the gesture type.



This enum was introduced or modified in  Qt 5.2.

*/
type Qt__NativeGestureType = int // core
// Sent before gesture event stream.
const Qt__BeginNativeGesture Qt__NativeGestureType = 0

// Sent after gesture event stream.
const Qt__EndNativeGesture Qt__NativeGestureType = 1

// Sent after a panning gesture. Similar to a click-and-drag mouse movement.
const Qt__PanNativeGesture Qt__NativeGestureType = 2

// Specifies the magnification delta in percent.
const Qt__ZoomNativeGesture Qt__NativeGestureType = 3

// Boolean magnification state.
const Qt__SmartZoomNativeGesture Qt__NativeGestureType = 4

// Rotation delta in degrees.
const Qt__RotateNativeGesture Qt__NativeGestureType = 5

// Sent after a swipe movements.
const Qt__SwipeNativeGesture Qt__NativeGestureType = 6

/*
This enum type describes the mode for moving focus.



Note: In 4.6, cursor navigation is only implemented for Symbian OS. On other platforms, it behaves as NavigationModeNone.

This enum was introduced or modified in  Qt 4.6.

See also QApplication::setNavigationMode() and QApplication::navigationMode().

*/
type Qt__NavigationMode = int // core
// Only the touch screen is used.
const Qt__NavigationModeNone Qt__NavigationMode = 0

// Qt::Key_Up and Qt::Key_Down are used to change focus.
const Qt__NavigationModeKeypadTabOrder Qt__NavigationMode = 1

// Qt::Key_Up, Qt::Key_Down, Qt::Key_Left and Qt::Key_Right are used to change focus.
const Qt__NavigationModeKeypadDirectional Qt__NavigationMode = 2

// The mouse cursor is used to change focus, it is displayed only on non touchscreen devices. The keypad is used to implement a virtual cursor, unless the device has an analog mouse type of input device (e.g. touchpad). This is the recommended setting for an application such as a web browser that needs pointer control on both touch and non-touch devices.
const Qt__NavigationModeCursorAuto Qt__NavigationMode = 3

// The mouse cursor is used to change focus, it is displayed regardless of device type. The keypad is used to implement a virtual cursor, unless the device has an analog mouse type of input device (e.g. touchpad)
const Qt__NavigationModeCursorForceVisible Qt__NavigationMode = 4

/*
This enum describes the movement style available to text cursors. The options are:


*/
type Qt__CursorMoveStyle = int // core
// Within a left-to-right text block, decrease cursor position when pressing left arrow key, increase cursor position when pressing the right arrow key. If the text block is right-to-left, the opposite behavior applies.
const Qt__LogicalMoveStyle Qt__CursorMoveStyle = 0

// Pressing the left arrow key will always cause the cursor to move left, regardless of the text's writing direction. Pressing the right arrow key will always cause the cursor to move right.
const Qt__VisualMoveStyle Qt__CursorMoveStyle = 1

/*
The timer type indicates how accurate a timer can be.




On Windows, Qt will use Windows's Multimedia timer facility (if available) for Qt::PreciseTimer and normal Windows timers for Qt::CoarseTimer and Qt::VeryCoarseTimer.


*/
type Qt__TimerType = int // core
// Precise timers try to keep millisecond accuracy
const Qt__PreciseTimer Qt__TimerType = 0

//
const Qt__CoarseTimer Qt__TimerType = 1

// Very coarse timers only keep full second accuracy
const Qt__VeryCoarseTimer Qt__TimerType = 2

/*
This enum describes the phase of scrolling.



This enum was introduced or modified in  Qt 5.2.

*/
type Qt__ScrollPhase = int // core
//
const Qt__NoScrollPhase Qt__ScrollPhase = 0

// Scrolling is about to begin, but the scrolling distance did not yet change.
const Qt__ScrollBegin Qt__ScrollPhase = 1

// The scrolling distance has changed (default).
const Qt__ScrollUpdate Qt__ScrollPhase = 2

// Scrolling has ended, but the scrolling distance did not change anymore.
const Qt__ScrollEnd Qt__ScrollPhase = 3

/*
This enum describes the source of a mouse event and can be useful to determine if the event is an artificial mouse event originating from another device such as a touchscreen.



This enum was introduced or modified in  Qt 5.3.

See also Qt::AA_SynthesizeMouseForUnhandledTouchEvents.

*/
type Qt__MouseEventSource = int // core
// The most common value. On platforms where such information is available this value indicates that the event was generated in response to a genuine mouse event in the system.
const Qt__MouseEventNotSynthesized Qt__MouseEventSource = 0

// Indicates that the mouse event was synthesized from a touch event by the platform.
const Qt__MouseEventSynthesizedBySystem Qt__MouseEventSource = 1

// Indicates that the mouse event was synthesized from an unhandled touch event by Qt.
const Qt__MouseEventSynthesizedByQt Qt__MouseEventSource = 2

//
const Qt__MouseEventSynthesizedByApplication Qt__MouseEventSource = 3

/*


 */
type Qt__MouseEventFlag = int // core
//
const Qt__MouseEventCreatedDoubleClick Qt__MouseEventFlag = 1

//
const Qt__MouseEventFlagMask Qt__MouseEventFlag = 255

/*
This enum describes the possible standards used by qChecksum().



This enum was introduced or modified in  Qt 5.9.

*/
type Qt__ChecksumType = int // core
//
const Qt__ChecksumIso3309 Qt__ChecksumType = 0

//
const Qt__ChecksumItuV41 Qt__ChecksumType = 1

/*


 */
type Qt__errc = int // stdglobal
//
const Qt__address_family_not_supported Qt__errc = 97

//
const Qt__address_in_use Qt__errc = 98

//
const Qt__address_not_available Qt__errc = 99

//
const Qt__already_connected Qt__errc = 106

//
const Qt__argument_list_too_long Qt__errc = 7

//
const Qt__argument_out_of_domain Qt__errc = 33

//
const Qt__bad_address Qt__errc = 14

//
const Qt__bad_file_descriptor Qt__errc = 9

//
const Qt__bad_message Qt__errc = 74

//
const Qt__broken_pipe Qt__errc = 32

//
const Qt__connection_aborted Qt__errc = 103

//
const Qt__connection_already_in_progress Qt__errc = 114

//
const Qt__connection_refused Qt__errc = 111

//
const Qt__connection_reset Qt__errc = 104

//
const Qt__cross_device_link Qt__errc = 18

//
const Qt__destination_address_required Qt__errc = 89

//
const Qt__device_or_resource_busy Qt__errc = 16

//
const Qt__directory_not_empty Qt__errc = 39

//
const Qt__executable_format_error Qt__errc = 8

//
const Qt__file_exists Qt__errc = 17

//
const Qt__file_too_large Qt__errc = 27

//
const Qt__filename_too_long Qt__errc = 36

//
const Qt__function_not_supported Qt__errc = 38

//
const Qt__host_unreachable Qt__errc = 113

//
const Qt__identifier_removed Qt__errc = 43

//
const Qt__illegal_byte_sequence Qt__errc = 84

//
const Qt__inappropriate_io_control_operation Qt__errc = 25

//
const Qt__interrupted Qt__errc = 4

//
const Qt__invalid_argument Qt__errc = 22

//
const Qt__invalid_seek Qt__errc = 29

//
const Qt__io_error Qt__errc = 5

//
const Qt__is_a_directory Qt__errc = 21

//
const Qt__message_size Qt__errc = 90

//
const Qt__network_down Qt__errc = 100

//
const Qt__network_reset Qt__errc = 102

//
const Qt__network_unreachable Qt__errc = 101

//
const Qt__no_buffer_space Qt__errc = 105

//
const Qt__no_child_process Qt__errc = 10

//
const Qt__no_link Qt__errc = 67

//
const Qt__no_lock_available Qt__errc = 37

//
const Qt__no_message_available Qt__errc = 61

//
const Qt__no_message Qt__errc = 42

//
const Qt__no_protocol_option Qt__errc = 92

//
const Qt__no_space_on_device Qt__errc = 28

//
const Qt__no_stream_resources Qt__errc = 63

//
const Qt__no_such_device_or_address Qt__errc = 6

//
const Qt__no_such_device Qt__errc = 19

//
const Qt__no_such_file_or_directory Qt__errc = 2

//
const Qt__no_such_process Qt__errc = 3

//
const Qt__not_a_directory Qt__errc = 20

//
const Qt__not_a_socket Qt__errc = 88

//
const Qt__not_a_stream Qt__errc = 60

//
const Qt__not_connected Qt__errc = 107

//
const Qt__not_enough_memory Qt__errc = 12

//
const Qt__not_supported Qt__errc = 95

//
const Qt__operation_canceled Qt__errc = 125

//
const Qt__operation_in_progress Qt__errc = 115

//
const Qt__operation_not_permitted Qt__errc = 1

//
const Qt__operation_not_supported Qt__errc = 95

//
const Qt__operation_would_block Qt__errc = 11

//
const Qt__owner_dead Qt__errc = 130

//
const Qt__permission_denied Qt__errc = 13

//
const Qt__protocol_error Qt__errc = 71

//
const Qt__protocol_not_supported Qt__errc = 93

//
const Qt__read_only_file_system Qt__errc = 30

//
const Qt__resource_deadlock_would_occur Qt__errc = 35

//
const Qt__resource_unavailable_try_again Qt__errc = 11

//
const Qt__result_out_of_range Qt__errc = 34

//
const Qt__state_not_recoverable Qt__errc = 131

//
const Qt__stream_timeout Qt__errc = 62

//
const Qt__text_file_busy Qt__errc = 26

//
const Qt__timed_out Qt__errc = 110

//
const Qt__too_many_files_open_in_system Qt__errc = 23

//
const Qt__too_many_files_open Qt__errc = 24

//
const Qt__too_many_links Qt__errc = 31

//
const Qt__too_many_symbolic_link_levels Qt__errc = 40

//
const Qt__value_too_large Qt__errc = 75

//
const Qt__wrong_protocol_type Qt__errc = 91

/*


 */
type Qt___Ios_Fmtflags = int // stdglobal
//
const Qt___S_boolalpha Qt___Ios_Fmtflags = 1

//
const Qt___S_dec Qt___Ios_Fmtflags = 2

//
const Qt___S_fixed Qt___Ios_Fmtflags = 4

//
const Qt___S_hex Qt___Ios_Fmtflags = 8

//
const Qt___S_internal Qt___Ios_Fmtflags = 16

//
const Qt___S_left Qt___Ios_Fmtflags = 32

//
const Qt___S_oct Qt___Ios_Fmtflags = 64

//
const Qt___S_right Qt___Ios_Fmtflags = 128

//
const Qt___S_scientific Qt___Ios_Fmtflags = 256

//
const Qt___S_showbase Qt___Ios_Fmtflags = 512

//
const Qt___S_showpoint Qt___Ios_Fmtflags = 1024

//
const Qt___S_showpos Qt___Ios_Fmtflags = 2048

//
const Qt___S_skipws Qt___Ios_Fmtflags = 4096

//
const Qt___S_unitbuf Qt___Ios_Fmtflags = 8192

//
const Qt___S_uppercase Qt___Ios_Fmtflags = 16384

//
const Qt___S_adjustfield Qt___Ios_Fmtflags = 176

//
const Qt___S_basefield Qt___Ios_Fmtflags = 74

//
const Qt___S_floatfield Qt___Ios_Fmtflags = 260

//
const Qt___S_ios_fmtflags_end Qt___Ios_Fmtflags = 65536

//
const Qt___S_ios_fmtflags_max Qt___Ios_Fmtflags = 2147483647

//
const Qt___S_ios_fmtflags_min Qt___Ios_Fmtflags = -2147483648

/*


 */
type Qt___Ios_Openmode = int // stdglobal
//
const Qt___S_app Qt___Ios_Openmode = 1

//
const Qt___S_ate Qt___Ios_Openmode = 2

//
const Qt___S_bin Qt___Ios_Openmode = 4

//
const Qt___S_in Qt___Ios_Openmode = 8

//
const Qt___S_out Qt___Ios_Openmode = 16

//
const Qt___S_trunc Qt___Ios_Openmode = 32

//
const Qt___S_ios_openmode_end Qt___Ios_Openmode = 65536

//
const Qt___S_ios_openmode_max Qt___Ios_Openmode = 2147483647

//
const Qt___S_ios_openmode_min Qt___Ios_Openmode = -2147483648

/*


 */
type Qt___Ios_Iostate = int // stdglobal
//
const Qt___S_goodbit Qt___Ios_Iostate = 0

//
const Qt___S_badbit Qt___Ios_Iostate = 1

//
const Qt___S_eofbit Qt___Ios_Iostate = 2

//
const Qt___S_failbit Qt___Ios_Iostate = 4

//
const Qt___S_ios_iostate_end Qt___Ios_Iostate = 65536

//
const Qt___S_ios_iostate_max Qt___Ios_Iostate = 2147483647

//
const Qt___S_ios_iostate_min Qt___Ios_Iostate = -2147483648

/*


 */
type Qt___Ios_Seekdir = int // stdglobal
//
const Qt___S_beg Qt___Ios_Seekdir = 0

//
const Qt___S_cur Qt___Ios_Seekdir = 1

//
const Qt___S_end Qt___Ios_Seekdir = 2

//
const Qt___S_ios_seekdir_end Qt___Ios_Seekdir = 65536

/*


 */
type Qt__io_errc = int // stdglobal
//
const Qt__stream Qt__io_errc = 1

/*


 */
type Qt___Rb_tree_color = int // stdglobal
//
const Qt___S_red Qt___Rb_tree_color = 0

//
const Qt___S_black Qt___Rb_tree_color = 1

/*


 */
type Qt__IteratorCapability = int // core
//
const Qt__ForwardCapability Qt__IteratorCapability = 1

//
const Qt__BiDirectionalCapability Qt__IteratorCapability = 2

//
const Qt__RandomAccessCapability Qt__IteratorCapability = 4

/*


 */
type Qt___Manager_operation = int // stdglobal
//
const Qt____get_type_info Qt___Manager_operation = 0

//
const Qt____get_functor_ptr Qt___Manager_operation = 1

//
const Qt____clone_functor Qt___Manager_operation = 2

//
const Qt____destroy_functor Qt___Manager_operation = 3

/*


 */
type Qt___Lock_policy = int // stdglobal
//
const Qt___S_single Qt___Lock_policy = 0

//
const Qt___S_mutex Qt___Lock_policy = 1

//
const Qt___S_atomic Qt___Lock_policy = 2

/*


 */
type Qt__pointer_safety = int // stdglobal
//
const Qt__relaxed Qt__pointer_safety = 0

//
const Qt__preferred Qt__pointer_safety = 1

//
const Qt__strict Qt__pointer_safety = 2

/*


 */
type Qt__future_errc = int // stdglobal
//
const Qt__future_already_retrieved Qt__future_errc = 1

//
const Qt__promise_already_satisfied Qt__future_errc = 2

//
const Qt__no_state Qt__future_errc = 3

//
const Qt__broken_promise Qt__future_errc = 4

/*


 */
type Qt__future_status = int // stdglobal
//
const Qt__ready Qt__future_status = 0

//
const Qt__timeout Qt__future_status = 1

//
const Qt__deferred Qt__future_status = 2

/*


 */
type Qt__DrawingHint = int // widgets
//
const Qt__OpaqueTopLeft Qt__DrawingHint = 1

//
const Qt__OpaqueTop Qt__DrawingHint = 2

//
const Qt__OpaqueTopRight Qt__DrawingHint = 4

//
const Qt__OpaqueLeft Qt__DrawingHint = 8

//
const Qt__OpaqueCenter Qt__DrawingHint = 16

//
const Qt__OpaqueRight Qt__DrawingHint = 32

//
const Qt__OpaqueBottomLeft Qt__DrawingHint = 64

//
const Qt__OpaqueBottom Qt__DrawingHint = 128

//
const Qt__OpaqueBottomRight Qt__DrawingHint = 256

//
const Qt__OpaqueCorners Qt__DrawingHint = 325

//
const Qt__OpaqueEdges Qt__DrawingHint = 170

//
const Qt__OpaqueFrame Qt__DrawingHint = 495

//
const Qt__OpaqueAll Qt__DrawingHint = 511

/*
Describes the two types of keys QSslKey supports.


*/
type Qt__KeyType = int // network
// A private key.
const Qt__PrivateKey Qt__KeyType = 0

// A public key.
const Qt__PublicKey Qt__KeyType = 1

/*
Describes supported encoding formats for certificates and keys.


*/
type Qt__EncodingFormat = int // network
// The PEM format.
const Qt__Pem Qt__EncodingFormat = 0

// The DER format.
const Qt__Der Qt__EncodingFormat = 1

/*
Describes the different key algorithms supported by QSslKey.



The opaque key facility allows applications to add support for facilities such as PKCS#11 that Qt does not currently offer natively.

*/
type Qt__KeyAlgorithm = int // network
// A key that should be treated as a 'black box' by QSslKey.
const Qt__Opaque Qt__KeyAlgorithm = 0

// The RSA algorithm.
const Qt__Rsa Qt__KeyAlgorithm = 1

// The DSA algorithm.
const Qt__Dsa Qt__KeyAlgorithm = 2

// The Elliptic Curve algorithm
const Qt__Ec Qt__KeyAlgorithm = 3

/*
Describes the key types for alternative name entries in QSslCertificate.



Note: In Qt 4, this enum was called AlternateNameEntryType. That name is deprecated in Qt 5.

See also QSslCertificate::subjectAlternativeNames().

*/
type Qt__AlternativeNameEntryType = int // network
// An email entry; the entry contains an email address that the certificate is valid for.
const Qt__EmailEntry Qt__AlternativeNameEntryType = 0

// A DNS host name entry; the entry contains a host name entry that the certificate is valid for. The entry may contain wildcards.
const Qt__DnsEntry Qt__AlternativeNameEntryType = 1

/*
Describes the protocol of the cipher.



Note: most servers understand both SSL and TLS, but it is recommended to use TLS only for security reasons. However, SSL and TLS are not compatible with each other: if you get unexpected handshake failures, verify that you chose the correct setting for your protocol.

*/
type Qt__SslProtocol = int // network
// SSLv3
const Qt__SslV3 Qt__SslProtocol = 0

// SSLv2
const Qt__SslV2 Qt__SslProtocol = 1

//
const Qt__TlsV1_0 Qt__SslProtocol = 2

//
const Qt__TlsV1_1 Qt__SslProtocol = 3

//
const Qt__TlsV1_2 Qt__SslProtocol = 4

//
const Qt__AnyProtocol Qt__SslProtocol = 5

//
const Qt__TlsV1SslV3 Qt__SslProtocol = 6

//
const Qt__SecureProtocols Qt__SslProtocol = 7

//
const Qt__TlsV1_0OrLater Qt__SslProtocol = 8

//
const Qt__TlsV1_1OrLater Qt__SslProtocol = 9

//
const Qt__TlsV1_2OrLater Qt__SslProtocol = 10

//
const Qt__UnknownProtocol Qt__SslProtocol = -1

/*


 */
type Qt__SslOption = int // network
//
const Qt__SslOptionDisableEmptyFragments Qt__SslOption = 1

//
const Qt__SslOptionDisableSessionTickets Qt__SslOption = 2

//
const Qt__SslOptionDisableCompression Qt__SslOption = 4

//
const Qt__SslOptionDisableServerNameIndication Qt__SslOption = 8

//
const Qt__SslOptionDisableLegacyRenegotiation Qt__SslOption = 16

//
const Qt__SslOptionDisableSessionSharing Qt__SslOption = 32

//
const Qt__SslOptionDisableSessionPersistence Qt__SslOption = 64

//
const Qt__SslOptionDisableServerCipherPreference Qt__SslOption = 128

/*


 */
type Qt__AutoParentResult = int // qml
//
const Qt__Parented Qt__AutoParentResult = 0

//
const Qt__IncompatibleObject Qt__AutoParentResult = 1

//
const Qt__IncompatibleParent Qt__AutoParentResult = 2

/*


 */
type Qt__RegistrationType = int // qml
//
const Qt__TypeRegistration Qt__RegistrationType = 0

//
const Qt__InterfaceRegistration Qt__RegistrationType = 1

//
const Qt__AutoParentRegistration Qt__RegistrationType = 2

//
const Qt__SingletonRegistration Qt__RegistrationType = 3

//
const Qt__CompositeRegistration Qt__RegistrationType = 4

//
const Qt__CompositeSingletonRegistration Qt__RegistrationType = 5

//
const Qt__QmlUnitCacheHookRegistration Qt__RegistrationType = 6

/*
Enumerates the levels of support a media service provider may have for a feature.


*/
type Qt__SupportEstimate = int // multimedia
// The feature is not supported.
const Qt__NotSupported Qt__SupportEstimate = 0

// The feature may be supported.
const Qt__MaybeSupported Qt__SupportEstimate = 1

// The feature is probably supported.
const Qt__ProbablySupported Qt__SupportEstimate = 2

// The service is the preferred provider of a service.
const Qt__PreferredService Qt__SupportEstimate = 3

/*
Enumerates quality encoding levels.

ConstantValue
QMultimedia::VeryLowQuality0
QMultimedia::LowQuality1
QMultimedia::NormalQuality2
QMultimedia::HighQuality3
QMultimedia::VeryHighQuality4

*/
type Qt__EncodingQuality = int // multimedia
//
const Qt__VeryLowQuality Qt__EncodingQuality = 0

//
const Qt__LowQuality Qt__EncodingQuality = 1

//
const Qt__NormalQuality Qt__EncodingQuality = 2

//
const Qt__HighQuality Qt__EncodingQuality = 3

//
const Qt__VeryHighQuality Qt__EncodingQuality = 4

/*
Enumerates encoding modes.


*/
type Qt__EncodingMode = int // multimedia
// Encoding will aim to have a constant quality, adjusting bitrate to fit.
const Qt__ConstantQualityEncoding Qt__EncodingMode = 0

// Encoding will use a constant bit rate, adjust quality to fit.
const Qt__ConstantBitRateEncoding Qt__EncodingMode = 1

// Encoding will try to keep an average bitrate setting, but will use more or less as needed.
const Qt__AverageBitRateEncoding Qt__EncodingMode = 2

// The media will first be processed to determine the characteristics, and then processed a second time allocating more bits to the areas that need it.
const Qt__TwoPassEncoding Qt__EncodingMode = 3

/*
Enumerates Service status errors.


*/
type Qt__AvailabilityStatus = int // multimedia
// The service is operating correctly.
const Qt__Available Qt__AvailabilityStatus = 0

// There is no service available to provide the requested functionality.
const Qt__ServiceMissing Qt__AvailabilityStatus = 1

// The service must wait for access to necessary resources.
const Qt__Busy Qt__AvailabilityStatus = 2

// The service could not allocate resources required to function correctly.
const Qt__ResourceError Qt__AvailabilityStatus = 3

/*

 */
type Qt__Error = int // multimedia
// No errors have occurred
const Qt__NoError Qt__Error = 0

// An error occurred opening the audio device
const Qt__OpenError Qt__Error = 1

// An error occurred during read/write of audio device
const Qt__IOError Qt__Error = 2

// Audio data is not being fed to the audio device at a fast enough rate
const Qt__UnderrunError Qt__Error = 3

// A non-recoverable error has occurred, the audio device is not usable at this time.
const Qt__FatalError Qt__Error = 4

/*

 */
type Qt__State = int // multimedia
// Audio data is being processed, this state is set after start() is called and while audio data is available to be processed.
const Qt__ActiveState Qt__State = 0

// The audio stream is in a suspended state. Entered after suspend() is called or when another stream takes control of the audio device. In the later case, a call to resume will return control of the audio device to this stream. This should usually only be done upon user request.
const Qt__SuspendedState Qt__State = 1

// The audio device is closed, and is not processing any audio data
const Qt__StoppedState Qt__State = 2

// The QIODevice passed in has no data and audio system's buffer is empty, this state is set after start() is called and while no audio data is available to be processed.
const Qt__IdleState Qt__State = 3

// This stream is in a suspended state because another higher priority stream currently has control of the audio device. Playback cannot resume until the higher priority stream relinquishes control of the audio device.
const Qt__InterruptedState Qt__State = 4

/*

 */
type Qt__Mode = int // multimedia
// audio input device
const Qt__AudioInput Qt__Mode = 0

// audio output device
const Qt__AudioOutput Qt__Mode = 1

/*
This enum describes the role of an audio stream.



This enum was introduced or modified in  Qt 5.6.

See also QMediaPlayer::setAudioRole().

*/
type Qt__Role = int // multimedia
// The role is unknown or undefined
const Qt__UnknownRole Qt__Role = 0

// Music
const Qt__MusicRole Qt__Role = 1

// Soundtrack from a movie or a video
const Qt__VideoRole Qt__Role = 2

// Voice communications, such as telephony
const Qt__VoiceCommunicationRole Qt__Role = 3

// Alarm
const Qt__AlarmRole Qt__Role = 4

// Notification, such as an incoming e-mail or a chat request
const Qt__NotificationRole Qt__Role = 5

// Ringtone
const Qt__RingtoneRole Qt__Role = 6

// For accessibility, such as with a screen reader
const Qt__AccessibilityRole Qt__Role = 7

// Sonification, such as with user interface sounds
const Qt__SonificationRole Qt__Role = 8

// Game audio
const Qt__GameRole Qt__Role = 9

/*
This enum defines the different audio volume scales.



This enum was introduced or modified in  Qt 5.8.

See also QAudio::convertVolume().

*/
type Qt__VolumeScale = int // multimedia
//
const Qt__LinearVolumeScale Qt__VolumeScale = 0

//
const Qt__CubicVolumeScale Qt__VolumeScale = 1

//
const Qt__LogarithmicVolumeScale Qt__VolumeScale = 2

//
const Qt__DecibelVolumeScale Qt__VolumeScale = 3

/*


 */
type Qt___EXCEPTION_DISPOSITION = int // stdglobal
//
const Qt__ExceptionContinueExecution Qt___EXCEPTION_DISPOSITION = 0

//
const Qt__ExceptionContinueSearch Qt___EXCEPTION_DISPOSITION = 1

//
const Qt__ExceptionNestedException Qt___EXCEPTION_DISPOSITION = 2

//
const Qt__ExceptionCollidedUnwind Qt___EXCEPTION_DISPOSITION = 3

/*


 */
type Qt__tagExtendedErrorParamTypes = int // stdglobal
//
const Qt__eeptAnsiString Qt__tagExtendedErrorParamTypes = 1

//
const Qt__eeptUnicodeString Qt__tagExtendedErrorParamTypes = 2

//
const Qt__eeptLongVal Qt__tagExtendedErrorParamTypes = 3

//
const Qt__eeptShortVal Qt__tagExtendedErrorParamTypes = 4

//
const Qt__eeptPointerVal Qt__tagExtendedErrorParamTypes = 5

//
const Qt__eeptNone Qt__tagExtendedErrorParamTypes = 6

//
const Qt__eeptBinary Qt__tagExtendedErrorParamTypes = 7

/*


 */
type Qt___RPC_NOTIFICATION_TYPES = int // stdglobal
//
const Qt__RpcNotificationTypeNone Qt___RPC_NOTIFICATION_TYPES = 0

//
const Qt__RpcNotificationTypeEvent Qt___RPC_NOTIFICATION_TYPES = 1

//
const Qt__RpcNotificationTypeApc Qt___RPC_NOTIFICATION_TYPES = 2

//
const Qt__RpcNotificationTypeIoc Qt___RPC_NOTIFICATION_TYPES = 3

//
const Qt__RpcNotificationTypeHwnd Qt___RPC_NOTIFICATION_TYPES = 4

//
const Qt__RpcNotificationTypeCallback Qt___RPC_NOTIFICATION_TYPES = 5

/*


 */
type Qt___RPC_ASYNC_EVENT = int // stdglobal
//
const Qt__RpcCallComplete Qt___RPC_ASYNC_EVENT = 0

//
const Qt__RpcSendComplete Qt___RPC_ASYNC_EVENT = 1

//
const Qt__RpcReceiveComplete Qt___RPC_ASYNC_EVENT = 2

//
const Qt__RpcClientDisconnect Qt___RPC_ASYNC_EVENT = 3

//
const Qt__RpcClientCancel Qt___RPC_ASYNC_EVENT = 4

/*


 */
type Qt___MEDIA_TYPE = int // stdglobal
//
const Qt__Unknown Qt___MEDIA_TYPE = 0

//
const Qt__F5_1Pt2_512 Qt___MEDIA_TYPE = 1

//
const Qt__F3_1Pt44_512 Qt___MEDIA_TYPE = 2

//
const Qt__F3_2Pt88_512 Qt___MEDIA_TYPE = 3

//
const Qt__F3_20Pt8_512 Qt___MEDIA_TYPE = 4

//
const Qt__F3_720_512 Qt___MEDIA_TYPE = 5

//
const Qt__F5_360_512 Qt___MEDIA_TYPE = 6

//
const Qt__F5_320_512 Qt___MEDIA_TYPE = 7

//
const Qt__F5_320_1024 Qt___MEDIA_TYPE = 8

//
const Qt__F5_180_512 Qt___MEDIA_TYPE = 9

//
const Qt__F5_160_512 Qt___MEDIA_TYPE = 10

//
const Qt__RemovableMedia Qt___MEDIA_TYPE = 11

//
const Qt__FixedMedia Qt___MEDIA_TYPE = 12

//
const Qt__F3_120M_512 Qt___MEDIA_TYPE = 13

//
const Qt__F3_640_512 Qt___MEDIA_TYPE = 14

//
const Qt__F5_640_512 Qt___MEDIA_TYPE = 15

//
const Qt__F5_720_512 Qt___MEDIA_TYPE = 16

//
const Qt__F3_1Pt2_512 Qt___MEDIA_TYPE = 17

//
const Qt__F3_1Pt23_1024 Qt___MEDIA_TYPE = 18

//
const Qt__F5_1Pt23_1024 Qt___MEDIA_TYPE = 19

//
const Qt__F3_128Mb_512 Qt___MEDIA_TYPE = 20

//
const Qt__F3_230Mb_512 Qt___MEDIA_TYPE = 21

//
const Qt__F8_256_128 Qt___MEDIA_TYPE = 22

/*


 */
type Qt__HBitmapFormat = int // winextras
//
const Qt__HBitmapNoAlpha Qt__HBitmapFormat = 0

//
const Qt__HBitmapPremultipliedAlpha Qt__HBitmapFormat = 1

//
const Qt__HBitmapAlpha Qt__HBitmapFormat = 2

/*


 */
type Qt__WindowFlip3DPolicy = int // winextras
//
const Qt__FlipDefault Qt__WindowFlip3DPolicy = 0

//
const Qt__FlipExcludeBelow Qt__WindowFlip3DPolicy = 1

//
const Qt__FlipExcludeAbove Qt__WindowFlip3DPolicy = 2

/*


 */
type Qt__jobjectRefType = int // stdglobal
//
const Qt__JNIInvalidRefType Qt__jobjectRefType = 0

//
const Qt__JNILocalRefType Qt__jobjectRefType = 1

//
const Qt__JNIGlobalRefType Qt__jobjectRefType = 2

//
const Qt__JNIWeakGlobalRefType Qt__jobjectRefType = 3

/*


 */
type Qt__BindFlag = int // androidextras
//
const Qt__None Qt__BindFlag = 0

//
const Qt__AutoCreate Qt__BindFlag = 1

//
const Qt__DebugUnbind Qt__BindFlag = 2

//
const Qt__NotForeground Qt__BindFlag = 4

//
const Qt__AboveClient Qt__BindFlag = 8

//
const Qt__AllowOomManagement Qt__BindFlag = 16

//
const Qt__WaivePriority Qt__BindFlag = 32

//
const Qt__Important Qt__BindFlag = 64

//
const Qt__AdjustWithActivity Qt__BindFlag = 128

//
const Qt__ExternalService Qt__BindFlag = -2147483648

/*


 */
type Qt__PermissionResult = int // androidextras
//
const Qt__Granted Qt__PermissionResult = 0

//
const Qt__Denied Qt__PermissionResult = 1

//  body block end

//  keep block begin

//  keep block end
