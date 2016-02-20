package qtcore
// auto generated, do not modify.
// created: Sat Feb 20 11:35:41 2016
// src-file: /QtCore/qstring.h
// dst-file: /src/core/qstring.go
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
  // proto:  QVector<uint> QString::toUcs4();
extern void C_ZNK7QString6toUcs4Ev(void* qthis); // 4
  // proto:  int QString::toWCharArray(wchar_t * array);
extern int32_t C_ZNK7QString12toWCharArrayEPw(void* qthis, void* arg0); // 2
  // proto:  const QChar * QString::unicode();
extern void* C_ZNK7QString7unicodeEv(void* qthis); // 2
  // proto: static QString QString::fromLatin1(const QByteArray & str);
extern void* C_ZN7QString10fromLatin1ERK10QByteArray(void* arg0); // 2
  // proto: static QString QString::fromLatin1(const char * str, int size);
extern void* C_ZN7QString10fromLatin1EPKci(void* arg0, int32_t arg1); // 2
  // proto:  void QString::reserve(int size);
extern void C_ZN7QString7reserveEi(void* qthis, int32_t arg0); // 2
  // proto:  void QString::swap(QString & other);
extern void C_ZN7QString4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QStringRef QString::rightRef(int n);
extern void C_ZNK7QString8rightRefEi(void* qthis, int32_t arg0); // 4
  // proto:  QStringRef QString::leftRef(int n);
extern void C_ZNK7QString7leftRefEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QString::rightJustified(int width, QChar fill, bool trunc);
extern void* C_ZNK7QString14rightJustifiedEi5QCharb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  void QString::detach();
extern void C_ZN7QString6detachEv(void* qthis); // 2
  // proto:  QString & QString::insert(int i, QChar c);
extern void* C_ZN7QString6insertEi5QChar(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString & QString::insert(int i, const QString & s);
extern void* C_ZN7QString6insertEiRKS_(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QString & QString::insert(int i, const QChar * uc, int len);
extern void* C_ZN7QString6insertEiPK5QChari(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  QStringRef QString::midRef(int position, int n);
extern void C_ZNK7QString6midRefEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::remove(const QRegExp & rx);
extern void* C_ZN7QString6removeERK7QRegExp(void* qthis, void* arg0); // 2
  // proto:  QString & QString::remove(const QRegularExpression & re);
extern void* C_ZN7QString6removeERK18QRegularExpression(void* qthis, void* arg0); // 2
  // proto:  QString & QString::remove(int i, int len);
extern void* C_ZN7QString6removeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QString QString::trimmed();
extern void* C_ZNKR7QString7trimmedEv(void* qthis); // 2
  // proto:  std::wstring QString::toStdWString();
extern void C_ZNK7QString12toStdWStringB5cxx11Ev(void* qthis); // 2
  // proto:  const_iterator QString::constEnd();
extern void* C_ZNK7QString8constEndEv(void* qthis); // 2
  // proto:  const ushort * QString::utf16();
extern void C_ZNK7QString5utf16Ev(void* qthis); // 4
  // proto:  QString QString::right(int n);
extern void* C_ZNK7QString5rightEi(void* qthis, int32_t arg0); // 4
  // proto:  QByteArray QString::toLocal8Bit();
extern void* C_ZNKR7QString11toLocal8BitEv(void* qthis); // 2
  // proto:  uint QString::toUInt(bool * ok, int base);
extern int32_t C_ZNK7QString6toUIntEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QString::QString();
extern void* C_ZN7QStringC2Ev(); // 1
  // proto:  void QString::QString(const QString & );
extern void* C_ZN7QStringC2ERKS_(void* arg0); // 1
  // proto:  void QString::QString(int size, QChar c);
extern void* C_ZN7QStringC2Ei5QChar(int32_t arg0, void* arg1); // 3
  // proto:  void QString::QString(const char * ch);
extern void* C_ZN7QStringC2EPKc(void* arg0); // 1
  // proto:  void QString::QString(const QChar * unicode, int size);
extern void* C_ZN7QStringC2EPK5QChari(void* arg0, int32_t arg1); // 3
  // proto:  void QString::QString(QChar c);
extern void* C_ZN7QStringC2E5QChar(void* arg0); // 3
  // proto:  void QString::QString(const QByteArray & a);
extern void* C_ZN7QStringC2ERK10QByteArray(void* arg0); // 1
  // proto:  QString QString::arg(const QString & a, int fieldWidth, QChar fillChar);
extern void* C_ZNK7QString3argERKS_i5QChar(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4);
extern void* C_ZNK7QString3argERKS_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3);
extern void* C_ZNK7QString3argERKS_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  QString QString::arg(qlonglong a, int fieldwidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argExii5QChar(void* qthis, int64_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 4
  // proto:  QString QString::arg(long a, int fieldwidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argElii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(ushort a, int fieldWidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argEtii5QChar(void* qthis, int16_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(int a, int fieldWidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argEiii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8);
extern void* C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7); // 2
  // proto:  QString QString::arg(short a, int fieldWidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argEsii5QChar(void* qthis, int16_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7);
extern void* C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5);
extern void* C_ZNK7QString3argERKS_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 2
  // proto:  QString QString::arg(uint a, int fieldWidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argEjii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2);
extern void* C_ZNK7QString3argERKS_S1_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QString QString::arg(double a, int fieldWidth, char fmt, int prec, QChar fillChar);
extern void* C_ZNK7QString3argEdici5QChar(void* qthis, double arg0, int32_t arg1, unsigned char arg2, int32_t arg3, void* arg4); // 4
  // proto:  QString QString::arg(qulonglong a, int fieldwidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argEyii5QChar(void* qthis, int64_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 4
  // proto:  QString QString::arg(char a, int fieldWidth, QChar fillChar);
extern void* C_ZNK7QString3argEci5QChar(void* qthis, unsigned char arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8, const QString & a9);
extern void* C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8); // 2
  // proto:  QString QString::arg(ulong a, int fieldwidth, int base, QChar fillChar);
extern void* C_ZNK7QString3argEmii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(QChar a, int fieldWidth, QChar fillChar);
extern void* C_ZNK7QString3argE5QChariS0_(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6);
extern void* C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5); // 2
  // proto:  QString & QString::append(QChar c);
extern void* C_ZN7QString6appendE5QChar(void* qthis, void* arg0); // 4
  // proto:  QString & QString::append(const char * s);
extern void* C_ZN7QString6appendEPKc(void* qthis, void* arg0); // 2
  // proto:  QString & QString::append(const QByteArray & s);
extern void* C_ZN7QString6appendERK10QByteArray(void* qthis, void* arg0); // 2
  // proto:  QString & QString::append(const QString & s);
extern void* C_ZN7QString6appendERKS_(void* qthis, void* arg0); // 4
  // proto:  QString & QString::append(const QChar * uc, int len);
extern void* C_ZN7QString6appendEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::capacity();
extern int32_t C_ZNK7QString8capacityEv(void* qthis); // 2
  // proto:  QString QString::toLower();
extern void* C_ZNKR7QString7toLowerEv(void* qthis); // 2
  // proto:  QString QString::toHtmlEscaped();
extern void* C_ZNK7QString13toHtmlEscapedEv(void* qthis); // 4
  // proto:  bool QString::isNull();
extern bool C_ZNK7QString6isNullEv(void* qthis); // 2
  // proto: static int QString::localeAwareCompare(const QString & s1, const QString & s2);
extern int32_t C_ZN7QString18localeAwareCompareERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  int QString::localeAwareCompare(const QString & s);
extern int32_t C_ZNK7QString18localeAwareCompareERKS_(void* qthis, void* arg0); // 4
  // proto: static QString QString::fromLocal8Bit(const QByteArray & str);
extern void* C_ZN7QString13fromLocal8BitERK10QByteArray(void* arg0); // 2
  // proto: static QString QString::fromLocal8Bit(const char * str, int size);
extern void* C_ZN7QString13fromLocal8BitEPKci(void* arg0, int32_t arg1); // 2
  // proto: static QString QString::fromWCharArray(const wchar_t * string, int size);
extern void* C_ZN7QString14fromWCharArrayEPKwi(void* arg0, int32_t arg1); // 2
  // proto:  std::string QString::toStdString();
extern void C_ZNK7QString11toStdStringB5cxx11Ev(void* qthis); // 2
  // proto:  int QString::indexOf(const QRegularExpression & re, int from, QRegularExpressionMatch * rmatch);
extern int32_t C_ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QString::indexOf(const QRegularExpression & re, int from);
extern int32_t C_ZNK7QString7indexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::indexOf(QRegExp & , int from);
extern int32_t C_ZNK7QString7indexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::indexOf(const QRegExp & , int from);
extern int32_t C_ZNK7QString7indexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  const_iterator QString::cbegin();
extern void* C_ZNK7QString6cbeginEv(void* qthis); // 2
  // proto:  QByteArray QString::toLatin1();
extern void* C_ZNKR7QString8toLatin1Ev(void* qthis); // 2
  // proto:  void QString::chop(int n);
extern void C_ZN7QString4chopEi(void* qthis, int32_t arg0); // 4
  // proto:  std::u32string QString::toStdU32String();
extern void C_ZNK7QString14toStdU32StringB5cxx11Ev(void* qthis); // 2
  // proto:  int QString::length();
extern int32_t C_ZNK7QString6lengthEv(void* qthis); // 2
  // proto:  double QString::toDouble(bool * ok);
extern double C_ZNK7QString8toDoubleEPb(void* qthis, void* arg0); // 4
  // proto:  bool QString::isSharedWith(const QString & other);
extern bool C_ZNK7QString12isSharedWithERKS_(void* qthis, void* arg0); // 2
  // proto:  ushort QString::toUShort(bool * ok, int base);
extern int16_t C_ZNK7QString8toUShortEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  const_iterator QString::constBegin();
extern void* C_ZNK7QString10constBeginEv(void* qthis); // 2
  // proto:  short QString::toShort(bool * ok, int base);
extern int16_t C_ZNK7QString7toShortEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto: static QString QString::number(uint , int base);
extern void* C_ZN7QString6numberEji(int32_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(double , char f, int prec);
extern void* C_ZN7QString6numberEdci(double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto: static QString QString::number(qulonglong , int base);
extern void* C_ZN7QString6numberEyi(int64_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(ulong , int base);
extern void* C_ZN7QString6numberEmi(int32_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(qlonglong , int base);
extern void* C_ZN7QString6numberExi(int64_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(long , int base);
extern void* C_ZN7QString6numberEli(int32_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(int , int base);
extern void* C_ZN7QString6numberEii(int32_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::replace(const QRegExp & rx, const QString & after);
extern void* C_ZN7QString7replaceERK7QRegExpRKS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString & QString::replace(int i, int len, QChar after);
extern void* C_ZN7QString7replaceEii5QChar(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString & QString::replace(const QRegularExpression & re, const QString & after);
extern void* C_ZN7QString7replaceERK18QRegularExpressionRKS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString & QString::replace(int i, int len, const QChar * s, int slen);
extern void* C_ZN7QString7replaceEiiPK5QChari(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  QString & QString::replace(int i, int len, const QString & after);
extern void* C_ZN7QString7replaceEiiRKS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QString::push_back(QChar c);
extern void C_ZN7QString9push_backE5QChar(void* qthis, void* arg0); // 2
  // proto:  void QString::push_back(const QString & s);
extern void C_ZN7QString9push_backERKS_(void* qthis, void* arg0); // 2
  // proto:  QString QString::simplified();
extern void* C_ZNKR7QString10simplifiedEv(void* qthis); // 2
  // proto:  int QString::size();
extern int32_t C_ZNK7QString4sizeEv(void* qthis); // 2
  // proto:  void QString::truncate(int pos);
extern void C_ZN7QString8truncateEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QString::contains(const QRegularExpression & re);
extern bool C_ZNK7QString8containsERK18QRegularExpression(void* qthis, void* arg0); // 4
  // proto:  bool QString::contains(const QRegularExpression & re, QRegularExpressionMatch * match);
extern bool C_ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QString::contains(QRegExp & rx);
extern bool C_ZNK7QString8containsER7QRegExp(void* qthis, void* arg0); // 2
  // proto:  bool QString::contains(const QRegExp & rx);
extern bool C_ZNK7QString8containsERK7QRegExp(void* qthis, void* arg0); // 2
  // proto:  bool QString::isSimpleText();
extern bool C_ZNK7QString12isSimpleTextEv(void* qthis); // 4
  // proto:  const_iterator QString::cend();
extern void* C_ZNK7QString4cendEv(void* qthis); // 2
  // proto:  QString QString::leftJustified(int width, QChar fill, bool trunc);
extern void* C_ZNK7QString13leftJustifiedEi5QCharb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  QString QString::repeated(int times);
extern void* C_ZNK7QString8repeatedEi(void* qthis, int32_t arg0); // 4
  // proto: static QString QString::fromRawData(const QChar * , int size);
extern void* C_ZN7QString11fromRawDataEPK5QChari(void* arg0, int32_t arg1); // 4
  // proto:  void QString::squeeze();
extern void C_ZN7QString7squeezeEv(void* qthis); // 2
  // proto:  void QString::resize(int size);
extern void C_ZN7QString6resizeEi(void* qthis, int32_t arg0); // 4
  // proto:  qulonglong QString::toULongLong(bool * ok, int base);
extern int64_t C_ZNK7QString11toULongLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::count(const QRegularExpression & re);
extern int32_t C_ZNK7QString5countERK18QRegularExpression(void* qthis, void* arg0); // 4
  // proto:  int QString::count();
extern int32_t C_ZNK7QString5countEv(void* qthis); // 2
  // proto:  int QString::count(const QRegExp & );
extern int32_t C_ZNK7QString5countERK7QRegExp(void* qthis, void* arg0); // 4
  // proto:  bool QString::isRightToLeft();
extern bool C_ZNK7QString13isRightToLeftEv(void* qthis); // 4
  // proto:  QByteArray QString::toUtf8();
extern void* C_ZNKR7QString6toUtf8Ev(void* qthis); // 2
  // proto: static QString QString::fromUtf8(const char * str, int size);
extern void* C_ZN7QString8fromUtf8EPKci(void* arg0, int32_t arg1); // 2
  // proto: static QString QString::fromUtf8(const QByteArray & str);
extern void* C_ZN7QString8fromUtf8ERK10QByteArray(void* arg0); // 2
  // proto:  std::u16string QString::toStdU16String();
extern void C_ZNK7QString14toStdU16StringB5cxx11Ev(void* qthis); // 2
  // proto:  float QString::toFloat(bool * ok);
extern float C_ZNK7QString7toFloatEPb(void* qthis, void* arg0); // 4
  // proto:  const QChar QString::at(int i);
extern void* C_ZNK7QString2atEi(void* qthis, int32_t arg0); // 2
  // proto:  QString & QString::fill(QChar c, int size);
extern void* C_ZN7QString4fillE5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  ulong QString::toULong(bool * ok, int base);
extern int32_t C_ZNK7QString7toULongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  iterator QString::end();
extern void* C_ZN7QString3endEv(void* qthis); // 2
  // proto: static QString QString::fromUcs4(const uint * , int size);
extern void* C_ZN7QString8fromUcs4EPKji(void* arg0, int32_t arg1); // 4
  // proto:  QString QString::mid(int position, int n);
extern void* C_ZNK7QString3midEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::prepend(const QString & s);
extern void* C_ZN7QString7prependERKS_(void* qthis, void* arg0); // 2
  // proto:  QString & QString::prepend(QChar c);
extern void* C_ZN7QString7prependE5QChar(void* qthis, void* arg0); // 2
  // proto:  QString & QString::prepend(const char * s);
extern void* C_ZN7QString7prependEPKc(void* qthis, void* arg0); // 2
  // proto:  QString & QString::prepend(const QByteArray & s);
extern void* C_ZN7QString7prependERK10QByteArray(void* qthis, void* arg0); // 2
  // proto:  const QChar * QString::constData();
extern void* C_ZNK7QString9constDataEv(void* qthis); // 2
  // proto:  qlonglong QString::toLongLong(bool * ok, int base);
extern int64_t C_ZNK7QString10toLongLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QString & QString::setUnicode(const QChar * unicode, int size);
extern void* C_ZN7QString10setUnicodeEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto: static QString QString::fromUtf16(const ushort * , int size);
extern void* C_ZN7QString9fromUtf16EPKti(void* arg0, int32_t arg1); // 4
  // proto:  bool QString::isEmpty();
extern bool C_ZNK7QString7isEmptyEv(void* qthis); // 2
  // proto:  QString & QString::setRawData(const QChar * unicode, int size);
extern void* C_ZN7QString10setRawDataEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QString::isDetached();
extern bool C_ZNK7QString10isDetachedEv(void* qthis); // 2
  // proto:  iterator QString::begin();
extern void* C_ZN7QString5beginEv(void* qthis); // 2
  // proto:  int QString::lastIndexOf(QRegExp & , int from);
extern int32_t C_ZNK7QString11lastIndexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::lastIndexOf(const QRegExp & , int from);
extern int32_t C_ZNK7QString11lastIndexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from);
extern int32_t C_ZNK7QString11lastIndexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from, QRegularExpressionMatch * rmatch);
extern int32_t C_ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString & QString::setNum(ushort , int base);
extern void* C_ZN7QString6setNumEti(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(qulonglong , int base);
extern void* C_ZN7QString6setNumEyi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::setNum(qlonglong , int base);
extern void* C_ZN7QString6setNumExi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::setNum(ulong , int base);
extern void* C_ZN7QString6setNumEmi(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(short , int base);
extern void* C_ZN7QString6setNumEsi(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(long , int base);
extern void* C_ZN7QString6setNumEli(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(uint , int base);
extern void* C_ZN7QString6setNumEji(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(int , int base);
extern void* C_ZN7QString6setNumEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(double , char f, int prec);
extern void* C_ZN7QString6setNumEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QString & QString::setNum(float , char f, int prec);
extern void* C_ZN7QString6setNumEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2); // 2
  // proto:  QString & QString::setUtf16(const ushort * utf16, int size);
extern void* C_ZN7QString8setUtf16EPKti(void* qthis, void* arg0, int32_t arg1); // 2
  // proto:  QChar * QString::data();
extern void* C_ZN7QString4dataEv(void* qthis); // 2
  // proto:  long QString::toLong(bool * ok, int base);
extern int32_t C_ZNK7QString6toLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  void QString::~QString();
extern void C_ZN7QStringD2Ev(void* qthis); // 2
  // proto:  void QString::clear();
extern void C_ZN7QString5clearEv(void* qthis); // 2
  // proto:  int QString::toInt(bool * ok, int base);
extern int32_t C_ZNK7QString5toIntEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QString QString::toCaseFolded();
extern void* C_ZNKR7QString12toCaseFoldedEv(void* qthis); // 2
  // proto:  QString QString::toUpper();
extern void* C_ZNKR7QString7toUpperEv(void* qthis); // 2
  // proto:  void QString::push_front(QChar c);
extern void C_ZN7QString10push_frontE5QChar(void* qthis, void* arg0); // 2
  // proto:  void QString::push_front(const QString & s);
extern void C_ZN7QString10push_frontERKS_(void* qthis, void* arg0); // 2
  // proto:  QString QString::left(int n);
extern void* C_ZNK7QString4leftEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLatin1String::QLatin1String(const QByteArray & s);
extern void* C_ZN13QLatin1StringC2ERK10QByteArray(void* arg0); // 1
  // proto:  void QLatin1String::QLatin1String(const char * s);
extern void* C_ZN13QLatin1StringC2EPKc(void* arg0); // 1
  // proto:  void QLatin1String::QLatin1String(const char * s, int sz);
extern void* C_ZN13QLatin1StringC2EPKci(void* arg0, int32_t arg1); // 1
  // proto:  const char * QLatin1String::latin1();
extern void* C_ZNK13QLatin1String6latin1Ev(void* qthis); // 2
  // proto:  const char * QLatin1String::data();
extern void* C_ZNK13QLatin1String4dataEv(void* qthis); // 2
  // proto:  int QLatin1String::size();
extern int32_t C_ZNK13QLatin1String4sizeEv(void* qthis); // 2
  // proto:  bool QCharRef::isLetterOrNumber();
extern bool C_ZN8QCharRef16isLetterOrNumberEv(void* qthis); // 2
  // proto:  bool QCharRef::isLetter();
extern bool C_ZNK8QCharRef8isLetterEv(void* qthis); // 2
  // proto:  bool QCharRef::hasMirrored();
extern bool C_ZNK8QCharRef11hasMirroredEv(void* qthis); // 2
  // proto:  QChar::UnicodeVersion QCharRef::unicodeVersion();
extern void C_ZNK8QCharRef14unicodeVersionEv(void* qthis); // 2
  // proto:  bool QCharRef::isTitleCase();
extern bool C_ZNK8QCharRef11isTitleCaseEv(void* qthis); // 2
  // proto:  void QCharRef::setCell(uchar cell);
extern void C_ZN8QCharRef7setCellEh(void* qthis, unsigned char arg0); // 2
  // proto:  ushort & QCharRef::unicode();
extern void* C_ZN8QCharRef7unicodeEv(void* qthis); // 2
  // proto:  bool QCharRef::isSpace();
extern bool C_ZNK8QCharRef7isSpaceEv(void* qthis); // 2
  // proto:  uchar QCharRef::row();
extern unsigned char C_ZNK8QCharRef3rowEv(void* qthis); // 2
  // proto:  QChar::Category QCharRef::category();
extern void C_ZNK8QCharRef8categoryEv(void* qthis); // 2
  // proto:  bool QCharRef::isUpper();
extern bool C_ZNK8QCharRef7isUpperEv(void* qthis); // 2
  // proto:  QChar QCharRef::toLower();
extern void* C_ZNK8QCharRef7toLowerEv(void* qthis); // 2
  // proto:  QChar::Script QCharRef::script();
extern void C_ZNK8QCharRef6scriptEv(void* qthis); // 2
  // proto:  QChar::JoiningType QCharRef::joiningType();
extern void C_ZNK8QCharRef11joiningTypeEv(void* qthis); // 2
  // proto:  uchar QCharRef::cell();
extern unsigned char C_ZNK8QCharRef4cellEv(void* qthis); // 2
  // proto:  int QCharRef::digitValue();
extern int32_t C_ZNK8QCharRef10digitValueEv(void* qthis); // 2
  // proto:  QString QCharRef::decomposition();
extern void* C_ZNK8QCharRef13decompositionEv(void* qthis); // 2
  // proto:  QChar::Direction QCharRef::direction();
extern void C_ZNK8QCharRef9directionEv(void* qthis); // 2
  // proto:  QChar QCharRef::mirroredChar();
extern void* C_ZNK8QCharRef12mirroredCharEv(void* qthis); // 2
  // proto:  uchar QCharRef::combiningClass();
extern unsigned char C_ZNK8QCharRef14combiningClassEv(void* qthis); // 2
  // proto:  QChar::Decomposition QCharRef::decompositionTag();
extern void C_ZNK8QCharRef16decompositionTagEv(void* qthis); // 2
  // proto:  bool QCharRef::isMark();
extern bool C_ZNK8QCharRef6isMarkEv(void* qthis); // 2
  // proto:  bool QCharRef::isLower();
extern bool C_ZNK8QCharRef7isLowerEv(void* qthis); // 2
  // proto:  bool QCharRef::isNumber();
extern bool C_ZNK8QCharRef8isNumberEv(void* qthis); // 2
  // proto:  QChar QCharRef::toTitleCase();
extern void* C_ZNK8QCharRef11toTitleCaseEv(void* qthis); // 2
  // proto:  char QCharRef::toLatin1();
extern unsigned char C_ZNK8QCharRef8toLatin1Ev(void* qthis); // 2
  // proto:  bool QCharRef::isNull();
extern bool C_ZNK8QCharRef6isNullEv(void* qthis); // 2
  // proto:  bool QCharRef::isPrint();
extern bool C_ZNK8QCharRef7isPrintEv(void* qthis); // 2
  // proto:  QChar QCharRef::toUpper();
extern void* C_ZNK8QCharRef7toUpperEv(void* qthis); // 2
  // proto:  bool QCharRef::isDigit();
extern bool C_ZNK8QCharRef7isDigitEv(void* qthis); // 2
  // proto:  void QCharRef::setRow(uchar row);
extern void C_ZN8QCharRef6setRowEh(void* qthis, unsigned char arg0); // 2
  // proto:  bool QCharRef::isPunct();
extern bool C_ZNK8QCharRef7isPunctEv(void* qthis); // 2
  // proto:  QChar::Joining QCharRef::joining();
extern void C_ZNK8QCharRef7joiningEv(void* qthis); // 2
  // proto:  QVector<uint> QStringRef::toUcs4();
extern void C_ZNK10QStringRef6toUcs4Ev(void* qthis); // 4
  // proto:  int QStringRef::localeAwareCompare(const QString & s);
extern int32_t C_ZNK10QStringRef18localeAwareCompareERK7QString(void* qthis, void* arg0); // 2
  // proto:  QByteArray QStringRef::toLocal8Bit();
extern void* C_ZNK10QStringRef11toLocal8BitEv(void* qthis); // 4
  // proto:  uint QStringRef::toUInt(bool * ok, int base);
extern int32_t C_ZNK10QStringRef6toUIntEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  short QStringRef::toShort(bool * ok, int base);
extern int16_t C_ZNK10QStringRef7toShortEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  ushort QStringRef::toUShort(bool * ok, int base);
extern int16_t C_ZNK10QStringRef8toUShortEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringRef QStringRef::right(int n);
extern void C_ZNK10QStringRef5rightEi(void* qthis, int32_t arg0); // 4
  // proto:  const QChar QStringRef::at(int i);
extern void* C_ZNK10QStringRef2atEi(void* qthis, int32_t arg0); // 2
  // proto:  const QChar * QStringRef::unicode();
extern void* C_ZNK10QStringRef7unicodeEv(void* qthis); // 2
  // proto:  int QStringRef::size();
extern int32_t C_ZNK10QStringRef4sizeEv(void* qthis); // 2
  // proto:  int QStringRef::length();
extern int32_t C_ZNK10QStringRef6lengthEv(void* qthis); // 2
  // proto:  QStringRef QStringRef::trimmed();
extern void C_ZNK10QStringRef7trimmedEv(void* qthis); // 4
  // proto:  QStringRef QStringRef::mid(int pos, int n);
extern void C_ZNK10QStringRef3midEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QChar * QStringRef::constData();
extern void* C_ZNK10QStringRef9constDataEv(void* qthis); // 2
  // proto:  qlonglong QStringRef::toLongLong(bool * ok, int base);
extern int64_t C_ZNK10QStringRef10toLongLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QStringRef QStringRef::appendTo(QString * string);
extern void C_ZNK10QStringRef8appendToEP7QString(void* qthis, void* arg0); // 4
  // proto:  bool QStringRef::isEmpty();
extern bool C_ZNK10QStringRef7isEmptyEv(void* qthis); // 2
  // proto:  QString QStringRef::toString();
extern void* C_ZNK10QStringRef8toStringEv(void* qthis); // 4
  // proto:  const QChar * QStringRef::cend();
extern void* C_ZNK10QStringRef4cendEv(void* qthis); // 2
  // proto:  void QStringRef::~QStringRef();
extern void C_ZN10QStringRefD2Ev(void* qthis); // 2
  // proto:  const QChar * QStringRef::begin();
extern void* C_ZNK10QStringRef5beginEv(void* qthis); // 2
  // proto:  const QString * QStringRef::string();
extern void* C_ZNK10QStringRef6stringEv(void* qthis); // 2
  // proto:  ulong QStringRef::toULong(bool * ok, int base);
extern int32_t C_ZNK10QStringRef7toULongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  float QStringRef::toFloat(bool * ok);
extern float C_ZNK10QStringRef7toFloatEPb(void* qthis, void* arg0); // 4
  // proto:  const QChar * QStringRef::cbegin();
extern void* C_ZNK10QStringRef6cbeginEv(void* qthis); // 2
  // proto:  const QChar * QStringRef::end();
extern void* C_ZNK10QStringRef3endEv(void* qthis); // 2
  // proto:  const QChar * QStringRef::data();
extern void* C_ZNK10QStringRef4dataEv(void* qthis); // 2
  // proto:  qulonglong QStringRef::toULongLong(bool * ok, int base);
extern int64_t C_ZNK10QStringRef11toULongLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QStringRef::count();
extern int32_t C_ZNK10QStringRef5countEv(void* qthis); // 2
  // proto:  void QStringRef::clear();
extern void C_ZN10QStringRef5clearEv(void* qthis); // 2
  // proto:  int QStringRef::toInt(bool * ok, int base);
extern int32_t C_ZNK10QStringRef5toIntEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QByteArray QStringRef::toLatin1();
extern void* C_ZNK10QStringRef8toLatin1Ev(void* qthis); // 4
  // proto:  bool QStringRef::isNull();
extern bool C_ZNK10QStringRef6isNullEv(void* qthis); // 2
  // proto:  double QStringRef::toDouble(bool * ok);
extern double C_ZNK10QStringRef8toDoubleEPb(void* qthis, void* arg0); // 4
  // proto:  void QStringRef::QStringRef(const QString * string);
extern void* C_ZN10QStringRefC2EPK7QString(void* arg0); // 1
  // proto:  void QStringRef::QStringRef();
extern void* C_ZN10QStringRefC2Ev(); // 1
  // proto:  void QStringRef::QStringRef(const QString * string, int position, int size);
extern void* C_ZN10QStringRefC2EPK7QStringii(void* arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  int QStringRef::position();
extern int32_t C_ZNK10QStringRef8positionEv(void* qthis); // 2
  // proto:  long QStringRef::toLong(bool * ok, int base);
extern int32_t C_ZNK10QStringRef6toLongEPbi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  QByteArray QStringRef::toUtf8();
extern void* C_ZNK10QStringRef6toUtf8Ev(void* qthis); // 4
  // proto:  QStringRef QStringRef::left(int n);
extern void C_ZNK10QStringRef4leftEi(void* qthis, int32_t arg0); // 4
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

// class sizeof(QStringDataPtr)=8
type QStringDataPtr struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QString)=8
type QString struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QLatin1String)=16
type QLatin1String struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QCharRef)=16
type QCharRef struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStringRef)=16
type QStringRef struct {
  // qbase: None;
  Qclsinst unsafe.Pointer /* *C.void */;
}

// toUcs4()
func (this *QString) Toucs4(args ...interface{}) () {
  // toUcs4()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6toUcs4Ev
    // invoke: QVector<uint> toUcs4()
    C.C_ZNK7QString6toUcs4Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toUcs4", args)
  }

  return
}

// toWCharArray(wchar_t *)
func (this *QString) Towchararray(args ...interface{}) (ret interface{}) {
  // toWCharArray(wchar_t *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.RuneTy(true) // "wchar_t *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12toWCharArrayEPw
    // invoke: int toWCharArray(wchar_t *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZNK7QString12toWCharArrayEPw(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toWCharArray", args)
  }

  return
}

// unicode()
func (this *QString) Unicode(args ...interface{}) (ret interface{}) {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7unicodeEv
    // invoke: const QChar * unicode()
    var ret0 = C.C_ZNK7QString7unicodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "unicode", args)
  }

  return
}

// fromLatin1(const class QByteArray &)
func (this *QString) Fromlatin1_S(args ...interface{}) (ret interface{}) {
  // fromLatin1(const class QByteArray &)
  // fromLatin1(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString10fromLatin1ERK10QByteArray
    // invoke: QString fromLatin1(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString10fromLatin1ERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString10fromLatin1EPKci
    // invoke: QString fromLatin1(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString10fromLatin1EPKci(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromLatin1", args)
  }

  return
}

// reserve(int)
func (this *QString) Reserve(args ...interface{}) () {
  // reserve(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7reserveEi
    // invoke: void reserve(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QString7reserveEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "reserve", args)
  }

  return
}

// swap(class QString &)
func (this *QString) Swap(args ...interface{}) () {
  // swap(class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4swapERS_
    // invoke: void swap(class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QString4swapERS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "swap", args)
  }

  return
}

// rightRef(int)
func (this *QString) Rightref(args ...interface{}) () {
  // rightRef(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8rightRefEi
    // invoke: QStringRef rightRef(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK7QString8rightRefEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "rightRef", args)
  }

  return
}

// leftRef(int)
func (this *QString) Leftref(args ...interface{}) () {
  // leftRef(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7leftRefEi
    // invoke: QStringRef leftRef(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK7QString7leftRefEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "leftRef", args)
  }

  return
}

// rightJustified(int, class QChar, _Bool)
func (this *QString) Rightjustified(args ...interface{}) (ret interface{}) {
  // rightJustified(int, class QChar, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString14rightJustifiedEi5QCharb
    // invoke: QString rightJustified(int, class QChar, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString14rightJustifiedEi5QCharb(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "rightJustified", args)
  }

  return
}

// detach()
func (this *QString) Detach(args ...interface{}) () {
  // detach()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6detachEv
    // invoke: void detach()
    C.C_ZN7QString6detachEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "detach", args)
  }

  return
}

// insert(int, class QChar)
func (this *QString) Insert(args ...interface{}) (ret interface{}) {
  // insert(int, class QChar)
  // insert(int, const class QString &)
  // insert(int, const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6insertEi5QChar
    // invoke: QString & insert(int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6insertEi5QChar(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString6insertEiRKS_
    // invoke: QString & insert(int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6insertEiRKS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString6insertEiPK5QChari
    // invoke: QString & insert(int, const class QChar *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QString6insertEiPK5QChari(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "insert", args)
  }

  return
}

// midRef(int, int)
func (this *QString) Midref(args ...interface{}) () {
  // midRef(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6midRefEii
    // invoke: QStringRef midRef(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK7QString6midRefEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "midRef", args)
  }

  return
}

// remove(const class QRegExp &)
func (this *QString) Remove(args ...interface{}) (ret interface{}) {
  // remove(const class QRegExp &)
  // remove(const class QRegularExpression &)
  // remove(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6removeERK7QRegExp
    // invoke: QString & remove(const class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString6removeERK7QRegExp(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString6removeERK18QRegularExpression
    // invoke: QString & remove(const class QRegularExpression &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString6removeERK18QRegularExpression(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString6removeEii
    // invoke: QString & remove(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6removeEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "remove", args)
  }

  return
}

// trimmed()
func (this *QString) Trimmed(args ...interface{}) (ret interface{}) {
  // trimmed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString7trimmedEv
    // invoke: QString trimmed()
    var ret0 = C.C_ZNKR7QString7trimmedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "trimmed", args)
  }

  return
}

// toStdWString()
func (this *QString) Tostdwstring(args ...interface{}) () {
  // toStdWString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12toStdWStringB5cxx11Ev
    // invoke: std::wstring toStdWString()
    C.C_ZNK7QString12toStdWStringB5cxx11Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdWString", args)
  }

  return
}

// constEnd()
func (this *QString) Constend(args ...interface{}) (ret interface{}) {
  // constEnd()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8constEndEv
    // invoke: const_iterator constEnd()
    var ret0 = C.C_ZNK7QString8constEndEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "constEnd", args)
  }

  return
}

// utf16()
func (this *QString) Utf16(args ...interface{}) () {
  // utf16()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5utf16Ev
    // invoke: const ushort * utf16()
    C.C_ZNK7QString5utf16Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "utf16", args)
  }

  return
}

// right(int)
func (this *QString) Right(args ...interface{}) (ret interface{}) {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5rightEi
    // invoke: QString right(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString5rightEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "right", args)
  }

  return
}

// toLocal8Bit()
func (this *QString) Tolocal8Bit(args ...interface{}) (ret interface{}) {
  // toLocal8Bit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString11toLocal8BitEv
    // invoke: QByteArray toLocal8Bit()
    var ret0 = C.C_ZNKR7QString11toLocal8BitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toLocal8Bit", args)
  }

  return
}

// toUInt(_Bool *, int)
func (this *QString) Touint(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString6toUIntEPbi
    // invoke: uint toUInt(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString6toUIntEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toUInt", args)
  }

  return
}

// QString()
func NewQString(args ...interface{}) *QString {
  // QString()
  // QString(const class QString &)
  // QString(int, class QChar)
  // QString(const char *)
  // QString(const class QChar *, int)
  // QString(class QChar)
  // QString(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.ByteTy(true) // "const char *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QStringC1Ev
    // invoke: void QString()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2Ev()
    return &QString{Qclsinst:qthis}
  case 1:
    // invoke: _ZN7QStringC1ERKS_
    // invoke: void QString(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2ERKS_(arg0)
    return &QString{Qclsinst:qthis}
  case 2:
    // invoke: _ZN7QStringC1Ei5QChar
    // invoke: void QString(int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2Ei5QChar(arg0, arg1)
    return &QString{Qclsinst:qthis}
  case 3:
    // invoke: _ZN7QStringC1EPKc
    // invoke: void QString(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[3][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2EPKc(arg0)
    return &QString{Qclsinst:qthis}
  case 4:
    // invoke: _ZN7QStringC1EPK5QChari
    // invoke: void QString(const class QChar *, int)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2EPK5QChari(arg0, arg1)
    return &QString{Qclsinst:qthis}
  case 5:
    // invoke: _ZN7QStringC1E5QChar
    // invoke: void QString(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2E5QChar(arg0)
    return &QString{Qclsinst:qthis}
  case 6:
    // invoke: _ZN7QStringC1ERK10QByteArray
    // invoke: void QString(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN7QStringC2ERK10QByteArray(arg0)
    return &QString{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QString", "QString", args)
  }

  return nil // QString{Qclsinst:qthis}
}

// arg(const class QString &, int, class QChar)
func (this *QString) Arg(args ...interface{}) (ret interface{}) {
  // arg(const class QString &, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &)
  // arg(qlonglong, int, int, class QChar)
  // arg(long, int, int, class QChar)
  // arg(ushort, int, int, class QChar)
  // arg(int, int, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(short, int, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(uint, int, int, class QChar)
  // arg(const class QString &, const class QString &)
  // arg(double, int, char, int, class QChar)
  // arg(qulonglong, int, int, class QChar)
  // arg(char, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(ulong, int, int, class QChar)
  // arg(class QChar, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "int"
  vtys[3][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "long"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = qtrt.Int32Ty(false) // "int"
  vtys[4][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "int"
  vtys[5][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "int"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = qtrt.Int32Ty(false) // "int"
  vtys[6][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][7] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int16Ty(false) // "short"
  vtys[8][1] = qtrt.Int32Ty(false) // "int"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[8][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "uint"
  vtys[11][1] = qtrt.Int32Ty(false) // "int"
  vtys[11][2] = qtrt.Int32Ty(false) // "int"
  vtys[11][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[12][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.DoubleTy(false) // "double"
  vtys[13][1] = qtrt.Int32Ty(false) // "int"
  vtys[13][2] = qtrt.ByteTy(false) // "char"
  vtys[13][3] = qtrt.Int32Ty(false) // "int"
  vtys[13][4] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[14][1] = qtrt.Int32Ty(false) // "int"
  vtys[14][2] = qtrt.Int32Ty(false) // "int"
  vtys[14][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = qtrt.ByteTy(false) // "char"
  vtys[15][1] = qtrt.Int32Ty(false) // "int"
  vtys[15][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][7] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[16][8] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[17][1] = qtrt.Int32Ty(false) // "int"
  vtys[17][2] = qtrt.Int32Ty(false) // "int"
  vtys[17][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[18][1] = qtrt.Int32Ty(false) // "int"
  vtys[18][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[19][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[19][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[19][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[19][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[19][5] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString3argERKS_i5QChar
    // invoke: QString arg(const class QString &, int, class QChar)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QChar).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString3argERKS_i5QChar(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_S1_(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK7QString3argERKS_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK7QString3argExii5QChar
    // invoke: QString arg(qlonglong, int, int, class QChar)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argExii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZNK7QString3argElii5QChar
    // invoke: QString arg(long, int, int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argElii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZNK7QString3argEtii5QChar
    // invoke: QString arg(ushort, int, int, class QChar)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argEtii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 6:
    // invoke: _ZNK7QString3argEiii5QChar
    // invoke: QString arg(int, int, int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argEiii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 7:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(*QString).Qclsinst
    if false {fmt.Println(arg6)}
    var arg7 = args[7].(*QString).Qclsinst
    if false {fmt.Println(arg7)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 8:
    // invoke: _ZNK7QString3argEsii5QChar
    // invoke: QString arg(short, int, int, class QChar)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argEsii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 9:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(*QString).Qclsinst
    if false {fmt.Println(arg6)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 10:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QString).Qclsinst
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_S1_S1_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 11:
    // invoke: _ZNK7QString3argEjii5QChar
    // invoke: QString arg(uint, int, int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argEjii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 12:
    // invoke: _ZNK7QString3argERKS_S1_
    // invoke: QString arg(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 13:
    // invoke: _ZNK7QString3argEdici5QChar
    // invoke: QString arg(double, int, char, int, class QChar)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.uchar(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QChar).Qclsinst
    if false {fmt.Println(arg4)}
    var ret0 = C.C_ZNK7QString3argEdici5QChar(this.Qclsinst, arg0, arg1, arg2, arg3, arg4)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 14:
    // invoke: _ZNK7QString3argEyii5QChar
    // invoke: QString arg(qulonglong, int, int, class QChar)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argEyii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 15:
    // invoke: _ZNK7QString3argEci5QChar
    // invoke: QString arg(char, int, class QChar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QChar).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString3argEci5QChar(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 16:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*QString).Qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(*QString).Qclsinst
    if false {fmt.Println(arg6)}
    var arg7 = args[7].(*QString).Qclsinst
    if false {fmt.Println(arg7)}
    var arg8 = args[8].(*QString).Qclsinst
    if false {fmt.Println(arg8)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 17:
    // invoke: _ZNK7QString3argEmii5QChar
    // invoke: QString arg(ulong, int, int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QChar).Qclsinst
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZNK7QString3argEmii5QChar(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 18:
    // invoke: _ZNK7QString3argE5QChariS0_
    // invoke: QString arg(class QChar, int, class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QChar).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString3argE5QChariS0_(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 19:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(*QString).Qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(*QString).Qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(*QString).Qclsinst
    if false {fmt.Println(arg5)}
    var ret0 = C.C_ZNK7QString3argERKS_S1_S1_S1_S1_S1_(this.Qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "arg", args)
  }

  return
}

// append(class QChar)
func (this *QString) Append(args ...interface{}) (ret interface{}) {
  // append(class QChar)
  // append(const char *)
  // append(const class QByteArray &)
  // append(const class QString &)
  // append(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6appendE5QChar
    // invoke: QString & append(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString6appendE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString6appendEPKc
    // invoke: QString & append(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN7QString6appendEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString6appendERK10QByteArray
    // invoke: QString & append(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString6appendERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN7QString6appendERKS_
    // invoke: QString & append(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString6appendERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN7QString6appendEPK5QChari
    // invoke: QString & append(const class QChar *, int)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6appendEPK5QChari(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "append", args)
  }

  return
}

// capacity()
func (this *QString) Capacity(args ...interface{}) (ret interface{}) {
  // capacity()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8capacityEv
    // invoke: int capacity()
    var ret0 = C.C_ZNK7QString8capacityEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "capacity", args)
  }

  return
}

// toLower()
func (this *QString) Tolower(args ...interface{}) (ret interface{}) {
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString7toLowerEv
    // invoke: QString toLower()
    var ret0 = C.C_ZNKR7QString7toLowerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toLower", args)
  }

  return
}

// toHtmlEscaped()
func (this *QString) Tohtmlescaped(args ...interface{}) (ret interface{}) {
  // toHtmlEscaped()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString13toHtmlEscapedEv
    // invoke: QString toHtmlEscaped()
    var ret0 = C.C_ZNK7QString13toHtmlEscapedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toHtmlEscaped", args)
  }

  return
}

// isNull()
func (this *QString) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK7QString6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "isNull", args)
  }

  return
}

// localeAwareCompare(const class QString &, const class QString &)
func (this *QString) Localeawarecompare_S(args ...interface{}) (ret interface{}) {
  // localeAwareCompare(const class QString &, const class QString &)
  // localeAwareCompare(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString18localeAwareCompareERKS_S1_
    // invoke: int localeAwareCompare(const class QString &, const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString18localeAwareCompareERKS_S1_(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QString18localeAwareCompareERKS_
    // invoke: int localeAwareCompare(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString18localeAwareCompareERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "localeAwareCompare", args)
  }

  return
}

// fromLocal8Bit(const class QByteArray &)
func (this *QString) Fromlocal8Bit_S(args ...interface{}) (ret interface{}) {
  // fromLocal8Bit(const class QByteArray &)
  // fromLocal8Bit(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString13fromLocal8BitERK10QByteArray
    // invoke: QString fromLocal8Bit(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString13fromLocal8BitERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString13fromLocal8BitEPKci
    // invoke: QString fromLocal8Bit(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString13fromLocal8BitEPKci(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromLocal8Bit", args)
  }

  return
}

// fromWCharArray(const wchar_t *, int)
func (this *QString) Fromwchararray_S(args ...interface{}) (ret interface{}) {
  // fromWCharArray(const wchar_t *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.RuneTy(true) // "const wchar_t *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString14fromWCharArrayEPKwi
    // invoke: QString fromWCharArray(const wchar_t *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString14fromWCharArrayEPKwi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromWCharArray", args)
  }

  return
}

// toStdString()
func (this *QString) Tostdstring(args ...interface{}) () {
  // toStdString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString11toStdStringB5cxx11Ev
    // invoke: std::string toStdString()
    C.C_ZNK7QString11toStdStringB5cxx11Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdString", args)
  }

  return
}

// indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
func (this *QString) Indexof(args ...interface{}) (ret interface{}) {
  // indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
  // indexOf(const class QRegularExpression &, int)
  // indexOf(class QRegExp &, int)
  // indexOf(const class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch
    // invoke: int indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QRegularExpressionMatch).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioni
    // invoke: int indexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString7indexOfERK18QRegularExpressioni(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK7QString7indexOfER7QRegExpi
    // invoke: int indexOf(class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString7indexOfER7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK7QString7indexOfERK7QRegExpi
    // invoke: int indexOf(const class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString7indexOfERK7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "indexOf", args)
  }

  return
}

// cbegin()
func (this *QString) Cbegin(args ...interface{}) (ret interface{}) {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6cbeginEv
    // invoke: const_iterator cbegin()
    var ret0 = C.C_ZNK7QString6cbeginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "cbegin", args)
  }

  return
}

// toLatin1()
func (this *QString) Tolatin1(args ...interface{}) (ret interface{}) {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString8toLatin1Ev
    // invoke: QByteArray toLatin1()
    var ret0 = C.C_ZNKR7QString8toLatin1Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toLatin1", args)
  }

  return
}

// chop(int)
func (this *QString) Chop(args ...interface{}) () {
  // chop(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4chopEi
    // invoke: void chop(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QString4chopEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "chop", args)
  }

  return
}

// toStdU32String()
func (this *QString) Tostdu32String(args ...interface{}) () {
  // toStdU32String()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString14toStdU32StringB5cxx11Ev
    // invoke: std::u32string toStdU32String()
    C.C_ZNK7QString14toStdU32StringB5cxx11Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdU32String", args)
  }

  return
}

// length()
func (this *QString) Length(args ...interface{}) (ret interface{}) {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString6lengthEv
    // invoke: int length()
    var ret0 = C.C_ZNK7QString6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "length", args)
  }

  return
}

// toDouble(_Bool *)
func (this *QString) Todouble(args ...interface{}) (ret interface{}) {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8toDoubleEPb
    // invoke: double toDouble(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString8toDoubleEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toDouble", args)
  }

  return
}

// isSharedWith(const class QString &)
func (this *QString) Issharedwith(args ...interface{}) (ret interface{}) {
  // isSharedWith(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12isSharedWithERKS_
    // invoke: bool isSharedWith(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString12isSharedWithERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "isSharedWith", args)
  }

  return
}

// toUShort(_Bool *, int)
func (this *QString) Toushort(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString8toUShortEPbi
    // invoke: ushort toUShort(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString8toUShortEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toUShort", args)
  }

  return
}

// constBegin()
func (this *QString) Constbegin(args ...interface{}) (ret interface{}) {
  // constBegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString10constBeginEv
    // invoke: const_iterator constBegin()
    var ret0 = C.C_ZNK7QString10constBeginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "constBegin", args)
  }

  return
}

// toShort(_Bool *, int)
func (this *QString) Toshort(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString7toShortEPbi
    // invoke: short toShort(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString7toShortEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "short"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toShort", args)
  }

  return
}

// number(uint, int)
func (this *QString) Number_S(args ...interface{}) (ret interface{}) {
  // number(uint, int)
  // number(double, char, int)
  // number(qulonglong, int)
  // number(ulong, int)
  // number(qlonglong, int)
  // number(long, int)
  // number(int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "uint"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.DoubleTy(false) // "double"
  vtys[1][1] = qtrt.ByteTy(false) // "char"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "long"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "int"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6numberEji
    // invoke: QString number(uint, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6numberEji(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString6numberEdci
    // invoke: QString number(double, char, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QString6numberEdci(arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString6numberEyi
    // invoke: QString number(qulonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6numberEyi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN7QString6numberEmi
    // invoke: QString number(ulong, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6numberEmi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN7QString6numberExi
    // invoke: QString number(qlonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6numberExi(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZN7QString6numberEli
    // invoke: QString number(long, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6numberEli(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 6:
    // invoke: _ZN7QString6numberEii
    // invoke: QString number(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6numberEii(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "number", args)
  }

  return
}

// replace(const class QRegExp &, const class QString &)
func (this *QString) Replace(args ...interface{}) (ret interface{}) {
  // replace(const class QRegExp &, const class QString &)
  // replace(int, int, class QChar)
  // replace(const class QRegularExpression &, const class QString &)
  // replace(int, int, const class QChar *, int)
  // replace(int, int, const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[3][3] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "int"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[4][2] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7replaceERK7QRegExpRKS_
    // invoke: QString & replace(const class QRegExp &, const class QString &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString7replaceERK7QRegExpRKS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString7replaceEii5QChar
    // invoke: QString & replace(int, int, class QChar)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QChar).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QString7replaceEii5QChar(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString7replaceERK18QRegularExpressionRKS_
    // invoke: QString & replace(const class QRegularExpression &, const class QString &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QString).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString7replaceERK18QRegularExpressionRKS_(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN7QString7replaceEiiPK5QChari
    // invoke: QString & replace(int, int, const class QChar *, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QChar).Qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(qtrt.PrimConv(args[3], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg3)}
    var ret0 = C.C_ZN7QString7replaceEiiPK5QChari(this.Qclsinst, arg0, arg1, arg2, arg3)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN7QString7replaceEiiRKS_
    // invoke: QString & replace(int, int, const class QString &)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QString).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QString7replaceEiiRKS_(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "replace", args)
  }

  return
}

// push_back(class QChar)
func (this *QString) Push_Back(args ...interface{}) () {
  // push_back(class QChar)
  // push_back(const class QString &)
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
    // invoke: _ZN7QString9push_backE5QChar
    // invoke: void push_back(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QString9push_backE5QChar(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString9push_backERKS_
    // invoke: void push_back(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QString9push_backERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "push_back", args)
  }

  return
}

// simplified()
func (this *QString) Simplified(args ...interface{}) (ret interface{}) {
  // simplified()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString10simplifiedEv
    // invoke: QString simplified()
    var ret0 = C.C_ZNKR7QString10simplifiedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "simplified", args)
  }

  return
}

// size()
func (this *QString) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK7QString4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "size", args)
  }

  return
}

// truncate(int)
func (this *QString) Truncate(args ...interface{}) () {
  // truncate(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString8truncateEi
    // invoke: void truncate(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QString8truncateEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "truncate", args)
  }

  return
}

// contains(const class QRegularExpression &)
func (this *QString) Contains(args ...interface{}) (ret interface{}) {
  // contains(const class QRegularExpression &)
  // contains(const class QRegularExpression &, class QRegularExpressionMatch *)
  // contains(class QRegExp &)
  // contains(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1][1] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8containsERK18QRegularExpression
    // invoke: bool contains(const class QRegularExpression &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString8containsERK18QRegularExpression(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch
    // invoke: bool contains(const class QRegularExpression &, class QRegularExpressionMatch *)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QRegularExpressionMatch).Qclsinst
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK7QString8containsER7QRegExp
    // invoke: bool contains(class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString8containsER7QRegExp(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK7QString8containsERK7QRegExp
    // invoke: bool contains(const class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString8containsERK7QRegExp(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "contains", args)
  }

  return
}

// isSimpleText()
func (this *QString) Issimpletext(args ...interface{}) (ret interface{}) {
  // isSimpleText()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12isSimpleTextEv
    // invoke: bool isSimpleText()
    var ret0 = C.C_ZNK7QString12isSimpleTextEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "isSimpleText", args)
  }

  return
}

// cend()
func (this *QString) Cend(args ...interface{}) (ret interface{}) {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString4cendEv
    // invoke: const_iterator cend()
    var ret0 = C.C_ZNK7QString4cendEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const_iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "cend", args)
  }

  return
}

// leftJustified(int, class QChar, _Bool)
func (this *QString) Leftjustified(args ...interface{}) (ret interface{}) {
  // leftJustified(int, class QChar, _Bool)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][2] = qtrt.BoolTy(false) // "bool"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString13leftJustifiedEi5QCharb
    // invoke: QString leftJustified(int, class QChar, _Bool)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(*QChar).Qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString13leftJustifiedEi5QCharb(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "leftJustified", args)
  }

  return
}

// repeated(int)
func (this *QString) Repeated(args ...interface{}) (ret interface{}) {
  // repeated(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8repeatedEi
    // invoke: QString repeated(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString8repeatedEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "repeated", args)
  }

  return
}

// fromRawData(const class QChar *, int)
func (this *QString) Fromrawdata_S(args ...interface{}) (ret interface{}) {
  // fromRawData(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString11fromRawDataEPK5QChari
    // invoke: QString fromRawData(const class QChar *, int)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString11fromRawDataEPK5QChari(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromRawData", args)
  }

  return
}

// squeeze()
func (this *QString) Squeeze(args ...interface{}) () {
  // squeeze()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7squeezeEv
    // invoke: void squeeze()
    C.C_ZN7QString7squeezeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "squeeze", args)
  }

  return
}

// resize(int)
func (this *QString) Resize(args ...interface{}) () {
  // resize(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6resizeEi
    // invoke: void resize(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZN7QString6resizeEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "resize", args)
  }

  return
}

// toULongLong(_Bool *, int)
func (this *QString) Toulonglong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString11toULongLongEPbi
    // invoke: qulonglong toULongLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString11toULongLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qulonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toULongLong", args)
  }

  return
}

// count(const class QRegularExpression &)
func (this *QString) Count(args ...interface{}) (ret interface{}) {
  // count(const class QRegularExpression &)
  // count()
  // count(const class QRegExp &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString5countERK18QRegularExpression
    // invoke: int count(const class QRegularExpression &)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString5countERK18QRegularExpression(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QString5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK7QString5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK7QString5countERK7QRegExp
    // invoke: int count(const class QRegExp &)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString5countERK7QRegExp(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "count", args)
  }

  return
}

// isRightToLeft()
func (this *QString) Isrighttoleft(args ...interface{}) (ret interface{}) {
  // isRightToLeft()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString13isRightToLeftEv
    // invoke: bool isRightToLeft()
    var ret0 = C.C_ZNK7QString13isRightToLeftEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "isRightToLeft", args)
  }

  return
}

// toUtf8()
func (this *QString) Toutf8(args ...interface{}) (ret interface{}) {
  // toUtf8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString6toUtf8Ev
    // invoke: QByteArray toUtf8()
    var ret0 = C.C_ZNKR7QString6toUtf8Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toUtf8", args)
  }

  return
}

// fromUtf8(const char *, int)
func (this *QString) Fromutf8_S(args ...interface{}) (ret interface{}) {
  // fromUtf8(const char *, int)
  // fromUtf8(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(true) // "const char *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString8fromUtf8EPKci
    // invoke: QString fromUtf8(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[0][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString8fromUtf8EPKci(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString8fromUtf8ERK10QByteArray
    // invoke: QString fromUtf8(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString8fromUtf8ERK10QByteArray(arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromUtf8", args)
  }

  return
}

// toStdU16String()
func (this *QString) Tostdu16String(args ...interface{}) () {
  // toStdU16String()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString14toStdU16StringB5cxx11Ev
    // invoke: std::u16string toStdU16String()
    C.C_ZNK7QString14toStdU16StringB5cxx11Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdU16String", args)
  }

  return
}

// toFloat(_Bool *)
func (this *QString) Tofloat(args ...interface{}) (ret interface{}) {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7toFloatEPb
    // invoke: float toFloat(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString7toFloatEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toFloat", args)
  }

  return
}

// at(int)
func (this *QString) At(args ...interface{}) (ret interface{}) {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString2atEi
    // invoke: const QChar at(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString2atEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "at", args)
  }

  return
}

// fill(class QChar, int)
func (this *QString) Fill(args ...interface{}) (ret interface{}) {
  // fill(class QChar, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4fillE5QChari
    // invoke: QString & fill(class QChar, int)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString4fillE5QChari(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fill", args)
  }

  return
}

// toULong(_Bool *, int)
func (this *QString) Toulong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString7toULongEPbi
    // invoke: ulong toULong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString7toULongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "ulong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toULong", args)
  }

  return
}

// end()
func (this *QString) End(args ...interface{}) (ret interface{}) {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString3endEv
    // invoke: iterator end()
    var ret0 = C.C_ZN7QString3endEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "end", args)
  }

  return
}

// fromUcs4(const uint *, int)
func (this *QString) Fromucs4_S(args ...interface{}) (ret interface{}) {
  // fromUcs4(const uint *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(true) // "const uint *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString8fromUcs4EPKji
    // invoke: QString fromUcs4(const uint *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString8fromUcs4EPKji(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromUcs4", args)
  }

  return
}

// mid(int, int)
func (this *QString) Mid(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString3midEii
    // invoke: QString mid(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString3midEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "mid", args)
  }

  return
}

// prepend(const class QString &)
func (this *QString) Prepend(args ...interface{}) (ret interface{}) {
  // prepend(const class QString &)
  // prepend(class QChar)
  // prepend(const char *)
  // prepend(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7prependERKS_
    // invoke: QString & prepend(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString7prependERKS_(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString7prependE5QChar
    // invoke: QString & prepend(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString7prependE5QChar(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString7prependEPKc
    // invoke: QString & prepend(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var ret0 = C.C_ZN7QString7prependEPKc(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN7QString7prependERK10QByteArray
    // invoke: QString & prepend(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZN7QString7prependERK10QByteArray(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "prepend", args)
  }

  return
}

// constData()
func (this *QString) Constdata(args ...interface{}) (ret interface{}) {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString9constDataEv
    // invoke: const QChar * constData()
    var ret0 = C.C_ZNK7QString9constDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "constData", args)
  }

  return
}

// toLongLong(_Bool *, int)
func (this *QString) Tolonglong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString10toLongLongEPbi
    // invoke: qlonglong toLongLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString10toLongLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qlonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toLongLong", args)
  }

  return
}

// setUnicode(const class QChar *, int)
func (this *QString) Setunicode(args ...interface{}) (ret interface{}) {
  // setUnicode(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString10setUnicodeEPK5QChari
    // invoke: QString & setUnicode(const class QChar *, int)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString10setUnicodeEPK5QChari(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "setUnicode", args)
  }

  return
}

// fromUtf16(const ushort *, int)
func (this *QString) Fromutf16_S(args ...interface{}) (ret interface{}) {
  // fromUtf16(const ushort *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(true) // "const ushort *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString9fromUtf16EPKti
    // invoke: QString fromUtf16(const ushort *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString9fromUtf16EPKti(arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "fromUtf16", args)
  }

  return
}

// isEmpty()
func (this *QString) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK7QString7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "isEmpty", args)
  }

  return
}

// setRawData(const class QChar *, int)
func (this *QString) Setrawdata(args ...interface{}) (ret interface{}) {
  // setRawData(const class QChar *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString10setRawDataEPK5QChari
    // invoke: QString & setRawData(const class QChar *, int)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString10setRawDataEPK5QChari(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "setRawData", args)
  }

  return
}

// isDetached()
func (this *QString) Isdetached(args ...interface{}) (ret interface{}) {
  // isDetached()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString10isDetachedEv
    // invoke: bool isDetached()
    var ret0 = C.C_ZNK7QString10isDetachedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "isDetached", args)
  }

  return
}

// begin()
func (this *QString) Begin(args ...interface{}) (ret interface{}) {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString5beginEv
    // invoke: iterator begin()
    var ret0 = C.C_ZN7QString5beginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "iterator"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "begin", args)
  }

  return
}

// lastIndexOf(class QRegExp &, int)
func (this *QString) Lastindexof(args ...interface{}) (ret interface{}) {
  // lastIndexOf(class QRegExp &, int)
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString11lastIndexOfER7QRegExpi
    // invoke: int lastIndexOf(class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString11lastIndexOfER7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZNK7QString11lastIndexOfERK7QRegExpi
    // invoke: int lastIndexOf(const class QRegExp &, int)
    var arg0 = args[0].(*QRegExp).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString11lastIndexOfERK7QRegExpi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioni
    // invoke: int lastIndexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString11lastIndexOfERK18QRegularExpressioni(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch
    // invoke: int lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
    var arg0 = args[0].(*QRegularExpression).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(*QRegularExpressionMatch).Qclsinst
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "lastIndexOf", args)
  }

  return
}

// setNum(ushort, int)
func (this *QString) Setnum(args ...interface{}) (ret interface{}) {
  // setNum(ushort, int)
  // setNum(qulonglong, int)
  // setNum(qlonglong, int)
  // setNum(ulong, int)
  // setNum(short, int)
  // setNum(long, int)
  // setNum(uint, int)
  // setNum(int, int)
  // setNum(double, char, int)
  // setNum(float, char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int16Ty(false) // "short"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "long"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "uint"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.DoubleTy(false) // "double"
  vtys[8][1] = qtrt.ByteTy(false) // "char"
  vtys[8][2] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.FloatTy(false) // "float"
  vtys[9][1] = qtrt.ByteTy(false) // "char"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6setNumEti
    // invoke: QString & setNum(ushort, int)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEti(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 1:
    // invoke: _ZN7QString6setNumEyi
    // invoke: QString & setNum(qulonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEyi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 2:
    // invoke: _ZN7QString6setNumExi
    // invoke: QString & setNum(qlonglong, int)
    var arg0 = C.int64_t(qtrt.PrimConv(args[0], qtrt.Int64Ty(false)).(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumExi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 3:
    // invoke: _ZN7QString6setNumEmi
    // invoke: QString & setNum(ulong, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEmi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 4:
    // invoke: _ZN7QString6setNumEsi
    // invoke: QString & setNum(short, int)
    var arg0 = C.int16_t(qtrt.PrimConv(args[0], qtrt.Int16Ty(false)).(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEsi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 5:
    // invoke: _ZN7QString6setNumEli
    // invoke: QString & setNum(long, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEli(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 6:
    // invoke: _ZN7QString6setNumEji
    // invoke: QString & setNum(uint, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEji(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 7:
    // invoke: _ZN7QString6setNumEii
    // invoke: QString & setNum(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString6setNumEii(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 8:
    // invoke: _ZN7QString6setNumEdci
    // invoke: QString & setNum(double, char, int)
    var arg0 = C.double(qtrt.PrimConv(args[0], qtrt.DoubleTy(false)).(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QString6setNumEdci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  case 9:
    // invoke: _ZN7QString6setNumEfci
    // invoke: QString & setNum(float, char, int)
    var arg0 = C.float(qtrt.PrimConv(args[0], qtrt.FloatTy(false)).(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var ret0 = C.C_ZN7QString6setNumEfci(this.Qclsinst, arg0, arg1, arg2)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "setNum", args)
  }

  return
}

// setUtf16(const ushort *, int)
func (this *QString) Setutf16(args ...interface{}) (ret interface{}) {
  // setUtf16(const ushort *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(true) // "const ushort *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString8setUtf16EPKti
    // invoke: QString & setUtf16(const ushort *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZN7QString8setUtf16EPKti(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "setUtf16", args)
  }

  return
}

// data()
func (this *QString) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString4dataEv
    // invoke: QChar * data()
    var ret0 = C.C_ZN7QString4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "data", args)
  }

  return
}

// toLong(_Bool *, int)
func (this *QString) Tolong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString6toLongEPbi
    // invoke: long toLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString6toLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "long"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toLong", args)
  }

  return
}

// ~QString()
func (this *QString) Freeqstring(args ...interface{}) () {
  // ~QString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QStringD0Ev
    // invoke: void ~QString()
    C.C_ZN7QStringD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "~QString", args)
  }

  return
}

// clear()
func (this *QString) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString5clearEv
    // invoke: void clear()
    C.C_ZN7QString5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QString", "clear", args)
  }

  return
}

// toInt(_Bool *, int)
func (this *QString) Toint(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK7QString5toIntEPbi
    // invoke: int toInt(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK7QString5toIntEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toInt", args)
  }

  return
}

// toCaseFolded()
func (this *QString) Tocasefolded(args ...interface{}) (ret interface{}) {
  // toCaseFolded()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString12toCaseFoldedEv
    // invoke: QString toCaseFolded()
    var ret0 = C.C_ZNKR7QString12toCaseFoldedEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toCaseFolded", args)
  }

  return
}

// toUpper()
func (this *QString) Toupper(args ...interface{}) (ret interface{}) {
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNKR7QString7toUpperEv
    // invoke: QString toUpper()
    var ret0 = C.C_ZNKR7QString7toUpperEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "toUpper", args)
  }

  return
}

// push_front(class QChar)
func (this *QString) Push_Front(args ...interface{}) () {
  // push_front(class QChar)
  // push_front(const class QString &)
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
    // invoke: _ZN7QString10push_frontE5QChar
    // invoke: void push_front(class QChar)
    var arg0 = args[0].(*QChar).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QString10push_frontE5QChar(this.Qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString10push_frontERKS_
    // invoke: void push_front(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZN7QString10push_frontERKS_(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "push_front", args)
  }

  return
}

// left(int)
func (this *QString) Left(args ...interface{}) (ret interface{}) {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString4leftEi
    // invoke: QString left(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK7QString4leftEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QString", "left", args)
  }

  return
}

// QLatin1String(const class QByteArray &)
func NewQLatin1String(args ...interface{}) *QLatin1String {
  // QLatin1String(const class QByteArray &)
  // QLatin1String(const char *)
  // QLatin1String(const char *, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN13QLatin1StringC1ERK10QByteArray
    // invoke: void QLatin1String(const class QByteArray &)
    var arg0 = args[0].(*QByteArray).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QLatin1StringC2ERK10QByteArray(arg0)
    return &QLatin1String{Qclsinst:qthis}
  case 1:
    // invoke: _ZN13QLatin1StringC1EPKc
    // invoke: void QLatin1String(const char *)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[1][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QLatin1StringC2EPKc(arg0)
    return &QLatin1String{Qclsinst:qthis}
  case 2:
    // invoke: _ZN13QLatin1StringC1EPKci
    // invoke: void QLatin1String(const char *, int)
    argif0, free0 := qtrt.HandyConvert2c(args[0], vtys[2][0])
    var arg0 = argif0.(unsafe.Pointer)
    if false {fmt.Println(argif0, arg0)}
    if free0 {defer C.free(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN13QLatin1StringC2EPKci(arg0, arg1)
    return &QLatin1String{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QLatin1String", "QLatin1String", args)
  }

  return nil // QLatin1String{Qclsinst:qthis}
}

// latin1()
func (this *QLatin1String) Latin1(args ...interface{}) (ret interface{}) {
  // latin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QLatin1String6latin1Ev
    // invoke: const char * latin1()
    var ret0 = C.C_ZNK13QLatin1String6latin1Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLatin1String", "latin1", args)
  }

  return
}

// data()
func (this *QLatin1String) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QLatin1String4dataEv
    // invoke: const char * data()
    var ret0 = C.C_ZNK13QLatin1String4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(true) // "const char *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLatin1String", "data", args)
  }

  return
}

// size()
func (this *QLatin1String) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK13QLatin1String4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK13QLatin1String4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QLatin1String", "size", args)
  }

  return
}

// isLetterOrNumber()
func (this *QCharRef) Isletterornumber(args ...interface{}) (ret interface{}) {
  // isLetterOrNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef16isLetterOrNumberEv
    // invoke: bool isLetterOrNumber()
    var ret0 = C.C_ZN8QCharRef16isLetterOrNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isLetterOrNumber", args)
  }

  return
}

// isLetter()
func (this *QCharRef) Isletter(args ...interface{}) (ret interface{}) {
  // isLetter()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8isLetterEv
    // invoke: bool isLetter()
    var ret0 = C.C_ZNK8QCharRef8isLetterEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isLetter", args)
  }

  return
}

// hasMirrored()
func (this *QCharRef) Hasmirrored(args ...interface{}) (ret interface{}) {
  // hasMirrored()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11hasMirroredEv
    // invoke: bool hasMirrored()
    var ret0 = C.C_ZNK8QCharRef11hasMirroredEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "hasMirrored", args)
  }

  return
}

// unicodeVersion()
func (this *QCharRef) Unicodeversion(args ...interface{}) () {
  // unicodeVersion()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef14unicodeVersionEv
    // invoke: QChar::UnicodeVersion unicodeVersion()
    C.C_ZNK8QCharRef14unicodeVersionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "unicodeVersion", args)
  }

  return
}

// isTitleCase()
func (this *QCharRef) Istitlecase(args ...interface{}) (ret interface{}) {
  // isTitleCase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11isTitleCaseEv
    // invoke: bool isTitleCase()
    var ret0 = C.C_ZNK8QCharRef11isTitleCaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isTitleCase", args)
  }

  return
}

// setCell(uchar)
func (this *QCharRef) Setcell(args ...interface{}) () {
  // setCell(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef7setCellEh
    // invoke: void setCell(uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN8QCharRef7setCellEh(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCharRef", "setCell", args)
  }

  return
}

// unicode()
func (this *QCharRef) Unicode(args ...interface{}) (ret interface{}) {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef7unicodeEv
    // invoke: ushort & unicode()
    var ret0 = C.C_ZN8QCharRef7unicodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "ushort &"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "unicode", args)
  }

  return
}

// isSpace()
func (this *QCharRef) Isspace(args ...interface{}) (ret interface{}) {
  // isSpace()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isSpaceEv
    // invoke: bool isSpace()
    var ret0 = C.C_ZNK8QCharRef7isSpaceEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isSpace", args)
  }

  return
}

// row()
func (this *QCharRef) Row(args ...interface{}) (ret interface{}) {
  // row()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef3rowEv
    // invoke: uchar row()
    var ret0 = C.C_ZNK8QCharRef3rowEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "row", args)
  }

  return
}

// category()
func (this *QCharRef) Category(args ...interface{}) () {
  // category()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8categoryEv
    // invoke: QChar::Category category()
    C.C_ZNK8QCharRef8categoryEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "category", args)
  }

  return
}

// isUpper()
func (this *QCharRef) Isupper(args ...interface{}) (ret interface{}) {
  // isUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isUpperEv
    // invoke: bool isUpper()
    var ret0 = C.C_ZNK8QCharRef7isUpperEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isUpper", args)
  }

  return
}

// toLower()
func (this *QCharRef) Tolower(args ...interface{}) (ret interface{}) {
  // toLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7toLowerEv
    // invoke: QChar toLower()
    var ret0 = C.C_ZNK8QCharRef7toLowerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "toLower", args)
  }

  return
}

// script()
func (this *QCharRef) Script(args ...interface{}) () {
  // script()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef6scriptEv
    // invoke: QChar::Script script()
    C.C_ZNK8QCharRef6scriptEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "script", args)
  }

  return
}

// joiningType()
func (this *QCharRef) Joiningtype(args ...interface{}) () {
  // joiningType()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11joiningTypeEv
    // invoke: QChar::JoiningType joiningType()
    C.C_ZNK8QCharRef11joiningTypeEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "joiningType", args)
  }

  return
}

// cell()
func (this *QCharRef) Cell(args ...interface{}) (ret interface{}) {
  // cell()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef4cellEv
    // invoke: uchar cell()
    var ret0 = C.C_ZNK8QCharRef4cellEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "cell", args)
  }

  return
}

// digitValue()
func (this *QCharRef) Digitvalue(args ...interface{}) (ret interface{}) {
  // digitValue()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef10digitValueEv
    // invoke: int digitValue()
    var ret0 = C.C_ZNK8QCharRef10digitValueEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "digitValue", args)
  }

  return
}

// decomposition()
func (this *QCharRef) Decomposition(args ...interface{}) (ret interface{}) {
  // decomposition()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef13decompositionEv
    // invoke: QString decomposition()
    var ret0 = C.C_ZNK8QCharRef13decompositionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "decomposition", args)
  }

  return
}

// direction()
func (this *QCharRef) Direction(args ...interface{}) () {
  // direction()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef9directionEv
    // invoke: QChar::Direction direction()
    C.C_ZNK8QCharRef9directionEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "direction", args)
  }

  return
}

// mirroredChar()
func (this *QCharRef) Mirroredchar(args ...interface{}) (ret interface{}) {
  // mirroredChar()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef12mirroredCharEv
    // invoke: QChar mirroredChar()
    var ret0 = C.C_ZNK8QCharRef12mirroredCharEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "mirroredChar", args)
  }

  return
}

// combiningClass()
func (this *QCharRef) Combiningclass(args ...interface{}) (ret interface{}) {
  // combiningClass()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef14combiningClassEv
    // invoke: uchar combiningClass()
    var ret0 = C.C_ZNK8QCharRef14combiningClassEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "uchar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "combiningClass", args)
  }

  return
}

// decompositionTag()
func (this *QCharRef) Decompositiontag(args ...interface{}) () {
  // decompositionTag()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef16decompositionTagEv
    // invoke: QChar::Decomposition decompositionTag()
    C.C_ZNK8QCharRef16decompositionTagEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "decompositionTag", args)
  }

  return
}

// isMark()
func (this *QCharRef) Ismark(args ...interface{}) (ret interface{}) {
  // isMark()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef6isMarkEv
    // invoke: bool isMark()
    var ret0 = C.C_ZNK8QCharRef6isMarkEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isMark", args)
  }

  return
}

// isLower()
func (this *QCharRef) Islower(args ...interface{}) (ret interface{}) {
  // isLower()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isLowerEv
    // invoke: bool isLower()
    var ret0 = C.C_ZNK8QCharRef7isLowerEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isLower", args)
  }

  return
}

// isNumber()
func (this *QCharRef) Isnumber(args ...interface{}) (ret interface{}) {
  // isNumber()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8isNumberEv
    // invoke: bool isNumber()
    var ret0 = C.C_ZNK8QCharRef8isNumberEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isNumber", args)
  }

  return
}

// toTitleCase()
func (this *QCharRef) Totitlecase(args ...interface{}) (ret interface{}) {
  // toTitleCase()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef11toTitleCaseEv
    // invoke: QChar toTitleCase()
    var ret0 = C.C_ZNK8QCharRef11toTitleCaseEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "toTitleCase", args)
  }

  return
}

// toLatin1()
func (this *QCharRef) Tolatin1(args ...interface{}) (ret interface{}) {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef8toLatin1Ev
    // invoke: char toLatin1()
    var ret0 = C.C_ZNK8QCharRef8toLatin1Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.ByteTy(false) // "char"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "toLatin1", args)
  }

  return
}

// isNull()
func (this *QCharRef) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK8QCharRef6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isNull", args)
  }

  return
}

// isPrint()
func (this *QCharRef) Isprint(args ...interface{}) (ret interface{}) {
  // isPrint()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isPrintEv
    // invoke: bool isPrint()
    var ret0 = C.C_ZNK8QCharRef7isPrintEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isPrint", args)
  }

  return
}

// toUpper()
func (this *QCharRef) Toupper(args ...interface{}) (ret interface{}) {
  // toUpper()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7toUpperEv
    // invoke: QChar toUpper()
    var ret0 = C.C_ZNK8QCharRef7toUpperEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "toUpper", args)
  }

  return
}

// isDigit()
func (this *QCharRef) Isdigit(args ...interface{}) (ret interface{}) {
  // isDigit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isDigitEv
    // invoke: bool isDigit()
    var ret0 = C.C_ZNK8QCharRef7isDigitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isDigit", args)
  }

  return
}

// setRow(uchar)
func (this *QCharRef) Setrow(args ...interface{}) () {
  // setRow(uchar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.ByteTy(false) // "uchar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef6setRowEh
    // invoke: void setRow(uchar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    C.C_ZN8QCharRef6setRowEh(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCharRef", "setRow", args)
  }

  return
}

// isPunct()
func (this *QCharRef) Ispunct(args ...interface{}) (ret interface{}) {
  // isPunct()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7isPunctEv
    // invoke: bool isPunct()
    var ret0 = C.C_ZNK8QCharRef7isPunctEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QCharRef", "isPunct", args)
  }

  return
}

// joining()
func (this *QCharRef) Joining(args ...interface{}) () {
  // joining()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK8QCharRef7joiningEv
    // invoke: QChar::Joining joining()
    C.C_ZNK8QCharRef7joiningEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "joining", args)
  }

  return
}

// toUcs4()
func (this *QStringRef) Toucs4(args ...interface{}) () {
  // toUcs4()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6toUcs4Ev
    // invoke: QVector<uint> toUcs4()
    C.C_ZNK10QStringRef6toUcs4Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "toUcs4", args)
  }

  return
}

// localeAwareCompare(const class QString &)
func (this *QStringRef) Localeawarecompare(args ...interface{}) (ret interface{}) {
  // localeAwareCompare(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef18localeAwareCompareERK7QString
    // invoke: int localeAwareCompare(const class QString &)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QStringRef18localeAwareCompareERK7QString(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "localeAwareCompare", args)
  }

  return
}

// toLocal8Bit()
func (this *QStringRef) Tolocal8Bit(args ...interface{}) (ret interface{}) {
  // toLocal8Bit()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef11toLocal8BitEv
    // invoke: QByteArray toLocal8Bit()
    var ret0 = C.C_ZNK10QStringRef11toLocal8BitEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toLocal8Bit", args)
  }

  return
}

// toUInt(_Bool *, int)
func (this *QStringRef) Touint(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef6toUIntEPbi
    // invoke: uint toUInt(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef6toUIntEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "uint"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toUInt", args)
  }

  return
}

// toShort(_Bool *, int)
func (this *QStringRef) Toshort(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef7toShortEPbi
    // invoke: short toShort(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef7toShortEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "short"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toShort", args)
  }

  return
}

// toUShort(_Bool *, int)
func (this *QStringRef) Toushort(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef8toUShortEPbi
    // invoke: ushort toUShort(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef8toUShortEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int16Ty(false) // "ushort"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toUShort", args)
  }

  return
}

// right(int)
func (this *QStringRef) Right(args ...interface{}) () {
  // right(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5rightEi
    // invoke: QStringRef right(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK10QStringRef5rightEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "right", args)
  }

  return
}

// at(int)
func (this *QStringRef) At(args ...interface{}) (ret interface{}) {
  // at(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef2atEi
    // invoke: const QChar at(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QStringRef2atEi(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "at", args)
  }

  return
}

// unicode()
func (this *QStringRef) Unicode(args ...interface{}) (ret interface{}) {
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7unicodeEv
    // invoke: const QChar * unicode()
    var ret0 = C.C_ZNK10QStringRef7unicodeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "unicode", args)
  }

  return
}

// size()
func (this *QStringRef) Size(args ...interface{}) (ret interface{}) {
  // size()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4sizeEv
    // invoke: int size()
    var ret0 = C.C_ZNK10QStringRef4sizeEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "size", args)
  }

  return
}

// length()
func (this *QStringRef) Length(args ...interface{}) (ret interface{}) {
  // length()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6lengthEv
    // invoke: int length()
    var ret0 = C.C_ZNK10QStringRef6lengthEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "length", args)
  }

  return
}

// trimmed()
func (this *QStringRef) Trimmed(args ...interface{}) () {
  // trimmed()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7trimmedEv
    // invoke: QStringRef trimmed()
    C.C_ZNK10QStringRef7trimmedEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "trimmed", args)
  }

  return
}

// mid(int, int)
func (this *QStringRef) Mid(args ...interface{}) () {
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
    // invoke: _ZNK10QStringRef3midEii
    // invoke: QStringRef mid(int, int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    C.C_ZNK10QStringRef3midEii(this.Qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "mid", args)
  }

  return
}

// constData()
func (this *QStringRef) Constdata(args ...interface{}) (ret interface{}) {
  // constData()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef9constDataEv
    // invoke: const QChar * constData()
    var ret0 = C.C_ZNK10QStringRef9constDataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "constData", args)
  }

  return
}

// toLongLong(_Bool *, int)
func (this *QStringRef) Tolonglong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef10toLongLongEPbi
    // invoke: qlonglong toLongLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef10toLongLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qlonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toLongLong", args)
  }

  return
}

// appendTo(class QString *)
func (this *QStringRef) Appendto(args ...interface{}) () {
  // appendTo(class QString *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "QString *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8appendToEP7QString
    // invoke: QStringRef appendTo(class QString *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    C.C_ZNK10QStringRef8appendToEP7QString(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "appendTo", args)
  }

  return
}

// isEmpty()
func (this *QStringRef) Isempty(args ...interface{}) (ret interface{}) {
  // isEmpty()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7isEmptyEv
    // invoke: bool isEmpty()
    var ret0 = C.C_ZNK10QStringRef7isEmptyEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "isEmpty", args)
  }

  return
}

// toString()
func (this *QStringRef) Tostring(args ...interface{}) (ret interface{}) {
  // toString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toStringEv
    // invoke: QString toString()
    var ret0 = C.C_ZNK10QStringRef8toStringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "QString"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toString", args)
  }

  return
}

// cend()
func (this *QStringRef) Cend(args ...interface{}) (ret interface{}) {
  // cend()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4cendEv
    // invoke: const QChar * cend()
    var ret0 = C.C_ZNK10QStringRef4cendEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "cend", args)
  }

  return
}

// ~QStringRef()
func (this *QStringRef) Freeqstringref(args ...interface{}) () {
  // ~QStringRef()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStringRefD0Ev
    // invoke: void ~QStringRef()
    C.C_ZN10QStringRefD2Ev(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "~QStringRef", args)
  }

  return
}

// begin()
func (this *QStringRef) Begin(args ...interface{}) (ret interface{}) {
  // begin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5beginEv
    // invoke: const QChar * begin()
    var ret0 = C.C_ZNK10QStringRef5beginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "begin", args)
  }

  return
}

// string()
func (this *QStringRef) String(args ...interface{}) (ret interface{}) {
  // string()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6stringEv
    // invoke: const QString * string()
    var ret0 = C.C_ZNK10QStringRef6stringEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QString{}) // "const QString *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "string", args)
  }

  return
}

// toULong(_Bool *, int)
func (this *QStringRef) Toulong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef7toULongEPbi
    // invoke: ulong toULong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef7toULongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "ulong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toULong", args)
  }

  return
}

// toFloat(_Bool *)
func (this *QStringRef) Tofloat(args ...interface{}) (ret interface{}) {
  // toFloat(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef7toFloatEPb
    // invoke: float toFloat(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QStringRef7toFloatEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.FloatTy(false) // "float"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toFloat", args)
  }

  return
}

// cbegin()
func (this *QStringRef) Cbegin(args ...interface{}) (ret interface{}) {
  // cbegin()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6cbeginEv
    // invoke: const QChar * cbegin()
    var ret0 = C.C_ZNK10QStringRef6cbeginEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "cbegin", args)
  }

  return
}

// end()
func (this *QStringRef) End(args ...interface{}) (ret interface{}) {
  // end()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef3endEv
    // invoke: const QChar * end()
    var ret0 = C.C_ZNK10QStringRef3endEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "end", args)
  }

  return
}

// data()
func (this *QStringRef) Data(args ...interface{}) (ret interface{}) {
  // data()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4dataEv
    // invoke: const QChar * data()
    var ret0 = C.C_ZNK10QStringRef4dataEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QChar{}) // "const QChar *"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "data", args)
  }

  return
}

// toULongLong(_Bool *, int)
func (this *QStringRef) Toulonglong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef11toULongLongEPbi
    // invoke: qulonglong toULongLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef11toULongLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int64Ty(false) // "qulonglong"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toULongLong", args)
  }

  return
}

// count()
func (this *QStringRef) Count(args ...interface{}) (ret interface{}) {
  // count()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5countEv
    // invoke: int count()
    var ret0 = C.C_ZNK10QStringRef5countEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "count", args)
  }

  return
}

// clear()
func (this *QStringRef) Clear(args ...interface{}) () {
  // clear()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStringRef5clearEv
    // invoke: void clear()
    C.C_ZN10QStringRef5clearEv(this.Qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "clear", args)
  }

  return
}

// toInt(_Bool *, int)
func (this *QStringRef) Toint(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef5toIntEPbi
    // invoke: int toInt(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef5toIntEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toInt", args)
  }

  return
}

// toLatin1()
func (this *QStringRef) Tolatin1(args ...interface{}) (ret interface{}) {
  // toLatin1()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toLatin1Ev
    // invoke: QByteArray toLatin1()
    var ret0 = C.C_ZNK10QStringRef8toLatin1Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toLatin1", args)
  }

  return
}

// isNull()
func (this *QStringRef) Isnull(args ...interface{}) (ret interface{}) {
  // isNull()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6isNullEv
    // invoke: bool isNull()
    var ret0 = C.C_ZNK10QStringRef6isNullEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.BoolTy(false) // "bool"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "isNull", args)
  }

  return
}

// toDouble(_Bool *)
func (this *QStringRef) Todouble(args ...interface{}) (ret interface{}) {
  // toDouble(_Bool *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.BoolTy(true) // "bool *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8toDoubleEPb
    // invoke: double toDouble(_Bool *)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var ret0 = C.C_ZNK10QStringRef8toDoubleEPb(this.Qclsinst, arg0)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.DoubleTy(false) // "double"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toDouble", args)
  }

  return
}

// QStringRef(const class QString *)
func NewQStringRef(args ...interface{}) *QStringRef {
  // QStringRef(const class QString *)
  // QStringRef()
  // QStringRef(const class QString *, int, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString *"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString *"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN10QStringRefC1EPK7QString
    // invoke: void QStringRef(const class QString *)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QStringRefC2EPK7QString(arg0)
    return &QStringRef{Qclsinst:qthis}
  case 1:
    // invoke: _ZN10QStringRefC1Ev
    // invoke: void QStringRef()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QStringRefC2Ev()
    return &QStringRef{Qclsinst:qthis}
  case 2:
    // invoke: _ZN10QStringRefC1EPK7QStringii
    // invoke: void QStringRef(const class QString *, int, int)
    var arg0 = args[0].(*QString).Qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(qtrt.PrimConv(args[2], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    qthis = C.C_ZN10QStringRefC2EPK7QStringii(arg0, arg1, arg2)
    return &QStringRef{Qclsinst:qthis}
  default:
    qtrt.ErrorResolve("QStringRef", "QStringRef", args)
  }

  return nil // QStringRef{Qclsinst:qthis}
}

// position()
func (this *QStringRef) Position(args ...interface{}) (ret interface{}) {
  // position()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef8positionEv
    // invoke: int position()
    var ret0 = C.C_ZNK10QStringRef8positionEv(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "int"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "position", args)
  }

  return
}

// toLong(_Bool *, int)
func (this *QStringRef) Tolong(args ...interface{}) (ret interface{}) {
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
    // invoke: _ZNK10QStringRef6toLongEPbi
    // invoke: long toLong(_Bool *, int)
    var arg0 = (unsafe.Pointer)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(qtrt.PrimConv(args[1], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg1)}
    var ret0 = C.C_ZNK10QStringRef6toLongEPbi(this.Qclsinst, arg0, arg1)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = qtrt.Int32Ty(false) // "long"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toLong", args)
  }

  return
}

// toUtf8()
func (this *QStringRef) Toutf8(args ...interface{}) (ret interface{}) {
  // toUtf8()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef6toUtf8Ev
    // invoke: QByteArray toUtf8()
    var ret0 = C.C_ZNK10QStringRef6toUtf8Ev(this.Qclsinst)
    if false {reflect.TypeOf(ret0)}
    ret = ret0
    var rety = reflect.TypeOf(QByteArray{}) // "QByteArray"
    if reflect.TypeOf(ret0).ConvertibleTo(rety) {
        ret = reflect.ValueOf(ret0).Convert(rety).Interface()
    } else {
        ret = qtrt.HandyConvert2go(ret0, rety)
    }
  default:
    qtrt.ErrorResolve("QStringRef", "toUtf8", args)
  }

  return
}

// left(int)
func (this *QStringRef) Left(args ...interface{}) () {
  // left(int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef4leftEi
    // invoke: QStringRef left(int)
    var arg0 = C.int32_t(qtrt.PrimConv(args[0], qtrt.Int32Ty(false)).(int32))
    if false {fmt.Println(arg0)}
    C.C_ZNK10QStringRef4leftEi(this.Qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "left", args)
  }

  return
}

// <= body block end

