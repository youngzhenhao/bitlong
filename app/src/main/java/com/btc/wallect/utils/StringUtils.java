package com.btc.wallect.utils;

import java.util.List;

public class StringUtils {
    /**
     * 判断str是否为空
     *
     * @param str String
     * @return boolean true:空;false:非空
     */
    public static boolean isStrEmpty(String str) {
        return ((str == null) || (str.trim().equals("")));
    }
    /**
     * 判断list是否为空
     * @param str String
     * @return boolean true:空;false:非空
     */
    public static boolean isListEmpty(List str) {
        return ((str == null) || (str.size()<=0));
    }

    //获取String
    public static String getString(String str) {
        if (str == null || "null".equalsIgnoreCase(str)) return "";
        return str;
    }
    //获取String
    public static String getString(Object str) {
        if (str == null || "null".equalsIgnoreCase(str.toString()))  return "";
        return str.toString();
    }

}
