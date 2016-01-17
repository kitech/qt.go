package qt5
// auto generated, do not modify.
// created: Sun Jan 17 14:31:14 2016
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
extern void _ZNK7QString6toUcs4Ev(void* qthis); // 4
  // proto:  int QString::toWCharArray(wchar_t * array);
extern void _ZNK7QString12toWCharArrayEPw(void* qthis, wchar_t* arg0); // 2
  // proto:  const QChar * QString::unicode();
extern void _ZNK7QString7unicodeEv(void* qthis); // 2
  // proto: static QString QString::fromLatin1(const QByteArray & str);
extern void _ZN7QString10fromLatin1ERK10QByteArray(void* arg0); // 2
  // proto: static QString QString::fromLatin1(const char * str, int size);
extern void _ZN7QString10fromLatin1EPKci(unsigned char* arg0, int32_t arg1); // 2
  // proto:  void QString::reserve(int size);
extern void _ZN7QString7reserveEi(void* qthis, int32_t arg0); // 2
  // proto:  void QString::swap(QString & other);
extern void _ZN7QString4swapERS_(void* qthis, void* arg0); // 2
  // proto:  QStringRef QString::rightRef(int n);
extern void _ZNK7QString8rightRefEi(void* qthis, int32_t arg0); // 4
  // proto:  QStringRef QString::leftRef(int n);
extern void _ZNK7QString7leftRefEi(void* qthis, int32_t arg0); // 4
  // proto:  QString QString::rightJustified(int width, QChar fill, bool trunc);
extern void _ZNK7QString14rightJustifiedEi5QCharb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  void QString::detach();
extern void _ZN7QString6detachEv(void* qthis); // 2
  // proto:  QString & QString::insert(int i, QChar c);
extern void _ZN7QString6insertEi5QChar(void* qthis, int32_t arg0, void* arg1); // 4
  // proto:  QString & QString::insert(int i, const QString & s);
extern void _ZN7QString6insertEiRKS_(void* qthis, int32_t arg0, void* arg1); // 2
  // proto:  QString & QString::insert(int i, const QChar * uc, int len);
extern void _ZN7QString6insertEiPK5QChari(void* qthis, int32_t arg0, void* arg1, int32_t arg2); // 4
  // proto:  QStringRef QString::midRef(int position, int n);
extern void _ZNK7QString6midRefEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::remove(const QRegExp & rx);
extern void _ZN7QString6removeERK7QRegExp(void* qthis, void* arg0); // 2
  // proto:  QString & QString::remove(const QRegularExpression & re);
extern void _ZN7QString6removeERK18QRegularExpression(void* qthis, void* arg0); // 2
  // proto:  QString & QString::remove(int i, int len);
extern void _ZN7QString6removeEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QString QString::trimmed();
extern void _ZNKR7QString7trimmedEv(void* qthis); // 2
  // proto:  std::wstring QString::toStdWString();
extern void _ZNK7QString12toStdWStringEv(void* qthis); // 2
  // proto:  const_iterator QString::constEnd();
extern void _ZNK7QString8constEndEv(void* qthis); // 2
  // proto:  const ushort * QString::utf16();
extern void _ZNK7QString5utf16Ev(void* qthis); // 4
  // proto:  QString QString::right(int n);
extern void _ZNK7QString5rightEi(void* qthis, int32_t arg0); // 4
  // proto:  QByteArray QString::toLocal8Bit();
extern void _ZNKR7QString11toLocal8BitEv(void* qthis); // 2
  // proto:  uint QString::toUInt(bool * ok, int base);
extern void _ZNK7QString6toUIntEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  void QString::QString();
extern void _ZN7QStringC2Ev(void* qthis); // 1
  // proto:  void QString::QString(const QString & );
extern void _ZN7QStringC2ERKS_(void* qthis, void* arg0); // 1
  // proto:  void QString::QString(int size, QChar c);
extern void _ZN7QStringC2Ei5QChar(void* qthis, int32_t arg0, void* arg1); // 3
  // proto:  void QString::QString(const char * ch);
extern void _ZN7QStringC2EPKc(void* qthis, unsigned char* arg0); // 1
  // proto:  void QString::QString(const QChar * unicode, int size);
extern void _ZN7QStringC2EPK5QChari(void* qthis, void* arg0, int32_t arg1); // 3
  // proto:  void QString::QString(QChar c);
extern void _ZN7QStringC2E5QChar(void* qthis, void* arg0); // 3
  // proto:  void QString::QString(const QByteArray & a);
extern void _ZN7QStringC2ERK10QByteArray(void* qthis, void* arg0); // 1
  // proto:  QString QString::arg(const QString & a, int fieldWidth, QChar fillChar);
extern void _ZNK7QString3argERKS_i5QChar(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4);
extern void _ZNK7QString3argERKS_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3);
extern void _ZNK7QString3argERKS_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2); // 2
  // proto:  QString QString::arg(qlonglong a, int fieldwidth, int base, QChar fillChar);
extern void _ZNK7QString3argExii5QChar(void* qthis, int64_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 4
  // proto:  QString QString::arg(long a, int fieldwidth, int base, QChar fillChar);
extern void _ZNK7QString3argElii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(ushort a, int fieldWidth, int base, QChar fillChar);
extern void _ZNK7QString3argEtii5QChar(void* qthis, int16_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(int a, int fieldWidth, int base, QChar fillChar);
extern void _ZNK7QString3argEiii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8);
extern void _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7); // 2
  // proto:  QString QString::arg(short a, int fieldWidth, int base, QChar fillChar);
extern void _ZNK7QString3argEsii5QChar(void* qthis, int16_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7);
extern void _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5);
extern void _ZNK7QString3argERKS_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4); // 2
  // proto:  QString QString::arg(uint a, int fieldWidth, int base, QChar fillChar);
extern void _ZNK7QString3argEjii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(const QString & a1, const QString & a2);
extern void _ZNK7QString3argERKS_S1_(void* qthis, void* arg0, void* arg1); // 2
  // proto:  QString QString::arg(double a, int fieldWidth, char fmt, int prec, QChar fillChar);
extern void _ZNK7QString3argEdici5QChar(void* qthis, double arg0, int32_t arg1, unsigned char arg2, int32_t arg3, void* arg4); // 4
  // proto:  QString QString::arg(qulonglong a, int fieldwidth, int base, QChar fillChar);
extern void _ZNK7QString3argEyii5QChar(void* qthis, int64_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 4
  // proto:  QString QString::arg(char a, int fieldWidth, QChar fillChar);
extern void _ZNK7QString3argEci5QChar(void* qthis, unsigned char arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8, const QString & a9);
extern void _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8); // 2
  // proto:  QString QString::arg(ulong a, int fieldwidth, int base, QChar fillChar);
extern void _ZNK7QString3argEmii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3); // 2
  // proto:  QString QString::arg(QChar a, int fieldWidth, QChar fillChar);
extern void _ZNK7QString3argE5QChariS0_(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6);
extern void _ZNK7QString3argERKS_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5); // 2
  // proto:  QString & QString::append(QChar c);
extern void _ZN7QString6appendE5QChar(void* qthis, void* arg0); // 4
  // proto:  QString & QString::append(const char * s);
extern void _ZN7QString6appendEPKc(void* qthis, unsigned char* arg0); // 2
  // proto:  QString & QString::append(const QByteArray & s);
extern void _ZN7QString6appendERK10QByteArray(void* qthis, void* arg0); // 2
  // proto:  QString & QString::append(const QString & s);
extern void _ZN7QString6appendERKS_(void* qthis, void* arg0); // 4
  // proto:  QString & QString::append(const QChar * uc, int len);
extern void _ZN7QString6appendEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::capacity();
extern void _ZNK7QString8capacityEv(void* qthis); // 2
  // proto:  QString QString::toLower();
extern void _ZNKR7QString7toLowerEv(void* qthis); // 2
  // proto:  QString QString::toHtmlEscaped();
extern void _ZNK7QString13toHtmlEscapedEv(void* qthis); // 4
  // proto:  bool QString::isNull();
extern void _ZNK7QString6isNullEv(void* qthis); // 2
  // proto: static int QString::localeAwareCompare(const QString & s1, const QString & s2);
extern void _ZN7QString18localeAwareCompareERKS_S1_(void* arg0, void* arg1); // 2
  // proto:  int QString::localeAwareCompare(const QString & s);
extern void _ZNK7QString18localeAwareCompareERKS_(void* qthis, void* arg0); // 4
  // proto: static QString QString::fromLocal8Bit(const QByteArray & str);
extern void _ZN7QString13fromLocal8BitERK10QByteArray(void* arg0); // 2
  // proto: static QString QString::fromLocal8Bit(const char * str, int size);
extern void _ZN7QString13fromLocal8BitEPKci(unsigned char* arg0, int32_t arg1); // 2
  // proto: static QString QString::fromWCharArray(const wchar_t * string, int size);
extern void _ZN7QString14fromWCharArrayEPKwi(wchar_t* arg0, int32_t arg1); // 2
  // proto:  std::string QString::toStdString();
extern void _ZNK7QString11toStdStringEv(void* qthis); // 2
  // proto:  int QString::indexOf(const QRegularExpression & re, int from, QRegularExpressionMatch * rmatch);
extern void _ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  int QString::indexOf(const QRegularExpression & re, int from);
extern void _ZNK7QString7indexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::indexOf(QRegExp & , int from);
extern void _ZNK7QString7indexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::indexOf(const QRegExp & , int from);
extern void _ZNK7QString7indexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  const_iterator QString::cbegin();
extern void _ZNK7QString6cbeginEv(void* qthis); // 2
  // proto:  QByteArray QString::toLatin1();
extern void _ZNKR7QString8toLatin1Ev(void* qthis); // 2
  // proto:  void QString::chop(int n);
extern void _ZN7QString4chopEi(void* qthis, int32_t arg0); // 4
  // proto:  std::u32string QString::toStdU32String();
extern void _ZNK7QString14toStdU32StringEv(void* qthis); // 2
  // proto:  int QString::length();
extern void _ZNK7QString6lengthEv(void* qthis); // 2
  // proto:  double QString::toDouble(bool * ok);
extern void _ZNK7QString8toDoubleEPb(void* qthis, bool* arg0); // 4
  // proto:  bool QString::isSharedWith(const QString & other);
extern void _ZNK7QString12isSharedWithERKS_(void* qthis, void* arg0); // 2
  // proto:  ushort QString::toUShort(bool * ok, int base);
extern void _ZNK7QString8toUShortEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  const_iterator QString::constBegin();
extern void _ZNK7QString10constBeginEv(void* qthis); // 2
  // proto:  short QString::toShort(bool * ok, int base);
extern void _ZNK7QString7toShortEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto: static QString QString::number(uint , int base);
extern void _ZN7QString6numberEji(int32_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(double , char f, int prec);
extern void _ZN7QString6numberEdci(double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto: static QString QString::number(qulonglong , int base);
extern void _ZN7QString6numberEyi(int64_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(ulong , int base);
extern void _ZN7QString6numberEmi(int32_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(qlonglong , int base);
extern void _ZN7QString6numberExi(int64_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(long , int base);
extern void _ZN7QString6numberEli(int32_t arg0, int32_t arg1); // 4
  // proto: static QString QString::number(int , int base);
extern void _ZN7QString6numberEii(int32_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::replace(const QRegExp & rx, const QString & after);
extern void _ZN7QString7replaceERK7QRegExpRKS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString & QString::replace(int i, int len, QChar after);
extern void _ZN7QString7replaceEii5QChar(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString & QString::replace(const QRegularExpression & re, const QString & after);
extern void _ZN7QString7replaceERK18QRegularExpressionRKS_(void* qthis, void* arg0, void* arg1); // 4
  // proto:  QString & QString::replace(int i, int len, const QChar * s, int slen);
extern void _ZN7QString7replaceEiiPK5QChari(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3); // 4
  // proto:  QString & QString::replace(int i, int len, const QString & after);
extern void _ZN7QString7replaceEiiRKS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2); // 4
  // proto:  void QString::push_back(QChar c);
extern void _ZN7QString9push_backE5QChar(void* qthis, void* arg0); // 2
  // proto:  void QString::push_back(const QString & s);
extern void _ZN7QString9push_backERKS_(void* qthis, void* arg0); // 2
  // proto:  QString QString::simplified();
extern void _ZNKR7QString10simplifiedEv(void* qthis); // 2
  // proto:  int QString::size();
extern void _ZNK7QString4sizeEv(void* qthis); // 2
  // proto:  void QString::truncate(int pos);
extern void _ZN7QString8truncateEi(void* qthis, int32_t arg0); // 4
  // proto:  bool QString::contains(const QRegularExpression & re);
extern void _ZNK7QString8containsERK18QRegularExpression(void* qthis, void* arg0); // 4
  // proto:  bool QString::contains(const QRegularExpression & re, QRegularExpressionMatch * match);
extern void _ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch(void* qthis, void* arg0, void* arg1); // 4
  // proto:  bool QString::contains(QRegExp & rx);
extern void _ZNK7QString8containsER7QRegExp(void* qthis, void* arg0); // 2
  // proto:  bool QString::contains(const QRegExp & rx);
extern void _ZNK7QString8containsERK7QRegExp(void* qthis, void* arg0); // 2
  // proto:  bool QString::isSimpleText();
extern void _ZNK7QString12isSimpleTextEv(void* qthis); // 4
  // proto:  const_iterator QString::cend();
extern void _ZNK7QString4cendEv(void* qthis); // 2
  // proto:  QString QString::leftJustified(int width, QChar fill, bool trunc);
extern void _ZNK7QString13leftJustifiedEi5QCharb(void* qthis, int32_t arg0, void* arg1, bool arg2); // 4
  // proto:  QString QString::repeated(int times);
extern void _ZNK7QString8repeatedEi(void* qthis, int32_t arg0); // 4
  // proto: static QString QString::fromRawData(const QChar * , int size);
extern void _ZN7QString11fromRawDataEPK5QChari(void* arg0, int32_t arg1); // 4
  // proto:  void QString::squeeze();
extern void _ZN7QString7squeezeEv(void* qthis); // 2
  // proto:  void QString::resize(int size);
extern void _ZN7QString6resizeEi(void* qthis, int32_t arg0); // 4
  // proto:  qulonglong QString::toULongLong(bool * ok, int base);
extern void _ZNK7QString11toULongLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  int QString::count(const QRegularExpression & re);
extern void _ZNK7QString5countERK18QRegularExpression(void* qthis, void* arg0); // 4
  // proto:  int QString::count();
extern void _ZNK7QString5countEv(void* qthis); // 2
  // proto:  int QString::count(const QRegExp & );
extern void _ZNK7QString5countERK7QRegExp(void* qthis, void* arg0); // 4
  // proto:  bool QString::isRightToLeft();
extern void _ZNK7QString13isRightToLeftEv(void* qthis); // 4
  // proto:  QByteArray QString::toUtf8();
extern void _ZNKR7QString6toUtf8Ev(void* qthis); // 2
  // proto: static QString QString::fromUtf8(const char * str, int size);
extern void _ZN7QString8fromUtf8EPKci(unsigned char* arg0, int32_t arg1); // 2
  // proto: static QString QString::fromUtf8(const QByteArray & str);
extern void _ZN7QString8fromUtf8ERK10QByteArray(void* arg0); // 2
  // proto:  std::u16string QString::toStdU16String();
extern void _ZNK7QString14toStdU16StringEv(void* qthis); // 2
  // proto:  float QString::toFloat(bool * ok);
extern void _ZNK7QString7toFloatEPb(void* qthis, bool* arg0); // 4
  // proto:  const QChar QString::at(int i);
extern void _ZNK7QString2atEi(void* qthis, int32_t arg0); // 2
  // proto:  QString & QString::fill(QChar c, int size);
extern void _ZN7QString4fillE5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  ulong QString::toULong(bool * ok, int base);
extern void _ZNK7QString7toULongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  iterator QString::end();
extern void _ZN7QString3endEv(void* qthis); // 2
  // proto: static QString QString::fromUcs4(const uint * , int size);
extern void _ZN7QString8fromUcs4EPKji(int32_t* arg0, int32_t arg1); // 4
  // proto:  QString QString::mid(int position, int n);
extern void _ZNK7QString3midEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::prepend(const QString & s);
extern void _ZN7QString7prependERKS_(void* qthis, void* arg0); // 2
  // proto:  QString & QString::prepend(QChar c);
extern void _ZN7QString7prependE5QChar(void* qthis, void* arg0); // 2
  // proto:  QString & QString::prepend(const char * s);
extern void _ZN7QString7prependEPKc(void* qthis, unsigned char* arg0); // 2
  // proto:  QString & QString::prepend(const QByteArray & s);
extern void _ZN7QString7prependERK10QByteArray(void* qthis, void* arg0); // 2
  // proto:  const QChar * QString::constData();
extern void _ZNK7QString9constDataEv(void* qthis); // 2
  // proto:  qlonglong QString::toLongLong(bool * ok, int base);
extern void _ZNK7QString10toLongLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QString & QString::setUnicode(const QChar * unicode, int size);
extern void _ZN7QString10setUnicodeEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto: static QString QString::fromUtf16(const ushort * , int size);
extern void _ZN7QString9fromUtf16EPKti(int16_t* arg0, int32_t arg1); // 4
  // proto:  bool QString::isEmpty();
extern void _ZNK7QString7isEmptyEv(void* qthis); // 2
  // proto:  QString & QString::setRawData(const QChar * unicode, int size);
extern void _ZN7QString10setRawDataEPK5QChari(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  bool QString::isDetached();
extern void _ZNK7QString10isDetachedEv(void* qthis); // 2
  // proto:  iterator QString::begin();
extern void _ZN7QString5beginEv(void* qthis); // 2
  // proto:  int QString::lastIndexOf(QRegExp & , int from);
extern void _ZNK7QString11lastIndexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::lastIndexOf(const QRegExp & , int from);
extern void _ZNK7QString11lastIndexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from);
extern void _ZNK7QString11lastIndexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1); // 4
  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from, QRegularExpressionMatch * rmatch);
extern void _ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch(void* qthis, void* arg0, int32_t arg1, void* arg2); // 4
  // proto:  QString & QString::setNum(ushort , int base);
extern void _ZN7QString6setNumEti(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(qulonglong , int base);
extern void _ZN7QString6setNumEyi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::setNum(qlonglong , int base);
extern void _ZN7QString6setNumExi(void* qthis, int64_t arg0, int32_t arg1); // 4
  // proto:  QString & QString::setNum(ulong , int base);
extern void _ZN7QString6setNumEmi(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(short , int base);
extern void _ZN7QString6setNumEsi(void* qthis, int16_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(long , int base);
extern void _ZN7QString6setNumEli(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(uint , int base);
extern void _ZN7QString6setNumEji(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(int , int base);
extern void _ZN7QString6setNumEii(void* qthis, int32_t arg0, int32_t arg1); // 2
  // proto:  QString & QString::setNum(double , char f, int prec);
extern void _ZN7QString6setNumEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2); // 4
  // proto:  QString & QString::setNum(float , char f, int prec);
extern void _ZN7QString6setNumEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2); // 2
  // proto:  QString & QString::setUtf16(const ushort * utf16, int size);
extern void _ZN7QString8setUtf16EPKti(void* qthis, int16_t* arg0, int32_t arg1); // 2
  // proto:  QChar * QString::data();
extern void _ZN7QString4dataEv(void* qthis); // 2
  // proto:  long QString::toLong(bool * ok, int base);
extern void _ZNK7QString6toLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  void QString::~QString();
extern void _ZN7QStringD2Ev(void* qthis); // 2
  // proto:  void QString::clear();
extern void _ZN7QString5clearEv(void* qthis); // 2
  // proto:  int QString::toInt(bool * ok, int base);
extern void _ZNK7QString5toIntEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QString QString::toCaseFolded();
extern void _ZNKR7QString12toCaseFoldedEv(void* qthis); // 2
  // proto:  QString QString::toUpper();
extern void _ZNKR7QString7toUpperEv(void* qthis); // 2
  // proto:  void QString::push_front(QChar c);
extern void _ZN7QString10push_frontE5QChar(void* qthis, void* arg0); // 2
  // proto:  void QString::push_front(const QString & s);
extern void _ZN7QString10push_frontERKS_(void* qthis, void* arg0); // 2
  // proto:  QString QString::left(int n);
extern void _ZNK7QString4leftEi(void* qthis, int32_t arg0); // 4
  // proto:  void QLatin1String::QLatin1String(const QByteArray & s);
extern void _ZN13QLatin1StringC2ERK10QByteArray(void* qthis, void* arg0); // 1
  // proto:  void QLatin1String::QLatin1String(const char * s);
extern void _ZN13QLatin1StringC2EPKc(void* qthis, unsigned char* arg0); // 1
  // proto:  void QLatin1String::QLatin1String(const char * s, int sz);
extern void _ZN13QLatin1StringC2EPKci(void* qthis, unsigned char* arg0, int32_t arg1); // 1
  // proto:  const char * QLatin1String::latin1();
extern void _ZNK13QLatin1String6latin1Ev(void* qthis); // 2
  // proto:  const char * QLatin1String::data();
extern void _ZNK13QLatin1String4dataEv(void* qthis); // 2
  // proto:  int QLatin1String::size();
extern void _ZNK13QLatin1String4sizeEv(void* qthis); // 2
  // proto:  bool QCharRef::isLetterOrNumber();
extern void _ZN8QCharRef16isLetterOrNumberEv(void* qthis); // 2
  // proto:  bool QCharRef::isLetter();
extern void _ZNK8QCharRef8isLetterEv(void* qthis); // 2
  // proto:  bool QCharRef::hasMirrored();
extern void _ZNK8QCharRef11hasMirroredEv(void* qthis); // 2
  // proto:  QChar::UnicodeVersion QCharRef::unicodeVersion();
extern void _ZNK8QCharRef14unicodeVersionEv(void* qthis); // 2
  // proto:  bool QCharRef::isTitleCase();
extern void _ZNK8QCharRef11isTitleCaseEv(void* qthis); // 2
  // proto:  void QCharRef::setCell(uchar cell);
extern void _ZN8QCharRef7setCellEh(void* qthis, unsigned char arg0); // 2
  // proto:  ushort & QCharRef::unicode();
extern void _ZN8QCharRef7unicodeEv(void* qthis); // 2
  // proto:  bool QCharRef::isSpace();
extern void _ZNK8QCharRef7isSpaceEv(void* qthis); // 2
  // proto:  uchar QCharRef::row();
extern void _ZNK8QCharRef3rowEv(void* qthis); // 2
  // proto:  QChar::Category QCharRef::category();
extern void _ZNK8QCharRef8categoryEv(void* qthis); // 2
  // proto:  bool QCharRef::isUpper();
extern void _ZNK8QCharRef7isUpperEv(void* qthis); // 2
  // proto:  QChar QCharRef::toLower();
extern void _ZNK8QCharRef7toLowerEv(void* qthis); // 2
  // proto:  QChar::Script QCharRef::script();
extern void _ZNK8QCharRef6scriptEv(void* qthis); // 2
  // proto:  QChar::JoiningType QCharRef::joiningType();
extern void _ZNK8QCharRef11joiningTypeEv(void* qthis); // 2
  // proto:  uchar QCharRef::cell();
extern void _ZNK8QCharRef4cellEv(void* qthis); // 2
  // proto:  int QCharRef::digitValue();
extern void _ZNK8QCharRef10digitValueEv(void* qthis); // 2
  // proto:  QString QCharRef::decomposition();
extern void _ZNK8QCharRef13decompositionEv(void* qthis); // 2
  // proto:  QChar::Direction QCharRef::direction();
extern void _ZNK8QCharRef9directionEv(void* qthis); // 2
  // proto:  QChar QCharRef::mirroredChar();
extern void _ZNK8QCharRef12mirroredCharEv(void* qthis); // 2
  // proto:  uchar QCharRef::combiningClass();
extern void _ZNK8QCharRef14combiningClassEv(void* qthis); // 2
  // proto:  QChar::Decomposition QCharRef::decompositionTag();
extern void _ZNK8QCharRef16decompositionTagEv(void* qthis); // 2
  // proto:  bool QCharRef::isMark();
extern void _ZNK8QCharRef6isMarkEv(void* qthis); // 2
  // proto:  bool QCharRef::isLower();
extern void _ZNK8QCharRef7isLowerEv(void* qthis); // 2
  // proto:  bool QCharRef::isNumber();
extern void _ZNK8QCharRef8isNumberEv(void* qthis); // 2
  // proto:  QChar QCharRef::toTitleCase();
extern void _ZNK8QCharRef11toTitleCaseEv(void* qthis); // 2
  // proto:  char QCharRef::toLatin1();
extern void _ZNK8QCharRef8toLatin1Ev(void* qthis); // 2
  // proto:  bool QCharRef::isNull();
extern void _ZNK8QCharRef6isNullEv(void* qthis); // 2
  // proto:  bool QCharRef::isPrint();
extern void _ZNK8QCharRef7isPrintEv(void* qthis); // 2
  // proto:  QChar QCharRef::toUpper();
extern void _ZNK8QCharRef7toUpperEv(void* qthis); // 2
  // proto:  bool QCharRef::isDigit();
extern void _ZNK8QCharRef7isDigitEv(void* qthis); // 2
  // proto:  void QCharRef::setRow(uchar row);
extern void _ZN8QCharRef6setRowEh(void* qthis, unsigned char arg0); // 2
  // proto:  bool QCharRef::isPunct();
extern void _ZNK8QCharRef7isPunctEv(void* qthis); // 2
  // proto:  QChar::Joining QCharRef::joining();
extern void _ZNK8QCharRef7joiningEv(void* qthis); // 2
  // proto:  QVector<uint> QStringRef::toUcs4();
extern void _ZNK10QStringRef6toUcs4Ev(void* qthis); // 4
  // proto:  int QStringRef::localeAwareCompare(const QString & s);
extern void _ZNK10QStringRef18localeAwareCompareERK7QString(void* qthis, void* arg0); // 2
  // proto:  QByteArray QStringRef::toLocal8Bit();
extern void _ZNK10QStringRef11toLocal8BitEv(void* qthis); // 4
  // proto:  uint QStringRef::toUInt(bool * ok, int base);
extern void _ZNK10QStringRef6toUIntEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  short QStringRef::toShort(bool * ok, int base);
extern void _ZNK10QStringRef7toShortEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  ushort QStringRef::toUShort(bool * ok, int base);
extern void _ZNK10QStringRef8toUShortEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QStringRef QStringRef::right(int n);
extern void _ZNK10QStringRef5rightEi(void* qthis, int32_t arg0); // 4
  // proto:  const QChar QStringRef::at(int i);
extern void _ZNK10QStringRef2atEi(void* qthis, int32_t arg0); // 2
  // proto:  const QChar * QStringRef::unicode();
extern void _ZNK10QStringRef7unicodeEv(void* qthis); // 2
  // proto:  int QStringRef::size();
extern void _ZNK10QStringRef4sizeEv(void* qthis); // 2
  // proto:  int QStringRef::length();
extern void _ZNK10QStringRef6lengthEv(void* qthis); // 2
  // proto:  QStringRef QStringRef::trimmed();
extern void _ZNK10QStringRef7trimmedEv(void* qthis); // 4
  // proto:  QStringRef QStringRef::mid(int pos, int n);
extern void _ZNK10QStringRef3midEii(void* qthis, int32_t arg0, int32_t arg1); // 4
  // proto:  const QChar * QStringRef::constData();
extern void _ZNK10QStringRef9constDataEv(void* qthis); // 2
  // proto:  qlonglong QStringRef::toLongLong(bool * ok, int base);
extern void _ZNK10QStringRef10toLongLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QStringRef QStringRef::appendTo(QString * string);
extern void _ZNK10QStringRef8appendToEP7QString(void* qthis, void* arg0); // 4
  // proto:  bool QStringRef::isEmpty();
extern void _ZNK10QStringRef7isEmptyEv(void* qthis); // 2
  // proto:  QString QStringRef::toString();
extern void _ZNK10QStringRef8toStringEv(void* qthis); // 4
  // proto:  const QChar * QStringRef::cend();
extern void _ZNK10QStringRef4cendEv(void* qthis); // 2
  // proto:  void QStringRef::~QStringRef();
extern void _ZN10QStringRefD2Ev(void* qthis); // 2
  // proto:  const QChar * QStringRef::begin();
extern void _ZNK10QStringRef5beginEv(void* qthis); // 2
  // proto:  const QString * QStringRef::string();
extern void _ZNK10QStringRef6stringEv(void* qthis); // 2
  // proto:  ulong QStringRef::toULong(bool * ok, int base);
extern void _ZNK10QStringRef7toULongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  float QStringRef::toFloat(bool * ok);
extern void _ZNK10QStringRef7toFloatEPb(void* qthis, bool* arg0); // 4
  // proto:  const QChar * QStringRef::cbegin();
extern void _ZNK10QStringRef6cbeginEv(void* qthis); // 2
  // proto:  const QChar * QStringRef::end();
extern void _ZNK10QStringRef3endEv(void* qthis); // 2
  // proto:  const QChar * QStringRef::data();
extern void _ZNK10QStringRef4dataEv(void* qthis); // 2
  // proto:  qulonglong QStringRef::toULongLong(bool * ok, int base);
extern void _ZNK10QStringRef11toULongLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  int QStringRef::count();
extern void _ZNK10QStringRef5countEv(void* qthis); // 2
  // proto:  void QStringRef::clear();
extern void _ZN10QStringRef5clearEv(void* qthis); // 2
  // proto:  int QStringRef::toInt(bool * ok, int base);
extern void _ZNK10QStringRef5toIntEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QByteArray QStringRef::toLatin1();
extern void _ZNK10QStringRef8toLatin1Ev(void* qthis); // 4
  // proto:  bool QStringRef::isNull();
extern void _ZNK10QStringRef6isNullEv(void* qthis); // 2
  // proto:  double QStringRef::toDouble(bool * ok);
extern void _ZNK10QStringRef8toDoubleEPb(void* qthis, bool* arg0); // 4
  // proto:  void QStringRef::QStringRef(const QString * string);
extern void _ZN10QStringRefC2EPK7QString(void* qthis, void* arg0); // 1
  // proto:  void QStringRef::QStringRef();
extern void _ZN10QStringRefC2Ev(void* qthis); // 1
  // proto:  void QStringRef::QStringRef(const QString * string, int position, int size);
extern void _ZN10QStringRefC2EPK7QStringii(void* qthis, void* arg0, int32_t arg1, int32_t arg2); // 1
  // proto:  int QStringRef::position();
extern void _ZNK10QStringRef8positionEv(void* qthis); // 2
  // proto:  long QStringRef::toLong(bool * ok, int base);
extern void _ZNK10QStringRef6toLongEPbi(void* qthis, bool* arg0, int32_t arg1); // 4
  // proto:  QByteArray QStringRef::toUtf8();
extern void _ZNK10QStringRef6toUtf8Ev(void* qthis); // 4
  // proto:  QStringRef QStringRef::left(int n);
extern void _ZNK10QStringRef4leftEi(void* qthis, int32_t arg0); // 4
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
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QString)=8
type QString struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QLatin1String)=16
type QLatin1String struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QCharRef)=16
type QCharRef struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// class sizeof(QStringRef)=16
type QStringRef struct {
  // qbase: None;
  qclsinst unsafe.Pointer /* *C.void */;
}

// toUcs4()
func (this *QString) toUcs4(args ...interface{}) () {
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
    C._ZNK7QString6toUcs4Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toUcs4", args)
  }

}

// toWCharArray(wchar_t *)
func (this *QString) toWCharArray(args ...interface{}) () {
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
    var arg0 = (*C.wchar_t)((unsafe.Pointer)(reflect.ValueOf(args[0].([]rune)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZNK7QString12toWCharArrayEPw(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "toWCharArray", args)
  }

}

// unicode()
func (this *QString) unicode(args ...interface{}) () {
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
    C._ZNK7QString7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "unicode", args)
  }

}

// fromLatin1(const class QByteArray &)
func (this *QString) fromLatin1_s(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString10fromLatin1ERK10QByteArray(arg0)
  case 1:
    // invoke: _ZN7QString10fromLatin1EPKci
    // invoke: QString fromLatin1(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString10fromLatin1EPKci(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fromLatin1", args)
  }

}

// reserve(int)
func (this *QString) reserve(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QString7reserveEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "reserve", args)
  }

}

// swap(class QString &)
func (this *QString) swap(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "swap", args)
  }

}

// rightRef(int)
func (this *QString) rightRef(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QString8rightRefEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "rightRef", args)
  }

}

// leftRef(int)
func (this *QString) leftRef(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QString7leftRefEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "leftRef", args)
  }

}

// rightJustified(int, class QChar, _Bool)
func (this *QString) rightJustified(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C._ZNK7QString14rightJustifiedEi5QCharb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "rightJustified", args)
  }

}

// detach()
func (this *QString) detach(args ...interface{}) () {
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
    C._ZN7QString6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "detach", args)
  }

}

// insert(int, class QChar)
func (this *QString) insert(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString6insertEi5QChar(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QString6insertEiRKS_
    // invoke: QString & insert(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString6insertEiRKS_(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN7QString6insertEiPK5QChari
    // invoke: QString & insert(int, const class QChar *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN7QString6insertEiPK5QChari(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "insert", args)
  }

}

// midRef(int, int)
func (this *QString) midRef(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString6midRefEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "midRef", args)
  }

}

// remove(const class QRegExp &)
func (this *QString) remove(args ...interface{}) () {
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
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6removeERK7QRegExp(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString6removeERK18QRegularExpression
    // invoke: QString & remove(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6removeERK18QRegularExpression(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN7QString6removeEii
    // invoke: QString & remove(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6removeEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "remove", args)
  }

}

// trimmed()
func (this *QString) trimmed(args ...interface{}) () {
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
    C._ZNKR7QString7trimmedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "trimmed", args)
  }

}

// toStdWString()
func (this *QString) toStdWString(args ...interface{}) () {
  // toStdWString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12toStdWStringEv
    // invoke: std::wstring toStdWString()
    C._ZNK7QString12toStdWStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdWString", args)
  }

}

// constEnd()
func (this *QString) constEnd(args ...interface{}) () {
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
    C._ZNK7QString8constEndEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "constEnd", args)
  }

}

// utf16()
func (this *QString) utf16(args ...interface{}) () {
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
    C._ZNK7QString5utf16Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "utf16", args)
  }

}

// right(int)
func (this *QString) right(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QString5rightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "right", args)
  }

}

// toLocal8Bit()
func (this *QString) toLocal8Bit(args ...interface{}) () {
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
    C._ZNKR7QString11toLocal8BitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toLocal8Bit", args)
  }

}

// toUInt(_Bool *, int)
func (this *QString) toUInt(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString6toUIntEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toUInt", args)
  }

}

// QString()
func NewQString(args ...interface{}) QString {
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
    C._ZN7QStringC2Ev(qthis)
  case 1:
    // invoke: _ZN7QStringC1ERKS_
    // invoke: void QString(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QStringC2ERKS_(qthis, arg0)
  case 2:
    // invoke: _ZN7QStringC1Ei5QChar
    // invoke: void QString(int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QStringC2Ei5QChar(qthis, arg0, arg1)
  case 3:
    // invoke: _ZN7QStringC1EPKc
    // invoke: void QString(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QStringC2EPKc(qthis, arg0)
  case 4:
    // invoke: _ZN7QStringC1EPK5QChari
    // invoke: void QString(const class QChar *, int)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QStringC2EPK5QChari(qthis, arg0, arg1)
  case 5:
    // invoke: _ZN7QStringC1E5QChar
    // invoke: void QString(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QStringC2E5QChar(qthis, arg0)
  case 6:
    // invoke: _ZN7QStringC1ERK10QByteArray
    // invoke: void QString(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN7QStringC2ERK10QByteArray(qthis, arg0)
  default:
    qtrt.ErrorResolve("QString", "QString", args)
  }

  return QString{}
}

// arg(const class QString &, int, class QChar)
func (this *QString) arg(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argERKS_i5QChar(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argERKS_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
    // invoke: _ZNK7QString3argERKS_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argERKS_S1_S1_(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZNK7QString3argExii5QChar
    // invoke: QString arg(qlonglong, int, int, class QChar)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argExii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZNK7QString3argElii5QChar
    // invoke: QString arg(long, int, int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argElii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 5:
    // invoke: _ZNK7QString3argEtii5QChar
    // invoke: QString arg(ushort, int, int, class QChar)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argEtii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 6:
    // invoke: _ZNK7QString3argEiii5QChar
    // invoke: QString arg(int, int, int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argEiii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 7:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QString).qclsinst
    if false {fmt.Println(arg6)}
    var arg7 = args[7].(QString).qclsinst
    if false {fmt.Println(arg7)}
    C._ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
  case 8:
    // invoke: _ZNK7QString3argEsii5QChar
    // invoke: QString arg(short, int, int, class QChar)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argEsii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 9:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QString).qclsinst
    if false {fmt.Println(arg6)}
    C._ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 10:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    C._ZNK7QString3argERKS_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 11:
    // invoke: _ZNK7QString3argEjii5QChar
    // invoke: QString arg(uint, int, int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argEjii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 12:
    // invoke: _ZNK7QString3argERKS_S1_
    // invoke: QString arg(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QString3argERKS_S1_(this.qclsinst, arg0, arg1)
  case 13:
    // invoke: _ZNK7QString3argEdici5QChar
    // invoke: QString arg(double, int, char, int, class QChar)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.uchar(args[2].(byte))
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QChar).qclsinst
    if false {fmt.Println(arg4)}
    C._ZNK7QString3argEdici5QChar(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 14:
    // invoke: _ZNK7QString3argEyii5QChar
    // invoke: QString arg(qulonglong, int, int, class QChar)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argEyii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 15:
    // invoke: _ZNK7QString3argEci5QChar
    // invoke: QString arg(char, int, class QChar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argEci5QChar(this.qclsinst, arg0, arg1, arg2)
  case 16:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    var arg6 = args[6].(QString).qclsinst
    if false {fmt.Println(arg6)}
    var arg7 = args[7].(QString).qclsinst
    if false {fmt.Println(arg7)}
    var arg8 = args[8].(QString).qclsinst
    if false {fmt.Println(arg8)}
    C._ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  case 17:
    // invoke: _ZNK7QString3argEmii5QChar
    // invoke: QString arg(ulong, int, int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QChar).qclsinst
    if false {fmt.Println(arg3)}
    C._ZNK7QString3argEmii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 18:
    // invoke: _ZNK7QString3argE5QChariS0_
    // invoke: QString arg(class QChar, int, class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argE5QChariS0_(this.qclsinst, arg0, arg1, arg2)
  case 19:
    // invoke: _ZNK7QString3argERKS_S1_S1_S1_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = args[3].(QString).qclsinst
    if false {fmt.Println(arg3)}
    var arg4 = args[4].(QString).qclsinst
    if false {fmt.Println(arg4)}
    var arg5 = args[5].(QString).qclsinst
    if false {fmt.Println(arg5)}
    C._ZNK7QString3argERKS_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  default:
    qtrt.ErrorResolve("QString", "arg", args)
  }

}

// append(class QChar)
func (this *QString) append(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6appendE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString6appendEPKc
    // invoke: QString & append(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN7QString6appendEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN7QString6appendERK10QByteArray
    // invoke: QString & append(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6appendERK10QByteArray(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN7QString6appendERKS_
    // invoke: QString & append(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6appendERKS_(this.qclsinst, arg0)
  case 4:
    // invoke: _ZN7QString6appendEPK5QChari
    // invoke: QString & append(const class QChar *, int)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6appendEPK5QChari(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "append", args)
  }

}

// capacity()
func (this *QString) capacity(args ...interface{}) () {
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
    C._ZNK7QString8capacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "capacity", args)
  }

}

// toLower()
func (this *QString) toLower(args ...interface{}) () {
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
    C._ZNKR7QString7toLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toLower", args)
  }

}

// toHtmlEscaped()
func (this *QString) toHtmlEscaped(args ...interface{}) () {
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
    C._ZNK7QString13toHtmlEscapedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toHtmlEscaped", args)
  }

}

// isNull()
func (this *QString) isNull(args ...interface{}) () {
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
    C._ZNK7QString6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isNull", args)
  }

}

// localeAwareCompare(const class QString &, const class QString &)
func (this *QString) localeAwareCompare_s(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString18localeAwareCompareERKS_S1_(arg0, arg1)
  case 1:
    // invoke: _ZNK7QString18localeAwareCompareERKS_
    // invoke: int localeAwareCompare(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString18localeAwareCompareERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "localeAwareCompare", args)
  }

}

// fromLocal8Bit(const class QByteArray &)
func (this *QString) fromLocal8Bit_s(args ...interface{}) () {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString13fromLocal8BitERK10QByteArray(arg0)
  case 1:
    // invoke: _ZN7QString13fromLocal8BitEPKci
    // invoke: QString fromLocal8Bit(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString13fromLocal8BitEPKci(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fromLocal8Bit", args)
  }

}

// fromWCharArray(const wchar_t *, int)
func (this *QString) fromWCharArray_s(args ...interface{}) () {
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
    var arg0 = (*C.wchar_t)((unsafe.Pointer)(reflect.ValueOf(args[0].([]rune)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString14fromWCharArrayEPKwi(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fromWCharArray", args)
  }

}

// toStdString()
func (this *QString) toStdString(args ...interface{}) () {
  // toStdString()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString11toStdStringEv
    // invoke: std::string toStdString()
    C._ZNK7QString11toStdStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdString", args)
  }

}

// indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
func (this *QString) indexOf(args ...interface{}) () {
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
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch(this.qclsinst, arg0, arg1, arg2)
  case 1:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioni
    // invoke: int indexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7indexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QString7indexOfER7QRegExpi
    // invoke: int indexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7indexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK7QString7indexOfERK7QRegExpi
    // invoke: int indexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7indexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "indexOf", args)
  }

}

// cbegin()
func (this *QString) cbegin(args ...interface{}) () {
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
    C._ZNK7QString6cbeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "cbegin", args)
  }

}

// toLatin1()
func (this *QString) toLatin1(args ...interface{}) () {
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
    C._ZNKR7QString8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toLatin1", args)
  }

}

// chop(int)
func (this *QString) chop(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QString4chopEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "chop", args)
  }

}

// toStdU32String()
func (this *QString) toStdU32String(args ...interface{}) () {
  // toStdU32String()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString14toStdU32StringEv
    // invoke: std::u32string toStdU32String()
    C._ZNK7QString14toStdU32StringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdU32String", args)
  }

}

// length()
func (this *QString) length(args ...interface{}) () {
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
    C._ZNK7QString6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "length", args)
  }

}

// toDouble(_Bool *)
func (this *QString) toDouble(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK7QString8toDoubleEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "toDouble", args)
  }

}

// isSharedWith(const class QString &)
func (this *QString) isSharedWith(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString12isSharedWithERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "isSharedWith", args)
  }

}

// toUShort(_Bool *, int)
func (this *QString) toUShort(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString8toUShortEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toUShort", args)
  }

}

// constBegin()
func (this *QString) constBegin(args ...interface{}) () {
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
    C._ZNK7QString10constBeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "constBegin", args)
  }

}

// toShort(_Bool *, int)
func (this *QString) toShort(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7toShortEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toShort", args)
  }

}

// number(uint, int)
func (this *QString) number_s(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6numberEji(arg0, arg1)
  case 1:
    // invoke: _ZN7QString6numberEdci
    // invoke: QString number(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN7QString6numberEdci(arg0, arg1, arg2)
  case 2:
    // invoke: _ZN7QString6numberEyi
    // invoke: QString number(qulonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6numberEyi(arg0, arg1)
  case 3:
    // invoke: _ZN7QString6numberEmi
    // invoke: QString number(ulong, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6numberEmi(arg0, arg1)
  case 4:
    // invoke: _ZN7QString6numberExi
    // invoke: QString number(qlonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6numberExi(arg0, arg1)
  case 5:
    // invoke: _ZN7QString6numberEli
    // invoke: QString number(long, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6numberEli(arg0, arg1)
  case 6:
    // invoke: _ZN7QString6numberEii
    // invoke: QString number(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6numberEii(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "number", args)
  }

}

// replace(const class QRegExp &, const class QString &)
func (this *QString) replace(args ...interface{}) () {
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
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString7replaceERK7QRegExpRKS_(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QString7replaceEii5QChar
    // invoke: QString & replace(int, int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN7QString7replaceEii5QChar(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN7QString7replaceERK18QRegularExpressionRKS_
    // invoke: QString & replace(const class QRegularExpression &, const class QString &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString7replaceERK18QRegularExpressionRKS_(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN7QString7replaceEiiPK5QChari
    // invoke: QString & replace(int, int, const class QChar *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    var arg3 = C.int32_t(args[3].(int32))
    if false {fmt.Println(arg3)}
    C._ZN7QString7replaceEiiPK5QChari(this.qclsinst, arg0, arg1, arg2, arg3)
  case 4:
    // invoke: _ZN7QString7replaceEiiRKS_
    // invoke: QString & replace(int, int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN7QString7replaceEiiRKS_(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "replace", args)
  }

}

// push_back(class QChar)
func (this *QString) push_back(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString9push_backE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString9push_backERKS_
    // invoke: void push_back(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString9push_backERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "push_back", args)
  }

}

// simplified()
func (this *QString) simplified(args ...interface{}) () {
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
    C._ZNKR7QString10simplifiedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "simplified", args)
  }

}

// size()
func (this *QString) size(args ...interface{}) () {
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
    C._ZNK7QString4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "size", args)
  }

}

// truncate(int)
func (this *QString) truncate(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QString8truncateEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "truncate", args)
  }

}

// contains(const class QRegularExpression &)
func (this *QString) contains(args ...interface{}) () {
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
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString8containsERK18QRegularExpression(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch
    // invoke: bool contains(const class QRegularExpression &, class QRegularExpressionMatch *)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg1)}
    C._ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QString8containsER7QRegExp
    // invoke: bool contains(class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString8containsER7QRegExp(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK7QString8containsERK7QRegExp
    // invoke: bool contains(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString8containsERK7QRegExp(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "contains", args)
  }

}

// isSimpleText()
func (this *QString) isSimpleText(args ...interface{}) () {
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
    C._ZNK7QString12isSimpleTextEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isSimpleText", args)
  }

}

// cend()
func (this *QString) cend(args ...interface{}) () {
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
    C._ZNK7QString4cendEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "cend", args)
  }

}

// leftJustified(int, class QChar, _Bool)
func (this *QString) leftJustified(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.bool(args[2].(bool))
    if false {fmt.Println(arg2)}
    C._ZNK7QString13leftJustifiedEi5QCharb(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "leftJustified", args)
  }

}

// repeated(int)
func (this *QString) repeated(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QString8repeatedEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "repeated", args)
  }

}

// fromRawData(const class QChar *, int)
func (this *QString) fromRawData_s(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString11fromRawDataEPK5QChari(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fromRawData", args)
  }

}

// squeeze()
func (this *QString) squeeze(args ...interface{}) () {
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
    C._ZN7QString7squeezeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "squeeze", args)
  }

}

// resize(int)
func (this *QString) resize(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZN7QString6resizeEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "resize", args)
  }

}

// toULongLong(_Bool *, int)
func (this *QString) toULongLong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11toULongLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toULongLong", args)
  }

}

// count(const class QRegularExpression &)
func (this *QString) count(args ...interface{}) () {
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
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString5countERK18QRegularExpression(this.qclsinst, arg0)
  case 1:
    // invoke: _ZNK7QString5countEv
    // invoke: int count()
    C._ZNK7QString5countEv(this.qclsinst)
  case 2:
    // invoke: _ZNK7QString5countERK7QRegExp
    // invoke: int count(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString5countERK7QRegExp(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "count", args)
  }

}

// isRightToLeft()
func (this *QString) isRightToLeft(args ...interface{}) () {
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
    C._ZNK7QString13isRightToLeftEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isRightToLeft", args)
  }

}

// toUtf8()
func (this *QString) toUtf8(args ...interface{}) () {
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
    C._ZNKR7QString6toUtf8Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toUtf8", args)
  }

}

// fromUtf8(const char *, int)
func (this *QString) fromUtf8_s(args ...interface{}) () {
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
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString8fromUtf8EPKci(arg0, arg1)
  case 1:
    // invoke: _ZN7QString8fromUtf8ERK10QByteArray
    // invoke: QString fromUtf8(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString8fromUtf8ERK10QByteArray(arg0)
  default:
    qtrt.ErrorResolve("QString", "fromUtf8", args)
  }

}

// toStdU16String()
func (this *QString) toStdU16String(args ...interface{}) () {
  // toStdU16String()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString14toStdU16StringEv
    // invoke: std::u16string toStdU16String()
    C._ZNK7QString14toStdU16StringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toStdU16String", args)
  }

}

// toFloat(_Bool *)
func (this *QString) toFloat(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK7QString7toFloatEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "toFloat", args)
  }

}

// at(int)
func (this *QString) at(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QString2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "at", args)
  }

}

// fill(class QChar, int)
func (this *QString) fill(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString4fillE5QChari(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fill", args)
  }

}

// toULong(_Bool *, int)
func (this *QString) toULong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7toULongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toULong", args)
  }

}

// end()
func (this *QString) end(args ...interface{}) () {
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
    C._ZN7QString3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "end", args)
  }

}

// fromUcs4(const uint *, int)
func (this *QString) fromUcs4_s(args ...interface{}) () {
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
    var arg0 = (*C.int32_t)(args[0].(*int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString8fromUcs4EPKji(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fromUcs4", args)
  }

}

// mid(int, int)
func (this *QString) mid(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString3midEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "mid", args)
  }

}

// prepend(const class QString &)
func (this *QString) prepend(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString7prependERKS_(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString7prependE5QChar
    // invoke: QString & prepend(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString7prependE5QChar(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN7QString7prependEPKc
    // invoke: QString & prepend(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    C._ZN7QString7prependEPKc(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN7QString7prependERK10QByteArray
    // invoke: QString & prepend(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString7prependERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "prepend", args)
  }

}

// constData()
func (this *QString) constData(args ...interface{}) () {
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
    C._ZNK7QString9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "constData", args)
  }

}

// toLongLong(_Bool *, int)
func (this *QString) toLongLong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString10toLongLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toLongLong", args)
  }

}

// setUnicode(const class QChar *, int)
func (this *QString) setUnicode(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString10setUnicodeEPK5QChari(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "setUnicode", args)
  }

}

// fromUtf16(const ushort *, int)
func (this *QString) fromUtf16_s(args ...interface{}) () {
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
    var arg0 = (*C.int16_t)(args[0].(*int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString9fromUtf16EPKti(arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "fromUtf16", args)
  }

}

// isEmpty()
func (this *QString) isEmpty(args ...interface{}) () {
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
    C._ZNK7QString7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isEmpty", args)
  }

}

// setRawData(const class QChar *, int)
func (this *QString) setRawData(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString10setRawDataEPK5QChari(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "setRawData", args)
  }

}

// isDetached()
func (this *QString) isDetached(args ...interface{}) () {
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
    C._ZNK7QString10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isDetached", args)
  }

}

// begin()
func (this *QString) begin(args ...interface{}) () {
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
    C._ZN7QString5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "begin", args)
  }

}

// lastIndexOf(class QRegExp &, int)
func (this *QString) lastIndexOf(args ...interface{}) () {
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
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11lastIndexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK7QString11lastIndexOfERK7QRegExpi
    // invoke: int lastIndexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11lastIndexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioni
    // invoke: int lastIndexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11lastIndexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch
    // invoke: int lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "lastIndexOf", args)
  }

}

// setNum(ushort, int)
func (this *QString) setNum(args ...interface{}) () {
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
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEti(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QString6setNumEyi
    // invoke: QString & setNum(qulonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEyi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN7QString6setNumExi
    // invoke: QString & setNum(qlonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumExi(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZN7QString6setNumEmi
    // invoke: QString & setNum(ulong, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEmi(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN7QString6setNumEsi
    // invoke: QString & setNum(short, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEsi(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN7QString6setNumEli
    // invoke: QString & setNum(long, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEli(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN7QString6setNumEji
    // invoke: QString & setNum(uint, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEji(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN7QString6setNumEii
    // invoke: QString & setNum(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEii(this.qclsinst, arg0, arg1)
  case 8:
    // invoke: _ZN7QString6setNumEdci
    // invoke: QString & setNum(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN7QString6setNumEdci(this.qclsinst, arg0, arg1, arg2)
  case 9:
    // invoke: _ZN7QString6setNumEfci
    // invoke: QString & setNum(float, char, int)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN7QString6setNumEfci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "setNum", args)
  }

}

// setUtf16(const ushort *, int)
func (this *QString) setUtf16(args ...interface{}) () {
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
    var arg0 = (*C.int16_t)(args[0].(*int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString8setUtf16EPKti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "setUtf16", args)
  }

}

// data()
func (this *QString) data(args ...interface{}) () {
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
    C._ZN7QString4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "data", args)
  }

}

// toLong(_Bool *, int)
func (this *QString) toLong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString6toLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toLong", args)
  }

}

// ~QString()
func (this *QString) FreeQString(args ...interface{}) () {
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
    C._ZN7QStringD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "~QString", args)
  }

}

// clear()
func (this *QString) clear(args ...interface{}) () {
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
    C._ZN7QString5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "clear", args)
  }

}

// toInt(_Bool *, int)
func (this *QString) toInt(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString5toIntEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "toInt", args)
  }

}

// toCaseFolded()
func (this *QString) toCaseFolded(args ...interface{}) () {
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
    C._ZNKR7QString12toCaseFoldedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toCaseFolded", args)
  }

}

// toUpper()
func (this *QString) toUpper(args ...interface{}) () {
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
    C._ZNKR7QString7toUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "toUpper", args)
  }

}

// push_front(class QChar)
func (this *QString) push_front(args ...interface{}) () {
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
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString10push_frontE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString10push_frontERKS_
    // invoke: void push_front(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString10push_frontERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "push_front", args)
  }

}

// left(int)
func (this *QString) left(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK7QString4leftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "left", args)
  }

}

// QLatin1String(const class QByteArray &)
func NewQLatin1String(args ...interface{}) QLatin1String {
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
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QLatin1StringC2ERK10QByteArray(qthis, arg0)
  case 1:
    // invoke: _ZN13QLatin1StringC1EPKc
    // invoke: void QLatin1String(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QLatin1StringC2EPKc(qthis, arg0)
  case 2:
    // invoke: _ZN13QLatin1StringC1EPKci
    // invoke: void QLatin1String(const char *, int)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).Pointer()))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN13QLatin1StringC2EPKci(qthis, arg0, arg1)
  default:
    qtrt.ErrorResolve("QLatin1String", "QLatin1String", args)
  }

  return QLatin1String{}
}

// latin1()
func (this *QLatin1String) latin1(args ...interface{}) () {
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
    C._ZNK13QLatin1String6latin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1String", "latin1", args)
  }

}

// data()
func (this *QLatin1String) data(args ...interface{}) () {
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
    C._ZNK13QLatin1String4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1String", "data", args)
  }

}

// size()
func (this *QLatin1String) size(args ...interface{}) () {
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
    C._ZNK13QLatin1String4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1String", "size", args)
  }

}

// isLetterOrNumber()
func (this *QCharRef) isLetterOrNumber(args ...interface{}) () {
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
    C._ZN8QCharRef16isLetterOrNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isLetterOrNumber", args)
  }

}

// isLetter()
func (this *QCharRef) isLetter(args ...interface{}) () {
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
    C._ZNK8QCharRef8isLetterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isLetter", args)
  }

}

// hasMirrored()
func (this *QCharRef) hasMirrored(args ...interface{}) () {
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
    C._ZNK8QCharRef11hasMirroredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "hasMirrored", args)
  }

}

// unicodeVersion()
func (this *QCharRef) unicodeVersion(args ...interface{}) () {
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
    C._ZNK8QCharRef14unicodeVersionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "unicodeVersion", args)
  }

}

// isTitleCase()
func (this *QCharRef) isTitleCase(args ...interface{}) () {
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
    C._ZNK8QCharRef11isTitleCaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isTitleCase", args)
  }

}

// setCell(uchar)
func (this *QCharRef) setCell(args ...interface{}) () {
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
    C._ZN8QCharRef7setCellEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCharRef", "setCell", args)
  }

}

// unicode()
func (this *QCharRef) unicode(args ...interface{}) () {
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
    C._ZN8QCharRef7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "unicode", args)
  }

}

// isSpace()
func (this *QCharRef) isSpace(args ...interface{}) () {
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
    C._ZNK8QCharRef7isSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isSpace", args)
  }

}

// row()
func (this *QCharRef) row(args ...interface{}) () {
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
    C._ZNK8QCharRef3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "row", args)
  }

}

// category()
func (this *QCharRef) category(args ...interface{}) () {
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
    C._ZNK8QCharRef8categoryEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "category", args)
  }

}

// isUpper()
func (this *QCharRef) isUpper(args ...interface{}) () {
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
    C._ZNK8QCharRef7isUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isUpper", args)
  }

}

// toLower()
func (this *QCharRef) toLower(args ...interface{}) () {
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
    C._ZNK8QCharRef7toLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toLower", args)
  }

}

// script()
func (this *QCharRef) script(args ...interface{}) () {
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
    C._ZNK8QCharRef6scriptEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "script", args)
  }

}

// joiningType()
func (this *QCharRef) joiningType(args ...interface{}) () {
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
    C._ZNK8QCharRef11joiningTypeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "joiningType", args)
  }

}

// cell()
func (this *QCharRef) cell(args ...interface{}) () {
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
    C._ZNK8QCharRef4cellEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "cell", args)
  }

}

// digitValue()
func (this *QCharRef) digitValue(args ...interface{}) () {
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
    C._ZNK8QCharRef10digitValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "digitValue", args)
  }

}

// decomposition()
func (this *QCharRef) decomposition(args ...interface{}) () {
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
    C._ZNK8QCharRef13decompositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "decomposition", args)
  }

}

// direction()
func (this *QCharRef) direction(args ...interface{}) () {
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
    C._ZNK8QCharRef9directionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "direction", args)
  }

}

// mirroredChar()
func (this *QCharRef) mirroredChar(args ...interface{}) () {
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
    C._ZNK8QCharRef12mirroredCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "mirroredChar", args)
  }

}

// combiningClass()
func (this *QCharRef) combiningClass(args ...interface{}) () {
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
    C._ZNK8QCharRef14combiningClassEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "combiningClass", args)
  }

}

// decompositionTag()
func (this *QCharRef) decompositionTag(args ...interface{}) () {
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
    C._ZNK8QCharRef16decompositionTagEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "decompositionTag", args)
  }

}

// isMark()
func (this *QCharRef) isMark(args ...interface{}) () {
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
    C._ZNK8QCharRef6isMarkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isMark", args)
  }

}

// isLower()
func (this *QCharRef) isLower(args ...interface{}) () {
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
    C._ZNK8QCharRef7isLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isLower", args)
  }

}

// isNumber()
func (this *QCharRef) isNumber(args ...interface{}) () {
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
    C._ZNK8QCharRef8isNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isNumber", args)
  }

}

// toTitleCase()
func (this *QCharRef) toTitleCase(args ...interface{}) () {
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
    C._ZNK8QCharRef11toTitleCaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toTitleCase", args)
  }

}

// toLatin1()
func (this *QCharRef) toLatin1(args ...interface{}) () {
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
    C._ZNK8QCharRef8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toLatin1", args)
  }

}

// isNull()
func (this *QCharRef) isNull(args ...interface{}) () {
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
    C._ZNK8QCharRef6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isNull", args)
  }

}

// isPrint()
func (this *QCharRef) isPrint(args ...interface{}) () {
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
    C._ZNK8QCharRef7isPrintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isPrint", args)
  }

}

// toUpper()
func (this *QCharRef) toUpper(args ...interface{}) () {
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
    C._ZNK8QCharRef7toUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toUpper", args)
  }

}

// isDigit()
func (this *QCharRef) isDigit(args ...interface{}) () {
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
    C._ZNK8QCharRef7isDigitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isDigit", args)
  }

}

// setRow(uchar)
func (this *QCharRef) setRow(args ...interface{}) () {
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
    C._ZN8QCharRef6setRowEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCharRef", "setRow", args)
  }

}

// isPunct()
func (this *QCharRef) isPunct(args ...interface{}) () {
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
    C._ZNK8QCharRef7isPunctEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isPunct", args)
  }

}

// joining()
func (this *QCharRef) joining(args ...interface{}) () {
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
    C._ZNK8QCharRef7joiningEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "joining", args)
  }

}

// toUcs4()
func (this *QStringRef) toUcs4(args ...interface{}) () {
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
    C._ZNK10QStringRef6toUcs4Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "toUcs4", args)
  }

}

// localeAwareCompare(const class QString &)
func (this *QStringRef) localeAwareCompare(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef18localeAwareCompareERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "localeAwareCompare", args)
  }

}

// toLocal8Bit()
func (this *QStringRef) toLocal8Bit(args ...interface{}) () {
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
    C._ZNK10QStringRef11toLocal8BitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "toLocal8Bit", args)
  }

}

// toUInt(_Bool *, int)
func (this *QStringRef) toUInt(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef6toUIntEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toUInt", args)
  }

}

// toShort(_Bool *, int)
func (this *QStringRef) toShort(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef7toShortEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toShort", args)
  }

}

// toUShort(_Bool *, int)
func (this *QStringRef) toUShort(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef8toUShortEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toUShort", args)
  }

}

// right(int)
func (this *QStringRef) right(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef5rightEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "right", args)
  }

}

// at(int)
func (this *QStringRef) at(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "at", args)
  }

}

// unicode()
func (this *QStringRef) unicode(args ...interface{}) () {
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
    C._ZNK10QStringRef7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "unicode", args)
  }

}

// size()
func (this *QStringRef) size(args ...interface{}) () {
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
    C._ZNK10QStringRef4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "size", args)
  }

}

// length()
func (this *QStringRef) length(args ...interface{}) () {
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
    C._ZNK10QStringRef6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "length", args)
  }

}

// trimmed()
func (this *QStringRef) trimmed(args ...interface{}) () {
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
    C._ZNK10QStringRef7trimmedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "trimmed", args)
  }

}

// mid(int, int)
func (this *QStringRef) mid(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef3midEii(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "mid", args)
  }

}

// constData()
func (this *QStringRef) constData(args ...interface{}) () {
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
    C._ZNK10QStringRef9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "constData", args)
  }

}

// toLongLong(_Bool *, int)
func (this *QStringRef) toLongLong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef10toLongLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toLongLong", args)
  }

}

// appendTo(class QString *)
func (this *QStringRef) appendTo(args ...interface{}) () {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef8appendToEP7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "appendTo", args)
  }

}

// isEmpty()
func (this *QStringRef) isEmpty(args ...interface{}) () {
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
    C._ZNK10QStringRef7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "isEmpty", args)
  }

}

// toString()
func (this *QStringRef) toString(args ...interface{}) () {
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
    C._ZNK10QStringRef8toStringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "toString", args)
  }

}

// cend()
func (this *QStringRef) cend(args ...interface{}) () {
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
    C._ZNK10QStringRef4cendEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "cend", args)
  }

}

// ~QStringRef()
func (this *QStringRef) FreeQStringRef(args ...interface{}) () {
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
    C._ZN10QStringRefD2Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "~QStringRef", args)
  }

}

// begin()
func (this *QStringRef) begin(args ...interface{}) () {
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
    C._ZNK10QStringRef5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "begin", args)
  }

}

// string()
func (this *QStringRef) string(args ...interface{}) () {
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
    C._ZNK10QStringRef6stringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "string", args)
  }

}

// toULong(_Bool *, int)
func (this *QStringRef) toULong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef7toULongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toULong", args)
  }

}

// toFloat(_Bool *)
func (this *QStringRef) toFloat(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef7toFloatEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "toFloat", args)
  }

}

// cbegin()
func (this *QStringRef) cbegin(args ...interface{}) () {
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
    C._ZNK10QStringRef6cbeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "cbegin", args)
  }

}

// end()
func (this *QStringRef) end(args ...interface{}) () {
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
    C._ZNK10QStringRef3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "end", args)
  }

}

// data()
func (this *QStringRef) data(args ...interface{}) () {
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
    C._ZNK10QStringRef4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "data", args)
  }

}

// toULongLong(_Bool *, int)
func (this *QStringRef) toULongLong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef11toULongLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toULongLong", args)
  }

}

// count()
func (this *QStringRef) count(args ...interface{}) () {
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
    C._ZNK10QStringRef5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "count", args)
  }

}

// clear()
func (this *QStringRef) clear(args ...interface{}) () {
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
    C._ZN10QStringRef5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "clear", args)
  }

}

// toInt(_Bool *, int)
func (this *QStringRef) toInt(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef5toIntEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toInt", args)
  }

}

// toLatin1()
func (this *QStringRef) toLatin1(args ...interface{}) () {
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
    C._ZNK10QStringRef8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "toLatin1", args)
  }

}

// isNull()
func (this *QStringRef) isNull(args ...interface{}) () {
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
    C._ZNK10QStringRef6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "isNull", args)
  }

}

// toDouble(_Bool *)
func (this *QStringRef) toDouble(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef8toDoubleEPb(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "toDouble", args)
  }

}

// QStringRef(const class QString *)
func NewQStringRef(args ...interface{}) QStringRef {
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
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QStringRefC2EPK7QString(qthis, arg0)
  case 1:
    // invoke: _ZN10QStringRefC1Ev
    // invoke: void QStringRef()
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QStringRefC2Ev(qthis)
  case 2:
    // invoke: _ZN10QStringRefC1EPK7QStringii
    // invoke: void QStringRef(const class QString *, int, int)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    var qthis = unsafe.Pointer(C.malloc(5))
    if false {reflect.TypeOf(qthis)}
    C._ZN10QStringRefC2EPK7QStringii(qthis, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QStringRef", "QStringRef", args)
  }

  return QStringRef{}
}

// position()
func (this *QStringRef) position(args ...interface{}) () {
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
    C._ZNK10QStringRef8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "position", args)
  }

}

// toLong(_Bool *, int)
func (this *QStringRef) toLong(args ...interface{}) () {
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
    var arg0 = (*C.bool)(args[0].(*bool))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK10QStringRef6toLongEPbi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QStringRef", "toLong", args)
  }

}

// toUtf8()
func (this *QStringRef) toUtf8(args ...interface{}) () {
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
    C._ZNK10QStringRef6toUtf8Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "toUtf8", args)
  }

}

// left(int)
func (this *QStringRef) left(args ...interface{}) () {
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
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    C._ZNK10QStringRef4leftEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "left", args)
  }

}

// <= body block end

