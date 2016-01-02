package qt5
// auto generated, do not modify.
// created: Sat Jan  2 16:11:29 2016
// src-file: /QtGui/qimagereader.h
// dst-file: /src/gui/qimagereader.go
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
  // proto:  QString QImageReader::errorString();
extern void _ZNK12QImageReader11errorStringEv(void* qthis);
  // proto:  bool QImageReader::canRead();
extern void _ZNK12QImageReader7canReadEv(void* qthis);
  // proto:  void QImageReader::~QImageReader();
extern void _ZN12QImageReaderD0Ev(void* qthis);
  // proto:  void QImageReader::setScaledSize(const QSize & size);
extern void _ZN12QImageReader13setScaledSizeERK5QSize(void* qthis, void* arg0);
  // proto:  bool QImageReader::read(QImage * image);
extern void _ZN12QImageReader4readEP6QImage(void* qthis, void* arg0);
  // proto:  void QImageReader::setScaledClipRect(const QRect & rect);
extern void _ZN12QImageReader17setScaledClipRectERK5QRect(void* qthis, void* arg0);
  // proto:  int QImageReader::imageCount();
extern void _ZNK12QImageReader10imageCountEv(void* qthis);
  // proto:  QStringList QImageReader::textKeys();
extern void _ZNK12QImageReader8textKeysEv(void* qthis);
  // proto:  bool QImageReader::decideFormatFromContent();
extern void _ZNK12QImageReader23decideFormatFromContentEv(void* qthis);
  // proto:  QIODevice * QImageReader::device();
extern void _ZNK12QImageReader6deviceEv(void* qthis);
  // proto:  bool QImageReader::autoTransform();
extern void _ZNK12QImageReader13autoTransformEv(void* qthis);
  // proto:  bool QImageReader::jumpToNextImage();
extern void _ZN12QImageReader15jumpToNextImageEv(void* qthis);
  // proto: static QByteArray QImageReader::imageFormat(const QString & fileName);
extern void _ZN12QImageReader11imageFormatERK7QString(void* arg0);
  // proto:  QList<QByteArray> QImageReader::supportedSubTypes();
extern void _ZNK12QImageReader17supportedSubTypesEv(void* qthis);
  // proto:  QSize QImageReader::size();
extern void _ZNK12QImageReader4sizeEv(void* qthis);
  // proto:  QColor QImageReader::backgroundColor();
extern void _ZNK12QImageReader15backgroundColorEv(void* qthis);
  // proto:  QByteArray QImageReader::subType();
extern void _ZNK12QImageReader7subTypeEv(void* qthis);
  // proto:  int QImageReader::currentImageNumber();
extern void _ZNK12QImageReader18currentImageNumberEv(void* qthis);
  // proto: static QList<QByteArray> QImageReader::supportedImageFormats();
extern void _ZN12QImageReader21supportedImageFormatsEv();
  // proto:  int QImageReader::loopCount();
extern void _ZNK12QImageReader9loopCountEv(void* qthis);
  // proto:  void QImageReader::setDecideFormatFromContent(bool ignored);
extern void _ZN12QImageReader26setDecideFormatFromContentEb(void* qthis, bool arg0);
  // proto:  QRect QImageReader::scaledClipRect();
extern void _ZNK12QImageReader14scaledClipRectEv(void* qthis);
  // proto: static QList<QByteArray> QImageReader::supportedMimeTypes();
extern void _ZN12QImageReader18supportedMimeTypesEv();
  // proto:  QString QImageReader::text(const QString & key);
extern void _ZNK12QImageReader4textERK7QString(void* qthis, void* arg0);
  // proto:  int QImageReader::nextImageDelay();
extern void _ZNK12QImageReader14nextImageDelayEv(void* qthis);
  // proto:  QImage QImageReader::read();
extern void _ZN12QImageReader4readEv(void* qthis);
  // proto:  bool QImageReader::supportsAnimation();
extern void _ZNK12QImageReader17supportsAnimationEv(void* qthis);
  // proto:  bool QImageReader::jumpToImage(int imageNumber);
extern void _ZN12QImageReader11jumpToImageEi(void* qthis, int arg0);
  // proto:  void QImageReader::setFileName(const QString & fileName);
extern void _ZN12QImageReader11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  void QImageReader::QImageReader(const QImageReader & );
extern void* dector_ZN12QImageReaderC1ERKS_(void* arg0);
extern void _ZN12QImageReaderC1ERKS_(void* qthis, void* arg0);
  // proto:  QSize QImageReader::scaledSize();
extern void _ZNK12QImageReader10scaledSizeEv(void* qthis);
  // proto:  void QImageReader::setAutoTransform(bool enabled);
extern void _ZN12QImageReader16setAutoTransformEb(void* qthis, bool arg0);
  // proto:  void QImageReader::setClipRect(const QRect & rect);
extern void _ZN12QImageReader11setClipRectERK5QRect(void* qthis, void* arg0);
  // proto:  bool QImageReader::autoDetectImageFormat();
extern void _ZNK12QImageReader21autoDetectImageFormatEv(void* qthis);
  // proto:  QRect QImageReader::currentImageRect();
extern void _ZNK12QImageReader16currentImageRectEv(void* qthis);
  // proto:  void QImageReader::QImageReader(const QString & fileName, const QByteArray & format);
extern void* dector_ZN12QImageReaderC1ERK7QStringRK10QByteArray(void* arg0, void* arg1);
extern void _ZN12QImageReaderC1ERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto: static QByteArray QImageReader::imageFormat(QIODevice * device);
extern void _ZN12QImageReader11imageFormatEP9QIODevice(void* arg0);
  // proto:  int QImageReader::quality();
extern void _ZNK12QImageReader7qualityEv(void* qthis);
  // proto:  void QImageReader::setDevice(QIODevice * device);
extern void _ZN12QImageReader9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QImageReader::setBackgroundColor(const QColor & color);
extern void _ZN12QImageReader18setBackgroundColorERK6QColor(void* qthis, void* arg0);
  // proto:  void QImageReader::setQuality(int quality);
extern void _ZN12QImageReader10setQualityEi(void* qthis, int arg0);
  // proto:  void QImageReader::QImageReader(QIODevice * device, const QByteArray & format);
extern void* dector_ZN12QImageReaderC1EP9QIODeviceRK10QByteArray(void* arg0, void* arg1);
extern void _ZN12QImageReaderC1EP9QIODeviceRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  void QImageReader::setAutoDetectImageFormat(bool enabled);
extern void _ZN12QImageReader24setAutoDetectImageFormatEb(void* qthis, bool arg0);
  // proto:  void QImageReader::QImageReader();
extern void* dector_ZN12QImageReaderC1Ev();
extern void _ZN12QImageReaderC1Ev(void* qthis);
  // proto:  void QImageReader::setFormat(const QByteArray & format);
extern void _ZN12QImageReader9setFormatERK10QByteArray(void* qthis, void* arg0);
  // proto:  QString QImageReader::fileName();
extern void _ZNK12QImageReader8fileNameEv(void* qthis);
  // proto:  QRect QImageReader::clipRect();
extern void _ZNK12QImageReader8clipRectEv(void* qthis);
  // proto:  QByteArray QImageReader::format();
extern void _ZNK12QImageReader6formatEv(void* qthis);
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

// class sizeof(QImageReader)=8
type QImageReader struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

  // proto:  QString QImageReader::errorString();
func (this *QImageReader) errorString(args ...interface{}) () {
  // errorString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader11errorStringEv
  default:
    qtrt.ErrorResolve("QImageReader", "errorString", args)
  }

}

  // proto:  bool QImageReader::canRead();
func (this *QImageReader) canRead(args ...interface{}) () {
  // canRead()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader7canReadEv
  default:
    qtrt.ErrorResolve("QImageReader", "canRead", args)
  }

}

  // proto:  void QImageReader::~QImageReader();
func (this *QImageReader) FreeQImageReader(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageReader", "~QImageReader", args)
  }

}

  // proto:  void QImageReader::setScaledSize(const QSize & size);
func (this *QImageReader) setScaledSize(args ...interface{}) () {
  // setScaledSize(const class QSize &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QSize{}) // "const QSize &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader13setScaledSizeERK5QSize
    var arg0 = args[0].(QSize).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setScaledSize", args)
  }

}

  // proto:  bool QImageReader::read(QImage * image);
func (this *QImageReader) read(args ...interface{}) () {
  // read(class QImage *)
  // read()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QImage{}) // "QImage *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader4readEP6QImage
    var arg0 = args[0].(QImage).qclsinst
    if false {fmt.Println(arg0)}
  case 1:
    // invoke: _ZN12QImageReader4readEv
  default:
    qtrt.ErrorResolve("QImageReader", "read", args)
  }

}

  // proto:  void QImageReader::setScaledClipRect(const QRect & rect);
func (this *QImageReader) setScaledClipRect(args ...interface{}) () {
  // setScaledClipRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader17setScaledClipRectERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setScaledClipRect", args)
  }

}

  // proto:  int QImageReader::imageCount();
func (this *QImageReader) imageCount(args ...interface{}) () {
  // imageCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader10imageCountEv
  default:
    qtrt.ErrorResolve("QImageReader", "imageCount", args)
  }

}

  // proto:  QStringList QImageReader::textKeys();
func (this *QImageReader) textKeys(args ...interface{}) () {
  // textKeys()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader8textKeysEv
  default:
    qtrt.ErrorResolve("QImageReader", "textKeys", args)
  }

}

  // proto:  bool QImageReader::decideFormatFromContent();
func (this *QImageReader) decideFormatFromContent(args ...interface{}) () {
  // decideFormatFromContent()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader23decideFormatFromContentEv
  default:
    qtrt.ErrorResolve("QImageReader", "decideFormatFromContent", args)
  }

}

  // proto:  QIODevice * QImageReader::device();
func (this *QImageReader) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader6deviceEv
  default:
    qtrt.ErrorResolve("QImageReader", "device", args)
  }

}

  // proto:  bool QImageReader::autoTransform();
func (this *QImageReader) autoTransform(args ...interface{}) () {
  // autoTransform()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader13autoTransformEv
  default:
    qtrt.ErrorResolve("QImageReader", "autoTransform", args)
  }

}

  // proto:  bool QImageReader::jumpToNextImage();
func (this *QImageReader) jumpToNextImage(args ...interface{}) () {
  // jumpToNextImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader15jumpToNextImageEv
  default:
    qtrt.ErrorResolve("QImageReader", "jumpToNextImage", args)
  }

}

  // proto: static QByteArray QImageReader::imageFormat(const QString & fileName);
func (this *QImageReader) imageFormat_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageReader", "imageFormat", args)
  }

}

  // proto:  QList<QByteArray> QImageReader::supportedSubTypes();
func (this *QImageReader) supportedSubTypes(args ...interface{}) () {
  // supportedSubTypes()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader17supportedSubTypesEv
  default:
    qtrt.ErrorResolve("QImageReader", "supportedSubTypes", args)
  }

}

  // proto:  QSize QImageReader::size();
func (this *QImageReader) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader4sizeEv
  default:
    qtrt.ErrorResolve("QImageReader", "size", args)
  }

}

  // proto:  QColor QImageReader::backgroundColor();
func (this *QImageReader) backgroundColor(args ...interface{}) () {
  // backgroundColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader15backgroundColorEv
  default:
    qtrt.ErrorResolve("QImageReader", "backgroundColor", args)
  }

}

  // proto:  QByteArray QImageReader::subType();
func (this *QImageReader) subType(args ...interface{}) () {
  // subType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader7subTypeEv
  default:
    qtrt.ErrorResolve("QImageReader", "subType", args)
  }

}

  // proto:  int QImageReader::currentImageNumber();
func (this *QImageReader) currentImageNumber(args ...interface{}) () {
  // currentImageNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader18currentImageNumberEv
  default:
    qtrt.ErrorResolve("QImageReader", "currentImageNumber", args)
  }

}

  // proto: static QList<QByteArray> QImageReader::supportedImageFormats();
func (this *QImageReader) supportedImageFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageReader", "supportedImageFormats", args)
  }

}

  // proto:  int QImageReader::loopCount();
func (this *QImageReader) loopCount(args ...interface{}) () {
  // loopCount()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader9loopCountEv
  default:
    qtrt.ErrorResolve("QImageReader", "loopCount", args)
  }

}

  // proto:  void QImageReader::setDecideFormatFromContent(bool ignored);
func (this *QImageReader) setDecideFormatFromContent(args ...interface{}) () {
  // setDecideFormatFromContent(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader26setDecideFormatFromContentEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setDecideFormatFromContent", args)
  }

}

  // proto:  QRect QImageReader::scaledClipRect();
func (this *QImageReader) scaledClipRect(args ...interface{}) () {
  // scaledClipRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader14scaledClipRectEv
  default:
    qtrt.ErrorResolve("QImageReader", "scaledClipRect", args)
  }

}

  // proto: static QList<QByteArray> QImageReader::supportedMimeTypes();
func (this *QImageReader) supportedMimeTypes_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QImageReader", "supportedMimeTypes", args)
  }

}

  // proto:  QString QImageReader::text(const QString & key);
func (this *QImageReader) text(args ...interface{}) () {
  // text(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader4textERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "text", args)
  }

}

  // proto:  int QImageReader::nextImageDelay();
func (this *QImageReader) nextImageDelay(args ...interface{}) () {
  // nextImageDelay()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader14nextImageDelayEv
  default:
    qtrt.ErrorResolve("QImageReader", "nextImageDelay", args)
  }

}

  // proto:  bool QImageReader::supportsAnimation();
func (this *QImageReader) supportsAnimation(args ...interface{}) () {
  // supportsAnimation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader17supportsAnimationEv
  default:
    qtrt.ErrorResolve("QImageReader", "supportsAnimation", args)
  }

}

  // proto:  bool QImageReader::jumpToImage(int imageNumber);
func (this *QImageReader) jumpToImage(args ...interface{}) () {
  // jumpToImage(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11jumpToImageEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "jumpToImage", args)
  }

}

  // proto:  void QImageReader::setFileName(const QString & fileName);
func (this *QImageReader) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11setFileNameERK7QString
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setFileName", args)
  }

}

  // proto:  void QImageReader::QImageReader(const QImageReader & );
func NewQImageReader(args ...interface{}) QImageReader {
  return QImageReader{}
}

  // proto:  QSize QImageReader::scaledSize();
func (this *QImageReader) scaledSize(args ...interface{}) () {
  // scaledSize()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader10scaledSizeEv
  default:
    qtrt.ErrorResolve("QImageReader", "scaledSize", args)
  }

}

  // proto:  void QImageReader::setAutoTransform(bool enabled);
func (this *QImageReader) setAutoTransform(args ...interface{}) () {
  // setAutoTransform(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader16setAutoTransformEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setAutoTransform", args)
  }

}

  // proto:  void QImageReader::setClipRect(const QRect & rect);
func (this *QImageReader) setClipRect(args ...interface{}) () {
  // setClipRect(const class QRect &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRect{}) // "const QRect &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader11setClipRectERK5QRect
    var arg0 = args[0].(QRect).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setClipRect", args)
  }

}

  // proto:  bool QImageReader::autoDetectImageFormat();
func (this *QImageReader) autoDetectImageFormat(args ...interface{}) () {
  // autoDetectImageFormat()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader21autoDetectImageFormatEv
  default:
    qtrt.ErrorResolve("QImageReader", "autoDetectImageFormat", args)
  }

}

  // proto:  QRect QImageReader::currentImageRect();
func (this *QImageReader) currentImageRect(args ...interface{}) () {
  // currentImageRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader16currentImageRectEv
  default:
    qtrt.ErrorResolve("QImageReader", "currentImageRect", args)
  }

}

  // proto:  int QImageReader::quality();
func (this *QImageReader) quality(args ...interface{}) () {
  // quality()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader7qualityEv
  default:
    qtrt.ErrorResolve("QImageReader", "quality", args)
  }

}

  // proto:  void QImageReader::setDevice(QIODevice * device);
func (this *QImageReader) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader9setDeviceEP9QIODevice
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setDevice", args)
  }

}

  // proto:  void QImageReader::setBackgroundColor(const QColor & color);
func (this *QImageReader) setBackgroundColor(args ...interface{}) () {
  // setBackgroundColor(const class QColor &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QColor{}) // "const QColor &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader18setBackgroundColorERK6QColor
    var arg0 = args[0].(QColor).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setBackgroundColor", args)
  }

}

  // proto:  void QImageReader::setQuality(int quality);
func (this *QImageReader) setQuality(args ...interface{}) () {
  // setQuality(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader10setQualityEi
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setQuality", args)
  }

}

  // proto:  void QImageReader::setAutoDetectImageFormat(bool enabled);
func (this *QImageReader) setAutoDetectImageFormat(args ...interface{}) () {
  // setAutoDetectImageFormat(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader24setAutoDetectImageFormatEb
    var arg0 = C.int8_t(args[0].(int8))
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setAutoDetectImageFormat", args)
  }

}

  // proto:  void QImageReader::setFormat(const QByteArray & format);
func (this *QImageReader) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QImageReader9setFormatERK10QByteArray
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
  default:
    qtrt.ErrorResolve("QImageReader", "setFormat", args)
  }

}

  // proto:  QString QImageReader::fileName();
func (this *QImageReader) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader8fileNameEv
  default:
    qtrt.ErrorResolve("QImageReader", "fileName", args)
  }

}

  // proto:  QRect QImageReader::clipRect();
func (this *QImageReader) clipRect(args ...interface{}) () {
  // clipRect()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader8clipRectEv
  default:
    qtrt.ErrorResolve("QImageReader", "clipRect", args)
  }

}

  // proto:  QByteArray QImageReader::format();
func (this *QImageReader) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QImageReader6formatEv
  default:
    qtrt.ErrorResolve("QImageReader", "format", args)
  }

}

// <= body block end

