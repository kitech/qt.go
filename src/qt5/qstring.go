package qt5
// auto generated, do not modify.
// created: Sun Jan  3 17:27:54 2016
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
  // proto:  qlonglong QString::toLongLong(bool * ok, int base);
extern void _ZNK7QString10toLongLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  bool QString::isNull();
extern void demth_ZNK7QString6isNullEv(void* qthis);
  // proto:  QString & QString::append(const QChar * uc, int len);
extern void _ZN7QString6appendEPK5QChari(void* qthis, void* arg0, int32_t arg1);
  // proto:  QString & QString::prepend(QChar c);
extern void demth_ZN7QString7prependE5QChar(void* qthis, void* arg0);
  // proto:  QString & QString::insert(int i, QChar c);
extern void _ZN7QString6insertEi5QChar(void* qthis, int32_t arg0, void* arg1);
  // proto:  QString QString::left(int n);
extern void _ZNK7QString4leftEi(void* qthis, int32_t arg0);
  // proto:  void QString::QString(QChar c);
extern void* dector_ZN7QStringC1E5QChar(void* arg0);
extern void _ZN7QStringC1E5QChar(void* qthis, void* arg0);
  // proto:  QString & QString::prepend(const char * s);
extern void demth_ZN7QString7prependEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from);
extern void _ZNK7QString11lastIndexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1);
  // proto: static QString QString::number(int , int base);
extern void _ZN7QString6numberEii(int32_t arg0, int32_t arg1);
  // proto:  void QString::resize(int size);
extern void _ZN7QString6resizeEi(void* qthis, int32_t arg0);
  // proto:  void QString::push_front(QChar c);
extern void demth_ZN7QString10push_frontE5QChar(void* qthis, void* arg0);
  // proto:  void QString::QString();
extern void* dector_ZN7QStringC1Ev();
extern void demth_ZN7QStringC1Ev(void* qthis);
  // proto:  double QString::toDouble(bool * ok);
extern void _ZNK7QString8toDoubleEPb(void* qthis, bool* arg0);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8);
extern void demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7);
  // proto:  QStringRef QString::rightRef(int n);
extern void _ZNK7QString8rightRefEi(void* qthis, int32_t arg0);
  // proto:  QString & QString::setNum(short , int base);
extern void demth_ZN7QString6setNumEsi(void* qthis, int16_t arg0, int32_t arg1);
  // proto:  void QString::QString(const QChar * unicode, int size);
extern void* dector_ZN7QStringC1EPK5QChari(void* arg0, int32_t arg1);
extern void _ZN7QStringC1EPK5QChari(void* qthis, void* arg0, int32_t arg1);
  // proto:  float QString::toFloat(bool * ok);
extern void _ZNK7QString7toFloatEPb(void* qthis, bool* arg0);
  // proto:  int QString::count(const QRegularExpression & re);
extern void _ZNK7QString5countERK18QRegularExpression(void* qthis, void* arg0);
  // proto:  QStringRef QString::midRef(int position, int n);
extern void _ZNK7QString6midRefEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QString::detach();
extern void demth_ZN7QString6detachEv(void* qthis);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4);
extern void demth_ZNK7QString3argERKS_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3);
  // proto:  QString QString::arg(long a, int fieldwidth, int base, QChar fillChar);
extern void demth_ZNK7QString3argElii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  int QString::count();
extern void demth_ZNK7QString5countEv(void* qthis);
  // proto:  QString & QString::setNum(qulonglong , int base);
extern void _ZN7QString6setNumEyi(void* qthis, int64_t arg0, int32_t arg1);
  // proto: static QString QString::fromStdWString(const std::wstring & s);
extern void demth_ZN7QString14fromStdWStringERKi(int32_t* arg0);
  // proto:  void QString::push_back(QChar c);
extern void demth_ZN7QString9push_backE5QChar(void* qthis, void* arg0);
  // proto:  QString & QString::setNum(float , char f, int prec);
extern void demth_ZN7QString6setNumEfci(void* qthis, float arg0, unsigned char arg1, int32_t arg2);
  // proto:  int QString::count(const QRegExp & );
extern void _ZNK7QString5countERK7QRegExp(void* qthis, void* arg0);
  // proto:  int QString::size();
extern void demth_ZNK7QString4sizeEv(void* qthis);
  // proto:  QString & QString::insert(int i, const QChar * uc, int len);
extern void _ZN7QString6insertEiPK5QChari(void* qthis, int32_t arg0, void* arg1, int32_t arg2);
  // proto:  QString & QString::replace(int i, int len, const QChar * s, int slen);
extern void _ZN7QString7replaceEiiPK5QChari(void* qthis, int32_t arg0, int32_t arg1, void* arg2, int32_t arg3);
  // proto: static QString QString::fromRawData(const QChar * , int size);
extern void _ZN7QString11fromRawDataEPK5QChari(void* arg0, int32_t arg1);
  // proto:  QString & QString::insert(int i, const QString & s);
extern void demth_ZN7QString6insertEiRKS_(void* qthis, int32_t arg0, void* arg1);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5);
extern void demth_ZNK7QString3argERKS_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4);
  // proto:  QString & QString::setRawData(const QChar * unicode, int size);
extern void _ZN7QString10setRawDataEPK5QChari(void* qthis, void* arg0, int32_t arg1);
  // proto:  QString & QString::prepend(const QString & s);
extern void demth_ZN7QString7prependERKS_(void* qthis, void* arg0);
  // proto:  ulong QString::toULong(bool * ok, int base);
extern void _ZNK7QString7toULongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  void QString::chop(int n);
extern void _ZN7QString4chopEi(void* qthis, int32_t arg0);
  // proto: static QString QString::fromUtf16(const ushort * , int size);
extern void _ZN7QString9fromUtf16EPKti(int16_t* arg0, int32_t arg1);
  // proto:  bool QString::isDetached();
extern void demth_ZNK7QString10isDetachedEv(void* qthis);
  // proto:  QString & QString::setNum(qlonglong , int base);
extern void _ZN7QString6setNumExi(void* qthis, int64_t arg0, int32_t arg1);
  // proto:  QString QString::mid(int position, int n);
extern void _ZNK7QString3midEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto: static QString QString::fromLocal8Bit(const char * str, int size);
extern void demth_ZN7QString13fromLocal8BitEPKci(unsigned char* arg0, int32_t arg1);
  // proto:  void QString::swap(QString & other);
extern void demth_ZN7QString4swapERS_(void* qthis, void* arg0);
  // proto: static QString QString::fromUtf8(const QByteArray & str);
extern void demth_ZN7QString8fromUtf8ERK10QByteArray(void* arg0);
  // proto: static QString QString::fromUcs4(const char32_t * str, int size);
extern void demth_ZN7QString8fromUcs4EPKDii(char32_t* arg0, int32_t arg1);
  // proto:  QString QString::leftJustified(int width, QChar fill, bool trunc);
extern void _ZNK7QString13leftJustifiedEi5QCharb(void* qthis, int32_t arg0, void* arg1, bool arg2);
  // proto:  int QString::indexOf(const QRegExp & , int from);
extern void _ZNK7QString7indexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1);
  // proto:  void QString::push_back(const QString & s);
extern void demth_ZN7QString9push_backERKS_(void* qthis, void* arg0);
  // proto:  int QString::lastIndexOf(QRegExp & , int from);
extern void _ZNK7QString11lastIndexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3);
extern void demth_ZNK7QString3argERKS_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2);
  // proto:  const ushort * QString::utf16();
extern void _ZNK7QString5utf16Ev(void* qthis);
  // proto:  int QString::toInt(bool * ok, int base);
extern void _ZNK7QString5toIntEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QString QString::arg(double a, int fieldWidth, char fmt, int prec, QChar fillChar);
extern void _ZNK7QString3argEdici5QChar(void* qthis, double arg0, int32_t arg1, unsigned char arg2, int32_t arg3, void* arg4);
  // proto:  QChar * QString::data();
extern void demth_ZN7QString4dataEv(void* qthis);
  // proto:  QString & QString::setNum(uint , int base);
extern void demth_ZN7QString6setNumEji(void* qthis, int32_t arg0, int32_t arg1);
  // proto: static int QString::localeAwareCompare(const QString & s1, const QString & s2);
extern void demth_ZN7QString18localeAwareCompareERKS_S1_(void* arg0, void* arg1);
  // proto:  void QString::QString(const char * ch);
extern void* dector_ZN7QStringC1EPKc(unsigned char* arg0);
extern void demth_ZN7QStringC1EPKc(void* qthis, unsigned char* arg0);
  // proto: static QString QString::fromUtf16(const char16_t * str, int size);
extern void demth_ZN7QString9fromUtf16EPKDsi(char16_t* arg0, int32_t arg1);
  // proto:  QString & QString::replace(const QRegExp & rx, const QString & after);
extern void _ZN7QString7replaceERK7QRegExpRKS_(void* qthis, void* arg0, void* arg1);
  // proto:  QString QString::repeated(int times);
extern void _ZNK7QString8repeatedEi(void* qthis, int32_t arg0);
  // proto:  QString & QString::setUtf16(const ushort * utf16, int size);
extern void demth_ZN7QString8setUtf16EPKti(void* qthis, int16_t* arg0, int32_t arg1);
  // proto: static QString QString::fromStdU32String(const std::u32string & s);
extern void demth_ZN7QString16fromStdU32StringERKi(int32_t* arg0);
  // proto:  void QString::clear();
extern void demth_ZN7QString5clearEv(void* qthis);
  // proto:  bool QString::contains(const QRegExp & rx);
extern void demth_ZNK7QString8containsERK7QRegExp(void* qthis, void* arg0);
  // proto:  bool QString::isSharedWith(const QString & other);
extern void demth_ZNK7QString12isSharedWithERKS_(void* qthis, void* arg0);
  // proto: static QString QString::fromLatin1(const QByteArray & str);
extern void demth_ZN7QString10fromLatin1ERK10QByteArray(void* arg0);
  // proto:  void QString::~QString();
extern void demth_ZN7QStringD0Ev(void* qthis);
  // proto:  QString & QString::remove(const QRegularExpression & re);
extern void demth_ZN7QString6removeERK18QRegularExpression(void* qthis, void* arg0);
  // proto:  QString & QString::setNum(int , int base);
extern void demth_ZN7QString6setNumEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QString QString::toHtmlEscaped();
extern void _ZNK7QString13toHtmlEscapedEv(void* qthis);
  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from, QRegularExpressionMatch * rmatch);
extern void _ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  QString & QString::append(const QByteArray & s);
extern void demth_ZN7QString6appendERK10QByteArray(void* qthis, void* arg0);
  // proto: static QString QString::fromLatin1(const char * str, int size);
extern void demth_ZN7QString10fromLatin1EPKci(unsigned char* arg0, int32_t arg1);
  // proto:  bool QString::contains(const QRegularExpression & re, QRegularExpressionMatch * match);
extern void _ZNK7QString8containsERK18QRegularExpressionP23QRegularExpressionMatch(void* qthis, void* arg0, void* arg1);
  // proto:  int QString::indexOf(const QRegularExpression & re, int from, QRegularExpressionMatch * rmatch);
extern void _ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  int QString::lastIndexOf(const QRegExp & , int from);
extern void _ZNK7QString11lastIndexOfERK7QRegExpi(void* qthis, void* arg0, int32_t arg1);
  // proto:  int QString::toWCharArray(wchar_t * array);
extern void demth_ZNK7QString12toWCharArrayEPw(void* qthis, wchar_t* arg0);
  // proto:  QString & QString::prepend(const QByteArray & s);
extern void demth_ZN7QString7prependERK10QByteArray(void* qthis, void* arg0);
  // proto:  QString & QString::replace(int i, int len, const QString & after);
extern void _ZN7QString7replaceEiiRKS_(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto: static QString QString::fromStdString(const std::string & s);
extern void demth_ZN7QString13fromStdStringERKi(int32_t* arg0);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7);
extern void demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6);
  // proto: static QString QString::fromWCharArray(const wchar_t * string, int size);
extern void demth_ZN7QString14fromWCharArrayEPKwi(wchar_t* arg0, int32_t arg1);
  // proto:  QString & QString::fill(QChar c, int size);
extern void _ZN7QString4fillE5QChari(void* qthis, void* arg0, int32_t arg1);
  // proto:  const QChar * QString::constData();
extern void demth_ZNK7QString9constDataEv(void* qthis);
  // proto: static QString QString::number(ulong , int base);
extern void _ZN7QString6numberEmi(int32_t arg0, int32_t arg1);
  // proto:  long QString::toLong(bool * ok, int base);
extern void _ZNK7QString6toLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  int QString::length();
extern void demth_ZNK7QString6lengthEv(void* qthis);
  // proto: static QString QString::fromUtf8(const char * str, int size);
extern void demth_ZN7QString8fromUtf8EPKci(unsigned char* arg0, int32_t arg1);
  // proto: static QString QString::number(qlonglong , int base);
extern void _ZN7QString6numberExi(int64_t arg0, int32_t arg1);
  // proto:  QStringRef QString::leftRef(int n);
extern void _ZNK7QString7leftRefEi(void* qthis, int32_t arg0);
  // proto:  QString & QString::setNum(long , int base);
extern void demth_ZN7QString6setNumEli(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  QString QString::arg(const QString & a1, const QString & a2);
extern void demth_ZNK7QString3argERKS_S1_(void* qthis, void* arg0, void* arg1);
  // proto:  bool QString::isSimpleText();
extern void _ZNK7QString12isSimpleTextEv(void* qthis);
  // proto: static QString QString::fromUcs4(const uint * , int size);
extern void _ZN7QString8fromUcs4EPKji(int32_t* arg0, int32_t arg1);
  // proto:  QString & QString::setUnicode(const QChar * unicode, int size);
extern void _ZN7QString10setUnicodeEPK5QChari(void* qthis, void* arg0, int32_t arg1);
  // proto:  bool QString::contains(QRegExp & rx);
extern void demth_ZNK7QString8containsER7QRegExp(void* qthis, void* arg0);
  // proto:  const QChar * QString::unicode();
extern void demth_ZNK7QString7unicodeEv(void* qthis);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8, const QString & a9);
extern void demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8);
  // proto:  int QString::indexOf(const QRegularExpression & re, int from);
extern void _ZNK7QString7indexOfERK18QRegularExpressioni(void* qthis, void* arg0, int32_t arg1);
  // proto: static QString QString::number(long , int base);
extern void _ZN7QString6numberEli(int32_t arg0, int32_t arg1);
  // proto: static QString QString::number(uint , int base);
extern void _ZN7QString6numberEji(int32_t arg0, int32_t arg1);
  // proto: static QString QString::fromLocal8Bit(const QByteArray & str);
extern void demth_ZN7QString13fromLocal8BitERK10QByteArray(void* arg0);
  // proto:  const QChar QString::at(int i);
extern void demth_ZNK7QString2atEi(void* qthis, int32_t arg0);
  // proto:  void QString::QString(int size, QChar c);
extern void* dector_ZN7QStringC1Ei5QChar(int32_t arg0, void* arg1);
extern void _ZN7QStringC1Ei5QChar(void* qthis, int32_t arg0, void* arg1);
  // proto:  QString & QString::setNum(ulong , int base);
extern void demth_ZN7QString6setNumEmi(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  void QString::push_front(const QString & s);
extern void demth_ZN7QString10push_frontERKS_(void* qthis, void* arg0);
  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6);
extern void demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_(void* qthis, void* arg0, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5);
  // proto: static QString QString::number(double , char f, int prec);
extern void _ZN7QString6numberEdci(double arg0, unsigned char arg1, int32_t arg2);
  // proto:  QString & QString::append(QChar c);
extern void _ZN7QString6appendE5QChar(void* qthis, void* arg0);
  // proto:  uint QString::toUInt(bool * ok, int base);
extern void _ZNK7QString6toUIntEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QString & QString::append(const QString & s);
extern void _ZN7QString6appendERKS_(void* qthis, void* arg0);
  // proto: static QString QString::fromStdU16String(const std::u16string & s);
extern void demth_ZN7QString16fromStdU16StringERKi(int32_t* arg0);
  // proto:  QString QString::arg(qlonglong a, int fieldwidth, int base, QChar fillChar);
extern void _ZNK7QString3argExii5QChar(void* qthis, int64_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  ushort QString::toUShort(bool * ok, int base);
extern void _ZNK7QString8toUShortEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QString QString::arg(uint a, int fieldWidth, int base, QChar fillChar);
extern void demth_ZNK7QString3argEjii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  QString & QString::setNum(ushort , int base);
extern void demth_ZN7QString6setNumEti(void* qthis, int16_t arg0, int32_t arg1);
  // proto:  QString & QString::replace(const QRegularExpression & re, const QString & after);
extern void _ZN7QString7replaceERK18QRegularExpressionRKS_(void* qthis, void* arg0, void* arg1);
  // proto:  QString & QString::setNum(double , char f, int prec);
extern void _ZN7QString6setNumEdci(void* qthis, double arg0, unsigned char arg1, int32_t arg2);
  // proto: static QString QString::number(qulonglong , int base);
extern void _ZN7QString6numberEyi(int64_t arg0, int32_t arg1);
  // proto:  QString QString::arg(ushort a, int fieldWidth, int base, QChar fillChar);
extern void demth_ZNK7QString3argEtii5QChar(void* qthis, int16_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  void QString::QString(const QString & );
extern void* dector_ZN7QStringC1ERKS_(void* arg0);
extern void demth_ZN7QStringC1ERKS_(void* qthis, void* arg0);
  // proto:  QString QString::arg(short a, int fieldWidth, int base, QChar fillChar);
extern void demth_ZNK7QString3argEsii5QChar(void* qthis, int16_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  void QString::QString(const QByteArray & a);
extern void* dector_ZN7QStringC1ERK10QByteArray(void* arg0);
extern void demth_ZN7QStringC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  qulonglong QString::toULongLong(bool * ok, int base);
extern void _ZNK7QString11toULongLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QString & QString::append(const char * s);
extern void demth_ZN7QString6appendEPKc(void* qthis, unsigned char* arg0);
  // proto:  int QString::capacity();
extern void demth_ZNK7QString8capacityEv(void* qthis);
  // proto:  void QString::squeeze();
extern void demth_ZN7QString7squeezeEv(void* qthis);
  // proto:  void QString::truncate(int pos);
extern void _ZN7QString8truncateEi(void* qthis, int32_t arg0);
  // proto:  QString QString::arg(int a, int fieldWidth, int base, QChar fillChar);
extern void demth_ZNK7QString3argEiii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  QString QString::arg(QChar a, int fieldWidth, QChar fillChar);
extern void _ZNK7QString3argE5QChariS0_(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  int QString::localeAwareCompare(const QString & s);
extern void _ZNK7QString18localeAwareCompareERKS_(void* qthis, void* arg0);
  // proto:  QString & QString::remove(const QRegExp & rx);
extern void demth_ZN7QString6removeERK7QRegExp(void* qthis, void* arg0);
  // proto:  bool QString::contains(const QRegularExpression & re);
extern void _ZNK7QString8containsERK18QRegularExpression(void* qthis, void* arg0);
  // proto:  int QString::indexOf(QRegExp & , int from);
extern void _ZNK7QString7indexOfER7QRegExpi(void* qthis, void* arg0, int32_t arg1);
  // proto:  QString & QString::replace(int i, int len, QChar after);
extern void _ZN7QString7replaceEii5QChar(void* qthis, int32_t arg0, int32_t arg1, void* arg2);
  // proto:  bool QString::isRightToLeft();
extern void _ZNK7QString13isRightToLeftEv(void* qthis);
  // proto:  QString QString::arg(char a, int fieldWidth, QChar fillChar);
extern void _ZNK7QString3argEci5QChar(void* qthis, unsigned char arg0, int32_t arg1, void* arg2);
  // proto:  QVector<uint> QString::toUcs4();
extern void _ZNK7QString6toUcs4Ev(void* qthis);
  // proto:  QString & QString::remove(int i, int len);
extern void _ZN7QString6removeEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  bool QString::isEmpty();
extern void demth_ZNK7QString7isEmptyEv(void* qthis);
  // proto:  QString QString::right(int n);
extern void _ZNK7QString5rightEi(void* qthis, int32_t arg0);
  // proto:  QString QString::rightJustified(int width, QChar fill, bool trunc);
extern void _ZNK7QString14rightJustifiedEi5QCharb(void* qthis, int32_t arg0, void* arg1, bool arg2);
  // proto:  QString QString::arg(const QString & a, int fieldWidth, QChar fillChar);
extern void _ZNK7QString3argERKS_i5QChar(void* qthis, void* arg0, int32_t arg1, void* arg2);
  // proto:  QString QString::arg(qulonglong a, int fieldwidth, int base, QChar fillChar);
extern void _ZNK7QString3argEyii5QChar(void* qthis, int64_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  void QString::reserve(int size);
extern void demth_ZN7QString7reserveEi(void* qthis, int32_t arg0);
  // proto:  short QString::toShort(bool * ok, int base);
extern void _ZNK7QString7toShortEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QString QString::arg(ulong a, int fieldwidth, int base, QChar fillChar);
extern void demth_ZNK7QString3argEmii5QChar(void* qthis, int32_t arg0, int32_t arg1, int32_t arg2, void* arg3);
  // proto:  const char * QLatin1String::data();
extern void demth_ZNK13QLatin1String4dataEv(void* qthis);
  // proto:  void QLatin1String::QLatin1String(const char * s);
extern void* dector_ZN13QLatin1StringC1EPKc(unsigned char* arg0);
extern void _ZN13QLatin1StringC1EPKc(void* qthis, unsigned char* arg0);
  // proto:  int QLatin1String::size();
extern void demth_ZNK13QLatin1String4sizeEv(void* qthis);
  // proto:  void QLatin1String::QLatin1String(const QByteArray & s);
extern void* dector_ZN13QLatin1StringC1ERK10QByteArray(void* arg0);
extern void demth_ZN13QLatin1StringC1ERK10QByteArray(void* qthis, void* arg0);
  // proto:  const char * QLatin1String::latin1();
extern void demth_ZNK13QLatin1String6latin1Ev(void* qthis);
  // proto:  void QLatin1String::QLatin1String(const char * s, int sz);
extern void* dector_ZN13QLatin1StringC1EPKci(unsigned char* arg0, int32_t arg1);
extern void _ZN13QLatin1StringC1EPKci(void* qthis, unsigned char* arg0, int32_t arg1);
  // proto:  bool QCharRef::isLetterOrNumber();
extern void demth_ZN8QCharRef16isLetterOrNumberEv(void* qthis);
  // proto:  bool QCharRef::isDigit();
extern void demth_ZNK8QCharRef7isDigitEv(void* qthis);
  // proto:  char QCharRef::toLatin1();
extern void demth_ZNK8QCharRef8toLatin1Ev(void* qthis);
  // proto:  void QCharRef::setCell(uchar cell);
extern void demth_ZN8QCharRef7setCellEh(void* qthis, unsigned char arg0);
  // proto:  bool QCharRef::isMark();
extern void demth_ZNK8QCharRef6isMarkEv(void* qthis);
  // proto:  void QCharRef::QCharRef(QString & str, int idx);
extern void* dector_ZN8QCharRefC1ER7QStringi(void* arg0, int32_t arg1);
extern void demth_ZN8QCharRefC1ER7QStringi(void* qthis, void* arg0, int32_t arg1);
  // proto:  int QCharRef::digitValue();
extern void demth_ZNK8QCharRef10digitValueEv(void* qthis);
  // proto:  bool QCharRef::isLetter();
extern void demth_ZNK8QCharRef8isLetterEv(void* qthis);
  // proto:  bool QCharRef::isNumber();
extern void demth_ZNK8QCharRef8isNumberEv(void* qthis);
  // proto:  bool QCharRef::isPrint();
extern void demth_ZNK8QCharRef7isPrintEv(void* qthis);
  // proto:  QChar QCharRef::toLower();
extern void demth_ZNK8QCharRef7toLowerEv(void* qthis);
  // proto:  void QCharRef::setRow(uchar row);
extern void demth_ZN8QCharRef6setRowEh(void* qthis, unsigned char arg0);
  // proto:  bool QCharRef::isNull();
extern void demth_ZNK8QCharRef6isNullEv(void* qthis);
  // proto:  QChar QCharRef::toTitleCase();
extern void demth_ZNK8QCharRef11toTitleCaseEv(void* qthis);
  // proto:  bool QCharRef::hasMirrored();
extern void demth_ZNK8QCharRef11hasMirroredEv(void* qthis);
  // proto:  uchar QCharRef::row();
extern void demth_ZNK8QCharRef3rowEv(void* qthis);
  // proto:  ushort & QCharRef::unicode();
extern void demth_ZN8QCharRef7unicodeEv(void* qthis);
  // proto:  bool QCharRef::isTitleCase();
extern void demth_ZNK8QCharRef11isTitleCaseEv(void* qthis);
  // proto:  bool QCharRef::isUpper();
extern void demth_ZNK8QCharRef7isUpperEv(void* qthis);
  // proto:  uchar QCharRef::cell();
extern void demth_ZNK8QCharRef4cellEv(void* qthis);
  // proto:  QString QCharRef::decomposition();
extern void demth_ZNK8QCharRef13decompositionEv(void* qthis);
  // proto:  uchar QCharRef::combiningClass();
extern void demth_ZNK8QCharRef14combiningClassEv(void* qthis);
  // proto:  QChar QCharRef::mirroredChar();
extern void demth_ZNK8QCharRef12mirroredCharEv(void* qthis);
  // proto:  bool QCharRef::isSpace();
extern void demth_ZNK8QCharRef7isSpaceEv(void* qthis);
  // proto:  bool QCharRef::isPunct();
extern void demth_ZNK8QCharRef7isPunctEv(void* qthis);
  // proto:  QChar QCharRef::toUpper();
extern void demth_ZNK8QCharRef7toUpperEv(void* qthis);
  // proto:  bool QCharRef::isLower();
extern void demth_ZNK8QCharRef7isLowerEv(void* qthis);
  // proto:  short QStringRef::toShort(bool * ok, int base);
extern void _ZNK10QStringRef7toShortEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  void QStringRef::QStringRef(const QString * string);
extern void* dector_ZN10QStringRefC1EPK7QString(void* arg0);
extern void demth_ZN10QStringRefC1EPK7QString(void* qthis, void* arg0);
  // proto:  qulonglong QStringRef::toULongLong(bool * ok, int base);
extern void _ZNK10QStringRef11toULongLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  void QStringRef::clear();
extern void demth_ZN10QStringRef5clearEv(void* qthis);
  // proto:  int QStringRef::position();
extern void demth_ZNK10QStringRef8positionEv(void* qthis);
  // proto:  long QStringRef::toLong(bool * ok, int base);
extern void _ZNK10QStringRef6toLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  const QChar * QStringRef::cbegin();
extern void demth_ZNK10QStringRef6cbeginEv(void* qthis);
  // proto:  ushort QStringRef::toUShort(bool * ok, int base);
extern void _ZNK10QStringRef8toUShortEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  uint QStringRef::toUInt(bool * ok, int base);
extern void _ZNK10QStringRef6toUIntEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  bool QStringRef::isEmpty();
extern void demth_ZNK10QStringRef7isEmptyEv(void* qthis);
  // proto:  int QStringRef::localeAwareCompare(const QString & s);
extern void demth_ZNK10QStringRef18localeAwareCompareERK7QString(void* qthis, void* arg0);
  // proto:  QByteArray QStringRef::toUtf8();
extern void _ZNK10QStringRef6toUtf8Ev(void* qthis);
  // proto:  int QStringRef::size();
extern void demth_ZNK10QStringRef4sizeEv(void* qthis);
  // proto:  const QChar * QStringRef::constData();
extern void demth_ZNK10QStringRef9constDataEv(void* qthis);
  // proto:  QStringRef QStringRef::left(int n);
extern void _ZNK10QStringRef4leftEi(void* qthis, int32_t arg0);
  // proto:  QVector<uint> QStringRef::toUcs4();
extern void _ZNK10QStringRef6toUcs4Ev(void* qthis);
  // proto:  int QStringRef::count();
extern void demth_ZNK10QStringRef5countEv(void* qthis);
  // proto:  void QStringRef::QStringRef(const QString * string, int position, int size);
extern void* dector_ZN10QStringRefC1EPK7QStringii(void* arg0, int32_t arg1, int32_t arg2);
extern void demth_ZN10QStringRefC1EPK7QStringii(void* qthis, void* arg0, int32_t arg1, int32_t arg2);
  // proto:  void QStringRef::QStringRef();
extern void* dector_ZN10QStringRefC1Ev();
extern void demth_ZN10QStringRefC1Ev(void* qthis);
  // proto:  QStringRef QStringRef::right(int n);
extern void _ZNK10QStringRef5rightEi(void* qthis, int32_t arg0);
  // proto:  const QChar QStringRef::at(int i);
extern void demth_ZNK10QStringRef2atEi(void* qthis, int32_t arg0);
  // proto:  double QStringRef::toDouble(bool * ok);
extern void _ZNK10QStringRef8toDoubleEPb(void* qthis, bool* arg0);
  // proto:  bool QStringRef::isNull();
extern void demth_ZNK10QStringRef6isNullEv(void* qthis);
  // proto:  const QChar * QStringRef::data();
extern void demth_ZNK10QStringRef4dataEv(void* qthis);
  // proto:  qlonglong QStringRef::toLongLong(bool * ok, int base);
extern void _ZNK10QStringRef10toLongLongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  QByteArray QStringRef::toLatin1();
extern void _ZNK10QStringRef8toLatin1Ev(void* qthis);
  // proto:  const QChar * QStringRef::begin();
extern void demth_ZNK10QStringRef5beginEv(void* qthis);
  // proto:  const QChar * QStringRef::unicode();
extern void demth_ZNK10QStringRef7unicodeEv(void* qthis);
  // proto:  QStringRef QStringRef::mid(int pos, int n);
extern void _ZNK10QStringRef3midEii(void* qthis, int32_t arg0, int32_t arg1);
  // proto:  float QStringRef::toFloat(bool * ok);
extern void _ZNK10QStringRef7toFloatEPb(void* qthis, bool* arg0);
  // proto:  const QString * QStringRef::string();
extern void demth_ZNK10QStringRef6stringEv(void* qthis);
  // proto:  QString QStringRef::toString();
extern void _ZNK10QStringRef8toStringEv(void* qthis);
  // proto:  QStringRef QStringRef::trimmed();
extern void _ZNK10QStringRef7trimmedEv(void* qthis);
  // proto:  int QStringRef::toInt(bool * ok, int base);
extern void _ZNK10QStringRef5toIntEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  const QChar * QStringRef::cend();
extern void demth_ZNK10QStringRef4cendEv(void* qthis);
  // proto:  QStringRef QStringRef::appendTo(QString * string);
extern void _ZNK10QStringRef8appendToEP7QString(void* qthis, void* arg0);
  // proto:  int QStringRef::length();
extern void demth_ZNK10QStringRef6lengthEv(void* qthis);
  // proto:  void QStringRef::~QStringRef();
extern void demth_ZN10QStringRefD0Ev(void* qthis);
  // proto:  QByteArray QStringRef::toLocal8Bit();
extern void _ZNK10QStringRef11toLocal8BitEv(void* qthis);
  // proto:  ulong QStringRef::toULong(bool * ok, int base);
extern void _ZNK10QStringRef7toULongEPbi(void* qthis, bool* arg0, int32_t arg1);
  // proto:  const QChar * QStringRef::end();
extern void demth_ZNK10QStringRef3endEv(void* qthis);
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

  // proto:  qlonglong QString::toLongLong(bool * ok, int base);
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

  // proto:  bool QString::isNull();
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
    C.demth_ZNK7QString6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isNull", args)
  }

}

  // proto:  QString & QString::append(const QChar * uc, int len);
func (this *QString) append(args ...interface{}) () {
  // append(const class QChar *, int)
  // append(const class QByteArray &)
  // append(const class QStringRef &)
  // append(class QChar)
  // append(const class QString &)
  // append(const char *)
  // append(class QLatin1String)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.ByteTy(true) // "const char *"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6appendEPK5QChari
    // invoke: QString & append(const class QChar *, int)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6appendEPK5QChari(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QString6appendERK10QByteArray
    // invoke: QString & append(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString6appendERK10QByteArray(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN7QString6appendE5QChar
    // invoke: QString & append(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6appendE5QChar(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN7QString6appendERKS_
    // invoke: QString & append(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C._ZN7QString6appendERKS_(this.qclsinst, arg0)
  case 4:
    // invoke: _ZN7QString6appendEPKc
    // invoke: QString & append(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString6appendEPKc(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "append", args)
  }

}

  // proto:  QString & QString::prepend(QChar c);
func (this *QString) prepend(args ...interface{}) () {
  // prepend(class QLatin1String)
  // prepend(class QChar)
  // prepend(const char *)
  // prepend(const class QString &)
  // prepend(const class QByteArray &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.ByteTy(true) // "const char *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QByteArray{}) // "const QByteArray &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString7prependE5QChar
    // invoke: QString & prepend(class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString7prependE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString7prependEPKc
    // invoke: QString & prepend(const char *)
    var arg0 = (*C.uchar)((unsafe.Pointer)(reflect.ValueOf(args[0].([]byte)).UnsafeAddr()))
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString7prependEPKc(this.qclsinst, arg0)
  case 2:
    // invoke: _ZN7QString7prependERKS_
    // invoke: QString & prepend(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString7prependERKS_(this.qclsinst, arg0)
  case 3:
    // invoke: _ZN7QString7prependERK10QByteArray
    // invoke: QString & prepend(const class QByteArray &)
    var arg0 = args[0].(QByteArray).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString7prependERK10QByteArray(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "prepend", args)
  }

}

  // proto:  QString & QString::insert(int i, QChar c);
func (this *QString) insert(args ...interface{}) () {
  // insert(int, class QChar)
  // insert(int, const class QChar *, int)
  // insert(int, const class QString &)
  // insert(int, class QLatin1String)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int32Ty(false) // "int"
  vtys[0][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][2] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "int"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"

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
    // invoke: _ZN7QString6insertEiPK5QChari
    // invoke: QString & insert(int, const class QChar *, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QChar).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN7QString6insertEiPK5QChari(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZN7QString6insertEiRKS_
    // invoke: QString & insert(int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6insertEiRKS_(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "insert", args)
  }

}

  // proto:  QString QString::left(int n);
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

  // proto:  void QString::QString(QChar c);
func NewQString(args ...interface{}) QString {
  return QString{}
}

  // proto:  int QString::lastIndexOf(const QRegularExpression & re, int from);
func (this *QString) lastIndexOf(args ...interface{}) () {
  // lastIndexOf(const class QRegularExpression &, int)
  // lastIndexOf(class QLatin1String, int, Qt::CaseSensitivity)
  // lastIndexOf(class QRegExp &, int)
  // lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
  // lastIndexOf(const class QRegExp &, int)
  // lastIndexOf(const class QStringRef &, int, Qt::CaseSensitivity)
  // lastIndexOf(const class QString &, int, Qt::CaseSensitivity)
  // lastIndexOf(class QChar, int, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioni
    // invoke: int lastIndexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11lastIndexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK7QString11lastIndexOfER7QRegExpi
    // invoke: int lastIndexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11lastIndexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch
    // invoke: int lastIndexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString11lastIndexOfERK18QRegularExpressioniP23QRegularExpressionMatch(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZNK7QString11lastIndexOfERK7QRegExpi
    // invoke: int lastIndexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString11lastIndexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "lastIndexOf", args)
  }

}

  // proto: static QString QString::number(int , int base);
func (this *QString) number_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "number", args)
  }

}

  // proto:  void QString::resize(int size);
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

  // proto:  void QString::push_front(QChar c);
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
    C.demth_ZN7QString10push_frontE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString10push_frontERKS_
    // invoke: void push_front(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString10push_frontERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "push_front", args)
  }

}

  // proto:  double QString::toDouble(bool * ok);
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

  // proto:  QString QString::arg(const QString & a1, const QString & a2, const QString & a3, const QString & a4, const QString & a5, const QString & a6, const QString & a7, const QString & a8);
func (this *QString) arg(args ...interface{}) () {
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(long, int, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &)
  // arg(double, int, char, int, class QChar)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(const class QString &, const class QString &, const class QString &, const class QString &, const class QString &, const class QString &)
  // arg(qlonglong, int, int, class QChar)
  // arg(uint, int, int, class QChar)
  // arg(ushort, int, int, class QChar)
  // arg(short, int, int, class QChar)
  // arg(int, int, int, class QChar)
  // arg(class QChar, int, class QChar)
  // arg(char, int, class QChar)
  // arg(const class QString &, int, class QChar)
  // arg(qulonglong, int, int, class QChar)
  // arg(ulong, int, int, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][7] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.Int32Ty(false) // "long"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[2][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.DoubleTy(false) // "double"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.ByteTy(false) // "char"
  vtys[5][3] = qtrt.Int32Ty(false) // "int"
  vtys[5][4] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][6] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][7] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8][8] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][3] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][4] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[9][5] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[10][1] = qtrt.Int32Ty(false) // "int"
  vtys[10][2] = qtrt.Int32Ty(false) // "int"
  vtys[10][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = qtrt.Int32Ty(false) // "uint"
  vtys[11][1] = qtrt.Int32Ty(false) // "int"
  vtys[11][2] = qtrt.Int32Ty(false) // "int"
  vtys[11][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[12][1] = qtrt.Int32Ty(false) // "int"
  vtys[12][2] = qtrt.Int32Ty(false) // "int"
  vtys[12][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[13] = make(map[int32]reflect.Type)
  vtys[13][0] = qtrt.Int16Ty(false) // "short"
  vtys[13][1] = qtrt.Int32Ty(false) // "int"
  vtys[13][2] = qtrt.Int32Ty(false) // "int"
  vtys[13][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[14] = make(map[int32]reflect.Type)
  vtys[14][0] = qtrt.Int32Ty(false) // "int"
  vtys[14][1] = qtrt.Int32Ty(false) // "int"
  vtys[14][2] = qtrt.Int32Ty(false) // "int"
  vtys[14][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[15] = make(map[int32]reflect.Type)
  vtys[15][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[15][1] = qtrt.Int32Ty(false) // "int"
  vtys[15][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[16] = make(map[int32]reflect.Type)
  vtys[16][0] = qtrt.ByteTy(false) // "char"
  vtys[16][1] = qtrt.Int32Ty(false) // "int"
  vtys[16][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[17] = make(map[int32]reflect.Type)
  vtys[17][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[17][1] = qtrt.Int32Ty(false) // "int"
  vtys[17][2] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[18] = make(map[int32]reflect.Type)
  vtys[18][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[18][1] = qtrt.Int32Ty(false) // "int"
  vtys[18][2] = qtrt.Int32Ty(false) // "int"
  vtys[18][3] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[19] = make(map[int32]reflect.Type)
  vtys[19][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[19][1] = qtrt.Int32Ty(false) // "int"
  vtys[19][2] = qtrt.Int32Ty(false) // "int"
  vtys[19][3] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
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
    C.demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
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
    C.demth_ZNK7QString3argERKS_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3)
  case 2:
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
    C.demth_ZNK7QString3argElii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 3:
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
    C.demth_ZNK7QString3argERKS_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4)
  case 4:
    // invoke: _ZNK7QString3argERKS_S1_S1_
    // invoke: QString arg(const class QString &, const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C.demth_ZNK7QString3argERKS_S1_S1_(this.qclsinst, arg0, arg1, arg2)
  case 5:
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
  case 6:
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
    C.demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6)
  case 7:
    // invoke: _ZNK7QString3argERKS_S1_
    // invoke: QString arg(const class QString &, const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C.demth_ZNK7QString3argERKS_S1_(this.qclsinst, arg0, arg1)
  case 8:
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
    C.demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
  case 9:
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
    C.demth_ZNK7QString3argERKS_S1_S1_S1_S1_S1_(this.qclsinst, arg0, arg1, arg2, arg3, arg4, arg5)
  case 10:
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
    C.demth_ZNK7QString3argEjii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 12:
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
    C.demth_ZNK7QString3argEtii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 13:
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
    C.demth_ZNK7QString3argEsii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 14:
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
    C.demth_ZNK7QString3argEiii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  case 15:
    // invoke: _ZNK7QString3argE5QChariS0_
    // invoke: QString arg(class QChar, int, class QChar)
    var arg0 = args[0].(QChar).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argE5QChariS0_(this.qclsinst, arg0, arg1, arg2)
  case 16:
    // invoke: _ZNK7QString3argEci5QChar
    // invoke: QString arg(char, int, class QChar)
    var arg0 = C.uchar(args[0].(byte))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argEci5QChar(this.qclsinst, arg0, arg1, arg2)
  case 17:
    // invoke: _ZNK7QString3argERKS_i5QChar
    // invoke: QString arg(const class QString &, int, class QChar)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString3argERKS_i5QChar(this.qclsinst, arg0, arg1, arg2)
  case 18:
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
  case 19:
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
    C.demth_ZNK7QString3argEmii5QChar(this.qclsinst, arg0, arg1, arg2, arg3)
  default:
    qtrt.ErrorResolve("QString", "arg", args)
  }

}

  // proto:  QStringRef QString::rightRef(int n);
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

  // proto:  QString & QString::setNum(short , int base);
func (this *QString) setNum(args ...interface{}) () {
  // setNum(short, int)
  // setNum(qulonglong, int)
  // setNum(float, char, int)
  // setNum(qlonglong, int)
  // setNum(uint, int)
  // setNum(int, int)
  // setNum(long, int)
  // setNum(ulong, int)
  // setNum(ushort, int)
  // setNum(double, char, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.Int16Ty(false) // "short"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int64Ty(false) // "qulonglong"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = qtrt.FloatTy(false) // "float"
  vtys[2][1] = qtrt.ByteTy(false) // "char"
  vtys[2][2] = qtrt.Int32Ty(false) // "int"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int64Ty(false) // "qlonglong"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = qtrt.Int32Ty(false) // "uint"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = qtrt.Int32Ty(false) // "int"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = qtrt.Int32Ty(false) // "long"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "ulong"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = qtrt.Int16Ty(false) // "ushort"
  vtys[8][1] = qtrt.Int32Ty(false) // "int"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = qtrt.DoubleTy(false) // "double"
  vtys[9][1] = qtrt.ByteTy(false) // "char"
  vtys[9][2] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6setNumEsi
    // invoke: QString & setNum(short, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6setNumEsi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZN7QString6setNumEyi
    // invoke: QString & setNum(qulonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumEyi(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN7QString6setNumEfci
    // invoke: QString & setNum(float, char, int)
    var arg0 = C.float(args[0].(float32))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C.demth_ZN7QString6setNumEfci(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN7QString6setNumExi
    // invoke: QString & setNum(qlonglong, int)
    var arg0 = C.int64_t(args[0].(int64))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZN7QString6setNumExi(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN7QString6setNumEji
    // invoke: QString & setNum(uint, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6setNumEji(this.qclsinst, arg0, arg1)
  case 5:
    // invoke: _ZN7QString6setNumEii
    // invoke: QString & setNum(int, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6setNumEii(this.qclsinst, arg0, arg1)
  case 6:
    // invoke: _ZN7QString6setNumEli
    // invoke: QString & setNum(long, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6setNumEli(this.qclsinst, arg0, arg1)
  case 7:
    // invoke: _ZN7QString6setNumEmi
    // invoke: QString & setNum(ulong, int)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6setNumEmi(this.qclsinst, arg0, arg1)
  case 8:
    // invoke: _ZN7QString6setNumEti
    // invoke: QString & setNum(ushort, int)
    var arg0 = C.int16_t(args[0].(int16))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C.demth_ZN7QString6setNumEti(this.qclsinst, arg0, arg1)
  case 9:
    // invoke: _ZN7QString6setNumEdci
    // invoke: QString & setNum(double, char, int)
    var arg0 = C.double(args[0].(float64))
    if false {fmt.Println(arg0)}
    var arg1 = C.uchar(args[1].(byte))
    if false {fmt.Println(arg1)}
    var arg2 = C.int32_t(args[2].(int32))
    if false {fmt.Println(arg2)}
    C._ZN7QString6setNumEdci(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "setNum", args)
  }

}

  // proto:  float QString::toFloat(bool * ok);
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

  // proto:  int QString::count(const QRegularExpression & re);
func (this *QString) count(args ...interface{}) () {
  // count(const class QStringRef &, Qt::CaseSensitivity)
  // count(const class QRegularExpression &)
  // count()
  // count(const class QRegExp &)
  // count(const class QString &, Qt::CaseSensitivity)
  // count(class QChar, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

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
    C.demth_ZNK7QString5countEv(this.qclsinst)
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

  // proto:  QStringRef QString::midRef(int position, int n);
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

  // proto:  void QString::detach();
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
    C.demth_ZN7QString6detachEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "detach", args)
  }

}

  // proto: static QString QString::fromStdWString(const std::wstring & s);
func (this *QString) fromStdWString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdWString", args)
  }

}

  // proto:  void QString::push_back(QChar c);
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
    C.demth_ZN7QString9push_backE5QChar(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString9push_backERKS_
    // invoke: void push_back(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString9push_backERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "push_back", args)
  }

}

  // proto:  int QString::size();
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
    C.demth_ZNK7QString4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "size", args)
  }

}

  // proto:  QString & QString::replace(int i, int len, const QChar * s, int slen);
func (this *QString) replace(args ...interface{}) () {
  // replace(const class QString &, const class QString &, Qt::CaseSensitivity)
  // replace(int, int, const class QChar *, int)
  // replace(const class QString &, class QLatin1String, Qt::CaseSensitivity)
  // replace(class QChar, class QLatin1String, Qt::CaseSensitivity)
  // replace(class QLatin1String, const class QString &, Qt::CaseSensitivity)
  // replace(const class QRegExp &, const class QString &)
  // replace(class QChar, const class QString &, Qt::CaseSensitivity)
  // replace(int, int, const class QString &)
  // replace(const class QChar *, int, const class QChar *, int, Qt::CaseSensitivity)
  // replace(class QLatin1String, class QLatin1String, Qt::CaseSensitivity)
  // replace(class QChar, class QChar, Qt::CaseSensitivity)
  // replace(const class QRegularExpression &, const class QString &)
  // replace(int, int, class QChar)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = qtrt.Int32Ty(false) // "int"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[1][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[1][3] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[2][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[2][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[3][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[4][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[5][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[6][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[6][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = qtrt.Int32Ty(false) // "int"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"
  vtys[7][2] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[8] = make(map[int32]reflect.Type)
  vtys[8][0] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[8][1] = qtrt.Int32Ty(false) // "int"
  vtys[8][2] = reflect.TypeOf(QChar{}) // "const QChar *"
  vtys[8][3] = qtrt.Int32Ty(false) // "int"
  vtys[8][4] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[9] = make(map[int32]reflect.Type)
  vtys[9][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[9][1] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[9][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[10] = make(map[int32]reflect.Type)
  vtys[10][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[10][1] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[10][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[11] = make(map[int32]reflect.Type)
  vtys[11][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[11][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[12] = make(map[int32]reflect.Type)
  vtys[12][0] = qtrt.Int32Ty(false) // "int"
  vtys[12][1] = qtrt.Int32Ty(false) // "int"
  vtys[12][2] = reflect.TypeOf(QChar{}) // "QChar"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
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
  case 1:
    // invoke: _ZN7QString7replaceERK7QRegExpRKS_
    // invoke: QString & replace(const class QRegExp &, const class QString &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString7replaceERK7QRegExpRKS_(this.qclsinst, arg0, arg1)
  case 2:
    // invoke: _ZN7QString7replaceEiiRKS_
    // invoke: QString & replace(int, int, const class QString &)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QString).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN7QString7replaceEiiRKS_(this.qclsinst, arg0, arg1, arg2)
  case 3:
    // invoke: _ZN7QString7replaceERK18QRegularExpressionRKS_
    // invoke: QString & replace(const class QRegularExpression &, const class QString &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = args[1].(QString).qclsinst
    if false {fmt.Println(arg1)}
    C._ZN7QString7replaceERK18QRegularExpressionRKS_(this.qclsinst, arg0, arg1)
  case 4:
    // invoke: _ZN7QString7replaceEii5QChar
    // invoke: QString & replace(int, int, class QChar)
    var arg0 = C.int32_t(args[0].(int32))
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QChar).qclsinst
    if false {fmt.Println(arg2)}
    C._ZN7QString7replaceEii5QChar(this.qclsinst, arg0, arg1, arg2)
  default:
    qtrt.ErrorResolve("QString", "replace", args)
  }

}

  // proto: static QString QString::fromRawData(const QChar * , int size);
func (this *QString) fromRawData_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromRawData", args)
  }

}

  // proto:  QString & QString::setRawData(const QChar * unicode, int size);
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

  // proto:  ulong QString::toULong(bool * ok, int base);
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

  // proto:  void QString::chop(int n);
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

  // proto: static QString QString::fromUtf16(const ushort * , int size);
func (this *QString) fromUtf16_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromUtf16", args)
  }

}

  // proto:  bool QString::isDetached();
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
    C.demth_ZNK7QString10isDetachedEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isDetached", args)
  }

}

  // proto:  QString QString::mid(int position, int n);
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

  // proto: static QString QString::fromLocal8Bit(const char * str, int size);
func (this *QString) fromLocal8Bit_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromLocal8Bit", args)
  }

}

  // proto:  void QString::swap(QString & other);
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
    C.demth_ZN7QString4swapERS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "swap", args)
  }

}

  // proto: static QString QString::fromUtf8(const QByteArray & str);
func (this *QString) fromUtf8_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromUtf8", args)
  }

}

  // proto: static QString QString::fromUcs4(const char32_t * str, int size);
func (this *QString) fromUcs4_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromUcs4", args)
  }

}

  // proto:  QString QString::leftJustified(int width, QChar fill, bool trunc);
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

  // proto:  int QString::indexOf(const QRegExp & , int from);
func (this *QString) indexOf(args ...interface{}) () {
  // indexOf(const class QString &, int, Qt::CaseSensitivity)
  // indexOf(const class QRegExp &, int)
  // indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
  // indexOf(const class QStringRef &, int, Qt::CaseSensitivity)
  // indexOf(const class QRegularExpression &, int)
  // indexOf(class QChar, int, Qt::CaseSensitivity)
  // indexOf(class QLatin1String, int, Qt::CaseSensitivity)
  // indexOf(class QRegExp &, int)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = qtrt.Int32Ty(false) // "int"
  vtys[0][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[1][1] = qtrt.Int32Ty(false) // "int"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2][1] = qtrt.Int32Ty(false) // "int"
  vtys[2][2] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[3][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[4][1] = qtrt.Int32Ty(false) // "int"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5][1] = qtrt.Int32Ty(false) // "int"
  vtys[5][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[6][1] = qtrt.Int32Ty(false) // "int"
  vtys[6][2] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[7][1] = qtrt.Int32Ty(false) // "int"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString7indexOfERK7QRegExpi
    // invoke: int indexOf(const class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7indexOfERK7QRegExpi(this.qclsinst, arg0, arg1)
  case 1:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch
    // invoke: int indexOf(const class QRegularExpression &, int, class QRegularExpressionMatch *)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    var arg2 = args[2].(QRegularExpressionMatch).qclsinst
    if false {fmt.Println(arg2)}
    C._ZNK7QString7indexOfERK18QRegularExpressioniP23QRegularExpressionMatch(this.qclsinst, arg0, arg1, arg2)
  case 2:
    // invoke: _ZNK7QString7indexOfERK18QRegularExpressioni
    // invoke: int indexOf(const class QRegularExpression &, int)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7indexOfERK18QRegularExpressioni(this.qclsinst, arg0, arg1)
  case 3:
    // invoke: _ZNK7QString7indexOfER7QRegExpi
    // invoke: int indexOf(class QRegExp &, int)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    var arg1 = C.int32_t(args[1].(int32))
    if false {fmt.Println(arg1)}
    C._ZNK7QString7indexOfER7QRegExpi(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "indexOf", args)
  }

}

  // proto:  const ushort * QString::utf16();
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

  // proto:  int QString::toInt(bool * ok, int base);
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

  // proto:  QChar * QString::data();
func (this *QString) data(args ...interface{}) () {
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
    // invoke: _ZN7QString4dataEv
    // invoke: QChar * data()
    C.demth_ZN7QString4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "data", args)
  }

}

  // proto: static int QString::localeAwareCompare(const QString & s1, const QString & s2);
func (this *QString) localeAwareCompare_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "localeAwareCompare", args)
  }

}

  // proto:  QString QString::repeated(int times);
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

  // proto:  QString & QString::setUtf16(const ushort * utf16, int size);
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
    C.demth_ZN7QString8setUtf16EPKti(this.qclsinst, arg0, arg1)
  default:
    qtrt.ErrorResolve("QString", "setUtf16", args)
  }

}

  // proto: static QString QString::fromStdU32String(const std::u32string & s);
func (this *QString) fromStdU32String_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdU32String", args)
  }

}

  // proto:  void QString::clear();
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
    C.demth_ZN7QString5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "clear", args)
  }

}

  // proto:  bool QString::contains(const QRegExp & rx);
func (this *QString) contains(args ...interface{}) () {
  // contains(class QLatin1String, Qt::CaseSensitivity)
  // contains(const class QRegExp &)
  // contains(const class QStringRef &, Qt::CaseSensitivity)
  // contains(const class QRegularExpression &, class QRegularExpressionMatch *)
  // contains(class QRegExp &)
  // contains(class QChar, Qt::CaseSensitivity)
  // contains(const class QRegularExpression &)
  // contains(const class QString &, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QLatin1String{}) // "QLatin1String"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[3][1] = reflect.TypeOf(QRegularExpressionMatch{}) // "QRegularExpressionMatch *"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QRegExp{}) // "QRegExp &"
  vtys[5] = make(map[int32]reflect.Type)
  vtys[5][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[5][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[6] = make(map[int32]reflect.Type)
  vtys[6][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[7] = make(map[int32]reflect.Type)
  vtys[7][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[7][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString8containsERK7QRegExp
    // invoke: bool contains(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QString8containsERK7QRegExp(this.qclsinst, arg0)
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
    C.demth_ZNK7QString8containsER7QRegExp(this.qclsinst, arg0)
  case 3:
    // invoke: _ZNK7QString8containsERK18QRegularExpression
    // invoke: bool contains(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C._ZNK7QString8containsERK18QRegularExpression(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "contains", args)
  }

}

  // proto:  bool QString::isSharedWith(const QString & other);
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
    C.demth_ZNK7QString12isSharedWithERKS_(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "isSharedWith", args)
  }

}

  // proto: static QString QString::fromLatin1(const QByteArray & str);
func (this *QString) fromLatin1_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromLatin1", args)
  }

}

  // proto:  void QString::~QString();
func (this *QString) FreeQString(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "~QString", args)
  }

}

  // proto:  QString & QString::remove(const QRegularExpression & re);
func (this *QString) remove(args ...interface{}) () {
  // remove(class QChar, Qt::CaseSensitivity)
  // remove(const class QRegularExpression &)
  // remove(const class QRegExp &)
  // remove(int, int)
  // remove(const class QString &, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QRegularExpression{}) // "const QRegularExpression &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QRegExp{}) // "const QRegExp &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = qtrt.Int32Ty(false) // "int"
  vtys[3][1] = qtrt.Int32Ty(false) // "int"
  vtys[4] = make(map[int32]reflect.Type)
  vtys[4][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[4][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN7QString6removeERK18QRegularExpression
    // invoke: QString & remove(const class QRegularExpression &)
    var arg0 = args[0].(QRegularExpression).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString6removeERK18QRegularExpression(this.qclsinst, arg0)
  case 1:
    // invoke: _ZN7QString6removeERK7QRegExp
    // invoke: QString & remove(const class QRegExp &)
    var arg0 = args[0].(QRegExp).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZN7QString6removeERK7QRegExp(this.qclsinst, arg0)
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

  // proto:  QString QString::toHtmlEscaped();
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

  // proto:  int QString::toWCharArray(wchar_t * array);
func (this *QString) toWCharArray(args ...interface{}) () {
  // toWCharArray(wchar_t *)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = qtrt.StringTy(false) // "wchar_t *"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK7QString12toWCharArrayEPw
    // invoke: int toWCharArray(wchar_t *)
    var arg0 = C.CString(args[0].(string))
    if false {fmt.Println(arg0)}
    C.demth_ZNK7QString12toWCharArrayEPw(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "toWCharArray", args)
  }

}

  // proto: static QString QString::fromStdString(const std::string & s);
func (this *QString) fromStdString_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdString", args)
  }

}

  // proto: static QString QString::fromWCharArray(const wchar_t * string, int size);
func (this *QString) fromWCharArray_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromWCharArray", args)
  }

}

  // proto:  QString & QString::fill(QChar c, int size);
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

  // proto:  const QChar * QString::constData();
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
    C.demth_ZNK7QString9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "constData", args)
  }

}

  // proto:  long QString::toLong(bool * ok, int base);
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

  // proto:  int QString::length();
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
    C.demth_ZNK7QString6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "length", args)
  }

}

  // proto:  QStringRef QString::leftRef(int n);
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

  // proto:  bool QString::isSimpleText();
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

  // proto:  QString & QString::setUnicode(const QChar * unicode, int size);
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

  // proto:  const QChar * QString::unicode();
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
    C.demth_ZNK7QString7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "unicode", args)
  }

}

  // proto:  const QChar QString::at(int i);
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
    C.demth_ZNK7QString2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "at", args)
  }

}

  // proto:  uint QString::toUInt(bool * ok, int base);
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

  // proto: static QString QString::fromStdU16String(const std::u16string & s);
func (this *QString) fromStdU16String_s(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QString", "fromStdU16String", args)
  }

}

  // proto:  ushort QString::toUShort(bool * ok, int base);
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

  // proto:  qulonglong QString::toULongLong(bool * ok, int base);
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

  // proto:  int QString::capacity();
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
    C.demth_ZNK7QString8capacityEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "capacity", args)
  }

}

  // proto:  void QString::squeeze();
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
    C.demth_ZN7QString7squeezeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "squeeze", args)
  }

}

  // proto:  void QString::truncate(int pos);
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

  // proto:  int QString::localeAwareCompare(const QString & s);
func (this *QString) localeAwareCompare(args ...interface{}) () {
  // localeAwareCompare(const class QString &, const class QString &)
  // localeAwareCompare(const class QString &, const class QStringRef &)
  // localeAwareCompare(const class QStringRef &)
  // localeAwareCompare(const class QString &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[0][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QString{}) // "const QString &"

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
    C.demth_ZN7QString18localeAwareCompareERKS_S1_(arg0, arg1)
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

  // proto:  bool QString::isRightToLeft();
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

  // proto:  QVector<uint> QString::toUcs4();
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

  // proto:  bool QString::isEmpty();
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
    C.demth_ZNK7QString7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QString", "isEmpty", args)
  }

}

  // proto:  QString QString::right(int n);
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

  // proto:  QString QString::rightJustified(int width, QChar fill, bool trunc);
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

  // proto:  void QString::reserve(int size);
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
    C.demth_ZN7QString7reserveEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QString", "reserve", args)
  }

}

  // proto:  short QString::toShort(bool * ok, int base);
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

  // proto:  const char * QLatin1String::data();
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
    C.demth_ZNK13QLatin1String4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1String", "data", args)
  }

}

  // proto:  void QLatin1String::QLatin1String(const char * s);
func NewQLatin1String(args ...interface{}) QLatin1String {
  return QLatin1String{}
}

  // proto:  int QLatin1String::size();
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
    C.demth_ZNK13QLatin1String4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1String", "size", args)
  }

}

  // proto:  const char * QLatin1String::latin1();
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
    C.demth_ZNK13QLatin1String6latin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QLatin1String", "latin1", args)
  }

}

  // proto:  bool QCharRef::isLetterOrNumber();
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
    C.demth_ZN8QCharRef16isLetterOrNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isLetterOrNumber", args)
  }

}

  // proto:  bool QCharRef::isDigit();
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
    C.demth_ZNK8QCharRef7isDigitEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isDigit", args)
  }

}

  // proto:  char QCharRef::toLatin1();
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
    C.demth_ZNK8QCharRef8toLatin1Ev(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toLatin1", args)
  }

}

  // proto:  void QCharRef::setCell(uchar cell);
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
    C.demth_ZN8QCharRef7setCellEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCharRef", "setCell", args)
  }

}

  // proto:  bool QCharRef::isMark();
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
    C.demth_ZNK8QCharRef6isMarkEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isMark", args)
  }

}

  // proto:  void QCharRef::QCharRef(QString & str, int idx);
func NewQCharRef(args ...interface{}) QCharRef {
  return QCharRef{}
}

  // proto:  int QCharRef::digitValue();
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
    C.demth_ZNK8QCharRef10digitValueEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "digitValue", args)
  }

}

  // proto:  bool QCharRef::isLetter();
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
    C.demth_ZNK8QCharRef8isLetterEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isLetter", args)
  }

}

  // proto:  bool QCharRef::isNumber();
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
    C.demth_ZNK8QCharRef8isNumberEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isNumber", args)
  }

}

  // proto:  bool QCharRef::isPrint();
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
    C.demth_ZNK8QCharRef7isPrintEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isPrint", args)
  }

}

  // proto:  QChar QCharRef::toLower();
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
    C.demth_ZNK8QCharRef7toLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toLower", args)
  }

}

  // proto:  void QCharRef::setRow(uchar row);
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
    C.demth_ZN8QCharRef6setRowEh(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QCharRef", "setRow", args)
  }

}

  // proto:  bool QCharRef::isNull();
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
    C.demth_ZNK8QCharRef6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isNull", args)
  }

}

  // proto:  QChar QCharRef::toTitleCase();
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
    C.demth_ZNK8QCharRef11toTitleCaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toTitleCase", args)
  }

}

  // proto:  bool QCharRef::hasMirrored();
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
    C.demth_ZNK8QCharRef11hasMirroredEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "hasMirrored", args)
  }

}

  // proto:  uchar QCharRef::row();
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
    C.demth_ZNK8QCharRef3rowEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "row", args)
  }

}

  // proto:  ushort & QCharRef::unicode();
func (this *QCharRef) unicode(args ...interface{}) () {
  // unicode()
  // unicode()
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[1] = make(map[int32]reflect.Type)

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZN8QCharRef7unicodeEv
    // invoke: ushort & unicode()
    C.demth_ZN8QCharRef7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "unicode", args)
  }

}

  // proto:  bool QCharRef::isTitleCase();
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
    C.demth_ZNK8QCharRef11isTitleCaseEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isTitleCase", args)
  }

}

  // proto:  bool QCharRef::isUpper();
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
    C.demth_ZNK8QCharRef7isUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isUpper", args)
  }

}

  // proto:  uchar QCharRef::cell();
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
    C.demth_ZNK8QCharRef4cellEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "cell", args)
  }

}

  // proto:  QString QCharRef::decomposition();
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
    C.demth_ZNK8QCharRef13decompositionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "decomposition", args)
  }

}

  // proto:  uchar QCharRef::combiningClass();
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
    C.demth_ZNK8QCharRef14combiningClassEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "combiningClass", args)
  }

}

  // proto:  QChar QCharRef::mirroredChar();
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
    C.demth_ZNK8QCharRef12mirroredCharEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "mirroredChar", args)
  }

}

  // proto:  bool QCharRef::isSpace();
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
    C.demth_ZNK8QCharRef7isSpaceEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isSpace", args)
  }

}

  // proto:  bool QCharRef::isPunct();
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
    C.demth_ZNK8QCharRef7isPunctEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isPunct", args)
  }

}

  // proto:  QChar QCharRef::toUpper();
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
    C.demth_ZNK8QCharRef7toUpperEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "toUpper", args)
  }

}

  // proto:  bool QCharRef::isLower();
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
    C.demth_ZNK8QCharRef7isLowerEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QCharRef", "isLower", args)
  }

}

  // proto:  short QStringRef::toShort(bool * ok, int base);
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

  // proto:  void QStringRef::QStringRef(const QString * string);
func NewQStringRef(args ...interface{}) QStringRef {
  return QStringRef{}
}

  // proto:  qulonglong QStringRef::toULongLong(bool * ok, int base);
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

  // proto:  void QStringRef::clear();
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
    C.demth_ZN10QStringRef5clearEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "clear", args)
  }

}

  // proto:  int QStringRef::position();
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
    C.demth_ZNK10QStringRef8positionEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "position", args)
  }

}

  // proto:  long QStringRef::toLong(bool * ok, int base);
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

  // proto:  const QChar * QStringRef::cbegin();
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
    C.demth_ZNK10QStringRef6cbeginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "cbegin", args)
  }

}

  // proto:  ushort QStringRef::toUShort(bool * ok, int base);
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

  // proto:  uint QStringRef::toUInt(bool * ok, int base);
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

  // proto:  bool QStringRef::isEmpty();
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
    C.demth_ZNK10QStringRef7isEmptyEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "isEmpty", args)
  }

}

  // proto:  int QStringRef::localeAwareCompare(const QString & s);
func (this *QStringRef) localeAwareCompare(args ...interface{}) () {
  // localeAwareCompare(const class QString &)
  // localeAwareCompare(const class QStringRef &)
  // localeAwareCompare(const class QStringRef &, const class QString &)
  // localeAwareCompare(const class QStringRef &, const class QStringRef &)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[2][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[2][1] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[3][1] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef18localeAwareCompareERK7QString
    // invoke: int localeAwareCompare(const class QString &)
    var arg0 = args[0].(QString).qclsinst
    if false {fmt.Println(arg0)}
    C.demth_ZNK10QStringRef18localeAwareCompareERK7QString(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "localeAwareCompare", args)
  }

}

  // proto:  QByteArray QStringRef::toUtf8();
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

  // proto:  int QStringRef::size();
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
    C.demth_ZNK10QStringRef4sizeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "size", args)
  }

}

  // proto:  const QChar * QStringRef::constData();
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
    C.demth_ZNK10QStringRef9constDataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "constData", args)
  }

}

  // proto:  QStringRef QStringRef::left(int n);
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

  // proto:  QVector<uint> QStringRef::toUcs4();
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

  // proto:  int QStringRef::count();
func (this *QStringRef) count(args ...interface{}) () {
  // count(const class QStringRef &, Qt::CaseSensitivity)
  // count(const class QString &, Qt::CaseSensitivity)
  // count()
  // count(class QChar, Qt::CaseSensitivity)
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}
  vtys[0] = make(map[int32]reflect.Type)
  vtys[0][0] = reflect.TypeOf(QStringRef{}) // "const QStringRef &"
  vtys[0][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[1] = make(map[int32]reflect.Type)
  vtys[1][0] = reflect.TypeOf(QString{}) // "const QString &"
  vtys[1][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"
  vtys[2] = make(map[int32]reflect.Type)
  vtys[3] = make(map[int32]reflect.Type)
  vtys[3][0] = reflect.TypeOf(QChar{}) // "QChar"
  vtys[3][1] = qtrt.Int32Ty(false) // "Qt::CaseSensitivity"

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  case 0:
    // invoke: _ZNK10QStringRef5countEv
    // invoke: int count()
    C.demth_ZNK10QStringRef5countEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "count", args)
  }

}

  // proto:  QStringRef QStringRef::right(int n);
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

  // proto:  const QChar QStringRef::at(int i);
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
    C.demth_ZNK10QStringRef2atEi(this.qclsinst, arg0)
  default:
    qtrt.ErrorResolve("QStringRef", "at", args)
  }

}

  // proto:  double QStringRef::toDouble(bool * ok);
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

  // proto:  bool QStringRef::isNull();
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
    C.demth_ZNK10QStringRef6isNullEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "isNull", args)
  }

}

  // proto:  const QChar * QStringRef::data();
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
    C.demth_ZNK10QStringRef4dataEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "data", args)
  }

}

  // proto:  qlonglong QStringRef::toLongLong(bool * ok, int base);
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

  // proto:  QByteArray QStringRef::toLatin1();
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

  // proto:  const QChar * QStringRef::begin();
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
    C.demth_ZNK10QStringRef5beginEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "begin", args)
  }

}

  // proto:  const QChar * QStringRef::unicode();
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
    C.demth_ZNK10QStringRef7unicodeEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "unicode", args)
  }

}

  // proto:  QStringRef QStringRef::mid(int pos, int n);
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

  // proto:  float QStringRef::toFloat(bool * ok);
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

  // proto:  const QString * QStringRef::string();
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
    C.demth_ZNK10QStringRef6stringEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "string", args)
  }

}

  // proto:  QString QStringRef::toString();
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

  // proto:  QStringRef QStringRef::trimmed();
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

  // proto:  int QStringRef::toInt(bool * ok, int base);
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

  // proto:  const QChar * QStringRef::cend();
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
    C.demth_ZNK10QStringRef4cendEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "cend", args)
  }

}

  // proto:  QStringRef QStringRef::appendTo(QString * string);
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

  // proto:  int QStringRef::length();
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
    C.demth_ZNK10QStringRef6lengthEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "length", args)
  }

}

  // proto:  void QStringRef::~QStringRef();
func (this *QStringRef) FreeQStringRef(args ...interface{}) () {
  var vtys = make(map[int32]map[int32]reflect.Type)
  if false {fmt.Println(vtys)}

  var matched_index = qtrt.SymbolResolve(args, vtys)
  if false {fmt.Println(matched_index)}
  switch matched_index {
  default:
    qtrt.ErrorResolve("QStringRef", "~QStringRef", args)
  }

}

  // proto:  QByteArray QStringRef::toLocal8Bit();
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

  // proto:  ulong QStringRef::toULong(bool * ok, int base);
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

  // proto:  const QChar * QStringRef::end();
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
    C.demth_ZNK10QStringRef3endEv(this.qclsinst)
  default:
    qtrt.ErrorResolve("QStringRef", "end", args)
  }

}

// <= body block end

