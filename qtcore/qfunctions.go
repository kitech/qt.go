package qtcore

import "unsafe"
import "github.com/kitech/qt.go/qtrt"

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
}

//  header block end

//  body block begin
// /usr/include/qt/QtCore/qnamespace.h:1754
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TimerType)

/*

 */
func Qt_getEnumMetaObject(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_9TimerTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1725
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::SortOrder)

/*

 */
func Qt_getEnumMetaObject_1(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_9SortOrderE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1710
// index:2
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::DayOfWeek)

/*

 */
func Qt_getEnumMetaObject_2(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_9DayOfWeekE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1679
// index:3
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ArrowType)

/*

 */
func Qt_getEnumMetaObject_3(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_9ArrowTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1709
// index:4
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TimeSpec)

/*

 */
func Qt_getEnumMetaObject_4(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8TimeSpecE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1698
// index:5
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TextFlag)

/*

 */
func Qt_getEnumMetaObject_5(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8TextFlagE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1689
// index:6
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::SizeMode)

/*

 */
func Qt_getEnumMetaObject_6(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8SizeModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1694
// index:7
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::SizeHint)

/*

 */
func Qt_getEnumMetaObject_7(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8SizeHintE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1681
// index:8
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::PenStyle)

/*

 */
func Qt_getEnumMetaObject_8(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8PenStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1686
// index:9
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::MaskMode)

/*

 */
func Qt_getEnumMetaObject_9(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8MaskModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1685
// index:10
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::FillRule)

/*

 */
func Qt_getEnumMetaObject_10(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_8FillRuleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1691
// index:11
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Corner)

/*

 */
func Qt_getEnumMetaObject_11(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_6CornerE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1687
// index:12
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::BGMode)

/*

 */
func Qt_getEnumMetaObject_12(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_6BGModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1692
// index:13
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Edge)

/*

 */
func Qt_getEnumMetaObject_13(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_4EdgeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1690
// index:14
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Axis)

/*

 */
func Qt_getEnumMetaObject_14(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_4AxisE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1716
// index:15
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Key)

/*

 */
func Qt_getEnumMetaObject_15(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_3KeyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1721
// index:16
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ItemSelectionOperation)

/*

 */
func Qt_getEnumMetaObject_16(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_22ItemSelectionOperationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1734
// index:17
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ApplicationAttribute)

/*

 */
func Qt_getEnumMetaObject_17(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_20ApplicationAttributeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1718
// index:18
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TextInteractionFlag)

/*

 */
func Qt_getEnumMetaObject_18(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_19TextInteractionFlagE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1714
// index:19
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TransformationMode)

/*

 */
func Qt_getEnumMetaObject_19(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_18TransformationModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1744
// index:20
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ScreenOrientation)

/*

 */
func Qt_getEnumMetaObject_20(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1751
// index:21
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::NativeGestureType)

/*

 */
func Qt_getEnumMetaObject_21(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_17NativeGestureTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1720
// index:22
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ItemSelectionMode)

/*

 */
func Qt_getEnumMetaObject_22(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1678
// index:23
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ContextMenuPolicy)

/*

 */
func Qt_getEnumMetaObject_23(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_17ContextMenuPolicyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1759
// index:24
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TabFocusBehavior)

/*

 */
func Qt_getEnumMetaObject_24(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_16TabFocusBehaviorE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1756
// index:25
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::MouseEventSource)

/*

 */
func Qt_getEnumMetaObject_25(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_16MouseEventSourceE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1739
// index:26
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::InputMethodQuery)

/*

 */
func Qt_getEnumMetaObject_26(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1747
// index:27
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ApplicationState)

/*

 */
func Qt_getEnumMetaObject_27(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_16ApplicationStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1733
// index:28
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::WidgetAttribute)

/*

 */
func Qt_getEnumMetaObject_28(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15WidgetAttributeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1680
// index:29
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ToolButtonStyle)

/*

 */
func Qt_getEnumMetaObject_29(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1717
// index:30
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ShortcutContext)

/*

 */
func Qt_getEnumMetaObject_30(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15ShortcutContextE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1676
// index:31
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ScrollBarPolicy)

/*

 */
func Qt_getEnumMetaObject_31(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15ScrollBarPolicyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1693
// index:32
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::LayoutDirection)

/*

 */
func Qt_getEnumMetaObject_32(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1738
// index:33
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::InputMethodHint)

/*

 */
func Qt_getEnumMetaObject_33(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15InputMethodHintE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1753
// index:34
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::CursorMoveStyle)

/*

 */
func Qt_getEnumMetaObject_34(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1726
// index:35
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::CaseSensitivity)

/*

 */
func Qt_getEnumMetaObject_35(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1713
// index:36
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::AspectRatioMode)

/*

 */
func Qt_getEnumMetaObject_36(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1732
// index:37
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::WindowModality)

/*

 */
func Qt_getEnumMetaObject_37(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_14WindowModalityE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1757
// index:38
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::MouseEventFlag)

/*

 */
func Qt_getEnumMetaObject_38(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_14MouseEventFlagE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1704
// index:39
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::DockWidgetArea)

/*

 */
func Qt_getEnumMetaObject_39(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1746
// index:40
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ConnectionType)

/*

 */
func Qt_getEnumMetaObject_40(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1707
// index:41
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TextElideMode)

/*

 */
func Qt_getEnumMetaObject_41(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_13TextElideModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1688
// index:42
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ClipOperation)

/*

 */
func Qt_getEnumMetaObject_42(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_13ClipOperationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1683
// index:43
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::PenJoinStyle)

/*

 */
func Qt_getEnumMetaObject_43(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1724
// index:44
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ItemDataRole)

/*

 */
func Qt_getEnumMetaObject_44(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_12ItemDataRoleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1749
// index:45
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::GestureState)

/*

 */
func Qt_getEnumMetaObject_45(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_12GestureStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1741
// index:46
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::EnterKeyType)

/*

 */
func Qt_getEnumMetaObject_46(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_12EnterKeyTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1758
// index:47
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ChecksumType)

/*

 */
func Qt_getEnumMetaObject_47(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_12ChecksumTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1731
// index:48
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::WindowState)

/*

 */
func Qt_getEnumMetaObject_48(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11WindowStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1705
// index:49
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ToolBarArea)

/*

 */
func Qt_getEnumMetaObject_49(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1755
// index:50
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ScrollPhase)

/*

 */
func Qt_getEnumMetaObject_50(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11ScrollPhaseE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1682
// index:51
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::PenCapStyle)

/*

 */
func Qt_getEnumMetaObject_51(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11PenCapStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1695
// index:52
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Orientation)

/*

 */
func Qt_getEnumMetaObject_52(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11OrientationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1712
// index:53
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::GlobalColor)

/*

 */
func Qt_getEnumMetaObject_53(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11GlobalColorE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1750
// index:54
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::GestureType)

/*

 */
func Qt_getEnumMetaObject_54(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11GestureTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1737
// index:55
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::FocusReason)

/*

 */
func Qt_getEnumMetaObject_55(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11FocusReasonE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1677
// index:56
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::FocusPolicy)

/*

 */
func Qt_getEnumMetaObject_56(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11FocusPolicyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1711
// index:57
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::CursorShape)

/*

 */
func Qt_getEnumMetaObject_57(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_11CursorShapeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1730
// index:58
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::WindowType)

/*

 */
func Qt_getEnumMetaObject_58(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_10WindowTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1706
// index:59
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TextFormat)

/*

 */
func Qt_getEnumMetaObject_59(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_10TextFormatE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1696
// index:60
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::DropAction)

/*

 */
func Qt_getEnumMetaObject_60(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_10DropActionE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1708
// index:61
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::DateFormat)

/*

 */
func Qt_getEnumMetaObject_61(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_10DateFormatE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1723
// index:62
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::CheckState)

/*

 */
func Qt_getEnumMetaObject_62(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_10CheckStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1684
// index:63
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::BrushStyle)

/*

 */
func Qt_getEnumMetaObject_63(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectENS_10BrushStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1727
// index:64
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::MatchFlags)

/*

 */
func Qt_getEnumMetaObject_64(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_9MatchFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1722
// index:65
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ItemFlags)

/*

 */
func Qt_getEnumMetaObject_65(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_8ItemFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1701
// index:66
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Edges)

/*

 */
func Qt_getEnumMetaObject_66(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_4EdgeEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1719
// index:67
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TextInteractionFlags)

/*

 */
func Qt_getEnumMetaObject_67(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1715
// index:68
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ImageConversionFlags)

/*

 */
func Qt_getEnumMetaObject_68(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1745
// index:69
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ScreenOrientations)

/*

 */
func Qt_getEnumMetaObject_69(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_17ScreenOrientationEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1728
// index:70
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::KeyboardModifiers)

/*

 */
func Qt_getEnumMetaObject_70(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1742
// index:71
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::InputMethodQueries)

/*

 */
func Qt_getEnumMetaObject_71(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1743
// index:72
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::TouchPointStates)

/*

 */
func Qt_getEnumMetaObject_72(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_15TouchPointStateEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1740
// index:73
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::InputMethodHints)

/*

 */
func Qt_getEnumMetaObject_73(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1702
// index:74
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::DockWidgetAreas)

/*

 */
func Qt_getEnumMetaObject_74(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_14DockWidgetAreaEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1697
// index:75
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Alignment)

/*

 */
func Qt_getEnumMetaObject_75(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1736
// index:76
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::WindowStates)

/*

 */
func Qt_getEnumMetaObject_76(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_11WindowStateEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1703
// index:77
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::ToolBarAreas)

/*

 */
func Qt_getEnumMetaObject_77(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_11ToolBarAreaEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1699
// index:78
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::Orientations)

/*

 */
func Qt_getEnumMetaObject_78(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_11OrientationEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1729
// index:79
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::MouseButtons)

/*

 */
func Qt_getEnumMetaObject_79(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_11MouseButtonEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1735
// index:80
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::WindowFlags)

/*

 */
func Qt_getEnumMetaObject_80(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_10WindowTypeEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1700
// index:81
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getEnumMetaObject(Qt::DropActions)

/*

 */
func Qt_getEnumMetaObject_81(arg0 int) *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20qt_getEnumMetaObjectE6QFlagsINS_10DropActionEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qnamespace.h:1754
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TimerType)

/*

 */
func Qt_getEnumName(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_9TimerTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1725
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::SortOrder)

/*

 */
func Qt_getEnumName_1(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_9SortOrderE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1710
// index:2
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::DayOfWeek)

/*

 */
func Qt_getEnumName_2(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_9DayOfWeekE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1679
// index:3
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ArrowType)

/*

 */
func Qt_getEnumName_3(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_9ArrowTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1709
// index:4
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TimeSpec)

/*

 */
func Qt_getEnumName_4(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8TimeSpecE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1698
// index:5
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TextFlag)

/*

 */
func Qt_getEnumName_5(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8TextFlagE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1689
// index:6
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::SizeMode)

/*

 */
func Qt_getEnumName_6(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8SizeModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1694
// index:7
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::SizeHint)

/*

 */
func Qt_getEnumName_7(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8SizeHintE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1681
// index:8
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::PenStyle)

/*

 */
func Qt_getEnumName_8(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8PenStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1686
// index:9
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::MaskMode)

/*

 */
func Qt_getEnumName_9(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8MaskModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1685
// index:10
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::FillRule)

/*

 */
func Qt_getEnumName_10(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_8FillRuleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1691
// index:11
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Corner)

/*

 */
func Qt_getEnumName_11(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_6CornerE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1687
// index:12
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::BGMode)

/*

 */
func Qt_getEnumName_12(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_6BGModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1692
// index:13
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Edge)

/*

 */
func Qt_getEnumName_13(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_4EdgeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1690
// index:14
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Axis)

/*

 */
func Qt_getEnumName_14(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_4AxisE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1716
// index:15
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Key)

/*

 */
func Qt_getEnumName_15(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_3KeyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1721
// index:16
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ItemSelectionOperation)

/*

 */
func Qt_getEnumName_16(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_22ItemSelectionOperationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1734
// index:17
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ApplicationAttribute)

/*

 */
func Qt_getEnumName_17(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_20ApplicationAttributeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1718
// index:18
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TextInteractionFlag)

/*

 */
func Qt_getEnumName_18(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_19TextInteractionFlagE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1714
// index:19
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TransformationMode)

/*

 */
func Qt_getEnumName_19(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_18TransformationModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1744
// index:20
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ScreenOrientation)

/*

 */
func Qt_getEnumName_20(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_17ScreenOrientationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1751
// index:21
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::NativeGestureType)

/*

 */
func Qt_getEnumName_21(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_17NativeGestureTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1720
// index:22
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ItemSelectionMode)

/*

 */
func Qt_getEnumName_22(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_17ItemSelectionModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1678
// index:23
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ContextMenuPolicy)

/*

 */
func Qt_getEnumName_23(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_17ContextMenuPolicyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1759
// index:24
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TabFocusBehavior)

/*

 */
func Qt_getEnumName_24(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_16TabFocusBehaviorE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1756
// index:25
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::MouseEventSource)

/*

 */
func Qt_getEnumName_25(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_16MouseEventSourceE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1739
// index:26
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::InputMethodQuery)

/*

 */
func Qt_getEnumName_26(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_16InputMethodQueryE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1747
// index:27
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ApplicationState)

/*

 */
func Qt_getEnumName_27(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_16ApplicationStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1733
// index:28
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::WidgetAttribute)

/*

 */
func Qt_getEnumName_28(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15WidgetAttributeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1680
// index:29
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ToolButtonStyle)

/*

 */
func Qt_getEnumName_29(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15ToolButtonStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1717
// index:30
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ShortcutContext)

/*

 */
func Qt_getEnumName_30(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15ShortcutContextE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1676
// index:31
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ScrollBarPolicy)

/*

 */
func Qt_getEnumName_31(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15ScrollBarPolicyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1693
// index:32
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::LayoutDirection)

/*

 */
func Qt_getEnumName_32(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15LayoutDirectionE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1738
// index:33
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::InputMethodHint)

/*

 */
func Qt_getEnumName_33(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15InputMethodHintE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1753
// index:34
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::CursorMoveStyle)

/*

 */
func Qt_getEnumName_34(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15CursorMoveStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1726
// index:35
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::CaseSensitivity)

/*

 */
func Qt_getEnumName_35(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15CaseSensitivityE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1713
// index:36
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::AspectRatioMode)

/*

 */
func Qt_getEnumName_36(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_15AspectRatioModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1732
// index:37
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::WindowModality)

/*

 */
func Qt_getEnumName_37(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_14WindowModalityE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1757
// index:38
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::MouseEventFlag)

/*

 */
func Qt_getEnumName_38(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_14MouseEventFlagE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1704
// index:39
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::DockWidgetArea)

/*

 */
func Qt_getEnumName_39(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_14DockWidgetAreaE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1746
// index:40
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ConnectionType)

/*

 */
func Qt_getEnumName_40(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_14ConnectionTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1707
// index:41
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TextElideMode)

/*

 */
func Qt_getEnumName_41(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_13TextElideModeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1688
// index:42
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ClipOperation)

/*

 */
func Qt_getEnumName_42(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_13ClipOperationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1683
// index:43
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::PenJoinStyle)

/*

 */
func Qt_getEnumName_43(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_12PenJoinStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1724
// index:44
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ItemDataRole)

/*

 */
func Qt_getEnumName_44(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_12ItemDataRoleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1749
// index:45
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::GestureState)

/*

 */
func Qt_getEnumName_45(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_12GestureStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1741
// index:46
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::EnterKeyType)

/*

 */
func Qt_getEnumName_46(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_12EnterKeyTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1758
// index:47
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ChecksumType)

/*

 */
func Qt_getEnumName_47(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_12ChecksumTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1731
// index:48
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::WindowState)

/*

 */
func Qt_getEnumName_48(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11WindowStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1705
// index:49
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ToolBarArea)

/*

 */
func Qt_getEnumName_49(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11ToolBarAreaE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1755
// index:50
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ScrollPhase)

/*

 */
func Qt_getEnumName_50(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11ScrollPhaseE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1682
// index:51
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::PenCapStyle)

/*

 */
func Qt_getEnumName_51(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11PenCapStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1695
// index:52
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Orientation)

/*

 */
func Qt_getEnumName_52(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11OrientationE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1712
// index:53
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::GlobalColor)

/*

 */
func Qt_getEnumName_53(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11GlobalColorE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1750
// index:54
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::GestureType)

/*

 */
func Qt_getEnumName_54(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11GestureTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1737
// index:55
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::FocusReason)

/*

 */
func Qt_getEnumName_55(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11FocusReasonE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1677
// index:56
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::FocusPolicy)

/*

 */
func Qt_getEnumName_56(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11FocusPolicyE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1711
// index:57
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::CursorShape)

/*

 */
func Qt_getEnumName_57(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_11CursorShapeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1730
// index:58
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::WindowType)

/*

 */
func Qt_getEnumName_58(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_10WindowTypeE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1706
// index:59
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TextFormat)

/*

 */
func Qt_getEnumName_59(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_10TextFormatE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1696
// index:60
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::DropAction)

/*

 */
func Qt_getEnumName_60(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_10DropActionE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1708
// index:61
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::DateFormat)

/*

 */
func Qt_getEnumName_61(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_10DateFormatE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1723
// index:62
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::CheckState)

/*

 */
func Qt_getEnumName_62(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_10CheckStateE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1684
// index:63
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::BrushStyle)

/*

 */
func Qt_getEnumName_63(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameENS_10BrushStyleE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1727
// index:64
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::MatchFlags)

/*

 */
func Qt_getEnumName_64(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_9MatchFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1722
// index:65
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ItemFlags)

/*

 */
func Qt_getEnumName_65(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_8ItemFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1701
// index:66
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Edges)

/*

 */
func Qt_getEnumName_66(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_4EdgeEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1719
// index:67
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TextInteractionFlags)

/*

 */
func Qt_getEnumName_67(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_19TextInteractionFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1715
// index:68
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ImageConversionFlags)

/*

 */
func Qt_getEnumName_68(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_19ImageConversionFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1745
// index:69
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ScreenOrientations)

/*

 */
func Qt_getEnumName_69(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_17ScreenOrientationEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1728
// index:70
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::KeyboardModifiers)

/*

 */
func Qt_getEnumName_70(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_16KeyboardModifierEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1742
// index:71
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::InputMethodQueries)

/*

 */
func Qt_getEnumName_71(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_16InputMethodQueryEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1743
// index:72
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::TouchPointStates)

/*

 */
func Qt_getEnumName_72(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_15TouchPointStateEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1740
// index:73
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::InputMethodHints)

/*

 */
func Qt_getEnumName_73(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_15InputMethodHintEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1702
// index:74
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::DockWidgetAreas)

/*

 */
func Qt_getEnumName_74(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_14DockWidgetAreaEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1697
// index:75
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Alignment)

/*

 */
func Qt_getEnumName_75(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_13AlignmentFlagEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1736
// index:76
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::WindowStates)

/*

 */
func Qt_getEnumName_76(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_11WindowStateEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1703
// index:77
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::ToolBarAreas)

/*

 */
func Qt_getEnumName_77(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_11ToolBarAreaEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1699
// index:78
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::Orientations)

/*

 */
func Qt_getEnumName_78(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_11OrientationEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1729
// index:79
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::MouseButtons)

/*

 */
func Qt_getEnumName_79(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_11MouseButtonEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1735
// index:80
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::WindowFlags)

/*

 */
func Qt_getEnumName_80(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_10WindowTypeEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qnamespace.h:1700
// index:81
// Invalid inline Visibility=Default Availability=Available
// [8] const char * qt_getEnumName(Qt::DropActions)

/*

 */
func Qt_getEnumName_81(arg0 int) string {
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt14qt_getEnumNameE6QFlagsINS_10DropActionEE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:92
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qstrcmp(const char *, const QByteArray &)

/*
A safe strcmp() function.

Compares str1 and str2. Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns an arbitrary non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrncmp(), qstricmp(), qstrnicmp(), and 8-bit Character Comparisons.
*/
func Qstrcmp(str1 string, str2 QByteArray_ITF) int {
	var convArg0 = qtrt.CString(str1)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if str2 != nil && str2.QByteArray_PTR() != nil {
		convArg1 = str2.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZL7qstrcmpPKcRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:843
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qIsNull(float)

/*

 */
func QIsNull(f float32) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZL7qIsNullf", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:827
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool qIsNull(double)

/*

 */
func QIsNull_1(d float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZL7qIsNulld", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:807
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(float, float)

/*

 */
func QFuzzyCompare(p1 float32, p2 float32) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZL13qFuzzyCompareff", qtrt.FFI_TYPE_POINTER, p1, p2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:802
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(double, double)

/*

 */
func QFuzzyCompare_1(p1 float64, p2 float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZL13qFuzzyComparedd", qtrt.FFI_TYPE_POINTER, p1, p2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:817
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyIsNull(float)

/*

 */
func QFuzzyIsNull(f float32) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZL12qFuzzyIsNullf", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:812
// index:1
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyIsNull(double)

/*

 */
func QFuzzyIsNull_1(d float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZL12qFuzzyIsNulld", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:1143
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qunsetenv(const char *)

/*

 */
func Qunsetenv(varName string) bool {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z9qunsetenvPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:732
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qt_assert(const char *, const char *, int)

/*

 */
func Qt_assert(assertion string, file string, line int) {
	var convArg0 = qtrt.CString(assertion)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(file)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z9qt_assertPKcS0_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, line)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:101
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qstrnicmp(const char *, const char *, uint)

/*
A safe strnicmp() function.

Compares at most len bytes of str1 and str2 ignoring the case of the characters. The encoding of the strings is assumed to be Latin-1.

Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns a random non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrcmp(), qstrncmp(), qstricmp(), and 8-bit Character Comparisons.
*/
func Qstrnicmp(arg0 string, arg1 string, len_ uint) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z9qstrnicmpPKcS0_j", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:53
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qIsFinite(float)

/*

 */
func QIsFinite(f float32) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z9qIsFinitef", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:50
// index:1
// Invalid Visibility=Default Availability=Available
// [1] bool qIsFinite(double)

/*

 */
func QIsFinite_1(d float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z9qIsFinited", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qhashfunctions.h:70
// index:0
// Invalid Visibility=Default Availability=Available
// [4] uint qHashBits(const void *, size_t, uint)

/*

 */
func QHashBits(p unsafe.Pointer /*666*/, size uint, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z9qHashBitsPKvmj", qtrt.FFI_TYPE_POINTER, p, size, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:685
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] QByteArray qCompress(const QByteArray &, int)

/*
This is an overloaded function.

Compresses the first nbytes of data at compression level compressionLevel and returns the compressed data in a new byte array.
*/
func QCompress(data QByteArray_ITF, compressionLevel int) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z9qCompressRK10QByteArrayi", qtrt.FFI_TYPE_POINTER, convArg0, compressionLevel)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:683
// index:1
// Invalid Visibility=Default Availability=Available
// [8] QByteArray qCompress(const uchar *, int, int)

/*
This is an overloaded function.

Compresses the first nbytes of data at compression level compressionLevel and returns the compressed data in a new byte array.
*/
func QCompress_1(data unsafe.Pointer /*666*/, nbytes int, compressionLevel int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z9qCompressPKhii", qtrt.FFI_TYPE_POINTER, data, nbytes, compressionLevel)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:109
// index:0
// Invalid Visibility=Default Availability=Available
// [2] quint16 qChecksum(const char *, uint, Qt::ChecksumType)

/*
Returns the CRC-16 checksum of the first len bytes of data.

The checksum is independent of the byte order (endianness) and will be calculated accorded to the algorithm published in ISO 3309 (Qt::ChecksumIso3309).

Note: This function is a 16-bit cache conserving (16 entry table) implementation of the CRC-16-CCITT algorithm.
*/
func QChecksum(s string, len_ uint, standard int) uint16 {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z9qChecksumPKcjN2Qt12ChecksumTypeE", qtrt.FFI_TYPE_POINTER, convArg0, len_, standard)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:108
// index:1
// Invalid Visibility=Default Availability=Available
// [2] quint16 qChecksum(const char *, uint)

/*
Returns the CRC-16 checksum of the first len bytes of data.

The checksum is independent of the byte order (endianness) and will be calculated accorded to the algorithm published in ISO 3309 (Qt::ChecksumIso3309).

Note: This function is a 16-bit cache conserving (16 entry table) implementation of the CRC-16-CCITT algorithm.
*/
func QChecksum_1(s string, len_ uint) uint16 {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z9qChecksumPKcj", qtrt.FFI_TYPE_POINTER, convArg0, len_)
	qtrt.ErrPrint(err, rv)
	return uint16(rv) // 222
}

// /usr/include/qt/QtCore/qglobal.h:781
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qBadAlloc()

/*

 */
func QBadAlloc() {
	rv, err := qtrt.InvokeQtFunc6("_Z9qBadAllocv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:76
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qstrnlen(const char *, uint)

/*
A safe strnlen() function.

Returns the number of characters that precede the terminating '\0', but at most maxlen. If str is 0, returns 0.

This function was introduced in  Qt 4.2.

See also qstrlen().
*/
func Qstrnlen(str string, maxlen uint) uint {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z8qstrnlenPKcj", qtrt.FFI_TYPE_POINTER, convArg0, maxlen)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:87
// index:0
// Invalid Visibility=Default Availability=Available
// [8] char * qstrncpy(char *, const char *, uint)

/*
A safe strncpy() function.

Copies at most len bytes from src (stopping at len or the terminating '\0' whichever comes first) into dst and returns a pointer to dst. Guarantees that dst is '\0'-terminated. If src or dst is 0, returns 0 immediately.

This function assumes that dst is at least len characters long.

Note: When compiling with Visual C++ compiler version 14.00 (Visual C++ 2005) or later, internally the function strncpy_s will be used.

See also qstrcpy().
*/
func Qstrncpy(dst string, src string, len_ uint) string {
	var convArg0 = qtrt.CString(dst)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(src)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z8qstrncpyPcPKcj", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:95
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qstrncmp(const char *, const char *, uint)

/*
A safe strncmp() function.

Compares at most len bytes of str1 and str2.

Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns a random non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrcmp(), qstricmp(), qstrnicmp(), and 8-bit Character Comparisons.
*/
func Qstrncmp(str1 string, str2 string, len_ uint) int {
	var convArg0 = qtrt.CString(str1)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(str2)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z8qstrncmpPKcS0_j", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, len_)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:100
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qstricmp(const char *, const char *)

/*
A safe stricmp() function.

Compares str1 and str2 ignoring the case of the characters. The encoding of the strings is assumed to be Latin-1.

Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns a random non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrcmp(), qstrncmp(), qstrnicmp(), and 8-bit Character Comparisons.
*/
func Qstricmp(arg0 string, arg1 string) int {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(arg1)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z8qstricmpPKcS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:544
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qint64 qRound64(float)

/*

 */
func QRound64(d float32) int64 {
	rv, err := qtrt.InvokeQtFunc6("_Z8qRound64f", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qglobal.h:542
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] qint64 qRound64(double)

/*

 */
func QRound64_1(d float64) int64 {
	rv, err := qtrt.InvokeQtFunc6("_Z8qRound64d", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return int64(rv) // 222
}

// /usr/include/qt/QtCore/qglobal.h:1150
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qIntCast(float)

/*

 */
func QIntCast(f float32) int {
	rv, err := qtrt.InvokeQtFunc6("_Z8qIntCastf", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:1149
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qIntCast(double)

/*

 */
func QIntCast_1(f float64) int {
	rv, err := qtrt.InvokeQtFunc6("_Z8qIntCastd", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmath.h:206
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qFastSin(qreal)

/*

 */
func QFastSin(x float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z8qFastSind", qtrt.FFI_TYPE_POINTER, x)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:216
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qFastCos(qreal)

/*

 */
func QFastCos(x float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z8qFastCosd", qtrt.FFI_TYPE_POINTER, x)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qcoreapplication.h:261
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qAppName()

/*

 */
func QAppName() string {
	rv, err := qtrt.InvokeQtFunc6("_Z8qAppNamev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qglobal.h:647
// index:0
// Invalid inline Visibility=Default Availability=Available
// [-2] void qt_noop()

/*

 */
func Qt_noop() {
	rv, err := qtrt.InvokeQtFunc6("_Z7qt_noopv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qhashfunctions.h:105
// index:0
// Invalid Visibility=Default Availability=Available
// [4] uint qt_hash(QStringView, uint)

/*

 */
func Qt_hash(key QStringView_ITF /*123*/, chained uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QStringView_PTR() != nil {
		convArg0 = key.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z7qt_hash11QStringViewj", qtrt.FFI_TYPE_POINTER, convArg0, chained)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:73
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qstrlen(const char *)

/*
A safe strlen() function.

Returns the number of characters that precede the terminating '\0', or 0 if str is 0.

See also qstrnlen().
*/
func Qstrlen(str string) uint {
	var convArg0 = qtrt.CString(str)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z7qstrlenPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qbytearray.h:71
// index:0
// Invalid Visibility=Default Availability=Available
// [8] char * qstrdup(const char *)

/*
Returns a duplicate string.

Allocates space for a copy of src, copies it, and returns a pointer to the copy. If src is 0, it immediately returns 0.

Ownership is passed to the caller, so the returned string must be deleted using delete[].
*/
func Qstrdup(arg0 string) string {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z7qstrdupPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:86
// index:0
// Invalid Visibility=Default Availability=Available
// [8] char * qstrcpy(char *, const char *)

/*
Copies all the characters up to and including the '\0' from src into dst and returns a pointer to dst. If src is 0, it immediately returns 0.

This function assumes that dst is large enough to hold the contents of src.

See also qstrncpy().
*/
func Qstrcpy(dst string, src string) string {
	var convArg0 = qtrt.CString(dst)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(src)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z7qstrcpyPcPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qbytearray.h:90
// index:1
// Invalid Visibility=Default Availability=Available
// [4] int qstrcmp(const QByteArray &, const QByteArray &)

/*
A safe strcmp() function.

Compares str1 and str2. Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns an arbitrary non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrncmp(), qstricmp(), qstrnicmp(), and 8-bit Character Comparisons.
*/
func Qstrcmp_1(str1 QByteArray_ITF, str2 QByteArray_ITF) int {
	var convArg0 unsafe.Pointer
	if str1 != nil && str1.QByteArray_PTR() != nil {
		convArg0 = str1.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if str2 != nil && str2.QByteArray_PTR() != nil {
		convArg1 = str2.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z7qstrcmpRK10QByteArrayS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:91
// index:2
// Invalid Visibility=Default Availability=Available
// [4] int qstrcmp(const QByteArray &, const char *)

/*
A safe strcmp() function.

Compares str1 and str2. Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns an arbitrary non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrncmp(), qstricmp(), qstrnicmp(), and 8-bit Character Comparisons.
*/
func Qstrcmp_2(str1 QByteArray_ITF, str2 string) int {
	var convArg0 unsafe.Pointer
	if str1 != nil && str1.QByteArray_PTR() != nil {
		convArg0 = str1.QByteArray_PTR().GetCthis()
	}
	var convArg1 = qtrt.CString(str2)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z7qstrcmpRK10QByteArrayPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qbytearray.h:89
// index:3
// Invalid Visibility=Default Availability=Available
// [4] int qstrcmp(const char *, const char *)

/*
A safe strcmp() function.

Compares str1 and str2. Returns a negative value if str1 is less than str2, 0 if str1 is equal to str2 or a positive value if str1 is greater than str2.

Special case 1: Returns 0 if str1 and str2 are both 0.

Special case 2: Returns an arbitrary non-zero value if str1 is 0 or str2 is 0 (but not both).

See also qstrncmp(), qstricmp(), qstrnicmp(), and 8-bit Character Comparisons.
*/
func Qstrcmp_3(str1 string, str2 string) int {
	var convArg0 = qtrt.CString(str1)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(str2)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_Z7qstrcmpPKcS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:1142
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qputenv(const char *, const QByteArray &)

/*

 */
func Qputenv(varName string, value QByteArray_ITF) bool {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	var convArg1 unsafe.Pointer
	if value != nil && value.QByteArray_PTR() != nil {
		convArg1 = value.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z7qputenvPKcRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:1134
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QByteArray qgetenv(const char *)

/*

 */
func Qgetenv(varName string) *QByteArray /*123*/ {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z7qgetenvPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qglobal.h:1051
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qtTrId(const char *, int)

/*

 */
func QtTrId(id string, n int) string {
	var convArg0 = qtrt.CString(id)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z6qtTrIdPKci", qtrt.FFI_TYPE_POINTER, convArg0, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qglobal.h:1155
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qsrand(uint)

/*

 */
func Qsrand(seed uint) {
	rv, err := qtrt.InvokeQtFunc6("_Z6qsrandj", qtrt.FFI_TYPE_POINTER, seed)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qglobal.h:539
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qRound(float)

/*

 */
func QRound(d float32) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qRoundf", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:537
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qRound(double)

/*

 */
func QRound_1(d float64) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qRoundd", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:52
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qIsNaN(float)

/*

 */
func QIsNaN(f float32) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z6qIsNaNf", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:49
// index:1
// Invalid Visibility=Default Availability=Available
// [1] bool qIsNaN(double)

/*

 */
func QIsNaN_1(d float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z6qIsNaNd", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:51
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qIsInf(float)

/*

 */
func QIsInf(f float32) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z6qIsInff", qtrt.FFI_TYPE_POINTER, f)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qnumeric.h:48
// index:1
// Invalid Visibility=Default Availability=Available
// [1] bool qIsInf(double)

/*

 */
func QIsInf_1(d float64) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z6qIsInfd", qtrt.FFI_TYPE_POINTER, d)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qmath.h:74
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qFloor(qreal)

/*

 */
func QFloor(v float64) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qFloord", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmath.h:122
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAtan2(qreal, qreal)

/*

 */
func QAtan2(y float64, x float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z6qAtan2dd", qtrt.FFI_TYPE_POINTER, y, x)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:1156
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qrand()

/*

 */
func Qrand() int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qrandv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmath.h:128
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qSqrt(qreal)

/*

 */
func QSqrt(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qSqrtd", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:54
// index:0
// Invalid Visibility=Default Availability=Available
// [8] double qSNaN()

/*

 */
func QSNaN() float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qSNaNv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:55
// index:0
// Invalid Visibility=Default Availability=Available
// [8] double qQNaN()

/*

 */
func QQNaN() float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qQNaNv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qhashfunctions.h:86
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(quint64, uint)

/*

 */
func QHash(key uint64, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashyj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:90
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(qint64, uint)

/*

 */
func QHash_1(key int64, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashxj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:75
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(ushort, uint)

/*

 */
func QHash_2(key uint16, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashtj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:76
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(short, uint)

/*

 */
func QHash_3(key int16, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashsj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:79
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(ulong, uint)

/*

 */
func QHash_4(key uint, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashmj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:85
// index:5
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(long, uint)

/*

 */
func QHash_5(key int, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashlj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:77
// index:6
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(uint, uint)

/*

 */
func QHash_6(key uint, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashjj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:78
// index:7
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(int, uint)

/*

 */
func QHash_7(key int, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashij", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:73
// index:8
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(uchar, uint)

/*

 */
func QHash_8(key byte, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashhj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:91
// index:9
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(float, uint)

/*

 */
func QHash_9(key float32, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashfj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:94
// index:10
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(long double, uint)

/*

 */
func QHash_10(key float64, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashej", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:92
// index:11
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(double, uint)

/*

 */
func QHash_11(key float64, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashdj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:72
// index:12
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(char, uint)

/*

 */
func QHash_12(key byte, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashcj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:74
// index:13
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(signed char, uint)

/*

 */
func QHash_13(key byte, seed uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashaj", qtrt.FFI_TYPE_POINTER, key, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qurlquery.h:53
// index:14
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QUrlQuery &, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash_14(key QUrlQuery_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QUrlQuery_PTR() != nil {
		convArg0 = key.QUrlQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QUrlQueryj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qmimetype.h:58
// index:15
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QMimeType &, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash_15(key QMimeType_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QMimeType_PTR() != nil {
		convArg0 = key.QMimeType_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QMimeTypej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:412
// index:16
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QDateTime &, uint)

/*

 */
func QHash_16(key QDateTime_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QDateTime_PTR() != nil {
		convArg0 = key.QDateTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QDateTimej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:103
// index:17
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QBitArray &, uint)

/*

 */
func QHash_17(key QBitArray_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QBitArray_PTR() != nil {
		convArg0 = key.QBitArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK9QBitArrayj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:99
// index:18
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QString &, uint)

/*

 */
func QHash_18(key string, seed uint) uint {
	var tmpArg0 = NewQString_5(key)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK7QStringj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qregexp.h:56
// index:19
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QRegExp &, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash_19(key QRegExp_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QRegExp_PTR() != nil {
		convArg0 = key.QRegExp_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK7QRegExpj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qlocale.h:62
// index:20
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QLocale &, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash_20(key QLocale_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLocale_PTR() != nil {
		convArg0 = key.QLocale_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK7QLocalej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/quuid.h:235
// index:21
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QUuid &, uint)

/*
Returns a hash of the UUID uuid, using seed to seed the calculation.

This function was introduced in  Qt 5.0.
*/
func QHash_21(uuid QUuid_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if uuid != nil && uuid.QUuid_PTR() != nil {
		convArg0 = uuid.QUuid_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK5QUuidj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:414
// index:22
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QTime &, uint)

/*

 */
func QHash_22(key QTime_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QTime_PTR() != nil {
		convArg0 = key.QTime_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK5QTimej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qdatetime.h:413
// index:23
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QDate &, uint)

/*

 */
func QHash_23(key QDate_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QDate_PTR() != nil {
		convArg0 = key.QDate_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK5QDatej", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qurl.h:122
// index:24
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QUrl &, uint)

/*

 */
func QHash_24(url QUrl_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK4QUrlj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:102
// index:25
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QPersistentModelIndex &, uint)

/*

 */
func QHash_25(index QPersistentModelIndex_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if index != nil && index.QPersistentModelIndex_PTR() != nil {
		convArg0 = index.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK21QPersistentModelIndexj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:149
// index:26
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QPersistentModelIndex &, uint)

/*

 */
func QHash_26(index QPersistentModelIndex_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if index != nil && index.QPersistentModelIndex_PTR() != nil {
		convArg0 = index.QPersistentModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK21QPersistentModelIndexj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qitemselectionmodel.h:227
// index:27
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QItemSelectionRange &)

/*

 */
func QHash_27(arg0 QItemSelectionRange_ITF) uint {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QItemSelectionRange_PTR() != nil {
		convArg0 = arg0.QItemSelectionRange_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK19QItemSelectionRange", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qregularexpression.h:62
// index:28
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QRegularExpression &, uint)

/*
Returns the hash value for key, using seed to seed the calculation.

This function was introduced in  Qt 5.6.
*/
func QHash_28(key QRegularExpression_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QRegularExpression_PTR() != nil {
		convArg0 = key.QRegularExpression_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK18QRegularExpressionj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qversionnumber.h:54
// index:29
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QVersionNumber &, uint)

/*

 */
func QHash_29(key QVersionNumber_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QVersionNumber_PTR() != nil {
		convArg0 = key.QVersionNumber_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK14QVersionNumberj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qabstractitemmodel.h:437
// index:30
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QModelIndex &)

/*

 */
func QHash_30(index QModelIndex_ITF) uint {
	var convArg0 unsafe.Pointer
	if index != nil && index.QModelIndex_PTR() != nil {
		convArg0 = index.QModelIndex_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK11QModelIndex", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:100
// index:31
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QStringRef &, uint)

/*

 */
func QHash_31(key QStringRef_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QStringRef_PTR() != nil {
		convArg0 = key.QStringRef_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK10QStringRefj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:97
// index:32
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(const QByteArray &, uint)

/*

 */
func QHash_32(key QByteArray_ITF, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QByteArray_PTR() != nil {
		convArg0 = key.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHashRK10QByteArrayj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:96
// index:33
// Invalid inline Visibility=Default Availability=Available
// [4] uint qHash(const QChar, uint)

/*

 */
func QHash_33(key QChar_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QChar_PTR() != nil {
		convArg0 = key.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash5QCharj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:104
// index:34
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(QLatin1String, uint)

/*

 */
func QHash_34(key QLatin1String_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QLatin1String_PTR() != nil {
		convArg0 = key.QLatin1String_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash13QLatin1Stringj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:102
// index:35
// Invalid Visibility=Default Availability=Available
// [4] uint qHash(QStringView, uint)

/*

 */
func QHash_35(key QStringView_ITF /*123*/, seed uint) uint {
	var convArg0 unsafe.Pointer
	if key != nil && key.QStringView_PTR() != nil {
		convArg0 = key.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z5qHash11QStringViewj", qtrt.FFI_TYPE_POINTER, convArg0, seed)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:80
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qFabs(qreal)

/*

 */
func QFabs(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qFabsd", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:68
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qCeil(qreal)

/*

 */
func QCeil(v float64) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qCeild", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qmath.h:116
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAtan(qreal)

/*

 */
func QAtan(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qAtand", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:110
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAsin(qreal)

/*

 */
func QAsin(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qAsind", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:104
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qAcos(qreal)

/*

 */
func QAcos(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z5qAcosd", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:98
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qTan(qreal)

/*

 */
func QTan(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z4qTand", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:86
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qSin(qreal)

/*

 */
func QSin(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z4qSind", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:146
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qPow(qreal, qreal)

/*

 */
func QPow(x float64, y float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z4qPowdd", qtrt.FFI_TYPE_POINTER, x, y)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qnumeric.h:56
// index:0
// Invalid Visibility=Default Availability=Available
// [8] double qInf()

/*

 */
func QInf() float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z4qInfv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:140
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qExp(qreal)

/*

 */
func QExp(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z4qExpd", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:92
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qCos(qreal)

/*

 */
func QCos(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z4qCosd", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qmath.h:134
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] qreal qLn(qreal)

/*

 */
func QLn(v float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z3qLnd", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qplugin.h:78
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRegisterStaticPluginFunction(QStaticPlugin)

/*

 */
func QRegisterStaticPluginFunction(staticPlugin QStaticPlugin_ITF /*123*/) {
	var convArg0 unsafe.Pointer
	if staticPlugin != nil && staticPlugin.QStaticPlugin_PTR() != nil {
		convArg0 = staticPlugin.QStaticPlugin_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z29qRegisterStaticPluginFunction13QStaticPlugin", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qglobal.h:1147
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qEnvironmentVariableIntValue(const char *, bool *)

/*

 */
func QEnvironmentVariableIntValue(varName string, ok *bool) int {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z28qEnvironmentVariableIntValuePKcPb", qtrt.FFI_TYPE_POINTER, convArg0, ok)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:1145
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qEnvironmentVariableIsEmpty(const char *)

/*

 */
func QEnvironmentVariableIsEmpty(varName string) bool {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z27qEnvironmentVariableIsEmptyPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:1146
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qEnvironmentVariableIsSet(const char *)

/*

 */
func QEnvironmentVariableIsSet(varName string) bool {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z25qEnvironmentVariableIsSetPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// ./bsheaders/QtCore/extra_export.h:7
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qUnregisterResourceData(int, const unsigned char *, const unsigned char *, const unsigned char *)

/*

 */
func QUnregisterResourceData(arg0 int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/, arg3 unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z23qUnregisterResourceDataiPKhS0_S0_", qtrt.FFI_TYPE_POINTER, arg0, arg1, arg2, arg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextstream.h:279
// index:0
// Invalid inline Visibility=Default Availability=Available
// [40] QTextStreamManipulator qSetRealNumberPrecision(int)

/*
Equivalent to QTextStream::setRealNumberPrecision(precision).
*/
func QSetRealNumberPrecision(precision int) *QTextStreamManipulator /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z23qSetRealNumberPrecisioni", qtrt.FFI_TYPE_POINTER, precision)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamManipulatorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStreamManipulator)
	return rv2
}

// /usr/include/qt/QtCore/qlogging.h:191
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QtMessageHandler qInstallMessageHandler(QtMessageHandler)

/*

 */
func QInstallMessageHandler(arg0 unsafe.Pointer /*666*/) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z22qInstallMessageHandlerPFv9QtMsgTypeRK18QMessageLogContextRK7QStringE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qalgorithms.h:791
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint64)

/*

 */
func QCountTrailingZeroBits(v uint64) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsy", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:775
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint16)

/*

 */
func QCountTrailingZeroBits_1(v uint16) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z22qCountTrailingZeroBitst", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:802
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(unsigned long)

/*

 */
func QCountTrailingZeroBits_2(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsm", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:742
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint32)

/*

 */
func QCountTrailingZeroBits_3(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsj", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:760
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountTrailingZeroBits(quint8)

/*

 */
func QCountTrailingZeroBits_4(v byte) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z22qCountTrailingZeroBitsh", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// ./bsheaders/QtCore/extra_export.h:6
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qRegisterResourceData(int, const unsigned char *, const unsigned char *, const unsigned char *)

/*

 */
func QRegisterResourceData(arg0 int, arg1 unsafe.Pointer /*666*/, arg2 unsafe.Pointer /*666*/, arg3 unsafe.Pointer /*666*/) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z21qRegisterResourceDataiPKhS0_S0_", qtrt.FFI_TYPE_POINTER, arg0, arg1, arg2, arg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qalgorithms.h:847
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint64)

/*

 */
func QCountLeadingZeroBits(v uint64) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsy", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:834
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint16)

/*

 */
func QCountLeadingZeroBits_1(v uint16) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z21qCountLeadingZeroBitst", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:862
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(unsigned long)

/*

 */
func QCountLeadingZeroBits_2(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsm", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:807
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint32)

/*

 */
func QCountLeadingZeroBits_3(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsj", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:822
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qCountLeadingZeroBits(quint8)

/*

 */
func QCountLeadingZeroBits_4(v byte) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z21qCountLeadingZeroBitsh", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qobject.h:93
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QObject * qt_qFindChild_helper(const QObject *, const QString &, const QMetaObject &, Qt::FindChildOptions)

/*

 */
func Qt_qFindChild_helper(parent QObject_ITF /*777 const QObject **/, name string, mo QMetaObject_ITF, options int) *QObject /*777 QObject **/ {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if mo != nil && mo.QMetaObject_PTR() != nil {
		convArg2 = mo.QMetaObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z20qt_qFindChild_helperPK7QObjectRK7QStringRK11QMetaObject6QFlagsIN2Qt15FindChildOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, options)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qglobal.h:1140
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qEnvironmentVariable(const char *, const QString &)

/*

 */
func QEnvironmentVariable(varName string, defaultValue string) string {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	var tmpArg1 = NewQString_5(defaultValue)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z20qEnvironmentVariablePKcRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qglobal.h:1139
// index:1
// Invalid Visibility=Default Availability=Available
// [8] QString qEnvironmentVariable(const char *)

/*

 */
func QEnvironmentVariable_1(varName string) string {
	var convArg0 = qtrt.CString(varName)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z20qEnvironmentVariablePKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qhashfunctions.h:68
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qSetGlobalQHashSeed(int)

/*

 */
func QSetGlobalQHashSeed(newSeed int) {
	rv, err := qtrt.InvokeQtFunc6("_Z19qSetGlobalQHashSeedi", qtrt.FFI_TYPE_POINTER, newSeed)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qnamespace.h:53
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getQtMetaObject()

/*

 */
func Qt_getQtMetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_Z18qt_getQtMetaObjectv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qobject.h:473
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] const QMetaObject * qt_getQtMetaObject()

/*

 */
func Qt_getQtMetaObject_1() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_Z18qt_getQtMetaObjectv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qlogging.h:193
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qSetMessagePattern(const QString &)

/*

 */
func QSetMessagePattern(messagePattern string) {
	var tmpArg0 = NewQString_5(messagePattern)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z18qSetMessagePatternRK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qcoreapplication.h:260
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qRemovePostRoutine(QtCleanUpFunction)

/*
Removes the cleanup routine specified by ptr from the list of routines called by the QCoreApplication destructor. The routine must have been previously added to the list by a call to qAddPostRoutine(), otherwise this function has no effect.

Note: This function has been thread-safe since Qt 5.10.

Note: This function is thread-safe.

This function was introduced in  Qt 5.3.

See also qAddPostRoutine().
*/
func QRemovePostRoutine(arg0 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_Z18qRemovePostRoutinePFvvE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlogging.h:179
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qt_message_output(QtMsgType, const QMessageLogContext &, const QString &)

/*

 */
func Qt_message_output(arg0 int, context QMessageLogContext_ITF, message string) {
	var convArg1 unsafe.Pointer
	if context != nil && context.QMessageLogContext_PTR() != nil {
		convArg1 = context.QMessageLogContext_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(message)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z17qt_message_output9QtMsgTypeRK18QMessageLogContextRK7QString", qtrt.FFI_TYPE_POINTER, arg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qmath.h:236
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] float qRadiansToDegrees(float)

/*

 */
func QRadiansToDegrees(radians float32) float32 {
	rv, err := qtrt.InvokeQtFunc6("_Z17qRadiansToDegreesf", qtrt.FFI_TYPE_POINTER, radians)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qmath.h:241
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] double qRadiansToDegrees(double)

/*

 */
func QRadiansToDegrees_1(radians float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z17qRadiansToDegreesd", qtrt.FFI_TYPE_POINTER, radians)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qlogging.h:194
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qFormatLogMessage(QtMsgType, const QMessageLogContext &, const QString &)

/*

 */
func QFormatLogMessage(type_ int, context QMessageLogContext_ITF, buf string) string {
	var convArg1 unsafe.Pointer
	if context != nil && context.QMessageLogContext_PTR() != nil {
		convArg1 = context.QMessageLogContext_PTR().GetCthis()
	}
	var tmpArg2 = NewQString_5(buf)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z17qFormatLogMessage9QtMsgTypeRK18QMessageLogContextRK7QString", qtrt.FFI_TYPE_POINTER, type_, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qmath.h:226
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] float qDegreesToRadians(float)

/*

 */
func QDegreesToRadians(degrees float32) float32 {
	rv, err := qtrt.InvokeQtFunc6("_Z17qDegreesToRadiansf", qtrt.FFI_TYPE_POINTER, degrees)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float32", rv).(float32) // 1111
}

// /usr/include/qt/QtCore/qmath.h:231
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] double qDegreesToRadians(double)

/*

 */
func QDegreesToRadians_1(degrees float64) float64 {
	rv, err := qtrt.InvokeQtFunc6("_Z17qDegreesToRadiansd", qtrt.FFI_TYPE_POINTER, degrees)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:780
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qt_check_pointer(const char *, int)

/*

 */
func Qt_check_pointer(arg0 string, arg1 int) {
	var convArg0 = qtrt.CString(arg0)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z16qt_check_pointerPKci", qtrt.FFI_TYPE_POINTER, convArg0, arg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qalgorithms.h:717
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint64)

/*

 */
func QPopulationCount(v uint64) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z16qPopulationCounty", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:706
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint16)

/*

 */
func QPopulationCount_1(v uint16) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z16qPopulationCountt", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:732
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(unsigned long)

/*

 */
func QPopulationCount_2(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z16qPopulationCountm", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:683
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint32)

/*

 */
func QPopulationCount_3(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z16qPopulationCountj", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qalgorithms.h:696
// index:4
// Invalid inline Visibility=Default Availability=Available
// [4] uint qPopulationCount(quint8)

/*

 */
func QPopulationCount_4(v byte) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z16qPopulationCounth", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qhashfunctions.h:67
// index:0
// Invalid Visibility=Default Availability=Available
// [4] int qGlobalQHashSeed()

/*

 */
func QGlobalQHashSeed() int {
	rv, err := qtrt.InvokeQtFunc6("_Z16qGlobalQHashSeedv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qglobal.h:727
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qt_error_string(int)

/*

 */
func Qt_error_string(errorCode int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z15qt_error_stringi", qtrt.FFI_TYPE_POINTER, errorCode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qglobal.h:901
// index:0
// Invalid Visibility=Default Availability=Available
// [8] void * qReallocAligned(void *, size_t, size_t, size_t)

/*

 */
func QReallocAligned(ptr unsafe.Pointer /*666*/, size uint, oldsize uint, alignment uint) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z15qReallocAlignedPvmmm", qtrt.FFI_TYPE_POINTER, ptr, size, oldsize, alignment)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qmath.h:264
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] quint64 qNextPowerOfTwo(quint64)

/*

 */
func QNextPowerOfTwo(v uint64) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_Z15qNextPowerOfTwoy", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:287
// index:1
// Invalid inline Visibility=Default Availability=Available
// [8] quint64 qNextPowerOfTwo(qint64)

/*

 */
func QNextPowerOfTwo_1(v int64) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_Z15qNextPowerOfTwox", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:247
// index:2
// Invalid inline Visibility=Default Availability=Available
// [4] quint32 qNextPowerOfTwo(quint32)

/*

 */
func QNextPowerOfTwo_2(v uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z15qNextPowerOfTwoj", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qmath.h:282
// index:3
// Invalid inline Visibility=Default Availability=Available
// [4] quint32 qNextPowerOfTwo(qint32)

/*

 */
func QNextPowerOfTwo_3(v int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z15qNextPowerOfTwoi", qtrt.FFI_TYPE_POINTER, v)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qcoreapplication.h:259
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qAddPostRoutine(QtCleanUpFunction)

/*
Adds a global routine that will be called from the QCoreApplication destructor. This function is normally used to add cleanup routines for program-wide functionality.

The cleanup routines are called in the reverse order of their addition.

The function specified by ptr should take no arguments and should return nothing. For example:


  static int *global_ptr = 0;

  static void cleanup_ptr()
  {
      delete [] global_ptr;
      global_ptr = 0;
  }

  void init_ptr()
  {
      global_ptr = new int[100];      // allocate data
      qAddPostRoutine(cleanup_ptr);   // delete later
  }



Note that for an application- or module-wide cleanup, qAddPostRoutine() is often not suitable. For example, if the program is split into dynamically loaded modules, the relevant module may be unloaded long before the QCoreApplication destructor is called. In such cases, if using qAddPostRoutine() is still desirable, qRemovePostRoutine() can be used to prevent a routine from being called by the QCoreApplication destructor. For example, if that routine was called before the module was unloaded.

For modules and libraries, using a reference-counted initialization manager or Qt's parent-child deletion mechanism may be better. Here is an example of a private class that uses the parent-child mechanism to call a cleanup function at the right time:


  class MyPrivateInitStuff : public QObject
  {
  public:
      static MyPrivateInitStuff *initStuff(QObject *parent)
      {
          if (!p)
              p = new MyPrivateInitStuff(parent);
          return p;
      }

      ~MyPrivateInitStuff()
      {
          // cleanup goes here
      }

  private:
      MyPrivateInitStuff(QObject *parent)
          : QObject(parent)
      {
          // initialization goes here
      }

      MyPrivateInitStuff *p;
  };



By selecting the right parent object, this can often be made to clean up the module's data at the right moment.

Note: This function has been thread-safe since Qt 5.10.

Note: This function is thread-safe.

See also qRemovePostRoutine().
*/
func QAddPostRoutine(arg0 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_Z15qAddPostRoutinePFvvE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qtextstream.h:267
// index:0
// Invalid inline Visibility=Default Availability=Available
// [40] QTextStreamManipulator qSetFieldWidth(int)

/*
Equivalent to QTextStream::setFieldWidth(width).
*/
func QSetFieldWidth(width int) *QTextStreamManipulator /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z14qSetFieldWidthi", qtrt.FFI_TYPE_POINTER, width)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamManipulatorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStreamManipulator)
	return rv2
}

// /usr/include/qt/QtCore/qglobal.h:900
// index:0
// Invalid Visibility=Default Availability=Available
// [8] void * qMallocAligned(size_t, size_t)

/*

 */
func QMallocAligned(size uint, alignment uint) unsafe.Pointer /*666*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z14qMallocAlignedmm", qtrt.FFI_TYPE_POINTER, size, alignment)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qnumeric.h:58
// index:0
// Invalid Visibility=Default Availability=Available
// [4] quint32 qFloatDistance(float, float)

/*

 */
func QFloatDistance(a float32, b float32) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z14qFloatDistanceff", qtrt.FFI_TYPE_POINTER, a, b)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtCore/qnumeric.h:59
// index:1
// Invalid Visibility=Default Availability=Available
// [8] quint64 qFloatDistance(double, double)

/*

 */
func QFloatDistance_1(a float64, b float64) uint64 {
	rv, err := qtrt.InvokeQtFunc6("_Z14qFloatDistancedd", qtrt.FFI_TYPE_POINTER, a, b)
	qtrt.ErrPrint(err, rv)
	return uint64(rv) // 222
}

// /usr/include/qt/QtCore/qcoreapplication.h:258
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qAddPreRoutine(QtStartUpFunction)

/*

 */
func QAddPreRoutine(arg0 unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_Z14qAddPreRoutinePFvvE", qtrt.FFI_TYPE_POINTER, arg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qobjectdefs.h:261
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qFlagLocation(const char *)

/*

 */
func QFlagLocation(method string) string {
	var convArg0 = qtrt.CString(method)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_Z13qFlagLocationPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtCore/qglobal.h:687
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool qSharedBuild()

/*

 */
func QSharedBuild() bool {
	rv, err := qtrt.InvokeQtFunc6("_Z12qSharedBuildv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qglobal.h:902
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qFreeAligned(void *)

/*

 */
func QFreeAligned(ptr unsafe.Pointer /*666*/) {
	rv, err := qtrt.InvokeQtFunc6("_Z12qFreeAlignedPv", qtrt.FFI_TYPE_POINTER, ptr)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qglobal.h:749
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qt_assert_x(const char *, const char *, const char *, int)

/*

 */
func Qt_assert_x(where string, what string, file string, line int) {
	var convArg0 = qtrt.CString(where)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(what)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(file)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_Z11qt_assert_xPKcS0_S0_i", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2, line)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qbytearray.h:687
// index:0
// Invalid inline Visibility=Default Availability=Available
// [8] QByteArray qUncompress(const QByteArray &)

/*
This is an overloaded function.

Uncompresses the first nbytes of data and returns a new byte array with the uncompressed data.
*/
func QUncompress(data QByteArray_ITF) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if data != nil && data.QByteArray_PTR() != nil {
		convArg0 = data.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z11qUncompressRK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qbytearray.h:684
// index:1
// Invalid Visibility=Default Availability=Available
// [8] QByteArray qUncompress(const uchar *, int)

/*
This is an overloaded function.

Uncompresses the first nbytes of data and returns a new byte array with the uncompressed data.
*/
func QUncompress_1(data unsafe.Pointer /*666*/, nbytes int) *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_Z11qUncompressPKhi", qtrt.FFI_TYPE_POINTER, data, nbytes)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextstream.h:273
// index:0
// Invalid inline Visibility=Default Availability=Available
// [40] QTextStreamManipulator qSetPadChar(QChar)

/*
Equivalent to QTextStream::setPadChar(ch).
*/
func QSetPadChar(ch QChar_ITF /*123*/) *QTextStreamManipulator /*123*/ {
	var convArg0 unsafe.Pointer
	if ch != nil && ch.QChar_PTR() != nil {
		convArg0 = ch.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z11qSetPadChar5QChar", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextStreamManipulatorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextStreamManipulator)
	return rv2
}

//  body block end
