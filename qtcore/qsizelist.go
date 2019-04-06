package qtcore

// /usr/include/qt/QtCore/qsize.h
// #include <qsize.h>
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

func init_unused_10091() {
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
type QSizeList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QSizeList) Operator_equal0() *QSizeList {
	// QSizeList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QSizeList) Operator_equal1() *QSizeList {
	// QSizeList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QSizeList) Swap0() {
	// QSizeList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QSizeList) Operator_equal_equal0() bool {
	// QSizeList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QSizeList) Operator_not_equal0() bool {
	// QSizeList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QSizeList) Size0() int {
	// QSizeList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QSizeList) Detach0() {
	// QSizeList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QSizeList) DetachShared0() {
	// QSizeList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QSizeList) IsDetached0() bool {
	// QSizeList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QSizeList) SetSharable0() {
	// QSizeList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QSizeList) IsSharedWith0() bool {
	// QSizeList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QSizeList) IsEmpty0() bool {
	// QSizeList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QSizeList) Clear0() {
	// QSizeList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QSizeList) At0() *QSize {
	// QSizeList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & operator[](int)
func (this *QSizeList) Operator_get_index0() *QSize {
	// QSizeList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// T & operator[](int)
func (this *QSizeList) Operator_get_index1() *QSize {
	// QSizeList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// void reserve(int)
func (this *QSizeList) Reserve0() {
	// QSizeList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QSizeList) Append0() {
	// QSizeList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QSizeList) Append1() {
	// QSizeList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QSizeList) Prepend0() {
	// QSizeList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QSizeList) Insert0() {
	// QSizeList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QSizeList) Replace0() {
	// QSizeList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QSizeList) RemoveAt0() {
	// QSizeList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QSizeList) RemoveAll0() int {
	// QSizeList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QSizeList) RemoveOne0() bool {
	// QSizeList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QSizeList) TakeAt0() *QSize {
	// QSizeList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// T takeFirst()
func (this *QSizeList) TakeFirst0() *QSize {
	// QSizeList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// T takeLast()
func (this *QSizeList) TakeLast0() *QSize {
	// QSizeList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// void move(int, int)
func (this *QSizeList) Move0() {
	// QSizeList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QSizeList) Swap1() {
	// QSizeList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QSizeList) IndexOf0() int {
	// QSizeList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QSizeList) LastIndexOf0() int {
	// QSizeList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QSizeList) Contains0() bool {
	// QSizeList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QSizeList) Count0() int {
	// QSizeList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QSizeList) Begin0() {
	// QSizeList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QSizeList) Begin1() {
	// QSizeList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QSizeList) Cbegin0() {
	// QSizeList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QSizeList) ConstBegin0() {
	// QSizeList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QSizeList) End0() {
	// QSizeList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QSizeList) End1() {
	// QSizeList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QSizeList) Cend0() {
	// QSizeList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QSizeList) ConstEnd0() {
	// QSizeList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QSizeList) Rbegin0() {
	// QSizeList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QSizeList) Rend0() {
	// QSizeList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QSizeList) Rbegin1() {
	// QSizeList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QSizeList) Rend1() {
	// QSizeList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QSizeList) Crbegin0() {
	// QSizeList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QSizeList) Crend0() {
	// QSizeList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QSizeList) Insert1() {
	// QSizeList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QSizeList) Erase0() {
	// QSizeList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QSizeList) Erase1() {
	// QSizeList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QSizeList) Count1() int {
	// QSizeList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QSizeList) Length0() int {
	// QSizeList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QSizeList) First0() *QSize {
	// QSizeList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & constFirst()
func (this *QSizeList) ConstFirst0() *QSize {
	// QSizeList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & first()
func (this *QSizeList) First1() *QSize {
	// QSizeList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// T & last()
func (this *QSizeList) Last0() *QSize {
	// QSizeList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & last()
func (this *QSizeList) Last1() *QSize {
	// QSizeList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & constLast()
func (this *QSizeList) ConstLast0() *QSize {
	// QSizeList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// void removeFirst()
func (this *QSizeList) RemoveFirst0() {
	// QSizeList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QSizeList) RemoveLast0() {
	// QSizeList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QSizeList) StartsWith0() bool {
	// QSizeList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QSizeList) EndsWith0() bool {
	// QSizeList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QSizeList) Mid0() *QSizeList {
	// QSizeList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QSizeList) Value0() *QSize {
	// QSizeList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// T value(int, const T &)
func (this *QSizeList) Value1() *QSize {
	// QSizeList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// void push_back(const T &)
func (this *QSizeList) Push_back0() {
	// QSizeList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QSizeList) Push_front0() {
	// QSizeList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QSizeList) Front0() *QSize {
	// QSizeList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & front()
func (this *QSizeList) Front1() *QSize {
	// QSizeList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// T & back()
func (this *QSizeList) Back0() *QSize {
	// QSizeList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// const T & back()
func (this *QSizeList) Back1() *QSize {
	// QSizeList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QSize{}
}

// void pop_front()
func (this *QSizeList) Pop_front0() {
	// QSizeList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QSizeList) Pop_back0() {
	// QSizeList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QSizeList) Empty0() bool {
	// QSizeList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QSizeList) Operator_add_equal0() *QSizeList {
	// QSizeList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QSizeList) Operator_add0() *QSizeList {
	// QSizeList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QSizeList) Operator_add_equal1() *QSizeList {
	// QSizeList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QSizeList) Operator_left_shift0() *QSizeList {
	// QSizeList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QSizeList) Operator_left_shift1() *QSizeList {
	// QSizeList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QSizeList) ToVector0() {
	// QSizeList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QSizeList) ToSet0() {
	// QSizeList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QSizeList) FromVector0() *QSizeList {
	// QSizeList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QSizeList) FromSet0() *QSizeList {
	// QSizeList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QSizeList) FromStdList0() *QSizeList {
	// QSizeList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QSizeList) ToStdList0() {
	// QSizeList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QSizeList) Detach_helper_grow0() {
	// QSizeList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QSizeList) Detach_helper0() {
	// QSizeList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QSizeList) Detach_helper1() {
	// QSizeList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QSizeList) Dealloc0() {
	// QSizeList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QSizeList) Node_construct0() {
	// QSizeList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QSizeList) Node_destruct0() {
	// QSizeList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QSizeList) Node_copy0() {
	// QSizeList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QSizeList) Node_destruct1() {
	// QSizeList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QSizeList) IsValidIterator0() bool {
	// QSizeList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QSizeList) Op_eq_impl0() bool {
	// QSizeList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QSizeList) Op_eq_impl1() bool {
	// QSizeList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QSizeList) Contains_impl0() bool {
	// QSizeList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QSizeList) Contains_impl1() bool {
	// QSizeList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QSizeList) Count_impl0() int {
	// QSizeList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QSizeList) Count_impl1() int {
	// QSizeList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QSizeList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
