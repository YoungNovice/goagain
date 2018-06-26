###
* 类型hash表或者字典 以key-value形式存储数据
* key 必须是支持 == != 比较运算符的类型 不可以是函数 map 或者 slice
* map的查比线性搜索快很多 但是比使用索引访问数据的类型慢100倍
* map 使用make创建

make(m[keyType]valueType, cap) cap表示容量 可以省略
超出容量自动扩容
* len获取元素的个数

* 键值对不存在是自动添加 使用delete删除用键值对
* for range 对map和slice遍历操作


