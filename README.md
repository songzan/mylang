## 为什么要设计自己的语言

为了追求语言的执行速度和更高的开发效率，不断有新的语言出现。

作为程序员，不得不学习多种语言，这样就可以在多种环境中书写代码。但各种语言千差万别，很多需要注意的地方，这些语言的个性化程度决定了程序员的开发效率。

这个项目的目的是帮助程序员设计自己的语言，隐藏底层语言的各种细节，专注于提高程序员的开发效率。

所以，设计自己的语言，并不需要同时设计底层运行模块，只需要设计前端语法描述，然后将语法树转换成目标代码，就完成了语言的设计目标。

拥有自己的语言，使用固定的风格编程，是一件快乐的事情。

这个项目本身就设计了一种计算机语言 MyLang，而且项目大部分的代码是用 MyLang 写的。

## MyLang 语言简介

没有返回函数类型的的语法，好像是关键字不够用了。用 sub 替代 func 做内置函数关键字。但现有的算法没有用到，等需要的时候再加。

函数定义没有使用大括号或是明确的结束标志。因为没有内置函数。

只是并没有专用解释器或编译器，只是把它转换成别的语言，以便完成代码的调试，扩展和编译。

已经实现了多种语言的转换模块：

  ToC ToGo ToJs ToLua ToPhp ToRuby ToPerl ToPython

也就是说，你可以用许多语言来测试这个项目。

## 为什么要用这么多语言来实现?

因为这个项目就是用来设计语法，解析语言，所以研究各种主流语言的语法是有益处的。另外，用不同的语言重写代码，会发现算法中的不通用之处。随着经历多种语言的重写，会发现以前引以为豪的对语言的细节的把握，实在是没有用的东西，从这些奇技淫巧中摆脱出来，才能真正写出稳定的代码。

## Spp 是什么意思？

Spp 是 String Prepare Parser 字符串预处理器的意思。

Spp 定义了一门类似 EBNF 的语法描述语言，而 Spp 解析器可以依据语法描述对文本进行解析，输出 JSON 数据结构。并不需要生成中间代码。

  > spp grammar.spp
  > This is grammar Repl, type 'exit' to exit

计算机语言可以根据程序员的习惯和爱好来定制，使用 Spp 语言用来描述语义。而代码只是一种特别的文本，可以被解析成数据结构，从而解析和计算，并生成其它的语言。

如果一个文本，遵循某种语法，就可以使用 Spp 进行解析, 通常用它来解析计算机代码。

## Spp 有什么用？

当代码可以被解析成数据结构的时候，代码本身就成了数据。可以用各种算法修改，计算，扩展，改变这个数据结构，然后再写成代码的时候，我们可以写能写代码的代码。

计算机语言有非常复杂的语法形式，所以说，能够解析代码的工具，也能解析大部分的文本。

当语言工具升级，项目代码因为不兼容，语言本身需要重新改写，绝大多数软件公司只能被动应对。

通过设计旧语言的语法，将代码升级到指定格式，能充分使用新语言工具的效能。因为大部分的代码不会变化，修改都是小改动，这就是代码迁移。

## Spp 已经做过什么项目？

Spp 项目本身的执行代码，就是用 Spp 工具解析, 从一种中间代码 MyLang 解析转换而来的，所以 Spp 项目本身就是自己的测试项目。

    # 以下命令会重新编译所有的中间代码
    > ./spp make

## 项目地址

可以使用 git 下载，或在线下载项目压缩包。

    > git clone https://www.gitee.com:str/myspp.git

Spp 默认的实现语言是 C.
  
    > cd myspp

    ## 显示 Spp 项目帮助：
    > ./c-spp.exe

    This is MySpp
    Copyright  2018-2020 Songzhiquan

    Usage:
    >> go run spp.go [help]
    >> go run spp.go make [go]
    >> go run spp.go spp [top]
    >> go run spp.go my [top]
    >> go run spp.go grammar [top]

    ## 进入 Spp 语法设计模式：
    $ ./c-spp.exe spp
    Spp <top> REPL, type 'exit' exit.
    >> a = b
    [[.rule[[.name.a][.rules[.name.b]]]]]
    [[.a[.ntoken.b]]]
    >>

    ## 调试语法文件模式
    > go run spp.spp grammar.spp

    ## 解析文本模式
    > go run spp.go grammar.spp text.file

使用 Spp 和 Antlr 不同，并不需要生成中间代码，会直接生成数据结构。

## Spp 的语法是一种上下文无关的语法

Spp 语言，通常叫 grammar，由一个或一个以上的规则 spec 组成。spec 规则包括规则名称 name 和一个或一个以上的规则 rule 组成:

    spec-one = rules-one ;
    spec-two = rules-two ;

规则名称是由字母组成的，大小写均可，可以用 _ 或 @ 开始。

    name = a ;
    _name = b ;
    Name = c ;
    @name = d ;

下面的语句是等价的！

    name = token ;
    name | token ;

名称和规则之间用 "|" 是为了多分支情况下，第一个分支的增加和变化不需要修改其它的字符。

试着比较下面的两个规则，如果删除或调整第一个分支，会有什么不同:

    rule
      = a
      | b
      | c
      ;

    rule
      | a
      | b
      | c
      ;

而单行规则用等号更直观:

    rule = a | b | c ;
    rule | a | b | c ;

Spp 的语法本身就是用自己定义的，所以用户可以定义自己的语法格式。

Spp 语言定义了一系列规则单元，用以描述文本的各种结构。

## 规则变量 token

因着引入变量，可以让语法能描述嵌套和递归的结构，规则变量是指向另外一个规则的引用.

    name = token ;
    token = other ;

有三种规则变量，命名规则变量 ntoken, 忽略规则变量 rtoken, 分组规则变量 gtoken.

.. 命名规则变量 ntoken (name token)

    ## 字母开头的名称是命名规则变量
    name Name

命名规则变量会对捕获的内容 match 进行命名:

    [name match at] # at 是捕获文本的位置

.. 忽略规则变量 rtoken (reject token)

    ## 下划线开头的名称是 rtoken
    _s = \s+ ;

如果匹配成功，不管匹配到的是什么，都只会忽略捕获的内容。这种结构常用来处理代码中的注释和空白。

.. 分组规则变量 gtoken (group token)
 
   ## 以 @ 开始的名称是分组规则变量
   @name = ab ?cd ;

如果没有匹配到 cd, 那么就返回 ab, 如果匹配到 cd, 就返回 @name.
这种规则通常作为临时规则，对其他组合规则进行命名。但不会影响输出的数据结构。

## string 字符串

   string = 'abc' ;

## range 区间

标记一段字符的范围是一种常见的字符规则定义方式。

    digit = [0-9];
    lower = [a-z];

## rules 规则串

    rules_define = rule1 rule2 rule3 ;
    rules_sample = rule1 rule2 rule3 ;

许多规则的串联，逐一匹配才算匹配，规则之间忽略空白。

## chars 分组
  
    group_sample = { rule1 rule2 rule3 } ;

chars 和 rules 相同，需要逐一匹配所有的 rule 才算匹配，不过 rule 之间不会自动忽略空白。这通常用于定义字符的串联，例如:

    Identifier = { [_\a] \w* }

Identifier 的第一个字符和连接的字符之间不允许有空格存在。

## branch 分支

  branch = case1 | case2 | case3

规则的分支，从上到下，只要匹配之中的任何一个，就算匹配。

## 匹配否定 not

  ! rule

如果匹配成功，则返回失败标志，如果不匹配，才算匹配成功。

## cclass 字符类

    :a   [ :l :u '_' ]
    :b   ' '
    :c   '`'
    :d   '0'-'9'
    :e   '\'
    :h   [ ' ' '\t' ]
    :l   'a'-'z' 
    :n   char \n
    :r   char \r
    :s   [ :b :n :r :t ]
    :t   char \t
    :u   'A'-'Z'
    v = [ :n :r :t ]
    w = [ - _ :a :d ]
    x = [\d a-f A-F];

## anychar 任意字符

    any-char = <ab-cd[]\>

anychar 中没有 range 和转义字符，所有的字符都按照本义，所以 anychar 不能包括 '>' 字符。

## 重复匹配 repeat match

.. more 要至少匹配 1 次，尽可能多

    more_sample = cd+ ;

.. many 不匹配也算匹配，尽可能多

    many_sample = rule* ;

.. maybe 可能

    maybe_sample = rule? ;

不匹配也算匹配，只能匹配一次

## keyword 关键字

通常用于声明语言的关键字，由于函数名称的前缀可能和关键字一样，为了能明确的区分它们，特增加关键字。

    return = :return 'return' ;


说是任意字符，其实它并不会匹配文本的末尾，只有 eof 会匹配。除此之外，它会匹配所以文本内的字符。

## eof 文件结束标记 用美元符号表示

    eof : EOF ;

eof 文件结束标记，在匹配文件的结尾。这样就表示整个文件都匹配完了。

避免回溯，提高文本解析速度

现在的正则表达式算法中有很多支持回溯的语法描述，这会严重降低算法的效率，也大大增加了算法的复杂度，更是让算法所耗用的缓存没有边界，这会成为程序的严重漏洞。其实设计良好的正则，是可以避免回溯的。

而回溯是语法设计错误的标志，就好像后悔一样，会耗费我们的生命。

使用文件句柄做输入，这样就能处理很大的文件，不需要很大的缓存就能很快速的解析！

文本的解析

计算机世界的信息是以文件的形式保存和传递的，而文件是以字符串的形式保存和展现的。

这个项目的作者，已经破产了，如果想要让这个项目继续下去，就加微信 134 34 425 524, 如果手机没停机，就会收到。

## end
