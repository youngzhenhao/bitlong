package com.btc.wallect.utils.threadutil;

import java.util.LinkedList;
import java.util.concurrent.atomic.AtomicInteger;

public class ThreadPool {
    // 定义一个集合来管理线程
    LinkedList<Runnable> runnables = new LinkedList<>();

    int threadCount = 3;// 定义线程池的容量最大为3
    // int count = 0;//定义当前的线程个数
    // 使用原子性的操作定义当前线程的个数，避免线程安全问题
    AtomicInteger count = new AtomicInteger(3);

    // 只能创建三个线程
    public  void execute(Runnable runnable) {
        // 每次执行的时候都把线程添加到集合中
        // count++; 这个操作会造成线程的安全问题
        runnables.add(runnable);
        // 这个方法表示得到当前数值加1的值
        if (count.incrementAndGet() < 3) {
            createThread();
        }
    }

    // 定义一个创建线程的方法
    public void createThread() {
        new Thread() {
            public void run() {
                while (true) {
                    // 先判断集合中是否有线程
                    if (runnables.size() > 0) {

                        // 即从集合中移除了一个线程，又取出了一个异步任务
                        Runnable remove = runnables.remove();
                        if (remove != null) {
                            remove.run();// 如果不为空就run
                        }

                    }else {
                        // 否则就等待 wake();
                    }
                }
            };
        }.start();
    }



}
