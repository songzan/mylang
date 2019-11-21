-- Lua implement of MySpp

local chr = string.char
local rep = string.rep
local find = string.find
local format = string.format
local insert = table.insert
local reverse = string.reverse
local itos = tostring
local lower = string.lower
local now = os.time
local ord = string.byte
local remove = table.remove
local rename = os.rename
local replace = string.gsub
local substr = string.sub
local toint = tonumber
local tostr = table.concat
local upper = string.upper

local FOF   = chr(0)
local EIN   = chr(1)
local OUT   = chr(2)
local QSTR  = chr(3)
local FAIL  = chr(0)
local PASS  = chr(1)

function error(x)
  print(x); os.exit()
end

function osargs()
  local osArgs = {}
  for i, v in pairs(arg) do
    if i > 0 then insert(osArgs, v) end
  end
  return osArgs
end

function readline()
  return io.read("*l")
end

function first(x)
  return substr(x,1,1)
end

function second(x)
  return substr(x,2,2)
end

function tochar(hexstr)
  return chr(tonumber("0x" .. hexstr))
end

function tail(x)
  return substr(x, #x)
end

function rest(x)
  return substr(x, 2)
end

function include(_str,_char)
  if find(_str,_char,1,true) == nil then
    return false
  end
  return true
end

function trim(_str)
  local _lstr = replace(_str, "^%s+", '')
  return replace(_lstr, "%s+$", '')
end

function charat(_str, _at)
  local off = _at + 1
  return substr(_str, off, off)
end

function arrayat(_arr, _at)
  local off = _at + 1
  return _arr[off]
end

function chars(_str)
  local _chars = {}
  for i=1, #_str do
    local _char = substr(_str,i,i)
    insert(_chars, _char)
  end
  return _chars
end

function split(_str,_sep)
  local _strs = {}
  local _chars = {}
  for _, _char in ipairs(chars(_str)) do
    if _char == _sep then
      insert(_strs, tostr(_chars))
      _chars = {}
    else
      insert(_chars, _char)
    end
  end
  insert(_strs, tostr(_chars))
  return _strs
end

function join(_table,_sep)
  return table.concat(_table, _sep)
end

function startwith(_str, _sub_str)
  local _len = #_sub_str
  if substr(_str, 0, _len) == _sub_str then
    return true
  end
  return false
end

function endwith(_str,_end)
  local r_str = reverse(_str)
  local r_end = reverse(_end)
  return startwith(r_str,r_end)
end

function map(_arr,_fn)
  local _map_arr = {}
  for _, _value in ipairs(_arr) do
    insert(_map_arr, _fn(_value))
  end
  return _map_arr
end

function restat(x,i)
  return substr(x,i+1)
end

function chop(s)
  return substr(s,1,-2)
end

function cut(s)
  return substr(s,2,-2)
end

function isin(_elem, _arr)
  for _, v in ipairs(_arr) do
    if v == _elem then return true end
  end
  return false
end

function restStrs(x)
  local rest_t = {}
  for i=2, #x do
    insert(rest_t, x[i])
  end
  return rest_t
end

function shift(_arr)
  return remove(_arr,1)
end

function has(_hash, _key)
  if _hash[_key] == nil then
    return false
  end
  return true
end

function toend(_str)
  if #_str < 50 then return _str end
  return substr(_str,0,50)
end

function getline(_str, _off)
  local _s = substr(_str,0,_off)
  local count = 0
  for _, ch in ipairs(chars(_s)) do
    if ch == "\n" then
      count = count + 1
    end
  end
  return itos(count)
end

function flat(_arr)
  return table.unpack(_arr)
end

function readfile(fname)
  local fh = assert(io.open(fname, 'r'))
  local my_text = fh:read('*all')
  fh:close()
  return my_text
end

function writefile(fname, text)
  local fh = assert(io.open(fname, 'w'))
  fh:write(text)
  fh:flush()
  fh:close()
end

function copy(from_file,to_file)
  writefile(to_file,readfile(from_file))
end

function isfile(_file)
  local f = io.open(_file,"r")
  if f == nil then return false end
  io.close(f)
  return true
end

function isupper(sc)
  if sc >= 'A' then
    if sc <= 'Z' then
      return true
    end
  end
  return false
end

function isvspace(sc)
  if sc == '\n' then
    return true
  end
  if sc == '\r' then
    return true
  end
  return false
end

function ishspace(sc)
  if sc == ' ' then
    return true
  end
  if sc == '\t' then
    return true
  end
  return false
end

function isspace(sc)
  if isvspace(sc) then
    return true
  end
  if ishspace(sc) then
    return true
  end
  return false
end

function isdigit(sc)
  if sc < '0' then
    return false
  end
  if sc > '9' then
    return false
  end
  return true
end

function islower(sc)
  if sc >= 'a' then
    if sc <= 'z' then
      return true
    end
  end
  return false
end

function isalpha(sc)
  if islower(sc) then
    return true
  end
  if isupper(sc) then
    return true
  end
  return false
end

function iswords(sc)
  if isalpha(sc) then
    return true
  end
  if isdigit(sc) then
    return true
  end
  if sc == '-' then return true end
  if sc == '_' then return true end
  return false
end

function isxdigit(sc)
  if isdigit(sc) then
    return true
  end
  if sc >= 'A' then
    if sc <= 'F' then
      return true
    end
  end
  if sc >= 'a' then
    if sc <= 'f' then
      return true
    end
  end
  return false
end

function isodigit(sc)
  if sc < '0' then
    return false
  end
  if sc > '7' then
    return false
  end
  return true
end

