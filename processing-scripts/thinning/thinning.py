# %%
import sys
import numpy as np
import laspy as lp
import matplotlib.pyplot as plt


args = sys.argv
print(args)

point_cloud = lp.read(
    sys.argv[1])


print(sys.argv[1])

print(point_cloud.point_format)

transponded_pts = np.vstack((point_cloud.x, point_cloud.y,
                             point_cloud.z)).transpose()

transponded_colors = np.vstack((point_cloud.red, point_cloud.green,
                                point_cloud.blue)).transpose()

screening_size = int(args[4])

decimated_points = transponded_pts[::screening_size]
decimated_colors = transponded_colors[::screening_size]


print(plt.show())


header = lp.LasHeader(point_format=3, version="1.2")
header.add_extra_dim(lp.ExtraBytesParams(name="random", type=np.int32))
header.scales = np.array([0.1, 0.1, 0.1])


new_las_file = lp.create(
    point_format=header.point_format,
    file_version=header.version)


new_las_file.xyz = transponded_pts[::screening_size]

new_las_file.red = point_cloud.red[::screening_size]
new_las_file.green = point_cloud.green[::screening_size]
new_las_file.blue = point_cloud.blue[::screening_size]


print(new_las_file.points)


file_save_string = args[2] + args[3]

new_las_file.write(file_save_string+'.las')
