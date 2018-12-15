package qtgui

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

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
	if false {
		qtcore.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtGui/qquaternion.h:293
// index:48
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator+(const QQuaternion &, const QQuaternion &)

/*

 */
func Operator_add48(q1 QQuaternion_ITF, q2 QQuaternion_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZplRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:336
// index:53
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QTextItem::RenderFlags::enum_type, int)

/*

 */
func Operator_or53(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN9QTextItem10RenderFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qpainter.h:497
// index:54
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QPainter::RenderHints::enum_type, int)

/*

 */
func Operator_or54(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN8QPainter10RenderHintEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopengldebug.h:147
// index:55
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLDebugMessage::Severities::enum_type, int)

/*

 */
func Operator_or55(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN19QOpenGLDebugMessage8SeverityEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopengldebug.h:145
// index:56
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLDebugMessage::Sources::enum_type, int)

/*

 */
func Operator_or56(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN19QOpenGLDebugMessage6SourceEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopengldebug.h:146
// index:57
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLDebugMessage::Types::enum_type, int)

/*

 */
func Operator_or57(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN19QOpenGLDebugMessage4TypeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopenglfunctions.h:445
// index:58
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLFunctions::OpenGLFeatures::enum_type, int)

/*

 */
func Operator_or58(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN16QOpenGLFunctions13OpenGLFeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qsurfaceformat.h:179
// index:59
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSurfaceFormat::FormatOptions::enum_type, int)

/*

 */
func Operator_or59(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN14QSurfaceFormat12FormatOptionEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopengltexture.h:627
// index:60
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLTexture::Features::enum_type, int)

/*

 */
func Operator_or60(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN14QOpenGLTexture7FeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qimageiohandler.h:157
// index:61
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QImageIOPlugin::Capabilities::enum_type, int)

/*

 */
func Operator_or61(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN14QImageIOPlugin10CapabilityEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:304
// index:62
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QTextDocument::FindFlags::enum_type, int)

/*

 */
func Operator_or62(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QTextDocument8FindFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopenglshaderprogram.h:108
// index:63
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLShader::ShaderType::enum_type, int)

/*

 */
func Operator_or63(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QOpenGLShader13ShaderTypeBitEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qopenglbuffer.h:139
// index:64
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QOpenGLBuffer::RangeAccessFlags::enum_type, int)

/*

 */
func Operator_or64(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QOpenGLBuffer15RangeAccessFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qtouchdevice.h:93
// index:65
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QTouchDevice::Capabilities::enum_type, int)

/*

 */
func Operator_or65(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN12QTouchDevice14CapabilityFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:338
// index:66
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QPaintEngine::DirtyFlags::enum_type, int)

/*

 */
func Operator_or66(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN12QPaintEngine9DirtyFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qpaintengine.h:337
// index:67
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QPaintEngine::PaintEngineFeatures::enum_type, int)

/*

 */
func Operator_or67(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN12QPaintEngine18PaintEngineFeatureEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qevent.h:986
// index:68
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QTouchEvent::TouchPoint::InfoFlags::enum_type, int)

/*

 */
func Operator_or68(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QTouchEvent10TouchPoint8InfoFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qtextoption.h:149
// index:69
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QTextOption::Flags::enum_type, int)

/*

 */
func Operator_or69(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QTextOption4FlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qtextformat.h:389
// index:70
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QTextFormat::PageBreakFlags::enum_type, int)

/*

 */
func Operator_or70(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QTextFormat13PageBreakFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:445
// index:71
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QAccessible::Relation::enum_type, int)

/*

 */
func Operator_or71(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QAccessible12RelationFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:313
// index:24
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator-(const QQuaternion &)

/*
Returns a QQuaternion object that is formed by subtracting q2 from q1; each component is subtracted separately.

See also QQuaternion::operator-=().
*/
func Operator_minus24(quaternion QQuaternion_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZngRK11QQuaternion", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qpagesize.h:301
// index:55
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QPageSize &, const QPageSize &)

/*

 */
func Operator_not_equal55(lhs QPageSize_ITF, rhs QPageSize_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QPageSize_PTR() != nil {
		convArg0 = lhs.QPageSize_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QPageSize_PTR() != nil {
		convArg1 = rhs.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK9QPageSizeS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qcursor.h:121
// index:56
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QCursor &, const QCursor &)

/*

 */
func Operator_not_equal56(lhs QCursor_ITF, rhs QCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QCursor_PTR() != nil {
		convArg0 = lhs.QCursor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QCursor_PTR() != nil {
		convArg1 = rhs.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK7QCursorS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:173
// index:57
// Invalid Visibility=Default Availability=Available
// [1] bool operator!=(const QSurfaceFormat &, const QSurfaceFormat &)

/*

 */
func Operator_not_equal57(arg0 QSurfaceFormat_ITF, arg1 QSurfaceFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSurfaceFormat_PTR() != nil {
		convArg0 = arg0.QSurfaceFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QSurfaceFormat_PTR() != nil {
		convArg1 = arg1.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK14QSurfaceFormatS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:288
// index:58
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QQuaternion &, const QQuaternion &)

/*

 */
func Operator_not_equal58(q1 QQuaternion_ITF, q2 QQuaternion_ITF) bool {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:144
// index:59
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator!=(const QPageLayout &, const QPageLayout &)

/*

 */
func Operator_not_equal59(lhs QPageLayout_ITF, rhs QPageLayout_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QPageLayout_PTR() != nil {
		convArg0 = lhs.QPageLayout_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QPageLayout_PTR() != nil {
		convArg1 = rhs.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZneRK11QPageLayoutS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:303
// index:27
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator*(float, const QQuaternion &)

/*

 */
func Operator_mul27(factor float32, quaternion QQuaternion_ITF) *QQuaternion /*123*/ {
	var convArg1 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg1 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZmlfRK11QQuaternion", qtrt.FFI_TYPE_POINTER, factor, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:308
// index:28
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator*(const QQuaternion &, float)

/*

 */
func Operator_mul28(quaternion QQuaternion_ITF, factor float32) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZmlRK11QQuaternionf", qtrt.FFI_TYPE_POINTER, convArg0, factor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:257
// index:29
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator*(const QQuaternion &, const QQuaternion &)

/*

 */
func Operator_mul29(q1 QQuaternion_ITF, q2 QQuaternion_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZmlRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qquaternion.h:298
// index:25
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator-(const QQuaternion &, const QQuaternion &)

/*
Returns a QQuaternion object that is formed by subtracting q2 from q1; each component is subtracted separately.

See also QQuaternion::operator-=().
*/
func Operator_minus25(q1 QQuaternion_ITF, q2 QQuaternion_ITF) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZmiRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qaccessible.h:443
// index:56
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QAccessible::State &, const QAccessible::State &)

/*

 */
func Operator_equal_equal56(first int, second int) bool {
	rv, err := qtrt.InvokeQtFunc6("_ZeqRKN11QAccessible5StateES2_", qtrt.FFI_TYPE_POINTER, &first, &second)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagesize.h:300
// index:57
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QPageSize &, const QPageSize &)

/*

 */
func Operator_equal_equal57(lhs QPageSize_ITF, rhs QPageSize_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QPageSize_PTR() != nil {
		convArg0 = lhs.QPageSize_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QPageSize_PTR() != nil {
		convArg1 = rhs.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK9QPageSizeS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qcursor.h:120
// index:58
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QCursor &, const QCursor &)

/*

 */
func Operator_equal_equal58(lhs QCursor_ITF, rhs QCursor_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QCursor_PTR() != nil {
		convArg0 = lhs.QCursor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QCursor_PTR() != nil {
		convArg1 = rhs.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK7QCursorS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qsurfaceformat.h:172
// index:59
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QSurfaceFormat &, const QSurfaceFormat &)

/*

 */
func Operator_equal_equal59(arg0 QSurfaceFormat_ITF, arg1 QSurfaceFormat_ITF) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QSurfaceFormat_PTR() != nil {
		convArg0 = arg0.QSurfaceFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if arg1 != nil && arg1.QSurfaceFormat_PTR() != nil {
		convArg1 = arg1.QSurfaceFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK14QSurfaceFormatS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:184
// index:60
// Invalid inline Visibility=Default Availability=Available
// [1] bool operator==(const QQuaternion &, const QQuaternion &)

/*

 */
func Operator_equal_equal60(q1 QQuaternion_ITF, q2 QQuaternion_ITF) bool {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qpagelayout.h:143
// index:61
// Invalid Visibility=Default Availability=Available
// [1] bool operator==(const QPageLayout &, const QPageLayout &)

/*

 */
func Operator_equal_equal61(lhs QPageLayout_ITF, rhs QPageLayout_ITF) bool {
	var convArg0 unsafe.Pointer
	if lhs != nil && lhs.QPageLayout_PTR() != nil {
		convArg0 = lhs.QPageLayout_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if rhs != nil && rhs.QPageLayout_PTR() != nil {
		convArg1 = rhs.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZeqRK11QPageLayoutS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qquaternion.h:318
// index:16
// Invalid inline Visibility=Default Availability=Available
// [16] const QQuaternion operator/(const QQuaternion &, float)

/*

 */
func Operator_div16(quaternion QQuaternion_ITF, divisor float32) *QQuaternion /*123*/ {
	var convArg0 unsafe.Pointer
	if quaternion != nil && quaternion.QQuaternion_PTR() != nil {
		convArg0 = quaternion.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZdvRK11QQuaternionf", qtrt.FFI_TYPE_POINTER, convArg0, divisor)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQQuaternionFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQQuaternion)
	return rv2
}

// /usr/include/qt/QtGui/qtextdocument.h:76
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString convertFromPlainText(const QString &, Qt::WhiteSpaceMode)

/*

 */
func ConvertFromPlainText(plain string, mode int) string {
	var tmpArg0 = qtcore.NewQString5(plain)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt20convertFromPlainTextERK7QStringNS_14WhiteSpaceModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qtextdocument.h:75
// index:0
// Invalid Visibility=Default Availability=Available
// [1] bool mightBeRichText(const QString &)

/*

 */
func MightBeRichText(arg0 string) bool {
	var tmpArg0 = qtcore.NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt15mightBeRichTextERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qtextdocument.h:79
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QTextCodec * codecForHtml(const QByteArray &)

/*

 */
func CodecForHtml(ba qtcore.QByteArray_ITF) *qtcore.QTextCodec /*777 QTextCodec **/ {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN2Qt12codecForHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qrgb.h:78
// index:0
// Invalid inline Visibility=Default Availability=Available
// [1] bool qIsGray(QRgb)

/*

 */
func QIsGray(rgb uint) bool {
	rv, err := qtrt.InvokeQtFunc6("_Z7qIsGrayj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrgb.h:57
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qGreen(QRgb)

/*

 */
func QGreen(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qGreenj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:63
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qAlpha(QRgb)

/*

 */
func QAlpha(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z6qAlphaj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:69
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qRgba(int, int, int, int)

/*

 */
func QRgba(r int, g int, b int, a int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z5qRgbaiiii", qtrt.FFI_TYPE_POINTER, r, g, b, a)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgb.h:75
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qGray(QRgb)

/*

 */
func QGray(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qGrayj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:72
// index:1
// Invalid inline Visibility=Default Availability=Available
// [4] int qGray(int, int, int)

/*

 */
func QGray1(r int, g int, b int) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qGrayiii", qtrt.FFI_TYPE_POINTER, r, g, b)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qrgb.h:60
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qBlue(QRgb)

/*

 */
func QBlue(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z5qBluej", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qpixmapcache.h:95
// index:32
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPixmapCache::Key &, QPixmapCache::Key &)

/*

 */
func Swap32(value1 int, value2 int) {
	rv, err := qtrt.InvokeQtFunc6("_Z4swapRN12QPixmapCache3KeyES1_", qtrt.FFI_TYPE_POINTER, &value1, &value2)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:183
// index:33
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPolygonF &, QPolygonF &)

/*
Swaps polygon other with this polygon. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap33(value1 QPolygonF_ITF, value2 QPolygonF_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPolygonF_PTR() != nil {
		convArg0 = value1.QPolygonF_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPolygonF_PTR() != nil {
		convArg1 = value2.QPolygonF_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR9QPolygonFS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagesize.h:298
// index:34
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPageSize &, QPageSize &)

/*
Swaps this QPageSize with other. This function is very fast and never fails.
*/
func Swap34(value1 QPageSize_ITF, value2 QPageSize_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPageSize_PTR() != nil {
		convArg0 = value1.QPageSize_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPageSize_PTR() != nil {
		convArg1 = value2.QPageSize_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR9QPageSizeS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qglyphrun.h:128
// index:35
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QGlyphRun &, QGlyphRun &)

/*
Swaps this glyph run instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap35(value1 QGlyphRun_ITF, value2 QGlyphRun_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QGlyphRun_PTR() != nil {
		convArg0 = value1.QGlyphRun_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QGlyphRun_PTR() != nil {
		convArg1 = value2.QGlyphRun_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR9QGlyphRunS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpolygon.h:104
// index:36
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPolygon &, QPolygon &)

/*
Swaps polygon other with this polygon. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap36(value1 QPolygon_ITF, value2 QPolygon_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPolygon_PTR() != nil {
		convArg0 = value1.QPolygon_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPolygon_PTR() != nil {
		convArg1 = value2.QPolygon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR8QPolygonS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpicture.h:121
// index:37
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPicture &, QPicture &)

/*
Swaps picture other with this picture. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap37(value1 QPicture_ITF, value2 QPicture_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPicture_PTR() != nil {
		convArg0 = value1.QPicture_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPicture_PTR() != nil {
		convArg1 = value2.QPicture_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR8QPictureS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpalette.h:193
// index:38
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPalette &, QPalette &)

/*
Swaps this palette instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap38(value1 QPalette_ITF, value2 QPalette_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPalette_PTR() != nil {
		convArg0 = value1.QPalette_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPalette_PTR() != nil {
		convArg1 = value2.QPalette_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR8QPaletteS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qregion.h:184
// index:39
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QRegion &, QRegion &)

/*
Swaps region other with this region. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap39(value1 QRegion_ITF, value2 QRegion_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QRegion_PTR() != nil {
		convArg0 = value1.QRegion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QRegion_PTR() != nil {
		convArg1 = value2.QRegion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR7QRegionS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpixmap.h:205
// index:40
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPixmap &, QPixmap &)

/*
Swaps pixmap other with this pixmap. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap40(value1 QPixmap_ITF, value2 QPixmap_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPixmap_PTR() != nil {
		convArg0 = value1.QPixmap_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPixmap_PTR() != nil {
		convArg1 = value2.QPixmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR7QPixmapS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcursor.h:118
// index:41
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QCursor &, QCursor &)

/*
Swaps this cursor with the other cursor.
*/
func Swap41(value1 QCursor_ITF, value2 QCursor_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QCursor_PTR() != nil {
		convArg0 = value1.QCursor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QCursor_PTR() != nil {
		convArg1 = value2.QCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR7QCursorS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbitmap.h:84
// index:42
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QBitmap &, QBitmap &)

/*
Swaps bitmap other with this bitmap. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap42(value1 QBitmap_ITF, value2 QBitmap_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QBitmap_PTR() != nil {
		convArg0 = value1.QBitmap_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QBitmap_PTR() != nil {
		convArg1 = value2.QBitmap_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR7QBitmapS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qimage.h:379
// index:43
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QImage &, QImage &)

/*
Swaps image other with this image. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap43(value1 QImage_ITF, value2 QImage_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QImage_PTR() != nil {
		convArg0 = value1.QImage_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QImage_PTR() != nil {
		convArg1 = value2.QImage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR6QImageS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qbrush.h:137
// index:44
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QBrush &, QBrush &)

/*
Swaps brush other with this brush. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap44(value1 QBrush_ITF, value2 QBrush_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QBrush_PTR() != nil {
		convArg0 = value1.QBrush_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QBrush_PTR() != nil {
		convArg1 = value2.QBrush_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR6QBrushS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qicon.h:144
// index:45
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QIcon &, QIcon &)

/*
Swaps icon other with this icon. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap45(value1 QIcon_ITF, value2 QIcon_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QIcon_PTR() != nil {
		convArg0 = value1.QIcon_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QIcon_PTR() != nil {
		convArg1 = value2.QIcon_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR5QIconS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpen.h:133
// index:46
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPen &, QPen &)

/*
Swaps pen other with this pen. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap46(value1 QPen_ITF, value2 QPen_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPen_PTR() != nil {
		convArg0 = value1.QPen_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPen_PTR() != nil {
		convArg1 = value2.QPen_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR4QPenS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qopenglpixeltransferoptions.h:96
// index:47
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QOpenGLPixelTransferOptions &, QOpenGLPixelTransferOptions &)

/*

 */
func Swap47(value1 QOpenGLPixelTransferOptions_ITF, value2 QOpenGLPixelTransferOptions_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QOpenGLPixelTransferOptions_PTR() != nil {
		convArg0 = value1.QOpenGLPixelTransferOptions_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QOpenGLPixelTransferOptions_PTR() != nil {
		convArg1 = value2.QOpenGLPixelTransferOptions_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR27QOpenGLPixelTransferOptionsS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:982
// index:48
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextTableCellFormat &, QTextTableCellFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap48(value1 QTextTableCellFormat_ITF, value2 QTextTableCellFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextTableCellFormat_PTR() != nil {
		convArg0 = value1.QTextTableCellFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextTableCellFormat_PTR() != nil {
		convArg1 = value2.QTextTableCellFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR20QTextTableCellFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qopengldebug.h:144
// index:49
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QOpenGLDebugMessage &, QOpenGLDebugMessage &)

/*

 */
func Swap49(value1 QOpenGLDebugMessage_ITF, value2 QOpenGLDebugMessage_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QOpenGLDebugMessage_PTR() != nil {
		convArg0 = value1.QOpenGLDebugMessage_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QOpenGLDebugMessage_PTR() != nil {
		convArg1 = value2.QOpenGLDebugMessage_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR19QOpenGLDebugMessageS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:941
// index:50
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextTableFormat &, QTextTableFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap50(value1 QTextTableFormat_ITF, value2 QTextTableFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextTableFormat_PTR() != nil {
		convArg0 = value1.QTextTableFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextTableFormat_PTR() != nil {
		convArg1 = value2.QTextTableFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR16QTextTableFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:767
// index:51
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextImageFormat &, QTextImageFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap51(value1 QTextImageFormat_ITF, value2 QTextImageFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextImageFormat_PTR() != nil {
		convArg0 = value1.QTextImageFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextImageFormat_PTR() != nil {
		convArg1 = value2.QTextImageFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR16QTextImageFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:870
// index:52
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextFrameFormat &, QTextFrameFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap52(value1 QTextFrameFormat_ITF, value2 QTextFrameFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextFrameFormat_PTR() != nil {
		convArg0 = value1.QTextFrameFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextFrameFormat_PTR() != nil {
		convArg1 = value2.QTextFrameFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR16QTextFrameFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:660
// index:53
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextBlockFormat &, QTextBlockFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap53(value1 QTextBlockFormat_ITF, value2 QTextBlockFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextBlockFormat_PTR() != nil {
		convArg0 = value1.QTextBlockFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextBlockFormat_PTR() != nil {
		convArg1 = value2.QTextBlockFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR16QTextBlockFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:725
// index:54
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextListFormat &, QTextListFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap54(value1 QTextListFormat_ITF, value2 QTextListFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextListFormat_PTR() != nil {
		convArg0 = value1.QTextListFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextListFormat_PTR() != nil {
		convArg1 = value2.QTextListFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR15QTextListFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:563
// index:55
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextCharFormat &, QTextCharFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap55(value1 QTextCharFormat_ITF, value2 QTextCharFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextCharFormat_PTR() != nil {
		convArg0 = value1.QTextCharFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextCharFormat_PTR() != nil {
		convArg1 = value2.QTextCharFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR15QTextCharFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpainterpath.h:234
// index:56
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPainterPath &, QPainterPath &)

/*
Swaps painter path other with this painter path. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
func Swap56(value1 QPainterPath_ITF, value2 QPainterPath_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPainterPath_PTR() != nil {
		convArg0 = value1.QPainterPath_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPainterPath_PTR() != nil {
		convArg1 = value2.QPainterPath_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR12QPainterPathS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextformat.h:384
// index:57
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextFormat &, QTextFormat &)

/*
Swaps this text format with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap57(value1 QTextFormat_ITF, value2 QTextFormat_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextFormat_PTR() != nil {
		convArg0 = value1.QTextFormat_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextFormat_PTR() != nil {
		convArg1 = value2.QTextFormat_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR11QTextFormatS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextcursor.h:239
// index:58
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QTextCursor &, QTextCursor &)

/*
Swaps this text cursor instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap58(value1 QTextCursor_ITF, value2 QTextCursor_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QTextCursor_PTR() != nil {
		convArg0 = value1.QTextCursor_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QTextCursor_PTR() != nil {
		convArg1 = value2.QTextCursor_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR11QTextCursorS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qstatictext.h:104
// index:59
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QStaticText &, QStaticText &)

/*
Swaps this static text instance with other. This function is very fast and never fails.

This function was introduced in  Qt 5.0.
*/
func Swap59(value1 QStaticText_ITF, value2 QStaticText_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QStaticText_PTR() != nil {
		convArg0 = value1.QStaticText_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QStaticText_PTR() != nil {
		convArg1 = value2.QStaticText_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR11QStaticTextS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qpagelayout.h:141
// index:60
// Invalid inline Visibility=Default Availability=Available
// [-2] void swap(QPageLayout &, QPageLayout &)

/*
Swaps this page layout with other. This function is very fast and never fails.
*/
func Swap60(value1 QPageLayout_ITF, value2 QPageLayout_ITF) {
	var convArg0 unsafe.Pointer
	if value1 != nil && value1.QPageLayout_PTR() != nil {
		convArg0 = value1.QPageLayout_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if value2 != nil && value2.QPageLayout_PTR() != nil {
		convArg1 = value2.QPageLayout_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z4swapR11QPageLayoutS0_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qrgb.h:66
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qRgb(int, int, int)

/*

 */
func QRgb(r int, g int, b int) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z4qRgbiii", qtrt.FFI_TYPE_POINTER, r, g, b)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qrgb.h:54
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] int qRed(QRgb)

/*

 */
func QRed(rgb uint) int {
	rv, err := qtrt.InvokeQtFunc6("_Z4qRedj", qtrt.FFI_TYPE_POINTER, rgb)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qaccessible.h:974
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qAccessibleLocalizedActionDescription(const QString &)

/*

 */
func QAccessibleLocalizedActionDescription(actionName string) string {
	var tmpArg0 = qtcore.NewQString5(actionName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z37qAccessibleLocalizedActionDescriptionRK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qaccessible.h:973
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qAccessibleEventString(QAccessible::Event)

/*

 */
func QAccessibleEventString(event int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z22qAccessibleEventStringN11QAccessible5EventE", qtrt.FFI_TYPE_POINTER, event)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qaccessible.h:972
// index:0
// Invalid Visibility=Default Availability=Available
// [8] const char * qAccessibleRoleString(QAccessible::Role)

/*

 */
func QAccessibleRoleString(role int) string {
	rv, err := qtrt.InvokeQtFunc6("_Z21qAccessibleRoleStringN11QAccessible4RoleE", qtrt.FFI_TYPE_POINTER, role)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}

// /usr/include/qt/QtGui/qicon.h:155
// index:0
// Invalid Visibility=Default Availability=Available
// [8] QString qt_findAtNxFile(const QString &, qreal, qreal *)

/*

 */
func Qt_findAtNxFile(baseFileName string, targetDevicePixelRatio float64, sourceDevicePixelRatio unsafe.Pointer /*666*/) string {
	var tmpArg0 = qtcore.NewQString5(baseFileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z15qt_findAtNxFileRK7QStringdPd", qtrt.FFI_TYPE_POINTER, convArg0, targetDevicePixelRatio, sourceDevicePixelRatio)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	qtcore.DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtGui/qrgb.h:96
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qUnpremultiply(QRgb)

/*

 */
func QUnpremultiply(p uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z14qUnpremultiplyj", qtrt.FFI_TYPE_POINTER, p)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qquaternion.h:323
// index:3
// Invalid inline Visibility=Default Availability=Available
// [1] bool qFuzzyCompare(const QQuaternion &, const QQuaternion &)

/*
Returns true if q1 and q2 are equal, allowing for a small fuzziness factor for floating-point comparisons; false otherwise.
*/
func QFuzzyCompare3(q1 QQuaternion_ITF, q2 QQuaternion_ITF) bool {
	var convArg0 unsafe.Pointer
	if q1 != nil && q1.QQuaternion_PTR() != nil {
		convArg0 = q1.QQuaternion_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if q2 != nil && q2.QQuaternion_PTR() != nil {
		convArg1 = q2.QQuaternion_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_Z13qFuzzyCompareRK11QQuaternionS1_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtGui/qrgb.h:81
// index:0
// Invalid inline Visibility=Default Availability=Available
// [4] QRgb qPremultiply(QRgb)

/*

 */
func QPremultiply(x uint) uint {
	rv, err := qtrt.InvokeQtFunc6("_Z12qPremultiplyj", qtrt.FFI_TYPE_POINTER, x)
	qtrt.ErrPrint(err, rv)
	return uint(rv) // 222
}

//  body block end
