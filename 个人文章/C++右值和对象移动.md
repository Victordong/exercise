**C++右值和对象移动**
#C++
## 什么是右值
c++的表达式要不然是左值 要不然是右值  

当一个对象被用作右值，用的是对象的值  
当一个对象备用作左值，用的是对象的在内存中的位置  

左值是可被操作（如取址，赋值）的“对象”，比如被定义好的变量  
右值是不可被操作（如取址，赋值）的“值”，比如字面常量，临时变量  

常见左值右值
```c++
std::string a = “string”;		// a为左值 “string”为右值
int b = 10;					// 10为右值 b为左值
int c = 10+b;					// 10+i为右值
b++							// 返回值为右值
++b							// 返回值为左值
```

函数返回的值可能是右值也可能是左值
下面这种返回的就是一个左值
```c++
int& leftValue(int& value) {
    return value;
}
```

而这种返回的是一个右值，因为返回的是一个拷贝值，在被使用完之后将会被删除
```c++
int rightValue(int& value) {
    return value;
}
```

**左值持久，右值短暂**
左值有持久的状态 右值是字面常量或求值过程中创建的临时变量，这种对象往往将要被销毁并且没有其他的用户在使用


## 右值引用
c++11 为支持移动操作引入了新的引用类型——右值引用
右值引用是绑定到右值上的引用，使用&&来获得右值引用
右值引用可以**延长对象的生命周期**

下面是常见的左值引用和右值引用
```c++
int i = 10;				
int& l1 = i;				// l1是左值引用
int&& r1 = 10;			// r1是右值引用
int&& r2 = i + 1;			// r2是右值引用
const int& l2 = i + 1;	// 可以将const 的左值引用绑定到一个右值上
int&&r3 = r1;				// 错误 因为变量本身是左值
int&& r3 = std::move(i);	// 使用move 生成右值
```

## 使用移动构造函数和移动赋值函数
```c++

class ValueClass {
   private:
    std::vector<std::string> strs;

   public:
    ValueClass(){};
	  ~ValueClass(){};
    ValueClass(std::vector<std::string>&& strs) : strs(strs){};
    ValueClass(ValueClass&& v) : strs(std::move(v.strs)){};
		// 移动构造函数
    ValueClass& operator=(ValueClass&& v) {
        if (this != &v) {
            strs = std::move(v.strs);
        }
        return *this;
    };
		// 移动赋值函数
    void print() {
        for (auto str : strs) {
            std::cout << str << “ “;
        }
    };
};

int main(int argc, char const* argv[]) {
    std::vector<std::string> a{“1”, “2”, “3”};
    ValueClass value_1(std::move(a));

    ValueClass value_2;
    value_2 = std::move(value_1);
    value_2.print();

    return 0;
}
```
运行结果
```shell
1 2 3 
```

**完成移动操作后 源对象应该处于销毁其是无害的状态**

### noexcept
有时我们需要使用noexcept来标明函数不会抛出异常，从而使标准库（类似vector），使用我们自己定义的移动函数

### 合成的移动操作
如果class已经定义了拷贝构造函数或拷贝赋值函数或析构函数，编译器就不会为其合成移动构造函数和移动赋值函数了 
当class中非static数据成员中有不可移动的，编译器也不会为其合成移动构造函数和移动赋值函数
在一定情况下编译器会将移动操作定义为delete

**对左值进行拷贝，对右值进行移动，如果没有移动操作，右值也将被拷贝**

同时支持左值和右值赋值
```c++
class ValueClass {
   public:
    ValueClass(const ValueClass& v) : strs(v.strs){};
		// 增加拷贝构造函数
    ValueClass& operator=(ValueClass v) {
        swap(*this, v);
        return *this;
    };
		// 更改赋值函数
    void swap(ValueClass& a, ValueClass& b) {
        using std::swap;
        swap(a.strs, b.strs);
    }
		// 增加swap函数
    void print() {
        for (auto str : strs) {
            std::cout << str << “ “;
        }
    };
    
};
```

## 移动的意义
移动语义
移动语义使得在 C++ 里返回大对象（如容器）的函数和运算符成为现实，因而可以提高代码的简洁性和可读性，提高程序员的生产率

## NROV
返回值优化：能把对象直接构造到调用者的栈上
C++11之后，返回值优化仍可以发生，但在没有返回值优化的情况下，编译器将试图把本地对象移动出去，而不是拷贝出去
