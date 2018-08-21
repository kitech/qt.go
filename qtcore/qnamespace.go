package qtcore

import "fmt"

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

func float_round_styleItemName(val int) string {
	switch val {
	case Qt__round_indeterminate: // -1
		return "round_indeterminate"
	case Qt__round_toward_zero: // 0
		return "round_toward_zero"
	case Qt__round_to_nearest: // 1
		return "round_to_nearest"
	case Qt__round_toward_infinity: // 2
		return "round_toward_infinity"
	case Qt__round_toward_neg_infinity: // 3
		return "round_toward_neg_infinity"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__float_denorm_style = int // stdglobal
//
const Qt__denorm_indeterminate Qt__float_denorm_style = -1

//
const Qt__denorm_absent Qt__float_denorm_style = 0

//
const Qt__denorm_present Qt__float_denorm_style = 1

func float_denorm_styleItemName(val int) string {
	switch val {
	case Qt__denorm_indeterminate: // -1
		return "denorm_indeterminate"
	case Qt__denorm_absent: // 0
		return "denorm_absent"
	case Qt__denorm_present: // 1
		return "denorm_present"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func QtMsgTypeItemName(val int) string {
	switch val {
	case Qt__QtDebugMsg: // 0
		return "QtDebugMsg"
	case Qt__QtWarningMsg: // 1
		return "QtWarningMsg"
	case Qt__QtCriticalMsg: // 2
		return "QtCriticalMsg,QtSystemMsg"
	case Qt__QtFatalMsg: // 3
		return "QtFatalMsg"
	case Qt__QtInfoMsg: // 4
		return "QtInfoMsg"
		// case Qt__QtSystemMsg: // 2
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func memory_orderItemName(val int) string {
	switch val {
	case Qt__memory_order_relaxed: // 0
		return "memory_order_relaxed"
	case Qt__memory_order_consume: // 1
		return "memory_order_consume"
	case Qt__memory_order_acquire: // 2
		return "memory_order_acquire"
	case Qt__memory_order_release: // 3
		return "memory_order_release"
	case Qt__memory_order_acq_rel: // 4
		return "memory_order_acq_rel"
	case Qt__memory_order_seq_cst: // 5
		return "memory_order_seq_cst"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func __memory_order_modifierItemName(val int) string {
	switch val {
	case Qt____memory_order_mask: // 65535
		return "__memory_order_mask"
	case Qt____memory_order_modifier_mask: // -65536
		return "__memory_order_modifier_mask"
	case Qt____memory_order_hle_acquire: // 65536
		return "__memory_order_hle_acquire"
	case Qt____memory_order_hle_release: // 131072
		return "__memory_order_hle_release"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func GlobalColorItemName(val int) string {
	switch val {
	case Qt__color0: // 0
		return "color0"
	case Qt__color1: // 1
		return "color1"
	case Qt__black: // 2
		return "black"
	case Qt__white: // 3
		return "white"
	case Qt__darkGray: // 4
		return "darkGray"
	case Qt__gray: // 5
		return "gray"
	case Qt__lightGray: // 6
		return "lightGray"
	case Qt__red: // 7
		return "red"
	case Qt__green: // 8
		return "green"
	case Qt__blue: // 9
		return "blue"
	case Qt__cyan: // 10
		return "cyan"
	case Qt__magenta: // 11
		return "magenta"
	case Qt__yellow: // 12
		return "yellow"
	case Qt__darkRed: // 13
		return "darkRed"
	case Qt__darkGreen: // 14
		return "darkGreen"
	case Qt__darkBlue: // 15
		return "darkBlue"
	case Qt__darkCyan: // 16
		return "darkCyan"
	case Qt__darkMagenta: // 17
		return "darkMagenta"
	case Qt__darkYellow: // 18
		return "darkYellow"
	case Qt__transparent: // 19
		return "transparent"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func KeyboardModifierItemName(val int) string {
	switch val {
	case Qt__NoModifier: // 0
		return "NoModifier"
	case Qt__ShiftModifier: // 33554432
		return "ShiftModifier"
	case Qt__ControlModifier: // 67108864
		return "ControlModifier"
	case Qt__AltModifier: // 134217728
		return "AltModifier"
	case Qt__MetaModifier: // 268435456
		return "MetaModifier"
	case Qt__KeypadModifier: // 536870912
		return "KeypadModifier"
	case Qt__GroupSwitchModifier: // 1073741824
		return "GroupSwitchModifier"
	case Qt__KeyboardModifierMask: // -33554432
		return "KeyboardModifierMask"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ModifierItemName(val int) string {
	switch val {
	case Qt__META: // 268435456
		return "META"
	case Qt__SHIFT: // 33554432
		return "SHIFT"
	case Qt__CTRL: // 67108864
		return "CTRL"
	case Qt__ALT: // 134217728
		return "ALT"
	case Qt__MODIFIER_MASK: // -33554432
		return "MODIFIER_MASK"
	case Qt__UNICODE_ACCEL: // 0
		return "UNICODE_ACCEL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func MouseButtonItemName(val int) string {
	switch val {
	case Qt__NoButton: // 0
		return "NoButton"
	case Qt__LeftButton: // 1
		return "LeftButton"
	case Qt__RightButton: // 2
		return "RightButton"
	case Qt__MidButton: // 4
		return "MidButton,MiddleButton"
		// case Qt__MiddleButton: // 4
		// return ""
	case Qt__BackButton: // 8
		return "BackButton,XButton1,ExtraButton1"
		// case Qt__XButton1: // 8
		// return ""
		// case Qt__ExtraButton1: // 8
		// return ""
	case Qt__ForwardButton: // 16
		return "ForwardButton,XButton2,ExtraButton2"
		// case Qt__XButton2: // 16
		// return ""
		// case Qt__ExtraButton2: // 16
		// return ""
	case Qt__TaskButton: // 32
		return "TaskButton,ExtraButton3"
		// case Qt__ExtraButton3: // 32
		// return ""
	case Qt__ExtraButton4: // 64
		return "ExtraButton4"
	case Qt__ExtraButton5: // 128
		return "ExtraButton5"
	case Qt__ExtraButton6: // 256
		return "ExtraButton6"
	case Qt__ExtraButton7: // 512
		return "ExtraButton7"
	case Qt__ExtraButton8: // 1024
		return "ExtraButton8"
	case Qt__ExtraButton9: // 2048
		return "ExtraButton9"
	case Qt__ExtraButton10: // 4096
		return "ExtraButton10"
	case Qt__ExtraButton11: // 8192
		return "ExtraButton11"
	case Qt__ExtraButton12: // 16384
		return "ExtraButton12"
	case Qt__ExtraButton13: // 32768
		return "ExtraButton13"
	case Qt__ExtraButton14: // 65536
		return "ExtraButton14"
	case Qt__ExtraButton15: // 131072
		return "ExtraButton15"
	case Qt__ExtraButton16: // 262144
		return "ExtraButton16"
	case Qt__ExtraButton17: // 524288
		return "ExtraButton17"
	case Qt__ExtraButton18: // 1048576
		return "ExtraButton18"
	case Qt__ExtraButton19: // 2097152
		return "ExtraButton19"
	case Qt__ExtraButton20: // 4194304
		return "ExtraButton20"
	case Qt__ExtraButton21: // 8388608
		return "ExtraButton21"
	case Qt__ExtraButton22: // 16777216
		return "ExtraButton22"
	case Qt__ExtraButton23: // 33554432
		return "ExtraButton23"
	case Qt__ExtraButton24: // 67108864
		return "ExtraButton24,MaxMouseButton"
	case Qt__AllButtons: // 134217727
		return "AllButtons"
		// case Qt__MaxMouseButton: // 67108864
		// return ""
	case Qt__MouseButtonMask: // -1
		return "MouseButtonMask"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__Orientation = int // core
//
const Qt__Horizontal Qt__Orientation = 1

//
const Qt__Vertical Qt__Orientation = 2

func OrientationItemName(val int) string {
	switch val {
	case Qt__Horizontal: // 1
		return "Horizontal"
	case Qt__Vertical: // 2
		return "Vertical"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func FocusPolicyItemName(val int) string {
	switch val {
	case Qt__NoFocus: // 0
		return "NoFocus"
	case Qt__TabFocus: // 1
		return "TabFocus"
	case Qt__ClickFocus: // 2
		return "ClickFocus"
	case Qt__StrongFocus: // 11
		return "StrongFocus"
	case Qt__WheelFocus: // 15
		return "WheelFocus"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TabFocusBehaviorItemName(val int) string {
	switch val {
	case Qt__NoTabFocus: // 0
		return "NoTabFocus"
	case Qt__TabFocusTextControls: // 1
		return "TabFocusTextControls"
	case Qt__TabFocusListControls: // 2
		return "TabFocusListControls"
	case Qt__TabFocusAllControls: // 255
		return "TabFocusAllControls"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum describes how the items in a widget are sorted.


*/
type Qt__SortOrder = int // core
//
const Qt__AscendingOrder Qt__SortOrder = 0

//
const Qt__DescendingOrder Qt__SortOrder = 1

func SortOrderItemName(val int) string {
	switch val {
	case Qt__AscendingOrder: // 0
		return "AscendingOrder"
	case Qt__DescendingOrder: // 1
		return "DescendingOrder"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TileRuleItemName(val int) string {
	switch val {
	case Qt__StretchTile: // 0
		return "StretchTile"
	case Qt__RepeatTile: // 1
		return "RepeatTile"
	case Qt__RoundTile: // 2
		return "RoundTile"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func AlignmentFlagItemName(val int) string {
	switch val {
	case Qt__AlignLeft: // 1
		return "AlignLeft,AlignLeading"
		// case Qt__AlignLeading: // 1
		// return ""
	case Qt__AlignRight: // 2
		return "AlignRight,AlignTrailing"
		// case Qt__AlignTrailing: // 2
		// return ""
	case Qt__AlignHCenter: // 4
		return "AlignHCenter"
	case Qt__AlignJustify: // 8
		return "AlignJustify"
	case Qt__AlignAbsolute: // 16
		return "AlignAbsolute"
	case Qt__AlignHorizontal_Mask: // 31
		return "AlignHorizontal_Mask"
	case Qt__AlignTop: // 32
		return "AlignTop"
	case Qt__AlignBottom: // 64
		return "AlignBottom"
	case Qt__AlignVCenter: // 128
		return "AlignVCenter"
	case Qt__AlignBaseline: // 256
		return "AlignBaseline"
	case Qt__AlignVertical_Mask: // 480
		return "AlignVertical_Mask"
	case Qt__AlignCenter: // 132
		return "AlignCenter"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TextFlagItemName(val int) string {
	switch val {
	case Qt__TextSingleLine: // 256
		return "TextSingleLine"
	case Qt__TextDontClip: // 512
		return "TextDontClip"
	case Qt__TextExpandTabs: // 1024
		return "TextExpandTabs"
	case Qt__TextShowMnemonic: // 2048
		return "TextShowMnemonic"
	case Qt__TextWordWrap: // 4096
		return "TextWordWrap"
	case Qt__TextWrapAnywhere: // 8192
		return "TextWrapAnywhere"
	case Qt__TextDontPrint: // 16384
		return "TextDontPrint"
	case Qt__TextIncludeTrailingSpaces: // 134217728
		return "TextIncludeTrailingSpaces"
	case Qt__TextHideMnemonic: // 32768
		return "TextHideMnemonic"
	case Qt__TextJustificationForced: // 65536
		return "TextJustificationForced"
	case Qt__TextForceLeftToRight: // 131072
		return "TextForceLeftToRight"
	case Qt__TextForceRightToLeft: // 262144
		return "TextForceRightToLeft"
	case Qt__TextLongestVariant: // 524288
		return "TextLongestVariant"
	case Qt__TextBypassShaping: // 1048576
		return "TextBypassShaping"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TextElideModeItemName(val int) string {
	switch val {
	case Qt__ElideLeft: // 0
		return "ElideLeft"
	case Qt__ElideRight: // 1
		return "ElideRight"
	case Qt__ElideMiddle: // 2
		return "ElideMiddle"
	case Qt__ElideNone: // 3
		return "ElideNone"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func WhiteSpaceModeItemName(val int) string {
	switch val {
	case Qt__WhiteSpaceNormal: // 0
		return "WhiteSpaceNormal"
	case Qt__WhiteSpacePre: // 1
		return "WhiteSpacePre"
	case Qt__WhiteSpaceNoWrap: // 2
		return "WhiteSpaceNoWrap"
	case Qt__WhiteSpaceModeUndefined: // -1
		return "WhiteSpaceModeUndefined"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum contains the types of accuracy that can be used by the QTextDocument class when testing for mouse clicks on text documents.


*/
type Qt__HitTestAccuracy = int // core
// The point at which input occurred must coincide exactly with input-sensitive parts of the document.
const Qt__ExactHit Qt__HitTestAccuracy = 0

// The point at which input occurred can lie close to input-sensitive parts of the document.
const Qt__FuzzyHit Qt__HitTestAccuracy = 1

func HitTestAccuracyItemName(val int) string {
	switch val {
	case Qt__ExactHit: // 0
		return "ExactHit"
	case Qt__FuzzyHit: // 1
		return "FuzzyHit"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func WindowTypeItemName(val int) string {
	switch val {
	case Qt__Widget: // 0
		return "Widget"
	case Qt__Window: // 1
		return "Window"
	case Qt__Dialog: // 3
		return "Dialog"
	case Qt__Sheet: // 5
		return "Sheet"
	case Qt__Drawer: // 7
		return "Drawer"
	case Qt__Popup: // 9
		return "Popup"
	case Qt__Tool: // 11
		return "Tool"
	case Qt__ToolTip: // 13
		return "ToolTip"
	case Qt__SplashScreen: // 15
		return "SplashScreen"
	case Qt__Desktop: // 17
		return "Desktop"
	case Qt__SubWindow: // 18
		return "SubWindow"
	case Qt__ForeignWindow: // 33
		return "ForeignWindow"
	case Qt__CoverWindow: // 65
		return "CoverWindow"
	case Qt__WindowType_Mask: // 255
		return "WindowType_Mask"
	case Qt__MSWindowsFixedSizeDialogHint: // 256
		return "MSWindowsFixedSizeDialogHint"
	case Qt__MSWindowsOwnDC: // 512
		return "MSWindowsOwnDC"
	case Qt__BypassWindowManagerHint: // 1024
		return "BypassWindowManagerHint,X11BypassWindowManagerHint"
		// case Qt__X11BypassWindowManagerHint: // 1024
		// return ""
	case Qt__FramelessWindowHint: // 2048
		return "FramelessWindowHint"
	case Qt__WindowTitleHint: // 4096
		return "WindowTitleHint"
	case Qt__WindowSystemMenuHint: // 8192
		return "WindowSystemMenuHint"
	case Qt__WindowMinimizeButtonHint: // 16384
		return "WindowMinimizeButtonHint"
	case Qt__WindowMaximizeButtonHint: // 32768
		return "WindowMaximizeButtonHint"
	case Qt__WindowMinMaxButtonsHint: // 49152
		return "WindowMinMaxButtonsHint"
	case Qt__WindowContextHelpButtonHint: // 65536
		return "WindowContextHelpButtonHint"
	case Qt__WindowShadeButtonHint: // 131072
		return "WindowShadeButtonHint"
	case Qt__WindowStaysOnTopHint: // 262144
		return "WindowStaysOnTopHint"
	case Qt__WindowTransparentForInput: // 524288
		return "WindowTransparentForInput"
	case Qt__WindowOverridesSystemGestures: // 1048576
		return "WindowOverridesSystemGestures"
	case Qt__WindowDoesNotAcceptFocus: // 2097152
		return "WindowDoesNotAcceptFocus"
	case Qt__MaximizeUsingFullscreenGeometryHint: // 4194304
		return "MaximizeUsingFullscreenGeometryHint"
	case Qt__CustomizeWindowHint: // 33554432
		return "CustomizeWindowHint"
	case Qt__WindowStaysOnBottomHint: // 67108864
		return "WindowStaysOnBottomHint"
	case Qt__WindowCloseButtonHint: // 134217728
		return "WindowCloseButtonHint"
	case Qt__MacWindowToolBarButtonHint: // 268435456
		return "MacWindowToolBarButtonHint"
	case Qt__BypassGraphicsProxyWidget: // 536870912
		return "BypassGraphicsProxyWidget"
	case Qt__NoDropShadowWindowHint: // 1073741824
		return "NoDropShadowWindowHint"
	case Qt__WindowFullscreenButtonHint: // -2147483648
		return "WindowFullscreenButtonHint"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func WindowStateItemName(val int) string {
	switch val {
	case Qt__WindowNoState: // 0
		return "WindowNoState"
	case Qt__WindowMinimized: // 1
		return "WindowMinimized"
	case Qt__WindowMaximized: // 2
		return "WindowMaximized"
	case Qt__WindowFullScreen: // 4
		return "WindowFullScreen"
	case Qt__WindowActive: // 8
		return "WindowActive"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ApplicationStateItemName(val int) string {
	switch val {
	case Qt__ApplicationSuspended: // 0
		return "ApplicationSuspended"
	case Qt__ApplicationHidden: // 1
		return "ApplicationHidden"
	case Qt__ApplicationInactive: // 2
		return "ApplicationInactive"
	case Qt__ApplicationActive: // 4
		return "ApplicationActive"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ScreenOrientationItemName(val int) string {
	switch val {
	case Qt__PrimaryOrientation: // 0
		return "PrimaryOrientation"
	case Qt__PortraitOrientation: // 1
		return "PortraitOrientation"
	case Qt__LandscapeOrientation: // 2
		return "LandscapeOrientation"
	case Qt__InvertedPortraitOrientation: // 4
		return "InvertedPortraitOrientation"
	case Qt__InvertedLandscapeOrientation: // 8
		return "InvertedLandscapeOrientation"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func WidgetAttributeItemName(val int) string {
	switch val {
	case Qt__WA_Disabled: // 0
		return "WA_Disabled"
	case Qt__WA_UnderMouse: // 1
		return "WA_UnderMouse"
	case Qt__WA_MouseTracking: // 2
		return "WA_MouseTracking"
	case Qt__WA_ContentsPropagated: // 3
		return "WA_ContentsPropagated"
	case Qt__WA_OpaquePaintEvent: // 4
		return "WA_OpaquePaintEvent,WA_NoBackground"
		// case Qt__WA_NoBackground: // 4
		// return ""
	case Qt__WA_StaticContents: // 5
		return "WA_StaticContents"
	case Qt__WA_LaidOut: // 7
		return "WA_LaidOut"
	case Qt__WA_PaintOnScreen: // 8
		return "WA_PaintOnScreen"
	case Qt__WA_NoSystemBackground: // 9
		return "WA_NoSystemBackground"
	case Qt__WA_UpdatesDisabled: // 10
		return "WA_UpdatesDisabled"
	case Qt__WA_Mapped: // 11
		return "WA_Mapped"
	case Qt__WA_MacNoClickThrough: // 12
		return "WA_MacNoClickThrough"
	case Qt__WA_InputMethodEnabled: // 14
		return "WA_InputMethodEnabled"
	case Qt__WA_WState_Visible: // 15
		return "WA_WState_Visible"
	case Qt__WA_WState_Hidden: // 16
		return "WA_WState_Hidden"
	case Qt__WA_ForceDisabled: // 32
		return "WA_ForceDisabled"
	case Qt__WA_KeyCompression: // 33
		return "WA_KeyCompression"
	case Qt__WA_PendingMoveEvent: // 34
		return "WA_PendingMoveEvent"
	case Qt__WA_PendingResizeEvent: // 35
		return "WA_PendingResizeEvent"
	case Qt__WA_SetPalette: // 36
		return "WA_SetPalette"
	case Qt__WA_SetFont: // 37
		return "WA_SetFont"
	case Qt__WA_SetCursor: // 38
		return "WA_SetCursor"
	case Qt__WA_NoChildEventsFromChildren: // 39
		return "WA_NoChildEventsFromChildren"
	case Qt__WA_WindowModified: // 41
		return "WA_WindowModified"
	case Qt__WA_Resized: // 42
		return "WA_Resized"
	case Qt__WA_Moved: // 43
		return "WA_Moved"
	case Qt__WA_PendingUpdate: // 44
		return "WA_PendingUpdate"
	case Qt__WA_InvalidSize: // 45
		return "WA_InvalidSize"
	case Qt__WA_MacBrushedMetal: // 46
		return "WA_MacBrushedMetal,WA_MacMetalStyle"
		// case Qt__WA_MacMetalStyle: // 46
		// return ""
	case Qt__WA_CustomWhatsThis: // 47
		return "WA_CustomWhatsThis"
	case Qt__WA_LayoutOnEntireRect: // 48
		return "WA_LayoutOnEntireRect"
	case Qt__WA_OutsideWSRange: // 49
		return "WA_OutsideWSRange"
	case Qt__WA_GrabbedShortcut: // 50
		return "WA_GrabbedShortcut"
	case Qt__WA_TransparentForMouseEvents: // 51
		return "WA_TransparentForMouseEvents"
	case Qt__WA_PaintUnclipped: // 52
		return "WA_PaintUnclipped"
	case Qt__WA_SetWindowIcon: // 53
		return "WA_SetWindowIcon"
	case Qt__WA_NoMouseReplay: // 54
		return "WA_NoMouseReplay"
	case Qt__WA_DeleteOnClose: // 55
		return "WA_DeleteOnClose"
	case Qt__WA_RightToLeft: // 56
		return "WA_RightToLeft"
	case Qt__WA_SetLayoutDirection: // 57
		return "WA_SetLayoutDirection"
	case Qt__WA_NoChildEventsForParent: // 58
		return "WA_NoChildEventsForParent"
	case Qt__WA_ForceUpdatesDisabled: // 59
		return "WA_ForceUpdatesDisabled"
	case Qt__WA_WState_Created: // 60
		return "WA_WState_Created"
	case Qt__WA_WState_CompressKeys: // 61
		return "WA_WState_CompressKeys"
	case Qt__WA_WState_InPaintEvent: // 62
		return "WA_WState_InPaintEvent"
	case Qt__WA_WState_Reparented: // 63
		return "WA_WState_Reparented"
	case Qt__WA_WState_ConfigPending: // 64
		return "WA_WState_ConfigPending"
	case Qt__WA_WState_Polished: // 66
		return "WA_WState_Polished"
	case Qt__WA_WState_DND: // 67
		return "WA_WState_DND"
	case Qt__WA_WState_OwnSizePolicy: // 68
		return "WA_WState_OwnSizePolicy"
	case Qt__WA_WState_ExplicitShowHide: // 69
		return "WA_WState_ExplicitShowHide"
	case Qt__WA_ShowModal: // 70
		return "WA_ShowModal"
	case Qt__WA_MouseNoMask: // 71
		return "WA_MouseNoMask"
	case Qt__WA_GroupLeader: // 72
		return "WA_GroupLeader"
	case Qt__WA_NoMousePropagation: // 73
		return "WA_NoMousePropagation"
	case Qt__WA_Hover: // 74
		return "WA_Hover"
	case Qt__WA_InputMethodTransparent: // 75
		return "WA_InputMethodTransparent"
	case Qt__WA_QuitOnClose: // 76
		return "WA_QuitOnClose"
	case Qt__WA_KeyboardFocusChange: // 77
		return "WA_KeyboardFocusChange"
	case Qt__WA_AcceptDrops: // 78
		return "WA_AcceptDrops"
	case Qt__WA_DropSiteRegistered: // 79
		return "WA_DropSiteRegistered,WA_ForceAcceptDrops"
		// case Qt__WA_ForceAcceptDrops: // 79
		// return ""
	case Qt__WA_WindowPropagation: // 80
		return "WA_WindowPropagation"
	case Qt__WA_NoX11EventCompression: // 81
		return "WA_NoX11EventCompression"
	case Qt__WA_TintedBackground: // 82
		return "WA_TintedBackground"
	case Qt__WA_X11OpenGLOverlay: // 83
		return "WA_X11OpenGLOverlay"
	case Qt__WA_AlwaysShowToolTips: // 84
		return "WA_AlwaysShowToolTips"
	case Qt__WA_MacOpaqueSizeGrip: // 85
		return "WA_MacOpaqueSizeGrip"
	case Qt__WA_SetStyle: // 86
		return "WA_SetStyle"
	case Qt__WA_SetLocale: // 87
		return "WA_SetLocale"
	case Qt__WA_MacShowFocusRect: // 88
		return "WA_MacShowFocusRect"
	case Qt__WA_MacNormalSize: // 89
		return "WA_MacNormalSize"
	case Qt__WA_MacSmallSize: // 90
		return "WA_MacSmallSize"
	case Qt__WA_MacMiniSize: // 91
		return "WA_MacMiniSize"
	case Qt__WA_LayoutUsesWidgetRect: // 92
		return "WA_LayoutUsesWidgetRect"
	case Qt__WA_StyledBackground: // 93
		return "WA_StyledBackground"
	case Qt__WA_MSWindowsUseDirect3D: // 94
		return "WA_MSWindowsUseDirect3D"
	case Qt__WA_CanHostQMdiSubWindowTitleBar: // 95
		return "WA_CanHostQMdiSubWindowTitleBar"
	case Qt__WA_MacAlwaysShowToolWindow: // 96
		return "WA_MacAlwaysShowToolWindow"
	case Qt__WA_StyleSheet: // 97
		return "WA_StyleSheet"
	case Qt__WA_ShowWithoutActivating: // 98
		return "WA_ShowWithoutActivating"
	case Qt__WA_X11BypassTransientForHint: // 99
		return "WA_X11BypassTransientForHint"
	case Qt__WA_NativeWindow: // 100
		return "WA_NativeWindow"
	case Qt__WA_DontCreateNativeAncestors: // 101
		return "WA_DontCreateNativeAncestors"
	case Qt__WA_MacVariableSize: // 102
		return "WA_MacVariableSize"
	case Qt__WA_DontShowOnScreen: // 103
		return "WA_DontShowOnScreen"
	case Qt__WA_X11NetWmWindowTypeDesktop: // 104
		return "WA_X11NetWmWindowTypeDesktop"
	case Qt__WA_X11NetWmWindowTypeDock: // 105
		return "WA_X11NetWmWindowTypeDock"
	case Qt__WA_X11NetWmWindowTypeToolBar: // 106
		return "WA_X11NetWmWindowTypeToolBar"
	case Qt__WA_X11NetWmWindowTypeMenu: // 107
		return "WA_X11NetWmWindowTypeMenu"
	case Qt__WA_X11NetWmWindowTypeUtility: // 108
		return "WA_X11NetWmWindowTypeUtility"
	case Qt__WA_X11NetWmWindowTypeSplash: // 109
		return "WA_X11NetWmWindowTypeSplash"
	case Qt__WA_X11NetWmWindowTypeDialog: // 110
		return "WA_X11NetWmWindowTypeDialog"
	case Qt__WA_X11NetWmWindowTypeDropDownMenu: // 111
		return "WA_X11NetWmWindowTypeDropDownMenu"
	case Qt__WA_X11NetWmWindowTypePopupMenu: // 112
		return "WA_X11NetWmWindowTypePopupMenu"
	case Qt__WA_X11NetWmWindowTypeToolTip: // 113
		return "WA_X11NetWmWindowTypeToolTip"
	case Qt__WA_X11NetWmWindowTypeNotification: // 114
		return "WA_X11NetWmWindowTypeNotification"
	case Qt__WA_X11NetWmWindowTypeCombo: // 115
		return "WA_X11NetWmWindowTypeCombo"
	case Qt__WA_X11NetWmWindowTypeDND: // 116
		return "WA_X11NetWmWindowTypeDND"
	case Qt__WA_MacFrameworkScaled: // 117
		return "WA_MacFrameworkScaled"
	case Qt__WA_SetWindowModality: // 118
		return "WA_SetWindowModality"
	case Qt__WA_WState_WindowOpacitySet: // 119
		return "WA_WState_WindowOpacitySet"
	case Qt__WA_TranslucentBackground: // 120
		return "WA_TranslucentBackground"
	case Qt__WA_AcceptTouchEvents: // 121
		return "WA_AcceptTouchEvents"
	case Qt__WA_WState_AcceptedTouchBeginEvent: // 122
		return "WA_WState_AcceptedTouchBeginEvent"
	case Qt__WA_TouchPadAcceptSingleTouchEvents: // 123
		return "WA_TouchPadAcceptSingleTouchEvents"
	case Qt__WA_X11DoNotAcceptFocus: // 126
		return "WA_X11DoNotAcceptFocus"
	case Qt__WA_MacNoShadow: // 127
		return "WA_MacNoShadow"
	case Qt__WA_AlwaysStackOnTop: // 128
		return "WA_AlwaysStackOnTop"
	case Qt__WA_TabletTracking: // 129
		return "WA_TabletTracking"
	case Qt__WA_ContentsMarginsRespectsSafeArea: // 130
		return "WA_ContentsMarginsRespectsSafeArea"
	case Qt__WA_AttributeCount: // 131
		return "WA_AttributeCount"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ApplicationAttributeItemName(val int) string {
	switch val {
	case Qt__AA_ImmediateWidgetCreation: // 0
		return "AA_ImmediateWidgetCreation"
	case Qt__AA_MSWindowsUseDirect3DByDefault: // 1
		return "AA_MSWindowsUseDirect3DByDefault"
	case Qt__AA_DontShowIconsInMenus: // 2
		return "AA_DontShowIconsInMenus"
	case Qt__AA_NativeWindows: // 3
		return "AA_NativeWindows"
	case Qt__AA_DontCreateNativeWidgetSiblings: // 4
		return "AA_DontCreateNativeWidgetSiblings"
	case Qt__AA_PluginApplication: // 5
		return "AA_PluginApplication,AA_MacPluginApplication"
		// case Qt__AA_MacPluginApplication: // 5
		// return ""
	case Qt__AA_DontUseNativeMenuBar: // 6
		return "AA_DontUseNativeMenuBar"
	case Qt__AA_MacDontSwapCtrlAndMeta: // 7
		return "AA_MacDontSwapCtrlAndMeta"
	case Qt__AA_Use96Dpi: // 8
		return "AA_Use96Dpi"
	case Qt__AA_X11InitThreads: // 10
		return "AA_X11InitThreads"
	case Qt__AA_SynthesizeTouchForUnhandledMouseEvents: // 11
		return "AA_SynthesizeTouchForUnhandledMouseEvents"
	case Qt__AA_SynthesizeMouseForUnhandledTouchEvents: // 12
		return "AA_SynthesizeMouseForUnhandledTouchEvents"
	case Qt__AA_UseHighDpiPixmaps: // 13
		return "AA_UseHighDpiPixmaps"
	case Qt__AA_ForceRasterWidgets: // 14
		return "AA_ForceRasterWidgets"
	case Qt__AA_UseDesktopOpenGL: // 15
		return "AA_UseDesktopOpenGL"
	case Qt__AA_UseOpenGLES: // 16
		return "AA_UseOpenGLES"
	case Qt__AA_UseSoftwareOpenGL: // 17
		return "AA_UseSoftwareOpenGL"
	case Qt__AA_ShareOpenGLContexts: // 18
		return "AA_ShareOpenGLContexts"
	case Qt__AA_SetPalette: // 19
		return "AA_SetPalette"
	case Qt__AA_EnableHighDpiScaling: // 20
		return "AA_EnableHighDpiScaling"
	case Qt__AA_DisableHighDpiScaling: // 21
		return "AA_DisableHighDpiScaling"
	case Qt__AA_UseStyleSheetPropagationInWidgetStyles: // 22
		return "AA_UseStyleSheetPropagationInWidgetStyles"
	case Qt__AA_DontUseNativeDialogs: // 23
		return "AA_DontUseNativeDialogs"
	case Qt__AA_SynthesizeMouseForUnhandledTabletEvents: // 24
		return "AA_SynthesizeMouseForUnhandledTabletEvents"
	case Qt__AA_CompressHighFrequencyEvents: // 25
		return "AA_CompressHighFrequencyEvents"
	case Qt__AA_DontCheckOpenGLContextThreadAffinity: // 26
		return "AA_DontCheckOpenGLContextThreadAffinity"
	case Qt__AA_DisableShaderDiskCache: // 27
		return "AA_DisableShaderDiskCache"
	case Qt__AA_DontShowShortcutsInContextMenus: // 28
		return "AA_DontShowShortcutsInContextMenus"
	case Qt__AA_CompressTabletEvents: // 29
		return "AA_CompressTabletEvents"
	case Qt__AA_DisableWindowContextHelpButton: // 30
		return "AA_DisableWindowContextHelpButton"
	case Qt__AA_AttributeCount: // 31
		return "AA_AttributeCount"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ImageConversionFlagItemName(val int) string {
	switch val {
	case Qt__ColorMode_Mask: // 3
		return "ColorMode_Mask,ColorOnly"
	case Qt__AutoColor: // 0
		return "AutoColor,ThresholdAlphaDither,DiffuseDither,AutoDither"
		// case Qt__ColorOnly: // 3
		// return ""
	case Qt__MonoOnly: // 2
		return "MonoOnly"
	case Qt__AlphaDither_Mask: // 12
		return "AlphaDither_Mask,NoAlpha"
		// case Qt__ThresholdAlphaDither: // 0
		// return ""
	case Qt__OrderedAlphaDither: // 4
		return "OrderedAlphaDither"
	case Qt__DiffuseAlphaDither: // 8
		return "DiffuseAlphaDither"
		// case Qt__NoAlpha: // 12
		// return ""
	case Qt__Dither_Mask: // 48
		return "Dither_Mask"
		// case Qt__DiffuseDither: // 0
		// return ""
	case Qt__OrderedDither: // 16
		return "OrderedDither"
	case Qt__ThresholdDither: // 32
		return "ThresholdDither"
	case Qt__DitherMode_Mask: // 192
		return "DitherMode_Mask"
		// case Qt__AutoDither: // 0
		// return ""
	case Qt__PreferDither: // 64
		return "PreferDither"
	case Qt__AvoidDither: // 128
		return "AvoidDither"
	case Qt__NoOpaqueDetection: // 256
		return "NoOpaqueDetection"
	case Qt__NoFormatConversion: // 512
		return "NoFormatConversion"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func BGModeItemName(val int) string {
	switch val {
	case Qt__TransparentMode: // 0
		return "TransparentMode"
	case Qt__OpaqueMode: // 1
		return "OpaqueMode"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func KeyItemName(val int) string {
	switch val {
	case Qt__Key_Escape: // 16777216
		return "Key_Escape"
	case Qt__Key_Tab: // 16777217
		return "Key_Tab"
	case Qt__Key_Backtab: // 16777218
		return "Key_Backtab"
	case Qt__Key_Backspace: // 16777219
		return "Key_Backspace"
	case Qt__Key_Return: // 16777220
		return "Key_Return"
	case Qt__Key_Enter: // 16777221
		return "Key_Enter"
	case Qt__Key_Insert: // 16777222
		return "Key_Insert"
	case Qt__Key_Delete: // 16777223
		return "Key_Delete"
	case Qt__Key_Pause: // 16777224
		return "Key_Pause"
	case Qt__Key_Print: // 16777225
		return "Key_Print"
	case Qt__Key_SysReq: // 16777226
		return "Key_SysReq"
	case Qt__Key_Clear: // 16777227
		return "Key_Clear"
	case Qt__Key_Home: // 16777232
		return "Key_Home"
	case Qt__Key_End: // 16777233
		return "Key_End"
	case Qt__Key_Left: // 16777234
		return "Key_Left"
	case Qt__Key_Up: // 16777235
		return "Key_Up"
	case Qt__Key_Right: // 16777236
		return "Key_Right"
	case Qt__Key_Down: // 16777237
		return "Key_Down"
	case Qt__Key_PageUp: // 16777238
		return "Key_PageUp"
	case Qt__Key_PageDown: // 16777239
		return "Key_PageDown"
	case Qt__Key_Shift: // 16777248
		return "Key_Shift"
	case Qt__Key_Control: // 16777249
		return "Key_Control"
	case Qt__Key_Meta: // 16777250
		return "Key_Meta"
	case Qt__Key_Alt: // 16777251
		return "Key_Alt"
	case Qt__Key_CapsLock: // 16777252
		return "Key_CapsLock"
	case Qt__Key_NumLock: // 16777253
		return "Key_NumLock"
	case Qt__Key_ScrollLock: // 16777254
		return "Key_ScrollLock"
	case Qt__Key_F1: // 16777264
		return "Key_F1"
	case Qt__Key_F2: // 16777265
		return "Key_F2"
	case Qt__Key_F3: // 16777266
		return "Key_F3"
	case Qt__Key_F4: // 16777267
		return "Key_F4"
	case Qt__Key_F5: // 16777268
		return "Key_F5"
	case Qt__Key_F6: // 16777269
		return "Key_F6"
	case Qt__Key_F7: // 16777270
		return "Key_F7"
	case Qt__Key_F8: // 16777271
		return "Key_F8"
	case Qt__Key_F9: // 16777272
		return "Key_F9"
	case Qt__Key_F10: // 16777273
		return "Key_F10"
	case Qt__Key_F11: // 16777274
		return "Key_F11"
	case Qt__Key_F12: // 16777275
		return "Key_F12"
	case Qt__Key_F13: // 16777276
		return "Key_F13"
	case Qt__Key_F14: // 16777277
		return "Key_F14"
	case Qt__Key_F15: // 16777278
		return "Key_F15"
	case Qt__Key_F16: // 16777279
		return "Key_F16"
	case Qt__Key_F17: // 16777280
		return "Key_F17"
	case Qt__Key_F18: // 16777281
		return "Key_F18"
	case Qt__Key_F19: // 16777282
		return "Key_F19"
	case Qt__Key_F20: // 16777283
		return "Key_F20"
	case Qt__Key_F21: // 16777284
		return "Key_F21"
	case Qt__Key_F22: // 16777285
		return "Key_F22"
	case Qt__Key_F23: // 16777286
		return "Key_F23"
	case Qt__Key_F24: // 16777287
		return "Key_F24"
	case Qt__Key_F25: // 16777288
		return "Key_F25"
	case Qt__Key_F26: // 16777289
		return "Key_F26"
	case Qt__Key_F27: // 16777290
		return "Key_F27"
	case Qt__Key_F28: // 16777291
		return "Key_F28"
	case Qt__Key_F29: // 16777292
		return "Key_F29"
	case Qt__Key_F30: // 16777293
		return "Key_F30"
	case Qt__Key_F31: // 16777294
		return "Key_F31"
	case Qt__Key_F32: // 16777295
		return "Key_F32"
	case Qt__Key_F33: // 16777296
		return "Key_F33"
	case Qt__Key_F34: // 16777297
		return "Key_F34"
	case Qt__Key_F35: // 16777298
		return "Key_F35"
	case Qt__Key_Super_L: // 16777299
		return "Key_Super_L"
	case Qt__Key_Super_R: // 16777300
		return "Key_Super_R"
	case Qt__Key_Menu: // 16777301
		return "Key_Menu"
	case Qt__Key_Hyper_L: // 16777302
		return "Key_Hyper_L"
	case Qt__Key_Hyper_R: // 16777303
		return "Key_Hyper_R"
	case Qt__Key_Help: // 16777304
		return "Key_Help"
	case Qt__Key_Direction_L: // 16777305
		return "Key_Direction_L"
	case Qt__Key_Direction_R: // 16777312
		return "Key_Direction_R"
	case Qt__Key_Space: // 32
		return "Key_Space,Key_Any"
		// case Qt__Key_Any: // 32
		// return ""
	case Qt__Key_Exclam: // 33
		return "Key_Exclam"
	case Qt__Key_QuoteDbl: // 34
		return "Key_QuoteDbl"
	case Qt__Key_NumberSign: // 35
		return "Key_NumberSign"
	case Qt__Key_Dollar: // 36
		return "Key_Dollar"
	case Qt__Key_Percent: // 37
		return "Key_Percent"
	case Qt__Key_Ampersand: // 38
		return "Key_Ampersand"
	case Qt__Key_Apostrophe: // 39
		return "Key_Apostrophe"
	case Qt__Key_ParenLeft: // 40
		return "Key_ParenLeft"
	case Qt__Key_ParenRight: // 41
		return "Key_ParenRight"
	case Qt__Key_Asterisk: // 42
		return "Key_Asterisk"
	case Qt__Key_Plus: // 43
		return "Key_Plus"
	case Qt__Key_Comma: // 44
		return "Key_Comma"
	case Qt__Key_Minus: // 45
		return "Key_Minus"
	case Qt__Key_Period: // 46
		return "Key_Period"
	case Qt__Key_Slash: // 47
		return "Key_Slash"
	case Qt__Key_0: // 48
		return "Key_0"
	case Qt__Key_1: // 49
		return "Key_1"
	case Qt__Key_2: // 50
		return "Key_2"
	case Qt__Key_3: // 51
		return "Key_3"
	case Qt__Key_4: // 52
		return "Key_4"
	case Qt__Key_5: // 53
		return "Key_5"
	case Qt__Key_6: // 54
		return "Key_6"
	case Qt__Key_7: // 55
		return "Key_7"
	case Qt__Key_8: // 56
		return "Key_8"
	case Qt__Key_9: // 57
		return "Key_9"
	case Qt__Key_Colon: // 58
		return "Key_Colon"
	case Qt__Key_Semicolon: // 59
		return "Key_Semicolon"
	case Qt__Key_Less: // 60
		return "Key_Less"
	case Qt__Key_Equal: // 61
		return "Key_Equal"
	case Qt__Key_Greater: // 62
		return "Key_Greater"
	case Qt__Key_Question: // 63
		return "Key_Question"
	case Qt__Key_At: // 64
		return "Key_At"
	case Qt__Key_A: // 65
		return "Key_A"
	case Qt__Key_B: // 66
		return "Key_B"
	case Qt__Key_C: // 67
		return "Key_C"
	case Qt__Key_D: // 68
		return "Key_D"
	case Qt__Key_E: // 69
		return "Key_E"
	case Qt__Key_F: // 70
		return "Key_F"
	case Qt__Key_G: // 71
		return "Key_G"
	case Qt__Key_H: // 72
		return "Key_H"
	case Qt__Key_I: // 73
		return "Key_I"
	case Qt__Key_J: // 74
		return "Key_J"
	case Qt__Key_K: // 75
		return "Key_K"
	case Qt__Key_L: // 76
		return "Key_L"
	case Qt__Key_M: // 77
		return "Key_M"
	case Qt__Key_N: // 78
		return "Key_N"
	case Qt__Key_O: // 79
		return "Key_O"
	case Qt__Key_P: // 80
		return "Key_P"
	case Qt__Key_Q: // 81
		return "Key_Q"
	case Qt__Key_R: // 82
		return "Key_R"
	case Qt__Key_S: // 83
		return "Key_S"
	case Qt__Key_T: // 84
		return "Key_T"
	case Qt__Key_U: // 85
		return "Key_U"
	case Qt__Key_V: // 86
		return "Key_V"
	case Qt__Key_W: // 87
		return "Key_W"
	case Qt__Key_X: // 88
		return "Key_X"
	case Qt__Key_Y: // 89
		return "Key_Y"
	case Qt__Key_Z: // 90
		return "Key_Z"
	case Qt__Key_BracketLeft: // 91
		return "Key_BracketLeft"
	case Qt__Key_Backslash: // 92
		return "Key_Backslash"
	case Qt__Key_BracketRight: // 93
		return "Key_BracketRight"
	case Qt__Key_AsciiCircum: // 94
		return "Key_AsciiCircum"
	case Qt__Key_Underscore: // 95
		return "Key_Underscore"
	case Qt__Key_QuoteLeft: // 96
		return "Key_QuoteLeft"
	case Qt__Key_BraceLeft: // 123
		return "Key_BraceLeft"
	case Qt__Key_Bar: // 124
		return "Key_Bar"
	case Qt__Key_BraceRight: // 125
		return "Key_BraceRight"
	case Qt__Key_AsciiTilde: // 126
		return "Key_AsciiTilde"
	case Qt__Key_nobreakspace: // 160
		return "Key_nobreakspace"
	case Qt__Key_exclamdown: // 161
		return "Key_exclamdown"
	case Qt__Key_cent: // 162
		return "Key_cent"
	case Qt__Key_sterling: // 163
		return "Key_sterling"
	case Qt__Key_currency: // 164
		return "Key_currency"
	case Qt__Key_yen: // 165
		return "Key_yen"
	case Qt__Key_brokenbar: // 166
		return "Key_brokenbar"
	case Qt__Key_section: // 167
		return "Key_section"
	case Qt__Key_diaeresis: // 168
		return "Key_diaeresis"
	case Qt__Key_copyright: // 169
		return "Key_copyright"
	case Qt__Key_ordfeminine: // 170
		return "Key_ordfeminine"
	case Qt__Key_guillemotleft: // 171
		return "Key_guillemotleft"
	case Qt__Key_notsign: // 172
		return "Key_notsign"
	case Qt__Key_hyphen: // 173
		return "Key_hyphen"
	case Qt__Key_registered: // 174
		return "Key_registered"
	case Qt__Key_macron: // 175
		return "Key_macron"
	case Qt__Key_degree: // 176
		return "Key_degree"
	case Qt__Key_plusminus: // 177
		return "Key_plusminus"
	case Qt__Key_twosuperior: // 178
		return "Key_twosuperior"
	case Qt__Key_threesuperior: // 179
		return "Key_threesuperior"
	case Qt__Key_acute: // 180
		return "Key_acute"
	case Qt__Key_mu: // 181
		return "Key_mu"
	case Qt__Key_paragraph: // 182
		return "Key_paragraph"
	case Qt__Key_periodcentered: // 183
		return "Key_periodcentered"
	case Qt__Key_cedilla: // 184
		return "Key_cedilla"
	case Qt__Key_onesuperior: // 185
		return "Key_onesuperior"
	case Qt__Key_masculine: // 186
		return "Key_masculine"
	case Qt__Key_guillemotright: // 187
		return "Key_guillemotright"
	case Qt__Key_onequarter: // 188
		return "Key_onequarter"
	case Qt__Key_onehalf: // 189
		return "Key_onehalf"
	case Qt__Key_threequarters: // 190
		return "Key_threequarters"
	case Qt__Key_questiondown: // 191
		return "Key_questiondown"
	case Qt__Key_Agrave: // 192
		return "Key_Agrave"
	case Qt__Key_Aacute: // 193
		return "Key_Aacute"
	case Qt__Key_Acircumflex: // 194
		return "Key_Acircumflex"
	case Qt__Key_Atilde: // 195
		return "Key_Atilde"
	case Qt__Key_Adiaeresis: // 196
		return "Key_Adiaeresis"
	case Qt__Key_Aring: // 197
		return "Key_Aring"
	case Qt__Key_AE: // 198
		return "Key_AE"
	case Qt__Key_Ccedilla: // 199
		return "Key_Ccedilla"
	case Qt__Key_Egrave: // 200
		return "Key_Egrave"
	case Qt__Key_Eacute: // 201
		return "Key_Eacute"
	case Qt__Key_Ecircumflex: // 202
		return "Key_Ecircumflex"
	case Qt__Key_Ediaeresis: // 203
		return "Key_Ediaeresis"
	case Qt__Key_Igrave: // 204
		return "Key_Igrave"
	case Qt__Key_Iacute: // 205
		return "Key_Iacute"
	case Qt__Key_Icircumflex: // 206
		return "Key_Icircumflex"
	case Qt__Key_Idiaeresis: // 207
		return "Key_Idiaeresis"
	case Qt__Key_ETH: // 208
		return "Key_ETH"
	case Qt__Key_Ntilde: // 209
		return "Key_Ntilde"
	case Qt__Key_Ograve: // 210
		return "Key_Ograve"
	case Qt__Key_Oacute: // 211
		return "Key_Oacute"
	case Qt__Key_Ocircumflex: // 212
		return "Key_Ocircumflex"
	case Qt__Key_Otilde: // 213
		return "Key_Otilde"
	case Qt__Key_Odiaeresis: // 214
		return "Key_Odiaeresis"
	case Qt__Key_multiply: // 215
		return "Key_multiply"
	case Qt__Key_Ooblique: // 216
		return "Key_Ooblique"
	case Qt__Key_Ugrave: // 217
		return "Key_Ugrave"
	case Qt__Key_Uacute: // 218
		return "Key_Uacute"
	case Qt__Key_Ucircumflex: // 219
		return "Key_Ucircumflex"
	case Qt__Key_Udiaeresis: // 220
		return "Key_Udiaeresis"
	case Qt__Key_Yacute: // 221
		return "Key_Yacute"
	case Qt__Key_THORN: // 222
		return "Key_THORN"
	case Qt__Key_ssharp: // 223
		return "Key_ssharp"
	case Qt__Key_division: // 247
		return "Key_division"
	case Qt__Key_ydiaeresis: // 255
		return "Key_ydiaeresis"
	case Qt__Key_AltGr: // 16781571
		return "Key_AltGr"
	case Qt__Key_Multi_key: // 16781600
		return "Key_Multi_key"
	case Qt__Key_Codeinput: // 16781623
		return "Key_Codeinput"
	case Qt__Key_SingleCandidate: // 16781628
		return "Key_SingleCandidate"
	case Qt__Key_MultipleCandidate: // 16781629
		return "Key_MultipleCandidate"
	case Qt__Key_PreviousCandidate: // 16781630
		return "Key_PreviousCandidate"
	case Qt__Key_Mode_switch: // 16781694
		return "Key_Mode_switch"
	case Qt__Key_Kanji: // 16781601
		return "Key_Kanji"
	case Qt__Key_Muhenkan: // 16781602
		return "Key_Muhenkan"
	case Qt__Key_Henkan: // 16781603
		return "Key_Henkan"
	case Qt__Key_Romaji: // 16781604
		return "Key_Romaji"
	case Qt__Key_Hiragana: // 16781605
		return "Key_Hiragana"
	case Qt__Key_Katakana: // 16781606
		return "Key_Katakana"
	case Qt__Key_Hiragana_Katakana: // 16781607
		return "Key_Hiragana_Katakana"
	case Qt__Key_Zenkaku: // 16781608
		return "Key_Zenkaku"
	case Qt__Key_Hankaku: // 16781609
		return "Key_Hankaku"
	case Qt__Key_Zenkaku_Hankaku: // 16781610
		return "Key_Zenkaku_Hankaku"
	case Qt__Key_Touroku: // 16781611
		return "Key_Touroku"
	case Qt__Key_Massyo: // 16781612
		return "Key_Massyo"
	case Qt__Key_Kana_Lock: // 16781613
		return "Key_Kana_Lock"
	case Qt__Key_Kana_Shift: // 16781614
		return "Key_Kana_Shift"
	case Qt__Key_Eisu_Shift: // 16781615
		return "Key_Eisu_Shift"
	case Qt__Key_Eisu_toggle: // 16781616
		return "Key_Eisu_toggle"
	case Qt__Key_Hangul: // 16781617
		return "Key_Hangul"
	case Qt__Key_Hangul_Start: // 16781618
		return "Key_Hangul_Start"
	case Qt__Key_Hangul_End: // 16781619
		return "Key_Hangul_End"
	case Qt__Key_Hangul_Hanja: // 16781620
		return "Key_Hangul_Hanja"
	case Qt__Key_Hangul_Jamo: // 16781621
		return "Key_Hangul_Jamo"
	case Qt__Key_Hangul_Romaja: // 16781622
		return "Key_Hangul_Romaja"
	case Qt__Key_Hangul_Jeonja: // 16781624
		return "Key_Hangul_Jeonja"
	case Qt__Key_Hangul_Banja: // 16781625
		return "Key_Hangul_Banja"
	case Qt__Key_Hangul_PreHanja: // 16781626
		return "Key_Hangul_PreHanja"
	case Qt__Key_Hangul_PostHanja: // 16781627
		return "Key_Hangul_PostHanja"
	case Qt__Key_Hangul_Special: // 16781631
		return "Key_Hangul_Special"
	case Qt__Key_Dead_Grave: // 16781904
		return "Key_Dead_Grave"
	case Qt__Key_Dead_Acute: // 16781905
		return "Key_Dead_Acute"
	case Qt__Key_Dead_Circumflex: // 16781906
		return "Key_Dead_Circumflex"
	case Qt__Key_Dead_Tilde: // 16781907
		return "Key_Dead_Tilde"
	case Qt__Key_Dead_Macron: // 16781908
		return "Key_Dead_Macron"
	case Qt__Key_Dead_Breve: // 16781909
		return "Key_Dead_Breve"
	case Qt__Key_Dead_Abovedot: // 16781910
		return "Key_Dead_Abovedot"
	case Qt__Key_Dead_Diaeresis: // 16781911
		return "Key_Dead_Diaeresis"
	case Qt__Key_Dead_Abovering: // 16781912
		return "Key_Dead_Abovering"
	case Qt__Key_Dead_Doubleacute: // 16781913
		return "Key_Dead_Doubleacute"
	case Qt__Key_Dead_Caron: // 16781914
		return "Key_Dead_Caron"
	case Qt__Key_Dead_Cedilla: // 16781915
		return "Key_Dead_Cedilla"
	case Qt__Key_Dead_Ogonek: // 16781916
		return "Key_Dead_Ogonek"
	case Qt__Key_Dead_Iota: // 16781917
		return "Key_Dead_Iota"
	case Qt__Key_Dead_Voiced_Sound: // 16781918
		return "Key_Dead_Voiced_Sound"
	case Qt__Key_Dead_Semivoiced_Sound: // 16781919
		return "Key_Dead_Semivoiced_Sound"
	case Qt__Key_Dead_Belowdot: // 16781920
		return "Key_Dead_Belowdot"
	case Qt__Key_Dead_Hook: // 16781921
		return "Key_Dead_Hook"
	case Qt__Key_Dead_Horn: // 16781922
		return "Key_Dead_Horn"
	case Qt__Key_Back: // 16777313
		return "Key_Back"
	case Qt__Key_Forward: // 16777314
		return "Key_Forward"
	case Qt__Key_Stop: // 16777315
		return "Key_Stop"
	case Qt__Key_Refresh: // 16777316
		return "Key_Refresh"
	case Qt__Key_VolumeDown: // 16777328
		return "Key_VolumeDown"
	case Qt__Key_VolumeMute: // 16777329
		return "Key_VolumeMute"
	case Qt__Key_VolumeUp: // 16777330
		return "Key_VolumeUp"
	case Qt__Key_BassBoost: // 16777331
		return "Key_BassBoost"
	case Qt__Key_BassUp: // 16777332
		return "Key_BassUp"
	case Qt__Key_BassDown: // 16777333
		return "Key_BassDown"
	case Qt__Key_TrebleUp: // 16777334
		return "Key_TrebleUp"
	case Qt__Key_TrebleDown: // 16777335
		return "Key_TrebleDown"
	case Qt__Key_MediaPlay: // 16777344
		return "Key_MediaPlay"
	case Qt__Key_MediaStop: // 16777345
		return "Key_MediaStop"
	case Qt__Key_MediaPrevious: // 16777346
		return "Key_MediaPrevious"
	case Qt__Key_MediaNext: // 16777347
		return "Key_MediaNext"
	case Qt__Key_MediaRecord: // 16777348
		return "Key_MediaRecord"
	case Qt__Key_MediaPause: // 16777349
		return "Key_MediaPause"
	case Qt__Key_MediaTogglePlayPause: // 16777350
		return "Key_MediaTogglePlayPause"
	case Qt__Key_HomePage: // 16777360
		return "Key_HomePage"
	case Qt__Key_Favorites: // 16777361
		return "Key_Favorites"
	case Qt__Key_Search: // 16777362
		return "Key_Search"
	case Qt__Key_Standby: // 16777363
		return "Key_Standby"
	case Qt__Key_OpenUrl: // 16777364
		return "Key_OpenUrl"
	case Qt__Key_LaunchMail: // 16777376
		return "Key_LaunchMail"
	case Qt__Key_LaunchMedia: // 16777377
		return "Key_LaunchMedia"
	case Qt__Key_Launch0: // 16777378
		return "Key_Launch0"
	case Qt__Key_Launch1: // 16777379
		return "Key_Launch1"
	case Qt__Key_Launch2: // 16777380
		return "Key_Launch2"
	case Qt__Key_Launch3: // 16777381
		return "Key_Launch3"
	case Qt__Key_Launch4: // 16777382
		return "Key_Launch4"
	case Qt__Key_Launch5: // 16777383
		return "Key_Launch5"
	case Qt__Key_Launch6: // 16777384
		return "Key_Launch6"
	case Qt__Key_Launch7: // 16777385
		return "Key_Launch7"
	case Qt__Key_Launch8: // 16777386
		return "Key_Launch8"
	case Qt__Key_Launch9: // 16777387
		return "Key_Launch9"
	case Qt__Key_LaunchA: // 16777388
		return "Key_LaunchA"
	case Qt__Key_LaunchB: // 16777389
		return "Key_LaunchB"
	case Qt__Key_LaunchC: // 16777390
		return "Key_LaunchC"
	case Qt__Key_LaunchD: // 16777391
		return "Key_LaunchD"
	case Qt__Key_LaunchE: // 16777392
		return "Key_LaunchE"
	case Qt__Key_LaunchF: // 16777393
		return "Key_LaunchF"
	case Qt__Key_MonBrightnessUp: // 16777394
		return "Key_MonBrightnessUp"
	case Qt__Key_MonBrightnessDown: // 16777395
		return "Key_MonBrightnessDown"
	case Qt__Key_KeyboardLightOnOff: // 16777396
		return "Key_KeyboardLightOnOff"
	case Qt__Key_KeyboardBrightnessUp: // 16777397
		return "Key_KeyboardBrightnessUp"
	case Qt__Key_KeyboardBrightnessDown: // 16777398
		return "Key_KeyboardBrightnessDown"
	case Qt__Key_PowerOff: // 16777399
		return "Key_PowerOff"
	case Qt__Key_WakeUp: // 16777400
		return "Key_WakeUp"
	case Qt__Key_Eject: // 16777401
		return "Key_Eject"
	case Qt__Key_ScreenSaver: // 16777402
		return "Key_ScreenSaver"
	case Qt__Key_WWW: // 16777403
		return "Key_WWW"
	case Qt__Key_Memo: // 16777404
		return "Key_Memo"
	case Qt__Key_LightBulb: // 16777405
		return "Key_LightBulb"
	case Qt__Key_Shop: // 16777406
		return "Key_Shop"
	case Qt__Key_History: // 16777407
		return "Key_History"
	case Qt__Key_AddFavorite: // 16777408
		return "Key_AddFavorite"
	case Qt__Key_HotLinks: // 16777409
		return "Key_HotLinks"
	case Qt__Key_BrightnessAdjust: // 16777410
		return "Key_BrightnessAdjust"
	case Qt__Key_Finance: // 16777411
		return "Key_Finance"
	case Qt__Key_Community: // 16777412
		return "Key_Community"
	case Qt__Key_AudioRewind: // 16777413
		return "Key_AudioRewind"
	case Qt__Key_BackForward: // 16777414
		return "Key_BackForward"
	case Qt__Key_ApplicationLeft: // 16777415
		return "Key_ApplicationLeft"
	case Qt__Key_ApplicationRight: // 16777416
		return "Key_ApplicationRight"
	case Qt__Key_Book: // 16777417
		return "Key_Book"
	case Qt__Key_CD: // 16777418
		return "Key_CD"
	case Qt__Key_Calculator: // 16777419
		return "Key_Calculator"
	case Qt__Key_ToDoList: // 16777420
		return "Key_ToDoList"
	case Qt__Key_ClearGrab: // 16777421
		return "Key_ClearGrab"
	case Qt__Key_Close: // 16777422
		return "Key_Close"
	case Qt__Key_Copy: // 16777423
		return "Key_Copy"
	case Qt__Key_Cut: // 16777424
		return "Key_Cut"
	case Qt__Key_Display: // 16777425
		return "Key_Display"
	case Qt__Key_DOS: // 16777426
		return "Key_DOS"
	case Qt__Key_Documents: // 16777427
		return "Key_Documents"
	case Qt__Key_Excel: // 16777428
		return "Key_Excel"
	case Qt__Key_Explorer: // 16777429
		return "Key_Explorer"
	case Qt__Key_Game: // 16777430
		return "Key_Game"
	case Qt__Key_Go: // 16777431
		return "Key_Go"
	case Qt__Key_iTouch: // 16777432
		return "Key_iTouch"
	case Qt__Key_LogOff: // 16777433
		return "Key_LogOff"
	case Qt__Key_Market: // 16777434
		return "Key_Market"
	case Qt__Key_Meeting: // 16777435
		return "Key_Meeting"
	case Qt__Key_MenuKB: // 16777436
		return "Key_MenuKB"
	case Qt__Key_MenuPB: // 16777437
		return "Key_MenuPB"
	case Qt__Key_MySites: // 16777438
		return "Key_MySites"
	case Qt__Key_News: // 16777439
		return "Key_News"
	case Qt__Key_OfficeHome: // 16777440
		return "Key_OfficeHome"
	case Qt__Key_Option: // 16777441
		return "Key_Option"
	case Qt__Key_Paste: // 16777442
		return "Key_Paste"
	case Qt__Key_Phone: // 16777443
		return "Key_Phone"
	case Qt__Key_Calendar: // 16777444
		return "Key_Calendar"
	case Qt__Key_Reply: // 16777445
		return "Key_Reply"
	case Qt__Key_Reload: // 16777446
		return "Key_Reload"
	case Qt__Key_RotateWindows: // 16777447
		return "Key_RotateWindows"
	case Qt__Key_RotationPB: // 16777448
		return "Key_RotationPB"
	case Qt__Key_RotationKB: // 16777449
		return "Key_RotationKB"
	case Qt__Key_Save: // 16777450
		return "Key_Save"
	case Qt__Key_Send: // 16777451
		return "Key_Send"
	case Qt__Key_Spell: // 16777452
		return "Key_Spell"
	case Qt__Key_SplitScreen: // 16777453
		return "Key_SplitScreen"
	case Qt__Key_Support: // 16777454
		return "Key_Support"
	case Qt__Key_TaskPane: // 16777455
		return "Key_TaskPane"
	case Qt__Key_Terminal: // 16777456
		return "Key_Terminal"
	case Qt__Key_Tools: // 16777457
		return "Key_Tools"
	case Qt__Key_Travel: // 16777458
		return "Key_Travel"
	case Qt__Key_Video: // 16777459
		return "Key_Video"
	case Qt__Key_Word: // 16777460
		return "Key_Word"
	case Qt__Key_Xfer: // 16777461
		return "Key_Xfer"
	case Qt__Key_ZoomIn: // 16777462
		return "Key_ZoomIn"
	case Qt__Key_ZoomOut: // 16777463
		return "Key_ZoomOut"
	case Qt__Key_Away: // 16777464
		return "Key_Away"
	case Qt__Key_Messenger: // 16777465
		return "Key_Messenger"
	case Qt__Key_WebCam: // 16777466
		return "Key_WebCam"
	case Qt__Key_MailForward: // 16777467
		return "Key_MailForward"
	case Qt__Key_Pictures: // 16777468
		return "Key_Pictures"
	case Qt__Key_Music: // 16777469
		return "Key_Music"
	case Qt__Key_Battery: // 16777470
		return "Key_Battery"
	case Qt__Key_Bluetooth: // 16777471
		return "Key_Bluetooth"
	case Qt__Key_WLAN: // 16777472
		return "Key_WLAN"
	case Qt__Key_UWB: // 16777473
		return "Key_UWB"
	case Qt__Key_AudioForward: // 16777474
		return "Key_AudioForward"
	case Qt__Key_AudioRepeat: // 16777475
		return "Key_AudioRepeat"
	case Qt__Key_AudioRandomPlay: // 16777476
		return "Key_AudioRandomPlay"
	case Qt__Key_Subtitle: // 16777477
		return "Key_Subtitle"
	case Qt__Key_AudioCycleTrack: // 16777478
		return "Key_AudioCycleTrack"
	case Qt__Key_Time: // 16777479
		return "Key_Time"
	case Qt__Key_Hibernate: // 16777480
		return "Key_Hibernate"
	case Qt__Key_View: // 16777481
		return "Key_View"
	case Qt__Key_TopMenu: // 16777482
		return "Key_TopMenu"
	case Qt__Key_PowerDown: // 16777483
		return "Key_PowerDown"
	case Qt__Key_Suspend: // 16777484
		return "Key_Suspend"
	case Qt__Key_ContrastAdjust: // 16777485
		return "Key_ContrastAdjust"
	case Qt__Key_LaunchG: // 16777486
		return "Key_LaunchG"
	case Qt__Key_LaunchH: // 16777487
		return "Key_LaunchH"
	case Qt__Key_TouchpadToggle: // 16777488
		return "Key_TouchpadToggle"
	case Qt__Key_TouchpadOn: // 16777489
		return "Key_TouchpadOn"
	case Qt__Key_TouchpadOff: // 16777490
		return "Key_TouchpadOff"
	case Qt__Key_MicMute: // 16777491
		return "Key_MicMute"
	case Qt__Key_Red: // 16777492
		return "Key_Red"
	case Qt__Key_Green: // 16777493
		return "Key_Green"
	case Qt__Key_Yellow: // 16777494
		return "Key_Yellow"
	case Qt__Key_Blue: // 16777495
		return "Key_Blue"
	case Qt__Key_ChannelUp: // 16777496
		return "Key_ChannelUp"
	case Qt__Key_ChannelDown: // 16777497
		return "Key_ChannelDown"
	case Qt__Key_Guide: // 16777498
		return "Key_Guide"
	case Qt__Key_Info: // 16777499
		return "Key_Info"
	case Qt__Key_Settings: // 16777500
		return "Key_Settings"
	case Qt__Key_MicVolumeUp: // 16777501
		return "Key_MicVolumeUp"
	case Qt__Key_MicVolumeDown: // 16777502
		return "Key_MicVolumeDown"
	case Qt__Key_New: // 16777504
		return "Key_New"
	case Qt__Key_Open: // 16777505
		return "Key_Open"
	case Qt__Key_Find: // 16777506
		return "Key_Find"
	case Qt__Key_Undo: // 16777507
		return "Key_Undo"
	case Qt__Key_Redo: // 16777508
		return "Key_Redo"
	case Qt__Key_MediaLast: // 16842751
		return "Key_MediaLast"
	case Qt__Key_Select: // 16842752
		return "Key_Select"
	case Qt__Key_Yes: // 16842753
		return "Key_Yes"
	case Qt__Key_No: // 16842754
		return "Key_No"
	case Qt__Key_Cancel: // 16908289
		return "Key_Cancel"
	case Qt__Key_Printer: // 16908290
		return "Key_Printer"
	case Qt__Key_Execute: // 16908291
		return "Key_Execute"
	case Qt__Key_Sleep: // 16908292
		return "Key_Sleep"
	case Qt__Key_Play: // 16908293
		return "Key_Play"
	case Qt__Key_Zoom: // 16908294
		return "Key_Zoom"
	case Qt__Key_Exit: // 16908298
		return "Key_Exit"
	case Qt__Key_Context1: // 17825792
		return "Key_Context1"
	case Qt__Key_Context2: // 17825793
		return "Key_Context2"
	case Qt__Key_Context3: // 17825794
		return "Key_Context3"
	case Qt__Key_Context4: // 17825795
		return "Key_Context4"
	case Qt__Key_Call: // 17825796
		return "Key_Call"
	case Qt__Key_Hangup: // 17825797
		return "Key_Hangup"
	case Qt__Key_Flip: // 17825798
		return "Key_Flip"
	case Qt__Key_ToggleCallHangup: // 17825799
		return "Key_ToggleCallHangup"
	case Qt__Key_VoiceDial: // 17825800
		return "Key_VoiceDial"
	case Qt__Key_LastNumberRedial: // 17825801
		return "Key_LastNumberRedial"
	case Qt__Key_Camera: // 17825824
		return "Key_Camera"
	case Qt__Key_CameraFocus: // 17825825
		return "Key_CameraFocus"
	case Qt__Key_unknown: // 33554431
		return "Key_unknown"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ArrowTypeItemName(val int) string {
	switch val {
	case Qt__NoArrow: // 0
		return "NoArrow"
	case Qt__UpArrow: // 1
		return "UpArrow"
	case Qt__DownArrow: // 2
		return "DownArrow"
	case Qt__LeftArrow: // 3
		return "LeftArrow"
	case Qt__RightArrow: // 4
		return "RightArrow"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func PenStyleItemName(val int) string {
	switch val {
	case Qt__NoPen: // 0
		return "NoPen"
	case Qt__SolidLine: // 1
		return "SolidLine"
	case Qt__DashLine: // 2
		return "DashLine"
	case Qt__DotLine: // 3
		return "DotLine"
	case Qt__DashDotLine: // 4
		return "DashDotLine"
	case Qt__DashDotDotLine: // 5
		return "DashDotDotLine"
	case Qt__CustomDashLine: // 6
		return "CustomDashLine"
	case Qt__MPenStyle: // 15
		return "MPenStyle"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func PenCapStyleItemName(val int) string {
	switch val {
	case Qt__FlatCap: // 0
		return "FlatCap"
	case Qt__SquareCap: // 16
		return "SquareCap"
	case Qt__RoundCap: // 32
		return "RoundCap"
	case Qt__MPenCapStyle: // 48
		return "MPenCapStyle"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func PenJoinStyleItemName(val int) string {
	switch val {
	case Qt__MiterJoin: // 0
		return "MiterJoin"
	case Qt__BevelJoin: // 64
		return "BevelJoin"
	case Qt__RoundJoin: // 128
		return "RoundJoin"
	case Qt__SvgMiterJoin: // 256
		return "SvgMiterJoin"
	case Qt__MPenJoinStyle: // 448
		return "MPenJoinStyle"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func BrushStyleItemName(val int) string {
	switch val {
	case Qt__NoBrush: // 0
		return "NoBrush"
	case Qt__SolidPattern: // 1
		return "SolidPattern"
	case Qt__Dense1Pattern: // 2
		return "Dense1Pattern"
	case Qt__Dense2Pattern: // 3
		return "Dense2Pattern"
	case Qt__Dense3Pattern: // 4
		return "Dense3Pattern"
	case Qt__Dense4Pattern: // 5
		return "Dense4Pattern"
	case Qt__Dense5Pattern: // 6
		return "Dense5Pattern"
	case Qt__Dense6Pattern: // 7
		return "Dense6Pattern"
	case Qt__Dense7Pattern: // 8
		return "Dense7Pattern"
	case Qt__HorPattern: // 9
		return "HorPattern"
	case Qt__VerPattern: // 10
		return "VerPattern"
	case Qt__CrossPattern: // 11
		return "CrossPattern"
	case Qt__BDiagPattern: // 12
		return "BDiagPattern"
	case Qt__FDiagPattern: // 13
		return "FDiagPattern"
	case Qt__DiagCrossPattern: // 14
		return "DiagCrossPattern"
	case Qt__LinearGradientPattern: // 15
		return "LinearGradientPattern"
	case Qt__RadialGradientPattern: // 16
		return "RadialGradientPattern"
	case Qt__ConicalGradientPattern: // 17
		return "ConicalGradientPattern"
	case Qt__TexturePattern: // 24
		return "TexturePattern"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum is used by QPainter::drawRoundedRect() and QPainterPath::addRoundedRect() functions to specify the radii of rectangle corners with respect to the dimensions of the bounding rectangles specified.



This enum was introduced or modified in  Qt 4.4.

*/
type Qt__SizeMode = int // core
// Specifies the size using absolute measurements.
const Qt__AbsoluteSize Qt__SizeMode = 0

// Specifies the size relative to the bounding rectangle, typically using percentage measurements.
const Qt__RelativeSize Qt__SizeMode = 1

func SizeModeItemName(val int) string {
	switch val {
	case Qt__AbsoluteSize: // 0
		return "AbsoluteSize"
	case Qt__RelativeSize: // 1
		return "RelativeSize"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func UIEffectItemName(val int) string {
	switch val {
	case Qt__UI_General: // 0
		return "UI_General"
	case Qt__UI_AnimateMenu: // 1
		return "UI_AnimateMenu"
	case Qt__UI_FadeMenu: // 2
		return "UI_FadeMenu"
	case Qt__UI_AnimateCombo: // 3
		return "UI_AnimateCombo"
	case Qt__UI_AnimateTooltip: // 4
		return "UI_AnimateTooltip"
	case Qt__UI_FadeTooltip: // 5
		return "UI_FadeTooltip"
	case Qt__UI_AnimateToolBox: // 6
		return "UI_AnimateToolBox"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func CursorShapeItemName(val int) string {
	switch val {
	case Qt__ArrowCursor: // 0
		return "ArrowCursor"
	case Qt__UpArrowCursor: // 1
		return "UpArrowCursor"
	case Qt__CrossCursor: // 2
		return "CrossCursor"
	case Qt__WaitCursor: // 3
		return "WaitCursor"
	case Qt__IBeamCursor: // 4
		return "IBeamCursor"
	case Qt__SizeVerCursor: // 5
		return "SizeVerCursor"
	case Qt__SizeHorCursor: // 6
		return "SizeHorCursor"
	case Qt__SizeBDiagCursor: // 7
		return "SizeBDiagCursor"
	case Qt__SizeFDiagCursor: // 8
		return "SizeFDiagCursor"
	case Qt__SizeAllCursor: // 9
		return "SizeAllCursor"
	case Qt__BlankCursor: // 10
		return "BlankCursor"
	case Qt__SplitVCursor: // 11
		return "SplitVCursor"
	case Qt__SplitHCursor: // 12
		return "SplitHCursor"
	case Qt__PointingHandCursor: // 13
		return "PointingHandCursor"
	case Qt__ForbiddenCursor: // 14
		return "ForbiddenCursor"
	case Qt__WhatsThisCursor: // 15
		return "WhatsThisCursor"
	case Qt__BusyCursor: // 16
		return "BusyCursor"
	case Qt__OpenHandCursor: // 17
		return "OpenHandCursor"
	case Qt__ClosedHandCursor: // 18
		return "ClosedHandCursor"
	case Qt__DragCopyCursor: // 19
		return "DragCopyCursor"
	case Qt__DragMoveCursor: // 20
		return "DragMoveCursor"
	case Qt__DragLinkCursor: // 21
		return "DragLinkCursor,LastCursor"
		// case Qt__LastCursor: // 21
		// return ""
	case Qt__BitmapCursor: // 24
		return "BitmapCursor"
	case Qt__CustomCursor: // 25
		return "CustomCursor"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TextFormatItemName(val int) string {
	switch val {
	case Qt__PlainText: // 0
		return "PlainText"
	case Qt__RichText: // 1
		return "RichText"
	case Qt__AutoText: // 2
		return "AutoText"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func AspectRatioModeItemName(val int) string {
	switch val {
	case Qt__IgnoreAspectRatio: // 0
		return "IgnoreAspectRatio"
	case Qt__KeepAspectRatio: // 1
		return "KeepAspectRatio"
	case Qt__KeepAspectRatioByExpanding: // 2
		return "KeepAspectRatioByExpanding"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func DockWidgetAreaItemName(val int) string {
	switch val {
	case Qt__LeftDockWidgetArea: // 1
		return "LeftDockWidgetArea"
	case Qt__RightDockWidgetArea: // 2
		return "RightDockWidgetArea"
	case Qt__TopDockWidgetArea: // 4
		return "TopDockWidgetArea"
	case Qt__BottomDockWidgetArea: // 8
		return "BottomDockWidgetArea"
	case Qt__DockWidgetArea_Mask: // 15
		return "DockWidgetArea_Mask,AllDockWidgetAreas"
		// case Qt__AllDockWidgetAreas: // 15
		// return ""
	case Qt__NoDockWidgetArea: // 0
		return "NoDockWidgetArea"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__DockWidgetAreaSizes = int // core
//
const Qt__NDockWidgetAreas Qt__DockWidgetAreaSizes = 4

func DockWidgetAreaSizesItemName(val int) string {
	switch val {
	case Qt__NDockWidgetAreas: // 4
		return "NDockWidgetAreas"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ToolBarAreaItemName(val int) string {
	switch val {
	case Qt__LeftToolBarArea: // 1
		return "LeftToolBarArea"
	case Qt__RightToolBarArea: // 2
		return "RightToolBarArea"
	case Qt__TopToolBarArea: // 4
		return "TopToolBarArea"
	case Qt__BottomToolBarArea: // 8
		return "BottomToolBarArea"
	case Qt__ToolBarArea_Mask: // 15
		return "ToolBarArea_Mask,AllToolBarAreas"
		// case Qt__AllToolBarAreas: // 15
		// return ""
	case Qt__NoToolBarArea: // 0
		return "NoToolBarArea"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__ToolBarAreaSizes = int // core
//
const Qt__NToolBarAreas Qt__ToolBarAreaSizes = 4

func ToolBarAreaSizesItemName(val int) string {
	switch val {
	case Qt__NToolBarAreas: // 4
		return "NToolBarAreas"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func DateFormatItemName(val int) string {
	switch val {
	case Qt__TextDate: // 0
		return "TextDate"
	case Qt__ISODate: // 1
		return "ISODate"
	case Qt__SystemLocaleDate: // 2
		return "SystemLocaleDate,LocalDate"
		// case Qt__LocalDate: // 2
		// return ""
	case Qt__LocaleDate: // 3
		return "LocaleDate"
	case Qt__SystemLocaleShortDate: // 4
		return "SystemLocaleShortDate"
	case Qt__SystemLocaleLongDate: // 5
		return "SystemLocaleLongDate"
	case Qt__DefaultLocaleShortDate: // 6
		return "DefaultLocaleShortDate"
	case Qt__DefaultLocaleLongDate: // 7
		return "DefaultLocaleLongDate"
	case Qt__RFC2822Date: // 8
		return "RFC2822Date"
	case Qt__ISODateWithMs: // 9
		return "ISODateWithMs"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TimeSpecItemName(val int) string {
	switch val {
	case Qt__LocalTime: // 0
		return "LocalTime"
	case Qt__UTC: // 1
		return "UTC"
	case Qt__OffsetFromUTC: // 2
		return "OffsetFromUTC"
	case Qt__TimeZone: // 3
		return "TimeZone"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func DayOfWeekItemName(val int) string {
	switch val {
	case Qt__Monday: // 1
		return "Monday"
	case Qt__Tuesday: // 2
		return "Tuesday"
	case Qt__Wednesday: // 3
		return "Wednesday"
	case Qt__Thursday: // 4
		return "Thursday"
	case Qt__Friday: // 5
		return "Friday"
	case Qt__Saturday: // 6
		return "Saturday"
	case Qt__Sunday: // 7
		return "Sunday"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ScrollBarPolicyItemName(val int) string {
	switch val {
	case Qt__ScrollBarAsNeeded: // 0
		return "ScrollBarAsNeeded"
	case Qt__ScrollBarAlwaysOff: // 1
		return "ScrollBarAlwaysOff"
	case Qt__ScrollBarAlwaysOn: // 2
		return "ScrollBarAlwaysOn"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func CaseSensitivityItemName(val int) string {
	switch val {
	case Qt__CaseInsensitive: // 0
		return "CaseInsensitive"
	case Qt__CaseSensitive: // 1
		return "CaseSensitive"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func CornerItemName(val int) string {
	switch val {
	case Qt__TopLeftCorner: // 0
		return "TopLeftCorner"
	case Qt__TopRightCorner: // 1
		return "TopRightCorner"
	case Qt__BottomLeftCorner: // 2
		return "BottomLeftCorner"
	case Qt__BottomRightCorner: // 3
		return "BottomRightCorner"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func EdgeItemName(val int) string {
	switch val {
	case Qt__TopEdge: // 1
		return "TopEdge"
	case Qt__LeftEdge: // 2
		return "LeftEdge"
	case Qt__RightEdge: // 4
		return "RightEdge"
	case Qt__BottomEdge: // 8
		return "BottomEdge"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ConnectionTypeItemName(val int) string {
	switch val {
	case Qt__AutoConnection: // 0
		return "AutoConnection"
	case Qt__DirectConnection: // 1
		return "DirectConnection"
	case Qt__QueuedConnection: // 2
		return "QueuedConnection"
	case Qt__BlockingQueuedConnection: // 3
		return "BlockingQueuedConnection"
	case Qt__UniqueConnection: // 128
		return "UniqueConnection"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ShortcutContextItemName(val int) string {
	switch val {
	case Qt__WidgetShortcut: // 0
		return "WidgetShortcut"
	case Qt__WindowShortcut: // 1
		return "WindowShortcut"
	case Qt__ApplicationShortcut: // 2
		return "ApplicationShortcut"
	case Qt__WidgetWithChildrenShortcut: // 3
		return "WidgetWithChildrenShortcut"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
Specifies which method should be used to fill the paths and polygons.


*/
type Qt__FillRule = int // core
// Specifies that the region is filled using the odd even fill rule. With this rule, we determine whether a point is inside the shape by using the following method. Draw a horizontal line from the point to a location outside the shape, and count the number of intersections. If the number of intersections is an odd number, the point is inside the shape. This mode is the default.
const Qt__OddEvenFill Qt__FillRule = 0

// Specifies that the region is filled using the non zero winding rule. With this rule, we determine whether a point is inside the shape by using the following method. Draw a horizontal line from the point to a location outside the shape. Determine whether the direction of the line at each intersection point is up or down. The winding number is determined by summing the direction of each intersection. If the number is non zero, the point is inside the shape. This fill mode can also in most cases be considered as the intersection of closed shapes.
const Qt__WindingFill Qt__FillRule = 1

func FillRuleItemName(val int) string {
	switch val {
	case Qt__OddEvenFill: // 0
		return "OddEvenFill"
	case Qt__WindingFill: // 1
		return "WindingFill"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum specifies the behavior of the QPixmap::createMaskFromColor() and QImage::createMaskFromColor() functions.


*/
type Qt__MaskMode = int // core
// Creates a mask where all pixels matching the given color are opaque.
const Qt__MaskInColor Qt__MaskMode = 0

// Creates a mask where all pixels matching the given color are transparent.
const Qt__MaskOutColor Qt__MaskMode = 1

func MaskModeItemName(val int) string {
	switch val {
	case Qt__MaskInColor: // 0
		return "MaskInColor"
	case Qt__MaskOutColor: // 1
		return "MaskOutColor"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*

 */
type Qt__ClipOperation = int // core
// This operation turns clipping off.
const Qt__NoClip Qt__ClipOperation = 0

// Replaces the current clip path/rect/region with the one supplied in the function call.
const Qt__ReplaceClip Qt__ClipOperation = 1

// Intersects the current clip path/rect/region with the one supplied in the function call.
const Qt__IntersectClip Qt__ClipOperation = 2

func ClipOperationItemName(val int) string {
	switch val {
	case Qt__NoClip: // 0
		return "NoClip"
	case Qt__ReplaceClip: // 1
		return "ReplaceClip"
	case Qt__IntersectClip: // 2
		return "IntersectClip"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ItemSelectionModeItemName(val int) string {
	switch val {
	case Qt__ContainsItemShape: // 0
		return "ContainsItemShape"
	case Qt__IntersectsItemShape: // 1
		return "IntersectsItemShape"
	case Qt__ContainsItemBoundingRect: // 2
		return "ContainsItemBoundingRect"
	case Qt__IntersectsItemBoundingRect: // 3
		return "IntersectsItemBoundingRect"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum is used in QGraphicsScene to specify what to do with currently selected items when setting a selection area.



See also QGraphicsScene::setSelectionArea().

*/
type Qt__ItemSelectionOperation = int // core
// The currently selected items are replaced by items in the selection area.
const Qt__ReplaceSelection Qt__ItemSelectionOperation = 0

// The items in the selection area are added to the currently selected items.
const Qt__AddToSelection Qt__ItemSelectionOperation = 1

func ItemSelectionOperationItemName(val int) string {
	switch val {
	case Qt__ReplaceSelection: // 0
		return "ReplaceSelection"
	case Qt__AddToSelection: // 1
		return "AddToSelection"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum type defines whether image transformations (e.g., scaling) should be smooth or not.



See also QImage::scaled().

*/
type Qt__TransformationMode = int // core
// The transformation is performed quickly, with no smoothing.
const Qt__FastTransformation Qt__TransformationMode = 0

// The resulting image is transformed using bilinear filtering.
const Qt__SmoothTransformation Qt__TransformationMode = 1

func TransformationModeItemName(val int) string {
	switch val {
	case Qt__FastTransformation: // 0
		return "FastTransformation"
	case Qt__SmoothTransformation: // 1
		return "SmoothTransformation"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func AxisItemName(val int) string {
	switch val {
	case Qt__XAxis: // 0
		return "XAxis"
	case Qt__YAxis: // 1
		return "YAxis"
	case Qt__ZAxis: // 2
		return "ZAxis"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func FocusReasonItemName(val int) string {
	switch val {
	case Qt__MouseFocusReason: // 0
		return "MouseFocusReason"
	case Qt__TabFocusReason: // 1
		return "TabFocusReason"
	case Qt__BacktabFocusReason: // 2
		return "BacktabFocusReason"
	case Qt__ActiveWindowFocusReason: // 3
		return "ActiveWindowFocusReason"
	case Qt__PopupFocusReason: // 4
		return "PopupFocusReason"
	case Qt__ShortcutFocusReason: // 5
		return "ShortcutFocusReason"
	case Qt__MenuBarFocusReason: // 6
		return "MenuBarFocusReason"
	case Qt__OtherFocusReason: // 7
		return "OtherFocusReason"
	case Qt__NoFocusReason: // 8
		return "NoFocusReason"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ContextMenuPolicyItemName(val int) string {
	switch val {
	case Qt__NoContextMenu: // 0
		return "NoContextMenu"
	case Qt__DefaultContextMenu: // 1
		return "DefaultContextMenu"
	case Qt__ActionsContextMenu: // 2
		return "ActionsContextMenu"
	case Qt__CustomContextMenu: // 3
		return "CustomContextMenu"
	case Qt__PreventContextMenu: // 4
		return "PreventContextMenu"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func InputMethodQueryItemName(val int) string {
	switch val {
	case Qt__ImEnabled: // 1
		return "ImEnabled"
	case Qt__ImCursorRectangle: // 2
		return "ImCursorRectangle,ImMicroFocus"
		// case Qt__ImMicroFocus: // 2
		// return ""
	case Qt__ImFont: // 4
		return "ImFont"
	case Qt__ImCursorPosition: // 8
		return "ImCursorPosition"
	case Qt__ImSurroundingText: // 16
		return "ImSurroundingText"
	case Qt__ImCurrentSelection: // 32
		return "ImCurrentSelection"
	case Qt__ImMaximumTextLength: // 64
		return "ImMaximumTextLength"
	case Qt__ImAnchorPosition: // 128
		return "ImAnchorPosition"
	case Qt__ImHints: // 256
		return "ImHints"
	case Qt__ImPreferredLanguage: // 512
		return "ImPreferredLanguage"
	case Qt__ImAbsolutePosition: // 1024
		return "ImAbsolutePosition"
	case Qt__ImTextBeforeCursor: // 2048
		return "ImTextBeforeCursor"
	case Qt__ImTextAfterCursor: // 4096
		return "ImTextAfterCursor"
	case Qt__ImEnterKeyType: // 8192
		return "ImEnterKeyType"
	case Qt__ImAnchorRectangle: // 16384
		return "ImAnchorRectangle"
	case Qt__ImInputItemClipRectangle: // 32768
		return "ImInputItemClipRectangle"
	case Qt__ImPlatformData: // -2147483648
		return "ImPlatformData"
	case Qt__ImQueryInput: // 16570
		return "ImQueryInput"
	case Qt__ImQueryAll: // -1
		return "ImQueryAll"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func InputMethodHintItemName(val int) string {
	switch val {
	case Qt__ImhNone: // 0
		return "ImhNone"
	case Qt__ImhHiddenText: // 1
		return "ImhHiddenText"
	case Qt__ImhSensitiveData: // 2
		return "ImhSensitiveData"
	case Qt__ImhNoAutoUppercase: // 4
		return "ImhNoAutoUppercase"
	case Qt__ImhPreferNumbers: // 8
		return "ImhPreferNumbers"
	case Qt__ImhPreferUppercase: // 16
		return "ImhPreferUppercase"
	case Qt__ImhPreferLowercase: // 32
		return "ImhPreferLowercase"
	case Qt__ImhNoPredictiveText: // 64
		return "ImhNoPredictiveText"
	case Qt__ImhDate: // 128
		return "ImhDate"
	case Qt__ImhTime: // 256
		return "ImhTime"
	case Qt__ImhPreferLatin: // 512
		return "ImhPreferLatin"
	case Qt__ImhMultiLine: // 1024
		return "ImhMultiLine"
	case Qt__ImhDigitsOnly: // 65536
		return "ImhDigitsOnly"
	case Qt__ImhFormattedNumbersOnly: // 131072
		return "ImhFormattedNumbersOnly"
	case Qt__ImhUppercaseOnly: // 262144
		return "ImhUppercaseOnly"
	case Qt__ImhLowercaseOnly: // 524288
		return "ImhLowercaseOnly"
	case Qt__ImhDialableCharactersOnly: // 1048576
		return "ImhDialableCharactersOnly"
	case Qt__ImhEmailCharactersOnly: // 2097152
		return "ImhEmailCharactersOnly"
	case Qt__ImhUrlCharactersOnly: // 4194304
		return "ImhUrlCharactersOnly"
	case Qt__ImhLatinOnly: // 8388608
		return "ImhLatinOnly"
	case Qt__ImhExclusiveInputMask: // -65536
		return "ImhExclusiveInputMask"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func EnterKeyTypeItemName(val int) string {
	switch val {
	case Qt__EnterKeyDefault: // 0
		return "EnterKeyDefault"
	case Qt__EnterKeyReturn: // 1
		return "EnterKeyReturn"
	case Qt__EnterKeyDone: // 2
		return "EnterKeyDone"
	case Qt__EnterKeyGo: // 3
		return "EnterKeyGo"
	case Qt__EnterKeySend: // 4
		return "EnterKeySend"
	case Qt__EnterKeySearch: // 5
		return "EnterKeySearch"
	case Qt__EnterKeyNext: // 6
		return "EnterKeyNext"
	case Qt__EnterKeyPrevious: // 7
		return "EnterKeyPrevious"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ToolButtonStyleItemName(val int) string {
	switch val {
	case Qt__ToolButtonIconOnly: // 0
		return "ToolButtonIconOnly"
	case Qt__ToolButtonTextOnly: // 1
		return "ToolButtonTextOnly"
	case Qt__ToolButtonTextBesideIcon: // 2
		return "ToolButtonTextBesideIcon"
	case Qt__ToolButtonTextUnderIcon: // 3
		return "ToolButtonTextUnderIcon"
	case Qt__ToolButtonFollowStyle: // 4
		return "ToolButtonFollowStyle"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func LayoutDirectionItemName(val int) string {
	switch val {
	case Qt__LeftToRight: // 0
		return "LeftToRight"
	case Qt__RightToLeft: // 1
		return "RightToLeft"
	case Qt__LayoutDirectionAuto: // 2
		return "LayoutDirectionAuto"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func AnchorPointItemName(val int) string {
	switch val {
	case Qt__AnchorLeft: // 0
		return "AnchorLeft"
	case Qt__AnchorHorizontalCenter: // 1
		return "AnchorHorizontalCenter"
	case Qt__AnchorRight: // 2
		return "AnchorRight"
	case Qt__AnchorTop: // 3
		return "AnchorTop"
	case Qt__AnchorVerticalCenter: // 4
		return "AnchorVerticalCenter"
	case Qt__AnchorBottom: // 5
		return "AnchorBottom"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__FindChildOption = int // core
//
const Qt__FindDirectChildrenOnly Qt__FindChildOption = 0

//
const Qt__FindChildrenRecursively Qt__FindChildOption = 1

func FindChildOptionItemName(val int) string {
	switch val {
	case Qt__FindDirectChildrenOnly: // 0
		return "FindDirectChildrenOnly"
	case Qt__FindChildrenRecursively: // 1
		return "FindChildrenRecursively"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func DropActionItemName(val int) string {
	switch val {
	case Qt__CopyAction: // 1
		return "CopyAction"
	case Qt__MoveAction: // 2
		return "MoveAction"
	case Qt__LinkAction: // 4
		return "LinkAction"
	case Qt__ActionMask: // 255
		return "ActionMask"
	case Qt__TargetMoveAction: // 32770
		return "TargetMoveAction"
	case Qt__IgnoreAction: // 0
		return "IgnoreAction"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func CheckStateItemName(val int) string {
	switch val {
	case Qt__Unchecked: // 0
		return "Unchecked"
	case Qt__PartiallyChecked: // 1
		return "PartiallyChecked"
	case Qt__Checked: // 2
		return "Checked"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ItemDataRoleItemName(val int) string {
	switch val {
	case Qt__DisplayRole: // 0
		return "DisplayRole"
	case Qt__DecorationRole: // 1
		return "DecorationRole"
	case Qt__EditRole: // 2
		return "EditRole"
	case Qt__ToolTipRole: // 3
		return "ToolTipRole"
	case Qt__StatusTipRole: // 4
		return "StatusTipRole"
	case Qt__WhatsThisRole: // 5
		return "WhatsThisRole"
	case Qt__FontRole: // 6
		return "FontRole"
	case Qt__TextAlignmentRole: // 7
		return "TextAlignmentRole"
	case Qt__BackgroundColorRole: // 8
		return "BackgroundColorRole,BackgroundRole"
		// case Qt__BackgroundRole: // 8
		// return ""
	case Qt__TextColorRole: // 9
		return "TextColorRole,ForegroundRole"
		// case Qt__ForegroundRole: // 9
		// return ""
	case Qt__CheckStateRole: // 10
		return "CheckStateRole"
	case Qt__AccessibleTextRole: // 11
		return "AccessibleTextRole"
	case Qt__AccessibleDescriptionRole: // 12
		return "AccessibleDescriptionRole"
	case Qt__SizeHintRole: // 13
		return "SizeHintRole"
	case Qt__InitialSortOrderRole: // 14
		return "InitialSortOrderRole"
	case Qt__DisplayPropertyRole: // 27
		return "DisplayPropertyRole"
	case Qt__DecorationPropertyRole: // 28
		return "DecorationPropertyRole"
	case Qt__ToolTipPropertyRole: // 29
		return "ToolTipPropertyRole"
	case Qt__StatusTipPropertyRole: // 30
		return "StatusTipPropertyRole"
	case Qt__WhatsThisPropertyRole: // 31
		return "WhatsThisPropertyRole"
	case Qt__UserRole: // 256
		return "UserRole"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ItemFlagItemName(val int) string {
	switch val {
	case Qt__NoItemFlags: // 0
		return "NoItemFlags"
	case Qt__ItemIsSelectable: // 1
		return "ItemIsSelectable"
	case Qt__ItemIsEditable: // 2
		return "ItemIsEditable"
	case Qt__ItemIsDragEnabled: // 4
		return "ItemIsDragEnabled"
	case Qt__ItemIsDropEnabled: // 8
		return "ItemIsDropEnabled"
	case Qt__ItemIsUserCheckable: // 16
		return "ItemIsUserCheckable"
	case Qt__ItemIsEnabled: // 32
		return "ItemIsEnabled"
	case Qt__ItemIsAutoTristate: // 64
		return "ItemIsAutoTristate,ItemIsTristate"
		// case Qt__ItemIsTristate: // 64
		// return ""
	case Qt__ItemNeverHasChildren: // 128
		return "ItemNeverHasChildren"
	case Qt__ItemIsUserTristate: // 256
		return "ItemIsUserTristate"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func MatchFlagItemName(val int) string {
	switch val {
	case Qt__MatchExactly: // 0
		return "MatchExactly"
	case Qt__MatchContains: // 1
		return "MatchContains"
	case Qt__MatchStartsWith: // 2
		return "MatchStartsWith"
	case Qt__MatchEndsWith: // 3
		return "MatchEndsWith"
	case Qt__MatchRegExp: // 4
		return "MatchRegExp"
	case Qt__MatchWildcard: // 5
		return "MatchWildcard"
	case Qt__MatchFixedString: // 8
		return "MatchFixedString"
	case Qt__MatchCaseSensitive: // 16
		return "MatchCaseSensitive"
	case Qt__MatchWrap: // 32
		return "MatchWrap"
	case Qt__MatchRecursive: // 64
		return "MatchRecursive"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func WindowModalityItemName(val int) string {
	switch val {
	case Qt__NonModal: // 0
		return "NonModal"
	case Qt__WindowModal: // 1
		return "WindowModal"
	case Qt__ApplicationModal: // 2
		return "ApplicationModal"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TextInteractionFlagItemName(val int) string {
	switch val {
	case Qt__NoTextInteraction: // 0
		return "NoTextInteraction"
	case Qt__TextSelectableByMouse: // 1
		return "TextSelectableByMouse"
	case Qt__TextSelectableByKeyboard: // 2
		return "TextSelectableByKeyboard"
	case Qt__LinksAccessibleByMouse: // 4
		return "LinksAccessibleByMouse"
	case Qt__LinksAccessibleByKeyboard: // 8
		return "LinksAccessibleByKeyboard"
	case Qt__TextEditable: // 16
		return "TextEditable"
	case Qt__TextEditorInteraction: // 19
		return "TextEditorInteraction"
	case Qt__TextBrowserInteraction: // 13
		return "TextBrowserInteraction"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func EventPriorityItemName(val int) string {
	switch val {
	case Qt__HighEventPriority: // 1
		return "HighEventPriority"
	case Qt__NormalEventPriority: // 0
		return "NormalEventPriority"
	case Qt__LowEventPriority: // -1
		return "LowEventPriority"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func SizeHintItemName(val int) string {
	switch val {
	case Qt__MinimumSize: // 0
		return "MinimumSize"
	case Qt__PreferredSize: // 1
		return "PreferredSize"
	case Qt__MaximumSize: // 2
		return "MaximumSize"
	case Qt__MinimumDescent: // 3
		return "MinimumDescent"
	case Qt__NSizeHints: // 4
		return "NSizeHints"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func WindowFrameSectionItemName(val int) string {
	switch val {
	case Qt__NoSection: // 0
		return "NoSection"
	case Qt__LeftSection: // 1
		return "LeftSection"
	case Qt__TopLeftSection: // 2
		return "TopLeftSection"
	case Qt__TopSection: // 3
		return "TopSection"
	case Qt__TopRightSection: // 4
		return "TopRightSection"
	case Qt__RightSection: // 5
		return "RightSection"
	case Qt__BottomRightSection: // 6
		return "BottomRightSection"
	case Qt__BottomSection: // 7
		return "BottomSection"
	case Qt__BottomLeftSection: // 8
		return "BottomLeftSection"
	case Qt__TitleBarArea: // 9
		return "TitleBarArea"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__Initialization = int // core
//
const Qt__Uninitialized Qt__Initialization = 0

func InitializationItemName(val int) string {
	switch val {
	case Qt__Uninitialized: // 0
		return "Uninitialized"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum specifies the coordinate system.



This enum was introduced or modified in  Qt 4.6.

*/
type Qt__CoordinateSystem = int // core
// Coordinates are relative to the top-left corner of the object's paint device.
const Qt__DeviceCoordinates Qt__CoordinateSystem = 0

// Coordinates are relative to the top-left corner of the object.
const Qt__LogicalCoordinates Qt__CoordinateSystem = 1

func CoordinateSystemItemName(val int) string {
	switch val {
	case Qt__DeviceCoordinates: // 0
		return "DeviceCoordinates"
	case Qt__LogicalCoordinates: // 1
		return "LogicalCoordinates"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TouchPointStateItemName(val int) string {
	switch val {
	case Qt__TouchPointPressed: // 1
		return "TouchPointPressed"
	case Qt__TouchPointMoved: // 2
		return "TouchPointMoved"
	case Qt__TouchPointStationary: // 4
		return "TouchPointStationary"
	case Qt__TouchPointReleased: // 8
		return "TouchPointReleased"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func GestureStateItemName(val int) string {
	switch val {
	case Qt__NoGesture: // 0
		return "NoGesture"
	case Qt__GestureStarted: // 1
		return "GestureStarted"
	case Qt__GestureUpdated: // 2
		return "GestureUpdated"
	case Qt__GestureFinished: // 3
		return "GestureFinished"
	case Qt__GestureCanceled: // 4
		return "GestureCanceled"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func GestureTypeItemName(val int) string {
	switch val {
	case Qt__TapGesture: // 1
		return "TapGesture"
	case Qt__TapAndHoldGesture: // 2
		return "TapAndHoldGesture"
	case Qt__PanGesture: // 3
		return "PanGesture"
	case Qt__PinchGesture: // 4
		return "PinchGesture"
	case Qt__SwipeGesture: // 5
		return "SwipeGesture"
	case Qt__CustomGesture: // 256
		return "CustomGesture"
	case Qt__LastGestureType: // -1
		return "LastGestureType"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__GestureFlag = int // core
//
const Qt__DontStartGestureOnChildren Qt__GestureFlag = 1

//
const Qt__ReceivePartialGestures Qt__GestureFlag = 2

//
const Qt__IgnoredGesturesPropagateToParent Qt__GestureFlag = 4

func GestureFlagItemName(val int) string {
	switch val {
	case Qt__DontStartGestureOnChildren: // 1
		return "DontStartGestureOnChildren"
	case Qt__ReceivePartialGestures: // 2
		return "ReceivePartialGestures"
	case Qt__IgnoredGesturesPropagateToParent: // 4
		return "IgnoredGesturesPropagateToParent"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func NativeGestureTypeItemName(val int) string {
	switch val {
	case Qt__BeginNativeGesture: // 0
		return "BeginNativeGesture"
	case Qt__EndNativeGesture: // 1
		return "EndNativeGesture"
	case Qt__PanNativeGesture: // 2
		return "PanNativeGesture"
	case Qt__ZoomNativeGesture: // 3
		return "ZoomNativeGesture"
	case Qt__SmartZoomNativeGesture: // 4
		return "SmartZoomNativeGesture"
	case Qt__RotateNativeGesture: // 5
		return "RotateNativeGesture"
	case Qt__SwipeNativeGesture: // 6
		return "SwipeNativeGesture"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func NavigationModeItemName(val int) string {
	switch val {
	case Qt__NavigationModeNone: // 0
		return "NavigationModeNone"
	case Qt__NavigationModeKeypadTabOrder: // 1
		return "NavigationModeKeypadTabOrder"
	case Qt__NavigationModeKeypadDirectional: // 2
		return "NavigationModeKeypadDirectional"
	case Qt__NavigationModeCursorAuto: // 3
		return "NavigationModeCursorAuto"
	case Qt__NavigationModeCursorForceVisible: // 4
		return "NavigationModeCursorForceVisible"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum describes the movement style available to text cursors. The options are:


*/
type Qt__CursorMoveStyle = int // core
// Within a left-to-right text block, decrease cursor position when pressing left arrow key, increase cursor position when pressing the right arrow key. If the text block is right-to-left, the opposite behavior applies.
const Qt__LogicalMoveStyle Qt__CursorMoveStyle = 0

// Pressing the left arrow key will always cause the cursor to move left, regardless of the text's writing direction. Pressing the right arrow key will always cause the cursor to move right.
const Qt__VisualMoveStyle Qt__CursorMoveStyle = 1

func CursorMoveStyleItemName(val int) string {
	switch val {
	case Qt__LogicalMoveStyle: // 0
		return "LogicalMoveStyle"
	case Qt__VisualMoveStyle: // 1
		return "VisualMoveStyle"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func TimerTypeItemName(val int) string {
	switch val {
	case Qt__PreciseTimer: // 0
		return "PreciseTimer"
	case Qt__CoarseTimer: // 1
		return "CoarseTimer"
	case Qt__VeryCoarseTimer: // 2
		return "VeryCoarseTimer"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ScrollPhaseItemName(val int) string {
	switch val {
	case Qt__NoScrollPhase: // 0
		return "NoScrollPhase"
	case Qt__ScrollBegin: // 1
		return "ScrollBegin"
	case Qt__ScrollUpdate: // 2
		return "ScrollUpdate"
	case Qt__ScrollEnd: // 3
		return "ScrollEnd"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func MouseEventSourceItemName(val int) string {
	switch val {
	case Qt__MouseEventNotSynthesized: // 0
		return "MouseEventNotSynthesized"
	case Qt__MouseEventSynthesizedBySystem: // 1
		return "MouseEventSynthesizedBySystem"
	case Qt__MouseEventSynthesizedByQt: // 2
		return "MouseEventSynthesizedByQt"
	case Qt__MouseEventSynthesizedByApplication: // 3
		return "MouseEventSynthesizedByApplication"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__MouseEventFlag = int // core
//
const Qt__MouseEventCreatedDoubleClick Qt__MouseEventFlag = 1

//
const Qt__MouseEventFlagMask Qt__MouseEventFlag = 255

func MouseEventFlagItemName(val int) string {
	switch val {
	case Qt__MouseEventCreatedDoubleClick: // 1
		return "MouseEventCreatedDoubleClick"
	case Qt__MouseEventFlagMask: // 255
		return "MouseEventFlagMask"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum describes the possible standards used by qChecksum().



This enum was introduced or modified in  Qt 5.9.

*/
type Qt__ChecksumType = int // core
//
const Qt__ChecksumIso3309 Qt__ChecksumType = 0

//
const Qt__ChecksumItuV41 Qt__ChecksumType = 1

func ChecksumTypeItemName(val int) string {
	switch val {
	case Qt__ChecksumIso3309: // 0
		return "ChecksumIso3309"
	case Qt__ChecksumItuV41: // 1
		return "ChecksumItuV41"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func errcItemName(val int) string {
	switch val {
	case Qt__address_family_not_supported: // 97
		return "address_family_not_supported"
	case Qt__address_in_use: // 98
		return "address_in_use"
	case Qt__address_not_available: // 99
		return "address_not_available"
	case Qt__already_connected: // 106
		return "already_connected"
	case Qt__argument_list_too_long: // 7
		return "argument_list_too_long"
	case Qt__argument_out_of_domain: // 33
		return "argument_out_of_domain"
	case Qt__bad_address: // 14
		return "bad_address"
	case Qt__bad_file_descriptor: // 9
		return "bad_file_descriptor"
	case Qt__bad_message: // 74
		return "bad_message"
	case Qt__broken_pipe: // 32
		return "broken_pipe"
	case Qt__connection_aborted: // 103
		return "connection_aborted"
	case Qt__connection_already_in_progress: // 114
		return "connection_already_in_progress"
	case Qt__connection_refused: // 111
		return "connection_refused"
	case Qt__connection_reset: // 104
		return "connection_reset"
	case Qt__cross_device_link: // 18
		return "cross_device_link"
	case Qt__destination_address_required: // 89
		return "destination_address_required"
	case Qt__device_or_resource_busy: // 16
		return "device_or_resource_busy"
	case Qt__directory_not_empty: // 39
		return "directory_not_empty"
	case Qt__executable_format_error: // 8
		return "executable_format_error"
	case Qt__file_exists: // 17
		return "file_exists"
	case Qt__file_too_large: // 27
		return "file_too_large"
	case Qt__filename_too_long: // 36
		return "filename_too_long"
	case Qt__function_not_supported: // 38
		return "function_not_supported"
	case Qt__host_unreachable: // 113
		return "host_unreachable"
	case Qt__identifier_removed: // 43
		return "identifier_removed"
	case Qt__illegal_byte_sequence: // 84
		return "illegal_byte_sequence"
	case Qt__inappropriate_io_control_operation: // 25
		return "inappropriate_io_control_operation"
	case Qt__interrupted: // 4
		return "interrupted"
	case Qt__invalid_argument: // 22
		return "invalid_argument"
	case Qt__invalid_seek: // 29
		return "invalid_seek"
	case Qt__io_error: // 5
		return "io_error"
	case Qt__is_a_directory: // 21
		return "is_a_directory"
	case Qt__message_size: // 90
		return "message_size"
	case Qt__network_down: // 100
		return "network_down"
	case Qt__network_reset: // 102
		return "network_reset"
	case Qt__network_unreachable: // 101
		return "network_unreachable"
	case Qt__no_buffer_space: // 105
		return "no_buffer_space"
	case Qt__no_child_process: // 10
		return "no_child_process"
	case Qt__no_link: // 67
		return "no_link"
	case Qt__no_lock_available: // 37
		return "no_lock_available"
	case Qt__no_message_available: // 61
		return "no_message_available"
	case Qt__no_message: // 42
		return "no_message"
	case Qt__no_protocol_option: // 92
		return "no_protocol_option"
	case Qt__no_space_on_device: // 28
		return "no_space_on_device"
	case Qt__no_stream_resources: // 63
		return "no_stream_resources"
	case Qt__no_such_device_or_address: // 6
		return "no_such_device_or_address"
	case Qt__no_such_device: // 19
		return "no_such_device"
	case Qt__no_such_file_or_directory: // 2
		return "no_such_file_or_directory"
	case Qt__no_such_process: // 3
		return "no_such_process"
	case Qt__not_a_directory: // 20
		return "not_a_directory"
	case Qt__not_a_socket: // 88
		return "not_a_socket"
	case Qt__not_a_stream: // 60
		return "not_a_stream"
	case Qt__not_connected: // 107
		return "not_connected"
	case Qt__not_enough_memory: // 12
		return "not_enough_memory"
	case Qt__not_supported: // 95
		return "not_supported,operation_not_supported"
	case Qt__operation_canceled: // 125
		return "operation_canceled"
	case Qt__operation_in_progress: // 115
		return "operation_in_progress"
	case Qt__operation_not_permitted: // 1
		return "operation_not_permitted"
		// case Qt__operation_not_supported: // 95
		// return ""
	case Qt__operation_would_block: // 11
		return "operation_would_block,resource_unavailable_try_again"
	case Qt__owner_dead: // 130
		return "owner_dead"
	case Qt__permission_denied: // 13
		return "permission_denied"
	case Qt__protocol_error: // 71
		return "protocol_error"
	case Qt__protocol_not_supported: // 93
		return "protocol_not_supported"
	case Qt__read_only_file_system: // 30
		return "read_only_file_system"
	case Qt__resource_deadlock_would_occur: // 35
		return "resource_deadlock_would_occur"
		// case Qt__resource_unavailable_try_again: // 11
		// return ""
	case Qt__result_out_of_range: // 34
		return "result_out_of_range"
	case Qt__state_not_recoverable: // 131
		return "state_not_recoverable"
	case Qt__stream_timeout: // 62
		return "stream_timeout"
	case Qt__text_file_busy: // 26
		return "text_file_busy"
	case Qt__timed_out: // 110
		return "timed_out"
	case Qt__too_many_files_open_in_system: // 23
		return "too_many_files_open_in_system"
	case Qt__too_many_files_open: // 24
		return "too_many_files_open"
	case Qt__too_many_links: // 31
		return "too_many_links"
	case Qt__too_many_symbolic_link_levels: // 40
		return "too_many_symbolic_link_levels"
	case Qt__value_too_large: // 75
		return "value_too_large"
	case Qt__wrong_protocol_type: // 91
		return "wrong_protocol_type"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _Ios_FmtflagsItemName(val int) string {
	switch val {
	case Qt___S_boolalpha: // 1
		return "_S_boolalpha"
	case Qt___S_dec: // 2
		return "_S_dec"
	case Qt___S_fixed: // 4
		return "_S_fixed"
	case Qt___S_hex: // 8
		return "_S_hex"
	case Qt___S_internal: // 16
		return "_S_internal"
	case Qt___S_left: // 32
		return "_S_left"
	case Qt___S_oct: // 64
		return "_S_oct"
	case Qt___S_right: // 128
		return "_S_right"
	case Qt___S_scientific: // 256
		return "_S_scientific"
	case Qt___S_showbase: // 512
		return "_S_showbase"
	case Qt___S_showpoint: // 1024
		return "_S_showpoint"
	case Qt___S_showpos: // 2048
		return "_S_showpos"
	case Qt___S_skipws: // 4096
		return "_S_skipws"
	case Qt___S_unitbuf: // 8192
		return "_S_unitbuf"
	case Qt___S_uppercase: // 16384
		return "_S_uppercase"
	case Qt___S_adjustfield: // 176
		return "_S_adjustfield"
	case Qt___S_basefield: // 74
		return "_S_basefield"
	case Qt___S_floatfield: // 260
		return "_S_floatfield"
	case Qt___S_ios_fmtflags_end: // 65536
		return "_S_ios_fmtflags_end"
	case Qt___S_ios_fmtflags_max: // 2147483647
		return "_S_ios_fmtflags_max"
	case Qt___S_ios_fmtflags_min: // -2147483648
		return "_S_ios_fmtflags_min"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _Ios_OpenmodeItemName(val int) string {
	switch val {
	case Qt___S_app: // 1
		return "_S_app"
	case Qt___S_ate: // 2
		return "_S_ate"
	case Qt___S_bin: // 4
		return "_S_bin"
	case Qt___S_in: // 8
		return "_S_in"
	case Qt___S_out: // 16
		return "_S_out"
	case Qt___S_trunc: // 32
		return "_S_trunc"
	case Qt___S_ios_openmode_end: // 65536
		return "_S_ios_openmode_end"
	case Qt___S_ios_openmode_max: // 2147483647
		return "_S_ios_openmode_max"
	case Qt___S_ios_openmode_min: // -2147483648
		return "_S_ios_openmode_min"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _Ios_IostateItemName(val int) string {
	switch val {
	case Qt___S_goodbit: // 0
		return "_S_goodbit"
	case Qt___S_badbit: // 1
		return "_S_badbit"
	case Qt___S_eofbit: // 2
		return "_S_eofbit"
	case Qt___S_failbit: // 4
		return "_S_failbit"
	case Qt___S_ios_iostate_end: // 65536
		return "_S_ios_iostate_end"
	case Qt___S_ios_iostate_max: // 2147483647
		return "_S_ios_iostate_max"
	case Qt___S_ios_iostate_min: // -2147483648
		return "_S_ios_iostate_min"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _Ios_SeekdirItemName(val int) string {
	switch val {
	case Qt___S_beg: // 0
		return "_S_beg"
	case Qt___S_cur: // 1
		return "_S_cur"
	case Qt___S_end: // 2
		return "_S_end"
	case Qt___S_ios_seekdir_end: // 65536
		return "_S_ios_seekdir_end"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__io_errc = int // stdglobal
//
const Qt__stream Qt__io_errc = 1

func io_errcItemName(val int) string {
	switch val {
	case Qt__stream: // 1
		return "stream"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___Rb_tree_color = int // stdglobal
//
const Qt___S_red Qt___Rb_tree_color = 0

//
const Qt___S_black Qt___Rb_tree_color = 1

func _Rb_tree_colorItemName(val int) string {
	switch val {
	case Qt___S_red: // 0
		return "_S_red"
	case Qt___S_black: // 1
		return "_S_black"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__IteratorCapability = int // core
//
const Qt__ForwardCapability Qt__IteratorCapability = 1

//
const Qt__BiDirectionalCapability Qt__IteratorCapability = 2

//
const Qt__RandomAccessCapability Qt__IteratorCapability = 4

func IteratorCapabilityItemName(val int) string {
	switch val {
	case Qt__ForwardCapability: // 1
		return "ForwardCapability"
	case Qt__BiDirectionalCapability: // 2
		return "BiDirectionalCapability"
	case Qt__RandomAccessCapability: // 4
		return "RandomAccessCapability"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _Manager_operationItemName(val int) string {
	switch val {
	case Qt____get_type_info: // 0
		return "__get_type_info"
	case Qt____get_functor_ptr: // 1
		return "__get_functor_ptr"
	case Qt____clone_functor: // 2
		return "__clone_functor"
	case Qt____destroy_functor: // 3
		return "__destroy_functor"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___Lock_policy = int // stdglobal
//
const Qt___S_single Qt___Lock_policy = 0

//
const Qt___S_mutex Qt___Lock_policy = 1

//
const Qt___S_atomic Qt___Lock_policy = 2

func _Lock_policyItemName(val int) string {
	switch val {
	case Qt___S_single: // 0
		return "_S_single"
	case Qt___S_mutex: // 1
		return "_S_mutex"
	case Qt___S_atomic: // 2
		return "_S_atomic"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__pointer_safety = int // stdglobal
//
const Qt__relaxed Qt__pointer_safety = 0

//
const Qt__preferred Qt__pointer_safety = 1

//
const Qt__strict Qt__pointer_safety = 2

func pointer_safetyItemName(val int) string {
	switch val {
	case Qt__relaxed: // 0
		return "relaxed"
	case Qt__preferred: // 1
		return "preferred"
	case Qt__strict: // 2
		return "strict"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func future_errcItemName(val int) string {
	switch val {
	case Qt__future_already_retrieved: // 1
		return "future_already_retrieved"
	case Qt__promise_already_satisfied: // 2
		return "promise_already_satisfied"
	case Qt__no_state: // 3
		return "no_state"
	case Qt__broken_promise: // 4
		return "broken_promise"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__future_status = int // stdglobal
//
const Qt__ready Qt__future_status = 0

//
const Qt__timeout Qt__future_status = 1

//
const Qt__deferred Qt__future_status = 2

func future_statusItemName(val int) string {
	switch val {
	case Qt__ready: // 0
		return "ready"
	case Qt__timeout: // 1
		return "timeout"
	case Qt__deferred: // 2
		return "deferred"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func DrawingHintItemName(val int) string {
	switch val {
	case Qt__OpaqueTopLeft: // 1
		return "OpaqueTopLeft"
	case Qt__OpaqueTop: // 2
		return "OpaqueTop"
	case Qt__OpaqueTopRight: // 4
		return "OpaqueTopRight"
	case Qt__OpaqueLeft: // 8
		return "OpaqueLeft"
	case Qt__OpaqueCenter: // 16
		return "OpaqueCenter"
	case Qt__OpaqueRight: // 32
		return "OpaqueRight"
	case Qt__OpaqueBottomLeft: // 64
		return "OpaqueBottomLeft"
	case Qt__OpaqueBottom: // 128
		return "OpaqueBottom"
	case Qt__OpaqueBottomRight: // 256
		return "OpaqueBottomRight"
	case Qt__OpaqueCorners: // 325
		return "OpaqueCorners"
	case Qt__OpaqueEdges: // 170
		return "OpaqueEdges"
	case Qt__OpaqueFrame: // 495
		return "OpaqueFrame"
	case Qt__OpaqueAll: // 511
		return "OpaqueAll"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
Describes the two types of keys QSslKey supports.


*/
type Qt__KeyType = int // network
// A private key.
const Qt__PrivateKey Qt__KeyType = 0

// A public key.
const Qt__PublicKey Qt__KeyType = 1

func KeyTypeItemName(val int) string {
	switch val {
	case Qt__PrivateKey: // 0
		return "PrivateKey"
	case Qt__PublicKey: // 1
		return "PublicKey"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
Describes supported encoding formats for certificates and keys.


*/
type Qt__EncodingFormat = int // network
// The PEM format.
const Qt__Pem Qt__EncodingFormat = 0

// The DER format.
const Qt__Der Qt__EncodingFormat = 1

func EncodingFormatItemName(val int) string {
	switch val {
	case Qt__Pem: // 0
		return "Pem"
	case Qt__Der: // 1
		return "Der"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func KeyAlgorithmItemName(val int) string {
	switch val {
	case Qt__Opaque: // 0
		return "Opaque"
	case Qt__Rsa: // 1
		return "Rsa"
	case Qt__Dsa: // 2
		return "Dsa"
	case Qt__Ec: // 3
		return "Ec"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func AlternativeNameEntryTypeItemName(val int) string {
	switch val {
	case Qt__EmailEntry: // 0
		return "EmailEntry"
	case Qt__DnsEntry: // 1
		return "DnsEntry"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func SslProtocolItemName(val int) string {
	switch val {
	case Qt__SslV3: // 0
		return "SslV3"
	case Qt__SslV2: // 1
		return "SslV2"
	case Qt__TlsV1_0: // 2
		return "TlsV1_0"
	case Qt__TlsV1_1: // 3
		return "TlsV1_1"
	case Qt__TlsV1_2: // 4
		return "TlsV1_2"
	case Qt__AnyProtocol: // 5
		return "AnyProtocol"
	case Qt__TlsV1SslV3: // 6
		return "TlsV1SslV3"
	case Qt__SecureProtocols: // 7
		return "SecureProtocols"
	case Qt__TlsV1_0OrLater: // 8
		return "TlsV1_0OrLater"
	case Qt__TlsV1_1OrLater: // 9
		return "TlsV1_1OrLater"
	case Qt__TlsV1_2OrLater: // 10
		return "TlsV1_2OrLater"
	case Qt__UnknownProtocol: // -1
		return "UnknownProtocol"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func SslOptionItemName(val int) string {
	switch val {
	case Qt__SslOptionDisableEmptyFragments: // 1
		return "SslOptionDisableEmptyFragments"
	case Qt__SslOptionDisableSessionTickets: // 2
		return "SslOptionDisableSessionTickets"
	case Qt__SslOptionDisableCompression: // 4
		return "SslOptionDisableCompression"
	case Qt__SslOptionDisableServerNameIndication: // 8
		return "SslOptionDisableServerNameIndication"
	case Qt__SslOptionDisableLegacyRenegotiation: // 16
		return "SslOptionDisableLegacyRenegotiation"
	case Qt__SslOptionDisableSessionSharing: // 32
		return "SslOptionDisableSessionSharing"
	case Qt__SslOptionDisableSessionPersistence: // 64
		return "SslOptionDisableSessionPersistence"
	case Qt__SslOptionDisableServerCipherPreference: // 128
		return "SslOptionDisableServerCipherPreference"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__AutoParentResult = int // qml
//
const Qt__Parented Qt__AutoParentResult = 0

//
const Qt__IncompatibleObject Qt__AutoParentResult = 1

//
const Qt__IncompatibleParent Qt__AutoParentResult = 2

func AutoParentResultItemName(val int) string {
	switch val {
	case Qt__Parented: // 0
		return "Parented"
	case Qt__IncompatibleObject: // 1
		return "IncompatibleObject"
	case Qt__IncompatibleParent: // 2
		return "IncompatibleParent"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func RegistrationTypeItemName(val int) string {
	switch val {
	case Qt__TypeRegistration: // 0
		return "TypeRegistration"
	case Qt__InterfaceRegistration: // 1
		return "InterfaceRegistration"
	case Qt__AutoParentRegistration: // 2
		return "AutoParentRegistration"
	case Qt__SingletonRegistration: // 3
		return "SingletonRegistration"
	case Qt__CompositeRegistration: // 4
		return "CompositeRegistration"
	case Qt__CompositeSingletonRegistration: // 5
		return "CompositeSingletonRegistration"
	case Qt__QmlUnitCacheHookRegistration: // 6
		return "QmlUnitCacheHookRegistration"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func SupportEstimateItemName(val int) string {
	switch val {
	case Qt__NotSupported: // 0
		return "NotSupported"
	case Qt__MaybeSupported: // 1
		return "MaybeSupported"
	case Qt__ProbablySupported: // 2
		return "ProbablySupported"
	case Qt__PreferredService: // 3
		return "PreferredService"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func EncodingQualityItemName(val int) string {
	switch val {
	case Qt__VeryLowQuality: // 0
		return "VeryLowQuality"
	case Qt__LowQuality: // 1
		return "LowQuality"
	case Qt__NormalQuality: // 2
		return "NormalQuality"
	case Qt__HighQuality: // 3
		return "HighQuality"
	case Qt__VeryHighQuality: // 4
		return "VeryHighQuality"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func EncodingModeItemName(val int) string {
	switch val {
	case Qt__ConstantQualityEncoding: // 0
		return "ConstantQualityEncoding"
	case Qt__ConstantBitRateEncoding: // 1
		return "ConstantBitRateEncoding"
	case Qt__AverageBitRateEncoding: // 2
		return "AverageBitRateEncoding"
	case Qt__TwoPassEncoding: // 3
		return "TwoPassEncoding"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func AvailabilityStatusItemName(val int) string {
	switch val {
	case Qt__Available: // 0
		return "Available"
	case Qt__ServiceMissing: // 1
		return "ServiceMissing"
	case Qt__Busy: // 2
		return "Busy"
	case Qt__ResourceError: // 3
		return "ResourceError"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func ErrorItemName(val int) string {
	switch val {
	case Qt__NoError: // 0
		return "NoError"
	case Qt__OpenError: // 1
		return "OpenError"
	case Qt__IOError: // 2
		return "IOError"
	case Qt__UnderrunError: // 3
		return "UnderrunError"
	case Qt__FatalError: // 4
		return "FatalError"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func StateItemName(val int) string {
	switch val {
	case Qt__ActiveState: // 0
		return "ActiveState"
	case Qt__SuspendedState: // 1
		return "SuspendedState"
	case Qt__StoppedState: // 2
		return "StoppedState"
	case Qt__IdleState: // 3
		return "IdleState"
	case Qt__InterruptedState: // 4
		return "InterruptedState"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*

 */
type Qt__Mode = int // multimedia
// audio input device
const Qt__AudioInput Qt__Mode = 0

// audio output device
const Qt__AudioOutput Qt__Mode = 1

func ModeItemName(val int) string {
	switch val {
	case Qt__AudioInput: // 0
		return "AudioInput"
	case Qt__AudioOutput: // 1
		return "AudioOutput"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func RoleItemName(val int) string {
	switch val {
	case Qt__UnknownRole: // 0
		return "UnknownRole"
	case Qt__MusicRole: // 1
		return "MusicRole"
	case Qt__VideoRole: // 2
		return "VideoRole"
	case Qt__VoiceCommunicationRole: // 3
		return "VoiceCommunicationRole"
	case Qt__AlarmRole: // 4
		return "AlarmRole"
	case Qt__NotificationRole: // 5
		return "NotificationRole"
	case Qt__RingtoneRole: // 6
		return "RingtoneRole"
	case Qt__AccessibilityRole: // 7
		return "AccessibilityRole"
	case Qt__SonificationRole: // 8
		return "SonificationRole"
	case Qt__GameRole: // 9
		return "GameRole"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func VolumeScaleItemName(val int) string {
	switch val {
	case Qt__LinearVolumeScale: // 0
		return "LinearVolumeScale"
	case Qt__CubicVolumeScale: // 1
		return "CubicVolumeScale"
	case Qt__LogarithmicVolumeScale: // 2
		return "LogarithmicVolumeScale"
	case Qt__DecibelVolumeScale: // 3
		return "DecibelVolumeScale"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _EXCEPTION_DISPOSITIONItemName(val int) string {
	switch val {
	case Qt__ExceptionContinueExecution: // 0
		return "ExceptionContinueExecution"
	case Qt__ExceptionContinueSearch: // 1
		return "ExceptionContinueSearch"
	case Qt__ExceptionNestedException: // 2
		return "ExceptionNestedException"
	case Qt__ExceptionCollidedUnwind: // 3
		return "ExceptionCollidedUnwind"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func tagExtendedErrorParamTypesItemName(val int) string {
	switch val {
	case Qt__eeptAnsiString: // 1
		return "eeptAnsiString"
	case Qt__eeptUnicodeString: // 2
		return "eeptUnicodeString"
	case Qt__eeptLongVal: // 3
		return "eeptLongVal"
	case Qt__eeptShortVal: // 4
		return "eeptShortVal"
	case Qt__eeptPointerVal: // 5
		return "eeptPointerVal"
	case Qt__eeptNone: // 6
		return "eeptNone"
	case Qt__eeptBinary: // 7
		return "eeptBinary"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _RPC_NOTIFICATION_TYPESItemName(val int) string {
	switch val {
	case Qt__RpcNotificationTypeNone: // 0
		return "RpcNotificationTypeNone"
	case Qt__RpcNotificationTypeEvent: // 1
		return "RpcNotificationTypeEvent"
	case Qt__RpcNotificationTypeApc: // 2
		return "RpcNotificationTypeApc"
	case Qt__RpcNotificationTypeIoc: // 3
		return "RpcNotificationTypeIoc"
	case Qt__RpcNotificationTypeHwnd: // 4
		return "RpcNotificationTypeHwnd"
	case Qt__RpcNotificationTypeCallback: // 5
		return "RpcNotificationTypeCallback"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _RPC_ASYNC_EVENTItemName(val int) string {
	switch val {
	case Qt__RpcCallComplete: // 0
		return "RpcCallComplete"
	case Qt__RpcSendComplete: // 1
		return "RpcSendComplete"
	case Qt__RpcReceiveComplete: // 2
		return "RpcReceiveComplete"
	case Qt__RpcClientDisconnect: // 3
		return "RpcClientDisconnect"
	case Qt__RpcClientCancel: // 4
		return "RpcClientCancel"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func _MEDIA_TYPEItemName(val int) string {
	switch val {
	case Qt__Unknown: // 0
		return "Unknown"
	case Qt__F5_1Pt2_512: // 1
		return "F5_1Pt2_512"
	case Qt__F3_1Pt44_512: // 2
		return "F3_1Pt44_512"
	case Qt__F3_2Pt88_512: // 3
		return "F3_2Pt88_512"
	case Qt__F3_20Pt8_512: // 4
		return "F3_20Pt8_512"
	case Qt__F3_720_512: // 5
		return "F3_720_512"
	case Qt__F5_360_512: // 6
		return "F5_360_512"
	case Qt__F5_320_512: // 7
		return "F5_320_512"
	case Qt__F5_320_1024: // 8
		return "F5_320_1024"
	case Qt__F5_180_512: // 9
		return "F5_180_512"
	case Qt__F5_160_512: // 10
		return "F5_160_512"
	case Qt__RemovableMedia: // 11
		return "RemovableMedia"
	case Qt__FixedMedia: // 12
		return "FixedMedia"
	case Qt__F3_120M_512: // 13
		return "F3_120M_512"
	case Qt__F3_640_512: // 14
		return "F3_640_512"
	case Qt__F5_640_512: // 15
		return "F5_640_512"
	case Qt__F5_720_512: // 16
		return "F5_720_512"
	case Qt__F3_1Pt2_512: // 17
		return "F3_1Pt2_512"
	case Qt__F3_1Pt23_1024: // 18
		return "F3_1Pt23_1024"
	case Qt__F5_1Pt23_1024: // 19
		return "F5_1Pt23_1024"
	case Qt__F3_128Mb_512: // 20
		return "F3_128Mb_512"
	case Qt__F3_230Mb_512: // 21
		return "F3_230Mb_512"
	case Qt__F8_256_128: // 22
		return "F8_256_128"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__HBitmapFormat = int // winextras
//
const Qt__HBitmapNoAlpha Qt__HBitmapFormat = 0

//
const Qt__HBitmapPremultipliedAlpha Qt__HBitmapFormat = 1

//
const Qt__HBitmapAlpha Qt__HBitmapFormat = 2

func HBitmapFormatItemName(val int) string {
	switch val {
	case Qt__HBitmapNoAlpha: // 0
		return "HBitmapNoAlpha"
	case Qt__HBitmapPremultipliedAlpha: // 1
		return "HBitmapPremultipliedAlpha"
	case Qt__HBitmapAlpha: // 2
		return "HBitmapAlpha"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__WindowFlip3DPolicy = int // winextras
//
const Qt__FlipDefault Qt__WindowFlip3DPolicy = 0

//
const Qt__FlipExcludeBelow Qt__WindowFlip3DPolicy = 1

//
const Qt__FlipExcludeAbove Qt__WindowFlip3DPolicy = 2

func WindowFlip3DPolicyItemName(val int) string {
	switch val {
	case Qt__FlipDefault: // 0
		return "FlipDefault"
	case Qt__FlipExcludeBelow: // 1
		return "FlipExcludeBelow"
	case Qt__FlipExcludeAbove: // 2
		return "FlipExcludeAbove"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func jobjectRefTypeItemName(val int) string {
	switch val {
	case Qt__JNIInvalidRefType: // 0
		return "JNIInvalidRefType"
	case Qt__JNILocalRefType: // 1
		return "JNILocalRefType"
	case Qt__JNIGlobalRefType: // 2
		return "JNIGlobalRefType"
	case Qt__JNIWeakGlobalRefType: // 3
		return "JNIWeakGlobalRefType"
	default:
		return fmt.Sprintf("%d", val)
	}
}

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

func BindFlagItemName(val int) string {
	switch val {
	case Qt__None: // 0
		return "None"
	case Qt__AutoCreate: // 1
		return "AutoCreate"
	case Qt__DebugUnbind: // 2
		return "DebugUnbind"
	case Qt__NotForeground: // 4
		return "NotForeground"
	case Qt__AboveClient: // 8
		return "AboveClient"
	case Qt__AllowOomManagement: // 16
		return "AllowOomManagement"
	case Qt__WaivePriority: // 32
		return "WaivePriority"
	case Qt__Important: // 64
		return "Important"
	case Qt__AdjustWithActivity: // 128
		return "AdjustWithActivity"
	case Qt__ExternalService: // -2147483648
		return "ExternalService"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__PermissionResult = int // androidextras
//
const Qt__Granted Qt__PermissionResult = 0

//
const Qt__Denied Qt__PermissionResult = 1

func PermissionResultItemName(val int) string {
	switch val {
	case Qt__Granted: // 0
		return "Granted"
	case Qt__Denied: // 1
		return "Denied"
	default:
		return fmt.Sprintf("%d", val)
	}
}

//  body block end

//  keep block begin

func make_sure_usepkg_qnamespace() {
	if false {
		fmt.Println(123)
	}
}

//  keep block end
