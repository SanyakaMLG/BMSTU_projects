#include <stdio.h>
#include <stdlib.h>

typedef struct elem elem;

struct elem{
    int key;
    int value;
    elem* next;
};

elem *Init(int k, int v) {
    elem *x = malloc(sizeof(elem));
    x->key = k;
    x->value = v;
    x->next = NULL;
    return x;
}

elem *Push(elem *x, elem *y) {
    y->next = x;
    return y;
}
elem *Search(elem *x, int k) {
    elem *y = x;
    while(!(x == NULL || x->key == k)) {
        x = x->next;
    }
    return x;
}


int Lookup(elem **t, int m, int k){
    elem *p = Search(t[k % m ], k);
    if(!p) return 0;
    return p->value;
}

void Freeelem(elem *t){
    elem *x = t;
    while(t){
        t = t->next;
        free(x);
        x = t;
    }
}


int main(){
    int n, m, key, value;
    scanf("%d", &n);
    scanf("%d", &m);
    elem **t = calloc(m, sizeof(elem*));
    for(int i = 0 ; i < n ; i++){
        char str[10];
        scanf("%s", str);
        if((str[1] != 'S')){
            scanf("%d", &key);
            printf("%d\n", Lookup(t,m,key));
        }
        else{
            scanf("%d %d", &key, &value);
            t[key % m] = Push(t[key % m],Init(key, value));
        }
    }
    for(int i = 0; i < m; ++i) {
        Freeelem(t[i]);
    }
    free(t);
    return 0;
}