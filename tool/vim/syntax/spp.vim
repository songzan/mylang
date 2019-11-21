" Vim synax file
" Language: Spp
" Maintainer: Michael Song
" URL: http://www.gitee.com:str/myspp.git
" Started at: 2018-9-12
" Last change: 2018-10-31

if exists("b:current_syntax")
  finish
endif

let s:cpo_save = &cpo
set cpo&vim

setlocal iskeyword+=-

syn match sppComment  '//.*'
syn match sppName /[\l\u_]\w*/
syn match sppRept /[?+*]/
syn match sppNot /!/
syn match sppTill /\~/
syn match sppBranch /\|/
syn match sppChar /\\\S/
syn keyword sppEof EOF
syn region sppStr start=/'/ end=/'/
syn region sppCclass start=/\[/ skip=/\\\]/ end=/]/
syn region sppChars start=/{/ end=/}/ contains=@sppRule
syn region sppGroup start=/(/ end=/)/ contains=@sppRule
syn cluster sppRule contains=sppCclass,sppChar,SppStr,sppEof,sppTill,sppNot,sppRept
" Comment Constant Identifier
" Ignore PreProc
" Underlined
hi def link sppComment  Comment
hi def link sppName     Identifier
hi def link sppStr      Constant
hi def link sppRept     Special
hi def link sppNot      Special
hi def link sppTill     Special
hi def link sppBranch   Special
hi def link sppChar     Statement
hi def link sppEof      Statement
hi def link sppCclass  Special
hi def link sppChars    Type
hi def link sppGroup    Type

let b:current_syntax = "spp"

let &cpo = s:cpo_save
unlet s:cpo_save
