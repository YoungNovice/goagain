### slice
1. 本身不是数组 ， 它指向底层的数组
2. 作为变长数组的替代方案 可以关联底层数组的局部或者全部
3. 是引用类型
4. 可以直接创建或者从底层数组生成
5. len() 个数 cap() 容量
6. 用make创建

make([]t, len, cap) cap 可以省略 其长度就是len

### Reslice 
reslice是索引以被slice的切片为准
索引不可以超过被slice的切片的容量的cap值
索引越界将导致错误

### append
可以在slice尾部追加元素
可以将一个slice追加在另一个slice尾部
如果最终长度未超过追加到slice的容量则返回原始的slice
如果超过追加到slice的容量将重新分配数组并拷贝原始数据
