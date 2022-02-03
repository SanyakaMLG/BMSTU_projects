#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max(int a, int b) {
    return a > b ? a : b;
}

int min(int a, int b) {
    return a < b ? a : b;
}

struct CircBuf {
    int *data;
    int count;
    int head;
    int tail;
    int cap;
};

struct CircBuf *InitQueue() {
    struct CircBuf *q = malloc(4 * sizeof(int) + 4 * sizeof(int));
    q->data = (int*)malloc(4 * sizeof(int));
    q->cap = 4;
    q->count = q->head = q->tail = 0;
    return q;
}

int QueueEmpty(struct CircBuf *q) {
    return q->count == 0 ? 1 : 0;
}

void Enqueue(struct CircBuf *q, int x) {
    q->data[q->tail] = x;
    q->tail++;
    if(q->tail == q->cap) {
        q->data = realloc(q->data, sizeof(int) * 2 * q->cap);
        q->cap *= 2;
    }
    q->count++;
}

int Dequeue(struct CircBuf *q) {
    int x = q->data[q->head];
    q->head++;
    if(q->head == q->cap) q->head = 0;
    q->count--;
    return x;
}

int main(int argc, char ** argv) {
    struct CircBuf *q = InitQueue();
    int n, x;
    scanf("%d", &n);
    char *expr = (char*)malloc(10);
    for(int i = 0; i < n; ++i) {
        scanf("%s", expr);
        if(!strcmp(expr, "ENQ")) {
            scanf("%d", &x);
            Enqueue(q, x);
        }
        if(!strcmp(expr, "DEQ")) printf("%d\n", Dequeue(q));
        if(!strcmp(expr, "EMPTY")) {
            QueueEmpty(q) ? printf("true\n") : printf("false\n");
        }
    }
    free(q->data);
    free(q);
    free(expr);
    return 0;
}