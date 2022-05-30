# %%
import sys

import numpy as np
import laspy as lp
import matplotlib.pyplot as plt


args = sys.argv
print(args)

point_cloud = lp.read(sys.argv[1])

voxel_size = int(sys.argv[4])


points = np.vstack((point_cloud.x, point_cloud.y,
                    point_cloud.z)).transpose()

nb_vox = np.ceil((np.max(points, axis=0) - np.min(points, axis=0))/voxel_size)

non_empty_voxel_keys, inverse, nb_pts_per_voxel = np.unique(
    ((points - np.min(points, axis=0)) // voxel_size).astype(int), axis=0, return_inverse=True, return_counts=True)

idx_pts_vox_sorted = np.argsort(inverse)

voxel_grid = {}
grid_barycenter = []
last_seen = 0

for index, voxel in enumerate(non_empty_voxel_keys):
    voxel_grid[tuple(
        voxel)] = points[idx_pts_vox_sorted[last_seen:last_seen+nb_pts_per_voxel[index]]]
    grid_barycenter.append(np.mean(voxel_grid[tuple(voxel)], axis=0))

    last_seen += nb_pts_per_voxel[index]

sampled = grid_barycenter

header = lp.LasHeader(point_format=3, version="1.2")
header.add_extra_dim(lp.ExtraBytesParams(name="random", type=np.int32))
header.scales = np.array([0.1, 0.1, 0.1])


new_las_file = lp.create(
    point_format=header.point_format,
    file_version=header.version)


new_las_file.xyz = sampled


print(new_las_file.points)


file_save_string = args[2] + args[3]

new_las_file.write(file_save_string+'.las')
