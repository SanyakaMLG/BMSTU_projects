#include "declaration.h"

Queue::DoubleStack::DoubleStack(int size) {
    this->arr = new int[size];
    this->cap = size;
    this->top1 = 0;
    this->top2 = size - 1;
}

void Queue::DoubleStack::Push1(int x) {
    this->arr[this->top1] = x;
    this->top1++;
}

void Queue::DoubleStack::Push2(int x) {
    this->arr[this->top2] = x;
    this->top2--;
}

int Queue::DoubleStack::Pop1() {
    this->top1--;
    return this->arr[this->top1];
}

int Queue::DoubleStack::Pop2() {
    this->top2++;
    return this->arr[this->top2];
}

bool Queue::DoubleStack::StackEmpty1() {
    return this->top1 == 0;
}

bool Queue::DoubleStack::StackEmpty2() {
    return this->top2 == this->cap - 1;
}

int &Queue::DoubleStack::operator[](int idx) {
    if(this->top1 > idx) {
        return this->arr[this->top1 - idx - 1];
    } else {
        idx = idx - this->top1;
        return this->arr[this->cap - 1 - idx];
    }
}


Queue::Queue(int size) {
    this->queue = DoubleStack(size);
}

void Queue::Enqueue(int x) {
    this->queue.Push1(x);
    this->numbers++;
}

int Queue::Dequeue() {
    if(this->queue.StackEmpty2()) {
        while(!this->queue.StackEmpty1()) {
            this->queue.Push2(this->queue.Pop1());
        }
    }
    this->numbers--;
    return this->queue.Pop2();
}

int Queue::getSize() {
    return this->numbers;
}

int &Queue::operator[](int idx) {
    return this->queue[idx];
}

