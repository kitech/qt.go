package qtnetwork

// /usr/include/qt/QtNetwork/qnetworkconfiguration.h
// #include <qnetworkconfiguration.h>
// #include <QtNetwork>

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
type QNetworkConfigurationList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QNetworkConfigurationList) Operator_equal_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QNetworkConfigurationList) Operator_equal_1() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QNetworkConfigurationList) Swap_0() {
	// QNetworkConfigurationList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QNetworkConfigurationList) Operator_equal_equal_0() bool {
	// QNetworkConfigurationList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QNetworkConfigurationList) Operator_not_equal_0() bool {
	// QNetworkConfigurationList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QNetworkConfigurationList) Size_0() int {
	// QNetworkConfigurationList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QNetworkConfigurationList) Detach_0() {
	// QNetworkConfigurationList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QNetworkConfigurationList) DetachShared_0() {
	// QNetworkConfigurationList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QNetworkConfigurationList) IsDetached_0() bool {
	// QNetworkConfigurationList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QNetworkConfigurationList) SetSharable_0() {
	// QNetworkConfigurationList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QNetworkConfigurationList) IsSharedWith_0() bool {
	// QNetworkConfigurationList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QNetworkConfigurationList) IsEmpty_0() bool {
	// QNetworkConfigurationList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QNetworkConfigurationList) Clear_0() {
	// QNetworkConfigurationList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QNetworkConfigurationList) At_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & operator[](int)
func (this *QNetworkConfigurationList) Operator_get_index_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// T & operator[](int)
func (this *QNetworkConfigurationList) Operator_get_index_1() *QNetworkConfiguration {
	// QNetworkConfigurationList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// void reserve(int)
func (this *QNetworkConfigurationList) Reserve_0() {
	// QNetworkConfigurationList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QNetworkConfigurationList) Append_0() {
	// QNetworkConfigurationList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QNetworkConfigurationList) Append_1() {
	// QNetworkConfigurationList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QNetworkConfigurationList) Prepend_0() {
	// QNetworkConfigurationList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QNetworkConfigurationList) Insert_0() {
	// QNetworkConfigurationList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QNetworkConfigurationList) Replace_0() {
	// QNetworkConfigurationList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QNetworkConfigurationList) RemoveAt_0() {
	// QNetworkConfigurationList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QNetworkConfigurationList) RemoveAll_0() int {
	// QNetworkConfigurationList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QNetworkConfigurationList) RemoveOne_0() bool {
	// QNetworkConfigurationList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QNetworkConfigurationList) TakeAt_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// T takeFirst()
func (this *QNetworkConfigurationList) TakeFirst_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// T takeLast()
func (this *QNetworkConfigurationList) TakeLast_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// void move(int, int)
func (this *QNetworkConfigurationList) Move_0() {
	// QNetworkConfigurationList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QNetworkConfigurationList) Swap_1() {
	// QNetworkConfigurationList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QNetworkConfigurationList) IndexOf_0() int {
	// QNetworkConfigurationList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QNetworkConfigurationList) LastIndexOf_0() int {
	// QNetworkConfigurationList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QNetworkConfigurationList) Contains_0() bool {
	// QNetworkConfigurationList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QNetworkConfigurationList) Count_0() int {
	// QNetworkConfigurationList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QNetworkConfigurationList) Begin_0() {
	// QNetworkConfigurationList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QNetworkConfigurationList) Begin_1() {
	// QNetworkConfigurationList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QNetworkConfigurationList) Cbegin_0() {
	// QNetworkConfigurationList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QNetworkConfigurationList) ConstBegin_0() {
	// QNetworkConfigurationList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QNetworkConfigurationList) End_0() {
	// QNetworkConfigurationList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QNetworkConfigurationList) End_1() {
	// QNetworkConfigurationList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QNetworkConfigurationList) Cend_0() {
	// QNetworkConfigurationList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QNetworkConfigurationList) ConstEnd_0() {
	// QNetworkConfigurationList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QNetworkConfigurationList) Rbegin_0() {
	// QNetworkConfigurationList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QNetworkConfigurationList) Rend_0() {
	// QNetworkConfigurationList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QNetworkConfigurationList) Rbegin_1() {
	// QNetworkConfigurationList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QNetworkConfigurationList) Rend_1() {
	// QNetworkConfigurationList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QNetworkConfigurationList) Crbegin_0() {
	// QNetworkConfigurationList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QNetworkConfigurationList) Crend_0() {
	// QNetworkConfigurationList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QNetworkConfigurationList) Insert_1() {
	// QNetworkConfigurationList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QNetworkConfigurationList) Erase_0() {
	// QNetworkConfigurationList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QNetworkConfigurationList) Erase_1() {
	// QNetworkConfigurationList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QNetworkConfigurationList) Count_1() int {
	// QNetworkConfigurationList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QNetworkConfigurationList) Length_0() int {
	// QNetworkConfigurationList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QNetworkConfigurationList) First_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & constFirst()
func (this *QNetworkConfigurationList) ConstFirst_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & first()
func (this *QNetworkConfigurationList) First_1() *QNetworkConfiguration {
	// QNetworkConfigurationList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// T & last()
func (this *QNetworkConfigurationList) Last_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & last()
func (this *QNetworkConfigurationList) Last_1() *QNetworkConfiguration {
	// QNetworkConfigurationList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & constLast()
func (this *QNetworkConfigurationList) ConstLast_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// void removeFirst()
func (this *QNetworkConfigurationList) RemoveFirst_0() {
	// QNetworkConfigurationList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QNetworkConfigurationList) RemoveLast_0() {
	// QNetworkConfigurationList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QNetworkConfigurationList) StartsWith_0() bool {
	// QNetworkConfigurationList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QNetworkConfigurationList) EndsWith_0() bool {
	// QNetworkConfigurationList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QNetworkConfigurationList) Mid_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QNetworkConfigurationList) Value_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// T value(int, const T &)
func (this *QNetworkConfigurationList) Value_1() *QNetworkConfiguration {
	// QNetworkConfigurationList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// void push_back(const T &)
func (this *QNetworkConfigurationList) Push_back_0() {
	// QNetworkConfigurationList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QNetworkConfigurationList) Push_front_0() {
	// QNetworkConfigurationList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QNetworkConfigurationList) Front_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & front()
func (this *QNetworkConfigurationList) Front_1() *QNetworkConfiguration {
	// QNetworkConfigurationList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// T & back()
func (this *QNetworkConfigurationList) Back_0() *QNetworkConfiguration {
	// QNetworkConfigurationList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// const T & back()
func (this *QNetworkConfigurationList) Back_1() *QNetworkConfiguration {
	// QNetworkConfigurationList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QNetworkConfiguration{}
}

// void pop_front()
func (this *QNetworkConfigurationList) Pop_front_0() {
	// QNetworkConfigurationList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QNetworkConfigurationList) Pop_back_0() {
	// QNetworkConfigurationList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QNetworkConfigurationList) Empty_0() bool {
	// QNetworkConfigurationList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QNetworkConfigurationList) Operator_add_equal_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QNetworkConfigurationList) Operator_add_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QNetworkConfigurationList) Operator_add_equal_1() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QNetworkConfigurationList) Operator_left_shift_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QNetworkConfigurationList) Operator_left_shift_1() *QNetworkConfigurationList {
	// QNetworkConfigurationList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QNetworkConfigurationList) ToVector_0() {
	// QNetworkConfigurationList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QNetworkConfigurationList) ToSet_0() {
	// QNetworkConfigurationList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QNetworkConfigurationList) FromVector_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QNetworkConfigurationList) FromSet_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QNetworkConfigurationList) FromStdList_0() *QNetworkConfigurationList {
	// QNetworkConfigurationList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QNetworkConfigurationList) ToStdList_0() {
	// QNetworkConfigurationList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QNetworkConfigurationList) Detach_helper_grow_0() {
	// QNetworkConfigurationList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QNetworkConfigurationList) Detach_helper_0() {
	// QNetworkConfigurationList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QNetworkConfigurationList) Detach_helper_1() {
	// QNetworkConfigurationList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QNetworkConfigurationList) Dealloc_0() {
	// QNetworkConfigurationList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QNetworkConfigurationList) Node_construct_0() {
	// QNetworkConfigurationList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QNetworkConfigurationList) Node_destruct_0() {
	// QNetworkConfigurationList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QNetworkConfigurationList) Node_copy_0() {
	// QNetworkConfigurationList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QNetworkConfigurationList) Node_destruct_1() {
	// QNetworkConfigurationList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QNetworkConfigurationList) IsValidIterator_0() bool {
	// QNetworkConfigurationList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QNetworkConfigurationList) Op_eq_impl_0() bool {
	// QNetworkConfigurationList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QNetworkConfigurationList) Op_eq_impl_1() bool {
	// QNetworkConfigurationList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QNetworkConfigurationList) Contains_impl_0() bool {
	// QNetworkConfigurationList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QNetworkConfigurationList) Contains_impl_1() bool {
	// QNetworkConfigurationList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QNetworkConfigurationList) Count_impl_0() int {
	// QNetworkConfigurationList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QNetworkConfigurationList) Count_impl_1() int {
	// QNetworkConfigurationList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QNetworkConfigurationList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
