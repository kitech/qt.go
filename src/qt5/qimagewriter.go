package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  void QImageWriter::setCompression(int compression);
extern void C_ZN12QImageWriter14setCompressionEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QImageWriter::canWrite();
extern void C_ZNK12QImageWriter8canWriteEv(void* qthis); // 4
  // proto: static QList<QByteArray> QImageWriter::supportedMimeTypes();
extern void C_ZN12QImageWriter18supportedMimeTypesEv(); // 4
  // proto:  QByteArray QImageWriter::subType();
extern void C_ZNK12QImageWriter7subTypeEv(void* qthis); // 4
  // proto:  void QImageWriter::setOptimizedWrite(bool optimize);
extern void C_ZN12QImageWriter17setOptimizedWriteEb(void* qthis, bool arg0); // 4
  // proto:  void QImageWriter::QImageWriter(const QString & fileName, const QByteArray & format);
extern void* C_ZN12QImageWriterC2ERK7QStringRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QImageWriter::QImageWriter(QIODevice * device, const QByteArray & format);
extern void* C_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QImageWriter::QImageWriter();
extern void* C_ZN12QImageWriterC2Ev(); // 3
  // proto:  int QImageWriter::quality();
extern void C_ZNK12QImageWriter7qualityEv(void* qthis); // 4
  // proto:  void QImageWriter::setDescription(const QString & description);
extern void C_ZN12QImageWriter14setDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QImageWriter::compression();
extern void C_ZNK12QImageWriter11compressionEv(void* qthis); // 4
  // proto:  void QImageWriter::setFormat(const QByteArray & format);
extern void C_ZN12QImageWriter9setFormatERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  bool QImageWriter::write(const QImage & image);
extern void C_ZN12QImageWriter5writeERK6QImage(void* qthis, void* arg0); // 4
  // proto:  void QImageWriter::setDevice(QIODevice * device);
extern void C_ZN12QImageWriter9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  void QImageWriter::setSubType(const QByteArray & type);
extern void C_ZN12QImageWriter10setSubTypeERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QImageIOHandler::Transformations QImageWriter::transformation();
extern void C_ZNK12QImageWriter14transformationEv(void* qthis); // 4
  // proto: static QList<QByteArray> QImageWriter::supportedImageFormats();
extern void C_ZN12QImageWriter21supportedImageFormatsEv(); // 4
  // proto:  void QImageWriter::setProgressiveScanWrite(bool progressive);
extern void C_ZN12QImageWriter23setProgressiveScanWriteEb(void* qthis, bool arg0); // 4
  // proto:  QString QImageWriter::description();
extern void C_ZNK12QImageWriter11descriptionEv(void* qthis); // 4
  // proto:  QString QImageWriter::errorString();
extern void C_ZNK12QImageWriter11errorStringEv(void* qthis); // 4
  // proto:  QByteArray QImageWriter::format();
extern void C_ZNK12QImageWriter6formatEv(void* qthis); // 4
  // proto:  QImageWriter::ImageWriterError QImageWriter::error();
extern void C_ZNK12QImageWriter5errorEv(void* qthis); // 4
  // proto:  QString QImageWriter::fileName();
extern void C_ZNK12QImageWriter8fileNameEv(void* qthis); // 4
  // proto:  void QImageWriter::setGamma(float gamma);
extern void C_ZN12QImageWriter8setGammaEf(void* qthis, float arg0); // 4
  // proto:  QIODevice * QImageWriter::device();
extern void C_ZNK12QImageWriter6deviceEv(void* qthis); // 4
  // proto:  QList<QByteArray> QImageWriter::supportedSubTypes();
extern void C_ZNK12QImageWriter17supportedSubTypesEv(void* qthis); // 4
  // proto:  bool QImageWriter::optimizedWrite();
extern void C_ZNK12QImageWriter14optimizedWriteEv(void* qthis); // 4
  // proto:  bool QImageWriter::progressiveScanWrite();
extern void C_ZNK12QImageWriter20progressiveScanWriteEv(void* qthis); // 4
  // proto:  void QImageWriter::setText(const QString & key, const QString & text);
extern void C_ZN12QImageWriter7setTextERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QImageWriter::setFileName(const QString & fileName);
extern void C_ZN12QImageWriter11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QImageWriter::setQuality(int quality);
extern void C_ZN12QImageWriter10setQualityEi(void* qthis, int32_t arg0); // 4
  // proto:  void QImageWriter::~QImageWriter();
extern void C_ZN12QImageWriterD2Ev(void* qthis); // 4
  // proto:  float QImageWriter::gamma();
extern void C_ZNK12QImageWriter5gammaEv(void* qthis); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// setCompression(int)
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
    // invoke: void setCompression(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter14setCompressionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setCompression", args)
  }

}

// canWrite()
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
    // invoke: bool canWrite()
    var ret = C.C_ZNK12QImageWriter8canWriteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "canWrite", args)
  }

}

// supportedMimeTypes()
func (this *QImageWriter) supportedMimeTypes_s(args ...interface{}) () {
  // supportedMimeTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter18supportedMimeTypesEv
    // invoke: QList<QByteArray> supportedMimeTypes()
    C.C_ZN12QImageWriter18supportedMimeTypesEv()
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedMimeTypes", args)
  }

}

// subType()
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
    // invoke: QByteArray subType()
    var ret = C.C_ZNK12QImageWriter7subTypeEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "subType", args)
  }

}

// setOptimizedWrite(_Bool)
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
    // invoke: void setOptimizedWrite(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter17setOptimizedWriteEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setOptimizedWrite", args)
  }

}

// QImageWriter(const class QString &, const class QByteArray &)
func NewQImageWriter(args ...interface{}) *QImageWriter {
  // QImageWriter(const class QString &, const class QByteArray &)
  // QImageWriter(class QIODevice *, const class QByteArray &)
  // QImageWriter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriterC1ERK7QStringRK10QByteArray
    // invoke: void QImageWriter(const class QString &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageWriterC2ERK7QStringRK10QByteArray(arg0, arg1)
    return &QImageWriter{qclsinst:qthis}
  case 1:
    // invoke: _ZN12QImageWriterC1EP9QIODeviceRK10QByteArray
    // invoke: void QImageWriter(class QIODevice *, const class QByteArray &)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray(arg0, arg1)
    return &QImageWriter{qclsinst:qthis}
  case 2:
    // invoke: _ZN12QImageWriterC1Ev
    // invoke: void QImageWriter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageWriterC2Ev()
    return &QImageWriter{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QImageWriter", "QImageWriter", args)
  }

  return nil // QImageWriter{qclsinst:qthis}
}

// quality()
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
    // invoke: int quality()
    var ret = C.C_ZNK12QImageWriter7qualityEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "quality", args)
  }

}

// setDescription(const class QString &)
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
    // invoke: void setDescription(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter14setDescriptionERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setDescription", args)
  }

}

// compression()
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
    // invoke: int compression()
    var ret = C.C_ZNK12QImageWriter11compressionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "compression", args)
  }

}

// setFormat(const class QByteArray &)
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
    // invoke: void setFormat(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter9setFormatERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setFormat", args)
  }

}

// write(const class QImage &)
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
    // invoke: bool write(const class QImage &)
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZN12QImageWriter5writeERK6QImage(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "write", args)
  }

}

// setDevice(class QIODevice *)
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
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setDevice", args)
  }

}

// setSubType(const class QByteArray &)
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
    // invoke: void setSubType(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter10setSubTypeERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setSubType", args)
  }

}

// transformation()
func (this *QImageWriter) transformation(args ...interface{}) () {
  // transformation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter14transformationEv
    // invoke: QImageIOHandler::Transformations transformation()
    C.C_ZNK12QImageWriter14transformationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "transformation", args)
  }

}

// supportedImageFormats()
func (this *QImageWriter) supportedImageFormats_s(args ...interface{}) () {
  // supportedImageFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriter21supportedImageFormatsEv
    // invoke: QList<QByteArray> supportedImageFormats()
    C.C_ZN12QImageWriter21supportedImageFormatsEv()
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedImageFormats", args)
  }

}

// setProgressiveScanWrite(_Bool)
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
    // invoke: void setProgressiveScanWrite(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter23setProgressiveScanWriteEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setProgressiveScanWrite", args)
  }

}

// description()
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
    // invoke: QString description()
    var ret = C.C_ZNK12QImageWriter11descriptionEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "description", args)
  }

}

// errorString()
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
    // invoke: QString errorString()
    var ret = C.C_ZNK12QImageWriter11errorStringEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "errorString", args)
  }

}

// format()
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
    // invoke: QByteArray format()
    var ret = C.C_ZNK12QImageWriter6formatEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "format", args)
  }

}

// error()
func (this *QImageWriter) error(args ...interface{}) () {
  // error()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageWriter5errorEv
    // invoke: QImageWriter::ImageWriterError error()
    C.C_ZNK12QImageWriter5errorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "error", args)
  }

}

// fileName()
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
    // invoke: QString fileName()
    var ret = C.C_ZNK12QImageWriter8fileNameEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "fileName", args)
  }

}

// setGamma(float)
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
    // invoke: void setGamma(float)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter8setGammaEf(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setGamma", args)
  }

}

// device()
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
    // invoke: QIODevice * device()
    var ret = C.C_ZNK12QImageWriter6deviceEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "device", args)
  }

}

// supportedSubTypes()
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
    // invoke: QList<QByteArray> supportedSubTypes()
    C.C_ZNK12QImageWriter17supportedSubTypesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedSubTypes", args)
  }

}

// optimizedWrite()
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
    // invoke: bool optimizedWrite()
    var ret = C.C_ZNK12QImageWriter14optimizedWriteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "optimizedWrite", args)
  }

}

// progressiveScanWrite()
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
    // invoke: bool progressiveScanWrite()
    var ret = C.C_ZNK12QImageWriter20progressiveScanWriteEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "progressiveScanWrite", args)
  }

}

// setText(const class QString &, const class QString &)
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
    // invoke: void setText(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QImageWriter7setTextERK7QStringS2_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImageWriter", "setText", args)
  }

}

// setFileName(const class QString &)
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
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setFileName", args)
  }

}

// setQuality(int)
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
    // invoke: void setQuality(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter10setQualityEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setQuality", args)
  }

}

// ~QImageWriter()
func (this *QImageWriter) FreeQImageWriter(args ...interface{}) () {
  // ~QImageWriter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageWriterD0Ev
    // invoke: void ~QImageWriter()
    C.C_ZN12QImageWriterD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "~QImageWriter", args)
  }

}

// gamma()
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
    // invoke: float gamma()
    var ret = C.C_ZNK12QImageWriter5gammaEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QImageWriter", "gamma", args)
  }

}

// <= body block end

