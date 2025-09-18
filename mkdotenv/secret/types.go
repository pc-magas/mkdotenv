package secret

const (
    ValueTypeNone  = ""
    ValueTypeKeepassX = "keepassx"
    // ValueTypeAWSSecrets      = "aws-secrets-manager"
    // ValueTypeAzureKeyVault   = "azure-key-vault"
    // ValueTypeGCPSecretManager = "gcp-secret-manager"
)

var ValueTypes = []string{
    ValueTypeNone,
    ValueTypeKeepassX,
}

// VerifyType checks if the given value is one of the supported value types.
func VerifyType(valueType string) bool {
    for _, t := range ValueTypes {
        if valueType == t {
            return true
        }
    }
    return false
}
