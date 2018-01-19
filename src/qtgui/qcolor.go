//  header block begin
// /usr/include/qt/QtGui/qcolor.h
// #include <qcolor.h>
// #include <QtGui>
package qtgui

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
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
	cthis unsafe.Pointer
}

// /usr/include/qt/QtGui/qcolor.h:70
// index:0
// inline
// void QColor()
func NewQColor() *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:71
// index:1
// void QColor(Qt::GlobalColor)
func NewQColor_1(color int) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2EN2Qt11GlobalColorE", ffiqt.FFI_TYPE_VOID, cthis, &color)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:72
// index:2
// inline
// void QColor(int, int, int, int)
func NewQColor_2(r int, g int, b int, a int) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2Eiiii", ffiqt.FFI_TYPE_VOID, cthis, &r, &g, &b, &a)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:73
// index:3
// void QColor(QRgb)
func NewQColor_3(rgb uint) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2Ej", ffiqt.FFI_TYPE_VOID, cthis, &rgb)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:74
// index:4
// void QColor(class QRgba64)
func NewQColor_4(rgba64 unsafe.Pointer) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2E7QRgba64", ffiqt.FFI_TYPE_VOID, cthis, rgba64)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:76
// index:5
// inline
// void QColor(const class QString &)
func NewQColor_5(name unsafe.Pointer) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2ERK7QString", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:78
// index:6
// inline
// void QColor(class QStringView)
func NewQColor_6(name unsafe.Pointer) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2E11QStringView", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:79
// index:7
// inline
// void QColor(const char *)
func NewQColor_7(aname unsafe.Pointer) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2EPKc", ffiqt.FFI_TYPE_VOID, cthis, aname)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:80
// index:8
// inline
// void QColor(class QLatin1String)
func NewQColor_8(name unsafe.Pointer) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2E13QLatin1String", ffiqt.FFI_TYPE_VOID, cthis, name)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:81
// index:9
// void QColor(enum QColor::Spec)
func NewQColor_9(spec int) *QColor {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColorC2ENS_4SpecE", ffiqt.FFI_TYPE_VOID, cthis, &spec)
	gopp.ErrPrint(err, rv)
	return &QColor{cthis}
}

// /usr/include/qt/QtGui/qcolor.h:95
// index:0
// bool isValid()
func (this *QColor) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7isValidEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:98
// index:0
// QString name()
func (this *QColor) Name() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4nameEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:99
// index:1
// QString name(enum QColor::NameFormat)
func (this *QColor) Name_1(format int) {
	// 1: (, QColor::NameFormat format), (&format)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4nameENS_10NameFormatE", ffiqt.FFI_TYPE_VOID, this.cthis, &format)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:102
// index:0
// void setNamedColor(const class QString &)
func (this *QColor) SetNamedColor(name unsafe.Pointer) {
	// 0: (, const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor13setNamedColorERK7QString", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:104
// index:1
// void setNamedColor(class QStringView)
func (this *QColor) SetNamedColor_1(name unsafe.Pointer) {
	// 1: (, QStringView name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor13setNamedColorE11QStringView", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:105
// index:2
// void setNamedColor(class QLatin1String)
func (this *QColor) SetNamedColor_2(name unsafe.Pointer) {
	// 2: (, QLatin1String name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor13setNamedColorE13QLatin1String", ffiqt.FFI_TYPE_VOID, this.cthis, name)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:107
// index:0
// static
// QStringList colorNames()
func (this *QColor) ColorNames() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor10colorNamesEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_ColorNames() {
	// 0: (), ()
	var nilthis *QColor
	nilthis.ColorNames()
}

// /usr/include/qt/QtGui/qcolor.h:109
// index:0
// inline
// QColor::Spec spec()
func (this *QColor) Spec() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4specEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:112
// index:0
// int alpha()
func (this *QColor) Alpha() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5alphaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:113
// index:0
// void setAlpha(int)
func (this *QColor) SetAlpha(alpha int) {
	// 0: (, int alpha), (&alpha)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setAlphaEi", ffiqt.FFI_TYPE_VOID, this.cthis, &alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:115
// index:0
// qreal alphaF()
func (this *QColor) AlphaF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6alphaFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:116
// index:0
// void setAlphaF(qreal)
func (this *QColor) SetAlphaF(alpha float64) {
	// 0: (, qreal alpha), (&alpha)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9setAlphaFEd", ffiqt.FFI_TYPE_VOID, this.cthis, &alpha)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:118
// index:0
// int red()
func (this *QColor) Red() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor3redEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:119
// index:0
// int green()
func (this *QColor) Green() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5greenEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:120
// index:0
// int blue()
func (this *QColor) Blue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4blueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:121
// index:0
// void setRed(int)
func (this *QColor) SetRed(red int) {
	// 0: (, int red), (&red)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setRedEi", ffiqt.FFI_TYPE_VOID, this.cthis, &red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:122
// index:0
// void setGreen(int)
func (this *QColor) SetGreen(green int) {
	// 0: (, int green), (&green)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setGreenEi", ffiqt.FFI_TYPE_VOID, this.cthis, &green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:123
// index:0
// void setBlue(int)
func (this *QColor) SetBlue(blue int) {
	// 0: (, int blue), (&blue)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setBlueEi", ffiqt.FFI_TYPE_VOID, this.cthis, &blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:125
// index:0
// qreal redF()
func (this *QColor) RedF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4redFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:126
// index:0
// qreal greenF()
func (this *QColor) GreenF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6greenFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:127
// index:0
// qreal blueF()
func (this *QColor) BlueF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5blueFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:128
// index:0
// void setRedF(qreal)
func (this *QColor) SetRedF(red float64) {
	// 0: (, qreal red), (&red)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setRedFEd", ffiqt.FFI_TYPE_VOID, this.cthis, &red)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:129
// index:0
// void setGreenF(qreal)
func (this *QColor) SetGreenF(green float64) {
	// 0: (, qreal green), (&green)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9setGreenFEd", ffiqt.FFI_TYPE_VOID, this.cthis, &green)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:130
// index:0
// void setBlueF(qreal)
func (this *QColor) SetBlueF(blue float64) {
	// 0: (, qreal blue), (&blue)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setBlueFEd", ffiqt.FFI_TYPE_VOID, this.cthis, &blue)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:132
// index:0
// void getRgb(int *, int *, int *, int *)
func (this *QColor) GetRgb(r unsafe.Pointer, g unsafe.Pointer, b unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, int * r, int * g, int * b, int * a), (r, g, b, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6getRgbEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, r, g, b, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:133
// index:0
// void setRgb(int, int, int, int)
func (this *QColor) SetRgb(r int, g int, b int, a int) {
	// 0: (, int r, int g, int b, int a), (&r, &g, &b, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setRgbEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &r, &g, &b, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:145
// index:1
// void setRgb(QRgb)
func (this *QColor) SetRgb_1(rgb uint) {
	// 1: (, QRgb rgb), (&rgb)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setRgbEj", ffiqt.FFI_TYPE_VOID, this.cthis, &rgb)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:135
// index:0
// void getRgbF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetRgbF(r unsafe.Pointer, g unsafe.Pointer, b unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, qreal * r, qreal * g, qreal * b, qreal * a), (r, g, b, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7getRgbFEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, r, g, b, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:136
// index:0
// void setRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) SetRgbF(r float64, g float64, b float64, a float64) {
	// 0: (, qreal r, qreal g, qreal b, qreal a), (&r, &g, &b, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setRgbFEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &r, &g, &b, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:138
// index:0
// QRgba64 rgba64()
func (this *QColor) Rgba64() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6rgba64Ev", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:139
// index:0
// void setRgba64(class QRgba64)
func (this *QColor) SetRgba64(rgba unsafe.Pointer) {
	// 0: (, QRgba64 rgba), (rgba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9setRgba64E7QRgba64", ffiqt.FFI_TYPE_VOID, this.cthis, rgba)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:141
// index:0
// QRgb rgba()
func (this *QColor) Rgba() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4rgbaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:142
// index:0
// void setRgba(QRgb)
func (this *QColor) SetRgba(rgba uint) {
	// 0: (, QRgb rgba), (&rgba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setRgbaEj", ffiqt.FFI_TYPE_VOID, this.cthis, &rgba)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:144
// index:0
// QRgb rgb()
func (this *QColor) Rgb() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor3rgbEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:147
// index:0
// int hue()
func (this *QColor) Hue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor3hueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:148
// index:0
// int saturation()
func (this *QColor) Saturation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor10saturationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:149
// index:0
// int hsvHue()
func (this *QColor) HsvHue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6hsvHueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:150
// index:0
// int hsvSaturation()
func (this *QColor) HsvSaturation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor13hsvSaturationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:151
// index:0
// int value()
func (this *QColor) Value() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5valueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:153
// index:0
// qreal hueF()
func (this *QColor) HueF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4hueFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:154
// index:0
// qreal saturationF()
func (this *QColor) SaturationF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor11saturationFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:155
// index:0
// qreal hsvHueF()
func (this *QColor) HsvHueF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7hsvHueFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:156
// index:0
// qreal hsvSaturationF()
func (this *QColor) HsvSaturationF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor14hsvSaturationFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:157
// index:0
// qreal valueF()
func (this *QColor) ValueF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6valueFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:159
// index:0
// void getHsv(int *, int *, int *, int *)
func (this *QColor) GetHsv(h unsafe.Pointer, s unsafe.Pointer, v unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, int * h, int * s, int * v, int * a), (h, s, v, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6getHsvEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, h, s, v, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:160
// index:0
// void setHsv(int, int, int, int)
func (this *QColor) SetHsv(h int, s int, v int, a int) {
	// 0: (, int h, int s, int v, int a), (&h, &s, &v, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setHsvEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &h, &s, &v, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:162
// index:0
// void getHsvF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetHsvF(h unsafe.Pointer, s unsafe.Pointer, v unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, qreal * h, qreal * s, qreal * v, qreal * a), (h, s, v, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7getHsvFEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, h, s, v, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:163
// index:0
// void setHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) SetHsvF(h float64, s float64, v float64, a float64) {
	// 0: (, qreal h, qreal s, qreal v, qreal a), (&h, &s, &v, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setHsvFEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &h, &s, &v, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:165
// index:0
// int cyan()
func (this *QColor) Cyan() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4cyanEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:166
// index:0
// int magenta()
func (this *QColor) Magenta() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7magentaEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:167
// index:0
// int yellow()
func (this *QColor) Yellow() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6yellowEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:168
// index:0
// int black()
func (this *QColor) Black() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5blackEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:170
// index:0
// qreal cyanF()
func (this *QColor) CyanF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5cyanFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:171
// index:0
// qreal magentaF()
func (this *QColor) MagentaF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor8magentaFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:172
// index:0
// qreal yellowF()
func (this *QColor) YellowF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7yellowFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:173
// index:0
// qreal blackF()
func (this *QColor) BlackF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6blackFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:175
// index:0
// void getCmyk(int *, int *, int *, int *, int *)
func (this *QColor) GetCmyk(c unsafe.Pointer, m unsafe.Pointer, y unsafe.Pointer, k unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, int * c, int * m, int * y, int * k, int * a), (c, m, y, k, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7getCmykEPiS0_S0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, c, m, y, k, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:176
// index:0
// void setCmyk(int, int, int, int, int)
func (this *QColor) SetCmyk(c int, m int, y int, k int, a int) {
	// 0: (, int c, int m, int y, int k, int a), (&c, &m, &y, &k, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setCmykEiiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &c, &m, &y, &k, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:178
// index:0
// void getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetCmykF(c unsafe.Pointer, m unsafe.Pointer, y unsafe.Pointer, k unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, qreal * c, qreal * m, qreal * y, qreal * k, qreal * a), (c, m, y, k, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8getCmykFEPdS0_S0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, c, m, y, k, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:179
// index:0
// void setCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) SetCmykF(c float64, m float64, y float64, k float64, a float64) {
	// 0: (, qreal c, qreal m, qreal y, qreal k, qreal a), (&c, &m, &y, &k, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8setCmykFEddddd", ffiqt.FFI_TYPE_VOID, this.cthis, &c, &m, &y, &k, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:181
// index:0
// int hslHue()
func (this *QColor) HslHue() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6hslHueEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:182
// index:0
// int hslSaturation()
func (this *QColor) HslSaturation() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor13hslSaturationEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:183
// index:0
// int lightness()
func (this *QColor) Lightness() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor9lightnessEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:185
// index:0
// qreal hslHueF()
func (this *QColor) HslHueF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7hslHueFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:186
// index:0
// qreal hslSaturationF()
func (this *QColor) HslSaturationF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor14hslSaturationFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:187
// index:0
// qreal lightnessF()
func (this *QColor) LightnessF() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor10lightnessFEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:189
// index:0
// void getHsl(int *, int *, int *, int *)
func (this *QColor) GetHsl(h unsafe.Pointer, s unsafe.Pointer, l unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, int * h, int * s, int * l, int * a), (h, s, l, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6getHslEPiS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, h, s, l, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:190
// index:0
// void setHsl(int, int, int, int)
func (this *QColor) SetHsl(h int, s int, l int, a int) {
	// 0: (, int h, int s, int l, int a), (&h, &s, &l, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor6setHslEiiii", ffiqt.FFI_TYPE_VOID, this.cthis, &h, &s, &l, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:192
// index:0
// void getHslF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetHslF(h unsafe.Pointer, s unsafe.Pointer, l unsafe.Pointer, a unsafe.Pointer) {
	// 0: (, qreal * h, qreal * s, qreal * l, qreal * a), (h, s, l, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7getHslFEPdS0_S0_S0_", ffiqt.FFI_TYPE_VOID, this.cthis, h, s, l, a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:193
// index:0
// void setHslF(qreal, qreal, qreal, qreal)
func (this *QColor) SetHslF(h float64, s float64, l float64, a float64) {
	// 0: (, qreal h, qreal s, qreal l, qreal a), (&h, &s, &l, &a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7setHslFEdddd", ffiqt.FFI_TYPE_VOID, this.cthis, &h, &s, &l, &a)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:195
// index:0
// QColor toRgb()
func (this *QColor) ToRgb() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5toRgbEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:196
// index:0
// QColor toHsv()
func (this *QColor) ToHsv() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5toHsvEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:197
// index:0
// QColor toCmyk()
func (this *QColor) ToCmyk() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6toCmykEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:198
// index:0
// QColor toHsl()
func (this *QColor) ToHsl() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5toHslEv", ffiqt.FFI_TYPE_VOID, this.cthis)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:200
// index:0
// QColor convertTo(enum QColor::Spec)
func (this *QColor) ConvertTo(colorSpec int) {
	// 0: (, QColor::Spec colorSpec), (&colorSpec)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor9convertToENS_4SpecE", ffiqt.FFI_TYPE_VOID, this.cthis, &colorSpec)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:202
// index:0
// static
// QColor fromRgb(QRgb)
func (this *QColor) FromRgb(rgb uint) {
	// 0: (QRgb rgb), (rgb)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromRgbEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromRgb(rgb uint) {
	// 0: (QRgb rgb), (rgb)
	var nilthis *QColor
	nilthis.FromRgb(rgb)
}

// /usr/include/qt/QtGui/qcolor.h:205
// index:1
// static
// QColor fromRgb(int, int, int, int)
func (this *QColor) FromRgb_1(r int, g int, b int, a int) {
	// 1: (int r, int g, int b, int a), (r, g, b, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromRgbEiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromRgb_1(r int, g int, b int, a int) {
	// 1: (int r, int g, int b, int a), (r, g, b, a)
	var nilthis *QColor
	nilthis.FromRgb_1(r, g, b, a)
}

// /usr/include/qt/QtGui/qcolor.h:203
// index:0
// static
// QColor fromRgba(QRgb)
func (this *QColor) FromRgba(rgba uint) {
	// 0: (QRgb rgba), (rgba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromRgbaEj", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromRgba(rgba uint) {
	// 0: (QRgb rgba), (rgba)
	var nilthis *QColor
	nilthis.FromRgba(rgba)
}

// /usr/include/qt/QtGui/qcolor.h:206
// index:0
// static
// QColor fromRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) FromRgbF(r float64, g float64, b float64, a float64) {
	// 0: (qreal r, qreal g, qreal b, qreal a), (r, g, b, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromRgbFEdddd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromRgbF(r float64, g float64, b float64, a float64) {
	// 0: (qreal r, qreal g, qreal b, qreal a), (r, g, b, a)
	var nilthis *QColor
	nilthis.FromRgbF(r, g, b, a)
}

// /usr/include/qt/QtGui/qcolor.h:208
// index:0
// static
// QColor fromRgba64(ushort, ushort, ushort, ushort)
func (this *QColor) FromRgba64(r uint16, g uint16, b uint16, a uint16) {
	// 0: (ushort r, ushort g, ushort b, ushort a), (r, g, b, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor10fromRgba64Etttt", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromRgba64(r uint16, g uint16, b uint16, a uint16) {
	// 0: (ushort r, ushort g, ushort b, ushort a), (r, g, b, a)
	var nilthis *QColor
	nilthis.FromRgba64(r, g, b, a)
}

// /usr/include/qt/QtGui/qcolor.h:209
// index:1
// static
// QColor fromRgba64(class QRgba64)
func (this *QColor) FromRgba64_1(rgba unsafe.Pointer) {
	// 1: (QRgba64 rgba), (rgba)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor10fromRgba64E7QRgba64", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromRgba64_1(rgba unsafe.Pointer) {
	// 1: (QRgba64 rgba), (rgba)
	var nilthis *QColor
	nilthis.FromRgba64_1(rgba)
}

// /usr/include/qt/QtGui/qcolor.h:211
// index:0
// static
// QColor fromHsv(int, int, int, int)
func (this *QColor) FromHsv(h int, s int, v int, a int) {
	// 0: (int h, int s, int v, int a), (h, s, v, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromHsvEiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromHsv(h int, s int, v int, a int) {
	// 0: (int h, int s, int v, int a), (h, s, v, a)
	var nilthis *QColor
	nilthis.FromHsv(h, s, v, a)
}

// /usr/include/qt/QtGui/qcolor.h:212
// index:0
// static
// QColor fromHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) FromHsvF(h float64, s float64, v float64, a float64) {
	// 0: (qreal h, qreal s, qreal v, qreal a), (h, s, v, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromHsvFEdddd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromHsvF(h float64, s float64, v float64, a float64) {
	// 0: (qreal h, qreal s, qreal v, qreal a), (h, s, v, a)
	var nilthis *QColor
	nilthis.FromHsvF(h, s, v, a)
}

// /usr/include/qt/QtGui/qcolor.h:214
// index:0
// static
// QColor fromCmyk(int, int, int, int, int)
func (this *QColor) FromCmyk(c int, m int, y int, k int, a int) {
	// 0: (int c, int m, int y, int k, int a), (c, m, y, k, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromCmykEiiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromCmyk(c int, m int, y int, k int, a int) {
	// 0: (int c, int m, int y, int k, int a), (c, m, y, k, a)
	var nilthis *QColor
	nilthis.FromCmyk(c, m, y, k, a)
}

// /usr/include/qt/QtGui/qcolor.h:215
// index:0
// static
// QColor fromCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) FromCmykF(c float64, m float64, y float64, k float64, a float64) {
	// 0: (qreal c, qreal m, qreal y, qreal k, qreal a), (c, m, y, k, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor9fromCmykFEddddd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromCmykF(c float64, m float64, y float64, k float64, a float64) {
	// 0: (qreal c, qreal m, qreal y, qreal k, qreal a), (c, m, y, k, a)
	var nilthis *QColor
	nilthis.FromCmykF(c, m, y, k, a)
}

// /usr/include/qt/QtGui/qcolor.h:217
// index:0
// static
// QColor fromHsl(int, int, int, int)
func (this *QColor) FromHsl(h int, s int, l int, a int) {
	// 0: (int h, int s, int l, int a), (h, s, l, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor7fromHslEiiii", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromHsl(h int, s int, l int, a int) {
	// 0: (int h, int s, int l, int a), (h, s, l, a)
	var nilthis *QColor
	nilthis.FromHsl(h, s, l, a)
}

// /usr/include/qt/QtGui/qcolor.h:218
// index:0
// static
// QColor fromHslF(qreal, qreal, qreal, qreal)
func (this *QColor) FromHslF(h float64, s float64, l float64, a float64) {
	// 0: (qreal h, qreal s, qreal l, qreal a), (h, s, l, a)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor8fromHslFEdddd", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_FromHslF(h float64, s float64, l float64, a float64) {
	// 0: (qreal h, qreal s, qreal l, qreal a), (h, s, l, a)
	var nilthis *QColor
	nilthis.FromHslF(h, s, l, a)
}

// /usr/include/qt/QtGui/qcolor.h:220
// index:0
// QColor light(int)
func (this *QColor) Light(f int) {
	// 0: (, int f), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor5lightEi", ffiqt.FFI_TYPE_VOID, this.cthis, &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:221
// index:0
// QColor lighter(int)
func (this *QColor) Lighter(f int) {
	// 0: (, int f), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor7lighterEi", ffiqt.FFI_TYPE_VOID, this.cthis, &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:222
// index:0
// QColor dark(int)
func (this *QColor) Dark(f int) {
	// 0: (, int f), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor4darkEi", ffiqt.FFI_TYPE_VOID, this.cthis, &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:223
// index:0
// QColor darker(int)
func (this *QColor) Darker(f int) {
	// 0: (, int f), (&f)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK6QColor6darkerEi", ffiqt.FFI_TYPE_VOID, this.cthis, &f)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qcolor.h:231
// index:0
// static
// bool isValidColor(const class QString &)
func (this *QColor) IsValidColor(name unsafe.Pointer) {
	// 0: (const QString & name), (name)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor12isValidColorERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_IsValidColor(name unsafe.Pointer) {
	// 0: (const QString & name), (name)
	var nilthis *QColor
	nilthis.IsValidColor(name)
}

// /usr/include/qt/QtGui/qcolor.h:233
// index:1
// static
// bool isValidColor(class QStringView)
func (this *QColor) IsValidColor_1(arg0 unsafe.Pointer) {
	// 1: (QStringView arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor12isValidColorE11QStringView", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_IsValidColor_1(arg0 unsafe.Pointer) {
	// 1: (QStringView arg0), (arg0)
	var nilthis *QColor
	nilthis.IsValidColor_1(arg0)
}

// /usr/include/qt/QtGui/qcolor.h:234
// index:2
// static
// bool isValidColor(class QLatin1String)
func (this *QColor) IsValidColor_2(arg0 unsafe.Pointer) {
	// 2: (QLatin1String arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN6QColor12isValidColorE13QLatin1String", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QColor_IsValidColor_2(arg0 unsafe.Pointer) {
	// 2: (QLatin1String arg0), (arg0)
	var nilthis *QColor
	nilthis.IsValidColor_2(arg0)
}

//  body block end
