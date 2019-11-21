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


func is_str(_s string) bool {
  if first(_s) > QSTR {
    return true
  }
  return false
}

func is_estr(_s string) bool {
  if first(_s) == EIN {
    return true
  }
  return false
}

func names(_s string) []string {
  if has_char(_s,'.') {
    return split(_s,'.')
  }
  return split(_s,':')
}

func is_atom(_s string) bool {
  if startwith(_s,EIN) {
    if second(_s) == QSTR {
      return true
    }
  }
  return false
}

func is_atom_name(_atom string,_name string) bool {
  if is_atom(_atom) {
    if name(_atom) == _name {
      return true
    }
  }
  return false
}

func is_atoms(_s string) bool {
  if first(_s) == EIN {
    if second(_s) == EIN {
      return true
    }
  }
  return false
}

func to_estr(_s string) string {
  if is_estr(_s) {
    return _s
  }
  return addcs(QSTR,_s)
}

func estr(_arr []string) string {
  _buf := new(bytes.Buffer)
  _buf.WriteRune(EIN)
  for _,_s := range _arr {
    _buf.WriteString(to_estr(_s))
  }
  _buf.WriteRune(OUT)
  return _buf.String()
}

func epush(_atoms string,_atom string) string {
  return chop(_atoms) + addsc(_atom,OUT)
}

func einsert(_atom string,_atoms string) string {
  return addcs(EIN,_atom + rest(_atoms))
}

func econcat(_a string,_b string) string {
  return chop(_a) + rest(_b)
}

func name(_estr string) string {
  return atoms(_estr)[0]
}

func value(_estr string) string {
  return atoms(_estr)[1]
}

func atoms(_s string) []string {
  if is_str(_s) {
    return names(_s)
  }
  if is_atom(_s) {
    return items_atom(rest(_s))
  }
  return items_atoms(rest(_s))
}

func items_atom(_estr string) []string {
  _arr := []string{}
  _buf := new(bytes.Buffer)
  _depth := 0
  _s := rest(_estr)
  for _,_c := range []char(_s) {
    if _depth == 0 {
      switch (_c) {
        case EIN:_depth += 1
          _arr = append(_arr,_buf.String())
          _buf = new(bytes.Buffer)
          _buf.WriteRune(_c)
        case QSTR:_arr = append(_arr,_buf.String())
          _buf = new(bytes.Buffer)
        case OUT:_arr = append(_arr,_buf.String())
        default: _buf.WriteRune(_c)
      }
    } else {
      switch (_c) {
        case EIN:_depth += 1
          _buf.WriteRune(_c)
        case OUT:_depth -= 1
          _buf.WriteRune(_c)
        default: _buf.WriteRune(_c)
      }
    }
  }
  return _arr
}

func items_atoms(_estr string) []string {
  _arr := []string{}
  _buf := new(bytes.Buffer)
  _buf.WriteRune(EIN)
  _depth := 1
  _s := rest(_estr)
  for _,_c := range []char(_s) {
    if _depth == 0 {
      switch (_c) {
        case EIN:_depth += 1
          _arr = append(_arr,_buf.String())
          _buf = new(bytes.Buffer)
          _buf.WriteRune(EIN)
        case OUT:_arr = append(_arr,_buf.String())
        default: _buf.WriteRune(_c)
      }
    } else {
      switch (_c) {
        case EIN:_depth += 1
          _buf.WriteRune(_c)
        case OUT:_depth -= 1
          _buf.WriteRune(_c)
        default: _buf.WriteRune(_c)
      }
    }
  }
  return _arr
}

func clean(_ast string) string {
  if is_atom(_ast) {
    return clean_atom(_ast)
  }
  return clean_atoms(_ast)
}

func clean_atom(_atom string) string {
  _name,_value := flat2(atoms(_atom))
  if is_str(_value) {
    return estr([]string{_name,_value})
  }
  if is_atom(_value) {
    return estr([]string{_name,clean_atom(_value)})
  }
  return estr([]string{_name,clean_atoms(_value)})
}

func clean_atoms(_args string) string {
  _atoms := atoms(_args)
  _clean_atoms := mapstrs(_atoms,clean_atom)
  _estr_clean_atoms := estr(_clean_atoms)
  return _estr_clean_atoms
}

func from_ejson(_ejson string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_c := range []char(_ejson) {
    if !isspace(_c) {
      switch (_mode) {
        case 1:_mode = 0
          switch (_c) {
            case 's':_buf.WriteRune(' ')
            case 't':_buf.WriteRune('\t')
            case 'r':_buf.WriteRune('\r')
            case 'n':_buf.WriteRune('\n')
            default: _buf.WriteRune(_c)
          }
        default: switch (_c) {
            case '[':_buf.WriteRune(EIN)
            case ']':_buf.WriteRune(OUT)
            case '.':_buf.WriteRune(QSTR)
            case ':':_mode = 1
            default: _buf.WriteRune(_c)
          }
      }
    }
  }
  return _buf.String()
}

func to_ejson(_estr string) string {
  _buf := new(bytes.Buffer)
  for _,_c := range []char(_estr) {
    switch (_c) {
      case EIN:_buf.WriteRune('[')
      case OUT:_buf.WriteRune(']')
      case QSTR:_buf.WriteRune('.')
      case '[':_buf.WriteString(":[")
      case ']':_buf.WriteString(":]")
      case '.':_buf.WriteString(":.")
      case ':':_buf.WriteString("::")
      case ' ':_buf.WriteString(":s")
      case '\t':_buf.WriteString(":t")
      case '\n':_buf.WriteString(":n")
      case '\r':_buf.WriteString(":r")
      default: _buf.WriteRune(_c)
    }
  }
  return _buf.String()
}

func to_json(_estr string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_c := range []char(_estr) {
    switch (_mode) {
      case 0:switch (_c) {
          case EIN:_buf.WriteRune('[')
          case QSTR:_buf.WriteRune('"')
            _mode = 1
          case OUT:_buf.WriteRune(']')
            _mode = 2
          default: _buf.WriteRune(_c)
        }
      case 1:switch (_c) {
          case EIN:_buf.WriteString("\",[")
            _mode = 0
          case QSTR:_buf.WriteString("\",\"")
          case OUT:_buf.WriteString("\"]")
            _mode = 2
          case ' ':_buf.WriteString("\t")
          case '\n':_buf.WriteString("\n")
          case '\r':_buf.WriteString("\r")
          case '\\':_buf.WriteString("\\")
          case '"':_buf.WriteString("\\\"")
          default: _buf.WriteRune(_c)
        }
      default: switch (_c) {
          case EIN:_buf.WriteString(",[")
            _mode = 0
          case QSTR:_buf.WriteString(",\"")
            _mode = 1
          case OUT:_buf.WriteRune(']')
          default: _buf.WriteRune(_c)
        }
    }
  }
  return _buf.String()
}

func get_spp_ast() string {
  return `[[.top[.rules[[.more[.branch[[.rtoken._s][.ntoken.spec]]]][.eof.EOF]]]][._s[.branch[[.more[.class.s]][.rules[[.str.//][.till[.class.v]]]]]]][.spec[.rules[[.ntoken.name][.cclass[[.char.|][.char.=]]][.gtoken.@branch][.branch[[.char.;][.eof.EOF]]]]]][.name[.chars[[.cclass[[.char.@][.class.a][.char._]]][.many[.cclass[[.char.-][.char._][.class.a]]]]]]][.@branch[.rules[[.gtoken.@rules][.many[.rules[[.char.|][.maybe[.gtoken.@rules]]]]]]]][.@rules[.more[.branch[[.class.s][.gtoken.@rept]]]]][.@rept[.rules[[.gtoken.@rule][.maybe[.ntoken.rept]]]]][.@rule[.branch[[.rtoken._s][.gtoken.@group][.gtoken.@chars][.ntoken.class][.ntoken.cclass][.ntoken.str][.ntoken.not][.ntoken.till][.ntoken.name][.ntoken.word][.ntoken.anychar]]]][.@group[.rules[[.char.(][.gtoken.@branch][.char.)]]]][.rept[.branch[[.cclass[[.char.?][.char.*][.char.+]]][.more[.class.d]]]]][.@chars[.rules[[.char.{][.more[.branch[[.rtoken._s][.gtoken.@rept]]]][.char.}]]]][.anychar[.chars[[.char.<][.till[.char.>]]]]][.str[.chars[[.char.'][.till[.char.']]]]][.class[.chars[[.class.e][.class.S]]]][.cclass[.chars[[.char.:[][.more[.branch[[.gtoken.@range][.ntoken.class]]]][.char.:]]]]][.@range[.chars[[.ntoken.char][.maybe[.rules[[.char.-][.ntoken.char]]]]]]][.char[.not[.cclass[[.char.:]][.class.e]]]]][.not[.rules[[.char.!][.branch[[.ntoken.cclass][.ntoken.class][.ntoken.str]]]]]][.till[.rules[[.char.~][.gtoken.@rule]]]][.word[.chars[[.char.::][.more[.class.w]]]]]]`
}

func new_cursor(_s string,_t table) *Cursor {
  _c := new(Cursor)
  _chars := []char(_s)
  _c.str = _s
  _c.input = _chars
  _c.len = len(_chars)
  _c.at = 0
  _c.output = []string{}
  _c.rtable = _t
  _c.depth = 0
  return _c
}

func match_table(_s string,_t table) string {
  return match_door(_s,_t,`top`)
}

func match_door(_s string,_t table,_name string) string {
  if !has_key(_t,_name) {
    error(fmt.Sprintf("not exists rule: %s",_name))
  }
  _rule := _t[_name]
  _c := new_cursor(_s,_t)
  _match := match_rule(_c,_rule)
  if _match == FAIL {
    fail_report(_c)
  }
  return estr(_c.output)
}

func readchar(_c *Cursor) char {
  _at := _c.at
  if _at >= _c.len {
    return FOF
  }
  return _c.input[_at]
}

func fail_report(_c *Cursor)  {
  _at := _c.at
  _s := _c.str
  _line := getline(_s,_at)
  _code := toend(restat(_s,_at))
  error(fmt.Sprintf("stop at line %s: |%s",_line,_code))
}

func match_rule(_c *Cursor,_atom string) string {
  _name,_rule := flat2(atoms(_atom))
  switch (_name) {
    case `rules`:return match_rules(_c,_rule)
    case `chars`:return match_chars(_c,_rule)
    case `branch`:return match_branch(_c,_rule)
    case `more`:return match_more(_c,_rule)
    case `many`:return match_many(_c,_rule)
    case `maybe`:return match_maybe(_c,_rule)
    case `time`:return match_time(_c,_rule)
    case `word`:return match_word(_c,_rule)
    case `ntoken`:return match_ntoken(_c,_rule)
    case `gtoken`:return match_gtoken(_c,_rule)
    case `rtoken`:return match_rtoken(_c,_rule)
    case `char`:return match_char(_c,_rule)
    case `class`:return match_class(_c,_rule)
    case `cclass`:return match_cclass(_c,_rule)
    case `range`:return match_range(_c,_rule)
    case `str`:return match_str(_c,_rule)
    case `anychar`:return match_anychar(_c,_rule)
    case `not`:return match_not(_c,_rule)
    case `till`:return match_till(_c,_rule)
    case `eof`:return match_eof(_c)
    case `any`:return match_any(_c)
    default: error(fmt.Sprintf("match spp rule miss %s",_name))
  }
  return ``
}

func match_rules(_c *Cursor,_rules string) string {
  _gather := PASS
  for _,_rule := range atoms(_rules) {
    for isspace(readchar(_c)) {
      _c.at += 1
    }
    _match := match_rule(_c,_rule)
    if _match == FAIL {
      return _match
    }
    _gather = gather(_gather,_match)
  }
  return _gather
}

func match_chars(_c *Cursor,_rules string) string {
  _gather := PASS
  for _,_rule := range atoms(_rules) {
    _match := match_rule(_c,_rule)
    if _match == FAIL {
      return _match
    }
    _gather = gather(_gather,_match)
  }
  return _gather
}

func match_branch(_c *Cursor,_rules string) string {
  _at := _c.at
  for _,_rule := range atoms(_rules) {
    _match := match_rule(_c,_rule)
    if _match != FAIL {
      return _match
    }
    if _at != _c.at {
      fail_report(_c)
    }
  }
  return FAIL
}

func match_more(_c *Cursor,_rule string) string {
  _gather := match_rule(_c,_rule)
  if _gather == FAIL {
    return _gather
  }
  _match := match_many(_c,_rule)
  return gather(_gather,_match)
}

func match_many(_c *Cursor,_rule string) string {
  _gather := PASS
  for true {
    _match := match_rule(_c,_rule)
    if _match == FAIL {
      return _gather
    }
    _gather = gather(_gather,_match)
  }
  return _gather
}

func match_time(_c *Cursor,_atom string) string {
  _time_str,_rule := flat2(atoms(_atom))
  _time := toint(_time_str)
  _gather := PASS
  _at := _c.at
  for _time > 0 {
    _match := match_rule(_c,_rule)
    if _match == FAIL {
      _c.at = _at
      return FAIL
    }
    _time -= 1
    _gather = gather(_gather,_match)
  }
  return _gather
}

func match_maybe(_c *Cursor,_rule string) string {
  _match := match_rule(_c,_rule)
  if _match == FAIL {
    return PASS
  }
  return _match
}

func match_token(_c *Cursor,_name string) string {
  _rule := _c.rtable[_name]
  _c.depth += 1
  _match := match_rule(_c,_rule)
  _c.depth -= 1
  return _match
}

func match_ntoken(_c *Cursor,_name string) string {
  _match := match_token(_c,_name)
  if _match == FAIL {
    return _match
  }
  if _match == PASS {
    return _match
  }
  _atom := estr([]string{_name,_match,itos(_c.at)})
  if _c.depth > 0 {
    return _atom
  }
  _c.output = append(_c.output,_atom)
  return PASS
}

func match_gtoken(_c *Cursor,_name string) string {
  _match := match_token(_c,_name)
  if _match == FAIL {
    return _match
  }
  if _match == PASS {
    return _match
  }
  if is_str(_match) {
    return _match
  }
  if is_atom(_match) {
    if _c.depth > 0 {
      return _match
    }
    _c.output = append(_c.output,_match)
    return PASS
  }
  _atom := estr([]string{_name,_match,itos(_c.at)})
  if _c.depth > 0 {
    return _atom
  }
  _c.output = append(_c.output,_atom)
  return PASS
}

func match_rtoken(_c *Cursor,_name string) string {
  _rule := _c.rtable[_name]
  _match := match_rule(_c,_rule)
  if _match == FAIL {
    return _match
  }
  return PASS
}

func match_str(_c *Cursor,_str string) string {
  _at := _c.at
  for _,_a := range []char(_str) {
    if readchar(_c) != _a {
      _c.at = _at
      return FAIL
    }
    _c.at += 1
  }
  return _str
}

func match_anychar(_c *Cursor,_str string) string {
  for _,_a := range []char(_str) {
    if readchar(_c) == _a {
      _c.at += 1
      return string(_a)
    }
  }
  return FAIL
}

func match_word(_c *Cursor,_name string) string {
  _at := _c.at
  for _,_a := range []char(_name) {
    if readchar(_c) != _a {
      _c.at = _at
      return FAIL
    }
    _c.at += 1
  }
  if iswords(readchar(_c)) {
    _c.at = _at
    return FAIL
  }
  return _name
}

func match_range(_c *Cursor,_range string) string {
  _from := first(_range)
  _a := readchar(_c)
  if _a < _from {
    return FAIL
  }
  _to := tail(_range)
  if _a > _to {
    return FAIL
  }
  _c.at += 1
  return string(_a)
}

func match_char(_c *Cursor,_char string) string {
  _a := readchar(_c)
  if first(_char) != _a {
    return FAIL
  }
  _c.at += 1
  return _char
}

func match_class(_c *Cursor,_cchar string) string {
  if is_match_class(_c,_cchar) {
    _a := readchar(_c)
    _c.at += 1
    return string(_a)
  }
  return FAIL
}

func match_cclass(_c *Cursor,_class string) string {
  _a := readchar(_c)
  if is_match_cclass(_c,_class) {
    _c.at += 1
    return string(_a)
  }
  return FAIL
}

func match_eof(_c *Cursor) string {
  _a := readchar(_c)
  if _a == FOF {
    return PASS
  }
  return FAIL
}

func match_any(_c *Cursor) string {
  _a := readchar(_c)
  if _a == FOF {
    return FAIL
  }
  _c.at += 1
  return string(_a)
}

func match_till(_c *Cursor,_rule string) string {
  _buf := new(bytes.Buffer)
  for is_match_not(_c,_rule) {
    _buf.WriteRune(readchar(_c))
    _c.at += 1
  }
  _gather := _buf.String()
  _match := match_rule(_c,_rule)
  if _match == FAIL {
    return _gather
  }
  if _buf.Len() == 0 {
    return _match
  }
  return gather(_gather,_match)
}

func match_not(_c *Cursor,_rule string) string {
  if is_match_rule(_c,_rule) {
    return FAIL
  }
  _a := readchar(_c)
  if _a == FOF {
    return FAIL
  }
  _c.at += 1
  return string(_a)
}

func is_match_not(_c *Cursor,_rule string) bool {
  _a := readchar(_c)
  if _a == FOF {
    return false
  }
  if is_match_rule(_c,_rule) {
    return false
  }
  return true
}

func gather(_gather string,_match string) string {
  if _match == PASS {
    return _gather
  }
  if _gather == PASS {
    return _match
  }
  if is_str(_match) {
    if is_str(_gather) {
      return _gather + _match
    }
    return _gather
  }
  if is_str(_gather) {
    return _match
  }
  if is_atom(_gather) {
    if is_atom(_match) {
      return estr([]string{_gather,_match})
    }
    return einsert(_gather,_match)
  }
  if is_atom(_match) {
    return epush(_gather,_match)
  }
  return econcat(_gather,_match)
}

func is_match_rule(_c *Cursor,_atom string) bool {
  _name,_rule := flat2(atoms(_atom))
  switch (_name) {
    case `cclass`:return is_match_cclass(_c,_rule)
    case `str`:return is_match_str(_c,_rule)
    case `range`:return is_match_range(_c,_rule)
    case `char`:return is_match_char(_c,_rule)
    case `class`:return is_match_class(_c,_rule)
    case `anychar`:return is_match_anychar(_c,_rule)
    default: error(fmt.Sprintf("is match rule miss %s",_name))
  }
  return false
}

func is_match_range(_c *Cursor,_range string) bool {
  _a := readchar(_c)
  if _a == FOF {
    return false
  }
  _from := first(_range)
  if _a < _from {
    return false
  }
  _to := tail(_range)
  if _a > _to {
    return false
  }
  return true
}

func is_match_char(_c *Cursor,_char string) bool {
  _a := readchar(_c)
  if first(_char) == _a {
    return true
  }
  return false
}

func is_match_class(_c *Cursor,_class string) bool {
  _a := readchar(_c)
  switch (_class) {
    case `a`:return isalpha(_a)
    case `A`:return !isalpha(_a)
    case `b`:return _a == ' '
    case `B`:return _a != ' '
    case `c`:return _a == '`'
    case `C`:return _a != '`'
    case `d`:return isdigit(_a)
    case `D`:return !isdigit(_a)
    case `e`:return _a == '\\'
    case `E`:return _a != '\\'
    case `h`:return ishspace(_a)
    case `H`:return !ishspace(_a)
    case `l`:return islower(_a)
    case `L`:return !islower(_a)
    case `n`:return _a == '\n'
    case `N`:return _a != '\n'
    case `q`:return _a == '"'
    case `Q`:return _a != '"'
    case `r`:return _a == '\r'
    case `R`:return _a != '\r'
    case `s`:return isspace(_a)
    case `S`:return !isspace(_a)
    case `t`:return _a == '\t'
    case `T`:return _a != '\t'
    case `u`:return isupper(_a)
    case `U`:return !isupper(_a)
    case `v`:return isvspace(_a)
    case `V`:return !isvspace(_a)
    case `w`:return iswords(_a)
    case `W`:return !iswords(_a)
    case `x`:return isxdigit(_a)
    case `X`:return !isxdigit(_a)
    default: return first(_class) == _a
  }
}

func is_match_cclass(_c *Cursor,_rules string) bool {
  for _,_rule := range atoms(_rules) {
    _match := is_match_rule(_c,_rule)
    if _match == true {
      return true
    }
  }
  return false
}

func is_match_str(_c *Cursor,_str string) bool {
  _at := _c.at
  for _,_char := range []char(_str) {
    if readchar(_c) != _char {
      _c.at = _at
      return false
    }
    _c.at += 1
  }
  _c.at = _at
  return true
}

func is_match_anychar(_c *Cursor,_str string) bool {
  for _,_char := range []char(_str) {
    if readchar(_c) == _char {
      return true
    }
  }
  return false
}

func opt_spp_match(_match string) string {
  _rules := atoms(_match)
  return estr(mapstrs(_rules,get_name_rule))
}

func get_name_rule(_ast string) string {
  _token,_atom := flat2(atoms(value(_ast)))
  _name := value(_token)
  _rule := opt_spp_atom(_atom)
  return estr([]string{_name,_rule})
}

func opt_spp_ast(_atoms string) string {
  if is_atom(_atoms) {
    return estr([]string{opt_spp_atom(_atoms)})
  }
  _rules := []string{}
  for _,_atom := range atoms(_atoms) {
    _rules = append(_rules,opt_spp_atom(_atom))
  }
  return estr(_rules)
}

func opt_spp_atom(_atom string) string {
  _name,_ast := flat2(atoms(_atom))
  switch (_name) {
    case `name`:return opt_spp_token(_ast)
    case `@rules`:return opt_spp_rules(_ast)
    case `@branch`:return opt_spp_branch(_ast)
    case `@chars`:return opt_spp_chars(_ast)
    case `cclass`:return opt_spp_cclass(_ast)
    case `@range`:return opt_spp_range(_ast)
    case `@rept`:return opt_spp_rept(_ast)
    case `str`:return opt_spp_str(_ast)
    case `char`:return opt_spp_char(_ast)
    case `class`:return opt_spp_class(_ast)
    case `anychar`:return opt_spp_anychar(_ast)
    case `not`:return estr([]string{`not`,opt_spp_atom(_ast)})
    case `till`:return estr([]string{`till`,opt_spp_atom(_ast)})
    case `word`:return estr([]string{`word`,rest(_ast)})
    default: return estr([]string{_name,_ast})
  }
}

func opt_spp_rules(_ast string) string {
  if is_str(_ast) {
    return ``
  }
  if is_atom(_ast) {
    return opt_spp_atom(_ast)
  }
  return estr([]string{`rules`,opt_spp_ast(_ast)})
}

func opt_spp_branch(_ast string) string {
  if is_atom(_ast) {
    return opt_spp_atom(_ast)
  }
  return estr([]string{`branch`,opt_spp_ast(_ast)})
}

func opt_spp_chars(_ast string) string {
  if is_str(_ast) {
    return ``
  }
  if is_atom(_ast) {
    return opt_spp_atom(_ast)
  }
  return estr([]string{`chars`,opt_spp_ast(_ast)})
}

func opt_spp_cclass(_ast string) string {
  if is_str(_ast) {
    return ``
  }
  if is_atom(_ast) {
    return opt_spp_atom(_ast)
  }
  return estr([]string{`cclass`,opt_spp_ast(_ast)})
}

func opt_spp_range(_ast string) string {
  _from,_to := flat2(atoms(opt_spp_ast(_ast)))
  _range := value(_from) + value(_to)
  return estr([]string{`range`,_range})
}

func opt_spp_token(_name string) string {
  if _name == `EOF` {
    return estr([]string{`eof`,_name})
  }
  _type := get_spp_token_type(_name)
  return estr([]string{_type,_name})
}

func get_spp_token_type(_name string) string {
  _char := first(_name)
  switch (_char) {
    case '@':return `gtoken`
    case '_':return `rtoken`
    default: return `ntoken`
  }
}

func opt_spp_rept(_ast string) string {
  _atom,_rept := flat2(atoms(_ast))
  _rule := opt_spp_atom(_atom)
  _flag := value(_rept)
  switch (_flag) {
    case `?`:return estr([]string{`maybe`,_rule})
    case `*`:return estr([]string{`many`,_rule})
    case `+`:return estr([]string{`more`,_rule})
    default: return estr([]string{`time`,estr([]string{_flag,_rule})})
  }
}

func opt_spp_str(_ast string) string {
  _str := cut(_ast)
  if len(_str) == 1 {
    return estr([]string{`char`,_str})
  }
  return estr([]string{`str`,_str})
}

func opt_spp_class(_ast string) string {
  _char := tail(_ast)
  if isalpha(_char) {
    return estr([]string{`class`,string(_char)})
  }
  return estr([]string{`char`,string(_char)})
}

func opt_spp_char(_ast string) string {
  _char := tail(_ast)
  return estr([]string{`char`,string(_char)})
}

func opt_spp_anychar(_ast string) string {
  _str := cut(_ast)
  return estr([]string{`anychar`,_str})
}

func opt_my_match(_match string) string {
  _atoms := atoms(_match)
  return estr(mapstrs(_atoms,opt_my_atom))
}

func opt_my_ast(_ast string) string {
  if is_atom(_ast) {
    return estr([]string{opt_my_atom(_ast)})
  }
  return opt_my_match(_ast)
}

func opt_my_atom(_atom string) string {
  _name,_ast,_at := flat3(atoms(_atom))
  switch (_name) {
    case `ns`:return estr([]string{`ns`,value(_ast),_at})
    case `use`:return estr([]string{`use`,value(_ast),_at})
    case `define`:return opt_my_define(_ast,_at)
    case `func`:return opt_my_func(_ast,_at)
    case `struct`:return opt_my_struct(_ast,_at)
    case `fn`:return opt_my_fn(_ast,_at)
    case `@value`:return opt_my_call(_ast,_at)
    case `my`:return opt_my_my(_ast,_at)
    case `for`:return opt_my_for(_ast,_at)
    case `@str`:return opt_my_str(_ast,_at)
    case `@exprs`:return opt_my_exprs(_ast,_at)
    case `ofthen`:return opt_my_ofthen(_ast,_at)
    case `set`:return estr([]string{_name,opt_my_match(_ast),_at})
    case `while`:return estr([]string{_name,opt_my_match(_ast),_at})
    case `of`:return estr([]string{_name,opt_my_match(_ast),_at})
    case `@if`:return estr([]string{`ifelse`,opt_my_match(_ast),_at})
    case `if`:return estr([]string{_name,opt_my_match(_ast),_at})
    case `case`:return estr([]string{_name,opt_my_match(_ast),_at})
    case `block`:return estr([]string{_name,opt_my_ast(_ast),_at})
    case `estr`:return estr([]string{_name,opt_my_ast(_ast),_at})
    case `return`:return estr([]string{_name,opt_my_atom(_ast),_at})
    case `else`:return estr([]string{_name,opt_my_atom(_ast),_at})
    case `then`:return estr([]string{_name,opt_my_atom(_ast),_at})
    case `say`:return estr([]string{_name,opt_my_atom(_ast),_at})
    case `print`:return estr([]string{_name,opt_my_atom(_ast),_at})
    case `error`:return estr([]string{_name,opt_my_atom(_ast),_at})
    case `kstr`:return estr([]string{`kstr`,rest(_ast),_at})
    case `lstr`:return estr([]string{`lstr`,cut(_ast),_at})
    case `const`:return estr([]string{`const`,_ast,_at})
    case `char`:return _atom
    case `name`:return _atom
    case `sym`:return _atom
    case `var`:return _atom
    case `int`:return _atom
    case `dstr`:return _atom
    case `nil`:return _atom
    case `bool`:return _atom
    case `init`:return _atom
    default: error(fmt.Sprintf("opt my atom miss %s",_name))
  }
  return _atom
}

func opt_my_exprs(_ast string,_at string) string {
  return estr([]string{`block`,opt_my_ast(_ast),_at})
}

func opt_my_ofthen(_ast string,_at string) string {
  return estr([]string{`ofthen`,opt_my_ast(_ast),_at})
}

func opt_my_str(_ast string,_at string) string {
  if is_str(_ast) {
    return estr([]string{`dstr`,``,_at})
  }
  return estr([]string{`str`,_ast,_at})
}

func opt_my_define(_args string,_at string) string {
  _value := mapstrs(atoms(_args),value)
  return estr([]string{`define`,estr(_value),_at})
}

func opt_my_func(_atoms string,_at string) string {
  _elems := mapstrs(atoms(_atoms),value)
  _call,_args_type,_ret_type := flat3(_elems)
  _name := fmt.Sprintf("%s.%s",_call,_args_type)
  return estr([]string{`func`,estr([]string{_name,_ret_type}),_at})
}

func opt_my_struct(_args string,_at string) string {
  _name_atom,_fields := flat2(atoms(_args))
  _name := value(_name_atom)
  _pairs := mapstrs(atoms(value(_fields)),value)
  return estr([]string{`struct`,estr([]string{_name,estr(_pairs)}),_at})
}

func opt_my_fn(_expr string,_at string) string {
  _call,_sign,_block := flat3(atoms(_expr))
  _name := value(_call)
  _fn_args,_ret_atom := flat2(atoms(value(_sign)))
  _ret := value(_ret_atom)
  _args := value(_fn_args)
  _type,_names := flat2(atoms(parse_args(_args)))
  _head := estr([]string{fmt.Sprintf("%s.%s",_name,_type),_ret})
  _block_atom := opt_my_atom(_block)
  _fn_atom := estr([]string{_head,estr([]string{_names,_block_atom})})
  return estr([]string{`fn`,_fn_atom,_at})
}

func parse_args(_args string) string {
  if is_str(_args) {
    return estr([]string{`nil`,`nil`})
  }
  if is_atom(_args) {
    return parse_arg(_args)
  }
  return parse_fn_args(_args)
}

func parse_arg(_args string) string {
  _fnarg := parse_fn_arg(value(_args))
  _name,_type := flat2(atoms(_fnarg))
  return estr([]string{_type,estr([]string{fmt.Sprintf("%s:%s",_name,_type)})})
}

func parse_fn_args(_args string) string {
  _names := []string{}
  _types := []string{}
  for _,_var := range atoms(_args) {
    _arg := value(_var)
    _name,_type := flat2(atoms(parse_fn_arg(_arg)))
    _types = append(_types,_type)
    _names = append(_names,fmt.Sprintf("%s:%s",_name,_type))
  }
  _types_str := joinchar(_types,':')
  _names_estr := estr(_names)
  return estr([]string{_types_str,_names_estr})
}

func parse_fn_arg(_arg string) string {
  if has_char(_arg,':') {
    return _arg
  }
  return fmt.Sprintf("%s:str",_arg)
}

func opt_my_my(_args string,_at string) string {
  _opt_args := opt_my_ast(_args)
  _var,_value := flat2(atoms(_opt_args))
  _type,_name := flat2(atoms(_var))
  if _type == `var` {
    return estr([]string{`my`,estr([]string{_name,_value}),_at})
  }
  _atoms := atoms(_name)
  _names := joinchar(mapstrs(_atoms,value),':')
  _my_atom := estr([]string{_names,_value})
  return estr([]string{`mys`,_my_atom,_at})
}

func opt_my_for(_args string,_at string) string {
  _opt_args := opt_my_ast(_args)
  _iter_atom,_set,_block := flat3(atoms(_opt_args))
  _iter := value(_iter_atom)
  return estr([]string{`for`,estr([]string{estr([]string{_iter,_set}),_block}),_at})
}

func opt_my_call(_ast string,_at string) string {
  _atom,_call := flat2(atoms(_ast))
  _atom = opt_my_atom(_atom)
  _type,_value := flat2(atoms(_call))
  return opt_my_calls(_atom,_type,_value,_at)
}

func opt_my_vars(_atom string,_call string,_at string) string {
  _var,_ast := flat2(atoms(_call))
  _atom = estr([]string{`index`,estr([]string{_atom,_var}),_at})
  _type,_value := flat2(atoms(_ast))
  return opt_my_calls(_atom,_type,_value,_at)
}

func opt_my_syms(_atom string,_call string,_at string) string {
  _sym,_ast := flat2(atoms(_call))
  _name := value(_sym)
  _atom = opt_my_ocall(_atom,_name,_at)
  _type,_value := flat2(atoms(_ast))
  return opt_my_calls(_atom,_type,_value,_at)
}

func opt_my_args(_atom string,_call string,_at string) string {
  _args,_ast := flat2(atoms(_call))
  _args = value(_args)
  _atom = opt_my_acall(_atom,_args,_at)
  _type,_value := flat2(atoms(_ast))
  return opt_my_calls(_atom,_type,_value,_at)
}

func opt_my_calls(_atom string,_type string,_value string,_at string) string {
  switch (_type) {
    case `sym`:return opt_my_ocall(_atom,_value,_at)
    case `var`:return opt_my_icall(_atom,_value,_at)
    case `args`:return opt_my_acall(_atom,_value,_at)
    case `pcall`:return opt_my_pcall(_atom,_value,_at)
    case `@vars`:return opt_my_vars(_atom,_value,_at)
    case `@args`:return opt_my_args(_atom,_value,_at)
    default: return opt_my_syms(_atom,_value,_at)
  }
}

func opt_my_acall(_atom string,_value string,_at string) string {
  _args := opt_my_ast(_value)
  _name,_var := flat2(atoms(value(_atom)))
  _args = einsert(_var,_args)
  return estr([]string{`call`,estr([]string{_name,_args}),_at})
}

func opt_my_ocall(_atom string,_name string,_at string) string {
  if name(_atom) == `nil` {
    return estr([]string{`ncall`,fmt.Sprintf("%s.nil",_name),_at})
  }
  return estr([]string{`ocall`,estr([]string{_name,_atom}),_at})
}

func opt_my_icall(_atom string,_value string,_at string) string {
  _var := estr([]string{`var`,_value,_at})
  return estr([]string{`index`,estr([]string{_atom,_var}),_at})
}

func opt_my_pcall(_atom string,_value string,_at string) string {
  _op,_call := flat2(atoms(_value))
  _args := estr([]string{_atom,opt_my_atom(_call)})
  _name := value(_op)
  return estr([]string{`pcall`,estr([]string{_name,_args}),_at})
}

func get_table(_lang string) table {
  _ast := get_spp_ast()
  _estr := from_ejson(_ast)
  _t := ast_to_table(_estr)
  return get_lang_table(_t,_lang)
}

func ast_to_table(_ast string) table {
  _t := table{}
  for _,_spec := range atoms(_ast) {
    _name,_rule := flat2(atoms(_spec))
    _t[_name] = _rule
  }
  return _t
}

func get_lang_table(_t table,_lang string) table {
  _file := get_grammar_file(_lang)
  _grammar := readfile(_file)
  _match := match_table(_grammar,_t)
  _ast := opt_spp_match(_match)
  lint_spp_ast(_ast,_lang)
  return ast_to_table(_ast)
}

func get_grammar_file(_lang string) string {
  _file := fmt.Sprintf("%s.spp",_lang)
  if isfile(_file) {
    return _file
  }
  return fmt.Sprintf("grammar/%s",_file)
}

func spp_repl(_top string)  {
  _t := get_table(`spp`)
  println(fmt.Sprintf("Spp <%s> REPL, type 'exit' exit.",_top))
  for true {
    print(">> ")
    _line := trim(readline())
    if _line == `exit` {
      return 
    }
    _match := match_door(_line,_t,_top)
    println(to_ejson(clean(_match)))
    println(to_ejson(opt_spp_match(_match)))
  }
}

func my_repl(_top string)  {
  _t := get_table(`my`)
  println(fmt.Sprintf("My <%s> REPL, type :exit exit.",_top))
  for true {
    print(">> ")
    _line := trim(readline())
    if _line == `exit` {
      return 
    }
    _match := match_door(_line,_t,_top)
    println(to_ejson(clean(_match)))
    println(to_ejson(clean(opt_my_match(_match))))
  }
}

func lang_repl(_lang string,_top string)  {
  _t := get_table(_lang)
  println(fmt.Sprintf("%s <%s> REPL, type :exit exit.",_lang,_top))
  for true {
    print(">> ")
    _line := trim(readline())
    if _line == `exit` {
      return 
    }
    _match := match_door(_line,_t,_top)
    println(to_ejson(clean(_match)))
  }
}

func get_ast(_lang string) string {
  _boot_ast := get_spp_ast()
  _estr := from_ejson(_boot_ast)
  _t := ast_to_table(_estr)
  _file := fmt.Sprintf("grammar/%s.spp",_lang)
  _grammar := readfile(_file)
  _match := match_table(_grammar,_t)
  _ast := opt_spp_match(_match)
  lint_spp_ast(_ast,_lang)
  return to_ejson(_ast)
}

func lint_spp_ast(_ast string,_lang string)  {
  _t := table{}
  _rules := []string{}
  for _,_atom := range atoms(_ast) {
    _rules = append(_rules,regist_rule(_t,_atom,_lang))
  }
  lint_spp_rules(_t,_rules,_lang)
  lint_spp_table(_t,_lang)
}

func lint_spp_rules(_t table,_rules []string,_lang string)  {
  for _,_rule := range _rules {
    lint_spp_rule(_t,_rule,_lang)
  }
}

func lint_spp_table(_t table,_lang string)  {
  for _name,_ := range _t {
    if _t[_name] == `ok` {
      println(fmt.Sprintf("%s token %s not use",_lang,_name))
    }
  }
}

func regist_rule(_t table,_ast string,_lang string) string {
  _name,_rule := flat2(atoms(_ast))
  if has_key(_t,_name) {
    error(fmt.Sprintf("%s re-define rule %s",_lang,_name))
  }
  if _name == `top` {
    _t[_name] = `use`
  } else {
    _t[_name] = `ok`
  }
  return _rule
}

func lint_spp_rule(_t table,_rule string,_lang string)  {
  _name,_value := flat2(atoms(_rule))
  switch (_name) {
    case `ntoken`:lint_spp_token(_t,_value,_lang)
    case `rtoken`:lint_spp_token(_t,_value,_lang)
    case `gtoken`:lint_spp_token(_t,_value,_lang)
    case `more`:lint_spp_rule(_t,_value,_lang)
    case `many`:lint_spp_rule(_t,_value,_lang)
    case `maybe`:lint_spp_rule(_t,_value,_lang)
    case `not`:lint_spp_rule(_t,_value,_lang)
    case `till`:lint_spp_rule(_t,_value,_lang)
    case `rules`:lint_spp_atoms(_t,_value,_lang)
    case `group`:lint_spp_atoms(_t,_value,_lang)
    case `chars`:lint_spp_atoms(_t,_value,_lang)
    case `branch`:lint_spp_atoms(_t,_value,_lang)
  }
}

func lint_spp_token(_t table,_name string,_lang string)  {
  if has_key(_t,_name) {
    _t[_name] = `use`
  } else {
    error(fmt.Sprintf("%s not define rule %s",_lang,_name))
  }
}

func lint_spp_atoms(_t table,_atoms string,_lang string)  {
  for _,_rule := range atoms(_atoms) {
    lint_spp_rule(_t,_rule,_lang)
  }
}

func lint_my_ast(_ast string) *Lint {
  _t := new(Lint)
  _t.indent = 0
  _t.counter = 0
  _t.stree = tree{}
  _t.stack = []string{}
  regist_ast(_t,_ast)
  lint_my_atoms(_t,_ast)
  reset_block(_t)
  return _t
}

func regist_ast(_t *Lint,_ast string)  {
  for _,_expr := range atoms(_ast) {
    regist_expr(_t,_expr)
  }
}

func regist_expr(_t *Lint,_expr string)  {
  _name,_args,_at := flat3(atoms(_expr))
  _t.at = _at
  switch (_name) {
    case `ns`:regist_ns(_t,_args)
    case `use`:regist_module(_t,_args)
    default: regist_atom(_t,_name,_args)
  }
}

func regist_atom(_t *Lint,_name string,_args string)  {
  switch (_name) {
    case `func`:regist_func(_t,_args)
    case `define`:regist_define(_t,_args)
    case `struct`:regist_struct(_t,_args)
    case `fn`:regist_fn(_t,_args)
  }
}

func regist_ns(_t *Lint,_ns string)  {
  in_ns(_t,_ns)
  _t.ns = _ns
}

func regist_module(_t *Lint,_module string)  {
  _ast_file := fmt.Sprintf("to/o/%s.o",_module)
  _ast := readfile(_ast_file)
  for _,_expr := range atoms(_ast) {
    regist_module_expr(_t,_expr)
  }
}

func regist_module_expr(_t *Lint,_expr string)  {
  _name,_args := flat2(atoms(_expr))
  regist_atom(_t,_name,_args)
}

func regist_define(_t *Lint,_args string)  {
  _name,_type := flat2(atoms(_args))
  set_name_type(_t,_name,_type)
}

func regist_struct(_t *Lint,_args string)  {
  _name,_pairs := flat2(atoms(_args))
  set_name_type(_t,_name,"struct")
  in_ns(_t,_name)
  for _,_pair := range atoms(_pairs) {
    regist_field(_t,_pair)
  }
  out_ns(_t)
}

func regist_field(_t *Lint,_field string)  {
  _name,_type := flat2(atoms(_field))
  set_name_type(_t,_name,_type)
}

func regist_func(_t *Lint,_args string)  {
  _name,_ret_type := flat2(atoms(_args))
  set_name_type(_t,_name,_ret_type)
}

func regist_fn(_t *Lint,_args string)  {
  _fn_head := name(_args)
  regist_func(_t,_fn_head)
}

func lint_my_atoms(_t *Lint,_atoms string)  {
  for _,_atom := range atoms(_atoms) {
    lint_my_atom(_t,_atom)
  }
}

func lint_my_atom(_t *Lint,_atom string)  {
  _name,_args,_at := flat3(atoms(_atom))
  _t.at = _at
  switch (_name) {
    case `fn`:lint_my_fn(_t,_args)
    case `my`:lint_my_my(_t,_args)
    case `mys`:lint_my_mys(_t,_args)
    case `for`:lint_my_for(_t,_args)
    case `else`:lint_my_atom(_t,_args)
    case `then`:lint_my_atom(_t,_args)
    case `say`:lint_my_atom(_t,_args)
    case `print`:lint_my_atom(_t,_args)
    case `ifelse`:lint_my_atoms(_t,_args)
    case `of`:lint_my_atoms(_t,_args)
    case `case`:lint_my_atoms(_t,_args)
    case `estr`:lint_my_atoms(_t,_args)
    case `str`:lint_my_atoms(_t,_args)
    case `if`:lint_my_cond_expr(_t,_args)
    case `while`:lint_my_cond_expr(_t,_args)
    case `block`:lint_my_atoms(_t,_args)
    case `ofthen`:lint_my_atoms(_t,_args)
    case `return`:lint_my_return(_t,_args)
    case `set`:lint_my_set(_t,_args)
    case `ocall`:lint_my_ocall(_t,_args)
    case `ncall`:lint_my_ncall(_t,_args)
    case `pcall`:lint_my_call(_t,_args)
    case `call`:lint_my_call(_t,_args)
    default: return 
  }
}

func lint_my_fn(_t *Lint,_ast string)  {
  _head,_exprs := flat2(atoms(_ast))
  _args,_block := flat2(atoms(_exprs))
  _name_type,_ret_type := flat2(atoms(_head))
  in_ns(_t,_name_type)
  _t.ret = _ret_type
  regist_fn_args(_t,_args)
  lint_my_atom(_t,_block)
  out_ns(_t)
}

func regist_fn_args(_t *Lint,_args string)  {
  if _args == `nil` {
    return 
  }
  for _,_arg := range atoms(_args) {
    regist_fn_arg(_t,_arg)
  }
}

func regist_fn_arg(_t *Lint,_arg string)  {
  _name,_type := flat2(atoms(_arg))
  set_name_type(_t,_name,_type)
}

func lint_my_my(_t *Lint,_args string)  {
  _var,_value := flat2(atoms(_args))
  lint_my_atom(_t,_value)
  _type := get_atom_type(_t,_value)
  set_name_type(_t,_var,_type)
}

func lint_my_mys(_t *Lint,_args string)  {
  _vars,_value := flat2(atoms(_args))
  lint_my_atom(_t,_value)
  for _,_name := range names(_vars) {
    set_name_type(_t,_name,`str`)
  }
}

func lint_my_for(_t *Lint,_args string)  {
  _iter,_block := flat2(atoms(_args))
  _var,_set := flat2(atoms(_iter))
  _set_type := get_atom_type(_t,_set)
  _var_type := get_iter_type(_t,_set_type)
  in_block(_t)
  set_name_type(_t,_var,_var_type)
  lint_my_atom(_t,_block)
  out_block(_t)
}

func lint_my_set(_t *Lint,_args string)  {
  _var,_value := flat2(atoms(_args))
  _atype := get_atom_type(_t,_var)
  _btype := get_atom_type(_t,_value)
  _call_name := fmt.Sprintf("set.%s:%s",_atype,_btype)
  if is_define(_t,_call_name) {
    return 
  }
  report(_t,fmt.Sprintf("no set %s %s!",_atype,_btype))
}

func lint_my_cond_expr(_t *Lint,_args string)  {
  _cond,_expr := flat2(atoms(_args))
  _type := get_atom_type(_t,_cond)
  lint_my_atom(_t,_expr)
  if _type == `bool` {
    return 
  }
  report(_t,"cond return type not bool")
}

func lint_my_ncall(_t *Lint,_call string)  {
  if is_define(_t,_call) {
    return 
  }
  report(_t,fmt.Sprintf("no define call %s",_call))
}

func lint_my_ocall(_t *Lint,_args string)  {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _name_type := fmt.Sprintf("%s.%s",_name,_type)
  if is_define(_t,_name_type) {
    return 
  }
  if is_field(_t,_type,_name) {
    return 
  }
  report(_t,fmt.Sprintf("no call %s.%s",_name,_type))
}

func lint_my_call(_t *Lint,_ast string)  {
  _name,_args := flat2(atoms(_ast))
  _args_type := get_args_type(_t,_args)
  _call := fmt.Sprintf("%s.%s",_name,_args_type)
  if is_define(_t,_call) {
    return 
  }
  report(_t,fmt.Sprintf("undefine call |%s|",_call))
}

func lint_my_return(_t *Lint,_args string)  {
  lint_my_atom(_t,_args)
  _rtype := _t.ret
  _atype := get_atom_type(_t,_args)
  if _atype == _rtype {
    return 
  }
  report(_t,fmt.Sprintf("return %s wish: %s",_atype,_rtype))
}

func get_atom_type(_t *Lint,_atom string) string {
  _name,_value,_at := flat3(atoms(_atom))
  _t.at = _at
  switch (_name) {
    case `init`:return _value
    case `int`:return _name
    case `char`:return _name
    case `str`:return _name
    case `bool`:return _name
    case `nil`:return _name
    case `kstr`:return `str`
    case `lstr`:return `str`
    case `estr`:return `str`
    case `dstr`:return `str`
    case `sym`:return get_sym_type(_t,_value)
    case `var`:return get_var_type(_t,_value)
    case `const`:return get_const_type(_t,_value)
    case `index`:return get_index_type(_t,_value)
    case `pcall`:return get_call_type(_t,_value)
    case `call`:return get_call_type(_t,_value)
    case `ocall`:return get_ocall_type(_t,_value)
    case `ncall`:return get_ncall_type(_t,_value)
    default: return `nil`
  }
}

func get_iter_type(_t *Lint,_type string) string {
  _call_name := fmt.Sprintf("iter.%s",_type)
  return get_name_type(_t,_call_name)
}

func get_index_type(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _call := fmt.Sprintf("index.%s",_type)
  return get_name_type(_t,_call)
}

func get_var_type(_t *Lint,_name string) string {
  return get_name_type(_t,_name)
}

func get_const_type(_t *Lint,_name string) string {
  _type := get_name_type(_t,_name)
  if _type == `struct` {
    return _name
  }
  return _type
}

func get_sym_type(_t *Lint,_name string) string {
  if is_define(_t,fmt.Sprintf("%s.str",_name)) {
    return "fn"
  }
  report(_t,fmt.Sprintf("no define fn %s:str",_name))
  return `nil`
}

func get_call_type(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _name_type := fmt.Sprintf("%s.%s",_name,_type)
  return get_name_type(_t,_name_type)
}

func is_struct(_t *Lint,_name string) bool {
  if !is_define(_t,_name) {
    return false
  }
  _type := get_name_type(_t,_name)
  if _type == `struct` {
    return true
  }
  return false
}

func is_field(_t *Lint,_name string,_field string) bool {
  if !is_struct(_t,_name) {
    return false
  }
  if has_key(_t.stree[_name],_field) {
    return true
  }
  return false
}

func get_field_type(_t *Lint,_name string,_field string) string {
  return _t.stree[_name][_field]
}

func get_ocall_type(_t *Lint,_atom string) string {
  _name,_value := flat2(atoms(_atom))
  _type := get_atom_type(_t,_value)
  _name_type := fmt.Sprintf("%s.%s",_name,_type)
  if is_define(_t,_name_type) {
    return get_name_type(_t,_name_type)
  }
  if is_field(_t,_type,_name) {
    return get_field_type(_t,_type,_name)
  }
  report(_t,fmt.Sprintf("no register call |%s:%s|",_name,_type))
  return `nil`
}

func get_args_type(_t *Lint,_args string) string {
  _types := []string{}
  for _,_atom := range atoms(_args) {
    _types = append(_types,get_atom_type(_t,_atom))
  }
  return joinchar(_types,':')
}

func get_ncall_type(_t *Lint,_call string) string {
  return get_name_type(_t,_call)
}

func is_define(_t *Lint,_name string) bool {
  _tree := _t.stree
  for _,_ns := range _t.stack {
    if has_key(_tree[_ns],_name) {
      return true
    }
  }
  return false
}

func set_name_type(_t *Lint,_name string,_type string)  {
  if is_define(_t,_name) {
    report(_t,fmt.Sprintf("redefine var %s",_name))
  }
  _ns := get_ns(_t)
  _t.stree[_ns][_name] = _type
}

func get_name_type(_t *Lint,_name string) string {
  _tree := _t.stree
  for _,_ns := range _t.stack {
    if has_key(_tree[_ns],_name) {
      return _tree[_ns][_name]
    }
  }
  report(_t,fmt.Sprintf("undefine name |%s|",_name))
  return `nil`
}

func get_indent(_t *Lint) string {
  _indent := _t.indent
  return repeat("  ",_indent)
}

func get_end(_t *Lint) string {
  _indent := get_indent(_t)
  return _indent + `end`
}

func get_ns(_t *Lint) string {
  if len(_t.stack) == 0 {
    return ``
  }
  return _t.stack[0]
}

func reset_block(_t *Lint)  {
  _t.counter = 0
}

func in_ns(_t *Lint,_ns string)  {
  if !has_node(_t.stree,_ns) {
    _t.stree[_ns] = table{}
  }
  if _ns != get_ns(_t) {
    _t.stack = append([]string{_ns},_t.stack...)
  }
}

func out_ns(_t *Lint)  {
  _t.stack = _t.stack[1:]
}

func in_block(_t *Lint)  {
  _ns := itos(_t.counter)
  _t.counter += 1
  in_ns(_t,_ns)
}

func out_block(_t *Lint)  {
  out_ns(_t)
}

func report(_t *Lint,_message string)  {
  _ns := _t.ns
  _code := readfile(fmt.Sprintf("my/%s.my",_ns))
  _line := getline(_code,toint(_t.at))
  error(fmt.Sprintf("** line %s %s",_line,_message))
}

func ast_tomy(_t *Lint,_ast string) string {
  _s := atoms_tomy(_t,_ast)
  _end := "\n\n// vim: ft=my"
  return _s + _end
}

func atoms_tomy(_t *Lint,_atoms string) string {
  _arr := atoms_tomys(_t,_atoms)
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func atoms_tomys(_t *Lint,_ast string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_ast) {
    _arr = append(_arr,atom_tomy(_t,_atom))
  }
  return _arr
}

func block_tomy(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_tomy(_t,_expr)
  _idt := get_indent(_t)
  _t.indent -= 1
  return fmt.Sprintf("%s%s",_idt,_s)
}

func call_args_tomy(_t *Lint,_args string) string {
  if is_atom(_args) {
    return atom_tomy(_t,_args)
  }
  return joinchar(atoms_tomys(_t,_args),' ')
}

func atom_tomy(_t *Lint,_atom string) string {
  _name,_ast := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return fmt.Sprintf("ns %s\n",_ast)
    case `use`:return fmt.Sprintf("use %s",_ast)
    case `fn`:return fn_tomy(_t,_ast)
    case `for`:return for_tomy(_t,_ast)
    case `ifelse`:return atoms_tomy(_t,_ast)
    case `block`:return atoms_tomy(_t,_ast)
    case `ofthen`:return atoms_tomy(_t,_ast)
    case `case`:return case_tomy(_t,_ast)
    case `my`:return my_tomy(_t,_ast)
    case `mys`:return mys_tomy(_t,_ast)
    case `estr`:return estr_tomy(_t,_ast)
    case `ocall`:return ocall_tomy(_t,_ast)
    case `index`:return index_tomy(_t,_ast)
    case `call`:return call_tomy(_t,_ast)
    case `pcall`:return pcall_tomy(_t,_ast)
    case `cond`:return cond_tomy(_t,_ast)
    case `str`:return str_tomy(_ast)
    case `dstr`:return fmt.Sprintf("\"%s\"",_ast)
    case `kstr`:return fmt.Sprintf(":%s",_ast)
    case `lstr`:return fmt.Sprintf("`%s`",_ast)
    case `func`:return func_tomy(_ast)
    case `define`:return define_tomy(_ast)
    case `struct`:return struct_tomy(_ast)
    case `ncall`:return ncall_tomy(_ast)
    case `sym`:return _ast
    case `var`:return _ast
    case `nil`:return _ast
    case `char`:return _ast
    case `bool`:return _ast
    case `const`:return _ast
    case `int`:return _ast
    case `init`:return _ast
    case `if`:return if_tomy(_t,_ast)
    case `else`:return else_tomy(_t,_ast)
    case `while`:return while_tomy(_t,_ast)
    case `of`:return of_tomy(_t,_ast)
    case `then`:return then_tomy(_t,_ast)
    case `return`:return return_tomy(_t,_ast)
    case `say`:return say_tomy(_t,_ast)
    case `print`:return print_tomy(_t,_ast)
    case `error`:return error_tomy(_t,_ast)
    case `set`:return set_tomy(_t,_ast)
    default: println(fmt.Sprintf("atom to mylang miss %s",_name))
  }
  return ``
}

func cond_tomy(_t *Lint,_args string) string {
  _ss := atoms_tomys(_t,_args)
  return joinchar(_ss,'|')
}

func ncall_tomy(_ast string) string {
  _name := name(_ast)
  return fmt.Sprintf("nil.%s",_name)
}

func func_tomy(_ast string) string {
  _call,_ret := flat2(atoms(_ast))
  _name,_args := flat2(atoms(_call))
  _ret = lower(_ret)
  _args = lower(_args)
  return fmt.Sprintf("func %s %s %s",_name,_args,_ret)
}

func define_tomy(_ast string) string {
  _new,_type := flat2(atoms(_ast))
  _type = lower(_type)
  return fmt.Sprintf("define %s %s",_new,_type)
}

func struct_tomy(_ast string) string {
  _name,_pairs := flat2(atoms(_ast))
  _arr := mapstrs(atoms(_pairs),pair_tomy)
  _sep := "\n  "
  _s := _sep + join(_arr,_sep)
  return fmt.Sprintf("\nstruct %s [%s\n]",_name,_s)
}

func pair_tomy(_pair string) string {
  _field,_type := flat2(atoms(_pair))
  return fmt.Sprintf("%s:%s",_field,_type)
}

func fn_tomy(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_exprs := flat2(atoms(_atoms))
  _name_type,_ret := flat2(atoms(_head))
  _name := name(_name_type)
  _args_str := fn_args_tomy(_args)
  _block := block_tomy(_t,_exprs)
  _declare := fmt.Sprintf("%s(%s) %s",_name,_args_str,_ret)
  return fmt.Sprintf("\nfn %s\n%s",_declare,_block)
}

func fn_args_tomy(_args string) string {
  if _args == `nil` {
    return ``
  }
  _arr := mapstrs(atoms(_args),fn_arg_tomy)
  return joinchar(_arr,' ')
}

func fn_arg_tomy(_arg string) string {
  _name,_type := flat2(atoms(_arg))
  if _type == `str` {
    return _name
  }
  return fmt.Sprintf("%s:%s",_name,_type)
}

func for_tomy(_t *Lint,_args string) string {
  _iter_atom,_exprs := flat2(atoms(_args))
  _it,_var_set := flat2(atoms(_iter_atom))
  _set := atom_tomy(_t,_var_set)
  _block := block_tomy(_t,_exprs)
  return fmt.Sprintf("for %s %s\n%s",_it,_set,_block)
}

func case_tomy(_t *Lint,_args string) string {
  _arr := atoms_tomys(_t,_args)
  _cond,_exprs := flat2(_arr)
  _idt := get_indent(_t)
  return fmt.Sprintf("case %s\n%s%s",_cond,_idt,_exprs)
}

func my_tomy(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _value_str := atom_tomy(_t,_value)
  return fmt.Sprintf("my %s %s",_name,_value_str)
}

func mys_tomy(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _vars_str := joinchar(names(_vars),' ')
  _value_str := atom_tomy(_t,_value)
  return fmt.Sprintf("my [%s] %s",_vars_str,_value_str)
}

func of_tomy(_t *Lint,_args string) string {
  _s := cond_block_tomy(_t,_args)
  return fmt.Sprintf("of %s",_s)
}

func cond_block_tomy(_t *Lint,_args string) string {
  _cond_atom,_exprs_atom := flat2(atoms(_args))
  _cond := atom_tomy(_t,_cond_atom)
  _block := block_tomy(_t,_exprs_atom)
  return fmt.Sprintf("%s\n%s",_cond,_block)
}

func then_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("then %s",_s)
}

func if_tomy(_t *Lint,_args string) string {
  _s := cond_block_tomy(_t,_args)
  return fmt.Sprintf("if %s",_s)
}

func else_tomy(_t *Lint,_args string) string {
  _s := block_tomy(_t,_args)
  return fmt.Sprintf("else\n%s",_s)
}

func while_tomy(_t *Lint,_args string) string {
  _s := cond_block_tomy(_t,_args)
  return fmt.Sprintf("while %s",_s)
}

func return_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("return %s",_s)
}

func say_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("say %s",_s)
}

func print_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("print %s",_s)
}

func error_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("error %s",_s)
}

func set_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("set %s",_s)
}

func estr_tomy(_t *Lint,_args string) string {
  _s := call_args_tomy(_t,_args)
  return fmt.Sprintf("[%s]",_s)
}

func str_tomy(_args string) string {
  _s := tostr(mapstrs(atoms(_args),value))
  return fmt.Sprintf("\"%s\"",_s)
}

func ocall_tomy(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _s := atom_tomy(_t,_value)
  return fmt.Sprintf("%s.%s",_s,_name)
}

func index_tomy(_t *Lint,_args string) string {
  _a,_b := flat2(atoms_tomys(_t,_args))
  return fmt.Sprintf("%s.%s",_a,_b)
}

func pcall_tomy(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _a,_b := flat2(atoms_tomys(_t,_args))
  return fmt.Sprintf("%s %s %s",_a,_name,_b)
}

func call_tomy(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _arr := atoms_tomys(_t,_args)
  _n := len(_arr)
  _obj := _arr[0]
  _arr = _arr[1:]
  if _n == 1 {
    return fmt.Sprintf("%s.%s",_obj,_name)
  }
  _s := joinchar(_arr,' ')
  return fmt.Sprintf("%s.%s(%s)",_obj,_name,_s)
}

func ast_togo(_t *Lint,_ast string) string {
  _arr := atoms_togos(_t,_ast)
  return joinchar(_arr,'\n')
}

func atoms_togos(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_togo(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_togo(_t *Lint,_atoms string) string {
  _arr := atoms_togos(_t,_atoms)
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_togo(_t *Lint,_expr string) string {
  _idt := get_indent(_t)
  _t.indent += 1
  _s := atom_togo(_t,_expr)
  _inc_idt := get_indent(_t)
  _t.indent -= 1
  return fmt.Sprintf("{\n%s%s\n%s}",_inc_idt,_s,_idt)
}

func exprs_togo(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_togo(_t,_expr)
  _t.indent -= 1
  return _s
}

func call_args_togo(_t *Lint,_args string) string {
  _arr := atoms_togos(_t,_args)
  return joinchar(_arr,',')
}

func name_togo(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_c := range []char(_name) {
    switch (_c) {
      case '-':_buf.WriteRune('_')
      case '$':_buf.WriteRune('_')
      default: _buf.WriteRune(_c)
    }
  }
  return _buf.String()
}

func atom_togo(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_togo(_t,_args)
    case `fn`:return fn_togo(_t,_args)
    case `for`:return for_togo(_t,_args)
    case `case`:return case_togo(_t,_args)
    case `if`:return if_togo(_t,_args)
    case `ifelse`:return ifelse_togo(_t,_args)
    case `else`:return else_togo(_t,_args)
    case `while`:return while_togo(_t,_args)
    case `block`:return atoms_togo(_t,_args)
    case `ofthen`:return atoms_togo(_t,_args)
    case `of`:return of_togo(_t,_args)
    case `then`:return then_togo(_t,_args)
    case `my`:return my_togo(_t,_args)
    case `mys`:return mys_togo(_t,_args)
    case `set`:return set_togo(_t,_args)
    case `return`:return return_togo(_t,_args)
    case `say`:return say_togo(_t,_args)
    case `print`:return print_togo(_t,_args)
    case `error`:return error_togo(_t,_args)
    case `estr`:return estr_togo(_t,_args)
    case `str`:return str_togo(_t,_args)
    case `ncall`:return ncall_togo(_t,_args)
    case `ocall`:return ocall_togo(_t,_args)
    case `pcall`:return pcall_togo(_t,_args)
    case `index`:return index_togo(_t,_args)
    case `call`:return call_togo(_t,_args)
    case `sym`:return name_togo(_args)
    case `var`:return name_togo(_args)
    case `init`:return init_togo(_args)
    case `const`:return const_togo(_t,_args)
    case `dstr`:return dstr_togo(_args)
    case `lstr`:return fmt.Sprintf("`%s`",_args)
    case `kstr`:return fmt.Sprintf("`%s`",_args)
    case `char`:return _args
    case `int`:return _args
    case `bool`:return _args
    case `use`:return ``
    case `nil`:return ``
    case `define`:return ``
    case `struct`:return ``
    case `func`:return ``
    default: error(fmt.Sprintf("atom to go miss %s",_name))
  }
  return ``
}

func ns_togo(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_togo(_t *Lint,_ast string) string {
  _head,_exprs := flat2(atoms(_ast))
  _args,_block := flat2(atoms(_exprs))
  _call_type,_ret_type := flat2(atoms(_head))
  _call := name(_call_type)
  in_ns(_t,_call_type)
  _ret := type_togo(_t,_ret_type)
  _name := name_togo(_call)
  _astr := fn_args_togo(_t,_args)
  _s := block_togo(_t,_block)
  out_ns(_t)
  if _name == `main` {
    return fmt.Sprintf("\nfunc %s() %s",_name,_s)
  }
  return fmt.Sprintf("\nfunc %s(%s) %s %s",_name,_astr,_ret,_s)
}

func fn_args_togo(_t *Lint,_fnargs string) string {
  if _fnargs == `nil` {
    return ``
  }
  _args := []string{}
  for _,_fnarg := range atoms(_fnargs) {
    _arg_str := fn_arg_togo(_t,_fnarg)
    _args = append(_args,_arg_str)
  }
  return joinchar(_args,',')
}

func fn_arg_togo(_t *Lint,_fnarg string) string {
  _name,_type := flat2(atoms(_fnarg))
  _name_str := name_togo(_name)
  _type_str := type_togo(_t,_type)
  return fmt.Sprintf("%s %s",_name_str,_type_str)
}

func type_togo(_t *Lint,_type string) string {
  if is_struct(_t,_type) {
    return fmt.Sprintf("*%s",_type)
  }
  switch (_type) {
    case `nil`:return ``
    case `str`:return "string"
    case `strs`:return "[]string"
    case `chars`:return "*Buffer"
    default: return _type
  }
}

func for_togo(_t *Lint,_args string) string {
  _iter_atom,_expr := flat2(atoms(_args))
  _iter := for_iter_togo(_t,_iter_atom)
  in_block(_t)
  _block := block_togo(_t,_expr)
  out_block(_t)
  return fmt.Sprintf("for %s %s",_iter,_block)
}

func for_iter_togo(_t *Lint,_atom string) string {
  _iter_name,_set_atom := flat2(atoms(_atom))
  _type := get_atom_type(_t,_set_atom)
  _it := name_togo(_iter_name)
  _set := atom_togo(_t,_set_atom)
  switch (_type) {
    case `str`:return fmt.Sprintf("_,%s := range []char(%s)",_it,_set)
    case `strs`:return fmt.Sprintf("_,%s := range %s",_it,_set)
    default: return fmt.Sprintf("%s,_ := range %s",_it,_set)
  }
}

func case_togo(_t *Lint,_args string) string {
  _cond_atom,_block_atom := flat2(atoms(_args))
  _cond := atom_togo(_t,_cond_atom)
  _block := block_togo(_t,_block_atom)
  return fmt.Sprintf("switch (%s) %s",_cond,_block)
}

func of_togo(_t *Lint,_args string) string {
  _cond,_block := flat2(atoms(_args))
  _cond_str := atom_togo(_t,_cond)
  _block_str := exprs_togo(_t,_block)
  return fmt.Sprintf("case %s:%s",_cond_str,_block_str)
}

func then_togo(_t *Lint,_block string) string {
  _s := exprs_togo(_t,_block)
  return fmt.Sprintf("default: %s",_s)
}

func if_togo(_t *Lint,_args string) string {
  _s := cond_block_togo(_t,_args)
  return fmt.Sprintf("if %s",_s)
}

func cond_block_togo(_t *Lint,_args string) string {
  _cond,_block := flat2(atoms(_args))
  _cond_str := atom_togo(_t,_cond)
  _block_str := block_togo(_t,_block)
  return fmt.Sprintf("%s %s",_cond_str,_block_str)
}

func ifelse_togo(_t *Lint,_exprs string) string {
  _arr := atoms_togos(_t,_exprs)
  return joinchar(_arr,' ')
}

func else_togo(_t *Lint,_block string) string {
  _block_str := block_togo(_t,_block)
  return fmt.Sprintf("else %s",_block_str)
}

func while_togo(_t *Lint,_args string) string {
  _s := cond_block_togo(_t,_args)
  return fmt.Sprintf("for %s",_s)
}

func my_togo(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _var_str := name_togo(_var)
  _vstr := atom_togo(_t,_value)
  return fmt.Sprintf("%s := %s",_var_str,_vstr)
}

func mys_togo(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _names := mapstrs(names(_vars),name_togo)
  _len := itos(len(_names))
  _nstr := joinchar(_names,',')
  _type := get_atom_type(_t,_value)
  _vstr := atom_togo(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("%s := flat%s(%s)",_nstr,_len,_vstr)
  }
  return fmt.Sprintf("%s := flat%s(atoms(%s))",_nstr,_len,_vstr)
}

func set_togo(_t *Lint,_args string) string {
  _arr := atoms_togos(_t,_args)
  return join(_arr," = ")
}

func return_togo(_t *Lint,_args string) string {
  _s := atom_togo(_t,_args)
  return fmt.Sprintf("return %s",_s)
}

func say_togo(_t *Lint,_args string) string {
  _s := atom_togo(_t,_args)
  return fmt.Sprintf("println(%s)",_s)
}

func print_togo(_t *Lint,_args string) string {
  _s := atom_togo(_t,_args)
  return fmt.Sprintf("print(%s)",_s)
}

func error_togo(_t *Lint,_args string) string {
  _s := atom_togo(_t,_args)
  return fmt.Sprintf("error(%s)",_s)
}

func estr_togo(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _s := call_args_togo(_t,_args)
  if _type == `strs` {
    return fmt.Sprintf("estr(%s)",_s)
  }
  return fmt.Sprintf("estr([]string{%s})",_s)
}

func str_togo(_t *Lint,_args string) string {
  _buf := new(bytes.Buffer)
  _names := []string{}
  for _,_atom := range atoms(_args) {
    _name,_value := flat2(atoms(_atom))
    if _name == `dstr` {
      _buf.WriteString(dchars_togo(_value))
    } else {
      _buf.WriteString(`%s`)
      _names = append(_names,name_togo(_value))
    }
  }
  _format := _buf.String()
  _nstr := joinchar(_names,',')
  return fmt.Sprintf("fmt.Sprintf(\"%s\",%s)",_format,_nstr)
}

func dstr_togo(_dstr string) string {
  _s := dchars_togo(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func dchars_togo(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_c := range []char(_s) {
    if _mode == 0 {
      if _c == '\\' {
        _mode = 1
      } else {
        _buf.WriteRune(_c)
      }
    } else {
      _mode = 0
      switch (_c) {
        case '$':_buf.WriteRune(_c)
        case '-':_buf.WriteRune(_c)
        case '%':_buf.WriteString(`%%`)
        default: _buf.WriteString(addcc('\\',_c))
      }
    }
  }
  return _buf.String()
}

func ncall_togo(_t *Lint,_call string) string {
  _name := name_togo(name(_call))
  return fmt.Sprintf("%s()",_name)
}

func index_togo(_t *Lint,_args string) string {
  _data,_at := flat2(atoms_togos(_t,_args))
  return fmt.Sprintf("%s[%s]",_data,_at)
}

func pcall_togo(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _a,_b := flat2(atoms_togos(_t,_args))
  switch (_name) {
    case `+`:return add_togo(_a,_b,_type)
    case `<<`:return push_togo(_a,_b,_type)
    case `>>`:return insert_togo(_a,_b)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func in_togo(_a string,_b string,_type string) string {
  switch (_type) {
    case `char:chars`:return fmt.Sprintf("inchars(%s,%s)",_a,_b)
    default: return fmt.Sprintf("instrs(%s,%s)",_a,_b)
  }
}

func add_togo(_a string,_b string,_type string) string {
  switch (_type) {
    case `char:char`:return fmt.Sprintf("addcc(%s,%s)",_a,_b)
    case `char:str`:return fmt.Sprintf("addcs(%s,%s)",_a,_b)
    case `str:char`:return fmt.Sprintf("addsc(%s,%s)",_a,_b)
    default: return fmt.Sprintf("%s + %s",_a,_b)
  }
}

func push_togo(_a string,_b string,_type string) string {
  switch (_type) {
    case `buffer:char`:return fmt.Sprintf("%s.WriteRune(%s)",_a,_b)
    case `buffer:str`:return fmt.Sprintf("%s.WriteString(%s)",_a,_b)
    default: return fmt.Sprintf("%s = append(%s,%s)",_a,_a,_b)
  }
}

func insert_togo(_a string,_b string) string {
  return fmt.Sprintf("%s = append([]string{%s},%s...)",_b,_a,_b)
}

func ocall_togo(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_togo(_t,_value)
  if is_field(_t,_type,_name) {
    return fmt.Sprintf("%s.%s",_s,_name)
  }
  return call_name_togo(_name,_s,_type)
}

func call_togo(_t *Lint,_atom string) string {
  _call,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _s := call_args_togo(_t,_args)
  return call_name_togo(_call,_s,_type)
}

func init_togo(_name string) string {
  switch (_name) {
    case `int`:return `0`
    case `str`:return "``"
    case `strs`:return "[]string{}"
    case `buffer`:return "new(bytes.Buffer)"
    case `table`:return "table{}"
    case `tree`:return "tree{}"
    default: error(fmt.Sprintf("init to go miss %s",_name))
  }
  return ``
}

func const_togo(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return fmt.Sprintf("new(%s)",_name)
  }
  return _name
}

func call_name_togo(_call string,_s string,_type string) string {
  _name := name_togo(_call)
  switch (_name) {
    case `join`:return join_togo(_s,_type)
    case `tostr`:return tostr_togo(_s,_type)
    case `has`:return has_togo(_s,_type)
    case `len`:return len_togo(_s,_type)
    case `first`:return first_togo(_s,_type)
    case `second`:return second_togo(_s,_type)
    case `shift`:return fmt.Sprintf("%s = %s[1:]",_s,_s)
    case `map`:return fmt.Sprintf("mapstrs(%s)",_s)
    case `inc`:return fmt.Sprintf("%s += 1",_s)
    case `dec`:return fmt.Sprintf("%s -= 1",_s)
    case `not`:return fmt.Sprintf("!%s",_s)
    case `tochars`:return fmt.Sprintf("[]char(%s)",_s)
    default: return fmt.Sprintf("%s(%s)",_name,_s)
  }
}

func len_togo(_s string,_type string) string {
  switch (_type) {
    case `buffer`:return fmt.Sprintf("%s.Len()",_s)
    default: return fmt.Sprintf("len(%s)",_s)
  }
}

func join_togo(_s string,_type string) string {
  if _type == `strs:str` {
    return fmt.Sprintf("join(%s)",_s)
  }
  return fmt.Sprintf("joinchar(%s)",_s)
}

func tostr_togo(_s string,_type string) string {
  switch (_type) {
    case `char`:return fmt.Sprintf("string(%s)",_s)
    case `chars`:return fmt.Sprintf("string(%s)",_s)
    case `buffer`:return fmt.Sprintf("%s.String()",_s)
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    default: return fmt.Sprintf("tostr(%s)",_s)
  }
}

func first_togo(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("first(%s)",_s)
    default: return fmt.Sprintf("%s[0]",_s)
  }
}

func second_togo(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("second(%s)",_s)
    default: return fmt.Sprintf("%s[1]",_s)
  }
}

func has_togo(_s string,_type string) string {
  switch (_type) {
    case `table:str`:return fmt.Sprintf("has_key(%s)",_s)
    case `tree:str`:return fmt.Sprintf("has_node(%s)",_s)
    case `str:char`:return fmt.Sprintf("has_char(%s)",_s)
    default: return fmt.Sprintf("has_str(%s)",_s)
  }
}

func ast_toc(_t *Lint,_ast string) string {
  _fn_str := get_fn_declare(_t,_ast)
  _exprs := atoms_tocs(_t,_ast)
  _ast_str := joinchar(_exprs,'\n')
  return _fn_str + _ast_str
}

func get_fn_declare(_t *Lint,_ast string) string {
  _arr := []string{}
  for _,_expr := range atoms(_ast) {
    _name,_args := flat2(atoms(_expr))
    if _name == `fn` {
      _arr = append(_arr,fn_to_declare(_t,_args))
    }
  }
  return joinchar(_arr,'\n')
}

func fn_to_declare(_t *Lint,_args string) string {
  _name_type,_ret_type := flat2(atoms(name(_args)))
  _call,_types := flat2(atoms(_name_type))
  if _call == `main` {
    return ``
  }
  _name := name_toc(_call)
  _ret := type_toc(_ret_type)
  if _types == `nil` {
    return fmt.Sprintf("%s %s();",_ret,_name)
  }
  _as := joinchar(mapstrs(names(_types),type_toc),',')
  return fmt.Sprintf("%s %s(%s);",_ret,_name,_as)
}

func atoms_tocs(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_toc(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_toc(_t *Lint,_exprs string) string {
  _arr := atoms_tocs(_t,_exprs)
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_toc(_t *Lint,_expr string) string {
  _t.indent += 1
  _block := atom_toc(_t,_expr)
  _idt := get_indent(_t)
  _t.indent -= 1
  _ident := get_indent(_t)
  return fmt.Sprintf("{\n%s%s\n%s}",_idt,_block,_ident)
}

func exprs_toc(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_toc(_t,_expr)
  _t.indent -= 1
  return _s
}

func call_args_toc(_t *Lint,_args string) string {
  _arr := atoms_tocs(_t,_args)
  return joinchar(_arr,',')
}

func atom_toc(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_toc(_t,_args)
    case `fn`:return fn_toc(_t,_args)
    case `for`:return for_toc(_t,_args)
    case `case`:return case_toc(_t,_args)
    case `ifelse`:return atoms_toc(_t,_args)
    case `block`:return atoms_toc(_t,_args)
    case `ofthen`:return atoms_toc(_t,_args)
    case `if`:return if_toc(_t,_args)
    case `else`:return else_toc(_t,_args)
    case `while`:return while_toc(_t,_args)
    case `my`:return my_toc(_t,_args)
    case `mys`:return mys_toc(_t,_args)
    case `return`:return return_toc(_t,_args)
    case `set`:return set_toc(_t,_args)
    case `say`:return say_toc(_t,_args)
    case `print`:return print_toc(_t,_args)
    case `error`:return error_toc(_t,_args)
    case `estr`:return estr_toc(_t,_args)
    case `ncall`:return ncall_toc(_t,_args)
    case `index`:return index_toc(_t,_args)
    case `pcall`:return pcall_toc(_t,_args)
    case `ocall`:return ocall_toc(_t,_args)
    case `call`:return call_toc(_t,_args)
    case `const`:return const_toc(_t,_args)
    case `sym`:return name_toc(_args)
    case `var`:return name_toc(_args)
    case `init`:return init_toc(_args)
    case `str`:return str_toc(_args)
    case `dstr`:return dstr_toc(_args)
    case `lstr`:return lstr_toc(_args)
    case `char`:return _args
    case `kstr`:return fmt.Sprintf("\"%s\"",_args)
    case `bool`:return _args
    case `int`:return _args
    case `nil`:return ``
    case `use`:return ``
    case `define`:return ``
    case `func`:return ``
    case `struct`:return ``
    default: println(fmt.Sprintf("atom to c miss %s",_name))
  }
  return ``
}

func lstr_toc(_lstr string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_lstr) {
    if _a != '\r' {
      switch (_a) {
        case '\n':_buf.WriteString("\\n")
        default: _buf.WriteRune(_a)
      }
    }
  }
  _s := _buf.String()
  return fmt.Sprintf("\"%s\"",_s)
}

func ns_toc(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_toc(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_exprs := flat2(atoms(_atoms))
  _name_type,_ret_type := flat2(atoms(_head))
  in_ns(_t,_name_type)
  _call := name(_name_type)
  _name := name_toc(_call)
  _block := exprs_toc(_t,_exprs)
  out_ns(_t)
  if _name == `main` {
    _hs := "int main(int argc,char* argv[])"
    _es := "  return 0;\n}\n"
    return fmt.Sprintf("\n%s{  %s\n%s",_hs,_block,_es)
  }
  _ret := type_toc(_ret_type)
  _args_str := fn_args_toc(_args)
  _str := fmt.Sprintf("%s %s(%s)",_ret,_name,_args_str)
  return fmt.Sprintf("\n%s {\n  %s\n}",_str,_block)
}

func fn_args_toc(_args string) string {
  if _args == `nil` {
    return ``
  }
  _arr := mapstrs(atoms(_args),fnarg_toc)
  return joinchar(_arr,',')
}

func fnarg_toc(_fnarg string) string {
  _name,_type := flat2(atoms(_fnarg))
  _name_str := name_toc(_name)
  _type_str := type_toc(_type)
  return fmt.Sprintf("%s %s",_type_str,_name_str)
}

func type_toc(_type string) string {
  switch (_type) {
    case `nil`:return `void`
    case `str`:return `char*`
    case `bool`:return `bool`
    case `int`:return `int`
    case `char`:return `char`
    case `chars`:return `char*`
    case `buffer`:return `Buffer*`
    case `strs`:return `Strs*`
    case `table`:return `Table*`
    case `tree`:return `Tree*`
    default: return fmt.Sprintf("%s*",_type)
  }
}

func for_toc(_t *Lint,_args string) string {
  _it_expr,_exprs := flat2(atoms(_args))
  _it_str := iter_toc(_t,_it_expr)
  in_block(_t)
  _exprs_str := exprs_toc(_t,_exprs)
  out_block(_t)
  _idt := get_indent(_t)
  return fmt.Sprintf("%s  %s\n%s}",_it_str,_exprs_str,_idt)
}

func iter_toc(_t *Lint,_expr string) string {
  _var,_it_atom := flat2(atoms(_expr))
  _type := get_atom_type(_t,_it_atom)
  _name := name_toc(_var)
  _it := atom_toc(_t,_it_atom)
  _idt := get_indent(_t)
  switch (_type) {
    case `str`:return iterstr_toc(_idt,_name,_it)
    case `strs`:return iter_strs_toc(_idt,_name,_it)
    default: return iter_table_toc(_idt,_name,_it)
  }
}

func iterstr_toc(_idt string,_var string,_it string) string {
  _bstr := fmt.Sprintf("int i = 0; int n = strlen(%s);",_it)
  _istr := "for (i=0; i < n; i++) {"
  _vstr := fmt.Sprintf("  char %s = %s[i];",_var,_it)
  return fmt.Sprintf("%s\n%s%s\n%s%s\n%s",_bstr,_idt,_istr,_idt,_vstr,_idt)
}

func iter_strs_toc(_idt string,_var string,_it string) string {
  _bstr := fmt.Sprintf("Snode* n = %s->head;",_it)
  _istr := "while (n != NULL) {"
  _vstr := fmt.Sprintf("  char* %s = n->str; n = n->next;",_var)
  return fmt.Sprintf("%s\n%s%s\n%s%s\n%s",_bstr,_idt,_istr,_idt,_vstr,_idt)
}

func iter_table_toc(_idt string,_var string,_it string) string {
  _bstr := fmt.Sprintf("Snode* n = %s->keys->head;",_it)
  _istr := "while (n != NULL) {"
  _vstr := fmt.Sprintf("  char* %s = n->str; n = n->next;",_var)
  return fmt.Sprintf("%s\n%s%s\n%s%s\n%s",_bstr,_idt,_istr,_idt,_vstr,_idt)
}

func case_toc(_t *Lint,_args string) string {
  _case_atom,_exprs := flat2(atoms(_args))
  _case := atom_toc(_t,_case_atom)
  _n := 0
  _if := `if`
  _arr := []string{}
  for _,_expr := range atoms(value(_exprs)) {
    _n += 1
    if _n == 2 {
      _if = "else if"
    }
    _arr = append(_arr,ofthen_toc(_t,_expr,_case,_if))
  }
  return joinchar(_arr,' ')
}

func ofthen_toc(_t *Lint,_expr string,_case string,_if string) string {
  _name,_exprs := flat2(atoms(_expr))
  switch (_name) {
    case `of`:_cond,_block := flat2(atoms(_exprs))
      _cs := cond_toc(_cond,_case)
      _bs := block_toc(_t,_block)
      return fmt.Sprintf("%s (%s) %s",_if,_cs,_bs)
    default: return else_toc(_t,_exprs)
  }
}

func cond_toc(_atom string,_case string) string {
  _name,_value := flat2(atoms(_atom))
  switch (_name) {
    case `kstr`:return fmt.Sprintf("eq(%s,\"%s\")",_case,_value)
    default: return fmt.Sprintf("%s == %s",_case,_value)
  }
}

func if_toc(_t *Lint,_exprs string) string {
  _s := cond_block_toc(_t,_exprs)
  return fmt.Sprintf("if %s",_s)
}

func else_toc(_t *Lint,_exprs string) string {
  _s := block_toc(_t,_exprs)
  return fmt.Sprintf("else %s",_s)
}

func while_toc(_t *Lint,_args string) string {
  _s := cond_block_toc(_t,_args)
  return fmt.Sprintf("while %s",_s)
}

func cond_block_toc(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(_args))
  _cond_str := atom_toc(_t,_cond)
  _exprs_str := block_toc(_t,_exprs)
  return fmt.Sprintf("(%s) %s",_cond_str,_exprs_str)
}

func my_toc(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _tstr := type_toc(_type)
  _vstr := atom_toc(_t,_value)
  _name := name_toc(_var)
  return fmt.Sprintf("%s %s = %s;",_tstr,_name,_vstr)
}

func return_toc(_t *Lint,_args string) string {
  _value := atom_toc(_t,_args)
  return fmt.Sprintf("return %s;",_value)
}

func say_toc(_t *Lint,_args string) string {
  _s := atom_toc(_t,_args)
  return fmt.Sprintf("say(%s);",_s)
}

func print_toc(_t *Lint,_args string) string {
  _s := atom_toc(_t,_args)
  return fmt.Sprintf("print(%s);",_s)
}

func error_toc(_t *Lint,_args string) string {
  _s := atom_toc(_t,_args)
  return fmt.Sprintf("error(%s);",_s)
}

func estr_toc(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _arr := atoms_tocs(_t,_args)
  _len := itos(len(_arr))
  _s := joinchar(_arr,',')
  if _type == `strs` {
    return fmt.Sprintf("estr(%s)",_s)
  }
  return fmt.Sprintf("estr(tostrs(%s,%s))",_len,_s)
}

func ncall_toc(_t *Lint,_call string) string {
  _type := get_ncall_type(_t,_call)
  _name := name_toc(name(_call))
  if _name == `osargs` {
    return "osargs(argc,argv)"
  }
  if _type == `nil` {
    return fmt.Sprintf("%s();",_name)
  }
  return fmt.Sprintf("%s()",_name)
}

func ocall_toc(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_toc(_t,_value)
  if is_field(_t,_type,_name) {
    return fmt.Sprintf("%s->%s",_s,_name)
  }
  return call_name_toc(_t,_name,_s,_type)
}

func pcall_toc(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _a,_b := flat2(atoms_tocs(_t,_args))
  switch (_name) {
    case `+`:return add_toc(_a,_b,_type)
    case `==`:return eq_toc(_a,_b,_type)
    case `!=`:return ne_toc(_a,_b,_type)
    case `<<`:return push_toc(_a,_b,_type)
    case `>>`:return fmt.Sprintf("insert(%s,%s);",_b,_a)
    case `+=`:return fmt.Sprintf("%s += %s;",_a,_b)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func add_toc(_a string,_b string,_type string) string {
  switch (_type) {
    case `char:char`:return fmt.Sprintf("charAddChar(%s,%s)",_a,_b)
    case `str:char`:return fmt.Sprintf("strAddChar(%s,%s)",_a,_b)
    case `char:str`:return fmt.Sprintf("charAddStr(%s,%s)",_a,_b)
    case `str:str`:return fmt.Sprintf("strAddStr(%s,%s)",_a,_b)
    default: return fmt.Sprintf("%s + %s",_a,_b)
  }
}

func eq_toc(_a string,_b string,_type string) string {
  switch (_type) {
    case `str:str`:return fmt.Sprintf("eq(%s,%s)",_a,_b)
    default: return fmt.Sprintf("%s == %s",_a,_b)
  }
}

func ne_toc(_a string,_b string,_type string) string {
  switch (_type) {
    case `str:str`:return fmt.Sprintf("ne(%s,%s)",_a,_b)
    default: return fmt.Sprintf("%s != %s",_a,_b)
  }
}

func push_toc(_a string,_b string,_type string) string {
  switch (_type) {
    case `buffer:char`:return fmt.Sprintf("pushChar(%s,%s);",_a,_b)
    case `buffer:str`:return fmt.Sprintf("pushStr(%s,%s);",_a,_b)
    default: return fmt.Sprintf("push(%s,%s);",_a,_b)
  }
}

func call_toc(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _s := call_args_toc(_t,_args)
  return call_name_toc(_t,_name,_s,_type)
}

func call_name_toc(_t *Lint,_name string,_s string,_type string) string {
  _name_type := fmt.Sprintf("%s.%s",_name,_type)
  _ctype := get_name_type(_t,_name_type)
  _pname := name_toc(_name)
  switch (_name) {
    case `has`:return has_toc(_s,_type)
    case `join`:return join_toc(_s,_type)
    case `len`:return len_toc(_s,_type)
    case `first`:return first_toc(_s,_type)
    case `second`:return second_toc(_s,_type)
    case `rest`:return rest_toc(_s,_type)
    case `tostr`:return tostr_toc(_s,_type)
    case `dec`:return fmt.Sprintf("%s--;",_s)
    case `inc`:return fmt.Sprintf("%s++;",_s)
    case `not`:return fmt.Sprintf("!%s",_s)
    case `toint`:return fmt.Sprintf("atoi(%s)",_s)
    case `tochars`:return _s
    case `shift`:return fmt.Sprintf("shift(%s);",_s)
    default: if _ctype == `nil` {
        return fmt.Sprintf("%s(%s);",_pname,_s)
      } else {
        return fmt.Sprintf("%s(%s)",_pname,_s)
      }
  }
}

func has_toc(_s string,_type string) string {
  switch (_type) {
    case `str:char`:return fmt.Sprintf("hasChar(%s)",_s)
    case `table:str`:return fmt.Sprintf("hasKey(%s)",_s)
    default: return fmt.Sprintf("hasNode(%s)",_s)
  }
}

func tostr_toc(_s string,_type string) string {
  switch (_type) {
    case `int`:return fmt.Sprintf("intToStr(%s)",_s)
    case `char`:return fmt.Sprintf("charToStr(%s)",_s)
    case `strs`:return fmt.Sprintf("strsToStr(%s)",_s)
    case `buffer`:return fmt.Sprintf("bufferToStr(%s)",_s)
    default: return _s
  }
}

func join_toc(_s string,_type string) string {
  switch (_type) {
    case `strs:char`:return fmt.Sprintf("join(%s)",_s)
    default: return fmt.Sprintf("joinStr(%s)",_s)
  }
}

func len_toc(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("strlen(%s)",_s)
    case `chars`:return fmt.Sprintf("strlen(%s)",_s)
    default: return fmt.Sprintf("%s->len",_s)
  }
}

func first_toc(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("%s[0]",_s)
    default: return fmt.Sprintf("first(%s)",_s)
  }
}

func second_toc(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("%s[1]",_s)
    default: return fmt.Sprintf("second(%s)",_s)
  }
}

func rest_toc(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("rest(%s)",_s)
    default: return fmt.Sprintf("restStrs(%s)",_s)
  }
}

func init_toc(_name string) string {
  switch (_name) {
    case `int`:return `0`
    case `str`:return "\"\""
    case `strs`:return "newStrs()"
    case `buffer`:return "newBuffer()"
    case `table`:return "newTable()"
    case `tree`:return "newTree()"
    default: return fmt.Sprintf("new%s()",_name)
  }
}

func const_toc(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return fmt.Sprintf("new%s()",_name)
  }
  return _name
}

func name_toc(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '$':_buf.WriteRune('a')
      case '-':_buf.WriteRune('_')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func str_toc(_atoms string) string {
  _args := []string{}
  for _,_atom := range atoms(_atoms) {
    _args = append(_args,satom_toc(_atom))
  }
  _amount := itos(len(_args))
  _args_str := joinchar(_args,',')
  return fmt.Sprintf("add(%s,%s)",_amount,_args_str)
}

func satom_toc(_atom string) string {
  _type,_value := flat2(atoms(_atom))
  if _type == `dstr` {
    return dstr_toc(_value)
  }
  return name_toc(_value)
}

func dstr_toc(_dstr string) string {
  _s := dchars_toc(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func dchars_toc(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '$':_buf.WriteRune(_a)
          case '-':_buf.WriteRune(_a)
          case '{':_buf.WriteRune(_a)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: switch (_a) {
          case '\\':_mode = 1
          default: _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func set_toc(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _type,_name := flat2(atoms(_var))
  _vstr := atom_toc(_t,_value)
  if _type == `index` {
    return set_index_toc(_t,_name,_vstr)
  }
  _var_str := atom_toc(_t,_var)
  return fmt.Sprintf("%s = %s;",_var_str,_vstr)
}

func set_index_toc(_t *Lint,_args string,_value string) string {
  _type := get_args_type(_t,_args)
  _data,_at := flat2(atoms_tocs(_t,_args))
  switch (_type) {
    case `chars:int`:return fmt.Sprintf("%s[%s] = %s;",_data,_at,_value)
    case `table:str`:return fmt.Sprintf("setKey(%s,%s,%s);",_data,_at,_value)
    default: return fmt.Sprintf("setNode(%s,%s,%s);",_data,_at,_value)
  }
}

func index_toc(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _data,_at := flat2(atoms_tocs(_t,_args))
  switch (_type) {
    case `chars:int`:return fmt.Sprintf("%s[%s]",_data,_at)
    case `table:str`:return fmt.Sprintf("getKey(%s,%s)",_data,_at)
    default: return fmt.Sprintf("getNode(%s,%s)",_data,_at)
  }
}

func mys_toc(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _syms := names(_vars)
  _names := mapstrs(_syms,name_toc)
  _len := itos(len(_names))
  _ds := join(_names,",*")
  _as := join(_names,",&")
  _vs := atom_toc(_t,_value)
  _type := get_atom_type(_t,_value)
  _dstr := fmt.Sprintf("char *%s;",_ds)
  if _type == `strs` {
    return fmt.Sprintf("%s set%s(%s,&%s);",_dstr,_len,_vs,_as)
  }
  return fmt.Sprintf("%s set%s(atoms(%s),&%s);",_dstr,_len,_vs,_as)
}

func ast_topl(_t *Lint,_ast string) string {
  _arr := atoms_topls(_t,_ast)
  return joinchar(_arr,'\n')
}

func atoms_topls(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_topl(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_topl(_t *Lint,_atoms string) string {
  _arr := atoms_topls(_t,_atoms)
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_topl(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_topl(_t,_expr)
  _idt := get_indent(_t)
  _t.indent -= 1
  _ident := get_indent(_t)
  return fmt.Sprintf("{\n%s%s\n%s}",_idt,_s,_ident)
}

func call_args_topl(_t *Lint,_args string) string {
  _arr := atoms_topls(_t,_args)
  return joinchar(_arr,',')
}

func atom_topl(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_topl(_t,_args)
    case `fn`:return fn_topl(_t,_args)
    case `for`:return for_topl(_t,_args)
    case `case`:return case_topl(_t,_args)
    case `ifelse`:return atoms_topl(_t,_args)
    case `block`:return atoms_topl(_t,_args)
    case `ofthen`:return atoms_topl(_t,_args)
    case `of`:return of_topl(_t,_args)
    case `then`:return then_topl(_t,_args)
    case `if`:return if_topl(_t,_args)
    case `else`:return else_topl(_t,_args)
    case `while`:return while_topl(_t,_args)
    case `my`:return my_topl(_t,_args)
    case `mys`:return mys_topl(_t,_args)
    case `set`:return set_topl(_t,_args)
    case `say`:return say_topl(_t,_args)
    case `print`:return print_topl(_t,_args)
    case `error`:return error_topl(_t,_args)
    case `return`:return return_topl(_t,_args)
    case `estr`:return estr_topl(_t,_args)
    case `index`:return index_topl(_t,_args)
    case `lstr`:return fmt.Sprintf("q`%s`",_args)
    case `ncall`:return ncall_topl(_t,_args)
    case `ocall`:return ocall_topl(_t,_args)
    case `pcall`:return pcall_topl(_t,_args)
    case `call`:return call_topl(_t,_args)
    case `sym`:return name_topl(_args)
    case `var`:return name_topl(_args)
    case `init`:return init_topl(_args)
    case `str`:return str_topl(_args)
    case `dstr`:return dstr_topl(_args)
    case `bool`:return bool_topl(_args)
    case `const`:return const_topl(_t,_args)
    case `kstr`:return fmt.Sprintf("'%s'",_args)
    case `int`:return _args
    case `char`:return char_topl(_args)
    case `nil`:return ``
    case `use`:return ``
    default: println(fmt.Sprintf("atom to perl miss %s",_name))
  }
  return ``
}

func char_topl(_str string) string {
  if len(_str) == 3 {
    return _str
  }
  _char := cut(_str)
  return fmt.Sprintf("\"%s\"",_char)
}

func ns_topl(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_topl(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_exprs := flat2(atoms(_atoms))
  _name_type := name(_head)
  in_ns(_t,_name_type)
  _call := name(_name_type)
  _name := name_topl(_call)
  _t.indent += 1
  _block := atom_topl(_t,_exprs)
  _t.indent -= 1
  out_ns(_t)
  if _name == `main` {
    return fmt.Sprintf("\ndo {\n%s\n}",_block)
  }
  if _args == `nil` {
    return fmt.Sprintf("\nsub %s {\n  %s\n}",_name,_block)
  }
  _args_str := fn_args_topl(_args)
  _dec := fmt.Sprintf("  my (%s) = @_;",_args_str)
  return fmt.Sprintf("\nsub %s {\n%s\n  %s\n}",_name,_dec,_block)
}

func fn_args_topl(_args string) string {
  if _args == `nil` {
    return ``
  }
  _arr := mapstrs(atoms(_args),name)
  return joinchar(mapstrs(_arr,name_topl),',')
}

func for_topl(_t *Lint,_args string) string {
  _iter_atom,_exprs := flat2(atoms(_args))
  _iter := iter_topl(_t,_iter_atom)
  in_block(_t)
  _block := block_topl(_t,_exprs)
  out_block(_t)
  return fmt.Sprintf("for %s %s",_iter,_block)
}

func iter_topl(_t *Lint,_atom string) string {
  _iter_name,_set_atom := flat2(atoms(_atom))
  _type := get_atom_type(_t,_set_atom)
  _it := name_topl(_iter_name)
  _set := atom_topl(_t,_set_atom)
  switch (_type) {
    case `strs`:return fmt.Sprintf("my %s (@{%s})",_it,_set)
    case `str`:return fmt.Sprintf("my %s (chars(%s))",_it,_set)
    default: return fmt.Sprintf("my %s (hkeys(%s))",_it,_set)
  }
}

func case_topl(_t *Lint,_args string) string {
  _s := cond_block_topl(_t,_args)
  return fmt.Sprintf("given %s",_s)
}

func of_topl(_t *Lint,_args string) string {
  _s := cond_block_topl(_t,_args)
  return fmt.Sprintf("when %s",_s)
}

func then_topl(_t *Lint,_atom string) string {
  _s := block_topl(_t,_atom)
  return fmt.Sprintf("default %s",_s)
}

func if_topl(_t *Lint,_args string) string {
  _s := cond_block_topl(_t,_args)
  return fmt.Sprintf("if %s",_s)
}

func cond_block_topl(_t *Lint,_args string) string {
  _cond_atom,_expr := flat2(atoms(_args))
  _cond := atom_topl(_t,_cond_atom)
  _block := block_topl(_t,_expr)
  return fmt.Sprintf("(%s) %s",_cond,_block)
}

func else_topl(_t *Lint,_atom string) string {
  _s := block_topl(_t,_atom)
  return fmt.Sprintf("else %s",_s)
}

func while_topl(_t *Lint,_args string) string {
  _s := cond_block_topl(_t,_args)
  return fmt.Sprintf("while %s",_s)
}

func my_topl(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _var_str := name_topl(_var)
  _value_str := atom_topl(_t,_value)
  return fmt.Sprintf("my %s = %s;",_var_str,_value_str)
}

func mys_topl(_t *Lint,_args string) string {
  _names,_value := flat2(atoms(_args))
  _vars := mapstrs(names(_names),name_topl)
  _vars_str := joinchar(_vars,',')
  _type := get_atom_type(_t,_value)
  _vstr := atom_topl(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("my (%s) = @{%s};",_vars_str,_vstr)
  }
  return fmt.Sprintf("my (%s) = @{atoms(%s)};",_vars_str,_vstr)
}

func set_topl(_t *Lint,_args string) string {
  _arr := atoms_topls(_t,_args)
  _var,_value := flat2(_arr)
  return fmt.Sprintf("%s = %s;",_var,_value)
}

func say_topl(_t *Lint,_atom string) string {
  _s := atom_topl(_t,_atom)
  return fmt.Sprintf("say %s;",_s)
}

func print_topl(_t *Lint,_atom string) string {
  _s := atom_topl(_t,_atom)
  return fmt.Sprintf("print %s;",_s)
}

func error_topl(_t *Lint,_atom string) string {
  _s := atom_topl(_t,_atom)
  return fmt.Sprintf("error(%s);",_s)
}

func return_topl(_t *Lint,_atom string) string {
  _s := atom_topl(_t,_atom)
  return fmt.Sprintf("return %s;",_s)
}

func estr_topl(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _s := call_args_topl(_t,_args)
  if _type == `strs` {
    return fmt.Sprintf("estr(%s)",_s)
  }
  return fmt.Sprintf("estr([%s])",_s)
}

func index_topl(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _data,_at := flat2(atoms_topls(_t,_args))
  switch (_type) {
    case `str:int`:return fmt.Sprintf("indexat(%s, %s)",_data,_at)
    case `strs:int`:return fmt.Sprintf("%s->[%s]",_data,_at)
    case `chars:int`:return fmt.Sprintf("%s->[%s]",_data,_at)
    default: return fmt.Sprintf("%s->{%s}",_data,_at)
  }
}

func ncall_topl(_t *Lint,_call string) string {
  _type := get_ncall_type(_t,_call)
  _name := name_topl(name(_call))
  if _type == `nil` {
    return fmt.Sprintf("%s();",_name)
  }
  return fmt.Sprintf("%s()",_name)
}

func ocall_topl(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_topl(_t,_value)
  if is_field(_t,_type,_name) {
    return fmt.Sprintf("%s->{%s}",_s,_name)
  }
  return call_name_topl(_t,_name,_s,_type)
}

func pcall_topl(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _a,_b := flat2(atoms_topls(_t,_args))
  switch (_name) {
    case `+`:return add_topl(_a,_b,_type)
    case `==`:return eq_topl(_a,_b,_type)
    case `!=`:return ne_topl(_a,_b,_type)
    case `>=`:return ge_topl(_a,_b,_type)
    case `<=`:return le_topl(_a,_b,_type)
    case `>`:return gt_topl(_a,_b,_type)
    case `<`:return lt_topl(_a,_b,_type)
    case `<<`:return fmt.Sprintf("push(@{%s},%s);",_a,_b)
    case `>>`:return fmt.Sprintf("unshift(@{%s},%s);",_b,_a)
    case `+=`:return fmt.Sprintf("%s += %s;",_a,_b)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func eq_topl(_a string,_b string,_type string) string {
  if _type == `int:int` {
    return fmt.Sprintf("%s == %s",_a,_b)
  }
  return fmt.Sprintf("%s eq %s",_a,_b)
}

func ne_topl(_a string,_b string,_type string) string {
  if _type == `int:int` {
    return fmt.Sprintf("%s != %s",_a,_b)
  }
  return fmt.Sprintf("%s ne %s",_a,_b)
}

func ge_topl(_a string,_b string,_type string) string {
  if _type == `int:int` {
    return fmt.Sprintf("%s >= %s",_a,_b)
  }
  return fmt.Sprintf("%s ge %s",_a,_b)
}

func le_topl(_a string,_b string,_type string) string {
  if _type == `int:int` {
    return fmt.Sprintf("%s <= %s",_a,_b)
  }
  return fmt.Sprintf("%s le %s",_a,_b)
}

func lt_topl(_a string,_b string,_type string) string {
  if _type == `int:int` {
    return fmt.Sprintf("%s < %s",_a,_b)
  }
  return fmt.Sprintf("%s lt %s",_a,_b)
}

func gt_topl(_a string,_b string,_type string) string {
  if _type == `int:int` {
    return fmt.Sprintf("%s > %s",_a,_b)
  }
  return fmt.Sprintf("%s gt %s",_a,_b)
}

func add_topl(_a string,_b string,_type string) string {
  switch (_type) {
    case `int:int`:return fmt.Sprintf("%s + %s",_a,_b)
    default: return fmt.Sprintf("%s . %s",_a,_b)
  }
}

func call_topl(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _arr := atoms_topls(_t,_args)
  _s := joinchar(_arr,',')
  if _name == `map` {
    return map_topl(_arr)
  }
  return call_name_topl(_t,_name,_s,_type)
}

func map_topl(_arr []string) string {
  _list,_fn := flat2(_arr)
  return fmt.Sprintf("[map { %s($_) } @{%s}]",_fn,_list)
}

func call_name_topl(_t *Lint,_name string,_s string,_type string) string {
  _name_type := fmt.Sprintf("%s.%s",_name,_type)
  _ctype := get_name_type(_t,_name_type)
  _pname := name_topl(_name)
  switch (_pname) {
    case `len`:return len_topl(_s,_type)
    case `has`:return has_topl(_s,_type)
    case `second`:return second_topl(_s,_type)
    case `tostr`:return tostr_topl(_s,_type)
    case `first`:return first_topl(_s,_type)
    case `rest`:return rest_topl(_s,_type)
    case `shift`:return fmt.Sprintf("shift(@{%s});",_s)
    case `dec`:return fmt.Sprintf("%s -= 1;",_s)
    case `inc`:return fmt.Sprintf("%s += 1;",_s)
    case `not`:return fmt.Sprintf("!%s",_s)
    case `chop`:return fmt.Sprintf("chopstr(%s)",_s)
    case `upper`:return fmt.Sprintf("uc(%s)",_s)
    case `lower`:return fmt.Sprintf("lc(%s)",_s)
    case `join`:return fmt.Sprintf("joinstrs(%s)",_s)
    case `split`:return fmt.Sprintf("splitstr(%s)",_s)
    default: if _ctype == `nil` {
        return fmt.Sprintf("%s(%s);",_pname,_s)
      } else {
        return fmt.Sprintf("%s(%s)",_pname,_s)
      }
  }
}

func len_topl(_s string,_type string) string {
  switch (_type) {
    case `strs`:return fmt.Sprintf("len(%s)",_s)
    case `chars`:return fmt.Sprintf("len(%s)",_s)
    default: return fmt.Sprintf("length(%s)",_s)
  }
}

func has_topl(_s string,_type string) string {
  switch (_type) {
    case `str:char`:return fmt.Sprintf("include(%s)",_s)
    default: return fmt.Sprintf("has(%s)",_s)
  }
}

func second_topl(_s string,_type string) string {
  switch (_type) {
    case `strs`:return fmt.Sprintf("%s->[1]",_s)
    default: return fmt.Sprintf("second(%s)",_s)
  }
}

func tostr_topl(_s string,_type string) string {
  switch (_type) {
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    case `char`:return _s
    default: return fmt.Sprintf("string(%s)",_s)
  }
}

func first_topl(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("first(%s)",_s)
    default: return fmt.Sprintf("%s->[0]",_s)
  }
}

func rest_topl(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("rest(%s)",_s)
    default: return fmt.Sprintf("reststrs(%s)",_s)
  }
}

func dstr_topl(_dstr string) string {
  _s := dchars_topl(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func bool_topl(_name string) string {
  if _name == `true` {
    return `1`
  }
  return `0`
}

func str_topl(_args string) string {
  _arr := mapstrs(atoms(_args),satom_topl)
  _s := tostr(_arr)
  return fmt.Sprintf("\"%s\"",_s)
}

func satom_topl(_atom string) string {
  _type,_value := flat2(atoms(_atom))
  if _type == `dstr` {
    return dchars_topl(_value)
  }
  return name_topl(_value)
}

func dchars_topl(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '%':_buf.WriteString(`%%`)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: switch (_a) {
          case '@':_buf.WriteString("\\@")
          case '[':_buf.WriteString("\\[")
          case '{':_buf.WriteString("\\{")
          case '\\':_mode = 1
          default: _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func init_topl(_name string) string {
  switch (_name) {
    case `str`:return "''"
    case `int`:return `0`
    case `strs`:return "[]"
    case `buffer`:return "[]"
    case `table`:return "{}"
    case `tree`:return "{}"
    default: return "{}"
  }
}

func const_topl(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return "{}"
  }
  return _name
}

func name_topl(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '-':_buf.WriteRune('_')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func ast_tolua(_t *Lint,_ast string) string {
  _arr := atoms_toluas(_t,_ast)
  return joinchar(_arr,'\n')
}

func atoms_toluas(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_tolua(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_tolua(_t *Lint,_atoms string) string {
  _arr := atoms_toluas(_t,_atoms)
  return atom_strs_tolua(_t,_arr)
}

func atom_strs_tolua(_t *Lint,_arr []string) string {
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_tolua(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_tolua(_t,_expr)
  _idt := get_indent(_t)
  _t.indent -= 1
  return _idt + _s
}

func call_args_tolua(_t *Lint,_args string) string {
  _arr := atoms_toluas(_t,_args)
  return joinchar(_arr,',')
}

func atom_tolua(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_tolua(_t,_args)
    case `block`:return atoms_tolua(_t,_args)
    case `fnblock`:return atoms_tolua(_t,_args)
    case `fn`:return fn_tolua(_t,_args)
    case `for`:return for_tolua(_t,_args)
    case `case`:return case_tolua(_t,_args)
    case `ifelse`:return ifelse_tolua(_t,_args)
    case `if`:return if_tolua(_t,_args)
    case `while`:return while_tolua(_t,_args)
    case `my`:return my_tolua(_t,_args)
    case `mys`:return mys_tolua(_t,_args)
    case `set`:return set_tolua(_t,_args)
    case `say`:return say_tolua(_t,_args)
    case `print`:return print_tolua(_t,_args)
    case `error`:return error_tolua(_t,_args)
    case `return`:return return_tolua(_t,_args)
    case `estr`:return estr_tolua(_t,_args)
    case `str`:return str_tolua(_t,_args)
    case `ncall`:return ncall_tolua(_args)
    case `ocall`:return ocall_tolua(_t,_args)
    case `index`:return index_tolua(_t,_args)
    case `pcall`:return pcall_tolua(_t,_args)
    case `call`:return call_tolua(_t,_args)
    case `sym`:return name_tolua(_args)
    case `var`:return name_tolua(_args)
    case `dstr`:return dstr_tolua(_args)
    case `const`:return const_tolua(_t,_args)
    case `init`:return init_tolua(_args)
    case `lstr`:return fmt.Sprintf("[=[%s]=]",_args)
    case `kstr`:return fmt.Sprintf("'%s'",_args)
    case `int`:return _args
    case `nil`:return _args
    case `bool`:return _args
    case `char`:return _args
    case `use`:return ``
    default: println(fmt.Sprintf("atom to lua miss %s",_name))
  }
  return ``
}

func const_tolua(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return "{}"
  }
  return _name
}

func ns_tolua(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_tolua(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_block := flat2(atoms(_atoms))
  _name_type := name(_head)
  _name := name_tolua(name(_name_type))
  in_ns(_t,_name_type)
  _args_str := fn_args_tolua(_args)
  _block_str := block_tolua(_t,_block)
  out_ns(_t)
  if _name == `main` {
    return fmt.Sprintf("do\n%s\nend",_block_str)
  }
  _func := fmt.Sprintf("\nfunction %s(%s)",_name,_args_str)
  return fmt.Sprintf("%s\n%s\nend",_func,_block_str)
}

func fn_args_tolua(_args string) string {
  if _args == `nil` {
    return ``
  }
  _names := mapstrs(atoms(_args),name)
  return joinchar(mapstrs(_names,name_tolua),',')
}

func for_tolua(_t *Lint,_args string) string {
  _it_atom,_expr := flat2(atoms(_args))
  _it := iter_tolua(_t,_it_atom)
  in_block(_t)
  _block := block_tolua(_t,_expr)
  out_block(_t)
  _end := get_end(_t)
  return fmt.Sprintf("for %s do\n%s\n%s",_it,_block,_end)
}

func iter_tolua(_t *Lint,_atom string) string {
  _it_name,_iter_atom := flat2(atoms(_atom))
  _type := get_atom_type(_t,_iter_atom)
  _it := name_tolua(_it_name)
  _iter := atom_tolua(_t,_iter_atom)
  switch (_type) {
    case `strs`:return fmt.Sprintf("_,%s in ipairs(%s)",_it,_iter)
    case `table`:return fmt.Sprintf("%s,_ in pairs(%s)",_it,_iter)
    default: return fmt.Sprintf("_,%s in ipairs(chars(%s))",_it,_iter)
  }
}

func case_tolua(_t *Lint,_args string) string {
  _case_atom,_block := flat2(atoms(_args))
  _case := atom_tolua(_t,_case_atom)
  _if := `if`
  _arr := []string{}
  for _,_expr := range atoms(value(_block)) {
    _name,_value := flat2(atoms(_expr))
    if _name == `of` {
      _arr = append(_arr,of_tolua(_t,_value,_if,_case))
      _if = `elseif`
    } else {
      _arr = append(_arr,else_tolua(_t,_value))
    }
  }
  _arr = append(_arr,`end`)
  return atom_strs_tolua(_t,_arr)
}

func of_tolua(_t *Lint,_args string,_if string,_case string) string {
  _cond,_block := flat2(atoms(_args))
  _cond_str := cond_tolua(_t,_cond,_case)
  _s := block_tolua(_t,_block)
  return fmt.Sprintf("%s %s then\n%s",_if,_cond_str,_s)
}

func cond_tolua(_t *Lint,_atom string,_case string) string {
  _name,_value := flat2(atoms(_atom))
  if _name == `cond` {
    _vstr := call_args_tolua(_t,_value)
    return fmt.Sprintf("isin(%s,{%s})",_case,_vstr)
  }
  _s := atom_tolua(_t,_atom)
  return fmt.Sprintf("%s == %s",_case,_s)
}

func ifelse_tolua(_t *Lint,_block string) string {
  _arr := []string{}
  for _,_expr := range atoms(_block) {
    _arr = append(_arr,if_expr_tolua(_t,_expr))
  }
  _arr = append(_arr,`end`)
  return atom_strs_tolua(_t,_arr)
}

func if_expr_tolua(_t *Lint,_expr string) string {
  _name,_value := flat2(atoms(_expr))
  if _name == `if` {
    return iif_tolua(_t,_value)
  }
  return else_tolua(_t,_value)
}

func iif_tolua(_t *Lint,_args string) string {
  _cblock := cond_block_tolua(_t,_args)
  _cond,_block := flat2(atoms(_cblock))
  return fmt.Sprintf("if %s then\n%s",_cond,_block)
}

func cond_block_tolua(_t *Lint,_args string) string {
  _cond_atom,_block_atom := flat2(atoms(_args))
  _cond := atom_tolua(_t,_cond_atom)
  _block := block_tolua(_t,_block_atom)
  return estr([]string{_cond,_block})
}

func else_tolua(_t *Lint,_expr string) string {
  _block := block_tolua(_t,_expr)
  return fmt.Sprintf("else\n%s",_block)
}

func if_tolua(_t *Lint,_args string) string {
  _cblock := cond_block_tolua(_t,_args)
  _cond,_block := flat2(atoms(_cblock))
  _end := get_end(_t)
  return fmt.Sprintf("if %s then\n%s\n%s",_cond,_block,_end)
}

func while_tolua(_t *Lint,_args string) string {
  _cblock := cond_block_tolua(_t,_args)
  _cond,_block := flat2(atoms(_cblock))
  _end := get_end(_t)
  return fmt.Sprintf("while %s do\n%s\n%s",_cond,_block,_end)
}

func my_tolua(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _name := name_tolua(_var)
  _value_str := atom_tolua(_t,_value)
  return fmt.Sprintf("local %s = %s",_name,_value_str)
}

func mys_tolua(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _names := mapstrs(names(_vars),name_tolua)
  _nstr := joinchar(_names,',')
  _type := get_atom_type(_t,_value)
  _vstr := atom_tolua(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("local %s = flat(%s)",_nstr,_vstr)
  }
  return fmt.Sprintf("local %s = flat(atoms(%s))",_nstr,_vstr)
}

func set_tolua(_t *Lint,_args string) string {
  _arr := atoms_toluas(_t,_args)
  _var,_value := flat2(_arr)
  return fmt.Sprintf("%s = %s",_var,_value)
}

func say_tolua(_t *Lint,_args string) string {
  _s := atom_tolua(_t,_args)
  return fmt.Sprintf("print(%s)",_s)
}

func print_tolua(_t *Lint,_args string) string {
  _s := atom_tolua(_t,_args)
  return fmt.Sprintf("io.write(%s)",_s)
}

func error_tolua(_t *Lint,_args string) string {
  _s := atom_tolua(_t,_args)
  return fmt.Sprintf("error(%s)",_s)
}

func return_tolua(_t *Lint,_atom string) string {
  _s := atom_tolua(_t,_atom)
  return fmt.Sprintf("return %s",_s)
}

func estr_tolua(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _s := call_args_tolua(_t,_args)
  switch (_type) {
    case `strs`:return fmt.Sprintf("estr(%s)",_s)
    default: return fmt.Sprintf("estr({%s})",_s)
  }
}

func ocall_tolua(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_tolua(_t,_value)
  if is_field(_t,_type,_name) {
    return fmt.Sprintf("%s.%s",_s,_name)
  }
  return call_value_tolua(_name,_s,_type)
}

func ncall_tolua(_call string) string {
  _name := name_tolua(name(_call))
  return fmt.Sprintf("%s()",_name)
}

func pcall_tolua(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _a,_b := flat2(atoms_toluas(_t,_args))
  switch (_name) {
    case `+`:return add_tolua(_a,_b,_type)
    case `>>`:return fmt.Sprintf("insert(%s,1,%s)",_b,_a)
    case `<<`:return fmt.Sprintf("insert(%s,%s)",_a,_b)
    case `!=`:return fmt.Sprintf("%s ~= %s",_a,_b)
    case `+=`:return fmt.Sprintf("%s = %s + %s",_a,_a,_b)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func add_tolua(_a string,_b string,_type string) string {
  switch (_type) {
    case `int:int`:return fmt.Sprintf("%s + %s",_a,_b)
    default: return fmt.Sprintf("%s .. %s",_a,_b)
  }
}

func index_tolua(_t *Lint,_args string) string {
  _data,_at := flat2(atoms_toluas(_t,_args))
  _type := get_args_type(_t,_args)
  switch (_type) {
    case `str:int`:return fmt.Sprintf("charat(%s,%s)",_data,_at)
    case `strs:int`:return fmt.Sprintf("arrayat(%s,%s)",_data,_at)
    default: return fmt.Sprintf("%s[%s]",_data,_at)
  }
}

func call_tolua(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _arr := atoms_toluas(_t,_args)
  _s := joinchar(_arr,',')
  return call_value_tolua(_name,_s,_type)
}

func call_value_tolua(_call string,_s string,_type string) string {
  _name := name_tolua(_call)
  switch (_name) {
    case `tostr`:return tostr_tolua(_s,_type)
    case `first`:return first_tolua(_s,_type)
    case `second`:return second_tolua(_s,_type)
    case `rest`:return rest_tolua(_s,_type)
    case `has`:return has_tolua(_s,_type)
    case `repeat`:return fmt.Sprintf("rep(%s)",_s)
    case `len`:return fmt.Sprintf("#%s",_s)
    case `dec`:return fmt.Sprintf("%s = %s - 1",_s,_s)
    case `inc`:return fmt.Sprintf("%s = %s + 1",_s,_s)
    case `not`:return fmt.Sprintf("not %s",_s)
    case `tochars`:return fmt.Sprintf("chars(%s)",_s)
    default: return fmt.Sprintf("%s(%s)",_name,_s)
  }
}

func tostr_tolua(_s string,_type string) string {
  switch (_type) {
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    case `char`:return _s
    default: return fmt.Sprintf("tostr(%s)",_s)
  }
}

func first_tolua(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("first(%s)",_s)
    default: return fmt.Sprintf("%s[1]",_s)
  }
}

func second_tolua(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("second(%s)",_s)
    default: return fmt.Sprintf("%s[2]",_s)
  }
}

func rest_tolua(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("rest(%s)",_s)
    default: return fmt.Sprintf("restStrs(%s)",_s)
  }
}

func has_tolua(_s string,_type string) string {
  switch (_type) {
    case `str:char`:return fmt.Sprintf("include(%s)",_s)
    case `str:str`:return fmt.Sprintf("include(%s)",_s)
    default: return fmt.Sprintf("has(%s)",_s)
  }
}

func init_tolua(_name string) string {
  switch (_name) {
    case `int`:return `1`
    case `str`:return "''"
    default: return "{}"
  }
}

func name_tolua(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '-':_buf.WriteRune('_')
      case '$':_buf.WriteRune('s')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func str_tolua(_t *Lint,_args string) string {
  _buf := new(bytes.Buffer)
  _names := []string{}
  for _,_atom := range atoms(_args) {
    _name,_value := flat2(atoms(_atom))
    if _name == `dstr` {
      _buf.WriteString(dchars_tolua(_value))
    } else {
      _buf.WriteString(`%s`)
      _names = append(_names,name_tolua(_value))
    }
  }
  _f := _buf.String()
  _s := joinchar(_names,',')
  return fmt.Sprintf("format(\"%s\",%s)",_f,_s)
}

func dstr_tolua(_dstr string) string {
  _s := dchars_tolua(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func dchars_tolua(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '$':_buf.WriteRune(_a)
          case '-':_buf.WriteRune(_a)
          case '%':_buf.WriteString(`%%`)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: if _a == '\\' {
          _mode = 1
        } else {
          _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func ast_torb(_t *Lint,_ast string) string {
  _arr := atoms_torbs(_t,_ast)
  return joinchar(_arr,'\n')
}

func atoms_torbs(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_torb(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_torb(_t *Lint,_atoms string) string {
  _arr := atoms_torbs(_t,_atoms)
  return atom_strs_torb(_t,_arr)
}

func atom_strs_torb(_t *Lint,_arr []string) string {
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func call_args_torb(_t *Lint,_args string) string {
  _arr := atoms_torbs(_t,_args)
  return joinchar(_arr,',')
}

func block_torb(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_torb(_t,_expr)
  _idt := get_indent(_t)
  _t.indent -= 1
  return fmt.Sprintf("%s%s",_idt,_s)
}

func atom_torb(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_torb(_t,_args)
    case `fn`:return fn_torb(_t,_args)
    case `block`:return atoms_torb(_t,_args)
    case `ofthen`:return atoms_torb(_t,_args)
    case `for`:return for_torb(_t,_args)
    case `case`:return case_torb(_t,_args)
    case `if`:return if_torb(_t,_args)
    case `while`:return while_torb(_t,_args)
    case `ifelse`:return ifelse_torb(_t,_args)
    case `of`:return of_torb(_t,_args)
    case `then`:return then_torb(_t,_args)
    case `set`:return set_torb(_t,_args)
    case `index`:return index_torb(_t,_args)
    case `my`:return my_torb(_t,_args)
    case `mys`:return mys_torb(_t,_args)
    case `return`:return return_torb(_t,_args)
    case `say`:return say_torb(_t,_args)
    case `print`:return print_torb(_t,_args)
    case `error`:return error_torb(_t,_args)
    case `estr`:return estr_torb(_t,_args)
    case `ncall`:return ncall_torb(_args)
    case `ocall`:return ocall_torb(_t,_args)
    case `pcall`:return pcall_torb(_t,_args)
    case `call`:return call_torb(_t,_args)
    case `const`:return const_torb(_t,_args)
    case `sym`:return name_torb(_args)
    case `var`:return name_torb(_args)
    case `init`:return init_torb(_args)
    case `str`:return str_torb(_args)
    case `dstr`:return dstr_torb(_args)
    case `lstr`:return fmt.Sprintf("<<STR\n%s\nSTR",_args)
    case `kstr`:return fmt.Sprintf("'%s'",_args)
    case `bool`:return _args
    case `nil`:return _args
    case `char`:return char_torb(_args)
    case `int`:return _args
    case `use`:return ``
    default: println(fmt.Sprintf("atom to ruby miss %s",_name))
  }
  return ``
}

func char_torb(_str string) string {
  if len(_str) == 3 {
    return _str
  }
  _char := cut(_str)
  return fmt.Sprintf("\"%s\"",_char)
}

func ns_torb(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_torb(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_exprs := flat2(atoms(_atoms))
  _name_type := name(_head)
  in_ns(_t,_name_type)
  _name := name_torb(name(_name_type))
  _args_str := fn_args_torb(_args)
  _block := block_torb(_t,_exprs)
  out_ns(_t)
  if _name == `main` {
    return fmt.Sprintf("begin\n%s\nend",_block)
  }
  _head_str := fmt.Sprintf("\ndef %s(%s)",_name,_args_str)
  return fmt.Sprintf("%s\n%s\nend",_head_str,_block)
}

func fn_args_torb(_args string) string {
  if _args == `nil` {
    return ``
  }
  _arr := mapstrs(atoms(_args),name)
  return joinchar(mapstrs(_arr,name_torb),',')
}

func for_torb(_t *Lint,_args string) string {
  _iter_expr,_exprs := flat2(atoms(_args))
  _iter_str := iter_torb(_t,_iter_expr)
  in_block(_t)
  _s := block_torb(_t,_exprs)
  out_block(_t)
  _end := get_end(_t)
  return fmt.Sprintf("%s\n%s\n%s",_iter_str,_s,_end)
}

func iter_torb(_t *Lint,_atom string) string {
  _iter_name,_set_atom := flat2(atoms(_atom))
  _type := get_atom_type(_t,_set_atom)
  _iter := name_torb(_iter_name)
  _set := atom_torb(_t,_set_atom)
  switch (_type) {
    case `strs`:return fmt.Sprintf("%s.each do |%s|",_set,_iter)
    case `str`:return fmt.Sprintf("%s.each_char do |%s|",_set,_iter)
    default: return fmt.Sprintf("%s.keys.each do |%s|",_set,_iter)
  }
}

func case_torb(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(cond_block_torb(_t,_args)))
  _end := get_end(_t)
  return fmt.Sprintf("case (%s)\n%s\n%s",_cond,_exprs,_end)
}

func of_torb(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(cond_block_torb(_t,_args)))
  return fmt.Sprintf("when %s then\n%s",_cond,_exprs)
}

func then_torb(_t *Lint,_args string) string {
  _s := block_torb(_t,_args)
  return fmt.Sprintf("else\n%s",_s)
}

func ifelse_torb(_t *Lint,_exprs string) string {
  _arr := []string{}
  for _,_expr := range atoms(_exprs) {
    _arr = append(_arr,expr_torb(_t,_expr))
  }
  _arr = append(_arr,`end`)
  return atom_strs_torb(_t,_arr)
}

func expr_torb(_t *Lint,_expr string) string {
  _name,_value := flat2(atoms(_expr))
  if _name == `if` {
    return if_if_torb(_t,_value)
  }
  return else_torb(_t,_value)
}

func if_torb(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(cond_block_torb(_t,_args)))
  _end := get_end(_t)
  return fmt.Sprintf("if %s then\n%s\n%s",_cond,_exprs,_end)
}

func if_if_torb(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(cond_block_torb(_t,_args)))
  return fmt.Sprintf("if %s then\n%s",_cond,_exprs)
}

func cond_block_torb(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(_args))
  _cond_str := atom_torb(_t,_cond)
  _exprs_str := block_torb(_t,_exprs)
  return estr([]string{_cond_str,_exprs_str})
}

func else_torb(_t *Lint,_exprs string) string {
  _s := block_torb(_t,_exprs)
  return fmt.Sprintf("else\n%s",_s)
}

func while_torb(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(cond_block_torb(_t,_args)))
  _end := get_end(_t)
  return fmt.Sprintf("while %s do\n%s\n%s",_cond,_exprs,_end)
}

func my_torb(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _var_str := name_torb(_var)
  _vstr := atom_torb(_t,_value)
  return fmt.Sprintf("%s = %s",_var_str,_vstr)
}

func set_torb(_t *Lint,_args string) string {
  _arr := atoms_torbs(_t,_args)
  _var,_value := flat2(_arr)
  return fmt.Sprintf("%s = %s",_var,_value)
}

func index_torb(_t *Lint,_args string) string {
  _data,_at := flat2(atoms_torbs(_t,_args))
  return fmt.Sprintf("%s[%s]",_data,_at)
}

func mys_torb(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _names := mapstrs(names(_vars),name_torb)
  _nstr := joinchar(_names,',')
  _vstr := atom_torb(_t,_value)
  _type := get_atom_type(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("%s = %s",_nstr,_vstr)
  }
  return fmt.Sprintf("%s = atoms(%s)",_nstr,_vstr)
}

func return_torb(_t *Lint,_atom string) string {
  _s := atom_torb(_t,_atom)
  return fmt.Sprintf("return %s",_s)
}

func say_torb(_t *Lint,_atom string) string {
  _s := atom_torb(_t,_atom)
  return fmt.Sprintf("p %s",_s)
}

func print_torb(_t *Lint,_atom string) string {
  _s := atom_torb(_t,_atom)
  return fmt.Sprintf("print(%s)",_s)
}

func error_torb(_t *Lint,_atom string) string {
  _s := atom_torb(_t,_atom)
  return fmt.Sprintf("error(%s)",_s)
}

func ncall_torb(_call string) string {
  _name := name_torb(name(_call))
  return fmt.Sprintf("%s()",_name)
}

func estr_torb(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _s := call_args_torb(_t,_args)
  switch (_type) {
    case `strs`:return fmt.Sprintf("estr(%s)",_s)
    default: return fmt.Sprintf("estr([%s])",_s)
  }
}

func ocall_torb(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_torb(_t,_value)
  if is_field(_t,_type,_name) {
    return fmt.Sprintf("%s['%s']",_s,_name)
  }
  return call_name_torb(_name,_s,_type)
}

func pcall_torb(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _a,_b := flat2(atoms_torbs(_t,_args))
  switch (_name) {
    case `>>`:return fmt.Sprintf("%s.unshift(%s)",_b,_a)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func call_torb(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _arr := atoms_torbs(_t,_args)
  if _name == `map` {
    return map_torb(_arr)
  }
  _s := joinchar(_arr,',')
  _type := get_args_type(_t,_args)
  return call_name_torb(_name,_s,_type)
}

func map_torb(_arr []string) string {
  _array,_fn := flat2(_arr)
  return fmt.Sprintf("%s.map { |n| %s(n) }",_array,_fn)
}

func call_name_torb(_name string,_s string,_type string) string {
  _rname := name_torb(_name)
  switch (_name) {
    case `tostr`:return tostr_torb(_s,_type)
    case `has`:return has_torb(_s,_type)
    case `trim`:return fmt.Sprintf("%s.strip",_s)
    case `len`:return fmt.Sprintf("%s.length",_s)
    case `shift`:return fmt.Sprintf("%s.shift",_s)
    case `dec`:return fmt.Sprintf("%s -= 1",_s)
    case `inc`:return fmt.Sprintf("%s += 1",_s)
    case `upper`:return fmt.Sprintf("%s.upcase",_s)
    case `lower`:return fmt.Sprintf("%s.downcase",_s)
    case `tochars`:return fmt.Sprintf("to_chars(%s)",_s)
    default: return fmt.Sprintf("%s(%s)",_rname,_s)
  }
}

func tostr_torb(_s string,_type string) string {
  switch (_type) {
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    case `char`:return _s
    default: return fmt.Sprintf("to_string(%s)",_s)
  }
}

func has_torb(_s string,_type string) string {
  switch (_type) {
    case `str:char`:return fmt.Sprintf("find(%s)",_s)
    case `str:str`:return fmt.Sprintf("find(%s)",_s)
    default: return fmt.Sprintf("has(%s)",_s)
  }
}

func const_torb(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return "{}"
  }
  return _name
}

func init_torb(_name string) string {
  switch (_name) {
    case `int`:return "0"
    case `str`:return "''"
    case `strs`:return "[]"
    case `buffer`:return "[]"
    case `table`:return "{}"
    case `tree`:return "{}"
    default: return "{}"
  }
}

func name_torb(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '-':_buf.WriteRune('_')
      case '$':_buf.WriteRune('_')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func str_torb(_args string) string {
  _arr := mapstrs(atoms(_args),satom_torb)
  _s := tostr(_arr)
  return fmt.Sprintf("\"%s\"",_s)
}

func satom_torb(_atom string) string {
  _type,_value := flat2(atoms(_atom))
  switch (_type) {
    case `var`:_name := name_torb(_value)
      return fmt.Sprintf("#{%s}",_name)
    default: return dchars_torb(_value)
  }
}

func dchars_torb(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '$':_buf.WriteRune(_a)
          case '-':_buf.WriteRune(_a)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: switch (_a) {
          case '#':_buf.WriteString("\\#")
          case '\\':_mode = 1
          default: _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func dstr_torb(_dstr string) string {
  _s := dchars_torb(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func ast_topy(_t *Lint,_ast string) string {
  _arr := atoms_topys(_t,_ast)
  return joinchar(_arr,'\n')
}

func atoms_topys(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_topy(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_topy(_t *Lint,_atoms string) string {
  _arr := atoms_topys(_t,_atoms)
  return atom_strs_topy(_t,_arr)
}

func atom_strs_topy(_t *Lint,_arr []string) string {
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_topy(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_topy(_t,_expr)
  _idt := get_indent(_t)
  _t.indent -= 1
  return fmt.Sprintf("\n%s%s",_idt,_s)
}

func call_args_topy(_t *Lint,_args string) string {
  _arr := atoms_topys(_t,_args)
  return joinchar(_arr,',')
}

func atom_topy(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_topy(_t,_args)
    case `fn`:return fn_topy(_t,_args)
    case `for`:return for_topy(_t,_args)
    case `case`:return case_topy(_t,_args)
    case `if`:return if_topy(_t,_args)
    case `else`:return else_topy(_t,_args)
    case `while`:return while_topy(_t,_args)
    case `ifelse`:return atoms_topy(_t,_args)
    case `block`:return atoms_topy(_t,_args)
    case `my`:return my_topy(_t,_args)
    case `mys`:return mys_topy(_t,_args)
    case `return`:return return_topy(_t,_args)
    case `say`:return say_topy(_t,_args)
    case `print`:return print_topy(_t,_args)
    case `error`:return error_topy(_t,_args)
    case `estr`:return estr_topy(_t,_args)
    case `ncall`:return ncall_topy(_args)
    case `ocall`:return ocall_topy(_t,_args)
    case `pcall`:return pcall_topy(_t,_args)
    case `call`:return call_topy(_t,_args)
    case `const`:return const_topy(_t,_args)
    case `set`:return set_topy(_t,_args)
    case `index`:return index_topy(_t,_args)
    case `sym`:return name_topy(_args)
    case `var`:return name_topy(_args)
    case `init`:return init_topy(_args)
    case `str`:return str_topy(_args)
    case `dstr`:return dstr_topy(_args)
    case `lstr`:return fmt.Sprintf("'''%s'''",_args)
    case `kstr`:return fmt.Sprintf("'%s'",_args)
    case `int`:return _args
    case `char`:return _args
    case `bool`:return bool_topy(_args)
    case `nil`:return `None`
    case `use`:return ``
    default: println(fmt.Sprintf("atom to python miss %s",_name))
  }
  return ``
}

func bool_topy(_name string) string {
  if _name == `true` {
    return `True`
  }
  return `False`
}

func ns_topy(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_topy(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_exprs := flat2(atoms(_atoms))
  _name_type := name(_head)
  in_ns(_t,_name_type)
  _name := name_topy(name(_name_type))
  _args_str := fn_args_topy(_args)
  _exprs_str := block_topy(_t,_exprs)
  out_ns(_t)
  _fnhead := "if __name__== '__main__':"
  if _name == `main` {
    return fmt.Sprintf("\n%s%s",_fnhead,_exprs_str)
  }
  _declare := fmt.Sprintf("\ndef %s(%s)",_name,_args_str)
  return fmt.Sprintf("%s:%s",_declare,_exprs_str)
}

func fn_args_topy(_args string) string {
  if _args == `nil` {
    return ``
  }
  _arr := mapstrs(atoms(_args),name)
  return joinchar(mapstrs(_arr,name_topy),',')
}

func for_topy(_t *Lint,_args string) string {
  _iter_expr,_exprs := flat2(atoms(_args))
  _iter_str := iter_topy(_t,_iter_expr)
  in_block(_t)
  _exprs_str := block_topy(_t,_exprs)
  out_block(_t)
  return fmt.Sprintf("for %s:%s",_iter_str,_exprs_str)
}

func iter_topy(_t *Lint,_atom string) string {
  _name,_set_atom := flat2(atoms(_atom))
  _type := get_atom_type(_t,_set_atom)
  _iter := name_topy(_name)
  _set := atom_topy(_t,_set_atom)
  switch (_type) {
    case `str`:return fmt.Sprintf("%s in %s",_iter,_set)
    case `strs`:return fmt.Sprintf("%s in %s",_iter,_set)
    default: return fmt.Sprintf("%s in %s.keys()",_iter,_set)
  }
}

func case_topy(_t *Lint,_args string) string {
  _case_atom,_exprs := flat2(atoms(_args))
  _case := atom_topy(_t,_case_atom)
  _if := `if`
  _arr := []string{}
  _n := 0
  for _,_expr := range atoms(value(_exprs)) {
    if _n > 0 {
      _if = `elif`
    }
    _n += 1
    _name,_value := flat2(atoms(_expr))
    if _name == `of` {
      _arr = append(_arr,of_topy(_t,_value,_if,_case))
    } else {
      _arr = append(_arr,else_topy(_t,_value))
    }
  }
  return atom_strs_topy(_t,_arr)
}

func of_topy(_t *Lint,_args string,_if string,_case string) string {
  _cond,_exprs := flat2(atoms(_args))
  _cs := atom_topy(_t,_cond)
  _s := block_topy(_t,_exprs)
  return fmt.Sprintf("%s %s == %s: %s",_if,_case,_cs,_s)
}

func if_topy(_t *Lint,_exprs string) string {
  _s := cond_block_topy(_t,_exprs)
  return fmt.Sprintf("if %s",_s)
}

func cond_block_topy(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(_args))
  _cond_str := atom_topy(_t,_cond)
  _exprs_str := block_topy(_t,_exprs)
  return fmt.Sprintf("%s: %s",_cond_str,_exprs_str)
}

func else_topy(_t *Lint,_exprs string) string {
  _s := block_topy(_t,_exprs)
  return fmt.Sprintf("else: %s",_s)
}

func while_topy(_t *Lint,_exprs string) string {
  _s := cond_block_topy(_t,_exprs)
  return fmt.Sprintf("while %s",_s)
}

func return_topy(_t *Lint,_atom string) string {
  _s := atom_topy(_t,_atom)
  return fmt.Sprintf("return %s",_s)
}

func say_topy(_t *Lint,_atom string) string {
  _s := atom_topy(_t,_atom)
  return fmt.Sprintf("say(%s)",_s)
}

func print_topy(_t *Lint,_atom string) string {
  _s := atom_topy(_t,_atom)
  return fmt.Sprintf("aprint(%s)",_s)
}

func error_topy(_t *Lint,_atom string) string {
  _s := atom_topy(_t,_atom)
  return fmt.Sprintf("error(%s)",_s)
}

func estr_topy(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _s := call_args_topy(_t,_args)
  switch (_type) {
    case `strs`:return fmt.Sprintf("estr(%s)",_s)
    default: return fmt.Sprintf("estr([%s])",_s)
  }
}

func my_topy(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _nstr := name_topy(_name)
  _vstr := atom_topy(_t,_value)
  return fmt.Sprintf("%s = %s",_nstr,_vstr)
}

func mys_topy(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _names := names(_vars)
  _len := itos(len(_names))
  _nstr := joinchar(mapstrs(_names,name_topy),',')
  _vstr := atom_topy(_t,_value)
  _type := get_atom_type(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("%s = get%s(%s)",_nstr,_len,_vstr)
  }
  return fmt.Sprintf("%s = get%s(atoms(%s))",_nstr,_len,_vstr)
}

func set_topy(_t *Lint,_args string) string {
  _arr := atoms_topys(_t,_args)
  _var,_value := flat2(_arr)
  return fmt.Sprintf("%s = %s",_var,_value)
}

func name_topy(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '-':_buf.WriteRune('_')
      case '$':_buf.WriteRune('_')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func dstr_topy(_dstr string) string {
  _s := dchars_topy(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func str_topy(_args string) string {
  _buf := new(bytes.Buffer)
  _vars := []string{}
  for _,_arg := range atoms(_args) {
    _type,_value := flat2(atoms(_arg))
    if _type == `dstr` {
      _buf.WriteString(dchars_topy(_value))
    } else {
      _buf.WriteString(`%s`)
      _vars = append(_vars,name_topy(_value))
    }
  }
  _s := _buf.String()
  _vars_str := joinchar(_vars,',')
  return fmt.Sprintf("\"%s\" %% (%s)",_s,_vars_str)
}

func dchars_topy(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '$':_buf.WriteRune(_a)
          case '-':_buf.WriteRune(_a)
          case '%':_buf.WriteString(`%%`)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: switch (_a) {
          case '\\':_mode = 1
          default: _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func const_topy(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return "{}"
  }
  return _name
}

func init_topy(_name string) string {
  switch (_name) {
    case `int`:return `0`
    case `str`:return "''"
    case `strs`:return "[]"
    case `buffer`:return "[]"
    case `table`:return "{}"
    case `tree`:return "{}"
    default: return "{}"
  }
}

func ncall_topy(_call string) string {
  _name := name_topy(name(_call))
  return fmt.Sprintf("%s()",_name)
}

func ocall_topy(_t *Lint,_args string) string {
  _call,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_topy(_t,_value)
  _name := name_topy(_call)
  if is_field(_t,_type,_call) {
    return fmt.Sprintf("%s['%s']",_s,_name)
  }
  return call_name_topy(_name,_s,_type)
}

func index_topy(_t *Lint,_args string) string {
  _data,_at := flat2(atoms_topys(_t,_args))
  return fmt.Sprintf("%s[%s]",_data,_at)
}

func pcall_topy(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _a,_b := flat2(atoms_topys(_t,_args))
  switch (_name) {
    case `<<`:return fmt.Sprintf("%s.append(%s)",_a,_b)
    case `>>`:return fmt.Sprintf("%s.insert(0,%s)",_b,_a)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func call_topy(_t *Lint,_atom string) string {
  _call,_args := flat2(atoms(_atom))
  _arr := atoms_topys(_t,_args)
  _s := joinchar(_arr,',')
  _type := get_args_type(_t,_args)
  _name := name_topy(_call)
  switch (_call) {
    case `map`:return map_topy(_arr)
    case `join`:return join_topy(_arr)
    case `split`:return split_topy(_arr)
    default: return call_name_topy(_name,_s,_type)
  }
}

func map_topy(_arr []string) string {
  _list,_fn := flat2(_arr)
  return fmt.Sprintf("list(map(%s,%s))",_fn,_list)
}

func join_topy(_arr []string) string {
  _list,_sep := flat2(_arr)
  return fmt.Sprintf("%s.join(%s)",_sep,_list)
}

func split_topy(_arr []string) string {
  _s,_sep := flat2(_arr)
  return fmt.Sprintf("%s.split(%s)",_s,_sep)
}

func call_name_topy(_name string,_s string,_type string) string {
  switch (_name) {
    case `tostr`:return tostr_topy(_s,_type)
    case `dec`:return fmt.Sprintf("%s -= 1",_s)
    case `inc`:return fmt.Sprintf("%s += 1",_s)
    case `not`:return fmt.Sprintf("not %s",_s)
    case `tochars`:return fmt.Sprintf("list(%s)",_s)
    case `next`:return fmt.Sprintf("second(%s)",_s)
    case `shift`:return fmt.Sprintf("%s.pop(0)",_s)
    case `copy`:return fmt.Sprintf("shutil.copy(%s)",_s)
    case `isfile`:return fmt.Sprintf("os.path.isfile(%s)",_s)
    case `filter`:return fmt.Sprintf("grep(%s)",_s)
    case `rename`:return fmt.Sprintf("os.rename(%s)",_s)
    case `isdigit`:return fmt.Sprintf("%s.isdigit()",_s)
    case `isalpha`:return fmt.Sprintf("%s.isalpha()",_s)
    case `isspace`:return fmt.Sprintf("%s.isspace()",_s)
    case `isupper`:return fmt.Sprintf("%s.isupper()",_s)
    case `islower`:return fmt.Sprintf("%s.islower()",_s)
    default: return fmt.Sprintf("%s(%s)",_name,_s)
  }
}

func tostr_topy(_s string,_type string) string {
  switch (_type) {
    case `char`:return _s
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    default: return fmt.Sprintf("string(%s)",_s)
  }
}

func ast_tojs(_t *Lint,_ast string) string {
  _arr := atoms_tojss(_t,_ast)
  return joinchar(_arr,'\n')
}

func atoms_tojss(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_tojs(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_tojs(_t *Lint,_atoms string) string {
  _arr := atoms_tojss(_t,_atoms)
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_tojs(_t *Lint,_expr string) string {
  _idt := get_indent(_t)
  _t.indent += 1
  _s := atom_tojs(_t,_expr)
  _indent := get_indent(_t)
  _t.indent -= 1
  return fmt.Sprintf("{\n%s%s\n%s}",_indent,_s,_idt)
}

func exprs_tojs(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_tojs(_t,_expr)
  _t.indent -= 1
  return _s
}

func call_args_tojs(_t *Lint,_args string) string {
  _arr := atoms_tojss(_t,_args)
  return joinchar(_arr,',')
}

func atom_tojs(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_tojs(_t,_args)
    case `fn`:return fn_tojs(_t,_args)
    case `for`:return for_tojs(_t,_args)
    case `case`:return case_tojs(_t,_args)
    case `of`:return of_tojs(_t,_args)
    case `then`:return then_tojs(_t,_args)
    case `if`:return if_tojs(_t,_args)
    case `ifelse`:return atoms_tojs(_t,_args)
    case `block`:return atoms_tojs(_t,_args)
    case `ofthen`:return atoms_tojs(_t,_args)
    case `else`:return else_tojs(_t,_args)
    case `while`:return while_tojs(_t,_args)
    case `set`:return set_tojs(_t,_args)
    case `my`:return my_tojs(_t,_args)
    case `mys`:return mys_tojs(_t,_args)
    case `return`:return return_tojs(_t,_args)
    case `say`:return say_tojs(_t,_args)
    case `print`:return print_tojs(_t,_args)
    case `error`:return error_tojs(_t,_args)
    case `estr`:return estr_tojs(_t,_args)
    case `const`:return const_tojs(_t,_args)
    case `index`:return index_tojs(_t,_args)
    case `init`:return init_tojs(_args)
    case `sym`:return name_tojs(_args)
    case `var`:return name_tojs(_args)
    case `str`:return str_tojs(_t,_args)
    case `dstr`:return dstr_tojs(_args)
    case `ncall`:return ncall_tojs(_args)
    case `ocall`:return ocall_tojs(_t,_args)
    case `call`:return call_tojs(_t,_args)
    case `pcall`:return pcall_tojs(_t,_args)
    case `lstr`:return fmt.Sprintf("`%s`",_args)
    case `kstr`:return fmt.Sprintf("'%s'",_args)
    case `int`:return _args
    case `bool`:return _args
    case `char`:return _args
    case `nil`:return "null"
    case `use`:return ``
    default: println(fmt.Sprintf("atom to Js miss %s",_name))
  }
  return ``
}

func ns_tojs(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_tojs(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_exprs := flat2(atoms(_atoms))
  _name_type := name(_head)
  in_ns(_t,_name_type)
  _name := name_tojs(name(_name_type))
  _args_str := fn_args_tojs(_args)
  _block := block_tojs(_t,_exprs)
  out_ns(_t)
  if _name == `main` {
    return fmt.Sprintf("do %s while (false)",_block)
  }
  return fmt.Sprintf("\nfunction %s(%s) %s",_name,_args_str,_block)
}

func fn_args_tojs(_args string) string {
  if _args == `nil` {
    return ``
  }
  _names := mapstrs(atoms(_args),name)
  _arr := mapstrs(_names,name_tojs)
  return joinchar(_arr,',')
}

func for_tojs(_t *Lint,_args string) string {
  _iter_atom,_expr := flat2(atoms(_args))
  _iter := iter_tojs(_t,_iter_atom)
  in_block(_t)
  _block := exprs_tojs(_t,_expr)
  out_block(_t)
  _idt := get_indent(_t)
  return fmt.Sprintf("for (%s) {%s\n%s}",_iter,_block,_idt)
}

func iter_tojs(_t *Lint,_atom string) string {
  _iter_name,_set_atom := flat2(atoms(_atom))
  _iter := name_tojs(_iter_name)
  _type := get_atom_type(_t,_set_atom)
  _set := atom_tojs(_t,_set_atom)
  switch (_type) {
    case `table`:return fmt.Sprintf("let %s of Object.keys(%s)",_iter,_set)
    case `tree`:return fmt.Sprintf("let %s of Object.keys(%s)",_iter,_set)
    default: return fmt.Sprintf("let %s of %s",_iter,_set)
  }
}

func case_tojs(_t *Lint,_args string) string {
  _s := cond_expr_tojs(_t,_args)
  return fmt.Sprintf("switch %s",_s)
}

func of_tojs(_t *Lint,_args string) string {
  _arr := atoms_tojss(_t,_args)
  _cond,_exprs := flat2(_arr)
  return fmt.Sprintf("case %s: %s; break;",_cond,_exprs)
}

func cond_tojs(_t *Lint,_args string) string {
  _arr := atoms_tojss(_t,_args)
  return join(_arr,": case ")
}

func then_tojs(_t *Lint,_exprs string) string {
  _s := exprs_tojs(_t,_exprs)
  return fmt.Sprintf("default: %s",_s)
}

func if_tojs(_t *Lint,_args string) string {
  _s := cond_expr_tojs(_t,_args)
  return fmt.Sprintf("if %s",_s)
}

func cond_expr_tojs(_t *Lint,_args string) string {
  _cond_atom,_expr := flat2(atoms(_args))
  _cond := atom_tojs(_t,_cond_atom)
  _block := block_tojs(_t,_expr)
  return fmt.Sprintf("(%s) %s",_cond,_block)
}

func else_tojs(_t *Lint,_exprs string) string {
  _s := block_tojs(_t,_exprs)
  return fmt.Sprintf("else %s",_s)
}

func while_tojs(_t *Lint,_args string) string {
  _s := cond_expr_tojs(_t,_args)
  return fmt.Sprintf("while %s",_s)
}

func set_tojs(_t *Lint,_args string) string {
  _arr := atoms_tojss(_t,_args)
  _var,_value := flat2(_arr)
  return fmt.Sprintf("%s = %s;",_var,_value)
}

func my_tojs(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _name := name_tojs(_var)
  _value_str := atom_tojs(_t,_value)
  return fmt.Sprintf("let %s = %s;",_name,_value_str)
}

func mys_tojs(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _names := mapstrs(names(_vars),name_tojs)
  _vars_str := joinchar(_names,',')
  _type := get_atom_type(_t,_value)
  _vstr := atom_tojs(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("let [%s] = %s",_vars_str,_vstr)
  }
  return fmt.Sprintf("let [%s] = atoms(%s)",_vars_str,_vstr)
}

func return_tojs(_t *Lint,_atom string) string {
  _s := atom_tojs(_t,_atom)
  return fmt.Sprintf("return %s",_s)
}

func say_tojs(_t *Lint,_atom string) string {
  _s := atom_tojs(_t,_atom)
  return fmt.Sprintf("say(%s);",_s)
}

func print_tojs(_t *Lint,_atom string) string {
  _s := atom_tojs(_t,_atom)
  return fmt.Sprintf("print(%s);",_s)
}

func error_tojs(_t *Lint,_atom string) string {
  _s := atom_tojs(_t,_atom)
  return fmt.Sprintf("error(%s);",_s)
}

func estr_tojs(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _s := call_args_tojs(_t,_args)
  switch (_type) {
    case `strs`:return fmt.Sprintf("estr(%s)",_s)
    default: return fmt.Sprintf("estr([%s])",_s)
  }
}

func const_tojs(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return "{}"
  }
  return _name
}

func init_tojs(_name string) string {
  switch (_name) {
    case `int`:return "0"
    case `str`:return "''"
    case `buffer`:return "[]"
    case `strs`:return "[]"
    case `table`:return "{}"
    case `tree`:return "{}"
    default: return "{}"
  }
}

func name_tojs(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '-':_buf.WriteRune('_')
      case '$':_buf.WriteRune('s')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func str_tojs(_t *Lint,_args string) string {
  _buf := new(bytes.Buffer)
  for _,_atom := range atoms(_args) {
    _buf.WriteString(satom_tojs(_atom))
  }
  _format := _buf.String()
  return fmt.Sprintf("`%s`",_format)
}

func satom_tojs(_atom string) string {
  _name,_value := flat2(atoms(_atom))
  switch (_name) {
    case `var`:_var := name_tojs(_value)
      return fmt.Sprintf("${%s}",_var)
    default: return dchars_tojs(_value)
  }
}

func dstr_tojs(_dstr string) string {
  _s := dchars_tojs(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func dchars_tojs(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '-':_buf.WriteRune(_a)
          case '[':_buf.WriteRune(_a)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: switch (_a) {
          case '\\':_mode = 1
          case '`':_buf.WriteString(addcc('\\',_a))
          default: _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func ncall_tojs(_call string) string {
  _name := name_tojs(name(_call))
  return fmt.Sprintf("%s()",_name)
}

func ocall_tojs(_t *Lint,_args string) string {
  _call,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_tojs(_t,_value)
  _name := name_tojs(_call)
  if is_field(_t,_type,_call) {
    return fmt.Sprintf("%s.%s",_s,_name)
  }
  return call_name_tojs(_name,_s,_type)
}

func call_tojs(_t *Lint,_atom string) string {
  _call,_args := flat2(atoms(_atom))
  _name := name_tojs(_call)
  _type := get_args_type(_t,_args)
  _s := call_args_tojs(_t,_args)
  return call_name_tojs(_name,_s,_type)
}

func index_tojs(_t *Lint,_args string) string {
  _a,_b := flat2(atoms_tojss(_t,_args))
  return fmt.Sprintf("%s[%s]",_a,_b)
}

func pcall_tojs(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _a,_b := flat2(atoms_tojss(_t,_args))
  switch (_name) {
    case `<<`:return fmt.Sprintf("%s.push(%s)",_a,_b)
    case `>>`:return fmt.Sprintf("%s.unshift(%s)",_b,_a)
    case `==`:return fmt.Sprintf("%s === %s",_a,_b)
    case `!=`:return fmt.Sprintf("%s !== %s",_a,_b)
    case `+=`:return fmt.Sprintf("%s += %s;",_a,_b)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func call_name_tojs(_name string,_s string,_type string) string {
  switch (_name) {
    case `tostr`:return tostr_tojs(_s,_type)
    case `has`:return has_tojs(_s,_type)
    case `rest`:return rest_tojs(_s,_type)
    case `first`:return fmt.Sprintf("%s[0]",_s)
    case `second`:return fmt.Sprintf("%s[1]",_s)
    case `dec`:return fmt.Sprintf("%s --;",_s)
    case `inc`:return fmt.Sprintf("%s++;",_s)
    case `not`:return fmt.Sprintf("!%s",_s)
    case `len`:return fmt.Sprintf("%s.length",_s)
    case `shift`:return fmt.Sprintf("%s.shift()",_s)
    case `toint`:return fmt.Sprintf("parseInt(%s)",_s)
    default: return fmt.Sprintf("%s(%s)",_name,_s)
  }
}

func tostr_tojs(_s string,_type string) string {
  switch (_type) {
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    case `char`:return _s
    default: return fmt.Sprintf("tostr(%s)",_s)
  }
}

func has_tojs(_s string,_type string) string {
  switch (_type) {
    case `str:char`:return fmt.Sprintf("hasStr(%s)",_s)
    case `strs:str`:return fmt.Sprintf("hasStr(%s)",_s)
    default: return fmt.Sprintf("hasKey(%s)",_s)
  }
}

func rest_tojs(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("rest(%s)",_s)
    default: return fmt.Sprintf("reststrs(%s)",_s)
  }
}

func ast_tophp(_t *Lint,_ast string) string {
  _arr := atoms_tophps(_t,_ast)
  _s := joinchar(_arr,'\n')
  return fmt.Sprintf("<?php\n%s\n?>",_s)
}

func atoms_tophps(_t *Lint,_atoms string) []string {
  _arr := []string{}
  for _,_atom := range atoms(_atoms) {
    _s := atom_tophp(_t,_atom)
    if _s != `` {
      _arr = append(_arr,_s)
    }
  }
  return _arr
}

func atoms_tophp(_t *Lint,_atoms string) string {
  _arr := atoms_tophps(_t,_atoms)
  return atom_strs_tophp(_t,_arr)
}

func atom_strs_tophp(_t *Lint,_arr []string) string {
  _idt := get_indent(_t)
  _sep := addcs('\n',_idt)
  return join(_arr,_sep)
}

func block_tophp(_t *Lint,_expr string) string {
  _idt := get_indent(_t)
  _t.indent += 1
  _s := atom_tophp(_t,_expr)
  _inc_idt := get_indent(_t)
  _t.indent -= 1
  return fmt.Sprintf("{\n%s%s\n%s}",_inc_idt,_s,_idt)
}

func exprs_tophp(_t *Lint,_expr string) string {
  _t.indent += 1
  _s := atom_tophp(_t,_expr)
  _t.indent -= 1
  return _s
}

func atom_tophp(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  switch (_name) {
    case `ns`:return ns_tophp(_t,_args)
    case `fn`:return fn_tophp(_t,_args)
    case `for`:return for_tophp(_t,_args)
    case `case`:return case_tophp(_t,_args)
    case `of`:return of_tophp(_t,_args)
    case `then`:return then_tophp(_t,_args)
    case `if`:return if_tophp(_t,_args)
    case `ifelse`:return atoms_tophp(_t,_args)
    case `block`:return atoms_tophp(_t,_args)
    case `ofthen`:return atoms_tophp(_t,_args)
    case `else`:return else_tophp(_t,_args)
    case `while`:return while_tophp(_t,_args)
    case `set`:return set_tophp(_t,_args)
    case `my`:return my_tophp(_t,_args)
    case `mys`:return mys_tophp(_t,_args)
    case `return`:return return_tophp(_t,_args)
    case `say`:return say_tophp(_t,_args)
    case `print`:return print_tophp(_t,_args)
    case `error`:return error_tophp(_t,_args)
    case `estr`:return estr_tophp(_t,_args)
    case `pcall`:return pcall_tophp(_t,_args)
    case `ocall`:return ocall_tophp(_t,_args)
    case `call`:return call_tophp(_t,_args)
    case `index`:return index_tophp(_t,_args)
    case `ncall`:return ncall_tophp(_t,_args)
    case `sym`:return name_tophp(_args)
    case `var`:return name_tophp(_args)
    case `init`:return init_tophp(_args)
    case `const`:return const_tophp(_t,_args)
    case `str`:return str_tophp(_args)
    case `dstr`:return dstr_tophp(_args)
    case `lstr`:return fmt.Sprintf("<<<'STR'\n%s\nSTR",_args)
    case `int`:return _args
    case `bool`:return _args
    case `char`:return char_tophp(_args)
    case `kstr`:return fmt.Sprintf("'%s'",_args)
    case `nil`:return `null`
    case `use`:return ``
    default: println(fmt.Sprintf("atom tophp miss %s",_name))
  }
  return ``
}

func char_tophp(_str string) string {
  if len(_str) == 3 {
    return _str
  }
  _char := cut(_str)
  return fmt.Sprintf("\"%s\"",_char)
}

func ns_tophp(_t *Lint,_ns string) string {
  in_ns(_t,_ns)
  return ``
}

func fn_tophp(_t *Lint,_ast string) string {
  _head,_atoms := flat2(atoms(_ast))
  _args,_block := flat2(atoms(_atoms))
  _name_type := name(_head)
  in_ns(_t,_name_type)
  _name := name_tophp(name(_name_type))
  _args_str := fn_args_tophp(_t,_args)
  _block_str := block_tophp(_t,_block)
  out_ns(_t)
  if _name == `main` {
    return fmt.Sprintf("do %s while (False);",_block_str)
  }
  _declare := fmt.Sprintf("function %s(%s)",_name,_args_str)
  return fmt.Sprintf("\n%s %s",_declare,_block_str)
}

func fn_args_tophp(_t *Lint,_args string) string {
  if _args == `nil` {
    return ``
  }
  _arr := mapstrs(atoms(_args),fn_arg_tophp)
  return joinchar(_arr,',')
}

func fn_arg_tophp(_arg string) string {
  _var,_type := flat2(atoms(_arg))
  _name := name_tophp(_var)
  switch (_type) {
    case `table`:return fmt.Sprintf("&%s",_name)
    case `tree`:return fmt.Sprintf("&%s",_name)
    case `Lint`:return fmt.Sprintf("&%s",_name)
    case `Cursor`:return fmt.Sprintf("&%s",_name)
    default: return _name
  }
}

func for_tophp(_t *Lint,_args string) string {
  _iter,_block := flat2(atoms(_args))
  _iter_str := iter_tophp(_t,_iter)
  in_block(_t)
  _block_str := block_tophp(_t,_block)
  out_block(_t)
  return fmt.Sprintf("foreach (%s) %s",_iter_str,_block_str)
}

func iter_tophp(_t *Lint,_atom string) string {
  _iter_name,_set_atom := flat2(atoms(_atom))
  _type := get_atom_type(_t,_set_atom)
  _iter := name_tophp(_iter_name)
  _set := atom_tophp(_t,_set_atom)
  switch (_type) {
    case `strs`:return fmt.Sprintf("%s as %s",_set,_iter)
    case `str`:return fmt.Sprintf("tochars(%s) as %s",_set,_iter)
    default: return fmt.Sprintf("array_keys(%s) as %s",_set,_iter)
  }
}

func case_tophp(_t *Lint,_args string) string {
  _s := cond_exprs_tophp(_t,_args)
  return fmt.Sprintf("switch %s",_s)
}

func of_tophp(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(_args))
  _cond_str := atom_tophp(_t,_cond)
  _block := exprs_tophp(_t,_exprs)
  return fmt.Sprintf("case %s: %s; break;",_cond_str,_block)
}

func then_tophp(_t *Lint,_args string) string {
  _s := exprs_tophp(_t,_args)
  return fmt.Sprintf("default: %s;",_s)
}

func if_tophp(_t *Lint,_exprs string) string {
  _s := cond_exprs_tophp(_t,_exprs)
  return fmt.Sprintf("if %s",_s)
}

func else_tophp(_t *Lint,_exprs string) string {
  _s := block_tophp(_t,_exprs)
  return fmt.Sprintf("else %s",_s)
}

func cond_exprs_tophp(_t *Lint,_args string) string {
  _cond,_exprs := flat2(atoms(_args))
  _cond_str := atom_tophp(_t,_cond)
  _exprs_str := block_tophp(_t,_exprs)
  return fmt.Sprintf("(%s) %s",_cond_str,_exprs_str)
}

func while_tophp(_t *Lint,_args string) string {
  _s := cond_exprs_tophp(_t,_args)
  return fmt.Sprintf("while %s",_s)
}

func set_tophp(_t *Lint,_args string) string {
  _arr := atoms_tophps(_t,_args)
  _a,_b := flat2(_arr)
  return fmt.Sprintf("%s = %s;",_a,_b)
}

func my_tophp(_t *Lint,_args string) string {
  _var,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _vstr := atom_tophp(_t,_value)
  _name := name_tophp(_var)
  if _vstr == "[]" {
    return fmt.Sprintf("%s = %s;",_name,_vstr)
  }
  switch (_type) {
    case `tree`:return fmt.Sprintf("%s = &%s;",_name,_vstr)
    case `table`:return fmt.Sprintf("%s = &%s;",_name,_vstr)
    default: return fmt.Sprintf("%s = %s;",_name,_vstr)
  }
}

func mys_tophp(_t *Lint,_args string) string {
  _vars,_value := flat2(atoms(_args))
  _names := mapstrs(names(_vars),name_tophp)
  _nstr := joinchar(_names,',')
  _vstr := atom_tophp(_t,_value)
  _type := get_atom_type(_t,_value)
  if _type == `strs` {
    return fmt.Sprintf("list(%s) = %s;",_nstr,_vstr)
  }
  return fmt.Sprintf("list(%s) = atoms(%s);",_nstr,_vstr)
}

func return_tophp(_t *Lint,_atom string) string {
  _s := atom_tophp(_t,_atom)
  return fmt.Sprintf("return %s;",_s)
}

func say_tophp(_t *Lint,_atom string) string {
  _s := atom_tophp(_t,_atom)
  return fmt.Sprintf("say(%s);",_s)
}

func print_tophp(_t *Lint,_atom string) string {
  _s := atom_tophp(_t,_atom)
  return fmt.Sprintf("print(%s);",_s)
}

func error_tophp(_t *Lint,_atom string) string {
  _s := atom_tophp(_t,_atom)
  return fmt.Sprintf("error(%s);",_s)
}

func estr_tophp(_t *Lint,_args string) string {
  _type := get_args_type(_t,_args)
  _arr := atoms_tophps(_t,_args)
  _s := joinchar(_arr,',')
  if _type == `strs` {
    return fmt.Sprintf("estr(%s)",_s)
  }
  return fmt.Sprintf("estr(array(%s))",_s)
}

func init_tophp(_name string) string {
  switch (_name) {
    case `int`:return `0`
    case `str`:return "''"
    default: return "[]"
  }
}

func const_tophp(_t *Lint,_name string) string {
  if is_struct(_t,_name) {
    return "[]"
  }
  return _name
}

func name_tophp(_name string) string {
  _buf := new(bytes.Buffer)
  for _,_a := range []char(_name) {
    switch (_a) {
      case '-':_buf.WriteRune('_')
      default: _buf.WriteRune(_a)
    }
  }
  return _buf.String()
}

func str_tophp(_args string) string {
  _arr := mapstrs(atoms(_args),satom_tophp)
  _s := tostr(_arr)
  return fmt.Sprintf("\"%s\"",_s)
}

func satom_tophp(_atom string) string {
  _type,_value := flat2(atoms(_atom))
  if _type == `dstr` {
    return dchars_tophp(_value)
  }
  return name_tophp(_value)
}

func dstr_tophp(_dstr string) string {
  _s := dchars_tophp(_dstr)
  return fmt.Sprintf("\"%s\"",_s)
}

func dchars_tophp(_s string) string {
  _buf := new(bytes.Buffer)
  _mode := 0
  for _,_a := range []char(_s) {
    switch (_mode) {
      case 1:_mode = 0
        switch (_a) {
          case '-':_buf.WriteRune(_a)
          case '$':_buf.WriteRune(_a)
          case '[':_buf.WriteRune(_a)
          case '{':_buf.WriteRune(_a)
          case '%':_buf.WriteString(`%%`)
          default: _buf.WriteString(addcc('\\',_a))
        }
      default: switch (_a) {
          case '\\':_mode = 1
          case '[':_buf.WriteString("\\[")
          case '{':_buf.WriteString("\\{")
          default: _buf.WriteRune(_a)
        }
    }
  }
  return _buf.String()
}

func ncall_tophp(_t *Lint,_call string) string {
  _type := get_ncall_type(_t,_call)
  _name := name_tophp(name(_call))
  if _name == `osargs` {
    return "array_slice($argv,1)"
  }
  if _type == `nil` {
    return fmt.Sprintf("%s();",_name)
  }
  return fmt.Sprintf("%s()",_name)
}

func ocall_tophp(_t *Lint,_args string) string {
  _name,_value := flat2(atoms(_args))
  _type := get_atom_type(_t,_value)
  _s := atom_tophp(_t,_value)
  if is_field(_t,_type,_name) {
    return fmt.Sprintf("%s['%s']",_s,_name)
  }
  return call_name_tophp(_t,_name,_s,_type)
}

func index_tophp(_t *Lint,_args string) string {
  _data,_at := flat2(atoms_tophps(_t,_args))
  return fmt.Sprintf("%s[%s]",_data,_at)
}

func pcall_tophp(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _a,_b := flat2(atoms_tophps(_t,_args))
  switch (_name) {
    case `+`:return add_tophp(_a,_b,_type)
    case `<<`:return fmt.Sprintf("array_push(%s,%s);",_a,_b)
    case `>>`:return fmt.Sprintf("array_unshift(%s,%s);",_b,_a)
    case `+=`:return fmt.Sprintf("%s += %s;",_a,_b)
    default: return fmt.Sprintf("%s %s %s",_a,_name,_b)
  }
}

func add_tophp(_a string,_b string,_type string) string {
  switch (_type) {
    case `int:int`:return fmt.Sprintf("%s + %s",_a,_b)
    default: return fmt.Sprintf("%s . %s",_a,_b)
  }
}

func call_tophp(_t *Lint,_atom string) string {
  _name,_args := flat2(atoms(_atom))
  _type := get_args_type(_t,_args)
  _arr := atoms_tophps(_t,_args)
  _s := joinchar(_arr,',')
  if _name == `map` {
    return map_tophp(_arr)
  }
  return call_name_tophp(_t,_name,_s,_type)
}

func map_tophp(_arr []string) string {
  _array,_fn := flat2(_arr)
  return fmt.Sprintf("array_map('%s',%s)",_fn,_array)
}

func call_name_tophp(_t *Lint,_name string,_s string,_type string) string {
  _name_type := fmt.Sprintf("%s.%s",_name,_type)
  _ctype := get_name_type(_t,_name_type)
  _pname := name_tophp(_name)
  switch (_name) {
    case `len`:return len_tophp(_s,_type)
    case `tostr`:return tostr_tophp(_s,_type)
    case `first`:return first_tophp(_s,_type)
    case `rest`:return rest_tophp(_s,_type)
    case `has`:return has_tophp(_s,_type)
    case `shift`:return fmt.Sprintf("array_shift(%s);",_s)
    case `dec`:return fmt.Sprintf("--%s;",_s)
    case `inc`:return fmt.Sprintf("%s++;",_s)
    case `not`:return fmt.Sprintf("!%s",_s)
    case `toint`:return fmt.Sprintf("(int)%s",_s)
    case `join`:return fmt.Sprintf("implode(%s)",_s)
    case `readfile`:return fmt.Sprintf("read_file(%s)",_s)
    case `next`:return fmt.Sprintf("second(%s)",_s)
    case `chop`:return fmt.Sprintf("pchop(%s)",_s)
    case `lower`:return fmt.Sprintf("strtolower(%s)",_s)
    case `upper`:return fmt.Sprintf("strtoupper(%s)",_s)
    case `range`:return fmt.Sprintf("rangestr(%s)",_s)
    case `isdigit`:return fmt.Sprintf("ctype_digit(%s)",_s)
    case `isspace`:return fmt.Sprintf("ctype_space(%s)",_s)
    case `isalpha`:return fmt.Sprintf("ctype_alpha(%s)",_s)
    case `isupper`:return fmt.Sprintf("ctype_upper(%s)",_s)
    case `islower`:return fmt.Sprintf("ctype_lower(%s)",_s)
    case `isxdigit`:return fmt.Sprintf("ctype_xdigit(%s)",_s)
    case `isfile`:return fmt.Sprintf("is_file(%s)",_s)
    case `repeat`:return fmt.Sprintf("str_repeat(%s)",_s)
    default: return pname_tophp(_pname,_s,_ctype)
  }
}

func pname_tophp(_pname string,_s string,_ctype string) string {
  if _ctype == `nil` {
    return fmt.Sprintf("%s(%s);",_pname,_s)
  }
  return fmt.Sprintf("%s(%s)",_pname,_s)
}

func len_tophp(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("strlen(%s)",_s)
    default: return fmt.Sprintf("count(%s)",_s)
  }
}

func tostr_tophp(_s string,_type string) string {
  switch (_type) {
    case `int`:return fmt.Sprintf("itos(%s)",_s)
    case `char`:return _s
    default: return fmt.Sprintf("tostr(%s)",_s)
  }
}

func first_tophp(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("first(%s)",_s)
    default: return fmt.Sprintf("(%s)[0]",_s)
  }
}

func rest_tophp(_s string,_type string) string {
  switch (_type) {
    case `str`:return fmt.Sprintf("rest(%s)",_s)
    default: return fmt.Sprintf("restStrs(%s)",_s)
  }
}

func has_tophp(_s string,_type string) string {
  switch (_type) {
    case `str:char`:return fmt.Sprintf("find(%s)",_s)
    case `str:str`:return fmt.Sprintf("find(%s)",_s)
    default: return fmt.Sprintf("has(%s)",_s)
  }
}

func make_spp(_lang string)  {
  _files := estr([]string{`Core`,`Estr`,`SppAst`,`Match`,`OptSppMatch`,`OptMyMatch`,`GetTable`,`LintMyAst`,`ToMy`,`ToGo`,`ToC`,`ToPerl`,`ToLua`,`ToRuby`,`ToPython`,`ToJs`,`ToPHP`,`Main`})
  compiler_files(_files)
  gather_files(_files,_lang,`spp`)
}

func compiler_files(_files string)  {
  _t := get_table(`my`)
  for _,_file := range atoms(_files) {
    compiler_file(_t,_file)
  }
}

func compiler_file(_table table,_file string)  {
  _myfile := fmt.Sprintf("my/%s.my",_file)
  println(fmt.Sprintf("compiler %s .. ",_myfile))
  _code := readfile(_myfile)
  _match := match_table(_code,_table)
  _ast := opt_my_match(_match)
  _ofile := fmt.Sprintf("to/o/%s.o",_file)
  writefile(_ofile,_ast)
  println(fmt.Sprintf("-> %s ",_ofile))
  _t := lint_my_ast(_ast)
  ast_to_lang(_t,_ast,_file,`go`)
  ast_to_lang(_t,_ast,_file,`c`)
  ast_to_lang(_t,_ast,_file,`js`)
  ast_to_lang(_t,_ast,_file,`pl`)
  ast_to_lang(_t,_ast,_file,`rb`)
  ast_to_lang(_t,_ast,_file,`php`)
  ast_to_lang(_t,_ast,_file,`lua`)
  ast_to_lang(_t,_ast,_file,`py`)
  ast_to_lang(_t,_ast,_file,`my`)
}

func ast_to_lang(_t *Lint,_ast string,_file string,_lang string)  {
  if _file == `Core` {
    return 
  }
  reset_block(_t)
  _lang_file := fmt.Sprintf("to/%s/%s.%s",_lang,_file,_lang)
  _code := ``
  switch (_lang) {
    case `go`:_code = ast_togo(_t,_ast)
    case `c`:_code = ast_toc(_t,_ast)
    case `js`:_code = ast_tojs(_t,_ast)
    case `pl`:_code = ast_topl(_t,_ast)
    case `rb`:_code = ast_torb(_t,_ast)
    case `php`:_code = ast_tophp(_t,_ast)
    case `lua`:_code = ast_tolua(_t,_ast)
    case `py`:_code = ast_topy(_t,_ast)
    case `my`:_code = ast_tomy(_t,_ast)
    default: error(fmt.Sprintf("ast to Lang miss %s",_lang))
  }
  writefile(_lang_file,_code)
  println(fmt.Sprintf("-> %s ",_lang_file))
}

func gather_files(_files string,_lang string,_item string)  {
  _arr := []string{}
  println(fmt.Sprintf("gather %s.%s require files .. ",_item,_lang))
  for _,_name := range atoms(_files) {
    _arr = append(_arr,get_file_str(_name,_lang))
  }
  _file := fmt.Sprintf("%s-%s.%s",_lang,_item,_lang)
  _time := now()
  _bakfile := fmt.Sprintf("bak/%s#%s.%s",_item,_time,_lang)
  if isfile(_file) {
    if _lang == `go` {
      rename(_file,_bakfile)
    }
  }
  _text := joinchar(_arr,'\n')
  writefile(_file,_text)
  println(fmt.Sprintf("create file %s",_file))
}

func get_file_str(_name string,_lang string) string {
  _file := get_lang_file(_lang,_name)
  return readfile(_file)
}

func get_lang_file(_lang string,_file string) string {
  if _file == `Core` {
    return fmt.Sprintf("core/Core.%s",_lang)
  }
  return fmt.Sprintf("to/%s/%s.%s",_lang,_file,_lang)
}

func help() string {
  return `
  This is MySpp
  Copyright  2018-2020 Songzhiquan

  Usage:
  >> go run go-spp.go [help]
  >> go run go-spp.go make [go]
  >> go run go-spp.go spp [top]
  >> go run go-spp.go my [top]
  >> go run go-spp.go grammar [top]
`
}

func main() {
  _osargs := osargs()
  _len := len(_osargs)
  if _len == 0 {
    print(help())
  }
  if _len == 1 {
    _name := _osargs[0]
    switch (_name) {
      case `help`:print(help())
      case `make`:make_spp(`go`)
      case `spp`:spp_repl(`top`)
      case `my`:my_repl(`top`)
      case `ast`:print(get_ast("spp"))
      default: lang_repl(_name,`top`)
    }
  }
  if _len == 2 {
    _lang,_door := flat2(_osargs)
    switch (_lang) {
      case `make`:make_spp(_door)
      case `spp`:spp_repl(_door)
      case `my`:my_repl(_door)
      case `ast`:print(get_ast(_door))
      default: lang_repl(_lang,_door)
    }
  }
}