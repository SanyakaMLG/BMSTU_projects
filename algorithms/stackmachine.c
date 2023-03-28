#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max(int a, int b) {
    return a > b ? a : b;
}

int min(int a, int b) {
    return a < b ? a : b;
}

struct Stack {
    int *data;
    int top;
    int size;
};

struct Stack *InitStack(int n) {
     struct Stack *stack = malloc(n * sizeof(int) + 100);
     stack->data = (int*)malloc(n * sizeof(int));
     stack->size = n;
     stack->top = 0;
     return stack;
}

void push(struct Stack *stack, int x) {
    stack->data[stack->top] = x;
    stack->top++;
}

int pop(struct Stack *stack) {
    stack->top--;
    return stack->data[stack->top];
}

void add(struct Stack *stack) {
    push(stack, (pop(stack) + pop(stack)));
}

void sub(struct Stack *stack) {
    push(stack, (pop(stack) - pop(stack)));
}

void mul(struct Stack *stack) {
    push(stack, (pop(stack) * pop(stack)));
}

void divide(struct Stack *stack) {
    push(stack, (pop(stack)/(pop(stack))));
}

void maxStack(struct Stack *stack) {
    push(stack, (max(pop(stack), pop(stack))));
}

void minStack(struct Stack *stack) {
    push(stack, (min(pop(stack), pop(stack))));
}

void neg(struct Stack *stack) {
    push(stack, -pop(stack));
}

void dup(struct Stack *stack) {
    int x = pop(stack);
    push(stack, x);
    push(stack, x);
}

void swap(struct Stack *stack) {
    int x = pop(stack);
    int y = pop(stack);
    push(stack, x);
    push(stack, y);
}

int main(int argc, char ** argv) {
    struct Stack *stack;
    int n, x;
    scanf("%d", &n);
    stack = InitStack(n);
    char *expr = (char*)malloc(10);
    for(int i = 0; i < n; ++i) {
        scanf("%s", expr);
        if(!strcmp(expr, "CONST")) {
            scanf("%d", &x);
            push(stack, x);
        }
        if(!strcmp(expr, "ADD")) add(stack);
        if(!strcmp(expr, "SUB")) sub(stack);
        if(!strcmp(expr, "MUL")) mul(stack);
        if(!strcmp(expr, "DIV")) divide(stack);
        if(!strcmp(expr, "MAX")) maxStack(stack);
        if(!strcmp(expr, "MIN")) minStack(stack);
        if(!strcmp(expr, "NEG")) neg(stack);
        if(!strcmp(expr, "DUP")) dup(stack);
        if(!strcmp(expr, "SWAP")) swap(stack);
    }
    printf("%d", stack->data[0]);
    free(stack->data);
    free(stack);
    free(expr);
    return 0;
}