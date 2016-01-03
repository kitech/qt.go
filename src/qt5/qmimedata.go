package qt5
// auto generated, do not modify.
// created: Sun Jan  3 20:07:07 2016
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
  // proto:  void QMimeData::setData(const QString & mimetype, const QByteArray & data);
extern void _ZN9QMimeData7setDataERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  QVariant QMimeData::colorData();
extern void _ZNK9QMimeData9colorDataEv(void* qthis);
  // proto:  void QMimeData::~QMimeData();
extern void _ZN9QMimeDataD0Ev(void* qthis);
  // proto:  bool QMimeData::hasHtml();
extern void _ZNK9QMimeData7hasHtmlEv(void* qthis);
  // proto:  void QMimeData::QMimeData(const QMimeData & );
extern void* dector_ZN9QMimeDataC1ERKS_(void* arg0);
extern void _ZN9QMimeDataC1ERKS_(void* qthis, void* arg0);
  // proto:  QVariant QMimeData::imageData();
extern void _ZNK9QMimeData9imageDataEv(void* qthis);
  // proto:  bool QMimeData::hasFormat(const QString & mimetype);
extern void _ZNK9QMimeData9hasFormatERK7QString(void* qthis, void* arg0);
  // proto:  void QMimeData::setText(const QString & text);
extern void _ZN9QMimeData7setTextERK7QString(void* qthis, void* arg0);
  // proto:  void QMimeData::clear();
extern void _ZN9QMimeData5clearEv(void* qthis);
  // proto:  QString QMimeData::text();
extern void _ZNK9QMimeData4textEv(void* qthis);
  // proto:  void QMimeData::setHtml(const QString & html);
extern void _ZN9QMimeData7setHtmlERK7QString(void* qthis, void* arg0);
  // proto:  void QMimeData::setImageData(const QVariant & image);
extern void _ZN9QMimeData12setImageDataERK8QVariant(void* qthis, void* arg0);
  // proto:  bool QMimeData::hasUrls();
extern void _ZNK9QMimeData7hasUrlsEv(void* qthis);
  // proto:  bool QMimeData::hasColor();
extern void _ZNK9QMimeData8hasColorEv(void* qthis);
  // proto:  void QMimeData::removeFormat(const QString & mimetype);
extern void _ZN9QMimeData12removeFormatERK7QString(void* qthis, void* arg0);
  // proto:  QString QMimeData::html();
extern void _ZNK9QMimeData4htmlEv(void* qthis);
  // proto:  void QMimeData::QMimeData();
extern void* dector_ZN9QMimeDataC1Ev();
extern void _ZN9QMimeDataC1Ev(void* qthis);
  // proto:  QList<QUrl> QMimeData::urls();
extern void _ZNK9QMimeData4urlsEv(void* qthis);
  // proto:  void QMimeData::setColorData(const QVariant & color);
extern void _ZN9QMimeData12setColorDataERK8QVariant(void* qthis, void* arg0);
  // proto:  bool QMimeData::hasText();
extern void _ZNK9QMimeData7hasTextEv(void* qthis);
  // proto:  const QMetaObject * QMimeData::metaObject();
extern void _ZNK9QMimeData10metaObjectEv(void* qthis);
  // proto:  QByteArray QMimeData::data(const QString & mimetype);
extern void _ZNK9QMimeData4dataERK7QString(void* qthis, void* arg0);
  // proto:  QStringList QMimeData::formats();
extern void _ZNK9QMimeData7formatsEv(void* qthis);
  // proto:  bool QMimeData::hasImage();
extern void _ZNK9QMimeData8hasImageEv(void* qthis);
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

  // proto:  void QMimeData::setData(const QString & mimetype, const QByteArray & data);
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
    C._ZN9QMimeData7setDataERK7QStringRK10QByteArray(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMimeData", "setData", args)
  }

}

  // proto:  QVariant QMimeData::colorData();
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
    C._ZNK9QMimeData9colorDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "colorData", args)
  }

}

  // proto:  void QMimeData::~QMimeData();
func (this *QMimeData) FreeQMimeData(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QMimeData", "~QMimeData", args)
  }

}

  // proto:  bool QMimeData::hasHtml();
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
    C._ZNK9QMimeData7hasHtmlEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "hasHtml", args)
  }

}

  // proto:  void QMimeData::QMimeData(const QMimeData & );
func NewQMimeData(args ...interface{}) QMimeData {
  return QMimeData{}
}

  // proto:  QVariant QMimeData::imageData();
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
    C._ZNK9QMimeData9imageDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "imageData", args)
  }

}

  // proto:  bool QMimeData::hasFormat(const QString & mimetype);
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
    C._ZNK9QMimeData9hasFormatERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "hasFormat", args)
  }

}

  // proto:  void QMimeData::setText(const QString & text);
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
    C._ZN9QMimeData7setTextERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setText", args)
  }

}

  // proto:  void QMimeData::clear();
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
    C._ZN9QMimeData5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "clear", args)
  }

}

  // proto:  QString QMimeData::text();
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
    C._ZNK9QMimeData4textEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "text", args)
  }

}

  // proto:  void QMimeData::setHtml(const QString & html);
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
    C._ZN9QMimeData7setHtmlERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setHtml", args)
  }

}

  // proto:  void QMimeData::setImageData(const QVariant & image);
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
    C._ZN9QMimeData12setImageDataERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setImageData", args)
  }

}

  // proto:  bool QMimeData::hasUrls();
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
    C._ZNK9QMimeData7hasUrlsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "hasUrls", args)
  }

}

  // proto:  bool QMimeData::hasColor();
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
    C._ZNK9QMimeData8hasColorEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "hasColor", args)
  }

}

  // proto:  void QMimeData::removeFormat(const QString & mimetype);
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
    C._ZN9QMimeData12removeFormatERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "removeFormat", args)
  }

}

  // proto:  QString QMimeData::html();
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
    C._ZNK9QMimeData4htmlEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "html", args)
  }

}

  // proto:  QList<QUrl> QMimeData::urls();
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
    C._ZNK9QMimeData4urlsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "urls", args)
  }

}

  // proto:  void QMimeData::setColorData(const QVariant & color);
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
    C._ZN9QMimeData12setColorDataERK8QVariant(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setColorData", args)
  }

}

  // proto:  bool QMimeData::hasText();
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
    C._ZNK9QMimeData7hasTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "hasText", args)
  }

}

  // proto:  const QMetaObject * QMimeData::metaObject();
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
    C._ZNK9QMimeData10metaObjectEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "metaObject", args)
  }

}

  // proto:  QByteArray QMimeData::data(const QString & mimetype);
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
    C._ZNK9QMimeData4dataERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "data", args)
  }

}

  // proto:  QStringList QMimeData::formats();
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
    C._ZNK9QMimeData7formatsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "formats", args)
  }

}

  // proto:  bool QMimeData::hasImage();
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
    C._ZNK9QMimeData8hasImageEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "hasImage", args)
  }

}

// <= body block end

