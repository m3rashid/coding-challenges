package huffman

import (
	"bufio"
	"fmt"
	"internal/tree"
	"math"
	"os"
	"sort"
)

type TreeNodeData struct {
	Character  rune
	Frequency  int
	IsInternal bool
}

func Encode(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	runeFrequencyMap := make(map[rune]int)
	for scanner.Scan() {
		ch := scanner.Text()[0]
		if _, ok := runeFrequencyMap[rune(ch)]; ok {
			runeFrequencyMap[rune(ch)]++
		} else {
			runeFrequencyMap[rune(ch)] = 1
		}
	}

	pow := 1
	for math.Pow(2, float64(pow)) < float64(len(runeFrequencyMap)) {
		pow++
	}

	totalCount := 0
	frequencies := []TreeNodeData{}
	for runeKey, frequency := range runeFrequencyMap {
		frequencies = append(frequencies, TreeNodeData{
			Character: runeKey,
			Frequency: frequency,
		})
		totalCount += frequency
	}

	// sort the frequencies
	sort.Slice(frequencies, func(i, j int) bool {
		return frequencies[i].Frequency < frequencies[j].Frequency
	})

	var tempFrequencies []TreeNodeData
	tempFrequencies = append(tempFrequencies, frequencies...)

	allNodes := []*tree.Node[TreeNodeData]{}

	for len(tempFrequencies) > 1 {
		sort.Slice(tempFrequencies, func(i, j int) bool {
			return tempFrequencies[i].Frequency < tempFrequencies[j].Frequency
		})

		newFreqNode := TreeNodeData{
			Frequency:  tempFrequencies[0].Frequency + tempFrequencies[1].Frequency,
			IsInternal: true,
		}

		newNode := tree.NewNode[TreeNodeData](newFreqNode, nil, nil, nil)
		leftNode := tree.NewNode[TreeNodeData](tempFrequencies[0], nil, nil, newNode)
		rightNode := tree.NewNode[TreeNodeData](tempFrequencies[1], nil, nil, newNode)
		newNode.Left = leftNode
		newNode.Right = rightNode
		tempFrequencies = tempFrequencies[2:]
		tempFrequencies = append(tempFrequencies, newFreqNode)
		allNodes = append(allNodes, newNode, leftNode, rightNode)
	}

	root := allNodes[0].GetRoot()
	root.Print()

	fmt.Println(len(allNodes), len(frequencies), len(tempFrequencies))
}
