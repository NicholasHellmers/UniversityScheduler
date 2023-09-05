package main

// This app is a time scheduler for a given set of tasks. It
// is meant to be a restful api that can be used by a front end
// to schedule tasks. It takes in a set of tasks, and prints out the
// shortest time to complete all tasks.

type Node struct {
	Name string
	Time int
}
