package qtgui

// /usr/include/qt/QtGui/qaccessible.h
// #include <qaccessible.h>
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
type QAccessibleInterfaceList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_equal_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QAccessibleInterfaceList) Operator_equal_1() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QAccessibleInterfaceList) Swap_0() {
	// QAccessibleInterfaceList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_equal_equal_0() bool {
	// QAccessibleInterfaceList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_not_equal_0() bool {
	// QAccessibleInterfaceList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QAccessibleInterfaceList) Size_0() int {
	// QAccessibleInterfaceList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QAccessibleInterfaceList) Detach_0() {
	// QAccessibleInterfaceList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QAccessibleInterfaceList) DetachShared_0() {
	// QAccessibleInterfaceList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QAccessibleInterfaceList) IsDetached_0() bool {
	// QAccessibleInterfaceList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QAccessibleInterfaceList) SetSharable_0() {
	// QAccessibleInterfaceList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QAccessibleInterfaceList) IsSharedWith_0() bool {
	// QAccessibleInterfaceList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QAccessibleInterfaceList) IsEmpty_0() bool {
	// QAccessibleInterfaceList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QAccessibleInterfaceList) Clear_0() {
	// QAccessibleInterfaceList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QAccessibleInterfaceList) At_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & operator[](int)
func (this *QAccessibleInterfaceList) Operator_get_index_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T & operator[](int)
func (this *QAccessibleInterfaceList) Operator_get_index_1() *QAccessibleInterface {
	// QAccessibleInterfaceList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void reserve(int)
func (this *QAccessibleInterfaceList) Reserve_0() {
	// QAccessibleInterfaceList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QAccessibleInterfaceList) Append_0() {
	// QAccessibleInterfaceList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QAccessibleInterfaceList) Append_1() {
	// QAccessibleInterfaceList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QAccessibleInterfaceList) Prepend_0() {
	// QAccessibleInterfaceList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QAccessibleInterfaceList) Insert_0() {
	// QAccessibleInterfaceList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QAccessibleInterfaceList) Replace_0() {
	// QAccessibleInterfaceList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QAccessibleInterfaceList) RemoveAt_0() {
	// QAccessibleInterfaceList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QAccessibleInterfaceList) RemoveAll_0() int {
	// QAccessibleInterfaceList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QAccessibleInterfaceList) RemoveOne_0() bool {
	// QAccessibleInterfaceList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QAccessibleInterfaceList) TakeAt_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T takeFirst()
func (this *QAccessibleInterfaceList) TakeFirst_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T takeLast()
func (this *QAccessibleInterfaceList) TakeLast_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void move(int, int)
func (this *QAccessibleInterfaceList) Move_0() {
	// QAccessibleInterfaceList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QAccessibleInterfaceList) Swap_1() {
	// QAccessibleInterfaceList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QAccessibleInterfaceList) IndexOf_0() int {
	// QAccessibleInterfaceList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QAccessibleInterfaceList) LastIndexOf_0() int {
	// QAccessibleInterfaceList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QAccessibleInterfaceList) Contains_0() bool {
	// QAccessibleInterfaceList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QAccessibleInterfaceList) Count_0() int {
	// QAccessibleInterfaceList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QAccessibleInterfaceList) Begin_0() {
	// QAccessibleInterfaceList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QAccessibleInterfaceList) Begin_1() {
	// QAccessibleInterfaceList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QAccessibleInterfaceList) Cbegin_0() {
	// QAccessibleInterfaceList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QAccessibleInterfaceList) ConstBegin_0() {
	// QAccessibleInterfaceList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QAccessibleInterfaceList) End_0() {
	// QAccessibleInterfaceList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QAccessibleInterfaceList) End_1() {
	// QAccessibleInterfaceList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QAccessibleInterfaceList) Cend_0() {
	// QAccessibleInterfaceList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QAccessibleInterfaceList) ConstEnd_0() {
	// QAccessibleInterfaceList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QAccessibleInterfaceList) Rbegin_0() {
	// QAccessibleInterfaceList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QAccessibleInterfaceList) Rend_0() {
	// QAccessibleInterfaceList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QAccessibleInterfaceList) Rbegin_1() {
	// QAccessibleInterfaceList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QAccessibleInterfaceList) Rend_1() {
	// QAccessibleInterfaceList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QAccessibleInterfaceList) Crbegin_0() {
	// QAccessibleInterfaceList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QAccessibleInterfaceList) Crend_0() {
	// QAccessibleInterfaceList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QAccessibleInterfaceList) Insert_1() {
	// QAccessibleInterfaceList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QAccessibleInterfaceList) Erase_0() {
	// QAccessibleInterfaceList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QAccessibleInterfaceList) Erase_1() {
	// QAccessibleInterfaceList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QAccessibleInterfaceList) Count_1() int {
	// QAccessibleInterfaceList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QAccessibleInterfaceList) Length_0() int {
	// QAccessibleInterfaceList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QAccessibleInterfaceList) First_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & constFirst()
func (this *QAccessibleInterfaceList) ConstFirst_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & first()
func (this *QAccessibleInterfaceList) First_1() *QAccessibleInterface {
	// QAccessibleInterfaceList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T & last()
func (this *QAccessibleInterfaceList) Last_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & last()
func (this *QAccessibleInterfaceList) Last_1() *QAccessibleInterface {
	// QAccessibleInterfaceList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & constLast()
func (this *QAccessibleInterfaceList) ConstLast_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void removeFirst()
func (this *QAccessibleInterfaceList) RemoveFirst_0() {
	// QAccessibleInterfaceList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QAccessibleInterfaceList) RemoveLast_0() {
	// QAccessibleInterfaceList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QAccessibleInterfaceList) StartsWith_0() bool {
	// QAccessibleInterfaceList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QAccessibleInterfaceList) EndsWith_0() bool {
	// QAccessibleInterfaceList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QAccessibleInterfaceList) Mid_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QAccessibleInterfaceList) Value_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T value(int, const T &)
func (this *QAccessibleInterfaceList) Value_1() *QAccessibleInterface {
	// QAccessibleInterfaceList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void push_back(const T &)
func (this *QAccessibleInterfaceList) Push_back_0() {
	// QAccessibleInterfaceList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QAccessibleInterfaceList) Push_front_0() {
	// QAccessibleInterfaceList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QAccessibleInterfaceList) Front_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & front()
func (this *QAccessibleInterfaceList) Front_1() *QAccessibleInterface {
	// QAccessibleInterfaceList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// T & back()
func (this *QAccessibleInterfaceList) Back_0() *QAccessibleInterface {
	// QAccessibleInterfaceList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// const T & back()
func (this *QAccessibleInterfaceList) Back_1() *QAccessibleInterface {
	// QAccessibleInterfaceList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QAccessibleInterface{}
}

// void pop_front()
func (this *QAccessibleInterfaceList) Pop_front_0() {
	// QAccessibleInterfaceList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QAccessibleInterfaceList) Pop_back_0() {
	// QAccessibleInterfaceList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QAccessibleInterfaceList) Empty_0() bool {
	// QAccessibleInterfaceList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_add_equal_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_add_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QAccessibleInterfaceList) Operator_add_equal_1() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QAccessibleInterfaceList) Operator_left_shift_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QAccessibleInterfaceList) Operator_left_shift_1() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QAccessibleInterfaceList) ToVector_0() {
	// QAccessibleInterfaceList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QAccessibleInterfaceList) ToSet_0() {
	// QAccessibleInterfaceList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QAccessibleInterfaceList) FromVector_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QAccessibleInterfaceList) FromSet_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QAccessibleInterfaceList) FromStdList_0() *QAccessibleInterfaceList {
	// QAccessibleInterfaceList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QAccessibleInterfaceList) ToStdList_0() {
	// QAccessibleInterfaceList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QAccessibleInterfaceList) Detach_helper_grow_0() {
	// QAccessibleInterfaceList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QAccessibleInterfaceList) Detach_helper_0() {
	// QAccessibleInterfaceList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QAccessibleInterfaceList) Detach_helper_1() {
	// QAccessibleInterfaceList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QAccessibleInterfaceList) Dealloc_0() {
	// QAccessibleInterfaceList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QAccessibleInterfaceList) Node_construct_0() {
	// QAccessibleInterfaceList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QAccessibleInterfaceList) Node_destruct_0() {
	// QAccessibleInterfaceList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QAccessibleInterfaceList) Node_copy_0() {
	// QAccessibleInterfaceList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QAccessibleInterfaceList) Node_destruct_1() {
	// QAccessibleInterfaceList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QAccessibleInterfaceList) IsValidIterator_0() bool {
	// QAccessibleInterfaceList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Op_eq_impl_0() bool {
	// QAccessibleInterfaceList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Op_eq_impl_1() bool {
	// QAccessibleInterfaceList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Contains_impl_0() bool {
	// QAccessibleInterfaceList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Contains_impl_1() bool {
	// QAccessibleInterfaceList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Count_impl_0() int {
	// QAccessibleInterfaceList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QAccessibleInterfaceList) Count_impl_1() int {
	// QAccessibleInterfaceList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QAccessibleInterfaceList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
