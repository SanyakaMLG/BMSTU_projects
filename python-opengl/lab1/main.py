from OpenGL.GL import *
import glfw
import math

delta = 0.0
angle = 0.0
posx = 0.0
posy = 0.0


def key_callback(window, key, scancode, action, mods):
    global delta
    global angle
    if action == glfw.PRESS:
        if key == glfw.KEY_RIGHT:
            delta += 0.1
        if key == glfw.KEY_LEFT:
            delta -= 0.1


def circle_draw(radius, center_x, center_y, color_r, color_g, color_b):
    x = 0
    y = 0
    line_amount = 200
    twice_pi = 2.0 * 3.1415
    glBegin(GL_POLYGON)
    for i in range(line_amount):
        glVertex2f(
            center_x + x + (radius * math.cos(i * twice_pi / line_amount)),
            center_y + y + (radius * math.sin(i * twice_pi / line_amount))
        )
        glColor3f(color_r, color_g, color_b)
    glEnd()


def display(window):
    global angle
    glClear(GL_COLOR_BUFFER_BIT)
    glLoadIdentity()
    glClearColor(1.0, 1.0, 1.0, 1.0)
    glPushMatrix()
    glRotatef(angle, 0, 0, 1)

    glBegin(GL_POLYGON)
    for x in range(-50, 51):
        glVertex2f(x/200, -((x/50) ** 2) + 0.5)
        glColor3f(0.5, 0.5, 0.5)
    glEnd()

    glBegin(GL_LINE_LOOP)
    glVertex2f(0, 0.5)
    glColor3f(0, 0, 0)
    glVertex2f(0, -0.5)
    glColor3f(0, 0, 0)
    glEnd()

    circle_draw(0.25, 0.25, -0.5, 0.5, 0.5, 0.5)
    circle_draw(0.25, -0.25, -0.5, 0.5, 0.5, 0.5)
    circle_draw(0.2, 0.25, -0.5, 255, 255, 255)
    circle_draw(0.2, -0.25, -0.5, 255, 255, 255)

    glPopMatrix()
    angle += delta
    glfw.swap_buffers(window)
    glfw.poll_events()


def main():
    if not glfw.init():
        return
    window = glfw.create_window(640, 640, "Lab1", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    while not glfw.window_should_close(window):
        display(window)
    glfw.destroy_window(window)
    glfw.terminate()


if __name__ == '__main__':
    main()
