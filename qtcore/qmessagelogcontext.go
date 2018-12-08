package qtcore

// /usr/include/qt/QtCore/qlogging.h
// #include <qlogging.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 11
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
type QMessageLogContext struct {
	*qtrt.CObject
}
type QMessageLogContext_ITF interface {
	QMessageLogContext_PTR() *QMessageLogContext
}

func (ptr *QMessageLogContext) QMessageLogContext_PTR() *QMessageLogContext { return ptr }

func (this *QMessageLogContext) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMessageLogContext) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMessageLogContextFromPointer(cthis unsafe.Pointer) *QMessageLogContext {
	return &QMessageLogContext{&qtrt.CObject{cthis}}
}
func (*QMessageLogContext) NewFromPointer(cthis unsafe.Pointer) *QMessageLogContext {
	return NewQMessageLogContextFromPointer(cthis)
}

// /usr/include/qt/QtCore/qlogging.h:66
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void QMessageLogContext()

/*

 */
func (*QMessageLogContext) NewForInherit() *QMessageLogContext {
	return NewQMessageLogContext()
}
func NewQMessageLogContext() *QMessageLogContext {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMessageLogContextC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageLogContextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageLogContext)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:68
// index:1
// Public inline Visibility=Default Availability=Available
// [-2] void QMessageLogContext(const char *, int, const char *, const char *)

/*

 */
func (*QMessageLogContext) NewForInherit1(fileName string, lineNumber int, functionName string, categoryName string) *QMessageLogContext {
	return NewQMessageLogContext1(fileName, lineNumber, functionName, categoryName)
}
func NewQMessageLogContext1(fileName string, lineNumber int, functionName string, categoryName string) *QMessageLogContext {
	var convArg0 = qtrt.CString(fileName)
	defer qtrt.FreeMem(convArg0)
	var convArg2 = qtrt.CString(functionName)
	defer qtrt.FreeMem(convArg2)
	var convArg3 = qtrt.CString(categoryName)
	defer qtrt.FreeMem(convArg3)
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMessageLogContextC2EPKciS1_S1_", qtrt.FFI_TYPE_POINTER, convArg0, lineNumber, convArg2, convArg3)
	qtrt.ErrPrint(err, rv)
	gothis := NewQMessageLogContextFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQMessageLogContext)
	return gothis
}

// /usr/include/qt/QtCore/qlogging.h:71
// index:0
// Public Visibility=Default Availability=Available
// [-2] void copy(const QMessageLogContext &)

/*

 */
func (this *QMessageLogContext) Copy(logContext QMessageLogContext_ITF) {
	var convArg0 unsafe.Pointer
	if logContext != nil && logContext.QMessageLogContext_PTR() != nil {
		convArg0 = logContext.QMessageLogContext_PTR().GetCthis()
	}
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMessageLogContext4copyERKS_", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

func DeleteQMessageLogContext(this *QMessageLogContext) {
	rv, err := qtrt.InvokeQtFunc6("_ZN18QMessageLogContextD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
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
