#ifndef LAB10_MATRIX_H
#define LAB10_MATRIX_H

int gcd (int a, int b) {
    if(a < 0 || b < 0) {
        a = abs(a);
        b = abs(b);
    }
    while (b) {
        a %= b;
        std::swap(a, b);
    }
    return a;
}

template<int M, int N>
class Matrix {
private:
    class Row;
    Row *column;
    class Row {
    public:
        int *row;
        Row() {
            row = new int[N];
            for(int i = 0; i < N; i++) {
                row[i] = 0;
            }
        }
        int &operator[](int i) {
            return row[i];
        }
    };
public:
    Matrix();
    Row &operator[](int i);
    void fillMatrix();
    void printMatrix();
    class Iterator;
    Iterator begin() {
        return column;
    }
    Iterator end() {
        return column + M;
    }
    class Iterator {
    private:
        Row *cur;
        int row = 0;
    public:
        Iterator(Row *first) : cur(first)
        {}
        Row& operator+ (int n) {
            return *(cur + n);
        }
        Row& operator++ (int) {
            return *cur++;
        }
        Row& operator++() {
            return *++cur;
        };
        bool operator==(const Iterator& it) {
            return cur == it.cur;
        }
        bool operator!=(const Iterator& it) {
            return cur != it.cur;
        }

        Iterator operator=(int x) {
            int oldGCD = cur[0][0];
            int newGCD = x;
            for(int i = 1; i < N; i++) {
                oldGCD = gcd(oldGCD, cur[0][i]);
            }
            for(int i = 0; i < N; i++) {
                cur[0][i] /= oldGCD;
                cur[0][i] *= newGCD;
            }
            return *this;
        }

        int operator*() {
            int curGCD = cur[0][0];
            for(int i = 1; i < N; i++) {
                curGCD = gcd(curGCD, cur[0][i]);
            }
            return curGCD;
        }
    };
};

template<int M, int N>
Matrix<M, N>::Matrix() {
    column = new Row[M];
    for(int i = 0; i < M; i++) {
        column[i] = Row();
    }
}

template<int M, int N>
typename Matrix<M, N>::Row &Matrix<M, N>::operator[](int i) {
    return column[i];
}

template<int M, int N>
void Matrix<M, N>::fillMatrix() {
    for(int i = 0; i < M; i++) {
        for(int j = 0; j < N; j++) {
            int x;
            std::cin >> x;
            column[i][j] = x;
        }
    }
}

template<int M, int N>
void Matrix<M, N>::printMatrix() {
    for(int i = 0; i < M; i++) {
        for(int j = 0; j < N; j++) {
            std::cout << column[i][j] << " ";
        }
        std::cout << std::endl;
    }
}

#endif //LAB10_MATRIX_H
