package rewards_handler

type rewardsUsecase interface {
}

type rewardsHandler struct {
	rewardsUsecase rewardsUsecase
}

func NewRewardsHandler(rewardsUsecase rewardsUsecase) *rewardsHandler {
	return &rewardsHandler{
		rewardsUsecase: rewardsUsecase,
	}
}
