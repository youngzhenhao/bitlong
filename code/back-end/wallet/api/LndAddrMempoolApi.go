package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// TODO: Click address to query transaction through the MemPool API.

// https://mempool.space/testnet/api/v1/difficulty-adjustment

type DifficultyResponse struct {
	ProgressPercent       float64 `json:"progressPercent"`
	DifficultyChange      float64 `json:"difficultyChange"`
	EstimatedRetargetDate int64   `json:"estimatedRetargetDate"`
	RemainingBlocks       int     `json:"remainingBlocks"`
	RemainingTime         int     `json:"remainingTime"`
	PreviousRetarget      float64 `json:"previousRetarget"`
	PreviousTime          int     `json:"previousTime"`
	NextRetargetHeight    int     `json:"nextRetargetHeight"`
	TimeAvg               int     `json:"timeAvg"`
	AdjustedTimeAvg       int     `json:"adjustedTimeAvg"`
	TimeOffset            int     `json:"timeOffset"`
	ExpectedBlocks        float64 `json:"expectedBlocks"`
}

func GetDifficultyAdjustment() string {
	targetUrl := "https://mempool.space/testnet/api/v1/difficulty-adjustment"
	response, err := http.Get(targetUrl)
	if err != nil {
		fmt.Printf("%s http.PostForm :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "http get fail.", "")
	}
	bodyBytes, _ := io.ReadAll(response.Body)
	var difficultyResponse DifficultyResponse
	if err := json.Unmarshal(bodyBytes, &difficultyResponse); err != nil {
		fmt.Printf("%s PSTUUI json.Unmarshal :%v\n", GetTimeNow(), err)
		return MakeJsonResult(false, "Unmarshal response body fail.", "")
	}
	return MakeJsonResult(true, "", difficultyResponse)
}
