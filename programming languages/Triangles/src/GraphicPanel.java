import javax.swing.*;
import java.awt.*;

public class GraphicPanel extends JPanel {
    private Triangle triangle;

    public GraphicPanel() {
        triangle = new Triangle();
    }

    public void rePainting(int x, int y, int z) {
        triangle.setX(x);
        triangle.setY(y);
        triangle.setZ(z);
        repaint();
    }

    @Override
    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        triangle.drawing(g);
    }
}
