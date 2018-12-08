package qtcore

// /usr/include/qt/QtCore/qtranslator.h
// #include <qtranslator.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 33
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
type QTranslator struct {
	*QObject
}
type QTranslator_ITF interface {
	QObject_ITF
	QTranslator_PTR() *QTranslator
}

func (ptr *QTranslator) QTranslator_PTR() *QTranslator { return ptr }

func (this *QTranslator) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QTranslator) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQTranslatorFromPointer(cthis unsafe.Pointer) *QTranslator {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QTranslator{bcthis0}
}
func (*QTranslator) NewFromPointer(cthis unsafe.Pointer) *QTranslator {
	return NewQTranslatorFromPointer(cthis)
}

// /usr/include/qt/QtCore/qtranslator.h:56
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QTranslator) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTranslator10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qtranslator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTranslator(QObject *)

/*
Constructs an empty message file object with parent parent that is not connected to any file.
*/
func (*QTranslator) NewForInherit(parent QObject_ITF /*777 QObject **/) *QTranslator {
	return NewQTranslator(parent)
}
func NewQTranslator(parent QObject_ITF /*777 QObject **/) *QTranslator {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTranslatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTranslator")
	return gothis
}

// /usr/include/qt/QtCore/qtranslator.h:58
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTranslator(QObject *)

/*
Constructs an empty message file object with parent parent that is not connected to any file.
*/
func (*QTranslator) NewForInheritp() *QTranslator {
	return NewQTranslatorp()
}
func NewQTranslatorp() *QTranslator {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslatorC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTranslatorFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QTranslator")
	return gothis
}

// /usr/include/qt/QtCore/qtranslator.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTranslator()

/*

 */
func DeleteQTranslator(this *QTranslator) {
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslatorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qtranslator.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int) const

/*
Returns the translation for the key (context, sourceText, disambiguation). If none is found, also tries (context, sourceText, ""). If that still fails, returns a null string.

Note: Incomplete translations may result in unexpected behavior: If no translation for (context, sourceText, "") is provided, the method might in this case actually return a translation for a different disambiguation.

If n is not -1, it is used to choose an appropriate form for the translation (e.g. "%n file found" vs. "%n files found").

If you need to programatically insert translations into a QTranslator, this function can be reimplemented.

See also load().
*/
func (this *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sourceText)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTranslator9translateEPKcS1_S1_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtranslator.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int) const

/*
Returns the translation for the key (context, sourceText, disambiguation). If none is found, also tries (context, sourceText, ""). If that still fails, returns a null string.

Note: Incomplete translations may result in unexpected behavior: If no translation for (context, sourceText, "") is provided, the method might in this case actually return a translation for a different disambiguation.

If n is not -1, it is used to choose an appropriate form for the translation (e.g. "%n file found" vs. "%n files found").

If you need to programatically insert translations into a QTranslator, this function can be reimplemented.

See also load().
*/
func (this *QTranslator) Translatep(context string, sourceText string) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sourceText)
	defer qtrt.FreeMem(convArg1)
	// arg: 2, const char *=Pointer, =Invalid, , Invalid
	var convArg2 unsafe.Pointer
	// arg: 3, int=Int, =Invalid, , Invalid
	n := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTranslator9translateEPKcS1_S1_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtranslator.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int) const

/*
Returns the translation for the key (context, sourceText, disambiguation). If none is found, also tries (context, sourceText, ""). If that still fails, returns a null string.

Note: Incomplete translations may result in unexpected behavior: If no translation for (context, sourceText, "") is provided, the method might in this case actually return a translation for a different disambiguation.

If n is not -1, it is used to choose an appropriate form for the translation (e.g. "%n file found" vs. "%n files found").

If you need to programatically insert translations into a QTranslator, this function can be reimplemented.

See also load().
*/
func (this *QTranslator) Translatep1(context string, sourceText string, disambiguation string) string {
	var convArg0 = qtrt.CString(context)
	defer qtrt.FreeMem(convArg0)
	var convArg1 = qtrt.CString(sourceText)
	defer qtrt.FreeMem(convArg1)
	var convArg2 = qtrt.CString(disambiguation)
	defer qtrt.FreeMem(convArg2)
	// arg: 3, int=Int, =Invalid, , Invalid
	n := int(-1)
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTranslator9translateEPKcS1_S1_i", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, n)
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qtranslator.h:64
// index:0
// Public virtual Visibility=Default Availability=Available
// [1] bool isEmpty() const

/*
Returns true if this translator is empty, otherwise returns false. This function works with stripped and unstripped translation files.
*/
func (this *QTranslator) IsEmpty() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK11QTranslator7isEmptyEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load(filename string, directory string, search_delimiters string, suffix string) bool {
	var tmpArg0 = NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(directory)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString5(search_delimiters)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString5(suffix)
	var convArg3 = tmpArg3.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Loadp(filename string) bool {
	var tmpArg0 = NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, const QString &=LValueReference, QString=Record, , Invalid
	var convArg1 = NewQString()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Loadp1(filename string, directory string) bool {
	var tmpArg0 = NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(directory)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:66
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load(const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Loadp2(filename string, directory string, search_delimiters string) bool {
	var tmpArg0 = NewQString5(filename)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString5(directory)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString5(search_delimiters)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QStringS2_S2_S2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QLocale &, const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load1(locale QLocale_ITF, filename string, prefix string, directory string, suffix string) bool {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(filename)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString5(prefix)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString5(directory)
	var convArg3 = tmpArg3.GetCthis()
	var tmpArg4 = NewQString5(suffix)
	var convArg4 = tmpArg4.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QLocale &, const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load1p(locale QLocale_ITF, filename string) bool {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(filename)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = NewQString()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QLocale &, const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load1p1(locale QLocale_ITF, filename string, prefix string) bool {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(filename)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString5(prefix)
	var convArg2 = tmpArg2.GetCthis()
	// arg: 3, const QString &=LValueReference, QString=Record, , Invalid
	var convArg3 = NewQString()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:70
// index:1
// Public Visibility=Default Availability=Available
// [1] bool load(const QLocale &, const QString &, const QString &, const QString &, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load1p2(locale QLocale_ITF, filename string, prefix string, directory string) bool {
	var convArg0 unsafe.Pointer
	if locale != nil && locale.QLocale_PTR() != nil {
		convArg0 = locale.QLocale_PTR().GetCthis()
	}
	var tmpArg1 = NewQString5(filename)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString5(prefix)
	var convArg2 = tmpArg2.GetCthis()
	var tmpArg3 = NewQString5(directory)
	var convArg3 = tmpArg3.GetCthis()
	// arg: 4, const QString &=LValueReference, QString=Record, , Invalid
	var convArg4 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadERK7QLocaleRK7QStringS5_S5_S5_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1, convArg2, convArg3, convArg4)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:75
// index:2
// Public Visibility=Default Availability=Available
// [1] bool load(const uchar *, int, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load2(data unsafe.Pointer /*666*/, len_ int, directory string) bool {
	var tmpArg2 = NewQString5(directory)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadEPKhiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), data, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qtranslator.h:75
// index:2
// Public Visibility=Default Availability=Available
// [1] bool load(const uchar *, int, const QString &)

/*
Loads filename + suffix (".qm" if the suffix is not specified), which may be an absolute file name or relative to directory. Returns true if the translation is successfully loaded; otherwise returns false.

If directory is not specified, the current directory is used (i.e., as currentPath()).

The previous contents of this translator object are discarded.

If the file name does not exist, other file names are tried in the following order:
*/
func (this *QTranslator) Load2p(data unsafe.Pointer /*666*/, len_ int) bool {
	// arg: 2, const QString &=LValueReference, QString=Record, , Invalid
	var convArg2 = NewQString()
	rv, err := qtrt.InvokeQtFunc6("_ZN11QTranslator4loadEPKhiRK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), data, len_, convArg2)
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
