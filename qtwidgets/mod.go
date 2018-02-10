package qtwidgets

/*
#cgo CFLAGS: -std=c11
// #cgo LDFLAGS: -lQt5Inline
//  -lQt5Core -lQt5Gui -lQt5Widgets
*/
// import "C"
import (
	"gopp"
	"unsafe"

	"qt.go/qtcore"
	"qt.go/qtgui"
	"qt.go/qtrt"
)

func init() {
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

func KeepMe() {}

/////
type QWidgetListx struct {
	*qtrt.CObject
}

func (this *QWidgetListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWidgetListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWidgetListxFromPointer(cthis unsafe.Pointer) *QWidgetListx {
	return &QWidgetListx{&qtrt.CObject{cthis}}
}
func NewQWidgetListFromPointer(cthis unsafe.Pointer) *QWidgetList {
	return &QWidgetList{&qtrt.CObject{cthis}}
}
func (*QWidgetListx) NewFromPointer(cthis unsafe.Pointer) *QWidgetListx {
	return NewQWidgetListxFromPointer(cthis)
}

func (this *QWidgetListx) At(i int) *QWidget {
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	return NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QWidgetListx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QWidgetListx) ConvertToSlice() (lst []*QWidget) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QWidgetListx) ConvertFromSlice(lst []*QWidget) *QWidgetListx {
	return nil
}

/////
type QWidgetSetx struct {
	*qtrt.CObject
}

func (this *QWidgetSetx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QWidgetSetx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQWidgetSetxFromPointer(cthis unsafe.Pointer) *QWidgetSetx {
	return &QWidgetSetx{&qtrt.CObject{cthis}}
}
func (*QWidgetSetx) NewFromPointer(cthis unsafe.Pointer) *QWidgetSetx {
	return NewQWidgetSetxFromPointer(cthis)
}

func (this *QWidgetSetx) At(i int) *QWidget {
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	gopp.ErrPrint(err, rv)
	return NewQWidgetFromPointer(unsafe.Pointer(uintptr(rv)))
}

func (this *QWidgetSetx) Count() int {
	rv, err := qtrt.InvokeQtFunc6("C_QWidgetSet_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	gopp.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QWidgetSetx) ConvertToSlice() (lst []*QWidget) {
	for i := 0; i < this.Count(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QWidgetSetx) ConvertFromSlice(lst []*QWidget) *QWidgetSetx {
	return nil
}
