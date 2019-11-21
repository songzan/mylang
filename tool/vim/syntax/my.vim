" Vim synax file
" Language: Mylang
" Maintainer: Michael Song
" URL: http://www.github.com:songzan/myspp.git
" Started at: 2016-9-12
" Last change: 2018-10-19

if exists("b:current_syntax")
  finish
endif

let s:cpo_save = &cpo
set cpo&vim

setlocal iskeyword=33,36-38,42,43,45-46,60-90,94,97-122,126

"" Operator
syn keyword myOper <=
syn keyword myOper <
syn keyword myOper ==
syn keyword myOper >=
syn keyword myOper >
syn keyword myOper !=
syn keyword myOper +=
syn keyword myOper +
syn keyword myOper !~
syn keyword myOper >>
syn keyword myOper <<
syn keyword myOper eq
syn keyword myOper in

"" built-in Function
syn keyword myFunc basename
syn keyword myFunc chop
syn keyword myFunc cleandir
syn keyword myFunc copy
syn keyword myFunc cut
syn keyword myFunc dec
syn keyword myFunc dectochar
syn keyword myFunc dirfiles
syn keyword myFunc endwith
syn keyword myFunc find
syn keyword myFunc first
syn keyword myFunc getline
syn keyword myFunc has
syn keyword myFunc hextochar
syn keyword myFunc in
syn keyword myFunc inc
syn keyword myFunc index
syn keyword myFunc isalpha
syn keyword myFunc isdigit
syn keyword myFunc isfile
syn keyword myFunc ishspace
syn keyword myFunc isletter
syn keyword myFunc islower
syn keyword myFunc isspace
syn keyword myFunc isupper
syn keyword myFunc isvspace
syn keyword myFunc iswords
syn keyword myFunc join
syn keyword myFunc len
syn keyword myFunc lower
syn keyword myFunc map
syn keyword myFunc next
syn keyword myFunc not
syn keyword myFunc now
syn keyword myFunc osargs
syn keyword myFunc shift
syn keyword myFunc push
syn keyword myFunc range
syn keyword myFunc readfile
syn keyword myFunc readline
syn keyword myFunc region
syn keyword myFunc rename
syn keyword myFunc reset
syn keyword myFunc rest
syn keyword myFunc restat
syn keyword myFunc second
syn keyword myFunc shift
syn keyword myFunc split
syn keyword myFunc startwith
syn keyword myFunc tail
syn keyword myFunc tochar
syn keyword myFunc tochars
syn keyword myFunc toend
syn keyword myFunc toint
syn keyword myFunc tostr
syn keyword myFunc trim
syn keyword myFunc upper
syn keyword myFunc writefile

"" type declare
syn keyword myType nil
syn keyword myType int
syn keyword myType char
syn keyword myType chars
syn keyword myType buffer
syn keyword myType str
syn keyword myType bool
syn keyword myType fn
syn keyword myType strs
syn keyword myType table
syn keyword myType tree
syn keyword myType Cursor
syn keyword myType Lint

"" constant
syn keyword myConst EIN
syn keyword myConst FOF
syn keyword myConst FAIL
syn keyword myConst OUT
syn keyword myConst PASS
syn keyword myConst QSTR

"" Spp Boolean
syn keyword myBool true
syn keyword myBool false

"" Spp keyword
syn keyword myKeyword case
syn keyword myKeyword define
syn keyword myKeyword set
syn keyword myKeyword func
syn keyword myKeyword else
syn keyword myKeyword fn
syn keyword myKeyword for
syn keyword myKeyword if
syn keyword myKeyword my
syn keyword myKeyword ns
syn keyword myKeyword of
syn keyword myKeyword return
syn keyword myKeyword then
syn keyword myKeyword say
syn keyword myKeyword error
syn keyword myKeyword struct
syn keyword myKeyword use
syn keyword myKeyword while
syn keyword myKeyword ok

syn match   myComment  /#.*/
syn match   myComment  /\/\/.*/
syn region  MyBcomment start="/\*" end="\*/"
syn region  myLstr     start=/`/ end=/`/
syn match   myChar     "'[^\\]'"
syn match   myChar     "'[^']'"
syn match   myChar     "'\\['\\abefnrtv]'"
syn region  myStr      start=/"/ skip=/\\"/ end=/"/

hi def link myComment  Comment
hi def link myBcomment Comment
hi def link myType     Type
hi def link myOper     Operator
hi def link myFunc     Type
hi def link myBool     Constant
hi def link myLstr     String
hi def link myChar     String
hi def link myStr      String
hi def link myKeyword  Keyword
hi def link myConst    Keyword

let b:current_syntax = "my"

let &cpo = s:cpo_save
unlet s:cpo_save
