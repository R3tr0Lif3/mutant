package db

import (
	"log"

	"mutant/models"

	"strconv"

	"github.com/gomodule/redigo/redis"
)

const server = "localhost:6379"

//SaveDNA set dna secuence for stats
func SaveDNA(dna []string, isMutant bool) error {

	conn, err := redis.Dial("tcp", server)

	if err != nil {

		log.Fatal(err)

	}

	defer conn.Close()

	ok, err := conn.Do("HMSET", dna, "verified", "ok", "isMutant", isMutant)

	if err != nil {

		log.Fatal(ok)

		return err

	}

	if isMutant {

		ok, err := conn.Do("INCR", "mutant")

		if err != nil {

			log.Fatal(ok)

			return err

		}

	} else {

		ok, err := conn.Do("INCR", "human")

		if err != nil {

			log.Fatal(ok)

			return err

		}

	}

	return nil

}

//DNAExist check if dna Exist
func DNAExist(dna []string) (models.VerifiedModel, error) {

	var verified = models.VerifiedModel{}

	conn, err := redis.Dial("tcp", server)

	if err != nil {

		log.Fatal(err)

	}

	defer conn.Close()

	found, err := redis.Strings(conn.Do("HMGET", dna, "verified", "isMutant"))

	if err != nil {

		return verified, err

	}

	if found[0] == "ok" {

		verified.Verified = true

		verified.IsMutannt, _ = strconv.ParseBool(found[1])

		return verified, nil

	}

	verified.Verified = false

	return verified, nil

}

//GetStatsBD get stats from redis
func GetStatsBD() (models.StatsModel, error) {

	var stats models.StatsModel

	conn, err := redis.Dial("tcp", server)

	if err != nil {

		return stats, err

	}

	defer conn.Close()

	human, err := redis.Int((conn.Do("GET", "human")))

	mutant, err := redis.Int((conn.Do("GET", "mutant")))

	//Get result for stats request

	mutantFloat := float64(mutant)

	humanFloat := float64(human)

	if mutantFloat == 0 || humanFloat == 0 {

		stats.CountMutantDNA = mutant

		stats.CountHumanDNA = human

		stats.Ratio = 0.0

		return stats, nil

	}

	var ratio = mutantFloat / humanFloat

	stats.CountMutantDNA = mutant

	stats.CountHumanDNA = human

	stats.Ratio = ratio

	return stats, nil

}
