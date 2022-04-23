package com.company;
import java.util.ArrayList;

public class Universe {
    private ArrayList<Particle> particles;
    public void addParticle(Particle inX) {
        this.particles.add(inX);
    }
    public Universe() {
        this.particles = new ArrayList<>();
    }
    public double getSumWeight() {
        double res = 0;
        for(int i = 0; i < this.particles.size(); i++) {
            Particle part = this.particles.get(i);
            double w = part.getWeight();
            res += w;
        }
        return res;
    }
    public double getAverageWeight() {
        return this.getSumWeight() / this.particles.size();
    }
    public double getAverageKinetic() {
        double res = 0;
        for(int i = 0; i < this.particles.size(); i++) {
            Particle part = this.particles.get(i);
            double velo = part.getV();
            double w = part.getWeight();
            double kinetic = w * velo * velo / 2;
            res += kinetic;
        }
        return res / this.particles.size();
    }
}
