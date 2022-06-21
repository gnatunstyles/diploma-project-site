import sys
import laspy as lp

args = sys.argv
# todo: give with sys args file saving dir to save point amount to it
print(args)

point_cloud = lp.read(args[1])


print(len(point_cloud))

point_file_name = args[2] + args[3]

print(point_file_name)

with open(point_file_name, 'w') as f:
    f.write(str(len(point_cloud)))
