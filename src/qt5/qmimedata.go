package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qmimedata.h
// dst-file: /src/core/qmimedata.go
//

// header block begin =>


// <= header block end

// main block begin =>
// <= main block end

// use block begin =>
// <= use block end

// ext block begin =>
// #[link(name = "Qt5Core")]
// #[link(name = "Qt5Gui")]
// #[link(name = "Qt5Widgets")]
// #[link(name = "QtInline")]

// extern {
import "fmt"
import "reflect"
import "qtrt"
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
}

// } // <= ext block end

// body block begin =>
// class sizeof(QMimeData)=1
type QMimeData struct {
  /*qbase*/ QObject;
  qclsinst uint64 /* *mut c_void*/;
}


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
  default:
    qtrt.ErrorResolve("QMimeData", "setData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "colorData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "hasHtml", args)
 }

}


func NewQMimeData(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QMimeData", "imageData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "hasFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "setText", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "clear", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "text", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "setHtml", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "setImageData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "hasUrls", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "hasColor", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "removeFormat", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "html", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "urls", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "setColorData", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "hasText", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "metaObject", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "data", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "formats", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QMimeData", "hasImage", args)
 }

}

// <= body block end

