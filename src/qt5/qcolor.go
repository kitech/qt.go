package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qcolor.h
// dst-file: /src/gui/qcolor.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
import "fmt"
import "reflect"
import "unsafe"
import "qtrt"
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QColor::getHsvF(qreal * h, qreal * s, qreal * v, qreal * a);
extern void _ZNK6QColor7getHsvFEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  int QColor::alpha();
extern void _ZNK6QColor5alphaEv(void* qthis);
  // proto:  qreal QColor::hslSaturationF();
extern void _ZNK6QColor14hslSaturationFEv(void* qthis);
  // proto:  void QColor::setAlphaF(qreal alpha);
extern void _ZN6QColor9setAlphaFEd(void* qthis, double arg0);
  // proto:  void QColor::getRgb(int * r, int * g, int * b, int * a);
extern void _ZNK6QColor6getRgbEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  int QColor::hslHue();
extern void _ZNK6QColor6hslHueEv(void* qthis);
  // proto:  int QColor::lightness();
extern void _ZNK6QColor9lightnessEv(void* qthis);
  // proto:  void QColor::setAlpha(int alpha);
extern void _ZN6QColor8setAlphaEi(void* qthis, int arg0);
  // proto: static QColor QColor::fromHslF(qreal h, qreal s, qreal l, qreal a);
extern void _ZN6QColor8fromHslFEdddd(double arg0, double arg1, double arg2, double arg3);
  // proto:  void QColor::getCmyk(int * c, int * m, int * y, int * k, int * a);
extern void _ZN6QColor7getCmykEPiS0_S0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3, int* arg4);
  // proto:  int QColor::green();
extern void _ZNK6QColor5greenEv(void* qthis);
  // proto:  int QColor::hsvSaturation();
extern void _ZNK6QColor13hsvSaturationEv(void* qthis);
  // proto:  QColor QColor::toHsl();
extern void _ZNK6QColor5toHslEv(void* qthis);
  // proto:  void QColor::QColor();
extern void* dector_ZN6QColorC1Ev();
extern void _ZN6QColorC1Ev(void* qthis);
  // proto:  void QColor::QColor(const char * name);
extern void* dector_ZN6QColorC1EPKc(char* arg0);
extern void _ZN6QColorC1EPKc(void* qthis, char* arg0);
  // proto:  void QColor::setBlue(int blue);
extern void _ZN6QColor7setBlueEi(void* qthis, int arg0);
  // proto:  int QColor::cyan();
extern void _ZNK6QColor4cyanEv(void* qthis);
  // proto:  void QColor::setCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
extern void _ZN6QColor8setCmykFEddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4);
  // proto: static QColor QColor::fromCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
extern void _ZN6QColor9fromCmykFEddddd(double arg0, double arg1, double arg2, double arg3, double arg4);
  // proto:  QColor QColor::light(int f);
extern void _ZNK6QColor5lightEi(void* qthis, int arg0);
  // proto:  void QColor::getHslF(qreal * h, qreal * s, qreal * l, qreal * a);
extern void _ZNK6QColor7getHslFEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto: static QColor QColor::fromRgb(QRgb rgb);
extern void _ZN6QColor7fromRgbEj(unsigned int arg0);
  // proto:  int QColor::yellow();
extern void _ZNK6QColor6yellowEv(void* qthis);
  // proto:  void QColor::getRgbF(qreal * r, qreal * g, qreal * b, qreal * a);
extern void _ZNK6QColor7getRgbFEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3);
  // proto:  void QColor::setRgb(int r, int g, int b, int a);
extern void _ZN6QColor6setRgbEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  QString QColor::name();
extern void _ZNK6QColor4nameEv(void* qthis);
  // proto:  qreal QColor::redF();
extern void _ZNK6QColor4redFEv(void* qthis);
  // proto:  qreal QColor::blackF();
extern void _ZNK6QColor6blackFEv(void* qthis);
  // proto:  void QColor::setHsvF(qreal h, qreal s, qreal v, qreal a);
extern void _ZN6QColor7setHsvFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  void QColor::setRgb(QRgb rgb);
extern void _ZN6QColor6setRgbEj(void* qthis, unsigned int arg0);
  // proto: static QColor QColor::fromRgb(int r, int g, int b, int a);
extern void _ZN6QColor7fromRgbEiiii(int arg0, int arg1, int arg2, int arg3);
  // proto:  qreal QColor::hsvHueF();
extern void _ZNK6QColor7hsvHueFEv(void* qthis);
  // proto:  qreal QColor::hsvSaturationF();
extern void _ZNK6QColor14hsvSaturationFEv(void* qthis);
  // proto:  qreal QColor::yellowF();
extern void _ZNK6QColor7yellowFEv(void* qthis);
  // proto:  int QColor::black();
extern void _ZNK6QColor5blackEv(void* qthis);
  // proto:  void QColor::setGreenF(qreal green);
extern void _ZN6QColor9setGreenFEd(void* qthis, double arg0);
  // proto:  QRgb QColor::rgba();
extern void _ZNK6QColor4rgbaEv(void* qthis);
  // proto:  QColor QColor::toCmyk();
extern void _ZNK6QColor6toCmykEv(void* qthis);
  // proto:  qreal QColor::greenF();
extern void _ZNK6QColor6greenFEv(void* qthis);
  // proto:  int QColor::red();
extern void _ZNK6QColor3redEv(void* qthis);
  // proto:  void QColor::setRgbF(qreal r, qreal g, qreal b, qreal a);
extern void _ZN6QColor7setRgbFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  qreal QColor::lightnessF();
extern void _ZNK6QColor10lightnessFEv(void* qthis);
  // proto:  QColor QColor::toHsv();
extern void _ZNK6QColor5toHsvEv(void* qthis);
  // proto:  void QColor::QColor(const QColor & color);
extern void* dector_ZN6QColorC1ERKS_(void* arg0);
extern void _ZN6QColorC1ERKS_(void* qthis, void* arg0);
  // proto: static QColor QColor::fromHsv(int h, int s, int v, int a);
extern void _ZN6QColor7fromHsvEiiii(int arg0, int arg1, int arg2, int arg3);
  // proto:  qreal QColor::hueF();
extern void _ZNK6QColor4hueFEv(void* qthis);
  // proto:  void QColor::setBlueF(qreal blue);
extern void _ZN6QColor8setBlueFEd(void* qthis, double arg0);
  // proto:  qreal QColor::saturationF();
extern void _ZNK6QColor11saturationFEv(void* qthis);
  // proto:  bool QColor::isValid();
extern void _ZNK6QColor7isValidEv(void* qthis);
  // proto:  QColor QColor::darker(int f);
extern void _ZNK6QColor6darkerEi(void* qthis, int arg0);
  // proto:  qreal QColor::blueF();
extern void _ZNK6QColor5blueFEv(void* qthis);
  // proto:  int QColor::hue();
extern void _ZNK6QColor3hueEv(void* qthis);
  // proto:  void QColor::setRgba(QRgb rgba);
extern void _ZN6QColor7setRgbaEj(void* qthis, unsigned int arg0);
  // proto:  void QColor::setNamedColor(const QString & name);
extern void _ZN6QColor13setNamedColorERK7QString(void* qthis, void* arg0);
  // proto:  int QColor::magenta();
extern void _ZNK6QColor7magentaEv(void* qthis);
  // proto:  QColor QColor::lighter(int f);
extern void _ZNK6QColor7lighterEi(void* qthis, int arg0);
  // proto:  QColor QColor::toRgb();
extern void _ZNK6QColor5toRgbEv(void* qthis);
  // proto:  qreal QColor::magentaF();
extern void _ZNK6QColor8magentaFEv(void* qthis);
  // proto:  qreal QColor::hslHueF();
extern void _ZNK6QColor7hslHueFEv(void* qthis);
  // proto: static QColor QColor::fromCmyk(int c, int m, int y, int k, int a);
extern void _ZN6QColor8fromCmykEiiiii(int arg0, int arg1, int arg2, int arg3, int arg4);
  // proto:  void QColor::setCmyk(int c, int m, int y, int k, int a);
extern void _ZN6QColor7setCmykEiiiii(void* qthis, int arg0, int arg1, int arg2, int arg3, int arg4);
  // proto: static QStringList QColor::colorNames();
extern void _ZN6QColor10colorNamesEv();
  // proto:  void QColor::getHsv(int * h, int * s, int * v, int * a);
extern void _ZNK6QColor6getHsvEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto:  void QColor::getCmykF(qreal * c, qreal * m, qreal * y, qreal * k, qreal * a);
extern void _ZN6QColor8getCmykFEPdS0_S0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3, double* arg4);
  // proto:  void QColor::setRed(int red);
extern void _ZN6QColor6setRedEi(void* qthis, int arg0);
  // proto:  void QColor::QColor(const QString & name);
extern void* dector_ZN6QColorC1ERK7QString(void* arg0);
extern void _ZN6QColorC1ERK7QString(void* qthis, void* arg0);
  // proto: static QColor QColor::fromRgba(QRgb rgba);
extern void _ZN6QColor8fromRgbaEj(unsigned int arg0);
  // proto:  void QColor::setHsv(int h, int s, int v, int a);
extern void _ZN6QColor6setHsvEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  QRgb QColor::rgb();
extern void _ZNK6QColor3rgbEv(void* qthis);
  // proto:  void QColor::setHslF(qreal h, qreal s, qreal l, qreal a);
extern void _ZN6QColor7setHslFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3);
  // proto:  int QColor::saturation();
extern void _ZNK6QColor10saturationEv(void* qthis);
  // proto:  void QColor::QColor(int r, int g, int b, int a);
extern void* dector_ZN6QColorC1Eiiii(int arg0, int arg1, int arg2, int arg3);
extern void _ZN6QColorC1Eiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  qreal QColor::alphaF();
extern void _ZNK6QColor6alphaFEv(void* qthis);
  // proto:  int QColor::value();
extern void _ZNK6QColor5valueEv(void* qthis);
  // proto: static QColor QColor::fromHsvF(qreal h, qreal s, qreal v, qreal a);
extern void _ZN6QColor8fromHsvFEdddd(double arg0, double arg1, double arg2, double arg3);
  // proto:  QColor QColor::dark(int f);
extern void _ZNK6QColor4darkEi(void* qthis, int arg0);
  // proto:  void QColor::setRedF(qreal red);
extern void _ZN6QColor7setRedFEd(void* qthis, double arg0);
  // proto: static QColor QColor::fromHsl(int h, int s, int l, int a);
extern void _ZN6QColor7fromHslEiiii(int arg0, int arg1, int arg2, int arg3);
  // proto:  void QColor::setHsl(int h, int s, int l, int a);
extern void _ZN6QColor6setHslEiiii(void* qthis, int arg0, int arg1, int arg2, int arg3);
  // proto:  void QColor::QColor(QRgb rgb);
extern void* dector_ZN6QColorC1Ej(unsigned int arg0);
extern void _ZN6QColorC1Ej(void* qthis, unsigned int arg0);
  // proto:  void QColor::setGreen(int green);
extern void _ZN6QColor8setGreenEi(void* qthis, int arg0);
  // proto:  void QColor::getHsl(int * h, int * s, int * l, int * a);
extern void _ZNK6QColor6getHslEPiS0_S0_S0_(void* qthis, int* arg0, int* arg1, int* arg2, int* arg3);
  // proto: static bool QColor::isValidColor(const QString & name);
extern void _ZN6QColor12isValidColorERK7QString(void* arg0);
  // proto:  int QColor::hslSaturation();
extern void _ZNK6QColor13hslSaturationEv(void* qthis);
  // proto: static QColor QColor::fromRgbF(qreal r, qreal g, qreal b, qreal a);
extern void _ZN6QColor8fromRgbFEdddd(double arg0, double arg1, double arg2, double arg3);
  // proto:  int QColor::blue();
extern void _ZNK6QColor4blueEv(void* qthis);
  // proto:  int QColor::hsvHue();
extern void _ZNK6QColor6hsvHueEv(void* qthis);
  // proto:  qreal QColor::valueF();
extern void _ZNK6QColor6valueFEv(void* qthis);
  // proto:  qreal QColor::cyanF();
extern void _ZNK6QColor5cyanFEv(void* qthis);
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
}

// class sizeof(QColor)=16
type QColor struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QColor::getHsvF(qreal * h, qreal * s, qreal * v, qreal * a);
func (this *QColor) getHsvF(args ...interface{}) () {
  // getHsvF(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7getHsvFEPdS0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getHsvF", args)
  }

}

  // proto:  int QColor::alpha();
func (this *QColor) alpha(args ...interface{}) () {
  // alpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5alphaEv
  default:
    qtrt.ErrorResolve("QColor", "alpha", args)
  }

}

  // proto:  qreal QColor::hslSaturationF();
func (this *QColor) hslSaturationF(args ...interface{}) () {
  // hslSaturationF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor14hslSaturationFEv
  default:
    qtrt.ErrorResolve("QColor", "hslSaturationF", args)
  }

}

  // proto:  void QColor::setAlphaF(qreal alpha);
func (this *QColor) setAlphaF(args ...interface{}) () {
  // setAlphaF(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor9setAlphaFEd
  default:
    qtrt.ErrorResolve("QColor", "setAlphaF", args)
  }

}

  // proto:  void QColor::getRgb(int * r, int * g, int * b, int * a);
func (this *QColor) getRgb(args ...interface{}) () {
  // getRgb(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6getRgbEPiS0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getRgb", args)
  }

}

  // proto:  int QColor::hslHue();
func (this *QColor) hslHue(args ...interface{}) () {
  // hslHue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6hslHueEv
  default:
    qtrt.ErrorResolve("QColor", "hslHue", args)
  }

}

  // proto:  int QColor::lightness();
func (this *QColor) lightness(args ...interface{}) () {
  // lightness()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor9lightnessEv
  default:
    qtrt.ErrorResolve("QColor", "lightness", args)
  }

}

  // proto:  void QColor::setAlpha(int alpha);
func (this *QColor) setAlpha(args ...interface{}) () {
  // setAlpha(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor8setAlphaEi
  default:
    qtrt.ErrorResolve("QColor", "setAlpha", args)
  }

}

  // proto: static QColor QColor::fromHslF(qreal h, qreal s, qreal l, qreal a);
func (this *QColor) fromHslF_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromHslF", args)
  }

}

  // proto:  void QColor::getCmyk(int * c, int * m, int * y, int * k, int * a);
func (this *QColor) getCmyk(args ...interface{}) () {
  // getCmyk(int *, int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"
  vtys[0][4] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7getCmykEPiS0_S0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getCmyk", args)
  }

}

  // proto:  int QColor::green();
func (this *QColor) green(args ...interface{}) () {
  // green()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5greenEv
  default:
    qtrt.ErrorResolve("QColor", "green", args)
  }

}

  // proto:  int QColor::hsvSaturation();
func (this *QColor) hsvSaturation(args ...interface{}) () {
  // hsvSaturation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor13hsvSaturationEv
  default:
    qtrt.ErrorResolve("QColor", "hsvSaturation", args)
  }

}

  // proto:  QColor QColor::toHsl();
func (this *QColor) toHsl(args ...interface{}) () {
  // toHsl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5toHslEv
  default:
    qtrt.ErrorResolve("QColor", "toHsl", args)
  }

}

  // proto:  void QColor::QColor();
func NewQColor(args ...interface{}) QColor {
  return QColor{}
}

  // proto:  void QColor::setBlue(int blue);
func (this *QColor) setBlue(args ...interface{}) () {
  // setBlue(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setBlueEi
  default:
    qtrt.ErrorResolve("QColor", "setBlue", args)
  }

}

  // proto:  int QColor::cyan();
func (this *QColor) cyan(args ...interface{}) () {
  // cyan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4cyanEv
  default:
    qtrt.ErrorResolve("QColor", "cyan", args)
  }

}

  // proto:  void QColor::setCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
func (this *QColor) setCmykF(args ...interface{}) () {
  // setCmykF(qreal, qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][4] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor8setCmykFEddddd
  default:
    qtrt.ErrorResolve("QColor", "setCmykF", args)
  }

}

  // proto: static QColor QColor::fromCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
func (this *QColor) fromCmykF_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromCmykF", args)
  }

}

  // proto:  QColor QColor::light(int f);
func (this *QColor) light(args ...interface{}) () {
  // light(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5lightEi
  default:
    qtrt.ErrorResolve("QColor", "light", args)
  }

}

  // proto:  void QColor::getHslF(qreal * h, qreal * s, qreal * l, qreal * a);
func (this *QColor) getHslF(args ...interface{}) () {
  // getHslF(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7getHslFEPdS0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getHslF", args)
  }

}

  // proto: static QColor QColor::fromRgb(QRgb rgb);
func (this *QColor) fromRgb_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromRgb", args)
  }

}

  // proto:  int QColor::yellow();
func (this *QColor) yellow(args ...interface{}) () {
  // yellow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6yellowEv
  default:
    qtrt.ErrorResolve("QColor", "yellow", args)
  }

}

  // proto:  void QColor::getRgbF(qreal * r, qreal * g, qreal * b, qreal * a);
func (this *QColor) getRgbF(args ...interface{}) () {
  // getRgbF(qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7getRgbFEPdS0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getRgbF", args)
  }

}

  // proto:  void QColor::setRgb(int r, int g, int b, int a);
func (this *QColor) setRgb(args ...interface{}) () {
  // setRgb(int, int, int, int)
  // setRgb(QRgb)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QRgb"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor6setRgbEiiii
  case 1:
    // invoke: _ZN6QColor6setRgbEj
  default:
    qtrt.ErrorResolve("QColor", "setRgb", args)
  }

}

  // proto:  QString QColor::name();
func (this *QColor) name(args ...interface{}) () {
  // name()
  // name(enum QColor::NameFormat)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QColor::NameFormat"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4nameEv
  case 1:
    // invoke: _ZNK6QColor4nameENS_10NameFormatE
  default:
    qtrt.ErrorResolve("QColor", "name", args)
  }

}

  // proto:  qreal QColor::redF();
func (this *QColor) redF(args ...interface{}) () {
  // redF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4redFEv
  default:
    qtrt.ErrorResolve("QColor", "redF", args)
  }

}

  // proto:  qreal QColor::blackF();
func (this *QColor) blackF(args ...interface{}) () {
  // blackF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6blackFEv
  default:
    qtrt.ErrorResolve("QColor", "blackF", args)
  }

}

  // proto:  void QColor::setHsvF(qreal h, qreal s, qreal v, qreal a);
func (this *QColor) setHsvF(args ...interface{}) () {
  // setHsvF(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setHsvFEdddd
  default:
    qtrt.ErrorResolve("QColor", "setHsvF", args)
  }

}

  // proto:  qreal QColor::hsvHueF();
func (this *QColor) hsvHueF(args ...interface{}) () {
  // hsvHueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7hsvHueFEv
  default:
    qtrt.ErrorResolve("QColor", "hsvHueF", args)
  }

}

  // proto:  qreal QColor::hsvSaturationF();
func (this *QColor) hsvSaturationF(args ...interface{}) () {
  // hsvSaturationF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor14hsvSaturationFEv
  default:
    qtrt.ErrorResolve("QColor", "hsvSaturationF", args)
  }

}

  // proto:  qreal QColor::yellowF();
func (this *QColor) yellowF(args ...interface{}) () {
  // yellowF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7yellowFEv
  default:
    qtrt.ErrorResolve("QColor", "yellowF", args)
  }

}

  // proto:  int QColor::black();
func (this *QColor) black(args ...interface{}) () {
  // black()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5blackEv
  default:
    qtrt.ErrorResolve("QColor", "black", args)
  }

}

  // proto:  void QColor::setGreenF(qreal green);
func (this *QColor) setGreenF(args ...interface{}) () {
  // setGreenF(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor9setGreenFEd
  default:
    qtrt.ErrorResolve("QColor", "setGreenF", args)
  }

}

  // proto:  QRgb QColor::rgba();
func (this *QColor) rgba(args ...interface{}) () {
  // rgba()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4rgbaEv
  default:
    qtrt.ErrorResolve("QColor", "rgba", args)
  }

}

  // proto:  QColor QColor::toCmyk();
func (this *QColor) toCmyk(args ...interface{}) () {
  // toCmyk()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6toCmykEv
  default:
    qtrt.ErrorResolve("QColor", "toCmyk", args)
  }

}

  // proto:  qreal QColor::greenF();
func (this *QColor) greenF(args ...interface{}) () {
  // greenF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6greenFEv
  default:
    qtrt.ErrorResolve("QColor", "greenF", args)
  }

}

  // proto:  int QColor::red();
func (this *QColor) red(args ...interface{}) () {
  // red()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor3redEv
  default:
    qtrt.ErrorResolve("QColor", "red", args)
  }

}

  // proto:  void QColor::setRgbF(qreal r, qreal g, qreal b, qreal a);
func (this *QColor) setRgbF(args ...interface{}) () {
  // setRgbF(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setRgbFEdddd
  default:
    qtrt.ErrorResolve("QColor", "setRgbF", args)
  }

}

  // proto:  qreal QColor::lightnessF();
func (this *QColor) lightnessF(args ...interface{}) () {
  // lightnessF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor10lightnessFEv
  default:
    qtrt.ErrorResolve("QColor", "lightnessF", args)
  }

}

  // proto:  QColor QColor::toHsv();
func (this *QColor) toHsv(args ...interface{}) () {
  // toHsv()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5toHsvEv
  default:
    qtrt.ErrorResolve("QColor", "toHsv", args)
  }

}

  // proto: static QColor QColor::fromHsv(int h, int s, int v, int a);
func (this *QColor) fromHsv_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromHsv", args)
  }

}

  // proto:  qreal QColor::hueF();
func (this *QColor) hueF(args ...interface{}) () {
  // hueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4hueFEv
  default:
    qtrt.ErrorResolve("QColor", "hueF", args)
  }

}

  // proto:  void QColor::setBlueF(qreal blue);
func (this *QColor) setBlueF(args ...interface{}) () {
  // setBlueF(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor8setBlueFEd
  default:
    qtrt.ErrorResolve("QColor", "setBlueF", args)
  }

}

  // proto:  qreal QColor::saturationF();
func (this *QColor) saturationF(args ...interface{}) () {
  // saturationF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor11saturationFEv
  default:
    qtrt.ErrorResolve("QColor", "saturationF", args)
  }

}

  // proto:  bool QColor::isValid();
func (this *QColor) isValid(args ...interface{}) () {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7isValidEv
  default:
    qtrt.ErrorResolve("QColor", "isValid", args)
  }

}

  // proto:  QColor QColor::darker(int f);
func (this *QColor) darker(args ...interface{}) () {
  // darker(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6darkerEi
  default:
    qtrt.ErrorResolve("QColor", "darker", args)
  }

}

  // proto:  qreal QColor::blueF();
func (this *QColor) blueF(args ...interface{}) () {
  // blueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5blueFEv
  default:
    qtrt.ErrorResolve("QColor", "blueF", args)
  }

}

  // proto:  int QColor::hue();
func (this *QColor) hue(args ...interface{}) () {
  // hue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor3hueEv
  default:
    qtrt.ErrorResolve("QColor", "hue", args)
  }

}

  // proto:  void QColor::setRgba(QRgb rgba);
func (this *QColor) setRgba(args ...interface{}) () {
  // setRgba(QRgb)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QRgb"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setRgbaEj
  default:
    qtrt.ErrorResolve("QColor", "setRgba", args)
  }

}

  // proto:  void QColor::setNamedColor(const QString & name);
func (this *QColor) setNamedColor(args ...interface{}) () {
  // setNamedColor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor13setNamedColorERK7QString
  default:
    qtrt.ErrorResolve("QColor", "setNamedColor", args)
  }

}

  // proto:  int QColor::magenta();
func (this *QColor) magenta(args ...interface{}) () {
  // magenta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7magentaEv
  default:
    qtrt.ErrorResolve("QColor", "magenta", args)
  }

}

  // proto:  QColor QColor::lighter(int f);
func (this *QColor) lighter(args ...interface{}) () {
  // lighter(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7lighterEi
  default:
    qtrt.ErrorResolve("QColor", "lighter", args)
  }

}

  // proto:  QColor QColor::toRgb();
func (this *QColor) toRgb(args ...interface{}) () {
  // toRgb()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5toRgbEv
  default:
    qtrt.ErrorResolve("QColor", "toRgb", args)
  }

}

  // proto:  qreal QColor::magentaF();
func (this *QColor) magentaF(args ...interface{}) () {
  // magentaF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor8magentaFEv
  default:
    qtrt.ErrorResolve("QColor", "magentaF", args)
  }

}

  // proto:  qreal QColor::hslHueF();
func (this *QColor) hslHueF(args ...interface{}) () {
  // hslHueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7hslHueFEv
  default:
    qtrt.ErrorResolve("QColor", "hslHueF", args)
  }

}

  // proto: static QColor QColor::fromCmyk(int c, int m, int y, int k, int a);
func (this *QColor) fromCmyk_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromCmyk", args)
  }

}

  // proto:  void QColor::setCmyk(int c, int m, int y, int k, int a);
func (this *QColor) setCmyk(args ...interface{}) () {
  // setCmyk(int, int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"
  vtys[0][4] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setCmykEiiiii
  default:
    qtrt.ErrorResolve("QColor", "setCmyk", args)
  }

}

  // proto: static QStringList QColor::colorNames();
func (this *QColor) colorNames_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "colorNames", args)
  }

}

  // proto:  void QColor::getHsv(int * h, int * s, int * v, int * a);
func (this *QColor) getHsv(args ...interface{}) () {
  // getHsv(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6getHsvEPiS0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getHsv", args)
  }

}

  // proto:  void QColor::getCmykF(qreal * c, qreal * m, qreal * y, qreal * k, qreal * a);
func (this *QColor) getCmykF(args ...interface{}) () {
  // getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][1] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][2] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][3] = qtrt.DoubleTy(true) // "qreal *"
  vtys[0][4] = qtrt.DoubleTy(true) // "qreal *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor8getCmykFEPdS0_S0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getCmykF", args)
  }

}

  // proto:  void QColor::setRed(int red);
func (this *QColor) setRed(args ...interface{}) () {
  // setRed(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor6setRedEi
  default:
    qtrt.ErrorResolve("QColor", "setRed", args)
  }

}

  // proto: static QColor QColor::fromRgba(QRgb rgba);
func (this *QColor) fromRgba_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromRgba", args)
  }

}

  // proto:  void QColor::setHsv(int h, int s, int v, int a);
func (this *QColor) setHsv(args ...interface{}) () {
  // setHsv(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor6setHsvEiiii
  default:
    qtrt.ErrorResolve("QColor", "setHsv", args)
  }

}

  // proto:  QRgb QColor::rgb();
func (this *QColor) rgb(args ...interface{}) () {
  // rgb()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor3rgbEv
  default:
    qtrt.ErrorResolve("QColor", "rgb", args)
  }

}

  // proto:  void QColor::setHslF(qreal h, qreal s, qreal l, qreal a);
func (this *QColor) setHslF(args ...interface{}) () {
  // setHslF(qreal, qreal, qreal, qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][1] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][2] = qtrt.DoubleTy(false) // "qreal"
  vtys[0][3] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setHslFEdddd
  default:
    qtrt.ErrorResolve("QColor", "setHslF", args)
  }

}

  // proto:  int QColor::saturation();
func (this *QColor) saturation(args ...interface{}) () {
  // saturation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor10saturationEv
  default:
    qtrt.ErrorResolve("QColor", "saturation", args)
  }

}

  // proto:  qreal QColor::alphaF();
func (this *QColor) alphaF(args ...interface{}) () {
  // alphaF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6alphaFEv
  default:
    qtrt.ErrorResolve("QColor", "alphaF", args)
  }

}

  // proto:  int QColor::value();
func (this *QColor) value(args ...interface{}) () {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5valueEv
  default:
    qtrt.ErrorResolve("QColor", "value", args)
  }

}

  // proto: static QColor QColor::fromHsvF(qreal h, qreal s, qreal v, qreal a);
func (this *QColor) fromHsvF_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromHsvF", args)
  }

}

  // proto:  QColor QColor::dark(int f);
func (this *QColor) dark(args ...interface{}) () {
  // dark(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4darkEi
  default:
    qtrt.ErrorResolve("QColor", "dark", args)
  }

}

  // proto:  void QColor::setRedF(qreal red);
func (this *QColor) setRedF(args ...interface{}) () {
  // setRedF(qreal)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.DoubleTy(false) // "qreal"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7setRedFEd
  default:
    qtrt.ErrorResolve("QColor", "setRedF", args)
  }

}

  // proto: static QColor QColor::fromHsl(int h, int s, int l, int a);
func (this *QColor) fromHsl_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromHsl", args)
  }

}

  // proto:  void QColor::setHsl(int h, int s, int l, int a);
func (this *QColor) setHsl(args ...interface{}) () {
  // setHsl(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[0][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor6setHslEiiii
  default:
    qtrt.ErrorResolve("QColor", "setHsl", args)
  }

}

  // proto:  void QColor::setGreen(int green);
func (this *QColor) setGreen(args ...interface{}) () {
  // setGreen(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor8setGreenEi
  default:
    qtrt.ErrorResolve("QColor", "setGreen", args)
  }

}

  // proto:  void QColor::getHsl(int * h, int * s, int * l, int * a);
func (this *QColor) getHsl(args ...interface{}) () {
  // getHsl(int *, int *, int *, int *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "int *"
  vtys[0][1] = qtrt.Int32Ty(true) // "int *"
  vtys[0][2] = qtrt.Int32Ty(true) // "int *"
  vtys[0][3] = qtrt.Int32Ty(true) // "int *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6getHslEPiS0_S0_S0_
  default:
    qtrt.ErrorResolve("QColor", "getHsl", args)
  }

}

  // proto: static bool QColor::isValidColor(const QString & name);
func (this *QColor) isValidColor_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "isValidColor", args)
  }

}

  // proto:  int QColor::hslSaturation();
func (this *QColor) hslSaturation(args ...interface{}) () {
  // hslSaturation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor13hslSaturationEv
  default:
    qtrt.ErrorResolve("QColor", "hslSaturation", args)
  }

}

  // proto: static QColor QColor::fromRgbF(qreal r, qreal g, qreal b, qreal a);
func (this *QColor) fromRgbF_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QColor", "fromRgbF", args)
  }

}

  // proto:  int QColor::blue();
func (this *QColor) blue(args ...interface{}) () {
  // blue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4blueEv
  default:
    qtrt.ErrorResolve("QColor", "blue", args)
  }

}

  // proto:  int QColor::hsvHue();
func (this *QColor) hsvHue(args ...interface{}) () {
  // hsvHue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6hsvHueEv
  default:
    qtrt.ErrorResolve("QColor", "hsvHue", args)
  }

}

  // proto:  qreal QColor::valueF();
func (this *QColor) valueF(args ...interface{}) () {
  // valueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6valueFEv
  default:
    qtrt.ErrorResolve("QColor", "valueF", args)
  }

}

  // proto:  qreal QColor::cyanF();
func (this *QColor) cyanF(args ...interface{}) () {
  // cyanF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5cyanFEv
  default:
    qtrt.ErrorResolve("QColor", "cyanF", args)
  }

}

// <= body block end

