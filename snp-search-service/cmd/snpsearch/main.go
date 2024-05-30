package main

import (
    "github.com/gin-gonic/gin"
    "github.com/MingzhiYe16/AnnoGO/snp-search-service/internal/api"
)

func main() {
    router := gin.Default()

    // Route definitions
    router.GET("/snps", api.SearchSNPs) // Search SNPs endpoint

    // Start server
    router.Run() // listens on :8080 by default
}
