package qtgui

// /usr/include/qt/QtGui/qcolor.h
// #include <qcolor.h>
// #include <QtGui>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 23
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QColor struct {
	*qtrt.CObject
}

func (this *QColor) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QColor) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQColorFromPointer(cthis unsafe.Pointer) *QColor {
	return &QColor{&qtrt.CObject{cthis}}
}
func (*QColor) NewFromPointer(cthis unsafe.Pointer) *QColor {
	return NewQColorFromPointer(cthis)
}

// /usr/include/qt/QtGui/qcolor.h:70
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QColor()
func NewQColor() *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:71
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QColor(Qt::GlobalColor)
func NewQColor_1(color int) *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2EN2Qt11GlobalColorE", ffiqt.FFI_TYPE_POINTER, color)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:72
// index:2
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(int, int, int, int)
func NewQColor_2(r int, g int, b int, a int) *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2Eiiii", ffiqt.FFI_TYPE_POINTER, r, g, b, a)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:73
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QColor(QRgb)
func NewQColor_3(rgb uint) *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2Ej", ffiqt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:74
// index:4
// Public Visibility=Default Availability=Available
// [-2] void QColor(QRgba64)
func NewQColor_4(rgba64 *QRgba64 /*123*/) *QColor {
	var convArg0 = rgba64.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2E7QRgba64", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:76
// index:5
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(const QString &)
func NewQColor_5(name *qtcore.QString) *QColor {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2ERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:78
// index:6
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(QStringView)
func NewQColor_6(name *qtcore.QStringView /*123*/) *QColor {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2E11QStringView", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:79
// index:7
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(const char *)
func NewQColor_7(aname string) *QColor {
	var convArg0 = qtrt.CString(aname)
	defer qtrt.FreeMem(convArg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2EPKc", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:80
// index:8
// Public inline Visibility=Default Availability=Available
// [-2] void QColor(QLatin1String)
func NewQColor_8(name *qtcore.QLatin1String /*123*/) *QColor {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2E13QLatin1String", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:81
// index:9
// Public Visibility=Default Availability=Available
// [-2] void QColor(enum QColor::Spec)
func NewQColor_9(spec int) *QColor {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2ENS_4SpecE", ffiqt.FFI_TYPE_POINTER, spec)
	gopp.ErrPrint(err, rv)
	gothis := NewQColorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQColor)
	return gothis
}

// /usr/include/qt/QtGui/qcolor.h:95
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QColor) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtGui/qcolor.h:98
// index:0
// Public Visibility=Default Availability=Available
// [8] QString name()
func (this *QColor) Name() *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4nameEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:99
// index:1
// Public Visibility=Default Availability=Available
// [8] QString name(enum QColor::NameFormat)
func (this *QColor) Name_1(format int) *qtcore.QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4nameENS_10NameFormatE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), format)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := qtcore.NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQString)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:102
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setNamedColor(const QString &)
func (this *QColor) SetNamedColor(name *qtcore.QString) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor13setNamedColorERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:104
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setNamedColor(QStringView)
func (this *QColor) SetNamedColor_1(name *qtcore.QStringView /*123*/) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor13setNamedColorE11QStringView", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:105
// index:2
// Public Visibility=Default Availability=Available
// [-2] void setNamedColor(QLatin1String)
func (this *QColor) SetNamedColor_2(name *qtcore.QLatin1String /*123*/) {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor13setNamedColorE13QLatin1String", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:109
// index:0
// Public inline Visibility=Default Availability=Available
// [4] QColor::Spec spec()
func (this *QColor) Spec() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4specEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return int(rv)
}

// /usr/include/qt/QtGui/qcolor.h:112
// index:0
// Public Visibility=Default Availability=Available
// [4] int alpha()
func (this *QColor) Alpha() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5alphaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:113
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlpha(int)
func (this *QColor) SetAlpha(alpha int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setAlphaEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:115
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal alphaF()
func (this *QColor) AlphaF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6alphaFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:116
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAlphaF(qreal)
func (this *QColor) SetAlphaF(alpha float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9setAlphaFEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:118
// index:0
// Public Visibility=Default Availability=Available
// [4] int red()
func (this *QColor) Red() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor3redEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:119
// index:0
// Public Visibility=Default Availability=Available
// [4] int green()
func (this *QColor) Green() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5greenEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:120
// index:0
// Public Visibility=Default Availability=Available
// [4] int blue()
func (this *QColor) Blue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4blueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:121
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRed(int)
func (this *QColor) SetRed(red int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setRedEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:122
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGreen(int)
func (this *QColor) SetGreen(green int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setGreenEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlue(int)
func (this *QColor) SetBlue(blue int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setBlueEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:125
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal redF()
func (this *QColor) RedF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4redFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:126
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal greenF()
func (this *QColor) GreenF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6greenFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:127
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal blueF()
func (this *QColor) BlueF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5blueFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:128
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRedF(qreal)
func (this *QColor) SetRedF(red float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setRedFEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:129
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setGreenF(qreal)
func (this *QColor) SetGreenF(green float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9setGreenFEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:130
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setBlueF(qreal)
func (this *QColor) SetBlueF(blue float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setBlueFEd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:132
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getRgb(int *, int *, int *, int *)
func (this *QColor) GetRgb(r unsafe.Pointer /*666*/, g unsafe.Pointer /*666*/, b unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6getRgbEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &r, &g, &b, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:133
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRgb(int, int, int, int)
func (this *QColor) SetRgb(r int, g int, b int, a int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setRgbEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), r, g, b, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:145
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setRgb(QRgb)
func (this *QColor) SetRgb_1(rgb uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setRgbEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rgb)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getRgbF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetRgbF(r unsafe.Pointer /*666*/, g unsafe.Pointer /*666*/, b unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7getRgbFEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &r, &g, &b, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:136
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) SetRgbF(r float64, g float64, b float64, a float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setRgbFEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), r, g, b, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QRgba64 rgba64()
func (this *QColor) Rgba64() *QRgba64 /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6rgba64Ev", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQRgba64FromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQRgba64)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:139
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRgba64(QRgba64)
func (this *QColor) SetRgba64(rgba *QRgba64 /*123*/) {
	var convArg0 = rgba.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9setRgba64E7QRgba64", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:141
// index:0
// Public Visibility=Default Availability=Available
// [4] QRgb rgba()
func (this *QColor) Rgba() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4rgbaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qcolor.h:142
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setRgba(QRgb)
func (this *QColor) SetRgba(rgba uint) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setRgbaEj", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), rgba)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:144
// index:0
// Public Visibility=Default Availability=Available
// [4] QRgb rgb()
func (this *QColor) Rgb() uint {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor3rgbEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return uint(rv) // 222
}

// /usr/include/qt/QtGui/qcolor.h:147
// index:0
// Public Visibility=Default Availability=Available
// [4] int hue()
func (this *QColor) Hue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor3hueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:148
// index:0
// Public Visibility=Default Availability=Available
// [4] int saturation()
func (this *QColor) Saturation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor10saturationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:149
// index:0
// Public Visibility=Default Availability=Available
// [4] int hsvHue()
func (this *QColor) HsvHue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6hsvHueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:150
// index:0
// Public Visibility=Default Availability=Available
// [4] int hsvSaturation()
func (this *QColor) HsvSaturation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor13hsvSaturationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:151
// index:0
// Public Visibility=Default Availability=Available
// [4] int value()
func (this *QColor) Value() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5valueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:153
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal hueF()
func (this *QColor) HueF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4hueFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:154
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal saturationF()
func (this *QColor) SaturationF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor11saturationFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:155
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal hsvHueF()
func (this *QColor) HsvHueF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7hsvHueFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:156
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal hsvSaturationF()
func (this *QColor) HsvSaturationF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor14hsvSaturationFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:157
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal valueF()
func (this *QColor) ValueF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6valueFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:159
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getHsv(int *, int *, int *, int *)
func (this *QColor) GetHsv(h unsafe.Pointer /*666*/, s unsafe.Pointer /*666*/, v unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6getHsvEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &h, &s, &v, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:160
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHsv(int, int, int, int)
func (this *QColor) SetHsv(h int, s int, v int, a int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setHsvEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h, s, v, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:162
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getHsvF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetHsvF(h unsafe.Pointer /*666*/, s unsafe.Pointer /*666*/, v unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7getHsvFEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &h, &s, &v, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:163
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) SetHsvF(h float64, s float64, v float64, a float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setHsvFEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h, s, v, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:165
// index:0
// Public Visibility=Default Availability=Available
// [4] int cyan()
func (this *QColor) Cyan() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4cyanEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:166
// index:0
// Public Visibility=Default Availability=Available
// [4] int magenta()
func (this *QColor) Magenta() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7magentaEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:167
// index:0
// Public Visibility=Default Availability=Available
// [4] int yellow()
func (this *QColor) Yellow() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6yellowEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:168
// index:0
// Public Visibility=Default Availability=Available
// [4] int black()
func (this *QColor) Black() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5blackEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:170
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal cyanF()
func (this *QColor) CyanF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5cyanFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:171
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal magentaF()
func (this *QColor) MagentaF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor8magentaFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:172
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal yellowF()
func (this *QColor) YellowF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7yellowFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:173
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal blackF()
func (this *QColor) BlackF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6blackFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:175
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getCmyk(int *, int *, int *, int *, int *)
func (this *QColor) GetCmyk(c unsafe.Pointer /*666*/, m unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, k unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7getCmykEPiS0_S0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c, &m, &y, &k, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCmyk(int, int, int, int, int)
func (this *QColor) SetCmyk(c int, m int, y int, k int, a int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setCmykEiiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), c, m, y, k, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:178
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetCmykF(c unsafe.Pointer /*666*/, m unsafe.Pointer /*666*/, y unsafe.Pointer /*666*/, k unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8getCmykFEPdS0_S0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &c, &m, &y, &k, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:179
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) SetCmykF(c float64, m float64, y float64, k float64, a float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setCmykFEddddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), c, m, y, k, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:181
// index:0
// Public Visibility=Default Availability=Available
// [4] int hslHue()
func (this *QColor) HslHue() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6hslHueEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:182
// index:0
// Public Visibility=Default Availability=Available
// [4] int hslSaturation()
func (this *QColor) HslSaturation() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor13hslSaturationEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:183
// index:0
// Public Visibility=Default Availability=Available
// [4] int lightness()
func (this *QColor) Lightness() int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor9lightnessEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:185
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal hslHueF()
func (this *QColor) HslHueF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7hslHueFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:186
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal hslSaturationF()
func (this *QColor) HslSaturationF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor14hslSaturationFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:187
// index:0
// Public Visibility=Default Availability=Available
// [8] qreal lightnessF()
func (this *QColor) LightnessF() float64 {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor10lightnessFEv", ffiqt.FFI_TYPE_DOUBLE, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("float64", rv).(float64) // 1111
}

// /usr/include/qt/QtGui/qcolor.h:189
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getHsl(int *, int *, int *, int *)
func (this *QColor) GetHsl(h unsafe.Pointer /*666*/, s unsafe.Pointer /*666*/, l unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6getHslEPiS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &h, &s, &l, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:190
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHsl(int, int, int, int)
func (this *QColor) SetHsl(h int, s int, l int, a int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setHslEiiii", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h, s, l, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:192
// index:0
// Public Visibility=Default Availability=Available
// [-2] void getHslF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetHslF(h unsafe.Pointer /*666*/, s unsafe.Pointer /*666*/, l unsafe.Pointer /*666*/, a unsafe.Pointer /*666*/) {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7getHslFEPdS0_S0_S0_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), &h, &s, &l, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:193
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHslF(qreal, qreal, qreal, qreal)
func (this *QColor) SetHslF(h float64, s float64, l float64, a float64) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setHslFEdddd", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), h, s, l, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:195
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor toRgb()
func (this *QColor) ToRgb() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5toRgbEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:196
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor toHsv()
func (this *QColor) ToHsv() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5toHsvEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:197
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor toCmyk()
func (this *QColor) ToCmyk() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6toCmykEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:198
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor toHsl()
func (this *QColor) ToHsl() *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5toHslEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:200
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor convertTo(enum QColor::Spec)
func (this *QColor) ConvertTo(colorSpec int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor9convertToENS_4SpecE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), colorSpec)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:202
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromRgb(QRgb)
func (this *QColor) FromRgb(rgb uint) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromRgbEj", ffiqt.FFI_TYPE_POINTER, rgb)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromRgb(rgb uint) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromRgb(rgb)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:205
// index:1
// Public static Visibility=Default Availability=Available
// [16] QColor fromRgb(int, int, int, int)
func (this *QColor) FromRgb_1(r int, g int, b int, a int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromRgbEiiii", ffiqt.FFI_TYPE_POINTER, r, g, b, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromRgb_1(r int, g int, b int, a int) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromRgb_1(r, g, b, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:203
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromRgba(QRgb)
func (this *QColor) FromRgba(rgba uint) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromRgbaEj", ffiqt.FFI_TYPE_POINTER, rgba)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromRgba(rgba uint) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromRgba(rgba)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:206
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) FromRgbF(r float64, g float64, b float64, a float64) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromRgbFEdddd", ffiqt.FFI_TYPE_POINTER, r, g, b, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromRgbF(r float64, g float64, b float64, a float64) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromRgbF(r, g, b, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:208
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromRgba64(ushort, ushort, ushort, ushort)
func (this *QColor) FromRgba64(r uint16, g uint16, b uint16, a uint16) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor10fromRgba64Etttt", ffiqt.FFI_TYPE_POINTER, r, g, b, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromRgba64(r uint16, g uint16, b uint16, a uint16) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromRgba64(r, g, b, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:209
// index:1
// Public static Visibility=Default Availability=Available
// [16] QColor fromRgba64(QRgba64)
func (this *QColor) FromRgba64_1(rgba *QRgba64 /*123*/) *QColor /*123*/ {
	var convArg0 = rgba.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor10fromRgba64E7QRgba64", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromRgba64_1(rgba *QRgba64 /*123*/) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromRgba64_1(rgba)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:211
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromHsv(int, int, int, int)
func (this *QColor) FromHsv(h int, s int, v int, a int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromHsvEiiii", ffiqt.FFI_TYPE_POINTER, h, s, v, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromHsv(h int, s int, v int, a int) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromHsv(h, s, v, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:212
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) FromHsvF(h float64, s float64, v float64, a float64) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromHsvFEdddd", ffiqt.FFI_TYPE_POINTER, h, s, v, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromHsvF(h float64, s float64, v float64, a float64) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromHsvF(h, s, v, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:214
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromCmyk(int, int, int, int, int)
func (this *QColor) FromCmyk(c int, m int, y int, k int, a int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromCmykEiiiii", ffiqt.FFI_TYPE_POINTER, c, m, y, k, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromCmyk(c int, m int, y int, k int, a int) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromCmyk(c, m, y, k, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:215
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) FromCmykF(c float64, m float64, y float64, k float64, a float64) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9fromCmykFEddddd", ffiqt.FFI_TYPE_POINTER, c, m, y, k, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromCmykF(c float64, m float64, y float64, k float64, a float64) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromCmykF(c, m, y, k, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:217
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromHsl(int, int, int, int)
func (this *QColor) FromHsl(h int, s int, l int, a int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromHslEiiii", ffiqt.FFI_TYPE_POINTER, h, s, l, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromHsl(h int, s int, l int, a int) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromHsl(h, s, l, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:218
// index:0
// Public static Visibility=Default Availability=Available
// [16] QColor fromHslF(qreal, qreal, qreal, qreal)
func (this *QColor) FromHslF(h float64, s float64, l float64, a float64) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromHslFEdddd", ffiqt.FFI_TYPE_POINTER, h, s, l, a)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}
func QColor_FromHslF(h float64, s float64, l float64, a float64) *QColor /*123*/ {
	var nilthis *QColor
	rv := nilthis.FromHslF(h, s, l, a)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:220
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor light(int)
func (this *QColor) Light(f int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5lightEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:221
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor lighter(int)
func (this *QColor) Lighter(f int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7lighterEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:222
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor dark(int)
func (this *QColor) Dark(f int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4darkEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:223
// index:0
// Public Visibility=Default Availability=Available
// [16] QColor darker(int)
func (this *QColor) Darker(f int) *QColor /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6darkerEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), f)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQColorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQColor)
	return rv2
}

// /usr/include/qt/QtGui/qcolor.h:231
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isValidColor(const QString &)
func (this *QColor) IsValidColor(name *qtcore.QString) bool {
	var convArg0 = name.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor12isValidColorERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QColor_IsValidColor(name *qtcore.QString) bool {
	var nilthis *QColor
	rv := nilthis.IsValidColor(name)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:233
// index:1
// Public static Visibility=Default Availability=Available
// [1] bool isValidColor(QStringView)
func (this *QColor) IsValidColor_1(arg0 *qtcore.QStringView /*123*/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor12isValidColorE11QStringView", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QColor_IsValidColor_1(arg0 *qtcore.QStringView /*123*/) bool {
	var nilthis *QColor
	rv := nilthis.IsValidColor_1(arg0)
	return rv
}

// /usr/include/qt/QtGui/qcolor.h:234
// index:2
// Public static Visibility=Default Availability=Available
// [1] bool isValidColor(QLatin1String)
func (this *QColor) IsValidColor_2(arg0 *qtcore.QLatin1String /*123*/) bool {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor12isValidColorE13QLatin1String", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	return rv != 0
}
func QColor_IsValidColor_2(arg0 *qtcore.QLatin1String /*123*/) bool {
	var nilthis *QColor
	rv := nilthis.IsValidColor_2(arg0)
	return rv
}

func DeleteQColor(this *QColor) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

type QColor__Spec = int

const QColor__Invalid QColor__Spec = 0
const QColor__Rgb QColor__Spec = 1
const QColor__Hsv QColor__Spec = 2
const QColor__Cmyk QColor__Spec = 3
const QColor__Hsl QColor__Spec = 4

type QColor__NameFormat = int

const QColor__HexRgb QColor__NameFormat = 0
const QColor__HexArgb QColor__NameFormat = 1

//  body block end
