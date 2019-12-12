#!/usr/bin/python
# Demo of OOP in Python

class Point():
    """Location of a point in 2-D space

        attributes: x, y
        """

class Rectangle():
    """Represent a rectangle

    attributes: width, height, corner
    """

def find_center(rect):
        p = Point()
        p.x = rect.corner.x + rect.width/2
        p.y = rect.corner.y + rect.height/2
        return p

def print_point(p):
        print p.x, p.y

def move_rectangle(rect, dx, dy):
        rect.corner.x += dx
        rect.corner.y += dy
        return rect

def compare(obj1, obj2):
        if obj1 is obj2:
                print "%s is identical to %s" % (obj1, obj2)
        else:
                print "%s and %s are different objects" % (obj1, obj2)

# Instantiate an object (instance)
box = Rectangle()

# Set attributes
box.width = 100.0
box.height = 200.0
box.corner = Point() # embedded object
box.corner.x = 0.0
box.corner.y = 0.0

print_point(find_center(box))
move_rectangle(box, 10.0, 10.0)
print_point(find_center(box))

# Create an object alias (another name, the same object)
box2 = box
compare(box2, box)

# Create an object copy (different object) ..
import copy
# .. shallow copy (embedded objects are not copied)
box3 = copy.copy(box)
compare(box3, box)
compare(box3.corner, box.corner)
# .. deep copy
box4 = copy.deepcopy(box)
compare(box4, box)
compare(box4.corner, box.corner)
