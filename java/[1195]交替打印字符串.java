package pri.leetcode.leetcode.editor.cn;
//编写一个可以从 1 到 n 输出代表这个数字的字符串的程序，但是：
//
// 
// 如果这个数字可以被 3 整除，输出 "fizz"。 
// 如果这个数字可以被 5 整除，输出 "buzz"。 
// 如果这个数字可以同时被 3 和 5 整除，输出 "fizzbuzz"。 
// 
//
// 例如，当 n = 15，输出： 1, 2, fizz, 4, buzz, fizz, 7, 8, fizz, buzz, 11, fizz, 13, 14, fizzbuzz。 
//
// 假设有这么一个类： 
//
// class FizzBuzz {
//  public FizzBuzz(int n) { ... }               // constructor
//  public void fizz(printFizz) { ... }          // only output "fizz"
//  public void buzz(printBuzz) { ... }          // only output "buzz"
//  public void fizzbuzz(printFizzBuzz) { ... }  // only output "fizzbuzz"
//  public void number(printNumber) { ... }      // only output the numbers
//} 
//
// 请你实现一个有四个线程的多线程版 FizzBuzz， 同一个 FizzBuzz 实例会被如下四个线程使用： 
//
// 
// 线程A将调用 fizz() 来判断是否能被 3 整除，如果可以，则输出 fizz。 
// 线程B将调用 buzz() 来判断是否能被 5 整除，如果可以，则输出 buzz。 
// 线程C将调用 fizzbuzz() 来判断是否同时能被 3 和 5 整除，如果可以，则输出 fizzbuzz。 
// 线程D将调用 number() 来实现输出既不能被 3 整除也不能被 5 整除的数字。 
// 
//

import java.util.function.IntConsumer;

//leetcode submit region begin(Prohibit modification and deletion)
class FizzBuzz {
    private int n;
    private volatile int i;
    public FizzBuzz(int n) {
        this.n = n;
        this.i = 1;
    }

    // printFizz.run() outputs "fizz".
    public void fizz(Runnable printFizz) throws InterruptedException {
        while (true) {
            synchronized (this) {
                if (this.i > n) {
                    break;
                }

                if (this.i % 3 == 0 && this.i % 5 != 0) {
                    printFizz.run();
                    this.i++;
                }
            }
        }
    }

    // printBuzz.run() outputs "buzz".
    public void buzz(Runnable printBuzz) throws InterruptedException {
        while (true) {
            synchronized (this) {
                if (this.i > n) {
                    break;
                }

                if (this.i % 5 == 0 && this.i % 3 != 0) {
                    printBuzz.run();
                    this.i++;
                }
            }
        }
    }

    // printFizzBuzz.run() outputs "fizzbuzz".
    public void fizzbuzz(Runnable printFizzBuzz) throws InterruptedException {

        while (true) {
            synchronized (this) {
                if (this.i > n) {
                    break;
                }

                if (this.i % 5 == 0 && this.i % 3 == 0) {
                    printFizzBuzz.run();
                    this.i++;
                }
            }
        }
    }

    // printNumber.accept(x) outputs "x", where x is an integer.
    public void number(IntConsumer printNumber) throws InterruptedException {

        while (true) {
            synchronized (this) {
                if (this.i > n) {
                    break;
                }

                if (this.i % 5 != 0 && this.i % 3 != 0) {
                    printNumber.accept(this.i);
                    this.i++;
                }
            }
        }
    }

    public static void main(String[] args) {
        Runnable fizz = () -> System.out.println("fizz");
        Runnable buzz = () -> System.out.println("buzz");
        Runnable fizzbuzz = () -> System.out.println("fizzbuzz");
        IntConsumer intConsumer = System.out::println;

        FizzBuzz fizzBuzz = new FizzBuzz(15);

        Thread a = new Thread(()-> {
            try {
                fizzBuzz.fizz(fizz);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });
        Thread b = new Thread(()-> {
            try {
                fizzBuzz.buzz(buzz);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });
        Thread c = new Thread(()-> {
            try {
                fizzBuzz.fizzbuzz(fizzbuzz);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });
        Thread d = new Thread(()-> {
            try {
                fizzBuzz.number(intConsumer);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        });

        a.start();
        b.start();
        c.start();
        d.start();
    }
}
//leetcode submit region end(Prohibit modification and deletion)
