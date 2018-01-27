package qtcore

// /usr/include/qt/QtCore/qfileinfo.h
// #include <qfileinfo.h>
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
import "gopp"
import "qt.go/cffiqt"
import "qt.go/qtrt"

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
}

//  ext block end

//  body block begin
type QFileInfoList struct {
	*qtrt.CObject
}

// QList<T> & operator=(const QList<T> &)
func (this *QFileInfoList) Operator_equal_0() *QFileInfoList {
	// QFileInfoList_operator_equal_0()
	return this
}

// QList<T> & operator=(QList<T> &&)
func (this *QFileInfoList) Operator_equal_1() *QFileInfoList {
	// QFileInfoList_operator_equal_1()
	return this
}

// void swap(QList<T> &)
func (this *QFileInfoList) Swap_0() {
	// QFileInfoList_swap_0()
}

// bool operator==(const QList<T> &)
func (this *QFileInfoList) Operator_equal_equal_0() bool {
	// QFileInfoList_operator_equal_equal_0()
	return 0 == 0
}

// bool operator!=(const QList<T> &)
func (this *QFileInfoList) Operator_not_equal_0() bool {
	// QFileInfoList_operator_not_equal_0()
	return 0 == 0
}

// int size()
func (this *QFileInfoList) Size_0() int {
	// QFileInfoList_size_0()
	return 0
}

// void detach()
func (this *QFileInfoList) Detach_0() {
	// QFileInfoList_detach_0()
}

// void detachShared()
func (this *QFileInfoList) DetachShared_0() {
	// QFileInfoList_detachShared_0()
}

// bool isDetached()
func (this *QFileInfoList) IsDetached_0() bool {
	// QFileInfoList_isDetached_0()
	return 0 == 0
}

// void setSharable(bool)
func (this *QFileInfoList) SetSharable_0() {
	// QFileInfoList_setSharable_0()
}

// bool isSharedWith(const QList<T> &)
func (this *QFileInfoList) IsSharedWith_0() bool {
	// QFileInfoList_isSharedWith_0()
	return 0 == 0
}

// bool isEmpty()
func (this *QFileInfoList) IsEmpty_0() bool {
	// QFileInfoList_isEmpty_0()
	return 0 == 0
}

// void clear()
func (this *QFileInfoList) Clear_0() {
	// QFileInfoList_clear_0()
}

// const T & at(int)
func (this *QFileInfoList) At_0() *QFileInfo {
	// QFileInfoList_at_0()
	return &QFileInfo{}
}

// const T & operator[](int)
func (this *QFileInfoList) Operator_get_index_0() *QFileInfo {
	// QFileInfoList_operator_get_index_0()
	return &QFileInfo{}
}

// T & operator[](int)
func (this *QFileInfoList) Operator_get_index_1() *QFileInfo {
	// QFileInfoList_operator_get_index_1()
	return &QFileInfo{}
}

// void reserve(int)
func (this *QFileInfoList) Reserve_0() {
	// QFileInfoList_reserve_0()
}

// void append(const T &)
func (this *QFileInfoList) Append_0() {
	// QFileInfoList_append_0()
}

// void append(const QList<T> &)
func (this *QFileInfoList) Append_1() {
	// QFileInfoList_append_1()
}

// void prepend(const T &)
func (this *QFileInfoList) Prepend_0() {
	// QFileInfoList_prepend_0()
}

// void insert(int, const T &)
func (this *QFileInfoList) Insert_0() {
	// QFileInfoList_insert_0()
}

// void replace(int, const T &)
func (this *QFileInfoList) Replace_0() {
	// QFileInfoList_replace_0()
}

// void removeAt(int)
func (this *QFileInfoList) RemoveAt_0() {
	// QFileInfoList_removeAt_0()
}

// int removeAll(const T &)
func (this *QFileInfoList) RemoveAll_0() int {
	// QFileInfoList_removeAll_0()
	return 0
}

// bool removeOne(const T &)
func (this *QFileInfoList) RemoveOne_0() bool {
	// QFileInfoList_removeOne_0()
	return 0 == 0
}

// T takeAt(int)
func (this *QFileInfoList) TakeAt_0() *QFileInfo {
	// QFileInfoList_takeAt_0()
	return &QFileInfo{}
}

// T takeFirst()
func (this *QFileInfoList) TakeFirst_0() *QFileInfo {
	// QFileInfoList_takeFirst_0()
	return &QFileInfo{}
}

// T takeLast()
func (this *QFileInfoList) TakeLast_0() *QFileInfo {
	// QFileInfoList_takeLast_0()
	return &QFileInfo{}
}

// void move(int, int)
func (this *QFileInfoList) Move_0() {
	// QFileInfoList_move_0()
}

// void swap(int, int)
func (this *QFileInfoList) Swap_1() {
	// QFileInfoList_swap_1()
}

// int indexOf(const T &, int)
func (this *QFileInfoList) IndexOf_0() int {
	// QFileInfoList_indexOf_0()
	return 0
}

// int lastIndexOf(const T &, int)
func (this *QFileInfoList) LastIndexOf_0() int {
	// QFileInfoList_lastIndexOf_0()
	return 0
}

// bool contains(const T &)
func (this *QFileInfoList) Contains_0() bool {
	// QFileInfoList_contains_0()
	return 0 == 0
}

// int count(const T &)
func (this *QFileInfoList) Count_0() int {
	// QFileInfoList_count_0()
	return 0
}

// QList::iterator begin()
func (this *QFileInfoList) Begin_0() {
	// QFileInfoList_begin_0()
}

// QList::const_iterator begin()
func (this *QFileInfoList) Begin_1() {
	// QFileInfoList_begin_1()
}

// QList::const_iterator cbegin()
func (this *QFileInfoList) Cbegin_0() {
	// QFileInfoList_cbegin_0()
}

// QList::const_iterator constBegin()
func (this *QFileInfoList) ConstBegin_0() {
	// QFileInfoList_constBegin_0()
}

// QList::iterator end()
func (this *QFileInfoList) End_0() {
	// QFileInfoList_end_0()
}

// QList::const_iterator end()
func (this *QFileInfoList) End_1() {
	// QFileInfoList_end_1()
}

// QList::const_iterator cend()
func (this *QFileInfoList) Cend_0() {
	// QFileInfoList_cend_0()
}

// QList::const_iterator constEnd()
func (this *QFileInfoList) ConstEnd_0() {
	// QFileInfoList_constEnd_0()
}

// QList::reverse_iterator rbegin()
func (this *QFileInfoList) Rbegin_0() {
	// QFileInfoList_rbegin_0()
}

// QList::reverse_iterator rend()
func (this *QFileInfoList) Rend_0() {
	// QFileInfoList_rend_0()
}

// QList::const_reverse_iterator rbegin()
func (this *QFileInfoList) Rbegin_1() {
	// QFileInfoList_rbegin_1()
}

// QList::const_reverse_iterator rend()
func (this *QFileInfoList) Rend_1() {
	// QFileInfoList_rend_1()
}

// QList::const_reverse_iterator crbegin()
func (this *QFileInfoList) Crbegin_0() {
	// QFileInfoList_crbegin_0()
}

// QList::const_reverse_iterator crend()
func (this *QFileInfoList) Crend_0() {
	// QFileInfoList_crend_0()
}

// QList::iterator insert(QList::iterator, const T &)
func (this *QFileInfoList) Insert_1() {
	// QFileInfoList_insert_1()
}

// QList::iterator erase(QList::iterator)
func (this *QFileInfoList) Erase_0() {
	// QFileInfoList_erase_0()
}

// QList::iterator erase(QList::iterator, QList::iterator)
func (this *QFileInfoList) Erase_1() {
	// QFileInfoList_erase_1()
}

// int count()
func (this *QFileInfoList) Count_1() int {
	// QFileInfoList_count_1()
	return 0
}

// int length()
func (this *QFileInfoList) Length_0() int {
	// QFileInfoList_length_0()
	return 0
}

// T & first()
func (this *QFileInfoList) First_0() *QFileInfo {
	// QFileInfoList_first_0()
	return &QFileInfo{}
}

// const T & constFirst()
func (this *QFileInfoList) ConstFirst_0() *QFileInfo {
	// QFileInfoList_constFirst_0()
	return &QFileInfo{}
}

// const T & first()
func (this *QFileInfoList) First_1() *QFileInfo {
	// QFileInfoList_first_1()
	return &QFileInfo{}
}

// T & last()
func (this *QFileInfoList) Last_0() *QFileInfo {
	// QFileInfoList_last_0()
	return &QFileInfo{}
}

// const T & last()
func (this *QFileInfoList) Last_1() *QFileInfo {
	// QFileInfoList_last_1()
	return &QFileInfo{}
}

// const T & constLast()
func (this *QFileInfoList) ConstLast_0() *QFileInfo {
	// QFileInfoList_constLast_0()
	return &QFileInfo{}
}

// void removeFirst()
func (this *QFileInfoList) RemoveFirst_0() {
	// QFileInfoList_removeFirst_0()
}

// void removeLast()
func (this *QFileInfoList) RemoveLast_0() {
	// QFileInfoList_removeLast_0()
}

// bool startsWith(const T &)
func (this *QFileInfoList) StartsWith_0() bool {
	// QFileInfoList_startsWith_0()
	return 0 == 0
}

// bool endsWith(const T &)
func (this *QFileInfoList) EndsWith_0() bool {
	// QFileInfoList_endsWith_0()
	return 0 == 0
}

// QList<T> mid(int, int)
func (this *QFileInfoList) Mid_0() *QFileInfoList {
	// QFileInfoList_mid_0()
	return this
}

// T value(int)
func (this *QFileInfoList) Value_0() *QFileInfo {
	// QFileInfoList_value_0()
	return &QFileInfo{}
}

// T value(int, const T &)
func (this *QFileInfoList) Value_1() *QFileInfo {
	// QFileInfoList_value_1()
	return &QFileInfo{}
}

// void push_back(const T &)
func (this *QFileInfoList) Push_back_0() {
	// QFileInfoList_push_back_0()
}

// void push_front(const T &)
func (this *QFileInfoList) Push_front_0() {
	// QFileInfoList_push_front_0()
}

// T & front()
func (this *QFileInfoList) Front_0() *QFileInfo {
	// QFileInfoList_front_0()
	return &QFileInfo{}
}

// const T & front()
func (this *QFileInfoList) Front_1() *QFileInfo {
	// QFileInfoList_front_1()
	return &QFileInfo{}
}

// T & back()
func (this *QFileInfoList) Back_0() *QFileInfo {
	// QFileInfoList_back_0()
	return &QFileInfo{}
}

// const T & back()
func (this *QFileInfoList) Back_1() *QFileInfo {
	// QFileInfoList_back_1()
	return &QFileInfo{}
}

// void pop_front()
func (this *QFileInfoList) Pop_front_0() {
	// QFileInfoList_pop_front_0()
}

// void pop_back()
func (this *QFileInfoList) Pop_back_0() {
	// QFileInfoList_pop_back_0()
}

// bool empty()
func (this *QFileInfoList) Empty_0() bool {
	// QFileInfoList_empty_0()
	return 0 == 0
}

// QList<T> & operator+=(const QList<T> &)
func (this *QFileInfoList) Operator_add_equal_0() *QFileInfoList {
	// QFileInfoList_operator_add_equal_0()
	return this
}

// QList<T> operator+(const QList<T> &)
func (this *QFileInfoList) Operator_add_0() *QFileInfoList {
	// QFileInfoList_operator_add_0()
	return this
}

// QList<T> & operator+=(const T &)
func (this *QFileInfoList) Operator_add_equal_1() *QFileInfoList {
	// QFileInfoList_operator_add_equal_1()
	return this
}

// QList<T> & operator<<(const T &)
func (this *QFileInfoList) Operator_left_shift_0() *QFileInfoList {
	// QFileInfoList_operator_left_shift_0()
	return this
}

// QList<T> & operator<<(const QList<T> &)
func (this *QFileInfoList) Operator_left_shift_1() *QFileInfoList {
	// QFileInfoList_operator_left_shift_1()
	return this
}

// QVector<T> toVector()
func (this *QFileInfoList) ToVector_0() {
	// QFileInfoList_toVector_0()
}

// QSet<T> toSet()
func (this *QFileInfoList) ToSet_0() {
	// QFileInfoList_toSet_0()
}

// QList<T> fromVector(const QVector<T> &)
func (this *QFileInfoList) FromVector_0() *QFileInfoList {
	// QFileInfoList_fromVector_0()
	return this
}

// QList<T> fromSet(const QSet<T> &)
func (this *QFileInfoList) FromSet_0() *QFileInfoList {
	// QFileInfoList_fromSet_0()
	return this
}

// QList<T> fromStdList(const std::list<T> &)
func (this *QFileInfoList) FromStdList_0() *QFileInfoList {
	// QFileInfoList_fromStdList_0()
	return this
}

// std::list<T> toStdList()
func (this *QFileInfoList) ToStdList_0() {
	// QFileInfoList_toStdList_0()
}

// QList::Node * detach_helper_grow(int, int)
func (this *QFileInfoList) Detach_helper_grow_0() {
	// QFileInfoList_detach_helper_grow_0()
}

// void detach_helper(int)
func (this *QFileInfoList) Detach_helper_0() {
	// QFileInfoList_detach_helper_0()
}

// void detach_helper()
func (this *QFileInfoList) Detach_helper_1() {
	// QFileInfoList_detach_helper_1()
}

// void dealloc(QListData::Data *)
func (this *QFileInfoList) Dealloc_0() {
	// QFileInfoList_dealloc_0()
}

// void node_construct(QList::Node *, const T &)
func (this *QFileInfoList) Node_construct_0() {
	// QFileInfoList_node_construct_0()
}

// void node_destruct(QList::Node *)
func (this *QFileInfoList) Node_destruct_0() {
	// QFileInfoList_node_destruct_0()
}

// void node_copy(QList::Node *, QList::Node *, QList::Node *)
func (this *QFileInfoList) Node_copy_0() {
	// QFileInfoList_node_copy_0()
}

// void node_destruct(QList::Node *, QList::Node *)
func (this *QFileInfoList) Node_destruct_1() {
	// QFileInfoList_node_destruct_1()
}

// bool isValidIterator(const QList::iterator &)
func (this *QFileInfoList) IsValidIterator_0() bool {
	// QFileInfoList_isValidIterator_0()
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::NotArrayCompatibleLayout)
func (this *QFileInfoList) Op_eq_impl_0() bool {
	// QFileInfoList_op_eq_impl_0()
	return 0 == 0
}

// bool op_eq_impl(const QList<T> &, QListData::ArrayCompatibleLayout)
func (this *QFileInfoList) Op_eq_impl_1() bool {
	// QFileInfoList_op_eq_impl_1()
	return 0 == 0
}

// bool contains_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QFileInfoList) Contains_impl_0() bool {
	// QFileInfoList_contains_impl_0()
	return 0 == 0
}

// bool contains_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QFileInfoList) Contains_impl_1() bool {
	// QFileInfoList_contains_impl_1()
	return 0 == 0
}

// int count_impl(const T &, QListData::NotArrayCompatibleLayout)
func (this *QFileInfoList) Count_impl_0() int {
	// QFileInfoList_count_impl_0()
	return 0
}

// int count_impl(const T &, QListData::ArrayCompatibleLayout)
func (this *QFileInfoList) Count_impl_1() int {
	// QFileInfoList_count_impl_1()
	return 0
}

//  body block end
