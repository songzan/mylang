#include <assert.h>
#include <ctype.h>
#include <stdarg.h>
#include <stdbool.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define  FOF   '\0'
#define  EIN   '\1'
#define  OUT   '\2'
#define  QSTR  '\3'
#define  FAIL  "\0"
#define  PASS  "\1"

void say(char* s) {
  printf("%s\n", s);
  fflush(stdout); 
}

void print(char* s) {
  printf("%s", s);
  fflush(stdout); 
}

void error(char* s) {
  printf("%s", s); exit(1);
}

bool eq(char* s1, char* s2) {
  return (strcmp(s1, s2) == 0);
}

bool ne(char* s1, char* s2) {
  return (strcmp(s1, s2) != 0);
}

bool hasChar(char* s, char ch) {
  return (strchr(s,ch) != NULL);
}

bool startwith(char* s, char c) {
  if (s[0] == c) return true;
  return false;
}

char tail(char* s) {
  return s[strlen(s)-1]; 
}

bool endwith(char* s, char c) {
  char lc = tail(s);
  if (lc == c) return true;
  return false;
}

void* apply_mem(int size) {
  void *ptr = malloc(size);
  if (ptr == NULL)
    error("apply memory fail!");
  return ptr;
}

char* upper(char *s) {
  int n = strlen(s);
  int i = 0;
  char *us = apply_mem(n+1);
  while (i < n) {
    us[i] = toupper(s[i]);
  }
  us[i] = '\0';
  return us;
}

char* lower(char *s) {
  int n = strlen(s);
  int i = 0;
  char *us = apply_mem(n+1);
  while (i < n) {
    us[i] = tolower(s[i]);
  }
  us[i] = '\0';
  return us;
}

char* charToStr(char c) {
  char* s = apply_mem(2);
  s[0] = c; s[1] = '\0'; return s;
}

char* charAddChar(char a, char b) {
  char* cc = apply_mem(3);
  cc[0] = a; cc[1] = b; cc[2] = '\0';
  return cc;
}

char* charAddStr(char ch, char* s) {
  int n = strlen(s);
  char* ns = apply_mem(n+2);
  ns[0] = ch;
  int i = 0;
  while (i < n) { ns[i+1] = s[i]; i++; }
  ns[n+1] = '\0';
  return ns;
}

char* strAddChar(char* s, char ch) {
  int n = strlen(s);
  char* ns = apply_mem(n+2);
  int i = 0;
  while (i < n) { ns[i] = s[i]; i++; }
  ns[n] = ch;
  ns[n+1] = '\0';
  return ns;
}

char* strAddStr(char* s1, char* s2) {
  int n1 = strlen(s1);
  int n2 = strlen(s2);
  char* ss = apply_mem(n1+n2+1);
  int i = 0; int m = 0;
  while (i < n1) { ss[i] = s1[i]; i++; }
  while (m < n2) { ss[i] = s2[m]; i++; m++; }
  ss[i] = '\0';
  return ss;
}

char* substr(char* s, int m, int n) { 
  if (n <= 0) {
    n = strlen(s) + n;
    return substr(s,m,n);
  }
  int size = n - m + 1;
  char* ns = apply_mem(size);
  int i = 0;
  while (m < n) {
    ns[i] = s[m]; m++; i++;
  }
  ns[i] = '\0';
  return ns;
}

char* chop(char* s) {
  return substr(s,0,-1);
}

char* cut(char* s) {
  return substr(s, 1, -1);
}

char* rest(char* s) {
  return substr(s, 1, 0);
}

char* restat(char* s, int at) {
  return substr(s, at, 0);
}

char* toend(char* s) {
  if (strlen(s) < 40) return s;
  return substr(s,0,40);
}

char* intToStr(int n) {
  n = n % 100000;
  char* s = apply_mem(5);
  int count = 5;
  do {
    count--;
    char c = (char)((n%10)+48);
    s[count] = c;
    n = n/10; 
  } while (n > 0);
  return substr(s,count,5);
}

int count(char* s, char ch) {
  int n = strlen(s);
  int c = 0;
  int i = 0;
  while (i < n) {
    if (s[i] == ch) { c++; }
    i++;
  }
  return c;
}

char* getline(char* cs,int off) {
  char* s = substr(cs, 0, off);
  int c = count(s,'\n') + 1;
  return intToStr(c);
}

char* repeat(char* s, int tn) {
  int sn = strlen(s);
  char* rs = apply_mem(sn*tn+1);
  int ri = 0;
  int m, i;
  for (m = 0; m < tn; m++) {
    for (i = 0; i < sn; i++) {
      rs[ri] = s[i]; ri++;
    }
  }
  rs[ri] = '\0';
  return rs;
}

char* trim(char* s) {
  int rn = 0;
  while (isspace(s[rn])) { rn++; }
  int tn = strlen(s)-1;
  while (isspace(s[tn])) { tn--; }
  return substr(s,rn,tn+1);
}

char* copyStr(char* s) {
  int size = strlen(s) + 1;
  char* cs = apply_mem(size);
  int i = 0;
  while (i < size) {
    cs[i] = s[i]; i++;
  }
  return cs;
}

struct snode {
  char* str;
  struct snode *next;
};

typedef struct snode Snode;

Snode *newSnode(void) {
  int size = sizeof(Snode);
  Snode *nnode = apply_mem(size);
  nnode->next = NULL;
  return nnode;
}

struct strs {
  Snode *head;
  Snode *last;
  int len;
};

typedef struct strs Strs;

Strs *newStrs(void) {
  int size = sizeof(Strs); 
  Strs *ss = apply_mem(size);
  ss->head = NULL;
  ss->len = 0;
  return ss;
}

void push(Strs *ss, char* s) {
  Snode *nnode = newSnode();
  nnode->str = s;
  if (ss->len == 0) {
    ss->last = nnode;
    ss->head = nnode;
  } else {
    ss->last->next = nnode;
    ss->last = nnode;
  }
  ss->len++;  
}

void insert(Strs* ss, char* str) {
  Snode *nnode = newSnode();
  nnode->str = str;
  if (ss->len == 0) {
    ss->head = nnode;
    ss->last = nnode;
  } else {
    nnode->next = ss->head;
    ss->head = nnode;
  }
  ss->len++;
}

void shift(Strs *ss) {
  Snode *fnode = ss->head;
  ss->head = fnode->next;
  ss->len--;
  free(fnode);
}

int getIndex(Strs *ss, char* key) {
  int idx = 0;
  Snode *fnode = ss->head;
  while (fnode != NULL) {
    char* s = fnode->str;
    if (eq(s,key)) break;
    fnode = fnode->next;
    idx++;
  }
  if (idx == ss->len) { return -1; }
  return idx;
}

char* getStr(Strs *ss, int idx) {
  Snode *fnode = ss->head;
  while (fnode != NULL) {
    if (idx == 0) break;
    fnode = fnode->next;
    idx--;
  }
  return fnode->str;
}

char* second(Strs *ss) {
  return getStr(ss,1);
}

char* first(Strs *ss) {
  return ss->head->str; 
}

Strs *restStrs(Strs *ss) {
  Strs* mss = newStrs();
  Snode* node = ss->head->next;
  while (node != NULL) {
    push(mss, node->str);
    node = node->next;
  }
  return mss;
}

int strslen(Strs *ss) { return ss->len; }

Strs *tostrs(int amount, ...) {
  Strs *ss = newStrs();
  va_list vl;
  va_start(vl,amount);
  int i = 0;
  while (i < amount) {
    push(ss,va_arg(vl, char*)); i++;
  }
  va_end(vl);
  return ss;
}

struct buffer {
  int off;
  int len;
  char* str;
  Strs *arr;
};

typedef struct buffer Buffer;

Buffer *newBuffer() {
  int size = sizeof(Buffer);
  Buffer *buf = apply_mem(size);
  buf->str = (char*)apply_mem(64);
  buf->off = 0;
  buf->len = 0;
  buf->arr = newStrs();
  return buf;
}

void pushChar(Buffer *buf, char ch) {
  buf->str[buf->off] = ch;
  buf->off++; buf->len++;
  if (buf->off == 64) {
    push(buf->arr, buf->str);
    buf->str = (char*)apply_mem(64);
    buf->off = 0;
  }
}

void pushStr(Buffer *buf, char* s) {
  int n = strlen(s);
  int i = 0;
  while (i < n) {
    pushChar(buf, s[i]); i++;
  }
}

char* bufferToStr(Buffer *buf) {
  char* str = apply_mem(buf->len + 1);
  int n = buf->arr->len;
  int si = 0; int i = 0; int m = 0;
  while (i < n) {
    char* s = first(buf->arr); i++;
    shift(buf->arr);
    for (m = 0; m < 64; m++) {
      str[si] = s[m]; si++;
    }
    free(s);
  }
  n = buf->off;
  char* cs = buf->str;
  for (i = 0; i < n; i++) {
    str[si] = cs[i]; si++;
  }
  str[si] = '\0';
  free(buf->str); free(buf); return str;
}

char* readline(void) {
  Buffer *buf = newBuffer();
  int i = 0;
  while (i < 99) {
    char c = getchar(); i++;
    if (c == '\r' || c == '\n') break;
    pushChar(buf,c);
  }
  return bufferToStr(buf);
}

bool isfile(char* fname) {
  FILE* fh = fopen(fname, "r");
  if (fh == NULL) return false;
  return true;
}

char* readfile(char* fname) {
  Buffer *buf = newBuffer();
  FILE* fh = fopen(fname, "r");
  if (fh == NULL) {
    printf("file %s not exists!\n", fname);
    exit(1);
  }
  while (!feof(fh)) {
    char ch = fgetc(fh);
    if (ch == '\r') continue;
    if (ch < 0) continue;
    pushChar(buf, ch);
  }
  fclose(fh);
  return bufferToStr(buf);
}

void writefile(char* fname, char* s) {
  FILE* fh = fopen(fname, "w");
  if (fh == NULL) {
    printf("file %s not exists!\n", fname);
  }
  int i = 0;
  while (true) {
    char ch = s[i]; i++;
    if (ch == '\r') continue;
    if (ch == '\0') break;
    if (ch < 0) break;
    fputc(ch, fh);
  }
  fclose(fh);
}

Strs *split(char* s, char ch) {
  int len = strlen(s);
  Strs *ss = newStrs();
  Buffer *buf = newBuffer();
  bool mode = false;
  int i = 0;
  while (i < len) {
    char c = s[i]; i++;
    if (c == ch) {
      push(ss, bufferToStr(buf));
      buf = newBuffer();
      mode = false;
    } else {
      pushChar(buf, c);
      mode = true;
    }
  }
  if (mode) { push(ss, bufferToStr(buf)); }
  return ss;
}

char* joinStr(Strs *ss, char* bs) {
  Buffer *buf = newBuffer();
  int n = ss->len;
  int i = 0;
  Snode* node = ss->head;
  for (i = 0; i < n; i++) {
    if (i > 0) pushStr(buf,bs);
    pushStr(buf, node->str);
    node = node->next;
  }
  return bufferToStr(buf);
}

char* join(Strs *ss, char ch) {
  Buffer *buf = newBuffer();
  int n = ss->len;
  int i = 0;
  Snode* node = ss->head;
  while (i < n) {
    if (i > 0) { pushChar(buf,ch); }
    pushStr(buf, node->str);
    node = node->next;
    i++;
  }
  return bufferToStr(buf);
}

char* strsToStr(Strs *ss) {
  Buffer *buf = newBuffer();
  int n = ss->len;
  int i = 0;
  Snode* node = ss->head;
  while (i < n) {
    pushStr(buf, node->str);
    node = node->next;
    i++;
  }
  return bufferToStr(buf);
}

char* add(int amount, ...) {
  Buffer *buf = newBuffer();
  va_list vl;
  va_start(vl,amount);
  int i = 0;
  while (i < amount) {
    pushStr(buf, va_arg(vl, char*)); i++;
  }
  va_end(vl);
  return bufferToStr(buf);
}

bool inStrs(char* s,Strs *ss) {
  if (getIndex(ss, s) > -1) return true;
  return false;
}

Strs *map(Strs *ss, char*(*f)(char*)) {
  int n = ss->len;
  Strs *mss = newStrs();
  Snode* node = ss->head;
  while (node != NULL) {
    push(mss, f(node->str));
    node = node->next;
  }
  return mss;
}

struct table {
  Strs *keys;
  Strs *values;
};

typedef struct table Table;

Table *newTable() {
  int size = sizeof(Table);
  Table *t = apply_mem(size);
  t->keys = newStrs();
  t->values = newStrs();
  return t;
}

void setKey(Table *t, char* key, char* value) {
  insert(t->keys,key);
  insert(t->values,value);
}

char* getKey(Table *t, char* key) {
  int idx = getIndex(t->keys,key);
  return getStr(t->values,idx);
}

bool hasKey(Table *t, char* key) {
  return inStrs(key, t->keys);
}

struct tnode {
  Table *table;
  struct tnode *next;
};

typedef struct tnode Tnode;

Tnode *newTnode(void) {
  int size = sizeof(Tnode);
  Tnode *nnode = apply_mem(size);
  nnode->next = NULL;
  return nnode;
}

struct tables {
  Tnode *head;
  int len;
};

typedef struct tables Tables;

Tables *newTables(void) {
  int size = sizeof(Tables);
  Tables *ts = apply_mem(size);
  ts->head = NULL;
  ts->len = 0;
  return ts;
}

void _insert(Tables *ts, Table *tbl) {
  Tnode *nnode = newTnode();
  nnode->table = tbl;
  if (ts->len == 0) {
    ts->head = nnode;
  } else {
    nnode->next = ts->head;
    ts->head = nnode;
  }
  ts->len++;
}

Table *getTable(Tables *ts, int idx) {
  Tnode *tnode = ts->head;
  while (tnode != NULL) {
   if (idx == 0) break;
   tnode = tnode->next;
   idx--;
  }
  return tnode->table;
}

struct tree {
  Strs *keys;
  Tables *values;
};

typedef struct tree Tree;

Tree *newTree() {
  int size = sizeof(Tree);
  Tree *t = apply_mem(size);
  t->keys = newStrs();
  t->values = newTables();
  return t;
}

void setNode(Tree *t,char* key,Table *tbl) {
  insert(t->keys,key);
  _insert(t->values,tbl);
}

Table *getNode(Tree *t, char* key) {
  int idx = getIndex(t->keys,key);
  return getTable(t->values,idx);
}

bool hasNode(Tree *t, char* key) {
  return inStrs(key, t->keys);
}

struct cursor {
  int    at;
  int    len;
  int    depth;
  char*  input;
  char*  str;
  Table  *rtable;
  Strs   *output;
};

typedef struct cursor Cursor;

Cursor *newCursor(void) {
  int size = sizeof(Cursor);
  Cursor *c = apply_mem(size);
  return c;
};

struct lint {
  int    indent;
  int    counter;
  char*  ret;
  char*  ns;
  char*  at;
  Strs   *stack;
  Tree   *stree;
};

typedef struct lint Lint;

Lint *newLint(void) {
  int size = sizeof(Lint);
  Lint *t = apply_mem(size);
  return t;
}

void set2(Strs* ss,char* *a,char* *b) {
  *a = first(ss);
  *b = second(ss);
}

void set3(Strs* ss,char* *a,char* *b,char* * c) {
  Snode *n = ss->head;
  *a = n->str;
  n = n->next;
  *b = n->str;
  n = n->next;
  *c = n->str;
}

bool isvspace(char c) {
  if (c == '\r') return true;
  if (c == '\n') return true;
  return false;
}

bool ishspace(char c) {
  if (c == '\t') return true;
  if (c == ' ') return true;
  return false;
}

bool iswords(char c) {
  if (isalnum(c)) return true;
  if (c == '-') return true;
  if (c == '_') return true;
  return false;
}

char* now(void) {
   time_t s = time(NULL);
   return intToStr(s);
}

Strs* osargs(int argc, char* argv[]) {
  Strs* ss = newStrs();
  int n = 1;
  while (n < argc) {
    push(ss,argv[n]); n++;
  }
  return ss;
}

// int main() { printf("hello world!"); }

