package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  qreal QColor::hueF();
extern void C_ZNK6QColor4hueFEv(void* qthis); // 4
  // proto:  int QColor::yellow();
extern void C_ZNK6QColor6yellowEv(void* qthis); // 4
  // proto:  qreal QColor::cyanF();
extern void C_ZNK6QColor5cyanFEv(void* qthis); // 4
  // proto:  qreal QColor::magentaF();
extern void C_ZNK6QColor8magentaFEv(void* qthis); // 4
  // proto: static bool QColor::isValidColor(const QString & name);
extern void C_ZN6QColor12isValidColorERK7QString(void* arg0); // 4
  // proto:  void QColor::setHsvF(qreal h, qreal s, qreal v, qreal a);
extern void C_ZN6QColor7setHsvFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QColor::setHsl(int h, int s, int l, int a);
extern void C_ZN6QColor6setHslEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QColor::setCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
extern void C_ZN6QColor8setCmykFEddddd(void* qthis, double arg0, double arg1, double arg2, double arg3, double arg4); // 4
  // proto:  void QColor::setHsv(int h, int s, int v, int a);
extern void C_ZN6QColor6setHsvEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  QRgb QColor::rgb();
extern void C_ZNK6QColor3rgbEv(void* qthis); // 4
  // proto:  int QColor::black();
extern void C_ZNK6QColor5blackEv(void* qthis); // 4
  // proto: static QColor QColor::fromCmyk(int c, int m, int y, int k, int a);
extern void C_ZN6QColor8fromCmykEiiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  QColor::Spec QColor::spec();
extern void C_ZNK6QColor4specEv(void* qthis); // 2
  // proto:  int QColor::saturation();
extern void C_ZNK6QColor10saturationEv(void* qthis); // 4
  // proto:  void QColor::setCmyk(int c, int m, int y, int k, int a);
extern void C_ZN6QColor7setCmykEiiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3, int32_t arg4); // 4
  // proto:  QColor QColor::lighter(int f);
extern void C_ZNK6QColor7lighterEi(void* qthis, int32_t arg0); // 2
  // proto:  qreal QColor::yellowF();
extern void C_ZNK6QColor7yellowFEv(void* qthis); // 4
  // proto: static QColor QColor::fromRgbF(qreal r, qreal g, qreal b, qreal a);
extern void C_ZN6QColor8fromRgbFEdddd(double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QColor::getHsvF(qreal * h, qreal * s, qreal * v, qreal * a);
extern void C_ZNK6QColor7getHsvFEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  QString QColor::name();
extern void C_ZNK6QColor4nameEv(void* qthis); // 4
  // proto:  QColor QColor::toRgb();
extern void C_ZNK6QColor5toRgbEv(void* qthis); // 4
  // proto:  qreal QColor::hsvHueF();
extern void C_ZNK6QColor7hsvHueFEv(void* qthis); // 4
  // proto: static QColor QColor::fromRgba(QRgb rgba);
extern void C_ZN6QColor8fromRgbaEj(int32_t arg0); // 4
  // proto:  void QColor::setAlphaF(qreal alpha);
extern void C_ZN6QColor9setAlphaFEd(void* qthis, double arg0); // 4
  // proto:  int QColor::hslSaturation();
extern void C_ZNK6QColor13hslSaturationEv(void* qthis); // 4
  // proto:  qreal QColor::blackF();
extern void C_ZNK6QColor6blackFEv(void* qthis); // 4
  // proto:  qreal QColor::saturationF();
extern void C_ZNK6QColor11saturationFEv(void* qthis); // 4
  // proto:  void QColor::setRgb(QRgb rgb);
extern void C_ZN6QColor6setRgbEj(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::setRgb(int r, int g, int b, int a);
extern void C_ZN6QColor6setRgbEiiii(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QColor::blue();
extern void C_ZNK6QColor4blueEv(void* qthis); // 4
  // proto:  qreal QColor::blueF();
extern void C_ZNK6QColor5blueFEv(void* qthis); // 4
  // proto:  QRgb QColor::rgba();
extern void C_ZNK6QColor4rgbaEv(void* qthis); // 4
  // proto:  int QColor::red();
extern void C_ZNK6QColor3redEv(void* qthis); // 4
  // proto: static QColor QColor::fromHsvF(qreal h, qreal s, qreal v, qreal a);
extern void C_ZN6QColor8fromHsvFEdddd(double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  qreal QColor::alphaF();
extern void C_ZNK6QColor6alphaFEv(void* qthis); // 4
  // proto:  void QColor::setRedF(qreal red);
extern void C_ZN6QColor7setRedFEd(void* qthis, double arg0); // 4
  // proto:  int QColor::lightness();
extern void C_ZNK6QColor9lightnessEv(void* qthis); // 4
  // proto: static QStringList QColor::colorNames();
extern void C_ZN6QColor10colorNamesEv(); // 4
  // proto:  void QColor::getCmykF(qreal * c, qreal * m, qreal * y, qreal * k, qreal * a);
extern void C_ZN6QColor8getCmykFEPdS0_S0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3, double* arg4); // 4
  // proto:  void QColor::setNamedColor(const QString & name);
extern void C_ZN6QColor13setNamedColorERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QColor::setGreen(int green);
extern void C_ZN6QColor8setGreenEi(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::setRed(int red);
extern void C_ZN6QColor6setRedEi(void* qthis, int32_t arg0); // 4
  // proto:  QColor QColor::toCmyk();
extern void C_ZNK6QColor6toCmykEv(void* qthis); // 4
  // proto:  qreal QColor::redF();
extern void C_ZNK6QColor4redFEv(void* qthis); // 4
  // proto:  int QColor::hue();
extern void C_ZNK6QColor3hueEv(void* qthis); // 4
  // proto:  qreal QColor::hslSaturationF();
extern void C_ZNK6QColor14hslSaturationFEv(void* qthis); // 4
  // proto:  qreal QColor::hsvSaturationF();
extern void C_ZNK6QColor14hsvSaturationFEv(void* qthis); // 4
  // proto:  QColor QColor::toHsl();
extern void C_ZNK6QColor5toHslEv(void* qthis); // 4
  // proto:  int QColor::value();
extern void C_ZNK6QColor5valueEv(void* qthis); // 4
  // proto:  QColor QColor::toHsv();
extern void C_ZNK6QColor5toHsvEv(void* qthis); // 4
  // proto:  QColor QColor::darker(int f);
extern void C_ZNK6QColor6darkerEi(void* qthis, int32_t arg0); // 2
  // proto:  void QColor::setBlue(int blue);
extern void C_ZN6QColor7setBlueEi(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::setHslF(qreal h, qreal s, qreal l, qreal a);
extern void C_ZN6QColor7setHslFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  qreal QColor::valueF();
extern void C_ZNK6QColor6valueFEv(void* qthis); // 4
  // proto: static QColor QColor::fromHsl(int h, int s, int l, int a);
extern void C_ZN6QColor7fromHslEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  int QColor::hslHue();
extern void C_ZNK6QColor6hslHueEv(void* qthis); // 4
  // proto: static QColor QColor::fromHsv(int h, int s, int v, int a);
extern void C_ZN6QColor7fromHsvEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QColor::getRgb(int * r, int * g, int * b, int * a);
extern void C_ZNK6QColor6getRgbEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  void QColor::setRgba(QRgb rgba);
extern void C_ZN6QColor7setRgbaEj(void* qthis, int32_t arg0); // 4
  // proto:  int QColor::magenta();
extern void C_ZNK6QColor7magentaEv(void* qthis); // 4
  // proto:  void QColor::QColor(const char * name);
extern void* C_ZN6QColorC2EPKc(unsigned char* arg0); // 1
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
extern void C_ZNK6QColor7hslHueFEv(void* qthis); // 4
  // proto: static QColor QColor::fromHslF(qreal h, qreal s, qreal l, qreal a);
extern void C_ZN6QColor8fromHslFEdddd(double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  int QColor::hsvSaturation();
extern void C_ZNK6QColor13hsvSaturationEv(void* qthis); // 4
  // proto:  void QColor::setRgbF(qreal r, qreal g, qreal b, qreal a);
extern void C_ZN6QColor7setRgbFEdddd(void* qthis, double arg0, double arg1, double arg2, double arg3); // 4
  // proto:  void QColor::setAlpha(int alpha);
extern void C_ZN6QColor8setAlphaEi(void* qthis, int32_t arg0); // 4
  // proto:  int QColor::cyan();
extern void C_ZNK6QColor4cyanEv(void* qthis); // 4
  // proto:  void QColor::getRgbF(qreal * r, qreal * g, qreal * b, qreal * a);
extern void C_ZNK6QColor7getRgbFEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  void QColor::setBlueF(qreal blue);
extern void C_ZN6QColor8setBlueFEd(void* qthis, double arg0); // 4
  // proto:  void QColor::getCmyk(int * c, int * m, int * y, int * k, int * a);
extern void C_ZN6QColor7getCmykEPiS0_S0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3, int32_t* arg4); // 4
  // proto:  bool QColor::isValid();
extern void C_ZNK6QColor7isValidEv(void* qthis); // 2
  // proto:  void QColor::getHsv(int * h, int * s, int * v, int * a);
extern void C_ZNK6QColor6getHsvEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto: static QColor QColor::fromCmykF(qreal c, qreal m, qreal y, qreal k, qreal a);
extern void C_ZN6QColor9fromCmykFEddddd(double arg0, double arg1, double arg2, double arg3, double arg4); // 4
  // proto:  int QColor::hsvHue();
extern void C_ZNK6QColor6hsvHueEv(void* qthis); // 4
  // proto:  void QColor::getHsl(int * h, int * s, int * l, int * a);
extern void C_ZNK6QColor6getHslEPiS0_S0_S0_(void* qthis, int32_t* arg0, int32_t* arg1, int32_t* arg2, int32_t* arg3); // 4
  // proto:  int QColor::alpha();
extern void C_ZNK6QColor5alphaEv(void* qthis); // 4
  // proto:  QColor QColor::dark(int f);
extern void C_ZNK6QColor4darkEi(void* qthis, int32_t arg0); // 4
  // proto:  qreal QColor::greenF();
extern void C_ZNK6QColor6greenFEv(void* qthis); // 4
  // proto:  QColor QColor::light(int f);
extern void C_ZNK6QColor5lightEi(void* qthis, int32_t arg0); // 4
  // proto:  void QColor::getHslF(qreal * h, qreal * s, qreal * l, qreal * a);
extern void C_ZNK6QColor7getHslFEPdS0_S0_S0_(void* qthis, double* arg0, double* arg1, double* arg2, double* arg3); // 4
  // proto:  qreal QColor::lightnessF();
extern void C_ZNK6QColor10lightnessFEv(void* qthis); // 4
  // proto:  int QColor::green();
extern void C_ZNK6QColor5greenEv(void* qthis); // 4
  // proto: static QColor QColor::fromRgb(QRgb rgb);
extern void C_ZN6QColor7fromRgbEj(int32_t arg0); // 4
  // proto: static QColor QColor::fromRgb(int r, int g, int b, int a);
extern void C_ZN6QColor7fromRgbEiiii(int32_t arg0, int32_t arg1, int32_t arg2, int32_t arg3); // 4
  // proto:  void QColor::setGreenF(qreal green);
extern void C_ZN6QColor9setGreenFEd(void* qthis, double arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// hueF()
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
    // invoke: qreal hueF()
    var ret = C.C_ZNK6QColor4hueFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hueF", args)
  }

}

// yellow()
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
    // invoke: int yellow()
    var ret = C.C_ZNK6QColor6yellowEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "yellow", args)
  }

}

// cyanF()
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
    // invoke: qreal cyanF()
    var ret = C.C_ZNK6QColor5cyanFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "cyanF", args)
  }

}

// magentaF()
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
    // invoke: qreal magentaF()
    var ret = C.C_ZNK6QColor8magentaFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "magentaF", args)
  }

}

// isValidColor(const class QString &)
func (this *QColor) isValidColor_s(args ...interface{}) () {
  // isValidColor(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColor12isValidColorERK7QString
    // invoke: bool isValidColor(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN6QColor12isValidColorERK7QString(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "isValidColor", args)
  }

}

// setHsvF(qreal, qreal, qreal, qreal)
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
    // invoke: void setHsvF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor7setHsvFEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHsvF", args)
  }

}

// setHsl(int, int, int, int)
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
    // invoke: void setHsl(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor6setHslEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHsl", args)
  }

}

// setCmykF(qreal, qreal, qreal, qreal, qreal)
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
    // invoke: void setCmykF(qreal, qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor8setCmykFEddddd(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "setCmykF", args)
  }

}

// setHsv(int, int, int, int)
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
    // invoke: void setHsv(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor6setHsvEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHsv", args)
  }

}

// rgb()
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
    // invoke: QRgb rgb()
    var ret = C.C_ZNK6QColor3rgbEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "rgb", args)
  }

}

// black()
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
    // invoke: int black()
    var ret = C.C_ZNK6QColor5blackEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "black", args)
  }

}

// fromCmyk(int, int, int, int, int)
func (this *QColor) fromCmyk_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZN6QColor8fromCmykEiiiii(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromCmyk", args)
  }

}

// spec()
func (this *QColor) spec(args ...interface{}) () {
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
    C.C_ZNK6QColor4specEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QColor", "spec", args)
  }

}

// saturation()
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
    // invoke: int saturation()
    var ret = C.C_ZNK6QColor10saturationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "saturation", args)
  }

}

// setCmyk(int, int, int, int, int)
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
    // invoke: void setCmyk(int, int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = C.int32_t(args[4].(int32))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor7setCmykEiiiii(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "setCmyk", args)
  }

}

// lighter(int)
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
    // invoke: QColor lighter(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK6QColor7lighterEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "lighter", args)
  }

}

// yellowF()
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
    // invoke: qreal yellowF()
    var ret = C.C_ZNK6QColor7yellowFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "yellowF", args)
  }

}

// fromRgbF(qreal, qreal, qreal, qreal)
func (this *QColor) fromRgbF_s(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN6QColor8fromRgbFEdddd(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromRgbF", args)
  }

}

// getHsvF(qreal *, qreal *, qreal *, qreal *)
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
    // invoke: void getHsvF(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor7getHsvFEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHsvF", args)
  }

}

// name()
func (this *QColor) name(args ...interface{}) () {
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
    var ret = C.C_ZNK6QColor4nameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "name", args)
  }

}

// toRgb()
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
    // invoke: QColor toRgb()
    var ret = C.C_ZNK6QColor5toRgbEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "toRgb", args)
  }

}

// hsvHueF()
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
    // invoke: qreal hsvHueF()
    var ret = C.C_ZNK6QColor7hsvHueFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hsvHueF", args)
  }

}

// fromRgba(QRgb)
func (this *QColor) fromRgba_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN6QColor8fromRgbaEj(arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromRgba", args)
  }

}

// setAlphaF(qreal)
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
    // invoke: void setAlphaF(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor9setAlphaFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setAlphaF", args)
  }

}

// hslSaturation()
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
    // invoke: int hslSaturation()
    var ret = C.C_ZNK6QColor13hslSaturationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hslSaturation", args)
  }

}

// blackF()
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
    // invoke: qreal blackF()
    var ret = C.C_ZNK6QColor6blackFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "blackF", args)
  }

}

// saturationF()
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
    // invoke: qreal saturationF()
    var ret = C.C_ZNK6QColor11saturationFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "saturationF", args)
  }

}

// setRgb(QRgb)
func (this *QColor) setRgb(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor6setRgbEj(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN6QColor6setRgbEiiii
    // invoke: void setRgb(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor6setRgbEiiii(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setRgb", args)
  }

}

// blue()
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
    // invoke: int blue()
    var ret = C.C_ZNK6QColor4blueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "blue", args)
  }

}

// blueF()
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
    // invoke: qreal blueF()
    var ret = C.C_ZNK6QColor5blueFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "blueF", args)
  }

}

// rgba()
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
    // invoke: QRgb rgba()
    var ret = C.C_ZNK6QColor4rgbaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "rgba", args)
  }

}

// red()
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
    // invoke: int red()
    var ret = C.C_ZNK6QColor3redEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "red", args)
  }

}

// fromHsvF(qreal, qreal, qreal, qreal)
func (this *QColor) fromHsvF_s(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN6QColor8fromHsvFEdddd(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromHsvF", args)
  }

}

// alphaF()
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
    // invoke: qreal alphaF()
    var ret = C.C_ZNK6QColor6alphaFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "alphaF", args)
  }

}

// setRedF(qreal)
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
    // invoke: void setRedF(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor7setRedFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setRedF", args)
  }

}

// lightness()
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
    // invoke: int lightness()
    var ret = C.C_ZNK6QColor9lightnessEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "lightness", args)
  }

}

// colorNames()
func (this *QColor) colorNames_s(args ...interface{}) () {
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

}

// getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
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
    // invoke: void getCmykF(qreal *, qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.double)(args[4].(*float64))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor8getCmykFEPdS0_S0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "getCmykF", args)
  }

}

// setNamedColor(const class QString &)
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
    // invoke: void setNamedColor(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor13setNamedColorERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setNamedColor", args)
  }

}

// setGreen(int)
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
    // invoke: void setGreen(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor8setGreenEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setGreen", args)
  }

}

// setRed(int)
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
    // invoke: void setRed(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor6setRedEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setRed", args)
  }

}

// toCmyk()
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
    // invoke: QColor toCmyk()
    var ret = C.C_ZNK6QColor6toCmykEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "toCmyk", args)
  }

}

// redF()
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
    // invoke: qreal redF()
    var ret = C.C_ZNK6QColor4redFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "redF", args)
  }

}

// hue()
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
    // invoke: int hue()
    var ret = C.C_ZNK6QColor3hueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hue", args)
  }

}

// hslSaturationF()
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
    // invoke: qreal hslSaturationF()
    var ret = C.C_ZNK6QColor14hslSaturationFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hslSaturationF", args)
  }

}

// hsvSaturationF()
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
    // invoke: qreal hsvSaturationF()
    var ret = C.C_ZNK6QColor14hsvSaturationFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hsvSaturationF", args)
  }

}

// toHsl()
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
    // invoke: QColor toHsl()
    var ret = C.C_ZNK6QColor5toHslEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "toHsl", args)
  }

}

// value()
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
    // invoke: int value()
    var ret = C.C_ZNK6QColor5valueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "value", args)
  }

}

// toHsv()
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
    // invoke: QColor toHsv()
    var ret = C.C_ZNK6QColor5toHsvEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "toHsv", args)
  }

}

// darker(int)
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
    // invoke: QColor darker(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK6QColor6darkerEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "darker", args)
  }

}

// setBlue(int)
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
    // invoke: void setBlue(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor7setBlueEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setBlue", args)
  }

}

// setHslF(qreal, qreal, qreal, qreal)
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
    // invoke: void setHslF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor7setHslFEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setHslF", args)
  }

}

// valueF()
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
    // invoke: qreal valueF()
    var ret = C.C_ZNK6QColor6valueFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "valueF", args)
  }

}

// fromHsl(int, int, int, int)
func (this *QColor) fromHsl_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN6QColor7fromHslEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromHsl", args)
  }

}

// hslHue()
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
    // invoke: int hslHue()
    var ret = C.C_ZNK6QColor6hslHueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hslHue", args)
  }

}

// fromHsv(int, int, int, int)
func (this *QColor) fromHsv_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN6QColor7fromHsvEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromHsv", args)
  }

}

// getRgb(int *, int *, int *, int *)
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
    // invoke: void getRgb(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor6getRgbEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getRgb", args)
  }

}

// setRgba(QRgb)
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
    // invoke: void setRgba(QRgb)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor7setRgbaEj(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setRgba", args)
  }

}

// magenta()
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
    // invoke: int magenta()
    var ret = C.C_ZNK6QColor7magentaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "magenta", args)
  }

}

// QColor(const char *)
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
  vtys[5][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN6QColorC1EPKc
    // invoke: void QColor(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2EPKc(arg0)
    return &QColor{qclsinst:qthis}
  case 1:
    // invoke: _ZN6QColorC1Ej
    // invoke: void QColor(QRgb)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2Ej(arg0)
    return &QColor{qclsinst:qthis}
  case 2:
    // invoke: _ZN6QColorC1ERKS_
    // invoke: void QColor(const class QColor &)
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2ERKS_(arg0)
    return &QColor{qclsinst:qthis}
  case 3:
    // invoke: _ZN6QColorC1Eiiii
    // invoke: void QColor(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2Eiiii(arg0, arg1, arg2, arg3)
    return &QColor{qclsinst:qthis}
  case 4:
    // invoke: _ZN6QColorC1Ev
    // invoke: void QColor()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2Ev()
    return &QColor{qclsinst:qthis}
  case 5:
    // invoke: _ZN6QColorC1ERK7QString
    // invoke: void QColor(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN6QColorC2ERK7QString(arg0)
    return &QColor{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QColor", "QColor", args)
  }

  return nil // QColor{qclsinst:qthis}
}

// hslHueF()
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
    // invoke: qreal hslHueF()
    var ret = C.C_ZNK6QColor7hslHueFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hslHueF", args)
  }

}

// fromHslF(qreal, qreal, qreal, qreal)
func (this *QColor) fromHslF_s(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN6QColor8fromHslFEdddd(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromHslF", args)
  }

}

// hsvSaturation()
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
    // invoke: int hsvSaturation()
    var ret = C.C_ZNK6QColor13hsvSaturationEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hsvSaturation", args)
  }

}

// setRgbF(qreal, qreal, qreal, qreal)
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
    // invoke: void setRgbF(qreal, qreal, qreal, qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    C.C_ZN6QColor7setRgbFEdddd(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "setRgbF", args)
  }

}

// setAlpha(int)
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
    // invoke: void setAlpha(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor8setAlphaEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setAlpha", args)
  }

}

// cyan()
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
    // invoke: int cyan()
    var ret = C.C_ZNK6QColor4cyanEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "cyan", args)
  }

}

// getRgbF(qreal *, qreal *, qreal *, qreal *)
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
    // invoke: void getRgbF(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor7getRgbFEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getRgbF", args)
  }

}

// setBlueF(qreal)
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
    // invoke: void setBlueF(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor8setBlueFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setBlueF", args)
  }

}

// getCmyk(int *, int *, int *, int *, int *)
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
    // invoke: void getCmyk(int *, int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    var arg4 = (*C.int32_t)(args[4].(*int32))
    if false {fmt.Println(arg4)}
    C.C_ZN6QColor7getCmykEPiS0_S0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  default:
    qtrt.ErrorResolve("QColor", "getCmyk", args)
  }

}

// isValid()
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
    // invoke: bool isValid()
    var ret = C.C_ZNK6QColor7isValidEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "isValid", args)
  }

}

// getHsv(int *, int *, int *, int *)
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
    // invoke: void getHsv(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor6getHsvEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHsv", args)
  }

}

// fromCmykF(qreal, qreal, qreal, qreal, qreal)
func (this *QColor) fromCmykF_s(args ...interface{}) () {
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
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.double(args[1].(float64))
    if false {fmt.Println(arg1)}
    var arg2 = C.double(args[2].(float64))
    if false {fmt.Println(arg2)}
    var arg3 = C.double(args[3].(float64))
    if false {fmt.Println(arg3)}
    var arg4 = C.double(args[4].(float64))
    if false {fmt.Println(arg4)}
    var ret = C.C_ZN6QColor9fromCmykFEddddd(arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromCmykF", args)
  }

}

// hsvHue()
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
    // invoke: int hsvHue()
    var ret = C.C_ZNK6QColor6hsvHueEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "hsvHue", args)
  }

}

// getHsl(int *, int *, int *, int *)
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
    // invoke: void getHsl(int *, int *, int *, int *)
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.int32_t)(args[1].(*int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.int32_t)(args[2].(*int32))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.int32_t)(args[3].(*int32))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor6getHslEPiS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHsl", args)
  }

}

// alpha()
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
    // invoke: int alpha()
    var ret = C.C_ZNK6QColor5alphaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "alpha", args)
  }

}

// dark(int)
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
    // invoke: QColor dark(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK6QColor4darkEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "dark", args)
  }

}

// greenF()
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
    // invoke: qreal greenF()
    var ret = C.C_ZNK6QColor6greenFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "greenF", args)
  }

}

// light(int)
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
    // invoke: QColor light(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK6QColor5lightEi(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "light", args)
  }

}

// getHslF(qreal *, qreal *, qreal *, qreal *)
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
    // invoke: void getHslF(qreal *, qreal *, qreal *, qreal *)
    var arg0 = (*C.double)(args[0].(*float64))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.double)(args[1].(*float64))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.double)(args[2].(*float64))
    if false {fmt.Println(arg2)}
    var arg3 = (*C.double)(args[3].(*float64))
    if false {fmt.Println(arg3)}
    C.C_ZNK6QColor7getHslFEPdS0_S0_S0_(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QColor", "getHslF", args)
  }

}

// lightnessF()
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
    // invoke: qreal lightnessF()
    var ret = C.C_ZNK6QColor10lightnessFEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "lightnessF", args)
  }

}

// green()
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
    // invoke: int green()
    var ret = C.C_ZNK6QColor5greenEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "green", args)
  }

}

// fromRgb(QRgb)
func (this *QColor) fromRgb_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN6QColor7fromRgbEj(arg0)
    if false {reflect.TypeOf(ret)}
  case 1:
    // invoke: _ZN6QColor7fromRgbEiiii
    // invoke: QColor fromRgb(int, int, int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var ret = C.C_ZN6QColor7fromRgbEiiii(arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QColor", "fromRgb", args)
  }

}

// setGreenF(qreal)
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
    // invoke: void setGreenF(qreal)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    C.C_ZN6QColor9setGreenFEd(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QColor", "setGreenF", args)
  }

}

// <= body block end

