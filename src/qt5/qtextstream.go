package qt5
// auto generated, do not modify.
// created: Sun Jan 31 23:40:52 2016
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
extern void* C_ZN22QTextStreamManipulatorC2EM11QTextStreamFviEi(void* arg0, int32_t arg1); // 1
  // proto:  void QTextStreamManipulator::QTextStreamManipulator(QTSMFC m, QChar c);
extern void* C_ZN22QTextStreamManipulatorC2EM11QTextStreamFv5QCharES1_(void* arg0, void* arg1); // 1
  // proto:  void QTextStreamManipulator::exec(QTextStream & s);
extern void C_ZN22QTextStreamManipulator4execER11QTextStream(void* qthis, void* arg0); // 2
  // proto:  bool QTextStream::autoDetectUnicode();
extern bool C_ZNK11QTextStream17autoDetectUnicodeEv(void* qthis); // 4
  // proto:  QLocale QTextStream::locale();
extern void* C_ZNK11QTextStream6localeEv(void* qthis); // 4
  // proto:  void QTextStream::skipWhiteSpace();
extern void C_ZN11QTextStream14skipWhiteSpaceEv(void* qthis); // 4
  // proto:  qint64 QTextStream::pos();
extern int64_t C_ZNK11QTextStream3posEv(void* qthis); // 4
  // proto:  void QTextStream::setRealNumberPrecision(int precision);
extern void C_ZN11QTextStream22setRealNumberPrecisionEi(void* qthis, int32_t arg0); // 4
  // proto:  void QTextStream::flush();
extern void C_ZN11QTextStream5flushEv(void* qthis); // 4
  // proto:  bool QTextStream::seek(qint64 pos);
extern bool C_ZN11QTextStream4seekEx(void* qthis, int64_t arg0); // 4
  // proto:  QTextStream::FieldAlignment QTextStream::fieldAlignment();
extern void C_ZNK11QTextStream14fieldAlignmentEv(void* qthis); // 4
  // proto:  int QTextStream::fieldWidth();
extern int32_t C_ZNK11QTextStream10fieldWidthEv(void* qthis); // 4
  // proto:  void QTextStream::setAutoDetectUnicode(bool enabled);
extern void C_ZN11QTextStream20setAutoDetectUnicodeEb(void* qthis, bool arg0); // 4
  // proto:  NumberFlags QTextStream::numberFlags();
extern void C_ZNK11QTextStream11numberFlagsEv(void* qthis); // 4
  // proto:  bool QTextStream::generateByteOrderMark();
extern bool C_ZNK11QTextStream21generateByteOrderMarkEv(void* qthis); // 4
  // proto:  QTextCodec * QTextStream::codec();
extern void* C_ZNK11QTextStream5codecEv(void* qthis); // 4
  // proto:  void QTextStream::setCodec(QTextCodec * codec);
extern void C_ZN11QTextStream8setCodecEP10QTextCodec(void* qthis, void* arg0); // 4
  // proto:  void QTextStream::setCodec(const char * codecName);
extern void C_ZN11QTextStream8setCodecEPKc(void* qthis, void* arg0); // 4
  // proto:  int QTextStream::realNumberPrecision();
extern int32_t C_ZNK11QTextStream19realNumberPrecisionEv(void* qthis); // 4
  // proto:  void QTextStream::setDevice(QIODevice * device);
extern void C_ZN11QTextStream9setDeviceEP9QIODevice(void* qthis, void* arg0); // 4
  // proto:  QChar QTextStream::padChar();
extern void* C_ZNK11QTextStream7padCharEv(void* qthis); // 4
  // proto:  void QTextStream::QTextStream(QIODevice * device);
extern void* C_ZN11QTextStreamC2EP9QIODevice(void* arg0); // 3
  // proto:  void QTextStream::QTextStream();
extern void* C_ZN11QTextStreamC2Ev(); // 3
  // proto:  QTextStream::Status QTextStream::status();
extern void C_ZNK11QTextStream6statusEv(void* qthis); // 4
  // proto:  QString * QTextStream::string();
extern void* C_ZNK11QTextStream6stringEv(void* qthis); // 4
  // proto:  void QTextStream::setGenerateByteOrderMark(bool generate);
extern void C_ZN11QTextStream24setGenerateByteOrderMarkEb(void* qthis, bool arg0); // 4
  // proto:  QString QTextStream::read(qint64 maxlen);
extern void* C_ZN11QTextStream4readEx(void* qthis, int64_t arg0); // 4
  // proto:  void QTextStream::~QTextStream();
extern void C_ZN11QTextStreamD2Ev(void* qthis); // 4
  // proto:  QTextStream::RealNumberNotation QTextStream::realNumberNotation();
extern void C_ZNK11QTextStream18realNumberNotationEv(void* qthis); // 4
  // proto:  void QTextStream::setPadChar(QChar ch);
extern void C_ZN11QTextStream10setPadCharE5QChar(void* qthis, void* arg0); // 4
  // proto:  void QTextStream::setFieldWidth(int width);
extern void C_ZN11QTextStream13setFieldWidthEi(void* qthis, int32_t arg0); // 4
  // proto:  QIODevice * QTextStream::device();
extern void* C_ZNK11QTextStream6deviceEv(void* qthis); // 4
  // proto:  QString QTextStream::readLine(qint64 maxlen);
extern void* C_ZN11QTextStream8readLineEx(void* qthis, int64_t arg0); // 4
  // proto:  void QTextStream::reset();
extern void C_ZN11QTextStream5resetEv(void* qthis); // 4
  // proto:  QString QTextStream::readAll();
extern void* C_ZN11QTextStream7readAllEv(void* qthis); // 4
  // proto:  bool QTextStream::readLineInto(QString * line, qint64 maxlen);
extern bool C_ZN11QTextStream12readLineIntoEP7QStringx(void* qthis, void* arg0, int64_t arg1); // 4
  // proto:  void QTextStream::setLocale(const QLocale & locale);
extern void C_ZN11QTextStream9setLocaleERK7QLocale(void* qthis, void* arg0); // 4
  // proto:  void QTextStream::resetStatus();
extern void C_ZN11QTextStream11resetStatusEv(void* qthis); // 4
  // proto:  bool QTextStream::atEnd();
extern bool C_ZNK11QTextStream5atEndEv(void* qthis); // 4
  // proto:  void QTextStream::setIntegerBase(int base);
extern void C_ZN11QTextStream14setIntegerBaseEi(void* qthis, int32_t arg0); // 4
  // proto:  int QTextStream::integerBase();
extern int32_t C_ZNK11QTextStream11integerBaseEv(void* qthis); // 4
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
func NewQTextStreamManipulator(args ...interface{}) *QTextStreamManipulator {
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
    qthis = C.C_ZN22QTextStreamManipulatorC2EM11QTextStreamFviEi(arg0, arg1)
    return &QTextStreamManipulator{qclsinst:qthis}
  case 1:
    // invoke: _ZN22QTextStreamManipulatorC1EM11QTextStreamFv5QCharES1_
    // invoke: void QTextStreamManipulator(QTSMFC, class QChar)
    var arg0 = args[0].(unsafe.Pointer)
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN22QTextStreamManipulatorC2EM11QTextStreamFv5QCharES1_(arg0, arg1)
    return &QTextStreamManipulator{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextStreamManipulator", "QTextStreamManipulator", args)
  }

  return nil // QTextStreamManipulator{qclsinst:qthis}
}

// exec(class QTextStream &)
func (this *QTextStreamManipulator) Exec(args ...interface{}) () {
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

  return
}

// autoDetectUnicode()
func (this *QTextStream) Autodetectunicode(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream17autoDetectUnicodeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "autoDetectUnicode", args)
  }

  return
}

// locale()
func (this *QTextStream) Locale(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream6localeEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QLocale{}) // "QLocale"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "locale", args)
  }

  return
}

// skipWhiteSpace()
func (this *QTextStream) Skipwhitespace(args ...interface{}) () {
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

  return
}

// pos()
func (this *QTextStream) Pos(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream3posEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qint64"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "pos", args)
  }

  return
}

// setRealNumberPrecision(int)
func (this *QTextStream) Setrealnumberprecision(args ...interface{}) () {
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

  return
}

// flush()
func (this *QTextStream) Flush(args ...interface{}) () {
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

  return
}

// seek(qint64)
func (this *QTextStream) Seek(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QTextStream4seekEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "seek", args)
  }

  return
}

// fieldAlignment()
func (this *QTextStream) Fieldalignment(args ...interface{}) () {
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

  return
}

// fieldWidth()
func (this *QTextStream) Fieldwidth(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream10fieldWidthEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "fieldWidth", args)
  }

  return
}

// setAutoDetectUnicode(_Bool)
func (this *QTextStream) Setautodetectunicode(args ...interface{}) () {
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

  return
}

// numberFlags()
func (this *QTextStream) Numberflags(args ...interface{}) () {
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

  return
}

// generateByteOrderMark()
func (this *QTextStream) Generatebyteordermark(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream21generateByteOrderMarkEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "generateByteOrderMark", args)
  }

  return
}

// codec()
func (this *QTextStream) Codec(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream5codecEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QTextCodec{}) // "QTextCodec *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "codec", args)
  }

  return
}

// setCodec(class QTextCodec *)
func (this *QTextStream) Setcodec(args ...interface{}) () {
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
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN11QTextStream8setCodecEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QTextStream", "setCodec", args)
  }

  return
}

// realNumberPrecision()
func (this *QTextStream) Realnumberprecision(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream19realNumberPrecisionEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "realNumberPrecision", args)
  }

  return
}

// setDevice(class QIODevice *)
func (this *QTextStream) Setdevice(args ...interface{}) () {
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

  return
}

// padChar()
func (this *QTextStream) Padchar(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream7padCharEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "padChar", args)
  }

  return
}

// QTextStream(class QIODevice *)
func NewQTextStream(args ...interface{}) *QTextStream {
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
    qthis = C.C_ZN11QTextStreamC2EP9QIODevice(arg0)
    return &QTextStream{qclsinst:qthis}
  case 1:
    // invoke: _ZN11QTextStreamC1Ev
    // invoke: void QTextStream()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN11QTextStreamC2Ev()
    return &QTextStream{qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QTextStream", "QTextStream", args)
  }

  return nil // QTextStream{qclsinst:qthis}
}

// status()
func (this *QTextStream) Status(args ...interface{}) () {
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

  return
}

// string()
func (this *QTextStream) String(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream6stringEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "string", args)
  }

  return
}

// setGenerateByteOrderMark(_Bool)
func (this *QTextStream) Setgeneratebyteordermark(args ...interface{}) () {
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

  return
}

// read(qint64)
func (this *QTextStream) Read(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QTextStream4readEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "read", args)
  }

  return
}

// ~QTextStream()
func (this *QTextStream) Freeqtextstream(args ...interface{}) () {
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

  return
}

// realNumberNotation()
func (this *QTextStream) Realnumbernotation(args ...interface{}) () {
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

  return
}

// setPadChar(class QChar)
func (this *QTextStream) Setpadchar(args ...interface{}) () {
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

  return
}

// setFieldWidth(int)
func (this *QTextStream) Setfieldwidth(args ...interface{}) () {
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

  return
}

// device()
func (this *QTextStream) Device(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream6deviceEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QIODevice{}) // "QIODevice *"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "device", args)
  }

  return
}

// readLine(qint64)
func (this *QTextStream) Readline(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QTextStream8readLineEx(this.qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "readLine", args)
  }

  return
}

// reset()
func (this *QTextStream) Reset(args ...interface{}) () {
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

  return
}

// readAll()
func (this *QTextStream) Readall(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QTextStream7readAllEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "readAll", args)
  }

  return
}

// readLineInto(class QString *, qint64)
func (this *QTextStream) Readlineinto(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZN11QTextStream12readLineIntoEP7QStringx(this.qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "readLineInto", args)
  }

  return
}

// setLocale(const class QLocale &)
func (this *QTextStream) Setlocale(args ...interface{}) () {
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

  return
}

// resetStatus()
func (this *QTextStream) Resetstatus(args ...interface{}) () {
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

  return
}

// atEnd()
func (this *QTextStream) Atend(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream5atEndEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "atEnd", args)
  }

  return
}

// setIntegerBase(int)
func (this *QTextStream) Setintegerbase(args ...interface{}) () {
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

  return
}

// integerBase()
func (this *QTextStream) Integerbase(args ...interface{}) (ret interface{}) {
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
    var ret0 = C.C_ZNK11QTextStream11integerBaseEv(this.qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    ret = reflect.ValueOf(ret0).Convert(rety).Interface()
  default:
    qtrt.ErrorResolve("QTextStream", "integerBase", args)
  }

  return
}

// <= body block end

