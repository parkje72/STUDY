class Solution {
    public static int solution(String s) {
        String answer = "";
        String[] arr = s.split("");
        String sign = "";
        
        if(arr.length >= 1 && arr.length <= 5){
            if(arr[0].equals("-") || arr[0].equals("+")){
                sign = arr[0];
                String[] arrCopy = new String[arr.length - 1]; 
                for(int j = 0; j < arrCopy.length; j++){
                    arrCopy[j] = arr[j+1];
                }
                answer = String.join("",arrCopy);
            } 
            else {
                for(int i = 0; i < arr.length; i++){
                    answer = arr[i];
                }
                answer = String.join("",arr);
            }
        }
        return Integer.parseInt(sign + answer);
    }
}
