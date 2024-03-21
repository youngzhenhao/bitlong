package com.btc.wallect.utils.threadutil;



import com.btc.wallect.utils.LogUntil;

import java.util.concurrent.ArrayBlockingQueue;
import java.util.concurrent.Callable;
import java.util.concurrent.Executors;
import java.util.concurrent.Future;
import java.util.concurrent.ThreadPoolExecutor;
import java.util.concurrent.TimeUnit;

public class ThreadPoolUtils {
    public static volatile ThreadPoolExecutor pool = null;


    private final static int CORE_POOL_SIZE = 10;//核心线程数
    private final static int MAXIMUM_POOL_SIZE = 50;//最大线程数
    private final static long KEEP_ALIVE_TIME = 5;//非核心线程空闲时间
    private final static int MAXIMUM_WORK_QUEUE = 400;//消息队列最大任务数

    private ThreadPoolUtils() {
        pool = new ThreadPoolExecutor(CORE_POOL_SIZE,// 1. 核心线程数
                MAXIMUM_POOL_SIZE,// 2. 最大线程数
                KEEP_ALIVE_TIME, // 3. 非核心线程空闲时间
                TimeUnit.SECONDS, // 4. 时间单位
                new ArrayBlockingQueue<>(MAXIMUM_WORK_QUEUE),// 5. 阻塞队列
                Executors.defaultThreadFactory(),// 6. 创建线程工厂
                new ThreadPoolExecutor.AbortPolicy());// 7. 拒绝策略
//        LogUntil.e( "PushThreadPoolUtils 线程池创建完成");
    }

    public static ThreadPoolExecutor getInstance() {
        //双重if单例
        if (null == pool) {
            synchronized (ThreadPoolUtils.class) {
                if (pool == null) {
                    new ThreadPoolUtils();
                    return pool;
                }
            }
        }
        return pool;
    }

    // 无响应执行
    public static void execute(Runnable runnable) {
        getInstance().execute(runnable);
    }

    // 有响应执行
    public static <T> Future<T> submit(Callable<T> callable) {
        return getInstance().submit(callable);
    }
}
