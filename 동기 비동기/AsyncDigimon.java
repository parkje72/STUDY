import java.util.ArrayList;

public class AsyncDigimon{
    public static void main(String[] args){
        for(int i=0;i<5;i++){
            Thread thread = new Thread(new PrintTask(i));
            thread.start();
        }
    }

    private static class PrintTask implements Runnable {
        final int taskNumber;
        ArrayList<String> digimon = new ArrayList<String>();
        
       

        public PrintTask(int taskNumber){
            this.taskNumber = taskNumber;
            digimon.add("안");
            digimon.add("녕");
            digimon.add("디");
            digimon.add("지");
            digimon.add("몬");
        }



        @Override
        public void run() {
            // TODO Auto-generated method stub
            throw new UnsupportedOperationException("Unimplemented method 'run'");
        }
    }

    // @Override
    // public void run() {
    //     Object taskNumber;
    //     System.out.println(digimon.get(taskNumber));
    // }
}
// 몬
// 안
// 지
// 디
// 녕
