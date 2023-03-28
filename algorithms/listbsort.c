#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Elem elem;

struct Elem {
    struct Elem *next;
    char *word;
};

void swap(elem *a, elem *b) {
    char *tmp = a->word;
    a->word = b->word;
    b->word = tmp;
}

int compare(elem *a, elem *b) {
    return strlen(a->word) > strlen(b->word) ? 1 : 0;
}

struct Elem *bsort(struct Elem *list) {
    if(!list) return list;
    int n = 0;
    struct Elem *q;
    struct Elem *head = list;
    for(struct Elem *p = head; p; p = p->next) {
        ++n;
    }
    int t = n - 1;
    while(t > 0) {
        int bound = t;
        t = 0;
        q = head;
        for (int i = 0; i < bound; ++i, q = q->next) {
            if(compare(q, q->next)) {
                swap(q->next, q);
                t = i;
            }
        }
    }
    return head;
}

void InsertBeforeHead(elem **head, char *s) {
    elem *tmp = (elem*)malloc(sizeof(elem));
    tmp->word = s;
    tmp->next = (*head);
    (*head) = tmp;
}

void PrintList(elem *head) {
    elem *p = head;
    while(p) {
        printf("%s ", p->word);
        p = p->next;
    }
}
void DeleteList(elem **list) {
    elem *prev = NULL;
    if(!(*list)) return;
    while((*list)->next) {
        prev = *list;
        *list = (*list)->next;
        free(prev);
    }
    free(*list);
}

elem *getLast(elem *head) {
    if(head == NULL) return NULL;
    while(head->next) {
        head=head->next;
    }
    return head;
}

void pushBack(elem *head, char *str) {
    elem *p = getLast(head);
    elem *tmp = (elem*)malloc(sizeof(elem));
    tmp->word = str;
    tmp->next = NULL;
    p->next = tmp;
}

int main(int argc, char ** argv) {
    char *str = malloc(1002);
    gets(str);
    char **s;
    s = (char**)malloc(1002 * sizeof(char*));
    int i = 0;
    char *pch = strtok(str, " ");
    while(pch != NULL) {
        s[i] = (char*)malloc(1002);
        strcpy(s[i], pch);
        pch = strtok(NULL, " ");
        ++i;
    }
    free(pch);
    int n = i;
    elem *list = NULL;
    if(n > 0) {
        InsertBeforeHead(&list, s[0]);
    }
    for(i = 1; i < n; ++i) {
        pushBack(list, s[i]);
    }
    list = bsort(list);
    PrintList(list);
    for(i = 0; i < n; ++i) {
        free(s[i]);
    }
    free(s);
    DeleteList(&list);
    free(str);
    return 0;
}