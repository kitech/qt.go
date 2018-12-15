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
const Qt__WA_AttributeCount Qt__WidgetAttribute = 130

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
	case Qt__WA_AttributeCount: // 130
		return "WA_AttributeCount"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*
This enum describes attributes that change the behavior of application-wide features. These are enabled and disabled using QCoreApplication::setAttribute(), and can be tested for with QCoreApplication::testAttribute().

Qt::AA_MacPluginApplicationAA_PluginApplicationThis attribute has been deprecated. Use AA_PluginApplication instead.


The following values are obsolete:


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
const Qt__AA_AttributeCount Qt__ApplicationAttribute = 28

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
	case Qt__AA_AttributeCount: // 28
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
// Coordinates are relative to the upper-left corner of the object's paint device.
const Qt__DeviceCoordinates Qt__CoordinateSystem = 0

// Coordinates are relative to the upper-left corner of the object.
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

// The audio device is in a suspended state, this state will only be entered after suspend() is called.
const Qt__SuspendedState Qt__State = 1

// The audio device is closed, and is not processing any audio data
const Qt__StoppedState Qt__State = 2

// The QIODevice passed in has no data and audio system's buffer is empty, this state is set after start() is called and while no audio data is available to be processed.
const Qt__IdleState Qt__State = 3

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
type Qt___HEAP_INFORMATION_CLASS = int // stdglobal
//
const Qt__HeapCompatibilityInformation Qt___HEAP_INFORMATION_CLASS = 0

func _HEAP_INFORMATION_CLASSItemName(val int) string {
	switch val {
	case Qt__HeapCompatibilityInformation: // 0
		return "HeapCompatibilityInformation"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__IMPORT_OBJECT_TYPE = int // stdglobal
//
const Qt__IMPORT_OBJECT_CODE Qt__IMPORT_OBJECT_TYPE = 0

//
const Qt__IMPORT_OBJECT_DATA Qt__IMPORT_OBJECT_TYPE = 1

//
const Qt__IMPORT_OBJECT_CONST Qt__IMPORT_OBJECT_TYPE = 2

func IMPORT_OBJECT_TYPEItemName(val int) string {
	switch val {
	case Qt__IMPORT_OBJECT_CODE: // 0
		return "IMPORT_OBJECT_CODE"
	case Qt__IMPORT_OBJECT_DATA: // 1
		return "IMPORT_OBJECT_DATA"
	case Qt__IMPORT_OBJECT_CONST: // 2
		return "IMPORT_OBJECT_CONST"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__IMPORT_OBJECT_NAME_TYPE = int // stdglobal
//
const Qt__IMPORT_OBJECT_ORDINAL Qt__IMPORT_OBJECT_NAME_TYPE = 0

//
const Qt__IMPORT_OBJECT_NAME Qt__IMPORT_OBJECT_NAME_TYPE = 1

//
const Qt__IMPORT_OBJECT_NAME_NO_PREFIX Qt__IMPORT_OBJECT_NAME_TYPE = 2

//
const Qt__IMPORT_OBJECT_NAME_UNDECORATE Qt__IMPORT_OBJECT_NAME_TYPE = 3

func IMPORT_OBJECT_NAME_TYPEItemName(val int) string {
	switch val {
	case Qt__IMPORT_OBJECT_ORDINAL: // 0
		return "IMPORT_OBJECT_ORDINAL"
	case Qt__IMPORT_OBJECT_NAME: // 1
		return "IMPORT_OBJECT_NAME"
	case Qt__IMPORT_OBJECT_NAME_NO_PREFIX: // 2
		return "IMPORT_OBJECT_NAME_NO_PREFIX"
	case Qt__IMPORT_OBJECT_NAME_UNDECORATE: // 3
		return "IMPORT_OBJECT_NAME_UNDECORATE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__ReplacesCorHdrNumericDefines = int // stdglobal
//
const Qt__COMIMAGE_FLAGS_ILONLY Qt__ReplacesCorHdrNumericDefines = 1

//
const Qt__COMIMAGE_FLAGS_32BITREQUIRED Qt__ReplacesCorHdrNumericDefines = 2

//
const Qt__COMIMAGE_FLAGS_IL_LIBRARY Qt__ReplacesCorHdrNumericDefines = 4

//
const Qt__COMIMAGE_FLAGS_STRONGNAMESIGNED Qt__ReplacesCorHdrNumericDefines = 8

//
const Qt__COMIMAGE_FLAGS_NATIVE_ENTRYPOINT Qt__ReplacesCorHdrNumericDefines = 16

//
const Qt__COMIMAGE_FLAGS_TRACKDEBUGDATA Qt__ReplacesCorHdrNumericDefines = 65536

//
const Qt__COMIMAGE_FLAGS_32BITPREFERRED Qt__ReplacesCorHdrNumericDefines = 131072

//
const Qt__COR_VERSION_MAJOR_V2 Qt__ReplacesCorHdrNumericDefines = 2

//
const Qt__COR_VERSION_MAJOR Qt__ReplacesCorHdrNumericDefines = 2

//
const Qt__COR_VERSION_MINOR Qt__ReplacesCorHdrNumericDefines = 5

//
const Qt__COR_DELETED_NAME_LENGTH Qt__ReplacesCorHdrNumericDefines = 8

//
const Qt__COR_VTABLEGAP_NAME_LENGTH Qt__ReplacesCorHdrNumericDefines = 8

//
const Qt__NATIVE_TYPE_MAX_CB Qt__ReplacesCorHdrNumericDefines = 1

//
const Qt__COR_ILMETHOD_SECT_SMALL_MAX_DATASIZE Qt__ReplacesCorHdrNumericDefines = 255

//
const Qt__IMAGE_COR_MIH_METHODRVA Qt__ReplacesCorHdrNumericDefines = 1

//
const Qt__IMAGE_COR_MIH_EHRVA Qt__ReplacesCorHdrNumericDefines = 2

//
const Qt__IMAGE_COR_MIH_BASICBLOCK Qt__ReplacesCorHdrNumericDefines = 8

//
const Qt__COR_VTABLE_32BIT Qt__ReplacesCorHdrNumericDefines = 1

//
const Qt__COR_VTABLE_64BIT Qt__ReplacesCorHdrNumericDefines = 2

//
const Qt__COR_VTABLE_FROM_UNMANAGED Qt__ReplacesCorHdrNumericDefines = 4

//
const Qt__COR_VTABLE_CALL_MOST_DERIVED Qt__ReplacesCorHdrNumericDefines = 16

//
const Qt__IMAGE_COR_EATJ_THUNK_SIZE Qt__ReplacesCorHdrNumericDefines = 32

//
const Qt__MAX_CLASS_NAME Qt__ReplacesCorHdrNumericDefines = 1024

//
const Qt__MAX_PACKAGE_NAME Qt__ReplacesCorHdrNumericDefines = 1024

func ReplacesCorHdrNumericDefinesItemName(val int) string {
	switch val {
	case Qt__COMIMAGE_FLAGS_ILONLY: // 1
		return "COMIMAGE_FLAGS_ILONLY,NATIVE_TYPE_MAX_CB,IMAGE_COR_MIH_METHODRVA,COR_VTABLE_32BIT"
	case Qt__COMIMAGE_FLAGS_32BITREQUIRED: // 2
		return "COMIMAGE_FLAGS_32BITREQUIRED,COR_VERSION_MAJOR_V2,COR_VERSION_MAJOR,IMAGE_COR_MIH_EHRVA,COR_VTABLE_64BIT"
	case Qt__COMIMAGE_FLAGS_IL_LIBRARY: // 4
		return "COMIMAGE_FLAGS_IL_LIBRARY,COR_VTABLE_FROM_UNMANAGED"
	case Qt__COMIMAGE_FLAGS_STRONGNAMESIGNED: // 8
		return "COMIMAGE_FLAGS_STRONGNAMESIGNED,COR_DELETED_NAME_LENGTH,COR_VTABLEGAP_NAME_LENGTH,IMAGE_COR_MIH_BASICBLOCK"
	case Qt__COMIMAGE_FLAGS_NATIVE_ENTRYPOINT: // 16
		return "COMIMAGE_FLAGS_NATIVE_ENTRYPOINT,COR_VTABLE_CALL_MOST_DERIVED"
	case Qt__COMIMAGE_FLAGS_TRACKDEBUGDATA: // 65536
		return "COMIMAGE_FLAGS_TRACKDEBUGDATA"
	case Qt__COMIMAGE_FLAGS_32BITPREFERRED: // 131072
		return "COMIMAGE_FLAGS_32BITPREFERRED"
		// case Qt__COR_VERSION_MAJOR_V2: // 2
		// return ""
		// case Qt__COR_VERSION_MAJOR: // 2
		// return ""
	case Qt__COR_VERSION_MINOR: // 5
		return "COR_VERSION_MINOR"
		// case Qt__COR_DELETED_NAME_LENGTH: // 8
		// return ""
		// case Qt__COR_VTABLEGAP_NAME_LENGTH: // 8
		// return ""
		// case Qt__NATIVE_TYPE_MAX_CB: // 1
		// return ""
	case Qt__COR_ILMETHOD_SECT_SMALL_MAX_DATASIZE: // 255
		return "COR_ILMETHOD_SECT_SMALL_MAX_DATASIZE"
		// case Qt__IMAGE_COR_MIH_METHODRVA: // 1
		// return ""
		// case Qt__IMAGE_COR_MIH_EHRVA: // 2
		// return ""
		// case Qt__IMAGE_COR_MIH_BASICBLOCK: // 8
		// return ""
		// case Qt__COR_VTABLE_32BIT: // 1
		// return ""
		// case Qt__COR_VTABLE_64BIT: // 2
		// return ""
		// case Qt__COR_VTABLE_FROM_UNMANAGED: // 4
		// return ""
		// case Qt__COR_VTABLE_CALL_MOST_DERIVED: // 16
		// return ""
	case Qt__IMAGE_COR_EATJ_THUNK_SIZE: // 32
		return "IMAGE_COR_EATJ_THUNK_SIZE"
	case Qt__MAX_CLASS_NAME: // 1024
		return "MAX_CLASS_NAME,MAX_PACKAGE_NAME"
		// case Qt__MAX_PACKAGE_NAME: // 1024
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___TOKEN_ELEVATION_TYPE = int // stdglobal
//
const Qt__TokenElevationTypeDefault Qt___TOKEN_ELEVATION_TYPE = 1

//
const Qt__TokenElevationTypeFull Qt___TOKEN_ELEVATION_TYPE = 2

//
const Qt__TokenElevationTypeLimited Qt___TOKEN_ELEVATION_TYPE = 3

func _TOKEN_ELEVATION_TYPEItemName(val int) string {
	switch val {
	case Qt__TokenElevationTypeDefault: // 1
		return "TokenElevationTypeDefault"
	case Qt__TokenElevationTypeFull: // 2
		return "TokenElevationTypeFull"
	case Qt__TokenElevationTypeLimited: // 3
		return "TokenElevationTypeLimited"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___TOKEN_INFORMATION_CLASS = int // stdglobal
//
const Qt__TokenUser Qt___TOKEN_INFORMATION_CLASS = 1

//
const Qt__TokenGroups Qt___TOKEN_INFORMATION_CLASS = 2

//
const Qt__TokenPrivileges Qt___TOKEN_INFORMATION_CLASS = 3

//
const Qt__TokenOwner Qt___TOKEN_INFORMATION_CLASS = 4

//
const Qt__TokenPrimaryGroup Qt___TOKEN_INFORMATION_CLASS = 5

//
const Qt__TokenDefaultDacl Qt___TOKEN_INFORMATION_CLASS = 6

//
const Qt__TokenSource Qt___TOKEN_INFORMATION_CLASS = 7

//
const Qt__TokenType Qt___TOKEN_INFORMATION_CLASS = 8

//
const Qt__TokenImpersonationLevel Qt___TOKEN_INFORMATION_CLASS = 9

//
const Qt__TokenStatistics Qt___TOKEN_INFORMATION_CLASS = 10

//
const Qt__TokenRestrictedSids Qt___TOKEN_INFORMATION_CLASS = 11

//
const Qt__TokenSessionId Qt___TOKEN_INFORMATION_CLASS = 12

//
const Qt__TokenGroupsAndPrivileges Qt___TOKEN_INFORMATION_CLASS = 13

//
const Qt__TokenSessionReference Qt___TOKEN_INFORMATION_CLASS = 14

//
const Qt__TokenSandBoxInert Qt___TOKEN_INFORMATION_CLASS = 15

//
const Qt__TokenAuditPolicy Qt___TOKEN_INFORMATION_CLASS = 16

//
const Qt__TokenOrigin Qt___TOKEN_INFORMATION_CLASS = 17

//
const Qt__TokenElevationType Qt___TOKEN_INFORMATION_CLASS = 18

//
const Qt__TokenLinkedToken Qt___TOKEN_INFORMATION_CLASS = 19

//
const Qt__TokenElevation Qt___TOKEN_INFORMATION_CLASS = 20

//
const Qt__TokenHasRestrictions Qt___TOKEN_INFORMATION_CLASS = 21

//
const Qt__TokenAccessInformation Qt___TOKEN_INFORMATION_CLASS = 22

//
const Qt__TokenVirtualizationAllowed Qt___TOKEN_INFORMATION_CLASS = 23

//
const Qt__TokenVirtualizationEnabled Qt___TOKEN_INFORMATION_CLASS = 24

//
const Qt__TokenIntegrityLevel Qt___TOKEN_INFORMATION_CLASS = 25

//
const Qt__TokenUIAccess Qt___TOKEN_INFORMATION_CLASS = 26

//
const Qt__TokenMandatoryPolicy Qt___TOKEN_INFORMATION_CLASS = 27

//
const Qt__TokenLogonSid Qt___TOKEN_INFORMATION_CLASS = 28

//
const Qt__TokenIsAppContainer Qt___TOKEN_INFORMATION_CLASS = 29

//
const Qt__TokenCapabilities Qt___TOKEN_INFORMATION_CLASS = 30

//
const Qt__TokenAppContainerSid Qt___TOKEN_INFORMATION_CLASS = 31

//
const Qt__TokenAppContainerNumber Qt___TOKEN_INFORMATION_CLASS = 32

//
const Qt__TokenUserClaimAttributes Qt___TOKEN_INFORMATION_CLASS = 33

//
const Qt__TokenDeviceClaimAttributes Qt___TOKEN_INFORMATION_CLASS = 34

//
const Qt__TokenRestrictedUserClaimAttributes Qt___TOKEN_INFORMATION_CLASS = 35

//
const Qt__TokenRestrictedDeviceClaimAttributes Qt___TOKEN_INFORMATION_CLASS = 36

//
const Qt__TokenDeviceGroups Qt___TOKEN_INFORMATION_CLASS = 37

//
const Qt__TokenRestrictedDeviceGroups Qt___TOKEN_INFORMATION_CLASS = 38

//
const Qt__TokenSecurityAttributes Qt___TOKEN_INFORMATION_CLASS = 39

//
const Qt__TokenIsRestricted Qt___TOKEN_INFORMATION_CLASS = 40

//
const Qt__TokenProcessTrustLevel Qt___TOKEN_INFORMATION_CLASS = 41

//
const Qt__MaxTokenInfoClass Qt___TOKEN_INFORMATION_CLASS = 42

func _TOKEN_INFORMATION_CLASSItemName(val int) string {
	switch val {
	case Qt__TokenUser: // 1
		return "TokenUser"
	case Qt__TokenGroups: // 2
		return "TokenGroups"
	case Qt__TokenPrivileges: // 3
		return "TokenPrivileges"
	case Qt__TokenOwner: // 4
		return "TokenOwner"
	case Qt__TokenPrimaryGroup: // 5
		return "TokenPrimaryGroup"
	case Qt__TokenDefaultDacl: // 6
		return "TokenDefaultDacl"
	case Qt__TokenSource: // 7
		return "TokenSource"
	case Qt__TokenType: // 8
		return "TokenType"
	case Qt__TokenImpersonationLevel: // 9
		return "TokenImpersonationLevel"
	case Qt__TokenStatistics: // 10
		return "TokenStatistics"
	case Qt__TokenRestrictedSids: // 11
		return "TokenRestrictedSids"
	case Qt__TokenSessionId: // 12
		return "TokenSessionId"
	case Qt__TokenGroupsAndPrivileges: // 13
		return "TokenGroupsAndPrivileges"
	case Qt__TokenSessionReference: // 14
		return "TokenSessionReference"
	case Qt__TokenSandBoxInert: // 15
		return "TokenSandBoxInert"
	case Qt__TokenAuditPolicy: // 16
		return "TokenAuditPolicy"
	case Qt__TokenOrigin: // 17
		return "TokenOrigin"
	case Qt__TokenElevationType: // 18
		return "TokenElevationType"
	case Qt__TokenLinkedToken: // 19
		return "TokenLinkedToken"
	case Qt__TokenElevation: // 20
		return "TokenElevation"
	case Qt__TokenHasRestrictions: // 21
		return "TokenHasRestrictions"
	case Qt__TokenAccessInformation: // 22
		return "TokenAccessInformation"
	case Qt__TokenVirtualizationAllowed: // 23
		return "TokenVirtualizationAllowed"
	case Qt__TokenVirtualizationEnabled: // 24
		return "TokenVirtualizationEnabled"
	case Qt__TokenIntegrityLevel: // 25
		return "TokenIntegrityLevel"
	case Qt__TokenUIAccess: // 26
		return "TokenUIAccess"
	case Qt__TokenMandatoryPolicy: // 27
		return "TokenMandatoryPolicy"
	case Qt__TokenLogonSid: // 28
		return "TokenLogonSid"
	case Qt__TokenIsAppContainer: // 29
		return "TokenIsAppContainer"
	case Qt__TokenCapabilities: // 30
		return "TokenCapabilities"
	case Qt__TokenAppContainerSid: // 31
		return "TokenAppContainerSid"
	case Qt__TokenAppContainerNumber: // 32
		return "TokenAppContainerNumber"
	case Qt__TokenUserClaimAttributes: // 33
		return "TokenUserClaimAttributes"
	case Qt__TokenDeviceClaimAttributes: // 34
		return "TokenDeviceClaimAttributes"
	case Qt__TokenRestrictedUserClaimAttributes: // 35
		return "TokenRestrictedUserClaimAttributes"
	case Qt__TokenRestrictedDeviceClaimAttributes: // 36
		return "TokenRestrictedDeviceClaimAttributes"
	case Qt__TokenDeviceGroups: // 37
		return "TokenDeviceGroups"
	case Qt__TokenRestrictedDeviceGroups: // 38
		return "TokenRestrictedDeviceGroups"
	case Qt__TokenSecurityAttributes: // 39
		return "TokenSecurityAttributes"
	case Qt__TokenIsRestricted: // 40
		return "TokenIsRestricted"
	case Qt__TokenProcessTrustLevel: // 41
		return "TokenProcessTrustLevel"
	case Qt__MaxTokenInfoClass: // 42
		return "MaxTokenInfoClass"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___ACL_INFORMATION_CLASS = int // stdglobal
//
const Qt__AclRevisionInformation Qt___ACL_INFORMATION_CLASS = 1

//
const Qt__AclSizeInformation Qt___ACL_INFORMATION_CLASS = 2

func _ACL_INFORMATION_CLASSItemName(val int) string {
	switch val {
	case Qt__AclRevisionInformation: // 1
		return "AclRevisionInformation"
	case Qt__AclSizeInformation: // 2
		return "AclSizeInformation"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagTOKEN_TYPE = int // stdglobal
//
const Qt__TokenPrimary Qt__tagTOKEN_TYPE = 1

//
const Qt__TokenImpersonation Qt__tagTOKEN_TYPE = 2

func tagTOKEN_TYPEItemName(val int) string {
	switch val {
	case Qt__TokenPrimary: // 1
		return "TokenPrimary"
	case Qt__TokenImpersonation: // 2
		return "TokenImpersonation"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___SECURITY_IMPERSONATION_LEVEL = int // stdglobal
//
const Qt__SecurityAnonymous Qt___SECURITY_IMPERSONATION_LEVEL = 0

//
const Qt__SecurityIdentification Qt___SECURITY_IMPERSONATION_LEVEL = 1

//
const Qt__SecurityImpersonation Qt___SECURITY_IMPERSONATION_LEVEL = 2

//
const Qt__SecurityDelegation Qt___SECURITY_IMPERSONATION_LEVEL = 3

func _SECURITY_IMPERSONATION_LEVELItemName(val int) string {
	switch val {
	case Qt__SecurityAnonymous: // 0
		return "SecurityAnonymous"
	case Qt__SecurityIdentification: // 1
		return "SecurityIdentification"
	case Qt__SecurityImpersonation: // 2
		return "SecurityImpersonation"
	case Qt__SecurityDelegation: // 3
		return "SecurityDelegation"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSID_NAME_USE = int // stdglobal
//
const Qt__SidTypeUser Qt__tagSID_NAME_USE = 1

//
const Qt__SidTypeGroup Qt__tagSID_NAME_USE = 2

//
const Qt__SidTypeDomain Qt__tagSID_NAME_USE = 3

//
const Qt__SidTypeAlias Qt__tagSID_NAME_USE = 4

//
const Qt__SidTypeWellKnownGroup Qt__tagSID_NAME_USE = 5

//
const Qt__SidTypeDeletedAccount Qt__tagSID_NAME_USE = 6

//
const Qt__SidTypeInvalid Qt__tagSID_NAME_USE = 7

//
const Qt__SidTypeUnknown Qt__tagSID_NAME_USE = 8

func tagSID_NAME_USEItemName(val int) string {
	switch val {
	case Qt__SidTypeUser: // 1
		return "SidTypeUser"
	case Qt__SidTypeGroup: // 2
		return "SidTypeGroup"
	case Qt__SidTypeDomain: // 3
		return "SidTypeDomain"
	case Qt__SidTypeAlias: // 4
		return "SidTypeAlias"
	case Qt__SidTypeWellKnownGroup: // 5
		return "SidTypeWellKnownGroup"
	case Qt__SidTypeDeletedAccount: // 6
		return "SidTypeDeletedAccount"
	case Qt__SidTypeInvalid: // 7
		return "SidTypeInvalid"
	case Qt__SidTypeUnknown: // 8
		return "SidTypeUnknown"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___LATENCY_TIME = int // stdglobal
//
const Qt__LT_DONT_CARE Qt___LATENCY_TIME = 0

//
const Qt__LT_LOWEST_LATENCY Qt___LATENCY_TIME = 1

func _LATENCY_TIMEItemName(val int) string {
	switch val {
	case Qt__LT_DONT_CARE: // 0
		return "LT_DONT_CARE"
	case Qt__LT_LOWEST_LATENCY: // 1
		return "LT_LOWEST_LATENCY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___POWER_ACTION = int // stdglobal
//
const Qt__PowerActionNone Qt___POWER_ACTION = 0

//
const Qt__PowerActionReserved Qt___POWER_ACTION = 1

//
const Qt__PowerActionSleep Qt___POWER_ACTION = 2

//
const Qt__PowerActionHibernate Qt___POWER_ACTION = 3

//
const Qt__PowerActionShutdown Qt___POWER_ACTION = 4

//
const Qt__PowerActionShutdownReset Qt___POWER_ACTION = 5

//
const Qt__PowerActionShutdownOff Qt___POWER_ACTION = 6

//
const Qt__PowerActionWarmEject Qt___POWER_ACTION = 7

func _POWER_ACTIONItemName(val int) string {
	switch val {
	case Qt__PowerActionNone: // 0
		return "PowerActionNone"
	case Qt__PowerActionReserved: // 1
		return "PowerActionReserved"
	case Qt__PowerActionSleep: // 2
		return "PowerActionSleep"
	case Qt__PowerActionHibernate: // 3
		return "PowerActionHibernate"
	case Qt__PowerActionShutdown: // 4
		return "PowerActionShutdown"
	case Qt__PowerActionShutdownReset: // 5
		return "PowerActionShutdownReset"
	case Qt__PowerActionShutdownOff: // 6
		return "PowerActionShutdownOff"
	case Qt__PowerActionWarmEject: // 7
		return "PowerActionWarmEject"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___POWER_PLATFORM_ROLE = int // stdglobal
//
const Qt__PlatformRoleUnspecified Qt___POWER_PLATFORM_ROLE = 0

//
const Qt__PlatformRoleDesktop Qt___POWER_PLATFORM_ROLE = 1

//
const Qt__PlatformRoleMobile Qt___POWER_PLATFORM_ROLE = 2

//
const Qt__PlatformRoleWorkstation Qt___POWER_PLATFORM_ROLE = 3

//
const Qt__PlatformRoleEnterpriseServer Qt___POWER_PLATFORM_ROLE = 4

//
const Qt__PlatformRoleSOHOServer Qt___POWER_PLATFORM_ROLE = 5

//
const Qt__PlatformRoleAppliancePC Qt___POWER_PLATFORM_ROLE = 6

//
const Qt__PlatformRolePerformanceServer Qt___POWER_PLATFORM_ROLE = 7

//
const Qt__PlatformRoleSlate Qt___POWER_PLATFORM_ROLE = 8

//
const Qt__PlatformRoleMaximum Qt___POWER_PLATFORM_ROLE = 9

func _POWER_PLATFORM_ROLEItemName(val int) string {
	switch val {
	case Qt__PlatformRoleUnspecified: // 0
		return "PlatformRoleUnspecified"
	case Qt__PlatformRoleDesktop: // 1
		return "PlatformRoleDesktop"
	case Qt__PlatformRoleMobile: // 2
		return "PlatformRoleMobile"
	case Qt__PlatformRoleWorkstation: // 3
		return "PlatformRoleWorkstation"
	case Qt__PlatformRoleEnterpriseServer: // 4
		return "PlatformRoleEnterpriseServer"
	case Qt__PlatformRoleSOHOServer: // 5
		return "PlatformRoleSOHOServer"
	case Qt__PlatformRoleAppliancePC: // 6
		return "PlatformRoleAppliancePC"
	case Qt__PlatformRolePerformanceServer: // 7
		return "PlatformRolePerformanceServer"
	case Qt__PlatformRoleSlate: // 8
		return "PlatformRoleSlate"
	case Qt__PlatformRoleMaximum: // 9
		return "PlatformRoleMaximum"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___SYSTEM_POWER_STATE = int // stdglobal
//
const Qt__PowerSystemUnspecified Qt___SYSTEM_POWER_STATE = 0

//
const Qt__PowerSystemWorking Qt___SYSTEM_POWER_STATE = 1

//
const Qt__PowerSystemSleeping1 Qt___SYSTEM_POWER_STATE = 2

//
const Qt__PowerSystemSleeping2 Qt___SYSTEM_POWER_STATE = 3

//
const Qt__PowerSystemSleeping3 Qt___SYSTEM_POWER_STATE = 4

//
const Qt__PowerSystemHibernate Qt___SYSTEM_POWER_STATE = 5

//
const Qt__PowerSystemShutdown Qt___SYSTEM_POWER_STATE = 6

//
const Qt__PowerSystemMaximum Qt___SYSTEM_POWER_STATE = 7

func _SYSTEM_POWER_STATEItemName(val int) string {
	switch val {
	case Qt__PowerSystemUnspecified: // 0
		return "PowerSystemUnspecified"
	case Qt__PowerSystemWorking: // 1
		return "PowerSystemWorking"
	case Qt__PowerSystemSleeping1: // 2
		return "PowerSystemSleeping1"
	case Qt__PowerSystemSleeping2: // 3
		return "PowerSystemSleeping2"
	case Qt__PowerSystemSleeping3: // 4
		return "PowerSystemSleeping3"
	case Qt__PowerSystemHibernate: // 5
		return "PowerSystemHibernate"
	case Qt__PowerSystemShutdown: // 6
		return "PowerSystemShutdown"
	case Qt__PowerSystemMaximum: // 7
		return "PowerSystemMaximum"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___DEVICE_POWER_STATE = int // stdglobal
//
const Qt__PowerDeviceUnspecified Qt___DEVICE_POWER_STATE = 0

//
const Qt__PowerDeviceD0 Qt___DEVICE_POWER_STATE = 1

//
const Qt__PowerDeviceD1 Qt___DEVICE_POWER_STATE = 2

//
const Qt__PowerDeviceD2 Qt___DEVICE_POWER_STATE = 3

//
const Qt__PowerDeviceD3 Qt___DEVICE_POWER_STATE = 4

//
const Qt__PowerDeviceMaximum Qt___DEVICE_POWER_STATE = 5

func _DEVICE_POWER_STATEItemName(val int) string {
	switch val {
	case Qt__PowerDeviceUnspecified: // 0
		return "PowerDeviceUnspecified"
	case Qt__PowerDeviceD0: // 1
		return "PowerDeviceD0"
	case Qt__PowerDeviceD1: // 2
		return "PowerDeviceD1"
	case Qt__PowerDeviceD2: // 3
		return "PowerDeviceD2"
	case Qt__PowerDeviceD3: // 4
		return "PowerDeviceD3"
	case Qt__PowerDeviceMaximum: // 5
		return "PowerDeviceMaximum"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___POWER_INFORMATION_LEVEL = int // stdglobal
//
const Qt__SystemPowerPolicyAc Qt___POWER_INFORMATION_LEVEL = 0

//
const Qt__SystemPowerPolicyDc Qt___POWER_INFORMATION_LEVEL = 1

//
const Qt__VerifySystemPolicyAc Qt___POWER_INFORMATION_LEVEL = 2

//
const Qt__VerifySystemPolicyDc Qt___POWER_INFORMATION_LEVEL = 3

//
const Qt__SystemPowerCapabilities Qt___POWER_INFORMATION_LEVEL = 4

//
const Qt__SystemBatteryState Qt___POWER_INFORMATION_LEVEL = 5

//
const Qt__SystemPowerStateHandler Qt___POWER_INFORMATION_LEVEL = 6

//
const Qt__ProcessorStateHandler Qt___POWER_INFORMATION_LEVEL = 7

//
const Qt__SystemPowerPolicyCurrent Qt___POWER_INFORMATION_LEVEL = 8

//
const Qt__AdministratorPowerPolicy Qt___POWER_INFORMATION_LEVEL = 9

//
const Qt__SystemReserveHiberFile Qt___POWER_INFORMATION_LEVEL = 10

//
const Qt__ProcessorInformation Qt___POWER_INFORMATION_LEVEL = 11

//
const Qt__SystemPowerInformation Qt___POWER_INFORMATION_LEVEL = 12

//
const Qt__ProcessorStateHandler2 Qt___POWER_INFORMATION_LEVEL = 13

//
const Qt__LastWakeTime Qt___POWER_INFORMATION_LEVEL = 14

//
const Qt__LastSleepTime Qt___POWER_INFORMATION_LEVEL = 15

//
const Qt__SystemExecutionState Qt___POWER_INFORMATION_LEVEL = 16

//
const Qt__SystemPowerStateNotifyHandler Qt___POWER_INFORMATION_LEVEL = 17

//
const Qt__ProcessorPowerPolicyAc Qt___POWER_INFORMATION_LEVEL = 18

//
const Qt__ProcessorPowerPolicyDc Qt___POWER_INFORMATION_LEVEL = 19

//
const Qt__VerifyProcessorPowerPolicyAc Qt___POWER_INFORMATION_LEVEL = 20

//
const Qt__VerifyProcessorPowerPolicyDc Qt___POWER_INFORMATION_LEVEL = 21

//
const Qt__ProcessorPowerPolicyCurrent Qt___POWER_INFORMATION_LEVEL = 22

func _POWER_INFORMATION_LEVELItemName(val int) string {
	switch val {
	case Qt__SystemPowerPolicyAc: // 0
		return "SystemPowerPolicyAc"
	case Qt__SystemPowerPolicyDc: // 1
		return "SystemPowerPolicyDc"
	case Qt__VerifySystemPolicyAc: // 2
		return "VerifySystemPolicyAc"
	case Qt__VerifySystemPolicyDc: // 3
		return "VerifySystemPolicyDc"
	case Qt__SystemPowerCapabilities: // 4
		return "SystemPowerCapabilities"
	case Qt__SystemBatteryState: // 5
		return "SystemBatteryState"
	case Qt__SystemPowerStateHandler: // 6
		return "SystemPowerStateHandler"
	case Qt__ProcessorStateHandler: // 7
		return "ProcessorStateHandler"
	case Qt__SystemPowerPolicyCurrent: // 8
		return "SystemPowerPolicyCurrent"
	case Qt__AdministratorPowerPolicy: // 9
		return "AdministratorPowerPolicy"
	case Qt__SystemReserveHiberFile: // 10
		return "SystemReserveHiberFile"
	case Qt__ProcessorInformation: // 11
		return "ProcessorInformation"
	case Qt__SystemPowerInformation: // 12
		return "SystemPowerInformation"
	case Qt__ProcessorStateHandler2: // 13
		return "ProcessorStateHandler2"
	case Qt__LastWakeTime: // 14
		return "LastWakeTime"
	case Qt__LastSleepTime: // 15
		return "LastSleepTime"
	case Qt__SystemExecutionState: // 16
		return "SystemExecutionState"
	case Qt__SystemPowerStateNotifyHandler: // 17
		return "SystemPowerStateNotifyHandler"
	case Qt__ProcessorPowerPolicyAc: // 18
		return "ProcessorPowerPolicyAc"
	case Qt__ProcessorPowerPolicyDc: // 19
		return "ProcessorPowerPolicyDc"
	case Qt__VerifyProcessorPowerPolicyAc: // 20
		return "VerifyProcessorPowerPolicyAc"
	case Qt__VerifyProcessorPowerPolicyDc: // 21
		return "VerifyProcessorPowerPolicyDc"
	case Qt__ProcessorPowerPolicyCurrent: // 22
		return "ProcessorPowerPolicyCurrent"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___CM_SERVICE_NODE_TYPE = int // stdglobal
//
const Qt__DriverType Qt___CM_SERVICE_NODE_TYPE = 1

//
const Qt__FileSystemType Qt___CM_SERVICE_NODE_TYPE = 2

//
const Qt__Win32ServiceOwnProcess Qt___CM_SERVICE_NODE_TYPE = 16

//
const Qt__Win32ServiceShareProcess Qt___CM_SERVICE_NODE_TYPE = 32

//
const Qt__AdapterType Qt___CM_SERVICE_NODE_TYPE = 4

//
const Qt__RecognizerType Qt___CM_SERVICE_NODE_TYPE = 8

func _CM_SERVICE_NODE_TYPEItemName(val int) string {
	switch val {
	case Qt__DriverType: // 1
		return "DriverType"
	case Qt__FileSystemType: // 2
		return "FileSystemType"
	case Qt__Win32ServiceOwnProcess: // 16
		return "Win32ServiceOwnProcess"
	case Qt__Win32ServiceShareProcess: // 32
		return "Win32ServiceShareProcess"
	case Qt__AdapterType: // 4
		return "AdapterType"
	case Qt__RecognizerType: // 8
		return "RecognizerType"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___CM_SERVICE_LOAD_TYPE = int // stdglobal
//
const Qt__BootLoad Qt___CM_SERVICE_LOAD_TYPE = 0

//
const Qt__SystemLoad Qt___CM_SERVICE_LOAD_TYPE = 1

//
const Qt__AutoLoad Qt___CM_SERVICE_LOAD_TYPE = 2

//
const Qt__DemandLoad Qt___CM_SERVICE_LOAD_TYPE = 3

//
const Qt__DisableLoad Qt___CM_SERVICE_LOAD_TYPE = 4

func _CM_SERVICE_LOAD_TYPEItemName(val int) string {
	switch val {
	case Qt__BootLoad: // 0
		return "BootLoad"
	case Qt__SystemLoad: // 1
		return "SystemLoad"
	case Qt__AutoLoad: // 2
		return "AutoLoad"
	case Qt__DemandLoad: // 3
		return "DemandLoad"
	case Qt__DisableLoad: // 4
		return "DisableLoad"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___CM_ERROR_CONTROL_TYPE = int // stdglobal
//
const Qt__IgnoreError Qt___CM_ERROR_CONTROL_TYPE = 0

//
const Qt__NormalError Qt___CM_ERROR_CONTROL_TYPE = 1

//
const Qt__SevereError Qt___CM_ERROR_CONTROL_TYPE = 2

//
const Qt__CriticalError Qt___CM_ERROR_CONTROL_TYPE = 3

func _CM_ERROR_CONTROL_TYPEItemName(val int) string {
	switch val {
	case Qt__IgnoreError: // 0
		return "IgnoreError"
	case Qt__NormalError: // 1
		return "NormalError"
	case Qt__SevereError: // 2
		return "SevereError"
	case Qt__CriticalError: // 3
		return "CriticalError"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___ACTIVATION_CONTEXT_INFO_CLASS = int // stdglobal
//
const Qt__ActivationContextBasicInformation Qt___ACTIVATION_CONTEXT_INFO_CLASS = 1

//
const Qt__ActivationContextDetailedInformation Qt___ACTIVATION_CONTEXT_INFO_CLASS = 2

//
const Qt__AssemblyDetailedInformationInActivationContext Qt___ACTIVATION_CONTEXT_INFO_CLASS = 3

//
const Qt__FileInformationInAssemblyOfAssemblyInActivationContext Qt___ACTIVATION_CONTEXT_INFO_CLASS = 4

//
const Qt__RunlevelInformationInActivationContext Qt___ACTIVATION_CONTEXT_INFO_CLASS = 5

//
const Qt__CompatibilityInformationInActivationContext Qt___ACTIVATION_CONTEXT_INFO_CLASS = 6

//
const Qt__ActivationContextManifestResourceName Qt___ACTIVATION_CONTEXT_INFO_CLASS = 7

//
const Qt__MaxActivationContextInfoClass Qt___ACTIVATION_CONTEXT_INFO_CLASS = 8

//
const Qt__AssemblyDetailedInformationInActivationContxt Qt___ACTIVATION_CONTEXT_INFO_CLASS = 3

//
const Qt__FileInformationInAssemblyOfAssemblyInActivationContxt Qt___ACTIVATION_CONTEXT_INFO_CLASS = 4

func _ACTIVATION_CONTEXT_INFO_CLASSItemName(val int) string {
	switch val {
	case Qt__ActivationContextBasicInformation: // 1
		return "ActivationContextBasicInformation"
	case Qt__ActivationContextDetailedInformation: // 2
		return "ActivationContextDetailedInformation"
	case Qt__AssemblyDetailedInformationInActivationContext: // 3
		return "AssemblyDetailedInformationInActivationContext,AssemblyDetailedInformationInActivationContxt"
	case Qt__FileInformationInAssemblyOfAssemblyInActivationContext: // 4
		return "FileInformationInAssemblyOfAssemblyInActivationContext,FileInformationInAssemblyOfAssemblyInActivationContxt"
	case Qt__RunlevelInformationInActivationContext: // 5
		return "RunlevelInformationInActivationContext"
	case Qt__CompatibilityInformationInActivationContext: // 6
		return "CompatibilityInformationInActivationContext"
	case Qt__ActivationContextManifestResourceName: // 7
		return "ActivationContextManifestResourceName"
	case Qt__MaxActivationContextInfoClass: // 8
		return "MaxActivationContextInfoClass"
		// case Qt__AssemblyDetailedInformationInActivationContxt: // 3
		// return ""
		// case Qt__FileInformationInAssemblyOfAssemblyInActivationContxt: // 4
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___JOBOBJECTINFOCLASS = int // stdglobal
//
const Qt__JobObjectBasicAccountingInformation Qt___JOBOBJECTINFOCLASS = 1

//
const Qt__JobObjectBasicLimitInformation Qt___JOBOBJECTINFOCLASS = 2

//
const Qt__JobObjectBasicProcessIdList Qt___JOBOBJECTINFOCLASS = 3

//
const Qt__JobObjectBasicUIRestrictions Qt___JOBOBJECTINFOCLASS = 4

//
const Qt__JobObjectSecurityLimitInformation Qt___JOBOBJECTINFOCLASS = 5

//
const Qt__JobObjectEndOfJobTimeInformation Qt___JOBOBJECTINFOCLASS = 6

//
const Qt__JobObjectAssociateCompletionPortInformation Qt___JOBOBJECTINFOCLASS = 7

//
const Qt__JobObjectBasicAndIoAccountingInformation Qt___JOBOBJECTINFOCLASS = 8

//
const Qt__JobObjectExtendedLimitInformation Qt___JOBOBJECTINFOCLASS = 9

//
const Qt__JobObjectJobSetInformation Qt___JOBOBJECTINFOCLASS = 10

//
const Qt__MaxJobObjectInfoClass Qt___JOBOBJECTINFOCLASS = 11

func _JOBOBJECTINFOCLASSItemName(val int) string {
	switch val {
	case Qt__JobObjectBasicAccountingInformation: // 1
		return "JobObjectBasicAccountingInformation"
	case Qt__JobObjectBasicLimitInformation: // 2
		return "JobObjectBasicLimitInformation"
	case Qt__JobObjectBasicProcessIdList: // 3
		return "JobObjectBasicProcessIdList"
	case Qt__JobObjectBasicUIRestrictions: // 4
		return "JobObjectBasicUIRestrictions"
	case Qt__JobObjectSecurityLimitInformation: // 5
		return "JobObjectSecurityLimitInformation"
	case Qt__JobObjectEndOfJobTimeInformation: // 6
		return "JobObjectEndOfJobTimeInformation"
	case Qt__JobObjectAssociateCompletionPortInformation: // 7
		return "JobObjectAssociateCompletionPortInformation"
	case Qt__JobObjectBasicAndIoAccountingInformation: // 8
		return "JobObjectBasicAndIoAccountingInformation"
	case Qt__JobObjectExtendedLimitInformation: // 9
		return "JobObjectExtendedLimitInformation"
	case Qt__JobObjectJobSetInformation: // 10
		return "JobObjectJobSetInformation"
	case Qt__MaxJobObjectInfoClass: // 11
		return "MaxJobObjectInfoClass"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___LOGICAL_PROCESSOR_RELATIONSHIP = int // stdglobal
//
const Qt__RelationProcessorCore Qt___LOGICAL_PROCESSOR_RELATIONSHIP = 0

//
const Qt__RelationNumaNode Qt___LOGICAL_PROCESSOR_RELATIONSHIP = 1

//
const Qt__RelationCache Qt___LOGICAL_PROCESSOR_RELATIONSHIP = 2

//
const Qt__RelationProcessorPackage Qt___LOGICAL_PROCESSOR_RELATIONSHIP = 3

//
const Qt__RelationGroup Qt___LOGICAL_PROCESSOR_RELATIONSHIP = 4

//
const Qt__RelationAll Qt___LOGICAL_PROCESSOR_RELATIONSHIP = 65535

func _LOGICAL_PROCESSOR_RELATIONSHIPItemName(val int) string {
	switch val {
	case Qt__RelationProcessorCore: // 0
		return "RelationProcessorCore"
	case Qt__RelationNumaNode: // 1
		return "RelationNumaNode"
	case Qt__RelationCache: // 2
		return "RelationCache"
	case Qt__RelationProcessorPackage: // 3
		return "RelationProcessorPackage"
	case Qt__RelationGroup: // 4
		return "RelationGroup"
	case Qt__RelationAll: // 65535
		return "RelationAll"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___PROCESSOR_CACHE_TYPE = int // stdglobal
//
const Qt__CacheUnified Qt___PROCESSOR_CACHE_TYPE = 0

//
const Qt__CacheInstruction Qt___PROCESSOR_CACHE_TYPE = 1

//
const Qt__CacheData Qt___PROCESSOR_CACHE_TYPE = 2

//
const Qt__CacheTrace Qt___PROCESSOR_CACHE_TYPE = 3

func _PROCESSOR_CACHE_TYPEItemName(val int) string {
	switch val {
	case Qt__CacheUnified: // 0
		return "CacheUnified"
	case Qt__CacheInstruction: // 1
		return "CacheInstruction"
	case Qt__CacheData: // 2
		return "CacheData"
	case Qt__CacheTrace: // 3
		return "CacheTrace"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___TP_CALLBACK_PRIORITY = int // stdglobal
//
const Qt__TP_CALLBACK_PRIORITY_HIGH Qt___TP_CALLBACK_PRIORITY = 0

//
const Qt__TP_CALLBACK_PRIORITY_NORMAL Qt___TP_CALLBACK_PRIORITY = 1

//
const Qt__TP_CALLBACK_PRIORITY_LOW Qt___TP_CALLBACK_PRIORITY = 2

//
const Qt__TP_CALLBACK_PRIORITY_INVALID Qt___TP_CALLBACK_PRIORITY = 3

//
const Qt__TP_CALLBACK_PRIORITY_COUNT Qt___TP_CALLBACK_PRIORITY = 3

func _TP_CALLBACK_PRIORITYItemName(val int) string {
	switch val {
	case Qt__TP_CALLBACK_PRIORITY_HIGH: // 0
		return "TP_CALLBACK_PRIORITY_HIGH"
	case Qt__TP_CALLBACK_PRIORITY_NORMAL: // 1
		return "TP_CALLBACK_PRIORITY_NORMAL"
	case Qt__TP_CALLBACK_PRIORITY_LOW: // 2
		return "TP_CALLBACK_PRIORITY_LOW"
	case Qt__TP_CALLBACK_PRIORITY_INVALID: // 3
		return "TP_CALLBACK_PRIORITY_INVALID,TP_CALLBACK_PRIORITY_COUNT"
		// case Qt__TP_CALLBACK_PRIORITY_COUNT: // 3
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___RTL_UMS_THREAD_INFO_CLASS = int // stdglobal
//
const Qt__UmsThreadInvalidInfoClass Qt___RTL_UMS_THREAD_INFO_CLASS = 0

//
const Qt__UmsThreadUserContext Qt___RTL_UMS_THREAD_INFO_CLASS = 1

//
const Qt__UmsThreadPriority Qt___RTL_UMS_THREAD_INFO_CLASS = 2

//
const Qt__UmsThreadAffinity Qt___RTL_UMS_THREAD_INFO_CLASS = 3

//
const Qt__UmsThreadTeb Qt___RTL_UMS_THREAD_INFO_CLASS = 4

//
const Qt__UmsThreadIsSuspended Qt___RTL_UMS_THREAD_INFO_CLASS = 5

//
const Qt__UmsThreadIsTerminated Qt___RTL_UMS_THREAD_INFO_CLASS = 6

//
const Qt__UmsThreadMaxInfoClass Qt___RTL_UMS_THREAD_INFO_CLASS = 7

func _RTL_UMS_THREAD_INFO_CLASSItemName(val int) string {
	switch val {
	case Qt__UmsThreadInvalidInfoClass: // 0
		return "UmsThreadInvalidInfoClass"
	case Qt__UmsThreadUserContext: // 1
		return "UmsThreadUserContext"
	case Qt__UmsThreadPriority: // 2
		return "UmsThreadPriority"
	case Qt__UmsThreadAffinity: // 3
		return "UmsThreadAffinity"
	case Qt__UmsThreadTeb: // 4
		return "UmsThreadTeb"
	case Qt__UmsThreadIsSuspended: // 5
		return "UmsThreadIsSuspended"
	case Qt__UmsThreadIsTerminated: // 6
		return "UmsThreadIsTerminated"
	case Qt__UmsThreadMaxInfoClass: // 7
		return "UmsThreadMaxInfoClass"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___RTL_UMS_SCHEDULER_REASON = int // stdglobal
//
const Qt__UmsSchedulerStartup Qt___RTL_UMS_SCHEDULER_REASON = 0

//
const Qt__UmsSchedulerThreadBlocked Qt___RTL_UMS_SCHEDULER_REASON = 1

//
const Qt__UmsSchedulerThreadYield Qt___RTL_UMS_SCHEDULER_REASON = 2

func _RTL_UMS_SCHEDULER_REASONItemName(val int) string {
	switch val {
	case Qt__UmsSchedulerStartup: // 0
		return "UmsSchedulerStartup"
	case Qt__UmsSchedulerThreadBlocked: // 1
		return "UmsSchedulerThreadBlocked"
	case Qt__UmsSchedulerThreadYield: // 2
		return "UmsSchedulerThreadYield"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___PROCESS_MITIGATION_POLICY = int // stdglobal
//
const Qt__ProcessDEPPolicy Qt___PROCESS_MITIGATION_POLICY = 0

//
const Qt__ProcessASLRPolicy Qt___PROCESS_MITIGATION_POLICY = 1

//
const Qt__ProcessDynamicCodePolicy Qt___PROCESS_MITIGATION_POLICY = 2

//
const Qt__ProcessStrictHandleCheckPolicy Qt___PROCESS_MITIGATION_POLICY = 3

//
const Qt__ProcessSystemCallDisablePolicy Qt___PROCESS_MITIGATION_POLICY = 4

//
const Qt__ProcessMitigationOptionsMask Qt___PROCESS_MITIGATION_POLICY = 5

//
const Qt__ProcessExtensionPointDisablePolicy Qt___PROCESS_MITIGATION_POLICY = 6

//
const Qt__ProcessControlFlowGuardPolicy Qt___PROCESS_MITIGATION_POLICY = 7

//
const Qt__ProcessSignaturePolicy Qt___PROCESS_MITIGATION_POLICY = 8

//
const Qt__ProcessFontDisablePolicy Qt___PROCESS_MITIGATION_POLICY = 9

//
const Qt__ProcessImageLoadPolicy Qt___PROCESS_MITIGATION_POLICY = 10

//
const Qt__ProcessSystemCallFilterPolicy Qt___PROCESS_MITIGATION_POLICY = 11

//
const Qt__ProcessPayloadRestrictionPolicy Qt___PROCESS_MITIGATION_POLICY = 12

//
const Qt__ProcessChildProcessPolicy Qt___PROCESS_MITIGATION_POLICY = 13

//
const Qt__ProcessSideChannelIsolationPolicy Qt___PROCESS_MITIGATION_POLICY = 14

//
const Qt__MaxProcessMitigationPolicy Qt___PROCESS_MITIGATION_POLICY = 15

func _PROCESS_MITIGATION_POLICYItemName(val int) string {
	switch val {
	case Qt__ProcessDEPPolicy: // 0
		return "ProcessDEPPolicy"
	case Qt__ProcessASLRPolicy: // 1
		return "ProcessASLRPolicy"
	case Qt__ProcessDynamicCodePolicy: // 2
		return "ProcessDynamicCodePolicy"
	case Qt__ProcessStrictHandleCheckPolicy: // 3
		return "ProcessStrictHandleCheckPolicy"
	case Qt__ProcessSystemCallDisablePolicy: // 4
		return "ProcessSystemCallDisablePolicy"
	case Qt__ProcessMitigationOptionsMask: // 5
		return "ProcessMitigationOptionsMask"
	case Qt__ProcessExtensionPointDisablePolicy: // 6
		return "ProcessExtensionPointDisablePolicy"
	case Qt__ProcessControlFlowGuardPolicy: // 7
		return "ProcessControlFlowGuardPolicy"
	case Qt__ProcessSignaturePolicy: // 8
		return "ProcessSignaturePolicy"
	case Qt__ProcessFontDisablePolicy: // 9
		return "ProcessFontDisablePolicy"
	case Qt__ProcessImageLoadPolicy: // 10
		return "ProcessImageLoadPolicy"
	case Qt__ProcessSystemCallFilterPolicy: // 11
		return "ProcessSystemCallFilterPolicy"
	case Qt__ProcessPayloadRestrictionPolicy: // 12
		return "ProcessPayloadRestrictionPolicy"
	case Qt__ProcessChildProcessPolicy: // 13
		return "ProcessChildProcessPolicy"
	case Qt__ProcessSideChannelIsolationPolicy: // 14
		return "ProcessSideChannelIsolationPolicy"
	case Qt__MaxProcessMitigationPolicy: // 15
		return "MaxProcessMitigationPolicy"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__DPI_AWARENESS = int // stdglobal
//
const Qt__DPI_AWARENESS_INVALID Qt__DPI_AWARENESS = -1

//
const Qt__DPI_AWARENESS_UNAWARE Qt__DPI_AWARENESS = 0

//
const Qt__DPI_AWARENESS_SYSTEM_AWARE Qt__DPI_AWARENESS = 1

//
const Qt__DPI_AWARENESS_PER_MONITOR_AWARE Qt__DPI_AWARENESS = 2

func DPI_AWARENESSItemName(val int) string {
	switch val {
	case Qt__DPI_AWARENESS_INVALID: // -1
		return "DPI_AWARENESS_INVALID"
	case Qt__DPI_AWARENESS_UNAWARE: // 0
		return "DPI_AWARENESS_UNAWARE"
	case Qt__DPI_AWARENESS_SYSTEM_AWARE: // 1
		return "DPI_AWARENESS_SYSTEM_AWARE"
	case Qt__DPI_AWARENESS_PER_MONITOR_AWARE: // 2
		return "DPI_AWARENESS_PER_MONITOR_AWARE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___FINDEX_INFO_LEVELS = int // stdglobal
//
const Qt__FindExInfoStandard Qt___FINDEX_INFO_LEVELS = 0

//
const Qt__FindExInfoBasic Qt___FINDEX_INFO_LEVELS = 1

//
const Qt__FindExInfoMaxInfoLevel Qt___FINDEX_INFO_LEVELS = 2

func _FINDEX_INFO_LEVELSItemName(val int) string {
	switch val {
	case Qt__FindExInfoStandard: // 0
		return "FindExInfoStandard"
	case Qt__FindExInfoBasic: // 1
		return "FindExInfoBasic"
	case Qt__FindExInfoMaxInfoLevel: // 2
		return "FindExInfoMaxInfoLevel"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___FINDEX_SEARCH_OPS = int // stdglobal
//
const Qt__FindExSearchNameMatch Qt___FINDEX_SEARCH_OPS = 0

//
const Qt__FindExSearchLimitToDirectories Qt___FINDEX_SEARCH_OPS = 1

//
const Qt__FindExSearchLimitToDevices Qt___FINDEX_SEARCH_OPS = 2

//
const Qt__FindExSearchMaxSearchOp Qt___FINDEX_SEARCH_OPS = 3

func _FINDEX_SEARCH_OPSItemName(val int) string {
	switch val {
	case Qt__FindExSearchNameMatch: // 0
		return "FindExSearchNameMatch"
	case Qt__FindExSearchLimitToDirectories: // 1
		return "FindExSearchLimitToDirectories"
	case Qt__FindExSearchLimitToDevices: // 2
		return "FindExSearchLimitToDevices"
	case Qt__FindExSearchMaxSearchOp: // 3
		return "FindExSearchMaxSearchOp"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___MEMORY_RESOURCE_NOTIFICATION_TYPE = int // stdglobal
//
const Qt__LowMemoryResourceNotification Qt___MEMORY_RESOURCE_NOTIFICATION_TYPE = 0

//
const Qt__HighMemoryResourceNotification Qt___MEMORY_RESOURCE_NOTIFICATION_TYPE = 1

func _MEMORY_RESOURCE_NOTIFICATION_TYPEItemName(val int) string {
	switch val {
	case Qt__LowMemoryResourceNotification: // 0
		return "LowMemoryResourceNotification"
	case Qt__HighMemoryResourceNotification: // 1
		return "HighMemoryResourceNotification"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___FILE_ID_TYPE = int // stdglobal
//
const Qt__FileIdType Qt___FILE_ID_TYPE = 0

//
const Qt__ObjectIdType Qt___FILE_ID_TYPE = 1

//
const Qt__ExtendedFileIdType Qt___FILE_ID_TYPE = 2

//
const Qt__MaximumFileIdType Qt___FILE_ID_TYPE = 3

func _FILE_ID_TYPEItemName(val int) string {
	switch val {
	case Qt__FileIdType: // 0
		return "FileIdType"
	case Qt__ObjectIdType: // 1
		return "ObjectIdType"
	case Qt__ExtendedFileIdType: // 2
		return "ExtendedFileIdType"
	case Qt__MaximumFileIdType: // 3
		return "MaximumFileIdType"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___FILE_INFO_BY_HANDLE_CLASS = int // stdglobal
//
const Qt__FileBasicInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 0

//
const Qt__FileStandardInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 1

//
const Qt__FileNameInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 2

//
const Qt__FileRenameInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 3

//
const Qt__FileDispositionInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 4

//
const Qt__FileAllocationInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 5

//
const Qt__FileEndOfFileInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 6

//
const Qt__FileStreamInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 7

//
const Qt__FileCompressionInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 8

//
const Qt__FileAttributeTagInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 9

//
const Qt__FileIdBothDirectoryInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 10

//
const Qt__FileIdBothDirectoryRestartInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 11

//
const Qt__FileIoPriorityHintInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 12

//
const Qt__FileRemoteProtocolInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 13

//
const Qt__FileFullDirectoryInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 14

//
const Qt__FileFullDirectoryRestartInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 15

//
const Qt__FileStorageInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 16

//
const Qt__FileAlignmentInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 17

//
const Qt__FileIdInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 18

//
const Qt__FileIdExtdDirectoryInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 19

//
const Qt__FileIdExtdDirectoryRestartInfo Qt___FILE_INFO_BY_HANDLE_CLASS = 20

//
const Qt__MaximumFileInfoByHandlesClass Qt___FILE_INFO_BY_HANDLE_CLASS = 21

func _FILE_INFO_BY_HANDLE_CLASSItemName(val int) string {
	switch val {
	case Qt__FileBasicInfo: // 0
		return "FileBasicInfo"
	case Qt__FileStandardInfo: // 1
		return "FileStandardInfo"
	case Qt__FileNameInfo: // 2
		return "FileNameInfo"
	case Qt__FileRenameInfo: // 3
		return "FileRenameInfo"
	case Qt__FileDispositionInfo: // 4
		return "FileDispositionInfo"
	case Qt__FileAllocationInfo: // 5
		return "FileAllocationInfo"
	case Qt__FileEndOfFileInfo: // 6
		return "FileEndOfFileInfo"
	case Qt__FileStreamInfo: // 7
		return "FileStreamInfo"
	case Qt__FileCompressionInfo: // 8
		return "FileCompressionInfo"
	case Qt__FileAttributeTagInfo: // 9
		return "FileAttributeTagInfo"
	case Qt__FileIdBothDirectoryInfo: // 10
		return "FileIdBothDirectoryInfo"
	case Qt__FileIdBothDirectoryRestartInfo: // 11
		return "FileIdBothDirectoryRestartInfo"
	case Qt__FileIoPriorityHintInfo: // 12
		return "FileIoPriorityHintInfo"
	case Qt__FileRemoteProtocolInfo: // 13
		return "FileRemoteProtocolInfo"
	case Qt__FileFullDirectoryInfo: // 14
		return "FileFullDirectoryInfo"
	case Qt__FileFullDirectoryRestartInfo: // 15
		return "FileFullDirectoryRestartInfo"
	case Qt__FileStorageInfo: // 16
		return "FileStorageInfo"
	case Qt__FileAlignmentInfo: // 17
		return "FileAlignmentInfo"
	case Qt__FileIdInfo: // 18
		return "FileIdInfo"
	case Qt__FileIdExtdDirectoryInfo: // 19
		return "FileIdExtdDirectoryInfo"
	case Qt__FileIdExtdDirectoryRestartInfo: // 20
		return "FileIdExtdDirectoryRestartInfo"
	case Qt__MaximumFileInfoByHandlesClass: // 21
		return "MaximumFileInfoByHandlesClass"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___PRIORITY_HINT = int // stdglobal
//
const Qt__IoPriorityHintVeryLow Qt___PRIORITY_HINT = 0

//
const Qt__IoPriorityHintLow Qt___PRIORITY_HINT = 1

//
const Qt__IoPriorityHintNormal Qt___PRIORITY_HINT = 2

//
const Qt__MaximumIoPriorityHintType Qt___PRIORITY_HINT = 3

func _PRIORITY_HINTItemName(val int) string {
	switch val {
	case Qt__IoPriorityHintVeryLow: // 0
		return "IoPriorityHintVeryLow"
	case Qt__IoPriorityHintLow: // 1
		return "IoPriorityHintLow"
	case Qt__IoPriorityHintNormal: // 2
		return "IoPriorityHintNormal"
	case Qt__MaximumIoPriorityHintType: // 3
		return "MaximumIoPriorityHintType"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___POWER_REQUEST_TYPE = int // stdglobal
//
const Qt__PowerRequestDisplayRequired Qt___POWER_REQUEST_TYPE = 0

//
const Qt__PowerRequestSystemRequired Qt___POWER_REQUEST_TYPE = 1

//
const Qt__PowerRequestAwayModeRequired Qt___POWER_REQUEST_TYPE = 2

func _POWER_REQUEST_TYPEItemName(val int) string {
	switch val {
	case Qt__PowerRequestDisplayRequired: // 0
		return "PowerRequestDisplayRequired"
	case Qt__PowerRequestSystemRequired: // 1
		return "PowerRequestSystemRequired"
	case Qt__PowerRequestAwayModeRequired: // 2
		return "PowerRequestAwayModeRequired"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___GET_FILEEX_INFO_LEVELS = int // stdglobal
//
const Qt__GetFileExInfoStandard Qt___GET_FILEEX_INFO_LEVELS = 0

func _GET_FILEEX_INFO_LEVELSItemName(val int) string {
	switch val {
	case Qt__GetFileExInfoStandard: // 0
		return "GetFileExInfoStandard"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___COPYFILE2_MESSAGE_TYPE = int // stdglobal
//
const Qt__COPYFILE2_CALLBACK_NONE Qt___COPYFILE2_MESSAGE_TYPE = 0

//
const Qt__COPYFILE2_CALLBACK_CHUNK_STARTED Qt___COPYFILE2_MESSAGE_TYPE = 1

//
const Qt__COPYFILE2_CALLBACK_CHUNK_FINISHED Qt___COPYFILE2_MESSAGE_TYPE = 2

//
const Qt__COPYFILE2_CALLBACK_STREAM_STARTED Qt___COPYFILE2_MESSAGE_TYPE = 3

//
const Qt__COPYFILE2_CALLBACK_STREAM_FINISHED Qt___COPYFILE2_MESSAGE_TYPE = 4

//
const Qt__COPYFILE2_CALLBACK_POLL_CONTINUE Qt___COPYFILE2_MESSAGE_TYPE = 5

//
const Qt__COPYFILE2_CALLBACK_ERROR Qt___COPYFILE2_MESSAGE_TYPE = 6

//
const Qt__COPYFILE2_CALLBACK_MAX Qt___COPYFILE2_MESSAGE_TYPE = 7

func _COPYFILE2_MESSAGE_TYPEItemName(val int) string {
	switch val {
	case Qt__COPYFILE2_CALLBACK_NONE: // 0
		return "COPYFILE2_CALLBACK_NONE"
	case Qt__COPYFILE2_CALLBACK_CHUNK_STARTED: // 1
		return "COPYFILE2_CALLBACK_CHUNK_STARTED"
	case Qt__COPYFILE2_CALLBACK_CHUNK_FINISHED: // 2
		return "COPYFILE2_CALLBACK_CHUNK_FINISHED"
	case Qt__COPYFILE2_CALLBACK_STREAM_STARTED: // 3
		return "COPYFILE2_CALLBACK_STREAM_STARTED"
	case Qt__COPYFILE2_CALLBACK_STREAM_FINISHED: // 4
		return "COPYFILE2_CALLBACK_STREAM_FINISHED"
	case Qt__COPYFILE2_CALLBACK_POLL_CONTINUE: // 5
		return "COPYFILE2_CALLBACK_POLL_CONTINUE"
	case Qt__COPYFILE2_CALLBACK_ERROR: // 6
		return "COPYFILE2_CALLBACK_ERROR"
	case Qt__COPYFILE2_CALLBACK_MAX: // 7
		return "COPYFILE2_CALLBACK_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___COPYFILE2_MESSAGE_ACTION = int // stdglobal
//
const Qt__COPYFILE2_PROGRESS_CONTINUE Qt___COPYFILE2_MESSAGE_ACTION = 0

//
const Qt__COPYFILE2_PROGRESS_CANCEL Qt___COPYFILE2_MESSAGE_ACTION = 1

//
const Qt__COPYFILE2_PROGRESS_STOP Qt___COPYFILE2_MESSAGE_ACTION = 2

//
const Qt__COPYFILE2_PROGRESS_QUIET Qt___COPYFILE2_MESSAGE_ACTION = 3

//
const Qt__COPYFILE2_PROGRESS_PAUSE Qt___COPYFILE2_MESSAGE_ACTION = 4

func _COPYFILE2_MESSAGE_ACTIONItemName(val int) string {
	switch val {
	case Qt__COPYFILE2_PROGRESS_CONTINUE: // 0
		return "COPYFILE2_PROGRESS_CONTINUE"
	case Qt__COPYFILE2_PROGRESS_CANCEL: // 1
		return "COPYFILE2_PROGRESS_CANCEL"
	case Qt__COPYFILE2_PROGRESS_STOP: // 2
		return "COPYFILE2_PROGRESS_STOP"
	case Qt__COPYFILE2_PROGRESS_QUIET: // 3
		return "COPYFILE2_PROGRESS_QUIET"
	case Qt__COPYFILE2_PROGRESS_PAUSE: // 4
		return "COPYFILE2_PROGRESS_PAUSE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___COPYFILE2_COPY_PHASE = int // stdglobal
//
const Qt__COPYFILE2_PHASE_NONE Qt___COPYFILE2_COPY_PHASE = 0

//
const Qt__COPYFILE2_PHASE_PREPARE_SOURCE Qt___COPYFILE2_COPY_PHASE = 1

//
const Qt__COPYFILE2_PHASE_PREPARE_DEST Qt___COPYFILE2_COPY_PHASE = 2

//
const Qt__COPYFILE2_PHASE_READ_SOURCE Qt___COPYFILE2_COPY_PHASE = 3

//
const Qt__COPYFILE2_PHASE_WRITE_DESTINATION Qt___COPYFILE2_COPY_PHASE = 4

//
const Qt__COPYFILE2_PHASE_SERVER_COPY Qt___COPYFILE2_COPY_PHASE = 5

//
const Qt__COPYFILE2_PHASE_NAMEGRAFT_COPY Qt___COPYFILE2_COPY_PHASE = 6

//
const Qt__COPYFILE2_PHASE_MAX Qt___COPYFILE2_COPY_PHASE = 7

func _COPYFILE2_COPY_PHASEItemName(val int) string {
	switch val {
	case Qt__COPYFILE2_PHASE_NONE: // 0
		return "COPYFILE2_PHASE_NONE"
	case Qt__COPYFILE2_PHASE_PREPARE_SOURCE: // 1
		return "COPYFILE2_PHASE_PREPARE_SOURCE"
	case Qt__COPYFILE2_PHASE_PREPARE_DEST: // 2
		return "COPYFILE2_PHASE_PREPARE_DEST"
	case Qt__COPYFILE2_PHASE_READ_SOURCE: // 3
		return "COPYFILE2_PHASE_READ_SOURCE"
	case Qt__COPYFILE2_PHASE_WRITE_DESTINATION: // 4
		return "COPYFILE2_PHASE_WRITE_DESTINATION"
	case Qt__COPYFILE2_PHASE_SERVER_COPY: // 5
		return "COPYFILE2_PHASE_SERVER_COPY"
	case Qt__COPYFILE2_PHASE_NAMEGRAFT_COPY: // 6
		return "COPYFILE2_PHASE_NAMEGRAFT_COPY"
	case Qt__COPYFILE2_PHASE_MAX: // 7
		return "COPYFILE2_PHASE_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___COMPUTER_NAME_FORMAT = int // stdglobal
//
const Qt__ComputerNameNetBIOS Qt___COMPUTER_NAME_FORMAT = 0

//
const Qt__ComputerNameDnsHostname Qt___COMPUTER_NAME_FORMAT = 1

//
const Qt__ComputerNameDnsDomain Qt___COMPUTER_NAME_FORMAT = 2

//
const Qt__ComputerNameDnsFullyQualified Qt___COMPUTER_NAME_FORMAT = 3

//
const Qt__ComputerNamePhysicalNetBIOS Qt___COMPUTER_NAME_FORMAT = 4

//
const Qt__ComputerNamePhysicalDnsHostname Qt___COMPUTER_NAME_FORMAT = 5

//
const Qt__ComputerNamePhysicalDnsDomain Qt___COMPUTER_NAME_FORMAT = 6

//
const Qt__ComputerNamePhysicalDnsFullyQualified Qt___COMPUTER_NAME_FORMAT = 7

//
const Qt__ComputerNameMax Qt___COMPUTER_NAME_FORMAT = 8

func _COMPUTER_NAME_FORMATItemName(val int) string {
	switch val {
	case Qt__ComputerNameNetBIOS: // 0
		return "ComputerNameNetBIOS"
	case Qt__ComputerNameDnsHostname: // 1
		return "ComputerNameDnsHostname"
	case Qt__ComputerNameDnsDomain: // 2
		return "ComputerNameDnsDomain"
	case Qt__ComputerNameDnsFullyQualified: // 3
		return "ComputerNameDnsFullyQualified"
	case Qt__ComputerNamePhysicalNetBIOS: // 4
		return "ComputerNamePhysicalNetBIOS"
	case Qt__ComputerNamePhysicalDnsHostname: // 5
		return "ComputerNamePhysicalDnsHostname"
	case Qt__ComputerNamePhysicalDnsDomain: // 6
		return "ComputerNamePhysicalDnsDomain"
	case Qt__ComputerNamePhysicalDnsFullyQualified: // 7
		return "ComputerNamePhysicalDnsFullyQualified"
	case Qt__ComputerNameMax: // 8
		return "ComputerNameMax"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___DEP_SYSTEM_POLICY_TYPE = int // stdglobal
//
const Qt__AlwaysOff Qt___DEP_SYSTEM_POLICY_TYPE = 0

//
const Qt__AlwaysOn Qt___DEP_SYSTEM_POLICY_TYPE = 1

//
const Qt__OptIn Qt___DEP_SYSTEM_POLICY_TYPE = 2

//
const Qt__OptOut Qt___DEP_SYSTEM_POLICY_TYPE = 3

func _DEP_SYSTEM_POLICY_TYPEItemName(val int) string {
	switch val {
	case Qt__AlwaysOff: // 0
		return "AlwaysOff"
	case Qt__AlwaysOn: // 1
		return "AlwaysOn"
	case Qt__OptIn: // 2
		return "OptIn"
	case Qt__OptOut: // 3
		return "OptOut"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___PROC_THREAD_ATTRIBUTE_NUM = int // stdglobal
//
const Qt__ProcThreadAttributeParentProcess Qt___PROC_THREAD_ATTRIBUTE_NUM = 0

//
const Qt__ProcThreadAttributeHandleList Qt___PROC_THREAD_ATTRIBUTE_NUM = 2

//
const Qt__ProcThreadAttributeGroupAffinity Qt___PROC_THREAD_ATTRIBUTE_NUM = 3

//
const Qt__ProcThreadAttributeIdealProcessor Qt___PROC_THREAD_ATTRIBUTE_NUM = 5

//
const Qt__ProcThreadAttributeUmsThread Qt___PROC_THREAD_ATTRIBUTE_NUM = 6

//
const Qt__ProcThreadAttributeMitigationPolicy Qt___PROC_THREAD_ATTRIBUTE_NUM = 7

//
const Qt__ProcThreadAttributeSecurityCapabilities Qt___PROC_THREAD_ATTRIBUTE_NUM = 9

//
const Qt__ProcThreadAttributeProtectionLevel Qt___PROC_THREAD_ATTRIBUTE_NUM = 11

//
const Qt__ProcThreadAttributeJobList Qt___PROC_THREAD_ATTRIBUTE_NUM = 13

//
const Qt__ProcThreadAttributeChildProcessPolicy Qt___PROC_THREAD_ATTRIBUTE_NUM = 14

//
const Qt__ProcThreadAttributeAllApplicationPackagesPolicy Qt___PROC_THREAD_ATTRIBUTE_NUM = 15

//
const Qt__ProcThreadAttributeWin32kFilter Qt___PROC_THREAD_ATTRIBUTE_NUM = 16

//
const Qt__ProcThreadAttributeSafeOpenPromptOriginClaim Qt___PROC_THREAD_ATTRIBUTE_NUM = 17

func _PROC_THREAD_ATTRIBUTE_NUMItemName(val int) string {
	switch val {
	case Qt__ProcThreadAttributeParentProcess: // 0
		return "ProcThreadAttributeParentProcess"
	case Qt__ProcThreadAttributeHandleList: // 2
		return "ProcThreadAttributeHandleList"
	case Qt__ProcThreadAttributeGroupAffinity: // 3
		return "ProcThreadAttributeGroupAffinity"
	case Qt__ProcThreadAttributeIdealProcessor: // 5
		return "ProcThreadAttributeIdealProcessor"
	case Qt__ProcThreadAttributeUmsThread: // 6
		return "ProcThreadAttributeUmsThread"
	case Qt__ProcThreadAttributeMitigationPolicy: // 7
		return "ProcThreadAttributeMitigationPolicy"
	case Qt__ProcThreadAttributeSecurityCapabilities: // 9
		return "ProcThreadAttributeSecurityCapabilities"
	case Qt__ProcThreadAttributeProtectionLevel: // 11
		return "ProcThreadAttributeProtectionLevel"
	case Qt__ProcThreadAttributeJobList: // 13
		return "ProcThreadAttributeJobList"
	case Qt__ProcThreadAttributeChildProcessPolicy: // 14
		return "ProcThreadAttributeChildProcessPolicy"
	case Qt__ProcThreadAttributeAllApplicationPackagesPolicy: // 15
		return "ProcThreadAttributeAllApplicationPackagesPolicy"
	case Qt__ProcThreadAttributeWin32kFilter: // 16
		return "ProcThreadAttributeWin32kFilter"
	case Qt__ProcThreadAttributeSafeOpenPromptOriginClaim: // 17
		return "ProcThreadAttributeSafeOpenPromptOriginClaim"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagINPUT_MESSAGE_DEVICE_TYPE = int // stdglobal
//
const Qt__IMDT_UNAVAILABLE Qt__tagINPUT_MESSAGE_DEVICE_TYPE = 0

//
const Qt__IMDT_KEYBOARD Qt__tagINPUT_MESSAGE_DEVICE_TYPE = 1

//
const Qt__IMDT_MOUSE Qt__tagINPUT_MESSAGE_DEVICE_TYPE = 2

//
const Qt__IMDT_TOUCH Qt__tagINPUT_MESSAGE_DEVICE_TYPE = 4

//
const Qt__IMDT_PEN Qt__tagINPUT_MESSAGE_DEVICE_TYPE = 8

//
const Qt__IMDT_TOUCHPAD Qt__tagINPUT_MESSAGE_DEVICE_TYPE = 16

func tagINPUT_MESSAGE_DEVICE_TYPEItemName(val int) string {
	switch val {
	case Qt__IMDT_UNAVAILABLE: // 0
		return "IMDT_UNAVAILABLE"
	case Qt__IMDT_KEYBOARD: // 1
		return "IMDT_KEYBOARD"
	case Qt__IMDT_MOUSE: // 2
		return "IMDT_MOUSE"
	case Qt__IMDT_TOUCH: // 4
		return "IMDT_TOUCH"
	case Qt__IMDT_PEN: // 8
		return "IMDT_PEN"
	case Qt__IMDT_TOUCHPAD: // 16
		return "IMDT_TOUCHPAD"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagINPUT_MESSAGE_ORIGIN_ID = int // stdglobal
//
const Qt__IMO_UNAVAILABLE Qt__tagINPUT_MESSAGE_ORIGIN_ID = 0

//
const Qt__IMO_HARDWARE Qt__tagINPUT_MESSAGE_ORIGIN_ID = 1

//
const Qt__IMO_INJECTED Qt__tagINPUT_MESSAGE_ORIGIN_ID = 2

//
const Qt__IMO_SYSTEM Qt__tagINPUT_MESSAGE_ORIGIN_ID = 4

func tagINPUT_MESSAGE_ORIGIN_IDItemName(val int) string {
	switch val {
	case Qt__IMO_UNAVAILABLE: // 0
		return "IMO_UNAVAILABLE"
	case Qt__IMO_HARDWARE: // 1
		return "IMO_HARDWARE"
	case Qt__IMO_INJECTED: // 2
		return "IMO_INJECTED"
	case Qt__IMO_SYSTEM: // 4
		return "IMO_SYSTEM"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagAR_STATE = int // stdglobal
//
const Qt__AR_ENABLED Qt__tagAR_STATE = 0

//
const Qt__AR_DISABLED Qt__tagAR_STATE = 1

//
const Qt__AR_SUPPRESSED Qt__tagAR_STATE = 2

//
const Qt__AR_REMOTESESSION Qt__tagAR_STATE = 4

//
const Qt__AR_MULTIMON Qt__tagAR_STATE = 8

//
const Qt__AR_NOSENSOR Qt__tagAR_STATE = 16

//
const Qt__AR_NOT_SUPPORTED Qt__tagAR_STATE = 32

//
const Qt__AR_DOCKED Qt__tagAR_STATE = 64

//
const Qt__AR_LAPTOP Qt__tagAR_STATE = 128

func tagAR_STATEItemName(val int) string {
	switch val {
	case Qt__AR_ENABLED: // 0
		return "AR_ENABLED"
	case Qt__AR_DISABLED: // 1
		return "AR_DISABLED"
	case Qt__AR_SUPPRESSED: // 2
		return "AR_SUPPRESSED"
	case Qt__AR_REMOTESESSION: // 4
		return "AR_REMOTESESSION"
	case Qt__AR_MULTIMON: // 8
		return "AR_MULTIMON"
	case Qt__AR_NOSENSOR: // 16
		return "AR_NOSENSOR"
	case Qt__AR_NOT_SUPPORTED: // 32
		return "AR_NOT_SUPPORTED"
	case Qt__AR_DOCKED: // 64
		return "AR_DOCKED"
	case Qt__AR_LAPTOP: // 128
		return "AR_LAPTOP"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__ORIENTATION_PREFERENCE = int // stdglobal
//
const Qt__ORIENTATION_PREFERENCE_NONE Qt__ORIENTATION_PREFERENCE = 0

//
const Qt__ORIENTATION_PREFERENCE_LANDSCAPE Qt__ORIENTATION_PREFERENCE = 1

//
const Qt__ORIENTATION_PREFERENCE_PORTRAIT Qt__ORIENTATION_PREFERENCE = 2

//
const Qt__ORIENTATION_PREFERENCE_LANDSCAPE_FLIPPED Qt__ORIENTATION_PREFERENCE = 4

//
const Qt__ORIENTATION_PREFERENCE_PORTRAIT_FLIPPED Qt__ORIENTATION_PREFERENCE = 8

func ORIENTATION_PREFERENCEItemName(val int) string {
	switch val {
	case Qt__ORIENTATION_PREFERENCE_NONE: // 0
		return "ORIENTATION_PREFERENCE_NONE"
	case Qt__ORIENTATION_PREFERENCE_LANDSCAPE: // 1
		return "ORIENTATION_PREFERENCE_LANDSCAPE"
	case Qt__ORIENTATION_PREFERENCE_PORTRAIT: // 2
		return "ORIENTATION_PREFERENCE_PORTRAIT"
	case Qt__ORIENTATION_PREFERENCE_LANDSCAPE_FLIPPED: // 4
		return "ORIENTATION_PREFERENCE_LANDSCAPE_FLIPPED"
	case Qt__ORIENTATION_PREFERENCE_PORTRAIT_FLIPPED: // 8
		return "ORIENTATION_PREFERENCE_PORTRAIT_FLIPPED"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagPOINTER_DEVICE_TYPE = int // stdglobal
//
const Qt__POINTER_DEVICE_TYPE_INTEGRATED_PEN Qt__tagPOINTER_DEVICE_TYPE = 1

//
const Qt__POINTER_DEVICE_TYPE_EXTERNAL_PEN Qt__tagPOINTER_DEVICE_TYPE = 2

//
const Qt__POINTER_DEVICE_TYPE_TOUCH Qt__tagPOINTER_DEVICE_TYPE = 3

//
const Qt__POINTER_DEVICE_TYPE_TOUCH_PAD Qt__tagPOINTER_DEVICE_TYPE = 4

//
const Qt__POINTER_DEVICE_TYPE_MAX Qt__tagPOINTER_DEVICE_TYPE = -1

func tagPOINTER_DEVICE_TYPEItemName(val int) string {
	switch val {
	case Qt__POINTER_DEVICE_TYPE_INTEGRATED_PEN: // 1
		return "POINTER_DEVICE_TYPE_INTEGRATED_PEN"
	case Qt__POINTER_DEVICE_TYPE_EXTERNAL_PEN: // 2
		return "POINTER_DEVICE_TYPE_EXTERNAL_PEN"
	case Qt__POINTER_DEVICE_TYPE_TOUCH: // 3
		return "POINTER_DEVICE_TYPE_TOUCH"
	case Qt__POINTER_DEVICE_TYPE_TOUCH_PAD: // 4
		return "POINTER_DEVICE_TYPE_TOUCH_PAD"
	case Qt__POINTER_DEVICE_TYPE_MAX: // -1
		return "POINTER_DEVICE_TYPE_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagPOINTER_INPUT_TYPE = int // stdglobal
//
const Qt__PT_POINTER Qt__tagPOINTER_INPUT_TYPE = 1

//
const Qt__PT_TOUCH Qt__tagPOINTER_INPUT_TYPE = 2

//
const Qt__PT_PEN Qt__tagPOINTER_INPUT_TYPE = 3

//
const Qt__PT_MOUSE Qt__tagPOINTER_INPUT_TYPE = 4

//
const Qt__PT_TOUCHPAD Qt__tagPOINTER_INPUT_TYPE = 5

func tagPOINTER_INPUT_TYPEItemName(val int) string {
	switch val {
	case Qt__PT_POINTER: // 1
		return "PT_POINTER"
	case Qt__PT_TOUCH: // 2
		return "PT_TOUCH"
	case Qt__PT_PEN: // 3
		return "PT_PEN"
	case Qt__PT_MOUSE: // 4
		return "PT_MOUSE"
	case Qt__PT_TOUCHPAD: // 5
		return "PT_TOUCHPAD"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___NORM_FORM = int // stdglobal
//
const Qt__NormalizationOther Qt___NORM_FORM = 0

//
const Qt__NormalizationC Qt___NORM_FORM = 1

//
const Qt__NormalizationD Qt___NORM_FORM = 2

//
const Qt__NormalizationKC Qt___NORM_FORM = 5

//
const Qt__NormalizationKD Qt___NORM_FORM = 6

func _NORM_FORMItemName(val int) string {
	switch val {
	case Qt__NormalizationOther: // 0
		return "NormalizationOther"
	case Qt__NormalizationC: // 1
		return "NormalizationC"
	case Qt__NormalizationD: // 2
		return "NormalizationD"
	case Qt__NormalizationKC: // 5
		return "NormalizationKC"
	case Qt__NormalizationKD: // 6
		return "NormalizationKD"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__SYSGEOTYPE = int // stdglobal
//
const Qt__GEO_NATION Qt__SYSGEOTYPE = 1

//
const Qt__GEO_LATITUDE Qt__SYSGEOTYPE = 2

//
const Qt__GEO_LONGITUDE Qt__SYSGEOTYPE = 3

//
const Qt__GEO_ISO2 Qt__SYSGEOTYPE = 4

//
const Qt__GEO_ISO3 Qt__SYSGEOTYPE = 5

//
const Qt__GEO_RFC1766 Qt__SYSGEOTYPE = 6

//
const Qt__GEO_LCID Qt__SYSGEOTYPE = 7

//
const Qt__GEO_FRIENDLYNAME Qt__SYSGEOTYPE = 8

//
const Qt__GEO_OFFICIALNAME Qt__SYSGEOTYPE = 9

//
const Qt__GEO_TIMEZONES Qt__SYSGEOTYPE = 10

//
const Qt__GEO_OFFICIALLANGUAGES Qt__SYSGEOTYPE = 11

//
const Qt__GEO_ISO_UN_NUMBER Qt__SYSGEOTYPE = 12

//
const Qt__GEO_PARENT Qt__SYSGEOTYPE = 13

//
const Qt__GEO_DIALINGCODE Qt__SYSGEOTYPE = 14

//
const Qt__GEO_CURRENCYCODE Qt__SYSGEOTYPE = 15

//
const Qt__GEO_CURRENCYSYMBOL Qt__SYSGEOTYPE = 16

//
const Qt__GEO_NAME Qt__SYSGEOTYPE = 17

//
const Qt__GEO_ID Qt__SYSGEOTYPE = 18

func SYSGEOTYPEItemName(val int) string {
	switch val {
	case Qt__GEO_NATION: // 1
		return "GEO_NATION"
	case Qt__GEO_LATITUDE: // 2
		return "GEO_LATITUDE"
	case Qt__GEO_LONGITUDE: // 3
		return "GEO_LONGITUDE"
	case Qt__GEO_ISO2: // 4
		return "GEO_ISO2"
	case Qt__GEO_ISO3: // 5
		return "GEO_ISO3"
	case Qt__GEO_RFC1766: // 6
		return "GEO_RFC1766"
	case Qt__GEO_LCID: // 7
		return "GEO_LCID"
	case Qt__GEO_FRIENDLYNAME: // 8
		return "GEO_FRIENDLYNAME"
	case Qt__GEO_OFFICIALNAME: // 9
		return "GEO_OFFICIALNAME"
	case Qt__GEO_TIMEZONES: // 10
		return "GEO_TIMEZONES"
	case Qt__GEO_OFFICIALLANGUAGES: // 11
		return "GEO_OFFICIALLANGUAGES"
	case Qt__GEO_ISO_UN_NUMBER: // 12
		return "GEO_ISO_UN_NUMBER"
	case Qt__GEO_PARENT: // 13
		return "GEO_PARENT"
	case Qt__GEO_DIALINGCODE: // 14
		return "GEO_DIALINGCODE"
	case Qt__GEO_CURRENCYCODE: // 15
		return "GEO_CURRENCYCODE"
	case Qt__GEO_CURRENCYSYMBOL: // 16
		return "GEO_CURRENCYSYMBOL"
	case Qt__GEO_NAME: // 17
		return "GEO_NAME"
	case Qt__GEO_ID: // 18
		return "GEO_ID"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__SYSGEOCLASS = int // stdglobal
//
const Qt__GEOCLASS_REGION Qt__SYSGEOCLASS = 14

//
const Qt__GEOCLASS_NATION Qt__SYSGEOCLASS = 16

func SYSGEOCLASSItemName(val int) string {
	switch val {
	case Qt__GEOCLASS_REGION: // 14
		return "GEOCLASS_REGION"
	case Qt__GEOCLASS_NATION: // 16
		return "GEOCLASS_NATION"
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
type Qt__SHSTOCKICONID = int // stdglobal
//
const Qt__SIID_INVALID Qt__SHSTOCKICONID = -1

//
const Qt__SIID_DOCNOASSOC Qt__SHSTOCKICONID = 0

//
const Qt__SIID_DOCASSOC Qt__SHSTOCKICONID = 1

//
const Qt__SIID_APPLICATION Qt__SHSTOCKICONID = 2

//
const Qt__SIID_FOLDER Qt__SHSTOCKICONID = 3

//
const Qt__SIID_FOLDEROPEN Qt__SHSTOCKICONID = 4

//
const Qt__SIID_DRIVE525 Qt__SHSTOCKICONID = 5

//
const Qt__SIID_DRIVE35 Qt__SHSTOCKICONID = 6

//
const Qt__SIID_DRIVERREMOVE Qt__SHSTOCKICONID = 7

//
const Qt__SIID_DRIVERFIXED Qt__SHSTOCKICONID = 8

//
const Qt__SIID_DRIVERNET Qt__SHSTOCKICONID = 9

//
const Qt__SIID_DRIVERNETDISABLE Qt__SHSTOCKICONID = 10

//
const Qt__SIID_DRIVERCD Qt__SHSTOCKICONID = 11

//
const Qt__SIID_DRIVERRAM Qt__SHSTOCKICONID = 12

//
const Qt__SIID_WORLD Qt__SHSTOCKICONID = 13

//
const Qt__SIID_SERVER Qt__SHSTOCKICONID = 15

//
const Qt__SIID_PRINTER Qt__SHSTOCKICONID = 16

//
const Qt__SIID_MYNETWORK Qt__SHSTOCKICONID = 17

//
const Qt__SIID_FIND Qt__SHSTOCKICONID = 22

//
const Qt__SIID_HELP Qt__SHSTOCKICONID = 23

//
const Qt__SIID_SHARE Qt__SHSTOCKICONID = 28

//
const Qt__SIID_LINK Qt__SHSTOCKICONID = 29

//
const Qt__SIID_SLOWFILE Qt__SHSTOCKICONID = 30

//
const Qt__SIID_RECYCLER Qt__SHSTOCKICONID = 31

//
const Qt__SIID_RECYCLERFULL Qt__SHSTOCKICONID = 32

//
const Qt__SIID_MEDIACDAUDIO Qt__SHSTOCKICONID = 40

//
const Qt__SIID_LOCK Qt__SHSTOCKICONID = 47

//
const Qt__SIID_AUTOLIST Qt__SHSTOCKICONID = 49

//
const Qt__SIID_PRINTERNET Qt__SHSTOCKICONID = 50

//
const Qt__SIID_SERVERSHARE Qt__SHSTOCKICONID = 51

//
const Qt__SIID_PRINTERFAX Qt__SHSTOCKICONID = 52

//
const Qt__SIID_PRINTERFAXNET Qt__SHSTOCKICONID = 53

//
const Qt__SIID_PRINTERFILE Qt__SHSTOCKICONID = 54

//
const Qt__SIID_STACK Qt__SHSTOCKICONID = 55

//
const Qt__SIID_MEDIASVCD Qt__SHSTOCKICONID = 56

//
const Qt__SIID_STUFFEDFOLDER Qt__SHSTOCKICONID = 57

//
const Qt__SIID_DRIVEUNKNOWN Qt__SHSTOCKICONID = 58

//
const Qt__SIID_DRIVEDVD Qt__SHSTOCKICONID = 59

//
const Qt__SIID_MEDIADVD Qt__SHSTOCKICONID = 60

//
const Qt__SIID_MEDIADVDRAM Qt__SHSTOCKICONID = 61

//
const Qt__SIID_MEDIADVDRW Qt__SHSTOCKICONID = 62

//
const Qt__SIID_MEDIADVDR Qt__SHSTOCKICONID = 63

//
const Qt__SIID_MEDIADVDROM Qt__SHSTOCKICONID = 64

//
const Qt__SIID_MEDIACDAUDIOPLUS Qt__SHSTOCKICONID = 65

//
const Qt__SIID_MEDIACDRW Qt__SHSTOCKICONID = 66

//
const Qt__SIID_MEDIACDR Qt__SHSTOCKICONID = 67

//
const Qt__SIID_MEDIACDBURN Qt__SHSTOCKICONID = 68

//
const Qt__SIID_MEDIABLANKCD Qt__SHSTOCKICONID = 69

//
const Qt__SIID_MEDIACDROM Qt__SHSTOCKICONID = 70

//
const Qt__SIID_AUDIOFILES Qt__SHSTOCKICONID = 71

//
const Qt__SIID_IMAGEFILES Qt__SHSTOCKICONID = 72

//
const Qt__SIID_VIDEOFILES Qt__SHSTOCKICONID = 73

//
const Qt__SIID_MIXEDFILES Qt__SHSTOCKICONID = 74

//
const Qt__SIID_FOLDERBACK Qt__SHSTOCKICONID = 75

//
const Qt__SIID_FOLDERFRONT Qt__SHSTOCKICONID = 76

//
const Qt__SIID_SHIELD Qt__SHSTOCKICONID = 77

//
const Qt__SIID_WARNING Qt__SHSTOCKICONID = 78

//
const Qt__SIID_INFO Qt__SHSTOCKICONID = 79

//
const Qt__SIID_ERROR Qt__SHSTOCKICONID = 80

//
const Qt__SIID_KEY Qt__SHSTOCKICONID = 81

//
const Qt__SIID_SOFTWARE Qt__SHSTOCKICONID = 82

//
const Qt__SIID_RENAME Qt__SHSTOCKICONID = 83

//
const Qt__SIID_DELETE Qt__SHSTOCKICONID = 84

//
const Qt__SIID_MEDIAAUDIODVD Qt__SHSTOCKICONID = 85

//
const Qt__SIID_MEDIAMOVIEDVD Qt__SHSTOCKICONID = 86

//
const Qt__SIID_MEDIAENHANCEDCD Qt__SHSTOCKICONID = 87

//
const Qt__SIID_MEDIAENHANCEDDVD Qt__SHSTOCKICONID = 88

//
const Qt__SIID_MEDIAHDDVD Qt__SHSTOCKICONID = 89

//
const Qt__SIID_MEDIABLUERAY Qt__SHSTOCKICONID = 90

//
const Qt__SIID_MEDIAVCD Qt__SHSTOCKICONID = 91

//
const Qt__SIID_MEDIADVDPLUSR Qt__SHSTOCKICONID = 92

//
const Qt__SIID_MEDIADVDPLUSRW Qt__SHSTOCKICONID = 93

//
const Qt__SIID_DESKTOPPC Qt__SHSTOCKICONID = 94

//
const Qt__SIID_MOBILEPC Qt__SHSTOCKICONID = 95

//
const Qt__SIID_USERS Qt__SHSTOCKICONID = 96

//
const Qt__SIID_MEDIASMARTMEDIA Qt__SHSTOCKICONID = 97

//
const Qt__SIID_MEDIACOMPACTFLASH Qt__SHSTOCKICONID = 98

//
const Qt__SIID_DEVICECELLPHONE Qt__SHSTOCKICONID = 99

//
const Qt__SIID_DEVICECAMERA Qt__SHSTOCKICONID = 100

//
const Qt__SIID_DEVICEVIDEOCAMERA Qt__SHSTOCKICONID = 101

//
const Qt__SIID_DEVICEAUDIOPLAYER Qt__SHSTOCKICONID = 102

//
const Qt__SIID_NETWORKCONNECT Qt__SHSTOCKICONID = 103

//
const Qt__SIID_INTERNET Qt__SHSTOCKICONID = 104

//
const Qt__SIID_ZIPFILE Qt__SHSTOCKICONID = 105

//
const Qt__SIID_SETTINGS Qt__SHSTOCKICONID = 106

//
const Qt__SIID_DRIVEHDDVD Qt__SHSTOCKICONID = 132

//
const Qt__SIID_DRIVEBD Qt__SHSTOCKICONID = 133

//
const Qt__SIID_MEDIAHDDVDROM Qt__SHSTOCKICONID = 134

//
const Qt__SIID_MEDIAHDDVDR Qt__SHSTOCKICONID = 135

//
const Qt__SIID_MEDIAHDDVDRAM Qt__SHSTOCKICONID = 136

//
const Qt__SIID_MEDIABDROM Qt__SHSTOCKICONID = 137

//
const Qt__SIID_MEDIABDR Qt__SHSTOCKICONID = 138

//
const Qt__SIID_MEDIABDRE Qt__SHSTOCKICONID = 139

//
const Qt__SIID_CLUSTEREDDRIVE Qt__SHSTOCKICONID = 140

//
const Qt__SIID_MAX_ICONS Qt__SHSTOCKICONID = 181

func SHSTOCKICONIDItemName(val int) string {
	switch val {
	case Qt__SIID_INVALID: // -1
		return "SIID_INVALID"
	case Qt__SIID_DOCNOASSOC: // 0
		return "SIID_DOCNOASSOC"
	case Qt__SIID_DOCASSOC: // 1
		return "SIID_DOCASSOC"
	case Qt__SIID_APPLICATION: // 2
		return "SIID_APPLICATION"
	case Qt__SIID_FOLDER: // 3
		return "SIID_FOLDER"
	case Qt__SIID_FOLDEROPEN: // 4
		return "SIID_FOLDEROPEN"
	case Qt__SIID_DRIVE525: // 5
		return "SIID_DRIVE525"
	case Qt__SIID_DRIVE35: // 6
		return "SIID_DRIVE35"
	case Qt__SIID_DRIVERREMOVE: // 7
		return "SIID_DRIVERREMOVE"
	case Qt__SIID_DRIVERFIXED: // 8
		return "SIID_DRIVERFIXED"
	case Qt__SIID_DRIVERNET: // 9
		return "SIID_DRIVERNET"
	case Qt__SIID_DRIVERNETDISABLE: // 10
		return "SIID_DRIVERNETDISABLE"
	case Qt__SIID_DRIVERCD: // 11
		return "SIID_DRIVERCD"
	case Qt__SIID_DRIVERRAM: // 12
		return "SIID_DRIVERRAM"
	case Qt__SIID_WORLD: // 13
		return "SIID_WORLD"
	case Qt__SIID_SERVER: // 15
		return "SIID_SERVER"
	case Qt__SIID_PRINTER: // 16
		return "SIID_PRINTER"
	case Qt__SIID_MYNETWORK: // 17
		return "SIID_MYNETWORK"
	case Qt__SIID_FIND: // 22
		return "SIID_FIND"
	case Qt__SIID_HELP: // 23
		return "SIID_HELP"
	case Qt__SIID_SHARE: // 28
		return "SIID_SHARE"
	case Qt__SIID_LINK: // 29
		return "SIID_LINK"
	case Qt__SIID_SLOWFILE: // 30
		return "SIID_SLOWFILE"
	case Qt__SIID_RECYCLER: // 31
		return "SIID_RECYCLER"
	case Qt__SIID_RECYCLERFULL: // 32
		return "SIID_RECYCLERFULL"
	case Qt__SIID_MEDIACDAUDIO: // 40
		return "SIID_MEDIACDAUDIO"
	case Qt__SIID_LOCK: // 47
		return "SIID_LOCK"
	case Qt__SIID_AUTOLIST: // 49
		return "SIID_AUTOLIST"
	case Qt__SIID_PRINTERNET: // 50
		return "SIID_PRINTERNET"
	case Qt__SIID_SERVERSHARE: // 51
		return "SIID_SERVERSHARE"
	case Qt__SIID_PRINTERFAX: // 52
		return "SIID_PRINTERFAX"
	case Qt__SIID_PRINTERFAXNET: // 53
		return "SIID_PRINTERFAXNET"
	case Qt__SIID_PRINTERFILE: // 54
		return "SIID_PRINTERFILE"
	case Qt__SIID_STACK: // 55
		return "SIID_STACK"
	case Qt__SIID_MEDIASVCD: // 56
		return "SIID_MEDIASVCD"
	case Qt__SIID_STUFFEDFOLDER: // 57
		return "SIID_STUFFEDFOLDER"
	case Qt__SIID_DRIVEUNKNOWN: // 58
		return "SIID_DRIVEUNKNOWN"
	case Qt__SIID_DRIVEDVD: // 59
		return "SIID_DRIVEDVD"
	case Qt__SIID_MEDIADVD: // 60
		return "SIID_MEDIADVD"
	case Qt__SIID_MEDIADVDRAM: // 61
		return "SIID_MEDIADVDRAM"
	case Qt__SIID_MEDIADVDRW: // 62
		return "SIID_MEDIADVDRW"
	case Qt__SIID_MEDIADVDR: // 63
		return "SIID_MEDIADVDR"
	case Qt__SIID_MEDIADVDROM: // 64
		return "SIID_MEDIADVDROM"
	case Qt__SIID_MEDIACDAUDIOPLUS: // 65
		return "SIID_MEDIACDAUDIOPLUS"
	case Qt__SIID_MEDIACDRW: // 66
		return "SIID_MEDIACDRW"
	case Qt__SIID_MEDIACDR: // 67
		return "SIID_MEDIACDR"
	case Qt__SIID_MEDIACDBURN: // 68
		return "SIID_MEDIACDBURN"
	case Qt__SIID_MEDIABLANKCD: // 69
		return "SIID_MEDIABLANKCD"
	case Qt__SIID_MEDIACDROM: // 70
		return "SIID_MEDIACDROM"
	case Qt__SIID_AUDIOFILES: // 71
		return "SIID_AUDIOFILES"
	case Qt__SIID_IMAGEFILES: // 72
		return "SIID_IMAGEFILES"
	case Qt__SIID_VIDEOFILES: // 73
		return "SIID_VIDEOFILES"
	case Qt__SIID_MIXEDFILES: // 74
		return "SIID_MIXEDFILES"
	case Qt__SIID_FOLDERBACK: // 75
		return "SIID_FOLDERBACK"
	case Qt__SIID_FOLDERFRONT: // 76
		return "SIID_FOLDERFRONT"
	case Qt__SIID_SHIELD: // 77
		return "SIID_SHIELD"
	case Qt__SIID_WARNING: // 78
		return "SIID_WARNING"
	case Qt__SIID_INFO: // 79
		return "SIID_INFO"
	case Qt__SIID_ERROR: // 80
		return "SIID_ERROR"
	case Qt__SIID_KEY: // 81
		return "SIID_KEY"
	case Qt__SIID_SOFTWARE: // 82
		return "SIID_SOFTWARE"
	case Qt__SIID_RENAME: // 83
		return "SIID_RENAME"
	case Qt__SIID_DELETE: // 84
		return "SIID_DELETE"
	case Qt__SIID_MEDIAAUDIODVD: // 85
		return "SIID_MEDIAAUDIODVD"
	case Qt__SIID_MEDIAMOVIEDVD: // 86
		return "SIID_MEDIAMOVIEDVD"
	case Qt__SIID_MEDIAENHANCEDCD: // 87
		return "SIID_MEDIAENHANCEDCD"
	case Qt__SIID_MEDIAENHANCEDDVD: // 88
		return "SIID_MEDIAENHANCEDDVD"
	case Qt__SIID_MEDIAHDDVD: // 89
		return "SIID_MEDIAHDDVD"
	case Qt__SIID_MEDIABLUERAY: // 90
		return "SIID_MEDIABLUERAY"
	case Qt__SIID_MEDIAVCD: // 91
		return "SIID_MEDIAVCD"
	case Qt__SIID_MEDIADVDPLUSR: // 92
		return "SIID_MEDIADVDPLUSR"
	case Qt__SIID_MEDIADVDPLUSRW: // 93
		return "SIID_MEDIADVDPLUSRW"
	case Qt__SIID_DESKTOPPC: // 94
		return "SIID_DESKTOPPC"
	case Qt__SIID_MOBILEPC: // 95
		return "SIID_MOBILEPC"
	case Qt__SIID_USERS: // 96
		return "SIID_USERS"
	case Qt__SIID_MEDIASMARTMEDIA: // 97
		return "SIID_MEDIASMARTMEDIA"
	case Qt__SIID_MEDIACOMPACTFLASH: // 98
		return "SIID_MEDIACOMPACTFLASH"
	case Qt__SIID_DEVICECELLPHONE: // 99
		return "SIID_DEVICECELLPHONE"
	case Qt__SIID_DEVICECAMERA: // 100
		return "SIID_DEVICECAMERA"
	case Qt__SIID_DEVICEVIDEOCAMERA: // 101
		return "SIID_DEVICEVIDEOCAMERA"
	case Qt__SIID_DEVICEAUDIOPLAYER: // 102
		return "SIID_DEVICEAUDIOPLAYER"
	case Qt__SIID_NETWORKCONNECT: // 103
		return "SIID_NETWORKCONNECT"
	case Qt__SIID_INTERNET: // 104
		return "SIID_INTERNET"
	case Qt__SIID_ZIPFILE: // 105
		return "SIID_ZIPFILE"
	case Qt__SIID_SETTINGS: // 106
		return "SIID_SETTINGS"
	case Qt__SIID_DRIVEHDDVD: // 132
		return "SIID_DRIVEHDDVD"
	case Qt__SIID_DRIVEBD: // 133
		return "SIID_DRIVEBD"
	case Qt__SIID_MEDIAHDDVDROM: // 134
		return "SIID_MEDIAHDDVDROM"
	case Qt__SIID_MEDIAHDDVDR: // 135
		return "SIID_MEDIAHDDVDR"
	case Qt__SIID_MEDIAHDDVDRAM: // 136
		return "SIID_MEDIAHDDVDRAM"
	case Qt__SIID_MEDIABDROM: // 137
		return "SIID_MEDIABDROM"
	case Qt__SIID_MEDIABDR: // 138
		return "SIID_MEDIABDR"
	case Qt__SIID_MEDIABDRE: // 139
		return "SIID_MEDIABDRE"
	case Qt__SIID_CLUSTEREDDRIVE: // 140
		return "SIID_CLUSTEREDDRIVE"
	case Qt__SIID_MAX_ICONS: // 181
		return "SIID_MAX_ICONS"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagMEMCTX = int // stdglobal
//
const Qt__MEMCTX_TASK Qt__tagMEMCTX = 1

//
const Qt__MEMCTX_SHARED Qt__tagMEMCTX = 2

//
const Qt__MEMCTX_MACSYSTEM Qt__tagMEMCTX = 3

//
const Qt__MEMCTX_UNKNOWN Qt__tagMEMCTX = -1

//
const Qt__MEMCTX_SAME Qt__tagMEMCTX = -2

func tagMEMCTXItemName(val int) string {
	switch val {
	case Qt__MEMCTX_TASK: // 1
		return "MEMCTX_TASK"
	case Qt__MEMCTX_SHARED: // 2
		return "MEMCTX_SHARED"
	case Qt__MEMCTX_MACSYSTEM: // 3
		return "MEMCTX_MACSYSTEM"
	case Qt__MEMCTX_UNKNOWN: // -1
		return "MEMCTX_UNKNOWN"
	case Qt__MEMCTX_SAME: // -2
		return "MEMCTX_SAME"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagCLSCTX = int // stdglobal
//
const Qt__CLSCTX_INPROC_SERVER Qt__tagCLSCTX = 1

//
const Qt__CLSCTX_INPROC_HANDLER Qt__tagCLSCTX = 2

//
const Qt__CLSCTX_LOCAL_SERVER Qt__tagCLSCTX = 4

//
const Qt__CLSCTX_INPROC_SERVER16 Qt__tagCLSCTX = 8

//
const Qt__CLSCTX_REMOTE_SERVER Qt__tagCLSCTX = 16

//
const Qt__CLSCTX_INPROC_HANDLER16 Qt__tagCLSCTX = 32

//
const Qt__CLSCTX_INPROC_SERVERX86 Qt__tagCLSCTX = 64

//
const Qt__CLSCTX_INPROC_HANDLERX86 Qt__tagCLSCTX = 128

//
const Qt__CLSCTX_ESERVER_HANDLER Qt__tagCLSCTX = 256

//
const Qt__CLSCTX_NO_CODE_DOWNLOAD Qt__tagCLSCTX = 1024

//
const Qt__CLSCTX_NO_CUSTOM_MARSHAL Qt__tagCLSCTX = 4096

//
const Qt__CLSCTX_ENABLE_CODE_DOWNLOAD Qt__tagCLSCTX = 8192

//
const Qt__CLSCTX_NO_FAILURE_LOG Qt__tagCLSCTX = 16384

//
const Qt__CLSCTX_DISABLE_AAA Qt__tagCLSCTX = 32768

//
const Qt__CLSCTX_ENABLE_AAA Qt__tagCLSCTX = 65536

//
const Qt__CLSCTX_FROM_DEFAULT_CONTEXT Qt__tagCLSCTX = 131072

func tagCLSCTXItemName(val int) string {
	switch val {
	case Qt__CLSCTX_INPROC_SERVER: // 1
		return "CLSCTX_INPROC_SERVER"
	case Qt__CLSCTX_INPROC_HANDLER: // 2
		return "CLSCTX_INPROC_HANDLER"
	case Qt__CLSCTX_LOCAL_SERVER: // 4
		return "CLSCTX_LOCAL_SERVER"
	case Qt__CLSCTX_INPROC_SERVER16: // 8
		return "CLSCTX_INPROC_SERVER16"
	case Qt__CLSCTX_REMOTE_SERVER: // 16
		return "CLSCTX_REMOTE_SERVER"
	case Qt__CLSCTX_INPROC_HANDLER16: // 32
		return "CLSCTX_INPROC_HANDLER16"
	case Qt__CLSCTX_INPROC_SERVERX86: // 64
		return "CLSCTX_INPROC_SERVERX86"
	case Qt__CLSCTX_INPROC_HANDLERX86: // 128
		return "CLSCTX_INPROC_HANDLERX86"
	case Qt__CLSCTX_ESERVER_HANDLER: // 256
		return "CLSCTX_ESERVER_HANDLER"
	case Qt__CLSCTX_NO_CODE_DOWNLOAD: // 1024
		return "CLSCTX_NO_CODE_DOWNLOAD"
	case Qt__CLSCTX_NO_CUSTOM_MARSHAL: // 4096
		return "CLSCTX_NO_CUSTOM_MARSHAL"
	case Qt__CLSCTX_ENABLE_CODE_DOWNLOAD: // 8192
		return "CLSCTX_ENABLE_CODE_DOWNLOAD"
	case Qt__CLSCTX_NO_FAILURE_LOG: // 16384
		return "CLSCTX_NO_FAILURE_LOG"
	case Qt__CLSCTX_DISABLE_AAA: // 32768
		return "CLSCTX_DISABLE_AAA"
	case Qt__CLSCTX_ENABLE_AAA: // 65536
		return "CLSCTX_ENABLE_AAA"
	case Qt__CLSCTX_FROM_DEFAULT_CONTEXT: // 131072
		return "CLSCTX_FROM_DEFAULT_CONTEXT"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagMSHLFLAGS = int // stdglobal
//
const Qt__MSHLFLAGS_NORMAL Qt__tagMSHLFLAGS = 0

//
const Qt__MSHLFLAGS_TABLESTRONG Qt__tagMSHLFLAGS = 1

//
const Qt__MSHLFLAGS_TABLEWEAK Qt__tagMSHLFLAGS = 2

//
const Qt__MSHLFLAGS_NOPING Qt__tagMSHLFLAGS = 4

func tagMSHLFLAGSItemName(val int) string {
	switch val {
	case Qt__MSHLFLAGS_NORMAL: // 0
		return "MSHLFLAGS_NORMAL"
	case Qt__MSHLFLAGS_TABLESTRONG: // 1
		return "MSHLFLAGS_TABLESTRONG"
	case Qt__MSHLFLAGS_TABLEWEAK: // 2
		return "MSHLFLAGS_TABLEWEAK"
	case Qt__MSHLFLAGS_NOPING: // 4
		return "MSHLFLAGS_NOPING"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagMSHCTX = int // stdglobal
//
const Qt__MSHCTX_LOCAL Qt__tagMSHCTX = 0

//
const Qt__MSHCTX_NOSHAREDMEM Qt__tagMSHCTX = 1

//
const Qt__MSHCTX_DIFFERENTMACHINE Qt__tagMSHCTX = 2

//
const Qt__MSHCTX_INPROC Qt__tagMSHCTX = 3

//
const Qt__MSHCTX_CROSSCTX Qt__tagMSHCTX = 4

func tagMSHCTXItemName(val int) string {
	switch val {
	case Qt__MSHCTX_LOCAL: // 0
		return "MSHCTX_LOCAL"
	case Qt__MSHCTX_NOSHAREDMEM: // 1
		return "MSHCTX_NOSHAREDMEM"
	case Qt__MSHCTX_DIFFERENTMACHINE: // 2
		return "MSHCTX_DIFFERENTMACHINE"
	case Qt__MSHCTX_INPROC: // 3
		return "MSHCTX_INPROC"
	case Qt__MSHCTX_CROSSCTX: // 4
		return "MSHCTX_CROSSCTX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagDVASPECT = int // stdglobal
//
const Qt__DVASPECT_CONTENT Qt__tagDVASPECT = 1

//
const Qt__DVASPECT_THUMBNAIL Qt__tagDVASPECT = 2

//
const Qt__DVASPECT_ICON Qt__tagDVASPECT = 4

//
const Qt__DVASPECT_DOCPRINT Qt__tagDVASPECT = 8

func tagDVASPECTItemName(val int) string {
	switch val {
	case Qt__DVASPECT_CONTENT: // 1
		return "DVASPECT_CONTENT"
	case Qt__DVASPECT_THUMBNAIL: // 2
		return "DVASPECT_THUMBNAIL"
	case Qt__DVASPECT_ICON: // 4
		return "DVASPECT_ICON"
	case Qt__DVASPECT_DOCPRINT: // 8
		return "DVASPECT_DOCPRINT"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSTGC = int // stdglobal
//
const Qt__STGC_DEFAULT Qt__tagSTGC = 0

//
const Qt__STGC_OVERWRITE Qt__tagSTGC = 1

//
const Qt__STGC_ONLYIFCURRENT Qt__tagSTGC = 2

//
const Qt__STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE Qt__tagSTGC = 4

//
const Qt__STGC_CONSOLIDATE Qt__tagSTGC = 8

func tagSTGCItemName(val int) string {
	switch val {
	case Qt__STGC_DEFAULT: // 0
		return "STGC_DEFAULT"
	case Qt__STGC_OVERWRITE: // 1
		return "STGC_OVERWRITE"
	case Qt__STGC_ONLYIFCURRENT: // 2
		return "STGC_ONLYIFCURRENT"
	case Qt__STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE: // 4
		return "STGC_DANGEROUSLYCOMMITMERELYTODISKCACHE"
	case Qt__STGC_CONSOLIDATE: // 8
		return "STGC_CONSOLIDATE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSTGMOVE = int // stdglobal
//
const Qt__STGMOVE_MOVE Qt__tagSTGMOVE = 0

//
const Qt__STGMOVE_COPY Qt__tagSTGMOVE = 1

//
const Qt__STGMOVE_SHALLOWCOPY Qt__tagSTGMOVE = 2

func tagSTGMOVEItemName(val int) string {
	switch val {
	case Qt__STGMOVE_MOVE: // 0
		return "STGMOVE_MOVE"
	case Qt__STGMOVE_COPY: // 1
		return "STGMOVE_COPY"
	case Qt__STGMOVE_SHALLOWCOPY: // 2
		return "STGMOVE_SHALLOWCOPY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSTATFLAG = int // stdglobal
//
const Qt__STATFLAG_DEFAULT Qt__tagSTATFLAG = 0

//
const Qt__STATFLAG_NONAME Qt__tagSTATFLAG = 1

//
const Qt__STATFLAG_NOOPEN Qt__tagSTATFLAG = 2

func tagSTATFLAGItemName(val int) string {
	switch val {
	case Qt__STATFLAG_DEFAULT: // 0
		return "STATFLAG_DEFAULT"
	case Qt__STATFLAG_NONAME: // 1
		return "STATFLAG_NONAME"
	case Qt__STATFLAG_NOOPEN: // 2
		return "STATFLAG_NOOPEN"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__VARENUM = int // stdglobal
//
const Qt__VT_EMPTY Qt__VARENUM = 0

//
const Qt__VT_NULL Qt__VARENUM = 1

//
const Qt__VT_I2 Qt__VARENUM = 2

//
const Qt__VT_I4 Qt__VARENUM = 3

//
const Qt__VT_R4 Qt__VARENUM = 4

//
const Qt__VT_R8 Qt__VARENUM = 5

//
const Qt__VT_CY Qt__VARENUM = 6

//
const Qt__VT_DATE Qt__VARENUM = 7

//
const Qt__VT_BSTR Qt__VARENUM = 8

//
const Qt__VT_DISPATCH Qt__VARENUM = 9

//
const Qt__VT_ERROR Qt__VARENUM = 10

//
const Qt__VT_BOOL Qt__VARENUM = 11

//
const Qt__VT_VARIANT Qt__VARENUM = 12

//
const Qt__VT_UNKNOWN Qt__VARENUM = 13

//
const Qt__VT_DECIMAL Qt__VARENUM = 14

//
const Qt__VT_I1 Qt__VARENUM = 16

//
const Qt__VT_UI1 Qt__VARENUM = 17

//
const Qt__VT_UI2 Qt__VARENUM = 18

//
const Qt__VT_UI4 Qt__VARENUM = 19

//
const Qt__VT_I8 Qt__VARENUM = 20

//
const Qt__VT_UI8 Qt__VARENUM = 21

//
const Qt__VT_INT Qt__VARENUM = 22

//
const Qt__VT_UINT Qt__VARENUM = 23

//
const Qt__VT_VOID Qt__VARENUM = 24

//
const Qt__VT_HRESULT Qt__VARENUM = 25

//
const Qt__VT_PTR Qt__VARENUM = 26

//
const Qt__VT_SAFEARRAY Qt__VARENUM = 27

//
const Qt__VT_CARRAY Qt__VARENUM = 28

//
const Qt__VT_USERDEFINED Qt__VARENUM = 29

//
const Qt__VT_LPSTR Qt__VARENUM = 30

//
const Qt__VT_LPWSTR Qt__VARENUM = 31

//
const Qt__VT_RECORD Qt__VARENUM = 36

//
const Qt__VT_INT_PTR Qt__VARENUM = 37

//
const Qt__VT_UINT_PTR Qt__VARENUM = 38

//
const Qt__VT_FILETIME Qt__VARENUM = 64

//
const Qt__VT_BLOB Qt__VARENUM = 65

//
const Qt__VT_STREAM Qt__VARENUM = 66

//
const Qt__VT_STORAGE Qt__VARENUM = 67

//
const Qt__VT_STREAMED_OBJECT Qt__VARENUM = 68

//
const Qt__VT_STORED_OBJECT Qt__VARENUM = 69

//
const Qt__VT_BLOB_OBJECT Qt__VARENUM = 70

//
const Qt__VT_CF Qt__VARENUM = 71

//
const Qt__VT_CLSID Qt__VARENUM = 72

//
const Qt__VT_VERSIONED_STREAM Qt__VARENUM = 73

//
const Qt__VT_BSTR_BLOB Qt__VARENUM = 4095

//
const Qt__VT_VECTOR Qt__VARENUM = 4096

//
const Qt__VT_ARRAY Qt__VARENUM = 8192

//
const Qt__VT_BYREF Qt__VARENUM = 16384

//
const Qt__VT_RESERVED Qt__VARENUM = 32768

//
const Qt__VT_ILLEGAL Qt__VARENUM = 65535

//
const Qt__VT_ILLEGALMASKED Qt__VARENUM = 4095

//
const Qt__VT_TYPEMASK Qt__VARENUM = 4095

func VARENUMItemName(val int) string {
	switch val {
	case Qt__VT_EMPTY: // 0
		return "VT_EMPTY"
	case Qt__VT_NULL: // 1
		return "VT_NULL"
	case Qt__VT_I2: // 2
		return "VT_I2"
	case Qt__VT_I4: // 3
		return "VT_I4"
	case Qt__VT_R4: // 4
		return "VT_R4"
	case Qt__VT_R8: // 5
		return "VT_R8"
	case Qt__VT_CY: // 6
		return "VT_CY"
	case Qt__VT_DATE: // 7
		return "VT_DATE"
	case Qt__VT_BSTR: // 8
		return "VT_BSTR"
	case Qt__VT_DISPATCH: // 9
		return "VT_DISPATCH"
	case Qt__VT_ERROR: // 10
		return "VT_ERROR"
	case Qt__VT_BOOL: // 11
		return "VT_BOOL"
	case Qt__VT_VARIANT: // 12
		return "VT_VARIANT"
	case Qt__VT_UNKNOWN: // 13
		return "VT_UNKNOWN"
	case Qt__VT_DECIMAL: // 14
		return "VT_DECIMAL"
	case Qt__VT_I1: // 16
		return "VT_I1"
	case Qt__VT_UI1: // 17
		return "VT_UI1"
	case Qt__VT_UI2: // 18
		return "VT_UI2"
	case Qt__VT_UI4: // 19
		return "VT_UI4"
	case Qt__VT_I8: // 20
		return "VT_I8"
	case Qt__VT_UI8: // 21
		return "VT_UI8"
	case Qt__VT_INT: // 22
		return "VT_INT"
	case Qt__VT_UINT: // 23
		return "VT_UINT"
	case Qt__VT_VOID: // 24
		return "VT_VOID"
	case Qt__VT_HRESULT: // 25
		return "VT_HRESULT"
	case Qt__VT_PTR: // 26
		return "VT_PTR"
	case Qt__VT_SAFEARRAY: // 27
		return "VT_SAFEARRAY"
	case Qt__VT_CARRAY: // 28
		return "VT_CARRAY"
	case Qt__VT_USERDEFINED: // 29
		return "VT_USERDEFINED"
	case Qt__VT_LPSTR: // 30
		return "VT_LPSTR"
	case Qt__VT_LPWSTR: // 31
		return "VT_LPWSTR"
	case Qt__VT_RECORD: // 36
		return "VT_RECORD"
	case Qt__VT_INT_PTR: // 37
		return "VT_INT_PTR"
	case Qt__VT_UINT_PTR: // 38
		return "VT_UINT_PTR"
	case Qt__VT_FILETIME: // 64
		return "VT_FILETIME"
	case Qt__VT_BLOB: // 65
		return "VT_BLOB"
	case Qt__VT_STREAM: // 66
		return "VT_STREAM"
	case Qt__VT_STORAGE: // 67
		return "VT_STORAGE"
	case Qt__VT_STREAMED_OBJECT: // 68
		return "VT_STREAMED_OBJECT"
	case Qt__VT_STORED_OBJECT: // 69
		return "VT_STORED_OBJECT"
	case Qt__VT_BLOB_OBJECT: // 70
		return "VT_BLOB_OBJECT"
	case Qt__VT_CF: // 71
		return "VT_CF"
	case Qt__VT_CLSID: // 72
		return "VT_CLSID"
	case Qt__VT_VERSIONED_STREAM: // 73
		return "VT_VERSIONED_STREAM"
	case Qt__VT_BSTR_BLOB: // 4095
		return "VT_BSTR_BLOB,VT_ILLEGALMASKED,VT_TYPEMASK"
	case Qt__VT_VECTOR: // 4096
		return "VT_VECTOR"
	case Qt__VT_ARRAY: // 8192
		return "VT_ARRAY"
	case Qt__VT_BYREF: // 16384
		return "VT_BYREF"
	case Qt__VT_RESERVED: // 32768
		return "VT_RESERVED"
	case Qt__VT_ILLEGAL: // 65535
		return "VT_ILLEGAL"
		// case Qt__VT_ILLEGALMASKED: // 4095
		// return ""
		// case Qt__VT_TYPEMASK: // 4095
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagTYSPEC = int // stdglobal
//
const Qt__TYSPEC_CLSID Qt__tagTYSPEC = 0

//
const Qt__TYSPEC_FILEEXT Qt__tagTYSPEC = 1

//
const Qt__TYSPEC_MIMETYPE Qt__tagTYSPEC = 2

//
const Qt__TYSPEC_PROGID Qt__tagTYSPEC = 3

//
const Qt__TYSPEC_FILENAME Qt__tagTYSPEC = 4

//
const Qt__TYSPEC_PACKAGENAME Qt__tagTYSPEC = 5

//
const Qt__TYSPEC_OBJECTID Qt__tagTYSPEC = 6

func tagTYSPECItemName(val int) string {
	switch val {
	case Qt__TYSPEC_CLSID: // 0
		return "TYSPEC_CLSID"
	case Qt__TYSPEC_FILEEXT: // 1
		return "TYSPEC_FILEEXT"
	case Qt__TYSPEC_MIMETYPE: // 2
		return "TYSPEC_MIMETYPE"
	case Qt__TYSPEC_PROGID: // 3
		return "TYSPEC_PROGID"
	case Qt__TYSPEC_FILENAME: // 4
		return "TYSPEC_FILENAME"
	case Qt__TYSPEC_PACKAGENAME: // 5
		return "TYSPEC_PACKAGENAME"
	case Qt__TYSPEC_OBJECTID: // 6
		return "TYSPEC_OBJECTID"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagEXTCONN = int // stdglobal
//
const Qt__EXTCONN_STRONG Qt__tagEXTCONN = 1

//
const Qt__EXTCONN_WEAK Qt__tagEXTCONN = 2

//
const Qt__EXTCONN_CALLABLE Qt__tagEXTCONN = 4

func tagEXTCONNItemName(val int) string {
	switch val {
	case Qt__EXTCONN_STRONG: // 1
		return "EXTCONN_STRONG"
	case Qt__EXTCONN_WEAK: // 2
		return "EXTCONN_WEAK"
	case Qt__EXTCONN_CALLABLE: // 4
		return "EXTCONN_CALLABLE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagBIND_FLAGS = int // stdglobal
//
const Qt__BIND_MAYBOTHERUSER Qt__tagBIND_FLAGS = 1

//
const Qt__BIND_JUSTTESTEXISTENCE Qt__tagBIND_FLAGS = 2

func tagBIND_FLAGSItemName(val int) string {
	switch val {
	case Qt__BIND_MAYBOTHERUSER: // 1
		return "BIND_MAYBOTHERUSER"
	case Qt__BIND_JUSTTESTEXISTENCE: // 2
		return "BIND_JUSTTESTEXISTENCE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagMKSYS = int // stdglobal
//
const Qt__MKSYS_NONE Qt__tagMKSYS = 0

//
const Qt__MKSYS_GENERICCOMPOSITE Qt__tagMKSYS = 1

//
const Qt__MKSYS_FILEMONIKER Qt__tagMKSYS = 2

//
const Qt__MKSYS_ANTIMONIKER Qt__tagMKSYS = 3

//
const Qt__MKSYS_ITEMMONIKER Qt__tagMKSYS = 4

//
const Qt__MKSYS_POINTERMONIKER Qt__tagMKSYS = 5

//
const Qt__MKSYS_CLASSMONIKER Qt__tagMKSYS = 7

func tagMKSYSItemName(val int) string {
	switch val {
	case Qt__MKSYS_NONE: // 0
		return "MKSYS_NONE"
	case Qt__MKSYS_GENERICCOMPOSITE: // 1
		return "MKSYS_GENERICCOMPOSITE"
	case Qt__MKSYS_FILEMONIKER: // 2
		return "MKSYS_FILEMONIKER"
	case Qt__MKSYS_ANTIMONIKER: // 3
		return "MKSYS_ANTIMONIKER"
	case Qt__MKSYS_ITEMMONIKER: // 4
		return "MKSYS_ITEMMONIKER"
	case Qt__MKSYS_POINTERMONIKER: // 5
		return "MKSYS_POINTERMONIKER"
	case Qt__MKSYS_CLASSMONIKER: // 7
		return "MKSYS_CLASSMONIKER"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagMKREDUCE = int // stdglobal
//
const Qt__MKRREDUCE_ONE Qt__tagMKREDUCE = 196608

//
const Qt__MKRREDUCE_TOUSER Qt__tagMKREDUCE = 131072

//
const Qt__MKRREDUCE_THROUGHUSER Qt__tagMKREDUCE = 65536

//
const Qt__MKRREDUCE_ALL Qt__tagMKREDUCE = 0

func tagMKREDUCEItemName(val int) string {
	switch val {
	case Qt__MKRREDUCE_ONE: // 196608
		return "MKRREDUCE_ONE"
	case Qt__MKRREDUCE_TOUSER: // 131072
		return "MKRREDUCE_TOUSER"
	case Qt__MKRREDUCE_THROUGHUSER: // 65536
		return "MKRREDUCE_THROUGHUSER"
	case Qt__MKRREDUCE_ALL: // 0
		return "MKRREDUCE_ALL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSTGTY = int // stdglobal
//
const Qt__STGTY_STORAGE Qt__tagSTGTY = 1

//
const Qt__STGTY_STREAM Qt__tagSTGTY = 2

//
const Qt__STGTY_LOCKBYTES Qt__tagSTGTY = 3

//
const Qt__STGTY_PROPERTY Qt__tagSTGTY = 4

func tagSTGTYItemName(val int) string {
	switch val {
	case Qt__STGTY_STORAGE: // 1
		return "STGTY_STORAGE"
	case Qt__STGTY_STREAM: // 2
		return "STGTY_STREAM"
	case Qt__STGTY_LOCKBYTES: // 3
		return "STGTY_LOCKBYTES"
	case Qt__STGTY_PROPERTY: // 4
		return "STGTY_PROPERTY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSTREAM_SEEK = int // stdglobal
//
const Qt__STREAM_SEEK_SET Qt__tagSTREAM_SEEK = 0

//
const Qt__STREAM_SEEK_CUR Qt__tagSTREAM_SEEK = 1

//
const Qt__STREAM_SEEK_END Qt__tagSTREAM_SEEK = 2

func tagSTREAM_SEEKItemName(val int) string {
	switch val {
	case Qt__STREAM_SEEK_SET: // 0
		return "STREAM_SEEK_SET"
	case Qt__STREAM_SEEK_CUR: // 1
		return "STREAM_SEEK_CUR"
	case Qt__STREAM_SEEK_END: // 2
		return "STREAM_SEEK_END"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagLOCKTYPE = int // stdglobal
//
const Qt__LOCK_WRITE Qt__tagLOCKTYPE = 1

//
const Qt__LOCK_EXCLUSIVE Qt__tagLOCKTYPE = 2

//
const Qt__LOCK_ONLYONCE Qt__tagLOCKTYPE = 4

func tagLOCKTYPEItemName(val int) string {
	switch val {
	case Qt__LOCK_WRITE: // 1
		return "LOCK_WRITE"
	case Qt__LOCK_EXCLUSIVE: // 2
		return "LOCK_EXCLUSIVE"
	case Qt__LOCK_ONLYONCE: // 4
		return "LOCK_ONLYONCE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagADVF = int // stdglobal
//
const Qt__ADVF_NODATA Qt__tagADVF = 1

//
const Qt__ADVF_PRIMEFIRST Qt__tagADVF = 2

//
const Qt__ADVF_ONLYONCE Qt__tagADVF = 4

//
const Qt__ADVF_DATAONSTOP Qt__tagADVF = 64

//
const Qt__ADVFCACHE_NOHANDLER Qt__tagADVF = 8

//
const Qt__ADVFCACHE_FORCEBUILTIN Qt__tagADVF = 16

//
const Qt__ADVFCACHE_ONSAVE Qt__tagADVF = 32

func tagADVFItemName(val int) string {
	switch val {
	case Qt__ADVF_NODATA: // 1
		return "ADVF_NODATA"
	case Qt__ADVF_PRIMEFIRST: // 2
		return "ADVF_PRIMEFIRST"
	case Qt__ADVF_ONLYONCE: // 4
		return "ADVF_ONLYONCE"
	case Qt__ADVF_DATAONSTOP: // 64
		return "ADVF_DATAONSTOP"
	case Qt__ADVFCACHE_NOHANDLER: // 8
		return "ADVFCACHE_NOHANDLER"
	case Qt__ADVFCACHE_FORCEBUILTIN: // 16
		return "ADVFCACHE_FORCEBUILTIN"
	case Qt__ADVFCACHE_ONSAVE: // 32
		return "ADVFCACHE_ONSAVE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagTYMED = int // stdglobal
//
const Qt__TYMED_HGLOBAL Qt__tagTYMED = 1

//
const Qt__TYMED_FILE Qt__tagTYMED = 2

//
const Qt__TYMED_ISTREAM Qt__tagTYMED = 4

//
const Qt__TYMED_ISTORAGE Qt__tagTYMED = 8

//
const Qt__TYMED_GDI Qt__tagTYMED = 16

//
const Qt__TYMED_MFPICT Qt__tagTYMED = 32

//
const Qt__TYMED_ENHMF Qt__tagTYMED = 64

//
const Qt__TYMED_NULL Qt__tagTYMED = 0

func tagTYMEDItemName(val int) string {
	switch val {
	case Qt__TYMED_HGLOBAL: // 1
		return "TYMED_HGLOBAL"
	case Qt__TYMED_FILE: // 2
		return "TYMED_FILE"
	case Qt__TYMED_ISTREAM: // 4
		return "TYMED_ISTREAM"
	case Qt__TYMED_ISTORAGE: // 8
		return "TYMED_ISTORAGE"
	case Qt__TYMED_GDI: // 16
		return "TYMED_GDI"
	case Qt__TYMED_MFPICT: // 32
		return "TYMED_MFPICT"
	case Qt__TYMED_ENHMF: // 64
		return "TYMED_ENHMF"
	case Qt__TYMED_NULL: // 0
		return "TYMED_NULL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagDATADIR = int // stdglobal
//
const Qt__DATADIR_GET Qt__tagDATADIR = 1

//
const Qt__DATADIR_SET Qt__tagDATADIR = 2

func tagDATADIRItemName(val int) string {
	switch val {
	case Qt__DATADIR_GET: // 1
		return "DATADIR_GET"
	case Qt__DATADIR_SET: // 2
		return "DATADIR_SET"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagCALLTYPE = int // stdglobal
//
const Qt__CALLTYPE_TOPLEVEL Qt__tagCALLTYPE = 1

//
const Qt__CALLTYPE_NESTED Qt__tagCALLTYPE = 2

//
const Qt__CALLTYPE_ASYNC Qt__tagCALLTYPE = 3

//
const Qt__CALLTYPE_TOPLEVEL_CALLPENDING Qt__tagCALLTYPE = 4

//
const Qt__CALLTYPE_ASYNC_CALLPENDING Qt__tagCALLTYPE = 5

func tagCALLTYPEItemName(val int) string {
	switch val {
	case Qt__CALLTYPE_TOPLEVEL: // 1
		return "CALLTYPE_TOPLEVEL"
	case Qt__CALLTYPE_NESTED: // 2
		return "CALLTYPE_NESTED"
	case Qt__CALLTYPE_ASYNC: // 3
		return "CALLTYPE_ASYNC"
	case Qt__CALLTYPE_TOPLEVEL_CALLPENDING: // 4
		return "CALLTYPE_TOPLEVEL_CALLPENDING"
	case Qt__CALLTYPE_ASYNC_CALLPENDING: // 5
		return "CALLTYPE_ASYNC_CALLPENDING"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSERVERCALL = int // stdglobal
//
const Qt__SERVERCALL_ISHANDLED Qt__tagSERVERCALL = 0

//
const Qt__SERVERCALL_REJECTED Qt__tagSERVERCALL = 1

//
const Qt__SERVERCALL_RETRYLATER Qt__tagSERVERCALL = 2

func tagSERVERCALLItemName(val int) string {
	switch val {
	case Qt__SERVERCALL_ISHANDLED: // 0
		return "SERVERCALL_ISHANDLED"
	case Qt__SERVERCALL_REJECTED: // 1
		return "SERVERCALL_REJECTED"
	case Qt__SERVERCALL_RETRYLATER: // 2
		return "SERVERCALL_RETRYLATER"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagPENDINGTYPE = int // stdglobal
//
const Qt__PENDINGTYPE_TOPLEVEL Qt__tagPENDINGTYPE = 1

//
const Qt__PENDINGTYPE_NESTED Qt__tagPENDINGTYPE = 2

func tagPENDINGTYPEItemName(val int) string {
	switch val {
	case Qt__PENDINGTYPE_TOPLEVEL: // 1
		return "PENDINGTYPE_TOPLEVEL"
	case Qt__PENDINGTYPE_NESTED: // 2
		return "PENDINGTYPE_NESTED"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagPENDINGMSG = int // stdglobal
//
const Qt__PENDINGMSG_CANCELCALL Qt__tagPENDINGMSG = 0

//
const Qt__PENDINGMSG_WAITNOPROCESS Qt__tagPENDINGMSG = 1

//
const Qt__PENDINGMSG_WAITDEFPROCESS Qt__tagPENDINGMSG = 2

func tagPENDINGMSGItemName(val int) string {
	switch val {
	case Qt__PENDINGMSG_CANCELCALL: // 0
		return "PENDINGMSG_CANCELCALL"
	case Qt__PENDINGMSG_WAITNOPROCESS: // 1
		return "PENDINGMSG_WAITNOPROCESS"
	case Qt__PENDINGMSG_WAITDEFPROCESS: // 2
		return "PENDINGMSG_WAITDEFPROCESS"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = int // stdglobal
//
const Qt__EOAC_NONE Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 0

//
const Qt__EOAC_MUTUAL_AUTH Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 1

//
const Qt__EOAC_SECURE_REFS Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 2

//
const Qt__EOAC_ACCESS_CONTROL Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 4

//
const Qt__EOAC_APPID Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 8

//
const Qt__EOAC_DYNAMIC Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 16

//
const Qt__EOAC_STATIC_CLOAKING Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 32

//
const Qt__EOAC_DYNAMIC_CLOAKING Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 64

//
const Qt__EOAC_ANY_AUTHORITY Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 128

//
const Qt__EOAC_MAKE_FULLSIC Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 256

//
const Qt__EOAC_REQUIRE_FULLSIC Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 512

//
const Qt__EOAC_AUTO_IMPERSONATE Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 1024

//
const Qt__EOAC_DEFAULT Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 2048

//
const Qt__EOAC_DISABLE_AAA Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 4096

//
const Qt__EOAC_NO_CUSTOM_MARSHAL Qt__tagEOLE_AUTHENTICATION_CAPABILITIES = 8192

func tagEOLE_AUTHENTICATION_CAPABILITIESItemName(val int) string {
	switch val {
	case Qt__EOAC_NONE: // 0
		return "EOAC_NONE"
	case Qt__EOAC_MUTUAL_AUTH: // 1
		return "EOAC_MUTUAL_AUTH"
	case Qt__EOAC_SECURE_REFS: // 2
		return "EOAC_SECURE_REFS"
	case Qt__EOAC_ACCESS_CONTROL: // 4
		return "EOAC_ACCESS_CONTROL"
	case Qt__EOAC_APPID: // 8
		return "EOAC_APPID"
	case Qt__EOAC_DYNAMIC: // 16
		return "EOAC_DYNAMIC"
	case Qt__EOAC_STATIC_CLOAKING: // 32
		return "EOAC_STATIC_CLOAKING"
	case Qt__EOAC_DYNAMIC_CLOAKING: // 64
		return "EOAC_DYNAMIC_CLOAKING"
	case Qt__EOAC_ANY_AUTHORITY: // 128
		return "EOAC_ANY_AUTHORITY"
	case Qt__EOAC_MAKE_FULLSIC: // 256
		return "EOAC_MAKE_FULLSIC"
	case Qt__EOAC_REQUIRE_FULLSIC: // 512
		return "EOAC_REQUIRE_FULLSIC"
	case Qt__EOAC_AUTO_IMPERSONATE: // 1024
		return "EOAC_AUTO_IMPERSONATE"
	case Qt__EOAC_DEFAULT: // 2048
		return "EOAC_DEFAULT"
	case Qt__EOAC_DISABLE_AAA: // 4096
		return "EOAC_DISABLE_AAA"
	case Qt__EOAC_NO_CUSTOM_MARSHAL: // 8192
		return "EOAC_NO_CUSTOM_MARSHAL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagDCOM_CALL_STATE = int // stdglobal
//
const Qt__DCOM_NONE Qt__tagDCOM_CALL_STATE = 0

//
const Qt__DCOM_CALL_COMPLETE Qt__tagDCOM_CALL_STATE = 1

//
const Qt__DCOM_CALL_CANCELED Qt__tagDCOM_CALL_STATE = 2

func tagDCOM_CALL_STATEItemName(val int) string {
	switch val {
	case Qt__DCOM_NONE: // 0
		return "DCOM_NONE"
	case Qt__DCOM_CALL_COMPLETE: // 1
		return "DCOM_CALL_COMPLETE"
	case Qt__DCOM_CALL_CANCELED: // 2
		return "DCOM_CALL_CANCELED"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___APTTYPE = int // stdglobal
//
const Qt__APTTYPE_CURRENT Qt___APTTYPE = -1

//
const Qt__APTTYPE_STA Qt___APTTYPE = 0

//
const Qt__APTTYPE_MTA Qt___APTTYPE = 1

//
const Qt__APTTYPE_NA Qt___APTTYPE = 2

//
const Qt__APTTYPE_MAINSTA Qt___APTTYPE = 3

func _APTTYPEItemName(val int) string {
	switch val {
	case Qt__APTTYPE_CURRENT: // -1
		return "APTTYPE_CURRENT"
	case Qt__APTTYPE_STA: // 0
		return "APTTYPE_STA"
	case Qt__APTTYPE_MTA: // 1
		return "APTTYPE_MTA"
	case Qt__APTTYPE_NA: // 2
		return "APTTYPE_NA"
	case Qt__APTTYPE_MAINSTA: // 3
		return "APTTYPE_MAINSTA"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___APTTYPEQUALIFIER = int // stdglobal
//
const Qt__APTTYPEQUALIFIER_NONE Qt___APTTYPEQUALIFIER = 0

//
const Qt__APTTYPEQUALIFIER_IMPLICIT_MTA Qt___APTTYPEQUALIFIER = 1

//
const Qt__APTTYPEQUALIFIER_NA_ON_MTA Qt___APTTYPEQUALIFIER = 2

//
const Qt__APTTYPEQUALIFIER_NA_ON_STA Qt___APTTYPEQUALIFIER = 3

//
const Qt__APTTYPEQUALIFIER_NA_ON_IMPLICIT_MTA Qt___APTTYPEQUALIFIER = 4

//
const Qt__APTTYPEQUALIFIER_NA_ON_MAINSTA Qt___APTTYPEQUALIFIER = 5

func _APTTYPEQUALIFIERItemName(val int) string {
	switch val {
	case Qt__APTTYPEQUALIFIER_NONE: // 0
		return "APTTYPEQUALIFIER_NONE"
	case Qt__APTTYPEQUALIFIER_IMPLICIT_MTA: // 1
		return "APTTYPEQUALIFIER_IMPLICIT_MTA"
	case Qt__APTTYPEQUALIFIER_NA_ON_MTA: // 2
		return "APTTYPEQUALIFIER_NA_ON_MTA"
	case Qt__APTTYPEQUALIFIER_NA_ON_STA: // 3
		return "APTTYPEQUALIFIER_NA_ON_STA"
	case Qt__APTTYPEQUALIFIER_NA_ON_IMPLICIT_MTA: // 4
		return "APTTYPEQUALIFIER_NA_ON_IMPLICIT_MTA"
	case Qt__APTTYPEQUALIFIER_NA_ON_MAINSTA: // 5
		return "APTTYPEQUALIFIER_NA_ON_MAINSTA"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___THDTYPE = int // stdglobal
//
const Qt__THDTYPE_BLOCKMESSAGES Qt___THDTYPE = 0

//
const Qt__THDTYPE_PROCESSMESSAGES Qt___THDTYPE = 1

func _THDTYPEItemName(val int) string {
	switch val {
	case Qt__THDTYPE_BLOCKMESSAGES: // 0
		return "THDTYPE_BLOCKMESSAGES"
	case Qt__THDTYPE_PROCESSMESSAGES: // 1
		return "THDTYPE_PROCESSMESSAGES"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagGLOBALOPT_PROPERTIES = int // stdglobal
//
const Qt__COMGLB_EXCEPTION_HANDLING Qt__tagGLOBALOPT_PROPERTIES = 1

//
const Qt__COMGLB_APPID Qt__tagGLOBALOPT_PROPERTIES = 2

//
const Qt__COMGLB_RPC_THREADPOOL_SETTING Qt__tagGLOBALOPT_PROPERTIES = 3

func tagGLOBALOPT_PROPERTIESItemName(val int) string {
	switch val {
	case Qt__COMGLB_EXCEPTION_HANDLING: // 1
		return "COMGLB_EXCEPTION_HANDLING"
	case Qt__COMGLB_APPID: // 2
		return "COMGLB_APPID"
	case Qt__COMGLB_RPC_THREADPOOL_SETTING: // 3
		return "COMGLB_RPC_THREADPOOL_SETTING"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagGLOBALOPT_EH_VALUES = int // stdglobal
//
const Qt__COMGLB_EXCEPTION_HANDLE Qt__tagGLOBALOPT_EH_VALUES = 0

//
const Qt__COMGLB_EXCEPTION_DONOT_HANDLE_FATAL Qt__tagGLOBALOPT_EH_VALUES = 1

//
const Qt__COMGLB_EXCEPTION_DONOT_HANDLE Qt__tagGLOBALOPT_EH_VALUES = 1

//
const Qt__COMGLB_EXCEPTION_DONOT_HANDLE_ANY Qt__tagGLOBALOPT_EH_VALUES = 2

func tagGLOBALOPT_EH_VALUESItemName(val int) string {
	switch val {
	case Qt__COMGLB_EXCEPTION_HANDLE: // 0
		return "COMGLB_EXCEPTION_HANDLE"
	case Qt__COMGLB_EXCEPTION_DONOT_HANDLE_FATAL: // 1
		return "COMGLB_EXCEPTION_DONOT_HANDLE_FATAL,COMGLB_EXCEPTION_DONOT_HANDLE"
		// case Qt__COMGLB_EXCEPTION_DONOT_HANDLE: // 1
		// return ""
	case Qt__COMGLB_EXCEPTION_DONOT_HANDLE_ANY: // 2
		return "COMGLB_EXCEPTION_DONOT_HANDLE_ANY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagGLOBALOPT_RPCTP_VALUES = int // stdglobal
//
const Qt__COMGLB_RPC_THREADPOOL_SETTING_DEFAULT_POOL Qt__tagGLOBALOPT_RPCTP_VALUES = 0

//
const Qt__COMGLB_RPC_THREADPOOL_SETTING_PRIVATE_POOL Qt__tagGLOBALOPT_RPCTP_VALUES = 1

func tagGLOBALOPT_RPCTP_VALUESItemName(val int) string {
	switch val {
	case Qt__COMGLB_RPC_THREADPOOL_SETTING_DEFAULT_POOL: // 0
		return "COMGLB_RPC_THREADPOOL_SETTING_DEFAULT_POOL"
	case Qt__COMGLB_RPC_THREADPOOL_SETTING_PRIVATE_POOL: // 1
		return "COMGLB_RPC_THREADPOOL_SETTING_PRIVATE_POOL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagCOINIT = int // stdglobal
//
const Qt__COINIT_APARTMENTTHREADED Qt__tagCOINIT = 2

//
const Qt__COINIT_MULTITHREADED Qt__tagCOINIT = 0

//
const Qt__COINIT_DISABLE_OLE1DDE Qt__tagCOINIT = 4

//
const Qt__COINIT_SPEED_OVER_MEMORY Qt__tagCOINIT = 8

func tagCOINITItemName(val int) string {
	switch val {
	case Qt__COINIT_APARTMENTTHREADED: // 2
		return "COINIT_APARTMENTTHREADED"
	case Qt__COINIT_MULTITHREADED: // 0
		return "COINIT_MULTITHREADED"
	case Qt__COINIT_DISABLE_OLE1DDE: // 4
		return "COINIT_DISABLE_OLE1DDE"
	case Qt__COINIT_SPEED_OVER_MEMORY: // 8
		return "COINIT_SPEED_OVER_MEMORY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagREGCLS = int // stdglobal
//
const Qt__REGCLS_SINGLEUSE Qt__tagREGCLS = 0

//
const Qt__REGCLS_MULTIPLEUSE Qt__tagREGCLS = 1

//
const Qt__REGCLS_MULTI_SEPARATE Qt__tagREGCLS = 2

//
const Qt__REGCLS_SUSPENDED Qt__tagREGCLS = 4

//
const Qt__REGCLS_SURROGATE Qt__tagREGCLS = 8

func tagREGCLSItemName(val int) string {
	switch val {
	case Qt__REGCLS_SINGLEUSE: // 0
		return "REGCLS_SINGLEUSE"
	case Qt__REGCLS_MULTIPLEUSE: // 1
		return "REGCLS_MULTIPLEUSE"
	case Qt__REGCLS_MULTI_SEPARATE: // 2
		return "REGCLS_MULTI_SEPARATE"
	case Qt__REGCLS_SUSPENDED: // 4
		return "REGCLS_SUSPENDED"
	case Qt__REGCLS_SURROGATE: // 8
		return "REGCLS_SURROGATE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagCOWAIT_FLAGS = int // stdglobal
//
const Qt__COWAIT_WAITALL Qt__tagCOWAIT_FLAGS = 1

//
const Qt__COWAIT_ALERTABLE Qt__tagCOWAIT_FLAGS = 2

//
const Qt__COWAIT_INPUTAVAILABLE Qt__tagCOWAIT_FLAGS = 4

func tagCOWAIT_FLAGSItemName(val int) string {
	switch val {
	case Qt__COWAIT_WAITALL: // 1
		return "COWAIT_WAITALL"
	case Qt__COWAIT_ALERTABLE: // 2
		return "COWAIT_ALERTABLE"
	case Qt__COWAIT_INPUTAVAILABLE: // 4
		return "COWAIT_INPUTAVAILABLE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLERENDER = int // stdglobal
//
const Qt__OLERENDER_NONE Qt__tagOLERENDER = 0

//
const Qt__OLERENDER_DRAW Qt__tagOLERENDER = 1

//
const Qt__OLERENDER_FORMAT Qt__tagOLERENDER = 2

//
const Qt__OLERENDER_ASIS Qt__tagOLERENDER = 3

func tagOLERENDERItemName(val int) string {
	switch val {
	case Qt__OLERENDER_NONE: // 0
		return "OLERENDER_NONE"
	case Qt__OLERENDER_DRAW: // 1
		return "OLERENDER_DRAW"
	case Qt__OLERENDER_FORMAT: // 2
		return "OLERENDER_FORMAT"
	case Qt__OLERENDER_ASIS: // 3
		return "OLERENDER_ASIS"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagBINDSPEED = int // stdglobal
//
const Qt__BINDSPEED_INDEFINITE Qt__tagBINDSPEED = 1

//
const Qt__BINDSPEED_MODERATE Qt__tagBINDSPEED = 2

//
const Qt__BINDSPEED_IMMEDIATE Qt__tagBINDSPEED = 3

func tagBINDSPEEDItemName(val int) string {
	switch val {
	case Qt__BINDSPEED_INDEFINITE: // 1
		return "BINDSPEED_INDEFINITE"
	case Qt__BINDSPEED_MODERATE: // 2
		return "BINDSPEED_MODERATE"
	case Qt__BINDSPEED_IMMEDIATE: // 3
		return "BINDSPEED_IMMEDIATE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLECONTF = int // stdglobal
//
const Qt__OLECONTF_EMBEDDINGS Qt__tagOLECONTF = 1

//
const Qt__OLECONTF_LINKS Qt__tagOLECONTF = 2

//
const Qt__OLECONTF_OTHERS Qt__tagOLECONTF = 4

//
const Qt__OLECONTF_ONLYUSER Qt__tagOLECONTF = 8

//
const Qt__OLECONTF_ONLYIFRUNNING Qt__tagOLECONTF = 16

func tagOLECONTFItemName(val int) string {
	switch val {
	case Qt__OLECONTF_EMBEDDINGS: // 1
		return "OLECONTF_EMBEDDINGS"
	case Qt__OLECONTF_LINKS: // 2
		return "OLECONTF_LINKS"
	case Qt__OLECONTF_OTHERS: // 4
		return "OLECONTF_OTHERS"
	case Qt__OLECONTF_ONLYUSER: // 8
		return "OLECONTF_ONLYUSER"
	case Qt__OLECONTF_ONLYIFRUNNING: // 16
		return "OLECONTF_ONLYIFRUNNING"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLEUPDATE = int // stdglobal
//
const Qt__OLEUPDATE_ALWAYS Qt__tagOLEUPDATE = 1

//
const Qt__OLEUPDATE_ONCALL Qt__tagOLEUPDATE = 3

func tagOLEUPDATEItemName(val int) string {
	switch val {
	case Qt__OLEUPDATE_ALWAYS: // 1
		return "OLEUPDATE_ALWAYS"
	case Qt__OLEUPDATE_ONCALL: // 3
		return "OLEUPDATE_ONCALL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLELINKBIND = int // stdglobal
//
const Qt__OLELINKBIND_EVENIFCLASSDIFF Qt__tagOLELINKBIND = 1

func tagOLELINKBINDItemName(val int) string {
	switch val {
	case Qt__OLELINKBIND_EVENIFCLASSDIFF: // 1
		return "OLELINKBIND_EVENIFCLASSDIFF"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagDISCARDCACHE = int // stdglobal
//
const Qt__DISCARDCACHE_SAVEIFDIRTY Qt__tagDISCARDCACHE = 0

//
const Qt__DISCARDCACHE_NOSAVE Qt__tagDISCARDCACHE = 1

func tagDISCARDCACHEItemName(val int) string {
	switch val {
	case Qt__DISCARDCACHE_SAVEIFDIRTY: // 0
		return "DISCARDCACHE_SAVEIFDIRTY"
	case Qt__DISCARDCACHE_NOSAVE: // 1
		return "DISCARDCACHE_NOSAVE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLEVERBATTRIB = int // stdglobal
//
const Qt__OLEVERBATTRIB_NEVERDIRTIES Qt__tagOLEVERBATTRIB = 1

//
const Qt__OLEVERBATTRIB_ONCONTAINERMENU Qt__tagOLEVERBATTRIB = 2

func tagOLEVERBATTRIBItemName(val int) string {
	switch val {
	case Qt__OLEVERBATTRIB_NEVERDIRTIES: // 1
		return "OLEVERBATTRIB_NEVERDIRTIES"
	case Qt__OLEVERBATTRIB_ONCONTAINERMENU: // 2
		return "OLEVERBATTRIB_ONCONTAINERMENU"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLEGETMONIKER = int // stdglobal
//
const Qt__OLEGETMONIKER_ONLYIFTHERE Qt__tagOLEGETMONIKER = 1

//
const Qt__OLEGETMONIKER_FORCEASSIGN Qt__tagOLEGETMONIKER = 2

//
const Qt__OLEGETMONIKER_UNASSIGN Qt__tagOLEGETMONIKER = 3

//
const Qt__OLEGETMONIKER_TEMPFORUSER Qt__tagOLEGETMONIKER = 4

func tagOLEGETMONIKERItemName(val int) string {
	switch val {
	case Qt__OLEGETMONIKER_ONLYIFTHERE: // 1
		return "OLEGETMONIKER_ONLYIFTHERE"
	case Qt__OLEGETMONIKER_FORCEASSIGN: // 2
		return "OLEGETMONIKER_FORCEASSIGN"
	case Qt__OLEGETMONIKER_UNASSIGN: // 3
		return "OLEGETMONIKER_UNASSIGN"
	case Qt__OLEGETMONIKER_TEMPFORUSER: // 4
		return "OLEGETMONIKER_TEMPFORUSER"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLEWHICHMK = int // stdglobal
//
const Qt__OLEWHICHMK_CONTAINER Qt__tagOLEWHICHMK = 1

//
const Qt__OLEWHICHMK_OBJREL Qt__tagOLEWHICHMK = 2

//
const Qt__OLEWHICHMK_OBJFULL Qt__tagOLEWHICHMK = 3

func tagOLEWHICHMKItemName(val int) string {
	switch val {
	case Qt__OLEWHICHMK_CONTAINER: // 1
		return "OLEWHICHMK_CONTAINER"
	case Qt__OLEWHICHMK_OBJREL: // 2
		return "OLEWHICHMK_OBJREL"
	case Qt__OLEWHICHMK_OBJFULL: // 3
		return "OLEWHICHMK_OBJFULL"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagUSERCLASSTYPE = int // stdglobal
//
const Qt__USERCLASSTYPE_FULL Qt__tagUSERCLASSTYPE = 1

//
const Qt__USERCLASSTYPE_SHORT Qt__tagUSERCLASSTYPE = 2

//
const Qt__USERCLASSTYPE_APPNAME Qt__tagUSERCLASSTYPE = 3

func tagUSERCLASSTYPEItemName(val int) string {
	switch val {
	case Qt__USERCLASSTYPE_FULL: // 1
		return "USERCLASSTYPE_FULL"
	case Qt__USERCLASSTYPE_SHORT: // 2
		return "USERCLASSTYPE_SHORT"
	case Qt__USERCLASSTYPE_APPNAME: // 3
		return "USERCLASSTYPE_APPNAME"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLEMISC = int // stdglobal
//
const Qt__OLEMISC_RECOMPOSEONRESIZE Qt__tagOLEMISC = 1

//
const Qt__OLEMISC_ONLYICONIC Qt__tagOLEMISC = 2

//
const Qt__OLEMISC_INSERTNOTREPLACE Qt__tagOLEMISC = 4

//
const Qt__OLEMISC_STATIC Qt__tagOLEMISC = 8

//
const Qt__OLEMISC_CANTLINKINSIDE Qt__tagOLEMISC = 16

//
const Qt__OLEMISC_CANLINKBYOLE1 Qt__tagOLEMISC = 32

//
const Qt__OLEMISC_ISLINKOBJECT Qt__tagOLEMISC = 64

//
const Qt__OLEMISC_INSIDEOUT Qt__tagOLEMISC = 128

//
const Qt__OLEMISC_ACTIVATEWHENVISIBLE Qt__tagOLEMISC = 256

//
const Qt__OLEMISC_RENDERINGISDEVICEINDEPENDENT Qt__tagOLEMISC = 512

//
const Qt__OLEMISC_INVISIBLEATRUNTIME Qt__tagOLEMISC = 1024

//
const Qt__OLEMISC_ALWAYSRUN Qt__tagOLEMISC = 2048

//
const Qt__OLEMISC_ACTSLIKEBUTTON Qt__tagOLEMISC = 4096

//
const Qt__OLEMISC_ACTSLIKELABEL Qt__tagOLEMISC = 8192

//
const Qt__OLEMISC_NOUIACTIVATE Qt__tagOLEMISC = 16384

//
const Qt__OLEMISC_ALIGNABLE Qt__tagOLEMISC = 32768

//
const Qt__OLEMISC_SIMPLEFRAME Qt__tagOLEMISC = 65536

//
const Qt__OLEMISC_SETCLIENTSITEFIRST Qt__tagOLEMISC = 131072

//
const Qt__OLEMISC_IMEMODE Qt__tagOLEMISC = 262144

//
const Qt__OLEMISC_IGNOREACTIVATEWHENVISIBLE Qt__tagOLEMISC = 524288

//
const Qt__OLEMISC_WANTSTOMENUMERGE Qt__tagOLEMISC = 1048576

//
const Qt__OLEMISC_SUPPORTSMULTILEVELUNDO Qt__tagOLEMISC = 2097152

func tagOLEMISCItemName(val int) string {
	switch val {
	case Qt__OLEMISC_RECOMPOSEONRESIZE: // 1
		return "OLEMISC_RECOMPOSEONRESIZE"
	case Qt__OLEMISC_ONLYICONIC: // 2
		return "OLEMISC_ONLYICONIC"
	case Qt__OLEMISC_INSERTNOTREPLACE: // 4
		return "OLEMISC_INSERTNOTREPLACE"
	case Qt__OLEMISC_STATIC: // 8
		return "OLEMISC_STATIC"
	case Qt__OLEMISC_CANTLINKINSIDE: // 16
		return "OLEMISC_CANTLINKINSIDE"
	case Qt__OLEMISC_CANLINKBYOLE1: // 32
		return "OLEMISC_CANLINKBYOLE1"
	case Qt__OLEMISC_ISLINKOBJECT: // 64
		return "OLEMISC_ISLINKOBJECT"
	case Qt__OLEMISC_INSIDEOUT: // 128
		return "OLEMISC_INSIDEOUT"
	case Qt__OLEMISC_ACTIVATEWHENVISIBLE: // 256
		return "OLEMISC_ACTIVATEWHENVISIBLE"
	case Qt__OLEMISC_RENDERINGISDEVICEINDEPENDENT: // 512
		return "OLEMISC_RENDERINGISDEVICEINDEPENDENT"
	case Qt__OLEMISC_INVISIBLEATRUNTIME: // 1024
		return "OLEMISC_INVISIBLEATRUNTIME"
	case Qt__OLEMISC_ALWAYSRUN: // 2048
		return "OLEMISC_ALWAYSRUN"
	case Qt__OLEMISC_ACTSLIKEBUTTON: // 4096
		return "OLEMISC_ACTSLIKEBUTTON"
	case Qt__OLEMISC_ACTSLIKELABEL: // 8192
		return "OLEMISC_ACTSLIKELABEL"
	case Qt__OLEMISC_NOUIACTIVATE: // 16384
		return "OLEMISC_NOUIACTIVATE"
	case Qt__OLEMISC_ALIGNABLE: // 32768
		return "OLEMISC_ALIGNABLE"
	case Qt__OLEMISC_SIMPLEFRAME: // 65536
		return "OLEMISC_SIMPLEFRAME"
	case Qt__OLEMISC_SETCLIENTSITEFIRST: // 131072
		return "OLEMISC_SETCLIENTSITEFIRST"
	case Qt__OLEMISC_IMEMODE: // 262144
		return "OLEMISC_IMEMODE"
	case Qt__OLEMISC_IGNOREACTIVATEWHENVISIBLE: // 524288
		return "OLEMISC_IGNOREACTIVATEWHENVISIBLE"
	case Qt__OLEMISC_WANTSTOMENUMERGE: // 1048576
		return "OLEMISC_WANTSTOMENUMERGE"
	case Qt__OLEMISC_SUPPORTSMULTILEVELUNDO: // 2097152
		return "OLEMISC_SUPPORTSMULTILEVELUNDO"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagOLECLOSE = int // stdglobal
//
const Qt__OLECLOSE_SAVEIFDIRTY Qt__tagOLECLOSE = 0

//
const Qt__OLECLOSE_NOSAVE Qt__tagOLECLOSE = 1

//
const Qt__OLECLOSE_PROMPTSAVE Qt__tagOLECLOSE = 2

func tagOLECLOSEItemName(val int) string {
	switch val {
	case Qt__OLECLOSE_SAVEIFDIRTY: // 0
		return "OLECLOSE_SAVEIFDIRTY"
	case Qt__OLECLOSE_NOSAVE: // 1
		return "OLECLOSE_NOSAVE"
	case Qt__OLECLOSE_PROMPTSAVE: // 2
		return "OLECLOSE_PROMPTSAVE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSF_TYPE = int // stdglobal
//
const Qt__SF_ERROR Qt__tagSF_TYPE = 10

//
const Qt__SF_I1 Qt__tagSF_TYPE = 16

//
const Qt__SF_I2 Qt__tagSF_TYPE = 2

//
const Qt__SF_I4 Qt__tagSF_TYPE = 3

//
const Qt__SF_I8 Qt__tagSF_TYPE = 20

//
const Qt__SF_BSTR Qt__tagSF_TYPE = 8

//
const Qt__SF_UNKNOWN Qt__tagSF_TYPE = 13

//
const Qt__SF_DISPATCH Qt__tagSF_TYPE = 9

//
const Qt__SF_VARIANT Qt__tagSF_TYPE = 12

//
const Qt__SF_RECORD Qt__tagSF_TYPE = 36

//
const Qt__SF_HAVEIID Qt__tagSF_TYPE = 32781

func tagSF_TYPEItemName(val int) string {
	switch val {
	case Qt__SF_ERROR: // 10
		return "SF_ERROR"
	case Qt__SF_I1: // 16
		return "SF_I1"
	case Qt__SF_I2: // 2
		return "SF_I2"
	case Qt__SF_I4: // 3
		return "SF_I4"
	case Qt__SF_I8: // 20
		return "SF_I8"
	case Qt__SF_BSTR: // 8
		return "SF_BSTR"
	case Qt__SF_UNKNOWN: // 13
		return "SF_UNKNOWN"
	case Qt__SF_DISPATCH: // 9
		return "SF_DISPATCH"
	case Qt__SF_VARIANT: // 12
		return "SF_VARIANT"
	case Qt__SF_RECORD: // 36
		return "SF_RECORD"
	case Qt__SF_HAVEIID: // 32781
		return "SF_HAVEIID"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagTYPEKIND = int // stdglobal
//
const Qt__TKIND_ENUM Qt__tagTYPEKIND = 0

//
const Qt__TKIND_RECORD Qt__tagTYPEKIND = 1

//
const Qt__TKIND_MODULE Qt__tagTYPEKIND = 2

//
const Qt__TKIND_INTERFACE Qt__tagTYPEKIND = 3

//
const Qt__TKIND_DISPATCH Qt__tagTYPEKIND = 4

//
const Qt__TKIND_COCLASS Qt__tagTYPEKIND = 5

//
const Qt__TKIND_ALIAS Qt__tagTYPEKIND = 6

//
const Qt__TKIND_UNION Qt__tagTYPEKIND = 7

//
const Qt__TKIND_MAX Qt__tagTYPEKIND = 8

func tagTYPEKINDItemName(val int) string {
	switch val {
	case Qt__TKIND_ENUM: // 0
		return "TKIND_ENUM"
	case Qt__TKIND_RECORD: // 1
		return "TKIND_RECORD"
	case Qt__TKIND_MODULE: // 2
		return "TKIND_MODULE"
	case Qt__TKIND_INTERFACE: // 3
		return "TKIND_INTERFACE"
	case Qt__TKIND_DISPATCH: // 4
		return "TKIND_DISPATCH"
	case Qt__TKIND_COCLASS: // 5
		return "TKIND_COCLASS"
	case Qt__TKIND_ALIAS: // 6
		return "TKIND_ALIAS"
	case Qt__TKIND_UNION: // 7
		return "TKIND_UNION"
	case Qt__TKIND_MAX: // 8
		return "TKIND_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagCALLCONV = int // stdglobal
//
const Qt__CC_FASTCALL Qt__tagCALLCONV = 0

//
const Qt__CC_CDECL Qt__tagCALLCONV = 1

//
const Qt__CC_MSCPASCAL Qt__tagCALLCONV = 2

//
const Qt__CC_PASCAL Qt__tagCALLCONV = 2

//
const Qt__CC_MACPASCAL Qt__tagCALLCONV = 3

//
const Qt__CC_STDCALL Qt__tagCALLCONV = 4

//
const Qt__CC_FPFASTCALL Qt__tagCALLCONV = 5

//
const Qt__CC_SYSCALL Qt__tagCALLCONV = 6

//
const Qt__CC_MPWCDECL Qt__tagCALLCONV = 7

//
const Qt__CC_MPWPASCAL Qt__tagCALLCONV = 8

//
const Qt__CC_MAX Qt__tagCALLCONV = 9

func tagCALLCONVItemName(val int) string {
	switch val {
	case Qt__CC_FASTCALL: // 0
		return "CC_FASTCALL"
	case Qt__CC_CDECL: // 1
		return "CC_CDECL"
	case Qt__CC_MSCPASCAL: // 2
		return "CC_MSCPASCAL,CC_PASCAL"
		// case Qt__CC_PASCAL: // 2
		// return ""
	case Qt__CC_MACPASCAL: // 3
		return "CC_MACPASCAL"
	case Qt__CC_STDCALL: // 4
		return "CC_STDCALL"
	case Qt__CC_FPFASTCALL: // 5
		return "CC_FPFASTCALL"
	case Qt__CC_SYSCALL: // 6
		return "CC_SYSCALL"
	case Qt__CC_MPWCDECL: // 7
		return "CC_MPWCDECL"
	case Qt__CC_MPWPASCAL: // 8
		return "CC_MPWPASCAL"
	case Qt__CC_MAX: // 9
		return "CC_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagFUNCKIND = int // stdglobal
//
const Qt__FUNC_VIRTUAL Qt__tagFUNCKIND = 0

//
const Qt__FUNC_PUREVIRTUAL Qt__tagFUNCKIND = 1

//
const Qt__FUNC_NONVIRTUAL Qt__tagFUNCKIND = 2

//
const Qt__FUNC_STATIC Qt__tagFUNCKIND = 3

//
const Qt__FUNC_DISPATCH Qt__tagFUNCKIND = 4

func tagFUNCKINDItemName(val int) string {
	switch val {
	case Qt__FUNC_VIRTUAL: // 0
		return "FUNC_VIRTUAL"
	case Qt__FUNC_PUREVIRTUAL: // 1
		return "FUNC_PUREVIRTUAL"
	case Qt__FUNC_NONVIRTUAL: // 2
		return "FUNC_NONVIRTUAL"
	case Qt__FUNC_STATIC: // 3
		return "FUNC_STATIC"
	case Qt__FUNC_DISPATCH: // 4
		return "FUNC_DISPATCH"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagINVOKEKIND = int // stdglobal
//
const Qt__INVOKE_FUNC Qt__tagINVOKEKIND = 1

//
const Qt__INVOKE_PROPERTYGET Qt__tagINVOKEKIND = 2

//
const Qt__INVOKE_PROPERTYPUT Qt__tagINVOKEKIND = 4

//
const Qt__INVOKE_PROPERTYPUTREF Qt__tagINVOKEKIND = 8

func tagINVOKEKINDItemName(val int) string {
	switch val {
	case Qt__INVOKE_FUNC: // 1
		return "INVOKE_FUNC"
	case Qt__INVOKE_PROPERTYGET: // 2
		return "INVOKE_PROPERTYGET"
	case Qt__INVOKE_PROPERTYPUT: // 4
		return "INVOKE_PROPERTYPUT"
	case Qt__INVOKE_PROPERTYPUTREF: // 8
		return "INVOKE_PROPERTYPUTREF"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagVARKIND = int // stdglobal
//
const Qt__VAR_PERINSTANCE Qt__tagVARKIND = 0

//
const Qt__VAR_STATIC Qt__tagVARKIND = 1

//
const Qt__VAR_CONST Qt__tagVARKIND = 2

//
const Qt__VAR_DISPATCH Qt__tagVARKIND = 3

func tagVARKINDItemName(val int) string {
	switch val {
	case Qt__VAR_PERINSTANCE: // 0
		return "VAR_PERINSTANCE"
	case Qt__VAR_STATIC: // 1
		return "VAR_STATIC"
	case Qt__VAR_CONST: // 2
		return "VAR_CONST"
	case Qt__VAR_DISPATCH: // 3
		return "VAR_DISPATCH"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagTYPEFLAGS = int // stdglobal
//
const Qt__TYPEFLAG_FAPPOBJECT Qt__tagTYPEFLAGS = 1

//
const Qt__TYPEFLAG_FCANCREATE Qt__tagTYPEFLAGS = 2

//
const Qt__TYPEFLAG_FLICENSED Qt__tagTYPEFLAGS = 4

//
const Qt__TYPEFLAG_FPREDECLID Qt__tagTYPEFLAGS = 8

//
const Qt__TYPEFLAG_FHIDDEN Qt__tagTYPEFLAGS = 16

//
const Qt__TYPEFLAG_FCONTROL Qt__tagTYPEFLAGS = 32

//
const Qt__TYPEFLAG_FDUAL Qt__tagTYPEFLAGS = 64

//
const Qt__TYPEFLAG_FNONEXTENSIBLE Qt__tagTYPEFLAGS = 128

//
const Qt__TYPEFLAG_FOLEAUTOMATION Qt__tagTYPEFLAGS = 256

//
const Qt__TYPEFLAG_FRESTRICTED Qt__tagTYPEFLAGS = 512

//
const Qt__TYPEFLAG_FAGGREGATABLE Qt__tagTYPEFLAGS = 1024

//
const Qt__TYPEFLAG_FREPLACEABLE Qt__tagTYPEFLAGS = 2048

//
const Qt__TYPEFLAG_FDISPATCHABLE Qt__tagTYPEFLAGS = 4096

//
const Qt__TYPEFLAG_FREVERSEBIND Qt__tagTYPEFLAGS = 8192

//
const Qt__TYPEFLAG_FPROXY Qt__tagTYPEFLAGS = 16384

func tagTYPEFLAGSItemName(val int) string {
	switch val {
	case Qt__TYPEFLAG_FAPPOBJECT: // 1
		return "TYPEFLAG_FAPPOBJECT"
	case Qt__TYPEFLAG_FCANCREATE: // 2
		return "TYPEFLAG_FCANCREATE"
	case Qt__TYPEFLAG_FLICENSED: // 4
		return "TYPEFLAG_FLICENSED"
	case Qt__TYPEFLAG_FPREDECLID: // 8
		return "TYPEFLAG_FPREDECLID"
	case Qt__TYPEFLAG_FHIDDEN: // 16
		return "TYPEFLAG_FHIDDEN"
	case Qt__TYPEFLAG_FCONTROL: // 32
		return "TYPEFLAG_FCONTROL"
	case Qt__TYPEFLAG_FDUAL: // 64
		return "TYPEFLAG_FDUAL"
	case Qt__TYPEFLAG_FNONEXTENSIBLE: // 128
		return "TYPEFLAG_FNONEXTENSIBLE"
	case Qt__TYPEFLAG_FOLEAUTOMATION: // 256
		return "TYPEFLAG_FOLEAUTOMATION"
	case Qt__TYPEFLAG_FRESTRICTED: // 512
		return "TYPEFLAG_FRESTRICTED"
	case Qt__TYPEFLAG_FAGGREGATABLE: // 1024
		return "TYPEFLAG_FAGGREGATABLE"
	case Qt__TYPEFLAG_FREPLACEABLE: // 2048
		return "TYPEFLAG_FREPLACEABLE"
	case Qt__TYPEFLAG_FDISPATCHABLE: // 4096
		return "TYPEFLAG_FDISPATCHABLE"
	case Qt__TYPEFLAG_FREVERSEBIND: // 8192
		return "TYPEFLAG_FREVERSEBIND"
	case Qt__TYPEFLAG_FPROXY: // 16384
		return "TYPEFLAG_FPROXY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagFUNCFLAGS = int // stdglobal
//
const Qt__FUNCFLAG_FRESTRICTED Qt__tagFUNCFLAGS = 1

//
const Qt__FUNCFLAG_FSOURCE Qt__tagFUNCFLAGS = 2

//
const Qt__FUNCFLAG_FBINDABLE Qt__tagFUNCFLAGS = 4

//
const Qt__FUNCFLAG_FREQUESTEDIT Qt__tagFUNCFLAGS = 8

//
const Qt__FUNCFLAG_FDISPLAYBIND Qt__tagFUNCFLAGS = 16

//
const Qt__FUNCFLAG_FDEFAULTBIND Qt__tagFUNCFLAGS = 32

//
const Qt__FUNCFLAG_FHIDDEN Qt__tagFUNCFLAGS = 64

//
const Qt__FUNCFLAG_FUSESGETLASTERROR Qt__tagFUNCFLAGS = 128

//
const Qt__FUNCFLAG_FDEFAULTCOLLELEM Qt__tagFUNCFLAGS = 256

//
const Qt__FUNCFLAG_FUIDEFAULT Qt__tagFUNCFLAGS = 512

//
const Qt__FUNCFLAG_FNONBROWSABLE Qt__tagFUNCFLAGS = 1024

//
const Qt__FUNCFLAG_FREPLACEABLE Qt__tagFUNCFLAGS = 2048

//
const Qt__FUNCFLAG_FIMMEDIATEBIND Qt__tagFUNCFLAGS = 4096

func tagFUNCFLAGSItemName(val int) string {
	switch val {
	case Qt__FUNCFLAG_FRESTRICTED: // 1
		return "FUNCFLAG_FRESTRICTED"
	case Qt__FUNCFLAG_FSOURCE: // 2
		return "FUNCFLAG_FSOURCE"
	case Qt__FUNCFLAG_FBINDABLE: // 4
		return "FUNCFLAG_FBINDABLE"
	case Qt__FUNCFLAG_FREQUESTEDIT: // 8
		return "FUNCFLAG_FREQUESTEDIT"
	case Qt__FUNCFLAG_FDISPLAYBIND: // 16
		return "FUNCFLAG_FDISPLAYBIND"
	case Qt__FUNCFLAG_FDEFAULTBIND: // 32
		return "FUNCFLAG_FDEFAULTBIND"
	case Qt__FUNCFLAG_FHIDDEN: // 64
		return "FUNCFLAG_FHIDDEN"
	case Qt__FUNCFLAG_FUSESGETLASTERROR: // 128
		return "FUNCFLAG_FUSESGETLASTERROR"
	case Qt__FUNCFLAG_FDEFAULTCOLLELEM: // 256
		return "FUNCFLAG_FDEFAULTCOLLELEM"
	case Qt__FUNCFLAG_FUIDEFAULT: // 512
		return "FUNCFLAG_FUIDEFAULT"
	case Qt__FUNCFLAG_FNONBROWSABLE: // 1024
		return "FUNCFLAG_FNONBROWSABLE"
	case Qt__FUNCFLAG_FREPLACEABLE: // 2048
		return "FUNCFLAG_FREPLACEABLE"
	case Qt__FUNCFLAG_FIMMEDIATEBIND: // 4096
		return "FUNCFLAG_FIMMEDIATEBIND"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagVARFLAGS = int // stdglobal
//
const Qt__VARFLAG_FREADONLY Qt__tagVARFLAGS = 1

//
const Qt__VARFLAG_FSOURCE Qt__tagVARFLAGS = 2

//
const Qt__VARFLAG_FBINDABLE Qt__tagVARFLAGS = 4

//
const Qt__VARFLAG_FREQUESTEDIT Qt__tagVARFLAGS = 8

//
const Qt__VARFLAG_FDISPLAYBIND Qt__tagVARFLAGS = 16

//
const Qt__VARFLAG_FDEFAULTBIND Qt__tagVARFLAGS = 32

//
const Qt__VARFLAG_FHIDDEN Qt__tagVARFLAGS = 64

//
const Qt__VARFLAG_FRESTRICTED Qt__tagVARFLAGS = 128

//
const Qt__VARFLAG_FDEFAULTCOLLELEM Qt__tagVARFLAGS = 256

//
const Qt__VARFLAG_FUIDEFAULT Qt__tagVARFLAGS = 512

//
const Qt__VARFLAG_FNONBROWSABLE Qt__tagVARFLAGS = 1024

//
const Qt__VARFLAG_FREPLACEABLE Qt__tagVARFLAGS = 2048

//
const Qt__VARFLAG_FIMMEDIATEBIND Qt__tagVARFLAGS = 4096

func tagVARFLAGSItemName(val int) string {
	switch val {
	case Qt__VARFLAG_FREADONLY: // 1
		return "VARFLAG_FREADONLY"
	case Qt__VARFLAG_FSOURCE: // 2
		return "VARFLAG_FSOURCE"
	case Qt__VARFLAG_FBINDABLE: // 4
		return "VARFLAG_FBINDABLE"
	case Qt__VARFLAG_FREQUESTEDIT: // 8
		return "VARFLAG_FREQUESTEDIT"
	case Qt__VARFLAG_FDISPLAYBIND: // 16
		return "VARFLAG_FDISPLAYBIND"
	case Qt__VARFLAG_FDEFAULTBIND: // 32
		return "VARFLAG_FDEFAULTBIND"
	case Qt__VARFLAG_FHIDDEN: // 64
		return "VARFLAG_FHIDDEN"
	case Qt__VARFLAG_FRESTRICTED: // 128
		return "VARFLAG_FRESTRICTED"
	case Qt__VARFLAG_FDEFAULTCOLLELEM: // 256
		return "VARFLAG_FDEFAULTCOLLELEM"
	case Qt__VARFLAG_FUIDEFAULT: // 512
		return "VARFLAG_FUIDEFAULT"
	case Qt__VARFLAG_FNONBROWSABLE: // 1024
		return "VARFLAG_FNONBROWSABLE"
	case Qt__VARFLAG_FREPLACEABLE: // 2048
		return "VARFLAG_FREPLACEABLE"
	case Qt__VARFLAG_FIMMEDIATEBIND: // 4096
		return "VARFLAG_FIMMEDIATEBIND"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagDESCKIND = int // stdglobal
//
const Qt__DESCKIND_NONE Qt__tagDESCKIND = 0

//
const Qt__DESCKIND_FUNCDESC Qt__tagDESCKIND = 1

//
const Qt__DESCKIND_VARDESC Qt__tagDESCKIND = 2

//
const Qt__DESCKIND_TYPECOMP Qt__tagDESCKIND = 3

//
const Qt__DESCKIND_IMPLICITAPPOBJ Qt__tagDESCKIND = 4

//
const Qt__DESCKIND_MAX Qt__tagDESCKIND = 5

func tagDESCKINDItemName(val int) string {
	switch val {
	case Qt__DESCKIND_NONE: // 0
		return "DESCKIND_NONE"
	case Qt__DESCKIND_FUNCDESC: // 1
		return "DESCKIND_FUNCDESC"
	case Qt__DESCKIND_VARDESC: // 2
		return "DESCKIND_VARDESC"
	case Qt__DESCKIND_TYPECOMP: // 3
		return "DESCKIND_TYPECOMP"
	case Qt__DESCKIND_IMPLICITAPPOBJ: // 4
		return "DESCKIND_IMPLICITAPPOBJ"
	case Qt__DESCKIND_MAX: // 5
		return "DESCKIND_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagSYSKIND = int // stdglobal
//
const Qt__SYS_WIN16 Qt__tagSYSKIND = 0

//
const Qt__SYS_WIN32 Qt__tagSYSKIND = 1

//
const Qt__SYS_MAC Qt__tagSYSKIND = 2

//
const Qt__SYS_WIN64 Qt__tagSYSKIND = 3

func tagSYSKINDItemName(val int) string {
	switch val {
	case Qt__SYS_WIN16: // 0
		return "SYS_WIN16"
	case Qt__SYS_WIN32: // 1
		return "SYS_WIN32"
	case Qt__SYS_MAC: // 2
		return "SYS_MAC"
	case Qt__SYS_WIN64: // 3
		return "SYS_WIN64"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagLIBFLAGS = int // stdglobal
//
const Qt__LIBFLAG_FRESTRICTED Qt__tagLIBFLAGS = 1

//
const Qt__LIBFLAG_FCONTROL Qt__tagLIBFLAGS = 2

//
const Qt__LIBFLAG_FHIDDEN Qt__tagLIBFLAGS = 4

//
const Qt__LIBFLAG_FHASDISKIMAGE Qt__tagLIBFLAGS = 8

func tagLIBFLAGSItemName(val int) string {
	switch val {
	case Qt__LIBFLAG_FRESTRICTED: // 1
		return "LIBFLAG_FRESTRICTED"
	case Qt__LIBFLAG_FCONTROL: // 2
		return "LIBFLAG_FCONTROL"
	case Qt__LIBFLAG_FHIDDEN: // 4
		return "LIBFLAG_FHIDDEN"
	case Qt__LIBFLAG_FHASDISKIMAGE: // 8
		return "LIBFLAG_FHASDISKIMAGE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagCHANGEKIND = int // stdglobal
//
const Qt__CHANGEKIND_ADDMEMBER Qt__tagCHANGEKIND = 0

//
const Qt__CHANGEKIND_DELETEMEMBER Qt__tagCHANGEKIND = 1

//
const Qt__CHANGEKIND_SETNAMES Qt__tagCHANGEKIND = 2

//
const Qt__CHANGEKIND_SETDOCUMENTATION Qt__tagCHANGEKIND = 3

//
const Qt__CHANGEKIND_GENERAL Qt__tagCHANGEKIND = 4

//
const Qt__CHANGEKIND_INVALIDATE Qt__tagCHANGEKIND = 5

//
const Qt__CHANGEKIND_CHANGEFAILED Qt__tagCHANGEKIND = 6

//
const Qt__CHANGEKIND_MAX Qt__tagCHANGEKIND = 7

func tagCHANGEKINDItemName(val int) string {
	switch val {
	case Qt__CHANGEKIND_ADDMEMBER: // 0
		return "CHANGEKIND_ADDMEMBER"
	case Qt__CHANGEKIND_DELETEMEMBER: // 1
		return "CHANGEKIND_DELETEMEMBER"
	case Qt__CHANGEKIND_SETNAMES: // 2
		return "CHANGEKIND_SETNAMES"
	case Qt__CHANGEKIND_SETDOCUMENTATION: // 3
		return "CHANGEKIND_SETDOCUMENTATION"
	case Qt__CHANGEKIND_GENERAL: // 4
		return "CHANGEKIND_GENERAL"
	case Qt__CHANGEKIND_INVALIDATE: // 5
		return "CHANGEKIND_INVALIDATE"
	case Qt__CHANGEKIND_CHANGEFAILED: // 6
		return "CHANGEKIND_CHANGEFAILED"
	case Qt__CHANGEKIND_MAX: // 7
		return "CHANGEKIND_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagDOMNodeType = int // stdglobal
//
const Qt__NODE_INVALID Qt__tagDOMNodeType = 0

//
const Qt__NODE_ELEMENT Qt__tagDOMNodeType = 1

//
const Qt__NODE_ATTRIBUTE Qt__tagDOMNodeType = 2

//
const Qt__NODE_TEXT Qt__tagDOMNodeType = 3

//
const Qt__NODE_CDATA_SECTION Qt__tagDOMNodeType = 4

//
const Qt__NODE_ENTITY_REFERENCE Qt__tagDOMNodeType = 5

//
const Qt__NODE_ENTITY Qt__tagDOMNodeType = 6

//
const Qt__NODE_PROCESSING_INSTRUCTION Qt__tagDOMNodeType = 7

//
const Qt__NODE_COMMENT Qt__tagDOMNodeType = 8

//
const Qt__NODE_DOCUMENT Qt__tagDOMNodeType = 9

//
const Qt__NODE_DOCUMENT_TYPE Qt__tagDOMNodeType = 10

//
const Qt__NODE_DOCUMENT_FRAGMENT Qt__tagDOMNodeType = 11

//
const Qt__NODE_NOTATION Qt__tagDOMNodeType = 12

func tagDOMNodeTypeItemName(val int) string {
	switch val {
	case Qt__NODE_INVALID: // 0
		return "NODE_INVALID"
	case Qt__NODE_ELEMENT: // 1
		return "NODE_ELEMENT"
	case Qt__NODE_ATTRIBUTE: // 2
		return "NODE_ATTRIBUTE"
	case Qt__NODE_TEXT: // 3
		return "NODE_TEXT"
	case Qt__NODE_CDATA_SECTION: // 4
		return "NODE_CDATA_SECTION"
	case Qt__NODE_ENTITY_REFERENCE: // 5
		return "NODE_ENTITY_REFERENCE"
	case Qt__NODE_ENTITY: // 6
		return "NODE_ENTITY"
	case Qt__NODE_PROCESSING_INSTRUCTION: // 7
		return "NODE_PROCESSING_INSTRUCTION"
	case Qt__NODE_COMMENT: // 8
		return "NODE_COMMENT"
	case Qt__NODE_DOCUMENT: // 9
		return "NODE_DOCUMENT"
	case Qt__NODE_DOCUMENT_TYPE: // 10
		return "NODE_DOCUMENT_TYPE"
	case Qt__NODE_DOCUMENT_FRAGMENT: // 11
		return "NODE_DOCUMENT_FRAGMENT"
	case Qt__NODE_NOTATION: // 12
		return "NODE_NOTATION"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagXMLEMEM_TYPE = int // stdglobal
//
const Qt__XMLELEMTYPE_ELEMENT Qt__tagXMLEMEM_TYPE = 0

//
const Qt__XMLELEMTYPE_TEXT Qt__tagXMLEMEM_TYPE = 1

//
const Qt__XMLELEMTYPE_COMMENT Qt__tagXMLEMEM_TYPE = 2

//
const Qt__XMLELEMTYPE_DOCUMENT Qt__tagXMLEMEM_TYPE = 3

//
const Qt__XMLELEMTYPE_DTD Qt__tagXMLEMEM_TYPE = 4

//
const Qt__XMLELEMTYPE_PI Qt__tagXMLEMEM_TYPE = 5

//
const Qt__XMLELEMTYPE_OTHER Qt__tagXMLEMEM_TYPE = 6

func tagXMLEMEM_TYPEItemName(val int) string {
	switch val {
	case Qt__XMLELEMTYPE_ELEMENT: // 0
		return "XMLELEMTYPE_ELEMENT"
	case Qt__XMLELEMTYPE_TEXT: // 1
		return "XMLELEMTYPE_TEXT"
	case Qt__XMLELEMTYPE_COMMENT: // 2
		return "XMLELEMTYPE_COMMENT"
	case Qt__XMLELEMTYPE_DOCUMENT: // 3
		return "XMLELEMTYPE_DOCUMENT"
	case Qt__XMLELEMTYPE_DTD: // 4
		return "XMLELEMTYPE_DTD"
	case Qt__XMLELEMTYPE_PI: // 5
		return "XMLELEMTYPE_PI"
	case Qt__XMLELEMTYPE_OTHER: // 6
		return "XMLELEMTYPE_OTHER"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000002 = int // stdglobal
//
const Qt__BINDVERB_GET Qt____WIDL_urlmon_generated_name_00000002 = 0

//
const Qt__BINDVERB_POST Qt____WIDL_urlmon_generated_name_00000002 = 1

//
const Qt__BINDVERB_PUT Qt____WIDL_urlmon_generated_name_00000002 = 2

//
const Qt__BINDVERB_CUSTOM Qt____WIDL_urlmon_generated_name_00000002 = 3

func __WIDL_urlmon_generated_name_00000002ItemName(val int) string {
	switch val {
	case Qt__BINDVERB_GET: // 0
		return "BINDVERB_GET"
	case Qt__BINDVERB_POST: // 1
		return "BINDVERB_POST"
	case Qt__BINDVERB_PUT: // 2
		return "BINDVERB_PUT"
	case Qt__BINDVERB_CUSTOM: // 3
		return "BINDVERB_CUSTOM"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000003 = int // stdglobal
//
const Qt__BINDINFOF_URLENCODESTGMEDDATA Qt____WIDL_urlmon_generated_name_00000003 = 1

//
const Qt__BINDINFOF_URLENCODEDEXTRAINFO Qt____WIDL_urlmon_generated_name_00000003 = 2

func __WIDL_urlmon_generated_name_00000003ItemName(val int) string {
	switch val {
	case Qt__BINDINFOF_URLENCODESTGMEDDATA: // 1
		return "BINDINFOF_URLENCODESTGMEDDATA"
	case Qt__BINDINFOF_URLENCODEDEXTRAINFO: // 2
		return "BINDINFOF_URLENCODEDEXTRAINFO"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000004 = int // stdglobal
//
const Qt__BINDF_ASYNCHRONOUS Qt____WIDL_urlmon_generated_name_00000004 = 1

//
const Qt__BINDF_ASYNCSTORAGE Qt____WIDL_urlmon_generated_name_00000004 = 2

//
const Qt__BINDF_NOPROGRESSIVERENDERING Qt____WIDL_urlmon_generated_name_00000004 = 4

//
const Qt__BINDF_OFFLINEOPERATION Qt____WIDL_urlmon_generated_name_00000004 = 8

//
const Qt__BINDF_GETNEWESTVERSION Qt____WIDL_urlmon_generated_name_00000004 = 16

//
const Qt__BINDF_NOWRITECACHE Qt____WIDL_urlmon_generated_name_00000004 = 32

//
const Qt__BINDF_NEEDFILE Qt____WIDL_urlmon_generated_name_00000004 = 64

//
const Qt__BINDF_PULLDATA Qt____WIDL_urlmon_generated_name_00000004 = 128

//
const Qt__BINDF_IGNORESECURITYPROBLEM Qt____WIDL_urlmon_generated_name_00000004 = 256

//
const Qt__BINDF_RESYNCHRONIZE Qt____WIDL_urlmon_generated_name_00000004 = 512

//
const Qt__BINDF_HYPERLINK Qt____WIDL_urlmon_generated_name_00000004 = 1024

//
const Qt__BINDF_NO_UI Qt____WIDL_urlmon_generated_name_00000004 = 2048

//
const Qt__BINDF_SILENTOPERATION Qt____WIDL_urlmon_generated_name_00000004 = 4096

//
const Qt__BINDF_PRAGMA_NO_CACHE Qt____WIDL_urlmon_generated_name_00000004 = 8192

//
const Qt__BINDF_GETCLASSOBJECT Qt____WIDL_urlmon_generated_name_00000004 = 16384

//
const Qt__BINDF_RESERVED_1 Qt____WIDL_urlmon_generated_name_00000004 = 32768

//
const Qt__BINDF_FREE_THREADED Qt____WIDL_urlmon_generated_name_00000004 = 65536

//
const Qt__BINDF_DIRECT_READ Qt____WIDL_urlmon_generated_name_00000004 = 131072

//
const Qt__BINDF_FORMS_SUBMIT Qt____WIDL_urlmon_generated_name_00000004 = 262144

//
const Qt__BINDF_GETFROMCACHE_IF_NET_FAIL Qt____WIDL_urlmon_generated_name_00000004 = 524288

//
const Qt__BINDF_FROMURLMON Qt____WIDL_urlmon_generated_name_00000004 = 1048576

//
const Qt__BINDF_FWD_BACK Qt____WIDL_urlmon_generated_name_00000004 = 2097152

//
const Qt__BINDF_PREFERDEFAULTHANDLER Qt____WIDL_urlmon_generated_name_00000004 = 4194304

//
const Qt__BINDF_ENFORCERESTRICTED Qt____WIDL_urlmon_generated_name_00000004 = 8388608

func __WIDL_urlmon_generated_name_00000004ItemName(val int) string {
	switch val {
	case Qt__BINDF_ASYNCHRONOUS: // 1
		return "BINDF_ASYNCHRONOUS"
	case Qt__BINDF_ASYNCSTORAGE: // 2
		return "BINDF_ASYNCSTORAGE"
	case Qt__BINDF_NOPROGRESSIVERENDERING: // 4
		return "BINDF_NOPROGRESSIVERENDERING"
	case Qt__BINDF_OFFLINEOPERATION: // 8
		return "BINDF_OFFLINEOPERATION"
	case Qt__BINDF_GETNEWESTVERSION: // 16
		return "BINDF_GETNEWESTVERSION"
	case Qt__BINDF_NOWRITECACHE: // 32
		return "BINDF_NOWRITECACHE"
	case Qt__BINDF_NEEDFILE: // 64
		return "BINDF_NEEDFILE"
	case Qt__BINDF_PULLDATA: // 128
		return "BINDF_PULLDATA"
	case Qt__BINDF_IGNORESECURITYPROBLEM: // 256
		return "BINDF_IGNORESECURITYPROBLEM"
	case Qt__BINDF_RESYNCHRONIZE: // 512
		return "BINDF_RESYNCHRONIZE"
	case Qt__BINDF_HYPERLINK: // 1024
		return "BINDF_HYPERLINK"
	case Qt__BINDF_NO_UI: // 2048
		return "BINDF_NO_UI"
	case Qt__BINDF_SILENTOPERATION: // 4096
		return "BINDF_SILENTOPERATION"
	case Qt__BINDF_PRAGMA_NO_CACHE: // 8192
		return "BINDF_PRAGMA_NO_CACHE"
	case Qt__BINDF_GETCLASSOBJECT: // 16384
		return "BINDF_GETCLASSOBJECT"
	case Qt__BINDF_RESERVED_1: // 32768
		return "BINDF_RESERVED_1"
	case Qt__BINDF_FREE_THREADED: // 65536
		return "BINDF_FREE_THREADED"
	case Qt__BINDF_DIRECT_READ: // 131072
		return "BINDF_DIRECT_READ"
	case Qt__BINDF_FORMS_SUBMIT: // 262144
		return "BINDF_FORMS_SUBMIT"
	case Qt__BINDF_GETFROMCACHE_IF_NET_FAIL: // 524288
		return "BINDF_GETFROMCACHE_IF_NET_FAIL"
	case Qt__BINDF_FROMURLMON: // 1048576
		return "BINDF_FROMURLMON"
	case Qt__BINDF_FWD_BACK: // 2097152
		return "BINDF_FWD_BACK"
	case Qt__BINDF_PREFERDEFAULTHANDLER: // 4194304
		return "BINDF_PREFERDEFAULTHANDLER"
	case Qt__BINDF_ENFORCERESTRICTED: // 8388608
		return "BINDF_ENFORCERESTRICTED"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000005 = int // stdglobal
//
const Qt__URL_ENCODING_NONE Qt____WIDL_urlmon_generated_name_00000005 = 0

//
const Qt__URL_ENCODING_ENABLE_UTF8 Qt____WIDL_urlmon_generated_name_00000005 = 268435456

//
const Qt__URL_ENCODING_DISABLE_UTF8 Qt____WIDL_urlmon_generated_name_00000005 = 536870912

func __WIDL_urlmon_generated_name_00000005ItemName(val int) string {
	switch val {
	case Qt__URL_ENCODING_NONE: // 0
		return "URL_ENCODING_NONE"
	case Qt__URL_ENCODING_ENABLE_UTF8: // 268435456
		return "URL_ENCODING_ENABLE_UTF8"
	case Qt__URL_ENCODING_DISABLE_UTF8: // 536870912
		return "URL_ENCODING_DISABLE_UTF8"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000006 = int // stdglobal
//
const Qt__BINDINFO_OPTIONS_WININETFLAG Qt____WIDL_urlmon_generated_name_00000006 = 65536

//
const Qt__BINDINFO_OPTIONS_ENABLE_UTF8 Qt____WIDL_urlmon_generated_name_00000006 = 131072

//
const Qt__BINDINFO_OPTIONS_DISABLE_UTF8 Qt____WIDL_urlmon_generated_name_00000006 = 262144

//
const Qt__BINDINFO_OPTIONS_USE_IE_ENCODING Qt____WIDL_urlmon_generated_name_00000006 = 524288

//
const Qt__BINDINFO_OPTIONS_BINDTOOBJECT Qt____WIDL_urlmon_generated_name_00000006 = 1048576

//
const Qt__BINDINFO_OPTIONS_SECURITYOPTOUT Qt____WIDL_urlmon_generated_name_00000006 = 2097152

//
const Qt__BINDINFO_OPTIONS_IGNOREMIMETEXTPLAIN Qt____WIDL_urlmon_generated_name_00000006 = 4194304

//
const Qt__BINDINFO_OPTIONS_USEBINDSTRINGCREDS Qt____WIDL_urlmon_generated_name_00000006 = 8388608

//
const Qt__BINDINFO_OPTIONS_IGNOREHTTPHTTPSREDIRECTS Qt____WIDL_urlmon_generated_name_00000006 = 16777216

//
const Qt__BINDINFO_OPTIONS_IGNORE_SSLERRORS_ONCE Qt____WIDL_urlmon_generated_name_00000006 = 33554432

//
const Qt__BINDINFO_WPC_DOWNLOADBLOCKED Qt____WIDL_urlmon_generated_name_00000006 = 134217728

//
const Qt__BINDINFO_WPC_LOGGING_ENABLED Qt____WIDL_urlmon_generated_name_00000006 = 268435456

//
const Qt__BINDINFO_OPTIONS_ALLOWCONNECTDATA Qt____WIDL_urlmon_generated_name_00000006 = 536870912

//
const Qt__BINDINFO_OPTIONS_DISABLEAUTOREDIRECTS Qt____WIDL_urlmon_generated_name_00000006 = 1073741824

//
const Qt__BINDINFO_OPTIONS_SHDOCVW_NAVIGATE Qt____WIDL_urlmon_generated_name_00000006 = -2147483648

func __WIDL_urlmon_generated_name_00000006ItemName(val int) string {
	switch val {
	case Qt__BINDINFO_OPTIONS_WININETFLAG: // 65536
		return "BINDINFO_OPTIONS_WININETFLAG"
	case Qt__BINDINFO_OPTIONS_ENABLE_UTF8: // 131072
		return "BINDINFO_OPTIONS_ENABLE_UTF8"
	case Qt__BINDINFO_OPTIONS_DISABLE_UTF8: // 262144
		return "BINDINFO_OPTIONS_DISABLE_UTF8"
	case Qt__BINDINFO_OPTIONS_USE_IE_ENCODING: // 524288
		return "BINDINFO_OPTIONS_USE_IE_ENCODING"
	case Qt__BINDINFO_OPTIONS_BINDTOOBJECT: // 1048576
		return "BINDINFO_OPTIONS_BINDTOOBJECT"
	case Qt__BINDINFO_OPTIONS_SECURITYOPTOUT: // 2097152
		return "BINDINFO_OPTIONS_SECURITYOPTOUT"
	case Qt__BINDINFO_OPTIONS_IGNOREMIMETEXTPLAIN: // 4194304
		return "BINDINFO_OPTIONS_IGNOREMIMETEXTPLAIN"
	case Qt__BINDINFO_OPTIONS_USEBINDSTRINGCREDS: // 8388608
		return "BINDINFO_OPTIONS_USEBINDSTRINGCREDS"
	case Qt__BINDINFO_OPTIONS_IGNOREHTTPHTTPSREDIRECTS: // 16777216
		return "BINDINFO_OPTIONS_IGNOREHTTPHTTPSREDIRECTS"
	case Qt__BINDINFO_OPTIONS_IGNORE_SSLERRORS_ONCE: // 33554432
		return "BINDINFO_OPTIONS_IGNORE_SSLERRORS_ONCE"
	case Qt__BINDINFO_WPC_DOWNLOADBLOCKED: // 134217728
		return "BINDINFO_WPC_DOWNLOADBLOCKED"
	case Qt__BINDINFO_WPC_LOGGING_ENABLED: // 268435456
		return "BINDINFO_WPC_LOGGING_ENABLED"
	case Qt__BINDINFO_OPTIONS_ALLOWCONNECTDATA: // 536870912
		return "BINDINFO_OPTIONS_ALLOWCONNECTDATA"
	case Qt__BINDINFO_OPTIONS_DISABLEAUTOREDIRECTS: // 1073741824
		return "BINDINFO_OPTIONS_DISABLEAUTOREDIRECTS"
	case Qt__BINDINFO_OPTIONS_SHDOCVW_NAVIGATE: // -2147483648
		return "BINDINFO_OPTIONS_SHDOCVW_NAVIGATE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000007 = int // stdglobal
//
const Qt__BSCF_FIRSTDATANOTIFICATION Qt____WIDL_urlmon_generated_name_00000007 = 1

//
const Qt__BSCF_INTERMEDIATEDATANOTIFICATION Qt____WIDL_urlmon_generated_name_00000007 = 2

//
const Qt__BSCF_LASTDATANOTIFICATION Qt____WIDL_urlmon_generated_name_00000007 = 4

//
const Qt__BSCF_DATAFULLYAVAILABLE Qt____WIDL_urlmon_generated_name_00000007 = 8

//
const Qt__BSCF_AVAILABLEDATASIZEUNKNOWN Qt____WIDL_urlmon_generated_name_00000007 = 16

//
const Qt__BSCF_SKIPDRAINDATAFORFILEURLS Qt____WIDL_urlmon_generated_name_00000007 = 32

//
const Qt__BSCF_64BITLENGTHDOWNLOAD Qt____WIDL_urlmon_generated_name_00000007 = 64

func __WIDL_urlmon_generated_name_00000007ItemName(val int) string {
	switch val {
	case Qt__BSCF_FIRSTDATANOTIFICATION: // 1
		return "BSCF_FIRSTDATANOTIFICATION"
	case Qt__BSCF_INTERMEDIATEDATANOTIFICATION: // 2
		return "BSCF_INTERMEDIATEDATANOTIFICATION"
	case Qt__BSCF_LASTDATANOTIFICATION: // 4
		return "BSCF_LASTDATANOTIFICATION"
	case Qt__BSCF_DATAFULLYAVAILABLE: // 8
		return "BSCF_DATAFULLYAVAILABLE"
	case Qt__BSCF_AVAILABLEDATASIZEUNKNOWN: // 16
		return "BSCF_AVAILABLEDATASIZEUNKNOWN"
	case Qt__BSCF_SKIPDRAINDATAFORFILEURLS: // 32
		return "BSCF_SKIPDRAINDATAFORFILEURLS"
	case Qt__BSCF_64BITLENGTHDOWNLOAD: // 64
		return "BSCF_64BITLENGTHDOWNLOAD"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__BINDSTATUS = int // stdglobal
//
const Qt__BINDSTATUS_FINDINGRESOURCE Qt__BINDSTATUS = 1

//
const Qt__BINDSTATUS_CONNECTING Qt__BINDSTATUS = 2

//
const Qt__BINDSTATUS_REDIRECTING Qt__BINDSTATUS = 3

//
const Qt__BINDSTATUS_BEGINDOWNLOADDATA Qt__BINDSTATUS = 4

//
const Qt__BINDSTATUS_DOWNLOADINGDATA Qt__BINDSTATUS = 5

//
const Qt__BINDSTATUS_ENDDOWNLOADDATA Qt__BINDSTATUS = 6

//
const Qt__BINDSTATUS_BEGINDOWNLOADCOMPONENTS Qt__BINDSTATUS = 7

//
const Qt__BINDSTATUS_INSTALLINGCOMPONENTS Qt__BINDSTATUS = 8

//
const Qt__BINDSTATUS_ENDDOWNLOADCOMPONENTS Qt__BINDSTATUS = 9

//
const Qt__BINDSTATUS_USINGCACHEDCOPY Qt__BINDSTATUS = 10

//
const Qt__BINDSTATUS_SENDINGREQUEST Qt__BINDSTATUS = 11

//
const Qt__BINDSTATUS_CLASSIDAVAILABLE Qt__BINDSTATUS = 12

//
const Qt__BINDSTATUS_MIMETYPEAVAILABLE Qt__BINDSTATUS = 13

//
const Qt__BINDSTATUS_CACHEFILENAMEAVAILABLE Qt__BINDSTATUS = 14

//
const Qt__BINDSTATUS_BEGINSYNCOPERATION Qt__BINDSTATUS = 15

//
const Qt__BINDSTATUS_ENDSYNCOPERATION Qt__BINDSTATUS = 16

//
const Qt__BINDSTATUS_BEGINUPLOADDATA Qt__BINDSTATUS = 17

//
const Qt__BINDSTATUS_UPLOADINGDATA Qt__BINDSTATUS = 18

//
const Qt__BINDSTATUS_ENDUPLOADINGDATA Qt__BINDSTATUS = 19

//
const Qt__BINDSTATUS_PROTOCOLCLASSID Qt__BINDSTATUS = 20

//
const Qt__BINDSTATUS_ENCODING Qt__BINDSTATUS = 21

//
const Qt__BINDSTATUS_VERIFIEDMIMETYPEAVAILABLE Qt__BINDSTATUS = 22

//
const Qt__BINDSTATUS_CLASSINSTALLLOCATION Qt__BINDSTATUS = 23

//
const Qt__BINDSTATUS_DECODING Qt__BINDSTATUS = 24

//
const Qt__BINDSTATUS_LOADINGMIMEHANDLER Qt__BINDSTATUS = 25

//
const Qt__BINDSTATUS_CONTENTDISPOSITIONATTACH Qt__BINDSTATUS = 26

//
const Qt__BINDSTATUS_FILTERREPORTMIMETYPE Qt__BINDSTATUS = 27

//
const Qt__BINDSTATUS_CLSIDCANINSTANTIATE Qt__BINDSTATUS = 28

//
const Qt__BINDSTATUS_IUNKNOWNAVAILABLE Qt__BINDSTATUS = 29

//
const Qt__BINDSTATUS_DIRECTBIND Qt__BINDSTATUS = 30

//
const Qt__BINDSTATUS_RAWMIMETYPE Qt__BINDSTATUS = 31

//
const Qt__BINDSTATUS_PROXYDETECTING Qt__BINDSTATUS = 32

//
const Qt__BINDSTATUS_ACCEPTRANGES Qt__BINDSTATUS = 33

//
const Qt__BINDSTATUS_COOKIE_SENT Qt__BINDSTATUS = 34

//
const Qt__BINDSTATUS_COMPACT_POLICY_RECEIVED Qt__BINDSTATUS = 35

//
const Qt__BINDSTATUS_COOKIE_SUPPRESSED Qt__BINDSTATUS = 36

//
const Qt__BINDSTATUS_COOKIE_STATE_UNKNOWN Qt__BINDSTATUS = 37

//
const Qt__BINDSTATUS_COOKIE_STATE_ACCEPT Qt__BINDSTATUS = 38

//
const Qt__BINDSTATUS_COOKIE_STATE_REJECT Qt__BINDSTATUS = 39

//
const Qt__BINDSTATUS_COOKIE_STATE_PROMPT Qt__BINDSTATUS = 40

//
const Qt__BINDSTATUS_COOKIE_STATE_LEASH Qt__BINDSTATUS = 41

//
const Qt__BINDSTATUS_COOKIE_STATE_DOWNGRADE Qt__BINDSTATUS = 42

//
const Qt__BINDSTATUS_POLICY_HREF Qt__BINDSTATUS = 43

//
const Qt__BINDSTATUS_P3P_HEADER Qt__BINDSTATUS = 44

//
const Qt__BINDSTATUS_SESSION_COOKIE_RECEIVED Qt__BINDSTATUS = 45

//
const Qt__BINDSTATUS_PERSISTENT_COOKIE_RECEIVED Qt__BINDSTATUS = 46

//
const Qt__BINDSTATUS_SESSION_COOKIES_ALLOWED Qt__BINDSTATUS = 47

//
const Qt__BINDSTATUS_CACHECONTROL Qt__BINDSTATUS = 48

//
const Qt__BINDSTATUS_CONTENTDISPOSITIONFILENAME Qt__BINDSTATUS = 49

//
const Qt__BINDSTATUS_MIMETEXTPLAINMISMATCH Qt__BINDSTATUS = 50

//
const Qt__BINDSTATUS_PUBLISHERAVAILABLE Qt__BINDSTATUS = 51

//
const Qt__BINDSTATUS_DISPLAYNAMEAVAILABLE Qt__BINDSTATUS = 52

//
const Qt__BINDSTATUS_SSLUX_NAVBLOCKED Qt__BINDSTATUS = 53

//
const Qt__BINDSTATUS_SERVER_MIMETYPEAVAILABLE Qt__BINDSTATUS = 54

//
const Qt__BINDSTATUS_SNIFFED_CLASSIDAVAILABLE Qt__BINDSTATUS = 55

//
const Qt__BINDSTATUS_64BIT_PROGRESS Qt__BINDSTATUS = 56

//
const Qt__BINDSTATUS_LAST Qt__BINDSTATUS = 56

//
const Qt__BINDSTATUS_RESERVED_0 Qt__BINDSTATUS = 57

//
const Qt__BINDSTATUS_RESERVED_1 Qt__BINDSTATUS = 58

//
const Qt__BINDSTATUS_RESERVED_2 Qt__BINDSTATUS = 59

//
const Qt__BINDSTATUS_RESERVED_3 Qt__BINDSTATUS = 60

//
const Qt__BINDSTATUS_RESERVED_4 Qt__BINDSTATUS = 61

//
const Qt__BINDSTATUS_RESERVED_5 Qt__BINDSTATUS = 62

//
const Qt__BINDSTATUS_RESERVED_6 Qt__BINDSTATUS = 63

//
const Qt__BINDSTATUS_RESERVED_7 Qt__BINDSTATUS = 64

//
const Qt__BINDSTATUS_RESERVED_8 Qt__BINDSTATUS = 65

//
const Qt__BINDSTATUS_RESERVED_9 Qt__BINDSTATUS = 66

//
const Qt__BINDSTATUS_LAST_PRIVATE Qt__BINDSTATUS = 66

func BINDSTATUSItemName(val int) string {
	switch val {
	case Qt__BINDSTATUS_FINDINGRESOURCE: // 1
		return "BINDSTATUS_FINDINGRESOURCE"
	case Qt__BINDSTATUS_CONNECTING: // 2
		return "BINDSTATUS_CONNECTING"
	case Qt__BINDSTATUS_REDIRECTING: // 3
		return "BINDSTATUS_REDIRECTING"
	case Qt__BINDSTATUS_BEGINDOWNLOADDATA: // 4
		return "BINDSTATUS_BEGINDOWNLOADDATA"
	case Qt__BINDSTATUS_DOWNLOADINGDATA: // 5
		return "BINDSTATUS_DOWNLOADINGDATA"
	case Qt__BINDSTATUS_ENDDOWNLOADDATA: // 6
		return "BINDSTATUS_ENDDOWNLOADDATA"
	case Qt__BINDSTATUS_BEGINDOWNLOADCOMPONENTS: // 7
		return "BINDSTATUS_BEGINDOWNLOADCOMPONENTS"
	case Qt__BINDSTATUS_INSTALLINGCOMPONENTS: // 8
		return "BINDSTATUS_INSTALLINGCOMPONENTS"
	case Qt__BINDSTATUS_ENDDOWNLOADCOMPONENTS: // 9
		return "BINDSTATUS_ENDDOWNLOADCOMPONENTS"
	case Qt__BINDSTATUS_USINGCACHEDCOPY: // 10
		return "BINDSTATUS_USINGCACHEDCOPY"
	case Qt__BINDSTATUS_SENDINGREQUEST: // 11
		return "BINDSTATUS_SENDINGREQUEST"
	case Qt__BINDSTATUS_CLASSIDAVAILABLE: // 12
		return "BINDSTATUS_CLASSIDAVAILABLE"
	case Qt__BINDSTATUS_MIMETYPEAVAILABLE: // 13
		return "BINDSTATUS_MIMETYPEAVAILABLE"
	case Qt__BINDSTATUS_CACHEFILENAMEAVAILABLE: // 14
		return "BINDSTATUS_CACHEFILENAMEAVAILABLE"
	case Qt__BINDSTATUS_BEGINSYNCOPERATION: // 15
		return "BINDSTATUS_BEGINSYNCOPERATION"
	case Qt__BINDSTATUS_ENDSYNCOPERATION: // 16
		return "BINDSTATUS_ENDSYNCOPERATION"
	case Qt__BINDSTATUS_BEGINUPLOADDATA: // 17
		return "BINDSTATUS_BEGINUPLOADDATA"
	case Qt__BINDSTATUS_UPLOADINGDATA: // 18
		return "BINDSTATUS_UPLOADINGDATA"
	case Qt__BINDSTATUS_ENDUPLOADINGDATA: // 19
		return "BINDSTATUS_ENDUPLOADINGDATA"
	case Qt__BINDSTATUS_PROTOCOLCLASSID: // 20
		return "BINDSTATUS_PROTOCOLCLASSID"
	case Qt__BINDSTATUS_ENCODING: // 21
		return "BINDSTATUS_ENCODING"
	case Qt__BINDSTATUS_VERIFIEDMIMETYPEAVAILABLE: // 22
		return "BINDSTATUS_VERIFIEDMIMETYPEAVAILABLE"
	case Qt__BINDSTATUS_CLASSINSTALLLOCATION: // 23
		return "BINDSTATUS_CLASSINSTALLLOCATION"
	case Qt__BINDSTATUS_DECODING: // 24
		return "BINDSTATUS_DECODING"
	case Qt__BINDSTATUS_LOADINGMIMEHANDLER: // 25
		return "BINDSTATUS_LOADINGMIMEHANDLER"
	case Qt__BINDSTATUS_CONTENTDISPOSITIONATTACH: // 26
		return "BINDSTATUS_CONTENTDISPOSITIONATTACH"
	case Qt__BINDSTATUS_FILTERREPORTMIMETYPE: // 27
		return "BINDSTATUS_FILTERREPORTMIMETYPE"
	case Qt__BINDSTATUS_CLSIDCANINSTANTIATE: // 28
		return "BINDSTATUS_CLSIDCANINSTANTIATE"
	case Qt__BINDSTATUS_IUNKNOWNAVAILABLE: // 29
		return "BINDSTATUS_IUNKNOWNAVAILABLE"
	case Qt__BINDSTATUS_DIRECTBIND: // 30
		return "BINDSTATUS_DIRECTBIND"
	case Qt__BINDSTATUS_RAWMIMETYPE: // 31
		return "BINDSTATUS_RAWMIMETYPE"
	case Qt__BINDSTATUS_PROXYDETECTING: // 32
		return "BINDSTATUS_PROXYDETECTING"
	case Qt__BINDSTATUS_ACCEPTRANGES: // 33
		return "BINDSTATUS_ACCEPTRANGES"
	case Qt__BINDSTATUS_COOKIE_SENT: // 34
		return "BINDSTATUS_COOKIE_SENT"
	case Qt__BINDSTATUS_COMPACT_POLICY_RECEIVED: // 35
		return "BINDSTATUS_COMPACT_POLICY_RECEIVED"
	case Qt__BINDSTATUS_COOKIE_SUPPRESSED: // 36
		return "BINDSTATUS_COOKIE_SUPPRESSED"
	case Qt__BINDSTATUS_COOKIE_STATE_UNKNOWN: // 37
		return "BINDSTATUS_COOKIE_STATE_UNKNOWN"
	case Qt__BINDSTATUS_COOKIE_STATE_ACCEPT: // 38
		return "BINDSTATUS_COOKIE_STATE_ACCEPT"
	case Qt__BINDSTATUS_COOKIE_STATE_REJECT: // 39
		return "BINDSTATUS_COOKIE_STATE_REJECT"
	case Qt__BINDSTATUS_COOKIE_STATE_PROMPT: // 40
		return "BINDSTATUS_COOKIE_STATE_PROMPT"
	case Qt__BINDSTATUS_COOKIE_STATE_LEASH: // 41
		return "BINDSTATUS_COOKIE_STATE_LEASH"
	case Qt__BINDSTATUS_COOKIE_STATE_DOWNGRADE: // 42
		return "BINDSTATUS_COOKIE_STATE_DOWNGRADE"
	case Qt__BINDSTATUS_POLICY_HREF: // 43
		return "BINDSTATUS_POLICY_HREF"
	case Qt__BINDSTATUS_P3P_HEADER: // 44
		return "BINDSTATUS_P3P_HEADER"
	case Qt__BINDSTATUS_SESSION_COOKIE_RECEIVED: // 45
		return "BINDSTATUS_SESSION_COOKIE_RECEIVED"
	case Qt__BINDSTATUS_PERSISTENT_COOKIE_RECEIVED: // 46
		return "BINDSTATUS_PERSISTENT_COOKIE_RECEIVED"
	case Qt__BINDSTATUS_SESSION_COOKIES_ALLOWED: // 47
		return "BINDSTATUS_SESSION_COOKIES_ALLOWED"
	case Qt__BINDSTATUS_CACHECONTROL: // 48
		return "BINDSTATUS_CACHECONTROL"
	case Qt__BINDSTATUS_CONTENTDISPOSITIONFILENAME: // 49
		return "BINDSTATUS_CONTENTDISPOSITIONFILENAME"
	case Qt__BINDSTATUS_MIMETEXTPLAINMISMATCH: // 50
		return "BINDSTATUS_MIMETEXTPLAINMISMATCH"
	case Qt__BINDSTATUS_PUBLISHERAVAILABLE: // 51
		return "BINDSTATUS_PUBLISHERAVAILABLE"
	case Qt__BINDSTATUS_DISPLAYNAMEAVAILABLE: // 52
		return "BINDSTATUS_DISPLAYNAMEAVAILABLE"
	case Qt__BINDSTATUS_SSLUX_NAVBLOCKED: // 53
		return "BINDSTATUS_SSLUX_NAVBLOCKED"
	case Qt__BINDSTATUS_SERVER_MIMETYPEAVAILABLE: // 54
		return "BINDSTATUS_SERVER_MIMETYPEAVAILABLE"
	case Qt__BINDSTATUS_SNIFFED_CLASSIDAVAILABLE: // 55
		return "BINDSTATUS_SNIFFED_CLASSIDAVAILABLE"
	case Qt__BINDSTATUS_64BIT_PROGRESS: // 56
		return "BINDSTATUS_64BIT_PROGRESS,BINDSTATUS_LAST"
		// case Qt__BINDSTATUS_LAST: // 56
		// return ""
	case Qt__BINDSTATUS_RESERVED_0: // 57
		return "BINDSTATUS_RESERVED_0"
	case Qt__BINDSTATUS_RESERVED_1: // 58
		return "BINDSTATUS_RESERVED_1"
	case Qt__BINDSTATUS_RESERVED_2: // 59
		return "BINDSTATUS_RESERVED_2"
	case Qt__BINDSTATUS_RESERVED_3: // 60
		return "BINDSTATUS_RESERVED_3"
	case Qt__BINDSTATUS_RESERVED_4: // 61
		return "BINDSTATUS_RESERVED_4"
	case Qt__BINDSTATUS_RESERVED_5: // 62
		return "BINDSTATUS_RESERVED_5"
	case Qt__BINDSTATUS_RESERVED_6: // 63
		return "BINDSTATUS_RESERVED_6"
	case Qt__BINDSTATUS_RESERVED_7: // 64
		return "BINDSTATUS_RESERVED_7"
	case Qt__BINDSTATUS_RESERVED_8: // 65
		return "BINDSTATUS_RESERVED_8"
	case Qt__BINDSTATUS_RESERVED_9: // 66
		return "BINDSTATUS_RESERVED_9,BINDSTATUS_LAST_PRIVATE"
		// case Qt__BINDSTATUS_LAST_PRIVATE: // 66
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000008 = int // stdglobal
//
const Qt__BINDF2_DISABLEBASICOVERHTTP Qt____WIDL_urlmon_generated_name_00000008 = 1

//
const Qt__BINDF2_DISABLEAUTOCOOKIEHANDLING Qt____WIDL_urlmon_generated_name_00000008 = 2

//
const Qt__BINDF2_READ_DATA_GREATER_THAN_4GB Qt____WIDL_urlmon_generated_name_00000008 = 4

//
const Qt__BINDF2_DISABLE_HTTP_REDIRECT_XSECURITYID Qt____WIDL_urlmon_generated_name_00000008 = 8

//
const Qt__BINDF2_RESERVED_3 Qt____WIDL_urlmon_generated_name_00000008 = 536870912

//
const Qt__BINDF2_RESERVED_2 Qt____WIDL_urlmon_generated_name_00000008 = 1073741824

//
const Qt__BINDF2_RESERVED_1 Qt____WIDL_urlmon_generated_name_00000008 = -2147483648

func __WIDL_urlmon_generated_name_00000008ItemName(val int) string {
	switch val {
	case Qt__BINDF2_DISABLEBASICOVERHTTP: // 1
		return "BINDF2_DISABLEBASICOVERHTTP"
	case Qt__BINDF2_DISABLEAUTOCOOKIEHANDLING: // 2
		return "BINDF2_DISABLEAUTOCOOKIEHANDLING"
	case Qt__BINDF2_READ_DATA_GREATER_THAN_4GB: // 4
		return "BINDF2_READ_DATA_GREATER_THAN_4GB"
	case Qt__BINDF2_DISABLE_HTTP_REDIRECT_XSECURITYID: // 8
		return "BINDF2_DISABLE_HTTP_REDIRECT_XSECURITYID"
	case Qt__BINDF2_RESERVED_3: // 536870912
		return "BINDF2_RESERVED_3"
	case Qt__BINDF2_RESERVED_2: // 1073741824
		return "BINDF2_RESERVED_2"
	case Qt__BINDF2_RESERVED_1: // -2147483648
		return "BINDF2_RESERVED_1"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000009 = int // stdglobal
//
const Qt__CIP_DISK_FULL Qt____WIDL_urlmon_generated_name_00000009 = 0

//
const Qt__CIP_ACCESS_DENIED Qt____WIDL_urlmon_generated_name_00000009 = 1

//
const Qt__CIP_NEWER_VERSION_EXISTS Qt____WIDL_urlmon_generated_name_00000009 = 2

//
const Qt__CIP_OLDER_VERSION_EXISTS Qt____WIDL_urlmon_generated_name_00000009 = 3

//
const Qt__CIP_NAME_CONFLICT Qt____WIDL_urlmon_generated_name_00000009 = 4

//
const Qt__CIP_TRUST_VERIFICATION_COMPONENT_MISSING Qt____WIDL_urlmon_generated_name_00000009 = 5

//
const Qt__CIP_EXE_SELF_REGISTERATION_TIMEOUT Qt____WIDL_urlmon_generated_name_00000009 = 6

//
const Qt__CIP_UNSAFE_TO_ABORT Qt____WIDL_urlmon_generated_name_00000009 = 7

//
const Qt__CIP_NEED_REBOOT Qt____WIDL_urlmon_generated_name_00000009 = 8

//
const Qt__CIP_NEED_REBOOT_UI_PERMISSION Qt____WIDL_urlmon_generated_name_00000009 = 9

func __WIDL_urlmon_generated_name_00000009ItemName(val int) string {
	switch val {
	case Qt__CIP_DISK_FULL: // 0
		return "CIP_DISK_FULL"
	case Qt__CIP_ACCESS_DENIED: // 1
		return "CIP_ACCESS_DENIED"
	case Qt__CIP_NEWER_VERSION_EXISTS: // 2
		return "CIP_NEWER_VERSION_EXISTS"
	case Qt__CIP_OLDER_VERSION_EXISTS: // 3
		return "CIP_OLDER_VERSION_EXISTS"
	case Qt__CIP_NAME_CONFLICT: // 4
		return "CIP_NAME_CONFLICT"
	case Qt__CIP_TRUST_VERIFICATION_COMPONENT_MISSING: // 5
		return "CIP_TRUST_VERIFICATION_COMPONENT_MISSING"
	case Qt__CIP_EXE_SELF_REGISTERATION_TIMEOUT: // 6
		return "CIP_EXE_SELF_REGISTERATION_TIMEOUT"
	case Qt__CIP_UNSAFE_TO_ABORT: // 7
		return "CIP_UNSAFE_TO_ABORT"
	case Qt__CIP_NEED_REBOOT: // 8
		return "CIP_NEED_REBOOT"
	case Qt__CIP_NEED_REBOOT_UI_PERMISSION: // 9
		return "CIP_NEED_REBOOT_UI_PERMISSION"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_0000000A = int // stdglobal
//
const Qt__MIMETYPEPROP Qt____WIDL_urlmon_generated_name_0000000A = 0

//
const Qt__USE_SRC_URL Qt____WIDL_urlmon_generated_name_0000000A = 1

//
const Qt__CLASSIDPROP Qt____WIDL_urlmon_generated_name_0000000A = 2

//
const Qt__TRUSTEDDOWNLOADPROP Qt____WIDL_urlmon_generated_name_0000000A = 3

//
const Qt__POPUPLEVELPROP Qt____WIDL_urlmon_generated_name_0000000A = 4

func __WIDL_urlmon_generated_name_0000000AItemName(val int) string {
	switch val {
	case Qt__MIMETYPEPROP: // 0
		return "MIMETYPEPROP"
	case Qt__USE_SRC_URL: // 1
		return "USE_SRC_URL"
	case Qt__CLASSIDPROP: // 2
		return "CLASSIDPROP"
	case Qt__TRUSTEDDOWNLOADPROP: // 3
		return "TRUSTEDDOWNLOADPROP"
	case Qt__POPUPLEVELPROP: // 4
		return "POPUPLEVELPROP"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagBINDSTRING = int // stdglobal
//
const Qt__BINDSTRING_HEADERS Qt__tagBINDSTRING = 1

//
const Qt__BINDSTRING_ACCEPT_MIMES Qt__tagBINDSTRING = 2

//
const Qt__BINDSTRING_EXTRA_URL Qt__tagBINDSTRING = 3

//
const Qt__BINDSTRING_LANGUAGE Qt__tagBINDSTRING = 4

//
const Qt__BINDSTRING_USERNAME Qt__tagBINDSTRING = 5

//
const Qt__BINDSTRING_PASSWORD Qt__tagBINDSTRING = 6

//
const Qt__BINDSTRING_UA_PIXELS Qt__tagBINDSTRING = 7

//
const Qt__BINDSTRING_UA_COLOR Qt__tagBINDSTRING = 8

//
const Qt__BINDSTRING_OS Qt__tagBINDSTRING = 9

//
const Qt__BINDSTRING_USER_AGENT Qt__tagBINDSTRING = 10

//
const Qt__BINDSTRING_ACCEPT_ENCODINGS Qt__tagBINDSTRING = 11

//
const Qt__BINDSTRING_POST_COOKIE Qt__tagBINDSTRING = 12

//
const Qt__BINDSTRING_POST_DATA_MIME Qt__tagBINDSTRING = 13

//
const Qt__BINDSTRING_URL Qt__tagBINDSTRING = 14

//
const Qt__BINDSTRING_IID Qt__tagBINDSTRING = 15

//
const Qt__BINDSTRING_FLAG_BIND_TO_OBJECT Qt__tagBINDSTRING = 16

//
const Qt__BINDSTRING_PTR_BIND_CONTEXT Qt__tagBINDSTRING = 17

//
const Qt__BINDSTRING_XDR_ORIGIN Qt__tagBINDSTRING = 18

//
const Qt__BINDSTRING_DOWNLOADPATH Qt__tagBINDSTRING = 19

//
const Qt__BINDSTRING_ROOTDOC_URL Qt__tagBINDSTRING = 20

//
const Qt__BINDSTRING_INITIAL_FILENAME Qt__tagBINDSTRING = 21

//
const Qt__BINDSTRING_PROXY_USERNAME Qt__tagBINDSTRING = 22

//
const Qt__BINDSTRING_PROXY_PASSWORD Qt__tagBINDSTRING = 23

//
const Qt__BINDSTRING_ENTERPRISE_ID Qt__tagBINDSTRING = 24

func tagBINDSTRINGItemName(val int) string {
	switch val {
	case Qt__BINDSTRING_HEADERS: // 1
		return "BINDSTRING_HEADERS"
	case Qt__BINDSTRING_ACCEPT_MIMES: // 2
		return "BINDSTRING_ACCEPT_MIMES"
	case Qt__BINDSTRING_EXTRA_URL: // 3
		return "BINDSTRING_EXTRA_URL"
	case Qt__BINDSTRING_LANGUAGE: // 4
		return "BINDSTRING_LANGUAGE"
	case Qt__BINDSTRING_USERNAME: // 5
		return "BINDSTRING_USERNAME"
	case Qt__BINDSTRING_PASSWORD: // 6
		return "BINDSTRING_PASSWORD"
	case Qt__BINDSTRING_UA_PIXELS: // 7
		return "BINDSTRING_UA_PIXELS"
	case Qt__BINDSTRING_UA_COLOR: // 8
		return "BINDSTRING_UA_COLOR"
	case Qt__BINDSTRING_OS: // 9
		return "BINDSTRING_OS"
	case Qt__BINDSTRING_USER_AGENT: // 10
		return "BINDSTRING_USER_AGENT"
	case Qt__BINDSTRING_ACCEPT_ENCODINGS: // 11
		return "BINDSTRING_ACCEPT_ENCODINGS"
	case Qt__BINDSTRING_POST_COOKIE: // 12
		return "BINDSTRING_POST_COOKIE"
	case Qt__BINDSTRING_POST_DATA_MIME: // 13
		return "BINDSTRING_POST_DATA_MIME"
	case Qt__BINDSTRING_URL: // 14
		return "BINDSTRING_URL"
	case Qt__BINDSTRING_IID: // 15
		return "BINDSTRING_IID"
	case Qt__BINDSTRING_FLAG_BIND_TO_OBJECT: // 16
		return "BINDSTRING_FLAG_BIND_TO_OBJECT"
	case Qt__BINDSTRING_PTR_BIND_CONTEXT: // 17
		return "BINDSTRING_PTR_BIND_CONTEXT"
	case Qt__BINDSTRING_XDR_ORIGIN: // 18
		return "BINDSTRING_XDR_ORIGIN"
	case Qt__BINDSTRING_DOWNLOADPATH: // 19
		return "BINDSTRING_DOWNLOADPATH"
	case Qt__BINDSTRING_ROOTDOC_URL: // 20
		return "BINDSTRING_ROOTDOC_URL"
	case Qt__BINDSTRING_INITIAL_FILENAME: // 21
		return "BINDSTRING_INITIAL_FILENAME"
	case Qt__BINDSTRING_PROXY_USERNAME: // 22
		return "BINDSTRING_PROXY_USERNAME"
	case Qt__BINDSTRING_PROXY_PASSWORD: // 23
		return "BINDSTRING_PROXY_PASSWORD"
	case Qt__BINDSTRING_ENTERPRISE_ID: // 24
		return "BINDSTRING_ENTERPRISE_ID"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___tagPI_FLAGS = int // stdglobal
//
const Qt__PI_PARSE_URL Qt___tagPI_FLAGS = 1

//
const Qt__PI_FILTER_MODE Qt___tagPI_FLAGS = 2

//
const Qt__PI_FORCE_ASYNC Qt___tagPI_FLAGS = 4

//
const Qt__PI_USE_WORKERTHREAD Qt___tagPI_FLAGS = 8

//
const Qt__PI_MIMEVERIFICATION Qt___tagPI_FLAGS = 16

//
const Qt__PI_CLSIDLOOKUP Qt___tagPI_FLAGS = 32

//
const Qt__PI_DATAPROGRESS Qt___tagPI_FLAGS = 64

//
const Qt__PI_SYNCHRONOUS Qt___tagPI_FLAGS = 128

//
const Qt__PI_APARTMENTTHREADED Qt___tagPI_FLAGS = 256

//
const Qt__PI_CLASSINSTALL Qt___tagPI_FLAGS = 512

//
const Qt__PD_FORCE_SWITCH Qt___tagPI_FLAGS = 65536

func _tagPI_FLAGSItemName(val int) string {
	switch val {
	case Qt__PI_PARSE_URL: // 1
		return "PI_PARSE_URL"
	case Qt__PI_FILTER_MODE: // 2
		return "PI_FILTER_MODE"
	case Qt__PI_FORCE_ASYNC: // 4
		return "PI_FORCE_ASYNC"
	case Qt__PI_USE_WORKERTHREAD: // 8
		return "PI_USE_WORKERTHREAD"
	case Qt__PI_MIMEVERIFICATION: // 16
		return "PI_MIMEVERIFICATION"
	case Qt__PI_CLSIDLOOKUP: // 32
		return "PI_CLSIDLOOKUP"
	case Qt__PI_DATAPROGRESS: // 64
		return "PI_DATAPROGRESS"
	case Qt__PI_SYNCHRONOUS: // 128
		return "PI_SYNCHRONOUS"
	case Qt__PI_APARTMENTTHREADED: // 256
		return "PI_APARTMENTTHREADED"
	case Qt__PI_CLASSINSTALL: // 512
		return "PI_CLASSINSTALL"
	case Qt__PD_FORCE_SWITCH: // 65536
		return "PD_FORCE_SWITCH"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___tagPARSEACTION = int // stdglobal
//
const Qt__PARSE_CANONICALIZE Qt___tagPARSEACTION = 1

//
const Qt__PARSE_FRIENDLY Qt___tagPARSEACTION = 2

//
const Qt__PARSE_SECURITY_URL Qt___tagPARSEACTION = 3

//
const Qt__PARSE_ROOTDOCUMENT Qt___tagPARSEACTION = 4

//
const Qt__PARSE_DOCUMENT Qt___tagPARSEACTION = 5

//
const Qt__PARSE_ANCHOR Qt___tagPARSEACTION = 6

//
const Qt__PARSE_ENCODE Qt___tagPARSEACTION = 7

//
const Qt__PARSE_DECODE Qt___tagPARSEACTION = 8

//
const Qt__PARSE_PATH_FROM_URL Qt___tagPARSEACTION = 9

//
const Qt__PARSE_URL_FROM_PATH Qt___tagPARSEACTION = 10

//
const Qt__PARSE_MIME Qt___tagPARSEACTION = 11

//
const Qt__PARSE_SERVER Qt___tagPARSEACTION = 12

//
const Qt__PARSE_SCHEMA Qt___tagPARSEACTION = 13

//
const Qt__PARSE_SITE Qt___tagPARSEACTION = 14

//
const Qt__PARSE_DOMAIN Qt___tagPARSEACTION = 15

//
const Qt__PARSE_LOCATION Qt___tagPARSEACTION = 16

//
const Qt__PARSE_SECURITY_DOMAIN Qt___tagPARSEACTION = 17

//
const Qt__PARSE_ESCAPE Qt___tagPARSEACTION = 18

//
const Qt__PARSE_UNESCAPE Qt___tagPARSEACTION = 19

func _tagPARSEACTIONItemName(val int) string {
	switch val {
	case Qt__PARSE_CANONICALIZE: // 1
		return "PARSE_CANONICALIZE"
	case Qt__PARSE_FRIENDLY: // 2
		return "PARSE_FRIENDLY"
	case Qt__PARSE_SECURITY_URL: // 3
		return "PARSE_SECURITY_URL"
	case Qt__PARSE_ROOTDOCUMENT: // 4
		return "PARSE_ROOTDOCUMENT"
	case Qt__PARSE_DOCUMENT: // 5
		return "PARSE_DOCUMENT"
	case Qt__PARSE_ANCHOR: // 6
		return "PARSE_ANCHOR"
	case Qt__PARSE_ENCODE: // 7
		return "PARSE_ENCODE"
	case Qt__PARSE_DECODE: // 8
		return "PARSE_DECODE"
	case Qt__PARSE_PATH_FROM_URL: // 9
		return "PARSE_PATH_FROM_URL"
	case Qt__PARSE_URL_FROM_PATH: // 10
		return "PARSE_URL_FROM_PATH"
	case Qt__PARSE_MIME: // 11
		return "PARSE_MIME"
	case Qt__PARSE_SERVER: // 12
		return "PARSE_SERVER"
	case Qt__PARSE_SCHEMA: // 13
		return "PARSE_SCHEMA"
	case Qt__PARSE_SITE: // 14
		return "PARSE_SITE"
	case Qt__PARSE_DOMAIN: // 15
		return "PARSE_DOMAIN"
	case Qt__PARSE_LOCATION: // 16
		return "PARSE_LOCATION"
	case Qt__PARSE_SECURITY_DOMAIN: // 17
		return "PARSE_SECURITY_DOMAIN"
	case Qt__PARSE_ESCAPE: // 18
		return "PARSE_ESCAPE"
	case Qt__PARSE_UNESCAPE: // 19
		return "PARSE_UNESCAPE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___tagPSUACTION = int // stdglobal
//
const Qt__PSU_DEFAULT Qt___tagPSUACTION = 1

//
const Qt__PSU_SECURITY_URL_ONLY Qt___tagPSUACTION = 2

func _tagPSUACTIONItemName(val int) string {
	switch val {
	case Qt__PSU_DEFAULT: // 1
		return "PSU_DEFAULT"
	case Qt__PSU_SECURITY_URL_ONLY: // 2
		return "PSU_SECURITY_URL_ONLY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___tagQUERYOPTION = int // stdglobal
//
const Qt__QUERY_EXPIRATION_DATE Qt___tagQUERYOPTION = 1

//
const Qt__QUERY_TIME_OF_LAST_CHANGE Qt___tagQUERYOPTION = 2

//
const Qt__QUERY_CONTENT_ENCODING Qt___tagQUERYOPTION = 3

//
const Qt__QUERY_CONTENT_TYPE Qt___tagQUERYOPTION = 4

//
const Qt__QUERY_REFRESH Qt___tagQUERYOPTION = 5

//
const Qt__QUERY_RECOMBINE Qt___tagQUERYOPTION = 6

//
const Qt__QUERY_CAN_NAVIGATE Qt___tagQUERYOPTION = 7

//
const Qt__QUERY_USES_NETWORK Qt___tagQUERYOPTION = 8

//
const Qt__QUERY_IS_CACHED Qt___tagQUERYOPTION = 9

//
const Qt__QUERY_IS_INSTALLEDENTRY Qt___tagQUERYOPTION = 10

//
const Qt__QUERY_IS_CACHED_OR_MAPPED Qt___tagQUERYOPTION = 11

//
const Qt__QUERY_USES_CACHE Qt___tagQUERYOPTION = 12

//
const Qt__QUERY_IS_SECURE Qt___tagQUERYOPTION = 13

//
const Qt__QUERY_IS_SAFE Qt___tagQUERYOPTION = 14

//
const Qt__QUERY_USES_HISTORYFOLDER Qt___tagQUERYOPTION = 15

//
const Qt__QUERY_IS_CACHED_AND_USABLE_OFFLINE Qt___tagQUERYOPTION = 16

func _tagQUERYOPTIONItemName(val int) string {
	switch val {
	case Qt__QUERY_EXPIRATION_DATE: // 1
		return "QUERY_EXPIRATION_DATE"
	case Qt__QUERY_TIME_OF_LAST_CHANGE: // 2
		return "QUERY_TIME_OF_LAST_CHANGE"
	case Qt__QUERY_CONTENT_ENCODING: // 3
		return "QUERY_CONTENT_ENCODING"
	case Qt__QUERY_CONTENT_TYPE: // 4
		return "QUERY_CONTENT_TYPE"
	case Qt__QUERY_REFRESH: // 5
		return "QUERY_REFRESH"
	case Qt__QUERY_RECOMBINE: // 6
		return "QUERY_RECOMBINE"
	case Qt__QUERY_CAN_NAVIGATE: // 7
		return "QUERY_CAN_NAVIGATE"
	case Qt__QUERY_USES_NETWORK: // 8
		return "QUERY_USES_NETWORK"
	case Qt__QUERY_IS_CACHED: // 9
		return "QUERY_IS_CACHED"
	case Qt__QUERY_IS_INSTALLEDENTRY: // 10
		return "QUERY_IS_INSTALLEDENTRY"
	case Qt__QUERY_IS_CACHED_OR_MAPPED: // 11
		return "QUERY_IS_CACHED_OR_MAPPED"
	case Qt__QUERY_USES_CACHE: // 12
		return "QUERY_USES_CACHE"
	case Qt__QUERY_IS_SECURE: // 13
		return "QUERY_IS_SECURE"
	case Qt__QUERY_IS_SAFE: // 14
		return "QUERY_IS_SAFE"
	case Qt__QUERY_USES_HISTORYFOLDER: // 15
		return "QUERY_USES_HISTORYFOLDER"
	case Qt__QUERY_IS_CACHED_AND_USABLE_OFFLINE: // 16
		return "QUERY_IS_CACHED_AND_USABLE_OFFLINE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___tagOIBDG_FLAGS = int // stdglobal
//
const Qt__OIBDG_APARTMENTTHREADED Qt___tagOIBDG_FLAGS = 256

//
const Qt__OIBDG_DATAONLY Qt___tagOIBDG_FLAGS = 4096

func _tagOIBDG_FLAGSItemName(val int) string {
	switch val {
	case Qt__OIBDG_APARTMENTTHREADED: // 256
		return "OIBDG_APARTMENTTHREADED"
	case Qt__OIBDG_DATAONLY: // 4096
		return "OIBDG_DATAONLY"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_0000000B = int // stdglobal
//
const Qt__PUAF_DEFAULT Qt____WIDL_urlmon_generated_name_0000000B = 0

//
const Qt__PUAF_NOUI Qt____WIDL_urlmon_generated_name_0000000B = 1

//
const Qt__PUAF_ISFILE Qt____WIDL_urlmon_generated_name_0000000B = 2

//
const Qt__PUAF_WARN_IF_DENIED Qt____WIDL_urlmon_generated_name_0000000B = 4

//
const Qt__PUAF_FORCEUI_FOREGROUND Qt____WIDL_urlmon_generated_name_0000000B = 8

//
const Qt__PUAF_CHECK_TIPS Qt____WIDL_urlmon_generated_name_0000000B = 16

func __WIDL_urlmon_generated_name_0000000BItemName(val int) string {
	switch val {
	case Qt__PUAF_DEFAULT: // 0
		return "PUAF_DEFAULT"
	case Qt__PUAF_NOUI: // 1
		return "PUAF_NOUI"
	case Qt__PUAF_ISFILE: // 2
		return "PUAF_ISFILE"
	case Qt__PUAF_WARN_IF_DENIED: // 4
		return "PUAF_WARN_IF_DENIED"
	case Qt__PUAF_FORCEUI_FOREGROUND: // 8
		return "PUAF_FORCEUI_FOREGROUND"
	case Qt__PUAF_CHECK_TIPS: // 16
		return "PUAF_CHECK_TIPS"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_0000000C = int // stdglobal
//
const Qt__SZM_CREATE Qt____WIDL_urlmon_generated_name_0000000C = 0

//
const Qt__SZM_DELETE Qt____WIDL_urlmon_generated_name_0000000C = 1

func __WIDL_urlmon_generated_name_0000000CItemName(val int) string {
	switch val {
	case Qt__SZM_CREATE: // 0
		return "SZM_CREATE"
	case Qt__SZM_DELETE: // 1
		return "SZM_DELETE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagURLZONE = int // stdglobal
//
const Qt__URLZONE_INVALID Qt__tagURLZONE = -1

//
const Qt__URLZONE_PREDEFINED_MIN Qt__tagURLZONE = 0

//
const Qt__URLZONE_LOCAL_MACHINE Qt__tagURLZONE = 0

//
const Qt__URLZONE_INTRANET Qt__tagURLZONE = 1

//
const Qt__URLZONE_TRUSTED Qt__tagURLZONE = 2

//
const Qt__URLZONE_INTERNET Qt__tagURLZONE = 3

//
const Qt__URLZONE_UNTRUSTED Qt__tagURLZONE = 4

//
const Qt__URLZONE_PREDEFINED_MAX Qt__tagURLZONE = 999

//
const Qt__URLZONE_USER_MIN Qt__tagURLZONE = 1000

//
const Qt__URLZONE_USER_MAX Qt__tagURLZONE = 10000

func tagURLZONEItemName(val int) string {
	switch val {
	case Qt__URLZONE_INVALID: // -1
		return "URLZONE_INVALID"
	case Qt__URLZONE_PREDEFINED_MIN: // 0
		return "URLZONE_PREDEFINED_MIN,URLZONE_LOCAL_MACHINE"
		// case Qt__URLZONE_LOCAL_MACHINE: // 0
		// return ""
	case Qt__URLZONE_INTRANET: // 1
		return "URLZONE_INTRANET"
	case Qt__URLZONE_TRUSTED: // 2
		return "URLZONE_TRUSTED"
	case Qt__URLZONE_INTERNET: // 3
		return "URLZONE_INTERNET"
	case Qt__URLZONE_UNTRUSTED: // 4
		return "URLZONE_UNTRUSTED"
	case Qt__URLZONE_PREDEFINED_MAX: // 999
		return "URLZONE_PREDEFINED_MAX"
	case Qt__URLZONE_USER_MIN: // 1000
		return "URLZONE_USER_MIN"
	case Qt__URLZONE_USER_MAX: // 10000
		return "URLZONE_USER_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagURLTEMPLATE = int // stdglobal
//
const Qt__URLTEMPLATE_CUSTOM Qt__tagURLTEMPLATE = 0

//
const Qt__URLTEMPLATE_PREDEFINED_MIN Qt__tagURLTEMPLATE = 65536

//
const Qt__URLTEMPLATE_LOW Qt__tagURLTEMPLATE = 65536

//
const Qt__URLTEMPLATE_MEDLOW Qt__tagURLTEMPLATE = 66816

//
const Qt__URLTEMPLATE_MEDIUM Qt__tagURLTEMPLATE = 69632

//
const Qt__URLTEMPLATE_MEDHIGH Qt__tagURLTEMPLATE = 70912

//
const Qt__URLTEMPLATE_HIGH Qt__tagURLTEMPLATE = 73728

//
const Qt__URLTEMPLATE_PREDEFINED_MAX Qt__tagURLTEMPLATE = 131072

func tagURLTEMPLATEItemName(val int) string {
	switch val {
	case Qt__URLTEMPLATE_CUSTOM: // 0
		return "URLTEMPLATE_CUSTOM"
	case Qt__URLTEMPLATE_PREDEFINED_MIN: // 65536
		return "URLTEMPLATE_PREDEFINED_MIN,URLTEMPLATE_LOW"
		// case Qt__URLTEMPLATE_LOW: // 65536
		// return ""
	case Qt__URLTEMPLATE_MEDLOW: // 66816
		return "URLTEMPLATE_MEDLOW"
	case Qt__URLTEMPLATE_MEDIUM: // 69632
		return "URLTEMPLATE_MEDIUM"
	case Qt__URLTEMPLATE_MEDHIGH: // 70912
		return "URLTEMPLATE_MEDHIGH"
	case Qt__URLTEMPLATE_HIGH: // 73728
		return "URLTEMPLATE_HIGH"
	case Qt__URLTEMPLATE_PREDEFINED_MAX: // 131072
		return "URLTEMPLATE_PREDEFINED_MAX"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_0000000D = int // stdglobal
//
const Qt__ZAFLAGS_CUSTOM_EDIT Qt____WIDL_urlmon_generated_name_0000000D = 1

//
const Qt__ZAFLAGS_ADD_SITES Qt____WIDL_urlmon_generated_name_0000000D = 2

//
const Qt__ZAFLAGS_REQUIRE_VERIFICATION Qt____WIDL_urlmon_generated_name_0000000D = 4

//
const Qt__ZAFLAGS_INCLUDE_PROXY_OVERRIDE Qt____WIDL_urlmon_generated_name_0000000D = 8

//
const Qt__ZAFLAGS_INCLUDE_INTRANET_SITES Qt____WIDL_urlmon_generated_name_0000000D = 16

//
const Qt__ZAFLAGS_NO_UI Qt____WIDL_urlmon_generated_name_0000000D = 32

//
const Qt__ZAFLAGS_SUPPORTS_VERIFICATION Qt____WIDL_urlmon_generated_name_0000000D = 64

//
const Qt__ZAFLAGS_UNC_AS_INTRANET Qt____WIDL_urlmon_generated_name_0000000D = 128

//
const Qt__ZAFLAGS_DETECT_INTRANET Qt____WIDL_urlmon_generated_name_0000000D = 256

//
const Qt__ZAFLAGS_USE_LOCKED_ZONES Qt____WIDL_urlmon_generated_name_0000000D = 65536

//
const Qt__ZAFLAGS_VERIFY_TEMPLATE_SETTINGS Qt____WIDL_urlmon_generated_name_0000000D = 131072

//
const Qt__ZAFLAGS_NO_CACHE Qt____WIDL_urlmon_generated_name_0000000D = 262144

func __WIDL_urlmon_generated_name_0000000DItemName(val int) string {
	switch val {
	case Qt__ZAFLAGS_CUSTOM_EDIT: // 1
		return "ZAFLAGS_CUSTOM_EDIT"
	case Qt__ZAFLAGS_ADD_SITES: // 2
		return "ZAFLAGS_ADD_SITES"
	case Qt__ZAFLAGS_REQUIRE_VERIFICATION: // 4
		return "ZAFLAGS_REQUIRE_VERIFICATION"
	case Qt__ZAFLAGS_INCLUDE_PROXY_OVERRIDE: // 8
		return "ZAFLAGS_INCLUDE_PROXY_OVERRIDE"
	case Qt__ZAFLAGS_INCLUDE_INTRANET_SITES: // 16
		return "ZAFLAGS_INCLUDE_INTRANET_SITES"
	case Qt__ZAFLAGS_NO_UI: // 32
		return "ZAFLAGS_NO_UI"
	case Qt__ZAFLAGS_SUPPORTS_VERIFICATION: // 64
		return "ZAFLAGS_SUPPORTS_VERIFICATION"
	case Qt__ZAFLAGS_UNC_AS_INTRANET: // 128
		return "ZAFLAGS_UNC_AS_INTRANET"
	case Qt__ZAFLAGS_DETECT_INTRANET: // 256
		return "ZAFLAGS_DETECT_INTRANET"
	case Qt__ZAFLAGS_USE_LOCKED_ZONES: // 65536
		return "ZAFLAGS_USE_LOCKED_ZONES"
	case Qt__ZAFLAGS_VERIFY_TEMPLATE_SETTINGS: // 131072
		return "ZAFLAGS_VERIFY_TEMPLATE_SETTINGS"
	case Qt__ZAFLAGS_NO_CACHE: // 262144
		return "ZAFLAGS_NO_CACHE"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___URLZONEREG = int // stdglobal
//
const Qt__URLZONEREG_DEFAULT Qt___URLZONEREG = 0

//
const Qt__URLZONEREG_HKLM Qt___URLZONEREG = 1

//
const Qt__URLZONEREG_HKCU Qt___URLZONEREG = 2

func _URLZONEREGItemName(val int) string {
	switch val {
	case Qt__URLZONEREG_DEFAULT: // 0
		return "URLZONEREG_DEFAULT"
	case Qt__URLZONEREG_HKLM: // 1
		return "URLZONEREG_HKLM"
	case Qt__URLZONEREG_HKCU: // 2
		return "URLZONEREG_HKCU"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___tagINTERNETFEATURELIST = int // stdglobal
//
const Qt__FEATURE_OBJECT_CACHING Qt___tagINTERNETFEATURELIST = 0

//
const Qt__FEATURE_ZONE_ELEVATION Qt___tagINTERNETFEATURELIST = 1

//
const Qt__FEATURE_MIME_HANDLING Qt___tagINTERNETFEATURELIST = 2

//
const Qt__FEATURE_MIME_SNIFFING Qt___tagINTERNETFEATURELIST = 3

//
const Qt__FEATURE_WINDOW_RESTRICTIONS Qt___tagINTERNETFEATURELIST = 4

//
const Qt__FEATURE_WEBOC_POPUPMANAGEMENT Qt___tagINTERNETFEATURELIST = 5

//
const Qt__FEATURE_BEHAVIORS Qt___tagINTERNETFEATURELIST = 6

//
const Qt__FEATURE_DISABLE_MK_PROTOCOL Qt___tagINTERNETFEATURELIST = 7

//
const Qt__FEATURE_LOCALMACHINE_LOCKDOWN Qt___tagINTERNETFEATURELIST = 8

//
const Qt__FEATURE_SECURITYBAND Qt___tagINTERNETFEATURELIST = 9

//
const Qt__FEATURE_RESTRICT_ACTIVEXINSTALL Qt___tagINTERNETFEATURELIST = 10

//
const Qt__FEATURE_VALIDATE_NAVIGATE_URL Qt___tagINTERNETFEATURELIST = 11

//
const Qt__FEATURE_RESTRICT_FILEDOWNLOAD Qt___tagINTERNETFEATURELIST = 12

//
const Qt__FEATURE_ADDON_MANAGEMENT Qt___tagINTERNETFEATURELIST = 13

//
const Qt__FEATURE_PROTOCOL_LOCKDOWN Qt___tagINTERNETFEATURELIST = 14

//
const Qt__FEATURE_HTTP_USERNAME_PASSWORD_DISABLE Qt___tagINTERNETFEATURELIST = 15

//
const Qt__FEATURE_SAFE_BINDTOOBJECT Qt___tagINTERNETFEATURELIST = 16

//
const Qt__FEATURE_UNC_SAVEDFILECHECK Qt___tagINTERNETFEATURELIST = 17

//
const Qt__FEATURE_GET_URL_DOM_FILEPATH_UNENCODED Qt___tagINTERNETFEATURELIST = 18

//
const Qt__FEATURE_TABBED_BROWSING Qt___tagINTERNETFEATURELIST = 19

//
const Qt__FEATURE_SSLUX Qt___tagINTERNETFEATURELIST = 20

//
const Qt__FEATURE_DISABLE_NAVIGATION_SOUNDS Qt___tagINTERNETFEATURELIST = 21

//
const Qt__FEATURE_DISABLE_LEGACY_COMPRESSION Qt___tagINTERNETFEATURELIST = 22

//
const Qt__FEATURE_FORCE_ADDR_AND_STATUS Qt___tagINTERNETFEATURELIST = 23

//
const Qt__FEATURE_XMLHTTP Qt___tagINTERNETFEATURELIST = 24

//
const Qt__FEATURE_DISABLE_TELNET_PROTOCOL Qt___tagINTERNETFEATURELIST = 25

//
const Qt__FEATURE_FEEDS Qt___tagINTERNETFEATURELIST = 26

//
const Qt__FEATURE_BLOCK_INPUT_PROMPTS Qt___tagINTERNETFEATURELIST = 27

//
const Qt__FEATURE_ENTRY_COUNT Qt___tagINTERNETFEATURELIST = 28

func _tagINTERNETFEATURELISTItemName(val int) string {
	switch val {
	case Qt__FEATURE_OBJECT_CACHING: // 0
		return "FEATURE_OBJECT_CACHING"
	case Qt__FEATURE_ZONE_ELEVATION: // 1
		return "FEATURE_ZONE_ELEVATION"
	case Qt__FEATURE_MIME_HANDLING: // 2
		return "FEATURE_MIME_HANDLING"
	case Qt__FEATURE_MIME_SNIFFING: // 3
		return "FEATURE_MIME_SNIFFING"
	case Qt__FEATURE_WINDOW_RESTRICTIONS: // 4
		return "FEATURE_WINDOW_RESTRICTIONS"
	case Qt__FEATURE_WEBOC_POPUPMANAGEMENT: // 5
		return "FEATURE_WEBOC_POPUPMANAGEMENT"
	case Qt__FEATURE_BEHAVIORS: // 6
		return "FEATURE_BEHAVIORS"
	case Qt__FEATURE_DISABLE_MK_PROTOCOL: // 7
		return "FEATURE_DISABLE_MK_PROTOCOL"
	case Qt__FEATURE_LOCALMACHINE_LOCKDOWN: // 8
		return "FEATURE_LOCALMACHINE_LOCKDOWN"
	case Qt__FEATURE_SECURITYBAND: // 9
		return "FEATURE_SECURITYBAND"
	case Qt__FEATURE_RESTRICT_ACTIVEXINSTALL: // 10
		return "FEATURE_RESTRICT_ACTIVEXINSTALL"
	case Qt__FEATURE_VALIDATE_NAVIGATE_URL: // 11
		return "FEATURE_VALIDATE_NAVIGATE_URL"
	case Qt__FEATURE_RESTRICT_FILEDOWNLOAD: // 12
		return "FEATURE_RESTRICT_FILEDOWNLOAD"
	case Qt__FEATURE_ADDON_MANAGEMENT: // 13
		return "FEATURE_ADDON_MANAGEMENT"
	case Qt__FEATURE_PROTOCOL_LOCKDOWN: // 14
		return "FEATURE_PROTOCOL_LOCKDOWN"
	case Qt__FEATURE_HTTP_USERNAME_PASSWORD_DISABLE: // 15
		return "FEATURE_HTTP_USERNAME_PASSWORD_DISABLE"
	case Qt__FEATURE_SAFE_BINDTOOBJECT: // 16
		return "FEATURE_SAFE_BINDTOOBJECT"
	case Qt__FEATURE_UNC_SAVEDFILECHECK: // 17
		return "FEATURE_UNC_SAVEDFILECHECK"
	case Qt__FEATURE_GET_URL_DOM_FILEPATH_UNENCODED: // 18
		return "FEATURE_GET_URL_DOM_FILEPATH_UNENCODED"
	case Qt__FEATURE_TABBED_BROWSING: // 19
		return "FEATURE_TABBED_BROWSING"
	case Qt__FEATURE_SSLUX: // 20
		return "FEATURE_SSLUX"
	case Qt__FEATURE_DISABLE_NAVIGATION_SOUNDS: // 21
		return "FEATURE_DISABLE_NAVIGATION_SOUNDS"
	case Qt__FEATURE_DISABLE_LEGACY_COMPRESSION: // 22
		return "FEATURE_DISABLE_LEGACY_COMPRESSION"
	case Qt__FEATURE_FORCE_ADDR_AND_STATUS: // 23
		return "FEATURE_FORCE_ADDR_AND_STATUS"
	case Qt__FEATURE_XMLHTTP: // 24
		return "FEATURE_XMLHTTP"
	case Qt__FEATURE_DISABLE_TELNET_PROTOCOL: // 25
		return "FEATURE_DISABLE_TELNET_PROTOCOL"
	case Qt__FEATURE_FEEDS: // 26
		return "FEATURE_FEEDS"
	case Qt__FEATURE_BLOCK_INPUT_PROMPTS: // 27
		return "FEATURE_BLOCK_INPUT_PROMPTS"
	case Qt__FEATURE_ENTRY_COUNT: // 28
		return "FEATURE_ENTRY_COUNT"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_0000000E = int // stdglobal
//
const Qt__Uri_PROPERTY_ABSOLUTE_URI Qt____WIDL_urlmon_generated_name_0000000E = 0

//
const Qt__Uri_PROPERTY_STRING_START Qt____WIDL_urlmon_generated_name_0000000E = 0

//
const Qt__Uri_PROPERTY_AUTHORITY Qt____WIDL_urlmon_generated_name_0000000E = 1

//
const Qt__Uri_PROPERTY_DISPLAY_URI Qt____WIDL_urlmon_generated_name_0000000E = 2

//
const Qt__Uri_PROPERTY_DOMAIN Qt____WIDL_urlmon_generated_name_0000000E = 3

//
const Qt__Uri_PROPERTY_EXTENSION Qt____WIDL_urlmon_generated_name_0000000E = 4

//
const Qt__Uri_PROPERTY_FRAGMENT Qt____WIDL_urlmon_generated_name_0000000E = 5

//
const Qt__Uri_PROPERTY_HOST Qt____WIDL_urlmon_generated_name_0000000E = 6

//
const Qt__Uri_PROPERTY_PASSWORD Qt____WIDL_urlmon_generated_name_0000000E = 7

//
const Qt__Uri_PROPERTY_PATH Qt____WIDL_urlmon_generated_name_0000000E = 8

//
const Qt__Uri_PROPERTY_PATH_AND_QUERY Qt____WIDL_urlmon_generated_name_0000000E = 9

//
const Qt__Uri_PROPERTY_QUERY Qt____WIDL_urlmon_generated_name_0000000E = 10

//
const Qt__Uri_PROPERTY_RAW_URI Qt____WIDL_urlmon_generated_name_0000000E = 11

//
const Qt__Uri_PROPERTY_SCHEME_NAME Qt____WIDL_urlmon_generated_name_0000000E = 12

//
const Qt__Uri_PROPERTY_USER_INFO Qt____WIDL_urlmon_generated_name_0000000E = 13

//
const Qt__Uri_PROPERTY_USER_NAME Qt____WIDL_urlmon_generated_name_0000000E = 14

//
const Qt__Uri_PROPERTY_STRING_LAST Qt____WIDL_urlmon_generated_name_0000000E = 14

//
const Qt__Uri_PROPERTY_HOST_TYPE Qt____WIDL_urlmon_generated_name_0000000E = 15

//
const Qt__Uri_PROPERTY_DWORD_START Qt____WIDL_urlmon_generated_name_0000000E = 15

//
const Qt__Uri_PROPERTY_PORT Qt____WIDL_urlmon_generated_name_0000000E = 16

//
const Qt__Uri_PROPERTY_SCHEME Qt____WIDL_urlmon_generated_name_0000000E = 17

//
const Qt__Uri_PROPERTY_ZONE Qt____WIDL_urlmon_generated_name_0000000E = 18

//
const Qt__Uri_PROPERTY_DWORD_LAST Qt____WIDL_urlmon_generated_name_0000000E = 18

func __WIDL_urlmon_generated_name_0000000EItemName(val int) string {
	switch val {
	case Qt__Uri_PROPERTY_ABSOLUTE_URI: // 0
		return "Uri_PROPERTY_ABSOLUTE_URI,Uri_PROPERTY_STRING_START"
		// case Qt__Uri_PROPERTY_STRING_START: // 0
		// return ""
	case Qt__Uri_PROPERTY_AUTHORITY: // 1
		return "Uri_PROPERTY_AUTHORITY"
	case Qt__Uri_PROPERTY_DISPLAY_URI: // 2
		return "Uri_PROPERTY_DISPLAY_URI"
	case Qt__Uri_PROPERTY_DOMAIN: // 3
		return "Uri_PROPERTY_DOMAIN"
	case Qt__Uri_PROPERTY_EXTENSION: // 4
		return "Uri_PROPERTY_EXTENSION"
	case Qt__Uri_PROPERTY_FRAGMENT: // 5
		return "Uri_PROPERTY_FRAGMENT"
	case Qt__Uri_PROPERTY_HOST: // 6
		return "Uri_PROPERTY_HOST"
	case Qt__Uri_PROPERTY_PASSWORD: // 7
		return "Uri_PROPERTY_PASSWORD"
	case Qt__Uri_PROPERTY_PATH: // 8
		return "Uri_PROPERTY_PATH"
	case Qt__Uri_PROPERTY_PATH_AND_QUERY: // 9
		return "Uri_PROPERTY_PATH_AND_QUERY"
	case Qt__Uri_PROPERTY_QUERY: // 10
		return "Uri_PROPERTY_QUERY"
	case Qt__Uri_PROPERTY_RAW_URI: // 11
		return "Uri_PROPERTY_RAW_URI"
	case Qt__Uri_PROPERTY_SCHEME_NAME: // 12
		return "Uri_PROPERTY_SCHEME_NAME"
	case Qt__Uri_PROPERTY_USER_INFO: // 13
		return "Uri_PROPERTY_USER_INFO"
	case Qt__Uri_PROPERTY_USER_NAME: // 14
		return "Uri_PROPERTY_USER_NAME,Uri_PROPERTY_STRING_LAST"
		// case Qt__Uri_PROPERTY_STRING_LAST: // 14
		// return ""
	case Qt__Uri_PROPERTY_HOST_TYPE: // 15
		return "Uri_PROPERTY_HOST_TYPE,Uri_PROPERTY_DWORD_START"
		// case Qt__Uri_PROPERTY_DWORD_START: // 15
		// return ""
	case Qt__Uri_PROPERTY_PORT: // 16
		return "Uri_PROPERTY_PORT"
	case Qt__Uri_PROPERTY_SCHEME: // 17
		return "Uri_PROPERTY_SCHEME"
	case Qt__Uri_PROPERTY_ZONE: // 18
		return "Uri_PROPERTY_ZONE,Uri_PROPERTY_DWORD_LAST"
		// case Qt__Uri_PROPERTY_DWORD_LAST: // 18
		// return ""
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_0000000F = int // stdglobal
//
const Qt__Uri_HOST_UNKNOWN Qt____WIDL_urlmon_generated_name_0000000F = 0

//
const Qt__Uri_HOST_DNS Qt____WIDL_urlmon_generated_name_0000000F = 1

//
const Qt__Uri_HOST_IPV4 Qt____WIDL_urlmon_generated_name_0000000F = 2

//
const Qt__Uri_HOST_IPV6 Qt____WIDL_urlmon_generated_name_0000000F = 3

//
const Qt__Uri_HOST_IDN Qt____WIDL_urlmon_generated_name_0000000F = 4

func __WIDL_urlmon_generated_name_0000000FItemName(val int) string {
	switch val {
	case Qt__Uri_HOST_UNKNOWN: // 0
		return "Uri_HOST_UNKNOWN"
	case Qt__Uri_HOST_DNS: // 1
		return "Uri_HOST_DNS"
	case Qt__Uri_HOST_IPV4: // 2
		return "Uri_HOST_IPV4"
	case Qt__Uri_HOST_IPV6: // 3
		return "Uri_HOST_IPV6"
	case Qt__Uri_HOST_IDN: // 4
		return "Uri_HOST_IDN"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt____WIDL_urlmon_generated_name_00000010 = int // stdglobal
//
const Qt__BINDHANDLETYPES_APPCACHE Qt____WIDL_urlmon_generated_name_00000010 = 0

//
const Qt__BINDHANDLETYPES_DEPENDENCY Qt____WIDL_urlmon_generated_name_00000010 = 1

//
const Qt__BINDHANDLETYPES_COUNT Qt____WIDL_urlmon_generated_name_00000010 = 2

func __WIDL_urlmon_generated_name_00000010ItemName(val int) string {
	switch val {
	case Qt__BINDHANDLETYPES_APPCACHE: // 0
		return "BINDHANDLETYPES_APPCACHE"
	case Qt__BINDHANDLETYPES_DEPENDENCY: // 1
		return "BINDHANDLETYPES_DEPENDENCY"
	case Qt__BINDHANDLETYPES_COUNT: // 2
		return "BINDHANDLETYPES_COUNT"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt__tagREGKIND = int // stdglobal
//
const Qt__REGKIND_DEFAULT Qt__tagREGKIND = 0

//
const Qt__REGKIND_REGISTER Qt__tagREGKIND = 1

//
const Qt__REGKIND_NONE Qt__tagREGKIND = 2

func tagREGKINDItemName(val int) string {
	switch val {
	case Qt__REGKIND_DEFAULT: // 0
		return "REGKIND_DEFAULT"
	case Qt__REGKIND_REGISTER: // 1
		return "REGKIND_REGISTER"
	case Qt__REGKIND_NONE: // 2
		return "REGKIND_NONE"
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
type Qt___SC_STATUS_TYPE = int // stdglobal
//
const Qt__SC_STATUS_PROCESS_INFO Qt___SC_STATUS_TYPE = 0

func _SC_STATUS_TYPEItemName(val int) string {
	switch val {
	case Qt__SC_STATUS_PROCESS_INFO: // 0
		return "SC_STATUS_PROCESS_INFO"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___SC_ENUM_TYPE = int // stdglobal
//
const Qt__SC_ENUM_PROCESS_INFO Qt___SC_ENUM_TYPE = 0

func _SC_ENUM_TYPEItemName(val int) string {
	switch val {
	case Qt__SC_ENUM_PROCESS_INFO: // 0
		return "SC_ENUM_PROCESS_INFO"
	default:
		return fmt.Sprintf("%d", val)
	}
}

/*


 */
type Qt___SC_ACTION_TYPE = int // stdglobal
//
const Qt__SC_ACTION_NONE Qt___SC_ACTION_TYPE = 0

//
const Qt__SC_ACTION_RESTART Qt___SC_ACTION_TYPE = 1

//
const Qt__SC_ACTION_REBOOT Qt___SC_ACTION_TYPE = 2

//
const Qt__SC_ACTION_RUN_COMMAND Qt___SC_ACTION_TYPE = 3

func _SC_ACTION_TYPEItemName(val int) string {
	switch val {
	case Qt__SC_ACTION_NONE: // 0
		return "SC_ACTION_NONE"
	case Qt__SC_ACTION_RESTART: // 1
		return "SC_ACTION_RESTART"
	case Qt__SC_ACTION_REBOOT: // 2
		return "SC_ACTION_REBOOT"
	case Qt__SC_ACTION_RUN_COMMAND: // 3
		return "SC_ACTION_RUN_COMMAND"
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

//  body block end

//  keep block begin

func make_sure_usepkg_qnamespace() {
	if false {
		fmt.Println(123)
	}
}

//  keep block end
