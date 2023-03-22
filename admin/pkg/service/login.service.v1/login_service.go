package login_service_v1

import (
	"context"
	"github.com/blog-admin/config"
	"github.com/blog-admin/internal/dao"
	"github.com/blog-admin/internal/database/tran"
	"github.com/blog-admin/internal/repo"
	"github.com/blog-admin/pkg/model"
	"github.com/blog-common/encrypts"
	"github.com/blog-common/errs"
	"github.com/blog-common/jwts"
	"github.com/blog-grpc/admin/login"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

type LoginService struct {
	login.UnimplementedLoginServiceServer
	cache       repo.Cache
	userRepo    repo.UserRepo
	transaction tran.Transaction
}

func New() *LoginService {
	return &LoginService{
		cache:       dao.Rc,
		userRepo:    dao.NewUserDao(),
		transaction: dao.NewTransaction(),
	}
}

func (l *LoginService) Login(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	c := context.Background()
	pwd := encrypts.Md5(msg.Password)
	user, err := l.userRepo.FindMember(c, msg.Username, pwd)
	if err != nil {
		zap.L().Error("Login db FindMember error", zap.Error(err))
		return nil, errs.GrpcError(model.DBError)
	}
	if user == nil {
		return nil, errs.GrpcError(model.AccountAndPwdError)
	}
	userMessage := &login.UserMessage{}
	err = copier.Copy(userMessage, user)
	memIdStr := strconv.FormatInt(user.Id, 10)
	exp := time.Duration(config.C.JwtConfig.AccessExp*3600*24) * time.Second
	rExp := time.Duration(config.C.JwtConfig.RefreshExp*3600*24) * time.Second
	token := jwts.CreateToken(memIdStr, exp, config.C.JwtConfig.AccessSecret, rExp, config.C.JwtConfig.RefreshSecret, msg.Ip)
	tokenList := &login.TokenMessage{
		AccessToken:    token.AccessToken,
		RefreshToken:   token.RefreshToken,
		AccessTokenExp: token.AccessExp,
		TokenType:      "bearer",
	}
	return &login.LoginResponse{
		User:  userMessage,
		Token: tokenList,
	}, nil
}
func (l *LoginService) TokenVerify(ctx context.Context, msg *login.LoginMessage) (*login.LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenVerify not implemented")
}
func (l *LoginService) FindMemInfoById(ctx context.Context, msg *login.UserQueryMessage) (*login.UserMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMemInfoById not implemented")
}
func (l *LoginService) FindMemInfoByIds(ctx context.Context, msg *login.UserQueryMessage) (*login.UserMessageList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindMemInfoByIds not implemented")
}
