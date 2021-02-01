package models

//StatsModel Model
type StatsModel struct {
	CountMutantDNA int     `json:"count_mutant_dna"`
	CountHumanDNA  int     `json:"count_human_dna"`
	Ratio          float64 `json:"ratio"`
}
