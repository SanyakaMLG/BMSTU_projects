package com.company;

public class Main {

    public static void main(String[] args) {
	    Universe MilkyWay = new Universe();
        for(int i = 0; i < 100000; ++i) {
            Vector X = new Vector(Math.random()*100, Math.random()*100, Math.random()*100);
            Vector Velo = new Vector(Math.random()*100, Math.random()*100, Math.random()*100);
            int weight = (int)(Math.random() * 100);
            Particle p = new Particle(X, weight, Velo);
            MilkyWay.addParticle(p);
        }
        System.out.println(MilkyWay.getAverageKinetic());
        System.out.println(MilkyWay.getSumWeight());
        System.out.print(MilkyWay.getAverageWeight());
    }
}
