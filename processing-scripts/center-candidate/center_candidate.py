# %%
import sys

import numpy as np
import laspy as lp


args = sys.argv
print(args)

point_cloud = lp.read(args[1])

voxel_size = int(args[4])


transponded_pts = np.vstack((point_cloud.x, point_cloud.y,
                             point_cloud.z)).transpose()

nb_vox = np.ceil((np.max(transponded_pts, axis=0) -
                 np.min(transponded_pts, axis=0))/voxel_size)

non_empty_voxel_keys, inverse, points_per_voxel = np.unique(
    ((transponded_pts - np.min(transponded_pts, axis=0)) // voxel_size).astype(int),
    axis=0, return_inverse=True, return_counts=True)

idx_pts_vox_sorted = np.argsort(inverse)

grid = {}
candidate_center_list = []
last = 0

for index, voxel in enumerate(non_empty_voxel_keys):
    grid[tuple(
        voxel)] = transponded_pts[idx_pts_vox_sorted[last:last+points_per_voxel[index]]]

    candidate_center_list.append(grid[tuple(voxel)][np.linalg.norm(
        grid[tuple(voxel)] - np.mean(grid[tuple(voxel)], axis=0), axis=1).argmin()])
    last += points_per_voxel[index]

sampled = candidate_center_list

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
