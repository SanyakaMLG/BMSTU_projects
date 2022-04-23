import javax.swing.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;

public class Frame extends JFrame {
    private JLabel label1, label2, label3;
    private JButton button;
    private JTextField field1, field2, field3;
    private GraphicPanel GraphicPanel;
    public Frame() {
        super("Triangle");
        setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        setLayout(null);
        setBounds(0, 0, 800, 800);
        button = new JButton("Draw");
        button.setBounds(710, 5, 70, 40);
        add(button);
        field1 = new JTextField("10");
        field2 = new JTextField("10");
        field3 = new JTextField();

        field1.setBounds(80, 5, 100, 20);
        field2.setBounds(80, 30, 100, 20);
        field3.setBounds(80, 55, 100, 20);
        GraphicPanel = new GraphicPanel();
        button.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                Triangle t = new Triangle();
                GraphicPanel.rePainting(Integer.parseInt(field1.getText()), Integer.parseInt(field2.getText()),
                                            Integer.parseInt(field3.getText()));
            }
        });
        GraphicPanel.setBounds(0, 100, 800, 700);
        add(GraphicPanel);
        label1 = new JLabel("First");
        label1.setBounds(5, 5, 70, 20);
        label2 = new JLabel("Second");
        label2.setBounds(5, 30, 70, 20);
        label3 = new JLabel("Third");
        label3.setBounds(5, 55, 70, 20);
        add(label1);
        add(label2);
        add(label3);

        add(field1);
        add(field2);
        add(field3);
        setVisible(true);
    }
}
