package com.company;

public class PrimeDiv implements Comparable<PrimeDiv> {
    private int number;
    private int countPrimeDevisors;
    public PrimeDiv (int inX) {
        this.number = inX;
        this.countPrimeDevisors = this.count();
    }
    private int count() {
        int n = this.number;
        int res = 0;
        for(int i = 2; i <= n; i++) {
            if(n % i == 0) {
                res++;
                n = n/i;
                i = 1;
            }
        }
        return res;
    }
    public String toString() {
        return Integer.toString(number);
    }
    public int compareTo(PrimeDiv obj) {
        return this.countPrimeDevisors - obj.countPrimeDevisors;
    }
}
