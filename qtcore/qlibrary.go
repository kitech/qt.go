package qtcore

// /usr/include/qt/QtCore/qlibrary.h
// #include <qlibrary.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 35
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
type QLibrary struct {
	*QObject
}
type QLibrary_ITF interface {
	QObject_ITF
	QLibrary_PTR() *QLibrary
}

func (ptr *QLibrary) QLibrary_PTR() *QLibrary { return ptr }

func (this *QLibrary) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QLibrary) SetCthis(cthis unsafe.Pointer) {
	this.QObject = NewQObjectFromPointer(cthis)
}
func NewQLibraryFromPointer(cthis unsafe.Pointer) *QLibrary {
	bcthis0 := NewQObjectFromPointer(cthis)
	return &QLibrary{bcthis0}
}
func (*QLibrary) NewFromPointer(cthis unsafe.Pointer) *QLibrary {
	return NewQLibraryFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlibrary.h:53
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject() const

/*

 */
func (this *QLibrary) MetaObject() *QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QLibrary10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtCore/qlibrary.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit(parent QObject_ITF /*777 QObject **/) *QLibrary {
	return NewQLibrary(parent)
}
func NewQLibrary(parent QObject_ITF /*777 QObject **/) *QLibrary {
	var convArg0 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg0 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:68
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit__() *QLibrary {
	return NewQLibrary__()
}
func NewQLibrary__() *QLibrary {
	// arg: 0, QObject *=Pointer, QObject=Record, , Invalid
	var convArg0 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2EP7QObject", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit_1(fileName string, parent QObject_ITF /*777 QObject **/) *QLibrary {
	return NewQLibrary_1(fileName, parent)
}
func NewQLibrary_1(fileName string, parent QObject_ITF /*777 QObject **/) *QLibrary {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg1 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:69
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit_1_(fileName string) *QLibrary {
	return NewQLibrary_1_(fileName)
}
func NewQLibrary_1_(fileName string) *QLibrary {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 1, QObject *=Pointer, QObject=Record, , Invalid
	var convArg1 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:70
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, int, QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit_2(fileName string, verNum int, parent QObject_ITF /*777 QObject **/) *QLibrary {
	return NewQLibrary_2(fileName, verNum, parent)
}
func NewQLibrary_2(fileName string, verNum int, parent QObject_ITF /*777 QObject **/) *QLibrary {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringiP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, verNum, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:70
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, int, QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit_2_(fileName string, verNum int) *QLibrary {
	return NewQLibrary_2_(fileName, verNum)
}
func NewQLibrary_2_(fileName string, verNum int) *QLibrary {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringiP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, verNum, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:71
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, const QString &, QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit_3(fileName string, version string, parent QObject_ITF /*777 QObject **/) *QLibrary {
	return NewQLibrary_3(fileName, version, parent)
}
func NewQLibrary_3(fileName string, version string, parent QObject_ITF /*777 QObject **/) *QLibrary {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(version)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 unsafe.Pointer
	if parent != nil && parent.QObject_PTR() != nil {
		convArg2 = parent.QObject_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringS2_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:71
// index:3
// Public Visibility=Default Availability=Available
// [-2] void QLibrary(const QString &, const QString &, QObject *)

/*
Constructs a library with the given parent.
*/
func (*QLibrary) NewForInherit_3_(fileName string, version string) *QLibrary {
	return NewQLibrary_3_(fileName, version)
}
func NewQLibrary_3_(fileName string, version string) *QLibrary {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(version)
	var convArg1 = tmpArg1.GetCthis()
	// arg: 2, QObject *=Pointer, QObject=Record, , Invalid
	var convArg2 unsafe.Pointer
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryC2ERK7QStringS2_P7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQLibraryFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QLibrary")
	return gothis
}

// /usr/include/qt/QtCore/qlibrary.h:72
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QLibrary()

/*

 */
func DeleteQLibrary(this *QLibrary) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibraryD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 32)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qlibrary.h:74
// index:0
// Public Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const char *)

/*
Returns the address of the exported symbol symbol. The library is loaded if necessary. The function returns 0 if the symbol could not be resolved or if the library could not be loaded.

Example:


  typedef int (*AvgFunction)(int, int);

  AvgFunction avg = (AvgFunction) library->resolve("avg");
  if (avg)
      return avg(5, 8);
  else
      return -1;



The symbol must be exported as a C function from the library. This means that the function must be wrapped in an extern "C" if the library is compiled with a C++ compiler. On Windows you must also explicitly export the function from the DLL using the __declspec(dllexport) compiler directive, for example:


  extern "C" MY_EXPORT int avg(int a, int b)
  {
      return (a + b) / 2;
  }



with MY_EXPORT defined as


  #ifdef Q_OS_WIN
  #define MY_EXPORT __declspec(dllexport)
  #else
  #define MY_EXPORT
  #endif
*/
func (this *QLibrary) Resolve(symbol string) unsafe.Pointer /*666*/ {
	var convArg0 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary7resolveEPKc", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtCore/qlibrary.h:75
// index:1
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const QString &, const char *)

/*
Returns the address of the exported symbol symbol. The library is loaded if necessary. The function returns 0 if the symbol could not be resolved or if the library could not be loaded.

Example:


  typedef int (*AvgFunction)(int, int);

  AvgFunction avg = (AvgFunction) library->resolve("avg");
  if (avg)
      return avg(5, 8);
  else
      return -1;



The symbol must be exported as a C function from the library. This means that the function must be wrapped in an extern "C" if the library is compiled with a C++ compiler. On Windows you must also explicitly export the function from the DLL using the __declspec(dllexport) compiler directive, for example:


  extern "C" MY_EXPORT int avg(int a, int b)
  {
      return (a + b) / 2;
  }



with MY_EXPORT defined as


  #ifdef Q_OS_WIN
  #define MY_EXPORT __declspec(dllexport)
  #else
  #define MY_EXPORT
  #endif
*/
func (this *QLibrary) Resolve_1(fileName string, symbol string) unsafe.Pointer /*666*/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg1 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg1)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringPKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QLibrary_Resolve_1(fileName string, symbol string) unsafe.Pointer /*666*/ {
	var nilthis *QLibrary
	rv := nilthis.Resolve_1(fileName, symbol)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:76
// index:2
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const QString &, int, const char *)

/*
Returns the address of the exported symbol symbol. The library is loaded if necessary. The function returns 0 if the symbol could not be resolved or if the library could not be loaded.

Example:


  typedef int (*AvgFunction)(int, int);

  AvgFunction avg = (AvgFunction) library->resolve("avg");
  if (avg)
      return avg(5, 8);
  else
      return -1;



The symbol must be exported as a C function from the library. This means that the function must be wrapped in an extern "C" if the library is compiled with a C++ compiler. On Windows you must also explicitly export the function from the DLL using the __declspec(dllexport) compiler directive, for example:


  extern "C" MY_EXPORT int avg(int a, int b)
  {
      return (a + b) / 2;
  }



with MY_EXPORT defined as


  #ifdef Q_OS_WIN
  #define MY_EXPORT __declspec(dllexport)
  #else
  #define MY_EXPORT
  #endif
*/
func (this *QLibrary) Resolve_2(fileName string, verNum int, symbol string) unsafe.Pointer /*666*/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var convArg2 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringiPKc", qtrt.FFI_TYPE_POINTER, convArg0, verNum, convArg2)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QLibrary_Resolve_2(fileName string, verNum int, symbol string) unsafe.Pointer /*666*/ {
	var nilthis *QLibrary
	rv := nilthis.Resolve_2(fileName, verNum, symbol)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:77
// index:3
// Public static Visibility=Default Availability=Available
// [8] QFunctionPointer resolve(const QString &, const QString &, const char *)

/*
Returns the address of the exported symbol symbol. The library is loaded if necessary. The function returns 0 if the symbol could not be resolved or if the library could not be loaded.

Example:


  typedef int (*AvgFunction)(int, int);

  AvgFunction avg = (AvgFunction) library->resolve("avg");
  if (avg)
      return avg(5, 8);
  else
      return -1;



The symbol must be exported as a C function from the library. This means that the function must be wrapped in an extern "C" if the library is compiled with a C++ compiler. On Windows you must also explicitly export the function from the DLL using the __declspec(dllexport) compiler directive, for example:


  extern "C" MY_EXPORT int avg(int a, int b)
  {
      return (a + b) / 2;
  }



with MY_EXPORT defined as


  #ifdef Q_OS_WIN
  #define MY_EXPORT __declspec(dllexport)
  #else
  #define MY_EXPORT
  #endif
*/
func (this *QLibrary) Resolve_3(fileName string, version string, symbol string) unsafe.Pointer /*666*/ {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(version)
	var convArg1 = tmpArg1.GetCthis()
	var convArg2 = qtrt.CString(symbol)
	defer qtrt.FreeMem(convArg2)
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary7resolveERK7QStringS2_PKc", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}
func QLibrary_Resolve_3(fileName string, version string, symbol string) unsafe.Pointer /*666*/ {
	var nilthis *QLibrary
	rv := nilthis.Resolve_3(fileName, version, symbol)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:79
// index:0
// Public Visibility=Default Availability=Available
// [1] bool load()

/*
Loads the library and returns true if the library was loaded successfully; otherwise returns false. Since resolve() always calls this function before resolving any symbols it is not necessary to call it explicitly. In some situations you might want the library loaded in advance, in which case you would use this function.

See also unload().
*/
func (this *QLibrary) Load() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary4loadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlibrary.h:80
// index:0
// Public Visibility=Default Availability=Available
// [1] bool unload()

/*
Unloads the library and returns true if the library could be unloaded; otherwise returns false.

This happens automatically on application termination, so you shouldn't normally need to call this function.

If other instances of QLibrary are using the same library, the call will fail, and unloading will only happen when every instance has called unload().

Note that on Mac OS X 10.3 (Panther), dynamic libraries cannot be unloaded.

See also resolve() and load().
*/
func (this *QLibrary) Unload() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary6unloadEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlibrary.h:81
// index:0
// Public Visibility=Default Availability=Available
// [1] bool isLoaded() const

/*
Returns true if the library is loaded; otherwise returns false.

See also load().
*/
func (this *QLibrary) IsLoaded() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QLibrary8isLoadedEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
}

// /usr/include/qt/QtCore/qlibrary.h:83
// index:0
// Public static Visibility=Default Availability=Available
// [1] bool isLibrary(const QString &)

/*
Returns true if fileName has a valid suffix for a loadable library; otherwise returns false.


 PlatformValid suffixes
Windows.dll, .DLL
Unix/Linux.so
AIX.a
HP-UX.sl, .so (HP-UXi)
macOS and iOS.dylib, .bundle, .so


Trailing versioning numbers on Unix are ignored.
*/
func (this *QLibrary) IsLibrary(fileName string) bool {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary9isLibraryERK7QString", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return rv != 0
}
func QLibrary_IsLibrary(fileName string) bool {
	var nilthis *QLibrary
	rv := nilthis.IsLibrary(fileName)
	return rv
}

// /usr/include/qt/QtCore/qlibrary.h:85
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileName(const QString &)

/*

 */
func (this *QLibrary) SetFileName(fileName string) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary11setFileNameERK7QString", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:86
// index:0
// Public Visibility=Default Availability=Available
// [8] QString fileName() const

/*

 */
func (this *QLibrary) FileName() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QLibrary8fileNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlibrary.h:88
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setFileNameAndVersion(const QString &, int)

/*
Sets the fileName property and major version number to fileName and versionNumber respectively. The versionNumber is ignored on Windows.

See also setFileName().
*/
func (this *QLibrary) SetFileNameAndVersion(fileName string, verNum int) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringi", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, verNum)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:89
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setFileNameAndVersion(const QString &, const QString &)

/*
Sets the fileName property and major version number to fileName and versionNumber respectively. The versionNumber is ignored on Windows.

See also setFileName().
*/
func (this *QLibrary) SetFileNameAndVersion_1(fileName string, version string) {
	var tmpArg0 = NewQString_5(fileName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(version)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary21setFileNameAndVersionERK7QStringS2_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:90
// index:0
// Public Visibility=Default Availability=Available
// [8] QString errorString() const

/*
Returns a text string with the description of the last error that occurred. Currently, errorString will only be set if load(), unload() or resolve() for some reason fails.

This function was introduced in  Qt 4.2.
*/
func (this *QLibrary) ErrorString() string {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QLibrary11errorStringEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringFromPointer(unsafe.Pointer(uintptr(rv)))
	rv3 := rv2.ToLocal8Bit().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qlibrary.h:92
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLoadHints(QLibrary::LoadHints)

/*

 */
func (this *QLibrary) SetLoadHints(hints int) {
	rv, err := qtrt.InvokeQtFunc6("_ZN8QLibrary12setLoadHintsE6QFlagsINS_8LoadHintEE", qtrt.FFI_TYPE_POINTER, this.GetCthis(), hints)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtCore/qlibrary.h:93
// index:0
// Public Visibility=Default Availability=Available
// [4] QLibrary::LoadHints loadHints() const

/*

 */
func (this *QLibrary) LoadHints() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK8QLibrary9loadHintsEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return int(rv)
}

/*


 */
type QLibrary__LoadHint = int

//
const QLibrary__ResolveAllSymbolsHint QLibrary__LoadHint = 1

//
const QLibrary__ExportExternalSymbolsHint QLibrary__LoadHint = 2

//
const QLibrary__LoadArchiveMemberHint QLibrary__LoadHint = 4

//
const QLibrary__PreventUnloadHint QLibrary__LoadHint = 8

//
const QLibrary__DeepBindHint QLibrary__LoadHint = 16

func (this *QLibrary) LoadHintItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QLibrary_LoadHintItemName(val int) string {
	var nilthis *QLibrary
	return nilthis.LoadHintItemName(val)
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
