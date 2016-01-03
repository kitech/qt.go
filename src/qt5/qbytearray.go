package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  void QByteRef::QByteRef(QByteArray & array, int idx);
extern void* dector_ZN8QByteRefC1ER10QByteArrayi(void* arg0, int32_t arg1);
extern void demth_ZN8QByteRefC1ER10QByteArrayi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::insert(int i, char c);
extern void _ZN10QByteArray6insertEic(void* qthis, int32_t arg0, unsigned char arg1);
  // proto:  int QByteArray::lastIndexOf(const QByteArray & a, int from);
extern void _ZNK10QByteArray11lastIndexOfERKS_i(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QByteArray::push_front(const QByteArray & a);
extern void demth_ZN10QByteArray10push_frontERKS_(void* qthis, void* arg0);
  // proto:  ulong QByteArray::toULong(bool * ok, int base);
extern void _ZNK10QByteArray7toULongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(const char * before, const char * after);
extern void demth_ZN10QByteArray7replaceEPKcS1_(void* qthis, unsigned char* arg0, unsigned char* arg1);
  // proto:  QByteArray & QByteArray::replace(const QByteArray & before, const QByteArray & after);
extern void _ZN10QByteArray7replaceERKS_S1_(void* qthis, void* arg0, void* arg1);
  // proto: static QByteArray QByteArray::fromHex(const QByteArray & hexEncoded);
extern void _ZN10QByteArray7fromHexERKS_(void* arg0);
  // proto:  QByteArray & QByteArray::prepend(const char * s);
extern void _ZN10QByteArray7prependEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QByteArray::count(const QByteArray & a);
extern void _ZNK10QByteArray5countERKS_(void* qthis, void* arg0);
  // proto:  void QByteArray::~QByteArray();
extern void demth_ZN10QByteArrayD0Ev(void* qthis);
  // proto:  void QByteArray::QByteArray();
extern void* dector_ZN10QByteArrayC1Ev();
extern void demth_ZN10QByteArrayC1Ev(void* qthis);
  // proto:  QByteArray & QByteArray::replace(const char * before, const QByteArray & after);
extern void _ZN10QByteArray7replaceEPKcRKS_(void* qthis, unsigned char* arg0, void* arg1);
  // proto:  float QByteArray::toFloat(bool * ok);
extern void _ZNK10QByteArray7toFloatEPb(void* qthis, bool* arg0);
  // proto:  void QByteArray::truncate(int pos);
extern void _ZN10QByteArray8truncateEi(void* qthis, int32_t arg0);
  // proto:  QByteArray QByteArray::toBase64();
extern void _ZNK10QByteArray8toBase64Ev(void* qthis);
  // proto:  bool QByteArray::isEmpty();
extern void demth_ZNK10QByteArray7isEmptyEv(void* qthis);
  // proto:  QByteArray & QByteArray::insert(int i, const char * s, int len);
extern void _ZN10QByteArray6insertEiPKci(void* qthis, int32_t arg0, unsigned char* arg1, int32_t arg2);
  // proto:  QByteArray & QByteArray::insert(int i, const QString & s);
extern void demth_ZN10QByteArray6insertEiRK7QString(void* qthis, int32_t arg0, void* arg1);
  // proto:  void QByteArray::resize(int size);
extern void _ZN10QByteArray6resizeEi(void* qthis, int32_t arg0);
  // proto:  QByteArray & QByteArray::replace(int index, int len, const char * s, int alen);
extern void _ZN10QByteArray7replaceEiiPKci(void* qthis, int32_t arg0, int32_t arg1, unsigned char* arg2, int32_t arg3);
  // proto:  int QByteArray::lastIndexOf(const QString & s, int from);
extern void demth_ZNK10QByteArray11lastIndexOfERK7QStringi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QByteArray QByteArray::toHex();
extern void _ZNK10QByteArray5toHexEv(void* qthis);
  // proto:  int QByteArray::indexOf(const char * c, int from);
extern void _ZNK10QByteArray7indexOfEPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(char before, const QByteArray & after);
extern void _ZN10QByteArray7replaceEcRKS_(void* qthis, unsigned char arg0, void* arg1);
  // proto:  uint QByteArray::toUInt(bool * ok, int base);
extern void _ZNK10QByteArray6toUIntEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto: static QByteArray QByteArray::fromStdString(const std::string & s);
extern void demth_ZN10QByteArray13fromStdStringERKi(int32_t* arg0);
  // proto:  bool QByteArray::isNull();
extern void _ZNK10QByteArray6isNullEv(void* qthis);
  // proto:  void QByteArray::reserve(int size);
extern void demth_ZN10QByteArray7reserveEi(void* qthis, int32_t arg0);
  // proto: static QByteArray QByteArray::fromRawData(const char * , int size);
extern void _ZN10QByteArray11fromRawDataEPKci(unsigned char* arg0, int32_t arg1);
  // proto:  bool QByteArray::contains(char c);
extern void demth_ZNK10QByteArray8containsEc(void* qthis, unsigned char arg0);
  // proto:  void QByteArray::QByteArray(int size, char c);
extern void* dector_ZN10QByteArrayC1Eic(int32_t arg0, unsigned char arg1);
extern void _ZN10QByteArrayC1Eic(void* qthis, int32_t arg0, unsigned char arg1);
  // proto:  int QByteArray::indexOf(const QByteArray & a, int from);
extern void _ZNK10QByteArray7indexOfERKS_i(void* qthis, void* arg0, int32_t arg1);
  // proto:  long QByteArray::toLong(bool * ok, int base);
extern void _ZNK10QByteArray6toLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  int QByteArray::indexOf(char c, int from);
extern void _ZNK10QByteArray7indexOfEci(void* qthis, unsigned char arg0, int32_t arg1);
  // proto:  char * QByteArray::data();
extern void demth_ZN10QByteArray4dataEv(void* qthis);
  // proto: static QByteArray QByteArray::number(double , char f, int prec);
extern void _ZN10QByteArray6numberEdci(double arg0, unsigned char arg1, int32_t arg2);
  // proto:  int QByteArray::capacity();
extern void demth_ZNK10QByteArray8capacityEv(void* qthis);
  // proto:  int QByteArray::count();
extern void demth_ZNK10QByteArray5countEv(void* qthis);
  // proto: static QByteArray QByteArray::fromBase64(const QByteArray & base64);
extern void _ZN10QByteArray10fromBase64ERKS_(void* arg0);
  // proto:  QByteArray QByteArray::left(int len);
extern void _ZNK10QByteArray4leftEi(void* qthis, int32_t arg0);
  // proto:  QByteArray & QByteArray::replace(char before, char after);
extern void _ZN10QByteArray7replaceEcc(void* qthis, unsigned char arg0, unsigned char arg1);
  // proto:  QByteArray & QByteArray::append(char c);
extern void _ZN10QByteArray6appendEc(void* qthis, unsigned char arg0);
  // proto:  bool QByteArray::startsWith(const char * c);
extern void _ZNK10QByteArray10startsWithEPKc(void* qthis, unsigned char* arg0);
  // proto:  QByteArray & QByteArray::remove(int index, int len);
extern void _ZN10QByteArray6removeEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  int QByteArray::lastIndexOf(char c, int from);
extern void _ZNK10QByteArray11lastIndexOfEci(void* qthis, unsigned char arg0, int32_t arg1);
  // proto:  bool QByteArray::startsWith(const QByteArray & a);
extern void _ZNK10QByteArray10startsWithERKS_(void* qthis, void* arg0);
  // proto:  bool QByteArray::contains(const char * a);
extern void demth_ZNK10QByteArray8containsEPKc(void* qthis, unsigned char* arg0);
  // proto:  bool QByteArray::endsWith(const char * c);
extern void _ZNK10QByteArray8endsWithEPKc(void* qthis, unsigned char* arg0);
  // proto:  void QByteArray::squeeze();
extern void demth_ZN10QByteArray7squeezeEv(void* qthis);
  // proto:  void QByteArray::QByteArray(const char * , int size);
extern void* dector_ZN10QByteArrayC1EPKci(unsigned char* arg0, int32_t arg1);
extern void _ZN10QByteArrayC1EPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  int QByteArray::indexOf(const QString & s, int from);
extern void demth_ZNK10QByteArray7indexOfERK7QStringi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QByteArray::detach();
extern void demth_ZN10QByteArray6detachEv(void* qthis);
  // proto:  QByteArray QByteArray::repeated(int times);
extern void _ZNK10QByteArray8repeatedEi(void* qthis, int32_t arg0);
  // proto:  QByteArray & QByteArray::setNum(float , char f, int prec);
extern void demth_ZN10QByteArray6setNumEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2);
  // proto:  short QByteArray::toShort(bool * ok, int base);
extern void _ZNK10QByteArray7toShortEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::prepend(char c);
extern void _ZN10QByteArray7prependEc(void* qthis, unsigned char arg0);
  // proto:  int QByteArray::toInt(bool * ok, int base);
extern void _ZNK10QByteArray5toIntEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  void QByteArray::push_back(char c);
extern void demth_ZN10QByteArray9push_backEc(void* qthis, unsigned char arg0);
  // proto:  bool QByteArray::isSharedWith(const QByteArray & other);
extern void demth_ZNK10QByteArray12isSharedWithERKS_(void* qthis, void* arg0);
  // proto:  int QByteArray::size();
extern void demth_ZNK10QByteArray4sizeEv(void* qthis);
  // proto:  bool QByteArray::endsWith(char c);
extern void _ZNK10QByteArray8endsWithEc(void* qthis, unsigned char arg0);
  // proto: static QByteArray QByteArray::number(uint , int base);
extern void _ZN10QByteArray6numberEji(int32_t arg0, int32_t arg1);
  // proto:  void QByteArray::push_front(char c);
extern void demth_ZN10QByteArray10push_frontEc(void* qthis, unsigned char arg0);
  // proto:  QByteArray QByteArray::leftJustified(int width, char fill, bool truncate);
extern void _ZNK10QByteArray13leftJustifiedEicb(void* qthis, int32_t arg0, unsigned char arg1, bool arg2);
  // proto: static QByteArray QByteArray::number(qulonglong , int base);
extern void _ZN10QByteArray6numberEyi(int64_t arg0, int32_t arg1);
  // proto:  int QByteArray::count(char c);
extern void _ZNK10QByteArray5countEc(void* qthis, unsigned char arg0);
  // proto:  double QByteArray::toDouble(bool * ok);
extern void _ZNK10QByteArray8toDoubleEPb(void* qthis, bool* arg0);
  // proto:  QByteArray & QByteArray::replace(int index, int len, const QByteArray & s);
extern void _ZN10QByteArray7replaceEiiRKS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  QByteArray & QByteArray::setNum(short , int base);
extern void demth_ZN10QByteArray6setNumEsi(void* qthis, int16_t arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::prepend(const QByteArray & a);
extern void _ZN10QByteArray7prependERKS_(void* qthis, void* arg0);
  // proto:  qulonglong QByteArray::toULongLong(bool * ok, int base);
extern void _ZNK10QByteArray11toULongLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(char c, const QString & after);
extern void demth_ZN10QByteArray7replaceEcRK7QString(void* qthis, unsigned char arg0, void* arg1);
  // proto: static QByteArray QByteArray::fromPercentEncoding(const QByteArray & pctEncoded, char percent);
extern void _ZN10QByteArray19fromPercentEncodingERKS_c(void* arg0, unsigned char arg1);
  // proto:  void QByteArray::push_front(const char * c);
extern void demth_ZN10QByteArray10push_frontEPKc(void* qthis, unsigned char* arg0);
  // proto:  void QByteArray::clear();
extern void _ZN10QByteArray5clearEv(void* qthis);
  // proto:  qlonglong QByteArray::toLongLong(bool * ok, int base);
extern void _ZNK10QByteArray10toLongLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::prepend(const char * s, int len);
extern void _ZN10QByteArray7prependEPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  const char * QByteArray::constData();
extern void demth_ZNK10QByteArray9constDataEv(void* qthis);
  // proto:  void QByteArray::QByteArray(const QByteArray & );
extern void* dector_ZN10QByteArrayC1ERKS_(void* arg0);
extern void demth_ZN10QByteArrayC1ERKS_(void* qthis, void* arg0);
  // proto:  int QByteArray::length();
extern void demth_ZNK10QByteArray6lengthEv(void* qthis);
  // proto: static QByteArray QByteArray::number(int , int base);
extern void _ZN10QByteArray6numberEii(int32_t arg0, int32_t arg1);
  // proto:  bool QByteArray::startsWith(char c);
extern void _ZNK10QByteArray10startsWithEc(void* qthis, unsigned char arg0);
  // proto:  QByteArray & QByteArray::setNum(double , char f, int prec);
extern void _ZN10QByteArray6setNumEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2);
  // proto: static QByteArray QByteArray::number(qlonglong , int base);
extern void _ZN10QByteArray6numberExi(int64_t arg0, int32_t arg1);
  // proto:  char QByteArray::at(int i);
extern void demth_ZNK10QByteArray2atEi(void* qthis, int32_t arg0);
  // proto:  QByteArray & QByteArray::setNum(ushort , int base);
extern void demth_ZN10QByteArray6setNumEti(void* qthis, int16_t arg0, int32_t arg1);
  // proto:  void QByteArray::swap(QByteArray & other);
extern void demth_ZN10QByteArray4swapERS_(void* qthis, void* arg0);
  // proto:  QByteArray & QByteArray::replace(const QString & before, const char * after);
extern void demth_ZN10QByteArray7replaceERK7QStringPKc(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  QByteArray & QByteArray::append(const QByteArray & a);
extern void _ZN10QByteArray6appendERKS_(void* qthis, void* arg0);
  // proto:  bool QByteArray::endsWith(const QByteArray & a);
extern void _ZNK10QByteArray8endsWithERKS_(void* qthis, void* arg0);
  // proto:  int QByteArray::count(const char * a);
extern void _ZNK10QByteArray5countEPKc(void* qthis, unsigned char* arg0);
  // proto:  QByteArray & QByteArray::replace(const char * before, int bsize, const char * after, int asize);
extern void _ZN10QByteArray7replaceEPKciS1_i(void* qthis, unsigned char* arg0, int32_t arg1, unsigned char* arg2, int32_t arg3);
  // proto:  QList<QByteArray> QByteArray::split(char sep);
extern void _ZNK10QByteArray5splitEc(void* qthis, unsigned char arg0);
  // proto:  QByteArray & QByteArray::setNum(qlonglong , int base);
extern void _ZN10QByteArray6setNumExi(void* qthis, int64_t arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(char before, const char * after);
extern void demth_ZN10QByteArray7replaceEcPKc(void* qthis, unsigned char arg0, unsigned char* arg1);
  // proto:  QByteArray & QByteArray::append(const char * s);
extern void _ZN10QByteArray6appendEPKc(void* qthis, unsigned char* arg0);
  // proto:  QByteArray QByteArray::right(int len);
extern void _ZNK10QByteArray5rightEi(void* qthis, int32_t arg0);
  // proto:  QByteArray & QByteArray::append(const QString & s);
extern void demth_ZN10QByteArray6appendERK7QString(void* qthis, void* arg0);
  // proto:  void QByteArray::chop(int n);
extern void _ZN10QByteArray4chopEi(void* qthis, int32_t arg0);
  // proto:  int QByteArray::lastIndexOf(const char * c, int from);
extern void _ZNK10QByteArray11lastIndexOfEPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(int index, int len, const char * s);
extern void _ZN10QByteArray7replaceEiiPKc(void* qthis, int32_t arg0, int32_t arg1, unsigned char* arg2);
  // proto:  void QByteArray::push_back(const QByteArray & a);
extern void demth_ZN10QByteArray9push_backERKS_(void* qthis, void* arg0);
  // proto:  QByteArray QByteArray::toPercentEncoding(const QByteArray & exclude, const QByteArray & include, char percent);
extern void _ZNK10QByteArray17toPercentEncodingERKS_S1_c(void* qthis, void* arg0, void* arg1, unsigned char arg2);
  // proto:  bool QByteArray::isDetached();
extern void demth_ZNK10QByteArray10isDetachedEv(void* qthis);
  // proto:  QByteArray & QByteArray::append(const char * s, int len);
extern void _ZN10QByteArray6appendEPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(const QByteArray & before, const char * after);
extern void demth_ZN10QByteArray7replaceERKS_PKc(void* qthis, void* arg0, unsigned char* arg1);
  // proto:  QByteArray & QByteArray::setNum(qulonglong , int base);
extern void _ZN10QByteArray6setNumEyi(void* qthis, int64_t arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::setRawData(const char * a, uint n);
extern void _ZN10QByteArray10setRawDataEPKcj(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::replace(const QString & before, const QByteArray & after);
extern void demth_ZN10QByteArray7replaceERK7QStringRKS_(void* qthis, void* arg0, void* arg1);
  // proto:  QByteArray & QByteArray::setNum(uint , int base);
extern void demth_ZN10QByteArray6setNumEji(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QByteArray QByteArray::mid(int index, int len);
extern void _ZNK10QByteArray3midEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::setNum(int , int base);
extern void demth_ZN10QByteArray6setNumEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QByteArray & QByteArray::insert(int i, const QByteArray & a);
extern void _ZN10QByteArray6insertEiRKS_(void* qthis, int32_t arg0, void* arg1);
  // proto:  QByteArray & QByteArray::insert(int i, const char * s);
extern void _ZN10QByteArray6insertEiPKc(void* qthis, int32_t arg0, unsigned char* arg1);
  // proto:  QByteArray & QByteArray::fill(char c, int size);
extern void _ZN10QByteArray4fillEci(void* qthis, unsigned char arg0, int32_t arg1);
  // proto:  ushort QByteArray::toUShort(bool * ok, int base);
extern void _ZNK10QByteArray8toUShortEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  void QByteArray::push_back(const char * c);
extern void demth_ZN10QByteArray9push_backEPKc(void* qthis, unsigned char* arg0);
  // proto:  QByteArray QByteArray::rightJustified(int width, char fill, bool truncate);
extern void _ZNK10QByteArray14rightJustifiedEicb(void* qthis, int32_t arg0, unsigned char arg1, bool arg2);
  // proto:  bool QByteArray::contains(const QByteArray & a);
extern void demth_ZNK10QByteArray8containsERKS_(void* qthis, void* arg0);
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

  // proto:  void QByteRef::QByteRef(QByteArray & array, int idx);
func NewQByteRef(args ...interface{}) QByteRef {
  return QByteRef{}
}

  // proto:  QByteArray & QByteArray::insert(int i, char c);
func (this *QByteArray) insert(args ...interface{}) () {
  // insert(int, char)
  // insert(int, const char *, int)
  // insert(int, const class QString &)
  // insert(int, const class QByteArray &)
  // insert(int, const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6insertEic
    // invoke: QByteArray & insert(int, char)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEic(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QByteArray6insertEiPKci
    // invoke: QByteArray & insert(int, const char *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray6insertEiPKci(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN10QByteArray6insertEiRK7QString
    // invoke: QByteArray & insert(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray6insertEiRK7QString(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArray6insertEiRKS_
    // invoke: QByteArray & insert(int, const class QByteArray &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEiRKS_(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN10QByteArray6insertEiPKc
    // invoke: QByteArray & insert(int, const char *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6insertEiPKc(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "insert", args)
  }

}

  // proto:  int QByteArray::lastIndexOf(const QByteArray & a, int from);
func (this *QByteArray) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QByteArray &, int)
  // lastIndexOf(const class QString &, int)
  // lastIndexOf(char, int)
  // lastIndexOf(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray11lastIndexOfERKS_i
    // invoke: int lastIndexOf(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfERKS_i(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK10QByteArray11lastIndexOfERK7QStringi
    // invoke: int lastIndexOf(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK10QByteArray11lastIndexOfERK7QStringi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK10QByteArray11lastIndexOfEci
    // invoke: int lastIndexOf(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfEci(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK10QByteArray11lastIndexOfEPKci
    // invoke: int lastIndexOf(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray11lastIndexOfEPKci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "lastIndexOf", args)
  }

}

  // proto:  void QByteArray::push_front(const QByteArray & a);
func (this *QByteArray) push_front(args ...interface{}) () {
  // push_front(const class QByteArray &)
  // push_front(char)
  // push_front(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray10push_frontERKS_
    // invoke: void push_front(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray10push_frontERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray10push_frontEc
    // invoke: void push_front(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray10push_frontEc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray10push_frontEPKc
    // invoke: void push_front(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray10push_frontEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "push_front", args)
  }

}

  // proto:  ulong QByteArray::toULong(bool * ok, int base);
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

  // proto:  QByteArray & QByteArray::replace(const char * before, const char * after);
func (this *QByteArray) replace(args ...interface{}) () {
  // replace(const char *, const char *)
  // replace(const class QByteArray &, const class QByteArray &)
  // replace(const char *, const class QByteArray &)
  // replace(int, int, const char *, int)
  // replace(char, const class QByteArray &)
  // replace(char, char)
  // replace(int, int, const class QByteArray &)
  // replace(char, const class QString &)
  // replace(const class QString &, const char *)
  // replace(const char *, int, const char *, int)
  // replace(char, const char *)
  // replace(int, int, const char *)
  // replace(const class QByteArray &, const char *)
  // replace(const class QString &, const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.ByteTy(true) // "const char *"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(false) // "char"
  vtys[4][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(false) // "char"
  vtys[5][1] = qtrt.ByteTy(false) // "char"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "int"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.ByteTy(false) // "char"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][1] = qtrt.ByteTy(true) // "const char *"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.ByteTy(true) // "const char *"
  vtys[9][1] = qtrt.Int32Ty(false) // "int"
  vtys[9][2] = qtrt.ByteTy(true) // "const char *"
  vtys[9][3] = qtrt.Int32Ty(false) // "int"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.ByteTy(false) // "char"
  vtys[10][1] = qtrt.ByteTy(true) // "const char *"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "int"
  vtys[11][1] = qtrt.Int32Ty(false) // "int"
  vtys[11][2] = qtrt.ByteTy(true) // "const char *"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[12][1] = qtrt.ByteTy(true) // "const char *"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[13][1] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7replaceEPKcS1_
    // invoke: QByteArray & replace(const char *, const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray7replaceEPKcS1_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN10QByteArray7replaceERKS_S1_
    // invoke: QByteArray & replace(const class QByteArray &, const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceERKS_S1_(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray7replaceEPKcRKS_
    // invoke: QByteArray & replace(const char *, const class QByteArray &)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEPKcRKS_(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN10QByteArray7replaceEiiPKci
    // invoke: QByteArray & replace(int, int, const char *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QByteArray7replaceEiiPKci(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN10QByteArray7replaceEcRKS_
    // invoke: QByteArray & replace(char, const class QByteArray &)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEcRKS_(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN10QByteArray7replaceEcc
    // invoke: QByteArray & replace(char, char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7replaceEcc(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN10QByteArray7replaceEiiRKS_
    // invoke: QByteArray & replace(int, int, const class QByteArray &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QByteArray).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray7replaceEiiRKS_(this.qclsinst, arg0, arg1, arg2)
  case 7:
    // invoke: _ZN10QByteArray7replaceEcRK7QString
    // invoke: QByteArray & replace(char, const class QString &)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray7replaceEcRK7QString(this.qclsinst, arg0, arg1)
  case 8:
    // invoke: _ZN10QByteArray7replaceERK7QStringPKc
    // invoke: QByteArray & replace(const class QString &, const char *)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray7replaceERK7QStringPKc(this.qclsinst, arg0, arg1)
  case 9:
    // invoke: _ZN10QByteArray7replaceEPKciS1_i
    // invoke: QByteArray & replace(const char *, int, const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN10QByteArray7replaceEPKciS1_i(this.qclsinst, arg0, arg1, arg2, arg3)
  case 10:
    // invoke: _ZN10QByteArray7replaceEcPKc
    // invoke: QByteArray & replace(char, const char *)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray7replaceEcPKc(this.qclsinst, arg0, arg1)
  case 11:
    // invoke: _ZN10QByteArray7replaceEiiPKc
    // invoke: QByteArray & replace(int, int, const char *)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[2].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray7replaceEiiPKc(this.qclsinst, arg0, arg1, arg2)
  case 12:
    // invoke: _ZN10QByteArray7replaceERKS_PKc
    // invoke: QByteArray & replace(const class QByteArray &, const char *)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[1].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray7replaceERKS_PKc(this.qclsinst, arg0, arg1)
  case 13:
    // invoke: _ZN10QByteArray7replaceERK7QStringRKS_
    // invoke: QByteArray & replace(const class QString &, const class QByteArray &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QByteArray).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray7replaceERK7QStringRKS_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "replace", args)
  }

}

  // proto: static QByteArray QByteArray::fromHex(const QByteArray & hexEncoded);
func (this *QByteArray) fromHex_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromHex", args)
  }

}

  // proto:  QByteArray & QByteArray::prepend(const char * s);
func (this *QByteArray) prepend(args ...interface{}) () {
  // prepend(const char *)
  // prepend(char)
  // prepend(const class QByteArray &)
  // prepend(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray7prependEPKc
    // invoke: QByteArray & prepend(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7prependEPKc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray7prependEc
    // invoke: QByteArray & prepend(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7prependEc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray7prependERKS_
    // invoke: QByteArray & prepend(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray7prependERKS_(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN10QByteArray7prependEPKci
    // invoke: QByteArray & prepend(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray7prependEPKci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "prepend", args)
  }

}

  // proto:  int QByteArray::count(const QByteArray & a);
func (this *QByteArray) count(args ...interface{}) () {
  // count(const class QByteArray &)
  // count()
  // count(char)
  // count(const char *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray5countERKS_
    // invoke: int count(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5countERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QByteArray5countEv
    // invoke: int count()
    C.demth_ZNK10QByteArray5countEv(this.qclsinst)
  case 2:
    // invoke: _ZNK10QByteArray5countEc
    // invoke: int count(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5countEc(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK10QByteArray5countEPKc
    // invoke: int count(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray5countEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "count", args)
  }

}

  // proto:  void QByteArray::~QByteArray();
func (this *QByteArray) FreeQByteArray(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "~QByteArray", args)
  }

}

  // proto:  void QByteArray::QByteArray();
func NewQByteArray(args ...interface{}) QByteArray {
  return QByteArray{}
}

  // proto:  float QByteArray::toFloat(bool * ok);
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

  // proto:  void QByteArray::truncate(int pos);
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

  // proto:  QByteArray QByteArray::toBase64();
func (this *QByteArray) toBase64(args ...interface{}) () {
  // toBase64(Base64Options)
  // toBase64()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int64Ty(false) // "Base64Options"
  vtys[1] = make(map[int32]reflect.Type)

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

  // proto:  bool QByteArray::isEmpty();
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
    C.demth_ZNK10QByteArray7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "isEmpty", args)
  }

}

  // proto:  void QByteArray::resize(int size);
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

  // proto:  QByteArray QByteArray::toHex();
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

  // proto:  int QByteArray::indexOf(const char * c, int from);
func (this *QByteArray) indexOf(args ...interface{}) () {
  // indexOf(const char *, int)
  // indexOf(const class QByteArray &, int)
  // indexOf(char, int)
  // indexOf(const class QString &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(false) // "char"
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfEPKci(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK10QByteArray7indexOfERKS_i
    // invoke: int indexOf(const class QByteArray &, int)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfERKS_i(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK10QByteArray7indexOfEci
    // invoke: int indexOf(char, int)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QByteArray7indexOfEci(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK10QByteArray7indexOfERK7QStringi
    // invoke: int indexOf(const class QString &, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZNK10QByteArray7indexOfERK7QStringi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "indexOf", args)
  }

}

  // proto:  uint QByteArray::toUInt(bool * ok, int base);
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

  // proto: static QByteArray QByteArray::fromStdString(const std::string & s);
func (this *QByteArray) fromStdString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromStdString", args)
  }

}

  // proto:  bool QByteArray::isNull();
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

  // proto:  void QByteArray::reserve(int size);
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
    C.demth_ZN10QByteArray7reserveEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "reserve", args)
  }

}

  // proto: static QByteArray QByteArray::fromRawData(const char * , int size);
func (this *QByteArray) fromRawData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromRawData", args)
  }

}

  // proto:  bool QByteArray::contains(char c);
func (this *QByteArray) contains(args ...interface{}) () {
  // contains(char)
  // contains(const char *)
  // contains(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8containsEc
    // invoke: bool contains(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QByteArray8containsEc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QByteArray8containsEPKc
    // invoke: bool contains(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QByteArray8containsEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QByteArray8containsERKS_
    // invoke: bool contains(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QByteArray8containsERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "contains", args)
  }

}

  // proto:  long QByteArray::toLong(bool * ok, int base);
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

  // proto:  char * QByteArray::data();
func (this *QByteArray) data(args ...interface{}) () {
  // data()
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray4dataEv
    // invoke: char * data()
    C.demth_ZN10QByteArray4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "data", args)
  }

}

  // proto: static QByteArray QByteArray::number(double , char f, int prec);
func (this *QByteArray) number_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "number", args)
  }

}

  // proto:  int QByteArray::capacity();
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
    C.demth_ZNK10QByteArray8capacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "capacity", args)
  }

}

  // proto: static QByteArray QByteArray::fromBase64(const QByteArray & base64);
func (this *QByteArray) fromBase64_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromBase64", args)
  }

}

  // proto:  QByteArray QByteArray::left(int len);
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

  // proto:  QByteArray & QByteArray::append(char c);
func (this *QByteArray) append(args ...interface{}) () {
  // append(char)
  // append(const class QByteArray &)
  // append(const char *)
  // append(const class QString &)
  // append(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "char"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"

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
    // invoke: _ZN10QByteArray6appendERKS_
    // invoke: QByteArray & append(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6appendERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray6appendEPKc
    // invoke: QByteArray & append(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZN10QByteArray6appendEPKc(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN10QByteArray6appendERK7QString
    // invoke: QByteArray & append(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray6appendERK7QString(this.qclsinst, arg0)
  case 4:
    // invoke: _ZN10QByteArray6appendEPKci
    // invoke: QByteArray & append(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6appendEPKci(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "append", args)
  }

}

  // proto:  bool QByteArray::startsWith(const char * c);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
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

  // proto:  QByteArray & QByteArray::remove(int index, int len);
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

  // proto:  bool QByteArray::endsWith(const char * c);
func (this *QByteArray) endsWith(args ...interface{}) () {
  // endsWith(const char *)
  // endsWith(char)
  // endsWith(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(false) // "char"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QByteArray8endsWithEPKc
    // invoke: bool endsWith(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8endsWithEPKc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK10QByteArray8endsWithEc
    // invoke: bool endsWith(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8endsWithEc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZNK10QByteArray8endsWithERKS_
    // invoke: bool endsWith(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QByteArray8endsWithERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "endsWith", args)
  }

}

  // proto:  void QByteArray::squeeze();
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
    C.demth_ZN10QByteArray7squeezeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "squeeze", args)
  }

}

  // proto:  void QByteArray::detach();
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
    C.demth_ZN10QByteArray6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "detach", args)
  }

}

  // proto:  QByteArray QByteArray::repeated(int times);
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

  // proto:  QByteArray & QByteArray::setNum(float , char f, int prec);
func (this *QByteArray) setNum(args ...interface{}) () {
  // setNum(float, char, int)
  // setNum(short, int)
  // setNum(double, char, int)
  // setNum(ushort, int)
  // setNum(qlonglong, int)
  // setNum(qulonglong, int)
  // setNum(uint, int)
  // setNum(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.FloatTy(false) // "float"
  vtys[0][1] = qtrt.ByteTy(false) // "char"
  vtys[0][2] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int16Ty(false) // "short"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.DoubleTy(false) // "double"
  vtys[2][1] = qtrt.ByteTy(false) // "char"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "uint"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QByteArray6setNumEfci
    // invoke: QByteArray & setNum(float, char, int)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN10QByteArray6setNumEfci(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZN10QByteArray6setNumEsi
    // invoke: QByteArray & setNum(short, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray6setNumEsi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN10QByteArray6setNumEdci
    // invoke: QByteArray & setNum(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN10QByteArray6setNumEdci(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN10QByteArray6setNumEti
    // invoke: QByteArray & setNum(ushort, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray6setNumEti(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN10QByteArray6setNumExi
    // invoke: QByteArray & setNum(qlonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumExi(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN10QByteArray6setNumEyi
    // invoke: QByteArray & setNum(qulonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray6setNumEyi(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN10QByteArray6setNumEji
    // invoke: QByteArray & setNum(uint, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray6setNumEji(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN10QByteArray6setNumEii
    // invoke: QByteArray & setNum(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN10QByteArray6setNumEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "setNum", args)
  }

}

  // proto:  short QByteArray::toShort(bool * ok, int base);
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

  // proto:  int QByteArray::toInt(bool * ok, int base);
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

  // proto:  void QByteArray::push_back(char c);
func (this *QByteArray) push_back(args ...interface{}) () {
  // push_back(char)
  // push_back(const class QByteArray &)
  // push_back(const char *)
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
    // invoke: _ZN10QByteArray9push_backEc
    // invoke: void push_back(char)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray9push_backEc(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN10QByteArray9push_backERKS_
    // invoke: void push_back(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray9push_backERKS_(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN10QByteArray9push_backEPKc
    // invoke: void push_back(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZN10QByteArray9push_backEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "push_back", args)
  }

}

  // proto:  bool QByteArray::isSharedWith(const QByteArray & other);
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
    C.demth_ZNK10QByteArray12isSharedWithERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "isSharedWith", args)
  }

}

  // proto:  int QByteArray::size();
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
    C.demth_ZNK10QByteArray4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "size", args)
  }

}

  // proto:  QByteArray QByteArray::leftJustified(int width, char fill, bool truncate);
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

  // proto:  double QByteArray::toDouble(bool * ok);
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

  // proto:  qulonglong QByteArray::toULongLong(bool * ok, int base);
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

  // proto: static QByteArray QByteArray::fromPercentEncoding(const QByteArray & pctEncoded, char percent);
func (this *QByteArray) fromPercentEncoding_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QByteArray", "fromPercentEncoding", args)
  }

}

  // proto:  void QByteArray::clear();
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

  // proto:  qlonglong QByteArray::toLongLong(bool * ok, int base);
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

  // proto:  const char * QByteArray::constData();
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
    C.demth_ZNK10QByteArray9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "constData", args)
  }

}

  // proto:  int QByteArray::length();
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
    C.demth_ZNK10QByteArray6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "length", args)
  }

}

  // proto:  char QByteArray::at(int i);
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
    C.demth_ZNK10QByteArray2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "at", args)
  }

}

  // proto:  void QByteArray::swap(QByteArray & other);
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
    C.demth_ZN10QByteArray4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QByteArray", "swap", args)
  }

}

  // proto:  QList<QByteArray> QByteArray::split(char sep);
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

  // proto:  QByteArray QByteArray::right(int len);
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

  // proto:  void QByteArray::chop(int n);
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

  // proto:  QByteArray QByteArray::toPercentEncoding(const QByteArray & exclude, const QByteArray & include, char percent);
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

  // proto:  bool QByteArray::isDetached();
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
    C.demth_ZNK10QByteArray10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QByteArray", "isDetached", args)
  }

}

  // proto:  QByteArray & QByteArray::setRawData(const char * a, uint n);
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN10QByteArray10setRawDataEPKcj(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QByteArray", "setRawData", args)
  }

}

  // proto:  QByteArray QByteArray::mid(int index, int len);
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

  // proto:  QByteArray & QByteArray::fill(char c, int size);
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

  // proto:  ushort QByteArray::toUShort(bool * ok, int base);
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

  // proto:  QByteArray QByteArray::rightJustified(int width, char fill, bool truncate);
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

// <= body block end

