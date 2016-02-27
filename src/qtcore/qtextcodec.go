package qtcore
// auto generated, do not modify.
// created: Sat Feb 27 18:05:15 2016
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
import "runtime"
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
extern void* C_ZN12QTextEncoderC2EPK10QTextCodec(void* arg0); // 1
  // proto:  bool QTextEncoder::hasFailure();
extern bool C_ZNK12QTextEncoder10hasFailureEv(void* qthis); // 4
  // proto:  QByteArray QTextEncoder::fromUnicode(const QChar * uc, int len);
extern void* C_ZN12QTextEncoder11fromUnicodeEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QByteArray QTextEncoder::fromUnicode(const QString & str);
extern void* C_ZN12QTextEncoder11fromUnicodeERK7QString(void* qthis, void* arg0); // 4
  // proto:  void QTextEncoder::~QTextEncoder();
extern void C_ZN12QTextEncoderD2Ev(void* qthis); // 4
  // proto: static void QTextCodec::setCodecForLocale(QTextCodec * c);
extern void C_ZN10QTextCodec17setCodecForLocaleEPS_(void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForName(const QByteArray & name);
extern void* C_ZN10QTextCodec12codecForNameERK10QByteArray(void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForName(const char * name);
extern void* C_ZN10QTextCodec12codecForNameEPKc(void* arg0); // 2
  // proto:  bool QTextCodec::canEncode(QChar );
extern bool C_ZNK10QTextCodec9canEncodeE5QChar(void* qthis, void* arg0); // 4
  // proto:  bool QTextCodec::canEncode(const QString & );
extern bool C_ZNK10QTextCodec9canEncodeERK7QString(void* qthis, void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForLocale();
extern void* C_ZN10QTextCodec14codecForLocaleEv(); // 4
  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba);
extern void* C_ZN10QTextCodec12codecForHtmlERK10QByteArray(void* arg0); // 4
  // proto: static QTextCodec * QTextCodec::codecForHtml(const QByteArray & ba, QTextCodec * defaultCodec);
extern void* C_ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_(void* arg0, void* arg1); // 4
  // proto: static QList<QByteArray> QTextCodec::availableCodecs();
extern void C_ZN10QTextCodec15availableCodecsEv(); // 4
  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba, QTextCodec * defaultCodec);
extern void* C_ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_(void* arg0, void* arg1); // 4
  // proto: static QTextCodec * QTextCodec::codecForUtfText(const QByteArray & ba);
extern void* C_ZN10QTextCodec15codecForUtfTextERK10QByteArray(void* arg0); // 4
  // proto: static QList<int> QTextCodec::availableMibs();
extern void C_ZN10QTextCodec13availableMibsEv(); // 4
  // proto:  QByteArray QTextCodec::fromUnicode(const QString & uc);
extern void* C_ZNK10QTextCodec11fromUnicodeERK7QString(void* qthis, void* arg0); // 4
  // proto:  QList<QByteArray> QTextCodec::aliases();
extern void C_ZNK10QTextCodec7aliasesEv(void* qthis); // 4
  // proto: static QTextCodec * QTextCodec::codecForMib(int mib);
extern void* C_ZN10QTextCodec11codecForMibEi(int32_t arg0); // 4
  // proto:  QString QTextCodec::toUnicode(const QByteArray & );
extern void* C_ZNK10QTextCodec9toUnicodeERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QString QTextCodec::toUnicode(const char * chars);
extern void* C_ZNK10QTextCodec9toUnicodeEPKc(void* qthis, void* arg0); // 4
  // proto:  void QTextDecoder::QTextDecoder(const QTextCodec * codec);
extern void* C_ZN12QTextDecoderC2EPK10QTextCodec(void* arg0); // 1
  // proto:  bool QTextDecoder::hasFailure();
extern bool C_ZNK12QTextDecoder10hasFailureEv(void* qthis); // 4
  // proto:  void QTextDecoder::~QTextDecoder();
extern void C_ZN12QTextDecoderD2Ev(void* qthis); // 4
  // proto:  QString QTextDecoder::toUnicode(const QByteArray & ba);
extern void* C_ZN12QTextDecoder9toUnicodeERK10QByteArray(void* qthis, void* arg0); // 4
  // proto:  QString QTextDecoder::toUnicode(const char * chars, int len);
extern void* C_ZN12QTextDecoder9toUnicodeEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QTextDecoder::toUnicode(QString * target, const char * chars, int len);
extern void C_ZN12QTextDecoder9toUnicodeEP7QStringPKci(void* qthis, void* arg0, void* arg1, int32_t arg2); // 4
*/
import "C"
// } // <= ext block end

// body block begin =>
func init() {
  if false {qtrt.KeepMe()}
  if false {fmt.Println(123)}
  if false {reflect.TypeOf(123)}
  if false {reflect.TypeOf(unsafe.Sizeof(0))}
  if false {reflect.TypeOf(runtime.Version)}
}

// class sizeof(QTextEncoder)=1
type QTextEncoder struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextCodec)=8
type QTextCodec struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextDecoder)=1
type QTextDecoder struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// QTextEncoder(const class QTextCodec *)
func GcfreeQTextEncoder(this *QTextEncoder) {
  qtrt.UniverseFree(this)
}
func NewQTextEncoder(args ...interface{}) *QTextEncoder {
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
    var arg0 = args[0].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTextEncoderC2EPK10QTextCodec(arg0)
    this := &QTextEncoder{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTextEncoder)
    return this
  default:
    qtrt.ErrorResolve("QTextEncoder", "QTextEncoder", args)
  }

  return nil // QTextEncoder{Qclsinst:qthis}
}

// hasFailure()
func (this *QTextEncoder) HasFailure(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTextEncoder10hasFailureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEncoder", "hasFailure", args)
  }

  return
}

// fromUnicode(const class QChar *, int)
func (this *QTextEncoder) FromUnicode(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QTextEncoder11fromUnicodeEPK5QChari(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QTextEncoder11fromUnicodeERK7QString
    // invoke: QByteArray fromUnicode(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QTextEncoder11fromUnicodeERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextEncoder", "fromUnicode", args)
  }

  return
}

// ~QTextEncoder()
func (this *QTextEncoder) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN12QTextEncoderD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTextEncoder", "~QTextEncoder", args)
  }

  return
}

// setCodecForLocale(class QTextCodec *)
func (this *QTextCodec) SetCodecForLocale_s(args ...interface{}) () {
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
    var arg0 = args[0].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QTextCodec17setCodecForLocaleEPS_(arg0)
  default:
    qtrt.ErrorResolve("QTextCodec", "setCodecForLocale", args)
  }

  return
}

// codecForName(const class QByteArray &)
func (this *QTextCodec) CodecForName_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QTextCodec12codecForNameERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QTextCodec12codecForNameEPKc
    // invoke: QTextCodec * codecForName(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN10QTextCodec12codecForNameEPKc(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForName", args)
  }

  return
}

// canEncode(class QChar)
func (this *QTextCodec) CanEncode(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextCodec9canEncodeE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QTextCodec9canEncodeERK7QString
    // invoke: bool canEncode(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextCodec9canEncodeERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "canEncode", args)
  }

  return
}

// codecForLocale()
func (this *QTextCodec) CodecForLocale_s(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN10QTextCodec14codecForLocaleEv()
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForLocale", args)
  }

  return
}

// codecForHtml(const class QByteArray &)
func (this *QTextCodec) CodecForHtml_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QTextCodec12codecForHtmlERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_
    // invoke: QTextCodec * codecForHtml(const class QByteArray &, class QTextCodec *)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForHtml", args)
  }

  return
}

// availableCodecs()
func (this *QTextCodec) AvailableCodecs_s(args ...interface{}) () {
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
    C.C_ZN10QTextCodec15availableCodecsEv()
  default:
    qtrt.ErrorResolve("QTextCodec", "availableCodecs", args)
  }

  return
}

// codecForUtfText(const class QByteArray &, class QTextCodec *)
func (this *QTextCodec) CodecForUtfText_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QTextCodec15codecForUtfTextERK10QByteArray
    // invoke: QTextCodec * codecForUtfText(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QTextCodec15codecForUtfTextERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForUtfText", args)
  }

  return
}

// availableMibs()
func (this *QTextCodec) AvailableMibs_s(args ...interface{}) () {
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
    C.C_ZN10QTextCodec13availableMibsEv()
  default:
    qtrt.ErrorResolve("QTextCodec", "availableMibs", args)
  }

  return
}

// fromUnicode(const class QString &)
func (this *QTextCodec) FromUnicode(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextCodec11fromUnicodeERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "fromUnicode", args)
  }

  return
}

// aliases()
func (this *QTextCodec) Aliases(args ...interface{}) () {
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
    C.C_ZNK10QTextCodec7aliasesEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QTextCodec", "aliases", args)
  }

  return
}

// codecForMib(int)
func (this *QTextCodec) CodecForMib_s(args ...interface{}) (ret interface{}) {
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
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QTextCodec11codecForMibEi(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForMib", args)
  }

  return
}

// toUnicode(const class QByteArray &)
func (this *QTextCodec) ToUnicode(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QTextCodec9toUnicodeERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QTextCodec9toUnicodeEPKc
    // invoke: QString toUnicode(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK10QTextCodec9toUnicodeEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextCodec", "toUnicode", args)
  }

  return
}

// QTextDecoder(const class QTextCodec *)
func GcfreeQTextDecoder(this *QTextDecoder) {
  qtrt.UniverseFree(this)
}
func NewQTextDecoder(args ...interface{}) *QTextDecoder {
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
    var arg0 = args[0].(*QTextCodec).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN12QTextDecoderC2EPK10QTextCodec(arg0)
    this := &QTextDecoder{Qclsinst:qthis}
    runtime.SetFinalizer(this, GcfreeQTextDecoder)
    return this
  default:
    qtrt.ErrorResolve("QTextDecoder", "QTextDecoder", args)
  }

  return nil // QTextDecoder{Qclsinst:qthis}
}

// hasFailure()
func (this *QTextDecoder) HasFailure(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK12QTextDecoder10hasFailureEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QTextDecoder", "hasFailure", args)
  }

  return
}

// ~QTextDecoder()
func (this *QTextDecoder) Free(args ...interface{}) () {
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
    if this != nil && this.Qclsinst != nil {
      C.C_ZN12QTextDecoderD2Ev(this.Qclsinst)
      this.Qclsinst = nil
    }
  default:
    qtrt.ErrorResolve("QTextDecoder", "~QTextDecoder", args)
  }

  return
}

// toUnicode(const class QByteArray &)
func (this *QTextDecoder) ToUnicode(args ...interface{}) (ret interface{}) {
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
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN12QTextDecoder9toUnicodeERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN12QTextDecoder9toUnicodeEPKci
    // invoke: QString toUnicode(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN12QTextDecoder9toUnicodeEPKci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN12QTextDecoder9toUnicodeEP7QStringPKci
    // invoke: void toUnicode(class QString *, const char *, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[2][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    C.C_ZN12QTextDecoder9toUnicodeEP7QStringPKci(this.Qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QTextDecoder", "toUnicode", args)
  }

  return
}

// <= body block end

