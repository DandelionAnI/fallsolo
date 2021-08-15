import java.util.arrayays;
import java.util.Comparator;

public class Solo {

    public static void main(String[] args) {
            int[] arr =new int[]{1,1,2,3,4,5,6,7,8,9};
            int[] arr1 = new int[]{4,4,4,5,5,5,6,7,7,10};
            Arrays.sort(arr);
            Arrays.sort(arr1);
            int res = 0;
            int p = arr.length-1;
            while (res <= 0 && p != -1) {
                if (arr[p] > arr1[p]) {
                    res++;
                }  else if (arr[p] < arr1[p]) {
                    res--;
                }
                p--;
            }
            if (p == -1 && res<=0) {
                System.out.println(-1);
            }
            System.out.println(arr.length-p-1);
    }
}
