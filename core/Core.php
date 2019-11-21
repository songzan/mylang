<?php

define("FOF",chr(0));
define("EIN",chr(1));
define("OUT",chr(2));
define("QSTR",chr(3));
define("FAIL",chr(0));
define("PASS",chr(1));

function say($s) { echo "$s\n"; }

function error($s) { say($s); exit(); }

function itos($int) { return (string)$int; }

function first($s) { return $s[0]; }

function second($s) { return $s[1]; }

function tail($s) { return $s[-1]; }

function cut($s) { return substr($s,1,-1); }

function rest($s) { return substr($s,1); }

function now() { return time(); }

function pchop($s) {
  return substr($s,0,-1); }

function restat($s,$at) {
  return substr($s,$at); }

function toend($s) {
  if (strlen($s) < 20) return $s;
  return substr($s,0,20);
}

function isin($s, $arr) {
  return in_array($s, $arr, TRUE);
}

function split($s, $sub) {
  return explode($sub, $s);
}

function find($s, $sub) {
  if (strstr($s, $sub)) { return true; }
  return false;
}

function has(&$array,$key) {
  return isset($array[$key]);
}

function startwith($s,$findstr) {
  return (strpos($s,$findstr) === 0);
}

function restStrs(&$arr) {
  return array_slice($arr,1);
}

function tochars($s) { return str_split($s); }

function tostr($arr) { return implode('', $arr); }

function getline($s, $off) {
  $s = substr($s,0,$off);
  $count = 0;
  foreach (tochars($s) as $ch) {
    if ($ch == "\n") { $count++; }
  }
  return $count;
}

function read_file($file) {
  return file_get_contents($file);
}

function writefile($file,$data) {
  return file_put_contents($file,$data);
}

function isvspace($c) {
  if ($c == "\r") { return true; }
  if ($c == "\n") { return true; }
  return false;
}

function ishspace($c) {
  if ($c == ' ') { return true; }
  if ($c == "\t") { return true; }
  return false;
}

function iswords($c) {
  if (ctype_alnum($c)) { return true; }
  if ($c == '_') { return true; }
  if ($c == '-') { return true; }
  return false;
}

?>
