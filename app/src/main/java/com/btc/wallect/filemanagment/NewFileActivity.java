package com.btc.wallect.filemanagment;

import android.Manifest;
import android.app.Activity;
import android.app.AlertDialog;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.os.Build;
import android.os.Bundle;
import android.view.View;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.core.app.ActivityCompat;


import com.btc.wallect.R;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.FileUtils;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.utils.threadutil.ThreadPool;
import com.btc.wallect.utils.threadutil.ThreadPoolUtils;
import com.btc.wallect.view.activity.ZipActivity;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

import api.Api;


public class NewFileActivity extends Activity {
    private static final int REQUEST_READ_EXTERNAL_STORAGE = 1;
    private static final String[] PERMISSIONS_STORAGE = {Manifest.permission.READ_EXTERNAL_STORAGE, Manifest.permission.WRITE_EXTERNAL_STORAGE};
    //请求状态码
    private static final int REQUEST_EXTERNAL_STORAGE = 1;

    private TextView mTVfilePath;
    private EditText medTxt, mEdTxtfileName, mEdTxtfileNameTxt;
    private String filePath;
    private AlertDialog dialog;
    private boolean havePermission;

    @Override
    protected void onCreate(@Nullable Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.act_newfile);
        initView();
//        final Intent intent = new Intent(this, FirstService.class);
//        startService(intent);
        requestStoragePermissions();
//        startService(new Intent(this, LocalService.class));
//        startService(new Intent(this, RemoteService.class));
    }

    public void initView() {
        mTVfilePath = findViewById(R.id.tv_filePath);
        medTxt = findViewById(R.id.ed_Txt);
        mEdTxtfileName = findViewById(R.id.ed_Txt_fileName);
        mEdTxtfileNameTxt = findViewById(R.id.ed_Txt_fileNameTxt);

        findViewById(R.id.btn_createFile).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String pathFile = FileUtils.CreeateFileTest(NewFileActivity.this);
                //  SpUtils.SavePath(NewFileActivity.this, pathFile);

            }
        });
        findViewById(R.id.btn_createFileTxt).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                //  String txt = "dirpath=" + SpUtils.readPath(NewFileActivity.this);
                String newTXT = ConStantUtil.ConfigTxt();

                String txt = "dirpath=" + getFilesDir() + newTXT;
                String TxtfileName = "config";

                FileUtils.CreeateFileTxt(NewFileActivity.this, TxtfileName, txt);


            }
        });
        findViewById(R.id.btn_readFileTxt).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
//                String txt = ReadFile();
//                medTxt.setText(txt);
                Test();

            }
        });
        findViewById(R.id.btn_modflyFileTxt).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                String edtxt = medTxt.getText().toString();

            }
        });
        findViewById(R.id.btn_deleteFileTxt).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

            }
        });
        findViewById(R.id.btn_compressFile).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent in = new Intent(NewFileActivity.this, ZipActivity.class);
                startActivity(in);


            }
        });
        findViewById(R.id.btn_moueFile).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                // String pathFile = SpUtils.readPath(NewFileActivity.this);

                String pathFile = FileUtils.CreeateFile(NewFileActivity.this);
                SharedPreferencesHelperUtil.getInstance().putStringValue(ConStantUtil.WALLECT_ACCOUT, pathFile);
            }
        });
        findViewById(R.id.btn_startaar).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {

                ThreadPoolUtils.execute(new Runnable() {
                    @Override
                    public void run() {
                        findViewById(R.id.btn_startaar).setClickable(false);
                         Api.starLnd();
                    }
                });

            }

        });
        findViewById(R.id.btn_testLndFile).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
//                String txt = "dirpath=" + getFilesDir()+"111111";
//                String TxtfileName = "TESTWWWW";
//                FileUtils.CreeateFileTxtTest(NewFileActivity.this, TxtfileName, txt);

            }
        });
        findViewById(R.id.btn_initwallet).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {


                boolean b = FileUtils.createFileConf(NewFileActivity.this);
                if (b) {

                    ThreadPool t = new ThreadPool();
                    t.execute(new Runnable() {
                        @Override
                        public void run() {
//                            boolean wa = Api.initwallet("12345678");
//                            LogUtil.i("lz>>>>in", wa + "");
                        }
                    });

                }
            }
        });
        findViewById(R.id.btn_UncompressFile).setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
//                Intent inte = new Intent(NewFileActivity.this, WebViewActivity.class);
//                startActivity(inte);


            }
        });
    }

    @Override
    public void onRequestPermissionsResult(int requestCode, @NonNull String[] permissions, @NonNull int[] grantResults) { //申请文件读写权限
        super.onRequestPermissionsResult(requestCode, permissions, grantResults);
        switch (requestCode) {
            case REQUEST_EXTERNAL_STORAGE:
                if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    havePermission = true;
                    //  Toast.makeText(this, "授权成功！", Toast.LENGTH_SHORT).show();

                } else {
                    havePermission = false;
                    Toast.makeText(this, "授权被拒绝！", Toast.LENGTH_SHORT).show();
                }
                break;
            case 33:
                Toast.makeText(this, "已授权！", Toast.LENGTH_SHORT).show();
                break;
        }
    }


    @Override
    protected void onResume() {
        super.onResume();


    }


    private void requestStoragePermissions() {
        String[] permissions;
        if (Build.VERSION.SDK_INT < Build.VERSION_CODES.TIRAMISU) {
            permissions = new String[]{Manifest.permission.READ_EXTERNAL_STORAGE, Manifest.permission.WRITE_EXTERNAL_STORAGE};
        } else {
            permissions = new String[]{Manifest.permission.READ_MEDIA_AUDIO, Manifest.permission.READ_MEDIA_IMAGES, Manifest.permission.READ_MEDIA_VIDEO, Manifest.permission.WRITE_EXTERNAL_STORAGE};
        }
        ActivityCompat.requestPermissions((Activity) this, permissions, REQUEST_EXTERNAL_STORAGE);
    }

    private void Test() {
//        String printTxtPath =Environment.getExternalStorageDirectory().getPath();
//        Log.e("lz>>>>>path:",printTxtPath+"");
//        String path=SpUtils.readPath(NewFileActivity.this);
//        String txt= Api.readFile(path);
        String path = getFilesDir().getPath() + "/NewFolderBit";

//        boolean isTXT = Api.configureFile(path);
//        Log.e("lz>>>>>path", path + ":" + isTXT);
//        Toast.makeText(this, "路径：" + path + ":返回状态：" + isTXT, Toast.LENGTH_SHORT).show();
//        medTxt.setText("路径：" + path + ":返回状态：" + isTXT);

        try {
            // 获取文件路径
            String filePath = getFilesDir().getPath() + "/NewFolderBit/config.txt";

            // 创建文件对象
            File file = new File(filePath);

            // 创建文件输入流
            FileInputStream fis = new FileInputStream(file);

            // 创建一个字节数组，用于存放文件内容
            byte[] buffer = new byte[fis.available()];

            // 读取文件内容到字节数组
            fis.read(buffer);

            // 关闭输入流
            fis.close();

            // 将字节数组转换为字符串
            String content = new String(buffer);
            // 输出文件内容
            System.out.println("LZ>>>>>txt读取结果:" + content);
        } catch (IOException e) {
            e.printStackTrace();
        }
    }


}
