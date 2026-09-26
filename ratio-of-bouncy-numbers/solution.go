package main

import (
  "errors"
  "strconv"
)

func bouncyRatio(ratio float64) (int, error) {
  if ratio < 0 || ratio > 99 {
    return 0, errors.New("invalid ratio")
  }
  
  bouncyCount := 0
  
  for n := 1; ; n++ {
    if isBouncy(n) {
      bouncyCount++
    }
    
    currentRatio := float64(bouncyCount) / float64(n) * 100
    if currentRatio >= ratio {
      return n, nil
    }
  }
}

func isBouncy(n int) bool {
  return !isIncreasing(n) && !isDecreasing(n)
}

func isIncreasing(n int) bool {
  s := strconv.Itoa(n)
  for i := 1; i < len(s); i++ {
    if s[i] < s[i-1] {
      return false
    }
  }
  return true
}

func isDecreasing(n int) bool {
  s := strconv.Itoa(n)
  for i := 1; i < len(s); i++ {
    if s[i] > s[i-1] {
      return false
    }
  }
  return true
}