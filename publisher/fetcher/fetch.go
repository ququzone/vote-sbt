package fetcher

import (
	"context"

	graphql "github.com/hasura/go-graphql-client"
)

type OrderDirection string

func Fetch(id string, skip uint64) ([]string, error) {
	client := graphql.NewClient("https://hub.snapshot.org/graphql", nil)

	var q struct {
		Votes []struct {
			Ipfs    string  `graphql:"ipfs"`
			Voter   string  `graphql:"voter"`
			Choice  uint64  `graphql:"choice"`
			Vp      float64 `graphql:"vp"`
			Reason  string
			Created uint64
		} `graphql:"votes(first: $first, skip: $skip, where: {proposal: $id, vp_gt: 0, space: $space}, orderBy: $orderBy, orderDirection: $orderDirection)"`
	}
	var orderDirection OrderDirection = "desc"
	variables := map[string]interface{}{
		"id":             id,
		"first":          10,
		"skip":           skip,
		"orderBy":        "vp",
		"orderDirection": orderDirection,
		"space":          "iotex.eth",
	}

	err := client.Query(context.Background(), &q, variables)
	if err != nil {
		return nil, err
	}
	result := make([]string, len(q.Votes))
	for i, v := range q.Votes {
		result[i] = v.Voter
	}
	return result, nil
}
