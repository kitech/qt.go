package qtqml

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
import "C"
import (
	"unsafe"

	"github.com/kitech/qt.go/qtrt"
)

func KeepMe() {}

/////
type QJSValueListx struct {
	*qtrt.CObject
}

func (this *QJSValueListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QJSValueListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQJSValueListxFromPointer(cthis unsafe.Pointer) *QJSValueListx {
	return &QJSValueListx{&qtrt.CObject{cthis}}
}
func (*QJSValueListx) NewFromPointer(cthis unsafe.Pointer) *QJSValueListx {
	return NewQJSValueListxFromPointer(cthis)
}

func (this *QJSValueListx) At(i int) *QJSValue {
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return NewQJSValueFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QJSValueListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QJSValueListx) ConvertToSlice() (lst []*QJSValue) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QJSValueListx) ConvertFromSlice(lst []*QJSValue) *QJSValueListx {
	return nil
}
