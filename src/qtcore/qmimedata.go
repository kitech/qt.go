package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
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
extern void* C_ZNK9QMimeData9imageDataEv(void* qthis); // 4
  // proto:  QString QMimeData::text();
extern void* C_ZNK9QMimeData4textEv(void* qthis); // 4
  // proto:  void QMimeData::setHtml(const QString & html);
extern void C_ZN9QMimeData7setHtmlERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QMimeData::hasHtml();
extern bool C_ZNK9QMimeData7hasHtmlEv(void* qthis); // 4
  // proto:  bool QMimeData::hasText();
extern bool C_ZNK9QMimeData7hasTextEv(void* qthis); // 4
  // proto:  void QMimeData::QMimeData();
extern void* C_ZN9QMimeDataC2Ev(); // 3
  // proto:  bool QMimeData::hasUrls();
extern bool C_ZNK9QMimeData7hasUrlsEv(void* qthis); // 4
  // proto:  void QMimeData::setImageData(const QVariant & image);
extern void C_ZN9QMimeData12setImageDataERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  QString QMimeData::html();
extern void* C_ZNK9QMimeData4htmlEv(void* qthis); // 4
  // proto:  void QMimeData::setColorData(const QVariant & color);
extern void C_ZN9QMimeData12setColorDataERK8QVariant(void* qthis, void* arg0); // 4
  // proto:  void QMimeData::setData(const QString & mimetype, const QByteArray & data);
extern void C_ZN9QMimeData7setDataERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1); // 4
  // proto:  void QMimeData::removeFormat(const QString & mimetype);
extern void C_ZN9QMimeData12removeFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  bool QMimeData::hasFormat(const QString & mimetype);
extern bool C_ZNK9QMimeData9hasFormatERK7QString(void* qthis, void* arg0); // 4
  // proto:  QByteArray QMimeData::data(const QString & mimetype);
extern void* C_ZNK9QMimeData4dataERK7QString(void* qthis, void* arg0); // 4
  // proto:  const QMetaObject * QMimeData::metaObject();
extern void C_ZNK9QMimeData10metaObjectEv(void* qthis); // 4
  // proto:  void QMimeData::setText(const QString & text);
extern void C_ZN9QMimeData7setTextERK7QString(void* qthis, void* arg0); // 4
  // proto:  QVariant QMimeData::colorData();
extern void* C_ZNK9QMimeData9colorDataEv(void* qthis); // 4
  // proto:  bool QMimeData::hasColor();
extern bool C_ZNK9QMimeData8hasColorEv(void* qthis); // 4
  // proto:  QList<QUrl> QMimeData::urls();
extern void C_ZNK9QMimeData4urlsEv(void* qthis); // 4
  // proto:  QStringList QMimeData::formats();
extern void C_ZNK9QMimeData7formatsEv(void* qthis); // 4
  // proto:  bool QMimeData::hasImage();
extern bool C_ZNK9QMimeData8hasImageEv(void* qthis); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// ~QMimeData()
func (this *QMimeData) Freeqmimedata(args ...interface{}) () {
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
    C.C_ZN9QMimeDataD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "~QMimeData", args)
  }

  return
}

// imageData()
func (this *QMimeData) Imagedata(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData9imageDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "imageData", args)
  }

  return
}

// text()
func (this *QMimeData) Text(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData4textEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "text", args)
  }

  return
}

// setHtml(const class QString &)
func (this *QMimeData) Sethtml(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData7setHtmlERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setHtml", args)
  }

  return
}

// hasHtml()
func (this *QMimeData) Hashtml(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData7hasHtmlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "hasHtml", args)
  }

  return
}

// hasText()
func (this *QMimeData) Hastext(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData7hasTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "hasText", args)
  }

  return
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
    return &QMimeData{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QMimeData", "QMimeData", args)
  }

  return nil // QMimeData{Qclsinst:qthis}
}

// hasUrls()
func (this *QMimeData) Hasurls(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData7hasUrlsEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "hasUrls", args)
  }

  return
}

// setImageData(const class QVariant &)
func (this *QMimeData) Setimagedata(args ...interface{}) () {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData12setImageDataERK8QVariant(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setImageData", args)
  }

  return
}

// html()
func (this *QMimeData) Html(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData4htmlEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "html", args)
  }

  return
}

// setColorData(const class QVariant &)
func (this *QMimeData) Setcolordata(args ...interface{}) () {
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
    var arg0 = args[0].(*QVariant).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData12setColorDataERK8QVariant(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setColorData", args)
  }

  return
}

// setData(const class QString &, const class QByteArray &)
func (this *QMimeData) Setdata(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    C.C_ZN9QMimeData7setDataERK7QStringRK10QByteArray(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QMimeData", "setData", args)
  }

  return
}

// removeFormat(const class QString &)
func (this *QMimeData) Removeformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData12removeFormatERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "removeFormat", args)
  }

  return
}

// hasFormat(const class QString &)
func (this *QMimeData) Hasformat(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMimeData9hasFormatERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "hasFormat", args)
  }

  return
}

// data(const class QString &)
func (this *QMimeData) Data(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK9QMimeData4dataERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "data", args)
  }

  return
}

// metaObject()
func (this *QMimeData) Metaobject(args ...interface{}) () {
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
    C.C_ZNK9QMimeData10metaObjectEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "metaObject", args)
  }

  return
}

// setText(const class QString &)
func (this *QMimeData) Settext(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN9QMimeData7setTextERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QMimeData", "setText", args)
  }

  return
}

// colorData()
func (this *QMimeData) Colordata(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData9colorDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QVariant{}) // "QVariant"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "colorData", args)
  }

  return
}

// hasColor()
func (this *QMimeData) Hascolor(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData8hasColorEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "hasColor", args)
  }

  return
}

// urls()
func (this *QMimeData) Urls(args ...interface{}) () {
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
    C.C_ZNK9QMimeData4urlsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "urls", args)
  }

  return
}

// formats()
func (this *QMimeData) Formats(args ...interface{}) () {
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
    C.C_ZNK9QMimeData7formatsEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "formats", args)
  }

  return
}

// hasImage()
func (this *QMimeData) Hasimage(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK9QMimeData8hasImageEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QMimeData", "hasImage", args)
  }

  return
}

// clear()
func (this *QMimeData) Clear(args ...interface{}) () {
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
    C.C_ZN9QMimeData5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QMimeData", "clear", args)
  }

  return
}

// <= body block end

