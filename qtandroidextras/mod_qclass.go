package qtandroidextras

import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

// this file is write by hand, don't delete

func (this *QAndroidJniObject) CallObjectMethod(methodName, signature string, args ...interface{}) *QAndroidJniObject {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject16callObjectMethodEPKcS1_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return NewQAndroidJniObjectFromPointer(unsafe.Pointer(uintptr(rv)))
}

//
func (this *QAndroidJniObject) CallMethod(methodName, signature string, args ...interface{}) {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIvEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err, rv)
}

func (this *QAndroidJniObject) CallMethod_2(methodName, signature string, args ...interface{}) byte {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIhEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return byte(rv)
}

func (this *QAndroidJniObject) CallMethod_3(methodName, signature string, args ...interface{}) int8 {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIaEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return int8(rv)
}

func (this *QAndroidJniObject) CallMethod_4(methodName, signature string, args ...interface{}) uint8 {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodItEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return uint8(rv)
}

func (this *QAndroidJniObject) CallMethod_5(methodName, signature string, args ...interface{}) int16 {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIsEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return int16(rv)
}

func (this *QAndroidJniObject) CallMethod_6(methodName, signature string, args ...interface{}) int {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIiEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return int(rv)
}

func (this *QAndroidJniObject) CallMethod_7(methodName, signature string, args ...interface{}) int64 {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIxEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return int64(rv)
}

func (this *QAndroidJniObject) CallMethod_8(methodName, signature string, args ...interface{}) float32 {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIfEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return float32(rv)
}

func (this *QAndroidJniObject) CallMethod_9(methodName, signature string, args ...interface{}) float64 {
	rv, err := qtrt.InvokeQtFunc6Var("C_ZNK17QAndroidJniObject10callMethodIdEET_PKcS3_z",
		qtrt.FFI_TYPE_POINTER, 3, append([]interface{}{this.GetCthis(), methodName, signature}, args...)...)
	qtrt.ErrPrint(err)
	return float64(rv)
}

/////
type JNIEnv struct {
	cthis unsafe.Pointer
}

func (this *QAndroidJniEnvironment) JNIEnv() *JNIEnv {
	rv, err := qtrt.InvokeQtFunc6("C_ZNK22QAndroidJniEnvironmentcvP7_JNIEnvEv",
		qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err)
	return &JNIEnv{unsafe.Pointer(uintptr(rv))}
}

func (this *JNIEnv) ExceptionCheck() bool {
	rv, err := qtrt.InvokeQtFunc6("C_JNIEnv_ExceptionCheck", qtrt.FFI_TYPE_POINTER, unsafe.Pointer(this))
	qtrt.ErrPrint(err, rv)
	return qtrt.IfElse(rv == 0, false, true).(bool)
}

func (this *JNIEnv) ExceptionClear() {
	rv, err := qtrt.InvokeQtFunc6("C_JNIEnv_ExceptionClear", qtrt.FFI_TYPE_POINTER, unsafe.Pointer(this))
	qtrt.ErrPrint(err, rv)
}
