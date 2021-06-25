package waiter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/batch"
	"github.com/cloudposse/terraform-provider-awsutils/internal/service/batch/finder"
	"github.com/cloudposse/terraform-provider-awsutils/internal/tfresource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func ComputeEnvironmentStatus(conn *batch.Batch, name string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		computeEnvironmentDetail, err := finder.ComputeEnvironmentDetailByName(conn, name)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return computeEnvironmentDetail, aws.StringValue(computeEnvironmentDetail.Status), nil
	}
}
