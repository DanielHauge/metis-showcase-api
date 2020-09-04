package main

type NodeData struct{
	Type int32 `json:"type"`
	Url string `json:"url"`
	Title string `json:"title"`
}

type Node struct{
	Data NodeData `json:"data"`
	Children map[string]Node `json:"children"`
}
