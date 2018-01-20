//  header block begin
// /usr/include/qt/QtCore/qurl.h
// #include <qurl.h>
// #include <QtCore>
package qtcore

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
}

//  ext block end

//  body block begin
type QUrl struct {
	*qtrt.CObject
}

func (this *QUrl) GetCthis() unsafe.Pointer {
	return this.Cthis
}

// /usr/include/qt/QtCore/qurl.h:176
// index:0
// void QUrl()
func NewQUrl() *QUrl {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrlC2Ev", ffiqt.FFI_TYPE_VOID, cthis)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(cthis)
	return gothis
}
func NewQUrlFromPointer(cthis unsafe.Pointer) *QUrl {
	return &QUrl{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtCore/qurl.h:182
// index:1
// void QUrl(const class QString &, enum QUrl::ParsingMode)
func NewQUrl_1(url unsafe.Pointer, mode int) *QUrl {
	cthis := qtrt.Calloc(1, 256)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrlC2ERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, cthis, url, &mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(cthis)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:191
// index:0
// void ~QUrl()
func DeleteQUrl(*QUrl) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrlD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:193
// index:0
// inline
// void swap(class QUrl &)
func (this *QUrl) Swap(other unsafe.Pointer) {
	// 0: (, other QUrl &), (other)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl4swapERS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), other)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:195
// index:0
// void setUrl(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUrl(url unsafe.Pointer, mode int) {
	// 0: (, url const QString &, mode QUrl::ParsingMode), (url, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), url, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:202
// index:0
// static
// QUrl fromEncoded(const class QByteArray &, enum QUrl::ParsingMode)
func (this *QUrl) FromEncoded(url unsafe.Pointer, mode int) {
	// 0: (url const QByteArray &, mode QUrl::ParsingMode), (url, mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11fromEncodedERK10QByteArrayNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromEncoded(url unsafe.Pointer, mode int) {
	// 0: (url const QByteArray &, mode QUrl::ParsingMode), (url, mode)
	var nilthis *QUrl
	nilthis.FromEncoded(url, mode)
}

// /usr/include/qt/QtCore/qurl.h:210
// index:0
// static
// QUrl fromUserInput(const class QString &)
func (this *QUrl) FromUserInput(userInput unsafe.Pointer) {
	// 0: (userInput const QString &), (userInput)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromUserInput(userInput unsafe.Pointer) {
	// 0: (userInput const QString &), (userInput)
	var nilthis *QUrl
	nilthis.FromUserInput(userInput)
}

// /usr/include/qt/QtCore/qurl.h:212
// index:1
// static
// QUrl fromUserInput(const class QString &, const class QString &, QUrl::UserInputResolutionOptions)
func (this *QUrl) FromUserInput_1(userInput unsafe.Pointer, workingDirectory unsafe.Pointer, options int) {
	// 1: (userInput const QString &, workingDirectory const QString &, QFlags<QUrl::UserInputResolutionOption> options), (userInput, workingDirectory, options)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QStringS2_6QFlagsINS_25UserInputResolutionOptionEE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromUserInput_1(userInput unsafe.Pointer, workingDirectory unsafe.Pointer, options int) {
	// 1: (userInput const QString &, workingDirectory const QString &, QFlags<QUrl::UserInputResolutionOption> options), (userInput, workingDirectory, options)
	var nilthis *QUrl
	nilthis.FromUserInput_1(userInput, workingDirectory, options)
}

// /usr/include/qt/QtCore/qurl.h:215
// index:0
// bool isValid()
func (this *QUrl) IsValid() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl7isValidEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:216
// index:0
// QString errorString()
func (this *QUrl) ErrorString() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11errorStringEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:218
// index:0
// bool isEmpty()
func (this *QUrl) IsEmpty() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl7isEmptyEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:219
// index:0
// void clear()
func (this *QUrl) Clear() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl5clearEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:221
// index:0
// void setScheme(const class QString &)
func (this *QUrl) SetScheme(scheme unsafe.Pointer) {
	// 0: (, scheme const QString &), (scheme)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl9setSchemeERK7QString", ffiqt.FFI_TYPE_VOID, this.GetCthis(), scheme)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:222
// index:0
// QString scheme()
func (this *QUrl) Scheme() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl6schemeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:224
// index:0
// void setAuthority(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetAuthority(authority unsafe.Pointer, mode int) {
	// 0: (, authority const QString &, mode QUrl::ParsingMode), (authority, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl12setAuthorityERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), authority, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:227
// index:0
// void setUserInfo(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserInfo(userInfo unsafe.Pointer, mode int) {
	// 0: (, userInfo const QString &, mode QUrl::ParsingMode), (userInfo, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setUserInfoERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), userInfo, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:230
// index:0
// void setUserName(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserName(userName unsafe.Pointer, mode int) {
	// 0: (, userName const QString &, mode QUrl::ParsingMode), (userName, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setUserNameERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), userName, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:233
// index:0
// void setPassword(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPassword(password unsafe.Pointer, mode int) {
	// 0: (, password const QString &, mode QUrl::ParsingMode), (password, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setPasswordERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), password, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:236
// index:0
// void setHost(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetHost(host unsafe.Pointer, mode int) {
	// 0: (, host const QString &, mode QUrl::ParsingMode), (host, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7setHostERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), host, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:242
// index:0
// void setPort(int)
func (this *QUrl) SetPort(port int) {
	// 0: (, port int), (&port)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7setPortEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:243
// index:0
// int port(int)
func (this *QUrl) Port(defaultPort int) {
	// 0: (, defaultPort int), (&defaultPort)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl4portEi", ffiqt.FFI_TYPE_VOID, this.GetCthis(), &defaultPort)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:245
// index:0
// void setPath(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPath(path unsafe.Pointer, mode int) {
	// 0: (, path const QString &, mode QUrl::ParsingMode), (path, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7setPathERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), path, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:249
// index:0
// bool hasQuery()
func (this *QUrl) HasQuery() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl8hasQueryEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:250
// index:0
// void setQuery(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetQuery(query unsafe.Pointer, mode int) {
	// 0: (, query const QString &, mode QUrl::ParsingMode), (query, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl8setQueryERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), query, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:251
// index:1
// void setQuery(const class QUrlQuery &)
func (this *QUrl) SetQuery_1(query unsafe.Pointer) {
	// 1: (, query const QUrlQuery &), (query)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl8setQueryERK9QUrlQuery", ffiqt.FFI_TYPE_VOID, this.GetCthis(), query)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:254
// index:0
// bool hasFragment()
func (this *QUrl) HasFragment() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11hasFragmentEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:256
// index:0
// void setFragment(const class QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetFragment(fragment unsafe.Pointer, mode int) {
	// 0: (, fragment const QString &, mode QUrl::ParsingMode), (fragment, &mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setFragmentERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID, this.GetCthis(), fragment, &mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:258
// index:0
// QUrl resolved(const class QUrl &)
func (this *QUrl) Resolved(relative unsafe.Pointer) {
	// 0: (, relative const QUrl &), (relative)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl8resolvedERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), relative)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:260
// index:0
// bool isRelative()
func (this *QUrl) IsRelative() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl10isRelativeEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:261
// index:0
// bool isParentOf(const class QUrl &)
func (this *QUrl) IsParentOf(url unsafe.Pointer) {
	// 0: (, url const QUrl &), (url)
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl10isParentOfERKS_", ffiqt.FFI_TYPE_VOID, this.GetCthis(), url)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:263
// index:0
// bool isLocalFile()
func (this *QUrl) IsLocalFile() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11isLocalFileEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:264
// index:0
// static
// QUrl fromLocalFile(const class QString &)
func (this *QUrl) FromLocalFile(localfile unsafe.Pointer) {
	// 0: (localfile const QString &), (localfile)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl13fromLocalFileERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromLocalFile(localfile unsafe.Pointer) {
	// 0: (localfile const QString &), (localfile)
	var nilthis *QUrl
	nilthis.FromLocalFile(localfile)
}

// /usr/include/qt/QtCore/qurl.h:265
// index:0
// QString toLocalFile()
func (this *QUrl) ToLocalFile() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11toLocalFileEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:267
// index:0
// void detach()
func (this *QUrl) Detach() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl6detachEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:268
// index:0
// bool isDetached()
func (this *QUrl) IsDetached() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl10isDetachedEv", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:276
// index:0
// static
// QString fromPercentEncoding(const class QByteArray &)
func (this *QUrl) FromPercentEncoding(arg0 unsafe.Pointer) {
	// 0: (const QByteArray & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl19fromPercentEncodingERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromPercentEncoding(arg0 unsafe.Pointer) {
	// 0: (const QByteArray & arg0), (arg0)
	var nilthis *QUrl
	nilthis.FromPercentEncoding(arg0)
}

// /usr/include/qt/QtCore/qurl.h:277
// index:0
// static
// QByteArray toPercentEncoding(const class QString &, const class QByteArray &, const class QByteArray &)
func (this *QUrl) ToPercentEncoding(arg0 unsafe.Pointer, exclude unsafe.Pointer, include unsafe.Pointer) {
	// 0: (const QString & arg0, exclude const QByteArray &, include const QByteArray &), (arg0, exclude, include)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_ToPercentEncoding(arg0 unsafe.Pointer, exclude unsafe.Pointer, include unsafe.Pointer) {
	// 0: (const QString & arg0, exclude const QByteArray &, include const QByteArray &), (arg0, exclude, include)
	var nilthis *QUrl
	nilthis.ToPercentEncoding(arg0, exclude, include)
}

// /usr/include/qt/QtCore/qurl.h:357
// index:0
// static
// QString fromAce(const class QByteArray &)
func (this *QUrl) FromAce(arg0 unsafe.Pointer) {
	// 0: (const QByteArray & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7fromAceERK10QByteArray", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromAce(arg0 unsafe.Pointer) {
	// 0: (const QByteArray & arg0), (arg0)
	var nilthis *QUrl
	nilthis.FromAce(arg0)
}

// /usr/include/qt/QtCore/qurl.h:358
// index:0
// static
// QByteArray toAce(const class QString &)
func (this *QUrl) ToAce(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl5toAceERK7QString", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_ToAce(arg0 unsafe.Pointer) {
	// 0: (const QString & arg0), (arg0)
	var nilthis *QUrl
	nilthis.ToAce(arg0)
}

// /usr/include/qt/QtCore/qurl.h:359
// index:0
// static
// QStringList idnWhitelist()
func (this *QUrl) IdnWhitelist() {
	// 0: (), ()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl12idnWhitelistEv", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_IdnWhitelist() {
	// 0: (), ()
	var nilthis *QUrl
	nilthis.IdnWhitelist()
}

// /usr/include/qt/QtCore/qurl.h:361
// index:0
// static
// QList<QUrl> fromStringList(const class QStringList &, enum QUrl::ParsingMode)
func (this *QUrl) FromStringList(uris unsafe.Pointer, mode int) {
	// 0: (uris const QStringList &, mode QUrl::ParsingMode), (uris, mode)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl14fromStringListERK11QStringListNS_11ParsingModeE", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_FromStringList(uris unsafe.Pointer, mode int) {
	// 0: (uris const QStringList &, mode QUrl::ParsingMode), (uris, mode)
	var nilthis *QUrl
	nilthis.FromStringList(uris, mode)
}

// /usr/include/qt/QtCore/qurl.h:363
// index:0
// static
// void setIdnWhitelist(const class QStringList &)
func (this *QUrl) SetIdnWhitelist(arg0 unsafe.Pointer) {
	// 0: (const QStringList & arg0), (arg0)
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl15setIdnWhitelistERK11QStringList", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}
func QUrl_SetIdnWhitelist(arg0 unsafe.Pointer) {
	// 0: (const QStringList & arg0), (arg0)
	var nilthis *QUrl
	nilthis.SetIdnWhitelist(arg0)
}

//  body block end
