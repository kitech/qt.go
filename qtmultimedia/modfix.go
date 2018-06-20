package qtmultimedia

import (
	"unsafe"

	"github.com/kitech/qt.go/qtcore"
	"github.com/kitech/qt.go/qtrt"
)

type QSizeList = qtcore.QSizeList

/////
type QCameraFocusZoneListx struct {
	*qtrt.CObject
}

func (this *QCameraFocusZoneListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCameraFocusZoneListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCameraFocusZoneListxFromPointer(cthis unsafe.Pointer) *QCameraFocusZoneListx {
	return &QCameraFocusZoneListx{&qtrt.CObject{cthis}}
}
func NewQCameraFocusZoneListFromPointer(cthis unsafe.Pointer) *QCameraFocusZoneList {
	return &QCameraFocusZoneList{&qtrt.CObject{cthis}}
}
func (*QCameraFocusZoneListx) NewFromPointer(cthis unsafe.Pointer) *QCameraFocusZoneListx {
	return NewQCameraFocusZoneListxFromPointer(cthis)
}
func DeleteQCameraFocusZoneListx(this *QCameraFocusZoneListx) {
	qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_dtor", qtrt.FFI_TYPE_POINTER, this.GetCthis())
}
func (this *QCameraFocusZoneListx) Release() { this.CObject = nil }

func (this *QCameraFocusZoneListx) At(i int) *QCameraFocusZone {
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	urlo := NewQCameraFocusZoneFromPointer(unsafe.Pointer(uintptr(rv)))
	return urlo
}

func (this *QCameraFocusZoneListx) Count(s string) int {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QCameraFocusZoneListx) Count_1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QCameraFocusZoneListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QCameraFocusZoneList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QCameraFocusZoneListx) ConvertToSlice() (lst []*QCameraFocusZone) {
	for i := 0; i < this.Size(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QCameraFocusZoneListx) ConvertFromSlice(lst []*QCameraFocusZone) *QCameraFocusZoneListx {
	return nil
}

/////
type QMediaResourceListx struct {
	*qtrt.CObject
}

func (this *QMediaResourceListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QMediaResourceListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQMediaResourceListxFromPointer(cthis unsafe.Pointer) *QMediaResourceListx {
	return &QMediaResourceListx{&qtrt.CObject{cthis}}
}
func NewQMediaResourceListFromPointer(cthis unsafe.Pointer) *QMediaResourceList {
	return &QMediaResourceList{&qtrt.CObject{cthis}}
}
func (*QMediaResourceListx) NewFromPointer(cthis unsafe.Pointer) *QMediaResourceListx {
	return NewQMediaResourceListxFromPointer(cthis)
}
func DeleteQMediaResourceListx(this *QMediaResourceListx) {
	qtrt.InvokeQtFunc6("C_QMediaResourceList_dtor", qtrt.FFI_TYPE_POINTER, this.GetCthis())
}
func (this *QMediaResourceListx) Release() { this.CObject = nil }

func (this *QMediaResourceListx) At(i int) *QMediaResource {
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	urlo := NewQMediaResourceFromPointer(unsafe.Pointer(uintptr(rv)))
	return urlo
}

func (this *QMediaResourceListx) Count(s string) int {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QMediaResourceListx) Count_1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QMediaResourceListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QMediaResourceList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QMediaResourceListx) ConvertToSlice() (lst []*QMediaResource) {
	for i := 0; i < this.Size(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QMediaResourceListx) ConvertFromSlice(lst []*QMediaResource) *QMediaResourceListx {
	return nil
}

/////
type QCameraInfoListx struct {
	*qtrt.CObject
}

func (this *QCameraInfoListx) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QCameraInfoListx) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQCameraInfoListxFromPointer(cthis unsafe.Pointer) *QCameraInfoListx {
	return &QCameraInfoListx{&qtrt.CObject{cthis}}
}
func NewQCameraInfoListFromPointer(cthis unsafe.Pointer) *QCameraInfoList {
	return &QCameraInfoList{&qtrt.CObject{cthis}}
}
func (*QCameraInfoListx) NewFromPointer(cthis unsafe.Pointer) *QCameraInfoListx {
	return NewQCameraInfoListxFromPointer(cthis)
}
func DeleteQCameraInfoListx(this *QCameraInfoListx) {
	qtrt.InvokeQtFunc6("C_QCameraInfoList_dtor", qtrt.FFI_TYPE_POINTER, this.GetCthis())
}
func (this *QCameraInfoListx) Release() { this.CObject = nil }

func (this *QCameraInfoListx) At(i int) *QCameraInfo {
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_at_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), i)
	qtrt.ErrPrint(err, rv)
	urlo := NewQCameraInfoFromPointer(unsafe.Pointer(uintptr(rv)))
	return urlo
}

func (this *QCameraInfoListx) Count(s string) int {
	var convArg0 = qtrt.CString(s)
	defer qtrt.FreeMem(convArg0)
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_count_0", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QCameraInfoListx) Count_1() int {
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_count_1", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QCameraInfoListx) Size() int {
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_size_0", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int)
}

func (this *QCameraInfoListx) ConvertToSlice() (lst []*QCameraInfo) {
	for i := 0; i < this.Size(); i++ {
		lst = append(lst, this.At(i))
	}
	return
}

func (*QCameraInfoListx) ConvertFromSlice(lst []*QCameraInfo) *QCameraInfoListx {
	return nil
}
