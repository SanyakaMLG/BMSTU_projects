#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max(int a, int b) {
    return a > b ? a : b;
}

typedef struct ElemWithMax {
    int val, max;
} elem;

struct DoubleStack {
    elem *data;
    int cap;
    int top1;
    int top2;
};

struct DoubleStack *InitDoubleStack(int n) {
    struct DoubleStack *stack = malloc(n * sizeof(elem) + 3 * sizeof(int));
    stack->data = (elem*)malloc(n * sizeof(elem));
    stack->cap = n;
    stack->top1 = 0;
    stack->top2 = n - 1;
}

int StackEmpty1(struct DoubleStack *stack) {
    return stack->top1 == 0 ? 1 : 0;
}

int StackEmpty2(struct DoubleStack *stack) {
    return stack->top2 == (stack->cap - 1) ? 1 : 0;
}

void push1(struct DoubleStack *stack, int x) {
    stack->data[stack->top1].val = x;
    if(stack->top1 != 0) {
        stack->data[stack->top1].max = max(stack->data[stack->top1 - 1].max, x);
    } else {
        stack->data[stack->top1].max = x;
    }
    stack->top1++;
}

void push2(struct DoubleStack *stack, int x) {
    stack->data[stack->top2].val = x;
    if(stack->top2 != stack->cap - 1) {
        stack->data[stack->top2].max = max(stack->data[stack->top2 + 1].max, x);
    } else {
        stack->data[stack->top2].max = x;
    }
    stack->top2--;
}

int pop1(struct DoubleStack *stack) {
    stack->top1--;
    return stack->data[stack->top1].val;
}

int pop2(struct DoubleStack *stack) {
    stack->top2++;
    return stack->data[stack->top2].val;
}

struct DoubleStack *InitQueue(int n) {
    struct DoubleStack *q = InitDoubleStack(n);
    return q;
}

int QueueEmpty(struct DoubleStack *q) {
    return StackEmpty1(q) && StackEmpty2(q) ? 1 : 0;
}

void Enqueue(struct DoubleStack *q, int x) {
    push1(q, x);
}

int Dequeue(struct DoubleStack *q) {
    if(StackEmpty2(q)) {
        while(!StackEmpty1(q)) {
            push2(q, pop1(q));
        }
    }
    int x = pop2(q);
    return x;
}

int maximum(struct DoubleStack *q) {
    if(!q->top1) {
        return q->data[q->top2 + 1].max;
    } else {
        if(q->top2 == q->cap - 1) {
            return q->data[q->top1 - 1].max;
        } else {
            int x = q->data[q->top1 - 1].max;
            int y = q->data[q->top2 + 1].max;
            return max(x, y);
        }
    }
}

int main(int argc, char ** argv) {
    struct DoubleStack *q;
    int n, x;
    scanf("%d", &n);
    q = InitQueue(n);
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
        if(!strcmp(expr, "MAX")) printf("%d\n", maximum(q));
    }
    free(q->data);
    free(q);
    free(expr);
    return 0;
}