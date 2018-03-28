package qtcore

// /usr/include/qt/QtCore/qabstractitemmodel.h
// #include <qabstractitemmodel.h>
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
import "log"
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
		log.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
}

//  keep block end

//  body block begin
type QModelIndexList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QModelIndexList) Operator_equal_0() *QModelIndexList {
	// QModelIndexList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QModelIndexList) Operator_equal_1() *QModelIndexList {
	// QModelIndexList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QModelIndexList) Swap_0() {
	// QModelIndexList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QModelIndexList) Operator_equal_equal_0() bool {
	// QModelIndexList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QModelIndexList) Operator_not_equal_0() bool {
	// QModelIndexList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QModelIndexList) Size_0() int {
	// QModelIndexList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QModelIndexList) Detach_0() {
	// QModelIndexList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QModelIndexList) DetachShared_0() {
	// QModelIndexList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QModelIndexList) IsDetached_0() bool {
	// QModelIndexList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QModelIndexList) SetSharable_0() {
	// QModelIndexList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QModelIndexList) IsSharedWith_0() bool {
	// QModelIndexList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QModelIndexList) IsEmpty_0() bool {
	// QModelIndexList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QModelIndexList) Clear_0() {
	// QModelIndexList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QModelIndexList) At_0() *QModelIndex {
	// QModelIndexList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & operator[](int)
func (this *QModelIndexList) Operator_get_index_0() *QModelIndex {
	// QModelIndexList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// T & operator[](int)
func (this *QModelIndexList) Operator_get_index_1() *QModelIndex {
	// QModelIndexList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// void reserve(int)
func (this *QModelIndexList) Reserve_0() {
	// QModelIndexList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QModelIndexList) Append_0() {
	// QModelIndexList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QModelIndexList) Append_1() {
	// QModelIndexList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QModelIndexList) Prepend_0() {
	// QModelIndexList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QModelIndexList) Insert_0() {
	// QModelIndexList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QModelIndexList) Replace_0() {
	// QModelIndexList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QModelIndexList) RemoveAt_0() {
	// QModelIndexList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QModelIndexList) RemoveAll_0() int {
	// QModelIndexList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QModelIndexList) RemoveOne_0() bool {
	// QModelIndexList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QModelIndexList) TakeAt_0() *QModelIndex {
	// QModelIndexList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// T takeFirst()
func (this *QModelIndexList) TakeFirst_0() *QModelIndex {
	// QModelIndexList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// T takeLast()
func (this *QModelIndexList) TakeLast_0() *QModelIndex {
	// QModelIndexList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// void move(int, int)
func (this *QModelIndexList) Move_0() {
	// QModelIndexList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QModelIndexList) Swap_1() {
	// QModelIndexList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QModelIndexList) IndexOf_0() int {
	// QModelIndexList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QModelIndexList) LastIndexOf_0() int {
	// QModelIndexList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QModelIndexList) Contains_0() bool {
	// QModelIndexList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QModelIndexList) Count_0() int {
	// QModelIndexList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QModelIndexList) Begin_0() {
	// QModelIndexList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QModelIndexList) Begin_1() {
	// QModelIndexList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QModelIndexList) Cbegin_0() {
	// QModelIndexList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QModelIndexList) ConstBegin_0() {
	// QModelIndexList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QModelIndexList) End_0() {
	// QModelIndexList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QModelIndexList) End_1() {
	// QModelIndexList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QModelIndexList) Cend_0() {
	// QModelIndexList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QModelIndexList) ConstEnd_0() {
	// QModelIndexList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QModelIndexList) Rbegin_0() {
	// QModelIndexList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QModelIndexList) Rend_0() {
	// QModelIndexList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QModelIndexList) Rbegin_1() {
	// QModelIndexList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QModelIndexList) Rend_1() {
	// QModelIndexList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QModelIndexList) Crbegin_0() {
	// QModelIndexList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QModelIndexList) Crend_0() {
	// QModelIndexList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QModelIndexList) Insert_1() {
	// QModelIndexList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QModelIndexList) Erase_0() {
	// QModelIndexList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QModelIndexList) Erase_1() {
	// QModelIndexList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QModelIndexList) Count_1() int {
	// QModelIndexList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QModelIndexList) Length_0() int {
	// QModelIndexList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QModelIndexList) First_0() *QModelIndex {
	// QModelIndexList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & constFirst()
func (this *QModelIndexList) ConstFirst_0() *QModelIndex {
	// QModelIndexList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & first()
func (this *QModelIndexList) First_1() *QModelIndex {
	// QModelIndexList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// T & last()
func (this *QModelIndexList) Last_0() *QModelIndex {
	// QModelIndexList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & last()
func (this *QModelIndexList) Last_1() *QModelIndex {
	// QModelIndexList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & constLast()
func (this *QModelIndexList) ConstLast_0() *QModelIndex {
	// QModelIndexList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// void removeFirst()
func (this *QModelIndexList) RemoveFirst_0() {
	// QModelIndexList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QModelIndexList) RemoveLast_0() {
	// QModelIndexList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QModelIndexList) StartsWith_0() bool {
	// QModelIndexList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QModelIndexList) EndsWith_0() bool {
	// QModelIndexList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QModelIndexList) Mid_0() *QModelIndexList {
	// QModelIndexList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QModelIndexList) Value_0() *QModelIndex {
	// QModelIndexList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// T value(int, const T &)
func (this *QModelIndexList) Value_1() *QModelIndex {
	// QModelIndexList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// void push_back(const T &)
func (this *QModelIndexList) Push_back_0() {
	// QModelIndexList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QModelIndexList) Push_front_0() {
	// QModelIndexList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QModelIndexList) Front_0() *QModelIndex {
	// QModelIndexList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & front()
func (this *QModelIndexList) Front_1() *QModelIndex {
	// QModelIndexList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// T & back()
func (this *QModelIndexList) Back_0() *QModelIndex {
	// QModelIndexList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// const T & back()
func (this *QModelIndexList) Back_1() *QModelIndex {
	// QModelIndexList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QModelIndex{}
}

// void pop_front()
func (this *QModelIndexList) Pop_front_0() {
	// QModelIndexList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QModelIndexList) Pop_back_0() {
	// QModelIndexList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QModelIndexList) Empty_0() bool {
	// QModelIndexList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QModelIndexList) Operator_add_equal_0() *QModelIndexList {
	// QModelIndexList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QModelIndexList) Operator_add_0() *QModelIndexList {
	// QModelIndexList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QModelIndexList) Operator_add_equal_1() *QModelIndexList {
	// QModelIndexList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QModelIndexList) Operator_left_shift_0() *QModelIndexList {
	// QModelIndexList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QModelIndexList) Operator_left_shift_1() *QModelIndexList {
	// QModelIndexList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QModelIndexList) ToVector_0() {
	// QModelIndexList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QModelIndexList) ToSet_0() {
	// QModelIndexList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QModelIndexList) FromVector_0() *QModelIndexList {
	// QModelIndexList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QModelIndexList) FromSet_0() *QModelIndexList {
	// QModelIndexList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QModelIndexList) FromStdList_0() *QModelIndexList {
	// QModelIndexList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QModelIndexList) ToStdList_0() {
	// QModelIndexList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QModelIndexList) Detach_helper_grow_0() {
	// QModelIndexList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QModelIndexList) Detach_helper_0() {
	// QModelIndexList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QModelIndexList) Detach_helper_1() {
	// QModelIndexList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QModelIndexList) Dealloc_0() {
	// QModelIndexList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QModelIndexList) Node_construct_0() {
	// QModelIndexList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QModelIndexList) Node_destruct_0() {
	// QModelIndexList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QModelIndexList) Node_copy_0() {
	// QModelIndexList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QModelIndexList) Node_destruct_1() {
	// QModelIndexList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QModelIndexList) IsValidIterator_0() bool {
	// QModelIndexList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QModelIndexList) Op_eq_impl_0() bool {
	// QModelIndexList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QModelIndexList) Op_eq_impl_1() bool {
	// QModelIndexList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QModelIndexList) Contains_impl_0() bool {
	// QModelIndexList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QModelIndexList) Contains_impl_1() bool {
	// QModelIndexList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QModelIndexList) Count_impl_0() int {
	// QModelIndexList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QModelIndexList) Count_impl_1() int {
	// QModelIndexList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QModelIndexList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
