package qtcore

// /usr/include/qt/QtCore/qobject.h
// #include <qobject.h>
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
type QObjectList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QObjectList) Operator_equal_0() *QObjectList {
	// QObjectList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QObjectList) Operator_equal_1() *QObjectList {
	// QObjectList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QObjectList) Swap_0() {
	// QObjectList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QObjectList) Operator_equal_equal_0() bool {
	// QObjectList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QObjectList) Operator_not_equal_0() bool {
	// QObjectList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QObjectList) Size_0() int {
	// QObjectList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QObjectList) Detach_0() {
	// QObjectList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QObjectList) DetachShared_0() {
	// QObjectList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QObjectList) IsDetached_0() bool {
	// QObjectList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QObjectList) SetSharable_0() {
	// QObjectList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QObjectList) IsSharedWith_0() bool {
	// QObjectList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QObjectList) IsEmpty_0() bool {
	// QObjectList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QObjectList) Clear_0() {
	// QObjectList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QObjectList) At_0() *QObject {
	// QObjectList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & operator[](int)
func (this *QObjectList) Operator_get_index_0() *QObject {
	// QObjectList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// T & operator[](int)
func (this *QObjectList) Operator_get_index_1() *QObject {
	// QObjectList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// void reserve(int)
func (this *QObjectList) Reserve_0() {
	// QObjectList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QObjectList) Append_0() {
	// QObjectList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QObjectList) Append_1() {
	// QObjectList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QObjectList) Prepend_0() {
	// QObjectList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QObjectList) Insert_0() {
	// QObjectList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QObjectList) Replace_0() {
	// QObjectList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QObjectList) RemoveAt_0() {
	// QObjectList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QObjectList) RemoveAll_0() int {
	// QObjectList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QObjectList) RemoveOne_0() bool {
	// QObjectList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QObjectList) TakeAt_0() *QObject {
	// QObjectList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// T takeFirst()
func (this *QObjectList) TakeFirst_0() *QObject {
	// QObjectList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// T takeLast()
func (this *QObjectList) TakeLast_0() *QObject {
	// QObjectList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// void move(int, int)
func (this *QObjectList) Move_0() {
	// QObjectList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QObjectList) Swap_1() {
	// QObjectList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QObjectList) IndexOf_0() int {
	// QObjectList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QObjectList) LastIndexOf_0() int {
	// QObjectList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QObjectList) Contains_0() bool {
	// QObjectList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QObjectList) Count_0() int {
	// QObjectList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QObjectList) Begin_0() {
	// QObjectList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QObjectList) Begin_1() {
	// QObjectList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QObjectList) Cbegin_0() {
	// QObjectList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QObjectList) ConstBegin_0() {
	// QObjectList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QObjectList) End_0() {
	// QObjectList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QObjectList) End_1() {
	// QObjectList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QObjectList) Cend_0() {
	// QObjectList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QObjectList) ConstEnd_0() {
	// QObjectList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QObjectList) Rbegin_0() {
	// QObjectList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QObjectList) Rend_0() {
	// QObjectList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QObjectList) Rbegin_1() {
	// QObjectList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QObjectList) Rend_1() {
	// QObjectList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QObjectList) Crbegin_0() {
	// QObjectList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QObjectList) Crend_0() {
	// QObjectList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(class QList::iterator, const T &)
func (this *QObjectList) Insert_1() {
	// QObjectList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator)
func (this *QObjectList) Erase_0() {
	// QObjectList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator, class QList::iterator)
func (this *QObjectList) Erase_1() {
	// QObjectList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QObjectList) Count_1() int {
	// QObjectList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QObjectList) Length_0() int {
	// QObjectList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QObjectList) First_0() *QObject {
	// QObjectList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & constFirst()
func (this *QObjectList) ConstFirst_0() *QObject {
	// QObjectList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & first()
func (this *QObjectList) First_1() *QObject {
	// QObjectList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// T & last()
func (this *QObjectList) Last_0() *QObject {
	// QObjectList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & last()
func (this *QObjectList) Last_1() *QObject {
	// QObjectList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & constLast()
func (this *QObjectList) ConstLast_0() *QObject {
	// QObjectList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// void removeFirst()
func (this *QObjectList) RemoveFirst_0() {
	// QObjectList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QObjectList) RemoveLast_0() {
	// QObjectList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QObjectList) StartsWith_0() bool {
	// QObjectList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QObjectList) EndsWith_0() bool {
	// QObjectList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QObjectList) Mid_0() *QObjectList {
	// QObjectList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QObjectList) Value_0() *QObject {
	// QObjectList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// T value(int, const T &)
func (this *QObjectList) Value_1() *QObject {
	// QObjectList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// void push_back(const T &)
func (this *QObjectList) Push_back_0() {
	// QObjectList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QObjectList) Push_front_0() {
	// QObjectList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QObjectList) Front_0() *QObject {
	// QObjectList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & front()
func (this *QObjectList) Front_1() *QObject {
	// QObjectList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// T & back()
func (this *QObjectList) Back_0() *QObject {
	// QObjectList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// const T & back()
func (this *QObjectList) Back_1() *QObject {
	// QObjectList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QObject{}
}

// void pop_front()
func (this *QObjectList) Pop_front_0() {
	// QObjectList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QObjectList) Pop_back_0() {
	// QObjectList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QObjectList) Empty_0() bool {
	// QObjectList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QObjectList) Operator_add_equal_0() *QObjectList {
	// QObjectList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QObjectList) Operator_add_0() *QObjectList {
	// QObjectList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QObjectList) Operator_add_equal_1() *QObjectList {
	// QObjectList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QObjectList) Operator_left_shift_0() *QObjectList {
	// QObjectList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QObjectList) Operator_left_shift_1() *QObjectList {
	// QObjectList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QObjectList) ToVector_0() {
	// QObjectList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QObjectList) ToSet_0() {
	// QObjectList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QObjectList) FromVector_0() *QObjectList {
	// QObjectList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QObjectList) FromSet_0() *QObjectList {
	// QObjectList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QObjectList) FromStdList_0() *QObjectList {
	// QObjectList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QObjectList) ToStdList_0() {
	// QObjectList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QObjectList) Detach_helper_grow_0() {
	// QObjectList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QObjectList) Detach_helper_0() {
	// QObjectList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QObjectList) Detach_helper_1() {
	// QObjectList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(struct QListData::Data *)
func (this *QObjectList) Dealloc_0() {
	// QObjectList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(struct QList::Node *, const T &)
func (this *QObjectList) Node_construct_0() {
	// QObjectList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *)
func (this *QObjectList) Node_destruct_0() {
	// QObjectList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(struct QList::Node *, struct QList::Node *, struct QList::Node *)
func (this *QObjectList) Node_copy_0() {
	// QObjectList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *, struct QList::Node *)
func (this *QObjectList) Node_destruct_1() {
	// QObjectList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const class QList::iterator &)
func (this *QObjectList) IsValidIterator_0() bool {
	// QObjectList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::NotArrayCompatibleLayout)
func (this *QObjectList) Op_eq_impl_0() bool {
	// QObjectList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::ArrayCompatibleLayout)
func (this *QObjectList) Op_eq_impl_1() bool {
	// QObjectList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QObjectList) Contains_impl_0() bool {
	// QObjectList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QObjectList) Contains_impl_1() bool {
	// QObjectList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QObjectList) Count_impl_0() int {
	// QObjectList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QObjectList) Count_impl_1() int {
	// QObjectList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QObjectList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
