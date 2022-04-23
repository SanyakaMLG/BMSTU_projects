#ifndef LAB8_EQUALITY_H
#define LAB8_EQUALITY_H
#include <iostream>
#include <vector>

using namespace std;

template<typename T>
class Equality {
private:
    vector<T> arr;
    T result;
public:
    Equality(int size, T res) {
        result = res;
    }
    bool isEqual(vector<T>);
    T operator[](int);
    int getSize();
    void setVector(int);
    void setVector(vector<T>);
};

template<typename T>
bool Equality<T>::isEqual(vector<T> x) {
    if(x.size() != arr.size()) {
        cout << "Error: Different dimensions.\n";
        return false;
    } else {
        T res = 0;
        for(int i = 0; i < arr.size(); i++) {
            res += arr[i] * x[i];
        }
        return res == result;
    }
}

template<typename T>
T Equality<T>::operator[](int i) {
    return arr[i];
}

template<typename T>
int Equality<T>::getSize() {
    return arr.size();
}

template<typename T>
void Equality<T>::setVector(int size) {
    for(int i = 0; i < size; i++) {
        T x;
        cin >> x;
        arr.push_back(x);
    }
}

template<typename T>
void Equality<T>::setVector(vector<T> x) {
    arr = x;
}

template<>
bool Equality<bool>::isEqual(vector<bool> x) {
    if(x.size() != arr.size()) {
        cout << "Error: Difference dimensions.\n";
        return false;
    } else {
        bool res = false;
        for(int i = 0; i < arr.size(); i++) {
            if(arr[i] && x[i]) {
                res = true;
                break;
            }
        }
        return res == result;
    }
}

#endif //LAB8_EQUALITY_H