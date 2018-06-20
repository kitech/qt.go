package qtcore

// /usr/include/qt/QtCore/qabstractstate.h
// #include <qabstractstate.h>
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
type QAbstractStateList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QAbstractStateList) Operator_equal_0() *QAbstractStateList {
	// QAbstractStateList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QAbstractStateList) Operator_equal_1() *QAbstractStateList {
	// QAbstractStateList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QAbstractStateList) Swap_0() {
	// QAbstractStateList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QAbstractStateList) Operator_equal_equal_0() bool {
	// QAbstractStateList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QAbstractStateList) Operator_not_equal_0() bool {
	// QAbstractStateList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QAbstractStateList) Size_0() int {
	// QAbstractStateList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QAbstractStateList) Detach_0() {
	// QAbstractStateList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QAbstractStateList) DetachShared_0() {
	// QAbstractStateList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QAbstractStateList) IsDetached_0() bool {
	// QAbstractStateList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QAbstractStateList) SetSharable_0() {
	// QAbstractStateList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QAbstractStateList) IsSharedWith_0() bool {
	// QAbstractStateList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QAbstractStateList) IsEmpty_0() bool {
	// QAbstractStateList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QAbstractStateList) Clear_0() {
	// QAbstractStateList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QAbstractStateList) At_0() *QAbstractState {
	// QAbstractStateList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & operator[](int)
func (this *QAbstractStateList) Operator_get_index_0() *QAbstractState {
	// QAbstractStateList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// T & operator[](int)
func (this *QAbstractStateList) Operator_get_index_1() *QAbstractState {
	// QAbstractStateList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// void reserve(int)
func (this *QAbstractStateList) Reserve_0() {
	// QAbstractStateList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QAbstractStateList) Append_0() {
	// QAbstractStateList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QAbstractStateList) Append_1() {
	// QAbstractStateList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QAbstractStateList) Prepend_0() {
	// QAbstractStateList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QAbstractStateList) Insert_0() {
	// QAbstractStateList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QAbstractStateList) Replace_0() {
	// QAbstractStateList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QAbstractStateList) RemoveAt_0() {
	// QAbstractStateList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QAbstractStateList) RemoveAll_0() int {
	// QAbstractStateList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QAbstractStateList) RemoveOne_0() bool {
	// QAbstractStateList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QAbstractStateList) TakeAt_0() *QAbstractState {
	// QAbstractStateList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// T takeFirst()
func (this *QAbstractStateList) TakeFirst_0() *QAbstractState {
	// QAbstractStateList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// T takeLast()
func (this *QAbstractStateList) TakeLast_0() *QAbstractState {
	// QAbstractStateList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// void move(int, int)
func (this *QAbstractStateList) Move_0() {
	// QAbstractStateList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QAbstractStateList) Swap_1() {
	// QAbstractStateList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QAbstractStateList) IndexOf_0() int {
	// QAbstractStateList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QAbstractStateList) LastIndexOf_0() int {
	// QAbstractStateList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QAbstractStateList) Contains_0() bool {
	// QAbstractStateList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QAbstractStateList) Count_0() int {
	// QAbstractStateList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QAbstractStateList) Begin_0() {
	// QAbstractStateList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QAbstractStateList) Begin_1() {
	// QAbstractStateList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QAbstractStateList) Cbegin_0() {
	// QAbstractStateList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QAbstractStateList) ConstBegin_0() {
	// QAbstractStateList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QAbstractStateList) End_0() {
	// QAbstractStateList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QAbstractStateList) End_1() {
	// QAbstractStateList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QAbstractStateList) Cend_0() {
	// QAbstractStateList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QAbstractStateList) ConstEnd_0() {
	// QAbstractStateList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QAbstractStateList) Rbegin_0() {
	// QAbstractStateList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QAbstractStateList) Rend_0() {
	// QAbstractStateList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QAbstractStateList) Rbegin_1() {
	// QAbstractStateList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QAbstractStateList) Rend_1() {
	// QAbstractStateList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QAbstractStateList) Crbegin_0() {
	// QAbstractStateList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QAbstractStateList) Crend_0() {
	// QAbstractStateList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QAbstractStateList) Insert_1() {
	// QAbstractStateList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QAbstractStateList) Erase_0() {
	// QAbstractStateList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QAbstractStateList) Erase_1() {
	// QAbstractStateList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QAbstractStateList) Count_1() int {
	// QAbstractStateList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QAbstractStateList) Length_0() int {
	// QAbstractStateList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QAbstractStateList) First_0() *QAbstractState {
	// QAbstractStateList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & constFirst()
func (this *QAbstractStateList) ConstFirst_0() *QAbstractState {
	// QAbstractStateList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & first()
func (this *QAbstractStateList) First_1() *QAbstractState {
	// QAbstractStateList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// T & last()
func (this *QAbstractStateList) Last_0() *QAbstractState {
	// QAbstractStateList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & last()
func (this *QAbstractStateList) Last_1() *QAbstractState {
	// QAbstractStateList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & constLast()
func (this *QAbstractStateList) ConstLast_0() *QAbstractState {
	// QAbstractStateList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// void removeFirst()
func (this *QAbstractStateList) RemoveFirst_0() {
	// QAbstractStateList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QAbstractStateList) RemoveLast_0() {
	// QAbstractStateList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QAbstractStateList) StartsWith_0() bool {
	// QAbstractStateList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QAbstractStateList) EndsWith_0() bool {
	// QAbstractStateList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QAbstractStateList) Mid_0() *QAbstractStateList {
	// QAbstractStateList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QAbstractStateList) Value_0() *QAbstractState {
	// QAbstractStateList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// T value(int, const T &)
func (this *QAbstractStateList) Value_1() *QAbstractState {
	// QAbstractStateList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// void push_back(const T &)
func (this *QAbstractStateList) Push_back_0() {
	// QAbstractStateList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QAbstractStateList) Push_front_0() {
	// QAbstractStateList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QAbstractStateList) Front_0() *QAbstractState {
	// QAbstractStateList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & front()
func (this *QAbstractStateList) Front_1() *QAbstractState {
	// QAbstractStateList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// T & back()
func (this *QAbstractStateList) Back_0() *QAbstractState {
	// QAbstractStateList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// const T & back()
func (this *QAbstractStateList) Back_1() *QAbstractState {
	// QAbstractStateList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAbstractState{}
}

// void pop_front()
func (this *QAbstractStateList) Pop_front_0() {
	// QAbstractStateList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QAbstractStateList) Pop_back_0() {
	// QAbstractStateList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QAbstractStateList) Empty_0() bool {
	// QAbstractStateList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QAbstractStateList) Operator_add_equal_0() *QAbstractStateList {
	// QAbstractStateList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QAbstractStateList) Operator_add_0() *QAbstractStateList {
	// QAbstractStateList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QAbstractStateList) Operator_add_equal_1() *QAbstractStateList {
	// QAbstractStateList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QAbstractStateList) Operator_left_shift_0() *QAbstractStateList {
	// QAbstractStateList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QAbstractStateList) Operator_left_shift_1() *QAbstractStateList {
	// QAbstractStateList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QAbstractStateList) ToVector_0() {
	// QAbstractStateList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QAbstractStateList) ToSet_0() {
	// QAbstractStateList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QAbstractStateList) FromVector_0() *QAbstractStateList {
	// QAbstractStateList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QAbstractStateList) FromSet_0() *QAbstractStateList {
	// QAbstractStateList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QAbstractStateList) FromStdList_0() *QAbstractStateList {
	// QAbstractStateList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QAbstractStateList) ToStdList_0() {
	// QAbstractStateList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QAbstractStateList) Detach_helper_grow_0() {
	// QAbstractStateList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QAbstractStateList) Detach_helper_0() {
	// QAbstractStateList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QAbstractStateList) Detach_helper_1() {
	// QAbstractStateList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QAbstractStateList) Dealloc_0() {
	// QAbstractStateList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QAbstractStateList) Node_construct_0() {
	// QAbstractStateList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QAbstractStateList) Node_destruct_0() {
	// QAbstractStateList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QAbstractStateList) Node_copy_0() {
	// QAbstractStateList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QAbstractStateList) Node_destruct_1() {
	// QAbstractStateList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QAbstractStateList) IsValidIterator_0() bool {
	// QAbstractStateList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractStateList) Op_eq_impl_0() bool {
	// QAbstractStateList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QAbstractStateList) Op_eq_impl_1() bool {
	// QAbstractStateList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractStateList) Contains_impl_0() bool {
	// QAbstractStateList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAbstractStateList) Contains_impl_1() bool {
	// QAbstractStateList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAbstractStateList) Count_impl_0() int {
	// QAbstractStateList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAbstractStateList) Count_impl_1() int {
	// QAbstractStateList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QAbstractStateList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
