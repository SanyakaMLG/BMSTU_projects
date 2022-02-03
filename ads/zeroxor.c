#include <stdio.h>
#include <stdlib.h>

typedef struct Stack_s {
    long key;
    int value;
    struct Stack_s *next;
} Stack_t;

typedef Stack_t** HashTable_t;

Stack_t *Push(Stack_t *stack, long key) {
    Stack_t *new = (Stack_t*)malloc(sizeof(Stack_t));
    new->next = stack;
    new->key = key;
    new->value = 0;
    return new;
}

Stack_t *Pop(Stack_t *stack) {
    Stack_t *new_head = stack->next;
    free(stack);
    return new_head;
}

Stack_t *Find(Stack_t *stack, long key) {
    while(stack && stack->key != key) {
        stack = stack->next;
    }
    return stack;
}

int Count(Stack_t *stack) {
    int count = 0;
    while (stack != NULL) {
        count += stack->value;
        stack = stack->next;
    }
    return count;
}

void free_stack(Stack_t *stack) {
    Stack_t *prev;
    while(stack) {
        prev = stack;
        stack = stack->next;
        free(prev);
    }
}

HashTable_t InitTable(int n) {
    HashTable_t table;
    table = (Stack_t**)calloc(n, sizeof(Stack_t *));
    return table;
}

void free_table(HashTable_t table, int n) {
    for(int i = 0; i < n; i++) {
        free_stack(table[i]);
    }
    free(table);
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    HashTable_t table = InitTable(n);
    table[0] = Push(table[0], 0);
    long value, tmp = 0;
    Stack_t *tmp_node;
    int result = 0;
    for (int i = 0; i < n; i++) {
        scanf("%ld", &value);
        tmp ^= value;
        tmp_node = Find(table[tmp % n], tmp);
        if (!tmp_node) {
            table[tmp % n] = Push(table[tmp % n], tmp);
        } else {
            result += tmp_node->value;
            tmp_node->value++;
        }
    }
    for (int i = 0; i < n; i++) {
        result += Count(table[i]);
    }
    printf("%d\n", result);
    free_table(table, n);
    return 0;
}