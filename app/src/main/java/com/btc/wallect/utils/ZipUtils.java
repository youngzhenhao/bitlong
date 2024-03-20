package com.btc.wallect.utils;

import android.app.Activity;
import android.util.Log;
import android.widget.Toast;

import java.io.BufferedInputStream;
import java.io.BufferedOutputStream;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileOutputStream;
import java.io.IOException;
import java.io.InputStream;
import java.util.zip.ZipEntry;
import java.util.zip.ZipInputStream;
import java.util.zip.ZipOutputStream;

public class ZipUtils {
    static String TAG = "lZ>>";




    public static void compressFolder(Activity activity, String sourcePath, String zipFile) throws IOException {
        File folder = new File(sourcePath);
        File zipPath = new File(zipFile);
        if (zipPath.exists()) {
            zipPath.delete();
        }

        // 创建输出流并指定要生成的ZIP文件路径
        try (ZipOutputStream zos = new ZipOutputStream(new FileOutputStream(zipFile))) {
            addFilesToZip(folder, folder.getName(), zos);

            System.out.println("文件夹已成功压缩为 " + zipFile);
            Toast.makeText(activity, "压缩已完成", Toast.LENGTH_SHORT).show();
        } catch (IOException e) {
            throw new RuntimeException("无法将文件夹压缩到ZIP文件中", e);
        }
    }

    private static void addFilesToZip(File file, String parentName, ZipOutputStream zos) throws IOException {
        if (!file.exists()) return;

        for (File child : file.listFiles()) {
            if (child.isDirectory()) {
                addFilesToZip(child, parentName + "/" + child.getName(), zos);
            } else {
                byte[] buffer = new byte[1024];

                // 获取当前文件相对于源文件夹的路径名称
                String entryName = parentName + "/" + child.getName();

                // 添加新条目到ZIP文件中
                zos.putNextEntry(new ZipEntry(entryName));

                // 读取文件内容并写入ZIP文件
                try (InputStream is = new FileInputStream(child)) {
                    int length;
                    while ((length = is.read(buffer)) > 0) {
                        zos.write(buffer, 0, length);
                    }
                } finally {
                    zos.closeEntry();
                }
            }
        }
    }

    public static void Unzip(String zipFile, String targetDir) {
        Log.e("lz>>","zipFile:"+zipFile+" targetDir:"+targetDir);
        int BUFFER = 4096; //这里缓冲区我们使用4KB，
        String strEntry; //保存每个zip的条目名称

        try {
            BufferedOutputStream dest = null; //缓冲输出流
            FileInputStream fis = new FileInputStream(zipFile);
            ZipInputStream zis = new ZipInputStream(new BufferedInputStream(fis));
            ZipEntry entry; //每个zip条目的实例

            while ((entry = zis.getNextEntry()) != null) {

                try {
                    Log.i("Unzip: ","="+ entry);
                    int count;
                    byte data[] = new byte[BUFFER];
                    strEntry = entry.getName();

                    File entryFile = new File(targetDir + strEntry);
                    File entryDir = new File(entryFile.getParent());
                    if (!entryDir.exists()) {
                        entryDir.mkdirs();
                    }

                    FileOutputStream fos = new FileOutputStream(entryFile);
                    dest = new BufferedOutputStream(fos, BUFFER);
                    while ((count = zis.read(data, 0, BUFFER)) != -1) {
                        dest.write(data, 0, count);
                    }
                    dest.flush();
                    dest.close();
                } catch (Exception ex) {
                    ex.printStackTrace();
                }
            }
            zis.close();
            Log.i("Unzip: ","="+ "结束");
        } catch (Exception cwj) {
            cwj.printStackTrace();
        }
    }
}
