package qtquick

import "unsafe"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

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
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtqml.KeepMe()
	}
}

//  header block end

//  body block begin
// /usr/include/qt/QtQuick/qsgnode.h:362
// index:90
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGNode::Flags::enum_type, int)

/*

 */
func Operator_or_90(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN7QSGNode4FlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:361
// index:91
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGNode::DirtyState::enum_type, int)

/*

 */
func Operator_or_91(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN7QSGNode13DirtyStateBitEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgsimpletexturenode.h:94
// index:92
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGSimpleTextureNode::TextureCoordinatesTransformMode::enum_type, int)

/*

 */
func Operator_or_92(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN20QSGSimpleTextureNode31TextureCoordinatesTransformFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:98
// index:93
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGRendererInterface::ShaderCompilationTypes::enum_type, int)

/*

 */
func Operator_or_93(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN20QSGRendererInterface21ShaderCompilationTypeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendererinterface.h:99
// index:94
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGRendererInterface::ShaderSourceTypes::enum_type, int)

/*

 */
func Operator_or_94(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN20QSGRendererInterface16ShaderSourceTypeEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgabstractrenderer.h:101
// index:95
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGAbstractRenderer::ClearMode::enum_type, int)

/*

 */
func Operator_or_95(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN19QSGAbstractRenderer12ClearModeBitEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgmaterial.h:163
// index:96
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGMaterialShader::RenderState::DirtyStates::enum_type, int)

/*

 */
func Operator_or_96(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN17QSGMaterialShader11RenderState10DirtyStateEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qquickpainteditem.h:134
// index:97
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QQuickPaintedItem::PerformanceHints::enum_type, int)

/*

 */
func Operator_or_97(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN17QQuickPaintedItem15PerformanceHintEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendernode.h:100
// index:98
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGRenderNode::StateFlags::enum_type, int)

/*

 */
func Operator_or_98(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QSGRenderNode9StateFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgrendernode.h:101
// index:99
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGRenderNode::RenderingFlags::enum_type, int)

/*

 */
func Operator_or_99(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN13QSGRenderNode13RenderingFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgimagenode.h:92
// index:100
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGImageNode::TextureCoordinatesTransformMode::enum_type, int)

/*

 */
func Operator_or_100(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN12QSGImageNode31TextureCoordinatesTransformFlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgmaterial.h:162
// index:101
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QSGMaterial::Flags::enum_type, int)

/*

 */
func Operator_or_101(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN11QSGMaterial4FlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qquickitem.h:464
// index:102
// Invalid inline Visibility=Default Availability=Available
// [4] QIncompatibleFlag operator|(QQuickItem::Flags::enum_type, int)

/*

 */
func Operator_or_102(f1 int, f2 int) *qtcore.QIncompatibleFlag /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZorN10QQuickItem4FlagEi", qtrt.FFI_TYPE_POINTER, f1, f2)
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQIncompatibleFlagFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQIncompatibleFlag)
	return rv2
}

// /usr/include/qt/QtQuick/qsgnode.h:193
// index:0
// Invalid Visibility=Default Availability=Available
// [-2] void qsgnode_set_description(QSGNode *, const QString &)

/*

 */
func Qsgnode_set_description(node QSGNode_ITF /*777 QSGNode **/, description string) {
	var convArg0 unsafe.Pointer
	if node != nil && node.QSGNode_PTR() != nil {
		convArg0 = node.QSGNode_PTR().GetCthis()
	}
	var tmpArg1 = qtcore.NewQString_5(description)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_Z23qsgnode_set_descriptionP7QSGNodeRK7QString", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

//  body block end
