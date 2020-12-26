package qtcore

// /usr/include/qt/QtCore/qcoreapplication.h
// #include <qcoreapplication.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 0
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
// size 16
type QCoreApplication struct {
	*QObject
}
type QCoreApplication_ITF interface {
	QObject_ITF
	QCoreApplication_PTR() *QCoreApplication
}

func (ptr *QCoreApplication) QCoreApplication_PTR() *QCoreApplication { return ptr }

// ignore GetCthis for 1 base
// ignore SetCthis for 1 base
// ignore GetCthis for 1 base
func QCoreApplicationFromptr(cthis Voidptr) *QCoreApplication {
	bcthis0 := QObjectFromptr(cthis)
	return &QCoreApplication{bcthis0}
}
func (*QCoreApplication) Fromptr(cthis Voidptr) *QCoreApplication {
	return QCoreApplicationFromptr(cthis)
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static Indirect Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int)

/*
 */
func (this *QCoreApplication) Translate(context string, key string, disambiguation string, n int) string {
	var convArg0 = qtrt.CStringRef(&context)
	var convArg1 = qtrt.CStringRef(&key)
	var convArg2 = qtrt.CStringRef(&disambiguation)
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2590451099, "_ZN16QCoreApplication9translateEPKcS1_S1_i", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&convArg2), Voidptr(&n))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}
func QCoreApplication_Translate(context string, key string, disambiguation string, n int) string {
	var nilthis *QCoreApplication
	rv := nilthis.Translate(context, key, disambiguation, n)
	return rv
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static Indirect Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int)

/*
 */
func (this *QCoreApplication) Translatep(context string, key string) string {
	var convArg0 = qtrt.CStringRef(&context)
	var convArg1 = qtrt.CStringRef(&key)
	// arg: 2, const char *=Pointer, =Invalid, , Invalid
	var convArg2 Voidptr
	// arg: 3, int=Int, =Invalid, , Invalid
	n := int(-1)
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2590451099, "_ZN16QCoreApplication9translateEPKcS1_S1_i", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&convArg2), Voidptr(&n))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

// /usr/include/qt/QtCore/qcoreapplication.h:156
// index:0
// Public static Indirect Visibility=Default Availability=Available
// [8] QString translate(const char *, const char *, const char *, int)

/*
 */
func (this *QCoreApplication) Translatep1(context string, key string, disambiguation string) string {
	var convArg0 = qtrt.CStringRef(&context)
	var convArg1 = qtrt.CStringRef(&key)
	var convArg2 = qtrt.CStringRef(&disambiguation)
	// arg: 3, int=Int, =Invalid, , Invalid
	n := int(-1)
	sretobj := qtrt.Malloc(8) // QString
	rv, err := qtrt.Qtcc3(2590451099, "_ZN16QCoreApplication9translateEPKcS1_S1_i", qtrt.FFITO_POINTER,
		qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_POINTER, qtrt.FFITO_INT, Voidptr(&sretobj), Voidptr(&convArg0), Voidptr(&convArg1), Voidptr(&convArg2), Voidptr(&n))
	qtrt.ErrPrint2(err, rv)
	rv = qtrt.VRetype(uintptr(sretobj))
	rv2 := /*==*/ QStringFromptr(Voidptr(uintptr(rv)))
	rv3 := rv2.ToUtf8().Data()
	/*==*/ DeleteQString(rv2)
	return rv3
}

func DeleteQCoreApplication(this *QCoreApplication) {
	rv, err := qtrt.Qtcc1(0, "_ZN16QCoreApplicationD2Ev", qtrt.FFITY_VOID, this.GetCthis())
	qtrt.ErrPrint2(err, rv)
	this.SetCthis(nil)
}

/*


 */
type QCoreApplication__ = int

//
const QCoreApplication__ApplicationFlags QCoreApplication__ = 331522

func (this *QCoreApplication) ItemName(val int) string {
	return qtrt.GetClassEnumItemName(this, val)
}
func QCoreApplication_ItemName(val int) string {
	var nilthis *QCoreApplication
	return nilthis.ItemName(val)
}

//  body block end

//  keep block begin

func init_unused_10037() {
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
