package qt5
// auto generated, do not modify.
// created: Sat Jan  2 12:23:25 2016
// src-file: /QtGui/qimagewriter.h
// dst-file: /src/gui/qimagewriter.go
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
  // proto:  void QImageWriter::setText(const QString & key, const QString & text);
extern void _ZN12QImageWriter7setTextERK7QStringS2_(void* qthis, void* arg0, void* arg1);
  // proto:  void QImageWriter::setGamma(float gamma);
extern void _ZN12QImageWriter8setGammaEf(void* qthis, float arg0);
  // proto:  void QImageWriter::setFileName(const QString & fileName);
extern void _ZN12QImageWriter11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  bool QImageWriter::optimizedWrite();
extern void _ZNK12QImageWriter14optimizedWriteEv(void* qthis);
  // proto:  void QImageWriter::~QImageWriter();
extern void _ZN12QImageWriterD0Ev(void* qthis);
  // proto:  QIODevice * QImageWriter::device();
extern void _ZNK12QImageWriter6deviceEv(void* qthis);
  // proto:  QByteArray QImageWriter::subType();
extern void _ZNK12QImageWriter7subTypeEv(void* qthis);
  // proto: static QList<QByteArray> QImageWriter::supportedMimeTypes();
extern void _ZN12QImageWriter18supportedMimeTypesEv();
  // proto:  int QImageWriter::quality();
extern void _ZNK12QImageWriter7qualityEv(void* qthis);
  // proto:  bool QImageWriter::write(const QImage & image);
extern void _ZN12QImageWriter5writeERK6QImage(void* qthis, void* arg0);
  // proto:  void QImageWriter::setCompression(int compression);
extern void _ZN12QImageWriter14setCompressionEi(void* qthis, int arg0);
  // proto: static QList<QByteArray> QImageWriter::supportedImageFormats();
extern void _ZN12QImageWriter21supportedImageFormatsEv();
  // proto:  QString QImageWriter::fileName();
extern void _ZNK12QImageWriter8fileNameEv(void* qthis);
  // proto:  void QImageWriter::setOptimizedWrite(bool optimize);
extern void _ZN12QImageWriter17setOptimizedWriteEb(void* qthis, bool arg0);
  // proto:  QString QImageWriter::errorString();
extern void _ZNK12QImageWriter11errorStringEv(void* qthis);
  // proto:  void QImageWriter::setQuality(int quality);
extern void _ZN12QImageWriter10setQualityEi(void* qthis, int arg0);
  // proto:  float QImageWriter::gamma();
extern void _ZNK12QImageWriter5gammaEv(void* qthis);
  // proto:  QString QImageWriter::description();
extern void _ZNK12QImageWriter11descriptionEv(void* qthis);
  // proto:  void QImageWriter::QImageWriter();
extern void* dector_ZN12QImageWriterC1Ev();
extern void _ZN12QImageWriterC1Ev(void* qthis);
  // proto:  void QImageWriter::setFormat(const QByteArray & format);
extern void _ZN12QImageWriter9setFormatERK10QByteArray(void* qthis, void* arg0);
  // proto:  void QImageWriter::QImageWriter(const QString & fileName, const QByteArray & format);
extern void* dector_ZN12QImageWriterC1ERK7QStringRK10QByteArray(void* arg0, void* arg1);
extern void _ZN12QImageWriterC1ERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  void QImageWriter::setDevice(QIODevice * device);
extern void _ZN12QImageWriter9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QImageWriter::setSubType(const QByteArray & type);
extern void _ZN12QImageWriter10setSubTypeERK10QByteArray(void* qthis, void* arg0);
  // proto:  bool QImageWriter::progressiveScanWrite();
extern void _ZNK12QImageWriter20progressiveScanWriteEv(void* qthis);
  // proto:  QByteArray QImageWriter::format();
extern void _ZNK12QImageWriter6formatEv(void* qthis);
  // proto:  QList<QByteArray> QImageWriter::supportedSubTypes();
extern void _ZNK12QImageWriter17supportedSubTypesEv(void* qthis);
  // proto:  bool QImageWriter::canWrite();
extern void _ZNK12QImageWriter8canWriteEv(void* qthis);
  // proto:  void QImageWriter::QImageWriter(QIODevice * device, const QByteArray & format);
extern void* dector_ZN12QImageWriterC1EP9QIODeviceRK10QByteArray(void* arg0, void* arg1);
extern void _ZN12QImageWriterC1EP9QIODeviceRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  int QImageWriter::compression();
extern void _ZNK12QImageWriter11compressionEv(void* qthis);
  // proto:  void QImageWriter::setProgressiveScanWrite(bool progressive);
extern void _ZN12QImageWriter23setProgressiveScanWriteEb(void* qthis, bool arg0);
  // proto:  void QImageWriter::setDescription(const QString & description);
extern void _ZN12QImageWriter14setDescriptionERK7QString(void* qthis, void* arg0);
  // proto:  void QImageWriter::QImageWriter(const QImageWriter & );
extern void* dector_ZN12QImageWriterC1ERKS_(void* arg0);
extern void _ZN12QImageWriterC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QImageWriter)=8
type QImageWriter struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  void QImageWriter::setText(const QString & key, const QString & text);
func (this *QImageWriter) setText(args ...interface{}) () {
  // setText(const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter7setTextERK7QStringS2_
  default:
    qtrt.ErrorResolve("QImageWriter", "setText", args)
  }

}

  // proto:  void QImageWriter::setGamma(float gamma);
func (this *QImageWriter) setGamma(args ...interface{}) () {
  // setGamma(float)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter8setGammaEf
  default:
    qtrt.ErrorResolve("QImageWriter", "setGamma", args)
  }

}

  // proto:  void QImageWriter::setFileName(const QString & fileName);
func (this *QImageWriter) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter11setFileNameERK7QString
  default:
    qtrt.ErrorResolve("QImageWriter", "setFileName", args)
  }

}

  // proto:  bool QImageWriter::optimizedWrite();
func (this *QImageWriter) optimizedWrite(args ...interface{}) () {
  // optimizedWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter14optimizedWriteEv
  default:
    qtrt.ErrorResolve("QImageWriter", "optimizedWrite", args)
  }

}

  // proto:  void QImageWriter::~QImageWriter();
func (this *QImageWriter) FreeQImageWriter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageWriter", "~QImageWriter", args)
  }

}

  // proto:  QIODevice * QImageWriter::device();
func (this *QImageWriter) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter6deviceEv
  default:
    qtrt.ErrorResolve("QImageWriter", "device", args)
  }

}

  // proto:  QByteArray QImageWriter::subType();
func (this *QImageWriter) subType(args ...interface{}) () {
  // subType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter7subTypeEv
  default:
    qtrt.ErrorResolve("QImageWriter", "subType", args)
  }

}

  // proto: static QList<QByteArray> QImageWriter::supportedMimeTypes();
func (this *QImageWriter) supportedMimeTypes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedMimeTypes", args)
  }

}

  // proto:  int QImageWriter::quality();
func (this *QImageWriter) quality(args ...interface{}) () {
  // quality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter7qualityEv
  default:
    qtrt.ErrorResolve("QImageWriter", "quality", args)
  }

}

  // proto:  bool QImageWriter::write(const QImage & image);
func (this *QImageWriter) write(args ...interface{}) () {
  // write(const class QImage &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "const QImage &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter5writeERK6QImage
  default:
    qtrt.ErrorResolve("QImageWriter", "write", args)
  }

}

  // proto:  void QImageWriter::setCompression(int compression);
func (this *QImageWriter) setCompression(args ...interface{}) () {
  // setCompression(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter14setCompressionEi
  default:
    qtrt.ErrorResolve("QImageWriter", "setCompression", args)
  }

}

  // proto: static QList<QByteArray> QImageWriter::supportedImageFormats();
func (this *QImageWriter) supportedImageFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedImageFormats", args)
  }

}

  // proto:  QString QImageWriter::fileName();
func (this *QImageWriter) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter8fileNameEv
  default:
    qtrt.ErrorResolve("QImageWriter", "fileName", args)
  }

}

  // proto:  void QImageWriter::setOptimizedWrite(bool optimize);
func (this *QImageWriter) setOptimizedWrite(args ...interface{}) () {
  // setOptimizedWrite(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter17setOptimizedWriteEb
  default:
    qtrt.ErrorResolve("QImageWriter", "setOptimizedWrite", args)
  }

}

  // proto:  QString QImageWriter::errorString();
func (this *QImageWriter) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter11errorStringEv
  default:
    qtrt.ErrorResolve("QImageWriter", "errorString", args)
  }

}

  // proto:  void QImageWriter::setQuality(int quality);
func (this *QImageWriter) setQuality(args ...interface{}) () {
  // setQuality(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter10setQualityEi
  default:
    qtrt.ErrorResolve("QImageWriter", "setQuality", args)
  }

}

  // proto:  float QImageWriter::gamma();
func (this *QImageWriter) gamma(args ...interface{}) () {
  // gamma()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter5gammaEv
  default:
    qtrt.ErrorResolve("QImageWriter", "gamma", args)
  }

}

  // proto:  QString QImageWriter::description();
func (this *QImageWriter) description(args ...interface{}) () {
  // description()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter11descriptionEv
  default:
    qtrt.ErrorResolve("QImageWriter", "description", args)
  }

}

  // proto:  void QImageWriter::QImageWriter();
func NewQImageWriter(args ...interface{}) QImageWriter {
  return QImageWriter{}
}

  // proto:  void QImageWriter::setFormat(const QByteArray & format);
func (this *QImageWriter) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter9setFormatERK10QByteArray
  default:
    qtrt.ErrorResolve("QImageWriter", "setFormat", args)
  }

}

  // proto:  void QImageWriter::setDevice(QIODevice * device);
func (this *QImageWriter) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QImageWriter", "setDevice", args)
  }

}

  // proto:  void QImageWriter::setSubType(const QByteArray & type);
func (this *QImageWriter) setSubType(args ...interface{}) () {
  // setSubType(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter10setSubTypeERK10QByteArray
  default:
    qtrt.ErrorResolve("QImageWriter", "setSubType", args)
  }

}

  // proto:  bool QImageWriter::progressiveScanWrite();
func (this *QImageWriter) progressiveScanWrite(args ...interface{}) () {
  // progressiveScanWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter20progressiveScanWriteEv
  default:
    qtrt.ErrorResolve("QImageWriter", "progressiveScanWrite", args)
  }

}

  // proto:  QByteArray QImageWriter::format();
func (this *QImageWriter) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter6formatEv
  default:
    qtrt.ErrorResolve("QImageWriter", "format", args)
  }

}

  // proto:  QList<QByteArray> QImageWriter::supportedSubTypes();
func (this *QImageWriter) supportedSubTypes(args ...interface{}) () {
  // supportedSubTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter17supportedSubTypesEv
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedSubTypes", args)
  }

}

  // proto:  bool QImageWriter::canWrite();
func (this *QImageWriter) canWrite(args ...interface{}) () {
  // canWrite()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter8canWriteEv
  default:
    qtrt.ErrorResolve("QImageWriter", "canWrite", args)
  }

}

  // proto:  int QImageWriter::compression();
func (this *QImageWriter) compression(args ...interface{}) () {
  // compression()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter11compressionEv
  default:
    qtrt.ErrorResolve("QImageWriter", "compression", args)
  }

}

  // proto:  void QImageWriter::setProgressiveScanWrite(bool progressive);
func (this *QImageWriter) setProgressiveScanWrite(args ...interface{}) () {
  // setProgressiveScanWrite(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter23setProgressiveScanWriteEb
  default:
    qtrt.ErrorResolve("QImageWriter", "setProgressiveScanWrite", args)
  }

}

  // proto:  void QImageWriter::setDescription(const QString & description);
func (this *QImageWriter) setDescription(args ...interface{}) () {
  // setDescription(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter14setDescriptionERK7QString
  default:
    qtrt.ErrorResolve("QImageWriter", "setDescription", args)
  }

}

// <= body block end

