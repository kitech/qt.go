package qt5
// auto generated, do not modify.
// created: Sun Jan 31 14:26:18 2016
// src-file: /QtCore/qmimedata.h
// dst-file: /src/core/qmimedata.go
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
  // proto:  void QMimeData::~QMimeData();
extern void C_ZN9QMimeDataD2Ev(void* qthis); // 4
  // proto:  QVariant QMimeData::imageData();
extern void C_ZNK9QMimeData9imageDataEv(void* qthis); // 4
  // proto:  QString QMimeData::text();
extern void C_ZNK9QMimeData4textEv(void* qthis); // 4
  // proto:  void QMimeData::setHtml(const QString & html);
extern void C_ZN9QMimeData7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QMimeData::hasHtml();
extern void C_ZNK9QMimeData7hasHtmlEv(void* qthis); // 4
  // proto:  bool QMimeData::hasText();
extern void C_ZNK9QMimeData7hasTextEv(void* qthis); // 4
  // proto:  void QMimeData::QMimeData();
extern void* C_ZN9QMimeDataC2Ev(); // 3
  // proto:  bool QMimeData::hasUrls();
extern void C_ZNK9QMimeData7hasUrlsEv(void* qthis); // 4
  // proto:  void QMimeData::setImageData(const QVariant & image);
extern void C_ZN9QMimeData12setImageDataERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  QString QMimeData::html();
extern void C_ZNK9QMimeData4htmlEv(void* qthis); // 4
  // proto:  void QMimeData::setColorData(const QVariant & color);
extern void C_ZN9QMimeData12setColorDataERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  void QMimeData::setData(const QString & mimetype, const QByteArray & data);
extern void C_ZN9QMimeData7setDataERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QMimeData::removeFormat(const QString & mimetype);
extern void C_ZN9QMimeData12removeFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QMimeData::hasFormat(const QString & mimetype);
extern void C_ZNK9QMimeData9hasFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  QByteArray QMimeData::data(const QString & mimetype);
extern void C_ZNK9QMimeData4dataERK7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QMimeData::metaObject();
extern void C_ZNK9QMimeData10metaObjectEv(void* qthis); // 4
  // proto:  void QMimeData::setText(const QString & text);
extern void C_ZN9QMimeData7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QVariant QMimeData::colorData();
extern void C_ZNK9QMimeData9colorDataEv(void* qthis); // 4
  // proto:  bool QMimeData::hasColor();
extern void C_ZNK9QMimeData8hasColorEv(void* qthis); // 4
  // proto:  QList<QUrl> QMimeData::urls();
extern void C_ZNK9QMimeData4urlsEv(void* qthis); // 4
  // proto:  QStringList QMimeData::formats();
extern void C_ZNK9QMimeData7formatsEv(void* qthis); // 4
  // proto:  bool QMimeData::hasImage();
extern void C_ZNK9QMimeData8hasImageEv(void* qthis); // 4
  // proto:  void QMimeData::clear();
extern void C_ZN9QMimeData5clearEv(void* qthis); // 4
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

// class sizeof(QMimeData)=1
type QMimeData struct {
  /*qbase*/ QObject;
  qclsinst unsafe.Pointer /* *C.void */;
}

// ~QMimeData()
func (this *QMimeData) FreeQMimeData(args ...interface{}) () {
  // ~QMimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeDataD0Ev
    // invoke: void ~QMimeData()
    C.C_ZN9QMimeDataD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "~QMimeData", args)
  }

}

// imageData()
func (this *QMimeData) imageData(args ...interface{}) () {
  // imageData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData9imageDataEv
    // invoke: QVariant imageData()
    var ret = C.C_ZNK9QMimeData9imageDataEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "imageData", args)
  }

}

// text()
func (this *QMimeData) text(args ...interface{}) () {
  // text()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData4textEv
    // invoke: QString text()
    var ret = C.C_ZNK9QMimeData4textEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "text", args)
  }

}

// setHtml(const class QString &)
func (this *QMimeData) setHtml(args ...interface{}) () {
  // setHtml(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData7setHtmlERK7QString
    // invoke: void setHtml(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData7setHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setHtml", args)
  }

}

// hasHtml()
func (this *QMimeData) hasHtml(args ...interface{}) () {
  // hasHtml()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData7hasHtmlEv
    // invoke: bool hasHtml()
    var ret = C.C_ZNK9QMimeData7hasHtmlEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "hasHtml", args)
  }

}

// hasText()
func (this *QMimeData) hasText(args ...interface{}) () {
  // hasText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData7hasTextEv
    // invoke: bool hasText()
    var ret = C.C_ZNK9QMimeData7hasTextEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "hasText", args)
  }

}

// QMimeData()
func NewQMimeData(args ...interface{}) *QMimeData {
  // QMimeData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeDataC1Ev
    // invoke: void QMimeData()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN9QMimeDataC2Ev()
    return &QMimeData{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMimeData", "QMimeData", args)
  }

  return nil // QMimeData{qclsinst:qthis}
}

// hasUrls()
func (this *QMimeData) hasUrls(args ...interface{}) () {
  // hasUrls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData7hasUrlsEv
    // invoke: bool hasUrls()
    var ret = C.C_ZNK9QMimeData7hasUrlsEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "hasUrls", args)
  }

}

// setImageData(const class QVariant &)
func (this *QMimeData) setImageData(args ...interface{}) () {
  // setImageData(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData12setImageDataERK8QVariant
    // invoke: void setImageData(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData12setImageDataERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setImageData", args)
  }

}

// html()
func (this *QMimeData) html(args ...interface{}) () {
  // html()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData4htmlEv
    // invoke: QString html()
    var ret = C.C_ZNK9QMimeData4htmlEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "html", args)
  }

}

// setColorData(const class QVariant &)
func (this *QMimeData) setColorData(args ...interface{}) () {
  // setColorData(const class QVariant &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QVariant{}) // "const QVariant &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData12setColorDataERK8QVariant
    // invoke: void setColorData(const class QVariant &)
    var arg0 = args[0].(QVariant).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData12setColorDataERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setColorData", args)
  }

}

// setData(const class QString &, const class QByteArray &)
func (this *QMimeData) setData(args ...interface{}) () {
  // setData(const class QString &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData7setDataERK7QStringRK10QByteArray
    // invoke: void setData(const class QString &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QMimeData7setDataERK7QStringRK10QByteArray(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMimeData", "setData", args)
  }

}

// removeFormat(const class QString &)
func (this *QMimeData) removeFormat(args ...interface{}) () {
  // removeFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData12removeFormatERK7QString
    // invoke: void removeFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData12removeFormatERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "removeFormat", args)
  }

}

// hasFormat(const class QString &)
func (this *QMimeData) hasFormat(args ...interface{}) () {
  // hasFormat(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData9hasFormatERK7QString
    // invoke: bool hasFormat(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK9QMimeData9hasFormatERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "hasFormat", args)
  }

}

// data(const class QString &)
func (this *QMimeData) data(args ...interface{}) () {
  // data(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData4dataERK7QString
    // invoke: QByteArray data(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var ret = C.C_ZNK9QMimeData4dataERK7QString(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "data", args)
  }

}

// metaObject()
func (this *QMimeData) metaObject(args ...interface{}) () {
  // metaObject()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData10metaObjectEv
    // invoke: const QMetaObject * metaObject()
    C.C_ZNK9QMimeData10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "metaObject", args)
  }

}

// setText(const class QString &)
func (this *QMimeData) setText(args ...interface{}) () {
  // setText(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData7setTextERK7QString
    // invoke: void setText(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setText", args)
  }

}

// colorData()
func (this *QMimeData) colorData(args ...interface{}) () {
  // colorData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData9colorDataEv
    // invoke: QVariant colorData()
    var ret = C.C_ZNK9QMimeData9colorDataEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "colorData", args)
  }

}

// hasColor()
func (this *QMimeData) hasColor(args ...interface{}) () {
  // hasColor()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData8hasColorEv
    // invoke: bool hasColor()
    var ret = C.C_ZNK9QMimeData8hasColorEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "hasColor", args)
  }

}

// urls()
func (this *QMimeData) urls(args ...interface{}) () {
  // urls()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData4urlsEv
    // invoke: QList<QUrl> urls()
    C.C_ZNK9QMimeData4urlsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "urls", args)
  }

}

// formats()
func (this *QMimeData) formats(args ...interface{}) () {
  // formats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData7formatsEv
    // invoke: QStringList formats()
    C.C_ZNK9QMimeData7formatsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "formats", args)
  }

}

// hasImage()
func (this *QMimeData) hasImage(args ...interface{}) () {
  // hasImage()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK9QMimeData8hasImageEv
    // invoke: bool hasImage()
    var ret = C.C_ZNK9QMimeData8hasImageEv(this.qclsinst)
    if false {reflect.TypeOf(ret)}
  default:
    qtrt.ErrorResolve("QMimeData", "hasImage", args)
  }

}

// clear()
func (this *QMimeData) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN9QMimeData5clearEv
    // invoke: void clear()
    C.C_ZN9QMimeData5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "clear", args)
  }

}

// <= body block end

