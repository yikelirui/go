package sort
//冒泡排序
func BubleSort(arr []int){
	flag:=false
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr)-1-i;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
				flag=true
			}

		}
		if !flag{
			return
		}
	}
}
//选择排序
func SelectSort(arr []int){
   for i:=0;i<len(arr);i++{
   	  index:=i
   	  j:=i+1
   	  for ;j<len(arr);j++{
   	  	if arr[j]>arr[index]{
   	  		index=j
		}
	  }
   	  arr[i],arr[index]=arr[index],arr[i]
   }
}
//插入排序
func InsertSort(arr []int){
   for i:=1;i<len(arr);i++{
   	 j:=i-1
   	 tmp:=arr[i]
   	 for ;j>=0 && arr[j]>tmp;j--{
   	 	arr[j+1]=arr[j]
	 }
   	 arr[j+1]=tmp
   }
}