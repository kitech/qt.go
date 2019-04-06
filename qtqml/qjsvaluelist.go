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

func init_unused_10067() {
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
func (this *QJSValueList) Operator_equal0() *QJSValueList {
	// QJSValueList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QJSValueList) Operator_equal1() *QJSValueList {
	// QJSValueList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QJSValueList) Swap0() {
	// QJSValueList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QJSValueList) Operator_equal_equal0() bool {
	// QJSValueList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QJSValueList) Operator_not_equal0() bool {
	// QJSValueList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QJSValueList) Size0() int {
	// QJSValueList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QJSValueList) Detach0() {
	// QJSValueList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QJSValueList) DetachShared0() {
	// QJSValueList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QJSValueList) IsDetached0() bool {
	// QJSValueList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QJSValueList) SetSharable0() {
	// QJSValueList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QJSValueList) IsSharedWith0() bool {
	// QJSValueList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QJSValueList) IsEmpty0() bool {
	// QJSValueList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QJSValueList) Clear0() {
	// QJSValueList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QJSValueList) At0() *QJSValue {
	// QJSValueList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & operator[](int)
func (this *QJSValueList) Operator_get_index0() *QJSValue {
	// QJSValueList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T & operator[](int)
func (this *QJSValueList) Operator_get_index1() *QJSValue {
	// QJSValueList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void reserve(int)
func (this *QJSValueList) Reserve0() {
	// QJSValueList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QJSValueList) Append0() {
	// QJSValueList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QJSValueList) Append1() {
	// QJSValueList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QJSValueList) Prepend0() {
	// QJSValueList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QJSValueList) Insert0() {
	// QJSValueList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QJSValueList) Replace0() {
	// QJSValueList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QJSValueList) RemoveAt0() {
	// QJSValueList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QJSValueList) RemoveAll0() int {
	// QJSValueList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QJSValueList) RemoveOne0() bool {
	// QJSValueList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QJSValueList) TakeAt0() *QJSValue {
	// QJSValueList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T takeFirst()
func (this *QJSValueList) TakeFirst0() *QJSValue {
	// QJSValueList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T takeLast()
func (this *QJSValueList) TakeLast0() *QJSValue {
	// QJSValueList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void move(int, int)
func (this *QJSValueList) Move0() {
	// QJSValueList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QJSValueList) Swap1() {
	// QJSValueList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QJSValueList) IndexOf0() int {
	// QJSValueList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QJSValueList) LastIndexOf0() int {
	// QJSValueList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QJSValueList) Contains0() bool {
	// QJSValueList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QJSValueList) Count0() int {
	// QJSValueList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QJSValueList) Begin0() {
	// QJSValueList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QJSValueList) Begin1() {
	// QJSValueList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QJSValueList) Cbegin0() {
	// QJSValueList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QJSValueList) ConstBegin0() {
	// QJSValueList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QJSValueList) End0() {
	// QJSValueList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QJSValueList) End1() {
	// QJSValueList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QJSValueList) Cend0() {
	// QJSValueList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QJSValueList) ConstEnd0() {
	// QJSValueList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QJSValueList) Rbegin0() {
	// QJSValueList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QJSValueList) Rend0() {
	// QJSValueList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QJSValueList) Rbegin1() {
	// QJSValueList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QJSValueList) Rend1() {
	// QJSValueList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QJSValueList) Crbegin0() {
	// QJSValueList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QJSValueList) Crend0() {
	// QJSValueList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QJSValueList) Insert1() {
	// QJSValueList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QJSValueList) Erase0() {
	// QJSValueList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QJSValueList) Erase1() {
	// QJSValueList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QJSValueList) Count1() int {
	// QJSValueList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QJSValueList) Length0() int {
	// QJSValueList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QJSValueList) First0() *QJSValue {
	// QJSValueList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & constFirst()
func (this *QJSValueList) ConstFirst0() *QJSValue {
	// QJSValueList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & first()
func (this *QJSValueList) First1() *QJSValue {
	// QJSValueList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T & last()
func (this *QJSValueList) Last0() *QJSValue {
	// QJSValueList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & last()
func (this *QJSValueList) Last1() *QJSValue {
	// QJSValueList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & constLast()
func (this *QJSValueList) ConstLast0() *QJSValue {
	// QJSValueList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void removeFirst()
func (this *QJSValueList) RemoveFirst0() {
	// QJSValueList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QJSValueList) RemoveLast0() {
	// QJSValueList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QJSValueList) StartsWith0() bool {
	// QJSValueList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QJSValueList) EndsWith0() bool {
	// QJSValueList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QJSValueList) Mid0() *QJSValueList {
	// QJSValueList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QJSValueList) Value0() *QJSValue {
	// QJSValueList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T value(int, const T &)
func (this *QJSValueList) Value1() *QJSValue {
	// QJSValueList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void push_back(const T &)
func (this *QJSValueList) Push_back0() {
	// QJSValueList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QJSValueList) Push_front0() {
	// QJSValueList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QJSValueList) Front0() *QJSValue {
	// QJSValueList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & front()
func (this *QJSValueList) Front1() *QJSValue {
	// QJSValueList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// T & back()
func (this *QJSValueList) Back0() *QJSValue {
	// QJSValueList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// const T & back()
func (this *QJSValueList) Back1() *QJSValue {
	// QJSValueList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QJSValue{}
}

// void pop_front()
func (this *QJSValueList) Pop_front0() {
	// QJSValueList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QJSValueList) Pop_back0() {
	// QJSValueList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QJSValueList) Empty0() bool {
	// QJSValueList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QJSValueList) Operator_add_equal0() *QJSValueList {
	// QJSValueList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QJSValueList) Operator_add0() *QJSValueList {
	// QJSValueList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QJSValueList) Operator_add_equal1() *QJSValueList {
	// QJSValueList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QJSValueList) Operator_left_shift0() *QJSValueList {
	// QJSValueList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QJSValueList) Operator_left_shift1() *QJSValueList {
	// QJSValueList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QJSValueList) ToVector0() {
	// QJSValueList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QJSValueList) ToSet0() {
	// QJSValueList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QJSValueList) FromVector0() *QJSValueList {
	// QJSValueList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QJSValueList) FromSet0() *QJSValueList {
	// QJSValueList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QJSValueList) FromStdList0() *QJSValueList {
	// QJSValueList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QJSValueList) ToStdList0() {
	// QJSValueList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QJSValueList) Detach_helper_grow0() {
	// QJSValueList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QJSValueList) Detach_helper0() {
	// QJSValueList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QJSValueList) Detach_helper1() {
	// QJSValueList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QJSValueList) Dealloc0() {
	// QJSValueList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QJSValueList) Node_construct0() {
	// QJSValueList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QJSValueList) Node_destruct0() {
	// QJSValueList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QJSValueList) Node_copy0() {
	// QJSValueList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QJSValueList) Node_destruct1() {
	// QJSValueList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QJSValueList) IsValidIterator0() bool {
	// QJSValueList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QJSValueList) Op_eq_impl0() bool {
	// QJSValueList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QJSValueList) Op_eq_impl1() bool {
	// QJSValueList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QJSValueList) Contains_impl0() bool {
	// QJSValueList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QJSValueList) Contains_impl1() bool {
	// QJSValueList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QJSValueList) Count_impl0() int {
	// QJSValueList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QJSValueList) Count_impl1() int {
	// QJSValueList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QJSValueList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
