#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

int gcd(int a, int b) {
    a = abs(a);
    b = abs(b);
    while(a && b) {
        if(a >= b) {
            a %= b;
        } else {
            b %= a;
        }
    }
    return a | b;
}

int *ComputeLogarithms(int n) {
    int m = log2(n);
    int size = (1 << (m + 1));
    int *lg = (int*)malloc(size * sizeof(int));
    int j = 0;
    for(int i = 1; i <= (m + 1); ++i) {
        int end = (1 << i);
        while(j < end) {
            lg[j] = i - 1;
            ++j;
        }
    }
    return lg;
}

int SparseTable_Query(int **ST, int L, int R, int *lg) {
    int j = lg[R - L + 1];
    int v = gcd(ST[L][j], ST[R - (1 << j) + 1][j]);
    return v;
}

int **SparseTable_Build(int *arr, int n, int *lg) {
    int m = lg[n] + 1;
    int **ST = (int**)malloc(n * sizeof(int*));
    for(int i = 0; i < n; ++i) {
        ST[i] = (int*)malloc(m * sizeof(int));
    }
    for(int i = 0; i < n; ++i) {
        ST[i][0] = arr[i];
    }
    for(int i = 1; i < m; ++i) {
        int end = n - (1 << i);
        for(int j = 0; j <= end; ++j) {
            ST[j][i] = gcd(ST[j][i - 1], ST[j + (1 << (i - 1))][i - 1]);
        }
    }
    return ST;
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    int *lg = ComputeLogarithms(n);
    int *arr = (int*)malloc(n * sizeof(int));
    for(int i = 0; i < n; ++i) {
        int a;
        scanf("%d", &a);
        arr[i] = a;
    }
    int **ST = SparseTable_Build(arr, n, lg);
    int m;
    scanf("%d", &m);
    int L, R;
    for(int i = 0; i < m; ++i) {
        scanf("%d %d", &L, &R);
        printf("%d ", SparseTable_Query(ST, L, R, lg));
    }
    for(int i = 0; i < n; ++i) {
        free(ST[i]);
    }
    free(ST);
    free(lg);
    free(arr);
    return 0;
}