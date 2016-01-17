package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:13 2016
// src-file: /QtCore/qtextcodec.h
// dst-file: /src/core/qtextcodec.go
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
  // proto:  void QTextEncoder::QTextEncoder(const QTextCodec * codec);
extern void _ZN12QTextEncoderC2EPK10QTextCodec(void* qthis, void* arg0); // 1
  // proto:  bool QTextEncoder::hasFailure();
extern void _ZNK12QTextEncoder10hasFailureEv(void* qthis); // 4
  // proto:  QByteArray QTextEncoder::fromUnicode(const QChar * uc, int len);
extern void _ZN12QTextEncoder11fromUnicodeEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QByteArray QTextEncoder::fromUnicode(const QString & str);
extern void _ZN12QTextEncoder11fromUnicodeERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEncoder::~QTextEncoder();
extern void _ZN12QTextEncoderD2Ev(void* qthis); // 4
  // proto: static void QTextCodec::setCodecForLocale(QTextCodec * c);
extern void _ZN10QTextCodec17setCodecForLocaleEPS_(void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForName(const QByteArray & name);
extern void _ZN10QTextCodec12codecForNameERK10QByteArray(void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForName(const char * name);
extern void _ZN10QTextCodec12codecForNameEPKc(unsigned char* arg0); // 2
  // proto:  bool QTextCodec::canEncode(QChar );
extern void _ZNK10QTextCodec9canEncodeE5QChar(void* qthis, void* arg0); // 4
  // proto:  bool QTextCodec::canEncode(const QString & );
extern void _ZNK10QTextCodec9canEncodeERK7QString(void* qthis, void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForLocale();
extern void _ZN10QTextCodec14codecForLocaleEv(); // 4
  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba);
extern void _ZN10QTextCodec12codecForHtmlERK10QByteArray(void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba, QTextCodec * defaultCodec);
extern void _ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_(void* arg0, void* arg1); // 4
  // proto: static QList<QByteArray> QTextCodec::availableCodecs();
extern void _ZN10QTextCodec15availableCodecsEv(); // 4
  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba, QTextCodec * defaultCodec);
extern void _ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_(void* arg0, void* arg1); // 4
  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba);
extern void _ZN10QTextCodec15codecForUtfTextERK10QByteArray(void* arg0); // 4
  // proto: static QList<int> QTextCodec::availableMibs();
extern void _ZN10QTextCodec13availableMibsEv(); // 4
  // proto:  QByteArray QTextCodec::fromUnicode(const QString & uc);
extern void _ZNK10QTextCodec11fromUnicodeERK7QString(void* qthis, void* arg0); // 4
  // proto:  QList<QByteArray> QTextCodec::aliases();
extern void _ZNK10QTextCodec7aliasesEv(void* qthis); // 4
  // proto: static QTextCodec * QTextCodec::codecForMib(int mib);
extern void _ZN10QTextCodec11codecForMibEi(int32_t arg0); // 4
  // proto:  QString QTextCodec::toUnicode(const QByteArray & );
extern void _ZNK10QTextCodec9toUnicodeERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QString QTextCodec::toUnicode(const char * chars);
extern void _ZNK10QTextCodec9toUnicodeEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  void QTextDecoder::QTextDecoder(const QTextCodec * codec);
extern void _ZN12QTextDecoderC2EPK10QTextCodec(void* qthis, void* arg0); // 1
  // proto:  bool QTextDecoder::hasFailure();
extern void _ZNK12QTextDecoder10hasFailureEv(void* qthis); // 4
  // proto:  void QTextDecoder::~QTextDecoder();
extern void _ZN12QTextDecoderD2Ev(void* qthis); // 4
  // proto:  QString QTextDecoder::toUnicode(const QByteArray & ba);
extern void _ZN12QTextDecoder9toUnicodeERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QString QTextDecoder::toUnicode(const char * chars, int len);
extern void _ZN12QTextDecoder9toUnicodeEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  void QTextDecoder::toUnicode(QString * target, const char * chars, int len);
extern void _ZN12QTextDecoder9toUnicodeEP7QStringPKci(void* qthis, void* arg0, unsigned char* arg1, int32_t arg2); // 4
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

// class sizeof(QTextEncoder)=1
type QTextEncoder struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextCodec)=8
type QTextCodec struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextDecoder)=1
type QTextDecoder struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QTextEncoder(const class QTextCodec *)
func NewQTextEncoder(args ...interface{}) QTextEncoder {
  // QTextEncoder(const class QTextCodec *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "const QTextCodec *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextEncoderC1EPK10QTextCodec
    // invoke: void QTextEncoder(const class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QTextEncoderC2EPK10QTextCodec(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextEncoder", "QTextEncoder", args)
  }

  return QTextEncoder{}
}

// hasFailure()
func (this *QTextEncoder) hasFailure(args ...interface{}) () {
  // hasFailure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextEncoder10hasFailureEv
    // invoke: bool hasFailure()
    C._ZNK12QTextEncoder10hasFailureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEncoder", "hasFailure", args)
  }

}

// fromUnicode(const class QChar *, int)
func (this *QTextEncoder) fromUnicode(args ...interface{}) () {
  // fromUnicode(const class QChar *, int)
  // fromUnicode(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextEncoder11fromUnicodeEPK5QChari
    // invoke: QByteArray fromUnicode(const class QChar *, int)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN12QTextEncoder11fromUnicodeEPK5QChari(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN12QTextEncoder11fromUnicodeERK7QString
    // invoke: QByteArray fromUnicode(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QTextEncoder11fromUnicodeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextEncoder", "fromUnicode", args)
  }

}

// ~QTextEncoder()
func (this *QTextEncoder) FreeQTextEncoder(args ...interface{}) () {
  // ~QTextEncoder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextEncoderD0Ev
    // invoke: void ~QTextEncoder()
    C._ZN12QTextEncoderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextEncoder", "~QTextEncoder", args)
  }

}

// setCodecForLocale(class QTextCodec *)
func (this *QTextCodec) setCodecForLocale_s(args ...interface{}) () {
  // setCodecForLocale(class QTextCodec *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec17setCodecForLocaleEPS_
    // invoke: void setCodecForLocale(class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextCodec17setCodecForLocaleEPS_(arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "setCodecForLocale", args)
  }

}

// codecForName(const class QByteArray &)
func (this *QTextCodec) codecForName_s(args ...interface{}) () {
  // codecForName(const class QByteArray &)
  // codecForName(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec12codecForNameERK10QByteArray
    // invoke: QTextCodec * codecForName(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextCodec12codecForNameERK10QByteArray(arg0)
  case 1:
    // invoke: _ZN10QTextCodec12codecForNameEPKc
    // invoke: QTextCodec * codecForName(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN10QTextCodec12codecForNameEPKc(arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForName", args)
  }

}

// canEncode(class QChar)
func (this *QTextCodec) canEncode(args ...interface{}) () {
  // canEncode(class QChar)
  // canEncode(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec9canEncodeE5QChar
    // invoke: bool canEncode(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextCodec9canEncodeE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QTextCodec9canEncodeERK7QString
    // invoke: bool canEncode(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextCodec9canEncodeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "canEncode", args)
  }

}

// codecForLocale()
func (this *QTextCodec) codecForLocale_s(args ...interface{}) () {
  // codecForLocale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec14codecForLocaleEv
    // invoke: QTextCodec * codecForLocale()
    C._ZN10QTextCodec14codecForLocaleEv()
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForLocale", args)
  }

}

// codecForHtml(const class QByteArray &)
func (this *QTextCodec) codecForHtml_s(args ...interface{}) () {
  // codecForHtml(const class QByteArray &)
  // codecForHtml(const class QByteArray &, class QTextCodec *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec12codecForHtmlERK10QByteArray
    // invoke: QTextCodec * codecForHtml(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextCodec12codecForHtmlERK10QByteArray(arg0)
  case 1:
    // invoke: _ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_
    // invoke: QTextCodec * codecForHtml(const class QByteArray &, class QTextCodec *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextCodec).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_(arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForHtml", args)
  }

}

// availableCodecs()
func (this *QTextCodec) availableCodecs_s(args ...interface{}) () {
  // availableCodecs()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec15availableCodecsEv
    // invoke: QList<QByteArray> availableCodecs()
    C._ZN10QTextCodec15availableCodecsEv()
  default:
    qtrt.ErrorResolve("QTextCodec", "availableCodecs", args)
  }

}

// codecForUtfText(const class QByteArray &, class QTextCodec *)
func (this *QTextCodec) codecForUtfText_s(args ...interface{}) () {
  // codecForUtfText(const class QByteArray &, class QTextCodec *)
  // codecForUtfText(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_
    // invoke: QTextCodec * codecForUtfText(const class QByteArray &, class QTextCodec *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QTextCodec).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_(arg0, arg1)
  case 1:
    // invoke: _ZN10QTextCodec15codecForUtfTextERK10QByteArray
    // invoke: QTextCodec * codecForUtfText(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QTextCodec15codecForUtfTextERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForUtfText", args)
  }

}

// availableMibs()
func (this *QTextCodec) availableMibs_s(args ...interface{}) () {
  // availableMibs()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec13availableMibsEv
    // invoke: QList<int> availableMibs()
    C._ZN10QTextCodec13availableMibsEv()
  default:
    qtrt.ErrorResolve("QTextCodec", "availableMibs", args)
  }

}

// fromUnicode(const class QString &)
func (this *QTextCodec) fromUnicode(args ...interface{}) () {
  // fromUnicode(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec11fromUnicodeERK7QString
    // invoke: QByteArray fromUnicode(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextCodec11fromUnicodeERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "fromUnicode", args)
  }

}

// aliases()
func (this *QTextCodec) aliases(args ...interface{}) () {
  // aliases()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec7aliasesEv
    // invoke: QList<QByteArray> aliases()
    C._ZNK10QTextCodec7aliasesEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextCodec", "aliases", args)
  }

}

// codecForMib(int)
func (this *QTextCodec) codecForMib_s(args ...interface{}) () {
  // codecForMib(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QTextCodec11codecForMibEi
    // invoke: QTextCodec * codecForMib(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QTextCodec11codecForMibEi(arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForMib", args)
  }

}

// toUnicode(const class QByteArray &)
func (this *QTextCodec) toUnicode(args ...interface{}) () {
  // toUnicode(const class QByteArray &)
  // toUnicode(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec9toUnicodeERK10QByteArray
    // invoke: QString toUnicode(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QTextCodec9toUnicodeERK10QByteArray(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QTextCodec9toUnicodeEPKc
    // invoke: QString toUnicode(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK10QTextCodec9toUnicodeEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "toUnicode", args)
  }

}

// QTextDecoder(const class QTextCodec *)
func NewQTextDecoder(args ...interface{}) QTextDecoder {
  // QTextDecoder(const class QTextCodec *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "const QTextCodec *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextDecoderC1EPK10QTextCodec
    // invoke: void QTextDecoder(const class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN12QTextDecoderC2EPK10QTextCodec(qthis, arg0)
  default:
    qtrt.ErrorResolve("QTextDecoder", "QTextDecoder", args)
  }

  return QTextDecoder{}
}

// hasFailure()
func (this *QTextDecoder) hasFailure(args ...interface{}) () {
  // hasFailure()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK12QTextDecoder10hasFailureEv
    // invoke: bool hasFailure()
    C._ZNK12QTextDecoder10hasFailureEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDecoder", "hasFailure", args)
  }

}

// ~QTextDecoder()
func (this *QTextDecoder) FreeQTextDecoder(args ...interface{}) () {
  // ~QTextDecoder()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextDecoderD0Ev
    // invoke: void ~QTextDecoder()
    C._ZN12QTextDecoderD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextDecoder", "~QTextDecoder", args)
  }

}

// toUnicode(const class QByteArray &)
func (this *QTextDecoder) toUnicode(args ...interface{}) () {
  // toUnicode(const class QByteArray &)
  // toUnicode(const char *, int)
  // toUnicode(class QString *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN12QTextDecoder9toUnicodeERK10QByteArray
    // invoke: QString toUnicode(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN12QTextDecoder9toUnicodeERK10QByteArray(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN12QTextDecoder9toUnicodeEPKci
    // invoke: QString toUnicode(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN12QTextDecoder9toUnicodeEPKci(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN12QTextDecoder9toUnicodeEP7QStringPKci
    // invoke: void toUnicode(class QString *, const char *, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN12QTextDecoder9toUnicodeEP7QStringPKci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTextDecoder", "toUnicode", args)
  }

}

// <= body block end

