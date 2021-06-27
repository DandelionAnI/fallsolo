import java.util.Arrays;
import java.util.Comparator;

public class Solo {

    public int maxEnvelopes(int[][] envelopes) {
        int n = envelopes.length;
        if (envelopes.length == 0)
            return 0;

        Arrays.sort(envelopes,new Comparator<int[]>(){
            public int compare(int[] e1, int[] e2) {
                if (e1[0] != e2[0]) 
                    return e1[0] - e2[0];
                else return e2[1] - e1[1];
            }
        });
            
    }

}
