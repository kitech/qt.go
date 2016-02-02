package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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
extern bool C_ZNK12QImageWriter8canWriteEv(void* qthis); // 4
  // proto: static QList<QByteArray> QImageWriter::supportedMimeTypes();
extern void C_ZN12QImageWriter18supportedMimeTypesEv(); // 4
  // proto:  QByteArray QImageWriter::subType();
extern void* C_ZNK12QImageWriter7subTypeEv(void* qthis); // 4
  // proto:  void QImageWriter::setOptimizedWrite(bool optimize);
extern void C_ZN12QImageWriter17setOptimizedWriteEb(void* qthis, bool arg0); // 4
  // proto:  void QImageWriter::QImageWriter(const QString & fileName, const QByteArray & format);
extern void* C_ZN12QImageWriterC2ERK7QStringRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QImageWriter::QImageWriter(QIODevice * device, const QByteArray & format);
extern void* C_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QImageWriter::QImageWriter();
extern void* C_ZN12QImageWriterC2Ev(); // 3
  // proto:  int QImageWriter::quality();
extern int32_t C_ZNK12QImageWriter7qualityEv(void* qthis); // 4
  // proto:  void QImageWriter::setDescription(const QString & description);
extern void C_ZN12QImageWriter14setDescriptionERK7QString(void* qthis, void* arg0); // 4
  // proto:  int QImageWriter::compression();
extern int32_t C_ZNK12QImageWriter11compressionEv(void* qthis); // 4
  // proto:  void QImageWriter::setFormat(const QByteArray & format);
extern void C_ZN12QImageWriter9setFormatERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  bool QImageWriter::write(const QImage & image);
extern bool C_ZN12QImageWriter5writeERK6QImage(void* qthis, void* arg0); // 4
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
extern void* C_ZNK12QImageWriter11descriptionEv(void* qthis); // 4
  // proto:  QString QImageWriter::errorString();
extern void* C_ZNK12QImageWriter11errorStringEv(void* qthis); // 4
  // proto:  QByteArray QImageWriter::format();
extern void* C_ZNK12QImageWriter6formatEv(void* qthis); // 4
  // proto:  QImageWriter::ImageWriterError QImageWriter::error();
extern void C_ZNK12QImageWriter5errorEv(void* qthis); // 4
  // proto:  QString QImageWriter::fileName();
extern void* C_ZNK12QImageWriter8fileNameEv(void* qthis); // 4
  // proto:  void QImageWriter::setGamma(float gamma);
extern void C_ZN12QImageWriter8setGammaEf(void* qthis, float arg0); // 4
  // proto:  QIODevice * QImageWriter::device();
extern void* C_ZNK12QImageWriter6deviceEv(void* qthis); // 4
  // proto:  QList<QByteArray> QImageWriter::supportedSubTypes();
extern void C_ZNK12QImageWriter17supportedSubTypesEv(void* qthis); // 4
  // proto:  bool QImageWriter::optimizedWrite();
extern bool C_ZNK12QImageWriter14optimizedWriteEv(void* qthis); // 4
  // proto:  bool QImageWriter::progressiveScanWrite();
extern bool C_ZNK12QImageWriter20progressiveScanWriteEv(void* qthis); // 4
  // proto:  void QImageWriter::setText(const QString & key, const QString & text);
extern void C_ZN12QImageWriter7setTextERK7QStringS2_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QImageWriter::setFileName(const QString & fileName);
extern void C_ZN12QImageWriter11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QImageWriter::setQuality(int quality);
extern void C_ZN12QImageWriter10setQualityEi(void* qthis, int32_t arg0); // 4
  // proto:  void QImageWriter::~QImageWriter();
extern void C_ZN12QImageWriterD2Ev(void* qthis); // 4
  // proto:  float QImageWriter::gamma();
extern float C_ZNK12QImageWriter5gammaEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// setCompression(int)
func (this *QImageWriter) Setcompression(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter14setCompressionEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setCompression", args)
  }

  return
}

// canWrite()
func (this *QImageWriter) Canwrite(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter8canWriteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "canWrite", args)
  }

  return
}

// supportedMimeTypes()
func (this *QImageWriter) Supportedmimetypes_S(args ...interface{}) () {
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

  return
}

// subType()
func (this *QImageWriter) Subtype(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter7subTypeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "subType", args)
  }

  return
}

// setOptimizedWrite(_Bool)
func (this *QImageWriter) Setoptimizedwrite(args ...interface{}) () {
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
    C.C_ZN12QImageWriter17setOptimizedWriteEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setOptimizedWrite", args)
  }

  return
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageWriterC2ERK7QStringRK10QByteArray(arg0, arg1)
    return &QImageWriter{Qclsinst:qthis}
  case 1:
    // invoke: _ZN12QImageWriterC1EP9QIODeviceRK10QByteArray
    // invoke: void QImageWriter(class QIODevice *, const class QByteArray &)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageWriterC2EP9QIODeviceRK10QByteArray(arg0, arg1)
    return &QImageWriter{Qclsinst:qthis}
  case 2:
    // invoke: _ZN12QImageWriterC1Ev
    // invoke: void QImageWriter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QImageWriterC2Ev()
    return &QImageWriter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QImageWriter", "QImageWriter", args)
  }

  return nil // QImageWriter{Qclsinst:qthis}
}

// quality()
func (this *QImageWriter) Quality(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter7qualityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "quality", args)
  }

  return
}

// setDescription(const class QString &)
func (this *QImageWriter) Setdescription(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter14setDescriptionERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setDescription", args)
  }

  return
}

// compression()
func (this *QImageWriter) Compression(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter11compressionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "compression", args)
  }

  return
}

// setFormat(const class QByteArray &)
func (this *QImageWriter) Setformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter9setFormatERK10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setFormat", args)
  }

  return
}

// write(const class QImage &)
func (this *QImageWriter) Write(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QImage).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QImageWriter5writeERK6QImage(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "write", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QImageWriter) Setdevice(args ...interface{}) () {
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
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter9setDeviceEP9QIODevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setDevice", args)
  }

  return
}

// setSubType(const class QByteArray &)
func (this *QImageWriter) Setsubtype(args ...interface{}) () {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter10setSubTypeERK10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setSubType", args)
  }

  return
}

// transformation()
func (this *QImageWriter) Transformation(args ...interface{}) () {
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
    C.C_ZNK12QImageWriter14transformationEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "transformation", args)
  }

  return
}

// supportedImageFormats()
func (this *QImageWriter) Supportedimageformats_S(args ...interface{}) () {
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

  return
}

// setProgressiveScanWrite(_Bool)
func (this *QImageWriter) Setprogressivescanwrite(args ...interface{}) () {
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
    C.C_ZN12QImageWriter23setProgressiveScanWriteEb(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setProgressiveScanWrite", args)
  }

  return
}

// description()
func (this *QImageWriter) Description(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter11descriptionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "description", args)
  }

  return
}

// errorString()
func (this *QImageWriter) Errorstring(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter11errorStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "errorString", args)
  }

  return
}

// format()
func (this *QImageWriter) Format(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "format", args)
  }

  return
}

// error()
func (this *QImageWriter) Error(args ...interface{}) () {
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
    C.C_ZNK12QImageWriter5errorEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "error", args)
  }

  return
}

// fileName()
func (this *QImageWriter) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "fileName", args)
  }

  return
}

// setGamma(float)
func (this *QImageWriter) Setgamma(args ...interface{}) () {
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
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter8setGammaEf(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setGamma", args)
  }

  return
}

// device()
func (this *QImageWriter) Device(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "device", args)
  }

  return
}

// supportedSubTypes()
func (this *QImageWriter) Supportedsubtypes(args ...interface{}) () {
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
    C.C_ZNK12QImageWriter17supportedSubTypesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "supportedSubTypes", args)
  }

  return
}

// optimizedWrite()
func (this *QImageWriter) Optimizedwrite(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter14optimizedWriteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "optimizedWrite", args)
  }

  return
}

// progressiveScanWrite()
func (this *QImageWriter) Progressivescanwrite(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter20progressiveScanWriteEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "progressiveScanWrite", args)
  }

  return
}

// setText(const class QString &, const class QString &)
func (this *QImageWriter) Settext(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN12QImageWriter7setTextERK7QStringS2_(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QImageWriter", "setText", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QImageWriter) Setfilename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setFileName", args)
  }

  return
}

// setQuality(int)
func (this *QImageWriter) Setquality(args ...interface{}) () {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN12QImageWriter10setQualityEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QImageWriter", "setQuality", args)
  }

  return
}

// ~QImageWriter()
func (this *QImageWriter) Freeqimagewriter(args ...interface{}) () {
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
    C.C_ZN12QImageWriterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QImageWriter", "~QImageWriter", args)
  }

  return
}

// gamma()
func (this *QImageWriter) Gamma(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QImageWriter5gammaEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QImageWriter", "gamma", args)
  }

  return
}

// <= body block end

