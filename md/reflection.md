### 反射的作用

1. 大大提高程序的灵活性， 使得interface{}有更大的发挥余地
2. 反射使用TypeOf 和 ValueOf函数从接口中获取目标信息
3. 反射会将匿名字段作为独立字段(匿名字段的本质)
4. 想要利用反射修改对象的状态 前提interface.data是settable即pointer-interface
5. 通过反射可以的动态的调用方法