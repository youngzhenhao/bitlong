package com.btc.wallect.utils;

import android.content.Context;
import android.util.Log;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileReader;
import java.io.RandomAccessFile;

/**
 * 工具类
 *
 * @author gph
 */
public class Tool {

    /**
     * 将字符串写入到文本文件中
     */
    public void writeTxtToFile(Context context,String strcontent, String filePath,
                               String fileName) {
        // 生成文件夹之后，再生成文件，不然会出错
        makeFilePath(filePath, fileName);// 生成文件

        String strFilePath = filePath + fileName;
SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.WALLECT_ACCOUT,strFilePath);

        Log.e("LZ>>>strFilePath:",strFilePath);
        // 每次写入时，都换行写
        String strContent = strcontent + "\r\n";
        try {
            File file = new File(strFilePath);
            if (!file.exists()) {
                Log.d("TestFile", "Create the file:" + strFilePath);
                file.getParentFile().mkdirs();
                file.createNewFile();
            }
            RandomAccessFile raf = new RandomAccessFile(file, "rwd");
            raf.seek(file.length());
            raf.write(strContent.getBytes());
            raf.close();
        } catch (Exception e) {
            Log.e("error:", e + "");
        }
    }

    /**
     * 生成文件
     */
    public File makeFilePath(String filePath, String fileName) {
        File file = null;
        makeRootDirectory(filePath);// 生成文件夹
        try {
            file = new File(filePath + fileName);
            if (!file.exists()) {
                file.createNewFile();
            }
        } catch (Exception e) {
            e.printStackTrace();
        }
        return file;
    }

    /**
     * 生成文件夹
     */
    public static void makeRootDirectory(String filePath) {
        File file = null;
        try {
            file = new File(filePath);
            if (!file.exists()) {
                file.mkdir();
            }
        } catch (Exception e) {
            Log.i("error:", e + "");
        }
    }

    public String readTxt(String file) {
        Log.e("LZ:", file);
        BufferedReader bre = null;
        String str = "";
        String returnstr = "";
        String a;
        try {

            bre = new BufferedReader(new FileReader(file));//此时获取到的bre就是整个文件的缓存流
            while ((str = bre.readLine()) != null) { // 判断最后一行不存在，为空结束循环

                Log.e("LZ", "readTxt: a------------" + str);

                String[] arr = str.split("\\s+");
                for (String ss : arr) {
                    a = arr[0];
                }

                Log.e("LZ-----str:", str);
                returnstr=str;
            }

        } catch (Exception e) {
            Log.e("LZ", "readTxt: ---------------" + e.toString());
        }
        return returnstr;
    }


}
