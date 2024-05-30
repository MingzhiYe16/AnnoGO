package models

// SNP represents the SNP information
type SNP struct {
    Chromosome  string
    Position    int
    SNP_ID      string
    Gene_ID     string
    Annotations map[string]map[string]interface{} // Nested map for categories
}
