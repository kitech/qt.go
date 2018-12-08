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

/*

 */
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

/*
Constructs an empty QUrl object.
*/
func (*QUrl) NewForInherit() *QUrl {
	return NewQUrl()
}
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
// [-2] void QUrl(const QString &, QUrl::ParsingMode)

/*
Constructs an empty QUrl object.
*/
func (*QUrl) NewForInherit1(url string, mode int) *QUrl {
	return NewQUrl1(url, mode)
}
func NewQUrl1(url string, mode int) *QUrl {
	var tmpArg0 = NewQString5(url)
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
// [-2] void QUrl(const QString &, QUrl::ParsingMode)

/*
Constructs an empty QUrl object.
*/
func (*QUrl) NewForInherit1p(url string) *QUrl {
	return NewQUrl1p(url)
}
func NewQUrl1p(url string) *QUrl {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
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

/*

 */
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

/*

 */
func (this *QUrl) Operator_equal1(url string) *QUrl {
	var tmpArg0 = NewQString5(url)
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

/*

 */
func (this *QUrl) Operator_equal2(other unsafe.Pointer /*333*/) *QUrl {
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

/*

 */
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

/*
Swaps URL other with this URL. This operation is very fast and never fails.

This function was introduced in  Qt 4.8.
*/
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
// [-2] void setUrl(const QString &, QUrl::ParsingMode)

/*
Parses url and sets this object to that value. QUrl will automatically percent encode all characters that are not allowed in a URL and decode the percent-encoded sequences that represent an unreserved character (letters, digits, hyphens, undercores, dots and tildes). All other characters are left in their original forms.

Parses the url using the parser mode parsingMode. In TolerantMode (the default), QUrl will correct certain mistakes, notably the presence of a percent character ('%') not followed by two hexadecimal digits, and it will accept any character in any position. In StrictMode, encoding mistakes will not be tolerated and QUrl will also check that certain forbidden characters are not present in unencoded form. If an error is detected in StrictMode, isValid() will return false. The parsing mode DecodedMode is not permitted in this context and will produce a run-time warning.

See also url() and toString().
*/
func (this *QUrl) SetUrl(url string, mode int) {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:195
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUrl(const QString &, QUrl::ParsingMode)

/*
Parses url and sets this object to that value. QUrl will automatically percent encode all characters that are not allowed in a URL and decode the percent-encoded sequences that represent an unreserved character (letters, digits, hyphens, undercores, dots and tildes). All other characters are left in their original forms.

Parses the url using the parser mode parsingMode. In TolerantMode (the default), QUrl will correct certain mistakes, notably the presence of a percent character ('%') not followed by two hexadecimal digits, and it will accept any character in any position. In StrictMode, encoding mistakes will not be tolerated and QUrl will also check that certain forbidden characters are not present in unencoded form. If an error is detected in StrictMode, isValid() will return false. The parsing mode DecodedMode is not permitted in this context and will produce a run-time warning.

See also url() and toString().
*/
func (this *QUrl) SetUrlp(url string) {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:202
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromEncoded(const QByteArray &, QUrl::ParsingMode)

/*
Parses input and returns the corresponding QUrl. input is assumed to be in encoded form, containing only ASCII characters.

Parses the URL using parsingMode. See setUrl() for more information on this parameter. QUrl::DecodedMode is not permitted in this context.

See also toEncoded() and setUrl().
*/
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
// [8] QUrl fromEncoded(const QByteArray &, QUrl::ParsingMode)

/*
Parses input and returns the corresponding QUrl. input is assumed to be in encoded form, containing only ASCII characters.

Parses the URL using parsingMode. See setUrl() for more information on this parameter. QUrl::DecodedMode is not permitted in this context.

See also toEncoded() and setUrl().
*/
func (this *QUrl) FromEncodedp(url QByteArray_ITF) *QUrl /*123*/ {
	var convArg0 unsafe.Pointer
	if url != nil && url.QByteArray_PTR() != nil {
		convArg0 = url.QByteArray_PTR().GetCthis()
	}
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
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

/*
Returns a valid URL from a user supplied userInput string if one can be deducted. In the case that is not possible, an invalid QUrl() is returned.

Most applications that can browse the web, allow the user to input a URL in the form of a plain string. This string can be manually typed into a location bar, obtained from the clipboard, or passed in via command line arguments.

When the string is not already a valid URL, a best guess is performed, making various web related assumptions.

In the case the string corresponds to a valid file path on the system, a file:// URL is constructed, using QUrl::fromLocalFile().

If that is not the case, an attempt is made to turn the string into a http:// or ftp:// URL. The latter in the case the string starts with 'ftp'. The result is then passed through QUrl's tolerant parser, and in the case or success, a valid QUrl is returned, or else a QUrl().
*/
func (this *QUrl) FromUserInput(userInput string) *QUrl /*123*/ {
	var tmpArg0 = NewQString5(userInput)
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

/*
Returns a valid URL from a user supplied userInput string if one can be deducted. In the case that is not possible, an invalid QUrl() is returned.

Most applications that can browse the web, allow the user to input a URL in the form of a plain string. This string can be manually typed into a location bar, obtained from the clipboard, or passed in via command line arguments.

When the string is not already a valid URL, a best guess is performed, making various web related assumptions.

In the case the string corresponds to a valid file path on the system, a file:// URL is constructed, using QUrl::fromLocalFile().

If that is not the case, an attempt is made to turn the string into a http:// or ftp:// URL. The latter in the case the string starts with 'ftp'. The result is then passed through QUrl's tolerant parser, and in the case or success, a valid QUrl is returned, or else a QUrl().
*/
func (this *QUrl) FromUserInput1(userInput string, workingDirectory string, options int) *QUrl /*123*/ {
	var tmpArg0 = NewQString5(userInput)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(workingDirectory)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl13fromUserInputERK7QStringS2_6QFlagsINS_25UserInputResolutionOptionEE", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, options)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQUrl)
	return rv2
}
func QUrl_FromUserInput1(userInput string, workingDirectory string, options int) *QUrl /*123*/ {
	var nilthis *QUrl
	rv := nilthis.FromUserInput1(userInput, workingDirectory, options)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:212
// index:1
// Public static Visibility=Default Availability=Available
// [8] QUrl fromUserInput(const QString &, const QString &, QUrl::UserInputResolutionOptions)

/*
Returns a valid URL from a user supplied userInput string if one can be deducted. In the case that is not possible, an invalid QUrl() is returned.

Most applications that can browse the web, allow the user to input a URL in the form of a plain string. This string can be manually typed into a location bar, obtained from the clipboard, or passed in via command line arguments.

When the string is not already a valid URL, a best guess is performed, making various web related assumptions.

In the case the string corresponds to a valid file path on the system, a file:// URL is constructed, using QUrl::fromLocalFile().

If that is not the case, an attempt is made to turn the string into a http:// or ftp:// URL. The latter in the case the string starts with 'ftp'. The result is then passed through QUrl's tolerant parser, and in the case or success, a valid QUrl is returned, or else a QUrl().
*/
func (this *QUrl) FromUserInput1p(userInput string, workingDirectory string) *QUrl /*123*/ {
	var tmpArg0 = NewQString5(userInput)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(workingDirectory)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QUrl::UserInputResolutionOptions=Typedef, QUrl::UserInputResolutionOptions=Typedef, QFlags<QUrl::UserInputResolutionOption>, Unexposed
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

/*
Returns true if the URL is non-empty and valid; otherwise returns false.

The URL is run through a conformance test. Every part of the URL must conform to the standard encoding rules of the URI standard for the URL to be reported as valid.


  bool checkUrl(const QUrl &url) {
      if (!url.isValid()) {
          qDebug("Invalid URL: %s", qUtf8Printable(url.toString()));
          return false;
      }

      return true;
  }
*/
func (this *QUrl) IsValid() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl7isValidEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:216
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns an error message if the last operation that modified this QUrl object ran into a parsing error. If no error was detected, this function returns an empty string and isValid() returns true.

The error message returned by this function is technical in nature and may not be understood by end users. It is mostly useful to developers trying to understand why QUrl will not accept some input.

This function was introduced in  Qt 4.2.

See also QUrl::ParsingMode.
*/
func (this *QUrl) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qurl.h:218
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if the URL has no data; otherwise returns false.

See also clear().
*/
func (this *QUrl) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:219
// index:0
// Public Visibility=Default Availability=Available
// [-2] void clear()

/*
Resets the content of the QUrl. After calling this function, the QUrl is equal to one that has been constructed with the default empty constructor.

See also isEmpty().
*/
func (this *QUrl) Clear() {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl5clearEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:221
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setScheme(const QString &)

/*
Sets the scheme of the URL to scheme. As a scheme can only contain ASCII characters, no conversion or decoding is done on the input. It must also start with an ASCII letter.

The scheme describes the type (or protocol) of the URL. It's represented by one or more ASCII characters at the start the URL.

A scheme is strictly RFC 3986-compliant: scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )

The following example shows a URL where the scheme is "ftp":



To set the scheme, the following call is used:


  QUrl url;
  url.setScheme("ftp");



The scheme can also be empty, in which case the URL is interpreted as relative.

See also scheme() and isRelative().
*/
func (this *QUrl) SetScheme(scheme string) {
	var tmpArg0 = NewQString5(scheme)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl9setSchemeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:222
// index:0
// Public Visibility=Default Availability=Available
// [8] QString scheme() const

/*
Returns the scheme of the URL. If an empty string is returned, this means the scheme is undefined and the URL is then relative.

The scheme can only contain US-ASCII letters or digits, which means it cannot contain any character that would otherwise require encoding. Additionally, schemes are always returned in lowercase form.

See also setScheme() and isRelative().
*/
func (this *QUrl) Scheme() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl6schemeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qurl.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAuthority(const QString &, QUrl::ParsingMode)

/*
Sets the authority of the URL to authority.

The authority of a URL is the combination of user info, a host name and a port. All of these elements are optional; an empty authority is therefore valid.

The user info and host are separated by a '@', and the host and port are separated by a ':'. If the user info is empty, the '@' must be omitted; although a stray ':' is permitted if the port is empty.

The following example shows a valid authority string:



The authority data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters.

This function does not allow mode to be QUrl::DecodedMode. To set fully decoded data, call setUserName(), setPassword(), setHost() and setPort() individually.

See also authority(), setUserInfo(), setHost(), and setPort().
*/
func (this *QUrl) SetAuthority(authority string, mode int) {
	var tmpArg0 = NewQString5(authority)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl12setAuthorityERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:224
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setAuthority(const QString &, QUrl::ParsingMode)

/*
Sets the authority of the URL to authority.

The authority of a URL is the combination of user info, a host name and a port. All of these elements are optional; an empty authority is therefore valid.

The user info and host are separated by a '@', and the host and port are separated by a ':'. If the user info is empty, the '@' must be omitted; although a stray ':' is permitted if the port is empty.

The following example shows a valid authority string:



The authority data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters.

This function does not allow mode to be QUrl::DecodedMode. To set fully decoded data, call setUserName(), setPassword(), setHost() and setPort() individually.

See also authority(), setUserInfo(), setHost(), and setPort().
*/
func (this *QUrl) SetAuthorityp(authority string) {
	var tmpArg0 = NewQString5(authority)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl12setAuthorityERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserInfo(const QString &, QUrl::ParsingMode)

/*
Sets the user info of the URL to userInfo. The user info is an optional part of the authority of the URL, as described in setAuthority().

The user info consists of a user name and optionally a password, separated by a ':'. If the password is empty, the colon must be omitted. The following example shows a valid user info string:



The userInfo data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters.

This function does not allow mode to be QUrl::DecodedMode. To set fully decoded data, call setUserName() and setPassword() individually.

See also userInfo(), setUserName(), setPassword(), and setAuthority().
*/
func (this *QUrl) SetUserInfo(userInfo string, mode int) {
	var tmpArg0 = NewQString5(userInfo)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserInfoERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:227
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserInfo(const QString &, QUrl::ParsingMode)

/*
Sets the user info of the URL to userInfo. The user info is an optional part of the authority of the URL, as described in setAuthority().

The user info consists of a user name and optionally a password, separated by a ':'. If the password is empty, the colon must be omitted. The following example shows a valid user info string:



The userInfo data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters.

This function does not allow mode to be QUrl::DecodedMode. To set fully decoded data, call setUserName() and setPassword() individually.

See also userInfo(), setUserName(), setPassword(), and setAuthority().
*/
func (this *QUrl) SetUserInfop(userInfo string) {
	var tmpArg0 = NewQString5(userInfo)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserInfoERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserName(const QString &, QUrl::ParsingMode)

/*
Sets the URL's user name to userName. The userName is part of the user info element in the authority of the URL, as described in setUserInfo().

The userName data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the user name from a data source which is not a URL, such as a password dialog shown to the user or with a user name obtained by calling userName() with the QUrl::FullyDecoded formatting option.

See also userName() and setUserInfo().
*/
func (this *QUrl) SetUserName(userName string, mode int) {
	var tmpArg0 = NewQString5(userName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserNameERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:230
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setUserName(const QString &, QUrl::ParsingMode)

/*
Sets the URL's user name to userName. The userName is part of the user info element in the authority of the URL, as described in setUserInfo().

The userName data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the user name from a data source which is not a URL, such as a password dialog shown to the user or with a user name obtained by calling userName() with the QUrl::FullyDecoded formatting option.

See also userName() and setUserInfo().
*/
func (this *QUrl) SetUserNamep(userName string) {
	var tmpArg0 = NewQString5(userName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setUserNameERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &, QUrl::ParsingMode)

/*
Sets the URL's password to password. The password is part of the user info element in the authority of the URL, as described in setUserInfo().

The password data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the password from a data source which is not a URL, such as a password dialog shown to the user or with a password obtained by calling password() with the QUrl::FullyDecoded formatting option.

See also password() and setUserInfo().
*/
func (this *QUrl) SetPassword(password string, mode int) {
	var tmpArg0 = NewQString5(password)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setPasswordERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:233
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPassword(const QString &, QUrl::ParsingMode)

/*
Sets the URL's password to password. The password is part of the user info element in the authority of the URL, as described in setUserInfo().

The password data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the password from a data source which is not a URL, such as a password dialog shown to the user or with a password obtained by calling password() with the QUrl::FullyDecoded formatting option.

See also password() and setUserInfo().
*/
func (this *QUrl) SetPasswordp(password string) {
	var tmpArg0 = NewQString5(password)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setPasswordERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, QUrl::ParsingMode)

/*
Sets the host of the URL to host. The host is part of the authority.

The host data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

Note that, in all cases, the result of the parsing must be a valid hostname according to STD 3 rules, as modified by the Internationalized Resource Identifiers specification (RFC 3987). Invalid hostnames are not permitted and will cause isValid() to become false.

See also host() and setAuthority().
*/
func (this *QUrl) SetHost(host string, mode int) {
	var tmpArg0 = NewQString5(host)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setHostERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:236
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setHost(const QString &, QUrl::ParsingMode)

/*
Sets the host of the URL to host. The host is part of the authority.

The host data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

Note that, in all cases, the result of the parsing must be a valid hostname according to STD 3 rules, as modified by the Internationalized Resource Identifiers specification (RFC 3987). Invalid hostnames are not permitted and will cause isValid() to become false.

See also host() and setAuthority().
*/
func (this *QUrl) SetHostp(host string) {
	var tmpArg0 = NewQString5(host)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setHostERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:242
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPort(int)

/*
Sets the port of the URL to port. The port is part of the authority of the URL, as described in setAuthority().

port must be between 0 and 65535 inclusive. Setting the port to -1 indicates that the port is unspecified.

See also port().
*/
func (this *QUrl) SetPort(port int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setPortEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), port)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int port(int) const

/*
Returns the port of the URL, or defaultPort if the port is unspecified.

Example:


  QTcpSocket sock;
  sock.connectToHost(url.host(), url.port(80));



This function was introduced in  Qt 4.1.

See also setPort().
*/
func (this *QUrl) Port(defaultPort int) int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl4portEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultPort)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qurl.h:243
// index:0
// Public Visibility=Default Availability=Available
// [4] int port(int) const

/*
Returns the port of the URL, or defaultPort if the port is unspecified.

Example:


  QTcpSocket sock;
  sock.connectToHost(url.host(), url.port(80));



This function was introduced in  Qt 4.1.

See also setPort().
*/
func (this *QUrl) Portp() int {
	// arg: 0, int=Int, =Invalid, , Invalid
	defaultPort := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl4portEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), defaultPort)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qurl.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &, QUrl::ParsingMode)

/*
Sets the path of the URL to path. The path is the part of the URL that comes after the authority but before the query string.



For non-hierarchical schemes, the path will be everything following the scheme declaration, as in the following example:



The path data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the path from a data source which is not a URL, such as a dialog shown to the user or with a path obtained by calling path() with the QUrl::FullyDecoded formatting option.

See also path().
*/
func (this *QUrl) SetPath(path string, mode int) {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setPathERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:245
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setPath(const QString &, QUrl::ParsingMode)

/*
Sets the path of the URL to path. The path is the part of the URL that comes after the authority but before the query string.



For non-hierarchical schemes, the path will be everything following the scheme declaration, as in the following example:



The path data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode (the default), all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the path from a data source which is not a URL, such as a dialog shown to the user or with a path obtained by calling path() with the QUrl::FullyDecoded formatting option.

See also path().
*/
func (this *QUrl) SetPathp(path string) {
	var tmpArg0 = NewQString5(path)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7setPathERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:249
// index:0
// Public Visibility=Default Availability=Available
// [1] bool hasQuery() const

/*
Returns true if this URL contains a Query (i.e., if ? was seen on it).

This function was introduced in  Qt 4.2.

See also setQuery(), query(), and hasFragment().
*/
func (this *QUrl) HasQuery() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl8hasQueryEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &, QUrl::ParsingMode)

/*
Sets the query string of the URL to query.

This function is useful if you need to pass a query string that does not fit into the key-value pattern, or that uses a different scheme for encoding special characters than what is suggested by QUrl.

Passing a value of QString() to query (a null QString) unsets the query completely. However, passing a value of QString("") will set the query to an empty value, as if the original URL had a lone "?".

The query data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

Query strings often contain percent-encoded sequences, so use of DecodedMode is discouraged. One special sequence to be aware of is that of the plus character ('+'). QUrl does not convert spaces to plus characters, even though HTML forms posted by web browsers do. In order to represent an actual plus character in a query, the sequence "%2B" is usually used. This function will leave "%2B" sequences untouched in TolerantMode or StrictMode.

See also query() and hasQuery().
*/
func (this *QUrl) SetQuery(query string, mode int) {
	var tmpArg0 = NewQString5(query)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl8setQueryERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:250
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QString &, QUrl::ParsingMode)

/*
Sets the query string of the URL to query.

This function is useful if you need to pass a query string that does not fit into the key-value pattern, or that uses a different scheme for encoding special characters than what is suggested by QUrl.

Passing a value of QString() to query (a null QString) unsets the query completely. However, passing a value of QString("") will set the query to an empty value, as if the original URL had a lone "?".

The query data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

Query strings often contain percent-encoded sequences, so use of DecodedMode is discouraged. One special sequence to be aware of is that of the plus character ('+'). QUrl does not convert spaces to plus characters, even though HTML forms posted by web browsers do. In order to represent an actual plus character in a query, the sequence "%2B" is usually used. This function will leave "%2B" sequences untouched in TolerantMode or StrictMode.

See also query() and hasQuery().
*/
func (this *QUrl) SetQueryp(query string) {
	var tmpArg0 = NewQString5(query)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl8setQueryERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:251
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setQuery(const QUrlQuery &)

/*
Sets the query string of the URL to query.

This function is useful if you need to pass a query string that does not fit into the key-value pattern, or that uses a different scheme for encoding special characters than what is suggested by QUrl.

Passing a value of QString() to query (a null QString) unsets the query completely. However, passing a value of QString("") will set the query to an empty value, as if the original URL had a lone "?".

The query data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

Query strings often contain percent-encoded sequences, so use of DecodedMode is discouraged. One special sequence to be aware of is that of the plus character ('+'). QUrl does not convert spaces to plus characters, even though HTML forms posted by web browsers do. In order to represent an actual plus character in a query, the sequence "%2B" is usually used. This function will leave "%2B" sequences untouched in TolerantMode or StrictMode.

See also query() and hasQuery().
*/
func (this *QUrl) SetQuery1(query QUrlQuery_ITF) {
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

/*
Returns true if this URL contains a fragment (i.e., if # was seen on it).

This function was introduced in  Qt 4.2.

See also fragment() and setFragment().
*/
func (this *QUrl) HasFragment() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11hasFragmentEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFragment(const QString &, QUrl::ParsingMode)

/*
Sets the fragment of the URL to fragment. The fragment is the last part of the URL, represented by a '#' followed by a string of characters. It is typically used in HTTP for referring to a certain link or point on a page:



The fragment is sometimes also referred to as the URL "reference".

Passing an argument of QString() (a null QString) will unset the fragment. Passing an argument of QString("") (an empty but not null QString) will set the fragment to an empty string (as if the original URL had a lone "#").

The fragment data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the fragment from a data source which is not a URL or with a fragment obtained by calling fragment() with the QUrl::FullyDecoded formatting option.

See also fragment() and hasFragment().
*/
func (this *QUrl) SetFragment(fragment string, mode int) {
	var tmpArg0 = NewQString5(fragment)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setFragmentERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:256
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFragment(const QString &, QUrl::ParsingMode)

/*
Sets the fragment of the URL to fragment. The fragment is the last part of the URL, represented by a '#' followed by a string of characters. It is typically used in HTTP for referring to a certain link or point on a page:



The fragment is sometimes also referred to as the URL "reference".

Passing an argument of QString() (a null QString) will unset the fragment. Passing an argument of QString("") (an empty but not null QString) will set the fragment to an empty string (as if the original URL had a lone "#").

The fragment data is interpreted according to mode: in StrictMode, any '%' characters must be followed by exactly two hexadecimal characters and some characters (including space) are not allowed in undecoded form. In TolerantMode, all characters are accepted in undecoded form and the tolerant parser will correct stray '%' not followed by two hex characters. In DecodedMode, '%' stand for themselves and encoded characters are not possible.

QUrl::DecodedMode should be used when setting the fragment from a data source which is not a URL or with a fragment obtained by calling fragment() with the QUrl::FullyDecoded formatting option.

See also fragment() and hasFragment().
*/
func (this *QUrl) SetFragmentp(fragment string) {
	var tmpArg0 = NewQString5(fragment)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl11setFragmentERK7QStringNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:258
// index:0
// Public Visibility=Default Availability=Available
// [8] QUrl resolved(const QUrl &) const

/*
Returns the result of the merge of this URL with relative. This URL is used as a base to convert relative to an absolute URL.

If relative is not a relative URL, this function will return relative directly. Otherwise, the paths of the two URLs are merged, and the new URL returned has the scheme and authority of the base URL, but with the merged path, as in the following example:


  QUrl baseUrl("http://qt.digia.com/Support/");
  QUrl relativeUrl("../Product/Library/");
  qDebug(baseUrl.resolved(relativeUrl).toString());
  // prints "http://qt.digia.com/Product/Library/"



Calling resolved() with ".." returns a QUrl whose directory is one level higher than the original. Similarly, calling resolved() with "../.." removes two levels from the path. If relative is "/", the path becomes "/".

See also isRelative().
*/
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

/*
Returns true if the URL is relative; otherwise returns false. A URL is relative reference if its scheme is undefined; this function is therefore equivalent to calling scheme().isEmpty().

Relative references are defined in RFC 3986 section 4.2.

See also Relative URLs vs Relative Paths.
*/
func (this *QUrl) IsRelative() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl10isRelativeEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:261
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isParentOf(const QUrl &) const

/*
Returns true if this URL is a parent of childUrl. childUrl is a child of this URL if the two URLs share the same scheme and authority, and this URL's path is a parent of the path of childUrl.
*/
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

/*
Returns true if this URL is pointing to a local file path. A URL is a local file path if the scheme is "file".

Note that this function considers URLs with hostnames to be local file paths, even if the eventual file path cannot be opened with QFile::open().

This function was introduced in  Qt 4.8.

See also fromLocalFile() and toLocalFile().
*/
func (this *QUrl) IsLocalFile() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11isLocalFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:264
// index:0
// Public static Visibility=Default Availability=Available
// [8] QUrl fromLocalFile(const QString &)

/*
Returns a QUrl representation of localFile, interpreted as a local file. This function accepts paths separated by slashes as well as the native separator for this platform.

This function also accepts paths with a doubled leading slash (or backslash) to indicate a remote file, as in "//servername/path/to/file.txt". Note that only certain platforms can actually open this file using QFile::open().

An empty localFile leads to an empty URL (since Qt 5.4).


  qDebug() << QUrl::fromLocalFile("file.txt");            // QUrl("file:file.txt")
  qDebug() << QUrl::fromLocalFile("/home/user/file.txt"); // QUrl("file:///home/user/file.txt")
  qDebug() << QUrl::fromLocalFile("file:file.txt");       // doesn't make sense; expects path, not url with scheme



In the first line in snippet above, a file URL is constructed from a local, relative path. A file URL with a relative path only makes sense if there is a base URL to resolve it against. For example:


  QUrl url = QUrl::fromLocalFile("file.txt");
  QUrl baseUrl = QUrl("file:/home/user/");
  // wrong: prints QUrl("file:file.txt"), as url already has a scheme
  qDebug() << baseUrl.resolved(url);



To resolve such a URL, it's necessary to remove the scheme beforehand:


  // correct: prints QUrl("file:///home/user/file.txt")
  url.setScheme(QString());
  qDebug() << baseUrl.resolved(url);



For this reason, it is better to use a relative URL (that is, no scheme) for relative file paths:


  QUrl url = QUrl("file.txt");
  QUrl baseUrl = QUrl("file:/home/user/");
  // prints QUrl("file:///home/user/file.txt")
  qDebug() << baseUrl.resolved(url);



See also toLocalFile(), isLocalFile(), and QDir::toNativeSeparators().
*/
func (this *QUrl) FromLocalFile(localfile string) *QUrl /*123*/ {
	var tmpArg0 = NewQString5(localfile)
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

/*
Returns the path of this URL formatted as a local file path. The path returned will use forward slashes, even if it was originally created from one with backslashes.

If this URL contains a non-empty hostname, it will be encoded in the returned value in the form found on SMB networks (for example, "//servername/path/to/file.txt").


  qDebug() << QUrl("file:file.txt").toLocalFile();            // "file:file.txt"
  qDebug() << QUrl("file:/home/user/file.txt").toLocalFile(); // "file:///home/user/file.txt"
  qDebug() << QUrl("file.txt").toLocalFile();                 // ""; wasn't a local file as it had no scheme



Note: if the path component of this URL contains a non-UTF-8 binary sequence (such as %80), the behaviour of this function is undefined.

See also fromLocalFile() and isLocalFile().
*/
func (this *QUrl) ToLocalFile() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl11toLocalFileEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qurl.h:267
// index:0
// Public Visibility=Default Availability=Available
// [-2] void detach()

/*

 */
func (this *QUrl) Detach() {
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl6detachEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:268
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isDetached() const

/*

 */
func (this *QUrl) IsDetached() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK4QUrl10isDetachedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qurl.h:270
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator<(const QUrl &) const

/*

 */
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

/*

 */
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

/*

 */
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

/*
Returns a decoded copy of input. input is first decoded from percent encoding, then converted from UTF-8 to unicode.

Note: Given invalid input (such as a string containing the sequence "%G5", which is not a valid hexadecimal number) the output will be invalid as well. As an example: the sequence "%G5" could be decoded to 'W'.
*/
func (this *QUrl) FromPercentEncoding(arg0 QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl19fromPercentEncodingERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
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

/*
Returns an encoded copy of input. input is first converted to UTF-8, and all ASCII-characters that are not in the unreserved group are percent encoded. To prevent characters from being percent encoded pass them to exclude. To force characters to be percent encoded pass them to include.

Unreserved is defined as: ALPHA / DIGIT / "-" / "." / "_" / "~"


  QByteArray ba = QUrl::toPercentEncoding("{a fishy string?}", "{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"
*/
func (this *QUrl) ToPercentEncoding(arg0 string, exclude QByteArray_ITF, include_ QByteArray_ITF) *QByteArray /*123*/ {
	var tmpArg0 = NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg1 = exclude.QByteArray_PTR().GetCthis()
	}
	var convArg2 unsafe.Pointer
	if include_ != nil && include_.QByteArray_PTR() != nil {
		convArg2 = include_.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl17toPercentEncodingERK7QStringRK10QByteArrayS5_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}
func QUrl_ToPercentEncoding(arg0 string, exclude QByteArray_ITF, include_ QByteArray_ITF) *QByteArray /*123*/ {
	var nilthis *QUrl
	rv := nilthis.ToPercentEncoding(arg0, exclude, include_)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:277
// index:0
// Public static Visibility=Default Availability=Available
// [8] QByteArray toPercentEncoding(const QString &, const QByteArray &, const QByteArray &)

/*
Returns an encoded copy of input. input is first converted to UTF-8, and all ASCII-characters that are not in the unreserved group are percent encoded. To prevent characters from being percent encoded pass them to exclude. To force characters to be percent encoded pass them to include.

Unreserved is defined as: ALPHA / DIGIT / "-" / "." / "_" / "~"


  QByteArray ba = QUrl::toPercentEncoding("{a fishy string?}", "{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"
*/
func (this *QUrl) ToPercentEncodingp(arg0 string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
	var convArg1 = NewQByteArray()
	// arg: 2, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
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

/*
Returns an encoded copy of input. input is first converted to UTF-8, and all ASCII-characters that are not in the unreserved group are percent encoded. To prevent characters from being percent encoded pass them to exclude. To force characters to be percent encoded pass them to include.

Unreserved is defined as: ALPHA / DIGIT / "-" / "." / "_" / "~"


  QByteArray ba = QUrl::toPercentEncoding("{a fishy string?}", "{}", "s");
  qDebug(ba.constData());
  // prints "{a fi%73hy %73tring%3F}"
*/
func (this *QUrl) ToPercentEncodingp1(arg0 string, exclude QByteArray_ITF) *QByteArray /*123*/ {
	var tmpArg0 = NewQString5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if exclude != nil && exclude.QByteArray_PTR() != nil {
		convArg1 = exclude.QByteArray_PTR().GetCthis()
	}
	// arg: 2, const QByteArray &=LValueReference, QByteArray=Record, , Invalid
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

/*
Returns the Unicode form of the given domain name domain, which is encoded in the ASCII Compatible Encoding (ACE). The result of this function is considered equivalent to domain.

If the value in domain cannot be encoded, it will be converted to QString and returned.

The ASCII Compatible Encoding (ACE) is defined by RFC 3490, RFC 3491 and RFC 3492. It is part of the Internationalizing Domain Names in Applications (IDNA) specification, which allows for domain names (like "example.com") to be written using international characters.

This function was introduced in  Qt 4.2.
*/
func (this *QUrl) FromAce(arg0 QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl7fromAceERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
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

/*
Returns the ASCII Compatible Encoding of the given domain name domain. The result of this function is considered equivalent to domain.

The ASCII-Compatible Encoding (ACE) is defined by RFC 3490, RFC 3491 and RFC 3492. It is part of the Internationalizing Domain Names in Applications (IDNA) specification, which allows for domain names (like "example.com") to be written using international characters.

This function returns an empty QByteArray if domain is not a valid hostname. Note, in particular, that IPv6 literals are not valid domain names.

This function was introduced in  Qt 4.2.
*/
func (this *QUrl) ToAce(arg0 string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString5(arg0)
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

/*
Returns the current whitelist of top-level domains that are allowed to have non-ASCII characters in their compositions.

See setIdnWhitelist() for the rationale of this list.

This function was introduced in  Qt 4.2.

See also setIdnWhitelist().
*/
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

// /usr/include/qt/QtCore/qurl.h:361
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> fromStringList(const QStringList &, QUrl::ParsingMode)

/*
Converts a list of strings representing urls into a list of urls, using QUrl(str, mode). Note that this means all strings must be urls, not for instance local paths.

This function was introduced in  Qt 5.1.
*/
func (this *QUrl) FromStringList(uris QStringList_ITF, mode int) *QUrlList /*lll*/ {
	var convArg0 unsafe.Pointer
	if uris != nil && uris.QStringList_PTR() != nil {
		convArg0 = uris.QStringList_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl14fromStringListERK11QStringListNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}
func QUrl_FromStringList(uris QStringList_ITF, mode int) *QUrlList /*lll*/ {
	var nilthis *QUrl
	rv := nilthis.FromStringList(uris, mode)
	return rv
}

// /usr/include/qt/QtCore/qurl.h:361
// index:0
// Public static Visibility=Default Availability=Available
// [-2] QList<QUrl> fromStringList(const QStringList &, QUrl::ParsingMode)

/*
Converts a list of strings representing urls into a list of urls, using QUrl(str, mode). Note that this means all strings must be urls, not for instance local paths.

This function was introduced in  Qt 5.1.
*/
func (this *QUrl) FromStringListp(uris QStringList_ITF) *QUrlList /*lll*/ {
	var convArg0 unsafe.Pointer
	if uris != nil && uris.QStringList_PTR() != nil {
		convArg0 = uris.QStringList_PTR().GetCthis()
	}
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN4QUrl14fromStringListERK11QStringListNS_11ParsingModeE", qtrt.FFI_TYPE_POINTER, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQUrlListFromPointer(unsafe.Pointer(uintptr(rv))) //5552
	return rv2
}

// /usr/include/qt/QtCore/qurl.h:363
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setIdnWhitelist(const QStringList &)

/*
Sets the whitelist of Top-Level Domains (TLDs) that are allowed to have non-ASCII characters in domains to the value of list.

Note that if you call this function, you need to do so before you start any threads that might access idnWhitelist().

Qt comes with a default list that contains the Internet top-level domains that have published support for Internationalized Domain Names (IDNs) and rules to guarantee that no deception can happen between similarly-looking characters (such as the Latin lowercase letter 'a' and the Cyrillic equivalent, which in most fonts are visually identical).

This list is periodically maintained, as registrars publish new rules.

This function is provided for those who need to manipulate the list, in order to add or remove a TLD. It is not recommended to change its value for purposes other than testing, as it may expose users to security risks.

This function was introduced in  Qt 4.2.

See also idnWhitelist().
*/
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

/*
The parsing mode controls the way QUrl parses strings.



In TolerantMode, the parser has the following behaviour:


Spaces and "%20": unencoded space characters will be accepted and will be treated as equivalent to "%20".
Single "%" characters: Any occurrences of a percent character "%" not followed by exactly two hexadecimal characters (e.g., "13% coverage.html") will be replaced by "%25". Note that one lone "%" character will trigger the correction mode for all percent characters.
Reserved and unreserved characters: An encoded URL should only contain a few characters as literals; all other characters should be percent-encoded. In TolerantMode, these characters will be accepted if they are found in the URL: space / double-quote / "<" / ">" / "" / "^" / "`" / "{" / "|" / "}" Those same characters can be decoded again by passing QUrl::DecodeReserved to toString() or toEncoded(). In the getters of individual components, those characters are often returned in decoded form.


When in StrictMode, if a parsing error is found, isValid() will return false and errorString() will return a message describing the error. If more than one error is detected, it is undefined which error gets reported.

Note that TolerantMode is not usually enough for parsing user input, which often contains more errors and expectations than the parser can deal with. When dealing with data coming directly from the user -- as opposed to data coming from data-transfer sources, such as other programs -- it is recommended to use fromUserInput().

See also fromUserInput(), setUrl(), toString(), toEncoded(), and QUrl::FormattingOptions.

*/
type QUrl__ParsingMode = int

// QUrl will try to correct some common errors in URLs. This mode is useful for parsing URLs coming from sources not known to be strictly standards-conforming.
const QUrl__TolerantMode QUrl__ParsingMode = 0

// Only valid URLs are accepted. This mode is useful for general URL validation.
const QUrl__StrictMode QUrl__ParsingMode = 1

// QUrl will interpret the URL component in the fully-decoded form, where percent characters stand for themselves, not as the beginning of a percent-encoded sequence. This mode is only valid for the setters setting components of a URL; it is not permitted in the QUrl constructor, in fromEncoded() or in setUrl(). For more information on this mode, see the documentation for QUrl::FullyDecoded.
const QUrl__DecodedMode QUrl__ParsingMode = 2

func (this *QUrl) ParsingModeItemName(val int) string {
	switch val {
	case QUrl__TolerantMode: // 0
		return "TolerantMode"
	case QUrl__StrictMode: // 1
		return "StrictMode"
	case QUrl__DecodedMode: // 2
		return "DecodedMode"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QUrl_ParsingModeItemName(val int) string {
	var nilthis *QUrl
	return nilthis.ParsingModeItemName(val)
}

/*


 */
type QUrl__UrlFormattingOption = int

//
const QUrl__None QUrl__UrlFormattingOption = 0

//
const QUrl__RemoveScheme QUrl__UrlFormattingOption = 1

//
const QUrl__RemovePassword QUrl__UrlFormattingOption = 2

//
const QUrl__RemoveUserInfo QUrl__UrlFormattingOption = 6

//
const QUrl__RemovePort QUrl__UrlFormattingOption = 8

//
const QUrl__RemoveAuthority QUrl__UrlFormattingOption = 30

//
const QUrl__RemovePath QUrl__UrlFormattingOption = 32

//
const QUrl__RemoveQuery QUrl__UrlFormattingOption = 64

//
const QUrl__RemoveFragment QUrl__UrlFormattingOption = 128

//
const QUrl__PreferLocalFile QUrl__UrlFormattingOption = 512

//
const QUrl__StripTrailingSlash QUrl__UrlFormattingOption = 1024

//
const QUrl__RemoveFilename QUrl__UrlFormattingOption = 2048

//
const QUrl__NormalizePathSegments QUrl__UrlFormattingOption = 4096

func (this *QUrl) UrlFormattingOptionItemName(val int) string {
	switch val {
	case QUrl__None: // 0
		return "None"
	case QUrl__RemoveScheme: // 1
		return "RemoveScheme"
	case QUrl__RemovePassword: // 2
		return "RemovePassword"
	case QUrl__RemoveUserInfo: // 6
		return "RemoveUserInfo"
	case QUrl__RemovePort: // 8
		return "RemovePort"
	case QUrl__RemoveAuthority: // 30
		return "RemoveAuthority"
	case QUrl__RemovePath: // 32
		return "RemovePath"
	case QUrl__RemoveQuery: // 64
		return "RemoveQuery"
	case QUrl__RemoveFragment: // 128
		return "RemoveFragment"
	case QUrl__PreferLocalFile: // 512
		return "PreferLocalFile"
	case QUrl__StripTrailingSlash: // 1024
		return "StripTrailingSlash"
	case QUrl__RemoveFilename: // 2048
		return "RemoveFilename"
	case QUrl__NormalizePathSegments: // 4096
		return "NormalizePathSegments"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QUrl_UrlFormattingOptionItemName(val int) string {
	var nilthis *QUrl
	return nilthis.UrlFormattingOptionItemName(val)
}

/*


 */
type QUrl__ComponentFormattingOption = int

//
const QUrl__PrettyDecoded QUrl__ComponentFormattingOption = 0

//
const QUrl__EncodeSpaces QUrl__ComponentFormattingOption = 1048576

//
const QUrl__EncodeUnicode QUrl__ComponentFormattingOption = 2097152

//
const QUrl__EncodeDelimiters QUrl__ComponentFormattingOption = 12582912

//
const QUrl__EncodeReserved QUrl__ComponentFormattingOption = 16777216

//
const QUrl__DecodeReserved QUrl__ComponentFormattingOption = 33554432

//
const QUrl__FullyEncoded QUrl__ComponentFormattingOption = 32505856

//
const QUrl__FullyDecoded QUrl__ComponentFormattingOption = 133169152

func (this *QUrl) ComponentFormattingOptionItemName(val int) string {
	switch val {
	case QUrl__PrettyDecoded: // 0
		return "PrettyDecoded"
	case QUrl__EncodeSpaces: // 1048576
		return "EncodeSpaces"
	case QUrl__EncodeUnicode: // 2097152
		return "EncodeUnicode"
	case QUrl__EncodeDelimiters: // 12582912
		return "EncodeDelimiters"
	case QUrl__EncodeReserved: // 16777216
		return "EncodeReserved"
	case QUrl__DecodeReserved: // 33554432
		return "DecodeReserved"
	case QUrl__FullyEncoded: // 32505856
		return "FullyEncoded"
	case QUrl__FullyDecoded: // 133169152
		return "FullyDecoded"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QUrl_ComponentFormattingOptionItemName(val int) string {
	var nilthis *QUrl
	return nilthis.ComponentFormattingOptionItemName(val)
}

/*


 */
type QUrl__UserInputResolutionOption = int

//
const QUrl__DefaultResolution QUrl__UserInputResolutionOption = 0

//
const QUrl__AssumeLocalFile QUrl__UserInputResolutionOption = 1

func (this *QUrl) UserInputResolutionOptionItemName(val int) string {
	switch val {
	case QUrl__DefaultResolution: // 0
		return "DefaultResolution"
	case QUrl__AssumeLocalFile: // 1
		return "AssumeLocalFile"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QUrl_UserInputResolutionOptionItemName(val int) string {
	var nilthis *QUrl
	return nilthis.UserInputResolutionOptionItemName(val)
}

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
