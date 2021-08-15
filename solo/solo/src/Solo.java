import java.util.arrayays;
import java.util.Comparator;

public class Solo {

    public int maxEnvelopes(int[][] envelopes) {
//         int n = a.length;
//         int sum = 0;
//         for (int i=0; i<n ;i++) {
//             sum += a[i];
//         }
//         int target = sum/2;
    
//     int ans = sum;
//     int l=0,r=0;
//     int cur =0;
//     while (r<n) {
//         cur += a[r];
//         if (cur >= target) {
//             ans = Math.min(ans,Math.abs(cur-(sum-cur)));
//             while (l<r && cur > target) {
//                 cur -= a[l];
//                 l++;
//                 ans = Math.min(ans, Math.abs(cur - (sum-cur)));
//             }
//             if (ans == target) {
//                 return ans;
//             }
//         }
//         r++;
//     }
//     return ans;
// }
}

    public int maxEnvelopes(int[] array, int k) {
        Arrays.sort(arr);
        int res = 0;
        for (int i=0; i< array.length;i++) {
            int r = k-array[i];
            if (r< array[i]) {
                break;
            }
            int pos = Arrays.binarySearch(array,r);

            if(pos >= 0) {
                while(pos+1<array.length && array[pos+1] == array[pos]) {
                    pos++;
                }
                res += Math.max(0, pos-i);
            } else {
                int rr = Math.min(array.length-1,-pos-1);
                res += Math.max(0,rr-i);
            }
        }
        return res;
    }
