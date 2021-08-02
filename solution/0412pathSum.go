package solution

import "fmt"

func pathSum(root *TreeNode, sum int) int {
    ans,road := 0,[][]int{}
    var dfs func(node *TreeNode, target int, add []int) 
    dfs = func(node *TreeNode, target int, add []int) {
        add = append(add, node.Val)
        target = target - node.Val
        if target == 0 {
            temp := []int{}
            copy(temp, add)
            road = append(road, temp)
            ans++
        }
        if node.Left != nil {
            dfs(node.Left, target,add)
            add = add[:len(add)-1]
        } 
        if node.Right != nil {
            dfs(node.Right, target,add)
            add = add[:len(add)-1]
        } 
    }
    var bianli func(node *TreeNode)
    bianli = func(node *TreeNode) {
        add := []int{}
        dfs(node,sum,add)
        if node.Left != nil {
            bianli(node.Left)
        } 
        if node.Right != nil {
            bianli(node.Right)
        }         
    }
    bianli(root)
    fmt.Println(road)
    return ans
}

