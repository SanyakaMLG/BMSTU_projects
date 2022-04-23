#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Elem {
    int nel, idx;
    int *value;
} elem;

typedef struct PriorityQueue {
    elem *heap;
    int cap;
    int count;
} queue;

void swap(elem *a, elem *b) {
    elem c = *a;
    *a = *b;
    *b = c;
}

queue *InitQueue(int n) {
    queue *q = (queue*)malloc(n * sizeof(elem) + 10 * sizeof(int));
    q->heap = (elem*)malloc(n * sizeof(elem));
    q->cap = n;
    q->count = 0;
    return q;
}

void Insert(queue *q, elem x) {
    int i = q->count;
    q->count = i + 1;
    q->heap[i] = x;
    while(i > 0 && q->heap[(i - 1)/2].value[q->heap[(i - 1)/2].idx] > q->heap[i].value[q->heap[i].idx]) {
        swap(&q->heap[(i - 1)/2], &q->heap[i]);
        i = (i - 1)/2;
    }
}

void heapify(queue *q, int nel, int i) {
    int L, R, j;
    while(1) {
        L = 2 * i + 1;
        R = 2 * i + 2;
        j = i;
        if(L < nel && q->heap[i].value[q->heap[i].idx] > q->heap[L].value[q->heap[L].idx]) i = L;
        if(R < nel && q->heap[i].value[q->heap[i].idx] > q->heap[R].value[q->heap[R].idx]) i = R;
        if(i == j) break;
        swap(&q->heap[i], &q->heap[j]);
    }
}

elem ExtractMax(queue *q) {
    elem p = q->heap[0];
    q->count--;
    if(q->count > 0) {
        q->heap[0] = q->heap[q->count];
        heapify(q, q->count, 0);
    }
    return p;
}

int QueueEmpty(queue *q) {
    return q->count == 0 ? 1 : 0;
}

int main(int argc, char ** argv) {
    int k, n = 0;
    scanf("%d", &k);
    int **arr = (int**)malloc(sizeof(int*) * k);
    int length[k];
    for(int i = 0; i < k; ++i) {
        scanf("%d", &length[i]);
        n += length[i];
    }
    for(int i = 0; i < k; ++i) {
        arr[i] = (int*)malloc(sizeof(int) * length[i]);
        for(int j = 0; j < length[i]; ++j) {
            scanf("%d", &arr[i][j]);
        }
    }
    queue *q = InitQueue(k);
    for(int i = 0; i < k; ++i) {
        elem x;
        x.nel = length[i];
        x.value = arr[i];
        x.idx = 0;
        Insert(q, x);
    }
    while(!QueueEmpty(q)) {
        elem tmp = ExtractMax(q);
        printf("%d ", tmp.value[tmp.idx]);
        int idx = ++tmp.idx;
        if(idx < tmp.nel) {
            tmp.value[tmp.idx] = tmp.value[tmp.idx];
            Insert(q, tmp);
        }
    }
    for(int i = 0; i < k; ++i) {
        free(arr[i]);
    }
    free(arr);
    free(q->heap);
    free(q);
    return 0;
}