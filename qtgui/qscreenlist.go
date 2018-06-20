package qtgui

// /usr/include/qt/QtGui/qscreen.h
// #include <qscreen.h>
// #include <QtGui>

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
}

//  keep block end

//  body block begin
type QScreenList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QScreenList) Operator_equal_0() *QScreenList {
	// QScreenList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QScreenList) Operator_equal_1() *QScreenList {
	// QScreenList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QScreenList) Swap_0() {
	// QScreenList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QScreenList) Operator_equal_equal_0() bool {
	// QScreenList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QScreenList) Operator_not_equal_0() bool {
	// QScreenList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QScreenList) Size_0() int {
	// QScreenList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QScreenList) Detach_0() {
	// QScreenList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QScreenList) DetachShared_0() {
	// QScreenList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QScreenList) IsDetached_0() bool {
	// QScreenList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QScreenList) SetSharable_0() {
	// QScreenList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QScreenList) IsSharedWith_0() bool {
	// QScreenList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QScreenList) IsEmpty_0() bool {
	// QScreenList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QScreenList) Clear_0() {
	// QScreenList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QScreenList) At_0() *QScreen {
	// QScreenList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & operator[](int)
func (this *QScreenList) Operator_get_index_0() *QScreen {
	// QScreenList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T & operator[](int)
func (this *QScreenList) Operator_get_index_1() *QScreen {
	// QScreenList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void reserve(int)
func (this *QScreenList) Reserve_0() {
	// QScreenList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QScreenList) Append_0() {
	// QScreenList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QScreenList) Append_1() {
	// QScreenList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QScreenList) Prepend_0() {
	// QScreenList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QScreenList) Insert_0() {
	// QScreenList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QScreenList) Replace_0() {
	// QScreenList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QScreenList) RemoveAt_0() {
	// QScreenList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QScreenList) RemoveAll_0() int {
	// QScreenList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QScreenList) RemoveOne_0() bool {
	// QScreenList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QScreenList) TakeAt_0() *QScreen {
	// QScreenList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T takeFirst()
func (this *QScreenList) TakeFirst_0() *QScreen {
	// QScreenList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T takeLast()
func (this *QScreenList) TakeLast_0() *QScreen {
	// QScreenList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void move(int, int)
func (this *QScreenList) Move_0() {
	// QScreenList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QScreenList) Swap_1() {
	// QScreenList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QScreenList) IndexOf_0() int {
	// QScreenList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QScreenList) LastIndexOf_0() int {
	// QScreenList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QScreenList) Contains_0() bool {
	// QScreenList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QScreenList) Count_0() int {
	// QScreenList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QScreenList) Begin_0() {
	// QScreenList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QScreenList) Begin_1() {
	// QScreenList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QScreenList) Cbegin_0() {
	// QScreenList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QScreenList) ConstBegin_0() {
	// QScreenList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QScreenList) End_0() {
	// QScreenList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QScreenList) End_1() {
	// QScreenList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QScreenList) Cend_0() {
	// QScreenList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QScreenList) ConstEnd_0() {
	// QScreenList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QScreenList) Rbegin_0() {
	// QScreenList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QScreenList) Rend_0() {
	// QScreenList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QScreenList) Rbegin_1() {
	// QScreenList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QScreenList) Rend_1() {
	// QScreenList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QScreenList) Crbegin_0() {
	// QScreenList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QScreenList) Crend_0() {
	// QScreenList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QScreenList) Insert_1() {
	// QScreenList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QScreenList) Erase_0() {
	// QScreenList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QScreenList) Erase_1() {
	// QScreenList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QScreenList) Count_1() int {
	// QScreenList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QScreenList) Length_0() int {
	// QScreenList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QScreenList) First_0() *QScreen {
	// QScreenList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & constFirst()
func (this *QScreenList) ConstFirst_0() *QScreen {
	// QScreenList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & first()
func (this *QScreenList) First_1() *QScreen {
	// QScreenList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T & last()
func (this *QScreenList) Last_0() *QScreen {
	// QScreenList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & last()
func (this *QScreenList) Last_1() *QScreen {
	// QScreenList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & constLast()
func (this *QScreenList) ConstLast_0() *QScreen {
	// QScreenList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void removeFirst()
func (this *QScreenList) RemoveFirst_0() {
	// QScreenList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QScreenList) RemoveLast_0() {
	// QScreenList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QScreenList) StartsWith_0() bool {
	// QScreenList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QScreenList) EndsWith_0() bool {
	// QScreenList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QScreenList) Mid_0() *QScreenList {
	// QScreenList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QScreenList) Value_0() *QScreen {
	// QScreenList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T value(int, const T &)
func (this *QScreenList) Value_1() *QScreen {
	// QScreenList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void push_back(const T &)
func (this *QScreenList) Push_back_0() {
	// QScreenList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QScreenList) Push_front_0() {
	// QScreenList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QScreenList) Front_0() *QScreen {
	// QScreenList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & front()
func (this *QScreenList) Front_1() *QScreen {
	// QScreenList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// T & back()
func (this *QScreenList) Back_0() *QScreen {
	// QScreenList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// const T & back()
func (this *QScreenList) Back_1() *QScreen {
	// QScreenList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QScreen{}
}

// void pop_front()
func (this *QScreenList) Pop_front_0() {
	// QScreenList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QScreenList) Pop_back_0() {
	// QScreenList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QScreenList) Empty_0() bool {
	// QScreenList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QScreenList) Operator_add_equal_0() *QScreenList {
	// QScreenList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QScreenList) Operator_add_0() *QScreenList {
	// QScreenList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QScreenList) Operator_add_equal_1() *QScreenList {
	// QScreenList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QScreenList) Operator_left_shift_0() *QScreenList {
	// QScreenList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QScreenList) Operator_left_shift_1() *QScreenList {
	// QScreenList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QScreenList) ToVector_0() {
	// QScreenList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QScreenList) ToSet_0() {
	// QScreenList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QScreenList) FromVector_0() *QScreenList {
	// QScreenList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QScreenList) FromSet_0() *QScreenList {
	// QScreenList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QScreenList) FromStdList_0() *QScreenList {
	// QScreenList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QScreenList) ToStdList_0() {
	// QScreenList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QScreenList) Detach_helper_grow_0() {
	// QScreenList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QScreenList) Detach_helper_0() {
	// QScreenList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QScreenList) Detach_helper_1() {
	// QScreenList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QScreenList) Dealloc_0() {
	// QScreenList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QScreenList) Node_construct_0() {
	// QScreenList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QScreenList) Node_destruct_0() {
	// QScreenList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QScreenList) Node_copy_0() {
	// QScreenList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QScreenList) Node_destruct_1() {
	// QScreenList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QScreenList) IsValidIterator_0() bool {
	// QScreenList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QScreenList) Op_eq_impl_0() bool {
	// QScreenList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QScreenList) Op_eq_impl_1() bool {
	// QScreenList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QScreenList) Contains_impl_0() bool {
	// QScreenList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QScreenList) Contains_impl_1() bool {
	// QScreenList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QScreenList) Count_impl_0() int {
	// QScreenList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QScreenList) Count_impl_1() int {
	// QScreenList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QScreenList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
