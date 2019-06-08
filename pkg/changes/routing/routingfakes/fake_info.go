// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	sync "sync"

	routing "github.com/calebamiles/keps/pkg/changes/routing"
)

type FakeInfo struct {
	ChangeReceiptStub        func() string
	changeReceiptMutex       sync.RWMutex
	changeReceiptArgsForCall []struct {
	}
	changeReceiptReturns struct {
		result1 string
	}
	changeReceiptReturnsOnCall map[int]struct {
		result1 string
	}
	ChangesUnderPathStub        func() string
	changesUnderPathMutex       sync.RWMutex
	changesUnderPathArgsForCall []struct {
	}
	changesUnderPathReturns struct {
		result1 string
	}
	changesUnderPathReturnsOnCall map[int]struct {
		result1 string
	}
	FullDescriptionStub        func() string
	fullDescriptionMutex       sync.RWMutex
	fullDescriptionArgsForCall []struct {
	}
	fullDescriptionReturns struct {
		result1 string
	}
	fullDescriptionReturnsOnCall map[int]struct {
		result1 string
	}
	OwningSIGStub        func() string
	owningSIGMutex       sync.RWMutex
	owningSIGArgsForCall []struct {
	}
	owningSIGReturns struct {
		result1 string
	}
	owningSIGReturnsOnCall map[int]struct {
		result1 string
	}
	PathToRepoStub        func() string
	pathToRepoMutex       sync.RWMutex
	pathToRepoArgsForCall []struct {
	}
	pathToRepoReturns struct {
		result1 string
	}
	pathToRepoReturnsOnCall map[int]struct {
		result1 string
	}
	PrincipalEmailStub        func() string
	principalEmailMutex       sync.RWMutex
	principalEmailArgsForCall []struct {
	}
	principalEmailReturns struct {
		result1 string
	}
	principalEmailReturnsOnCall map[int]struct {
		result1 string
	}
	PrincipalNameStub        func() string
	principalNameMutex       sync.RWMutex
	principalNameArgsForCall []struct {
	}
	principalNameReturns struct {
		result1 string
	}
	principalNameReturnsOnCall map[int]struct {
		result1 string
	}
	ShortSummaryStub        func() string
	shortSummaryMutex       sync.RWMutex
	shortSummaryArgsForCall []struct {
	}
	shortSummaryReturns struct {
		result1 string
	}
	shortSummaryReturnsOnCall map[int]struct {
		result1 string
	}
	SourceBranchStub        func() string
	sourceBranchMutex       sync.RWMutex
	sourceBranchArgsForCall []struct {
	}
	sourceBranchReturns struct {
		result1 string
	}
	sourceBranchReturnsOnCall map[int]struct {
		result1 string
	}
	SourceRepositoryStub        func() string
	sourceRepositoryMutex       sync.RWMutex
	sourceRepositoryArgsForCall []struct {
	}
	sourceRepositoryReturns struct {
		result1 string
	}
	sourceRepositoryReturnsOnCall map[int]struct {
		result1 string
	}
	SourceRepositoryOwnerStub        func() string
	sourceRepositoryOwnerMutex       sync.RWMutex
	sourceRepositoryOwnerArgsForCall []struct {
	}
	sourceRepositoryOwnerReturns struct {
		result1 string
	}
	sourceRepositoryOwnerReturnsOnCall map[int]struct {
		result1 string
	}
	TargetBranchStub        func() string
	targetBranchMutex       sync.RWMutex
	targetBranchArgsForCall []struct {
	}
	targetBranchReturns struct {
		result1 string
	}
	targetBranchReturnsOnCall map[int]struct {
		result1 string
	}
	TargetRepositoryStub        func() string
	targetRepositoryMutex       sync.RWMutex
	targetRepositoryArgsForCall []struct {
	}
	targetRepositoryReturns struct {
		result1 string
	}
	targetRepositoryReturnsOnCall map[int]struct {
		result1 string
	}
	TargetRepositoryOwnerStub        func() string
	targetRepositoryOwnerMutex       sync.RWMutex
	targetRepositoryOwnerArgsForCall []struct {
	}
	targetRepositoryOwnerReturns struct {
		result1 string
	}
	targetRepositoryOwnerReturnsOnCall map[int]struct {
		result1 string
	}
	TitleStub        func() string
	titleMutex       sync.RWMutex
	titleArgsForCall []struct {
	}
	titleReturns struct {
		result1 string
	}
	titleReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInfo) ChangeReceipt() string {
	fake.changeReceiptMutex.Lock()
	ret, specificReturn := fake.changeReceiptReturnsOnCall[len(fake.changeReceiptArgsForCall)]
	fake.changeReceiptArgsForCall = append(fake.changeReceiptArgsForCall, struct {
	}{})
	fake.recordInvocation("ChangeReceipt", []interface{}{})
	fake.changeReceiptMutex.Unlock()
	if fake.ChangeReceiptStub != nil {
		return fake.ChangeReceiptStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.changeReceiptReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) ChangeReceiptCallCount() int {
	fake.changeReceiptMutex.RLock()
	defer fake.changeReceiptMutex.RUnlock()
	return len(fake.changeReceiptArgsForCall)
}

func (fake *FakeInfo) ChangeReceiptCalls(stub func() string) {
	fake.changeReceiptMutex.Lock()
	defer fake.changeReceiptMutex.Unlock()
	fake.ChangeReceiptStub = stub
}

func (fake *FakeInfo) ChangeReceiptReturns(result1 string) {
	fake.changeReceiptMutex.Lock()
	defer fake.changeReceiptMutex.Unlock()
	fake.ChangeReceiptStub = nil
	fake.changeReceiptReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) ChangeReceiptReturnsOnCall(i int, result1 string) {
	fake.changeReceiptMutex.Lock()
	defer fake.changeReceiptMutex.Unlock()
	fake.ChangeReceiptStub = nil
	if fake.changeReceiptReturnsOnCall == nil {
		fake.changeReceiptReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.changeReceiptReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) ChangesUnderPath() string {
	fake.changesUnderPathMutex.Lock()
	ret, specificReturn := fake.changesUnderPathReturnsOnCall[len(fake.changesUnderPathArgsForCall)]
	fake.changesUnderPathArgsForCall = append(fake.changesUnderPathArgsForCall, struct {
	}{})
	fake.recordInvocation("ChangesUnderPath", []interface{}{})
	fake.changesUnderPathMutex.Unlock()
	if fake.ChangesUnderPathStub != nil {
		return fake.ChangesUnderPathStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.changesUnderPathReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) ChangesUnderPathCallCount() int {
	fake.changesUnderPathMutex.RLock()
	defer fake.changesUnderPathMutex.RUnlock()
	return len(fake.changesUnderPathArgsForCall)
}

func (fake *FakeInfo) ChangesUnderPathCalls(stub func() string) {
	fake.changesUnderPathMutex.Lock()
	defer fake.changesUnderPathMutex.Unlock()
	fake.ChangesUnderPathStub = stub
}

func (fake *FakeInfo) ChangesUnderPathReturns(result1 string) {
	fake.changesUnderPathMutex.Lock()
	defer fake.changesUnderPathMutex.Unlock()
	fake.ChangesUnderPathStub = nil
	fake.changesUnderPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) ChangesUnderPathReturnsOnCall(i int, result1 string) {
	fake.changesUnderPathMutex.Lock()
	defer fake.changesUnderPathMutex.Unlock()
	fake.ChangesUnderPathStub = nil
	if fake.changesUnderPathReturnsOnCall == nil {
		fake.changesUnderPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.changesUnderPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) FullDescription() string {
	fake.fullDescriptionMutex.Lock()
	ret, specificReturn := fake.fullDescriptionReturnsOnCall[len(fake.fullDescriptionArgsForCall)]
	fake.fullDescriptionArgsForCall = append(fake.fullDescriptionArgsForCall, struct {
	}{})
	fake.recordInvocation("FullDescription", []interface{}{})
	fake.fullDescriptionMutex.Unlock()
	if fake.FullDescriptionStub != nil {
		return fake.FullDescriptionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.fullDescriptionReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) FullDescriptionCallCount() int {
	fake.fullDescriptionMutex.RLock()
	defer fake.fullDescriptionMutex.RUnlock()
	return len(fake.fullDescriptionArgsForCall)
}

func (fake *FakeInfo) FullDescriptionCalls(stub func() string) {
	fake.fullDescriptionMutex.Lock()
	defer fake.fullDescriptionMutex.Unlock()
	fake.FullDescriptionStub = stub
}

func (fake *FakeInfo) FullDescriptionReturns(result1 string) {
	fake.fullDescriptionMutex.Lock()
	defer fake.fullDescriptionMutex.Unlock()
	fake.FullDescriptionStub = nil
	fake.fullDescriptionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) FullDescriptionReturnsOnCall(i int, result1 string) {
	fake.fullDescriptionMutex.Lock()
	defer fake.fullDescriptionMutex.Unlock()
	fake.FullDescriptionStub = nil
	if fake.fullDescriptionReturnsOnCall == nil {
		fake.fullDescriptionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.fullDescriptionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) OwningSIG() string {
	fake.owningSIGMutex.Lock()
	ret, specificReturn := fake.owningSIGReturnsOnCall[len(fake.owningSIGArgsForCall)]
	fake.owningSIGArgsForCall = append(fake.owningSIGArgsForCall, struct {
	}{})
	fake.recordInvocation("OwningSIG", []interface{}{})
	fake.owningSIGMutex.Unlock()
	if fake.OwningSIGStub != nil {
		return fake.OwningSIGStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.owningSIGReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) OwningSIGCallCount() int {
	fake.owningSIGMutex.RLock()
	defer fake.owningSIGMutex.RUnlock()
	return len(fake.owningSIGArgsForCall)
}

func (fake *FakeInfo) OwningSIGCalls(stub func() string) {
	fake.owningSIGMutex.Lock()
	defer fake.owningSIGMutex.Unlock()
	fake.OwningSIGStub = stub
}

func (fake *FakeInfo) OwningSIGReturns(result1 string) {
	fake.owningSIGMutex.Lock()
	defer fake.owningSIGMutex.Unlock()
	fake.OwningSIGStub = nil
	fake.owningSIGReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) OwningSIGReturnsOnCall(i int, result1 string) {
	fake.owningSIGMutex.Lock()
	defer fake.owningSIGMutex.Unlock()
	fake.OwningSIGStub = nil
	if fake.owningSIGReturnsOnCall == nil {
		fake.owningSIGReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.owningSIGReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) PathToRepo() string {
	fake.pathToRepoMutex.Lock()
	ret, specificReturn := fake.pathToRepoReturnsOnCall[len(fake.pathToRepoArgsForCall)]
	fake.pathToRepoArgsForCall = append(fake.pathToRepoArgsForCall, struct {
	}{})
	fake.recordInvocation("PathToRepo", []interface{}{})
	fake.pathToRepoMutex.Unlock()
	if fake.PathToRepoStub != nil {
		return fake.PathToRepoStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.pathToRepoReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) PathToRepoCallCount() int {
	fake.pathToRepoMutex.RLock()
	defer fake.pathToRepoMutex.RUnlock()
	return len(fake.pathToRepoArgsForCall)
}

func (fake *FakeInfo) PathToRepoCalls(stub func() string) {
	fake.pathToRepoMutex.Lock()
	defer fake.pathToRepoMutex.Unlock()
	fake.PathToRepoStub = stub
}

func (fake *FakeInfo) PathToRepoReturns(result1 string) {
	fake.pathToRepoMutex.Lock()
	defer fake.pathToRepoMutex.Unlock()
	fake.PathToRepoStub = nil
	fake.pathToRepoReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) PathToRepoReturnsOnCall(i int, result1 string) {
	fake.pathToRepoMutex.Lock()
	defer fake.pathToRepoMutex.Unlock()
	fake.PathToRepoStub = nil
	if fake.pathToRepoReturnsOnCall == nil {
		fake.pathToRepoReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.pathToRepoReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) PrincipalEmail() string {
	fake.principalEmailMutex.Lock()
	ret, specificReturn := fake.principalEmailReturnsOnCall[len(fake.principalEmailArgsForCall)]
	fake.principalEmailArgsForCall = append(fake.principalEmailArgsForCall, struct {
	}{})
	fake.recordInvocation("PrincipalEmail", []interface{}{})
	fake.principalEmailMutex.Unlock()
	if fake.PrincipalEmailStub != nil {
		return fake.PrincipalEmailStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.principalEmailReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) PrincipalEmailCallCount() int {
	fake.principalEmailMutex.RLock()
	defer fake.principalEmailMutex.RUnlock()
	return len(fake.principalEmailArgsForCall)
}

func (fake *FakeInfo) PrincipalEmailCalls(stub func() string) {
	fake.principalEmailMutex.Lock()
	defer fake.principalEmailMutex.Unlock()
	fake.PrincipalEmailStub = stub
}

func (fake *FakeInfo) PrincipalEmailReturns(result1 string) {
	fake.principalEmailMutex.Lock()
	defer fake.principalEmailMutex.Unlock()
	fake.PrincipalEmailStub = nil
	fake.principalEmailReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) PrincipalEmailReturnsOnCall(i int, result1 string) {
	fake.principalEmailMutex.Lock()
	defer fake.principalEmailMutex.Unlock()
	fake.PrincipalEmailStub = nil
	if fake.principalEmailReturnsOnCall == nil {
		fake.principalEmailReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.principalEmailReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) PrincipalName() string {
	fake.principalNameMutex.Lock()
	ret, specificReturn := fake.principalNameReturnsOnCall[len(fake.principalNameArgsForCall)]
	fake.principalNameArgsForCall = append(fake.principalNameArgsForCall, struct {
	}{})
	fake.recordInvocation("PrincipalName", []interface{}{})
	fake.principalNameMutex.Unlock()
	if fake.PrincipalNameStub != nil {
		return fake.PrincipalNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.principalNameReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) PrincipalNameCallCount() int {
	fake.principalNameMutex.RLock()
	defer fake.principalNameMutex.RUnlock()
	return len(fake.principalNameArgsForCall)
}

func (fake *FakeInfo) PrincipalNameCalls(stub func() string) {
	fake.principalNameMutex.Lock()
	defer fake.principalNameMutex.Unlock()
	fake.PrincipalNameStub = stub
}

func (fake *FakeInfo) PrincipalNameReturns(result1 string) {
	fake.principalNameMutex.Lock()
	defer fake.principalNameMutex.Unlock()
	fake.PrincipalNameStub = nil
	fake.principalNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) PrincipalNameReturnsOnCall(i int, result1 string) {
	fake.principalNameMutex.Lock()
	defer fake.principalNameMutex.Unlock()
	fake.PrincipalNameStub = nil
	if fake.principalNameReturnsOnCall == nil {
		fake.principalNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.principalNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) ShortSummary() string {
	fake.shortSummaryMutex.Lock()
	ret, specificReturn := fake.shortSummaryReturnsOnCall[len(fake.shortSummaryArgsForCall)]
	fake.shortSummaryArgsForCall = append(fake.shortSummaryArgsForCall, struct {
	}{})
	fake.recordInvocation("ShortSummary", []interface{}{})
	fake.shortSummaryMutex.Unlock()
	if fake.ShortSummaryStub != nil {
		return fake.ShortSummaryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.shortSummaryReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) ShortSummaryCallCount() int {
	fake.shortSummaryMutex.RLock()
	defer fake.shortSummaryMutex.RUnlock()
	return len(fake.shortSummaryArgsForCall)
}

func (fake *FakeInfo) ShortSummaryCalls(stub func() string) {
	fake.shortSummaryMutex.Lock()
	defer fake.shortSummaryMutex.Unlock()
	fake.ShortSummaryStub = stub
}

func (fake *FakeInfo) ShortSummaryReturns(result1 string) {
	fake.shortSummaryMutex.Lock()
	defer fake.shortSummaryMutex.Unlock()
	fake.ShortSummaryStub = nil
	fake.shortSummaryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) ShortSummaryReturnsOnCall(i int, result1 string) {
	fake.shortSummaryMutex.Lock()
	defer fake.shortSummaryMutex.Unlock()
	fake.ShortSummaryStub = nil
	if fake.shortSummaryReturnsOnCall == nil {
		fake.shortSummaryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.shortSummaryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) SourceBranch() string {
	fake.sourceBranchMutex.Lock()
	ret, specificReturn := fake.sourceBranchReturnsOnCall[len(fake.sourceBranchArgsForCall)]
	fake.sourceBranchArgsForCall = append(fake.sourceBranchArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceBranch", []interface{}{})
	fake.sourceBranchMutex.Unlock()
	if fake.SourceBranchStub != nil {
		return fake.SourceBranchStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceBranchReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) SourceBranchCallCount() int {
	fake.sourceBranchMutex.RLock()
	defer fake.sourceBranchMutex.RUnlock()
	return len(fake.sourceBranchArgsForCall)
}

func (fake *FakeInfo) SourceBranchCalls(stub func() string) {
	fake.sourceBranchMutex.Lock()
	defer fake.sourceBranchMutex.Unlock()
	fake.SourceBranchStub = stub
}

func (fake *FakeInfo) SourceBranchReturns(result1 string) {
	fake.sourceBranchMutex.Lock()
	defer fake.sourceBranchMutex.Unlock()
	fake.SourceBranchStub = nil
	fake.sourceBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) SourceBranchReturnsOnCall(i int, result1 string) {
	fake.sourceBranchMutex.Lock()
	defer fake.sourceBranchMutex.Unlock()
	fake.SourceBranchStub = nil
	if fake.sourceBranchReturnsOnCall == nil {
		fake.sourceBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) SourceRepository() string {
	fake.sourceRepositoryMutex.Lock()
	ret, specificReturn := fake.sourceRepositoryReturnsOnCall[len(fake.sourceRepositoryArgsForCall)]
	fake.sourceRepositoryArgsForCall = append(fake.sourceRepositoryArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceRepository", []interface{}{})
	fake.sourceRepositoryMutex.Unlock()
	if fake.SourceRepositoryStub != nil {
		return fake.SourceRepositoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceRepositoryReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) SourceRepositoryCallCount() int {
	fake.sourceRepositoryMutex.RLock()
	defer fake.sourceRepositoryMutex.RUnlock()
	return len(fake.sourceRepositoryArgsForCall)
}

func (fake *FakeInfo) SourceRepositoryCalls(stub func() string) {
	fake.sourceRepositoryMutex.Lock()
	defer fake.sourceRepositoryMutex.Unlock()
	fake.SourceRepositoryStub = stub
}

func (fake *FakeInfo) SourceRepositoryReturns(result1 string) {
	fake.sourceRepositoryMutex.Lock()
	defer fake.sourceRepositoryMutex.Unlock()
	fake.SourceRepositoryStub = nil
	fake.sourceRepositoryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) SourceRepositoryReturnsOnCall(i int, result1 string) {
	fake.sourceRepositoryMutex.Lock()
	defer fake.sourceRepositoryMutex.Unlock()
	fake.SourceRepositoryStub = nil
	if fake.sourceRepositoryReturnsOnCall == nil {
		fake.sourceRepositoryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceRepositoryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) SourceRepositoryOwner() string {
	fake.sourceRepositoryOwnerMutex.Lock()
	ret, specificReturn := fake.sourceRepositoryOwnerReturnsOnCall[len(fake.sourceRepositoryOwnerArgsForCall)]
	fake.sourceRepositoryOwnerArgsForCall = append(fake.sourceRepositoryOwnerArgsForCall, struct {
	}{})
	fake.recordInvocation("SourceRepositoryOwner", []interface{}{})
	fake.sourceRepositoryOwnerMutex.Unlock()
	if fake.SourceRepositoryOwnerStub != nil {
		return fake.SourceRepositoryOwnerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sourceRepositoryOwnerReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) SourceRepositoryOwnerCallCount() int {
	fake.sourceRepositoryOwnerMutex.RLock()
	defer fake.sourceRepositoryOwnerMutex.RUnlock()
	return len(fake.sourceRepositoryOwnerArgsForCall)
}

func (fake *FakeInfo) SourceRepositoryOwnerCalls(stub func() string) {
	fake.sourceRepositoryOwnerMutex.Lock()
	defer fake.sourceRepositoryOwnerMutex.Unlock()
	fake.SourceRepositoryOwnerStub = stub
}

func (fake *FakeInfo) SourceRepositoryOwnerReturns(result1 string) {
	fake.sourceRepositoryOwnerMutex.Lock()
	defer fake.sourceRepositoryOwnerMutex.Unlock()
	fake.SourceRepositoryOwnerStub = nil
	fake.sourceRepositoryOwnerReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) SourceRepositoryOwnerReturnsOnCall(i int, result1 string) {
	fake.sourceRepositoryOwnerMutex.Lock()
	defer fake.sourceRepositoryOwnerMutex.Unlock()
	fake.SourceRepositoryOwnerStub = nil
	if fake.sourceRepositoryOwnerReturnsOnCall == nil {
		fake.sourceRepositoryOwnerReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.sourceRepositoryOwnerReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TargetBranch() string {
	fake.targetBranchMutex.Lock()
	ret, specificReturn := fake.targetBranchReturnsOnCall[len(fake.targetBranchArgsForCall)]
	fake.targetBranchArgsForCall = append(fake.targetBranchArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetBranch", []interface{}{})
	fake.targetBranchMutex.Unlock()
	if fake.TargetBranchStub != nil {
		return fake.TargetBranchStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetBranchReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) TargetBranchCallCount() int {
	fake.targetBranchMutex.RLock()
	defer fake.targetBranchMutex.RUnlock()
	return len(fake.targetBranchArgsForCall)
}

func (fake *FakeInfo) TargetBranchCalls(stub func() string) {
	fake.targetBranchMutex.Lock()
	defer fake.targetBranchMutex.Unlock()
	fake.TargetBranchStub = stub
}

func (fake *FakeInfo) TargetBranchReturns(result1 string) {
	fake.targetBranchMutex.Lock()
	defer fake.targetBranchMutex.Unlock()
	fake.TargetBranchStub = nil
	fake.targetBranchReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TargetBranchReturnsOnCall(i int, result1 string) {
	fake.targetBranchMutex.Lock()
	defer fake.targetBranchMutex.Unlock()
	fake.TargetBranchStub = nil
	if fake.targetBranchReturnsOnCall == nil {
		fake.targetBranchReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetBranchReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TargetRepository() string {
	fake.targetRepositoryMutex.Lock()
	ret, specificReturn := fake.targetRepositoryReturnsOnCall[len(fake.targetRepositoryArgsForCall)]
	fake.targetRepositoryArgsForCall = append(fake.targetRepositoryArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetRepository", []interface{}{})
	fake.targetRepositoryMutex.Unlock()
	if fake.TargetRepositoryStub != nil {
		return fake.TargetRepositoryStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetRepositoryReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) TargetRepositoryCallCount() int {
	fake.targetRepositoryMutex.RLock()
	defer fake.targetRepositoryMutex.RUnlock()
	return len(fake.targetRepositoryArgsForCall)
}

func (fake *FakeInfo) TargetRepositoryCalls(stub func() string) {
	fake.targetRepositoryMutex.Lock()
	defer fake.targetRepositoryMutex.Unlock()
	fake.TargetRepositoryStub = stub
}

func (fake *FakeInfo) TargetRepositoryReturns(result1 string) {
	fake.targetRepositoryMutex.Lock()
	defer fake.targetRepositoryMutex.Unlock()
	fake.TargetRepositoryStub = nil
	fake.targetRepositoryReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TargetRepositoryReturnsOnCall(i int, result1 string) {
	fake.targetRepositoryMutex.Lock()
	defer fake.targetRepositoryMutex.Unlock()
	fake.TargetRepositoryStub = nil
	if fake.targetRepositoryReturnsOnCall == nil {
		fake.targetRepositoryReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetRepositoryReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TargetRepositoryOwner() string {
	fake.targetRepositoryOwnerMutex.Lock()
	ret, specificReturn := fake.targetRepositoryOwnerReturnsOnCall[len(fake.targetRepositoryOwnerArgsForCall)]
	fake.targetRepositoryOwnerArgsForCall = append(fake.targetRepositoryOwnerArgsForCall, struct {
	}{})
	fake.recordInvocation("TargetRepositoryOwner", []interface{}{})
	fake.targetRepositoryOwnerMutex.Unlock()
	if fake.TargetRepositoryOwnerStub != nil {
		return fake.TargetRepositoryOwnerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.targetRepositoryOwnerReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) TargetRepositoryOwnerCallCount() int {
	fake.targetRepositoryOwnerMutex.RLock()
	defer fake.targetRepositoryOwnerMutex.RUnlock()
	return len(fake.targetRepositoryOwnerArgsForCall)
}

func (fake *FakeInfo) TargetRepositoryOwnerCalls(stub func() string) {
	fake.targetRepositoryOwnerMutex.Lock()
	defer fake.targetRepositoryOwnerMutex.Unlock()
	fake.TargetRepositoryOwnerStub = stub
}

func (fake *FakeInfo) TargetRepositoryOwnerReturns(result1 string) {
	fake.targetRepositoryOwnerMutex.Lock()
	defer fake.targetRepositoryOwnerMutex.Unlock()
	fake.TargetRepositoryOwnerStub = nil
	fake.targetRepositoryOwnerReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TargetRepositoryOwnerReturnsOnCall(i int, result1 string) {
	fake.targetRepositoryOwnerMutex.Lock()
	defer fake.targetRepositoryOwnerMutex.Unlock()
	fake.TargetRepositoryOwnerStub = nil
	if fake.targetRepositoryOwnerReturnsOnCall == nil {
		fake.targetRepositoryOwnerReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.targetRepositoryOwnerReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) Title() string {
	fake.titleMutex.Lock()
	ret, specificReturn := fake.titleReturnsOnCall[len(fake.titleArgsForCall)]
	fake.titleArgsForCall = append(fake.titleArgsForCall, struct {
	}{})
	fake.recordInvocation("Title", []interface{}{})
	fake.titleMutex.Unlock()
	if fake.TitleStub != nil {
		return fake.TitleStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.titleReturns
	return fakeReturns.result1
}

func (fake *FakeInfo) TitleCallCount() int {
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	return len(fake.titleArgsForCall)
}

func (fake *FakeInfo) TitleCalls(stub func() string) {
	fake.titleMutex.Lock()
	defer fake.titleMutex.Unlock()
	fake.TitleStub = stub
}

func (fake *FakeInfo) TitleReturns(result1 string) {
	fake.titleMutex.Lock()
	defer fake.titleMutex.Unlock()
	fake.TitleStub = nil
	fake.titleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) TitleReturnsOnCall(i int, result1 string) {
	fake.titleMutex.Lock()
	defer fake.titleMutex.Unlock()
	fake.TitleStub = nil
	if fake.titleReturnsOnCall == nil {
		fake.titleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.titleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeInfo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.changeReceiptMutex.RLock()
	defer fake.changeReceiptMutex.RUnlock()
	fake.changesUnderPathMutex.RLock()
	defer fake.changesUnderPathMutex.RUnlock()
	fake.fullDescriptionMutex.RLock()
	defer fake.fullDescriptionMutex.RUnlock()
	fake.owningSIGMutex.RLock()
	defer fake.owningSIGMutex.RUnlock()
	fake.pathToRepoMutex.RLock()
	defer fake.pathToRepoMutex.RUnlock()
	fake.principalEmailMutex.RLock()
	defer fake.principalEmailMutex.RUnlock()
	fake.principalNameMutex.RLock()
	defer fake.principalNameMutex.RUnlock()
	fake.shortSummaryMutex.RLock()
	defer fake.shortSummaryMutex.RUnlock()
	fake.sourceBranchMutex.RLock()
	defer fake.sourceBranchMutex.RUnlock()
	fake.sourceRepositoryMutex.RLock()
	defer fake.sourceRepositoryMutex.RUnlock()
	fake.sourceRepositoryOwnerMutex.RLock()
	defer fake.sourceRepositoryOwnerMutex.RUnlock()
	fake.targetBranchMutex.RLock()
	defer fake.targetBranchMutex.RUnlock()
	fake.targetRepositoryMutex.RLock()
	defer fake.targetRepositoryMutex.RUnlock()
	fake.targetRepositoryOwnerMutex.RLock()
	defer fake.targetRepositoryOwnerMutex.RUnlock()
	fake.titleMutex.RLock()
	defer fake.titleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInfo) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ routing.Info = new(FakeInfo)
