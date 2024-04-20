package server

import (
	"context"
	"log"
	db "usermanagmentService/database"
	proto "usermanagmentService/proto"
)

func (s *UserManagementServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.RegisterResponse, error) {
	log.Println("Register user invoked")
	err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Username:     req.Username,
		PasswordHash: req.Password,
		Role:         req.Role,
	})
	if err != nil {
		return &proto.RegisterResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to create user",
			},
		}, err
	}

	return &proto.RegisterResponse{
		Response: &proto.Response{
			Success: true,
			Message: "User created successfully",
		},
	}, nil
}

func (s *UserManagementServer) UpdateUser(ctx context.Context, req *proto.UpdateUserRequest) (*proto.UpdateUserResponse, error) {
	log.Println("Update user invoked", req.User)
	err := s.queries.UpdateUser(ctx, db.UpdateUserParams{
		UserID:       req.User.UserId,
		Username:     req.User.Username,
		PasswordHash: req.User.PasswordHash,
		Role: req.User.Role,
	})
	if err != nil {
		return &proto.UpdateUserResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to update user",
			},
		}, err
	}

	return &proto.UpdateUserResponse{
		Response: &proto.Response{
			Success: true,
			Message: "User updated successfully",
		},
	}, nil
}

func (s *UserManagementServer) GetUser(ctx context.Context, req *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	log.Println("Get user invoked")
	user, err := s.queries.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserResponse{
		User: &proto.UserResponse{
			UserId:   user.UserID,
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}

func (s *UserManagementServer) GetUsers(ctx context.Context, req *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	log.Println("Get users invoked")
	users, err := s.queries.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	var userResponses []*proto.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, &proto.UserResponse{
			UserId:   user.UserID,
			Username: user.Username,
			Role:     user.Role,
		})
	}
	return &proto.GetUsersResponse{
		Users: userResponses,
	}, nil
}

func (s *UserManagementServer) GetUserByUsername(ctx context.Context, req *proto.GetUserByUsernameRequest) (*proto.GetUserByUsernameResponse, error) {
	log.Println("Get user by username invoked")
	user, err := s.queries.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &proto.GetUserByUsernameResponse{
		User: &proto.UserResponse{
			UserId:   user.UserID,
			Username: user.Username,
			Role:     user.Role,
		},
	}, nil
}

func (s *UserManagementServer) DeleteUser(ctx context.Context, req *proto.DeleteUserRequest) (*proto.DeleteUserResponse, error) {
	log.Println("Delete user invoked")
	err := s.queries.DeleteUser(ctx, req.UserId)
	if err != nil {
		return &proto.DeleteUserResponse{
			Response: &proto.Response{
				Success: false,
				Message: "Failed to delete user",
			},
		}, err
	}

	return &proto.DeleteUserResponse{
		Response: &proto.Response{
			Success: true,
			Message: "User deleted successfully",
		},
	}, nil
}
