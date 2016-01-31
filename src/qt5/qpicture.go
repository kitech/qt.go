package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
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
extern void C_ZN10QPictureIO13setParametersEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  void QPictureIO::QPictureIO(const QString & fileName, const char * format);
extern void C_ZN10QPictureIOC2ERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1); // 3
  // proto:  void QPictureIO::QPictureIO(QIODevice * ioDevice, const char * format);
extern void C_ZN10QPictureIOC2EP9QIODevicePKc(void* qthis, void* arg0, unsigned char* arg1); // 3
  // proto:  void QPictureIO::QPictureIO();
extern void C_ZN10QPictureIOC2Ev(void* qthis); // 3
  // proto:  void QPictureIO::~QPictureIO();
extern void C_ZN10QPictureIOD2Ev(void* qthis); // 4
  // proto:  int QPictureIO::quality();
extern void C_ZNK10QPictureIO7qualityEv(void* qthis); // 4
  // proto:  void QPictureIO::setDescription(const QString & );
extern void C_ZN10QPictureIO14setDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  const char * QPictureIO::parameters();
extern void C_ZNK10QPictureIO10parametersEv(void* qthis); // 4
  // proto:  void QPictureIO::setFormat(const char * );
extern void C_ZN10QPictureIO9setFormatEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  QIODevice * QPictureIO::ioDevice();
extern void C_ZNK10QPictureIO8ioDeviceEv(void* qthis); // 4
  // proto:  void QPictureIO::setPicture(const QPicture & );
extern void C_ZN10QPictureIO10setPictureERK8QPicture(void* qthis, void* arg0); // 4
  // proto:  void QPictureIO::setStatus(int );
extern void C_ZN10QPictureIO9setStatusEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QPictureIO::write();
extern void C_ZN10QPictureIO5writeEv(void* qthis); // 4
  // proto: static QList<QByteArray> QPictureIO::outputFormats();
extern void C_ZN10QPictureIO13outputFormatsEv(); // 4
  // proto: static QByteArray QPictureIO::pictureFormat(const QString & fileName);
extern void C_ZN10QPictureIO13pictureFormatERK7QString(void* arg0); // 4
  // proto: static QByteArray QPictureIO::pictureFormat(QIODevice * );
extern void C_ZN10QPictureIO13pictureFormatEP9QIODevice(void* arg0); // 4
  // proto:  int QPictureIO::status();
extern void C_ZNK10QPictureIO6statusEv(void* qthis); // 4
  // proto:  const QPicture & QPictureIO::picture();
extern void C_ZNK10QPictureIO7pictureEv(void* qthis); // 4
  // proto:  QString QPictureIO::description();
extern void C_ZNK10QPictureIO11descriptionEv(void* qthis); // 4
  // proto:  const char * QPictureIO::format();
extern void C_ZNK10QPictureIO6formatEv(void* qthis); // 4
  // proto:  bool QPictureIO::read();
extern void C_ZN10QPictureIO4readEv(void* qthis); // 4
  // proto:  QString QPictureIO::fileName();
extern void C_ZNK10QPictureIO8fileNameEv(void* qthis); // 4
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
extern void C_ZNK10QPictureIO5gammaEv(void* qthis); // 4
  // proto:  bool QPicture::load(const QString & fileName, const char * format);
extern void C_ZN8QPicture4loadERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  bool QPicture::load(QIODevice * dev, const char * format);
extern void C_ZN8QPicture4loadEP9QIODevicePKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  QRect QPicture::boundingRect();
extern void C_ZNK8QPicture12boundingRectEv(void* qthis); // 4
  // proto:  void QPicture::~QPicture();
extern void C_ZN8QPictureD2Ev(void* qthis); // 4
  // proto:  void QPicture::QPicture(int formatVersion);
extern void C_ZN8QPictureC2Ei(void* qthis, int32_t arg0); // 3
  // proto:  void QPicture::QPicture(const QPicture & );
extern void C_ZN8QPictureC2ERKS_(void* qthis, void* arg0); // 3
  // proto:  QPaintEngine * QPicture::paintEngine();
extern void C_ZNK8QPicture11paintEngineEv(void* qthis); // 4
  // proto:  uint QPicture::size();
extern void C_ZNK8QPicture4sizeEv(void* qthis); // 4
  // proto: static QStringList QPicture::outputFormatList();
extern void C_ZN8QPicture16outputFormatListEv(); // 4
  // proto:  void QPicture::setBoundingRect(const QRect & r);
extern void C_ZN8QPicture15setBoundingRectERK5QRect(void* qthis, void* arg0); // 4
  // proto: static QList<QByteArray> QPicture::outputFormats();
extern void C_ZN8QPicture13outputFormatsEv(); // 4
  // proto:  void QPicture::swap(QPicture & other);
extern void C_ZN8QPicture4swapERS_(void* qthis, void* arg0); // 2
  // proto: static const char * QPicture::pictureFormat(const QString & fileName);
extern void C_ZN8QPicture13pictureFormatERK7QString(void* arg0); // 4
  // proto:  bool QPicture::save(QIODevice * dev, const char * format);
extern void C_ZN8QPicture4saveEP9QIODevicePKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  bool QPicture::save(const QString & fileName, const char * format);
extern void C_ZN8QPicture4saveERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1); // 4
  // proto:  void QPicture::setData(const char * data, uint size);
extern void C_ZN8QPicture7setDataEPKcj(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  bool QPicture::isDetached();
extern void C_ZNK8QPicture10isDetachedEv(void* qthis); // 4
  // proto:  bool QPicture::play(QPainter * p);
extern void C_ZN8QPicture4playEP8QPainter(void* qthis, void* arg0); // 4
  // proto: static QList<QByteArray> QPicture::inputFormats();
extern void C_ZN8QPicture12inputFormatsEv(); // 4
  // proto: static QStringList QPicture::inputFormatList();
extern void C_ZN8QPicture15inputFormatListEv(); // 4
  // proto:  void QPicture::detach();
extern void C_ZN8QPicture6detachEv(void* qthis); // 4
  // proto:  const char * QPicture::data();
extern void C_ZNK8QPicture4dataEv(void* qthis); // 4
  // proto:  bool QPicture::isNull();
extern void C_ZNK8QPicture6isNullEv(void* qthis); // 4
  // proto:  int QPicture::devType();
extern void C_ZNK8QPicture7devTypeEv(void* qthis); // 4
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
func (this *QPictureIO) setParameters(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO13setParametersEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setParameters", args)
  }

}

// QPictureIO(const class QString &, const char *)
func NewQPictureIO(args ...interface{}) QPictureIO {
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
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QPictureIOC2ERK7QStringPKc(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN10QPictureIOC1EP9QIODevicePKc
    // invoke: void QPictureIO(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QPictureIOC2EP9QIODevicePKc(qthis, arg0, arg1)
  case 2:
    // invoke: _ZN10QPictureIOC1Ev
    // invoke: void QPictureIO()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN10QPictureIOC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QPictureIO", "QPictureIO", args)
  }

  return QPictureIO{}
}

// ~QPictureIO()
func (this *QPictureIO) FreeQPictureIO(args ...interface{}) () {
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

}

// quality()
func (this *QPictureIO) quality(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO7qualityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "quality", args)
  }

}

// setDescription(const class QString &)
func (this *QPictureIO) setDescription(args ...interface{}) () {
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

}

// parameters()
func (this *QPictureIO) parameters(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO10parametersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "parameters", args)
  }

}

// setFormat(const char *)
func (this *QPictureIO) setFormat(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO9setFormatEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setFormat", args)
  }

}

// ioDevice()
func (this *QPictureIO) ioDevice(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO8ioDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "ioDevice", args)
  }

}

// setPicture(const class QPicture &)
func (this *QPictureIO) setPicture(args ...interface{}) () {
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

}

// setStatus(int)
func (this *QPictureIO) setStatus(args ...interface{}) () {
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

}

// write()
func (this *QPictureIO) write(args ...interface{}) () {
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
    C.C_ZN10QPictureIO5writeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "write", args)
  }

}

// outputFormats()
func (this *QPictureIO) outputFormats_s(args ...interface{}) () {
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

}

// pictureFormat(const class QString &)
func (this *QPictureIO) pictureFormat_s(args ...interface{}) () {
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
    C.C_ZN10QPictureIO13pictureFormatERK7QString(arg0)
  case 1:
    // invoke: _ZN10QPictureIO13pictureFormatEP9QIODevice
    // invoke: QByteArray pictureFormat(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QPictureIO13pictureFormatEP9QIODevice(arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "pictureFormat", args)
  }

}

// status()
func (this *QPictureIO) status(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO6statusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "status", args)
  }

}

// picture()
func (this *QPictureIO) picture(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO7pictureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "picture", args)
  }

}

// description()
func (this *QPictureIO) description(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO11descriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "description", args)
  }

}

// format()
func (this *QPictureIO) format(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "format", args)
  }

}

// read()
func (this *QPictureIO) read(args ...interface{}) () {
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
    C.C_ZN10QPictureIO4readEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "read", args)
  }

}

// fileName()
func (this *QPictureIO) fileName(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "fileName", args)
  }

}

// setGamma(float)
func (this *QPictureIO) setGamma(args ...interface{}) () {
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

}

// inputFormats()
func (this *QPictureIO) inputFormats_s(args ...interface{}) () {
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

}

// setIODevice(class QIODevice *)
func (this *QPictureIO) setIODevice(args ...interface{}) () {
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

}

// setFileName(const class QString &)
func (this *QPictureIO) setFileName(args ...interface{}) () {
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

}

// setQuality(int)
func (this *QPictureIO) setQuality(args ...interface{}) () {
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

}

// gamma()
func (this *QPictureIO) gamma(args ...interface{}) () {
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
    C.C_ZNK10QPictureIO5gammaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "gamma", args)
  }

}

// load(const class QString &, const char *)
func (this *QPicture) load(args ...interface{}) () {
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
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPicture4loadERK7QStringPKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPicture4loadEP9QIODevicePKc
    // invoke: bool load(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPicture4loadEP9QIODevicePKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "load", args)
  }

}

// boundingRect()
func (this *QPicture) boundingRect(args ...interface{}) () {
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
    C.C_ZNK8QPicture12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "boundingRect", args)
  }

}

// ~QPicture()
func (this *QPicture) FreeQPicture(args ...interface{}) () {
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

}

// QPicture(int)
func NewQPicture(args ...interface{}) QPicture {
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
    C.C_ZN8QPictureC2Ei(qthis, arg0)
  case 1:
    // invoke: _ZN8QPictureC1ERKS_
    // invoke: void QPicture(const class QPicture &)
    var arg0 = args[0].(QPicture).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN8QPictureC2ERKS_(qthis, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "QPicture", args)
  }

  return QPicture{}
}

// paintEngine()
func (this *QPicture) paintEngine(args ...interface{}) () {
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
    C.C_ZNK8QPicture11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "paintEngine", args)
  }

}

// size()
func (this *QPicture) size(args ...interface{}) () {
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
    C.C_ZNK8QPicture4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "size", args)
  }

}

// outputFormatList()
func (this *QPicture) outputFormatList_s(args ...interface{}) () {
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

}

// setBoundingRect(const class QRect &)
func (this *QPicture) setBoundingRect(args ...interface{}) () {
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

}

// outputFormats()
func (this *QPicture) outputFormats_s(args ...interface{}) () {
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

}

// swap(class QPicture &)
func (this *QPicture) swap(args ...interface{}) () {
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

}

// pictureFormat(const class QString &)
func (this *QPicture) pictureFormat_s(args ...interface{}) () {
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
    C.C_ZN8QPicture13pictureFormatERK7QString(arg0)
  default:
    qtrt.ErrorResolve("QPicture", "pictureFormat", args)
  }

}

// save(class QIODevice *, const char *)
func (this *QPicture) save(args ...interface{}) () {
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
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPicture4saveEP9QIODevicePKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPicture4saveERK7QStringPKc
    // invoke: bool save(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPicture4saveERK7QStringPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "save", args)
  }

}

// setData(const char *, uint)
func (this *QPicture) setData(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.C_ZN8QPicture7setDataEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "setData", args)
  }

}

// isDetached()
func (this *QPicture) isDetached(args ...interface{}) () {
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
    C.C_ZNK8QPicture10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "isDetached", args)
  }

}

// play(class QPainter *)
func (this *QPicture) play(args ...interface{}) () {
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
    C.C_ZN8QPicture4playEP8QPainter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "play", args)
  }

}

// inputFormats()
func (this *QPicture) inputFormats_s(args ...interface{}) () {
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

}

// inputFormatList()
func (this *QPicture) inputFormatList_s(args ...interface{}) () {
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

}

// detach()
func (this *QPicture) detach(args ...interface{}) () {
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

}

// data()
func (this *QPicture) data(args ...interface{}) () {
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
    C.C_ZNK8QPicture4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "data", args)
  }

}

// isNull()
func (this *QPicture) isNull(args ...interface{}) () {
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
    C.C_ZNK8QPicture6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "isNull", args)
  }

}

// devType()
func (this *QPicture) devType(args ...interface{}) () {
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
    C.C_ZNK8QPicture7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "devType", args)
  }

}

// <= body block end

