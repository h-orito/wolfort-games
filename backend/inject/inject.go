package inject

import (
	"net/http"
	"wolfort-games/adaptor/auth"
	"wolfort-games/adaptor/graphql"
	"wolfort-games/application/app_service"
	"wolfort-games/application/usecase"
	"wolfort-games/domain/dom_service"
	"wolfort-games/domain/model"
	db "wolfort-games/infrastructure/rdb"
	"wolfort-games/middleware/auth0"
	"wolfort-games/middleware/graph"

	"github.com/99designs/gqlgen/graphql/handler"
)

func InjectServer() http.Handler {
	database := injectDb()
	userRepository := injectUserRepository(database)
	resolver := injectResolver(database, userRepository)
	srv := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers:  &resolver,
				Directives: graphql.Directive,
			},
		),
	)
	handler := auth.AuthMiddleware(srv, userRepository)
	return auth0.JwtMiddleware()(handler)
}

// resolver
func injectResolver(
	database db.DB,
	userRepository model.UserRepository,
) graphql.Resolver {
	tx := db.NewTransaction(database.Connection)
	// repository
	chinchiroRoomRepository := injectChinchiroRoomRepository(database)
	chinchiroRoomParticipantRepository := injectChinchiroRoomParticipantRepository(database)
	chinchiroGameRepository := injectChinchiroGameRepository(database)
	playerRepository := injectPlayerRepository(database)
	// domain service
	chinchiroRoomMasterDomainService := injectChinchiroRoomMasterDomainService()
	chinchiroRoomParticipateDomainService := injectChinchiroRoomParticipateDomainService(chinchiroRoomMasterDomainService)
	chinchiroGameDomainService := injectChinchiroGameDomainService()
	// application service
	chinchiroRoomService := injectChinchiroRoomService(chinchiroRoomRepository, chinchiroRoomParticipantRepository)
	chinchiroGameService := injectChinchiroGameService(chinchiroRoomParticipantRepository, chinchiroGameRepository, chinchiroGameDomainService)
	playerService := injectPlayerService(playerRepository, userRepository)
	// usecase
	chinchiroRoomUsecase := injectChinchiroRoomUsecase(
		chinchiroRoomService,
		playerService,
		chinchiroRoomMasterDomainService,
		chinchiroRoomParticipateDomainService,
		tx,
	)
	chinchiroGameUsecase := injectChinchiroGameUsecase(
		chinchiroRoomService,
		chinchiroGameService,
		playerService,
		chinchiroRoomMasterDomainService,
		tx,
	)
	playerUsecase := injectPlayerUsecase(playerService, tx)
	loaders := injectLoaders(playerUsecase, chinchiroRoomUsecase, chinchiroGameUsecase)
	return graphql.NewResolver(
		playerUsecase,
		chinchiroRoomUsecase,
		chinchiroGameUsecase,
		loaders,
	)
}

// loader
func injectLoaders(
	playerUsecase usecase.PlayerUsecase,
	chinchiroRoomUsecase usecase.ChinchiroRoomUsecase,
	chinchiroGameUsecase usecase.ChinchiroGameUsecase,
) *graphql.Loaders {
	return graphql.NewLoaders(playerUsecase, chinchiroRoomUsecase, chinchiroGameUsecase)
}

// usecase
func injectChinchiroRoomUsecase(
	chinchiroRoomService app_service.ChinchiroRoomService,
	playerService app_service.PlayerService,
	chinchiroRoomMasterDomainService dom_service.ChinchiroRoomMasterDomainService,
	chinchiroRoomParticipateDomainService dom_service.ChinchiroRoomParticipateDomainService,
	tx usecase.Transaction,
) usecase.ChinchiroRoomUsecase {
	return usecase.NewChinchiroRoomUsecase(
		chinchiroRoomService,
		playerService,
		chinchiroRoomMasterDomainService,
		chinchiroRoomParticipateDomainService,
		tx,
	)
}

func injectChinchiroGameUsecase(
	chinchiroRoomService app_service.ChinchiroRoomService,
	chinchiroGameService app_service.ChinchiroGameService,
	playerService app_service.PlayerService,
	chinchiroRoomMasterDomainService dom_service.ChinchiroRoomMasterDomainService,
	tx usecase.Transaction,
) usecase.ChinchiroGameUsecase {
	return usecase.NewChinchiroGameUsecase(
		chinchiroRoomService,
		chinchiroGameService,
		playerService,
		chinchiroRoomMasterDomainService,
		tx,
	)
}

func injectPlayerUsecase(playerService app_service.PlayerService,
	tx usecase.Transaction,
) usecase.PlayerUsecase {
	return usecase.NewPlayerUsecase(playerService, tx)
}

// service
func injectChinchiroRoomService(
	chinchiroRoomRepository model.ChinchiroRoomRepository,
	chinchiroRoomParticipantRepository model.ChinchiroRoomParticipantRepository,
) app_service.ChinchiroRoomService {
	return app_service.NewChinchiroRoomService(chinchiroRoomRepository, chinchiroRoomParticipantRepository)
}

func injectChinchiroGameService(
	chinchiroRoomParticipantRepository model.ChinchiroRoomParticipantRepository,
	chinchiroGameRepository model.ChinchiroGameRepository,
	chinchiroGameDomainService dom_service.ChinchiroGameDomainService,
) app_service.ChinchiroGameService {
	return app_service.NewChinchiroGameService(chinchiroRoomParticipantRepository, chinchiroGameRepository, chinchiroGameDomainService)
}

func injectPlayerService(
	playerRepository model.PlayerRepository,
	userRepository model.UserRepository,
) app_service.PlayerService {
	return app_service.NewPlayerService(playerRepository, userRepository)
}

// domain service
func injectChinchiroRoomMasterDomainService() dom_service.ChinchiroRoomMasterDomainService {
	return dom_service.NewChinchiroRoomMasterDomainService()
}

func injectChinchiroRoomParticipateDomainService(
	masterDomainService dom_service.ChinchiroRoomMasterDomainService,
) dom_service.ChinchiroRoomParticipateDomainService {
	return dom_service.NewChinchiroRoomParticipateDomainService(masterDomainService)
}

func injectChinchiroGameDomainService() dom_service.ChinchiroGameDomainService {
	return dom_service.NewChinchiroGameDomainService()
}

// repository
func injectChinchiroRoomRepository(database db.DB) model.ChinchiroRoomRepository {
	return db.NewChinchiroRoomRepository(&database)
}

func injectChinchiroRoomParticipantRepository(database db.DB) model.ChinchiroRoomParticipantRepository {
	return db.NewChinchiroRoomParticipantRepository(&database)
}

func injectChinchiroGameRepository(database db.DB) model.ChinchiroGameRepository {
	return db.NewChinchiroGameRepository(&database)
}

func injectPlayerRepository(database db.DB) model.PlayerRepository {
	return db.NewPlayerRepository(&database)
}

func injectUserRepository(database db.DB) model.UserRepository {
	return db.NewUserRepository(&database)
}

// database
func injectDb() db.DB {
	return db.NewDB()
}
