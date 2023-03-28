package com.company;

public class Particle {
    private Vector coord;
    private double weight;
    private Vector velocity;
    public Particle (Vector inCoord, int inWeight, Vector inVelocity) {
        this.coord = inCoord;
        this.weight = inWeight;
        this.velocity = inVelocity;
    }
    public double getV () {
        return Math.sqrt(coord.getX() * coord.getX() + coord.getY() * coord.getY() +
                coord.getZ() * coord.getZ());
    }
    public double getWeight () {
        return this.weight;
    }
    public double getX() {
        return this.coord.getX();
    }
    public double getY() {
        return this.coord.getY();
    }
    public double getZ() {
        return this.coord.getZ();
    }
}
