package qtcore

import "github.com/kitech/qt.go/qtrt"

func (this *QByteArray) Data_fix() string {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QByteArray4dataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringIN(rv, this.Size())
}

func QVersion() string {
	rv, err := qtrt.InvokeQtFunc6("C_qVersion", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	return qtrt.GoStringI(rv)
}
