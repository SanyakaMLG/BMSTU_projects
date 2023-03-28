#ifndef LAB9_SPARSEMATRIX_H
#define LAB9_SPARSEMATRIX_H
#include <utility>
#include <map>
#include <iostream>

using namespace std;

template<typename T, int M, int N>
class SparseMatrix{
private:
    map<pair<int, int>, T> Matrix;
public:
    SparseMatrix() = default;
    void setElem(int row, int column, T x);
    void plusElem(int row, int column, T x);
    void show();
    T& operator()(int, int);
    SparseMatrix operator+(SparseMatrix &other);
    SparseMatrix operator*(SparseMatrix &other);
    map<pair<int, int>, T> getMatrix();
    bool operator==(SparseMatrix &other);
    bool operator!=(SparseMatrix &other);
};

template<typename T, int M, int N>
map<pair<int, int>, T> SparseMatrix<T, M, N>::getMatrix() {
    return Matrix;
}

template<typename T, int M, int N>
void SparseMatrix<T, M, N>::plusElem(int row, int column, T x) {
    pair<int, int> idx = {row, column};
    if(Matrix.find(idx) == Matrix.end()) {
        Matrix[idx] = x;
    } else {
        Matrix[idx] += x;
    }
}

template<typename T, int M, int N>
void SparseMatrix<T, M, N>::setElem(int row, int column, T x) {
    pair<int, int> idx;
    idx.first = row;
    idx.second = column;
    if(x == 0 && Matrix.find(idx) != Matrix.end()) {
        Matrix.erase(idx);
    } else {
        if(x != 0) {
            Matrix[idx] = x;
        }
    }
}

template<typename T, int M, int N>
void SparseMatrix<T, M, N>::show() {
    for(int i = 0; i < M; i++) {
        for(int j = 0; j < N; j++) {
            pair<int, int> idx;
            idx.first = i;
            idx.second = j;
            auto search = Matrix.find(idx);
            if(search != Matrix.end()) {
                cout << Matrix[idx] << " ";
            } else {
                cout << 0 << " ";
            }
        }
        cout << endl;
    }
}

template<typename T, int M, int N>
T &SparseMatrix<T, M, N>::operator()(int row, int column) {
    pair<int, int> idx;
    idx.first = row;
    idx.second = column;
    if(Matrix.find(idx) != Matrix.end()) {
        return Matrix[idx];
    } else {
        Matrix[idx] = 0;
        return Matrix[idx];
    }
}

template<typename T, int M, int N>
SparseMatrix<T, M, N> SparseMatrix<T, M, N>::operator+(SparseMatrix<T, M, N> &other) {
    SparseMatrix<T, M, N> tmpMatrix;
    for(int i = 0; i < M; i++) {
        for(int j = 0; j < N; j++) {
            pair<int, int> idx;
            auto tmp1 = (*this)(i, j);
            auto tmp2 = other(i, j);
            if(tmp1 + tmp2 != 0) {
                tmpMatrix.setElem(i, j, tmp1 + tmp2);
            }
        }
    }
    return tmpMatrix;
}

template<typename T, int M, int N>
SparseMatrix<T, M, N> operator*(int k, SparseMatrix<T, M, N> &m) {
    SparseMatrix<T, M, N> tmpMatrix;
    map<pair<int, int>, T> tmpMap = m.getMatrix();
    for(auto elem : tmpMap) {
        T x = elem.second;
        x *= k;
        int row = (elem.first).first;
        int column = (elem.first).second;
        tmpMatrix.setElem(row, column, x);
    }
    return tmpMatrix;
}

template<typename T, int M, int N>
SparseMatrix<T, M, N> operator*(SparseMatrix<T, M, N> &m, int k) {
    SparseMatrix<T, M, N> tmpMatrix;
    map<pair<int, int>, T> tmpMap = m.getMatrix();
    for(auto elem : tmpMap) {
        T x = elem.second;
        x *= k;
        int row = (elem.first).first;
        int column = (elem.first).second;
        tmpMatrix.setElem(row, column, x);
    }
    return tmpMatrix;
}

template<typename T, int M, int N>
bool SparseMatrix<T, M, N>::operator==(SparseMatrix<T, M, N> &other) {
    if((*this).getMatrix().size() != other.getMatrix().size()) {
        return false;
    }
    for(auto elem : (*this).getMatrix()) {
        if(other.getMatrix().find(elem.first) == other.getMatrix().end() ||
           other.getMatrix()[elem.first] != elem.second) {
            return false;
        }
    }
    return true;
}

template<typename T, int M, int N>
bool SparseMatrix<T, M, N>::operator!=(SparseMatrix<T, M, N> &other) {
    return !(*this == other);
}

template<typename T, int M, int N>
SparseMatrix<T, M, N> SparseMatrix<T, M, N>::operator*(SparseMatrix<T, M, N> &other) {
    SparseMatrix<T, M, N> tmpMatrix;
    for(auto elem : this->getMatrix()) {
        for(int i = 0; i < N; i++) {
            tmpMatrix.plusElem(elem.first.first, i, elem.second *
                                                                   other(elem.first.second, i));
        }
    }
    return tmpMatrix;
}

#endif //LAB9_SPARSEMATRIX_H
