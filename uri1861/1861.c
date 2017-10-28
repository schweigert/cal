#include <stdio.h>
#include <stdlib.h>

#define SZ_DATAMAP 10

#define uint unsigned long long int

#define string char*

#define bool char
#define true 1
#define false 0

struct Data {
  string name;
  bool alive;
  int kd;
};

struct List {
  uint code;
  struct List* other;
  struct Data* data;
};

struct List* datamap[SZ_DATAMAP] = {0};


uint power(uint a, uint b) {
  if(a == 1 || b == 0) return 1;

  b--;
  uint c = a;

  while(b--){
    c *= a;
  }

  return c;
}
uint hash (string str) {
  uint r = 0;
  uint pos;

  for(pos = 0 ; str[pos] != '\0' ; pos++){
    r += (str[pos])*power(255,pos);
  }

  return r;
}
uint pos  (uint code){
  return code % SZ_DATAMAP;
}

int str_len(string a){
  int r;
  for(r = 0; r != '\0'; r++);
  return r;
}

uint data_size = 0;

struct List* create_data(string name) {
  struct Data* d = malloc(sizeof(struct Data));
  d->name = name;
  d->alive = true;
  d->kd = 0;

  struct List* r = malloc(sizeof(struct List));
  r->code = hash(name);
  r->other = NULL;
  r->data = d;

  data_size++;

  return r;
}

struct List* find(string b){
  uint code = hash(b);
  uint position = pos(code);

  struct List* l = datamap[position];

  if(l == NULL){
    datamap[position] = create_data(b);
    l = datamap[position];
    goto ret;
  }

  while(true){
    if(l->code == code){
      goto ret;
    }

    if(l->other == NULL){
      l->other = create_data(b);
      l = l->other;
      goto ret;
    }
    return l->other;
  }

  ret:
  return l;
}


void killer(string b){
  struct List* l = find(b);
  l->data->kd++;
}

void kill(string b){
  struct List* l = find(b);
  l->data->alive = false;
}

string create_name(){
  return (string)(malloc(sizeof(char)*11));
}

void print_alive_list(struct List* l){
  if(l == NULL)
    return;

  if(l->data->alive) printf("%s\t%llu\n", l->data->name, l->code);
  print_alive_list(l->other);
}

void print_alive(){
  uint i;
  for(i = 0; i < SZ_DATAMAP; i++){
    struct List* l = datamap[i];

    if(l == NULL)
      continue;
    printf("Lista: %llu\n",i);
    print_alive_list(l);
  }
}

struct List* get_last(struct List* l){
    if(l->other == NULL)
      return l;

    return get_last(l->other);
}

int comp(const void* elm1, const void* elm2){
  struct List* a = *((struct List**) elm1);
  struct List* b = *((struct List**) elm2);
  return (int)(a->code - b->code);
}

void print_sorted() {
  struct List *l = NULL, *j = NULL, *k = NULL;
  uint i;

  for(i = 0; i < SZ_DATAMAP; i++){
    j = datamap[i];
    if(j == NULL) continue;
    if(k == NULL){
      k = j;
      l = j;
      continue;
    }

    l = get_last(l);
    l->other = j;
  }

  struct List **master = malloc(sizeof(struct List *)*data_size);

  for(i = 0; i < data_size; i++){
    if(k == NULL){
      master[i] = NULL;
    } else {
      master[i] = k;
      k = k->other;
    }
  }

  qsort(master, data_size, sizeof(struct List*), comp);

  for(i = 0; i < data_size; i++){
    if(master[i]->data->alive) printf("%s %d\n",master[i]->data->name,master[i]->data->kd);
  }
}

int main(void){

  while(true){
    string a;
    string b;
    a = create_name();
    b = create_name();
    if(scanf("%s %s",a,b) == EOF){
      break;
    }
    kill(b);
    killer(a);
  }

  printf("HALL OF MURDERERS\n");
  print_sorted();
}
