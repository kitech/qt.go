package qtwebenginecore

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h
// #include <qwebengineurlscheme.h>
// #include <QtWebEngineCore>

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
import "github.com/kitech/qt.go/qtqml"
import "github.com/kitech/qt.go/qtpositioning"
import "github.com/kitech/qt.go/qtwebchannel"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtquick"

//  ext block end

//  body block begin

/*

 */
type QWebEngineUrlScheme struct {
	*qtrt.CObject
}
type QWebEngineUrlScheme_ITF interface {
	QWebEngineUrlScheme_PTR() *QWebEngineUrlScheme
}

func (ptr *QWebEngineUrlScheme) QWebEngineUrlScheme_PTR() *QWebEngineUrlScheme { return ptr }

func (this *QWebEngineUrlScheme) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWebEngineUrlScheme) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWebEngineUrlSchemeFromPointer(cthis unsafe.Pointer) *QWebEngineUrlScheme {
	return &QWebEngineUrlScheme{&qtrt.CObject{cthis}}
}
func (*QWebEngineUrlScheme) NewFromPointer(cthis unsafe.Pointer) *QWebEngineUrlScheme {
	return NewQWebEngineUrlSchemeFromPointer(cthis)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:79
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineUrlScheme()

/*

 */
func (*QWebEngineUrlScheme) NewForInherit() *QWebEngineUrlScheme {
	return NewQWebEngineUrlScheme()
}
func NewQWebEngineUrlScheme() *QWebEngineUrlScheme {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlSchemeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineUrlSchemeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineUrlScheme)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:80
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QWebEngineUrlScheme(const QByteArray &)

/*

 */
func (*QWebEngineUrlScheme) NewForInherit1(name qtcore.QByteArray_ITF) *QWebEngineUrlScheme {
	return NewQWebEngineUrlScheme1(name)
}
func NewQWebEngineUrlScheme1(name qtcore.QByteArray_ITF) *QWebEngineUrlScheme {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlSchemeC2ERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQWebEngineUrlSchemeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQWebEngineUrlScheme)
	return gothis
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:83
// index:0
// Public Visibility=Default Availability=Available
// [8] QWebEngineUrlScheme & operator=(const QWebEngineUrlScheme &)

/*

 */
func (this *QWebEngineUrlScheme) Operator_equal(that QWebEngineUrlScheme_ITF) *QWebEngineUrlScheme {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineUrlScheme_PTR() != nil {
		convArg0 = that.QWebEngineUrlScheme_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlSchemeaSERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineUrlSchemeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineUrlScheme)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:86
// index:1
// Public Visibility=Default Availability=Available
// [8] QWebEngineUrlScheme & operator=(QWebEngineUrlScheme &&)

/*

 */
func (this *QWebEngineUrlScheme) Operator_equal1(that unsafe.Pointer /*333*/) *QWebEngineUrlScheme {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlSchemeaSEOS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), that)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineUrlSchemeFromPointer(unsafe.Pointer(uintptr(rv))) // 4441
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineUrlScheme)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QWebEngineUrlScheme()

/*

 */
func DeleteQWebEngineUrlScheme(this *QWebEngineUrlScheme) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlSchemeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:90
// index:0
// Public Visibility=Default Availability=Available
// [1] bool operator==(const QWebEngineUrlScheme &) const

/*

 */
func (this *QWebEngineUrlScheme) Operator_equal_equal(that QWebEngineUrlScheme_ITF) bool {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineUrlScheme_PTR() != nil {
		convArg0 = that.QWebEngineUrlScheme_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWebEngineUrlSchemeeqERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:91
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool operator!=(const QWebEngineUrlScheme &) const

/*

 */
func (this *QWebEngineUrlScheme) Operator_not_equal(that QWebEngineUrlScheme_ITF) bool {
	var convArg0 unsafe.Pointer
	if that != nil && that.QWebEngineUrlScheme_PTR() != nil {
		convArg0 = that.QWebEngineUrlScheme_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWebEngineUrlSchemeneERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:93
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray name() const

/*

 */
func (this *QWebEngineUrlScheme) Name() *qtcore.QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWebEngineUrlScheme4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := qtcore.NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2, qtcore.DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:94
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setName(const QByteArray &)

/*

 */
func (this *QWebEngineUrlScheme) SetName(newValue qtcore.QByteArray_ITF) {
	var convArg0 unsafe.Pointer
	if newValue != nil && newValue.QByteArray_PTR() != nil {
		convArg0 = newValue.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlScheme7setNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:96
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineUrlScheme::Syntax syntax() const

/*

 */
func (this *QWebEngineUrlScheme) Syntax() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWebEngineUrlScheme6syntaxEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:97
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSyntax(QWebEngineUrlScheme::Syntax)

/*

 */
func (this *QWebEngineUrlScheme) SetSyntax(newValue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlScheme9setSyntaxENS_6SyntaxE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newValue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:99
// index:0
// Public Visibility=Default Availability=Available
// [4] int defaultPort() const

/*

 */
func (this *QWebEngineUrlScheme) DefaultPort() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWebEngineUrlScheme11defaultPortEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:100
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setDefaultPort(int)

/*

 */
func (this *QWebEngineUrlScheme) SetDefaultPort(newValue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlScheme14setDefaultPortEi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newValue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:102
// index:0
// Public Visibility=Default Availability=Available
// [4] QWebEngineUrlScheme::Flags flags() const

/*

 */
func (this *QWebEngineUrlScheme) Flags() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QWebEngineUrlScheme5flagsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:103
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFlags(QWebEngineUrlScheme::Flags)

/*

 */
func (this *QWebEngineUrlScheme) SetFlags(newValue int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlScheme8setFlagsE6QFlagsINS_4FlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), newValue)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:105
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void registerScheme(const QWebEngineUrlScheme &)

/*

 */
func (this *QWebEngineUrlScheme) RegisterScheme(scheme QWebEngineUrlScheme_ITF) {
	var convArg0 unsafe.Pointer
	if scheme != nil && scheme.QWebEngineUrlScheme_PTR() != nil {
		convArg0 = scheme.QWebEngineUrlScheme_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlScheme14registerSchemeERKS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QWebEngineUrlScheme_RegisterScheme(scheme QWebEngineUrlScheme_ITF) {
	var nilthis *QWebEngineUrlScheme
	nilthis.RegisterScheme(scheme)
}

// /usr/include/qt/QtWebEngineCore/qwebengineurlscheme.h:106
// index:0
// Public static Visibility=Default Availability=Available
// [8] QWebEngineUrlScheme schemeByName(const QByteArray &)

/*

 */
func (this *QWebEngineUrlScheme) SchemeByName(name qtcore.QByteArray_ITF) *QWebEngineUrlScheme /*123*/ {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN19QWebEngineUrlScheme12schemeByNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQWebEngineUrlSchemeFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQWebEngineUrlScheme)
	return rv2
}
func QWebEngineUrlScheme_SchemeByName(name qtcore.QByteArray_ITF) *QWebEngineUrlScheme /*123*/ {
	var nilthis *QWebEngineUrlScheme
	rv := nilthis.SchemeByName(name)
	return rv
}

/*


 */
type QWebEngineUrlScheme__Syntax = int

//
const QWebEngineUrlScheme__HostPortAndUserInformation QWebEngineUrlScheme__Syntax = 0

//
const QWebEngineUrlScheme__HostAndPort QWebEngineUrlScheme__Syntax = 1

//
const QWebEngineUrlScheme__Host QWebEngineUrlScheme__Syntax = 2

//
const QWebEngineUrlScheme__Path QWebEngineUrlScheme__Syntax = 3

func (this *QWebEngineUrlScheme) SyntaxItemName(val int) string {
	switch val {
	case QWebEngineUrlScheme__HostPortAndUserInformation: // 0
		return "HostPortAndUserInformation"
	case QWebEngineUrlScheme__HostAndPort: // 1
		return "HostAndPort"
	case QWebEngineUrlScheme__Host: // 2
		return "Host"
	case QWebEngineUrlScheme__Path: // 3
		return "Path"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineUrlScheme_SyntaxItemName(val int) string {
	var nilthis *QWebEngineUrlScheme
	return nilthis.SyntaxItemName(val)
}

/*


 */
type QWebEngineUrlScheme__SpecialPort = int

//
const QWebEngineUrlScheme__PortUnspecified QWebEngineUrlScheme__SpecialPort = -1

func (this *QWebEngineUrlScheme) SpecialPortItemName(val int) string {
	switch val {
	case QWebEngineUrlScheme__PortUnspecified: // -1
		return "PortUnspecified"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineUrlScheme_SpecialPortItemName(val int) string {
	var nilthis *QWebEngineUrlScheme
	return nilthis.SpecialPortItemName(val)
}

/*


 */
type QWebEngineUrlScheme__Flag = int

//
const QWebEngineUrlScheme__SecureScheme QWebEngineUrlScheme__Flag = 1

//
const QWebEngineUrlScheme__LocalScheme QWebEngineUrlScheme__Flag = 2

//
const QWebEngineUrlScheme__LocalAccessAllowed QWebEngineUrlScheme__Flag = 4

//
const QWebEngineUrlScheme__NoAccessAllowed QWebEngineUrlScheme__Flag = 8

//
const QWebEngineUrlScheme__ServiceWorkersAllowed QWebEngineUrlScheme__Flag = 16

//
const QWebEngineUrlScheme__ViewSourceAllowed QWebEngineUrlScheme__Flag = 32

//
const QWebEngineUrlScheme__ContentSecurityPolicyIgnored QWebEngineUrlScheme__Flag = 64

func (this *QWebEngineUrlScheme) FlagItemName(val int) string {
	switch val {
	case QWebEngineUrlScheme__SecureScheme: // 1
		return "SecureScheme"
	case QWebEngineUrlScheme__LocalScheme: // 2
		return "LocalScheme"
	case QWebEngineUrlScheme__LocalAccessAllowed: // 4
		return "LocalAccessAllowed"
	case QWebEngineUrlScheme__NoAccessAllowed: // 8
		return "NoAccessAllowed"
	case QWebEngineUrlScheme__ServiceWorkersAllowed: // 16
		return "ServiceWorkersAllowed"
	case QWebEngineUrlScheme__ViewSourceAllowed: // 32
		return "ViewSourceAllowed"
	case QWebEngineUrlScheme__ContentSecurityPolicyIgnored: // 64
		return "ContentSecurityPolicyIgnored"
	default:
		return fmt.Sprintf("%d", val)
	}
}
func QWebEngineUrlScheme_FlagItemName(val int) string {
	var nilthis *QWebEngineUrlScheme
	return nilthis.FlagItemName(val)
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
	if false {
		qtqml.KeepMe()
	}
	if false {
		qtpositioning.KeepMe()
	}
	if false {
		qtwebchannel.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtquick.KeepMe()
	}
}

//  keep block end
