package server

import (
	"context"
	"errors"
	"habits/api"
	"habits/api/pkg/habit"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateHabit(ctx context.Context, request *api.CreateHabitRequest) (*api.CreateHabitResponse, error) {
	var freq uint
	if request.WeeklyFrequency != nil {
		freq = uint(*request.WeeklyFrequency)
	}

	h := habit.Habit {
		Name: habit.Name(request.Name),
		WeeklyFrequency: habit.WeeklyFrequency(freq),
	}

	createdHabit, err := habit.Create(ctx, s.db, h)
	if err != nil {
		var invalidErr habit.InvalidInputError
		if errors.As(err, &invalidErr) {
			return nil, status.Error(codes.InvalidArgument, invalidErr.Error())
		}
		return nil, status.Errorf(codes.Internal, "cannot save habit %v: %s", h, err.Error())
	}

	s.lgr.Logf("Habit %s successfully registered", createdHabit.ID)

	return &api.CreateHabitResponse{
		Habit: &api.Habit{
			Id: string(createdHabit.ID),
			Name: string(createdHabit.Name),
			WeeklyFrequency: int32(createdHabit.WeeklyFrequency),
		},
	}, nil
}