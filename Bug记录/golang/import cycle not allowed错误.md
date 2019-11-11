# golang中("import cycle not allowed")错误

写代码遇见的这种错误import cycle not allowed，这是由于代码保重循环引用导致的，具体情况为

package A 引用了 package B 和package C ，而在package B中也引用了package C，导致包存在死循环。报错。

#### 解决方案

将package B 引用的package C 的内容独立出来，形成package D 即可

#### 

####  