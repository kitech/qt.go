package qt5
// auto generated, do not modify.
// created: Mon Feb  1 16:24:50 2016
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

/*
#include <stdlib.h>
#include <stdbool.h>
#include <stdint.h>
#include <wchar.h>
#include <uchar.h>
  // proto:  QByteArray QTextDocumentWriter::format();
extern void* C_ZNK19QTextDocumentWriter6formatEv(void* qthis); // 4
  // proto:  void QTextDocumentWriter::setFormat(const QByteArray & format);
extern void C_ZN19QTextDocumentWriter9setFormatERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  void QTextDocumentWriter::setFileName(const QString & fileName);
extern void C_ZN19QTextDocumentWriter11setFileNameERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextDocumentWriter::QTextDocumentWriter();
extern void* C_ZN19QTextDocumentWriterC2Ev(); // 3
  // proto:  void QTextDocumentWriter::QTextDocumentWriter(QIODevice * device, const QByteArray & format);
extern void* C_ZN19QTextDocumentWriterC2EP9QIODeviceRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  void QTextDocumentWriter::QTextDocumentWriter(const QString & fileName, const QByteArray & format);
extern void* C_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray(void* arg0, void* arg1); // 3
  // proto:  QString QTextDocumentWriter::fileName();
extern void* C_ZNK19QTextDocumentWriter8fileNameEv(void* qthis); // 4
  // proto:  bool QTextDocumentWriter::write(const QTextDocumentFragment & fragment);
extern bool C_ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment(void* qthis, void* arg0); // 4
  // proto:  bool QTextDocumentWriter::write(const QTextDocument * document);
extern bool C_ZN19QTextDocumentWriter5writeEPK13QTextDocument(void* qthis, void* arg0); // 4
  // proto:  QTextCodec * QTextDocumentWriter::codec();
extern void* C_ZNK19QTextDocumentWriter5codecEv(void* qthis); // 4
  // proto:  void QTextDocumentWriter::~QTextDocumentWriter();
extern void C_ZN19QTextDocumentWriterD2Ev(void* qthis); // 4
  // proto:  void QTextDocumentWriter::setDevice(QIODevice * device);
extern void C_ZN19QTextDocumentWriter9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QIODevice * QTextDocumentWriter::device();
extern void* C_ZNK19QTextDocumentWriter6deviceEv(void* qthis); // 4
  // proto:  void QTextDocumentWriter::setCodec(QTextCodec * codec);
extern void C_ZN19QTextDocumentWriter8setCodecEP10QTextCodec(void* qthis, void* arg0); // 4
  // proto: static QList<QByteArray> QTextDocumentWriter::supportedDocumentFormats();
extern void C_ZN19QTextDocumentWriter24supportedDocumentFormatsEv(); // 4
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
  Qclsinst unsafe.Pointer /* *C.void */;
}

// format()
func (this *QTextDocumentWriter) Format(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextDocumentWriter6formatEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "format", args)
  }

  return
}

// setFormat(const class QByteArray &)
func (this *QTextDocumentWriter) Setformat(args ...interface{}) () {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QTextDocumentWriter9setFormatERK10QByteArray(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setFormat", args)
  }

  return
}

// setFileName(const class QString &)
func (this *QTextDocumentWriter) Setfilename(args ...interface{}) () {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QTextDocumentWriter11setFileNameERK7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setFileName", args)
  }

  return
}

// QTextDocumentWriter()
func NewQTextDocumentWriter(args ...interface{}) *QTextDocumentWriter {
  // QTextDocumentWriter()
  // QTextDocumentWriter(class QIODevice *, const class QByteArray &)
  // QTextDocumentWriter(const class QString &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriterC1Ev
    // invoke: void QTextDocumentWriter()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QTextDocumentWriterC2Ev()
    return &QTextDocumentWriter{Qclsinst:qthis}
  case 1:
    // invoke: _ZN19QTextDocumentWriterC1EP9QIODeviceRK10QByteArray
    // invoke: void QTextDocumentWriter(class QIODevice *, const class QByteArray &)
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QTextDocumentWriterC2EP9QIODeviceRK10QByteArray(arg0, arg1)
    return &QTextDocumentWriter{Qclsinst:qthis}
  case 2:
    // invoke: _ZN19QTextDocumentWriterC1ERK7QStringRK10QByteArray
    // invoke: void QTextDocumentWriter(const class QString &, const class QByteArray &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN19QTextDocumentWriterC2ERK7QStringRK10QByteArray(arg0, arg1)
    return &QTextDocumentWriter{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "QTextDocumentWriter", args)
  }

  return nil // QTextDocumentWriter{Qclsinst:qthis}
}

// fileName()
func (this *QTextDocumentWriter) Filename(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextDocumentWriter8fileNameEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "fileName", args)
  }

  return
}

// write(const class QTextDocumentFragment &)
func (this *QTextDocumentWriter) Write(args ...interface{}) (ret interface{}) {
  // write(const class QTextDocumentFragment &)
  // write(const class QTextDocument *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextDocumentFragment{}) // "const QTextDocumentFragment &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QTextDocument{}) // "const QTextDocument *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment
    // invoke: bool write(const class QTextDocumentFragment &)
    var arg0 = args[0].(*QTextDocumentFragment).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN19QTextDocumentWriter5writeERK21QTextDocumentFragment(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN19QTextDocumentWriter5writeEPK13QTextDocument
    // invoke: bool write(const class QTextDocument *)
    var arg0 = args[0].(*QTextDocument).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN19QTextDocumentWriter5writeEPK13QTextDocument(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "write", args)
  }

  return
}

// codec()
func (this *QTextDocumentWriter) Codec(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextDocumentWriter5codecEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "codec", args)
  }

  return
}

// ~QTextDocumentWriter()
func (this *QTextDocumentWriter) Freeqtextdocumentwriter(args ...interface{}) () {
  // ~QTextDocumentWriter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriterD0Ev
    // invoke: void ~QTextDocumentWriter()
    C.C_ZN19QTextDocumentWriterD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "~QTextDocumentWriter", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QTextDocumentWriter) Setdevice(args ...interface{}) () {
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
    var arg0 = args[0].(*QIODevice).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QTextDocumentWriter9setDeviceEP9QIODevice(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setDevice", args)
  }

  return
}

// device()
func (this *QTextDocumentWriter) Device(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK19QTextDocumentWriter6deviceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "device", args)
  }

  return
}

// setCodec(class QTextCodec *)
func (this *QTextDocumentWriter) Setcodec(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN19QTextDocumentWriter8setCodecEP10QTextCodec(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "setCodec", args)
  }

  return
}

// supportedDocumentFormats()
func (this *QTextDocumentWriter) Supporteddocumentformats_S(args ...interface{}) () {
  // supportedDocumentFormats()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN19QTextDocumentWriter24supportedDocumentFormatsEv
    // invoke: QList<QByteArray> supportedDocumentFormats()
    C.C_ZN19QTextDocumentWriter24supportedDocumentFormatsEv()
  default:
    qtrt.ErrorResolve("QTextDocumentWriter", "supportedDocumentFormats", args)
  }

  return
}

// <= body block end

