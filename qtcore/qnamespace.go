package qtcore

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

//  ext block end

//  body block begin

type Qt__float_round_style = int

const Qt__round_indeterminate Qt__float_round_style = -1
const Qt__round_toward_zero Qt__float_round_style = 0
const Qt__round_to_nearest Qt__float_round_style = 1
const Qt__round_toward_infinity Qt__float_round_style = 2
const Qt__round_toward_neg_infinity Qt__float_round_style = 3

type Qt__float_denorm_style = int

const Qt__denorm_indeterminate Qt__float_denorm_style = -1
const Qt__denorm_absent Qt__float_denorm_style = 0
const Qt__denorm_present Qt__float_denorm_style = 1

type Qt__QtMsgType = int

const Qt__QtDebugMsg Qt__QtMsgType = 0
const Qt__QtWarningMsg Qt__QtMsgType = 1
const Qt__QtCriticalMsg Qt__QtMsgType = 2
const Qt__QtFatalMsg Qt__QtMsgType = 3
const Qt__QtInfoMsg Qt__QtMsgType = 4
const Qt__QtSystemMsg Qt__QtMsgType = 2

type Qt__memory_order = int

const Qt__memory_order_relaxed Qt__memory_order = 0
const Qt__memory_order_consume Qt__memory_order = 1
const Qt__memory_order_acquire Qt__memory_order = 2
const Qt__memory_order_release Qt__memory_order = 3
const Qt__memory_order_acq_rel Qt__memory_order = 4
const Qt__memory_order_seq_cst Qt__memory_order = 5

type Qt____memory_order_modifier = int

const Qt____memory_order_mask Qt____memory_order_modifier = 65535
const Qt____memory_order_modifier_mask Qt____memory_order_modifier = -65536
const Qt____memory_order_hle_acquire Qt____memory_order_modifier = 65536
const Qt____memory_order_hle_release Qt____memory_order_modifier = 131072

type Qt__GlobalColor = int

const Qt__color0 Qt__GlobalColor = 0
const Qt__color1 Qt__GlobalColor = 1
const Qt__black Qt__GlobalColor = 2
const Qt__white Qt__GlobalColor = 3
const Qt__darkGray Qt__GlobalColor = 4
const Qt__gray Qt__GlobalColor = 5
const Qt__lightGray Qt__GlobalColor = 6
const Qt__red Qt__GlobalColor = 7
const Qt__green Qt__GlobalColor = 8
const Qt__blue Qt__GlobalColor = 9
const Qt__cyan Qt__GlobalColor = 10
const Qt__magenta Qt__GlobalColor = 11
const Qt__yellow Qt__GlobalColor = 12
const Qt__darkRed Qt__GlobalColor = 13
const Qt__darkGreen Qt__GlobalColor = 14
const Qt__darkBlue Qt__GlobalColor = 15
const Qt__darkCyan Qt__GlobalColor = 16
const Qt__darkMagenta Qt__GlobalColor = 17
const Qt__darkYellow Qt__GlobalColor = 18
const Qt__transparent Qt__GlobalColor = 19

type Qt__KeyboardModifier = int

const Qt__NoModifier Qt__KeyboardModifier = 0
const Qt__ShiftModifier Qt__KeyboardModifier = 33554432
const Qt__ControlModifier Qt__KeyboardModifier = 67108864
const Qt__AltModifier Qt__KeyboardModifier = 134217728
const Qt__MetaModifier Qt__KeyboardModifier = 268435456
const Qt__KeypadModifier Qt__KeyboardModifier = 536870912
const Qt__GroupSwitchModifier Qt__KeyboardModifier = 1073741824
const Qt__KeyboardModifierMask Qt__KeyboardModifier = -33554432

type Qt__Modifier = int

const Qt__META Qt__Modifier = 268435456
const Qt__SHIFT Qt__Modifier = 33554432
const Qt__CTRL Qt__Modifier = 67108864
const Qt__ALT Qt__Modifier = 134217728
const Qt__MODIFIER_MASK Qt__Modifier = -33554432
const Qt__UNICODE_ACCEL Qt__Modifier = 0

type Qt__MouseButton = int

const Qt__NoButton Qt__MouseButton = 0
const Qt__LeftButton Qt__MouseButton = 1
const Qt__RightButton Qt__MouseButton = 2
const Qt__MidButton Qt__MouseButton = 4
const Qt__MiddleButton Qt__MouseButton = 4
const Qt__BackButton Qt__MouseButton = 8
const Qt__XButton1 Qt__MouseButton = 8
const Qt__ExtraButton1 Qt__MouseButton = 8
const Qt__ForwardButton Qt__MouseButton = 16
const Qt__XButton2 Qt__MouseButton = 16
const Qt__ExtraButton2 Qt__MouseButton = 16
const Qt__TaskButton Qt__MouseButton = 32
const Qt__ExtraButton3 Qt__MouseButton = 32
const Qt__ExtraButton4 Qt__MouseButton = 64
const Qt__ExtraButton5 Qt__MouseButton = 128
const Qt__ExtraButton6 Qt__MouseButton = 256
const Qt__ExtraButton7 Qt__MouseButton = 512
const Qt__ExtraButton8 Qt__MouseButton = 1024
const Qt__ExtraButton9 Qt__MouseButton = 2048
const Qt__ExtraButton10 Qt__MouseButton = 4096
const Qt__ExtraButton11 Qt__MouseButton = 8192
const Qt__ExtraButton12 Qt__MouseButton = 16384
const Qt__ExtraButton13 Qt__MouseButton = 32768
const Qt__ExtraButton14 Qt__MouseButton = 65536
const Qt__ExtraButton15 Qt__MouseButton = 131072
const Qt__ExtraButton16 Qt__MouseButton = 262144
const Qt__ExtraButton17 Qt__MouseButton = 524288
const Qt__ExtraButton18 Qt__MouseButton = 1048576
const Qt__ExtraButton19 Qt__MouseButton = 2097152
const Qt__ExtraButton20 Qt__MouseButton = 4194304
const Qt__ExtraButton21 Qt__MouseButton = 8388608
const Qt__ExtraButton22 Qt__MouseButton = 16777216
const Qt__ExtraButton23 Qt__MouseButton = 33554432
const Qt__ExtraButton24 Qt__MouseButton = 67108864
const Qt__AllButtons Qt__MouseButton = 134217727
const Qt__MaxMouseButton Qt__MouseButton = 67108864
const Qt__MouseButtonMask Qt__MouseButton = -1

type Qt__Orientation = int

const Qt__Horizontal Qt__Orientation = 1
const Qt__Vertical Qt__Orientation = 2

type Qt__FocusPolicy = int

const Qt__NoFocus Qt__FocusPolicy = 0
const Qt__TabFocus Qt__FocusPolicy = 1
const Qt__ClickFocus Qt__FocusPolicy = 2
const Qt__StrongFocus Qt__FocusPolicy = 11
const Qt__WheelFocus Qt__FocusPolicy = 15

type Qt__TabFocusBehavior = int

const Qt__NoTabFocus Qt__TabFocusBehavior = 0
const Qt__TabFocusTextControls Qt__TabFocusBehavior = 1
const Qt__TabFocusListControls Qt__TabFocusBehavior = 2
const Qt__TabFocusAllControls Qt__TabFocusBehavior = 255

type Qt__SortOrder = int

const Qt__AscendingOrder Qt__SortOrder = 0
const Qt__DescendingOrder Qt__SortOrder = 1

type Qt__TileRule = int

const Qt__StretchTile Qt__TileRule = 0
const Qt__RepeatTile Qt__TileRule = 1
const Qt__RoundTile Qt__TileRule = 2

type Qt__AlignmentFlag = int

const Qt__AlignLeft Qt__AlignmentFlag = 1
const Qt__AlignLeading Qt__AlignmentFlag = 1
const Qt__AlignRight Qt__AlignmentFlag = 2
const Qt__AlignTrailing Qt__AlignmentFlag = 2
const Qt__AlignHCenter Qt__AlignmentFlag = 4
const Qt__AlignJustify Qt__AlignmentFlag = 8
const Qt__AlignAbsolute Qt__AlignmentFlag = 16
const Qt__AlignHorizontal_Mask Qt__AlignmentFlag = 31
const Qt__AlignTop Qt__AlignmentFlag = 32
const Qt__AlignBottom Qt__AlignmentFlag = 64
const Qt__AlignVCenter Qt__AlignmentFlag = 128
const Qt__AlignBaseline Qt__AlignmentFlag = 256
const Qt__AlignVertical_Mask Qt__AlignmentFlag = 480
const Qt__AlignCenter Qt__AlignmentFlag = 132

type Qt__TextFlag = int

const Qt__TextSingleLine Qt__TextFlag = 256
const Qt__TextDontClip Qt__TextFlag = 512
const Qt__TextExpandTabs Qt__TextFlag = 1024
const Qt__TextShowMnemonic Qt__TextFlag = 2048
const Qt__TextWordWrap Qt__TextFlag = 4096
const Qt__TextWrapAnywhere Qt__TextFlag = 8192
const Qt__TextDontPrint Qt__TextFlag = 16384
const Qt__TextIncludeTrailingSpaces Qt__TextFlag = 134217728
const Qt__TextHideMnemonic Qt__TextFlag = 32768
const Qt__TextJustificationForced Qt__TextFlag = 65536
const Qt__TextForceLeftToRight Qt__TextFlag = 131072
const Qt__TextForceRightToLeft Qt__TextFlag = 262144
const Qt__TextLongestVariant Qt__TextFlag = 524288
const Qt__TextBypassShaping Qt__TextFlag = 1048576

type Qt__TextElideMode = int

const Qt__ElideLeft Qt__TextElideMode = 0
const Qt__ElideRight Qt__TextElideMode = 1
const Qt__ElideMiddle Qt__TextElideMode = 2
const Qt__ElideNone Qt__TextElideMode = 3

type Qt__WhiteSpaceMode = int

const Qt__WhiteSpaceNormal Qt__WhiteSpaceMode = 0
const Qt__WhiteSpacePre Qt__WhiteSpaceMode = 1
const Qt__WhiteSpaceNoWrap Qt__WhiteSpaceMode = 2
const Qt__WhiteSpaceModeUndefined Qt__WhiteSpaceMode = -1

type Qt__HitTestAccuracy = int

const Qt__ExactHit Qt__HitTestAccuracy = 0
const Qt__FuzzyHit Qt__HitTestAccuracy = 1

type Qt__WindowType = int

const Qt__Widget Qt__WindowType = 0
const Qt__Window Qt__WindowType = 1
const Qt__Dialog Qt__WindowType = 3
const Qt__Sheet Qt__WindowType = 5
const Qt__Drawer Qt__WindowType = 7
const Qt__Popup Qt__WindowType = 9
const Qt__Tool Qt__WindowType = 11
const Qt__ToolTip Qt__WindowType = 13
const Qt__SplashScreen Qt__WindowType = 15
const Qt__Desktop Qt__WindowType = 17
const Qt__SubWindow Qt__WindowType = 18
const Qt__ForeignWindow Qt__WindowType = 33
const Qt__CoverWindow Qt__WindowType = 65
const Qt__WindowType_Mask Qt__WindowType = 255
const Qt__MSWindowsFixedSizeDialogHint Qt__WindowType = 256
const Qt__MSWindowsOwnDC Qt__WindowType = 512
const Qt__BypassWindowManagerHint Qt__WindowType = 1024
const Qt__X11BypassWindowManagerHint Qt__WindowType = 1024
const Qt__FramelessWindowHint Qt__WindowType = 2048
const Qt__WindowTitleHint Qt__WindowType = 4096
const Qt__WindowSystemMenuHint Qt__WindowType = 8192
const Qt__WindowMinimizeButtonHint Qt__WindowType = 16384
const Qt__WindowMaximizeButtonHint Qt__WindowType = 32768
const Qt__WindowMinMaxButtonsHint Qt__WindowType = 49152
const Qt__WindowContextHelpButtonHint Qt__WindowType = 65536
const Qt__WindowShadeButtonHint Qt__WindowType = 131072
const Qt__WindowStaysOnTopHint Qt__WindowType = 262144
const Qt__WindowTransparentForInput Qt__WindowType = 524288
const Qt__WindowOverridesSystemGestures Qt__WindowType = 1048576
const Qt__WindowDoesNotAcceptFocus Qt__WindowType = 2097152
const Qt__MaximizeUsingFullscreenGeometryHint Qt__WindowType = 4194304
const Qt__CustomizeWindowHint Qt__WindowType = 33554432
const Qt__WindowStaysOnBottomHint Qt__WindowType = 67108864
const Qt__WindowCloseButtonHint Qt__WindowType = 134217728
const Qt__MacWindowToolBarButtonHint Qt__WindowType = 268435456
const Qt__BypassGraphicsProxyWidget Qt__WindowType = 536870912
const Qt__NoDropShadowWindowHint Qt__WindowType = 1073741824
const Qt__WindowFullscreenButtonHint Qt__WindowType = -2147483648

type Qt__WindowState = int

const Qt__WindowNoState Qt__WindowState = 0
const Qt__WindowMinimized Qt__WindowState = 1
const Qt__WindowMaximized Qt__WindowState = 2
const Qt__WindowFullScreen Qt__WindowState = 4
const Qt__WindowActive Qt__WindowState = 8

type Qt__ApplicationState = int

const Qt__ApplicationSuspended Qt__ApplicationState = 0
const Qt__ApplicationHidden Qt__ApplicationState = 1
const Qt__ApplicationInactive Qt__ApplicationState = 2
const Qt__ApplicationActive Qt__ApplicationState = 4

type Qt__ScreenOrientation = int

const Qt__PrimaryOrientation Qt__ScreenOrientation = 0
const Qt__PortraitOrientation Qt__ScreenOrientation = 1
const Qt__LandscapeOrientation Qt__ScreenOrientation = 2
const Qt__InvertedPortraitOrientation Qt__ScreenOrientation = 4
const Qt__InvertedLandscapeOrientation Qt__ScreenOrientation = 8

type Qt__WidgetAttribute = int

const Qt__WA_Disabled Qt__WidgetAttribute = 0
const Qt__WA_UnderMouse Qt__WidgetAttribute = 1
const Qt__WA_MouseTracking Qt__WidgetAttribute = 2
const Qt__WA_ContentsPropagated Qt__WidgetAttribute = 3
const Qt__WA_OpaquePaintEvent Qt__WidgetAttribute = 4
const Qt__WA_NoBackground Qt__WidgetAttribute = 4
const Qt__WA_StaticContents Qt__WidgetAttribute = 5
const Qt__WA_LaidOut Qt__WidgetAttribute = 7
const Qt__WA_PaintOnScreen Qt__WidgetAttribute = 8
const Qt__WA_NoSystemBackground Qt__WidgetAttribute = 9
const Qt__WA_UpdatesDisabled Qt__WidgetAttribute = 10
const Qt__WA_Mapped Qt__WidgetAttribute = 11
const Qt__WA_MacNoClickThrough Qt__WidgetAttribute = 12
const Qt__WA_InputMethodEnabled Qt__WidgetAttribute = 14
const Qt__WA_WState_Visible Qt__WidgetAttribute = 15
const Qt__WA_WState_Hidden Qt__WidgetAttribute = 16
const Qt__WA_ForceDisabled Qt__WidgetAttribute = 32
const Qt__WA_KeyCompression Qt__WidgetAttribute = 33
const Qt__WA_PendingMoveEvent Qt__WidgetAttribute = 34
const Qt__WA_PendingResizeEvent Qt__WidgetAttribute = 35
const Qt__WA_SetPalette Qt__WidgetAttribute = 36
const Qt__WA_SetFont Qt__WidgetAttribute = 37
const Qt__WA_SetCursor Qt__WidgetAttribute = 38
const Qt__WA_NoChildEventsFromChildren Qt__WidgetAttribute = 39
const Qt__WA_WindowModified Qt__WidgetAttribute = 41
const Qt__WA_Resized Qt__WidgetAttribute = 42
const Qt__WA_Moved Qt__WidgetAttribute = 43
const Qt__WA_PendingUpdate Qt__WidgetAttribute = 44
const Qt__WA_InvalidSize Qt__WidgetAttribute = 45
const Qt__WA_MacBrushedMetal Qt__WidgetAttribute = 46
const Qt__WA_MacMetalStyle Qt__WidgetAttribute = 46
const Qt__WA_CustomWhatsThis Qt__WidgetAttribute = 47
const Qt__WA_LayoutOnEntireRect Qt__WidgetAttribute = 48
const Qt__WA_OutsideWSRange Qt__WidgetAttribute = 49
const Qt__WA_GrabbedShortcut Qt__WidgetAttribute = 50
const Qt__WA_TransparentForMouseEvents Qt__WidgetAttribute = 51
const Qt__WA_PaintUnclipped Qt__WidgetAttribute = 52
const Qt__WA_SetWindowIcon Qt__WidgetAttribute = 53
const Qt__WA_NoMouseReplay Qt__WidgetAttribute = 54
const Qt__WA_DeleteOnClose Qt__WidgetAttribute = 55
const Qt__WA_RightToLeft Qt__WidgetAttribute = 56
const Qt__WA_SetLayoutDirection Qt__WidgetAttribute = 57
const Qt__WA_NoChildEventsForParent Qt__WidgetAttribute = 58
const Qt__WA_ForceUpdatesDisabled Qt__WidgetAttribute = 59
const Qt__WA_WState_Created Qt__WidgetAttribute = 60
const Qt__WA_WState_CompressKeys Qt__WidgetAttribute = 61
const Qt__WA_WState_InPaintEvent Qt__WidgetAttribute = 62
const Qt__WA_WState_Reparented Qt__WidgetAttribute = 63
const Qt__WA_WState_ConfigPending Qt__WidgetAttribute = 64
const Qt__WA_WState_Polished Qt__WidgetAttribute = 66
const Qt__WA_WState_DND Qt__WidgetAttribute = 67
const Qt__WA_WState_OwnSizePolicy Qt__WidgetAttribute = 68
const Qt__WA_WState_ExplicitShowHide Qt__WidgetAttribute = 69
const Qt__WA_ShowModal Qt__WidgetAttribute = 70
const Qt__WA_MouseNoMask Qt__WidgetAttribute = 71
const Qt__WA_GroupLeader Qt__WidgetAttribute = 72
const Qt__WA_NoMousePropagation Qt__WidgetAttribute = 73
const Qt__WA_Hover Qt__WidgetAttribute = 74
const Qt__WA_InputMethodTransparent Qt__WidgetAttribute = 75
const Qt__WA_QuitOnClose Qt__WidgetAttribute = 76
const Qt__WA_KeyboardFocusChange Qt__WidgetAttribute = 77
const Qt__WA_AcceptDrops Qt__WidgetAttribute = 78
const Qt__WA_DropSiteRegistered Qt__WidgetAttribute = 79
const Qt__WA_ForceAcceptDrops Qt__WidgetAttribute = 79
const Qt__WA_WindowPropagation Qt__WidgetAttribute = 80
const Qt__WA_NoX11EventCompression Qt__WidgetAttribute = 81
const Qt__WA_TintedBackground Qt__WidgetAttribute = 82
const Qt__WA_X11OpenGLOverlay Qt__WidgetAttribute = 83
const Qt__WA_AlwaysShowToolTips Qt__WidgetAttribute = 84
const Qt__WA_MacOpaqueSizeGrip Qt__WidgetAttribute = 85
const Qt__WA_SetStyle Qt__WidgetAttribute = 86
const Qt__WA_SetLocale Qt__WidgetAttribute = 87
const Qt__WA_MacShowFocusRect Qt__WidgetAttribute = 88
const Qt__WA_MacNormalSize Qt__WidgetAttribute = 89
const Qt__WA_MacSmallSize Qt__WidgetAttribute = 90
const Qt__WA_MacMiniSize Qt__WidgetAttribute = 91
const Qt__WA_LayoutUsesWidgetRect Qt__WidgetAttribute = 92
const Qt__WA_StyledBackground Qt__WidgetAttribute = 93
const Qt__WA_MSWindowsUseDirect3D Qt__WidgetAttribute = 94
const Qt__WA_CanHostQMdiSubWindowTitleBar Qt__WidgetAttribute = 95
const Qt__WA_MacAlwaysShowToolWindow Qt__WidgetAttribute = 96
const Qt__WA_StyleSheet Qt__WidgetAttribute = 97
const Qt__WA_ShowWithoutActivating Qt__WidgetAttribute = 98
const Qt__WA_X11BypassTransientForHint Qt__WidgetAttribute = 99
const Qt__WA_NativeWindow Qt__WidgetAttribute = 100
const Qt__WA_DontCreateNativeAncestors Qt__WidgetAttribute = 101
const Qt__WA_MacVariableSize Qt__WidgetAttribute = 102
const Qt__WA_DontShowOnScreen Qt__WidgetAttribute = 103
const Qt__WA_X11NetWmWindowTypeDesktop Qt__WidgetAttribute = 104
const Qt__WA_X11NetWmWindowTypeDock Qt__WidgetAttribute = 105
const Qt__WA_X11NetWmWindowTypeToolBar Qt__WidgetAttribute = 106
const Qt__WA_X11NetWmWindowTypeMenu Qt__WidgetAttribute = 107
const Qt__WA_X11NetWmWindowTypeUtility Qt__WidgetAttribute = 108
const Qt__WA_X11NetWmWindowTypeSplash Qt__WidgetAttribute = 109
const Qt__WA_X11NetWmWindowTypeDialog Qt__WidgetAttribute = 110
const Qt__WA_X11NetWmWindowTypeDropDownMenu Qt__WidgetAttribute = 111
const Qt__WA_X11NetWmWindowTypePopupMenu Qt__WidgetAttribute = 112
const Qt__WA_X11NetWmWindowTypeToolTip Qt__WidgetAttribute = 113
const Qt__WA_X11NetWmWindowTypeNotification Qt__WidgetAttribute = 114
const Qt__WA_X11NetWmWindowTypeCombo Qt__WidgetAttribute = 115
const Qt__WA_X11NetWmWindowTypeDND Qt__WidgetAttribute = 116
const Qt__WA_MacFrameworkScaled Qt__WidgetAttribute = 117
const Qt__WA_SetWindowModality Qt__WidgetAttribute = 118
const Qt__WA_WState_WindowOpacitySet Qt__WidgetAttribute = 119
const Qt__WA_TranslucentBackground Qt__WidgetAttribute = 120
const Qt__WA_AcceptTouchEvents Qt__WidgetAttribute = 121
const Qt__WA_WState_AcceptedTouchBeginEvent Qt__WidgetAttribute = 122
const Qt__WA_TouchPadAcceptSingleTouchEvents Qt__WidgetAttribute = 123
const Qt__WA_X11DoNotAcceptFocus Qt__WidgetAttribute = 126
const Qt__WA_MacNoShadow Qt__WidgetAttribute = 127
const Qt__WA_AlwaysStackOnTop Qt__WidgetAttribute = 128
const Qt__WA_TabletTracking Qt__WidgetAttribute = 129
const Qt__WA_AttributeCount Qt__WidgetAttribute = 130

type Qt__ApplicationAttribute = int

const Qt__AA_ImmediateWidgetCreation Qt__ApplicationAttribute = 0
const Qt__AA_MSWindowsUseDirect3DByDefault Qt__ApplicationAttribute = 1
const Qt__AA_DontShowIconsInMenus Qt__ApplicationAttribute = 2
const Qt__AA_NativeWindows Qt__ApplicationAttribute = 3
const Qt__AA_DontCreateNativeWidgetSiblings Qt__ApplicationAttribute = 4
const Qt__AA_PluginApplication Qt__ApplicationAttribute = 5
const Qt__AA_MacPluginApplication Qt__ApplicationAttribute = 5
const Qt__AA_DontUseNativeMenuBar Qt__ApplicationAttribute = 6
const Qt__AA_MacDontSwapCtrlAndMeta Qt__ApplicationAttribute = 7
const Qt__AA_Use96Dpi Qt__ApplicationAttribute = 8
const Qt__AA_X11InitThreads Qt__ApplicationAttribute = 10
const Qt__AA_SynthesizeTouchForUnhandledMouseEvents Qt__ApplicationAttribute = 11
const Qt__AA_SynthesizeMouseForUnhandledTouchEvents Qt__ApplicationAttribute = 12
const Qt__AA_UseHighDpiPixmaps Qt__ApplicationAttribute = 13
const Qt__AA_ForceRasterWidgets Qt__ApplicationAttribute = 14
const Qt__AA_UseDesktopOpenGL Qt__ApplicationAttribute = 15
const Qt__AA_UseOpenGLES Qt__ApplicationAttribute = 16
const Qt__AA_UseSoftwareOpenGL Qt__ApplicationAttribute = 17
const Qt__AA_ShareOpenGLContexts Qt__ApplicationAttribute = 18
const Qt__AA_SetPalette Qt__ApplicationAttribute = 19
const Qt__AA_EnableHighDpiScaling Qt__ApplicationAttribute = 20
const Qt__AA_DisableHighDpiScaling Qt__ApplicationAttribute = 21
const Qt__AA_UseStyleSheetPropagationInWidgetStyles Qt__ApplicationAttribute = 22
const Qt__AA_DontUseNativeDialogs Qt__ApplicationAttribute = 23
const Qt__AA_SynthesizeMouseForUnhandledTabletEvents Qt__ApplicationAttribute = 24
const Qt__AA_CompressHighFrequencyEvents Qt__ApplicationAttribute = 25
const Qt__AA_DontCheckOpenGLContextThreadAffinity Qt__ApplicationAttribute = 26
const Qt__AA_DisableShaderDiskCache Qt__ApplicationAttribute = 27
const Qt__AA_DontShowShortcutsInContextMenus Qt__ApplicationAttribute = 28
const Qt__AA_CompressTabletEvents Qt__ApplicationAttribute = 29
const Qt__AA_DisableWindowContextHelpButton Qt__ApplicationAttribute = 30
const Qt__AA_AttributeCount Qt__ApplicationAttribute = 31

type Qt__ImageConversionFlag = int

const Qt__ColorMode_Mask Qt__ImageConversionFlag = 3
const Qt__AutoColor Qt__ImageConversionFlag = 0
const Qt__ColorOnly Qt__ImageConversionFlag = 3
const Qt__MonoOnly Qt__ImageConversionFlag = 2
const Qt__AlphaDither_Mask Qt__ImageConversionFlag = 12
const Qt__ThresholdAlphaDither Qt__ImageConversionFlag = 0
const Qt__OrderedAlphaDither Qt__ImageConversionFlag = 4
const Qt__DiffuseAlphaDither Qt__ImageConversionFlag = 8
const Qt__NoAlpha Qt__ImageConversionFlag = 12
const Qt__Dither_Mask Qt__ImageConversionFlag = 48
const Qt__DiffuseDither Qt__ImageConversionFlag = 0
const Qt__OrderedDither Qt__ImageConversionFlag = 16
const Qt__ThresholdDither Qt__ImageConversionFlag = 32
const Qt__DitherMode_Mask Qt__ImageConversionFlag = 192
const Qt__AutoDither Qt__ImageConversionFlag = 0
const Qt__PreferDither Qt__ImageConversionFlag = 64
const Qt__AvoidDither Qt__ImageConversionFlag = 128
const Qt__NoOpaqueDetection Qt__ImageConversionFlag = 256
const Qt__NoFormatConversion Qt__ImageConversionFlag = 512

type Qt__BGMode = int

const Qt__TransparentMode Qt__BGMode = 0
const Qt__OpaqueMode Qt__BGMode = 1

type Qt__Key = int

const Qt__Key_Escape Qt__Key = 16777216
const Qt__Key_Tab Qt__Key = 16777217
const Qt__Key_Backtab Qt__Key = 16777218
const Qt__Key_Backspace Qt__Key = 16777219
const Qt__Key_Return Qt__Key = 16777220
const Qt__Key_Enter Qt__Key = 16777221
const Qt__Key_Insert Qt__Key = 16777222
const Qt__Key_Delete Qt__Key = 16777223
const Qt__Key_Pause Qt__Key = 16777224
const Qt__Key_Print Qt__Key = 16777225
const Qt__Key_SysReq Qt__Key = 16777226
const Qt__Key_Clear Qt__Key = 16777227
const Qt__Key_Home Qt__Key = 16777232
const Qt__Key_End Qt__Key = 16777233
const Qt__Key_Left Qt__Key = 16777234
const Qt__Key_Up Qt__Key = 16777235
const Qt__Key_Right Qt__Key = 16777236
const Qt__Key_Down Qt__Key = 16777237
const Qt__Key_PageUp Qt__Key = 16777238
const Qt__Key_PageDown Qt__Key = 16777239
const Qt__Key_Shift Qt__Key = 16777248
const Qt__Key_Control Qt__Key = 16777249
const Qt__Key_Meta Qt__Key = 16777250
const Qt__Key_Alt Qt__Key = 16777251
const Qt__Key_CapsLock Qt__Key = 16777252
const Qt__Key_NumLock Qt__Key = 16777253
const Qt__Key_ScrollLock Qt__Key = 16777254
const Qt__Key_F1 Qt__Key = 16777264
const Qt__Key_F2 Qt__Key = 16777265
const Qt__Key_F3 Qt__Key = 16777266
const Qt__Key_F4 Qt__Key = 16777267
const Qt__Key_F5 Qt__Key = 16777268
const Qt__Key_F6 Qt__Key = 16777269
const Qt__Key_F7 Qt__Key = 16777270
const Qt__Key_F8 Qt__Key = 16777271
const Qt__Key_F9 Qt__Key = 16777272
const Qt__Key_F10 Qt__Key = 16777273
const Qt__Key_F11 Qt__Key = 16777274
const Qt__Key_F12 Qt__Key = 16777275
const Qt__Key_F13 Qt__Key = 16777276
const Qt__Key_F14 Qt__Key = 16777277
const Qt__Key_F15 Qt__Key = 16777278
const Qt__Key_F16 Qt__Key = 16777279
const Qt__Key_F17 Qt__Key = 16777280
const Qt__Key_F18 Qt__Key = 16777281
const Qt__Key_F19 Qt__Key = 16777282
const Qt__Key_F20 Qt__Key = 16777283
const Qt__Key_F21 Qt__Key = 16777284
const Qt__Key_F22 Qt__Key = 16777285
const Qt__Key_F23 Qt__Key = 16777286
const Qt__Key_F24 Qt__Key = 16777287
const Qt__Key_F25 Qt__Key = 16777288
const Qt__Key_F26 Qt__Key = 16777289
const Qt__Key_F27 Qt__Key = 16777290
const Qt__Key_F28 Qt__Key = 16777291
const Qt__Key_F29 Qt__Key = 16777292
const Qt__Key_F30 Qt__Key = 16777293
const Qt__Key_F31 Qt__Key = 16777294
const Qt__Key_F32 Qt__Key = 16777295
const Qt__Key_F33 Qt__Key = 16777296
const Qt__Key_F34 Qt__Key = 16777297
const Qt__Key_F35 Qt__Key = 16777298
const Qt__Key_Super_L Qt__Key = 16777299
const Qt__Key_Super_R Qt__Key = 16777300
const Qt__Key_Menu Qt__Key = 16777301
const Qt__Key_Hyper_L Qt__Key = 16777302
const Qt__Key_Hyper_R Qt__Key = 16777303
const Qt__Key_Help Qt__Key = 16777304
const Qt__Key_Direction_L Qt__Key = 16777305
const Qt__Key_Direction_R Qt__Key = 16777312
const Qt__Key_Space Qt__Key = 32
const Qt__Key_Any Qt__Key = 32
const Qt__Key_Exclam Qt__Key = 33
const Qt__Key_QuoteDbl Qt__Key = 34
const Qt__Key_NumberSign Qt__Key = 35
const Qt__Key_Dollar Qt__Key = 36
const Qt__Key_Percent Qt__Key = 37
const Qt__Key_Ampersand Qt__Key = 38
const Qt__Key_Apostrophe Qt__Key = 39
const Qt__Key_ParenLeft Qt__Key = 40
const Qt__Key_ParenRight Qt__Key = 41
const Qt__Key_Asterisk Qt__Key = 42
const Qt__Key_Plus Qt__Key = 43
const Qt__Key_Comma Qt__Key = 44
const Qt__Key_Minus Qt__Key = 45
const Qt__Key_Period Qt__Key = 46
const Qt__Key_Slash Qt__Key = 47
const Qt__Key_0 Qt__Key = 48
const Qt__Key_1 Qt__Key = 49
const Qt__Key_2 Qt__Key = 50
const Qt__Key_3 Qt__Key = 51
const Qt__Key_4 Qt__Key = 52
const Qt__Key_5 Qt__Key = 53
const Qt__Key_6 Qt__Key = 54
const Qt__Key_7 Qt__Key = 55
const Qt__Key_8 Qt__Key = 56
const Qt__Key_9 Qt__Key = 57
const Qt__Key_Colon Qt__Key = 58
const Qt__Key_Semicolon Qt__Key = 59
const Qt__Key_Less Qt__Key = 60
const Qt__Key_Equal Qt__Key = 61
const Qt__Key_Greater Qt__Key = 62
const Qt__Key_Question Qt__Key = 63
const Qt__Key_At Qt__Key = 64
const Qt__Key_A Qt__Key = 65
const Qt__Key_B Qt__Key = 66
const Qt__Key_C Qt__Key = 67
const Qt__Key_D Qt__Key = 68
const Qt__Key_E Qt__Key = 69
const Qt__Key_F Qt__Key = 70
const Qt__Key_G Qt__Key = 71
const Qt__Key_H Qt__Key = 72
const Qt__Key_I Qt__Key = 73
const Qt__Key_J Qt__Key = 74
const Qt__Key_K Qt__Key = 75
const Qt__Key_L Qt__Key = 76
const Qt__Key_M Qt__Key = 77
const Qt__Key_N Qt__Key = 78
const Qt__Key_O Qt__Key = 79
const Qt__Key_P Qt__Key = 80
const Qt__Key_Q Qt__Key = 81
const Qt__Key_R Qt__Key = 82
const Qt__Key_S Qt__Key = 83
const Qt__Key_T Qt__Key = 84
const Qt__Key_U Qt__Key = 85
const Qt__Key_V Qt__Key = 86
const Qt__Key_W Qt__Key = 87
const Qt__Key_X Qt__Key = 88
const Qt__Key_Y Qt__Key = 89
const Qt__Key_Z Qt__Key = 90
const Qt__Key_BracketLeft Qt__Key = 91
const Qt__Key_Backslash Qt__Key = 92
const Qt__Key_BracketRight Qt__Key = 93
const Qt__Key_AsciiCircum Qt__Key = 94
const Qt__Key_Underscore Qt__Key = 95
const Qt__Key_QuoteLeft Qt__Key = 96
const Qt__Key_BraceLeft Qt__Key = 123
const Qt__Key_Bar Qt__Key = 124
const Qt__Key_BraceRight Qt__Key = 125
const Qt__Key_AsciiTilde Qt__Key = 126
const Qt__Key_nobreakspace Qt__Key = 160
const Qt__Key_exclamdown Qt__Key = 161
const Qt__Key_cent Qt__Key = 162
const Qt__Key_sterling Qt__Key = 163
const Qt__Key_currency Qt__Key = 164
const Qt__Key_yen Qt__Key = 165
const Qt__Key_brokenbar Qt__Key = 166
const Qt__Key_section Qt__Key = 167
const Qt__Key_diaeresis Qt__Key = 168
const Qt__Key_copyright Qt__Key = 169
const Qt__Key_ordfeminine Qt__Key = 170
const Qt__Key_guillemotleft Qt__Key = 171
const Qt__Key_notsign Qt__Key = 172
const Qt__Key_hyphen Qt__Key = 173
const Qt__Key_registered Qt__Key = 174
const Qt__Key_macron Qt__Key = 175
const Qt__Key_degree Qt__Key = 176
const Qt__Key_plusminus Qt__Key = 177
const Qt__Key_twosuperior Qt__Key = 178
const Qt__Key_threesuperior Qt__Key = 179
const Qt__Key_acute Qt__Key = 180
const Qt__Key_mu Qt__Key = 181
const Qt__Key_paragraph Qt__Key = 182
const Qt__Key_periodcentered Qt__Key = 183
const Qt__Key_cedilla Qt__Key = 184
const Qt__Key_onesuperior Qt__Key = 185
const Qt__Key_masculine Qt__Key = 186
const Qt__Key_guillemotright Qt__Key = 187
const Qt__Key_onequarter Qt__Key = 188
const Qt__Key_onehalf Qt__Key = 189
const Qt__Key_threequarters Qt__Key = 190
const Qt__Key_questiondown Qt__Key = 191
const Qt__Key_Agrave Qt__Key = 192
const Qt__Key_Aacute Qt__Key = 193
const Qt__Key_Acircumflex Qt__Key = 194
const Qt__Key_Atilde Qt__Key = 195
const Qt__Key_Adiaeresis Qt__Key = 196
const Qt__Key_Aring Qt__Key = 197
const Qt__Key_AE Qt__Key = 198
const Qt__Key_Ccedilla Qt__Key = 199
const Qt__Key_Egrave Qt__Key = 200
const Qt__Key_Eacute Qt__Key = 201
const Qt__Key_Ecircumflex Qt__Key = 202
const Qt__Key_Ediaeresis Qt__Key = 203
const Qt__Key_Igrave Qt__Key = 204
const Qt__Key_Iacute Qt__Key = 205
const Qt__Key_Icircumflex Qt__Key = 206
const Qt__Key_Idiaeresis Qt__Key = 207
const Qt__Key_ETH Qt__Key = 208
const Qt__Key_Ntilde Qt__Key = 209
const Qt__Key_Ograve Qt__Key = 210
const Qt__Key_Oacute Qt__Key = 211
const Qt__Key_Ocircumflex Qt__Key = 212
const Qt__Key_Otilde Qt__Key = 213
const Qt__Key_Odiaeresis Qt__Key = 214
const Qt__Key_multiply Qt__Key = 215
const Qt__Key_Ooblique Qt__Key = 216
const Qt__Key_Ugrave Qt__Key = 217
const Qt__Key_Uacute Qt__Key = 218
const Qt__Key_Ucircumflex Qt__Key = 219
const Qt__Key_Udiaeresis Qt__Key = 220
const Qt__Key_Yacute Qt__Key = 221
const Qt__Key_THORN Qt__Key = 222
const Qt__Key_ssharp Qt__Key = 223
const Qt__Key_division Qt__Key = 247
const Qt__Key_ydiaeresis Qt__Key = 255
const Qt__Key_AltGr Qt__Key = 16781571
const Qt__Key_Multi_key Qt__Key = 16781600
const Qt__Key_Codeinput Qt__Key = 16781623
const Qt__Key_SingleCandidate Qt__Key = 16781628
const Qt__Key_MultipleCandidate Qt__Key = 16781629
const Qt__Key_PreviousCandidate Qt__Key = 16781630
const Qt__Key_Mode_switch Qt__Key = 16781694
const Qt__Key_Kanji Qt__Key = 16781601
const Qt__Key_Muhenkan Qt__Key = 16781602
const Qt__Key_Henkan Qt__Key = 16781603
const Qt__Key_Romaji Qt__Key = 16781604
const Qt__Key_Hiragana Qt__Key = 16781605
const Qt__Key_Katakana Qt__Key = 16781606
const Qt__Key_Hiragana_Katakana Qt__Key = 16781607
const Qt__Key_Zenkaku Qt__Key = 16781608
const Qt__Key_Hankaku Qt__Key = 16781609
const Qt__Key_Zenkaku_Hankaku Qt__Key = 16781610
const Qt__Key_Touroku Qt__Key = 16781611
const Qt__Key_Massyo Qt__Key = 16781612
const Qt__Key_Kana_Lock Qt__Key = 16781613
const Qt__Key_Kana_Shift Qt__Key = 16781614
const Qt__Key_Eisu_Shift Qt__Key = 16781615
const Qt__Key_Eisu_toggle Qt__Key = 16781616
const Qt__Key_Hangul Qt__Key = 16781617
const Qt__Key_Hangul_Start Qt__Key = 16781618
const Qt__Key_Hangul_End Qt__Key = 16781619
const Qt__Key_Hangul_Hanja Qt__Key = 16781620
const Qt__Key_Hangul_Jamo Qt__Key = 16781621
const Qt__Key_Hangul_Romaja Qt__Key = 16781622
const Qt__Key_Hangul_Jeonja Qt__Key = 16781624
const Qt__Key_Hangul_Banja Qt__Key = 16781625
const Qt__Key_Hangul_PreHanja Qt__Key = 16781626
const Qt__Key_Hangul_PostHanja Qt__Key = 16781627
const Qt__Key_Hangul_Special Qt__Key = 16781631
const Qt__Key_Dead_Grave Qt__Key = 16781904
const Qt__Key_Dead_Acute Qt__Key = 16781905
const Qt__Key_Dead_Circumflex Qt__Key = 16781906
const Qt__Key_Dead_Tilde Qt__Key = 16781907
const Qt__Key_Dead_Macron Qt__Key = 16781908
const Qt__Key_Dead_Breve Qt__Key = 16781909
const Qt__Key_Dead_Abovedot Qt__Key = 16781910
const Qt__Key_Dead_Diaeresis Qt__Key = 16781911
const Qt__Key_Dead_Abovering Qt__Key = 16781912
const Qt__Key_Dead_Doubleacute Qt__Key = 16781913
const Qt__Key_Dead_Caron Qt__Key = 16781914
const Qt__Key_Dead_Cedilla Qt__Key = 16781915
const Qt__Key_Dead_Ogonek Qt__Key = 16781916
const Qt__Key_Dead_Iota Qt__Key = 16781917
const Qt__Key_Dead_Voiced_Sound Qt__Key = 16781918
const Qt__Key_Dead_Semivoiced_Sound Qt__Key = 16781919
const Qt__Key_Dead_Belowdot Qt__Key = 16781920
const Qt__Key_Dead_Hook Qt__Key = 16781921
const Qt__Key_Dead_Horn Qt__Key = 16781922
const Qt__Key_Back Qt__Key = 16777313
const Qt__Key_Forward Qt__Key = 16777314
const Qt__Key_Stop Qt__Key = 16777315
const Qt__Key_Refresh Qt__Key = 16777316
const Qt__Key_VolumeDown Qt__Key = 16777328
const Qt__Key_VolumeMute Qt__Key = 16777329
const Qt__Key_VolumeUp Qt__Key = 16777330
const Qt__Key_BassBoost Qt__Key = 16777331
const Qt__Key_BassUp Qt__Key = 16777332
const Qt__Key_BassDown Qt__Key = 16777333
const Qt__Key_TrebleUp Qt__Key = 16777334
const Qt__Key_TrebleDown Qt__Key = 16777335
const Qt__Key_MediaPlay Qt__Key = 16777344
const Qt__Key_MediaStop Qt__Key = 16777345
const Qt__Key_MediaPrevious Qt__Key = 16777346
const Qt__Key_MediaNext Qt__Key = 16777347
const Qt__Key_MediaRecord Qt__Key = 16777348
const Qt__Key_MediaPause Qt__Key = 16777349
const Qt__Key_MediaTogglePlayPause Qt__Key = 16777350
const Qt__Key_HomePage Qt__Key = 16777360
const Qt__Key_Favorites Qt__Key = 16777361
const Qt__Key_Search Qt__Key = 16777362
const Qt__Key_Standby Qt__Key = 16777363
const Qt__Key_OpenUrl Qt__Key = 16777364
const Qt__Key_LaunchMail Qt__Key = 16777376
const Qt__Key_LaunchMedia Qt__Key = 16777377
const Qt__Key_Launch0 Qt__Key = 16777378
const Qt__Key_Launch1 Qt__Key = 16777379
const Qt__Key_Launch2 Qt__Key = 16777380
const Qt__Key_Launch3 Qt__Key = 16777381
const Qt__Key_Launch4 Qt__Key = 16777382
const Qt__Key_Launch5 Qt__Key = 16777383
const Qt__Key_Launch6 Qt__Key = 16777384
const Qt__Key_Launch7 Qt__Key = 16777385
const Qt__Key_Launch8 Qt__Key = 16777386
const Qt__Key_Launch9 Qt__Key = 16777387
const Qt__Key_LaunchA Qt__Key = 16777388
const Qt__Key_LaunchB Qt__Key = 16777389
const Qt__Key_LaunchC Qt__Key = 16777390
const Qt__Key_LaunchD Qt__Key = 16777391
const Qt__Key_LaunchE Qt__Key = 16777392
const Qt__Key_LaunchF Qt__Key = 16777393
const Qt__Key_MonBrightnessUp Qt__Key = 16777394
const Qt__Key_MonBrightnessDown Qt__Key = 16777395
const Qt__Key_KeyboardLightOnOff Qt__Key = 16777396
const Qt__Key_KeyboardBrightnessUp Qt__Key = 16777397
const Qt__Key_KeyboardBrightnessDown Qt__Key = 16777398
const Qt__Key_PowerOff Qt__Key = 16777399
const Qt__Key_WakeUp Qt__Key = 16777400
const Qt__Key_Eject Qt__Key = 16777401
const Qt__Key_ScreenSaver Qt__Key = 16777402
const Qt__Key_WWW Qt__Key = 16777403
const Qt__Key_Memo Qt__Key = 16777404
const Qt__Key_LightBulb Qt__Key = 16777405
const Qt__Key_Shop Qt__Key = 16777406
const Qt__Key_History Qt__Key = 16777407
const Qt__Key_AddFavorite Qt__Key = 16777408
const Qt__Key_HotLinks Qt__Key = 16777409
const Qt__Key_BrightnessAdjust Qt__Key = 16777410
const Qt__Key_Finance Qt__Key = 16777411
const Qt__Key_Community Qt__Key = 16777412
const Qt__Key_AudioRewind Qt__Key = 16777413
const Qt__Key_BackForward Qt__Key = 16777414
const Qt__Key_ApplicationLeft Qt__Key = 16777415
const Qt__Key_ApplicationRight Qt__Key = 16777416
const Qt__Key_Book Qt__Key = 16777417
const Qt__Key_CD Qt__Key = 16777418
const Qt__Key_Calculator Qt__Key = 16777419
const Qt__Key_ToDoList Qt__Key = 16777420
const Qt__Key_ClearGrab Qt__Key = 16777421
const Qt__Key_Close Qt__Key = 16777422
const Qt__Key_Copy Qt__Key = 16777423
const Qt__Key_Cut Qt__Key = 16777424
const Qt__Key_Display Qt__Key = 16777425
const Qt__Key_DOS Qt__Key = 16777426
const Qt__Key_Documents Qt__Key = 16777427
const Qt__Key_Excel Qt__Key = 16777428
const Qt__Key_Explorer Qt__Key = 16777429
const Qt__Key_Game Qt__Key = 16777430
const Qt__Key_Go Qt__Key = 16777431
const Qt__Key_iTouch Qt__Key = 16777432
const Qt__Key_LogOff Qt__Key = 16777433
const Qt__Key_Market Qt__Key = 16777434
const Qt__Key_Meeting Qt__Key = 16777435
const Qt__Key_MenuKB Qt__Key = 16777436
const Qt__Key_MenuPB Qt__Key = 16777437
const Qt__Key_MySites Qt__Key = 16777438
const Qt__Key_News Qt__Key = 16777439
const Qt__Key_OfficeHome Qt__Key = 16777440
const Qt__Key_Option Qt__Key = 16777441
const Qt__Key_Paste Qt__Key = 16777442
const Qt__Key_Phone Qt__Key = 16777443
const Qt__Key_Calendar Qt__Key = 16777444
const Qt__Key_Reply Qt__Key = 16777445
const Qt__Key_Reload Qt__Key = 16777446
const Qt__Key_RotateWindows Qt__Key = 16777447
const Qt__Key_RotationPB Qt__Key = 16777448
const Qt__Key_RotationKB Qt__Key = 16777449
const Qt__Key_Save Qt__Key = 16777450
const Qt__Key_Send Qt__Key = 16777451
const Qt__Key_Spell Qt__Key = 16777452
const Qt__Key_SplitScreen Qt__Key = 16777453
const Qt__Key_Support Qt__Key = 16777454
const Qt__Key_TaskPane Qt__Key = 16777455
const Qt__Key_Terminal Qt__Key = 16777456
const Qt__Key_Tools Qt__Key = 16777457
const Qt__Key_Travel Qt__Key = 16777458
const Qt__Key_Video Qt__Key = 16777459
const Qt__Key_Word Qt__Key = 16777460
const Qt__Key_Xfer Qt__Key = 16777461
const Qt__Key_ZoomIn Qt__Key = 16777462
const Qt__Key_ZoomOut Qt__Key = 16777463
const Qt__Key_Away Qt__Key = 16777464
const Qt__Key_Messenger Qt__Key = 16777465
const Qt__Key_WebCam Qt__Key = 16777466
const Qt__Key_MailForward Qt__Key = 16777467
const Qt__Key_Pictures Qt__Key = 16777468
const Qt__Key_Music Qt__Key = 16777469
const Qt__Key_Battery Qt__Key = 16777470
const Qt__Key_Bluetooth Qt__Key = 16777471
const Qt__Key_WLAN Qt__Key = 16777472
const Qt__Key_UWB Qt__Key = 16777473
const Qt__Key_AudioForward Qt__Key = 16777474
const Qt__Key_AudioRepeat Qt__Key = 16777475
const Qt__Key_AudioRandomPlay Qt__Key = 16777476
const Qt__Key_Subtitle Qt__Key = 16777477
const Qt__Key_AudioCycleTrack Qt__Key = 16777478
const Qt__Key_Time Qt__Key = 16777479
const Qt__Key_Hibernate Qt__Key = 16777480
const Qt__Key_View Qt__Key = 16777481
const Qt__Key_TopMenu Qt__Key = 16777482
const Qt__Key_PowerDown Qt__Key = 16777483
const Qt__Key_Suspend Qt__Key = 16777484
const Qt__Key_ContrastAdjust Qt__Key = 16777485
const Qt__Key_LaunchG Qt__Key = 16777486
const Qt__Key_LaunchH Qt__Key = 16777487
const Qt__Key_TouchpadToggle Qt__Key = 16777488
const Qt__Key_TouchpadOn Qt__Key = 16777489
const Qt__Key_TouchpadOff Qt__Key = 16777490
const Qt__Key_MicMute Qt__Key = 16777491
const Qt__Key_Red Qt__Key = 16777492
const Qt__Key_Green Qt__Key = 16777493
const Qt__Key_Yellow Qt__Key = 16777494
const Qt__Key_Blue Qt__Key = 16777495
const Qt__Key_ChannelUp Qt__Key = 16777496
const Qt__Key_ChannelDown Qt__Key = 16777497
const Qt__Key_Guide Qt__Key = 16777498
const Qt__Key_Info Qt__Key = 16777499
const Qt__Key_Settings Qt__Key = 16777500
const Qt__Key_MicVolumeUp Qt__Key = 16777501
const Qt__Key_MicVolumeDown Qt__Key = 16777502
const Qt__Key_New Qt__Key = 16777504
const Qt__Key_Open Qt__Key = 16777505
const Qt__Key_Find Qt__Key = 16777506
const Qt__Key_Undo Qt__Key = 16777507
const Qt__Key_Redo Qt__Key = 16777508
const Qt__Key_MediaLast Qt__Key = 16842751
const Qt__Key_Select Qt__Key = 16842752
const Qt__Key_Yes Qt__Key = 16842753
const Qt__Key_No Qt__Key = 16842754
const Qt__Key_Cancel Qt__Key = 16908289
const Qt__Key_Printer Qt__Key = 16908290
const Qt__Key_Execute Qt__Key = 16908291
const Qt__Key_Sleep Qt__Key = 16908292
const Qt__Key_Play Qt__Key = 16908293
const Qt__Key_Zoom Qt__Key = 16908294
const Qt__Key_Exit Qt__Key = 16908298
const Qt__Key_Context1 Qt__Key = 17825792
const Qt__Key_Context2 Qt__Key = 17825793
const Qt__Key_Context3 Qt__Key = 17825794
const Qt__Key_Context4 Qt__Key = 17825795
const Qt__Key_Call Qt__Key = 17825796
const Qt__Key_Hangup Qt__Key = 17825797
const Qt__Key_Flip Qt__Key = 17825798
const Qt__Key_ToggleCallHangup Qt__Key = 17825799
const Qt__Key_VoiceDial Qt__Key = 17825800
const Qt__Key_LastNumberRedial Qt__Key = 17825801
const Qt__Key_Camera Qt__Key = 17825824
const Qt__Key_CameraFocus Qt__Key = 17825825
const Qt__Key_unknown Qt__Key = 33554431

type Qt__ArrowType = int

const Qt__NoArrow Qt__ArrowType = 0
const Qt__UpArrow Qt__ArrowType = 1
const Qt__DownArrow Qt__ArrowType = 2
const Qt__LeftArrow Qt__ArrowType = 3
const Qt__RightArrow Qt__ArrowType = 4

type Qt__PenStyle = int

const Qt__NoPen Qt__PenStyle = 0
const Qt__SolidLine Qt__PenStyle = 1
const Qt__DashLine Qt__PenStyle = 2
const Qt__DotLine Qt__PenStyle = 3
const Qt__DashDotLine Qt__PenStyle = 4
const Qt__DashDotDotLine Qt__PenStyle = 5
const Qt__CustomDashLine Qt__PenStyle = 6
const Qt__MPenStyle Qt__PenStyle = 15

type Qt__PenCapStyle = int

const Qt__FlatCap Qt__PenCapStyle = 0
const Qt__SquareCap Qt__PenCapStyle = 16
const Qt__RoundCap Qt__PenCapStyle = 32
const Qt__MPenCapStyle Qt__PenCapStyle = 48

type Qt__PenJoinStyle = int

const Qt__MiterJoin Qt__PenJoinStyle = 0
const Qt__BevelJoin Qt__PenJoinStyle = 64
const Qt__RoundJoin Qt__PenJoinStyle = 128
const Qt__SvgMiterJoin Qt__PenJoinStyle = 256
const Qt__MPenJoinStyle Qt__PenJoinStyle = 448

type Qt__BrushStyle = int

const Qt__NoBrush Qt__BrushStyle = 0
const Qt__SolidPattern Qt__BrushStyle = 1
const Qt__Dense1Pattern Qt__BrushStyle = 2
const Qt__Dense2Pattern Qt__BrushStyle = 3
const Qt__Dense3Pattern Qt__BrushStyle = 4
const Qt__Dense4Pattern Qt__BrushStyle = 5
const Qt__Dense5Pattern Qt__BrushStyle = 6
const Qt__Dense6Pattern Qt__BrushStyle = 7
const Qt__Dense7Pattern Qt__BrushStyle = 8
const Qt__HorPattern Qt__BrushStyle = 9
const Qt__VerPattern Qt__BrushStyle = 10
const Qt__CrossPattern Qt__BrushStyle = 11
const Qt__BDiagPattern Qt__BrushStyle = 12
const Qt__FDiagPattern Qt__BrushStyle = 13
const Qt__DiagCrossPattern Qt__BrushStyle = 14
const Qt__LinearGradientPattern Qt__BrushStyle = 15
const Qt__RadialGradientPattern Qt__BrushStyle = 16
const Qt__ConicalGradientPattern Qt__BrushStyle = 17
const Qt__TexturePattern Qt__BrushStyle = 24

type Qt__SizeMode = int

const Qt__AbsoluteSize Qt__SizeMode = 0
const Qt__RelativeSize Qt__SizeMode = 1

type Qt__UIEffect = int

const Qt__UI_General Qt__UIEffect = 0
const Qt__UI_AnimateMenu Qt__UIEffect = 1
const Qt__UI_FadeMenu Qt__UIEffect = 2
const Qt__UI_AnimateCombo Qt__UIEffect = 3
const Qt__UI_AnimateTooltip Qt__UIEffect = 4
const Qt__UI_FadeTooltip Qt__UIEffect = 5
const Qt__UI_AnimateToolBox Qt__UIEffect = 6

type Qt__CursorShape = int

const Qt__ArrowCursor Qt__CursorShape = 0
const Qt__UpArrowCursor Qt__CursorShape = 1
const Qt__CrossCursor Qt__CursorShape = 2
const Qt__WaitCursor Qt__CursorShape = 3
const Qt__IBeamCursor Qt__CursorShape = 4
const Qt__SizeVerCursor Qt__CursorShape = 5
const Qt__SizeHorCursor Qt__CursorShape = 6
const Qt__SizeBDiagCursor Qt__CursorShape = 7
const Qt__SizeFDiagCursor Qt__CursorShape = 8
const Qt__SizeAllCursor Qt__CursorShape = 9
const Qt__BlankCursor Qt__CursorShape = 10
const Qt__SplitVCursor Qt__CursorShape = 11
const Qt__SplitHCursor Qt__CursorShape = 12
const Qt__PointingHandCursor Qt__CursorShape = 13
const Qt__ForbiddenCursor Qt__CursorShape = 14
const Qt__WhatsThisCursor Qt__CursorShape = 15
const Qt__BusyCursor Qt__CursorShape = 16
const Qt__OpenHandCursor Qt__CursorShape = 17
const Qt__ClosedHandCursor Qt__CursorShape = 18
const Qt__DragCopyCursor Qt__CursorShape = 19
const Qt__DragMoveCursor Qt__CursorShape = 20
const Qt__DragLinkCursor Qt__CursorShape = 21
const Qt__LastCursor Qt__CursorShape = 21
const Qt__BitmapCursor Qt__CursorShape = 24
const Qt__CustomCursor Qt__CursorShape = 25

type Qt__TextFormat = int

const Qt__PlainText Qt__TextFormat = 0
const Qt__RichText Qt__TextFormat = 1
const Qt__AutoText Qt__TextFormat = 2

type Qt__AspectRatioMode = int

const Qt__IgnoreAspectRatio Qt__AspectRatioMode = 0
const Qt__KeepAspectRatio Qt__AspectRatioMode = 1
const Qt__KeepAspectRatioByExpanding Qt__AspectRatioMode = 2

type Qt__DockWidgetArea = int

const Qt__LeftDockWidgetArea Qt__DockWidgetArea = 1
const Qt__RightDockWidgetArea Qt__DockWidgetArea = 2
const Qt__TopDockWidgetArea Qt__DockWidgetArea = 4
const Qt__BottomDockWidgetArea Qt__DockWidgetArea = 8
const Qt__DockWidgetArea_Mask Qt__DockWidgetArea = 15
const Qt__AllDockWidgetAreas Qt__DockWidgetArea = 15
const Qt__NoDockWidgetArea Qt__DockWidgetArea = 0

type Qt__DockWidgetAreaSizes = int

const Qt__NDockWidgetAreas Qt__DockWidgetAreaSizes = 4

type Qt__ToolBarArea = int

const Qt__LeftToolBarArea Qt__ToolBarArea = 1
const Qt__RightToolBarArea Qt__ToolBarArea = 2
const Qt__TopToolBarArea Qt__ToolBarArea = 4
const Qt__BottomToolBarArea Qt__ToolBarArea = 8
const Qt__ToolBarArea_Mask Qt__ToolBarArea = 15
const Qt__AllToolBarAreas Qt__ToolBarArea = 15
const Qt__NoToolBarArea Qt__ToolBarArea = 0

type Qt__ToolBarAreaSizes = int

const Qt__NToolBarAreas Qt__ToolBarAreaSizes = 4

type Qt__DateFormat = int

const Qt__TextDate Qt__DateFormat = 0
const Qt__ISODate Qt__DateFormat = 1
const Qt__SystemLocaleDate Qt__DateFormat = 2
const Qt__LocalDate Qt__DateFormat = 2
const Qt__LocaleDate Qt__DateFormat = 3
const Qt__SystemLocaleShortDate Qt__DateFormat = 4
const Qt__SystemLocaleLongDate Qt__DateFormat = 5
const Qt__DefaultLocaleShortDate Qt__DateFormat = 6
const Qt__DefaultLocaleLongDate Qt__DateFormat = 7
const Qt__RFC2822Date Qt__DateFormat = 8
const Qt__ISODateWithMs Qt__DateFormat = 9

type Qt__TimeSpec = int

const Qt__LocalTime Qt__TimeSpec = 0
const Qt__UTC Qt__TimeSpec = 1
const Qt__OffsetFromUTC Qt__TimeSpec = 2
const Qt__TimeZone Qt__TimeSpec = 3

type Qt__DayOfWeek = int

const Qt__Monday Qt__DayOfWeek = 1
const Qt__Tuesday Qt__DayOfWeek = 2
const Qt__Wednesday Qt__DayOfWeek = 3
const Qt__Thursday Qt__DayOfWeek = 4
const Qt__Friday Qt__DayOfWeek = 5
const Qt__Saturday Qt__DayOfWeek = 6
const Qt__Sunday Qt__DayOfWeek = 7

type Qt__ScrollBarPolicy = int

const Qt__ScrollBarAsNeeded Qt__ScrollBarPolicy = 0
const Qt__ScrollBarAlwaysOff Qt__ScrollBarPolicy = 1
const Qt__ScrollBarAlwaysOn Qt__ScrollBarPolicy = 2

type Qt__CaseSensitivity = int

const Qt__CaseInsensitive Qt__CaseSensitivity = 0
const Qt__CaseSensitive Qt__CaseSensitivity = 1

type Qt__Corner = int

const Qt__TopLeftCorner Qt__Corner = 0
const Qt__TopRightCorner Qt__Corner = 1
const Qt__BottomLeftCorner Qt__Corner = 2
const Qt__BottomRightCorner Qt__Corner = 3

type Qt__Edge = int

const Qt__TopEdge Qt__Edge = 1
const Qt__LeftEdge Qt__Edge = 2
const Qt__RightEdge Qt__Edge = 4
const Qt__BottomEdge Qt__Edge = 8

type Qt__ConnectionType = int

const Qt__AutoConnection Qt__ConnectionType = 0
const Qt__DirectConnection Qt__ConnectionType = 1
const Qt__QueuedConnection Qt__ConnectionType = 2
const Qt__BlockingQueuedConnection Qt__ConnectionType = 3
const Qt__UniqueConnection Qt__ConnectionType = 128

type Qt__ShortcutContext = int

const Qt__WidgetShortcut Qt__ShortcutContext = 0
const Qt__WindowShortcut Qt__ShortcutContext = 1
const Qt__ApplicationShortcut Qt__ShortcutContext = 2
const Qt__WidgetWithChildrenShortcut Qt__ShortcutContext = 3

type Qt__FillRule = int

const Qt__OddEvenFill Qt__FillRule = 0
const Qt__WindingFill Qt__FillRule = 1

type Qt__MaskMode = int

const Qt__MaskInColor Qt__MaskMode = 0
const Qt__MaskOutColor Qt__MaskMode = 1

type Qt__ClipOperation = int

const Qt__NoClip Qt__ClipOperation = 0
const Qt__ReplaceClip Qt__ClipOperation = 1
const Qt__IntersectClip Qt__ClipOperation = 2

type Qt__ItemSelectionMode = int

const Qt__ContainsItemShape Qt__ItemSelectionMode = 0
const Qt__IntersectsItemShape Qt__ItemSelectionMode = 1
const Qt__ContainsItemBoundingRect Qt__ItemSelectionMode = 2
const Qt__IntersectsItemBoundingRect Qt__ItemSelectionMode = 3

type Qt__ItemSelectionOperation = int

const Qt__ReplaceSelection Qt__ItemSelectionOperation = 0
const Qt__AddToSelection Qt__ItemSelectionOperation = 1

type Qt__TransformationMode = int

const Qt__FastTransformation Qt__TransformationMode = 0
const Qt__SmoothTransformation Qt__TransformationMode = 1

type Qt__Axis = int

const Qt__XAxis Qt__Axis = 0
const Qt__YAxis Qt__Axis = 1
const Qt__ZAxis Qt__Axis = 2

type Qt__FocusReason = int

const Qt__MouseFocusReason Qt__FocusReason = 0
const Qt__TabFocusReason Qt__FocusReason = 1
const Qt__BacktabFocusReason Qt__FocusReason = 2
const Qt__ActiveWindowFocusReason Qt__FocusReason = 3
const Qt__PopupFocusReason Qt__FocusReason = 4
const Qt__ShortcutFocusReason Qt__FocusReason = 5
const Qt__MenuBarFocusReason Qt__FocusReason = 6
const Qt__OtherFocusReason Qt__FocusReason = 7
const Qt__NoFocusReason Qt__FocusReason = 8

type Qt__ContextMenuPolicy = int

const Qt__NoContextMenu Qt__ContextMenuPolicy = 0
const Qt__DefaultContextMenu Qt__ContextMenuPolicy = 1
const Qt__ActionsContextMenu Qt__ContextMenuPolicy = 2
const Qt__CustomContextMenu Qt__ContextMenuPolicy = 3
const Qt__PreventContextMenu Qt__ContextMenuPolicy = 4

type Qt__InputMethodQuery = int

const Qt__ImEnabled Qt__InputMethodQuery = 1
const Qt__ImCursorRectangle Qt__InputMethodQuery = 2
const Qt__ImMicroFocus Qt__InputMethodQuery = 2
const Qt__ImFont Qt__InputMethodQuery = 4
const Qt__ImCursorPosition Qt__InputMethodQuery = 8
const Qt__ImSurroundingText Qt__InputMethodQuery = 16
const Qt__ImCurrentSelection Qt__InputMethodQuery = 32
const Qt__ImMaximumTextLength Qt__InputMethodQuery = 64
const Qt__ImAnchorPosition Qt__InputMethodQuery = 128
const Qt__ImHints Qt__InputMethodQuery = 256
const Qt__ImPreferredLanguage Qt__InputMethodQuery = 512
const Qt__ImAbsolutePosition Qt__InputMethodQuery = 1024
const Qt__ImTextBeforeCursor Qt__InputMethodQuery = 2048
const Qt__ImTextAfterCursor Qt__InputMethodQuery = 4096
const Qt__ImEnterKeyType Qt__InputMethodQuery = 8192
const Qt__ImAnchorRectangle Qt__InputMethodQuery = 16384
const Qt__ImInputItemClipRectangle Qt__InputMethodQuery = 32768
const Qt__ImPlatformData Qt__InputMethodQuery = -2147483648
const Qt__ImQueryInput Qt__InputMethodQuery = 16570
const Qt__ImQueryAll Qt__InputMethodQuery = -1

type Qt__InputMethodHint = int

const Qt__ImhNone Qt__InputMethodHint = 0
const Qt__ImhHiddenText Qt__InputMethodHint = 1
const Qt__ImhSensitiveData Qt__InputMethodHint = 2
const Qt__ImhNoAutoUppercase Qt__InputMethodHint = 4
const Qt__ImhPreferNumbers Qt__InputMethodHint = 8
const Qt__ImhPreferUppercase Qt__InputMethodHint = 16
const Qt__ImhPreferLowercase Qt__InputMethodHint = 32
const Qt__ImhNoPredictiveText Qt__InputMethodHint = 64
const Qt__ImhDate Qt__InputMethodHint = 128
const Qt__ImhTime Qt__InputMethodHint = 256
const Qt__ImhPreferLatin Qt__InputMethodHint = 512
const Qt__ImhMultiLine Qt__InputMethodHint = 1024
const Qt__ImhDigitsOnly Qt__InputMethodHint = 65536
const Qt__ImhFormattedNumbersOnly Qt__InputMethodHint = 131072
const Qt__ImhUppercaseOnly Qt__InputMethodHint = 262144
const Qt__ImhLowercaseOnly Qt__InputMethodHint = 524288
const Qt__ImhDialableCharactersOnly Qt__InputMethodHint = 1048576
const Qt__ImhEmailCharactersOnly Qt__InputMethodHint = 2097152
const Qt__ImhUrlCharactersOnly Qt__InputMethodHint = 4194304
const Qt__ImhLatinOnly Qt__InputMethodHint = 8388608
const Qt__ImhExclusiveInputMask Qt__InputMethodHint = -65536

type Qt__EnterKeyType = int

const Qt__EnterKeyDefault Qt__EnterKeyType = 0
const Qt__EnterKeyReturn Qt__EnterKeyType = 1
const Qt__EnterKeyDone Qt__EnterKeyType = 2
const Qt__EnterKeyGo Qt__EnterKeyType = 3
const Qt__EnterKeySend Qt__EnterKeyType = 4
const Qt__EnterKeySearch Qt__EnterKeyType = 5
const Qt__EnterKeyNext Qt__EnterKeyType = 6
const Qt__EnterKeyPrevious Qt__EnterKeyType = 7

type Qt__ToolButtonStyle = int

const Qt__ToolButtonIconOnly Qt__ToolButtonStyle = 0
const Qt__ToolButtonTextOnly Qt__ToolButtonStyle = 1
const Qt__ToolButtonTextBesideIcon Qt__ToolButtonStyle = 2
const Qt__ToolButtonTextUnderIcon Qt__ToolButtonStyle = 3
const Qt__ToolButtonFollowStyle Qt__ToolButtonStyle = 4

type Qt__LayoutDirection = int

const Qt__LeftToRight Qt__LayoutDirection = 0
const Qt__RightToLeft Qt__LayoutDirection = 1
const Qt__LayoutDirectionAuto Qt__LayoutDirection = 2

type Qt__AnchorPoint = int

const Qt__AnchorLeft Qt__AnchorPoint = 0
const Qt__AnchorHorizontalCenter Qt__AnchorPoint = 1
const Qt__AnchorRight Qt__AnchorPoint = 2
const Qt__AnchorTop Qt__AnchorPoint = 3
const Qt__AnchorVerticalCenter Qt__AnchorPoint = 4
const Qt__AnchorBottom Qt__AnchorPoint = 5

type Qt__FindChildOption = int

const Qt__FindDirectChildrenOnly Qt__FindChildOption = 0
const Qt__FindChildrenRecursively Qt__FindChildOption = 1

type Qt__DropAction = int

const Qt__CopyAction Qt__DropAction = 1
const Qt__MoveAction Qt__DropAction = 2
const Qt__LinkAction Qt__DropAction = 4
const Qt__ActionMask Qt__DropAction = 255
const Qt__TargetMoveAction Qt__DropAction = 32770
const Qt__IgnoreAction Qt__DropAction = 0

type Qt__CheckState = int

const Qt__Unchecked Qt__CheckState = 0
const Qt__PartiallyChecked Qt__CheckState = 1
const Qt__Checked Qt__CheckState = 2

type Qt__ItemDataRole = int

const Qt__DisplayRole Qt__ItemDataRole = 0
const Qt__DecorationRole Qt__ItemDataRole = 1
const Qt__EditRole Qt__ItemDataRole = 2
const Qt__ToolTipRole Qt__ItemDataRole = 3
const Qt__StatusTipRole Qt__ItemDataRole = 4
const Qt__WhatsThisRole Qt__ItemDataRole = 5
const Qt__FontRole Qt__ItemDataRole = 6
const Qt__TextAlignmentRole Qt__ItemDataRole = 7
const Qt__BackgroundColorRole Qt__ItemDataRole = 8
const Qt__BackgroundRole Qt__ItemDataRole = 8
const Qt__TextColorRole Qt__ItemDataRole = 9
const Qt__ForegroundRole Qt__ItemDataRole = 9
const Qt__CheckStateRole Qt__ItemDataRole = 10
const Qt__AccessibleTextRole Qt__ItemDataRole = 11
const Qt__AccessibleDescriptionRole Qt__ItemDataRole = 12
const Qt__SizeHintRole Qt__ItemDataRole = 13
const Qt__InitialSortOrderRole Qt__ItemDataRole = 14
const Qt__DisplayPropertyRole Qt__ItemDataRole = 27
const Qt__DecorationPropertyRole Qt__ItemDataRole = 28
const Qt__ToolTipPropertyRole Qt__ItemDataRole = 29
const Qt__StatusTipPropertyRole Qt__ItemDataRole = 30
const Qt__WhatsThisPropertyRole Qt__ItemDataRole = 31
const Qt__UserRole Qt__ItemDataRole = 256

type Qt__ItemFlag = int

const Qt__NoItemFlags Qt__ItemFlag = 0
const Qt__ItemIsSelectable Qt__ItemFlag = 1
const Qt__ItemIsEditable Qt__ItemFlag = 2
const Qt__ItemIsDragEnabled Qt__ItemFlag = 4
const Qt__ItemIsDropEnabled Qt__ItemFlag = 8
const Qt__ItemIsUserCheckable Qt__ItemFlag = 16
const Qt__ItemIsEnabled Qt__ItemFlag = 32
const Qt__ItemIsAutoTristate Qt__ItemFlag = 64
const Qt__ItemIsTristate Qt__ItemFlag = 64
const Qt__ItemNeverHasChildren Qt__ItemFlag = 128
const Qt__ItemIsUserTristate Qt__ItemFlag = 256

type Qt__MatchFlag = int

const Qt__MatchExactly Qt__MatchFlag = 0
const Qt__MatchContains Qt__MatchFlag = 1
const Qt__MatchStartsWith Qt__MatchFlag = 2
const Qt__MatchEndsWith Qt__MatchFlag = 3
const Qt__MatchRegExp Qt__MatchFlag = 4
const Qt__MatchWildcard Qt__MatchFlag = 5
const Qt__MatchFixedString Qt__MatchFlag = 8
const Qt__MatchCaseSensitive Qt__MatchFlag = 16
const Qt__MatchWrap Qt__MatchFlag = 32
const Qt__MatchRecursive Qt__MatchFlag = 64

type Qt__WindowModality = int

const Qt__NonModal Qt__WindowModality = 0
const Qt__WindowModal Qt__WindowModality = 1
const Qt__ApplicationModal Qt__WindowModality = 2

type Qt__TextInteractionFlag = int

const Qt__NoTextInteraction Qt__TextInteractionFlag = 0
const Qt__TextSelectableByMouse Qt__TextInteractionFlag = 1
const Qt__TextSelectableByKeyboard Qt__TextInteractionFlag = 2
const Qt__LinksAccessibleByMouse Qt__TextInteractionFlag = 4
const Qt__LinksAccessibleByKeyboard Qt__TextInteractionFlag = 8
const Qt__TextEditable Qt__TextInteractionFlag = 16
const Qt__TextEditorInteraction Qt__TextInteractionFlag = 19
const Qt__TextBrowserInteraction Qt__TextInteractionFlag = 13

type Qt__EventPriority = int

const Qt__HighEventPriority Qt__EventPriority = 1
const Qt__NormalEventPriority Qt__EventPriority = 0
const Qt__LowEventPriority Qt__EventPriority = -1

type Qt__SizeHint = int

const Qt__MinimumSize Qt__SizeHint = 0
const Qt__PreferredSize Qt__SizeHint = 1
const Qt__MaximumSize Qt__SizeHint = 2
const Qt__MinimumDescent Qt__SizeHint = 3
const Qt__NSizeHints Qt__SizeHint = 4

type Qt__WindowFrameSection = int

const Qt__NoSection Qt__WindowFrameSection = 0
const Qt__LeftSection Qt__WindowFrameSection = 1
const Qt__TopLeftSection Qt__WindowFrameSection = 2
const Qt__TopSection Qt__WindowFrameSection = 3
const Qt__TopRightSection Qt__WindowFrameSection = 4
const Qt__RightSection Qt__WindowFrameSection = 5
const Qt__BottomRightSection Qt__WindowFrameSection = 6
const Qt__BottomSection Qt__WindowFrameSection = 7
const Qt__BottomLeftSection Qt__WindowFrameSection = 8
const Qt__TitleBarArea Qt__WindowFrameSection = 9

type Qt__Initialization = int

const Qt__Uninitialized Qt__Initialization = 0

type Qt__CoordinateSystem = int

const Qt__DeviceCoordinates Qt__CoordinateSystem = 0
const Qt__LogicalCoordinates Qt__CoordinateSystem = 1

type Qt__TouchPointState = int

const Qt__TouchPointPressed Qt__TouchPointState = 1
const Qt__TouchPointMoved Qt__TouchPointState = 2
const Qt__TouchPointStationary Qt__TouchPointState = 4
const Qt__TouchPointReleased Qt__TouchPointState = 8

type Qt__GestureState = int

const Qt__NoGesture Qt__GestureState = 0
const Qt__GestureStarted Qt__GestureState = 1
const Qt__GestureUpdated Qt__GestureState = 2
const Qt__GestureFinished Qt__GestureState = 3
const Qt__GestureCanceled Qt__GestureState = 4

type Qt__GestureType = int

const Qt__TapGesture Qt__GestureType = 1
const Qt__TapAndHoldGesture Qt__GestureType = 2
const Qt__PanGesture Qt__GestureType = 3
const Qt__PinchGesture Qt__GestureType = 4
const Qt__SwipeGesture Qt__GestureType = 5
const Qt__CustomGesture Qt__GestureType = 256
const Qt__LastGestureType Qt__GestureType = -1

type Qt__GestureFlag = int

const Qt__DontStartGestureOnChildren Qt__GestureFlag = 1
const Qt__ReceivePartialGestures Qt__GestureFlag = 2
const Qt__IgnoredGesturesPropagateToParent Qt__GestureFlag = 4

type Qt__NativeGestureType = int

const Qt__BeginNativeGesture Qt__NativeGestureType = 0
const Qt__EndNativeGesture Qt__NativeGestureType = 1
const Qt__PanNativeGesture Qt__NativeGestureType = 2
const Qt__ZoomNativeGesture Qt__NativeGestureType = 3
const Qt__SmartZoomNativeGesture Qt__NativeGestureType = 4
const Qt__RotateNativeGesture Qt__NativeGestureType = 5
const Qt__SwipeNativeGesture Qt__NativeGestureType = 6

type Qt__NavigationMode = int

const Qt__NavigationModeNone Qt__NavigationMode = 0
const Qt__NavigationModeKeypadTabOrder Qt__NavigationMode = 1
const Qt__NavigationModeKeypadDirectional Qt__NavigationMode = 2
const Qt__NavigationModeCursorAuto Qt__NavigationMode = 3
const Qt__NavigationModeCursorForceVisible Qt__NavigationMode = 4

type Qt__CursorMoveStyle = int

const Qt__LogicalMoveStyle Qt__CursorMoveStyle = 0
const Qt__VisualMoveStyle Qt__CursorMoveStyle = 1

type Qt__TimerType = int

const Qt__PreciseTimer Qt__TimerType = 0
const Qt__CoarseTimer Qt__TimerType = 1
const Qt__VeryCoarseTimer Qt__TimerType = 2

type Qt__ScrollPhase = int

const Qt__NoScrollPhase Qt__ScrollPhase = 0
const Qt__ScrollBegin Qt__ScrollPhase = 1
const Qt__ScrollUpdate Qt__ScrollPhase = 2
const Qt__ScrollEnd Qt__ScrollPhase = 3

type Qt__MouseEventSource = int

const Qt__MouseEventNotSynthesized Qt__MouseEventSource = 0
const Qt__MouseEventSynthesizedBySystem Qt__MouseEventSource = 1
const Qt__MouseEventSynthesizedByQt Qt__MouseEventSource = 2
const Qt__MouseEventSynthesizedByApplication Qt__MouseEventSource = 3

type Qt__MouseEventFlag = int

const Qt__MouseEventCreatedDoubleClick Qt__MouseEventFlag = 1
const Qt__MouseEventFlagMask Qt__MouseEventFlag = 255

type Qt__ChecksumType = int

const Qt__ChecksumIso3309 Qt__ChecksumType = 0
const Qt__ChecksumItuV41 Qt__ChecksumType = 1

type Qt__errc = int

const Qt__address_family_not_supported Qt__errc = 97
const Qt__address_in_use Qt__errc = 98
const Qt__address_not_available Qt__errc = 99
const Qt__already_connected Qt__errc = 106
const Qt__argument_list_too_long Qt__errc = 7
const Qt__argument_out_of_domain Qt__errc = 33
const Qt__bad_address Qt__errc = 14
const Qt__bad_file_descriptor Qt__errc = 9
const Qt__bad_message Qt__errc = 74
const Qt__broken_pipe Qt__errc = 32
const Qt__connection_aborted Qt__errc = 103
const Qt__connection_already_in_progress Qt__errc = 114
const Qt__connection_refused Qt__errc = 111
const Qt__connection_reset Qt__errc = 104
const Qt__cross_device_link Qt__errc = 18
const Qt__destination_address_required Qt__errc = 89
const Qt__device_or_resource_busy Qt__errc = 16
const Qt__directory_not_empty Qt__errc = 39
const Qt__executable_format_error Qt__errc = 8
const Qt__file_exists Qt__errc = 17
const Qt__file_too_large Qt__errc = 27
const Qt__filename_too_long Qt__errc = 36
const Qt__function_not_supported Qt__errc = 38
const Qt__host_unreachable Qt__errc = 113
const Qt__identifier_removed Qt__errc = 43
const Qt__illegal_byte_sequence Qt__errc = 84
const Qt__inappropriate_io_control_operation Qt__errc = 25
const Qt__interrupted Qt__errc = 4
const Qt__invalid_argument Qt__errc = 22
const Qt__invalid_seek Qt__errc = 29
const Qt__io_error Qt__errc = 5
const Qt__is_a_directory Qt__errc = 21
const Qt__message_size Qt__errc = 90
const Qt__network_down Qt__errc = 100
const Qt__network_reset Qt__errc = 102
const Qt__network_unreachable Qt__errc = 101
const Qt__no_buffer_space Qt__errc = 105
const Qt__no_child_process Qt__errc = 10
const Qt__no_link Qt__errc = 67
const Qt__no_lock_available Qt__errc = 37
const Qt__no_message_available Qt__errc = 61
const Qt__no_message Qt__errc = 42
const Qt__no_protocol_option Qt__errc = 92
const Qt__no_space_on_device Qt__errc = 28
const Qt__no_stream_resources Qt__errc = 63
const Qt__no_such_device_or_address Qt__errc = 6
const Qt__no_such_device Qt__errc = 19
const Qt__no_such_file_or_directory Qt__errc = 2
const Qt__no_such_process Qt__errc = 3
const Qt__not_a_directory Qt__errc = 20
const Qt__not_a_socket Qt__errc = 88
const Qt__not_a_stream Qt__errc = 60
const Qt__not_connected Qt__errc = 107
const Qt__not_enough_memory Qt__errc = 12
const Qt__not_supported Qt__errc = 95
const Qt__operation_canceled Qt__errc = 125
const Qt__operation_in_progress Qt__errc = 115
const Qt__operation_not_permitted Qt__errc = 1
const Qt__operation_not_supported Qt__errc = 95
const Qt__operation_would_block Qt__errc = 11
const Qt__owner_dead Qt__errc = 130
const Qt__permission_denied Qt__errc = 13
const Qt__protocol_error Qt__errc = 71
const Qt__protocol_not_supported Qt__errc = 93
const Qt__read_only_file_system Qt__errc = 30
const Qt__resource_deadlock_would_occur Qt__errc = 35
const Qt__resource_unavailable_try_again Qt__errc = 11
const Qt__result_out_of_range Qt__errc = 34
const Qt__state_not_recoverable Qt__errc = 131
const Qt__stream_timeout Qt__errc = 62
const Qt__text_file_busy Qt__errc = 26
const Qt__timed_out Qt__errc = 110
const Qt__too_many_files_open_in_system Qt__errc = 23
const Qt__too_many_files_open Qt__errc = 24
const Qt__too_many_links Qt__errc = 31
const Qt__too_many_symbolic_link_levels Qt__errc = 40
const Qt__value_too_large Qt__errc = 75
const Qt__wrong_protocol_type Qt__errc = 91

type Qt___Ios_Fmtflags = int

const Qt___S_boolalpha Qt___Ios_Fmtflags = 1
const Qt___S_dec Qt___Ios_Fmtflags = 2
const Qt___S_fixed Qt___Ios_Fmtflags = 4
const Qt___S_hex Qt___Ios_Fmtflags = 8
const Qt___S_internal Qt___Ios_Fmtflags = 16
const Qt___S_left Qt___Ios_Fmtflags = 32
const Qt___S_oct Qt___Ios_Fmtflags = 64
const Qt___S_right Qt___Ios_Fmtflags = 128
const Qt___S_scientific Qt___Ios_Fmtflags = 256
const Qt___S_showbase Qt___Ios_Fmtflags = 512
const Qt___S_showpoint Qt___Ios_Fmtflags = 1024
const Qt___S_showpos Qt___Ios_Fmtflags = 2048
const Qt___S_skipws Qt___Ios_Fmtflags = 4096
const Qt___S_unitbuf Qt___Ios_Fmtflags = 8192
const Qt___S_uppercase Qt___Ios_Fmtflags = 16384
const Qt___S_adjustfield Qt___Ios_Fmtflags = 176
const Qt___S_basefield Qt___Ios_Fmtflags = 74
const Qt___S_floatfield Qt___Ios_Fmtflags = 260
const Qt___S_ios_fmtflags_end Qt___Ios_Fmtflags = 65536
const Qt___S_ios_fmtflags_max Qt___Ios_Fmtflags = 2147483647
const Qt___S_ios_fmtflags_min Qt___Ios_Fmtflags = -2147483648

type Qt___Ios_Openmode = int

const Qt___S_app Qt___Ios_Openmode = 1
const Qt___S_ate Qt___Ios_Openmode = 2
const Qt___S_bin Qt___Ios_Openmode = 4
const Qt___S_in Qt___Ios_Openmode = 8
const Qt___S_out Qt___Ios_Openmode = 16
const Qt___S_trunc Qt___Ios_Openmode = 32
const Qt___S_ios_openmode_end Qt___Ios_Openmode = 65536
const Qt___S_ios_openmode_max Qt___Ios_Openmode = 2147483647
const Qt___S_ios_openmode_min Qt___Ios_Openmode = -2147483648

type Qt___Ios_Iostate = int

const Qt___S_goodbit Qt___Ios_Iostate = 0
const Qt___S_badbit Qt___Ios_Iostate = 1
const Qt___S_eofbit Qt___Ios_Iostate = 2
const Qt___S_failbit Qt___Ios_Iostate = 4
const Qt___S_ios_iostate_end Qt___Ios_Iostate = 65536
const Qt___S_ios_iostate_max Qt___Ios_Iostate = 2147483647
const Qt___S_ios_iostate_min Qt___Ios_Iostate = -2147483648

type Qt___Ios_Seekdir = int

const Qt___S_beg Qt___Ios_Seekdir = 0
const Qt___S_cur Qt___Ios_Seekdir = 1
const Qt___S_end Qt___Ios_Seekdir = 2
const Qt___S_ios_seekdir_end Qt___Ios_Seekdir = 65536

type Qt__io_errc = int

const Qt__stream Qt__io_errc = 1

type Qt___Rb_tree_color = int

const Qt___S_red Qt___Rb_tree_color = 0
const Qt___S_black Qt___Rb_tree_color = 1

type Qt__IteratorCapability = int

const Qt__ForwardCapability Qt__IteratorCapability = 1
const Qt__BiDirectionalCapability Qt__IteratorCapability = 2
const Qt__RandomAccessCapability Qt__IteratorCapability = 4

type Qt___Manager_operation = int

const Qt____get_type_info Qt___Manager_operation = 0
const Qt____get_functor_ptr Qt___Manager_operation = 1
const Qt____clone_functor Qt___Manager_operation = 2
const Qt____destroy_functor Qt___Manager_operation = 3

type Qt___Lock_policy = int

const Qt___S_single Qt___Lock_policy = 0
const Qt___S_mutex Qt___Lock_policy = 1
const Qt___S_atomic Qt___Lock_policy = 2

type Qt__pointer_safety = int

const Qt__relaxed Qt__pointer_safety = 0
const Qt__preferred Qt__pointer_safety = 1
const Qt__strict Qt__pointer_safety = 2

type Qt__future_errc = int

const Qt__future_already_retrieved Qt__future_errc = 1
const Qt__promise_already_satisfied Qt__future_errc = 2
const Qt__no_state Qt__future_errc = 3
const Qt__broken_promise Qt__future_errc = 4

type Qt__future_status = int

const Qt__ready Qt__future_status = 0
const Qt__timeout Qt__future_status = 1
const Qt__deferred Qt__future_status = 2

type Qt__DrawingHint = int

const Qt__OpaqueTopLeft Qt__DrawingHint = 1
const Qt__OpaqueTop Qt__DrawingHint = 2
const Qt__OpaqueTopRight Qt__DrawingHint = 4
const Qt__OpaqueLeft Qt__DrawingHint = 8
const Qt__OpaqueCenter Qt__DrawingHint = 16
const Qt__OpaqueRight Qt__DrawingHint = 32
const Qt__OpaqueBottomLeft Qt__DrawingHint = 64
const Qt__OpaqueBottom Qt__DrawingHint = 128
const Qt__OpaqueBottomRight Qt__DrawingHint = 256
const Qt__OpaqueCorners Qt__DrawingHint = 325
const Qt__OpaqueEdges Qt__DrawingHint = 170
const Qt__OpaqueFrame Qt__DrawingHint = 495
const Qt__OpaqueAll Qt__DrawingHint = 511

//  body block end
