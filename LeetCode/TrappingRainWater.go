package main
import (
    "fmt"
);
func max(a, b int) int{
    // fmt.Println(a,b)
    if a> b{
        return a
    }
    return b
}
func min(a, b int) int{
    // fmt.Println(a,b)
    if a< b{
        return a
    }
    return b
}
func trap(height []int) int {
    heightLength := len(height)
    lmax := make([]int, heightLength);
    rmax := make([]int, heightLength);
    lmax[0] = height[0]
    rmax[heightLength-1] = height[heightLength-1]
    for i := 1; i < heightLength; i++{
        lmax[i] = max(height[i], lmax[i-1])
    }
    for i := heightLength-2; i >=0; i--{
        rmax[i] = max(height[i], rmax[i+1])
    }
    // fmt.Println(lmax)
    var a int;
    for i := 0; i<heightLength; i++{
        a += (min(lmax[i], rmax[i])-height[i])
    }
    return a

}
func main(){
	a := []int{0,1,0,2,1,0,1,3,2,1,2,1} 
	fmt.Println(trap(a))
}

