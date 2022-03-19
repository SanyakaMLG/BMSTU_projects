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
        g.setColor(Color.BLACK);
        g.drawRect((int)(Math.random() * 400), (int)(Math.random() * 400), x, y);
    }
}
