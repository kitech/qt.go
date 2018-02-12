package qtcore

// /usr/include/qt/QtCore/qvariant.h
// #include <qvariant.h>
// #include <QtCore>

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
import "github.com/kitech/qt.go/qtrt"

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
		qtrt.KeepMe()
	}
}

//  keep block end

//  body block begin
type QVariantList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QVariantList) Operator_equal_0() *QVariantList {
	// QVariantList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QVariantList) Operator_equal_1() *QVariantList {
	// QVariantList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QVariantList) Swap_0() {
	// QVariantList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QVariantList) Operator_equal_equal_0() bool {
	// QVariantList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QVariantList) Operator_not_equal_0() bool {
	// QVariantList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QVariantList) Size_0() int {
	// QVariantList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QVariantList) Detach_0() {
	// QVariantList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QVariantList) DetachShared_0() {
	// QVariantList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QVariantList) IsDetached_0() bool {
	// QVariantList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QVariantList) SetSharable_0() {
	// QVariantList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QVariantList) IsSharedWith_0() bool {
	// QVariantList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QVariantList) IsEmpty_0() bool {
	// QVariantList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QVariantList) Clear_0() {
	// QVariantList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QVariantList) At_0() *QVariant {
	// QVariantList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & operator[](int)
func (this *QVariantList) Operator_get_index_0() *QVariant {
	// QVariantList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & operator[](int)
func (this *QVariantList) Operator_get_index_1() *QVariant {
	// QVariantList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void reserve(int)
func (this *QVariantList) Reserve_0() {
	// QVariantList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QVariantList) Append_0() {
	// QVariantList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QVariantList) Append_1() {
	// QVariantList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QVariantList) Prepend_0() {
	// QVariantList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QVariantList) Insert_0() {
	// QVariantList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QVariantList) Replace_0() {
	// QVariantList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QVariantList) RemoveAt_0() {
	// QVariantList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QVariantList) RemoveAll_0() int {
	// QVariantList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QVariantList) RemoveOne_0() bool {
	// QVariantList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QVariantList) TakeAt_0() *QVariant {
	// QVariantList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T takeFirst()
func (this *QVariantList) TakeFirst_0() *QVariant {
	// QVariantList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T takeLast()
func (this *QVariantList) TakeLast_0() *QVariant {
	// QVariantList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void move(int, int)
func (this *QVariantList) Move_0() {
	// QVariantList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QVariantList) Swap_1() {
	// QVariantList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QVariantList) IndexOf_0() int {
	// QVariantList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QVariantList) LastIndexOf_0() int {
	// QVariantList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QVariantList) Contains_0() bool {
	// QVariantList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QVariantList) Count_0() int {
	// QVariantList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QVariantList) Begin_0() {
	// QVariantList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QVariantList) Begin_1() {
	// QVariantList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QVariantList) Cbegin_0() {
	// QVariantList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QVariantList) ConstBegin_0() {
	// QVariantList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QVariantList) End_0() {
	// QVariantList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QVariantList) End_1() {
	// QVariantList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QVariantList) Cend_0() {
	// QVariantList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QVariantList) ConstEnd_0() {
	// QVariantList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QVariantList) Rbegin_0() {
	// QVariantList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QVariantList) Rend_0() {
	// QVariantList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QVariantList) Rbegin_1() {
	// QVariantList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QVariantList) Rend_1() {
	// QVariantList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QVariantList) Crbegin_0() {
	// QVariantList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QVariantList) Crend_0() {
	// QVariantList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(class QList::iterator, const T &)
func (this *QVariantList) Insert_1() {
	// QVariantList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator)
func (this *QVariantList) Erase_0() {
	// QVariantList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator, class QList::iterator)
func (this *QVariantList) Erase_1() {
	// QVariantList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QVariantList) Count_1() int {
	// QVariantList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QVariantList) Length_0() int {
	// QVariantList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QVariantList) First_0() *QVariant {
	// QVariantList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & constFirst()
func (this *QVariantList) ConstFirst_0() *QVariant {
	// QVariantList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & first()
func (this *QVariantList) First_1() *QVariant {
	// QVariantList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & last()
func (this *QVariantList) Last_0() *QVariant {
	// QVariantList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & last()
func (this *QVariantList) Last_1() *QVariant {
	// QVariantList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & constLast()
func (this *QVariantList) ConstLast_0() *QVariant {
	// QVariantList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void removeFirst()
func (this *QVariantList) RemoveFirst_0() {
	// QVariantList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QVariantList) RemoveLast_0() {
	// QVariantList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QVariantList) StartsWith_0() bool {
	// QVariantList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QVariantList) EndsWith_0() bool {
	// QVariantList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QVariantList) Mid_0() *QVariantList {
	// QVariantList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QVariantList) Value_0() *QVariant {
	// QVariantList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T value(int, const T &)
func (this *QVariantList) Value_1() *QVariant {
	// QVariantList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void push_back(const T &)
func (this *QVariantList) Push_back_0() {
	// QVariantList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QVariantList) Push_front_0() {
	// QVariantList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QVariantList) Front_0() *QVariant {
	// QVariantList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & front()
func (this *QVariantList) Front_1() *QVariant {
	// QVariantList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & back()
func (this *QVariantList) Back_0() *QVariant {
	// QVariantList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & back()
func (this *QVariantList) Back_1() *QVariant {
	// QVariantList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void pop_front()
func (this *QVariantList) Pop_front_0() {
	// QVariantList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QVariantList) Pop_back_0() {
	// QVariantList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QVariantList) Empty_0() bool {
	// QVariantList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QVariantList) Operator_add_equal_0() *QVariantList {
	// QVariantList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QVariantList) Operator_add_0() *QVariantList {
	// QVariantList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QVariantList) Operator_add_equal_1() *QVariantList {
	// QVariantList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QVariantList) Operator_left_shift_0() *QVariantList {
	// QVariantList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QVariantList) Operator_left_shift_1() *QVariantList {
	// QVariantList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QVariantList) ToVector_0() {
	// QVariantList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QVariantList) ToSet_0() {
	// QVariantList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QVariantList) FromVector_0() *QVariantList {
	// QVariantList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QVariantList) FromSet_0() *QVariantList {
	// QVariantList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QVariantList) FromStdList_0() *QVariantList {
	// QVariantList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QVariantList) ToStdList_0() {
	// QVariantList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QVariantList) Detach_helper_grow_0() {
	// QVariantList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QVariantList) Detach_helper_0() {
	// QVariantList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QVariantList) Detach_helper_1() {
	// QVariantList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(struct QListData::Data *)
func (this *QVariantList) Dealloc_0() {
	// QVariantList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(struct QList::Node *, const T &)
func (this *QVariantList) Node_construct_0() {
	// QVariantList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *)
func (this *QVariantList) Node_destruct_0() {
	// QVariantList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(struct QList::Node *, struct QList::Node *, struct QList::Node *)
func (this *QVariantList) Node_copy_0() {
	// QVariantList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *, struct QList::Node *)
func (this *QVariantList) Node_destruct_1() {
	// QVariantList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const class QList::iterator &)
func (this *QVariantList) IsValidIterator_0() bool {
	// QVariantList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Op_eq_impl_0() bool {
	// QVariantList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::ArrayCompatibleLayout)
func (this *QVariantList) Op_eq_impl_1() bool {
	// QVariantList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Contains_impl_0() bool {
	// QVariantList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QVariantList) Contains_impl_1() bool {
	// QVariantList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Count_impl_0() int {
	// QVariantList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QVariantList) Count_impl_1() int {
	// QVariantList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QVariantList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
