import glfw
from OpenGL.GL import *
from math import *

mode = 0
angle = 0
xRot = 0.0
yRot = 0.0


def main():
    if not glfw.init():
        return
    window = glfw.create_window(640, 640, "Lab2", None, None)
    if not window:
        glfw.terminate()
        return
    glfw.make_context_current(window)
    glfw.set_key_callback(window, key_callback)
    glfw.set_scroll_callback(window, scroll_callback)
    while not glfw.window_should_close(window):
        display(window)
    glfw.destroy_window(window)
    glfw.terminate()


def cube():
    # front
    glBegin(GL_POLYGON)

    glColor3f(1.0, 0.0, 0.0)
    glVertex3f(0.5, -0.5, -0.5)
    glVertex3f(0.5, 0.5, -0.5)
    glVertex3f(-0.5, 0.5, -0.5)
    glVertex3f(-0.5, -0.5, -0.5)
    glEnd()
    # back
    glBegin(GL_POLYGON)
    glColor3f(0.0, 1.0, 0)
    glVertex3f(0.5, -0.5, 0.5)
    glVertex3f(0.5, 0.5, 0.5)
    glVertex3f(-0.5, 0.5, 0.5)
    glVertex3f(-0.5, -0.5, 0.5)
    glEnd()
    # r
    glBegin(GL_POLYGON)
    glColor3f(0.0, 0.0, 1.0)
    glVertex3f(0.5, -0.5, -0.5)
    glVertex3f(0.5, 0.5, -0.5)
    glVertex3f(0.5, 0.5, 0.5)
    glVertex3f(0.5, -0.5, 0.5)
    glEnd()
    # l
    glBegin(GL_POLYGON)
    glColor3f(2, 1.0, 0.0)
    glVertex3f(-0.5, -0.5, 0.5)
    glVertex3f(-0.5, 0.5, 0.5)
    glVertex3f(-0.5, 0.5, -0.5)
    glVertex3f(-0.5, -0.5, -0.5)
    glEnd()
    # t
    glBegin(GL_POLYGON)
    glColor3f(0.0, 2.0, 5.0)
    glVertex3f(0.5, 0.5, 0.5)
    glVertex3f(0.5, 0.5, -0.5)
    glVertex3f(-0.5, 0.5, -0.5)
    glVertex3f(-0.5, 0.5, 0.5)
    glEnd()
    # b
    glBegin(GL_POLYGON)
    glColor3f(1, 0.0, 2)
    glVertex3f(0.5, -0.5, -0.5)
    glVertex3f(0.5, -0.5, 0.5)
    glVertex3f(-0.5, -0.5, 0.5)
    glVertex3f(-0.5, -0.5, -0.5)
    glEnd()


def display(window):
    glEnable(GL_DEPTH_TEST)
    glClear(GL_COLOR_BUFFER_BIT | GL_DEPTH_BUFFER_BIT)
    glViewport(640 - 160, 640 - 160, 160, 160)
    glLoadIdentity()
    glClearColor(1.0, 1.0, 1.0, 1.0)

    rotate_x = [
        1, 0, 0, 0,
        0, cos(xRot), sin(xRot), 0,
        0, -sin(xRot), cos(xRot), 0,
        0, 0, 0, 1
    ]

    rotate_y = [
        cos(yRot), 0, -sin(yRot), 0, 0, 1, 0, 0, sin(yRot), 0, cos(yRot), 0, 0, 0, 0, 1
    ]

    mat1 = [1, 0, 0, 0,
            0, 1, 0, 0,
            0, 0, -1, 0,
            0, 0, 0, 1]
    mat2 = [0, 0, -1, 0,
            0, 1, 0, 0,
            -1, 0, 0, 0,
            0, 0, 0, 1]
    mat3 = [1, 0, 0, 0,
            0, 0, -1, 0,
            0, -1, 0, 0,
            0, 0, 0, 1]

    if mode == 0:
        glPolygonMode(GL_FRONT_AND_BACK, GL_LINE)
    else:
        glPolygonMode(GL_FRONT_AND_BACK, GL_FILL)

    glPushMatrix()
    glMultMatrixf(rotate_x)
    glMultMatrixf(rotate_y)
    cube()
    glPopMatrix()

    glViewport(640 - 3 * 160, 640 - 160, 160, 160)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glMultMatrixf(mat1)

    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()
    glMultMatrixf(rotate_x)
    glMultMatrixf(rotate_y)
    cube()

    glViewport(640 - 3 * 160, 640 - 3 * 160, 160, 160)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glMultMatrixf(mat2)

    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()
    glMultMatrixf(rotate_x)
    glMultMatrixf(rotate_y)
    cube()

    glViewport(640 - 160, 640 - 3 * 160, 160, 160)
    glMatrixMode(GL_PROJECTION)
    glLoadIdentity()
    glMultMatrixf(mat3)

    glMatrixMode(GL_MODELVIEW)
    glLoadIdentity()
    glMultMatrixf(rotate_x)
    glMultMatrixf(rotate_y)
    cube()

    glfw.swap_buffers(window)
    glfw.poll_events()


def key_callback(window, key, scancode, action, mods):
    global mode, xRot, yRot
    if key == glfw.KEY_RIGHT:
        yRot += 0.25

    if key == glfw.KEY_LEFT:
        yRot -= 0.25

    if key == glfw.KEY_UP:
        xRot += 0.25

    if key == glfw.KEY_DOWN:
        xRot -= 0.25
    if action == glfw.PRESS:
        if key == glfw.KEY_SPACE:
            mode = (mode + 1) % 2


def scroll_callback(window, xoffset, yoffset):
    pass


if __name__ == '__main__':
    main()