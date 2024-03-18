package com.btc.wallect.utils;

import android.util.Log;

import com.btc.wallect.BuildConfig;

public class LogUntil {
    private static String className;        //类名
    private static String methodName;       //方法名
    private static int lineNumber;          //行数
    private static final boolean openLog = BuildConfig.DEBUG;

    /**
     * methodName 方法名
     * className  类名
     * lineNumber 行数
     */
    private static void getMethodNames(StackTraceElement[] sElements) {
        className = sElements[1].getFileName();
        methodName = sElements[1].getMethodName();
        lineNumber = sElements[1].getLineNumber();
    }

    /**
     * `
     * 封装 log 信息
     *
     * @param log 打印信息
     */
    private static String createLog(String log) {
        return methodName +
                "(" + className + ":" + lineNumber + ")" +
                log;
    }

    public static String i(String s) {
        getMethodNames(new Throwable().getStackTrace());
        String ss = createLog(s);
        Log.i("TAG", ss);
        return ss;
    }

    public static void i(String s, Throwable e) {
        if (openLog) {
            Log.i("TAG", s, e);
        }
    }


    public static String w(String s) {
        getMethodNames(new Throwable().getStackTrace());
        String ss = createLog(s);
        Log.w("TAG.W", ss);
        return ss;

    }

    public static void w(String s, Throwable e) {
        if (openLog) {
            Log.w("TAG.W", s, e);
        }
    }


    public static String e(String s) {
        getMethodNames(new Throwable().getStackTrace());
        String ss = createLog(s);
        Log.e("TAG.E", ss);
        return ss;
    }

    public static void e(String s, Throwable e) {
        if (openLog) {
            Log.e("TAG.E", s, e);
        }
    }

    public static String d(String s) {
        getMethodNames(new Throwable().getStackTrace());
        String ss = createLog(s);
        Log.d("TAG.D", ss);
        return ss;
    }

    public static void d(String s, Throwable e) {
        if (openLog) {
            Log.d("TAG.D", s, e);
        }
    }

    public static String v(String s) {
        getMethodNames(new Throwable().getStackTrace());
        String ss = createLog(s);
        Log.v("TAG.V", ss);
        return ss;
    }

    public static void v(String s, Throwable e) {
        if (openLog) {
            Log.v("TAG.V", s, e);
        }
    }

    public static String wtf(String s) {
        getMethodNames(new Throwable().getStackTrace());
        String ss = createLog(s);
        Log.wtf("TAG.WTF", ss);
        return ss;
    }

    public static void wtf(String s, Throwable e) {
        if (openLog) {
            Log.wtf("TAG.WTF", s, e);
        }
    }

}
