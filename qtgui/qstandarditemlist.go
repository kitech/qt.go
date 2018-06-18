package qtgui

// /usr/include/qt/QtGui/qstandarditemmodel.h
// #include <qstandarditemmodel.h>
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
type QStandardItemList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QStandardItemList) Operator_equal_0() *QStandardItemList {
	// QStandardItemList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QStandardItemList) Operator_equal_1() *QStandardItemList {
	// QStandardItemList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QStandardItemList) Swap_0() {
	// QStandardItemList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QStandardItemList) Operator_equal_equal_0() bool {
	// QStandardItemList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QStandardItemList) Operator_not_equal_0() bool {
	// QStandardItemList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QStandardItemList) Size_0() int {
	// QStandardItemList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QStandardItemList) Detach_0() {
	// QStandardItemList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QStandardItemList) DetachShared_0() {
	// QStandardItemList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QStandardItemList) IsDetached_0() bool {
	// QStandardItemList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QStandardItemList) SetSharable_0() {
	// QStandardItemList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QStandardItemList) IsSharedWith_0() bool {
	// QStandardItemList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QStandardItemList) IsEmpty_0() bool {
	// QStandardItemList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QStandardItemList) Clear_0() {
	// QStandardItemList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QStandardItemList) At_0() *QStandardItem {
	// QStandardItemList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & operator[](int)
func (this *QStandardItemList) Operator_get_index_0() *QStandardItem {
	// QStandardItemList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// T & operator[](int)
func (this *QStandardItemList) Operator_get_index_1() *QStandardItem {
	// QStandardItemList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// void reserve(int)
func (this *QStandardItemList) Reserve_0() {
	// QStandardItemList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QStandardItemList) Append_0() {
	// QStandardItemList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QStandardItemList) Append_1() {
	// QStandardItemList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QStandardItemList) Prepend_0() {
	// QStandardItemList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QStandardItemList) Insert_0() {
	// QStandardItemList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QStandardItemList) Replace_0() {
	// QStandardItemList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QStandardItemList) RemoveAt_0() {
	// QStandardItemList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QStandardItemList) RemoveAll_0() int {
	// QStandardItemList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QStandardItemList) RemoveOne_0() bool {
	// QStandardItemList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QStandardItemList) TakeAt_0() *QStandardItem {
	// QStandardItemList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// T takeFirst()
func (this *QStandardItemList) TakeFirst_0() *QStandardItem {
	// QStandardItemList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// T takeLast()
func (this *QStandardItemList) TakeLast_0() *QStandardItem {
	// QStandardItemList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// void move(int, int)
func (this *QStandardItemList) Move_0() {
	// QStandardItemList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QStandardItemList) Swap_1() {
	// QStandardItemList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QStandardItemList) IndexOf_0() int {
	// QStandardItemList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QStandardItemList) LastIndexOf_0() int {
	// QStandardItemList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QStandardItemList) Contains_0() bool {
	// QStandardItemList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QStandardItemList) Count_0() int {
	// QStandardItemList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QStandardItemList) Begin_0() {
	// QStandardItemList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QStandardItemList) Begin_1() {
	// QStandardItemList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QStandardItemList) Cbegin_0() {
	// QStandardItemList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QStandardItemList) ConstBegin_0() {
	// QStandardItemList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QStandardItemList) End_0() {
	// QStandardItemList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QStandardItemList) End_1() {
	// QStandardItemList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QStandardItemList) Cend_0() {
	// QStandardItemList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QStandardItemList) ConstEnd_0() {
	// QStandardItemList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QStandardItemList) Rbegin_0() {
	// QStandardItemList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QStandardItemList) Rend_0() {
	// QStandardItemList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QStandardItemList) Rbegin_1() {
	// QStandardItemList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QStandardItemList) Rend_1() {
	// QStandardItemList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QStandardItemList) Crbegin_0() {
	// QStandardItemList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QStandardItemList) Crend_0() {
	// QStandardItemList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QStandardItemList) Insert_1() {
	// QStandardItemList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QStandardItemList) Erase_0() {
	// QStandardItemList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QStandardItemList) Erase_1() {
	// QStandardItemList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QStandardItemList) Count_1() int {
	// QStandardItemList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QStandardItemList) Length_0() int {
	// QStandardItemList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QStandardItemList) First_0() *QStandardItem {
	// QStandardItemList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & constFirst()
func (this *QStandardItemList) ConstFirst_0() *QStandardItem {
	// QStandardItemList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & first()
func (this *QStandardItemList) First_1() *QStandardItem {
	// QStandardItemList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// T & last()
func (this *QStandardItemList) Last_0() *QStandardItem {
	// QStandardItemList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & last()
func (this *QStandardItemList) Last_1() *QStandardItem {
	// QStandardItemList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & constLast()
func (this *QStandardItemList) ConstLast_0() *QStandardItem {
	// QStandardItemList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// void removeFirst()
func (this *QStandardItemList) RemoveFirst_0() {
	// QStandardItemList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QStandardItemList) RemoveLast_0() {
	// QStandardItemList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QStandardItemList) StartsWith_0() bool {
	// QStandardItemList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QStandardItemList) EndsWith_0() bool {
	// QStandardItemList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QStandardItemList) Mid_0() *QStandardItemList {
	// QStandardItemList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QStandardItemList) Value_0() *QStandardItem {
	// QStandardItemList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// T value(int, const T &)
func (this *QStandardItemList) Value_1() *QStandardItem {
	// QStandardItemList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// void push_back(const T &)
func (this *QStandardItemList) Push_back_0() {
	// QStandardItemList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QStandardItemList) Push_front_0() {
	// QStandardItemList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QStandardItemList) Front_0() *QStandardItem {
	// QStandardItemList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & front()
func (this *QStandardItemList) Front_1() *QStandardItem {
	// QStandardItemList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// T & back()
func (this *QStandardItemList) Back_0() *QStandardItem {
	// QStandardItemList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// const T & back()
func (this *QStandardItemList) Back_1() *QStandardItem {
	// QStandardItemList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QStandardItem{}
}

// void pop_front()
func (this *QStandardItemList) Pop_front_0() {
	// QStandardItemList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QStandardItemList) Pop_back_0() {
	// QStandardItemList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QStandardItemList) Empty_0() bool {
	// QStandardItemList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QStandardItemList) Operator_add_equal_0() *QStandardItemList {
	// QStandardItemList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QStandardItemList) Operator_add_0() *QStandardItemList {
	// QStandardItemList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QStandardItemList) Operator_add_equal_1() *QStandardItemList {
	// QStandardItemList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QStandardItemList) Operator_left_shift_0() *QStandardItemList {
	// QStandardItemList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QStandardItemList) Operator_left_shift_1() *QStandardItemList {
	// QStandardItemList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QStandardItemList) ToVector_0() {
	// QStandardItemList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QStandardItemList) ToSet_0() {
	// QStandardItemList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QStandardItemList) FromVector_0() *QStandardItemList {
	// QStandardItemList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QStandardItemList) FromSet_0() *QStandardItemList {
	// QStandardItemList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QStandardItemList) FromStdList_0() *QStandardItemList {
	// QStandardItemList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QStandardItemList) ToStdList_0() {
	// QStandardItemList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QStandardItemList) Detach_helper_grow_0() {
	// QStandardItemList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QStandardItemList) Detach_helper_0() {
	// QStandardItemList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QStandardItemList) Detach_helper_1() {
	// QStandardItemList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QStandardItemList) Dealloc_0() {
	// QStandardItemList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QStandardItemList) Node_construct_0() {
	// QStandardItemList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QStandardItemList) Node_destruct_0() {
	// QStandardItemList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QStandardItemList) Node_copy_0() {
	// QStandardItemList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QStandardItemList) Node_destruct_1() {
	// QStandardItemList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QStandardItemList) IsValidIterator_0() bool {
	// QStandardItemList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QStandardItemList) Op_eq_impl_0() bool {
	// QStandardItemList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QStandardItemList) Op_eq_impl_1() bool {
	// QStandardItemList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QStandardItemList) Contains_impl_0() bool {
	// QStandardItemList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QStandardItemList) Contains_impl_1() bool {
	// QStandardItemList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QStandardItemList) Count_impl_0() int {
	// QStandardItemList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QStandardItemList) Count_impl_1() int {
	// QStandardItemList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QStandardItemList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
