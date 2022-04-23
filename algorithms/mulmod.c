#include <stdio.h>

int main(int argc, char ** argv) {
    long long a, b, m, res = 0;
    scanf("%lld%lld%lld", &a, &b, &m);
    if(m == 1) {
        printf("%d", 0);
        return 0;
    }
    for (int i = 63; i >= 0; --i) {
        if(res >= m) {
            res = (res % m) * 2;
        } else {
            res = res * 2;
        }
        if(res >= m) {
            res = res % m + (a * ((b >> i) % 2)) % m;
        } else {
            res += a * ((b >> i) % 2);
        }
    }
    printf("%lld", res % m);
    return 0;
}
