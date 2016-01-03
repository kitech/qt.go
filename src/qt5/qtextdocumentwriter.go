package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
// src-file: /QtGui/qtextdocumentwriter.h
// dst-file: /src/gui/qtextdocumentwriter.go
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
  // proto:  void QTextDocumentWriter::setCodec(QTextCodec * codec);
extern void _ZN19QTextDocumentWriter8setCodecEP10QTextCodec(void* qthis, void* arg0);
  // proto:  void QTextDocumentWriter::QTextDocumentWriter(QIODevice * device, const QByteArray & format);
extern void* dector_ZN19QTextDocumentWriterC1EP9QIODeviceRK10QByteArray(void* arg0, void* arg1);
extern void _ZN19QTextDocumentWriterC1EP9QIODeviceRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  void QTextDocumentWriter::setFileName(const QString & fileName);
extern void _ZN19QTextDocumentWriter11setFileNameERK7QString(void* qthis, void* arg0);
  // proto:  QByteArray QTextDocumentWriter::format();
extern void _ZNK19QTextDocumentWriter6formatEv(void* qthis);
  // proto:  void QTextDocumentWriter::setDevice(QIODevice * device);
extern void _ZN19QTextDocumentWriter9setDeviceEP9QIODevice(void* qthis, void* arg0);
  // proto:  void QTextDocumentWriter::QTextDocumentWriter(const QString & fileName, const QByteArray & format);
extern void* dector_ZN19QTextDocumentWriterC1ERK7QStringRK10QByteArray(void* arg0, void* arg1);
extern void _ZN19QTextDocumentWriterC1ERK7QStringRK10QByteArray(void* qthis, void* arg0, void* arg1);
  // proto:  void QTextDocumentWriter::setFormat(const QByteArray & format);
extern void _ZN19QTextDocumentWriter9setFormatERK10QByteArray(void* qthis, void* arg0);
  // proto:  bool QTextDocumentWriter::write(const QTextDocument * document);
extern void _ZN19QTextDocumentWriter5writeEPK13QTextDocument(void* qthis, void* arg0);
  // proto:  bool QTextDocumentWriter::write(const QTextDocumentFragment & fragment);
extern void _ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment(void* qthis, void* arg0);
  // proto:  void QTextDocumentWriter::QTextDocumentWriter();
extern void* dector_ZN19QTextDocumentWriterC1Ev();
extern void _ZN19QTextDocumentWriterC1Ev(void* qthis);
  // proto:  QTextCodec * QTextDocumentWriter::codec();
extern void _ZNK19QTextDocumentWriter5codecEv(void* qthis);
  // proto:  QString QTextDocumentWriter::fileName();
extern void _ZNK19QTextDocumentWriter8fileNameEv(void* qthis);
  // proto: static QList<QByteArray> QTextDocumentWriter::supportedDocumentFormats();
extern void _ZN19QTextDocumentWriter24supportedDocumentFormatsEv();
  // proto:  QIODevice * QTextDocumentWriter::device();
extern void _ZNK19QTextDocumentWriter6deviceEv(void* qthis);
  // proto:  void QTextDocumentWriter::~QTextDocumentWriter();
extern void _ZN19QTextDocumentWriterD0Ev(void* qthis);
  // proto:  void QTextDocumentWriter::QTextDocumentWriter(const QTextDocumentWriter & );
extern void* dector_ZN19QTextDocumentWriterC1ERKS_(void* arg0);
extern void _ZN19QTextDocumentWriterC1ERKS_(void* qthis, void* arg0);
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

// class sizeof(QTextDocumentWriter)=8
type QTextDocumentWriter struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

  // proto:  void QTextDocumentWriter::setCodec(QTextCodec * codec);
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
    // invoke: void setCodec(class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QTextDocumentWriter8setCodecEP10QTextCodec(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setCodec", args)
  }

}

  // proto:  void QTextDocumentWriter::QTextDocumentWriter(QIODevice * device, const QByteArray & format);
func NewQTextDocumentWriter(args ...interface{}) QTextDocumentWriter {
  return QTextDocumentWriter{}
}

  // proto:  void QTextDocumentWriter::setFileName(const QString & fileName);
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
    // invoke: void setFileName(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QTextDocumentWriter11setFileNameERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setFileName", args)
  }

}

  // proto:  QByteArray QTextDocumentWriter::format();
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
    // invoke: QByteArray format()
    C._ZNK19QTextDocumentWriter6formatEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "format", args)
  }

}

  // proto:  void QTextDocumentWriter::setDevice(QIODevice * device);
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
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QTextDocumentWriter9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setDevice", args)
  }

}

  // proto:  void QTextDocumentWriter::setFormat(const QByteArray & format);
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
    // invoke: void setFormat(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QTextDocumentWriter9setFormatERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setFormat", args)
  }

}

  // proto:  bool QTextDocumentWriter::write(const QTextDocument * document);
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
    // invoke: bool write(const class QTextDocument *)
    var arg0 = args[0].(QTextDocument).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QTextDocumentWriter5writeEPK13QTextDocument(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment
    // invoke: bool write(const class QTextDocumentFragment &)
    var arg0 = args[0].(QTextDocumentFragment).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "write", args)
  }

}

  // proto:  QTextCodec * QTextDocumentWriter::codec();
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
    // invoke: QTextCodec * codec()
    C._ZNK19QTextDocumentWriter5codecEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "codec", args)
  }

}

  // proto:  QString QTextDocumentWriter::fileName();
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
    // invoke: QString fileName()
    C._ZNK19QTextDocumentWriter8fileNameEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "fileName", args)
  }

}

  // proto: static QList<QByteArray> QTextDocumentWriter::supportedDocumentFormats();
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

  // proto:  QIODevice * QTextDocumentWriter::device();
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
    // invoke: QIODevice * device()
    C._ZNK19QTextDocumentWriter6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "device", args)
  }

}

  // proto:  void QTextDocumentWriter::~QTextDocumentWriter();
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

