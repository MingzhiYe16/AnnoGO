package db

import (
    "github.com/gocql/gocql"
    "github.com/MingzhiYe16/ANNOGO/pkg/models"
)

// FindSNPsByChromosome fetches SNP data for a specific chromosome
func FindSNPsByChromosome(chromosome string) ([]model.SNP, error) {
    var snps []model.SNP
    query := "SELECT chromosome, position, snp_id, gene_id, annotations FROM snps WHERE chromosome = ?"
    iter := Session.Query(query, chromosome).Iter()

    var snp model.SNP
    for iter.Scan(&snp.Chromosome, &snp.Position, &snp.SNP_ID, &snp.Gene_ID, &snp.Annotations) {
        snps = append(snps, snp)
    }

    if err := iter.Close(); err != nil {
        return nil, err
    }

    return snps, nil
}
