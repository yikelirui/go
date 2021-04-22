package sort
//冒泡排序
func BubleSort(arr []int){
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr)-i-1;j++{
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
}
//选择排序
func SelectSort(arr []int){
	for  i:=0;i<len(arr)-1;i++{
		index:=i
		j:=i+1
		for ;j<len(arr);j++{
			if arr[index]<arr[j]{
				index=j//选出最小值下标
			}
		}
		arr[index],arr[i]=arr[i],arr[index]
	}
}
//插入排序
func InsertSort(arr []int){
   for i:=1;i<len(arr);i++{
   	 tmp:=arr[i]
   	 j:=i-1
   	 for ;j>=0 && arr[j]>tmp;j--{
   	 	arr[j+1]=arr[j]
	 }
   	 arr[j+1]=tmp
   }
}