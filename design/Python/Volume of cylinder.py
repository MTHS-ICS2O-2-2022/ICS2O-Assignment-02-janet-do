import math


def calculate():
    # input
    radius = float(input("Enter the radius of the cylinder (in cm): "))
    height = float(input("Enter the height of the cylinder (in cm): "))

    # process
    volume = math.pi * radius**2 * height

    # output
    print(f"The volume of the cylinder is: {volume:.2f} cm³")


calculate()