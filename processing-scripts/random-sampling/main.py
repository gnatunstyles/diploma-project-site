# %%
import sys
import numpy as np
import laspy as lp
import matplotlib.pyplot as plt
from mpl_toolkits import mplot3d


args = sys.argv
print(args)

point_cloud = lp.read(
    sys.argv[1])


print(sys.argv[1])
print(point_cloud.point_format)

points = np.vstack((point_cloud.x, point_cloud.y,
                    point_cloud.z)).transpose()

colors = np.vstack((point_cloud.red, point_cloud.green,
                   point_cloud.blue)).transpose()

factor = int(args[4])

decimated_points = points[::factor]
decimated_colors = colors[::factor]

ax = plt.axes(projection='3d')

ax.scatter(decimated_points[:, 0], decimated_points[:, 1],
           decimated_points[:, 2], c=decimated_colors/65535, s=0.01)

print(plt.show())


header = lp.LasHeader(point_format=3, version="1.2")
header.add_extra_dim(lp.ExtraBytesParams(name="random", type=np.int32))
header.scales = np.array([0.1, 0.1, 0.1])


new_las_file = lp.create(
    point_format=header.point_format,
    file_version=header.version)


new_las_file.xyz = decimated_points

new_las_file.red = point_cloud.red[::factor]
new_las_file.green = point_cloud.green[::factor]
new_las_file.blue = point_cloud.blue[::factor]


print(new_las_file.points)


file_save_string = args[2] + args[3]

new_las_file.write(file_save_string+'.las')


# %%
