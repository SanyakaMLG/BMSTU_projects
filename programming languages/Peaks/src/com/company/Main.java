package com.company;
import java.util.Arrays;

public class Main {

    public static void main(String[] args) {
	    Sequence[] Sequences = new Sequence[10];
        for(int i = 0; i < 10; i++) {
            Sequences[i] = new Sequence(10);
        }
        Arrays.sort(Sequences);
        for(int i = 0; i < 10; i++) {
            System.out.println(Sequences[i]);
        }
    }
}
