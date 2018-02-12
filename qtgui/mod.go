package qtgui

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
// import "C"
import (
	"unsafe"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

func init() {
	if false {
		qtcore.KeepMe()
	}
}

func KeepMe() {}

/////
type QWindowListx struct {
	*qtrt.CObject
}

func (this *QWindowListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWindowListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWindowListxFromPointer(cthis unsafe.Pointer) *QWindowListx {
	return &QWindowListx{&qtrt.CObject{cthis}}
}
func NewQWindowListFromPointer(cthis unsafe.Pointer) *QWindowList {
	return &QWindowList{&qtrt.CObject{cthis}}
}
func (*QWindowListx) NewFromPointer(cthis unsafe.Pointer) *QWindowListx {
	return NewQWindowListxFromPointer(cthis)
}

func (this *QWindowListx) At(i int) *QWindow {
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	return NewQWindowFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QWindowListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QWindowList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QWindowListx) ConvertToSlice() (lst []*QWindow) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QWindowListx) ConvertFromSlice(lst []*QWindow) *QWindowListx {
	return nil
}
