import java.awt.*;

public class Triangle {
    private int x, y, z;

    public void setX(int x) {
        this.x = x;
    }
    public void setY(int y) {
        this.y = y;
    }
    public void setZ(int z) {
        this.z = z;
    }

    public void drawing(Graphics g) {
        double radius; double p;
        p = (double)(x + y + z) / 2;
        radius = (double)(x * y * z) / (4 * Math.sqrt(p) * Math.sqrt(p - x) * Math.sqrt(p - y) * Math.sqrt(p - z));
        double alfa, beta, gamma;
        alfa = 0;
        if(x > 2 * radius) {
            beta = Math.toRadians(180);
            gamma = Math.toRadians(180 + 2 * Math.toDegrees(Math.asin((double)y / (2 * radius))));
        } else {
            if(y > 2 * radius) {
                beta = Math.toRadians(2 * Math.toDegrees(Math.asin((double)x / (2 * radius))));
                gamma = Math.toRadians(180 + 2 * Math.toDegrees(Math.asin((double)x / (2 * radius))));
            } else {
                beta = Math.toRadians(2 * Math.toDegrees(Math.asin((double)x / (2 * radius))));
                gamma = Math.toRadians(2 * Math.toDegrees(Math.asin((double)y / (2 * radius))) + 2 * Math.toDegrees(Math.asin((double)x / (2 * radius))));
            }
        }
        int[] arrX = new int[3];
        int[] arrY = new int[3];
        arrX[0] = (int)(400 + 10 * radius * Math.cos(alfa));
        arrX[1] = (int)(400 + 10 * radius * Math.cos(beta));
        arrX[2] = (int)(400 + 10 * radius * Math.cos(gamma));
        arrY[0] = (int)(350 + 10 * radius * Math.sin(alfa));
        arrY[1] = (int)(350 + 10 * radius * Math.sin(beta));
        arrY[2] = (int)(350 + 10 * radius * Math.sin(gamma));
        int red, green, blue;
        red = (int)( Math.toDegrees(Math.asin((double)x / (2 * radius))) / 0.703);
        green = (int)( Math.toDegrees(Math.asin((double)y / (2 * radius))) / 0.703);
        blue = (int)( Math.toDegrees(Math.asin((double)z / (2 * radius))) / 0.703);
        if(x > 2 * radius) {
            red = 128;
        }
        if(y > 2 * radius) {
            green = 128;
        }
        if(z > 2 * radius) {
            blue = 128;
        }
        g.setColor(new Color(red, green, blue));
        g.fillPolygon(arrX, arrY, 3);
    }
}
