#!/usr/bin/env python

import shutil
import time
import sys
import os

FOF   = chr(0)
EIN   = chr(1)
OUT   = chr(2)
QSTR  = chr(3)
FAIL  = chr(0)
PASS  = chr(1)

def hextochar(s): return int(s,16)

def dectochar(s): return int(s,10)

def now():
  return time.strftime("%Y%m%d%H%M%S",time.localtime())

def upper(_str): return _str.upper()

def lower(_str): return _str.lower()

def say(_str):
    sys.stdout.write(_str + '\n')
    sys.stdout.flush()

def readline():
  return sys.stdin.readline()

def aprint(s):
  sys.stdout.write(s)
  sys.stdout.flush()

def error(s):
  print(s); exit()

def itos(_int):
  return str(_int)

def has(_hash, _key): return _key in _hash

def trim(_str): return _str.strip()

def tochars(_s): return list(_s)

def startwith(_str, _sub_str):
  return _str.startswith(_sub_str)

def endwith(_str, _sub_str):
  return _str.endswith(_sub_str)

def first(_str): return _str[0]
def tail(_str): return _str[-1]
def rest(_str): return _str[1:]
def second(_str): return _str[1]

def restat(_str,_at): return _str[_at:]

def getline(_str,_at):
  _s = restat(_str,_at)
  return _s.count("\n")

def toend(vs):
  if len(vs) < 20: return vs
  return vs[0:20]

def cut(_str): return _str[1:-1]

def chop(_str): return _str[:-1]

def find(_sub,_str):
  return _str.find(_sub) >= 0

def string(_array):
  return ''.join(_array)

def readfile(file):
  fh = open(file)
  file_str = fh.read()
  fh.close
  return file_str

def writefile(file,s):
  fh = open(file, 'w')
  fh.write(s)
  fh.close
  return True

def osargs(): return rest(sys.argv)

def get2(arr):
  if len(arr) > 2: return arr[0:2]
  return arr

def get3(arr):
  if len(arr) > 3: return arr[0:3]
  return arr

def isvspace(_c):
  if _c == '\n': return True
  if _c == '\r': return True
  return False

def ishspace(_c):
  if _c == ' ': return True
  if _c == '\t': return True
  return False

def iswords(_c):
  if _c.isalpha(): return True
  if _c.isdigit(): return True
  if _c == '_': return True
  if _c == '-': return True
  return False

def isxdigit(_c):
  if _c.isdigit(): return True
  if _c >= 'A':
    if _c <= 'F': return True
  if _c >= 'a':
    if _c <= 'f': return True
  return False  

def repeat(_s,_n):
    return _s * _n
