package qtcore

// /usr/include/qt/QtCore/qurl.h
// #include <qurl.h>
// #include <QtCore>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QUrl) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQUrlFromPointer(cthis unsafe.Pointer) *QUrl {
	return &QUrl{&qtrt.CObject{cthis}}
}
func (*QUrl) NewFromPointer(cthis unsafe.Pointer) *QUrl {
	return NewQUrlFromPointer(cthis)
}

// /usr/include/qt/QtCore/qurl.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUrl()
func NewQUrl() *QUrl {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrlC2Ev", ffiqt.FFI_TYPE_POINTER)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:182
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrl(const QString &, enum QUrl::ParsingMode)
func NewQUrl_1(url *QString, mode int) *QUrl {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrlC2ERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, convArg0, mode)
	gopp.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QUrl()
func DeleteQUrl(this *QUrl) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrlD2Ev", ffiqt.FFI_TYPE_VOID, this.GetCthis())
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qurl.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QUrl &)
func (this *QUrl) Swap(other *QUrl) {
	var convArg0 = other.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl4swapERS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUrl(url *QString, mode int) {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:202
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromEncoded(const QByteArray &, enum QUrl::ParsingMode)
func (this *QUrl) FromEncoded(url *QByteArray, mode int) *QUrl /*123*/ {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11fromEncodedERK10QByteArrayNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, convArg0, mode)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromEncoded(url *QByteArray, mode int) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromEncoded(url, mode)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromUserInput(const QString &)
func (this *QUrl) FromUserInput(userInput *QString) *QUrl /*123*/ {
	var convArg0 = userInput.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromUserInput(userInput *QString) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromUserInput(userInput)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:215
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid()
func (this *QUrl) IsValid() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl7isValidEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:216
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString()
func (this *QUrl) ErrorString() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11errorStringEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:218
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty()
func (this *QUrl) IsEmpty() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl7isEmptyEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QUrl) Clear() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl5clearEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScheme(const QString &)
func (this *QUrl) SetScheme(scheme *QString) {
	var convArg0 = scheme.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl9setSchemeERK7QString", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] QString scheme()
func (this *QUrl) Scheme() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl6schemeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAuthority(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetAuthority(authority *QString, mode int) {
	var convArg0 = authority.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl12setAuthorityERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserInfo(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserInfo(userInfo *QString, mode int) {
	var convArg0 = userInfo.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setUserInfoERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserName(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserName(userName *QString, mode int) {
	var convArg0 = userName.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setUserNameERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPassword(password *QString, mode int) {
	var convArg0 = password.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setPasswordERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetHost(host *QString, mode int) {
	var convArg0 = host.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7setHostERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPort(int)
func (this *QUrl) SetPort(port int) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7setPortEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), port)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int port(int)
func (this *QUrl) Port(defaultPort int) int {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl4portEi", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), defaultPort)
	gopp.ErrPrint(err, rv)
	//  return rv
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qurl.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPath(path *QString, mode int) {
	var convArg0 = path.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7setPathERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:249
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasQuery()
func (this *QUrl) HasQuery() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl8hasQueryEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetQuery(query *QString, mode int) {
	var convArg0 = query.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl8setQueryERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:251
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QUrlQuery &)
func (this *QUrl) SetQuery_1(query *QUrlQuery) {
	var convArg0 = query.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl8setQueryERK9QUrlQuery", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFragment()
func (this *QUrl) HasFragment() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11hasFragmentEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFragment(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetFragment(fragment *QString, mode int) {
	var convArg0 = fragment.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl11setFragmentERK7QStringNS_11ParsingModeE", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:258
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl resolved(const QUrl &)
func (this *QUrl) Resolved(relative *QUrl) *QUrl /*123*/ {
	var convArg0 = relative.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl8resolvedERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRelative()
func (this *QUrl) IsRelative() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl10isRelativeEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isParentOf(const QUrl &)
func (this *QUrl) IsParentOf(url *QUrl) bool {
	var convArg0 = url.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl10isParentOfERKS_", ffiqt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:263
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLocalFile()
func (this *QUrl) IsLocalFile() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11isLocalFileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:264
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromLocalFile(const QString &)
func (this *QUrl) FromLocalFile(localfile *QString) *QUrl /*123*/ {
	var convArg0 = localfile.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl13fromLocalFileERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromLocalFile(localfile *QString) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromLocalFile(localfile)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:265
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toLocalFile()
func (this *QUrl) ToLocalFile() *QString /*123*/ {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl11toLocalFileEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QUrl) Detach() {
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl6detachEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:268
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached()
func (this *QUrl) IsDetached() bool {
	rv, err := ffiqt.InvokeQtFunc6("_ZNK4QUrl10isDetachedEv", ffiqt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	//  return rv
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:276
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromPercentEncoding(const QByteArray &)
func (this *QUrl) FromPercentEncoding(arg0 *QByteArray) *QString /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl19fromPercentEncodingERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QUrl_FromPercentEncoding(arg0 *QByteArray) *QString /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromPercentEncoding(arg0)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QString &, const QByteArray &, const QByteArray &)
func (this *QUrl) ToPercentEncoding(arg0 *QString, exclude *QByteArray, include *QByteArray) *QByteArray /*123*/ {
	var convArg0 = arg0.GetCthis()
	var convArg1 = exclude.GetCthis()
	var convArg2 = include.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_", ffiqt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QUrl_ToPercentEncoding(arg0 *QString, exclude *QByteArray, include *QByteArray) *QByteArray /*123*/ {
	var nilthis *QUrl
	rv := nilthis.ToPercentEncoding(arg0, exclude, include)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:357
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromAce(const QByteArray &)
func (this *QUrl) FromAce(arg0 *QByteArray) *QString /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl7fromAceERK10QByteArray", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQString)
	return rv2
}
func QUrl_FromAce(arg0 *QByteArray) *QString /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromAce(arg0)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:358
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toAce(const QString &)
func (this *QUrl) ToAce(arg0 *QString) *QByteArray /*123*/ {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl5toAceERK7QString", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
	// return rv
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QUrl_ToAce(arg0 *QString) *QByteArray /*123*/ {
	var nilthis *QUrl
	rv := nilthis.ToAce(arg0)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:363
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setIdnWhitelist(const QStringList &)
func (this *QUrl) SetIdnWhitelist(arg0 *QStringList) {
	var convArg0 = arg0.GetCthis()
	rv, err := ffiqt.InvokeQtFunc6("_ZN4QUrl15setIdnWhitelistERK11QStringList", ffiqt.FFI_TYPE_POINTER, convArg0)
	gopp.ErrPrint(err, rv)
}
func QUrl_SetIdnWhitelist(arg0 *QStringList) {
	var nilthis *QUrl
	nilthis.SetIdnWhitelist(arg0)
}

type QUrl__ParsingMode = int

const QUrl__TolerantMode QUrl__ParsingMode = 0
const QUrl__StrictMode QUrl__ParsingMode = 1
const QUrl__DecodedMode QUrl__ParsingMode = 2

type QUrl__UrlFormattingOption = int

const QUrl__None QUrl__UrlFormattingOption = 0
const QUrl__RemoveScheme QUrl__UrlFormattingOption = 1
const QUrl__RemovePassword QUrl__UrlFormattingOption = 2
const QUrl__RemoveUserInfo QUrl__UrlFormattingOption = 6
const QUrl__RemovePort QUrl__UrlFormattingOption = 8
const QUrl__RemoveAuthority QUrl__UrlFormattingOption = 30
const QUrl__RemovePath QUrl__UrlFormattingOption = 32
const QUrl__RemoveQuery QUrl__UrlFormattingOption = 64
const QUrl__RemoveFragment QUrl__UrlFormattingOption = 128
const QUrl__PreferLocalFile QUrl__UrlFormattingOption = 512
const QUrl__StripTrailingSlash QUrl__UrlFormattingOption = 1024
const QUrl__RemoveFilename QUrl__UrlFormattingOption = 2048
const QUrl__NormalizePathSegments QUrl__UrlFormattingOption = 4096

type QUrl__ComponentFormattingOption = int

const QUrl__PrettyDecoded QUrl__ComponentFormattingOption = 0
const QUrl__EncodeSpaces QUrl__ComponentFormattingOption = 1048576
const QUrl__EncodeUnicode QUrl__ComponentFormattingOption = 2097152
const QUrl__EncodeDelimiters QUrl__ComponentFormattingOption = 12582912
const QUrl__EncodeReserved QUrl__ComponentFormattingOption = 16777216
const QUrl__DecodeReserved QUrl__ComponentFormattingOption = 33554432
const QUrl__FullyEncoded QUrl__ComponentFormattingOption = 32505856
const QUrl__FullyDecoded QUrl__ComponentFormattingOption = 133169152

type QUrl__UserInputResolutionOption = int

const QUrl__DefaultResolution QUrl__UserInputResolutionOption = 0
const QUrl__AssumeLocalFile QUrl__UserInputResolutionOption = 1

//  body block end
