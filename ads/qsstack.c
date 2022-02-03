#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int compare(int *a, int *b) {
    return *a - *b;
}

void swap(int *a, int *b) {
    int tmp = *a;
    *a = *b;
    *b = tmp;
}

struct Task {
    int low, high;
};

typedef struct Task task;

struct Stack {
    task *data;
    int cap;
    int top;
};

struct Stack *InitStack(int n) {
    struct Stack *s = malloc(n * sizeof(task) + 3 * sizeof(int));
    s->data = malloc(n * sizeof(task));
    s->cap = n;
    s->top = 0;
    return s;
}

int StackEmpty(struct Stack *s) {
    return s->top == 0 ? 1 : 0;
}

void Push(struct Stack *s, task x) {
    s->data[s->top] = x;
    s->top++;
}

task Pop(struct Stack *s) {
    s->top--;
    return s->data[s->top];
}

int partition(int *arr, int n, int low, int high) {
    int i = low;
    for(int j = i; j < high; ++j) {
        if(arr[j]<arr[high]) {
            swap(&arr[i], &arr[j]);
            ++i;
        }
    }
    swap(&arr[i], &arr[high]);
    return i;
}

void qsstack(struct Stack *s, int *arr, int nel) {
    int low, high, q;
    task t;
    t.low = 0;
    t.high = nel - 1;
    Push(s, t);
    while(s->top != 0) {
        t = Pop(s);
        low = t.low;
        high = t.high;
        while(low < high) {
            q = partition(arr, nel, low, high);
            t.low = low;
            t.high = q - 1;
            Push(s, t);
            low = q + 1;
        }
    }
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    int arr[n];
    for(int i = 0; i < n; ++i) {
        scanf("%d", &arr[i]);
    }
    struct Stack *s = InitStack(2 * n);
    qsstack(s, arr, n);
    for(int i = 0; i < n; ++i) {
        printf("%d ", arr[i]);
    }
    free(s->data);
    free(s);
    return 0;
}