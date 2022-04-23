#include <stdio.h>
#include <stdlib.h>
#include <string.h>

int max(int a, int b) {
    return a > b ? a : b;
}

int query(int *T, int L, int R, int i, int a, int b) {
    if(L == a && R == b) {
        return T[i];
    } else {
        int m = (a + b)/2;
        if(R <= m) {
            return query(T, L, R, 2*i + 1, a, m);
        } else {
            if(L > m) {
                return query(T, L, R, 2*i + 2, m + 1, b);
            } else {
                return max(query(T, L, m, 2*i + 1, a, m), query(T, m + 1, R, 2*i + 2, m + 1, b));
            }
        }
    }
}

int SegmentTree_query(int *T, int n, int L, int R) {
    return query(T, L, R, 0, 0, n - 1);
}

void build(int *arr, int i, int a, int b, int *T) {
    if(a == b) {
        T[i] = arr[a];
    } else {
        int m = (a + b)/2;
        build(arr, 2*i + 1, a, m, T);
        build(arr, 2*i + 2, m + 1, b, T);
        T[i] = max(T[2*i + 1], T[2*i + 2]);
    }
}

void SegmentTree_Build(int *arr, int n, int *T) {
    build(arr, 0, 0, n - 1, T);
}

void update(int j, int x, int i, int a, int b, int *T) {
    if(a == b) {
        T[i] = x;
    } else {
        int m = (a + b) / 2;
        if(j <= m) {
            update(j, x, 2*i + 1, a, m, T);
        } else {
            update(j, x, 2*i + 2, m + 1, b, T);
        }
        T[i] = max(T[2*i + 1], T[2*i + 2]);
    }
}

void SegmentTree_Update(int j, int x, int n, int *T) {
    update(j, x, 0, 0, n - 1, T);
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    int arr[n];
    for(int i = 0; i < n; ++i) {
        scanf("%d", &arr[i]);
    }
    int m;
    scanf("%d", &m);
    int *T = (int*)malloc(4 * n * sizeof(int));
    SegmentTree_Build(arr, n, T);
    char *expr = (char*)malloc(4);
    int a, b;
    for(int i = 0; i < m; ++i) {
        scanf("%s %d %d", expr, &a, &b);
        if(!strcmp(expr, "MAX")) {
            printf("%d ", SegmentTree_query(T, n, a, b));
        } else {
            SegmentTree_Update(a, b, n, T);
        }
    }
    free(T);
    free(expr);
    return 0;
}