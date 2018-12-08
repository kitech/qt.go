package qtquick

// /usr/include/qt/QtQuick/qsgmaterial.h
// #include <qsgmaterial.h>
// #include <QtQuick>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtqml"

//  ext block end

//  body block begin

/*

 */
type QSGMaterial struct {
	*qtrt.CObject
}
type QSGMaterial_ITF interface {
	QSGMaterial_PTR() *QSGMaterial
}

func (ptr *QSGMaterial) QSGMaterial_PTR() *QSGMaterial { return ptr }

func (this *QSGMaterial) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QSGMaterial) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQSGMaterialFromPointer(cthis unsafe.Pointer) *QSGMaterial {
	return &QSGMaterial{&qtrt.CObject{cthis}}
}
func (*QSGMaterial) NewFromPointer(cthis unsafe.Pointer) *QSGMaterial {
	return NewQSGMaterialFromPointer(cthis)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:146
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QSGMaterial()

/*

 */
func (*QSGMaterial) NewForInherit() *QSGMaterial {
	return NewQSGMaterial()
}
func NewQSGMaterial() *QSGMaterial {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterialC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQSGMaterialFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQSGMaterial)
	return gothis
}

// /usr/include/qt/QtQuick/qsgmaterial.h:147
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QSGMaterial()

/*

 */
func DeleteQSGMaterial(this *QSGMaterial) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterialD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:149
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGMaterialType * type() const

/*
This function is called by the scene graph to return a unique instance per subclass.
*/
func (this *QSGMaterial) Type() *QSGMaterialType /*777 QSGMaterialType **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial4typeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialTypeFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgmaterial.h:150
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QSGMaterialShader * createShader() const

/*
This function returns a new instance of a the QSGMaterialShader implementatation used to render geometry for a specific implementation of QSGMaterial.

The function will be called only once for each material type that exists in the scene graph and will be cached internally.
*/
func (this *QSGMaterial) CreateShader() *QSGMaterialShader /*777 QSGMaterialShader **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial12createShaderEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQSGMaterialShaderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQuick/qsgmaterial.h:151
// index:0
// Public virtual Visibility=Default Availability=Available
// [4] int compare(const QSGMaterial *) const

/*
Compares this material to other and returns 0 if they are equal; -1 if this material should sort before other and 1 if other should sort before.

The scene graph can reorder geometry nodes to minimize state changes. The compare function is called during the sorting process so that the materials can be sorted to minimize state changes in each call to QSGMaterialShader::updateState().

The this pointer and other is guaranteed to have the same type().
*/
func (this *QSGMaterial) Compare(other QSGMaterial_ITF /*777 const QSGMaterial **/) int {
	var convArg0 unsafe.Pointer
	if other != nil && other.QSGMaterial_PTR() != nil {
		convArg0 = other.QSGMaterial_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial7compareEPKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtQuick/qsgmaterial.h:153
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QSGMaterial::Flags flags() const

/*
Returns the material's flags.
*/
func (this *QSGMaterial) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QSGMaterial5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QSGMaterial::Flags, bool)

/*
Sets the flags flags on this material if on is true; otherwise clears the attribute.
*/
func (this *QSGMaterial) SetFlag(flags int, on bool) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterial7setFlagE6QFlagsINS_4FlagEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, on)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQuick/qsgmaterial.h:154
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlag(QSGMaterial::Flags, bool)

/*
Sets the flags flags on this material if on is true; otherwise clears the attribute.
*/
func (this *QSGMaterial) SetFlagp(flags int) {
	// arg: 1, bool=Bool, =Invalid, , Invalid
	on := true
	rv, err := qtrt.InvokeQtFunc6("_ZN11QSGMaterial7setFlagE6QFlagsINS_4FlagEEb", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags, on)
	qtrt.ErrPrint(err, rv)
}

/*


 */
type QSGMaterial__Flag = int

//
const QSGMaterial__Blending QSGMaterial__Flag = 1

//
const QSGMaterial__RequiresDeterminant QSGMaterial__Flag = 2

//
const QSGMaterial__RequiresFullMatrixExceptTranslate QSGMaterial__Flag = 6

//
const QSGMaterial__RequiresFullMatrix QSGMaterial__Flag = 14

//
const QSGMaterial__CustomCompileStep QSGMaterial__Flag = 16

func (this *QSGMaterial) FlagItemName(val int) string {
	switch val {
	case QSGMaterial__Blending: // 1
		return "Blending"
	case QSGMaterial__RequiresDeterminant: // 2
		return "RequiresDeterminant"
	case QSGMaterial__RequiresFullMatrixExceptTranslate: // 6
		return "RequiresFullMatrixExceptTranslate"
	case QSGMaterial__RequiresFullMatrix: // 14
		return "RequiresFullMatrix"
	case QSGMaterial__CustomCompileStep: // 16
		return "CustomCompileStep"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QSGMaterial_FlagItemName(val int) string {
	var nilthis *QSGMaterial
	return nilthis.FlagItemName(val)
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

//  keep block end
