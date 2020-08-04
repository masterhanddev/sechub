// SPDX-License-Identifier: MIT

package cli

import (
	"testing"

	sechubTestUtil "daimler.com/sechub/testutil"
)

// https://localhost:8443/api/project/testproject/job/
func TestBuildCreateNewSecHubJobAPICall(t *testing.T) {
	/* prepare */
	context := new(Context)
	config := new(Config)

	context.config = config
	config.projectID = "testproject"
	config.server = "https://localhost:8443"

	/* execute */
	result := buildCreateNewSecHubJobAPICall(context)

	/* test*/
	sechubTestUtil.AssertEquals("https://localhost:8443/api/project/testproject/job", result, t)

}

// https://localhost:8443/api/project/testproject/job/e21b13fc-591e-4abd-b119-755d473c5625
func TestBuildGetSecHubJobStatusAPICall(t *testing.T) {
	/* prepare */
	context := new(Context)
	config := new(Config)

	context.config = config
	config.projectID = "testproject"
	config.server = "https://localhost:8443"

	config.secHubJobUUID = "e21b13fc-591e-4abd-b119-755d473c5625"

	/* execute */
	result := buildGetSecHubJobStatusAPICall(context)

	/* test*/
	sechubTestUtil.AssertEquals("https://localhost:8443/api/project/testproject/job/e21b13fc-591e-4abd-b119-755d473c5625", result, t)

}

// https://localhost:8443/api/project/testproject/report/e21b13fc-591e-4abd-b119-755d473c5625
func TestBuildGetSecHubJobReportAPICall(t *testing.T) {
	/* prepare */
	context := new(Context)
	config := new(Config)

	context.config = config
	config.projectID = "testproject"
	config.server = "https://localhost:8443"

	config.secHubJobUUID = "e21b13fc-591e-4abd-b119-755d473c5625"

	/* execute */
	result := buildGetSecHubJobReportAPICall(context)

	/* test*/
	sechubTestUtil.AssertEquals("https://localhost:8443/api/project/testproject/report/e21b13fc-591e-4abd-b119-755d473c5625", result, t)

}

// https://localhost:8443/api/project/testproject/job/e21b13fc-591e-4abd-b119-755d473c5625
func TestBuildPostSecHubUploadSourceCodeAPICall(t *testing.T) {
	/* prepare */
	context := new(Context)
	config := new(Config)

	context.config = config
	config.projectID = "testproject"
	config.server = "https://localhost:8443"

	config.secHubJobUUID = "e21b13fc-591e-4abd-b119-755d473c5625"

	/* execute */
	result := buildUploadSourceCodeAPICall(context)

	/* test*/
	sechubTestUtil.AssertEquals("https://localhost:8443/api/project/testproject/job/e21b13fc-591e-4abd-b119-755d473c5625/sourcecode", result, t)

}
