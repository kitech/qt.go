package qtcore
// auto generated, do not modify.
// created: Sun Aug  7 10:49:52 2016
// src-file: /QtCore/qbytearray.h
// dst-file: /src/core/qbytearray.go
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
  // proto:  void QByteArray::detach();
extern void C_ZN10QByteArray6detachEv(void* qthis); // 2
  // proto:  void QByteArray::swap(QByteArray & other);
extern void C_ZN10QByteArray4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QByteArray QByteArray::rightJustified(int width, char fill, bool truncate);
extern void* C_ZNK10QByteArray14rightJustifiedEicb(void* qthis, int32_t arg0, unsigned char arg1, bool arg2); // 4
  // proto:  QByteArray & QByteArray::remove(int index, int len);
extern void* C_ZN10QByteArray6removeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArray::trimmed();
extern void* C_ZNKR10QByteArray7trimmedEv(void* qthis); // 2
  // proto:  const_iterator QByteArray::constEnd();
extern void* C_ZNK10QByteArray8constEndEv(void* qthis); // 2
  // proto:  bool QByteArray::isSharedWith(const QByteArray & other);
extern bool C_ZNK10QByteArray12isSharedWithERKS_(void* qthis, void* arg0); // 2
  // proto:  QByteArray QByteArray::right(int len);
extern void* C_ZNK10QByteArray5rightEi(void* qthis, int32_t arg0); // 4
  // proto:  uint QByteArray::toUInt(bool * ok, int base);
extern int32_t C_ZNK10QByteArray6toUIntEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::startsWith(const char * c);
extern bool C_ZNK10QByteArray10startsWithEPKc(void* qthis, void* arg0); // 4
  // proto:  bool QByteArray::startsWith(const QByteArray & a);
extern bool C_ZNK10QByteArray10startsWithERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QByteArray::startsWith(char c);
extern bool C_ZNK10QByteArray10startsWithEc(void* qthis, unsigned char arg0); // 4
  // proto:  int QByteArray::capacity();
extern int32_t C_ZNK10QByteArray8capacityEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::toLower();
extern void* C_ZNKR10QByteArray7toLowerEv(void* qthis); // 2
  // proto:  bool QByteArray::isNull();
extern bool C_ZNK10QByteArray6isNullEv(void* qthis); // 4
  // proto:  std::string QByteArray::toStdString();
extern void C_ZNK10QByteArray11toStdStringB5cxx11Ev(void* qthis); // 2
  // proto:  QByteArray QByteArray::toHex();
extern void* C_ZNK10QByteArray5toHexEv(void* qthis); // 4
  // proto:  int QByteArray::indexOf(const char * c, int from);
extern int32_t C_ZNK10QByteArray7indexOfEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::indexOf(char c, int from);
extern int32_t C_ZNK10QByteArray7indexOfEci(void* qthis, unsigned char arg0, int32_t arg1); // 4
  // proto:  int QByteArray::indexOf(const QByteArray & a, int from);
extern int32_t C_ZNK10QByteArray7indexOfERKS_i(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::indexOf(const QString & s, int from);
extern int32_t C_ZNK10QByteArray7indexOfERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  const_iterator QByteArray::cbegin();
extern void* C_ZNK10QByteArray6cbeginEv(void* qthis); // 2
  // proto:  void QByteArray::QByteArray();
extern void* C_ZN10QByteArrayC2Ev(); // 1
  // proto:  void QByteArray::QByteArray(const QByteArray & );
extern void* C_ZN10QByteArrayC2ERKS_(void* arg0); // 1
  // proto:  void QByteArray::QByteArray(const char * , int size);
extern void* C_ZN10QByteArrayC2EPKci(void* arg0, int32_t arg1); // 3
  // proto:  void QByteArray::QByteArray(int size, char c);
extern void* C_ZN10QByteArrayC2Eic(int32_t arg0, unsigned char arg1); // 3
  // proto: static QByteArray QByteArray::fromBase64(const QByteArray & base64);
extern void* C_ZN10QByteArray10fromBase64ERKS_(void* arg0); // 4
  // proto: static QByteArray QByteArray::fromHex(const QByteArray & hexEncoded);
extern void* C_ZN10QByteArray7fromHexERKS_(void* arg0); // 4
  // proto:  void QByteArray::chop(int n);
extern void C_ZN10QByteArray4chopEi(void* qthis, int32_t arg0); // 4
  // proto:  int QByteArray::length();
extern int32_t C_ZNK10QByteArray6lengthEv(void* qthis); // 2
  // proto:  double QByteArray::toDouble(bool * ok);
extern double C_ZNK10QByteArray8toDoubleEPb(void* qthis, void* arg0); // 4
  // proto:  long QByteArray::toLong(bool * ok, int base);
extern int32_t C_ZNK10QByteArray6toLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArray::toPercentEncoding(const QByteArray & exclude, const QByteArray & include, char percent);
extern void* C_ZNK10QByteArray17toPercentEncodingERKS_S1_c(void* qthis, void* arg0, void* arg1, unsigned char arg2); // 4
  // proto:  ushort QByteArray::toUShort(bool * ok, int base);
extern int16_t C_ZNK10QByteArray8toUShortEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  const_iterator QByteArray::constBegin();
extern void* C_ZNK10QByteArray10constBeginEv(void* qthis); // 2
  // proto:  short QByteArray::toShort(bool * ok, int base);
extern int16_t C_ZNK10QByteArray7toShortEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(qlonglong , int base);
extern void* C_ZN10QByteArray6numberExi(int64_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(qulonglong , int base);
extern void* C_ZN10QByteArray6numberEyi(int64_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(int , int base);
extern void* C_ZN10QByteArray6numberEii(int32_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(uint , int base);
extern void* C_ZN10QByteArray6numberEji(int32_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(double , char f, int prec);
extern void* C_ZN10QByteArray6numberEdci(double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QByteArray & QByteArray::replace(const QString & before, const char * after);
extern void* C_ZN10QByteArray7replaceERK7QStringPKc(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(char before, const QByteArray & after);
extern void* C_ZN10QByteArray7replaceEcRKS_(void* qthis, unsigned char arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::replace(int index, int len, const QByteArray & s);
extern void* C_ZN10QByteArray7replaceEiiRKS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QByteArray & QByteArray::replace(char before, const char * after);
extern void* C_ZN10QByteArray7replaceEcPKc(void* qthis, unsigned char arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(char before, char after);
extern void* C_ZN10QByteArray7replaceEcc(void* qthis, unsigned char arg0, unsigned char arg1); // 4
  // proto:  QByteArray & QByteArray::replace(const QByteArray & before, const QByteArray & after);
extern void* C_ZN10QByteArray7replaceERKS_S1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::replace(const QByteArray & before, const char * after);
extern void* C_ZN10QByteArray7replaceERKS_PKc(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(int index, int len, const char * s);
extern void* C_ZN10QByteArray7replaceEiiPKc(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QByteArray & QByteArray::replace(int index, int len, const char * s, int alen);
extern void* C_ZN10QByteArray7replaceEiiPKci(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  QByteArray & QByteArray::replace(const char * before, const QByteArray & after);
extern void* C_ZN10QByteArray7replaceEPKcRKS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::replace(const char * before, const char * after);
extern void* C_ZN10QByteArray7replaceEPKcS1_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(char c, const QString & after);
extern void* C_ZN10QByteArray7replaceEcRK7QString(void* qthis, unsigned char arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(const QString & before, const QByteArray & after);
extern void* C_ZN10QByteArray7replaceERK7QStringRKS_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(const char * before, int bsize, const char * after, int asize);
extern void* C_ZN10QByteArray7replaceEPKciS1_i(void* qthis, void* arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  void QByteArray::push_back(const QByteArray & a);
extern void C_ZN10QByteArray9push_backERKS_(void* qthis, void* arg0); // 2
  // proto:  void QByteArray::push_back(const char * c);
extern void C_ZN10QByteArray9push_backEPKc(void* qthis, void* arg0); // 2
  // proto:  void QByteArray::push_back(char c);
extern void C_ZN10QByteArray9push_backEc(void* qthis, unsigned char arg0); // 2
  // proto:  QByteArray QByteArray::simplified();
extern void* C_ZNKR10QByteArray10simplifiedEv(void* qthis); // 2
  // proto:  int QByteArray::size();
extern int32_t C_ZNK10QByteArray4sizeEv(void* qthis); // 2
  // proto:  bool QByteArray::contains(char c);
extern bool C_ZNK10QByteArray8containsEc(void* qthis, unsigned char arg0); // 2
  // proto:  bool QByteArray::contains(const QByteArray & a);
extern bool C_ZNK10QByteArray8containsERKS_(void* qthis, void* arg0); // 2
  // proto:  bool QByteArray::contains(const char * a);
extern bool C_ZNK10QByteArray8containsEPKc(void* qthis, void* arg0); // 2
  // proto: static QByteArray QByteArray::fromPercentEncoding(const QByteArray & pctEncoded, char percent);
extern void* C_ZN10QByteArray19fromPercentEncodingERKS_c(void* arg0, unsigned char arg1); // 4
  // proto:  const_iterator QByteArray::cend();
extern void* C_ZNK10QByteArray4cendEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::leftJustified(int width, char fill, bool truncate);
extern void* C_ZNK10QByteArray13leftJustifiedEicb(void* qthis, int32_t arg0, unsigned char arg1, bool arg2); // 4
  // proto:  QByteArray QByteArray::repeated(int times);
extern void* C_ZNK10QByteArray8repeatedEi(void* qthis, int32_t arg0); // 4
  // proto: static QByteArray QByteArray::fromRawData(const char * , int size);
extern void* C_ZN10QByteArray11fromRawDataEPKci(void* arg0, int32_t arg1); // 4
  // proto:  void QByteArray::squeeze();
extern void C_ZN10QByteArray7squeezeEv(void* qthis); // 2
  // proto:  void QByteArray::resize(int size);
extern void C_ZN10QByteArray6resizeEi(void* qthis, int32_t arg0); // 4
  // proto:  qulonglong QByteArray::toULongLong(bool * ok, int base);
extern int64_t C_ZNK10QByteArray11toULongLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::count();
extern int32_t C_ZNK10QByteArray5countEv(void* qthis); // 2
  // proto:  int QByteArray::count(const char * a);
extern int32_t C_ZNK10QByteArray5countEPKc(void* qthis, void* arg0); // 4
  // proto:  int QByteArray::count(char c);
extern int32_t C_ZNK10QByteArray5countEc(void* qthis, unsigned char arg0); // 4
  // proto:  int QByteArray::count(const QByteArray & a);
extern int32_t C_ZNK10QByteArray5countERKS_(void* qthis, void* arg0); // 4
  // proto:  void QByteArray::~QByteArray();
extern void C_ZN10QByteArrayD2Ev(void* qthis); // 2
  // proto:  float QByteArray::toFloat(bool * ok);
extern float C_ZNK10QByteArray7toFloatEPb(void* qthis, void* arg0); // 4
  // proto:  char QByteArray::at(int i);
extern unsigned char C_ZNK10QByteArray2atEi(void* qthis, int32_t arg0); // 2
  // proto:  QByteArray & QByteArray::fill(char c, int size);
extern void* C_ZN10QByteArray4fillEci(void* qthis, unsigned char arg0, int32_t arg1); // 4
  // proto:  ulong QByteArray::toULong(bool * ok, int base);
extern int32_t C_ZNK10QByteArray7toULongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  iterator QByteArray::end();
extern void* C_ZN10QByteArray3endEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::mid(int index, int len);
extern void* C_ZNK10QByteArray3midEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const char * QByteArray::constData();
extern void* C_ZNK10QByteArray9constDataEv(void* qthis); // 2
  // proto:  qlonglong QByteArray::toLongLong(bool * ok, int base);
extern int64_t C_ZNK10QByteArray10toLongLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::isEmpty();
extern bool C_ZNK10QByteArray7isEmptyEv(void* qthis); // 2
  // proto:  QList<QByteArray> QByteArray::split(char sep);
extern void C_ZNK10QByteArray5splitEc(void* qthis, unsigned char arg0); // 4
  // proto:  QByteArray & QByteArray::setRawData(const char * a, uint n);
extern void* C_ZN10QByteArray10setRawDataEPKcj(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::isDetached();
extern bool C_ZNK10QByteArray10isDetachedEv(void* qthis); // 2
  // proto:  iterator QByteArray::begin();
extern void* C_ZN10QByteArray5beginEv(void* qthis); // 2
  // proto:  int QByteArray::lastIndexOf(const char * c, int from);
extern int32_t C_ZNK10QByteArray11lastIndexOfEPKci(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::lastIndexOf(const QString & s, int from);
extern int32_t C_ZNK10QByteArray11lastIndexOfERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QByteArray::lastIndexOf(const QByteArray & a, int from);
extern int32_t C_ZNK10QByteArray11lastIndexOfERKS_i(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::lastIndexOf(char c, int from);
extern int32_t C_ZNK10QByteArray11lastIndexOfEci(void* qthis, unsigned char arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArray::toUpper();
extern void* C_ZNKR10QByteArray7toUpperEv(void* qthis); // 2
  // proto:  void QByteArray::truncate(int pos);
extern void C_ZN10QByteArray8truncateEi(void* qthis, int32_t arg0); // 4
  // proto:  QByteArray QByteArray::toBase64();
extern void* C_ZNK10QByteArray8toBase64Ev(void* qthis); // 4
  // proto:  char * QByteArray::data();
extern void* C_ZN10QByteArray4dataEv(void* qthis); // 2
  // proto:  void QByteArray::clear();
extern void C_ZN10QByteArray5clearEv(void* qthis); // 4
  // proto:  int QByteArray::toInt(bool * ok, int base);
extern int32_t C_ZNK10QByteArray5toIntEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::endsWith(const QByteArray & a);
extern bool C_ZNK10QByteArray8endsWithERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QByteArray::endsWith(const char * c);
extern bool C_ZNK10QByteArray8endsWithEPKc(void* qthis, void* arg0); // 4
  // proto:  bool QByteArray::endsWith(char c);
extern bool C_ZNK10QByteArray8endsWithEc(void* qthis, unsigned char arg0); // 4
  // proto:  QByteArray & QByteArray::setNum(uint , int base);
extern void* C_ZN10QByteArray6setNumEji(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QByteArray & QByteArray::setNum(int , int base);
extern void* C_ZN10QByteArray6setNumEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QByteArray & QByteArray::setNum(short , int base);
extern void* C_ZN10QByteArray6setNumEsi(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  QByteArray & QByteArray::setNum(double , char f, int prec);
extern void* C_ZN10QByteArray6setNumEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QByteArray & QByteArray::setNum(float , char f, int prec);
extern void* C_ZN10QByteArray6setNumEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2); // 2
  // proto:  QByteArray & QByteArray::setNum(qulonglong , int base);
extern void* C_ZN10QByteArray6setNumEyi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::setNum(qlonglong , int base);
extern void* C_ZN10QByteArray6setNumExi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::setNum(ushort , int base);
extern void* C_ZN10QByteArray6setNumEti(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  void QByteArray::reserve(int size);
extern void C_ZN10QByteArray7reserveEi(void* qthis, int32_t arg0); // 2
  // proto:  void QByteArray::push_front(char c);
extern void C_ZN10QByteArray10push_frontEc(void* qthis, unsigned char arg0); // 2
  // proto:  void QByteArray::push_front(const QByteArray & a);
extern void C_ZN10QByteArray10push_frontERKS_(void* qthis, void* arg0); // 2
  // proto:  void QByteArray::push_front(const char * c);
extern void C_ZN10QByteArray10push_frontEPKc(void* qthis, void* arg0); // 2
  // proto:  QByteArray QByteArray::left(int len);
extern void* C_ZNK10QByteArray4leftEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QByteRef)=16
type QByteRef struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QByteArray)=8
type QByteArray struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QByteArrayDataPtr)=8
type QByteArrayDataPtr struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// detach()
func (this *QByteArray) Detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6detachEv
    // invoke: void detach()
    C.C_ZN10QByteArray6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "detach", args)
  }

  return
}

// swap(class QByteArray &)
func (this *QByteArray) Swap(args ...interface{}) () {
  // swap(class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4swapERS_
    // invoke: void swap(class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "swap", args)
  }

  return
}

// rightJustified(int, char, _Bool)
func (this *QByteArray) Rightjustified(args ...interface{}) (ret interface{}) {
  // rightJustified(int, char, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray14rightJustifiedEicb
    // invoke: QByteArray rightJustified(int, char, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK10QByteArray14rightJustifiedEicb(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "rightJustified", args)
  }

  return
}

// remove(int, int)
func (this *QByteArray) Remove(args ...interface{}) (ret interface{}) {
  // remove(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6removeEii
    // invoke: QByteArray & remove(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6removeEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "remove", args)
  }

  return
}

// trimmed()
func (this *QByteArray) Trimmed(args ...interface{}) (ret interface{}) {
  // trimmed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray7trimmedEv
    // invoke: QByteArray trimmed()
    var ret0 = C.C_ZNKR10QByteArray7trimmedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "trimmed", args)
  }

  return
}

// constEnd()
func (this *QByteArray) Constend(args ...interface{}) (ret interface{}) {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8constEndEv
    // invoke: const_iterator constEnd()
    var ret0 = C.C_ZNK10QByteArray8constEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "constEnd", args)
  }

  return
}

// isSharedWith(const class QByteArray &)
func (this *QByteArray) Issharedwith(args ...interface{}) (ret interface{}) {
  // isSharedWith(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray12isSharedWithERKS_
    // invoke: bool isSharedWith(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray12isSharedWithERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "isSharedWith", args)
  }

  return
}

// right(int)
func (this *QByteArray) Right(args ...interface{}) (ret interface{}) {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5rightEi
    // invoke: QByteArray right(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray5rightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "right", args)
  }

  return
}

// toUInt(_Bool *, int)
func (this *QByteArray) Touint(args ...interface{}) (ret interface{}) {
  // toUInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6toUIntEPbi
    // invoke: uint toUInt(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray6toUIntEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toUInt", args)
  }

  return
}

// startsWith(const char *)
func (this *QByteArray) Startswith(args ...interface{}) (ret interface{}) {
  // startsWith(const char *)
  // startsWith(const class QByteArray &)
  // startsWith(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10startsWithEPKc
    // invoke: bool startsWith(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK10QByteArray10startsWithEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QByteArray10startsWithERKS_
    // invoke: bool startsWith(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray10startsWithERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QByteArray10startsWithEc
    // invoke: bool startsWith(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray10startsWithEc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "startsWith", args)
  }

  return
}

// capacity()
func (this *QByteArray) Capacity(args ...interface{}) (ret interface{}) {
  // capacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8capacityEv
    // invoke: int capacity()
    var ret0 = C.C_ZNK10QByteArray8capacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "capacity", args)
  }

  return
}

// toLower()
func (this *QByteArray) Tolower(args ...interface{}) (ret interface{}) {
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray7toLowerEv
    // invoke: QByteArray toLower()
    var ret0 = C.C_ZNKR10QByteArray7toLowerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toLower", args)
  }

  return
}

// isNull()
func (this *QByteArray) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK10QByteArray6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "isNull", args)
  }

  return
}

// toStdString()
func (this *QByteArray) Tostdstring(args ...interface{}) () {
  // toStdString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11toStdStringB5cxx11Ev
    // invoke: std::string toStdString()
    C.C_ZNK10QByteArray11toStdStringB5cxx11Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "toStdString", args)
  }

  return
}

// toHex()
func (this *QByteArray) Tohex(args ...interface{}) (ret interface{}) {
  // toHex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5toHexEv
    // invoke: QByteArray toHex()
    var ret0 = C.C_ZNK10QByteArray5toHexEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toHex", args)
  }

  return
}

// indexOf(const char *, int)
func (this *QByteArray) Indexof(args ...interface{}) (ret interface{}) {
  // indexOf(const char *, int)
  // indexOf(char, int)
  // indexOf(const class QByteArray &, int)
  // indexOf(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7indexOfEPKci
    // invoke: int indexOf(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray7indexOfEPKci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QByteArray7indexOfEci
    // invoke: int indexOf(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray7indexOfEci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QByteArray7indexOfERKS_i
    // invoke: int indexOf(const class QByteArray &, int)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray7indexOfERKS_i(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK10QByteArray7indexOfERK7QStringi
    // invoke: int indexOf(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray7indexOfERK7QStringi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "indexOf", args)
  }

  return
}

// cbegin()
func (this *QByteArray) Cbegin(args ...interface{}) (ret interface{}) {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6cbeginEv
    // invoke: const_iterator cbegin()
    var ret0 = C.C_ZNK10QByteArray6cbeginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "cbegin", args)
  }

  return
}

// QByteArray()
func NewQByteArray(args ...interface{}) *QByteArray {
  // QByteArray()
  // QByteArray(const class QByteArray &)
  // QByteArray(const char *, int)
  // QByteArray(int, char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArrayC1Ev
    // invoke: void QByteArray()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QByteArrayC2Ev()
    return &QByteArray{Qclsinst:qthis}
  case 1:
    // invoke: _ZN10QByteArrayC1ERKS_
    // invoke: void QByteArray(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QByteArrayC2ERKS_(arg0)
    return &QByteArray{Qclsinst:qthis}
  case 2:
    // invoke: _ZN10QByteArrayC1EPKci
    // invoke: void QByteArray(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QByteArrayC2EPKci(arg0, arg1)
    return &QByteArray{Qclsinst:qthis}
  case 3:
    // invoke: _ZN10QByteArrayC1Eic
    // invoke: void QByteArray(int, char)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QByteArrayC2Eic(arg0, arg1)
    return &QByteArray{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QByteArray", "QByteArray", args)
  }

  return nil // QByteArray{Qclsinst:qthis}
}

// fromBase64(const class QByteArray &)
func (this *QByteArray) Frombase64_S(args ...interface{}) (ret interface{}) {
  // fromBase64(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10fromBase64ERKS_
    // invoke: QByteArray fromBase64(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QByteArray10fromBase64ERKS_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "fromBase64", args)
  }

  return
}

// fromHex(const class QByteArray &)
func (this *QByteArray) Fromhex_S(args ...interface{}) (ret interface{}) {
  // fromHex(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7fromHexERKS_
    // invoke: QByteArray fromHex(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN10QByteArray7fromHexERKS_(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "fromHex", args)
  }

  return
}

// chop(int)
func (this *QByteArray) Chop(args ...interface{}) () {
  // chop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4chopEi
    // invoke: void chop(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray4chopEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "chop", args)
  }

  return
}

// length()
func (this *QByteArray) Length(args ...interface{}) (ret interface{}) {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6lengthEv
    // invoke: int length()
    var ret0 = C.C_ZNK10QByteArray6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "length", args)
  }

  return
}

// toDouble(_Bool *)
func (this *QByteArray) Todouble(args ...interface{}) (ret interface{}) {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toDoubleEPb
    // invoke: double toDouble(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray8toDoubleEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toDouble", args)
  }

  return
}

// toLong(_Bool *, int)
func (this *QByteArray) Tolong(args ...interface{}) (ret interface{}) {
  // toLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6toLongEPbi
    // invoke: long toLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray6toLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "long"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toLong", args)
  }

  return
}

// toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
func (this *QByteArray) Topercentencoding(args ...interface{}) (ret interface{}) {
  // toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][2] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray17toPercentEncodingERKS_S1_c
    // invoke: QByteArray toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.uchar(args[2].(byte))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK10QByteArray17toPercentEncodingERKS_S1_c(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toPercentEncoding", args)
  }

  return
}

// toUShort(_Bool *, int)
func (this *QByteArray) Toushort(args ...interface{}) (ret interface{}) {
  // toUShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toUShortEPbi
    // invoke: ushort toUShort(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray8toUShortEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toUShort", args)
  }

  return
}

// constBegin()
func (this *QByteArray) Constbegin(args ...interface{}) (ret interface{}) {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10constBeginEv
    // invoke: const_iterator constBegin()
    var ret0 = C.C_ZNK10QByteArray10constBeginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "constBegin", args)
  }

  return
}

// toShort(_Bool *, int)
func (this *QByteArray) Toshort(args ...interface{}) (ret interface{}) {
  // toShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toShortEPbi
    // invoke: short toShort(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray7toShortEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "short"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toShort", args)
  }

  return
}

// number(qlonglong, int)
func (this *QByteArray) Number_S(args ...interface{}) (ret interface{}) {
  // number(qlonglong, int)
  // number(qulonglong, int)
  // number(int, int)
  // number(uint, int)
  // number(double, char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "uint"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.DoubleTy(false) // "double"
  vtys[4][1] = qtrt.ByteTy(false) // "char"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6numberExi
    // invoke: QByteArray number(qlonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6numberExi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QByteArray6numberEyi
    // invoke: QByteArray number(qulonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6numberEyi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN10QByteArray6numberEii
    // invoke: QByteArray number(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6numberEii(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN10QByteArray6numberEji
    // invoke: QByteArray number(uint, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6numberEji(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN10QByteArray6numberEdci
    // invoke: QByteArray number(double, char, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QByteArray6numberEdci(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "number", args)
  }

  return
}

// replace(const class QString &, const char *)
func (this *QByteArray) Replace(args ...interface{}) (ret interface{}) {
  // replace(const class QString &, const char *)
  // replace(char, const class QByteArray &)
  // replace(int, int, const class QByteArray &)
  // replace(char, const char *)
  // replace(char, char)
  // replace(const class QByteArray &, const class QByteArray &)
  // replace(const class QByteArray &, const char *)
  // replace(int, int, const char *)
  // replace(int, int, const char *, int)
  // replace(const char *, const class QByteArray &)
  // replace(const char *, const char *)
  // replace(char, const class QString &)
  // replace(const class QString &, const class QByteArray &)
  // replace(const char *, int, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(false) // "char"
  vtys[3][1] = qtrt.ByteTy(true) // "const char *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(false) // "char"
  vtys[4][1] = qtrt.ByteTy(false) // "char"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[5][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[6][1] = qtrt.ByteTy(true) // "const char *"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.ByteTy(true) // "const char *"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int32Ty(false) // "int"
  vtys[8][1] = qtrt.Int32Ty(false) // "int"
  vtys[8][2] = qtrt.ByteTy(true) // "const char *"
  vtys[8][3] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(true) // "const char *"
  vtys[10][1] = qtrt.ByteTy(true) // "const char *"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.ByteTy(false) // "char"
  vtys[11][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[12][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.ByteTy(true) // "const char *"
  vtys[13][1] = qtrt.Int32Ty(false) // "int"
  vtys[13][2] = qtrt.ByteTy(true) // "const char *"
  vtys[13][3] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7replaceERK7QStringPKc
    // invoke: QByteArray & replace(const class QString &, const char *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[0][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceERK7QStringPKc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QByteArray7replaceEcRKS_
    // invoke: QByteArray & replace(char, const class QByteArray &)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceEcRKS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN10QByteArray7replaceEiiRKS_
    // invoke: QByteArray & replace(int, int, const class QByteArray &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QByteArray).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QByteArray7replaceEiiRKS_(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN10QByteArray7replaceEcPKc
    // invoke: QByteArray & replace(char, const char *)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[3][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceEcPKc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN10QByteArray7replaceEcc
    // invoke: QByteArray & replace(char, char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceEcc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZN10QByteArray7replaceERKS_S1_
    // invoke: QByteArray & replace(const class QByteArray &, const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceERKS_S1_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 6:
    // invoke: _ZN10QByteArray7replaceERKS_PKc
    // invoke: QByteArray & replace(const class QByteArray &, const char *)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[6][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceERKS_PKc(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 7:
    // invoke: _ZN10QByteArray7replaceEiiPKc
    // invoke: QByteArray & replace(int, int, const char *)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[7][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var ret0 = C.C_ZN10QByteArray7replaceEiiPKc(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 8:
    // invoke: _ZN10QByteArray7replaceEiiPKci
    // invoke: QByteArray & replace(int, int, const char *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[8][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN10QByteArray7replaceEiiPKci(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 9:
    // invoke: _ZN10QByteArray7replaceEPKcRKS_
    // invoke: QByteArray & replace(const char *, const class QByteArray &)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[9][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceEPKcRKS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 10:
    // invoke: _ZN10QByteArray7replaceEPKcS1_
    // invoke: QByteArray & replace(const char *, const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[10][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    argif1, free1 := qtrt.HandyConvert2c(args[1], vtys[10][1])
    var arg1 = argif1.(unsafe.Pointer)
    if false {fmt.Println(argif1, arg1)}
    if free1 {defer C.free(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceEPKcS1_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 11:
    // invoke: _ZN10QByteArray7replaceEcRK7QString
    // invoke: QByteArray & replace(char, const class QString &)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceEcRK7QString(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 12:
    // invoke: _ZN10QByteArray7replaceERK7QStringRKS_
    // invoke: QByteArray & replace(const class QString &, const class QByteArray &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QByteArray).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray7replaceERK7QStringRKS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 13:
    // invoke: _ZN10QByteArray7replaceEPKciS1_i
    // invoke: QByteArray & replace(const char *, int, const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[13][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    argif2, free2 := qtrt.HandyConvert2c(args[2], vtys[13][2])
    var arg2 = argif2.(unsafe.Pointer)
    if false {fmt.Println(argif2, arg2)}
    if free2 {defer C.free(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN10QByteArray7replaceEPKciS1_i(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "replace", args)
  }

  return
}

// push_back(const class QByteArray &)
func (this *QByteArray) Push_Back(args ...interface{}) () {
  // push_back(const class QByteArray &)
  // push_back(const char *)
  // push_back(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray9push_backERKS_
    // invoke: void push_back(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray9push_backERKS_(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray9push_backEPKc
    // invoke: void push_back(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN10QByteArray9push_backEPKc(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray9push_backEc
    // invoke: void push_back(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray9push_backEc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "push_back", args)
  }

  return
}

// simplified()
func (this *QByteArray) Simplified(args ...interface{}) (ret interface{}) {
  // simplified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray10simplifiedEv
    // invoke: QByteArray simplified()
    var ret0 = C.C_ZNKR10QByteArray10simplifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "simplified", args)
  }

  return
}

// size()
func (this *QByteArray) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK10QByteArray4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "size", args)
  }

  return
}

// contains(char)
func (this *QByteArray) Contains(args ...interface{}) (ret interface{}) {
  // contains(char)
  // contains(const class QByteArray &)
  // contains(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8containsEc
    // invoke: bool contains(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray8containsEc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QByteArray8containsERKS_
    // invoke: bool contains(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray8containsERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QByteArray8containsEPKc
    // invoke: bool contains(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK10QByteArray8containsEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "contains", args)
  }

  return
}

// fromPercentEncoding(const class QByteArray &, char)
func (this *QByteArray) Frompercentencoding_S(args ...interface{}) (ret interface{}) {
  // fromPercentEncoding(const class QByteArray &, char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray19fromPercentEncodingERKS_c
    // invoke: QByteArray fromPercentEncoding(const class QByteArray &, char)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray19fromPercentEncodingERKS_c(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "fromPercentEncoding", args)
  }

  return
}

// cend()
func (this *QByteArray) Cend(args ...interface{}) (ret interface{}) {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4cendEv
    // invoke: const_iterator cend()
    var ret0 = C.C_ZNK10QByteArray4cendEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "cend", args)
  }

  return
}

// leftJustified(int, char, _Bool)
func (this *QByteArray) Leftjustified(args ...interface{}) (ret interface{}) {
  // leftJustified(int, char, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray13leftJustifiedEicb
    // invoke: QByteArray leftJustified(int, char, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK10QByteArray13leftJustifiedEicb(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "leftJustified", args)
  }

  return
}

// repeated(int)
func (this *QByteArray) Repeated(args ...interface{}) (ret interface{}) {
  // repeated(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8repeatedEi
    // invoke: QByteArray repeated(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray8repeatedEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "repeated", args)
  }

  return
}

// fromRawData(const char *, int)
func (this *QByteArray) Fromrawdata_S(args ...interface{}) (ret interface{}) {
  // fromRawData(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray11fromRawDataEPKci
    // invoke: QByteArray fromRawData(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray11fromRawDataEPKci(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "fromRawData", args)
  }

  return
}

// squeeze()
func (this *QByteArray) Squeeze(args ...interface{}) () {
  // squeeze()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7squeezeEv
    // invoke: void squeeze()
    C.C_ZN10QByteArray7squeezeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "squeeze", args)
  }

  return
}

// resize(int)
func (this *QByteArray) Resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6resizeEi
    // invoke: void resize(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray6resizeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "resize", args)
  }

  return
}

// toULongLong(_Bool *, int)
func (this *QByteArray) Toulonglong(args ...interface{}) (ret interface{}) {
  // toULongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11toULongLongEPbi
    // invoke: qulonglong toULongLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray11toULongLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qulonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toULongLong", args)
  }

  return
}

// count()
func (this *QByteArray) Count(args ...interface{}) (ret interface{}) {
  // count()
  // count(const char *)
  // count(char)
  // count(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK10QByteArray5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QByteArray5countEPKc
    // invoke: int count(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK10QByteArray5countEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QByteArray5countEc
    // invoke: int count(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray5countEc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK10QByteArray5countERKS_
    // invoke: int count(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray5countERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "count", args)
  }

  return
}

// ~QByteArray()
func (this *QByteArray) Freeqbytearray(args ...interface{}) () {
  // ~QByteArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArrayD0Ev
    // invoke: void ~QByteArray()
    C.C_ZN10QByteArrayD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "~QByteArray", args)
  }

  return
}

// toFloat(_Bool *)
func (this *QByteArray) Tofloat(args ...interface{}) (ret interface{}) {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toFloatEPb
    // invoke: float toFloat(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray7toFloatEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toFloat", args)
  }

  return
}

// at(int)
func (this *QByteArray) At(args ...interface{}) (ret interface{}) {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray2atEi
    // invoke: char at(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray2atEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "char"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "at", args)
  }

  return
}

// fill(char, int)
func (this *QByteArray) Fill(args ...interface{}) (ret interface{}) {
  // fill(char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4fillEci
    // invoke: QByteArray & fill(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray4fillEci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "fill", args)
  }

  return
}

// toULong(_Bool *, int)
func (this *QByteArray) Toulong(args ...interface{}) (ret interface{}) {
  // toULong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toULongEPbi
    // invoke: ulong toULong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray7toULongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "ulong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toULong", args)
  }

  return
}

// end()
func (this *QByteArray) End(args ...interface{}) (ret interface{}) {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray3endEv
    // invoke: iterator end()
    var ret0 = C.C_ZN10QByteArray3endEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "end", args)
  }

  return
}

// mid(int, int)
func (this *QByteArray) Mid(args ...interface{}) (ret interface{}) {
  // mid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray3midEii
    // invoke: QByteArray mid(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray3midEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "mid", args)
  }

  return
}

// constData()
func (this *QByteArray) Constdata(args ...interface{}) (ret interface{}) {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray9constDataEv
    // invoke: const char * constData()
    var ret0 = C.C_ZNK10QByteArray9constDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "constData", args)
  }

  return
}

// toLongLong(_Bool *, int)
func (this *QByteArray) Tolonglong(args ...interface{}) (ret interface{}) {
  // toLongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10toLongLongEPbi
    // invoke: qlonglong toLongLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray10toLongLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qlonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toLongLong", args)
  }

  return
}

// isEmpty()
func (this *QByteArray) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK10QByteArray7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "isEmpty", args)
  }

  return
}

// split(char)
func (this *QByteArray) Split(args ...interface{}) () {
  // split(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5splitEc
    // invoke: QList<QByteArray> split(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZNK10QByteArray5splitEc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "split", args)
  }

  return
}

// setRawData(const char *, uint)
func (this *QByteArray) Setrawdata(args ...interface{}) (ret interface{}) {
  // setRawData(const char *, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10setRawDataEPKcj
    // invoke: QByteArray & setRawData(const char *, uint)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray10setRawDataEPKcj(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "setRawData", args)
  }

  return
}

// isDetached()
func (this *QByteArray) Isdetached(args ...interface{}) (ret interface{}) {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10isDetachedEv
    // invoke: bool isDetached()
    var ret0 = C.C_ZNK10QByteArray10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "isDetached", args)
  }

  return
}

// begin()
func (this *QByteArray) Begin(args ...interface{}) (ret interface{}) {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray5beginEv
    // invoke: iterator begin()
    var ret0 = C.C_ZN10QByteArray5beginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "begin", args)
  }

  return
}

// lastIndexOf(const char *, int)
func (this *QByteArray) Lastindexof(args ...interface{}) (ret interface{}) {
  // lastIndexOf(const char *, int)
  // lastIndexOf(const class QString &, int)
  // lastIndexOf(const class QByteArray &, int)
  // lastIndexOf(char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(false) // "char"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11lastIndexOfEPKci
    // invoke: int lastIndexOf(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray11lastIndexOfEPKci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QByteArray11lastIndexOfERK7QStringi
    // invoke: int lastIndexOf(const class QString &, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray11lastIndexOfERK7QStringi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QByteArray11lastIndexOfERKS_i
    // invoke: int lastIndexOf(const class QByteArray &, int)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray11lastIndexOfERKS_i(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK10QByteArray11lastIndexOfEci
    // invoke: int lastIndexOf(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray11lastIndexOfEci(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "lastIndexOf", args)
  }

  return
}

// toUpper()
func (this *QByteArray) Toupper(args ...interface{}) (ret interface{}) {
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray7toUpperEv
    // invoke: QByteArray toUpper()
    var ret0 = C.C_ZNKR10QByteArray7toUpperEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toUpper", args)
  }

  return
}

// truncate(int)
func (this *QByteArray) Truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray8truncateEi
    // invoke: void truncate(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray8truncateEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "truncate", args)
  }

  return
}

// toBase64()
func (this *QByteArray) Tobase64(args ...interface{}) (ret interface{}) {
  // toBase64()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toBase64Ev
    // invoke: QByteArray toBase64()
    var ret0 = C.C_ZNK10QByteArray8toBase64Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toBase64", args)
  }

  return
}

// data()
func (this *QByteArray) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4dataEv
    // invoke: char * data()
    var ret0 = C.C_ZN10QByteArray4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "data", args)
  }

  return
}

// clear()
func (this *QByteArray) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray5clearEv
    // invoke: void clear()
    C.C_ZN10QByteArray5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "clear", args)
  }

  return
}

// toInt(_Bool *, int)
func (this *QByteArray) Toint(args ...interface{}) (ret interface{}) {
  // toInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5toIntEPbi
    // invoke: int toInt(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QByteArray5toIntEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "toInt", args)
  }

  return
}

// endsWith(const class QByteArray &)
func (this *QByteArray) Endswith(args ...interface{}) (ret interface{}) {
  // endsWith(const class QByteArray &)
  // endsWith(const char *)
  // endsWith(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8endsWithERKS_
    // invoke: bool endsWith(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray8endsWithERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK10QByteArray8endsWithEPKc
    // invoke: bool endsWith(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK10QByteArray8endsWithEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK10QByteArray8endsWithEc
    // invoke: bool endsWith(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray8endsWithEc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "endsWith", args)
  }

  return
}

// setNum(uint, int)
func (this *QByteArray) Setnum(args ...interface{}) (ret interface{}) {
  // setNum(uint, int)
  // setNum(int, int)
  // setNum(short, int)
  // setNum(double, char, int)
  // setNum(float, char, int)
  // setNum(qulonglong, int)
  // setNum(qlonglong, int)
  // setNum(ushort, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int16Ty(false) // "short"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.DoubleTy(false) // "double"
  vtys[3][1] = qtrt.ByteTy(false) // "char"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.FloatTy(false) // "float"
  vtys[4][1] = qtrt.ByteTy(false) // "char"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6setNumEji
    // invoke: QByteArray & setNum(uint, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6setNumEji(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN10QByteArray6setNumEii
    // invoke: QByteArray & setNum(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6setNumEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN10QByteArray6setNumEsi
    // invoke: QByteArray & setNum(short, int)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6setNumEsi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN10QByteArray6setNumEdci
    // invoke: QByteArray & setNum(double, char, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QByteArray6setNumEdci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN10QByteArray6setNumEfci
    // invoke: QByteArray & setNum(float, char, int)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN10QByteArray6setNumEfci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZN10QByteArray6setNumEyi
    // invoke: QByteArray & setNum(qulonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6setNumEyi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 6:
    // invoke: _ZN10QByteArray6setNumExi
    // invoke: QByteArray & setNum(qlonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6setNumExi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 7:
    // invoke: _ZN10QByteArray6setNumEti
    // invoke: QByteArray & setNum(ushort, int)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN10QByteArray6setNumEti(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "setNum", args)
  }

  return
}

// reserve(int)
func (this *QByteArray) Reserve(args ...interface{}) () {
  // reserve(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7reserveEi
    // invoke: void reserve(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray7reserveEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "reserve", args)
  }

  return
}

// push_front(char)
func (this *QByteArray) Push_Front(args ...interface{}) () {
  // push_front(char)
  // push_front(const class QByteArray &)
  // push_front(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10push_frontEc
    // invoke: void push_front(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray10push_frontEc(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray10push_frontERKS_
    // invoke: void push_front(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN10QByteArray10push_frontERKS_(this.Qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray10push_frontEPKc
    // invoke: void push_front(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    C.C_ZN10QByteArray10push_frontEPKc(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "push_front", args)
  }

  return
}

// left(int)
func (this *QByteArray) Left(args ...interface{}) (ret interface{}) {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  var dargExists = make(map[int32]map[int32]bool)
  if false {fmt.Println(dargExists)}
  var dargValues = make(map[int32]map[int32]interface{})
  if false {fmt.Println(dargValues)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4leftEi
    // invoke: QByteArray left(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QByteArray4leftEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QByteArray", "left", args)
  }

  return
}

// <= body block end

