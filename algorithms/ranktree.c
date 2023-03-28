#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int maxInt(int a, int b) {
    return a > b ? a : b;
}

typedef struct {
    int key;
    char *value;
    int count;
} pair;

typedef struct BinarySearchTree {
    pair elem;
    struct BinarySearchTree *parent;
    struct BinarySearchTree *left;
    struct BinarySearchTree *right;
} tree;

tree *InitTree() {
    tree *t = NULL;
    return t;
}

int MapEmpty(tree *t) {
    return t == NULL;
}

tree *Descend(tree *t, int key) {
    tree *x = t;
    while(x != NULL && x->elem.key != key) {
        if(key < x->elem.key) {
            x = x->left;
        } else {
            x = x->right;
        }
    }
    return x;
}

tree *Minimum(tree *t) {
    tree *x;
    if(t == NULL) {
        x = NULL;
    } else {
        x = t;
        while(x->left) {
            x = x->left;
        }
    }
    return x;
}

tree *Succ(tree *x) {
    tree *y;
    if(x->right) {
        y = Minimum(x->right);
    } else {
        y = x->parent;
        while(y && x == y->right) {
            x = y;
            y = y->parent;
        }
    }
    return y;
}

char *LookUp(tree *t, int key) {
    tree *x = Descend(t, key);
    return x->elem.value;
}

void Insert(tree **t, int key, char *str) {
    tree *y = (tree*)malloc(sizeof(pair) + 3 * sizeof(tree*));
    y->elem.key = key;
    y->elem.value = str;
    y->elem.count = 1;
    y->parent = NULL;
    y->left = NULL;
    y->right = NULL;
    if(*t == NULL) {
        //printf("%p\n", t);
        *t = y;
        //printf("%p", t);
    } else {
        tree *x = *t;
        while(1) {
            x->elem.count++;
            if(key < x->elem.key) {
                if(x->left == NULL) {
                    x->left = y;
                    y->parent = x;
                    return;
                }
                x = x->left;
            } else {
                if(x->right == NULL) {
                    x->right = y;
                    y->parent = x;
                    return;
                }
                x = x->right;
            }
        }
    }
}

void RecalcCounts(tree **t, tree *x) {
    int count = x->elem.count;
    tree *y = x;
    if(x->parent->left && x->parent->right || x->parent->elem.count > count + 1 ||
            x->parent->left->elem.count == x->parent->right->elem.count) return;
    else {
        y = y->parent;
        while(y) {
            y->elem.count--;
            y = y->parent;
        }
    }
}

void ReplaceNode(tree **t, tree *x, tree *y) {
    if(x == *t) {
        *t = y;
        if(y) y->parent = NULL;
    } else {
        tree *p = x;
        if(y) y->parent = p;
        if(p->left == x) {
            p->left = y;
        } else {
            p->right = y;
        }
    }
}

void Delete(tree **t, int key) {
    tree *x = Descend(*t, key);
    if(x->left == NULL && x->right == NULL) {
        RecalcCounts(t, x);
        ReplaceNode(t, x, NULL);
    } else {
        if(x->left == NULL) {
            RecalcCounts(t, x);
            ReplaceNode(t, x, x->right);
        } else {
            if(x->right == NULL) {
                RecalcCounts(t, x);
                ReplaceNode(t, x, x->left);
            } else {
                tree *y = Succ(x);
                RecalcCounts(t, y);
                ReplaceNode(t, y, y->right);
                x->left->parent = y;
                y->left = x->left;
                if(x->right) x->right->parent = y;
                y->right = x->right;
                int count = x->elem.count;
                ReplaceNode(t, x, y);
                y->elem.count = count;
            }
        }
    }
}

char *SearchByRank(tree *t, int rank)
{
    int k = t->left == NULL ? 1 : t->left->elem.count + 1;
    if(k == rank) return t->elem.value;
    else if(rank < k) return SearchByRank(t->left, rank);
    else return SearchByRank(t->right, rank - k);
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    tree *T = InitTree();
    char str[10];
    int key, x;
    char *value = malloc(15);
    for(int i = 0; i < n; ++i) {
        scanf("%s", &str);
        if(!strcmp(str, "INSERT")) {
            scanf("%d %s", &key, value);
            Insert(&T, key, value);
        }
        if(!strcmp(str, "LOOKUP")) {
            scanf("%d", &key);
            printf("%s\n", LookUp(T, key));
        }
        if(!strcmp(str, "DELETE")) {
            scanf("%d", &key);
            Delete(&T, key);
        }
        if(!strcmp(str, "SEARCH")) {
            scanf("%d", &x);
            printf("%s\n", SearchByRank(T, x));
        }
        if(!strcmp(str, "PRINTF")) {
            printf("%s %d\n%s %d", T->elem.value, T->elem.count, T->right->elem.value, T->right->elem.count);
        }
    }
    return 0;
}