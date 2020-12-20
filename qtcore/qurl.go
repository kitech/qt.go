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
// extern C begin: 1
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
// size 8
type QUrl struct {
	*qtrt.CObject
}
type QUrl_ITF interface {
	QUrl_PTR() *QUrl
}

func (ptr *QUrl) QUrl_PTR() *QUrl { return ptr }

// ignore GetCthis for 0 base
// ignore SetCthis for 0 base
func QUrlFromptr(cthis Voidptr) *QUrl {
	return &QUrl{&qtrt.CObject{cthis}}
}
func (*QUrl) Fromptr(cthis Voidptr) *QUrl {
	return QUrlFromptr(cthis)
}

// /usr/include/qt/QtCore/qurl.h:176
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QUrl()

/*
 */
func (*QUrl) NewForInherit() *QUrl {
	return NewQUrl()
}
func NewQUrl() *QUrl {
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc1(887786197, "_ZN4QUrlC2Ev", qtrt.FFITY_POINTER, cthis)
	qtrt.ErrPrint(err, rv)
	gothis := QUrlFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:182
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrl(const QString &, QUrl::ParsingMode)

/*
 */
func (*QUrl) NewForInherit1(url string, mode int) *QUrl {
	return NewQUrl1(url, mode)
}
func NewQUrl1(url string, mode int) *QUrl {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc1(2077668047, "_ZN4QUrlC2ERK7QStringNS_11ParsingModeE", qtrt.FFITY_POINTER, cthis, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	gothis := QUrlFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:182
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QUrl(const QString &, QUrl::ParsingMode)

/*
 */
func (*QUrl) NewForInherit1p(url string) *QUrl {
	return NewQUrl1p(url)
}
func NewQUrl1p(url string) *QUrl {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	cthis := qtrt.Malloc(8)
	rv, err := qtrt.Qtcc1(2077668047, "_ZN4QUrlC2ERK7QStringNS_11ParsingModeE", qtrt.FFITY_POINTER, cthis, convArg0, mode)
	qtrt.ErrPrint(err, rv)
	gothis := QUrlFromptr(cthis)
	qtrt.SetFinalizer(gothis, DeleteQUrl)
	return gothis
}

// /usr/include/qt/QtCore/qurl.h:193
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setUrl(const QString &, QUrl::ParsingMode)

/*
 */
func (this *QUrl) SetUrl(url string, mode int) {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.Qtcc1(241840493, "_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qurl.h:193
// index:0
// Public Ignore Visibility=Default Availability=Available
// [-2] void setUrl(const QString &, QUrl::ParsingMode)

/*
 */
func (this *QUrl) SetUrlp(url string) {
	var tmpArg0 = NewQString5(url)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QUrl::ParsingMode=Enum, QUrl::ParsingMode=Enum, , Invalid
	mode := 0
	rv, err := qtrt.Qtcc1(241840493, "_ZN4QUrl6setUrlERK7QStringNS_11ParsingModeE", qtrt.FFITY_POINTER, this.GetCthis(), convArg0, mode)
	qtrt.ErrPrint(err, rv)
}

func DeleteQUrl(this *QUrl) {
	rv, err := qtrt.Qtcc1(0, "_ZN4QUrlD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QUrl__ParsingMode = int

//
const QUrl__TolerantMode QUrl__ParsingMode = 0

//
const QUrl__StrictMode QUrl__ParsingMode = 1

//
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

func init_unused_10025() {
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
