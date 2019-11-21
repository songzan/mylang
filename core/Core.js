// Javascript Implement of MySpp

"use strict"
var fs = require("fs");

const FOF  = chr(0);
const EIN  = chr(1);
const OUT  = chr(2);
const QSTR = chr(3);
const FAIL = chr(0);
const PASS = chr(1);

function now() { return Date.now(); }

function itos(n) { return n.toString() }

function say(s) { console.log(s) }

function exit() { process.exit(1) }

function error(s) { say(s); exit() }

function print(s) {
  process.stdout.write(s)
}

function startwith(s,b) {
  if (s[0] === b) return true;
  return false;
}

function endwith(s,b) {
  if (last(a) === b) return true;
  return false;
}

function tail(a) {
  return a[a.length-1] 
}

function cut(str) {
 return str.slice(1,-1) 
}

function chop(str) {
  return str.slice(0,-1) 
}

function hasStr(str,ch) {
  return str.includes(ch);
}

function split(str,p) {
  return str.split(p); 
}

function join(arr,ch) {
  return arr.join(ch); 
}

function tostr(arr) {
  return arr.join(''); 
}

function rest(str) {
  return str.substr(1); 
}

function restat(str,n) {
  return str.substr(n) 
}

function reststrs(a) {
  return a.slice(1); 
}

function tochars(s) {
  return s.split(''); 
}

function upper(s) {
  return s.toUpperCase(); 
}

function lower(s) {
  return s.toLowerCase(); 
}

function trim(s) { return s.trim() }

function repeat(a,b) { return a.repeat(b) }

function map(arr,fn) { return arr.map(fn) }

function chr(n) {
  return String.fromCharCode(n);
}

function osargs() {
  return process.argv.splice(2);
}

function hasKey(obj, key) {
  return obj.hasOwnProperty(key);
}

function getline(cs, off) {
  let s = cs.substr(0,off);
  let c = 0;
  for (let ch of s) {
    if (ch ==='\n') c++;
  }
  return c;
}

function toend(str) {
  if (str.length < 60) return str;
  return str.substr(0,60);
}

function readfile(file) {
  let data = fs.readFileSync(file);
  return data.toString();
}

function writefile(file,data) {
  fs.writeFileSync(file,data)
}

function isfile(file) {
  fs.exists(file, function(err) {
    if (err) { return true; }
    return false;
  })
}

function isupper(sc) {
  if (sc >= 'A') {
    if (sc <= 'Z') { return true }
  }
  return false
}

function isvspace(sc) {
  if (sc === '\n') {
    return true
  }
  if (sc === '\r') {
    return true
  }
  return false
}

function ishspace(sc) {
  if (sc === ' ') return true;
  if (sc === '\t') return true;
  return false
}

function isspace(sc) {
  if (isvspace(sc)) return true;
  if (ishspace(sc)) return true;
  return false;
}

function isdigit(sc) {
  if (sc < '0') return false;
  if (sc > '9') return false;
  return true
}

function islower(sc) {
  if (sc >= 'a') {
    if (sc <= 'z') {
      return true
    }
  }
  return false
}

function isalpha(sc) {
  if (islower(sc)) return true;
  if (isupper(sc)) return true;
  return false;
}

function iswords(sc) {
  if (isalpha(sc)) {
    return true
  }
  if (isdigit(sc)) {
    return true
  }
  if (sc === '_') {
    return true
  }
  if (sc === '-') {
    return true
  }
  return false
}

function isxdigit(sc) {
  if (isdigit(sc)) {
    return true
  }
  if (sc >= 'A') {
    if (sc <= 'F') {
      return true
    }
  }
  if (sc >= 'a') {
    if (sc <= 'f') {
      return true
    }
  }
  return false
}
