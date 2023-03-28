#include <stdio.h>
#include <stdlib.h>

long long f(long long n, long long arr[], long long x0, long long k) {
    long long res = 0;
    for (long long i = n; i >= k; --i) {
        res = res * x0 + arr[i];
    }
    return res;
}

int main(int argc, char ** argv)
{
    long long n, x0;
    scanf("%lld%lld", &n, &x0);
    long long arr[n + 1];
    for (long long i = n; i >= 0; --i) {
        scanf("%lld", &arr[i]);              // ������ �������� ������������� ������� � an �� a0
    }
    printf("%lld\n", f(n, arr, x0, 0));      // ��������� �������� �������� � ����� �0
/*  for(int i = 0; i < k; ++i) {           - �������� ������������� ��� k-�� �����������
        for(int j = n; j >= 0; --j) {
            arr[j] *= j - i;
        }
    } */
    for (long long i = n; i >= 0; --i) {
        arr[i] *= i;
    }
    printf("%lld\n", f(n, arr, x0, 1));     // ��������� � ������� �������� ����������� � ����� �0
    return 0;                               // � ����� ������ ��� k-�� �����������: f(n, arr, x0, k)
}
