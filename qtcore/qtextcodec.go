package qtcore

// /usr/include/qt/QtCore/qtextcodec.h
// #include <qtextcodec.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 16
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
type QTextCodec struct {
	*qtrt.CObject
}
type QTextCodec_ITF interface {
	QTextCodec_PTR() *QTextCodec
}

func (ptr *QTextCodec) QTextCodec_PTR() *QTextCodec { return ptr }

func (this *QTextCodec) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QTextCodec) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQTextCodecFromPointer(cthis unsafe.Pointer) *QTextCodec {
	return &QTextCodec{&qtrt.CObject{cthis}}
}
func (*QTextCodec) NewFromPointer(cthis unsafe.Pointer) *QTextCodec {
	return NewQTextCodecFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtextcodec.h:61
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForName(const QByteArray &)

/*
Searches all installed QTextCodec objects and returns the one which best matches name; the match is case-insensitive. Returns 0 if no codec matching the name name could be found.

Note: This function is thread-safe.
*/
func (this *QTextCodec) CodecForName(name QByteArray_ITF) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 unsafe.Pointer
	if name != nil && name.QByteArray_PTR() != nil {
		convArg0 = name.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForName(name QByteArray_ITF) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForName(name)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:62
// index:1
// Public static inline Visibility=Default Availability=Available
// [8] QTextCodec * codecForName(const char *)

/*
Searches all installed QTextCodec objects and returns the one which best matches name; the match is case-insensitive. Returns 0 if no codec matching the name name could be found.

Note: This function is thread-safe.
*/
func (this *QTextCodec) CodecForName_1(name string) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 = qtrt.CString(name)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec12codecForNameEPKc", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForName_1(name string) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForName_1(name)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForMib(int)

/*
Returns the QTextCodec which matches the MIBenum mib.

Note: This function is thread-safe.
*/
func (this *QTextCodec) CodecForMib(mib int) *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec11codecForMibEi", qtrt.FFI_TYPE_POINTER, mib)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForMib(mib int) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForMib(mib)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:68
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForLocale()

/*
Returns a pointer to the codec most suitable for this locale.

On Windows, the codec will be based on a system locale. On Unix systems, the codec will might fall back to using the iconv library if no builtin codec for the locale can be found.

Note that in these cases the codec's name will be "System".

Note: This function is thread-safe.

See also setCodecForLocale().
*/
func (this *QTextCodec) CodecForLocale() *QTextCodec /*777 QTextCodec **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec14codecForLocaleEv", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForLocale() *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForLocale()
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:69
// index:0
// Public static Visibility=Default Availability=Available
// [-2] void setCodecForLocale(QTextCodec *)

/*
Set the codec to c; this will be returned by codecForLocale(). If c is a null pointer, the codec is reset to the default.

This might be needed for some applications that want to use their own mechanism for setting the locale.

Warning: This function is not reentrant.

See also codecForLocale().
*/
func (this *QTextCodec) SetCodecForLocale(c QTextCodec_ITF /*777 QTextCodec **/) {
	var convArg0 unsafe.Pointer
	if c != nil && c.QTextCodec_PTR() != nil {
		convArg0 = c.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec17setCodecForLocaleEPS_", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
}
func QTextCodec_SetCodecForLocale(c QTextCodec_ITF /*777 QTextCodec **/) {
	var nilthis *QTextCodec
	nilthis.SetCodecForLocale(c)
}

// /usr/include/qt/QtCore/qtextcodec.h:75
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForHtml(const QByteArray &)

/*
Tries to detect the encoding of the provided snippet of HTML in the given byte array, ba, by checking the BOM (Byte Order Mark) and the content-type meta header and returns a QTextCodec instance that is capable of decoding the html to unicode. If the codec cannot be detected from the content provided, defaultCodec is returned.

This function was introduced in  Qt 4.4.

See also codecForUtfText().
*/
func (this *QTextCodec) CodecForHtml(ba QByteArray_ITF) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForHtml(ba QByteArray_ITF) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForHtml(ba)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:76
// index:1
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForHtml(const QByteArray &, QTextCodec *)

/*
Tries to detect the encoding of the provided snippet of HTML in the given byte array, ba, by checking the BOM (Byte Order Mark) and the content-type meta header and returns a QTextCodec instance that is capable of decoding the html to unicode. If the codec cannot be detected from the content provided, defaultCodec is returned.

This function was introduced in  Qt 4.4.

See also codecForUtfText().
*/
func (this *QTextCodec) CodecForHtml_1(ba QByteArray_ITF, defaultCodec QTextCodec_ITF /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if defaultCodec != nil && defaultCodec.QTextCodec_PTR() != nil {
		convArg1 = defaultCodec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec12codecForHtmlERK10QByteArrayPS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForHtml_1(ba QByteArray_ITF, defaultCodec QTextCodec_ITF /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForHtml_1(ba, defaultCodec)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:78
// index:0
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForUtfText(const QByteArray &)

/*
Tries to detect the encoding of the provided snippet ba by using the BOM (Byte Order Mark) and returns a QTextCodec instance that is capable of decoding the text to unicode. If the codec cannot be detected from the content provided, defaultCodec is returned.

This function was introduced in  Qt 4.6.

See also codecForHtml().
*/
func (this *QTextCodec) CodecForUtfText(ba QByteArray_ITF) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArray", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForUtfText(ba QByteArray_ITF) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForUtfText(ba)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:79
// index:1
// Public static Visibility=Default Availability=Available
// [8] QTextCodec * codecForUtfText(const QByteArray &, QTextCodec *)

/*
Tries to detect the encoding of the provided snippet ba by using the BOM (Byte Order Mark) and returns a QTextCodec instance that is capable of decoding the text to unicode. If the codec cannot be detected from the content provided, defaultCodec is returned.

This function was introduced in  Qt 4.6.

See also codecForHtml().
*/
func (this *QTextCodec) CodecForUtfText_1(ba QByteArray_ITF, defaultCodec QTextCodec_ITF /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var convArg0 unsafe.Pointer
	if ba != nil && ba.QByteArray_PTR() != nil {
		convArg0 = ba.QByteArray_PTR().GetCthis()
	}
	var convArg1 unsafe.Pointer
	if defaultCodec != nil && defaultCodec.QTextCodec_PTR() != nil {
		convArg1 = defaultCodec.QTextCodec_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodec15codecForUtfTextERK10QByteArrayPS_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QTextCodec_CodecForUtfText_1(ba QByteArray_ITF, defaultCodec QTextCodec_ITF /*777 QTextCodec **/) *QTextCodec /*777 QTextCodec **/ {
	var nilthis *QTextCodec
	rv := nilthis.CodecForUtfText_1(ba, defaultCodec)
	return rv
}

// /usr/include/qt/QtCore/qtextcodec.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool canEncode(QChar) const

/*
Returns true if the Unicode character ch can be fully encoded with this codec; otherwise returns false.
*/
func (this *QTextCodec) CanEncode(arg0 QChar_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QChar_PTR() != nil {
		convArg0 = arg0.QChar_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE5QChar", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextcodec.h:83
// index:1
// Public Visibility=Default Availability=Available
// [1] bool canEncode(const QString &) const

/*
Returns true if the Unicode character ch can be fully encoded with this codec; otherwise returns false.
*/
func (this *QTextCodec) CanEncode_1(arg0 string) bool {
	var tmpArg0 = NewQString_5(arg0)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextcodec.h:85
// index:2
// Public Visibility=Default Availability=Available
// [1] bool canEncode(QStringView) const

/*
Returns true if the Unicode character ch can be fully encoded with this codec; otherwise returns false.
*/
func (this *QTextCodec) CanEncode_2(arg0 QStringView_ITF /*123*/) bool {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QStringView_PTR() != nil {
		convArg0 = arg0.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec9canEncodeE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtextcodec.h:87
// index:0
// Public Visibility=Default Availability=Available
// [8] QString toUnicode(const QByteArray &) const

/*
Converts a from the encoding of this codec to Unicode, and returns the result in a QString.
*/
func (this *QTextCodec) ToUnicode(arg0 QByteArray_ITF) string {
	var convArg0 unsafe.Pointer
	if arg0 != nil && arg0.QByteArray_PTR() != nil {
		convArg0 = arg0.QByteArray_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeERK10QByteArray", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextcodec.h:88
// index:1
// Public Visibility=Default Availability=Available
// [8] QString toUnicode(const char *) const

/*
Converts a from the encoding of this codec to Unicode, and returns the result in a QString.
*/
func (this *QTextCodec) ToUnicode_1(chars string) string {
	var convArg0 = qtrt.CString(chars)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec9toUnicodeEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtextcodec.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QByteArray fromUnicode(const QString &) const

/*
Converts str from Unicode to the encoding of this codec, and returns the result in a QByteArray.
*/
func (this *QTextCodec) FromUnicode(uc string) *QByteArray /*123*/ {
	var tmpArg0 = NewQString_5(uc)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:92
// index:1
// Public Visibility=Default Availability=Available
// [8] QByteArray fromUnicode(QStringView) const

/*
Converts str from Unicode to the encoding of this codec, and returns the result in a QByteArray.
*/
func (this *QTextCodec) FromUnicode_1(uc QStringView_ITF /*123*/) *QByteArray /*123*/ {
	var convArg0 unsafe.Pointer
	if uc != nil && uc.QStringView_PTR() != nil {
		convArg0 = uc.QStringView_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec11fromUnicodeE11QStringView", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDecoder * makeDecoder(QTextCodec::ConversionFlags) const

/*
Creates a QTextDecoder with a specified flags to decode chunks of char * data to create chunks of Unicode data.

The caller is responsible for deleting the returned object.

This function was introduced in  Qt 4.7.
*/
func (this *QTextCodec) MakeDecoder(flags int) *QTextDecoder /*777 QTextDecoder **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec11makeDecoderE6QFlagsINS_14ConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDecoderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextcodec.h:119
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextDecoder * makeDecoder(QTextCodec::ConversionFlags) const

/*
Creates a QTextDecoder with a specified flags to decode chunks of char * data to create chunks of Unicode data.

The caller is responsible for deleting the returned object.

This function was introduced in  Qt 4.7.
*/
func (this *QTextCodec) MakeDecoder__() *QTextDecoder /*777 QTextDecoder **/ {
	// arg: 0, QTextCodec::ConversionFlags=Typedef, QTextCodec::ConversionFlags=Typedef, QFlags<QTextCodec::ConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec11makeDecoderE6QFlagsINS_14ConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextDecoderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextcodec.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextEncoder * makeEncoder(QTextCodec::ConversionFlags) const

/*
Creates a QTextEncoder with a specified flags to encode chunks of Unicode data as char * data.

The caller is responsible for deleting the returned object.

This function was introduced in  Qt 4.7.
*/
func (this *QTextCodec) MakeEncoder(flags int) *QTextEncoder /*777 QTextEncoder **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec11makeEncoderE6QFlagsINS_14ConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextEncoderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextcodec.h:120
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextEncoder * makeEncoder(QTextCodec::ConversionFlags) const

/*
Creates a QTextEncoder with a specified flags to encode chunks of Unicode data as char * data.

The caller is responsible for deleting the returned object.

This function was introduced in  Qt 4.7.
*/
func (this *QTextCodec) MakeEncoder__() *QTextEncoder /*777 QTextEncoder **/ {
	// arg: 0, QTextCodec::ConversionFlags=Typedef, QTextCodec::ConversionFlags=Typedef, QFlags<QTextCodec::ConversionFlag>, Unexposed
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec11makeEncoderE6QFlagsINS_14ConversionFlagEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), flags)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextEncoderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtextcodec.h:122
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [8] QByteArray name() const

/*
QTextCodec subclasses must reimplement this function. It returns the name of the encoding supported by the subclass.

If the codec is registered as a character set in the IANA character-sets encoding file this method should return the preferred mime name for the codec if defined, otherwise its name.
*/
func (this *QTextCodec) Name() *QByteArray /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQByteArrayFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQByteArray)
	return rv2
}

// /usr/include/qt/QtCore/qtextcodec.h:124
// index:0
// Public purevirtual virtual Visibility=Default Availability=Available
// [4] int mibEnum() const

/*
Subclasses of QTextCodec must reimplement this function. It returns the MIBenum (see IANA character-sets encoding file for more information). It is important that each QTextCodec subclass returns the correct unique value for this function.
*/
func (this *QTextCodec) MibEnum() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextCodec7mibEnumEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtCore/qtextcodec.h:130
// index:0
// Protected Visibility=Default Availability=Available
// [-2] void QTextCodec()

/*
Constructs a QTextCodec, and gives it the highest precedence. The QTextCodec should always be constructed on the heap (i.e. with new). Qt takes ownership and will delete it when the application terminates.
*/
func NewQTextCodec() *QTextCodec {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodecC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextCodecFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQTextCodec)
	return gothis
}

// /usr/include/qt/QtCore/qtextcodec.h:131
// index:0
// Protected virtual Visibility=Default Availability=Available
// [-2] void ~QTextCodec()

/*

 */
func DeleteQTextCodec(this *QTextCodec) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextCodecD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 8)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QTextCodec__ConversionFlag = int

//
const QTextCodec__DefaultConversion QTextCodec__ConversionFlag = 0

//
const QTextCodec__ConvertInvalidToNull QTextCodec__ConversionFlag = -2147483648

//
const QTextCodec__IgnoreHeader QTextCodec__ConversionFlag = 1

//
const QTextCodec__FreeFunction QTextCodec__ConversionFlag = 2

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
