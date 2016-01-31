package qt5
// auto generated, do not modify.
// created: Sun Jan 31 12:22:27 2016
// src-file: /QtCore/qtextstream.h
// dst-file: /src/core/qtextstream.go
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
  // proto:  void QTextStreamManipulator::QTextStreamManipulator(QTSMFI m, int a);
extern void C_ZN22QTextStreamManipulatorC2EM11QTextStreamFviEi(void* qthis, void* arg0, int32_t arg1); // 1
  // proto:  void QTextStreamManipulator::QTextStreamManipulator(QTSMFC m, QChar c);
extern void C_ZN22QTextStreamManipulatorC2EM11QTextStreamFv5QCharES1_(void* qthis, void* arg0, void* arg1); // 1
  // proto:  void QTextStreamManipulator::exec(QTextStream & s);
extern void C_ZN22QTextStreamManipulator4execER11QTextStream(void* qthis, void* arg0); // 2
  // proto:  bool QTextStream::autoDetectUnicode();
extern void C_ZNK11QTextStream17autoDetectUnicodeEv(void* qthis); // 4
  // proto:  QLocale QTextStream::locale();
extern void C_ZNK11QTextStream6localeEv(void* qthis); // 4
  // proto:  void QTextStream::skipWhiteSpace();
extern void C_ZN11QTextStream14skipWhiteSpaceEv(void* qthis); // 4
  // proto:  qint64 QTextStream::pos();
extern void C_ZNK11QTextStream3posEv(void* qthis); // 4
  // proto:  void QTextStream::setRealNumberPrecision(int precision);
extern void C_ZN11QTextStream22setRealNumberPrecisionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextStream::flush();
extern void C_ZN11QTextStream5flushEv(void* qthis); // 4
  // proto:  bool QTextStream::seek(qint64 pos);
extern void C_ZN11QTextStream4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  QTextStream::FieldAlignment QTextStream::fieldAlignment();
extern void C_ZNK11QTextStream14fieldAlignmentEv(void* qthis); // 4
  // proto:  int QTextStream::fieldWidth();
extern void C_ZNK11QTextStream10fieldWidthEv(void* qthis); // 4
  // proto:  void QTextStream::setAutoDetectUnicode(bool enabled);
extern void C_ZN11QTextStream20setAutoDetectUnicodeEb(void* qthis, bool arg0); // 4
  // proto:  NumberFlags QTextStream::numberFlags();
extern void C_ZNK11QTextStream11numberFlagsEv(void* qthis); // 4
  // proto:  bool QTextStream::generateByteOrderMark();
extern void C_ZNK11QTextStream21generateByteOrderMarkEv(void* qthis); // 4
  // proto:  QTextCodec * QTextStream::codec();
extern void C_ZNK11QTextStream5codecEv(void* qthis); // 4
  // proto:  void QTextStream::setCodec(QTextCodec * codec);
extern void C_ZN11QTextStream8setCodecEP10QTextCodec(void* qthis, void* arg0); // 4
  // proto:  void QTextStream::setCodec(const char * codecName);
extern void C_ZN11QTextStream8setCodecEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  int QTextStream::realNumberPrecision();
extern void C_ZNK11QTextStream19realNumberPrecisionEv(void* qthis); // 4
  // proto:  void QTextStream::setDevice(QIODevice * device);
extern void C_ZN11QTextStream9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QChar QTextStream::padChar();
extern void C_ZNK11QTextStream7padCharEv(void* qthis); // 4
  // proto:  void QTextStream::QTextStream(QIODevice * device);
extern void C_ZN11QTextStreamC2EP9QIODevice(void* qthis, void* arg0); // 3
  // proto:  void QTextStream::QTextStream();
extern void C_ZN11QTextStreamC2Ev(void* qthis); // 3
  // proto:  QTextStream::Status QTextStream::status();
extern void C_ZNK11QTextStream6statusEv(void* qthis); // 4
  // proto:  QString * QTextStream::string();
extern void C_ZNK11QTextStream6stringEv(void* qthis); // 4
  // proto:  void QTextStream::setGenerateByteOrderMark(bool generate);
extern void C_ZN11QTextStream24setGenerateByteOrderMarkEb(void* qthis, bool arg0); // 4
  // proto:  QString QTextStream::read(qint64 maxlen);
extern void C_ZN11QTextStream4readEx(void* qthis, int64_t arg0); // 4
  // proto:  void QTextStream::~QTextStream();
extern void C_ZN11QTextStreamD2Ev(void* qthis); // 4
  // proto:  QTextStream::RealNumberNotation QTextStream::realNumberNotation();
extern void C_ZNK11QTextStream18realNumberNotationEv(void* qthis); // 4
  // proto:  void QTextStream::setPadChar(QChar ch);
extern void C_ZN11QTextStream10setPadCharE5QChar(void* qthis, void* arg0); // 4
  // proto:  void QTextStream::setFieldWidth(int width);
extern void C_ZN11QTextStream13setFieldWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QIODevice * QTextStream::device();
extern void C_ZNK11QTextStream6deviceEv(void* qthis); // 4
  // proto:  QString QTextStream::readLine(qint64 maxlen);
extern void C_ZN11QTextStream8readLineEx(void* qthis, int64_t arg0); // 4
  // proto:  void QTextStream::reset();
extern void C_ZN11QTextStream5resetEv(void* qthis); // 4
  // proto:  QString QTextStream::readAll();
extern void C_ZN11QTextStream7readAllEv(void* qthis); // 4
  // proto:  bool QTextStream::readLineInto(QString * line, qint64 maxlen);
extern void C_ZN11QTextStream12readLineIntoEP7QStringx(void* qthis, void* arg0, int64_t arg1); // 4
  // proto:  void QTextStream::setLocale(const QLocale & locale);
extern void C_ZN11QTextStream9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  void QTextStream::resetStatus();
extern void C_ZN11QTextStream11resetStatusEv(void* qthis); // 4
  // proto:  bool QTextStream::atEnd();
extern void C_ZNK11QTextStream5atEndEv(void* qthis); // 4
  // proto:  void QTextStream::setIntegerBase(int base);
extern void C_ZN11QTextStream14setIntegerBaseEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextStream::integerBase();
extern void C_ZNK11QTextStream11integerBaseEv(void* qthis); // 4
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

// class sizeof(QTextStreamManipulator)=40
type QTextStreamManipulator struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QTextStream)=1
type QTextStream struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// QTextStreamManipulator(QTSMFI, int)
func NewQTextStreamManipulator(args ...interface{}) QTextStreamManipulator {
  // QTextStreamManipulator(QTSMFI, int)
  // QTextStreamManipulator(QTSMFC, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.VoidpTy() // "QTSMFI"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.VoidpTy() // "QTSMFC"
  vtys[1][1] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN22QTextStreamManipulatorC1EM11QTextStreamFviEi
    // invoke: void QTextStreamManipulator(QTSMFI, int)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN22QTextStreamManipulatorC2EM11QTextStreamFviEi(qthis, arg0, arg1)
  case 1:
    // invoke: _ZN22QTextStreamManipulatorC1EM11QTextStreamFv5QCharES1_
    // invoke: void QTextStreamManipulator(QTSMFC, class QChar)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN22QTextStreamManipulatorC2EM11QTextStreamFv5QCharES1_(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextStreamManipulator", "QTextStreamManipulator", args)
  }

  return QTextStreamManipulator{}
}

// exec(class QTextStream &)
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
    // invoke: void exec(class QTextStream &)
    var arg0 = args[0].(QTextStream).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN22QTextStreamManipulator4execER11QTextStream(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStreamManipulator", "exec", args)
  }

}

// autoDetectUnicode()
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
    // invoke: bool autoDetectUnicode()
    C.C_ZNK11QTextStream17autoDetectUnicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "autoDetectUnicode", args)
  }

}

// locale()
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
    // invoke: QLocale locale()
    C.C_ZNK11QTextStream6localeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "locale", args)
  }

}

// skipWhiteSpace()
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
    // invoke: void skipWhiteSpace()
    C.C_ZN11QTextStream14skipWhiteSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "skipWhiteSpace", args)
  }

}

// pos()
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
    // invoke: qint64 pos()
    C.C_ZNK11QTextStream3posEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "pos", args)
  }

}

// setRealNumberPrecision(int)
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
    // invoke: void setRealNumberPrecision(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream22setRealNumberPrecisionEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setRealNumberPrecision", args)
  }

}

// flush()
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
    // invoke: void flush()
    C.C_ZN11QTextStream5flushEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "flush", args)
  }

}

// seek(qint64)
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
    // invoke: bool seek(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream4seekEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "seek", args)
  }

}

// fieldAlignment()
func (this *QTextStream) fieldAlignment(args ...interface{}) () {
  // fieldAlignment()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream14fieldAlignmentEv
    // invoke: QTextStream::FieldAlignment fieldAlignment()
    C.C_ZNK11QTextStream14fieldAlignmentEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "fieldAlignment", args)
  }

}

// fieldWidth()
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
    // invoke: int fieldWidth()
    C.C_ZNK11QTextStream10fieldWidthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "fieldWidth", args)
  }

}

// setAutoDetectUnicode(_Bool)
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
    // invoke: void setAutoDetectUnicode(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream20setAutoDetectUnicodeEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setAutoDetectUnicode", args)
  }

}

// numberFlags()
func (this *QTextStream) numberFlags(args ...interface{}) () {
  // numberFlags()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream11numberFlagsEv
    // invoke: NumberFlags numberFlags()
    C.C_ZNK11QTextStream11numberFlagsEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "numberFlags", args)
  }

}

// generateByteOrderMark()
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
    // invoke: bool generateByteOrderMark()
    C.C_ZNK11QTextStream21generateByteOrderMarkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "generateByteOrderMark", args)
  }

}

// codec()
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
    // invoke: QTextCodec * codec()
    C.C_ZNK11QTextStream5codecEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "codec", args)
  }

}

// setCodec(class QTextCodec *)
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
    // invoke: void setCodec(class QTextCodec *)
    var arg0 = args[0].(QTextCodec).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream8setCodecEP10QTextCodec(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN11QTextStream8setCodecEPKc
    // invoke: void setCodec(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream8setCodecEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setCodec", args)
  }

}

// realNumberPrecision()
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
    // invoke: int realNumberPrecision()
    C.C_ZNK11QTextStream19realNumberPrecisionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "realNumberPrecision", args)
  }

}

// setDevice(class QIODevice *)
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
    // invoke: void setDevice(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream9setDeviceEP9QIODevice(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setDevice", args)
  }

}

// padChar()
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
    // invoke: QChar padChar()
    C.C_ZNK11QTextStream7padCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "padChar", args)
  }

}

// QTextStream(class QIODevice *)
func NewQTextStream(args ...interface{}) QTextStream {
  // QTextStream(class QIODevice *)
  // QTextStream()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QIODevice{}) // "QIODevice *"
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStreamC1EP9QIODevice
    // invoke: void QTextStream(class QIODevice *)
    var arg0 = args[0].(QIODevice).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextStreamC2EP9QIODevice(qthis, arg0)
  case 1:
    // invoke: _ZN11QTextStreamC1Ev
    // invoke: void QTextStream()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C.C_ZN11QTextStreamC2Ev(qthis)
  default:
    qtrt.ErrorResolve("QTextStream", "QTextStream", args)
  }

  return QTextStream{}
}

// status()
func (this *QTextStream) status(args ...interface{}) () {
  // status()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream6statusEv
    // invoke: QTextStream::Status status()
    C.C_ZNK11QTextStream6statusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "status", args)
  }

}

// string()
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
    // invoke: QString * string()
    C.C_ZNK11QTextStream6stringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "string", args)
  }

}

// setGenerateByteOrderMark(_Bool)
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
    // invoke: void setGenerateByteOrderMark(_Bool)
    var arg0 = C.bool(args[0].(bool))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream24setGenerateByteOrderMarkEb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setGenerateByteOrderMark", args)
  }

}

// read(qint64)
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
    // invoke: QString read(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream4readEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "read", args)
  }

}

// ~QTextStream()
func (this *QTextStream) FreeQTextStream(args ...interface{}) () {
  // ~QTextStream()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN11QTextStreamD0Ev
    // invoke: void ~QTextStream()
    C.C_ZN11QTextStreamD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "~QTextStream", args)
  }

}

// realNumberNotation()
func (this *QTextStream) realNumberNotation(args ...interface{}) () {
  // realNumberNotation()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK11QTextStream18realNumberNotationEv
    // invoke: QTextStream::RealNumberNotation realNumberNotation()
    C.C_ZNK11QTextStream18realNumberNotationEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "realNumberNotation", args)
  }

}

// setPadChar(class QChar)
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
    // invoke: void setPadChar(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream10setPadCharE5QChar(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setPadChar", args)
  }

}

// setFieldWidth(int)
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
    // invoke: void setFieldWidth(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream13setFieldWidthEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setFieldWidth", args)
  }

}

// device()
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
    // invoke: QIODevice * device()
    C.C_ZNK11QTextStream6deviceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "device", args)
  }

}

// readLine(qint64)
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
    // invoke: QString readLine(qint64)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream8readLineEx(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "readLine", args)
  }

}

// reset()
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
    // invoke: void reset()
    C.C_ZN11QTextStream5resetEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "reset", args)
  }

}

// readAll()
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
    // invoke: QString readAll()
    C.C_ZN11QTextStream7readAllEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "readAll", args)
  }

}

// readLineInto(class QString *, qint64)
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
    // invoke: bool readLineInto(class QString *, qint64)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int64_t(args[1].(int64))
    if false {fmt.Println(arg1)}
    C.C_ZN11QTextStream12readLineIntoEP7QStringx(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QTextStream", "readLineInto", args)
  }

}

// setLocale(const class QLocale &)
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
    // invoke: void setLocale(const class QLocale &)
    var arg0 = args[0].(QLocale).qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream9setLocaleERK7QLocale(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setLocale", args)
  }

}

// resetStatus()
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
    // invoke: void resetStatus()
    C.C_ZN11QTextStream11resetStatusEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "resetStatus", args)
  }

}

// atEnd()
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
    // invoke: bool atEnd()
    C.C_ZNK11QTextStream5atEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "atEnd", args)
  }

}

// setIntegerBase(int)
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
    // invoke: void setIntegerBase(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN11QTextStream14setIntegerBaseEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setIntegerBase", args)
  }

}

// integerBase()
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
    // invoke: int integerBase()
    C.C_ZNK11QTextStream11integerBaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QTextStream", "integerBase", args)
  }

}

// <= body block end

