package qtcore

// /usr/include/qt/QtCore/qurl.h
// #include <qurl.h>
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

func init_unused_10079() {
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
type QUrlList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QUrlList) Operator_equal0() *QUrlList {
	// QUrlList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QUrlList) Operator_equal1() *QUrlList {
	// QUrlList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QUrlList) Swap0() {
	// QUrlList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QUrlList) Operator_equal_equal0() bool {
	// QUrlList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QUrlList) Operator_not_equal0() bool {
	// QUrlList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QUrlList) Size0() int {
	// QUrlList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QUrlList) Detach0() {
	// QUrlList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QUrlList) DetachShared0() {
	// QUrlList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QUrlList) IsDetached0() bool {
	// QUrlList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QUrlList) SetSharable0() {
	// QUrlList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QUrlList) IsSharedWith0() bool {
	// QUrlList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QUrlList) IsEmpty0() bool {
	// QUrlList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QUrlList) Clear0() {
	// QUrlList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QUrlList) At0() *QUrl {
	// QUrlList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & operator[](int)
func (this *QUrlList) Operator_get_index0() *QUrl {
	// QUrlList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// T & operator[](int)
func (this *QUrlList) Operator_get_index1() *QUrl {
	// QUrlList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// void reserve(int)
func (this *QUrlList) Reserve0() {
	// QUrlList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QUrlList) Append0() {
	// QUrlList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QUrlList) Append1() {
	// QUrlList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QUrlList) Prepend0() {
	// QUrlList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QUrlList) Insert0() {
	// QUrlList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QUrlList) Replace0() {
	// QUrlList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QUrlList) RemoveAt0() {
	// QUrlList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QUrlList) RemoveAll0() int {
	// QUrlList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QUrlList) RemoveOne0() bool {
	// QUrlList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QUrlList) TakeAt0() *QUrl {
	// QUrlList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// T takeFirst()
func (this *QUrlList) TakeFirst0() *QUrl {
	// QUrlList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// T takeLast()
func (this *QUrlList) TakeLast0() *QUrl {
	// QUrlList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// void move(int, int)
func (this *QUrlList) Move0() {
	// QUrlList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QUrlList) Swap1() {
	// QUrlList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QUrlList) IndexOf0() int {
	// QUrlList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QUrlList) LastIndexOf0() int {
	// QUrlList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QUrlList) Contains0() bool {
	// QUrlList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QUrlList) Count0() int {
	// QUrlList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QUrlList) Begin0() {
	// QUrlList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QUrlList) Begin1() {
	// QUrlList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QUrlList) Cbegin0() {
	// QUrlList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QUrlList) ConstBegin0() {
	// QUrlList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QUrlList) End0() {
	// QUrlList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QUrlList) End1() {
	// QUrlList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QUrlList) Cend0() {
	// QUrlList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QUrlList) ConstEnd0() {
	// QUrlList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QUrlList) Rbegin0() {
	// QUrlList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QUrlList) Rend0() {
	// QUrlList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QUrlList) Rbegin1() {
	// QUrlList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QUrlList) Rend1() {
	// QUrlList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QUrlList) Crbegin0() {
	// QUrlList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QUrlList) Crend0() {
	// QUrlList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QUrlList) Insert1() {
	// QUrlList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QUrlList) Erase0() {
	// QUrlList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QUrlList) Erase1() {
	// QUrlList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QUrlList) Count1() int {
	// QUrlList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QUrlList) Length0() int {
	// QUrlList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QUrlList) First0() *QUrl {
	// QUrlList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & constFirst()
func (this *QUrlList) ConstFirst0() *QUrl {
	// QUrlList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & first()
func (this *QUrlList) First1() *QUrl {
	// QUrlList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// T & last()
func (this *QUrlList) Last0() *QUrl {
	// QUrlList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & last()
func (this *QUrlList) Last1() *QUrl {
	// QUrlList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & constLast()
func (this *QUrlList) ConstLast0() *QUrl {
	// QUrlList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// void removeFirst()
func (this *QUrlList) RemoveFirst0() {
	// QUrlList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QUrlList) RemoveLast0() {
	// QUrlList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QUrlList) StartsWith0() bool {
	// QUrlList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QUrlList) EndsWith0() bool {
	// QUrlList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QUrlList) Mid0() *QUrlList {
	// QUrlList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QUrlList) Value0() *QUrl {
	// QUrlList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// T value(int, const T &)
func (this *QUrlList) Value1() *QUrl {
	// QUrlList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// void push_back(const T &)
func (this *QUrlList) Push_back0() {
	// QUrlList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QUrlList) Push_front0() {
	// QUrlList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QUrlList) Front0() *QUrl {
	// QUrlList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & front()
func (this *QUrlList) Front1() *QUrl {
	// QUrlList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// T & back()
func (this *QUrlList) Back0() *QUrl {
	// QUrlList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// const T & back()
func (this *QUrlList) Back1() *QUrl {
	// QUrlList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QUrl{}
}

// void pop_front()
func (this *QUrlList) Pop_front0() {
	// QUrlList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QUrlList) Pop_back0() {
	// QUrlList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QUrlList) Empty0() bool {
	// QUrlList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QUrlList) Operator_add_equal0() *QUrlList {
	// QUrlList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QUrlList) Operator_add0() *QUrlList {
	// QUrlList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QUrlList) Operator_add_equal1() *QUrlList {
	// QUrlList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QUrlList) Operator_left_shift0() *QUrlList {
	// QUrlList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QUrlList) Operator_left_shift1() *QUrlList {
	// QUrlList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QUrlList) ToVector0() {
	// QUrlList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QUrlList) ToSet0() {
	// QUrlList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QUrlList) FromVector0() *QUrlList {
	// QUrlList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QUrlList) FromSet0() *QUrlList {
	// QUrlList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QUrlList) FromStdList0() *QUrlList {
	// QUrlList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QUrlList) ToStdList0() {
	// QUrlList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QUrlList) Detach_helper_grow0() {
	// QUrlList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QUrlList) Detach_helper0() {
	// QUrlList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QUrlList) Detach_helper1() {
	// QUrlList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QUrlList) Dealloc0() {
	// QUrlList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QUrlList) Node_construct0() {
	// QUrlList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QUrlList) Node_destruct0() {
	// QUrlList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QUrlList) Node_copy0() {
	// QUrlList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QUrlList) Node_destruct1() {
	// QUrlList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QUrlList) IsValidIterator0() bool {
	// QUrlList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QUrlList) Op_eq_impl0() bool {
	// QUrlList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QUrlList) Op_eq_impl1() bool {
	// QUrlList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QUrlList) Contains_impl0() bool {
	// QUrlList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QUrlList) Contains_impl1() bool {
	// QUrlList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QUrlList) Count_impl0() int {
	// QUrlList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QUrlList) Count_impl1() int {
	// QUrlList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QUrlList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
