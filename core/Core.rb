#!/usr/bin/env ruby

FOF   = 0.chr
EIN   = 1.chr
OUT   = 2.chr
QSTR  = 3.chr
FAIL  = 0.chr
PASS  = 1.chr

def error(x)
  puts(x)
  exit
end

def readline()
  return STDIN.getc()
end

def now()
  time = Time.now()
  return time.strftime("%Y%m%d%H%M%S")
end

def itos(_int)
  return _int.to_s
end

def substr(_str,_from,_to)
  if _to > 0 then
    return _str[_from,_to]
  end
  _to = _to - 1
  return _str[_from .. _to]
end

def first(_str)
  return _str[0]
end

def second(_str)
  return _str[1]
end

def tail(_str)
  return _str[-1]
end

def rest(_str)
  return _str[1..-1]
end

def cut(_str)
  return _str[1..-2]
end

def chop(_str)
  return _str[0..-2]
end

def restat(_str,_at)
  return _str[_at..-1]
end

def toend(ss)
  if ss.length < 40 then
    return ss
  end
  return substr(ss,0,40)
end

def find(_str,_char)
  if _str.index(_char) == nil
    return false
  end
  return true
end

def repeat(_s,_int)
  return _s * _int 
end

def startwith(_str,_sub)
  return _str.start_with?(_sub)
end

def endwith(_str,_sub)
 return _str.end_with?(_sub)
end

def instrs(_elem, _array)
  return _array.include?(_elem)
end

def to_chars(_s)
  _chars = []
  _s.each_char {|c| _chars << c }
  return _chars
end

def split(_str,_sep)
  return _str.split(_sep) 
end

def to_string(x) return x.join end

def aflat(_array)
  return _array[0], _array[1] 
end

def join(_strs, _sep)
  return _strs.join(_sep) 
end

def has(_hash, _key)
  return _hash.has_key?(_key)
end

def readfile(_file)
  return IO.read(_file) 
end

def writefile(_file,_str)
  fh = File.new(_file, 'w')
  fh.puts(_str)
  fh.close
  return true
end

def getline(_str, _off)
  _s = restat(_str,_off)
  _count = 0
  for c in to_chars(_s) do
    if c == NLN then _count += 1 end
  end
  return _count
end

def osargs()
  return ARGV
end

def rename(_a,_b)
  File.rename(_a,_b)
end

def isfile(_f)
  if File::exists?(_f) then return true end
  return false
end

def isupper(_c)
  if _c >= 'A' then 
    if _c <= 'Z' then return true end
  end
  return false
end

def isvspace(_c)
  if _c == "\n" then return true end
  if _c == "\r" then return true end
  return false
end

def ishspace(_c)
  if _c == ' ' then return true end
  if _c == "\t" then return true end
  return false
end

def isspace(_c)
  if isvspace(_c) then return true end
  if ishspace(_c) then return true end
  return false
end

def isdigit(_c)
  if _c < '0' then return false end
  if _c > '9' then return false end
  return true
end

def islower(_c)
  if _c >= 'a' then 
    if _c <= 'z' then return true end
  end
  return false
end

def isalpha(_c)
  if islower(_c) then return true end
  if isupper(_c) then return true end
  return false
end

def iswords(_c)
  if isalpha(_c) then return true end
  if isdigit(_c) then return true end
  if _c == '_' then return true end
  if _c == '-' then return true end
  return false
end

def isxdigit(_c)
  if isdigit(_c) then return true end
  if _c >= 'A' then 
    if _c <= 'F' then return true end
  end
  if _c >= 'a' then 
    if _c <= 'f' then return true end
  end
  return false
end
