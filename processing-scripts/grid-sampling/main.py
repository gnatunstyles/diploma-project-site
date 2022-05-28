import numpy as np
import laspy as lp
import matplotlib.pyplot as plt

point_cloud = lp.read(
    "cloud.las")

voxel_size = 2


points = np.vstack((point_cloud.x, point_cloud.y,
                    point_cloud.z)).transpose()

# colors = np.vstack((point_cloud.red, point_cloud.green,
#                    point_cloud.blue)).transpose()

nb_vox = np.ceil((np.max(points, axis=0) - np.min(points, axis=0))/voxel_size)

non_empty_voxel_keys, inverse, nb_pts_per_voxel = np.unique(
    ((points - np.min(points, axis=0)) // voxel_size).astype(int), axis=0, return_inverse=True, return_counts=True)

idx_pts_vox_sorted = np.argsort(inverse)

voxel_grid = {}
grid_barycenter, grid_candidate_center = [], []
last_seen = 0

for idx, vox in enumerate(non_empty_voxel_keys):
    voxel_grid[tuple(vox)] = points[idx_pts_vox_sorted[
        last_seen:last_seen+nb_pts_per_voxel[idx]]]

    grid_barycenter.append(np.mean(voxel_grid[tuple(vox)], axis=0))
    grid_candidate_center.append(
        voxel_grid[tuple(vox)][np.linalg.norm(voxel_grid[tuple(vox)] -
                                              np.mean(voxel_grid[tuple(vox)], axis=0), axis=1).argmin()])
    last_seen += nb_pts_per_voxel[idx]

sampled = grid_candidate_center

header = lp.LasHeader(point_format=3, version="1.2")
header.add_extra_dim(lp.ExtraBytesParams(name="random", type=np.int32))
header.scales = np.array([0.1, 0.1, 0.1])


new_las_file = lp.create(
    point_format=header.point_format,
    file_version=header.version)


new_las_file.xyz = sampled


# new_las_file.red = point_cloud.red[::10]
# new_las_file.green = point_cloud.green[::10]
# new_las_file.blue = point_cloud.blue[::10]


print(new_las_file.points)

new_las_file.write('cloud_vox.las')
