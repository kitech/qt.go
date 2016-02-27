package qtgui
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
import "qtrt"
import "qtcore"
// <= use block end

// ext block begin =>

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QColor::hueF();
extern double C_ZNK6QColor4hueFEv(void* qthis); // 4
  // proto:  int QColor::yellow();
extern int32_t C_ZNK6QColor6yellowEv(void* qthis); // 4
  // proto:  qreal QColor::cyanF();
extern double C_ZNK6QColor5cyanFEv(void* qthis); // 4
  // proto:  qreal QColor::magentaF();
extern double C_ZNK6QColor8magentaFEv(void* qthis); // 4
  // proto: static bool QColor::isValidColor(const QString & name);
extern bool C_ZN6QColor12isValidColorERK7QString(void* arg0); // 4
  // proto:  void QColor::setHsvF(qreal h, qreal s, qreal v, qreal a);
extern void C_ZN6QColor7setHsvFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QColor::setHsl(int h, int s, int l, int a);
extern void C_ZN6QColor6setHslEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QColor::setCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
extern void C_ZN6QColor8setCmykFEddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4); // 4
  // proto:  void QColor::setHsv(int h, int s, int v, int a);
extern void C_ZN6QColor6setHsvEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QRgb QColor::rgb();
extern int32_t C_ZNK6QColor3rgbEv(void* qthis); // 4
  // proto:  int QColor::black();
extern int32_t C_ZNK6QColor5blackEv(void* qthis); // 4
  // proto: static QColor QColor::fromCmyk(int c, int m, int y, int k, int a);
extern void* C_ZN6QColor8fromCmykEiiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  QColor::Spec QColor::spec();
extern void C_ZNK6QColor4specEv(void* qthis); // 2
  // proto:  int QColor::saturation();
extern int32_t C_ZNK6QColor10saturationEv(void* qthis); // 4
  // proto:  void QColor::setCmyk(int c, int m, int y, int k, int a);
extern void C_ZN6QColor7setCmykEiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  QColor QColor::lighter(int f);
extern void* C_ZNK6QColor7lighterEi(void* qthis, int32_t arg0); // 2
  // proto:  qreal QColor::yellowF();
extern double C_ZNK6QColor7yellowFEv(void* qthis); // 4
  // proto: static QColor QColor::fromRgbF(qreal r, qreal g, qreal b, qreal a);
extern void* C_ZN6QColor8fromRgbFEdddd(double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QColor::getHsvF(qreal * h, qreal * s, qreal * v, qreal * a);
extern void C_ZNK6QColor7getHsvFEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  QString QColor::name();
extern void* C_ZNK6QColor4nameEv(void* qthis); // 4
  // proto:  QColor QColor::toRgb();
extern void* C_ZNK6QColor5toRgbEv(void* qthis); // 4
  // proto:  qreal QColor::hsvHueF();
extern double C_ZNK6QColor7hsvHueFEv(void* qthis); // 4
  // proto: static QColor QColor::fromRgba(QRgb rgba);
extern void* C_ZN6QColor8fromRgbaEj(int32_t arg0); // 4
  // proto:  void QColor::setAlphaF(qreal alpha);
extern void C_ZN6QColor9setAlphaFEd(void* qthis, double arg0); // 4
  // proto:  int QColor::hslSaturation();
extern int32_t C_ZNK6QColor13hslSaturationEv(void* qthis); // 4
  // proto:  qreal QColor::blackF();
extern double C_ZNK6QColor6blackFEv(void* qthis); // 4
  // proto:  qreal QColor::saturationF();
extern double C_ZNK6QColor11saturationFEv(void* qthis); // 4
  // proto:  void QColor::setRgb(QRgb rgb);
extern void C_ZN6QColor6setRgbEj(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::setRgb(int r, int g, int b, int a);
extern void C_ZN6QColor6setRgbEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QColor::blue();
extern int32_t C_ZNK6QColor4blueEv(void* qthis); // 4
  // proto:  qreal QColor::blueF();
extern double C_ZNK6QColor5blueFEv(void* qthis); // 4
  // proto:  QRgb QColor::rgba();
extern int32_t C_ZNK6QColor4rgbaEv(void* qthis); // 4
  // proto:  int QColor::red();
extern int32_t C_ZNK6QColor3redEv(void* qthis); // 4
  // proto: static QColor QColor::fromHsvF(qreal h, qreal s, qreal v, qreal a);
extern void* C_ZN6QColor8fromHsvFEdddd(double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  qreal QColor::alphaF();
extern double C_ZNK6QColor6alphaFEv(void* qthis); // 4
  // proto:  void QColor::setRedF(qreal red);
extern void C_ZN6QColor7setRedFEd(void* qthis, double arg0); // 4
  // proto:  int QColor::lightness();
extern int32_t C_ZNK6QColor9lightnessEv(void* qthis); // 4
  // proto: static QStringList QColor::colorNames();
extern void C_ZN6QColor10colorNamesEv(); // 4
  // proto:  void QColor::getCmykF(qreal * c, qreal * m, qreal * y, qreal * k, qreal * a);
extern void C_ZN6QColor8getCmykFEPdS0_S0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 4
  // proto:  void QColor::setNamedColor(const QString & name);
extern void C_ZN6QColor13setNamedColorERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QColor::setGreen(int green);
extern void C_ZN6QColor8setGreenEi(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::setRed(int red);
extern void C_ZN6QColor6setRedEi(void* qthis, int32_t arg0); // 4
  // proto:  QColor QColor::toCmyk();
extern void* C_ZNK6QColor6toCmykEv(void* qthis); // 4
  // proto:  qreal QColor::redF();
extern double C_ZNK6QColor4redFEv(void* qthis); // 4
  // proto:  int QColor::hue();
extern int32_t C_ZNK6QColor3hueEv(void* qthis); // 4
  // proto:  qreal QColor::hslSaturationF();
extern double C_ZNK6QColor14hslSaturationFEv(void* qthis); // 4
  // proto:  qreal QColor::hsvSaturationF();
extern double C_ZNK6QColor14hsvSaturationFEv(void* qthis); // 4
  // proto:  QColor QColor::toHsl();
extern void* C_ZNK6QColor5toHslEv(void* qthis); // 4
  // proto:  int QColor::value();
extern int32_t C_ZNK6QColor5valueEv(void* qthis); // 4
  // proto:  QColor QColor::toHsv();
extern void* C_ZNK6QColor5toHsvEv(void* qthis); // 4
  // proto:  QColor QColor::darker(int f);
extern void* C_ZNK6QColor6darkerEi(void* qthis, int32_t arg0); // 2
  // proto:  void QColor::setBlue(int blue);
extern void C_ZN6QColor7setBlueEi(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::setHslF(qreal h, qreal s, qreal l, qreal a);
extern void C_ZN6QColor7setHslFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  qreal QColor::valueF();
extern double C_ZNK6QColor6valueFEv(void* qthis); // 4
  // proto: static QColor QColor::fromHsl(int h, int s, int l, int a);
extern void* C_ZN6QColor7fromHslEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QColor::hslHue();
extern int32_t C_ZNK6QColor6hslHueEv(void* qthis); // 4
  // proto: static QColor QColor::fromHsv(int h, int s, int v, int a);
extern void* C_ZN6QColor7fromHsvEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QColor::getRgb(int * r, int * g, int * b, int * a);
extern void C_ZNK6QColor6getRgbEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  void QColor::setRgba(QRgb rgba);
extern void C_ZN6QColor7setRgbaEj(void* qthis, int32_t arg0); // 4
  // proto:  int QColor::magenta();
extern int32_t C_ZNK6QColor7magentaEv(void* qthis); // 4
  // proto:  void QColor::QColor(const char * name);
extern void* C_ZN6QColorC2EPKc(void* arg0); // 1
  // proto:  void QColor::QColor(QRgb rgb);
extern void* C_ZN6QColorC2Ej(int32_t arg0); // 3
  // proto:  void QColor::QColor(const QColor & color);
extern void* C_ZN6QColorC2ERKS_(void* arg0); // 1
  // proto:  void QColor::QColor(int r, int g, int b, int a);
extern void* C_ZN6QColorC2Eiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 1
  // proto:  void QColor::QColor();
extern void* C_ZN6QColorC2Ev(); // 1
  // proto:  void QColor::QColor(const QString & name);
extern void* C_ZN6QColorC2ERK7QString(void* arg0); // 1
  // proto:  qreal QColor::hslHueF();
extern double C_ZNK6QColor7hslHueFEv(void* qthis); // 4
  // proto: static QColor QColor::fromHslF(qreal h, qreal s, qreal l, qreal a);
extern void* C_ZN6QColor8fromHslFEdddd(double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  int QColor::hsvSaturation();
extern int32_t C_ZNK6QColor13hsvSaturationEv(void* qthis); // 4
  // proto:  void QColor::setRgbF(qreal r, qreal g, qreal b, qreal a);
extern void C_ZN6QColor7setRgbFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QColor::setAlpha(int alpha);
extern void C_ZN6QColor8setAlphaEi(void* qthis, int32_t arg0); // 4
  // proto:  int QColor::cyan();
extern int32_t C_ZNK6QColor4cyanEv(void* qthis); // 4
  // proto:  void QColor::getRgbF(qreal * r, qreal * g, qreal * b, qreal * a);
extern void C_ZNK6QColor7getRgbFEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  void QColor::setBlueF(qreal blue);
extern void C_ZN6QColor8setBlueFEd(void* qthis, double arg0); // 4
  // proto:  void QColor::getCmyk(int * c, int * m, int * y, int * k, int * a);
extern void C_ZN6QColor7getCmykEPiS0_S0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 4
  // proto:  bool QColor::isValid();
extern bool C_ZNK6QColor7isValidEv(void* qthis); // 2
  // proto:  void QColor::getHsv(int * h, int * s, int * v, int * a);
extern void C_ZNK6QColor6getHsvEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto: static QColor QColor::fromCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
extern void* C_ZN6QColor9fromCmykFEddddd(double arg0, double arg1, double arg2, double arg3, double arg4); // 4
  // proto:  int QColor::hsvHue();
extern int32_t C_ZNK6QColor6hsvHueEv(void* qthis); // 4
  // proto:  void QColor::getHsl(int * h, int * s, int * l, int * a);
extern void C_ZNK6QColor6getHslEPiS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  int QColor::alpha();
extern int32_t C_ZNK6QColor5alphaEv(void* qthis); // 4
  // proto:  QColor QColor::dark(int f);
extern void* C_ZNK6QColor4darkEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QColor::greenF();
extern double C_ZNK6QColor6greenFEv(void* qthis); // 4
  // proto:  QColor QColor::light(int f);
extern void* C_ZNK6QColor5lightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::getHslF(qreal * h, qreal * s, qreal * l, qreal * a);
extern void C_ZNK6QColor7getHslFEPdS0_S0_S0_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 4
  // proto:  qreal QColor::lightnessF();
extern double C_ZNK6QColor10lightnessFEv(void* qthis); // 4
  // proto:  int QColor::green();
extern int32_t C_ZNK6QColor5greenEv(void* qthis); // 4
  // proto: static QColor QColor::fromRgb(QRgb rgb);
extern void* C_ZN6QColor7fromRgbEj(int32_t arg0); // 4
  // proto: static QColor QColor::fromRgb(int r, int g, int b, int a);
extern void* C_ZN6QColor7fromRgbEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QColor::setGreenF(qreal green);
extern void C_ZN6QColor9setGreenFEd(void* qthis, double arg0); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {qtcore.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QColor)=16
type QColor struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// hueF()
func (this *QColor) HueF(args ...interface{}) (ret interface{}) {
  // hueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4hueFEv
    // invoke: qreal hueF()
    var ret0 = C.C_ZNK6QColor4hueFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hueF", args)
  }

  return
}

// yellow()
func (this *QColor) Yellow(args ...interface{}) (ret interface{}) {
  // yellow()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6yellowEv
    // invoke: int yellow()
    var ret0 = C.C_ZNK6QColor6yellowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "yellow", args)
  }

  return
}

// cyanF()
func (this *QColor) CyanF(args ...interface{}) (ret interface{}) {
  // cyanF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5cyanFEv
    // invoke: qreal cyanF()
    var ret0 = C.C_ZNK6QColor5cyanFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "cyanF", args)
  }

  return
}

// magentaF()
func (this *QColor) MagentaF(args ...interface{}) (ret interface{}) {
  // magentaF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor8magentaFEv
    // invoke: qreal magentaF()
    var ret0 = C.C_ZNK6QColor8magentaFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "magentaF", args)
  }

  return
}

// isValidColor(const class QString &)
func (this *QColor) IsValidColor_s(args ...interface{}) (ret interface{}) {
  // isValidColor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor12isValidColorERK7QString
    // invoke: bool isValidColor(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QColor12isValidColorERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "isValidColor", args)
  }

  return
}

// setHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) SetHsvF(args ...interface{}) () {
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
    // invoke: void setHsvF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor7setHsvFEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHsvF", args)
  }

  return
}

// setHsl(int, int, int, int)
func (this *QColor) SetHsl(args ...interface{}) () {
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
    // invoke: void setHsl(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor6setHslEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHsl", args)
  }

  return
}

// setCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) SetCmykF(args ...interface{}) () {
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
    // invoke: void setCmykF(qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor8setCmykFEddddd(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "setCmykF", args)
  }

  return
}

// setHsv(int, int, int, int)
func (this *QColor) SetHsv(args ...interface{}) () {
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
    // invoke: void setHsv(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor6setHsvEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHsv", args)
  }

  return
}

// rgb()
func (this *QColor) Rgb(args ...interface{}) (ret interface{}) {
  // rgb()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor3rgbEv
    // invoke: QRgb rgb()
    var ret0 = C.C_ZNK6QColor3rgbEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "QRgb"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "rgb", args)
  }

  return
}

// black()
func (this *QColor) Black(args ...interface{}) (ret interface{}) {
  // black()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5blackEv
    // invoke: int black()
    var ret0 = C.C_ZNK6QColor5blackEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "black", args)
  }

  return
}

// fromCmyk(int, int, int, int, int)
func (this *QColor) FromCmyk_s(args ...interface{}) (ret interface{}) {
  // fromCmyk(int, int, int, int, int)
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
    // invoke: _ZN6QColor8fromCmykEiiiii
    // invoke: QColor fromCmyk(int, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN6QColor8fromCmykEiiiii(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromCmyk", args)
  }

  return
}

// spec()
func (this *QColor) Spec(args ...interface{}) () {
  // spec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4specEv
    // invoke: QColor::Spec spec()
    C.C_ZNK6QColor4specEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QColor", "spec", args)
  }

  return
}

// saturation()
func (this *QColor) Saturation(args ...interface{}) (ret interface{}) {
  // saturation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor10saturationEv
    // invoke: int saturation()
    var ret0 = C.C_ZNK6QColor10saturationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "saturation", args)
  }

  return
}

// setCmyk(int, int, int, int, int)
func (this *QColor) SetCmyk(args ...interface{}) () {
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
    // invoke: void setCmyk(int, int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(qtrt.PrimConv(args[4], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor7setCmykEiiiii(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "setCmyk", args)
  }

  return
}

// lighter(int)
func (this *QColor) Lighter(args ...interface{}) (ret interface{}) {
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
    // invoke: QColor lighter(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QColor7lighterEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "lighter", args)
  }

  return
}

// yellowF()
func (this *QColor) YellowF(args ...interface{}) (ret interface{}) {
  // yellowF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7yellowFEv
    // invoke: qreal yellowF()
    var ret0 = C.C_ZNK6QColor7yellowFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "yellowF", args)
  }

  return
}

// fromRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) FromRgbF_s(args ...interface{}) (ret interface{}) {
  // fromRgbF(qreal, qreal, qreal, qreal)
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
    // invoke: _ZN6QColor8fromRgbFEdddd
    // invoke: QColor fromRgbF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN6QColor8fromRgbFEdddd(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromRgbF", args)
  }

  return
}

// getHsvF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetHsvF(args ...interface{}) () {
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
    // invoke: void getHsvF(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor7getHsvFEPdS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHsvF", args)
  }

  return
}

// name()
func (this *QColor) Name(args ...interface{}) (ret interface{}) {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4nameEv
    // invoke: QString name()
    var ret0 = C.C_ZNK6QColor4nameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(qtcore.QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "name", args)
  }

  return
}

// toRgb()
func (this *QColor) ToRgb(args ...interface{}) (ret interface{}) {
  // toRgb()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5toRgbEv
    // invoke: QColor toRgb()
    var ret0 = C.C_ZNK6QColor5toRgbEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "toRgb", args)
  }

  return
}

// hsvHueF()
func (this *QColor) HsvHueF(args ...interface{}) (ret interface{}) {
  // hsvHueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7hsvHueFEv
    // invoke: qreal hsvHueF()
    var ret0 = C.C_ZNK6QColor7hsvHueFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hsvHueF", args)
  }

  return
}

// fromRgba(QRgb)
func (this *QColor) FromRgba_s(args ...interface{}) (ret interface{}) {
  // fromRgba(QRgb)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QRgb"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor8fromRgbaEj
    // invoke: QColor fromRgba(QRgb)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QColor8fromRgbaEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromRgba", args)
  }

  return
}

// setAlphaF(qreal)
func (this *QColor) SetAlphaF(args ...interface{}) () {
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
    // invoke: void setAlphaF(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor9setAlphaFEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setAlphaF", args)
  }

  return
}

// hslSaturation()
func (this *QColor) HslSaturation(args ...interface{}) (ret interface{}) {
  // hslSaturation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor13hslSaturationEv
    // invoke: int hslSaturation()
    var ret0 = C.C_ZNK6QColor13hslSaturationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hslSaturation", args)
  }

  return
}

// blackF()
func (this *QColor) BlackF(args ...interface{}) (ret interface{}) {
  // blackF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6blackFEv
    // invoke: qreal blackF()
    var ret0 = C.C_ZNK6QColor6blackFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "blackF", args)
  }

  return
}

// saturationF()
func (this *QColor) SaturationF(args ...interface{}) (ret interface{}) {
  // saturationF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor11saturationFEv
    // invoke: qreal saturationF()
    var ret0 = C.C_ZNK6QColor11saturationFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "saturationF", args)
  }

  return
}

// setRgb(QRgb)
func (this *QColor) SetRgb(args ...interface{}) () {
  // setRgb(QRgb)
  // setRgb(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QRgb"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor6setRgbEj
    // invoke: void setRgb(QRgb)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor6setRgbEj(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN6QColor6setRgbEiiii
    // invoke: void setRgb(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor6setRgbEiiii(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setRgb", args)
  }

  return
}

// blue()
func (this *QColor) Blue(args ...interface{}) (ret interface{}) {
  // blue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4blueEv
    // invoke: int blue()
    var ret0 = C.C_ZNK6QColor4blueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "blue", args)
  }

  return
}

// blueF()
func (this *QColor) BlueF(args ...interface{}) (ret interface{}) {
  // blueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5blueFEv
    // invoke: qreal blueF()
    var ret0 = C.C_ZNK6QColor5blueFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "blueF", args)
  }

  return
}

// rgba()
func (this *QColor) Rgba(args ...interface{}) (ret interface{}) {
  // rgba()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4rgbaEv
    // invoke: QRgb rgba()
    var ret0 = C.C_ZNK6QColor4rgbaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "QRgb"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "rgba", args)
  }

  return
}

// red()
func (this *QColor) Red(args ...interface{}) (ret interface{}) {
  // red()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor3redEv
    // invoke: int red()
    var ret0 = C.C_ZNK6QColor3redEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "red", args)
  }

  return
}

// fromHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) FromHsvF_s(args ...interface{}) (ret interface{}) {
  // fromHsvF(qreal, qreal, qreal, qreal)
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
    // invoke: _ZN6QColor8fromHsvFEdddd
    // invoke: QColor fromHsvF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN6QColor8fromHsvFEdddd(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromHsvF", args)
  }

  return
}

// alphaF()
func (this *QColor) AlphaF(args ...interface{}) (ret interface{}) {
  // alphaF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6alphaFEv
    // invoke: qreal alphaF()
    var ret0 = C.C_ZNK6QColor6alphaFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "alphaF", args)
  }

  return
}

// setRedF(qreal)
func (this *QColor) SetRedF(args ...interface{}) () {
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
    // invoke: void setRedF(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor7setRedFEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setRedF", args)
  }

  return
}

// lightness()
func (this *QColor) Lightness(args ...interface{}) (ret interface{}) {
  // lightness()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor9lightnessEv
    // invoke: int lightness()
    var ret0 = C.C_ZNK6QColor9lightnessEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "lightness", args)
  }

  return
}

// colorNames()
func (this *QColor) ColorNames_s(args ...interface{}) () {
  // colorNames()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor10colorNamesEv
    // invoke: QStringList colorNames()
    C.C_ZN6QColor10colorNamesEv()
  default:
    qtrt.ErrorResolve("QColor", "colorNames", args)
  }

  return
}

// getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetCmykF(args ...interface{}) () {
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
    // invoke: void getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    var arg4 = (unsafe.Pointer)(args[4].(*float64))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor8getCmykFEPdS0_S0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "getCmykF", args)
  }

  return
}

// setNamedColor(const class QString &)
func (this *QColor) SetNamedColor(args ...interface{}) () {
  // setNamedColor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor13setNamedColorERK7QString
    // invoke: void setNamedColor(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor13setNamedColorERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setNamedColor", args)
  }

  return
}

// setGreen(int)
func (this *QColor) SetGreen(args ...interface{}) () {
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
    // invoke: void setGreen(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor8setGreenEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setGreen", args)
  }

  return
}

// setRed(int)
func (this *QColor) SetRed(args ...interface{}) () {
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
    // invoke: void setRed(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor6setRedEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setRed", args)
  }

  return
}

// toCmyk()
func (this *QColor) ToCmyk(args ...interface{}) (ret interface{}) {
  // toCmyk()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6toCmykEv
    // invoke: QColor toCmyk()
    var ret0 = C.C_ZNK6QColor6toCmykEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "toCmyk", args)
  }

  return
}

// redF()
func (this *QColor) RedF(args ...interface{}) (ret interface{}) {
  // redF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4redFEv
    // invoke: qreal redF()
    var ret0 = C.C_ZNK6QColor4redFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "redF", args)
  }

  return
}

// hue()
func (this *QColor) Hue(args ...interface{}) (ret interface{}) {
  // hue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor3hueEv
    // invoke: int hue()
    var ret0 = C.C_ZNK6QColor3hueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hue", args)
  }

  return
}

// hslSaturationF()
func (this *QColor) HslSaturationF(args ...interface{}) (ret interface{}) {
  // hslSaturationF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor14hslSaturationFEv
    // invoke: qreal hslSaturationF()
    var ret0 = C.C_ZNK6QColor14hslSaturationFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hslSaturationF", args)
  }

  return
}

// hsvSaturationF()
func (this *QColor) HsvSaturationF(args ...interface{}) (ret interface{}) {
  // hsvSaturationF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor14hsvSaturationFEv
    // invoke: qreal hsvSaturationF()
    var ret0 = C.C_ZNK6QColor14hsvSaturationFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hsvSaturationF", args)
  }

  return
}

// toHsl()
func (this *QColor) ToHsl(args ...interface{}) (ret interface{}) {
  // toHsl()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5toHslEv
    // invoke: QColor toHsl()
    var ret0 = C.C_ZNK6QColor5toHslEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "toHsl", args)
  }

  return
}

// value()
func (this *QColor) Value(args ...interface{}) (ret interface{}) {
  // value()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5valueEv
    // invoke: int value()
    var ret0 = C.C_ZNK6QColor5valueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "value", args)
  }

  return
}

// toHsv()
func (this *QColor) ToHsv(args ...interface{}) (ret interface{}) {
  // toHsv()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5toHsvEv
    // invoke: QColor toHsv()
    var ret0 = C.C_ZNK6QColor5toHsvEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "toHsv", args)
  }

  return
}

// darker(int)
func (this *QColor) Darker(args ...interface{}) (ret interface{}) {
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
    // invoke: QColor darker(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QColor6darkerEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "darker", args)
  }

  return
}

// setBlue(int)
func (this *QColor) SetBlue(args ...interface{}) () {
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
    // invoke: void setBlue(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor7setBlueEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setBlue", args)
  }

  return
}

// setHslF(qreal, qreal, qreal, qreal)
func (this *QColor) SetHslF(args ...interface{}) () {
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
    // invoke: void setHslF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor7setHslFEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHslF", args)
  }

  return
}

// valueF()
func (this *QColor) ValueF(args ...interface{}) (ret interface{}) {
  // valueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6valueFEv
    // invoke: qreal valueF()
    var ret0 = C.C_ZNK6QColor6valueFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "valueF", args)
  }

  return
}

// fromHsl(int, int, int, int)
func (this *QColor) FromHsl_s(args ...interface{}) (ret interface{}) {
  // fromHsl(int, int, int, int)
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
    // invoke: _ZN6QColor7fromHslEiiii
    // invoke: QColor fromHsl(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN6QColor7fromHslEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromHsl", args)
  }

  return
}

// hslHue()
func (this *QColor) HslHue(args ...interface{}) (ret interface{}) {
  // hslHue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6hslHueEv
    // invoke: int hslHue()
    var ret0 = C.C_ZNK6QColor6hslHueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hslHue", args)
  }

  return
}

// fromHsv(int, int, int, int)
func (this *QColor) FromHsv_s(args ...interface{}) (ret interface{}) {
  // fromHsv(int, int, int, int)
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
    // invoke: _ZN6QColor7fromHsvEiiii
    // invoke: QColor fromHsv(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN6QColor7fromHsvEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromHsv", args)
  }

  return
}

// getRgb(int *, int *, int *, int *)
func (this *QColor) GetRgb(args ...interface{}) () {
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
    // invoke: void getRgb(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor6getRgbEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getRgb", args)
  }

  return
}

// setRgba(QRgb)
func (this *QColor) SetRgba(args ...interface{}) () {
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
    // invoke: void setRgba(QRgb)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor7setRgbaEj(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setRgba", args)
  }

  return
}

// magenta()
func (this *QColor) Magenta(args ...interface{}) (ret interface{}) {
  // magenta()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7magentaEv
    // invoke: int magenta()
    var ret0 = C.C_ZNK6QColor7magentaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "magenta", args)
  }

  return
}

// QColor(const char *)
func GcfreeQColor(this *QColor) {
  qtrt.UniverseFree(this)
}
func NewQColor(args ...interface{}) *QColor {
  // QColor(const char *)
  // QColor(QRgb)
  // QColor(const class QColor &)
  // QColor(int, int, int, int)
  // QColor()
  // QColor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "QRgb"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QColor{}) // "const QColor &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(qtcore.QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColorC1EPKc
    // invoke: void QColor(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2EPKc(arg0)
    this := &QColor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColor)
    return this
  case 1:
    // invoke: _ZN6QColorC1Ej
    // invoke: void QColor(QRgb)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2Ej(arg0)
    this := &QColor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColor)
    return this
  case 2:
    // invoke: _ZN6QColorC1ERKS_
    // invoke: void QColor(const class QColor &)
    var arg0 = args[0].(*QColor).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2ERKS_(arg0)
    this := &QColor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColor)
    return this
  case 3:
    // invoke: _ZN6QColorC1Eiiii
    // invoke: void QColor(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2Eiiii(arg0, arg1, arg2, arg3)
    this := &QColor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColor)
    return this
  case 4:
    // invoke: _ZN6QColorC1Ev
    // invoke: void QColor()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2Ev()
    this := &QColor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColor)
    return this
  case 5:
    // invoke: _ZN6QColorC1ERK7QString
    // invoke: void QColor(const class QString &)
    var arg0 = args[0].(*qtcore.QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2ERK7QString(arg0)
    this := &QColor{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQColor)
    return this
  default:
    qtrt.ErrorResolve("QColor", "QColor", args)
  }

  return nil // QColor{Qclsinst:qthis}
}

// hslHueF()
func (this *QColor) HslHueF(args ...interface{}) (ret interface{}) {
  // hslHueF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7hslHueFEv
    // invoke: qreal hslHueF()
    var ret0 = C.C_ZNK6QColor7hslHueFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hslHueF", args)
  }

  return
}

// fromHslF(qreal, qreal, qreal, qreal)
func (this *QColor) FromHslF_s(args ...interface{}) (ret interface{}) {
  // fromHslF(qreal, qreal, qreal, qreal)
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
    // invoke: _ZN6QColor8fromHslFEdddd
    // invoke: QColor fromHslF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN6QColor8fromHslFEdddd(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromHslF", args)
  }

  return
}

// hsvSaturation()
func (this *QColor) HsvSaturation(args ...interface{}) (ret interface{}) {
  // hsvSaturation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor13hsvSaturationEv
    // invoke: int hsvSaturation()
    var ret0 = C.C_ZNK6QColor13hsvSaturationEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hsvSaturation", args)
  }

  return
}

// setRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) SetRgbF(args ...interface{}) () {
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
    // invoke: void setRgbF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor7setRgbFEdddd(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setRgbF", args)
  }

  return
}

// setAlpha(int)
func (this *QColor) SetAlpha(args ...interface{}) () {
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
    // invoke: void setAlpha(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor8setAlphaEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setAlpha", args)
  }

  return
}

// cyan()
func (this *QColor) Cyan(args ...interface{}) (ret interface{}) {
  // cyan()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor4cyanEv
    // invoke: int cyan()
    var ret0 = C.C_ZNK6QColor4cyanEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "cyan", args)
  }

  return
}

// getRgbF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetRgbF(args ...interface{}) () {
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
    // invoke: void getRgbF(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor7getRgbFEPdS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getRgbF", args)
  }

  return
}

// setBlueF(qreal)
func (this *QColor) SetBlueF(args ...interface{}) () {
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
    // invoke: void setBlueF(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor8setBlueFEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setBlueF", args)
  }

  return
}

// getCmyk(int *, int *, int *, int *, int *)
func (this *QColor) GetCmyk(args ...interface{}) () {
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
    // invoke: void getCmyk(int *, int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var arg4 = (unsafe.Pointer)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor7getCmykEPiS0_S0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "getCmyk", args)
  }

  return
}

// isValid()
func (this *QColor) IsValid(args ...interface{}) (ret interface{}) {
  // isValid()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor7isValidEv
    // invoke: bool isValid()
    var ret0 = C.C_ZNK6QColor7isValidEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "isValid", args)
  }

  return
}

// getHsv(int *, int *, int *, int *)
func (this *QColor) GetHsv(args ...interface{}) () {
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
    // invoke: void getHsv(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor6getHsvEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHsv", args)
  }

  return
}

// fromCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) FromCmykF_s(args ...interface{}) (ret interface{}) {
  // fromCmykF(qreal, qreal, qreal, qreal, qreal)
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
    // invoke: _ZN6QColor9fromCmykFEddddd
    // invoke: QColor fromCmykF(qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(qtrt.PrimConv(args[1], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(qtrt.PrimConv(args[2], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(qtrt.PrimConv(args[3], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(qtrt.PrimConv(args[4], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZN6QColor9fromCmykFEddddd(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromCmykF", args)
  }

  return
}

// hsvHue()
func (this *QColor) HsvHue(args ...interface{}) (ret interface{}) {
  // hsvHue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6hsvHueEv
    // invoke: int hsvHue()
    var ret0 = C.C_ZNK6QColor6hsvHueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "hsvHue", args)
  }

  return
}

// getHsl(int *, int *, int *, int *)
func (this *QColor) GetHsl(args ...interface{}) () {
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
    // invoke: void getHsl(int *, int *, int *, int *)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor6getHslEPiS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHsl", args)
  }

  return
}

// alpha()
func (this *QColor) Alpha(args ...interface{}) (ret interface{}) {
  // alpha()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5alphaEv
    // invoke: int alpha()
    var ret0 = C.C_ZNK6QColor5alphaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "alpha", args)
  }

  return
}

// dark(int)
func (this *QColor) Dark(args ...interface{}) (ret interface{}) {
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
    // invoke: QColor dark(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QColor4darkEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "dark", args)
  }

  return
}

// greenF()
func (this *QColor) GreenF(args ...interface{}) (ret interface{}) {
  // greenF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor6greenFEv
    // invoke: qreal greenF()
    var ret0 = C.C_ZNK6QColor6greenFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "greenF", args)
  }

  return
}

// light(int)
func (this *QColor) Light(args ...interface{}) (ret interface{}) {
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
    // invoke: QColor light(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK6QColor5lightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "light", args)
  }

  return
}

// getHslF(qreal *, qreal *, qreal *, qreal *)
func (this *QColor) GetHslF(args ...interface{}) () {
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
    // invoke: void getHslF(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (unsafe.Pointer)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (unsafe.Pointer)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (unsafe.Pointer)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (unsafe.Pointer)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor7getHslFEPdS0_S0_S0_(this.Qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHslF", args)
  }

  return
}

// lightnessF()
func (this *QColor) LightnessF(args ...interface{}) (ret interface{}) {
  // lightnessF()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor10lightnessFEv
    // invoke: qreal lightnessF()
    var ret0 = C.C_ZNK6QColor10lightnessFEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "qreal"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "lightnessF", args)
  }

  return
}

// green()
func (this *QColor) Green(args ...interface{}) (ret interface{}) {
  // green()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK6QColor5greenEv
    // invoke: int green()
    var ret0 = C.C_ZNK6QColor5greenEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "green", args)
  }

  return
}

// fromRgb(QRgb)
func (this *QColor) FromRgb_s(args ...interface{}) (ret interface{}) {
  // fromRgb(QRgb)
  // fromRgb(int, int, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "QRgb"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor7fromRgbEj
    // invoke: QColor fromRgb(QRgb)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN6QColor7fromRgbEj(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN6QColor7fromRgbEiiii
    // invoke: QColor fromRgb(int, int, int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN6QColor7fromRgbEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QColor{}) // "QColor"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QColor", "fromRgb", args)
  }

  return
}

// setGreenF(qreal)
func (this *QColor) SetGreenF(args ...interface{}) () {
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
    // invoke: void setGreenF(qreal)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor9setGreenFEd(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setGreenF", args)
  }

  return
}

// <= body block end

