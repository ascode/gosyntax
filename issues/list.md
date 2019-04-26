# issues

### 为什么我的项目编译提示找不到项目自己这个模块？
问题描述：

我的项目名称叫做gosyntax，想在项目根目录执行go build，为什么却报错提示
```cassandraql
can't load package: package github.com/JINSCOP/gosyntax: unknown import path "github.com/JINSCOP/gosyntax": cannot find module providing package github.com/JINSCOP/gosyntax

```
我的项目是使用go mod进行包管理的。

解答：

经过尝试发现，在根目录下面如果没有go文件，那么编译的时候就会提示找不到项目自己这个模块，就是如上的问题。只要在根目录下面放一个go文件，哪怕是个只带有package声明的空的go文件，只要后缀名是go，就能够编译了。

---