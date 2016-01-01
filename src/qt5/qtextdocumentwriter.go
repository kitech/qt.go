package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtGui/qtextdocumentwriter.h
// dst-file: /src/gui/qtextdocumentwriter.go
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
// class sizeof(QTextDocumentWriter)=8
type QTextDocumentWriter struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextDocumentWriter) setCodec(args ...interface{}) () {
  // setCodec(class QTextCodec *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter8setCodecEP10QTextCodec
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setCodec", args)
 }

}


func NewQTextDocumentWriter(args ...interface{})() {
}


func (this *QTextDocumentWriter) setFileName(args ...interface{}) () {
  // setFileName(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter11setFileNameERK7QString
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setFileName", args)
 }

}


func (this *QTextDocumentWriter) format(args ...interface{}) () {
  // format()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextDocumentWriter6formatEv
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "format", args)
 }

}


func (this *QTextDocumentWriter) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setDevice", args)
 }

}


func (this *QTextDocumentWriter) setFormat(args ...interface{}) () {
  // setFormat(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter9setFormatERK10QByteArray
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setFormat", args)
 }

}


func (this *QTextDocumentWriter) write(args ...interface{}) () {
  // write(const class QTextDocument *)
  // write(const class QTextDocumentFragment &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocument{}) // "const QTextDocument *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextDocumentFragment{}) // "const QTextDocumentFragment &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter5writeEPK13QTextDocument
  case 1:
    // invoke: _ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "write", args)
 }

}


func (this *QTextDocumentWriter) codec(args ...interface{}) () {
  // codec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextDocumentWriter5codecEv
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "codec", args)
 }

}


func (this *QTextDocumentWriter) fileName(args ...interface{}) () {
  // fileName()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextDocumentWriter8fileNameEv
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "fileName", args)
 }

}


func (this *QTextDocumentWriter) supportedDocumentFormats_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "supportedDocumentFormats", args)
 }

}


func (this *QTextDocumentWriter) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK19QTextDocumentWriter6deviceEv
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "device", args)
 }

}


func (this *QTextDocumentWriter) FreeQTextDocumentWriter(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "~QTextDocumentWriter", args)
 }

}

// <= body block end

