package qtcore

// /usr/include/qt/QtCore/qurl.h
// #include <qurl.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 8
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QUrl struct {
	*qtrt.CObject
}
type QUrl_ITF interface {
	QUrl_PTR() *QUrl
}

func (ptr *QUrl) QUrl_PTR() *QUrl { return ptr }

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
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:182
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrl(const QString &, enum QUrl::ParsingMode)
func NewQUrl_1(url string, mode int) *QUrl {
	var tmpArg0 = NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlC2ERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:182
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrl(const QString &, enum QUrl::ParsingMode)
func NewQUrl_1_(url string) *QUrl {
	var tmpArg0 = NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlC2ERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	gothis := NewQUrlFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:178
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl & operator=(const QUrl &)
func (this *QUrl) Operator_equal(copy QUrl_ITF) *QUrl {
	var convArg0 unsafe.Pointer
	if copy != nil && copy.QUrl_PTR() != nil {
		convArg0 = copy.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:183
// index:1
// Public Visibility=Default Availability=Available
// [8] QUrl & operator=(const QString &)
func (this *QUrl) Operator_equal_1(url string) *QUrl {
	var tmpArg0 = NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlaSERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:188
// index:2
// Public inline Visibility=Default Availability=Available
// [8] QUrl & operator=(QUrl &&)
func (this *QUrl) Operator_equal_2(other unsafe.Pointer /*333*/) *QUrl {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), other)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:191
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QUrl()
func DeleteQUrl(this *QUrl) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrlD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qurl.h:193
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void swap(QUrl &)
func (this *QUrl) Swap(other QUrl_ITF) {
	var convArg0 unsafe.Pointer
	if other != nil && other.QUrl_PTR() != nil {
		convArg0 = other.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl4swapERS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUrl(url string, mode int) {
	var tmpArg0 = NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUrl__(url string) {
	var tmpArg0 = NewQString_5(url)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:202
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromEncoded(const QByteArray &, enum QUrl::ParsingMode)
func (this *QUrl) FromEncoded(url QByteArray_ITF, mode int) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QByteArray_PTR() != nil {
		convArg0 = url.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11fromEncodedERK10QByteArrayNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromEncoded(url QByteArray_ITF, mode int) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromEncoded(url, mode)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:202
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromEncoded(const QByteArray &, enum QUrl::ParsingMode)
func (this *QUrl) FromEncoded__(url QByteArray_ITF) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QByteArray_PTR() != nil {
		convArg0 = url.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11fromEncodedERK10QByteArrayNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:210
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromUserInput(const QString &)
func (this *QUrl) FromUserInput(userInput string) *QUrl /*123*/ {
	var tmpArg0 = NewQString_5(userInput)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromUserInput(userInput string) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromUserInput(userInput)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:212
// index:1
// Public static Visibility=Default Availability=Available
// [8] QUrl fromUserInput(const QString &, const QString &, QUrl::UserInputResolutionOptions)
func (this *QUrl) FromUserInput_1(userInput string, workingDirectory string, options int) *QUrl /*123*/ {
	var tmpArg0 = NewQString_5(userInput)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(workingDirectory)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QStringS2_6QFlagsINS_25UserInputResolutionOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromUserInput_1(userInput string, workingDirectory string, options int) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromUserInput_1(userInput, workingDirectory, options)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:212
// index:1
// Public static Visibility=Default Availability=Available
// [8] QUrl fromUserInput(const QString &, const QString &, QUrl::UserInputResolutionOptions)
func (this *QUrl) FromUserInput_1_(userInput string, workingDirectory string) *QUrl /*123*/ {
	var tmpArg0 = NewQString_5(userInput)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(workingDirectory)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QUrl::UserInputResolutionOptions=Typedef, QUrl::UserInputResolutionOptions=Typedef, QFlags<QUrl::UserInputResolutionOption>
	options := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QStringS2_6QFlagsINS_25UserInputResolutionOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:215
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isValid() const
func (this *QUrl) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:216
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const
func (this *QUrl) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qurl.h:218
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const
func (this *QUrl) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()
func (this *QUrl) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScheme(const QString &)
func (this *QUrl) SetScheme(scheme string) {
	var tmpArg0 = NewQString_5(scheme)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl9setSchemeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] QString scheme() const
func (this *QUrl) Scheme() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl6schemeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qurl.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAuthority(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetAuthority(authority string, mode int) {
	var tmpArg0 = NewQString_5(authority)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl12setAuthorityERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAuthority(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetAuthority__(authority string) {
	var tmpArg0 = NewQString_5(authority)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl12setAuthorityERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserInfo(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserInfo(userInfo string, mode int) {
	var tmpArg0 = NewQString_5(userInfo)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserInfoERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserInfo(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserInfo__(userInfo string) {
	var tmpArg0 = NewQString_5(userInfo)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserInfoERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserName(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserName(userName string, mode int) {
	var tmpArg0 = NewQString_5(userName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserNameERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserName(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetUserName__(userName string) {
	var tmpArg0 = NewQString_5(userName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserNameERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPassword(password string, mode int) {
	var tmpArg0 = NewQString_5(password)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setPasswordERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPassword__(password string) {
	var tmpArg0 = NewQString_5(password)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setPasswordERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetHost(host string, mode int) {
	var tmpArg0 = NewQString_5(host)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setHostERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetHost__(host string) {
	var tmpArg0 = NewQString_5(host)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setHostERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPort(int)
func (this *QUrl) SetPort(port int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setPortEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int port(int) const
func (this *QUrl) Port(defaultPort int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl4portEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultPort)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qurl.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int port(int) const
func (this *QUrl) Port__() int {
	// arg: 0, int=Int, =Invalid,
	defaultPort := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl4portEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultPort)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qurl.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPath(path string, mode int) {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setPathERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetPath__(path string) {
	var tmpArg0 = NewQString_5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setPathERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:249
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasQuery() const
func (this *QUrl) HasQuery() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl8hasQueryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetQuery(query string, mode int) {
	var tmpArg0 = NewQString_5(query)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl8setQueryERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetQuery__(query string) {
	var tmpArg0 = NewQString_5(query)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl8setQueryERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:251
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QUrlQuery &)
func (this *QUrl) SetQuery_1(query QUrlQuery_ITF) {
	var convArg0 unsafe.Pointer
	if query != nil && query.QUrlQuery_PTR() != nil {
		convArg0 = query.QUrlQuery_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl8setQueryERK9QUrlQuery", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:254
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasFragment() const
func (this *QUrl) HasFragment() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11hasFragmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFragment(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetFragment(fragment string, mode int) {
	var tmpArg0 = NewQString_5(fragment)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setFragmentERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFragment(const QString &, enum QUrl::ParsingMode)
func (this *QUrl) SetFragment__(fragment string) {
	var tmpArg0 = NewQString_5(fragment)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum,
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setFragmentERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:258
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl resolved(const QUrl &) const
func (this *QUrl) Resolved(relative QUrl_ITF) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if relative != nil && relative.QUrl_PTR() != nil {
		convArg0 = relative.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl8resolvedERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:260
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isRelative() const
func (this *QUrl) IsRelative() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl10isRelativeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isParentOf(const QUrl &) const
func (this *QUrl) IsParentOf(url QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl10isParentOfERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:263
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLocalFile() const
func (this *QUrl) IsLocalFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11isLocalFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:264
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromLocalFile(const QString &)
func (this *QUrl) FromLocalFile(localfile string) *QUrl /*123*/ {
	var tmpArg0 = NewQString_5(localfile)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl13fromLocalFileERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromLocalFile(localfile string) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromLocalFile(localfile)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:265
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toLocalFile() const
func (this *QUrl) ToLocalFile() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11toLocalFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qurl.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()
func (this *QUrl) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:268
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const
func (this *QUrl) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:270
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QUrl &) const
func (this *QUrl) Operator_less_than(url QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrlltERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:271
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QUrl &) const
func (this *QUrl) Operator_equal_equal(url QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrleqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:272
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator!=(const QUrl &) const
func (this *QUrl) Operator_not_equal(url QUrl_ITF) bool {
	var convArg0 unsafe.Pointer
	if url != nil && url.QUrl_PTR() != nil {
		convArg0 = url.QUrl_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrlneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:276
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromPercentEncoding(const QByteArray &)
func (this *QUrl) FromPercentEncoding(arg0 QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl19fromPercentEncodingERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QUrl_FromPercentEncoding(arg0 QByteArray_ITF) string {
	var nilthis *QUrl
	rv := nilthis.FromPercentEncoding(arg0)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QString &, const QByteArray &, const QByteArray &)
func (this *QUrl) ToPercentEncoding(arg0 string, exclude QByteArray_ITF, include QByteArray_ITF) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg1 = exclude.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if include != nil && include.QByteArray_PTR() != nil {
		convArg2 = include.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QUrl_ToPercentEncoding(arg0 string, exclude QByteArray_ITF, include QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QUrl
	rv := nilthis.ToPercentEncoding(arg0, exclude, include)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QString &, const QByteArray &, const QByteArray &)
func (this *QUrl) ToPercentEncoding__(arg0 string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg1 = NewQByteArray()
	// arg: 2, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg2 = NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QString &, const QByteArray &, const QByteArray &)
func (this *QUrl) ToPercentEncoding__1(arg0 string, exclude QByteArray_ITF) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg1 = exclude.QByteArray_PTR().GetCthis()
	}
	// arg: 2, const QByteArray &=LValueReference, QByteArray=Record,
	var convArg2 = NewQByteArray()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:357
// index:0
// Public static Visibility=Default Availability=Available
// [8] QString fromAce(const QByteArray &)
func (this *QUrl) FromAce(arg0 QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7fromAceERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QUrl_FromAce(arg0 QByteArray_ITF) string {
	var nilthis *QUrl
	rv := nilthis.FromAce(arg0)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:358
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toAce(const QString &)
func (this *QUrl) ToAce(arg0 string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl5toAceERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QUrl_ToAce(arg0 string) *QByteArray /*123*/ {
	var nilthis *QUrl
	rv := nilthis.ToAce(arg0)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:359
// index:0
// Public static Visibility=Default Availability=Available
// [8] QStringList idnWhitelist()
func (this *QUrl) IdnWhitelist() *QStringList /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl12idnWhitelistEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringListFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringList)
	return rv2
}
func QUrl_IdnWhitelist() *QStringList /*123*/ {
	var nilthis *QUrl
	rv := nilthis.IdnWhitelist()
	return rv
}

// /usr/include/qt/QtCore/qurl.h:363
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setIdnWhitelist(const QStringList &)
func (this *QUrl) SetIdnWhitelist(arg0 QStringList_ITF) {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStringList_PTR() != nil {
		convArg0 = arg0.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl15setIdnWhitelistERK11QStringList", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QUrl_SetIdnWhitelist(arg0 QStringList_ITF) {
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

//  keep block begin

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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end
