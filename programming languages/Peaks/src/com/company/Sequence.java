package com.company;
import java.util.Random;

public class Sequence implements Comparable<Sequence> {
    private int[] array;
    private int peaks;
    public Sequence(int size) {
        this.array = new int[size];
        this.peaks = 0;
        Random rnd = new Random();
        for(int i = 0; i < size; i++) {
            this.array[i] = rnd.nextInt(101);
        }
        if(size == 1) {
            this.peaks = 1;
        } else {
            for (int i = 0; i < size; i++) {
                if (i == 0) {
                    if(this.array[i] > this.array[i + 1]) this.peaks++;
                } else {
                    if (i == size - 1) {
                        if(this.array[i] > this.array[i - 1]) this.peaks++;
                    } else {
                        if (this.array[i] > this.array[i - 1] && this.array[i] > this.array[i + 1]) {
                            this.peaks++;
                        }
                    }
                }
            }
        }
    }
    public int compareTo(Sequence obj) {
        return this.peaks - obj.peaks;
    }
    public String toString() {
        String str = "array: (";
        for(int i = 0; i < this.array.length - 1; i++) {
            str += Integer.toString(this.array[i]);
            str += ", ";
        }
        str += Integer.toString(this.array[this.array.length - 1]);
        str += "), peaks: ";
        str += Integer.toString(this.peaks);
        return str;
    }
}
