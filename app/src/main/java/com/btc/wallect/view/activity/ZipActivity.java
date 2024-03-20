package com.btc.wallect.view.activity;

import android.app.Activity;
import android.app.AlertDialog;
import android.content.DialogInterface;
import android.os.Bundle;
import android.os.Environment;
import android.util.Log;
import android.view.View;
import android.widget.EditText;

import androidx.annotation.Nullable;


import com.btc.wallect.R;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.FileUtils;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.utils.ZipUtils;

import java.io.File;


public class ZipActivity extends Activity {
    private EditText mEtzipPath, mEtunZipPath;
    private AlertDialog   alertDialog;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_zip);
        initView();
    }

    public void initView() {
        mEtzipPath = findViewById(R.id.et_user_zipPath);
        mEtunZipPath = findViewById(R.id.et_unZipPath);
        File zipfile = Environment.getExternalStorageDirectory();
        String zipWalletfile = zipfile.getAbsolutePath();
        mEtzipPath.setText("压缩文件路径：" + zipWalletfile + "/wallet.zip");
        String unzipPath=getFilesDir().getPath();
        SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.WALLECT_UNZIP,unzipPath+"/.lnd");


        mEtunZipPath.setText("解压路径：" + SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.WALLECT_UNZIP,""));

        findViewById(R.id.btn_zip).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                File appDir = getFilesDir();
                String filePath = appDir + "/.lnd";
                File zipfile = Environment.getExternalStorageDirectory();
                String zipWalletfile = zipfile.getAbsolutePath();
                Log.e("LZ>>>", "压缩路径：" + filePath + " 压缩完成的路径：" + zipWalletfile + "/wallet.zip");
                try {
                    ZipUtils.compressFolder(ZipActivity.this, filePath, zipWalletfile + "/wallet.zip");
                } catch (Exception e) {
                    throw new RuntimeException(e);
                }
            }
        });
        findViewById(R.id.btn_unzip).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                File appDir = getFilesDir();
                String filePath = appDir + "/.lnd";
                File file=new File(filePath);
                if(file.exists()){
                    isDeleteFile();

                }else {
                    UNzip();
                }

            }
        });


        findViewById(R.id.btn_Testzip).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String pathFile = FileUtils.CreeateZipFile(ZipActivity.this);
//                SpUtils.SaveUnZipPath(ZipActivity.this, pathFile);
//                mEtunZipPath.setText("解压路径：" + pathFile);
            }
        });
    }
    public void isDeleteFile(){

           alertDialog=new AlertDialog.Builder(this)
                .setIcon(R.drawable.newlogo)
                .setTitle("提示")
                .setMessage("检测到文件还存在，是否删除现有文件夹？")
                .setNegativeButton("是", new DialogInterface.OnClickListener() {
                    public void onClick(DialogInterface dialogInterface, int i) {
                        File appDir = getFilesDir();
                        String filePath = appDir + "/.lnd";
                        File file=new File(filePath);
                        boolean result = FileUtils.deleteDirectory(file);
                        if (result) {
                      Log.e("LZ>>>>","删除成功");
                        } else {
                            Log.e("LZ>>>>","删除失败");

                        }
                        alertDialog.dismiss();//销毁对话框
                        UNzip();
                    }
                })
                .setPositiveButton("否", new DialogInterface.OnClickListener() {
                    public void onClick(DialogInterface dialogInterface, int i) {
                        alertDialog.dismiss();
                    }
                })
                .create();

        alertDialog.show();
    }
    public void UNzip(){
        File zipfile = Environment.getExternalStorageDirectory();
        String zipWalletfile = zipfile.getAbsolutePath();
        String unzipPath=getFilesDir().getPath();
        ZipUtils.Unzip(zipWalletfile + "/wallet.zip",unzipPath+"/");
        mEtunZipPath.setText("解压路径：" +SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.WALLECT_UNZIP,""));
    }
}
