package com.company;
import java.util.Arrays;

public class Main {

    public static void main(String[] args) {
	    PrimeDiv[] numbers = new PrimeDiv[] {
                new PrimeDiv(5),
                new PrimeDiv(7),
                new PrimeDiv(20),
                new PrimeDiv(194),
                new PrimeDiv(2),
                new PrimeDiv(100),
                new PrimeDiv(50),
        };
        Arrays.sort(numbers);
        for (PrimeDiv a : numbers) System.out.println(a);
    }
}
