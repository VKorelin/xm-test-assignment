package authorization

const adminRole string = "admin"

const companyServicePath = "/xm.api.company.v1.CompanyService/"

// Not good because RPC methods names can be changed which will lead to to fix this code also
var accessabilityRoles = map[string]string{
	companyServicePath + "Create": adminRole,
	companyServicePath + "Delete": adminRole,
	companyServicePath + "Patch":  adminRole,
}
