package qt5
// auto generated, do not modify.
// created: Sat Jan  2 20:07:20 2016
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
  // proto:  void QPictureIO::QPictureIO(const QString & fileName, const char * format);
extern void* dector_ZN10QPictureIOC1ERK7QStringPKc(void* arg0, char* arg1);
extern void _ZN10QPictureIOC1ERK7QStringPKc(void* qthis, void* arg0, char* arg1);
  // proto:  QString QPictureIO::description();
extern void _ZNK10QPictureIO11descriptionEv(void* qthis);
  // proto: static QList<QByteArray> QPictureIO::inputFormats();
extern void _ZN10QPictureIO12inputFormatsEv();
  // proto:  void QPictureIO::setGamma(float );
extern void _ZN10QPictureIO8setGammaEf(void* qthis, float arg0);
  // proto:  int QPictureIO::status();
extern void _ZNK10QPictureIO6statusEv(void* qthis);
  // proto:  int QPictureIO::quality();
extern void _ZNK10QPictureIO7qualityEv(void* qthis);
  // proto:  bool QPictureIO::write();
extern void _ZN10QPictureIO5writeEv(void* qthis);
  // proto:  void QPictureIO::setFileName(const QString & );
extern void _ZN10QPictureIO11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  void QPictureIO::~QPictureIO();
extern void _ZN10QPictureIOD0Ev(void* qthis);
  // proto:  const char * QPictureIO::parameters();
extern void _ZNK10QPictureIO10parametersEv(void* qthis);
  // proto: static QByteArray QPictureIO::pictureFormat(QIODevice * );
extern void _ZN10QPictureIO13pictureFormatEP9QIODevice(void* arg0);
  // proto:  void QPictureIO::QPictureIO(const QPictureIO & );
extern void* dector_ZN10QPictureIOC1ERKS_(void* arg0);
extern void _ZN10QPictureIOC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QPictureIO::read();
extern void _ZN10QPictureIO4readEv(void* qthis);
  // proto:  QString QPictureIO::fileName();
extern void _ZNK10QPictureIO8fileNameEv(void* qthis);
  // proto:  void QPictureIO::QPictureIO(QIODevice * ioDevice, const char * format);
extern void* dector_ZN10QPictureIOC1EP9QIODevicePKc(void* arg0, char* arg1);
extern void _ZN10QPictureIOC1EP9QIODevicePKc(void* qthis, void* arg0, char* arg1);
  // proto:  const char * QPictureIO::format();
extern void _ZNK10QPictureIO6formatEv(void* qthis);
  // proto:  void QPictureIO::setQuality(int );
extern void _ZN10QPictureIO10setQualityEi(void* qthis, int arg0);
  // proto:  const QPicture & QPictureIO::picture();
extern void _ZNK10QPictureIO7pictureEv(void* qthis);
  // proto:  void QPictureIO::setFormat(const char * );
extern void _ZN10QPictureIO9setFormatEPKc(void* qthis, char* arg0);
  // proto:  void QPictureIO::setDescription(const QString & );
extern void _ZN10QPictureIO14setDescriptionERK7QString(void* qthis, void* arg0);
  // proto: static QByteArray QPictureIO::pictureFormat(const QString & fileName);
extern void _ZN10QPictureIO13pictureFormatERK7QString(void* arg0);
  // proto:  void QPictureIO::setIODevice(QIODevice * );
extern void _ZN10QPictureIO11setIODeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QPictureIO::setStatus(int );
extern void _ZN10QPictureIO9setStatusEi(void* qthis, int arg0);
  // proto:  QIODevice * QPictureIO::ioDevice();
extern void _ZNK10QPictureIO8ioDeviceEv(void* qthis);
  // proto:  float QPictureIO::gamma();
extern void _ZNK10QPictureIO5gammaEv(void* qthis);
  // proto: static QList<QByteArray> QPictureIO::outputFormats();
extern void _ZN10QPictureIO13outputFormatsEv();
  // proto:  void QPictureIO::setPicture(const QPicture & );
extern void _ZN10QPictureIO10setPictureERK8QPicture(void* qthis, void* arg0);
  // proto:  void QPictureIO::setParameters(const char * );
extern void _ZN10QPictureIO13setParametersEPKc(void* qthis, char* arg0);
  // proto:  void QPictureIO::QPictureIO();
extern void* dector_ZN10QPictureIOC1Ev();
extern void _ZN10QPictureIOC1Ev(void* qthis);
  // proto:  const char * QPicture::data();
extern void _ZNK8QPicture4dataEv(void* qthis);
  // proto: static QStringList QPicture::inputFormatList();
extern void _ZN8QPicture15inputFormatListEv();
  // proto:  void QPicture::swap(QPicture & other);
extern void demth_ZN8QPicture4swapERS_(void* qthis, void* arg0);
  // proto:  uint QPicture::size();
extern void _ZNK8QPicture4sizeEv(void* qthis);
  // proto:  bool QPicture::isNull();
extern void _ZNK8QPicture6isNullEv(void* qthis);
  // proto:  bool QPicture::save(QIODevice * dev, const char * format);
extern void _ZN8QPicture4saveEP9QIODevicePKc(void* qthis, void* arg0, char* arg1);
  // proto:  void QPicture::detach();
extern void _ZN8QPicture6detachEv(void* qthis);
  // proto: static QList<QByteArray> QPicture::inputFormats();
extern void _ZN8QPicture12inputFormatsEv();
  // proto:  void QPicture::QPicture(int formatVersion);
extern void* dector_ZN8QPictureC1Ei(int arg0);
extern void _ZN8QPictureC1Ei(void* qthis, int arg0);
  // proto:  void QPicture::QPicture(const QPicture & );
extern void* dector_ZN8QPictureC1ERKS_(void* arg0);
extern void _ZN8QPictureC1ERKS_(void* qthis, void* arg0);
  // proto:  bool QPicture::isDetached();
extern void _ZNK8QPicture10isDetachedEv(void* qthis);
  // proto: static QStringList QPicture::outputFormatList();
extern void _ZN8QPicture16outputFormatListEv();
  // proto:  void QPicture::setData(const char * data, uint size);
extern void _ZN8QPicture7setDataEPKcj(void* qthis, char* arg0, unsigned int arg1);
  // proto: static QList<QByteArray> QPicture::outputFormats();
extern void _ZN8QPicture13outputFormatsEv();
  // proto:  int QPicture::devType();
extern void _ZNK8QPicture7devTypeEv(void* qthis);
  // proto: static const char * QPicture::pictureFormat(const QString & fileName);
extern void _ZN8QPicture13pictureFormatERK7QString(void* arg0);
  // proto:  bool QPicture::save(const QString & fileName, const char * format);
extern void _ZN8QPicture4saveERK7QStringPKc(void* qthis, void* arg0, char* arg1);
  // proto:  bool QPicture::load(const QString & fileName, const char * format);
extern void _ZN8QPicture4loadERK7QStringPKc(void* qthis, void* arg0, char* arg1);
  // proto:  void QPicture::~QPicture();
extern void _ZN8QPictureD0Ev(void* qthis);
  // proto:  void QPicture::setBoundingRect(const QRect & r);
extern void _ZN8QPicture15setBoundingRectERK5QRect(void* qthis, void* arg0);
  // proto:  bool QPicture::load(QIODevice * dev, const char * format);
extern void _ZN8QPicture4loadEP9QIODevicePKc(void* qthis, void* arg0, char* arg1);
  // proto:  QRect QPicture::boundingRect();
extern void _ZNK8QPicture12boundingRectEv(void* qthis);
  // proto:  bool QPicture::play(QPainter * p);
extern void _ZN8QPicture4playEP8QPainter(void* qthis, void* arg0);
  // proto:  QPaintEngine * QPicture::paintEngine();
extern void _ZNK8QPicture11paintEngineEv(void* qthis);
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
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QPicture)=1
type QPicture struct {
  /*qbase*/ QPaintDevice;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QPictureIO::QPictureIO(const QString & fileName, const char * format);
func NewQPictureIO(args ...interface{}) QPictureIO {
  return QPictureIO{}
}

  // proto:  QString QPictureIO::description();
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
    C._ZNK10QPictureIO11descriptionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "description", args)
  }

}

  // proto: static QList<QByteArray> QPictureIO::inputFormats();
func (this *QPictureIO) inputFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPictureIO", "inputFormats", args)
  }

}

  // proto:  void QPictureIO::setGamma(float );
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
    C._ZN10QPictureIO8setGammaEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setGamma", args)
  }

}

  // proto:  int QPictureIO::status();
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
    C._ZNK10QPictureIO6statusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "status", args)
  }

}

  // proto:  int QPictureIO::quality();
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
    C._ZNK10QPictureIO7qualityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "quality", args)
  }

}

  // proto:  bool QPictureIO::write();
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
    C._ZN10QPictureIO5writeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "write", args)
  }

}

  // proto:  void QPictureIO::setFileName(const QString & );
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
    C._ZN10QPictureIO11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setFileName", args)
  }

}

  // proto:  void QPictureIO::~QPictureIO();
func (this *QPictureIO) FreeQPictureIO(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPictureIO", "~QPictureIO", args)
  }

}

  // proto:  const char * QPictureIO::parameters();
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
    C._ZNK10QPictureIO10parametersEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "parameters", args)
  }

}

  // proto: static QByteArray QPictureIO::pictureFormat(QIODevice * );
func (this *QPictureIO) pictureFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPictureIO", "pictureFormat", args)
  }

}

  // proto:  bool QPictureIO::read();
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
    C._ZN10QPictureIO4readEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "read", args)
  }

}

  // proto:  QString QPictureIO::fileName();
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
    C._ZNK10QPictureIO8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "fileName", args)
  }

}

  // proto:  const char * QPictureIO::format();
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
    C._ZNK10QPictureIO6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "format", args)
  }

}

  // proto:  void QPictureIO::setQuality(int );
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
    C._ZN10QPictureIO10setQualityEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setQuality", args)
  }

}

  // proto:  const QPicture & QPictureIO::picture();
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
    C._ZNK10QPictureIO7pictureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "picture", args)
  }

}

  // proto:  void QPictureIO::setFormat(const char * );
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    C._ZN10QPictureIO9setFormatEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setFormat", args)
  }

}

  // proto:  void QPictureIO::setDescription(const QString & );
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
    C._ZN10QPictureIO14setDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setDescription", args)
  }

}

  // proto:  void QPictureIO::setIODevice(QIODevice * );
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
    C._ZN10QPictureIO11setIODeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setIODevice", args)
  }

}

  // proto:  void QPictureIO::setStatus(int );
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
    C._ZN10QPictureIO9setStatusEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setStatus", args)
  }

}

  // proto:  QIODevice * QPictureIO::ioDevice();
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
    C._ZNK10QPictureIO8ioDeviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "ioDevice", args)
  }

}

  // proto:  float QPictureIO::gamma();
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
    C._ZNK10QPictureIO5gammaEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPictureIO", "gamma", args)
  }

}

  // proto: static QList<QByteArray> QPictureIO::outputFormats();
func (this *QPictureIO) outputFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPictureIO", "outputFormats", args)
  }

}

  // proto:  void QPictureIO::setPicture(const QPicture & );
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
    C._ZN10QPictureIO10setPictureERK8QPicture(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setPicture", args)
  }

}

  // proto:  void QPictureIO::setParameters(const char * );
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    C._ZN10QPictureIO13setParametersEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPictureIO", "setParameters", args)
  }

}

  // proto:  const char * QPicture::data();
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
    C._ZNK8QPicture4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "data", args)
  }

}

  // proto: static QStringList QPicture::inputFormatList();
func (this *QPicture) inputFormatList_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPicture", "inputFormatList", args)
  }

}

  // proto:  void QPicture::swap(QPicture & other);
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
    C.demth_ZN8QPicture4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "swap", args)
  }

}

  // proto:  uint QPicture::size();
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
    C._ZNK8QPicture4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "size", args)
  }

}

  // proto:  bool QPicture::isNull();
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
    C._ZNK8QPicture6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "isNull", args)
  }

}

  // proto:  bool QPicture::save(QIODevice * dev, const char * format);
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
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN8QPicture4saveEP9QIODevicePKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPicture4saveERK7QStringPKc
    // invoke: bool save(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN8QPicture4saveERK7QStringPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "save", args)
  }

}

  // proto:  void QPicture::detach();
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
    C._ZN8QPicture6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "detach", args)
  }

}

  // proto: static QList<QByteArray> QPicture::inputFormats();
func (this *QPicture) inputFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPicture", "inputFormats", args)
  }

}

  // proto:  void QPicture::QPicture(int formatVersion);
func NewQPicture(args ...interface{}) QPicture {
  return QPicture{}
}

  // proto:  bool QPicture::isDetached();
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
    C._ZNK8QPicture10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "isDetached", args)
  }

}

  // proto: static QStringList QPicture::outputFormatList();
func (this *QPicture) outputFormatList_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPicture", "outputFormatList", args)
  }

}

  // proto:  void QPicture::setData(const char * data, uint size);
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
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN8QPicture7setDataEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "setData", args)
  }

}

  // proto: static QList<QByteArray> QPicture::outputFormats();
func (this *QPicture) outputFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPicture", "outputFormats", args)
  }

}

  // proto:  int QPicture::devType();
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
    C._ZNK8QPicture7devTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "devType", args)
  }

}

  // proto: static const char * QPicture::pictureFormat(const QString & fileName);
func (this *QPicture) pictureFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPicture", "pictureFormat", args)
  }

}

  // proto:  bool QPicture::load(const QString & fileName, const char * format);
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
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN8QPicture4loadERK7QStringPKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN8QPicture4loadEP9QIODevicePKc
    // invoke: bool load(class QIODevice *, const char *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.CString(args[1].(string))
    if false {fmt.Println(arg1)}
    C._ZN8QPicture4loadEP9QIODevicePKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QPicture", "load", args)
  }

}

  // proto:  void QPicture::~QPicture();
func (this *QPicture) FreeQPicture(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QPicture", "~QPicture", args)
  }

}

  // proto:  void QPicture::setBoundingRect(const QRect & r);
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
    C._ZN8QPicture15setBoundingRectERK5QRect(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "setBoundingRect", args)
  }

}

  // proto:  QRect QPicture::boundingRect();
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
    C._ZNK8QPicture12boundingRectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "boundingRect", args)
  }

}

  // proto:  bool QPicture::play(QPainter * p);
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
    C._ZN8QPicture4playEP8QPainter(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QPicture", "play", args)
  }

}

  // proto:  QPaintEngine * QPicture::paintEngine();
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
    C._ZNK8QPicture11paintEngineEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QPicture", "paintEngine", args)
  }

}

// <= body block end

