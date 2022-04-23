#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct Task {
    int t1, t2;
} time;

typedef struct PriorityQueue {
    time *heap;
    int count;
    int cap;
} queue;

struct PriorityQueue *InitQueue(int n) {
    queue *q = (queue*)malloc(n * sizeof(time) + 10 * sizeof(int));
    q->heap = (time*)malloc(n * sizeof(time));
    q->cap = n;
    q->count = 0;
    return q;
}

void swap(time *task1, time *task2) {
    time tmp = *task1;
    *task1 = *task2;
    *task2 = tmp;
}

void Insert(queue *q, time task) {
    int i = q->count;
    q->count = i + 1;
    q->heap[i] = task;
    while(i > 0 && (q->heap[(i - 1)/2].t1 + q->heap[(i - 1)/2].t2) > (q->heap[i].t1 + q->heap[i].t2)) {
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
        if(L < nel && (q->heap[i].t1 + q->heap[i].t2) > (q->heap[L].t1 + q->heap[L].t2)) i = L;
        if(R < nel && (q->heap[i].t1 + q->heap[i].t2) > (q->heap[R].t1 + q->heap[R].t2)) i = R;
        if(i == j) break;
        swap(&q->heap[i], &q->heap[j]);
    }
}

time ExtractMin(queue *q) {
    time p = q->heap[0];
    q->count--;
    if(q->count > 0) {
        q->heap[0] = q->heap[q->count];
        heapify(q, q->count, 0);
    }
    return p;
}

int main(int argc, char ** argv) {
    int n, m, res = 0, time1, time2;
    queue *q;
    scanf("%d%d", &n, &m);
    q = InitQueue(n);
    if(n < m) {
        for(int i = 0; i < n; ++i) {
            scanf("%d %d", &time1, &time2);
            time x;
            x.t1 = time1;
            x.t2 = time2;
            if(time1 + time2 > res) res = time1 + time2;
            Insert(q, x);
        }
        for(int i = n; i < m; ++i) {
            time task = ExtractMin(q);
            time x;
            scanf("%d %d", &time1, &time2);
            x.t1 = time1;
            x.t2 = time2;
            if(task.t1 + task.t2 <= x.t1) {
                task.t1 = x.t1;
            } else {
                task.t1 += task.t2;
            }
            task.t2 = x.t2;
            if((task.t1 + task.t2) > res) {
                res = task.t1 + task.t2;
            }
            Insert(q, task);
        }
    } else {
        for(int i = 0; i < m; ++i) {
            scanf("%d %d", &time1, &time2);
            if(time1 + time2 > res) res = time1 + time2;
        }
    }
    printf("%d", res);
    free(q->heap);
    free(q);
    return 0;
}