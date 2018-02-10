package qtandroidextras

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidservice.h
// #include <qandroidservice.h>
// #include <QtAndroidExtras>

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "gopp"
import "qt.go/qtrt"
import "qt.go/qtcore"

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
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin

type QAndroidService struct {
	*qtcore.QCoreApplication
}

func (this *QAndroidService) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QCoreApplication.GetCthis()
	}
}
func (this *QAndroidService) SetCthis(cthis unsafe.Pointer) {
	this.QCoreApplication = qtcore.NewQCoreApplicationFromPointer(cthis)
}
func NewQAndroidServiceFromPointer(cthis unsafe.Pointer) *QAndroidService {
	bcthis0 := qtcore.NewQCoreApplicationFromPointer(cthis)
	return &QAndroidService{bcthis0}
}
func (*QAndroidService) NewFromPointer(cthis unsafe.Pointer) *QAndroidService {
	return NewQAndroidServiceFromPointer(cthis)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidservice.h:54
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidService(int &, char **, int)
func NewQAndroidService(argc int, argv []string, flags int) *QAndroidService {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidServiceC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, flags)
	gopp.ErrPrint(err, rv)
	gothis := NewQAndroidServiceFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidservice.h:59
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidService()
func DeleteQAndroidService(this *QAndroidService) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidServiceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 1)
	gopp.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/../../src/androidextras/android/qandroidservice.h:61
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAndroidBinder * onBind(const QAndroidIntent &)
func (this *QAndroidService) OnBind(intent *QAndroidIntent) *QAndroidBinder /*777 QAndroidBinder **/ {
	var convArg0 = intent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidService6onBindERK14QAndroidIntent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	gopp.ErrPrint(err, rv)
	return /*==*/ NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end
