// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package monitorv1iface provides an interface to enable mocking the monitorv1 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package monitorv1iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/thinry/ksc-sdk-go/service/monitorv1"
)

// Monitorv1API provides an interface to enable mocking the
// monitorv1.Monitorv1 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // monitorv1.
//    func myFunc(svc monitorv1iface.Monitorv1API) bool {
//        // Make svc.GetMetricStatisticsBatch request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := monitorv1.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockMonitorv1Client struct {
//        monitorv1iface.Monitorv1API
//    }
//    func (m *mockMonitorv1Client) GetMetricStatisticsBatch(input *map[string]interface{}) (*map[string]interface{}, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockMonitorv1Client{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type Monitorv1API interface {
	GetMetricStatisticsBatch(*map[string]interface{}) (*map[string]interface{}, error)
	GetMetricStatisticsBatchWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	GetMetricStatisticsBatchRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ Monitorv1API = (*monitorv1.Monitorv1)(nil)
