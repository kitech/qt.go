package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 9
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

/*
 */
// size 16
type QVariant struct {
	*qtrt.CObject
}
type QVariant_ITF interface {
	QVariant_PTR() *QVariant
}

func (ptr *QVariant) QVariant_PTR() *QVariant { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QVariantFromptr(cthis Voidptr) *QVariant {
	return &QVariant{&qtrt.CObject{cthis}}
}
func (*QVariant) Fromptr(cthis Voidptr) *QVariant {
	return QVariantFromptr(cthis)
}

// /usr/include/qt/QtCore/qvariant.h:208
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QVariant()

/*
 */
func (*QVariant) NewForInherit() *QVariant {
	return NewQVariant()
}
func NewQVariant() *QVariant {
	cthis := qtrt.Malloc(16)
	rv, err := qtrt.Qtcc1(93067653, "_ZN8QVariantC2Ev", qtrt.FFITY_POINTER, cthis)
	qtrt.ErrPrint(err, rv)
	gothis := QVariantFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQVariant)
	return gothis
}

func DeleteQVariant(this *QVariant) {
	rv, err := qtrt.Qtcc1(0, "_ZN8QVariantD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QVariant__Type = int

//
const QVariant__Invalid QVariant__Type = 0

//
const QVariant__Bool QVariant__Type = 1

//
const QVariant__Int QVariant__Type = 2

//
const QVariant__UInt QVariant__Type = 3

//
const QVariant__LongLong QVariant__Type = 4

//
const QVariant__ULongLong QVariant__Type = 5

//
const QVariant__Double QVariant__Type = 6

//
const QVariant__Char QVariant__Type = 7

//
const QVariant__Map QVariant__Type = 8

//
const QVariant__List QVariant__Type = 9

//
const QVariant__String QVariant__Type = 10

//
const QVariant__StringList QVariant__Type = 11

//
const QVariant__ByteArray QVariant__Type = 12

//
const QVariant__BitArray QVariant__Type = 13

//
const QVariant__Date QVariant__Type = 14

//
const QVariant__Time QVariant__Type = 15

//
const QVariant__DateTime QVariant__Type = 16

//
const QVariant__Url QVariant__Type = 17

//
const QVariant__Locale QVariant__Type = 18

//
const QVariant__Rect QVariant__Type = 19

//
const QVariant__RectF QVariant__Type = 20

//
const QVariant__Size QVariant__Type = 21

//
const QVariant__SizeF QVariant__Type = 22

//
const QVariant__Line QVariant__Type = 23

//
const QVariant__LineF QVariant__Type = 24

//
const QVariant__Point QVariant__Type = 25

//
const QVariant__PointF QVariant__Type = 26

//
const QVariant__RegExp QVariant__Type = 27

//
const QVariant__RegularExpression QVariant__Type = 44

//
const QVariant__Hash QVariant__Type = 28

//
const QVariant__EasingCurve QVariant__Type = 29

//
const QVariant__Uuid QVariant__Type = 30

//
const QVariant__ModelIndex QVariant__Type = 42

//
const QVariant__PersistentModelIndex QVariant__Type = 50

//
const QVariant__LastCoreType QVariant__Type = 55

//
const QVariant__Font QVariant__Type = 64

//
const QVariant__Pixmap QVariant__Type = 65

//
const QVariant__Brush QVariant__Type = 66

//
const QVariant__Color QVariant__Type = 67

//
const QVariant__Palette QVariant__Type = 68

//
const QVariant__Image QVariant__Type = 70

//
const QVariant__Polygon QVariant__Type = 71

//
const QVariant__Region QVariant__Type = 72

//
const QVariant__Bitmap QVariant__Type = 73

//
const QVariant__Cursor QVariant__Type = 74

//
const QVariant__KeySequence QVariant__Type = 75

//
const QVariant__Pen QVariant__Type = 76

//
const QVariant__TextLength QVariant__Type = 77

//
const QVariant__TextFormat QVariant__Type = 78

//
const QVariant__Matrix QVariant__Type = 79

//
const QVariant__Transform QVariant__Type = 80

//
const QVariant__Matrix4x4 QVariant__Type = 81

//
const QVariant__Vector2D QVariant__Type = 82

//
const QVariant__Vector3D QVariant__Type = 83

//
const QVariant__Vector4D QVariant__Type = 84

//
const QVariant__Quaternion QVariant__Type = 85

//
const QVariant__PolygonF QVariant__Type = 86

//
const QVariant__Icon QVariant__Type = 69

//
const QVariant__LastGuiType QVariant__Type = 87

//
const QVariant__SizePolicy QVariant__Type = 121

//
const QVariant__UserType QVariant__Type = 1024

//
const QVariant__LastType QVariant__Type = -1

func (this *QVariant) TypeItemName(val int) string {
	switch val {
	case QVariant__Invalid: // 0
		return "Invalid"
	case QVariant__Bool: // 1
		return "Bool"
	case QVariant__Int: // 2
		return "Int"
	case QVariant__UInt: // 3
		return "UInt"
	case QVariant__LongLong: // 4
		return "LongLong"
	case QVariant__ULongLong: // 5
		return "ULongLong"
	case QVariant__Double: // 6
		return "Double"
	case QVariant__Char: // 7
		return "Char"
	case QVariant__Map: // 8
		return "Map"
	case QVariant__List: // 9
		return "List"
	case QVariant__String: // 10
		return "String"
	case QVariant__StringList: // 11
		return "StringList"
	case QVariant__ByteArray: // 12
		return "ByteArray"
	case QVariant__BitArray: // 13
		return "BitArray"
	case QVariant__Date: // 14
		return "Date"
	case QVariant__Time: // 15
		return "Time"
	case QVariant__DateTime: // 16
		return "DateTime"
	case QVariant__Url: // 17
		return "Url"
	case QVariant__Locale: // 18
		return "Locale"
	case QVariant__Rect: // 19
		return "Rect"
	case QVariant__RectF: // 20
		return "RectF"
	case QVariant__Size: // 21
		return "Size"
	case QVariant__SizeF: // 22
		return "SizeF"
	case QVariant__Line: // 23
		return "Line"
	case QVariant__LineF: // 24
		return "LineF"
	case QVariant__Point: // 25
		return "Point"
	case QVariant__PointF: // 26
		return "PointF"
	case QVariant__RegExp: // 27
		return "RegExp"
	case QVariant__RegularExpression: // 44
		return "RegularExpression"
	case QVariant__Hash: // 28
		return "Hash"
	case QVariant__EasingCurve: // 29
		return "EasingCurve"
	case QVariant__Uuid: // 30
		return "Uuid"
	case QVariant__ModelIndex: // 42
		return "ModelIndex"
	case QVariant__PersistentModelIndex: // 50
		return "PersistentModelIndex"
	case QVariant__LastCoreType: // 55
		return "LastCoreType"
	case QVariant__Font: // 64
		return "Font"
	case QVariant__Pixmap: // 65
		return "Pixmap"
	case QVariant__Brush: // 66
		return "Brush"
	case QVariant__Color: // 67
		return "Color"
	case QVariant__Palette: // 68
		return "Palette"
	case QVariant__Image: // 70
		return "Image"
	case QVariant__Polygon: // 71
		return "Polygon"
	case QVariant__Region: // 72
		return "Region"
	case QVariant__Bitmap: // 73
		return "Bitmap"
	case QVariant__Cursor: // 74
		return "Cursor"
	case QVariant__KeySequence: // 75
		return "KeySequence"
	case QVariant__Pen: // 76
		return "Pen"
	case QVariant__TextLength: // 77
		return "TextLength"
	case QVariant__TextFormat: // 78
		return "TextFormat"
	case QVariant__Matrix: // 79
		return "Matrix"
	case QVariant__Transform: // 80
		return "Transform"
	case QVariant__Matrix4x4: // 81
		return "Matrix4x4"
	case QVariant__Vector2D: // 82
		return "Vector2D"
	case QVariant__Vector3D: // 83
		return "Vector3D"
	case QVariant__Vector4D: // 84
		return "Vector4D"
	case QVariant__Quaternion: // 85
		return "Quaternion"
	case QVariant__PolygonF: // 86
		return "PolygonF"
	case QVariant__Icon: // 69
		return "Icon"
	case QVariant__LastGuiType: // 87
		return "LastGuiType"
	case QVariant__SizePolicy: // 121
		return "SizePolicy"
	case QVariant__UserType: // 1024
		return "UserType"
	case QVariant__LastType: // -1
		return "LastType"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QVariant_TypeItemName(val int) string {
	var nilthis *QVariant
	return nilthis.TypeItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10013() {
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
}

//  keep block end
