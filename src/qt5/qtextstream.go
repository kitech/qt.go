package qt5
// auto generated, do not modify.
// created: Sat Jan  2 00:56:29 2016
// src-file: /QtCore/qtextstream.h
// dst-file: /src/core/qtextstream.go
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
// class sizeof(QTextStreamManipulator)=40
type QTextStreamManipulator struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}

// class sizeof(QTextStream)=1
type QTextStream struct {
  // qbase: None;
  qclsinst uint64 /* *mut c_void*/;
}


func (this *QTextStreamManipulator) exec(args ...interface{}) () {
  // exec(class QTextStream &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextStream{}) // "QTextStream &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN22QTextStreamManipulator4execER11QTextStream
  default:
    qtrt.ErrorResolve("QTextStreamManipulator", "exec", args)
 }

}


func NewQTextStreamManipulator(args ...interface{})() {
}


func (this *QTextStream) codec(args ...interface{}) () {
  // codec()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream5codecEv
  default:
    qtrt.ErrorResolve("QTextStream", "codec", args)
 }

}


func NewQTextStream(args ...interface{})() {
}


func (this *QTextStream) setLocale(args ...interface{}) () {
  // setLocale(const class QLocale &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLocale{}) // "const QLocale &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream9setLocaleERK7QLocale
  default:
    qtrt.ErrorResolve("QTextStream", "setLocale", args)
 }

}


func (this *QTextStream) atEnd(args ...interface{}) () {
  // atEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream5atEndEv
  default:
    qtrt.ErrorResolve("QTextStream", "atEnd", args)
 }

}


func (this *QTextStream) readLineInto(args ...interface{}) () {
  // readLineInto(class QString *, qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString *"
  vtys[0][1] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream12readLineIntoEP7QStringx
  default:
    qtrt.ErrorResolve("QTextStream", "readLineInto", args)
 }

}


func (this *QTextStream) setRealNumberPrecision(args ...interface{}) () {
  // setRealNumberPrecision(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream22setRealNumberPrecisionEi
  default:
    qtrt.ErrorResolve("QTextStream", "setRealNumberPrecision", args)
 }

}


func (this *QTextStream) setDevice(args ...interface{}) () {
  // setDevice(class QIODevice *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream9setDeviceEP9QIODevice
  default:
    qtrt.ErrorResolve("QTextStream", "setDevice", args)
 }

}


func (this *QTextStream) reset(args ...interface{}) () {
  // reset()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream5resetEv
  default:
    qtrt.ErrorResolve("QTextStream", "reset", args)
 }

}


func (this *QTextStream) seek(args ...interface{}) () {
  // seek(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream4seekEx
  default:
    qtrt.ErrorResolve("QTextStream", "seek", args)
 }

}


func (this *QTextStream) string(args ...interface{}) () {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6stringEv
  default:
    qtrt.ErrorResolve("QTextStream", "string", args)
 }

}


func (this *QTextStream) setAutoDetectUnicode(args ...interface{}) () {
  // setAutoDetectUnicode(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream20setAutoDetectUnicodeEb
  default:
    qtrt.ErrorResolve("QTextStream", "setAutoDetectUnicode", args)
 }

}


func (this *QTextStream) padChar(args ...interface{}) () {
  // padChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream7padCharEv
  default:
    qtrt.ErrorResolve("QTextStream", "padChar", args)
 }

}


func (this *QTextStream) device(args ...interface{}) () {
  // device()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6deviceEv
  default:
    qtrt.ErrorResolve("QTextStream", "device", args)
 }

}


func (this *QTextStream) resetStatus(args ...interface{}) () {
  // resetStatus()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream11resetStatusEv
  default:
    qtrt.ErrorResolve("QTextStream", "resetStatus", args)
 }

}


func (this *QTextStream) autoDetectUnicode(args ...interface{}) () {
  // autoDetectUnicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream17autoDetectUnicodeEv
  default:
    qtrt.ErrorResolve("QTextStream", "autoDetectUnicode", args)
 }

}


func (this *QTextStream) fieldWidth(args ...interface{}) () {
  // fieldWidth()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream10fieldWidthEv
  default:
    qtrt.ErrorResolve("QTextStream", "fieldWidth", args)
 }

}


func (this *QTextStream) generateByteOrderMark(args ...interface{}) () {
  // generateByteOrderMark()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream21generateByteOrderMarkEv
  default:
    qtrt.ErrorResolve("QTextStream", "generateByteOrderMark", args)
 }

}


func (this *QTextStream) setGenerateByteOrderMark(args ...interface{}) () {
  // setGenerateByteOrderMark(_Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream24setGenerateByteOrderMarkEb
  default:
    qtrt.ErrorResolve("QTextStream", "setGenerateByteOrderMark", args)
 }

}


func (this *QTextStream) setCodec(args ...interface{}) () {
  // setCodec(class QTextCodec *)
  // setCodec(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream8setCodecEP10QTextCodec
  case 1:
    // invoke: _ZN11QTextStream8setCodecEPKc
  default:
    qtrt.ErrorResolve("QTextStream", "setCodec", args)
 }

}


func (this *QTextStream) flush(args ...interface{}) () {
  // flush()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream5flushEv
  default:
    qtrt.ErrorResolve("QTextStream", "flush", args)
 }

}


func (this *QTextStream) setIntegerBase(args ...interface{}) () {
  // setIntegerBase(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream14setIntegerBaseEi
  default:
    qtrt.ErrorResolve("QTextStream", "setIntegerBase", args)
 }

}


func (this *QTextStream) FreeQTextStream(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  default:
    qtrt.ErrorResolve("QTextStream", "~QTextStream", args)
 }

}


func (this *QTextStream) locale(args ...interface{}) () {
  // locale()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6localeEv
  default:
    qtrt.ErrorResolve("QTextStream", "locale", args)
 }

}


func (this *QTextStream) read(args ...interface{}) () {
  // read(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream4readEx
  default:
    qtrt.ErrorResolve("QTextStream", "read", args)
 }

}


func (this *QTextStream) setPadChar(args ...interface{}) () {
  // setPadChar(class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream10setPadCharE5QChar
  default:
    qtrt.ErrorResolve("QTextStream", "setPadChar", args)
 }

}


func (this *QTextStream) realNumberPrecision(args ...interface{}) () {
  // realNumberPrecision()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream19realNumberPrecisionEv
  default:
    qtrt.ErrorResolve("QTextStream", "realNumberPrecision", args)
 }

}


func (this *QTextStream) pos(args ...interface{}) () {
  // pos()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream3posEv
  default:
    qtrt.ErrorResolve("QTextStream", "pos", args)
 }

}


func (this *QTextStream) readAll(args ...interface{}) () {
  // readAll()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream7readAllEv
  default:
    qtrt.ErrorResolve("QTextStream", "readAll", args)
 }

}


func (this *QTextStream) skipWhiteSpace(args ...interface{}) () {
  // skipWhiteSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream14skipWhiteSpaceEv
  default:
    qtrt.ErrorResolve("QTextStream", "skipWhiteSpace", args)
 }

}


func (this *QTextStream) setFieldWidth(args ...interface{}) () {
  // setFieldWidth(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream13setFieldWidthEi
  default:
    qtrt.ErrorResolve("QTextStream", "setFieldWidth", args)
 }

}


func (this *QTextStream) integerBase(args ...interface{}) () {
  // integerBase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream11integerBaseEv
  default:
    qtrt.ErrorResolve("QTextStream", "integerBase", args)
 }

}


func (this *QTextStream) readLine(args ...interface{}) () {
  // readLine(qint64)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qint64"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
 switch matched_index {
  case 0:
    // invoke: _ZN11QTextStream8readLineEx
  default:
    qtrt.ErrorResolve("QTextStream", "readLine", args)
 }

}

// <= body block end

