package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qtextcodec.h
// dst-file: /src/core/qtextcodec.go
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
// class sizeof(QTextEncoder)=1
type QTextEncoder struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextCodec)=8
type QTextCodec struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextDecoder)=1
type QTextDecoder struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextEncoder) FreeQTextEncoder(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextEncoder", "~QTextEncoder", args)
 }

}


func (this *QTextEncoder) fromUnicode(args ...interface{}) () {
  // fromUnicode(const class QString &)
  // fromUnicode(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTextEncoder11fromUnicodeERK7QString
  case 1:
    // invoke: _ZN12QTextEncoder11fromUnicodeEPK5QChari
  default:
    qtrt.ErrorResolve("QTextEncoder", "fromUnicode", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextEncoder", "hasFailure", args)
 }

}


func NewQTextEncoder(args ...interface{})() {
}


func (this *QTextCodec) name(args ...interface{}) () {
  // name()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec4nameEv
  default:
    qtrt.ErrorResolve("QTextCodec", "name", args)
 }

}


func (this *QTextCodec) toUnicode(args ...interface{}) () {
  // toUnicode(const char *, int, struct QTextCodec::ConverterState *)
  // toUnicode(const class QByteArray &)
  // toUnicode(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  // vtys[0][2] = reflect.TypeOf(QTextCodec::ConverterState{}) // "QTextCodec::ConverterState *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec9toUnicodeEPKciPNS_14ConverterStateE
  case 1:
    // invoke: _ZNK10QTextCodec9toUnicodeERK10QByteArray
  case 2:
    // invoke: _ZNK10QTextCodec9toUnicodeEPKc
  default:
    qtrt.ErrorResolve("QTextCodec", "toUnicode", args)
 }

}


func (this *QTextCodec) fromUnicode(args ...interface{}) () {
  // fromUnicode(const class QString &)
  // fromUnicode(const class QChar *, int, struct QTextCodec::ConverterState *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  // vtys[1][2] = reflect.TypeOf(QTextCodec::ConverterState{}) // "QTextCodec::ConverterState *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec11fromUnicodeERK7QString
  case 1:
    // invoke: _ZNK10QTextCodec11fromUnicodeEPK5QChariPNS_14ConverterStateE
  default:
    qtrt.ErrorResolve("QTextCodec", "fromUnicode", args)
 }

}


func (this *QTextCodec) codecForLocale_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForLocale", args)
 }

}


func (this *QTextCodec) availableMibs_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "availableMibs", args)
 }

}


func (this *QTextCodec) codecForHtml_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForHtml", args)
 }

}


func NewQTextCodec(args ...interface{})() {
}


func (this *QTextCodec) setCodecForLocale_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "setCodecForLocale", args)
 }

}


func (this *QTextCodec) codecForUtfText_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForUtfText", args)
 }

}


func (this *QTextCodec) mibEnum(args ...interface{}) () {
  // mibEnum()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec7mibEnumEv
  default:
    qtrt.ErrorResolve("QTextCodec", "mibEnum", args)
 }

}


func (this *QTextCodec) codecForName_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForName", args)
 }

}


func (this *QTextCodec) canEncode(args ...interface{}) () {
  // canEncode(const class QString &)
  // canEncode(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK10QTextCodec9canEncodeERK7QString
  case 1:
    // invoke: _ZNK10QTextCodec9canEncodeE5QChar
  default:
    qtrt.ErrorResolve("QTextCodec", "canEncode", args)
 }

}


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
  default:
    qtrt.ErrorResolve("QTextCodec", "aliases", args)
 }

}


func (this *QTextCodec) availableCodecs_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "availableCodecs", args)
 }

}


func (this *QTextCodec) FreeQTextCodec(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "~QTextCodec", args)
 }

}


func (this *QTextCodec) codecForMib_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextCodec", "codecForMib", args)
 }

}


func (this *QTextDecoder) toUnicode(args ...interface{}) () {
  // toUnicode(const char *, int)
  // toUnicode(const class QByteArray &)
  // toUnicode(class QString *, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[2][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN12QTextDecoder9toUnicodeEPKci
  case 1:
    // invoke: _ZN12QTextDecoder9toUnicodeERK10QByteArray
  case 2:
    // invoke: _ZN12QTextDecoder9toUnicodeEP7QStringPKci
  default:
    qtrt.ErrorResolve("QTextDecoder", "toUnicode", args)
 }

}


func NewQTextDecoder(args ...interface{})() {
}


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
  default:
    qtrt.ErrorResolve("QTextDecoder", "hasFailure", args)
 }

}


func (this *QTextDecoder) FreeQTextDecoder(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextDecoder", "~QTextDecoder", args)
 }

}

// <= body block end

