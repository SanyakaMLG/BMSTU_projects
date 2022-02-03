#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Elem {
    struct Elem *prev, *next;
    int v;
} elem;

elem *InitList() {
    elem *list = malloc(sizeof(elem));
    list->next = list;
    list->prev = list;
    return list;
}

void InsertAfter(elem *list, elem *new) {
    elem *p = list->next;
    list->next = new;
    new->prev = list;
    new->next = p;
    p->prev = new;
}

void Delete(elem *del) {
    elem *p = del->prev;
    elem *q = del->next;
    p->next = q;
    q->prev = p;
    del->prev = NULL;
    del->next = NULL;
    free(del);
}

void InsertionSort(elem *head, int n) {
    elem *p = head->next->next;
    while(p != head) {
        int el = p->v;
        elem *loc = p->prev;
        while(loc != head && loc->v > el) {
            loc = loc->prev;
        }
        elem *x = malloc(sizeof(elem));
        x->v = p->v;
        InsertAfter(loc, x);
        p = p->next;
        Delete(p->prev);
    }
}

void PrintList(elem *head) {
    elem *p = head->next;
    while(p != head) {
        printf("%d ", p->v);
        p = p->next;
    }
}

void DelList(elem *head) {
    elem *p = head->next;
    while(p != head) {
        p = p->next;
        free(p->prev);
    }
    free(p);
}

int main(int argc, char ** argv) {
    int n, a;
    elem *list;
    scanf("%d", &n);
    list = InitList();
    for(int i = 0; i < n; ++i) {
        scanf("%d", &a);
        elem *x = malloc(sizeof(elem));
        x->v = a;
        InsertAfter(list, x);
    }
    InsertionSort(list, n);
    PrintList(list);
    DelList(list);
    return 0;
}