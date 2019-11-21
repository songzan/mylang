package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "bufio"
  "strconv"
  "strings"
  "time"
  "bytes"
  "unicode"
)

type char = rune
type Buffer = bytes.Buffer
type table = map[string]string
type tree  = map[string]table

type Cursor struct {
  str    string
  input  []char
  len    int
  rtable table
  output []string
  depth  int
  at     int
}

type Lint struct {
  counter int
  indent  int
  at      string
  ns      string
  ret     string
  stack   []string
  stree   tree
}

const (
  FOF   char = 0
  EIN   char = 1
  OUT   char = 2
  QSTR  char = 3
  FAIL  = string(0)
  PASS  = string(1)
)

func osargs() []string { return os.Args[1:] }

func instrs(s string, ss []string) bool {
  for _, elem := range ss {
    if elem == s { return true }
  }
  return false
}

func hextostr(hex string) string {
  i, err := strconv.ParseInt(hex, 16, 32)
  if err != nil { panic(err) }
  return string(i)
}

func now() string {
  return time.Now().Format("20060102150405")
}

func error(s string) {
  println(s); os.Exit(1)
}

func readline() string {
  reader := bufio.NewReader(os.Stdin)
  line, _ := reader.ReadString('\n')
  return line
}

func first(s string) char { return []char(s)[0] }

func second(s string) char { return []char(s)[1] }

func rest(s string) string {
  return string([]char(s)[1:])
}

func tail(s string) char {
  cs := []char(s)
  return cs[len(cs)-1]
}

func cut(s string) string {
  cs := []char(s)
  return string(cs[1:len(cs)-1])
}

func chop(s string) string {
  cs := []char(s)
  return string(cs[0:len(cs)-1])
}

func restat(s string, at int) string {
  cs := []char(s)
  return string(cs[at:])
}

func toint(s string) int {
  i, _ := strconv.Atoi(s)
  return i
}

func trim(s string) string {
  return strings.TrimSpace(s)
}

func upper(s string) string {
  return strings.ToUpper(s)
}

func lower(s string) string {
  return strings.ToLower(s)
}

func tostr(arr []string) string {
  return strings.Join(arr,``)
}

func has_char(s string,c char) bool {
  for _, sc := range []char(s) {
    if sc == c { return true }
  }
  return false
}

func has_str(ss []string,s string) bool {
  for _, sc := range ss {
    if sc == s { return true }
  }
  return false
}

func has_key(t table, key string) bool {
   _, yes := t[key]
  return yes
}

func has_node(t tree, key string) bool {
  _, yes := t[key]
  return yes
}

func startwith(s string, c char) bool {
  return strings.HasPrefix(s, string(c))
}

func endwith(s, suffix string) bool {
  return strings.HasSuffix(s, suffix)
}

func repeat(s string, n int) string {
  return strings.Repeat(s, n)
}

func itos(i int) string {
  return strconv.Itoa(i)
}

func split(s string, c char) []string {
  return strings.Split(s,string(c))
}

func join(ss []string, s string) string {
  return strings.Join(ss, s)
}

func joinchar(ss []string, c char) string {
  return strings.Join(ss, string(c))
}

func toend(s string) string {
  if len(s) < 60 { return s }
  return s[0:60]
}

func addcc(a,b char) string {
  return string(a) + string(b)
}

func addcs(a char, s string) string {
  return string(a) + s
}

func addsc(s string, a char) string {
  return s + string(a)
}

func getline(cs string,off int) string {
  s := cs[0:off]
  c := strings.Count(s,"\n") + 1
  return itos(c)
}

func mapstrs(ra []string, f func(string) string) []string {
  var ra_map []string
  for _, x := range ra {
    ra_map = append(ra_map, f(x))
  }
  return ra_map
}

func flat2(_arr []string) (string, string) {
  if len(_arr) < 2 {
    fmt.Printf("%v", _arr)
    panic(`flat2 less element`)
  }
  return _arr[0], _arr[1]
}

func flat3(_arr []string) (string, string, string) {
  if len(_arr) < 3 {
    fmt.Printf("%v", _arr)
    panic(`flat3 less element`)
  }
  return _arr[0], _arr[1], _arr[2]
}

func readfile(file string) string {
  buf, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Printf("file %s not exists!", file)
    os.Exit(1)
  }
  return string(buf)
}

func writefile(file,s string) {
  x_data := []byte(s)
  err := ioutil.WriteFile(file, x_data, 0777)
  if err != nil {
    fmt.Printf("file %s not exists!", file)
    os.Exit(1)
  }
}

func copy(from,to string) {
  if isfile(from) {
    s := readfile(from)
    rs := strings.Replace(s,"\r","",-1)
    writefile(rs,to)
    return
  }
  fmt.Printf("%s not exists!\n", from)
}

func isfile(path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	return false
}

func rename(from_file,to_file string) {
  if isfile(from_file) {
    os.Rename(from_file,to_file)
    return
  }
  fmt.Printf("file %s not exists!\n", from_file)
}

func mkdir(dir string) {
  if isfile(dir) { return }
  os.Mkdir(dir,0777)
  os.Chmod(dir,0777)
}

func isvspace(c rune) bool {
  if c == '\n' { return true }
  if c == '\r' { return true }
  return false
}

func ishspace(_c char) bool {
  if _c == ' ' { return true }
  if _c == '\t' { return true }
  return false
}

func isletter(c rune) bool {
  return unicode.IsLetter(c)
}

func isspace(c rune) bool {
  return unicode.IsSpace(c)
}

func islower(c rune) bool {
  return unicode.IsLower(c)
}

func isupper(c rune) bool {
  return unicode.IsUpper(c)
}

func isdigit(c rune) bool {
  return unicode.IsDigit(c)
}

func isalpha(_c char) bool {
  return unicode.IsLetter(_c)
}

func iswords(_c char) bool {
  if isalpha(_c) { return true }
  if isdigit(_c) { return true }
  if _c == '_' { return true }
  if _c == '-' { return true }
  return false
}

func isxdigit(_c char) bool {
  if isdigit(_c) { return true }
  if _c >= 'A' {
    if _c <= 'F' { return true }
  }
  if _c >= 'a' {
    if _c <= 'f' { return true }
  }
  return false
}

// func main() { print("This is Core.go!") }
