#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

void Kadane(double *arr, int n) {
    double maxprod = arr[0], prod = 0.0;
    int start = 0, i = 0, L = 0, R = 0;
    while(i < n) {
        prod += log2(arr[i]);
        if(prod > maxprod) {
            maxprod = prod;
            L = start;
            R = i;
        }
        i++;
        if(prod < 0) {
            prod = 0;
            start = i;
        }
    }
    printf("%d %d", L, R);
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    double *arr = (double*)malloc((n + 1) * sizeof(double));
    for(int i = 0; i < n; ++i) {
        int a, b;
        scanf("%d/%d", &a, &b);
        double c = 1.0 * a / b;
        arr[i] = c;
    }
    Kadane(arr, n);
    free(arr);
    return 0;
}