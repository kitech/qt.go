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
type QVariantList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QVariantList) Operator_equal0() *QVariantList {
	// QVariantList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QVariantList) Operator_equal1() *QVariantList {
	// QVariantList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QVariantList) Swap0() {
	// QVariantList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QVariantList) Operator_equal_equal0() bool {
	// QVariantList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QVariantList) Operator_not_equal0() bool {
	// QVariantList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QVariantList) Size0() int {
	// QVariantList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QVariantList) Detach0() {
	// QVariantList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QVariantList) DetachShared0() {
	// QVariantList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QVariantList) IsDetached0() bool {
	// QVariantList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QVariantList) SetSharable0() {
	// QVariantList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QVariantList) IsSharedWith0() bool {
	// QVariantList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QVariantList) IsEmpty0() bool {
	// QVariantList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QVariantList) Clear0() {
	// QVariantList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QVariantList) At0() *QVariant {
	// QVariantList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & operator[](int)
func (this *QVariantList) Operator_get_index0() *QVariant {
	// QVariantList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & operator[](int)
func (this *QVariantList) Operator_get_index1() *QVariant {
	// QVariantList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void reserve(int)
func (this *QVariantList) Reserve0() {
	// QVariantList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QVariantList) Append0() {
	// QVariantList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QVariantList) Append1() {
	// QVariantList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QVariantList) Prepend0() {
	// QVariantList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QVariantList) Insert0() {
	// QVariantList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QVariantList) Replace0() {
	// QVariantList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QVariantList) RemoveAt0() {
	// QVariantList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QVariantList) RemoveAll0() int {
	// QVariantList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QVariantList) RemoveOne0() bool {
	// QVariantList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QVariantList) TakeAt0() *QVariant {
	// QVariantList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T takeFirst()
func (this *QVariantList) TakeFirst0() *QVariant {
	// QVariantList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T takeLast()
func (this *QVariantList) TakeLast0() *QVariant {
	// QVariantList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void move(int, int)
func (this *QVariantList) Move0() {
	// QVariantList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QVariantList) Swap1() {
	// QVariantList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QVariantList) IndexOf0() int {
	// QVariantList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QVariantList) LastIndexOf0() int {
	// QVariantList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QVariantList) Contains0() bool {
	// QVariantList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QVariantList) Count0() int {
	// QVariantList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QVariantList) Begin0() {
	// QVariantList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QVariantList) Begin1() {
	// QVariantList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QVariantList) Cbegin0() {
	// QVariantList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QVariantList) ConstBegin0() {
	// QVariantList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QVariantList) End0() {
	// QVariantList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QVariantList) End1() {
	// QVariantList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QVariantList) Cend0() {
	// QVariantList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QVariantList) ConstEnd0() {
	// QVariantList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QVariantList) Rbegin0() {
	// QVariantList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QVariantList) Rend0() {
	// QVariantList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QVariantList) Rbegin1() {
	// QVariantList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QVariantList) Rend1() {
	// QVariantList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QVariantList) Crbegin0() {
	// QVariantList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QVariantList) Crend0() {
	// QVariantList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QVariantList) Insert1() {
	// QVariantList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QVariantList) Erase0() {
	// QVariantList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QVariantList) Erase1() {
	// QVariantList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QVariantList) Count1() int {
	// QVariantList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QVariantList) Length0() int {
	// QVariantList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QVariantList) First0() *QVariant {
	// QVariantList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & constFirst()
func (this *QVariantList) ConstFirst0() *QVariant {
	// QVariantList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & first()
func (this *QVariantList) First1() *QVariant {
	// QVariantList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & last()
func (this *QVariantList) Last0() *QVariant {
	// QVariantList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & last()
func (this *QVariantList) Last1() *QVariant {
	// QVariantList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & constLast()
func (this *QVariantList) ConstLast0() *QVariant {
	// QVariantList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void removeFirst()
func (this *QVariantList) RemoveFirst0() {
	// QVariantList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QVariantList) RemoveLast0() {
	// QVariantList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QVariantList) StartsWith0() bool {
	// QVariantList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QVariantList) EndsWith0() bool {
	// QVariantList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QVariantList) Mid0() *QVariantList {
	// QVariantList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QVariantList) Value0() *QVariant {
	// QVariantList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T value(int, const T &)
func (this *QVariantList) Value1() *QVariant {
	// QVariantList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void push_back(const T &)
func (this *QVariantList) Push_back0() {
	// QVariantList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QVariantList) Push_front0() {
	// QVariantList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QVariantList) Front0() *QVariant {
	// QVariantList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & front()
func (this *QVariantList) Front1() *QVariant {
	// QVariantList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// T & back()
func (this *QVariantList) Back0() *QVariant {
	// QVariantList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// const T & back()
func (this *QVariantList) Back1() *QVariant {
	// QVariantList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QVariant{}
}

// void pop_front()
func (this *QVariantList) Pop_front0() {
	// QVariantList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QVariantList) Pop_back0() {
	// QVariantList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QVariantList) Empty0() bool {
	// QVariantList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QVariantList) Operator_add_equal0() *QVariantList {
	// QVariantList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QVariantList) Operator_add0() *QVariantList {
	// QVariantList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QVariantList) Operator_add_equal1() *QVariantList {
	// QVariantList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QVariantList) Operator_left_shift0() *QVariantList {
	// QVariantList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QVariantList) Operator_left_shift1() *QVariantList {
	// QVariantList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QVariantList) ToVector0() {
	// QVariantList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QVariantList) ToSet0() {
	// QVariantList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QVariantList) FromVector0() *QVariantList {
	// QVariantList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QVariantList) FromSet0() *QVariantList {
	// QVariantList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QVariantList) FromStdList0() *QVariantList {
	// QVariantList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QVariantList) ToStdList0() {
	// QVariantList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QVariantList) Detach_helper_grow0() {
	// QVariantList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QVariantList) Detach_helper0() {
	// QVariantList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QVariantList) Detach_helper1() {
	// QVariantList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QVariantList) Dealloc0() {
	// QVariantList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QVariantList) Node_construct0() {
	// QVariantList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QVariantList) Node_destruct0() {
	// QVariantList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QVariantList) Node_copy0() {
	// QVariantList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QVariantList) Node_destruct1() {
	// QVariantList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QVariantList) IsValidIterator0() bool {
	// QVariantList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Op_eq_impl0() bool {
	// QVariantList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QVariantList) Op_eq_impl1() bool {
	// QVariantList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Contains_impl0() bool {
	// QVariantList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QVariantList) Contains_impl1() bool {
	// QVariantList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QVariantList) Count_impl0() int {
	// QVariantList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QVariantList) Count_impl1() int {
	// QVariantList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QVariantList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
