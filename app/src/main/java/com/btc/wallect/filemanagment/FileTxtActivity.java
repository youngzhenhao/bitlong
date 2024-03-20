package com.btc.wallect.filemanagment;

import android.Manifest;
import android.app.Activity;
import android.app.AlertDialog;
import android.content.DialogInterface;
import android.content.pm.PackageManager;
import android.os.Build;
import android.os.Bundle;
import android.os.Environment;
import android.text.TextUtils;
import android.view.View;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;
import android.widget.Toast;

import androidx.annotation.NonNull;
import androidx.core.app.ActivityCompat;

import com.btc.wallect.R;
import com.btc.wallect.utils.ConStantUtil;
import com.btc.wallect.utils.SharedPreferencesHelperUtil;
import com.btc.wallect.utils.Tool;


/**
 * By abc
 */

public class FileTxtActivity extends Activity implements OnClickListener {

    private EditText mPathName;
    private EditText mFileName;
    private EditText mContent;
    private Button mTest;
    private Button mBtnRead;
    private TextView mShowTet;
    String filePath;
    String fileName1;
    private static final int REQUEST_READ_EXTERNAL_STORAGE = 1;
    private static final String[] PERMISSIONS_STORAGE = {Manifest.permission.READ_EXTERNAL_STORAGE, Manifest.permission.WRITE_EXTERNAL_STORAGE};
    //请求状态码
    private static final int REQUEST_EXTERNAL_STORAGE = 1;

    private TextView mTVfilePath;
    private EditText medTxt;

    private AlertDialog dialog;
    private boolean havePermission;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.newfile);
        findid();
       // checkPermission();
      //  hasStoragePermissions();


    }



    private void findid() {
        mPathName = (EditText) findViewById(R.id.et_pathName);
        mFileName = (EditText) findViewById(R.id.et_fileName);
        mContent = (EditText) findViewById(R.id.et_content);
        mTest = (Button) findViewById(R.id.btn_test);

        mBtnRead = (Button) findViewById(R.id.btn_read);
        mShowTet = (TextView) findViewById(R.id.Show_tet);
        mTest.setOnClickListener(this);
        mBtnRead.setOnClickListener(this);
    }

    private void getcontent() {
        //动态创建文件
        String pathName = mPathName.getText().toString();
        String fileName = mFileName.getText().toString();
        String content = mContent.getText().toString();
        if (TextUtils.isEmpty(pathName)) {
            Toast.makeText(this, "文件夹名不能为空", Toast.LENGTH_SHORT).show();
        } else if (TextUtils.isEmpty(fileName)) {
            Toast.makeText(this, "文件名不能为空", Toast.LENGTH_SHORT).show();
        } else {
            Tool tool = new Tool();
            filePath = Environment.getExternalStorageDirectory()
                    .getPath() + "/" + pathName + "/";
            fileName1 = fileName + ".txt";

            tool.writeTxtToFile(FileTxtActivity.this,content, filePath, fileName1);// 将字符串写入到文本文件中
            Toast.makeText(this, "创建成功", Toast.LENGTH_SHORT).show();
        }
    }

    @Override
    public void onClick(View view) {
        switch (view.getId()) {
            case R.id.btn_test:
                getcontent();
                break;
            case R.id.btn_read:
                Tool tool = new Tool();
              //  String path= SpUtils.readPath(FileTxtActivity.this);
                String path= SharedPreferencesHelperUtil.getInstance().getStringValue(ConStantUtil.WALLECT_ACCOUT,"");
                mShowTet.setText(tool.readTxt(path));

                break;
        }
    }

    private void checkPermission() {



        if (Build.VERSION.SDK_INT > Build.VERSION_CODES.M) {
            if (ActivityCompat.checkSelfPermission(this, Manifest.permission.WRITE_EXTERNAL_STORAGE) != PackageManager.PERMISSION_GRANTED) {
                //申请权限
                if (dialog != null) {
                    dialog.dismiss();
                    dialog = null;
                }
                dialog = new AlertDialog.Builder(this)
                        .setTitle("提示")//设置标题
                        .setMessage("请开启文件访问权限，否则无法正常使用应用！")
                        .setPositiveButton("确定", new DialogInterface.OnClickListener() {
                            @Override
                            public void onClick(DialogInterface dialog, int which) {
                                dialog.dismiss();
                                ActivityCompat.requestPermissions(FileTxtActivity.this, PERMISSIONS_STORAGE, REQUEST_EXTERNAL_STORAGE);

                            }
                        }).create();
                dialog.show();
            } else {
                havePermission = true;

            }
        } else {

            havePermission = true;

        }

    }


    @Override
    public void onRequestPermissionsResult(int requestCode, @NonNull String[] permissions, @NonNull int[] grantResults) { //申请文件读写权限
        super.onRequestPermissionsResult(requestCode, permissions, grantResults);

        switch (requestCode) {
            case REQUEST_EXTERNAL_STORAGE:
                if (grantResults.length > 0 && grantResults[0] == PackageManager.PERMISSION_GRANTED) {
                    havePermission = true;
                    Toast.makeText(this, "授权成功！", Toast.LENGTH_SHORT).show();

                } else {
                    havePermission = false;
                    Toast.makeText(this, "授权被拒绝！", Toast.LENGTH_SHORT).show();
                }
                break;
            case 33:

                Toast.makeText(this, "授权成功！", Toast.LENGTH_SHORT).show();
                break;
        }
    }

}
