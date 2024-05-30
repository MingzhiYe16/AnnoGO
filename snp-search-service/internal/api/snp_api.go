package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/MingzhiYe16/AnnoGO/snp-search-service/internal/db"
    "github.com/MingzhiYe16/ANNOGO/pkg/models"
)

// SearchSNPs handles the search requests for SNP data
func SearchSNPs(c *gin.Context) {
    // You might want to accept parameters like chromosome, position, etc.
    chromosome := c.Query("chromosome")
    position := c.Query("position") // You need to convert this to an int

    // For now, let's assume we're just using the chromosome to filter results
    snps, err := db.FindSNPsByChromosome(chromosome)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, snps)
}
