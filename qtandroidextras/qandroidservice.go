package qtandroidextras

// /usr/include/qt/QtAndroidExtras/qandroidservice.h
// #include <qandroidservice.h>
// #include <QtAndroidExtras>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 12
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "log"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

/*

 */
type QAndroidService struct {
	*qtcore.QCoreApplication
}
type QAndroidService_ITF interface {
	qtcore.QCoreApplication_ITF
	QAndroidService_PTR() *QAndroidService
}

func (ptr *QAndroidService) QAndroidService_PTR() *QAndroidService { return ptr }

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

// /usr/include/qt/QtAndroidExtras/qandroidservice.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidService(int &, char **, int)

/*
Creates a new Android service, passing argc and argv as parameters.

See also QCoreApplication.
*/
func (*QAndroidService) NewForInherit(argc int, argv []string, flags int) *QAndroidService {
	return NewQAndroidService(argc, argv, flags)
}
func NewQAndroidService(argc int, argv []string, flags int) *QAndroidService {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidServiceC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidServiceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAndroidService")
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidservice.h:55
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QAndroidService(int &, char **, int)

/*
Creates a new Android service, passing argc and argv as parameters.

See also QCoreApplication.
*/
func (*QAndroidService) NewForInheritp(argc int, argv []string) *QAndroidService {
	return NewQAndroidServicep(argc, argv)
}
func NewQAndroidServicep(argc int, argv []string) *QAndroidService {
	var convArg1 = qtrt.StringSliceToCCharPP(argv)
	// arg: 2, int=Int, =Invalid, , Invalid
	flags := 0
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidServiceC2ERiPPci", qtrt.FFI_TYPE_POINTER, &argc, convArg1, flags)
	qtrt.ErrPrint(err, rv)
	gothis := NewQAndroidServiceFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.ConnectDestroyed(gothis, "QAndroidService")
	return gothis
}

// /usr/include/qt/QtAndroidExtras/qandroidservice.h:66
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QAndroidService()

/*

 */
func DeleteQAndroidService(this *QAndroidService) {
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidServiceD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 24)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtAndroidExtras/qandroidservice.h:68
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] QAndroidBinder * onBind(const QAndroidIntent &)

/*
The user must override this method and to return a binder.

The intent parameter contains all the caller information.

The returned binder is used by the caller to perform IPC calls.

Warning: This method is called from Binder's thread which is different from the thread that this object was created.

See also QAndroidBinder::onTransact and QAndroidBinder::transact.
*/
func (this *QAndroidService) OnBind(intent QAndroidIntent_ITF) *QAndroidBinder /*777 QAndroidBinder **/ {
	var convArg0 unsafe.Pointer
	if intent != nil && intent.QAndroidIntent_PTR() != nil {
		convArg0 = intent.QAndroidIntent_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN15QAndroidService6onBindERK14QAndroidIntent", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQAndroidBinderFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

//  body block end

//  keep block begin

func init_unused_14917() {
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
		qtcore.KeepMe()
	}
}

//  keep block end
