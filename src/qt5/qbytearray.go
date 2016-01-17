package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZN10QByteArray6detachEv(void* qthis); // 2
  // proto:  void QByteArray::swap(QByteArray & other);
extern void _ZN10QByteArray4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QByteArray QByteArray::rightJustified(int width, char fill, bool truncate);
extern void _ZNK10QByteArray14rightJustifiedEicb(void* qthis, int32_t arg0, unsigned char arg1, bool arg2); // 4
  // proto:  QByteArray & QByteArray::insert(int i, const QString & s);
extern void _ZN10QByteArray6insertEiRK7QString(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::insert(int i, const char * s);
extern void _ZN10QByteArray6insertEiPKc(void* qthis, int32_t arg0, unsigned char* arg1); // 4
  // proto:  QByteArray & QByteArray::insert(int i, const QByteArray & a);
extern void _ZN10QByteArray6insertEiRKS_(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::insert(int i, char c);
extern void _ZN10QByteArray6insertEic(void* qthis, int32_t arg0, unsigned char arg1); // 4
  // proto:  QByteArray & QByteArray::insert(int i, const char * s, int len);
extern void _ZN10QByteArray6insertEiPKci(void* qthis, int32_t arg0, unsigned char* arg1, int32_t arg2); // 4
  // proto:  QByteArray & QByteArray::remove(int index, int len);
extern void _ZN10QByteArray6removeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArray::trimmed();
extern void _ZNKR10QByteArray7trimmedEv(void* qthis); // 2
  // proto:  const_iterator QByteArray::constEnd();
extern void _ZNK10QByteArray8constEndEv(void* qthis); // 2
  // proto:  bool QByteArray::isSharedWith(const QByteArray & other);
extern void _ZNK10QByteArray12isSharedWithERKS_(void* qthis, void* arg0); // 2
  // proto:  QByteArray QByteArray::right(int len);
extern void _ZNK10QByteArray5rightEi(void* qthis, int32_t arg0); // 4
  // proto:  uint QByteArray::toUInt(bool * ok, int base);
extern void _ZNK10QByteArray6toUIntEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::append(char c);
extern void _ZN10QByteArray6appendEc(void* qthis, unsigned char arg0); // 4
  // proto:  QByteArray & QByteArray::append(const char * s, int len);
extern void _ZN10QByteArray6appendEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::append(const QByteArray & a);
extern void _ZN10QByteArray6appendERKS_(void* qthis, void* arg0); // 4
  // proto:  QByteArray & QByteArray::append(const QString & s);
extern void _ZN10QByteArray6appendERK7QString(void* qthis, void* arg0); // 2
  // proto:  QByteArray & QByteArray::append(const char * s);
extern void _ZN10QByteArray6appendEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  bool QByteArray::startsWith(const char * c);
extern void _ZNK10QByteArray10startsWithEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  bool QByteArray::startsWith(const QByteArray & a);
extern void _ZNK10QByteArray10startsWithERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QByteArray::startsWith(char c);
extern void _ZNK10QByteArray10startsWithEc(void* qthis, unsigned char arg0); // 4
  // proto:  int QByteArray::capacity();
extern void _ZNK10QByteArray8capacityEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::toLower();
extern void _ZNKR10QByteArray7toLowerEv(void* qthis); // 2
  // proto:  bool QByteArray::isNull();
extern void _ZNK10QByteArray6isNullEv(void* qthis); // 4
  // proto:  std::string QByteArray::toStdString();
extern void _ZNK10QByteArray11toStdStringEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::toHex();
extern void _ZNK10QByteArray5toHexEv(void* qthis); // 4
  // proto:  int QByteArray::indexOf(const char * c, int from);
extern void _ZNK10QByteArray7indexOfEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::indexOf(char c, int from);
extern void _ZNK10QByteArray7indexOfEci(void* qthis, unsigned char arg0, int32_t arg1); // 4
  // proto:  int QByteArray::indexOf(const QByteArray & a, int from);
extern void _ZNK10QByteArray7indexOfERKS_i(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::indexOf(const QString & s, int from);
extern void _ZNK10QByteArray7indexOfERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  const_iterator QByteArray::cbegin();
extern void _ZNK10QByteArray6cbeginEv(void* qthis); // 2
  // proto:  void QByteArray::QByteArray();
extern void _ZN10QByteArrayC2Ev(void* qthis); // 1
  // proto:  void QByteArray::QByteArray(const QByteArray & );
extern void _ZN10QByteArrayC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  void QByteArray::QByteArray(const char * , int size);
extern void _ZN10QByteArrayC2EPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 3
  // proto:  void QByteArray::QByteArray(int size, char c);
extern void _ZN10QByteArrayC2Eic(void* qthis, int32_t arg0, unsigned char arg1); // 3
  // proto: static QByteArray QByteArray::fromBase64(const QByteArray & base64);
extern void _ZN10QByteArray10fromBase64ERKS_(void* arg0); // 4
  // proto: static QByteArray QByteArray::fromHex(const QByteArray & hexEncoded);
extern void _ZN10QByteArray7fromHexERKS_(void* arg0); // 4
  // proto:  void QByteArray::chop(int n);
extern void _ZN10QByteArray4chopEi(void* qthis, int32_t arg0); // 4
  // proto:  int QByteArray::length();
extern void _ZNK10QByteArray6lengthEv(void* qthis); // 2
  // proto:  double QByteArray::toDouble(bool * ok);
extern void _ZNK10QByteArray8toDoubleEPb(void* qthis, bool* arg0); // 4
  // proto:  long QByteArray::toLong(bool * ok, int base);
extern void _ZNK10QByteArray6toLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArray::toPercentEncoding(const QByteArray & exclude, const QByteArray & include, char percent);
extern void _ZNK10QByteArray17toPercentEncodingERKS_S1_c(void* qthis, void* arg0, void* arg1, unsigned char arg2); // 4
  // proto:  ushort QByteArray::toUShort(bool * ok, int base);
extern void _ZNK10QByteArray8toUShortEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  const_iterator QByteArray::constBegin();
extern void _ZNK10QByteArray10constBeginEv(void* qthis); // 2
  // proto:  short QByteArray::toShort(bool * ok, int base);
extern void _ZNK10QByteArray7toShortEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(qlonglong , int base);
extern void _ZN10QByteArray6numberExi(int64_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(qulonglong , int base);
extern void _ZN10QByteArray6numberEyi(int64_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(int , int base);
extern void _ZN10QByteArray6numberEii(int32_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(uint , int base);
extern void _ZN10QByteArray6numberEji(int32_t arg0, int32_t arg1); // 4
  // proto: static QByteArray QByteArray::number(double , char f, int prec);
extern void _ZN10QByteArray6numberEdci(double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QByteArray & QByteArray::replace(const QString & before, const char * after);
extern void _ZN10QByteArray7replaceERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(char before, const QByteArray & after);
extern void _ZN10QByteArray7replaceEcRKS_(void* qthis, unsigned char arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::replace(int index, int len, const QByteArray & s);
extern void _ZN10QByteArray7replaceEiiRKS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QByteArray & QByteArray::replace(char before, const char * after);
extern void _ZN10QByteArray7replaceEcPKc(void* qthis, unsigned char arg0, unsigned char* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(char before, char after);
extern void _ZN10QByteArray7replaceEcc(void* qthis, unsigned char arg0, unsigned char arg1); // 4
  // proto:  QByteArray & QByteArray::replace(const QByteArray & before, const QByteArray & after);
extern void _ZN10QByteArray7replaceERKS_S1_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::replace(const QByteArray & before, const char * after);
extern void _ZN10QByteArray7replaceERKS_PKc(void* qthis, void* arg0, unsigned char* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(int index, int len, const char * s);
extern void _ZN10QByteArray7replaceEiiPKc(void* qthis, int32_t arg0, int32_t arg1, unsigned char* arg2); // 4
  // proto:  QByteArray & QByteArray::replace(int index, int len, const char * s, int alen);
extern void _ZN10QByteArray7replaceEiiPKci(void* qthis, int32_t arg0, int32_t arg1, unsigned char* arg2, int32_t arg3); // 4
  // proto:  QByteArray & QByteArray::replace(const char * before, const QByteArray & after);
extern void _ZN10QByteArray7replaceEPKcRKS_(void* qthis, unsigned char* arg0, void* arg1); // 4
  // proto:  QByteArray & QByteArray::replace(const char * before, const char * after);
extern void _ZN10QByteArray7replaceEPKcS1_(void* qthis, unsigned char* arg0, unsigned char* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(char c, const QString & after);
extern void _ZN10QByteArray7replaceEcRK7QString(void* qthis, unsigned char arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(const QString & before, const QByteArray & after);
extern void _ZN10QByteArray7replaceERK7QStringRKS_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QByteArray & QByteArray::replace(const char * before, int bsize, const char * after, int asize);
extern void _ZN10QByteArray7replaceEPKciS1_i(void* qthis, unsigned char* arg0, int32_t arg1, unsigned char* arg2, int32_t arg3); // 4
  // proto:  void QByteArray::push_back(const QByteArray & a);
extern void _ZN10QByteArray9push_backERKS_(void* qthis, void* arg0); // 2
  // proto:  void QByteArray::push_back(const char * c);
extern void _ZN10QByteArray9push_backEPKc(void* qthis, unsigned char* arg0); // 2
  // proto:  void QByteArray::push_back(char c);
extern void _ZN10QByteArray9push_backEc(void* qthis, unsigned char arg0); // 2
  // proto:  QByteArray QByteArray::simplified();
extern void _ZNKR10QByteArray10simplifiedEv(void* qthis); // 2
  // proto:  int QByteArray::size();
extern void _ZNK10QByteArray4sizeEv(void* qthis); // 2
  // proto:  bool QByteArray::contains(char c);
extern void _ZNK10QByteArray8containsEc(void* qthis, unsigned char arg0); // 2
  // proto:  bool QByteArray::contains(const QByteArray & a);
extern void _ZNK10QByteArray8containsERKS_(void* qthis, void* arg0); // 2
  // proto:  bool QByteArray::contains(const char * a);
extern void _ZNK10QByteArray8containsEPKc(void* qthis, unsigned char* arg0); // 2
  // proto: static QByteArray QByteArray::fromPercentEncoding(const QByteArray & pctEncoded, char percent);
extern void _ZN10QByteArray19fromPercentEncodingERKS_c(void* arg0, unsigned char arg1); // 4
  // proto:  const_iterator QByteArray::cend();
extern void _ZNK10QByteArray4cendEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::leftJustified(int width, char fill, bool truncate);
extern void _ZNK10QByteArray13leftJustifiedEicb(void* qthis, int32_t arg0, unsigned char arg1, bool arg2); // 4
  // proto:  QByteArray QByteArray::repeated(int times);
extern void _ZNK10QByteArray8repeatedEi(void* qthis, int32_t arg0); // 4
  // proto: static QByteArray QByteArray::fromRawData(const char * , int size);
extern void _ZN10QByteArray11fromRawDataEPKci(unsigned char* arg0, int32_t arg1); // 4
  // proto:  void QByteArray::squeeze();
extern void _ZN10QByteArray7squeezeEv(void* qthis); // 2
  // proto:  void QByteArray::resize(int size);
extern void _ZN10QByteArray6resizeEi(void* qthis, int32_t arg0); // 4
  // proto:  qulonglong QByteArray::toULongLong(bool * ok, int base);
extern void _ZNK10QByteArray11toULongLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::count();
extern void _ZNK10QByteArray5countEv(void* qthis); // 2
  // proto:  int QByteArray::count(const char * a);
extern void _ZNK10QByteArray5countEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  int QByteArray::count(char c);
extern void _ZNK10QByteArray5countEc(void* qthis, unsigned char arg0); // 4
  // proto:  int QByteArray::count(const QByteArray & a);
extern void _ZNK10QByteArray5countERKS_(void* qthis, void* arg0); // 4
  // proto:  void QByteArray::~QByteArray();
extern void _ZN10QByteArrayD2Ev(void* qthis); // 2
  // proto:  float QByteArray::toFloat(bool * ok);
extern void _ZNK10QByteArray7toFloatEPb(void* qthis, bool* arg0); // 4
  // proto:  char QByteArray::at(int i);
extern void _ZNK10QByteArray2atEi(void* qthis, int32_t arg0); // 2
  // proto:  QByteArray & QByteArray::fill(char c, int size);
extern void _ZN10QByteArray4fillEci(void* qthis, unsigned char arg0, int32_t arg1); // 4
  // proto:  ulong QByteArray::toULong(bool * ok, int base);
extern void _ZNK10QByteArray7toULongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  iterator QByteArray::end();
extern void _ZN10QByteArray3endEv(void* qthis); // 2
  // proto:  QByteArray QByteArray::mid(int index, int len);
extern void _ZNK10QByteArray3midEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::prepend(const char * s);
extern void _ZN10QByteArray7prependEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  QByteArray & QByteArray::prepend(char c);
extern void _ZN10QByteArray7prependEc(void* qthis, unsigned char arg0); // 4
  // proto:  QByteArray & QByteArray::prepend(const char * s, int len);
extern void _ZN10QByteArray7prependEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::prepend(const QByteArray & a);
extern void _ZN10QByteArray7prependERKS_(void* qthis, void* arg0); // 4
  // proto:  const char * QByteArray::constData();
extern void _ZNK10QByteArray9constDataEv(void* qthis); // 2
  // proto:  qlonglong QByteArray::toLongLong(bool * ok, int base);
extern void _ZNK10QByteArray10toLongLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::isEmpty();
extern void _ZNK10QByteArray7isEmptyEv(void* qthis); // 2
  // proto:  QList<QByteArray> QByteArray::split(char sep);
extern void _ZNK10QByteArray5splitEc(void* qthis, unsigned char arg0); // 4
  // proto:  QByteArray & QByteArray::setRawData(const char * a, uint n);
extern void _ZN10QByteArray10setRawDataEPKcj(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::isDetached();
extern void _ZNK10QByteArray10isDetachedEv(void* qthis); // 2
  // proto:  iterator QByteArray::begin();
extern void _ZN10QByteArray5beginEv(void* qthis); // 2
  // proto:  int QByteArray::lastIndexOf(const char * c, int from);
extern void _ZNK10QByteArray11lastIndexOfEPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::lastIndexOf(const QString & s, int from);
extern void _ZNK10QByteArray11lastIndexOfERK7QStringi(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  int QByteArray::lastIndexOf(const QByteArray & a, int from);
extern void _ZNK10QByteArray11lastIndexOfERKS_i(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QByteArray::lastIndexOf(char c, int from);
extern void _ZNK10QByteArray11lastIndexOfEci(void* qthis, unsigned char arg0, int32_t arg1); // 4
  // proto:  QByteArray QByteArray::toUpper();
extern void _ZNKR10QByteArray7toUpperEv(void* qthis); // 2
  // proto:  void QByteArray::truncate(int pos);
extern void _ZN10QByteArray8truncateEi(void* qthis, int32_t arg0); // 4
  // proto:  QByteArray QByteArray::toBase64();
extern void _ZNK10QByteArray8toBase64Ev(void* qthis); // 4
  // proto:  char * QByteArray::data();
extern void _ZN10QByteArray4dataEv(void* qthis); // 2
  // proto:  void QByteArray::clear();
extern void _ZN10QByteArray5clearEv(void* qthis); // 4
  // proto:  int QByteArray::toInt(bool * ok, int base);
extern void _ZNK10QByteArray5toIntEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  bool QByteArray::endsWith(const QByteArray & a);
extern void _ZNK10QByteArray8endsWithERKS_(void* qthis, void* arg0); // 4
  // proto:  bool QByteArray::endsWith(const char * c);
extern void _ZNK10QByteArray8endsWithEPKc(void* qthis, unsigned char* arg0); // 4
  // proto:  bool QByteArray::endsWith(char c);
extern void _ZNK10QByteArray8endsWithEc(void* qthis, unsigned char arg0); // 4
  // proto:  QByteArray & QByteArray::setNum(uint , int base);
extern void _ZN10QByteArray6setNumEji(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QByteArray & QByteArray::setNum(int , int base);
extern void _ZN10QByteArray6setNumEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QByteArray & QByteArray::setNum(short , int base);
extern void _ZN10QByteArray6setNumEsi(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  QByteArray & QByteArray::setNum(double , char f, int prec);
extern void _ZN10QByteArray6setNumEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QByteArray & QByteArray::setNum(float , char f, int prec);
extern void _ZN10QByteArray6setNumEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2); // 2
  // proto:  QByteArray & QByteArray::setNum(qulonglong , int base);
extern void _ZN10QByteArray6setNumEyi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::setNum(qlonglong , int base);
extern void _ZN10QByteArray6setNumExi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QByteArray & QByteArray::setNum(ushort , int base);
extern void _ZN10QByteArray6setNumEti(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  void QByteArray::reserve(int size);
extern void _ZN10QByteArray7reserveEi(void* qthis, int32_t arg0); // 2
  // proto:  void QByteArray::push_front(char c);
extern void _ZN10QByteArray10push_frontEc(void* qthis, unsigned char arg0); // 2
  // proto:  void QByteArray::push_front(const QByteArray & a);
extern void _ZN10QByteArray10push_frontERKS_(void* qthis, void* arg0); // 2
  // proto:  void QByteArray::push_front(const char * c);
extern void _ZN10QByteArray10push_frontEPKc(void* qthis, unsigned char* arg0); // 2
  // proto:  QByteArray QByteArray::left(int len);
extern void _ZNK10QByteArray4leftEi(void* qthis, int32_t arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QByteArray)=8
type QByteArray struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QByteArrayDataPtr)=8
type QByteArrayDataPtr struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// detach()
func (this *QByteArray) detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6detachEv
    // invoke: void detach()
    C._ZN10QByteArray6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "detach", args)
  }

}

// swap(class QByteArray &)
func (this *QByteArray) swap(args ...interface{}) () {
  // swap(class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4swapERS_
    // invoke: void swap(class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "swap", args)
  }

}

// rightJustified(int, char, _Bool)
func (this *QByteArray) rightJustified(args ...interface{}) () {
  // rightJustified(int, char, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray14rightJustifiedEicb
    // invoke: QByteArray rightJustified(int, char, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C._ZNK10QByteArray14rightJustifiedEicb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QByteArray", "rightJustified", args)
  }

}

// insert(int, const class QString &)
func (this *QByteArray) insert(args ...interface{}) () {
  // insert(int, const class QString &)
  // insert(int, const char *)
  // insert(int, const class QByteArray &)
  // insert(int, char)
  // insert(int, const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.ByteTy(false) // "char"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6insertEiRK7QString
    // invoke: QByteArray & insert(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEiRK7QString(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QByteArray6insertEiPKc
    // invoke: QByteArray & insert(int, const char *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEiPKc(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray6insertEiRKS_
    // invoke: QByteArray & insert(int, const class QByteArray &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEiRKS_(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArray6insertEic
    // invoke: QByteArray & insert(int, char)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEic(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN10QByteArray6insertEiPKci
    // invoke: QByteArray & insert(int, const char *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray6insertEiPKci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QByteArray", "insert", args)
  }

}

// remove(int, int)
func (this *QByteArray) remove(args ...interface{}) () {
  // remove(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6removeEii
    // invoke: QByteArray & remove(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6removeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "remove", args)
  }

}

// trimmed()
func (this *QByteArray) trimmed(args ...interface{}) () {
  // trimmed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray7trimmedEv
    // invoke: QByteArray trimmed()
    C._ZNKR10QByteArray7trimmedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "trimmed", args)
  }

}

// constEnd()
func (this *QByteArray) constEnd(args ...interface{}) () {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8constEndEv
    // invoke: const_iterator constEnd()
    C._ZNK10QByteArray8constEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "constEnd", args)
  }

}

// isSharedWith(const class QByteArray &)
func (this *QByteArray) isSharedWith(args ...interface{}) () {
  // isSharedWith(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray12isSharedWithERKS_
    // invoke: bool isSharedWith(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray12isSharedWithERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "isSharedWith", args)
  }

}

// right(int)
func (this *QByteArray) right(args ...interface{}) () {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5rightEi
    // invoke: QByteArray right(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5rightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "right", args)
  }

}

// toUInt(_Bool *, int)
func (this *QByteArray) toUInt(args ...interface{}) () {
  // toUInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6toUIntEPbi
    // invoke: uint toUInt(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray6toUIntEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toUInt", args)
  }

}

// append(char)
func (this *QByteArray) append(args ...interface{}) () {
  // append(char)
  // append(const char *, int)
  // append(const class QByteArray &)
  // append(const class QString &)
  // append(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6appendEc
    // invoke: QByteArray & append(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6appendEc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray6appendEPKci
    // invoke: QByteArray & append(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6appendEPKci(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray6appendERKS_
    // invoke: QByteArray & append(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6appendERKS_(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN10QByteArray6appendERK7QString
    // invoke: QByteArray & append(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6appendERK7QString(this.qclsinst, arg0)
  case 4:
    // invoke: _ZN10QByteArray6appendEPKc
    // invoke: QByteArray & append(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6appendEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "append", args)
  }

}

// startsWith(const char *)
func (this *QByteArray) startsWith(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10startsWithEPKc
    // invoke: bool startsWith(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray10startsWithEPKc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QByteArray10startsWithERKS_
    // invoke: bool startsWith(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray10startsWithERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QByteArray10startsWithEc
    // invoke: bool startsWith(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray10startsWithEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "startsWith", args)
  }

}

// capacity()
func (this *QByteArray) capacity(args ...interface{}) () {
  // capacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8capacityEv
    // invoke: int capacity()
    C._ZNK10QByteArray8capacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "capacity", args)
  }

}

// toLower()
func (this *QByteArray) toLower(args ...interface{}) () {
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray7toLowerEv
    // invoke: QByteArray toLower()
    C._ZNKR10QByteArray7toLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "toLower", args)
  }

}

// isNull()
func (this *QByteArray) isNull(args ...interface{}) () {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6isNullEv
    // invoke: bool isNull()
    C._ZNK10QByteArray6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "isNull", args)
  }

}

// toStdString()
func (this *QByteArray) toStdString(args ...interface{}) () {
  // toStdString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11toStdStringEv
    // invoke: std::string toStdString()
    C._ZNK10QByteArray11toStdStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "toStdString", args)
  }

}

// toHex()
func (this *QByteArray) toHex(args ...interface{}) () {
  // toHex()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5toHexEv
    // invoke: QByteArray toHex()
    C._ZNK10QByteArray5toHexEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "toHex", args)
  }

}

// indexOf(const char *, int)
func (this *QByteArray) indexOf(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7indexOfEPKci
    // invoke: int indexOf(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfEPKci(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK10QByteArray7indexOfEci
    // invoke: int indexOf(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfEci(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK10QByteArray7indexOfERKS_i
    // invoke: int indexOf(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfERKS_i(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK10QByteArray7indexOfERK7QStringi
    // invoke: int indexOf(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "indexOf", args)
  }

}

// cbegin()
func (this *QByteArray) cbegin(args ...interface{}) () {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6cbeginEv
    // invoke: const_iterator cbegin()
    C._ZNK10QByteArray6cbeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "cbegin", args)
  }

}

// QByteArray()
func NewQByteArray(args ...interface{}) QByteArray {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArrayC1Ev
    // invoke: void QByteArray()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QByteArrayC2Ev(qthis)
  case 1:
    // invoke: _ZN10QByteArrayC1ERKS_
    // invoke: void QByteArray(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QByteArrayC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN10QByteArrayC1EPKci
    // invoke: void QByteArray(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QByteArrayC2EPKci(qthis, arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArrayC1Eic
    // invoke: void QByteArray(int, char)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QByteArrayC2Eic(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "QByteArray", args)
  }

  return QByteArray{}
}

// fromBase64(const class QByteArray &)
func (this *QByteArray) fromBase64_s(args ...interface{}) () {
  // fromBase64(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10fromBase64ERKS_
    // invoke: QByteArray fromBase64(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray10fromBase64ERKS_(arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "fromBase64", args)
  }

}

// fromHex(const class QByteArray &)
func (this *QByteArray) fromHex_s(args ...interface{}) () {
  // fromHex(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7fromHexERKS_
    // invoke: QByteArray fromHex(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7fromHexERKS_(arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "fromHex", args)
  }

}

// chop(int)
func (this *QByteArray) chop(args ...interface{}) () {
  // chop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4chopEi
    // invoke: void chop(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray4chopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "chop", args)
  }

}

// length()
func (this *QByteArray) length(args ...interface{}) () {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6lengthEv
    // invoke: int length()
    C._ZNK10QByteArray6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "length", args)
  }

}

// toDouble(_Bool *)
func (this *QByteArray) toDouble(args ...interface{}) () {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toDoubleEPb
    // invoke: double toDouble(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8toDoubleEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "toDouble", args)
  }

}

// toLong(_Bool *, int)
func (this *QByteArray) toLong(args ...interface{}) () {
  // toLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray6toLongEPbi
    // invoke: long toLong(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray6toLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toLong", args)
  }

}

// toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
func (this *QByteArray) toPercentEncoding(args ...interface{}) () {
  // toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][2] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray17toPercentEncodingERKS_S1_c
    // invoke: QByteArray toPercentEncoding(const class QByteArray &, const class QByteArray &, char)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.uchar(args[2].(byte))
    if false {fmt.Println(arg2)}
    C._ZNK10QByteArray17toPercentEncodingERKS_S1_c(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QByteArray", "toPercentEncoding", args)
  }

}

// toUShort(_Bool *, int)
func (this *QByteArray) toUShort(args ...interface{}) () {
  // toUShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toUShortEPbi
    // invoke: ushort toUShort(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray8toUShortEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toUShort", args)
  }

}

// constBegin()
func (this *QByteArray) constBegin(args ...interface{}) () {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10constBeginEv
    // invoke: const_iterator constBegin()
    C._ZNK10QByteArray10constBeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "constBegin", args)
  }

}

// toShort(_Bool *, int)
func (this *QByteArray) toShort(args ...interface{}) () {
  // toShort(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toShortEPbi
    // invoke: short toShort(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7toShortEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toShort", args)
  }

}

// number(qlonglong, int)
func (this *QByteArray) number_s(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6numberExi
    // invoke: QByteArray number(qlonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6numberExi(arg0, arg1)
  case 1:
    // invoke: _ZN10QByteArray6numberEyi
    // invoke: QByteArray number(qulonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6numberEyi(arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray6numberEii
    // invoke: QByteArray number(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6numberEii(arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArray6numberEji
    // invoke: QByteArray number(uint, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6numberEji(arg0, arg1)
  case 4:
    // invoke: _ZN10QByteArray6numberEdci
    // invoke: QByteArray number(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray6numberEdci(arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QByteArray", "number", args)
  }

}

// replace(const class QString &, const char *)
func (this *QByteArray) replace(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7replaceERK7QStringPKc
    // invoke: QByteArray & replace(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceERK7QStringPKc(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QByteArray7replaceEcRKS_
    // invoke: QByteArray & replace(char, const class QByteArray &)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEcRKS_(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray7replaceEiiRKS_
    // invoke: QByteArray & replace(int, int, const class QByteArray &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QByteArray).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray7replaceEiiRKS_(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN10QByteArray7replaceEcPKc
    // invoke: QByteArray & replace(char, const char *)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEcPKc(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN10QByteArray7replaceEcc
    // invoke: QByteArray & replace(char, char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEcc(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN10QByteArray7replaceERKS_S1_
    // invoke: QByteArray & replace(const class QByteArray &, const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceERKS_S1_(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN10QByteArray7replaceERKS_PKc
    // invoke: QByteArray & replace(const class QByteArray &, const char *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceERKS_PKc(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN10QByteArray7replaceEiiPKc
    // invoke: QByteArray & replace(int, int, const char *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray7replaceEiiPKc(this.qclsinst, arg0, arg1, arg2)
  case 8:
    // invoke: _ZN10QByteArray7replaceEiiPKci
    // invoke: QByteArray & replace(int, int, const char *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QByteArray7replaceEiiPKci(this.qclsinst, arg0, arg1, arg2, arg3)
  case 9:
    // invoke: _ZN10QByteArray7replaceEPKcRKS_
    // invoke: QByteArray & replace(const char *, const class QByteArray &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEPKcRKS_(this.qclsinst, arg0, arg1)
  case 10:
    // invoke: _ZN10QByteArray7replaceEPKcS1_
    // invoke: QByteArray & replace(const char *, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).Pointer()))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEPKcS1_(this.qclsinst, arg0, arg1)
  case 11:
    // invoke: _ZN10QByteArray7replaceEcRK7QString
    // invoke: QByteArray & replace(char, const class QString &)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEcRK7QString(this.qclsinst, arg0, arg1)
  case 12:
    // invoke: _ZN10QByteArray7replaceERK7QStringRKS_
    // invoke: QByteArray & replace(const class QString &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceERK7QStringRKS_(this.qclsinst, arg0, arg1)
  case 13:
    // invoke: _ZN10QByteArray7replaceEPKciS1_i
    // invoke: QByteArray & replace(const char *, int, const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).Pointer()))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QByteArray7replaceEPKciS1_i(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QByteArray", "replace", args)
  }

}

// push_back(const class QByteArray &)
func (this *QByteArray) push_back(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray9push_backERKS_
    // invoke: void push_back(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray9push_backERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray9push_backEPKc
    // invoke: void push_back(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray9push_backEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray9push_backEc
    // invoke: void push_back(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray9push_backEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "push_back", args)
  }

}

// simplified()
func (this *QByteArray) simplified(args ...interface{}) () {
  // simplified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray10simplifiedEv
    // invoke: QByteArray simplified()
    C._ZNKR10QByteArray10simplifiedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "simplified", args)
  }

}

// size()
func (this *QByteArray) size(args ...interface{}) () {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4sizeEv
    // invoke: int size()
    C._ZNK10QByteArray4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "size", args)
  }

}

// contains(char)
func (this *QByteArray) contains(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8containsEc
    // invoke: bool contains(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8containsEc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QByteArray8containsERKS_
    // invoke: bool contains(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8containsERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QByteArray8containsEPKc
    // invoke: bool contains(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8containsEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "contains", args)
  }

}

// fromPercentEncoding(const class QByteArray &, char)
func (this *QByteArray) fromPercentEncoding_s(args ...interface{}) () {
  // fromPercentEncoding(const class QByteArray &, char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray19fromPercentEncodingERKS_c
    // invoke: QByteArray fromPercentEncoding(const class QByteArray &, char)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray19fromPercentEncodingERKS_c(arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "fromPercentEncoding", args)
  }

}

// cend()
func (this *QByteArray) cend(args ...interface{}) () {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4cendEv
    // invoke: const_iterator cend()
    C._ZNK10QByteArray4cendEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "cend", args)
  }

}

// leftJustified(int, char, _Bool)
func (this *QByteArray) leftJustified(args ...interface{}) () {
  // leftJustified(int, char, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray13leftJustifiedEicb
    // invoke: QByteArray leftJustified(int, char, _Bool)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C._ZNK10QByteArray13leftJustifiedEicb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QByteArray", "leftJustified", args)
  }

}

// repeated(int)
func (this *QByteArray) repeated(args ...interface{}) () {
  // repeated(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8repeatedEi
    // invoke: QByteArray repeated(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8repeatedEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "repeated", args)
  }

}

// fromRawData(const char *, int)
func (this *QByteArray) fromRawData_s(args ...interface{}) () {
  // fromRawData(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray11fromRawDataEPKci
    // invoke: QByteArray fromRawData(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray11fromRawDataEPKci(arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "fromRawData", args)
  }

}

// squeeze()
func (this *QByteArray) squeeze(args ...interface{}) () {
  // squeeze()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7squeezeEv
    // invoke: void squeeze()
    C._ZN10QByteArray7squeezeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "squeeze", args)
  }

}

// resize(int)
func (this *QByteArray) resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6resizeEi
    // invoke: void resize(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6resizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "resize", args)
  }

}

// toULongLong(_Bool *, int)
func (this *QByteArray) toULongLong(args ...interface{}) () {
  // toULongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11toULongLongEPbi
    // invoke: qulonglong toULongLong(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11toULongLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toULongLong", args)
  }

}

// count()
func (this *QByteArray) count(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5countEv
    // invoke: int count()
    C._ZNK10QByteArray5countEv(this.qclsinst)
  case 1:
    // invoke: _ZNK10QByteArray5countEPKc
    // invoke: int count(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5countEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QByteArray5countEc
    // invoke: int count(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5countEc(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK10QByteArray5countERKS_
    // invoke: int count(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5countERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "count", args)
  }

}

// ~QByteArray()
func (this *QByteArray) FreeQByteArray(args ...interface{}) () {
  // ~QByteArray()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArrayD0Ev
    // invoke: void ~QByteArray()
    C._ZN10QByteArrayD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "~QByteArray", args)
  }

}

// toFloat(_Bool *)
func (this *QByteArray) toFloat(args ...interface{}) () {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toFloatEPb
    // invoke: float toFloat(_Bool *)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray7toFloatEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "toFloat", args)
  }

}

// at(int)
func (this *QByteArray) at(args ...interface{}) () {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray2atEi
    // invoke: char at(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "at", args)
  }

}

// fill(char, int)
func (this *QByteArray) fill(args ...interface{}) () {
  // fill(char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4fillEci
    // invoke: QByteArray & fill(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray4fillEci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "fill", args)
  }

}

// toULong(_Bool *, int)
func (this *QByteArray) toULong(args ...interface{}) () {
  // toULong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7toULongEPbi
    // invoke: ulong toULong(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7toULongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toULong", args)
  }

}

// end()
func (this *QByteArray) end(args ...interface{}) () {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray3endEv
    // invoke: iterator end()
    C._ZN10QByteArray3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "end", args)
  }

}

// mid(int, int)
func (this *QByteArray) mid(args ...interface{}) () {
  // mid(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray3midEii
    // invoke: QByteArray mid(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray3midEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "mid", args)
  }

}

// prepend(const char *)
func (this *QByteArray) prepend(args ...interface{}) () {
  // prepend(const char *)
  // prepend(char)
  // prepend(const char *, int)
  // prepend(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7prependEPKc
    // invoke: QByteArray & prepend(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7prependEPKc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray7prependEc
    // invoke: QByteArray & prepend(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7prependEc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray7prependEPKci
    // invoke: QByteArray & prepend(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7prependEPKci(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArray7prependERKS_
    // invoke: QByteArray & prepend(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7prependERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "prepend", args)
  }

}

// constData()
func (this *QByteArray) constData(args ...interface{}) () {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray9constDataEv
    // invoke: const char * constData()
    C._ZNK10QByteArray9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "constData", args)
  }

}

// toLongLong(_Bool *, int)
func (this *QByteArray) toLongLong(args ...interface{}) () {
  // toLongLong(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10toLongLongEPbi
    // invoke: qlonglong toLongLong(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray10toLongLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toLongLong", args)
  }

}

// isEmpty()
func (this *QByteArray) isEmpty(args ...interface{}) () {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray7isEmptyEv
    // invoke: bool isEmpty()
    C._ZNK10QByteArray7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "isEmpty", args)
  }

}

// split(char)
func (this *QByteArray) split(args ...interface{}) () {
  // split(char)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5splitEc
    // invoke: QList<QByteArray> split(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5splitEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "split", args)
  }

}

// setRawData(const char *, uint)
func (this *QByteArray) setRawData(args ...interface{}) () {
  // setRawData(const char *, uint)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "uint"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10setRawDataEPKcj
    // invoke: QByteArray & setRawData(const char *, uint)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray10setRawDataEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "setRawData", args)
  }

}

// isDetached()
func (this *QByteArray) isDetached(args ...interface{}) () {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray10isDetachedEv
    // invoke: bool isDetached()
    C._ZNK10QByteArray10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "isDetached", args)
  }

}

// begin()
func (this *QByteArray) begin(args ...interface{}) () {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray5beginEv
    // invoke: iterator begin()
    C._ZN10QByteArray5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "begin", args)
  }

}

// lastIndexOf(const char *, int)
func (this *QByteArray) lastIndexOf(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11lastIndexOfEPKci
    // invoke: int lastIndexOf(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfEPKci(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK10QByteArray11lastIndexOfERK7QStringi
    // invoke: int lastIndexOf(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfERK7QStringi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK10QByteArray11lastIndexOfERKS_i
    // invoke: int lastIndexOf(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfERKS_i(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK10QByteArray11lastIndexOfEci
    // invoke: int lastIndexOf(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfEci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "lastIndexOf", args)
  }

}

// toUpper()
func (this *QByteArray) toUpper(args ...interface{}) () {
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR10QByteArray7toUpperEv
    // invoke: QByteArray toUpper()
    C._ZNKR10QByteArray7toUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "toUpper", args)
  }

}

// truncate(int)
func (this *QByteArray) truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray8truncateEi
    // invoke: void truncate(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray8truncateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "truncate", args)
  }

}

// toBase64()
func (this *QByteArray) toBase64(args ...interface{}) () {
  // toBase64()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8toBase64Ev
    // invoke: QByteArray toBase64()
    C._ZNK10QByteArray8toBase64Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "toBase64", args)
  }

}

// data()
func (this *QByteArray) data(args ...interface{}) () {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4dataEv
    // invoke: char * data()
    C._ZN10QByteArray4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "data", args)
  }

}

// clear()
func (this *QByteArray) clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray5clearEv
    // invoke: void clear()
    C._ZN10QByteArray5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "clear", args)
  }

}

// toInt(_Bool *, int)
func (this *QByteArray) toInt(args ...interface{}) () {
  // toInt(_Bool *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5toIntEPbi
    // invoke: int toInt(_Bool *, int)
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray5toIntEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "toInt", args)
  }

}

// endsWith(const class QByteArray &)
func (this *QByteArray) endsWith(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8endsWithERKS_
    // invoke: bool endsWith(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8endsWithERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QByteArray8endsWithEPKc
    // invoke: bool endsWith(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8endsWithEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QByteArray8endsWithEc
    // invoke: bool endsWith(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8endsWithEc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "endsWith", args)
  }

}

// setNum(uint, int)
func (this *QByteArray) setNum(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6setNumEji
    // invoke: QByteArray & setNum(uint, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumEji(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QByteArray6setNumEii
    // invoke: QByteArray & setNum(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumEii(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray6setNumEsi
    // invoke: QByteArray & setNum(short, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumEsi(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArray6setNumEdci
    // invoke: QByteArray & setNum(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray6setNumEdci(this.qclsinst, arg0, arg1, arg2)
  case 4:
    // invoke: _ZN10QByteArray6setNumEfci
    // invoke: QByteArray & setNum(float, char, int)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray6setNumEfci(this.qclsinst, arg0, arg1, arg2)
  case 5:
    // invoke: _ZN10QByteArray6setNumEyi
    // invoke: QByteArray & setNum(qulonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumEyi(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN10QByteArray6setNumExi
    // invoke: QByteArray & setNum(qlonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumExi(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN10QByteArray6setNumEti
    // invoke: QByteArray & setNum(ushort, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumEti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "setNum", args)
  }

}

// reserve(int)
func (this *QByteArray) reserve(args ...interface{}) () {
  // reserve(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7reserveEi
    // invoke: void reserve(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7reserveEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "reserve", args)
  }

}

// push_front(char)
func (this *QByteArray) push_front(args ...interface{}) () {
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

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10push_frontEc
    // invoke: void push_front(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray10push_frontEc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray10push_frontERKS_
    // invoke: void push_front(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray10push_frontERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray10push_frontEPKc
    // invoke: void push_front(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray10push_frontEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "push_front", args)
  }

}

// left(int)
func (this *QByteArray) left(args ...interface{}) () {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray4leftEi
    // invoke: QByteArray left(int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray4leftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "left", args)
  }

}

// <= body block end

