package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
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
type QTextBlockList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QTextBlockList) Operator_equal0() *QTextBlockList {
	// QTextBlockList_operator_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QTextBlockList) Operator_equal1() *QTextBlockList {
	// QTextBlockList_operator_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// void swap(QList<T> &)
func (this *QTextBlockList) Swap0() {
	// QTextBlockList_swap_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_swap_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool operator==(const QList<T> &)
func (this *QTextBlockList) Operator_equal_equal0() bool {
	// QTextBlockList_operator_equal_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_equal_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QTextBlockList) Operator_not_equal0() bool {
	// QTextBlockList_operator_not_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_not_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int size()
func (this *QTextBlockList) Size0() int {
	// QTextBlockList_size_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_size_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// void detach()
func (this *QTextBlockList) Detach0() {
	// QTextBlockList_detach_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_detach_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detachShared()
func (this *QTextBlockList) DetachShared0() {
	// QTextBlockList_detachShared_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_detachShared_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isDetached()
func (this *QTextBlockList) IsDetached0() bool {
	// QTextBlockList_isDetached_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_isDetached_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void setSharable(bool)
func (this *QTextBlockList) SetSharable0() {
	// QTextBlockList_setSharable_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_setSharable_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isSharedWith(const QList<T> &)
func (this *QTextBlockList) IsSharedWith0() bool {
	// QTextBlockList_isSharedWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_isSharedWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool isEmpty()
func (this *QTextBlockList) IsEmpty0() bool {
	// QTextBlockList_isEmpty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_isEmpty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// void clear()
func (this *QTextBlockList) Clear0() {
	// QTextBlockList_clear_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_clear_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// const T & at(int)
func (this *QTextBlockList) At0() *QTextBlock {
	// QTextBlockList_at_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_at_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & operator[](int)
func (this *QTextBlockList) Operator_get_index0() *QTextBlock {
	// QTextBlockList_operator_get_index_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_get_index_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// T & operator[](int)
func (this *QTextBlockList) Operator_get_index1() *QTextBlock {
	// QTextBlockList_operator_get_index_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_get_index_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// void reserve(int)
func (this *QTextBlockList) Reserve0() {
	// QTextBlockList_reserve_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_reserve_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const T &)
func (this *QTextBlockList) Append0() {
	// QTextBlockList_append_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_append_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void append(const QList<T> &)
func (this *QTextBlockList) Append1() {
	// QTextBlockList_append_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_append_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void prepend(const T &)
func (this *QTextBlockList) Prepend0() {
	// QTextBlockList_prepend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_prepend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void insert(int, const T &)
func (this *QTextBlockList) Insert0() {
	// QTextBlockList_insert_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_insert_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void replace(int, const T &)
func (this *QTextBlockList) Replace0() {
	// QTextBlockList_replace_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_replace_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeAt(int)
func (this *QTextBlockList) RemoveAt0() {
	// QTextBlockList_removeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_removeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int removeAll(const T &)
func (this *QTextBlockList) RemoveAll0() int {
	// QTextBlockList_removeAll_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_removeAll_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool removeOne(const T &)
func (this *QTextBlockList) RemoveOne0() bool {
	// QTextBlockList_removeOne_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_removeOne_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// T takeAt(int)
func (this *QTextBlockList) TakeAt0() *QTextBlock {
	// QTextBlockList_takeAt_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_takeAt_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// T takeFirst()
func (this *QTextBlockList) TakeFirst0() *QTextBlock {
	// QTextBlockList_takeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_takeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// T takeLast()
func (this *QTextBlockList) TakeLast0() *QTextBlock {
	// QTextBlockList_takeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_takeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// void move(int, int)
func (this *QTextBlockList) Move0() {
	// QTextBlockList_move_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_move_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void swap(int, int)
func (this *QTextBlockList) Swap1() {
	// QTextBlockList_swap_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_swap_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int indexOf(const T &, int)
func (this *QTextBlockList) IndexOf0() int {
	// QTextBlockList_indexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_indexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QTextBlockList) LastIndexOf0() int {
	// QTextBlockList_lastIndexOf_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_lastIndexOf_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// bool contains(const T &)
func (this *QTextBlockList) Contains0() bool {
	// QTextBlockList_contains_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_contains_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count(const T &)
func (this *QTextBlockList) Count0() int {
	// QTextBlockList_count_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_count_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// QList::iterator begin()
func (this *QTextBlockList) Begin0() {
	// QTextBlockList_begin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_begin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator begin()
func (this *QTextBlockList) Begin1() {
	// QTextBlockList_begin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_begin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cbegin()
func (this *QTextBlockList) Cbegin0() {
	// QTextBlockList_cbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_cbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constBegin()
func (this *QTextBlockList) ConstBegin0() {
	// QTextBlockList_constBegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_constBegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator end()
func (this *QTextBlockList) End0() {
	// QTextBlockList_end_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_end_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator end()
func (this *QTextBlockList) End1() {
	// QTextBlockList_end_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_end_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator cend()
func (this *QTextBlockList) Cend0() {
	// QTextBlockList_cend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_cend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_iterator constEnd()
func (this *QTextBlockList) ConstEnd0() {
	// QTextBlockList_constEnd_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_constEnd_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rbegin()
func (this *QTextBlockList) Rbegin0() {
	// QTextBlockList_rbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_rbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::reverse_iterator rend()
func (this *QTextBlockList) Rend0() {
	// QTextBlockList_rend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_rend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rbegin()
func (this *QTextBlockList) Rbegin1() {
	// QTextBlockList_rbegin_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_rbegin_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator rend()
func (this *QTextBlockList) Rend1() {
	// QTextBlockList_rend_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_rend_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crbegin()
func (this *QTextBlockList) Crbegin0() {
	// QTextBlockList_crbegin_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_crbegin_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::const_reverse_iterator crend()
func (this *QTextBlockList) Crend0() {
	// QTextBlockList_crend_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_crend_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QTextBlockList) Insert1() {
	// QTextBlockList_insert_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_insert_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator)
func (this *QTextBlockList) Erase0() {
	// QTextBlockList_erase_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_erase_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QTextBlockList) Erase1() {
	// QTextBlockList_erase_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_erase_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// int count()
func (this *QTextBlockList) Count1() int {
	// QTextBlockList_count_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_count_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int length()
func (this *QTextBlockList) Length0() int {
	// QTextBlockList_length_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_length_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// T & first()
func (this *QTextBlockList) First0() *QTextBlock {
	// QTextBlockList_first_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_first_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & constFirst()
func (this *QTextBlockList) ConstFirst0() *QTextBlock {
	// QTextBlockList_constFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_constFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & first()
func (this *QTextBlockList) First1() *QTextBlock {
	// QTextBlockList_first_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_first_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// T & last()
func (this *QTextBlockList) Last0() *QTextBlock {
	// QTextBlockList_last_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_last_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & last()
func (this *QTextBlockList) Last1() *QTextBlock {
	// QTextBlockList_last_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_last_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & constLast()
func (this *QTextBlockList) ConstLast0() *QTextBlock {
	// QTextBlockList_constLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_constLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// void removeFirst()
func (this *QTextBlockList) RemoveFirst0() {
	// QTextBlockList_removeFirst_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_removeFirst_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void removeLast()
func (this *QTextBlockList) RemoveLast0() {
	// QTextBlockList_removeLast_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_removeLast_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool startsWith(const T &)
func (this *QTextBlockList) StartsWith0() bool {
	// QTextBlockList_startsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_startsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QTextBlockList) EndsWith0() bool {
	// QTextBlockList_endsWith_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_endsWith_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QTextBlockList) Mid0() *QTextBlockList {
	// QTextBlockList_mid_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_mid_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// T value(int)
func (this *QTextBlockList) Value0() *QTextBlock {
	// QTextBlockList_value_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_value_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// T value(int, const T &)
func (this *QTextBlockList) Value1() *QTextBlock {
	// QTextBlockList_value_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_value_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// void push_back(const T &)
func (this *QTextBlockList) Push_back0() {
	// QTextBlockList_push_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_push_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void push_front(const T &)
func (this *QTextBlockList) Push_front0() {
	// QTextBlockList_push_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_push_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// T & front()
func (this *QTextBlockList) Front0() *QTextBlock {
	// QTextBlockList_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & front()
func (this *QTextBlockList) Front1() *QTextBlock {
	// QTextBlockList_front_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_front_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// T & back()
func (this *QTextBlockList) Back0() *QTextBlock {
	// QTextBlockList_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// const T & back()
func (this *QTextBlockList) Back1() *QTextBlock {
	// QTextBlockList_back_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_back_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return &QTextBlock{}
}

// void pop_front()
func (this *QTextBlockList) Pop_front0() {
	// QTextBlockList_pop_front_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_pop_front_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void pop_back()
func (this *QTextBlockList) Pop_back0() {
	// QTextBlockList_pop_back_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_pop_back_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool empty()
func (this *QTextBlockList) Empty0() bool {
	// QTextBlockList_empty_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_empty_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QTextBlockList) Operator_add_equal0() *QTextBlockList {
	// QTextBlockList_operator_add_equal_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_add_equal_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QTextBlockList) Operator_add0() *QTextBlockList {
	// QTextBlockList_operator_add_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_add_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator+=(const T &)
func (this *QTextBlockList) Operator_add_equal1() *QTextBlockList {
	// QTextBlockList_operator_add_equal_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_add_equal_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const T &)
func (this *QTextBlockList) Operator_left_shift0() *QTextBlockList {
	// QTextBlockList_operator_left_shift_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_left_shift_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QTextBlockList) Operator_left_shift1() *QTextBlockList {
	// QTextBlockList_operator_left_shift_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_operator_left_shift_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QVector<T> toVector()
func (this *QTextBlockList) ToVector0() {
	// QTextBlockList_toVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_toVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QSet<T> toSet()
func (this *QTextBlockList) ToSet0() {
	// QTextBlockList_toSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_toSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList<T> fromVector(const QVector<T> &)
func (this *QTextBlockList) FromVector0() *QTextBlockList {
	// QTextBlockList_fromVector_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_fromVector_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QTextBlockList) FromSet0() *QTextBlockList {
	// QTextBlockList_fromSet_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_fromSet_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QTextBlockList) FromStdList0() *QTextBlockList {
	// QTextBlockList_fromStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_fromStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return this
}

// std::list<T> toStdList()
func (this *QTextBlockList) ToStdList0() {
	// QTextBlockList_toStdList_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_toStdList_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// QList::Node * detach_helper_grow(int, int)
func (this *QTextBlockList) Detach_helper_grow0() {
	// QTextBlockList_detach_helper_grow_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_detach_helper_grow_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper(int)
func (this *QTextBlockList) Detach_helper0() {
	// QTextBlockList_detach_helper_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_detach_helper_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void detach_helper()
func (this *QTextBlockList) Detach_helper1() {
	// QTextBlockList_detach_helper_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_detach_helper_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void dealloc(QListData::Data *)
func (this *QTextBlockList) Dealloc0() {
	// QTextBlockList_dealloc_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_dealloc_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_construct(QList::Node *, const T &)
func (this *QTextBlockList) Node_construct0() {
	// QTextBlockList_node_construct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_node_construct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *)
func (this *QTextBlockList) Node_destruct0() {
	// QTextBlockList_node_destruct_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_node_destruct_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QTextBlockList) Node_copy0() {
	// QTextBlockList_node_copy_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_node_copy_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QTextBlockList) Node_destruct1() {
	// QTextBlockList_node_destruct_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_node_destruct_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
}

// bool isValidIterator(const QList::iterator &)
func (this *QTextBlockList) IsValidIterator0() bool {
	// QTextBlockList_isValidIterator_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_isValidIterator_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QTextBlockList) Op_eq_impl0() bool {
	// QTextBlockList_op_eq_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_op_eq_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QTextBlockList) Op_eq_impl1() bool {
	// QTextBlockList_op_eq_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_op_eq_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QTextBlockList) Contains_impl0() bool {
	// QTextBlockList_contains_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_contains_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QTextBlockList) Contains_impl1() bool {
	// QTextBlockList_contains_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_contains_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QTextBlockList) Count_impl0() int {
	// QTextBlockList_count_impl_0()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_count_impl_0", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QTextBlockList) Count_impl1() int {
	// QTextBlockList_count_impl_1()
	rv, err := qtrt.InvokeQtFunc6("C_QTextBlockList_count_impl_1", qtrt.FFI_TYPE_POINTER, this.Cthis)
	qtrt.ErrPrint(err, rv)
	return 0
}

//  body block end
