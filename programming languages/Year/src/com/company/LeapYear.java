package com.company;

public class LeapYear {
    private int year;
    private boolean leap;
    public LeapYear(int inYear) {
        this.year = inYear;
        if(year % 4 == 0 && year % 100 != 0 || year % 400 == 0) {
            this.leap = true;
        } else {
            this.leap = false;
        }
    }
    public boolean getLeap() {
        return leap;
    }
    public String toString() {
        return "(" + year + ", " + leap + ")";
    }
}
