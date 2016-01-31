package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
// src-file: /QtGui/qpicture.h
// dst-file: /src/gui/qpicture.go
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
  // proto:  void QPictureIO::setParameters(const char * );
extern void C_ZN10QPictureIO13setParametersEPKc(void* qthis, void* arg0); // 4
  // proto:  void QPictureIO::QPictureIO(const QString & fileName, const char * format);
extern void* C_ZN10QPictureIOC2ERK7QStringPKc(void* arg0, void* arg1); // 3
  // proto:  void QPictureIO::QPictureIO(QIODevice * ioDevice, const char * format);
extern void* C_ZN10QPictureIOC2EP9QIODevicePKc(void* arg0, void* arg1); // 3
  // proto:  void QPictureIO::QPictureIO();
extern void* C_ZN10QPictureIOC2Ev(); // 3
  // proto:  void QPictureIO::~QPictureIO();
extern void C_ZN10QPictureIOD2Ev(void* qthis); // 4
  // proto:  int QPictureIO::quality();
extern int32_t C_ZNK10QPictureIO7qualityEv(void* qthis); // 4
  // proto:  void QPictureIO::setDescription(const QString & );
extern void C_ZN10QPictureIO14setDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  const char * QPictureIO::parameters();
extern void* C_ZNK10QPictureIO10parametersEv(void* qthis); // 4
  // proto:  void QPictureIO::setFormat(const char * );
extern void C_ZN10QPictureIO9setFormatEPKc(void* qthis, void* arg0); // 4
  // proto:  QIODevice * QPictureIO::ioDevice();
extern void* C_ZNK10QPictureIO8ioDeviceEv(void* qthis); // 4
  // proto:  void QPictureIO::setPicture(const QPicture & );
extern void C_ZN10QPictureIO10setPictureERK8QPicture(void* qthis, void* arg0); // 4
  // proto:  void QPictureIO::setStatus(int );
extern void C_ZN10QPictureIO9setStatusEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QPictureIO::write();
extern bool C_ZN10QPictureIO5writeEv(void* qthis); // 4
  // proto: static QList<QByteArray> QPictureIO::outputFormats();
extern void C_ZN10QPictureIO13outputFormatsEv(); // 4
  // proto: static QByteArray QPictureIO::pictureFormat(const QString & fileName);
extern void* C_ZN10QPictureIO13pictureFormatERK7QString(void* arg0); // 4
  // proto: static QByteArray QPictureIO::pictureFormat(QIODevice * );
extern void* C_ZN10QPictureIO13pictureFormatEP9QIODevice(void* arg0); // 4
  // proto:  int QPictureIO::status();
extern int32_t C_ZNK10QPictureIO6statusEv(void* qthis); // 4
  // proto:  const QPicture & QPictureIO::picture();
extern void* C_ZNK10QPictureIO7pictureEv(void* qthis); // 4
  // proto:  QString QPictureIO::description();
extern void* C_ZNK10QPictureIO11descriptionEv(void* qthis); // 4
  // proto:  const char * QPictureIO::format();
extern void* C_ZNK10QPictureIO6formatEv(void* qthis); // 4
  // proto:  bool QPictureIO::read();
extern bool C_ZN10QPictureIO4readEv(void* qthis); // 4
  // proto:  QString QPictureIO::fileName();
extern void* C_ZNK10QPictureIO8fileNameEv(void* qthis); // 4
  // proto:  void QPictureIO::setGamma(float );
extern void C_ZN10QPictureIO8setGammaEf(void* qthis, float arg0); // 4
  // proto: static QList<QByteArray> QPictureIO::inputFormats();
extern void C_ZN10QPictureIO12inputFormatsEv(); // 4
  // proto:  void QPictureIO::setIODevice(QIODevice * );
extern void C_ZN10QPictureIO11setIODeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  void QPictureIO::setFileName(const QString & );
extern void C_ZN10QPictureIO11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QPictureIO::setQuality(int );
extern void C_ZN10QPictureIO10setQualityEi(void* qthis, int32_t arg0); // 4
  // proto:  float QPictureIO::gamma();
extern float C_ZNK10QPictureIO5gammaEv(void* qthis); // 4
  // proto:  bool QPicture::load(const QString & fileName, const char * format);
extern bool C_ZN8QPicture4loadERK7QStringPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QPicture::load(QIODevice * dev, const char * format);
extern bool C_ZN8QPicture4loadEP9QIODevicePKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QRect QPicture::boundingRect();
extern void* C_ZNK8QPicture12boundingRectEv(void* qthis); // 4
  // proto:  void QPicture::~QPicture();
extern void C_ZN8QPictureD2Ev(void* qthis); // 4
  // proto:  void QPicture::QPicture(int formatVersion);
extern void* C_ZN8QPictureC2Ei(int32_t arg0); // 3
  // proto:  void QPicture::QPicture(const QPicture & );
extern void* C_ZN8QPictureC2ERKS_(void* arg0); // 3
  // proto:  QPaintEngine * QPicture::paintEngine();
extern void* C_ZNK8QPicture11paintEngineEv(void* qthis); // 4
  // proto:  uint QPicture::size();
extern int32_t C_ZNK8QPicture4sizeEv(void* qthis); // 4
  // proto: static QStringList QPicture::outputFormatList();
extern void C_ZN8QPicture16outputFormatListEv(); // 4
  // proto:  void QPicture::setBoundingRect(const QRect & r);
extern void C_ZN8QPicture15setBoundingRectERK5QRect(void* qthis, void* arg0); // 4
  // proto: static QList<QByteArray> QPicture::outputFormats();
extern void C_ZN8QPicture13outputFormatsEv(); // 4
  // proto:  void QPicture::swap(QPicture & other);
extern void C_ZN8QPicture4swapERS_(void* qthis, void* arg0); // 2
  // proto: static const char * QPicture::pictureFormat(const QString & fileName);
extern void* C_ZN8QPicture13pictureFormatERK7QString(void* arg0); // 4
  // proto:  bool QPicture::save(QIODevice * dev, const char * format);
extern bool C_ZN8QPicture4saveEP9QIODevicePKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QPicture::save(const QString & fileName, const char * format);
extern bool C_ZN8QPicture4saveERK7QStringPKc(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QPicture::setData(const char * data, uint size);
extern void C_ZN8QPicture7setDataEPKcj(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QPicture::isDetached();
extern bool C_ZNK8QPicture10isDetachedEv(void* qthis); // 4
  // proto:  bool QPicture::play(QPainter * p);
extern bool C_ZN8QPicture4playEP8QPainter(void* qthis, void* arg0); // 4
  // proto: static QList<QByteArray> QPicture::inputFormats();
extern void C_ZN8QPicture12inputFormatsEv(); // 4
  // proto: static QStringList QPicture::inputFormatList();
extern void C_ZN8QPicture15inputFormatListEv(); // 4
  // proto:  void QPicture::detach();
extern void C_ZN8QPicture6detachEv(void* qthis); // 4
  // proto:  const char * QPicture::data();
extern void* C_ZNK8QPicture4dataEv(void* qthis); // 4
  // proto:  bool QPicture::isNull();
extern bool C_ZNK8QPicture6isNullEv(void* qthis); // 4
  // proto:  int QPicture::devType();
extern int32_t C_ZNK8QPicture7devTypeEv(void* qthis); // 4
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

// class sizeof(QPictureIO)=8
type QPictureIO struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QPicture)=1
type QPicture struct {
  /*qbase*/ QPaintDevice;
  qclsinst unsafe.Pointer /* *C.void */;
}

// setParameters(const char *)
func (this *QPictureIO) Setparameters(args ...interface{}) () {
  // setParameters(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO13setParametersEPKc
    // invoke: void setParameters(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN10QPictureIO13setParametersEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setParameters", args)
  }

  return
}

// QPictureIO(const class QString &, const char *)
func NewQPictureIO(args ...interface{}) *QPictureIO {
  // QPictureIO(const class QString &, const char *)
  // QPictureIO(class QIODevice *, const char *)
  // QPictureIO()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIOC1ERK7QStringPKc
    // invoke: void QPictureIO(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QPictureIOC2ERK7QStringPKc(arg0, arg1)
    return &QPictureIO{qclsinst:qthis}
  case 1:
    // invoke: _ZN10QPictureIOC1EP9QIODevicePKc
    // invoke: void QPictureIO(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QPictureIOC2EP9QIODevicePKc(arg0, arg1)
    return &QPictureIO{qclsinst:qthis}
  case 2:
    // invoke: _ZN10QPictureIOC1Ev
    // invoke: void QPictureIO()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QPictureIOC2Ev()
    return &QPictureIO{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPictureIO", "QPictureIO", args)
  }

  return nil // QPictureIO{qclsinst:qthis}
}

// ~QPictureIO()
func (this *QPictureIO) Freeqpictureio(args ...interface{}) () {
  // ~QPictureIO()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIOD0Ev
    // invoke: void ~QPictureIO()
    C.C_ZN10QPictureIOD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "~QPictureIO", args)
  }

  return
}

// quality()
func (this *QPictureIO) Quality(args ...interface{}) (ret interface{}) {
  // quality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO7qualityEv
    // invoke: int quality()
    var ret0 = C.C_ZNK10QPictureIO7qualityEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "quality", args)
  }

  return
}

// setDescription(const class QString &)
func (this *QPictureIO) Setdescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO14setDescriptionERK7QString
    // invoke: void setDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO14setDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setDescription", args)
  }

  return
}

// parameters()
func (this *QPictureIO) Parameters(args ...interface{}) (ret interface{}) {
  // parameters()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO10parametersEv
    // invoke: const char * parameters()
    var ret0 = C.C_ZNK10QPictureIO10parametersEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "parameters", args)
  }

  return
}

// setFormat(const char *)
func (this *QPictureIO) Setformat(args ...interface{}) () {
  // setFormat(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO9setFormatEPKc
    // invoke: void setFormat(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN10QPictureIO9setFormatEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setFormat", args)
  }

  return
}

// ioDevice()
func (this *QPictureIO) Iodevice(args ...interface{}) (ret interface{}) {
  // ioDevice()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO8ioDeviceEv
    // invoke: QIODevice * ioDevice()
    var ret0 = C.C_ZNK10QPictureIO8ioDeviceEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "ioDevice", args)
  }

  return
}

// setPicture(const class QPicture &)
func (this *QPictureIO) Setpicture(args ...interface{}) () {
  // setPicture(const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO10setPictureERK8QPicture
    // invoke: void setPicture(const class QPicture &)
    var arg0 = args[0].(QPicture).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO10setPictureERK8QPicture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setPicture", args)
  }

  return
}

// setStatus(int)
func (this *QPictureIO) Setstatus(args ...interface{}) () {
  // setStatus(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO9setStatusEi
    // invoke: void setStatus(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO9setStatusEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setStatus", args)
  }

  return
}

// write()
func (this *QPictureIO) Write(args ...interface{}) (ret interface{}) {
  // write()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO5writeEv
    // invoke: bool write()
    var ret0 = C.C_ZN10QPictureIO5writeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "write", args)
  }

  return
}

// outputFormats()
func (this *QPictureIO) Outputformats_S(args ...interface{}) () {
  // outputFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO13outputFormatsEv
    // invoke: QList<QByteArray> outputFormats()
    C.C_ZN10QPictureIO13outputFormatsEv()
  default:
    qtrt.ErrorResolve("QPictureIO", "outputFormats", args)
  }

  return
}

// pictureFormat(const class QString &)
func (this *QPictureIO) Pictureformat_S(args ...interface{}) (ret interface{}) {
  // pictureFormat(const class QString &)
  // pictureFormat(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO13pictureFormatERK7QString
    // invoke: QByteArray pictureFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QPictureIO13pictureFormatERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN10QPictureIO13pictureFormatEP9QIODevice
    // invoke: QByteArray pictureFormat(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QPictureIO13pictureFormatEP9QIODevice(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "pictureFormat", args)
  }

  return
}

// status()
func (this *QPictureIO) Status(args ...interface{}) (ret interface{}) {
  // status()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO6statusEv
    // invoke: int status()
    var ret0 = C.C_ZNK10QPictureIO6statusEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "status", args)
  }

  return
}

// picture()
func (this *QPictureIO) Picture(args ...interface{}) (ret interface{}) {
  // picture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO7pictureEv
    // invoke: const QPicture & picture()
    var ret0 = C.C_ZNK10QPictureIO7pictureEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPicture{}) // "const QPicture &"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "picture", args)
  }

  return
}

// description()
func (this *QPictureIO) Description(args ...interface{}) (ret interface{}) {
  // description()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO11descriptionEv
    // invoke: QString description()
    var ret0 = C.C_ZNK10QPictureIO11descriptionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "description", args)
  }

  return
}

// format()
func (this *QPictureIO) Format(args ...interface{}) (ret interface{}) {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO6formatEv
    // invoke: const char * format()
    var ret0 = C.C_ZNK10QPictureIO6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "format", args)
  }

  return
}

// read()
func (this *QPictureIO) Read(args ...interface{}) (ret interface{}) {
  // read()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO4readEv
    // invoke: bool read()
    var ret0 = C.C_ZN10QPictureIO4readEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "read", args)
  }

  return
}

// fileName()
func (this *QPictureIO) Filename(args ...interface{}) (ret interface{}) {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO8fileNameEv
    // invoke: QString fileName()
    var ret0 = C.C_ZNK10QPictureIO8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "fileName", args)
  }

  return
}

// setGamma(float)
func (this *QPictureIO) Setgamma(args ...interface{}) () {
  // setGamma(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO8setGammaEf
    // invoke: void setGamma(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO8setGammaEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setGamma", args)
  }

  return
}

// inputFormats()
func (this *QPictureIO) Inputformats_S(args ...interface{}) () {
  // inputFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO12inputFormatsEv
    // invoke: QList<QByteArray> inputFormats()
    C.C_ZN10QPictureIO12inputFormatsEv()
  default:
    qtrt.ErrorResolve("QPictureIO", "inputFormats", args)
  }

  return
}

// setIODevice(class QIODevice *)
func (this *QPictureIO) Setiodevice(args ...interface{}) () {
  // setIODevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO11setIODeviceEP9QIODevice
    // invoke: void setIODevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO11setIODeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setIODevice", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QPictureIO) Setfilename(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO11setFileNameERK7QString
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setFileName", args)
  }

  return
}

// setQuality(int)
func (this *QPictureIO) Setquality(args ...interface{}) () {
  // setQuality(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QPictureIO10setQualityEi
    // invoke: void setQuality(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO10setQualityEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setQuality", args)
  }

  return
}

// gamma()
func (this *QPictureIO) Gamma(args ...interface{}) (ret interface{}) {
  // gamma()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QPictureIO5gammaEv
    // invoke: float gamma()
    var ret0 = C.C_ZNK10QPictureIO5gammaEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPictureIO", "gamma", args)
  }

  return
}

// load(const class QString &, const char *)
func (this *QPicture) Load(args ...interface{}) (ret interface{}) {
  // load(const class QString &, const char *)
  // load(class QIODevice *, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture4loadERK7QStringPKc
    // invoke: bool load(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN8QPicture4loadERK7QStringPKc(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN8QPicture4loadEP9QIODevicePKc
    // invoke: bool load(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN8QPicture4loadEP9QIODevicePKc(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "load", args)
  }

  return
}

// boundingRect()
func (this *QPicture) Boundingrect(args ...interface{}) (ret interface{}) {
  // boundingRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture12boundingRectEv
    // invoke: QRect boundingRect()
    var ret0 = C.C_ZNK8QPicture12boundingRectEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QRect{}) // "QRect"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "boundingRect", args)
  }

  return
}

// ~QPicture()
func (this *QPicture) Freeqpicture(args ...interface{}) () {
  // ~QPicture()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPictureD0Ev
    // invoke: void ~QPicture()
    C.C_ZN8QPictureD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "~QPicture", args)
  }

  return
}

// QPicture(int)
func NewQPicture(args ...interface{}) *QPicture {
  // QPicture(int)
  // QPicture(const class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QPicture{}) // "const QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPictureC1Ei
    // invoke: void QPicture(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPictureC2Ei(arg0)
    return &QPicture{qclsinst:qthis}
  case 1:
    // invoke: _ZN8QPictureC1ERKS_
    // invoke: void QPicture(const class QPicture &)
    var arg0 = args[0].(QPicture).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN8QPictureC2ERKS_(arg0)
    return &QPicture{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QPicture", "QPicture", args)
  }

  return nil // QPicture{qclsinst:qthis}
}

// paintEngine()
func (this *QPicture) Paintengine(args ...interface{}) (ret interface{}) {
  // paintEngine()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture11paintEngineEv
    // invoke: QPaintEngine * paintEngine()
    var ret0 = C.C_ZNK8QPicture11paintEngineEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QPaintEngine{}) // "QPaintEngine *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "paintEngine", args)
  }

  return
}

// size()
func (this *QPicture) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture4sizeEv
    // invoke: uint size()
    var ret0 = C.C_ZNK8QPicture4sizeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "size", args)
  }

  return
}

// outputFormatList()
func (this *QPicture) Outputformatlist_S(args ...interface{}) () {
  // outputFormatList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture16outputFormatListEv
    // invoke: QStringList outputFormatList()
    C.C_ZN8QPicture16outputFormatListEv()
  default:
    qtrt.ErrorResolve("QPicture", "outputFormatList", args)
  }

  return
}

// setBoundingRect(const class QRect &)
func (this *QPicture) Setboundingrect(args ...interface{}) () {
  // setBoundingRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture15setBoundingRectERK5QRect
    // invoke: void setBoundingRect(const class QRect &)
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPicture15setBoundingRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "setBoundingRect", args)
  }

  return
}

// outputFormats()
func (this *QPicture) Outputformats_S(args ...interface{}) () {
  // outputFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture13outputFormatsEv
    // invoke: QList<QByteArray> outputFormats()
    C.C_ZN8QPicture13outputFormatsEv()
  default:
    qtrt.ErrorResolve("QPicture", "outputFormats", args)
  }

  return
}

// swap(class QPicture &)
func (this *QPicture) Swap(args ...interface{}) () {
  // swap(class QPicture &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPicture{}) // "QPicture &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture4swapERS_
    // invoke: void swap(class QPicture &)
    var arg0 = args[0].(QPicture).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN8QPicture4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "swap", args)
  }

  return
}

// pictureFormat(const class QString &)
func (this *QPicture) Pictureformat_S(args ...interface{}) (ret interface{}) {
  // pictureFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture13pictureFormatERK7QString
    // invoke: const char * pictureFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QPicture13pictureFormatERK7QString(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "pictureFormat", args)
  }

  return
}

// save(class QIODevice *, const char *)
func (this *QPicture) Save(args ...interface{}) (ret interface{}) {
  // save(class QIODevice *, const char *)
  // save(const class QString &, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture4saveEP9QIODevicePKc
    // invoke: bool save(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN8QPicture4saveEP9QIODevicePKc(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  case 1:
    // invoke: _ZN8QPicture4saveERK7QStringPKc
    // invoke: bool save(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[1][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN8QPicture4saveERK7QStringPKc(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "save", args)
  }

  return
}

// setData(const char *, uint)
func (this *QPicture) Setdata(args ...interface{}) () {
  // setData(const char *, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture7setDataEPKcj
    // invoke: void setData(const char *, uint)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPicture7setDataEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "setData", args)
  }

  return
}

// isDetached()
func (this *QPicture) Isdetached(args ...interface{}) (ret interface{}) {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture10isDetachedEv
    // invoke: bool isDetached()
    var ret0 = C.C_ZNK8QPicture10isDetachedEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "isDetached", args)
  }

  return
}

// play(class QPainter *)
func (this *QPicture) Play(args ...interface{}) (ret interface{}) {
  // play(class QPainter *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QPainter{}) // "QPainter *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture4playEP8QPainter
    // invoke: bool play(class QPainter *)
    var arg0 = args[0].(QPainter).qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN8QPicture4playEP8QPainter(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "play", args)
  }

  return
}

// inputFormats()
func (this *QPicture) Inputformats_S(args ...interface{}) () {
  // inputFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture12inputFormatsEv
    // invoke: QList<QByteArray> inputFormats()
    C.C_ZN8QPicture12inputFormatsEv()
  default:
    qtrt.ErrorResolve("QPicture", "inputFormats", args)
  }

  return
}

// inputFormatList()
func (this *QPicture) Inputformatlist_S(args ...interface{}) () {
  // inputFormatList()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture15inputFormatListEv
    // invoke: QStringList inputFormatList()
    C.C_ZN8QPicture15inputFormatListEv()
  default:
    qtrt.ErrorResolve("QPicture", "inputFormatList", args)
  }

  return
}

// detach()
func (this *QPicture) Detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QPicture6detachEv
    // invoke: void detach()
    C.C_ZN8QPicture6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "detach", args)
  }

  return
}

// data()
func (this *QPicture) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture4dataEv
    // invoke: const char * data()
    var ret0 = C.C_ZNK8QPicture4dataEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "data", args)
  }

  return
}

// isNull()
func (this *QPicture) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK8QPicture6isNullEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "isNull", args)
  }

  return
}

// devType()
func (this *QPicture) Devtype(args ...interface{}) (ret interface{}) {
  // devType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QPicture7devTypeEv
    // invoke: int devType()
    var ret0 = C.C_ZNK8QPicture7devTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QPicture", "devType", args)
  }

  return
}

// <= body block end

