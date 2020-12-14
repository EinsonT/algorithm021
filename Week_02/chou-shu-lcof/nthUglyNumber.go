package chou_shu_lcof
func nthUglyNumber(n int) int {
	dp:=make([]int,n)
	dp[0]=1
	a,b,c :=0,0,0
	for i:=1;i<n;i++{
		t1,t2,t3:=dp[a]*2,dp[b]*3,dp[c]*5
		dp[i]=min(min(t1,t2),t3)
		if dp[i]==t1{
			a++
		}
		if dp[i]==t2{
			b++
		}
		if dp[i]==t3{
			c++
		}
	}
	return dp[n-1]
}

func min(a,b int)int{
	if a<b{
		return a
	}
	return b
}