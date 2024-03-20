package com.btc.wallect.utils;

import android.app.Activity;
import android.content.Context;
import android.os.Environment;
import android.text.TextUtils;
import android.util.Log;
import android.widget.Toast;

import java.io.BufferedReader;
import java.io.BufferedWriter;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.io.RandomAccessFile;
import java.io.UnsupportedEncodingException;
import java.util.ArrayList;
import java.util.List;

public class FileUtils {
    private static final String TAG = "FileUtils";

    /**
     * 创建文件
     *
     * @param filePath 文件地址
     * @param fileName 文件名
     * @return
     */
    public static void CreeateTxt(Activity context, String pathName, String fileName, String content, String path) {
        //动态创建文件
        String filePath;
        String fileName1;
        File appDir = context.getFilesDir();
        if (TextUtils.isEmpty(pathName)) {
            Toast.makeText(context, "文件夹名不能为空", Toast.LENGTH_SHORT).show();
        } else if (TextUtils.isEmpty(fileName)) {
            Toast.makeText(context, "文件名不能为空", Toast.LENGTH_SHORT).show();
        } else {
            Tool tool = new Tool();
            filePath = path + "/" + pathName + "/";
            fileName1 = fileName + ".txt";

            tool.writeTxtToFile(context, content, filePath, fileName1);// 将字符串写入到文本文件中
            Toast.makeText(context, "创建成功", Toast.LENGTH_SHORT).show();
        }
    }

    public static void CreeateFileTxt(Activity context, String fileName, String content) {
        //动态创建文件
        String filePath;
        String fileName1;
        File appDir = context.getFilesDir();
        if (TextUtils.isEmpty(fileName)) {
            Toast.makeText(context, "文件名不能为空", Toast.LENGTH_SHORT).show();
        } else {
            Tool tool = new Tool();

            filePath = appDir + "/" + "NewFolderBit" + "/";
            fileName1 = fileName + ".txt";
            tool.writeTxtToFile(context, content, filePath, fileName1);// 将字符串写入到文本文件中
            Toast.makeText(context, "创建成功", Toast.LENGTH_SHORT).show();
        }
    }

    public static void CreeateFileTxtTest(Activity context, String fileName, String content) {
        //动态创建文件
        String filePath;
        String fileName1;
        File appDir = context.getFilesDir();

        String name = appDir + "/.lnd";

        if (TextUtils.isEmpty(fileName)) {
            Toast.makeText(context, "文件名不能为空", Toast.LENGTH_SHORT).show();
        } else {
            Tool tool = new Tool();

            filePath = name + "/";
            fileName1 = fileName + ".txt";
            tool.writeTxtToFile(context, content, filePath, fileName1);// 将字符串写入到文本文件中
            Toast.makeText(context, "创建成功", Toast.LENGTH_SHORT).show();
        }
    }

    public static String CreeateFile(Activity context) {
        File appDir = context.getFilesDir();
        String name = appDir + "/.lnd";
        File tempFile = new File(name);
        if (tempFile.exists()) {
            Toast.makeText(context, "lnd已有文件夹", Toast.LENGTH_SHORT).show();
        } else {
            tempFile.mkdir();
            Toast.makeText(context, "lnd创建成功", Toast.LENGTH_SHORT).show();

        }
        Log.e("LZ>>>lnd路径：", tempFile.getPath());
        return tempFile.getPath();

    }

    public static String CreeateFileTest(Activity context) {
        File appDir = Environment.getExternalStorageDirectory();

        String name = appDir + "/.lnd";
        File tempFile = new File(name);
        if (tempFile.exists()) {
            Toast.makeText(context, "lnd已有文件夹", Toast.LENGTH_SHORT).show();
        } else {
            tempFile.mkdir();
            Toast.makeText(context, "lnd创建成功", Toast.LENGTH_SHORT).show();

        }
        Log.e("LZ>>>lnd路径test：", tempFile.getPath());
        return tempFile.getPath();

    }

    public static String CreeateZipFile(Activity context) {
        File appDir = Environment.getExternalStorageDirectory();
        String name = appDir + "/.lnd";
        File tempFile = new File(name);
        if (tempFile.exists()) {
            tempFile.delete();
            //  Toast.makeText(context, "testUnZip已有文件夹", Toast.LENGTH_SHORT).show();
        } else {
            tempFile.mkdir();
            Toast.makeText(context, "testind创建成功", Toast.LENGTH_SHORT).show();

        }
        Log.e("LZ>>>testUnZip路径：", tempFile.getPath());
        return tempFile.getPath();

    }

    public static String filePath(Context context) {
        String folderName = "NewFolderBit"; // 要创建的文件夹名称
        //   String appDir =Environment.getExternalStorageDirectory();
        //  File appDir = Environment.getExternalStorageDirectory();
        // File appDir = Environment.getDataDirectory();
        // 获取当前APP目录
        File appDir = context.getFilesDir();
        File newFolder = new File(appDir, folderName);

        if (!newFolder.exists()) {
            if (newFolder.mkdirs()) {
                Log.d("TAG", "成功创建文件夹");
                Toast.makeText(context, "成功创建文件夹！" + newFolder.getPath(), Toast.LENGTH_SHORT).show();
            } else {
                Log.e("TAG", "无法创建文件夹");
                Toast.makeText(context, "无法创建文件夹！", Toast.LENGTH_SHORT).show();
            }
        } else {
            Log.w("TAG", "文件夹已存在");
            Toast.makeText(context, "文件夹已存在！", Toast.LENGTH_SHORT).show();
        }


        return newFolder.getPath();
    }

    public static boolean createFile(String filePath, String fileName) {
        String strFilePath = filePath + fileName;
        File file = new File(filePath);
        if (!file.exists()) {
            file.mkdirs();
        }
        File subfile = new File(strFilePath);
        if (!subfile.exists()) {
            try {
                boolean b = subfile.createNewFile();
                return b;
            } catch (IOException e) {
                e.printStackTrace();
            }
        } else {
            return true;
        }
        return false;
    }

    /**
     * 遍历文件夹下的文件
     *
     * @param file 地址
     */
    public static List<File> getFile(File file) {
        List<File> list = new ArrayList<>();
        File[] fileArray = file.listFiles();
        if (fileArray == null) {
            return null;
        } else {
            for (File f : fileArray) {
                if (f.isFile()) {
                    list.add(0, f);
                } else {
                    getFile(f);
                }
            }
        }
        return list;
    }

    /**
     * 删除文件
     *
     * @param filePath 文件地址
     * @return
     */
    public static boolean deleteFiles(String filePath) {
        List<File> files = getFile(new File(filePath));
        if (files.size() != 0) {
            for (int i = 0; i < files.size(); i++) {
                File file = files.get(i);
/** 如果是文件则删除 如果都删除可不必判断 */
                if (file.isFile()) {
                    file.delete();
                }
            }
        }
        return true;
    }

    /**
     * 向文件中添加内容
     *
     * @param strcontent 内容
     * @param filePath   地址
     * @param fileName   文件名
     */
    public static void writeToFile(String strcontent, String filePath, String fileName) {
//生成文件夹之后，再生成文件，不然会出错
        String strFilePath = filePath + fileName;
// 每次写入时，都换行写
        File subfile = new File(strFilePath);

        RandomAccessFile raf = null;
        try {
/** 构造函数 第二个是读写方式 */
            raf = new RandomAccessFile(subfile, "rw");
/** 将记录指针移动到该文件的最后 */
            raf.seek(subfile.length());
/** 向文件末尾追加内容 */
            raf.write(strcontent.getBytes());
            raf.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
     * 修改文件内容（覆盖或者添加）
     *
     * @param path    文件地址
     * @param content 覆盖内容
     * @param append  指定了写入的方式，是覆盖写还是追加写(true=追加)(false=覆盖)
     */
    public static void modifyFile(String path, String content, boolean append) {
        try {
            FileWriter fileWriter = new FileWriter(path, append);
            BufferedWriter writer = new BufferedWriter(fileWriter);
            writer.append(content);
            writer.flush();
            writer.close();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    /**
     * 读取文件内容
     *
     * @param filePath 地址
     * @param filename 名称
     * @return 返回内容
     */
    public static String getString(String filePath, String filename) {
        FileInputStream inputStream = null;
        try {
            inputStream = new FileInputStream(new File(filePath + filename));
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        }
        InputStreamReader inputStreamReader = null;
        try {
            inputStreamReader = new InputStreamReader(inputStream, "UTF-8");
        } catch (UnsupportedEncodingException e1) {
            e1.printStackTrace();
        }
        BufferedReader reader = new BufferedReader(inputStreamReader);
        StringBuffer sb = new StringBuffer("");
        String line;
        try {
            while ((line = reader.readLine()) != null) {
                sb.append(line);
                sb.append("\n");
            }
        } catch (IOException e) {
            e.printStackTrace();
        }
        return sb.toString();
    }

    /**
     * 重命名文件
     *
     * @param oldPath 原来的文件地址
     * @param newPath 新的文件地址
     */
    public static void renameFile(String oldPath, String newPath) {
        File oleFile = new File(oldPath);
        File newFile = new File(newPath);
//执行重命名
        oleFile.renameTo(newFile);
    }

    /**
     * 复制文件
     *
     * @param fromFile 要复制的文件目录
     * @param toFile   要粘贴的文件目录
     * @return 是否复制成功
     */
    public static boolean copy(String fromFile, String toFile) {
//要复制的文件目录
        File[] currentFiles;
        File root = new File(fromFile);
//如同判断SD卡是否存在或者文件是否存在
//如果不存在则 return出去
        if (!root.exists()) {
            return false;
        }
//如果存在则获取当前目录下的全部文件 填充数组
        currentFiles = root.listFiles();
//目标目录
        File targetDir = new File(toFile);
//创建目录
        if (!targetDir.exists()) {
            targetDir.mkdirs();
        }
//遍历要复制该目录下的全部文件
        for (int i = 0; i < currentFiles.length; i++) {
            if (currentFiles[i].isDirectory())//如果当前项为子目录 进行递归
            {
                copy(currentFiles[i].getPath() + "/", toFile + currentFiles[i].getName() + "/");
            } else//如果当前项为文件则进行文件拷贝
            {
                CopySdcardFile(currentFiles[i].getPath(), toFile + currentFiles[i].getName());
            }
        }
        return true;
    }

    //文件拷贝
//要复制的目录下的所有非子目录(文件夹)文件拷贝
    public static boolean CopySdcardFile(String fromFile, String toFile) {
        try {
            InputStream fosfrom = new FileInputStream(fromFile);
            OutputStream fosto = new FileOutputStream(toFile);
            byte bt[] = new byte[1024];
            int c;
            while ((c = fosfrom.read(bt)) > 0) {
                fosto.write(bt, 0, c);
            }
            fosfrom.close();
            fosto.close();
            return true;
        } catch (Exception ex) {
            return false;
        }
    }

    public static boolean deleteDirectory(File dir) {
        if (dir.isDirectory()) {
            String[] children = dir.list();
            for (int i = 0; i < children.length; i++) {
                boolean success = deleteDirectory(new File(dir, children[i]));
                if (!success) {
                    // 处理删除失败的情况
                }
            }
        }
        // 递归遍历完毕后，删除目录自身
        return dir.delete();
    }

    /**
     * 创建txt文件
     */
    public static boolean createFileConf(Activity activity) {
        File appDir = activity.getFilesDir();
        //  File appDir =Environment.getExternalStorageDirectory();
        String name = appDir + "/.lnd";

        File filePath = new File(name);
        if (!filePath.exists()) {
            Toast.makeText(activity, "暂无lnd文件，无法生成conf配置文件", Toast.LENGTH_SHORT).show();
            return false;
        }
        String appConf = name + "/lnd.conf";
        File filePathConf = new File(appConf);
        if (!filePathConf.exists()) {

            //  filePathConf.createNewFile();
            Tool tool = new Tool();
            String filePathcof = name + "/";
            String fileName1 = "lnd.conf";
            tool.writeTxtToFile(activity, ConStantUtil.ConfTxt(), filePathcof, fileName1);// 将字符串写入到文本文件中
            Toast.makeText(activity, "创建成功", Toast.LENGTH_SHORT).show();

            LogUntil.i( "CONF文件创建成功");
            return true;
        }
        return true;

    }

}
