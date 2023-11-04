// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecase

import (
	"context"
	"github.com/kngnkg/tunetrail/backend/entity"
	"github.com/kngnkg/tunetrail/backend/infra/repository"
	"sync"
)

// Ensure, that UserRepositoryMock does implement UserRepository.
// If this is not the case, regenerate this file with moq.
var _ UserRepository = &UserRepositoryMock{}

// UserRepositoryMock is a mock implementation of UserRepository.
//
//	func TestSomethingThatUsesUserRepository(t *testing.T) {
//
//		// make and configure a mocked UserRepository
//		mockedUserRepository := &UserRepositoryMock{
//			GetUserByImmutableIdFunc: func(ctx context.Context, db repository.Executor, immutableId entity.ImmutableId) (*entity.User, error) {
//				panic("mock out the GetUserByImmutableId method")
//			},
//			GetUserByUsernameFunc: func(ctx context.Context, db repository.Executor, username entity.Username) (*entity.User, error) {
//				panic("mock out the GetUserByUsername method")
//			},
//			ListUsersFunc: func(ctx context.Context, db repository.Executor, immutableId entity.ImmutableId, limit int) ([]*entity.User, error) {
//				panic("mock out the ListUsers method")
//			},
//			ListUsersByIdFunc: func(ctx context.Context, db repository.Executor, userIds []entity.ImmutableId) ([]*entity.User, error) {
//				panic("mock out the ListUsersById method")
//			},
//			StoreUserFunc: func(ctx context.Context, db repository.Executor, user *entity.User) (*entity.User, error) {
//				panic("mock out the StoreUser method")
//			},
//		}
//
//		// use mockedUserRepository in code that requires UserRepository
//		// and then make assertions.
//
//	}
type UserRepositoryMock struct {
	// GetUserByImmutableIdFunc mocks the GetUserByImmutableId method.
	GetUserByImmutableIdFunc func(ctx context.Context, db repository.Executor, immutableId entity.ImmutableId) (*entity.User, error)

	// GetUserByUsernameFunc mocks the GetUserByUsername method.
	GetUserByUsernameFunc func(ctx context.Context, db repository.Executor, username entity.Username) (*entity.User, error)

	// ListUsersFunc mocks the ListUsers method.
	ListUsersFunc func(ctx context.Context, db repository.Executor, immutableId entity.ImmutableId, limit int) ([]*entity.User, error)

	// ListUsersByIdFunc mocks the ListUsersById method.
	ListUsersByIdFunc func(ctx context.Context, db repository.Executor, userIds []entity.ImmutableId) ([]*entity.User, error)

	// StoreUserFunc mocks the StoreUser method.
	StoreUserFunc func(ctx context.Context, db repository.Executor, user *entity.User) (*entity.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// GetUserByImmutableId holds details about calls to the GetUserByImmutableId method.
		GetUserByImmutableId []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Executor
			// ImmutableId is the immutableId argument value.
			ImmutableId entity.ImmutableId
		}
		// GetUserByUsername holds details about calls to the GetUserByUsername method.
		GetUserByUsername []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Executor
			// Username is the username argument value.
			Username entity.Username
		}
		// ListUsers holds details about calls to the ListUsers method.
		ListUsers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Executor
			// ImmutableId is the immutableId argument value.
			ImmutableId entity.ImmutableId
			// Limit is the limit argument value.
			Limit int
		}
		// ListUsersById holds details about calls to the ListUsersById method.
		ListUsersById []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Executor
			// UserIds is the userIds argument value.
			UserIds []entity.ImmutableId
		}
		// StoreUser holds details about calls to the StoreUser method.
		StoreUser []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Db is the db argument value.
			Db repository.Executor
			// User is the user argument value.
			User *entity.User
		}
	}
	lockGetUserByImmutableId sync.RWMutex
	lockGetUserByUsername    sync.RWMutex
	lockListUsers            sync.RWMutex
	lockListUsersById        sync.RWMutex
	lockStoreUser            sync.RWMutex
}

// GetUserByImmutableId calls GetUserByImmutableIdFunc.
func (mock *UserRepositoryMock) GetUserByImmutableId(ctx context.Context, db repository.Executor, immutableId entity.ImmutableId) (*entity.User, error) {
	if mock.GetUserByImmutableIdFunc == nil {
		panic("UserRepositoryMock.GetUserByImmutableIdFunc: method is nil but UserRepository.GetUserByImmutableId was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Db          repository.Executor
		ImmutableId entity.ImmutableId
	}{
		Ctx:         ctx,
		Db:          db,
		ImmutableId: immutableId,
	}
	mock.lockGetUserByImmutableId.Lock()
	mock.calls.GetUserByImmutableId = append(mock.calls.GetUserByImmutableId, callInfo)
	mock.lockGetUserByImmutableId.Unlock()
	return mock.GetUserByImmutableIdFunc(ctx, db, immutableId)
}

// GetUserByImmutableIdCalls gets all the calls that were made to GetUserByImmutableId.
// Check the length with:
//
//	len(mockedUserRepository.GetUserByImmutableIdCalls())
func (mock *UserRepositoryMock) GetUserByImmutableIdCalls() []struct {
	Ctx         context.Context
	Db          repository.Executor
	ImmutableId entity.ImmutableId
} {
	var calls []struct {
		Ctx         context.Context
		Db          repository.Executor
		ImmutableId entity.ImmutableId
	}
	mock.lockGetUserByImmutableId.RLock()
	calls = mock.calls.GetUserByImmutableId
	mock.lockGetUserByImmutableId.RUnlock()
	return calls
}

// GetUserByUsername calls GetUserByUsernameFunc.
func (mock *UserRepositoryMock) GetUserByUsername(ctx context.Context, db repository.Executor, username entity.Username) (*entity.User, error) {
	if mock.GetUserByUsernameFunc == nil {
		panic("UserRepositoryMock.GetUserByUsernameFunc: method is nil but UserRepository.GetUserByUsername was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Db       repository.Executor
		Username entity.Username
	}{
		Ctx:      ctx,
		Db:       db,
		Username: username,
	}
	mock.lockGetUserByUsername.Lock()
	mock.calls.GetUserByUsername = append(mock.calls.GetUserByUsername, callInfo)
	mock.lockGetUserByUsername.Unlock()
	return mock.GetUserByUsernameFunc(ctx, db, username)
}

// GetUserByUsernameCalls gets all the calls that were made to GetUserByUsername.
// Check the length with:
//
//	len(mockedUserRepository.GetUserByUsernameCalls())
func (mock *UserRepositoryMock) GetUserByUsernameCalls() []struct {
	Ctx      context.Context
	Db       repository.Executor
	Username entity.Username
} {
	var calls []struct {
		Ctx      context.Context
		Db       repository.Executor
		Username entity.Username
	}
	mock.lockGetUserByUsername.RLock()
	calls = mock.calls.GetUserByUsername
	mock.lockGetUserByUsername.RUnlock()
	return calls
}

// ListUsers calls ListUsersFunc.
func (mock *UserRepositoryMock) ListUsers(ctx context.Context, db repository.Executor, immutableId entity.ImmutableId, limit int) ([]*entity.User, error) {
	if mock.ListUsersFunc == nil {
		panic("UserRepositoryMock.ListUsersFunc: method is nil but UserRepository.ListUsers was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		Db          repository.Executor
		ImmutableId entity.ImmutableId
		Limit       int
	}{
		Ctx:         ctx,
		Db:          db,
		ImmutableId: immutableId,
		Limit:       limit,
	}
	mock.lockListUsers.Lock()
	mock.calls.ListUsers = append(mock.calls.ListUsers, callInfo)
	mock.lockListUsers.Unlock()
	return mock.ListUsersFunc(ctx, db, immutableId, limit)
}

// ListUsersCalls gets all the calls that were made to ListUsers.
// Check the length with:
//
//	len(mockedUserRepository.ListUsersCalls())
func (mock *UserRepositoryMock) ListUsersCalls() []struct {
	Ctx         context.Context
	Db          repository.Executor
	ImmutableId entity.ImmutableId
	Limit       int
} {
	var calls []struct {
		Ctx         context.Context
		Db          repository.Executor
		ImmutableId entity.ImmutableId
		Limit       int
	}
	mock.lockListUsers.RLock()
	calls = mock.calls.ListUsers
	mock.lockListUsers.RUnlock()
	return calls
}

// ListUsersById calls ListUsersByIdFunc.
func (mock *UserRepositoryMock) ListUsersById(ctx context.Context, db repository.Executor, userIds []entity.ImmutableId) ([]*entity.User, error) {
	if mock.ListUsersByIdFunc == nil {
		panic("UserRepositoryMock.ListUsersByIdFunc: method is nil but UserRepository.ListUsersById was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Db      repository.Executor
		UserIds []entity.ImmutableId
	}{
		Ctx:     ctx,
		Db:      db,
		UserIds: userIds,
	}
	mock.lockListUsersById.Lock()
	mock.calls.ListUsersById = append(mock.calls.ListUsersById, callInfo)
	mock.lockListUsersById.Unlock()
	return mock.ListUsersByIdFunc(ctx, db, userIds)
}

// ListUsersByIdCalls gets all the calls that were made to ListUsersById.
// Check the length with:
//
//	len(mockedUserRepository.ListUsersByIdCalls())
func (mock *UserRepositoryMock) ListUsersByIdCalls() []struct {
	Ctx     context.Context
	Db      repository.Executor
	UserIds []entity.ImmutableId
} {
	var calls []struct {
		Ctx     context.Context
		Db      repository.Executor
		UserIds []entity.ImmutableId
	}
	mock.lockListUsersById.RLock()
	calls = mock.calls.ListUsersById
	mock.lockListUsersById.RUnlock()
	return calls
}

// StoreUser calls StoreUserFunc.
func (mock *UserRepositoryMock) StoreUser(ctx context.Context, db repository.Executor, user *entity.User) (*entity.User, error) {
	if mock.StoreUserFunc == nil {
		panic("UserRepositoryMock.StoreUserFunc: method is nil but UserRepository.StoreUser was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Db   repository.Executor
		User *entity.User
	}{
		Ctx:  ctx,
		Db:   db,
		User: user,
	}
	mock.lockStoreUser.Lock()
	mock.calls.StoreUser = append(mock.calls.StoreUser, callInfo)
	mock.lockStoreUser.Unlock()
	return mock.StoreUserFunc(ctx, db, user)
}

// StoreUserCalls gets all the calls that were made to StoreUser.
// Check the length with:
//
//	len(mockedUserRepository.StoreUserCalls())
func (mock *UserRepositoryMock) StoreUserCalls() []struct {
	Ctx  context.Context
	Db   repository.Executor
	User *entity.User
} {
	var calls []struct {
		Ctx  context.Context
		Db   repository.Executor
		User *entity.User
	}
	mock.lockStoreUser.RLock()
	calls = mock.calls.StoreUser
	mock.lockStoreUser.RUnlock()
	return calls
}
