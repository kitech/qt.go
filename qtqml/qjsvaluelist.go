package qtqml

// /usr/include/qt/QtQml/qjsvalue.h
// #include <qjsvalue.h>
// #include <QtQml>

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
		qtnetwork.KeepMe()
	}
}

//  keep block end

//  body block begin
type QJSValueList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QJSValueList) Operator_equal_0() *QJSValueList {
	// QJSValueList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QJSValueList) Operator_equal_1() *QJSValueList {
	// QJSValueList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QJSValueList) Swap_0() {
	// QJSValueList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QJSValueList) Operator_equal_equal_0() bool {
	// QJSValueList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QJSValueList) Operator_not_equal_0() bool {
	// QJSValueList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QJSValueList) Size_0() int {
	// QJSValueList_size_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QJSValueList) Detach_0() {
	// QJSValueList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QJSValueList) DetachShared_0() {
	// QJSValueList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QJSValueList) IsDetached_0() bool {
	// QJSValueList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(_Bool)
func (this *QJSValueList) SetSharable_0() {
	// QJSValueList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QJSValueList) IsSharedWith_0() bool {
	// QJSValueList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QJSValueList) IsEmpty_0() bool {
	// QJSValueList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QJSValueList) Clear_0() {
	// QJSValueList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QJSValueList) At_0() *QJSValue {
	// QJSValueList_at_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & operator[](int)
func (this *QJSValueList) Operator_get_index_0() *QJSValue {
	// QJSValueList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T & operator[](int)
func (this *QJSValueList) Operator_get_index_1() *QJSValue {
	// QJSValueList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void reserve(int)
func (this *QJSValueList) Reserve_0() {
	// QJSValueList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QJSValueList) Append_0() {
	// QJSValueList_append_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QJSValueList) Append_1() {
	// QJSValueList_append_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QJSValueList) Prepend_0() {
	// QJSValueList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QJSValueList) Insert_0() {
	// QJSValueList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QJSValueList) Replace_0() {
	// QJSValueList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QJSValueList) RemoveAt_0() {
	// QJSValueList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QJSValueList) RemoveAll_0() int {
	// QJSValueList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QJSValueList) RemoveOne_0() bool {
	// QJSValueList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QJSValueList) TakeAt_0() *QJSValue {
	// QJSValueList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T takeFirst()
func (this *QJSValueList) TakeFirst_0() *QJSValue {
	// QJSValueList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T takeLast()
func (this *QJSValueList) TakeLast_0() *QJSValue {
	// QJSValueList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void move(int, int)
func (this *QJSValueList) Move_0() {
	// QJSValueList_move_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QJSValueList) Swap_1() {
	// QJSValueList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QJSValueList) IndexOf_0() int {
	// QJSValueList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QJSValueList) LastIndexOf_0() int {
	// QJSValueList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QJSValueList) Contains_0() bool {
	// QJSValueList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QJSValueList) Count_0() int {
	// QJSValueList_count_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QJSValueList) Begin_0() {
	// QJSValueList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QJSValueList) Begin_1() {
	// QJSValueList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QJSValueList) Cbegin_0() {
	// QJSValueList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QJSValueList) ConstBegin_0() {
	// QJSValueList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QJSValueList) End_0() {
	// QJSValueList_end_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QJSValueList) End_1() {
	// QJSValueList_end_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QJSValueList) Cend_0() {
	// QJSValueList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QJSValueList) ConstEnd_0() {
	// QJSValueList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QJSValueList) Rbegin_0() {
	// QJSValueList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QJSValueList) Rend_0() {
	// QJSValueList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QJSValueList) Rbegin_1() {
	// QJSValueList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QJSValueList) Rend_1() {
	// QJSValueList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QJSValueList) Crbegin_0() {
	// QJSValueList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QJSValueList) Crend_0() {
	// QJSValueList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(class QList::iterator, const T &)
func (this *QJSValueList) Insert_1() {
	// QJSValueList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator)
func (this *QJSValueList) Erase_0() {
	// QJSValueList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(class QList::iterator, class QList::iterator)
func (this *QJSValueList) Erase_1() {
	// QJSValueList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QJSValueList) Count_1() int {
	// QJSValueList_count_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QJSValueList) Length_0() int {
	// QJSValueList_length_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QJSValueList) First_0() *QJSValue {
	// QJSValueList_first_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & constFirst()
func (this *QJSValueList) ConstFirst_0() *QJSValue {
	// QJSValueList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & first()
func (this *QJSValueList) First_1() *QJSValue {
	// QJSValueList_first_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T & last()
func (this *QJSValueList) Last_0() *QJSValue {
	// QJSValueList_last_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & last()
func (this *QJSValueList) Last_1() *QJSValue {
	// QJSValueList_last_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & constLast()
func (this *QJSValueList) ConstLast_0() *QJSValue {
	// QJSValueList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void removeFirst()
func (this *QJSValueList) RemoveFirst_0() {
	// QJSValueList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QJSValueList) RemoveLast_0() {
	// QJSValueList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QJSValueList) StartsWith_0() bool {
	// QJSValueList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QJSValueList) EndsWith_0() bool {
	// QJSValueList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QJSValueList) Mid_0() *QJSValueList {
	// QJSValueList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QJSValueList) Value_0() *QJSValue {
	// QJSValueList_value_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T value(int, const T &)
func (this *QJSValueList) Value_1() *QJSValue {
	// QJSValueList_value_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void push_back(const T &)
func (this *QJSValueList) Push_back_0() {
	// QJSValueList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QJSValueList) Push_front_0() {
	// QJSValueList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QJSValueList) Front_0() *QJSValue {
	// QJSValueList_front_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & front()
func (this *QJSValueList) Front_1() *QJSValue {
	// QJSValueList_front_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T & back()
func (this *QJSValueList) Back_0() *QJSValue {
	// QJSValueList_back_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & back()
func (this *QJSValueList) Back_1() *QJSValue {
	// QJSValueList_back_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void pop_front()
func (this *QJSValueList) Pop_front_0() {
	// QJSValueList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QJSValueList) Pop_back_0() {
	// QJSValueList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QJSValueList) Empty_0() bool {
	// QJSValueList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QJSValueList) Operator_add_equal_0() *QJSValueList {
	// QJSValueList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QJSValueList) Operator_add_0() *QJSValueList {
	// QJSValueList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QJSValueList) Operator_add_equal_1() *QJSValueList {
	// QJSValueList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QJSValueList) Operator_left_shift_0() *QJSValueList {
	// QJSValueList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QJSValueList) Operator_left_shift_1() *QJSValueList {
	// QJSValueList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QJSValueList) ToVector_0() {
	// QJSValueList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QJSValueList) ToSet_0() {
	// QJSValueList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QJSValueList) FromVector_0() *QJSValueList {
	// QJSValueList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QJSValueList) FromSet_0() *QJSValueList {
	// QJSValueList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QJSValueList) FromStdList_0() *QJSValueList {
	// QJSValueList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QJSValueList) ToStdList_0() {
	// QJSValueList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QJSValueList) Detach_helper_grow_0() {
	// QJSValueList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QJSValueList) Detach_helper_0() {
	// QJSValueList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QJSValueList) Detach_helper_1() {
	// QJSValueList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(struct QListData::Data *)
func (this *QJSValueList) Dealloc_0() {
	// QJSValueList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(struct QList::Node *, const T &)
func (this *QJSValueList) Node_construct_0() {
	// QJSValueList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *)
func (this *QJSValueList) Node_destruct_0() {
	// QJSValueList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(struct QList::Node *, struct QList::Node *, struct QList::Node *)
func (this *QJSValueList) Node_copy_0() {
	// QJSValueList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(struct QList::Node *, struct QList::Node *)
func (this *QJSValueList) Node_destruct_1() {
	// QJSValueList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const class QList::iterator &)
func (this *QJSValueList) IsValidIterator_0() bool {
	// QJSValueList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::NotArrayCompatibleLayout)
func (this *QJSValueList) Op_eq_impl_0() bool {
	// QJSValueList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, struct QListData::ArrayCompatibleLayout)
func (this *QJSValueList) Op_eq_impl_1() bool {
	// QJSValueList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QJSValueList) Contains_impl_0() bool {
	// QJSValueList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QJSValueList) Contains_impl_1() bool {
	// QJSValueList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, struct QListData::NotArrayCompatibleLayout)
func (this *QJSValueList) Count_impl_0() int {
	// QJSValueList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, struct QListData::ArrayCompatibleLayout)
func (this *QJSValueList) Count_impl_1() int {
	// QJSValueList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("QJSValueList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
