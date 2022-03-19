import javax.swing.*;
import java.awt.*;

public class GraphicPanel extends JPanel {
    private Triangle triangle;
    public GraphicPanel() {
        triangle = new Triangle();
    }

    public void rePainting(int x, int y) {
        triangle.setX(x);
        triangle.setY(y);
        repaint();
    }

    @Override
    protected void paintComponent(Graphics g) {
        super.paintComponent(g);
        triangle.drawing(g);
    }
}
