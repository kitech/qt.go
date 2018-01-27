package qtwidgets

// /usr/include/qt/QtWidgets/qwidget.h
// #include <qwidget.h>
// #include <QtWidgets>

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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"
import "qt.go/qtcore"
import "qt.go/qtgui"

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
		qtrt.KeepMe()
	}
	if false {
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
	if false {
		qtgui.KeepMe()
	}
}

//  ext block end

//  body block begin
type QWidgetList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QWidgetList) Operator_equal_0() *QWidgetList {
	// QWidgetList_operator_equal_0()
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QWidgetList) Operator_equal_1() *QWidgetList {
	// QWidgetList_operator_equal_1()
	return this
}

// void swap(QList<T> &)
func (this *QWidgetList) Swap_0() {
	// QWidgetList_swap_0()
}

// bool operator==(const QList<T> &)
func (this *QWidgetList) Operator_equal_equal_0() bool {
	// QWidgetList_operator_equal_equal_0()
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QWidgetList) Operator_not_equal_0() bool {
	// QWidgetList_operator_not_equal_0()
	return 0 == 0
}

// int size()
func (this *QWidgetList) Size_0() int {
	// QWidgetList_size_0()
	return 0
}

// void detach()
func (this *QWidgetList) Detach_0() {
	// QWidgetList_detach_0()
}

// void detachShared()
func (this *QWidgetList) DetachShared_0() {
	// QWidgetList_detachShared_0()
}

// bool isDetached()
func (this *QWidgetList) IsDetached_0() bool {
	// QWidgetList_isDetached_0()
	return 0 == 0
}

// void setSharable(bool)
func (this *QWidgetList) SetSharable_0() {
	// QWidgetList_setSharable_0()
}

// bool isSharedWith(const QList<T> &)
func (this *QWidgetList) IsSharedWith_0() bool {
	// QWidgetList_isSharedWith_0()
	return 0 == 0
}

// bool isEmpty()
func (this *QWidgetList) IsEmpty_0() bool {
	// QWidgetList_isEmpty_0()
	return 0 == 0
}

// void clear()
func (this *QWidgetList) Clear_0() {
	// QWidgetList_clear_0()
}

// const T & at(int)
func (this *QWidgetList) At_0() *QWidget {
	// QWidgetList_at_0()
	return &QWidget{}
}

// const T & operator[](int)
func (this *QWidgetList) Operator_get_index_0() *QWidget {
	// QWidgetList_operator_get_index_0()
	return &QWidget{}
}

// T & operator[](int)
func (this *QWidgetList) Operator_get_index_1() *QWidget {
	// QWidgetList_operator_get_index_1()
	return &QWidget{}
}

// void reserve(int)
func (this *QWidgetList) Reserve_0() {
	// QWidgetList_reserve_0()
}

// void append(const T &)
func (this *QWidgetList) Append_0() {
	// QWidgetList_append_0()
}

// void append(const QList<T> &)
func (this *QWidgetList) Append_1() {
	// QWidgetList_append_1()
}

// void prepend(const T &)
func (this *QWidgetList) Prepend_0() {
	// QWidgetList_prepend_0()
}

// void insert(int, const T &)
func (this *QWidgetList) Insert_0() {
	// QWidgetList_insert_0()
}

// void replace(int, const T &)
func (this *QWidgetList) Replace_0() {
	// QWidgetList_replace_0()
}

// void removeAt(int)
func (this *QWidgetList) RemoveAt_0() {
	// QWidgetList_removeAt_0()
}

// int removeAll(const T &)
func (this *QWidgetList) RemoveAll_0() int {
	// QWidgetList_removeAll_0()
	return 0
}

// bool removeOne(const T &)
func (this *QWidgetList) RemoveOne_0() bool {
	// QWidgetList_removeOne_0()
	return 0 == 0
}

// T takeAt(int)
func (this *QWidgetList) TakeAt_0() *QWidget {
	// QWidgetList_takeAt_0()
	return &QWidget{}
}

// T takeFirst()
func (this *QWidgetList) TakeFirst_0() *QWidget {
	// QWidgetList_takeFirst_0()
	return &QWidget{}
}

// T takeLast()
func (this *QWidgetList) TakeLast_0() *QWidget {
	// QWidgetList_takeLast_0()
	return &QWidget{}
}

// void move(int, int)
func (this *QWidgetList) Move_0() {
	// QWidgetList_move_0()
}

// void swap(int, int)
func (this *QWidgetList) Swap_1() {
	// QWidgetList_swap_1()
}

// int indexOf(const T &, int)
func (this *QWidgetList) IndexOf_0() int {
	// QWidgetList_indexOf_0()
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QWidgetList) LastIndexOf_0() int {
	// QWidgetList_lastIndexOf_0()
	return 0
}

// bool contains(const T &)
func (this *QWidgetList) Contains_0() bool {
	// QWidgetList_contains_0()
	return 0 == 0
}

// int count(const T &)
func (this *QWidgetList) Count_0() int {
	// QWidgetList_count_0()
	return 0
}

// QList::iterator begin()
func (this *QWidgetList) Begin_0() {
	// QWidgetList_begin_0()
}

// QList::const_iterator begin()
func (this *QWidgetList) Begin_1() {
	// QWidgetList_begin_1()
}

// QList::const_iterator cbegin()
func (this *QWidgetList) Cbegin_0() {
	// QWidgetList_cbegin_0()
}

// QList::const_iterator constBegin()
func (this *QWidgetList) ConstBegin_0() {
	// QWidgetList_constBegin_0()
}

// QList::iterator end()
func (this *QWidgetList) End_0() {
	// QWidgetList_end_0()
}

// QList::const_iterator end()
func (this *QWidgetList) End_1() {
	// QWidgetList_end_1()
}

// QList::const_iterator cend()
func (this *QWidgetList) Cend_0() {
	// QWidgetList_cend_0()
}

// QList::const_iterator constEnd()
func (this *QWidgetList) ConstEnd_0() {
	// QWidgetList_constEnd_0()
}

// QList::reverse_iterator rbegin()
func (this *QWidgetList) Rbegin_0() {
	// QWidgetList_rbegin_0()
}

// QList::reverse_iterator rend()
func (this *QWidgetList) Rend_0() {
	// QWidgetList_rend_0()
}

// QList::const_reverse_iterator rbegin()
func (this *QWidgetList) Rbegin_1() {
	// QWidgetList_rbegin_1()
}

// QList::const_reverse_iterator rend()
func (this *QWidgetList) Rend_1() {
	// QWidgetList_rend_1()
}

// QList::const_reverse_iterator crbegin()
func (this *QWidgetList) Crbegin_0() {
	// QWidgetList_crbegin_0()
}

// QList::const_reverse_iterator crend()
func (this *QWidgetList) Crend_0() {
	// QWidgetList_crend_0()
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QWidgetList) Insert_1() {
	// QWidgetList_insert_1()
}

// QList::iterator erase(QList::iterator)
func (this *QWidgetList) Erase_0() {
	// QWidgetList_erase_0()
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QWidgetList) Erase_1() {
	// QWidgetList_erase_1()
}

// int count()
func (this *QWidgetList) Count_1() int {
	// QWidgetList_count_1()
	return 0
}

// int length()
func (this *QWidgetList) Length_0() int {
	// QWidgetList_length_0()
	return 0
}

// T & first()
func (this *QWidgetList) First_0() *QWidget {
	// QWidgetList_first_0()
	return &QWidget{}
}

// const T & constFirst()
func (this *QWidgetList) ConstFirst_0() *QWidget {
	// QWidgetList_constFirst_0()
	return &QWidget{}
}

// const T & first()
func (this *QWidgetList) First_1() *QWidget {
	// QWidgetList_first_1()
	return &QWidget{}
}

// T & last()
func (this *QWidgetList) Last_0() *QWidget {
	// QWidgetList_last_0()
	return &QWidget{}
}

// const T & last()
func (this *QWidgetList) Last_1() *QWidget {
	// QWidgetList_last_1()
	return &QWidget{}
}

// const T & constLast()
func (this *QWidgetList) ConstLast_0() *QWidget {
	// QWidgetList_constLast_0()
	return &QWidget{}
}

// void removeFirst()
func (this *QWidgetList) RemoveFirst_0() {
	// QWidgetList_removeFirst_0()
}

// void removeLast()
func (this *QWidgetList) RemoveLast_0() {
	// QWidgetList_removeLast_0()
}

// bool startsWith(const T &)
func (this *QWidgetList) StartsWith_0() bool {
	// QWidgetList_startsWith_0()
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QWidgetList) EndsWith_0() bool {
	// QWidgetList_endsWith_0()
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QWidgetList) Mid_0() *QWidgetList {
	// QWidgetList_mid_0()
	return this
}

// T value(int)
func (this *QWidgetList) Value_0() *QWidget {
	// QWidgetList_value_0()
	return &QWidget{}
}

// T value(int, const T &)
func (this *QWidgetList) Value_1() *QWidget {
	// QWidgetList_value_1()
	return &QWidget{}
}

// void push_back(const T &)
func (this *QWidgetList) Push_back_0() {
	// QWidgetList_push_back_0()
}

// void push_front(const T &)
func (this *QWidgetList) Push_front_0() {
	// QWidgetList_push_front_0()
}

// T & front()
func (this *QWidgetList) Front_0() *QWidget {
	// QWidgetList_front_0()
	return &QWidget{}
}

// const T & front()
func (this *QWidgetList) Front_1() *QWidget {
	// QWidgetList_front_1()
	return &QWidget{}
}

// T & back()
func (this *QWidgetList) Back_0() *QWidget {
	// QWidgetList_back_0()
	return &QWidget{}
}

// const T & back()
func (this *QWidgetList) Back_1() *QWidget {
	// QWidgetList_back_1()
	return &QWidget{}
}

// void pop_front()
func (this *QWidgetList) Pop_front_0() {
	// QWidgetList_pop_front_0()
}

// void pop_back()
func (this *QWidgetList) Pop_back_0() {
	// QWidgetList_pop_back_0()
}

// bool empty()
func (this *QWidgetList) Empty_0() bool {
	// QWidgetList_empty_0()
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QWidgetList) Operator_add_equal_0() *QWidgetList {
	// QWidgetList_operator_add_equal_0()
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QWidgetList) Operator_add_0() *QWidgetList {
	// QWidgetList_operator_add_0()
	return this
}

// QList<T> & operator+=(const T &)
func (this *QWidgetList) Operator_add_equal_1() *QWidgetList {
	// QWidgetList_operator_add_equal_1()
	return this
}

// QList<T> & operator<<(const T &)
func (this *QWidgetList) Operator_left_shift_0() *QWidgetList {
	// QWidgetList_operator_left_shift_0()
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QWidgetList) Operator_left_shift_1() *QWidgetList {
	// QWidgetList_operator_left_shift_1()
	return this
}

// QVector<T> toVector()
func (this *QWidgetList) ToVector_0() {
	// QWidgetList_toVector_0()
}

// QSet<T> toSet()
func (this *QWidgetList) ToSet_0() {
	// QWidgetList_toSet_0()
}

// QList<T> fromVector(const QVector<T> &)
func (this *QWidgetList) FromVector_0() *QWidgetList {
	// QWidgetList_fromVector_0()
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QWidgetList) FromSet_0() *QWidgetList {
	// QWidgetList_fromSet_0()
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QWidgetList) FromStdList_0() *QWidgetList {
	// QWidgetList_fromStdList_0()
	return this
}

// std::list<T> toStdList()
func (this *QWidgetList) ToStdList_0() {
	// QWidgetList_toStdList_0()
}

// QList::Node * detach_helper_grow(int, int)
func (this *QWidgetList) Detach_helper_grow_0() {
	// QWidgetList_detach_helper_grow_0()
}

// void detach_helper(int)
func (this *QWidgetList) Detach_helper_0() {
	// QWidgetList_detach_helper_0()
}

// void detach_helper()
func (this *QWidgetList) Detach_helper_1() {
	// QWidgetList_detach_helper_1()
}

// void dealloc(QListData::Data *)
func (this *QWidgetList) Dealloc_0() {
	// QWidgetList_dealloc_0()
}

// void node_construct(QList::Node *, const T &)
func (this *QWidgetList) Node_construct_0() {
	// QWidgetList_node_construct_0()
}

// void node_destruct(QList::Node *)
func (this *QWidgetList) Node_destruct_0() {
	// QWidgetList_node_destruct_0()
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QWidgetList) Node_copy_0() {
	// QWidgetList_node_copy_0()
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QWidgetList) Node_destruct_1() {
	// QWidgetList_node_destruct_1()
}

// bool isValidIterator(const QList::iterator &)
func (this *QWidgetList) IsValidIterator_0() bool {
	// QWidgetList_isValidIterator_0()
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QWidgetList) Op_eq_impl_0() bool {
	// QWidgetList_op_eq_impl_0()
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QWidgetList) Op_eq_impl_1() bool {
	// QWidgetList_op_eq_impl_1()
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QWidgetList) Contains_impl_0() bool {
	// QWidgetList_contains_impl_0()
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QWidgetList) Contains_impl_1() bool {
	// QWidgetList_contains_impl_1()
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QWidgetList) Count_impl_0() int {
	// QWidgetList_count_impl_0()
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QWidgetList) Count_impl_1() int {
	// QWidgetList_count_impl_1()
	return 0
}

//  body block end
