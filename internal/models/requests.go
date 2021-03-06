package models

type SignUpRequest struct {
	Email    string
	Username string
	Password string
}

type SignInRequest struct {
	Email    string
	Password string
}

type ProjectUpdateRequest struct {
	NewName string `json:"new_project_name"`
	NewInfo string `json:"new_info"`
}

type ProjectUploadRequest struct {
	ProjectName string `json:"project_name"`
}

type UserUpdateRequest struct {
	NewUsername string `json:"new_username"`
	NewEmail    string `json:"new_email"`
}

type ProcessingRandRequest struct {
	ProjectName string `json:"project_name"`
	UserId      uint64 `json:"user_id"`
	FilePath    string `json:"file_path"`
	Factor      uint64 `json:"factor"`
}
type ProcessingGridRequest struct {
	ProjectName string `json:"project_name"`
	UserId      uint64 `json:"user_id"`
	FilePath    string `json:"file_path"`
	Voxel       uint64 `json:"voxel_size"`
}

type GetPointsAmountRequest struct {
	ProjectName string `json:"project_name"`
	UserId      uint64 `json:"user_id"`
	FilePath    string `json:"file_path"`
}
