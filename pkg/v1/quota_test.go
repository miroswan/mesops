package v1

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/miroswan/mesops/pkg/v1/master"
	"github.com/miroswan/mesops/pkg/v1/mesos"
	"github.com/miroswan/mesops/pkg/v1/quota"
)

func TestGetQuota(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	responseType := master.Response_GET_QUOTA
	role := "test-role"
	principal := "test-principal"
	resourceName := "test-mem"
	resourceValue := 1.0
	valueType := mesos.Value_Type(1.0)
	response := master.Response{
		Type: &responseType,
		GetQuota: &master.Response_GetQuota{
			Status: &quota.QuotaStatus{
				Infos: []*quota.QuotaInfo{
					&quota.QuotaInfo{
						Role:      &role,
						Principal: &principal,
						Guarantee: []*mesos.Resource{
							&mesos.Resource{
								Name: &resourceName,
								Type: &valueType,
								Scalar: &mesos.Value_Scalar{
									Value: &resourceValue,
								},
							},
						},
					},
				},
			},
		},
	}

	output, err := proto.Marshal(&response)
	if err != nil {
		t.Fatal(err)
	}

	s.SetOutput(output).Handle()

	data, err := s.Master().GetQuota(s.Ctx())
	if err != nil {
		t.Fatal(err)
	}

	respRole := data.GetGetQuota().GetStatus().GetInfos()[0].GetRole()
	respPrincipal := data.GetGetQuota().GetStatus().GetInfos()[0].GetPrincipal()
	respResourceName := data.GetGetQuota().GetStatus().GetInfos()[0].GetGuarantee()[0].GetName()
	respResourceValue := data.GetGetQuota().GetStatus().GetInfos()[0].GetGuarantee()[0].GetScalar().GetValue()

	if role != respRole {
		t.Errorf("expected %s, got %s", role, respRole)
	}
	if principal != respPrincipal {
		t.Errorf("expected %s, got %s", principal, respPrincipal)
	}
	if resourceName != respResourceName {
		t.Errorf("expected %s, got %s", resourceName, respResourceName)
	}
	if resourceValue != respResourceValue {
		t.Errorf("expected %s, got %s", resourceValue, respResourceValue)
	}
}

func TestSetQuota(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	force := false
	role := "test-role"
	resourceName := "test-mem"
	valueType := mesos.Value_Type(1.0)
	resourceValue := 1.0
	call := &master.Call_SetQuota{
		QuotaRequest: &quota.QuotaRequest{
			Force: &force,
			Role:  &role,
			Guarantee: []*mesos.Resource{
				&mesos.Resource{
					Name: &resourceName,
					Type: &valueType,
					Scalar: &mesos.Value_Scalar{
						Value: &resourceValue,
					},
				},
			},
		},
	}
	err := s.Master().SetQuota(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}

func TestRemoveQuota(t *testing.T) {
	s := NewTestProtobufServer(MasterClient)
	defer s.Teardown()

	s.Handle()

	role := "test-role"
	call := &master.Call_RemoveQuota{Role: &role}

	err := s.Master().RemoveQuota(s.Ctx(), call)
	if err != nil {
		t.Error(err)
	}
}
