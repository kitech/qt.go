package qtmultimedia

// /usr/include/qt/QtMultimedia/qcamerainfo.h
// #include <qcamerainfo.h>
// #include <QtMultimedia>

//  header block end

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
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtgui"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

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
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end

//  body block begin
type QCameraInfoList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QCameraInfoList) Operator_equal_0() *QCameraInfoList {
	// QCameraInfoList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QCameraInfoList) Operator_equal_1() *QCameraInfoList {
	// QCameraInfoList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QCameraInfoList) Swap_0() {
	// QCameraInfoList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QCameraInfoList) Operator_equal_equal_0() bool {
	// QCameraInfoList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QCameraInfoList) Operator_not_equal_0() bool {
	// QCameraInfoList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QCameraInfoList) Size_0() int {
	// QCameraInfoList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QCameraInfoList) Detach_0() {
	// QCameraInfoList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QCameraInfoList) DetachShared_0() {
	// QCameraInfoList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QCameraInfoList) IsDetached_0() bool {
	// QCameraInfoList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QCameraInfoList) SetSharable_0() {
	// QCameraInfoList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QCameraInfoList) IsSharedWith_0() bool {
	// QCameraInfoList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QCameraInfoList) IsEmpty_0() bool {
	// QCameraInfoList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QCameraInfoList) Clear_0() {
	// QCameraInfoList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QCameraInfoList) At_0() *QCameraInfo {
	// QCameraInfoList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & operator[](int)
func (this *QCameraInfoList) Operator_get_index_0() *QCameraInfo {
	// QCameraInfoList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// T & operator[](int)
func (this *QCameraInfoList) Operator_get_index_1() *QCameraInfo {
	// QCameraInfoList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// void reserve(int)
func (this *QCameraInfoList) Reserve_0() {
	// QCameraInfoList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QCameraInfoList) Append_0() {
	// QCameraInfoList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QCameraInfoList) Append_1() {
	// QCameraInfoList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QCameraInfoList) Prepend_0() {
	// QCameraInfoList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QCameraInfoList) Insert_0() {
	// QCameraInfoList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QCameraInfoList) Replace_0() {
	// QCameraInfoList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QCameraInfoList) RemoveAt_0() {
	// QCameraInfoList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QCameraInfoList) RemoveAll_0() int {
	// QCameraInfoList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QCameraInfoList) RemoveOne_0() bool {
	// QCameraInfoList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QCameraInfoList) TakeAt_0() *QCameraInfo {
	// QCameraInfoList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// T takeFirst()
func (this *QCameraInfoList) TakeFirst_0() *QCameraInfo {
	// QCameraInfoList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// T takeLast()
func (this *QCameraInfoList) TakeLast_0() *QCameraInfo {
	// QCameraInfoList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// void move(int, int)
func (this *QCameraInfoList) Move_0() {
	// QCameraInfoList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QCameraInfoList) Swap_1() {
	// QCameraInfoList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QCameraInfoList) IndexOf_0() int {
	// QCameraInfoList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QCameraInfoList) LastIndexOf_0() int {
	// QCameraInfoList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QCameraInfoList) Contains_0() bool {
	// QCameraInfoList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QCameraInfoList) Count_0() int {
	// QCameraInfoList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QCameraInfoList) Begin_0() {
	// QCameraInfoList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QCameraInfoList) Begin_1() {
	// QCameraInfoList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QCameraInfoList) Cbegin_0() {
	// QCameraInfoList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QCameraInfoList) ConstBegin_0() {
	// QCameraInfoList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QCameraInfoList) End_0() {
	// QCameraInfoList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QCameraInfoList) End_1() {
	// QCameraInfoList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QCameraInfoList) Cend_0() {
	// QCameraInfoList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QCameraInfoList) ConstEnd_0() {
	// QCameraInfoList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QCameraInfoList) Rbegin_0() {
	// QCameraInfoList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QCameraInfoList) Rend_0() {
	// QCameraInfoList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QCameraInfoList) Rbegin_1() {
	// QCameraInfoList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QCameraInfoList) Rend_1() {
	// QCameraInfoList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QCameraInfoList) Crbegin_0() {
	// QCameraInfoList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QCameraInfoList) Crend_0() {
	// QCameraInfoList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QCameraInfoList) Insert_1() {
	// QCameraInfoList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QCameraInfoList) Erase_0() {
	// QCameraInfoList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QCameraInfoList) Erase_1() {
	// QCameraInfoList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QCameraInfoList) Count_1() int {
	// QCameraInfoList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QCameraInfoList) Length_0() int {
	// QCameraInfoList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QCameraInfoList) First_0() *QCameraInfo {
	// QCameraInfoList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & constFirst()
func (this *QCameraInfoList) ConstFirst_0() *QCameraInfo {
	// QCameraInfoList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & first()
func (this *QCameraInfoList) First_1() *QCameraInfo {
	// QCameraInfoList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// T & last()
func (this *QCameraInfoList) Last_0() *QCameraInfo {
	// QCameraInfoList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & last()
func (this *QCameraInfoList) Last_1() *QCameraInfo {
	// QCameraInfoList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & constLast()
func (this *QCameraInfoList) ConstLast_0() *QCameraInfo {
	// QCameraInfoList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// void removeFirst()
func (this *QCameraInfoList) RemoveFirst_0() {
	// QCameraInfoList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QCameraInfoList) RemoveLast_0() {
	// QCameraInfoList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QCameraInfoList) StartsWith_0() bool {
	// QCameraInfoList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QCameraInfoList) EndsWith_0() bool {
	// QCameraInfoList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QCameraInfoList) Mid_0() *QCameraInfoList {
	// QCameraInfoList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QCameraInfoList) Value_0() *QCameraInfo {
	// QCameraInfoList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// T value(int, const T &)
func (this *QCameraInfoList) Value_1() *QCameraInfo {
	// QCameraInfoList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// void push_back(const T &)
func (this *QCameraInfoList) Push_back_0() {
	// QCameraInfoList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QCameraInfoList) Push_front_0() {
	// QCameraInfoList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QCameraInfoList) Front_0() *QCameraInfo {
	// QCameraInfoList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & front()
func (this *QCameraInfoList) Front_1() *QCameraInfo {
	// QCameraInfoList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// T & back()
func (this *QCameraInfoList) Back_0() *QCameraInfo {
	// QCameraInfoList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// const T & back()
func (this *QCameraInfoList) Back_1() *QCameraInfo {
	// QCameraInfoList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QCameraInfo{}
}

// void pop_front()
func (this *QCameraInfoList) Pop_front_0() {
	// QCameraInfoList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QCameraInfoList) Pop_back_0() {
	// QCameraInfoList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QCameraInfoList) Empty_0() bool {
	// QCameraInfoList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QCameraInfoList) Operator_add_equal_0() *QCameraInfoList {
	// QCameraInfoList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QCameraInfoList) Operator_add_0() *QCameraInfoList {
	// QCameraInfoList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QCameraInfoList) Operator_add_equal_1() *QCameraInfoList {
	// QCameraInfoList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QCameraInfoList) Operator_left_shift_0() *QCameraInfoList {
	// QCameraInfoList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QCameraInfoList) Operator_left_shift_1() *QCameraInfoList {
	// QCameraInfoList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QCameraInfoList) ToVector_0() {
	// QCameraInfoList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QCameraInfoList) ToSet_0() {
	// QCameraInfoList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QCameraInfoList) FromVector_0() *QCameraInfoList {
	// QCameraInfoList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QCameraInfoList) FromSet_0() *QCameraInfoList {
	// QCameraInfoList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QCameraInfoList) FromStdList_0() *QCameraInfoList {
	// QCameraInfoList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QCameraInfoList) ToStdList_0() {
	// QCameraInfoList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QCameraInfoList) Detach_helper_grow_0() {
	// QCameraInfoList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QCameraInfoList) Detach_helper_0() {
	// QCameraInfoList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QCameraInfoList) Detach_helper_1() {
	// QCameraInfoList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QCameraInfoList) Dealloc_0() {
	// QCameraInfoList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QCameraInfoList) Node_construct_0() {
	// QCameraInfoList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QCameraInfoList) Node_destruct_0() {
	// QCameraInfoList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QCameraInfoList) Node_copy_0() {
	// QCameraInfoList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QCameraInfoList) Node_destruct_1() {
	// QCameraInfoList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QCameraInfoList) IsValidIterator_0() bool {
	// QCameraInfoList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QCameraInfoList) Op_eq_impl_0() bool {
	// QCameraInfoList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QCameraInfoList) Op_eq_impl_1() bool {
	// QCameraInfoList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QCameraInfoList) Contains_impl_0() bool {
	// QCameraInfoList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QCameraInfoList) Contains_impl_1() bool {
	// QCameraInfoList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QCameraInfoList) Count_impl_0() int {
	// QCameraInfoList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QCameraInfoList) Count_impl_1() int {
	// QCameraInfoList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QCameraInfoList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
