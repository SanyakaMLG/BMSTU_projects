#include <stdio.h>

int abs(int a) {
    return a < 0 ? -a : a;
}

void insertionSort(int *arr, int l, int r) {
    for (int step = l + 1; step <= r; step++) {
        int key = arr[step];
        int j = step - 1;
        while (abs(key) < abs(arr[j]) && j >= 0) {
            arr[j + 1] = arr[j];
            --j;
        }
        arr[j + 1] = key;
    }
}

void merge(int *arr, int p, int q, int r) {
    int n1 = q - p + 1;
    int n2 = r - q;
    int L[n1], M[n2];
    for (int i = 0; i < n1; i++)
        L[i] = arr[p + i];
    for (int j = 0; j < n2; j++)
        M[j] = arr[q + 1 + j];
    int i, j, k;
    i = 0;
    j = 0;
    k = p;
    while (i < n1 && j < n2) {
        if (abs(L[i]) <= abs(M[j])) {
            arr[k] = L[i];
            i++;
        } else {
            arr[k] = M[j];
            j++;
        }
        k++;
    }
    while (i < n1) {
        arr[k] = L[i];
        i++;
        k++;
    }
    while (j < n2) {
        arr[k] = M[j];
        j++;
        k++;
    }
}

void mergeSort(int *arr, int l, int r) {
    if (l - r + 1 > 5) {
        int m = l + (r - l) / 2;
        mergeSort(arr, l, m);
        mergeSort(arr, m + 1, r);
        merge(arr, l, m, r);
    } else {
        insertionSort(arr, l, r);
    }
}

int main(int argc, char ** argv) {
    int n;
    scanf("%d", &n);
    int arr[n];
    for(int i = 0; i < n; ++i) {
        scanf("%d", &arr[i]);
    }
    mergeSort(arr, 0, n - 1);
    for(int i = 0; i < n; ++i) {
        printf("%d ", arr[i]);
    }
    return 0;
}